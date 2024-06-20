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
        if True:
            return default__.safeModuloInt(-138, len(_dafny.Map({False: False})))
        elif True:
            return 541

    @staticmethod
    def fm1(p0, p1, globalState):
        return True

    @staticmethod
    def fm2(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ko"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_0_i0_ in range(default__.abs(296))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1_i1_ in range(default__.abs(78))]))

    @staticmethod
    def fm3(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s')])).Elements:
                d_2_v0_: str = compr_0_
                if (d_2_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s')])):
                    coll0_[d_2_v0_] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ac"))])) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dubvfe")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_3_i0_ in range(default__.abs(152))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pt")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))]))
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True, False, False, True, False]))))

    @staticmethod
    def fm5(p0, globalState):
        return _dafny.Map({len(_dafny.Set({_dafny.Set({not(not(False))}), _dafny.Set({True, False, True, True, (D9_DC25(True, _dafny.Set({_dafny.CodePoint('d')}))).cf41}), _dafny.Set({False}), _dafny.Set({True})})): (_dafny.MultiSet([not(True), False])) != (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))})

    @staticmethod
    def fm6(p0, globalState):
        return _dafny.CodePoint('b')

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({False: not(True)}))) | ((_dafny.Map({False: True})) | (_dafny.Map({not(True): not(False)})))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return _dafny.Map({(_dafny.MultiSet([False, False, False, False, True])).intersection(_dafny.MultiSet([False, False])): _dafny.CodePoint('l')})

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        if False:
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(228, 940):
                    d_5_v0_: int = compr_1_
                    if ((228) <= (d_5_v0_)) and ((d_5_v0_) < (940)):
                        coll1_[default__.safeModuloInt(d_5_v0_, 160)] = 791
                return _dafny.Map(coll1_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_4_i0_ in range(default__.abs(671))])): _dafny.Map({273: (_dafny.MultiSet([not(False), False])).cardinality})})) | (_dafny.Map({-33: iife1_()
            }))
        elif True:
            return _dafny.Map({(0) - (-777): _dafny.Map({len(_dafny.Map({False: False})): -658})})

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return (_dafny.MultiSet([_dafny.Map({False: 459}), _dafny.Map({False: -480})])) | (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.Set({True}))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 207})]))))

    @staticmethod
    def fm15(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])]) if True else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        if (988) != (len(_dafny.Map({True: False}))):
            return _dafny.Map({478: 394})
        elif True:
            return (_dafny.Map({105: 451})) | (_dafny.Map({251: -748}))

    @staticmethod
    def fm17(p0, p1, globalState):
        if (True) or (False):
            return _dafny.MultiSet([_dafny.CodePoint('s')])
        elif True:
            return _dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('o'), _dafny.CodePoint('i')])

    @staticmethod
    def fm19(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: D0
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([D0_DC2(D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_6_i0_ in range(default__.abs(633))])))])).Elements:
                d_7_v0_: D0 = compr_2_
                if (d_7_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC2(D0_DC0(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_6_i0_ in range(default__.abs(633))])))])):
                    coll2_ = coll2_.union(_dafny.Set([d_7_v0_]))
            return _dafny.Set(coll2_)
        return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('q'), _dafny.CodePoint('y')])).cardinality])) + ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwrvsluv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtfd"))}))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(iife2_()
        )]))).cardinality, len(_dafny.Set({False}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrcmuol")))])))

    @staticmethod
    def fm21(p0, globalState):
        return (_dafny.Set({175, 652, -976, 896})) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))}))

    @staticmethod
    def fm22(p0, globalState):
        return D3_DC12((_dafny.SeqWithoutIsStrInference([not(True), False])) + (_dafny.SeqWithoutIsStrInference([not(True)])))

    @staticmethod
    def fm23(globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(673, 245):
                d_9_v0_: int = compr_3_
                if ((673) <= (d_9_v0_)) and ((d_9_v0_) < (245)):
                    coll3_ = coll3_.union(_dafny.Set([(d_9_v0_) - (937)]))
            return _dafny.Set(coll3_)
        return D1_DC5((0) - ((182) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(not(False))]) for d_8_i0_ in range(default__.abs(357))])))), (-715) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lucmkk")))), -424, len((iife3_()
) | (_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.Map({False: False})}))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_10_i1_ in range(default__.abs(329))]))])).cardinality}))))

    @staticmethod
    def fm24(p0, p1, globalState):
        return D9_DC26((0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yglvl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emwcyv"))))))

    @staticmethod
    def fm25(p0, p1, globalState):
        return (_dafny.Map({True: (0) - (-269)})) | (_dafny.Map({True: 773}))

    @staticmethod
    def fm26(p0, globalState):
        return D0_DC2(D0_DC1(873, not(True), len(_dafny.SeqWithoutIsStrInference([380]))))

    @staticmethod
    def fm27(globalState):
        if not((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), False, not(True)]))).ispropersubset(_dafny.MultiSet([False, False]))):
            return _dafny.Map({not(False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivfd"))})
        elif True:
            return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbqnyc"))})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjxglfc"))}))

    @staticmethod
    def fm28(p0, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])})).Elements:
                d_11_v0_: _dafny.Seq = compr_4_
                if (d_11_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])})):
                    coll4_[d_11_v0_] = 200
            return _dafny.Map(coll4_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): 868})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([True, False]): (0) - (len(_dafny.SeqWithoutIsStrInference([False, False])))})) | (iife4_()
        ))

    @staticmethod
    def fm29(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False, True]) for d_12_i0_ in range(default__.abs(171))])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(False)]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False])), _dafny.MultiSet([True]), _dafny.MultiSet([True]), _dafny.MultiSet([True, False])]))

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm31(p0, globalState):
        if not(not(True)):
            return D12_DC36(True, False, 125, 710, _dafny.Map({True: _dafny.SeqWithoutIsStrInference([-438 for d_13_i0_ in range(default__.abs(674))])}))
        elif not(True):
            return D12_DC36(False, True, 770, len(_dafny.SeqWithoutIsStrInference([988])), _dafny.Map({False: _dafny.SeqWithoutIsStrInference([735, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), -224])}))
        elif True:
            return D12_DC36(False, not(False), 977, len(_dafny.SeqWithoutIsStrInference([False, False])), _dafny.Map({True: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])), -67, len(_dafny.Map({not(not(True)): (0) - (len(_dafny.SeqWithoutIsStrInference([True])))}))]))])}))

    @staticmethod
    def fm32(globalState):
        return ((_dafny.SeqWithoutIsStrInference([D11_DC31(-732)])) + (_dafny.SeqWithoutIsStrInference([D11_DC31(582) for d_14_i0_ in range(default__.abs(681))]))) + (_dafny.SeqWithoutIsStrInference([D11_DC31((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D0_DC1((0) - (len(_dafny.Set({True, False}))), True, 333)).cf2, False]))).cardinality), D11_DC31(len(_dafny.Set({not(True)}))), D11_DC31(len(_dafny.SeqWithoutIsStrInference([True]))), D11_DC31(len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')]))]))])))]))

    @staticmethod
    def fm33(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([D11_DC31(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to"))))]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bffhds")))})).keys.Elements:
                d_16_v0_: _dafny.Seq = compr_5_
                if (d_16_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([D11_DC31(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "to"))))]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bffhds")))})):
                    coll5_[d_16_v0_] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: -691}))])
            return _dafny.Map(coll5_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([D11_DC31(807) for d_15_i0_ in range(default__.abs(-70))]): _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aguorpcjs"))))])})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([D11_DC31(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "habj"))))]): _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nv")))), 200])}))) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([D11_DC31(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdusdcx"))))]): _dafny.SeqWithoutIsStrInference([582])})) | (iife5_()
        ))

    @staticmethod
    def fm34(p0, globalState):
        return D0_DC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlwjdjdmf")))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        return ((_dafny.Set({not(not(True)), False})) | (_dafny.Set({True}))) | (_dafny.Set({False, True}))

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        r2 = p2
        d_17_v0_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 20)
        d_17_v0_ = nw0_
        d_18_v1_: bool
        d_18_v1_ = False
        d_19_v2_: _dafny.Map
        d_19_v2_ = _dafny.Map({d_17_v0_: d_18_v1_})
        d_20_v3_: _dafny.Seq
        d_20_v3_ = _dafny.SeqWithoutIsStrInference([p2])
        d_21_v4_: D8
        d_21_v4_ = D8_DC20(True, (d_20_v3_)[default__.safeIndex(p2, len(d_20_v3_))])
        d_22_v5_: C5
        nw1_ = C5()
        nw1_.ctor__(len(d_19_v2_), default__.safeModuloInt(p2, p2), (d_21_v4_).cf34)
        d_22_v5_ = nw1_
        d_23_v6_: _dafny.Seq
        d_23_v6_ = _dafny.SeqWithoutIsStrInference([p1])
        d_24_v7_: str
        d_24_v7_ = _dafny.CodePoint('x')
        d_25_v8_: C4
        nw2_ = C4()
        nw2_.ctor__(d_17_v0_, (d_22_v5_).f25)
        d_25_v8_ = nw2_
        d_26_v9_: _dafny.Seq
        d_26_v9_ = _dafny.SeqWithoutIsStrInference([d_25_v8_, d_25_v8_])
        d_27_v10_: _dafny.MultiSet
        d_27_v10_ = _dafny.MultiSet([(d_26_v9_)[default__.safeIndex((d_22_v5_).f25, len(d_26_v9_))]])
        d_28_v11_: _dafny.Seq
        d_28_v11_ = _dafny.SeqWithoutIsStrInference([default__.fm1(895, d_18_v1_, globalState)])
        d_29_v12_: _dafny.Set
        d_29_v12_ = _dafny.Set({p2})
        d_30_v13_: _dafny.Map
        d_30_v13_ = _dafny.Map({(d_22_v5_).f25: (d_22_v5_).f26})
        rhs0_ = (len((p1).set(default__.safeIndex((D8_DC21(d_18_v1_, d_18_v1_, (0) - (len((d_23_v6_)[default__.safeIndex(p2, len(d_23_v6_))])))).cf37, len(p1)), d_24_v7_))) < (default__.safeModuloInt((0) - ((0) - ((d_27_v10_).cardinality)), len(d_28_v11_)))
        rhs1_ = (d_22_v5_).f25
        rhs2_ = (d_29_v12_).issubset(d_29_v12_)
        rhs3_ = default__.safeModuloInt((0) - (len(d_30_v13_)), p2)
        lhs0_ = globalState
        lhs1_ = globalState
        lhs0_.f7 = rhs0_
        r2 = rhs1_
        d_18_v1_ = rhs2_
        lhs1_.f8 = rhs3_
        d_31_v14_: _dafny.Map
        d_31_v14_ = _dafny.Map({d_18_v1_: d_18_v1_})
        d_32_v15_: C0
        nw3_ = C0()
        nw3_.ctor__()
        d_32_v15_ = nw3_
        d_33_v16_: _dafny.Map
        d_33_v16_ = _dafny.Map({d_32_v15_: 797})
        d_34_v17_: _dafny.Map
        d_34_v17_ = _dafny.Map({d_33_v16_: d_31_v14_})
        d_35_v18_: _dafny.Array
        nw4_ = _dafny.Array(None, 7)
        nw4_[int(0)] = d_31_v14_
        nw4_[int(1)] = d_31_v14_
        nw4_[int(2)] = _dafny.Map({not(d_18_v1_): d_18_v1_})
        nw4_[int(3)] = _dafny.Map({not(d_18_v1_): d_18_v1_})
        nw4_[int(4)] = ((d_34_v17_)[_dafny.Map({d_32_v15_: (d_22_v5_).f26})] if (_dafny.Map({d_32_v15_: (d_22_v5_).f26})) in (d_34_v17_) else d_31_v14_)
        nw4_[int(5)] = d_31_v14_
        nw4_[int(6)] = d_31_v14_
        d_35_v18_ = nw4_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_35_v18_).length(0)):
            d_36_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_36_i0_)) and ((d_36_i0_) < ((d_35_v18_).length(0)))):
                (d_35_v18_)[(d_36_i0_)] = (_dafny.Map({False: d_18_v1_})) | ((D19_DC52(d_31_v14_)).cf85)
        hi0_ = p2
        for d_37_i1_ in range((d_22_v5_).f25, hi0_):
            d_38_v19_: _dafny.Array
            nw5_ = _dafny.Array(False, 16)
            d_38_v19_ = nw5_
            d_38_v19_ = d_38_v19_
            r2 = default__.fm0((d_22_v5_).f25, d_18_v1_, globalState)
            (globalState).f7 = d_18_v1_
            (globalState).f6 = p1
        d_18_v1_ = d_18_v1_
        d_39_v20_: C2
        nw6_ = C2()
        nw6_.ctor__(d_30_v13_, d_18_v1_, d_18_v1_, d_20_v3_, p2)
        d_39_v20_ = nw6_
        d_40_v21_: _dafny.Array
        nw7_ = _dafny.Array(None, 1)
        nw7_[int(0)] = d_39_v20_
        d_40_v21_ = nw7_
        d_41_v22_: _dafny.MultiSet
        d_41_v22_ = _dafny.MultiSet([d_40_v21_, d_40_v21_, d_40_v21_, d_40_v21_])
        d_42_v23_: D20
        d_42_v23_ = D20_DC55(d_41_v22_)
        d_43_v24_: D0
        d_43_v24_ = D0_DC1(default__.fm0(p2, False, globalState), (d_41_v22_).issubset((d_42_v23_).cf89), (d_22_v5_).f26)
        r0 = d_43_v24_
        d_44_v25_: _dafny.Map
        d_44_v25_ = _dafny.Map({d_18_v1_: d_20_v3_})
        d_45_v26_: D12
        d_45_v26_ = D12_DC36(not((d_39_v20_).f32), d_18_v1_, len(p1), p2, d_44_v25_)
        pat_let_tv0_ = d_39_v20_
        pat_let_tv1_ = d_22_v5_
        pat_let_tv2_ = d_39_v20_
        pat_let_tv3_ = d_24_v7_
        pat_let_tv4_ = d_24_v7_
        pat_let_tv5_ = d_39_v20_
        pat_let_tv6_ = d_22_v5_
        pat_let_tv7_ = p1
        pat_let_tv8_ = d_18_v1_
        pat_let_tv9_ = d_22_v5_
        pat_let_tv10_ = d_22_v5_
        pat_let_tv11_ = d_18_v1_
        def lambda0_(source0_):
            if source0_.is_DC35:
                return (_dafny.Map({(0) - (len((pat_let_tv0_).f31)): True})).set((pat_let_tv1_).f26, (pat_let_tv2_).f32)
            elif source0_.is_DC36:
                d_46___mcc_h0_ = source0_.cf56
                d_47___mcc_h1_ = source0_.cf57
                d_48___mcc_h2_ = source0_.cf58
                d_49___mcc_h3_ = source0_.cf59
                d_50___mcc_h4_ = source0_.cf60
                d_51_cf60_ = d_50___mcc_h4_
                d_52_cf59_ = d_49___mcc_h3_
                d_53_cf58_ = d_48___mcc_h2_
                d_54_cf57_ = d_47___mcc_h1_
                d_55_cf56_ = d_46___mcc_h0_
                def iife6_():
                    coll6_ = _dafny.Map()
                    compr_6_: int
                    for compr_6_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([pat_let_tv3_ for d_56_i2_ in range(default__.abs(842))])), len(_dafny.SeqWithoutIsStrInference([-950]))])).Elements:
                        d_57_v27_: int = compr_6_
                        if (d_57_v27_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([pat_let_tv4_ for d_56_i2_ in range(default__.abs(842))])), len(_dafny.SeqWithoutIsStrInference([-950]))])):
                            coll6_[(d_57_v27_) * ((pat_let_tv6_).f26)] = (pat_let_tv5_).f32
                    return _dafny.Map(coll6_)
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(-214, 460):
                        d_58_v28_: int = compr_7_
                        if ((-214) <= (d_58_v28_)) and ((d_58_v28_) < (460)):
                            coll7_[default__.safeDivisionInt(d_58_v28_, d_53_cf58_)] = d_55_cf56_
                    return _dafny.Map(coll7_)
                return (iife6_()
                ) | ((iife7_()
                ).set(len(_dafny.Map({True: len(pat_let_tv7_)})), d_54_cf57_))
            elif source0_.is_DC34:
                d_59___mcc_h5_ = source0_.cf55
                d_60_cf55_ = d_59___mcc_h5_
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(989, 320):
                        d_61_v29_: int = compr_8_
                        if ((989) <= (d_61_v29_)) and ((d_61_v29_) < (320)):
                            coll8_[default__.safeModuloInt(d_61_v29_, (pat_let_tv9_).f25)] = pat_let_tv8_
                    return _dafny.Map(coll8_)
                return iife8_()
                
            elif True:
                d_62___mcc_h6_ = source0_.cf61
                d_63_cf61_ = d_62___mcc_h6_
                return _dafny.Map({(pat_let_tv10_).f25: pat_let_tv11_})

        r1 = lambda0_(d_45_v26_)
        r2 = (d_22_v5_).f25
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_64_v0_: _dafny.Seq
        d_64_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hg"))
        d_65_v1_: D0
        d_65_v1_ = D0_DC0(d_64_v0_)
        d_66_v2_: str
        d_66_v2_ = _dafny.CodePoint('t')
        d_67_v3_: _dafny.Array
        nw8_ = _dafny.Array(None, 12)
        nw8_[int(0)] = d_64_v0_
        nw8_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mexrhnq"))
        nw8_[int(2)] = d_64_v0_
        nw8_[int(3)] = d_64_v0_
        nw8_[int(4)] = d_64_v0_
        nw8_[int(5)] = (d_65_v1_).cf0
        nw8_[int(6)] = d_64_v0_
        nw8_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qu"))
        nw8_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wggsbwf"))
        nw8_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnjar"))
        nw8_[int(10)] = d_64_v0_
        nw8_[int(11)] = (d_64_v0_).set(default__.safeIndex(-282, len(d_64_v0_)), d_66_v2_)
        d_67_v3_ = nw8_
        d_68_v4_: bool
        d_68_v4_ = True
        d_69_v5_: _dafny.MultiSet
        d_69_v5_ = _dafny.MultiSet([d_68_v4_, True, d_68_v4_])
        d_70_v6_: _dafny.MultiSet
        d_70_v6_ = _dafny.MultiSet([d_69_v5_])
        d_71_v7_: _dafny.Array
        nw9_ = _dafny.Array(_dafny.Set({}), 5)
        d_71_v7_ = nw9_
        d_72_v8_: _dafny.Array
        nw10_ = _dafny.Array(int(0), 16)
        d_72_v8_ = nw10_
        d_73_v9_: int
        d_73_v9_ = -20
        d_74_v10_: _dafny.Map
        d_74_v10_ = _dafny.Map({d_72_v8_: d_73_v9_})
        d_75_globalState_: GlobalState
        nw11_ = GlobalState()
        nw11_.ctor__(True, 907, d_67_v3_, -486, False, d_70_v6_, d_64_v0_, False, -977, False, 632, _dafny.CodePoint('f'), 179, d_71_v7_, 435, 814, True, 558, (d_74_v10_) | (d_74_v10_), False, d_72_v8_, False, 634)
        d_75_globalState_ = nw11_
        d_76_i0_: int
        d_76_i0_ = 0
        with _dafny.label("0"):
            while d_68_v4_:
                with _dafny.c_label("0"):
                    if (d_76_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_76_i0_ = (d_76_i0_) + (1)
                    d_77_v11_: _dafny.Array
                    nw12_ = _dafny.Array(False, 16)
                    d_77_v11_ = nw12_
                    d_78_v12_: D0
                    d_79_v13_: _dafny.Map
                    d_80_v14_: int
                    out0_: D0
                    out1_: _dafny.Map
                    out2_: int
                    out0_, out1_, out2_ = default__.m0(d_77_v11_, d_64_v0_, d_73_v9_, d_75_globalState_)
                    d_78_v12_ = out0_
                    d_79_v13_ = out1_
                    d_80_v14_ = out2_
                    d_81_v15_: _dafny.Map
                    d_81_v15_ = _dafny.Map({d_73_v9_: 921})
                    index0_ = default__.safeIndex(91, (d_72_v8_).length(0))
                    (d_72_v8_)[index0_] = ((d_81_v15_)[len(d_64_v0_)] if (len(d_64_v0_)) in (d_81_v15_) else 973)
                    index1_ = default__.safeIndex(91, (d_72_v8_).length(0))
                    (d_72_v8_)[index1_] = default__.fm0(d_73_v9_, d_68_v4_, d_75_globalState_)
                    d_82_v16_: _dafny.Array
                    nw13_ = _dafny.Array(None, 22)
                    nw13_[int(0)] = d_67_v3_
                    nw13_[int(1)] = d_67_v3_
                    nw13_[int(2)] = d_67_v3_
                    nw13_[int(3)] = d_67_v3_
                    nw13_[int(4)] = d_67_v3_
                    nw13_[int(5)] = (d_67_v3_ if d_68_v4_ else d_67_v3_)
                    nw13_[int(6)] = (d_67_v3_ if d_68_v4_ else d_67_v3_)
                    nw13_[int(7)] = d_67_v3_
                    nw13_[int(8)] = d_67_v3_
                    nw13_[int(9)] = d_67_v3_
                    nw13_[int(10)] = d_67_v3_
                    nw13_[int(11)] = d_67_v3_
                    nw13_[int(12)] = d_67_v3_
                    nw13_[int(13)] = d_67_v3_
                    nw13_[int(14)] = d_67_v3_
                    nw13_[int(15)] = d_67_v3_
                    nw13_[int(16)] = d_67_v3_
                    nw13_[int(17)] = d_67_v3_
                    nw13_[int(18)] = d_67_v3_
                    nw13_[int(19)] = d_67_v3_
                    nw13_[int(20)] = d_67_v3_
                    nw13_[int(21)] = d_67_v3_
                    d_82_v16_ = nw13_
                    index2_ = default__.safeIndex(99, (d_82_v16_).length(0))
                    (d_82_v16_)[index2_] = d_67_v3_
                    index3_ = default__.safeIndex(99, (d_82_v16_).length(0))
                    (d_82_v16_)[index3_] = d_67_v3_
                    d_83_v17_: _dafny.Seq
                    d_83_v17_ = _dafny.SeqWithoutIsStrInference([d_68_v4_, (len(d_64_v0_)) > (d_80_v14_), d_68_v4_])
                    d_83_v17_ = d_83_v17_
                    pass
            pass
        d_84_v18_: _dafny.Array
        nw14_ = _dafny.Array(_dafny.Map({}), 20)
        d_84_v18_ = nw14_
        d_85_v19_: _dafny.Map
        d_85_v19_ = _dafny.Map({378: (_dafny.MultiSet([d_73_v9_])).cardinality})
        index4_ = default__.safeIndex(957, (d_84_v18_).length(0))
        (d_84_v18_)[index4_] = d_85_v19_
        d_86_v21_: _dafny.Seq
        d_86_v21_ = _dafny.SeqWithoutIsStrInference([467])
        index5_ = default__.safeIndex(957, (d_84_v18_).length(0))
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in (d_86_v21_).Elements:
                d_87_v20_: int = compr_9_
                if (d_87_v20_) in (d_86_v21_):
                    coll9_[(d_87_v20_) - (732)] = d_73_v9_
            return _dafny.Map(coll9_)
        (d_84_v18_)[index5_] = ((d_85_v19_).set(688, d_73_v9_)) | ((iife9_()
        ) | (d_85_v19_))
        d_88_v22_: _dafny.Array
        nw15_ = _dafny.Array(False, 20)
        d_88_v22_ = nw15_
        d_89_v23_: _dafny.Set
        d_89_v23_ = _dafny.Set({d_68_v4_, d_68_v4_})
        index6_ = default__.safeIndex(547, (d_88_v22_).length(0))
        (d_88_v22_)[index6_] = (_dafny.Set({d_68_v4_, d_68_v4_, d_68_v4_, d_68_v4_})).ispropersubset(d_89_v23_)
        d_90_v24_: _dafny.Map
        d_90_v24_ = _dafny.Map({d_66_v2_: default__.fm1(d_73_v9_, d_68_v4_, d_75_globalState_)})
        d_91_v25_: _dafny.MultiSet
        d_91_v25_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxuh")), default__.fm2(d_75_globalState_), d_64_v0_])
        index7_ = default__.safeIndex(547, (d_88_v22_).length(0))
        rhs4_ = not (d_68_v4_) or ((_dafny.Map({d_66_v2_: not(not(d_68_v4_))})) != (d_90_v24_))
        rhs5_ = (d_73_v9_) * (((d_91_v25_)[d_64_v0_] if (d_64_v0_) in (d_91_v25_) else d_73_v9_))
        rhs6_ = default__.safeModuloInt(d_73_v9_, d_73_v9_)
        rhs7_ = not(True)
        lhs2_ = d_75_globalState_
        lhs3_ = d_75_globalState_
        lhs4_ = d_75_globalState_
        lhs5_ = d_88_v22_
        lhs6_ = default__.safeIndex(547, (d_88_v22_).length(0))
        lhs2_.f7 = rhs4_
        lhs3_.f14 = rhs5_
        lhs4_.f14 = rhs6_
        lhs5_[lhs6_] = rhs7_
        d_92_v26_: _dafny.Seq
        d_92_v26_ = _dafny.SeqWithoutIsStrInference([not(d_68_v4_)])
        hi1_ = len(((d_92_v26_).set(default__.safeIndex(d_73_v9_, len(d_92_v26_)), (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]) if True else d_92_v26_))
        for d_93_i1_ in range(d_73_v9_, hi1_):
            (d_75_globalState_).f14 = d_73_v9_
            d_94_v27_: _dafny.Array
            nw16_ = _dafny.Array(_dafny.Seq({}), 9)
            d_94_v27_ = nw16_
            d_95_v28_: _dafny.Map
            d_95_v28_ = _dafny.Map({d_93_i1_: d_94_v27_})
            d_95_v28_ = (d_95_v28_).set(default__.safeModuloInt(d_93_i1_, d_73_v9_), d_94_v27_)
            d_68_v4_ = (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]
            index8_ = default__.safeIndex(240, (d_72_v8_).length(0))
            (d_72_v8_)[index8_] = d_73_v9_
            index9_ = default__.safeIndex(240, (d_72_v8_).length(0))
            (d_72_v8_)[index9_] = default__.safeDivisionInt(len(_dafny.Set({d_68_v4_})), (len(d_86_v21_)) + (d_73_v9_))
        d_96_i2_: int
        d_96_i2_ = 0
        with _dafny.label("1"):
            while ((d_73_v9_) * (d_73_v9_)) != (default__.safeDivisionInt(d_73_v9_, (0) - (d_73_v9_))):
                with _dafny.c_label("1"):
                    if (d_96_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_96_i2_ = (d_96_i2_) + (1)
                    (d_75_globalState_).f8 = len(_dafny.SeqWithoutIsStrInference([(d_89_v23_) | (d_89_v23_), d_89_v23_, d_89_v23_, d_89_v23_]))
                    d_97_v29_: D0
                    d_98_v30_: _dafny.Map
                    d_99_v31_: int
                    out3_: D0
                    out4_: _dafny.Map
                    out5_: int
                    out3_, out4_, out5_ = default__.m0(d_88_v22_, d_64_v0_, d_73_v9_, d_75_globalState_)
                    d_97_v29_ = out3_
                    d_98_v30_ = out4_
                    d_99_v31_ = out5_
                    d_100_v32_: _dafny.Seq
                    d_100_v32_ = _dafny.SeqWithoutIsStrInference([default__.fm3(d_73_v9_, d_75_globalState_), d_90_v24_, _dafny.Map({d_66_v2_: d_68_v4_}), d_90_v24_, _dafny.Map({d_66_v2_: d_68_v4_})])
                    d_101_v33_: D0
                    d_102_v34_: _dafny.Map
                    d_103_v35_: int
                    out6_: D0
                    out7_: _dafny.Map
                    out8_: int
                    out6_, out7_, out8_ = default__.m0(d_88_v22_, _dafny.SeqWithoutIsStrInference([d_66_v2_ for d_104_i3_ in range(default__.abs(751))]), (len(_dafny.Map({len(d_100_v32_): d_99_v31_}))) * ((default__.fm4(D0_DC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))), 998, d_92_v26_, d_66_v2_, d_75_globalState_)).cardinality), d_75_globalState_)
                    d_101_v33_ = out6_
                    d_102_v34_ = out7_
                    d_103_v35_ = out8_
                    index10_ = default__.safeIndex(279, (d_72_v8_).length(0))
                    (d_72_v8_)[index10_] = -24
                    d_105_v36_: _dafny.Map
                    d_105_v36_ = _dafny.Map({(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]: d_72_v8_})
                    index11_ = default__.safeIndex(279, (d_72_v8_).length(0))
                    (d_72_v8_)[index11_] = default__.fm0(d_99_v31_, (d_73_v9_) < (len(d_105_v36_)), d_75_globalState_)
                    pass
            pass
        (d_75_globalState_).f6 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahdth"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syrkxpe")))
        (d_75_globalState_).f14 = (0) - (default__.safeModuloInt((d_73_v9_ if d_68_v4_ else d_73_v9_), (0) - (len(default__.fm2(d_75_globalState_)))))
        if (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]:
            (d_75_globalState_).f7 = d_68_v4_
            (d_75_globalState_).f8 = default__.safeDivisionInt(len(d_64_v0_), (d_73_v9_) + (30))
            d_106_v37_: _dafny.Set
            d_106_v37_ = _dafny.Set({d_73_v9_})
            d_107_v38_: _dafny.Seq
            d_107_v38_ = _dafny.SeqWithoutIsStrInference([default__.fm5(d_106_v37_, d_75_globalState_)])
            (d_75_globalState_).f8 = default__.safeModuloInt(d_73_v9_, default__.safeDivisionInt((0) - (len(d_64_v0_)), len(d_107_v38_)))
            d_108_v39_: _dafny.Map
            d_108_v39_ = _dafny.Map({d_68_v4_: d_73_v9_})
            d_109_v40_: D0
            d_109_v40_ = D0_DC1(d_73_v9_, (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))], d_73_v9_)
            d_108_v39_ = ((_dafny.Map({True: len((d_64_v0_).set(default__.safeIndex(d_73_v9_, len(d_64_v0_)), d_66_v2_))})).set((d_109_v40_).cf2, -355)) | (d_108_v39_)
            d_86_v21_ = d_86_v21_
        elif True:
            d_110_v41_: _dafny.MultiSet
            d_110_v41_ = _dafny.MultiSet([d_73_v9_, len(d_86_v21_)])
            index12_ = default__.safeIndex(547, (d_88_v22_).length(0))
            rhs8_ = d_73_v9_
            rhs9_ = ((d_64_v0_).set(default__.safeIndex(d_73_v9_, len(d_64_v0_)), d_66_v2_)).set(default__.safeIndex(d_73_v9_, len((d_64_v0_).set(default__.safeIndex(d_73_v9_, len(d_64_v0_)), d_66_v2_))), d_66_v2_)
            rhs10_ = d_68_v4_
            rhs11_ = (d_110_v41_).issubset(d_110_v41_)
            rhs12_ = (d_92_v26_)[default__.safeIndex((len(d_92_v26_)) * (len((default__.fm2(d_75_globalState_)).set(default__.safeIndex(len(d_86_v21_), len(default__.fm2(d_75_globalState_))), d_66_v2_))), len(d_92_v26_))]
            lhs7_ = d_75_globalState_
            lhs8_ = d_75_globalState_
            lhs9_ = d_88_v22_
            lhs10_ = default__.safeIndex(547, (d_88_v22_).length(0))
            lhs11_ = d_75_globalState_
            lhs7_.f14 = rhs8_
            lhs8_.f6 = rhs9_
            lhs9_[lhs10_] = rhs10_
            d_68_v4_ = rhs11_
            lhs11_.f21 = rhs12_
            (d_75_globalState_).f11 = default__.fm6(d_73_v9_, d_75_globalState_)
            (d_75_globalState_).f9 = (D0_DC1(d_73_v9_, (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))], d_73_v9_)).cf2
            index13_ = default__.safeIndex(547, (d_88_v22_).length(0))
            index14_ = default__.safeIndex(547, (d_88_v22_).length(0))
            rhs13_ = ((d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))] if (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))] else d_68_v4_)
            rhs14_ = d_68_v4_
            rhs15_ = (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]
            lhs12_ = d_75_globalState_
            lhs13_ = d_88_v22_
            lhs14_ = default__.safeIndex(547, (d_88_v22_).length(0))
            lhs15_ = d_88_v22_
            lhs16_ = default__.safeIndex(547, (d_88_v22_).length(0))
            lhs12_.f7 = rhs13_
            lhs13_[lhs14_] = rhs14_
            lhs15_[lhs16_] = rhs15_
            d_111_v42_: D0
            d_112_v43_: _dafny.Map
            d_113_v44_: int
            out9_: D0
            out10_: _dafny.Map
            out11_: int
            out9_, out10_, out11_ = default__.m0(d_88_v22_, _dafny.SeqWithoutIsStrInference([d_66_v2_ for d_114_i4_ in range(default__.abs(-195))]), (0) - (d_73_v9_), d_75_globalState_)
            d_111_v42_ = out9_
            d_112_v43_ = out10_
            d_113_v44_ = out11_
        hi2_ = d_73_v9_
        for d_115_i5_ in range(d_73_v9_, hi2_):
            index15_ = default__.safeIndex(547, (d_88_v22_).length(0))
            (d_88_v22_)[index15_] = (d_64_v0_) != (d_64_v0_)
            (d_75_globalState_).f21 = d_68_v4_
            (d_75_globalState_).f8 = (d_115_i5_) + (default__.fm0(892, d_68_v4_, d_75_globalState_))
            d_116_v45_: D0
            d_117_v46_: _dafny.Map
            d_118_v47_: int
            out12_: D0
            out13_: _dafny.Map
            out14_: int
            out12_, out13_, out14_ = default__.m0((d_88_v22_ if (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))] else d_88_v22_), d_64_v0_, d_115_i5_, d_75_globalState_)
            d_116_v45_ = out12_
            d_117_v46_ = out13_
            d_118_v47_ = out14_
        d_119_i6_: int
        d_119_i6_ = 0
        with _dafny.label("2"):
            while (d_73_v9_) >= (-886):
                with _dafny.c_label("2"):
                    if (d_119_i6_) >= (100):
                        raise _dafny.Break("2")
                    d_119_i6_ = (d_119_i6_) + (1)
                    d_120_v48_: _dafny.Map
                    d_120_v48_ = _dafny.Map({d_73_v9_: d_68_v4_})
                    rhs16_ = d_72_v8_
                    rhs17_ = (len(d_120_v48_)) - ((0) - (d_73_v9_))
                    rhs18_ = (default__.fm0(d_73_v9_, d_68_v4_, d_75_globalState_)) == (743)
                    rhs19_ = (d_68_v4_ if d_68_v4_ else d_68_v4_)
                    rhs20_ = d_73_v9_
                    lhs17_ = d_75_globalState_
                    lhs18_ = d_75_globalState_
                    lhs19_ = d_75_globalState_
                    lhs20_ = d_75_globalState_
                    d_72_v8_ = rhs16_
                    lhs17_.f14 = rhs17_
                    lhs18_.f21 = rhs18_
                    lhs19_.f7 = rhs19_
                    lhs20_.f14 = rhs20_
                    (d_75_globalState_).f9 = (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]
                    d_121_v49_: _dafny.Seq
                    d_121_v49_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]]), d_92_v26_])
                    (d_75_globalState_).f14 = default__.safeModuloInt(d_73_v9_, (d_73_v9_) - (len((d_121_v49_)[default__.safeIndex(d_73_v9_, len(d_121_v49_))])))
                    d_122_v50_: _dafny.Array
                    nw17_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                    d_122_v50_ = nw17_
                    index16_ = default__.safeIndex(963, (d_122_v50_).length(0))
                    (d_122_v50_)[index16_] = d_66_v2_
                    d_123_v51_: _dafny.MultiSet
                    d_123_v51_ = _dafny.MultiSet([d_73_v9_])
                    index17_ = default__.safeIndex(963, (d_122_v50_).length(0))
                    rhs21_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghjmshpn"))) + (d_64_v0_)
                    rhs22_ = (d_64_v0_)[default__.safeIndex((d_123_v51_).cardinality, len(d_64_v0_))]
                    rhs23_ = d_66_v2_
                    lhs21_ = d_75_globalState_
                    lhs22_ = d_75_globalState_
                    lhs23_ = d_122_v50_
                    lhs24_ = default__.safeIndex(963, (d_122_v50_).length(0))
                    lhs21_.f6 = rhs21_
                    lhs22_.f11 = rhs22_
                    lhs23_[lhs24_] = rhs23_
                    pass
            pass
        d_124_v52_: D0
        d_125_v53_: _dafny.Map
        d_126_v54_: int
        out15_: D0
        out16_: _dafny.Map
        out17_: int
        out15_, out16_, out17_ = default__.m0(d_88_v22_, default__.fm2(d_75_globalState_), d_73_v9_, d_75_globalState_)
        d_124_v52_ = out15_
        d_125_v53_ = out16_
        d_126_v54_ = out17_
        d_127_v55_: _dafny.Set
        d_127_v55_ = _dafny.Set({len(d_64_v0_), d_73_v9_})
        hi3_ = len(d_127_v55_)
        for d_128_i7_ in range(((d_69_v5_) - (d_69_v5_)).cardinality, hi3_):
            rhs24_ = (0) - ((0) - (len(default__.fm2(d_75_globalState_))))
            rhs25_ = default__.safeDivisionInt(-737, d_128_i7_)
            rhs26_ = d_72_v8_
            lhs25_ = d_75_globalState_
            lhs25_.f8 = rhs24_
            d_126_v54_ = rhs25_
            d_72_v8_ = rhs26_
            d_129_v56_: _dafny.Array
            nw18_ = _dafny.Array(_dafny.Map({}), 28)
            d_129_v56_ = nw18_
            d_129_v56_ = d_129_v56_
            d_130_v57_: _dafny.Array
            nw19_ = _dafny.Array(_dafny.Seq({}), 21)
            d_130_v57_ = nw19_
            d_131_v58_: _dafny.Seq
            d_131_v58_ = _dafny.SeqWithoutIsStrInference([d_72_v8_])
            d_132_v59_: D1
            d_132_v59_ = D1_DC3(d_131_v58_)
            index18_ = default__.safeIndex(996, (d_130_v57_).length(0))
            (d_130_v57_)[index18_] = (d_132_v59_).cf5
            index19_ = default__.safeIndex(996, (d_130_v57_).length(0))
            (d_130_v57_)[index19_] = d_131_v58_
            d_133_v60_: D1
            d_133_v60_ = D1_DC5(d_126_v54_, True, len(d_127_v55_), default__.fm0(d_73_v9_, d_68_v4_, d_75_globalState_))
            index20_ = default__.safeIndex(957, (d_84_v18_).length(0))
            (d_84_v18_)[index20_] = ((d_84_v18_)[default__.safeIndex(957, (d_84_v18_).length(0))]).set((default__.fm0(default__.fm0(d_126_v54_, d_68_v4_, d_75_globalState_), (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))], d_75_globalState_)) + ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([d_73_v9_, d_128_i7_, (d_133_v60_).cf12, d_128_i7_]))))), d_73_v9_)
        d_134_v61_: _dafny.MultiSet
        d_134_v61_ = _dafny.MultiSet([d_73_v9_, len((_dafny.SeqWithoutIsStrInference([75])).set(default__.safeIndex(d_73_v9_, len(_dafny.SeqWithoutIsStrInference([75]))), d_126_v54_))])
        if (d_68_v4_) == (((d_134_v61_).set(len(_dafny.Set({d_65_v1_, D0_DC0(d_64_v0_)})), default__.abs(len(d_64_v0_)))).isdisjoint(d_134_v61_)):
            (d_75_globalState_).f9 = d_68_v4_
            index21_ = default__.safeIndex(268, (d_72_v8_).length(0))
            (d_72_v8_)[index21_] = (0) - (d_73_v9_)
            index22_ = default__.safeIndex(268, (d_72_v8_).length(0))
            (d_72_v8_)[index22_] = -38
            d_135_v62_: C0
            nw20_ = C0()
            nw20_.ctor__()
            d_135_v62_ = nw20_
            d_136_v63_: _dafny.Seq
            d_136_v63_ = _dafny.SeqWithoutIsStrInference([d_69_v5_, _dafny.MultiSet([(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]])])
            d_137_v64_: C3
            nw21_ = C3()
            nw21_.ctor__(d_136_v63_, (d_72_v8_)[default__.safeIndex(268, (d_72_v8_).length(0))], not(d_68_v4_), _dafny.SeqWithoutIsStrInference([d_73_v9_]))
            d_137_v64_ = nw21_
            d_138_v65_: _dafny.MultiSet
            d_138_v65_ = _dafny.MultiSet([_dafny.Map({d_135_v62_: d_137_v64_})])
            d_139_v66_: _dafny.Map
            d_139_v66_ = _dafny.Map({d_135_v62_: d_137_v64_})
            d_140_v67_: C3
            nw22_ = C3()
            nw22_.ctor__(default__.fm29(d_73_v9_, _dafny.Set({d_126_v54_, d_73_v9_}), d_75_globalState_), (d_72_v8_)[default__.safeIndex(268, (d_72_v8_).length(0))], ((d_72_v8_)[default__.safeIndex(268, (d_72_v8_).length(0))]) not in (d_86_v21_), (d_86_v21_).set(default__.safeIndex(((d_138_v65_)[d_139_v66_] if (d_139_v66_) in (d_138_v65_) else d_126_v54_), len(d_86_v21_)), d_73_v9_))
            d_140_v67_ = nw22_
            index23_ = default__.safeIndex(268, (d_72_v8_).length(0))
            rhs27_ = d_88_v22_
            rhs28_ = d_126_v54_
            rhs29_ = default__.fm1(d_73_v9_, (not(True) if False else (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]), d_75_globalState_)
            lhs26_ = d_72_v8_
            lhs27_ = default__.safeIndex(268, (d_72_v8_).length(0))
            lhs28_ = d_75_globalState_
            d_88_v22_ = rhs27_
            lhs26_[lhs27_] = rhs28_
            lhs28_.f7 = rhs29_
            index24_ = default__.safeIndex(547, (d_88_v22_).length(0))
            (d_88_v22_)[index24_] = (d_64_v0_) < (d_64_v0_)
        elif True:
            d_141_v68_: _dafny.Map
            d_141_v68_ = _dafny.Map({d_68_v4_: not((d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))])})
            d_142_v69_: _dafny.Map
            d_142_v69_ = _dafny.Map({(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]: len(d_141_v68_)})
            d_73_v9_ = ((d_69_v5_)[(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]] if ((d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]) in (d_69_v5_) else ((d_142_v69_)[(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]] if ((d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]) in (d_142_v69_) else d_73_v9_))
            d_143_v70_: _dafny.Map
            d_143_v70_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_66_v2_ for d_144_i8_ in range(default__.abs(347))]): (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]})
            d_143_v70_ = (d_143_v70_).set(d_64_v0_, d_68_v4_)
            d_145_v71_: _dafny.Seq
            d_145_v71_ = _dafny.SeqWithoutIsStrInference([d_92_v26_, d_92_v26_])
            d_92_v26_ = (d_92_v26_) + (((d_145_v71_)[default__.safeIndex((0) - (len(d_64_v0_)), len(d_145_v71_))]) + (d_92_v26_))
            d_146_v72_: _dafny.Seq
            d_146_v72_ = _dafny.SeqWithoutIsStrInference([d_141_v68_])
            d_146_v72_ = d_146_v72_
            index25_ = default__.safeIndex(201, (d_72_v8_).length(0))
            (d_72_v8_)[index25_] = d_126_v54_
            index26_ = default__.safeIndex(201, (d_72_v8_).length(0))
            rhs30_ = d_68_v4_
            rhs31_ = d_73_v9_
            rhs32_ = d_73_v9_
            rhs33_ = d_73_v9_
            lhs29_ = d_75_globalState_
            lhs30_ = d_72_v8_
            lhs31_ = default__.safeIndex(201, (d_72_v8_).length(0))
            lhs29_.f7 = rhs30_
            lhs30_[lhs31_] = rhs31_
            d_73_v9_ = rhs32_
            d_73_v9_ = rhs33_
        d_147_v73_: _dafny.Map
        d_147_v73_ = _dafny.Map({not(d_68_v4_): (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]})
        d_147_v73_ = (d_147_v73_).set(d_68_v4_, d_68_v4_)
        d_148_v74_: _dafny.Array
        nw23_ = _dafny.Array(None, 4)
        d_148_v74_ = nw23_
        d_149_v75_: _dafny.Map
        d_149_v75_ = _dafny.Map({(d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))]: d_148_v74_})
        d_150_v76_: T1
        nw24_ = C2()
        nw24_.ctor__((d_84_v18_)[default__.safeIndex(957, (d_84_v18_).length(0))], (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))], d_68_v4_, default__.fm19(d_73_v9_, (d_88_v22_)[default__.safeIndex(547, (d_88_v22_).length(0))], d_75_globalState_), len(d_149_v75_))
        d_150_v76_ = nw24_
        d_151_v77_: _dafny.Array
        nw25_ = _dafny.Array(_dafny.MultiSet({}), 11)
        d_151_v77_ = nw25_
        d_152_v78_: _dafny.Map
        d_152_v78_ = _dafny.Map({d_150_v76_: d_151_v77_})
        d_152_v78_ = (d_152_v78_).set(d_150_v76_, d_151_v77_)
        (d_75_globalState_).f7 = (default__.safeModuloInt(d_126_v54_, len((d_64_v0_).set(default__.safeIndex(d_73_v9_, len(d_64_v0_)), d_66_v2_)))) <= (d_73_v9_)
        _dafny.print((d_64_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_65_v1_).cf0).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_66_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_67_v3_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_68_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v5_) == (_dafny.MultiSet([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_70_v6_) == (_dafny.MultiSet([_dafny.MultiSet([True, True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_72_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_72_v8_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_72_v8_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_72_v8_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_73_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_74_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_75_globalState_).f2)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_globalState_).f5) == (_dafny.MultiSet([_dafny.MultiSet([True, True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_75_globalState_.f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_75_globalState_).f18)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_globalState_).f20)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_globalState_).f20)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_globalState_).f20)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_globalState_).f20)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_75_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_76_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_v18_)[17]) == (_dafny.Map({378: 1, 688: -20, -265: -20, 4: -20}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_v19_) == (_dafny.Map({378: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_86_v21_) == (_dafny.SeqWithoutIsStrInference([467]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_88_v22_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v23_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_v24_) == (_dafny.Map({_dafny.CodePoint('t'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v25_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxuh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hg"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v26_) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_96_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_119_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v52_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v52_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v52_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v53_) == (_dafny.Map({0: False, -1: False, 1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_v54_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_127_v55_) == (_dafny.Set({2, -20}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v61_) == (_dafny.MultiSet([-20, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v73_) == (_dafny.Map({False: False, True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_149_v75_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v76_).f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v76_.f29) == (_dafny.SeqWithoutIsStrInference([3, -2, 1, 1, 7]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v76_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_152_v78_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), False, int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.Map({}), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(int(0), int(0), _dafny.CodePoint('D'), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D3_DC13)

class D3_DC10(D3, NamedTuple('DC10', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC13(D3, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC15(int(0), _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)

class D4_DC15(D4, NamedTuple('DC15', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC16(D5, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC18(D7, NamedTuple('DC18', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC20(D8, NamedTuple('DC20', [('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(False, _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)

class D9_DC25(D9, NamedTuple('DC25', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC28(D9, NamedTuple('DC28', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC29(D10, NamedTuple('DC29', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)

class D11_DC31(D11, NamedTuple('DC31', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({self.cf53.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC35()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)

class D12_DC35(D12, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC39(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)

class D13_DC39(D13, NamedTuple('DC39', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC40(D13, NamedTuple('DC40', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC42(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC42(D14, NamedTuple('DC42', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC44(False, _dafny.Seq({}), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D15_DC45)

class D15_DC44(D15, NamedTuple('DC44', [('cf68', Any), ('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC43(D15, NamedTuple('DC43', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC45(D15, NamedTuple('DC45', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC45({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC45) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC47(int(0), int(0), None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D16_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)

class D16_DC47(D16, NamedTuple('DC47', [('cf75', Any), ('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC47({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC47) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC46(D16, NamedTuple('DC46', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC49(int(0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)

class D17_DC49(D17, NamedTuple('DC49', [('cf80', Any), ('cf81', Any), ('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC51()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)

class D18_DC51(D18, NamedTuple('DC51', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC53(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D19_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D19_DC54)

class D19_DC53(D19, NamedTuple('DC53', [('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC53({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC53) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC52(D19, NamedTuple('DC52', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC54(D19, NamedTuple('DC54', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC54({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC54) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC56(False, int(0), _dafny.Array(None, 0), _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D20_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D20_DC57)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D20_DC58)

class D20_DC56(D20, NamedTuple('DC56', [('cf90', Any), ('cf91', Any), ('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC56({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC56) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC57(D20, NamedTuple('DC57', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC57'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC57)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC55(D20, NamedTuple('DC55', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC58(D20, NamedTuple('DC58', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC58({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC58) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, globalState):
        pass

    def m2(self, p0, globalState):
        pass


class T1:
    pass
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value

class GlobalState:
    def  __init__(self):
        self.f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f7: bool = False
        self.f8: int = int(0)
        self.f9: bool = False
        self.f11: str = _dafny.CodePoint('D')
        self.f14: int = int(0)
        self.f21: bool = False
        self._f0: bool = False
        self._f1: int = int(0)
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        self._f3: int = int(0)
        self._f4: bool = False
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f10: int = int(0)
        self._f12: int = int(0)
        self._f13: _dafny.Array = _dafny.Array(None, 0)
        self._f15: int = int(0)
        self._f16: bool = False
        self._f17: int = int(0)
        self._f18: _dafny.Map = _dafny.Map({})
        self._f19: bool = False
        self._f20: _dafny.Array = _dafny.Array(None, 0)
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self)._f22 = f22

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
    @property
    def f5(self):
        return self._f5
    @property
    def f10(self):
        return self._f10
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
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f22(self):
        return self._f22

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm20(self, p0, p1, p2, p3, globalState):
        return False


class C1(T1, T0):
    def  __init__(self):
        self._f29: _dafny.Seq = _dafny.Seq({})
        self._f23: int = int(0)
        self._f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f23(self):
        return self._f23
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f28, f29, f23):
        (self)._f28 = f28
        (self).f29 = f29
        (self)._f23 = f23

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        index27_ = default__.safeIndex(216, (p0).length(0))
        (p0)[index27_] = p1
        index28_ = default__.safeIndex(216, (p0).length(0))
        (p0)[index28_] = (self).f28
        d_153_v0_: _dafny.Seq
        d_153_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pilipchus"))
        d_154_v1_: _dafny.Set
        d_154_v1_ = _dafny.Set({d_153_v0_})
        (self).f29 = default__.fm19(len(d_153_v0_), (d_153_v0_) not in (d_154_v1_), globalState)
        d_155_v2_: str
        d_155_v2_ = _dafny.CodePoint('n')
        d_156_v3_: _dafny.MultiSet
        d_156_v3_ = _dafny.MultiSet([d_155_v2_])
        index29_ = default__.safeIndex(216, (p0).length(0))
        rhs34_ = default__.fm0((self).f23, (p0)[default__.safeIndex(216, (p0).length(0))], globalState)
        rhs35_ = (self).f28
        rhs36_ = (self).f23
        rhs37_ = d_156_v3_
        rhs38_ = p1
        lhs32_ = globalState
        lhs33_ = globalState
        lhs34_ = p0
        lhs35_ = default__.safeIndex(216, (p0).length(0))
        lhs32_.f14 = rhs34_
        lhs33_.f7 = rhs35_
        r0 = rhs36_
        d_156_v3_ = rhs37_
        lhs34_[lhs35_] = rhs38_
        d_157_v5_: _dafny.Map
        d_157_v5_ = _dafny.Map({p0: (self).f23})
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([(self).f23 for d_185_i1_ in range(default__.abs(965))])).Elements:
                d_186_v4_: int = compr_10_
                if (d_186_v4_) in (_dafny.SeqWithoutIsStrInference([(self).f23 for d_185_i1_ in range(default__.abs(965))])):
                    coll10_ = coll10_.union(_dafny.Set([(d_186_v4_) - (783)]))
            return _dafny.Set(coll10_)
        hi4_ = (self).f23
        for d_158_i0_ in range(len(_dafny.SeqWithoutIsStrInference([len(iife10_()
        ), (self).f23, ((d_157_v5_)[p0] if (p0) in (d_157_v5_) else (self).f23)])), hi4_):
            if p1:
                (globalState).f14 = ((d_158_i0_) * ((self).f23) if (d_153_v0_) <= (d_153_v0_) else d_158_i0_)
                d_159_v6_: _dafny.Seq
                d_159_v6_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([p1, (self).f28, (p0)[default__.safeIndex(216, (p0).length(0))], default__.fm1((self).f23, (self).f28, globalState), p1])).set(default__.safeIndex(d_158_i0_, len(_dafny.SeqWithoutIsStrInference([p1, (self).f28, (p0)[default__.safeIndex(216, (p0).length(0))], default__.fm1((self).f23, (self).f28, globalState), p1]))), (self).f28)])
                d_160_v7_: _dafny.Map
                d_160_v7_ = _dafny.Map({d_158_i0_: (self).f23})
                d_161_v8_: D1
                d_161_v8_ = D1_DC4(d_160_v7_, (p0)[default__.safeIndex(216, (p0).length(0))], (self).f28)
                d_162_v9_: _dafny.Map
                d_162_v9_ = _dafny.Map({(d_159_v6_) + ((d_159_v6_).set(default__.safeIndex((self).f23, len(d_159_v6_)), _dafny.SeqWithoutIsStrInference([p1, (p0)[default__.safeIndex(216, (p0).length(0))]]))): (d_161_v8_).cf8})
                d_162_v9_ = (d_162_v9_).set(d_159_v6_, (self).f28)
                d_163_v10_: _dafny.Map
                d_163_v10_ = _dafny.Map({(p0)[default__.safeIndex(216, (p0).length(0))]: -375})
                d_164_v11_: _dafny.Seq
                d_164_v11_ = _dafny.SeqWithoutIsStrInference([(self).f28])
                rhs39_ = ((self).f23) + (((d_163_v10_)[default__.fm1(len(d_164_v11_), (self).f28, globalState)] if (default__.fm1(len(d_164_v11_), (self).f28, globalState)) in (d_163_v10_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_165_i2_ in range(default__.abs(-960))]))))
                rhs40_ = ((self).f23) > ((self).f23)
                lhs36_ = globalState
                lhs37_ = globalState
                lhs36_.f14 = rhs39_
                lhs37_.f9 = rhs40_
                d_166_v12_: _dafny.Array
                def lambda1_(d_167_i3_):
                    return (d_167_i3_) * (635)

                init0_ = lambda1_
                nw26_ = _dafny.Array(None, 13)
                for i0_0_ in range(nw26_.length(0)):
                    nw26_[i0_0_] = init0_(i0_0_)
                d_166_v12_ = nw26_
                d_168_v13_: _dafny.Seq
                d_168_v13_ = _dafny.SeqWithoutIsStrInference([d_166_v12_])
                d_169_v14_: D1
                d_169_v14_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_166_v12_, d_166_v12_, d_166_v12_]))
                d_170_v15_: _dafny.Array
                nw27_ = _dafny.Array(None, 10)
                nw27_[int(0)] = D1_DC3(d_168_v13_)
                nw27_[int(1)] = d_169_v14_
                nw27_[int(2)] = d_169_v14_
                nw27_[int(3)] = d_169_v14_
                nw27_[int(4)] = D1_DC3(d_168_v13_)
                nw27_[int(5)] = d_169_v14_
                nw27_[int(6)] = d_169_v14_
                nw27_[int(7)] = d_169_v14_
                nw27_[int(8)] = d_169_v14_
                nw27_[int(9)] = d_169_v14_
                d_170_v15_ = nw27_
                d_171_v16_: _dafny.MultiSet
                d_171_v16_ = _dafny.MultiSet([d_170_v15_, d_170_v15_])
                d_172_v17_: D3
                d_172_v17_ = D3_DC9(d_171_v16_)
                d_173_v18_: _dafny.Seq
                d_173_v18_ = _dafny.SeqWithoutIsStrInference([d_172_v17_, d_172_v17_, d_172_v17_, d_172_v17_])
                d_173_v18_ = ((d_173_v18_ if (self).f28 else d_173_v18_)) + (d_173_v18_)
                (globalState).f21 = (self).f28
            elif True:
                d_174_v19_: D0
                d_175_v20_: _dafny.Map
                d_176_v21_: int
                out18_: D0
                out19_: _dafny.Map
                out20_: int
                out18_, out19_, out20_ = default__.m0(p0, (default__.fm2(globalState)) + (d_153_v0_), (self.f29)[default__.safeIndex((self).f23, len(self.f29))], globalState)
                d_174_v19_ = out18_
                d_175_v20_ = out19_
                d_176_v21_ = out20_
                d_177_v22_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                d_177_v22_ = nw28_
                nw29_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_177_v22_ = nw29_
                d_178_v23_: _dafny.Map
                d_178_v23_ = _dafny.Map({len(d_153_v0_): d_176_v21_})
                d_179_v24_: _dafny.Set
                d_179_v24_ = _dafny.Set({len(d_178_v23_), (self).f23})
                index30_ = default__.safeIndex(216, (p0).length(0))
                rhs41_ = p1
                rhs42_ = not((True if (d_179_v24_).ispropersubset(d_179_v24_) else (d_153_v0_) != (_dafny.SeqWithoutIsStrInference([d_155_v2_ for d_180_i4_ in range(default__.abs(196))]))))
                rhs43_ = 595
                lhs38_ = p0
                lhs39_ = default__.safeIndex(216, (p0).length(0))
                lhs40_ = globalState
                lhs41_ = globalState
                lhs38_[lhs39_] = rhs41_
                lhs40_.f7 = rhs42_
                lhs41_.f8 = rhs43_
                d_181_v25_: _dafny.Map
                d_181_v25_ = _dafny.Map({len(self.f29): d_155_v2_})
                (globalState).f11 = ((d_181_v25_)[(self).f23] if ((self).f23) in (d_181_v25_) else default__.fm6(d_158_i0_, globalState))
                d_182_v26_: _dafny.Array
                nw30_ = _dafny.Array(False, 17)
                d_182_v26_ = nw30_
                d_182_v26_ = d_182_v26_
            d_183_v27_: _dafny.Seq
            d_183_v27_ = _dafny.SeqWithoutIsStrInference([(self).f28])
            d_184_v28_: _dafny.Seq
            d_184_v28_ = _dafny.SeqWithoutIsStrInference([d_183_v27_, d_183_v27_])
            index31_ = default__.safeIndex(216, (p0).length(0))
            rhs44_ = p1
            rhs45_ = (d_184_v28_) + ((d_184_v28_).set(default__.safeIndex(d_158_i0_, len(d_184_v28_)), d_183_v27_))
            rhs46_ = (d_183_v27_)[default__.safeIndex(default__.safeDivisionInt((self).f23, d_158_i0_), len(d_183_v27_))]
            lhs42_ = p0
            lhs43_ = default__.safeIndex(216, (p0).length(0))
            lhs44_ = globalState
            lhs42_[lhs43_] = rhs44_
            d_184_v28_ = rhs45_
            lhs44_.f21 = rhs46_
            (globalState).f8 = ((950) + ((self).f23) if (p0)[default__.safeIndex(216, (p0).length(0))] else d_158_i0_)
            (globalState).f21 = (self).f28
        d_187_v29_: _dafny.Seq
        d_187_v29_ = _dafny.SeqWithoutIsStrInference([False])
        d_188_v30_: _dafny.Map
        d_188_v30_ = _dafny.Map({(self).f28: (self).f23})
        d_189_v31_: _dafny.Seq
        d_189_v31_ = _dafny.SeqWithoutIsStrInference([True, not((len(d_187_v29_)) >= ((self).f23)), p1, (self).f28, (False) in (d_188_v30_)])
        if (d_189_v31_)[default__.safeIndex((self).f23, len(d_189_v31_))]:
            d_190_v32_: _dafny.Array
            nw31_ = _dafny.Array(int(0), 7)
            d_190_v32_ = nw31_
            index32_ = default__.safeIndex(532, (d_190_v32_).length(0))
            (d_190_v32_)[index32_] = len(self.f29)
            d_191_v33_: _dafny.Map
            d_191_v33_ = _dafny.Map({(self).f23: d_190_v32_})
            d_192_v34_: _dafny.Map
            d_192_v34_ = _dafny.Map({d_191_v33_: True})
            index33_ = default__.safeIndex(532, (d_190_v32_).length(0))
            rhs47_ = (0) - ((self).f23)
            rhs48_ = (self.f29) + (_dafny.SeqWithoutIsStrInference([715 for d_193_i5_ in range(default__.abs(288))]))
            rhs49_ = (0) - ((self).f23)
            rhs50_ = ((d_192_v34_)[d_191_v33_] if (d_191_v33_) in (d_192_v34_) else default__.fm1((self).f23, p1, globalState))
            rhs51_ = (len(d_189_v31_) if not (not(not((p0)[default__.safeIndex(216, (p0).length(0))]))) or ((self).f28) else (self).f23)
            lhs45_ = globalState
            lhs46_ = self
            lhs47_ = d_190_v32_
            lhs48_ = default__.safeIndex(532, (d_190_v32_).length(0))
            lhs49_ = globalState
            lhs45_.f14 = rhs47_
            lhs46_.f29 = rhs48_
            lhs47_[lhs48_] = rhs49_
            lhs49_.f21 = rhs50_
            r0 = rhs51_
            (globalState).f14 = (d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))]
            d_194_v35_: _dafny.Seq
            d_194_v35_ = _dafny.SeqWithoutIsStrInference([d_153_v0_, default__.fm2(globalState), d_153_v0_, d_153_v0_, d_153_v0_])
            d_195_v37_: _dafny.Map
            d_195_v37_ = _dafny.Map({831: (self).f28})
            def iife11_():
                def iife14_():
                    coll14_ = _dafny.Set()
                    compr_14_: _dafny.Seq
                    for compr_14_ in (d_194_v35_).Elements:
                        d_199_v36_: _dafny.Seq = compr_14_
                        if (d_199_v36_) in (d_194_v35_):
                            coll14_ = coll14_.union(_dafny.Set([d_199_v36_]))
                    return _dafny.Set(coll14_)
                def iife15_():
                    coll15_ = _dafny.Set()
                    compr_15_: _dafny.Seq
                    for compr_15_ in (_dafny.Map({d_153_v0_: ((d_195_v37_)[len(d_188_v30_)] if (len(d_188_v30_)) in (d_195_v37_) else p1)})).keys.Elements:
                        d_200_v38_: _dafny.Seq = compr_15_
                        if (d_200_v38_) in (_dafny.Map({d_153_v0_: ((d_195_v37_)[len(d_188_v30_)] if (len(d_188_v30_)) in (d_195_v37_) else p1)})):
                            coll15_ = coll15_.union(_dafny.Set([d_200_v38_]))
                    return _dafny.Set(coll15_)
                coll11_ = _dafny.Set()
                def iife12_():
                    coll12_ = _dafny.Set()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (d_194_v35_).Elements:
                        d_196_v36_: _dafny.Seq = compr_12_
                        if (d_196_v36_) in (d_194_v35_):
                            coll12_ = coll12_.union(_dafny.Set([d_196_v36_]))
                    return _dafny.Set(coll12_)
                def iife13_():
                    coll13_ = _dafny.Set()
                    compr_13_: _dafny.Seq
                    for compr_13_ in (_dafny.Map({d_153_v0_: ((d_195_v37_)[len(d_188_v30_)] if (len(d_188_v30_)) in (d_195_v37_) else p1)})).keys.Elements:
                        d_197_v38_: _dafny.Seq = compr_13_
                        if (d_197_v38_) in (_dafny.Map({d_153_v0_: ((d_195_v37_)[len(d_188_v30_)] if (len(d_188_v30_)) in (d_195_v37_) else p1)})):
                            coll13_ = coll13_.union(_dafny.Set([d_197_v38_]))
                    return _dafny.Set(coll13_)
                compr_11_: _dafny.Seq
                for compr_11_ in ((iife12_()
                ) - (iife13_()
                )).Elements:
                    d_198_v39_: _dafny.Seq = compr_11_
                    if (d_198_v39_) in ((iife14_()
                    ) - (iife15_()
                    )):
                        coll11_ = coll11_.union(_dafny.Set([d_198_v39_]))
                return _dafny.Set(coll11_)
            d_154_v1_ = iife11_()
            
            d_201_v41_: _dafny.Set
            def iife16_():
                coll16_ = _dafny.Set()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(784, 480):
                    d_202_v40_: int = compr_16_
                    if ((784) <= (d_202_v40_)) and ((d_202_v40_) < (480)):
                        coll16_ = coll16_.union(_dafny.Set([default__.safeDivisionInt(d_202_v40_, len(_dafny.Map({(self).f23: _dafny.MultiSet(self.f29)})))]))
                return _dafny.Set(coll16_)
            d_201_v41_ = _dafny.Set({len(iife16_()
            ), (d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))]})
            r0 = default__.safeDivisionInt(default__.fm0(len(default__.fm5(d_201_v41_, globalState)), (p0)[default__.safeIndex(216, (p0).length(0))], globalState), ((d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))] if default__.fm1((d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))], p1, globalState) else (self).f23))
            if (p0)[default__.safeIndex(216, (p0).length(0))]:
                index34_ = default__.safeIndex(216, (p0).length(0))
                (p0)[index34_] = (p0)[default__.safeIndex(216, (p0).length(0))]
                d_203_v42_: C0
                nw32_ = C0()
                nw32_.ctor__()
                d_203_v42_ = nw32_
                d_204_v43_: _dafny.Array
                nw33_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_204_v43_ = nw33_
                d_205_v44_: _dafny.Seq
                d_205_v44_ = _dafny.SeqWithoutIsStrInference([d_187_v29_, d_189_v31_, _dafny.SeqWithoutIsStrInference([p1])])
                rhs52_ = (True) not in (((d_205_v44_)[default__.safeIndex((d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))], len(d_205_v44_))]) + (d_187_v29_))
                rhs53_ = d_204_v43_
                lhs50_ = globalState
                lhs50_.f9 = rhs52_
                d_204_v43_ = rhs53_
                (globalState).f9 = p1
                d_206_v45_: _dafny.Set
                d_206_v45_ = _dafny.Set({d_155_v2_, d_155_v2_})
                d_207_v46_: D2
                d_207_v46_ = D2_DC7(d_206_v45_)
                d_208_v47_: _dafny.Map
                d_208_v47_ = _dafny.Map({d_207_v46_: d_153_v0_})
                d_209_v48_: _dafny.MultiSet
                d_209_v48_ = _dafny.MultiSet([(d_203_v42_).fm20(d_208_v47_, (d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))], len(d_153_v0_), (d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))], globalState)])
                d_210_v49_: _dafny.Map
                d_210_v49_ = _dafny.Map({d_209_v48_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dux"))})
                d_211_v50_: D0
                d_211_v50_ = D0_DC0(d_153_v0_)
                d_212_v51_: _dafny.Seq
                d_212_v51_ = _dafny.SeqWithoutIsStrInference([default__.fm4(d_211_v50_, (self).f23, _dafny.SeqWithoutIsStrInference([(self).f28]), d_155_v2_, globalState), d_209_v48_])
                d_210_v49_ = (d_210_v49_).set((d_212_v51_)[default__.safeIndex((self).f23, len(d_212_v51_))], d_153_v0_)
            elif True:
                d_213_v52_: _dafny.Map
                d_213_v52_ = _dafny.Map({(self).f23: 275})
                d_213_v52_ = (d_213_v52_).set((d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))], (d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))])
                d_201_v41_ = default__.fm21(d_194_v35_, globalState)
                d_214_v53_: D0
                d_215_v54_: _dafny.Map
                d_216_v55_: int
                out21_: D0
                out22_: _dafny.Map
                out23_: int
                out21_, out22_, out23_ = default__.m0(p0, ((d_153_v0_).set(default__.safeIndex((self).f23, len(d_153_v0_)), d_155_v2_)) + (d_153_v0_), (self).f23, globalState)
                d_214_v53_ = out21_
                d_215_v54_ = out22_
                d_216_v55_ = out23_
                index35_ = default__.safeIndex(532, (d_190_v32_).length(0))
                (d_190_v32_)[index35_] = (self).f23
                (globalState).f9 = not((((d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))]) != (d_216_v55_)) == (((d_214_v53_).cf3) == ((d_190_v32_)[default__.safeIndex(532, (d_190_v32_).length(0))])))
        elif True:
            d_217_v56_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_217_v56_ = nw34_
            d_218_v57_: _dafny.Array
            def lambda2_(d_219_i6_):
                return self.f29

            init1_ = lambda2_
            nw35_ = _dafny.Array(None, 19)
            for i0_1_ in range(nw35_.length(0)):
                nw35_[i0_1_] = init1_(i0_1_)
            d_218_v57_ = nw35_
            d_220_v58_: _dafny.Seq
            d_220_v58_ = _dafny.SeqWithoutIsStrInference([d_218_v57_, d_218_v57_])
            index36_ = default__.safeIndex(242, (d_217_v56_).length(0))
            (d_217_v56_)[index36_] = (d_220_v58_)[default__.safeIndex((self).f23, len(d_220_v58_))]
            index37_ = default__.safeIndex(242, (d_217_v56_).length(0))
            (d_217_v56_)[index37_] = d_218_v57_
            d_221_v59_: _dafny.Array
            nw36_ = _dafny.Array(int(0), 18)
            d_221_v59_ = nw36_
            index38_ = default__.safeIndex(461, (d_221_v59_).length(0))
            (d_221_v59_)[index38_] = (self).f23
            index39_ = default__.safeIndex(461, (d_221_v59_).length(0))
            (d_221_v59_)[index39_] = (self).f23
            d_222_v60_: _dafny.Set
            d_222_v60_ = _dafny.Set({not((self).f28), (p0)[default__.safeIndex(216, (p0).length(0))], False, p1})
            index40_ = default__.safeIndex(216, (p0).length(0))
            index41_ = default__.safeIndex(461, (d_221_v59_).length(0))
            rhs54_ = default__.fm1(default__.fm0((0) - (len(d_222_v60_)), p1, globalState), (self).f28, globalState)
            rhs55_ = not((self).f28)
            rhs56_ = (self).f23
            rhs57_ = (p0)[default__.safeIndex(216, (p0).length(0))]
            rhs58_ = default__.safeModuloInt(default__.safeModuloInt((d_221_v59_)[default__.safeIndex(461, (d_221_v59_).length(0))], len(d_153_v0_)), default__.safeDivisionInt((self).f23, default__.fm0((d_221_v59_)[default__.safeIndex(461, (d_221_v59_).length(0))], p1, globalState)))
            lhs51_ = globalState
            lhs52_ = p0
            lhs53_ = default__.safeIndex(216, (p0).length(0))
            lhs54_ = globalState
            lhs55_ = d_221_v59_
            lhs56_ = default__.safeIndex(461, (d_221_v59_).length(0))
            lhs51_.f9 = rhs54_
            lhs52_[lhs53_] = rhs55_
            r0 = rhs56_
            lhs54_.f9 = rhs57_
            lhs55_[lhs56_] = rhs58_
            d_223_v61_: D3
            d_223_v61_ = D3_DC12(d_189_v31_)
            d_224_v62_: D3
            d_224_v62_ = D3_DC10(True)
            d_223_v61_ = default__.fm22(d_224_v62_, globalState)
            index42_ = default__.safeIndex(461, (d_221_v59_).length(0))
            (d_221_v59_)[index42_] = (self.f29)[default__.safeIndex((d_221_v59_)[default__.safeIndex(461, (d_221_v59_).length(0))], len(self.f29))]
        hi5_ = (self).f23
        for d_225_i7_ in range((self).f23, hi5_):
            d_226_v63_: _dafny.Array
            nw37_ = _dafny.Array(None, 1)
            nw37_[int(0)] = (p0)[default__.safeIndex(216, (p0).length(0))]
            d_226_v63_ = nw37_
            d_226_v63_ = d_226_v63_
            (globalState).f11 = d_155_v2_
            d_188_v30_ = (d_188_v30_) | (d_188_v30_)
            index43_ = default__.safeIndex(216, (p0).length(0))
            rhs59_ = default__.fm1((self).f23, (p0)[default__.safeIndex(216, (p0).length(0))], globalState)
            rhs60_ = (self).f28
            rhs61_ = (d_187_v29_)[default__.safeIndex(d_225_i7_, len(d_187_v29_))]
            lhs57_ = p0
            lhs58_ = default__.safeIndex(216, (p0).length(0))
            lhs59_ = globalState
            lhs60_ = globalState
            lhs57_[lhs58_] = rhs59_
            lhs59_.f7 = rhs60_
            lhs60_.f7 = rhs61_
        r0 = (self).f23
        d_227_v64_: _dafny.Set
        d_227_v64_ = _dafny.Set({(0) - (((self).f23) * ((self).f23)), (0) - (len(d_153_v0_))})
        r1 = d_227_v64_
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        rhs62_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoiflwc"))) + (p0)
        rhs63_ = (p0) + (p0)
        rhs64_ = 641
        rhs65_ = ((self).f23) + ((self).f23)
        lhs61_ = globalState
        lhs62_ = globalState
        lhs61_.f6 = rhs62_
        lhs62_.f6 = rhs63_
        r0 = rhs64_
        r0 = rhs65_
        d_228_v0_: _dafny.MultiSet
        d_228_v0_ = _dafny.MultiSet([(self).f28])
        d_229_v1_: _dafny.Map
        d_229_v1_ = _dafny.Map({(d_228_v0_).issubset(_dafny.MultiSet([(self).f28])): (0) - ((self).f23)})
        d_229_v1_ = (d_229_v1_).set(((self).f28) or ((self).f28), 922)
        (globalState).f14 = (0) - ((self).f23)
        d_230_v2_: _dafny.MultiSet
        d_230_v2_ = _dafny.MultiSet([(self).f23])
        d_231_v3_: _dafny.Seq
        d_231_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({308: d_230_v2_})])
        d_232_v4_: _dafny.Map
        d_232_v4_ = _dafny.Map({(self).f23: d_230_v2_})
        d_233_v5_: _dafny.Seq
        d_233_v5_ = _dafny.SeqWithoutIsStrInference([(d_231_v3_)[default__.safeIndex((self).f23, len(d_231_v3_))], (d_232_v4_).set(324, d_230_v2_), d_232_v4_])
        d_234_v6_: _dafny.Map
        d_234_v6_ = _dafny.Map({(self).f23: len((d_233_v5_)[default__.safeIndex((0) - ((self).f23), len(d_233_v5_))])})
        d_235_v7_: str
        d_235_v7_ = _dafny.CodePoint('n')
        d_236_v8_: D2
        d_236_v8_ = D2_DC8((self).f23, len(d_234_v6_), d_235_v7_, (self).f23, (self).f28)
        d_229_v1_ = (d_229_v1_).set((d_236_v8_).cf19, (0) - ((self).f23))
        d_237_v9_: _dafny.Array
        nw38_ = _dafny.Array(int(0), 23)
        d_237_v9_ = nw38_
        hi6_ = (self).f23
        for d_238_i0_ in range((0) - (len(_dafny.SeqWithoutIsStrInference([d_237_v9_]))), hi6_):
            (globalState).f14 = (d_238_i0_) * ((self).f23)
            if not(default__.fm1(d_238_i0_, (self).f28, globalState)):
                (globalState).f21 = not((self).f28)
                d_239_v10_: _dafny.Map
                d_239_v10_ = _dafny.Map({(self).f28: (self).f28})
                d_240_v11_: _dafny.Seq
                d_240_v11_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28])
                d_239_v10_ = (d_239_v10_).set((self).f28, (d_240_v11_)[default__.safeIndex(d_238_i0_, len(d_240_v11_))])
                d_241_v12_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_241_v12_ = nw39_
                d_242_v13_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_242_v13_ = nw40_
                index44_ = default__.safeIndex(536, (d_241_v12_).length(0))
                (d_241_v12_)[index44_] = d_242_v13_
                index45_ = default__.safeIndex(536, (d_241_v12_).length(0))
                nw41_ = _dafny.Array(_dafny.CodePoint('D'), 23)
                (d_241_v12_)[index45_] = nw41_
                r0 = (self).f23
                d_243_v14_: _dafny.Array
                def lambda3_(d_244_p0_, d_245_i0_, d_246_v7_):
                    def lambda4_(d_247_i1_):
                        return (d_244_p0_).set(default__.safeIndex(d_245_i0_, len(d_244_p0_)), d_246_v7_)

                    return lambda4_

                init2_ = lambda3_(p0, d_238_i0_, d_235_v7_)
                nw42_ = _dafny.Array(None, 20)
                for i0_2_ in range(nw42_.length(0)):
                    nw42_[i0_2_] = init2_(i0_2_)
                d_243_v14_ = nw42_
                index46_ = default__.safeIndex(690, (d_243_v14_).length(0))
                (d_243_v14_)[index46_] = ((p0).set(default__.safeIndex(d_238_i0_, len(p0)), d_235_v7_)) + (p0)
                d_248_v15_: _dafny.Seq
                d_248_v15_ = _dafny.SeqWithoutIsStrInference([((p0).set(default__.safeIndex(d_238_i0_, len(p0)), d_235_v7_)).set(default__.safeIndex(d_238_i0_, len((p0).set(default__.safeIndex(d_238_i0_, len(p0)), d_235_v7_))), d_235_v7_)])
                index47_ = default__.safeIndex(363, (d_243_v14_).length(0))
                (d_243_v14_)[index47_] = ((d_248_v15_)[default__.safeIndex(d_238_i0_, len(d_248_v15_))]) + (p0)
                d_249_v16_: _dafny.Seq
                d_249_v16_ = _dafny.SeqWithoutIsStrInference([d_228_v0_])
                d_250_v17_: _dafny.Array
                nw43_ = _dafny.Array(None, 21)
                nw43_[int(0)] = d_228_v0_
                nw43_[int(1)] = d_228_v0_
                nw43_[int(2)] = _dafny.MultiSet([(self).f28, (self).f28, (self).f28, (self).f28])
                nw43_[int(3)] = d_228_v0_
                nw43_[int(4)] = d_228_v0_
                nw43_[int(5)] = d_228_v0_
                nw43_[int(6)] = (((d_228_v0_).set((self).f28, default__.abs((self).f23))).set((self).f28, default__.abs((self).f23))) | (_dafny.MultiSet([(self).f28, True]))
                nw43_[int(7)] = d_228_v0_
                nw43_[int(8)] = d_228_v0_
                nw43_[int(9)] = d_228_v0_
                nw43_[int(10)] = d_228_v0_
                nw43_[int(11)] = (d_228_v0_).intersection(d_228_v0_)
                nw43_[int(12)] = (d_228_v0_).set((self).f28, default__.abs((0) - ((self).f23)))
                nw43_[int(13)] = (d_228_v0_) | (_dafny.MultiSet([(self).f28]))
                nw43_[int(14)] = d_228_v0_
                nw43_[int(15)] = d_228_v0_
                nw43_[int(16)] = d_228_v0_
                nw43_[int(17)] = (d_228_v0_) - (d_228_v0_)
                nw43_[int(18)] = (d_249_v16_)[default__.safeIndex(d_238_i0_, len(d_249_v16_))]
                nw43_[int(19)] = d_228_v0_
                nw43_[int(20)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f28]))).intersection(d_228_v0_)
                d_250_v17_ = nw43_
                index48_ = default__.safeIndex(314, (d_250_v17_).length(0))
                (d_250_v17_)[index48_] = d_228_v0_
                d_251_v18_: _dafny.Set
                d_251_v18_ = _dafny.Set({not ((self).f28) or (not((self).f28))})
                index49_ = default__.safeIndex(690, (d_243_v14_).length(0))
                index50_ = default__.safeIndex(363, (d_243_v14_).length(0))
                index51_ = default__.safeIndex(314, (d_250_v17_).length(0))
                rhs66_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bshdxkt"))) + (p0)
                rhs67_ = (_dafny.SeqWithoutIsStrInference([d_235_v7_ for d_252_i2_ in range(default__.abs(128))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgs")))
                rhs68_ = d_228_v0_
                rhs69_ = d_251_v18_
                rhs70_ = p0
                lhs63_ = d_243_v14_
                lhs64_ = default__.safeIndex(690, (d_243_v14_).length(0))
                lhs65_ = d_243_v14_
                lhs66_ = default__.safeIndex(363, (d_243_v14_).length(0))
                lhs67_ = d_250_v17_
                lhs68_ = default__.safeIndex(314, (d_250_v17_).length(0))
                lhs69_ = globalState
                lhs63_[lhs64_] = rhs66_
                lhs65_[lhs66_] = rhs67_
                lhs67_[lhs68_] = rhs68_
                d_251_v18_ = rhs69_
                lhs69_.f6 = rhs70_
            elif True:
                (globalState).f21 = (self).f28
                index52_ = default__.safeIndex(258, (d_237_v9_).length(0))
                (d_237_v9_)[index52_] = d_238_i0_
                index53_ = default__.safeIndex(258, (d_237_v9_).length(0))
                rhs71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "celkv"))
                rhs72_ = len(_dafny.SeqWithoutIsStrInference([(self).f28]))
                lhs70_ = globalState
                lhs71_ = d_237_v9_
                lhs72_ = default__.safeIndex(258, (d_237_v9_).length(0))
                lhs70_.f6 = rhs71_
                lhs71_[lhs72_] = rhs72_
                d_253_v19_: D0
                d_253_v19_ = D0_DC1((self).f23, (self).f28, ((d_230_v2_)[(self).f23] if ((self).f23) in (d_230_v2_) else (d_230_v2_).cardinality))
                (globalState).f9 = (d_253_v19_).cf2
                (globalState).f21 = (self).f28
                d_254_v20_: _dafny.Map
                d_254_v20_ = _dafny.Map({(self).f23: ((self).f28 if default__.fm1((0) - ((d_237_v9_)[default__.safeIndex(258, (d_237_v9_).length(0))]), (self).f28, globalState) else (self).f28)})
                d_254_v20_ = (d_254_v20_).set((self).f23, (self).f28)
            d_255_v21_: D4
            d_255_v21_ = D4_DC15(-165, d_235_v7_, default__.safeModuloInt((self.f29)[default__.safeIndex((d_230_v2_).cardinality, len(self.f29))], (self).f23))
            d_255_v21_ = d_255_v21_
            d_256_v22_: _dafny.Array
            nw44_ = _dafny.Array(False, 22)
            d_256_v22_ = nw44_
            index54_ = default__.safeIndex(424, (d_256_v22_).length(0))
            (d_256_v22_)[index54_] = (self).f28
            index55_ = default__.safeIndex(424, (d_256_v22_).length(0))
            (d_256_v22_)[index55_] = (d_236_v8_).cf19
        d_257_i3_: int
        d_257_i3_ = 0
        with _dafny.label("3"):
            while not(((628) != ((D1_DC5((self).f23, (self).f28, (self).f23, (self).f23)).cf9) if (self).f28 else (self).f28)):
                with _dafny.c_label("3"):
                    if (d_257_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_257_i3_ = (d_257_i3_) + (1)
                    d_258_v23_: _dafny.Seq
                    d_258_v23_ = _dafny.SeqWithoutIsStrInference([(self).f28])
                    d_259_v24_: D3
                    d_259_v24_ = D3_DC10((d_258_v23_)[default__.safeIndex((self).f23, len(d_258_v23_))])
                    d_260_v25_: _dafny.Set
                    d_260_v25_ = _dafny.Set({(self).f28})
                    d_261_v26_: _dafny.Map
                    d_261_v26_ = _dafny.Map({329: (self).f28})
                    d_262_v27_: _dafny.Map
                    d_262_v27_ = _dafny.Map({(self).f28: ((d_261_v26_)[(self).f23] if ((self).f23) in (d_261_v26_) else (self).f28)})
                    d_263_v28_: _dafny.Array
                    nw45_ = _dafny.Array(None, 29)
                    nw45_[int(0)] = (self).f28
                    nw45_[int(1)] = default__.fm1((self).f23, (self).f28, globalState)
                    nw45_[int(2)] = (self).f28
                    nw45_[int(3)] = ((self).f28 if (self).f28 else (self).f28)
                    nw45_[int(4)] = (self).f28
                    nw45_[int(5)] = (self).f28
                    nw45_[int(6)] = (d_259_v24_).cf21
                    nw45_[int(7)] = (self).f28
                    nw45_[int(8)] = (self).f28
                    nw45_[int(9)] = (self).f28
                    nw45_[int(10)] = (self).f28
                    nw45_[int(11)] = (self).f28
                    nw45_[int(12)] = not ((d_258_v23_)[default__.safeIndex((self).f23, len(d_258_v23_))]) or (default__.fm1((self).f23, (self).f28, globalState))
                    nw45_[int(13)] = ((self).f28 if (self).f28 else (self).f28)
                    nw45_[int(14)] = (d_260_v25_) == (_dafny.Set({(self).f28}))
                    nw45_[int(15)] = (self).f28
                    nw45_[int(16)] = (self).f28
                    nw45_[int(17)] = (self).f28
                    nw45_[int(18)] = (self).f28
                    nw45_[int(19)] = ((d_262_v27_)[(self).f28] if ((self).f28) in (d_262_v27_) else (self).f28)
                    nw45_[int(20)] = ((self).f23) not in ((self.f29).set(default__.safeIndex((self).f23, len(self.f29)), 128))
                    nw45_[int(21)] = not(((0) - ((self).f23)) < ((self).f23))
                    nw45_[int(22)] = (self).f28
                    nw45_[int(23)] = (self).f28
                    nw45_[int(24)] = (((d_229_v1_)[(self).f28] if ((self).f28) in (d_229_v1_) else (self).f23)) > (-581)
                    nw45_[int(25)] = (self).f28
                    nw45_[int(26)] = (d_258_v23_)[default__.safeIndex((self).f23, len(d_258_v23_))]
                    nw45_[int(27)] = not((self).f28)
                    nw45_[int(28)] = ((self).f23) < ((self).f23)
                    d_263_v28_ = nw45_
                    index56_ = default__.safeIndex(547, (d_263_v28_).length(0))
                    (d_263_v28_)[index56_] = ((self).f28 if default__.fm1((self).f23, (self).f28, globalState) else (self).f28)
                    d_264_v29_: _dafny.MultiSet
                    d_264_v29_ = _dafny.MultiSet([d_237_v9_, d_237_v9_, d_237_v9_])
                    index57_ = default__.safeIndex(547, (d_263_v28_).length(0))
                    rhs73_ = (_dafny.MultiSet([d_237_v9_, d_237_v9_])).issubset(d_264_v29_)
                    rhs74_ = (self).f23
                    rhs75_ = default__.fm1((self).f23, (self).f28, globalState)
                    rhs76_ = not ((self).f28) or ((self).f28)
                    rhs77_ = default__.fm1((self).f23, (self).f28, globalState)
                    lhs73_ = d_263_v28_
                    lhs74_ = default__.safeIndex(547, (d_263_v28_).length(0))
                    lhs75_ = globalState
                    lhs76_ = globalState
                    lhs77_ = globalState
                    lhs73_[lhs74_] = rhs73_
                    r0 = rhs74_
                    lhs75_.f21 = rhs75_
                    lhs76_.f7 = rhs76_
                    lhs77_.f21 = rhs77_
                    d_265_v30_: _dafny.Map
                    d_265_v30_ = _dafny.Map({(d_258_v23_ if not(True) else _dafny.SeqWithoutIsStrInference([(self).f28])): (self).f28})
                    d_266_v31_: _dafny.Map
                    d_266_v31_ = _dafny.Map({d_235_v7_: p0})
                    d_265_v30_ = (d_265_v30_).set(d_258_v23_, not ((d_263_v28_)[default__.safeIndex(547, (d_263_v28_).length(0))]) or ((d_258_v23_)[default__.safeIndex(len(d_266_v31_), len(d_258_v23_))]))
                    d_267_v32_: C0
                    nw46_ = C0()
                    nw46_.ctor__()
                    d_267_v32_ = nw46_
                    d_268_v33_: C0
                    nw47_ = C0()
                    nw47_.ctor__()
                    d_268_v33_ = nw47_
                    pass
            pass
        r0 = (((self).f23) - ((self).f23)) + (712)
        r1 = (self.f29) + ((self.f29 if default__.fm1((self).f23, not((self).f28), globalState) else self.f29))
        return r0, r1


class C2(T1, T0):
    def  __init__(self):
        self._f29: _dafny.Seq = _dafny.Seq({})
        self._f23: int = int(0)
        self._f28: bool = False
        self._f31: _dafny.Map = _dafny.Map({})
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f23(self):
        return self._f23
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f31, f32, f28, f29, f23):
        (self)._f31 = f31
        (self)._f32 = f32
        (self)._f28 = f28
        (self).f29 = f29
        (self)._f23 = f23

    def fm18(self, p0, p1, p2, globalState):
        return (0) - ((((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_269_i0_ in range(default__.abs(779))])) if (self).f28 else _dafny.MultiSet([_dafny.CodePoint('k'), _dafny.CodePoint('g')]))).intersection((_dafny.MultiSet([_dafny.CodePoint('s')])).intersection(_dafny.MultiSet([_dafny.CodePoint('n')])))).cardinality)

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_270_v0_: _dafny.Seq
        d_270_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rovoqj"))
        d_271_v1_: _dafny.Map
        d_271_v1_ = _dafny.Map({d_270_v0_: True})
        d_272_v2_: _dafny.Map
        d_272_v2_ = _dafny.Map({(d_270_v0_) + (d_270_v0_): (d_271_v1_) | (d_271_v1_)})
        d_272_v2_ = (d_272_v2_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_273_i0_ in range(default__.abs(-392))]), (d_271_v1_) | (d_271_v1_))
        d_274_v3_: _dafny.Array
        def lambda5_(d_275_p1_):
            def lambda6_(d_276_i1_):
                return (_dafny.SeqWithoutIsStrInference([(self).f32, d_275_p1_])) + (_dafny.SeqWithoutIsStrInference([d_275_p1_]))

            return lambda6_

        init3_ = lambda5_(p1)
        nw48_ = _dafny.Array(None, 26)
        for i0_3_ in range(nw48_.length(0)):
            nw48_[i0_3_] = init3_(i0_3_)
        d_274_v3_ = nw48_
        d_274_v3_ = d_274_v3_
        d_277_i2_: int
        d_277_i2_ = 0
        with _dafny.label("4"):
            while ((self).f23) <= ((self).f23):
                with _dafny.c_label("4"):
                    if (d_277_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_277_i2_ = (d_277_i2_) + (1)
                    d_278_v4_: _dafny.MultiSet
                    d_278_v4_ = _dafny.MultiSet([d_270_v0_, d_270_v0_])
                    d_279_v5_: T0
                    nw49_ = C1()
                    nw49_.ctor__(False, self.f29, len(d_270_v0_))
                    d_279_v5_ = nw49_
                    d_280_v6_: _dafny.MultiSet
                    d_280_v6_ = _dafny.MultiSet([d_279_v5_, d_279_v5_, d_279_v5_])
                    d_281_v7_: str
                    d_281_v7_ = _dafny.CodePoint('f')
                    d_282_v8_: _dafny.Array
                    nw50_ = _dafny.Array(None, 11)
                    nw50_[int(0)] = (self).f23
                    nw50_[int(1)] = (0) - (((d_278_v4_)[d_270_v0_] if (d_270_v0_) in (d_278_v4_) else (self).f23))
                    nw50_[int(2)] = len(_dafny.Set({-334, (d_280_v6_).cardinality}))
                    nw50_[int(3)] = 118
                    nw50_[int(4)] = (d_279_v5_).f23
                    nw50_[int(5)] = (self.f29)[default__.safeIndex(248, len(self.f29))]
                    nw50_[int(6)] = (self).f23
                    nw50_[int(7)] = (self).f23
                    nw50_[int(8)] = (self).f23
                    nw50_[int(9)] = (self).f23
                    nw50_[int(10)] = len(_dafny.SeqWithoutIsStrInference([d_281_v7_]))
                    d_282_v8_ = nw50_
                    d_283_v9_: _dafny.Seq
                    d_283_v9_ = _dafny.SeqWithoutIsStrInference([d_282_v8_, d_282_v8_, d_282_v8_, d_282_v8_, d_282_v8_])
                    d_284_v10_: _dafny.Map
                    d_284_v10_ = _dafny.Map({d_283_v9_: (d_279_v5_).f23})
                    d_284_v10_ = (d_284_v10_).set(((_dafny.SeqWithoutIsStrInference([d_282_v8_])) + (d_283_v9_)).set(default__.safeIndex((d_279_v5_).f23, len((_dafny.SeqWithoutIsStrInference([d_282_v8_])) + (d_283_v9_))), d_282_v8_), (self).f23)
                    d_285_v11_: _dafny.Array
                    nw51_ = _dafny.Array(_dafny.Seq({}), 16)
                    d_285_v11_ = nw51_
                    d_286_v12_: C0
                    nw52_ = C0()
                    nw52_.ctor__()
                    d_286_v12_ = nw52_
                    d_287_v13_: _dafny.Seq
                    d_287_v13_ = _dafny.SeqWithoutIsStrInference([d_286_v12_])
                    index58_ = default__.safeIndex(624, (d_285_v11_).length(0))
                    (d_285_v11_)[index58_] = d_287_v13_
                    index59_ = default__.safeIndex(624, (d_285_v11_).length(0))
                    (d_285_v11_)[index59_] = (d_287_v13_ if (self).f28 else d_287_v13_)
                    d_288_v14_: D1
                    d_288_v14_ = D1_DC4((self).f31, p1, True)
                    if (d_288_v14_).cf8:
                        d_289_v15_: D1
                        d_289_v15_ = D1_DC3(d_283_v9_)
                        d_289_v15_ = d_289_v15_
                        (globalState).f9 = p1
                        d_290_v16_: _dafny.MultiSet
                        d_290_v16_ = _dafny.MultiSet([(self).f28])
                        (globalState).f9 = ((d_290_v16_).issubset(d_290_v16_)) or (p1)
                        rhs78_ = not((default__.safeModuloInt((d_279_v5_).f23, (d_279_v5_).f23)) <= ((d_279_v5_).f23))
                        rhs79_ = (0) - ((self).f23)
                        rhs80_ = (d_279_v5_).f23
                        lhs78_ = globalState
                        lhs79_ = globalState
                        lhs80_ = globalState
                        lhs78_.f7 = rhs78_
                        lhs79_.f8 = rhs79_
                        lhs80_.f8 = rhs80_
                        (globalState).f7 = ((self).f32) == ((d_270_v0_) < ((d_270_v0_).set(default__.safeIndex((d_279_v5_).f23, len(d_270_v0_)), d_281_v7_)))
                    elif True:
                        d_291_v18_: _dafny.Seq
                        d_291_v18_ = _dafny.SeqWithoutIsStrInference([d_270_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tocc"))])
                        d_292_v19_: _dafny.Map
                        def iife17_():
                            coll17_ = _dafny.Map()
                            compr_17_: _dafny.Seq
                            for compr_17_ in ((d_291_v18_).set(default__.safeIndex(474, len(d_291_v18_)), d_270_v0_)).Elements:
                                d_293_v17_: _dafny.Seq = compr_17_
                                if (d_293_v17_) in ((d_291_v18_).set(default__.safeIndex(474, len(d_291_v18_)), d_270_v0_)):
                                    coll17_[d_293_v17_] = (_dafny.MultiSet([(self).f28])).cardinality
                            return _dafny.Map(coll17_)
                        d_292_v19_ = _dafny.Map({default__.safeDivisionInt((d_279_v5_).f23, (d_279_v5_).f23): iife17_()
                        })
                        d_294_v20_: _dafny.Map
                        d_294_v20_ = _dafny.Map({d_270_v0_: (0) - ((d_279_v5_).f23)})
                        d_292_v19_ = (d_292_v19_).set((self).f23, d_294_v20_)
                        d_295_v21_: _dafny.Map
                        d_295_v21_ = _dafny.Map({default__.fm0((d_279_v5_).f23, p1, globalState): p1})
                        d_295_v21_ = (d_295_v21_).set(((self).f23) * ((d_279_v5_).f23), default__.fm1((d_279_v5_).f23, (self).f28, globalState))
                        (globalState).f14 = (default__.fm0((d_279_v5_).f23, p1, globalState)) * ((-677) * ((d_279_v5_).f23))
                        d_296_v22_: _dafny.Map
                        d_296_v22_ = _dafny.Map({(self).f32: 368})
                        d_296_v22_ = (d_296_v22_).set(p1, (len(default__.fm2(globalState))) - ((d_279_v5_).f23))
                        rhs81_ = ((d_270_v0_) + (_dafny.SeqWithoutIsStrInference([d_281_v7_ for d_297_i3_ in range(default__.abs(-846))]))) + (d_270_v0_)
                        rhs82_ = (self).f23
                        lhs81_ = globalState
                        d_270_v0_ = rhs81_
                        lhs81_.f14 = rhs82_
                    d_298_v23_: _dafny.Array
                    nw53_ = _dafny.Array(D0.default()(), 2)
                    d_298_v23_ = nw53_
                    d_299_v24_: D0
                    d_299_v24_ = D0_DC1((self).f23, (self).f32, (self).f23)
                    index60_ = default__.safeIndex(326, (d_298_v23_).length(0))
                    (d_298_v23_)[index60_] = d_299_v24_
                    index61_ = default__.safeIndex(326, (d_298_v23_).length(0))
                    (d_298_v23_)[index61_] = d_299_v24_
                    pass
            pass
        d_300_v25_: _dafny.Set
        d_300_v25_ = _dafny.Set({p1})
        d_300_v25_ = d_300_v25_
        d_301_i4_: int
        d_301_i4_ = 0
        with _dafny.label("5"):
            while (self).f28:
                with _dafny.c_label("5"):
                    if (d_301_i4_) >= (100):
                        raise _dafny.Break("5")
                    d_301_i4_ = (d_301_i4_) + (1)
                    d_302_v26_: _dafny.Seq
                    d_302_v26_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])
                    d_303_v27_: _dafny.MultiSet
                    d_303_v27_ = _dafny.MultiSet([default__.fm0((self).f23, p1, globalState), (self).f23])
                    d_304_v28_: _dafny.Seq
                    d_304_v28_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_305_v29_: _dafny.MultiSet
                    d_305_v29_ = _dafny.MultiSet([d_270_v0_])
                    d_306_v30_: _dafny.Array
                    nw54_ = _dafny.Array(None, 26)
                    nw54_[int(0)] = default__.fm0((self).f23, (self).f32, globalState)
                    nw54_[int(1)] = (self).f23
                    nw54_[int(2)] = len((d_270_v0_ if (d_302_v26_)[default__.safeIndex((self).f23, len(d_302_v26_))] else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_307_i5_ in range(default__.abs(11))])))
                    nw54_[int(3)] = (self).f23
                    nw54_[int(4)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpcxes")))
                    nw54_[int(5)] = (self).f23
                    nw54_[int(6)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_308_i6_ in range(default__.abs(481))])) + (d_270_v0_))
                    nw54_[int(7)] = (self).f23
                    nw54_[int(8)] = (self).f23
                    nw54_[int(9)] = (self).f23
                    nw54_[int(10)] = default__.safeDivisionInt((self).f23, (self).f23)
                    nw54_[int(11)] = (self).f23
                    nw54_[int(12)] = ((self).f23 if p1 else 222)
                    nw54_[int(13)] = (0) - ((self).f23)
                    nw54_[int(14)] = 583
                    nw54_[int(15)] = (self).f23
                    nw54_[int(16)] = (len(d_270_v0_)) * (((_dafny.MultiSet([p1])).set(p1, default__.abs((self).f23))).cardinality)
                    nw54_[int(17)] = (self).f23
                    nw54_[int(18)] = (self).f23
                    nw54_[int(19)] = ((self).f23) * ((self).f23)
                    nw54_[int(20)] = len(d_300_v25_)
                    nw54_[int(21)] = (self).f23
                    nw54_[int(22)] = ((d_303_v27_)[len(d_304_v28_)] if (len(d_304_v28_)) in (d_303_v27_) else (self).f23)
                    nw54_[int(23)] = (d_305_v29_).cardinality
                    nw54_[int(24)] = (self).f23
                    nw54_[int(25)] = (self).f23
                    d_306_v30_ = nw54_
                    d_306_v30_ = d_306_v30_
                    d_309_v31_: _dafny.MultiSet
                    d_309_v31_ = _dafny.MultiSet([(self).f32, False])
                    d_310_v32_: _dafny.Seq
                    d_310_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thg")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suocs"))).set(default__.safeIndex((d_309_v31_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suocs")))), _dafny.CodePoint('p')), _dafny.SeqWithoutIsStrInference([(D4_DC15((self).f23, _dafny.CodePoint('f'), len(d_270_v0_))).cf27 for d_311_i7_ in range(default__.abs(247))]), d_270_v0_, d_270_v0_])
                    d_312_v33_: _dafny.Seq
                    d_312_v33_ = _dafny.SeqWithoutIsStrInference([d_302_v26_])
                    d_313_v34_: _dafny.Map
                    d_313_v34_ = _dafny.Map({(d_310_v32_)[default__.safeIndex(default__.fm0(348, not((self).f28), globalState), len(d_310_v32_))]: d_312_v33_})
                    d_313_v34_ = (d_313_v34_).set(d_270_v0_, d_312_v33_)
                    d_314_v35_: _dafny.Map
                    d_314_v35_ = _dafny.Map({(self).f23: d_306_v30_})
                    d_315_v36_: _dafny.Array
                    nw55_ = _dafny.Array(None, 13)
                    nw55_[int(0)] = (d_306_v30_ if (self).f32 else d_306_v30_)
                    nw55_[int(1)] = d_306_v30_
                    nw55_[int(2)] = d_306_v30_
                    nw55_[int(3)] = d_306_v30_
                    nw55_[int(4)] = d_306_v30_
                    nw55_[int(5)] = ((d_314_v35_)[(self).f23] if ((self).f23) in (d_314_v35_) else d_306_v30_)
                    nw55_[int(6)] = d_306_v30_
                    nw55_[int(7)] = d_306_v30_
                    nw55_[int(8)] = d_306_v30_
                    nw55_[int(9)] = d_306_v30_
                    nw55_[int(10)] = d_306_v30_
                    nw55_[int(11)] = d_306_v30_
                    nw55_[int(12)] = d_306_v30_
                    d_315_v36_ = nw55_
                    index62_ = default__.safeIndex(502, (d_315_v36_).length(0))
                    (d_315_v36_)[index62_] = d_306_v30_
                    index63_ = default__.safeIndex(502, (d_315_v36_).length(0))
                    (d_315_v36_)[index63_] = d_306_v30_
                    d_316_v37_: C0
                    nw56_ = C0()
                    nw56_.ctor__()
                    d_316_v37_ = nw56_
                    pass
            pass
        d_317_v38_: _dafny.Seq
        d_317_v38_ = _dafny.SeqWithoutIsStrInference([(self).f32])
        d_318_v39_: _dafny.MultiSet
        d_318_v39_ = _dafny.MultiSet([(self).f23])
        d_317_v38_ = ((d_317_v38_).set(default__.safeIndex(default__.fm0(len(_dafny.Map({(d_318_v39_).cardinality: _dafny.MultiSet(d_317_v38_)})), not(p1), globalState), len(d_317_v38_)), default__.fm1((self).f23, (self).f32, globalState))) + (d_317_v38_)
        r0 = (self).f23
        d_319_v40_: _dafny.Set
        d_319_v40_ = _dafny.Set({(self).f23, len(self.f29)})
        r1 = (d_319_v40_) | (d_319_v40_)
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_320_v0_: _dafny.Array
        nw57_ = _dafny.Array(False, 29)
        d_320_v0_ = nw57_
        d_321_v1_: _dafny.Seq
        d_321_v1_ = _dafny.SeqWithoutIsStrInference([p0, p0, (p0) + (p0)])
        d_322_v2_: D0
        d_322_v2_ = D0_DC0(p0)
        d_323_v3_: D1
        d_323_v3_ = D1_DC5((self).f23, (self).f32, (self).f23, (self).f23)
        pat_let_tv12_ = p0
        pat_let_tv13_ = p0
        pat_let_tv14_ = p0
        pat_let_tv15_ = p0
        pat_let_tv16_ = p0
        pat_let_tv17_ = p0
        pat_let_tv18_ = p0
        pat_let_tv19_ = p0
        pat_let_tv20_ = d_321_v1_
        pat_let_tv21_ = p0
        pat_let_tv22_ = p0
        pat_let_tv23_ = p0
        pat_let_tv24_ = p0
        pat_let_tv25_ = d_321_v1_
        def lambda7_(source1_):
            if source1_.is_DC1:
                d_324___mcc_h0_ = source1_.cf1
                d_325___mcc_h1_ = source1_.cf2
                d_326___mcc_h2_ = source1_.cf3
                d_327_cf3_ = d_326___mcc_h2_
                d_328_cf2_ = d_325___mcc_h1_
                d_329_cf1_ = d_324___mcc_h0_
                if d_328_cf2_:
                    return d_328_cf2_
                elif True:
                    return (self).f32
            elif source1_.is_DC0:
                d_330___mcc_h3_ = source1_.cf0
                d_331_cf0_ = d_330___mcc_h3_
                return (self).f28
            elif True:
                d_332___mcc_h4_ = source1_.cf4
                d_333_cf4_ = d_332___mcc_h4_
                return ((pat_let_tv12_).set(default__.safeIndex((self).f23, len(pat_let_tv13_)), (D2_DC8((self).f23, (self).f23, _dafny.CodePoint('o'), len(pat_let_tv14_), (self).f28)).cf17)) != (((pat_let_tv15_).set(default__.safeIndex((self).f23, len(pat_let_tv16_)), _dafny.CodePoint('i'))).set(default__.safeIndex((self).f23, len((pat_let_tv17_).set(default__.safeIndex((self).f23, len(pat_let_tv18_)), _dafny.CodePoint('i')))), _dafny.CodePoint('b')))

        def lambda8_(source2_):
            if source2_.is_DC4:
                d_334___mcc_h5_ = source2_.cf6
                d_335___mcc_h6_ = source2_.cf7
                d_336___mcc_h7_ = source2_.cf8
                d_337_cf8_ = d_336___mcc_h7_
                d_338_cf7_ = d_335___mcc_h6_
                d_339_cf6_ = d_334___mcc_h5_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv19_])
            elif source2_.is_DC5:
                d_340___mcc_h8_ = source2_.cf9
                d_341___mcc_h9_ = source2_.cf10
                d_342___mcc_h10_ = source2_.cf11
                d_343___mcc_h11_ = source2_.cf12
                d_344_cf12_ = d_343___mcc_h11_
                d_345_cf11_ = d_342___mcc_h10_
                d_346_cf10_ = d_341___mcc_h9_
                d_347_cf9_ = d_340___mcc_h8_
                return ((pat_let_tv20_)) + (_dafny.SeqWithoutIsStrInference([pat_let_tv21_]))
            elif source2_.is_DC6:
                d_348___mcc_h12_ = source2_.cf13
                d_349_cf13_ = d_348___mcc_h12_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv22_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_350_i0_ in range(default__.abs(-44))]), pat_let_tv23_, pat_let_tv24_])
            elif True:
                d_351___mcc_h13_ = source2_.cf5
                d_352_cf5_ = d_351___mcc_h13_
                return pat_let_tv25_

        rhs83_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcapto"))) == (p0)
        rhs84_ = d_320_v0_
        rhs85_ = lambda7_(d_322_v2_)
        rhs86_ = lambda8_(d_323_v3_)
        lhs82_ = globalState
        lhs83_ = globalState
        lhs82_.f7 = rhs83_
        d_320_v0_ = rhs84_
        lhs83_.f9 = rhs85_
        d_321_v1_ = rhs86_
        d_353_v4_: D1
        d_353_v4_ = D1_DC6(d_320_v0_)
        d_354_v5_: _dafny.MultiSet
        d_354_v5_ = _dafny.MultiSet([(d_353_v4_).cf13])
        d_354_v5_ = (d_354_v5_) | (d_354_v5_)
        (globalState).f21 = not ((self).f32) or ((self).f28)
        hi7_ = (self).f23
        for d_355_i1_ in range(((self).f23) - ((self).f23), hi7_):
            d_356_v6_: _dafny.MultiSet
            d_356_v6_ = _dafny.MultiSet([(self).f23])
            if (d_355_i1_) in ((d_356_v6_) - (d_356_v6_)):
                (globalState).f7 = True
                d_357_v7_: _dafny.Map
                d_357_v7_ = _dafny.Map({(0) - ((self).f23): (self).f32})
                d_357_v7_ = (d_357_v7_ if not(False) else d_357_v7_)
                d_358_v8_: _dafny.Map
                d_358_v8_ = _dafny.Map({d_355_i1_: (0) - (((self).f23) * ((self).f23))})
                d_358_v8_ = (d_358_v8_).set(d_355_i1_, len((self).f31))
                (globalState).f8 = len(((p0) + ((p0) + (p0))).set(default__.safeIndex((d_355_i1_) - (((d_354_v5_).set(d_320_v0_, default__.abs((self).f23))).cardinality), len((p0) + ((p0) + (p0)))), _dafny.CodePoint('b')))
                d_359_v9_: _dafny.MultiSet
                d_359_v9_ = _dafny.MultiSet([(self).f32, (self).f28])
                d_360_v10_: _dafny.Seq
                d_360_v10_ = _dafny.SeqWithoutIsStrInference([((d_357_v7_)[(self).f23] if ((self).f23) in (d_357_v7_) else (self).f28), (self).f32])
                d_361_v11_: _dafny.Map
                d_361_v11_ = _dafny.Map({not((self).f28): (self).f28})
                d_362_v12_: _dafny.Seq
                d_362_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f28: (self).f28}), d_361_v11_])
                d_363_v13_: _dafny.Array
                nw58_ = _dafny.Array(None, 24)
                nw58_[int(0)] = default__.fm0(d_355_i1_, (self).f32, globalState)
                nw58_[int(1)] = (self).f23
                nw58_[int(2)] = d_355_i1_
                nw58_[int(3)] = d_355_i1_
                nw58_[int(4)] = d_355_i1_
                nw58_[int(5)] = (self).f23
                nw58_[int(6)] = len(d_360_v10_)
                nw58_[int(7)] = 362
                nw58_[int(8)] = d_355_i1_
                nw58_[int(9)] = len((d_362_v12_)[default__.safeIndex(d_355_i1_, len(d_362_v12_))])
                nw58_[int(10)] = len(p0)
                nw58_[int(11)] = (self).f23
                nw58_[int(12)] = ((d_359_v9_).set(True, default__.abs(d_355_i1_))).cardinality
                nw58_[int(13)] = default__.fm0((self).f23, (self).f28, globalState)
                nw58_[int(14)] = (self).f23
                nw58_[int(15)] = 384
                nw58_[int(16)] = len(p0)
                nw58_[int(17)] = d_355_i1_
                nw58_[int(18)] = 501
                nw58_[int(19)] = (d_359_v9_).cardinality
                nw58_[int(20)] = (self).f23
                nw58_[int(21)] = len(d_360_v10_)
                nw58_[int(22)] = d_355_i1_
                nw58_[int(23)] = d_355_i1_
                d_363_v13_ = nw58_
                d_364_v14_: C1
                nw59_ = C1()
                nw59_.ctor__((self).f32, self.f29, ((d_359_v9_)[not((self).f28)] if (not((self).f28)) in (d_359_v9_) else (0) - (len(_dafny.Map({(self).f32: d_363_v13_})))))
                d_364_v14_ = nw59_
            elif True:
                d_365_v15_: _dafny.Array
                def lambda9_(d_366_i2_):
                    return default__.safeModuloInt(d_366_i2_, (0) - ((self).f23))

                init4_ = lambda9_
                nw60_ = _dafny.Array(None, 27)
                for i0_4_ in range(nw60_.length(0)):
                    nw60_[i0_4_] = init4_(i0_4_)
                d_365_v15_ = nw60_
                index64_ = default__.safeIndex(445, (d_365_v15_).length(0))
                (d_365_v15_)[index64_] = d_355_i1_
                index65_ = default__.safeIndex(445, (d_365_v15_).length(0))
                (d_365_v15_)[index65_] = (self).f23
                (globalState).f21 = ((self).f32) or (not((self).f32))
                (globalState).f7 = ((p0) + (p0)) != (default__.fm2(globalState))
                d_367_v16_: _dafny.Map
                d_367_v16_ = _dafny.Map({((self).f32) == ((self).f28): d_365_v15_})
                d_367_v16_ = (d_367_v16_) | (d_367_v16_)
                d_368_v17_: _dafny.Map
                d_368_v17_ = _dafny.Map({d_355_i1_: (_dafny.MultiSet([True, (self).f32, (self).f32])).isdisjoint(_dafny.MultiSet([True]))})
                (globalState).f9 = not(((d_368_v17_)[d_355_i1_] if (d_355_i1_) in (d_368_v17_) else (self).f32))
            d_369_v18_: _dafny.Map
            d_369_v18_ = _dafny.Map({(self).f32: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_370_i3_ in range(default__.abs(-469))])})
            d_371_v19_: C1
            nw61_ = C1()
            nw61_.ctor__((self).f32, self.f29, len(((d_369_v18_)[(self).f32] if ((self).f32) in (d_369_v18_) else p0)))
            d_371_v19_ = nw61_
            d_372_v20_: _dafny.Seq
            d_372_v20_ = _dafny.SeqWithoutIsStrInference([(self).f28, True, (self).f32, (self).f32, default__.fm1(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_373_i4_ in range(default__.abs(131))])), (self).f32, globalState)])
            d_374_v21_: _dafny.Map
            d_374_v21_ = _dafny.Map({(self).f23: p0})
            d_375_v22_: _dafny.Map
            d_375_v22_ = _dafny.Map({d_372_v20_: (p0) + (((d_374_v21_)[d_355_i1_] if (d_355_i1_) in (d_374_v21_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkctqqp"))))})
            d_375_v22_ = (d_375_v22_).set(d_372_v20_, p0)
            d_376_v23_: _dafny.Array
            nw62_ = _dafny.Array(D1.default()(), 15)
            d_376_v23_ = nw62_
            index66_ = default__.safeIndex(72, (d_376_v23_).length(0))
            (d_376_v23_)[index66_] = d_323_v3_
            index67_ = default__.safeIndex(72, (d_376_v23_).length(0))
            (d_376_v23_)[index67_] = (default__.fm23(globalState) if not ((self).f32) or ((self).f32) else d_323_v3_)
        (globalState).f14 = (self).f23
        d_377_v24_: _dafny.Array
        nw63_ = _dafny.Array(D1.default()(), 11)
        d_377_v24_ = nw63_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_377_v24_).length(0)):
            d_378_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_378_i5_)) and ((d_378_i5_) < ((d_377_v24_).length(0)))):
                (d_377_v24_)[(d_378_i5_)] = D1_DC4(_dafny.Map({(self).f23: (self).f23}), (self).f28, not((self).f32))
        r0 = (self).f23
        r1 = ((_dafny.SeqWithoutIsStrInference([len(p0) for d_379_i6_ in range(default__.abs(-922))])) + (_dafny.SeqWithoutIsStrInference([(self).f23 for d_380_i7_ in range(default__.abs(454))]))).set(default__.safeIndex((self).f23, len((_dafny.SeqWithoutIsStrInference([len(p0) for d_381_i6_ in range(default__.abs(-922))])) + (_dafny.SeqWithoutIsStrInference([(self).f23 for d_382_i7_ in range(default__.abs(454))])))), (self).f23)
        return r0, r1

    def m6(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: D1 = D1.default()()
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_383_v0_: _dafny.MultiSet
        d_383_v0_ = _dafny.MultiSet([(self).f28, (self).f28, p0, True])
        d_384_v1_: _dafny.Map
        d_384_v1_ = _dafny.Map({False: (d_383_v0_).cardinality})
        d_385_v2_: _dafny.Seq
        d_385_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnnsolgl"))
        d_386_v3_: D2
        d_386_v3_ = D2_DC8((self).fm18(d_384_v1_, d_385_v2_, (self).f28, globalState), (self).f23, _dafny.CodePoint('a'), 812, False)
        d_387_v4_: D2
        d_387_v4_ = D2_DC8((self).f23, (self).f23, _dafny.CodePoint('t'), (d_386_v3_).cf16, (d_386_v3_).cf19)
        d_388_v5_: _dafny.Set
        d_388_v5_ = _dafny.Set({d_387_v4_, d_386_v3_})
        d_388_v5_ = d_388_v5_
        d_389_v6_: _dafny.Array
        nw64_ = _dafny.Array(int(0), 8)
        d_389_v6_ = nw64_
        index68_ = default__.safeIndex(669, (d_389_v6_).length(0))
        (d_389_v6_)[index68_] = 528
        index69_ = default__.safeIndex(669, (d_389_v6_).length(0))
        (d_389_v6_)[index69_] = (self).f23
        hi8_ = ((d_389_v6_)[default__.safeIndex(669, (d_389_v6_).length(0))]) - ((d_389_v6_)[default__.safeIndex(669, (d_389_v6_).length(0))])
        for d_390_i0_ in range((len(self.f29)) - ((self).f23), hi8_):
            (globalState).f8 = (d_389_v6_)[default__.safeIndex(669, (d_389_v6_).length(0))]
            d_391_v7_: str
            d_391_v7_ = _dafny.CodePoint('w')
            d_392_v8_: _dafny.Map
            d_392_v8_ = _dafny.Map({p0: d_391_v7_})
            index70_ = default__.safeIndex(669, (d_389_v6_).length(0))
            (d_389_v6_)[index70_] = (((d_384_v1_)[p0] if (p0) in (d_384_v1_) else (self).f23) if ((d_392_v8_).set((self).f32, d_391_v7_)) != (d_392_v8_) else ((self).f23) * (-826))
            d_393_v9_: C1
            nw65_ = C1()
            nw65_.ctor__((self).f28, (_dafny.SeqWithoutIsStrInference([d_390_i0_])) + (self.f29), d_390_i0_)
            d_393_v9_ = nw65_
            (self).f29 = self.f29
        d_394_v10_: _dafny.Seq
        d_394_v10_ = _dafny.SeqWithoutIsStrInference([p0, (self).f28])
        d_395_v11_: D3
        d_395_v11_ = D3_DC12(d_394_v10_)
        d_395_v11_ = (d_395_v11_ if (p0) == (p0) else d_395_v11_)
        d_396_v12_: _dafny.Array
        nw66_ = _dafny.Array(False, 25)
        d_396_v12_ = nw66_
        index71_ = default__.safeIndex(779, (d_396_v12_).length(0))
        (d_396_v12_)[index71_] = p0
        index72_ = default__.safeIndex(779, (d_396_v12_).length(0))
        (d_396_v12_)[index72_] = (self).f32
        d_397_i1_: int
        d_397_i1_ = 0
        with _dafny.label("6"):
            while (d_396_v12_)[default__.safeIndex(779, (d_396_v12_).length(0))]:
                with _dafny.c_label("6"):
                    if (d_397_i1_) >= (100):
                        raise _dafny.Break("6")
                    d_397_i1_ = (d_397_i1_) + (1)
                    r0 = not((d_396_v12_)[default__.safeIndex(779, (d_396_v12_).length(0))])
                    d_398_v13_: _dafny.Seq
                    d_398_v13_ = _dafny.SeqWithoutIsStrInference([d_385_v2_, d_385_v2_])
                    (globalState).f21 = (d_385_v2_) not in (d_398_v13_)
                    d_399_v14_: _dafny.Seq
                    d_399_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({False: len((d_384_v1_))})])
                    (globalState).f14 = (self).fm18((d_399_v14_)[default__.safeIndex((self).f23, len(d_399_v14_))], (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_400_i2_ in range(default__.abs(339))])) + (d_385_v2_), (d_396_v12_)[default__.safeIndex(779, (d_396_v12_).length(0))], globalState)
                    d_384_v1_ = (d_384_v1_).set((self).f32, (self).f23)
                    pass
            pass
        r0 = (self).f28
        d_401_v15_: _dafny.Map
        d_401_v15_ = _dafny.Map({((self).f23) * ((self).f23): d_396_v12_})
        r1 = d_401_v15_
        d_402_v16_: D1
        d_402_v16_ = D1_DC4((self).f31, (d_396_v12_)[default__.safeIndex(779, (d_396_v12_).length(0))], (d_396_v12_)[default__.safeIndex(779, (d_396_v12_).length(0))])
        r2 = d_402_v16_
        r3 = d_396_v12_
        return r0, r1, r2, r3

    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32

class C3(T0, T1):
    def  __init__(self):
        self._f29: _dafny.Seq = _dafny.Seq({})
        self._f23: int = int(0)
        self._f28: bool = False
        self._f30: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f23(self):
        return self._f23
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f30, f23, f28, f29):
        (self)._f30 = f30
        (self)._f23 = f23
        (self)._f28 = f28
        (self).f29 = f29

    def fm12(self, p0, p1, p2, globalState):
        return (0) - (((self).f23) + (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_403_i0_ in range(default__.abs(-585))])), 974)))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_404_v0_: _dafny.Array
        nw67_ = _dafny.Array(False, 22)
        d_404_v0_ = nw67_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_404_v0_).length(0)):
            d_405_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_405_i0_)) and ((d_405_i0_) < ((d_404_v0_).length(0)))):
                (d_404_v0_)[(d_405_i0_)] = (D2_DC8((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f23: (self).f23})]))).cardinality, (self).f23, _dafny.CodePoint('p'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nasts"))), True)).cf19
        if (default__.safeDivisionInt((self).f23, (self).f23)) < (((self).f23) * ((self).f23)):
            d_406_v1_: _dafny.Set
            d_406_v1_ = _dafny.Set({p1, p1})
            rhs87_ = (d_406_v1_).ispropersubset((d_406_v1_) - (_dafny.Set({(self).f28})))
            rhs88_ = (self).f28
            lhs84_ = globalState
            lhs85_ = globalState
            lhs84_.f21 = rhs87_
            lhs85_.f21 = rhs88_
            d_407_v2_: _dafny.Map
            d_407_v2_ = _dafny.Map({(self).f28: False})
            index73_ = default__.safeIndex(868, (p0).length(0))
            (p0)[index73_] = ((d_407_v2_)[p1] if (p1) in (d_407_v2_) else p1)
            index74_ = default__.safeIndex(868, (p0).length(0))
            (p0)[index74_] = (self).f28
            d_408_v3_: _dafny.Array
            nw68_ = _dafny.Array(_dafny.Map({}), 20)
            d_408_v3_ = nw68_
            d_409_v4_: _dafny.Map
            d_409_v4_ = _dafny.Map({(self).f23: True})
            index75_ = default__.safeIndex(243, (d_408_v3_).length(0))
            (d_408_v3_)[index75_] = d_409_v4_
            index76_ = default__.safeIndex(243, (d_408_v3_).length(0))
            (d_408_v3_)[index76_] = d_409_v4_
            index77_ = default__.safeIndex(868, (p0).length(0))
            (p0)[index77_] = ((p0)[default__.safeIndex(868, (p0).length(0))]) not in (d_407_v2_)
            d_410_v5_: D1
            d_410_v5_ = D1_DC5((self).f23, (p0)[default__.safeIndex(868, (p0).length(0))], (self).f23, (self).f23)
            d_411_v6_: _dafny.Seq
            d_411_v6_ = _dafny.SeqWithoutIsStrInference([d_410_v5_])
            d_412_v7_: _dafny.Map
            d_412_v7_ = _dafny.Map({d_411_v6_: p1})
            d_413_v8_: _dafny.Map
            d_413_v8_ = _dafny.Map({(self).f23: _dafny.Map({-225: (self).f23})})
            d_414_v9_: _dafny.MultiSet
            d_414_v9_ = _dafny.MultiSet([(self).f28])
            d_415_v10_: D3
            d_415_v10_ = D3_DC11(default__.fm0((d_414_v9_).cardinality, (self).f28, globalState))
            d_416_v11_: _dafny.Map
            d_416_v11_ = _dafny.Map({(d_414_v9_).cardinality: (d_415_v10_).cf22})
            d_417_v12_: _dafny.Set
            d_417_v12_ = _dafny.Set({(self).f23})
            d_418_v13_: D1
            d_418_v13_ = D1_DC4(d_416_v11_, ((self).f28) or (p1), (len(d_417_v12_)) != ((self).f23))
            rhs89_ = d_412_v7_
            rhs90_ = (self).f23
            rhs91_ = default__.fm13(_dafny.CodePoint('q'), default__.fm1((self).f23, (self).f28, globalState), ((self).f23 if default__.fm1((self).f23, p1, globalState) else (self).f23), ((p0)[default__.safeIndex(868, (p0).length(0))] if p1 else default__.fm1((self).f23, (p0)[default__.safeIndex(868, (p0).length(0))], globalState)), globalState)
            rhs92_ = d_418_v13_
            rhs93_ = 473
            lhs86_ = globalState
            lhs87_ = globalState
            d_412_v7_ = rhs89_
            lhs86_.f14 = rhs90_
            d_413_v8_ = rhs91_
            d_418_v13_ = rhs92_
            lhs87_.f14 = rhs93_
        elif True:
            (globalState).f14 = (self).f23
            d_419_v14_: _dafny.Set
            d_419_v14_ = _dafny.Set({p1})
            d_420_v15_: str
            d_420_v15_ = _dafny.CodePoint('h')
            d_421_v16_: _dafny.Map
            d_421_v16_ = _dafny.Map({p1: (self).f23})
            d_422_v17_: _dafny.MultiSet
            d_422_v17_ = _dafny.MultiSet([(self).f23, (self).fm12((self).f23, (self).f28, p1, globalState), (self).f23])
            d_423_v18_: _dafny.Seq
            d_423_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htcv"))
            d_424_v19_: _dafny.Array
            nw69_ = _dafny.Array(None, 28)
            nw69_[int(0)] = (self).f23
            nw69_[int(1)] = default__.safeDivisionInt((self).f23, (self).f23)
            nw69_[int(2)] = (self).f23
            nw69_[int(3)] = (self).f23
            nw69_[int(4)] = (self).f23
            nw69_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))
            nw69_[int(6)] = (len(d_419_v14_)) * ((self).f23)
            nw69_[int(7)] = (self).f23
            nw69_[int(8)] = (default__.fm14((self).f23, d_420_v15_, (self).f28, globalState)).cardinality
            nw69_[int(9)] = (self).f23
            nw69_[int(10)] = (self).f23
            nw69_[int(11)] = (self).f23
            nw69_[int(12)] = ((self).fm12((self).f23, default__.fm1((self).f23, p1, globalState), not(p1), globalState)) + (447)
            nw69_[int(13)] = (self).f23
            nw69_[int(14)] = 629
            nw69_[int(15)] = (self).fm12(default__.fm0((self).f23, p1, globalState), (self).f28, (self).f28, globalState)
            nw69_[int(16)] = (self).f23
            nw69_[int(17)] = (len(d_421_v16_)) * (((d_422_v17_)[876] if (876) in (d_422_v17_) else (self).f23))
            nw69_[int(18)] = ((self).f23) + ((self).fm12(len(d_423_v18_), (self).f28, p1, globalState))
            nw69_[int(19)] = (547) + ((self).f23)
            nw69_[int(20)] = len(d_423_v18_)
            nw69_[int(21)] = (len(self.f29) if p1 else (self).f23)
            nw69_[int(22)] = (self).f23
            nw69_[int(23)] = (self).f23
            nw69_[int(24)] = default__.safeModuloInt((self).f23, (self).f23)
            nw69_[int(25)] = ((d_422_v17_).set((self).f23, default__.abs((self).fm12((self).f23, (self).f28, (self).f28, globalState)))).cardinality
            nw69_[int(26)] = 966
            nw69_[int(27)] = (self).f23
            d_424_v19_ = nw69_
            index78_ = default__.safeIndex(726, (d_424_v19_).length(0))
            (d_424_v19_)[index78_] = (self).f23
            d_425_v20_: _dafny.Map
            d_425_v20_ = _dafny.Map({(self).f28: ((self).f23) <= (141)})
            index79_ = default__.safeIndex(726, (d_424_v19_).length(0))
            (d_424_v19_)[index79_] = (0) - (len(d_425_v20_))
            (globalState).f8 = default__.fm0((self).f23, p1, globalState)
            if ((self).f28 if (self).f28 else p1):
                d_426_v21_: _dafny.Seq
                d_426_v21_ = _dafny.SeqWithoutIsStrInference([False, p1, (self).f28])
                d_427_v22_: _dafny.MultiSet
                d_427_v22_ = _dafny.MultiSet([(d_426_v21_) + (d_426_v21_)])
                d_428_v23_: _dafny.MultiSet
                d_428_v23_ = _dafny.MultiSet([p1])
                d_429_v24_: _dafny.Map
                d_429_v24_ = _dafny.Map({(d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]: d_428_v23_})
                pat_let_tv26_ = p1
                def iife18_(_pat_let0_0):
                    def iife19_(d_430_dt__update__tmp_h0_):
                        def iife20_(_pat_let1_0):
                            def iife21_(d_431_dt__update_hcf21_h0_):
                                return D3_DC10(d_431_dt__update_hcf21_h0_)
                            return iife21_(_pat_let1_0)
                        return iife20_(pat_let_tv26_)
                    return iife19_(_pat_let0_0)
                d_427_v22_ = default__.fm15((iife18_(D3_DC10(p1))).cf21, d_429_v24_, globalState)
                d_432_v25_: _dafny.Seq
                d_432_v25_ = _dafny.SeqWithoutIsStrInference([d_404_v0_, d_404_v0_])
                d_433_v26_: _dafny.Array
                nw70_ = _dafny.Array(None, 14)
                nw70_[int(0)] = d_404_v0_
                nw70_[int(1)] = p0
                nw70_[int(2)] = p0
                nw70_[int(3)] = d_404_v0_
                nw70_[int(4)] = (p0 if (self).f28 else d_404_v0_)
                nw70_[int(5)] = d_404_v0_
                nw70_[int(6)] = (d_432_v25_)[default__.safeIndex((_dafny.MultiSet(self.f29)).cardinality, len(d_432_v25_))]
                nw70_[int(7)] = d_404_v0_
                nw70_[int(8)] = (p0 if p1 else d_404_v0_)
                nw70_[int(9)] = p0
                nw70_[int(10)] = p0
                nw70_[int(11)] = d_404_v0_
                nw70_[int(12)] = p0
                nw70_[int(13)] = p0
                d_433_v26_ = nw70_
                index80_ = default__.safeIndex(281, (d_433_v26_).length(0))
                (d_433_v26_)[index80_] = p0
                index81_ = default__.safeIndex(281, (d_433_v26_).length(0))
                (d_433_v26_)[index81_] = p0
                d_419_v14_ = d_419_v14_
                d_434_v27_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_434_v27_ = nw71_
                index82_ = default__.safeIndex(602, (d_434_v27_).length(0))
                (d_434_v27_)[index82_] = d_424_v19_
                d_435_v28_: _dafny.Map
                d_435_v28_ = _dafny.Map({(d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]: p1})
                d_436_v29_: _dafny.Map
                d_436_v29_ = _dafny.Map({d_423_v18_: (_dafny.Map({(d_426_v21_)[default__.safeIndex((d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))], len(d_426_v21_))]: (d_428_v23_).cardinality}) if p1 else _dafny.Map({((d_435_v28_)[len(d_423_v18_)] if (len(d_423_v18_)) in (d_435_v28_) else (self).f28): (d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]}))})
                index83_ = default__.safeIndex(602, (d_434_v27_).length(0))
                index84_ = default__.safeIndex(726, (d_424_v19_).length(0))
                rhs94_ = ((d_436_v29_)[d_423_v18_] if (d_423_v18_) in (d_436_v29_) else d_421_v16_)
                rhs95_ = d_424_v19_
                rhs96_ = (self).fm12((d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))], False, p1, globalState)
                rhs97_ = (self).f28
                rhs98_ = ((d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]) - ((self).f23)
                lhs88_ = d_434_v27_
                lhs89_ = default__.safeIndex(602, (d_434_v27_).length(0))
                lhs90_ = d_424_v19_
                lhs91_ = default__.safeIndex(726, (d_424_v19_).length(0))
                lhs92_ = globalState
                lhs93_ = globalState
                d_421_v16_ = rhs94_
                lhs88_[lhs89_] = rhs95_
                lhs90_[lhs91_] = rhs96_
                lhs92_.f9 = rhs97_
                lhs93_.f14 = rhs98_
                (globalState).f8 = (d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]
            elif True:
                d_437_v30_: _dafny.Map
                d_437_v30_ = _dafny.Map({d_404_v0_: ((self).f23) * (434)})
                d_437_v30_ = (d_437_v30_).set(d_404_v0_, (self).f23)
                index85_ = default__.safeIndex(726, (d_424_v19_).length(0))
                (d_424_v19_)[index85_] = (d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]
                d_438_v31_: _dafny.Seq
                d_438_v31_ = _dafny.SeqWithoutIsStrInference([(self.f29) + (self.f29), self.f29])
                d_439_v32_: _dafny.Map
                d_439_v32_ = _dafny.Map({391: self.f29})
                d_438_v31_ = (((d_438_v31_).set(default__.safeIndex(((d_422_v17_)[(self).f23] if ((self).f23) in (d_422_v17_) else (self).f23), len(d_438_v31_)), self.f29)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p1])), len((d_438_v31_).set(default__.safeIndex(((d_422_v17_)[(self).f23] if ((self).f23) in (d_422_v17_) else (self).f23), len(d_438_v31_)), self.f29))), ((d_439_v32_)[(d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]] if ((d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]) in (d_439_v32_) else _dafny.SeqWithoutIsStrInference([(d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))]])))) + ((d_438_v31_) + (d_438_v31_))
                (self).m5(globalState)
                d_440_v33_: D4
                d_440_v33_ = D4_DC14(d_424_v19_)
                d_424_v19_ = (d_440_v33_).cf25
            d_441_v34_: _dafny.Map
            d_441_v34_ = _dafny.Map({(D0_DC1((d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))], (self).f28, (d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))])).cf1: (self).f23})
            d_442_v35_: _dafny.Seq
            d_442_v35_ = _dafny.SeqWithoutIsStrInference([default__.fm16((self).f28, (self).f23, (self).f23, globalState)])
            d_441_v34_ = (d_441_v34_).set((self).f23, len(((d_442_v35_)[default__.safeIndex((d_424_v19_)[default__.safeIndex(726, (d_424_v19_).length(0))], len(d_442_v35_))]) | (_dafny.Map({len(d_423_v18_): (self).f23}))))
        (globalState).f14 = (self).f23
        d_443_v36_: _dafny.Seq
        d_443_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjvi"))
        d_444_v37_: str
        d_444_v37_ = _dafny.CodePoint('y')
        if not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ii"))) == ((d_443_v36_).set(default__.safeIndex((self).f23, len(d_443_v36_)), d_444_v37_))):
            index86_ = default__.safeIndex(956, (d_404_v0_).length(0))
            (d_404_v0_)[index86_] = (self).f28
            index87_ = default__.safeIndex(956, (d_404_v0_).length(0))
            (d_404_v0_)[index87_] = not((D3_DC10((self).f28)).cf21)
            d_445_v38_: _dafny.Array
            def lambda10_(d_446_v0_):
                def lambda11_(d_447_i1_):
                    return _dafny.MultiSet([(d_446_v0_)[default__.safeIndex(956, (d_446_v0_).length(0))]])

                return lambda11_

            init5_ = lambda10_(d_404_v0_)
            nw72_ = _dafny.Array(None, 14)
            for i0_5_ in range(nw72_.length(0)):
                nw72_[i0_5_] = init5_(i0_5_)
            d_445_v38_ = nw72_
            d_445_v38_ = d_445_v38_
            d_448_v40_: D1
            def iife22_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(57, 984):
                    d_449_v39_: int = compr_18_
                    if ((57) <= (d_449_v39_)) and ((d_449_v39_) < (984)):
                        coll18_[(d_449_v39_) + ((self).f23)] = (self).f23
                return _dafny.Map(coll18_)
            d_448_v40_ = D1_DC4(iife22_()
, (d_404_v0_)[default__.safeIndex(956, (d_404_v0_).length(0))], p1)
            index88_ = default__.safeIndex(956, (d_404_v0_).length(0))
            (d_404_v0_)[index88_] = (default__.fm1((self).f23, False, globalState)) == ((d_448_v40_).cf7)
            (globalState).f6 = (d_443_v36_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))
            (globalState).f8 = (self.f29)[default__.safeIndex((self).f23, len(self.f29))]
        elif True:
            index89_ = default__.safeIndex(296, (d_404_v0_).length(0))
            (d_404_v0_)[index89_] = p1
            index90_ = default__.safeIndex(296, (d_404_v0_).length(0))
            (d_404_v0_)[index90_] = p1
            d_450_v41_: _dafny.Array
            def lambda12_(d_451_i2_):
                return (d_451_i2_) + ((self).f23)

            init6_ = lambda12_
            nw73_ = _dafny.Array(None, 11)
            for i0_6_ in range(nw73_.length(0)):
                nw73_[i0_6_] = init6_(i0_6_)
            d_450_v41_ = nw73_
            d_452_v42_: _dafny.Seq
            d_452_v42_ = _dafny.SeqWithoutIsStrInference([(d_404_v0_)[default__.safeIndex(296, (d_404_v0_).length(0))], p1])
            index91_ = default__.safeIndex(509, (d_450_v41_).length(0))
            (d_450_v41_)[index91_] = len(d_452_v42_)
            index92_ = default__.safeIndex(509, (d_450_v41_).length(0))
            (d_450_v41_)[index92_] = (self).fm12((self).f23, (d_404_v0_)[default__.safeIndex(296, (d_404_v0_).length(0))], (d_404_v0_)[default__.safeIndex(296, (d_404_v0_).length(0))], globalState)
            d_453_v43_: _dafny.Array
            def lambda13_(d_454_v36_):
                def lambda14_(d_455_i3_):
                    return (d_454_v36_) + (d_454_v36_)

                return lambda14_

            init7_ = lambda13_(d_443_v36_)
            nw74_ = _dafny.Array(None, 18)
            for i0_7_ in range(nw74_.length(0)):
                nw74_[i0_7_] = init7_(i0_7_)
            d_453_v43_ = nw74_
            index93_ = default__.safeIndex(35, (d_453_v43_).length(0))
            (d_453_v43_)[index93_] = d_443_v36_
            d_456_v44_: _dafny.MultiSet
            d_456_v44_ = _dafny.MultiSet([d_444_v37_])
            d_457_v45_: _dafny.Map
            d_457_v45_ = _dafny.Map({True: (d_450_v41_)[default__.safeIndex(509, (d_450_v41_).length(0))]})
            d_458_v46_: _dafny.Seq
            d_458_v46_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkbs")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxecrdjl"))])
            d_459_v47_: _dafny.Set
            d_459_v47_ = _dafny.Set({True, (self).f28})
            d_460_v48_: _dafny.MultiSet
            d_460_v48_ = _dafny.MultiSet([d_459_v47_])
            d_461_v49_: D4
            d_461_v49_ = D4_DC15((d_460_v48_).cardinality, d_444_v37_, (self).f23)
            index94_ = default__.safeIndex(509, (d_450_v41_).length(0))
            index95_ = default__.safeIndex(509, (d_450_v41_).length(0))
            index96_ = default__.safeIndex(35, (d_453_v43_).length(0))
            index97_ = default__.safeIndex(509, (d_450_v41_).length(0))
            rhs99_ = (d_450_v41_)[default__.safeIndex(509, (d_450_v41_).length(0))]
            rhs100_ = default__.fm0((self).f23, (default__.fm17((self).f28, len(d_443_v36_), globalState)).ispropersubset((d_456_v44_).set(d_444_v37_, default__.abs(len(d_457_v45_)))), globalState)
            rhs101_ = (d_443_v36_) + (d_443_v36_)
            rhs102_ = (((d_461_v49_).cf28 if (d_404_v0_)[default__.safeIndex(296, (d_404_v0_).length(0))] else (self).fm12((self).f23, (self).f28, default__.fm1((self).f23, p1, globalState), globalState)) if (_dafny.SeqWithoutIsStrInference([d_443_v36_ for d_462_i4_ in range(default__.abs(39))])) < (d_458_v46_) else ((d_450_v41_)[default__.safeIndex(509, (d_450_v41_).length(0))]) - ((self).f23))
            lhs94_ = d_450_v41_
            lhs95_ = default__.safeIndex(509, (d_450_v41_).length(0))
            lhs96_ = d_450_v41_
            lhs97_ = default__.safeIndex(509, (d_450_v41_).length(0))
            lhs98_ = d_453_v43_
            lhs99_ = default__.safeIndex(35, (d_453_v43_).length(0))
            lhs100_ = d_450_v41_
            lhs101_ = default__.safeIndex(509, (d_450_v41_).length(0))
            lhs94_[lhs95_] = rhs99_
            lhs96_[lhs97_] = rhs100_
            lhs98_[lhs99_] = rhs101_
            lhs100_[lhs101_] = rhs102_
            (self).m5(globalState)
            index98_ = default__.safeIndex(296, (d_404_v0_).length(0))
            (d_404_v0_)[index98_] = ((d_450_v41_)[default__.safeIndex(509, (d_450_v41_).length(0))]) == (len(d_457_v45_))
        hi9_ = 581
        for d_463_i5_ in range((0) - ((self).f23), hi9_):
            d_464_v50_: _dafny.Map
            d_464_v50_ = _dafny.Map({p1: not(p1)})
            d_465_v51_: _dafny.Map
            d_465_v51_ = _dafny.Map({len(d_464_v50_): p0})
            d_466_v52_: _dafny.Array
            def lambda15_(d_467_i6_):
                return (d_467_i6_) + (len(_dafny.Set({self.f29})))

            init8_ = lambda15_
            nw75_ = _dafny.Array(None, 17)
            for i0_8_ in range(nw75_.length(0)):
                nw75_[i0_8_] = init8_(i0_8_)
            d_466_v52_ = nw75_
            d_468_v53_: _dafny.Map
            d_468_v53_ = _dafny.Map({((d_465_v51_)[(self).f23] if ((self).f23) in (d_465_v51_) else p0): d_466_v52_})
            d_469_v54_: D4
            d_469_v54_ = D4_DC14(((d_468_v53_)[p0] if (p0) in (d_468_v53_) else d_466_v52_))
            d_469_v54_ = D4_DC14(d_466_v52_)
            d_470_v55_: _dafny.Map
            d_470_v55_ = _dafny.Map({(self).f23: (self).f23})
            d_471_v56_: D1
            d_471_v56_ = D1_DC4(d_470_v55_, p1, p1)
            source3_ = d_471_v56_
            if source3_.is_DC4:
                d_472___mcc_h0_ = source3_.cf6
                d_473___mcc_h1_ = source3_.cf7
                d_474___mcc_h2_ = source3_.cf8
                d_475_cf8_ = d_474___mcc_h2_
                d_476_cf7_ = d_473___mcc_h1_
                d_477_cf6_ = d_472___mcc_h0_
                (globalState).f14 = (len(d_443_v36_)) * (((self).f23) * ((self).f23))
                d_478_v57_: _dafny.Map
                d_478_v57_ = _dafny.Map({(self).f28: (self).f23})
                d_479_v58_: _dafny.MultiSet
                d_479_v58_ = _dafny.MultiSet([len(d_478_v57_)])
                rhs103_ = (d_479_v58_).issubset(d_479_v58_)
                rhs104_ = d_469_v54_
                rhs105_ = d_444_v37_
                rhs106_ = (d_463_i5_) + ((self).f23)
                rhs107_ = d_469_v54_
                lhs102_ = globalState
                d_476_cf7_ = rhs103_
                d_469_v54_ = rhs104_
                d_444_v37_ = rhs105_
                lhs102_.f14 = rhs106_
                d_469_v54_ = rhs107_
                d_480_v59_: _dafny.Map
                d_480_v59_ = _dafny.Map({not (default__.fm1(631, d_475_cf8_, globalState)) or (not(d_475_cf8_)): d_464_v50_})
                d_481_v60_: _dafny.Seq
                d_481_v60_ = _dafny.SeqWithoutIsStrInference([d_464_v50_])
                d_480_v59_ = (d_480_v59_).set((self).f28, ((d_464_v50_).set(p1, d_475_cf8_)) | ((d_481_v60_)[default__.safeIndex((self).f23, len(d_481_v60_))]))
                d_482_v61_: _dafny.Map
                d_482_v61_ = _dafny.Map({_dafny.Map({d_466_v52_: _dafny.Set({d_444_v37_, d_444_v37_})}): (0) - ((self).f23)})
                d_483_v62_: _dafny.Set
                d_483_v62_ = _dafny.Set({d_444_v37_, _dafny.CodePoint('e'), _dafny.CodePoint('a'), d_444_v37_})
                d_484_v63_: _dafny.Map
                d_484_v63_ = _dafny.Map({d_466_v52_: d_483_v62_})
                d_482_v61_ = (d_482_v61_).set((d_484_v63_) | (d_484_v63_), 568)
            elif source3_.is_DC5:
                d_485___mcc_h3_ = source3_.cf9
                d_486___mcc_h4_ = source3_.cf10
                d_487___mcc_h5_ = source3_.cf11
                d_488___mcc_h6_ = source3_.cf12
                d_489_cf12_ = d_488___mcc_h6_
                d_490_cf11_ = d_487___mcc_h5_
                d_491_cf10_ = d_486___mcc_h4_
                d_492_cf9_ = d_485___mcc_h3_
                d_493_v64_: _dafny.Array
                nw76_ = _dafny.Array(None, 21)
                d_493_v64_ = nw76_
                d_494_v65_: T1
                nw77_ = C2()
                nw77_.ctor__(d_470_v55_, False, d_491_cf10_, _dafny.SeqWithoutIsStrInference([970]), (self.f29)[default__.safeIndex((0) - ((self).f23), len(self.f29))])
                d_494_v65_ = nw77_
                d_495_v66_: T1
                d_495_v66_ = d_494_v65_
                index99_ = default__.safeIndex(821, (d_493_v64_).length(0))
                (d_493_v64_)[index99_] = (d_495_v66_)
                index100_ = default__.safeIndex(821, (d_493_v64_).length(0))
                (d_493_v64_)[index100_] = d_494_v65_
                d_496_v67_: _dafny.Seq
                d_496_v67_ = _dafny.SeqWithoutIsStrInference([(self).f28, p1, d_491_cf10_, (self).f28])
                d_490_cf11_ = (d_490_cf11_) - (default__.safeDivisionInt(len(d_496_v67_), (self).f23))
                (globalState).f14 = d_489_cf12_
                index101_ = default__.safeIndex(836, (d_404_v0_).length(0))
                (d_404_v0_)[index101_] = not(True)
                index102_ = default__.safeIndex(836, (d_404_v0_).length(0))
                (d_404_v0_)[index102_] = d_491_cf10_
            elif source3_.is_DC6:
                d_497___mcc_h7_ = source3_.cf13
                d_498_cf13_ = d_497___mcc_h7_
                d_499_v68_: C1
                nw78_ = C1()
                nw78_.ctor__((self).f28, self.f29, default__.fm0((self).f23, (self).f28, globalState))
                d_499_v68_ = nw78_
                d_499_v68_ = d_499_v68_
                d_464_v50_ = (d_464_v50_).set(p1, not(p1))
                d_500_v69_: _dafny.Map
                d_500_v69_ = _dafny.Map({(self).f28: d_444_v37_})
                d_500_v69_ = (d_500_v69_).set((self).f28, d_444_v37_)
                d_501_v70_: _dafny.Map
                d_501_v70_ = _dafny.Map({False: d_464_v50_})
                (globalState).f9 = ((self).f28) in (d_501_v70_)
            elif True:
                d_502___mcc_h8_ = source3_.cf5
                d_503_cf5_ = d_502___mcc_h8_
                rhs108_ = p1
                rhs109_ = 408
                lhs103_ = globalState
                lhs103_.f7 = rhs108_
                r0 = rhs109_
                index103_ = default__.safeIndex(936, (d_466_v52_).length(0))
                (d_466_v52_)[index103_] = (self).f23
                index104_ = default__.safeIndex(936, (d_466_v52_).length(0))
                (d_466_v52_)[index104_] = (self).fm12((self).f23, p1, (self).f28, globalState)
                d_504_v71_: D0
                d_505_v72_: _dafny.Map
                d_506_v73_: int
                out24_: D0
                out25_: _dafny.Map
                out26_: int
                out24_, out25_, out26_ = default__.m0((p0 if (self).f28 else d_404_v0_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nplonay")), d_463_i5_, globalState)
                d_504_v71_ = out24_
                d_505_v72_ = out25_
                d_506_v73_ = out26_
                (globalState).f14 = (d_463_i5_) + (d_506_v73_)
            (globalState).f14 = (d_463_i5_) - (len(d_443_v36_))
            index105_ = default__.safeIndex(107, (d_404_v0_).length(0))
            (d_404_v0_)[index105_] = p1
            index106_ = default__.safeIndex(107, (d_404_v0_).length(0))
            (d_404_v0_)[index106_] = default__.fm1((self).f23, not(p1), globalState)
        d_507_v74_: _dafny.MultiSet
        d_507_v74_ = _dafny.MultiSet([not(True)])
        d_508_v75_: _dafny.Set
        d_508_v75_ = _dafny.Set({(self).f23, default__.fm0((self).f23, (self).f28, globalState)})
        if ((0) - ((0) - (((d_507_v74_).cardinality) - (len(d_508_v75_))))) == ((self).f23):
            if (d_508_v75_).ispropersubset(d_508_v75_):
                (globalState).f7 = False
                d_509_v76_: _dafny.Array
                def lambda16_(d_510_i7_):
                    return default__.safeDivisionInt(d_510_i7_, (self).f23)

                init9_ = lambda16_
                nw79_ = _dafny.Array(None, 21)
                for i0_9_ in range(nw79_.length(0)):
                    nw79_[i0_9_] = init9_(i0_9_)
                d_509_v76_ = nw79_
                d_511_v77_: _dafny.Array
                nw80_ = _dafny.Array(None, 20)
                nw80_[int(0)] = d_509_v76_
                nw80_[int(1)] = d_509_v76_
                nw80_[int(2)] = d_509_v76_
                nw80_[int(3)] = d_509_v76_
                nw80_[int(4)] = d_509_v76_
                nw80_[int(5)] = d_509_v76_
                nw80_[int(6)] = d_509_v76_
                nw80_[int(7)] = d_509_v76_
                nw80_[int(8)] = d_509_v76_
                nw80_[int(9)] = d_509_v76_
                nw80_[int(10)] = d_509_v76_
                nw80_[int(11)] = d_509_v76_
                nw80_[int(12)] = d_509_v76_
                nw80_[int(13)] = d_509_v76_
                nw80_[int(14)] = d_509_v76_
                nw80_[int(15)] = (d_509_v76_ if (self).f28 else d_509_v76_)
                nw80_[int(16)] = d_509_v76_
                nw80_[int(17)] = d_509_v76_
                nw80_[int(18)] = d_509_v76_
                nw80_[int(19)] = d_509_v76_
                d_511_v77_ = nw80_
                index107_ = default__.safeIndex(71, (d_511_v77_).length(0))
                (d_511_v77_)[index107_] = d_509_v76_
                index108_ = default__.safeIndex(71, (d_511_v77_).length(0))
                (d_511_v77_)[index108_] = d_509_v76_
                d_512_v78_: _dafny.MultiSet
                d_512_v78_ = _dafny.MultiSet([default__.safeDivisionInt((self).f23, (self).f23)])
                (globalState).f8 = ((d_512_v78_)[(self).f23] if ((self).f23) in (d_512_v78_) else (self).f23)
                (globalState).f21 = (self).f28
                d_513_v79_: _dafny.MultiSet
                d_513_v79_ = _dafny.MultiSet([d_509_v76_, d_509_v76_, d_509_v76_])
                d_513_v79_ = d_513_v79_
            elif True:
                d_514_v80_: D8
                d_514_v80_ = D8_DC19((self.f29).set(default__.safeIndex((self).f23, len(self.f29)), (self).f23))
                d_515_v81_: _dafny.MultiSet
                d_515_v81_ = _dafny.MultiSet([self.f29, (d_514_v80_).cf32])
                (globalState).f21 = (d_515_v81_) == (d_515_v81_)
                (globalState).f6 = default__.fm2(globalState)
                d_404_v0_ = p0
                d_516_v82_: _dafny.Array
                nw81_ = _dafny.Array(int(0), 21)
                d_516_v82_ = nw81_
                index109_ = default__.safeIndex(127, (d_516_v82_).length(0))
                (d_516_v82_)[index109_] = (self).f23
                index110_ = default__.safeIndex(127, (d_516_v82_).length(0))
                (d_516_v82_)[index110_] = (self).f23
                d_517_v83_: D3
                d_517_v83_ = D3_DC11((self).f23)
                d_518_v84_: _dafny.Map
                d_518_v84_ = _dafny.Map({d_508_v75_: d_517_v83_})
                d_519_v86_: _dafny.Seq
                d_519_v86_ = _dafny.SeqWithoutIsStrInference([d_508_v75_])
                index111_ = default__.safeIndex(127, (d_516_v82_).length(0))
                def iife23_():
                    coll19_ = _dafny.Map()
                    compr_19_: _dafny.Set
                    for compr_19_ in (d_519_v86_).Elements:
                        d_520_v85_: _dafny.Set = compr_19_
                        if (d_520_v85_) in (d_519_v86_):
                            coll19_[d_520_v85_] = d_517_v83_
                    return _dafny.Map(coll19_)
                rhs110_ = (iife23_()
                ) | (d_518_v84_)
                rhs111_ = default__.safeModuloInt((0) - (len((d_443_v36_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udxyoam"))))), (d_516_v82_)[default__.safeIndex(127, (d_516_v82_).length(0))])
                lhs104_ = d_516_v82_
                lhs105_ = default__.safeIndex(127, (d_516_v82_).length(0))
                d_518_v84_ = rhs110_
                lhs104_[lhs105_] = rhs111_
            (globalState).f21 = ((self).f23) == (805)
            d_521_v87_: _dafny.Seq
            d_521_v87_ = _dafny.SeqWithoutIsStrInference([p1])
            index112_ = default__.safeIndex(347, (p0).length(0))
            (p0)[index112_] = ((d_521_v87_).set(default__.safeIndex((self).f23, len(d_521_v87_)), (self).f28)) < (_dafny.SeqWithoutIsStrInference([True, (self).f28]))
            index113_ = default__.safeIndex(347, (p0).length(0))
            (p0)[index113_] = p1
            d_522_v89_: _dafny.Map
            def iife24_():
                coll20_ = _dafny.Map()
                compr_20_: str
                for compr_20_ in (d_443_v36_).Elements:
                    d_523_v88_: str = compr_20_
                    if (d_523_v88_) in (d_443_v36_):
                        coll20_[d_523_v88_] = (self).f23
                return _dafny.Map(coll20_)
            d_522_v89_ = _dafny.Map({True: iife24_()
            })
            d_524_v90_: _dafny.Set
            d_524_v90_ = _dafny.Set({(p0)[default__.safeIndex(347, (p0).length(0))]})
            d_525_v91_: _dafny.Map
            d_525_v91_ = _dafny.Map({default__.fm0(len(_dafny.Map({d_444_v37_: len(d_522_v89_)})), True, globalState): ((self.f29).set(default__.safeIndex(len(d_524_v90_), len(self.f29)), (self).f23)).set(default__.safeIndex((self).f23, len((self.f29).set(default__.safeIndex(len(d_524_v90_), len(self.f29)), (self).f23))), 943)})
            d_526_v92_: _dafny.Array
            nw82_ = _dafny.Array(int(0), 17)
            d_526_v92_ = nw82_
            d_527_v93_: _dafny.Map
            d_527_v93_ = _dafny.Map({d_526_v92_: self.f29})
            d_528_v94_: C1
            nw83_ = C1()
            nw83_.ctor__(False, ((d_527_v93_)[d_526_v92_] if (d_526_v92_) in (d_527_v93_) else self.f29), (self).f23)
            d_528_v94_ = nw83_
            d_529_v95_: _dafny.Map
            d_529_v95_ = _dafny.Map({((d_525_v91_)[(self).f23] if ((self).f23) in (d_525_v91_) else _dafny.SeqWithoutIsStrInference([len(d_525_v91_)])): d_528_v94_})
            d_529_v95_ = (d_529_v95_).set((self.f29) + (_dafny.SeqWithoutIsStrInference([(self).f23])), d_528_v94_)
            (globalState).f21 = True
        elif True:
            d_530_v96_: _dafny.Seq
            d_530_v96_ = _dafny.SeqWithoutIsStrInference([(self).f28, False, p1])
            rhs112_ = (self).f23
            rhs113_ = d_530_v96_
            rhs114_ = p1
            lhs106_ = globalState
            r0 = rhs112_
            d_530_v96_ = rhs113_
            lhs106_.f9 = rhs114_
            d_531_v97_: _dafny.Map
            d_531_v97_ = _dafny.Map({(self).f23: (self).f28})
            d_531_v97_ = (d_531_v97_).set(len(d_443_v36_), not(True))
            d_532_v98_: _dafny.Array
            nw84_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_532_v98_ = nw84_
            d_533_v99_: _dafny.Array
            def lambda17_(d_534_v37_):
                def lambda18_(d_535_i8_):
                    return d_534_v37_

                return lambda18_

            init10_ = lambda17_(d_444_v37_)
            nw85_ = _dafny.Array(None, 28)
            for i0_10_ in range(nw85_.length(0)):
                nw85_[i0_10_] = init10_(i0_10_)
            d_533_v99_ = nw85_
            index114_ = default__.safeIndex(139, (d_532_v98_).length(0))
            (d_532_v98_)[index114_] = d_533_v99_
            index115_ = default__.safeIndex(139, (d_532_v98_).length(0))
            rhs115_ = (0) - ((self.f29)[default__.safeIndex((self).f23, len(self.f29))])
            rhs116_ = d_533_v99_
            rhs117_ = not(p1)
            rhs118_ = (self).f23
            lhs107_ = globalState
            lhs108_ = d_532_v98_
            lhs109_ = default__.safeIndex(139, (d_532_v98_).length(0))
            lhs110_ = globalState
            lhs111_ = globalState
            lhs107_.f14 = rhs115_
            lhs108_[lhs109_] = rhs116_
            lhs110_.f9 = rhs117_
            lhs111_.f14 = rhs118_
            d_536_v100_: _dafny.Map
            d_536_v100_ = _dafny.Map({(self).f23: (self).f23})
            d_537_v101_: C2
            nw86_ = C2()
            nw86_.ctor__(d_536_v100_, p1, (not(p1)) or (p1), (_dafny.SeqWithoutIsStrInference([(self).f23, (self).f23])) + (self.f29), ((self).f23 if (self).f28 else (self).f23))
            d_537_v101_ = nw86_
            (globalState).f8 = (self).f23
        r0 = (self).f23
        d_538_v102_: _dafny.MultiSet
        d_538_v102_ = _dafny.MultiSet([(self).f23])
        d_539_v103_: _dafny.Map
        d_539_v103_ = _dafny.Map({(self).f23: (self).f23})
        d_540_v104_: _dafny.Map
        d_540_v104_ = _dafny.Map({d_444_v37_: ((d_538_v102_)[len(d_539_v103_)] if (len(d_539_v103_)) in (d_538_v102_) else (self).f23)})
        def iife25_():
            coll21_ = _dafny.Set()
            compr_21_: str
            for compr_21_ in (d_540_v104_).keys.Elements:
                d_541_v105_: str = compr_21_
                if (d_541_v105_) in (d_540_v104_):
                    coll21_ = coll21_.union(_dafny.Set([d_541_v105_]))
            return _dafny.Set(coll21_)
        r1 = (d_508_v75_) - (_dafny.Set({len(iife25_()
        ), (self).f23, len(d_443_v36_)}))
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_542_v0_: _dafny.Array
        def lambda19_(d_543_i0_):
            return default__.safeDivisionInt(d_543_i0_, (self).f23)

        init11_ = lambda19_
        nw87_ = _dafny.Array(None, 26)
        for i0_11_ in range(nw87_.length(0)):
            nw87_[i0_11_] = init11_(i0_11_)
        d_542_v0_ = nw87_
        d_544_v1_: _dafny.Seq
        d_544_v1_ = _dafny.SeqWithoutIsStrInference([d_542_v0_])
        d_545_v2_: D1
        d_545_v2_ = D1_DC3(d_544_v1_)
        d_546_v3_: _dafny.MultiSet
        d_546_v3_ = _dafny.MultiSet([d_545_v2_, d_545_v2_, d_545_v2_])
        d_547_v4_: _dafny.Map
        d_547_v4_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([False]): d_546_v3_})
        d_548_v5_: _dafny.Seq
        d_548_v5_ = _dafny.SeqWithoutIsStrInference([(self).f28])
        d_547_v4_ = (d_547_v4_).set(d_548_v5_, d_546_v3_)
        index116_ = default__.safeIndex(260, (d_542_v0_).length(0))
        (d_542_v0_)[index116_] = (self).f23
        d_549_v6_: _dafny.Seq
        d_549_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_550_i1_ in range(default__.abs(317))])])
        d_551_v7_: _dafny.Seq
        d_551_v7_ = d_549_v6_
        index117_ = default__.safeIndex(260, (d_542_v0_).length(0))
        def lambda20_(source4_):
            d_552___mcc_h0_ = source4_
            d_553_cf30_ = d_552___mcc_h0_
            return (self).f23

        (d_542_v0_)[index117_] = lambda20_((d_551_v7_ if (self).f28 else d_549_v6_))
        d_554_i2_: int
        d_554_i2_ = 0
        with _dafny.label("7"):
            while not ((self).f28) or (((self).f23) in (self.f29)):
                with _dafny.c_label("7"):
                    if (d_554_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_554_i2_ = (d_554_i2_) + (1)
                    d_555_v8_: _dafny.MultiSet
                    d_555_v8_ = _dafny.MultiSet([(self).f28])
                    d_556_v9_: str
                    d_556_v9_ = _dafny.CodePoint('g')
                    d_557_v10_: D2
                    d_557_v10_ = D2_DC8((self).f23, (d_555_v8_).cardinality, d_556_v9_, len(p0), (self).f28)
                    rhs119_ = not(not((d_557_v10_).cf19))
                    rhs120_ = ((self).f23) + ((692) + ((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))]))
                    rhs121_ = ((d_555_v8_).set((self).f28, default__.abs((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))]))) != (d_555_v8_)
                    lhs112_ = globalState
                    lhs113_ = globalState
                    lhs114_ = globalState
                    lhs112_.f7 = rhs119_
                    lhs113_.f8 = rhs120_
                    lhs114_.f21 = rhs121_
                    (globalState).f21 = False
                    (globalState).f11 = d_556_v9_
                    (globalState).f8 = (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))]
                    pass
            pass
        d_558_v11_: _dafny.Array
        def lambda21_(d_559_i3_):
            return _dafny.CodePoint('p')

        init12_ = lambda21_
        nw88_ = _dafny.Array(None, 15)
        for i0_12_ in range(nw88_.length(0)):
            nw88_[i0_12_] = init12_(i0_12_)
        d_558_v11_ = nw88_
        d_560_v12_: _dafny.Map
        d_560_v12_ = _dafny.Map({d_558_v11_: (self).f28})
        d_560_v12_ = (d_560_v12_).set(d_558_v11_, (self).f28)
        if False:
            d_561_v14_: _dafny.Map
            d_561_v14_ = _dafny.Map({len(d_548_v5_): (self).f28})
            d_562_v15_: _dafny.Seq
            d_562_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_561_v14_)])])
            d_563_v16_: C2
            nw89_ = C2()
            def iife26_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in ((d_562_v15_)[default__.safeIndex((self).f23, len(d_562_v15_))]).Elements:
                    d_564_v13_: int = compr_22_
                    if (d_564_v13_) in ((d_562_v15_)[default__.safeIndex((self).f23, len(d_562_v15_))]):
                        coll22_[(d_564_v13_) - ((self).f23)] = (self).f23
                return _dafny.Map(coll22_)
            nw89_.ctor__(iife26_()
            , (self).f28, (self).f28, self.f29, (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
            d_563_v16_ = nw89_
            d_565_v17_: _dafny.MultiSet
            d_565_v17_ = _dafny.MultiSet([(d_563_v16_).f32, (d_563_v16_).f32, (self).f28])
            (globalState).f6 = (d_549_v6_)[default__.safeIndex((d_565_v17_).cardinality, len(d_549_v6_))]
            d_566_v18_: _dafny.Array
            def lambda22_(d_567_p0_):
                def lambda23_(d_568_i4_):
                    return d_567_p0_

                return lambda23_

            init13_ = lambda22_(p0)
            nw90_ = _dafny.Array(None, 17)
            for i0_13_ in range(nw90_.length(0)):
                nw90_[i0_13_] = init13_(i0_13_)
            d_566_v18_ = nw90_
            d_566_v18_ = d_566_v18_
            d_569_v19_: _dafny.Array
            nw91_ = _dafny.Array(False, 21)
            d_569_v19_ = nw91_
            d_570_v20_: _dafny.Map
            d_570_v20_ = _dafny.Map({d_569_v19_: (945) >= ((0) - (-555))})
            (globalState).f21 = ((d_570_v20_)[d_569_v19_] if (d_569_v19_) in (d_570_v20_) else (self).f28)
            index118_ = default__.safeIndex(409, (d_569_v19_).length(0))
            (d_569_v19_)[index118_] = (d_563_v16_).f32
            d_571_v21_: _dafny.Map
            d_571_v21_ = _dafny.Map({(d_563_v16_).f32: (d_563_v16_).f32})
            d_572_v22_: _dafny.Map
            d_572_v22_ = _dafny.Map({(self).f28: len(d_571_v21_)})
            d_573_v23_: _dafny.MultiSet
            d_573_v23_ = _dafny.MultiSet([39, (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))]])
            index119_ = default__.safeIndex(409, (d_569_v19_).length(0))
            (d_569_v19_)[index119_] = ((len(d_572_v22_)) - ((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])) != (((d_573_v23_)[(self).f23] if ((self).f23) in (d_573_v23_) else len(d_548_v5_)))
        elif True:
            d_574_v24_: _dafny.Map
            d_574_v24_ = _dafny.Map({(0) - (len(d_549_v6_)): 63})
            d_575_v25_: T1
            nw92_ = C2()
            nw92_.ctor__(d_574_v24_, (self).f28, (self).f28, self.f29, (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
            d_575_v25_ = nw92_
            d_576_v26_: T1
            d_576_v26_ = d_575_v25_
            source5_ = d_576_v26_
            d_577___mcc_h1_ = source5_
            d_578_cf29_ = d_577___mcc_h1_
            d_579_v27_: _dafny.MultiSet
            d_579_v27_ = _dafny.MultiSet([(d_575_v25_).f23, (d_578_cf29_).f23, (0) - (len(d_548_v5_))])
            d_580_v28_: _dafny.MultiSet
            d_580_v28_ = _dafny.MultiSet([d_579_v27_, _dafny.MultiSet(self.f29), (d_579_v27_) | (d_579_v27_), (d_579_v27_).set((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))], default__.abs((d_575_v25_).f23)), d_579_v27_])
            d_581_v29_: _dafny.MultiSet
            d_581_v29_ = _dafny.MultiSet([d_558_v11_, d_558_v11_, d_558_v11_])
            rhs122_ = default__.safeModuloInt((0) - ((self).f23), ((self).f23) + (len(p0)))
            rhs123_ = _dafny.MultiSet([(_dafny.MultiSet([(d_575_v25_).f23, (d_581_v29_).cardinality, (d_575_v25_).f23, (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))], len(p0)])).set((d_575_v25_).f23, default__.abs((d_578_cf29_).f23)), d_579_v27_, (d_579_v27_).intersection(d_579_v27_), d_579_v27_])
            r0 = rhs122_
            d_580_v28_ = rhs123_
            d_582_v30_: _dafny.Array
            def lambda24_(d_583_v25_, d_584_cf29_):
                def lambda25_(d_585_i5_):
                    return (_dafny.MultiSet([False, not((d_583_v25_).f28), (d_584_cf29_).f28, (d_584_cf29_).f28])) | (_dafny.MultiSet([(d_583_v25_).f28, False, not(False), (d_583_v25_).f28]))

                return lambda25_

            init14_ = lambda24_(d_575_v25_, d_578_cf29_)
            nw93_ = _dafny.Array(None, 27)
            for i0_14_ in range(nw93_.length(0)):
                nw93_[i0_14_] = init14_(i0_14_)
            d_582_v30_ = nw93_
            rhs124_ = (d_575_v25_).f28
            rhs125_ = d_582_v30_
            lhs115_ = globalState
            lhs115_.f21 = rhs124_
            d_582_v30_ = rhs125_
            d_586_v31_: D3
            d_586_v31_ = D3_DC11((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
            pat_let_tv27_ = d_575_v25_
            d_587_v32_: _dafny.Map
            def iife27_(_pat_let2_0):
                def iife28_(d_588_dt__update__tmp_h0_):
                    def iife29_(_pat_let3_0):
                        def iife30_(d_589_dt__update_hcf22_h0_):
                            return D3_DC11(d_589_dt__update_hcf22_h0_)
                        return iife30_(_pat_let3_0)
                    return iife29_((pat_let_tv27_).f23)
                return iife28_(_pat_let2_0)
            d_587_v32_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([iife27_(d_586_v31_)])) + (_dafny.SeqWithoutIsStrInference([d_586_v31_ for d_590_i6_ in range(default__.abs(490))])): -173})
            d_591_v33_: _dafny.Seq
            d_591_v33_ = _dafny.SeqWithoutIsStrInference([D3_DC11((d_578_cf29_).f23), d_586_v31_])
            d_587_v32_ = (d_587_v32_).set((d_591_v33_ if (d_578_cf29_).f28 else d_591_v33_), (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
            d_592_v34_: _dafny.MultiSet
            d_592_v34_ = _dafny.MultiSet([d_542_v0_, d_542_v0_, d_542_v0_])
            (globalState).f7 = ((_dafny.MultiSet([d_542_v0_])).set((d_544_v1_)[default__.safeIndex((d_578_cf29_).f23, len(d_544_v1_))], default__.abs((d_575_v25_).f23))).ispropersubset(d_592_v34_)
            d_593_v35_: _dafny.Map
            d_593_v35_ = _dafny.Map({d_542_v0_: d_575_v25_.f29})
            d_594_v36_: _dafny.Map
            d_594_v36_ = _dafny.Map({False: (d_593_v35_).set(d_542_v0_, self.f29)})
            d_594_v36_ = (d_594_v36_).set((d_575_v25_).f28, d_593_v35_)
            if not(default__.fm1((d_575_v25_).f23, (d_575_v25_).f28, globalState)):
                (globalState).f7 = (d_575_v25_).f28
                d_558_v11_ = d_558_v11_
                (globalState).f9 = False
                d_575_v25_ = d_575_v25_
                d_595_v37_: _dafny.Array
                nw94_ = _dafny.Array(None, 5)
                nw94_[int(0)] = (d_575_v25_).f28
                nw94_[int(1)] = (d_575_v25_).f28
                nw94_[int(2)] = default__.fm1((0) - ((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))]), (d_575_v25_).f28, globalState)
                nw94_[int(3)] = (d_575_v25_).f28
                nw94_[int(4)] = default__.fm1((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))], (d_575_v25_).f28, globalState)
                d_595_v37_ = nw94_
                index120_ = default__.safeIndex(829, (d_595_v37_).length(0))
                (d_595_v37_)[index120_] = (self).f28
                d_596_v38_: str
                d_596_v38_ = _dafny.CodePoint('a')
                d_597_v39_: D9
                d_597_v39_ = D9_DC24(d_558_v11_)
                index121_ = default__.safeIndex(829, (d_595_v37_).length(0))
                rhs126_ = (d_596_v38_) not in (p0)
                rhs127_ = -741
                rhs128_ = (d_597_v39_).cf40
                rhs129_ = (_dafny.SeqWithoutIsStrInference([d_596_v38_ for d_598_i7_ in range(default__.abs(729))])) == (p0)
                rhs130_ = d_595_v37_
                lhs116_ = globalState
                lhs117_ = globalState
                lhs118_ = d_595_v37_
                lhs119_ = default__.safeIndex(829, (d_595_v37_).length(0))
                lhs116_.f21 = rhs126_
                lhs117_.f14 = rhs127_
                d_558_v11_ = rhs128_
                lhs118_[lhs119_] = rhs129_
                d_595_v37_ = rhs130_
            elif True:
                d_599_v40_: _dafny.Map
                d_599_v40_ = _dafny.Map({d_548_v5_: (self).f28})
                d_600_v42_: str
                d_600_v42_ = _dafny.CodePoint('w')
                d_601_v43_: _dafny.Map
                d_601_v43_ = _dafny.Map({d_600_v42_: (d_575_v25_).f28})
                def iife31_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_602_i9_ in range(default__.abs(177))])) for d_603_i8_ in range(default__.abs(770))])).Elements:
                        d_604_v41_: int = compr_23_
                        if (d_604_v41_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_602_i9_ in range(default__.abs(177))])) for d_603_i8_ in range(default__.abs(770))])):
                            coll23_[default__.safeModuloInt(d_604_v41_, (d_575_v25_).f23)] = (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))]
                    return _dafny.Map(coll23_)
                d_599_v40_ = (_dafny.Map({d_548_v5_: (self).f28}) if (iife31_()
                ) == (default__.fm16((self).f28, len(d_601_v43_), (self).f23, globalState)) else d_599_v40_)
                d_605_v44_: D9
                d_605_v44_ = D9_DC26(548)
                pat_let_tv28_ = d_542_v0_
                pat_let_tv29_ = d_542_v0_
                d_606_v45_: _dafny.Array
                nw95_ = _dafny.Array(None, 22)
                nw95_[int(0)] = D9_DC26((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
                nw95_[int(1)] = d_605_v44_
                nw95_[int(2)] = d_605_v44_
                nw95_[int(3)] = d_605_v44_
                nw95_[int(4)] = d_605_v44_
                nw95_[int(5)] = d_605_v44_
                nw95_[int(6)] = d_605_v44_
                nw95_[int(7)] = d_605_v44_
                nw95_[int(8)] = d_605_v44_
                nw95_[int(9)] = d_605_v44_
                nw95_[int(10)] = d_605_v44_
                nw95_[int(11)] = d_605_v44_
                nw95_[int(12)] = d_605_v44_
                nw95_[int(13)] = D9_DC26(604)
                def iife32_(_pat_let4_0):
                    def iife33_(d_607_dt__update__tmp_h1_):
                        def iife34_(_pat_let5_0):
                            def iife35_(d_608_dt__update_hcf43_h0_):
                                return D9_DC26(d_608_dt__update_hcf43_h0_)
                            return iife35_(_pat_let5_0)
                        return iife34_((self).f23)
                    return iife33_(_pat_let4_0)
                nw95_[int(14)] = iife32_(d_605_v44_)
                nw95_[int(15)] = d_605_v44_
                nw95_[int(16)] = d_605_v44_
                nw95_[int(17)] = d_605_v44_
                def iife36_(_pat_let6_0):
                    def iife37_(d_609_dt__update__tmp_h2_):
                        def iife38_(_pat_let7_0):
                            def iife39_(d_610_dt__update_hcf43_h1_):
                                return D9_DC26(d_610_dt__update_hcf43_h1_)
                            return iife39_(_pat_let7_0)
                        return iife38_((pat_let_tv29_)[default__.safeIndex(260, (pat_let_tv28_).length(0))])
                    return iife37_(_pat_let6_0)
                nw95_[int(18)] = iife36_(d_605_v44_)
                nw95_[int(19)] = d_605_v44_
                nw95_[int(20)] = d_605_v44_
                nw95_[int(21)] = d_605_v44_
                d_606_v45_ = nw95_
                index122_ = default__.safeIndex(94, (d_606_v45_).length(0))
                (d_606_v45_)[index122_] = d_605_v44_
                index123_ = default__.safeIndex(94, (d_606_v45_).length(0))
                (d_606_v45_)[index123_] = d_605_v44_
                index124_ = default__.safeIndex(260, (d_542_v0_).length(0))
                (d_542_v0_)[index124_] = 721
                d_611_v46_: _dafny.Array
                nw96_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_611_v46_ = nw96_
                d_612_v47_: _dafny.MultiSet
                d_612_v47_ = _dafny.MultiSet([True])
                d_613_v48_: _dafny.Seq
                d_613_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f28})])
                d_614_v49_: D0
                d_614_v49_ = D0_DC0(p0)
                d_615_v50_: _dafny.MultiSet
                d_615_v50_ = _dafny.MultiSet([len(p0)])
                d_616_v51_: _dafny.Array
                nw97_ = _dafny.Array(None, 22)
                nw97_[int(0)] = d_612_v47_
                nw97_[int(1)] = d_612_v47_
                nw97_[int(2)] = d_612_v47_
                nw97_[int(3)] = d_612_v47_
                nw97_[int(4)] = d_612_v47_
                nw97_[int(5)] = d_612_v47_
                nw97_[int(6)] = d_612_v47_
                nw97_[int(7)] = _dafny.MultiSet([(d_575_v25_).f28, (self).f28])
                nw97_[int(8)] = _dafny.MultiSet([(d_575_v25_).f28, (self).f28])
                nw97_[int(9)] = d_612_v47_
                nw97_[int(10)] = d_612_v47_
                nw97_[int(11)] = d_612_v47_
                nw97_[int(12)] = (((self).f30)[default__.safeIndex((self).f23, len((self).f30))]).set((self).f28, default__.abs(len(d_613_v48_)))
                nw97_[int(13)] = d_612_v47_
                nw97_[int(14)] = d_612_v47_
                nw97_[int(15)] = (d_612_v47_).set((d_575_v25_).f28, default__.abs((d_575_v25_).f23))
                nw97_[int(16)] = _dafny.MultiSet([(d_575_v25_).f28, (d_575_v25_).f28])
                nw97_[int(17)] = default__.fm4(d_614_v49_, (self).f23, d_548_v5_, d_600_v42_, globalState)
                nw97_[int(18)] = _dafny.MultiSet(d_548_v5_)
                nw97_[int(19)] = _dafny.MultiSet([(d_575_v25_).f28, (self).f28])
                nw97_[int(20)] = d_612_v47_
                nw97_[int(21)] = default__.fm4(d_614_v49_, ((d_615_v50_)[(d_612_v47_).cardinality] if ((d_612_v47_).cardinality) in (d_615_v50_) else (d_575_v25_).f23), d_548_v5_, d_600_v42_, globalState)
                d_616_v51_ = nw97_
                index125_ = default__.safeIndex(491, (d_611_v46_).length(0))
                (d_611_v46_)[index125_] = d_616_v51_
                d_617_v52_: _dafny.Map
                d_617_v52_ = _dafny.Map({(d_575_v25_).f23: (self).f28})
                d_618_v53_: _dafny.Array
                def lambda26_(d_619_i10_):
                    return False

                init15_ = lambda26_
                nw98_ = _dafny.Array(None, 27)
                for i0_15_ in range(nw98_.length(0)):
                    nw98_[i0_15_] = init15_(i0_15_)
                d_618_v53_ = nw98_
                d_620_v54_: _dafny.Map
                d_620_v54_ = _dafny.Map({d_618_v53_: default__.fm1((d_575_v25_).f23, (self).f28, globalState)})
                index126_ = default__.safeIndex(491, (d_611_v46_).length(0))
                rhs131_ = (len(d_617_v52_)) >= (len(d_620_v54_))
                rhs132_ = (d_575_v25_).f23
                rhs133_ = d_616_v51_
                lhs120_ = globalState
                lhs121_ = globalState
                lhs122_ = d_611_v46_
                lhs123_ = default__.safeIndex(491, (d_611_v46_).length(0))
                lhs120_.f21 = rhs131_
                lhs121_.f8 = rhs132_
                lhs122_[lhs123_] = rhs133_
                d_621_v55_: D0
                d_622_v56_: _dafny.Map
                d_623_v57_: int
                out27_: D0
                out28_: _dafny.Map
                out29_: int
                out27_, out28_, out29_ = default__.m0(d_618_v53_, (p0) + ((p0).set(default__.safeIndex((self).f23, len(p0)), d_600_v42_)), (d_575_v25_).f23, globalState)
                d_621_v55_ = out27_
                d_622_v56_ = out28_
                d_623_v57_ = out29_
            d_624_v58_: D8
            d_624_v58_ = D8_DC22((self).f28)
            (globalState).f8 = ((self).f23 if (d_624_v58_).cf38 else (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
            (globalState).f14 = (self).f23
        d_625_v59_: C0
        nw99_ = C0()
        nw99_.ctor__()
        d_625_v59_ = nw99_
        d_626_v60_: _dafny.Map
        d_626_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtk")): (self).f23})
        d_627_v61_: _dafny.Map
        d_627_v61_ = _dafny.Map({499: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gun"))})
        d_628_v62_: str
        d_628_v62_ = _dafny.CodePoint('a')
        d_629_v63_: _dafny.Set
        d_629_v63_ = _dafny.Set({(self).f23, -620, (self).f23, (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))], (self).f23})
        d_630_v64_: _dafny.Map
        d_630_v64_ = _dafny.Map({((d_626_v60_)[(((d_627_v61_)[-191] if (-191) in (d_627_v61_) else p0)).set(default__.safeIndex((self).f23, len(((d_627_v61_)[-191] if (-191) in (d_627_v61_) else p0))), d_628_v62_)] if ((((d_627_v61_)[-191] if (-191) in (d_627_v61_) else p0)).set(default__.safeIndex((self).f23, len(((d_627_v61_)[-191] if (-191) in (d_627_v61_) else p0))), d_628_v62_)) in (d_626_v60_) else len(d_629_v63_)): (0) - ((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])})
        r0 = ((d_630_v64_)[(0) - ((self).f23)] if ((0) - ((self).f23)) in (d_630_v64_) else (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
        r1 = ((self.f29) + (self.f29)).set(default__.safeIndex((d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))], len((self.f29) + (self.f29))), (d_542_v0_)[default__.safeIndex(260, (d_542_v0_).length(0))])
        return r0, r1

    def m5(self, globalState):
        d_631_v0_: _dafny.Array
        nw100_ = _dafny.Array(int(0), 1)
        d_631_v0_ = nw100_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_631_v0_).length(0)):
            d_632_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_632_i0_)) and ((d_632_i0_) < ((d_631_v0_).length(0)))):
                (d_631_v0_)[(d_632_i0_)] = default__.safeModuloInt(d_632_i0_, (len(self.f29)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgr")))))
        d_633_i1_: int
        d_633_i1_ = 0
        with _dafny.label("8"):
            while not((self).f28):
                with _dafny.c_label("8"):
                    if (d_633_i1_) >= (100):
                        raise _dafny.Break("8")
                    d_633_i1_ = (d_633_i1_) + (1)
                    d_634_v1_: C0
                    nw101_ = C0()
                    nw101_.ctor__()
                    d_634_v1_ = nw101_
                    d_635_v2_: _dafny.Array
                    nw102_ = _dafny.Array(False, 23)
                    d_635_v2_ = nw102_
                    d_636_v3_: _dafny.Set
                    d_636_v3_ = _dafny.Set({(0) - ((self).f23)})
                    d_637_v4_: _dafny.Map
                    d_637_v4_ = _dafny.Map({(self).f23: (self).f23})
                    index127_ = default__.safeIndex(174, (d_635_v2_).length(0))
                    (d_635_v2_)[index127_] = (d_636_v3_).ispropersubset(_dafny.Set({((d_637_v4_)[(self.f29)[default__.safeIndex((self).f23, len(self.f29))]] if ((self.f29)[default__.safeIndex((self).f23, len(self.f29))]) in (d_637_v4_) else (self).f23)}))
                    index128_ = default__.safeIndex(174, (d_635_v2_).length(0))
                    def iife40_():
                        coll24_ = _dafny.Set()
                        compr_24_: int
                        for compr_24_ in _dafny.IntegerRange(979, 687):
                            d_638_v5_: int = compr_24_
                            if ((979) <= (d_638_v5_)) and ((d_638_v5_) < (687)):
                                coll24_ = coll24_.union(_dafny.Set([(d_638_v5_) + ((self).f23)]))
                        return _dafny.Set(coll24_)
                    (d_635_v2_)[index128_] = not(((iife40_()
                    ).isdisjoint(d_636_v3_) if not((self).f28) else ((self).f28 if (self).f28 else (self).f28)))
                    d_639_v6_: _dafny.Set
                    d_639_v6_ = _dafny.Set({False, (self).f28})
                    d_640_v7_: _dafny.Seq
                    d_640_v7_ = _dafny.SeqWithoutIsStrInference([d_639_v6_])
                    rhs134_ = (self).f23
                    rhs135_ = (d_640_v7_)[default__.safeIndex((self).f23, len(d_640_v7_))]
                    rhs136_ = (self.f29)[default__.safeIndex((self).f23, len(self.f29))]
                    lhs124_ = globalState
                    lhs125_ = globalState
                    lhs124_.f14 = rhs134_
                    d_639_v6_ = rhs135_
                    lhs125_.f8 = rhs136_
                    index129_ = default__.safeIndex(174, (d_635_v2_).length(0))
                    (d_635_v2_)[index129_] = (self).f28
                    pass
            pass
        d_641_v8_: _dafny.Seq
        d_641_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgw"))
        d_642_v9_: _dafny.MultiSet
        d_642_v9_ = _dafny.MultiSet([len(d_641_v8_), (self).f23])
        d_643_v10_: _dafny.Map
        d_643_v10_ = _dafny.Map({(self).f23: self.f29})
        hi10_ = ((d_642_v9_)[685] if (685) in (d_642_v9_) else len(((d_643_v10_)[260] if (260) in (d_643_v10_) else _dafny.SeqWithoutIsStrInference([(self).f23, (self).f23]))))
        for d_644_i2_ in range(len((_dafny.SeqWithoutIsStrInference([(default__.fm17((self).f28, default__.fm0((self).f23, (self).f28, globalState), globalState)).cardinality, (self).f23])) + (self.f29)), hi10_):
            d_645_v11_: _dafny.Seq
            d_645_v11_ = _dafny.SeqWithoutIsStrInference([(self).f28])
            d_646_v12_: _dafny.MultiSet
            d_646_v12_ = _dafny.MultiSet([not((self).f28), (self).f28])
            d_647_v13_: _dafny.Map
            d_647_v13_ = _dafny.Map({False: True})
            d_648_v14_: _dafny.Map
            d_648_v14_ = _dafny.Map({(self).f28: d_647_v13_})
            d_649_v15_: _dafny.Set
            d_649_v15_ = _dafny.Set({d_644_i2_, default__.safeModuloInt((self).f23, (d_642_v9_).cardinality), (d_644_i2_) - (len(d_645_v11_)), (((d_646_v12_).set((self).f28, default__.abs(d_644_i2_))).cardinality) * ((self).f23), len(d_648_v14_)})
            d_649_v15_ = (d_649_v15_ if False else d_649_v15_)
            d_650_v16_: str
            d_650_v16_ = _dafny.CodePoint('e')
            d_651_v17_: _dafny.Map
            d_651_v17_ = _dafny.Map({len(self.f29): len(_dafny.SeqWithoutIsStrInference([((d_641_v8_).set(default__.safeIndex((self).f23, len(d_641_v8_)), d_650_v16_)).set(default__.safeIndex((self).f23, len((d_641_v8_).set(default__.safeIndex((self).f23, len(d_641_v8_)), d_650_v16_))), d_650_v16_), d_641_v8_]))})
            d_652_v18_: D4
            d_652_v18_ = D4_DC14(d_631_v0_)
            d_653_v19_: C2
            nw103_ = C2()
            nw103_.ctor__(d_651_v17_, default__.fm1(d_644_i2_, (self).f28, globalState), (self).f28, ((self.f29).set(default__.safeIndex(d_644_i2_, len(self.f29)), d_644_i2_)) + (self.f29), len(_dafny.SeqWithoutIsStrInference([d_652_v18_, d_652_v18_])))
            d_653_v19_ = nw103_
            d_651_v17_ = (d_651_v17_).set(len(d_641_v8_), (default__.fm24((d_653_v19_).f32, d_644_i2_, globalState)).cf43)
            if (self).f28:
                index130_ = default__.safeIndex(888, (d_631_v0_).length(0))
                (d_631_v0_)[index130_] = d_644_i2_
                d_654_v20_: _dafny.Array
                nw104_ = _dafny.Array(None, 6)
                nw104_[int(0)] = _dafny.CodePoint('n')
                nw104_[int(1)] = d_650_v16_
                nw104_[int(2)] = d_650_v16_
                nw104_[int(3)] = d_650_v16_
                nw104_[int(4)] = d_650_v16_
                nw104_[int(5)] = d_650_v16_
                d_654_v20_ = nw104_
                d_655_v21_: D9
                d_655_v21_ = D9_DC24(d_654_v20_)
                d_656_v22_: _dafny.Array
                nw105_ = _dafny.Array(None, 19)
                nw105_[int(0)] = d_654_v20_
                nw105_[int(1)] = d_654_v20_
                nw105_[int(2)] = d_654_v20_
                nw105_[int(3)] = d_654_v20_
                nw105_[int(4)] = d_654_v20_
                nw105_[int(5)] = d_654_v20_
                nw105_[int(6)] = d_654_v20_
                nw105_[int(7)] = d_654_v20_
                nw105_[int(8)] = d_654_v20_
                nw105_[int(9)] = d_654_v20_
                nw105_[int(10)] = d_654_v20_
                nw105_[int(11)] = d_654_v20_
                nw105_[int(12)] = d_654_v20_
                nw105_[int(13)] = d_654_v20_
                nw105_[int(14)] = d_654_v20_
                nw105_[int(15)] = d_654_v20_
                nw105_[int(16)] = (d_655_v21_).cf40
                nw105_[int(17)] = d_654_v20_
                nw105_[int(18)] = d_654_v20_
                d_656_v22_ = nw105_
                d_657_v23_: _dafny.Array
                nw106_ = _dafny.Array(_dafny.Seq({}), 11)
                d_657_v23_ = nw106_
                index131_ = default__.safeIndex(742, (d_657_v23_).length(0))
                (d_657_v23_)[index131_] = self.f29
                index132_ = default__.safeIndex(888, (d_631_v0_).length(0))
                index133_ = default__.safeIndex(742, (d_657_v23_).length(0))
                rhs137_ = 938
                rhs138_ = d_656_v22_
                rhs139_ = self.f29
                rhs140_ = (self).f28
                lhs126_ = d_631_v0_
                lhs127_ = default__.safeIndex(888, (d_631_v0_).length(0))
                lhs128_ = d_657_v23_
                lhs129_ = default__.safeIndex(742, (d_657_v23_).length(0))
                lhs130_ = globalState
                lhs126_[lhs127_] = rhs137_
                d_656_v22_ = rhs138_
                lhs128_[lhs129_] = rhs139_
                lhs130_.f21 = rhs140_
                d_658_v24_: D1
                d_658_v24_ = D1_DC5(d_644_i2_, (d_653_v19_).f32, 284, (self).f23)
                (globalState).f9 = (d_658_v24_).cf10
                (globalState).f11 = d_650_v16_
                d_659_v25_: _dafny.Map
                d_659_v25_ = _dafny.Map({(self).f23: (d_653_v19_).f32})
                d_660_v26_: _dafny.Map
                d_660_v26_ = _dafny.Map({(self).f28: len(d_659_v25_)})
                pat_let_tv30_ = d_641_v8_
                pat_let_tv31_ = d_631_v0_
                pat_let_tv32_ = d_631_v0_
                pat_let_tv33_ = d_641_v8_
                pat_let_tv34_ = d_650_v16_
                def iife41_(_pat_let8_0):
                    def iife42_(d_661_dt__update__tmp_h0_):
                        def iife43_(_pat_let9_0):
                            def iife44_(d_662_dt__update_hcf0_h0_):
                                return D0_DC0(d_662_dt__update_hcf0_h0_)
                            return iife44_(_pat_let9_0)
                        return iife43_((pat_let_tv30_).set(default__.safeIndex((0) - ((pat_let_tv32_)[default__.safeIndex(888, (pat_let_tv31_).length(0))]), len(pat_let_tv33_)), pat_let_tv34_))
                    return iife42_(_pat_let8_0)
                (globalState).f14 = (d_653_v19_).fm18((d_660_v26_ if (self).f28 else default__.fm25((_dafny.MultiSet([d_650_v16_])).cardinality, (d_653_v19_).f32, globalState)), ((iife41_(D0_DC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvnjskw"))))).cf0) + (_dafny.SeqWithoutIsStrInference([d_650_v16_ for d_663_i3_ in range(default__.abs(387))])), (d_653_v19_).f32, globalState)
                (globalState).f7 = (d_653_v19_).f32
            elif True:
                d_664_v27_: _dafny.Map
                d_664_v27_ = _dafny.Map({(self).f28: d_644_i2_})
                d_665_v28_: _dafny.Map
                d_665_v28_ = d_664_v27_
                d_665_v28_ = d_664_v27_
                (self).f29 = _dafny.SeqWithoutIsStrInference([(d_653_v19_).fm18(d_664_v27_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epl")), (d_645_v11_)[default__.safeIndex(d_644_i2_, len(d_645_v11_))], globalState), default__.fm0(d_644_i2_, (d_653_v19_).f32, globalState), (self).f23])
                index134_ = default__.safeIndex(661, (d_631_v0_).length(0))
                (d_631_v0_)[index134_] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnttsu"))))
                index135_ = default__.safeIndex(661, (d_631_v0_).length(0))
                rhs141_ = default__.safeModuloInt(len(d_641_v8_), (self).f23)
                rhs142_ = default__.safeModuloInt((self).f23, default__.safeDivisionInt((self).f23, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqnxnel")))))
                rhs143_ = (self).f23
                rhs144_ = (default__.safeDivisionInt((self).f23, d_644_i2_) if default__.fm1(len((d_653_v19_).f31), (d_653_v19_).f32, globalState) else d_644_i2_)
                lhs131_ = globalState
                lhs132_ = d_631_v0_
                lhs133_ = default__.safeIndex(661, (d_631_v0_).length(0))
                lhs134_ = globalState
                lhs135_ = globalState
                lhs131_.f8 = rhs141_
                lhs132_[lhs133_] = rhs142_
                lhs134_.f14 = rhs143_
                lhs135_.f14 = rhs144_
                d_666_v29_: _dafny.Array
                nw107_ = _dafny.Array(_dafny.Map({}), 8)
                d_666_v29_ = nw107_
                d_667_v30_: _dafny.Map
                d_667_v30_ = _dafny.Map({d_641_v8_: (d_653_v19_).f32})
                index136_ = default__.safeIndex(532, (d_666_v29_).length(0))
                (d_666_v29_)[index136_] = (d_667_v30_ if (d_653_v19_).f32 else _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pinimjpwv")): (self).f28}))
                d_668_v31_: _dafny.Array
                def lambda27_(d_669_v0_):
                    def lambda28_(d_670_i4_):
                        return (d_670_i4_) * ((d_669_v0_)[default__.safeIndex(661, (d_669_v0_).length(0))])

                    return lambda28_

                init16_ = lambda27_(d_631_v0_)
                nw108_ = _dafny.Array(None, 25)
                for i0_16_ in range(nw108_.length(0)):
                    nw108_[i0_16_] = init16_(i0_16_)
                d_668_v31_ = nw108_
                index137_ = default__.safeIndex(532, (d_666_v29_).length(0))
                rhs145_ = (d_667_v30_).set(d_641_v8_, not((d_653_v19_).f32))
                rhs146_ = d_668_v31_
                rhs147_ = default__.safeDivisionInt(d_644_i2_, (0) - ((d_631_v0_)[default__.safeIndex(661, (d_631_v0_).length(0))]))
                lhs136_ = d_666_v29_
                lhs137_ = default__.safeIndex(532, (d_666_v29_).length(0))
                lhs138_ = globalState
                lhs136_[lhs137_] = rhs145_
                d_668_v31_ = rhs146_
                lhs138_.f14 = rhs147_
                d_671_v32_: _dafny.Map
                d_671_v32_ = _dafny.Map({default__.fm1(564, (self).f28, globalState): d_650_v16_})
                d_672_v33_: _dafny.Map
                d_672_v33_ = _dafny.Map({(d_653_v19_).f32: d_650_v16_})
                d_673_v34_: _dafny.MultiSet
                d_673_v34_ = _dafny.MultiSet([(d_671_v32_) | (d_671_v32_), (d_672_v33_), (d_671_v32_ if default__.fm1(d_644_i2_, (d_653_v19_).f32, globalState) else d_671_v32_)])
                d_673_v34_ = d_673_v34_
        (globalState).f14 = (self).f23
        (globalState).f14 = (self).f23
        (globalState).f9 = ((self).f28 if (self).f28 else (self).f28)

    @property
    def f30(self):
        return self._f30

class C4(T0):
    def  __init__(self):
        self._f23: int = int(0)
        self._f27: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f23(self):
        return self._f23
    def ctor__(self, f27, f23):
        (self)._f27 = f27
        (self)._f23 = f23

    def fm10(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({True, not(True), False}))).issubset((_dafny.Set({False, False})) - (_dafny.Set({True, False})))

    def fm11(self, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False]) if True else _dafny.SeqWithoutIsStrInference([False])))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_674_v0_: _dafny.Array
        nw109_ = _dafny.Array(D1.default()(), 15)
        d_674_v0_ = nw109_
        d_675_v1_: _dafny.Seq
        d_675_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_674_v0_])])
        d_676_v2_: _dafny.Seq
        d_676_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ru"))
        d_677_v3_: _dafny.Map
        d_677_v3_ = _dafny.Map({p1: (d_675_v1_)[default__.safeIndex(len(d_676_v2_), len(d_675_v1_))]})
        d_678_v4_: _dafny.MultiSet
        d_678_v4_ = _dafny.MultiSet([d_674_v0_])
        d_679_v5_: D3
        d_679_v5_ = D3_DC9(d_678_v4_)
        rhs148_ = (((d_677_v3_)[p1] if (p1) in (d_677_v3_) else (d_679_v5_).cf20)).cardinality
        rhs149_ = (self).f23
        rhs150_ = (self).f23
        rhs151_ = default__.safeModuloInt((self).f23, (default__.fm0((self).f23, p1, globalState) if p1 else (self).f23))
        rhs152_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "taruujebt"))) + (d_676_v2_)) + (d_676_v2_)
        lhs139_ = globalState
        lhs140_ = globalState
        lhs141_ = globalState
        r0 = rhs148_
        lhs139_.f8 = rhs149_
        r0 = rhs150_
        lhs140_.f14 = rhs151_
        lhs141_.f6 = rhs152_
        d_680_v6_: _dafny.Map
        d_680_v6_ = _dafny.Map({p1: (self).f23})
        d_681_v7_: _dafny.Map
        d_681_v7_ = d_680_v6_
        d_682_v8_: _dafny.MultiSet
        d_682_v8_ = _dafny.MultiSet([d_681_v7_, d_681_v7_, d_681_v7_])
        d_683_v9_: _dafny.Seq
        d_683_v9_ = _dafny.SeqWithoutIsStrInference([(d_682_v8_).cardinality])
        d_684_v10_: C1
        nw110_ = C1()
        nw110_.ctor__(default__.fm1((self).f23, True, globalState), (d_683_v9_) + (d_683_v9_), (self).f23)
        d_684_v10_ = nw110_
        d_685_v11_: _dafny.Array
        def lambda29_(d_686_v9_):
            def lambda30_(d_687_i0_):
                return d_686_v9_

            return lambda30_

        init17_ = lambda29_(d_683_v9_)
        nw111_ = _dafny.Array(None, 15)
        for i0_17_ in range(nw111_.length(0)):
            nw111_[i0_17_] = init17_(i0_17_)
        d_685_v11_ = nw111_
        index138_ = default__.safeIndex(122, (d_685_v11_).length(0))
        (d_685_v11_)[index138_] = (_dafny.SeqWithoutIsStrInference([(self).f23])) + (d_683_v9_)
        index139_ = default__.safeIndex(428, (p0).length(0))
        (p0)[index139_] = ((self).f23) <= ((self).f23)
        index140_ = default__.safeIndex(122, (d_685_v11_).length(0))
        index141_ = default__.safeIndex(428, (p0).length(0))
        rhs153_ = (D9_DC27(p1, d_683_v9_, p0, p1, p1)).cf45
        rhs154_ = p1
        rhs155_ = p1
        rhs156_ = p1
        rhs157_ = ((self).f23) > (((self).f23) + ((self).f23))
        lhs142_ = d_685_v11_
        lhs143_ = default__.safeIndex(122, (d_685_v11_).length(0))
        lhs144_ = globalState
        lhs145_ = globalState
        lhs146_ = p0
        lhs147_ = default__.safeIndex(428, (p0).length(0))
        lhs148_ = globalState
        lhs142_[lhs143_] = rhs153_
        lhs144_.f7 = rhs154_
        lhs145_.f21 = rhs155_
        lhs146_[lhs147_] = rhs156_
        lhs148_.f21 = rhs157_
        d_688_i1_: int
        d_688_i1_ = 0
        with _dafny.label("9"):
            while (p0)[default__.safeIndex(428, (p0).length(0))]:
                with _dafny.c_label("9"):
                    if (d_688_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_688_i1_ = (d_688_i1_) + (1)
                    (globalState).f8 = (self).f23
                    d_689_v12_: _dafny.Array
                    nw112_ = _dafny.Array(None, 11)
                    d_689_v12_ = nw112_
                    d_690_v13_: D3
                    d_690_v13_ = D3_DC10((p0)[default__.safeIndex(428, (p0).length(0))])
                    def iife45_(_pat_let10_0):
                        def iife46_(d_691_dt__update__tmp_h0_):
                            def iife47_(_pat_let11_0):
                                def iife48_(d_692_dt__update_hcf21_h0_):
                                    return D3_DC10(d_692_dt__update_hcf21_h0_)
                                return iife48_(_pat_let11_0)
                            return iife47_((-625) > (856))
                        return iife46_(_pat_let10_0)
                    rhs158_ = d_689_v12_
                    rhs159_ = (self).f23
                    rhs160_ = iife45_(D3_DC10((p0)[default__.safeIndex(428, (p0).length(0))]))
                    lhs149_ = globalState
                    d_689_v12_ = rhs158_
                    lhs149_.f8 = rhs159_
                    d_690_v13_ = rhs160_
                    d_693_v14_: _dafny.Set
                    d_693_v14_ = _dafny.Set({(self).f27})
                    d_694_v15_: _dafny.Map
                    d_694_v15_ = _dafny.Map({d_693_v14_: p1})
                    (globalState).f21 = ((d_694_v15_)[d_693_v14_] if (d_693_v14_) in (d_694_v15_) else True)
                    d_695_v16_: _dafny.Array
                    nw113_ = _dafny.Array(int(0), 8)
                    d_695_v16_ = nw113_
                    d_695_v16_ = d_695_v16_
                    pass
            pass
        d_696_v17_: _dafny.Array
        def lambda31_(d_697_p1_, d_698_p0_):
            def lambda32_(d_699_i2_):
                return (_dafny.Map({_dafny.Set({False, d_697_p1_, (d_698_p0_)[default__.safeIndex(428, (d_698_p0_).length(0))]}): (self).f23})) | (_dafny.Map({_dafny.Set({not((d_698_p0_)[default__.safeIndex(428, (d_698_p0_).length(0))]), (d_698_p0_)[default__.safeIndex(428, (d_698_p0_).length(0))], True, (d_698_p0_)[default__.safeIndex(428, (d_698_p0_).length(0))], d_697_p1_}): (self).f23}))

            return lambda32_

        init18_ = lambda31_(p1, p0)
        nw114_ = _dafny.Array(None, 16)
        for i0_18_ in range(nw114_.length(0)):
            nw114_[i0_18_] = init18_(i0_18_)
        d_696_v17_ = nw114_
        d_700_v18_: _dafny.Set
        d_700_v18_ = _dafny.Set({not((p0)[default__.safeIndex(428, (p0).length(0))])})
        index142_ = default__.safeIndex(476, (d_696_v17_).length(0))
        (d_696_v17_)[index142_] = _dafny.Map({d_700_v18_: (self).f23})
        d_701_v20_: _dafny.Map
        d_701_v20_ = _dafny.Map({d_700_v18_: d_683_v9_})
        index143_ = default__.safeIndex(476, (d_696_v17_).length(0))
        def iife49_():
            coll25_ = _dafny.Map()
            compr_25_: _dafny.Set
            for compr_25_ in (d_701_v20_).keys.Elements:
                d_702_v19_: _dafny.Set = compr_25_
                if (d_702_v19_) in (d_701_v20_):
                    coll25_[d_702_v19_] = (self).f23
            return _dafny.Map(coll25_)
        (d_696_v17_)[index143_] = iife49_()
        
        if (p0)[default__.safeIndex(428, (p0).length(0))]:
            d_703_v21_: _dafny.Map
            d_703_v21_ = _dafny.Map({(self).f23: (self).f23})
            r0 = default__.fm0(((d_703_v21_)[len(d_676_v2_)] if (len(d_676_v2_)) in (d_703_v21_) else (self).f23), False, globalState)
            (globalState).f9 = (p0)[default__.safeIndex(428, (p0).length(0))]
            d_704_v22_: _dafny.MultiSet
            d_704_v22_ = _dafny.MultiSet([(p0)[default__.safeIndex(428, (p0).length(0))], not(p1), not(p1)])
            d_705_v23_: _dafny.Seq
            d_705_v23_ = _dafny.SeqWithoutIsStrInference([d_704_v22_])
            d_706_v24_: T0
            nw115_ = C3()
            nw115_.ctor__(d_705_v23_, default__.fm0((self).f23, p1, globalState), (p0)[default__.safeIndex(428, (p0).length(0))], (d_685_v11_)[default__.safeIndex(122, (d_685_v11_).length(0))])
            d_706_v24_ = nw115_
            d_707_v25_: _dafny.Map
            d_707_v25_ = _dafny.Map({d_706_v24_: (p0)[default__.safeIndex(428, (p0).length(0))]})
            d_708_v26_: C3
            nw116_ = C3()
            nw116_.ctor__((d_705_v23_ if (p0)[default__.safeIndex(428, (p0).length(0))] else d_705_v23_), ((0) - ((self).f23)) - (len(d_707_v25_)), (p0)[default__.safeIndex(428, (p0).length(0))], (d_685_v11_)[default__.safeIndex(122, (d_685_v11_).length(0))])
            d_708_v26_ = nw116_
            d_709_v29_: _dafny.Seq
            d_709_v29_ = _dafny.SeqWithoutIsStrInference([d_703_v21_, d_703_v21_])
            d_710_v30_: _dafny.Set
            d_710_v30_ = _dafny.Set({(self).f23, len(d_703_v21_)})
            d_711_v31_: _dafny.Array
            nw117_ = _dafny.Array(None, 16)
            nw117_[int(0)] = d_703_v21_
            nw117_[int(1)] = d_703_v21_
            nw117_[int(2)] = d_703_v21_
            nw117_[int(3)] = d_703_v21_
            nw117_[int(4)] = (d_703_v21_ if default__.fm1(((d_704_v22_)[(p0)[default__.safeIndex(428, (p0).length(0))]] if ((p0)[default__.safeIndex(428, (p0).length(0))]) in (d_704_v22_) else 651), (p0)[default__.safeIndex(428, (p0).length(0))], globalState) else d_703_v21_)
            nw117_[int(5)] = d_703_v21_
            nw117_[int(6)] = d_703_v21_
            nw117_[int(7)] = d_703_v21_
            nw117_[int(8)] = d_703_v21_
            nw117_[int(9)] = d_703_v21_
            nw117_[int(10)] = (d_703_v21_) | (_dafny.Map({len(d_703_v21_): (d_706_v24_).f23}))
            def iife50_():
                def iife52_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(118, 492):
                        d_712_v28_: int = compr_28_
                        if ((118) <= (d_712_v28_)) and ((d_712_v28_) < (492)):
                            coll28_[default__.safeDivisionInt(d_712_v28_, len(_dafny.Set({D3_DC10(not(p1))})))] = _dafny.MultiSet([len(d_700_v18_)])
                    return _dafny.Map(coll28_)
                coll26_ = _dafny.Map()
                def iife51_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(118, 492):
                        d_712_v28_: int = compr_27_
                        if ((118) <= (d_712_v28_)) and ((d_712_v28_) < (492)):
                            coll27_[default__.safeDivisionInt(d_712_v28_, len(_dafny.Set({D3_DC10(not(p1))})))] = _dafny.MultiSet([len(d_700_v18_)])
                    return _dafny.Map(coll27_)
                compr_26_: int
                for compr_26_ in (iife51_()
                ).keys.Elements:
                    d_713_v27_: int = compr_26_
                    if (d_713_v27_) in (iife52_()
                    ):
                        coll26_[(d_713_v27_) - ((d_706_v24_).f23)] = (d_706_v24_).f23
                return _dafny.Map(coll26_)
            nw117_[int(11)] = iife50_()
            
            nw117_[int(12)] = (d_709_v29_)[default__.safeIndex(len(d_676_v2_), len(d_709_v29_))]
            nw117_[int(13)] = (_dafny.Map({(0) - ((d_706_v24_).f23): (d_706_v24_).f23})).set(len(d_710_v30_), (self).f23)
            nw117_[int(14)] = d_703_v21_
            nw117_[int(15)] = d_703_v21_
            d_711_v31_ = nw117_
            d_714_v32_: _dafny.Seq
            d_714_v32_ = _dafny.SeqWithoutIsStrInference([(self).fm10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_715_i3_ in range(default__.abs(868))]), False, d_680_v6_, (p0)[default__.safeIndex(428, (p0).length(0))], globalState), p1, p1, p1, not((p0)[default__.safeIndex(428, (p0).length(0))])])
            index144_ = default__.safeIndex(32, ((self).f27).length(0))
            ((self).f27)[index144_] = (0) - (default__.fm0(((d_685_v11_)[default__.safeIndex(122, (d_685_v11_).length(0))])[default__.safeIndex(len(d_714_v32_), len((d_685_v11_)[default__.safeIndex(122, (d_685_v11_).length(0))]))], (p0)[default__.safeIndex(428, (p0).length(0))], globalState))
            d_716_v33_: D11
            d_716_v33_ = D11_DC30(d_711_v31_)
            index145_ = default__.safeIndex(32, ((self).f27).length(0))
            rhs161_ = (d_706_v24_).f23
            rhs162_ = (d_716_v33_).cf51
            rhs163_ = (self).f23
            rhs164_ = (self).f23
            lhs150_ = globalState
            lhs151_ = (self).f27
            lhs152_ = default__.safeIndex(32, ((self).f27).length(0))
            lhs153_ = globalState
            lhs150_.f14 = rhs161_
            d_711_v31_ = rhs162_
            lhs151_[lhs152_] = rhs163_
            lhs153_.f14 = rhs164_
            (globalState).f14 = len(d_683_v9_)
        elif True:
            d_717_v34_: _dafny.Map
            d_717_v34_ = _dafny.Map({default__.fm0((self).f23, (p0)[default__.safeIndex(428, (p0).length(0))], globalState): default__.fm6((0) - ((self).f23), globalState)})
            (globalState).f8 = (0) - (len(d_717_v34_))
            d_718_v35_: _dafny.Map
            d_718_v35_ = _dafny.Map({(self).f23: (self).f23})
            d_719_v36_: C2
            nw118_ = C2()
            nw118_.ctor__(d_718_v35_, (p0)[default__.safeIndex(428, (p0).length(0))], (len(d_676_v2_)) != ((0) - ((self).f23)), (d_683_v9_) + (_dafny.SeqWithoutIsStrInference([(self).f23 for d_720_i4_ in range(default__.abs(397))])), (self).f23)
            d_719_v36_ = nw118_
            d_721_v37_: _dafny.Array
            nw119_ = _dafny.Array(False, 21)
            d_721_v37_ = nw119_
            d_721_v37_ = p0
            d_722_v38_: _dafny.Set
            d_722_v38_ = _dafny.Set({d_719_v36_})
            d_723_v39_: _dafny.Seq
            d_723_v39_ = _dafny.SeqWithoutIsStrInference([p1, (p0)[default__.safeIndex(428, (p0).length(0))], default__.fm1((self).f23, (d_719_v36_).f32, globalState), p1, (d_722_v38_).isdisjoint(d_722_v38_)])
            if (d_723_v39_)[default__.safeIndex((0) - ((self).f23), len(d_723_v39_))]:
                index146_ = default__.safeIndex(428, (p0).length(0))
                (p0)[index146_] = ((self).f23) < ((self).f23)
                d_724_v40_: _dafny.Array
                nw120_ = _dafny.Array(None, 24)
                nw120_[int(0)] = d_680_v6_
                nw120_[int(1)] = _dafny.Map({(d_719_v36_).f32: (self).f23})
                nw120_[int(2)] = _dafny.Map({(p0)[default__.safeIndex(428, (p0).length(0))]: (self).f23})
                nw120_[int(3)] = (d_680_v6_) | (d_680_v6_)
                nw120_[int(4)] = d_680_v6_
                nw120_[int(5)] = (d_680_v6_).set((d_719_v36_).f32, (self).f23)
                nw120_[int(6)] = d_680_v6_
                nw120_[int(7)] = d_680_v6_
                nw120_[int(8)] = (d_680_v6_ if (d_719_v36_).f32 else _dafny.Map({(d_719_v36_).f32: (self).f23}))
                nw120_[int(9)] = d_680_v6_
                nw120_[int(10)] = (d_680_v6_) | (d_680_v6_)
                nw120_[int(11)] = (d_680_v6_) | (d_680_v6_)
                nw120_[int(12)] = ((d_680_v6_).set(False, (self).f23)) | (d_680_v6_)
                nw120_[int(13)] = d_680_v6_
                nw120_[int(14)] = (d_680_v6_) | (d_680_v6_)
                nw120_[int(15)] = (_dafny.Map({p1: (self).f23})).set((d_719_v36_).f32, len(d_676_v2_))
                nw120_[int(16)] = d_680_v6_
                nw120_[int(17)] = d_680_v6_
                nw120_[int(18)] = d_680_v6_
                nw120_[int(19)] = (d_680_v6_).set((p0)[default__.safeIndex(428, (p0).length(0))], len(d_680_v6_))
                nw120_[int(20)] = d_680_v6_
                nw120_[int(21)] = d_680_v6_
                nw120_[int(22)] = d_680_v6_
                nw120_[int(23)] = _dafny.Map({(p0)[default__.safeIndex(428, (p0).length(0))]: (self).f23})
                d_724_v40_ = nw120_
                index147_ = default__.safeIndex(449, (d_724_v40_).length(0))
                (d_724_v40_)[index147_] = d_680_v6_
                d_725_v41_: str
                d_725_v41_ = _dafny.CodePoint('v')
                d_726_v42_: _dafny.Map
                d_726_v42_ = _dafny.Map({d_725_v41_: p1})
                d_727_v43_: D12
                d_727_v43_ = D12_DC34(d_726_v42_)
                index148_ = default__.safeIndex(449, (d_724_v40_).length(0))
                (d_724_v40_)[index148_] = (_dafny.Map({p1: len((d_727_v43_).cf55)})) | (d_680_v6_)
                index149_ = default__.safeIndex(259, ((self).f27).length(0))
                ((self).f27)[index149_] = (self).f23
                d_728_v44_: _dafny.Map
                d_728_v44_ = _dafny.Map({(self).f23: (p0)[default__.safeIndex(428, (p0).length(0))]})
                d_729_v45_: D8
                d_729_v45_ = D8_DC21((d_719_v36_).f32, (p0)[default__.safeIndex(428, (p0).length(0))], len(d_728_v44_))
                index150_ = default__.safeIndex(259, ((self).f27).length(0))
                ((self).f27)[index150_] = default__.safeDivisionInt(798, ((d_729_v45_).cf37) * ((self).f23))
                d_730_v46_: bool
                d_731_v47_: _dafny.Map
                d_732_v48_: D1
                d_733_v49_: _dafny.Array
                out30_: bool
                out31_: _dafny.Map
                out32_: D1
                out33_: _dafny.Array
                out30_, out31_, out32_, out33_ = (d_719_v36_).m6((_dafny.Set({True, True, (d_719_v36_).f32})).issubset(d_700_v18_), globalState)
                d_730_v46_ = out30_
                d_731_v47_ = out31_
                d_732_v48_ = out32_
                d_733_v49_ = out33_
                index151_ = default__.safeIndex(428, (p0).length(0))
                (p0)[index151_] = p1
            elif True:
                index152_ = default__.safeIndex(428, (p0).length(0))
                (p0)[index152_] = ((self).f23) > ((33) * (127))
                index153_ = default__.safeIndex(914, ((self).f27).length(0))
                ((self).f27)[index153_] = (self).f23
                index154_ = default__.safeIndex(914, ((self).f27).length(0))
                ((self).f27)[index154_] = default__.safeDivisionInt((self).f23, default__.safeModuloInt((self).f23, (self).f23))
                d_734_v50_: str
                d_734_v50_ = _dafny.CodePoint('n')
                (globalState).f21 = ((d_676_v2_).set(default__.safeIndex(((self).f27)[default__.safeIndex(914, ((self).f27).length(0))], len(d_676_v2_)), d_734_v50_)) <= ((d_676_v2_) + (d_676_v2_))
                d_735_v51_: D3
                d_735_v51_ = D3_DC11(453)
                index155_ = default__.safeIndex(428, (p0).length(0))
                (p0)[index155_] = (d_723_v39_)[default__.safeIndex((0) - ((((self).f27)[default__.safeIndex(914, ((self).f27).length(0))] if (d_719_v36_).f32 else (d_735_v51_).cf22)), len(d_723_v39_))]
                (globalState).f21 = False
            d_736_v52_: _dafny.Set
            d_736_v52_ = _dafny.Set({(self).f23})
            d_737_v53_: C2
            nw121_ = C2()
            nw121_.ctor__(((d_719_v36_).f31).set((self).f23, 797), False, (self).fm10(d_676_v2_, p1, (d_680_v6_).set((p0)[default__.safeIndex(428, (p0).length(0))], (self).f23), default__.fm1((0) - ((self).f23), (d_719_v36_).f32, globalState), globalState), d_683_v9_, default__.safeModuloInt((0) - (len(d_736_v52_)), (0) - ((self).f23)))
            d_737_v53_ = nw121_
        r0 = (self).f23
        d_738_v54_: _dafny.Set
        d_738_v54_ = _dafny.Set({len(d_680_v6_), (self).f23})
        r1 = (d_738_v54_) | ((d_738_v54_).intersection(d_738_v54_))
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        (globalState).f8 = (self).f23
        d_739_v0_: bool
        d_739_v0_ = False
        d_740_i0_: int
        d_740_i0_ = 0
        with _dafny.label("10"):
            while (((self).fm10(p0, True, _dafny.Map({d_739_v0_: (self).f23}), d_739_v0_, globalState) if False else d_739_v0_)) or (d_739_v0_):
                with _dafny.c_label("10"):
                    if (d_740_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_740_i0_ = (d_740_i0_) + (1)
                    index156_ = default__.safeIndex(562, ((self).f27).length(0))
                    ((self).f27)[index156_] = default__.safeDivisionInt((self).f23, (self).f23)
                    d_741_v1_: _dafny.Seq
                    d_741_v1_ = _dafny.SeqWithoutIsStrInference([d_739_v0_])
                    d_742_v2_: _dafny.Array
                    nw122_ = _dafny.Array(None, 3)
                    nw122_[int(0)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uovdktobe"))) + (p0))
                    nw122_[int(1)] = default__.safeModuloInt(len(d_741_v1_), (self).f23)
                    nw122_[int(2)] = (self).f23
                    d_742_v2_ = nw122_
                    d_743_v3_: _dafny.Seq
                    d_743_v3_ = _dafny.SeqWithoutIsStrInference([(self).f23])
                    d_744_v4_: _dafny.Map
                    d_744_v4_ = _dafny.Map({(self).f23: len(d_743_v3_)})
                    d_745_v5_: _dafny.Set
                    d_745_v5_ = _dafny.Set({d_739_v0_})
                    d_746_v6_: _dafny.Map
                    d_746_v6_ = _dafny.Map({not(d_739_v0_): len((D13_DC38(d_745_v5_)).cf62)})
                    d_747_v7_: C2
                    nw123_ = C2()
                    nw123_.ctor__(d_744_v4_, (self).fm10(p0, d_739_v0_, d_746_v6_, True, globalState), d_739_v0_, d_743_v3_, 465)
                    d_747_v7_ = nw123_
                    index157_ = default__.safeIndex(562, ((self).f27).length(0))
                    rhs165_ = ((self).f23) - ((self).f23)
                    rhs166_ = d_739_v0_
                    rhs167_ = (self).f27
                    rhs168_ = (d_747_v7_ if d_739_v0_ else d_747_v7_)
                    lhs154_ = (self).f27
                    lhs155_ = default__.safeIndex(562, ((self).f27).length(0))
                    lhs156_ = globalState
                    lhs154_[lhs155_] = rhs165_
                    lhs156_.f21 = rhs166_
                    d_742_v2_ = rhs167_
                    d_747_v7_ = rhs168_
                    d_748_v8_: str
                    d_748_v8_ = _dafny.CodePoint('n')
                    d_749_v9_: _dafny.Set
                    d_749_v9_ = _dafny.Set({(D8_DC20(d_739_v0_, ((self).f27)[default__.safeIndex(562, ((self).f27).length(0))])).cf34, len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxmrtsdjn"))).set(default__.safeIndex(((self).f27)[default__.safeIndex(562, ((self).f27).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxmrtsdjn")))), d_748_v8_)).set(default__.safeIndex(512, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxmrtsdjn"))).set(default__.safeIndex(((self).f27)[default__.safeIndex(562, ((self).f27).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxmrtsdjn")))), d_748_v8_))), d_748_v8_))})
                    d_750_v10_: _dafny.Map
                    d_750_v10_ = _dafny.Map({d_739_v0_: d_739_v0_})
                    d_751_v11_: _dafny.Map
                    d_751_v11_ = _dafny.Map({len(d_749_v9_): ((d_750_v10_)[(d_747_v7_).f32] if ((d_747_v7_).f32) in (d_750_v10_) else d_739_v0_)})
                    d_752_v12_: _dafny.Array
                    nw124_ = _dafny.Array(None, 12)
                    nw124_[int(0)] = True
                    nw124_[int(1)] = (d_747_v7_).f32
                    nw124_[int(2)] = True
                    nw124_[int(3)] = (d_747_v7_).f32
                    nw124_[int(4)] = d_739_v0_
                    nw124_[int(5)] = ((d_747_v7_).f32) or (d_739_v0_)
                    nw124_[int(6)] = (d_741_v1_) != (d_741_v1_)
                    nw124_[int(7)] = d_739_v0_
                    nw124_[int(8)] = (d_739_v0_) == (d_739_v0_)
                    nw124_[int(9)] = ((d_747_v7_).f32 if (d_747_v7_).f32 else (d_747_v7_).f32)
                    nw124_[int(10)] = False
                    nw124_[int(11)] = (((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]) == (len(_dafny.SeqWithoutIsStrInference([d_739_v0_, d_739_v0_, (d_747_v7_).f32, ((d_751_v11_)[(0) - ((self).f23)] if ((0) - ((self).f23)) in (d_751_v11_) else (d_747_v7_).f32)])))
                    d_752_v12_ = nw124_
                    d_752_v12_ = d_752_v12_
                    d_753_v13_: D9
                    d_753_v13_ = D9_DC27((len(d_749_v9_)) != ((self).f23), d_743_v3_, d_752_v12_, d_739_v0_, (d_747_v7_).f32)
                    d_754_v14_: _dafny.Map
                    d_754_v14_ = _dafny.Map({(d_747_v7_).f32: 957})
                    index158_ = default__.safeIndex(562, ((self).f27).length(0))
                    rhs169_ = default__.fm0(len(p0), ((d_751_v11_)[((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]] if (((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]) in (d_751_v11_) else ((d_750_v10_)[(d_747_v7_).f32] if ((d_747_v7_).f32) in (d_750_v10_) else (self).fm10(p0, d_739_v0_, (d_754_v14_), d_739_v0_, globalState))), globalState)
                    rhs170_ = d_753_v13_
                    lhs157_ = (self).f27
                    lhs158_ = default__.safeIndex(562, ((self).f27).length(0))
                    lhs157_[lhs158_] = rhs169_
                    d_753_v13_ = rhs170_
                    d_755_v15_: _dafny.Map
                    d_755_v15_ = _dafny.Map({((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]: ((d_741_v1_) + (d_741_v1_)).set(default__.safeIndex(len(p0), len((d_741_v1_) + (d_741_v1_))), (d_747_v7_).f32)})
                    rhs171_ = ((self).f23) == (((_dafny.MultiSet([((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]])).cardinality) + (((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]))
                    rhs172_ = ((_dafny.Map({((self).f27)[default__.safeIndex(562, ((self).f27).length(0))]: d_741_v1_})) | (d_755_v15_)) | ((d_755_v15_).set(len(d_741_v1_), (self).fm11(globalState)))
                    rhs173_ = (0) - ((self).f23)
                    lhs159_ = globalState
                    lhs159_.f9 = rhs171_
                    d_755_v15_ = rhs172_
                    r0 = rhs173_
                    pass
            pass
        d_756_v16_: _dafny.Set
        d_756_v16_ = _dafny.Set({not(d_739_v0_), not(not(d_739_v0_))})
        def lambda33_(source6_):
            if source6_.is_DC1:
                d_757___mcc_h0_ = source6_.cf1
                d_758___mcc_h1_ = source6_.cf2
                d_759___mcc_h2_ = source6_.cf3
                d_760_cf3_ = d_759___mcc_h2_
                d_761_cf2_ = d_758___mcc_h1_
                d_762_cf1_ = d_757___mcc_h0_
                return (d_760_cf3_) - (d_762_cf1_)
            elif source6_.is_DC0:
                d_763___mcc_h3_ = source6_.cf0
                d_764_cf0_ = d_763___mcc_h3_
                return (11) - ((self).f23)
            elif True:
                d_765___mcc_h4_ = source6_.cf4
                d_766_cf4_ = d_765___mcc_h4_
                return (D9_DC26((self).f23)).cf43

        (globalState).f8 = (0) - (lambda33_(default__.fm26(d_756_v16_, globalState)))
        d_767_i1_: int
        d_767_i1_ = 0
        with _dafny.label("11"):
            while not(d_739_v0_):
                with _dafny.c_label("11"):
                    if (d_767_i1_) >= (100):
                        raise _dafny.Break("11")
                    d_767_i1_ = (d_767_i1_) + (1)
                    if d_739_v0_:
                        d_768_v17_: str
                        d_768_v17_ = _dafny.CodePoint('s')
                        d_769_v18_: _dafny.MultiSet
                        d_769_v18_ = _dafny.MultiSet([d_739_v0_, d_739_v0_, False, (d_739_v0_) or (default__.fm1(len(_dafny.SeqWithoutIsStrInference([d_768_v17_])), d_739_v0_, globalState)), ((self).f23) >= ((self).f23)])
                        d_769_v18_ = (_dafny.MultiSet([d_739_v0_, d_739_v0_])) - (d_769_v18_)
                        d_770_v19_: _dafny.Map
                        d_770_v19_ = _dafny.Map({((self).f23) == ((self).f23): d_768_v17_})
                        d_771_v20_: _dafny.Seq
                        d_771_v20_ = _dafny.SeqWithoutIsStrInference([(self).f23, 419, (self).f23])
                        d_772_v21_: _dafny.Map
                        d_772_v21_ = _dafny.Map({d_739_v0_: _dafny.SeqWithoutIsStrInference([(self).f23 for d_773_i2_ in range(default__.abs(-174))])})
                        d_774_v22_: D12
                        d_774_v22_ = D12_DC36(d_739_v0_, True, (self).f23, (self).f23, d_772_v21_)
                        d_775_v23_: _dafny.Array
                        nw125_ = _dafny.Array(None, 26)
                        nw125_[int(0)] = (597) >= ((self).f23)
                        nw125_[int(1)] = d_739_v0_
                        nw125_[int(2)] = d_739_v0_
                        nw125_[int(3)] = d_739_v0_
                        nw125_[int(4)] = not (d_739_v0_) or (d_739_v0_)
                        nw125_[int(5)] = True
                        nw125_[int(6)] = not (d_739_v0_) or (d_739_v0_)
                        nw125_[int(7)] = d_739_v0_
                        nw125_[int(8)] = True
                        nw125_[int(9)] = d_739_v0_
                        nw125_[int(10)] = ((self).f23) > ((0) - ((self).f23))
                        nw125_[int(11)] = d_739_v0_
                        nw125_[int(12)] = not(d_739_v0_)
                        nw125_[int(13)] = d_739_v0_
                        nw125_[int(14)] = d_739_v0_
                        nw125_[int(15)] = (d_768_v17_) not in (p0)
                        nw125_[int(16)] = d_739_v0_
                        nw125_[int(17)] = d_739_v0_
                        nw125_[int(18)] = not((d_771_v20_) == (d_771_v20_))
                        nw125_[int(19)] = (d_739_v0_) or (d_739_v0_)
                        nw125_[int(20)] = (d_774_v22_).cf57
                        nw125_[int(21)] = d_739_v0_
                        nw125_[int(22)] = not(True)
                        nw125_[int(23)] = d_739_v0_
                        nw125_[int(24)] = d_739_v0_
                        nw125_[int(25)] = d_739_v0_
                        d_775_v23_ = nw125_
                        index159_ = default__.safeIndex(36, (d_775_v23_).length(0))
                        (d_775_v23_)[index159_] = not(d_739_v0_)
                        d_776_v24_: _dafny.Array
                        def lambda34_(d_777_i3_):
                            return _dafny.CodePoint('o')

                        init19_ = lambda34_
                        nw126_ = _dafny.Array(None, 6)
                        for i0_19_ in range(nw126_.length(0)):
                            nw126_[i0_19_] = init19_(i0_19_)
                        d_776_v24_ = nw126_
                        index160_ = default__.safeIndex(605, (d_776_v24_).length(0))
                        (d_776_v24_)[index160_] = d_768_v17_
                        index161_ = default__.safeIndex(36, (d_775_v23_).length(0))
                        index162_ = default__.safeIndex(605, (d_776_v24_).length(0))
                        rhs174_ = (d_770_v19_) | (d_770_v19_)
                        rhs175_ = not (d_739_v0_) or (not(d_739_v0_))
                        rhs176_ = d_775_v23_
                        rhs177_ = d_768_v17_
                        lhs160_ = d_775_v23_
                        lhs161_ = default__.safeIndex(36, (d_775_v23_).length(0))
                        lhs162_ = d_776_v24_
                        lhs163_ = default__.safeIndex(605, (d_776_v24_).length(0))
                        d_770_v19_ = rhs174_
                        lhs160_[lhs161_] = rhs175_
                        d_775_v23_ = rhs176_
                        lhs162_[lhs163_] = rhs177_
                        index163_ = default__.safeIndex(36, (d_775_v23_).length(0))
                        (d_775_v23_)[index163_] = not (default__.fm1((self).f23, not((d_775_v23_)[default__.safeIndex(36, (d_775_v23_).length(0))]), globalState)) or (d_739_v0_)
                        d_778_v25_: _dafny.Map
                        d_778_v25_ = _dafny.Map({d_739_v0_: False})
                        d_779_v26_: _dafny.Seq
                        d_779_v26_ = _dafny.SeqWithoutIsStrInference([d_769_v18_, d_769_v18_, ((d_769_v18_).set((d_775_v23_)[default__.safeIndex(36, (d_775_v23_).length(0))], default__.abs(len(d_778_v25_)))).set((d_775_v23_)[default__.safeIndex(36, (d_775_v23_).length(0))], default__.abs(678)), d_769_v18_, d_769_v18_])
                        d_780_v27_: C3
                        nw127_ = C3()
                        nw127_.ctor__(d_779_v26_, (self).f23, not((d_775_v23_)[default__.safeIndex(36, (d_775_v23_).length(0))]), d_771_v20_)
                        d_780_v27_ = nw127_
                        (globalState).f21 = d_739_v0_
                    elif True:
                        d_781_v28_: str
                        d_781_v28_ = _dafny.CodePoint('x')
                        (globalState).f11 = d_781_v28_
                        d_782_v29_: _dafny.Array
                        nw128_ = _dafny.Array(_dafny.Array(None, 0), 15)
                        d_782_v29_ = nw128_
                        d_782_v29_ = d_782_v29_
                        d_783_v31_: _dafny.Map
                        d_783_v31_ = _dafny.Map({d_739_v0_: (self).f23})
                        d_784_v32_: _dafny.Set
                        d_784_v32_ = _dafny.Set({default__.fm0((self).f23, d_739_v0_, globalState), (self).f23, (self).f23, default__.fm0((0) - ((self).f23), (self).fm10(_dafny.SeqWithoutIsStrInference([d_781_v28_ for d_785_i4_ in range(default__.abs(532))]), d_739_v0_, d_783_v31_, d_739_v0_, globalState), globalState), (self).f23})
                        d_786_v33_: _dafny.Array
                        nw129_ = _dafny.Array(None, 2)
                        def iife53_():
                            coll29_ = _dafny.Set()
                            compr_29_: int
                            for compr_29_ in _dafny.IntegerRange(397, -231):
                                d_787_v30_: int = compr_29_
                                if ((397) <= (d_787_v30_)) and ((d_787_v30_) < (-231)):
                                    coll29_ = coll29_.union(_dafny.Set([(d_787_v30_) - ((self).f23)]))
                            return _dafny.Set(coll29_)
                        nw129_[int(0)] = (iife53_()
                        ) | (_dafny.Set({(self).f23}))
                        nw129_[int(1)] = d_784_v32_
                        d_786_v33_ = nw129_
                        index164_ = default__.safeIndex(754, (d_786_v33_).length(0))
                        (d_786_v33_)[index164_] = d_784_v32_
                        index165_ = default__.safeIndex(754, (d_786_v33_).length(0))
                        (d_786_v33_)[index165_] = d_784_v32_
                        d_788_v34_: _dafny.Map
                        d_788_v34_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(d_739_v0_)])): d_739_v0_})
                        d_788_v34_ = (d_788_v34_) | (d_788_v34_)
                        d_789_v35_: _dafny.Map
                        d_789_v35_ = _dafny.Map({(self).f23: (self).f23})
                        d_789_v35_ = (d_789_v35_).set(default__.fm0(-241, not(d_739_v0_), globalState), (self).f23)
                    d_790_v36_: _dafny.Seq
                    d_790_v36_ = _dafny.SeqWithoutIsStrInference([(self).f23])
                    d_791_v37_: D8
                    d_791_v37_ = D8_DC20(d_739_v0_, (self).f23)
                    d_792_v38_: _dafny.Seq
                    d_792_v38_ = _dafny.SeqWithoutIsStrInference([d_739_v0_, d_739_v0_, (d_791_v37_).cf33, d_739_v0_, d_739_v0_])
                    d_793_v39_: _dafny.Array
                    nw130_ = _dafny.Array(None, 18)
                    nw130_[int(0)] = default__.fm1((self).f23, True, globalState)
                    nw130_[int(1)] = d_739_v0_
                    nw130_[int(2)] = (d_756_v16_).issubset(d_756_v16_)
                    nw130_[int(3)] = d_739_v0_
                    nw130_[int(4)] = not(d_739_v0_)
                    nw130_[int(5)] = (len(d_790_v36_)) >= ((self).f23)
                    nw130_[int(6)] = d_739_v0_
                    nw130_[int(7)] = (d_792_v38_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hm"))), len(d_792_v38_))]
                    nw130_[int(8)] = d_739_v0_
                    nw130_[int(9)] = d_739_v0_
                    nw130_[int(10)] = d_739_v0_
                    nw130_[int(11)] = d_739_v0_
                    nw130_[int(12)] = False
                    nw130_[int(13)] = d_739_v0_
                    nw130_[int(14)] = default__.fm1(len(p0), d_739_v0_, globalState)
                    nw130_[int(15)] = not(not(not(d_739_v0_)))
                    nw130_[int(16)] = not (d_739_v0_) or (d_739_v0_)
                    nw130_[int(17)] = d_739_v0_
                    d_793_v39_ = nw130_
                    index166_ = default__.safeIndex(808, (d_793_v39_).length(0))
                    (d_793_v39_)[index166_] = d_739_v0_
                    index167_ = default__.safeIndex(808, (d_793_v39_).length(0))
                    (d_793_v39_)[index167_] = not(d_739_v0_)
                    if d_739_v0_:
                        d_794_v40_: _dafny.Map
                        d_794_v40_ = _dafny.Map({d_790_v36_: d_739_v0_})
                        d_794_v40_ = ((d_794_v40_) | (d_794_v40_)) | (d_794_v40_)
                        d_795_v41_: _dafny.Map
                        d_795_v41_ = _dafny.Map({default__.fm1((self).f23, (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], globalState): ((D0_DC1(967, (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], (self).f23)).cf3) <= (872)})
                        (globalState).f7 = ((d_795_v41_)[((0) - ((self).f23)) <= ((self).f23)] if (((0) - ((self).f23)) <= ((self).f23)) in (d_795_v41_) else ((self).f23) <= (341))
                        (globalState).f14 = default__.fm0((self).f23, (80) <= ((self).f23), globalState)
                        d_796_v42_: _dafny.Map
                        d_796_v42_ = _dafny.Map({(d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]: (self).f23})
                        d_797_v43_: _dafny.Map
                        d_797_v43_ = (d_796_v42_) | (d_796_v42_)
                        d_797_v43_ = d_797_v43_
                        (globalState).f21 = not (True) or (d_739_v0_)
                    elif True:
                        d_798_v44_: D0
                        d_799_v45_: _dafny.Map
                        d_800_v46_: int
                        out34_: D0
                        out35_: _dafny.Map
                        out36_: int
                        out34_, out35_, out36_ = default__.m0(d_793_v39_, p0, (self).f23, globalState)
                        d_798_v44_ = out34_
                        d_799_v45_ = out35_
                        d_800_v46_ = out36_
                        d_801_v47_: _dafny.Array
                        nw131_ = _dafny.Array(_dafny.Seq({}), 17)
                        d_801_v47_ = nw131_
                        d_802_v48_: C1
                        nw132_ = C1()
                        nw132_.ctor__((d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], d_790_v36_, (self).f23)
                        d_802_v48_ = nw132_
                        d_803_v49_: _dafny.Map
                        d_803_v49_ = _dafny.Map({d_802_v48_: d_801_v47_})
                        d_804_v50_: _dafny.Array
                        nw133_ = _dafny.Array(None, 23)
                        nw133_[int(0)] = d_801_v47_
                        nw133_[int(1)] = d_801_v47_
                        nw133_[int(2)] = d_801_v47_
                        nw133_[int(3)] = d_801_v47_
                        nw133_[int(4)] = d_801_v47_
                        nw133_[int(5)] = (d_801_v47_ if (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))] else ((d_803_v49_)[d_802_v48_] if (d_802_v48_) in (d_803_v49_) else d_801_v47_))
                        nw133_[int(6)] = d_801_v47_
                        nw133_[int(7)] = d_801_v47_
                        nw133_[int(8)] = d_801_v47_
                        nw133_[int(9)] = d_801_v47_
                        nw133_[int(10)] = d_801_v47_
                        nw133_[int(11)] = d_801_v47_
                        nw133_[int(12)] = d_801_v47_
                        nw133_[int(13)] = d_801_v47_
                        nw133_[int(14)] = d_801_v47_
                        nw133_[int(15)] = d_801_v47_
                        nw133_[int(16)] = d_801_v47_
                        nw133_[int(17)] = d_801_v47_
                        nw133_[int(18)] = d_801_v47_
                        nw133_[int(19)] = d_801_v47_
                        nw133_[int(20)] = d_801_v47_
                        nw133_[int(21)] = d_801_v47_
                        nw133_[int(22)] = d_801_v47_
                        d_804_v50_ = nw133_
                        index168_ = default__.safeIndex(507, (d_804_v50_).length(0))
                        (d_804_v50_)[index168_] = d_801_v47_
                        d_805_v51_: str
                        d_805_v51_ = _dafny.CodePoint('r')
                        d_806_v52_: _dafny.MultiSet
                        d_806_v52_ = _dafny.MultiSet([d_805_v51_, _dafny.CodePoint('o'), d_805_v51_])
                        index169_ = default__.safeIndex(136, ((self).f27).length(0))
                        ((self).f27)[index169_] = ((d_806_v52_)[d_805_v51_] if (d_805_v51_) in (d_806_v52_) else len(_dafny.SeqWithoutIsStrInference([(d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]])))
                        index170_ = default__.safeIndex(507, (d_804_v50_).length(0))
                        index171_ = default__.safeIndex(136, ((self).f27).length(0))
                        rhs178_ = d_801_v47_
                        rhs179_ = ((0) - (d_800_v46_)) + (((self).f23 if (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))] else (self).f23))
                        lhs164_ = d_804_v50_
                        lhs165_ = default__.safeIndex(507, (d_804_v50_).length(0))
                        lhs166_ = (self).f27
                        lhs167_ = default__.safeIndex(136, ((self).f27).length(0))
                        lhs164_[lhs165_] = rhs178_
                        lhs166_[lhs167_] = rhs179_
                        d_807_v53_: _dafny.MultiSet
                        d_807_v53_ = _dafny.MultiSet([(self).f23])
                        d_808_v54_: _dafny.Seq
                        d_808_v54_ = _dafny.SeqWithoutIsStrInference([p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxaeugrwk"))])
                        d_809_v55_: _dafny.Map
                        d_809_v55_ = _dafny.Map({(((d_807_v53_)[(_dafny.MultiSet([d_800_v46_, ((self).f27)[default__.safeIndex(136, ((self).f27).length(0))]])).cardinality] if ((_dafny.MultiSet([d_800_v46_, ((self).f27)[default__.safeIndex(136, ((self).f27).length(0))]])).cardinality) in (d_807_v53_) else (self).f23)) - (((self).f27)[default__.safeIndex(136, ((self).f27).length(0))]): len((d_808_v54_)[default__.safeIndex(((self).f27)[default__.safeIndex(136, ((self).f27).length(0))], len(d_808_v54_))])})
                        d_809_v55_ = d_809_v55_
                        d_810_v56_: _dafny.Array
                        def lambda35_(d_811_v45_, d_812_v0_):
                            def lambda36_(d_813_i5_):
                                return _dafny.Map({d_811_v45_: d_812_v0_})

                            return lambda36_

                        init20_ = lambda35_(d_799_v45_, d_739_v0_)
                        nw134_ = _dafny.Array(None, 25)
                        for i0_20_ in range(nw134_.length(0)):
                            nw134_[i0_20_] = init20_(i0_20_)
                        d_810_v56_ = nw134_
                        d_814_v57_: _dafny.Map
                        d_814_v57_ = _dafny.Map({d_799_v45_: False})
                        index172_ = default__.safeIndex(452, (d_810_v56_).length(0))
                        (d_810_v56_)[index172_] = d_814_v57_
                        d_815_v58_: _dafny.Seq
                        d_815_v58_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")) for d_816_i6_ in range(default__.abs(975))])])
                        d_817_v59_: _dafny.Seq
                        d_817_v59_ = _dafny.SeqWithoutIsStrInference([default__.fm21(((d_815_v58_)[default__.safeIndex(d_800_v46_, len(d_815_v58_))]).set(default__.safeIndex(d_800_v46_, len((d_815_v58_)[default__.safeIndex(d_800_v46_, len(d_815_v58_))])), p0), globalState)])
                        index173_ = default__.safeIndex(452, (d_810_v56_).length(0))
                        rhs180_ = _dafny.Map({d_799_v45_: d_739_v0_})
                        rhs181_ = False
                        rhs182_ = d_817_v59_
                        rhs183_ = (self).f23
                        lhs168_ = d_810_v56_
                        lhs169_ = default__.safeIndex(452, (d_810_v56_).length(0))
                        lhs170_ = globalState
                        lhs171_ = globalState
                        lhs168_[lhs169_] = rhs180_
                        lhs170_.f21 = rhs181_
                        d_817_v59_ = rhs182_
                        lhs171_.f8 = rhs183_
                        (globalState).f14 = default__.fm0((D0_DC1(d_800_v46_, not((d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]), d_800_v46_)).cf3, not((((self).f27)[default__.safeIndex(136, ((self).f27).length(0))]) <= (d_800_v46_)), globalState)
                    if True:
                        d_818_v60_: _dafny.Array
                        nw135_ = _dafny.Array(_dafny.Map({}), 13)
                        d_818_v60_ = nw135_
                        d_819_v61_: D11
                        d_819_v61_ = D11_DC33(D11_DC30(d_818_v60_))
                        d_820_v62_: D11
                        d_820_v62_ = D11_DC30(d_818_v60_)
                        d_821_v63_: _dafny.Array
                        nw136_ = _dafny.Array(None, 18)
                        nw136_[int(0)] = d_819_v61_
                        nw136_[int(1)] = d_819_v61_
                        nw136_[int(2)] = d_819_v61_
                        nw136_[int(3)] = d_819_v61_
                        nw136_[int(4)] = d_819_v61_
                        nw136_[int(5)] = d_819_v61_
                        nw136_[int(6)] = d_819_v61_
                        nw136_[int(7)] = d_819_v61_
                        nw136_[int(8)] = D11_DC33(d_820_v62_)
                        nw136_[int(9)] = d_819_v61_
                        nw136_[int(10)] = d_819_v61_
                        nw136_[int(11)] = d_819_v61_
                        nw136_[int(12)] = d_819_v61_
                        nw136_[int(13)] = d_819_v61_
                        nw136_[int(14)] = d_819_v61_
                        nw136_[int(15)] = d_819_v61_
                        nw136_[int(16)] = d_819_v61_
                        nw136_[int(17)] = d_819_v61_
                        d_821_v63_ = nw136_
                        d_822_v64_: _dafny.Map
                        d_822_v64_ = _dafny.Map({(0) - (default__.safeModuloInt(len(d_790_v36_), (0) - ((d_790_v36_)[default__.safeIndex(303, len(d_790_v36_))]))): d_821_v63_})
                        d_822_v64_ = (d_822_v64_).set(885, d_821_v63_)
                        d_823_v65_: C1
                        nw137_ = C1()
                        nw137_.ctor__((d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], d_790_v36_, (self).f23)
                        d_823_v65_ = nw137_
                        d_824_v66_: _dafny.Seq
                        d_824_v66_ = _dafny.SeqWithoutIsStrInference([d_823_v65_])
                        d_825_v67_: _dafny.Map
                        d_825_v67_ = _dafny.Map({(d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]: default__.fm1(len(d_824_v66_), default__.fm1(len(d_756_v16_), (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], globalState), globalState)})
                        d_826_v68_: _dafny.Map
                        d_826_v68_ = _dafny.Map({len(d_825_v67_): d_756_v16_})
                        d_827_v69_: _dafny.Map
                        d_827_v69_ = _dafny.Map({((d_826_v68_).set((self).f23, d_756_v16_)) | (d_826_v68_): (self).f23})
                        (globalState).f14 = ((d_827_v69_)[(d_826_v68_ if (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))] else _dafny.Map({(self).f23: d_756_v16_}))] if ((d_826_v68_ if (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))] else _dafny.Map({(self).f23: d_756_v16_}))) in (d_827_v69_) else (self).f23)
                        d_828_v70_: _dafny.Array
                        nw138_ = _dafny.Array(None, 10)
                        d_828_v70_ = nw138_
                        d_829_v71_: _dafny.Map
                        d_829_v71_ = _dafny.Map({(self).f23: (self).f23})
                        d_830_v72_: _dafny.Map
                        d_830_v72_ = _dafny.Map({(d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eys")))})
                        d_831_v73_: _dafny.MultiSet
                        d_831_v73_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-445 for d_832_i7_ in range(default__.abs(-251))]))])
                        d_833_v74_: C2
                        nw139_ = C2()
                        nw139_.ctor__(d_829_v71_, d_739_v0_, (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], d_790_v36_, (0) - (((d_830_v72_)[(self).fm10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bms")), d_739_v0_, d_830_v72_, True, globalState)] if ((self).fm10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bms")), d_739_v0_, d_830_v72_, True, globalState)) in (d_830_v72_) else (d_831_v73_).cardinality)))
                        d_833_v74_ = nw139_
                        index174_ = default__.safeIndex(156, (d_828_v70_).length(0))
                        (d_828_v70_)[index174_] = d_833_v74_
                        d_834_v75_: str
                        d_834_v75_ = _dafny.CodePoint('o')
                        d_835_v76_: _dafny.Map
                        d_835_v76_ = _dafny.Map({d_833_v74_: len(_dafny.Map({not(d_739_v0_): d_834_v75_}))})
                        index175_ = default__.safeIndex(156, (d_828_v70_).length(0))
                        rhs184_ = not(d_739_v0_)
                        rhs185_ = (d_833_v74_).f32
                        rhs186_ = default__.safeDivisionInt((self).f23, ((d_835_v76_)[d_833_v74_] if (d_833_v74_) in (d_835_v76_) else (self).f23))
                        rhs187_ = d_833_v74_
                        lhs172_ = globalState
                        lhs173_ = globalState
                        lhs174_ = globalState
                        lhs175_ = d_828_v70_
                        lhs176_ = default__.safeIndex(156, (d_828_v70_).length(0))
                        lhs172_.f9 = rhs184_
                        lhs173_.f21 = rhs185_
                        lhs174_.f14 = rhs186_
                        lhs175_[lhs176_] = rhs187_
                        index176_ = default__.safeIndex(808, (d_793_v39_).length(0))
                        (d_793_v39_)[index176_] = d_739_v0_
                        d_836_v77_: _dafny.Map
                        d_836_v77_ = _dafny.Map({True: d_823_v65_})
                        d_836_v77_ = d_836_v77_
                    elif True:
                        d_837_v78_: _dafny.Map
                        d_837_v78_ = _dafny.Map({((d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))] if (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))] else (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]): (self).f23})
                        d_838_v79_: _dafny.Map
                        d_838_v79_ = _dafny.Map({default__.safeDivisionInt((self).f23, (self).f23): (self).f23})
                        rhs188_ = d_739_v0_
                        rhs189_ = ((d_838_v79_)[(self).f23] if ((self).f23) in (d_838_v79_) else (self).f23)
                        rhs190_ = len((d_837_v78_) | ((d_837_v78_).set(d_739_v0_, 437)))
                        rhs191_ = (d_837_v78_) | (d_837_v78_)
                        lhs177_ = globalState
                        lhs178_ = globalState
                        lhs177_.f7 = rhs188_
                        r0 = rhs189_
                        lhs178_.f8 = rhs190_
                        d_837_v78_ = rhs191_
                        d_839_v80_: _dafny.Map
                        d_839_v80_ = _dafny.Map({(d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]: p0})
                        d_840_v81_: D11
                        d_840_v81_ = D11_DC32(p0)
                        d_841_v82_: _dafny.Array
                        nw140_ = _dafny.Array(None, 11)
                        nw140_[int(0)] = d_839_v80_
                        nw140_[int(1)] = d_839_v80_
                        nw140_[int(2)] = d_839_v80_
                        nw140_[int(3)] = _dafny.Map({not(default__.fm1(len((_dafny.Map({(d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]: d_739_v0_})).set((d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))], d_739_v0_)), d_739_v0_, globalState)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqs"))})
                        nw140_[int(4)] = d_839_v80_
                        nw140_[int(5)] = d_839_v80_
                        nw140_[int(6)] = default__.fm27(globalState)
                        nw140_[int(7)] = d_839_v80_
                        nw140_[int(8)] = _dafny.Map({d_739_v0_: (d_840_v81_).cf53})
                        nw140_[int(9)] = (d_839_v80_) | (d_839_v80_)
                        nw140_[int(10)] = d_839_v80_
                        d_841_v82_ = nw140_
                        index177_ = default__.safeIndex(933, (d_841_v82_).length(0))
                        (d_841_v82_)[index177_] = d_839_v80_
                        d_842_v83_: _dafny.Map
                        d_842_v83_ = _dafny.Map({d_792_v38_: (d_839_v80_) | ((d_839_v80_).set(d_739_v0_, p0))})
                        index178_ = default__.safeIndex(933, (d_841_v82_).length(0))
                        (d_841_v82_)[index178_] = ((d_842_v83_)[(d_792_v38_) + (d_792_v38_)] if ((d_792_v38_) + (d_792_v38_)) in (d_842_v83_) else default__.fm27(globalState))
                        d_843_v84_: D3
                        d_843_v84_ = D3_DC12(d_792_v38_)
                        d_844_v85_: _dafny.Array
                        def lambda37_(d_845_p0_):
                            def lambda38_(d_846_i8_):
                                return _dafny.SeqWithoutIsStrInference([d_845_p0_, d_845_p0_, d_845_p0_])

                            return lambda38_

                        init21_ = lambda37_(p0)
                        nw141_ = _dafny.Array(None, 6)
                        for i0_21_ in range(nw141_.length(0)):
                            nw141_[i0_21_] = init21_(i0_21_)
                        d_844_v85_ = nw141_
                        d_847_v86_: _dafny.Seq
                        d_847_v86_ = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])
                        d_848_v87_: _dafny.Seq
                        d_848_v87_ = d_847_v86_
                        index179_ = default__.safeIndex(292, (d_844_v85_).length(0))
                        (d_844_v85_)[index179_] = d_848_v87_
                        index180_ = default__.safeIndex(292, (d_844_v85_).length(0))
                        rhs192_ = d_843_v84_
                        rhs193_ = d_848_v87_
                        rhs194_ = (self).f23
                        lhs179_ = d_844_v85_
                        lhs180_ = default__.safeIndex(292, (d_844_v85_).length(0))
                        lhs181_ = globalState
                        d_843_v84_ = rhs192_
                        lhs179_[lhs180_] = rhs193_
                        lhs181_.f8 = rhs194_
                        (globalState).f14 = (self).f23
                        d_849_v88_: _dafny.Map
                        d_849_v88_ = _dafny.Map({(self).f23: (d_793_v39_)[default__.safeIndex(808, (d_793_v39_).length(0))]})
                        d_838_v79_ = (d_838_v79_).set(len(d_849_v88_), 933)
                    pass
            pass
        d_850_v89_: D13
        d_850_v89_ = D13_DC38(d_756_v16_)
        d_851_v90_: D13
        d_851_v90_ = D13_DC40(d_850_v89_)
        pat_let_tv35_ = d_850_v89_
        pat_let_tv36_ = d_850_v89_
        def iife55_(_pat_let13_0):
            def iife56_(d_852_dt__update__tmp_h0_):
                def iife57_(_pat_let14_0):
                    def iife58_(d_853_dt__update_hcf64_h0_):
                        return D13_DC40(d_853_dt__update_hcf64_h0_)
                    return iife58_(_pat_let14_0)
                return iife57_(D13_DC40(pat_let_tv35_))
            return iife56_(_pat_let13_0)
        def iife54_(_pat_let12_0):
            def iife59_(d_854_dt__update__tmp_h1_):
                def iife60_(_pat_let15_0):
                    def iife61_(d_855_dt__update_hcf64_h1_):
                        return D13_DC40(d_855_dt__update_hcf64_h1_)
                    return iife61_(_pat_let15_0)
                return iife60_(pat_let_tv36_)
            return iife59_(_pat_let12_0)
        source7_ = iife54_(iife55_(d_851_v90_))
        if source7_.is_DC39:
            d_856___mcc_h5_ = source7_.cf63
            d_857_cf63_ = d_856___mcc_h5_
            d_858_v91_: _dafny.Array
            nw142_ = _dafny.Array(False, 22)
            d_858_v91_ = nw142_
            d_859_v92_: _dafny.Set
            d_859_v92_ = _dafny.Set({d_858_v91_})
            d_860_v93_: _dafny.Seq
            d_860_v93_ = _dafny.SeqWithoutIsStrInference([d_858_v91_])
            d_861_v94_: _dafny.Seq
            d_861_v94_ = _dafny.SeqWithoutIsStrInference([d_859_v92_, _dafny.Set({d_858_v91_})])
            d_862_v95_: _dafny.Array
            nw143_ = _dafny.Array(None, 19)
            nw143_[int(0)] = (d_859_v92_ if d_739_v0_ else d_859_v92_)
            nw143_[int(1)] = d_859_v92_
            nw143_[int(2)] = d_859_v92_
            nw143_[int(3)] = _dafny.Set({d_858_v91_, (d_860_v93_)[default__.safeIndex(d_857_cf63_, len(d_860_v93_))]})
            nw143_[int(4)] = d_859_v92_
            nw143_[int(5)] = d_859_v92_
            nw143_[int(6)] = (d_859_v92_) | (d_859_v92_)
            nw143_[int(7)] = d_859_v92_
            nw143_[int(8)] = d_859_v92_
            nw143_[int(9)] = d_859_v92_
            nw143_[int(10)] = d_859_v92_
            nw143_[int(11)] = (d_861_v94_)[default__.safeIndex(d_857_cf63_, len(d_861_v94_))]
            nw143_[int(12)] = d_859_v92_
            nw143_[int(13)] = d_859_v92_
            nw143_[int(14)] = d_859_v92_
            nw143_[int(15)] = d_859_v92_
            nw143_[int(16)] = d_859_v92_
            nw143_[int(17)] = d_859_v92_
            nw143_[int(18)] = d_859_v92_
            d_862_v95_ = nw143_
            index181_ = default__.safeIndex(446, (d_862_v95_).length(0))
            (d_862_v95_)[index181_] = d_859_v92_
            d_863_v96_: _dafny.Array
            nw144_ = _dafny.Array(_dafny.CodePoint('D'), 23)
            d_863_v96_ = nw144_
            d_864_v97_: _dafny.Seq
            d_864_v97_ = _dafny.SeqWithoutIsStrInference([D9_DC24(d_863_v96_)])
            d_865_v98_: _dafny.Seq
            d_865_v98_ = _dafny.SeqWithoutIsStrInference([len(d_864_v97_)])
            d_866_v99_: D9
            d_866_v99_ = D9_DC27(d_739_v0_, d_865_v98_, d_858_v91_, False, True)
            index182_ = default__.safeIndex(446, (d_862_v95_).length(0))
            (d_862_v95_)[index182_] = _dafny.Set({(d_866_v99_).cf46, d_858_v91_})
            index183_ = default__.safeIndex(197, ((self).f27).length(0))
            ((self).f27)[index183_] = (self).f23
            d_867_v100_: C1
            nw145_ = C1()
            nw145_.ctor__(not(d_739_v0_), d_865_v98_, (self).f23)
            d_867_v100_ = nw145_
            d_868_v101_: _dafny.Map
            d_868_v101_ = _dafny.Map({d_739_v0_: d_867_v100_})
            d_869_v102_: _dafny.Seq
            d_869_v102_ = _dafny.SeqWithoutIsStrInference([d_739_v0_])
            index184_ = default__.safeIndex(197, ((self).f27).length(0))
            ((self).f27)[index184_] = default__.fm0(len(d_868_v101_), (d_869_v102_)[default__.safeIndex((self).f23, len(d_869_v102_))], globalState)
            d_870_v103_: _dafny.Map
            d_870_v103_ = _dafny.Map({default__.fm0(((self).f27)[default__.safeIndex(197, ((self).f27).length(0))], d_739_v0_, globalState): d_739_v0_})
            d_871_v104_: _dafny.MultiSet
            d_871_v104_ = _dafny.MultiSet([(self).f23, len(d_870_v103_), (self).f23])
            if (((d_871_v104_).cardinality) + (((self).f27)[default__.safeIndex(197, ((self).f27).length(0))])) < (((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]):
                d_869_v102_ = ((d_869_v102_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_869_v102_) for d_872_i9_ in range(default__.abs(878))])), len(d_869_v102_)), d_739_v0_)) + ((_dafny.SeqWithoutIsStrInference([d_739_v0_])) + (d_869_v102_))
                index185_ = default__.safeIndex(197, ((self).f27).length(0))
                ((self).f27)[index185_] = len(((p0) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chtuw"))) + (p0))).set(default__.safeIndex(default__.safeModuloInt(((self).f27)[default__.safeIndex(197, ((self).f27).length(0))], ((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]), len((p0) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chtuw"))) + (p0)))), (p0)[default__.safeIndex(((self).f27)[default__.safeIndex(197, ((self).f27).length(0))], len(p0))]))
                d_873_v105_: D3
                d_873_v105_ = D3_DC11(((d_871_v104_)[d_857_cf63_] if (d_857_cf63_) in (d_871_v104_) else ((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]))
                d_873_v105_ = d_873_v105_
                d_874_v106_: _dafny.Map
                d_874_v106_ = _dafny.Map({((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]: (self).f23})
                def iife62_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(311, 520):
                        d_875_v107_: int = compr_30_
                        if ((311) <= (d_875_v107_)) and ((d_875_v107_) < (520)):
                            coll30_[(d_875_v107_) + ((self).f23)] = d_857_cf63_
                    return _dafny.Map(coll30_)
                d_874_v106_ = (iife62_()
                ) | (d_874_v106_)
                d_876_v108_: C0
                nw146_ = C0()
                nw146_.ctor__()
                d_876_v108_ = nw146_
            elif True:
                d_877_v109_: str
                d_877_v109_ = _dafny.CodePoint('n')
                d_878_v110_: _dafny.MultiSet
                d_878_v110_ = _dafny.MultiSet([d_877_v109_])
                d_879_v112_: D9
                def iife63_():
                    coll31_ = _dafny.Set()
                    compr_31_: str
                    for compr_31_ in (d_878_v110_).Elements:
                        d_880_v111_: str = compr_31_
                        if (d_880_v111_) in (d_878_v110_):
                            coll31_ = coll31_.union(_dafny.Set([d_880_v111_]))
                    return _dafny.Set(coll31_)
                d_879_v112_ = D9_DC25(d_739_v0_, iife63_()
)
                d_881_v113_: D9
                d_881_v113_ = D9_DC28(d_879_v112_)
                d_881_v113_ = d_881_v113_
                d_882_v114_: _dafny.Map
                d_882_v114_ = _dafny.Map({d_739_v0_: ((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]})
                d_883_v115_: D9
                d_883_v115_ = D9_DC26((self).f23)
                rhs195_ = p0
                rhs196_ = ((((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]) + (len(d_882_v114_))) >= (default__.safeDivisionInt((self).f23, d_857_cf63_))
                rhs197_ = d_739_v0_
                rhs198_ = d_739_v0_
                rhs199_ = default__.safeModuloInt(-409, (d_883_v115_).cf43)
                lhs182_ = globalState
                lhs183_ = globalState
                lhs184_ = globalState
                lhs185_ = globalState
                lhs186_ = globalState
                lhs182_.f6 = rhs195_
                lhs183_.f9 = rhs196_
                lhs184_.f9 = rhs197_
                lhs185_.f9 = rhs198_
                lhs186_.f8 = rhs199_
                d_739_v0_ = not ((self).fm10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrkt")), d_739_v0_, _dafny.Map({d_739_v0_: ((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]}), d_739_v0_, globalState)) or (d_739_v0_)
                index186_ = default__.safeIndex(197, ((self).f27).length(0))
                ((self).f27)[index186_] = ((self).f27)[default__.safeIndex(197, ((self).f27).length(0))]
                d_884_v116_: _dafny.MultiSet
                d_884_v116_ = _dafny.MultiSet([d_739_v0_])
                d_885_v117_: _dafny.Seq
                d_885_v117_ = _dafny.SeqWithoutIsStrInference([d_884_v116_, d_884_v116_])
                d_886_v118_: C3
                nw147_ = C3()
                nw147_.ctor__((d_885_v117_) + (d_885_v117_), d_857_cf63_, default__.fm1(d_857_cf63_, d_739_v0_, globalState), d_865_v98_)
                d_886_v118_ = nw147_
            d_887_v119_: D9
            d_887_v119_ = D9_DC26((self).f23)
            d_888_v120_: D9
            d_888_v120_ = D9_DC28(d_887_v119_)
            d_889_v121_: _dafny.Seq
            d_889_v121_ = _dafny.SeqWithoutIsStrInference([(d_888_v120_ if d_739_v0_ else d_888_v120_)])
            d_890_v122_: D14
            d_890_v122_ = D14_DC41(d_889_v121_)
            d_889_v121_ = (d_889_v121_) + ((d_890_v122_).cf65)
        elif source7_.is_DC38:
            d_891___mcc_h6_ = source7_.cf62
            d_892_cf62_ = d_891___mcc_h6_
            d_893_v123_: _dafny.Array
            def lambda39_(d_894_v0_):
                def lambda40_(d_895_i10_):
                    return d_894_v0_

                return lambda40_

            init22_ = lambda39_(d_739_v0_)
            nw148_ = _dafny.Array(None, 21)
            for i0_22_ in range(nw148_.length(0)):
                nw148_[i0_22_] = init22_(i0_22_)
            d_893_v123_ = nw148_
            d_896_v124_: D0
            d_897_v125_: _dafny.Map
            d_898_v126_: int
            out37_: D0
            out38_: _dafny.Map
            out39_: int
            out37_, out38_, out39_ = default__.m0(d_893_v123_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_899_i11_ in range(default__.abs(779))]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwigmupbn"))), globalState)
            d_896_v124_ = out37_
            d_897_v125_ = out38_
            d_898_v126_ = out39_
            if d_739_v0_:
                d_900_v127_: str
                d_900_v127_ = _dafny.CodePoint('v')
                (globalState).f11 = d_900_v127_
                d_901_v128_: _dafny.Seq
                d_901_v128_ = _dafny.SeqWithoutIsStrInference([p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fq"))])
                d_902_v129_: _dafny.Seq
                d_902_v129_ = d_901_v128_
                d_901_v128_ = (d_901_v128_) + (((d_902_v129_)) + (d_901_v128_))
                d_903_v130_: _dafny.Map
                d_903_v130_ = _dafny.Map({(d_893_v123_ if d_739_v0_ else d_893_v123_): (d_739_v0_ if d_739_v0_ else d_739_v0_)})
                d_903_v130_ = (d_903_v130_).set(d_893_v123_, (p0) != (p0))
                d_904_v131_: T1
                nw149_ = C1()
                nw149_.ctor__(d_739_v0_, _dafny.SeqWithoutIsStrInference([-205 for d_905_i12_ in range(default__.abs(771))]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfnpyft"))))
                d_904_v131_ = nw149_
                d_904_v131_ = d_904_v131_
                index187_ = default__.safeIndex(913, ((self).f27).length(0))
                ((self).f27)[index187_] = (self).f23
                index188_ = default__.safeIndex(913, ((self).f27).length(0))
                ((self).f27)[index188_] = (self).f23
            elif True:
                d_906_v132_: _dafny.Map
                d_906_v132_ = _dafny.Map({d_898_v126_: default__.fm6(962, globalState)})
                d_907_v133_: str
                d_907_v133_ = _dafny.CodePoint('j')
                d_908_v134_: _dafny.Map
                d_908_v134_ = _dafny.Map({d_907_v133_: not(d_739_v0_)})
                d_909_v135_: _dafny.Seq
                d_909_v135_ = _dafny.SeqWithoutIsStrInference([(self).f23])
                d_910_v136_: _dafny.Seq
                d_910_v136_ = _dafny.SeqWithoutIsStrInference([len(d_908_v134_), (self).f23, (d_909_v135_)[default__.safeIndex(d_898_v126_, len(d_909_v135_))]])
                rhs200_ = ((d_906_v132_)[(self).f23] if ((self).f23) in (d_906_v132_) else d_907_v133_)
                rhs201_ = (d_909_v135_) + (_dafny.SeqWithoutIsStrInference([(self).f23 for d_911_i13_ in range(default__.abs(336))]))
                lhs187_ = globalState
                lhs187_.f11 = rhs200_
                r1 = rhs201_
                d_912_v137_: _dafny.Seq
                d_912_v137_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f23: 396})])
                d_913_v138_: _dafny.Map
                d_913_v138_ = _dafny.Map({d_912_v137_: _dafny.SeqWithoutIsStrInference([d_907_v133_ for d_914_i14_ in range(default__.abs(-831))])})
                d_913_v138_ = (d_913_v138_).set((d_912_v137_) + (d_912_v137_), (p0) + (_dafny.SeqWithoutIsStrInference([d_907_v133_ for d_915_i15_ in range(default__.abs(271))])))
                index189_ = default__.safeIndex(403, (d_893_v123_).length(0))
                (d_893_v123_)[index189_] = not(((0) - (d_898_v126_)) <= (495))
                index190_ = default__.safeIndex(403, (d_893_v123_).length(0))
                (d_893_v123_)[index190_] = not(d_739_v0_)
                d_916_v139_: _dafny.MultiSet
                d_916_v139_ = _dafny.MultiSet([_dafny.CodePoint('w'), d_907_v133_, d_907_v133_])
                def iife64_():
                    coll32_ = _dafny.Set()
                    compr_32_: str
                    for compr_32_ in (d_908_v134_).keys.Elements:
                        d_917_v140_: str = compr_32_
                        if (d_917_v140_) in (d_908_v134_):
                            coll32_ = coll32_.union(_dafny.Set([d_917_v140_]))
                    return _dafny.Set(coll32_)
                rhs202_ = p0
                rhs203_ = default__.fm1((d_909_v135_)[default__.safeIndex(((d_916_v139_)[d_907_v133_] if (d_907_v133_) in (d_916_v139_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ew")))), len(d_909_v135_))], default__.fm1(len(iife64_()
                ), True, globalState), globalState)
                rhs204_ = d_898_v126_
                lhs188_ = globalState
                lhs189_ = globalState
                lhs190_ = globalState
                lhs188_.f6 = rhs202_
                lhs189_.f7 = rhs203_
                lhs190_.f8 = rhs204_
                d_918_v141_: _dafny.Seq
                d_918_v141_ = _dafny.SeqWithoutIsStrInference([(d_898_v126_) == (-100)])
                index191_ = default__.safeIndex(403, (d_893_v123_).length(0))
                (d_893_v123_)[index191_] = (d_918_v141_)[default__.safeIndex(len(p0), len(d_918_v141_))]
            if d_739_v0_:
                d_919_v142_: _dafny.Seq
                d_919_v142_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufjwpiq")))])
                (globalState).f11 = default__.fm6((0) - ((d_919_v142_)[default__.safeIndex((self).f23, len(d_919_v142_))]), globalState)
                d_920_v143_: C1
                nw150_ = C1()
                nw150_.ctor__(not (default__.fm1((self).f23, not(d_739_v0_), globalState)) or (d_739_v0_), d_919_v142_, (self).f23)
                d_920_v143_ = nw150_
                (globalState).f8 = d_898_v126_
                d_921_v144_: _dafny.MultiSet
                d_921_v144_ = _dafny.MultiSet([not(d_739_v0_), not (d_739_v0_) or (d_739_v0_), d_739_v0_, d_739_v0_])
                d_921_v144_ = d_921_v144_
                (globalState).f7 = not(d_739_v0_)
            elif True:
                d_922_v145_: _dafny.Seq
                d_922_v145_ = _dafny.SeqWithoutIsStrInference([d_898_v126_])
                (globalState).f8 = ((d_922_v145_)[default__.safeIndex(d_898_v126_, len(d_922_v145_))]) - ((self).f23)
                index192_ = default__.safeIndex(193, (d_893_v123_).length(0))
                (d_893_v123_)[index192_] = d_739_v0_
                d_923_v146_: _dafny.Seq
                d_923_v146_ = _dafny.SeqWithoutIsStrInference([d_739_v0_, d_739_v0_])
                d_924_v147_: _dafny.Map
                d_924_v147_ = _dafny.Map({d_923_v146_: 267})
                index193_ = default__.safeIndex(193, (d_893_v123_).length(0))
                (d_893_v123_)[index193_] = ((default__.fm28(d_898_v126_, globalState) if True else d_924_v147_)) != (d_924_v147_)
                d_925_v148_: str
                d_925_v148_ = _dafny.CodePoint('e')
                d_926_v149_: D4
                d_926_v149_ = D4_DC15((self).f23, (p0)[default__.safeIndex((self).f23, len(p0))], (self).f23)
                d_927_v150_: _dafny.Set
                d_927_v150_ = _dafny.Set({d_925_v148_})
                d_928_v151_: _dafny.Map
                d_928_v151_ = _dafny.Map({D2_DC7(d_927_v150_): d_925_v148_})
                d_929_v152_: D2
                d_929_v152_ = D2_DC7(d_927_v150_)
                d_930_v153_: D2
                d_930_v153_ = D2_DC8(d_898_v126_, len(d_897_v125_), d_925_v148_, (0) - ((self).f23), ((d_897_v125_)[d_898_v126_] if (d_898_v126_) in (d_897_v125_) else (d_893_v123_)[default__.safeIndex(193, (d_893_v123_).length(0))]))
                d_931_v154_: _dafny.Array
                nw151_ = _dafny.Array(None, 19)
                nw151_[int(0)] = d_925_v148_
                nw151_[int(1)] = d_925_v148_
                nw151_[int(2)] = d_925_v148_
                nw151_[int(3)] = default__.fm6(531, globalState)
                nw151_[int(4)] = d_925_v148_
                nw151_[int(5)] = (d_926_v149_).cf27
                nw151_[int(6)] = d_925_v148_
                nw151_[int(7)] = _dafny.CodePoint('y')
                nw151_[int(8)] = d_925_v148_
                nw151_[int(9)] = d_925_v148_
                nw151_[int(10)] = d_925_v148_
                nw151_[int(11)] = d_925_v148_
                nw151_[int(12)] = default__.fm6(len(d_923_v146_), globalState)
                nw151_[int(13)] = ((d_928_v151_)[d_929_v152_] if (d_929_v152_) in (d_928_v151_) else d_925_v148_)
                nw151_[int(14)] = d_925_v148_
                nw151_[int(15)] = _dafny.CodePoint('u')
                nw151_[int(16)] = d_925_v148_
                nw151_[int(17)] = (d_930_v153_).cf17
                nw151_[int(18)] = _dafny.CodePoint('y')
                d_931_v154_ = nw151_
                d_931_v154_ = d_931_v154_
                d_932_v155_: _dafny.Map
                d_932_v155_ = _dafny.Map({((d_893_v123_)[default__.safeIndex(193, (d_893_v123_).length(0))]) == (d_739_v0_): p0})
                d_932_v155_ = (d_932_v155_).set((d_893_v123_)[default__.safeIndex(193, (d_893_v123_).length(0))], p0)
                d_933_v157_: C3
                nw152_ = C3()
                def iife65_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(899, 229):
                        d_934_v156_: int = compr_33_
                        if ((899) <= (d_934_v156_)) and ((d_934_v156_) < (229)):
                            coll33_ = coll33_.union(_dafny.Set([(d_934_v156_) - ((self).f23)]))
                    return _dafny.Set(coll33_)
                nw152_.ctor__(default__.fm29(d_898_v126_, iife65_()
                , globalState), (self).f23, (d_893_v123_)[default__.safeIndex(193, (d_893_v123_).length(0))], (d_922_v145_) + (d_922_v145_))
                d_933_v157_ = nw152_
                d_933_v157_ = d_933_v157_
            d_739_v0_ = not (not(False)) or (d_739_v0_)
        elif True:
            d_935___mcc_h7_ = source7_.cf64
            d_936_cf64_ = d_935___mcc_h7_
            d_937_v158_: _dafny.Array
            nw153_ = _dafny.Array(D1.default()(), 25)
            d_937_v158_ = nw153_
            d_938_v159_: _dafny.MultiSet
            d_938_v159_ = _dafny.MultiSet([d_937_v158_])
            d_939_v160_: _dafny.Seq
            d_939_v160_ = _dafny.SeqWithoutIsStrInference([D3_DC9(d_938_v159_)])
            d_940_v161_: _dafny.MultiSet
            d_940_v161_ = _dafny.MultiSet([393])
            d_941_v162_: _dafny.Map
            d_941_v162_ = _dafny.Map({_dafny.MultiSet([(self).f23]): default__.fm1((self).f23, d_739_v0_, globalState)})
            d_942_v163_: _dafny.Seq
            d_942_v163_ = _dafny.SeqWithoutIsStrInference([not(d_739_v0_)])
            d_943_v164_: _dafny.Set
            d_943_v164_ = _dafny.Set({(0) - ((self).f23)})
            d_944_v165_: str
            d_944_v165_ = _dafny.CodePoint('p')
            d_945_v166_: _dafny.Map
            d_945_v166_ = _dafny.Map({d_944_v165_: (self).f23})
            d_946_v167_: _dafny.MultiSet
            d_946_v167_ = _dafny.MultiSet([d_739_v0_, False])
            d_947_v168_: _dafny.Map
            d_947_v168_ = _dafny.Map({d_739_v0_: d_739_v0_})
            d_948_v169_: _dafny.Map
            d_948_v169_ = _dafny.Map({d_739_v0_: len(d_947_v168_)})
            d_949_v170_: _dafny.Array
            nw154_ = _dafny.Array(None, 27)
            nw154_[int(0)] = d_739_v0_
            nw154_[int(1)] = (d_939_v160_) < (d_939_v160_)
            nw154_[int(2)] = (d_940_v161_).issubset((d_940_v161_).set((self).f23, default__.abs((self).f23)))
            nw154_[int(3)] = d_739_v0_
            nw154_[int(4)] = ((d_941_v162_)[_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_942_v163_)]))] if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_942_v163_)]))) in (d_941_v162_) else d_739_v0_)
            nw154_[int(5)] = d_739_v0_
            nw154_[int(6)] = (d_943_v164_).ispropersubset(d_943_v164_)
            nw154_[int(7)] = (-888) != ((0) - ((self).f23))
            nw154_[int(8)] = True
            nw154_[int(9)] = ((0) - (((d_945_v166_)[d_944_v165_] if (d_944_v165_) in (d_945_v166_) else (d_940_v161_).cardinality))) < ((self).f23)
            nw154_[int(10)] = ((self).f23) > ((self).f23)
            nw154_[int(11)] = d_739_v0_
            nw154_[int(12)] = not((d_739_v0_) == (d_739_v0_))
            nw154_[int(13)] = ((self).f23) <= ((self).f23)
            nw154_[int(14)] = d_739_v0_
            nw154_[int(15)] = not(d_739_v0_)
            nw154_[int(16)] = d_739_v0_
            nw154_[int(17)] = d_739_v0_
            nw154_[int(18)] = d_739_v0_
            nw154_[int(19)] = not(d_739_v0_)
            nw154_[int(20)] = ((d_946_v167_).cardinality) != ((self).f23)
            nw154_[int(21)] = (self).fm10(p0, d_739_v0_, d_948_v169_, d_739_v0_, globalState)
            nw154_[int(22)] = (self).fm10(p0, d_739_v0_, d_948_v169_, d_739_v0_, globalState)
            nw154_[int(23)] = d_739_v0_
            nw154_[int(24)] = d_739_v0_
            nw154_[int(25)] = d_739_v0_
            nw154_[int(26)] = d_739_v0_
            d_949_v170_ = nw154_
            d_949_v170_ = d_949_v170_
            d_950_v171_: _dafny.Array
            def lambda41_(d_951_i16_):
                return (d_951_i16_) * ((self).f23)

            init23_ = lambda41_
            nw155_ = _dafny.Array(None, 6)
            for i0_23_ in range(nw155_.length(0)):
                nw155_[i0_23_] = init23_(i0_23_)
            d_950_v171_ = nw155_
            d_952_v172_: _dafny.Map
            d_952_v172_ = _dafny.Map({d_739_v0_: (self).f23})
            rhs205_ = default__.fm0((0) - ((self).f23), True, globalState)
            rhs206_ = d_950_v171_
            rhs207_ = d_948_v169_
            rhs208_ = (self).f23
            r0 = rhs205_
            d_950_v171_ = rhs206_
            d_952_v172_ = rhs207_
            r0 = rhs208_
            d_953_v173_: _dafny.Array
            nw156_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_953_v173_ = nw156_
            d_953_v173_ = d_953_v173_
            index194_ = default__.safeIndex(618, (d_949_v170_).length(0))
            (d_949_v170_)[index194_] = (True if not(d_739_v0_) else d_739_v0_)
            index195_ = default__.safeIndex(618, (d_949_v170_).length(0))
            (d_949_v170_)[index195_] = d_739_v0_
        d_954_v174_: D8
        d_954_v174_ = D8_DC22(d_739_v0_)
        pat_let_tv37_ = d_739_v0_
        def iife66_(_pat_let16_0):
            def iife67_(d_955_dt__update__tmp_h3_):
                def iife68_(_pat_let17_0):
                    def iife69_(d_956_dt__update_hcf38_h0_):
                        return D8_DC22(d_956_dt__update_hcf38_h0_)
                    return iife69_(_pat_let17_0)
                return iife68_(pat_let_tv37_)
            return iife67_(_pat_let16_0)
        (globalState).f7 = (iife66_(d_954_v174_)).cf38
        r0 = (self).f23
        r1 = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f23), (self).f23])), ((self).f23) * ((self).f23), (self).f23])
        return r0, r1

    @property
    def f27(self):
        return self._f27

class C5(T0):
    def  __init__(self):
        self._f23: int = int(0)
        self._f25: int = int(0)
        self._f26: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f23(self):
        return self._f23
    def ctor__(self, f25, f26, f23):
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f23 = f23

    def fm8(self, p0, p1, p2, p3, globalState):
        def iife70_():
            def iife75_():
                def iife76_():
                    def iife78_():
                        coll42_ = _dafny.Set()
                        compr_42_: str
                        for compr_42_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])).Elements:
                            d_964_v3_: str = compr_42_
                            if (d_964_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])):
                                coll42_ = coll42_.union(_dafny.Set([d_964_v3_]))
                        return _dafny.Set(coll42_)
                    coll40_ = _dafny.Map()
                    def iife77_():
                        coll41_ = _dafny.Set()
                        compr_41_: str
                        for compr_41_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])).Elements:
                            d_963_v3_: str = compr_41_
                            if (d_963_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])):
                                coll41_ = coll41_.union(_dafny.Set([d_963_v3_]))
                        return _dafny.Set(coll41_)
                    compr_40_: str
                    for compr_40_ in ((D2_DC7(iife77_()
)).cf14).Elements:
                        d_960_v2_: str = compr_40_
                        if (d_960_v2_) in ((D2_DC7(iife78_()
)).cf14):
                            coll40_[d_960_v2_] = True
                    return _dafny.Map(coll40_)
                coll39_ = _dafny.Map()
                compr_39_: int
                for compr_39_ in (_dafny.MultiSet([(self).f26, (self).f26])).Elements:
                    d_957_v1_: int = compr_39_
                    if (d_957_v1_) in (_dafny.MultiSet([(self).f26, (self).f26])):
                        coll39_[(d_957_v1_) - (len(iife76_()
                        ))] = False
                return _dafny.Map(coll39_)
            coll34_ = _dafny.Map()
            def iife71_():
                def iife72_():
                    def iife74_():
                        coll38_ = _dafny.Set()
                        compr_38_: str
                        for compr_38_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])).Elements:
                            d_961_v3_: str = compr_38_
                            if (d_961_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])):
                                coll38_ = coll38_.union(_dafny.Set([d_961_v3_]))
                        return _dafny.Set(coll38_)
                    coll36_ = _dafny.Map()
                    def iife73_():
                        coll37_ = _dafny.Set()
                        compr_37_: str
                        for compr_37_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])).Elements:
                            d_959_v3_: str = compr_37_
                            if (d_959_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_958_i0_ in range(default__.abs(547))])):
                                coll37_ = coll37_.union(_dafny.Set([d_959_v3_]))
                        return _dafny.Set(coll37_)
                    compr_36_: str
                    for compr_36_ in ((D2_DC7(iife73_()
)).cf14).Elements:
                        d_960_v2_: str = compr_36_
                        if (d_960_v2_) in ((D2_DC7(iife74_()
)).cf14):
                            coll36_[d_960_v2_] = True
                    return _dafny.Map(coll36_)
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in (_dafny.MultiSet([(self).f26, (self).f26])).Elements:
                    d_957_v1_: int = compr_35_
                    if (d_957_v1_) in (_dafny.MultiSet([(self).f26, (self).f26])):
                        coll35_[(d_957_v1_) - (len(iife72_()
                        ))] = False
                return _dafny.Map(coll35_)
            compr_34_: int
            for compr_34_ in (iife71_()
            ).keys.Elements:
                d_962_v0_: int = compr_34_
                if (d_962_v0_) in (iife75_()
                ):
                    coll34_[default__.safeModuloInt(d_962_v0_, (self).f25)] = False
            return _dafny.Map(coll34_)
        return _dafny.Map({(iife70_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdal"))): False})): (True) and (True)})

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_965_v0_: _dafny.Map
        d_965_v0_ = _dafny.Map({p1: -317})
        rhs209_ = p1
        rhs210_ = len(default__.fm9(p1, (0) - ((self).f25), (_dafny.Map({p1: (self).f26})) | (d_965_v0_), globalState))
        lhs191_ = globalState
        lhs191_.f21 = rhs209_
        r0 = rhs210_
        d_966_v1_: _dafny.Seq
        d_966_v1_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        d_967_v2_: C1
        nw157_ = C1()
        nw157_.ctor__((default__.fm1((self).f26, p1, globalState)) or (p1), (d_966_v1_) + (_dafny.SeqWithoutIsStrInference([-818 for d_968_i0_ in range(default__.abs(12))])), (0) - ((self).f26))
        d_967_v2_ = nw157_
        d_969_v3_: str
        d_969_v3_ = _dafny.CodePoint('m')
        d_970_v4_: _dafny.Map
        d_970_v4_ = _dafny.Map({d_969_v3_: p0})
        d_971_v5_: D2
        d_971_v5_ = D2_DC8(340, -216, d_969_v3_, (self).f26, p1)
        d_970_v4_ = (d_970_v4_).set((d_971_v5_).cf17, p0)
        d_972_v6_: D1
        d_972_v6_ = D1_DC4(default__.fm16(p1, (self).f23, default__.fm0((self).f26, False, globalState), globalState), p1, p1)
        d_973_i1_: int
        d_973_i1_ = 0
        with _dafny.label("12"):
            while (d_972_v6_).cf7:
                with _dafny.c_label("12"):
                    if (d_973_i1_) >= (100):
                        raise _dafny.Break("12")
                    d_973_i1_ = (d_973_i1_) + (1)
                    (globalState).f21 = p1
                    d_974_v7_: _dafny.Array
                    nw158_ = _dafny.Array(int(0), 4)
                    d_974_v7_ = nw158_
                    d_975_v8_: _dafny.Map
                    d_975_v8_ = _dafny.Map({(self).f23: (self).f23})
                    index196_ = default__.safeIndex(507, (d_974_v7_).length(0))
                    (d_974_v7_)[index196_] = default__.safeDivisionInt((0) - ((self).f23), len(d_975_v8_))
                    index197_ = default__.safeIndex(507, (d_974_v7_).length(0))
                    (d_974_v7_)[index197_] = (0) - (((self).f23 if ((self).f26) > ((self).f25) else (self).f25))
                    index198_ = default__.safeIndex(281, (p0).length(0))
                    (p0)[index198_] = p1
                    index199_ = default__.safeIndex(281, (p0).length(0))
                    (p0)[index199_] = ((self).f26) <= ((0) - ((0) - ((0) - (default__.fm0((self).f25, p1, globalState)))))
                    d_976_v9_: _dafny.Seq
                    d_976_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyx"))
                    d_977_v10_: _dafny.Map
                    d_977_v10_ = _dafny.Map({762: (p0)[default__.safeIndex(281, (p0).length(0))]})
                    d_978_v11_: _dafny.Map
                    d_978_v11_ = _dafny.Map({d_976_v9_: d_977_v10_})
                    index200_ = default__.safeIndex(507, (d_974_v7_).length(0))
                    rhs211_ = ((d_978_v11_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxiuegtec"))).set(default__.safeIndex(-119, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxiuegtec")))), d_969_v3_), d_977_v10_)).set(default__.fm2(globalState), d_977_v10_)
                    rhs212_ = (0) - (len(d_976_v9_))
                    lhs192_ = d_974_v7_
                    lhs193_ = default__.safeIndex(507, (d_974_v7_).length(0))
                    d_978_v11_ = rhs211_
                    lhs192_[lhs193_] = rhs212_
                    pass
            pass
        d_965_v0_ = (d_965_v0_).set(True, (self).f26)
        hi11_ = (self).f25
        for d_979_i2_ in range(len(d_966_v1_), hi11_):
            d_980_v12_: _dafny.Seq
            d_980_v12_ = _dafny.SeqWithoutIsStrInference([((self).f23) > ((0) - (-840))])
            d_981_v13_: D8
            d_981_v13_ = D8_DC19(d_966_v1_)
            d_982_v14_: _dafny.Seq
            d_982_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxqooq"))
            rhs213_ = not((4) != (((0) - ((self).f23)) * ((self).f25)))
            rhs214_ = (default__.safeModuloInt((self).f26, (self).f23)) != (default__.safeDivisionInt((self).f26, (self).f26))
            rhs215_ = ((default__.fm30((self).f23, d_981_v13_, d_982_v14_, (self).f25, globalState)) + (d_980_v12_)) + ((d_980_v12_) + (d_980_v12_))
            rhs216_ = ((self).f23) > ((self).f26)
            lhs194_ = globalState
            lhs195_ = globalState
            lhs196_ = globalState
            lhs194_.f7 = rhs213_
            lhs195_.f21 = rhs214_
            d_980_v12_ = rhs215_
            lhs196_.f21 = rhs216_
            d_983_v15_: D3
            d_983_v15_ = D3_DC10(p1)
            d_984_v16_: _dafny.Map
            d_984_v16_ = _dafny.Map({D3_DC10(p1): default__.fm0((self).f25, p1, globalState)})
            if (d_983_v15_) in (d_984_v16_):
                d_985_v17_: _dafny.Map
                d_985_v17_ = _dafny.Map({False: (d_966_v1_).set(default__.safeIndex(d_979_i2_, len(d_966_v1_)), (self).f26)})
                d_986_v18_: D12
                d_986_v18_ = D12_DC36(False, p1, (self).f25, (self).f25, d_985_v17_)
                d_987_v19_: _dafny.Map
                d_987_v19_ = _dafny.Map({len(_dafny.Map({p1: d_982_v14_})): d_969_v3_})
                d_988_v20_: _dafny.Seq
                d_988_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f26: d_969_v3_}), d_987_v19_])
                d_989_v21_: _dafny.Set
                d_989_v21_ = _dafny.Set({(self).f26})
                rhs217_ = (self).f23
                rhs218_ = p1
                rhs219_ = d_982_v14_
                rhs220_ = default__.safeModuloInt(len((d_988_v20_)[default__.safeIndex((self).f26, len(d_988_v20_))]), len((d_989_v21_) | (d_989_v21_)))
                rhs221_ = default__.fm31((self).f25, globalState)
                lhs197_ = globalState
                lhs198_ = globalState
                lhs199_ = globalState
                lhs200_ = globalState
                lhs197_.f8 = rhs217_
                lhs198_.f7 = rhs218_
                lhs199_.f6 = rhs219_
                lhs200_.f8 = rhs220_
                d_986_v18_ = rhs221_
                d_990_v22_: _dafny.Map
                d_990_v22_ = _dafny.Map({344: p1})
                d_991_v23_: _dafny.Map
                d_991_v23_ = _dafny.Map({p1: p1})
                d_992_v24_: D0
                d_992_v24_ = D0_DC1((self).f26, p1, len(d_991_v23_))
                d_990_v22_ = (d_990_v22_).set((self).f25, (d_992_v24_).cf2)
                d_993_v25_: C0
                nw159_ = C0()
                nw159_.ctor__()
                d_993_v25_ = nw159_
                index201_ = default__.safeIndex(200, (p0).length(0))
                (p0)[index201_] = p1
                index202_ = default__.safeIndex(200, (p0).length(0))
                (p0)[index202_] = p1
                d_994_v26_: _dafny.Array
                nw160_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_994_v26_ = nw160_
                d_995_v27_: _dafny.Array
                nw161_ = _dafny.Array(False, 15)
                d_995_v27_ = nw161_
                index203_ = default__.safeIndex(402, (d_994_v26_).length(0))
                (d_994_v26_)[index203_] = d_995_v27_
                index204_ = default__.safeIndex(402, (d_994_v26_).length(0))
                (d_994_v26_)[index204_] = d_995_v27_
            elif True:
                d_996_v28_: _dafny.MultiSet
                d_996_v28_ = _dafny.MultiSet([len(d_982_v14_)])
                (globalState).f7 = (p1) == ((d_996_v28_).issubset(_dafny.MultiSet([(self).f26])))
                d_997_v29_: _dafny.Array
                def lambda42_(d_998_i2_):
                    def lambda43_(d_999_i3_):
                        return (d_999_i3_) * (d_998_i2_)

                    return lambda43_

                init24_ = lambda42_(d_979_i2_)
                nw162_ = _dafny.Array(None, 27)
                for i0_24_ in range(nw162_.length(0)):
                    nw162_[i0_24_] = init24_(i0_24_)
                d_997_v29_ = nw162_
                d_1000_v30_: _dafny.Map
                d_1000_v30_ = _dafny.Map({p1: d_997_v29_})
                d_1001_v31_: _dafny.Map
                d_1001_v31_ = _dafny.Map({default__.safeModuloInt((self).f26, d_979_i2_): d_1000_v30_})
                d_1001_v31_ = (d_1001_v31_).set((self).f25, d_1000_v30_)
                d_1002_v32_: _dafny.Map
                d_1002_v32_ = _dafny.Map({(0) - (default__.fm0((self).f25, default__.fm1((self).f23, p1, globalState), globalState)): (self).f26})
                d_1003_v33_: _dafny.Map
                d_1003_v33_ = _dafny.Map({d_966_v1_: p1})
                d_1004_v34_: C2
                nw163_ = C2()
                nw163_.ctor__(d_1002_v32_, (d_1003_v33_) == (d_1003_v33_), p1, (_dafny.SeqWithoutIsStrInference([(self).f25, (self).f26])) + (d_966_v1_), (self).f25)
                d_1004_v34_ = nw163_
                d_1005_v35_: _dafny.Array
                def lambda44_(d_1006_i4_):
                    return D12_DC37(D12_DC35())

                init25_ = lambda44_
                nw164_ = _dafny.Array(None, 18)
                for i0_25_ in range(nw164_.length(0)):
                    nw164_[i0_25_] = init25_(i0_25_)
                d_1005_v35_ = nw164_
                d_1007_v36_: D12
                d_1007_v36_ = D12_DC35()
                d_1008_v37_: D12
                d_1008_v37_ = D12_DC37(d_1007_v36_)
                index205_ = default__.safeIndex(391, (d_1005_v35_).length(0))
                (d_1005_v35_)[index205_] = d_1008_v37_
                index206_ = default__.safeIndex(391, (d_1005_v35_).length(0))
                (d_1005_v35_)[index206_] = d_1008_v37_
                index207_ = default__.safeIndex(528, (d_997_v29_).length(0))
                (d_997_v29_)[index207_] = (self).f25
                index208_ = default__.safeIndex(528, (d_997_v29_).length(0))
                (d_997_v29_)[index208_] = ((self).f23) - ((self).f26)
            r0 = ((0) - ((self).f26)) * ((self).f25)
            d_1009_v38_: _dafny.Array
            def lambda45_(d_1010_i5_):
                return (d_1010_i5_) - ((self).f25)

            init26_ = lambda45_
            nw165_ = _dafny.Array(None, 7)
            for i0_26_ in range(nw165_.length(0)):
                nw165_[i0_26_] = init26_(i0_26_)
            d_1009_v38_ = nw165_
            d_1009_v38_ = d_1009_v38_
        r0 = (self).f23
        d_1011_v39_: _dafny.Set
        d_1011_v39_ = _dafny.Set({(self).f23})
        r1 = ((d_1011_v39_).intersection(d_1011_v39_)).intersection(d_1011_v39_)
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_1012_v0_: bool
        d_1012_v0_ = False
        d_1013_v1_: _dafny.Set
        d_1013_v1_ = _dafny.Set({d_1012_v0_, d_1012_v0_})
        d_1014_v2_: _dafny.Map
        d_1014_v2_ = _dafny.Map({(self).f25: (self).f25})
        d_1015_v3_: D13
        d_1015_v3_ = D13_DC39(((d_1014_v2_)[(self).f23] if ((self).f23) in (d_1014_v2_) else (self).f23))
        d_1016_v4_: _dafny.Set
        d_1016_v4_ = _dafny.Set({len(d_1013_v1_)})
        d_1017_v5_: _dafny.Array
        nw166_ = _dafny.Array(None, 12)
        nw166_[int(0)] = (self).f23
        nw166_[int(1)] = len(p0)
        nw166_[int(2)] = (self).f23
        nw166_[int(3)] = (self).f23
        nw166_[int(4)] = 700
        nw166_[int(5)] = (self).f25
        nw166_[int(6)] = (d_1015_v3_).cf63
        nw166_[int(7)] = (0) - ((self).f25)
        nw166_[int(8)] = (self).f26
        nw166_[int(9)] = (self).f25
        nw166_[int(10)] = len(d_1016_v4_)
        nw166_[int(11)] = 947
        d_1017_v5_ = nw166_
        d_1018_v6_: _dafny.Map
        d_1018_v6_ = _dafny.Map({d_1013_v1_: d_1017_v5_})
        d_1019_v7_: D13
        d_1019_v7_ = D13_DC38(d_1013_v1_)
        d_1020_v8_: _dafny.Map
        d_1020_v8_ = _dafny.Map({default__.fm1((self).f23, not(d_1012_v0_), globalState): (d_1019_v7_).cf62})
        d_1021_v9_: D0
        d_1021_v9_ = D0_DC0(p0)
        rhs222_ = _dafny.Map({((d_1020_v8_)[False] if (False) in (d_1020_v8_) else d_1013_v1_): d_1017_v5_})
        rhs223_ = (d_1013_v1_) - (d_1013_v1_)
        rhs224_ = ((d_1021_v9_).cf0) + (p0)
        lhs201_ = globalState
        d_1018_v6_ = rhs222_
        d_1013_v1_ = rhs223_
        lhs201_.f6 = rhs224_
        d_1013_v1_ = d_1013_v1_
        d_1022_v10_: str
        d_1022_v10_ = _dafny.CodePoint('k')
        (globalState).f11 = d_1022_v10_
        r0 = default__.safeModuloInt(((self).f26) - ((self).f25), (self).f26)
        d_1023_v11_: _dafny.Array
        nw167_ = _dafny.Array(None, 3)
        nw167_[int(0)] = d_1022_v10_
        nw167_[int(1)] = d_1022_v10_
        nw167_[int(2)] = d_1022_v10_
        d_1023_v11_ = nw167_
        d_1024_v12_: D9
        d_1024_v12_ = D9_DC24(d_1023_v11_)
        d_1025_v13_: _dafny.MultiSet
        d_1025_v13_ = _dafny.MultiSet([d_1024_v12_])
        d_1026_v14_: _dafny.Seq
        d_1026_v14_ = _dafny.SeqWithoutIsStrInference([D9_DC24(d_1023_v11_), d_1024_v12_, d_1024_v12_])
        (globalState).f8 = (((d_1025_v13_) - (d_1025_v13_)) - (_dafny.MultiSet(d_1026_v14_))).cardinality
        d_1027_v15_: _dafny.Map
        d_1027_v15_ = _dafny.Map({_dafny.CodePoint('i'): d_1012_v0_})
        d_1027_v15_ = (d_1027_v15_).set(d_1022_v10_, d_1012_v0_)
        r0 = default__.safeModuloInt((0) - ((self).f25), len(p0))
        r1 = _dafny.SeqWithoutIsStrInference([(self).f25])
        return r0, r1

    def m3(self, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Map = _dafny.Map({})
        d_1028_v0_: _dafny.Map
        d_1028_v0_ = _dafny.Map({(self).f26: (self).f26})
        d_1029_v1_: bool
        d_1029_v1_ = False
        d_1030_v2_: D1
        d_1030_v2_ = D1_DC5(954, d_1029_v1_, (self).f23, (self).f26)
        d_1031_v3_: _dafny.Seq
        d_1031_v3_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_1032_v4_: _dafny.Map
        d_1032_v4_ = _dafny.Map({d_1029_v1_: d_1029_v1_})
        d_1033_v5_: T1
        nw168_ = C2()
        nw168_.ctor__(d_1028_v0_, d_1029_v1_, (d_1030_v2_).cf10, d_1031_v3_, len(d_1032_v4_))
        d_1033_v5_ = nw168_
        d_1034_v6_: _dafny.Seq
        d_1034_v6_ = _dafny.SeqWithoutIsStrInference([d_1033_v5_])
        d_1035_v7_: _dafny.Seq
        d_1035_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uckqsq"))
        d_1036_v8_: _dafny.Seq
        d_1036_v8_ = _dafny.SeqWithoutIsStrInference([d_1035_v7_])
        hi12_ = (len((d_1036_v8_)[default__.safeIndex((self).f23, len(d_1036_v8_))]) if False else (self).f26)
        for d_1037_i0_ in range(len(((d_1034_v6_) + ((d_1034_v6_).set(default__.safeIndex(-858, len(d_1034_v6_)), d_1033_v5_))).set(default__.safeIndex((d_1033_v5_).f23, len((d_1034_v6_) + ((d_1034_v6_).set(default__.safeIndex(-858, len(d_1034_v6_)), d_1033_v5_)))), d_1033_v5_)), hi12_):
            (globalState).f9 = ((self).f25) != ((self).f26)
            d_1038_v9_: _dafny.Array
            nw169_ = _dafny.Array(int(0), 10)
            d_1038_v9_ = nw169_
            def lambda46_(d_1039_i1_):
                return default__.safeDivisionInt(d_1039_i1_, (self).f23)

            init27_ = lambda46_
            nw170_ = _dafny.Array(None, 28)
            for i0_27_ in range(nw170_.length(0)):
                nw170_[i0_27_] = init27_(i0_27_)
            d_1038_v9_ = nw170_
            r0 = ((d_1033_v5_).f28 if True else (d_1033_v5_).f28)
            d_1038_v9_ = d_1038_v9_
        if d_1029_v1_:
            d_1040_v10_: _dafny.Map
            d_1040_v10_ = _dafny.Map({d_1035_v7_: (d_1033_v5_).f28})
            d_1040_v10_ = (d_1040_v10_).set(d_1035_v7_, d_1029_v1_)
            (d_1033_v5_).f29 = d_1031_v3_
            (globalState).f7 = False
            d_1041_v11_: _dafny.MultiSet
            d_1041_v11_ = _dafny.MultiSet([(d_1033_v5_).f28, (d_1033_v5_).f28])
            d_1042_v12_: _dafny.Seq
            d_1042_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_1033_v5_).f28])])
            d_1043_v13_: C3
            nw171_ = C3()
            nw171_.ctor__((_dafny.SeqWithoutIsStrInference([d_1041_v11_, d_1041_v11_])) + (d_1042_v12_), (self).f25, (d_1033_v5_).f28, d_1031_v3_)
            d_1043_v13_ = nw171_
            d_1044_v14_: str
            d_1044_v14_ = _dafny.CodePoint('a')
            d_1045_v15_: _dafny.Seq
            d_1045_v15_ = _dafny.SeqWithoutIsStrInference([not(d_1029_v1_)])
            d_1046_v16_: _dafny.Array
            nw172_ = _dafny.Array(None, 27)
            nw172_[int(0)] = (d_1033_v5_).f28
            nw172_[int(1)] = (d_1033_v5_).f28
            nw172_[int(2)] = (d_1033_v5_).f28
            nw172_[int(3)] = True
            nw172_[int(4)] = (d_1033_v5_).f28
            nw172_[int(5)] = (d_1033_v5_).f28
            nw172_[int(6)] = (d_1033_v5_).f28
            nw172_[int(7)] = not(d_1029_v1_)
            nw172_[int(8)] = True
            nw172_[int(9)] = (d_1033_v5_).f28
            nw172_[int(10)] = (d_1033_v5_).f28
            nw172_[int(11)] = d_1029_v1_
            nw172_[int(12)] = d_1029_v1_
            nw172_[int(13)] = (d_1045_v15_)[default__.safeIndex((self).f26, len(d_1045_v15_))]
            nw172_[int(14)] = d_1029_v1_
            nw172_[int(15)] = d_1029_v1_
            nw172_[int(16)] = default__.fm1(-246, d_1029_v1_, globalState)
            nw172_[int(17)] = d_1029_v1_
            nw172_[int(18)] = (d_1033_v5_).f28
            nw172_[int(19)] = (d_1033_v5_).f28
            nw172_[int(20)] = False
            nw172_[int(21)] = (d_1033_v5_).f28
            nw172_[int(22)] = d_1029_v1_
            nw172_[int(23)] = (d_1045_v15_)[default__.safeIndex(356, len(d_1045_v15_))]
            nw172_[int(24)] = True
            nw172_[int(25)] = (d_1033_v5_).f28
            nw172_[int(26)] = d_1029_v1_
            d_1046_v16_ = nw172_
            d_1047_v17_: D9
            d_1047_v17_ = D9_DC27((d_1033_v5_).f28, _dafny.SeqWithoutIsStrInference([(self).f23 for d_1048_i2_ in range(default__.abs(19))]), d_1046_v16_, (d_1033_v5_).f28, (d_1045_v15_)[default__.safeIndex((self).f23, len(d_1045_v15_))])
            d_1049_v18_: D9
            d_1049_v18_ = D9_DC28(d_1047_v17_)
            d_1050_v19_: _dafny.Map
            d_1050_v19_ = _dafny.Map({d_1044_v14_: d_1049_v18_})
            d_1050_v19_ = (d_1050_v19_).set(d_1044_v14_, d_1049_v18_)
        elif True:
            d_1051_v20_: C0
            nw173_ = C0()
            nw173_.ctor__()
            d_1051_v20_ = nw173_
            (globalState).f14 = (0) - (((d_1033_v5_).f23) + ((d_1033_v5_).f23))
            d_1052_v21_: D11
            d_1052_v21_ = D11_DC32(d_1035_v7_)
            source8_ = d_1052_v21_
            if source8_.is_DC31:
                d_1053___mcc_h0_ = source8_.cf52
                d_1054_cf52_ = d_1053___mcc_h0_
                d_1055_v22_: _dafny.Set
                d_1055_v22_ = _dafny.Set({(d_1033_v5_).f23, 335, len(d_1032_v4_), (self).f23, (d_1031_v3_)[default__.safeIndex(d_1054_cf52_, len(d_1031_v3_))]})
                d_1056_v23_: C3
                nw174_ = C3()
                nw174_.ctor__(default__.fm29(default__.fm0(383, (d_1033_v5_).f28, globalState), d_1055_v22_, globalState), (d_1033_v5_).f23, (d_1033_v5_).f28, (_dafny.SeqWithoutIsStrInference([(self).f23])) + (d_1033_v5_.f29))
                d_1056_v23_ = nw174_
                d_1032_v4_ = (d_1032_v4_).set(default__.fm1(len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mnssog")): 387})), (d_1033_v5_).f28, globalState), d_1029_v1_)
                d_1057_v24_: _dafny.Set
                d_1057_v24_ = _dafny.Set({(d_1033_v5_).f28, d_1029_v1_})
                d_1057_v24_ = d_1057_v24_
                d_1058_v25_: _dafny.Map
                d_1058_v25_ = _dafny.Map({default__.fm24((d_1033_v5_).f28, len(d_1035_v7_), globalState): (d_1033_v5_).f23})
                d_1059_v26_: D9
                d_1059_v26_ = D9_DC26((d_1033_v5_).f23)
                d_1058_v25_ = (d_1058_v25_).set(d_1059_v26_, (d_1033_v5_).f23)
            elif source8_.is_DC32:
                d_1060___mcc_h1_ = source8_.cf53
                d_1061_cf53_ = d_1060___mcc_h1_
                rhs225_ = (d_1033_v5_).f23
                rhs226_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1062_i3_ in range(default__.abs(695))])
                lhs202_ = globalState
                lhs203_ = globalState
                lhs202_.f14 = rhs225_
                lhs203_.f6 = rhs226_
                d_1063_v27_: _dafny.Seq
                d_1063_v27_ = _dafny.SeqWithoutIsStrInference([(d_1033_v5_).f28])
                d_1029_v1_ = ((d_1033_v5_).f28) not in (d_1063_v27_)
                d_1064_v28_: _dafny.Set
                d_1064_v28_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1065_i4_ in range(default__.abs(915))])})
                (globalState).f7 = (d_1064_v28_) != (_dafny.Set({d_1061_cf53_, d_1035_v7_}))
                d_1031_v3_ = d_1033_v5_.f29
            elif source8_.is_DC30:
                d_1066___mcc_h2_ = source8_.cf51
                d_1067_cf51_ = d_1066___mcc_h2_
                (globalState).f8 = (default__.safeModuloInt((self).f26, (0) - ((d_1033_v5_).f23))) * (len(d_1035_v7_))
                d_1068_v29_: _dafny.Set
                d_1068_v29_ = _dafny.Set({(d_1033_v5_).f23})
                d_1069_v30_: _dafny.Map
                d_1069_v30_ = _dafny.Map({d_1029_v1_: len(_dafny.Map({d_1051_v20_: (d_1033_v5_).f28}))})
                d_1070_v31_: _dafny.Map
                d_1070_v31_ = _dafny.Map({False: d_1069_v30_})
                (globalState).f21 = (((0) - (default__.fm0(len(d_1068_v29_), not((d_1033_v5_).f28), globalState))) != ((self).f23)) in (d_1070_v31_)
                (globalState).f7 = (d_1033_v5_).f28
                d_1071_v32_: _dafny.Map
                d_1071_v32_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))): False})
                d_1072_v33_: C2
                nw175_ = C2()
                nw175_.ctor__(d_1028_v0_, not((d_1033_v5_).f28), d_1029_v1_, d_1033_v5_.f29, len((d_1071_v32_).set((d_1033_v5_).f23, (d_1033_v5_).f28)))
                d_1072_v33_ = nw175_
                d_1073_v34_: _dafny.MultiSet
                d_1073_v34_ = _dafny.MultiSet([not(True)])
                d_1074_v35_: D9
                d_1074_v35_ = D9_DC26((d_1033_v5_).f23)
                d_1075_v36_: _dafny.Set
                d_1075_v36_ = _dafny.Set({False})
                d_1076_v37_: _dafny.Array
                nw176_ = _dafny.Array(None, 17)
                nw176_[int(0)] = D9_DC26((d_1033_v5_).f23)
                nw176_[int(1)] = D9_DC26(((d_1073_v34_)[(d_1033_v5_).f28] if ((d_1033_v5_).f28) in (d_1073_v34_) else 418))
                nw176_[int(2)] = d_1074_v35_
                nw176_[int(3)] = d_1074_v35_
                nw176_[int(4)] = default__.fm24((d_1033_v5_).f28, 159, globalState)
                nw176_[int(5)] = D9_DC26((self).f25)
                nw176_[int(6)] = d_1074_v35_
                nw176_[int(7)] = d_1074_v35_
                nw176_[int(8)] = d_1074_v35_
                nw176_[int(9)] = d_1074_v35_
                nw176_[int(10)] = d_1074_v35_
                nw176_[int(11)] = d_1074_v35_
                nw176_[int(12)] = D9_DC26(len(d_1075_v36_))
                nw176_[int(13)] = d_1074_v35_
                nw176_[int(14)] = d_1074_v35_
                nw176_[int(15)] = d_1074_v35_
                nw176_[int(16)] = d_1074_v35_
                d_1076_v37_ = nw176_
                d_1077_v38_: _dafny.Map
                d_1077_v38_ = _dafny.Map({d_1072_v33_: d_1076_v37_})
                d_1077_v38_ = (d_1077_v38_).set(d_1072_v33_, d_1076_v37_)
            elif True:
                d_1078___mcc_h3_ = source8_.cf54
                d_1079_cf54_ = d_1078___mcc_h3_
                d_1080_v39_: D9
                d_1080_v39_ = D9_DC25(d_1029_v1_, _dafny.Set({_dafny.CodePoint('v')}))
                d_1081_v40_: D2
                d_1081_v40_ = D2_DC7((d_1080_v39_).cf42)
                d_1082_v41_: _dafny.Map
                d_1082_v41_ = _dafny.Map({d_1081_v40_: d_1035_v7_})
                d_1083_v42_: _dafny.Set
                d_1083_v42_ = _dafny.Set({d_1051_v20_})
                d_1084_v43_: _dafny.Set
                d_1084_v43_ = _dafny.Set({d_1029_v1_, (d_1033_v5_).f28, (d_1033_v5_).f28})
                d_1085_v44_: _dafny.Seq
                d_1085_v44_ = _dafny.SeqWithoutIsStrInference([not ((d_1051_v20_).fm20(d_1082_v41_, (d_1033_v5_).f23, (self).f26, len(d_1083_v42_), globalState)) or ((d_1033_v5_).f28), (d_1033_v5_).f28, ((d_1033_v5_).f28 if (d_1033_v5_).f28 else (d_1033_v5_).f28), (len(d_1084_v43_)) >= ((d_1033_v5_).f23)])
                (globalState).f21 = (d_1085_v44_)[default__.safeIndex(((self).f25 if (d_1033_v5_).f28 else 414), len(d_1085_v44_))]
                d_1086_v45_: _dafny.Array
                nw177_ = _dafny.Array(_dafny.Seq({}), 20)
                d_1086_v45_ = nw177_
                d_1087_v46_: _dafny.Seq
                d_1087_v46_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1033_v5_).f23, (0) - ((d_1033_v5_).f23)])])
                index209_ = default__.safeIndex(694, (d_1086_v45_).length(0))
                (d_1086_v45_)[index209_] = d_1087_v46_
                index210_ = default__.safeIndex(694, (d_1086_v45_).length(0))
                (d_1086_v45_)[index210_] = (_dafny.SeqWithoutIsStrInference([d_1033_v5_.f29 for d_1088_i5_ in range(default__.abs(506))])).set(default__.safeIndex(((d_1033_v5_).f23 if d_1029_v1_ else len(d_1036_v8_)), len(_dafny.SeqWithoutIsStrInference([d_1033_v5_.f29 for d_1089_i5_ in range(default__.abs(506))]))), d_1033_v5_.f29)
                d_1090_v47_: _dafny.Array
                nw178_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_1090_v47_ = nw178_
                d_1091_v48_: D9
                d_1091_v48_ = D9_DC24(d_1090_v47_)
                d_1092_v49_: D9
                d_1092_v49_ = D9_DC28(d_1091_v48_)
                d_1092_v49_ = d_1092_v49_
                d_1093_v50_: _dafny.Array
                nw179_ = _dafny.Array(_dafny.Set({}), 11)
                d_1093_v50_ = nw179_
                d_1094_v51_: str
                d_1094_v51_ = _dafny.CodePoint('l')
                d_1095_v52_: _dafny.Map
                d_1095_v52_ = _dafny.Map({d_1094_v51_: (d_1033_v5_).f28})
                d_1096_v53_: _dafny.Set
                d_1096_v53_ = _dafny.Set({d_1095_v52_})
                d_1097_v54_: _dafny.Map
                d_1097_v54_ = _dafny.Map({d_1095_v52_: (d_1033_v5_).f23})
                index211_ = default__.safeIndex(119, (d_1093_v50_).length(0))
                def iife79_():
                    coll43_ = _dafny.Set()
                    compr_43_: _dafny.Map
                    for compr_43_ in (d_1097_v54_).keys.Elements:
                        d_1098_v55_: _dafny.Map = compr_43_
                        if (d_1098_v55_) in (d_1097_v54_):
                            coll43_ = coll43_.union(_dafny.Set([d_1098_v55_]))
                    return _dafny.Set(coll43_)
                (d_1093_v50_)[index211_] = (d_1096_v53_) | (iife79_()
                )
                index212_ = default__.safeIndex(119, (d_1093_v50_).length(0))
                rhs227_ = (d_1033_v5_).f23
                rhs228_ = (d_1096_v53_) - (d_1096_v53_)
                rhs229_ = (d_1035_v7_).set(default__.safeIndex(len(d_1032_v4_), len(d_1035_v7_)), d_1094_v51_)
                lhs204_ = globalState
                lhs205_ = d_1093_v50_
                lhs206_ = default__.safeIndex(119, (d_1093_v50_).length(0))
                lhs207_ = globalState
                lhs204_.f14 = rhs227_
                lhs205_[lhs206_] = rhs228_
                lhs207_.f6 = rhs229_
            d_1099_v56_: _dafny.Seq
            d_1099_v56_ = _dafny.SeqWithoutIsStrInference([d_1029_v1_, (d_1033_v5_).f28, (d_1033_v5_).f28])
            d_1100_v57_: C1
            nw180_ = C1()
            nw180_.ctor__(not(not((d_1099_v56_)[default__.safeIndex((d_1033_v5_).f23, len(d_1099_v56_))])), d_1033_v5_.f29, ((self).f23) + ((self).f23))
            d_1100_v57_ = nw180_
            d_1101_v58_: str
            d_1101_v58_ = _dafny.CodePoint('y')
            d_1102_v59_: _dafny.Array
            def lambda47_(d_1103_v5_):
                def lambda48_(d_1104_i6_):
                    return (d_1104_i6_) * ((d_1103_v5_).f23)

                return lambda48_

            init28_ = lambda47_(d_1033_v5_)
            nw181_ = _dafny.Array(None, 17)
            for i0_28_ in range(nw181_.length(0)):
                nw181_[i0_28_] = init28_(i0_28_)
            d_1102_v59_ = nw181_
            d_1105_v60_: _dafny.Seq
            d_1105_v60_ = _dafny.SeqWithoutIsStrInference([d_1102_v59_, d_1102_v59_])
            d_1106_v61_: D1
            d_1106_v61_ = D1_DC3(d_1105_v60_)
            d_1107_v62_: _dafny.Map
            d_1107_v62_ = _dafny.Map({d_1101_v58_: d_1106_v61_})
            d_1108_v63_: _dafny.Set
            d_1108_v63_ = _dafny.Set({(d_1033_v5_).f28, not((d_1033_v5_).f28), (d_1033_v5_).f28})
            d_1109_v64_: _dafny.Map
            d_1109_v64_ = _dafny.Map({d_1108_v63_: (d_1107_v62_) | (d_1107_v62_)})
            d_1107_v62_ = ((d_1109_v64_)[(d_1108_v63_) - (d_1108_v63_)] if ((d_1108_v63_) - (d_1108_v63_)) in (d_1109_v64_) else d_1107_v62_)
        (globalState).f7 = (d_1033_v5_).f28
        hi13_ = ((d_1033_v5_).f23) - ((d_1033_v5_).f23)
        for d_1110_i7_ in range((self).f23, hi13_):
            d_1111_v65_: _dafny.Array
            def lambda49_(d_1112_i8_):
                return default__.safeDivisionInt(d_1112_i8_, (self).f23)

            init29_ = lambda49_
            nw182_ = _dafny.Array(None, 18)
            for i0_29_ in range(nw182_.length(0)):
                nw182_[i0_29_] = init29_(i0_29_)
            d_1111_v65_ = nw182_
            d_1113_v66_: C4
            nw183_ = C4()
            nw183_.ctor__(d_1111_v65_, len(d_1035_v7_))
            d_1113_v66_ = nw183_
            d_1114_v67_: _dafny.MultiSet
            d_1114_v67_ = _dafny.MultiSet([d_1113_v66_])
            d_1115_v68_: D11
            d_1115_v68_ = D11_DC32(d_1035_v7_)
            pat_let_tv38_ = d_1035_v7_
            d_1116_v69_: _dafny.MultiSet
            def iife80_(_pat_let18_0):
                def iife81_(d_1117_dt__update__tmp_h0_):
                    def iife82_(_pat_let19_0):
                        def iife83_(d_1118_dt__update_hcf53_h0_):
                            return D11_DC32(d_1118_dt__update_hcf53_h0_)
                        return iife83_(_pat_let19_0)
                    return iife82_(pat_let_tv38_)
                return iife81_(_pat_let18_0)
            d_1116_v69_ = _dafny.MultiSet([(iife80_(d_1115_v68_)).cf53, d_1035_v7_])
            if default__.fm1(((d_1114_v67_)[d_1113_v66_] if (d_1113_v66_) in (d_1114_v67_) else (self).f23), (d_1116_v69_).issubset(d_1116_v69_), globalState):
                d_1119_v70_: _dafny.Array
                def lambda50_(d_1120_v5_):
                    def lambda51_(d_1121_i9_):
                        return ((d_1120_v5_).f28) == (not((d_1120_v5_).f28))

                    return lambda51_

                init30_ = lambda50_(d_1033_v5_)
                nw184_ = _dafny.Array(None, 21)
                for i0_30_ in range(nw184_.length(0)):
                    nw184_[i0_30_] = init30_(i0_30_)
                d_1119_v70_ = nw184_
                index213_ = default__.safeIndex(470, (d_1119_v70_).length(0))
                (d_1119_v70_)[index213_] = (d_1033_v5_).f28
                index214_ = default__.safeIndex(470, (d_1119_v70_).length(0))
                (d_1119_v70_)[index214_] = (d_1033_v5_).f28
                d_1122_v71_: _dafny.Array
                nw185_ = _dafny.Array(_dafny.Map({}), 24)
                d_1122_v71_ = nw185_
                d_1123_v72_: _dafny.Map
                d_1123_v72_ = _dafny.Map({(d_1033_v5_).f28: d_1122_v71_})
                d_1123_v72_ = (d_1123_v72_).set(not (not(True)) or (True), d_1122_v71_)
                d_1124_v73_: str
                d_1124_v73_ = _dafny.CodePoint('b')
                (globalState).f11 = d_1124_v73_
                d_1125_v74_: _dafny.Array
                nw186_ = _dafny.Array(None, 7)
                nw186_[int(0)] = d_1035_v7_
                nw186_[int(1)] = (d_1035_v7_) + (d_1035_v7_)
                nw186_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrrj"))).set(default__.safeIndex(default__.fm0((self).f26, (d_1033_v5_).f28, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrrj")))), d_1124_v73_)
                nw186_[int(3)] = d_1035_v7_
                nw186_[int(4)] = default__.fm2(globalState)
                nw186_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1124_v73_ for d_1126_i10_ in range(default__.abs(379))])
                nw186_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etxm"))) + (d_1035_v7_)
                d_1125_v74_ = nw186_
                d_1127_v75_: _dafny.Map
                d_1127_v75_ = _dafny.Map({(d_1033_v5_).f28: d_1110_i7_})
                index215_ = default__.safeIndex(926, (d_1125_v74_).length(0))
                (d_1125_v74_)[index215_] = ((d_1035_v7_) + (d_1035_v7_)).set(default__.safeIndex(((d_1127_v75_)[d_1029_v1_] if (d_1029_v1_) in (d_1127_v75_) else len(d_1127_v75_)), len((d_1035_v7_) + (d_1035_v7_))), d_1124_v73_)
                index216_ = default__.safeIndex(926, (d_1125_v74_).length(0))
                (d_1125_v74_)[index216_] = d_1035_v7_
                d_1128_v76_: _dafny.MultiSet
                d_1128_v76_ = _dafny.MultiSet([d_1124_v73_])
                d_1129_v77_: D8
                d_1129_v77_ = D8_DC20((d_1033_v5_).f28, d_1110_i7_)
                d_1130_v78_: C1
                nw187_ = C1()
                nw187_.ctor__((d_1033_v5_).f28, d_1031_v3_, (d_1033_v5_).f23)
                d_1130_v78_ = nw187_
                d_1131_v79_: _dafny.Array
                nw188_ = _dafny.Array(None, 18)
                nw188_[int(0)] = d_1130_v78_
                nw188_[int(1)] = d_1130_v78_
                nw188_[int(2)] = d_1130_v78_
                nw188_[int(3)] = d_1130_v78_
                nw188_[int(4)] = d_1130_v78_
                nw188_[int(5)] = d_1130_v78_
                nw188_[int(6)] = d_1130_v78_
                nw188_[int(7)] = d_1130_v78_
                nw188_[int(8)] = d_1130_v78_
                nw188_[int(9)] = d_1130_v78_
                nw188_[int(10)] = d_1130_v78_
                nw188_[int(11)] = d_1130_v78_
                nw188_[int(12)] = d_1130_v78_
                nw188_[int(13)] = d_1130_v78_
                nw188_[int(14)] = d_1130_v78_
                nw188_[int(15)] = d_1130_v78_
                nw188_[int(16)] = d_1130_v78_
                nw188_[int(17)] = d_1130_v78_
                d_1131_v79_ = nw188_
                d_1132_v80_: _dafny.Map
                d_1132_v80_ = _dafny.Map({(d_1033_v5_).f28: d_1131_v79_})
                d_1133_v81_: _dafny.Map
                d_1133_v81_ = _dafny.Map({not((d_1119_v70_)[default__.safeIndex(470, (d_1119_v70_).length(0))]): d_1033_v5_.f29})
                d_1134_v82_: _dafny.MultiSet
                d_1134_v82_ = _dafny.MultiSet([len(d_1133_v81_)])
                (globalState).f14 = default__.fm0((((_dafny.MultiSet(d_1033_v5_.f29)).set((d_1128_v76_).cardinality, default__.abs((d_1129_v77_).cf34))).cardinality) - ((0) - (len(default__.fm2(globalState)))), not((_dafny.MultiSet([len(d_1132_v80_), (d_1033_v5_).f23])).ispropersubset(d_1134_v82_)), globalState)
            elif True:
                index217_ = default__.safeIndex(610, ((d_1113_v66_).f27).length(0))
                ((d_1113_v66_).f27)[index217_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoqqdqjyb"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1135_i11_ in range(default__.abs(-545))])))
                index218_ = default__.safeIndex(610, ((d_1113_v66_).f27).length(0))
                ((d_1113_v66_).f27)[index218_] = ((d_1033_v5_).f23) + (len(d_1035_v7_))
                d_1136_v83_: C0
                nw189_ = C0()
                nw189_.ctor__()
                d_1136_v83_ = nw189_
                d_1137_v84_: _dafny.Seq
                d_1137_v84_ = _dafny.SeqWithoutIsStrInference([False])
                d_1138_v85_: str
                d_1138_v85_ = _dafny.CodePoint('l')
                d_1139_v86_: _dafny.Set
                d_1139_v86_ = _dafny.Set({d_1138_v85_})
                d_1140_v87_: D9
                d_1140_v87_ = D9_DC25(True, d_1139_v86_)
                d_1141_v88_: _dafny.Map
                d_1141_v88_ = _dafny.Map({(d_1033_v5_).f28: (self).f26})
                d_1142_v89_: D8
                d_1142_v89_ = D8_DC21((d_1113_v66_).fm10(_dafny.SeqWithoutIsStrInference([d_1138_v85_ for d_1143_i12_ in range(default__.abs(224))]), (d_1033_v5_).f28, d_1141_v88_, (d_1137_v84_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixco"))), len(d_1137_v84_))], globalState), (d_1033_v5_).f28, (d_1033_v5_).f23)
                d_1144_v90_: _dafny.MultiSet
                d_1144_v90_ = _dafny.MultiSet([_dafny.CodePoint('q')])
                d_1145_v92_: D2
                def iife84_():
                    coll44_ = _dafny.Set()
                    compr_44_: str
                    for compr_44_ in (d_1144_v90_).Elements:
                        d_1146_v91_: str = compr_44_
                        if (d_1146_v91_) in (d_1144_v90_):
                            coll44_ = coll44_.union(_dafny.Set([d_1146_v91_]))
                    return _dafny.Set(coll44_)
                d_1145_v92_ = D2_DC7(iife84_()
)
                d_1147_v93_: _dafny.Map
                d_1147_v93_ = _dafny.Map({d_1145_v92_: d_1035_v7_})
                d_1148_v94_: _dafny.Array
                def lambda52_(d_1149_v1_):
                    def lambda53_(d_1150_i13_):
                        return d_1149_v1_

                    return lambda53_

                init31_ = lambda52_(d_1029_v1_)
                nw190_ = _dafny.Array(None, 21)
                for i0_31_ in range(nw190_.length(0)):
                    nw190_[i0_31_] = init31_(i0_31_)
                d_1148_v94_ = nw190_
                d_1151_v95_: _dafny.Set
                d_1151_v95_ = _dafny.Set({d_1031_v3_, d_1033_v5_.f29, d_1031_v3_})
                d_1152_v96_: _dafny.Array
                nw191_ = _dafny.Array(None, 27)
                nw191_[int(0)] = d_1029_v1_
                nw191_[int(1)] = d_1029_v1_
                nw191_[int(2)] = (d_1137_v84_)[default__.safeIndex(len(d_1035_v7_), len(d_1137_v84_))]
                nw191_[int(3)] = d_1029_v1_
                nw191_[int(4)] = False
                nw191_[int(5)] = (d_1140_v87_).cf41
                nw191_[int(6)] = (d_1029_v1_ if (d_1033_v5_).f28 else False)
                nw191_[int(7)] = (d_1033_v5_).f28
                nw191_[int(8)] = (d_1033_v5_).f28
                nw191_[int(9)] = (d_1033_v5_).f28
                nw191_[int(10)] = (d_1033_v5_).f28
                nw191_[int(11)] = (d_1033_v5_).f28
                nw191_[int(12)] = (d_1033_v5_).f28
                nw191_[int(13)] = False
                nw191_[int(14)] = (d_1035_v7_) != (d_1035_v7_)
                nw191_[int(15)] = True
                nw191_[int(16)] = (d_1033_v5_).f28
                nw191_[int(17)] = (d_1142_v89_).cf36
                nw191_[int(18)] = d_1029_v1_
                nw191_[int(19)] = (d_1033_v5_).f28
                nw191_[int(20)] = not(((d_1136_v83_).fm20(d_1147_v93_, len(d_1035_v7_), (d_1033_v5_).f23, (d_1033_v5_).f23, globalState)) not in (_dafny.Map({(d_1033_v5_).f28: d_1148_v94_})))
                nw191_[int(21)] = (d_1033_v5_).f28
                nw191_[int(22)] = (d_1033_v5_).f28
                nw191_[int(23)] = (d_1033_v5_).f28
                nw191_[int(24)] = True
                nw191_[int(25)] = (d_1033_v5_).f28
                nw191_[int(26)] = (d_1151_v95_).issubset(d_1151_v95_)
                d_1152_v96_ = nw191_
                d_1153_v97_: _dafny.Seq
                d_1153_v97_ = _dafny.SeqWithoutIsStrInference([d_1152_v96_, d_1152_v96_, d_1152_v96_])
                d_1152_v96_ = (d_1153_v97_)[default__.safeIndex((d_1033_v5_).f23, len(d_1153_v97_))]
                (globalState).f8 = default__.fm0(len(_dafny.SeqWithoutIsStrInference([(self).f26 for d_1154_i14_ in range(default__.abs(545))])), d_1029_v1_, globalState)
                d_1155_v98_: _dafny.MultiSet
                d_1155_v98_ = _dafny.MultiSet([d_1137_v84_])
                (globalState).f8 = (((d_1155_v98_)[d_1137_v84_] if (d_1137_v84_) in (d_1155_v98_) else (self).f23)) + (((d_1113_v66_).f27)[default__.safeIndex(610, ((d_1113_v66_).f27).length(0))])
            r0 = d_1029_v1_
            d_1156_v99_: str
            d_1156_v99_ = _dafny.CodePoint('g')
            (globalState).f21 = (d_1156_v99_) in (_dafny.SeqWithoutIsStrInference([(D4_DC15((self).f26, d_1156_v99_, len(d_1035_v7_))).cf27 for d_1157_i15_ in range(default__.abs(931))]))
            source9_ = D8_DC22(d_1029_v1_)
            if source9_.is_DC20:
                d_1158___mcc_h4_ = source9_.cf33
                d_1159___mcc_h5_ = source9_.cf34
                d_1160_cf34_ = d_1159___mcc_h5_
                d_1161_cf33_ = d_1158___mcc_h4_
                (globalState).f21 = default__.fm1((self).f26, default__.fm1((d_1033_v5_).f23, not(d_1161_cf33_), globalState), globalState)
                index219_ = default__.safeIndex(182, ((d_1113_v66_).f27).length(0))
                ((d_1113_v66_).f27)[index219_] = 953
                d_1162_v100_: _dafny.Set
                d_1162_v100_ = _dafny.Set({d_1029_v1_})
                index220_ = default__.safeIndex(182, ((d_1113_v66_).f27).length(0))
                rhs230_ = ((d_1035_v7_)[default__.safeIndex(d_1110_i7_, len(d_1035_v7_))] if d_1161_cf33_ else _dafny.CodePoint('b'))
                rhs231_ = default__.fm0(default__.safeDivisionInt(len(d_1162_v100_), len(d_1033_v5_.f29)), d_1161_cf33_, globalState)
                rhs232_ = default__.fm6((self).f25, globalState)
                lhs208_ = (d_1113_v66_).f27
                lhs209_ = default__.safeIndex(182, ((d_1113_v66_).f27).length(0))
                lhs210_ = globalState
                d_1156_v99_ = rhs230_
                lhs208_[lhs209_] = rhs231_
                lhs210_.f11 = rhs232_
                d_1163_v101_: _dafny.Array
                def lambda54_(d_1164_v5_):
                    def lambda55_(d_1165_i16_):
                        return (d_1164_v5_).f28

                    return lambda55_

                init32_ = lambda54_(d_1033_v5_)
                nw192_ = _dafny.Array(None, 5)
                for i0_32_ in range(nw192_.length(0)):
                    nw192_[i0_32_] = init32_(i0_32_)
                d_1163_v101_ = nw192_
                d_1166_v102_: _dafny.Map
                d_1166_v102_ = _dafny.Map({d_1163_v101_: (((d_1113_v66_).f27)[default__.safeIndex(182, ((d_1113_v66_).f27).length(0))]) <= ((_dafny.MultiSet([d_1029_v1_, (d_1033_v5_).f28])).cardinality)})
                d_1167_v103_: _dafny.Map
                d_1167_v103_ = _dafny.Map({(d_1033_v5_).f28: ((d_1113_v66_).f27)[default__.safeIndex(182, ((d_1113_v66_).f27).length(0))]})
                d_1166_v102_ = (d_1166_v102_).set(d_1163_v101_, (d_1113_v66_).fm10(d_1035_v7_, (d_1033_v5_).f28, d_1167_v103_, d_1029_v1_, globalState))
                rhs233_ = _dafny.CodePoint('v')
                rhs234_ = (0) - ((self).f26)
                lhs211_ = globalState
                lhs212_ = globalState
                lhs211_.f11 = rhs233_
                lhs212_.f8 = rhs234_
            elif source9_.is_DC21:
                d_1168___mcc_h6_ = source9_.cf35
                d_1169___mcc_h7_ = source9_.cf36
                d_1170___mcc_h8_ = source9_.cf37
                d_1171_cf37_ = d_1170___mcc_h8_
                d_1172_cf36_ = d_1169___mcc_h7_
                d_1173_cf35_ = d_1168___mcc_h6_
                index221_ = default__.safeIndex(129, ((d_1113_v66_).f27).length(0))
                ((d_1113_v66_).f27)[index221_] = d_1171_cf37_
                index222_ = default__.safeIndex(129, ((d_1113_v66_).f27).length(0))
                ((d_1113_v66_).f27)[index222_] = (0) - (d_1110_i7_)
                d_1174_v104_: _dafny.Array
                def lambda56_(d_1175_v5_):
                    def lambda57_(d_1176_i17_):
                        return (d_1175_v5_).f28

                    return lambda57_

                init33_ = lambda56_(d_1033_v5_)
                nw193_ = _dafny.Array(None, 6)
                for i0_33_ in range(nw193_.length(0)):
                    nw193_[i0_33_] = init33_(i0_33_)
                d_1174_v104_ = nw193_
                index223_ = default__.safeIndex(284, (d_1174_v104_).length(0))
                (d_1174_v104_)[index223_] = (d_1033_v5_).f28
                index224_ = default__.safeIndex(284, (d_1174_v104_).length(0))
                (d_1174_v104_)[index224_] = True
                index225_ = default__.safeIndex(284, (d_1174_v104_).length(0))
                (d_1174_v104_)[index225_] = not((d_1156_v99_) not in (d_1035_v7_))
                d_1177_v105_: _dafny.Map
                d_1177_v105_ = _dafny.Map({d_1174_v104_: d_1173_cf35_})
                d_1177_v105_ = d_1177_v105_
            elif source9_.is_DC22:
                d_1178___mcc_h9_ = source9_.cf38
                d_1179_cf38_ = d_1178___mcc_h9_
                d_1180_v106_: _dafny.Array
                def lambda58_(d_1181_v5_):
                    def lambda59_(d_1182_i18_):
                        return (d_1181_v5_).f28

                    return lambda59_

                init34_ = lambda58_(d_1033_v5_)
                nw194_ = _dafny.Array(None, 8)
                for i0_34_ in range(nw194_.length(0)):
                    nw194_[i0_34_] = init34_(i0_34_)
                d_1180_v106_ = nw194_
                index226_ = default__.safeIndex(436, (d_1180_v106_).length(0))
                (d_1180_v106_)[index226_] = d_1029_v1_
                index227_ = default__.safeIndex(436, (d_1180_v106_).length(0))
                (d_1180_v106_)[index227_] = d_1179_cf38_
                (globalState).f11 = _dafny.CodePoint('l')
                d_1183_v107_: C0
                nw195_ = C0()
                nw195_.ctor__()
                d_1183_v107_ = nw195_
                d_1032_v4_ = (d_1032_v4_).set(not(True), (d_1110_i7_) in (d_1033_v5_.f29))
            elif source9_.is_DC19:
                d_1184___mcc_h10_ = source9_.cf32
                d_1185_cf32_ = d_1184___mcc_h10_
                d_1186_v108_: _dafny.Map
                d_1186_v108_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vflk")): (d_1033_v5_).f28})
                d_1187_v109_: _dafny.Array
                nw196_ = _dafny.Array(None, 21)
                nw196_[int(0)] = d_1029_v1_
                nw196_[int(1)] = (d_1033_v5_).f28
                nw196_[int(2)] = (d_1033_v5_).f28
                nw196_[int(3)] = (d_1033_v5_).f28
                nw196_[int(4)] = (d_1033_v5_).f28
                nw196_[int(5)] = (d_1033_v5_).f28
                nw196_[int(6)] = d_1029_v1_
                nw196_[int(7)] = (d_1033_v5_).f28
                nw196_[int(8)] = (d_1033_v5_).f28
                nw196_[int(9)] = (d_1033_v5_).f28
                nw196_[int(10)] = (D8_DC20((d_1033_v5_).f28, (self).f23)).cf33
                nw196_[int(11)] = False
                nw196_[int(12)] = d_1029_v1_
                nw196_[int(13)] = True
                nw196_[int(14)] = d_1029_v1_
                nw196_[int(15)] = (d_1033_v5_).f28
                nw196_[int(16)] = ((d_1186_v108_)[d_1035_v7_] if (d_1035_v7_) in (d_1186_v108_) else (d_1033_v5_).f28)
                nw196_[int(17)] = default__.fm1(d_1110_i7_, d_1029_v1_, globalState)
                nw196_[int(18)] = d_1029_v1_
                nw196_[int(19)] = (d_1033_v5_).f28
                nw196_[int(20)] = not((d_1033_v5_).f28)
                d_1187_v109_ = nw196_
                d_1188_v110_: _dafny.Map
                d_1188_v110_ = _dafny.Map({d_1187_v109_: d_1156_v99_})
                (globalState).f11 = ((d_1188_v110_)[d_1187_v109_] if (d_1187_v109_) in (d_1188_v110_) else d_1156_v99_)
                d_1189_v111_: D0
                d_1190_v112_: _dafny.Map
                d_1191_v113_: int
                out40_: D0
                out41_: _dafny.Map
                out42_: int
                out40_, out41_, out42_ = default__.m0(d_1187_v109_, (d_1035_v7_) + (d_1035_v7_), d_1110_i7_, globalState)
                d_1189_v111_ = out40_
                d_1190_v112_ = out41_
                d_1191_v113_ = out42_
                (globalState).f14 = ((self).f25) - (d_1110_i7_)
                d_1192_v114_: _dafny.Array
                def lambda60_(d_1193_v99_):
                    def lambda61_(d_1194_i19_):
                        return d_1193_v99_

                    return lambda61_

                init35_ = lambda60_(d_1156_v99_)
                nw197_ = _dafny.Array(None, 5)
                for i0_35_ in range(nw197_.length(0)):
                    nw197_[i0_35_] = init35_(i0_35_)
                d_1192_v114_ = nw197_
                d_1195_v115_: _dafny.Seq
                d_1195_v115_ = _dafny.SeqWithoutIsStrInference([(d_1033_v5_).f28])
                d_1196_v116_: D3
                d_1196_v116_ = D3_DC12(d_1195_v115_)
                d_1197_v117_: _dafny.Map
                d_1197_v117_ = _dafny.Map({D9_DC24(d_1192_v114_): d_1196_v116_})
                d_1197_v117_ = d_1197_v117_
            elif True:
                d_1198___mcc_h11_ = source9_.cf39
                d_1199_cf39_ = d_1198___mcc_h11_
                d_1200_v118_: _dafny.Map
                d_1200_v118_ = _dafny.Map({(d_1033_v5_).f28: (d_1033_v5_).f23})
                d_1201_v119_: _dafny.Seq
                d_1201_v119_ = _dafny.SeqWithoutIsStrInference([d_1200_v118_])
                d_1032_v4_ = (d_1032_v4_).set((d_1113_v66_).fm10(d_1035_v7_, d_1029_v1_, (_dafny.Map({(d_1113_v66_).fm10(_dafny.SeqWithoutIsStrInference([d_1156_v99_ for d_1202_i20_ in range(default__.abs(-438))]), (d_1033_v5_).f28, d_1200_v118_, True, globalState): 922})).set((d_1033_v5_).f28, default__.fm0(len((d_1201_v119_)[default__.safeIndex((self).f25, len(d_1201_v119_))]), (d_1033_v5_).f28, globalState)), (d_1033_v5_).f28, globalState), d_1029_v1_)
                d_1203_v120_: _dafny.Array
                nw198_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_1203_v120_ = nw198_
                d_1204_v121_: _dafny.Array
                nw199_ = _dafny.Array(False, 25)
                d_1204_v121_ = nw199_
                index228_ = default__.safeIndex(181, (d_1203_v120_).length(0))
                (d_1203_v120_)[index228_] = d_1204_v121_
                d_1205_v122_: _dafny.MultiSet
                d_1205_v122_ = _dafny.MultiSet([d_1156_v99_])
                d_1206_v123_: _dafny.Map
                d_1206_v123_ = _dafny.Map({(self).f23: (d_1033_v5_).f28})
                index229_ = default__.safeIndex(181, (d_1203_v120_).length(0))
                def iife85_():
                    coll45_ = _dafny.Map()
                    compr_45_: int
                    for compr_45_ in _dafny.IntegerRange(-391, 524):
                        d_1207_v124_: int = compr_45_
                        if ((-391) <= (d_1207_v124_)) and ((d_1207_v124_) < (524)):
                            coll45_[default__.safeModuloInt(d_1207_v124_, 909)] = (0) - ((d_1033_v5_).f23)
                    return _dafny.Map(coll45_)
                rhs235_ = default__.safeModuloInt(default__.safeDivisionInt(len(d_1036_v8_), (self).f25), 491)
                rhs236_ = ((d_1032_v4_)[not((d_1033_v5_).f28)] if (not((d_1033_v5_).f28)) in (d_1032_v4_) else (d_1156_v99_) in ((d_1035_v7_).set(default__.safeIndex((d_1033_v5_).f23, len(d_1035_v7_)), d_1156_v99_)))
                rhs237_ = d_1204_v121_
                rhs238_ = ((d_1205_v122_)[d_1156_v99_] if (d_1156_v99_) in (d_1205_v122_) else 645)
                rhs239_ = (((d_1206_v123_).set(-753, (d_1033_v5_).f28)) | ((_dafny.Map({d_1110_i7_: (d_1033_v5_).f28})).set((self).f25, d_1029_v1_))).set(default__.safeDivisionInt(len(iife85_()
                ), (d_1033_v5_).f23), True)
                lhs213_ = globalState
                lhs214_ = d_1203_v120_
                lhs215_ = default__.safeIndex(181, (d_1203_v120_).length(0))
                lhs216_ = globalState
                lhs213_.f14 = rhs235_
                d_1029_v1_ = rhs236_
                lhs214_[lhs215_] = rhs237_
                lhs216_.f14 = rhs238_
                r2 = rhs239_
                d_1206_v123_ = (d_1206_v123_).set(d_1110_i7_, (d_1033_v5_).f28)
                d_1208_v125_: _dafny.Seq
                d_1208_v125_ = _dafny.SeqWithoutIsStrInference([(d_1033_v5_).f28])
                d_1209_v126_: _dafny.Map
                d_1209_v126_ = _dafny.Map({d_1208_v125_: ((self).f23) < ((self).f23)})
                rhs240_ = d_1035_v7_
                rhs241_ = (d_1033_v5_).f28
                rhs242_ = (d_1209_v126_) | (_dafny.Map({d_1208_v125_: (d_1033_v5_).f28}))
                rhs243_ = not((d_1033_v5_).f28)
                rhs244_ = ((-332) + ((self).f23)) + ((d_1033_v5_).f23)
                lhs217_ = globalState
                lhs218_ = globalState
                lhs219_ = globalState
                lhs217_.f6 = rhs240_
                r0 = rhs241_
                d_1209_v126_ = rhs242_
                lhs218_.f7 = rhs243_
                lhs219_.f8 = rhs244_
        hi14_ = default__.safeDivisionInt((d_1033_v5_).f23, (self).f26)
        for d_1210_i21_ in range((0) - ((d_1033_v5_).f23), hi14_):
            (globalState).f14 = default__.fm0((self).f25, ((d_1032_v4_)[(d_1033_v5_).f28] if ((d_1033_v5_).f28) in (d_1032_v4_) else (d_1033_v5_).f28), globalState)
            d_1211_v127_: D12
            d_1211_v127_ = D12_DC35()
            d_1212_v128_: _dafny.Array
            nw200_ = _dafny.Array(False, 5)
            d_1212_v128_ = nw200_
            d_1213_v129_: _dafny.Map
            d_1213_v129_ = _dafny.Map({d_1211_v127_: d_1212_v128_})
            d_1214_v130_: int
            d_1215_v131_: _dafny.Set
            out43_: int
            out44_: _dafny.Set
            out43_, out44_ = (d_1033_v5_).m1(((d_1213_v129_)[d_1211_v127_] if (d_1211_v127_) in (d_1213_v129_) else d_1212_v128_), ((d_1033_v5_).f28) and ((d_1033_v5_).f28), globalState)
            d_1214_v130_ = out43_
            d_1215_v131_ = out44_
            d_1216_v132_: _dafny.Array
            def lambda62_(d_1217_v130_):
                def lambda63_(d_1218_i22_):
                    return default__.safeDivisionInt(d_1218_i22_, d_1217_v130_)

                return lambda63_

            init36_ = lambda62_(d_1214_v130_)
            nw201_ = _dafny.Array(None, 15)
            for i0_36_ in range(nw201_.length(0)):
                nw201_[i0_36_] = init36_(i0_36_)
            d_1216_v132_ = nw201_
            d_1219_v133_: _dafny.Seq
            d_1219_v133_ = _dafny.SeqWithoutIsStrInference([d_1216_v132_])
            d_1220_v134_: D1
            d_1220_v134_ = D1_DC3(d_1219_v133_)
            source10_ = d_1220_v134_
            if source10_.is_DC4:
                d_1221___mcc_h12_ = source10_.cf6
                d_1222___mcc_h13_ = source10_.cf7
                d_1223___mcc_h14_ = source10_.cf8
                d_1224_cf8_ = d_1223___mcc_h14_
                d_1225_cf7_ = d_1222___mcc_h13_
                d_1226_cf6_ = d_1221___mcc_h12_
                d_1227_v135_: C3
                nw202_ = C3()
                nw202_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True]) for d_1228_i23_ in range(default__.abs(67))]), (d_1033_v5_).f23, (d_1033_v5_).f28, d_1033_v5_.f29)
                d_1227_v135_ = nw202_
                (globalState).f14 = (self).f23
                d_1229_v136_: C4
                nw203_ = C4()
                nw203_.ctor__(d_1216_v132_, d_1210_i21_)
                d_1229_v136_ = nw203_
                d_1230_v137_: D3
                d_1230_v137_ = D3_DC12(_dafny.SeqWithoutIsStrInference([d_1029_v1_]))
                d_1231_v138_: _dafny.Map
                d_1231_v138_ = _dafny.Map({d_1230_v137_: (d_1033_v5_).f28})
                d_1232_v139_: _dafny.Seq
                d_1232_v139_ = _dafny.SeqWithoutIsStrInference([not((d_1033_v5_).f28), not(d_1224_cf8_)])
                pat_let_tv39_ = d_1232_v139_
                pat_let_tv40_ = d_1232_v139_
                def iife86_(_pat_let20_0):
                    def iife87_(d_1233_dt__update__tmp_h2_):
                        def iife88_(_pat_let21_0):
                            def iife89_(d_1234_dt__update_hcf23_h1_):
                                return D3_DC12(d_1234_dt__update_hcf23_h1_)
                            return iife89_(_pat_let21_0)
                        return iife88_(pat_let_tv39_)
                    return iife87_(_pat_let20_0)
                def iife90_(_pat_let22_0):
                    def iife91_(d_1235_dt__update__tmp_h1_):
                        def iife92_(_pat_let23_0):
                            def iife93_(d_1236_dt__update_hcf23_h0_):
                                return D3_DC12(d_1236_dt__update_hcf23_h0_)
                            return iife93_(_pat_let23_0)
                        return iife92_(pat_let_tv40_)
                    return iife91_(_pat_let22_0)
                (globalState).f21 = ((d_1231_v138_)[iife86_(d_1230_v137_)] if (iife90_(d_1230_v137_)) in (d_1231_v138_) else True)
            elif source10_.is_DC5:
                d_1237___mcc_h15_ = source10_.cf9
                d_1238___mcc_h16_ = source10_.cf10
                d_1239___mcc_h17_ = source10_.cf11
                d_1240___mcc_h18_ = source10_.cf12
                d_1241_cf12_ = d_1240___mcc_h18_
                d_1242_cf11_ = d_1239___mcc_h17_
                d_1243_cf10_ = d_1238___mcc_h16_
                d_1244_cf9_ = d_1237___mcc_h15_
                (globalState).f14 = (d_1210_i21_) - ((d_1033_v5_).f23)
                d_1245_v140_: _dafny.Seq
                d_1245_v140_ = _dafny.SeqWithoutIsStrInference([d_1029_v1_, (d_1033_v5_).f28])
                (globalState).f21 = (d_1243_cf10_ if (d_1245_v140_) <= (d_1245_v140_) else ((self).f26) <= (d_1210_i21_))
                d_1036_v8_ = d_1036_v8_
                d_1246_v141_: C0
                nw204_ = C0()
                nw204_.ctor__()
                d_1246_v141_ = nw204_
            elif source10_.is_DC6:
                d_1247___mcc_h19_ = source10_.cf13
                d_1248_cf13_ = d_1247___mcc_h19_
                (globalState).f8 = (0) - ((self).f23)
                (globalState).f9 = d_1029_v1_
                d_1248_cf13_ = (d_1248_cf13_ if (d_1033_v5_).f28 else d_1248_cf13_)
                d_1249_v142_: _dafny.Map
                d_1249_v142_ = _dafny.Map({d_1033_v5_.f29: d_1210_i21_})
                d_1249_v142_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([default__.fm0(455, d_1029_v1_, globalState), (self).f25])) + (_dafny.SeqWithoutIsStrInference([(d_1033_v5_).f23])): 993})
            elif True:
                d_1250___mcc_h20_ = source10_.cf5
                d_1251_cf5_ = d_1250___mcc_h20_
                d_1252_v143_: _dafny.Map
                d_1252_v143_ = _dafny.Map({(d_1033_v5_).f28: d_1214_v130_})
                d_1253_v144_: _dafny.Seq
                d_1253_v144_ = _dafny.SeqWithoutIsStrInference([(d_1033_v5_).f28])
                d_1252_v143_ = (d_1252_v143_).set(d_1029_v1_, len(d_1253_v144_))
                d_1254_v145_: _dafny.MultiSet
                d_1254_v145_ = _dafny.MultiSet([(d_1033_v5_).f28])
                (globalState).f14 = default__.safeDivisionInt((d_1254_v145_).cardinality, ((0) - ((self).f23)) + ((_dafny.MultiSet([(d_1033_v5_).f23])).cardinality))
                d_1255_v146_: _dafny.Array
                def lambda64_(d_1256_v7_):
                    def lambda65_(d_1257_i24_):
                        return d_1256_v7_

                    return lambda65_

                init37_ = lambda64_(d_1035_v7_)
                nw205_ = _dafny.Array(None, 13)
                for i0_37_ in range(nw205_.length(0)):
                    nw205_[i0_37_] = init37_(i0_37_)
                d_1255_v146_ = nw205_
                index230_ = default__.safeIndex(404, (d_1255_v146_).length(0))
                (d_1255_v146_)[index230_] = d_1035_v7_
                index231_ = default__.safeIndex(404, (d_1255_v146_).length(0))
                (d_1255_v146_)[index231_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1258_i25_ in range(default__.abs(-84))])) + (d_1035_v7_)
                d_1035_v7_ = ((d_1035_v7_) + ((d_1255_v146_)[default__.safeIndex(404, (d_1255_v146_).length(0))])) + (d_1035_v7_)
            if d_1029_v1_:
                d_1259_v147_: _dafny.Array
                def lambda66_(d_1260_v131_):
                    def lambda67_(d_1261_i26_):
                        return d_1260_v131_

                    return lambda67_

                init38_ = lambda66_(d_1215_v131_)
                nw206_ = _dafny.Array(None, 18)
                for i0_38_ in range(nw206_.length(0)):
                    nw206_[i0_38_] = init38_(i0_38_)
                d_1259_v147_ = nw206_
                index232_ = default__.safeIndex(867, (d_1259_v147_).length(0))
                def iife94_():
                    def iife96_():
                        coll48_ = _dafny.Set()
                        compr_48_: int
                        for compr_48_ in _dafny.IntegerRange(836, 653):
                            d_1264_v148_: int = compr_48_
                            if ((836) <= (d_1264_v148_)) and ((d_1264_v148_) < (653)):
                                coll48_ = coll48_.union(_dafny.Set([(d_1264_v148_) * (25)]))
                        return _dafny.Set(coll48_)
                    coll46_ = _dafny.Set()
                    def iife95_():
                        coll47_ = _dafny.Set()
                        compr_47_: int
                        for compr_47_ in _dafny.IntegerRange(836, 653):
                            d_1262_v148_: int = compr_47_
                            if ((836) <= (d_1262_v148_)) and ((d_1262_v148_) < (653)):
                                coll47_ = coll47_.union(_dafny.Set([(d_1262_v148_) * (25)]))
                        return _dafny.Set(coll47_)
                    compr_46_: int
                    for compr_46_ in (iife95_()
                    ).Elements:
                        d_1263_v149_: int = compr_46_
                        if (d_1263_v149_) in (iife96_()
                        ):
                            coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1263_v149_, (_dafny.MultiSet([_dafny.CodePoint('b'), _dafny.CodePoint('e')])).cardinality)]))
                    return _dafny.Set(coll46_)
                (d_1259_v147_)[index232_] = iife94_()
                
                index233_ = default__.safeIndex(867, (d_1259_v147_).length(0))
                (d_1259_v147_)[index233_] = d_1215_v131_
                d_1265_v150_: int
                d_1266_v151_: _dafny.Seq
                out45_: int
                out46_: _dafny.Seq
                out45_, out46_ = (d_1033_v5_).m2((d_1035_v7_) + (d_1035_v7_), globalState)
                d_1265_v150_ = out45_
                d_1266_v151_ = out46_
                d_1267_v152_: D0
                d_1267_v152_ = D0_DC1((0) - (d_1214_v130_), d_1029_v1_, (self).f26)
                rhs245_ = (d_1033_v5_).f23
                rhs246_ = (d_1033_v5_).f28
                rhs247_ = (d_1267_v152_).cf2
                lhs220_ = globalState
                lhs221_ = globalState
                lhs220_.f14 = rhs245_
                d_1029_v1_ = rhs246_
                lhs221_.f21 = rhs247_
                d_1268_v153_: _dafny.Map
                d_1268_v153_ = _dafny.Map({(d_1033_v5_).f23: (d_1033_v5_).f28})
                d_1269_v154_: _dafny.Map
                d_1269_v154_ = _dafny.Map({d_1029_v1_: d_1214_v130_})
                index234_ = default__.safeIndex(568, (d_1212_v128_).length(0))
                (d_1212_v128_)[index234_] = not((len(d_1268_v153_)) > ((0) - (((d_1269_v154_)[(d_1033_v5_).f28] if ((d_1033_v5_).f28) in (d_1269_v154_) else (d_1033_v5_).f23))))
                d_1270_v155_: str
                d_1270_v155_ = _dafny.CodePoint('u')
                d_1271_v156_: D11
                d_1271_v156_ = D11_DC32(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxdqyef")))
                d_1272_v157_: _dafny.Array
                nw207_ = _dafny.Array(None, 26)
                nw207_[int(0)] = d_1035_v7_
                nw207_[int(1)] = (d_1035_v7_) + (d_1035_v7_)
                nw207_[int(2)] = d_1035_v7_
                nw207_[int(3)] = (d_1036_v8_)[default__.safeIndex(default__.fm0((d_1033_v5_).f23, (d_1033_v5_).f28, globalState), len(d_1036_v8_))]
                nw207_[int(4)] = (d_1035_v7_).set(default__.safeIndex(len((d_1036_v8_)[default__.safeIndex((self).f23, len(d_1036_v8_))]), len(d_1035_v7_)), d_1270_v155_)
                nw207_[int(5)] = d_1035_v7_
                nw207_[int(6)] = (d_1035_v7_) + ((d_1035_v7_).set(default__.safeIndex(d_1265_v150_, len(d_1035_v7_)), d_1270_v155_))
                nw207_[int(7)] = d_1035_v7_
                nw207_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1273_i27_ in range(default__.abs(525))])
                nw207_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfynwpxj"))) + (d_1035_v7_)
                nw207_[int(10)] = (d_1035_v7_) + (d_1035_v7_)
                nw207_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1270_v155_ for d_1274_i28_ in range(default__.abs(62))])
                nw207_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1270_v155_ for d_1275_i29_ in range(default__.abs(-834))])
                nw207_[int(13)] = (d_1271_v156_).cf53
                nw207_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                nw207_[int(15)] = default__.fm2(globalState)
                nw207_[int(16)] = d_1035_v7_
                nw207_[int(17)] = (d_1035_v7_) + (d_1035_v7_)
                nw207_[int(18)] = d_1035_v7_
                nw207_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
                nw207_[int(20)] = d_1035_v7_
                nw207_[int(21)] = d_1035_v7_
                nw207_[int(22)] = d_1035_v7_
                nw207_[int(23)] = d_1035_v7_
                nw207_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esqnlyvg"))
                nw207_[int(25)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkq"))) + (d_1035_v7_)
                d_1272_v157_ = nw207_
                index235_ = default__.safeIndex(881, (d_1272_v157_).length(0))
                (d_1272_v157_)[index235_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxukkyffw")) if (d_1033_v5_).f28 else d_1035_v7_)
                d_1276_v158_: _dafny.Array
                nw208_ = _dafny.Array(None, 1)
                nw208_[int(0)] = d_1028_v0_
                d_1276_v158_ = nw208_
                d_1277_v159_: D11
                d_1277_v159_ = D11_DC30(d_1276_v158_)
                d_1278_v160_: _dafny.Map
                d_1278_v160_ = _dafny.Map({(d_1270_v155_ if True else _dafny.CodePoint('e')): (d_1033_v5_).f28})
                d_1279_v161_: _dafny.Map
                d_1279_v161_ = _dafny.Map({d_1035_v7_: d_1035_v7_})
                index236_ = default__.safeIndex(568, (d_1212_v128_).length(0))
                index237_ = default__.safeIndex(881, (d_1272_v157_).length(0))
                rhs248_ = not(((d_1033_v5_).f23) < ((d_1033_v5_).f23))
                rhs249_ = ((((d_1279_v161_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "culpysq"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "culpysq"))) in (d_1279_v161_) else d_1035_v7_)) + (d_1035_v7_)) + ((d_1035_v7_) + (d_1035_v7_))
                rhs250_ = d_1277_v159_
                rhs251_ = (d_1278_v160_) | (d_1278_v160_)
                lhs222_ = d_1212_v128_
                lhs223_ = default__.safeIndex(568, (d_1212_v128_).length(0))
                lhs224_ = d_1272_v157_
                lhs225_ = default__.safeIndex(881, (d_1272_v157_).length(0))
                lhs222_[lhs223_] = rhs248_
                lhs224_[lhs225_] = rhs249_
                d_1277_v159_ = rhs250_
                d_1278_v160_ = rhs251_
                d_1035_v7_ = (d_1272_v157_)[default__.safeIndex(881, (d_1272_v157_).length(0))]
            elif True:
                d_1036_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp")) for d_1280_i30_ in range(default__.abs(217))])
                (globalState).f8 = ((self).f23) + ((d_1033_v5_).f23)
                d_1281_v162_: _dafny.Map
                d_1281_v162_ = _dafny.Map({(d_1033_v5_).f28: (self).f26})
                d_1282_v163_: C1
                nw209_ = C1()
                nw209_.ctor__((d_1033_v5_).f28, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1281_v162_])).cardinality]), (d_1033_v5_).f23)
                d_1282_v163_ = nw209_
                index238_ = default__.safeIndex(872, (d_1216_v132_).length(0))
                (d_1216_v132_)[index238_] = (self).f25
                d_1283_v164_: str
                d_1283_v164_ = _dafny.CodePoint('i')
                d_1284_v165_: _dafny.MultiSet
                d_1284_v165_ = _dafny.MultiSet([d_1283_v164_, d_1283_v164_])
                index239_ = default__.safeIndex(872, (d_1216_v132_).length(0))
                (d_1216_v132_)[index239_] = (((self).f26) + (((d_1284_v165_)[d_1283_v164_] if (d_1283_v164_) in (d_1284_v165_) else (self).f26))) + ((self).f26)
                d_1285_v166_: D11
                d_1285_v166_ = D11_DC31(d_1210_i21_)
                d_1286_v167_: _dafny.Seq
                d_1286_v167_ = _dafny.SeqWithoutIsStrInference([d_1285_v166_, d_1285_v166_, d_1285_v166_])
                d_1287_v168_: _dafny.Map
                d_1287_v168_ = _dafny.Map({d_1286_v167_: d_1033_v5_.f29})
                d_1288_v170_: _dafny.Set
                d_1288_v170_ = _dafny.Set({default__.fm32(globalState)})
                d_1289_v172_: _dafny.MultiSet
                d_1289_v172_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1285_v166_]), d_1286_v167_, (_dafny.SeqWithoutIsStrInference([d_1285_v166_])).set(default__.safeIndex((d_1033_v5_).f23, len(_dafny.SeqWithoutIsStrInference([d_1285_v166_]))), d_1285_v166_), d_1286_v167_])
                d_1290_v173_: _dafny.Array
                nw210_ = _dafny.Array(None, 26)
                nw210_[int(0)] = (d_1287_v168_).set(d_1286_v167_, d_1033_v5_.f29)
                nw210_[int(1)] = d_1287_v168_
                nw210_[int(2)] = _dafny.Map({d_1286_v167_: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f26]), _dafny.MultiSet(d_1033_v5_.f29), _dafny.MultiSet([d_1214_v130_, 383]), _dafny.MultiSet([d_1210_i21_, (d_1216_v132_)[default__.safeIndex(872, (d_1216_v132_).length(0))]])])) for d_1291_i31_ in range(default__.abs(428))])})
                nw210_[int(3)] = (d_1287_v168_) | ((D15_DC43(d_1287_v168_)).cf67)
                def iife97_():
                    coll49_ = _dafny.Map()
                    compr_49_: _dafny.Seq
                    for compr_49_ in (d_1288_v170_).Elements:
                        d_1292_v169_: _dafny.Seq = compr_49_
                        if (d_1292_v169_) in (d_1288_v170_):
                            coll49_[d_1292_v169_] = d_1033_v5_.f29
                    return _dafny.Map(coll49_)
                nw210_[int(4)] = iife97_()
                
                nw210_[int(5)] = d_1287_v168_
                nw210_[int(6)] = d_1287_v168_
                nw210_[int(7)] = _dafny.Map({d_1286_v167_: d_1033_v5_.f29})
                nw210_[int(8)] = d_1287_v168_
                nw210_[int(9)] = d_1287_v168_
                nw210_[int(10)] = d_1287_v168_
                nw210_[int(11)] = d_1287_v168_
                def iife98_():
                    coll50_ = _dafny.Map()
                    compr_50_: _dafny.Seq
                    for compr_50_ in (d_1289_v172_).Elements:
                        d_1293_v171_: _dafny.Seq = compr_50_
                        if (d_1293_v171_) in (d_1289_v172_):
                            coll50_[d_1293_v171_] = d_1031_v3_
                    return _dafny.Map(coll50_)
                nw210_[int(12)] = iife98_()
                
                nw210_[int(13)] = d_1287_v168_
                nw210_[int(14)] = (d_1287_v168_ if (d_1033_v5_).f28 else d_1287_v168_)
                nw210_[int(15)] = default__.fm33(len(_dafny.Set({len(d_1033_v5_.f29), d_1210_i21_, (0) - ((self).f23)})), (d_1033_v5_).f28, globalState)
                nw210_[int(16)] = d_1287_v168_
                nw210_[int(17)] = d_1287_v168_
                nw210_[int(18)] = d_1287_v168_
                nw210_[int(19)] = d_1287_v168_
                nw210_[int(20)] = d_1287_v168_
                nw210_[int(21)] = (d_1287_v168_) | (d_1287_v168_)
                nw210_[int(22)] = _dafny.Map({d_1286_v167_: d_1031_v3_})
                nw210_[int(23)] = d_1287_v168_
                nw210_[int(24)] = (d_1287_v168_) | ((d_1287_v168_).set(d_1286_v167_, d_1033_v5_.f29))
                nw210_[int(25)] = d_1287_v168_
                d_1290_v173_ = nw210_
                index240_ = default__.safeIndex(41, (d_1290_v173_).length(0))
                (d_1290_v173_)[index240_] = d_1287_v168_
                d_1294_v174_: _dafny.Seq
                d_1294_v174_ = _dafny.SeqWithoutIsStrInference([(d_1033_v5_).f28])
                index241_ = default__.safeIndex(41, (d_1290_v173_).length(0))
                (d_1290_v173_)[index241_] = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_1285_v166_ for d_1295_i32_ in range(default__.abs(836))])).set(default__.safeIndex(len(d_1294_v174_), len(_dafny.SeqWithoutIsStrInference([d_1285_v166_ for d_1296_i32_ in range(default__.abs(836))]))), d_1285_v166_): (d_1033_v5_.f29) + (_dafny.SeqWithoutIsStrInference([(d_1216_v132_)[default__.safeIndex(872, (d_1216_v132_).length(0))]]))})
        d_1297_v175_: _dafny.Set
        d_1297_v175_ = _dafny.Set({(d_1035_v7_ if (d_1033_v5_).f28 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "seef")))})
        d_1298_v176_: _dafny.Map
        d_1298_v176_ = _dafny.Map({(self).f25: d_1297_v175_})
        d_1299_v177_: D16
        d_1299_v177_ = D16_DC46(d_1297_v175_)
        d_1297_v175_ = ((d_1298_v176_)[(self).f26] if ((self).f26) in (d_1298_v176_) else (d_1297_v175_ if (d_1033_v5_).f28 else (d_1299_v177_).cf74))
        d_1300_v178_: _dafny.Map
        d_1300_v178_ = _dafny.Map({d_1035_v7_: (d_1033_v5_).f28})
        r0 = not(((d_1300_v178_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulwvqx"))) + (d_1035_v7_)] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulwvqx"))) + (d_1035_v7_)) in (d_1300_v178_) else (d_1033_v5_).f28))
        d_1301_v179_: _dafny.Map
        d_1301_v179_ = _dafny.Map({d_1029_v1_: (self).f23})
        d_1302_v180_: _dafny.Seq
        d_1302_v180_ = _dafny.SeqWithoutIsStrInference([d_1301_v179_, d_1301_v179_])
        d_1303_v181_: _dafny.Set
        d_1303_v181_ = _dafny.Set({len(d_1036_v8_)})
        d_1304_v182_: _dafny.Set
        d_1304_v182_ = _dafny.Set({_dafny.Set({(0) - (len(d_1303_v181_)), 957, 843, (self).f26})})
        r1 = (((d_1302_v180_) + (_dafny.SeqWithoutIsStrInference([d_1301_v179_ for d_1305_i33_ in range(default__.abs(681))]))) + ((d_1302_v180_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1033_v5_).f28: (self).f25}), d_1301_v179_, d_1301_v179_, d_1301_v179_])))).set(default__.safeIndex(len(d_1304_v182_), len(((d_1302_v180_) + (_dafny.SeqWithoutIsStrInference([d_1301_v179_ for d_1306_i33_ in range(default__.abs(681))]))) + ((d_1302_v180_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1033_v5_).f28: (self).f25}), d_1301_v179_, d_1301_v179_, d_1301_v179_]))))), d_1301_v179_)
        r2 = default__.fm5(d_1303_v181_, globalState)
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        (globalState).f7 = p0
        d_1307_v0_: D11
        d_1307_v0_ = D11_DC31((self).f25)
        d_1308_v1_: _dafny.Seq
        d_1308_v1_ = _dafny.SeqWithoutIsStrInference([d_1307_v0_, d_1307_v0_])
        d_1309_v2_: _dafny.Seq
        d_1309_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvpo"))
        d_1310_v3_: _dafny.Seq
        d_1310_v3_ = _dafny.SeqWithoutIsStrInference([len(d_1309_v2_), (self).f25])
        d_1311_v4_: _dafny.Map
        d_1311_v4_ = _dafny.Map({d_1308_v1_: d_1310_v3_})
        d_1312_v5_: D15
        d_1312_v5_ = D15_DC43((d_1311_v4_) | (d_1311_v4_))
        source11_ = d_1312_v5_
        if source11_.is_DC44:
            d_1313___mcc_h0_ = source11_.cf68
            d_1314___mcc_h1_ = source11_.cf69
            d_1315___mcc_h2_ = source11_.cf70
            d_1316___mcc_h3_ = source11_.cf71
            d_1317___mcc_h4_ = source11_.cf72
            d_1318_cf72_ = d_1317___mcc_h4_
            d_1319_cf71_ = d_1316___mcc_h3_
            d_1320_cf70_ = d_1315___mcc_h2_
            d_1321_cf69_ = d_1314___mcc_h1_
            d_1322_cf68_ = d_1313___mcc_h0_
            d_1323_v6_: _dafny.Array
            nw211_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1323_v6_ = nw211_
            d_1324_v7_: _dafny.Map
            d_1324_v7_ = _dafny.Map({_dafny.CodePoint('u'): p0})
            d_1325_v8_: D12
            d_1325_v8_ = D12_DC34(d_1324_v7_)
            pat_let_tv41_ = d_1324_v7_
            pat_let_tv42_ = d_1324_v7_
            pat_let_tv43_ = d_1324_v7_
            d_1326_v9_: _dafny.Array
            nw212_ = _dafny.Array(None, 21)
            nw212_[int(0)] = D12_DC34(d_1324_v7_)
            nw212_[int(1)] = D12_DC34(d_1324_v7_)
            nw212_[int(2)] = d_1325_v8_
            nw212_[int(3)] = d_1325_v8_
            def iife99_(_pat_let24_0):
                def iife100_(d_1327_dt__update__tmp_h0_):
                    def iife101_(_pat_let25_0):
                        def iife102_(d_1328_dt__update_hcf55_h0_):
                            return D12_DC34(d_1328_dt__update_hcf55_h0_)
                        return iife102_(_pat_let25_0)
                    return iife101_(pat_let_tv41_)
                return iife100_(_pat_let24_0)
            nw212_[int(4)] = iife99_(d_1325_v8_)
            nw212_[int(5)] = d_1325_v8_
            nw212_[int(6)] = d_1325_v8_
            def iife104_(_pat_let27_0):
                def iife105_(d_1329_dt__update__tmp_h1_):
                    def iife106_(_pat_let28_0):
                        def iife107_(d_1330_dt__update_hcf55_h1_):
                            return D12_DC34(d_1330_dt__update_hcf55_h1_)
                        return iife107_(_pat_let28_0)
                    return iife106_(pat_let_tv42_)
                return iife105_(_pat_let27_0)
            def iife103_(_pat_let26_0):
                def iife108_(d_1331_dt__update__tmp_h2_):
                    def iife109_(_pat_let29_0):
                        def iife110_(d_1332_dt__update_hcf55_h2_):
                            return D12_DC34(d_1332_dt__update_hcf55_h2_)
                        return iife110_(_pat_let29_0)
                    return iife109_(pat_let_tv43_)
                return iife108_(_pat_let26_0)
            nw212_[int(7)] = iife103_(iife104_(d_1325_v8_))
            nw212_[int(8)] = d_1325_v8_
            nw212_[int(9)] = D12_DC34(d_1324_v7_)
            nw212_[int(10)] = d_1325_v8_
            nw212_[int(11)] = d_1325_v8_
            nw212_[int(12)] = d_1325_v8_
            nw212_[int(13)] = d_1325_v8_
            nw212_[int(14)] = D12_DC34(d_1324_v7_)
            nw212_[int(15)] = d_1325_v8_
            nw212_[int(16)] = d_1325_v8_
            nw212_[int(17)] = d_1325_v8_
            nw212_[int(18)] = d_1325_v8_
            nw212_[int(19)] = d_1325_v8_
            nw212_[int(20)] = d_1325_v8_
            d_1326_v9_ = nw212_
            index242_ = default__.safeIndex(643, (d_1323_v6_).length(0))
            (d_1323_v6_)[index242_] = d_1326_v9_
            index243_ = default__.safeIndex(643, (d_1323_v6_).length(0))
            (d_1323_v6_)[index243_] = d_1326_v9_
            d_1333_v10_: _dafny.Map
            d_1333_v10_ = _dafny.Map({d_1318_cf72_: d_1322_cf68_})
            d_1334_v11_: _dafny.Seq
            d_1334_v11_ = _dafny.SeqWithoutIsStrInference([d_1309_v2_])
            d_1335_v12_: _dafny.Seq
            d_1335_v12_ = _dafny.SeqWithoutIsStrInference([d_1334_v11_])
            d_1336_v13_: _dafny.Map
            d_1336_v13_ = _dafny.Map({(self).f25: d_1319_cf71_})
            d_1337_v14_: _dafny.Set
            d_1337_v14_ = _dafny.Set({d_1318_cf72_})
            d_1338_v15_: _dafny.Array
            nw213_ = _dafny.Array(None, 27)
            nw213_[int(0)] = len(d_1333_v10_)
            nw213_[int(1)] = (self).f26
            nw213_[int(2)] = len((d_1335_v12_)[default__.safeIndex((_dafny.MultiSet([not(d_1319_cf71_)])).cardinality, len(d_1335_v12_))])
            nw213_[int(3)] = len(d_1336_v13_)
            nw213_[int(4)] = (self).f25
            nw213_[int(5)] = (self).f23
            nw213_[int(6)] = (self).f26
            nw213_[int(7)] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (len(d_1337_v14_))])): (self).f26}))
            nw213_[int(8)] = 353
            nw213_[int(9)] = (self).f25
            nw213_[int(10)] = (self).f25
            nw213_[int(11)] = -646
            nw213_[int(12)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rocyxmwas")))
            nw213_[int(13)] = d_1320_cf70_
            nw213_[int(14)] = len(d_1336_v13_)
            nw213_[int(15)] = 390
            nw213_[int(16)] = (self).f25
            nw213_[int(17)] = (self).f23
            nw213_[int(18)] = d_1320_cf70_
            nw213_[int(19)] = (self).f23
            nw213_[int(20)] = (self).f23
            nw213_[int(21)] = (self).f26
            nw213_[int(22)] = (self).f25
            nw213_[int(23)] = (self).f23
            nw213_[int(24)] = (self).f23
            nw213_[int(25)] = (self).f23
            nw213_[int(26)] = (self).f23
            d_1338_v15_ = nw213_
            d_1339_v16_: _dafny.Map
            d_1339_v16_ = _dafny.Map({d_1338_v15_: d_1319_cf71_})
            d_1340_v17_: _dafny.Array
            def lambda68_(d_1341_cf71_):
                def lambda69_(d_1342_i0_):
                    return d_1341_cf71_

                return lambda69_

            init39_ = lambda68_(d_1319_cf71_)
            nw214_ = _dafny.Array(None, 4)
            for i0_39_ in range(nw214_.length(0)):
                nw214_[i0_39_] = init39_(i0_39_)
            d_1340_v17_ = nw214_
            index244_ = default__.safeIndex(644, (d_1340_v17_).length(0))
            (d_1340_v17_)[index244_] = d_1319_cf71_
            index245_ = default__.safeIndex(644, (d_1340_v17_).length(0))
            rhs252_ = d_1339_v16_
            rhs253_ = ((d_1333_v10_)[d_1318_cf72_] if (d_1318_cf72_) in (d_1333_v10_) else not(((self).f25) <= ((self).f23)))
            lhs226_ = d_1340_v17_
            lhs227_ = default__.safeIndex(644, (d_1340_v17_).length(0))
            d_1339_v16_ = rhs252_
            lhs226_[lhs227_] = rhs253_
            d_1343_v18_: _dafny.Map
            d_1343_v18_ = _dafny.Map({d_1310_v3_: d_1322_cf68_})
            rhs254_ = d_1318_cf72_
            rhs255_ = _dafny.Map({d_1310_v3_: d_1318_cf72_})
            rhs256_ = d_1320_cf70_
            d_1322_cf68_ = rhs254_
            d_1343_v18_ = rhs255_
            r1 = rhs256_
            (globalState).f6 = default__.fm2(globalState)
        elif source11_.is_DC43:
            d_1344___mcc_h5_ = source11_.cf67
            d_1345_cf67_ = d_1344___mcc_h5_
            d_1346_v19_: _dafny.MultiSet
            d_1346_v19_ = _dafny.MultiSet([p0, p0])
            d_1347_v20_: _dafny.Seq
            d_1347_v20_ = _dafny.SeqWithoutIsStrInference([d_1346_v19_])
            d_1348_v21_: C3
            nw215_ = C3()
            nw215_.ctor__((_dafny.SeqWithoutIsStrInference([d_1346_v19_])) + (d_1347_v20_), (983) + ((self).f23), p0, d_1310_v3_)
            d_1348_v21_ = nw215_
            d_1349_v22_: _dafny.Set
            d_1349_v22_ = _dafny.Set({(d_1346_v19_).set(p0, default__.abs(190)), d_1346_v19_})
            def iife111_():
                coll51_ = _dafny.Set()
                compr_51_: _dafny.MultiSet
                for compr_51_ in (_dafny.Map({d_1346_v19_: p0})).keys.Elements:
                    d_1350_v23_: _dafny.MultiSet = compr_51_
                    if (d_1350_v23_) in (_dafny.Map({d_1346_v19_: p0})):
                        coll51_ = coll51_.union(_dafny.Set([d_1350_v23_]))
                return _dafny.Set(coll51_)
            if (d_1349_v22_) == (iife111_()
            ):
                d_1351_v24_: _dafny.Array
                nw216_ = _dafny.Array(None, 9)
                nw216_[int(0)] = p0
                nw216_[int(1)] = True
                nw216_[int(2)] = default__.fm1((self).f26, p0, globalState)
                nw216_[int(3)] = p0
                nw216_[int(4)] = p0
                nw216_[int(5)] = p0
                nw216_[int(6)] = default__.fm1((self).f25, p0, globalState)
                nw216_[int(7)] = p0
                nw216_[int(8)] = p0
                d_1351_v24_ = nw216_
                d_1352_v25_: _dafny.Seq
                d_1352_v25_ = _dafny.SeqWithoutIsStrInference([d_1351_v24_, d_1351_v24_, d_1351_v24_, d_1351_v24_])
                d_1352_v25_ = d_1352_v25_
                d_1353_v26_: str
                d_1353_v26_ = _dafny.CodePoint('j')
                d_1354_v27_: _dafny.Map
                d_1354_v27_ = _dafny.Map({(0) - ((self).f26): d_1309_v2_})
                d_1355_v28_: _dafny.Map
                d_1355_v28_ = _dafny.Map({p0: p0})
                d_1356_v29_: _dafny.Seq
                d_1356_v29_ = _dafny.SeqWithoutIsStrInference([d_1309_v2_])
                d_1357_v30_: _dafny.Array
                nw217_ = _dafny.Array(None, 23)
                nw217_[int(0)] = d_1309_v2_
                nw217_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfojm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmwijrbtt")))
                nw217_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1358_i1_ in range(default__.abs(172))])
                nw217_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vabaqredp"))
                nw217_[int(4)] = d_1309_v2_
                nw217_[int(5)] = (d_1309_v2_).set(default__.safeIndex((self).f25, len(d_1309_v2_)), d_1353_v26_)
                nw217_[int(6)] = ((d_1354_v27_)[len(d_1355_v28_)] if (len(d_1355_v28_)) in (d_1354_v27_) else d_1309_v2_)
                nw217_[int(7)] = ((d_1356_v29_)[default__.safeIndex(195, len(d_1356_v29_))]) + (d_1309_v2_)
                nw217_[int(8)] = d_1309_v2_
                nw217_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "untl"))
                nw217_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfldre"))) + (d_1309_v2_)
                nw217_[int(11)] = d_1309_v2_
                nw217_[int(12)] = d_1309_v2_
                nw217_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1353_v26_ for d_1359_i2_ in range(default__.abs(349))])
                nw217_[int(14)] = d_1309_v2_
                nw217_[int(15)] = d_1309_v2_
                nw217_[int(16)] = d_1309_v2_
                nw217_[int(17)] = d_1309_v2_
                nw217_[int(18)] = d_1309_v2_
                nw217_[int(19)] = (d_1309_v2_) + ((default__.fm34(500, globalState)).cf0)
                nw217_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                nw217_[int(21)] = (d_1309_v2_ if p0 else d_1309_v2_)
                nw217_[int(22)] = (d_1309_v2_) + ((d_1309_v2_).set(default__.safeIndex((self).f23, len(d_1309_v2_)), d_1353_v26_))
                d_1357_v30_ = nw217_
                index246_ = default__.safeIndex(464, (d_1357_v30_).length(0))
                (d_1357_v30_)[index246_] = d_1309_v2_
                index247_ = default__.safeIndex(464, (d_1357_v30_).length(0))
                (d_1357_v30_)[index247_] = d_1309_v2_
                (globalState).f6 = ((d_1357_v30_)[default__.safeIndex(464, (d_1357_v30_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwammv")))
                d_1360_v31_: C0
                nw218_ = C0()
                nw218_.ctor__()
                d_1360_v31_ = nw218_
                d_1361_v32_: _dafny.Seq
                d_1361_v32_ = _dafny.SeqWithoutIsStrInference([d_1360_v31_, d_1360_v31_, d_1360_v31_])
                d_1362_v33_: D17
                d_1362_v33_ = D17_DC48(d_1346_v19_)
                d_1363_v34_: _dafny.Set
                d_1363_v34_ = _dafny.Set({p0})
                d_1361_v32_ = ((d_1361_v32_) + (d_1361_v32_) if (default__.fm35(p0, False, (d_1362_v33_).cf79, globalState)).isdisjoint(d_1363_v34_) else d_1361_v32_)
                index248_ = default__.safeIndex(464, (d_1357_v30_).length(0))
                (d_1357_v30_)[index248_] = d_1309_v2_
            elif True:
                r0 = default__.fm1((self).f23, p0, globalState)
                d_1364_v35_: _dafny.Map
                d_1364_v35_ = _dafny.Map({(self).f26: (self).f26})
                d_1365_v36_: _dafny.Array
                nw219_ = _dafny.Array(int(0), 15)
                d_1365_v36_ = nw219_
                d_1366_v37_: _dafny.Seq
                d_1366_v37_ = _dafny.SeqWithoutIsStrInference([d_1365_v36_, d_1365_v36_, d_1365_v36_])
                d_1367_v38_: C1
                nw220_ = C1()
                nw220_.ctor__(p0, _dafny.SeqWithoutIsStrInference([(self).f26 for d_1368_i3_ in range(default__.abs(-756))]), default__.safeModuloInt(((d_1364_v35_)[len(d_1366_v37_)] if (len(d_1366_v37_)) in (d_1364_v35_) else (self).f26), (self).f23))
                d_1367_v38_ = nw220_
                r2 = -370
                (globalState).f9 = p0
                d_1369_v39_: _dafny.Map
                d_1369_v39_ = _dafny.Map({p0: p0})
                d_1370_v40_: _dafny.Seq
                d_1370_v40_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
                d_1371_v41_: _dafny.Seq
                d_1371_v41_ = _dafny.SeqWithoutIsStrInference([p0, (d_1370_v40_)[default__.safeIndex((_dafny.MultiSet([(self).f25])).cardinality, len(d_1370_v40_))], p0, p0])
                d_1369_v39_ = (d_1369_v39_).set(p0, (d_1370_v40_)[default__.safeIndex((self).f25, len(d_1370_v40_))])
            r2 = default__.safeDivisionInt((self).f25, (self).f25)
            d_1372_v42_: _dafny.Array
            def lambda70_(d_1373_i4_):
                return False

            init40_ = lambda70_
            nw221_ = _dafny.Array(None, 10)
            for i0_40_ in range(nw221_.length(0)):
                nw221_[i0_40_] = init40_(i0_40_)
            d_1372_v42_ = nw221_
            d_1374_v43_: D1
            d_1374_v43_ = D1_DC4(_dafny.Map({(self).f25: (self).f25}), p0, p0)
            index249_ = default__.safeIndex(604, (d_1372_v42_).length(0))
            (d_1372_v42_)[index249_] = (d_1374_v43_).cf8
            index250_ = default__.safeIndex(604, (d_1372_v42_).length(0))
            (d_1372_v42_)[index250_] = p0
        elif True:
            d_1375___mcc_h6_ = source11_.cf73
            d_1376_cf73_ = d_1375___mcc_h6_
            (globalState).f14 = ((0) - ((self).f25)) - ((self).f25)
            d_1377_v44_: _dafny.Array
            nw222_ = _dafny.Array(None, 2)
            nw222_[int(0)] = (self).f25
            nw222_[int(1)] = 79
            d_1377_v44_ = nw222_
            d_1378_v45_: _dafny.Seq
            d_1378_v45_ = _dafny.SeqWithoutIsStrInference([d_1377_v44_, d_1377_v44_, d_1377_v44_])
            d_1378_v45_ = d_1378_v45_
            d_1379_v46_: _dafny.Map
            d_1379_v46_ = _dafny.Map({(self).f23: default__.fm0((self).f23, False, globalState)})
            d_1380_v47_: C2
            nw223_ = C2()
            nw223_.ctor__(d_1379_v46_, (p0 if p0 else p0), (p0 if True else p0), (d_1310_v3_) + (d_1310_v3_), (self).f23)
            d_1380_v47_ = nw223_
            d_1381_v48_: C0
            nw224_ = C0()
            nw224_.ctor__()
            d_1381_v48_ = nw224_
        (globalState).f7 = (p0) and (p0)
        (globalState).f9 = not (p0) or (p0)
        hi15_ = (self).f23
        for d_1382_i5_ in range((self).f23, hi15_):
            d_1383_v49_: _dafny.Seq
            d_1383_v49_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1384_v50_: _dafny.Seq
            d_1384_v50_ = _dafny.SeqWithoutIsStrInference([d_1383_v49_, (_dafny.SeqWithoutIsStrInference([default__.fm1((_dafny.MultiSet(d_1310_v3_)).cardinality, p0, globalState)])).set(default__.safeIndex(default__.fm0((self).f26, p0, globalState), len(_dafny.SeqWithoutIsStrInference([default__.fm1((_dafny.MultiSet(d_1310_v3_)).cardinality, p0, globalState)]))), p0), (d_1383_v49_).set(default__.safeIndex((self).f25, len(d_1383_v49_)), p0)])
            d_1385_v51_: _dafny.Map
            d_1385_v51_ = _dafny.Map({default__.fm0((self).f23, p0, globalState): (self).f26})
            d_1386_v52_: D8
            d_1386_v52_ = D8_DC19(d_1310_v3_)
            d_1387_v53_: _dafny.Array
            nw225_ = _dafny.Array(None, 27)
            nw225_[int(0)] = d_1383_v49_
            nw225_[int(1)] = (d_1384_v50_)[default__.safeIndex((self).f25, len(d_1384_v50_))]
            nw225_[int(2)] = d_1383_v49_
            nw225_[int(3)] = d_1383_v49_
            nw225_[int(4)] = d_1383_v49_
            nw225_[int(5)] = (d_1383_v49_) + (d_1383_v49_)
            nw225_[int(6)] = d_1383_v49_
            nw225_[int(7)] = (d_1383_v49_) + (d_1383_v49_)
            nw225_[int(8)] = d_1383_v49_
            nw225_[int(9)] = (_dafny.SeqWithoutIsStrInference([True]) if p0 else d_1383_v49_)
            nw225_[int(10)] = d_1383_v49_
            nw225_[int(11)] = default__.fm30(((d_1385_v51_)[len((d_1310_v3_).set(default__.safeIndex((0) - ((self).f23), len(d_1310_v3_)), 612))] if (len((d_1310_v3_).set(default__.safeIndex((0) - ((self).f23), len(d_1310_v3_)), 612))) in (d_1385_v51_) else (self).f26), d_1386_v52_, d_1309_v2_, d_1382_i5_, globalState)
            nw225_[int(12)] = d_1383_v49_
            nw225_[int(13)] = d_1383_v49_
            nw225_[int(14)] = default__.fm30((self).f23, d_1386_v52_, d_1309_v2_, (self).f25, globalState)
            nw225_[int(15)] = d_1383_v49_
            nw225_[int(16)] = (d_1383_v49_ if True else d_1383_v49_)
            nw225_[int(17)] = d_1383_v49_
            nw225_[int(18)] = (d_1383_v49_).set(default__.safeIndex(261, len(d_1383_v49_)), p0)
            nw225_[int(19)] = d_1383_v49_
            nw225_[int(20)] = _dafny.SeqWithoutIsStrInference([p0, p0])
            nw225_[int(21)] = d_1383_v49_
            nw225_[int(22)] = (d_1383_v49_) + (d_1383_v49_)
            nw225_[int(23)] = d_1383_v49_
            nw225_[int(24)] = (default__.fm30((self).f23, d_1386_v52_, d_1309_v2_, d_1382_i5_, globalState)) + (_dafny.SeqWithoutIsStrInference([p0, True, not(p0)]))
            nw225_[int(25)] = (d_1383_v49_).set(default__.safeIndex((0) - (len(d_1309_v2_)), len(d_1383_v49_)), True)
            nw225_[int(26)] = d_1383_v49_
            d_1387_v53_ = nw225_
            d_1387_v53_ = d_1387_v53_
            rhs257_ = (self).f26
            rhs258_ = ((d_1309_v2_) + (d_1309_v2_)) + (d_1309_v2_)
            rhs259_ = p0
            lhs228_ = globalState
            lhs229_ = globalState
            lhs230_ = globalState
            lhs228_.f8 = rhs257_
            lhs229_.f6 = rhs258_
            lhs230_.f9 = rhs259_
            d_1388_v54_: _dafny.Array
            nw226_ = _dafny.Array(int(0), 25)
            d_1388_v54_ = nw226_
            d_1389_v55_: D8
            d_1389_v55_ = D8_DC20(True, d_1382_i5_)
            d_1390_v56_: D8
            d_1390_v56_ = D8_DC23(d_1389_v55_)
            d_1391_v57_: T1
            nw227_ = C2()
            nw227_.ctor__(d_1385_v51_, p0, not(p0), d_1310_v3_, (self).f25)
            d_1391_v57_ = nw227_
            d_1392_v58_: _dafny.Seq
            d_1392_v58_ = _dafny.SeqWithoutIsStrInference([d_1391_v57_])
            d_1393_v59_: _dafny.Seq
            d_1393_v59_ = _dafny.SeqWithoutIsStrInference([d_1388_v54_, d_1388_v54_])
            d_1394_v60_: _dafny.Set
            d_1394_v60_ = _dafny.Set({(self).f26, (self).f26})
            d_1395_v61_: _dafny.Map
            d_1395_v61_ = _dafny.Map({(d_1391_v57_).f28: d_1391_v57_})
            d_1396_v62_: T1
            d_1396_v62_ = d_1391_v57_
            d_1397_v63_: _dafny.Array
            nw228_ = _dafny.Array(None, 28)
            nw228_[int(0)] = d_1391_v57_
            nw228_[int(1)] = (d_1392_v58_)[default__.safeIndex(len(d_1393_v59_), len(d_1392_v58_))]
            nw228_[int(2)] = d_1391_v57_
            nw228_[int(3)] = d_1391_v57_
            nw228_[int(4)] = d_1391_v57_
            nw228_[int(5)] = d_1391_v57_
            nw228_[int(6)] = d_1391_v57_
            nw228_[int(7)] = d_1391_v57_
            nw228_[int(8)] = d_1391_v57_
            nw228_[int(9)] = (d_1392_v58_)[default__.safeIndex(d_1382_i5_, len(d_1392_v58_))]
            nw228_[int(10)] = d_1391_v57_
            nw228_[int(11)] = (d_1392_v58_)[default__.safeIndex(len(d_1394_v60_), len(d_1392_v58_))]
            nw228_[int(12)] = (d_1392_v58_)[default__.safeIndex((self).f25, len(d_1392_v58_))]
            nw228_[int(13)] = ((d_1395_v61_)[not((d_1391_v57_).f28)] if (not((d_1391_v57_).f28)) in (d_1395_v61_) else d_1391_v57_)
            nw228_[int(14)] = d_1391_v57_
            nw228_[int(15)] = (d_1396_v62_)
            nw228_[int(16)] = d_1391_v57_
            nw228_[int(17)] = d_1391_v57_
            nw228_[int(18)] = d_1391_v57_
            nw228_[int(19)] = d_1391_v57_
            nw228_[int(20)] = d_1391_v57_
            nw228_[int(21)] = (d_1396_v62_)
            nw228_[int(22)] = d_1391_v57_
            nw228_[int(23)] = d_1391_v57_
            nw228_[int(24)] = d_1391_v57_
            nw228_[int(25)] = d_1391_v57_
            nw228_[int(26)] = d_1391_v57_
            nw228_[int(27)] = d_1391_v57_
            d_1397_v63_ = nw228_
            index251_ = default__.safeIndex(638, (d_1397_v63_).length(0))
            (d_1397_v63_)[index251_] = d_1391_v57_
            index252_ = default__.safeIndex(638, (d_1397_v63_).length(0))
            rhs260_ = ((self).f25) * ((d_1391_v57_).f23)
            rhs261_ = d_1388_v54_
            rhs262_ = p0
            rhs263_ = d_1390_v56_
            rhs264_ = d_1391_v57_
            lhs231_ = globalState
            lhs232_ = globalState
            lhs233_ = d_1397_v63_
            lhs234_ = default__.safeIndex(638, (d_1397_v63_).length(0))
            lhs231_.f14 = rhs260_
            d_1388_v54_ = rhs261_
            lhs232_.f9 = rhs262_
            d_1390_v56_ = rhs263_
            lhs233_[lhs234_] = rhs264_
            (globalState).f7 = ((self).f26) == ((len(d_1383_v49_) if (d_1391_v57_).f28 else (self).f25))
        d_1398_i6_: int
        d_1398_i6_ = 0
        with _dafny.label("13"):
            while ((self).f26) > ((self).f23):
                with _dafny.c_label("13"):
                    if (d_1398_i6_) >= (100):
                        raise _dafny.Break("13")
                    d_1398_i6_ = (d_1398_i6_) + (1)
                    d_1399_v64_: _dafny.Map
                    d_1399_v64_ = _dafny.Map({p0: (self).f26})
                    d_1400_v65_: _dafny.Map
                    d_1400_v65_ = (d_1399_v64_) | (d_1399_v64_)
                    source12_ = d_1400_v65_
                    d_1401___mcc_h7_ = source12_
                    d_1402_cf31_ = d_1401___mcc_h7_
                    d_1403_v66_: _dafny.Seq
                    d_1403_v66_ = _dafny.SeqWithoutIsStrInference([True])
                    d_1404_v67_: _dafny.Array
                    nw229_ = _dafny.Array(False, 14)
                    d_1404_v67_ = nw229_
                    d_1405_v68_: _dafny.Map
                    d_1405_v68_ = _dafny.Map({d_1403_v66_: d_1404_v67_})
                    d_1406_v69_: _dafny.Seq
                    d_1406_v69_ = _dafny.SeqWithoutIsStrInference([d_1405_v68_])
                    d_1407_v70_: _dafny.Map
                    d_1407_v70_ = _dafny.Map({(d_1406_v69_)[default__.safeIndex((self).f23, len(d_1406_v69_))]: (0) - ((-861) - ((self).f26))})
                    d_1407_v70_ = (d_1407_v70_).set(d_1405_v68_, (self).f23)
                    (globalState).f21 = True
                    d_1408_v71_: _dafny.Array
                    nw230_ = _dafny.Array(_dafny.Map({}), 18)
                    d_1408_v71_ = nw230_
                    d_1409_v72_: _dafny.Map
                    d_1409_v72_ = _dafny.Map({300: not(p0)})
                    index253_ = default__.safeIndex(978, (d_1408_v71_).length(0))
                    (d_1408_v71_)[index253_] = (d_1409_v72_).set((self).f23, p0)
                    index254_ = default__.safeIndex(978, (d_1408_v71_).length(0))
                    (d_1408_v71_)[index254_] = d_1409_v72_
                    d_1410_v73_: _dafny.Map
                    d_1410_v73_ = _dafny.Map({(0) - (((self).f25) - (len(_dafny.SeqWithoutIsStrInference([(self).f25 for d_1411_i7_ in range(default__.abs(705))])))): _dafny.Map({(self).f23: not(p0)})})
                    d_1410_v73_ = (d_1410_v73_).set(len(_dafny.SeqWithoutIsStrInference([-201, 598])), _dafny.Map({(self).f25: p0}))
                    d_1412_v74_: _dafny.Seq
                    d_1412_v74_ = _dafny.SeqWithoutIsStrInference([d_1312_v5_, d_1312_v5_, d_1312_v5_])
                    d_1412_v74_ = d_1412_v74_
                    d_1413_v75_: C0
                    nw231_ = C0()
                    nw231_.ctor__()
                    d_1413_v75_ = nw231_
                    d_1414_v76_: _dafny.Map
                    d_1414_v76_ = _dafny.Map({(self).f23: d_1413_v75_})
                    d_1414_v76_ = (d_1414_v76_).set(default__.safeModuloInt((self).f23, (self).f25), d_1413_v75_)
                    (globalState).f8 = ((d_1310_v3_)[default__.safeIndex((self).f26, len(d_1310_v3_))]) + (((self).f23 if p0 else len((d_1399_v64_).set(p0, (self).f26))))
                    pass
            pass
        r0 = p0
        r1 = (self).f25
        d_1415_v77_: D0
        d_1415_v77_ = D0_DC0(d_1309_v2_)
        d_1416_v78_: D0
        d_1416_v78_ = D0_DC2(d_1415_v77_)
        pat_let_tv44_ = d_1310_v3_
        pat_let_tv45_ = d_1310_v3_
        pat_let_tv46_ = d_1310_v3_
        pat_let_tv47_ = d_1310_v3_
        def lambda71_(source13_):
            if source13_.is_DC1:
                d_1417___mcc_h8_ = source13_.cf1
                d_1418___mcc_h9_ = source13_.cf2
                d_1419___mcc_h10_ = source13_.cf3
                d_1420_cf3_ = d_1419___mcc_h10_
                d_1421_cf2_ = d_1418___mcc_h9_
                d_1422_cf1_ = d_1417___mcc_h8_
                return ((self).f26) + ((self).f23)
            elif source13_.is_DC0:
                d_1423___mcc_h11_ = source13_.cf0
                d_1424_cf0_ = d_1423___mcc_h11_
                def iife112_():
                    def iife114_():
                        coll54_ = _dafny.Set()
                        compr_54_: _dafny.Seq
                        for compr_54_ in (_dafny.SeqWithoutIsStrInference([pat_let_tv46_ for d_1425_i8_ in range(default__.abs(-298))])).Elements:
                            d_1428_v80_: _dafny.Seq = compr_54_
                            if (d_1428_v80_) in (_dafny.SeqWithoutIsStrInference([pat_let_tv47_ for d_1425_i8_ in range(default__.abs(-298))])):
                                coll54_ = coll54_.union(_dafny.Set([d_1428_v80_]))
                        return _dafny.Set(coll54_)
                    coll52_ = _dafny.Map()
                    def iife113_():
                        coll53_ = _dafny.Set()
                        compr_53_: _dafny.Seq
                        for compr_53_ in (_dafny.SeqWithoutIsStrInference([pat_let_tv44_ for d_1425_i8_ in range(default__.abs(-298))])).Elements:
                            d_1426_v80_: _dafny.Seq = compr_53_
                            if (d_1426_v80_) in (_dafny.SeqWithoutIsStrInference([pat_let_tv45_ for d_1425_i8_ in range(default__.abs(-298))])):
                                coll53_ = coll53_.union(_dafny.Set([d_1426_v80_]))
                        return _dafny.Set(coll53_)
                    compr_52_: _dafny.Seq
                    for compr_52_ in (iife113_()
                    ).Elements:
                        d_1427_v79_: _dafny.Seq = compr_52_
                        if (d_1427_v79_) in (iife114_()
                        ):
                            coll52_[d_1427_v79_] = (self).f25
                    return _dafny.Map(coll52_)
                return default__.safeModuloInt(len(iife112_()
                ), 362)
            elif True:
                d_1429___mcc_h12_ = source13_.cf4
                d_1430_cf4_ = d_1429___mcc_h12_
                return (self).f26

        r2 = lambda71_(d_1416_v78_)
        return r0, r1, r2

    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26

class C6(T0):
    def  __init__(self):
        self._f23: int = int(0)
        self._f24: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f23(self):
        return self._f23
    def ctor__(self, f24, f23):
        (self)._f24 = f24
        (self)._f23 = f23

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_1431_v0_: _dafny.Map
        d_1431_v0_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f24, (self).f23])): p0})
        hi16_ = len(d_1431_v0_)
        for d_1432_i0_ in range((self).f23, hi16_):
            r0 = (self).f24
            d_1433_v1_: _dafny.Map
            d_1433_v1_ = _dafny.Map({p1: not(p1)})
            d_1434_v2_: _dafny.Set
            d_1434_v2_ = _dafny.Set({(d_1433_v1_ if p1 else _dafny.Map({not(p1): p1})), d_1433_v1_, d_1433_v1_, d_1433_v1_})
            d_1435_v3_: str
            d_1435_v3_ = _dafny.CodePoint('j')
            d_1436_v4_: _dafny.MultiSet
            d_1436_v4_ = _dafny.MultiSet([_dafny.CodePoint('n'), d_1435_v3_])
            d_1434_v2_ = _dafny.Set({(d_1433_v1_).set(default__.fm1((d_1436_v4_).cardinality, p1, globalState), p1), (default__.fm7((self).f24, p1, p1, globalState)) | (d_1433_v1_), d_1433_v1_, (_dafny.Map({p1: p1})) | (d_1433_v1_)})
            index255_ = default__.safeIndex(897, (p0).length(0))
            (p0)[index255_] = p1
            d_1437_v5_: _dafny.MultiSet
            d_1437_v5_ = _dafny.MultiSet([p1, p1])
            d_1438_v6_: _dafny.Seq
            d_1438_v6_ = _dafny.SeqWithoutIsStrInference([p1])
            index256_ = default__.safeIndex(897, (p0).length(0))
            (p0)[index256_] = (d_1437_v5_).isdisjoint(default__.fm4(D0_DC0(_dafny.SeqWithoutIsStrInference([d_1435_v3_ for d_1439_i1_ in range(default__.abs(591))])), d_1432_i0_, d_1438_v6_, d_1435_v3_, globalState))
            d_1440_v7_: _dafny.Seq
            d_1440_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cduaj"))
            d_1441_v8_: _dafny.Seq
            d_1441_v8_ = _dafny.SeqWithoutIsStrInference([d_1440_v7_])
            (globalState).f14 = (0) - (len(d_1441_v8_))
        d_1442_v9_: _dafny.Seq
        d_1442_v9_ = _dafny.SeqWithoutIsStrInference([(self).f23])
        (globalState).f14 = (d_1442_v9_)[default__.safeIndex(439, len(d_1442_v9_))]
        index257_ = default__.safeIndex(487, (p0).length(0))
        (p0)[index257_] = (default__.fm0((self).f23, p1, globalState)) != ((self).f24)
        d_1443_v10_: _dafny.Map
        d_1443_v10_ = _dafny.Map({(self).f23: (self).f24})
        d_1444_v11_: _dafny.Seq
        d_1444_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
        d_1445_v12_: _dafny.Set
        d_1445_v12_ = _dafny.Set({(self).f24, ((d_1443_v10_)[(self).f24] if ((self).f24) in (d_1443_v10_) else len(d_1444_v11_))})
        d_1446_v13_: _dafny.MultiSet
        d_1446_v13_ = _dafny.MultiSet([default__.fm1(len(d_1445_v12_), False, globalState), p1])
        d_1447_v14_: _dafny.Set
        d_1447_v14_ = _dafny.Set({p1, p1})
        d_1448_v15_: T0
        nw232_ = C3()
        nw232_.ctor__(_dafny.SeqWithoutIsStrInference([(d_1446_v13_).set(True, default__.abs(len(d_1447_v14_))), _dafny.MultiSet([p1, p1]), d_1446_v13_, d_1446_v13_, _dafny.MultiSet([p1])]), (self).f23, p1, d_1442_v9_)
        d_1448_v15_ = nw232_
        d_1449_v16_: _dafny.Seq
        d_1449_v16_ = _dafny.SeqWithoutIsStrInference([d_1448_v15_])
        index258_ = default__.safeIndex(487, (p0).length(0))
        (p0)[index258_] = (d_1448_v15_) not in (((d_1449_v16_).set(default__.safeIndex(len(_dafny.Map({p1: p1})), len(d_1449_v16_)), d_1448_v15_)) + (d_1449_v16_))
        d_1450_v17_: _dafny.Map
        d_1450_v17_ = _dafny.Map({p1: 904})
        d_1451_v18_: _dafny.Seq
        d_1451_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(p0)[default__.safeIndex(487, (p0).length(0))]: (d_1448_v15_).f23}), d_1450_v17_])
        r0 = (((self).f23) * ((self).f24) if (d_1450_v17_) in (d_1451_v18_) else -954)
        hi17_ = default__.safeDivisionInt((self).f24, (self).f23)
        for d_1452_i2_ in range((d_1448_v15_).f23, hi17_):
            d_1453_v19_: _dafny.Map
            d_1453_v19_ = _dafny.Map({False: not(not(p1))})
            d_1453_v19_ = (d_1453_v19_).set(p1, default__.fm1((d_1448_v15_).f23, p1, globalState))
            d_1454_v20_: _dafny.Seq
            d_1454_v20_ = _dafny.SeqWithoutIsStrInference([default__.fm1((self).f23, (p0)[default__.safeIndex(487, (p0).length(0))], globalState), p1, (d_1446_v13_).isdisjoint(d_1446_v13_), p1, (not(p1) if p1 else True)])
            (globalState).f7 = (d_1454_v20_)[default__.safeIndex(len(default__.fm2(globalState)), len(d_1454_v20_))]
            d_1455_v21_: str
            d_1455_v21_ = _dafny.CodePoint('n')
            d_1456_v22_: _dafny.Map
            d_1456_v22_ = _dafny.Map({d_1455_v21_: (d_1442_v9_) + (_dafny.SeqWithoutIsStrInference([(d_1448_v15_).f23]))})
            d_1456_v22_ = (d_1456_v22_).set((d_1444_v11_)[default__.safeIndex((d_1448_v15_).f23, len(d_1444_v11_))], d_1442_v9_)
            d_1454_v20_ = (d_1454_v20_) + (d_1454_v20_)
        d_1457_v23_: _dafny.Array
        nw233_ = _dafny.Array(False, 15)
        d_1457_v23_ = nw233_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1457_v23_).length(0)):
            d_1458_i3_: int = guard_loop_4_
            if (True) and (((0) <= (d_1458_i3_)) and ((d_1458_i3_) < ((d_1457_v23_).length(0)))):
                (d_1457_v23_)[(d_1458_i3_)] = (d_1444_v11_) <= (d_1444_v11_)
        r0 = (self).f24
        d_1459_v24_: _dafny.MultiSet
        d_1459_v24_ = _dafny.MultiSet([(self).f23])
        d_1460_v25_: _dafny.Map
        d_1460_v25_ = _dafny.Map({d_1459_v24_: (d_1446_v13_).cardinality})
        d_1461_v26_: D18
        d_1461_v26_ = D18_DC50(d_1460_v25_)
        r1 = _dafny.Set({(d_1448_v15_).f23, len(((d_1461_v26_).cf84) | (d_1460_v25_)), (self).f23})
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_1462_v0_: bool
        d_1462_v0_ = True
        d_1463_v1_: _dafny.Map
        d_1463_v1_ = _dafny.Map({d_1462_v0_: (self).f23})
        d_1464_v2_: str
        d_1464_v2_ = _dafny.CodePoint('u')
        d_1465_v3_: _dafny.Set
        d_1465_v3_ = _dafny.Set({d_1464_v2_})
        d_1466_v4_: _dafny.Map
        d_1466_v4_ = _dafny.Map({d_1465_v3_: p0})
        d_1467_v5_: _dafny.Set
        d_1467_v5_ = _dafny.Set({len(((d_1466_v4_)[d_1465_v3_] if (d_1465_v3_) in (d_1466_v4_) else ((p0).set(default__.safeIndex((self).f23, len(p0)), d_1464_v2_)).set(default__.safeIndex((self).f23, len((p0).set(default__.safeIndex((self).f23, len(p0)), d_1464_v2_))), d_1464_v2_))), (self).f23})
        d_1463_v1_ = (d_1463_v1_).set((d_1467_v5_).ispropersubset(_dafny.Set({(self).f23})), (self).f24)
        d_1464_v2_ = d_1464_v2_
        if d_1462_v0_:
            d_1468_v6_: D3
            d_1468_v6_ = D3_DC10(not(d_1462_v0_))
            source14_ = d_1468_v6_
            if source14_.is_DC10:
                d_1469___mcc_h0_ = source14_.cf21
                d_1470_cf21_ = d_1469___mcc_h0_
                d_1471_v7_: _dafny.Array
                nw234_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_1471_v7_ = nw234_
                d_1472_v8_: _dafny.Array
                def lambda72_(d_1473_cf21_):
                    def lambda73_(d_1474_i0_):
                        return d_1473_cf21_

                    return lambda73_

                init41_ = lambda72_(d_1470_cf21_)
                nw235_ = _dafny.Array(None, 13)
                for i0_41_ in range(nw235_.length(0)):
                    nw235_[i0_41_] = init41_(i0_41_)
                d_1472_v8_ = nw235_
                index259_ = default__.safeIndex(540, (d_1471_v7_).length(0))
                (d_1471_v7_)[index259_] = d_1472_v8_
                d_1475_v9_: _dafny.Seq
                d_1475_v9_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1476_v10_: C3
                nw236_ = C3()
                nw236_.ctor__(default__.fm29(len(d_1475_v9_), _dafny.Set({(self).f23, (self).f23}), globalState), (self).f24, False, (d_1475_v9_) + (d_1475_v9_))
                d_1476_v10_ = nw236_
                index260_ = default__.safeIndex(540, (d_1471_v7_).length(0))
                rhs265_ = (0) - ((self).f23)
                rhs266_ = d_1464_v2_
                rhs267_ = d_1472_v8_
                rhs268_ = d_1476_v10_
                rhs269_ = d_1472_v8_
                lhs235_ = globalState
                lhs236_ = d_1471_v7_
                lhs237_ = default__.safeIndex(540, (d_1471_v7_).length(0))
                lhs235_.f14 = rhs265_
                d_1464_v2_ = rhs266_
                lhs236_[lhs237_] = rhs267_
                d_1476_v10_ = rhs268_
                d_1472_v8_ = rhs269_
                (globalState).f6 = default__.fm2(globalState)
                (globalState).f21 = d_1462_v0_
                index261_ = default__.safeIndex(682, (d_1472_v8_).length(0))
                (d_1472_v8_)[index261_] = d_1470_cf21_
                index262_ = default__.safeIndex(682, (d_1472_v8_).length(0))
                (d_1472_v8_)[index262_] = d_1470_cf21_
            elif source14_.is_DC11:
                d_1477___mcc_h1_ = source14_.cf22
                d_1478_cf22_ = d_1477___mcc_h1_
                d_1479_v11_: _dafny.Map
                d_1479_v11_ = _dafny.Map({d_1467_v5_: (p0) <= ((p0).set(default__.safeIndex((self).f23, len(p0)), d_1464_v2_))})
                d_1479_v11_ = (d_1479_v11_).set(d_1467_v5_, d_1462_v0_)
                d_1480_v12_: _dafny.MultiSet
                d_1480_v12_ = _dafny.MultiSet([d_1462_v0_])
                d_1481_v13_: _dafny.Array
                nw237_ = _dafny.Array(None, 8)
                nw237_[int(0)] = (self).f24
                nw237_[int(1)] = d_1478_cf22_
                nw237_[int(2)] = d_1478_cf22_
                nw237_[int(3)] = len(d_1463_v1_)
                nw237_[int(4)] = (self).f24
                nw237_[int(5)] = (self).f24
                nw237_[int(6)] = (d_1480_v12_).cardinality
                nw237_[int(7)] = (self).f24
                d_1481_v13_ = nw237_
                d_1482_v14_: _dafny.Map
                d_1482_v14_ = _dafny.Map({d_1481_v13_: d_1462_v0_})
                d_1483_v15_: _dafny.MultiSet
                d_1483_v15_ = _dafny.MultiSet([len(_dafny.Map({d_1462_v0_: ((d_1482_v14_)[d_1481_v13_] if (d_1481_v13_) in (d_1482_v14_) else d_1462_v0_)})), len(p0), (self).f23, d_1478_cf22_, d_1478_cf22_])
                d_1484_v16_: _dafny.Array
                nw238_ = _dafny.Array(False, 11)
                d_1484_v16_ = nw238_
                d_1485_v17_: _dafny.Map
                d_1485_v17_ = _dafny.Map({d_1483_v15_: d_1484_v16_})
                d_1486_v18_: _dafny.Map
                d_1486_v18_ = _dafny.Map({(0) - ((self).f24): d_1485_v17_})
                d_1486_v18_ = (d_1486_v18_).set((self).f24, d_1485_v17_)
                d_1487_v19_: D2
                d_1487_v19_ = D2_DC7(d_1465_v3_)
                d_1488_v20_: _dafny.Map
                d_1488_v20_ = _dafny.Map({d_1487_v19_: d_1462_v0_})
                d_1489_v21_: _dafny.Map
                d_1489_v21_ = _dafny.Map({(d_1488_v20_) | (_dafny.Map({d_1487_v19_: d_1462_v0_})): (self).f24})
                d_1489_v21_ = (d_1489_v21_).set(d_1488_v20_, (self).f24)
                d_1490_v22_: C5
                nw239_ = C5()
                nw239_.ctor__(len(p0), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))), (self).f24)
                d_1490_v22_ = nw239_
                d_1490_v22_ = d_1490_v22_
            elif source14_.is_DC12:
                d_1491___mcc_h2_ = source14_.cf23
                d_1492_cf23_ = d_1491___mcc_h2_
                d_1493_v23_: _dafny.Map
                d_1493_v23_ = _dafny.Map({(self).f23: (self).f24})
                d_1494_v24_: _dafny.Array
                nw240_ = _dafny.Array(None, 12)
                nw240_[int(0)] = (self).f23
                nw240_[int(1)] = default__.fm0((self).f24, d_1462_v0_, globalState)
                nw240_[int(2)] = (self).f24
                nw240_[int(3)] = ((self).f24 if False else (self).f24)
                nw240_[int(4)] = ((d_1493_v23_)[(self).f24] if ((self).f24) in (d_1493_v23_) else (self).f23)
                nw240_[int(5)] = (self).f24
                nw240_[int(6)] = ((self).f23) - (len(p0))
                nw240_[int(7)] = (len(d_1492_cf23_)) + ((0) - (default__.fm0((self).f23, d_1462_v0_, globalState)))
                nw240_[int(8)] = default__.fm0((0) - (len(d_1493_v23_)), d_1462_v0_, globalState)
                nw240_[int(9)] = (self).f24
                nw240_[int(10)] = (self).f24
                nw240_[int(11)] = (self).f24
                d_1494_v24_ = nw240_
                d_1494_v24_ = d_1494_v24_
                d_1495_v25_: _dafny.Seq
                d_1495_v25_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1496_v26_: _dafny.Seq
                d_1496_v26_ = _dafny.SeqWithoutIsStrInference([(p0).set(default__.safeIndex(len(d_1495_v25_), len(p0)), d_1464_v2_)])
                d_1497_v27_: D11
                d_1497_v27_ = D11_DC32((d_1496_v26_)[default__.safeIndex((self).f24, len(d_1496_v26_))])
                d_1498_v28_: _dafny.MultiSet
                d_1498_v28_ = _dafny.MultiSet([d_1497_v27_, d_1497_v27_])
                d_1499_v29_: _dafny.Map
                d_1499_v29_ = _dafny.Map({d_1498_v28_: (_dafny.SeqWithoutIsStrInference([(self).f23, (self).f24, len(d_1492_cf23_), (self).f24, (self).f24])).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference([(self).f23, (self).f24, len(d_1492_cf23_), (self).f24, (self).f24]))), default__.fm0((self).f24, d_1462_v0_, globalState))})
                d_1500_v30_: _dafny.Seq
                d_1500_v30_ = _dafny.SeqWithoutIsStrInference([(self).f23, (self).f24, (self).f23, (self).f23, (self).f24])
                d_1499_v29_ = (d_1499_v29_).set(d_1498_v28_, (d_1500_v30_) + (d_1500_v30_))
                d_1501_v31_: _dafny.Map
                d_1501_v31_ = _dafny.Map({d_1462_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxnn"))})
                d_1501_v31_ = ((d_1501_v31_) | (d_1501_v31_) if not(d_1462_v0_) else d_1501_v31_)
                (globalState).f8 = (0) - ((self).f23)
            elif source14_.is_DC9:
                d_1502___mcc_h3_ = source14_.cf20
                d_1503_cf20_ = d_1502___mcc_h3_
                (globalState).f9 = d_1462_v0_
                (globalState).f8 = (self).f23
                (globalState).f6 = p0
                d_1504_v32_: _dafny.MultiSet
                d_1504_v32_ = _dafny.MultiSet([d_1462_v0_, True, True])
                d_1505_v33_: _dafny.Seq
                d_1505_v33_ = _dafny.SeqWithoutIsStrInference([d_1462_v0_, d_1462_v0_, d_1462_v0_])
                d_1506_v34_: _dafny.Array
                nw241_ = _dafny.Array(None, 21)
                nw241_[int(0)] = d_1462_v0_
                nw241_[int(1)] = ((0) - (len(p0))) >= ((self).f23)
                nw241_[int(2)] = not(d_1462_v0_)
                nw241_[int(3)] = not(d_1462_v0_)
                nw241_[int(4)] = d_1462_v0_
                nw241_[int(5)] = not(d_1462_v0_)
                nw241_[int(6)] = default__.fm1((self).f23, d_1462_v0_, globalState)
                nw241_[int(7)] = (d_1464_v2_) in (p0)
                nw241_[int(8)] = d_1462_v0_
                nw241_[int(9)] = (d_1462_v0_) == (not(d_1462_v0_))
                nw241_[int(10)] = True
                nw241_[int(11)] = default__.fm1((self).f23, d_1462_v0_, globalState)
                nw241_[int(12)] = d_1462_v0_
                nw241_[int(13)] = (d_1504_v32_).ispropersubset(d_1504_v32_)
                nw241_[int(14)] = d_1462_v0_
                nw241_[int(15)] = d_1462_v0_
                nw241_[int(16)] = d_1462_v0_
                nw241_[int(17)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thw"))) <= (p0)
                nw241_[int(18)] = d_1462_v0_
                nw241_[int(19)] = (d_1505_v33_)[default__.safeIndex((self).f23, len(d_1505_v33_))]
                nw241_[int(20)] = d_1462_v0_
                d_1506_v34_ = nw241_
                d_1506_v34_ = d_1506_v34_
            elif True:
                d_1507___mcc_h4_ = source14_.cf24
                d_1508_cf24_ = d_1507___mcc_h4_
                d_1509_v35_: _dafny.MultiSet
                d_1509_v35_ = _dafny.MultiSet([d_1462_v0_])
                d_1510_v36_: _dafny.Set
                d_1510_v36_ = _dafny.Set({default__.fm35(d_1462_v0_, not(d_1462_v0_), d_1509_v35_, globalState)})
                d_1511_v37_: _dafny.Seq
                d_1511_v37_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1512_v38_: _dafny.Seq
                d_1512_v38_ = _dafny.SeqWithoutIsStrInference([d_1511_v37_])
                d_1513_v39_: _dafny.Map
                d_1513_v39_ = _dafny.Map({d_1462_v0_: _dafny.SeqWithoutIsStrInference([(self).f23, (self).f24])})
                d_1514_v40_: D12
                d_1514_v40_ = D12_DC36(True, d_1462_v0_, (self).f23, (self).f24, d_1513_v39_)
                d_1515_v41_: _dafny.MultiSet
                d_1515_v41_ = _dafny.MultiSet([(self).f24, (self).f24, (0) - ((self).f24), (self).f24])
                d_1516_v42_: _dafny.Array
                nw242_ = _dafny.Array(None, 17)
                nw242_[int(0)] = (self).f24
                nw242_[int(1)] = (self).f23
                nw242_[int(2)] = ((self).f23) + ((self).f23)
                nw242_[int(3)] = (self).f24
                nw242_[int(4)] = (self).f23
                nw242_[int(5)] = (self).f23
                nw242_[int(6)] = ((self).f23) * ((self).f24)
                nw242_[int(7)] = (self).f24
                nw242_[int(8)] = default__.safeDivisionInt(196, (self).f24)
                nw242_[int(9)] = (self).f24
                nw242_[int(10)] = (self).f23
                nw242_[int(11)] = len(d_1512_v38_)
                nw242_[int(12)] = (d_1514_v40_).cf59
                nw242_[int(13)] = default__.fm0(((d_1515_v41_)[(self).f24] if ((self).f24) in (d_1515_v41_) else -104), d_1462_v0_, globalState)
                nw242_[int(14)] = (self).f23
                nw242_[int(15)] = ((self).f23) * ((_dafny.MultiSet([(self).f23, (self).f24, (self).f23, len(p0)])).cardinality)
                nw242_[int(16)] = (self).f23
                d_1516_v42_ = nw242_
                index263_ = default__.safeIndex(645, (d_1516_v42_).length(0))
                (d_1516_v42_)[index263_] = (self).f24
                d_1517_v43_: _dafny.Set
                d_1517_v43_ = _dafny.Set({d_1462_v0_})
                index264_ = default__.safeIndex(645, (d_1516_v42_).length(0))
                rhs270_ = _dafny.Set({d_1517_v43_})
                rhs271_ = len(p0)
                rhs272_ = len((_dafny.Set({d_1462_v0_, d_1462_v0_})).intersection(d_1517_v43_))
                lhs238_ = d_1516_v42_
                lhs239_ = default__.safeIndex(645, (d_1516_v42_).length(0))
                d_1510_v36_ = rhs270_
                lhs238_[lhs239_] = rhs271_
                r0 = rhs272_
                d_1518_v44_: _dafny.Array
                def lambda74_(d_1519_i1_):
                    return True

                init42_ = lambda74_
                nw243_ = _dafny.Array(None, 15)
                for i0_42_ in range(nw243_.length(0)):
                    nw243_[i0_42_] = init42_(i0_42_)
                d_1518_v44_ = nw243_
                d_1520_v45_: D0
                d_1521_v46_: _dafny.Map
                d_1522_v47_: int
                out47_: D0
                out48_: _dafny.Map
                out49_: int
                out47_, out48_, out49_ = default__.m0(d_1518_v44_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqminyem")), (self).f23, globalState)
                d_1520_v45_ = out47_
                d_1521_v46_ = out48_
                d_1522_v47_ = out49_
                d_1462_v0_ = ((self).f23) != (((d_1516_v42_)[default__.safeIndex(645, (d_1516_v42_).length(0))]) * (len(_dafny.SeqWithoutIsStrInference([d_1464_v2_ for d_1523_i2_ in range(default__.abs(-520))]))))
                d_1463_v1_ = ((d_1463_v1_ if d_1462_v0_ else d_1463_v1_) if (d_1515_v41_).ispropersubset(d_1515_v41_) else d_1463_v1_)
            d_1524_v48_: _dafny.Array
            nw244_ = _dafny.Array(None, 2)
            nw244_[int(0)] = d_1462_v0_
            nw244_[int(1)] = False
            d_1524_v48_ = nw244_
            d_1524_v48_ = d_1524_v48_
            d_1525_v49_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_1525_v49_ = nw245_
            d_1526_v50_: _dafny.Map
            d_1526_v50_ = _dafny.Map({d_1462_v0_: d_1464_v2_})
            nw246_ = _dafny.Array(None, 6)
            nw246_[int(0)] = d_1464_v2_
            nw246_[int(1)] = d_1464_v2_
            nw246_[int(2)] = d_1464_v2_
            nw246_[int(3)] = ((d_1526_v50_)[False] if (False) in (d_1526_v50_) else default__.fm6((self).f23, globalState))
            nw246_[int(4)] = ((d_1526_v50_)[d_1462_v0_] if (d_1462_v0_) in (d_1526_v50_) else d_1464_v2_)
            nw246_[int(5)] = default__.fm6((self).f23, globalState)
            d_1525_v49_ = nw246_
            (globalState).f14 = 558
            rhs273_ = True
            rhs274_ = False
            lhs240_ = globalState
            lhs241_ = globalState
            lhs240_.f9 = rhs273_
            lhs241_.f7 = rhs274_
        elif True:
            (globalState).f8 = (self).f23
            d_1527_v51_: D13
            d_1527_v51_ = D13_DC39((0) - (default__.fm0((self).f24, d_1462_v0_, globalState)))
            d_1527_v51_ = (d_1527_v51_ if d_1462_v0_ else d_1527_v51_)
            d_1528_v52_: _dafny.MultiSet
            d_1528_v52_ = _dafny.MultiSet([True, d_1462_v0_])
            d_1529_v53_: T1
            nw247_ = C3()
            nw247_.ctor__(_dafny.SeqWithoutIsStrInference([d_1528_v52_, d_1528_v52_, d_1528_v52_, d_1528_v52_]), (self).f23, (d_1462_v0_) == (d_1462_v0_), _dafny.SeqWithoutIsStrInference([len(p0) for d_1530_i3_ in range(default__.abs(205))]))
            d_1529_v53_ = nw247_
            rhs275_ = d_1464_v2_
            rhs276_ = (d_1529_v53_).f23
            rhs277_ = default__.fm1((self).f23, (d_1529_v53_).f28, globalState)
            rhs278_ = (d_1529_v53_ if (d_1529_v53_).f28 else d_1529_v53_)
            lhs242_ = globalState
            lhs243_ = globalState
            lhs244_ = globalState
            lhs242_.f11 = rhs275_
            lhs243_.f14 = rhs276_
            lhs244_.f21 = rhs277_
            d_1529_v53_ = rhs278_
            (globalState).f7 = (d_1529_v53_).f28
            d_1531_v54_: _dafny.Map
            d_1531_v54_ = _dafny.Map({(d_1529_v53_).f28: p0})
            (globalState).f6 = ((_dafny.SeqWithoutIsStrInference([d_1464_v2_ for d_1532_i4_ in range(default__.abs(350))])) + (p0)) + (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1533_i5_ in range(default__.abs(635))])) + (((d_1531_v54_)[d_1462_v0_] if (d_1462_v0_) in (d_1531_v54_) else p0))).set(default__.safeIndex((d_1529_v53_).f23, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1534_i5_ in range(default__.abs(635))])) + (((d_1531_v54_)[d_1462_v0_] if (d_1462_v0_) in (d_1531_v54_) else p0)))), default__.fm6(-66, globalState)))
        r0 = ((self).f23) * ((self).f23)
        d_1535_v55_: _dafny.Array
        nw248_ = _dafny.Array(int(0), 3)
        d_1535_v55_ = nw248_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1535_v55_).length(0)):
            d_1536_i6_: int = guard_loop_5_
            if (True) and (((0) <= (d_1536_i6_)) and ((d_1536_i6_) < ((d_1535_v55_).length(0)))):
                (d_1535_v55_)[(d_1536_i6_)] = default__.safeModuloInt(d_1536_i6_, ((self).f24) * ((self).f24))
        d_1537_v56_: _dafny.Map
        d_1537_v56_ = _dafny.Map({(self).f24: d_1462_v0_})
        d_1538_v57_: _dafny.Map
        d_1538_v57_ = _dafny.Map({d_1537_v56_: (self).f24})
        (globalState).f9 = ((d_1464_v2_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))) or (((self).f23) < (len(d_1538_v57_)))
        r0 = (166) - ((self).f23)
        d_1539_v58_: _dafny.Seq
        d_1539_v58_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f23])
        r1 = d_1539_v58_
        return r0, r1

    @property
    def f24(self):
        return self._f24
