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
    def fm0(p0, p1, globalState):
        return 991

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(409, -296):
                d_0_v0_: int = compr_0_
                if ((409) <= (d_0_v0_)) and ((d_0_v0_) < (-296)):
                    coll0_[(d_0_v0_) + (260)] = D0_DC3(D0_DC3(D0_DC3(D0_DC2(False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({202}))])), len(_dafny.Map({False: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([65 for d_1_i0_ in range(default__.abs(-123))])): not(True)}))})), True))))
            return _dafny.Map(coll0_)
        return (iife0_()
        ) | ((_dafny.Map({944: D0_DC3(D0_DC2(False, 754, 561, False))})) | (_dafny.Map({(_dafny.MultiSet([-113])).cardinality: D0_DC3(D0_DC3(D0_DC2(True, 111, -539, not(True))))})))

    @staticmethod
    def fm2(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(466, len(_dafny.Map({_dafny.CodePoint('v'): 696})))])

    @staticmethod
    def fm3(p0, globalState):
        return (D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkxmuj")))).cf10

    @staticmethod
    def fm5(globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm6(p0, p1, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([False, True])) + (_dafny.SeqWithoutIsStrInference([not(False)]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm7(globalState):
        return ((_dafny.Set({not(True)})) | (_dafny.Set({False}))) - (_dafny.Set({True}))

    @staticmethod
    def fm10(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, not(not(False)), False])) + (_dafny.SeqWithoutIsStrInference([True, False, False, not(False)]))) + (_dafny.SeqWithoutIsStrInference([not(True)]))

    @staticmethod
    def fm11(p0, p1, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbvst")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmftxyldf")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2_i1_ in range(default__.abs(111))])])).issubset((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_3_i0_ in range(default__.abs(-760))])]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgiqbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxnfl"))])))

    @staticmethod
    def fm12(globalState):
        if (True if not(False) else False):
            return _dafny.Map({-835: len(_dafny.SeqWithoutIsStrInference([not(False), not(False)]))})
        elif True:
            return _dafny.Map({156: (_dafny.MultiSet([not(True), not(True)])).cardinality})

    @staticmethod
    def fm13(p0, p1, globalState):
        return _dafny.Set({_dafny.CodePoint('b'), _dafny.CodePoint('g')})

    @staticmethod
    def m2(p0, p1, globalState):
        if p0:
            d_4_v0_: int
            d_4_v0_ = 274
            d_4_v0_ = d_4_v0_
            d_5_v1_: _dafny.Set
            d_5_v1_ = _dafny.Set({d_4_v0_})
            if ((d_5_v1_).ispropersubset(d_5_v1_)) or (not (not(p0)) or (p0)):
                d_6_v2_: str
                d_6_v2_ = _dafny.CodePoint('e')
                (globalState).f18 = d_6_v2_
                d_7_v3_: _dafny.Seq
                d_7_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0])])
                d_8_v4_: C0
                nw0_ = C0()
                nw0_.ctor__(not(p0), _dafny.Map({d_4_v0_: (d_7_v3_)[default__.safeIndex(d_4_v0_, len(d_7_v3_))]}))
                d_8_v4_ = nw0_
                d_9_v5_: _dafny.Set
                d_9_v5_ = _dafny.Set({d_8_v4_})
                (globalState).f19 = not((d_9_v5_).ispropersubset(d_9_v5_))
                d_10_v6_: _dafny.MultiSet
                d_10_v6_ = _dafny.MultiSet([d_8_v4_.f24, p1])
                d_11_v7_: _dafny.Array
                nw1_ = _dafny.Array(None, 22)
                nw1_[int(0)] = p1
                nw1_[int(1)] = True
                nw1_[int(2)] = p0
                nw1_[int(3)] = not(True)
                nw1_[int(4)] = p1
                nw1_[int(5)] = d_8_v4_.f24
                nw1_[int(6)] = p0
                nw1_[int(7)] = not(p0)
                nw1_[int(8)] = False
                nw1_[int(9)] = p1
                nw1_[int(10)] = p0
                nw1_[int(11)] = d_8_v4_.f24
                nw1_[int(12)] = default__.fm11(d_10_v6_, d_4_v0_, globalState)
                nw1_[int(13)] = p0
                nw1_[int(14)] = p0
                nw1_[int(15)] = False
                nw1_[int(16)] = p0
                nw1_[int(17)] = d_8_v4_.f24
                nw1_[int(18)] = d_8_v4_.f24
                nw1_[int(19)] = p0
                nw1_[int(20)] = p0
                nw1_[int(21)] = p1
                d_11_v7_ = nw1_
                d_12_v8_: _dafny.Map
                d_12_v8_ = _dafny.Map({d_11_v7_: d_11_v7_})
                d_12_v8_ = (d_12_v8_).set(d_11_v7_, d_11_v7_)
                d_13_v9_: _dafny.Seq
                d_13_v9_ = _dafny.SeqWithoutIsStrInference([p1, p0])
                d_14_v10_: _dafny.Seq
                d_14_v10_ = _dafny.SeqWithoutIsStrInference([-330, len(d_13_v9_)])
                d_15_v11_: _dafny.Set
                d_15_v11_ = _dafny.Set({d_14_v10_, _dafny.SeqWithoutIsStrInference([(0) - (d_4_v0_)]), _dafny.SeqWithoutIsStrInference([d_4_v0_]), d_14_v10_, _dafny.SeqWithoutIsStrInference([d_4_v0_, d_4_v0_, d_4_v0_])})
                index0_ = default__.safeIndex(790, (d_11_v7_).length(0))
                (d_11_v7_)[index0_] = (d_15_v11_).ispropersubset(d_15_v11_)
                d_16_v12_: _dafny.Seq
                d_16_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlttakuyp"))
                index1_ = default__.safeIndex(790, (d_11_v7_).length(0))
                rhs0_ = p1
                rhs1_ = ((d_16_v12_) + (d_16_v12_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfdjcpnpq")))
                rhs2_ = (575) * (d_4_v0_)
                rhs3_ = d_8_v4_
                lhs0_ = d_11_v7_
                lhs1_ = default__.safeIndex(790, (d_11_v7_).length(0))
                lhs0_[lhs1_] = rhs0_
                d_16_v12_ = rhs1_
                d_4_v0_ = rhs2_
                d_8_v4_ = rhs3_
                d_8_v4_ = (d_8_v4_ if (d_8_v4_.f24) or ((d_11_v7_)[default__.safeIndex(790, (d_11_v7_).length(0))]) else d_8_v4_)
            elif True:
                d_17_v13_: _dafny.Seq
                d_17_v13_ = _dafny.SeqWithoutIsStrInference([p1])
                d_18_v14_: _dafny.Map
                d_18_v14_ = _dafny.Map({d_4_v0_: d_17_v13_})
                d_19_v15_: C0
                nw2_ = C0()
                nw2_.ctor__(not (p1) or (p0), d_18_v14_)
                d_19_v15_ = nw2_
                d_20_v16_: _dafny.Seq
                d_20_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqabmwu"))
                d_21_v17_: str
                d_21_v17_ = _dafny.CodePoint('u')
                d_22_v18_: _dafny.MultiSet
                d_22_v18_ = _dafny.MultiSet([p0, p1])
                d_23_v19_: _dafny.Map
                d_23_v19_ = _dafny.Map({d_21_v17_: len((_dafny.Map({p0: d_4_v0_})).set(default__.fm11(d_22_v18_, 870, globalState), d_4_v0_))})
                d_17_v13_ = ((default__.fm10(len(d_20_v16_), 751, globalState)) + ((d_17_v13_) + (d_17_v13_))).set(default__.safeIndex(len(d_23_v19_), len((default__.fm10(len(d_20_v16_), 751, globalState)) + ((d_17_v13_) + (d_17_v13_)))), not(p0))
                d_4_v0_ = 825
                d_24_v20_: _dafny.Array
                nw3_ = _dafny.Array(int(0), 1)
                d_24_v20_ = nw3_
                index2_ = default__.safeIndex(390, (d_24_v20_).length(0))
                (d_24_v20_)[index2_] = d_4_v0_
                index3_ = default__.safeIndex(390, (d_24_v20_).length(0))
                (d_24_v20_)[index3_] = ((d_4_v0_ if p1 else d_4_v0_)) * (d_4_v0_)
                (globalState).f19 = (d_17_v13_)[default__.safeIndex(d_4_v0_, len(d_17_v13_))]
            (globalState).f6 = p1
            (globalState).f6 = p1
            d_25_v21_: _dafny.Map
            d_25_v21_ = _dafny.Map({p1: p1})
            d_26_v22_: _dafny.Map
            d_26_v22_ = _dafny.Map({d_4_v0_: ((d_25_v21_)[False] if (False) in (d_25_v21_) else p0)})
            d_27_v23_: _dafny.Seq
            d_27_v23_ = _dafny.SeqWithoutIsStrInference([p0, p1, p0, p0])
            (globalState).f19 = ((d_27_v23_)[default__.safeIndex(default__.fm0(p0, p1, globalState), len(d_27_v23_))] if not(((d_26_v22_)[(0) - ((0) - (d_4_v0_))] if ((0) - ((0) - (d_4_v0_))) in (d_26_v22_) else False)) else p1)
        elif True:
            d_28_v24_: int
            d_28_v24_ = 191
            d_29_v25_: _dafny.Seq
            d_29_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvmshmd"))
            d_28_v24_ = (len(d_29_v25_)) + (d_28_v24_)
            d_30_v26_: _dafny.Array
            nw4_ = _dafny.Array(_dafny.Map({}), 29)
            d_30_v26_ = nw4_
            d_30_v26_ = d_30_v26_
            d_31_v27_: _dafny.Seq
            d_31_v27_ = _dafny.SeqWithoutIsStrInference([False, p0])
            d_32_v28_: D0
            d_32_v28_ = D0_DC1(d_31_v27_, -405)
            d_33_v29_: _dafny.Map
            d_33_v29_ = _dafny.Map({d_28_v24_: (d_32_v28_ if p0 else d_32_v28_)})
            d_34_v30_: _dafny.Map
            d_34_v30_ = _dafny.Map({d_28_v24_: p1})
            d_35_v31_: _dafny.MultiSet
            d_35_v31_ = _dafny.MultiSet([default__.fm0(p0, p0, globalState), len(d_34_v30_)])
            d_36_v32_: _dafny.Set
            d_36_v32_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([p1]))})
            d_33_v29_ = (d_33_v29_).set(((d_35_v31_)[d_28_v24_] if (d_28_v24_) in (d_35_v31_) else len(d_36_v32_)), d_32_v28_)
            d_37_v33_: _dafny.Array
            nw5_ = _dafny.Array(_dafny.MultiSet({}), 25)
            d_37_v33_ = nw5_
            d_37_v33_ = d_37_v33_
            d_38_v34_: _dafny.Array
            nw6_ = _dafny.Array(False, 3)
            d_38_v34_ = nw6_
            d_39_v35_: _dafny.MultiSet
            d_39_v35_ = _dafny.MultiSet([d_38_v34_, d_38_v34_])
            d_40_v36_: _dafny.MultiSet
            d_40_v36_ = (d_39_v35_ if not(p1) else d_39_v35_)
            source0_ = d_40_v36_
            d_41___mcc_h0_ = source0_
            d_42_cf18_ = d_41___mcc_h0_
            d_43_v37_: _dafny.Map
            d_43_v37_ = _dafny.Map({p0: -759})
            d_44_v38_: _dafny.Seq
            d_44_v38_ = _dafny.SeqWithoutIsStrInference([d_38_v34_, d_38_v34_])
            d_45_v39_: D3
            d_45_v39_ = D3_DC7(d_43_v37_, p0, p1, p1, (d_44_v38_)[default__.safeIndex((d_32_v28_).cf2, len(d_44_v38_))])
            d_46_v40_: C1
            nw7_ = C1()
            nw7_.ctor__(p1, (d_45_v39_).cf14)
            d_46_v40_ = nw7_
            d_47_v41_: str
            d_47_v41_ = _dafny.CodePoint('t')
            (globalState).f18 = d_47_v41_
            (globalState).f19 = not ((d_46_v40_).fm4(373, 230, d_46_v40_.f23, globalState)) or ((p0 if p1 else (d_45_v39_).cf12))
            (globalState).f19 = (d_46_v40_).f22
        d_48_v42_: _dafny.Array
        nw8_ = _dafny.Array(int(0), 24)
        d_48_v42_ = nw8_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_48_v42_).length(0)):
            d_49_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_49_i0_)) and ((d_49_i0_) < ((d_48_v42_).length(0)))):
                (d_48_v42_)[(d_49_i0_)] = default__.safeDivisionInt(d_49_i0_, -822)
        d_50_i1_: int
        d_50_i1_ = 0
        with _dafny.label("0"):
            while not(p0):
                with _dafny.c_label("0"):
                    if (d_50_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_50_i1_ = (d_50_i1_) + (1)
                    d_51_v43_: int
                    d_51_v43_ = 493
                    d_51_v43_ = d_51_v43_
                    d_52_v44_: _dafny.MultiSet
                    d_52_v44_ = _dafny.MultiSet([p1])
                    d_53_v45_: C1
                    nw9_ = C1()
                    nw9_.ctor__(True, p1)
                    d_53_v45_ = nw9_
                    d_54_v46_: _dafny.MultiSet
                    d_54_v46_ = _dafny.MultiSet([d_53_v45_])
                    d_55_v47_: D0
                    d_55_v47_ = D0_DC2(p0, len(_dafny.Map({not(p1): (D0_DC2(default__.fm11(d_52_v44_, d_51_v43_, globalState), d_51_v43_, d_51_v43_, True)).cf3})), (d_54_v46_).cardinality, p0)
                    d_56_v48_: _dafny.Set
                    d_56_v48_ = _dafny.Set({(d_55_v47_).cf3, p1, (d_53_v45_).f22, d_53_v45_.f23})
                    d_57_v49_: _dafny.Map
                    d_57_v49_ = _dafny.Map({d_53_v45_.f23: p0})
                    d_58_v50_: _dafny.Map
                    d_58_v50_ = _dafny.Map({len(d_56_v48_): ((d_57_v49_)[d_53_v45_.f23] if (d_53_v45_.f23) in (d_57_v49_) else d_53_v45_.f23)})
                    d_59_v51_: _dafny.Map
                    d_59_v51_ = _dafny.Map({d_58_v50_: p0})
                    d_59_v51_ = (d_59_v51_).set(d_58_v50_, (d_53_v45_).f22)
                    d_60_v52_: str
                    d_60_v52_ = _dafny.CodePoint('q')
                    d_61_v53_: _dafny.Seq
                    d_61_v53_ = _dafny.SeqWithoutIsStrInference([d_60_v52_, d_60_v52_])
                    d_62_v54_: _dafny.Map
                    d_62_v54_ = _dafny.Map({d_53_v45_.f23: d_61_v53_})
                    d_63_v55_: _dafny.Array
                    nw10_ = _dafny.Array(None, 1)
                    nw10_[int(0)] = default__.fm2(_dafny.MultiSet(((d_62_v54_)[d_53_v45_.f23] if (d_53_v45_.f23) in (d_62_v54_) else _dafny.SeqWithoutIsStrInference([d_60_v52_]))), globalState)
                    d_63_v55_ = nw10_
                    d_64_v56_: _dafny.Set
                    d_64_v56_ = _dafny.Set({default__.fm5(globalState)})
                    d_65_v57_: _dafny.Seq
                    d_65_v57_ = _dafny.SeqWithoutIsStrInference([d_64_v56_, d_64_v56_, default__.fm13(len(d_61_v53_), d_51_v43_, globalState)])
                    d_66_v58_: _dafny.Seq
                    d_66_v58_ = _dafny.SeqWithoutIsStrInference([len(d_65_v57_)])
                    index4_ = default__.safeIndex(816, (d_63_v55_).length(0))
                    (d_63_v55_)[index4_] = d_66_v58_
                    d_67_v59_: _dafny.MultiSet
                    d_67_v59_ = _dafny.MultiSet([_dafny.CodePoint('m')])
                    d_68_v60_: _dafny.Seq
                    d_68_v60_ = _dafny.SeqWithoutIsStrInference([(d_66_v58_).set(default__.safeIndex(756, len(d_66_v58_)), d_51_v43_), _dafny.SeqWithoutIsStrInference([865 for d_69_i2_ in range(default__.abs(3))])])
                    index5_ = default__.safeIndex(816, (d_63_v55_).length(0))
                    (d_63_v55_)[index5_] = (default__.fm2(d_67_v59_, globalState)) + ((d_68_v60_)[default__.safeIndex(d_51_v43_, len(d_68_v60_))])
                    if (d_53_v45_).f22:
                        d_70_v61_: _dafny.Map
                        d_70_v61_ = _dafny.Map({p1: d_51_v43_})
                        d_71_v62_: D3
                        d_71_v62_ = D3_DC8(((d_70_v61_)[(d_53_v45_).f22] if ((d_53_v45_).f22) in (d_70_v61_) else d_51_v43_), (d_51_v43_ if (d_53_v45_).f22 else d_51_v43_))
                        pat_let_tv0_ = d_51_v43_
                        def iife1_(_pat_let0_0):
                            def iife2_(d_72_dt__update__tmp_h0_):
                                def iife3_(_pat_let1_0):
                                    def iife4_(d_73_dt__update_hcf17_h0_):
                                        return D3_DC8((d_72_dt__update__tmp_h0_).cf16, d_73_dt__update_hcf17_h0_)
                                    return iife4_(_pat_let1_0)
                                return iife3_(pat_let_tv0_)
                            return iife2_(_pat_let0_0)
                        d_71_v62_ = (d_71_v62_ if (d_53_v45_).f22 else iife1_(d_71_v62_))
                        d_74_v64_: T0
                        nw11_ = C0()
                        def iife5_():
                            coll1_ = _dafny.Map()
                            compr_1_: int
                            for compr_1_ in _dafny.IntegerRange(-478, 307):
                                d_75_v63_: int = compr_1_
                                if ((-478) <= (d_75_v63_)) and ((d_75_v63_) < (307)):
                                    coll1_[default__.safeModuloInt(d_75_v63_, d_51_v43_)] = _dafny.SeqWithoutIsStrInference([False])
                            return _dafny.Map(coll1_)
                        nw11_.ctor__(p1, iife5_()
                        )
                        d_74_v64_ = nw11_
                        d_76_v65_: _dafny.Array
                        nw12_ = _dafny.Array(None, 4)
                        nw12_[int(0)] = d_74_v64_
                        nw12_[int(1)] = d_74_v64_
                        nw12_[int(2)] = d_74_v64_
                        nw12_[int(3)] = d_74_v64_
                        d_76_v65_ = nw12_
                        d_76_v65_ = d_76_v65_
                        d_77_v66_: C1
                        nw13_ = C1()
                        nw13_.ctor__(not(False), p1)
                        d_77_v66_ = nw13_
                        rhs4_ = d_60_v52_
                        rhs5_ = d_51_v43_
                        lhs2_ = globalState
                        lhs2_.f18 = rhs4_
                        d_51_v43_ = rhs5_
                        d_78_v67_: _dafny.Array
                        nw14_ = _dafny.Array(False, 10)
                        d_78_v67_ = nw14_
                        d_78_v67_ = d_78_v67_
                    elif True:
                        d_51_v43_ = len((d_61_v53_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))))
                        d_79_v68_: _dafny.Seq
                        d_79_v68_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                        d_80_v69_: _dafny.Map
                        d_80_v69_ = _dafny.Map({d_79_v68_: d_53_v45_.f23})
                        d_51_v43_ = (default__.safeModuloInt(d_51_v43_, d_51_v43_) if d_53_v45_.f23 else len(d_80_v69_))
                        d_81_v70_: _dafny.Array
                        nw15_ = _dafny.Array(False, 26)
                        d_81_v70_ = nw15_
                        index6_ = default__.safeIndex(65, (d_81_v70_).length(0))
                        (d_81_v70_)[index6_] = p0
                        index7_ = default__.safeIndex(65, (d_81_v70_).length(0))
                        (d_81_v70_)[index7_] = ((-514) <= (861)) or ((d_53_v45_).f22)
                        d_82_v71_: _dafny.Array
                        def lambda0_(d_83_v52_):
                            def lambda1_(d_84_i3_):
                                return d_83_v52_

                            return lambda1_

                        init0_ = lambda0_(d_60_v52_)
                        nw16_ = _dafny.Array(None, 22)
                        for i0_0_ in range(nw16_.length(0)):
                            nw16_[i0_0_] = init0_(i0_0_)
                        d_82_v71_ = nw16_
                        index8_ = default__.safeIndex(691, (d_82_v71_).length(0))
                        (d_82_v71_)[index8_] = d_60_v52_
                        index9_ = default__.safeIndex(691, (d_82_v71_).length(0))
                        (d_82_v71_)[index9_] = d_60_v52_
                        d_85_v72_: _dafny.Array
                        nw17_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_85_v72_ = nw17_
                        d_86_v73_: _dafny.Array
                        nw18_ = _dafny.Array(None, 1)
                        nw18_[int(0)] = d_85_v72_
                        d_86_v73_ = nw18_
                        index10_ = default__.safeIndex(739, (d_86_v73_).length(0))
                        (d_86_v73_)[index10_] = d_85_v72_
                        index11_ = default__.safeIndex(739, (d_86_v73_).length(0))
                        rhs6_ = d_85_v72_
                        rhs7_ = (0) - (len(((d_79_v68_ if p0 else _dafny.SeqWithoutIsStrInference([p0, (d_53_v45_).f22, (d_53_v45_).f22]))) + (d_79_v68_)))
                        rhs8_ = ((d_63_v55_)[default__.safeIndex(816, (d_63_v55_).length(0))]) == (d_66_v58_)
                        lhs3_ = d_86_v73_
                        lhs4_ = default__.safeIndex(739, (d_86_v73_).length(0))
                        lhs5_ = d_53_v45_
                        lhs3_[lhs4_] = rhs6_
                        d_51_v43_ = rhs7_
                        lhs5_.f23 = rhs8_
                    pass
            pass
        d_87_i4_: int
        d_87_i4_ = 0
        with _dafny.label("1"):
            while p0:
                with _dafny.c_label("1"):
                    if (d_87_i4_) >= (100):
                        raise _dafny.Break("1")
                    d_87_i4_ = (d_87_i4_) + (1)
                    d_88_v74_: int
                    d_88_v74_ = 120
                    index12_ = default__.safeIndex(926, (d_48_v42_).length(0))
                    (d_48_v42_)[index12_] = d_88_v74_
                    index13_ = default__.safeIndex(926, (d_48_v42_).length(0))
                    (d_48_v42_)[index13_] = (0) - (default__.safeModuloInt(d_88_v74_, -138))
                    d_89_v75_: _dafny.Seq
                    d_89_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqxh"))
                    d_90_v76_: _dafny.Map
                    d_90_v76_ = _dafny.Map({p0: d_89_v75_})
                    d_91_v77_: _dafny.Map
                    d_91_v77_ = _dafny.Map({d_88_v74_: (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]})
                    d_92_v78_: str
                    d_92_v78_ = _dafny.CodePoint('h')
                    d_93_v79_: D5
                    d_93_v79_ = D5_DC10(d_92_v78_)
                    d_94_v80_: _dafny.Seq
                    d_94_v80_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_95_v81_: C0
                    nw19_ = C0()
                    nw19_.ctor__(p0, _dafny.Map({(d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]: d_94_v80_}))
                    d_95_v81_ = nw19_
                    d_96_v82_: _dafny.Seq
                    d_96_v82_ = _dafny.SeqWithoutIsStrInference([d_95_v81_])
                    d_97_v83_: _dafny.Array
                    nw20_ = _dafny.Array(None, 24)
                    nw20_[int(0)] = -478
                    nw20_[int(1)] = len(d_91_v77_)
                    nw20_[int(2)] = d_88_v74_
                    nw20_[int(3)] = len(_dafny.Map({p0: (d_93_v79_).cf19}))
                    nw20_[int(4)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    nw20_[int(5)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    nw20_[int(6)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    nw20_[int(7)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    nw20_[int(8)] = d_88_v74_
                    nw20_[int(9)] = d_88_v74_
                    nw20_[int(10)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    nw20_[int(11)] = len(d_96_v82_)
                    nw20_[int(12)] = d_88_v74_
                    nw20_[int(13)] = len(d_89_v75_)
                    nw20_[int(14)] = d_88_v74_
                    nw20_[int(15)] = d_88_v74_
                    nw20_[int(16)] = d_88_v74_
                    nw20_[int(17)] = d_88_v74_
                    nw20_[int(18)] = len(_dafny.Map({(d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]: p1}))
                    nw20_[int(19)] = default__.fm0(d_95_v81_.f24, p0, globalState)
                    nw20_[int(20)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    nw20_[int(21)] = d_88_v74_
                    nw20_[int(22)] = d_88_v74_
                    nw20_[int(23)] = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    d_97_v83_ = nw20_
                    d_98_v84_: _dafny.Map
                    d_98_v84_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no")): d_97_v83_})
                    d_99_v85_: _dafny.Set
                    d_99_v85_ = _dafny.Set({d_88_v74_, (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))], (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))], len(d_98_v84_), len(d_94_v80_)})
                    d_100_v86_: _dafny.Seq
                    d_100_v86_ = _dafny.SeqWithoutIsStrInference([len(d_99_v85_), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(d_89_v75_)[default__.safeIndex(293, len(d_89_v75_))] for d_101_i5_ in range(default__.abs(4))])), (0) - (len(d_89_v75_))}))])
                    d_102_v87_: T0
                    nw21_ = C0()
                    nw21_.ctor__(d_95_v81_.f24, (d_95_v81_).f25)
                    d_102_v87_ = nw21_
                    d_103_v88_: _dafny.Map
                    d_103_v88_ = _dafny.Map({(d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]: d_102_v87_})
                    index14_ = default__.safeIndex(926, (d_48_v42_).length(0))
                    (d_48_v42_)[index14_] = default__.safeDivisionInt((len(((d_90_v76_)[p0] if (p0) in (d_90_v76_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovddoibvv"))))) * (len((d_100_v86_).set(default__.safeIndex(len(d_103_v88_), len(d_100_v86_)), d_88_v74_))), (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))])
                    d_104_v89_: _dafny.Array
                    nw22_ = _dafny.Array(False, 1)
                    d_104_v89_ = nw22_
                    index15_ = default__.safeIndex(379, (d_104_v89_).length(0))
                    (d_104_v89_)[index15_] = not(p1)
                    index16_ = default__.safeIndex(379, (d_104_v89_).length(0))
                    (d_104_v89_)[index16_] = d_95_v81_.f24
                    index17_ = default__.safeIndex(926, (d_48_v42_).length(0))
                    rhs9_ = d_102_v87_
                    rhs10_ = d_95_v81_.f24
                    rhs11_ = (d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))]
                    rhs12_ = ((0) - (d_88_v74_)) + ((d_48_v42_)[default__.safeIndex(926, (d_48_v42_).length(0))])
                    lhs6_ = globalState
                    lhs7_ = d_48_v42_
                    lhs8_ = default__.safeIndex(926, (d_48_v42_).length(0))
                    d_102_v87_ = rhs9_
                    lhs6_.f19 = rhs10_
                    lhs7_[lhs8_] = rhs11_
                    d_88_v74_ = rhs12_
                    pass
            pass
        d_105_v90_: int
        d_105_v90_ = 568
        d_105_v90_ = default__.safeModuloInt(-273, d_105_v90_)
        d_106_v91_: _dafny.Seq
        d_106_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obf"))
        d_107_v92_: _dafny.Array
        nw23_ = _dafny.Array(False, 8)
        d_107_v92_ = nw23_
        d_108_v93_: _dafny.MultiSet
        d_108_v93_ = _dafny.MultiSet([p1])
        d_109_v94_: _dafny.Seq
        d_109_v94_ = _dafny.SeqWithoutIsStrInference([default__.fm11(d_108_v93_, d_105_v90_, globalState)])
        d_110_v95_: T0
        nw24_ = C0()
        nw24_.ctor__(p1, _dafny.Map({d_105_v90_: d_109_v94_}))
        d_110_v95_ = nw24_
        d_111_v96_: D5
        d_111_v96_ = D5_DC11(p1, p0, d_107_v92_, d_110_v95_)
        d_112_v97_: _dafny.Seq
        d_112_v97_ = _dafny.SeqWithoutIsStrInference([d_111_v96_, D5_DC11(p1, p0, d_107_v92_, d_110_v95_), d_111_v96_, D5_DC11(p0, p1, d_107_v92_, d_110_v95_)])
        d_113_v98_: _dafny.Map
        d_113_v98_ = _dafny.Map({d_106_v91_: d_112_v97_})
        d_114_v99_: _dafny.Array
        nw25_ = _dafny.Array(_dafny.Array(None, 0), 7)
        d_114_v99_ = nw25_
        index18_ = default__.safeIndex(278, (d_114_v99_).length(0))
        (d_114_v99_)[index18_] = d_48_v42_
        d_115_v100_: _dafny.Array
        d_115_v100_ = d_48_v42_
        index19_ = default__.safeIndex(278, (d_114_v99_).length(0))
        rhs13_ = (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_116_i6_ in range(default__.abs(423))]): d_112_v97_})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rntowef")): d_112_v97_}))
        rhs14_ = (d_115_v100_)
        lhs9_ = d_114_v99_
        lhs10_ = default__.safeIndex(278, (d_114_v99_).length(0))
        d_113_v98_ = rhs13_
        lhs9_[lhs10_] = rhs14_

    @staticmethod
    def Main(noArgsParameter__):
        d_117_v0_: bool
        d_117_v0_ = True
        d_118_v1_: _dafny.Seq
        d_118_v1_ = _dafny.SeqWithoutIsStrInference([d_117_v0_])
        d_119_v2_: _dafny.Seq
        d_119_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg"))
        d_120_v3_: D0
        d_120_v3_ = D0_DC1(d_118_v1_, len(d_119_v2_))
        d_121_v4_: _dafny.Map
        d_121_v4_ = _dafny.Map({d_117_v0_: (d_120_v3_).cf2})
        d_122_v5_: _dafny.Map
        d_122_v5_ = _dafny.Map({d_120_v3_: _dafny.MultiSet([d_117_v0_])})
        d_123_v6_: _dafny.Set
        d_123_v6_ = _dafny.Set({d_117_v0_})
        d_124_v7_: _dafny.MultiSet
        d_124_v7_ = _dafny.MultiSet([d_117_v0_])
        d_125_v8_: _dafny.Seq
        d_125_v8_ = _dafny.SeqWithoutIsStrInference([d_119_v2_, d_119_v2_, d_119_v2_])
        d_126_v9_: _dafny.Array
        def lambda2_(d_127_v2_):
            def lambda3_(d_128_i0_):
                return default__.safeModuloInt(d_128_i0_, len(d_127_v2_))

            return lambda3_

        init1_ = lambda2_(d_119_v2_)
        nw26_ = _dafny.Array(None, 19)
        for i0_1_ in range(nw26_.length(0)):
            nw26_[i0_1_] = init1_(i0_1_)
        d_126_v9_ = nw26_
        d_129_v11_: str
        d_129_v11_ = _dafny.CodePoint('j')
        d_130_v12_: _dafny.Set
        d_130_v12_ = _dafny.Set({d_129_v11_})
        d_131_v13_: _dafny.Map
        d_131_v13_ = _dafny.Map({d_124_v7_: d_117_v0_})
        d_132_v14_: int
        d_132_v14_ = 588
        d_133_v15_: _dafny.Map
        d_133_v15_ = _dafny.Map({d_129_v11_: (_dafny.Map({d_117_v0_: d_132_v14_})).set(d_117_v0_, d_132_v14_)})
        d_134_globalState_: GlobalState
        nw27_ = GlobalState()
        def iife6_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(729, 212):
                d_135_v10_: int = compr_2_
                if ((729) <= (d_135_v10_)) and ((d_135_v10_) < (212)):
                    coll2_[default__.safeDivisionInt(d_135_v10_, -701)] = _dafny.SeqWithoutIsStrInference([-6, (0) - (-208)])
            return _dafny.Map(coll2_)
        nw27_.ctor__(d_121_v4_, True, True, ((d_122_v5_)[D0_DC1(d_118_v1_, len(d_123_v6_))] if (D0_DC1(d_118_v1_, len(d_123_v6_))) in (d_122_v5_) else d_124_v7_), True, 869, False, 924, False, -924, (d_125_v8_) + (d_125_v8_), d_119_v2_, 380, d_126_v9_, 486, iife6_()
        , d_130_v12_, (d_131_v13_) | (d_131_v13_), _dafny.CodePoint('y'), False, d_133_v15_, 964)
        d_134_globalState_ = nw27_
        d_136_v16_: _dafny.Set
        d_136_v16_ = _dafny.Set({d_118_v1_, d_118_v1_, (d_118_v1_) + (d_118_v1_)})
        d_136_v16_ = (d_136_v16_) | ((d_136_v16_) | (d_136_v16_))
        d_137_v17_: _dafny.Map
        d_137_v17_ = _dafny.Map({(d_117_v0_ if d_117_v0_ else not(d_117_v0_)): _dafny.SeqWithoutIsStrInference([(0) - (d_132_v14_) for d_138_i1_ in range(default__.abs(243))])})
        d_139_v18_: _dafny.Seq
        d_139_v18_ = _dafny.SeqWithoutIsStrInference([d_124_v7_])
        d_140_v19_: _dafny.Seq
        d_140_v19_ = _dafny.SeqWithoutIsStrInference([d_132_v14_, ((d_139_v18_)[default__.safeIndex(d_132_v14_, len(d_139_v18_))]).cardinality])
        d_137_v17_ = (d_137_v17_).set((d_117_v0_) or (not(d_117_v0_)), d_140_v19_)
        d_141_v20_: _dafny.Array
        nw28_ = _dafny.Array(False, 27)
        d_141_v20_ = nw28_
        d_141_v20_ = d_141_v20_
        d_142_v21_: D0
        d_142_v21_ = D0_DC2(d_117_v0_, d_132_v14_, len(d_121_v4_), d_117_v0_)
        d_143_v22_: _dafny.Map
        d_143_v22_ = _dafny.Map({d_132_v14_: D0_DC3(d_142_v21_)})
        d_144_v23_: _dafny.Array
        nw29_ = _dafny.Array(_dafny.Map({}), 12)
        d_144_v23_ = nw29_
        d_145_v24_: _dafny.Map
        d_145_v24_ = _dafny.Map({_dafny.MultiSet([d_129_v11_]): default__.fm0(d_117_v0_, d_117_v0_, d_134_globalState_)})
        index20_ = default__.safeIndex(746, (d_144_v23_).length(0))
        (d_144_v23_)[index20_] = d_145_v24_
        index21_ = default__.safeIndex(562, (d_126_v9_).length(0))
        (d_126_v9_)[index21_] = d_132_v14_
        d_146_v25_: _dafny.Set
        d_146_v25_ = _dafny.Set({len(d_119_v2_)})
        d_147_v26_: _dafny.MultiSet
        d_147_v26_ = _dafny.MultiSet([d_129_v11_])
        index22_ = default__.safeIndex(746, (d_144_v23_).length(0))
        index23_ = default__.safeIndex(562, (d_126_v9_).length(0))
        rhs15_ = (d_119_v2_) + (d_119_v2_)
        rhs16_ = default__.fm1(d_117_v0_, d_117_v0_, default__.safeModuloInt(len(d_146_v25_), d_132_v14_), d_132_v14_, d_134_globalState_)
        rhs17_ = _dafny.Map({d_147_v26_: (d_132_v14_) + (d_132_v14_)})
        rhs18_ = d_123_v6_
        rhs19_ = (0) - ((default__.safeModuloInt(d_132_v14_, d_132_v14_)) - (d_132_v14_))
        lhs11_ = d_144_v23_
        lhs12_ = default__.safeIndex(746, (d_144_v23_).length(0))
        lhs13_ = d_126_v9_
        lhs14_ = default__.safeIndex(562, (d_126_v9_).length(0))
        d_119_v2_ = rhs15_
        d_143_v22_ = rhs16_
        lhs11_[lhs12_] = rhs17_
        d_123_v6_ = rhs18_
        lhs13_[lhs14_] = rhs19_
        d_140_v19_ = d_140_v19_
        d_119_v2_ = (d_119_v2_) + ((d_119_v2_).set(default__.safeIndex(d_132_v14_, len(d_119_v2_)), d_129_v11_))
        index24_ = default__.safeIndex(562, (d_126_v9_).length(0))
        rhs20_ = d_120_v3_
        rhs21_ = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
        lhs15_ = d_126_v9_
        lhs16_ = default__.safeIndex(562, (d_126_v9_).length(0))
        d_120_v3_ = rhs20_
        lhs15_[lhs16_] = rhs21_
        d_148_i2_: int
        d_148_i2_ = 0
        with _dafny.label("2"):
            while not(d_117_v0_):
                with _dafny.c_label("2"):
                    if (d_148_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_148_i2_ = (d_148_i2_) + (1)
                    if (_dafny.SeqWithoutIsStrInference([d_129_v11_ for d_149_i3_ in range(default__.abs(138))])) != ((d_119_v2_) + (d_119_v2_)):
                        d_140_v19_ = default__.fm2((_dafny.MultiSet(d_119_v2_)).intersection(d_147_v26_), d_134_globalState_)
                        d_119_v2_ = default__.fm3(d_132_v14_, d_134_globalState_)
                        index25_ = default__.safeIndex(698, (d_141_v20_).length(0))
                        (d_141_v20_)[index25_] = d_117_v0_
                        index26_ = default__.safeIndex(698, (d_141_v20_).length(0))
                        (d_141_v20_)[index26_] = d_117_v0_
                        d_150_v27_: _dafny.Map
                        d_150_v27_ = _dafny.Map({len(d_146_v25_): d_118_v1_})
                        d_151_v28_: C0
                        nw30_ = C0()
                        nw30_.ctor__(not(not((d_141_v20_)[default__.safeIndex(698, (d_141_v20_).length(0))])), d_150_v27_)
                        d_151_v28_ = nw30_
                        default__.m2(((d_118_v1_)[default__.safeIndex(d_132_v14_, len(d_118_v1_))] if d_151_v28_.f24 else d_117_v0_), ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]) < ((0) - ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))])), d_134_globalState_)
                    elif True:
                        d_132_v14_ = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
                        (d_134_globalState_).f19 = ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]) != (d_132_v14_)
                        d_152_v29_: _dafny.Map
                        d_152_v29_ = _dafny.Map({len(d_140_v19_): (d_118_v1_).set(default__.safeIndex(((d_121_v4_)[d_117_v0_] if (d_117_v0_) in (d_121_v4_) else d_132_v14_), len(d_118_v1_)), d_117_v0_)})
                        d_153_v30_: T0
                        nw31_ = C0()
                        nw31_.ctor__(d_117_v0_, d_152_v29_)
                        d_153_v30_ = nw31_
                        d_154_v31_: _dafny.Map
                        d_154_v31_ = _dafny.Map({d_129_v11_: d_153_v30_})
                        d_155_v32_: _dafny.Map
                        d_155_v32_ = _dafny.Map({759: d_153_v30_})
                        d_156_v33_: _dafny.Array
                        nw32_ = _dafny.Array(None, 20)
                        nw32_[int(0)] = d_153_v30_
                        nw32_[int(1)] = d_153_v30_
                        nw32_[int(2)] = d_153_v30_
                        nw32_[int(3)] = d_153_v30_
                        nw32_[int(4)] = d_153_v30_
                        nw32_[int(5)] = d_153_v30_
                        nw32_[int(6)] = d_153_v30_
                        nw32_[int(7)] = (((d_154_v31_)[d_129_v11_] if (d_129_v11_) in (d_154_v31_) else d_153_v30_) if d_117_v0_ else d_153_v30_)
                        nw32_[int(8)] = d_153_v30_
                        nw32_[int(9)] = d_153_v30_
                        nw32_[int(10)] = d_153_v30_
                        nw32_[int(11)] = d_153_v30_
                        nw32_[int(12)] = d_153_v30_
                        nw32_[int(13)] = (d_153_v30_ if not(d_117_v0_) else d_153_v30_)
                        nw32_[int(14)] = d_153_v30_
                        nw32_[int(15)] = d_153_v30_
                        nw32_[int(16)] = d_153_v30_
                        nw32_[int(17)] = ((d_155_v32_)[(d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]] if ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]) in (d_155_v32_) else d_153_v30_)
                        nw32_[int(18)] = d_153_v30_
                        nw32_[int(19)] = d_153_v30_
                        d_156_v33_ = nw32_
                        index27_ = default__.safeIndex(48, (d_156_v33_).length(0))
                        (d_156_v33_)[index27_] = d_153_v30_
                        index28_ = default__.safeIndex(48, (d_156_v33_).length(0))
                        nw33_ = C0()
                        nw33_.ctor__(not(d_117_v0_), (d_152_v29_).set(len(d_118_v1_), d_118_v1_))
                        (d_156_v33_)[index28_] = nw33_
                        default__.m2(d_117_v0_, not(default__.fm11(d_124_v7_, len(_dafny.SeqWithoutIsStrInference([(d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]])), d_134_globalState_)), d_134_globalState_)
                        d_132_v14_ = (d_153_v30_).fm8(d_134_globalState_)
                    index29_ = default__.safeIndex(28, (d_141_v20_).length(0))
                    (d_141_v20_)[index29_] = not(default__.fm11(d_124_v7_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvyhg"))), d_134_globalState_))
                    index30_ = default__.safeIndex(28, (d_141_v20_).length(0))
                    (d_141_v20_)[index30_] = default__.fm11((_dafny.MultiSet([default__.fm11(_dafny.MultiSet([d_117_v0_, d_117_v0_]), (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))], d_134_globalState_), d_117_v0_])) | (d_124_v7_), d_132_v14_, d_134_globalState_)
                    if not(not((d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))])):
                        d_157_v34_: D0
                        d_157_v34_ = D0_DC2(d_117_v0_, d_132_v14_, (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))], d_117_v0_)
                        d_158_v35_: _dafny.Map
                        d_158_v35_ = _dafny.Map({(d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]: (d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))]})
                        d_159_v36_: _dafny.Array
                        nw34_ = _dafny.Array(None, 23)
                        nw34_[int(0)] = (d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))]
                        nw34_[int(1)] = (d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))]
                        nw34_[int(2)] = (d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))]
                        nw34_[int(3)] = d_117_v0_
                        nw34_[int(4)] = False
                        nw34_[int(5)] = True
                        nw34_[int(6)] = d_117_v0_
                        nw34_[int(7)] = d_117_v0_
                        nw34_[int(8)] = d_117_v0_
                        nw34_[int(9)] = d_117_v0_
                        nw34_[int(10)] = (d_157_v34_).cf6
                        nw34_[int(11)] = True
                        nw34_[int(12)] = d_117_v0_
                        nw34_[int(13)] = True
                        nw34_[int(14)] = (d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))]
                        nw34_[int(15)] = d_117_v0_
                        nw34_[int(16)] = d_117_v0_
                        nw34_[int(17)] = True
                        nw34_[int(18)] = not(False)
                        nw34_[int(19)] = (d_157_v34_).cf3
                        nw34_[int(20)] = not(d_117_v0_)
                        nw34_[int(21)] = d_117_v0_
                        nw34_[int(22)] = ((d_158_v35_)[293] if (293) in (d_158_v35_) else d_117_v0_)
                        d_159_v36_ = nw34_
                        d_160_v37_: _dafny.Map
                        d_160_v37_ = _dafny.Map({d_159_v36_: False})
                        d_160_v37_ = (d_160_v37_).set(d_159_v36_, ((d_158_v35_)[788] if (788) in (d_158_v35_) else d_117_v0_))
                        d_161_v38_: D3
                        d_161_v38_ = D3_DC8(d_132_v14_, d_132_v14_)
                        pat_let_tv1_ = d_126_v9_
                        pat_let_tv2_ = d_126_v9_
                        def iife7_(_pat_let2_0):
                            def iife8_(d_162_dt__update__tmp_h0_):
                                def iife9_(_pat_let3_0):
                                    def iife10_(d_163_dt__update_hcf17_h0_):
                                        return D3_DC8((d_162_dt__update__tmp_h0_).cf16, d_163_dt__update_hcf17_h0_)
                                    return iife10_(_pat_let3_0)
                                return iife9_((pat_let_tv2_)[default__.safeIndex(562, (pat_let_tv1_).length(0))])
                            return iife8_(_pat_let2_0)
                        d_132_v14_ = (default__.safeModuloInt((0) - ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]), (iife7_(d_161_v38_)).cf16) if (d_118_v1_)[default__.safeIndex(d_132_v14_, len(d_118_v1_))] else 803)
                        (d_134_globalState_).f6 = ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]) < (d_132_v14_)
                        d_164_v39_: _dafny.Seq
                        d_164_v39_ = _dafny.SeqWithoutIsStrInference([d_159_v36_])
                        d_159_v36_ = (d_164_v39_)[default__.safeIndex(default__.safeDivisionInt(d_132_v14_, (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]), len(d_164_v39_))]
                        d_165_v40_: C1
                        nw35_ = C1()
                        nw35_.ctor__(default__.fm11(d_124_v7_, len(d_118_v1_), d_134_globalState_), d_117_v0_)
                        d_165_v40_ = nw35_
                    elif True:
                        default__.m2(((d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))] if True else d_117_v0_), True, d_134_globalState_)
                        (d_134_globalState_).f19 = d_117_v0_
                        index31_ = default__.safeIndex(28, (d_141_v20_).length(0))
                        (d_141_v20_)[index31_] = default__.fm11(_dafny.MultiSet([(d_141_v20_)[default__.safeIndex(28, (d_141_v20_).length(0))], d_117_v0_, d_117_v0_]), d_132_v14_, d_134_globalState_)
                        default__.m2(d_117_v0_, d_117_v0_, d_134_globalState_)
                        index32_ = default__.safeIndex(28, (d_141_v20_).length(0))
                        (d_141_v20_)[index32_] = d_117_v0_
                    d_132_v14_ = (default__.safeModuloInt(len(_dafny.Map({d_132_v14_: (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]})), len(d_119_v2_))) - (len((d_123_v6_) | (d_123_v6_)))
                    pass
            pass
        d_166_v41_: _dafny.Map
        d_166_v41_ = _dafny.Map({(d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]: _dafny.SeqWithoutIsStrInference([d_117_v0_])})
        d_167_v42_: T0
        nw36_ = C0()
        nw36_.ctor__(True, d_166_v41_)
        d_167_v42_ = nw36_
        d_168_v43_: _dafny.Seq
        d_168_v43_ = _dafny.SeqWithoutIsStrInference([d_167_v42_, d_167_v42_, d_167_v42_])
        (d_134_globalState_).f19 = (d_168_v43_) < (d_168_v43_)
        default__.m2(d_117_v0_, d_117_v0_, d_134_globalState_)
        d_169_v44_: C1
        nw37_ = C1()
        nw37_.ctor__((d_117_v0_ if d_117_v0_ else d_117_v0_), d_117_v0_)
        d_169_v44_ = nw37_
        d_169_v44_ = d_169_v44_
        d_170_v45_: _dafny.Array
        nw38_ = _dafny.Array(None, 11)
        nw38_[int(0)] = d_126_v9_
        nw38_[int(1)] = d_126_v9_
        nw38_[int(2)] = d_126_v9_
        nw38_[int(3)] = d_126_v9_
        nw38_[int(4)] = d_126_v9_
        nw38_[int(5)] = d_126_v9_
        nw38_[int(6)] = d_126_v9_
        nw38_[int(7)] = d_126_v9_
        nw38_[int(8)] = d_126_v9_
        nw38_[int(9)] = d_126_v9_
        nw38_[int(10)] = d_126_v9_
        d_170_v45_ = nw38_
        index33_ = default__.safeIndex(408, (d_170_v45_).length(0))
        (d_170_v45_)[index33_] = d_126_v9_
        index34_ = default__.safeIndex(408, (d_170_v45_).length(0))
        (d_170_v45_)[index34_] = d_126_v9_
        d_132_v14_ = ((0) - (d_132_v14_) if (len(default__.fm3(((d_121_v4_)[(d_169_v44_).fm4(d_132_v14_, (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))], d_169_v44_.f23, d_134_globalState_)] if ((d_169_v44_).fm4(d_132_v14_, (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))], d_169_v44_.f23, d_134_globalState_)) in (d_121_v4_) else (0) - (d_132_v14_)), d_134_globalState_))) < (278) else default__.fm0(d_169_v44_.f23, d_117_v0_, d_134_globalState_))
        hi0_ = (0) - (default__.safeDivisionInt(d_132_v14_, (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]))
        for d_171_i4_ in range((d_132_v14_) - ((0) - ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))])), hi0_):
            d_132_v14_ = 673
            index35_ = default__.safeIndex(562, (d_126_v9_).length(0))
            index36_ = default__.safeIndex(562, (d_126_v9_).length(0))
            rhs22_ = ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]) != ((d_171_i4_) - (d_171_i4_))
            rhs23_ = d_126_v9_
            rhs24_ = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            rhs25_ = d_171_i4_
            lhs17_ = d_126_v9_
            lhs18_ = default__.safeIndex(562, (d_126_v9_).length(0))
            lhs19_ = d_126_v9_
            lhs20_ = default__.safeIndex(562, (d_126_v9_).length(0))
            d_117_v0_ = rhs22_
            d_126_v9_ = rhs23_
            lhs17_[lhs18_] = rhs24_
            lhs19_[lhs20_] = rhs25_
            (d_134_globalState_).f18 = _dafny.CodePoint('e')
            d_117_v0_ = not(not (((0) - ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))])) >= ((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))])) or (d_117_v0_))
        d_172_v47_: C0
        nw39_ = C0()
        def iife11_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([len(d_140_v19_) for d_173_i6_ in range(default__.abs(294))])).Elements:
                d_174_v46_: int = compr_3_
                if (d_174_v46_) in (_dafny.SeqWithoutIsStrInference([len(d_140_v19_) for d_173_i6_ in range(default__.abs(294))])):
                    coll3_[(d_174_v46_) + (len(d_121_v4_))] = d_118_v1_
            return _dafny.Map(coll3_)
        nw39_.ctor__(d_169_v44_.f23, iife11_()
        )
        d_172_v47_ = nw39_
        d_175_v48_: _dafny.Map
        d_175_v48_ = _dafny.Map({d_172_v47_: (d_169_v44_).f22})
        hi1_ = len(d_140_v19_)
        for d_176_i5_ in range(len(d_175_v48_), hi1_):
            d_177_v49_: _dafny.Seq
            d_178_v50_: bool
            out0_: _dafny.Seq
            out1_: bool
            out0_, out1_ = (d_169_v44_).m1(default__.safeDivisionInt(len(d_119_v2_), d_176_i5_), d_172_v47_.f24, (d_169_v44_).f22, len((d_119_v2_) + (d_119_v2_)), d_134_globalState_)
            d_177_v49_ = out0_
            d_178_v50_ = out1_
            d_179_v51_: bool
            d_180_v52_: _dafny.Map
            out2_: bool
            out3_: _dafny.Map
            out2_, out3_ = (d_169_v44_).m0(313, d_134_globalState_)
            d_179_v51_ = out2_
            d_180_v52_ = out3_
            d_181_v53_: _dafny.Map
            d_181_v53_ = _dafny.Map({len(d_140_v19_): d_141_v20_})
            d_182_v54_: _dafny.Seq
            d_182_v54_ = _dafny.SeqWithoutIsStrInference([d_141_v20_, d_141_v20_])
            d_183_v55_: _dafny.MultiSet
            d_183_v55_ = _dafny.MultiSet([d_141_v20_, d_141_v20_, ((d_181_v53_)[d_176_i5_] if (d_176_i5_) in (d_181_v53_) else (d_182_v54_)[default__.safeIndex(d_132_v14_, len(d_182_v54_))]), d_141_v20_, d_141_v20_])
            d_184_v56_: _dafny.MultiSet
            d_184_v56_ = _dafny.MultiSet([d_141_v20_])
            d_185_v57_: _dafny.Seq
            d_185_v57_ = _dafny.SeqWithoutIsStrInference([(d_184_v56_)])
            d_183_v55_ = (d_185_v57_)[default__.safeIndex(default__.safeModuloInt(d_176_i5_, 508), len(d_185_v57_))]
            d_124_v7_ = d_124_v7_
        if (d_169_v44_).f22:
            d_186_v58_: _dafny.Map
            d_186_v58_ = _dafny.Map({(d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]: len(d_146_v25_)})
            d_187_v59_: _dafny.Map
            d_187_v59_ = d_186_v58_
            d_187_v59_ = default__.fm12(d_134_globalState_)
            d_188_v60_: _dafny.Array
            d_188_v60_ = (d_170_v45_)[default__.safeIndex(408, (d_170_v45_).length(0))]
            d_189_v61_: _dafny.Set
            d_189_v61_ = _dafny.Set({d_188_v60_, d_188_v60_})
            d_190_v62_: _dafny.Map
            d_190_v62_ = _dafny.Map({d_172_v47_: d_189_v61_})
            (d_169_v44_).f23 = ((d_189_v61_) - (d_189_v61_)).isdisjoint(((d_190_v62_)[d_172_v47_] if (d_172_v47_) in (d_190_v62_) else d_189_v61_))
            d_132_v14_ = ((d_132_v14_) * (52)) + (len((_dafny.Map({d_172_v47_: d_169_v44_.f23})) | (d_175_v48_)))
            d_117_v0_ = default__.fm11(d_124_v7_, d_132_v14_, d_134_globalState_)
            (d_134_globalState_).f19 = (d_169_v44_).f22
        elif True:
            d_191_v63_: _dafny.MultiSet
            d_191_v63_ = _dafny.MultiSet([d_141_v20_, d_141_v20_])
            d_192_v64_: _dafny.MultiSet
            d_192_v64_ = (d_191_v63_) | (d_191_v63_)
            source1_ = d_192_v64_
            d_193___mcc_h0_ = source1_
            d_194_cf18_ = d_193___mcc_h0_
            d_195_v65_: _dafny.Map
            d_195_v65_ = _dafny.Map({_dafny.CodePoint('o'): 938})
            d_196_v66_: _dafny.Map
            d_196_v66_ = _dafny.Map({len(d_195_v65_): (d_170_v45_)[default__.safeIndex(408, (d_170_v45_).length(0))]})
            d_196_v66_ = (d_196_v66_).set(751, (d_170_v45_)[default__.safeIndex(408, (d_170_v45_).length(0))])
            nw40_ = _dafny.Array(None, 14)
            nw40_[int(0)] = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            nw40_[int(1)] = default__.safeModuloInt((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))], d_132_v14_)
            nw40_[int(2)] = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            nw40_[int(3)] = default__.fm0(not(True), d_169_v44_.f23, d_134_globalState_)
            nw40_[int(4)] = (338) * (d_132_v14_)
            nw40_[int(5)] = ((d_124_v7_)[d_169_v44_.f23] if (d_169_v44_.f23) in (d_124_v7_) else d_132_v14_)
            nw40_[int(6)] = ((0) - (d_132_v14_)) + (len(d_119_v2_))
            nw40_[int(7)] = d_132_v14_
            nw40_[int(8)] = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            nw40_[int(9)] = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            nw40_[int(10)] = default__.safeModuloInt((d_140_v19_)[default__.safeIndex((d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))], len(d_140_v19_))], d_132_v14_)
            nw40_[int(11)] = (0) - (d_132_v14_)
            nw40_[int(12)] = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            nw40_[int(13)] = (d_126_v9_)[default__.safeIndex(562, (d_126_v9_).length(0))]
            d_126_v9_ = nw40_
            index37_ = default__.safeIndex(408, (d_170_v45_).length(0))
            def lambda4_(d_197_v14_):
                def lambda5_(d_198_i7_):
                    return (d_198_i7_) + (d_197_v14_)

                return lambda5_

            init2_ = lambda4_(d_132_v14_)
            nw41_ = _dafny.Array(None, 6)
            for i0_2_ in range(nw41_.length(0)):
                nw41_[i0_2_] = init2_(i0_2_)
            (d_170_v45_)[index37_] = nw41_
            d_199_v67_: _dafny.Set
            d_199_v67_ = _dafny.Set({d_169_v44_})
            d_199_v67_ = d_199_v67_
            d_200_v68_: C0
            nw42_ = C0()
            nw42_.ctor__(d_172_v47_.f24, (d_172_v47_).f25)
            d_200_v68_ = nw42_
            (d_172_v47_).f24 = (default__.safeDivisionInt((d_167_v42_).fm8(d_134_globalState_), -823)) == (d_132_v14_)
            (d_172_v47_).f24 = d_169_v44_.f23
            d_201_v69_: bool
            d_202_v70_: _dafny.Map
            out4_: bool
            out5_: _dafny.Map
            out4_, out5_ = (d_169_v44_).m0(len(d_140_v19_), d_134_globalState_)
            d_201_v69_ = out4_
            d_202_v70_ = out5_
        _dafny.print(_dafny.string_of(d_117_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v1_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_119_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_120_v3_).cf1) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v3_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v4_) == (_dafny.Map({True: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v5_) == (_dafny.Map({D0_DC1(_dafny.SeqWithoutIsStrInference([True]), 9): _dafny.MultiSet([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v6_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v7_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v8_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v9_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_129_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v12_) == (_dafny.Set({_dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v13_) == (_dafny.Map({_dafny.MultiSet([True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_v14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v15_) == (_dafny.Map({_dafny.CodePoint('j'): _dafny.Map({True: 588})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f0) == (_dafny.Map({True: 36}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f3) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f10) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvnnsysbg"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_134_globalState_).f11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_.f13)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f15) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f16) == (_dafny.Set({_dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f17) == (_dafny.Map({_dafny.MultiSet([True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_134_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_134_globalState_).f20) == (_dafny.Map({_dafny.CodePoint('j'): _dafny.Map({True: 588})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v16_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v17_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([588, 1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v18_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v19_) == (_dafny.SeqWithoutIsStrInference([588, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v21_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v21_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v21_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v21_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v22_) == (_dafny.Map({944: D0_DC3(D0_DC2(False, 754, 561, False)), 1: D0_DC3(D0_DC3(D0_DC2(True, 111, -539, False)))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v23_)[2]) == (_dafny.Map({_dafny.MultiSet([_dafny.CodePoint('j')]): 1176}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_v24_) == (_dafny.Map({_dafny.MultiSet([_dafny.CodePoint('j')]): 991}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_v25_) == (_dafny.Set({9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v26_) == (_dafny.MultiSet([_dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_148_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v41_) == (_dafny.Map({588: _dafny.SeqWithoutIsStrInference([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_168_v43_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_v44_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_169_v44_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[0])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[1])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[2])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[3])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[4])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[5])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[6])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[7])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[8])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[9])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v45_)[10])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_172_v47_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v47_).f25) == (_dafny.Map({3: _dafny.SeqWithoutIsStrInference([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_175_v48_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Seq({}), int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf3', Any), ('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(_dafny.Map({}), False, False, False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({self.cf10.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC9(D4, NamedTuple('DC9', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC11(False, False, _dafny.Array(None, 0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC11(D5, NamedTuple('DC11', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC10(D5, NamedTuple('DC10', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Map = _dafny.Map({})
        self.f6: bool = False
        self.f13: _dafny.Array = _dafny.Array(None, 0)
        self.f18: str = _dafny.CodePoint('D')
        self.f19: bool = False
        self._f1: bool = False
        self._f2: bool = False
        self._f3: _dafny.MultiSet = _dafny.MultiSet({})
        self._f4: bool = False
        self._f5: int = int(0)
        self._f7: int = int(0)
        self._f8: bool = False
        self._f9: int = int(0)
        self._f10: _dafny.Seq = _dafny.Seq({})
        self._f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: _dafny.Map = _dafny.Map({})
        self._f16: _dafny.Set = _dafny.Set({})
        self._f17: _dafny.Map = _dafny.Map({})
        self._f20: _dafny.Map = _dafny.Map({})
        self._f21: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21

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
    @property
    def f5(self):
        return self._f5
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
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C0(T0):
    def  __init__(self):
        self.f24: bool = False
        self._f25: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f24, f25):
        (self).f24 = f24
        (self)._f25 = f25

    def fm8(self, globalState):
        return 694

    def fm9(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('h')

    @property
    def f25(self):
        return self._f25

class C1:
    def  __init__(self):
        self.f23: bool = False
        self._f22: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f22, f23):
        (self)._f22 = f22
        (self).f23 = f23

    def fm4(self, p0, p1, p2, globalState):
        return (self).f22

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_203_v0_: str
        d_203_v0_ = _dafny.CodePoint('a')
        d_204_v1_: _dafny.Map
        d_204_v1_ = _dafny.Map({d_203_v0_: False})
        d_205_v2_: _dafny.Seq
        d_205_v2_ = _dafny.SeqWithoutIsStrInference([d_204_v1_])
        d_205_v2_ = d_205_v2_
        d_206_v3_: _dafny.Seq
        d_206_v3_ = _dafny.SeqWithoutIsStrInference([(self).f22])
        d_207_v4_: _dafny.Seq
        d_207_v4_ = _dafny.SeqWithoutIsStrInference([d_206_v3_, d_206_v3_])
        d_208_v5_: D0
        d_208_v5_ = D0_DC1((d_207_v4_)[default__.safeIndex(p0, len(d_207_v4_))], p0)
        pat_let_tv3_ = p0
        def iife12_(_pat_let4_0):
            def iife13_(d_209_dt__update__tmp_h0_):
                def iife14_(_pat_let5_0):
                    def iife15_(d_210_dt__update_hcf2_h0_):
                        return D0_DC1((d_209_dt__update__tmp_h0_).cf1, d_210_dt__update_hcf2_h0_)
                    return iife15_(_pat_let5_0)
                return iife14_(pat_let_tv3_)
            return iife13_(_pat_let4_0)
        source2_ = iife12_(d_208_v5_)
        if source2_.is_DC1:
            d_211___mcc_h0_ = source2_.cf1
            d_212___mcc_h1_ = source2_.cf2
            d_213_cf2_ = d_212___mcc_h1_
            d_214_cf1_ = d_211___mcc_h0_
            d_215_v6_: _dafny.Seq
            d_215_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgrcsul"))
            d_216_v7_: _dafny.Map
            d_216_v7_ = _dafny.Map({d_215_v6_: (0) - (default__.safeDivisionInt(d_213_cf2_, len(d_214_cf1_)))})
            d_216_v7_ = (d_216_v7_).set((d_215_v6_) + (d_215_v6_), (0) - (d_213_cf2_))
            d_217_v8_: _dafny.Array
            def lambda6_(d_218_cf2_):
                def lambda7_(d_219_i0_):
                    return default__.safeModuloInt(d_219_i0_, d_218_cf2_)

                return lambda7_

            init3_ = lambda6_(d_213_cf2_)
            nw43_ = _dafny.Array(None, 8)
            for i0_3_ in range(nw43_.length(0)):
                nw43_[i0_3_] = init3_(i0_3_)
            d_217_v8_ = nw43_
            d_220_v9_: _dafny.MultiSet
            d_220_v9_ = _dafny.MultiSet([d_217_v8_, d_217_v8_])
            if ((d_220_v9_).ispropersubset(d_220_v9_)) and ((d_206_v3_) <= (_dafny.SeqWithoutIsStrInference([(self).fm4(-705, d_213_cf2_, self.f23, globalState)]))):
                (globalState).f19 = not(not ((self).f22) or ((self).f22))
                rhs26_ = d_217_v8_
                rhs27_ = ((d_216_v7_)[d_215_v6_] if (d_215_v6_) in (d_216_v7_) else d_213_cf2_)
                lhs21_ = globalState
                lhs21_.f13 = rhs26_
                d_213_cf2_ = rhs27_
                d_213_cf2_ = (0) - (default__.safeModuloInt(default__.safeModuloInt(d_213_cf2_, d_213_cf2_), p0))
                (globalState).f18 = d_203_v0_
                d_215_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upfp"))
            elif True:
                (self).f23 = (self).f22
                (globalState).f6 = not(False)
                d_221_v10_: _dafny.Map
                d_221_v10_ = _dafny.Map({(self).f22: d_214_cf1_})
                d_222_v11_: _dafny.Set
                d_222_v11_ = _dafny.Set({d_203_v0_, d_203_v0_, default__.fm5(globalState), d_203_v0_})
                d_223_v12_: _dafny.Seq
                d_223_v12_ = _dafny.SeqWithoutIsStrInference([d_222_v11_, d_222_v11_])
                rhs28_ = (0) - (len((((d_206_v3_) + (((d_221_v10_)[self.f23] if (self.f23) in (d_221_v10_) else d_214_cf1_))).set(default__.safeIndex(len(d_223_v12_), len((d_206_v3_) + (((d_221_v10_)[self.f23] if (self.f23) in (d_221_v10_) else d_214_cf1_)))), self.f23)) + (_dafny.SeqWithoutIsStrInference([self.f23, self.f23]))))
                rhs29_ = self.f23
                lhs22_ = globalState
                d_213_cf2_ = rhs28_
                lhs22_.f19 = rhs29_
                d_224_v13_: _dafny.MultiSet
                d_224_v13_ = _dafny.MultiSet([False])
                d_225_v14_: _dafny.Seq
                d_225_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtobwgfoy"))])
                d_226_v15_: _dafny.Seq
                d_226_v15_ = _dafny.SeqWithoutIsStrInference([d_224_v13_, ((d_224_v13_).set((self).f22, default__.abs(len(d_225_v14_)))) | (d_224_v13_), _dafny.MultiSet([True, (self).f22])])
                d_227_v16_: _dafny.Set
                d_227_v16_ = _dafny.Set({d_213_cf2_})
                rhs30_ = ((d_220_v9_).issubset(d_220_v9_) if (d_213_cf2_) in (d_227_v16_) else (not((self).f22) if self.f23 else self.f23))
                rhs31_ = (d_226_v15_) + (_dafny.SeqWithoutIsStrInference([d_224_v13_ for d_228_i1_ in range(default__.abs(632))]))
                lhs23_ = globalState
                lhs23_.f6 = rhs30_
                d_226_v15_ = rhs31_
                d_229_v17_: _dafny.Array
                def lambda8_(d_230_i2_):
                    return ((self).f22 if (self).f22 else (self).f22)

                init4_ = lambda8_
                nw44_ = _dafny.Array(None, 12)
                for i0_4_ in range(nw44_.length(0)):
                    nw44_[i0_4_] = init4_(i0_4_)
                d_229_v17_ = nw44_
                d_229_v17_ = d_229_v17_
            d_231_v18_: _dafny.Array
            nw45_ = _dafny.Array(False, 24)
            d_231_v18_ = nw45_
            index38_ = default__.safeIndex(192, (d_231_v18_).length(0))
            (d_231_v18_)[index38_] = self.f23
            d_232_v19_: _dafny.Set
            d_232_v19_ = _dafny.Set({d_213_cf2_, p0})
            index39_ = default__.safeIndex(192, (d_231_v18_).length(0))
            rhs32_ = len(d_215_v6_)
            rhs33_ = (d_206_v3_)[default__.safeIndex(len(d_232_v19_), len(d_206_v3_))]
            rhs34_ = self.f23
            rhs35_ = d_206_v3_
            lhs24_ = globalState
            lhs25_ = d_231_v18_
            lhs26_ = default__.safeIndex(192, (d_231_v18_).length(0))
            d_213_cf2_ = rhs32_
            lhs24_.f19 = rhs33_
            lhs25_[lhs26_] = rhs34_
            d_206_v3_ = rhs35_
            d_233_v20_: _dafny.Seq
            d_233_v20_ = _dafny.SeqWithoutIsStrInference([p0])
            index40_ = default__.safeIndex(782, (d_217_v8_).length(0))
            (d_217_v8_)[index40_] = len(d_233_v20_)
            index41_ = default__.safeIndex(192, (d_231_v18_).length(0))
            index42_ = default__.safeIndex(782, (d_217_v8_).length(0))
            rhs36_ = (d_231_v18_)[default__.safeIndex(192, (d_231_v18_).length(0))]
            rhs37_ = (0) - (d_213_cf2_)
            rhs38_ = ((self).f22) in (((d_214_cf1_) + (d_214_cf1_)).set(default__.safeIndex(len(d_215_v6_), len((d_214_cf1_) + (d_214_cf1_))), self.f23))
            lhs27_ = d_231_v18_
            lhs28_ = default__.safeIndex(192, (d_231_v18_).length(0))
            lhs29_ = d_217_v8_
            lhs30_ = default__.safeIndex(782, (d_217_v8_).length(0))
            lhs31_ = globalState
            lhs27_[lhs28_] = rhs36_
            lhs29_[lhs30_] = rhs37_
            lhs31_.f19 = rhs38_
        elif source2_.is_DC2:
            d_234___mcc_h2_ = source2_.cf3
            d_235___mcc_h3_ = source2_.cf4
            d_236___mcc_h4_ = source2_.cf5
            d_237___mcc_h5_ = source2_.cf6
            d_238_cf6_ = d_237___mcc_h5_
            d_239_cf5_ = d_236___mcc_h4_
            d_240_cf4_ = d_235___mcc_h3_
            d_241_cf3_ = d_234___mcc_h2_
            if self.f23:
                d_242_v21_: _dafny.Array
                def lambda9_(d_243_cf4_):
                    def lambda10_(d_244_i3_):
                        return (d_244_i3_) + (d_243_cf4_)

                    return lambda10_

                init5_ = lambda9_(d_240_cf4_)
                nw46_ = _dafny.Array(None, 9)
                for i0_5_ in range(nw46_.length(0)):
                    nw46_[i0_5_] = init5_(i0_5_)
                d_242_v21_ = nw46_
                d_245_v22_: _dafny.Map
                d_245_v22_ = _dafny.Map({self.f23: d_242_v21_})
                d_245_v22_ = (d_245_v22_).set(d_238_cf6_, d_242_v21_)
                d_246_v23_: _dafny.Array
                nw47_ = _dafny.Array(_dafny.Map({}), 4)
                d_246_v23_ = nw47_
                d_247_v24_: _dafny.Map
                d_247_v24_ = _dafny.Map({d_238_cf6_: p0})
                index43_ = default__.safeIndex(962, (d_246_v23_).length(0))
                (d_246_v23_)[index43_] = d_247_v24_
                d_248_v25_: _dafny.Seq
                d_248_v25_ = _dafny.SeqWithoutIsStrInference([d_247_v24_])
                d_249_v26_: _dafny.Seq
                d_249_v26_ = _dafny.SeqWithoutIsStrInference([len(d_247_v24_)])
                d_250_v27_: _dafny.Seq
                d_250_v27_ = _dafny.SeqWithoutIsStrInference([(d_249_v26_)[default__.safeIndex(d_240_cf4_, len(d_249_v26_))]])
                index44_ = default__.safeIndex(962, (d_246_v23_).length(0))
                (d_246_v23_)[index44_] = ((d_247_v24_) | (d_247_v24_)) | ((d_248_v25_)[default__.safeIndex(len(d_250_v27_), len(d_248_v25_))])
                d_251_v28_: _dafny.Seq
                d_252_v29_: bool
                out6_: _dafny.Seq
                out7_: bool
                out6_, out7_ = (self).m1(len((d_246_v23_)[default__.safeIndex(962, (d_246_v23_).length(0))]), self.f23, (self).f22, d_240_cf4_, globalState)
                d_251_v28_ = out6_
                d_252_v29_ = out7_
                d_240_cf4_ = d_240_cf4_
                d_240_cf4_ = 669
            elif True:
                d_239_cf5_ = p0
                d_253_v30_: _dafny.Array
                nw48_ = _dafny.Array(int(0), 6)
                d_253_v30_ = nw48_
                index45_ = default__.safeIndex(828, (d_253_v30_).length(0))
                (d_253_v30_)[index45_] = (0) - ((d_239_cf5_) + (d_239_cf5_))
                d_254_v31_: _dafny.Map
                d_254_v31_ = _dafny.Map({d_238_cf6_: 450})
                d_255_v32_: _dafny.Seq
                d_255_v32_ = _dafny.SeqWithoutIsStrInference([(0) - (d_239_cf5_), p0])
                d_256_v33_: _dafny.Map
                d_256_v33_ = _dafny.Map({d_254_v31_: (self).fm4(-224, (d_255_v32_)[default__.safeIndex(d_240_cf4_, len(d_255_v32_))], d_238_cf6_, globalState)})
                d_257_v34_: _dafny.Map
                d_257_v34_ = _dafny.Map({d_240_cf4_: d_241_cf3_})
                index46_ = default__.safeIndex(828, (d_253_v30_).length(0))
                rhs39_ = d_240_cf4_
                rhs40_ = len(d_256_v33_)
                rhs41_ = (len(d_257_v34_)) + ((0) - (d_239_cf5_))
                rhs42_ = (274) * (len(d_255_v32_))
                rhs43_ = not(d_238_cf6_)
                lhs32_ = d_253_v30_
                lhs33_ = default__.safeIndex(828, (d_253_v30_).length(0))
                lhs34_ = globalState
                d_239_cf5_ = rhs39_
                d_239_cf5_ = rhs40_
                lhs32_[lhs33_] = rhs41_
                d_239_cf5_ = rhs42_
                lhs34_.f19 = rhs43_
                d_258_v35_: _dafny.MultiSet
                d_258_v35_ = _dafny.MultiSet([d_238_cf6_])
                d_259_v36_: D0
                d_259_v36_ = D0_DC2((self).f22, p0, (d_258_v35_).cardinality, d_238_cf6_)
                d_260_v37_: _dafny.Map
                d_260_v37_ = _dafny.Map({d_259_v36_: (self).fm4(d_239_cf5_, (d_253_v30_)[default__.safeIndex(828, (d_253_v30_).length(0))], d_238_cf6_, globalState)})
                d_260_v37_ = (d_260_v37_).set(d_259_v36_, (self).f22)
                d_261_v38_: _dafny.Seq
                d_261_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvttpt"))
                d_262_v39_: _dafny.Array
                nw49_ = _dafny.Array(None, 13)
                nw49_[int(0)] = (p0) == ((d_255_v32_)[default__.safeIndex((d_253_v30_)[default__.safeIndex(828, (d_253_v30_).length(0))], len(d_255_v32_))])
                nw49_[int(1)] = d_238_cf6_
                nw49_[int(2)] = (d_241_cf3_) == (d_238_cf6_)
                nw49_[int(3)] = (d_261_v38_) < (d_261_v38_)
                nw49_[int(4)] = not(d_238_cf6_)
                nw49_[int(5)] = (self).f22
                nw49_[int(6)] = (self).fm4(d_239_cf5_, default__.fm0(not(not(d_238_cf6_)), self.f23, globalState), False, globalState)
                nw49_[int(7)] = (d_206_v3_) == (d_206_v3_)
                nw49_[int(8)] = self.f23
                nw49_[int(9)] = (self).f22
                nw49_[int(10)] = d_241_cf3_
                nw49_[int(11)] = self.f23
                nw49_[int(12)] = self.f23
                d_262_v39_ = nw49_
                d_263_v40_: D0
                d_263_v40_ = D0_DC0(554)
                d_264_v41_: _dafny.Seq
                d_264_v41_ = _dafny.SeqWithoutIsStrInference([d_262_v39_, d_262_v39_])
                rhs44_ = (0) - ((d_263_v40_).cf0)
                rhs45_ = True
                rhs46_ = not(((self).fm4(d_239_cf5_, (_dafny.MultiSet(d_255_v32_)).cardinality, d_238_cf6_, globalState) if (self).f22 else d_238_cf6_))
                rhs47_ = (d_264_v41_)[default__.safeIndex(p0, len(d_264_v41_))]
                lhs35_ = self
                lhs36_ = self
                d_239_cf5_ = rhs44_
                lhs35_.f23 = rhs45_
                lhs36_.f23 = rhs46_
                d_262_v39_ = rhs47_
                d_265_v42_: _dafny.MultiSet
                d_265_v42_ = _dafny.MultiSet([len(d_206_v3_), ((default__.fm6(d_240_cf4_, d_241_cf3_, globalState)) - ((d_258_v35_).set(False, default__.abs(p0)))).cardinality, d_239_cf5_])
                d_265_v42_ = d_265_v42_
            d_266_v43_: _dafny.Seq
            d_266_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbjypvnv"))
            d_267_v44_: _dafny.Seq
            d_267_v44_ = _dafny.SeqWithoutIsStrInference([(len((d_266_v43_).set(default__.safeIndex(p0, len(d_266_v43_)), d_203_v0_))) * (d_240_cf4_), p0, d_239_cf5_, d_239_cf5_, p0])
            d_267_v44_ = (d_267_v44_) + (d_267_v44_)
            if ((d_266_v43_) + (d_266_v43_)) != (d_266_v43_):
                d_268_v45_: _dafny.MultiSet
                d_268_v45_ = _dafny.MultiSet([d_241_cf3_, not(not(self.f23))])
                d_269_v46_: _dafny.Map
                d_269_v46_ = _dafny.Map({d_268_v45_: d_203_v0_})
                d_269_v46_ = (d_269_v46_).set((d_268_v45_) | (d_268_v45_), d_203_v0_)
                d_270_v47_: _dafny.Array
                nw50_ = _dafny.Array(False, 21)
                d_270_v47_ = nw50_
                index47_ = default__.safeIndex(126, (d_270_v47_).length(0))
                (d_270_v47_)[index47_] = (d_239_cf5_) == (p0)
                d_271_v48_: _dafny.Array
                nw51_ = _dafny.Array(_dafny.Map({}), 3)
                d_271_v48_ = nw51_
                d_272_v49_: _dafny.Map
                d_272_v49_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([d_267_v44_]))): p0})
                d_273_v50_: _dafny.Map
                d_273_v50_ = d_272_v49_
                index48_ = default__.safeIndex(466, (d_271_v48_).length(0))
                (d_271_v48_)[index48_] = (d_272_v49_) | ((d_273_v50_))
                index49_ = default__.safeIndex(126, (d_270_v47_).length(0))
                index50_ = default__.safeIndex(466, (d_271_v48_).length(0))
                rhs48_ = not((self).f22)
                rhs49_ = d_272_v49_
                lhs37_ = d_270_v47_
                lhs38_ = default__.safeIndex(126, (d_270_v47_).length(0))
                lhs39_ = d_271_v48_
                lhs40_ = default__.safeIndex(466, (d_271_v48_).length(0))
                lhs37_[lhs38_] = rhs48_
                lhs39_[lhs40_] = rhs49_
                r0 = (self).fm4(485, (0) - ((d_239_cf5_) * (d_239_cf5_)), self.f23, globalState)
                d_274_v51_: _dafny.Map
                d_274_v51_ = _dafny.Map({d_241_cf3_: d_238_cf6_})
                d_275_v52_: _dafny.Map
                d_275_v52_ = _dafny.Map({d_208_v5_: (((d_274_v51_)[d_238_cf6_] if (d_238_cf6_) in (d_274_v51_) else not(d_238_cf6_))) == (self.f23)})
                r0 = ((d_275_v52_)[d_208_v5_] if (d_208_v5_) in (d_275_v52_) else self.f23)
                index51_ = default__.safeIndex(126, (d_270_v47_).length(0))
                (d_270_v47_)[index51_] = d_238_cf6_
            elif True:
                d_276_v53_: _dafny.Set
                d_276_v53_ = _dafny.Set({738, d_239_cf5_, d_240_cf4_, len(d_266_v43_)})
                d_277_v54_: D0
                d_277_v54_ = D0_DC0(818)
                d_276_v53_ = _dafny.Set({(d_277_v54_).cf0, d_240_cf4_, p0})
                d_239_cf5_ = d_239_cf5_
                (globalState).f19 = self.f23
                d_238_cf6_ = (self).fm4(default__.fm0(False, (self).f22, globalState), d_240_cf4_, (d_267_v44_) < (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))) for d_278_i4_ in range(default__.abs(157))])), globalState)
                d_279_v55_: _dafny.Set
                d_279_v55_ = _dafny.Set({(self).f22})
                d_280_v56_: _dafny.MultiSet
                d_280_v56_ = _dafny.MultiSet([d_241_cf3_, d_238_cf6_])
                d_281_v57_: D0
                d_281_v57_ = D0_DC2(not((self).f22), p0, 301, (self).f22)
                d_282_v58_: _dafny.Array
                nw52_ = _dafny.Array(None, 21)
                nw52_[int(0)] = d_239_cf5_
                nw52_[int(1)] = d_239_cf5_
                nw52_[int(2)] = len(default__.fm7(globalState))
                nw52_[int(3)] = d_239_cf5_
                nw52_[int(4)] = len(d_279_v55_)
                nw52_[int(5)] = p0
                nw52_[int(6)] = d_239_cf5_
                nw52_[int(7)] = d_240_cf4_
                nw52_[int(8)] = (d_280_v56_).cardinality
                nw52_[int(9)] = p0
                nw52_[int(10)] = d_239_cf5_
                nw52_[int(11)] = (d_267_v44_)[default__.safeIndex(595, len(d_267_v44_))]
                nw52_[int(12)] = d_240_cf4_
                nw52_[int(13)] = d_239_cf5_
                nw52_[int(14)] = len(_dafny.SeqWithoutIsStrInference([(self).fm4(d_240_cf4_, 923, (self).f22, globalState)]))
                nw52_[int(15)] = d_239_cf5_
                nw52_[int(16)] = d_240_cf4_
                nw52_[int(17)] = d_239_cf5_
                nw52_[int(18)] = default__.fm0(not(self.f23), (d_281_v57_).cf6, globalState)
                nw52_[int(19)] = d_239_cf5_
                nw52_[int(20)] = d_240_cf4_
                d_282_v58_ = nw52_
                d_283_v59_: _dafny.Map
                d_283_v59_ = _dafny.Map({d_282_v58_: d_280_v56_})
                d_283_v59_ = (d_283_v59_).set(d_282_v58_, d_280_v56_)
            d_284_v60_: _dafny.Array
            nw53_ = _dafny.Array(None, 4)
            nw53_[int(0)] = d_238_cf6_
            nw53_[int(1)] = d_238_cf6_
            nw53_[int(2)] = (self).f22
            nw53_[int(3)] = not(False)
            d_284_v60_ = nw53_
            index52_ = default__.safeIndex(134, (d_284_v60_).length(0))
            (d_284_v60_)[index52_] = d_238_cf6_
            index53_ = default__.safeIndex(134, (d_284_v60_).length(0))
            (d_284_v60_)[index53_] = True
        elif source2_.is_DC0:
            d_285___mcc_h6_ = source2_.cf0
            d_286_cf0_ = d_285___mcc_h6_
            d_287_v61_: _dafny.Set
            d_287_v61_ = _dafny.Set({d_206_v3_, d_206_v3_})
            (globalState).f6 = (d_287_v61_).ispropersubset(d_287_v61_)
            if self.f23:
                d_288_v62_: _dafny.Seq
                d_288_v62_ = _dafny.SeqWithoutIsStrInference([d_286_cf0_, d_286_cf0_])
                d_289_v63_: _dafny.Map
                d_289_v63_ = _dafny.Map({self.f23: d_288_v62_})
                d_290_v64_: _dafny.Set
                d_290_v64_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_286_cf0_ for d_291_i5_ in range(default__.abs(923))]), d_288_v62_, (((d_289_v63_)[True] if (True) in (d_289_v63_) else d_288_v62_)).set(default__.safeIndex(d_286_cf0_, len(((d_289_v63_)[True] if (True) in (d_289_v63_) else d_288_v62_))), 279)})
                d_292_v65_: _dafny.Map
                d_292_v65_ = _dafny.Map({(self).f22: self.f23})
                d_293_v66_: _dafny.Map
                d_293_v66_ = _dafny.Map({d_290_v64_: (d_292_v65_) | (d_292_v65_)})
                d_293_v66_ = (d_293_v66_).set((d_290_v64_) | (d_290_v64_), d_292_v65_)
                d_294_v67_: _dafny.Seq
                d_294_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jowiusrj"))
                d_295_v68_: _dafny.Array
                nw54_ = _dafny.Array(False, 15)
                d_295_v68_ = nw54_
                index54_ = default__.safeIndex(315, (d_295_v68_).length(0))
                (d_295_v68_)[index54_] = (self).f22
                index55_ = default__.safeIndex(315, (d_295_v68_).length(0))
                rhs50_ = ((d_294_v67_ if self.f23 else d_294_v67_)).set(default__.safeIndex(909, len((d_294_v67_ if self.f23 else d_294_v67_))), d_203_v0_)
                rhs51_ = d_203_v0_
                rhs52_ = self.f23
                lhs41_ = globalState
                lhs42_ = d_295_v68_
                lhs43_ = default__.safeIndex(315, (d_295_v68_).length(0))
                d_294_v67_ = rhs50_
                lhs41_.f18 = rhs51_
                lhs42_[lhs43_] = rhs52_
                d_296_v69_: _dafny.MultiSet
                d_296_v69_ = _dafny.MultiSet([self.f23])
                r0 = (len(d_294_v67_)) != ((((d_296_v69_)[self.f23] if (self.f23) in (d_296_v69_) else p0)) + (d_286_cf0_))
                d_286_cf0_ = d_286_cf0_
                index56_ = default__.safeIndex(315, (d_295_v68_).length(0))
                (d_295_v68_)[index56_] = (self).f22
            elif True:
                d_297_v70_: _dafny.Seq
                d_297_v70_ = _dafny.SeqWithoutIsStrInference([p0, p0, d_286_cf0_, p0, 850])
                d_286_cf0_ = default__.safeDivisionInt(700, len((d_297_v70_) + (d_297_v70_)))
                d_298_v71_: _dafny.Array
                nw55_ = _dafny.Array(False, 5)
                d_298_v71_ = nw55_
                d_298_v71_ = d_298_v71_
                d_286_cf0_ = (p0) * (d_286_cf0_)
                d_286_cf0_ = (p0) * (d_286_cf0_)
                d_299_v72_: _dafny.MultiSet
                d_299_v72_ = _dafny.MultiSet([self.f23, True])
                d_300_v73_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_300_v73_ = nw56_
                d_301_v74_: _dafny.Array
                def lambda11_(d_302_p0_):
                    def lambda12_(d_303_i6_):
                        return (d_303_i6_) + (d_302_p0_)

                    return lambda12_

                init6_ = lambda11_(p0)
                nw57_ = _dafny.Array(None, 2)
                for i0_6_ in range(nw57_.length(0)):
                    nw57_[i0_6_] = init6_(i0_6_)
                d_301_v74_ = nw57_
                index57_ = default__.safeIndex(517, (d_300_v73_).length(0))
                (d_300_v73_)[index57_] = d_301_v74_
                index58_ = default__.safeIndex(517, (d_300_v73_).length(0))
                rhs53_ = d_299_v72_
                rhs54_ = d_286_cf0_
                rhs55_ = d_208_v5_
                rhs56_ = d_301_v74_
                rhs57_ = d_298_v71_
                lhs44_ = d_300_v73_
                lhs45_ = default__.safeIndex(517, (d_300_v73_).length(0))
                d_299_v72_ = rhs53_
                d_286_cf0_ = rhs54_
                d_208_v5_ = rhs55_
                lhs44_[lhs45_] = rhs56_
                d_298_v71_ = rhs57_
            (globalState).f19 = (self.f23 if (self).f22 else (self).f22)
            d_304_v75_: _dafny.Array
            nw58_ = _dafny.Array(False, 6)
            d_304_v75_ = nw58_
            index59_ = default__.safeIndex(161, (d_304_v75_).length(0))
            (d_304_v75_)[index59_] = False
            index60_ = default__.safeIndex(161, (d_304_v75_).length(0))
            (d_304_v75_)[index60_] = (default__.fm0(self.f23, (self).fm4(len(d_206_v3_), len(_dafny.Set({(self).f22, (self).f22})), not((self).f22), globalState), globalState)) < ((_dafny.MultiSet([self.f23])).cardinality)
        elif True:
            d_305___mcc_h7_ = source2_.cf7
            d_306_cf7_ = d_305___mcc_h7_
            d_307_v76_: _dafny.Seq
            d_307_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfhewcmh"))
            if (not((d_307_v76_) == (_dafny.SeqWithoutIsStrInference([d_203_v0_ for d_308_i7_ in range(default__.abs(754))]))) if not(self.f23) else self.f23):
                d_309_v77_: _dafny.Map
                d_309_v77_ = _dafny.Map({945: d_206_v3_})
                d_310_v78_: C0
                nw59_ = C0()
                nw59_.ctor__(False, (d_309_v77_) | (d_309_v77_))
                d_310_v78_ = nw59_
                d_311_v79_: _dafny.Seq
                d_312_v80_: bool
                out8_: _dafny.Seq
                out9_: bool
                out8_, out9_ = (self).m1(p0, (self).f22, ((0) - (p0)) >= (p0), len(_dafny.SeqWithoutIsStrInference([d_203_v0_ for d_313_i8_ in range(default__.abs(790))])), globalState)
                d_311_v79_ = out8_
                d_312_v80_ = out9_
                d_314_v81_: C0
                nw60_ = C0()
                nw60_.ctor__((self).fm4(p0, p0, d_310_v78_.f24, globalState), d_309_v77_)
                d_314_v81_ = nw60_
                d_315_v82_: _dafny.MultiSet
                d_315_v82_ = _dafny.MultiSet([p0, default__.safeModuloInt(default__.fm0(d_314_v81_.f24, d_312_v80_, globalState), p0)])
                rhs58_ = d_310_v78_
                rhs59_ = (d_315_v82_).set((d_208_v5_).cf2, default__.abs(p0))
                rhs60_ = not((self).f22)
                rhs61_ = d_203_v0_
                rhs62_ = (self).fm4(len(d_307_v76_), p0, not(d_314_v81_.f24), globalState)
                lhs46_ = globalState
                lhs47_ = globalState
                lhs48_ = self
                d_314_v81_ = rhs58_
                d_315_v82_ = rhs59_
                lhs46_.f19 = rhs60_
                lhs47_.f18 = rhs61_
                lhs48_.f23 = rhs62_
                d_311_v79_ = d_307_v76_
                d_316_v83_: _dafny.Map
                d_316_v83_ = _dafny.Map({(self).f22: 271})
                r1 = _dafny.Map({d_316_v83_: _dafny.SeqWithoutIsStrInference([d_203_v0_ for d_317_i9_ in range(default__.abs(306))])})
            elif True:
                d_318_v84_: int
                d_318_v84_ = -212
                d_318_v84_ = d_318_v84_
                d_319_v85_: _dafny.Array
                def lambda13_(d_320_p0_):
                    def lambda14_(d_321_i10_):
                        return (d_321_i10_) * (d_320_p0_)

                    return lambda14_

                init7_ = lambda13_(p0)
                nw61_ = _dafny.Array(None, 21)
                for i0_7_ in range(nw61_.length(0)):
                    nw61_[i0_7_] = init7_(i0_7_)
                d_319_v85_ = nw61_
                d_322_v86_: _dafny.Seq
                d_322_v86_ = _dafny.SeqWithoutIsStrInference([p0])
                index61_ = default__.safeIndex(457, (d_319_v85_).length(0))
                (d_319_v85_)[index61_] = default__.safeModuloInt((0) - (p0), len(d_322_v86_))
                d_323_v87_: _dafny.Array
                d_323_v87_ = d_319_v85_
                index62_ = default__.safeIndex(457, (d_319_v85_).length(0))
                rhs63_ = (d_318_v84_) * ((p0 if (self).f22 else 407))
                rhs64_ = (d_323_v87_)
                rhs65_ = 35
                lhs49_ = d_319_v85_
                lhs50_ = default__.safeIndex(457, (d_319_v85_).length(0))
                lhs49_[lhs50_] = rhs63_
                d_319_v85_ = rhs64_
                d_318_v84_ = rhs65_
                d_324_v88_: _dafny.Map
                d_324_v88_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjfewtn")): (self).f22})
                d_324_v88_ = (d_324_v88_).set(d_307_v76_, (self).f22)
                d_325_v89_: _dafny.Array
                nw62_ = _dafny.Array(None, 4)
                nw62_[int(0)] = self.f23
                nw62_[int(1)] = self.f23
                nw62_[int(2)] = not((self).f22)
                nw62_[int(3)] = self.f23
                d_325_v89_ = nw62_
                d_326_v90_: _dafny.Set
                d_326_v90_ = _dafny.Set({d_208_v5_})
                index63_ = default__.safeIndex(857, (d_325_v89_).length(0))
                (d_325_v89_)[index63_] = (d_326_v90_).issubset(d_326_v90_)
                index64_ = default__.safeIndex(857, (d_325_v89_).length(0))
                (d_325_v89_)[index64_] = (self).fm4(d_318_v84_, default__.fm0(not((self).f22), self.f23, globalState), (self).f22, globalState)
                (globalState).f6 = self.f23
            d_327_v91_: _dafny.Map
            d_327_v91_ = _dafny.Map({(p0) > ((d_208_v5_).cf2): d_307_v76_})
            d_327_v91_ = d_327_v91_
            d_328_v92_: D3
            d_328_v92_ = D3_DC6(d_307_v76_)
            d_329_v93_: _dafny.Seq
            d_329_v93_ = _dafny.SeqWithoutIsStrInference([d_307_v76_, d_307_v76_, d_307_v76_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), (d_328_v92_).cf10])
            if (d_307_v76_) <= ((d_329_v93_)[default__.safeIndex(-908, len(d_329_v93_))]):
                r0 = not((self).f22)
                (globalState).f19 = self.f23
                d_330_v94_: _dafny.Map
                d_330_v94_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference([p0 for d_331_i11_ in range(default__.abs(20))])})
                d_332_v95_: _dafny.Seq
                d_332_v95_ = _dafny.SeqWithoutIsStrInference([p0])
                d_333_v96_: _dafny.Map
                d_333_v96_ = _dafny.Map({len(((d_330_v94_)[(self).f22] if ((self).f22) in (d_330_v94_) else d_332_v95_)): d_206_v3_})
                d_334_v97_: C0
                nw63_ = C0()
                nw63_.ctor__(self.f23, d_333_v96_)
                d_334_v97_ = nw63_
                (globalState).f6 = self.f23
                d_335_v98_: _dafny.Array
                nw64_ = _dafny.Array(False, 5)
                d_335_v98_ = nw64_
                d_336_v99_: _dafny.Map
                d_336_v99_ = _dafny.Map({_dafny.Map({d_335_v98_: d_334_v97_.f24}): (_dafny.MultiSet([p0])).cardinality})
                d_337_v100_: _dafny.Map
                d_337_v100_ = _dafny.Map({d_335_v98_: self.f23})
                d_336_v99_ = (d_336_v99_).set(d_337_v100_, p0)
            elif True:
                (globalState).f6 = (self).f22
                (globalState).f6 = not(((self).f22) == (self.f23))
                d_338_v101_: D0
                d_338_v101_ = D0_DC0(p0)
                (globalState).f19 = (p0) <= ((d_338_v101_).cf0)
                (globalState).f6 = (self).f22
                d_339_v102_: _dafny.Map
                d_339_v102_ = _dafny.Map({p0: d_206_v3_})
                d_340_v103_: C0
                nw65_ = C0()
                nw65_.ctor__(False, d_339_v102_)
                d_340_v103_ = nw65_
                d_341_v104_: _dafny.Seq
                d_341_v104_ = _dafny.SeqWithoutIsStrInference([d_340_v103_, d_340_v103_, d_340_v103_, d_340_v103_])
                d_342_v105_: _dafny.MultiSet
                d_342_v105_ = _dafny.MultiSet([p0])
                d_343_v106_: _dafny.Set
                d_343_v106_ = _dafny.Set({p0})
                d_344_v107_: _dafny.Map
                d_344_v107_ = _dafny.Map({d_340_v103_.f24: len(d_343_v106_)})
                d_345_v108_: _dafny.Array
                nw66_ = _dafny.Array(None, 29)
                nw66_[int(0)] = p0
                nw66_[int(1)] = 173
                nw66_[int(2)] = p0
                nw66_[int(3)] = p0
                nw66_[int(4)] = 978
                nw66_[int(5)] = p0
                nw66_[int(6)] = p0
                nw66_[int(7)] = p0
                nw66_[int(8)] = p0
                nw66_[int(9)] = p0
                nw66_[int(10)] = len(d_341_v104_)
                nw66_[int(11)] = default__.safeModuloInt(p0, p0)
                nw66_[int(12)] = (p0) + (p0)
                nw66_[int(13)] = p0
                nw66_[int(14)] = p0
                nw66_[int(15)] = ((0) - (p0)) - (len(_dafny.SeqWithoutIsStrInference([d_203_v0_ for d_346_i12_ in range(default__.abs(929))])))
                nw66_[int(16)] = p0
                nw66_[int(17)] = (0) - (p0)
                nw66_[int(18)] = ((d_342_v105_).set(p0, default__.abs(-355))).cardinality
                nw66_[int(19)] = p0
                nw66_[int(20)] = len((d_206_v3_).set(default__.safeIndex(-292, len(d_206_v3_)), d_340_v103_.f24))
                nw66_[int(21)] = p0
                nw66_[int(22)] = p0
                nw66_[int(23)] = p0
                nw66_[int(24)] = p0
                nw66_[int(25)] = default__.safeDivisionInt(-346, p0)
                nw66_[int(26)] = default__.safeModuloInt(p0, p0)
                nw66_[int(27)] = p0
                nw66_[int(28)] = ((d_344_v107_)[d_340_v103_.f24] if (d_340_v103_.f24) in (d_344_v107_) else len(d_307_v76_))
                d_345_v108_ = nw66_
                index65_ = default__.safeIndex(195, (d_345_v108_).length(0))
                (d_345_v108_)[index65_] = p0
                d_347_v109_: _dafny.Seq
                d_347_v109_ = _dafny.SeqWithoutIsStrInference([-72, p0, (0) - (p0), p0])
                index66_ = default__.safeIndex(594, (d_345_v108_).length(0))
                (d_345_v108_)[index66_] = len(d_347_v109_)
                d_348_v110_: _dafny.Map
                d_348_v110_ = _dafny.Map({p0: (self).f22})
                index67_ = default__.safeIndex(195, (d_345_v108_).length(0))
                index68_ = default__.safeIndex(594, (d_345_v108_).length(0))
                rhs66_ = ((len((d_307_v76_).set(default__.safeIndex(p0, len(d_307_v76_)), d_203_v0_))) not in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - (p0)])), p0])) if (self).f22 else d_340_v103_.f24)
                rhs67_ = p0
                rhs68_ = len(d_348_v110_)
                rhs69_ = (default__.fm2(_dafny.MultiSet(d_307_v76_), globalState) if (self).fm4((d_338_v101_).cf0, p0, self.f23, globalState) else _dafny.SeqWithoutIsStrInference([p0]))
                lhs51_ = globalState
                lhs52_ = d_345_v108_
                lhs53_ = default__.safeIndex(195, (d_345_v108_).length(0))
                lhs54_ = d_345_v108_
                lhs55_ = default__.safeIndex(594, (d_345_v108_).length(0))
                lhs51_.f19 = rhs66_
                lhs52_[lhs53_] = rhs67_
                lhs54_[lhs55_] = rhs68_
                d_347_v109_ = rhs69_
            r0 = (not(True)) and ((75) > (p0))
        d_203_v0_ = default__.fm5(globalState)
        (globalState).f18 = d_203_v0_
        if (True if (self).f22 else (d_206_v3_) <= (d_206_v3_)):
            d_349_v111_: _dafny.Map
            d_349_v111_ = _dafny.Map({default__.fm0((self).fm4(p0, p0, (self).f22, globalState), False, globalState): (d_206_v3_).set(default__.safeIndex(p0, len(d_206_v3_)), self.f23)})
            d_350_v112_: C0
            nw67_ = C0()
            nw67_.ctor__(False, d_349_v111_)
            d_350_v112_ = nw67_
            d_351_v113_: D3
            d_351_v113_ = D3_DC6(default__.fm3(p0, globalState))
            d_352_v114_: _dafny.Map
            d_352_v114_ = _dafny.Map({(self).f22: d_351_v113_})
            d_352_v114_ = (d_352_v114_).set(d_350_v112_.f24, d_351_v113_)
            (globalState).f6 = d_350_v112_.f24
            d_353_v115_: _dafny.Array
            def lambda15_(d_354_v0_):
                def lambda16_(d_355_i13_):
                    return _dafny.SeqWithoutIsStrInference([d_354_v0_ for d_356_i14_ in range(default__.abs(132))])

                return lambda16_

            init8_ = lambda15_(d_203_v0_)
            nw68_ = _dafny.Array(None, 17)
            for i0_8_ in range(nw68_.length(0)):
                nw68_[i0_8_] = init8_(i0_8_)
            d_353_v115_ = nw68_
            index69_ = default__.safeIndex(867, (d_353_v115_).length(0))
            (d_353_v115_)[index69_] = (d_351_v113_).cf10
            d_357_v116_: _dafny.Seq
            d_357_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccagyoi"))
            index70_ = default__.safeIndex(867, (d_353_v115_).length(0))
            (d_353_v115_)[index70_] = (d_357_v116_) + ((d_351_v113_).cf10)
            d_358_v117_: C0
            nw69_ = C0()
            nw69_.ctor__((p0) != (p0), (d_350_v112_).f25)
            d_358_v117_ = nw69_
        elif True:
            d_203_v0_ = d_203_v0_
            d_359_v118_: _dafny.Seq
            d_359_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbspy"))
            (self).f23 = (len(d_359_v118_)) < (len(d_359_v118_))
            d_360_v119_: _dafny.MultiSet
            d_360_v119_ = _dafny.MultiSet([d_359_v118_, d_359_v118_, (d_359_v118_) + (_dafny.SeqWithoutIsStrInference([d_203_v0_ for d_361_i15_ in range(default__.abs(469))]))])
            d_362_v120_: _dafny.MultiSet
            d_362_v120_ = _dafny.MultiSet([p0])
            rhs70_ = ((582) - (p0)) != ((0) - (p0))
            rhs71_ = False
            rhs72_ = (d_362_v120_).ispropersubset((d_362_v120_).intersection(_dafny.MultiSet([p0])))
            rhs73_ = d_360_v119_
            lhs56_ = globalState
            lhs57_ = globalState
            lhs58_ = globalState
            lhs56_.f6 = rhs70_
            lhs57_.f6 = rhs71_
            lhs58_.f19 = rhs72_
            d_360_v119_ = rhs73_
            d_363_v121_: _dafny.Array
            nw70_ = _dafny.Array(False, 9)
            d_363_v121_ = nw70_
            index71_ = default__.safeIndex(971, (d_363_v121_).length(0))
            (d_363_v121_)[index71_] = (self).f22
            d_364_v122_: _dafny.Array
            nw71_ = _dafny.Array(int(0), 8)
            d_364_v122_ = nw71_
            index72_ = default__.safeIndex(971, (d_363_v121_).length(0))
            rhs74_ = d_364_v122_
            rhs75_ = self.f23
            lhs59_ = globalState
            lhs60_ = d_363_v121_
            lhs61_ = default__.safeIndex(971, (d_363_v121_).length(0))
            lhs59_.f13 = rhs74_
            lhs60_[lhs61_] = rhs75_
            d_365_v123_: _dafny.Seq
            d_365_v123_ = _dafny.SeqWithoutIsStrInference([p0])
            d_366_v124_: _dafny.Map
            d_366_v124_ = _dafny.Map({(d_363_v121_)[default__.safeIndex(971, (d_363_v121_).length(0))]: (d_365_v123_)[default__.safeIndex(p0, len(d_365_v123_))]})
            d_366_v124_ = (d_366_v124_).set(self.f23, (p0 if (self).f22 else p0))
        d_367_v125_: _dafny.Map
        d_367_v125_ = _dafny.Map({len(d_206_v3_): d_206_v3_})
        d_368_v126_: C0
        nw72_ = C0()
        nw72_.ctor__((self).f22, d_367_v125_)
        d_368_v126_ = nw72_
        d_369_v127_: _dafny.Seq
        d_369_v127_ = _dafny.SeqWithoutIsStrInference([d_368_v126_])
        d_370_v128_: _dafny.Array
        nw73_ = _dafny.Array(None, 8)
        nw73_[int(0)] = (self).f22
        nw73_[int(1)] = False
        nw73_[int(2)] = (self).f22
        nw73_[int(3)] = not(((self).f22) == ((self).fm4(len(d_369_v127_), p0, (self).f22, globalState)))
        nw73_[int(4)] = self.f23
        nw73_[int(5)] = (self).f22
        nw73_[int(6)] = self.f23
        nw73_[int(7)] = (p0) == (p0)
        d_370_v128_ = nw73_
        index73_ = default__.safeIndex(292, (d_370_v128_).length(0))
        (d_370_v128_)[index73_] = (self).f22
        index74_ = default__.safeIndex(292, (d_370_v128_).length(0))
        (d_370_v128_)[index74_] = ((p0) + ((0) - (p0))) == (p0)
        r0 = (self).f22
        d_371_v129_: _dafny.Map
        d_371_v129_ = _dafny.Map({(self).f22: p0})
        r1 = _dafny.Map({(d_371_v129_) | (d_371_v129_): _dafny.SeqWithoutIsStrInference([d_203_v0_ for d_372_i16_ in range(default__.abs(961))])})
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        (globalState).f0 = _dafny.Map({p1: p0})
        d_373_v0_: _dafny.Array
        nw74_ = _dafny.Array(None, 29)
        d_373_v0_ = nw74_
        d_373_v0_ = d_373_v0_
        if self.f23:
            d_374_v1_: _dafny.Map
            d_374_v1_ = _dafny.Map({self.f23: p2})
            d_374_v1_ = _dafny.Map({p2: not (True) or ((self).f22)})
            d_375_v2_: _dafny.Seq
            d_375_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahbq"))
            d_376_v3_: str
            d_376_v3_ = _dafny.CodePoint('h')
            d_377_v4_: _dafny.Seq
            d_377_v4_ = _dafny.SeqWithoutIsStrInference([(d_375_v2_) <= ((d_375_v2_).set(default__.safeIndex(p3, len(d_375_v2_)), d_376_v3_))])
            d_377_v4_ = (d_377_v4_) + (d_377_v4_)
            d_378_v5_: C0
            nw75_ = C0()
            nw75_.ctor__(p1, _dafny.Map({789: d_377_v4_}))
            d_378_v5_ = nw75_
            d_379_v6_: int
            d_379_v6_ = -403
            d_380_v7_: _dafny.Seq
            d_380_v7_ = _dafny.SeqWithoutIsStrInference([d_379_v6_])
            rhs76_ = d_378_v5_
            rhs77_ = default__.safeModuloInt(p3, (0) - (len(d_380_v7_)))
            d_378_v5_ = rhs76_
            d_379_v6_ = rhs77_
            (d_378_v5_).f24 = self.f23
            d_381_v9_: _dafny.Array
            nw76_ = _dafny.Array(None, 25)
            nw76_[int(0)] = (self).f22
            nw76_[int(1)] = True
            nw76_[int(2)] = not(not(p1))
            nw76_[int(3)] = d_378_v5_.f24
            nw76_[int(4)] = p1
            nw76_[int(5)] = True
            nw76_[int(6)] = False
            nw76_[int(7)] = p2
            nw76_[int(8)] = p2
            nw76_[int(9)] = self.f23
            nw76_[int(10)] = False
            nw76_[int(11)] = p1
            nw76_[int(12)] = d_378_v5_.f24
            nw76_[int(13)] = p2
            nw76_[int(14)] = (self).f22
            nw76_[int(15)] = (self).f22
            nw76_[int(16)] = self.f23
            nw76_[int(17)] = (self).f22
            nw76_[int(18)] = d_378_v5_.f24
            nw76_[int(19)] = p1
            nw76_[int(20)] = True
            nw76_[int(21)] = p1
            nw76_[int(22)] = p2
            nw76_[int(23)] = p1
            def iife16_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(36, 816):
                    d_382_v8_: int = compr_4_
                    if ((36) <= (d_382_v8_)) and ((d_382_v8_) < (816)):
                        coll4_[(d_382_v8_) - (len(d_375_v2_))] = True
                return _dafny.Map(coll4_)
            nw76_[int(24)] = (self).fm4(len(iife16_()
            ), p3, not(p2), globalState)
            d_381_v9_ = nw76_
            d_383_v10_: _dafny.MultiSet
            d_383_v10_ = _dafny.MultiSet([d_381_v9_])
            d_379_v6_ = ((d_383_v10_) | (d_383_v10_)).cardinality
        elif True:
            d_384_v11_: int
            d_384_v11_ = 478
            d_385_v12_: _dafny.MultiSet
            d_385_v12_ = _dafny.MultiSet([p3])
            d_386_v13_: _dafny.MultiSet
            d_386_v13_ = _dafny.MultiSet([self.f23])
            d_387_v14_: _dafny.Map
            d_387_v14_ = _dafny.Map({d_386_v13_: self.f23})
            d_388_v15_: _dafny.Map
            d_388_v15_ = _dafny.Map({p2: d_384_v11_})
            d_389_v17_: _dafny.Seq
            def iife17_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-789, 579):
                    d_390_v16_: int = compr_5_
                    if ((-789) <= (d_390_v16_)) and ((d_390_v16_) < (579)):
                        coll5_[(d_390_v16_) * (-899)] = p0
                return _dafny.Map(coll5_)
            d_389_v17_ = _dafny.SeqWithoutIsStrInference([len(d_388_v15_), len(iife17_()
            ), p0, d_384_v11_])
            d_391_v18_: _dafny.Seq
            d_391_v18_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p1: False})), len(d_387_v14_), (d_389_v17_)[default__.safeIndex(p0, len(d_389_v17_))]])
            d_384_v11_ = ((d_385_v12_).cardinality) + ((p3) * (len(d_391_v18_)))
            d_392_v19_: _dafny.Set
            d_392_v19_ = _dafny.Set({p0, p3, p3, d_384_v11_, (0) - (default__.safeModuloInt(d_384_v11_, p0))})
            d_392_v19_ = _dafny.Set({p3})
            d_393_v20_: D0
            d_393_v20_ = D0_DC2((self).f22, -817, d_384_v11_, p1)
            d_394_v21_: _dafny.Seq
            d_394_v21_ = _dafny.SeqWithoutIsStrInference([d_393_v20_, d_393_v20_, d_393_v20_, d_393_v20_])
            d_395_v23_: _dafny.Map
            d_395_v23_ = _dafny.Map({(self).f22: p2})
            d_396_v24_: _dafny.Seq
            d_396_v24_ = _dafny.SeqWithoutIsStrInference([True])
            def iife18_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(263, 685):
                    d_398_v22_: int = compr_6_
                    if ((263) <= (d_398_v22_)) and ((d_398_v22_) < (685)):
                        coll6_[default__.safeModuloInt(d_398_v22_, len(_dafny.Map({d_384_v11_: d_384_v11_})))] = p3
                return _dafny.Map(coll6_)
            rhs78_ = _dafny.SeqWithoutIsStrInference([d_393_v20_ for d_397_i0_ in range(default__.abs(999))])
            rhs79_ = p2
            rhs80_ = len(((_dafny.SeqWithoutIsStrInference([(self).fm4(len((iife18_()
            ).set(p0, len(d_395_v23_))), d_384_v11_, True, globalState), p2, p2, False]) if False else d_396_v24_)) + ((d_396_v24_) + (_dafny.SeqWithoutIsStrInference([(self).f22]))))
            lhs62_ = self
            d_394_v21_ = rhs78_
            lhs62_.f23 = rhs79_
            d_384_v11_ = rhs80_
            d_399_v25_: _dafny.Map
            d_399_v25_ = _dafny.Map({p0: default__.fm10(p3, d_384_v11_, globalState)})
            d_400_v26_: C0
            nw77_ = C0()
            nw77_.ctor__(p2, d_399_v25_)
            d_400_v26_ = nw77_
            d_401_v27_: _dafny.Array
            def lambda17_(d_402_p3_):
                def lambda18_(d_403_i1_):
                    return default__.safeDivisionInt(d_403_i1_, d_402_p3_)

                return lambda18_

            init9_ = lambda17_(p3)
            nw78_ = _dafny.Array(None, 17)
            for i0_9_ in range(nw78_.length(0)):
                nw78_[i0_9_] = init9_(i0_9_)
            d_401_v27_ = nw78_
            d_404_v28_: _dafny.Array
            nw79_ = _dafny.Array(None, 13)
            nw79_[int(0)] = d_401_v27_
            nw79_[int(1)] = d_401_v27_
            nw79_[int(2)] = d_401_v27_
            nw79_[int(3)] = d_401_v27_
            nw79_[int(4)] = d_401_v27_
            nw79_[int(5)] = d_401_v27_
            nw79_[int(6)] = d_401_v27_
            nw79_[int(7)] = d_401_v27_
            nw79_[int(8)] = d_401_v27_
            nw79_[int(9)] = d_401_v27_
            nw79_[int(10)] = d_401_v27_
            nw79_[int(11)] = d_401_v27_
            nw79_[int(12)] = d_401_v27_
            d_404_v28_ = nw79_
            index75_ = default__.safeIndex(390, (d_404_v28_).length(0))
            (d_404_v28_)[index75_] = d_401_v27_
            index76_ = default__.safeIndex(390, (d_404_v28_).length(0))
            (d_404_v28_)[index76_] = d_401_v27_
        d_405_v29_: _dafny.Array
        def lambda19_(d_406_p0_):
            def lambda20_(d_407_i2_):
                return D0_DC3(D0_DC0(d_406_p0_))

            return lambda20_

        init10_ = lambda19_(p0)
        nw80_ = _dafny.Array(None, 13)
        for i0_10_ in range(nw80_.length(0)):
            nw80_[i0_10_] = init10_(i0_10_)
        d_405_v29_ = nw80_
        d_408_v30_: _dafny.Map
        d_408_v30_ = _dafny.Map({(0) - (p0): (0) - (p0)})
        d_409_v31_: _dafny.Map
        d_409_v31_ = _dafny.Map({d_408_v30_: d_405_v29_})
        d_410_v32_: _dafny.Seq
        d_410_v32_ = _dafny.SeqWithoutIsStrInference([d_405_v29_, d_405_v29_, d_405_v29_])
        d_411_v33_: _dafny.Seq
        d_411_v33_ = _dafny.SeqWithoutIsStrInference([self.f23, p2])
        d_412_v34_: _dafny.Map
        d_412_v34_ = _dafny.Map({p0: d_411_v33_})
        d_413_v35_: T0
        nw81_ = C0()
        nw81_.ctor__(p1, d_412_v34_)
        d_413_v35_ = nw81_
        d_414_v36_: _dafny.Map
        d_414_v36_ = _dafny.Map({p2: d_413_v35_})
        d_405_v29_ = ((d_409_v31_)[d_408_v30_] if (d_408_v30_) in (d_409_v31_) else (d_410_v32_)[default__.safeIndex(len(d_414_v36_), len(d_410_v32_))])
        d_415_v37_: _dafny.Array
        nw82_ = _dafny.Array(False, 2)
        d_415_v37_ = nw82_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_415_v37_).length(0)):
            d_416_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_416_i3_)) and ((d_416_i3_) < ((d_415_v37_).length(0)))):
                (d_415_v37_)[(d_416_i3_)] = (_dafny.SeqWithoutIsStrInference([d_411_v33_])) == ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])) + (_dafny.SeqWithoutIsStrInference([d_411_v33_])))
        d_415_v37_ = d_415_v37_
        d_417_v38_: _dafny.Seq
        d_417_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crax"))
        r0 = d_417_v38_
        r1 = (p2) or (p1)
        return r0, r1

    @property
    def f22(self):
        return self._f22
