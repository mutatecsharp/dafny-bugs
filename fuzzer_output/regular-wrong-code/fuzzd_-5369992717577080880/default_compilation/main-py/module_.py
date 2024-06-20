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
    def fm0(p0, globalState):
        if (D4_DC15(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuw"))), True)).cf24:
            return D0_DC0(-581)
        elif True:
            return D0_DC0((0) - (len(_dafny.Map({412: True}))))

    @staticmethod
    def fm1(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(904, 148):
                d_0_v0_: int = compr_0_
                if ((904) <= (d_0_v0_)) and ((d_0_v0_) < (148)):
                    coll0_[(d_0_v0_) + (len(_dafny.Map({not(True): _dafny.Map({635: True})})))] = True
            return _dafny.Map(coll0_)
        return D0_DC2(True, ((0) - ((0) - (len(_dafny.Map({True: 16}))))) + ((0) - (len(iife0_()
))), (_dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dh")))])])) | (_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-400]))])))

    @staticmethod
    def fm2(globalState):
        return not (False) or (False)

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (0) - (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference([True, (D0_DC2(True, 163, _dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.Map({-701: len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))}))}) for d_1_i0_ in range(default__.abs(713))]))])]))).cf2]))), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxdaqqlt"))))))

    @staticmethod
    def fm10(globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(True), True]))) + (_dafny.SeqWithoutIsStrInference([(D5_DC18(-672, not(not(True)))).cf28, False]))

    @staticmethod
    def fm11(p0, p1, globalState):
        if (_dafny.MultiSet([True])) == (_dafny.MultiSet([False, True])):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yo"))
        elif True:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjeyx"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2_i0_ in range(default__.abs(600))]))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([False])).intersection((_dafny.MultiSet([True])) - (_dafny.MultiSet([not((D0_DC1(True)).cf1)])))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: str
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_4_i1_ in range(default__.abs(421))])).Elements:
                d_5_v0_: str = compr_1_
                if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_4_i1_ in range(default__.abs(421))])):
                    coll1_[d_5_v0_] = True
            return _dafny.Map(coll1_)
        return ((_dafny.SeqWithoutIsStrInference([398 for d_3_i0_ in range(default__.abs(987))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrmaqmxx"))))]))) + ((_dafny.SeqWithoutIsStrInference([42, 627]) if not(False) else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(not(False)): iife1_()
        })), -641, 944])))

    @staticmethod
    def fm14(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): not(True)}), _dafny.Map({False: (D0_DC1((D2_DC8(True, 488, False)).cf7)).cf1})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: True}), _dafny.Map({False: not(True)})]))

    @staticmethod
    def fm15(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")) for d_6_i0_ in range(default__.abs(982))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hchl")) for d_7_i1_ in range(default__.abs(641))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_8_i3_ in range(default__.abs(278))]) for d_9_i2_ in range(default__.abs(-593))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_10_i4_ in range(default__.abs(990))])])))

    @staticmethod
    def fm16(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(480, -305):
                d_11_v0_: int = compr_2_
                if ((480) <= (d_11_v0_)) and ((d_11_v0_) < (-305)):
                    coll2_[(d_11_v0_) * (len(_dafny.SeqWithoutIsStrInference([False])))] = True
            return _dafny.Map(coll2_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlvvv"))): False})) | (iife2_()
        )) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, not(True), True, False])): not(True)}))

    @staticmethod
    def fm17(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-549, 598):
                d_12_v0_: int = compr_3_
                if ((-549) <= (d_12_v0_)) and ((d_12_v0_) < (598)):
                    coll3_[(d_12_v0_) + ((0) - (-485))] = False
            return _dafny.Map(coll3_)
        return _dafny.MultiSet([default__.safeModuloInt(len(_dafny.Set({False})), 837), (0) - ((437) + (len(iife3_()
        ))), 368, len((_dafny.Map({True: True})) | (_dafny.Map({True: True}))), (0) - (len(_dafny.Set({not(False), True})))])

    @staticmethod
    def fm18(globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))): 356})).keys.Elements:
                    d_15_v1_: int = compr_6_
                    if (d_15_v1_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))): 356})):
                        coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_15_v1_, -91)]))
                return _dafny.Set(coll6_)
            coll4_ = _dafny.Map()
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: int
                for compr_5_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))): 356})).keys.Elements:
                    d_13_v1_: int = compr_5_
                    if (d_13_v1_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))): 356})):
                        coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_13_v1_, -91)]))
                return _dafny.Set(coll5_)
            compr_4_: int
            for compr_4_ in (iife5_()
            ).Elements:
                d_14_v0_: int = compr_4_
                if (d_14_v0_) in (iife6_()
                ):
                    coll4_[(d_14_v0_) * (len(_dafny.Set({False, not(True), False, True})))] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_16_i0_ in range(default__.abs(972))])
            return _dafny.Map(coll4_)
        if (len(iife4_()
        )) <= ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "delwixswv"))))):
            return _dafny.CodePoint('v')
        elif True:
            return _dafny.CodePoint('m')

    @staticmethod
    def fm19(p0, globalState):
        return ((D7_DC25(_dafny.SeqWithoutIsStrInference([D3_DC9(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('f')])) for d_17_i0_ in range(default__.abs(-252))]))).cf41) + (_dafny.SeqWithoutIsStrInference([D3_DC9(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_18_i1_ in range(default__.abs(356))])), D3_DC9(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')]))]))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxqwlk"))): _dafny.CodePoint('w')})).keys.Elements:
                d_19_v0_: int = compr_7_
                if (d_19_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxqwlk"))): _dafny.CodePoint('w')})):
                    coll7_[(d_19_v0_) * (757)] = True
            return _dafny.Map(coll7_)
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife7_()
 for d_20_i0_ in range(default__.abs(599))]))

    @staticmethod
    def fm21(p0, globalState):
        return D6_DC20((_dafny.Map({not(True): not(False)})) | (_dafny.Map({False: False})))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        hi0_ = p0
        for d_21_i0_ in range(p0, hi0_):
            d_22_v0_: str
            d_22_v0_ = _dafny.CodePoint('x')
            d_23_v1_: _dafny.Map
            d_23_v1_ = _dafny.Map({d_22_v0_: d_22_v0_})
            d_24_v2_: _dafny.Map
            d_24_v2_ = _dafny.Map({d_23_v1_: p2})
            index0_ = default__.safeIndex(412, (p1).length(0))
            (p1)[index0_] = p0
            d_25_v3_: _dafny.Seq
            d_25_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nodnggll"))
            index1_ = default__.safeIndex(412, (p1).length(0))
            rhs0_ = d_24_v2_
            rhs1_ = len(_dafny.SeqWithoutIsStrInference([d_22_v0_ for d_26_i1_ in range(default__.abs(246))]))
            rhs2_ = len(d_25_v3_)
            lhs0_ = p1
            lhs1_ = default__.safeIndex(412, (p1).length(0))
            d_24_v2_ = rhs0_
            r0 = rhs1_
            lhs0_[lhs1_] = rhs2_
            d_27_v4_: D5
            d_27_v4_ = D5_DC18(p0, p2)
            pat_let_tv0_ = p2
            pat_let_tv1_ = p0
            pat_let_tv2_ = p2
            index2_ = default__.safeIndex(412, (p1).length(0))
            def iife10_(_pat_let2_0):
                def iife11_(d_28_dt__update__tmp_h0_):
                    def iife12_(_pat_let3_0):
                        def iife13_(d_29_dt__update_hcf28_h0_):
                            return D5_DC18((d_28_dt__update__tmp_h0_).cf27, d_29_dt__update_hcf28_h0_)
                        return iife13_(_pat_let3_0)
                    return iife12_(pat_let_tv0_)
                return iife11_(_pat_let2_0)
            def iife9_(_pat_let1_0):
                def iife14_(d_30_dt__update__tmp_h1_):
                    def iife15_(_pat_let4_0):
                        def iife16_(d_31_dt__update_hcf27_h0_):
                            return D5_DC18(d_31_dt__update_hcf27_h0_, (d_30_dt__update__tmp_h1_).cf28)
                        return iife16_(_pat_let4_0)
                    return iife15_((0) - (pat_let_tv1_))
                return iife14_(_pat_let1_0)
            def iife8_(_pat_let0_0):
                def iife17_(d_32_dt__update__tmp_h2_):
                    def iife18_(_pat_let5_0):
                        def iife19_(d_33_dt__update_hcf28_h1_):
                            return D5_DC18((d_32_dt__update__tmp_h2_).cf27, d_33_dt__update_hcf28_h1_)
                        return iife19_(_pat_let5_0)
                    return iife18_(pat_let_tv2_)
                return iife17_(_pat_let0_0)
            rhs3_ = p0
            rhs4_ = default__.safeDivisionInt((d_27_v4_).cf27, (p1)[default__.safeIndex(412, (p1).length(0))])
            rhs5_ = iife8_(iife9_(iife10_(d_27_v4_)))
            lhs2_ = p1
            lhs3_ = default__.safeIndex(412, (p1).length(0))
            lhs2_[lhs3_] = rhs3_
            r0 = rhs4_
            d_27_v4_ = rhs5_
            d_34_v5_: C1
            nw0_ = C1()
            nw0_.ctor__(p2, (p1)[default__.safeIndex(412, (p1).length(0))], p2)
            d_34_v5_ = nw0_
            d_35_v6_: _dafny.Array
            nw1_ = _dafny.Array(None, 5)
            nw1_[int(0)] = d_34_v5_
            nw1_[int(1)] = d_34_v5_
            nw1_[int(2)] = d_34_v5_
            nw1_[int(3)] = d_34_v5_
            nw1_[int(4)] = d_34_v5_
            d_35_v6_ = nw1_
            index3_ = default__.safeIndex(201, (d_35_v6_).length(0))
            (d_35_v6_)[index3_] = d_34_v5_
            index4_ = default__.safeIndex(201, (d_35_v6_).length(0))
            (d_35_v6_)[index4_] = d_34_v5_
            (globalState).f3 = not((d_22_v0_) in (d_25_v3_))
        hi1_ = p0
        for d_36_i2_ in range(p0, hi1_):
            d_37_v7_: _dafny.Array
            nw2_ = _dafny.Array(None, 14)
            nw2_[int(0)] = p2
            nw2_[int(1)] = True
            nw2_[int(2)] = p2
            nw2_[int(3)] = p2
            nw2_[int(4)] = default__.fm2(globalState)
            nw2_[int(5)] = p2
            nw2_[int(6)] = p2
            nw2_[int(7)] = p2
            nw2_[int(8)] = p2
            nw2_[int(9)] = p2
            nw2_[int(10)] = p2
            nw2_[int(11)] = True
            nw2_[int(12)] = p2
            nw2_[int(13)] = p2
            d_37_v7_ = nw2_
            d_38_v8_: _dafny.Seq
            d_38_v8_ = _dafny.SeqWithoutIsStrInference([d_37_v7_, d_37_v7_])
            d_39_v9_: _dafny.Array
            nw3_ = _dafny.Array(_dafny.Map({}), 14)
            d_39_v9_ = nw3_
            d_40_v10_: _dafny.Map
            d_40_v10_ = _dafny.Map({d_38_v8_: d_39_v9_})
            d_40_v10_ = (d_40_v10_).set((d_38_v8_) + (d_38_v8_), d_39_v9_)
            d_41_v11_: _dafny.Seq
            d_41_v11_ = _dafny.SeqWithoutIsStrInference([d_36_i2_])
            d_42_v12_: _dafny.Set
            d_42_v12_ = _dafny.Set({(0) - (p0), p0, (d_41_v11_)[default__.safeIndex(d_36_i2_, len(d_41_v11_))], p0, d_36_i2_})
            d_43_v13_: D3
            d_43_v13_ = D3_DC10(len(d_42_v12_), p1, p2)
            source0_ = d_43_v13_
            if source0_.is_DC10:
                d_44___mcc_h0_ = source0_.cf11
                d_45___mcc_h1_ = source0_.cf12
                d_46___mcc_h2_ = source0_.cf13
                d_47_cf13_ = d_46___mcc_h2_
                d_48_cf12_ = d_45___mcc_h1_
                d_49_cf11_ = d_44___mcc_h0_
                d_50_v14_: _dafny.Map
                d_50_v14_ = _dafny.Map({p2: p0})
                index5_ = default__.safeIndex(352, (d_48_cf12_).length(0))
                (d_48_cf12_)[index5_] = (d_49_cf11_ if p2 else len(_dafny.Map({p2: len(d_50_v14_)})))
                index6_ = default__.safeIndex(352, (d_48_cf12_).length(0))
                (d_48_cf12_)[index6_] = ((d_36_i2_) * (p0) if False else p0)
                d_51_v15_: C2
                nw4_ = C2()
                nw4_.ctor__(True, d_47_cf13_)
                d_51_v15_ = nw4_
                d_51_v15_ = d_51_v15_
                index7_ = default__.safeIndex(365, (d_37_v7_).length(0))
                (d_37_v7_)[index7_] = d_51_v15_.f10
                index8_ = default__.safeIndex(365, (d_37_v7_).length(0))
                (d_37_v7_)[index8_] = p2
                d_52_v16_: _dafny.Array
                nw5_ = _dafny.Array(False, 14)
                d_52_v16_ = nw5_
                d_53_v17_: D0
                d_53_v17_ = D0_DC1(d_51_v15_.f10)
                d_54_v18_: _dafny.Set
                d_54_v18_ = _dafny.Set({(d_53_v17_).cf1, d_51_v15_.f10})
                d_55_v19_: bool
                d_56_v20_: bool
                d_57_v21_: _dafny.Map
                out0_: bool
                out1_: bool
                out2_: _dafny.Map
                out0_, out1_, out2_ = (d_51_v15_).m3((p0) - ((d_48_cf12_)[default__.safeIndex(352, (d_48_cf12_).length(0))]), d_52_v16_, True, d_54_v18_, globalState)
                d_55_v19_ = out0_
                d_56_v20_ = out1_
                d_57_v21_ = out2_
            elif source0_.is_DC11:
                d_58___mcc_h3_ = source0_.cf14
                d_59___mcc_h4_ = source0_.cf15
                d_60___mcc_h5_ = source0_.cf16
                d_61___mcc_h6_ = source0_.cf17
                d_62_cf17_ = d_61___mcc_h6_
                d_63_cf16_ = d_60___mcc_h5_
                d_64_cf15_ = d_59___mcc_h4_
                d_65_cf14_ = d_58___mcc_h3_
                d_66_v23_: _dafny.MultiSet
                def iife20_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(-756, 716):
                        d_67_v22_: int = compr_8_
                        if ((-756) <= (d_67_v22_)) and ((d_67_v22_) < (716)):
                            coll8_[default__.safeModuloInt(d_67_v22_, len(d_64_cf15_))] = not(p2)
                    return _dafny.Map(coll8_)
                d_66_v23_ = _dafny.MultiSet([iife20_()
                ])
                d_68_v24_: D2
                d_68_v24_ = D2_DC8((d_62_cf17_).f13, d_63_cf16_, (d_62_cf17_).f13)
                pat_let_tv3_ = d_62_cf17_
                pat_let_tv4_ = d_63_cf16_
                pat_let_tv5_ = d_62_cf17_
                def iife22_(_pat_let7_0):
                    def iife23_(d_69_dt__update__tmp_h3_):
                        def iife24_(_pat_let8_0):
                            def iife25_(d_70_dt__update_hcf7_h0_):
                                def iife26_(_pat_let9_0):
                                    def iife27_(d_71_dt__update_hcf8_h0_):
                                        return D2_DC8(d_70_dt__update_hcf7_h0_, d_71_dt__update_hcf8_h0_, (d_69_dt__update__tmp_h3_).cf9)
                                    return iife27_(_pat_let9_0)
                                return iife26_(pat_let_tv4_)
                            return iife25_(_pat_let8_0)
                        return iife24_((pat_let_tv3_).f13)
                    return iife23_(_pat_let7_0)
                def iife21_(_pat_let6_0):
                    def iife28_(d_72_dt__update__tmp_h4_):
                        def iife29_(_pat_let10_0):
                            def iife30_(d_73_dt__update_hcf7_h1_):
                                return D2_DC8(d_73_dt__update_hcf7_h1_, (d_72_dt__update__tmp_h4_).cf8, (d_72_dt__update__tmp_h4_).cf9)
                            return iife30_(_pat_let10_0)
                        return iife29_((pat_let_tv5_).f13)
                    return iife28_(_pat_let6_0)
                rhs6_ = (iife21_(iife22_(d_68_v24_))).cf7
                rhs7_ = default__.fm20((d_62_cf17_).f13, d_41_v11_, (d_63_cf16_) == (d_62_cf17_.f14), globalState)
                rhs8_ = (d_64_cf15_) < ((d_64_cf15_) + (d_64_cf15_))
                lhs4_ = globalState
                lhs5_ = globalState
                lhs4_.f3 = rhs6_
                d_66_v23_ = rhs7_
                lhs5_.f3 = rhs8_
                d_74_v25_: _dafny.Array
                nw6_ = _dafny.Array(int(0), 8)
                d_74_v25_ = nw6_
                d_74_v25_ = p1
                d_63_cf16_ = (d_36_i2_ if (p2 if (d_62_cf17_).f13 else (d_62_cf17_).f13) else d_65_cf14_)
                d_75_v26_: _dafny.Map
                d_75_v26_ = _dafny.Map({p2: (d_62_cf17_).f13})
                d_76_v27_: _dafny.Map
                d_76_v27_ = _dafny.Map({(len(d_75_v26_)) <= (d_65_cf14_): not(p2)})
                d_77_v28_: _dafny.MultiSet
                d_77_v28_ = _dafny.MultiSet([p2])
                d_76_v27_ = (d_76_v27_).set((p2) not in (d_77_v28_), (d_43_v13_).cf13)
            elif source0_.is_DC9:
                d_78___mcc_h7_ = source0_.cf10
                d_79_cf10_ = d_78___mcc_h7_
                d_80_v29_: _dafny.Set
                d_80_v29_ = _dafny.Set({p2, p2})
                (globalState).f0 = not(((d_41_v11_)[default__.safeIndex(p0, len(d_41_v11_))]) >= ((p0) * (len(d_80_v29_))))
                d_81_v30_: _dafny.Array
                def lambda0_(d_82_p2_):
                    def lambda1_(d_83_i3_):
                        return (_dafny.CodePoint('c') if d_82_p2_ else _dafny.CodePoint('x'))

                    return lambda1_

                init0_ = lambda0_(p2)
                nw7_ = _dafny.Array(None, 27)
                for i0_0_ in range(nw7_.length(0)):
                    nw7_[i0_0_] = init0_(i0_0_)
                d_81_v30_ = nw7_
                index9_ = default__.safeIndex(82, (d_81_v30_).length(0))
                (d_81_v30_)[index9_] = _dafny.CodePoint('q')
                d_84_v31_: str
                d_84_v31_ = _dafny.CodePoint('u')
                index10_ = default__.safeIndex(82, (d_81_v30_).length(0))
                (d_81_v30_)[index10_] = d_84_v31_
                d_39_v9_ = d_39_v9_
                index11_ = default__.safeIndex(451, (p1).length(0))
                (p1)[index11_] = p0
                index12_ = default__.safeIndex(451, (p1).length(0))
                (p1)[index12_] = p0
            elif True:
                d_85___mcc_h8_ = source0_.cf18
                d_86_cf18_ = d_85___mcc_h8_
                d_37_v7_ = d_37_v7_
                index13_ = default__.safeIndex(870, (d_37_v7_).length(0))
                (d_37_v7_)[index13_] = not(p2)
                d_87_v32_: _dafny.MultiSet
                d_87_v32_ = _dafny.MultiSet([p2, p2])
                d_88_v33_: _dafny.Seq
                d_88_v33_ = _dafny.SeqWithoutIsStrInference([d_87_v32_])
                index14_ = default__.safeIndex(870, (d_37_v7_).length(0))
                (d_37_v7_)[index14_] = ((d_88_v33_) + (d_88_v33_)) < (d_88_v33_)
                r0 = p0
                d_89_v34_: _dafny.Array
                nw8_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_89_v34_ = nw8_
                d_90_v35_: _dafny.Array
                nw9_ = _dafny.Array(False, 21)
                d_90_v35_ = nw9_
                index15_ = default__.safeIndex(666, (d_89_v34_).length(0))
                (d_89_v34_)[index15_] = d_90_v35_
                d_91_v36_: str
                d_91_v36_ = _dafny.CodePoint('l')
                index16_ = default__.safeIndex(666, (d_89_v34_).length(0))
                (d_89_v34_)[index16_] = (d_38_v8_)[default__.safeIndex(default__.fm3(d_91_v36_, True, p0, globalState), len(d_38_v8_))]
            d_92_v37_: D4
            d_92_v37_ = D4_DC15((p0) + (d_36_i2_), p2)
            rhs9_ = p2
            rhs10_ = d_92_v37_
            lhs6_ = globalState
            lhs6_.f3 = rhs9_
            d_92_v37_ = rhs10_
            d_93_v38_: _dafny.Seq
            d_93_v38_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_94_v39_: C2
            nw10_ = C2()
            nw10_.ctor__((d_93_v38_)[default__.safeIndex(len(d_93_v38_), len(d_93_v38_))], p2)
            d_94_v39_ = nw10_
        hi2_ = p0
        for d_95_i4_ in range((p0) + (p0), hi2_):
            (globalState).f0 = not(p2)
            r0 = (d_95_i4_) - (len(default__.fm11(False, p2, globalState)))
            (globalState).f0 = ((p2 if not(not(p2)) else p2)) and (not((d_95_i4_) != (p0)))
            (globalState).f3 = p2
        hi3_ = p0
        for d_96_i5_ in range(p0, hi3_):
            d_97_v40_: _dafny.Seq
            d_97_v40_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
            d_98_v41_: C1
            nw11_ = C1()
            nw11_.ctor__(not (p2) or (default__.fm2(globalState)), default__.safeModuloInt(len(d_97_v40_), d_96_i5_), (p2 if p2 else p2))
            d_98_v41_ = nw11_
            d_98_v41_ = d_98_v41_
            d_99_v42_: _dafny.Array
            def lambda2_(d_100_v41_):
                def lambda3_(d_101_i6_):
                    return (d_100_v41_).f11

                return lambda3_

            init1_ = lambda2_(d_98_v41_)
            nw12_ = _dafny.Array(None, 15)
            for i0_1_ in range(nw12_.length(0)):
                nw12_[i0_1_] = init1_(i0_1_)
            d_99_v42_ = nw12_
            d_102_v43_: _dafny.Map
            d_102_v43_ = _dafny.Map({d_99_v42_: (d_98_v41_).f11})
            d_102_v43_ = (d_102_v43_).set(d_99_v42_, p2)
            d_103_v44_: _dafny.MultiSet
            d_103_v44_ = _dafny.MultiSet([_dafny.Set({p0})])
            (globalState).f3 = ((d_103_v44_).issubset(d_103_v44_)) == ((d_98_v41_).f11)
            d_104_v45_: C3
            nw13_ = C3()
            nw13_.ctor__((d_98_v41_).f11)
            d_104_v45_ = nw13_
            d_105_v46_: _dafny.Map
            d_105_v46_ = _dafny.Map({d_99_v42_: _dafny.Map({d_104_v45_: (d_98_v41_).f11})})
            d_106_v47_: _dafny.Map
            d_106_v47_ = _dafny.Map({d_104_v45_: (d_98_v41_).f11})
            d_105_v46_ = (d_105_v46_).set(d_99_v42_, d_106_v47_)
        d_107_i7_: int
        d_107_i7_ = 0
        with _dafny.label("0"):
            while not(not(p2)):
                with _dafny.c_label("0"):
                    if (d_107_i7_) >= (100):
                        raise _dafny.Break("0")
                    d_107_i7_ = (d_107_i7_) + (1)
                    (globalState).f0 = p2
                    d_108_v48_: _dafny.MultiSet
                    d_108_v48_ = _dafny.MultiSet([p0])
                    d_109_v49_: _dafny.MultiSet
                    d_109_v49_ = _dafny.MultiSet([d_108_v48_])
                    d_110_v50_: D0
                    d_110_v50_ = D0_DC2(p2, p0, d_109_v49_)
                    d_111_v51_: _dafny.Array
                    nw14_ = _dafny.Array(None, 21)
                    nw14_[int(0)] = p2
                    nw14_[int(1)] = p2
                    nw14_[int(2)] = p2
                    nw14_[int(3)] = p2
                    nw14_[int(4)] = p2
                    nw14_[int(5)] = p2
                    nw14_[int(6)] = p2
                    nw14_[int(7)] = p2
                    nw14_[int(8)] = p2
                    nw14_[int(9)] = p2
                    nw14_[int(10)] = p2
                    nw14_[int(11)] = p2
                    nw14_[int(12)] = p2
                    nw14_[int(13)] = p2
                    nw14_[int(14)] = not(p2)
                    nw14_[int(15)] = (d_110_v50_).cf2
                    nw14_[int(16)] = p2
                    nw14_[int(17)] = True
                    nw14_[int(18)] = p2
                    nw14_[int(19)] = p2
                    nw14_[int(20)] = p2
                    d_111_v51_ = nw14_
                    d_112_v52_: _dafny.Seq
                    d_112_v52_ = _dafny.SeqWithoutIsStrInference([d_111_v51_, d_111_v51_, d_111_v51_])
                    d_113_v53_: _dafny.Seq
                    d_113_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nw"))
                    d_114_v54_: _dafny.Seq
                    d_114_v54_ = _dafny.SeqWithoutIsStrInference([p0, len((d_112_v52_) + ((d_112_v52_).set(default__.safeIndex(len(d_113_v53_), len(d_112_v52_)), d_111_v51_)))])
                    d_115_v55_: _dafny.Seq
                    d_115_v55_ = _dafny.SeqWithoutIsStrInference([d_114_v54_])
                    d_114_v54_ = (d_115_v55_)[default__.safeIndex(p0, len(d_115_v55_))]
                    d_116_v56_: _dafny.Seq
                    d_116_v56_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                    if (default__.fm10(globalState)) == (d_116_v56_):
                        (globalState).f0 = p2
                        d_117_v57_: _dafny.Map
                        d_117_v57_ = _dafny.Map({(p0) == (p0): p2})
                        d_118_v58_: str
                        d_118_v58_ = _dafny.CodePoint('k')
                        d_119_v59_: _dafny.Map
                        d_119_v59_ = _dafny.Map({p2: d_118_v58_})
                        rhs11_ = d_111_v51_
                        rhs12_ = default__.fm13(p2, (p2 if p2 else p2), p2, p0, globalState)
                        rhs13_ = (default__.fm21(p2, globalState)).cf30
                        rhs14_ = p2
                        rhs15_ = d_119_v59_
                        lhs7_ = globalState
                        d_111_v51_ = rhs11_
                        d_114_v54_ = rhs12_
                        d_117_v57_ = rhs13_
                        lhs7_.f3 = rhs14_
                        d_119_v59_ = rhs15_
                        d_120_v60_: _dafny.Array
                        nw15_ = _dafny.Array(_dafny.Array(None, 0), 26)
                        d_120_v60_ = nw15_
                        d_121_v61_: _dafny.Array
                        nw16_ = _dafny.Array(None, 7)
                        d_121_v61_ = nw16_
                        index17_ = default__.safeIndex(723, (d_120_v60_).length(0))
                        (d_120_v60_)[index17_] = d_121_v61_
                        index18_ = default__.safeIndex(723, (d_120_v60_).length(0))
                        (d_120_v60_)[index18_] = d_121_v61_
                        d_122_v62_: _dafny.MultiSet
                        d_122_v62_ = _dafny.MultiSet([p2, p2, p2, p2, p2])
                        d_123_v63_: _dafny.Set
                        d_123_v63_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([p0, (d_122_v62_).cardinality]))})
                        d_124_v64_: C0
                        nw17_ = C0()
                        nw17_.ctor__(True, (d_114_v54_)[default__.safeIndex(p0, len(d_114_v54_))], (d_123_v63_).ispropersubset(d_123_v63_))
                        d_124_v64_ = nw17_
                        d_124_v64_ = d_124_v64_
                        d_125_v65_: _dafny.Set
                        d_125_v65_ = _dafny.Set({d_108_v48_, d_108_v48_, (_dafny.MultiSet(d_114_v54_)) | (d_108_v48_)})
                        d_126_v66_: _dafny.Array
                        nw18_ = _dafny.Array(_dafny.Map({}), 19)
                        d_126_v66_ = nw18_
                        d_127_v67_: _dafny.Map
                        d_127_v67_ = _dafny.Map({(d_124_v64_).f13: d_116_v56_})
                        index19_ = default__.safeIndex(214, (d_126_v66_).length(0))
                        (d_126_v66_)[index19_] = d_127_v67_
                        index20_ = default__.safeIndex(214, (d_126_v66_).length(0))
                        rhs16_ = default__.safeDivisionInt(p0, (0) - (p0))
                        rhs17_ = d_125_v65_
                        rhs18_ = (d_113_v53_) + (d_113_v53_)
                        rhs19_ = _dafny.MultiSet([d_124_v64_.f14, default__.safeModuloInt(d_124_v64_.f14, d_124_v64_.f14)])
                        rhs20_ = ((d_127_v67_) | (d_127_v67_)) | (d_127_v67_)
                        lhs8_ = d_124_v64_
                        lhs9_ = d_126_v66_
                        lhs10_ = default__.safeIndex(214, (d_126_v66_).length(0))
                        lhs8_.f14 = rhs16_
                        d_125_v65_ = rhs17_
                        d_113_v53_ = rhs18_
                        d_108_v48_ = rhs19_
                        lhs9_[lhs10_] = rhs20_
                    elif True:
                        d_128_v68_: _dafny.MultiSet
                        d_128_v68_ = _dafny.MultiSet([d_114_v54_, _dafny.SeqWithoutIsStrInference([p0])])
                        (globalState).f0 = not((p0) >= ((d_128_v68_).cardinality))
                        d_114_v54_ = d_114_v54_
                        r0 = (p0) * (p0)
                        d_129_v69_: str
                        d_129_v69_ = _dafny.CodePoint('l')
                        index21_ = default__.safeIndex(606, (p1).length(0))
                        (p1)[index21_] = (len(d_113_v53_)) + (default__.fm3(d_129_v69_, True, len(d_113_v53_), globalState))
                        index22_ = default__.safeIndex(606, (p1).length(0))
                        (p1)[index22_] = p0
                        d_130_v70_: C0
                        nw19_ = C0()
                        nw19_.ctor__(False, p0, p2)
                        d_130_v70_ = nw19_
                        d_131_v71_: T0
                        nw20_ = C3()
                        nw20_.ctor__((d_130_v70_).f13)
                        d_131_v71_ = nw20_
                        d_132_v72_: _dafny.Map
                        d_132_v72_ = _dafny.Map({d_130_v70_: d_131_v71_})
                        d_133_v73_: _dafny.Map
                        d_133_v73_ = _dafny.Map({p2: (d_132_v72_) | (_dafny.Map({d_130_v70_: d_131_v71_}))})
                        d_133_v73_ = (d_133_v73_).set(not((d_130_v70_).f13), _dafny.Map({d_130_v70_: d_131_v71_}))
                    index23_ = default__.safeIndex(607, (d_111_v51_).length(0))
                    (d_111_v51_)[index23_] = p2
                    index24_ = default__.safeIndex(607, (d_111_v51_).length(0))
                    (d_111_v51_)[index24_] = ((p0) >= (p0)) or (p2)
                    pass
            pass
        hi4_ = p0
        for d_134_i8_ in range(702, hi4_):
            (globalState).f0 = (d_134_i8_) <= (d_134_i8_)
            d_135_v74_: _dafny.Array
            nw21_ = _dafny.Array(False, 4)
            d_135_v74_ = nw21_
            index25_ = default__.safeIndex(556, (d_135_v74_).length(0))
            (d_135_v74_)[index25_] = p2
            index26_ = default__.safeIndex(733, (d_135_v74_).length(0))
            (d_135_v74_)[index26_] = False
            d_136_v75_: _dafny.Array
            nw22_ = _dafny.Array(_dafny.Set({}), 21)
            d_136_v75_ = nw22_
            d_137_v76_: _dafny.Seq
            d_137_v76_ = _dafny.SeqWithoutIsStrInference([p2])
            d_138_v77_: _dafny.Set
            d_138_v77_ = _dafny.Set({p2, p2, p2})
            index27_ = default__.safeIndex(960, (d_136_v75_).length(0))
            (d_136_v75_)[index27_] = (_dafny.Set({p2, p2, (d_137_v76_)[default__.safeIndex(p0, len(d_137_v76_))], p2})) | (d_138_v77_)
            d_139_v78_: T0
            nw23_ = C1()
            nw23_.ctor__(p2, default__.safeModuloInt(d_134_i8_, p0), p2)
            d_139_v78_ = nw23_
            d_140_v79_: _dafny.Set
            d_140_v79_ = _dafny.Set({d_134_i8_, p0, p0})
            d_141_v80_: _dafny.MultiSet
            d_141_v80_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdndrwawf")))])
            d_142_v81_: _dafny.Map
            d_142_v81_ = _dafny.Map({(((d_141_v80_)[d_134_i8_] if (d_134_i8_) in (d_141_v80_) else d_134_i8_)) - (d_134_i8_): (p0) * (-532)})
            d_143_v82_: _dafny.Seq
            d_143_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koccov"))
            d_144_v83_: _dafny.Map
            d_144_v83_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([True, True, p2, (d_139_v78_).f9])})
            d_145_v84_: C2
            nw24_ = C2()
            nw24_.ctor__((d_139_v78_).f9, (d_139_v78_).f9)
            d_145_v84_ = nw24_
            d_146_v85_: _dafny.MultiSet
            d_146_v85_ = _dafny.MultiSet([d_145_v84_])
            index28_ = default__.safeIndex(556, (d_135_v74_).length(0))
            index29_ = default__.safeIndex(733, (d_135_v74_).length(0))
            index30_ = default__.safeIndex(960, (d_136_v75_).length(0))
            rhs21_ = (d_140_v79_).ispropersubset(d_140_v79_)
            rhs22_ = p2
            rhs23_ = ((d_142_v81_)[default__.safeModuloInt((0) - (-199), len(d_143_v82_))] if (default__.safeModuloInt((0) - (-199), len(d_143_v82_))) in (d_142_v81_) else ((d_139_v78_).fm4(d_144_v83_, p2, (d_139_v78_).f9, p0, globalState)) - ((d_146_v85_).cardinality))
            rhs24_ = d_138_v77_
            rhs25_ = d_139_v78_
            lhs11_ = d_135_v74_
            lhs12_ = default__.safeIndex(556, (d_135_v74_).length(0))
            lhs13_ = d_135_v74_
            lhs14_ = default__.safeIndex(733, (d_135_v74_).length(0))
            lhs15_ = d_136_v75_
            lhs16_ = default__.safeIndex(960, (d_136_v75_).length(0))
            lhs11_[lhs12_] = rhs21_
            lhs13_[lhs14_] = rhs22_
            r0 = rhs23_
            lhs15_[lhs16_] = rhs24_
            d_139_v78_ = rhs25_
            d_147_v86_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_147_v86_ = nw25_
            d_148_v87_: D4
            d_148_v87_ = D4_DC13(d_147_v86_)
            d_149_v88_: D4
            d_149_v88_ = D4_DC16(d_148_v87_)
            d_149_v88_ = d_149_v88_
            index31_ = default__.safeIndex(556, (d_135_v74_).length(0))
            (d_135_v74_)[index31_] = (d_134_i8_) == (-200)
        r0 = default__.safeDivisionInt((p0) * (p0), p0)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_150_v0_: int
        d_150_v0_ = 636
        d_151_v1_: _dafny.Seq
        d_151_v1_ = _dafny.SeqWithoutIsStrInference([d_150_v0_, d_150_v0_])
        d_152_v2_: bool
        d_152_v2_ = True
        d_153_v3_: _dafny.Map
        d_153_v3_ = _dafny.Map({d_152_v2_: d_151_v1_})
        d_154_v4_: _dafny.MultiSet
        d_154_v4_ = _dafny.MultiSet([d_151_v1_, d_151_v1_, ((d_153_v3_)[not(False)] if (not(False)) in (d_153_v3_) else _dafny.SeqWithoutIsStrInference([d_150_v0_])), d_151_v1_, d_151_v1_])
        d_155_v5_: _dafny.Array
        nw26_ = _dafny.Array(False, 26)
        d_155_v5_ = nw26_
        d_156_v6_: _dafny.MultiSet
        d_156_v6_ = _dafny.MultiSet([d_155_v5_, d_155_v5_])
        d_157_v7_: D0
        d_157_v7_ = D0_DC0((0) - (d_150_v0_))
        d_158_v8_: _dafny.Seq
        d_158_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyuhsbbd"))
        d_159_v9_: _dafny.Map
        d_159_v9_ = _dafny.Map({d_152_v2_: d_150_v0_})
        d_160_v10_: _dafny.Array
        nw27_ = _dafny.Array(None, 17)
        nw27_[int(0)] = d_150_v0_
        nw27_[int(1)] = d_150_v0_
        nw27_[int(2)] = len(_dafny.Map({d_152_v2_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqsopcbd")))}))
        nw27_[int(3)] = (0) - (d_150_v0_)
        nw27_[int(4)] = d_150_v0_
        nw27_[int(5)] = d_150_v0_
        nw27_[int(6)] = (d_156_v6_).cardinality
        nw27_[int(7)] = (d_157_v7_).cf0
        nw27_[int(8)] = len(d_151_v1_)
        nw27_[int(9)] = len(d_158_v8_)
        nw27_[int(10)] = d_150_v0_
        nw27_[int(11)] = d_150_v0_
        nw27_[int(12)] = d_150_v0_
        nw27_[int(13)] = d_150_v0_
        nw27_[int(14)] = d_150_v0_
        nw27_[int(15)] = len(d_159_v9_)
        nw27_[int(16)] = d_150_v0_
        d_160_v10_ = nw27_
        d_161_globalState_: GlobalState
        nw28_ = GlobalState()
        nw28_.ctor__(True, 463, True, False, False, d_154_v4_, d_160_v10_, _dafny.CodePoint('s'), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_162_i0_ in range(default__.abs(798))]))
        d_161_globalState_ = nw28_
        d_163_v11_: str
        d_163_v11_ = _dafny.CodePoint('p')
        pat_let_tv6_ = d_158_v8_
        pat_let_tv7_ = d_158_v8_
        pat_let_tv8_ = d_158_v8_
        pat_let_tv9_ = d_158_v8_
        pat_let_tv10_ = d_158_v8_
        pat_let_tv11_ = d_158_v8_
        pat_let_tv12_ = d_158_v8_
        pat_let_tv13_ = d_158_v8_
        pat_let_tv14_ = d_158_v8_
        pat_let_tv15_ = d_158_v8_
        pat_let_tv16_ = d_158_v8_
        pat_let_tv17_ = d_158_v8_
        def lambda4_(source1_):
            if source1_.is_DC1:
                d_164___mcc_h0_ = source1_.cf1
                d_165_cf1_ = d_164___mcc_h0_
                return pat_let_tv6_
            elif source1_.is_DC2:
                d_166___mcc_h1_ = source1_.cf2
                d_167___mcc_h2_ = source1_.cf3
                d_168___mcc_h3_ = source1_.cf4
                d_169_cf4_ = d_168___mcc_h3_
                d_170_cf3_ = d_167___mcc_h2_
                d_171_cf2_ = d_166___mcc_h1_
                return (pat_let_tv7_) + (pat_let_tv8_)
            elif source1_.is_DC3:
                return (pat_let_tv9_) + (pat_let_tv10_)
            elif True:
                d_172___mcc_h4_ = source1_.cf0
                d_173_cf0_ = d_172___mcc_h4_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbtq"))) + (pat_let_tv11_)

        def lambda5_(source2_):
            if source2_.is_DC1:
                d_174___mcc_h5_ = source2_.cf1
                d_175_cf1_ = d_174___mcc_h5_
                return pat_let_tv12_
            elif source2_.is_DC2:
                d_176___mcc_h6_ = source2_.cf2
                d_177___mcc_h7_ = source2_.cf3
                d_178___mcc_h8_ = source2_.cf4
                d_179_cf4_ = d_178___mcc_h8_
                d_180_cf3_ = d_177___mcc_h7_
                d_181_cf2_ = d_176___mcc_h6_
                return (pat_let_tv13_) + (pat_let_tv14_)
            elif source2_.is_DC3:
                return (pat_let_tv15_) + (pat_let_tv16_)
            elif True:
                d_182___mcc_h9_ = source2_.cf0
                d_183_cf0_ = d_182___mcc_h9_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbtq"))) + (pat_let_tv17_)

        d_158_v8_ = (lambda4_(default__.fm0(d_152_v2_, d_161_globalState_))).set(default__.safeIndex(d_150_v0_, len(lambda5_(default__.fm0(d_152_v2_, d_161_globalState_)))), d_163_v11_)
        hi5_ = d_150_v0_
        for d_184_i1_ in range(default__.safeModuloInt(d_150_v0_, d_150_v0_), hi5_):
            source3_ = default__.fm1(d_158_v8_, True, d_161_globalState_)
            if source3_.is_DC1:
                d_185___mcc_h10_ = source3_.cf1
                d_186_cf1_ = d_185___mcc_h10_
                index32_ = default__.safeIndex(185, (d_155_v5_).length(0))
                (d_155_v5_)[index32_] = d_186_cf1_
                index33_ = default__.safeIndex(185, (d_155_v5_).length(0))
                (d_155_v5_)[index33_] = default__.fm2(d_161_globalState_)
                d_187_v12_: _dafny.Map
                d_187_v12_ = _dafny.Map({d_150_v0_: d_158_v8_})
                d_188_v13_: D0
                d_188_v13_ = D0_DC1((d_187_v12_) != (d_187_v12_))
                d_188_v13_ = d_188_v13_
                d_189_v14_: _dafny.Map
                d_189_v14_ = _dafny.Map({413: d_150_v0_})
                d_190_v15_: _dafny.Map
                d_190_v15_ = _dafny.Map({d_184_i1_: d_189_v14_})
                d_191_v16_: _dafny.Set
                d_191_v16_ = _dafny.Set({len(d_158_v8_)})
                d_150_v0_ = default__.safeDivisionInt((d_151_v1_)[default__.safeIndex(len((((d_190_v15_)[d_150_v0_] if (d_150_v0_) in (d_190_v15_) else d_189_v14_)).set(len(d_191_v16_), d_184_i1_)), len(d_151_v1_))], d_150_v0_)
                d_192_v17_: _dafny.MultiSet
                d_192_v17_ = _dafny.MultiSet([(d_155_v5_)[default__.safeIndex(185, (d_155_v5_).length(0))], (d_155_v5_)[default__.safeIndex(185, (d_155_v5_).length(0))]])
                d_193_v18_: _dafny.Seq
                d_193_v18_ = _dafny.SeqWithoutIsStrInference([(d_192_v17_).issubset(d_192_v17_)])
                d_150_v0_ = (0) - (len(d_193_v18_))
            elif source3_.is_DC2:
                d_194___mcc_h11_ = source3_.cf2
                d_195___mcc_h12_ = source3_.cf3
                d_196___mcc_h13_ = source3_.cf4
                d_197_cf4_ = d_196___mcc_h13_
                d_198_cf3_ = d_195___mcc_h12_
                d_199_cf2_ = d_194___mcc_h11_
                d_200_v19_: _dafny.Map
                d_200_v19_ = _dafny.Map({len(_dafny.Set({d_199_cf2_, d_199_cf2_})): d_199_cf2_})
                d_201_v20_: int
                out3_: int
                out3_ = default__.m0(d_198_cf3_, d_160_v10_, d_152_v2_, (d_200_v19_) | (d_200_v19_), d_161_globalState_)
                d_201_v20_ = out3_
                nw29_ = _dafny.Array(False, 20)
                d_155_v5_ = nw29_
                (d_161_globalState_).f7 = d_163_v11_
                (d_161_globalState_).f0 = (-274) == (len(_dafny.Map({d_152_v2_: (d_151_v1_)[default__.safeIndex(d_201_v20_, len(d_151_v1_))]})))
            elif source3_.is_DC3:
                d_202_v21_: _dafny.Map
                d_202_v21_ = _dafny.Map({363: d_152_v2_})
                d_203_v22_: _dafny.Seq
                d_203_v22_ = _dafny.SeqWithoutIsStrInference([d_152_v2_])
                d_204_v23_: _dafny.Set
                d_204_v23_ = _dafny.Set({d_150_v0_, d_184_i1_, len(d_203_v22_)})
                d_205_v24_: int
                out4_: int
                out4_ = default__.m0((d_184_i1_) * (d_150_v0_), d_160_v10_, d_152_v2_, (d_202_v21_) | (_dafny.Map({default__.fm3(d_163_v11_, d_152_v2_, len(d_204_v23_), d_161_globalState_): d_152_v2_})), d_161_globalState_)
                d_205_v24_ = out4_
                d_152_v2_ = (d_203_v22_)[default__.safeIndex((0) - ((0) - (d_205_v24_)), len(d_203_v22_))]
                (d_161_globalState_).f0 = d_152_v2_
                (d_161_globalState_).f3 = default__.fm2(d_161_globalState_)
            elif True:
                d_206___mcc_h14_ = source3_.cf0
                d_207_cf0_ = d_206___mcc_h14_
                index34_ = default__.safeIndex(746, (d_160_v10_).length(0))
                (d_160_v10_)[index34_] = d_150_v0_
                index35_ = default__.safeIndex(746, (d_160_v10_).length(0))
                (d_160_v10_)[index35_] = len(d_158_v8_)
                d_208_v25_: C3
                nw30_ = C3()
                nw30_.ctor__(d_152_v2_)
                d_208_v25_ = nw30_
                d_209_v26_: _dafny.MultiSet
                d_209_v26_ = _dafny.MultiSet([d_150_v0_])
                d_210_v27_: _dafny.Seq
                d_210_v27_ = _dafny.SeqWithoutIsStrInference([d_152_v2_])
                d_211_v28_: _dafny.Set
                d_211_v28_ = _dafny.Set({(d_210_v27_)[default__.safeIndex(d_150_v0_, len(d_210_v27_))], d_152_v2_, d_152_v2_, d_152_v2_})
                d_212_v29_: _dafny.Array
                nw31_ = _dafny.Array(None, 23)
                nw31_[int(0)] = (default__.fm17(d_151_v1_, d_152_v2_, d_161_globalState_)) - (d_209_v26_)
                nw31_[int(1)] = d_209_v26_
                nw31_[int(2)] = d_209_v26_
                nw31_[int(3)] = d_209_v26_
                nw31_[int(4)] = d_209_v26_
                nw31_[int(5)] = (d_209_v26_).intersection(d_209_v26_)
                nw31_[int(6)] = d_209_v26_
                nw31_[int(7)] = d_209_v26_
                nw31_[int(8)] = d_209_v26_
                nw31_[int(9)] = (d_209_v26_) - (d_209_v26_)
                nw31_[int(10)] = d_209_v26_
                nw31_[int(11)] = _dafny.MultiSet([d_150_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpfylyv")))])
                nw31_[int(12)] = d_209_v26_
                nw31_[int(13)] = d_209_v26_
                nw31_[int(14)] = (d_209_v26_) | (d_209_v26_)
                nw31_[int(15)] = _dafny.MultiSet(default__.fm13(d_152_v2_, d_152_v2_, d_152_v2_, len(d_211_v28_), d_161_globalState_))
                nw31_[int(16)] = (d_209_v26_).intersection(_dafny.MultiSet([d_207_cf0_]))
                nw31_[int(17)] = d_209_v26_
                nw31_[int(18)] = default__.fm17(_dafny.SeqWithoutIsStrInference([d_207_cf0_ for d_213_i2_ in range(default__.abs(424))]), d_152_v2_, d_161_globalState_)
                nw31_[int(19)] = d_209_v26_
                nw31_[int(20)] = d_209_v26_
                nw31_[int(21)] = d_209_v26_
                nw31_[int(22)] = d_209_v26_
                d_212_v29_ = nw31_
                index36_ = default__.safeIndex(283, (d_212_v29_).length(0))
                (d_212_v29_)[index36_] = (d_209_v26_).intersection(d_209_v26_)
                index37_ = default__.safeIndex(746, (d_160_v10_).length(0))
                index38_ = default__.safeIndex(283, (d_212_v29_).length(0))
                rhs26_ = d_152_v2_
                rhs27_ = d_184_i1_
                rhs28_ = d_209_v26_
                rhs29_ = default__.fm2(d_161_globalState_)
                lhs17_ = d_161_globalState_
                lhs18_ = d_160_v10_
                lhs19_ = default__.safeIndex(746, (d_160_v10_).length(0))
                lhs20_ = d_212_v29_
                lhs21_ = default__.safeIndex(283, (d_212_v29_).length(0))
                lhs22_ = d_161_globalState_
                lhs17_.f3 = rhs26_
                lhs18_[lhs19_] = rhs27_
                lhs20_[lhs21_] = rhs28_
                lhs22_.f0 = rhs29_
                index39_ = default__.safeIndex(746, (d_160_v10_).length(0))
                (d_160_v10_)[index39_] = d_207_cf0_
            d_214_v30_: _dafny.Map
            d_214_v30_ = _dafny.Map({not(not(d_152_v2_)): d_155_v5_})
            d_214_v30_ = (d_214_v30_).set(d_152_v2_, d_155_v5_)
            d_215_v31_: _dafny.Set
            d_215_v31_ = _dafny.Set({d_152_v2_, d_152_v2_, False})
            d_215_v31_ = _dafny.Set({d_152_v2_, d_152_v2_})
            d_216_v32_: C3
            nw32_ = C3()
            nw32_.ctor__((d_184_i1_) >= (d_150_v0_))
            d_216_v32_ = nw32_
        (d_161_globalState_).f0 = d_152_v2_
        d_217_v33_: _dafny.Set
        d_217_v33_ = _dafny.Set({True, False, d_152_v2_})
        hi6_ = d_150_v0_
        for d_218_i3_ in range(len(d_217_v33_), hi6_):
            index40_ = default__.safeIndex(23, (d_155_v5_).length(0))
            (d_155_v5_)[index40_] = default__.fm2(d_161_globalState_)
            index41_ = default__.safeIndex(23, (d_155_v5_).length(0))
            (d_155_v5_)[index41_] = d_152_v2_
            index42_ = default__.safeIndex(13, (d_160_v10_).length(0))
            (d_160_v10_)[index42_] = (0) - (d_150_v0_)
            d_219_v34_: _dafny.Set
            d_219_v34_ = _dafny.Set({(0) - ((0) - (d_218_i3_)), d_150_v0_, d_218_i3_, (0) - (d_218_i3_)})
            index43_ = default__.safeIndex(13, (d_160_v10_).length(0))
            (d_160_v10_)[index43_] = default__.fm3(d_163_v11_, d_152_v2_, len(d_219_v34_), d_161_globalState_)
            index44_ = default__.safeIndex(13, (d_160_v10_).length(0))
            (d_160_v10_)[index44_] = d_218_i3_
            if (d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))]:
                (d_161_globalState_).f3 = ((d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))] if d_152_v2_ else d_152_v2_)
                d_150_v0_ = d_150_v0_
                d_220_v35_: C0
                nw33_ = C0()
                nw33_.ctor__((d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))], (d_160_v10_)[default__.safeIndex(13, (d_160_v10_).length(0))], d_152_v2_)
                d_220_v35_ = nw33_
                d_221_v36_: D3
                d_221_v36_ = D3_DC11(d_150_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")), len(_dafny.Map({d_152_v2_: d_218_i3_})), d_220_v35_)
                d_222_v37_: C1
                nw34_ = C1()
                nw34_.ctor__(False, (d_221_v36_).cf14, (d_220_v35_).f13)
                d_222_v37_ = nw34_
                d_223_v38_: _dafny.Map
                d_223_v38_ = _dafny.Map({d_222_v37_: len(d_158_v8_)})
                index45_ = default__.safeIndex(13, (d_160_v10_).length(0))
                (d_160_v10_)[index45_] = (((d_160_v10_)[default__.safeIndex(13, (d_160_v10_).length(0))] if (d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))] else len(d_223_v38_)) if d_152_v2_ else d_220_v35_.f14)
                (d_220_v35_).f14 = d_220_v35_.f14
                d_224_v39_: T0
                nw35_ = C1()
                nw35_.ctor__(d_152_v2_, (d_160_v10_)[default__.safeIndex(13, (d_160_v10_).length(0))], not(not(d_152_v2_)))
                d_224_v39_ = nw35_
                d_225_v40_: D4
                d_225_v40_ = D4_DC15(d_220_v35_.f14, d_152_v2_)
                nw36_ = C2()
                nw36_.ctor__((d_225_v40_).cf24, (d_224_v39_).f9)
                d_224_v39_ = nw36_
            elif True:
                d_150_v0_ = d_150_v0_
                d_226_v41_: _dafny.Array
                def lambda6_(d_227_v0_):
                    def lambda7_(d_228_i4_):
                        return default__.safeModuloInt(d_228_i4_, d_227_v0_)

                    return lambda7_

                init2_ = lambda6_(d_150_v0_)
                nw37_ = _dafny.Array(None, 13)
                for i0_2_ in range(nw37_.length(0)):
                    nw37_[i0_2_] = init2_(i0_2_)
                d_226_v41_ = nw37_
                d_229_v42_: D0
                d_229_v42_ = D0_DC1((d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))])
                d_230_v43_: _dafny.Map
                d_230_v43_ = _dafny.Map({d_150_v0_: (d_229_v42_).cf1})
                d_231_v44_: int
                out5_: int
                out5_ = default__.m0((d_160_v10_)[default__.safeIndex(13, (d_160_v10_).length(0))], d_226_v41_, not((d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))]), d_230_v43_, d_161_globalState_)
                d_231_v44_ = out5_
                d_232_v45_: _dafny.Map
                d_232_v45_ = _dafny.Map({default__.fm10(d_161_globalState_): d_157_v7_})
                d_233_v46_: _dafny.Seq
                d_233_v46_ = _dafny.SeqWithoutIsStrInference([(d_155_v5_)[default__.safeIndex(23, (d_155_v5_).length(0))], d_152_v2_])
                d_232_v45_ = ((d_232_v45_).set(d_233_v46_, D0_DC0(d_218_i3_))).set(default__.fm10(d_161_globalState_), d_157_v7_)
                d_150_v0_ = -74
                (d_161_globalState_).f7 = d_163_v11_
        d_234_v47_: _dafny.Seq
        d_234_v47_ = _dafny.SeqWithoutIsStrInference([d_152_v2_])
        (d_161_globalState_).f3 = (d_234_v47_) <= (d_234_v47_)
        d_235_v48_: C0
        nw38_ = C0()
        nw38_.ctor__(d_152_v2_, d_150_v0_, d_152_v2_)
        d_235_v48_ = nw38_
        if (d_152_v2_ if d_152_v2_ else (len(_dafny.Map({d_155_v5_: _dafny.Set({d_235_v48_, d_235_v48_})}))) < (d_150_v0_)):
            index46_ = default__.safeIndex(648, (d_160_v10_).length(0))
            (d_160_v10_)[index46_] = d_235_v48_.f14
            index47_ = default__.safeIndex(648, (d_160_v10_).length(0))
            (d_160_v10_)[index47_] = default__.safeModuloInt((0) - (d_235_v48_.f14), (d_235_v48_.f14) + (-739))
            index48_ = default__.safeIndex(89, (d_155_v5_).length(0))
            (d_155_v5_)[index48_] = (d_235_v48_).f13
            index49_ = default__.safeIndex(89, (d_155_v5_).length(0))
            (d_155_v5_)[index49_] = True
            d_236_v49_: T0
            nw39_ = C3()
            nw39_.ctor__(((d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))]) < ((d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))]))
            d_236_v49_ = nw39_
            d_236_v49_ = d_236_v49_
            if (26) != (default__.safeModuloInt(d_235_v48_.f14, d_235_v48_.f14)):
                d_237_v50_: _dafny.Array
                nw40_ = _dafny.Array(int(0), 13)
                d_237_v50_ = nw40_
                d_238_v51_: _dafny.Seq
                d_238_v51_ = _dafny.SeqWithoutIsStrInference([d_237_v50_, d_237_v50_, d_237_v50_])
                d_239_v52_: _dafny.Seq
                d_239_v52_ = _dafny.SeqWithoutIsStrInference([(d_238_v51_)[default__.safeIndex(len(d_158_v8_), len(d_238_v51_))], d_237_v50_])
                d_240_v53_: _dafny.Array
                nw41_ = _dafny.Array(None, 8)
                nw41_[int(0)] = d_237_v50_
                nw41_[int(1)] = d_237_v50_
                nw41_[int(2)] = (d_237_v50_ if d_152_v2_ else d_237_v50_)
                nw41_[int(3)] = (d_239_v52_)[default__.safeIndex((d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))], len(d_239_v52_))]
                nw41_[int(4)] = d_237_v50_
                nw41_[int(5)] = d_237_v50_
                nw41_[int(6)] = d_237_v50_
                nw41_[int(7)] = d_237_v50_
                d_240_v53_ = nw41_
                index50_ = default__.safeIndex(567, (d_240_v53_).length(0))
                (d_240_v53_)[index50_] = d_237_v50_
                index51_ = default__.safeIndex(567, (d_240_v53_).length(0))
                nw42_ = _dafny.Array(int(0), 22)
                (d_240_v53_)[index51_] = nw42_
                index52_ = default__.safeIndex(648, (d_160_v10_).length(0))
                (d_160_v10_)[index52_] = d_235_v48_.f14
                d_241_v54_: _dafny.Map
                d_241_v54_ = _dafny.Map({(d_155_v5_)[default__.safeIndex(89, (d_155_v5_).length(0))]: d_237_v50_})
                d_241_v54_ = (d_241_v54_ if not (default__.fm2(d_161_globalState_)) or (d_152_v2_) else (d_241_v54_).set(True, d_237_v50_))
                d_242_v55_: _dafny.MultiSet
                d_242_v55_ = _dafny.MultiSet([default__.safeDivisionInt((d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))], d_150_v0_)])
                d_243_v56_: _dafny.Map
                d_243_v56_ = _dafny.Map({d_235_v48_.f14: len(d_159_v9_)})
                d_244_v57_: _dafny.Seq
                d_244_v57_ = _dafny.SeqWithoutIsStrInference([d_234_v47_])
                index53_ = default__.safeIndex(648, (d_160_v10_).length(0))
                (d_160_v10_)[index53_] = (0) - ((0) - (((d_242_v55_)[((d_243_v56_)[(d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))]] if ((d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))]) in (d_243_v56_) else d_235_v48_.f14)] if (((d_243_v56_)[(d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))]] if ((d_160_v10_)[default__.safeIndex(648, (d_160_v10_).length(0))]) in (d_243_v56_) else d_235_v48_.f14)) in (d_242_v55_) else len((d_234_v47_) + ((d_244_v57_)[default__.safeIndex((0) - (d_150_v0_), len(d_244_v57_))])))))
                d_245_v58_: _dafny.Map
                d_245_v58_ = _dafny.Map({(d_235_v48_).f13: (d_235_v48_).f13})
                d_246_v59_: C2
                nw43_ = C2()
                nw43_.ctor__((d_245_v58_) != (d_245_v58_), (d_235_v48_).f13)
                d_246_v59_ = nw43_
                d_247_v60_: _dafny.Map
                d_247_v60_ = _dafny.Map({default__.fm2(d_161_globalState_): d_246_v59_})
                d_246_v59_ = (d_246_v59_ if d_152_v2_ else ((d_247_v60_)[(d_235_v48_).f13] if ((d_235_v48_).f13) in (d_247_v60_) else d_246_v59_))
            elif True:
                d_248_v61_: _dafny.Map
                d_248_v61_ = _dafny.Map({False: (d_235_v48_).f13})
                d_249_v62_: _dafny.Map
                d_249_v62_ = _dafny.Map({len(d_248_v61_): (d_235_v48_).f13})
                d_250_v63_: _dafny.Map
                d_250_v63_ = _dafny.Map({d_159_v9_: (d_249_v62_) != (d_249_v62_)})
                d_250_v63_ = d_250_v63_
                d_235_v48_ = d_235_v48_
                d_251_v64_: _dafny.Array
                nw44_ = _dafny.Array(int(0), 11)
                d_251_v64_ = nw44_
                d_252_v65_: C1
                nw45_ = C1()
                nw45_.ctor__(True, (0) - (d_235_v48_.f14), (d_155_v5_)[default__.safeIndex(89, (d_155_v5_).length(0))])
                d_252_v65_ = nw45_
                d_253_v66_: _dafny.Map
                d_253_v66_ = _dafny.Map({d_251_v64_: d_252_v65_})
                d_253_v66_ = (d_253_v66_).set(d_251_v64_, d_252_v65_)
                (d_235_v48_).f14 = default__.safeModuloInt(778, d_235_v48_.f14)
                d_254_v67_: _dafny.MultiSet
                d_254_v67_ = _dafny.MultiSet([(d_252_v65_).f12])
                d_255_v68_: C0
                nw46_ = C0()
                nw46_.ctor__((550) in (d_254_v67_), 114, (d_236_v49_).f9)
                d_255_v68_ = nw46_
            d_256_v69_: C2
            nw47_ = C2()
            nw47_.ctor__((d_163_v11_) in (d_158_v8_), (d_155_v5_)[default__.safeIndex(89, (d_155_v5_).length(0))])
            d_256_v69_ = nw47_
            d_256_v69_ = d_256_v69_
        elif True:
            index54_ = default__.safeIndex(109, (d_155_v5_).length(0))
            (d_155_v5_)[index54_] = (d_235_v48_).f13
            index55_ = default__.safeIndex(109, (d_155_v5_).length(0))
            (d_155_v5_)[index55_] = (d_235_v48_).f13
            (d_161_globalState_).f3 = (d_235_v48_).f13
            d_257_v70_: D0
            d_257_v70_ = D0_DC3()
            source4_ = (d_257_v70_ if (d_235_v48_).f13 else d_257_v70_)
            if source4_.is_DC1:
                d_258___mcc_h15_ = source4_.cf1
                d_259_cf1_ = d_258___mcc_h15_
                d_260_v71_: C1
                nw48_ = C1()
                nw48_.ctor__((d_235_v48_).f13, d_235_v48_.f14, not((d_235_v48_).f13))
                d_260_v71_ = nw48_
                d_261_v72_: _dafny.Map
                d_261_v72_ = _dafny.Map({186: d_260_v71_})
                d_262_v73_: _dafny.Map
                d_262_v73_ = _dafny.Map({len((d_261_v72_) | (_dafny.Map({366: d_260_v71_}))): d_234_v47_})
                rhs30_ = (d_262_v73_) | (d_262_v73_)
                rhs31_ = (0) - (d_235_v48_.f14)
                lhs23_ = d_235_v48_
                d_262_v73_ = rhs30_
                lhs23_.f14 = rhs31_
                d_158_v8_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))).set(default__.safeIndex((d_260_v71_).f12, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))), default__.fm18(d_161_globalState_))
                d_263_v74_: _dafny.Array
                def lambda8_(d_264_v8_):
                    def lambda9_(d_265_i5_):
                        return _dafny.SeqWithoutIsStrInference([D3_DC9(d_264_v8_) for d_266_i6_ in range(default__.abs(904))])

                    return lambda9_

                init3_ = lambda8_(d_158_v8_)
                nw49_ = _dafny.Array(None, 25)
                for i0_3_ in range(nw49_.length(0)):
                    nw49_[i0_3_] = init3_(i0_3_)
                d_263_v74_ = nw49_
                d_267_v75_: _dafny.Map
                d_267_v75_ = _dafny.Map({(d_260_v71_).f12: -277})
                d_268_v76_: D3
                d_268_v76_ = D3_DC9(d_158_v8_)
                d_269_v77_: _dafny.Seq
                d_269_v77_ = _dafny.SeqWithoutIsStrInference([d_268_v76_, d_268_v76_, d_268_v76_])
                index56_ = default__.safeIndex(111, (d_263_v74_).length(0))
                (d_263_v74_)[index56_] = (default__.fm19(d_267_v75_, d_161_globalState_)) + (d_269_v77_)
                index57_ = default__.safeIndex(111, (d_263_v74_).length(0))
                (d_263_v74_)[index57_] = d_269_v77_
                d_270_v79_: int
                out6_: int
                def iife31_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(515, 497):
                        d_271_v78_: int = compr_9_
                        if ((515) <= (d_271_v78_)) and ((d_271_v78_) < (497)):
                            coll9_[(d_271_v78_) - (d_235_v48_.f14)] = False
                    return _dafny.Map(coll9_)
                out6_ = default__.m0(d_235_v48_.f14, d_160_v10_, (d_158_v8_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gretekrl"))), iife31_()
                , d_161_globalState_)
                d_270_v79_ = out6_
            elif source4_.is_DC2:
                d_272___mcc_h16_ = source4_.cf2
                d_273___mcc_h17_ = source4_.cf3
                d_274___mcc_h18_ = source4_.cf4
                d_275_cf4_ = d_274___mcc_h18_
                d_276_cf3_ = d_273___mcc_h17_
                d_277_cf2_ = d_272___mcc_h16_
                d_278_v80_: C1
                nw50_ = C1()
                nw50_.ctor__((d_155_v5_)[default__.safeIndex(109, (d_155_v5_).length(0))], (d_276_cf3_) * (d_276_cf3_), (d_235_v48_).f13)
                d_278_v80_ = nw50_
                d_159_v9_ = (d_159_v9_).set((d_278_v80_).f11, len(d_159_v9_))
                d_279_v81_: C2
                nw51_ = C2()
                nw51_.ctor__((d_235_v48_).f13, (d_151_v1_) <= (d_151_v1_))
                d_279_v81_ = nw51_
                nw52_ = C2()
                nw52_.ctor__((d_150_v0_) == (d_235_v48_.f14), ((d_278_v80_).f12) >= (d_276_cf3_))
                d_279_v81_ = nw52_
                rhs32_ = default__.fm18(d_161_globalState_)
                rhs33_ = 712
                lhs24_ = d_161_globalState_
                lhs24_.f7 = rhs32_
                d_276_cf3_ = rhs33_
            elif source4_.is_DC3:
                (d_235_v48_).f14 = d_235_v48_.f14
                d_280_v82_: _dafny.Map
                d_280_v82_ = _dafny.Map({d_152_v2_: (d_158_v8_) != (d_158_v8_)})
                d_281_v83_: _dafny.Array
                nw53_ = _dafny.Array(False, 29)
                d_281_v83_ = nw53_
                d_282_v84_: _dafny.Map
                d_282_v84_ = _dafny.Map({D2_DC7(d_281_v83_): d_158_v8_})
                d_283_v85_: _dafny.Map
                d_283_v85_ = _dafny.Map({((d_282_v84_)[D2_DC7(d_281_v83_)] if (D2_DC7(d_281_v83_)) in (d_282_v84_) else d_158_v8_): d_234_v47_})
                rhs34_ = d_280_v82_
                rhs35_ = (d_283_v85_) | (d_283_v85_)
                rhs36_ = d_150_v0_
                lhs25_ = d_235_v48_
                d_280_v82_ = rhs34_
                d_283_v85_ = rhs35_
                lhs25_.f14 = rhs36_
                d_284_v86_: _dafny.Map
                d_284_v86_ = _dafny.Map({len(d_158_v8_): d_150_v0_})
                d_285_v87_: _dafny.Map
                d_285_v87_ = _dafny.Map({len((d_284_v86_).set(len(d_158_v8_), d_235_v48_.f14)): default__.fm2(d_161_globalState_)})
                d_286_v88_: int
                out7_: int
                out7_ = default__.m0((0) - (len(d_158_v8_)), d_160_v10_, d_152_v2_, d_285_v87_, d_161_globalState_)
                d_286_v88_ = out7_
                d_287_v89_: int
                out8_: int
                out8_ = default__.m0((0) - ((d_150_v0_) + (514)), d_160_v10_, (d_235_v48_).f13, d_285_v87_, d_161_globalState_)
                d_287_v89_ = out8_
            elif True:
                d_288___mcc_h19_ = source4_.cf0
                d_289_cf0_ = d_288___mcc_h19_
                (d_161_globalState_).f0 = (d_235_v48_).f13
                d_290_v90_: _dafny.Array
                def lambda10_(d_291_v11_, d_292_v8_):
                    def lambda11_(d_293_i7_):
                        return (_dafny.SeqWithoutIsStrInference([d_291_v11_ for d_294_i8_ in range(default__.abs(-815))])) + (d_292_v8_)

                    return lambda11_

                init4_ = lambda10_(d_163_v11_, d_158_v8_)
                nw54_ = _dafny.Array(None, 27)
                for i0_4_ in range(nw54_.length(0)):
                    nw54_[i0_4_] = init4_(i0_4_)
                d_290_v90_ = nw54_
                d_290_v90_ = d_290_v90_
                d_295_v91_: C0
                nw55_ = C0()
                nw55_.ctor__(not((d_235_v48_).f13), d_235_v48_.f14, d_152_v2_)
                d_295_v91_ = nw55_
                d_296_v92_: C1
                nw56_ = C1()
                nw56_.ctor__(True, d_235_v48_.f14, (d_235_v48_).f13)
                d_296_v92_ = nw56_
                d_297_v93_: _dafny.Seq
                d_297_v93_ = _dafny.SeqWithoutIsStrInference([d_296_v92_])
                d_152_v2_ = ((428 if (d_235_v48_).f13 else len(d_297_v93_))) < (d_150_v0_)
            d_298_v94_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.MultiSet({}), 15)
            d_298_v94_ = nw57_
            index58_ = default__.safeIndex(326, (d_298_v94_).length(0))
            (d_298_v94_)[index58_] = _dafny.MultiSet(((d_234_v47_).set(default__.safeIndex(d_150_v0_, len(d_234_v47_)), (d_235_v48_).f13)) + (d_234_v47_))
            d_299_v95_: _dafny.MultiSet
            d_299_v95_ = _dafny.MultiSet([True])
            index59_ = default__.safeIndex(326, (d_298_v94_).length(0))
            (d_298_v94_)[index59_] = d_299_v95_
            d_150_v0_ = d_235_v48_.f14
        d_300_v96_: _dafny.MultiSet
        d_300_v96_ = _dafny.MultiSet([d_150_v0_])
        d_150_v0_ = (len(_dafny.Map({(d_235_v48_).f13: (d_300_v96_).cardinality}))) + (771)
        if ((default__.fm2(d_161_globalState_)) == ((d_235_v48_).f13) if d_152_v2_ else d_152_v2_):
            (d_235_v48_).f14 = d_235_v48_.f14
            d_301_v97_: D0
            d_301_v97_ = D0_DC3()
            d_302_v98_: _dafny.Set
            d_302_v98_ = _dafny.Set({d_301_v97_, d_301_v97_})
            (d_161_globalState_).f3 = ((d_300_v96_).cardinality) == (len(d_302_v98_))
            d_303_v99_: C2
            nw58_ = C2()
            nw58_.ctor__((_dafny.Set({d_152_v2_})).ispropersubset(d_217_v33_), (d_235_v48_).f13)
            d_303_v99_ = nw58_
            d_303_v99_ = d_303_v99_
            d_304_v100_: C2
            nw59_ = C2()
            nw59_.ctor__(False, d_152_v2_)
            d_304_v100_ = nw59_
            d_235_v48_ = d_235_v48_
        elif True:
            (d_161_globalState_).f0 = (d_235_v48_).f13
            (d_161_globalState_).f7 = d_163_v11_
            d_151_v1_ = _dafny.SeqWithoutIsStrInference([(default__.fm3(d_163_v11_, (d_235_v48_).f13, default__.fm3(d_163_v11_, (d_235_v48_).f13, d_235_v48_.f14, d_161_globalState_), d_161_globalState_)) - (d_235_v48_.f14), 978, d_150_v0_])
            if (d_235_v48_).f13:
                d_305_v101_: C3
                nw60_ = C3()
                nw60_.ctor__(False)
                d_305_v101_ = nw60_
                d_306_v102_: _dafny.Seq
                d_306_v102_ = _dafny.SeqWithoutIsStrInference([d_160_v10_, d_160_v10_, d_160_v10_])
                (d_235_v48_).f14 = (len(d_306_v102_)) + ((d_150_v0_) + (default__.fm3(d_163_v11_, (d_235_v48_).f13, d_235_v48_.f14, d_161_globalState_)))
                d_307_v103_: _dafny.Map
                d_307_v103_ = _dafny.Map({d_235_v48_.f14: (d_235_v48_).f13})
                d_308_v104_: int
                out9_: int
                out9_ = default__.m0(d_235_v48_.f14, d_160_v10_, (d_235_v48_).f13, (_dafny.Map({d_235_v48_.f14: (d_235_v48_).f13})) | (d_307_v103_), d_161_globalState_)
                d_308_v104_ = out9_
                (d_235_v48_).f14 = d_235_v48_.f14
                d_309_v105_: C3
                nw61_ = C3()
                nw61_.ctor__((d_235_v48_).f13)
                d_309_v105_ = nw61_
            elif True:
                d_310_v106_: C3
                nw62_ = C3()
                nw62_.ctor__(d_152_v2_)
                d_310_v106_ = nw62_
                d_311_v107_: _dafny.Seq
                d_311_v107_ = _dafny.SeqWithoutIsStrInference([d_235_v48_])
                d_312_v108_: D3
                d_312_v108_ = D3_DC11(len((_dafny.SeqWithoutIsStrInference([d_310_v106_, d_310_v106_])).set(default__.safeIndex((0) - (d_150_v0_), len(_dafny.SeqWithoutIsStrInference([d_310_v106_, d_310_v106_]))), d_310_v106_)), _dafny.SeqWithoutIsStrInference([d_163_v11_ for d_313_i9_ in range(default__.abs(310))]), d_235_v48_.f14, (d_311_v107_)[default__.safeIndex(353, len(d_311_v107_))])
                d_312_v108_ = d_312_v108_
                d_310_v106_ = (D5_DC17(d_310_v106_)).cf26
                d_314_v109_: T0
                nw63_ = C3()
                nw63_.ctor__((d_235_v48_).f13)
                d_314_v109_ = nw63_
                index60_ = default__.safeIndex(403, (d_155_v5_).length(0))
                (d_155_v5_)[index60_] = (d_235_v48_).f13
                index61_ = default__.safeIndex(324, (d_155_v5_).length(0))
                (d_155_v5_)[index61_] = (d_235_v48_).f13
                index62_ = default__.safeIndex(403, (d_155_v5_).length(0))
                index63_ = default__.safeIndex(324, (d_155_v5_).length(0))
                rhs37_ = d_314_v109_
                rhs38_ = not (((d_235_v48_).f13 if d_152_v2_ else (d_235_v48_).f13)) or ((d_235_v48_).f13)
                rhs39_ = (d_235_v48_).f13
                rhs40_ = (d_235_v48_).f13
                lhs26_ = d_155_v5_
                lhs27_ = default__.safeIndex(403, (d_155_v5_).length(0))
                lhs28_ = d_155_v5_
                lhs29_ = default__.safeIndex(324, (d_155_v5_).length(0))
                d_314_v109_ = rhs37_
                lhs26_[lhs27_] = rhs38_
                d_152_v2_ = rhs39_
                lhs28_[lhs29_] = rhs40_
                (d_235_v48_).f14 = (len((d_151_v1_) + (d_151_v1_))) * (default__.safeDivisionInt((d_151_v1_)[default__.safeIndex(d_235_v48_.f14, len(d_151_v1_))], d_235_v48_.f14))
                d_315_v110_: D4
                d_315_v110_ = D4_DC15(718, True)
                d_316_v111_: _dafny.Map
                d_316_v111_ = _dafny.Map({(d_150_v0_ if (d_315_v110_).cf24 else len(_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_235_v48_).f13, (d_314_v109_).f9])}))): (d_235_v48_).f13})
                rhs41_ = d_160_v10_
                rhs42_ = d_316_v111_
                rhs43_ = len(d_217_v33_)
                rhs44_ = d_163_v11_
                lhs30_ = d_161_globalState_
                d_160_v10_ = rhs41_
                d_316_v111_ = rhs42_
                d_150_v0_ = rhs43_
                lhs30_.f7 = rhs44_
            d_155_v5_ = d_155_v5_
        d_317_v112_: C3
        nw64_ = C3()
        nw64_.ctor__(d_152_v2_)
        d_317_v112_ = nw64_
        d_318_v113_: _dafny.Set
        d_318_v113_ = _dafny.Set({(0) - (len(d_217_v33_))})
        d_319_v114_: D5
        d_319_v114_ = D5_DC18(len(d_318_v113_), (d_235_v48_).f13)
        d_150_v0_ = (d_150_v0_ if not((d_235_v48_.f14) >= ((d_319_v114_).cf27)) else d_150_v0_)
        if False:
            d_320_v115_: _dafny.Map
            d_320_v115_ = _dafny.Map({d_235_v48_.f14: d_235_v48_.f14})
            d_321_v116_: _dafny.Map
            d_321_v116_ = _dafny.Map({d_150_v0_: (d_235_v48_).f13})
            d_322_v117_: int
            out10_: int
            out10_ = default__.m0((_dafny.MultiSet([d_320_v115_, _dafny.Map({d_150_v0_: d_235_v48_.f14})])).cardinality, d_160_v10_, (d_318_v113_) != (d_318_v113_), d_321_v116_, d_161_globalState_)
            d_322_v117_ = out10_
            d_323_v118_: D0
            d_323_v118_ = D0_DC3()
            d_323_v118_ = d_323_v118_
            d_324_v119_: _dafny.Seq
            d_324_v119_ = _dafny.SeqWithoutIsStrInference([d_160_v10_, d_160_v10_])
            d_325_v120_: _dafny.Array
            nw65_ = _dafny.Array(None, 22)
            nw65_[int(0)] = d_160_v10_
            nw65_[int(1)] = d_160_v10_
            nw65_[int(2)] = d_160_v10_
            nw65_[int(3)] = d_160_v10_
            nw65_[int(4)] = d_160_v10_
            nw65_[int(5)] = d_160_v10_
            nw65_[int(6)] = d_160_v10_
            nw65_[int(7)] = d_160_v10_
            nw65_[int(8)] = (d_324_v119_)[default__.safeIndex(d_235_v48_.f14, len(d_324_v119_))]
            nw65_[int(9)] = d_160_v10_
            nw65_[int(10)] = d_160_v10_
            nw65_[int(11)] = d_160_v10_
            nw65_[int(12)] = d_160_v10_
            nw65_[int(13)] = d_160_v10_
            nw65_[int(14)] = d_160_v10_
            nw65_[int(15)] = d_160_v10_
            nw65_[int(16)] = d_160_v10_
            nw65_[int(17)] = d_160_v10_
            nw65_[int(18)] = d_160_v10_
            nw65_[int(19)] = d_160_v10_
            nw65_[int(20)] = d_160_v10_
            nw65_[int(21)] = (d_324_v119_)[default__.safeIndex(d_150_v0_, len(d_324_v119_))]
            d_325_v120_ = nw65_
            index64_ = default__.safeIndex(488, (d_325_v120_).length(0))
            (d_325_v120_)[index64_] = (d_324_v119_)[default__.safeIndex(d_322_v117_, len(d_324_v119_))]
            index65_ = default__.safeIndex(488, (d_325_v120_).length(0))
            (d_325_v120_)[index65_] = d_160_v10_
            d_326_v121_: _dafny.MultiSet
            d_326_v121_ = _dafny.MultiSet([d_235_v48_])
            d_327_v122_: C2
            nw66_ = C2()
            nw66_.ctor__(False, (d_235_v48_).f13)
            d_327_v122_ = nw66_
            d_328_v123_: _dafny.Seq
            d_328_v123_ = _dafny.SeqWithoutIsStrInference([d_327_v122_])
            d_329_v124_: _dafny.Map
            d_329_v124_ = _dafny.Map({d_326_v121_: (d_328_v123_)[default__.safeIndex(d_322_v117_, len(d_328_v123_))]})
            d_329_v124_ = _dafny.Map({_dafny.MultiSet([d_235_v48_]): d_327_v122_})
            d_330_v125_: T0
            nw67_ = C3()
            nw67_.ctor__((d_235_v48_).f13)
            d_330_v125_ = nw67_
            d_331_v126_: C0
            nw68_ = C0()
            nw68_.ctor__((len(_dafny.Map({d_330_v125_: (d_330_v125_).f9}))) < (len((d_234_v47_).set(default__.safeIndex(len(_dafny.Map({d_235_v48_.f14: d_322_v117_})), len(d_234_v47_)), (d_235_v48_).f13))), d_235_v48_.f14, True)
            d_331_v126_ = nw68_
        elif True:
            d_332_v127_: _dafny.MultiSet
            d_332_v127_ = _dafny.MultiSet([not((d_235_v48_).f13), (d_235_v48_).f13, not(True)])
            (d_161_globalState_).f3 = not(not ((True) == ((d_235_v48_).f13)) or ((d_332_v127_).isdisjoint(default__.fm12(not((d_235_v48_).f13), (d_235_v48_).f13, d_235_v48_.f14, (d_235_v48_).f13, d_161_globalState_))))
            d_333_v128_: _dafny.Map
            d_333_v128_ = _dafny.Map({(d_235_v48_).f13: True})
            d_333_v128_ = (d_333_v128_).set((d_235_v48_).f13, (d_235_v48_).f13)
            d_334_v129_: D3
            d_334_v129_ = D3_DC10(d_235_v48_.f14, d_160_v10_, False)
            d_160_v10_ = (d_334_v129_).cf12
            d_335_v130_: C3
            nw69_ = C3()
            nw69_.ctor__((d_235_v48_.f14) != (d_235_v48_.f14))
            d_335_v130_ = nw69_
            d_336_v131_: _dafny.Map
            d_336_v131_ = _dafny.Map({d_235_v48_.f14: (d_158_v8_) == ((d_158_v8_).set(default__.safeIndex(d_235_v48_.f14, len(d_158_v8_)), d_163_v11_))})
            d_336_v131_ = (d_336_v131_).set((d_235_v48_.f14) - (d_235_v48_.f14), d_152_v2_)
        hi7_ = len(d_217_v33_)
        for d_337_i10_ in range(d_150_v0_, hi7_):
            d_338_v132_: D4
            d_338_v132_ = D4_DC15(631, d_152_v2_)
            d_339_v133_: _dafny.Map
            d_339_v133_ = _dafny.Map({d_338_v132_: d_155_v5_})
            d_340_v134_: D2
            d_340_v134_ = D2_DC7(((d_339_v133_)[d_338_v132_] if (d_338_v132_) in (d_339_v133_) else d_155_v5_))
            source5_ = d_340_v134_
            if source5_.is_DC8:
                d_341___mcc_h20_ = source5_.cf7
                d_342___mcc_h21_ = source5_.cf8
                d_343___mcc_h22_ = source5_.cf9
                d_344_cf9_ = d_343___mcc_h22_
                d_345_cf8_ = d_342___mcc_h21_
                d_346_cf7_ = d_341___mcc_h20_
                d_347_v135_: _dafny.Array
                nw70_ = _dafny.Array(None, 5)
                nw70_[int(0)] = d_163_v11_
                nw70_[int(1)] = d_163_v11_
                nw70_[int(2)] = d_163_v11_
                nw70_[int(3)] = d_163_v11_
                nw70_[int(4)] = _dafny.CodePoint('n')
                d_347_v135_ = nw70_
                d_348_v136_: _dafny.Set
                d_348_v136_ = _dafny.Set({d_347_v135_})
                d_349_v137_: int
                d_350_v138_: D0
                d_351_v139_: int
                out11_: int
                out12_: D0
                out13_: int
                out11_, out12_, out13_ = (d_317_v112_).m1(d_348_v136_, d_161_globalState_)
                d_349_v137_ = out11_
                d_350_v138_ = out12_
                d_351_v139_ = out13_
                index66_ = default__.safeIndex(960, (d_160_v10_).length(0))
                (d_160_v10_)[index66_] = d_150_v0_
                index67_ = default__.safeIndex(960, (d_160_v10_).length(0))
                (d_160_v10_)[index67_] = (353) * (((d_154_v4_)[d_151_v1_] if (d_151_v1_) in (d_154_v4_) else d_337_i10_))
                d_152_v2_ = not (not ((d_235_v48_).f13) or ((d_319_v114_).cf28)) or ((d_235_v48_).f13)
                (d_235_v48_).f14 = d_235_v48_.f14
            elif True:
                d_352___mcc_h23_ = source5_.cf6
                d_353_cf6_ = d_352___mcc_h23_
                (d_161_globalState_).f0 = d_152_v2_
                (d_161_globalState_).f0 = d_152_v2_
                d_354_v140_: _dafny.Seq
                d_354_v140_ = _dafny.SeqWithoutIsStrInference([d_160_v10_, d_160_v10_, d_160_v10_, d_160_v10_])
                d_355_v141_: _dafny.Map
                d_355_v141_ = _dafny.Map({d_163_v11_: d_158_v8_})
                d_356_v142_: int
                out14_: int
                out14_ = default__.m0((209) - (d_235_v48_.f14), (d_354_v140_)[default__.safeIndex(d_235_v48_.f14, len(d_354_v140_))], not((d_235_v48_).f13), default__.fm16(d_355_v141_, d_337_i10_, d_161_globalState_), d_161_globalState_)
                d_356_v142_ = out14_
                (d_235_v48_).f14 = (0) - (default__.safeModuloInt((d_235_v48_.f14) - (d_150_v0_), default__.safeModuloInt(684, (0) - (d_235_v48_.f14))))
            d_357_v143_: _dafny.Map
            d_357_v143_ = _dafny.Map({d_337_i10_: d_235_v48_.f14})
            d_357_v143_ = d_357_v143_
            d_358_v144_: _dafny.Array
            def lambda12_(d_359_v11_):
                def lambda13_(d_360_i11_):
                    return d_359_v11_

                return lambda13_

            init5_ = lambda12_(d_163_v11_)
            nw71_ = _dafny.Array(None, 25)
            for i0_5_ in range(nw71_.length(0)):
                nw71_[i0_5_] = init5_(i0_5_)
            d_358_v144_ = nw71_
            d_361_v145_: _dafny.Set
            d_361_v145_ = _dafny.Set({d_358_v144_})
            d_362_v146_: int
            d_363_v147_: D0
            d_364_v148_: int
            out15_: int
            out16_: D0
            out17_: int
            out15_, out16_, out17_ = (d_317_v112_).m1(d_361_v145_, d_161_globalState_)
            d_362_v146_ = out15_
            d_363_v147_ = out16_
            d_364_v148_ = out17_
            index68_ = default__.safeIndex(984, (d_155_v5_).length(0))
            (d_155_v5_)[index68_] = (d_235_v48_).f13
            d_365_v149_: _dafny.Seq
            d_365_v149_ = _dafny.SeqWithoutIsStrInference([d_235_v48_])
            index69_ = default__.safeIndex(984, (d_155_v5_).length(0))
            (d_155_v5_)[index69_] = ((d_365_v149_) + (d_365_v149_)) < ((d_365_v149_) + (d_365_v149_))
        d_366_v150_: D5
        d_366_v150_ = D5_DC19(d_217_v33_)
        d_367_v151_: _dafny.Map
        d_367_v151_ = _dafny.Map({d_366_v150_: d_155_v5_})
        d_368_v152_: _dafny.Array
        nw72_ = _dafny.Array(None, 27)
        nw72_[int(0)] = d_367_v151_
        nw72_[int(1)] = d_367_v151_
        nw72_[int(2)] = d_367_v151_
        nw72_[int(3)] = (d_367_v151_) | (d_367_v151_)
        nw72_[int(4)] = d_367_v151_
        nw72_[int(5)] = (d_367_v151_).set(d_366_v150_, d_155_v5_)
        nw72_[int(6)] = d_367_v151_
        nw72_[int(7)] = (d_367_v151_) | (_dafny.Map({d_366_v150_: d_155_v5_}))
        nw72_[int(8)] = d_367_v151_
        nw72_[int(9)] = (_dafny.Map({d_366_v150_: d_155_v5_})) | (d_367_v151_)
        nw72_[int(10)] = d_367_v151_
        nw72_[int(11)] = d_367_v151_
        nw72_[int(12)] = d_367_v151_
        nw72_[int(13)] = (d_367_v151_) | (d_367_v151_)
        nw72_[int(14)] = d_367_v151_
        nw72_[int(15)] = _dafny.Map({d_366_v150_: d_155_v5_})
        nw72_[int(16)] = (d_367_v151_) | (d_367_v151_)
        nw72_[int(17)] = d_367_v151_
        nw72_[int(18)] = (_dafny.Map({d_366_v150_: d_155_v5_})) | (d_367_v151_)
        nw72_[int(19)] = d_367_v151_
        nw72_[int(20)] = d_367_v151_
        nw72_[int(21)] = ((d_367_v151_).set(d_366_v150_, d_155_v5_)) | (d_367_v151_)
        nw72_[int(22)] = d_367_v151_
        nw72_[int(23)] = (d_367_v151_) | (d_367_v151_)
        nw72_[int(24)] = _dafny.Map({d_366_v150_: d_155_v5_})
        nw72_[int(25)] = d_367_v151_
        nw72_[int(26)] = d_367_v151_
        d_368_v152_ = nw72_
        index70_ = default__.safeIndex(310, (d_368_v152_).length(0))
        (d_368_v152_)[index70_] = d_367_v151_
        index71_ = default__.safeIndex(310, (d_368_v152_).length(0))
        (d_368_v152_)[index71_] = d_367_v151_
        d_369_v153_: _dafny.MultiSet
        d_369_v153_ = _dafny.MultiSet([d_152_v2_])
        d_370_v154_: _dafny.Map
        d_370_v154_ = _dafny.Map({not((D0_DC1(False)).cf1): ((d_235_v48_).f13) == ((d_235_v48_).f13)})
        rhs45_ = (d_217_v33_).issubset(d_217_v33_)
        rhs46_ = ((d_370_v154_)[default__.fm2(d_161_globalState_)] if (default__.fm2(d_161_globalState_)) in (d_370_v154_) else not((d_235_v48_).f13))
        rhs47_ = (_dafny.MultiSet([d_152_v2_, (d_235_v48_).f13])) - (d_369_v153_)
        rhs48_ = d_155_v5_
        rhs49_ = d_235_v48_.f14
        lhs31_ = d_161_globalState_
        lhs32_ = d_235_v48_
        d_152_v2_ = rhs45_
        lhs31_.f3 = rhs46_
        d_369_v153_ = rhs47_
        d_155_v5_ = rhs48_
        lhs32_.f14 = rhs49_
        if d_152_v2_:
            d_371_v155_: _dafny.MultiSet
            d_371_v155_ = _dafny.MultiSet([d_163_v11_])
            d_372_v156_: C2
            nw73_ = C2()
            nw73_.ctor__((d_235_v48_.f14) > (((d_371_v155_)[d_163_v11_] if (d_163_v11_) in (d_371_v155_) else d_150_v0_)), (d_235_v48_).f13)
            d_372_v156_ = nw73_
            index72_ = default__.safeIndex(152, (d_160_v10_).length(0))
            (d_160_v10_)[index72_] = d_235_v48_.f14
            index73_ = default__.safeIndex(152, (d_160_v10_).length(0))
            (d_160_v10_)[index73_] = len((d_318_v113_) - (d_318_v113_))
            d_373_v157_: T0
            nw74_ = C1()
            nw74_.ctor__((d_235_v48_).f13, d_235_v48_.f14, (d_235_v48_).f13)
            d_373_v157_ = nw74_
            d_374_v158_: _dafny.MultiSet
            d_374_v158_ = _dafny.MultiSet([d_373_v157_, d_373_v157_])
            d_375_v159_: C0
            nw75_ = C0()
            nw75_.ctor__((d_235_v48_).f13, (0) - ((d_374_v158_).cardinality), (len(_dafny.Map({(d_160_v10_)[default__.safeIndex(152, (d_160_v10_).length(0))]: d_155_v5_}))) != (d_235_v48_.f14))
            d_375_v159_ = nw75_
            d_376_v160_: _dafny.Array
            nw76_ = _dafny.Array(_dafny.Map({}), 24)
            d_376_v160_ = nw76_
            d_377_v161_: _dafny.Map
            d_377_v161_ = _dafny.Map({d_376_v160_: (d_373_v157_).f9})
            d_378_v162_: C1
            nw77_ = C1()
            nw77_.ctor__(((d_377_v161_)[d_376_v160_] if (d_376_v160_) in (d_377_v161_) else (d_375_v159_).f13), len(default__.fm11(default__.fm2(d_161_globalState_), (d_235_v48_).f13, d_161_globalState_)), default__.fm2(d_161_globalState_))
            d_378_v162_ = nw77_
            d_379_v163_: _dafny.Map
            d_379_v163_ = _dafny.Map({len(d_158_v8_): d_235_v48_.f14})
            d_380_v164_: _dafny.Seq
            d_380_v164_ = _dafny.SeqWithoutIsStrInference([d_378_v162_, d_378_v162_, d_378_v162_])
            rhs50_ = (d_380_v164_)[default__.safeIndex((d_160_v10_)[default__.safeIndex(152, (d_160_v10_).length(0))], len(d_380_v164_))]
            rhs51_ = default__.fm2(d_161_globalState_)
            rhs52_ = _dafny.Map({d_150_v0_: len(d_158_v8_)})
            lhs33_ = d_372_v156_
            d_378_v162_ = rhs50_
            lhs33_.f10 = rhs51_
            d_379_v163_ = rhs52_
            d_381_v165_: _dafny.Map
            d_381_v165_ = _dafny.Map({_dafny.Map({d_235_v48_.f14: d_235_v48_.f14}): d_375_v159_.f14})
            (d_235_v48_).f14 = ((d_381_v165_)[d_379_v163_] if (d_379_v163_) in (d_381_v165_) else d_235_v48_.f14)
        elif True:
            d_159_v9_ = (d_159_v9_).set(d_152_v2_, d_235_v48_.f14)
            d_382_v166_: _dafny.Set
            d_382_v166_ = _dafny.Set({d_155_v5_, d_155_v5_})
            d_382_v166_ = d_382_v166_
            (d_161_globalState_).f3 = not (d_152_v2_) or ((d_235_v48_).f13)
            (d_235_v48_).f14 = len(_dafny.SeqWithoutIsStrInference([d_163_v11_ for d_383_i12_ in range(default__.abs(522))]))
            d_384_v167_: C1
            nw78_ = C1()
            nw78_.ctor__((d_235_v48_).f13, d_235_v48_.f14, d_152_v2_)
            d_384_v167_ = nw78_
            d_385_v168_: _dafny.Array
            nw79_ = _dafny.Array(None, 3)
            nw79_[int(0)] = d_384_v167_
            nw79_[int(1)] = d_384_v167_
            nw79_[int(2)] = d_384_v167_
            d_385_v168_ = nw79_
            d_386_v169_: _dafny.MultiSet
            d_386_v169_ = _dafny.MultiSet([_dafny.MultiSet([(d_235_v48_).f13, (d_235_v48_).f13, (d_235_v48_).f13, (d_235_v48_).f13])])
            index74_ = default__.safeIndex(820, (d_155_v5_).length(0))
            (d_155_v5_)[index74_] = d_152_v2_
            d_387_v170_: T0
            nw80_ = C3()
            nw80_.ctor__((d_235_v48_).f13)
            d_387_v170_ = nw80_
            index75_ = default__.safeIndex(820, (d_155_v5_).length(0))
            rhs53_ = d_385_v168_
            rhs54_ = (d_386_v169_).intersection((d_386_v169_) | (_dafny.MultiSet([d_369_v153_])))
            rhs55_ = ((d_235_v48_).f13 if (d_235_v48_).f13 else (d_384_v167_).fm8((d_387_v170_).f9, d_161_globalState_))
            rhs56_ = d_387_v170_
            lhs34_ = d_155_v5_
            lhs35_ = default__.safeIndex(820, (d_155_v5_).length(0))
            d_385_v168_ = rhs53_
            d_386_v169_ = rhs54_
            lhs34_[lhs35_] = rhs55_
            d_387_v170_ = rhs56_
        hi8_ = d_235_v48_.f14
        for d_388_i13_ in range((0) - (d_235_v48_.f14), hi8_):
            d_152_v2_ = not ((d_235_v48_).f13) or ((d_235_v48_).f13)
            index76_ = default__.safeIndex(222, (d_155_v5_).length(0))
            (d_155_v5_)[index76_] = (d_235_v48_).f13
            index77_ = default__.safeIndex(222, (d_155_v5_).length(0))
            (d_155_v5_)[index77_] = (False if (d_235_v48_).f13 else (d_235_v48_).f13)
            d_389_v171_: _dafny.Map
            d_389_v171_ = _dafny.Map({True: d_234_v47_})
            d_390_v172_: _dafny.Map
            d_390_v172_ = _dafny.Map({d_388_i13_: (d_155_v5_)[default__.safeIndex(222, (d_155_v5_).length(0))]})
            nw81_ = _dafny.Array(None, 9)
            nw81_[int(0)] = -787
            nw81_[int(1)] = (0) - ((0) - (d_235_v48_.f14))
            nw81_[int(2)] = (d_235_v48_).fm4(d_389_v171_, d_152_v2_, (d_235_v48_).f13, (d_317_v112_).fm4(d_389_v171_, (d_235_v48_).f13, False, d_235_v48_.f14, d_161_globalState_), d_161_globalState_)
            nw81_[int(3)] = d_235_v48_.f14
            nw81_[int(4)] = (d_235_v48_.f14) - ((d_369_v153_).cardinality)
            nw81_[int(5)] = (0) - ((0) - (d_235_v48_.f14))
            nw81_[int(6)] = len(_dafny.Map({((d_370_v154_)[(d_155_v5_)[default__.safeIndex(222, (d_155_v5_).length(0))]] if ((d_155_v5_)[default__.safeIndex(222, (d_155_v5_).length(0))]) in (d_370_v154_) else default__.fm2(d_161_globalState_)): d_390_v172_}))
            nw81_[int(7)] = d_388_i13_
            nw81_[int(8)] = d_235_v48_.f14
            d_160_v10_ = nw81_
            (d_161_globalState_).f7 = default__.fm18(d_161_globalState_)
        _dafny.print(_dafny.string_of(d_150_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v1_) == (_dafny.SeqWithoutIsStrInference([636, 636]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_152_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v3_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([636, 636])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v4_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v6_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v7_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_158_v8_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v9_) == (_dafny.Map({True: 636}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f5) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636]), _dafny.SeqWithoutIsStrInference([636, 636])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f6)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_globalState_).f8) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_163_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v33_) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v47_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_235_v48_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_235_v48_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v96_) == (_dafny.MultiSet([636]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v113_) == (_dafny.Set({-2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_319_v114_).cf27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_319_v114_).cf28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_366_v150_).cf29) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_367_v151_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[0])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[1])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[2])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[3])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[4])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[5])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[6])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[7])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[8])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[9])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[10])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[11])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[12])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[13])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[14])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[15])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[16])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[17])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[18])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[19])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[20])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[21])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[22])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[23])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[24])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[25])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_368_v152_)[26])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_369_v153_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_370_v154_) == (_dafny.Map({True: True}))))
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
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
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
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
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

class D1_DC6(D1, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0), _dafny.Array(None, 0), False)
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
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC10(D3, NamedTuple('DC10', [('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf14)}, {self.cf15.VerbatimString(True)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({self.cf10.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)

class D4_DC14(D4, NamedTuple('DC14', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC16(D4, NamedTuple('DC16', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC18(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)

class D5_DC18(D5, NamedTuple('DC18', [('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, None, _dafny.Map({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D6_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D6_DC23)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D6_DC24)

class D6_DC21(D6, NamedTuple('DC21', [('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({self.cf31.VerbatimString(True)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC22(D6, NamedTuple('DC22', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC24(D6, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC26(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D7_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D7_DC25)

class D7_DC26(D7, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC25(D7, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f3: bool = False
        self.f7: str = _dafny.CodePoint('D')
        self._f1: int = int(0)
        self._f2: bool = False
        self._f4: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8

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
    def f6(self):
        return self._f6
    @property
    def f8(self):
        return self._f8

class C0(T0):
    def  __init__(self):
        self._f9: bool = False
        self.f14: int = int(0)
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f13, f14, f9):
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f9 = f9

    def fm4(self, p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(self.f14, (0) - (-993))

    @property
    def f13(self):
        return self._f13

class C1(T0):
    def  __init__(self):
        self._f9: bool = False
        self._f11: bool = False
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f11, f12, f9):
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f9 = f9

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((self).f12) * ((self).f12)

    def fm8(self, p0, globalState):
        return (default__.safeModuloInt(len(_dafny.Map({(self).f12: not(False)})), (self).f12)) >= ((self).f12)

    def fm9(self, p0, p1, p2, globalState):
        return (self).f11

    def m4(self, p0, p1, p2, globalState):
        r0: str = _dafny.CodePoint('D')
        d_391_v0_: _dafny.Array
        nw82_ = _dafny.Array(D0.default()(), 15)
        d_391_v0_ = nw82_
        d_392_v1_: D0
        d_392_v1_ = D0_DC3()
        index78_ = default__.safeIndex(626, (d_391_v0_).length(0))
        (d_391_v0_)[index78_] = d_392_v1_
        d_393_v2_: _dafny.Array
        def lambda14_(d_394_i0_):
            return (self).f11

        init6_ = lambda14_
        nw83_ = _dafny.Array(None, 24)
        for i0_6_ in range(nw83_.length(0)):
            nw83_[i0_6_] = init6_(i0_6_)
        d_393_v2_ = nw83_
        index79_ = default__.safeIndex(328, (d_393_v2_).length(0))
        (d_393_v2_)[index79_] = ((self).f12) > ((self).f12)
        d_395_v3_: int
        d_395_v3_ = 229
        d_396_v4_: _dafny.MultiSet
        d_396_v4_ = _dafny.MultiSet([(self).f9, (self).f9, (self).f9])
        d_397_v5_: _dafny.Seq
        d_397_v5_ = _dafny.SeqWithoutIsStrInference([d_396_v4_])
        index80_ = default__.safeIndex(626, (d_391_v0_).length(0))
        index81_ = default__.safeIndex(328, (d_393_v2_).length(0))
        rhs57_ = (self).fm8((self).f9, globalState)
        rhs58_ = d_392_v1_
        rhs59_ = (self).f9
        rhs60_ = ((d_397_v5_)[default__.safeIndex(((self).f12) - (-958), len(d_397_v5_))]).cardinality
        lhs36_ = globalState
        lhs37_ = d_391_v0_
        lhs38_ = default__.safeIndex(626, (d_391_v0_).length(0))
        lhs39_ = d_393_v2_
        lhs40_ = default__.safeIndex(328, (d_393_v2_).length(0))
        lhs36_.f0 = rhs57_
        lhs37_[lhs38_] = rhs58_
        lhs39_[lhs40_] = rhs59_
        d_395_v3_ = rhs60_
        source6_ = ((d_391_v0_)[default__.safeIndex(626, (d_391_v0_).length(0))] if (self).f11 else d_392_v1_)
        if source6_.is_DC1:
            d_398___mcc_h0_ = source6_.cf1
            d_399_cf1_ = d_398___mcc_h0_
            d_400_v6_: C0
            nw84_ = C0()
            nw84_.ctor__((self).f11, p1, (self).f9)
            d_400_v6_ = nw84_
            d_401_v7_: _dafny.Array
            nw85_ = _dafny.Array(int(0), 25)
            d_401_v7_ = nw85_
            index82_ = default__.safeIndex(574, (d_401_v7_).length(0))
            (d_401_v7_)[index82_] = (d_395_v3_) - (316)
            index83_ = default__.safeIndex(574, (d_401_v7_).length(0))
            (d_401_v7_)[index83_] = d_400_v6_.f14
            d_402_v8_: _dafny.Map
            d_402_v8_ = _dafny.Map({p1: True})
            d_403_v9_: int
            out18_: int
            out18_ = default__.m0(d_400_v6_.f14, d_401_v7_, (self).fm9((d_401_v7_)[default__.safeIndex(574, (d_401_v7_).length(0))], (d_400_v6_).f13, (self).f12, globalState), d_402_v8_, globalState)
            d_403_v9_ = out18_
            (globalState).f0 = (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]
        elif source6_.is_DC2:
            d_404___mcc_h1_ = source6_.cf2
            d_405___mcc_h2_ = source6_.cf3
            d_406___mcc_h3_ = source6_.cf4
            d_407_cf4_ = d_406___mcc_h3_
            d_408_cf3_ = d_405___mcc_h2_
            d_409_cf2_ = d_404___mcc_h1_
            d_409_cf2_ = (self).f11
            d_393_v2_ = d_393_v2_
            d_410_v10_: _dafny.Map
            d_410_v10_ = _dafny.Map({_dafny.Set({d_396_v4_}): (self).f11})
            d_411_v11_: _dafny.Seq
            d_411_v11_ = _dafny.SeqWithoutIsStrInference([True])
            d_412_v12_: _dafny.Set
            d_412_v12_ = _dafny.Set({_dafny.MultiSet(d_411_v11_), d_396_v4_})
            d_410_v10_ = (d_410_v10_).set(d_412_v12_, not((self).f9))
            d_413_v13_: D0
            d_413_v13_ = D0_DC1((d_411_v11_) <= (d_411_v11_))
            source7_ = d_413_v13_
            if source7_.is_DC1:
                d_414___mcc_h5_ = source7_.cf1
                d_415_cf1_ = d_414___mcc_h5_
                d_395_v3_ = d_408_cf3_
                (globalState).f0 = (d_411_v11_)[default__.safeIndex((self).f12, len(d_411_v11_))]
                d_416_v14_: _dafny.Map
                d_416_v14_ = _dafny.Map({(self).f12: d_409_cf2_})
                d_417_v15_: _dafny.Seq
                d_417_v15_ = _dafny.SeqWithoutIsStrInference([d_411_v11_, d_411_v11_, _dafny.SeqWithoutIsStrInference([((d_416_v14_)[874] if (874) in (d_416_v14_) else (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])])])
                d_418_v16_: _dafny.Array
                nw86_ = _dafny.Array(None, 4)
                nw86_[int(0)] = d_411_v11_
                nw86_[int(1)] = (d_411_v11_) + (d_411_v11_)
                nw86_[int(2)] = d_411_v11_
                nw86_[int(3)] = (d_417_v15_)[default__.safeIndex(d_408_cf3_, len(d_417_v15_))]
                d_418_v16_ = nw86_
                d_418_v16_ = d_418_v16_
                rhs61_ = d_416_v14_
                rhs62_ = (((0) - ((0) - (len(_dafny.Set({d_409_cf2_})))) if (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))] else (self).f12)) * ((self).f12)
                rhs63_ = (d_393_v2_ if (self).fm8((self).f11, globalState) else d_393_v2_)
                d_416_v14_ = rhs61_
                d_408_cf3_ = rhs62_
                d_393_v2_ = rhs63_
            elif source7_.is_DC2:
                d_419___mcc_h6_ = source7_.cf2
                d_420___mcc_h7_ = source7_.cf3
                d_421___mcc_h8_ = source7_.cf4
                d_422_cf4_ = d_421___mcc_h8_
                d_423_cf3_ = d_420___mcc_h7_
                d_424_cf2_ = d_419___mcc_h6_
                d_425_v17_: _dafny.Map
                d_425_v17_ = _dafny.Map({not((self).f9): (self).f11})
                d_426_v18_: _dafny.Map
                d_426_v18_ = _dafny.Map({(self).f12: d_395_v3_})
                index84_ = default__.safeIndex(328, (d_393_v2_).length(0))
                index85_ = default__.safeIndex(328, (d_393_v2_).length(0))
                rhs64_ = ((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]) == ((self).f9)
                rhs65_ = (d_411_v11_)[default__.safeIndex(len(d_425_v17_), len(d_411_v11_))]
                rhs66_ = (self).f11
                rhs67_ = d_395_v3_
                rhs68_ = (((_dafny.SeqWithoutIsStrInference([False])).set(default__.safeIndex(d_395_v3_, len(_dafny.SeqWithoutIsStrInference([False]))), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])) + (d_411_v11_)) + ((_dafny.SeqWithoutIsStrInference([d_424_cf2_, (self).f9, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]]) if (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))] else ((default__.fm10(globalState)).set(default__.safeIndex(len(d_426_v18_), len(default__.fm10(globalState))), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])).set(default__.safeIndex(d_395_v3_, len((default__.fm10(globalState)).set(default__.safeIndex(len(d_426_v18_), len(default__.fm10(globalState))), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]))), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])))
                lhs41_ = globalState
                lhs42_ = d_393_v2_
                lhs43_ = default__.safeIndex(328, (d_393_v2_).length(0))
                lhs44_ = d_393_v2_
                lhs45_ = default__.safeIndex(328, (d_393_v2_).length(0))
                lhs41_.f3 = rhs64_
                lhs42_[lhs43_] = rhs65_
                lhs44_[lhs45_] = rhs66_
                d_423_cf3_ = rhs67_
                d_411_v11_ = rhs68_
                (globalState).f0 = d_424_cf2_
                d_395_v3_ = p1
                d_427_v19_: _dafny.Map
                d_427_v19_ = _dafny.Map({d_424_cf2_: p2})
                d_428_v21_: _dafny.Set
                def iife32_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in _dafny.IntegerRange(-922, 289):
                        d_429_v20_: int = compr_10_
                        if ((-922) <= (d_429_v20_)) and ((d_429_v20_) < (289)):
                            coll10_[(d_429_v20_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aee")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_430_i1_ in range(default__.abs(640))])])))] = (self).f9
                    return _dafny.Map(coll10_)
                d_428_v21_ = _dafny.Set({d_395_v3_, (self).f12, len(iife32_()
                ), d_423_cf3_, (self).f12})
                d_431_v22_: _dafny.Seq
                d_431_v22_ = _dafny.SeqWithoutIsStrInference([d_428_v21_])
                d_427_v19_ = (d_427_v19_).set((d_428_v21_).ispropersubset((d_431_v22_)[default__.safeIndex((self).f12, len(d_431_v22_))]), _dafny.SeqWithoutIsStrInference([d_395_v3_ for d_432_i2_ in range(default__.abs(913))]))
            elif source7_.is_DC3:
                d_433_v23_: str
                d_433_v23_ = _dafny.CodePoint('q')
                d_434_v24_: _dafny.Map
                d_434_v24_ = _dafny.Map({(self).f9: d_411_v11_})
                d_435_v25_: _dafny.Map
                d_435_v25_ = _dafny.Map({d_433_v23_: (self).fm4(d_434_v24_, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (self).fm8(d_409_cf2_, globalState), (self).f12, globalState)})
                d_435_v25_ = (d_435_v25_).set(d_433_v23_, (self).f12)
                d_436_v26_: _dafny.Map
                d_436_v26_ = _dafny.Map({(d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]: not(d_409_cf2_)})
                d_436_v26_ = (d_436_v26_).set((self).fm8((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], globalState), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])
                d_437_v27_: _dafny.Array
                nw87_ = _dafny.Array(_dafny.Map({}), 16)
                d_437_v27_ = nw87_
                d_437_v27_ = d_437_v27_
                d_438_v28_: _dafny.Seq
                d_438_v28_ = _dafny.SeqWithoutIsStrInference([d_395_v3_, (self).f12, default__.safeModuloInt(d_395_v3_, (self).f12)])
                d_438_v28_ = p2
            elif True:
                d_439___mcc_h9_ = source7_.cf0
                d_440_cf0_ = d_439___mcc_h9_
                d_441_v29_: _dafny.Map
                d_441_v29_ = _dafny.Map({(d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]: d_408_cf3_})
                d_442_v30_: _dafny.Map
                d_442_v30_ = _dafny.Map({(self).f11: d_441_v29_})
                d_443_v31_: D0
                d_443_v31_ = D0_DC2((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], len(d_442_v30_), d_407_cf4_)
                d_444_v32_: _dafny.MultiSet
                d_444_v32_ = _dafny.MultiSet([d_443_v31_, d_443_v31_])
                d_445_v33_: _dafny.Map
                d_445_v33_ = _dafny.Map({d_408_cf3_: d_409_cf2_})
                d_446_v34_: _dafny.Set
                d_446_v34_ = _dafny.Set({d_408_cf3_, len(d_445_v33_)})
                d_447_v35_: str
                d_447_v35_ = _dafny.CodePoint('k')
                d_448_v36_: D1
                d_448_v36_ = D1_DC4(_dafny.SeqWithoutIsStrInference([p1 for d_449_i3_ in range(default__.abs(106))]))
                rhs69_ = (d_444_v32_).ispropersubset((d_444_v32_ if (self).f9 else d_444_v32_))
                rhs70_ = ((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]) == ((d_446_v34_) == (d_446_v34_))
                rhs71_ = d_447_v35_
                rhs72_ = (self).f9
                rhs73_ = len((d_448_v36_).cf5)
                lhs46_ = globalState
                lhs47_ = globalState
                lhs48_ = globalState
                d_409_cf2_ = rhs69_
                lhs46_.f3 = rhs70_
                lhs47_.f7 = rhs71_
                lhs48_.f3 = rhs72_
                d_408_cf3_ = rhs73_
                d_450_v37_: _dafny.Set
                d_450_v37_ = _dafny.Set({d_393_v2_, d_393_v2_, d_393_v2_})
                d_451_v38_: _dafny.Seq
                d_451_v38_ = _dafny.SeqWithoutIsStrInference([d_450_v37_, d_450_v37_, d_450_v37_, (d_450_v37_) | (d_450_v37_)])
                d_452_v39_: _dafny.Seq
                d_452_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrmm"))
                d_450_v37_ = (d_451_v38_)[default__.safeIndex((0) - (default__.fm3((d_452_v39_)[default__.safeIndex(d_408_cf3_, len(d_452_v39_))], (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], -349, globalState)), len(d_451_v38_))]
                index86_ = default__.safeIndex(328, (d_393_v2_).length(0))
                (d_393_v2_)[index86_] = (d_443_v31_).cf2
                d_453_v40_: _dafny.Map
                d_453_v40_ = _dafny.Map({(d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]: (default__.fm11((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], globalState)) + (d_452_v39_)})
                d_440_cf0_ = len(d_453_v40_)
        elif source6_.is_DC3:
            d_454_v41_: _dafny.MultiSet
            d_454_v41_ = _dafny.MultiSet([d_395_v3_, p1, d_395_v3_])
            d_455_v42_: D0
            d_455_v42_ = D0_DC1((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])
            d_456_v43_: C0
            nw88_ = C0()
            nw88_.ctor__((self).fm9(p1, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], p1, globalState), (((d_454_v41_)[973] if (973) in (d_454_v41_) else d_395_v3_) if (self).f11 else d_395_v3_), (d_455_v42_).cf1)
            d_456_v43_ = nw88_
            d_456_v43_ = d_456_v43_
            d_457_v44_: _dafny.Seq
            d_457_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            d_458_v45_: _dafny.Map
            d_458_v45_ = _dafny.Map({d_457_v44_: d_456_v43_.f14})
            d_459_v46_: _dafny.Map
            d_459_v46_ = _dafny.Map({(d_458_v45_) | (d_458_v45_): (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]})
            d_459_v46_ = (d_459_v46_).set(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lh")): d_395_v3_}), (self).f11)
            if ((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgfq"))) < (d_457_v44_) else (self).f11):
                d_460_v47_: _dafny.Seq
                d_460_v47_ = _dafny.SeqWithoutIsStrInference([True])
                d_461_v48_: _dafny.Map
                d_461_v48_ = _dafny.Map({False: (d_460_v47_)[default__.safeIndex((self).f12, len(d_460_v47_))]})
                d_462_v49_: _dafny.Array
                nw89_ = _dafny.Array(None, 16)
                nw89_[int(0)] = ((d_454_v41_) - (d_454_v41_)).cardinality
                nw89_[int(1)] = (d_395_v3_ if (d_456_v43_).f13 else len(d_457_v44_))
                nw89_[int(2)] = d_395_v3_
                nw89_[int(3)] = len(d_457_v44_)
                nw89_[int(4)] = (0) - (d_395_v3_)
                nw89_[int(5)] = -221
                nw89_[int(6)] = (d_456_v43_.f14 if (d_456_v43_).f13 else d_456_v43_.f14)
                nw89_[int(7)] = len(d_457_v44_)
                nw89_[int(8)] = d_456_v43_.f14
                nw89_[int(9)] = d_456_v43_.f14
                nw89_[int(10)] = d_456_v43_.f14
                nw89_[int(11)] = d_456_v43_.f14
                nw89_[int(12)] = (len(d_461_v48_) if (self).f11 else len(_dafny.Set({d_456_v43_, d_456_v43_})))
                nw89_[int(13)] = (0) - (d_456_v43_.f14)
                nw89_[int(14)] = p1
                nw89_[int(15)] = 728
                d_462_v49_ = nw89_
                index87_ = default__.safeIndex(401, (d_462_v49_).length(0))
                (d_462_v49_)[index87_] = ((self).f12) - (p1)
                index88_ = default__.safeIndex(401, (d_462_v49_).length(0))
                (d_462_v49_)[index88_] = (0) - (default__.safeDivisionInt(d_456_v43_.f14, len((d_457_v44_) + (d_457_v44_))))
                d_463_v50_: D2
                d_463_v50_ = D2_DC7(d_393_v2_)
                d_464_v51_: _dafny.Seq
                d_464_v51_ = _dafny.SeqWithoutIsStrInference([d_393_v2_, d_393_v2_, (d_463_v50_).cf6])
                d_464_v51_ = ((d_464_v51_ if (self).f9 else d_464_v51_) if (self).f11 else d_464_v51_)
                d_465_v52_: _dafny.MultiSet
                d_465_v52_ = _dafny.MultiSet([d_454_v41_, d_454_v41_, d_454_v41_])
                d_466_v53_: D0
                d_466_v53_ = D0_DC2((d_456_v43_).f13, p1, d_465_v52_)
                d_467_v54_: D2
                d_467_v54_ = D2_DC8((d_466_v53_).cf2, (self).f12, (d_456_v43_).f13)
                d_468_v55_: C0
                nw90_ = C0()
                nw90_.ctor__((self).f11, (d_467_v54_).cf8, (d_456_v43_).f13)
                d_468_v55_ = nw90_
                d_469_v56_: str
                d_469_v56_ = _dafny.CodePoint('i')
                (globalState).f7 = d_469_v56_
                (d_456_v43_).f14 = d_456_v43_.f14
            elif True:
                d_470_v57_: _dafny.Map
                d_470_v57_ = _dafny.Map({d_456_v43_.f14: default__.safeModuloInt(186, d_456_v43_.f14)})
                d_471_v58_: str
                d_471_v58_ = _dafny.CodePoint('r')
                d_472_v59_: _dafny.Map
                d_472_v59_ = _dafny.Map({_dafny.CodePoint('m'): default__.fm12(True, (d_456_v43_).f13, d_456_v43_.f14, True, globalState)})
                d_473_v60_: _dafny.Seq
                d_473_v60_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11, (d_456_v43_).f13, (_dafny.Map({d_471_v58_: d_396_v4_})) == (d_472_v59_)])
                d_474_v61_: _dafny.Map
                d_474_v61_ = _dafny.Map({(d_456_v43_).f13: d_471_v58_})
                d_475_v62_: _dafny.Set
                d_475_v62_ = _dafny.Set({(self).fm8((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], globalState)})
                index89_ = default__.safeIndex(328, (d_393_v2_).length(0))
                rhs74_ = not((d_473_v60_)[default__.safeIndex(default__.safeDivisionInt(d_456_v43_.f14, p1), len(d_473_v60_))])
                rhs75_ = (len((d_474_v61_) | (d_474_v61_))) + (len(d_475_v62_))
                rhs76_ = (d_456_v43_).f13
                rhs77_ = _dafny.Map({p1: (len(d_457_v44_)) * (d_456_v43_.f14)})
                lhs49_ = globalState
                lhs50_ = d_456_v43_
                lhs51_ = d_393_v2_
                lhs52_ = default__.safeIndex(328, (d_393_v2_).length(0))
                lhs49_.f0 = rhs74_
                lhs50_.f14 = rhs75_
                lhs51_[lhs52_] = rhs76_
                d_470_v57_ = rhs77_
                (d_456_v43_).f14 = d_456_v43_.f14
                (d_456_v43_).f14 = (self).f12
                d_457_v44_ = d_457_v44_
                d_476_v63_: _dafny.Map
                d_476_v63_ = _dafny.Map({d_456_v43_.f14: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xncydgy"))})
                d_477_v64_: D0
                d_477_v64_ = D0_DC0(d_456_v43_.f14)
                d_476_v63_ = _dafny.Map({(d_477_v64_).cf0: default__.fm11((d_456_v43_).f13, (self).f11, globalState)})
            d_457_v44_ = (d_457_v44_) + (d_457_v44_)
        elif True:
            d_478___mcc_h4_ = source6_.cf0
            d_479_cf0_ = d_478___mcc_h4_
            d_480_v65_: str
            d_480_v65_ = _dafny.CodePoint('j')
            d_481_v66_: _dafny.Seq
            d_481_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
            d_482_v67_: _dafny.MultiSet
            d_482_v67_ = _dafny.MultiSet([default__.fm3(d_480_v65_, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (self).f12, globalState), len(d_481_v66_)])
            d_483_v68_: D0
            d_483_v68_ = D0_DC0(d_479_cf0_)
            d_395_v3_ = ((D0_DC0(((d_482_v67_)[278] if (278) in (d_482_v67_) else (self).f12)) if (self).f9 else d_483_v68_)).cf0
            d_484_v69_: C0
            nw91_ = C0()
            nw91_.ctor__(False, ((0) - ((0) - ((0) - (d_479_cf0_))) if (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))] else d_395_v3_), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])
            d_484_v69_ = nw91_
            d_485_v70_: _dafny.Set
            d_485_v70_ = _dafny.Set({True})
            d_485_v70_ = (d_485_v70_) | (d_485_v70_)
            d_486_v71_: _dafny.Map
            d_486_v71_ = _dafny.Map({(d_484_v69_).f13: d_484_v69_.f14})
            d_487_v72_: D3
            d_487_v72_ = D3_DC11(d_479_cf0_, d_481_v66_, len(d_486_v71_), d_484_v69_)
            d_488_v73_: _dafny.Seq
            d_488_v73_ = _dafny.SeqWithoutIsStrInference([(d_487_v72_).cf15, d_481_v66_, d_481_v66_])
            d_489_v74_: _dafny.Map
            d_489_v74_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_395_v3_ for d_490_i4_ in range(default__.abs(692))])) + (p2): d_488_v73_})
            d_488_v73_ = ((d_489_v74_)[_dafny.SeqWithoutIsStrInference([p1 for d_491_i5_ in range(default__.abs(-175))])] if (_dafny.SeqWithoutIsStrInference([p1 for d_492_i5_ in range(default__.abs(-175))])) in (d_489_v74_) else d_488_v73_)
        d_493_v75_: _dafny.Array
        def lambda15_(d_494_i6_):
            return _dafny.CodePoint('d')

        init7_ = lambda15_
        nw92_ = _dafny.Array(None, 21)
        for i0_7_ in range(nw92_.length(0)):
            nw92_[i0_7_] = init7_(i0_7_)
        d_493_v75_ = nw92_
        d_495_v76_: str
        d_495_v76_ = _dafny.CodePoint('e')
        index90_ = default__.safeIndex(535, (d_493_v75_).length(0))
        (d_493_v75_)[index90_] = d_495_v76_
        index91_ = default__.safeIndex(535, (d_493_v75_).length(0))
        (d_493_v75_)[index91_] = d_495_v76_
        d_496_i7_: int
        d_496_i7_ = 0
        with _dafny.label("1"):
            while (self).f9:
                with _dafny.c_label("1"):
                    if (d_496_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_496_i7_ = (d_496_i7_) + (1)
                    d_497_v77_: _dafny.MultiSet
                    d_497_v77_ = _dafny.MultiSet([len(_dafny.Set({d_395_v3_}))])
                    d_498_v78_: _dafny.MultiSet
                    d_498_v78_ = _dafny.MultiSet([d_497_v77_])
                    d_499_v79_: D0
                    d_499_v79_ = D0_DC2((self).f11, p1, d_498_v78_)
                    d_500_v80_: C0
                    nw93_ = C0()
                    nw93_.ctor__(((self).f9) and ((self).f9), (-293) + ((d_499_v79_).cf3), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])
                    d_500_v80_ = nw93_
                    d_501_v81_: _dafny.Array
                    nw94_ = _dafny.Array(None, 1)
                    nw94_[int(0)] = p2
                    d_501_v81_ = nw94_
                    index92_ = default__.safeIndex(207, (d_501_v81_).length(0))
                    (d_501_v81_)[index92_] = default__.fm13((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], True, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], d_500_v80_.f14, globalState)
                    d_502_v82_: _dafny.Seq
                    d_502_v82_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                    index93_ = default__.safeIndex(207, (d_501_v81_).length(0))
                    rhs78_ = default__.fm3((d_493_v75_)[default__.safeIndex(535, (d_493_v75_).length(0))], not((d_500_v80_).f13), len((d_502_v82_) + (d_502_v82_)), globalState)
                    rhs79_ = d_395_v3_
                    rhs80_ = p2
                    lhs53_ = d_500_v80_
                    lhs54_ = d_500_v80_
                    lhs55_ = d_501_v81_
                    lhs56_ = default__.safeIndex(207, (d_501_v81_).length(0))
                    lhs53_.f14 = rhs78_
                    lhs54_.f14 = rhs79_
                    lhs55_[lhs56_] = rhs80_
                    d_503_v83_: D1
                    d_503_v83_ = D1_DC6()
                    source8_ = (d_503_v83_ if (d_500_v80_).f13 else d_503_v83_)
                    if source8_.is_DC5:
                        d_504_v84_: _dafny.Map
                        d_504_v84_ = _dafny.Map({d_393_v2_: d_500_v80_.f14})
                        d_505_v85_: _dafny.Map
                        d_505_v85_ = _dafny.Map({len(d_502_v82_): d_500_v80_.f14})
                        d_504_v84_ = (d_504_v84_).set(d_393_v2_, len((d_505_v85_) | (d_505_v85_)))
                        d_506_v86_: _dafny.Array
                        nw95_ = _dafny.Array(int(0), 12)
                        d_506_v86_ = nw95_
                        index94_ = default__.safeIndex(480, (d_506_v86_).length(0))
                        (d_506_v86_)[index94_] = default__.safeDivisionInt(950, d_500_v80_.f14)
                        index95_ = default__.safeIndex(480, (d_506_v86_).length(0))
                        (d_506_v86_)[index95_] = default__.safeDivisionInt(d_500_v80_.f14, (d_500_v80_.f14 if True else p1))
                        d_507_v87_: _dafny.Map
                        d_507_v87_ = _dafny.Map({True: d_495_v76_})
                        d_508_v88_: D0
                        d_508_v88_ = D0_DC0(p1)
                        d_509_v89_: _dafny.Map
                        d_509_v89_ = _dafny.Map({((d_507_v87_)[(self).f11] if ((self).f11) in (d_507_v87_) else d_495_v76_): d_508_v88_})
                        d_509_v89_ = (d_509_v89_).set((d_493_v75_)[default__.safeIndex(535, (d_493_v75_).length(0))], d_508_v88_)
                        index96_ = default__.safeIndex(480, (d_506_v86_).length(0))
                        (d_506_v86_)[index96_] = ((self).f12) * (p1)
                    elif source8_.is_DC6:
                        d_510_v90_: _dafny.Seq
                        d_510_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpamgbonn"))
                        d_511_v91_: _dafny.Set
                        d_511_v91_ = _dafny.Set({(_dafny.MultiSet([(d_500_v80_).f13])).isdisjoint(d_396_v4_), (d_500_v80_).f13, (d_500_v80_).f13})
                        d_512_v92_: _dafny.Set
                        d_512_v92_ = _dafny.Set({(0) - (d_500_v80_.f14)})
                        d_513_v93_: _dafny.Seq
                        d_513_v93_ = _dafny.SeqWithoutIsStrInference([d_512_v92_, d_512_v92_, d_512_v92_, d_512_v92_])
                        d_514_v94_: _dafny.Map
                        d_514_v94_ = _dafny.Map({(d_500_v80_).f13: d_511_v91_})
                        rhs81_ = (((d_510_v90_) + (d_510_v90_)).set(default__.safeIndex(753, len((d_510_v90_) + (d_510_v90_))), d_495_v76_)) + (d_510_v90_)
                        rhs82_ = (_dafny.MultiSet((d_513_v93_ if (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))] else d_513_v93_))).issubset(_dafny.MultiSet([d_512_v92_, _dafny.Set({p1})]))
                        rhs83_ = (((d_514_v94_)[True] if (True) in (d_514_v94_) else _dafny.Set({(self).f11, (self).f9}))) - (d_511_v91_)
                        lhs57_ = globalState
                        d_510_v90_ = rhs81_
                        lhs57_.f3 = rhs82_
                        d_511_v91_ = rhs83_
                        d_515_v95_: _dafny.Map
                        d_515_v95_ = _dafny.Map({(self).f12: (d_500_v80_).f13})
                        (globalState).f0 = ((d_500_v80_).f13) or (((d_515_v95_)[(0) - (p1)] if ((0) - (p1)) in (d_515_v95_) else (self).f9))
                        (globalState).f7 = d_495_v76_
                        d_516_v96_: _dafny.Seq
                        d_516_v96_ = _dafny.SeqWithoutIsStrInference([d_510_v90_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx")), d_510_v90_])
                        d_517_v97_: T0
                        nw96_ = C0()
                        nw96_.ctor__((d_516_v96_) < (d_516_v96_), p1, (d_500_v80_).f13)
                        d_517_v97_ = nw96_
                        index97_ = default__.safeIndex(626, (d_391_v0_).length(0))
                        def iife33_():
                            coll11_ = _dafny.Map()
                            compr_11_: int
                            for compr_11_ in _dafny.IntegerRange(736, -955):
                                d_518_v98_: int = compr_11_
                                if ((736) <= (d_518_v98_)) and ((d_518_v98_) < (-955)):
                                    def iife34_():
                                        def iife36_():
                                            coll14_ = _dafny.Map()
                                            compr_14_: int
                                            for compr_14_ in ((d_501_v81_)[default__.safeIndex(207, (d_501_v81_).length(0))]).Elements:
                                                d_519_v100_: int = compr_14_
                                                if (d_519_v100_) in ((d_501_v81_)[default__.safeIndex(207, (d_501_v81_).length(0))]):
                                                    coll14_[(d_519_v100_) + ((0) - (p1))] = d_500_v80_.f14
                                            return _dafny.Map(coll14_)
                                        coll12_ = _dafny.Map()
                                        def iife35_():
                                            coll13_ = _dafny.Map()
                                            compr_13_: int
                                            for compr_13_ in ((d_501_v81_)[default__.safeIndex(207, (d_501_v81_).length(0))]).Elements:
                                                d_519_v100_: int = compr_13_
                                                if (d_519_v100_) in ((d_501_v81_)[default__.safeIndex(207, (d_501_v81_).length(0))]):
                                                    coll13_[(d_519_v100_) + ((0) - (p1))] = d_500_v80_.f14
                                            return _dafny.Map(coll13_)
                                        compr_12_: D0
                                        for compr_12_ in (_dafny.Map({D0_DC0((self).f12): len(iife35_()
                                        )})).keys.Elements:
                                            d_520_v99_: D0 = compr_12_
                                            if (d_520_v99_) in (_dafny.Map({D0_DC0((self).f12): len(iife36_()
                                            )})):
                                                coll12_[d_520_v99_] = d_395_v3_
                                        return _dafny.Map(coll12_)
                                    coll11_[default__.safeDivisionInt(d_518_v98_, d_500_v80_.f14)] = iife34_()

                            return _dafny.Map(coll11_)
                        rhs84_ = d_393_v2_
                        rhs85_ = d_517_v97_
                        rhs86_ = (len(d_511_v91_)) not in (iife33_()
                        )
                        rhs87_ = (d_391_v0_)[default__.safeIndex(626, (d_391_v0_).length(0))]
                        rhs88_ = ((self).f11) and ((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])
                        lhs58_ = globalState
                        lhs59_ = d_391_v0_
                        lhs60_ = default__.safeIndex(626, (d_391_v0_).length(0))
                        lhs61_ = globalState
                        d_393_v2_ = rhs84_
                        d_517_v97_ = rhs85_
                        lhs58_.f0 = rhs86_
                        lhs59_[lhs60_] = rhs87_
                        lhs61_.f0 = rhs88_
                    elif True:
                        d_521___mcc_h10_ = source8_.cf5
                        d_522_cf5_ = d_521___mcc_h10_
                        (globalState).f0 = (self).f9
                        d_523_v101_: _dafny.Array
                        nw97_ = _dafny.Array(int(0), 8)
                        d_523_v101_ = nw97_
                        index98_ = default__.safeIndex(758, (d_523_v101_).length(0))
                        (d_523_v101_)[index98_] = default__.safeModuloInt(p1, d_500_v80_.f14)
                        index99_ = default__.safeIndex(758, (d_523_v101_).length(0))
                        (d_523_v101_)[index99_] = d_395_v3_
                        d_395_v3_ = default__.safeModuloInt(148, d_395_v3_)
                        d_524_v102_: D0
                        d_524_v102_ = D0_DC1((self).fm8(True, globalState))
                        d_525_v103_: _dafny.Seq
                        d_525_v103_ = _dafny.SeqWithoutIsStrInference([(d_501_v81_)[default__.safeIndex(207, (d_501_v81_).length(0))]])
                        d_526_v104_: _dafny.Map
                        d_526_v104_ = _dafny.Map({d_502_v82_: d_500_v80_.f14})
                        d_527_v105_: _dafny.Set
                        d_527_v105_ = _dafny.Set({(0) - ((d_500_v80_.f14) * (len(d_525_v103_))), len(d_526_v104_)})
                        pat_let_tv18_ = d_500_v80_
                        def iife37_(_pat_let11_0):
                            def iife38_(d_528_dt__update__tmp_h0_):
                                def iife39_(_pat_let12_0):
                                    def iife40_(d_529_dt__update_hcf1_h0_):
                                        return D0_DC1(d_529_dt__update_hcf1_h0_)
                                    return iife40_(_pat_let12_0)
                                return iife39_((pat_let_tv18_).f13)
                            return iife38_(_pat_let11_0)
                        rhs89_ = iife37_(d_524_v102_)
                        rhs90_ = d_527_v105_
                        d_524_v102_ = rhs89_
                        d_527_v105_ = rhs90_
                    d_530_v106_: _dafny.Seq
                    d_530_v106_ = _dafny.SeqWithoutIsStrInference([d_495_v76_])
                    d_531_v107_: D3
                    d_531_v107_ = D3_DC9(d_530_v106_)
                    d_532_v108_: _dafny.Map
                    d_532_v108_ = _dafny.Map({(self).f11: d_530_v106_})
                    index100_ = default__.safeIndex(328, (d_393_v2_).length(0))
                    index101_ = default__.safeIndex(328, (d_393_v2_).length(0))
                    rhs91_ = (d_500_v80_.f14) > ((d_500_v80_.f14) + (len((d_531_v107_).cf10)))
                    rhs92_ = (d_532_v108_) != (d_532_v108_)
                    rhs93_ = p1
                    rhs94_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mm")))) - (len((d_501_v81_)[default__.safeIndex(207, (d_501_v81_).length(0))]))
                    lhs62_ = d_393_v2_
                    lhs63_ = default__.safeIndex(328, (d_393_v2_).length(0))
                    lhs64_ = d_393_v2_
                    lhs65_ = default__.safeIndex(328, (d_393_v2_).length(0))
                    lhs62_[lhs63_] = rhs91_
                    lhs64_[lhs65_] = rhs92_
                    d_395_v3_ = rhs93_
                    d_395_v3_ = rhs94_
                    pass
            pass
        d_533_v109_: _dafny.Map
        d_533_v109_ = _dafny.Map({(self).f12: p1})
        d_534_v110_: _dafny.MultiSet
        d_534_v110_ = _dafny.MultiSet([336, (0) - ((self).f12)])
        d_535_v111_: _dafny.Seq
        d_535_v111_ = _dafny.SeqWithoutIsStrInference([d_534_v110_])
        d_536_v112_: _dafny.Array
        nw98_ = _dafny.Array(None, 21)
        nw98_[int(0)] = len(d_533_v109_)
        nw98_[int(1)] = d_395_v3_
        nw98_[int(2)] = p1
        nw98_[int(3)] = (0) - ((self).f12)
        nw98_[int(4)] = (self).f12
        nw98_[int(5)] = ((d_396_v4_)[(d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]] if ((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]) in (d_396_v4_) else d_395_v3_)
        nw98_[int(6)] = (self).f12
        nw98_[int(7)] = 759
        nw98_[int(8)] = d_395_v3_
        nw98_[int(9)] = (0) - (p1)
        nw98_[int(10)] = p1
        nw98_[int(11)] = p1
        nw98_[int(12)] = -591
        nw98_[int(13)] = ((d_535_v111_)[default__.safeIndex(p1, len(d_535_v111_))]).cardinality
        nw98_[int(14)] = d_395_v3_
        nw98_[int(15)] = p1
        nw98_[int(16)] = d_395_v3_
        nw98_[int(17)] = (self).f12
        nw98_[int(18)] = len(p2)
        nw98_[int(19)] = (self).f12
        nw98_[int(20)] = (0) - ((self).f12)
        d_536_v112_ = nw98_
        d_537_v113_: _dafny.Map
        d_537_v113_ = _dafny.Map({d_536_v112_: (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]})
        d_537_v113_ = (d_537_v113_).set(d_536_v112_, not(not (default__.fm2(globalState)) or ((self).f11)))
        d_538_v114_: D2
        d_538_v114_ = D2_DC7(d_393_v2_)
        source9_ = d_538_v114_
        if source9_.is_DC8:
            d_539___mcc_h11_ = source9_.cf7
            d_540___mcc_h12_ = source9_.cf8
            d_541___mcc_h13_ = source9_.cf9
            d_542_cf9_ = d_541___mcc_h13_
            d_543_cf8_ = d_540___mcc_h12_
            d_544_cf7_ = d_539___mcc_h11_
            d_543_cf8_ = (self).f12
            d_545_v115_: _dafny.Seq
            d_545_v115_ = _dafny.SeqWithoutIsStrInference([False, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]])
            d_546_v116_: _dafny.Map
            d_546_v116_ = _dafny.Map({(self).f9: d_545_v115_})
            d_547_v117_: C0
            nw99_ = C0()
            nw99_.ctor__(d_542_cf9_, default__.safeModuloInt((self).fm4(d_546_v116_, (self).f9, (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (self).f12, globalState), d_543_cf8_), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))])
            d_547_v117_ = nw99_
            d_544_cf7_ = default__.fm2(globalState)
            d_548_v118_: _dafny.Set
            d_548_v118_ = _dafny.Set({(d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (336) <= (p1), (self).f9})
            d_548_v118_ = (d_548_v118_) | (_dafny.Set({not(d_542_cf9_), (self).f11, (d_547_v117_).f13, not((d_547_v117_).f13), (d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))]}))
        elif True:
            d_549___mcc_h14_ = source9_.cf6
            d_550_cf6_ = d_549___mcc_h14_
            d_551_v119_: C0
            nw100_ = C0()
            nw100_.ctor__(False, len(default__.fm11((d_393_v2_)[default__.safeIndex(328, (d_393_v2_).length(0))], (self).f11, globalState)), (self).f9)
            d_551_v119_ = nw100_
            d_552_v120_: _dafny.Seq
            d_552_v120_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ji"))
            d_553_v121_: D3
            d_553_v121_ = D3_DC11(d_551_v119_.f14, d_552_v120_, d_551_v119_.f14, d_551_v119_)
            d_554_v122_: D3
            d_554_v122_ = D3_DC11((0) - (d_551_v119_.f14), d_552_v120_, d_551_v119_.f14, (d_553_v121_).cf17)
            d_551_v119_ = (d_553_v121_).cf17
            d_555_v123_: _dafny.Array
            nw101_ = _dafny.Array(_dafny.Seq({}), 13)
            d_555_v123_ = nw101_
            index102_ = default__.safeIndex(882, (d_555_v123_).length(0))
            (d_555_v123_)[index102_] = (D1_DC4(p2)).cf5
            index103_ = default__.safeIndex(882, (d_555_v123_).length(0))
            (d_555_v123_)[index103_] = _dafny.SeqWithoutIsStrInference([len(d_552_v120_)])
            d_552_v120_ = d_552_v120_
            d_556_v124_: _dafny.Seq
            d_556_v124_ = _dafny.SeqWithoutIsStrInference([(d_395_v3_) != (282), (self).f11, (d_551_v119_).f13])
            (globalState).f3 = (d_556_v124_)[default__.safeIndex((p1) + (p1), len(d_556_v124_))]
        r0 = (d_493_v75_)[default__.safeIndex(535, (d_493_v75_).length(0))]
        return r0

    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12

class C2(T0):
    def  __init__(self):
        self._f9: bool = False
        self.f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f10, f9):
        (self).f10 = f10
        (self)._f9 = f9

    def fm4(self, p0, p1, p2, p3, globalState):
        if (self).f9:
            return len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f9, not((self).f9)})), 550, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofopjn")))})), 983]): (self).f9}))
        elif True:
            return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwhmbs"))))

    def fm6(self, p0, p1, globalState):
        return 266

    def fm7(self, p0, p1, p2, globalState):
        source10_ = D0_DC2((self).f9, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekckaebw"))), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([268, 682])) for d_557_i0_ in range(default__.abs(690))])))
        if source10_.is_DC1:
            d_558___mcc_h0_ = source10_.cf1
            d_559_cf1_ = d_558___mcc_h0_
            return (_dafny.Map({558: self.f10})) | (_dafny.Map({-997: (self).f9}))
        elif source10_.is_DC2:
            d_560___mcc_h1_ = source10_.cf2
            d_561___mcc_h2_ = source10_.cf3
            d_562___mcc_h3_ = source10_.cf4
            d_563_cf4_ = d_562___mcc_h3_
            d_564_cf3_ = d_561___mcc_h2_
            d_565_cf2_ = d_560___mcc_h1_
            return _dafny.Map({(0) - (d_564_cf3_): self.f10})
        elif source10_.is_DC3:
            return (_dafny.Map({(_dafny.MultiSet([self.f10, (self).f9, True])).cardinality: self.f10})) | (_dafny.Map({(0) - (-643): self.f10}))
        elif True:
            d_566___mcc_h4_ = source10_.cf0
            d_567_cf0_ = d_566___mcc_h4_
            return _dafny.Map({d_567_cf0_: (self).f9})

    def m2(self, p0, p1, p2, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        d_568_v0_: _dafny.Seq
        d_568_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqqru"))
        d_568_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mysvhxlt"))
        d_569_v1_: _dafny.Seq
        d_569_v1_ = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])
        d_570_v2_: int
        d_570_v2_ = 579
        d_571_v3_: _dafny.Map
        d_571_v3_ = _dafny.Map({d_570_v2_: 21})
        d_572_v4_: _dafny.MultiSet
        d_572_v4_ = _dafny.MultiSet([self.f10, (d_569_v1_)[default__.safeIndex(len((d_569_v1_).set(default__.safeIndex(len(d_571_v3_), len(d_569_v1_)), (self).f9)), len(d_569_v1_))], p0])
        d_573_v5_: _dafny.MultiSet
        d_573_v5_ = _dafny.MultiSet([d_570_v2_])
        d_574_v6_: T0
        nw102_ = C0()
        nw102_.ctor__(self.f10, (((d_572_v4_)[p0] if (p0) in (d_572_v4_) else d_570_v2_)) + (d_570_v2_), (d_573_v5_).ispropersubset(d_573_v5_))
        d_574_v6_ = nw102_
        d_574_v6_ = (d_574_v6_ if not((d_574_v6_).f9) else d_574_v6_)
        (self).f10 = self.f10
        d_575_v7_: _dafny.Set
        d_575_v7_ = _dafny.Set({(d_574_v6_).f9})
        d_570_v2_ = (d_574_v6_).fm4(_dafny.Map({p0: d_569_v1_}), (d_570_v2_) <= (d_570_v2_), p0, len(d_575_v7_), globalState)
        d_576_v8_: _dafny.Seq
        d_576_v8_ = _dafny.SeqWithoutIsStrInference([default__.fm11(True, (d_574_v6_).f9, globalState)])
        (globalState).f0 = (len(d_576_v8_)) < (748)
        hi9_ = d_570_v2_
        for d_577_i0_ in range(d_570_v2_, hi9_):
            (globalState).f0 = not(self.f10)
            d_570_v2_ = ((default__.fm3((p2)[default__.safeIndex(d_570_v2_, len(p2))], p0, d_570_v2_, globalState)) + (d_577_i0_)) * (len(p2))
            d_578_v9_: _dafny.Array
            nw103_ = _dafny.Array(int(0), 1)
            d_578_v9_ = nw103_
            index104_ = default__.safeIndex(665, (d_578_v9_).length(0))
            (d_578_v9_)[index104_] = d_577_i0_
            index105_ = default__.safeIndex(665, (d_578_v9_).length(0))
            (d_578_v9_)[index105_] = d_577_i0_
            d_579_v10_: _dafny.Map
            d_579_v10_ = _dafny.Map({d_574_v6_: True})
            d_580_v11_: _dafny.Array
            nw104_ = _dafny.Array(None, 6)
            nw104_[int(0)] = _dafny.Map({d_574_v6_: p1})
            nw104_[int(1)] = d_579_v10_
            nw104_[int(2)] = (d_579_v10_ if True else d_579_v10_)
            nw104_[int(3)] = d_579_v10_
            nw104_[int(4)] = d_579_v10_
            nw104_[int(5)] = d_579_v10_
            d_580_v11_ = nw104_
            index106_ = default__.safeIndex(175, (d_580_v11_).length(0))
            (d_580_v11_)[index106_] = d_579_v10_
            d_581_v12_: _dafny.Array
            nw105_ = _dafny.Array(None, 2)
            nw105_[int(0)] = (self.f10) not in (d_569_v1_)
            nw105_[int(1)] = (p2) < (d_568_v0_)
            d_581_v12_ = nw105_
            index107_ = default__.safeIndex(688, (d_581_v12_).length(0))
            (d_581_v12_)[index107_] = default__.fm2(globalState)
            d_582_v13_: _dafny.Seq
            d_582_v13_ = _dafny.SeqWithoutIsStrInference([171, d_570_v2_])
            d_583_v14_: _dafny.Map
            d_583_v14_ = _dafny.Map({d_582_v13_: _dafny.SeqWithoutIsStrInference([len(d_568_v0_) for d_584_i1_ in range(default__.abs(390))])})
            index108_ = default__.safeIndex(175, (d_580_v11_).length(0))
            index109_ = default__.safeIndex(688, (d_581_v12_).length(0))
            rhs95_ = d_579_v10_
            rhs96_ = (d_574_v6_).f9
            rhs97_ = d_569_v1_
            rhs98_ = (self.f10) or ((d_574_v6_).f9)
            rhs99_ = (len(d_583_v14_)) - (d_577_i0_)
            lhs66_ = d_580_v11_
            lhs67_ = default__.safeIndex(175, (d_580_v11_).length(0))
            lhs68_ = d_581_v12_
            lhs69_ = default__.safeIndex(688, (d_581_v12_).length(0))
            lhs66_[lhs67_] = rhs95_
            r1 = rhs96_
            d_569_v1_ = rhs97_
            lhs68_[lhs69_] = rhs98_
            d_570_v2_ = rhs99_
        r0 = d_575_v7_
        r1 = (self.f10 if self.f10 else (self).f9)
        return r0, r1

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_585_v1_: _dafny.Array
        def lambda16_(d_586_p0_, d_587_p2_):
            def lambda17_(d_588_i0_):
                def iife41_():
                    coll15_ = _dafny.Set()
                    compr_15_: int
                    for compr_15_ in _dafny.IntegerRange(427, 24):
                        d_589_v0_: int = compr_15_
                        if ((427) <= (d_589_v0_)) and ((d_589_v0_) < (24)):
                            coll15_ = coll15_.union(_dafny.Set([(d_589_v0_) + (786)]))
                    return _dafny.Set(coll15_)
                return (iife41_()
                ).ispropersubset(_dafny.Set({d_586_p0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_590_i1_ in range(default__.abs(407))])), 514, len(_dafny.Map({d_587_p2_: self.f10})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_591_i2_ in range(default__.abs(-939))]))}))

            return lambda17_

        init8_ = lambda16_(p0, p2)
        nw106_ = _dafny.Array(None, 25)
        for i0_8_ in range(nw106_.length(0)):
            nw106_[i0_8_] = init8_(i0_8_)
        d_585_v1_ = nw106_
        d_592_v2_: _dafny.Map
        d_592_v2_ = _dafny.Map({(self).f9: (self).f9})
        d_593_v3_: _dafny.MultiSet
        d_593_v3_ = _dafny.MultiSet([len(d_592_v2_)])
        d_594_v4_: _dafny.Seq
        d_594_v4_ = _dafny.SeqWithoutIsStrInference([p0, (d_593_v3_).cardinality, 333, p0])
        d_595_v5_: _dafny.Map
        d_595_v5_ = _dafny.Map({((d_593_v3_)[(d_594_v4_)[default__.safeIndex(p0, len(d_594_v4_))]] if ((d_594_v4_)[default__.safeIndex(p0, len(d_594_v4_))]) in (d_593_v3_) else p0): True})
        d_596_v6_: _dafny.Seq
        d_596_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isc"))
        d_597_v7_: _dafny.Map
        d_597_v7_ = _dafny.Map({len(d_595_v5_): d_596_v6_})
        d_598_v8_: _dafny.MultiSet
        d_598_v8_ = _dafny.MultiSet([(d_596_v6_ if self.f10 else d_596_v6_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djiib")), d_596_v6_])
        d_599_v9_: _dafny.Map
        d_599_v9_ = _dafny.Map({self.f10: d_597_v7_})
        d_600_v10_: _dafny.Seq
        d_600_v10_ = _dafny.SeqWithoutIsStrInference([self.f10, self.f10])
        d_601_v11_: D2
        d_601_v11_ = D2_DC8(p2, 618, (self).f9)
        pat_let_tv19_ = d_598_v8_
        pat_let_tv20_ = d_598_v8_
        def lambda18_(source11_):
            if source11_.is_DC8:
                d_602___mcc_h0_ = source11_.cf7
                d_603___mcc_h1_ = source11_.cf8
                d_604___mcc_h2_ = source11_.cf9
                d_605_cf9_ = d_604___mcc_h2_
                d_606_cf8_ = d_603___mcc_h1_
                d_607_cf7_ = d_602___mcc_h0_
                return pat_let_tv19_
            elif True:
                d_608___mcc_h3_ = source11_.cf6
                d_609_cf6_ = d_608___mcc_h3_
                return pat_let_tv20_

        rhs100_ = p1
        rhs101_ = (self).f9
        rhs102_ = p2
        rhs103_ = (d_597_v7_) | (((d_599_v9_)[(d_600_v10_)[default__.safeIndex(p0, len(d_600_v10_))]] if ((d_600_v10_)[default__.safeIndex(p0, len(d_600_v10_))]) in (d_599_v9_) else d_597_v7_))
        rhs104_ = lambda18_(d_601_v11_)
        d_585_v1_ = rhs100_
        r0 = rhs101_
        r1 = rhs102_
        d_597_v7_ = rhs103_
        d_598_v8_ = rhs104_
        d_610_v12_: _dafny.Array
        nw107_ = _dafny.Array(int(0), 29)
        d_610_v12_ = nw107_
        d_610_v12_ = (D3_DC10(p0, d_610_v12_, (d_600_v10_)[default__.safeIndex(p0, len(d_600_v10_))])).cf12
        d_611_v13_: _dafny.Seq
        d_611_v13_ = _dafny.SeqWithoutIsStrInference([d_585_v1_, p1, p1, d_585_v1_, d_585_v1_])
        d_612_v14_: C0
        nw108_ = C0()
        nw108_.ctor__(True, len(d_611_v13_), (self.f10) or (p2))
        d_612_v14_ = nw108_
        d_613_v15_: str
        d_613_v15_ = _dafny.CodePoint('k')
        pat_let_tv21_ = d_612_v14_
        pat_let_tv22_ = d_612_v14_
        pat_let_tv23_ = d_612_v14_
        pat_let_tv24_ = d_612_v14_
        def lambda19_(source12_):
            if source12_.is_DC10:
                d_614___mcc_h4_ = source12_.cf11
                d_615___mcc_h5_ = source12_.cf12
                d_616___mcc_h6_ = source12_.cf13
                d_617_cf13_ = d_616___mcc_h6_
                d_618_cf12_ = d_615___mcc_h5_
                d_619_cf11_ = d_614___mcc_h4_
                return True
            elif source12_.is_DC11:
                d_620___mcc_h7_ = source12_.cf14
                d_621___mcc_h8_ = source12_.cf15
                d_622___mcc_h9_ = source12_.cf16
                d_623___mcc_h10_ = source12_.cf17
                d_624_cf17_ = d_623___mcc_h10_
                d_625_cf16_ = d_622___mcc_h9_
                d_626_cf15_ = d_621___mcc_h8_
                d_627_cf14_ = d_620___mcc_h7_
                return (pat_let_tv21_).f13
            elif source12_.is_DC9:
                d_628___mcc_h11_ = source12_.cf10
                d_629_cf10_ = d_628___mcc_h11_
                return (self).f9
            elif True:
                d_630___mcc_h12_ = source12_.cf18
                d_631_cf18_ = d_630___mcc_h12_
                return not((_dafny.Set({pat_let_tv22_.f14})).ispropersubset(_dafny.Set({pat_let_tv23_.f14, (0) - (pat_let_tv24_.f14)})))

        (self).f10 = lambda19_(D3_DC9(_dafny.SeqWithoutIsStrInference([d_613_v15_, d_613_v15_])))
        d_632_v16_: _dafny.Array
        nw109_ = _dafny.Array(None, 15)
        nw109_[int(0)] = p1
        nw109_[int(1)] = p1
        nw109_[int(2)] = d_585_v1_
        nw109_[int(3)] = p1
        nw109_[int(4)] = p1
        nw109_[int(5)] = p1
        nw109_[int(6)] = p1
        nw109_[int(7)] = p1
        nw109_[int(8)] = p1
        nw109_[int(9)] = p1
        nw109_[int(10)] = p1
        nw109_[int(11)] = p1
        nw109_[int(12)] = p1
        nw109_[int(13)] = p1
        nw109_[int(14)] = d_585_v1_
        d_632_v16_ = nw109_
        d_633_v17_: D4
        d_633_v17_ = D4_DC13(d_632_v16_)
        d_634_v18_: _dafny.Array
        nw110_ = _dafny.Array(None, 28)
        nw110_[int(0)] = d_632_v16_
        nw110_[int(1)] = d_632_v16_
        nw110_[int(2)] = d_632_v16_
        nw110_[int(3)] = d_632_v16_
        nw110_[int(4)] = d_632_v16_
        nw110_[int(5)] = d_632_v16_
        nw110_[int(6)] = d_632_v16_
        nw110_[int(7)] = d_632_v16_
        nw110_[int(8)] = d_632_v16_
        nw110_[int(9)] = d_632_v16_
        nw110_[int(10)] = (d_633_v17_).cf19
        nw110_[int(11)] = d_632_v16_
        nw110_[int(12)] = d_632_v16_
        nw110_[int(13)] = d_632_v16_
        nw110_[int(14)] = d_632_v16_
        nw110_[int(15)] = d_632_v16_
        nw110_[int(16)] = d_632_v16_
        nw110_[int(17)] = d_632_v16_
        nw110_[int(18)] = d_632_v16_
        nw110_[int(19)] = d_632_v16_
        nw110_[int(20)] = d_632_v16_
        nw110_[int(21)] = d_632_v16_
        nw110_[int(22)] = d_632_v16_
        nw110_[int(23)] = d_632_v16_
        nw110_[int(24)] = d_632_v16_
        nw110_[int(25)] = d_632_v16_
        nw110_[int(26)] = (d_633_v17_).cf19
        nw110_[int(27)] = d_632_v16_
        d_634_v18_ = nw110_
        index110_ = default__.safeIndex(833, (d_634_v18_).length(0))
        (d_634_v18_)[index110_] = d_632_v16_
        index111_ = default__.safeIndex(833, (d_634_v18_).length(0))
        (d_634_v18_)[index111_] = d_632_v16_
        if (len((d_596_v6_) + (_dafny.SeqWithoutIsStrInference([d_613_v15_ for d_635_i3_ in range(default__.abs(919))])))) < (len(default__.fm10(globalState))):
            (globalState).f7 = d_613_v15_
            d_636_v19_: _dafny.MultiSet
            d_636_v19_ = _dafny.MultiSet([p2])
            d_637_v20_: _dafny.Map
            d_637_v20_ = _dafny.Map({d_612_v14_: d_636_v19_})
            d_638_v21_: _dafny.Map
            d_638_v21_ = _dafny.Map({d_637_v20_: default__.safeModuloInt(66, 853)})
            d_638_v21_ = (d_638_v21_).set(d_637_v20_, default__.safeDivisionInt(d_612_v14_.f14, d_612_v14_.f14))
            d_639_v22_: _dafny.Array
            def lambda20_(d_640_v2_):
                def lambda21_(d_641_i4_):
                    return d_640_v2_

                return lambda21_

            init9_ = lambda20_(d_592_v2_)
            nw111_ = _dafny.Array(None, 12)
            for i0_9_ in range(nw111_.length(0)):
                nw111_[i0_9_] = init9_(i0_9_)
            d_639_v22_ = nw111_
            d_639_v22_ = d_639_v22_
            d_642_v23_: _dafny.Array
            nw112_ = _dafny.Array(_dafny.MultiSet({}), 9)
            d_642_v23_ = nw112_
            d_643_v24_: _dafny.MultiSet
            d_643_v24_ = _dafny.MultiSet([d_585_v1_])
            index112_ = default__.safeIndex(718, (d_642_v23_).length(0))
            (d_642_v23_)[index112_] = (d_643_v24_) - (_dafny.MultiSet([p1]))
            index113_ = default__.safeIndex(718, (d_642_v23_).length(0))
            (d_642_v23_)[index113_] = d_643_v24_
            d_644_v25_: D1
            d_644_v25_ = D1_DC4(d_594_v4_)
            source13_ = d_644_v25_
            if source13_.is_DC5:
                (d_612_v14_).f14 = 421
                d_645_v26_: _dafny.Seq
                d_645_v26_ = _dafny.SeqWithoutIsStrInference([d_592_v2_])
                d_646_v27_: _dafny.Map
                d_646_v27_ = _dafny.Map({d_610_v12_: default__.fm3(d_613_v15_, True, p0, globalState)})
                rhs105_ = d_612_v14_.f14
                rhs106_ = (default__.fm14(-443, globalState)).set(default__.safeIndex((len(_dafny.SeqWithoutIsStrInference([(d_612_v14_).f13]))) - (d_612_v14_.f14), len(default__.fm14(-443, globalState))), d_592_v2_)
                rhs107_ = (d_646_v27_) != ((d_646_v27_) | (d_646_v27_))
                rhs108_ = d_613_v15_
                rhs109_ = p1
                lhs70_ = d_612_v14_
                lhs71_ = globalState
                lhs70_.f14 = rhs105_
                d_645_v26_ = rhs106_
                r0 = rhs107_
                lhs71_.f7 = rhs108_
                d_585_v1_ = rhs109_
                d_647_v28_: C1
                nw113_ = C1()
                nw113_.ctor__((d_612_v14_).f13, (self).fm6((d_612_v14_).f13, (self).f9, globalState), p2)
                d_647_v28_ = nw113_
                d_647_v28_ = d_647_v28_
                d_648_v29_: _dafny.Array
                nw114_ = _dafny.Array(D2.default()(), 29)
                d_648_v29_ = nw114_
                index114_ = default__.safeIndex(523, (d_648_v29_).length(0))
                (d_648_v29_)[index114_] = d_601_v11_
                index115_ = default__.safeIndex(523, (d_648_v29_).length(0))
                (d_648_v29_)[index115_] = D2_DC8(False, (d_647_v28_).f12, (d_612_v14_).f13)
            elif source13_.is_DC6:
                (self).f10 = (not(p2) if self.f10 else (d_612_v14_).f13)
                d_649_v30_: _dafny.Seq
                d_649_v30_ = _dafny.SeqWithoutIsStrInference([d_596_v6_, d_596_v6_])
                d_650_v31_: _dafny.Seq
                d_650_v31_ = _dafny.SeqWithoutIsStrInference([d_598_v8_])
                d_651_v32_: _dafny.Array
                nw115_ = _dafny.Array(None, 9)
                nw115_[int(0)] = d_598_v8_
                nw115_[int(1)] = (d_598_v8_).set(d_596_v6_, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnmnf")))))
                nw115_[int(2)] = d_598_v8_
                nw115_[int(3)] = (_dafny.MultiSet(d_649_v30_)).intersection(d_598_v8_)
                nw115_[int(4)] = (d_598_v8_) | (d_598_v8_)
                nw115_[int(5)] = (d_598_v8_) - (d_598_v8_)
                nw115_[int(6)] = _dafny.MultiSet(default__.fm15((self).f9, globalState))
                nw115_[int(7)] = (_dafny.MultiSet([d_596_v6_])).intersection((d_650_v31_)[default__.safeIndex(p0, len(d_650_v31_))])
                nw115_[int(8)] = (d_598_v8_) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aonb"))]))
                d_651_v32_ = nw115_
                index116_ = default__.safeIndex(482, (d_651_v32_).length(0))
                (d_651_v32_)[index116_] = d_598_v8_
                index117_ = default__.safeIndex(321, (d_585_v1_).length(0))
                (d_585_v1_)[index117_] = ((0) - (d_612_v14_.f14)) > (len(d_595_v5_))
                d_652_v33_: _dafny.Set
                d_652_v33_ = _dafny.Set({d_612_v14_.f14, 141})
                index118_ = default__.safeIndex(482, (d_651_v32_).length(0))
                index119_ = default__.safeIndex(321, (d_585_v1_).length(0))
                rhs110_ = _dafny.MultiSet([(d_652_v33_).isdisjoint(d_652_v33_)])
                rhs111_ = not(p2)
                rhs112_ = p0
                rhs113_ = d_598_v8_
                rhs114_ = (d_612_v14_).f13
                lhs72_ = globalState
                lhs73_ = d_612_v14_
                lhs74_ = d_651_v32_
                lhs75_ = default__.safeIndex(482, (d_651_v32_).length(0))
                lhs76_ = d_585_v1_
                lhs77_ = default__.safeIndex(321, (d_585_v1_).length(0))
                d_636_v19_ = rhs110_
                lhs72_.f0 = rhs111_
                lhs73_.f14 = rhs112_
                lhs74_[lhs75_] = rhs113_
                lhs76_[lhs77_] = rhs114_
                index120_ = default__.safeIndex(597, (d_610_v12_).length(0))
                (d_610_v12_)[index120_] = p0
                d_653_v34_: _dafny.MultiSet
                d_653_v34_ = _dafny.MultiSet([d_593_v3_])
                d_654_v35_: T0
                nw116_ = C1()
                nw116_.ctor__(self.f10, default__.fm3(d_613_v15_, (d_612_v14_).f13, -346, globalState), (d_653_v34_).ispropersubset((d_653_v34_).set(d_593_v3_, default__.abs(-782))))
                d_654_v35_ = nw116_
                d_655_v36_: _dafny.Map
                d_655_v36_ = _dafny.Map({d_636_v19_: d_594_v4_})
                index121_ = default__.safeIndex(597, (d_610_v12_).length(0))
                rhs115_ = len(((d_655_v36_)[_dafny.MultiSet((d_600_v10_ if self.f10 else d_600_v10_))] if (_dafny.MultiSet((d_600_v10_ if self.f10 else d_600_v10_))) in (d_655_v36_) else ((d_644_v25_).cf5) + (d_594_v4_)))
                rhs116_ = ((d_612_v14_).f13) and ((d_585_v1_)[default__.safeIndex(321, (d_585_v1_).length(0))])
                rhs117_ = 243
                rhs118_ = d_654_v35_
                rhs119_ = not((d_612_v14_.f14) < (default__.safeDivisionInt(p0, 858)))
                lhs78_ = d_612_v14_
                lhs79_ = d_610_v12_
                lhs80_ = default__.safeIndex(597, (d_610_v12_).length(0))
                lhs81_ = globalState
                lhs78_.f14 = rhs115_
                r0 = rhs116_
                lhs79_[lhs80_] = rhs117_
                d_654_v35_ = rhs118_
                lhs81_.f0 = rhs119_
                d_656_v37_: _dafny.Seq
                d_656_v37_ = _dafny.SeqWithoutIsStrInference([d_594_v4_, d_594_v4_])
                d_657_v38_: _dafny.Map
                d_657_v38_ = _dafny.Map({not (self.f10) or (False): ((d_656_v37_).set(default__.safeIndex(p0, len(d_656_v37_)), d_594_v4_)) + (((d_656_v37_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), len(d_656_v37_)), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqssgr"))) for d_658_i5_ in range(default__.abs(184))]))).set(default__.safeIndex(len(d_594_v4_), len((d_656_v37_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), len(d_656_v37_)), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqssgr"))) for d_659_i5_ in range(default__.abs(184))])))), d_594_v4_))})
                d_657_v38_ = (d_657_v38_).set(not(((d_585_v1_)[default__.safeIndex(321, (d_585_v1_).length(0))]) and ((d_654_v35_).f9)), d_656_v37_)
            elif True:
                d_660___mcc_h13_ = source13_.cf5
                d_661_cf5_ = d_660___mcc_h13_
                d_662_v39_: _dafny.Map
                d_662_v39_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([(d_612_v14_).f13])})
                (d_612_v14_).f14 = default__.safeModuloInt((d_612_v14_.f14) - (p0), (self).fm4(d_662_v39_, (d_612_v14_).f13, self.f10, d_612_v14_.f14, globalState))
                d_585_v1_ = d_585_v1_
                d_663_v40_: _dafny.Set
                d_663_v40_ = _dafny.Set({p1, d_585_v1_})
                d_663_v40_ = d_663_v40_
                (d_612_v14_).f14 = (0) - (p0)
        elif True:
            d_595_v5_ = (d_595_v5_).set(p0, (d_612_v14_).f13)
            d_664_v41_: _dafny.Seq
            d_664_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({default__.fm2(globalState): (d_612_v14_).f13})])
            (globalState).f0 = (True if not(False) else (_dafny.SeqWithoutIsStrInference([_dafny.Map({not((d_612_v14_).f13): (d_600_v10_)[default__.safeIndex(p0, len(d_600_v10_))]}), d_592_v2_, d_592_v2_])) == (d_664_v41_))
            d_665_v42_: C0
            nw117_ = C0()
            nw117_.ctor__((d_612_v14_).f13, p0, True)
            d_665_v42_ = nw117_
            d_666_v43_: _dafny.Map
            d_666_v43_ = _dafny.Map({d_612_v14_.f14: d_594_v4_})
            d_667_v44_: _dafny.MultiSet
            d_667_v44_ = _dafny.MultiSet([self.f10, (d_612_v14_).f13, self.f10, not(self.f10)])
            d_666_v43_ = (d_666_v43_).set((((d_667_v44_)[(d_612_v14_).f13] if ((d_612_v14_).f13) in (d_667_v44_) else (0) - (d_612_v14_.f14))) * (default__.fm3(_dafny.CodePoint('e'), False, d_612_v14_.f14, globalState)), (default__.fm13((self).f9, (d_665_v42_).f13, (d_612_v14_).f13, len(d_596_v6_), globalState)) + (_dafny.SeqWithoutIsStrInference([d_665_v42_.f14 for d_668_i6_ in range(default__.abs(-594))])))
            (globalState).f7 = d_613_v15_
        r0 = (d_612_v14_).f13
        d_669_v45_: D0
        d_669_v45_ = D0_DC0((0) - (p0))
        pat_let_tv25_ = d_600_v10_
        pat_let_tv26_ = d_612_v14_
        pat_let_tv27_ = d_600_v10_
        pat_let_tv28_ = d_612_v14_
        pat_let_tv29_ = p2
        def lambda22_(source14_):
            if source14_.is_DC1:
                d_670___mcc_h14_ = source14_.cf1
                d_671_cf1_ = d_670___mcc_h14_
                return (pat_let_tv25_)[default__.safeIndex(pat_let_tv26_.f14, len(pat_let_tv27_))]
            elif source14_.is_DC2:
                d_672___mcc_h15_ = source14_.cf2
                d_673___mcc_h16_ = source14_.cf3
                d_674___mcc_h17_ = source14_.cf4
                d_675_cf4_ = d_674___mcc_h17_
                d_676_cf3_ = d_673___mcc_h16_
                d_677_cf2_ = d_672___mcc_h15_
                return (pat_let_tv28_.f14) >= (d_676_cf3_)
            elif source14_.is_DC3:
                return self.f10
            elif True:
                d_678___mcc_h18_ = source14_.cf0
                d_679_cf0_ = d_678___mcc_h18_
                return pat_let_tv29_

        r1 = lambda22_(d_669_v45_)
        d_680_v46_: _dafny.Map
        d_680_v46_ = _dafny.Map({p0: d_610_v12_})
        r2 = (d_680_v46_).set((0) - (d_612_v14_.f14), d_610_v12_)
        return r0, r1, r2


class C3(T0):
    def  __init__(self):
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f9(self):
        return self._f9
    def ctor__(self, f9):
        (self)._f9 = f9

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouhfhsh"))), -999])).cardinality)) - ((636) * (-559))

    def fm5(self, globalState):
        return _dafny.CodePoint('a')

    def m1(self, p0, globalState):
        r0: int = int(0)
        r1: D0 = D0.default()()
        r2: int = int(0)
        d_681_v0_: int
        d_681_v0_ = -2
        d_682_v1_: _dafny.Map
        d_682_v1_ = _dafny.Map({(self).f9: (d_681_v0_) - (757)})
        d_683_v2_: _dafny.Array
        def lambda23_(d_684_v0_):
            def lambda24_(d_685_i0_):
                return (d_685_i0_) - (d_684_v0_)

            return lambda24_

        init10_ = lambda23_(d_681_v0_)
        nw118_ = _dafny.Array(None, 20)
        for i0_10_ in range(nw118_.length(0)):
            nw118_[i0_10_] = init10_(i0_10_)
        d_683_v2_ = nw118_
        d_686_v3_: _dafny.Set
        d_686_v3_ = _dafny.Set({d_683_v2_, d_683_v2_, d_683_v2_, d_683_v2_})
        d_682_v1_ = (d_682_v1_).set((d_686_v3_).ispropersubset(_dafny.Set({d_683_v2_, d_683_v2_, d_683_v2_})), d_681_v0_)
        d_687_v4_: _dafny.Seq
        d_687_v4_ = _dafny.SeqWithoutIsStrInference([d_681_v0_, d_681_v0_, d_681_v0_, d_681_v0_, d_681_v0_])
        d_688_v5_: _dafny.Seq
        d_688_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ps"))
        d_689_v6_: _dafny.MultiSet
        d_689_v6_ = _dafny.MultiSet([(0) - (d_681_v0_), d_681_v0_])
        d_690_v7_: _dafny.MultiSet
        d_690_v7_ = _dafny.MultiSet([_dafny.MultiSet([d_681_v0_]), _dafny.MultiSet([len(d_688_v5_), d_681_v0_]), d_689_v6_])
        d_691_v8_: _dafny.Map
        d_691_v8_ = _dafny.Map({d_687_v4_: (D0_DC2((self).f9, d_681_v0_, d_690_v7_)).cf3})
        d_692_v9_: _dafny.MultiSet
        d_692_v9_ = _dafny.MultiSet([len(d_691_v8_)])
        d_693_v10_: _dafny.Array
        nw119_ = _dafny.Array(None, 6)
        nw119_[int(0)] = d_689_v6_
        nw119_[int(1)] = d_692_v9_
        nw119_[int(2)] = _dafny.MultiSet([d_681_v0_])
        nw119_[int(3)] = d_689_v6_
        nw119_[int(4)] = _dafny.MultiSet(d_687_v4_)
        nw119_[int(5)] = (d_692_v9_) - (d_689_v6_)
        d_693_v10_ = nw119_
        index122_ = default__.safeIndex(96, (d_693_v10_).length(0))
        (d_693_v10_)[index122_] = (_dafny.MultiSet([(0) - (d_681_v0_), d_681_v0_]) if (self).f9 else d_689_v6_)
        d_694_v11_: _dafny.Map
        d_694_v11_ = _dafny.Map({d_681_v0_: d_692_v9_})
        index123_ = default__.safeIndex(96, (d_693_v10_).length(0))
        (d_693_v10_)[index123_] = (d_692_v9_).intersection(((d_694_v11_)[d_681_v0_] if (d_681_v0_) in (d_694_v11_) else d_692_v9_))
        d_695_i1_: int
        d_695_i1_ = 0
        with _dafny.label("2"):
            while (self).f9:
                with _dafny.c_label("2"):
                    if (d_695_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_695_i1_ = (d_695_i1_) + (1)
                    (globalState).f3 = default__.fm2(globalState)
                    d_696_v12_: C1
                    nw120_ = C1()
                    nw120_.ctor__((self).f9, d_681_v0_, (self).f9)
                    d_696_v12_ = nw120_
                    d_697_v13_: str
                    d_697_v13_ = _dafny.CodePoint('g')
                    d_698_v14_: T0
                    nw121_ = C0()
                    nw121_.ctor__((self).f9, d_681_v0_, (self).f9)
                    d_698_v14_ = nw121_
                    d_699_v15_: _dafny.Map
                    d_699_v15_ = _dafny.Map({d_698_v14_: (d_696_v12_).f11})
                    d_700_v16_: C1
                    nw122_ = C1()
                    nw122_.ctor__((default__.fm3(d_697_v13_, (self).f9, (d_696_v12_).f12, globalState)) in (_dafny.MultiSet([d_681_v0_])), (_dafny.MultiSet([d_681_v0_, (d_696_v12_).f12])).cardinality, ((d_699_v15_)[d_698_v14_] if (d_698_v14_) in (d_699_v15_) else (self).f9))
                    d_700_v16_ = nw122_
                    d_701_v17_: _dafny.Map
                    d_701_v17_ = _dafny.Map({(d_696_v12_).f12: d_696_v12_})
                    d_702_v18_: _dafny.Seq
                    d_702_v18_ = _dafny.SeqWithoutIsStrInference([d_696_v12_, d_700_v16_, d_696_v12_])
                    d_700_v16_ = ((d_701_v17_)[(d_700_v16_).f12] if ((d_700_v16_).f12) in (d_701_v17_) else (d_702_v18_)[default__.safeIndex((d_696_v12_).f12, len(d_702_v18_))])
                    d_703_v19_: _dafny.MultiSet
                    d_703_v19_ = _dafny.MultiSet([not((d_696_v12_).f11), (self).f9, (self).f9, (d_698_v14_).f9])
                    d_704_v20_: C2
                    nw123_ = C2()
                    nw123_.ctor__((d_703_v19_).isdisjoint(d_703_v19_), (d_697_v13_) not in (d_688_v5_))
                    d_704_v20_ = nw123_
                    pass
            pass
        hi10_ = d_681_v0_
        for d_705_i2_ in range(d_681_v0_, hi10_):
            r0 = len((d_688_v5_) + (d_688_v5_))
            d_706_v21_: _dafny.Map
            d_706_v21_ = _dafny.Map({(self).f9: d_688_v5_})
            d_706_v21_ = (d_706_v21_).set(not (default__.fm2(globalState)) or ((self).f9), (d_688_v5_) + (d_688_v5_))
            d_707_v22_: _dafny.Map
            d_707_v22_ = _dafny.Map({d_688_v5_: not(False)})
            d_708_v23_: _dafny.Set
            d_708_v23_ = _dafny.Set({(self).f9, (self).f9, (self).f9, (self).f9, (self).f9})
            d_709_v24_: _dafny.Map
            d_709_v24_ = _dafny.Map({d_707_v22_: (d_708_v23_).intersection(_dafny.Set({(self).f9, (self).f9, (self).f9, (self).f9, not((self).f9)}))})
            d_709_v24_ = (d_709_v24_).set(_dafny.Map({d_688_v5_: (self).f9}), (d_708_v23_) | (d_708_v23_))
            (globalState).f3 = (self).f9
        d_710_v25_: str
        d_710_v25_ = _dafny.CodePoint('t')
        d_711_v26_: _dafny.Map
        d_711_v26_ = _dafny.Map({(self).f9: d_710_v25_})
        d_712_v27_: _dafny.Map
        d_712_v27_ = _dafny.Map({d_681_v0_: len(d_711_v26_)})
        d_712_v27_ = (d_712_v27_).set(d_681_v0_, default__.safeDivisionInt((d_689_v6_).cardinality, d_681_v0_))
        if (self).f9:
            d_713_v28_: _dafny.Array
            nw124_ = _dafny.Array(None, 3)
            nw124_[int(0)] = d_683_v2_
            nw124_[int(1)] = d_683_v2_
            nw124_[int(2)] = d_683_v2_
            d_713_v28_ = nw124_
            d_713_v28_ = d_713_v28_
            d_714_v29_: _dafny.Seq
            d_714_v29_ = _dafny.SeqWithoutIsStrInference([d_682_v1_])
            d_715_v30_: D2
            d_715_v30_ = D2_DC8((self).f9, (d_681_v0_ if (self).f9 else len(d_712_v27_)), (d_714_v29_) < (d_714_v29_))
            source15_ = d_715_v30_
            if source15_.is_DC8:
                d_716___mcc_h0_ = source15_.cf7
                d_717___mcc_h1_ = source15_.cf8
                d_718___mcc_h2_ = source15_.cf9
                d_719_cf9_ = d_718___mcc_h2_
                d_720_cf8_ = d_717___mcc_h1_
                d_721_cf7_ = d_716___mcc_h0_
                (globalState).f3 = (self).f9
                d_722_v31_: _dafny.MultiSet
                d_722_v31_ = _dafny.MultiSet([not(True), d_719_cf9_])
                (globalState).f3 = (default__.safeDivisionInt(d_681_v0_, (d_722_v31_).cardinality)) <= ((0) - ((d_720_cf8_) * (d_681_v0_)))
                (globalState).f0 = (self).f9
                d_723_v32_: _dafny.Map
                d_723_v32_ = _dafny.Map({d_710_v25_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uathenfi"))})
                d_724_v33_: int
                out19_: int
                out19_ = default__.m0((d_720_cf8_ if False else len(d_687_v4_)), (d_683_v2_ if d_719_cf9_ else d_683_v2_), d_721_cf7_, default__.fm16(d_723_v32_, d_720_cf8_, globalState), globalState)
                d_724_v33_ = out19_
            elif True:
                d_725___mcc_h3_ = source15_.cf6
                d_726_cf6_ = d_725___mcc_h3_
                d_727_v34_: _dafny.Map
                d_727_v34_ = _dafny.Map({(0) - (d_681_v0_): (self).f9})
                d_728_v35_: int
                out20_: int
                out20_ = default__.m0(default__.fm3(d_710_v25_, (self).f9, d_681_v0_, globalState), d_683_v2_, False, d_727_v34_, globalState)
                d_728_v35_ = out20_
                d_729_v36_: T0
                nw125_ = C0()
                nw125_.ctor__((d_710_v25_) not in (d_688_v5_), d_681_v0_, ((self).f9) or (not(False)))
                d_729_v36_ = nw125_
                d_730_v37_: _dafny.Map
                d_730_v37_ = _dafny.Map({d_688_v5_: d_681_v0_})
                d_731_v38_: _dafny.Map
                d_731_v38_ = _dafny.Map({len(_dafny.Map({970: d_683_v2_})): d_688_v5_})
                rhs120_ = d_729_v36_
                rhs121_ = (d_731_v38_) != (d_731_v38_)
                rhs122_ = len(d_687_v4_)
                rhs123_ = d_730_v37_
                lhs82_ = globalState
                d_729_v36_ = rhs120_
                lhs82_.f3 = rhs121_
                r2 = rhs122_
                d_730_v37_ = rhs123_
                r2 = default__.safeModuloInt(d_681_v0_, (len(d_687_v4_)) + (d_681_v0_))
                d_688_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmguhhp"))
            d_732_v39_: C0
            nw126_ = C0()
            nw126_.ctor__((d_681_v0_) < (-788), d_681_v0_, (self).f9)
            d_732_v39_ = nw126_
            d_732_v39_ = d_732_v39_
            d_733_v40_: _dafny.Map
            d_733_v40_ = _dafny.Map({False: (self).f9})
            (globalState).f3 = ((d_733_v40_)[not((d_732_v39_).f13)] if (not((d_732_v39_).f13)) in (d_733_v40_) else False)
            d_734_v41_: _dafny.Map
            d_734_v41_ = _dafny.Map({((d_693_v10_)[default__.safeIndex(96, (d_693_v10_).length(0))]).cardinality: ((self).fm5(globalState) if (d_732_v39_).f13 else d_710_v25_)})
            d_734_v41_ = (d_734_v41_).set(d_681_v0_, d_710_v25_)
        elif True:
            d_735_v42_: _dafny.Map
            d_735_v42_ = _dafny.Map({d_710_v25_: default__.safeModuloInt(d_681_v0_, d_681_v0_)})
            d_735_v42_ = d_735_v42_
            d_688_v5_ = d_688_v5_
            d_681_v0_ = (150) * (d_681_v0_)
            r2 = d_681_v0_
            d_736_v43_: _dafny.Array
            nw127_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_736_v43_ = nw127_
            d_737_v44_: _dafny.Array
            def lambda25_(d_738_v25_):
                def lambda26_(d_739_i3_):
                    return d_738_v25_

                return lambda26_

            init11_ = lambda25_(d_710_v25_)
            nw128_ = _dafny.Array(None, 8)
            for i0_11_ in range(nw128_.length(0)):
                nw128_[i0_11_] = init11_(i0_11_)
            d_737_v44_ = nw128_
            index124_ = default__.safeIndex(626, (d_736_v43_).length(0))
            (d_736_v43_)[index124_] = d_737_v44_
            index125_ = default__.safeIndex(626, (d_736_v43_).length(0))
            (d_736_v43_)[index125_] = d_737_v44_
        r0 = ((0) - (d_681_v0_)) + (default__.safeDivisionInt(d_681_v0_, d_681_v0_))
        d_740_v45_: D1
        d_740_v45_ = D1_DC5()
        pat_let_tv30_ = d_681_v0_
        pat_let_tv31_ = d_681_v0_
        pat_let_tv32_ = d_712_v27_
        pat_let_tv33_ = d_681_v0_
        pat_let_tv34_ = d_681_v0_
        pat_let_tv35_ = d_712_v27_
        pat_let_tv36_ = d_681_v0_
        def lambda27_(source16_):
            if source16_.is_DC5:
                return D0_DC0(pat_let_tv30_)
            elif source16_.is_DC6:
                return D0_DC0(pat_let_tv31_)
            elif True:
                d_741___mcc_h4_ = source16_.cf5
                d_742_cf5_ = d_741___mcc_h4_
                return D0_DC0(len(_dafny.Set({((pat_let_tv32_)[pat_let_tv33_] if (pat_let_tv34_) in (pat_let_tv35_) else pat_let_tv36_)})))

        r1 = lambda27_(d_740_v45_)
        d_743_v46_: _dafny.Set
        d_743_v46_ = _dafny.Set({(self).f9, True, True})
        r2 = len((d_743_v46_).intersection(d_743_v46_))
        return r0, r1, r2

