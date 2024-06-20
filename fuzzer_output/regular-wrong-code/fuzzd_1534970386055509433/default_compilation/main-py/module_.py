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
    def fm4(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('s')

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcjolb"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvvpxlotu")))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (42) - (-345)

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_0_i0_ in range(default__.abs(32))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uus"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1_i1_ in range(default__.abs(725))])))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([692, len(_dafny.Map({False: True}))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) for d_2_i0_ in range(default__.abs(604))])) + (_dafny.SeqWithoutIsStrInference([206 for d_3_i1_ in range(default__.abs(74))])))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        if True:
            return _dafny.SeqWithoutIsStrInference([not(not(not(not(False)))), not(False)])
        elif True:
            return _dafny.SeqWithoutIsStrInference([True])

    @staticmethod
    def m0(globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: bool = False
        d_4_v0_: int
        d_4_v0_ = 341
        d_5_v1_: _dafny.Seq
        d_5_v1_ = _dafny.SeqWithoutIsStrInference([d_4_v0_])
        d_6_v2_: str
        d_6_v2_ = _dafny.CodePoint('n')
        d_7_v3_: bool
        d_7_v3_ = False
        d_8_v4_: _dafny.Array
        nw0_ = _dafny.Array(None, 12)
        nw0_[int(0)] = d_5_v1_
        nw0_[int(1)] = d_5_v1_
        nw0_[int(2)] = d_5_v1_
        nw0_[int(3)] = d_5_v1_
        nw0_[int(4)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuojgho"))), d_4_v0_])
        nw0_[int(5)] = d_5_v1_
        nw0_[int(6)] = _dafny.SeqWithoutIsStrInference([(0) - (d_4_v0_) for d_9_i0_ in range(default__.abs(439))])
        nw0_[int(7)] = d_5_v1_
        nw0_[int(8)] = d_5_v1_
        nw0_[int(9)] = (d_5_v1_) + (d_5_v1_)
        nw0_[int(10)] = default__.fm8(d_6_v2_, d_7_v3_, d_5_v1_, globalState)
        nw0_[int(11)] = d_5_v1_
        d_8_v4_ = nw0_
        d_10_v5_: D3
        d_10_v5_ = D3_DC11((_dafny.SeqWithoutIsStrInference([d_4_v0_])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axdp"))) for d_11_i1_ in range(default__.abs(153))])))
        d_12_v6_: _dafny.Set
        d_12_v6_ = _dafny.Set({d_7_v3_, d_7_v3_, d_7_v3_, d_7_v3_, d_7_v3_})
        d_13_v7_: _dafny.Seq
        d_13_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekcwr"))
        rhs0_ = default__.fm5(d_7_v3_, (len(d_12_v6_)) < (len(_dafny.SeqWithoutIsStrInference([d_6_v2_ for d_14_i2_ in range(default__.abs(648))]))), d_4_v0_, d_4_v0_, globalState)
        rhs1_ = ((d_4_v0_ if d_7_v3_ else d_4_v0_)) == ((0) - (default__.fm6(d_7_v3_, d_7_v3_, False, globalState)))
        rhs2_ = d_8_v4_
        rhs3_ = (default__.fm6(True, False, d_7_v3_, globalState)) + (len(d_13_v7_))
        rhs4_ = d_10_v5_
        lhs0_ = globalState
        lhs1_ = globalState
        lhs0_.f20 = rhs0_
        lhs1_.f20 = rhs1_
        d_8_v4_ = rhs2_
        r2 = rhs3_
        d_10_v5_ = rhs4_
        d_15_v8_: _dafny.Array
        nw1_ = _dafny.Array(int(0), 25)
        d_15_v8_ = nw1_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_15_v8_).length(0)):
            d_16_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_16_i3_)) and ((d_16_i3_) < ((d_15_v8_).length(0)))):
                (d_15_v8_)[(d_16_i3_)] = default__.safeDivisionInt(d_16_i3_, d_4_v0_)
        d_17_v9_: _dafny.Array
        nw2_ = _dafny.Array(_dafny.Map({}), 10)
        d_17_v9_ = nw2_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_17_v9_).length(0)):
            d_18_i4_: int = guard_loop_1_
            if (True) and (((0) <= (d_18_i4_)) and ((d_18_i4_) < ((d_17_v9_).length(0)))):
                (d_17_v9_)[(d_18_i4_)] = _dafny.Map({d_4_v0_: D4_DC18()})
        (globalState).f20 = not(d_7_v3_)
        if (d_4_v0_) < (default__.safeDivisionInt(d_4_v0_, d_4_v0_)):
            d_19_v10_: D0
            d_19_v10_ = D0_DC1(not(d_7_v3_), d_7_v3_, d_4_v0_)
            d_20_v11_: _dafny.Seq
            d_20_v11_ = _dafny.SeqWithoutIsStrInference([d_7_v3_])
            d_21_v12_: _dafny.Map
            d_21_v12_ = _dafny.Map({d_4_v0_: d_4_v0_})
            d_22_v13_: D5
            d_22_v13_ = D5_DC21(d_21_v12_)
            d_23_v14_: D5
            d_23_v14_ = D5_DC22(d_22_v13_)
            d_24_v15_: _dafny.Map
            d_24_v15_ = _dafny.Map({-355: False})
            d_25_v16_: _dafny.Array
            nw3_ = _dafny.Array(None, 28)
            nw3_[int(0)] = not (True) or (d_7_v3_)
            nw3_[int(1)] = d_7_v3_
            nw3_[int(2)] = d_7_v3_
            nw3_[int(3)] = d_7_v3_
            nw3_[int(4)] = (d_7_v3_ if d_7_v3_ else d_7_v3_)
            nw3_[int(5)] = d_7_v3_
            nw3_[int(6)] = d_7_v3_
            nw3_[int(7)] = False
            nw3_[int(8)] = d_7_v3_
            nw3_[int(9)] = d_7_v3_
            nw3_[int(10)] = (d_7_v3_ if (d_19_v10_).cf1 else d_7_v3_)
            nw3_[int(11)] = d_7_v3_
            nw3_[int(12)] = (d_7_v3_) == (d_7_v3_)
            nw3_[int(13)] = True
            nw3_[int(14)] = not(not((len(_dafny.Map({len(d_20_v11_): d_23_v14_}))) != (len(_dafny.SeqWithoutIsStrInference([d_6_v2_ for d_26_i5_ in range(default__.abs(302))])))))
            nw3_[int(15)] = not (d_7_v3_) or (d_7_v3_)
            nw3_[int(16)] = d_7_v3_
            nw3_[int(17)] = d_7_v3_
            nw3_[int(18)] = d_7_v3_
            nw3_[int(19)] = d_7_v3_
            nw3_[int(20)] = d_7_v3_
            nw3_[int(21)] = d_7_v3_
            nw3_[int(22)] = True
            nw3_[int(23)] = not(d_7_v3_)
            nw3_[int(24)] = d_7_v3_
            nw3_[int(25)] = (d_6_v2_) not in (default__.fm7(d_7_v3_, 110, d_4_v0_, False, globalState))
            nw3_[int(26)] = (((d_24_v15_)[d_4_v0_] if (d_4_v0_) in (d_24_v15_) else True)) and (d_7_v3_)
            nw3_[int(27)] = False
            d_25_v16_ = nw3_
            index0_ = default__.safeIndex(698, (d_25_v16_).length(0))
            (d_25_v16_)[index0_] = True
            index1_ = default__.safeIndex(698, (d_25_v16_).length(0))
            rhs5_ = (d_4_v0_) - ((d_4_v0_) - (d_4_v0_))
            rhs6_ = d_7_v3_
            lhs2_ = globalState
            lhs3_ = d_25_v16_
            lhs4_ = default__.safeIndex(698, (d_25_v16_).length(0))
            lhs2_.f5 = rhs5_
            lhs3_[lhs4_] = rhs6_
            d_27_v17_: _dafny.Array
            nw4_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
            d_27_v17_ = nw4_
            d_27_v17_ = (d_27_v17_ if (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))] else d_27_v17_)
            index2_ = default__.safeIndex(698, (d_25_v16_).length(0))
            (d_25_v16_)[index2_] = (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))]
            index3_ = default__.safeIndex(698, (d_25_v16_).length(0))
            (d_25_v16_)[index3_] = ((0) - (d_4_v0_)) != (d_4_v0_)
            if (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))]:
                index4_ = default__.safeIndex(698, (d_25_v16_).length(0))
                (d_25_v16_)[index4_] = ((d_4_v0_) - (d_4_v0_)) < (d_4_v0_)
                index5_ = default__.safeIndex(698, (d_25_v16_).length(0))
                (d_25_v16_)[index5_] = d_7_v3_
                d_28_v18_: _dafny.Set
                d_28_v18_ = _dafny.Set({d_4_v0_, len(d_5_v1_)})
                d_7_v3_ = (d_28_v18_).ispropersubset(d_28_v18_)
                d_29_v19_: C0
                nw5_ = C0()
                nw5_.ctor__(d_4_v0_, not(default__.fm5((d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))], d_7_v3_, d_4_v0_, len(d_20_v11_), globalState)), (d_13_v7_) + (d_13_v7_))
                d_29_v19_ = nw5_
                d_30_v20_: _dafny.Map
                d_30_v20_ = _dafny.Map({True: d_29_v19_})
                d_31_v21_: D1
                d_31_v21_ = D1_DC4(((d_30_v20_)[d_7_v3_] if (d_7_v3_) in (d_30_v20_) else d_29_v19_))
                d_31_v21_ = d_31_v21_
            elif True:
                d_20_v11_ = (_dafny.SeqWithoutIsStrInference([(d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))], default__.fm5(d_7_v3_, d_7_v3_, d_4_v0_, d_4_v0_, globalState)])) + (default__.fm9((d_20_v11_).set(default__.safeIndex(d_4_v0_, len(d_20_v11_)), (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))]), len(_dafny.SeqWithoutIsStrInference([d_6_v2_ for d_32_i6_ in range(default__.abs(655))])), (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))], globalState))
                d_33_v22_: C0
                nw6_ = C0()
                nw6_.ctor__(len(d_12_v6_), False, _dafny.SeqWithoutIsStrInference([d_6_v2_ for d_34_i7_ in range(default__.abs(-589))]))
                d_33_v22_ = nw6_
                d_35_v23_: _dafny.Seq
                d_35_v23_ = _dafny.SeqWithoutIsStrInference([d_33_v22_])
                d_36_v24_: _dafny.Array
                nw7_ = _dafny.Array(None, 18)
                nw7_[int(0)] = d_33_v22_
                nw7_[int(1)] = d_33_v22_
                nw7_[int(2)] = d_33_v22_
                nw7_[int(3)] = d_33_v22_
                nw7_[int(4)] = d_33_v22_
                nw7_[int(5)] = d_33_v22_
                nw7_[int(6)] = (d_35_v23_)[default__.safeIndex((d_33_v22_).f26, len(d_35_v23_))]
                nw7_[int(7)] = (d_33_v22_ if (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))] else d_33_v22_)
                nw7_[int(8)] = d_33_v22_
                nw7_[int(9)] = (d_33_v22_ if d_7_v3_ else d_33_v22_)
                nw7_[int(10)] = d_33_v22_
                nw7_[int(11)] = d_33_v22_
                nw7_[int(12)] = d_33_v22_
                nw7_[int(13)] = d_33_v22_
                nw7_[int(14)] = d_33_v22_
                nw7_[int(15)] = d_33_v22_
                nw7_[int(16)] = d_33_v22_
                nw7_[int(17)] = d_33_v22_
                d_36_v24_ = nw7_
                index6_ = default__.safeIndex(929, (d_36_v24_).length(0))
                (d_36_v24_)[index6_] = (d_33_v22_ if (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))] else d_33_v22_)
                d_37_v25_: _dafny.Map
                d_37_v25_ = _dafny.Map({d_27_v17_: d_33_v22_})
                index7_ = default__.safeIndex(929, (d_36_v24_).length(0))
                (d_36_v24_)[index7_] = ((d_33_v22_ if d_7_v3_ else d_33_v22_) if ((d_33_v22_).f26) < (d_4_v0_) else ((d_37_v25_)[d_27_v17_] if (d_27_v17_) in (d_37_v25_) else d_33_v22_))
                d_38_v26_: _dafny.MultiSet
                d_38_v26_ = _dafny.MultiSet([(d_36_v24_)[default__.safeIndex(929, (d_36_v24_).length(0))], d_33_v22_])
                d_39_v27_: C0
                nw8_ = C0()
                nw8_.ctor__((365) + ((d_38_v26_).cardinality), d_7_v3_, d_13_v7_)
                d_39_v27_ = nw8_
                d_40_v28_: _dafny.MultiSet
                d_40_v28_ = _dafny.MultiSet([d_15_v8_])
                d_41_v29_: T0
                nw9_ = C0()
                nw9_.ctor__((d_40_v28_).cardinality, (d_25_v16_)[default__.safeIndex(698, (d_25_v16_).length(0))], _dafny.SeqWithoutIsStrInference([d_6_v2_ for d_42_i8_ in range(default__.abs(126))]))
                d_41_v29_ = nw9_
                d_43_v30_: _dafny.Map
                d_43_v30_ = _dafny.Map({d_41_v29_: (d_39_v27_).f26})
                d_43_v30_ = (d_43_v30_).set(d_41_v29_, len(_dafny.SeqWithoutIsStrInference([d_6_v2_ for d_44_i9_ in range(default__.abs(-65))])))
                d_45_v31_: C0
                nw10_ = C0()
                nw10_.ctor__((d_39_v27_).f26, d_7_v3_, d_13_v7_)
                d_45_v31_ = nw10_
        elif True:
            (globalState).f5 = d_4_v0_
            d_5_v1_ = d_5_v1_
            d_46_v32_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Set({}), 19)
            d_46_v32_ = nw11_
            d_47_v33_: _dafny.Array
            def lambda0_(d_48_i10_):
                return True

            init0_ = lambda0_
            nw12_ = _dafny.Array(None, 13)
            for i0_0_ in range(nw12_.length(0)):
                nw12_[i0_0_] = init0_(i0_0_)
            d_47_v33_ = nw12_
            d_49_v34_: D2
            d_49_v34_ = D2_DC8(d_47_v33_)
            d_50_v35_: D2
            d_50_v35_ = D2_DC10(d_49_v34_)
            d_51_v36_: _dafny.Map
            d_51_v36_ = _dafny.Map({d_50_v35_: d_4_v0_})
            d_52_v37_: C0
            nw13_ = C0()
            nw13_.ctor__((0) - (len(d_51_v36_)), default__.fm5(d_7_v3_, d_7_v3_, d_4_v0_, d_4_v0_, globalState), d_13_v7_)
            d_52_v37_ = nw13_
            d_53_v38_: _dafny.Array
            d_53_v38_ = d_46_v32_
            rhs7_ = (d_53_v38_)
            rhs8_ = d_52_v37_
            rhs9_ = ((d_52_v37_).f26) * (len(d_5_v1_))
            lhs5_ = globalState
            d_46_v32_ = rhs7_
            d_52_v37_ = rhs8_
            lhs5_.f5 = rhs9_
            d_54_v39_: _dafny.Map
            d_54_v39_ = _dafny.Map({d_7_v3_: d_7_v3_})
            d_55_v40_: _dafny.Seq
            d_55_v40_ = _dafny.SeqWithoutIsStrInference([d_7_v3_])
            d_54_v39_ = _dafny.Map({d_7_v3_: (d_55_v40_)[default__.safeIndex((d_52_v37_).f26, len(d_55_v40_))]})
            d_56_v41_: _dafny.Map
            d_56_v41_ = _dafny.Map({True: d_6_v2_})
            d_57_v42_: D4
            d_57_v42_ = D4_DC17(len(d_56_v41_), d_52_v37_)
            d_58_v43_: _dafny.Map
            d_58_v43_ = _dafny.Map({(d_57_v42_).cf29: d_7_v3_})
            d_58_v43_ = (d_58_v43_).set(d_4_v0_, d_7_v3_)
        d_59_v44_: T0
        nw14_ = C0()
        nw14_.ctor__(d_4_v0_, default__.fm5(d_7_v3_, d_7_v3_, -857, default__.fm6(d_7_v3_, d_7_v3_, d_7_v3_, globalState), globalState), d_13_v7_)
        d_59_v44_ = nw14_
        d_60_v45_: _dafny.Map
        d_60_v45_ = _dafny.Map({d_59_v44_: -276})
        d_61_v46_: _dafny.Map
        d_61_v46_ = _dafny.Map({d_60_v45_: _dafny.SeqWithoutIsStrInference([d_6_v2_ for d_62_i11_ in range(default__.abs(857))])})
        (globalState).f4 = default__.fm7(d_7_v3_, (0) - (len(d_61_v46_)), d_4_v0_, d_7_v3_, globalState)
        r0 = (d_13_v7_) + ((d_13_v7_) + (_dafny.SeqWithoutIsStrInference([d_6_v2_ for d_63_i12_ in range(default__.abs(217))])))
        d_64_v47_: _dafny.Array
        nw15_ = _dafny.Array(_dafny.Array(None, 0), 9)
        d_64_v47_ = nw15_
        r1 = d_64_v47_
        r2 = d_4_v0_
        d_65_v48_: C0
        nw16_ = C0()
        nw16_.ctor__(d_4_v0_, d_7_v3_, d_13_v7_)
        d_65_v48_ = nw16_
        d_66_v49_: _dafny.MultiSet
        d_66_v49_ = _dafny.MultiSet([d_65_v48_, d_65_v48_, d_65_v48_])
        d_67_v50_: _dafny.Seq
        d_67_v50_ = _dafny.SeqWithoutIsStrInference([d_65_v48_, d_65_v48_])
        r3 = (((d_66_v49_)[(d_67_v50_)[default__.safeIndex((d_65_v48_).f26, len(d_67_v50_))]] if ((d_67_v50_)[default__.safeIndex((d_65_v48_).f26, len(d_67_v50_))]) in (d_66_v49_) else d_4_v0_)) > ((d_65_v48_).f26)
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_68_v0_: _dafny.Seq
        d_68_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkyq"))
        d_69_v1_: bool
        d_69_v1_ = False
        d_70_v2_: _dafny.Set
        d_70_v2_ = _dafny.Set({d_69_v1_, d_69_v1_})
        d_71_v3_: _dafny.Array
        nw17_ = _dafny.Array(None, 15)
        nw17_[int(0)] = not(False)
        nw17_[int(1)] = d_69_v1_
        nw17_[int(2)] = d_69_v1_
        nw17_[int(3)] = d_69_v1_
        nw17_[int(4)] = d_69_v1_
        nw17_[int(5)] = d_69_v1_
        nw17_[int(6)] = True
        nw17_[int(7)] = d_69_v1_
        nw17_[int(8)] = not(d_69_v1_)
        nw17_[int(9)] = d_69_v1_
        nw17_[int(10)] = d_69_v1_
        nw17_[int(11)] = d_69_v1_
        nw17_[int(12)] = d_69_v1_
        nw17_[int(13)] = d_69_v1_
        nw17_[int(14)] = d_69_v1_
        d_71_v3_ = nw17_
        d_72_v4_: D0
        d_72_v4_ = D0_DC0(d_68_v0_)
        d_73_v6_: int
        d_73_v6_ = 462
        d_74_v7_: _dafny.Seq
        d_74_v7_ = _dafny.SeqWithoutIsStrInference([d_73_v6_, d_73_v6_, d_73_v6_])
        d_75_v8_: str
        d_75_v8_ = _dafny.CodePoint('x')
        d_76_v9_: _dafny.Array
        nw18_ = _dafny.Array(None, 27)
        nw18_[int(0)] = d_75_v8_
        nw18_[int(1)] = d_75_v8_
        nw18_[int(2)] = d_75_v8_
        nw18_[int(3)] = _dafny.CodePoint('u')
        nw18_[int(4)] = d_75_v8_
        nw18_[int(5)] = d_75_v8_
        nw18_[int(6)] = d_75_v8_
        nw18_[int(7)] = d_75_v8_
        nw18_[int(8)] = d_75_v8_
        nw18_[int(9)] = d_75_v8_
        nw18_[int(10)] = d_75_v8_
        nw18_[int(11)] = d_75_v8_
        nw18_[int(12)] = d_75_v8_
        nw18_[int(13)] = (d_68_v0_)[default__.safeIndex(d_73_v6_, len(d_68_v0_))]
        nw18_[int(14)] = d_75_v8_
        nw18_[int(15)] = d_75_v8_
        nw18_[int(16)] = d_75_v8_
        nw18_[int(17)] = d_75_v8_
        nw18_[int(18)] = _dafny.CodePoint('e')
        nw18_[int(19)] = d_75_v8_
        nw18_[int(20)] = d_75_v8_
        nw18_[int(21)] = d_75_v8_
        nw18_[int(22)] = d_75_v8_
        nw18_[int(23)] = d_75_v8_
        nw18_[int(24)] = d_75_v8_
        nw18_[int(25)] = d_75_v8_
        nw18_[int(26)] = d_75_v8_
        d_76_v9_ = nw18_
        d_77_v10_: _dafny.Seq
        d_77_v10_ = _dafny.SeqWithoutIsStrInference([d_76_v9_, d_76_v9_, d_76_v9_, d_76_v9_, d_76_v9_])
        d_78_v11_: _dafny.MultiSet
        d_78_v11_ = _dafny.MultiSet([d_73_v6_])
        d_79_v12_: _dafny.Seq
        d_79_v12_ = _dafny.SeqWithoutIsStrInference([d_78_v11_])
        d_80_v13_: _dafny.Array
        def lambda1_(d_81_v1_):
            def lambda2_(d_82_i1_):
                return _dafny.Map({d_81_v1_: -176})

            return lambda2_

        init1_ = lambda1_(d_69_v1_)
        nw19_ = _dafny.Array(None, 29)
        for i0_1_ in range(nw19_.length(0)):
            nw19_[i0_1_] = init1_(i0_1_)
        d_80_v13_ = nw19_
        d_83_v14_: _dafny.Seq
        d_83_v14_ = _dafny.SeqWithoutIsStrInference([d_71_v3_, d_71_v3_])
        d_84_globalState_: GlobalState
        nw20_ = GlobalState()
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (d_74_v7_).Elements:
                d_85_v5_: int = compr_0_
                if (d_85_v5_) in (d_74_v7_):
                    coll0_[(d_85_v5_) * (d_73_v6_)] = _dafny.CodePoint('j')
            return _dafny.Map(coll0_)
        nw20_.ctor__(d_68_v0_, (d_70_v2_).intersection(d_70_v2_), d_71_v3_, False, (d_72_v4_).cf0, 202, (iife0_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_73_v6_ for d_86_i0_ in range(default__.abs(668))])): d_75_v8_})), False, (d_77_v10_) + (d_77_v10_), False, 525, False, False, True, (d_79_v12_)[default__.safeIndex(d_73_v6_, len(d_79_v12_))], d_80_v13_, 588, False, (d_83_v14_)[default__.safeIndex(d_73_v6_, len(d_83_v14_))], d_68_v0_, True, True, 173, True)
        d_84_globalState_ = nw20_
        source0_ = d_72_v4_
        if source0_.is_DC1:
            d_87___mcc_h0_ = source0_.cf1
            d_88___mcc_h1_ = source0_.cf2
            d_89___mcc_h2_ = source0_.cf3
            d_90_cf3_ = d_89___mcc_h2_
            d_91_cf2_ = d_88___mcc_h1_
            d_92_cf1_ = d_87___mcc_h0_
            d_93_v16_: _dafny.MultiSet
            d_93_v16_ = _dafny.MultiSet([d_68_v0_, d_68_v0_])
            d_94_v17_: D0
            d_94_v17_ = D0_DC1(d_69_v1_, d_69_v1_, d_90_cf3_)
            pat_let_tv0_ = d_91_cf2_
            pat_let_tv1_ = d_73_v6_
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(-99, 854):
                    d_95_v15_: int = compr_1_
                    if ((-99) <= (d_95_v15_)) and ((d_95_v15_) < (854)):
                        coll1_[(d_95_v15_) * (d_73_v6_)] = len(_dafny.SeqWithoutIsStrInference([d_75_v8_ for d_96_i2_ in range(default__.abs(-372))]))
                return _dafny.Map(coll1_)
            def iife3_(_pat_let1_0):
                def iife4_(d_97_dt__update__tmp_h0_):
                    def iife5_(_pat_let2_0):
                        def iife6_(d_98_dt__update_hcf2_h0_):
                            return D0_DC1((d_97_dt__update__tmp_h0_).cf1, d_98_dt__update_hcf2_h0_, (d_97_dt__update__tmp_h0_).cf3)
                        return iife6_(_pat_let2_0)
                    return iife5_(pat_let_tv0_)
                return iife4_(_pat_let1_0)
            def iife2_(_pat_let0_0):
                def iife7_(d_99_dt__update__tmp_h1_):
                    def iife8_(_pat_let3_0):
                        def iife9_(d_100_dt__update_hcf3_h0_):
                            def iife10_(_pat_let4_0):
                                def iife11_(d_101_dt__update_hcf2_h1_):
                                    return D0_DC1((d_99_dt__update__tmp_h1_).cf1, d_101_dt__update_hcf2_h1_, d_100_dt__update_hcf3_h0_)
                                return iife11_(_pat_let4_0)
                            return iife10_(False)
                        return iife9_(_pat_let3_0)
                    return iife8_((0) - (pat_let_tv1_))
                return iife7_(_pat_let0_0)
            rhs10_ = (len(iife1_()
            )) - ((d_93_v16_).cardinality)
            rhs11_ = (default__.safeModuloInt(d_73_v6_, d_73_v6_) if d_91_cf2_ else 620)
            rhs12_ = default__.safeModuloInt((d_73_v6_ if d_92_cf1_ else (d_74_v7_)[default__.safeIndex(d_73_v6_, len(d_74_v7_))]), d_90_cf3_)
            rhs13_ = (iife2_(iife3_(d_94_v17_))).cf3
            rhs14_ = (d_71_v3_ if not(d_69_v1_) else d_71_v3_)
            lhs6_ = d_84_globalState_
            lhs7_ = d_84_globalState_
            lhs6_.f5 = rhs10_
            d_90_cf3_ = rhs11_
            lhs7_.f5 = rhs12_
            d_90_cf3_ = rhs13_
            d_71_v3_ = rhs14_
            (d_84_globalState_).f5 = d_73_v6_
            d_102_v18_: _dafny.Map
            d_102_v18_ = _dafny.Map({d_91_cf2_: -284})
            d_103_v19_: C0
            nw21_ = C0()
            nw21_.ctor__(len((_dafny.Map({d_69_v1_: -22})) | (d_102_v18_)), d_91_cf2_, d_68_v0_)
            d_103_v19_ = nw21_
            d_104_v20_: T0
            nw22_ = C0()
            nw22_.ctor__(d_73_v6_, d_91_cf2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlwarcww")))
            d_104_v20_ = nw22_
            d_105_v21_: _dafny.Map
            d_105_v21_ = _dafny.Map({d_104_v20_: d_90_cf3_})
            (d_84_globalState_).f5 = ((d_105_v21_)[d_104_v20_] if (d_104_v20_) in (d_105_v21_) else d_90_cf3_)
        elif source0_.is_DC2:
            d_106___mcc_h3_ = source0_.cf4
            d_107___mcc_h4_ = source0_.cf5
            d_108_cf5_ = d_107___mcc_h4_
            d_109_cf4_ = d_106___mcc_h3_
            d_110_v22_: C0
            nw23_ = C0()
            nw23_.ctor__(len(d_74_v7_), d_69_v1_, _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_111_i3_ in range(default__.abs(284))]))
            d_110_v22_ = nw23_
            d_112_v23_: _dafny.Map
            d_112_v23_ = _dafny.Map({d_110_v22_: d_73_v6_})
            d_113_v24_: C0
            nw24_ = C0()
            nw24_.ctor__(len(d_112_v23_), d_69_v1_, d_68_v0_)
            d_113_v24_ = nw24_
            d_114_v25_: D1
            d_114_v25_ = D1_DC4(d_110_v22_)
            d_115_v26_: _dafny.Map
            d_115_v26_ = _dafny.Map({(d_114_v25_).cf7: d_110_v22_})
            d_116_v27_: _dafny.Seq
            d_116_v27_ = _dafny.SeqWithoutIsStrInference([d_115_v26_, d_115_v26_, (d_115_v26_) | (d_115_v26_), (d_115_v26_) | (d_115_v26_), d_115_v26_])
            d_116_v27_ = d_116_v27_
            (d_84_globalState_).f20 = d_69_v1_
            index8_ = default__.safeIndex(895, (d_71_v3_).length(0))
            (d_71_v3_)[index8_] = d_69_v1_
            d_117_v28_: _dafny.Seq
            d_117_v28_ = _dafny.SeqWithoutIsStrInference([d_69_v1_, d_69_v1_])
            index9_ = default__.safeIndex(895, (d_71_v3_).length(0))
            (d_71_v3_)[index9_] = (d_117_v28_)[default__.safeIndex(d_73_v6_, len(d_117_v28_))]
            d_118_v29_: _dafny.Set
            d_118_v29_ = _dafny.Set({d_73_v6_, (d_110_v22_).f26})
            d_119_v30_: _dafny.Map
            d_119_v30_ = _dafny.Map({d_108_cf5_: not((d_71_v3_)[default__.safeIndex(895, (d_71_v3_).length(0))])})
            rhs15_ = default__.fm4(len(d_118_v29_), d_73_v6_, _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_120_i4_ in range(default__.abs(102))]), (_dafny.SeqWithoutIsStrInference([d_69_v1_, d_69_v1_])) + (d_117_v28_), d_84_globalState_)
            rhs16_ = ((d_113_v24_).f26) * (len(d_119_v30_))
            lhs8_ = d_84_globalState_
            d_75_v8_ = rhs15_
            lhs8_.f5 = rhs16_
        elif source0_.is_DC0:
            d_121___mcc_h5_ = source0_.cf0
            d_122_cf0_ = d_121___mcc_h5_
            (d_84_globalState_).f5 = len(d_68_v0_)
            d_123_v31_: C0
            nw25_ = C0()
            nw25_.ctor__(d_73_v6_, d_69_v1_, d_122_cf0_)
            d_123_v31_ = nw25_
            d_124_v32_: _dafny.Seq
            d_124_v32_ = _dafny.SeqWithoutIsStrInference([d_123_v31_])
            d_125_v33_: _dafny.Seq
            d_125_v33_ = _dafny.SeqWithoutIsStrInference([d_124_v32_])
            d_125_v33_ = _dafny.SeqWithoutIsStrInference([d_124_v32_, (d_124_v32_).set(default__.safeIndex(477, len(d_124_v32_)), d_123_v31_), d_124_v32_])
            d_71_v3_ = (d_71_v3_ if not(d_69_v1_) else d_71_v3_)
            d_126_v34_: _dafny.Seq
            d_127_v35_: _dafny.Array
            d_128_v36_: int
            d_129_v37_: bool
            out0_: _dafny.Seq
            out1_: _dafny.Array
            out2_: int
            out3_: bool
            out0_, out1_, out2_, out3_ = default__.m0(d_84_globalState_)
            d_126_v34_ = out0_
            d_127_v35_ = out1_
            d_128_v36_ = out2_
            d_129_v37_ = out3_
        elif True:
            d_130___mcc_h6_ = source0_.cf6
            d_131_cf6_ = d_130___mcc_h6_
            source1_ = D0_DC0((d_68_v0_ if True else d_68_v0_))
            if source1_.is_DC1:
                d_132___mcc_h7_ = source1_.cf1
                d_133___mcc_h8_ = source1_.cf2
                d_134___mcc_h9_ = source1_.cf3
                d_135_cf3_ = d_134___mcc_h9_
                d_136_cf2_ = d_133___mcc_h8_
                d_137_cf1_ = d_132___mcc_h7_
                d_138_v38_: _dafny.Seq
                d_138_v38_ = _dafny.SeqWithoutIsStrInference([not(d_136_cf2_)])
                d_68_v0_ = ((d_68_v0_) + (d_68_v0_)).set(default__.safeIndex(default__.safeModuloInt(d_73_v6_, (0) - (d_73_v6_)), len((d_68_v0_) + (d_68_v0_))), default__.fm4(len(d_68_v0_), d_135_cf3_, _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_139_i5_ in range(default__.abs(911))]), d_138_v38_, d_84_globalState_))
                d_136_cf2_ = default__.fm5(not (d_69_v1_) or (d_137_cf1_), d_137_cf1_, d_135_cf3_, d_73_v6_, d_84_globalState_)
                d_136_cf2_ = d_69_v1_
                d_140_v39_: T1
                nw26_ = C0()
                nw26_.ctor__(d_73_v6_, d_137_cf1_, d_68_v0_)
                d_140_v39_ = nw26_
                d_141_v40_: _dafny.Map
                d_141_v40_ = _dafny.Map({d_140_v39_: default__.fm5(True, d_136_cf2_, d_135_cf3_, d_135_cf3_, d_84_globalState_)})
                (d_84_globalState_).f5 = default__.safeModuloInt(default__.fm6(not(d_69_v1_), d_136_cf2_, d_136_cf2_, d_84_globalState_), len((d_141_v40_).set(d_140_v39_, False)))
            elif source1_.is_DC2:
                d_142___mcc_h10_ = source1_.cf4
                d_143___mcc_h11_ = source1_.cf5
                d_144_cf5_ = d_143___mcc_h11_
                d_145_cf4_ = d_142___mcc_h10_
                d_146_v41_: _dafny.MultiSet
                d_146_v41_ = _dafny.MultiSet([d_69_v1_, False])
                d_147_v42_: T1
                nw27_ = C0()
                nw27_.ctor__((d_146_v41_).cardinality, d_69_v1_, d_68_v0_)
                d_147_v42_ = nw27_
                d_148_v43_: D1
                d_148_v43_ = D1_DC6(d_144_cf5_, d_147_v42_, not(True), d_69_v1_)
                d_149_v44_: _dafny.Map
                d_149_v44_ = _dafny.Map({False: d_75_v8_})
                d_150_v45_: _dafny.Map
                d_150_v45_ = _dafny.Map({d_73_v6_: default__.fm5(d_69_v1_, (d_147_v42_).f24, 609, d_144_cf5_, d_84_globalState_)})
                d_151_v46_: _dafny.Map
                d_151_v46_ = _dafny.Map({d_144_cf5_: d_144_cf5_})
                d_152_v48_: C0
                nw28_ = C0()
                nw28_.ctor__(d_144_cf5_, not(not(d_69_v1_)), (d_147_v42_).f25)
                d_152_v48_ = nw28_
                d_153_v49_: _dafny.Seq
                d_153_v49_ = _dafny.SeqWithoutIsStrInference([d_152_v48_, d_152_v48_, d_152_v48_])
                d_154_v50_: D0
                d_154_v50_ = D0_DC2(d_145_cf4_, d_73_v6_)
                d_155_v51_: _dafny.MultiSet
                def iife12_():
                    coll2_ = _dafny.Map()
                    compr_2_: int
                    for compr_2_ in (_dafny.Set({len(d_153_v49_), len(_dafny.SeqWithoutIsStrInference([(d_154_v50_).cf5])), (d_152_v48_).f26, (d_148_v43_).cf10, (d_152_v48_).f26})).Elements:
                        d_156_v47_: int = compr_2_
                        if (d_156_v47_) in (_dafny.Set({len(d_153_v49_), len(_dafny.SeqWithoutIsStrInference([(d_154_v50_).cf5])), (d_152_v48_).f26, (d_148_v43_).cf10, (d_152_v48_).f26})):
                            coll2_[default__.safeDivisionInt(d_156_v47_, d_144_cf5_)] = d_73_v6_
                    return _dafny.Map(coll2_)
                d_155_v51_ = _dafny.MultiSet([d_151_v46_, _dafny.Map({d_73_v6_: len(d_74_v7_)}), iife12_()
                , d_151_v46_])
                d_157_v52_: _dafny.Map
                d_157_v52_ = _dafny.Map({d_69_v1_: (d_155_v51_).cardinality})
                pat_let_tv2_ = d_147_v42_
                pat_let_tv3_ = d_69_v1_
                pat_let_tv4_ = d_147_v42_
                pat_let_tv5_ = d_144_cf5_
                def iife14_(_pat_let6_0):
                    def iife15_(d_158_dt__update__tmp_h2_):
                        def iife16_(_pat_let7_0):
                            def iife17_(d_159_dt__update_hcf11_h0_):
                                def iife18_(_pat_let8_0):
                                    def iife19_(d_160_dt__update_hcf10_h0_):
                                        return D1_DC6(d_160_dt__update_hcf10_h0_, d_159_dt__update_hcf11_h0_, (d_158_dt__update__tmp_h2_).cf12, (d_158_dt__update__tmp_h2_).cf13)
                                    return iife19_(_pat_let8_0)
                                return iife18_(660)
                            return iife17_(_pat_let7_0)
                        return iife16_(pat_let_tv2_)
                    return iife15_(_pat_let6_0)
                def iife13_(_pat_let5_0):
                    def iife20_(d_161_dt__update__tmp_h3_):
                        def iife21_(_pat_let9_0):
                            def iife22_(d_162_dt__update_hcf13_h0_):
                                def iife23_(_pat_let10_0):
                                    def iife24_(d_163_dt__update_hcf11_h1_):
                                        def iife25_(_pat_let11_0):
                                            def iife26_(d_164_dt__update_hcf10_h1_):
                                                return D1_DC6(d_164_dt__update_hcf10_h1_, d_163_dt__update_hcf11_h1_, (d_161_dt__update__tmp_h3_).cf12, d_162_dt__update_hcf13_h0_)
                                            return iife26_(_pat_let11_0)
                                        return iife25_(pat_let_tv5_)
                                    return iife24_(_pat_let10_0)
                                return iife23_(pat_let_tv4_)
                            return iife22_(_pat_let9_0)
                        return iife21_(not(pat_let_tv3_))
                    return iife20_(_pat_let5_0)
                (d_84_globalState_).f20 = ((d_147_v42_).f24 if default__.fm5((iife13_(iife14_(d_148_v43_))).cf12, d_69_v1_, default__.fm6((d_147_v42_).fm1(d_149_v44_, d_150_v45_, d_74_v7_, d_84_globalState_), d_69_v1_, d_69_v1_, d_84_globalState_), d_144_cf5_, d_84_globalState_) else (d_157_v52_) != (d_157_v52_))
                d_165_v53_: _dafny.Set
                d_165_v53_ = _dafny.Set({d_144_cf5_})
                d_166_v54_: _dafny.Seq
                d_166_v54_ = _dafny.SeqWithoutIsStrInference([(d_147_v42_).f24, (d_147_v42_).f24])
                rhs17_ = (d_165_v53_ if not(d_69_v1_) else (_dafny.Set({(d_152_v48_).f26})) - (d_165_v53_))
                rhs18_ = False
                rhs19_ = (d_166_v54_) == ((d_166_v54_) + (d_166_v54_))
                rhs20_ = (d_147_v42_).f24
                lhs9_ = d_84_globalState_
                lhs10_ = d_84_globalState_
                lhs11_ = d_84_globalState_
                d_165_v53_ = rhs17_
                lhs9_.f20 = rhs18_
                lhs10_.f21 = rhs19_
                lhs11_.f21 = rhs20_
                d_167_v55_: _dafny.Map
                d_167_v55_ = _dafny.Map({d_165_v53_: (d_152_v48_).f26})
                d_168_v56_: C0
                nw29_ = C0()
                nw29_.ctor__((d_152_v48_).f26, (d_167_v55_) == (d_167_v55_), d_68_v0_)
                d_168_v56_ = nw29_
                d_169_v57_: _dafny.Seq
                d_170_v58_: _dafny.Array
                d_171_v59_: int
                d_172_v60_: bool
                out4_: _dafny.Seq
                out5_: _dafny.Array
                out6_: int
                out7_: bool
                out4_, out5_, out6_, out7_ = default__.m0(d_84_globalState_)
                d_169_v57_ = out4_
                d_170_v58_ = out5_
                d_171_v59_ = out6_
                d_172_v60_ = out7_
            elif source1_.is_DC0:
                d_173___mcc_h12_ = source1_.cf0
                d_174_cf0_ = d_173___mcc_h12_
                (d_84_globalState_).f20 = d_69_v1_
                d_175_v61_: _dafny.Array
                nw30_ = _dafny.Array(int(0), 10)
                d_175_v61_ = nw30_
                index10_ = default__.safeIndex(299, (d_175_v61_).length(0))
                (d_175_v61_)[index10_] = d_73_v6_
                d_176_v62_: C0
                nw31_ = C0()
                nw31_.ctor__(d_73_v6_, False, d_174_cf0_)
                d_176_v62_ = nw31_
                d_177_v63_: _dafny.Seq
                d_177_v63_ = _dafny.SeqWithoutIsStrInference([d_176_v62_, d_176_v62_])
                d_178_v64_: _dafny.Seq
                d_178_v64_ = _dafny.SeqWithoutIsStrInference([d_177_v63_])
                index11_ = default__.safeIndex(299, (d_175_v61_).length(0))
                (d_175_v61_)[index11_] = (d_73_v6_) + ((len(((d_178_v64_)[default__.safeIndex(-355, len(d_178_v64_))]).set(default__.safeIndex((d_176_v62_).fm3(d_73_v6_, _dafny.SeqWithoutIsStrInference([d_74_v7_ for d_179_i6_ in range(default__.abs(609))]), d_73_v6_, (d_174_cf0_).set(default__.safeIndex(d_73_v6_, len(d_174_cf0_)), d_75_v8_), d_84_globalState_), len((d_178_v64_)[default__.safeIndex(-355, len(d_178_v64_))])), d_176_v62_)) if False else (0) - (len(d_174_cf0_))))
                d_69_v1_ = (659) > (len(d_174_cf0_))
                (d_84_globalState_).f21 = d_69_v1_
            elif True:
                d_180___mcc_h13_ = source1_.cf6
                d_181_cf6_ = d_180___mcc_h13_
                d_182_v65_: _dafny.Seq
                d_183_v66_: _dafny.Array
                d_184_v67_: int
                d_185_v68_: bool
                out8_: _dafny.Seq
                out9_: _dafny.Array
                out10_: int
                out11_: bool
                out8_, out9_, out10_, out11_ = default__.m0(d_84_globalState_)
                d_182_v65_ = out8_
                d_183_v66_ = out9_
                d_184_v67_ = out10_
                d_185_v68_ = out11_
                (d_84_globalState_).f5 = d_184_v67_
                d_69_v1_ = not(d_69_v1_)
                d_73_v6_ = default__.safeModuloInt(d_184_v67_, 874)
            d_69_v1_ = False
            d_186_v69_: _dafny.Seq
            d_186_v69_ = _dafny.SeqWithoutIsStrInference([d_69_v1_])
            d_187_v71_: C0
            nw32_ = C0()
            nw32_.ctor__(953, d_69_v1_, _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_188_i7_ in range(default__.abs(-753))]))
            d_187_v71_ = nw32_
            d_189_v72_: _dafny.Seq
            d_189_v72_ = _dafny.SeqWithoutIsStrInference([d_187_v71_, d_187_v71_, d_187_v71_, d_187_v71_])
            d_190_v73_: _dafny.Map
            d_190_v73_ = _dafny.Map({(d_187_v71_).f26: d_73_v6_})
            d_191_v74_: _dafny.Array
            nw33_ = _dafny.Array(None, 25)
            nw33_[int(0)] = d_73_v6_
            nw33_[int(1)] = d_73_v6_
            nw33_[int(2)] = d_73_v6_
            nw33_[int(3)] = d_73_v6_
            nw33_[int(4)] = len(d_186_v69_)
            nw33_[int(5)] = d_73_v6_
            nw33_[int(6)] = 225
            nw33_[int(7)] = len(d_68_v0_)
            nw33_[int(8)] = 373
            def iife27_():
                coll3_ = _dafny.Set()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(970, 840):
                    d_192_v70_: int = compr_3_
                    if ((970) <= (d_192_v70_)) and ((d_192_v70_) < (840)):
                        coll3_ = coll3_.union(_dafny.Set([(d_192_v70_) * (d_73_v6_)]))
                return _dafny.Set(coll3_)
            nw33_[int(9)] = len(iife27_()
            )
            nw33_[int(10)] = d_73_v6_
            nw33_[int(11)] = (d_78_v11_).cardinality
            nw33_[int(12)] = len(d_68_v0_)
            nw33_[int(13)] = d_73_v6_
            nw33_[int(14)] = len(d_189_v72_)
            nw33_[int(15)] = (d_187_v71_).f26
            nw33_[int(16)] = d_73_v6_
            nw33_[int(17)] = 768
            nw33_[int(18)] = d_73_v6_
            nw33_[int(19)] = (d_187_v71_).f26
            nw33_[int(20)] = d_73_v6_
            nw33_[int(21)] = (d_187_v71_).f26
            nw33_[int(22)] = len(d_190_v73_)
            nw33_[int(23)] = -224
            nw33_[int(24)] = d_73_v6_
            d_191_v74_ = nw33_
            d_193_v75_: D0
            d_193_v75_ = D0_DC2(d_191_v74_, d_73_v6_)
            source2_ = d_193_v75_
            if source2_.is_DC1:
                d_194___mcc_h14_ = source2_.cf1
                d_195___mcc_h15_ = source2_.cf2
                d_196___mcc_h16_ = source2_.cf3
                d_197_cf3_ = d_196___mcc_h16_
                d_198_cf2_ = d_195___mcc_h15_
                d_199_cf1_ = d_194___mcc_h14_
                d_200_v76_: _dafny.Map
                d_200_v76_ = _dafny.Map({((d_187_v71_).f26) * ((d_74_v7_)[default__.safeIndex(d_197_cf3_, len(d_74_v7_))]): d_71_v3_})
                d_201_v77_: _dafny.Map
                d_201_v77_ = _dafny.Map({d_199_cf1_: -730})
                d_200_v76_ = (d_200_v76_).set(len(d_201_v77_), d_71_v3_)
                d_202_v78_: _dafny.Map
                d_202_v78_ = _dafny.Map({((d_187_v71_).f26) - (d_197_cf3_): d_198_cf2_})
                d_202_v78_ = (d_202_v78_).set((d_187_v71_).f26, not(d_199_cf1_))
                d_198_cf2_ = d_199_cf1_
                d_203_v79_: _dafny.Seq
                d_204_v80_: _dafny.Array
                d_205_v81_: int
                d_206_v82_: bool
                out12_: _dafny.Seq
                out13_: _dafny.Array
                out14_: int
                out15_: bool
                out12_, out13_, out14_, out15_ = default__.m0(d_84_globalState_)
                d_203_v79_ = out12_
                d_204_v80_ = out13_
                d_205_v81_ = out14_
                d_206_v82_ = out15_
            elif source2_.is_DC2:
                d_207___mcc_h17_ = source2_.cf4
                d_208___mcc_h18_ = source2_.cf5
                d_209_cf5_ = d_208___mcc_h18_
                d_210_cf4_ = d_207___mcc_h17_
                d_211_v83_: _dafny.Map
                d_211_v83_ = _dafny.Map({d_70_v2_: len(d_70_v2_)})
                d_212_v84_: _dafny.Map
                d_212_v84_ = _dafny.Map({d_69_v1_: d_73_v6_})
                d_213_v85_: _dafny.MultiSet
                d_213_v85_ = _dafny.MultiSet([d_212_v84_])
                d_214_v86_: _dafny.Map
                d_214_v86_ = _dafny.Map({((d_211_v83_)[d_70_v2_] if (d_70_v2_) in (d_211_v83_) else len(d_74_v7_)): (d_213_v85_).issubset((d_213_v85_).set(d_212_v84_, default__.abs((d_187_v71_).f26)))})
                d_214_v86_ = (d_214_v86_).set((d_187_v71_).f26, False)
                d_215_v87_: C0
                nw34_ = C0()
                nw34_.ctor__((0) - ((d_187_v71_).f26), d_69_v1_, d_68_v0_)
                d_215_v87_ = nw34_
                index12_ = default__.safeIndex(208, (d_76_v9_).length(0))
                (d_76_v9_)[index12_] = d_75_v8_
                index13_ = default__.safeIndex(208, (d_76_v9_).length(0))
                (d_76_v9_)[index13_] = d_75_v8_
                d_70_v2_ = d_70_v2_
            elif source2_.is_DC0:
                d_216___mcc_h19_ = source2_.cf0
                d_217_cf0_ = d_216___mcc_h19_
                d_218_v88_: _dafny.Seq
                d_219_v89_: _dafny.Array
                d_220_v90_: int
                d_221_v91_: bool
                out16_: _dafny.Seq
                out17_: _dafny.Array
                out18_: int
                out19_: bool
                out16_, out17_, out18_, out19_ = default__.m0(d_84_globalState_)
                d_218_v88_ = out16_
                d_219_v89_ = out17_
                d_220_v90_ = out18_
                d_221_v91_ = out19_
                d_222_v92_: C0
                nw35_ = C0()
                nw35_.ctor__(d_220_v90_, d_69_v1_, default__.fm7(d_69_v1_, d_220_v90_, (d_74_v7_)[default__.safeIndex(d_220_v90_, len(d_74_v7_))], d_221_v91_, d_84_globalState_))
                d_222_v92_ = nw35_
                d_223_v93_: _dafny.Map
                d_223_v93_ = _dafny.Map({d_220_v90_: d_221_v91_})
                d_223_v93_ = (d_223_v93_).set(((d_187_v71_).f26) + (d_73_v6_), not((d_73_v6_) > (len(d_186_v69_))))
                (d_84_globalState_).f20 = d_221_v91_
            elif True:
                d_224___mcc_h20_ = source2_.cf6
                d_225_cf6_ = d_224___mcc_h20_
                d_226_v94_: _dafny.Array
                nw36_ = _dafny.Array(D1.default()(), 13)
                d_226_v94_ = nw36_
                d_227_v95_: T1
                nw37_ = C0()
                nw37_.ctor__((d_187_v71_).f26, not(d_69_v1_), d_68_v0_)
                d_227_v95_ = nw37_
                index14_ = default__.safeIndex(301, (d_226_v94_).length(0))
                (d_226_v94_)[index14_] = D1_DC6((d_187_v71_).f26, d_227_v95_, d_69_v1_, (d_227_v95_).f24)
                d_228_v96_: _dafny.Map
                d_228_v96_ = _dafny.Map({(d_187_v71_).f26: True})
                index15_ = default__.safeIndex(301, (d_226_v94_).length(0))
                (d_226_v94_)[index15_] = D1_DC6((d_187_v71_).f26, d_227_v95_, (d_227_v95_).f24, (d_227_v95_).fm1(_dafny.Map({(d_227_v95_).f24: d_75_v8_}), d_228_v96_, d_74_v7_, d_84_globalState_))
                d_229_v98_: _dafny.Map
                def iife28_():
                    coll4_ = _dafny.Map()
                    compr_4_: int
                    for compr_4_ in _dafny.IntegerRange(-822, 224):
                        d_230_v97_: int = compr_4_
                        if ((-822) <= (d_230_v97_)) and ((d_230_v97_) < (224)):
                            coll4_[default__.safeDivisionInt(d_230_v97_, (d_187_v71_).f26)] = (d_227_v95_).f24
                    return _dafny.Map(coll4_)
                d_229_v98_ = _dafny.Map({iife28_()
                : _dafny.MultiSet([(d_227_v95_).f24, d_69_v1_, d_69_v1_])})
                d_231_v99_: _dafny.MultiSet
                d_231_v99_ = _dafny.MultiSet([(d_68_v0_).set(default__.safeIndex(len(d_229_v98_), len(d_68_v0_)), _dafny.CodePoint('k'))])
                d_231_v99_ = (d_231_v99_).intersection(d_231_v99_)
                d_69_v1_ = (551) <= (((d_187_v71_).f26) * (564))
                d_232_v100_: _dafny.Array
                nw38_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_232_v100_ = nw38_
                d_233_v101_: D2
                d_233_v101_ = D2_DC8(d_71_v3_)
                d_234_v102_: _dafny.Array
                nw39_ = _dafny.Array(None, 14)
                nw39_[int(0)] = d_71_v3_
                nw39_[int(1)] = d_71_v3_
                nw39_[int(2)] = d_71_v3_
                nw39_[int(3)] = d_71_v3_
                nw39_[int(4)] = d_71_v3_
                nw39_[int(5)] = (d_233_v101_).cf15
                nw39_[int(6)] = d_71_v3_
                nw39_[int(7)] = d_71_v3_
                nw39_[int(8)] = d_71_v3_
                nw39_[int(9)] = d_71_v3_
                nw39_[int(10)] = d_71_v3_
                nw39_[int(11)] = d_71_v3_
                nw39_[int(12)] = d_71_v3_
                nw39_[int(13)] = d_71_v3_
                d_234_v102_ = nw39_
                index16_ = default__.safeIndex(397, (d_232_v100_).length(0))
                (d_232_v100_)[index16_] = d_234_v102_
                index17_ = default__.safeIndex(397, (d_232_v100_).length(0))
                (d_232_v100_)[index17_] = (d_234_v102_ if (d_227_v95_).f24 else d_234_v102_)
            d_235_v103_: _dafny.Map
            d_235_v103_ = _dafny.Map({d_69_v1_: d_75_v8_})
            d_236_v104_: _dafny.Map
            d_236_v104_ = _dafny.Map({(0) - (default__.fm6(d_69_v1_, d_69_v1_, d_69_v1_, d_84_globalState_)): d_186_v69_})
            (d_84_globalState_).f21 = ((d_69_v1_ if (d_187_v71_).fm1(d_235_v103_, _dafny.Map({len(d_236_v104_): d_69_v1_}), _dafny.SeqWithoutIsStrInference([d_73_v6_ for d_237_i8_ in range(default__.abs(277))]), d_84_globalState_) else not(False))) and (d_69_v1_)
        d_238_v105_: _dafny.Array
        nw40_ = _dafny.Array(_dafny.Map({}), 24)
        d_238_v105_ = nw40_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_238_v105_).length(0)):
            d_239_i9_: int = guard_loop_2_
            if (True) and (((0) <= (d_239_i9_)) and ((d_239_i9_) < ((d_238_v105_).length(0)))):
                (d_238_v105_)[(d_239_i9_)] = (_dafny.Map({d_69_v1_: d_75_v8_})) | ((_dafny.Map({d_69_v1_: _dafny.CodePoint('t')})) | (_dafny.Map({d_69_v1_: d_75_v8_})))
        d_240_v106_: _dafny.Seq
        d_240_v106_ = _dafny.SeqWithoutIsStrInference([d_69_v1_, d_69_v1_])
        (d_84_globalState_).f21 = (not((d_240_v106_)[default__.safeIndex(len(d_68_v0_), len(d_240_v106_))])) == (d_69_v1_)
        d_241_v107_: _dafny.Map
        d_241_v107_ = _dafny.Map({49: d_73_v6_})
        d_242_v108_: D1
        d_242_v108_ = D1_DC5(d_69_v1_, d_73_v6_)
        pat_let_tv6_ = d_73_v6_
        pat_let_tv7_ = d_73_v6_
        pat_let_tv8_ = d_73_v6_
        pat_let_tv9_ = d_73_v6_
        pat_let_tv10_ = d_73_v6_
        pat_let_tv11_ = d_73_v6_
        def lambda3_(source3_):
            if source3_.is_DC5:
                d_243___mcc_h21_ = source3_.cf8
                d_244___mcc_h22_ = source3_.cf9
                d_245_cf9_ = d_244___mcc_h22_
                d_246_cf8_ = d_243___mcc_h21_
                return (pat_let_tv6_) * (pat_let_tv7_)
            elif source3_.is_DC6:
                d_247___mcc_h23_ = source3_.cf10
                d_248___mcc_h24_ = source3_.cf11
                d_249___mcc_h25_ = source3_.cf12
                d_250___mcc_h26_ = source3_.cf13
                d_251_cf13_ = d_250___mcc_h26_
                d_252_cf12_ = d_249___mcc_h25_
                d_253_cf11_ = d_248___mcc_h24_
                d_254_cf10_ = d_247___mcc_h23_
                return (pat_let_tv8_) * (pat_let_tv9_)
            elif source3_.is_DC4:
                d_255___mcc_h27_ = source3_.cf7
                d_256_cf7_ = d_255___mcc_h27_
                return (d_256_cf7_).f26
            elif True:
                d_257___mcc_h28_ = source3_.cf14
                d_258_cf14_ = d_257___mcc_h28_
                def iife29_():
                    coll5_ = _dafny.Set()
                    compr_5_: int
                    for compr_5_ in (_dafny.Set({pat_let_tv10_})).Elements:
                        d_259_v109_: int = compr_5_
                        if (d_259_v109_) in (_dafny.Set({pat_let_tv11_})):
                            coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_259_v109_, 53)]))
                    return _dafny.Set(coll5_)
                return len(iife29_()
                )

        rhs21_ = (d_69_v1_) == (d_69_v1_)
        rhs22_ = ((d_241_v107_)[d_73_v6_] if (d_73_v6_) in (d_241_v107_) else d_73_v6_)
        rhs23_ = lambda3_(d_242_v108_)
        lhs12_ = d_84_globalState_
        lhs12_.f20 = rhs21_
        d_73_v6_ = rhs22_
        d_73_v6_ = rhs23_
        d_260_v110_: T1
        nw41_ = C0()
        nw41_.ctor__(d_73_v6_, d_69_v1_, d_68_v0_)
        d_260_v110_ = nw41_
        d_261_v111_: D1
        d_261_v111_ = D1_DC6(d_73_v6_, d_260_v110_, True, (d_260_v110_).f24)
        d_262_v112_: _dafny.Map
        d_262_v112_ = _dafny.Map({False: ((d_261_v111_).cf10) > (d_73_v6_)})
        d_262_v112_ = (d_262_v112_).set((d_260_v110_).f24, (d_73_v6_) != (d_73_v6_))
        hi0_ = d_73_v6_
        for d_263_i10_ in range(658, hi0_):
            d_73_v6_ = d_263_i10_
            d_264_v113_: _dafny.Array
            nw42_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_264_v113_ = nw42_
            d_264_v113_ = d_264_v113_
            d_265_v114_: _dafny.Map
            d_265_v114_ = _dafny.Map({d_73_v6_: d_260_v110_})
            d_265_v114_ = (d_265_v114_).set((0) - (len(d_74_v7_)), d_260_v110_)
            d_262_v112_ = (d_262_v112_).set(d_69_v1_, (d_240_v106_)[default__.safeIndex(d_73_v6_, len(d_240_v106_))])
        d_266_v115_: _dafny.Set
        d_266_v115_ = _dafny.Set({d_75_v8_})
        d_267_v116_: _dafny.Set
        d_267_v116_ = _dafny.Set({d_266_v115_})
        d_268_i11_: int
        d_268_i11_ = 0
        with _dafny.label("0"):
            while ((d_267_v116_) | (_dafny.Set({d_266_v115_, d_266_v115_}))).ispropersubset(d_267_v116_):
                with _dafny.c_label("0"):
                    if (d_268_i11_) >= (100):
                        raise _dafny.Break("0")
                    d_268_i11_ = (d_268_i11_) + (1)
                    d_269_v117_: _dafny.Array
                    nw43_ = _dafny.Array(None, 1)
                    nw43_[int(0)] = d_73_v6_
                    d_269_v117_ = nw43_
                    index18_ = default__.safeIndex(6, (d_269_v117_).length(0))
                    (d_269_v117_)[index18_] = d_73_v6_
                    d_270_v118_: _dafny.Map
                    d_270_v118_ = _dafny.Map({((d_260_v110_).f25) + (d_68_v0_): d_73_v6_})
                    index19_ = default__.safeIndex(6, (d_269_v117_).length(0))
                    rhs24_ = (d_240_v106_)[default__.safeIndex(default__.safeDivisionInt(d_73_v6_, d_73_v6_), len(d_240_v106_))]
                    rhs25_ = ((d_270_v118_)[(d_260_v110_).f25] if ((d_260_v110_).f25) in (d_270_v118_) else d_73_v6_)
                    rhs26_ = (674) == (default__.fm6(d_69_v1_, (d_260_v110_).f24, True, d_84_globalState_))
                    lhs13_ = d_84_globalState_
                    lhs14_ = d_269_v117_
                    lhs15_ = default__.safeIndex(6, (d_269_v117_).length(0))
                    lhs16_ = d_84_globalState_
                    lhs13_.f21 = rhs24_
                    lhs14_[lhs15_] = rhs25_
                    lhs16_.f21 = rhs26_
                    d_271_v119_: _dafny.Seq
                    d_272_v120_: _dafny.Array
                    d_273_v121_: int
                    d_274_v122_: bool
                    out20_: _dafny.Seq
                    out21_: _dafny.Array
                    out22_: int
                    out23_: bool
                    out20_, out21_, out22_, out23_ = default__.m0(d_84_globalState_)
                    d_271_v119_ = out20_
                    d_272_v120_ = out21_
                    d_273_v121_ = out22_
                    d_274_v122_ = out23_
                    d_69_v1_ = ((54) + (len(d_74_v7_))) == (d_273_v121_)
                    d_275_v123_: C0
                    nw44_ = C0()
                    nw44_.ctor__((0) - (d_73_v6_), d_274_v122_, d_271_v119_)
                    d_275_v123_ = nw44_
                    d_276_v124_: _dafny.Array
                    nw45_ = _dafny.Array(_dafny.Map({}), 19)
                    d_276_v124_ = nw45_
                    d_277_v125_: _dafny.Map
                    d_277_v125_ = _dafny.Map({d_78_v11_: not((d_260_v110_).f24)})
                    index20_ = default__.safeIndex(695, (d_276_v124_).length(0))
                    (d_276_v124_)[index20_] = d_277_v125_
                    index21_ = default__.safeIndex(695, (d_276_v124_).length(0))
                    def iife30_():
                        coll6_ = _dafny.Map()
                        compr_6_: _dafny.MultiSet
                        for compr_6_ in (d_79_v12_).Elements:
                            d_278_v126_: _dafny.MultiSet = compr_6_
                            if (d_278_v126_) in (d_79_v12_):
                                coll6_[d_278_v126_] = (d_260_v110_).f24
                        return _dafny.Map(coll6_)
                    rhs27_ = d_275_v123_
                    rhs28_ = (iife30_()
                    ) | ((d_277_v125_) | (d_277_v125_))
                    rhs29_ = d_71_v3_
                    rhs30_ = (d_269_v117_)[default__.safeIndex(6, (d_269_v117_).length(0))]
                    lhs17_ = d_276_v124_
                    lhs18_ = default__.safeIndex(695, (d_276_v124_).length(0))
                    d_275_v123_ = rhs27_
                    lhs17_[lhs18_] = rhs28_
                    d_71_v3_ = rhs29_
                    d_273_v121_ = rhs30_
                    pass
            pass
        d_279_v127_: _dafny.Array
        nw46_ = _dafny.Array(None, 16)
        nw46_[int(0)] = (d_260_v110_).f25
        nw46_[int(1)] = (d_68_v0_) + (((d_260_v110_).f25).set(default__.safeIndex(d_73_v6_, len((d_260_v110_).f25)), d_75_v8_))
        nw46_[int(2)] = (d_260_v110_).f25
        nw46_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhh"))) + (d_68_v0_)
        nw46_[int(4)] = (((d_260_v110_).f25) + (d_68_v0_)).set(default__.safeIndex(d_73_v6_, len(((d_260_v110_).f25) + (d_68_v0_))), d_75_v8_)
        nw46_[int(5)] = ((d_260_v110_).f25) + ((d_260_v110_).f25)
        nw46_[int(6)] = (d_260_v110_).f25
        nw46_[int(7)] = (d_260_v110_).f25
        nw46_[int(8)] = (d_260_v110_).f25
        nw46_[int(9)] = (d_260_v110_).f25
        nw46_[int(10)] = default__.fm7((d_260_v110_).f24, d_73_v6_, d_73_v6_, (d_260_v110_).f24, d_84_globalState_)
        nw46_[int(11)] = d_68_v0_
        nw46_[int(12)] = d_68_v0_
        nw46_[int(13)] = d_68_v0_
        nw46_[int(14)] = (d_260_v110_).f25
        nw46_[int(15)] = _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_280_i12_ in range(default__.abs(13))])
        d_279_v127_ = nw46_
        index22_ = default__.safeIndex(285, (d_279_v127_).length(0))
        (d_279_v127_)[index22_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chxreot"))).set(default__.safeIndex(d_73_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chxreot")))), d_75_v8_)
        index23_ = default__.safeIndex(285, (d_279_v127_).length(0))
        (d_279_v127_)[index23_] = (d_260_v110_).f25
        if d_69_v1_:
            d_71_v3_ = d_71_v3_
            d_281_v128_: _dafny.Seq
            d_282_v129_: _dafny.Array
            d_283_v130_: int
            d_284_v131_: bool
            out24_: _dafny.Seq
            out25_: _dafny.Array
            out26_: int
            out27_: bool
            out24_, out25_, out26_, out27_ = default__.m0(d_84_globalState_)
            d_281_v128_ = out24_
            d_282_v129_ = out25_
            d_283_v130_ = out26_
            d_284_v131_ = out27_
            (d_84_globalState_).f21 = not (not((d_73_v6_) > (991))) or (d_69_v1_)
            d_285_v132_: _dafny.Seq
            d_286_v133_: _dafny.Array
            d_287_v134_: int
            d_288_v135_: bool
            out28_: _dafny.Seq
            out29_: _dafny.Array
            out30_: int
            out31_: bool
            out28_, out29_, out30_, out31_ = default__.m0(d_84_globalState_)
            d_285_v132_ = out28_
            d_286_v133_ = out29_
            d_287_v134_ = out30_
            d_288_v135_ = out31_
            d_289_v136_: T0
            nw47_ = C0()
            nw47_.ctor__(len(d_285_v132_), True, (d_260_v110_).f25)
            d_289_v136_ = nw47_
            d_290_v137_: _dafny.Map
            d_290_v137_ = _dafny.Map({d_69_v1_: d_289_v136_})
            d_291_v138_: _dafny.MultiSet
            d_291_v138_ = _dafny.MultiSet([(d_290_v137_).set(d_284_v131_, d_289_v136_)])
            d_292_v139_: _dafny.Seq
            d_292_v139_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcoudebbi")), d_285_v132_])
            d_293_v140_: _dafny.Seq
            d_293_v140_ = _dafny.SeqWithoutIsStrInference([(d_292_v139_)[default__.safeIndex(d_287_v134_, len(d_292_v139_))], _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_294_i13_ in range(default__.abs(524))])])
            d_295_v141_: _dafny.Map
            d_295_v141_ = _dafny.Map({d_71_v3_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yckgtwu"))})
            d_296_v142_: _dafny.Array
            nw48_ = _dafny.Array(None, 17)
            nw48_[int(0)] = d_73_v6_
            nw48_[int(1)] = len((d_68_v0_) + ((d_260_v110_).f25))
            nw48_[int(2)] = (d_287_v134_) + (d_287_v134_)
            nw48_[int(3)] = (d_291_v138_).cardinality
            nw48_[int(4)] = d_287_v134_
            nw48_[int(5)] = len((d_260_v110_).f25)
            nw48_[int(6)] = (len((d_292_v139_)[default__.safeIndex(d_73_v6_, len(d_292_v139_))]) if d_69_v1_ else 995)
            nw48_[int(7)] = default__.safeDivisionInt(d_283_v130_, d_283_v130_)
            nw48_[int(8)] = (782) + (d_283_v130_)
            nw48_[int(9)] = (d_73_v6_) - (d_283_v130_)
            nw48_[int(10)] = 929
            nw48_[int(11)] = len((d_260_v110_).f25)
            nw48_[int(12)] = 507
            nw48_[int(13)] = default__.safeDivisionInt(len(d_295_v141_), d_287_v134_)
            nw48_[int(14)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_75_v8_ for d_297_i14_ in range(default__.abs(555))])), d_287_v134_)
            nw48_[int(15)] = default__.fm6((d_260_v110_).f24, False, d_288_v135_, d_84_globalState_)
            nw48_[int(16)] = len((d_74_v7_) + (_dafny.SeqWithoutIsStrInference([d_283_v130_])))
            d_296_v142_ = nw48_
            index24_ = default__.safeIndex(254, (d_296_v142_).length(0))
            (d_296_v142_)[index24_] = d_287_v134_
            index25_ = default__.safeIndex(254, (d_296_v142_).length(0))
            (d_296_v142_)[index25_] = (d_73_v6_) * (d_73_v6_)
        elif True:
            d_298_v143_: D2
            d_298_v143_ = D2_DC9(d_69_v1_, d_73_v6_)
            d_299_v144_: C0
            nw49_ = C0()
            nw49_.ctor__(len(_dafny.Map({d_72_v4_: (d_260_v110_).f24})), (d_260_v110_).f24, (d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])
            d_299_v144_ = nw49_
            d_300_v145_: _dafny.Set
            d_300_v145_ = _dafny.Set({d_299_v144_})
            d_301_v146_: _dafny.Map
            d_301_v146_ = _dafny.Map({d_69_v1_: d_300_v145_})
            d_302_v147_: _dafny.Map
            d_302_v147_ = _dafny.Map({(d_260_v110_).f25: _dafny.Map({d_298_v143_: len(d_301_v146_)})})
            d_303_v149_: _dafny.Map
            d_303_v149_ = _dafny.Map({d_68_v0_: _dafny.Map({(d_260_v110_).f24: 428})})
            def iife31_():
                coll7_ = _dafny.Map()
                compr_7_: _dafny.Seq
                for compr_7_ in (d_303_v149_).keys.Elements:
                    d_304_v148_: _dafny.Seq = compr_7_
                    if (d_304_v148_) in (d_303_v149_):
                        coll7_[d_304_v148_] = _dafny.Map({d_298_v143_: len((d_260_v110_).f25)})
                return _dafny.Map(coll7_)
            d_302_v147_ = iife31_()
            
            d_305_v150_: _dafny.Seq
            d_306_v151_: _dafny.Array
            d_307_v152_: int
            d_308_v153_: bool
            out32_: _dafny.Seq
            out33_: _dafny.Array
            out34_: int
            out35_: bool
            out32_, out33_, out34_, out35_ = default__.m0(d_84_globalState_)
            d_305_v150_ = out32_
            d_306_v151_ = out33_
            d_307_v152_ = out34_
            d_308_v153_ = out35_
            d_309_v154_: _dafny.Map
            d_309_v154_ = _dafny.Map({d_307_v152_: d_75_v8_})
            d_310_v155_: D0
            d_310_v155_ = D0_DC1(not(d_69_v1_), (d_260_v110_).f24, d_307_v152_)
            d_311_v156_: _dafny.Seq
            d_311_v156_ = _dafny.SeqWithoutIsStrInference([D1_DC6((d_310_v155_).cf3, d_260_v110_, True, (d_260_v110_).f24)])
            d_309_v154_ = (d_309_v154_).set(len(d_311_v156_), default__.fm4(len(d_309_v154_), (d_299_v144_).f26, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")), d_240_v106_, d_84_globalState_))
            if (d_260_v110_).f24:
                d_312_v157_: _dafny.Seq
                d_313_v158_: _dafny.Array
                d_314_v159_: int
                d_315_v160_: bool
                out36_: _dafny.Seq
                out37_: _dafny.Array
                out38_: int
                out39_: bool
                out36_, out37_, out38_, out39_ = default__.m0(d_84_globalState_)
                d_312_v157_ = out36_
                d_313_v158_ = out37_
                d_314_v159_ = out38_
                d_315_v160_ = out39_
                d_316_v161_: C0
                nw50_ = C0()
                nw50_.ctor__(d_314_v159_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))) < ((d_260_v110_).f25), (d_312_v157_) + (_dafny.SeqWithoutIsStrInference([d_75_v8_ for d_317_i15_ in range(default__.abs(728))])))
                d_316_v161_ = nw50_
                d_318_v162_: D0
                d_318_v162_ = D0_DC1(d_69_v1_, (d_240_v106_)[default__.safeIndex(d_307_v152_, len(d_240_v106_))], len(d_262_v112_))
                d_319_v163_: D0
                d_319_v163_ = D0_DC3(d_318_v162_)
                d_320_v164_: D0
                d_320_v164_ = D0_DC3(d_318_v162_)
                d_321_v165_: D0
                d_321_v165_ = D0_DC3(d_318_v162_)
                pat_let_tv12_ = d_319_v163_
                def iife32_(_pat_let12_0):
                    def iife33_(d_322_dt__update__tmp_h4_):
                        def iife34_(_pat_let13_0):
                            def iife35_(d_323_dt__update_hcf6_h0_):
                                return D0_DC3(d_323_dt__update_hcf6_h0_)
                            return iife35_(_pat_let13_0)
                        return iife34_(D0_DC3(pat_let_tv12_))
                    return iife33_(_pat_let12_0)
                rhs31_ = d_315_v160_
                rhs32_ = (d_260_v110_).f24
                rhs33_ = (d_307_v152_ if not(d_69_v1_) else d_314_v159_)
                rhs34_ = -766
                rhs35_ = iife32_(d_321_v165_)
                lhs19_ = d_84_globalState_
                lhs20_ = d_84_globalState_
                lhs21_ = d_84_globalState_
                d_315_v160_ = rhs31_
                lhs19_.f20 = rhs32_
                lhs20_.f5 = rhs33_
                lhs21_.f5 = rhs34_
                d_321_v165_ = rhs35_
                (d_84_globalState_).f20 = d_69_v1_
                d_324_v166_: D3
                d_324_v166_ = D3_DC11(d_74_v7_)
                def iife36_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(636, 445):
                        d_325_v167_: int = compr_8_
                        if ((636) <= (d_325_v167_)) and ((d_325_v167_) < (445)):
                            coll8_[(d_325_v167_) * ((d_316_v161_).f26)] = (d_260_v110_).f24
                    return _dafny.Map(coll8_)
                d_315_v160_ = (((d_324_v166_).cf19) + (d_74_v7_)) < (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.Map({(d_316_v161_).f26: (d_260_v110_).f24}), iife36_()
])).cardinality for d_326_i16_ in range(default__.abs(796))]))
            elif True:
                d_327_v168_: _dafny.Seq
                d_328_v169_: _dafny.Array
                d_329_v170_: int
                d_330_v171_: bool
                out40_: _dafny.Seq
                out41_: _dafny.Array
                out42_: int
                out43_: bool
                out40_, out41_, out42_, out43_ = default__.m0(d_84_globalState_)
                d_327_v168_ = out40_
                d_328_v169_ = out41_
                d_329_v170_ = out42_
                d_330_v171_ = out43_
                d_331_v172_: _dafny.Seq
                d_332_v173_: _dafny.Array
                d_333_v174_: int
                d_334_v175_: bool
                out44_: _dafny.Seq
                out45_: _dafny.Array
                out46_: int
                out47_: bool
                out44_, out45_, out46_, out47_ = default__.m0(d_84_globalState_)
                d_331_v172_ = out44_
                d_332_v173_ = out45_
                d_333_v174_ = out46_
                d_334_v175_ = out47_
                d_335_v176_: _dafny.Seq
                d_336_v177_: _dafny.Array
                d_337_v178_: int
                d_338_v179_: bool
                out48_: _dafny.Seq
                out49_: _dafny.Array
                out50_: int
                out51_: bool
                out48_, out49_, out50_, out51_ = default__.m0(d_84_globalState_)
                d_335_v176_ = out48_
                d_336_v177_ = out49_
                d_337_v178_ = out50_
                d_338_v179_ = out51_
                (d_84_globalState_).f20 = (d_298_v143_).cf16
                (d_84_globalState_).f21 = not((d_260_v110_).f24)
            d_339_v180_: T0
            nw51_ = C0()
            nw51_.ctor__(d_307_v152_, (d_261_v111_).cf13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwdsggla")))
            d_339_v180_ = nw51_
            d_339_v180_ = d_339_v180_
        d_340_v181_: _dafny.Map
        d_340_v181_ = _dafny.Map({d_75_v8_: d_73_v6_})
        d_341_v182_: _dafny.MultiSet
        d_341_v182_ = _dafny.MultiSet([(d_260_v110_).f24])
        d_342_v183_: _dafny.Seq
        d_342_v183_ = _dafny.SeqWithoutIsStrInference([d_341_v182_])
        d_343_v185_: _dafny.Array
        nw52_ = _dafny.Array(None, 18)
        nw52_[int(0)] = d_73_v6_
        nw52_[int(1)] = len(d_340_v181_)
        nw52_[int(2)] = 37
        nw52_[int(3)] = d_73_v6_
        nw52_[int(4)] = d_73_v6_
        nw52_[int(5)] = (((d_241_v107_)[len((d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])] if (len((d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])) in (d_241_v107_) else d_73_v6_)) * (d_73_v6_)
        nw52_[int(6)] = d_73_v6_
        nw52_[int(7)] = default__.safeDivisionInt(d_73_v6_, 655)
        nw52_[int(8)] = d_73_v6_
        nw52_[int(9)] = default__.safeModuloInt(d_73_v6_, d_73_v6_)
        nw52_[int(10)] = d_73_v6_
        nw52_[int(11)] = len((d_260_v110_).f25)
        nw52_[int(12)] = 201
        def iife37_():
            coll9_ = _dafny.Set()
            compr_9_: _dafny.MultiSet
            for compr_9_ in (d_342_v183_).Elements:
                d_344_v184_: _dafny.MultiSet = compr_9_
                if (d_344_v184_) in (d_342_v183_):
                    coll9_ = coll9_.union(_dafny.Set([d_344_v184_]))
            return _dafny.Set(coll9_)
        nw52_[int(13)] = (d_73_v6_ if (d_260_v110_).f24 else len(iife37_()
        ))
        nw52_[int(14)] = len(d_74_v7_)
        nw52_[int(15)] = len((d_74_v7_ if d_69_v1_ else d_74_v7_))
        nw52_[int(16)] = d_73_v6_
        nw52_[int(17)] = d_73_v6_
        d_343_v185_ = nw52_
        rhs36_ = d_73_v6_
        rhs37_ = (d_260_v110_).f24
        rhs38_ = (d_73_v6_) - (((d_341_v182_)[d_69_v1_] if (d_69_v1_) in (d_341_v182_) else d_73_v6_))
        rhs39_ = d_343_v185_
        lhs22_ = d_84_globalState_
        d_73_v6_ = rhs36_
        d_69_v1_ = rhs37_
        lhs22_.f5 = rhs38_
        d_343_v185_ = rhs39_
        d_345_i17_: int
        d_345_i17_ = 0
        with _dafny.label("1"):
            while d_69_v1_:
                with _dafny.c_label("1"):
                    if (d_345_i17_) >= (100):
                        raise _dafny.Break("1")
                    d_345_i17_ = (d_345_i17_) + (1)
                    (d_84_globalState_).f21 = not(((d_260_v110_).f24) or ((d_260_v110_).f24))
                    (d_84_globalState_).f5 = d_73_v6_
                    (d_84_globalState_).f21 = (d_260_v110_).f24
                    (d_84_globalState_).f5 = ((0) - (len(_dafny.SeqWithoutIsStrInference([d_73_v6_, d_73_v6_, 529, d_73_v6_, d_73_v6_])))) + (d_73_v6_)
                    pass
            pass
        d_346_v186_: C0
        nw53_ = C0()
        nw53_.ctor__(d_73_v6_, d_69_v1_, (d_260_v110_).f25)
        d_346_v186_ = nw53_
        d_347_v187_: _dafny.Seq
        d_347_v187_ = _dafny.SeqWithoutIsStrInference([d_346_v186_])
        d_73_v6_ = (default__.safeModuloInt(d_73_v6_, d_73_v6_)) - ((0) - (len((default__.fm7((d_260_v110_).f24, d_73_v6_, 25, d_69_v1_, d_84_globalState_)).set(default__.safeIndex(len(d_347_v187_), len(default__.fm7((d_260_v110_).f24, d_73_v6_, 25, d_69_v1_, d_84_globalState_))), d_75_v8_))))
        if d_69_v1_:
            (d_84_globalState_).f5 = (len((d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])) * (((_dafny.MultiSet([(d_74_v7_)[default__.safeIndex((d_346_v186_).f26, len(d_74_v7_))], (d_346_v186_).f26, (d_346_v186_).f26]) if d_69_v1_ else d_78_v11_)).cardinality)
            d_348_v188_: _dafny.Map
            d_348_v188_ = _dafny.Map({d_343_v185_: d_346_v186_})
            d_348_v188_ = (d_348_v188_) | (d_348_v188_)
            d_349_v189_: _dafny.Map
            d_349_v189_ = _dafny.Map({(d_260_v110_).f24: d_262_v112_})
            d_350_v190_: _dafny.Set
            d_350_v190_ = _dafny.Set({(d_346_v186_).f26})
            rhs40_ = len((d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])
            rhs41_ = ((d_349_v189_)[(d_70_v2_).ispropersubset(d_70_v2_)] if ((d_70_v2_).ispropersubset(d_70_v2_)) in (d_349_v189_) else d_262_v112_)
            rhs42_ = default__.fm6((d_260_v110_).f24, (d_260_v110_).f24, not((d_260_v110_).f24), d_84_globalState_)
            rhs43_ = ((d_346_v186_).f26) - ((len(d_350_v190_)) - (d_73_v6_))
            lhs23_ = d_84_globalState_
            lhs24_ = d_84_globalState_
            lhs25_ = d_84_globalState_
            lhs23_.f5 = rhs40_
            d_262_v112_ = rhs41_
            lhs24_.f5 = rhs42_
            lhs25_.f5 = rhs43_
            (d_84_globalState_).f5 = (d_73_v6_ if ((d_260_v110_).f24 if (d_260_v110_).f24 else (d_260_v110_).f24) else (d_346_v186_).f26)
            d_351_v191_: _dafny.Seq
            d_351_v191_ = _dafny.SeqWithoutIsStrInference([d_260_v110_, d_260_v110_, d_260_v110_, d_260_v110_, d_260_v110_])
            d_352_v192_: _dafny.Set
            d_352_v192_ = _dafny.Set({d_70_v2_})
            d_353_v193_: T0
            nw54_ = C0()
            nw54_.ctor__(len(d_240_v106_), d_69_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bacmjqi")))
            d_353_v193_ = nw54_
            d_354_v194_: _dafny.Seq
            d_354_v194_ = _dafny.SeqWithoutIsStrInference([d_353_v193_])
            d_355_v195_: _dafny.Map
            d_355_v195_ = _dafny.Map({(0) - (len(d_352_v192_)): (d_354_v194_)[default__.safeIndex((d_346_v186_).f26, len(d_354_v194_))]})
            rhs44_ = default__.safeDivisionInt((d_346_v186_).f26, len((_dafny.SeqWithoutIsStrInference([d_260_v110_])) + (d_351_v191_)))
            rhs45_ = default__.safeDivisionInt(((d_346_v186_).f26) - ((d_346_v186_).f26), len((d_355_v195_) | (_dafny.Map({d_73_v6_: d_353_v193_}))))
            rhs46_ = (d_260_v110_).f24
            lhs26_ = d_84_globalState_
            lhs27_ = d_84_globalState_
            lhs28_ = d_84_globalState_
            lhs26_.f5 = rhs44_
            lhs27_.f5 = rhs45_
            lhs28_.f21 = rhs46_
        elif True:
            d_356_v196_: _dafny.Map
            d_356_v196_ = _dafny.Map({d_68_v0_: d_69_v1_})
            (d_84_globalState_).f20 = ((d_356_v196_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxd"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxd"))) in (d_356_v196_) else d_69_v1_)
            d_357_v197_: D3
            d_357_v197_ = D3_DC11(d_74_v7_)
            d_358_v198_: D3
            d_358_v198_ = D3_DC11((d_357_v197_).cf19)
            source4_ = d_357_v197_
            if source4_.is_DC12:
                d_359___mcc_h29_ = source4_.cf20
                d_360___mcc_h30_ = source4_.cf21
                d_361_cf21_ = d_360___mcc_h30_
                d_362_cf20_ = d_359___mcc_h29_
                (d_84_globalState_).f5 = d_73_v6_
                d_73_v6_ = len(d_83_v14_)
                d_71_v3_ = (d_71_v3_ if d_69_v1_ else d_71_v3_)
                d_343_v185_ = d_343_v185_
            elif source4_.is_DC13:
                d_363___mcc_h31_ = source4_.cf22
                d_364___mcc_h32_ = source4_.cf23
                d_365_cf23_ = d_364___mcc_h32_
                d_366_cf22_ = d_363___mcc_h31_
                d_367_v199_: _dafny.Array
                nw55_ = _dafny.Array(_dafny.Seq({}), 9)
                d_367_v199_ = nw55_
                d_368_v200_: D4
                d_368_v200_ = D4_DC15(d_240_v106_)
                index26_ = default__.safeIndex(220, (d_367_v199_).length(0))
                (d_367_v199_)[index26_] = (d_368_v200_).cf25
                index27_ = default__.safeIndex(220, (d_367_v199_).length(0))
                (d_367_v199_)[index27_] = d_240_v106_
                d_369_v201_: _dafny.Seq
                d_370_v202_: _dafny.Array
                d_371_v203_: int
                d_372_v204_: bool
                out52_: _dafny.Seq
                out53_: _dafny.Array
                out54_: int
                out55_: bool
                out52_, out53_, out54_, out55_ = default__.m0(d_84_globalState_)
                d_369_v201_ = out52_
                d_370_v202_ = out53_
                d_371_v203_ = out54_
                d_372_v204_ = out55_
                (d_84_globalState_).f5 = len((d_260_v110_).f25)
                index28_ = default__.safeIndex(11, (d_71_v3_).length(0))
                (d_71_v3_)[index28_] = (d_260_v110_).f24
                d_373_v205_: _dafny.Map
                d_373_v205_ = _dafny.Map({(d_260_v110_).f24: d_75_v8_})
                index29_ = default__.safeIndex(11, (d_71_v3_).length(0))
                (d_71_v3_)[index29_] = (d_366_cf22_).fm1(d_373_v205_, _dafny.Map({d_371_v203_: (d_260_v110_).f24}), (d_74_v7_) + (d_74_v7_), d_84_globalState_)
            elif source4_.is_DC11:
                d_374___mcc_h33_ = source4_.cf19
                d_375_cf19_ = d_374___mcc_h33_
                d_376_v206_: _dafny.Map
                d_376_v206_ = _dafny.Map({(d_260_v110_).f24: (d_346_v186_).f26})
                (d_84_globalState_).f5 = len((_dafny.Map({(d_260_v110_).f24: (d_346_v186_).f26})) | ((_dafny.Map({(d_260_v110_).f24: d_73_v6_})) | (d_376_v206_)))
                d_377_v207_: _dafny.Seq
                d_378_v208_: _dafny.Array
                d_379_v209_: int
                d_380_v210_: bool
                out56_: _dafny.Seq
                out57_: _dafny.Array
                out58_: int
                out59_: bool
                out56_, out57_, out58_, out59_ = default__.m0(d_84_globalState_)
                d_377_v207_ = out56_
                d_378_v208_ = out57_
                d_379_v209_ = out58_
                d_380_v210_ = out59_
                d_381_v211_: _dafny.Seq
                d_382_v212_: _dafny.Array
                d_383_v213_: int
                d_384_v214_: bool
                out60_: _dafny.Seq
                out61_: _dafny.Array
                out62_: int
                out63_: bool
                out60_, out61_, out62_, out63_ = default__.m0(d_84_globalState_)
                d_381_v211_ = out60_
                d_382_v212_ = out61_
                d_383_v213_ = out62_
                d_384_v214_ = out63_
                d_385_v215_: _dafny.Seq
                d_385_v215_ = _dafny.SeqWithoutIsStrInference([d_74_v7_])
                d_386_v216_: _dafny.Map
                d_386_v216_ = _dafny.Map({(d_260_v110_).f24: d_75_v8_})
                d_387_v217_: _dafny.Map
                d_387_v217_ = _dafny.Map({d_73_v6_: d_74_v7_})
                (d_84_globalState_).f5 = len(_dafny.Map({(d_385_v215_)[default__.safeIndex(929, len(d_385_v215_))]: (d_346_v186_).fm1((d_386_v216_).set(d_380_v210_, d_75_v8_), _dafny.Map({(d_346_v186_).f26: d_380_v210_}), (((d_387_v217_)[len(d_70_v2_)] if (len(d_70_v2_)) in (d_387_v217_) else d_375_cf19_)).set(default__.safeIndex((d_346_v186_).f26, len(((d_387_v217_)[len(d_70_v2_)] if (len(d_70_v2_)) in (d_387_v217_) else d_375_cf19_))), default__.fm6(d_69_v1_, not(d_380_v210_), (d_260_v110_).f24, d_84_globalState_)), d_84_globalState_)}))
            elif True:
                d_388___mcc_h34_ = source4_.cf24
                d_389_cf24_ = d_388___mcc_h34_
                d_390_v218_: _dafny.Map
                d_390_v218_ = _dafny.Map({False: d_75_v8_})
                d_391_v219_: _dafny.Map
                d_391_v219_ = _dafny.Map({d_73_v6_: (d_260_v110_).f24})
                d_392_v220_: T0
                nw56_ = C0()
                nw56_.ctor__(d_73_v6_, (d_260_v110_).fm1(d_390_v218_, d_391_v219_, (d_74_v7_).set(default__.safeIndex(d_73_v6_, len(d_74_v7_)), (d_346_v186_).f26), d_84_globalState_), _dafny.SeqWithoutIsStrInference([d_75_v8_ for d_393_i18_ in range(default__.abs(940))]))
                d_392_v220_ = nw56_
                d_394_v221_: D4
                d_394_v221_ = D4_DC16((d_346_v186_).f26, (d_240_v106_)[default__.safeIndex(d_73_v6_, len(d_240_v106_))], d_392_v220_)
                d_395_v222_: _dafny.Map
                d_395_v222_ = _dafny.Map({d_394_v221_: (len(d_262_v112_)) + ((0) - (len((d_260_v110_).f25)))})
                d_395_v222_ = (d_395_v222_).set(d_394_v221_, len(d_240_v106_))
                nw57_ = C0()
                def iife38_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in (d_74_v7_).Elements:
                        d_396_v223_: int = compr_10_
                        if (d_396_v223_) in (d_74_v7_):
                            coll10_[default__.safeDivisionInt(d_396_v223_, (d_346_v186_).f26)] = (d_260_v110_).f24
                    return _dafny.Map(coll10_)
                nw57_.ctor__(((d_346_v186_).f26) - ((d_346_v186_).fm3((d_346_v186_).f26, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(iife38_()
                )])]), 638, (d_260_v110_).f25, d_84_globalState_)), (d_69_v1_ if (d_260_v110_).f24 else False), (d_260_v110_).f25)
                d_346_v186_ = nw57_
                d_397_v224_: _dafny.Set
                d_397_v224_ = _dafny.Set({d_357_v197_})
                (d_84_globalState_).f5 = (len(d_262_v112_)) * (len(d_397_v224_))
                d_398_v225_: _dafny.Map
                d_398_v225_ = _dafny.Map({d_241_v107_: (d_260_v110_).f24})
                d_399_v226_: D5
                d_399_v226_ = D5_DC20(d_398_v225_)
                rhs47_ = (d_241_v107_) in ((d_399_v226_).cf32)
                rhs48_ = (d_346_v186_).f26
                lhs29_ = d_84_globalState_
                lhs30_ = d_84_globalState_
                lhs29_.f20 = rhs47_
                lhs30_.f5 = rhs48_
            d_400_v227_: _dafny.Map
            d_400_v227_ = _dafny.Map({d_71_v3_: (d_346_v186_).f26})
            d_401_v228_: _dafny.Map
            d_401_v228_ = _dafny.Map({(d_73_v6_ if (d_260_v110_).f24 else len(d_400_v227_)): (d_73_v6_) != (len(d_83_v14_))})
            d_401_v228_ = (d_401_v228_).set(488, d_69_v1_)
            index30_ = default__.safeIndex(800, (d_71_v3_).length(0))
            (d_71_v3_)[index30_] = (d_260_v110_).f24
            d_402_v229_: T0
            nw58_ = C0()
            nw58_.ctor__((d_346_v186_).f26, d_69_v1_, (d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])
            d_402_v229_ = nw58_
            d_403_v230_: _dafny.Seq
            d_403_v230_ = _dafny.SeqWithoutIsStrInference([d_402_v229_, d_402_v229_, d_402_v229_])
            index31_ = default__.safeIndex(800, (d_71_v3_).length(0))
            rhs49_ = (d_346_v186_).f26
            rhs50_ = (-804) <= ((d_346_v186_).f26)
            rhs51_ = (d_74_v7_) != ((d_74_v7_).set(default__.safeIndex(len(d_403_v230_), len(d_74_v7_)), d_73_v6_))
            rhs52_ = (d_240_v106_)[default__.safeIndex(277, len(d_240_v106_))]
            lhs31_ = d_71_v3_
            lhs32_ = default__.safeIndex(800, (d_71_v3_).length(0))
            d_73_v6_ = rhs49_
            d_69_v1_ = rhs50_
            lhs31_[lhs32_] = rhs51_
            d_69_v1_ = rhs52_
            (d_84_globalState_).f5 = (d_346_v186_).f26
        d_404_v231_: _dafny.Array
        nw59_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_404_v231_ = nw59_
        d_405_v232_: _dafny.Map
        d_405_v232_ = _dafny.Map({d_404_v231_: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, (d_260_v110_).f24]))})
        d_405_v232_ = (d_405_v232_).set(d_404_v231_, d_341_v182_)
        d_406_v233_: C0
        nw60_ = C0()
        nw60_.ctor__(d_73_v6_, (d_260_v110_).f24, (d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))])
        d_406_v233_ = nw60_
        index32_ = default__.safeIndex(413, (d_71_v3_).length(0))
        (d_71_v3_)[index32_] = (d_260_v110_).f24
        d_407_v234_: _dafny.Set
        d_407_v234_ = _dafny.Set({d_71_v3_, d_71_v3_, d_71_v3_})
        index33_ = default__.safeIndex(413, (d_71_v3_).length(0))
        index34_ = default__.safeIndex(285, (d_279_v127_).length(0))
        rhs53_ = (d_407_v234_).isdisjoint(d_407_v234_)
        rhs54_ = d_68_v0_
        rhs55_ = ((d_279_v127_)[default__.safeIndex(285, (d_279_v127_).length(0))]) + (d_68_v0_)
        lhs33_ = d_71_v3_
        lhs34_ = default__.safeIndex(413, (d_71_v3_).length(0))
        lhs35_ = d_279_v127_
        lhs36_ = default__.safeIndex(285, (d_279_v127_).length(0))
        lhs37_ = d_84_globalState_
        lhs33_[lhs34_] = rhs53_
        lhs35_[lhs36_] = rhs54_
        lhs37_.f19 = rhs55_
        _dafny.print((d_68_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_69_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_70_v2_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v3_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_72_v4_).cf0).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_73_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v7_) == (_dafny.SeqWithoutIsStrInference([462, 462, 462]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_77_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v11_) == (_dafny.MultiSet([462]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_v12_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([462])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[0]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[1]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[2]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[3]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[4]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[5]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[6]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[7]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[8]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[9]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[10]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[11]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[12]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[13]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[14]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[15]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[16]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[17]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[18]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[19]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[20]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[21]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[22]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[23]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[24]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[25]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[26]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[27]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_80_v13_)[28]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_83_v14_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_84_globalState_).f0).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_.f1) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_84_globalState_.f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f6) == (_dafny.Map({213444: _dafny.CodePoint('j'), 668: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_84_globalState_.f8)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_.f14) == (_dafny.MultiSet([462]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[0]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[1]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[2]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[3]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[4]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[5]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[6]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[7]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[8]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[9]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[10]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[11]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[12]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[13]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[14]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[15]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[16]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[17]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[18]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[19]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[20]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[21]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[22]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[23]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[24]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[25]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[26]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[27]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_84_globalState_).f15)[28]) == (_dafny.Map({False: -176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_84_globalState_.f19).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[0]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[1]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[2]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[3]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[4]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[5]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[6]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[7]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[8]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[9]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[10]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[11]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[12]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[13]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[14]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[15]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[16]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[17]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[18]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[19]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[20]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[21]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[22]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_238_v105_)[23]) == (_dafny.Map({False: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v106_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v107_) == (_dafny.Map({49: 462}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v108_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v108_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v110_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_260_v110_).f25).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v111_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v111_).cf11).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_261_v111_).cf11).f25).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v111_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v111_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v112_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v115_) == (_dafny.Set({_dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_267_v116_) == (_dafny.Set({_dafny.Set({_dafny.CodePoint('x')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_268_i11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_279_v127_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_279_v127_)[15]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_340_v181_) == (_dafny.Map({_dafny.CodePoint('x'): 213443}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v182_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v183_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v185_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_345_i17_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_346_v186_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_347_v187_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_405_v232_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_406_v233_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_407_v234_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False, int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC9(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)

class D2_DC9(D2, NamedTuple('DC9', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC10(D2, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC12(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D3_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D3_DC14)

class D3_DC12(D3, NamedTuple('DC12', [('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC13(D3, NamedTuple('DC13', [('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC13({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC13) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC14(D3, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC16(int(0), False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D4_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D4_DC18)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D4_DC19)

class D4_DC16(D4, NamedTuple('DC16', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC17(D4, NamedTuple('DC17', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC17({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC17) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC18(D4, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC19(D4, NamedTuple('DC19', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC19({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC19) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC21(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D5_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D5_DC20)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D5_DC22)

class D5_DC21(D5, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC20(D5, NamedTuple('DC20', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC20({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC20) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC22(D5, NamedTuple('DC22', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC22({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC22) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D6_DC23)

class D6_DC23(D6, NamedTuple('DC23', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f1: _dafny.Set = _dafny.Set({})
        self.f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f5: int = int(0)
        self.f8: _dafny.Seq = _dafny.Seq({})
        self.f14: _dafny.MultiSet = _dafny.MultiSet({})
        self.f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f20: bool = False
        self.f21: bool = False
        self._f0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        self._f3: bool = False
        self._f6: _dafny.Map = _dafny.Map({})
        self._f7: bool = False
        self._f9: bool = False
        self._f10: int = int(0)
        self._f11: bool = False
        self._f12: bool = False
        self._f13: bool = False
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        self._f16: int = int(0)
        self._f17: bool = False
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f22: int = int(0)
        self._f23: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f6(self):
        return self._f6
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
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
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
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23

class C0(T0, T1):
    def  __init__(self):
        self._f24: bool = False
        self._f25: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f26: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f26, f24, f25):
        (self)._f26 = f26
        (self)._f24 = f24
        (self)._f25 = f25

    def fm0(self, p0, p1, p2, p3, globalState):
        return (_dafny.Map({-128: (self).f24})) | (_dafny.Map({(self).f26: (self).f24}))

    def fm1(self, p0, p1, p2, globalState):
        return ((self).f24) and ((self).f24)

    def fm2(self, p0, p1, globalState):
        return D0_DC1((self).f24, (self).f24, (self).f26)

    def fm3(self, p0, p1, p2, p3, globalState):
        return (self).f26

    @property
    def f26(self):
        return self._f26
