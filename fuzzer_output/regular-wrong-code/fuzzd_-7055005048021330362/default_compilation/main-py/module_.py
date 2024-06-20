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
    def fm4(p0, p1, globalState):
        if (104) == (len(_dafny.Set({-395}))):
            return _dafny.CodePoint('m')
        elif True:
            return _dafny.CodePoint('p')

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}) for d_0_i1_ in range(default__.abs(412))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}), _dafny.Map({False: False}), _dafny.Map({True: not(False)})]))) for d_1_i0_ in range(default__.abs(33))])

    @staticmethod
    def fm6(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-139, 404):
                d_2_v0_: int = compr_0_
                if ((-139) <= (d_2_v0_)) and ((d_2_v0_) < (404)):
                    def iife1_():
                        coll1_ = _dafny.Map()
                        compr_1_: int
                        for compr_1_ in (_dafny.SeqWithoutIsStrInference([741])).Elements:
                            d_3_v1_: int = compr_1_
                            if (d_3_v1_) in (_dafny.SeqWithoutIsStrInference([741])):
                                coll1_[default__.safeDivisionInt(d_3_v1_, 293)] = not(True)
                        return _dafny.Map(coll1_)
                    coll0_[default__.safeModuloInt(d_2_v0_, 153)] = iife1_()

            return _dafny.Map(coll0_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.Map({22: 203})).keys.Elements:
                d_4_v2_: int = compr_2_
                if (d_4_v2_) in (_dafny.Map({22: 203})):
                    coll2_[(d_4_v2_) - (len(_dafny.Map({True: True})))] = False
            return _dafny.Map(coll2_)
        return default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nasbdl")))) * ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('k')])))), len((iife0_()
        ) | (_dafny.Map({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))))})): iife2_()
        }))))

    @staticmethod
    def fm9(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, False]) for d_5_i0_ in range(default__.abs(-545))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, False])]))).Elements:
                d_6_v0_: _dafny.Seq = compr_3_
                if (d_6_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, False]) for d_5_i0_ in range(default__.abs(-545))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, False])]))):
                    coll3_[d_6_v0_] = not((not(False)) == (True))
            return _dafny.Map(coll3_)
        return iife3_()
        

    @staticmethod
    def fm10(globalState):
        return D1_DC3(_dafny.Set({_dafny.Set({_dafny.CodePoint('x'), _dafny.CodePoint('e')})}))

    @staticmethod
    def fm11(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_7_i0_ in range(default__.abs(555))])

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('h')])) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')]))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlwcc"))])) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dulq")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouob")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wntsnxr"))])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oikjsmub"))])))

    @staticmethod
    def fm14(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-141, 493):
                d_8_v0_: int = compr_4_
                if ((-141) <= (d_8_v0_)) and ((d_8_v0_) < (493)):
                    coll4_ = coll4_.union(_dafny.Set([(d_8_v0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([True])])))]))
            return _dafny.Set(coll4_)
        return iife4_()
        

    @staticmethod
    def fm15(globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: str
            for compr_5_ in (_dafny.Set({_dafny.CodePoint('r')})).Elements:
                d_9_v0_: str = compr_5_
                if (d_9_v0_) in (_dafny.Set({_dafny.CodePoint('r')})):
                    coll5_ = coll5_.union(_dafny.Set([d_9_v0_]))
            return _dafny.Set(coll5_)
        return (_dafny.Set({_dafny.CodePoint('a'), _dafny.CodePoint('a'), _dafny.CodePoint('r'), _dafny.CodePoint('a')})) | ((_dafny.Set({_dafny.CodePoint('j')})) - (iife5_()
        ))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False)])), (_dafny.MultiSet([_dafny.Set({not(False)})])).cardinality])])

    @staticmethod
    def fm17(p0, p1, globalState):
        return (_dafny.MultiSet([not((D9_DC20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lt")), False, 130)).cf38), False, True, False])) - ((_dafny.MultiSet([not(True)])) | ((D3_DC7(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True])))).cf12))

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkakqfc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhejkofsg"))): len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olxnmwt")) if False else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_10_i0_ in range(default__.abs(461))])))})

    @staticmethod
    def fm19(globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: _dafny.MultiSet
            for compr_6_ in ((_dafny.Map({_dafny.MultiSet([-997]): D3_DC7(_dafny.MultiSet([not(True), True]))})) | (_dafny.Map({_dafny.MultiSet([-629]): D3_DC7(_dafny.MultiSet([True]))}))).keys.Elements:
                d_11_v0_: _dafny.MultiSet = compr_6_
                if (d_11_v0_) in ((_dafny.Map({_dafny.MultiSet([-997]): D3_DC7(_dafny.MultiSet([not(True), True]))})) | (_dafny.Map({_dafny.MultiSet([-629]): D3_DC7(_dafny.MultiSet([True]))}))):
                    coll6_[d_11_v0_] = ((0) - (-516)) <= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsjog"))))
            return _dafny.Map(coll6_)
        return iife6_()
        

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        d_12_v0_: _dafny.Map
        d_12_v0_ = _dafny.Map({-251: p1})
        d_13_v1_: int
        d_13_v1_ = -477
        (globalState).f3 = (len(d_12_v0_)) - (d_13_v1_)
        d_14_v2_: _dafny.Seq
        d_14_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqmiaea"))
        d_15_v3_: _dafny.Map
        d_15_v3_ = _dafny.Map({d_14_v2_: d_13_v1_})
        d_16_v4_: _dafny.Seq
        d_16_v4_ = _dafny.SeqWithoutIsStrInference([382, d_13_v1_])
        d_15_v3_ = ((_dafny.Map({d_14_v2_: d_13_v1_})) | (d_15_v3_)) | ((default__.fm18(878, d_16_v4_, d_13_v1_, d_13_v1_, globalState)) | (d_15_v3_))
        d_17_v5_: _dafny.MultiSet
        d_17_v5_ = _dafny.MultiSet([(d_16_v4_)[default__.safeIndex(-426, len(d_16_v4_))]])
        d_18_v6_: str
        d_18_v6_ = _dafny.CodePoint('t')
        hi0_ = (len(d_14_v2_) if p1 else d_13_v1_)
        for d_19_i0_ in range((0) - ((len(_dafny.Map({(d_17_v5_).cardinality: d_18_v6_}))) * (582)), hi0_):
            if p1:
                index0_ = default__.safeIndex(666, (p0).length(0))
                (p0)[index0_] = not((p1) == (p1))
                index1_ = default__.safeIndex(666, (p0).length(0))
                (p0)[index1_] = (p1) or (not((len(default__.fm19(globalState))) != (d_19_i0_)))
                d_20_v7_: _dafny.Set
                d_20_v7_ = _dafny.Set({d_14_v2_, d_14_v2_, d_14_v2_, d_14_v2_})
                d_21_v8_: C0
                nw0_ = C0()
                nw0_.ctor__(d_20_v7_)
                d_21_v8_ = nw0_
                d_22_v9_: _dafny.Seq
                d_22_v9_ = _dafny.SeqWithoutIsStrInference([d_21_v8_])
                d_23_v10_: _dafny.MultiSet
                d_23_v10_ = _dafny.MultiSet([d_21_v8_, d_21_v8_])
                d_24_v11_: C3
                nw1_ = C3()
                nw1_.ctor__((p0)[default__.safeIndex(666, (p0).length(0))], p1, d_14_v2_, p1)
                d_24_v11_ = nw1_
                d_25_v12_: D7
                d_25_v12_ = D7_DC16(p1, False, d_24_v11_, (d_24_v11_).f27)
                d_26_v13_: _dafny.Map
                d_26_v13_ = _dafny.Map({d_25_v12_: d_18_v6_})
                d_27_v14_: _dafny.MultiSet
                d_27_v14_ = _dafny.MultiSet([(d_24_v11_).f27])
                d_28_v15_: _dafny.Array
                nw2_ = _dafny.Array(False, 17)
                d_28_v15_ = nw2_
                d_29_v16_: _dafny.Map
                d_29_v16_ = _dafny.Map({(d_16_v4_)[default__.safeIndex(d_13_v1_, len(d_16_v4_))]: d_28_v15_})
                d_30_v17_: _dafny.Map
                d_30_v17_ = d_29_v16_
                d_31_v18_: _dafny.Map
                d_31_v18_ = _dafny.Map({d_13_v1_: d_13_v1_})
                d_32_v19_: _dafny.Array
                nw3_ = _dafny.Array(None, 20)
                nw3_[int(0)] = (d_13_v1_) + ((0) - (d_19_i0_))
                nw3_[int(1)] = d_19_i0_
                nw3_[int(2)] = (d_13_v1_) + (d_19_i0_)
                nw3_[int(3)] = d_19_i0_
                nw3_[int(4)] = default__.fm6(globalState)
                nw3_[int(5)] = len((d_22_v9_) + (d_22_v9_))
                nw3_[int(6)] = ((d_23_v10_)[d_21_v8_] if (d_21_v8_) in (d_23_v10_) else d_13_v1_)
                nw3_[int(7)] = d_19_i0_
                nw3_[int(8)] = d_13_v1_
                nw3_[int(9)] = (d_13_v1_) - (len(d_26_v13_))
                nw3_[int(10)] = 628
                nw3_[int(11)] = ((d_17_v5_)[d_13_v1_] if (d_13_v1_) in (d_17_v5_) else (d_27_v14_).cardinality)
                nw3_[int(12)] = len((d_29_v16_) | ((d_30_v17_)))
                nw3_[int(13)] = d_19_i0_
                nw3_[int(14)] = d_13_v1_
                nw3_[int(15)] = d_13_v1_
                nw3_[int(16)] = default__.safeDivisionInt(d_13_v1_, d_19_i0_)
                nw3_[int(17)] = ((d_31_v18_)[d_13_v1_] if (d_13_v1_) in (d_31_v18_) else d_13_v1_)
                nw3_[int(18)] = (d_13_v1_) * (len(d_21_v8_.f26))
                nw3_[int(19)] = d_19_i0_
                d_32_v19_ = nw3_
                index2_ = default__.safeIndex(378, (d_32_v19_).length(0))
                (d_32_v19_)[index2_] = d_19_i0_
                d_33_v20_: _dafny.Array
                nw4_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                d_33_v20_ = nw4_
                index3_ = default__.safeIndex(84, (d_33_v20_).length(0))
                (d_33_v20_)[index3_] = d_18_v6_
                index4_ = default__.safeIndex(378, (d_32_v19_).length(0))
                index5_ = default__.safeIndex(84, (d_33_v20_).length(0))
                index6_ = default__.safeIndex(666, (p0).length(0))
                index7_ = default__.safeIndex(666, (p0).length(0))
                rhs0_ = (d_19_i0_) + ((d_13_v1_ if d_24_v11_.f28 else 468))
                rhs1_ = default__.safeModuloInt(d_19_i0_, len(d_14_v2_))
                rhs2_ = d_18_v6_
                rhs3_ = not((((d_14_v2_) + (d_14_v2_)).set(default__.safeIndex(d_13_v1_, len((d_14_v2_) + (d_14_v2_))), d_18_v6_)) <= (d_14_v2_))
                rhs4_ = (d_24_v11_).f27
                lhs0_ = d_32_v19_
                lhs1_ = default__.safeIndex(378, (d_32_v19_).length(0))
                lhs2_ = globalState
                lhs3_ = d_33_v20_
                lhs4_ = default__.safeIndex(84, (d_33_v20_).length(0))
                lhs5_ = p0
                lhs6_ = default__.safeIndex(666, (p0).length(0))
                lhs7_ = p0
                lhs8_ = default__.safeIndex(666, (p0).length(0))
                lhs0_[lhs1_] = rhs0_
                lhs2_.f20 = rhs1_
                lhs3_[lhs4_] = rhs2_
                lhs5_[lhs6_] = rhs3_
                lhs7_[lhs8_] = rhs4_
                d_34_v21_: T0
                nw5_ = C3()
                nw5_.ctor__(p1, True, d_14_v2_, (p0)[default__.safeIndex(666, (p0).length(0))])
                d_34_v21_ = nw5_
                d_35_v22_: _dafny.Map
                d_35_v22_ = _dafny.Map({d_34_v21_: p1})
                d_36_v23_: _dafny.Seq
                d_36_v23_ = _dafny.SeqWithoutIsStrInference([d_34_v21_.f22])
                d_37_v24_: _dafny.Map
                d_37_v24_ = _dafny.Map({d_24_v11_.f28: d_13_v1_})
                d_38_v25_: D1
                d_38_v25_ = D1_DC5(d_34_v21_.f22, d_37_v24_, d_13_v1_, (_dafny.MultiSet([379])).set((d_32_v19_)[default__.safeIndex(378, (d_32_v19_).length(0))], default__.abs(len(_dafny.SeqWithoutIsStrInference([(d_33_v20_)[default__.safeIndex(84, (d_33_v20_).length(0))] for d_39_i1_ in range(default__.abs(627))])))))
                d_40_v26_: D3
                d_40_v26_ = D3_DC8(len(d_36_v23_), _dafny.Map({d_27_v14_: len(d_36_v23_)}), d_38_v25_, p1)
                d_41_v27_: D1
                d_41_v27_ = D1_DC5((d_40_v26_).cf16, d_37_v24_, (0) - (d_13_v1_), _dafny.MultiSet([d_19_i0_, (d_32_v19_)[default__.safeIndex(378, (d_32_v19_).length(0))]]))
                d_42_v28_: _dafny.Seq
                d_42_v28_ = _dafny.SeqWithoutIsStrInference([d_16_v4_])
                d_43_v29_: T2
                nw6_ = C2()
                nw6_.ctor__(d_42_v28_, _dafny.SeqWithoutIsStrInference([(d_33_v20_)[default__.safeIndex(84, (d_33_v20_).length(0))] for d_44_i2_ in range(default__.abs(604))]), d_24_v11_.f28)
                d_43_v29_ = nw6_
                d_45_v30_: _dafny.Seq
                d_45_v30_ = _dafny.SeqWithoutIsStrInference([d_43_v29_, d_43_v29_])
                d_46_v31_: _dafny.Map
                d_46_v31_ = _dafny.Map({(d_35_v22_) != (d_35_v22_): (d_45_v30_ if (d_41_v27_).cf7 else d_45_v30_)})
                d_46_v31_ = (d_46_v31_).set(not(d_34_v21_.f22), (d_45_v30_) + (d_45_v30_))
                d_43_v29_ = d_43_v29_
                d_47_v32_: _dafny.Set
                d_47_v32_ = _dafny.Set({d_43_v29_})
                d_47_v32_ = (d_47_v32_) | (d_47_v32_)
            elif True:
                d_48_v33_: _dafny.MultiSet
                d_48_v33_ = _dafny.MultiSet([p1, False])
                d_49_v34_: _dafny.MultiSet
                d_49_v34_ = _dafny.MultiSet([d_48_v33_, d_48_v33_, d_48_v33_, d_48_v33_, d_48_v33_])
                d_50_v35_: _dafny.Seq
                d_50_v35_ = _dafny.SeqWithoutIsStrInference([default__.fm12(d_13_v1_, default__.fm17(d_13_v1_, default__.fm12(d_13_v1_, d_48_v33_, p1, d_49_v34_, globalState), globalState), p1, d_49_v34_, globalState)])
                r0 = (d_50_v35_)[default__.safeIndex(default__.fm6(globalState), len(d_50_v35_))]
                r0 = (d_14_v2_) == (d_14_v2_)
                d_51_v36_: _dafny.Array
                def lambda0_(d_52_v33_):
                    def lambda1_(d_53_i3_):
                        return d_52_v33_

                    return lambda1_

                init0_ = lambda0_(d_48_v33_)
                nw7_ = _dafny.Array(None, 12)
                for i0_0_ in range(nw7_.length(0)):
                    nw7_[i0_0_] = init0_(i0_0_)
                d_51_v36_ = nw7_
                d_54_v37_: C1
                nw8_ = C1()
                nw8_.ctor__(d_51_v36_, p1)
                d_54_v37_ = nw8_
                d_54_v37_ = d_54_v37_
                d_55_v38_: _dafny.Set
                d_55_v38_ = _dafny.Set({(d_54_v37_).f25})
                d_56_v39_: _dafny.Set
                d_56_v39_ = _dafny.Set({d_55_v38_, _dafny.Set({True, (d_54_v37_).f25})})
                d_57_v40_: _dafny.Map
                d_57_v40_ = _dafny.Map({(D9_DC18(d_56_v39_)).cf31: p1})
                d_57_v40_ = (d_57_v40_).set(d_56_v39_, p1)
                d_58_v41_: _dafny.Map
                d_58_v41_ = _dafny.Map({d_19_i0_: d_18_v6_})
                d_59_v42_: _dafny.Seq
                d_59_v42_ = _dafny.SeqWithoutIsStrInference([d_48_v33_, d_48_v33_])
                r0 = (((d_58_v41_)[((d_59_v42_)[default__.safeIndex(d_13_v1_, len(d_59_v42_))]).cardinality] if (((d_59_v42_)[default__.safeIndex(d_13_v1_, len(d_59_v42_))]).cardinality) in (d_58_v41_) else d_18_v6_)) in (d_14_v2_)
            d_60_v43_: _dafny.Map
            d_60_v43_ = _dafny.Map({p0: d_13_v1_})
            d_13_v1_ = len(d_60_v43_)
            d_13_v1_ = -286
            index8_ = default__.safeIndex(286, (p0).length(0))
            (p0)[index8_] = p1
            d_61_v44_: _dafny.Array
            nw9_ = _dafny.Array(False, 10)
            d_61_v44_ = nw9_
            d_62_v45_: _dafny.Seq
            d_62_v45_ = _dafny.SeqWithoutIsStrInference([p0])
            index9_ = default__.safeIndex(286, (p0).length(0))
            rhs5_ = ((d_13_v1_) - (d_19_i0_)) in (d_16_v4_)
            rhs6_ = (p0 if not(True) else (d_62_v45_)[default__.safeIndex(513, len(d_62_v45_))])
            lhs9_ = p0
            lhs10_ = default__.safeIndex(286, (p0).length(0))
            lhs9_[lhs10_] = rhs5_
            d_61_v44_ = rhs6_
        d_63_v46_: _dafny.MultiSet
        d_63_v46_ = _dafny.MultiSet([p1])
        d_64_v47_: _dafny.Map
        d_64_v47_ = _dafny.Map({d_63_v46_: default__.fm6(globalState)})
        d_65_v48_: D1
        d_65_v48_ = D1_DC5(p1, _dafny.Map({p1: (0) - (d_13_v1_)}), d_13_v1_, d_17_v5_)
        pat_let_tv0_ = d_16_v4_
        pat_let_tv1_ = d_16_v4_
        pat_let_tv2_ = d_16_v4_
        def lambda2_(source0_):
            if source0_.is_DC8:
                d_66___mcc_h0_ = source0_.cf13
                d_67___mcc_h1_ = source0_.cf14
                d_68___mcc_h2_ = source0_.cf15
                d_69___mcc_h3_ = source0_.cf16
                d_70_cf16_ = d_69___mcc_h3_
                d_71_cf15_ = d_68___mcc_h2_
                d_72_cf14_ = d_67___mcc_h1_
                d_73_cf13_ = d_66___mcc_h0_
                return pat_let_tv0_
            elif source0_.is_DC9:
                d_74___mcc_h4_ = source0_.cf17
                d_75___mcc_h5_ = source0_.cf18
                d_76___mcc_h6_ = source0_.cf19
                d_77_cf19_ = d_76___mcc_h6_
                d_78_cf18_ = d_75___mcc_h5_
                d_79_cf17_ = d_74___mcc_h4_
                return pat_let_tv1_
            elif True:
                d_80___mcc_h7_ = source0_.cf12
                d_81_cf12_ = d_80___mcc_h7_
                return pat_let_tv2_

        rhs7_ = lambda2_(D3_DC8(d_13_v1_, d_64_v47_, d_65_v48_, not(p1)))
        rhs8_ = p1
        rhs9_ = p1
        d_16_v4_ = rhs7_
        r0 = rhs8_
        r0 = rhs9_
        d_82_v49_: _dafny.Set
        d_82_v49_ = _dafny.Set({len(d_12_v0_)})
        d_83_v50_: _dafny.Set
        d_83_v50_ = d_82_v49_
        d_84_v51_: _dafny.Seq
        d_84_v51_ = _dafny.SeqWithoutIsStrInference([(d_16_v4_).set(default__.safeIndex(d_13_v1_, len(d_16_v4_)), d_13_v1_), d_16_v4_, d_16_v4_])
        d_85_v52_: T1
        nw10_ = C4()
        nw10_.ctor__(d_83_v50_, p1, d_84_v51_, d_14_v2_, (len(d_82_v49_)) == (d_13_v1_))
        d_85_v52_ = nw10_
        d_86_v53_: _dafny.Set
        d_86_v53_ = _dafny.Set({False})
        d_87_v54_: _dafny.Map
        d_87_v54_ = _dafny.Map({d_86_v53_: p1})
        d_88_v56_: D9
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: _dafny.Set
            for compr_7_ in (d_87_v54_).keys.Elements:
                d_89_v55_: _dafny.Set = compr_7_
                if (d_89_v55_) in (d_87_v54_):
                    coll7_ = coll7_.union(_dafny.Set([d_89_v55_]))
            return _dafny.Set(coll7_)
        d_88_v56_ = D9_DC18(iife7_()
)
        d_90_v57_: _dafny.Map
        d_90_v57_ = _dafny.Map({True: p1})
        d_91_v58_: D9
        d_91_v58_ = D9_DC20(d_85_v52_.f21, p1, len(d_90_v57_))
        d_92_v59_: D9
        d_92_v59_ = D9_DC20((d_91_v58_).cf37, True, d_13_v1_)
        pat_let_tv3_ = d_13_v1_
        pat_let_tv4_ = d_65_v48_
        pat_let_tv5_ = d_13_v1_
        pat_let_tv6_ = d_13_v1_
        pat_let_tv7_ = d_63_v46_
        pat_let_tv8_ = d_63_v46_
        pat_let_tv9_ = p1
        pat_let_tv10_ = p1
        pat_let_tv11_ = p1
        pat_let_tv12_ = d_63_v46_
        pat_let_tv13_ = d_63_v46_
        def lambda3_(source1_):
            if source1_.is_DC19:
                d_93___mcc_h8_ = source1_.cf32
                d_94___mcc_h9_ = source1_.cf33
                d_95___mcc_h10_ = source1_.cf34
                d_96___mcc_h11_ = source1_.cf35
                d_97___mcc_h12_ = source1_.cf36
                d_98_cf36_ = d_97___mcc_h12_
                d_99_cf35_ = d_96___mcc_h11_
                d_100_cf34_ = d_95___mcc_h10_
                d_101_cf33_ = d_94___mcc_h9_
                d_102_cf32_ = d_93___mcc_h8_
                return pat_let_tv3_
            elif source1_.is_DC20:
                d_103___mcc_h13_ = source1_.cf37
                d_104___mcc_h14_ = source1_.cf38
                d_105___mcc_h15_ = source1_.cf39
                d_106_cf39_ = d_105___mcc_h15_
                d_107_cf38_ = d_104___mcc_h14_
                d_108_cf37_ = d_103___mcc_h13_
                return len(((pat_let_tv4_).cf8) | (_dafny.Map({d_107_cf38_: pat_let_tv5_})))
            elif source1_.is_DC21:
                d_109___mcc_h16_ = source1_.cf40
                d_110___mcc_h17_ = source1_.cf41
                d_111_cf41_ = d_110___mcc_h17_
                d_112_cf40_ = d_109___mcc_h16_
                return 995
            elif True:
                d_113___mcc_h18_ = source1_.cf31
                d_114_cf31_ = d_113___mcc_h18_
                return pat_let_tv6_

        def lambda4_(source2_):
            if source2_.is_DC19:
                d_115___mcc_h19_ = source2_.cf32
                d_116___mcc_h20_ = source2_.cf33
                d_117___mcc_h21_ = source2_.cf34
                d_118___mcc_h22_ = source2_.cf35
                d_119___mcc_h23_ = source2_.cf36
                d_120_cf36_ = d_119___mcc_h23_
                d_121_cf35_ = d_118___mcc_h22_
                d_122_cf34_ = d_117___mcc_h21_
                d_123_cf33_ = d_116___mcc_h20_
                d_124_cf32_ = d_115___mcc_h19_
                return pat_let_tv7_
            elif source2_.is_DC20:
                d_125___mcc_h24_ = source2_.cf37
                d_126___mcc_h25_ = source2_.cf38
                d_127___mcc_h26_ = source2_.cf39
                d_128_cf39_ = d_127___mcc_h26_
                d_129_cf38_ = d_126___mcc_h25_
                d_130_cf37_ = d_125___mcc_h24_
                return pat_let_tv8_
            elif source2_.is_DC21:
                d_131___mcc_h27_ = source2_.cf40
                d_132___mcc_h28_ = source2_.cf41
                d_133_cf41_ = d_132___mcc_h28_
                d_134_cf40_ = d_131___mcc_h27_
                return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([pat_let_tv9_, pat_let_tv10_])) + (_dafny.SeqWithoutIsStrInference([d_133_cf41_, not(pat_let_tv11_), False, True, d_133_cf41_])))
            elif True:
                d_135___mcc_h29_ = source2_.cf31
                d_136_cf31_ = d_135___mcc_h29_
                return (pat_let_tv12_) | (pat_let_tv13_)

        rhs10_ = lambda3_(d_88_v56_)
        rhs11_ = d_85_v52_
        rhs12_ = lambda4_(d_92_v59_)
        lhs11_ = globalState
        lhs11_.f3 = rhs10_
        d_85_v52_ = rhs11_
        d_63_v46_ = rhs12_
        (d_85_v52_).f22 = p1
        d_137_v60_: _dafny.Seq
        d_137_v60_ = _dafny.SeqWithoutIsStrInference([d_85_v52_.f22, d_85_v52_.f22, d_85_v52_.f22, d_85_v52_.f22, True])
        r0 = (not(not((d_85_v52_.f22) not in (d_137_v60_)))) and (d_85_v52_.f22)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_138_v1_: _dafny.Seq
        d_138_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "truxk"))
        d_139_v2_: _dafny.Set
        d_139_v2_ = _dafny.Set({d_138_v1_, d_138_v1_})
        d_140_v3_: int
        d_140_v3_ = -610
        d_141_v4_: str
        d_141_v4_ = _dafny.CodePoint('g')
        d_142_v5_: _dafny.Array
        nw11_ = _dafny.Array(False, 29)
        d_142_v5_ = nw11_
        d_143_v6_: _dafny.Array
        nw12_ = _dafny.Array(int(0), 5)
        d_143_v6_ = nw12_
        d_144_v7_: _dafny.MultiSet
        d_144_v7_ = _dafny.MultiSet([d_140_v3_])
        d_145_globalState_: GlobalState
        nw13_ = GlobalState()
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: _dafny.Seq
            for compr_8_ in (d_139_v2_).Elements:
                d_146_v0_: _dafny.Seq = compr_8_
                if (d_146_v0_) in (d_139_v2_):
                    coll8_[d_146_v0_] = len(_dafny.SeqWithoutIsStrInference([281]))
            return _dafny.Map(coll8_)
        nw13_.ctor__(False, iife8_()
        , -658, 2, -712, True, ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))).set(default__.safeIndex((0) - (d_140_v3_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))), d_141_v4_)).set(default__.safeIndex(d_140_v3_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))).set(default__.safeIndex((0) - (d_140_v3_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))), d_141_v4_))), d_141_v4_), True, d_138_v1_, d_138_v1_, 896, True, _dafny.CodePoint('c'), -158, 316, -654, d_142_v5_, d_143_v6_, False, d_144_v7_, 308)
        d_145_globalState_ = nw13_
        d_147_v8_: bool
        d_147_v8_ = False
        d_148_i0_: int
        d_148_i0_ = 0
        with _dafny.label("0"):
            while not(not(d_147_v8_)):
                with _dafny.c_label("0"):
                    if (d_148_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_148_i0_ = (d_148_i0_) + (1)
                    d_149_v9_: _dafny.Array
                    nw14_ = _dafny.Array(_dafny.Array(None, 0), 27)
                    d_149_v9_ = nw14_
                    index10_ = default__.safeIndex(512, (d_149_v9_).length(0))
                    (d_149_v9_)[index10_] = d_142_v5_
                    index11_ = default__.safeIndex(512, (d_149_v9_).length(0))
                    rhs13_ = (d_140_v3_) - ((d_140_v3_ if d_147_v8_ else d_140_v3_))
                    rhs14_ = d_142_v5_
                    rhs15_ = d_140_v3_
                    lhs12_ = d_149_v9_
                    lhs13_ = default__.safeIndex(512, (d_149_v9_).length(0))
                    lhs14_ = d_145_globalState_
                    d_140_v3_ = rhs13_
                    lhs12_[lhs13_] = rhs14_
                    lhs14_.f20 = rhs15_
                    (d_145_globalState_).f20 = d_140_v3_
                    d_140_v3_ = d_140_v3_
                    d_150_v10_: bool
                    out0_: bool
                    out0_ = default__.m0(d_142_v5_, d_147_v8_, d_145_globalState_)
                    d_150_v10_ = out0_
                    pass
            pass
        if d_147_v8_:
            d_151_v11_: _dafny.Array
            def lambda5_(d_152_i1_):
                return _dafny.MultiSet([True])

            init1_ = lambda5_
            nw15_ = _dafny.Array(None, 23)
            for i0_1_ in range(nw15_.length(0)):
                nw15_[i0_1_] = init1_(i0_1_)
            d_151_v11_ = nw15_
            d_153_v12_: C1
            nw16_ = C1()
            nw16_.ctor__(d_151_v11_, not (False) or (d_147_v8_))
            d_153_v12_ = nw16_
            d_154_v13_: _dafny.MultiSet
            d_154_v13_ = _dafny.MultiSet([True])
            d_155_v14_: _dafny.MultiSet
            d_155_v14_ = _dafny.MultiSet([d_154_v13_])
            d_156_v15_: _dafny.Set
            d_156_v15_ = _dafny.Set({(d_153_v12_).f25, False})
            d_157_v16_: C2
            nw17_ = C2()
            nw17_.ctor__(default__.fm16(d_156_v15_, (d_153_v12_).f25, d_147_v8_, d_145_globalState_), d_138_v1_, d_147_v8_)
            d_157_v16_ = nw17_
            d_158_v17_: _dafny.Map
            d_158_v17_ = _dafny.Map({default__.fm12(d_140_v3_, d_154_v13_, (d_153_v12_).f25, (d_155_v14_).set(d_154_v13_, default__.abs(len(d_138_v1_))), d_145_globalState_): d_157_v16_})
            d_159_v18_: _dafny.Map
            d_159_v18_ = _dafny.Map({d_138_v1_: d_158_v17_})
            d_160_v19_: _dafny.Map
            d_160_v19_ = _dafny.Map({d_147_v8_: False})
            d_161_v20_: _dafny.Seq
            d_161_v20_ = _dafny.SeqWithoutIsStrInference([d_160_v19_, d_160_v19_, _dafny.Map({d_147_v8_: False})])
            rhs16_ = d_140_v3_
            rhs17_ = d_159_v18_
            rhs18_ = default__.safeDivisionInt(d_140_v3_, d_140_v3_)
            rhs19_ = (d_160_v19_) not in (d_161_v20_)
            lhs15_ = d_145_globalState_
            lhs16_ = d_145_globalState_
            lhs15_.f3 = rhs16_
            d_159_v18_ = rhs17_
            lhs16_.f3 = rhs18_
            d_147_v8_ = rhs19_
            d_162_v21_: bool
            out1_: bool
            out1_ = default__.m0(d_142_v5_, not((d_153_v12_).f25), d_145_globalState_)
            d_162_v21_ = out1_
            d_147_v8_ = ((d_153_v12_).f25) or ((d_153_v12_).f25)
            d_140_v3_ = (d_140_v3_ if (d_157_v16_).fm3(d_145_globalState_) else d_140_v3_)
        elif True:
            index12_ = default__.safeIndex(347, (d_143_v6_).length(0))
            (d_143_v6_)[index12_] = -385
            d_163_v22_: _dafny.Seq
            d_163_v22_ = _dafny.SeqWithoutIsStrInference([d_140_v3_])
            index13_ = default__.safeIndex(347, (d_143_v6_).length(0))
            (d_143_v6_)[index13_] = (0) - (((d_163_v22_)[default__.safeIndex(d_140_v3_, len(d_163_v22_))]) * (d_140_v3_))
            d_164_v23_: _dafny.MultiSet
            d_164_v23_ = _dafny.MultiSet([d_147_v8_, d_147_v8_])
            d_165_v24_: _dafny.Map
            d_165_v24_ = _dafny.Map({(d_143_v6_)[default__.safeIndex(347, (d_143_v6_).length(0))]: d_147_v8_})
            d_166_v25_: _dafny.MultiSet
            d_166_v25_ = _dafny.MultiSet([d_164_v23_])
            d_167_v26_: _dafny.Array
            nw18_ = _dafny.Array(None, 27)
            nw18_[int(0)] = d_164_v23_
            nw18_[int(1)] = _dafny.MultiSet([True])
            nw18_[int(2)] = d_164_v23_
            nw18_[int(3)] = d_164_v23_
            nw18_[int(4)] = d_164_v23_
            nw18_[int(5)] = d_164_v23_
            nw18_[int(6)] = d_164_v23_
            nw18_[int(7)] = d_164_v23_
            nw18_[int(8)] = d_164_v23_
            nw18_[int(9)] = d_164_v23_
            nw18_[int(10)] = d_164_v23_
            nw18_[int(11)] = d_164_v23_
            nw18_[int(12)] = d_164_v23_
            nw18_[int(13)] = _dafny.MultiSet([d_147_v8_, d_147_v8_, d_147_v8_])
            nw18_[int(14)] = d_164_v23_
            nw18_[int(15)] = d_164_v23_
            nw18_[int(16)] = _dafny.MultiSet([d_147_v8_])
            nw18_[int(17)] = _dafny.MultiSet([d_147_v8_, d_147_v8_])
            nw18_[int(18)] = _dafny.MultiSet([d_147_v8_, d_147_v8_])
            nw18_[int(19)] = _dafny.MultiSet([d_147_v8_, d_147_v8_, default__.fm12(d_140_v3_, _dafny.MultiSet([d_147_v8_, True, ((d_165_v24_)[d_140_v3_] if (d_140_v3_) in (d_165_v24_) else d_147_v8_)]), not(default__.fm12(d_140_v3_, _dafny.MultiSet([d_147_v8_, d_147_v8_]), d_147_v8_, d_166_v25_, d_145_globalState_)), d_166_v25_, d_145_globalState_)])
            nw18_[int(20)] = d_164_v23_
            nw18_[int(21)] = default__.fm17((d_143_v6_)[default__.safeIndex(347, (d_143_v6_).length(0))], d_147_v8_, d_145_globalState_)
            nw18_[int(22)] = _dafny.MultiSet([d_147_v8_, d_147_v8_, d_147_v8_])
            nw18_[int(23)] = d_164_v23_
            nw18_[int(24)] = d_164_v23_
            nw18_[int(25)] = _dafny.MultiSet([d_147_v8_])
            nw18_[int(26)] = d_164_v23_
            d_167_v26_ = nw18_
            d_168_v27_: _dafny.Map
            d_168_v27_ = _dafny.Map({default__.fm12(len(_dafny.Map({d_147_v8_: d_140_v3_})), d_164_v23_, d_147_v8_, d_166_v25_, d_145_globalState_): d_140_v3_})
            d_169_v28_: C1
            nw19_ = C1()
            nw19_.ctor__(d_167_v26_, ((d_143_v6_)[default__.safeIndex(347, (d_143_v6_).length(0))]) >= (len(d_168_v27_)))
            d_169_v28_ = nw19_
            d_170_v29_: _dafny.Seq
            d_170_v29_ = _dafny.SeqWithoutIsStrInference([True, d_147_v8_])
            d_171_v30_: D3
            d_171_v30_ = D3_DC7(_dafny.MultiSet(d_170_v29_))
            index14_ = default__.safeIndex(347, (d_143_v6_).length(0))
            rhs20_ = d_169_v28_
            rhs21_ = (len((d_139_v2_) | (d_139_v2_)) if (d_140_v3_) >= (d_140_v3_) else -794)
            rhs22_ = d_171_v30_
            rhs23_ = (default__.safeModuloInt(len(d_163_v22_), d_140_v3_)) >= (d_140_v3_)
            lhs17_ = d_143_v6_
            lhs18_ = default__.safeIndex(347, (d_143_v6_).length(0))
            d_169_v28_ = rhs20_
            lhs17_[lhs18_] = rhs21_
            d_171_v30_ = rhs22_
            d_147_v8_ = rhs23_
            d_172_v31_: _dafny.Array
            def lambda6_(d_173_i2_):
                return _dafny.Set({True})

            init2_ = lambda6_
            nw20_ = _dafny.Array(None, 24)
            for i0_2_ in range(nw20_.length(0)):
                nw20_[i0_2_] = init2_(i0_2_)
            d_172_v31_ = nw20_
            d_174_v32_: _dafny.Array
            nw21_ = _dafny.Array(None, 4)
            nw21_[int(0)] = d_172_v31_
            nw21_[int(1)] = d_172_v31_
            nw21_[int(2)] = d_172_v31_
            nw21_[int(3)] = d_172_v31_
            d_174_v32_ = nw21_
            index15_ = default__.safeIndex(134, (d_174_v32_).length(0))
            (d_174_v32_)[index15_] = d_172_v31_
            index16_ = default__.safeIndex(134, (d_174_v32_).length(0))
            (d_174_v32_)[index16_] = d_172_v31_
            d_175_v33_: _dafny.Seq
            d_175_v33_ = _dafny.SeqWithoutIsStrInference([d_163_v22_, default__.fm5(d_140_v3_, d_140_v3_, d_140_v3_, d_145_globalState_)])
            d_176_v34_: C2
            nw22_ = C2()
            nw22_.ctor__(d_175_v33_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), d_147_v8_)
            d_176_v34_ = nw22_
            d_177_v35_: _dafny.Set
            d_177_v35_ = _dafny.Set({d_176_v34_})
            d_177_v35_ = _dafny.Set({d_176_v34_, d_176_v34_, d_176_v34_, d_176_v34_, d_176_v34_})
            d_178_v36_: D0
            d_178_v36_ = D0_DC1((d_169_v28_).f25)
            d_179_v37_: C1
            nw23_ = C1()
            nw23_.ctor__(d_169_v28_.f24, (d_178_v36_).cf1)
            d_179_v37_ = nw23_
        d_147_v8_ = d_147_v8_
        d_180_v38_: _dafny.MultiSet
        d_180_v38_ = _dafny.MultiSet([False])
        d_181_v39_: _dafny.MultiSet
        d_181_v39_ = _dafny.MultiSet([d_180_v38_, d_180_v38_])
        d_182_v40_: D0
        d_182_v40_ = D0_DC0(d_142_v5_)
        d_183_v41_: D0
        d_183_v41_ = D0_DC2(d_182_v40_)
        source3_ = (d_183_v41_ if default__.fm12(d_140_v3_, d_180_v38_, d_147_v8_, d_181_v39_, d_145_globalState_) else d_183_v41_)
        if source3_.is_DC1:
            d_184___mcc_h0_ = source3_.cf1
            d_185_cf1_ = d_184___mcc_h0_
            index17_ = default__.safeIndex(881, (d_143_v6_).length(0))
            (d_143_v6_)[index17_] = 203
            d_186_v42_: _dafny.Seq
            d_186_v42_ = _dafny.SeqWithoutIsStrInference([d_140_v3_, 555, d_140_v3_, d_140_v3_])
            d_187_v43_: _dafny.Seq
            d_187_v43_ = _dafny.SeqWithoutIsStrInference([d_147_v8_, d_147_v8_])
            d_188_v44_: _dafny.Array
            nw24_ = _dafny.Array(None, 19)
            nw24_[int(0)] = d_180_v38_
            nw24_[int(1)] = d_180_v38_
            nw24_[int(2)] = d_180_v38_
            nw24_[int(3)] = d_180_v38_
            nw24_[int(4)] = d_180_v38_
            nw24_[int(5)] = d_180_v38_
            nw24_[int(6)] = d_180_v38_
            nw24_[int(7)] = _dafny.MultiSet([d_185_cf1_, d_185_cf1_, d_185_cf1_])
            nw24_[int(8)] = d_180_v38_
            nw24_[int(9)] = default__.fm17(len(d_186_v42_), d_147_v8_, d_145_globalState_)
            nw24_[int(10)] = d_180_v38_
            nw24_[int(11)] = d_180_v38_
            nw24_[int(12)] = d_180_v38_
            nw24_[int(13)] = d_180_v38_
            nw24_[int(14)] = _dafny.MultiSet(d_187_v43_)
            nw24_[int(15)] = d_180_v38_
            nw24_[int(16)] = d_180_v38_
            nw24_[int(17)] = d_180_v38_
            nw24_[int(18)] = default__.fm17(d_140_v3_, True, d_145_globalState_)
            d_188_v44_ = nw24_
            d_189_v45_: _dafny.Seq
            d_189_v45_ = _dafny.SeqWithoutIsStrInference([d_188_v44_])
            d_190_v46_: C1
            nw25_ = C1()
            nw25_.ctor__((d_189_v45_)[default__.safeIndex(616, len(d_189_v45_))], d_147_v8_)
            d_190_v46_ = nw25_
            d_191_v47_: _dafny.Seq
            d_191_v47_ = _dafny.SeqWithoutIsStrInference([d_190_v46_, d_190_v46_, d_190_v46_])
            index18_ = default__.safeIndex(881, (d_143_v6_).length(0))
            rhs24_ = (d_140_v3_) + (len(d_191_v47_))
            rhs25_ = (d_140_v3_) - (-786)
            lhs19_ = d_143_v6_
            lhs20_ = default__.safeIndex(881, (d_143_v6_).length(0))
            lhs21_ = d_145_globalState_
            lhs19_[lhs20_] = rhs24_
            lhs21_.f20 = rhs25_
            index19_ = default__.safeIndex(881, (d_143_v6_).length(0))
            (d_143_v6_)[index19_] = 723
            d_192_v48_: _dafny.Array
            nw26_ = _dafny.Array(None, 19)
            d_192_v48_ = nw26_
            d_193_v49_: T0
            nw27_ = C3()
            nw27_.ctor__((d_190_v46_).f25, d_147_v8_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))).set(default__.safeIndex((0) - (d_140_v3_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))), d_141_v4_), False)
            d_193_v49_ = nw27_
            index20_ = default__.safeIndex(947, (d_192_v48_).length(0))
            (d_192_v48_)[index20_] = d_193_v49_
            index21_ = default__.safeIndex(947, (d_192_v48_).length(0))
            (d_192_v48_)[index21_] = d_193_v49_
            d_194_v50_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_194_v50_ = nw28_
            d_194_v50_ = d_194_v50_
        elif source3_.is_DC0:
            d_195___mcc_h1_ = source3_.cf0
            d_196_cf0_ = d_195___mcc_h1_
            d_197_v51_: _dafny.Seq
            d_197_v51_ = _dafny.SeqWithoutIsStrInference([d_140_v3_])
            d_198_v52_: _dafny.Seq
            d_198_v52_ = _dafny.SeqWithoutIsStrInference([d_197_v51_])
            d_199_v53_: T2
            nw29_ = C2()
            nw29_.ctor__(d_198_v52_, d_138_v1_, (d_147_v8_ if d_147_v8_ else d_147_v8_))
            d_199_v53_ = nw29_
            d_199_v53_ = d_199_v53_
            d_200_v54_: _dafny.Map
            d_200_v54_ = _dafny.Map({d_138_v1_: d_140_v3_})
            d_201_v56_: C0
            nw30_ = C0()
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: _dafny.Seq
                for compr_9_ in (d_200_v54_).keys.Elements:
                    d_202_v55_: _dafny.Seq = compr_9_
                    if (d_202_v55_) in (d_200_v54_):
                        coll9_ = coll9_.union(_dafny.Set([d_202_v55_]))
                return _dafny.Set(coll9_)
            nw30_.ctor__(iife9_()
            )
            d_201_v56_ = nw30_
            (d_145_globalState_).f8 = _dafny.SeqWithoutIsStrInference([d_141_v4_ for d_203_i3_ in range(default__.abs(140))])
            d_204_v57_: _dafny.Seq
            d_204_v57_ = _dafny.SeqWithoutIsStrInference([not(d_199_v53_.f22), True, d_199_v53_.f22])
            if (d_204_v57_) < (d_204_v57_):
                index22_ = default__.safeIndex(657, (d_143_v6_).length(0))
                (d_143_v6_)[index22_] = 842
                index23_ = default__.safeIndex(483, (d_142_v5_).length(0))
                (d_142_v5_)[index23_] = not(d_147_v8_)
                d_205_v58_: _dafny.Seq
                d_205_v58_ = _dafny.SeqWithoutIsStrInference([d_199_v53_])
                d_206_v59_: _dafny.Set
                d_206_v59_ = _dafny.Set({_dafny.Map({True: _dafny.CodePoint('y')})})
                d_207_v60_: _dafny.Map
                d_207_v60_ = _dafny.Map({d_206_v59_: d_199_v53_})
                d_208_v61_: C3
                nw31_ = C3()
                nw31_.ctor__(d_199_v53_.f22, d_147_v8_, d_199_v53_.f21, False)
                d_208_v61_ = nw31_
                d_209_v62_: _dafny.MultiSet
                d_209_v62_ = _dafny.MultiSet([d_208_v61_])
                d_210_v63_: C3
                d_210_v63_ = d_208_v61_
                d_211_v64_: _dafny.Array
                nw32_ = _dafny.Array(None, 28)
                nw32_[int(0)] = d_199_v53_
                nw32_[int(1)] = d_199_v53_
                nw32_[int(2)] = d_199_v53_
                nw32_[int(3)] = d_199_v53_
                nw32_[int(4)] = d_199_v53_
                nw32_[int(5)] = (d_205_v58_)[default__.safeIndex((0) - (d_140_v3_), len(d_205_v58_))]
                nw32_[int(6)] = d_199_v53_
                nw32_[int(7)] = d_199_v53_
                nw32_[int(8)] = d_199_v53_
                nw32_[int(9)] = d_199_v53_
                nw32_[int(10)] = ((d_207_v60_)[d_206_v59_] if (d_206_v59_) in (d_207_v60_) else d_199_v53_)
                nw32_[int(11)] = d_199_v53_
                nw32_[int(12)] = d_199_v53_
                nw32_[int(13)] = (d_199_v53_ if d_199_v53_.f22 else d_199_v53_)
                nw32_[int(14)] = d_199_v53_
                nw32_[int(15)] = d_199_v53_
                nw32_[int(16)] = d_199_v53_
                nw32_[int(17)] = d_199_v53_
                nw32_[int(18)] = (d_205_v58_)[default__.safeIndex(((d_209_v62_)[(d_210_v63_)] if ((d_210_v63_)) in (d_209_v62_) else d_140_v3_), len(d_205_v58_))]
                nw32_[int(19)] = (d_205_v58_)[default__.safeIndex(len(d_199_v53_.f21), len(d_205_v58_))]
                nw32_[int(20)] = d_199_v53_
                nw32_[int(21)] = d_199_v53_
                nw32_[int(22)] = d_199_v53_
                nw32_[int(23)] = d_199_v53_
                nw32_[int(24)] = d_199_v53_
                nw32_[int(25)] = d_199_v53_
                nw32_[int(26)] = d_199_v53_
                nw32_[int(27)] = d_199_v53_
                d_211_v64_ = nw32_
                index24_ = default__.safeIndex(357, (d_211_v64_).length(0))
                (d_211_v64_)[index24_] = d_199_v53_
                d_212_v65_: _dafny.Seq
                d_212_v65_ = _dafny.SeqWithoutIsStrInference([d_199_v53_.f21, d_138_v1_])
                d_213_v66_: _dafny.Map
                d_213_v66_ = _dafny.Map({d_208_v61_.f28: d_140_v3_})
                d_214_v67_: _dafny.Map
                d_214_v67_ = _dafny.Map({d_213_v66_: d_140_v3_})
                d_215_v68_: _dafny.Array
                nw33_ = _dafny.Array(None, 14)
                nw33_[int(0)] = (d_212_v65_)[default__.safeIndex(len(d_214_v67_), len(d_212_v65_))]
                nw33_[int(1)] = d_138_v1_
                nw33_[int(2)] = d_199_v53_.f21
                nw33_[int(3)] = d_138_v1_
                nw33_[int(4)] = d_199_v53_.f21
                nw33_[int(5)] = default__.fm11((d_208_v61_).f27, d_145_globalState_)
                nw33_[int(6)] = d_199_v53_.f21
                nw33_[int(7)] = d_199_v53_.f21
                nw33_[int(8)] = d_199_v53_.f21
                nw33_[int(9)] = d_138_v1_
                nw33_[int(10)] = default__.fm11(d_208_v61_.f28, d_145_globalState_)
                nw33_[int(11)] = d_199_v53_.f21
                nw33_[int(12)] = _dafny.SeqWithoutIsStrInference([d_141_v4_ for d_216_i4_ in range(default__.abs(434))])
                nw33_[int(13)] = d_199_v53_.f21
                d_215_v68_ = nw33_
                d_217_v69_: _dafny.MultiSet
                d_217_v69_ = _dafny.MultiSet([d_215_v68_])
                d_218_v70_: _dafny.Set
                d_218_v70_ = _dafny.Set({d_143_v6_, d_143_v6_, d_143_v6_, d_143_v6_, d_143_v6_})
                index25_ = default__.safeIndex(657, (d_143_v6_).length(0))
                index26_ = default__.safeIndex(483, (d_142_v5_).length(0))
                index27_ = default__.safeIndex(357, (d_211_v64_).length(0))
                rhs26_ = ((d_217_v69_)[d_215_v68_] if (d_215_v68_) in (d_217_v69_) else d_140_v3_)
                rhs27_ = (d_218_v70_).isdisjoint((d_218_v70_).intersection(_dafny.Set({d_143_v6_})))
                rhs28_ = d_199_v53_
                lhs22_ = d_143_v6_
                lhs23_ = default__.safeIndex(657, (d_143_v6_).length(0))
                lhs24_ = d_142_v5_
                lhs25_ = default__.safeIndex(483, (d_142_v5_).length(0))
                lhs26_ = d_211_v64_
                lhs27_ = default__.safeIndex(357, (d_211_v64_).length(0))
                lhs22_[lhs23_] = rhs26_
                lhs24_[lhs25_] = rhs27_
                lhs26_[lhs27_] = rhs28_
                d_219_v71_: _dafny.Array
                nw34_ = _dafny.Array(int(0), 9)
                d_219_v71_ = nw34_
                d_219_v71_ = d_219_v71_
                d_220_v72_: D1
                d_220_v72_ = D1_DC5(d_147_v8_, d_213_v66_, d_140_v3_, d_144_v7_)
                d_221_v73_: T0
                nw35_ = C3()
                nw35_.ctor__((d_140_v3_) == (d_140_v3_), d_208_v61_.f28, d_199_v53_.f21, (d_140_v3_) <= ((d_220_v72_).cf9))
                d_221_v73_ = nw35_
                d_222_v74_: _dafny.Set
                d_222_v74_ = _dafny.Set({d_140_v3_})
                d_223_v75_: _dafny.Set
                d_223_v75_ = d_222_v74_
                d_224_v76_: T1
                nw36_ = C4()
                nw36_.ctor__(d_223_v75_, (d_208_v61_).f27, d_198_v52_, d_138_v1_, (d_140_v3_) < (len(d_204_v57_)))
                d_224_v76_ = nw36_
                rhs29_ = default__.fm11((d_221_v73_).fm1(d_138_v1_, (d_142_v5_)[default__.safeIndex(483, (d_142_v5_).length(0))], (d_208_v61_).f27, d_145_globalState_), d_145_globalState_)
                rhs30_ = d_221_v73_
                rhs31_ = d_224_v76_
                lhs28_ = d_199_v53_
                lhs28_.f21 = rhs29_
                d_221_v73_ = rhs30_
                d_224_v76_ = rhs31_
                d_225_v77_: _dafny.Map
                d_225_v77_ = _dafny.Map({(d_143_v6_)[default__.safeIndex(657, (d_143_v6_).length(0))]: d_221_v73_.f22})
                d_226_v78_: _dafny.Map
                d_226_v78_ = _dafny.Map({d_225_v77_: d_196_cf0_})
                d_227_v79_: int
                out2_: int
                out2_ = (d_199_v53_).m1(d_221_v73_.f22, d_226_v78_, d_145_globalState_)
                d_227_v79_ = out2_
                d_228_v80_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                d_228_v80_ = nw37_
                index28_ = default__.safeIndex(193, (d_228_v80_).length(0))
                (d_228_v80_)[index28_] = d_141_v4_
                d_229_v81_: _dafny.Set
                d_229_v81_ = _dafny.Set({not((d_208_v61_).f27), d_147_v8_})
                d_230_v82_: C4
                nw38_ = C4()
                nw38_.ctor__(d_223_v75_, not (True) or (d_147_v8_), default__.fm16(d_229_v81_, d_208_v61_.f28, d_147_v8_, d_145_globalState_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpt"))).set(default__.safeIndex(d_140_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpt")))), d_141_v4_), d_199_v53_.f22)
                d_230_v82_ = nw38_
                index29_ = default__.safeIndex(483, (d_142_v5_).length(0))
                index30_ = default__.safeIndex(193, (d_228_v80_).length(0))
                rhs32_ = (d_230_v82_).fm1(d_138_v1_, not((d_230_v82_).f30), d_208_v61_.f28, d_145_globalState_)
                rhs33_ = d_141_v4_
                rhs34_ = d_230_v82_
                lhs29_ = d_142_v5_
                lhs30_ = default__.safeIndex(483, (d_142_v5_).length(0))
                lhs31_ = d_228_v80_
                lhs32_ = default__.safeIndex(193, (d_228_v80_).length(0))
                lhs29_[lhs30_] = rhs32_
                lhs31_[lhs32_] = rhs33_
                d_230_v82_ = rhs34_
            elif True:
                (d_199_v53_).f22 = (d_204_v57_)[default__.safeIndex(len((d_197_v51_ if d_199_v53_.f22 else default__.fm5((0) - (default__.fm6(d_145_globalState_)), d_140_v3_, 109, d_145_globalState_))), len(d_204_v57_))]
                (d_145_globalState_).f20 = d_140_v3_
                d_147_v8_ = d_147_v8_
                d_231_v83_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_231_v83_ = nw39_
                d_232_v84_: C1
                nw40_ = C1()
                nw40_.ctor__(d_231_v83_, False)
                d_232_v84_ = nw40_
                d_233_v85_: _dafny.Map
                d_233_v85_ = _dafny.Map({d_199_v53_.f22: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))})
                d_234_v86_: _dafny.Map
                d_234_v86_ = _dafny.Map({d_232_v84_: default__.fm12(d_140_v3_, ((d_233_v85_)[False] if (False) in (d_233_v85_) else _dafny.MultiSet([d_199_v53_.f22])), d_199_v53_.f22, d_181_v39_, d_145_globalState_)})
                d_234_v86_ = (d_234_v86_).set(d_232_v84_, (d_204_v57_)[default__.safeIndex(default__.fm6(d_145_globalState_), len(d_204_v57_))])
                d_180_v38_ = _dafny.MultiSet([not(d_199_v53_.f22)])
        elif True:
            d_235___mcc_h2_ = source3_.cf2
            d_236_cf2_ = d_235___mcc_h2_
            index31_ = default__.safeIndex(775, (d_143_v6_).length(0))
            (d_143_v6_)[index31_] = d_140_v3_
            index32_ = default__.safeIndex(775, (d_143_v6_).length(0))
            (d_143_v6_)[index32_] = d_140_v3_
            (d_145_globalState_).f20 = default__.fm6(d_145_globalState_)
            d_237_v87_: _dafny.Seq
            d_237_v87_ = _dafny.SeqWithoutIsStrInference([(d_143_v6_)[default__.safeIndex(775, (d_143_v6_).length(0))], d_140_v3_])
            rhs35_ = (d_140_v3_) + (d_140_v3_)
            rhs36_ = (d_237_v87_)[default__.safeIndex(d_140_v3_, len(d_237_v87_))]
            lhs33_ = d_145_globalState_
            lhs34_ = d_145_globalState_
            lhs33_.f20 = rhs35_
            lhs34_.f3 = rhs36_
            d_238_v88_: _dafny.Set
            d_238_v88_ = _dafny.Set({(d_143_v6_)[default__.safeIndex(775, (d_143_v6_).length(0))]})
            d_239_v89_: _dafny.Seq
            d_239_v89_ = _dafny.SeqWithoutIsStrInference([d_238_v88_])
            d_240_v90_: _dafny.Set
            d_240_v90_ = (d_239_v89_)[default__.safeIndex(d_140_v3_, len(d_239_v89_))]
            d_241_v91_: _dafny.Seq
            d_241_v91_ = _dafny.SeqWithoutIsStrInference([d_237_v87_, d_237_v87_])
            d_242_v92_: T1
            nw41_ = C4()
            nw41_.ctor__(_dafny.Set({(d_143_v6_)[default__.safeIndex(775, (d_143_v6_).length(0))]}), d_147_v8_, d_241_v91_, _dafny.SeqWithoutIsStrInference([d_141_v4_ for d_243_i5_ in range(default__.abs(-37))]), (default__.fm6(d_145_globalState_)) in (d_237_v87_))
            d_242_v92_ = nw41_
            rhs37_ = d_242_v92_.f22
            rhs38_ = (d_242_v92_ if not (d_147_v8_) or (d_147_v8_) else d_242_v92_)
            rhs39_ = d_140_v3_
            lhs35_ = d_145_globalState_
            d_147_v8_ = rhs37_
            d_242_v92_ = rhs38_
            lhs35_.f3 = rhs39_
        d_244_v93_: _dafny.MultiSet
        d_244_v93_ = _dafny.MultiSet([d_142_v5_, d_142_v5_])
        d_245_v94_: _dafny.Map
        d_245_v94_ = _dafny.Map({d_140_v3_: d_147_v8_})
        d_246_v95_: _dafny.Seq
        d_246_v95_ = _dafny.SeqWithoutIsStrInference([d_147_v8_, d_147_v8_])
        d_247_v96_: _dafny.Array
        nw42_ = _dafny.Array(None, 25)
        nw42_[int(0)] = d_244_v93_
        nw42_[int(1)] = d_244_v93_
        nw42_[int(2)] = d_244_v93_
        nw42_[int(3)] = (d_244_v93_).intersection(((D7_DC15(d_244_v93_)).cf25).set(d_142_v5_, default__.abs(d_140_v3_)))
        nw42_[int(4)] = d_244_v93_
        nw42_[int(5)] = _dafny.MultiSet([d_142_v5_, d_142_v5_])
        nw42_[int(6)] = (d_244_v93_) - (d_244_v93_)
        nw42_[int(7)] = d_244_v93_
        nw42_[int(8)] = d_244_v93_
        nw42_[int(9)] = d_244_v93_
        nw42_[int(10)] = d_244_v93_
        nw42_[int(11)] = (d_244_v93_) | (d_244_v93_)
        nw42_[int(12)] = d_244_v93_
        nw42_[int(13)] = (d_244_v93_ if ((d_245_v94_)[d_140_v3_] if (d_140_v3_) in (d_245_v94_) else not(d_147_v8_)) else d_244_v93_)
        nw42_[int(14)] = d_244_v93_
        nw42_[int(15)] = _dafny.MultiSet([d_142_v5_])
        nw42_[int(16)] = d_244_v93_
        nw42_[int(17)] = (d_244_v93_ if default__.fm12(572, _dafny.MultiSet([d_147_v8_, d_147_v8_, d_147_v8_, d_147_v8_, d_147_v8_]), (d_246_v95_)[default__.safeIndex(len(d_138_v1_), len(d_246_v95_))], _dafny.MultiSet([d_180_v38_]), d_145_globalState_) else d_244_v93_)
        nw42_[int(18)] = (D7_DC15(d_244_v93_)).cf25
        nw42_[int(19)] = d_244_v93_
        nw42_[int(20)] = d_244_v93_
        nw42_[int(21)] = (d_244_v93_).set(d_142_v5_, default__.abs(d_140_v3_))
        nw42_[int(22)] = d_244_v93_
        nw42_[int(23)] = (d_244_v93_) | (d_244_v93_)
        nw42_[int(24)] = (_dafny.MultiSet([d_142_v5_])) | (d_244_v93_)
        d_247_v96_ = nw42_
        index33_ = default__.safeIndex(543, (d_247_v96_).length(0))
        (d_247_v96_)[index33_] = (_dafny.MultiSet([d_142_v5_, d_142_v5_])) | (d_244_v93_)
        index34_ = default__.safeIndex(543, (d_247_v96_).length(0))
        (d_247_v96_)[index34_] = (((d_244_v93_).set(d_142_v5_, default__.abs(len(d_245_v94_)))) | (d_244_v93_)) - (d_244_v93_)
        d_248_v97_: _dafny.Set
        d_248_v97_ = _dafny.Set({d_140_v3_, d_140_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ly")))})
        d_249_v98_: _dafny.Set
        d_249_v98_ = d_248_v97_
        d_250_v99_: _dafny.Seq
        d_250_v99_ = _dafny.SeqWithoutIsStrInference([d_140_v3_])
        d_251_v100_: C4
        nw43_ = C4()
        nw43_.ctor__(d_249_v98_, d_147_v8_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_140_v3_]), d_250_v99_, (_dafny.SeqWithoutIsStrInference([d_140_v3_])).set(default__.safeIndex(d_140_v3_, len(_dafny.SeqWithoutIsStrInference([d_140_v3_]))), d_140_v3_)]), d_138_v1_, True)
        d_251_v100_ = nw43_
        d_252_v101_: _dafny.Set
        d_252_v101_ = _dafny.Set({d_251_v100_, d_251_v100_, d_251_v100_})
        d_253_v102_: _dafny.Map
        d_253_v102_ = _dafny.Map({default__.fm12(len(d_138_v1_), _dafny.MultiSet([d_147_v8_]), d_147_v8_, d_181_v39_, d_145_globalState_): d_252_v101_})
        hi1_ = (0) - (len(d_253_v102_))
        for d_254_i6_ in range(d_140_v3_, hi1_):
            d_255_v103_: _dafny.Seq
            d_255_v103_ = _dafny.SeqWithoutIsStrInference([d_250_v99_])
            d_256_v104_: _dafny.Seq
            d_256_v104_ = _dafny.SeqWithoutIsStrInference([(d_255_v103_)[default__.safeIndex(d_140_v3_, len(d_255_v103_))]])
            d_257_v105_: C2
            nw44_ = C2()
            nw44_.ctor__(d_255_v103_, (d_138_v1_) + ((d_138_v1_).set(default__.safeIndex(d_254_i6_, len(d_138_v1_)), _dafny.CodePoint('u'))), (d_140_v3_) >= (d_140_v3_))
            d_257_v105_ = nw44_
            d_258_v106_: _dafny.Seq
            d_258_v106_ = _dafny.SeqWithoutIsStrInference([d_245_v94_])
            d_259_v107_: int
            out3_: int
            out3_ = (d_257_v105_).m1((d_251_v100_).f30, _dafny.Map({(d_258_v106_)[default__.safeIndex(840, len(d_258_v106_))]: d_142_v5_}), d_145_globalState_)
            d_259_v107_ = out3_
            d_147_v8_ = (d_251_v100_).f30
            d_250_v99_ = d_250_v99_
        d_147_v8_ = (d_251_v100_).f30
        d_260_v108_: _dafny.Array
        nw45_ = _dafny.Array(None, 8)
        nw45_[int(0)] = d_141_v4_
        nw45_[int(1)] = (d_141_v4_ if d_147_v8_ else d_141_v4_)
        nw45_[int(2)] = d_141_v4_
        nw45_[int(3)] = d_141_v4_
        nw45_[int(4)] = d_141_v4_
        nw45_[int(5)] = d_141_v4_
        nw45_[int(6)] = d_141_v4_
        nw45_[int(7)] = d_141_v4_
        d_260_v108_ = nw45_
        d_260_v108_ = d_260_v108_
        d_250_v99_ = d_250_v99_
        d_261_v109_: D3
        d_261_v109_ = D3_DC7(d_180_v38_)
        pat_let_tv14_ = d_180_v38_
        d_262_v110_: _dafny.Array
        nw46_ = _dafny.Array(None, 25)
        nw46_[int(0)] = _dafny.MultiSet([False])
        nw46_[int(1)] = default__.fm17(d_140_v3_, (d_251_v100_).f30, d_145_globalState_)
        nw46_[int(2)] = d_180_v38_
        nw46_[int(3)] = d_180_v38_
        nw46_[int(4)] = d_180_v38_
        def iife10_(_pat_let0_0):
            def iife11_(d_263_dt__update__tmp_h1_):
                def iife12_(_pat_let1_0):
                    def iife13_(d_264_dt__update_hcf12_h0_):
                        return D3_DC7(d_264_dt__update_hcf12_h0_)
                    return iife13_(_pat_let1_0)
                return iife12_(pat_let_tv14_)
            return iife11_(_pat_let0_0)
        nw46_[int(5)] = (iife10_(d_261_v109_)).cf12
        nw46_[int(6)] = d_180_v38_
        nw46_[int(7)] = d_180_v38_
        nw46_[int(8)] = d_180_v38_
        nw46_[int(9)] = _dafny.MultiSet([d_147_v8_, (d_251_v100_).f30])
        nw46_[int(10)] = d_180_v38_
        nw46_[int(11)] = d_180_v38_
        nw46_[int(12)] = d_180_v38_
        nw46_[int(13)] = _dafny.MultiSet((d_246_v95_).set(default__.safeIndex(d_140_v3_, len(d_246_v95_)), (d_251_v100_).fm1(d_138_v1_, d_147_v8_, (d_251_v100_).f30, d_145_globalState_)))
        nw46_[int(14)] = _dafny.MultiSet([not((d_251_v100_).f30)])
        nw46_[int(15)] = d_180_v38_
        nw46_[int(16)] = d_180_v38_
        nw46_[int(17)] = d_180_v38_
        nw46_[int(18)] = _dafny.MultiSet([(d_251_v100_).f30])
        nw46_[int(19)] = d_180_v38_
        nw46_[int(20)] = (d_180_v38_).set(d_147_v8_, default__.abs((0) - (d_140_v3_)))
        nw46_[int(21)] = d_180_v38_
        nw46_[int(22)] = d_180_v38_
        nw46_[int(23)] = default__.fm17(889, d_147_v8_, d_145_globalState_)
        nw46_[int(24)] = _dafny.MultiSet(d_246_v95_)
        d_262_v110_ = nw46_
        d_265_v111_: C1
        nw47_ = C1()
        nw47_.ctor__(d_262_v110_, (d_251_v100_).f30)
        d_265_v111_ = nw47_
        d_266_v112_: _dafny.Map
        d_266_v112_ = _dafny.Map({d_140_v3_: d_140_v3_})
        d_267_v113_: _dafny.Map
        d_267_v113_ = _dafny.Map({d_180_v38_: d_140_v3_})
        d_268_v114_: D1
        d_268_v114_ = D1_DC5(False, _dafny.Map({(d_265_v111_).f25: d_140_v3_}), d_140_v3_, d_144_v7_)
        d_140_v3_ = default__.safeModuloInt(len(d_266_v112_), (D3_DC8(d_140_v3_, d_267_v113_, d_268_v114_, d_147_v8_)).cf13)
        hi2_ = d_140_v3_
        for d_269_i7_ in range(default__.safeDivisionInt(86, d_140_v3_), hi2_):
            d_270_v115_: _dafny.Seq
            d_270_v115_ = _dafny.SeqWithoutIsStrInference([default__.fm11((d_265_v111_).f25, d_145_globalState_), d_138_v1_, default__.fm11(False, d_145_globalState_)])
            d_138_v1_ = ((d_270_v115_)[default__.safeIndex(d_140_v3_, len(d_270_v115_))]) + (d_138_v1_)
            (d_145_globalState_).f20 = 761
            index35_ = default__.safeIndex(578, (d_260_v108_).length(0))
            (d_260_v108_)[index35_] = d_141_v4_
            d_271_v116_: _dafny.MultiSet
            d_271_v116_ = _dafny.MultiSet([d_138_v1_])
            index36_ = default__.safeIndex(578, (d_260_v108_).length(0))
            (d_260_v108_)[index36_] = (d_138_v1_)[default__.safeIndex(((_dafny.MultiSet([d_138_v1_])).intersection(d_271_v116_)).cardinality, len(d_138_v1_))]
            d_140_v3_ = ((d_180_v38_).set((d_265_v111_).f25, default__.abs((0) - ((d_140_v3_) + (984))))).cardinality
        d_272_v117_: C3
        nw48_ = C3()
        nw48_.ctor__((d_265_v111_).f25, d_147_v8_, d_138_v1_, (d_265_v111_).f25)
        d_272_v117_ = nw48_
        d_140_v3_ = d_140_v3_
        (d_145_globalState_).f12 = (default__.fm4(d_272_v117_.f28, d_180_v38_, d_145_globalState_) if (d_272_v117_).f27 else _dafny.CodePoint('q'))
        (d_272_v117_).f28 = (d_265_v111_).f25
        _dafny.print((d_138_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v2_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "truxk"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_140_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_141_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_144_v7_) == (_dafny.MultiSet([-610]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f1) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "truxk")): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_145_globalState_.f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_145_globalState_.f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_145_globalState_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f17)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_145_globalState_).f17)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_globalState_.f19) == (_dafny.MultiSet([-610]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_145_globalState_.f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_147_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_148_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v38_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_v39_) == (_dafny.MultiSet([_dafny.MultiSet([False]), _dafny.MultiSet([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v93_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v94_) == (_dafny.Map({-610: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v95_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[0]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[1]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[2]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[3]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[4]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[5]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[6]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[7]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[8]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[9]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[10]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[11]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[12]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[13]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[14]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[15]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[16]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[17]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[18]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[19]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[20]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[21]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[22]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[23]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v96_)[24]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v97_) == (_dafny.Set({-610, 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_249_v98_)) == (_dafny.Set({-610, 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v99_) == (_dafny.SeqWithoutIsStrInference([-610]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_251_v100_.f29)) == (_dafny.Set({-610, 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v100_).f30))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_252_v101_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_253_v102_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v108_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v109_).cf12) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[0]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[1]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[2]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[3]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[4]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[5]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[6]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[7]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[8]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[9]) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[10]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[11]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[12]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[13]) == (_dafny.MultiSet([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[14]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[15]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[16]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[17]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[18]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[19]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[20]) == (_dafny.MultiSet([False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[21]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[22]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[23]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v110_)[24]) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[0]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[1]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[2]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[3]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[4]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[5]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[6]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[7]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[8]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[9]) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[10]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[11]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[12]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[13]) == (_dafny.MultiSet([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[14]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[15]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[16]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[17]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[18]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[19]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[20]) == (_dafny.MultiSet([False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[21]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[22]) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[23]) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_265_v111_.f24)[24]) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v111_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v112_) == (_dafny.Map({-610: -610}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_267_v113_) == (_dafny.Map({_dafny.MultiSet([False]): -610}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v114_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_268_v114_).cf8) == (_dafny.Map({False: -610}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v114_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_268_v114_).cf10) == (_dafny.MultiSet([-610]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v117_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_272_v117_.f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.CodePoint('D'), _dafny.Map({}), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(int(0), _dafny.Map({}), D1.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC11(D4, NamedTuple('DC11', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC13(D5, NamedTuple('DC13', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC14(D6, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(False, False, None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC16(D7, NamedTuple('DC16', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)

class D8_DC17(D8, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC19(int(0), int(0), False, int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)

class D9_DC19(D9, NamedTuple('DC19', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({self.cf37.VerbatimString(True)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC18(D9, NamedTuple('DC18', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f22(self):
        return self._f22
    @f22.setter
    def f22(self, value):
        self._f22 = value

class T1:
    pass
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value

class T2:
    pass
    def m1(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f3: int = int(0)
        self.f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f12: str = _dafny.CodePoint('D')
        self.f19: _dafny.MultiSet = _dafny.MultiSet({})
        self.f20: int = int(0)
        self._f0: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: int = int(0)
        self._f4: int = int(0)
        self._f5: bool = False
        self._f7: bool = False
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: int = int(0)
        self._f11: bool = False
        self._f13: int = int(0)
        self._f14: int = int(0)
        self._f15: int = int(0)
        self._f16: _dafny.Array = _dafny.Array(None, 0)
        self._f17: _dafny.Array = _dafny.Array(None, 0)
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self).f20 = f20

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
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f7(self):
        return self._f7
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
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
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C0:
    def  __init__(self):
        self.f26: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f26):
        (self).f26 = f26

    def fm7(self, p0, globalState):
        return not (False) or ((False) and (True))

    def fm8(self, globalState):
        return (D1_DC3(_dafny.Set({_dafny.Set({_dafny.CodePoint('r')})}))).cf3


class C1:
    def  __init__(self):
        self.f24: _dafny.Array = _dafny.Array(None, 0)
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f24, f25):
        (self).f24 = f24
        (self)._f25 = f25

    def m3(self, p0, p1, p2, globalState):
        d_273_v0_: bool
        d_273_v0_ = True
        d_274_v1_: int
        d_274_v1_ = 976
        d_275_v2_: _dafny.Set
        d_275_v2_ = _dafny.Set({d_274_v1_, d_274_v1_, d_274_v1_})
        d_273_v0_ = (p1) or ((d_275_v2_) != (d_275_v2_))
        d_276_v3_: _dafny.Array
        nw49_ = _dafny.Array(False, 21)
        d_276_v3_ = nw49_
        d_277_v4_: D0
        d_277_v4_ = D0_DC0(d_276_v3_)
        pat_let_tv15_ = d_276_v3_
        def iife14_(_pat_let2_0):
            def iife15_(d_278_dt__update__tmp_h0_):
                def iife16_(_pat_let3_0):
                    def iife17_(d_279_dt__update_hcf0_h0_):
                        return D0_DC0(d_279_dt__update_hcf0_h0_)
                    return iife17_(_pat_let3_0)
                return iife16_(pat_let_tv15_)
            return iife15_(_pat_let2_0)
        d_277_v4_ = iife14_(d_277_v4_)
        d_280_v5_: _dafny.Seq
        d_280_v5_ = _dafny.SeqWithoutIsStrInference([p0, (self).f25])
        d_281_v6_: _dafny.MultiSet
        d_281_v6_ = _dafny.MultiSet([not(not(d_273_v0_)), (self).f25, p0, (d_280_v5_)[default__.safeIndex(default__.fm6(globalState), len(d_280_v5_))], (self).f25])
        d_282_v7_: _dafny.Seq
        d_282_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekkjfqs"))
        index37_ = default__.safeIndex(442, (d_276_v3_).length(0))
        (d_276_v3_)[index37_] = (d_282_v7_) < (d_282_v7_)
        d_283_v8_: _dafny.Seq
        d_283_v8_ = _dafny.SeqWithoutIsStrInference([d_280_v5_, d_280_v5_, d_280_v5_])
        index38_ = default__.safeIndex(442, (d_276_v3_).length(0))
        rhs40_ = d_281_v6_
        rhs41_ = (d_280_v5_) == ((d_283_v8_)[default__.safeIndex(len(d_282_v7_), len(d_283_v8_))])
        rhs42_ = p0
        lhs36_ = d_276_v3_
        lhs37_ = default__.safeIndex(442, (d_276_v3_).length(0))
        d_281_v6_ = rhs40_
        d_273_v0_ = rhs41_
        lhs36_[lhs37_] = rhs42_
        d_284_i0_: int
        d_284_i0_ = 0
        with _dafny.label("1"):
            while (d_276_v3_)[default__.safeIndex(442, (d_276_v3_).length(0))]:
                with _dafny.c_label("1"):
                    if (d_284_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_284_i0_ = (d_284_i0_) + (1)
                    d_285_v9_: _dafny.Array
                    nw50_ = _dafny.Array(_dafny.Map({}), 19)
                    d_285_v9_ = nw50_
                    d_285_v9_ = d_285_v9_
                    d_276_v3_ = d_276_v3_
                    d_286_v10_: D0
                    d_286_v10_ = D0_DC1((self).f25)
                    if (d_286_v10_).cf1:
                        (globalState).f3 = (len(_dafny.Set({(0) - (d_274_v1_)}))) + ((_dafny.MultiSet([d_274_v1_])).cardinality)
                        d_273_v0_ = p0
                        d_287_v11_: _dafny.Set
                        d_287_v11_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "px"))})
                        d_288_v12_: C0
                        nw51_ = C0()
                        nw51_.ctor__(d_287_v11_)
                        d_288_v12_ = nw51_
                        (globalState).f3 = len((d_282_v7_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_289_i1_ in range(default__.abs(687))])))
                        d_290_v13_: _dafny.Map
                        d_290_v13_ = _dafny.Map({p1: d_274_v1_})
                        d_291_v14_: _dafny.Seq
                        d_291_v14_ = _dafny.SeqWithoutIsStrInference([d_274_v1_, len(default__.fm9(d_273_v0_, p2, globalState))])
                        d_292_v15_: _dafny.Map
                        d_292_v15_ = _dafny.Map({d_274_v1_: (_dafny.MultiSet(d_291_v14_)).set(d_274_v1_, default__.abs(d_274_v1_))})
                        d_293_v16_: _dafny.MultiSet
                        d_293_v16_ = _dafny.MultiSet([len(d_280_v5_)])
                        d_294_v17_: D1
                        d_294_v17_ = D1_DC5((d_286_v10_).cf1, d_290_v13_, (0) - (len(d_282_v7_)), ((d_292_v15_)[d_274_v1_] if (d_274_v1_) in (d_292_v15_) else d_293_v16_))
                        d_294_v17_ = d_294_v17_
                    elif True:
                        index39_ = default__.safeIndex(442, (d_276_v3_).length(0))
                        (d_276_v3_)[index39_] = not (p0) or (p0)
                        d_274_v1_ = (515 if p1 else d_274_v1_)
                        d_295_v18_: _dafny.Set
                        d_295_v18_ = _dafny.Set({d_282_v7_})
                        d_296_v19_: C0
                        nw52_ = C0()
                        nw52_.ctor__(d_295_v18_)
                        d_296_v19_ = nw52_
                        (globalState).f20 = default__.safeModuloInt(219, d_274_v1_)
                        d_273_v0_ = d_273_v0_
                    source4_ = default__.fm10(globalState)
                    if source4_.is_DC4:
                        d_297___mcc_h0_ = source4_.cf4
                        d_298___mcc_h1_ = source4_.cf5
                        d_299___mcc_h2_ = source4_.cf6
                        d_300_cf6_ = d_299___mcc_h2_
                        d_301_cf5_ = d_298___mcc_h1_
                        d_302_cf4_ = d_297___mcc_h0_
                        d_303_v20_: _dafny.Array
                        def lambda7_(d_304_v10_, d_305_p1_):
                            def lambda8_(d_306_i2_):
                                def iife18_(_pat_let4_0):
                                    def iife19_(d_307_dt__update__tmp_h1_):
                                        def iife20_(_pat_let5_0):
                                            def iife21_(d_308_dt__update_hcf1_h0_):
                                                return D0_DC1(d_308_dt__update_hcf1_h0_)
                                            return iife21_(_pat_let5_0)
                                        return iife20_(d_305_p1_)
                                    return iife19_(_pat_let4_0)
                                return iife18_(d_304_v10_)

                            return lambda8_

                        init3_ = lambda7_(d_286_v10_, p1)
                        nw53_ = _dafny.Array(None, 20)
                        for i0_3_ in range(nw53_.length(0)):
                            nw53_[i0_3_] = init3_(i0_3_)
                        d_303_v20_ = nw53_
                        d_309_v21_: _dafny.Map
                        d_309_v21_ = _dafny.Map({d_303_v20_: d_274_v1_})
                        d_310_v22_: _dafny.Map
                        d_310_v22_ = _dafny.Map({(d_276_v3_)[default__.safeIndex(442, (d_276_v3_).length(0))]: d_309_v21_})
                        d_310_v22_ = (d_310_v22_).set((self).f25, d_309_v21_)
                        d_311_v23_: _dafny.Array
                        nw54_ = _dafny.Array(int(0), 5)
                        d_311_v23_ = nw54_
                        d_312_v24_: _dafny.Map
                        d_312_v24_ = _dafny.Map({False: d_311_v23_})
                        d_312_v24_ = (d_312_v24_).set(p0, d_311_v23_)
                        d_313_v25_: _dafny.Array
                        def lambda9_(d_314_cf4_):
                            def lambda10_(d_315_i3_):
                                return _dafny.SeqWithoutIsStrInference([d_314_cf4_ for d_316_i4_ in range(default__.abs(-865))])

                            return lambda10_

                        init4_ = lambda9_(d_302_cf4_)
                        nw55_ = _dafny.Array(None, 8)
                        for i0_4_ in range(nw55_.length(0)):
                            nw55_[i0_4_] = init4_(i0_4_)
                        d_313_v25_ = nw55_
                        index40_ = default__.safeIndex(96, (d_313_v25_).length(0))
                        (d_313_v25_)[index40_] = default__.fm11(False, globalState)
                        index41_ = default__.safeIndex(96, (d_313_v25_).length(0))
                        (d_313_v25_)[index41_] = d_282_v7_
                        d_274_v1_ = (0) - (((d_274_v1_) - (d_274_v1_)) - ((d_300_cf6_)[default__.safeIndex(d_274_v1_, len(d_300_cf6_))]))
                    elif source4_.is_DC5:
                        d_317___mcc_h3_ = source4_.cf7
                        d_318___mcc_h4_ = source4_.cf8
                        d_319___mcc_h5_ = source4_.cf9
                        d_320___mcc_h6_ = source4_.cf10
                        d_321_cf10_ = d_320___mcc_h6_
                        d_322_cf9_ = d_319___mcc_h5_
                        d_323_cf8_ = d_318___mcc_h4_
                        d_324_cf7_ = d_317___mcc_h3_
                        d_325_v26_: str
                        d_325_v26_ = _dafny.CodePoint('t')
                        d_273_v0_ = ((d_325_v26_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdncgjj"))) if p1 else (d_276_v3_)[default__.safeIndex(442, (d_276_v3_).length(0))])
                        index42_ = default__.safeIndex(442, (d_276_v3_).length(0))
                        (d_276_v3_)[index42_] = ((d_275_v2_) - (d_275_v2_)).ispropersubset(d_275_v2_)
                        d_326_v27_: _dafny.Seq
                        d_326_v27_ = _dafny.SeqWithoutIsStrInference([d_322_cf9_, d_322_cf9_])
                        d_327_v28_: _dafny.MultiSet
                        d_327_v28_ = _dafny.MultiSet([d_281_v6_])
                        d_328_v29_: _dafny.Map
                        d_328_v29_ = _dafny.Map({d_274_v1_: d_322_cf9_})
                        index43_ = default__.safeIndex(442, (d_276_v3_).length(0))
                        (d_276_v3_)[index43_] = default__.fm12(len((d_326_v27_) + (d_326_v27_)), ((d_281_v6_).set(d_273_v0_, default__.abs(d_322_cf9_))).intersection(d_281_v6_), (d_276_v3_)[default__.safeIndex(442, (d_276_v3_).length(0))], _dafny.MultiSet([_dafny.MultiSet([default__.fm12(d_322_cf9_, d_281_v6_, p2, d_327_v28_, globalState), p2]), _dafny.MultiSet(d_280_v5_), _dafny.MultiSet([(d_280_v5_)[default__.safeIndex((d_326_v27_)[default__.safeIndex(len(d_328_v29_), len(d_326_v27_))], len(d_280_v5_))], d_273_v0_]), _dafny.MultiSet([p0, (self).f25])]), globalState)
                        d_329_v30_: _dafny.Map
                        d_329_v30_ = _dafny.Map({d_322_cf9_: p2})
                        d_330_v31_: _dafny.Map
                        d_330_v31_ = _dafny.Map({d_321_cf10_: default__.fm11(((d_329_v30_)[d_274_v1_] if (d_274_v1_) in (d_329_v30_) else d_273_v0_), globalState)})
                        d_330_v31_ = (_dafny.Map({d_321_cf10_: _dafny.SeqWithoutIsStrInference([d_325_v26_ for d_331_i5_ in range(default__.abs(786))])})) | (d_330_v31_)
                    elif True:
                        d_332___mcc_h7_ = source4_.cf3
                        d_333_cf3_ = d_332___mcc_h7_
                        d_273_v0_ = ((d_274_v1_) * (d_274_v1_)) <= (d_274_v1_)
                        (globalState).f20 = default__.fm6(globalState)
                        (globalState).f3 = d_274_v1_
                        d_334_v32_: _dafny.Set
                        d_334_v32_ = _dafny.Set({d_282_v7_})
                        d_335_v33_: C0
                        nw56_ = C0()
                        nw56_.ctor__((d_334_v32_).intersection(d_334_v32_))
                        d_335_v33_ = nw56_
                    pass
            pass
        d_336_v34_: _dafny.Map
        d_336_v34_ = _dafny.Map({len(d_280_v5_): len(d_282_v7_)})
        d_337_v35_: _dafny.Seq
        d_337_v35_ = _dafny.SeqWithoutIsStrInference([d_336_v34_])
        index44_ = default__.safeIndex(442, (d_276_v3_).length(0))
        def iife22_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in ((d_337_v35_)[default__.safeIndex(d_274_v1_, len(d_337_v35_))]).keys.Elements:
                d_338_v36_: int = compr_10_
                if (d_338_v36_) in ((d_337_v35_)[default__.safeIndex(d_274_v1_, len(d_337_v35_))]):
                    coll10_ = coll10_.union(_dafny.Set([(d_338_v36_) - (len(_dafny.Set({(D0_DC1(True)).cf1})))]))
            return _dafny.Set(coll10_)
        (d_276_v3_)[index44_] = (iife22_()
        ) != ((d_275_v2_) | (d_275_v2_))
        d_339_v37_: _dafny.Map
        d_339_v37_ = _dafny.Map({p2: d_274_v1_})
        d_340_v38_: _dafny.MultiSet
        d_340_v38_ = _dafny.MultiSet([len(d_275_v2_)])
        d_341_v39_: D1
        d_341_v39_ = D1_DC5((d_276_v3_)[default__.safeIndex(442, (d_276_v3_).length(0))], d_339_v37_, d_274_v1_, d_340_v38_)
        pat_let_tv16_ = d_276_v3_
        pat_let_tv17_ = d_276_v3_
        pat_let_tv18_ = d_274_v1_
        pat_let_tv19_ = d_274_v1_
        def lambda11_(source5_):
            if source5_.is_DC4:
                d_342___mcc_h8_ = source5_.cf4
                d_343___mcc_h9_ = source5_.cf5
                d_344___mcc_h10_ = source5_.cf6
                d_345_cf6_ = d_344___mcc_h10_
                d_346_cf5_ = d_343___mcc_h9_
                d_347_cf4_ = d_342___mcc_h8_
                return len((_dafny.SeqWithoutIsStrInference([d_345_cf6_ for d_348_i6_ in range(default__.abs(478))]) if (pat_let_tv17_)[default__.safeIndex(442, (pat_let_tv16_).length(0))] else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_347_cf4_}))]) for d_349_i7_ in range(default__.abs(437))])))
            elif source5_.is_DC5:
                d_350___mcc_h11_ = source5_.cf7
                d_351___mcc_h12_ = source5_.cf8
                d_352___mcc_h13_ = source5_.cf9
                d_353___mcc_h14_ = source5_.cf10
                d_354_cf10_ = d_353___mcc_h14_
                d_355_cf9_ = d_352___mcc_h13_
                d_356_cf8_ = d_351___mcc_h12_
                d_357_cf7_ = d_350___mcc_h11_
                return pat_let_tv18_
            elif True:
                d_358___mcc_h15_ = source5_.cf3
                d_359_cf3_ = d_358___mcc_h15_
                return (pat_let_tv19_) - (585)

        (globalState).f3 = lambda11_(d_341_v39_)

    @property
    def f25(self):
        return self._f25

class C2(T2, T1, T0):
    def  __init__(self):
        self._f21: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f22: bool = False
        self._f23: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f22(self):
        return self._f22
    @f22.setter
    def f22(self, value):
        self._f22 = value
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    def ctor__(self, f23, f21, f22):
        (self).f23 = f23
        (self).f21 = f21
        (self).f22 = f22

    def fm2(self, p0, globalState):
        return not ((self.f22 if False else self.f22)) or (not(False))

    def fm0(self, globalState):
        return _dafny.Set({self.f22, self.f22, False, self.f22, self.f22})

    def fm1(self, p0, p1, p2, globalState):
        return ((99) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([self.f22])).cardinality])))) == (len((_dafny.Set({self.f22, self.f22})) | (_dafny.Set({True}))))

    def fm3(self, globalState):
        return not(((-141 if self.f22 else 419)) <= (len(_dafny.SeqWithoutIsStrInference([self.f21]))))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_360_v0_: int
        d_360_v0_ = 191
        d_361_v1_: _dafny.Seq
        d_361_v1_ = _dafny.SeqWithoutIsStrInference([d_360_v0_, d_360_v0_])
        d_362_v2_: _dafny.MultiSet
        d_362_v2_ = _dafny.MultiSet([(self.f23)[default__.safeIndex(d_360_v0_, len(self.f23))], d_361_v1_])
        d_363_v3_: _dafny.Set
        d_363_v3_ = _dafny.Set({((d_362_v2_)[d_361_v1_] if (d_361_v1_) in (d_362_v2_) else d_360_v0_)})
        d_364_v5_: _dafny.Set
        def iife23_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(132, 989):
                d_365_v4_: int = compr_11_
                if ((132) <= (d_365_v4_)) and ((d_365_v4_) < (989)):
                    coll11_ = coll11_.union(_dafny.Set([(d_365_v4_) + (len(_dafny.Map({d_360_v0_: p0})))]))
            return _dafny.Set(coll11_)
        d_364_v5_ = _dafny.Set({d_363_v3_, d_363_v3_, d_363_v3_, (iife23_()
        ) - (d_363_v3_)})
        d_364_v5_ = ((_dafny.Set({d_363_v3_}) if p0 else d_364_v5_)) | (d_364_v5_)
        if True:
            (globalState).f20 = d_360_v0_
            d_366_v6_: _dafny.Map
            d_366_v6_ = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_360_v0_, len(_dafny.SeqWithoutIsStrInference([self.f22])), d_360_v0_]))).cardinality: self.f21})
            d_367_v7_: str
            d_367_v7_ = _dafny.CodePoint('u')
            d_368_v8_: _dafny.Seq
            d_368_v8_ = _dafny.SeqWithoutIsStrInference([self.f21, ((d_366_v6_)[d_360_v0_] if (d_360_v0_) in (d_366_v6_) else self.f21), _dafny.SeqWithoutIsStrInference([d_367_v7_])])
            (self).f21 = ((d_368_v8_)[default__.safeIndex(d_360_v0_, len(d_368_v8_))]).set(default__.safeIndex(d_360_v0_, len((d_368_v8_)[default__.safeIndex(d_360_v0_, len(d_368_v8_))])), d_367_v7_)
            d_369_v9_: _dafny.Array
            def lambda12_(d_370_i0_):
                return True

            init5_ = lambda12_
            nw57_ = _dafny.Array(None, 19)
            for i0_5_ in range(nw57_.length(0)):
                nw57_[i0_5_] = init5_(i0_5_)
            d_369_v9_ = nw57_
            d_371_v10_: D0
            d_371_v10_ = D0_DC0(d_369_v9_)
            d_372_v11_: _dafny.Seq
            d_372_v11_ = _dafny.SeqWithoutIsStrInference([d_369_v9_])
            d_373_v12_: _dafny.Array
            nw58_ = _dafny.Array(None, 6)
            nw58_[int(0)] = d_369_v9_
            nw58_[int(1)] = d_369_v9_
            nw58_[int(2)] = d_369_v9_
            nw58_[int(3)] = (d_371_v10_).cf0
            nw58_[int(4)] = d_369_v9_
            nw58_[int(5)] = (d_372_v11_)[default__.safeIndex(d_360_v0_, len(d_372_v11_))]
            d_373_v12_ = nw58_
            d_374_v13_: D0
            d_374_v13_ = D0_DC1(self.f22)
            d_375_v14_: _dafny.MultiSet
            d_375_v14_ = _dafny.MultiSet([self.f22])
            d_376_v15_: _dafny.Map
            d_376_v15_ = _dafny.Map({default__.fm4(self.f22, d_375_v14_, globalState): d_373_v12_})
            def iife24_(_pat_let6_0):
                def iife25_(d_377_dt__update__tmp_h0_):
                    def iife26_(_pat_let7_0):
                        def iife27_(d_378_dt__update_hcf1_h0_):
                            return D0_DC1(d_378_dt__update_hcf1_h0_)
                        return iife27_(_pat_let7_0)
                    return iife26_(False)
                return iife25_(_pat_let6_0)
            d_373_v12_ = (((d_376_v15_)[_dafny.CodePoint('x')] if (_dafny.CodePoint('x')) in (d_376_v15_) else d_373_v12_) if (iife24_(d_374_v13_)).cf1 else d_373_v12_)
            index45_ = default__.safeIndex(623, (d_373_v12_).length(0))
            (d_373_v12_)[index45_] = d_369_v9_
            index46_ = default__.safeIndex(623, (d_373_v12_).length(0))
            (d_373_v12_)[index46_] = d_369_v9_
            d_379_v16_: _dafny.MultiSet
            d_379_v16_ = _dafny.MultiSet([d_360_v0_])
            r0 = ((d_379_v16_).cardinality) * ((d_361_v1_)[default__.safeIndex(849, len(d_361_v1_))])
        elif True:
            d_380_v17_: _dafny.Map
            d_380_v17_ = _dafny.Map({p0: d_360_v0_})
            d_381_v18_: _dafny.Map
            d_381_v18_ = _dafny.Map({d_380_v17_: not (self.f22) or (self.f22)})
            d_381_v18_ = _dafny.Map({d_380_v17_: True})
            (self).f22 = self.f22
            d_382_v19_: _dafny.Array
            nw59_ = _dafny.Array(False, 16)
            d_382_v19_ = nw59_
            d_383_v20_: bool
            out4_: bool
            out4_ = default__.m0(d_382_v19_, (True) and (p0), globalState)
            d_383_v20_ = out4_
            (self).f22 = d_383_v20_
            d_384_v21_: _dafny.Seq
            d_384_v21_ = _dafny.SeqWithoutIsStrInference([d_363_v3_, d_363_v3_])
            (globalState).f20 = len((d_384_v21_) + ((d_384_v21_).set(default__.safeIndex(182, len(d_384_v21_)), _dafny.Set({d_360_v0_, (0) - (d_360_v0_)}))))
        d_385_v22_: D0
        d_385_v22_ = D0_DC1(True)
        d_386_v23_: D0
        d_386_v23_ = D0_DC2(d_385_v22_)
        d_387_v24_: _dafny.Seq
        d_387_v24_ = _dafny.SeqWithoutIsStrInference([self.f21])
        d_388_v25_: _dafny.Seq
        d_388_v25_ = _dafny.SeqWithoutIsStrInference([self.f22, self.f22, not(p0)])
        d_389_v26_: _dafny.Array
        nw60_ = _dafny.Array(None, 15)
        nw60_[int(0)] = len(d_387_v24_)
        nw60_[int(1)] = d_360_v0_
        nw60_[int(2)] = d_360_v0_
        nw60_[int(3)] = d_360_v0_
        nw60_[int(4)] = 426
        nw60_[int(5)] = default__.safeModuloInt(d_360_v0_, d_360_v0_)
        nw60_[int(6)] = (0) - (d_360_v0_)
        nw60_[int(7)] = 834
        nw60_[int(8)] = d_360_v0_
        nw60_[int(9)] = len(d_387_v24_)
        nw60_[int(10)] = d_360_v0_
        nw60_[int(11)] = -412
        nw60_[int(12)] = d_360_v0_
        nw60_[int(13)] = d_360_v0_
        nw60_[int(14)] = len(d_388_v25_)
        d_389_v26_ = nw60_
        index47_ = default__.safeIndex(925, (d_389_v26_).length(0))
        (d_389_v26_)[index47_] = (len(d_361_v1_)) * (d_360_v0_)
        index48_ = default__.safeIndex(925, (d_389_v26_).length(0))
        rhs43_ = len(self.f23)
        rhs44_ = d_360_v0_
        rhs45_ = (0) - ((0) - (d_360_v0_))
        rhs46_ = d_386_v23_
        rhs47_ = default__.safeDivisionInt(d_360_v0_, (0) - ((0) - (len((d_388_v25_) + (d_388_v25_)))))
        lhs38_ = globalState
        lhs39_ = globalState
        lhs40_ = d_389_v26_
        lhs41_ = default__.safeIndex(925, (d_389_v26_).length(0))
        lhs38_.f20 = rhs43_
        d_360_v0_ = rhs44_
        lhs39_.f20 = rhs45_
        d_386_v23_ = rhs46_
        lhs40_[lhs41_] = rhs47_
        r0 = (d_389_v26_)[default__.safeIndex(925, (d_389_v26_).length(0))]
        d_390_i1_: int
        d_390_i1_ = 0
        with _dafny.label("2"):
            while False:
                with _dafny.c_label("2"):
                    if (d_390_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_390_i1_ = (d_390_i1_) + (1)
                    d_391_v27_: str
                    d_391_v27_ = _dafny.CodePoint('i')
                    d_392_v28_: _dafny.Map
                    d_392_v28_ = _dafny.Map({(d_389_v26_)[default__.safeIndex(925, (d_389_v26_).length(0))]: d_391_v27_})
                    d_393_v29_: _dafny.Map
                    d_393_v29_ = _dafny.Map({self.f22: len(d_392_v28_)})
                    d_394_v30_: _dafny.Set
                    d_394_v30_ = _dafny.Set({default__.fm5((d_389_v26_)[default__.safeIndex(925, (d_389_v26_).length(0))], len(_dafny.SeqWithoutIsStrInference([True])), d_360_v0_, globalState)})
                    index49_ = default__.safeIndex(925, (d_389_v26_).length(0))
                    (d_389_v26_)[index49_] = ((len(d_393_v29_)) - (914)) - (len((d_394_v30_ if False else _dafny.Set({d_361_v1_, _dafny.SeqWithoutIsStrInference([d_360_v0_, (d_389_v26_)[default__.safeIndex(925, (d_389_v26_).length(0))], d_360_v0_])}))))
                    d_395_v31_: _dafny.Array
                    def lambda13_(d_396_p0_):
                        def lambda14_(d_397_i2_):
                            return _dafny.MultiSet([not(self.f22), d_396_p0_])

                        return lambda14_

                    init6_ = lambda13_(p0)
                    nw61_ = _dafny.Array(None, 25)
                    for i0_6_ in range(nw61_.length(0)):
                        nw61_[i0_6_] = init6_(i0_6_)
                    d_395_v31_ = nw61_
                    d_398_v32_: C1
                    nw62_ = C1()
                    nw62_.ctor__(d_395_v31_, True)
                    d_398_v32_ = nw62_
                    d_399_v33_: _dafny.Set
                    d_399_v33_ = _dafny.Set({d_391_v27_})
                    d_400_v34_: D1
                    d_400_v34_ = D1_DC3(_dafny.Set({d_399_v33_}))
                    d_401_v35_: _dafny.Set
                    d_401_v35_ = _dafny.Set({d_400_v34_, d_400_v34_})
                    d_402_v36_: _dafny.Set
                    d_402_v36_ = _dafny.Set({(d_401_v35_) - (d_401_v35_), _dafny.Set({default__.fm10(globalState), d_400_v34_, default__.fm10(globalState), d_400_v34_})})
                    d_402_v36_ = d_402_v36_
                    d_403_v37_: _dafny.Array
                    nw63_ = _dafny.Array(_dafny.MultiSet({}), 20)
                    d_403_v37_ = nw63_
                    d_404_v38_: _dafny.MultiSet
                    d_404_v38_ = _dafny.MultiSet([self.f21, (self.f21).set(default__.safeIndex(len(d_388_v25_), len(self.f21)), d_391_v27_), self.f21, self.f21])
                    index50_ = default__.safeIndex(214, (d_403_v37_).length(0))
                    (d_403_v37_)[index50_] = d_404_v38_
                    d_405_v39_: D0
                    d_405_v39_ = D0_DC1((d_398_v32_).f25)
                    d_406_v40_: _dafny.Set
                    d_406_v40_ = _dafny.Set({(d_405_v39_).cf1, False, not(self.f22), p0})
                    d_407_v41_: _dafny.Map
                    d_407_v41_ = _dafny.Map({d_398_v32_: d_406_v40_})
                    index51_ = default__.safeIndex(214, (d_403_v37_).length(0))
                    (d_403_v37_)[index51_] = default__.fm13(self.f22, d_406_v40_, default__.safeDivisionInt(len(self.f21), (_dafny.MultiSet([((d_407_v41_)[d_398_v32_] if (d_398_v32_) in (d_407_v41_) else d_406_v40_), d_406_v40_, _dafny.Set({False, p0, (d_398_v32_).f25, (d_398_v32_).f25, (d_398_v32_).f25})])).cardinality), p0, globalState)
                    pass
            pass
        d_408_v43_: D1
        def iife28_():
            coll12_ = _dafny.Set()
            compr_12_: str
            for compr_12_ in (_dafny.MultiSet(self.f21)).Elements:
                d_409_v42_: str = compr_12_
                if (d_409_v42_) in (_dafny.MultiSet(self.f21)):
                    coll12_ = coll12_.union(_dafny.Set([d_409_v42_]))
            return _dafny.Set(coll12_)
        d_408_v43_ = D1_DC3(_dafny.Set({iife28_()
}))
        d_410_v44_: _dafny.Map
        d_410_v44_ = _dafny.Map({d_360_v0_: p0})
        pat_let_tv20_ = d_389_v26_
        pat_let_tv21_ = d_389_v26_
        pat_let_tv22_ = d_360_v0_
        pat_let_tv23_ = d_360_v0_
        pat_let_tv24_ = d_389_v26_
        pat_let_tv25_ = d_389_v26_
        def lambda15_(source6_):
            if source6_.is_DC4:
                d_411___mcc_h0_ = source6_.cf4
                d_412___mcc_h1_ = source6_.cf5
                d_413___mcc_h2_ = source6_.cf6
                d_414_cf6_ = d_413___mcc_h2_
                d_415_cf5_ = d_412___mcc_h1_
                d_416_cf4_ = d_411___mcc_h0_
                return ((pat_let_tv21_)[default__.safeIndex(925, (pat_let_tv20_).length(0))]) - ((D1_DC5(self.f22, _dafny.Map({not(self.f22): len(d_414_cf6_)}), len(_dafny.Set({self.f22, False})), _dafny.MultiSet([pat_let_tv22_, len(self.f21)]))).cf9)
            elif source6_.is_DC5:
                d_417___mcc_h3_ = source6_.cf7
                d_418___mcc_h4_ = source6_.cf8
                d_419___mcc_h5_ = source6_.cf9
                d_420___mcc_h6_ = source6_.cf10
                d_421_cf10_ = d_420___mcc_h6_
                d_422_cf9_ = d_419___mcc_h5_
                d_423_cf8_ = d_418___mcc_h4_
                d_424_cf7_ = d_417___mcc_h3_
                return pat_let_tv23_
            elif True:
                d_425___mcc_h7_ = source6_.cf3
                d_426_cf3_ = d_425___mcc_h7_
                return ((pat_let_tv25_)[default__.safeIndex(925, (pat_let_tv24_).length(0))]) + (443)

        rhs48_ = (self.f21) + (self.f21)
        rhs49_ = lambda15_(d_408_v43_)
        rhs50_ = d_389_v26_
        rhs51_ = ((d_389_v26_)[default__.safeIndex(925, (d_389_v26_).length(0))]) - (d_360_v0_)
        rhs52_ = len((d_410_v44_) | (d_410_v44_))
        lhs42_ = self
        lhs43_ = globalState
        lhs44_ = globalState
        lhs42_.f21 = rhs48_
        r0 = rhs49_
        d_389_v26_ = rhs50_
        lhs43_.f20 = rhs51_
        lhs44_.f20 = rhs52_
        d_427_v45_: str
        d_427_v45_ = _dafny.CodePoint('s')
        d_428_v46_: _dafny.Set
        d_428_v46_ = _dafny.Set({self.f22, self.f22, p0})
        d_429_v47_: _dafny.Map
        d_429_v47_ = _dafny.Map({len(d_428_v46_): default__.fm6(globalState)})
        pat_let_tv26_ = p0
        pat_let_tv27_ = p0
        pat_let_tv28_ = p0
        pat_let_tv29_ = p0
        pat_let_tv30_ = p0
        def lambda16_(source7_):
            if source7_.is_DC4:
                d_430___mcc_h8_ = source7_.cf4
                d_431___mcc_h9_ = source7_.cf5
                d_432___mcc_h10_ = source7_.cf6
                d_433_cf6_ = d_432___mcc_h10_
                d_434_cf5_ = d_431___mcc_h9_
                d_435_cf4_ = d_430___mcc_h8_
                return (_dafny.Map({pat_let_tv26_: self.f22})) | (_dafny.Map({pat_let_tv27_: False}))
            elif source7_.is_DC5:
                d_436___mcc_h11_ = source7_.cf7
                d_437___mcc_h12_ = source7_.cf8
                d_438___mcc_h13_ = source7_.cf9
                d_439___mcc_h14_ = source7_.cf10
                d_440_cf10_ = d_439___mcc_h14_
                d_441_cf9_ = d_438___mcc_h13_
                d_442_cf8_ = d_437___mcc_h12_
                d_443_cf7_ = d_436___mcc_h11_
                return (_dafny.Map({pat_let_tv28_: d_443_cf7_})) | (_dafny.Map({False: self.f22}))
            elif True:
                d_444___mcc_h15_ = source7_.cf3
                d_445_cf3_ = d_444___mcc_h15_
                return (_dafny.Map({pat_let_tv29_: True})) | (_dafny.Map({self.f22: pat_let_tv30_}))

        r0 = len(lambda16_(D1_DC4(d_427_v45_, d_429_v47_, d_361_v1_)))
        return r0

    def m2(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_446_v0_: _dafny.Array
        nw64_ = _dafny.Array(int(0), 13)
        d_446_v0_ = nw64_
        index52_ = default__.safeIndex(931, (d_446_v0_).length(0))
        (d_446_v0_)[index52_] = len(self.f21)
        index53_ = default__.safeIndex(931, (d_446_v0_).length(0))
        (d_446_v0_)[index53_] = (p0) - (p0)
        d_447_v1_: _dafny.Seq
        d_447_v1_ = _dafny.SeqWithoutIsStrInference([self.f22])
        d_448_v2_: _dafny.Set
        d_448_v2_ = _dafny.Set({True})
        d_447_v1_ = _dafny.SeqWithoutIsStrInference([((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]) not in (_dafny.Map({350: len(d_448_v2_)})), self.f22, (self.f22) and (self.f22)])
        if self.f22:
            d_449_v3_: _dafny.MultiSet
            d_449_v3_ = _dafny.MultiSet([self.f22])
            (globalState).f3 = (default__.safeModuloInt(default__.fm6(globalState), (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]) if not((self.f22 if not(self.f22) else self.f22)) else (d_449_v3_).cardinality)
            d_450_v4_: str
            d_450_v4_ = _dafny.CodePoint('u')
            d_451_v5_: _dafny.Map
            d_451_v5_ = _dafny.Map({(d_449_v3_).cardinality: (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]})
            d_452_v6_: _dafny.Seq
            d_452_v6_ = _dafny.SeqWithoutIsStrInference([(d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]])
            d_453_v7_: D1
            d_453_v7_ = D1_DC4(d_450_v4_, d_451_v5_, d_452_v6_)
            d_453_v7_ = d_453_v7_
            d_454_v8_: _dafny.Map
            d_454_v8_ = _dafny.Map({default__.fm12((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))], _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f22])), self.f22, _dafny.MultiSet([_dafny.MultiSet([False, not(True)]), d_449_v3_, d_449_v3_, _dafny.MultiSet([self.f22, self.f22, self.f22]), d_449_v3_]), globalState): (0) - ((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))])})
            d_455_v9_: _dafny.Map
            d_455_v9_ = _dafny.Map({self.f22: False})
            d_456_v10_: D0
            d_456_v10_ = D0_DC1(((d_455_v9_)[True] if (True) in (d_455_v9_) else self.f22))
            d_454_v8_ = (d_454_v8_).set((d_456_v10_).cf1, (0) - (p0))
            d_457_v11_: _dafny.Array
            nw65_ = _dafny.Array(None, 27)
            nw65_[int(0)] = self.f22
            nw65_[int(1)] = self.f22
            nw65_[int(2)] = self.f22
            nw65_[int(3)] = self.f22
            nw65_[int(4)] = self.f22
            nw65_[int(5)] = self.f22
            nw65_[int(6)] = self.f22
            nw65_[int(7)] = self.f22
            nw65_[int(8)] = False
            nw65_[int(9)] = self.f22
            nw65_[int(10)] = self.f22
            nw65_[int(11)] = True
            nw65_[int(12)] = self.f22
            nw65_[int(13)] = self.f22
            nw65_[int(14)] = self.f22
            nw65_[int(15)] = self.f22
            nw65_[int(16)] = self.f22
            nw65_[int(17)] = True
            nw65_[int(18)] = self.f22
            nw65_[int(19)] = False
            nw65_[int(20)] = self.f22
            nw65_[int(21)] = self.f22
            nw65_[int(22)] = not(self.f22)
            nw65_[int(23)] = self.f22
            nw65_[int(24)] = self.f22
            nw65_[int(25)] = False
            nw65_[int(26)] = self.f22
            d_457_v11_ = nw65_
            d_458_v12_: D0
            d_458_v12_ = D0_DC0(d_457_v11_)
            d_459_v13_: D0
            d_459_v13_ = D0_DC2(d_458_v12_)
            d_460_v14_: _dafny.Map
            d_460_v14_ = _dafny.Map({d_459_v13_: len(self.f21)})
            d_460_v14_ = (d_460_v14_).set(d_459_v13_, p0)
            r1 = self.f22
        elif True:
            (globalState).f20 = ((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etxvn"))))
            (globalState).f12 = _dafny.CodePoint('d')
            d_461_v15_: _dafny.Array
            def lambda17_(d_462_v0_):
                def lambda18_(d_463_i0_):
                    return _dafny.Set({(d_462_v0_)[default__.safeIndex(931, (d_462_v0_).length(0))]})

                return lambda18_

            init7_ = lambda17_(d_446_v0_)
            nw66_ = _dafny.Array(None, 15)
            for i0_7_ in range(nw66_.length(0)):
                nw66_[i0_7_] = init7_(i0_7_)
            d_461_v15_ = nw66_
            d_464_v16_: _dafny.Set
            d_464_v16_ = _dafny.Set({(d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))], p0})
            index54_ = default__.safeIndex(983, (d_461_v15_).length(0))
            (d_461_v15_)[index54_] = (d_464_v16_)
            d_465_v17_: _dafny.Map
            d_465_v17_ = _dafny.Map({(self).fm1(self.f21, True, False, globalState): self.f22})
            index55_ = default__.safeIndex(983, (d_461_v15_).length(0))
            (d_461_v15_)[index55_] = default__.fm14(d_465_v17_, p0, globalState)
            source8_ = d_464_v16_
            d_466___mcc_h0_ = source8_
            d_467_cf11_ = d_466___mcc_h0_
            (globalState).f20 = (default__.fm6(globalState)) + (default__.safeDivisionInt(len(self.f21), (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]))
            d_464_v16_ = d_464_v16_
            d_465_v17_ = d_465_v17_
            (self).f22 = ((self.f22 if self.f22 else self.f22)) and ((self).fm2(self.f22, globalState))
            d_468_v18_: _dafny.Array
            nw67_ = _dafny.Array(_dafny.MultiSet({}), 4)
            d_468_v18_ = nw67_
            d_469_v19_: _dafny.Map
            d_469_v19_ = _dafny.Map({d_468_v18_: (_dafny.Map({p0: len(self.f21)})) | (_dafny.Map({len(self.f21): p0}))})
            d_470_v20_: _dafny.Map
            d_470_v20_ = _dafny.Map({(d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]: p0})
            d_469_v19_ = (d_469_v19_).set(d_468_v18_, d_470_v20_)
        if True:
            d_471_v21_: _dafny.Set
            d_471_v21_ = _dafny.Set({(d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]})
            d_472_v22_: _dafny.Set
            d_472_v22_ = (d_471_v21_).intersection(d_471_v21_)
            d_472_v22_ = d_472_v22_
            d_473_v23_: _dafny.Set
            d_473_v23_ = _dafny.Set({self.f21, self.f21})
            d_474_v24_: C0
            nw68_ = C0()
            nw68_.ctor__(d_473_v23_)
            d_474_v24_ = nw68_
            d_474_v24_ = d_474_v24_
            d_475_v25_: _dafny.Array
            def lambda19_(d_476_i1_):
                return (self.f22 if self.f22 else self.f22)

            init8_ = lambda19_
            nw69_ = _dafny.Array(None, 15)
            for i0_8_ in range(nw69_.length(0)):
                nw69_[i0_8_] = init8_(i0_8_)
            d_475_v25_ = nw69_
            index56_ = default__.safeIndex(184, (d_475_v25_).length(0))
            (d_475_v25_)[index56_] = self.f22
            d_477_v26_: _dafny.MultiSet
            d_477_v26_ = _dafny.MultiSet([self.f22, (self).fm1(self.f21, self.f22, self.f22, globalState)])
            index57_ = default__.safeIndex(184, (d_475_v25_).length(0))
            (d_475_v25_)[index57_] = not(((211) == ((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))])) == ((d_477_v26_).isdisjoint(_dafny.MultiSet(d_447_v1_))))
            r1 = self.f22
            d_478_v27_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.Map({}), 14)
            d_478_v27_ = nw70_
            d_479_v28_: _dafny.MultiSet
            d_479_v28_ = _dafny.MultiSet([self.f21, default__.fm11(False, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdbdchmi"))])
            index58_ = default__.safeIndex(184, (d_475_v25_).length(0))
            rhs53_ = ((d_479_v28_)[self.f21] if (self.f21) in (d_479_v28_) else (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))])
            rhs54_ = (d_475_v25_)[default__.safeIndex(184, (d_475_v25_).length(0))]
            rhs55_ = (d_475_v25_)[default__.safeIndex(184, (d_475_v25_).length(0))]
            rhs56_ = d_478_v27_
            lhs45_ = globalState
            lhs46_ = d_475_v25_
            lhs47_ = default__.safeIndex(184, (d_475_v25_).length(0))
            lhs48_ = self
            lhs45_.f20 = rhs53_
            lhs46_[lhs47_] = rhs54_
            lhs48_.f22 = rhs55_
            d_478_v27_ = rhs56_
        elif True:
            d_480_v29_: _dafny.Array
            nw71_ = _dafny.Array(False, 8)
            d_480_v29_ = nw71_
            d_480_v29_ = d_480_v29_
            d_481_v30_: _dafny.MultiSet
            d_481_v30_ = _dafny.MultiSet([self.f22, False])
            d_482_v32_: _dafny.Set
            d_482_v32_ = _dafny.Set({d_481_v30_})
            d_483_v33_: _dafny.Seq
            def iife29_():
                coll13_ = _dafny.Map()
                compr_13_: _dafny.MultiSet
                for compr_13_ in (d_482_v32_).Elements:
                    d_484_v31_: _dafny.MultiSet = compr_13_
                    if (d_484_v31_) in (d_482_v32_):
                        coll13_[d_484_v31_] = (0) - (p0)
                return _dafny.Map(coll13_)
            d_483_v33_ = _dafny.SeqWithoutIsStrInference([iife29_()
            ])
            d_485_v34_: _dafny.Map
            d_485_v34_ = _dafny.Map({d_481_v30_: (d_483_v33_)[default__.safeIndex((0) - ((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]), len(d_483_v33_))]})
            d_485_v34_ = (d_485_v34_).set((_dafny.MultiSet([self.f22])).intersection(d_481_v30_), (d_483_v33_)[default__.safeIndex((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))], len(d_483_v33_))])
            d_447_v1_ = (d_447_v1_) + (d_447_v1_)
            if not(self.f22):
                index59_ = default__.safeIndex(931, (d_446_v0_).length(0))
                rhs57_ = (0) - (((((d_481_v30_)[self.f22] if (self.f22) in (d_481_v30_) else -936)) + ((d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]) if self.f22 else p0))
                rhs58_ = self.f22
                rhs59_ = d_446_v0_
                lhs49_ = d_446_v0_
                lhs50_ = default__.safeIndex(931, (d_446_v0_).length(0))
                lhs51_ = self
                lhs49_[lhs50_] = rhs57_
                lhs51_.f22 = rhs58_
                d_446_v0_ = rhs59_
                d_486_v35_: _dafny.Map
                d_486_v35_ = _dafny.Map({self.f22: d_480_v29_})
                d_487_v36_: D0
                d_487_v36_ = D0_DC1(self.f22)
                d_488_v37_: _dafny.Map
                d_488_v37_ = _dafny.Map({((d_486_v35_)[self.f22] if (self.f22) in (d_486_v35_) else d_480_v29_): d_487_v36_})
                (globalState).f20 = len((d_488_v37_) | (d_488_v37_))
                r1 = (d_447_v1_)[default__.safeIndex(p0, len(d_447_v1_))]
                d_489_v38_: str
                d_489_v38_ = _dafny.CodePoint('x')
                d_490_v39_: _dafny.MultiSet
                d_490_v39_ = _dafny.MultiSet([d_489_v38_, d_489_v38_])
                rhs60_ = (d_490_v39_) - (d_490_v39_)
                rhs61_ = self.f22
                lhs52_ = self
                d_490_v39_ = rhs60_
                lhs52_.f22 = rhs61_
                d_491_v40_: _dafny.Map
                d_491_v40_ = _dafny.Map({(self).fm3(globalState): self.f22})
                d_492_v41_: _dafny.Seq
                d_492_v41_ = _dafny.SeqWithoutIsStrInference([p0])
                d_493_v42_: _dafny.Set
                d_493_v42_ = _dafny.Set({len(d_492_v41_)})
                d_491_v40_ = (d_491_v40_).set((_dafny.Set({28, p0, (d_481_v30_).cardinality})).isdisjoint(d_493_v42_), (self).fm3(globalState))
            elif True:
                d_494_v43_: _dafny.Map
                d_494_v43_ = _dafny.Map({(d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]: self.f22})
                (self).f22 = (len(self.f21)) <= (default__.safeModuloInt(len(d_494_v43_), (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))]))
                r1 = not((not(self.f22) if self.f22 else (self.f22 if ((d_494_v43_)[643] if (643) in (d_494_v43_) else self.f22) else self.f22)))
                d_495_v44_: str
                d_495_v44_ = _dafny.CodePoint('g')
                d_496_v45_: _dafny.Set
                d_496_v45_ = _dafny.Set({d_495_v44_, d_495_v44_})
                d_497_v46_: D1
                d_497_v46_ = D1_DC3(_dafny.Set({d_496_v45_, _dafny.Set({d_495_v44_})}))
                d_498_v47_: _dafny.Seq
                d_498_v47_ = _dafny.SeqWithoutIsStrInference([d_497_v46_, default__.fm10(globalState)])
                (globalState).f20 = len(d_498_v47_)
                d_499_v48_: _dafny.Array
                nw72_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_499_v48_ = nw72_
                d_500_v49_: _dafny.Array
                def lambda20_(d_501_v30_):
                    def lambda21_(d_502_i2_):
                        return d_501_v30_

                    return lambda21_

                init9_ = lambda20_(d_481_v30_)
                nw73_ = _dafny.Array(None, 6)
                for i0_9_ in range(nw73_.length(0)):
                    nw73_[i0_9_] = init9_(i0_9_)
                d_500_v49_ = nw73_
                index60_ = default__.safeIndex(40, (d_499_v48_).length(0))
                (d_499_v48_)[index60_] = d_500_v49_
                index61_ = default__.safeIndex(40, (d_499_v48_).length(0))
                (d_499_v48_)[index61_] = d_500_v49_
                d_480_v29_ = d_480_v29_
            d_503_v50_: _dafny.Array
            def lambda22_(d_504_v30_, d_505_p0_):
                def lambda23_(d_506_i3_):
                    return (d_504_v30_).set(self.f22, default__.abs(d_505_p0_))

                return lambda23_

            init10_ = lambda22_(d_481_v30_, p0)
            nw74_ = _dafny.Array(None, 2)
            for i0_10_ in range(nw74_.length(0)):
                nw74_[i0_10_] = init10_(i0_10_)
            d_503_v50_ = nw74_
            d_507_v51_: C1
            nw75_ = C1()
            nw75_.ctor__(d_503_v50_, self.f22)
            d_507_v51_ = nw75_
            d_508_v52_: _dafny.Seq
            d_508_v52_ = _dafny.SeqWithoutIsStrInference([d_507_v51_, d_507_v51_])
            r1 = (not(self.f22)) == ((_dafny.Set({d_508_v52_, d_508_v52_})).ispropersubset(_dafny.Set({d_508_v52_, d_508_v52_, d_508_v52_})))
        d_509_v53_: _dafny.Set
        d_509_v53_ = _dafny.Set({self.f21, self.f21})
        d_510_v54_: C0
        nw76_ = C0()
        nw76_.ctor__(d_509_v53_)
        d_510_v54_ = nw76_
        d_511_i4_: int
        d_511_i4_ = 0
        with _dafny.label("3"):
            while (p0) >= (p0):
                with _dafny.c_label("3"):
                    if (d_511_i4_) >= (100):
                        raise _dafny.Break("3")
                    d_511_i4_ = (d_511_i4_) + (1)
                    d_512_v55_: _dafny.MultiSet
                    d_512_v55_ = _dafny.MultiSet([self.f22])
                    d_513_v56_: D3
                    d_513_v56_ = D3_DC7(_dafny.MultiSet([self.f22]))
                    d_512_v55_ = (d_513_v56_).cf12
                    index62_ = default__.safeIndex(931, (d_446_v0_).length(0))
                    (d_446_v0_)[index62_] = p0
                    (globalState).f20 = default__.fm6(globalState)
                    d_514_v57_: _dafny.Set
                    d_514_v57_ = _dafny.Set({p0, (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))], p0})
                    d_515_v58_: _dafny.Map
                    d_515_v58_ = _dafny.Map({(self.f22 if self.f22 else self.f22): self.f22})
                    d_516_v59_: D0
                    d_516_v59_ = D0_DC1(True)
                    d_517_v60_: D0
                    d_517_v60_ = D0_DC2(d_516_v59_)
                    rhs62_ = (default__.fm14(_dafny.Map({((d_515_v58_)[not(self.f22)] if (not(self.f22)) in (d_515_v58_) else self.f22): self.f22}), len(self.f21), globalState)) - (d_514_v57_)
                    rhs63_ = (d_447_v1_)[default__.safeIndex(len(default__.fm15(globalState)), len(d_447_v1_))]
                    rhs64_ = (d_515_v58_ if self.f22 else d_515_v58_)
                    rhs65_ = d_517_v60_
                    d_514_v57_ = rhs62_
                    r1 = rhs63_
                    d_515_v58_ = rhs64_
                    d_517_v60_ = rhs65_
                    pass
            pass
        d_518_v62_: _dafny.Set
        def iife30_():
            coll14_ = _dafny.Set()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(-25, 428):
                d_519_v61_: int = compr_14_
                if ((-25) <= (d_519_v61_)) and ((d_519_v61_) < (428)):
                    coll14_ = coll14_.union(_dafny.Set([default__.safeModuloInt(d_519_v61_, (d_446_v0_)[default__.safeIndex(931, (d_446_v0_).length(0))])]))
            return _dafny.Set(coll14_)
        d_518_v62_ = iife30_()
        
        def lambda24_(source9_):
            d_520___mcc_h1_ = source9_
            d_521_cf11_ = d_520___mcc_h1_
            return self.f21

        r0 = lambda24_(d_518_v62_)
        d_522_v63_: _dafny.MultiSet
        d_522_v63_ = _dafny.MultiSet([self.f22, self.f22, True, self.f22, self.f22])
        d_523_v64_: _dafny.Map
        d_523_v64_ = _dafny.Map({d_447_v1_: p0})
        d_524_v65_: _dafny.Array
        nw77_ = _dafny.Array(False, 1)
        d_524_v65_ = nw77_
        d_525_v66_: _dafny.Seq
        d_525_v66_ = _dafny.SeqWithoutIsStrInference([d_524_v65_, d_524_v65_])
        d_526_v67_: _dafny.Map
        d_526_v67_ = _dafny.Map({len(d_523_v64_): (d_525_v66_)[default__.safeIndex(p0, len(d_525_v66_))]})
        d_527_v68_: _dafny.MultiSet
        d_527_v68_ = _dafny.MultiSet([_dafny.MultiSet(d_447_v1_)])
        r1 = default__.fm12(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsldowxme"))), (d_522_v63_ if not(self.f22) else d_522_v63_), (d_526_v67_) == (_dafny.Map({88: d_524_v65_})), d_527_v68_, globalState)
        return r0, r1


class C3(T0):
    def  __init__(self):
        self._f21: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f22: bool = False
        self.f28: bool = False
        self._f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f22(self):
        return self._f22
    @f22.setter
    def f22(self, value):
        self._f22 = value
    def ctor__(self, f27, f28, f21, f22):
        (self)._f27 = f27
        (self).f28 = f28
        (self).f21 = f21
        (self).f22 = f22

    def fm0(self, globalState):
        return ((_dafny.Set({self.f22, True}) if False else _dafny.Set({(self).f27, self.f28}))) | (_dafny.Set({(self).f27}))

    def fm1(self, p0, p1, p2, globalState):
        return (self).f27

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        (self).f22 = False
        hi3_ = p1
        for d_528_i0_ in range(len(self.f21), hi3_):
            d_529_v0_: _dafny.Array
            nw78_ = _dafny.Array(None, 3)
            d_529_v0_ = nw78_
            d_530_v1_: _dafny.MultiSet
            d_530_v1_ = _dafny.MultiSet([d_529_v0_, d_529_v0_])
            d_530_v1_ = ((d_530_v1_).intersection(d_530_v1_)) | ((d_530_v1_).intersection(d_530_v1_))
            (globalState).f6 = self.f21
            if (self).f27:
                d_531_v2_: _dafny.Map
                d_531_v2_ = _dafny.Map({self.f22: p1})
                d_532_v3_: _dafny.Set
                d_532_v3_ = _dafny.Set({d_531_v2_, d_531_v2_})
                (self).f22 = ((_dafny.Set({_dafny.Map({self.f28: p3}), d_531_v2_, d_531_v2_, _dafny.Map({self.f28: p2})})) | (d_532_v3_)).isdisjoint(d_532_v3_)
                d_533_v4_: _dafny.Set
                d_533_v4_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cafmfqmaf")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_534_i1_ in range(default__.abs(-524))])})
                d_535_v5_: C0
                nw79_ = C0()
                nw79_.ctor__(d_533_v4_)
                d_535_v5_ = nw79_
                d_536_v6_: D4
                d_536_v6_ = D4_DC10(d_535_v5_)
                rhs66_ = p3
                rhs67_ = (d_536_v6_).cf20
                lhs53_ = globalState
                lhs53_.f20 = rhs66_
                d_535_v5_ = rhs67_
                d_537_v7_: _dafny.Seq
                d_537_v7_ = _dafny.SeqWithoutIsStrInference([p1, default__.fm6(globalState), p2])
                d_537_v7_ = d_537_v7_
                d_538_v8_: _dafny.Array
                nw80_ = _dafny.Array(False, 10)
                d_538_v8_ = nw80_
                rhs68_ = d_538_v8_
                rhs69_ = (default__.safeDivisionInt(d_528_i0_, d_528_i0_) if True else default__.safeDivisionInt(d_528_i0_, p3))
                lhs54_ = globalState
                d_538_v8_ = rhs68_
                lhs54_.f3 = rhs69_
                (self).f22 = False
            elif True:
                d_539_v9_: _dafny.MultiSet
                d_539_v9_ = _dafny.MultiSet([p3, p0])
                d_540_v10_: _dafny.Map
                d_540_v10_ = _dafny.Map({self.f22: (d_539_v9_).cardinality})
                d_541_v11_: _dafny.Seq
                d_541_v11_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({(self).f27: p0})) | (d_540_v10_)])
                d_541_v11_ = d_541_v11_
                d_542_v12_: _dafny.Set
                d_542_v12_ = _dafny.Set({p1, -257, 555, 879, d_528_i0_})
                d_543_v14_: _dafny.Map
                def iife31_():
                    coll15_ = _dafny.Set()
                    compr_15_: int
                    for compr_15_ in (_dafny.SeqWithoutIsStrInference([711 for d_544_i2_ in range(default__.abs(502))])).Elements:
                        d_545_v13_: int = compr_15_
                        if (d_545_v13_) in (_dafny.SeqWithoutIsStrInference([711 for d_544_i2_ in range(default__.abs(502))])):
                            coll15_ = coll15_.union(_dafny.Set([(d_545_v13_) + (-762)]))
                    return _dafny.Set(coll15_)
                d_543_v14_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkm"))): iife31_()
                })
                d_546_v15_: _dafny.Set
                d_546_v15_ = _dafny.Set({d_542_v12_, ((d_543_v14_)[p1] if (p1) in (d_543_v14_) else d_542_v12_)})
                (globalState).f3 = len((d_546_v15_) | (_dafny.Set({_dafny.Set({p0, len(d_542_v12_), p2})})))
                d_547_v16_: D3
                d_547_v16_ = D3_DC7(_dafny.MultiSet([(self).f27, self.f28]))
                d_548_v17_: _dafny.Map
                d_548_v17_ = _dafny.Map({p3: d_547_v16_})
                d_548_v17_ = (d_548_v17_) | (d_548_v17_)
                (self).f22 = ((self).f27) or ((self).f27)
                d_549_v18_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_549_v18_ = nw81_
                index63_ = default__.safeIndex(391, (d_549_v18_).length(0))
                (d_549_v18_)[index63_] = self.f21
                index64_ = default__.safeIndex(391, (d_549_v18_).length(0))
                (d_549_v18_)[index64_] = self.f21
            d_550_v19_: _dafny.Array
            def lambda25_(d_551_i3_):
                return (self.f21) == (self.f21)

            init11_ = lambda25_
            nw82_ = _dafny.Array(None, 5)
            for i0_11_ in range(nw82_.length(0)):
                nw82_[i0_11_] = init11_(i0_11_)
            d_550_v19_ = nw82_
            index65_ = default__.safeIndex(792, (d_550_v19_).length(0))
            (d_550_v19_)[index65_] = (self).f27
            index66_ = default__.safeIndex(792, (d_550_v19_).length(0))
            (d_550_v19_)[index66_] = (self.f22) or ((self).f27)
        d_552_v20_: _dafny.Array
        nw83_ = _dafny.Array(int(0), 6)
        d_552_v20_ = nw83_
        d_553_v21_: _dafny.Array
        nw84_ = _dafny.Array(False, 23)
        d_553_v21_ = nw84_
        d_554_v22_: _dafny.Seq
        d_554_v22_ = _dafny.SeqWithoutIsStrInference([self.f28, False, True, self.f22, self.f22])
        index67_ = default__.safeIndex(690, (d_553_v21_).length(0))
        (d_553_v21_)[index67_] = (d_554_v22_)[default__.safeIndex(p0, len(d_554_v22_))]
        index68_ = default__.safeIndex(791, (d_553_v21_).length(0))
        (d_553_v21_)[index68_] = (self).f27
        index69_ = default__.safeIndex(690, (d_553_v21_).length(0))
        index70_ = default__.safeIndex(791, (d_553_v21_).length(0))
        rhs70_ = d_552_v20_
        rhs71_ = self.f22
        rhs72_ = (d_554_v22_)[default__.safeIndex(-531, len(d_554_v22_))]
        rhs73_ = ((self.f28) or (self.f22) if self.f22 else True)
        lhs55_ = d_553_v21_
        lhs56_ = default__.safeIndex(690, (d_553_v21_).length(0))
        lhs57_ = d_553_v21_
        lhs58_ = default__.safeIndex(791, (d_553_v21_).length(0))
        lhs59_ = self
        d_552_v20_ = rhs70_
        lhs55_[lhs56_] = rhs71_
        lhs57_[lhs58_] = rhs72_
        lhs59_.f28 = rhs73_
        d_555_v23_: _dafny.Set
        d_555_v23_ = _dafny.Set({((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2, p3]))).set(p2, default__.abs(p0))).cardinality, p1})
        d_556_v24_: _dafny.Set
        d_556_v24_ = d_555_v23_
        source10_ = d_556_v24_
        d_557___mcc_h0_ = source10_
        d_558_cf11_ = d_557___mcc_h0_
        (globalState).f3 = p3
        d_559_v25_: _dafny.Seq
        d_559_v25_ = _dafny.SeqWithoutIsStrInference([d_552_v20_, d_552_v20_, (d_552_v20_ if self.f28 else d_552_v20_), d_552_v20_, d_552_v20_])
        d_559_v25_ = d_559_v25_
        (self).f28 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_560_i4_ in range(default__.abs(336))])) <= (self.f21)
        d_561_v26_: str
        d_561_v26_ = _dafny.CodePoint('t')
        d_562_v27_: _dafny.Set
        d_562_v27_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_561_v26_ for d_563_i5_ in range(default__.abs(468))])})
        d_564_v28_: _dafny.Map
        d_564_v28_ = _dafny.Map({d_561_v26_: d_562_v27_})
        d_565_v29_: C0
        nw85_ = C0()
        nw85_.ctor__(((d_564_v28_)[d_561_v26_] if (d_561_v26_) in (d_564_v28_) else _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ce"))})))
        d_565_v29_ = nw85_
        d_566_v31_: _dafny.Map
        d_566_v31_ = _dafny.Map({p0: False})
        def iife32_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(680, 613):
                d_567_v30_: int = compr_16_
                if ((680) <= (d_567_v30_)) and ((d_567_v30_) < (613)):
                    coll16_[(d_567_v30_) + (p2)] = (self).f27
            return _dafny.Map(coll16_)
        if (_dafny.SeqWithoutIsStrInference([iife32_()
 for d_568_i6_ in range(default__.abs(526))])) < (_dafny.SeqWithoutIsStrInference([d_566_v31_, d_566_v31_, _dafny.Map({p1: (self).f27}), d_566_v31_])):
            d_569_v32_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Seq({}), 5)
            d_569_v32_ = nw86_
            index71_ = default__.safeIndex(617, (d_569_v32_).length(0))
            (d_569_v32_)[index71_] = (d_554_v22_) + (_dafny.SeqWithoutIsStrInference([self.f22, (d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))], (d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))], (d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]]))
            index72_ = default__.safeIndex(617, (d_569_v32_).length(0))
            (d_569_v32_)[index72_] = ((d_554_v22_ if not(self.f22) else d_554_v22_)) + (d_554_v22_)
            d_570_v33_: _dafny.Array
            def lambda26_(d_571_i7_):
                return _dafny.MultiSet([not(self.f22), self.f22])

            init12_ = lambda26_
            nw87_ = _dafny.Array(None, 26)
            for i0_12_ in range(nw87_.length(0)):
                nw87_[i0_12_] = init12_(i0_12_)
            d_570_v33_ = nw87_
            d_572_v34_: C1
            nw88_ = C1()
            nw88_.ctor__(d_570_v33_, (self).f27)
            d_572_v34_ = nw88_
            index73_ = default__.safeIndex(331, (d_552_v20_).length(0))
            (d_552_v20_)[index73_] = p2
            d_573_v35_: _dafny.Seq
            d_573_v35_ = _dafny.SeqWithoutIsStrInference([len(d_554_v22_), len(self.f21)])
            d_574_v36_: _dafny.Seq
            d_574_v36_ = _dafny.SeqWithoutIsStrInference([244, p3, (_dafny.MultiSet(d_573_v35_)).cardinality, len(d_554_v22_), p2])
            d_575_v37_: _dafny.MultiSet
            d_575_v37_ = _dafny.MultiSet([886])
            d_576_v38_: _dafny.Map
            d_576_v38_ = _dafny.Map({(d_575_v37_).cardinality: p0})
            d_577_v39_: _dafny.MultiSet
            d_577_v39_ = _dafny.MultiSet([(d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]])
            d_578_v40_: _dafny.MultiSet
            d_578_v40_ = _dafny.MultiSet([d_577_v39_, d_577_v39_, d_577_v39_, d_577_v39_, d_577_v39_])
            d_579_v41_: _dafny.Map
            d_579_v41_ = _dafny.Map({(self).f27: False})
            index74_ = default__.safeIndex(331, (d_552_v20_).length(0))
            rhs74_ = d_572_v34_
            rhs75_ = ((self.f21)[default__.safeIndex(p0, len(self.f21))] if (d_574_v36_) < (default__.fm5(p3, len(d_576_v38_), p3, globalState)) else _dafny.CodePoint('k'))
            rhs76_ = (default__.fm6(globalState)) - (len((d_569_v32_)[default__.safeIndex(617, (d_569_v32_).length(0))]))
            rhs77_ = (default__.fm6(globalState) if default__.fm12(p1, d_577_v39_, self.f28, d_578_v40_, globalState) else ((d_574_v36_)[default__.safeIndex(((d_577_v39_)[(d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]] if ((d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]) in (d_577_v39_) else (0) - (len(d_579_v41_))), len(d_574_v36_))]) + (p3))
            lhs60_ = globalState
            lhs61_ = d_552_v20_
            lhs62_ = default__.safeIndex(331, (d_552_v20_).length(0))
            lhs63_ = globalState
            d_572_v34_ = rhs74_
            lhs60_.f12 = rhs75_
            lhs61_[lhs62_] = rhs76_
            lhs63_.f3 = rhs77_
            d_580_v42_: str
            d_580_v42_ = _dafny.CodePoint('e')
            d_581_v43_: _dafny.Set
            d_581_v43_ = _dafny.Set({_dafny.CodePoint('e'), d_580_v42_, d_580_v42_, d_580_v42_})
            d_582_v44_: _dafny.Set
            d_582_v44_ = _dafny.Set({d_581_v43_})
            d_583_v45_: D1
            d_583_v45_ = D1_DC3(d_582_v44_)
            source11_ = d_583_v45_
            if source11_.is_DC4:
                d_584___mcc_h1_ = source11_.cf4
                d_585___mcc_h2_ = source11_.cf5
                d_586___mcc_h3_ = source11_.cf6
                d_587_cf6_ = d_586___mcc_h3_
                d_588_cf5_ = d_585___mcc_h2_
                d_589_cf4_ = d_584___mcc_h1_
                index75_ = default__.safeIndex(690, (d_553_v21_).length(0))
                (d_553_v21_)[index75_] = (d_572_v34_).f25
                d_589_cf4_ = d_580_v42_
                pat_let_tv31_ = d_582_v44_
                def iife33_(_pat_let8_0):
                    def iife34_(d_590_dt__update__tmp_h0_):
                        def iife35_(_pat_let9_0):
                            def iife36_(d_591_dt__update_hcf3_h0_):
                                return D1_DC3(d_591_dt__update_hcf3_h0_)
                            return iife36_(_pat_let9_0)
                        return iife35_(pat_let_tv31_)
                    return iife34_(_pat_let8_0)
                d_583_v45_ = iife33_(d_583_v45_)
                d_592_v46_: C1
                d_592_v46_ = d_572_v34_
                rhs78_ = (d_592_v46_)
                rhs79_ = p3
                rhs80_ = ((self).f27 if self.f28 else (self.f21) < (_dafny.SeqWithoutIsStrInference([d_580_v42_ for d_593_i8_ in range(default__.abs(-738))])))
                lhs64_ = globalState
                lhs65_ = self
                d_572_v34_ = rhs78_
                lhs64_.f3 = rhs79_
                lhs65_.f22 = rhs80_
            elif source11_.is_DC5:
                d_594___mcc_h4_ = source11_.cf7
                d_595___mcc_h5_ = source11_.cf8
                d_596___mcc_h6_ = source11_.cf9
                d_597___mcc_h7_ = source11_.cf10
                d_598_cf10_ = d_597___mcc_h7_
                d_599_cf9_ = d_596___mcc_h6_
                d_600_cf8_ = d_595___mcc_h5_
                d_601_cf7_ = d_594___mcc_h4_
                (self).f22 = (d_577_v39_).ispropersubset(_dafny.MultiSet(((d_569_v32_)[default__.safeIndex(617, (d_569_v32_).length(0))]) + ((d_569_v32_)[default__.safeIndex(617, (d_569_v32_).length(0))])))
                index76_ = default__.safeIndex(690, (d_553_v21_).length(0))
                (d_553_v21_)[index76_] = self.f22
                (self).f28 = d_601_cf7_
                d_602_v47_: _dafny.Seq
                d_602_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_554_v22_: self.f22}))]), d_574_v36_, d_574_v36_])
                d_603_v48_: C2
                nw89_ = C2()
                nw89_.ctor__(d_602_v47_, self.f21, ((d_574_v36_)[default__.safeIndex(len((d_569_v32_)[default__.safeIndex(617, (d_569_v32_).length(0))]), len(d_574_v36_))]) >= (645))
                d_603_v48_ = nw89_
            elif True:
                d_604___mcc_h8_ = source11_.cf3
                d_605_cf3_ = d_604___mcc_h8_
                d_606_v49_: _dafny.Set
                d_606_v49_ = _dafny.Set({not((self).f27), False, not((d_572_v34_).f25)})
                index77_ = default__.safeIndex(331, (d_552_v20_).length(0))
                (d_552_v20_)[index77_] = ((len(d_606_v49_)) + (-765)) - (p1)
                index78_ = default__.safeIndex(690, (d_553_v21_).length(0))
                rhs81_ = (0) - ((((d_552_v20_)[default__.safeIndex(331, (d_552_v20_).length(0))]) * (876)) - (default__.safeModuloInt(len(_dafny.Map({False: len(d_554_v22_)})), default__.fm6(globalState))))
                rhs82_ = (d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]
                rhs83_ = default__.fm4(((0) - ((d_552_v20_)[default__.safeIndex(331, (d_552_v20_).length(0))])) not in (_dafny.SeqWithoutIsStrInference([p1])), _dafny.MultiSet([(d_572_v34_).f25, not(self.f28), (self).f27, (self).f27, (d_572_v34_).f25]), globalState)
                lhs66_ = globalState
                lhs67_ = d_553_v21_
                lhs68_ = default__.safeIndex(690, (d_553_v21_).length(0))
                lhs66_.f20 = rhs81_
                lhs67_[lhs68_] = rhs82_
                d_580_v42_ = rhs83_
                (self).f22 = not((d_572_v34_).f25)
                (globalState).f3 = p0
            d_607_v50_: _dafny.Map
            d_607_v50_ = _dafny.Map({d_580_v42_: len(self.f21)})
            d_607_v50_ = (d_607_v50_).set(d_580_v42_, default__.safeModuloInt(p0, p0))
            (self).f22 = self.f28
        elif True:
            d_608_v51_: str
            d_608_v51_ = _dafny.CodePoint('r')
            d_609_v52_: _dafny.Map
            d_609_v52_ = _dafny.Map({p1: d_608_v51_})
            d_610_v53_: _dafny.Array
            nw90_ = _dafny.Array(None, 19)
            nw90_[int(0)] = d_608_v51_
            nw90_[int(1)] = d_608_v51_
            nw90_[int(2)] = d_608_v51_
            nw90_[int(3)] = d_608_v51_
            nw90_[int(4)] = d_608_v51_
            nw90_[int(5)] = d_608_v51_
            nw90_[int(6)] = d_608_v51_
            nw90_[int(7)] = d_608_v51_
            nw90_[int(8)] = d_608_v51_
            nw90_[int(9)] = d_608_v51_
            nw90_[int(10)] = d_608_v51_
            nw90_[int(11)] = d_608_v51_
            nw90_[int(12)] = d_608_v51_
            nw90_[int(13)] = d_608_v51_
            nw90_[int(14)] = d_608_v51_
            nw90_[int(15)] = (_dafny.CodePoint('v') if self.f28 else ((d_609_v52_)[default__.fm6(globalState)] if (default__.fm6(globalState)) in (d_609_v52_) else d_608_v51_))
            nw90_[int(16)] = d_608_v51_
            nw90_[int(17)] = d_608_v51_
            nw90_[int(18)] = _dafny.CodePoint('b')
            d_610_v53_ = nw90_
            index79_ = default__.safeIndex(376, (d_610_v53_).length(0))
            (d_610_v53_)[index79_] = d_608_v51_
            index80_ = default__.safeIndex(376, (d_610_v53_).length(0))
            (d_610_v53_)[index80_] = d_608_v51_
            d_611_v54_: _dafny.Seq
            d_611_v54_ = _dafny.SeqWithoutIsStrInference([(0) - (p1)])
            d_612_v55_: _dafny.Seq
            d_612_v55_ = _dafny.SeqWithoutIsStrInference([(d_611_v54_)[default__.safeIndex(p0, len(d_611_v54_))]])
            d_613_v56_: _dafny.Map
            d_613_v56_ = _dafny.Map({(self).f27: p2})
            d_614_v57_: D1
            d_614_v57_ = D1_DC5(self.f22, d_613_v56_, p1, _dafny.MultiSet(d_611_v54_))
            d_615_v58_: _dafny.Map
            d_615_v58_ = _dafny.Map({d_608_v51_: p2})
            d_616_v59_: _dafny.Seq
            d_616_v59_ = _dafny.SeqWithoutIsStrInference([default__.fm5(p1, ((d_615_v58_)[(d_610_v53_)[default__.safeIndex(376, (d_610_v53_).length(0))]] if ((d_610_v53_)[default__.safeIndex(376, (d_610_v53_).length(0))]) in (d_615_v58_) else p0), p2, globalState), _dafny.SeqWithoutIsStrInference([p1]), d_612_v55_])
            d_617_v60_: _dafny.Array
            nw91_ = _dafny.Array(_dafny.MultiSet({}), 17)
            d_617_v60_ = nw91_
            d_618_v61_: C1
            nw92_ = C1()
            nw92_.ctor__(d_617_v60_, self.f28)
            d_618_v61_ = nw92_
            d_619_v62_: _dafny.Map
            d_619_v62_ = _dafny.Map({(d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]: d_618_v61_})
            d_620_v63_: _dafny.MultiSet
            d_620_v63_ = _dafny.MultiSet([(d_618_v61_).f25])
            d_621_v64_: _dafny.MultiSet
            d_621_v64_ = _dafny.MultiSet([d_620_v63_])
            d_622_v65_: _dafny.Map
            d_622_v65_ = _dafny.Map({d_554_v22_: ((d_566_v31_)[p1] if (p1) in (d_566_v31_) else default__.fm12(len(d_619_v62_), d_620_v63_, False, d_621_v64_, globalState))})
            d_623_v66_: _dafny.Array
            nw93_ = _dafny.Array(None, 27)
            nw93_[int(0)] = d_611_v54_
            nw93_[int(1)] = _dafny.SeqWithoutIsStrInference([p2 for d_624_i9_ in range(default__.abs(-664))])
            nw93_[int(2)] = d_612_v55_
            nw93_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_614_v57_).cf9, p2, p0, len(d_611_v54_)])
            nw93_[int(4)] = (d_611_v54_) + ((d_616_v59_)[default__.safeIndex(p1, len(d_616_v59_))])
            nw93_[int(5)] = d_612_v55_
            nw93_[int(6)] = d_611_v54_
            nw93_[int(7)] = default__.fm5(p3, p0, -369, globalState)
            nw93_[int(8)] = (d_611_v54_) + (d_611_v54_)
            nw93_[int(9)] = d_612_v55_
            nw93_[int(10)] = default__.fm5(p3, p3, -379, globalState)
            nw93_[int(11)] = _dafny.SeqWithoutIsStrInference([919])
            nw93_[int(12)] = _dafny.SeqWithoutIsStrInference([p1 for d_625_i10_ in range(default__.abs(588))])
            nw93_[int(13)] = d_612_v55_
            nw93_[int(14)] = d_611_v54_
            nw93_[int(15)] = d_611_v54_
            nw93_[int(16)] = d_612_v55_
            nw93_[int(17)] = (((_dafny.SeqWithoutIsStrInference([p2 for d_626_i11_ in range(default__.abs(766))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p2 for d_627_i11_ in range(default__.abs(766))]))), len(d_622_v65_))).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([p2 for d_628_i11_ in range(default__.abs(766))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p2 for d_629_i11_ in range(default__.abs(766))]))), len(d_622_v65_)))), p1)).set(default__.safeIndex(len(self.f21), len(((_dafny.SeqWithoutIsStrInference([p2 for d_630_i11_ in range(default__.abs(766))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p2 for d_631_i11_ in range(default__.abs(766))]))), len(d_622_v65_))).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([p2 for d_632_i11_ in range(default__.abs(766))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p2 for d_633_i11_ in range(default__.abs(766))]))), len(d_622_v65_)))), p1))), p0)
            nw93_[int(18)] = d_611_v54_
            nw93_[int(19)] = d_611_v54_
            nw93_[int(20)] = d_612_v55_
            nw93_[int(21)] = d_611_v54_
            nw93_[int(22)] = _dafny.SeqWithoutIsStrInference([p1, len(_dafny.Map({(self).f27: p1})), p3, p2])
            nw93_[int(23)] = _dafny.SeqWithoutIsStrInference([len(d_613_v56_), p1])
            nw93_[int(24)] = (d_611_v54_) + (d_612_v55_)
            nw93_[int(25)] = d_611_v54_
            nw93_[int(26)] = (d_611_v54_) + (d_612_v55_)
            d_623_v66_ = nw93_
            index81_ = default__.safeIndex(543, (d_623_v66_).length(0))
            (d_623_v66_)[index81_] = _dafny.SeqWithoutIsStrInference([p0])
            index82_ = default__.safeIndex(543, (d_623_v66_).length(0))
            (d_623_v66_)[index82_] = (d_612_v55_) + (d_612_v55_)
            d_634_v67_: D3
            d_634_v67_ = D3_DC9((p3) + (p1), p0, p0)
            source12_ = d_634_v67_
            if source12_.is_DC8:
                d_635___mcc_h9_ = source12_.cf13
                d_636___mcc_h10_ = source12_.cf14
                d_637___mcc_h11_ = source12_.cf15
                d_638___mcc_h12_ = source12_.cf16
                d_639_cf16_ = d_638___mcc_h12_
                d_640_cf15_ = d_637___mcc_h11_
                d_641_cf14_ = d_636___mcc_h10_
                d_642_cf13_ = d_635___mcc_h9_
                (globalState).f3 = (p3) * (default__.fm6(globalState))
                d_643_v68_: _dafny.Map
                d_643_v68_ = _dafny.Map({p1: d_642_cf13_})
                d_644_v69_: _dafny.Set
                d_644_v69_ = _dafny.Set({self.f21})
                d_645_v70_: C0
                nw94_ = C0()
                nw94_.ctor__(d_644_v69_)
                d_645_v70_ = nw94_
                d_646_v71_: _dafny.Seq
                d_646_v71_ = _dafny.SeqWithoutIsStrInference([d_645_v70_, d_645_v70_])
                d_647_v72_: _dafny.Seq
                d_647_v72_ = _dafny.SeqWithoutIsStrInference([d_646_v71_, d_646_v71_, _dafny.SeqWithoutIsStrInference([d_645_v70_, d_645_v70_])])
                d_648_v73_: _dafny.MultiSet
                d_648_v73_ = _dafny.MultiSet([(_dafny.MultiSet(d_647_v72_)).cardinality, d_642_cf13_, d_642_cf13_])
                pat_let_tv32_ = d_613_v56_
                pat_let_tv33_ = d_613_v56_
                pat_let_tv34_ = d_613_v56_
                pat_let_tv35_ = p1
                pat_let_tv36_ = d_648_v73_
                pat_let_tv37_ = p2
                def iife37_(_pat_let10_0):
                    def iife38_(d_649_dt__update__tmp_h1_):
                        def iife39_(_pat_let11_0):
                            def iife40_(d_650_dt__update_hcf8_h0_):
                                def iife41_(_pat_let12_0):
                                    def iife42_(d_651_dt__update_hcf10_h0_):
                                        def iife43_(_pat_let13_0):
                                            def iife44_(d_652_dt__update_hcf9_h0_):
                                                return D1_DC5((d_649_dt__update__tmp_h1_).cf7, d_650_dt__update_hcf8_h0_, d_652_dt__update_hcf9_h0_, d_651_dt__update_hcf10_h0_)
                                            return iife44_(_pat_let13_0)
                                        return iife43_(pat_let_tv37_)
                                    return iife42_(_pat_let12_0)
                                return iife41_(pat_let_tv36_)
                            return iife40_(_pat_let11_0)
                        return iife39_((pat_let_tv32_).set(self.f28, ((pat_let_tv33_)[False] if (False) in (pat_let_tv34_) else pat_let_tv35_)))
                    return iife38_(_pat_let10_0)
                rhs84_ = ((d_643_v68_)[(len(self.f21)) * (p1)] if ((len(self.f21)) * (p1)) in (d_643_v68_) else (0) - (p3))
                rhs85_ = self.f22
                rhs86_ = iife37_(d_640_cf15_)
                rhs87_ = (d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]
                rhs88_ = self.f22
                lhs69_ = globalState
                lhs70_ = self
                lhs69_.f20 = rhs84_
                d_639_cf16_ = rhs85_
                d_614_v57_ = rhs86_
                d_639_cf16_ = rhs87_
                lhs70_.f22 = rhs88_
                d_653_v74_: _dafny.Map
                d_653_v74_ = _dafny.Map({not(d_639_cf16_): (len(d_643_v68_)) != (p1)})
                rhs89_ = d_653_v74_
                rhs90_ = d_645_v70_
                rhs91_ = d_639_cf16_
                rhs92_ = (d_653_v74_) | (d_653_v74_)
                lhs71_ = self
                d_653_v74_ = rhs89_
                d_645_v70_ = rhs90_
                lhs71_.f28 = rhs91_
                d_653_v74_ = rhs92_
                rhs93_ = ((d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]) and (((d_618_v61_).f25 if self.f28 else self.f28))
                rhs94_ = (self).f27
                lhs72_ = self
                lhs73_ = self
                lhs72_.f28 = rhs93_
                lhs73_.f22 = rhs94_
            elif source12_.is_DC9:
                d_654___mcc_h13_ = source12_.cf17
                d_655___mcc_h14_ = source12_.cf18
                d_656___mcc_h15_ = source12_.cf19
                d_657_cf19_ = d_656___mcc_h15_
                d_658_cf18_ = d_655___mcc_h14_
                d_659_cf17_ = d_654___mcc_h13_
                d_660_v75_: _dafny.Seq
                d_660_v75_ = _dafny.SeqWithoutIsStrInference([d_617_v60_])
                d_661_v76_: C1
                nw95_ = C1()
                nw95_.ctor__((d_660_v75_)[default__.safeIndex(d_658_cf18_, len(d_660_v75_))], ((d_553_v21_)[default__.safeIndex(690, (d_553_v21_).length(0))]) or ((self).f27))
                d_661_v76_ = nw95_
                d_662_v77_: bool
                out5_: bool
                out5_ = default__.m0(d_553_v21_, self.f22, globalState)
                d_662_v77_ = out5_
                d_663_v78_: bool
                out6_: bool
                out6_ = default__.m0(d_553_v21_, (d_618_v61_).f25, globalState)
                d_663_v78_ = out6_
                (self).f22 = d_662_v77_
            elif True:
                d_664___mcc_h16_ = source12_.cf12
                d_665_cf12_ = d_664___mcc_h16_
                (globalState).f20 = p3
                d_666_v79_: _dafny.Map
                d_666_v79_ = _dafny.Map({(self).f27: (True if self.f22 else (d_618_v61_).f25)})
                d_666_v79_ = (d_666_v79_).set((self).f27, (d_618_v61_).f25)
                d_667_v80_: _dafny.Array
                nw96_ = _dafny.Array(None, 11)
                nw96_[int(0)] = self.f21
                nw96_[int(1)] = (default__.fm11(self.f22, globalState)).set(default__.safeIndex(668, len(default__.fm11(self.f22, globalState))), d_608_v51_)
                nw96_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnsquovwx"))
                nw96_[int(3)] = (self.f21) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ittgfr")))
                nw96_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqonwvi"))) + (self.f21)
                nw96_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwfxaqtqq"))
                nw96_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mt"))
                nw96_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_668_i12_ in range(default__.abs(242))])
                nw96_[int(8)] = (_dafny.SeqWithoutIsStrInference([d_608_v51_ for d_669_i13_ in range(default__.abs(-696))])) + (self.f21)
                nw96_[int(9)] = self.f21
                nw96_[int(10)] = self.f21
                d_667_v80_ = nw96_
                index83_ = default__.safeIndex(733, (d_667_v80_).length(0))
                (d_667_v80_)[index83_] = (_dafny.SeqWithoutIsStrInference([d_608_v51_ for d_670_i14_ in range(default__.abs(-122))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmvw")))
                index84_ = default__.safeIndex(733, (d_667_v80_).length(0))
                (d_667_v80_)[index84_] = self.f21
                d_612_v55_ = _dafny.SeqWithoutIsStrInference([222 for d_671_i15_ in range(default__.abs(301))])
            index85_ = default__.safeIndex(690, (d_553_v21_).length(0))
            (d_553_v21_)[index85_] = True
            d_672_v81_: _dafny.Seq
            d_672_v81_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_620_v63_])])
            d_619_v62_ = (d_619_v62_).set(default__.fm12(p2, _dafny.MultiSet([(self).f27, self.f22, (d_618_v61_).f25, default__.fm12(p1, d_620_v63_, (self).f27, _dafny.MultiSet([d_620_v63_]), globalState), default__.fm12(len(self.f21), _dafny.MultiSet([self.f22, self.f28]), self.f22, d_621_v64_, globalState)]), self.f22, (d_672_v81_)[default__.safeIndex(-979, len(d_672_v81_))], globalState), d_618_v61_)
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_552_v20_).length(0)):
            d_673_i16_: int = guard_loop_0_
            if (True) and (((0) <= (d_673_i16_)) and ((d_673_i16_) < ((d_552_v20_).length(0)))):
                (d_552_v20_)[(d_673_i16_)] = (d_673_i16_) * (p3)
        d_674_v82_: _dafny.Array
        nw97_ = _dafny.Array(_dafny.Array(None, 0), 24)
        d_674_v82_ = nw97_
        r0 = d_674_v82_
        return r0

    @property
    def f27(self):
        return self._f27

class C4(T1, T0):
    def  __init__(self):
        self._f21: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f22: bool = False
        self._f23: _dafny.Seq = _dafny.Seq({})
        self.f29: _dafny.Set = _dafny.Set({})
        self._f30: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f21(self):
        return self._f21
    @f21.setter
    def f21(self, value):
        self._f21 = value
    @property
    def f22(self):
        return self._f22
    @f22.setter
    def f22(self, value):
        self._f22 = value
    @property
    def f23(self):
        return self._f23
    @f23.setter
    def f23(self, value):
        self._f23 = value
    def ctor__(self, f29, f30, f23, f21, f22):
        (self).f29 = f29
        (self)._f30 = f30
        (self).f23 = f23
        (self).f21 = f21
        (self).f22 = f22

    def fm2(self, p0, globalState):
        if (self).f30:
            return False
        elif True:
            return self.f22

    def fm0(self, globalState):
        return (_dafny.Set({False})).intersection((_dafny.Set({self.f22, (self).f30})).intersection(_dafny.Set({(self).f30})))

    def fm1(self, p0, p1, p2, globalState):
        return self.f22

    @property
    def f30(self):
        return self._f30
