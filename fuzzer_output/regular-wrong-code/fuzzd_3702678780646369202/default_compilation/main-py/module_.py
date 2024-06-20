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
        return not((567) >= ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')]))) * (len((D6_DC14(_dafny.SeqWithoutIsStrInference([True, False]))).cf20))))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.SeqWithoutIsStrInference([151])).Elements:
                    d_0_v1_: int = compr_2_
                    if (d_0_v1_) in (_dafny.SeqWithoutIsStrInference([151])):
                        coll2_[(d_0_v1_) - (829)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i0_ in range(default__.abs(80))]))})), 523]))
                return _dafny.Map(coll2_)
            coll0_ = _dafny.Map()
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([151])).Elements:
                    d_0_v1_: int = compr_1_
                    if (d_0_v1_) in (_dafny.SeqWithoutIsStrInference([151])):
                        coll1_[(d_0_v1_) - (829)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i0_ in range(default__.abs(80))]))})), 523]))
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_2_v0_: int = compr_0_
                if (d_2_v0_) in (iife2_()
                ):
                    coll0_[(d_2_v0_) - (542)] = 921
            return _dafny.Map(coll0_)
        return (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([iife0_()
        , _dafny.Map({627: 464}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 638})])): 126}), _dafny.Map({-220: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: not(False)})) for d_3_i1_ in range(default__.abs(487))]))})])), -511)) + (535)

    @staticmethod
    def fm4(globalState):
        return ((_dafny.MultiSet([462])) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnvmueh"))), len(_dafny.Map({614: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_4_i0_ in range(default__.abs(822))]))}))])), (0) - (len(_dafny.Map({True: 239}))), -650, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality]))) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([-585])) + (_dafny.SeqWithoutIsStrInference([913, len(_dafny.Set({_dafny.Map({_dafny.CodePoint('e'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))}), _dafny.Map({_dafny.CodePoint('j'): len(_dafny.Map({301: False}))})}))]))))

    @staticmethod
    def fm5(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(963, 663):
                d_5_v0_: int = compr_3_
                if ((963) <= (d_5_v0_)) and ((d_5_v0_) < (663)):
                    coll3_[default__.safeModuloInt(d_5_v0_, 441)] = False
            return _dafny.Map(coll3_)
        return ((_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([True]))): False})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-620: (_dafny.MultiSet([False])).cardinality}))])): True}))) | (iife3_()
        )

    @staticmethod
    def fm6(p0, p1, globalState):
        return ((_dafny.MultiSet([False, False, True]) if True else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, False, not(False), False])))) | (_dafny.MultiSet([not(True)]))

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        return _dafny.Map({True: False})

    @staticmethod
    def fm8(p0, globalState):
        return _dafny.Set({(90) >= ((0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjl"))): -53}))))})

    @staticmethod
    def fm9(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usw"))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(335, 539):
                d_6_v0_: int = compr_4_
                if ((335) <= (d_6_v0_)) and ((d_6_v0_) < (539)):
                    coll4_[(d_6_v0_) * ((0) - (len(_dafny.Set({False}))))] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_7_i0_ in range(default__.abs(129))])))
            return _dafny.Map(coll4_)
        return _dafny.SeqWithoutIsStrInference([(iife4_()
        ) != (_dafny.Map({186: len(_dafny.Map({True: True}))})), (True) not in (_dafny.SeqWithoutIsStrInference([True, True, not(True), True, True]))])

    @staticmethod
    def fm11(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqmuuiom")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofdbvnl"))]))) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bc"))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_8_i1_ in range(default__.abs(42))]) for d_9_i0_ in range(default__.abs(379))]))))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.Map({False: _dafny.CodePoint('v')})])) - (_dafny.MultiSet([_dafny.Map({True: _dafny.CodePoint('c')}), _dafny.Map({False: _dafny.CodePoint('s')}), _dafny.Map({True: _dafny.CodePoint('s')})]))

    @staticmethod
    def m0(globalState):
        d_10_v0_: bool
        d_10_v0_ = False
        d_11_v1_: int
        d_11_v1_ = -798
        d_12_v2_: _dafny.Map
        d_12_v2_ = _dafny.Map({d_10_v0_: d_10_v0_})
        d_13_v3_: _dafny.Seq
        d_13_v3_ = _dafny.SeqWithoutIsStrInference([not(d_10_v0_)])
        d_14_v4_: _dafny.Set
        d_14_v4_ = _dafny.Set({d_11_v1_, (0) - (d_11_v1_)})
        d_15_v5_: _dafny.Set
        d_15_v5_ = _dafny.Set({d_10_v0_})
        d_16_v6_: _dafny.Array
        nw0_ = _dafny.Array(None, 22)
        nw0_[int(0)] = d_11_v1_
        nw0_[int(1)] = len(d_12_v2_)
        nw0_[int(2)] = (0) - (d_11_v1_)
        nw0_[int(3)] = len(d_13_v3_)
        nw0_[int(4)] = len(d_14_v4_)
        nw0_[int(5)] = d_11_v1_
        nw0_[int(6)] = d_11_v1_
        nw0_[int(7)] = d_11_v1_
        nw0_[int(8)] = d_11_v1_
        nw0_[int(9)] = d_11_v1_
        nw0_[int(10)] = d_11_v1_
        nw0_[int(11)] = d_11_v1_
        nw0_[int(12)] = d_11_v1_
        nw0_[int(13)] = d_11_v1_
        nw0_[int(14)] = d_11_v1_
        nw0_[int(15)] = len(d_15_v5_)
        nw0_[int(16)] = len(d_13_v3_)
        nw0_[int(17)] = -292
        nw0_[int(18)] = d_11_v1_
        nw0_[int(19)] = d_11_v1_
        nw0_[int(20)] = d_11_v1_
        nw0_[int(21)] = d_11_v1_
        d_16_v6_ = nw0_
        d_17_v7_: C0
        nw1_ = C0()
        nw1_.ctor__(d_10_v0_, (d_16_v6_ if d_10_v0_ else d_16_v6_))
        d_17_v7_ = nw1_
        d_18_i0_: int
        d_18_i0_ = 0
        with _dafny.label("0"):
            while (d_17_v7_).f28:
                with _dafny.c_label("0"):
                    if (d_18_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_18_i0_ = (d_18_i0_) + (1)
                    if default__.fm0(globalState):
                        d_19_v8_: str
                        d_19_v8_ = _dafny.CodePoint('d')
                        d_19_v8_ = (d_19_v8_ if d_10_v0_ else d_19_v8_)
                        d_20_v9_: _dafny.Array
                        nw2_ = _dafny.Array(None, 19)
                        nw2_[int(0)] = d_10_v0_
                        nw2_[int(1)] = (d_17_v7_).f28
                        nw2_[int(2)] = d_10_v0_
                        nw2_[int(3)] = (d_17_v7_).f28
                        nw2_[int(4)] = (d_17_v7_).f28
                        nw2_[int(5)] = (d_17_v7_).fm2((d_17_v7_).f28, globalState)
                        nw2_[int(6)] = (d_17_v7_).f28
                        nw2_[int(7)] = (d_17_v7_).f28
                        nw2_[int(8)] = not((d_17_v7_).f28)
                        nw2_[int(9)] = d_10_v0_
                        nw2_[int(10)] = (d_17_v7_).f28
                        nw2_[int(11)] = d_10_v0_
                        nw2_[int(12)] = (d_17_v7_).f28
                        nw2_[int(13)] = (d_17_v7_).f28
                        nw2_[int(14)] = (d_17_v7_).f28
                        nw2_[int(15)] = (d_17_v7_).f28
                        nw2_[int(16)] = not((d_17_v7_).f28)
                        nw2_[int(17)] = d_10_v0_
                        nw2_[int(18)] = (d_17_v7_).f28
                        d_20_v9_ = nw2_
                        d_21_v10_: _dafny.Array
                        d_21_v10_ = d_20_v9_
                        d_22_v11_: _dafny.Seq
                        d_22_v11_ = _dafny.SeqWithoutIsStrInference([d_20_v9_, d_20_v9_, d_20_v9_, d_20_v9_, d_20_v9_])
                        d_23_v12_: _dafny.Seq
                        d_23_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btlenkt"))
                        d_24_v13_: _dafny.Set
                        d_24_v13_ = _dafny.Set({d_20_v9_, (d_21_v10_), d_20_v9_, (d_22_v11_)[default__.safeIndex((0) - (len(d_23_v12_)), len(d_22_v11_))]})
                        d_25_v14_: D0
                        d_25_v14_ = D0_DC1((d_17_v7_).f28, d_11_v1_)
                        d_26_v15_: D0
                        d_26_v15_ = D0_DC2(d_25_v14_)
                        d_27_v16_: _dafny.MultiSet
                        d_27_v16_ = _dafny.MultiSet([d_10_v0_, not((d_17_v7_).f28)])
                        d_28_v17_: _dafny.Map
                        d_28_v17_ = _dafny.Map({d_11_v1_: d_20_v9_})
                        d_29_v18_: _dafny.Seq
                        d_29_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({((d_28_v17_)[len(_dafny.Set({d_11_v1_, d_11_v1_, len(d_15_v5_), d_11_v1_}))] if (len(_dafny.Set({d_11_v1_, d_11_v1_, len(d_15_v5_), d_11_v1_}))) in (d_28_v17_) else d_20_v9_)}), d_24_v13_, (d_24_v13_) - (d_24_v13_), _dafny.Set({d_20_v9_}), (d_24_v13_) | (d_24_v13_)])
                        d_30_v19_: D1
                        d_30_v19_ = D1_DC4((d_17_v7_).f28)
                        rhs0_ = (d_29_v18_)[default__.safeIndex(d_11_v1_, len(d_29_v18_))]
                        rhs1_ = d_26_v15_
                        rhs2_ = d_20_v9_
                        rhs3_ = ((d_27_v16_) | (d_27_v16_)).intersection(default__.fm6((d_30_v19_).cf5, d_11_v1_, globalState))
                        d_24_v13_ = rhs0_
                        d_26_v15_ = rhs1_
                        d_20_v9_ = rhs2_
                        d_27_v16_ = rhs3_
                        d_31_v20_: C0
                        nw3_ = C0()
                        nw3_.ctor__((d_17_v7_).f28, (d_17_v7_).f29)
                        d_31_v20_ = nw3_
                        d_32_v21_: _dafny.Map
                        d_32_v21_ = _dafny.Map({(d_17_v7_).f28: d_13_v3_})
                        d_33_v22_: _dafny.Seq
                        d_33_v22_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_17_v7_).f28]), d_13_v3_, (((d_32_v21_)[d_10_v0_] if (d_10_v0_) in (d_32_v21_) else d_13_v3_)) + (d_13_v3_), _dafny.SeqWithoutIsStrInference([(d_31_v20_).f28])])
                        d_34_v23_: _dafny.Map
                        d_34_v23_ = _dafny.Map({d_11_v1_: d_11_v1_})
                        d_35_v24_: D3
                        d_35_v24_ = D3_DC10((d_17_v7_).fm2((d_17_v7_).f28, globalState), len(d_34_v23_), not((d_17_v7_).f28), 282, (d_17_v7_).f29)
                        index0_ = default__.safeIndex(970, (d_20_v9_).length(0))
                        (d_20_v9_)[index0_] = (d_35_v24_).cf14
                        index1_ = default__.safeIndex(970, (d_20_v9_).length(0))
                        rhs4_ = False
                        rhs5_ = (d_11_v1_) > ((d_11_v1_) * (len(_dafny.Map({d_11_v1_: -924}))))
                        rhs6_ = d_33_v22_
                        rhs7_ = not((d_17_v7_).f28)
                        lhs0_ = globalState
                        lhs1_ = globalState
                        lhs2_ = d_20_v9_
                        lhs3_ = default__.safeIndex(970, (d_20_v9_).length(0))
                        lhs0_.f0 = rhs4_
                        lhs1_.f0 = rhs5_
                        d_33_v22_ = rhs6_
                        lhs2_[lhs3_] = rhs7_
                        d_36_v25_: D2
                        d_36_v25_ = D2_DC6(d_27_v16_)
                        d_37_v26_: _dafny.Map
                        d_37_v26_ = _dafny.Map({(d_36_v25_ if (d_31_v20_).fm2((d_17_v7_).f28, globalState) else d_36_v25_): 841})
                        d_37_v26_ = (d_37_v26_).set(D2_DC6(d_27_v16_), d_11_v1_)
                    elif True:
                        d_38_v27_: _dafny.MultiSet
                        d_38_v27_ = _dafny.MultiSet([(d_11_v1_) <= (d_11_v1_), (d_11_v1_) == (d_11_v1_)])
                        rhs8_ = d_11_v1_
                        rhs9_ = ((d_38_v27_)[(d_17_v7_).f28] if ((d_17_v7_).f28) in (d_38_v27_) else 591)
                        rhs10_ = (d_17_v7_).f29
                        d_11_v1_ = rhs8_
                        d_11_v1_ = rhs9_
                        d_16_v6_ = rhs10_
                        d_39_v28_: _dafny.Array
                        nw4_ = _dafny.Array(False, 25)
                        d_39_v28_ = nw4_
                        index2_ = default__.safeIndex(941, (d_39_v28_).length(0))
                        (d_39_v28_)[index2_] = (d_17_v7_).f28
                        d_40_v29_: str
                        d_40_v29_ = _dafny.CodePoint('o')
                        d_41_v30_: _dafny.Seq
                        d_41_v30_ = _dafny.SeqWithoutIsStrInference([d_40_v29_, d_40_v29_, d_40_v29_])
                        d_42_v31_: D0
                        d_42_v31_ = D0_DC1(d_10_v0_, -832)
                        index3_ = default__.safeIndex(941, (d_39_v28_).length(0))
                        def iife5_():
                            coll5_ = _dafny.Set()
                            compr_5_: int
                            for compr_5_ in (_dafny.SeqWithoutIsStrInference([d_11_v1_ for d_43_i1_ in range(default__.abs(944))])).Elements:
                                d_44_v32_: int = compr_5_
                                if (d_44_v32_) in (_dafny.SeqWithoutIsStrInference([d_11_v1_ for d_43_i1_ in range(default__.abs(944))])):
                                    coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_44_v32_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbrnrtwv"))))]))
                            return _dafny.Set(coll5_)
                        rhs11_ = (len(d_41_v30_)) + ((d_42_v31_).cf2)
                        rhs12_ = (d_17_v7_).f28
                        rhs13_ = ((d_11_v1_) in (iife5_()
                        ) if (d_17_v7_).f28 else (d_11_v1_) < (d_11_v1_))
                        lhs4_ = d_39_v28_
                        lhs5_ = default__.safeIndex(941, (d_39_v28_).length(0))
                        d_11_v1_ = rhs11_
                        lhs4_[lhs5_] = rhs12_
                        d_10_v0_ = rhs13_
                        d_40_v29_ = d_40_v29_
                        d_45_v33_: _dafny.MultiSet
                        d_45_v33_ = _dafny.MultiSet([d_17_v7_, d_17_v7_, d_17_v7_])
                        d_11_v1_ = ((d_45_v33_)[d_17_v7_] if (d_17_v7_) in (d_45_v33_) else (d_11_v1_) + (d_11_v1_))
                        d_46_v34_: _dafny.Map
                        d_46_v34_ = _dafny.Map({d_41_v30_: d_10_v0_})
                        d_47_v35_: _dafny.Array
                        nw5_ = _dafny.Array(None, 7)
                        nw5_[int(0)] = (d_12_v2_).set(d_10_v0_, False)
                        nw5_[int(1)] = d_12_v2_
                        nw5_[int(2)] = d_12_v2_
                        nw5_[int(3)] = (d_12_v2_) | (default__.fm7((d_17_v7_).f28, (d_17_v7_).f28, (d_39_v28_)[default__.safeIndex(941, (d_39_v28_).length(0))], d_11_v1_, globalState))
                        nw5_[int(4)] = d_12_v2_
                        nw5_[int(5)] = (default__.fm7(d_10_v0_, default__.fm0(globalState), (d_17_v7_).f28, len(d_46_v34_), globalState)).set(d_10_v0_, d_10_v0_)
                        nw5_[int(6)] = d_12_v2_
                        d_47_v35_ = nw5_
                        index4_ = default__.safeIndex(382, (d_47_v35_).length(0))
                        (d_47_v35_)[index4_] = (_dafny.Map({(d_17_v7_).f28: (d_17_v7_).fm2((d_17_v7_).f28, globalState)})) | (_dafny.Map({(d_39_v28_)[default__.safeIndex(941, (d_39_v28_).length(0))]: (d_17_v7_).f28}))
                        index5_ = default__.safeIndex(382, (d_47_v35_).length(0))
                        (d_47_v35_)[index5_] = d_12_v2_
                    if not (d_10_v0_) or (not ((d_17_v7_).f28) or (not((d_17_v7_).f28))):
                        d_48_v36_: _dafny.Seq
                        d_48_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vithblii"))
                        (globalState).f0 = (_dafny.CodePoint('i')) in ((d_48_v36_) + (d_48_v36_))
                        d_49_v37_: _dafny.Array
                        def lambda0_(d_50_i2_):
                            return _dafny.CodePoint('y')

                        init0_ = lambda0_
                        nw6_ = _dafny.Array(None, 7)
                        for i0_0_ in range(nw6_.length(0)):
                            nw6_[i0_0_] = init0_(i0_0_)
                        d_49_v37_ = nw6_
                        d_51_v38_: _dafny.Seq
                        d_51_v38_ = d_48_v36_
                        rhs14_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_52_i3_ in range(default__.abs(680))])) <= (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_53_i4_ in range(default__.abs(870))]))) + (d_48_v36_))
                        rhs15_ = (d_49_v37_ if (d_17_v7_).f28 else (d_49_v37_ if (d_17_v7_).f28 else d_49_v37_))
                        lhs6_ = globalState
                        lhs6_.f0 = rhs14_
                        d_49_v37_ = rhs15_
                        d_10_v0_ = d_10_v0_
                        d_11_v1_ = (default__.safeDivisionInt(d_11_v1_, d_11_v1_)) - (default__.fm3(d_11_v1_, (d_17_v7_).f28, d_11_v1_, _dafny.Map({d_11_v1_: (d_17_v7_).f28}), globalState))
                        d_54_v39_: _dafny.Array
                        nw7_ = _dafny.Array(_dafny.Set({}), 23)
                        d_54_v39_ = nw7_
                        d_55_v40_: _dafny.Map
                        d_55_v40_ = _dafny.Map({d_11_v1_: d_17_v7_})
                        d_56_v41_: _dafny.Map
                        d_56_v41_ = _dafny.Map({len(d_55_v40_): default__.fm8(d_11_v1_, globalState)})
                        index6_ = default__.safeIndex(908, (d_54_v39_).length(0))
                        (d_54_v39_)[index6_] = (_dafny.Set({(d_17_v7_).f28, d_10_v0_, d_10_v0_, (d_17_v7_).f28, (d_17_v7_).f28})) - (((d_56_v41_)[330] if (330) in (d_56_v41_) else default__.fm8(d_11_v1_, globalState)))
                        index7_ = default__.safeIndex(908, (d_54_v39_).length(0))
                        (d_54_v39_)[index7_] = (d_15_v5_) - ((d_15_v5_) | (d_15_v5_))
                    elif True:
                        d_57_v42_: str
                        d_57_v42_ = _dafny.CodePoint('l')
                        d_58_v43_: D3
                        d_58_v43_ = D3_DC10(d_10_v0_, d_11_v1_, (d_17_v7_).f28, d_11_v1_, (d_17_v7_).f29)
                        d_59_v44_: _dafny.Map
                        d_59_v44_ = _dafny.Map({(d_58_v43_).cf14: d_17_v7_})
                        d_60_v45_: _dafny.Seq
                        d_60_v45_ = _dafny.SeqWithoutIsStrInference([((d_59_v44_)[not((d_17_v7_).fm2((d_17_v7_).f28, globalState))] if (not((d_17_v7_).fm2((d_17_v7_).f28, globalState))) in (d_59_v44_) else d_17_v7_), d_17_v7_, d_17_v7_])
                        d_61_v46_: _dafny.Seq
                        d_61_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "squrxyb"))
                        d_62_v47_: _dafny.Seq
                        d_62_v47_ = d_61_v46_
                        d_63_v48_: _dafny.Map
                        d_63_v48_ = _dafny.Map({d_62_v47_: False})
                        d_64_v49_: D6
                        d_64_v49_ = D6_DC13(d_57_v42_)
                        pat_let_tv0_ = d_57_v42_
                        d_65_v50_: D6
                        def iife6_(_pat_let0_0):
                            def iife7_(d_66_dt__update__tmp_h1_):
                                def iife8_(_pat_let1_0):
                                    def iife9_(d_67_dt__update_hcf19_h0_):
                                        return D6_DC13(d_67_dt__update_hcf19_h0_)
                                    return iife9_(_pat_let1_0)
                                return iife8_(pat_let_tv0_)
                            return iife7_(_pat_let0_0)
                        d_65_v50_ = D6_DC13((iife6_(d_64_v49_)).cf19)
                        rhs16_ = d_11_v1_
                        rhs17_ = (d_60_v45_)[default__.safeIndex(len((d_63_v48_) | (d_63_v48_)), len(d_60_v45_))]
                        rhs18_ = (d_64_v49_).cf19
                        d_11_v1_ = rhs16_
                        d_17_v7_ = rhs17_
                        d_57_v42_ = rhs18_
                        d_68_v51_: _dafny.Array
                        nw8_ = _dafny.Array(None, 5)
                        nw8_[int(0)] = (d_17_v7_).f28
                        nw8_[int(1)] = (d_17_v7_).f28
                        nw8_[int(2)] = d_10_v0_
                        nw8_[int(3)] = (d_17_v7_).f28
                        nw8_[int(4)] = d_10_v0_
                        d_68_v51_ = nw8_
                        d_69_v52_: _dafny.Array
                        nw9_ = _dafny.Array(None, 18)
                        nw9_[int(0)] = (d_17_v7_).f29
                        nw9_[int(1)] = d_16_v6_
                        nw9_[int(2)] = (d_17_v7_).f29
                        nw9_[int(3)] = (d_17_v7_).f29
                        nw9_[int(4)] = (d_17_v7_).f29
                        nw9_[int(5)] = (d_17_v7_).f29
                        nw9_[int(6)] = (d_17_v7_).f29
                        nw9_[int(7)] = (d_17_v7_).f29
                        nw9_[int(8)] = (d_17_v7_).f29
                        nw9_[int(9)] = (d_17_v7_).f29
                        nw9_[int(10)] = (d_17_v7_).f29
                        nw9_[int(11)] = (d_17_v7_).f29
                        nw9_[int(12)] = (d_17_v7_).f29
                        nw9_[int(13)] = (d_17_v7_).f29
                        nw9_[int(14)] = d_16_v6_
                        nw9_[int(15)] = (d_17_v7_).f29
                        nw9_[int(16)] = (d_17_v7_).f29
                        nw9_[int(17)] = d_16_v6_
                        d_69_v52_ = nw9_
                        d_70_v53_: _dafny.Map
                        d_70_v53_ = _dafny.Map({d_68_v51_: d_69_v52_})
                        d_70_v53_ = (d_70_v53_).set(d_68_v51_, d_69_v52_)
                        d_71_v54_: C0
                        nw10_ = C0()
                        nw10_.ctor__(d_10_v0_, (d_17_v7_).f29)
                        d_71_v54_ = nw10_
                        d_72_v55_: _dafny.MultiSet
                        d_72_v55_ = _dafny.MultiSet([d_11_v1_])
                        d_73_v56_: _dafny.Map
                        d_73_v56_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_57_v42_ for d_74_i5_ in range(default__.abs(-735))]): (d_72_v55_).cardinality})
                        d_73_v56_ = (d_73_v56_).set(default__.fm9(globalState), d_11_v1_)
                        d_75_v57_: _dafny.Map
                        d_75_v57_ = _dafny.Map({275: len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))) + (d_61_v46_))})
                        d_76_v58_: _dafny.Map
                        d_76_v58_ = _dafny.Map({d_11_v1_: (d_17_v7_).f28})
                        d_77_v59_: _dafny.Seq
                        d_77_v59_ = _dafny.SeqWithoutIsStrInference([d_76_v58_])
                        d_75_v57_ = (d_75_v57_).set((d_11_v1_ if (d_13_v3_)[default__.safeIndex(765, len(d_13_v3_))] else default__.fm3(d_11_v1_, (d_17_v7_).f28, d_11_v1_, (d_77_v59_)[default__.safeIndex(len(d_61_v46_), len(d_77_v59_))], globalState)), d_11_v1_)
                    d_11_v1_ = (d_11_v1_ if (d_17_v7_).fm2(d_10_v0_, globalState) else d_11_v1_)
                    (globalState).f0 = (d_17_v7_).f28
                    pass
            pass
        d_78_v60_: _dafny.MultiSet
        d_78_v60_ = _dafny.MultiSet([d_10_v0_, (d_17_v7_).f28, d_10_v0_, (d_17_v7_).f28])
        d_79_v61_: _dafny.Array
        nw11_ = _dafny.Array(False, 8)
        d_79_v61_ = nw11_
        d_80_v62_: _dafny.Map
        d_80_v62_ = _dafny.Map({618: d_79_v61_})
        d_81_v63_: _dafny.Seq
        d_81_v63_ = _dafny.SeqWithoutIsStrInference([d_80_v62_, d_80_v62_, d_80_v62_])
        rhs19_ = (d_78_v60_).isdisjoint(d_78_v60_)
        rhs20_ = True
        rhs21_ = (d_17_v7_).f28
        rhs22_ = (len(((d_81_v63_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_11_v1_, d_11_v1_])), len(d_81_v63_))]) | (d_80_v62_))) - (d_11_v1_)
        lhs7_ = globalState
        lhs8_ = globalState
        lhs7_.f0 = rhs19_
        d_10_v0_ = rhs20_
        lhs8_.f0 = rhs21_
        d_11_v1_ = rhs22_
        if not((d_17_v7_).f28):
            d_11_v1_ = d_11_v1_
            (globalState).f0 = (d_17_v7_).f28
            d_82_v64_: _dafny.Seq
            d_82_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_10_v0_])])
            d_83_v65_: D0
            d_83_v65_ = D0_DC1(d_10_v0_, len(d_82_v64_))
            if (d_83_v65_).cf1:
                index8_ = default__.safeIndex(769, (d_79_v61_).length(0))
                (d_79_v61_)[index8_] = (d_17_v7_).f28
                index9_ = default__.safeIndex(769, (d_79_v61_).length(0))
                (d_79_v61_)[index9_] = ((d_17_v7_).f28) == ((d_17_v7_).f28)
                d_84_v66_: _dafny.Array
                nw12_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_84_v66_ = nw12_
                d_85_v67_: _dafny.Set
                d_85_v67_ = _dafny.Set({d_84_v66_})
                (globalState).f0 = (_dafny.Set({d_84_v66_})).issubset(d_85_v67_)
                d_86_v68_: D6
                d_86_v68_ = D6_DC14(default__.fm10(d_11_v1_, 673, d_11_v1_, globalState))
                d_87_v69_: D1
                d_87_v69_ = D1_DC3((d_17_v7_).f29)
                d_88_v70_: _dafny.Map
                d_88_v70_ = _dafny.Map({len((d_86_v68_).cf20): d_87_v69_})
                d_89_v71_: D1
                d_89_v71_ = D1_DC5(((d_88_v70_)[d_11_v1_] if (d_11_v1_) in (d_88_v70_) else d_87_v69_))
                d_90_v72_: D6
                d_90_v72_ = D6_DC16(_dafny.SeqWithoutIsStrInference([d_89_v71_]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "am")))
                d_91_v73_: _dafny.Seq
                d_91_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrtiwitvl"))
                pat_let_tv1_ = d_87_v69_
                d_92_v74_: _dafny.Seq
                def iife10_(_pat_let2_0):
                    def iife11_(d_93_dt__update__tmp_h2_):
                        def iife12_(_pat_let3_0):
                            def iife13_(d_94_dt__update_hcf6_h0_):
                                return D1_DC5(d_94_dt__update_hcf6_h0_)
                            return iife13_(_pat_let3_0)
                        return iife12_(pat_let_tv1_)
                    return iife11_(_pat_let2_0)
                d_92_v74_ = _dafny.SeqWithoutIsStrInference([iife10_(d_89_v71_), d_89_v71_])
                d_95_v75_: str
                d_95_v75_ = _dafny.CodePoint('i')
                pat_let_tv2_ = d_91_v73_
                pat_let_tv3_ = d_92_v74_
                pat_let_tv4_ = d_91_v73_
                pat_let_tv5_ = d_12_v2_
                pat_let_tv6_ = d_91_v73_
                pat_let_tv7_ = d_95_v75_
                d_96_v76_: _dafny.Array
                nw13_ = _dafny.Array(None, 10)
                nw13_[int(0)] = d_90_v72_
                def iife15_(_pat_let5_0):
                    def iife16_(d_97_dt__update__tmp_h3_):
                        def iife17_(_pat_let6_0):
                            def iife18_(d_98_dt__update_hcf25_h0_):
                                return D6_DC16((d_97_dt__update__tmp_h3_).cf24, d_98_dt__update_hcf25_h0_)
                            return iife18_(_pat_let6_0)
                        return iife17_(pat_let_tv2_)
                    return iife16_(_pat_let5_0)
                def iife14_(_pat_let4_0):
                    def iife19_(d_99_dt__update__tmp_h4_):
                        def iife20_(_pat_let7_0):
                            def iife21_(d_100_dt__update_hcf24_h0_):
                                return D6_DC16(d_100_dt__update_hcf24_h0_, (d_99_dt__update__tmp_h4_).cf25)
                            return iife21_(_pat_let7_0)
                        return iife20_(pat_let_tv3_)
                    return iife19_(_pat_let4_0)
                nw13_[int(1)] = iife14_(iife15_(d_90_v72_))
                nw13_[int(2)] = d_90_v72_
                nw13_[int(3)] = d_90_v72_
                nw13_[int(4)] = d_90_v72_
                nw13_[int(5)] = d_90_v72_
                def iife22_(_pat_let8_0):
                    def iife23_(d_101_dt__update__tmp_h5_):
                        def iife24_(_pat_let9_0):
                            def iife25_(d_102_dt__update_hcf25_h1_):
                                return D6_DC16((d_101_dt__update__tmp_h5_).cf24, d_102_dt__update_hcf25_h1_)
                            return iife25_(_pat_let9_0)
                        return iife24_((pat_let_tv4_).set(default__.safeIndex(len(pat_let_tv5_), len(pat_let_tv6_)), pat_let_tv7_))
                    return iife23_(_pat_let8_0)
                nw13_[int(6)] = iife22_(d_90_v72_)
                nw13_[int(7)] = d_90_v72_
                nw13_[int(8)] = d_90_v72_
                nw13_[int(9)] = d_90_v72_
                d_96_v76_ = nw13_
                index10_ = default__.safeIndex(929, (d_96_v76_).length(0))
                (d_96_v76_)[index10_] = D6_DC16(d_92_v74_, d_91_v73_)
                index11_ = default__.safeIndex(929, (d_96_v76_).length(0))
                (d_96_v76_)[index11_] = (d_90_v72_ if (d_17_v7_).f28 else d_90_v72_)
                (globalState).f0 = (d_17_v7_).f28
                d_90_v72_ = d_90_v72_
            elif True:
                d_103_v77_: _dafny.Set
                d_103_v77_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsxaj"))})
                (globalState).f0 = (d_103_v77_).issubset(d_103_v77_)
                d_104_v78_: _dafny.MultiSet
                d_104_v78_ = _dafny.MultiSet([777])
                d_105_v79_: _dafny.Map
                d_105_v79_ = _dafny.Map({(d_11_v1_) * (len(d_14_v4_)): d_104_v78_})
                d_104_v78_ = ((d_105_v79_)[len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysstjicth"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_106_i6_ in range(default__.abs(-758))])))] if (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysstjicth"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_107_i6_ in range(default__.abs(-758))])))) in (d_105_v79_) else d_104_v78_)
                d_108_v80_: _dafny.Seq
                d_108_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipw"))
                d_109_v81_: _dafny.Map
                d_109_v81_ = _dafny.Map({d_11_v1_: d_108_v80_})
                rhs23_ = default__.fm10((-680) + (d_11_v1_), d_11_v1_, d_11_v1_, globalState)
                rhs24_ = ((d_108_v80_) + (((d_109_v81_)[d_11_v1_] if (d_11_v1_) in (d_109_v81_) else d_108_v80_))) + (d_108_v80_)
                rhs25_ = default__.safeModuloInt(491, d_11_v1_)
                rhs26_ = (d_17_v7_).f28
                lhs9_ = globalState
                lhs10_ = globalState
                d_13_v3_ = rhs23_
                lhs9_.f2 = rhs24_
                d_11_v1_ = rhs25_
                lhs10_.f0 = rhs26_
                index12_ = default__.safeIndex(738, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index12_] = d_11_v1_
                index13_ = default__.safeIndex(738, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index13_] = d_11_v1_
                d_110_v82_: _dafny.Map
                d_110_v82_ = _dafny.Map({d_17_v7_: (d_17_v7_).f29})
                d_110_v82_ = (d_110_v82_).set(d_17_v7_, (d_17_v7_).f29)
            d_111_v83_: str
            d_111_v83_ = _dafny.CodePoint('p')
            d_111_v83_ = d_111_v83_
            d_112_v84_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.Set({}), 14)
            d_112_v84_ = nw14_
            index14_ = default__.safeIndex(57, (d_112_v84_).length(0))
            def iife26_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(-913, 328):
                    d_113_v85_: int = compr_6_
                    if ((-913) <= (d_113_v85_)) and ((d_113_v85_) < (328)):
                        coll6_ = coll6_.union(_dafny.Set([(d_113_v85_) - ((_dafny.MultiSet([d_11_v1_, d_11_v1_])).cardinality)]))
                return _dafny.Set(coll6_)
            (d_112_v84_)[index14_] = iife26_()
            
            index15_ = default__.safeIndex(57, (d_112_v84_).length(0))
            (d_112_v84_)[index15_] = d_14_v4_
        elif True:
            d_114_v86_: _dafny.Map
            d_114_v86_ = _dafny.Map({d_10_v0_: d_11_v1_})
            d_11_v1_ = (((d_114_v86_)[not(not((d_17_v7_).f28))] if (not(not((d_17_v7_).f28))) in (d_114_v86_) else d_11_v1_)) + (d_11_v1_)
            (globalState).f0 = True
            index16_ = default__.safeIndex(488, (d_79_v61_).length(0))
            (d_79_v61_)[index16_] = not ((d_17_v7_).f28) or ((d_17_v7_).f28)
            d_115_v87_: _dafny.Seq
            d_115_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dicp"))
            d_116_v88_: _dafny.Seq
            d_116_v88_ = _dafny.SeqWithoutIsStrInference([d_11_v1_])
            index17_ = default__.safeIndex(488, (d_79_v61_).length(0))
            rhs27_ = d_13_v3_
            rhs28_ = (len((d_115_v87_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqnhuuat")))) if d_10_v0_ else (0) - ((d_116_v88_)[default__.safeIndex(d_11_v1_, len(d_116_v88_))]))
            rhs29_ = (d_17_v7_).f28
            lhs11_ = d_79_v61_
            lhs12_ = default__.safeIndex(488, (d_79_v61_).length(0))
            d_13_v3_ = rhs27_
            d_11_v1_ = rhs28_
            lhs11_[lhs12_] = rhs29_
            d_117_v89_: _dafny.Array
            nw15_ = _dafny.Array(None, 13)
            nw15_[int(0)] = (d_79_v61_)[default__.safeIndex(488, (d_79_v61_).length(0))]
            nw15_[int(1)] = (d_78_v60_).isdisjoint(_dafny.MultiSet([d_10_v0_, d_10_v0_, (d_17_v7_).f28, d_10_v0_]))
            nw15_[int(2)] = d_10_v0_
            nw15_[int(3)] = False
            nw15_[int(4)] = (d_17_v7_).f28
            nw15_[int(5)] = (d_79_v61_)[default__.safeIndex(488, (d_79_v61_).length(0))]
            nw15_[int(6)] = True
            nw15_[int(7)] = (d_17_v7_).f28
            nw15_[int(8)] = (d_79_v61_)[default__.safeIndex(488, (d_79_v61_).length(0))]
            nw15_[int(9)] = (_dafny.MultiSet([672])) == (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_116_v88_) for d_118_i7_ in range(default__.abs(752))])))
            nw15_[int(10)] = (d_79_v61_)[default__.safeIndex(488, (d_79_v61_).length(0))]
            nw15_[int(11)] = (d_17_v7_).f28
            nw15_[int(12)] = (d_17_v7_).f28
            d_117_v89_ = nw15_
            d_117_v89_ = d_117_v89_
            d_119_v90_: _dafny.Seq
            d_119_v90_ = _dafny.SeqWithoutIsStrInference([d_17_v7_, d_17_v7_, d_17_v7_])
            d_120_v91_: _dafny.Map
            d_120_v91_ = _dafny.Map({d_11_v1_: (d_79_v61_)[default__.safeIndex(488, (d_79_v61_).length(0))]})
            d_17_v7_ = (d_119_v90_)[default__.safeIndex(default__.fm3(502, False, d_11_v1_, d_120_v91_, globalState), len(d_119_v90_))]
        d_121_v92_: D0
        d_121_v92_ = D0_DC1(((d_17_v7_).f28 if not((d_17_v7_).f28) else d_10_v0_), d_11_v1_)
        source0_ = d_121_v92_
        if source0_.is_DC1:
            d_122___mcc_h0_ = source0_.cf1
            d_123___mcc_h1_ = source0_.cf2
            d_124_cf2_ = d_123___mcc_h1_
            d_125_cf1_ = d_122___mcc_h0_
            d_11_v1_ = d_11_v1_
            d_126_v93_: _dafny.Map
            d_126_v93_ = _dafny.Map({d_124_cf2_: not(d_125_cf1_)})
            d_124_cf2_ = (0) - (default__.fm3((d_124_cf2_) - (d_124_cf2_), (d_17_v7_).f28, d_11_v1_, d_126_v93_, globalState))
            pat_let_tv8_ = d_11_v1_
            d_127_v94_: _dafny.Map
            def iife27_(_pat_let10_0):
                def iife28_(d_128_dt__update__tmp_h6_):
                    def iife29_(_pat_let11_0):
                        def iife30_(d_129_dt__update_hcf2_h0_):
                            return D0_DC1((d_128_dt__update__tmp_h6_).cf1, d_129_dt__update_hcf2_h0_)
                        return iife30_(_pat_let11_0)
                    return iife29_(pat_let_tv8_)
                return iife28_(_pat_let10_0)
            d_127_v94_ = _dafny.Map({iife27_(d_121_v92_): d_126_v93_})
            d_130_v95_: _dafny.Seq
            d_130_v95_ = _dafny.SeqWithoutIsStrInference([d_14_v4_, d_14_v4_])
            d_131_v96_: _dafny.Seq
            d_131_v96_ = _dafny.SeqWithoutIsStrInference([d_11_v1_, d_124_cf2_, d_11_v1_, d_11_v1_])
            rhs30_ = (d_127_v94_) | ((d_127_v94_) | (d_127_v94_))
            rhs31_ = not (((d_130_v95_)[default__.safeIndex(d_124_cf2_, len(d_130_v95_))]) == (d_14_v4_)) or (not((d_17_v7_).f28))
            rhs32_ = default__.safeDivisionInt(d_124_cf2_, (0) - (len(d_131_v96_)))
            rhs33_ = d_124_cf2_
            lhs13_ = globalState
            d_127_v94_ = rhs30_
            lhs13_.f0 = rhs31_
            d_124_cf2_ = rhs32_
            d_11_v1_ = rhs33_
            if (d_11_v1_) != (d_11_v1_):
                d_132_v97_: C0
                nw16_ = C0()
                nw16_.ctor__((d_17_v7_).f28, (d_17_v7_).f29)
                d_132_v97_ = nw16_
                index18_ = default__.safeIndex(804, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index18_] = (507) * (d_124_cf2_)
                index19_ = default__.safeIndex(804, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index19_] = len((d_12_v2_) | (d_12_v2_))
                d_133_v98_: _dafny.Seq
                d_133_v98_ = _dafny.SeqWithoutIsStrInference([d_79_v61_])
                d_134_v99_: _dafny.MultiSet
                d_134_v99_ = _dafny.MultiSet([d_11_v1_])
                d_135_v100_: _dafny.Seq
                d_135_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckgrfoe"))
                d_136_v101_: _dafny.Map
                d_136_v101_ = _dafny.Map({d_11_v1_: d_135_v100_})
                d_137_v102_: D0
                d_137_v102_ = D0_DC0(d_134_v99_)
                d_138_v103_: _dafny.Map
                d_138_v103_ = _dafny.Map({((d_17_v7_).f29)[default__.safeIndex(804, ((d_17_v7_).f29).length(0))]: d_134_v99_})
                d_139_v104_: _dafny.Array
                nw17_ = _dafny.Array(None, 18)
                nw17_[int(0)] = (d_134_v99_) | (d_134_v99_)
                nw17_[int(1)] = ((d_134_v99_).set(187, default__.abs(((d_17_v7_).f29)[default__.safeIndex(804, ((d_17_v7_).f29).length(0))]))) - (d_134_v99_)
                nw17_[int(2)] = d_134_v99_
                nw17_[int(3)] = (_dafny.MultiSet([len(d_135_v100_), d_124_cf2_, d_11_v1_, len(d_135_v100_), d_11_v1_])).intersection(d_134_v99_)
                nw17_[int(4)] = _dafny.MultiSet([(0) - ((0) - ((0) - (len(d_136_v101_))))])
                nw17_[int(5)] = d_134_v99_
                nw17_[int(6)] = (_dafny.MultiSet(d_131_v96_)).set(d_11_v1_, default__.abs(d_11_v1_))
                nw17_[int(7)] = d_134_v99_
                nw17_[int(8)] = d_134_v99_
                nw17_[int(9)] = _dafny.MultiSet(d_131_v96_)
                nw17_[int(10)] = (d_134_v99_).intersection(d_134_v99_)
                nw17_[int(11)] = (d_137_v102_).cf0
                nw17_[int(12)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_11_v1_]))
                nw17_[int(13)] = d_134_v99_
                nw17_[int(14)] = d_134_v99_
                nw17_[int(15)] = d_134_v99_
                nw17_[int(16)] = (((d_138_v103_)[705] if (705) in (d_138_v103_) else d_134_v99_)).set(d_124_cf2_, default__.abs(d_11_v1_))
                nw17_[int(17)] = _dafny.MultiSet([d_124_cf2_, -663])
                d_139_v104_ = nw17_
                index20_ = default__.safeIndex(414, (d_139_v104_).length(0))
                (d_139_v104_)[index20_] = _dafny.MultiSet(d_131_v96_)
                d_140_v105_: str
                d_140_v105_ = _dafny.CodePoint('b')
                d_141_v106_: _dafny.Map
                d_141_v106_ = _dafny.Map({True: d_140_v105_})
                d_142_v107_: _dafny.MultiSet
                d_142_v107_ = _dafny.MultiSet([_dafny.Map({d_125_cf1_: d_140_v105_}), _dafny.Map({True: _dafny.CodePoint('k')}), (d_141_v106_).set((d_17_v7_).f28, d_140_v105_), d_141_v106_, (d_141_v106_ if (d_132_v97_).f28 else d_141_v106_)])
                d_143_v108_: _dafny.Map
                d_143_v108_ = _dafny.Map({len(d_135_v100_): d_132_v97_})
                d_144_v109_: D7
                d_144_v109_ = D7_DC17(((d_143_v108_)[len(d_13_v3_)] if (len(d_13_v3_)) in (d_143_v108_) else d_17_v7_))
                d_145_v110_: _dafny.Seq
                d_145_v110_ = _dafny.SeqWithoutIsStrInference([(d_144_v109_).cf26])
                d_146_v111_: _dafny.Seq
                d_146_v111_ = _dafny.SeqWithoutIsStrInference([d_135_v100_, d_135_v100_])
                d_147_v112_: _dafny.MultiSet
                d_147_v112_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")), d_135_v100_, d_135_v100_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), _dafny.SeqWithoutIsStrInference([d_140_v105_ for d_148_i8_ in range(default__.abs(55))])])
                d_149_v113_: D7
                d_149_v113_ = D7_DC18((d_17_v7_).f28)
                index21_ = default__.safeIndex(804, ((d_17_v7_).f29).length(0))
                index22_ = default__.safeIndex(414, (d_139_v104_).length(0))
                rhs34_ = d_133_v98_
                rhs35_ = len(((d_145_v110_) + (_dafny.SeqWithoutIsStrInference([d_132_v97_, d_132_v97_, d_132_v97_]))).set(default__.safeIndex(381, len((d_145_v110_) + (_dafny.SeqWithoutIsStrInference([d_132_v97_, d_132_v97_, d_132_v97_])))), d_132_v97_))
                rhs36_ = d_134_v99_
                rhs37_ = ((default__.fm11(-642, _dafny.CodePoint('a'), globalState) if (d_132_v97_).f28 else d_147_v112_)).ispropersubset((_dafny.MultiSet(d_146_v111_)) - (d_147_v112_))
                rhs38_ = default__.fm12(d_11_v1_, ((d_17_v7_).f29)[default__.safeIndex(804, ((d_17_v7_).f29).length(0))], ((d_78_v60_)[(d_132_v97_).f28] if ((d_132_v97_).f28) in (d_78_v60_) else ((d_17_v7_).f29)[default__.safeIndex(804, ((d_17_v7_).f29).length(0))]), d_149_v113_, globalState)
                lhs14_ = (d_17_v7_).f29
                lhs15_ = default__.safeIndex(804, ((d_17_v7_).f29).length(0))
                lhs16_ = d_139_v104_
                lhs17_ = default__.safeIndex(414, (d_139_v104_).length(0))
                lhs18_ = globalState
                d_133_v98_ = rhs34_
                lhs14_[lhs15_] = rhs35_
                lhs16_[lhs17_] = rhs36_
                lhs18_.f0 = rhs37_
                d_142_v107_ = rhs38_
                index23_ = default__.safeIndex(991, (d_79_v61_).length(0))
                (d_79_v61_)[index23_] = (d_17_v7_).f28
                index24_ = default__.safeIndex(991, (d_79_v61_).length(0))
                (d_79_v61_)[index24_] = (d_17_v7_).f28
                index25_ = default__.safeIndex(804, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index25_] = (d_124_cf2_) + ((d_11_v1_) * (((d_17_v7_).f29)[default__.safeIndex(804, ((d_17_v7_).f29).length(0))]))
            elif True:
                d_126_v93_ = (d_126_v93_).set(d_11_v1_, (d_17_v7_).f28)
                d_17_v7_ = d_17_v7_
                d_150_v114_: _dafny.Array
                d_150_v114_ = d_79_v61_
                d_151_v115_: _dafny.Array
                nw18_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_151_v115_ = nw18_
                index26_ = default__.safeIndex(346, (d_151_v115_).length(0))
                (d_151_v115_)[index26_] = _dafny.CodePoint('k')
                d_152_v116_: str
                d_152_v116_ = _dafny.CodePoint('p')
                index27_ = default__.safeIndex(346, (d_151_v115_).length(0))
                rhs39_ = d_79_v61_
                rhs40_ = d_152_v116_
                lhs19_ = d_151_v115_
                lhs20_ = default__.safeIndex(346, (d_151_v115_).length(0))
                d_150_v114_ = rhs39_
                lhs19_[lhs20_] = rhs40_
                index28_ = default__.safeIndex(767, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index28_] = 21
                d_153_v117_: _dafny.MultiSet
                d_153_v117_ = _dafny.MultiSet([-324])
                index29_ = default__.safeIndex(767, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index29_] = (0) - (((d_153_v117_)[d_11_v1_] if (d_11_v1_) in (d_153_v117_) else d_11_v1_))
                d_154_v118_: _dafny.Seq
                d_154_v118_ = _dafny.SeqWithoutIsStrInference([d_79_v61_, (d_79_v61_ if (d_17_v7_).f28 else d_79_v61_), d_79_v61_])
                index30_ = default__.safeIndex(767, ((d_17_v7_).f29).length(0))
                rhs41_ = d_11_v1_
                rhs42_ = (d_154_v118_)[default__.safeIndex(((d_17_v7_).f29)[default__.safeIndex(767, ((d_17_v7_).f29).length(0))], len(d_154_v118_))]
                rhs43_ = d_11_v1_
                lhs21_ = (d_17_v7_).f29
                lhs22_ = default__.safeIndex(767, ((d_17_v7_).f29).length(0))
                lhs21_[lhs22_] = rhs41_
                d_79_v61_ = rhs42_
                d_124_cf2_ = rhs43_
        elif source0_.is_DC0:
            d_155___mcc_h2_ = source0_.cf0
            d_156_cf0_ = d_155___mcc_h2_
            d_11_v1_ = default__.fm3((d_11_v1_) + (d_11_v1_), (d_17_v7_).f28, d_11_v1_, _dafny.Map({d_11_v1_: (d_17_v7_).f28}), globalState)
            d_157_v119_: D6
            d_157_v119_ = D6_DC15((d_17_v7_).f28, (d_17_v7_).f28, d_11_v1_)
            d_11_v1_ = (d_157_v119_).cf23
            d_11_v1_ = default__.safeDivisionInt(d_11_v1_, d_11_v1_)
            d_158_v120_: _dafny.Seq
            d_158_v120_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sudyiqf"))
            (globalState).f2 = (d_158_v120_) + ((d_158_v120_) + (d_158_v120_))
        elif True:
            d_159___mcc_h3_ = source0_.cf3
            d_160_cf3_ = d_159___mcc_h3_
            d_161_v121_: D1
            d_161_v121_ = D1_DC5(D1_DC3((d_17_v7_).f29))
            d_162_v122_: D1
            d_162_v122_ = D1_DC5(d_161_v121_)
            d_163_v123_: D1
            d_163_v123_ = D1_DC5(d_162_v122_)
            d_164_v124_: _dafny.Seq
            d_164_v124_ = _dafny.SeqWithoutIsStrInference([d_163_v123_])
            d_165_v125_: _dafny.Seq
            d_165_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxedcgbw"))
            d_166_v126_: D6
            d_166_v126_ = D6_DC16(d_164_v124_, d_165_v125_)
            source1_ = d_166_v126_
            if source1_.is_DC14:
                d_167___mcc_h4_ = source1_.cf20
                d_168_cf20_ = d_167___mcc_h4_
                d_169_v127_: _dafny.Map
                d_169_v127_ = _dafny.Map({(d_17_v7_).f28: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmsp")))})
                d_11_v1_ = len(((_dafny.Map({(d_17_v7_).f28: d_11_v1_})) | (d_169_v127_)) | (d_169_v127_))
                index31_ = default__.safeIndex(878, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index31_] = d_11_v1_
                index32_ = default__.safeIndex(878, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index32_] = d_11_v1_
                d_11_v1_ = (((d_17_v7_).f29)[default__.safeIndex(878, ((d_17_v7_).f29).length(0))]) - (((d_17_v7_).f29)[default__.safeIndex(878, ((d_17_v7_).f29).length(0))])
                d_170_v128_: _dafny.Seq
                d_170_v128_ = _dafny.SeqWithoutIsStrInference([d_165_v125_])
                (globalState).f2 = (d_170_v128_)[default__.safeIndex(d_11_v1_, len(d_170_v128_))]
            elif source1_.is_DC15:
                d_171___mcc_h5_ = source1_.cf21
                d_172___mcc_h6_ = source1_.cf22
                d_173___mcc_h7_ = source1_.cf23
                d_174_cf23_ = d_173___mcc_h7_
                d_175_cf22_ = d_172___mcc_h6_
                d_176_cf21_ = d_171___mcc_h5_
                index33_ = default__.safeIndex(646, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index33_] = -981
                index34_ = default__.safeIndex(646, ((d_17_v7_).f29).length(0))
                ((d_17_v7_).f29)[index34_] = d_11_v1_
                (globalState).f2 = d_165_v125_
                (globalState).f0 = True
                (globalState).f0 = not((d_17_v7_).f28)
            elif source1_.is_DC16:
                d_177___mcc_h8_ = source1_.cf24
                d_178___mcc_h9_ = source1_.cf25
                d_179_cf25_ = d_178___mcc_h9_
                d_180_cf24_ = d_177___mcc_h8_
                d_181_v129_: _dafny.Set
                d_181_v129_ = _dafny.Set({d_17_v7_})
                d_182_v130_: _dafny.Map
                d_182_v130_ = _dafny.Map({d_181_v129_: (d_17_v7_).f28})
                d_183_v131_: _dafny.Array
                nw19_ = _dafny.Array(None, 13)
                nw19_[int(0)] = _dafny.Map({d_181_v129_: (d_17_v7_).f28})
                nw19_[int(1)] = (_dafny.Map({_dafny.Set({d_17_v7_}): not((d_17_v7_).f28)})) | (d_182_v130_)
                nw19_[int(2)] = (d_182_v130_) | (d_182_v130_)
                nw19_[int(3)] = d_182_v130_
                nw19_[int(4)] = d_182_v130_
                nw19_[int(5)] = d_182_v130_
                nw19_[int(6)] = (d_182_v130_) | (d_182_v130_)
                nw19_[int(7)] = d_182_v130_
                nw19_[int(8)] = (_dafny.Map({d_181_v129_: (d_17_v7_).f28})).set(_dafny.Set({d_17_v7_}), (d_17_v7_).f28)
                nw19_[int(9)] = (d_182_v130_).set(d_181_v129_, (d_17_v7_).f28)
                nw19_[int(10)] = d_182_v130_
                nw19_[int(11)] = d_182_v130_
                nw19_[int(12)] = d_182_v130_
                d_183_v131_ = nw19_
                index35_ = default__.safeIndex(24, (d_183_v131_).length(0))
                (d_183_v131_)[index35_] = (d_182_v130_) | (d_182_v130_)
                d_184_v132_: D7
                d_184_v132_ = D7_DC18((d_17_v7_).f28)
                d_185_v133_: D7
                d_185_v133_ = D7_DC19(d_184_v132_)
                d_186_v134_: D7
                d_186_v134_ = D7_DC19(d_184_v132_)
                index36_ = default__.safeIndex(24, (d_183_v131_).length(0))
                rhs44_ = d_182_v130_
                rhs45_ = (d_17_v7_ if d_10_v0_ else d_17_v7_)
                rhs46_ = default__.fm9(globalState)
                rhs47_ = d_186_v134_
                lhs23_ = d_183_v131_
                lhs24_ = default__.safeIndex(24, (d_183_v131_).length(0))
                lhs25_ = globalState
                lhs23_[lhs24_] = rhs44_
                d_17_v7_ = rhs45_
                lhs25_.f2 = rhs46_
                d_186_v134_ = rhs47_
                d_14_v4_ = d_14_v4_
                d_11_v1_ = default__.safeDivisionInt(len(d_179_cf25_), d_11_v1_)
                d_187_v135_: _dafny.Map
                d_187_v135_ = _dafny.Map({(0) - (default__.safeDivisionInt(d_11_v1_, d_11_v1_)): (d_17_v7_).f28})
                d_188_v136_: _dafny.Seq
                d_188_v136_ = _dafny.SeqWithoutIsStrInference([d_11_v1_])
                d_187_v135_ = (d_187_v135_).set(len(_dafny.Set({_dafny.SeqWithoutIsStrInference([294]), d_188_v136_, d_188_v136_})), (d_17_v7_).f28)
            elif True:
                d_189___mcc_h10_ = source1_.cf19
                d_190_cf19_ = d_189___mcc_h10_
                (globalState).f2 = (d_165_v125_) + ((d_166_v126_).cf25)
                d_191_v137_: C0
                nw20_ = C0()
                nw20_.ctor__((d_17_v7_).f28, d_16_v6_)
                d_191_v137_ = nw20_
                d_192_v138_: D1
                d_192_v138_ = D1_DC4((d_11_v1_) <= (d_11_v1_))
                rhs48_ = default__.safeModuloInt(d_11_v1_, d_11_v1_)
                rhs49_ = d_192_v138_
                rhs50_ = (d_17_v7_).fm2((d_17_v7_).f28, globalState)
                rhs51_ = (not((d_17_v7_).f28)) or ((d_191_v137_).f28)
                rhs52_ = (d_191_v137_).f29
                lhs26_ = globalState
                lhs27_ = globalState
                d_11_v1_ = rhs48_
                d_192_v138_ = rhs49_
                lhs26_.f0 = rhs50_
                lhs27_.f0 = rhs51_
                d_16_v6_ = rhs52_
                d_11_v1_ = (d_11_v1_) - ((0) - (d_11_v1_))
            d_193_v139_: str
            d_193_v139_ = _dafny.CodePoint('g')
            d_194_v140_: _dafny.Array
            nw21_ = _dafny.Array(None, 2)
            nw21_[int(0)] = d_193_v139_
            nw21_[int(1)] = d_193_v139_
            d_194_v140_ = nw21_
            d_195_v141_: _dafny.Map
            d_195_v141_ = _dafny.Map({d_193_v139_: d_194_v140_})
            d_196_v142_: _dafny.Seq
            d_196_v142_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_193_v139_: d_194_v140_})])
            d_197_v143_: _dafny.Seq
            d_197_v143_ = _dafny.SeqWithoutIsStrInference([d_17_v7_])
            d_198_v144_: _dafny.MultiSet
            d_198_v144_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_199_i9_ in range(default__.abs(597))])), len(d_197_v143_)])
            d_200_v145_: D8
            d_200_v145_ = D8_DC20(d_15_v5_)
            d_201_v146_: _dafny.Array
            nw22_ = _dafny.Array(None, 24)
            nw22_[int(0)] = d_195_v141_
            nw22_[int(1)] = d_195_v141_
            nw22_[int(2)] = (d_196_v142_)[default__.safeIndex(((d_198_v144_)[d_11_v1_] if (d_11_v1_) in (d_198_v144_) else len((d_200_v145_).cf29)), len(d_196_v142_))]
            nw22_[int(3)] = d_195_v141_
            nw22_[int(4)] = d_195_v141_
            nw22_[int(5)] = _dafny.Map({d_193_v139_: d_194_v140_})
            nw22_[int(6)] = d_195_v141_
            nw22_[int(7)] = d_195_v141_
            nw22_[int(8)] = (d_195_v141_) | (d_195_v141_)
            nw22_[int(9)] = (d_196_v142_)[default__.safeIndex(-425, len(d_196_v142_))]
            nw22_[int(10)] = d_195_v141_
            nw22_[int(11)] = d_195_v141_
            nw22_[int(12)] = (d_196_v142_)[default__.safeIndex(267, len(d_196_v142_))]
            nw22_[int(13)] = (_dafny.Map({d_193_v139_: d_194_v140_})) | (d_195_v141_)
            nw22_[int(14)] = _dafny.Map({d_193_v139_: d_194_v140_})
            nw22_[int(15)] = d_195_v141_
            nw22_[int(16)] = d_195_v141_
            nw22_[int(17)] = d_195_v141_
            nw22_[int(18)] = (d_195_v141_).set(d_193_v139_, d_194_v140_)
            nw22_[int(19)] = d_195_v141_
            nw22_[int(20)] = d_195_v141_
            nw22_[int(21)] = d_195_v141_
            nw22_[int(22)] = (_dafny.Map({d_193_v139_: d_194_v140_})) | ((d_195_v141_).set(d_193_v139_, d_194_v140_))
            nw22_[int(23)] = ((d_195_v141_).set(d_193_v139_, d_194_v140_)) | (d_195_v141_)
            d_201_v146_ = nw22_
            index37_ = default__.safeIndex(957, (d_201_v146_).length(0))
            (d_201_v146_)[index37_] = d_195_v141_
            index38_ = default__.safeIndex(957, (d_201_v146_).length(0))
            (d_201_v146_)[index38_] = (d_195_v141_) | (_dafny.Map({d_193_v139_: d_194_v140_}))
            d_202_v147_: _dafny.Map
            d_202_v147_ = _dafny.Map({d_11_v1_: (d_17_v7_).f28})
            d_11_v1_ = default__.safeModuloInt(d_11_v1_, (0) - ((0) - ((len(d_202_v147_)) - (d_11_v1_))))
            index39_ = default__.safeIndex(649, ((d_17_v7_).f29).length(0))
            ((d_17_v7_).f29)[index39_] = 346
            index40_ = default__.safeIndex(649, ((d_17_v7_).f29).length(0))
            ((d_17_v7_).f29)[index40_] = (d_11_v1_) - ((d_121_v92_).cf2)
        d_203_v148_: _dafny.Map
        d_203_v148_ = _dafny.Map({d_11_v1_: d_11_v1_})
        d_204_v149_: _dafny.Map
        d_204_v149_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_203_v148_, d_203_v148_]): (d_17_v7_).f28})
        d_205_v150_: str
        d_205_v150_ = _dafny.CodePoint('q')
        d_206_v151_: _dafny.Map
        d_206_v151_ = _dafny.Map({(d_17_v7_).f28: d_205_v150_})
        d_204_v149_ = (d_204_v149_).set(_dafny.SeqWithoutIsStrInference([(d_203_v148_).set(d_11_v1_, len(d_206_v151_))]), (d_17_v7_).f28)

    @staticmethod
    def Main(noArgsParameter__):
        d_207_v0_: bool
        d_207_v0_ = False
        d_208_v1_: _dafny.Map
        d_208_v1_ = _dafny.Map({d_207_v0_: d_207_v0_})
        d_209_v2_: _dafny.Set
        d_209_v2_ = _dafny.Set({d_207_v0_})
        d_210_v3_: _dafny.Seq
        d_210_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk"))
        d_211_v4_: _dafny.Seq
        d_211_v4_ = _dafny.SeqWithoutIsStrInference([d_207_v0_, d_207_v0_, d_207_v0_, d_207_v0_])
        d_212_v5_: int
        d_212_v5_ = 228
        d_213_v6_: _dafny.Seq
        d_213_v6_ = _dafny.SeqWithoutIsStrInference([d_212_v5_])
        d_214_v7_: _dafny.Set
        d_214_v7_ = _dafny.Set({len(d_213_v6_)})
        d_215_v8_: _dafny.Seq
        d_215_v8_ = _dafny.SeqWithoutIsStrInference([(0) - (d_212_v5_), len(d_214_v7_)])
        d_216_v9_: str
        d_216_v9_ = _dafny.CodePoint('v')
        d_217_v10_: _dafny.Map
        d_217_v10_ = _dafny.Map({(0) - (d_212_v5_): d_216_v9_})
        d_218_v11_: _dafny.MultiSet
        d_218_v11_ = _dafny.MultiSet([d_217_v10_])
        d_219_v12_: _dafny.Array
        def lambda1_(d_220_v5_):
            def lambda2_(d_221_i0_):
                return (d_221_i0_) + (d_220_v5_)

            return lambda2_

        init1_ = lambda1_(d_212_v5_)
        nw23_ = _dafny.Array(None, 13)
        for i0_1_ in range(nw23_.length(0)):
            nw23_[i0_1_] = init1_(i0_1_)
        d_219_v12_ = nw23_
        d_222_v13_: _dafny.Map
        d_222_v13_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybllip")): _dafny.MultiSet([d_212_v5_, d_212_v5_])})
        d_223_v15_: _dafny.Seq
        d_223_v15_ = _dafny.SeqWithoutIsStrInference([d_213_v6_])
        d_224_v16_: D0
        d_224_v16_ = D0_DC1(d_207_v0_, len(d_208_v1_))
        d_225_v17_: D0
        d_225_v17_ = D0_DC2(d_224_v16_)
        d_226_v18_: _dafny.MultiSet
        d_226_v18_ = _dafny.MultiSet([d_225_v17_, d_225_v17_])
        d_227_v19_: _dafny.MultiSet
        d_227_v19_ = _dafny.MultiSet([d_226_v18_])
        d_228_v20_: _dafny.Map
        d_228_v20_ = _dafny.Map({((d_227_v19_)[d_226_v18_] if (d_226_v18_) in (d_227_v19_) else d_212_v5_): d_207_v0_})
        d_229_globalState_: GlobalState
        nw24_ = GlobalState()
        def iife31_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "incqt")) for d_230_i1_ in range(default__.abs(-191))])).set(default__.safeIndex(d_212_v5_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "incqt")) for d_231_i1_ in range(default__.abs(-191))]))), d_210_v3_)).Elements:
                d_232_v14_: _dafny.Seq = compr_7_
                if (d_232_v14_) in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "incqt")) for d_230_i1_ in range(default__.abs(-191))])).set(default__.safeIndex(d_212_v5_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "incqt")) for d_231_i1_ in range(default__.abs(-191))]))), d_210_v3_)):
                    coll7_[d_232_v14_] = (D0_DC0(_dafny.MultiSet([d_212_v5_]))).cf0
            return _dafny.Map(coll7_)
        nw24_.ctor__(True, (_dafny.Set({((d_208_v1_)[d_207_v0_] if (d_207_v0_) in (d_208_v1_) else d_207_v0_)})) - (d_209_v2_), d_210_v3_, d_211_v4_, d_215_v8_, -929, 851, True, d_218_v11_, 90, -909, 917, 336, d_219_v12_, 206, 664, -432, (d_222_v13_) | (iife31_()
        ), True, d_223_v15_, (d_210_v3_).set(default__.safeIndex(d_212_v5_, len(d_210_v3_)), d_216_v9_), -661, -89, _dafny.SeqWithoutIsStrInference([d_212_v5_]), True, 849, False, d_228_v20_)
        d_229_globalState_ = nw24_
        default__.m0(d_229_globalState_)
        if (len(d_211_v4_)) == (len(d_228_v20_)):
            d_212_v5_ = len(d_210_v3_)
            (d_229_globalState_).f0 = not(not(default__.fm0(d_229_globalState_)))
            d_233_v21_: C0
            nw25_ = C0()
            nw25_.ctor__(d_207_v0_, d_219_v12_)
            d_233_v21_ = nw25_
            d_234_v22_: D0
            d_234_v22_ = D0_DC1(d_207_v0_, default__.fm3(d_212_v5_, d_207_v0_, default__.fm3(d_212_v5_, (d_233_v21_).f28, 544, d_228_v20_, d_229_globalState_), d_228_v20_, d_229_globalState_))
            d_235_v23_: D1
            d_235_v23_ = D1_DC3(d_219_v12_)
            d_236_v24_: C0
            nw26_ = C0()
            nw26_.ctor__((d_234_v22_).cf1, (d_235_v23_).cf4)
            d_236_v24_ = nw26_
            d_236_v24_ = d_236_v24_
            d_216_v9_ = (d_236_v24_).fm1(D0_DC1((d_233_v21_).f28, d_212_v5_), d_212_v5_, d_212_v5_, len(d_210_v3_), d_229_globalState_)
        elif True:
            index41_ = default__.safeIndex(59, (d_219_v12_).length(0))
            (d_219_v12_)[index41_] = len(d_214_v7_)
            index42_ = default__.safeIndex(59, (d_219_v12_).length(0))
            (d_219_v12_)[index42_] = (0) - (len((d_208_v1_) | ((d_208_v1_).set(d_207_v0_, d_207_v0_))))
            d_237_v25_: _dafny.Array
            def lambda3_(d_238_v8_):
                def lambda4_(d_239_i2_):
                    return default__.safeModuloInt(d_239_i2_, len(d_238_v8_))

                return lambda4_

            init2_ = lambda3_(d_215_v8_)
            nw27_ = _dafny.Array(None, 21)
            for i0_2_ in range(nw27_.length(0)):
                nw27_[i0_2_] = init2_(i0_2_)
            d_237_v25_ = nw27_
            d_240_v26_: C0
            nw28_ = C0()
            nw28_.ctor__(d_207_v0_, d_237_v25_)
            d_240_v26_ = nw28_
            (d_229_globalState_).f0 = (d_240_v26_).f28
            d_212_v5_ = (d_219_v12_)[default__.safeIndex(59, (d_219_v12_).length(0))]
            d_241_v27_: _dafny.Array
            nw29_ = _dafny.Array(False, 11)
            d_241_v27_ = nw29_
            d_242_v29_: _dafny.MultiSet
            def iife32_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([((D2_DC6(_dafny.MultiSet([(d_240_v26_).f28]))).cf7).cardinality for d_243_i3_ in range(default__.abs(-392))])).Elements:
                    d_244_v28_: int = compr_8_
                    if (d_244_v28_) in (_dafny.SeqWithoutIsStrInference([((D2_DC6(_dafny.MultiSet([(d_240_v26_).f28]))).cf7).cardinality for d_243_i3_ in range(default__.abs(-392))])):
                        coll8_[default__.safeDivisionInt(d_244_v28_, d_212_v5_)] = (d_240_v26_).f28
                return _dafny.Map(coll8_)
            d_242_v29_ = _dafny.MultiSet([d_228_v20_, _dafny.Map({(d_219_v12_)[default__.safeIndex(59, (d_219_v12_).length(0))]: (d_240_v26_).f28}), iife32_()
            , d_228_v20_, d_228_v20_])
            index43_ = default__.safeIndex(414, (d_241_v27_).length(0))
            (d_241_v27_)[index43_] = ((d_242_v29_).cardinality) < ((d_219_v12_)[default__.safeIndex(59, (d_219_v12_).length(0))])
            index44_ = default__.safeIndex(414, (d_241_v27_).length(0))
            (d_241_v27_)[index44_] = (d_240_v26_).f28
        d_207_v0_ = d_207_v0_
        d_212_v5_ = d_212_v5_
        d_245_v30_: _dafny.Seq
        d_245_v30_ = _dafny.SeqWithoutIsStrInference([d_210_v3_, d_210_v3_])
        d_245_v30_ = (d_245_v30_) + ((d_245_v30_) + (d_245_v30_))
        d_246_v31_: _dafny.MultiSet
        d_246_v31_ = _dafny.MultiSet([default__.safeDivisionInt((0) - (d_212_v5_), (0) - (d_212_v5_))])
        d_246_v31_ = (default__.fm4(d_229_globalState_)) - (d_246_v31_)
        d_247_v32_: D3
        d_247_v32_ = D3_DC9(d_228_v20_)
        (d_229_globalState_).f0 = (not((len(_dafny.Map({d_207_v0_: True}))) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_212_v5_])))) if (d_212_v5_) <= (default__.fm3((0) - (d_212_v5_), d_207_v0_, d_212_v5_, (d_247_v32_).cf11, d_229_globalState_)) else not (d_207_v0_) or (not(d_207_v0_)))
        d_248_v33_: _dafny.Array
        def lambda5_(d_249_i4_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skjjyean"))

        init3_ = lambda5_
        nw30_ = _dafny.Array(None, 20)
        for i0_3_ in range(nw30_.length(0)):
            nw30_[i0_3_] = init3_(i0_3_)
        d_248_v33_ = nw30_
        d_250_v34_: _dafny.Map
        d_250_v34_ = _dafny.Map({d_248_v33_: d_212_v5_})
        d_250_v34_ = (d_250_v34_).set(d_248_v33_, d_212_v5_)
        d_212_v5_ = (d_212_v5_) + (len(d_210_v3_))
        (d_229_globalState_).f0 = d_207_v0_
        (d_229_globalState_).f0 = not(((d_208_v1_)[d_207_v0_] if (d_207_v0_) in (d_208_v1_) else (not(d_207_v0_)) == (d_207_v0_)))
        d_251_v35_: _dafny.Map
        d_251_v35_ = _dafny.Map({d_207_v0_: d_219_v12_})
        d_252_v36_: C0
        nw31_ = C0()
        nw31_.ctor__(d_207_v0_, d_219_v12_)
        d_252_v36_ = nw31_
        d_253_v37_: _dafny.Set
        d_253_v37_ = _dafny.Set({d_252_v36_, d_252_v36_, d_252_v36_})
        d_254_v38_: _dafny.Map
        d_254_v38_ = _dafny.Map({d_251_v35_: (d_253_v37_).issubset(d_253_v37_)})
        d_255_v39_: _dafny.Seq
        d_255_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_252_v36_).f28: d_219_v12_})])
        d_256_v40_: _dafny.Seq
        d_256_v40_ = _dafny.SeqWithoutIsStrInference([(d_255_v39_)[default__.safeIndex(d_212_v5_, len(d_255_v39_))]])
        d_254_v38_ = (d_254_v38_).set(((d_256_v40_)[default__.safeIndex(d_212_v5_, len(d_256_v40_))]) | (d_251_v35_), True)
        d_257_v41_: _dafny.Map
        d_257_v41_ = _dafny.Map({(_dafny.Map({d_207_v0_: (d_252_v36_).f28})).set(True, d_207_v0_): d_252_v36_})
        hi0_ = len(d_257_v41_)
        for d_258_i5_ in range(len((d_245_v30_) + (d_245_v30_)), hi0_):
            (d_229_globalState_).f0 = not(not((d_214_v7_) != (d_214_v7_)))
            (d_229_globalState_).f0 = (d_252_v36_).f28
            d_259_v42_: _dafny.MultiSet
            d_259_v42_ = _dafny.MultiSet([d_252_v36_])
            d_260_v43_: C0
            nw32_ = C0()
            nw32_.ctor__((_dafny.MultiSet([d_252_v36_])).issubset(d_259_v42_), d_219_v12_)
            d_260_v43_ = nw32_
            d_261_v44_: _dafny.Map
            d_261_v44_ = _dafny.Map({d_258_i5_: _dafny.SeqWithoutIsStrInference([(d_252_v36_).f28, d_207_v0_, not((d_252_v36_).f28)])})
            d_211_v4_ = (((d_261_v44_)[len((_dafny.Map({d_260_v43_: d_212_v5_})) | ((_dafny.Map({d_252_v36_: -199})).set(d_260_v43_, d_258_i5_)))] if (len((_dafny.Map({d_260_v43_: d_212_v5_})) | ((_dafny.Map({d_252_v36_: -199})).set(d_260_v43_, d_258_i5_)))) in (d_261_v44_) else _dafny.SeqWithoutIsStrInference([not((d_252_v36_).f28)]))).set(default__.safeIndex((d_258_i5_) + (d_212_v5_), len(((d_261_v44_)[len((_dafny.Map({d_260_v43_: d_212_v5_})) | ((_dafny.Map({d_252_v36_: -199})).set(d_260_v43_, d_258_i5_)))] if (len((_dafny.Map({d_260_v43_: d_212_v5_})) | ((_dafny.Map({d_252_v36_: -199})).set(d_260_v43_, d_258_i5_)))) in (d_261_v44_) else _dafny.SeqWithoutIsStrInference([not((d_252_v36_).f28)])))), (d_252_v36_).f28)
        d_247_v32_ = d_247_v32_
        d_262_v45_: D3
        d_262_v45_ = D3_DC10((d_252_v36_).f28, 281, d_207_v0_, d_212_v5_, (d_252_v36_).f29)
        d_212_v5_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iy"))), default__.fm3(d_212_v5_, (d_252_v36_).f28, len(_dafny.SeqWithoutIsStrInference([not((d_262_v45_).cf14)])), default__.fm5(d_216_v9_, d_229_globalState_), d_229_globalState_))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_248_v33_).length(0)):
            d_263_i6_: int = guard_loop_0_
            if (True) and (((0) <= (d_263_i6_)) and ((d_263_i6_) < ((d_248_v33_).length(0)))):
                (d_248_v33_)[(d_263_i6_)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnqf"))
        _dafny.print(_dafny.string_of(d_207_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v1_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v2_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_210_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v4_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_212_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v6_) == (_dafny.SeqWithoutIsStrInference([228]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v7_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_215_v8_) == (_dafny.SeqWithoutIsStrInference([-228, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_216_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v10_) == (_dafny.Map({-228: _dafny.CodePoint('v')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v11_) == (_dafny.MultiSet([_dafny.Map({-228: _dafny.CodePoint('v')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v12_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v13_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybllip")): _dafny.MultiSet([228, 228])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_223_v15_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([228])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v16_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v16_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_225_v17_).cf3).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_225_v17_).cf3).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v18_) == (_dafny.MultiSet([D0_DC2(D0_DC1(False, 1)), D0_DC2(D0_DC1(False, 1))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v19_) == (_dafny.MultiSet([_dafny.MultiSet([D0_DC2(D0_DC1(False, 1)), D0_DC2(D0_DC1(False, 1))])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v20_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_229_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f1) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_229_globalState_.f2).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f3) == (_dafny.SeqWithoutIsStrInference([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_.f4) == (_dafny.SeqWithoutIsStrInference([-228, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f8) == (_dafny.MultiSet([_dafny.Map({-228: _dafny.CodePoint('v')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f13)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_.f17) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybllip")): _dafny.MultiSet([228, 228]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "incqt")): _dafny.MultiSet([228]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk")): _dafny.MultiSet([228])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f19) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([228])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_229_globalState_).f20).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f23) == (_dafny.SeqWithoutIsStrInference([228]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_globalState_).f27) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v30_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v31_) == (_dafny.MultiSet([462, -1, -650]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v32_).cf11) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_248_v33_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_250_v34_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_251_v35_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v36_).f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v36_).f29)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_253_v37_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_254_v38_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_255_v39_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_256_v40_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_257_v41_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v45_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v45_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v45_).cf14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v45_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v45_).cf16)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0))
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
        return lambda: D1_DC4(False)
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

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(int(0), int(0))
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

class D2_DC7(D2, NamedTuple('DC7', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, int(0), False, int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({self.cf18.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf24)}, {self.cf25.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC18(D7, NamedTuple('DC18', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f4: _dafny.Seq = _dafny.Seq({})
        self.f17: _dafny.Map = _dafny.Map({})
        self._f1: _dafny.Set = _dafny.Set({})
        self._f3: _dafny.Seq = _dafny.Seq({})
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: _dafny.MultiSet = _dafny.MultiSet({})
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f13: _dafny.Array = _dafny.Array(None, 0)
        self._f14: int = int(0)
        self._f15: int = int(0)
        self._f16: int = int(0)
        self._f18: bool = False
        self._f19: _dafny.Seq = _dafny.Seq({})
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f21: int = int(0)
        self._f22: int = int(0)
        self._f23: _dafny.Seq = _dafny.Seq({})
        self._f24: bool = False
        self._f25: int = int(0)
        self._f26: bool = False
        self._f27: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f27 = f27

    @property
    def f1(self):
        return self._f1
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
    @property
    def f8(self):
        return self._f8
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
    def f16(self):
        return self._f16
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
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27

class C0:
    def  __init__(self):
        self._f28: bool = False
        self._f29: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f28, f29):
        (self)._f28 = f28
        (self)._f29 = f29

    def fm1(self, p0, p1, p2, p3, globalState):
        if not(True):
            return _dafny.CodePoint('u')
        elif True:
            return _dafny.CodePoint('o')

    def fm2(self, p0, globalState):
        return (self).f28

    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29
