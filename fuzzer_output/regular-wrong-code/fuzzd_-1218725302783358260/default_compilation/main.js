// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, p1, globalState) {
      return _module.__default.safeModuloInt(new BigNumber(856), new BigNumber(-444));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return (false) && ((new BigNumber(((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-618)), function (_0_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }))).length)).isLessThan(new BigNumber(-520)));
    };
    static fm3(p0, p1, p2, globalState) {
      return function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-158), new BigNumber(228))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(-158)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(228)))) {
            _coll0.push([_module.__default.safeModuloInt(_1_v0, new BigNumber(-841)),!(true) || (false)]);
          }
        }
        return _coll0;
      }();
    };
    static fm9(p0, globalState) {
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.of(true)) : (_dafny.Seq.of(true))), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true, true)));
    };
    static fm10(p0, globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(968),true)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-115),true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-981),!(false)));
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.Seq.UnicodeFromString("t");
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_module.D2.create_DC7(new BigNumber(644))).dtor_cf10),new BigNumber(76));
    };
    static fm13(p0, p1, globalState) {
      return (function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(507), new BigNumber(23))) {
          let _2_v0 = _compr_1;
          if (((new BigNumber(507)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(23)))) {
            _coll1.add(_module.__default.safeModuloInt(_2_v0, new BigNumber((_dafny.Seq.of(true, false)).length)));
          }
        }
        return _coll1;
      }()).Union((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gj")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false, false)).length))).length))).Intersect(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('q'.codePointAt(0)))).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(383)), function (_3_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }),true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dyck")).length), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("r"))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-66)), function (_4_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-114)), function (_5_i2) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }))).cardinality()))).length))).length))));
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Set.fromElements(true, (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(572)), function (_6_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(535), new BigNumber(768))) {
          let _7_v0 = _compr_2;
          if (((new BigNumber(535)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(768)))) {
            _coll2.add(_module.__default.safeModuloInt(_7_v0, new BigNumber(455)));
          }
        }
        return _coll2;
      }()).length))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("sqbhaywk")).length), new BigNumber(160))));
    };
    static fm15(p0, p1, globalState) {
      return new _dafny.CodePoint('e'.codePointAt(0));
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(false, true), _dafny.Seq.of(true, true, false, false), _dafny.Seq.of(false, true), _dafny.Seq.of(false))).Difference(_dafny.MultiSet.FromArray(((!(true)) ? (_dafny.Seq.of((_module.D8.create_DC24(_dafny.Set.fromElements(!(true), true), _dafny.Seq.of(!(false)))).dtor_cf31, _dafny.Seq.of(true), _dafny.Seq.of(true, true))) : (_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(!(false)), _dafny.Seq.of(true), _dafny.Seq.of(false, true, true, false), _dafny.Seq.of(true, true))))));
    };
    static fm17(p0, p1, globalState) {
      let _source0 = _module.D6.create_DC18(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("hqs")).length))).length));
      if (_source0.is_DC16) {
        let _8___mcc_h0 = (_source0).cf22;
        let _9___mcc_h1 = (_source0).cf23;
        let _10_cf23 = _9___mcc_h1;
        let _11_cf22 = _8___mcc_h0;
        return _module.D2.create_DC8(false, new BigNumber(((_module.D12.create_DC36(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rtpervwnl")))).dtor_cf52).length), false);
      } else if (_source0.is_DC17) {
        let _12___mcc_h2 = (_source0).cf24;
        let _13_cf24 = _12___mcc_h2;
        return _module.D2.create_DC8(_13_cf24, new BigNumber(877), _13_cf24);
      } else if (_source0.is_DC18) {
        let _14___mcc_h3 = (_source0).cf25;
        let _15_cf25 = _14___mcc_h3;
        return _module.D2.create_DC8(false, _15_cf25, !(!(!(true))));
      } else if (_source0.is_DC15) {
        let _16___mcc_h4 = (_source0).cf21;
        let _17_cf21 = _16___mcc_h4;
        return _module.D2.create_DC8(false, new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.Set.fromElements(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(647)), function (_18_i0) {
    return new _dafny.CodePoint('b'.codePointAt(0));
  })).length))), _dafny.Set.fromElements(new BigNumber(-146), new BigNumber(401), new BigNumber(-293), new BigNumber((_17_cf21).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(471))).cardinality())), new BigNumber(386))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).cardinality())))).Elements) {
    let _19_v0 = _compr_3;
    if ((_dafny.Set.fromElements(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(647)), function (_18_i0) {
      return new _dafny.CodePoint('b'.codePointAt(0));
    })).length))), _dafny.Set.fromElements(new BigNumber(-146), new BigNumber(401), new BigNumber(-293), new BigNumber((_17_cf21).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(471))).cardinality())), new BigNumber(386))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).cardinality())))).contains(_19_v0)) {
      _coll3.push([_19_v0,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(708)), function (_20_i1) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)]);
    }
  }
  return _coll3;
}()).length), false);
      } else {
        let _21___mcc_h5 = (_source0).cf26;
        let _22_cf26 = _21___mcc_h5;
        return _module.D2.create_DC8(!(true), new BigNumber(-990), false);
      }
    };
    static fm18(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((!(!(false))) || (false),false);
    };
    static fm19(p0, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(true));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), ((true) ? ((_module.D9.create_DC25(new _dafny.CodePoint('k'.codePointAt(0)))).dtor_cf32) : (new _dafny.CodePoint('w'.codePointAt(0)))));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D8.create_DC23(new BigNumber(663));
      if (_source1.is_DC23) {
        let _23___mcc_h0 = (_source1).cf29;
        let _24_cf29 = _23___mcc_h0;
        return (_dafny.MultiSet.fromElements(_24_cf29, _24_cf29, new BigNumber((_dafny.Seq.of(true, false)).length))).Union(_dafny.MultiSet.fromElements(new BigNumber(240)));
      } else if (_source1.is_DC24) {
        let _25___mcc_h1 = (_source1).cf30;
        let _26___mcc_h2 = (_source1).cf31;
        let _27_cf31 = _26___mcc_h2;
        let _28_cf30 = _25___mcc_h1;
        return _dafny.MultiSet.fromElements(new BigNumber(267));
      } else {
        let _29___mcc_h3 = (_source1).cf28;
        let _30_cf28 = _29___mcc_h3;
        return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber(992))).length))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("jgetcs")).length))));
      }
    };
    static fm22(p0, globalState) {
      return _module.D8.create_DC24(_dafny.Set.fromElements(false, true), _dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(false)));
    };
    static fm23(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(((!(false)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(459)), function (_31_i0) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), function (_32_i1) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).length);
      })) : (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-19),false),!(false))).length), new BigNumber(49)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-953)), function (_33_i2) {
        return new BigNumber(-289);
      }));
    };
    static fm24(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D5.create_DC14(true), _module.D5.create_DC14(true)), _dafny.Seq.of(_module.D5.create_DC14(true), _module.D5.create_DC14(false))), _dafny.Seq.of(_module.D5.create_DC14(true), _module.D5.create_DC14(true)));
    };
    static fm25(p0, p1, p2, globalState) {
      return _module.D5.create_DC14(!(false));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), function (_34_i0) {
        return _module.D5.create_DC13(true, false);
      });
    };
    static fm27(p0, p1, p2, globalState) {
      return _module.D9.create_DC26();
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.Seq.UnicodeFromString("");
      if (p3) {
        let _35_v0;
        _35_v0 = _dafny.Seq.UnicodeFromString("gnxtuopx");
        let _36_v1;
        _36_v1 = _dafny.MultiSet.fromElements(p3);
        let _37_v2;
        let _init0 = ((_38_v1) => function (_39_i2) {
          return _module.D2.create_DC7(new BigNumber((_dafny.MultiSet.fromElements(_38_v1)).cardinality()));
        })(_36_v1);
        let _nw0 = Array((new BigNumber(22)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _37_v2 = _nw0;
        let _40_v3;
        let _nw1 = new _module.C3();
        _nw1.__ctor(_37_v2, new _dafny.CodePoint('v'.codePointAt(0)), new BigNumber(419));
        _40_v3 = _nw1;
        let _41_v4;
        _41_v4 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(668)), function (_42_i1) {
          return new BigNumber(-194);
        }),_40_v3);
        let _43_v5;
        _43_v5 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-89)), function (_44_i3) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }));
        let _45_v6;
        let _nw2 = Array((new BigNumber(13)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _35_v0;
        _nw2[(_dafny.ONE).toNumber()] = _module.__default.fm11((((_36_v1).contains(p3)) ? ((_36_v1).get(p3)) : (new BigNumber((_35_v0).length))), p2, p3, globalState);
        _nw2[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_46_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        });
        _nw2[(new BigNumber(3)).toNumber()] = _35_v0;
        _nw2[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_35_v0, _module.__default.fm11(new BigNumber((_41_v4).length), p0, p3, globalState));
        _nw2[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("exf");
        _nw2[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("ivgkwls");
        _nw2[(new BigNumber(7)).toNumber()] = (((_43_v5).contains(p3)) ? ((_43_v5).get(p3)) : (_35_v0));
        _nw2[(new BigNumber(8)).toNumber()] = _35_v0;
        _nw2[(new BigNumber(9)).toNumber()] = _35_v0;
        _nw2[(new BigNumber(10)).toNumber()] = _35_v0;
        _nw2[(new BigNumber(11)).toNumber()] = _35_v0;
        _nw2[(new BigNumber(12)).toNumber()] = _35_v0;
        _45_v6 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length));
        (_45_v6)[_index0] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(473)), function (_47_i4) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        });
        let _48_v7;
        _48_v7 = new _dafny.CodePoint('c'.codePointAt(0));
        let _index1 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length));
        let _rhs0 = (p1).multipliedBy(p0);
        let _rhs1 = _dafny.Seq.update(_35_v0, _module.__default.safeIndex(_40_v3.f12, new BigNumber((_35_v0).length)), _48_v7);
        let _lhs0 = globalState;
        let _lhs1 = _45_v6;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length));
        _lhs0.f5 = _rhs0;
        _lhs1[_lhs2] = _rhs1;
        let _49_v8;
        _49_v8 = _dafny.Seq.of(_40_v3.f12);
        let _50_v9;
        _50_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]).length),p0);
        let _51_v10;
        let _nw3 = Array((new BigNumber(13)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _49_v8;
        _nw3[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_49_v8, _49_v8);
        _nw3[(new BigNumber(2)).toNumber()] = _49_v8;
        _nw3[(new BigNumber(3)).toNumber()] = _49_v8;
        _nw3[(new BigNumber(4)).toNumber()] = _49_v8;
        _nw3[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_49_v8, _dafny.Seq.of(p1));
        _nw3[(new BigNumber(6)).toNumber()] = _49_v8;
        _nw3[(new BigNumber(7)).toNumber()] = _49_v8;
        _nw3[(new BigNumber(8)).toNumber()] = _module.__default.fm23(p3, new BigNumber(16), new BigNumber(-473), globalState);
        _nw3[(new BigNumber(9)).toNumber()] = _49_v8;
        _nw3[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(new BigNumber((_50_v9).length));
        _nw3[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_49_v8, _49_v8);
        _nw3[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(p2);
        _51_v10 = _nw3;
        let _index2 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_51_v10).length));
        (_51_v10)[_index2] = _49_v8;
        let _52_v11;
        _52_v11 = _dafny.Map.Empty.slice().updateUnsafe(p3,_49_v8);
        let _index3 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_51_v10).length));
        (_51_v10)[_index3] = (((_52_v11).contains(true)) ? ((_52_v11).get(true)) : (_49_v8));
        let _53_v12;
        _53_v12 = _module.D0.create_DC0(p3);
        let _source2 = _53_v12;
        if (_source2.is_DC1) {
          let _54___mcc_h0 = (_source2).cf1;
          let _55_cf1 = _54___mcc_h0;
          let _56_v13;
          _56_v13 = _dafny.Map.Empty.slice().updateUnsafe(_40_v3.f12,p3);
          _56_v13 = (_56_v13).update(p2, p3);
          (_40_v3).f12 = p1;
          let _57_v14;
          _57_v14 = _module.D5.create_DC14(false);
          let _58_v15;
          _58_v15 = _dafny.Seq.of(_module.D5.create_DC14(p3), _57_v14);
          let _59_v16;
          _59_v16 = _dafny.Seq.of(_58_v15, _58_v15);
          let _60_v17;
          _60_v17 = _dafny.Seq.of(_37_v2, _37_v2);
          let _61_v18;
          let _nw4 = new _module.C3();
          _nw4.__ctor((_60_v17)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(p2))).cardinality()), new BigNumber((_60_v17).length))], new _dafny.CodePoint('s'.codePointAt(0)), new BigNumber((_dafny.MultiSet.fromElements(p3, p3)).cardinality()));
          _61_v18 = _nw4;
          let _62_v19;
          _62_v19 = _dafny.MultiSet.fromElements(_61_v18, _61_v18);
          let _63_v20;
          _63_v20 = _dafny.Map.Empty.slice().updateUnsafe((_62_v19).update(_61_v18, _module.__default.abs(_40_v3.f12)),_48_v7);
          let _64_v21;
          _64_v21 = _dafny.MultiSet.fromElements(p2, _40_v3.f12);
          let _65_v22;
          _65_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat((_59_v16)[_module.__default.safeIndex(new BigNumber((_63_v20).length), new BigNumber((_59_v16).length))], _module.__default.fm24(p3, _40_v3.f12, globalState)), _module.__default.safeIndex(new BigNumber((_64_v21).cardinality()), new BigNumber((_dafny.Seq.Concat((_59_v16)[_module.__default.safeIndex(new BigNumber((_63_v20).length), new BigNumber((_59_v16).length))], _module.__default.fm24(p3, _40_v3.f12, globalState))).length)), _module.D5.create_DC14(p3)),_40_v3);
          _65_v22 = (_65_v22).update(_dafny.Seq.Concat(_58_v15, _dafny.Seq.update(_58_v15, _module.__default.safeIndex(_61_v18.f12, new BigNumber((_58_v15).length)), _57_v14)), _40_v3);
          let _66_v23;
          _66_v23 = false;
          _66_v23 = p3;
        } else if (_source2.is_DC2) {
          let _67___mcc_h1 = (_source2).cf2;
          let _68___mcc_h2 = (_source2).cf3;
          let _69___mcc_h3 = (_source2).cf4;
          let _70___mcc_h4 = (_source2).cf5;
          let _71_cf5 = _70___mcc_h4;
          let _72_cf4 = _69___mcc_h3;
          let _73_cf3 = _68___mcc_h2;
          let _74_cf2 = _67___mcc_h1;
          _49_v8 = _dafny.Seq.Concat(_dafny.Seq.of(p2, (_dafny.ZERO).minus(p0), new BigNumber(973)), _dafny.Seq.of(new BigNumber((_35_v0).length)));
          let _75_v24;
          _75_v24 = _dafny.Map.Empty.slice().updateUnsafe(_48_v7,p3);
          let _76_v25;
          let _nw5 = Array((new BigNumber(18)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = (_40_v3.f12).minus(p2);
          _nw5[(_dafny.ONE).toNumber()] = new BigNumber(((_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]).length);
          _nw5[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_40_v3.f12, p1));
          _nw5[(new BigNumber(3)).toNumber()] = p0;
          _nw5[(new BigNumber(4)).toNumber()] = _40_v3.f12;
          _nw5[(new BigNumber(5)).toNumber()] = (p0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("die")).length));
          _nw5[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_module.__default.fm0(new BigNumber(493), new BigNumber((_75_v24).length), globalState)).multipliedBy(new BigNumber(685)));
          _nw5[(new BigNumber(7)).toNumber()] = p1;
          _nw5[(new BigNumber(8)).toNumber()] = new BigNumber(((_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]).length);
          _nw5[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_40_v3.f12);
          _nw5[(new BigNumber(10)).toNumber()] = p2;
          _nw5[(new BigNumber(11)).toNumber()] = p1;
          _nw5[(new BigNumber(12)).toNumber()] = p1;
          _nw5[(new BigNumber(13)).toNumber()] = p1;
          _nw5[(new BigNumber(14)).toNumber()] = new BigNumber((_35_v0).length);
          _nw5[(new BigNumber(15)).toNumber()] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(540)), ((_77_v8) => function (_78_i5) {
            return new BigNumber((_77_v8).length);
          })(_49_v8))).length)).minus(p0);
          _nw5[(new BigNumber(16)).toNumber()] = _40_v3.f12;
          _nw5[(new BigNumber(17)).toNumber()] = _40_v3.f12;
          _76_v25 = _nw5;
          let _index4 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_76_v25).length));
          (_76_v25)[_index4] = (_40_v3.f12).multipliedBy(_40_v3.f12);
          let _79_v26;
          _79_v26 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("scckxsky"), (_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]);
          let _index5 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_76_v25).length));
          (_76_v25)[_index5] = ((_71_cf5) ? (p2) : (new BigNumber((_79_v26).length)));
          _48_v7 = _48_v7;
          _71_cf5 = !((p2).isLessThan(p0));
        } else if (_source2.is_DC0) {
          let _80___mcc_h5 = (_source2).cf0;
          let _81_cf0 = _80___mcc_h5;
          let _82_v28;
          _82_v28 = _dafny.Set.fromElements(p1);
          let _nw6 = new _module.C0();
          _nw6.__ctor(function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_module.__default.fm23(_81_cf0, _40_v3.f12, p2, globalState)).Elements) {
              let _83_v27 = _compr_4;
              if (_dafny.Seq.contains(_module.__default.fm23(_81_cf0, _40_v3.f12, p2, globalState), _83_v27)) {
                _coll4.push([(_83_v27).multipliedBy(_40_v3.f12),_81_cf0]);
              }
            }
            return _coll4;
          }(), new BigNumber((_82_v28).length));
          _40_v3 = _nw6;
          let _84_v29;
          let _nw7 = Array((new BigNumber(22)).toNumber()).fill(false);
          _84_v29 = _nw7;
          let _index6 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_84_v29).length));
          (_84_v29)[_index6] = _81_cf0;
          let _85_v30;
          _85_v30 = _module.D5.create_DC14(false);
          let _86_v31;
          _86_v31 = _dafny.Seq.of((_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]);
          let _index7 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_84_v29).length));
          let _rhs2 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("y"), _35_v0), _86_v31);
          let _rhs3 = _module.__default.fm25(_81_cf0, _81_cf0, _module.__default.safeModuloInt(_module.__default.fm0(p2, (_dafny.ZERO).minus(p2), globalState), _40_v3.f12), globalState);
          let _rhs4 = _81_cf0;
          let _lhs3 = _84_v29;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_84_v29).length));
          _lhs3[_lhs4] = _rhs2;
          _85_v30 = _rhs3;
          _81_cf0 = _rhs4;
          let _index8 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_84_v29).length));
          (_84_v29)[_index8] = _81_cf0;
          let _rhs5 = _40_v3.f12;
          let _rhs6 = (((_84_v29)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_84_v29).length))]) ? (new BigNumber(((_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]).length)) : (((_51_v10)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_51_v10).length))])[_module.__default.safeIndex(p2, new BigNumber(((_51_v10)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_51_v10).length))]).length))]));
          let _rhs7 = (_dafny.ZERO).minus(_40_v3.f12);
          let _rhs8 = _48_v7;
          let _lhs5 = globalState;
          let _lhs6 = globalState;
          let _lhs7 = globalState;
          _lhs5.f5 = _rhs5;
          _lhs6.f5 = _rhs6;
          _lhs7.f5 = _rhs7;
          _48_v7 = _rhs8;
        } else {
          let _87___mcc_h6 = (_source2).cf6;
          let _88_cf6 = _87___mcc_h6;
          let _89_v32;
          let _nw8 = Array((new BigNumber(14)).toNumber());
          _89_v32 = _nw8;
          let _90_v33;
          let _nw9 = new _module.C3();
          _nw9.__ctor(_37_v2, _48_v7, p1);
          _90_v33 = _nw9;
          let _index9 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_89_v32).length));
          (_89_v32)[_index9] = _90_v33;
          let _index10 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_89_v32).length));
          let _nw10 = new _module.C3();
          _nw10.__ctor(_37_v2, _48_v7, (_dafny.ZERO).minus(p1));
          (_89_v32)[_index10] = _nw10;
          let _91_v34;
          let _nw11 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _91_v34 = _nw11;
          let _92_v35;
          _92_v35 = _dafny.MultiSet.fromElements(_91_v34, _91_v34);
          (_40_v3).f12 = (((_92_v35).contains(_91_v34)) ? ((_92_v35).get(_91_v34)) : (new BigNumber(432)));
          let _93_v36;
          _93_v36 = _dafny.Map.Empty.slice().updateUnsafe((_40_v3.f12).multipliedBy(_40_v3.f12),p3);
          _93_v36 = (_93_v36).update((_40_v3.f12).multipliedBy(new BigNumber((_36_v1).cardinality())), p3);
          let _94_v37;
          _94_v37 = _dafny.Set.fromElements(_40_v3.f12);
          let _95_v38;
          let _nw12 = new _module.C4();
          _nw12.__ctor(new BigNumber(((_94_v37).Union(_94_v37)).length), ((_module.__default.fm1(new BigNumber(-148), p3, p0, p3, globalState)) ? (_40_v3.f12) : (_40_v3.f12)));
          _95_v38 = _nw12;
        }
        let _96_v39;
        _96_v39 = _dafny.Seq.of(p3);
        if (!_dafny.areEqual(_dafny.Seq.of(p3), _dafny.Seq.Concat(_96_v39, _96_v39))) {
          (_40_v3).f12 = (((_50_v9).contains(p0)) ? ((_50_v9).get(p0)) : (p0));
          let _97_v40;
          _97_v40 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
          let _98_v41;
          let _nw13 = Array((new BigNumber(15)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _module.__default.fm1(p0, false, (_49_v8)[_module.__default.safeIndex(new BigNumber((_36_v1).cardinality()), new BigNumber((_49_v8).length))], p3, globalState);
          _nw13[(_dafny.ONE).toNumber()] = p3;
          _nw13[(new BigNumber(2)).toNumber()] = p3;
          _nw13[(new BigNumber(3)).toNumber()] = p3;
          _nw13[(new BigNumber(4)).toNumber()] = p3;
          _nw13[(new BigNumber(5)).toNumber()] = p3;
          _nw13[(new BigNumber(6)).toNumber()] = (p3) || (!(false));
          _nw13[(new BigNumber(7)).toNumber()] = false;
          _nw13[(new BigNumber(8)).toNumber()] = p3;
          _nw13[(new BigNumber(9)).toNumber()] = !(_dafny.Seq.IsPrefixOf(_module.__default.fm23(p3, _module.__default.fm0(new BigNumber((_97_v40).length), p2, globalState), p1, globalState), (_51_v10)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_51_v10).length))]));
          _nw13[(new BigNumber(10)).toNumber()] = p3;
          _nw13[(new BigNumber(11)).toNumber()] = p3;
          _nw13[(new BigNumber(12)).toNumber()] = p3;
          _nw13[(new BigNumber(13)).toNumber()] = p3;
          _nw13[(new BigNumber(14)).toNumber()] = (_96_v39)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_96_v39).length))];
          _98_v41 = _nw13;
          let _index11 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length));
          (_98_v41)[_index11] = p3;
          let _index12 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length));
          (_98_v41)[_index12] = p3;
          let _index13 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length));
          (_98_v41)[_index13] = (_98_v41)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length))];
          let _index14 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length));
          (_98_v41)[_index14] = p3;
          (_40_v3).f12 = (((_97_v40).contains((_98_v41)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length))])) ? ((_97_v40).get((_98_v41)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_98_v41).length))])) : (p2));
        } else {
          let _99_v42;
          let _nw14 = Array((new BigNumber(14)).toNumber());
          _99_v42 = _nw14;
          let _100_v43;
          let _init1 = ((_101_p3) => function (_102_i6) {
            return _101_p3;
          })(p3);
          let _nw15 = Array((new BigNumber(4)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw15.length); _i0_1++) {
            _nw15[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _100_v43 = _nw15;
          let _103_v44;
          _103_v44 = _dafny.Map.Empty.slice().updateUnsafe(_99_v42,_100_v43);
          _103_v44 = _103_v44;
          let _104_v45;
          _104_v45 = _module.D5.create_DC13(p3, p3);
          let _105_v46;
          _105_v46 = _dafny.Seq.of(_104_v45, _104_v45, _104_v45, _104_v45);
          let _106_v47;
          _106_v47 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
          let _107_v48;
          _107_v48 = _dafny.Set.fromElements(_105_v46, _dafny.Seq.Concat(_dafny.Seq.of(_104_v45, _104_v45), _105_v46), _module.__default.fm26(p0, _module.__default.fm15(_106_v47, _40_v3.f12, globalState), p3, _48_v7, globalState));
          _107_v48 = _107_v48;
          let _108_v49;
          _108_v49 = false;
          let _109_v50;
          let _nw16 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _109_v50 = _nw16;
          let _110_v51;
          _110_v51 = _dafny.Map.Empty.slice().updateUnsafe(p1,_100_v43);
          let _index15 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_109_v50).length));
          (_109_v50)[_index15] = _110_v51;
          let _index16 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_109_v50).length));
          let _rhs9 = _module.__default.fm1(p0, _108_v49, (p0).plus(new BigNumber(-415)), p3, globalState);
          let _rhs10 = _110_v51;
          let _lhs8 = _109_v50;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_109_v50).length));
          _108_v49 = _rhs9;
          _lhs8[_lhs9] = _rhs10;
          let _111_v52;
          _111_v52 = _module.D10.create_DC29(_40_v3);
          let _112_v53;
          let _nw17 = Array((new BigNumber(26)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = _40_v3;
          _nw17[(_dafny.ONE).toNumber()] = _40_v3;
          _nw17[(new BigNumber(2)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(3)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(4)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(5)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(6)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(7)).toNumber()] = ((!(true)) ? (_40_v3) : (_40_v3));
          _nw17[(new BigNumber(8)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(9)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(10)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(11)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(12)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(13)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(14)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(15)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(16)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(17)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(18)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(19)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(20)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(21)).toNumber()] = ((!(p3)) ? (_40_v3) : (_40_v3));
          _nw17[(new BigNumber(22)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(23)).toNumber()] = _40_v3;
          _nw17[(new BigNumber(24)).toNumber()] = ((_108_v49) ? ((_111_v52).dtor_cf39) : (_40_v3));
          _nw17[(new BigNumber(25)).toNumber()] = _40_v3;
          _112_v53 = _nw17;
          let _index17 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_112_v53).length));
          (_112_v53)[_index17] = _40_v3;
          let _113_v54;
          _113_v54 = _module.D8.create_DC23(_40_v3.f12);
          let _114_v55;
          _114_v55 = _module.D9.create_DC26();
          let _115_v56;
          let _nw18 = new _module.C1();
          _nw18.__ctor(new _dafny.CodePoint('g'.codePointAt(0)), _48_v7, new BigNumber(696));
          _115_v56 = _nw18;
          let _116_v57;
          _116_v57 = _dafny.Set.fromElements(_115_v56, _115_v56);
          let _117_v58;
          _117_v58 = _dafny.Map.Empty.slice().updateUnsafe(_40_v3.f12,p3);
          let _index18 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_112_v53).length));
          let _rhs11 = _40_v3;
          let _rhs12 = ((_108_v49) ? (_40_v3) : (_40_v3));
          let _rhs13 = _48_v7;
          let _rhs14 = _module.D8.create_DC23(_module.__default.safeModuloInt(p2, new BigNumber((_116_v57).length)));
          let _rhs15 = _module.__default.fm27(_117_v58, p3, _dafny.Seq.Concat(_35_v0, (_45_v6)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_45_v6).length))]), globalState);
          let _lhs10 = _112_v53;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_112_v53).length));
          _40_v3 = _rhs11;
          _lhs10[_lhs11] = _rhs12;
          _48_v7 = _rhs13;
          _113_v54 = _rhs14;
          _114_v55 = _rhs15;
          _35_v0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_118_v56) => function (_119_i7) {
            return (_118_v56).f15;
          })(_115_v56));
        }
        let _120_v59;
        _120_v59 = _dafny.MultiSet.fromElements(p2);
        let _121_v60;
        _121_v60 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
        r0 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_40_v3.f12, new BigNumber(-35)), (((_120_v59).contains((_dafny.ZERO).minus(p1))) ? ((_120_v59).get((_dafny.ZERO).minus(p1))) : (new BigNumber((_121_v60).length))));
      } else {
        let _122_v61;
        _122_v61 = true;
        let _123_v62;
        _123_v62 = _dafny.MultiSet.fromElements(p3);
        let _124_v63;
        _124_v63 = _dafny.MultiSet.fromElements(p1, new BigNumber((_123_v62).cardinality()));
        _122_v61 = (_dafny.MultiSet.fromElements(_module.__default.fm0(new BigNumber(864), p1, globalState))).equals((_124_v63).Union(_124_v63));
        let _125_v64;
        _125_v64 = _module.D10.create_DC30(p0);
        (globalState).f5 = ((_125_v64).dtor_cf40).minus((p0).multipliedBy(p1));
        let _126_v65;
        let _nw19 = Array((new BigNumber(24)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _122_v61;
        _nw19[(_dafny.ONE).toNumber()] = _122_v61;
        _nw19[(new BigNumber(2)).toNumber()] = p3;
        _nw19[(new BigNumber(3)).toNumber()] = _122_v61;
        _nw19[(new BigNumber(4)).toNumber()] = _122_v61;
        _nw19[(new BigNumber(5)).toNumber()] = true;
        _nw19[(new BigNumber(6)).toNumber()] = false;
        _nw19[(new BigNumber(7)).toNumber()] = !(_122_v61);
        _nw19[(new BigNumber(8)).toNumber()] = !(_122_v61);
        _nw19[(new BigNumber(9)).toNumber()] = p3;
        _nw19[(new BigNumber(10)).toNumber()] = _122_v61;
        _nw19[(new BigNumber(11)).toNumber()] = _122_v61;
        _nw19[(new BigNumber(12)).toNumber()] = true;
        _nw19[(new BigNumber(13)).toNumber()] = p3;
        _nw19[(new BigNumber(14)).toNumber()] = p3;
        _nw19[(new BigNumber(15)).toNumber()] = p3;
        _nw19[(new BigNumber(16)).toNumber()] = p3;
        _nw19[(new BigNumber(17)).toNumber()] = _122_v61;
        _nw19[(new BigNumber(18)).toNumber()] = p3;
        _nw19[(new BigNumber(19)).toNumber()] = false;
        _nw19[(new BigNumber(20)).toNumber()] = p3;
        _nw19[(new BigNumber(21)).toNumber()] = _122_v61;
        _nw19[(new BigNumber(22)).toNumber()] = (_module.D5.create_DC13(_122_v61, p3)).dtor_cf19;
        _nw19[(new BigNumber(23)).toNumber()] = p3;
        _126_v65 = _nw19;
        let _127_v66;
        _127_v66 = _dafny.Map.Empty.slice().updateUnsafe(_122_v61,_126_v65);
        let _128_v67;
        _128_v67 = _module.D11.create_DC32(_127_v66);
        _127_v66 = ((_128_v67).dtor_cf42).Merge((_127_v66).Merge(_127_v66));
        (globalState).f5 = ((p3) ? (p2) : ((p2).multipliedBy(p1)));
        let _129_v68;
        let _nw20 = new _module.C4();
        _nw20.__ctor(p1, p0);
        _129_v68 = _nw20;
      }
      let _130_v69;
      _130_v69 = _dafny.Set.fromElements(p3);
      let _131_v70;
      _131_v70 = _dafny.Seq.of(p2, p0);
      let _132_v71;
      let _nw21 = Array((new BigNumber(17)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = (new BigNumber(646)).minus(p1);
      _nw21[(_dafny.ONE).toNumber()] = p2;
      _nw21[(new BigNumber(2)).toNumber()] = (p1).plus(new BigNumber((_130_v69).length));
      _nw21[(new BigNumber(3)).toNumber()] = ((false) ? (p1) : (new BigNumber((_dafny.Seq.of(p0)).length)));
      _nw21[(new BigNumber(4)).toNumber()] = p0;
      _nw21[(new BigNumber(5)).toNumber()] = p2;
      _nw21[(new BigNumber(6)).toNumber()] = p1;
      _nw21[(new BigNumber(7)).toNumber()] = (p1).plus(p1);
      _nw21[(new BigNumber(8)).toNumber()] = p1;
      _nw21[(new BigNumber(9)).toNumber()] = new BigNumber((_131_v70).length);
      _nw21[(new BigNumber(10)).toNumber()] = p1;
      _nw21[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.of(p3, p3, p3)).length);
      _nw21[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(652), p0);
      _nw21[(new BigNumber(13)).toNumber()] = p0;
      _nw21[(new BigNumber(14)).toNumber()] = p0;
      _nw21[(new BigNumber(15)).toNumber()] = p2;
      _nw21[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(p2, p2);
      _132_v71 = _nw21;
      _132_v71 = _132_v71;
      let _133_v72;
      _133_v72 = false;
      let _134_v73;
      _134_v73 = new _dafny.CodePoint('y'.codePointAt(0));
      let _135_v74;
      let _nw22 = new _module.C1();
      _nw22.__ctor(_134_v73, _134_v73, p2);
      _135_v74 = _nw22;
      let _136_v75;
      _136_v75 = _dafny.Seq.UnicodeFromString("g");
      let _rhs16 = ((_133_v72) ? (_133_v72) : ((p1).isLessThan((_dafny.ZERO).minus(new BigNumber((_136_v75).length)))));
      let _rhs17 = ((p3) ? (_133_v72) : (_133_v72));
      let _rhs18 = _133_v72;
      let _rhs19 = _135_v74;
      let _rhs20 = p2;
      let _lhs12 = globalState;
      _133_v72 = _rhs16;
      _133_v72 = _rhs17;
      _133_v72 = _rhs18;
      _135_v74 = _rhs19;
      _lhs12.f5 = _rhs20;
      let _137_v76;
      let _init2 = ((_138_v75, _139_p1, _140_v74) => function (_141_i8) {
        return (new BigNumber(-476)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_138_v75, _module.__default.safeIndex((_dafny.ZERO).minus(_139_p1), new BigNumber((_138_v75).length)), (_140_v74).f15)).length));
      })(_136_v75, p1, _135_v74);
      let _nw23 = Array((new BigNumber(21)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw23.length); _i0_2++) {
        _nw23[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _137_v76 = _nw23;
      let _index19 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_137_v76).length));
      (_137_v76)[_index19] = p3;
      let _142_v77;
      _142_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-649)), ((_143_v74) => function (_144_i9) {
        return (_143_v74).f15;
      })(_135_v74))).length),_136_v75);
      let _145_v78;
      _145_v78 = _module.D1.create_DC4(_132_v71);
      let _146_v79;
      _146_v79 = _dafny.Set.fromElements((_145_v78).dtor_cf7, _132_v71);
      let _index20 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_137_v76).length));
      let _rhs21 = _module.__default.fm1(p0, _133_v72, (_131_v70)[_module.__default.safeIndex(p0, new BigNumber((_131_v70).length))], p3, globalState);
      let _rhs22 = _module.__default.fm0(_module.__default.safeModuloInt(new BigNumber(((((_142_v77).contains(new BigNumber(-120))) ? ((_142_v77).get(new BigNumber(-120))) : (_136_v75))).length), p2), new BigNumber((((p3) ? (_146_v79) : (_146_v79))).length), globalState);
      let _lhs13 = _137_v76;
      let _lhs14 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_137_v76).length));
      _lhs13[_lhs14] = _rhs21;
      r0 = _rhs22;
      let _147_v80;
      let _nw24 = new _module.C4();
      _nw24.__ctor((new BigNumber(-263)).plus(p0), p2);
      _147_v80 = _nw24;
      _133_v72 = true;
      r0 = _module.__default.safeModuloInt(_147_v80.f11, (_147_v80).f10);
      let _148_v81;
      _148_v81 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(821));
      r1 = _148_v81;
      let _149_v82;
      _149_v82 = _dafny.Map.Empty.slice().updateUnsafe(_147_v80.f11,p0);
      let _150_v83;
      _150_v83 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(347)), function (_151_i10) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _136_v75, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_136_v75, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_136_v75).length)), _134_v73), _136_v75), _module.__default.safeIndex(new BigNumber((_149_v82).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_136_v75, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_136_v75).length)), _134_v73), _136_v75)).length)), _134_v73));
      r2 = (_150_v83)[_module.__default.safeIndex(new BigNumber(-890), new BigNumber((_150_v83).length))];
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _152_v0;
      _152_v0 = new BigNumber(172);
      let _153_v1;
      _153_v1 = _dafny.Seq.of(_152_v0);
      let _154_v2;
      _154_v2 = _dafny.Map.Empty.slice().updateUnsafe(_152_v0,_153_v1);
      let _155_v3;
      _155_v3 = _dafny.MultiSet.fromElements(_152_v0);
      let _156_v4;
      _156_v4 = new _dafny.CodePoint('s'.codePointAt(0));
      let _157_v5;
      _157_v5 = _dafny.Seq.UnicodeFromString("o");
      let _158_v6;
      let _nw25 = Array((new BigNumber(11)).toNumber());
      _nw25[(_dafny.ZERO).toNumber()] = _156_v4;
      _nw25[(_dafny.ONE).toNumber()] = _156_v4;
      _nw25[(new BigNumber(2)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(3)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(4)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(5)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(6)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(7)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(8)).toNumber()] = _156_v4;
      _nw25[(new BigNumber(9)).toNumber()] = (_157_v5)[_module.__default.safeIndex(_152_v0, new BigNumber((_157_v5).length))];
      _nw25[(new BigNumber(10)).toNumber()] = _156_v4;
      _158_v6 = _nw25;
      let _159_globalState;
      let _nw26 = new _module.GlobalState();
      _nw26.__ctor((_dafny.MultiSet.FromArray((((_154_v2).contains(new BigNumber((_153_v1).length))) ? ((_154_v2).get(new BigNumber((_153_v1).length))) : (_dafny.Seq.of(new BigNumber(138)))))).Difference(_155_v3), true, false, true, new BigNumber(841), new BigNumber(338), true, false, new BigNumber(-192), _158_v6);
      _159_globalState = _nw26;
      let _160_v7;
      _160_v7 = _dafny.Map.Empty.slice().updateUnsafe(_152_v0,new BigNumber(447));
      let _161_v8;
      _161_v8 = true;
      let _162_v9;
      let _163_v10;
      let _164_v11;
      let _out0;
      let _out1;
      let _out2;
      let _outcollector0 = _module.__default.m0((((_160_v7).contains(new BigNumber(49))) ? ((_160_v7).get(new BigNumber(49))) : (_152_v0)), _152_v0, _152_v0, _161_v8, _159_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _162_v9 = _out0;
      _163_v10 = _out1;
      _164_v11 = _out2;
      _161_v8 = _161_v8;
      let _hi0 = _152_v0;
      for (let _165_i0 = _162_v9; _165_i0.isLessThan(_hi0); _165_i0 = _165_i0.plus(_dafny.ONE)) {
        let _166_v12;
        _166_v12 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-269)).minus(_module.__default.fm0(_162_v9, _162_v9, _159_globalState)),_dafny.Seq.UnicodeFromString("ytxvk"));
        _157_v5 = (((_166_v12).contains((_dafny.ZERO).minus(_162_v9))) ? ((_166_v12).get((_dafny.ZERO).minus(_162_v9))) : (_164_v11));
        _161_v8 = _161_v8;
        let _167_v13;
        _167_v13 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("xrlcvxj"), _157_v5, _dafny.Seq.UnicodeFromString("enoyjic"));
        let _168_v14;
        _168_v14 = _dafny.Seq.of(_161_v8);
        let _169_v15;
        _169_v15 = _module.D0.create_DC0(_161_v8);
        let _170_v16;
        _170_v16 = _module.D0.create_DC0((_169_v15).dtor_cf0);
        let _rhs23 = _dafny.Seq.Concat(_167_v13, _dafny.Seq.Concat(_167_v13, _167_v13));
        let _rhs24 = !(_161_v8);
        let _rhs25 = !((_165_i0).isLessThan(new BigNumber((_168_v14).length))) || (_161_v8);
        let _rhs26 = _165_i0;
        let _rhs27 = (_169_v15).dtor_cf0;
        _167_v13 = _rhs23;
        _161_v8 = _rhs24;
        _161_v8 = _rhs25;
        _152_v0 = _rhs26;
        _161_v8 = _rhs27;
        let _171_v17;
        let _nw27 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Set.Empty);
        _171_v17 = _nw27;
        _171_v17 = _171_v17;
      }
      let _172_v18;
      let _nw28 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _172_v18 = _nw28;
      let _173_v19;
      let _nw29 = Array((new BigNumber(14)).toNumber());
      _nw29[(_dafny.ZERO).toNumber()] = _172_v18;
      _nw29[(_dafny.ONE).toNumber()] = _172_v18;
      _nw29[(new BigNumber(2)).toNumber()] = (_module.D1.create_DC4(_172_v18)).dtor_cf7;
      _nw29[(new BigNumber(3)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(4)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(5)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(6)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(7)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(8)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(9)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(10)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(11)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(12)).toNumber()] = _172_v18;
      _nw29[(new BigNumber(13)).toNumber()] = _172_v18;
      _173_v19 = _nw29;
      let _index21 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_173_v19).length));
      (_173_v19)[_index21] = _172_v18;
      let _index22 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_173_v19).length));
      (_173_v19)[_index22] = _172_v18;
      let _174_v20;
      _174_v20 = _dafny.Set.fromElements(new BigNumber(-144));
      let _175_v21;
      let _176_v22;
      let _177_v23;
      let _out3;
      let _out4;
      let _out5;
      let _outcollector1 = _module.__default.m0(_152_v0, _162_v9, (_dafny.ZERO).minus(((true) ? ((_153_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_174_v20).length)), new BigNumber((_153_v1).length))]) : (new BigNumber((_157_v5).length)))), _module.__default.fm1((_dafny.ZERO).minus(_152_v0), _161_v8, new BigNumber(-267), !(true), _159_globalState), _159_globalState);
      _out3 = _outcollector1[0];
      _out4 = _outcollector1[1];
      _out5 = _outcollector1[2];
      _175_v21 = _out3;
      _176_v22 = _out4;
      _177_v23 = _out5;
      let _178_v24;
      let _init3 = ((_179_v0) => function (_180_i1) {
        return _module.D2.create_DC7(_179_v0);
      })(_152_v0);
      let _nw30 = Array((new BigNumber(5)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw30.length); _i0_3++) {
        _nw30[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _178_v24 = _nw30;
      let _181_v25;
      let _nw31 = new _module.C3();
      _nw31.__ctor(_178_v24, new _dafny.CodePoint('x'.codePointAt(0)), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_175_v21, _175_v21)));
      _181_v25 = _nw31;
      let _182_v26;
      let _nw32 = Array((new BigNumber(23)).toNumber());
      _182_v26 = _nw32;
      let _183_v27;
      let _nw33 = new _module.C1();
      _nw33.__ctor(_156_v4, _156_v4, _162_v9);
      _183_v27 = _nw33;
      let _184_v28;
      _184_v28 = _dafny.Map.Empty.slice().updateUnsafe(_152_v0,_161_v8);
      let _185_v29;
      _185_v29 = _dafny.Map.Empty.slice().updateUnsafe(_184_v28,(_183_v27).f15);
      let _186_v30;
      _186_v30 = _module.D9.create_DC27(_183_v27, new BigNumber((_155_v3).cardinality()), new BigNumber((_157_v5).length), new BigNumber(-208), _185_v29);
      let _index23 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_182_v26).length));
      (_182_v26)[_index23] = (_186_v30).dtor_cf33;
      let _index24 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_182_v26).length));
      let _nw34 = new _module.C1();
      _nw34.__ctor(_156_v4, (_183_v27).f15, (new BigNumber((_164_v11).length)).minus(_152_v0));
      (_182_v26)[_index24] = _nw34;
      let _index25 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
      (_172_v18)[_index25] = _152_v0;
      let _index26 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
      (_172_v18)[_index26] = (_dafny.ZERO).minus((_181_v25).fm4(_159_globalState));
      let _187_v31;
      let _nw35 = Array((new BigNumber(24)).toNumber());
      _187_v31 = _nw35;
      let _index27 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_187_v31).length));
      (_187_v31)[_index27] = _181_v25;
      let _index28 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_187_v31).length));
      (_187_v31)[_index28] = ((_161_v8) ? (_181_v25) : (_181_v25));
      let _188_v32;
      let _189_v33;
      let _190_v34;
      let _191_v35;
      let _out6;
      let _out7;
      let _out8;
      let _out9;
      let _outcollector2 = (_183_v27).m2(_159_globalState);
      _out6 = _outcollector2[0];
      _out7 = _outcollector2[1];
      _out8 = _outcollector2[2];
      _out9 = _outcollector2[3];
      _188_v32 = _out6;
      _189_v33 = _out7;
      _190_v34 = _out8;
      _191_v35 = _out9;
      let _192_v36;
      _192_v36 = _module.D5.create_DC14(_161_v8);
      if ((_192_v36).dtor_cf20) {
        let _193_v37;
        let _nw36 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
        _193_v37 = _nw36;
        let _194_v38;
        let _nw37 = new _module.C2();
        _nw37.__ctor(_188_v32, (_181_v25).fm7(_159_globalState));
        _194_v38 = _nw37;
        let _195_v39;
        _195_v39 = _dafny.Map.Empty.slice().updateUnsafe(_193_v37,_194_v38);
        _195_v39 = (_195_v39).update(_193_v37, _194_v38);
        let _196_v40;
        let _nw38 = Array((new BigNumber(5)).toNumber());
        _nw38[(_dafny.ZERO).toNumber()] = !(_188_v32);
        _nw38[(_dafny.ONE).toNumber()] = !(_161_v8);
        _nw38[(new BigNumber(2)).toNumber()] = _188_v32;
        _nw38[(new BigNumber(3)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(_191_v35)).cardinality())).isLessThan(_175_v21);
        _nw38[(new BigNumber(4)).toNumber()] = _161_v8;
        _196_v40 = _nw38;
        let _index29 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length));
        (_196_v40)[_index29] = _161_v8;
        let _197_v41;
        _197_v41 = _module.D2.create_DC8((_162_v9).isLessThanOrEqualTo((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]), ((!(_188_v32)) ? ((_153_v1)[_module.__default.safeIndex(_175_v21, new BigNumber((_153_v1).length))]) : (new BigNumber((_155_v3).cardinality()))), _dafny.Seq.IsPrefixOf(_dafny.Seq.of((_183_v27).f15, new _dafny.CodePoint('n'.codePointAt(0))), _177_v23));
        let _198_v42;
        _198_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4((_173_v19)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_173_v19).length))]),_189_v33);
        let _199_v43;
        _199_v43 = _dafny.Map.Empty.slice().updateUnsafe(_188_v32,_dafny.Set.fromElements(new BigNumber(-880), _175_v21));
        let _200_v44;
        let _nw39 = Array((new BigNumber(12)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = (_dafny.Set.fromElements(_191_v35)).Difference(_174_v20);
        _nw39[(_dafny.ONE).toNumber()] = _module.__default.fm13(false, _161_v8, _159_globalState);
        _nw39[(new BigNumber(2)).toNumber()] = _174_v20;
        _nw39[(new BigNumber(3)).toNumber()] = (_dafny.Set.fromElements((_dafny.ZERO).minus(_175_v21))).Intersect(_174_v20);
        _nw39[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_191_v35);
        _nw39[(new BigNumber(5)).toNumber()] = _174_v20;
        _nw39[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(new BigNumber((_198_v42).length))).Difference((((_199_v43).contains(_188_v32)) ? ((_199_v43).get(_188_v32)) : (_174_v20)));
        _nw39[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_194_v38.f12);
        _nw39[(new BigNumber(8)).toNumber()] = (_174_v20).Union(_174_v20);
        _nw39[(new BigNumber(9)).toNumber()] = (_174_v20).Difference(_174_v20);
        _nw39[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]);
        _nw39[(new BigNumber(11)).toNumber()] = _174_v20;
        _200_v44 = _nw39;
        let _index30 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_200_v44).length));
        (_200_v44)[_index30] = _dafny.Set.fromElements((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))], (_dafny.ZERO).minus(_152_v0));
        let _index31 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length));
        let _index32 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_200_v44).length));
        let _rhs28 = _161_v8;
        let _rhs29 = (_181_v25).fm7(_159_globalState);
        let _rhs30 = function (_pat_let0_0) {
          return function (_201_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_202_dt__update_hcf13_h0) {
                return function (_pat_let2_0) {
                  return function (_203_dt__update_hcf11_h0) {
                    return _module.D2.create_DC8(_203_dt__update_hcf11_h0, (_201_dt__update__tmp_h0).dtor_cf12, _202_dt__update_hcf13_h0);
                  }(_pat_let2_0);
                }(false);
              }(_pat_let1_0);
            }(true);
          }(_pat_let0_0);
        }(_197_v41);
        let _rhs31 = _174_v20;
        let _lhs15 = _196_v40;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length));
        let _lhs17 = _200_v44;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_200_v44).length));
        _lhs15[_lhs16] = _rhs28;
        _162_v9 = _rhs29;
        _197_v41 = _rhs30;
        _lhs17[_lhs18] = _rhs31;
        let _204_v45;
        _204_v45 = _dafny.Map.Empty.slice().updateUnsafe(_161_v8,true);
        let _205_v46;
        _205_v46 = _dafny.MultiSet.fromElements(_161_v8);
        let _206_v47;
        _206_v47 = _dafny.Seq.of(_205_v46);
        let _207_v48;
        _207_v48 = _module.D5.create_DC12(_dafny.Seq.of(_dafny.MultiSet.fromElements(_188_v32, !(_188_v32))));
        let _208_v49;
        _208_v49 = _dafny.Seq.of(_module.D5.create_DC12(_206_v47), _module.D5.create_DC12(_206_v47), _207_v48, _207_v48, _module.D5.create_DC12(_206_v47));
        let _209_v50;
        _209_v50 = _module.D6.create_DC15((_183_v27).fm5((((_204_v45).contains(!(false))) ? ((_204_v45).get(!(false))) : ((_196_v40)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length))])), _161_v8, _188_v32, _module.__default.fm1(_191_v35, true, (_153_v1)[_module.__default.safeIndex(new BigNumber((_208_v49).length), new BigNumber((_153_v1).length))], _188_v32, _159_globalState), _159_globalState));
        let _pat_let_tv0 = _153_v1;
        _209_v50 = function (_pat_let3_0) {
          return function (_210_dt__update__tmp_h1) {
            return function (_pat_let4_0) {
              return function (_211_dt__update_hcf21_h0) {
                return _module.D6.create_DC15(_211_dt__update_hcf21_h0);
              }(_pat_let4_0);
            }(_pat_let_tv0);
          }(_pat_let3_0);
        }(_module.D6.create_DC15(_153_v1));
        let _212_v51;
        let _init4 = ((_213_v40, _214_v11, _215_v23) => function (_216_i2) {
          return (((_213_v40)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_213_v40).length))]) ? (_214_v11) : (_215_v23));
        })(_196_v40, _164_v11, _177_v23);
        let _nw40 = Array((new BigNumber(24)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw40.length); _i0_4++) {
          _nw40[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _212_v51 = _nw40;
        _212_v51 = _212_v51;
        let _217_v52;
        _217_v52 = _dafny.Set.fromElements(true, (_196_v40)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length))], _188_v32);
        if (_module.__default.fm1((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))], _module.__default.fm1(_162_v9, (((_204_v45).contains(_188_v32)) ? ((_204_v45).get(_188_v32)) : (_188_v32)), _194_v38.f12, _module.__default.fm1(new BigNumber(125), (_196_v40)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length))], new BigNumber((_217_v52).length), _188_v32, _159_globalState), _159_globalState), (_181_v25).fm7(_159_globalState), _188_v32, _159_globalState)) {
          let _index33 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length));
          (_196_v40)[_index33] = !(_161_v8) || ((_dafny.MultiSet.FromArray(_153_v1)).IsSubsetOf(_module.__default.fm21(_152_v0, _175_v21, _156_v4, (_dafny.ZERO).minus(_191_v35), _159_globalState)));
          _161_v8 = _188_v32;
          let _218_v53;
          let _219_v54;
          let _220_v55;
          let _221_v56;
          let _out10;
          let _out11;
          let _out12;
          let _out13;
          let _outcollector3 = (_181_v25).m2(_159_globalState);
          _out10 = _outcollector3[0];
          _out11 = _outcollector3[1];
          _out12 = _outcollector3[2];
          _out13 = _outcollector3[3];
          _218_v53 = _out10;
          _219_v54 = _out11;
          _220_v55 = _out12;
          _221_v56 = _out13;
          let _222_v57;
          _222_v57 = _dafny.Map.Empty.slice().updateUnsafe(_194_v38,_161_v8);
          let _223_v58;
          let _nw41 = new _module.C2();
          _nw41.__ctor(_161_v8, new BigNumber(((_222_v57).Merge(_222_v57)).length));
          _223_v58 = _nw41;
          _196_v40 = _196_v40;
        } else {
          (_159_globalState).f5 = (_dafny.ZERO).minus((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]);
          let _224_v59;
          _224_v59 = _dafny.Map.Empty.slice().updateUnsafe((_181_v25).fm7(_159_globalState),_157_v5);
          let _225_v60;
          _225_v60 = _dafny.Set.fromElements(_183_v27);
          _224_v59 = (_224_v59).update((new BigNumber((_225_v60).length)).multipliedBy(new BigNumber(627)), _164_v11);
          _188_v32 = !((_196_v40)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length))]) || ((_196_v40)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_196_v40).length))]);
          let _index34 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
          (_172_v18)[_index34] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_152_v0, new BigNumber(241))), _module.__default.fm0((((_160_v7).contains((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))])) ? ((_160_v7).get((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))])) : (_152_v0)), new BigNumber(658), _159_globalState));
          _177_v23 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-635)), ((_226_v11, _227_v9) => function (_228_i3) {
            return (_226_v11)[_module.__default.safeIndex(_227_v9, new BigNumber((_226_v11).length))];
          })(_164_v11, _162_v9));
        }
      } else {
        let _229_v61;
        let _nw42 = new _module.C3();
        _nw42.__ctor(_178_v24, new _dafny.CodePoint('j'.codePointAt(0)), _152_v0);
        _229_v61 = _nw42;
        let _230_v62;
        _230_v62 = _dafny.Seq.of(_188_v32);
        _230_v62 = _dafny.Seq.update(_230_v62, _module.__default.safeIndex(_175_v21, new BigNumber((_230_v62).length)), ((true) ? (_188_v32) : (_module.__default.fm1(_152_v0, _161_v8, _152_v0, _188_v32, _159_globalState))));
        let _231_v63;
        _231_v63 = _dafny.Map.Empty.slice().updateUnsafe(_189_v33,_175_v21);
        if (!((_231_v63).update(_172_v18, _175_v21)).equals(_231_v63)) {
          let _232_v64;
          _232_v64 = _dafny.Map.Empty.slice().updateUnsafe(_188_v32,_182_v26);
          let _233_v65;
          let _nw43 = Array((new BigNumber(3)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = _174_v20;
          _nw43[(_dafny.ONE).toNumber()] = _174_v20;
          _nw43[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_module.__default.fm19(_152_v0, _159_globalState)).cardinality()), (_dafny.ZERO).minus((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]), _175_v21);
          _233_v65 = _nw43;
          let _index35 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_233_v65).length));
          (_233_v65)[_index35] = _174_v20;
          let _234_v66;
          _234_v66 = _module.D2.create_DC9(_152_v0);
          let _index36 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_233_v65).length));
          let _rhs32 = (_232_v64).update((_155_v3).IsSubsetOf(_155_v3), _182_v26);
          let _rhs33 = _dafny.Set.fromElements((new BigNumber(((_229_v61).fm5(_188_v32, _188_v32, _188_v32, _161_v8, _159_globalState)).length)).plus((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]), (_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]);
          let _rhs34 = _234_v66;
          let _rhs35 = _177_v23;
          let _lhs19 = _233_v65;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_233_v65).length));
          _232_v64 = _rhs32;
          _lhs19[_lhs20] = _rhs33;
          _234_v66 = _rhs34;
          _177_v23 = _rhs35;
          let _235_v67;
          let _nw44 = new _module.C2();
          _nw44.__ctor(_161_v8, _162_v9);
          _235_v67 = _nw44;
          let _236_v68;
          let _nw45 = Array((new BigNumber(11)).toNumber());
          _236_v68 = _nw45;
          let _237_v69;
          let _nw46 = new _module.C3();
          _nw46.__ctor((_229_v61).f14, new _dafny.CodePoint('r'.codePointAt(0)), new BigNumber(268));
          _237_v69 = _nw46;
          let _index37 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_236_v68).length));
          (_236_v68)[_index37] = _237_v69;
          let _index38 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_236_v68).length));
          let _rhs36 = ((_230_v62)[_module.__default.safeIndex(_237_v69.f12, new BigNumber((_230_v62).length))]) === (_188_v32);
          let _rhs37 = ((_161_v8) ? (_237_v69) : (_237_v69));
          let _lhs21 = _236_v68;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_236_v68).length));
          _161_v8 = _rhs36;
          _lhs21[_lhs22] = _rhs37;
          let _238_v70;
          let _239_v71;
          let _240_v72;
          let _241_v73;
          let _out14;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector4 = (_181_v25).m2(_159_globalState);
          _out14 = _outcollector4[0];
          _out15 = _outcollector4[1];
          _out16 = _outcollector4[2];
          _out17 = _outcollector4[3];
          _238_v70 = _out14;
          _239_v71 = _out15;
          _240_v72 = _out16;
          _241_v73 = _out17;
          _177_v23 = _157_v5;
        } else {
          let _242_v74;
          let _nw47 = new _module.C4();
          _nw47.__ctor((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))], _162_v9);
          _242_v74 = _nw47;
          let _243_v75;
          _243_v75 = _dafny.Seq.of(_242_v74, _242_v74, _242_v74);
          let _244_v76;
          _244_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_176_v22).length),_243_v75);
          _175_v21 = _module.__default.safeDivisionInt(new BigNumber(((((_244_v76).contains((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))])) ? ((_244_v76).get((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))])) : (_243_v75))).length), _191_v35);
          let _245_v77;
          let _nw48 = new _module.C0();
          _nw48.__ctor(_184_v28, _191_v35);
          _245_v77 = _nw48;
          let _246_v78;
          _246_v78 = _dafny.Map.Empty.slice().updateUnsafe(_245_v77,_dafny.Seq.update(_157_v5, _module.__default.safeIndex(_162_v9, new BigNumber((_157_v5).length)), (_183_v27).f15));
          _191_v35 = new BigNumber((((false) ? ((_246_v78).Merge(_246_v78)) : ((_246_v78).Merge(_246_v78)))).length);
          _161_v8 = _188_v32;
          _161_v8 = false;
          let _index39 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_173_v19).length));
          let _init5 = ((_247_v74) => function (_248_i4) {
            return _module.__default.safeModuloInt(_248_i4, (_247_v74).f10);
          })(_242_v74);
          let _nw49 = Array((new BigNumber(6)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw49.length); _i0_5++) {
            _nw49[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          (_173_v19)[_index39] = _nw49;
        }
        let _249_v79;
        _249_v79 = _module.D9.create_DC26();
        _249_v79 = ((!(_188_v32)) ? (_249_v79) : (_249_v79));
        let _250_v80;
        let _251_v81;
        let _252_v82;
        let _out18;
        let _out19;
        let _out20;
        let _outcollector5 = (_181_v25).m3(_dafny.MultiSet.fromElements(false, !(_188_v32)), _159_globalState);
        _out18 = _outcollector5[0];
        _out19 = _outcollector5[1];
        _out20 = _outcollector5[2];
        _250_v80 = _out18;
        _251_v81 = _out19;
        _252_v82 = _out20;
      }
      let _253_v83;
      _253_v83 = _dafny.MultiSet.fromElements(_161_v8);
      if (_module.__default.fm1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_188_v32,_253_v83)).length), _161_v8, new BigNumber(844), (_175_v21).isLessThan(new BigNumber((_253_v83).cardinality())), _159_globalState)) {
        _191_v35 = new BigNumber(214);
        _184_v28 = (_184_v28).update(((_161_v8) ? ((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]) : ((_dafny.ZERO).minus((_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))]))), false);
        let _254_v84;
        _254_v84 = _dafny.Map.Empty.slice().updateUnsafe(_161_v8,_187_v31);
        let _255_v85;
        _255_v85 = _dafny.Map.Empty.slice().updateUnsafe((_184_v28).Merge(_module.__default.fm3(_188_v32, _161_v8, (_181_v25).fm7(_159_globalState), _159_globalState)),(((_254_v84).contains(_161_v8)) ? ((_254_v84).get(_161_v8)) : (_187_v31)));
        let _256_v87;
        _256_v87 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber(46), _188_v32, new BigNumber(959), _161_v8, _159_globalState),(_183_v27).f15);
        let _257_v88;
        _257_v88 = _module.D0.create_DC2(_161_v8, function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(601), new BigNumber(587))) {
    let _258_v86 = _compr_5;
    if (((new BigNumber(601)).isLessThanOrEqualTo(_258_v86)) && ((_258_v86).isLessThan(new BigNumber(587)))) {
      _coll5.push([_module.__default.safeDivisionInt(_258_v86, new BigNumber((_157_v5).length)),_188_v32]);
    }
  }
  return _coll5;
}(), _256_v87, _161_v8);
        let _259_v89;
        _259_v89 = _module.D0.create_DC2(_161_v8, (_257_v88).dtor_cf3, (_256_v87).update(_161_v8, (_183_v27).f15), _161_v8);
        let _260_v90;
        _260_v90 = _dafny.Map.Empty.slice().updateUnsafe(_191_v35,_187_v31);
        _255_v85 = _dafny.Map.Empty.slice().updateUnsafe((_257_v88).dtor_cf3,(((_260_v90).contains(new BigNumber(702))) ? ((_260_v90).get(new BigNumber(702))) : (_187_v31)));
        _161_v8 = _161_v8;
        let _261_v91;
        _261_v91 = _dafny.Set.fromElements(_188_v32);
        let _262_v92;
        _262_v92 = _dafny.Seq.of(_188_v32);
        let _263_v93;
        _263_v93 = _dafny.Seq.of(_262_v92, _dafny.Seq.of(_161_v8, _188_v32, _188_v32), _262_v92);
        let _264_v94;
        _264_v94 = _module.D8.create_DC24(_261_v91, (_263_v93)[_module.__default.safeIndex(_175_v21, new BigNumber((_263_v93).length))]);
        let _265_v95;
        _265_v95 = _dafny.Set.fromElements(_module.__default.fm22(_188_v32, _159_globalState), _264_v94);
        let _266_v96;
        _266_v96 = _dafny.Map.Empty.slice().updateUnsafe((_265_v95).Intersect(_265_v95),((_188_v32) ? (new BigNumber((_153_v1).length)) : (_191_v35)));
        let _index40 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_187_v31).length));
        let _rhs38 = (_266_v96).Merge(_266_v96);
        let _rhs39 = _181_v25;
        let _rhs40 = true;
        let _lhs23 = _187_v31;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_187_v31).length));
        _266_v96 = _rhs38;
        _lhs23[_lhs24] = _rhs39;
        _161_v8 = _rhs40;
      } else {
        _188_v32 = _161_v8;
        _156_v4 = _156_v4;
        _156_v4 = ((!(_188_v32)) ? (new _dafny.CodePoint('m'.codePointAt(0))) : ((_164_v11)[_module.__default.safeIndex(_175_v21, new BigNumber((_164_v11).length))]));
        (_159_globalState).f5 = _152_v0;
        (_159_globalState).f5 = new BigNumber(578);
      }
      let _267_v97;
      let _268_v98;
      let _269_v99;
      let _270_v100;
      let _out21;
      let _out22;
      let _out23;
      let _out24;
      let _outcollector6 = (_183_v27).m2(_159_globalState);
      _out21 = _outcollector6[0];
      _out22 = _outcollector6[1];
      _out23 = _outcollector6[2];
      _out24 = _outcollector6[3];
      _267_v97 = _out21;
      _268_v98 = _out22;
      _269_v99 = _out23;
      _270_v100 = _out24;
      let _271_v101;
      let _nw50 = new _module.C3();
      _nw50.__ctor((_181_v25).f14, _156_v4, new BigNumber((_dafny.Seq.UnicodeFromString("m")).length));
      _271_v101 = _nw50;
      let _272_i5;
      _272_i5 = _dafny.ZERO;
      L0: {
        while (true) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_272_i5)) {
              break L0;
            }
            _272_i5 = (_272_i5).plus(_dafny.ONE);
            (_159_globalState).f5 = _175_v21;
            if (!(_188_v32) || (_267_v97)) {
              let _273_v102;
              let _274_v103;
              let _275_v104;
              let _276_v105;
              let _out25;
              let _out26;
              let _out27;
              let _out28;
              let _outcollector7 = (_183_v27).m2(_159_globalState);
              _out25 = _outcollector7[0];
              _out26 = _outcollector7[1];
              _out27 = _outcollector7[2];
              _out28 = _outcollector7[3];
              _273_v102 = _out25;
              _274_v103 = _out26;
              _275_v104 = _out27;
              _276_v105 = _out28;
              _191_v35 = ((_267_v97) ? ((_181_v25).fm4(_159_globalState)) : ((_162_v9).minus(_162_v9)));
              let _277_v106;
              let _init6 = ((_278_v1) => function (_279_i6) {
                return _dafny.Seq.IsProperPrefixOf(_278_v1, _278_v1);
              })(_153_v1);
              let _nw51 = Array((new BigNumber(4)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw51.length); _i0_6++) {
                _nw51[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _277_v106 = _nw51;
              let _index41 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_277_v106).length));
              (_277_v106)[_index41] = _267_v97;
              let _index42 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_277_v106).length));
              let _rhs41 = (_173_v19)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_173_v19).length))];
              let _rhs42 = !(_267_v97);
              let _lhs25 = _277_v106;
              let _lhs26 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_277_v106).length));
              _274_v103 = _rhs41;
              _lhs25[_lhs26] = _rhs42;
              let _index43 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
              let _rhs43 = _157_v5;
              let _rhs44 = _191_v35;
              let _rhs45 = (_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))];
              let _rhs46 = true;
              let _lhs27 = _172_v18;
              let _lhs28 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
              _177_v23 = _rhs43;
              _175_v21 = _rhs44;
              _lhs27[_lhs28] = _rhs45;
              _267_v97 = _rhs46;
              let _280_v107;
              let _nw52 = new _module.C3();
              _nw52.__ctor((_181_v25).f14, (_183_v27).f15, (_dafny.ZERO).minus((_276_v105).minus(_191_v35)));
              _280_v107 = _nw52;
            } else {
              _164_v11 = ((_161_v8) ? (_dafny.Seq.Concat(_177_v23, _157_v5)) : (_dafny.Seq.Concat(_157_v5, _177_v23)));
              let _index44 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
              (_172_v18)[_index44] = _191_v35;
              let _281_v108;
              _281_v108 = _dafny.Set.fromElements(_156_v4);
              let _rhs47 = ((_dafny.Set.fromElements((_183_v27).f15, _156_v4)).Union(_dafny.Set.fromElements(_156_v4))).Union(_281_v108);
              let _rhs48 = !((_270_v100).minus(_152_v0)).isEqualTo((_153_v1)[_module.__default.safeIndex(_162_v9, new BigNumber((_153_v1).length))]);
              let _rhs49 = _162_v9;
              let _rhs50 = _153_v1;
              let _lhs29 = _159_globalState;
              _281_v108 = _rhs47;
              _267_v97 = _rhs48;
              _lhs29.f5 = _rhs49;
              _153_v1 = _rhs50;
              let _282_v109;
              _282_v109 = _dafny.Seq.Concat(_157_v5, _157_v5);
              _282_v109 = _282_v109;
              let _283_v110;
              let _284_v111;
              let _285_v112;
              let _286_v113;
              let _out29;
              let _out30;
              let _out31;
              let _out32;
              let _outcollector8 = (_183_v27).m2(_159_globalState);
              _out29 = _outcollector8[0];
              _out30 = _outcollector8[1];
              _out31 = _outcollector8[2];
              _out32 = _outcollector8[3];
              _283_v110 = _out29;
              _284_v111 = _out30;
              _285_v112 = _out31;
              _286_v113 = _out32;
            }
            if (_161_v8) {
              let _287_v114;
              _287_v114 = _module.D6.create_DC15(_dafny.Seq.update(_153_v1, _module.__default.safeIndex(_191_v35, new BigNumber((_153_v1).length)), _175_v21));
              _287_v114 = ((!(_161_v8)) ? (_287_v114) : (_287_v114));
              _153_v1 = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_174_v20).length)), _153_v1);
              _153_v1 = _dafny.Seq.Concat(_dafny.Seq.of(_270_v100, (_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))], _162_v9, _152_v0), _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(_152_v0))));
              let _index45 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length));
              (_172_v18)[_index45] = new BigNumber(399);
              _152_v0 = _175_v21;
            } else {
              let _288_v115;
              let _nw53 = new _module.C4();
              _nw53.__ctor((_162_v9).plus(_191_v35), _module.__default.safeModuloInt(new BigNumber((_176_v22).length), _175_v21));
              _288_v115 = _nw53;
              let _289_v116;
              let _290_v117;
              let _291_v118;
              let _out33;
              let _out34;
              let _out35;
              let _outcollector9 = (_271_v101).m4(_162_v9, _161_v8, _dafny.Seq.Concat(_157_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-669)), ((_292_v27) => function (_293_i7) {
                return (_292_v27).f15;
              })(_183_v27))), _159_globalState);
              _out33 = _outcollector9[0];
              _out34 = _outcollector9[1];
              _out35 = _outcollector9[2];
              _289_v116 = _out33;
              _290_v117 = _out34;
              _291_v118 = _out35;
              let _294_v119;
              let _nw54 = Array((new BigNumber(16)).toNumber()).fill(false);
              _294_v119 = _nw54;
              let _index46 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_294_v119).length));
              (_294_v119)[_index46] = !(new BigNumber((_174_v20).length)).isEqualTo((_288_v115).f10);
              let _295_v120;
              _295_v120 = _dafny.Map.Empty.slice().updateUnsafe(_152_v0,_dafny.Set.fromElements(_175_v21));
              let _index47 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_294_v119).length));
              (_294_v119)[_index47] = (_174_v20).IsSubsetOf((((_295_v120).contains((_dafny.ZERO).minus(new BigNumber((_177_v23).length)))) ? ((_295_v120).get((_dafny.ZERO).minus(new BigNumber((_177_v23).length)))) : (_174_v20)));
              let _296_v121;
              let _297_v122;
              let _298_v123;
              let _out36;
              let _out37;
              let _out38;
              let _outcollector10 = (_271_v101).m3(_253_v83, _159_globalState);
              _out36 = _outcollector10[0];
              _out37 = _outcollector10[1];
              _out38 = _outcollector10[2];
              _296_v121 = _out36;
              _297_v122 = _out37;
              _298_v123 = _out38;
              let _index48 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_294_v119).length));
              (_294_v119)[_index48] = false;
            }
            _152_v0 = (_172_v18)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_172_v18).length))];
          }
        }
      }
      let _rhs51 = (_dafny.ZERO).minus(_162_v9);
      let _rhs52 = _module.__default.fm1(new BigNumber((_163_v10).length), _267_v97, _270_v100, _267_v97, _159_globalState);
      _191_v35 = _rhs51;
      _188_v32 = _rhs52;
      process.stdout.write(_dafny.toString(_152_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_153_v1, _dafny.Seq.of(new BigNumber(6), new BigNumber(399), new BigNumber(583), new BigNumber(399), new BigNumber(399)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(172),_dafny.Seq.of(new BigNumber(172))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v3).equals(_dafny.MultiSet.fromElements(new BigNumber(172)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_156_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_157_v5, _dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f0).equals(_dafny.MultiSet.fromElements(new BigNumber(138)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(172),new BigNumber(447)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(821)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_164_v11, _dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v18)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(8)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(9)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(10)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(11)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(12)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v19)[new BigNumber(13)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v20).equals(_dafny.Set.fromElements(new BigNumber(-144)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_v21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v22).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(821)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_177_v23, _dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v24)[_dafny.ZERO]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v24)[_dafny.ONE]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v24)[new BigNumber(2)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v24)[new BigNumber(3)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v24)[new BigNumber(4)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_181_v25).f14)[_dafny.ZERO]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_181_v25).f14)[_dafny.ONE]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_181_v25).f14)[new BigNumber(2)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_181_v25).f14)[new BigNumber(3)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_181_v25).f14)[new BigNumber(4)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v26)[new BigNumber(11)]).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v26)[new BigNumber(11)]).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v26)[new BigNumber(11)].f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_183_v27).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v28).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(171),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_v29).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(171),true),new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_186_v30).dtor_cf33).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_186_v30).dtor_cf33).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v30).dtor_cf33.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v30).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v30).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v30).dtor_cf36));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_186_v30).dtor_cf37).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(171),true),new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_187_v31)[new BigNumber(4)]).f14)[_dafny.ZERO]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_187_v31)[new BigNumber(4)]).f14)[_dafny.ONE]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_187_v31)[new BigNumber(4)]).f14)[new BigNumber(2)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_187_v31)[new BigNumber(4)]).f14)[new BigNumber(3)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_187_v31)[new BigNumber(4)]).f14)[new BigNumber(4)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_187_v31)[new BigNumber(4)]).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v31)[new BigNumber(4)].f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_188_v32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v33)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v34).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_191_v35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v36).dtor_cf20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v83).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_267_v97));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v98)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v99).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_270_v100));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_271_v101).f14)[_dafny.ZERO]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_271_v101).f14)[_dafny.ONE]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_271_v101).f14)[new BigNumber(2)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_271_v101).f14)[new BigNumber(3)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_271_v101).f14)[new BigNumber(4)]).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_272_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC2(cf2, cf3, cf4, cf5) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D0(3);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1([]);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC5(cf8) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC6(cf9) {
      let $dt = new D1(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC8(cf11, cf12, cf13) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D2(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC7(cf10) {
      let $dt = new D2(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(false, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf15) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + this.cf15.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf16) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf18, cf19) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC14(cf20) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC12(cf17) {
      let $dt = new D5(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf22, cf23) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC17(cf24) {
      let $dt = new D6(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC18(cf25) {
      let $dt = new D6(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC15(cf21) {
      let $dt = new D6(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC19(cf26) {
      let $dt = new D6(4);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get is_DC19() { return this.$tag === 4; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16([], _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC20(cf27) {
      let $dt = new D7(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf29) {
      let $dt = new D8(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC24(cf30, cf31) {
      let $dt = new D8(1);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC22(cf28) {
      let $dt = new D8(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC27(cf33, cf34, cf35, cf36, cf37) {
      let $dt = new D9(1);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC25(cf32) {
      let $dt = new D9(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC28(cf38) {
      let $dt = new D9(3);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26";
      } else if (this.$tag === 1) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf40) {
      let $dt = new D10(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC29(cf39) {
      let $dt = new D10(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC31(cf41) {
      let $dt = new D10(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC31" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC30(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC33(cf43, cf44, cf45, cf46, cf47) {
      let $dt = new D11(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC34(cf48, cf49, cf50) {
      let $dt = new D11(1);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC32(cf42) {
      let $dt = new D11(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC35(cf51) {
      let $dt = new D11(3);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get is_DC35() { return this.$tag === 3; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC34" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC35" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44) && this.cf45 === other.cf45 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48) && _dafny.areEqual(this.cf49, other.cf49) && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC33(_dafny.Map.Empty, _dafny.ZERO, false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37() {
      let $dt = new D12(0);
      return $dt;
    }
    static create_DC38(cf53, cf54, cf55) {
      let $dt = new D12(1);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC39(cf56) {
      let $dt = new D12(2);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC36(cf52) {
      let $dt = new D12(3);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get is_DC36() { return this.$tag === 3; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC37";
      } else if (this.$tag === 1) {
        return "D12.DC38" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC39" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54) && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf56 === other.cf56;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC37();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f5 = _dafny.ZERO;
      this._f0 = _dafny.MultiSet.Empty;
      this._f1 = false;
      this._f2 = false;
      this._f3 = false;
      this._f4 = _dafny.ZERO;
      this._f6 = false;
      this._f7 = false;
      this._f8 = _dafny.ZERO;
      this._f9 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = _dafny.ZERO;
      this.f16 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f16, f12) {
      let _this = this;
      (_this).f16 = f16;
      (_this)._f12 = f12;
      return;
    }
    fm6(globalState) {
      let _this = this;
      return _module.D0.create_DC0(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("qxlwa"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), function (_299_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})));
    };
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lnadqpvyl"), ((false) ? (_dafny.Seq.UnicodeFromString("yutm")) : (_dafny.Seq.UnicodeFromString("slatllus"))));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f12 = _dafny.ZERO;
      this._f13 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f15 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f15, f13, f12) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f13 = f13;
      (_this)._f12 = f12;
      return;
    }
    fm4(globalState) {
      let _this = this;
      return _this.f12;
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (false) {
        return _dafny.Seq.of(_this.f12);
      } else if (true) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(155)), function (_300_i0) {
          return _this.f12;
        });
      } else {
        return _dafny.Seq.of(new BigNumber(234), _this.f12);
      }
    };
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _301_v0;
      let _init7 = function (_302_i0) {
        return (_302_i0).multipliedBy(_this.f12);
      };
      let _nw55 = Array((new BigNumber(12)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw55.length); _i0_7++) {
        _nw55[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _301_v0 = _nw55;
      let _index49 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length));
      (_301_v0)[_index49] = _this.f12;
      let _index50 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length));
      (_301_v0)[_index50] = _this.f12;
      let _303_v1;
      _303_v1 = false;
      let _304_v2;
      _304_v2 = _module.D1.create_DC5(_303_v1);
      _304_v2 = _module.D1.create_DC5(_303_v1);
      if (((_303_v1) ? (_303_v1) : (_303_v1))) {
        let _305_v3;
        _305_v3 = _dafny.Set.fromElements(_303_v1, _303_v1);
        let _306_v4;
        _306_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_305_v3).length),false);
        let _307_v5;
        let _nw56 = new _module.C0();
        _nw56.__ctor(_306_v4, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
        _307_v5 = _nw56;
        let _308_v6;
        _308_v6 = _dafny.Seq.of(_303_v1, _303_v1, _303_v1);
        let _309_v7;
        _309_v7 = _dafny.Seq.of(_301_v0);
        let _rhs53 = true;
        let _rhs54 = !(_303_v1) || (_dafny.Seq.contains(_dafny.Seq.of(_307_v5), _307_v5));
        let _rhs55 = !(_303_v1) || (!((_308_v6)[_module.__default.safeIndex((_dafny.ZERO).minus(_307_v5.f12), new BigNumber((_308_v6).length))]));
        let _rhs56 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_309_v7, _309_v7), _309_v7), _module.__default.safeIndex(_this.f12, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_309_v7, _309_v7), _309_v7)).length)), _301_v0)).length);
        let _lhs30 = globalState;
        r0 = _rhs53;
        r0 = _rhs54;
        r0 = _rhs55;
        _lhs30.f5 = _rhs56;
        let _310_v8;
        _310_v8 = _dafny.Seq.of(_this.f12);
        let _311_v9;
        _311_v9 = _dafny.MultiSet.fromElements(_307_v5.f12);
        let _312_v10;
        _312_v10 = _dafny.Map.Empty.slice().updateUnsafe(_303_v1,_311_v9);
        let _313_v11;
        _313_v11 = _dafny.Map.Empty.slice().updateUnsafe(_312_v10,_303_v1);
        let _314_v12;
        _314_v12 = _dafny.Map.Empty.slice().updateUnsafe(_303_v1,(_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
        let _315_v13;
        _315_v13 = _module.D2.create_DC7(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_311_v9).contains(_this.f12)) ? ((_311_v9).get(_this.f12)) : ((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])),(((_314_v12).contains(_303_v1)) ? ((_314_v12).get(_303_v1)) : (_this.f12)))).length));
        let _rhs57 = _303_v1;
        let _rhs58 = _307_v5.f12;
        let _rhs59 = (((_313_v11).contains(_312_v10)) ? ((_313_v11).get(_312_v10)) : (!((_303_v1) || (_303_v1))));
        let _rhs60 = _dafny.Seq.update(_310_v8, _module.__default.safeIndex(((_303_v1) ? ((_315_v13).dtor_cf10) : (new BigNumber(-345))), new BigNumber((_310_v8).length)), _307_v5.f12);
        let _lhs31 = globalState;
        _303_v1 = _rhs57;
        _lhs31.f5 = _rhs58;
        r0 = _rhs59;
        _310_v8 = _rhs60;
        if (false) {
          let _rhs61 = ((_303_v1) ? (_303_v1) : (!_dafny.Seq.contains(_308_v6, _303_v1)));
          let _rhs62 = _303_v1;
          _303_v1 = _rhs61;
          r0 = _rhs62;
          let _316_v14;
          let _nw57 = new _module.C0();
          _nw57.__ctor(_306_v4, (_module.D2.create_DC9(new BigNumber(-834))).dtor_cf14);
          _316_v14 = _nw57;
          let _317_v15;
          _317_v15 = new _dafny.CodePoint('n'.codePointAt(0));
          _317_v15 = (_this).f13;
          let _318_v16;
          _318_v16 = _dafny.Seq.UnicodeFromString("iiotodfiq");
          _318_v16 = _dafny.Seq.Concat(_dafny.Seq.Concat(_318_v16, _318_v16), _318_v16);
          _308_v6 = _dafny.Seq.Concat(_module.__default.fm9(!(_303_v1), globalState), _dafny.Seq.of(_303_v1, true, _303_v1));
        } else {
          let _319_v17;
          _319_v17 = _dafny.Seq.UnicodeFromString("q");
          let _320_v18;
          _320_v18 = _dafny.Seq.of(_319_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(718)), function (_321_i1) {
            return (_this).f15;
          }), _dafny.Seq.Concat(_319_v17, _319_v17));
          let _322_v19;
          _322_v19 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_307_v5.f12),new BigNumber(210));
          _319_v17 = (_320_v18)[_module.__default.safeIndex((new BigNumber(440)).plus((((_322_v19).contains(new BigNumber((_319_v17).length))) ? ((_322_v19).get(new BigNumber((_319_v17).length))) : ((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]))), new BigNumber((_320_v18).length))];
          let _323_v20;
          let _nw58 = new _module.C0();
          _nw58.__ctor(_306_v4, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
          _323_v20 = _nw58;
          _323_v20 = _323_v20;
          let _324_v21;
          let _nw59 = Array((new BigNumber(22)).toNumber());
          _324_v21 = _nw59;
          _324_v21 = _324_v21;
          let _325_v22;
          _325_v22 = _dafny.MultiSet.fromElements(_308_v6);
          let _326_v23;
          _326_v23 = _dafny.Seq.of(_308_v6);
          _325_v22 = _dafny.MultiSet.FromArray(_326_v23);
          (_323_v20).f16 = _323_v20.f16;
        }
        _303_v1 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("h"), _dafny.Seq.UnicodeFromString("lbyqqyufh"));
        let _327_v24;
        let _nw60 = new _module.C0();
        _nw60.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_this.f12,_303_v1), (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
        _327_v24 = _nw60;
        let _328_v25;
        _328_v25 = _dafny.MultiSet.fromElements(_327_v24, _327_v24, _327_v24);
        let _329_v26;
        _329_v26 = _dafny.Map.Empty.slice().updateUnsafe(_303_v1,(_this).f15);
        let _330_v27;
        _330_v27 = _module.D0.create_DC2(_303_v1, _module.__default.fm10(new BigNumber((_328_v25).cardinality()), globalState), _329_v26, true);
        let _331_v28;
        _331_v28 = _dafny.Map.Empty.slice().updateUnsafe(_330_v27,_this.f12);
        (globalState).f5 = _module.__default.safeDivisionInt(((true) ? ((((_331_v28).contains(_330_v27)) ? ((_331_v28).get(_330_v27)) : ((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]))) : (_307_v5.f12)), (_307_v5.f12).minus((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]));
      } else {
        if (((_303_v1) ? (_303_v1) : (false))) {
          let _332_v29;
          _332_v29 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),_303_v1);
          let _333_v30;
          _333_v30 = _dafny.Seq.UnicodeFromString("ynbd");
          _303_v1 = !((((_332_v29).contains(new _dafny.CodePoint('w'.codePointAt(0)))) ? ((_332_v29).get(new _dafny.CodePoint('w'.codePointAt(0)))) : (_303_v1))) || (_module.__default.fm1(new BigNumber((_333_v30).length), _module.__default.fm1(_this.f12, !(_303_v1), (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _303_v1, globalState), (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _303_v1, globalState));
          let _334_v31;
          _334_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-503),_303_v1);
          let _335_v32;
          let _nw61 = new _module.C0();
          _nw61.__ctor((_334_v31).Merge(_334_v31), (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
          _335_v32 = _nw61;
          _333_v30 = _333_v30;
          r0 = _303_v1;
          r0 = _module.__default.fm1(new BigNumber((_dafny.Seq.Concat(_333_v30, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(198)), function (_336_i2) {
            return (_this).f13;
          }), _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_this.f12)).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(198)), function (_337_i2) {
            return (_this).f13;
          })).length)), new _dafny.CodePoint('q'.codePointAt(0))))).length), (_this.f12).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length))), (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], (_303_v1) && (_303_v1), globalState);
        } else {
          let _338_v33;
          _338_v33 = _dafny.Seq.of(_303_v1, _303_v1, _303_v1);
          let _339_v34;
          _339_v34 = _dafny.Seq.UnicodeFromString("cqtlrd");
          let _340_v35;
          _340_v35 = _dafny.Seq.of(_339_v34);
          let _341_v36;
          _341_v36 = _module.D0.create_DC0(true);
          let _342_v37;
          _342_v37 = _module.D0.create_DC3(_341_v36);
          let _343_v38;
          let _nw62 = Array((new BigNumber(23)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _303_v1;
          _nw62[(_dafny.ONE).toNumber()] = _303_v1;
          _nw62[(new BigNumber(2)).toNumber()] = false;
          _nw62[(new BigNumber(3)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(4)).toNumber()] = _module.__default.fm1(_this.f12, _303_v1, new BigNumber(-989), _303_v1, globalState);
          _nw62[(new BigNumber(5)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(6)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(7)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(8)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(9)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(10)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(11)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(12)).toNumber()] = false;
          _nw62[(new BigNumber(13)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(14)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(15)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(16)).toNumber()] = false;
          _nw62[(new BigNumber(17)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(18)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(19)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(20)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(21)).toNumber()] = _303_v1;
          _nw62[(new BigNumber(22)).toNumber()] = _303_v1;
          _343_v38 = _nw62;
          let _344_v39;
          let _nw63 = Array((new BigNumber(19)).toNumber());
          _nw63[(_dafny.ZERO).toNumber()] = _303_v1;
          _nw63[(_dafny.ONE).toNumber()] = _303_v1;
          _nw63[(new BigNumber(2)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(3)).toNumber()] = (_338_v33)[_module.__default.safeIndex(_this.f12, new BigNumber((_338_v33).length))];
          _nw63[(new BigNumber(4)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(5)).toNumber()] = _module.__default.fm1(_this.f12, true, new BigNumber(716), _303_v1, globalState);
          _nw63[(new BigNumber(6)).toNumber()] = !(_303_v1) || (_303_v1);
          _nw63[(new BigNumber(7)).toNumber()] = !(_303_v1) || (_303_v1);
          _nw63[(new BigNumber(8)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(9)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsPrefixOf(_module.__default.fm11((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _this.f12, _303_v1, globalState), _339_v34);
          _nw63[(new BigNumber(11)).toNumber()] = !_dafny.areEqual(_339_v34, (_340_v35)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_338_v33, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])), new BigNumber((_338_v33).length)), false)).length), new BigNumber((_340_v35).length))]);
          _nw63[(new BigNumber(12)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(13)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_342_v37,_343_v38)).contains(_342_v37);
          _nw63[(new BigNumber(15)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(16)).toNumber()] = _303_v1;
          _nw63[(new BigNumber(17)).toNumber()] = !(_303_v1);
          _nw63[(new BigNumber(18)).toNumber()] = _303_v1;
          _344_v39 = _nw63;
          let _rhs63 = _303_v1;
          let _rhs64 = new BigNumber(-754);
          let _rhs65 = ((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]).multipliedBy((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
          let _rhs66 = _344_v39;
          let _rhs67 = !(!(!(_303_v1)) || (true));
          let _lhs32 = globalState;
          let _lhs33 = globalState;
          r0 = _rhs63;
          _lhs32.f5 = _rhs64;
          _lhs33.f5 = _rhs65;
          _344_v39 = _rhs66;
          r0 = _rhs67;
          let _345_v40;
          _345_v40 = _dafny.MultiSet.fromElements(_this.f12, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
          let _346_v41;
          _346_v41 = _dafny.Map.Empty.slice().updateUnsafe(_303_v1,_303_v1);
          let _347_v42;
          _347_v42 = _dafny.Map.Empty.slice().updateUnsafe((((_345_v40).contains(new BigNumber((_346_v41).length))) ? ((_345_v40).get(new BigNumber((_346_v41).length))) : ((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])),false);
          let _348_v43;
          _348_v43 = _module.__default.fm11(_module.__default.fm0(_this.f12, _this.f12, globalState), (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _303_v1, globalState);
          let _349_v44;
          _349_v44 = _dafny.Map.Empty.slice().updateUnsafe(_303_v1,new BigNumber(((_348_v43)).length));
          let _350_v45;
          let _nw64 = new _module.C0();
          _nw64.__ctor(_347_v42, new BigNumber((_349_v44).length));
          _350_v45 = _nw64;
          let _351_v46;
          let _init8 = ((_352_v41) => function (_353_i3) {
            return _352_v41;
          })(_346_v41);
          let _nw65 = Array((new BigNumber(7)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw65.length); _i0_8++) {
            _nw65[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _351_v46 = _nw65;
          let _index51 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_351_v46).length));
          (_351_v46)[_index51] = (_346_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe(_303_v1,_303_v1));
          let _index52 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_351_v46).length));
          (_351_v46)[_index52] = _346_v41;
          let _354_v47;
          _354_v47 = _dafny.Map.Empty.slice().updateUnsafe(_338_v33,_this.f12);
          _354_v47 = (_354_v47).update(_dafny.Seq.of(_module.__default.fm1((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _303_v1, _this.f12, _303_v1, globalState)), _this.f12);
          let _355_v48;
          _355_v48 = _dafny.Set.fromElements(true);
          let _356_v49;
          _356_v49 = _dafny.Map.Empty.slice().updateUnsafe(_355_v48,(((_345_v40).contains((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])) ? ((_345_v40).get((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])) : (new BigNumber((_339_v34).length))));
          let _357_v50;
          _357_v50 = new _dafny.CodePoint('s'.codePointAt(0));
          let _358_v51;
          let _init9 = ((_359_v44) => function (_360_i4) {
            return _359_v44;
          })(_349_v44);
          let _nw66 = Array((new BigNumber(2)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw66.length); _i0_9++) {
            _nw66[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _358_v51 = _nw66;
          let _index53 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_358_v51).length));
          (_358_v51)[_index53] = _349_v44;
          let _361_v53;
          _361_v53 = _dafny.Seq.of(_355_v48);
          let _362_v54;
          _362_v54 = _355_v48;
          let _363_v55;
          _363_v55 = _module.D2.create_DC8(_303_v1, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _303_v1);
          let _index54 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_358_v51).length));
          let _rhs68 = ((function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (_361_v53).Elements) {
              let _364_v52 = _compr_6;
              if (_dafny.Seq.contains(_361_v53, _364_v52)) {
                _coll6.push([_364_v52,(_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]]);
              }
            }
            return _coll6;
          }()).update((_362_v54), (_363_v55).dtor_cf12)).update(_dafny.Set.fromElements(_303_v1, _303_v1, _303_v1), _module.__default.fm0(_this.f12, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], globalState));
          let _rhs69 = _357_v50;
          let _rhs70 = _349_v44;
          let _lhs34 = _358_v51;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_358_v51).length));
          _356_v49 = _rhs68;
          _357_v50 = _rhs69;
          _lhs34[_lhs35] = _rhs70;
        }
        let _365_v56;
        _365_v56 = _dafny.Seq.of(_303_v1);
        let _366_v57;
        _366_v57 = _module.D2.create_DC7(new BigNumber(706));
        let _367_v58;
        _367_v58 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_365_v56, _dafny.Seq.of(false)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(97)), ((_368_v56) => function (_369_i5) {
          return _368_v56;
        })(_365_v56)))).length), _this.f12, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], (_366_v57).dtor_cf10, (_this.f12).minus(_this.f12));
        _367_v58 = _367_v58;
        _303_v1 = ((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]).isLessThanOrEqualTo(_this.f12);
        let _370_v59;
        _370_v59 = _dafny.Map.Empty.slice().updateUnsafe((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))],_303_v1);
        let _371_v60;
        _371_v60 = _dafny.MultiSet.fromElements(_303_v1, false, _303_v1, _303_v1, _303_v1);
        let _372_v63;
        let _nw67 = Array((new BigNumber(13)).toNumber());
        _nw67[(_dafny.ZERO).toNumber()] = (_370_v59).Merge(_dafny.Map.Empty.slice().updateUnsafe((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))],_303_v1));
        _nw67[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))],_303_v1);
        _nw67[(new BigNumber(2)).toNumber()] = (_370_v59).Merge(_370_v59);
        _nw67[(new BigNumber(3)).toNumber()] = _370_v59;
        _nw67[(new BigNumber(4)).toNumber()] = _370_v59;
        _nw67[(new BigNumber(5)).toNumber()] = _module.__default.fm10(new BigNumber((_371_v60).cardinality()), globalState);
        _nw67[(new BigNumber(6)).toNumber()] = _370_v59;
        _nw67[(new BigNumber(7)).toNumber()] = (_370_v59).Merge(_370_v59);
        _nw67[(new BigNumber(8)).toNumber()] = _370_v59;
        _nw67[(new BigNumber(9)).toNumber()] = _370_v59;
        _nw67[(new BigNumber(10)).toNumber()] = function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(14), new BigNumber(690))) {
            let _373_v61 = _compr_7;
            if (((new BigNumber(14)).isLessThanOrEqualTo(_373_v61)) && ((_373_v61).isLessThan(new BigNumber(690)))) {
              _coll7.push([_module.__default.safeModuloInt(_373_v61, _this.f12),!(_303_v1)]);
            }
          }
          return _coll7;
        }();
        _nw67[(new BigNumber(11)).toNumber()] = _370_v59;
        _nw67[(new BigNumber(12)).toNumber()] = function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of _dafny.IntegerRange(new BigNumber(382), new BigNumber(85))) {
            let _374_v62 = _compr_8;
            if (((new BigNumber(382)).isLessThanOrEqualTo(_374_v62)) && ((_374_v62).isLessThan(new BigNumber(85)))) {
              _coll8.push([(_374_v62).plus(_this.f12),_303_v1]);
            }
          }
          return _coll8;
        }();
        _372_v63 = _nw67;
        let _index55 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_372_v63).length));
        (_372_v63)[_index55] = (_370_v59).update(_this.f12, _303_v1);
        let _index56 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_372_v63).length));
        (_372_v63)[_index56] = _370_v59;
        let _375_v64;
        let _init10 = function (_376_i6) {
          return _module.D2.create_DC9(new BigNumber(524));
        };
        let _nw68 = Array((new BigNumber(4)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw68.length); _i0_10++) {
          _nw68[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _375_v64 = _nw68;
        let _377_v65;
        _377_v65 = _dafny.Map.Empty.slice().updateUnsafe(_367_v58,_375_v64);
        _377_v65 = (_377_v65).update(_dafny.MultiSet.fromElements(_this.f12), _375_v64);
      }
      let _378_v66;
      _378_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,!(!(true)));
      let _379_v67;
      _379_v67 = _dafny.MultiSet.fromElements(_303_v1, _303_v1);
      let _380_v68;
      _380_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(20),new BigNumber((_379_v67).cardinality()));
      let _381_v71;
      _381_v71 = _dafny.Seq.of(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_378_v66)).length))), function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_380_v68).Keys.Elements) {
          let _382_v69 = _compr_9;
          if ((_380_v68).contains(_382_v69)) {
            _coll9.add(_module.__default.safeModuloInt(_382_v69, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC2(!(true), function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),false)).length), new BigNumber(192))).Elements) {
    let _383_v70 = _compr_10;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),false)).length), new BigNumber(192)), _383_v70)) {
      _coll10.push([(_383_v70).plus(new BigNumber(851)),!(true)]);
    }
  }
  return _coll10;
}(), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('f'.codePointAt(0))), true)).dtor_cf2,true)).length)));
          }
        }
        return _coll9;
      }());
      let _384_v72;
      _384_v72 = _dafny.Seq.UnicodeFromString("lvadps");
      let _385_v73;
      _385_v73 = _dafny.Set.fromElements(_this.f12);
      if (((_381_v71)[_module.__default.safeIndex(new BigNumber((_384_v72).length), new BigNumber((_381_v71).length))]).equals((_385_v73).Difference(_385_v73))) {
        let _386_v74;
        _386_v74 = _dafny.Seq.of((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]);
        r0 = _module.__default.fm1((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], false, (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], (new BigNumber((_386_v74).length)).isLessThan((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]), globalState);
        let _387_v75;
        let _nw69 = Array((new BigNumber(14)).toNumber()).fill(false);
        _387_v75 = _nw69;
        let _388_v76;
        _388_v76 = _module.D0.create_DC0(_303_v1);
        let _index57 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_387_v75).length));
        (_387_v75)[_index57] = (_388_v76).dtor_cf0;
        let _389_v77;
        let _nw70 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _389_v77 = _nw70;
        let _390_v78;
        _390_v78 = _dafny.Set.fromElements(_389_v77, _389_v77);
        let _index58 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length));
        let _index59 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_387_v75).length));
        let _rhs71 = new BigNumber((_390_v78).length);
        let _rhs72 = ((_dafny.ZERO).minus((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])).isLessThanOrEqualTo(_module.__default.safeModuloInt((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]));
        let _rhs73 = _this.f12;
        let _rhs74 = (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))];
        let _lhs36 = _301_v0;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length));
        let _lhs38 = _387_v75;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_387_v75).length));
        _lhs36[_lhs37] = _rhs71;
        _lhs38[_lhs39] = _rhs72;
        r3 = _rhs73;
        r3 = _rhs74;
        let _391_v79;
        let _nw71 = new _module.C0();
        _nw71.__ctor(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(550),true), new BigNumber(688));
        _391_v79 = _nw71;
        let _392_v80;
        let _nw72 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
        _392_v80 = _nw72;
        let _393_v81;
        _393_v81 = _dafny.Seq.of(_dafny.Seq.of((_387_v75)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_387_v75).length))]), _dafny.Seq.of((_387_v75)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_387_v75).length))], _303_v1));
        let _index60 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_392_v80).length));
        (_392_v80)[_index60] = _393_v81;
        let _index61 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_392_v80).length));
        (_392_v80)[_index61] = _393_v81;
        _387_v75 = _387_v75;
      } else {
        let _394_v82;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_384_v72).length),true), _this.f12);
        _394_v82 = _nw73;
        let _395_v83;
        let _nw74 = new _module.C0();
        _nw74.__ctor(_378_v66, _this.f12);
        _395_v83 = _nw74;
        _395_v83 = _395_v83;
        let _396_v84;
        let _init11 = ((_397_v72) => function (_398_i7) {
          return _dafny.Seq.Concat(_397_v72, _397_v72);
        })(_384_v72);
        let _nw75 = Array((new BigNumber(11)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw75.length); _i0_11++) {
          _nw75[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _396_v84 = _nw75;
        let _index62 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_396_v84).length));
        (_396_v84)[_index62] = _dafny.Seq.UnicodeFromString("r");
        let _399_v85;
        _399_v85 = _384_v72;
        let _index63 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_396_v84).length));
        (_396_v84)[_index63] = _dafny.Seq.Concat(((_303_v1) ? (_384_v72) : (_dafny.Seq.UnicodeFromString("hu"))), (_399_v85));
        let _index64 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_396_v84).length));
        (_396_v84)[_index64] = _dafny.Seq.Concat(_dafny.Seq.Concat(_384_v72, _dafny.Seq.UnicodeFromString("y")), _dafny.Seq.UnicodeFromString("bsdwgeit"));
      }
      let _400_v86;
      _400_v86 = _dafny.Seq.of(_303_v1);
      let _401_v87;
      _401_v87 = _module.D2.create_DC7(new BigNumber((_400_v86).length));
      let _402_v88;
      _402_v88 = _dafny.Seq.of((_401_v87).dtor_cf10);
      let _403_v89;
      _403_v89 = _module.D1.create_DC4(_301_v0);
      let _404_v90;
      _404_v90 = _dafny.Map.Empty.slice().updateUnsafe((_402_v88)[_module.__default.safeIndex(_this.f12, new BigNumber((_402_v88).length))],(_403_v89).dtor_cf7);
      _404_v90 = (_404_v90).update(((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]).plus((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]), _301_v0);
      let _hi1 = _this.f12;
      for (let _405_i8 = (_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]; _405_i8.isLessThan(_hi1); _405_i8 = _405_i8.plus(_dafny.ONE)) {
        let _406_v91;
        _406_v91 = _dafny.Set.fromElements(_400_v86);
        let _407_v92;
        _407_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_303_v1, _303_v1),_dafny.MultiSet.fromElements(_303_v1, _303_v1, false, _module.__default.fm1(new BigNumber((_385_v73).length), _303_v1, _this.f12, _303_v1, globalState)));
        if (((true) ? ((function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_407_v92).Keys.Elements) {
            let _408_v93 = _compr_11;
            if ((_407_v92).contains(_408_v93)) {
              _coll11.add(_408_v93);
            }
          }
          return _coll11;
        }()).IsProperSubsetOf(_406_v91)) : (!(!(((_303_v1) ? (_303_v1) : (_303_v1))))))) {
          _384_v72 = _dafny.Seq.Concat(_384_v72, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-898)), function (_409_i9) {
            return (_this).f15;
          }));
          _384_v72 = _384_v72;
          r0 = _303_v1;
          let _410_v94;
          _410_v94 = new _dafny.CodePoint('p'.codePointAt(0));
          _410_v94 = (_this).f13;
          let _411_v95;
          let _nw76 = Array((new BigNumber(18)).toNumber()).fill(_module.D0.Default());
          _411_v95 = _nw76;
          _411_v95 = _411_v95;
        } else {
          let _412_v96;
          let _nw77 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _412_v96 = _nw77;
          let _413_v97;
          _413_v97 = _dafny.Map.Empty.slice().updateUnsafe(_412_v96,_303_v1);
          _413_v97 = (_413_v97).update(_412_v96, !(_module.__default.fm1(_this.f12, !(_303_v1), _405_i8, _303_v1, globalState)) || (_303_v1));
          let _414_v98;
          let _nw78 = new _module.C0();
          _nw78.__ctor(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_402_v88).length),_303_v1), _this.f12);
          _414_v98 = _nw78;
          let _415_v99;
          _415_v99 = _dafny.Seq.of(_414_v98, _414_v98, _414_v98, _414_v98, _414_v98);
          let _416_v100;
          _416_v100 = _dafny.Map.Empty.slice().updateUnsafe(_415_v99,_414_v98.f12);
          _416_v100 = (_416_v100).update(_dafny.Seq.of(_414_v98, _414_v98, _414_v98, _414_v98), (_this.f12).multipliedBy((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]));
          let _index65 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length));
          (_301_v0)[_index65] = new BigNumber((_dafny.Seq.update(_384_v72, _module.__default.safeIndex((_this).fm4(globalState), new BigNumber((_384_v72).length)), (_this).f15)).length);
          _380_v68 = _module.__default.fm12((_dafny.ZERO).minus((_this.f12).minus((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))])), _414_v98.f12, _this.f12, (_module.__default.fm0(new BigNumber((_400_v86).length), new BigNumber((_380_v68).length), globalState)).plus((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))]), globalState);
          let _417_v101;
          _417_v101 = _dafny.Set.fromElements(_380_v68, _380_v68, _380_v68);
          let _418_v102;
          let _nw79 = Array((new BigNumber(13)).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = _303_v1;
          _nw79[(_dafny.ONE).toNumber()] = _303_v1;
          _nw79[(new BigNumber(2)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(3)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(4)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(5)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(6)).toNumber()] = true;
          _nw79[(new BigNumber(7)).toNumber()] = !(_303_v1);
          _nw79[(new BigNumber(8)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(9)).toNumber()] = _dafny.Seq.contains(_400_v86, _303_v1);
          _nw79[(new BigNumber(10)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(11)).toNumber()] = _303_v1;
          _nw79[(new BigNumber(12)).toNumber()] = (_417_v101).IsSubsetOf(_417_v101);
          _418_v102 = _nw79;
          _418_v102 = (((((_378_v66).contains(new BigNumber((_380_v68).length))) ? ((_378_v66).get(new BigNumber((_380_v68).length))) : (_303_v1))) ? (_418_v102) : (_418_v102));
        }
        _402_v88 = _402_v88;
        _303_v1 = !(true);
        let _419_v103;
        _419_v103 = _module.D0.create_DC0(_303_v1);
        _303_v1 = !((_419_v103).dtor_cf0);
      }
      r0 = !(_module.__default.fm1(_this.f12, _303_v1, (_this.f12).multipliedBy(_this.f12), false, globalState));
      r1 = _301_v0;
      let _420_v104;
      _420_v104 = _dafny.Set.fromElements(_402_v88);
      let _421_v105;
      _421_v105 = _dafny.MultiSet.fromElements((_301_v0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_301_v0).length))], _this.f12);
      let _422_v106;
      _422_v106 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),!((((_378_v66).contains(new BigNumber((_420_v104).length))) ? ((_378_v66).get(new BigNumber((_420_v104).length))) : (_module.__default.fm1(_this.f12, _303_v1, new BigNumber((_421_v105).cardinality()), _303_v1, globalState)))));
      r2 = (_422_v106).Merge(_422_v106);
      r3 = new BigNumber((_384_v72).length);
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let _423_v0;
      _423_v0 = _dafny.Set.fromElements(_this.f12, new BigNumber((_dafny.Seq.UnicodeFromString("lqyamhoax")).length));
      r1 = (_423_v0).IsSubsetOf(_423_v0);
      (_this).f12 = function (_source3) {
        if (_source3.is_DC8) {
          let _424___mcc_h0 = (_source3).cf11;
          let _425___mcc_h1 = (_source3).cf12;
          let _426___mcc_h2 = (_source3).cf13;
          let _427_cf13 = _426___mcc_h2;
          let _428_cf12 = _425___mcc_h1;
          let _429_cf11 = _424___mcc_h0;
          return _428_cf12;
        } else if (_source3.is_DC9) {
          let _430___mcc_h3 = (_source3).cf14;
          let _431_cf14 = _430___mcc_h3;
          return _431_cf14;
        } else {
          let _432___mcc_h4 = (_source3).cf10;
          let _433_cf10 = _432___mcc_h4;
          return _433_cf10;
        }
      }(_module.D2.create_DC9(_this.f12));
      let _434_v1;
      _434_v1 = true;
      let _435_i0;
      _435_i0 = _dafny.ZERO;
      L1: {
        while (_434_v1) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_435_i0)) {
              break L1;
            }
            _435_i0 = (_435_i0).plus(_dafny.ONE);
            let _436_v2;
            _436_v2 = _dafny.Seq.of(_434_v1, _434_v1);
            let _437_v3;
            _437_v3 = _dafny.Seq.UnicodeFromString("iy");
            let _438_v4;
            _438_v4 = _dafny.Map.Empty.slice().updateUnsafe(((_434_v1) ? (_this.f12) : (_this.f12)),(_module.__default.fm0(new BigNumber((_436_v2).length), new BigNumber((_437_v3).length), globalState)).isLessThanOrEqualTo(_this.f12));
            _438_v4 = (_438_v4).update(_this.f12, !(_434_v1));
            let _439_v5;
            let _nw80 = Array((new BigNumber(10)).toNumber()).fill([]);
            _439_v5 = _nw80;
            let _440_v6;
            _440_v6 = _module.D2.create_DC8(_434_v1, _this.f12, _434_v1);
            let _441_v7;
            let _nw81 = Array((new BigNumber(5)).toNumber());
            _nw81[(_dafny.ZERO).toNumber()] = (_440_v6).dtor_cf11;
            _nw81[(_dafny.ONE).toNumber()] = _434_v1;
            _nw81[(new BigNumber(2)).toNumber()] = _434_v1;
            _nw81[(new BigNumber(3)).toNumber()] = _434_v1;
            _nw81[(new BigNumber(4)).toNumber()] = _434_v1;
            _441_v7 = _nw81;
            let _index66 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_439_v5).length));
            (_439_v5)[_index66] = _441_v7;
            let _index67 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_439_v5).length));
            let _rhs75 = !((_module.__default.fm13(_434_v1, _434_v1, globalState)).contains(_this.f12)) || (!(_434_v1));
            let _rhs76 = _441_v7;
            let _lhs40 = _439_v5;
            let _lhs41 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_439_v5).length));
            _434_v1 = _rhs75;
            _lhs40[_lhs41] = _rhs76;
            let _442_v8;
            _442_v8 = _dafny.Set.fromElements(_434_v1);
            let _443_v9;
            _443_v9 = _module.D2.create_DC9(_this.f12);
            _442_v8 = _module.__default.fm14(_this.f12, _443_v9, globalState);
            let _444_v10;
            _444_v10 = _dafny.Seq.of(_this.f12);
            let _445_v11;
            _445_v11 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), (_this).f15, (_this).f15);
            (globalState).f5 = _module.__default.safeModuloInt(_this.f12, (_444_v10)[_module.__default.safeIndex((((_445_v11).contains((_this).f15)) ? ((_445_v11).get((_this).f15)) : (_this.f12)), new BigNumber((_444_v10).length))]);
          }
        }
      }
      r1 = ((_434_v1) ? (((_434_v1) ? (_434_v1) : (_434_v1))) : (!(_434_v1)));
      let _446_v12;
      _446_v12 = _dafny.Seq.of(_this.f12, _module.__default.safeDivisionInt(new BigNumber(164), _this.f12), _this.f12);
      let _447_v13;
      _447_v13 = _dafny.Seq.of(true, _434_v1);
      let _448_v14;
      _448_v14 = _dafny.Seq.UnicodeFromString("iiunb");
      _446_v12 = _dafny.Seq.update((_this).fm5(_dafny.Seq.IsPrefixOf(_447_v13, _447_v13), _434_v1, _434_v1, true, globalState), _module.__default.safeIndex((new BigNumber((_448_v14).length)).minus(_this.f12), new BigNumber(((_this).fm5(_dafny.Seq.IsPrefixOf(_447_v13, _447_v13), _434_v1, _434_v1, true, globalState)).length)), (_dafny.ZERO).minus((_this.f12).multipliedBy(_this.f12)));
      let _449_v15;
      let _nw82 = Array((new BigNumber(26)).toNumber()).fill([]);
      _449_v15 = _nw82;
      let _450_v16;
      let _init12 = ((_451_v0) => function (_452_i1) {
        return _451_v0;
      })(_423_v0);
      let _nw83 = Array((new BigNumber(25)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw83.length); _i0_12++) {
        _nw83[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _450_v16 = _nw83;
      let _index68 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_449_v15).length));
      (_449_v15)[_index68] = _450_v16;
      let _453_v17;
      _453_v17 = _dafny.Map.Empty.slice().updateUnsafe(_434_v1,_450_v16);
      let _index69 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_449_v15).length));
      (_449_v15)[_index69] = (((_453_v17).contains(_module.__default.fm1((_dafny.ZERO).minus(new BigNumber((_448_v14).length)), _434_v1, _this.f12, _434_v1, globalState))) ? ((_453_v17).get(_module.__default.fm1((_dafny.ZERO).minus(new BigNumber((_448_v14).length)), _434_v1, _this.f12, _434_v1, globalState))) : (_450_v16));
      r0 = (_dafny.ZERO).minus(function (_source4) {
        if (_source4.is_DC1) {
          let _454___mcc_h5 = (_source4).cf1;
          let _455_cf1 = _454___mcc_h5;
          return _module.__default.safeDivisionInt(new BigNumber(592), _this.f12);
        } else if (_source4.is_DC2) {
          let _456___mcc_h6 = (_source4).cf2;
          let _457___mcc_h7 = (_source4).cf3;
          let _458___mcc_h8 = (_source4).cf4;
          let _459___mcc_h9 = (_source4).cf5;
          let _460_cf5 = _459___mcc_h9;
          let _461_cf4 = _458___mcc_h8;
          let _462_cf3 = _457___mcc_h7;
          let _463_cf2 = _456___mcc_h6;
          return (_this.f12).plus(_this.f12);
        } else if (_source4.is_DC0) {
          let _464___mcc_h10 = (_source4).cf0;
          let _465_cf0 = _464___mcc_h10;
          return _this.f12;
        } else {
          let _466___mcc_h11 = (_source4).cf6;
          let _467_cf6 = _466___mcc_h11;
          return new BigNumber(103);
        }
      }(_module.D0.create_DC0(_module.__default.fm1(_this.f12, !(_434_v1), _this.f12, _434_v1, globalState))));
      let _468_v18;
      _468_v18 = _module.D1.create_DC5(_434_v1);
      let _469_v19;
      _469_v19 = _dafny.Map.Empty.slice().updateUnsafe((_468_v18).dtor_cf8,new BigNumber(623));
      r1 = ((!(_434_v1)) ? (_434_v1) : (_module.__default.fm1(_this.f12, !(_434_v1), (_dafny.ZERO).minus(new BigNumber((_469_v19).length)), _434_v1, globalState)));
      let _470_v20;
      _470_v20 = _dafny.Set.fromElements(_434_v1, false, _434_v1, _434_v1);
      let _471_v21;
      _471_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_434_v1);
      let _472_v22;
      _472_v22 = _dafny.Map.Empty.slice().updateUnsafe(true,(((_471_v21).contains(new BigNumber(421))) ? ((_471_v21).get(new BigNumber(421))) : (!(_434_v1))));
      let _473_v23;
      _473_v23 = _dafny.MultiSet.fromElements(_this.f12, new BigNumber((_472_v22).length));
      let _474_v24;
      _474_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_434_v1),_473_v23);
      let _475_v26;
      _475_v26 = _dafny.Seq.of(_470_v20);
      r2 = (_dafny.Map.Empty.slice().updateUnsafe(_470_v20,_473_v23)).Merge((_474_v24).Merge(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_475_v26).Elements) {
          let _476_v25 = _compr_12;
          if (_dafny.Seq.contains(_475_v26, _476_v25)) {
            _coll12.push([_476_v25,_473_v23]);
          }
        }
        return _coll12;
      }()));
      return [r0, r1, r2];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f12 = _dafny.ZERO;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f17, f12) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f12 = f12;
      return;
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f12 = _dafny.ZERO;
      this._f13 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f14 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f14, f13, f12) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f13 = f13;
      (_this)._f12 = f12;
      return;
    }
    fm4(globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(_this.f12, _module.__default.safeModuloInt(new BigNumber(632), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(_this.f12), _dafny.Seq.of(_this.f12));
    };
    fm6(globalState) {
      let _this = this;
      return _module.D0.create_DC0(false);
    };
    fm7(globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(((false) ? (_dafny.Seq.UnicodeFromString("kdgs")) : (_dafny.Seq.UnicodeFromString("xerenpimg"))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sugwvjq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), function (_477_i0) {
        return (_this).f13;
      })))).length);
    };
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _478_v0;
      let _init13 = function (_479_i0) {
        return (_479_i0).plus(_this.f12);
      };
      let _nw84 = Array((new BigNumber(4)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw84.length); _i0_13++) {
        _nw84[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _478_v0 = _nw84;
      let _index70 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length));
      (_478_v0)[_index70] = _module.__default.safeModuloInt(_this.f12, _this.f12);
      let _480_v1;
      _480_v1 = true;
      let _481_v2;
      _481_v2 = _module.D2.create_DC8(_480_v1, _this.f12, _480_v1);
      let _index71 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length));
      (_478_v0)[_index71] = function (_source5) {
        if (_source5.is_DC8) {
          let _482___mcc_h0 = (_source5).cf11;
          let _483___mcc_h1 = (_source5).cf12;
          let _484___mcc_h2 = (_source5).cf13;
          let _485_cf13 = _484___mcc_h2;
          let _486_cf12 = _483___mcc_h1;
          let _487_cf11 = _482___mcc_h0;
          return (_486_cf12).plus(new BigNumber((_dafny.Seq.of(true)).length));
        } else if (_source5.is_DC9) {
          let _488___mcc_h3 = (_source5).cf14;
          let _489_cf14 = _488___mcc_h3;
          return _this.f12;
        } else {
          let _490___mcc_h4 = (_source5).cf10;
          let _491_cf10 = _490___mcc_h4;
          return _this.f12;
        }
      }(_481_v2);
      let _492_i1;
      _492_i1 = _dafny.ZERO;
      L2: {
        while (_480_v1) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_492_i1)) {
              break L2;
            }
            _492_i1 = (_492_i1).plus(_dafny.ONE);
            let _493_v3;
            _493_v3 = _dafny.Seq.UnicodeFromString("dii");
            (globalState).f5 = (_module.__default.fm0(new BigNumber((_493_v3).length), (_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))], globalState)).plus(new BigNumber((_dafny.Seq.update(_493_v3, _module.__default.safeIndex((_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))], new BigNumber((_493_v3).length)), (_this).f13)).length));
            let _494_v4;
            _494_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(858),true);
            let _495_v5;
            let _nw85 = new _module.C0();
            _nw85.__ctor(_494_v4, _module.__default.safeDivisionInt(_module.__default.fm0((_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))], (_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))], globalState), (_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))]));
            _495_v5 = _nw85;
            let _496_v6;
            _496_v6 = _dafny.Seq.of(_480_v1);
            _496_v6 = ((_480_v1) ? (_496_v6) : (_dafny.Seq.update(_dafny.Seq.of(_480_v1, _480_v1, _480_v1, _480_v1, _480_v1), _module.__default.safeIndex((_dafny.ZERO).minus(_this.f12), new BigNumber((_dafny.Seq.of(_480_v1, _480_v1, _480_v1, _480_v1, _480_v1)).length)), _480_v1)));
            let _497_v7;
            let _nw86 = Array((new BigNumber(8)).toNumber()).fill([]);
            _497_v7 = _nw86;
            let _498_v8;
            let _nw87 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _498_v8 = _nw87;
            let _index72 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_497_v7).length));
            (_497_v7)[_index72] = _498_v8;
            let _499_v9;
            _499_v9 = _dafny.Seq.of(_498_v8);
            let _500_v10;
            _500_v10 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13, (_this).f13, (_this).f13);
            let _index73 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_497_v7).length));
            (_497_v7)[_index73] = (_499_v9)[_module.__default.safeIndex((((_500_v10).contains((_this).f13)) ? ((_500_v10).get((_this).f13)) : (_this.f12)), new BigNumber((_499_v9).length))];
          }
        }
      }
      let _501_v11;
      _501_v11 = _dafny.Map.Empty.slice().updateUnsafe((_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))],_478_v0);
      _501_v11 = (_501_v11).update((_dafny.ZERO).minus((new BigNumber(497)).minus((_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))])), _478_v0);
      r3 = (_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))];
      _480_v1 = ((_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))]).isLessThan((_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))]);
      r0 = _480_v1;
      let _502_v12;
      _502_v12 = _dafny.MultiSet.fromElements(_480_v1, _480_v1);
      let _503_v13;
      _503_v13 = _module.D5.create_DC12(_dafny.Seq.of(_502_v12, _502_v12));
      r0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_503_v13).dtor_cf17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), ((_504_v12) => function (_505_i2) {
        return _504_v12;
      })(_502_v12))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), ((_506_v12) => function (_507_i3) {
        return _506_v12;
      })(_502_v12)));
      r1 = _478_v0;
      let _508_v14;
      _508_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_480_v1);
      r2 = (_508_v14).Merge(_508_v14);
      let _509_v15;
      _509_v15 = _dafny.Set.fromElements(_480_v1, _module.__default.fm1(_this.f12, _480_v1, (_478_v0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_478_v0).length))], _480_v1, globalState));
      let _510_v16;
      _510_v16 = _dafny.Set.fromElements(_dafny.Set.fromElements(_480_v1, _480_v1, _480_v1, false), _dafny.Set.fromElements(_480_v1, false, _480_v1, _480_v1), _509_v15);
      r3 = new BigNumber((_510_v16).length);
      return [r0, r1, r2, r3];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _511_v0;
      _511_v0 = true;
      let _512_i0;
      _512_i0 = _dafny.ZERO;
      L3: {
        while (_511_v0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_512_i0)) {
              break L3;
            }
            _512_i0 = (_512_i0).plus(_dafny.ONE);
            let _513_v1;
            _513_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(902),_this.f12);
            let _514_v2;
            let _nw88 = Array((new BigNumber(7)).toNumber()).fill([]);
            _514_v2 = _nw88;
            let _515_v3;
            _515_v3 = _dafny.Seq.of(_this.f12, new BigNumber(842));
            let _516_v4;
            _516_v4 = _dafny.Seq.of(_this.f12, new BigNumber((_515_v3).length));
            let _517_v5;
            let _nw89 = Array((new BigNumber(22)).toNumber());
            _nw89[(_dafny.ZERO).toNumber()] = _511_v0;
            _nw89[(_dafny.ONE).toNumber()] = true;
            _nw89[(new BigNumber(2)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(3)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(4)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(5)).toNumber()] = _module.__default.fm1(_this.f12, _511_v0, _this.f12, _511_v0, globalState);
            _nw89[(new BigNumber(6)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(7)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(8)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(9)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(10)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(11)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(12)).toNumber()] = false;
            _nw89[(new BigNumber(13)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(14)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(15)).toNumber()] = _module.__default.fm1(new BigNumber((_515_v3).length), _511_v0, new BigNumber(129), _511_v0, globalState);
            _nw89[(new BigNumber(16)).toNumber()] = _module.__default.fm1(_this.f12, _511_v0, _this.f12, _511_v0, globalState);
            _nw89[(new BigNumber(17)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(18)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(19)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(20)).toNumber()] = _511_v0;
            _nw89[(new BigNumber(21)).toNumber()] = _511_v0;
            _517_v5 = _nw89;
            let _index74 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_514_v2).length));
            (_514_v2)[_index74] = _517_v5;
            let _index75 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_514_v2).length));
            let _rhs77 = _513_v1;
            let _rhs78 = ((_511_v0) ? (_517_v5) : (_517_v5));
            let _lhs42 = _514_v2;
            let _lhs43 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_514_v2).length));
            _513_v1 = _rhs77;
            _lhs42[_lhs43] = _rhs78;
            let _518_v6;
            _518_v6 = _dafny.Seq.UnicodeFromString("doktuoubt");
            if ((_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_518_v6).length)),_this.f12)).length),_511_v0)).length), new BigNumber(11))).isEqualTo(new BigNumber((_dafny.Seq.Concat(_518_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), ((_519_v6) => function (_520_i1) {
              return (_519_v6)[_module.__default.safeIndex(_this.f12, new BigNumber((_519_v6).length))];
            })(_518_v6)))).length))) {
              let _521_v7;
              _521_v7 = _dafny.Map.Empty.slice().updateUnsafe(_511_v0,_this.f12);
              (globalState).f5 = (_module.__default.safeDivisionInt(new BigNumber((_521_v7).length), _this.f12)).multipliedBy(_this.f12);
              (_this).f12 = _this.f12;
              let _522_v8;
              let _nw90 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _522_v8 = _nw90;
              let _index76 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length));
              (_522_v8)[_index76] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_522_v8)).length), _this.f12);
              let _index77 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length));
              (_522_v8)[_index77] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f12));
              let _523_v9;
              _523_v9 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-655)), function (_524_i2) {
                return (_this).f13;
              }));
              let _525_v10;
              _525_v10 = _module.D2.create_DC8(_511_v0, (_522_v8)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length))], _511_v0);
              let _526_v11;
              _526_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_523_v9)[_module.__default.safeIndex((_522_v8)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length))], new BigNumber((_523_v9).length))]).length),_525_v10);
              let _527_v12;
              _527_v12 = _dafny.Seq.of(_511_v0);
              let _pat_let_tv1 = _511_v0;
              let _index78 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length));
              (_522_v8)[_index78] = new BigNumber(((((_526_v11).update(_this.f12, function (_pat_let5_0) {
                return function (_528_dt__update__tmp_h0) {
                  return function (_pat_let6_0) {
                    return function (_529_dt__update_hcf11_h0) {
                      return _module.D2.create_DC8(_529_dt__update_hcf11_h0, (_528_dt__update__tmp_h0).dtor_cf12, (_528_dt__update__tmp_h0).dtor_cf13);
                    }(_pat_let6_0);
                  }(_pat_let_tv1);
                }(_pat_let5_0);
              }(_525_v10))).update(new BigNumber((_518_v6).length), _525_v10)).update((_522_v8)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length))], _module.D2.create_DC8((_527_v12)[_module.__default.safeIndex(_this.f12, new BigNumber((_527_v12).length))], (_522_v8)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length))], !(!(_511_v0))))).length);
              let _530_v13;
              _530_v13 = _dafny.Set.fromElements((_this).f13, (_this).f13, (_518_v6)[_module.__default.safeIndex(new BigNumber(363), new BigNumber((_518_v6).length))], new _dafny.CodePoint('x'.codePointAt(0)));
              let _531_v14;
              _531_v14 = _dafny.Set.fromElements(_511_v0);
              let _532_v15;
              _532_v15 = _dafny.Seq.of(_531_v14, _531_v14);
              let _533_v16;
              _533_v16 = _module.D6.create_DC15(_516_v4);
              let _534_v17;
              _534_v17 = _dafny.Map.Empty.slice().updateUnsafe(_530_v13,((_511_v0) ? (_dafny.Seq.of(new BigNumber(((_532_v15)[_module.__default.safeIndex(_this.f12, new BigNumber((_532_v15).length))]).length), (_dafny.ZERO).minus(_module.__default.fm0((_522_v8)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_522_v8).length))], _this.f12, globalState)))) : ((_533_v16).dtor_cf21)));
              _534_v17 = _dafny.Map.Empty.slice().updateUnsafe((_530_v13).Union(_530_v13),_516_v4);
            } else {
              let _535_v18;
              _535_v18 = _dafny.Set.fromElements(_515_v3);
              let _536_v19;
              _536_v19 = _dafny.Map.Empty.slice().updateUnsafe(_535_v18,_518_v6);
              _511_v0 = !(!(_511_v0) || (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("hvvjigs"), (((_536_v19).contains(_535_v18)) ? ((_536_v19).get(_535_v18)) : (_518_v6)))));
              let _537_v20;
              let _nw91 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
              _537_v20 = _nw91;
              let _index79 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_537_v20).length));
              (_537_v20)[_index79] = _this.f12;
              let _538_v21;
              _538_v21 = _module.D0.create_DC0(!(_511_v0));
              let _539_v22;
              _539_v22 = _module.D0.create_DC3(_538_v21);
              let _index80 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_537_v20).length));
              (_537_v20)[_index80] = new BigNumber(894);
              let _index81 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_537_v20).length));
              (_537_v20)[_index81] = _this.f12;
              let _index82 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_537_v20).length));
              (_537_v20)[_index82] = (_dafny.ZERO).minus((((_513_v1).contains(_this.f12)) ? ((_513_v1).get(_this.f12)) : (new BigNumber((_518_v6).length))));
              let _pat_let_tv2 = _538_v21;
              let _index83 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_537_v20).length));
              let _index84 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_537_v20).length));
              let _index85 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_537_v20).length));
              let _index86 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_537_v20).length));
              let _rhs79 = _this.f12;
              let _rhs80 = function (_pat_let7_0) {
                return function (_540_dt__update__tmp_h1) {
                  return function (_pat_let8_0) {
                    return function (_541_dt__update_hcf6_h0) {
                      return _module.D0.create_DC3(_541_dt__update_hcf6_h0);
                    }(_pat_let8_0);
                  }(_module.D0.create_DC3(_pat_let_tv2));
                }(_pat_let7_0);
              }(_539_v22);
              let _rhs81 = _this.f12;
              let _rhs82 = _this.f12;
              let _rhs83 = new BigNumber(((((_511_v0) || (!(_511_v0))) ? (p0) : (p0))).cardinality());
              let _lhs44 = _537_v20;
              let _lhs45 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_537_v20).length));
              let _lhs46 = _537_v20;
              let _lhs47 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_537_v20).length));
              let _lhs48 = _537_v20;
              let _lhs49 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_537_v20).length));
              let _lhs50 = _537_v20;
              let _lhs51 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_537_v20).length));
              _lhs44[_lhs45] = _rhs79;
              _539_v22 = _rhs80;
              _lhs46[_lhs47] = _rhs81;
              _lhs48[_lhs49] = _rhs82;
              _lhs50[_lhs51] = _rhs83;
              r1 = _511_v0;
              let _542_v23;
              _542_v23 = _dafny.Set.fromElements(_511_v0, _511_v0);
              (_this).f12 = ((_511_v0) ? ((_537_v20)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_537_v20).length))]) : (new BigNumber((_dafny.Seq.Concat(_516_v4, (_this).fm5(_511_v0, _511_v0, _511_v0, _module.__default.fm1((((_513_v1).contains(_this.f12)) ? ((_513_v1).get(_this.f12)) : (new BigNumber((_542_v23).length))), _511_v0, (_537_v20)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_537_v20).length))], true, globalState), globalState))).length)));
              r1 = !(_511_v0);
            }
            _511_v0 = false;
            let _543_v24;
            _543_v24 = _module.D5.create_DC14(_511_v0);
            let _544_v25;
            _544_v25 = _dafny.Set.fromElements(_511_v0);
            let _545_v26;
            _545_v26 = _dafny.MultiSet.fromElements(_this.f12, new BigNumber((_544_v25).length));
            let _pat_let_tv3 = _545_v26;
            let _pat_let_tv4 = _545_v26;
            let _source6 = function (_pat_let9_0) {
              return function (_546_dt__update__tmp_h2) {
                return function (_pat_let10_0) {
                  return function (_547_dt__update_hcf20_h0) {
                    return _module.D5.create_DC14(_547_dt__update_hcf20_h0);
                  }(_pat_let10_0);
                }((_pat_let_tv3).IsSubsetOf(_pat_let_tv4));
              }(_pat_let9_0);
            }(_543_v24);
            if (_source6.is_DC13) {
              let _548___mcc_h0 = (_source6).cf18;
              let _549___mcc_h1 = (_source6).cf19;
              let _550_cf19 = _549___mcc_h1;
              let _551_cf18 = _548___mcc_h0;
              let _552_v27;
              let _nw92 = Array((_dafny.ONE).toNumber());
              _nw92[(_dafny.ZERO).toNumber()] = (_this.f12).multipliedBy((_516_v4)[_module.__default.safeIndex(_this.f12, new BigNumber((_516_v4).length))]);
              _552_v27 = _nw92;
              let _index87 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_552_v27).length));
              (_552_v27)[_index87] = _this.f12;
              let _index88 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_552_v27).length));
              (_552_v27)[_index88] = _this.f12;
              _518_v6 = _dafny.Seq.UnicodeFromString("nbfyp");
              (globalState).f5 = (_dafny.ZERO).minus(_this.f12);
              let _rhs84 = _551_cf18;
              let _rhs85 = (_dafny.ZERO).minus(_this.f12);
              r1 = _rhs84;
              r2 = _rhs85;
            } else if (_source6.is_DC14) {
              let _553___mcc_h2 = (_source6).cf20;
              let _554_cf20 = _553___mcc_h2;
              let _555_v28;
              _555_v28 = _dafny.Seq.of(_dafny.Seq.contains(_518_v6, new _dafny.CodePoint('n'.codePointAt(0))));
              _511_v0 = (_555_v28)[_module.__default.safeIndex(_this.f12, new BigNumber((_555_v28).length))];
              let _556_v29;
              _556_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.UnicodeFromString("tf"), _dafny.Seq.UnicodeFromString("tbjwfs"), _518_v6);
              r1 = _module.__default.fm1(_this.f12, !(_module.__default.fm1(_this.f12, _554_cf20, _this.f12, _511_v0, globalState)), new BigNumber(((_556_v29).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(653)), function (_557_i3) {
                return (_this).f13;
              }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-737)), function (_558_i4) {
                return (_this).f13;
              })))).cardinality()), _554_cf20, globalState);
              let _559_v30;
              _559_v30 = _dafny.Map.Empty.slice().updateUnsafe(false,_511_v0);
              let _560_v31;
              _560_v31 = _dafny.Seq.of(_544_v25, _544_v25, _544_v25, _544_v25, _dafny.Set.fromElements(true, _554_cf20, _511_v0));
              let _561_v32;
              _561_v32 = _dafny.Map.Empty.slice().updateUnsafe(_560_v31,(_this).f13);
              let _562_v33;
              let _nw93 = Array((new BigNumber(10)).toNumber());
              _nw93[(_dafny.ZERO).toNumber()] = (_this).f13;
              _nw93[(_dafny.ONE).toNumber()] = _module.__default.fm15(_559_v30, new BigNumber((_518_v6).length), globalState);
              _nw93[(new BigNumber(2)).toNumber()] = (((_561_v32).contains(_560_v31)) ? ((_561_v32).get(_560_v31)) : ((_this).f13));
              _nw93[(new BigNumber(3)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(4)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(5)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(6)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(7)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(8)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(9)).toNumber()] = (_this).f13;
              _562_v33 = _nw93;
              let _rhs86 = _562_v33;
              let _rhs87 = ((_dafny.ZERO).minus(_this.f12)).minus(_this.f12);
              let _lhs52 = globalState;
              _562_v33 = _rhs86;
              _lhs52.f5 = _rhs87;
              _511_v0 = _554_cf20;
            } else {
              let _563___mcc_h3 = (_source6).cf17;
              let _564_cf17 = _563___mcc_h3;
              _518_v6 = _dafny.Seq.update(_518_v6, _module.__default.safeIndex(_this.f12, new BigNumber((_518_v6).length)), (_this).f13);
              let _565_v34;
              let _init14 = function (_566_i5) {
                return (_566_i5).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_567_i6) {
                  return (_this).f13;
                })).length)));
              };
              let _nw94 = Array((_dafny.ONE).toNumber());
              for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw94.length); _i0_14++) {
                _nw94[_i0_14] = _init14(new BigNumber(_i0_14));
              }
              _565_v34 = _nw94;
              _565_v34 = (((_511_v0) === (_511_v0)) ? (_565_v34) : (((_511_v0) ? (_565_v34) : (_565_v34))));
              r1 = !(!(_511_v0));
              let _568_v35;
              let _nw95 = new _module.C1();
              _nw95.__ctor((_this).f13, (_this).f13, _this.f12);
              _568_v35 = _nw95;
            }
          }
        }
      }
      let _569_v36;
      _569_v36 = _dafny.Map.Empty.slice().updateUnsafe(_511_v0,_511_v0);
      let _570_v37;
      _570_v37 = _dafny.Seq.UnicodeFromString("gpstgyprm");
      let _571_v38;
      _571_v38 = _dafny.Seq.of(_module.__default.fm1(new BigNumber((_569_v36).length), _511_v0, new BigNumber((_570_v37).length), _511_v0, globalState));
      let _572_v39;
      _572_v39 = _dafny.Set.fromElements(_module.__default.fm1(_this.f12, false, new BigNumber((_571_v38).length), _511_v0, globalState), _511_v0, (((_569_v36).contains(_511_v0)) ? ((_569_v36).get(_511_v0)) : (_511_v0)), _511_v0);
      let _source7 = _572_v39;
      let _573___mcc_h4 = _source7;
      let _574_cf16 = _573___mcc_h4;
      let _575_v40;
      let _nw96 = Array((new BigNumber(8)).toNumber());
      _nw96[(_dafny.ZERO).toNumber()] = !(_module.__default.fm1(new BigNumber(783), _511_v0, _this.f12, _511_v0, globalState));
      _nw96[(_dafny.ONE).toNumber()] = _dafny.Seq.IsPrefixOf(_570_v37, _570_v37);
      _nw96[(new BigNumber(2)).toNumber()] = true;
      _nw96[(new BigNumber(3)).toNumber()] = _511_v0;
      _nw96[(new BigNumber(4)).toNumber()] = _511_v0;
      _nw96[(new BigNumber(5)).toNumber()] = (_this.f12).isLessThan(_this.f12);
      _nw96[(new BigNumber(6)).toNumber()] = !((!(_511_v0)) || (_511_v0));
      _nw96[(new BigNumber(7)).toNumber()] = _511_v0;
      _575_v40 = _nw96;
      let _index89 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_575_v40).length));
      (_575_v40)[_index89] = _511_v0;
      let _576_v41;
      _576_v41 = _dafny.Seq.of(new BigNumber((_570_v37).length));
      let _577_v42;
      _577_v42 = _dafny.Map.Empty.slice().updateUnsafe(_511_v0,(_this).f13);
      let _578_v44;
      _578_v44 = _dafny.Set.fromElements((_this).f13);
      let _index90 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_575_v40).length));
      (_575_v40)[_index90] = ((_511_v0) ? ((new BigNumber((_576_v41).length)).isEqualTo(_this.f12)) : (!(function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_578_v44).Elements) {
          let _579_v43 = _compr_13;
          if ((_578_v44).contains(_579_v43)) {
            _coll13.push([_579_v43,_511_v0]);
          }
        }
        return _coll13;
      }()).contains((((_577_v42).contains(_511_v0)) ? ((_577_v42).get(_511_v0)) : ((_this).f13)))));
      let _580_v46;
      _580_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12);
      let _581_v47;
      _581_v47 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,new BigNumber((_580_v46).length));
      let _582_v48;
      _582_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_581_v47).length),_511_v0);
      let _583_v49;
      let _nw97 = new _module.C0();
      _nw97.__ctor((function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of (_576_v41).Elements) {
          let _584_v45 = _compr_14;
          if (_dafny.Seq.contains(_576_v41, _584_v45)) {
            _coll14.push([(_584_v45).minus((_dafny.ZERO).minus(new BigNumber(-485))),!((_575_v40)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_575_v40).length))])]);
          }
        }
        return _coll14;
      }()).Merge(_582_v48), _this.f12);
      _583_v49 = _nw97;
      let _585_v50;
      let _init15 = function (_586_i7) {
        return (_586_i7).minus(_this.f12);
      };
      let _nw98 = Array((new BigNumber(22)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw98.length); _i0_15++) {
        _nw98[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _585_v50 = _nw98;
      let _index91 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_585_v50).length));
      (_585_v50)[_index91] = _this.f12;
      let _index92 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_585_v50).length));
      (_585_v50)[_index92] = (_this.f12).plus(new BigNumber((_574_cf16).length));
      (_583_v49).f16 = (_583_v49.f16).update(new BigNumber(437), !_dafny.Seq.contains((_this).fm5(!(_511_v0), _511_v0, _511_v0, _511_v0, globalState), _this.f12));
      let _587_v51;
      let _nw99 = Array((new BigNumber(27)).toNumber()).fill(false);
      _587_v51 = _nw99;
      let _index93 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length));
      (_587_v51)[_index93] = _511_v0;
      let _588_v52;
      _588_v52 = _dafny.Seq.of(p0, p0);
      let _589_v53;
      _589_v53 = _module.D5.create_DC12(_588_v52);
      let _590_v54;
      _590_v54 = _dafny.Seq.of(_589_v53);
      let _index94 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length));
      (_587_v51)[_index94] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(_589_v53), _dafny.Seq.update(_590_v54, _module.__default.safeIndex(_this.f12, new BigNumber((_590_v54).length)), _589_v53)), _590_v54);
      let _591_v55;
      _591_v55 = _dafny.MultiSet.fromElements(_571_v38, _dafny.Seq.of((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]), _571_v38);
      let _592_v56;
      _592_v56 = _dafny.Set.fromElements(_this.f12, _this.f12);
      let _593_v57;
      _593_v57 = _dafny.Seq.of(_this.f12, _this.f12, new BigNumber((_592_v56).length), _this.f12);
      let _594_v58;
      _594_v58 = _module.D6.create_DC15(_593_v57);
      _511_v0 = ((_591_v55).Intersect(_591_v55)).IsSubsetOf((_591_v55).Union(_module.__default.fm16(_511_v0, _this.f12, _594_v58, _this.f12, globalState)));
      if (!_dafny.areEqual(((true) ? (_570_v37) : (_570_v37)), _570_v37)) {
        let _595_v59;
        _595_v59 = _module.D2.create_DC8((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))], _this.f12, _511_v0);
        let _pat_let_tv5 = _511_v0;
        let _596_v60;
        let _nw100 = Array((new BigNumber(29)).toNumber());
        _nw100[(_dafny.ZERO).toNumber()] = _595_v59;
        _nw100[(_dafny.ONE).toNumber()] = _595_v59;
        _nw100[(new BigNumber(2)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(3)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(4)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(5)).toNumber()] = _module.__default.fm17(_511_v0, _module.__default.fm1(_this.f12, _511_v0, _this.f12, _511_v0, globalState), globalState);
        _nw100[(new BigNumber(6)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(7)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(8)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(9)).toNumber()] = function (_pat_let11_0) {
          return function (_597_dt__update__tmp_h3) {
            return function (_pat_let12_0) {
              return function (_598_dt__update_hcf13_h0) {
                return _module.D2.create_DC8((_597_dt__update__tmp_h3).dtor_cf11, (_597_dt__update__tmp_h3).dtor_cf12, _598_dt__update_hcf13_h0);
              }(_pat_let12_0);
            }(_pat_let_tv5);
          }(_pat_let11_0);
        }(_595_v59);
        _nw100[(new BigNumber(10)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(11)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(12)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(13)).toNumber()] = _module.D2.create_DC8(false, new BigNumber((_571_v38).length), _511_v0);
        _nw100[(new BigNumber(14)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(15)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(16)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(17)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(18)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(19)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(20)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(21)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(22)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(23)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(24)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(25)).toNumber()] = _module.__default.fm17(_511_v0, _511_v0, globalState);
        _nw100[(new BigNumber(26)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(27)).toNumber()] = _595_v59;
        _nw100[(new BigNumber(28)).toNumber()] = _595_v59;
        _596_v60 = _nw100;
        let _599_v61;
        _599_v61 = _dafny.Seq.of(_596_v60, _596_v60);
        _599_v61 = (((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]) ? (_599_v61) : (_599_v61));
        let _600_v62;
        let _601_v63;
        let _602_v64;
        let _out39;
        let _out40;
        let _out41;
        let _outcollector11 = _module.__default.m0(_this.f12, (new BigNumber(113)).minus((_dafny.ZERO).minus(_this.f12)), _this.f12, (_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))], globalState);
        _out39 = _outcollector11[0];
        _out40 = _outcollector11[1];
        _out41 = _outcollector11[2];
        _600_v62 = _out39;
        _601_v63 = _out40;
        _602_v64 = _out41;
        _511_v0 = (_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))];
        let _603_v65;
        _603_v65 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(_511_v0)),_511_v0);
        _603_v65 = (_603_v65).update(_571_v38, (_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]);
        let _index95 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length));
        (_587_v51)[_index95] = _511_v0;
      } else {
        let _index96 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length));
        (_587_v51)[_index96] = (_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))];
        let _604_v66;
        let _init16 = function (_605_i8) {
          return _dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_dafny.ZERO).minus(_this.f12));
        };
        let _nw101 = Array((new BigNumber(27)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw101.length); _i0_16++) {
          _nw101[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _604_v66 = _nw101;
        let _606_v67;
        _606_v67 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,false);
        let _607_v68;
        _607_v68 = _dafny.Seq.of(_571_v38, _dafny.Seq.of((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))], false, false, _511_v0));
        let _608_v69;
        _608_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_606_v67).length),(_module.D2.create_DC9(_module.__default.fm0(_this.f12, new BigNumber((_dafny.MultiSet.FromArray((_607_v68)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_607_v68).length))])).cardinality()), globalState))).dtor_cf14);
        let _index97 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_604_v66).length));
        (_604_v66)[_index97] = (_608_v69).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12));
        let _609_v71;
        _609_v71 = _dafny.MultiSet.fromElements((_this).f13);
        let _index98 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_604_v66).length));
        let _rhs88 = function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of (_593_v57).Elements) {
            let _610_v70 = _compr_15;
            if (_dafny.Seq.contains(_593_v57, _610_v70)) {
              _coll15.push([(_610_v70).multipliedBy(new BigNumber(901)),_this.f12]);
            }
          }
          return _coll15;
        }();
        let _rhs89 = _570_v37;
        let _rhs90 = (_dafny.MultiSet.FromArray(_module.__default.fm11(_this.f12, new BigNumber((_dafny.MultiSet.FromArray(_593_v57)).cardinality()), (_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))], globalState))).IsSubsetOf((_dafny.MultiSet.fromElements((_this).f13, new _dafny.CodePoint('f'.codePointAt(0)), (_this).f13)).Difference(_609_v71));
        let _rhs91 = _570_v37;
        let _lhs53 = _604_v66;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_604_v66).length));
        _lhs53[_lhs54] = _rhs88;
        _570_v37 = _rhs89;
        _511_v0 = _rhs90;
        _570_v37 = _rhs91;
        let _611_v72;
        let _nw102 = Array((new BigNumber(9)).toNumber()).fill([]);
        _611_v72 = _nw102;
        let _612_v73;
        _612_v73 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_this.f12, false, _this.f12, false, globalState),_611_v72);
        _612_v73 = (_612_v73).update(_511_v0, (((_612_v73).contains(_511_v0)) ? ((_612_v73).get(_511_v0)) : (_611_v72)));
        let _613_v74;
        _613_v74 = new _dafny.CodePoint('o'.codePointAt(0));
        _613_v74 = _module.__default.fm15(_569_v36, (_593_v57)[_module.__default.safeIndex(new BigNumber((_570_v37).length), new BigNumber((_593_v57).length))], globalState);
        _592_v56 = _592_v56;
      }
      let _614_i9;
      _614_i9 = _dafny.ZERO;
      L4: {
        while ((_571_v38)[_module.__default.safeIndex(((_dafny.ZERO).minus(_this.f12)).plus(_this.f12), new BigNumber((_571_v38).length))]) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_614_i9)) {
              break L4;
            }
            _614_i9 = (_614_i9).plus(_dafny.ONE);
            (globalState).f5 = new BigNumber(479);
            let _615_v75;
            _615_v75 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]);
            _615_v75 = (_615_v75).update(_this.f12, !((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]) || ((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]));
            let _616_v76;
            _616_v76 = new _dafny.CodePoint('u'.codePointAt(0));
            _616_v76 = _616_v76;
            _511_v0 = (((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))]) ? (!(_dafny.Set.fromElements(_511_v0, _511_v0)).contains((_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))])) : ((_this.f12).isEqualTo((_dafny.ZERO).minus(_this.f12))));
          }
        }
      }
      r0 = (new BigNumber(832)).multipliedBy(_this.f12);
      r1 = (_587_v51)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((_587_v51).length))];
      let _617_v77;
      _617_v77 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_570_v37, _module.__default.safeIndex(_this.f12, new BigNumber((_570_v37).length)), (_this).f13), _570_v37);
      r2 = _module.__default.safeModuloInt((((_617_v77).contains(_dafny.Seq.UnicodeFromString("ptyrk"))) ? ((_617_v77).get(_dafny.Seq.UnicodeFromString("ptyrk"))) : (_this.f12)), (_this.f12).multipliedBy(_this.f12));
      return [r0, r1, r2];
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _hi2 = (_this.f12).minus(_this.f12);
      for (let _618_i0 = p0; _618_i0.isLessThan(_hi2); _618_i0 = _618_i0.plus(_dafny.ONE)) {
        let _619_v0;
        _619_v0 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f12);
        let _620_v1;
        _620_v1 = _dafny.MultiSet.fromElements(_619_v0, _619_v0);
        _620_v1 = _dafny.MultiSet.FromArray(_dafny.Seq.of(_619_v0));
        let _621_v2;
        let _nw103 = Array((new BigNumber(29)).toNumber()).fill([]);
        _621_v2 = _nw103;
        let _622_v3;
        _622_v3 = _module.D2.create_DC8(p1, new BigNumber(-238), p1);
        let _623_v4;
        _623_v4 = _dafny.Set.fromElements((_622_v3).dtor_cf11);
        let _624_v5;
        _624_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _625_v6;
        let _nw104 = new _module.C2();
        _nw104.__ctor(true, (_dafny.ZERO).minus(_618_i0));
        _625_v6 = _nw104;
        let _626_v7;
        _626_v7 = _dafny.Seq.of(p0, p0, p0, p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_625_v6)).length));
        let _627_v8;
        _627_v8 = _dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC8(p1, p0, p1)).dtor_cf12,_this.f12);
        let _628_v9;
        let _nw105 = Array((new BigNumber(25)).toNumber());
        _nw105[(_dafny.ZERO).toNumber()] = _this.f12;
        _nw105[(_dafny.ONE).toNumber()] = p0;
        _nw105[(new BigNumber(2)).toNumber()] = new BigNumber((_module.__default.fm18(new BigNumber((p2).length), new BigNumber(421), p1, globalState)).length);
        _nw105[(new BigNumber(3)).toNumber()] = _this.f12;
        _nw105[(new BigNumber(4)).toNumber()] = _this.f12;
        _nw105[(new BigNumber(5)).toNumber()] = p0;
        _nw105[(new BigNumber(6)).toNumber()] = new BigNumber((_623_v4).length);
        _nw105[(new BigNumber(7)).toNumber()] = new BigNumber((p2).length);
        _nw105[(new BigNumber(8)).toNumber()] = _this.f12;
        _nw105[(new BigNumber(9)).toNumber()] = _this.f12;
        _nw105[(new BigNumber(10)).toNumber()] = _618_i0;
        _nw105[(new BigNumber(11)).toNumber()] = new BigNumber(203);
        _nw105[(new BigNumber(12)).toNumber()] = p0;
        _nw105[(new BigNumber(13)).toNumber()] = _618_i0;
        _nw105[(new BigNumber(14)).toNumber()] = p0;
        _nw105[(new BigNumber(15)).toNumber()] = _618_i0;
        _nw105[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((p2).length));
        _nw105[(new BigNumber(17)).toNumber()] = new BigNumber(968);
        _nw105[(new BigNumber(18)).toNumber()] = _this.f12;
        _nw105[(new BigNumber(19)).toNumber()] = _618_i0;
        _nw105[(new BigNumber(20)).toNumber()] = new BigNumber(((_624_v5).update(new BigNumber((_626_v7).length), p1)).length);
        _nw105[(new BigNumber(21)).toNumber()] = p0;
        _nw105[(new BigNumber(22)).toNumber()] = p0;
        _nw105[(new BigNumber(23)).toNumber()] = new BigNumber((_627_v8).length);
        _nw105[(new BigNumber(24)).toNumber()] = _618_i0;
        _628_v9 = _nw105;
        let _index99 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_621_v2).length));
        (_621_v2)[_index99] = _628_v9;
        let _629_v10;
        _629_v10 = _module.D6.create_DC15(((p1) ? (_626_v7) : (_626_v7)));
        let _index100 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_621_v2).length));
        let _rhs92 = _625_v6.f12;
        let _rhs93 = _628_v9;
        let _rhs94 = p1;
        let _rhs95 = p1;
        let _rhs96 = _module.D6.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), ((_630_p2) => function (_631_i1) {
  return new BigNumber((_630_p2).length);
})(p2)));
        let _lhs55 = _621_v2;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_621_v2).length));
        r2 = _rhs92;
        _lhs55[_lhs56] = _rhs93;
        r1 = _rhs94;
        r1 = _rhs95;
        _629_v10 = _rhs96;
        let _index101 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_621_v2).length));
        let _nw106 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        (_621_v2)[_index101] = _nw106;
        let _632_v11;
        _632_v11 = _dafny.Seq.of(p1, p1);
        let _633_v12;
        let _nw107 = Array((new BigNumber(16)).toNumber());
        _nw107[(_dafny.ZERO).toNumber()] = p1;
        _nw107[(_dafny.ONE).toNumber()] = _module.__default.fm1(p0, p1, new BigNumber(-195), p1, globalState);
        _nw107[(new BigNumber(2)).toNumber()] = p1;
        _nw107[(new BigNumber(3)).toNumber()] = p1;
        _nw107[(new BigNumber(4)).toNumber()] = p1;
        _nw107[(new BigNumber(5)).toNumber()] = p1;
        _nw107[(new BigNumber(6)).toNumber()] = p1;
        _nw107[(new BigNumber(7)).toNumber()] = p1;
        _nw107[(new BigNumber(8)).toNumber()] = false;
        _nw107[(new BigNumber(9)).toNumber()] = p1;
        _nw107[(new BigNumber(10)).toNumber()] = (_632_v11)[_module.__default.safeIndex(p0, new BigNumber((_632_v11).length))];
        _nw107[(new BigNumber(11)).toNumber()] = p1;
        _nw107[(new BigNumber(12)).toNumber()] = p1;
        _nw107[(new BigNumber(13)).toNumber()] = p1;
        _nw107[(new BigNumber(14)).toNumber()] = p1;
        _nw107[(new BigNumber(15)).toNumber()] = p1;
        _633_v12 = _nw107;
        let _634_v13;
        _634_v13 = _dafny.Map.Empty.slice().updateUnsafe(_618_i0,_633_v12);
        let _635_v14;
        _635_v14 = _dafny.Seq.of(_633_v12);
        _634_v13 = (_634_v13).update(_this.f12, (_635_v14)[_module.__default.safeIndex((_626_v7)[_module.__default.safeIndex(_this.f12, new BigNumber((_626_v7).length))], new BigNumber((_635_v14).length))]);
      }
      if (p1) {
        let _636_v15;
        _636_v15 = _dafny.Set.fromElements(_this.f12);
        let _637_v16;
        _637_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f12);
        let _638_v17;
        _638_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,p1);
        let _639_v18;
        let _nw108 = new _module.C0();
        _nw108.__ctor(_638_v17, _this.f12);
        _639_v18 = _nw108;
        let _640_v19;
        _640_v19 = _dafny.Seq.of(_639_v18);
        let _641_v20;
        _641_v20 = _dafny.MultiSet.fromElements(false);
        let _642_v21;
        let _nw109 = Array((new BigNumber(29)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = p0;
        _nw109[(_dafny.ONE).toNumber()] = _this.f12;
        _nw109[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_this.f12);
        _nw109[(new BigNumber(3)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(4)).toNumber()] = new BigNumber((_636_v15).length);
        _nw109[(new BigNumber(5)).toNumber()] = p0;
        _nw109[(new BigNumber(6)).toNumber()] = p0;
        _nw109[(new BigNumber(7)).toNumber()] = p0;
        _nw109[(new BigNumber(8)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(9)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(10)).toNumber()] = p0;
        _nw109[(new BigNumber(11)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(12)).toNumber()] = (((_637_v16).contains(p1)) ? ((_637_v16).get(p1)) : (p0));
        _nw109[(new BigNumber(13)).toNumber()] = p0;
        _nw109[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_640_v19, _module.__default.safeIndex(new BigNumber((_641_v20).cardinality()), new BigNumber((_640_v19).length)), _639_v18)).length);
        _nw109[(new BigNumber(15)).toNumber()] = p0;
        _nw109[(new BigNumber(16)).toNumber()] = p0;
        _nw109[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw109[(new BigNumber(18)).toNumber()] = _639_v18.f12;
        _nw109[(new BigNumber(19)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(20)).toNumber()] = p0;
        _nw109[(new BigNumber(21)).toNumber()] = _639_v18.f12;
        _nw109[(new BigNumber(22)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw109[(new BigNumber(24)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(25)).toNumber()] = p0;
        _nw109[(new BigNumber(26)).toNumber()] = p0;
        _nw109[(new BigNumber(27)).toNumber()] = _this.f12;
        _nw109[(new BigNumber(28)).toNumber()] = _639_v18.f12;
        _642_v21 = _nw109;
        let _643_v22;
        _643_v22 = _module.D1.create_DC4(_642_v21);
        let _644_v23;
        _644_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p1,(_643_v22).dtor_cf7),((p1) ? (p1) : (p1)));
        let _645_v24;
        _645_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,_642_v21);
        let _646_v25;
        _646_v25 = _module.D7.create_DC20(_645_v24);
        let _647_v26;
        _647_v26 = _dafny.Seq.of(!(p1));
        _644_v23 = (_644_v23).update((_646_v25).dtor_cf27, !(_module.__default.fm1(new BigNumber((_dafny.MultiSet.FromArray(_647_v26)).cardinality()), false, p0, p1, globalState)));
        r2 = _this.f12;
        r0 = p1;
        let _648_v27;
        _648_v27 = new _dafny.CodePoint('b'.codePointAt(0));
        _648_v27 = (_this).f13;
        let _649_v28;
        let _nw110 = Array((new BigNumber(4)).toNumber()).fill(false);
        _649_v28 = _nw110;
        let _index102 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_649_v28).length));
        (_649_v28)[_index102] = p1;
        let _index103 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_649_v28).length));
        (_649_v28)[_index103] = p1;
      } else {
        let _650_v29;
        _650_v29 = _dafny.Seq.UnicodeFromString("cggpcap");
        _650_v29 = _650_v29;
        let _651_v30;
        let _nw111 = new _module.C0();
        _nw111.__ctor((_dafny.Map.Empty.slice().updateUnsafe(_this.f12,true)).update(new BigNumber(25), p1), _this.f12);
        _651_v30 = _nw111;
        _651_v30 = _651_v30;
        if (p1) {
          let _652_v31;
          _652_v31 = new _dafny.CodePoint('e'.codePointAt(0));
          _652_v31 = (_this).f13;
          let _653_v32;
          _653_v32 = _dafny.Seq.of(!(p1));
          let _654_v33;
          _654_v33 = _dafny.Map.Empty.slice().updateUnsafe((_653_v32)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_653_v32).length))],p1);
          let _655_v34;
          _655_v34 = _dafny.Set.fromElements(p1, false, p1);
          let _656_v35;
          let _nw112 = Array((new BigNumber(22)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = p1;
          _nw112[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).equals(_654_v33);
          _nw112[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_this.f12, p1, p0, p1, globalState);
          _nw112[(new BigNumber(3)).toNumber()] = p1;
          _nw112[(new BigNumber(4)).toNumber()] = p1;
          _nw112[(new BigNumber(5)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
          _nw112[(new BigNumber(6)).toNumber()] = !(p1) || (p1);
          _nw112[(new BigNumber(7)).toNumber()] = p1;
          _nw112[(new BigNumber(8)).toNumber()] = p1;
          _nw112[(new BigNumber(9)).toNumber()] = p1;
          _nw112[(new BigNumber(10)).toNumber()] = p1;
          _nw112[(new BigNumber(11)).toNumber()] = p1;
          _nw112[(new BigNumber(12)).toNumber()] = p1;
          _nw112[(new BigNumber(13)).toNumber()] = p1;
          _nw112[(new BigNumber(14)).toNumber()] = p1;
          _nw112[(new BigNumber(15)).toNumber()] = ((p1) ? (p1) : (p1));
          _nw112[(new BigNumber(16)).toNumber()] = p1;
          _nw112[(new BigNumber(17)).toNumber()] = (_655_v34).IsDisjointFrom(_655_v34);
          _nw112[(new BigNumber(18)).toNumber()] = true;
          _nw112[(new BigNumber(19)).toNumber()] = p1;
          _nw112[(new BigNumber(20)).toNumber()] = p1;
          _nw112[(new BigNumber(21)).toNumber()] = ((p1) ? (p1) : (p1));
          _656_v35 = _nw112;
          let _index104 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_656_v35).length));
          (_656_v35)[_index104] = p1;
          let _657_v36;
          _657_v36 = _dafny.Seq.of(new BigNumber(-803));
          let _index105 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_656_v35).length));
          (_656_v35)[_index105] = !(new BigNumber((_dafny.Set.fromElements(_656_v35)).length)).isEqualTo(new BigNumber((_657_v36).length));
          let _658_v37;
          let _init17 = ((_659_v29, _660_p2) => function (_661_i2) {
            return _dafny.Seq.Concat(_659_v29, _660_p2);
          })(_650_v29, p2);
          let _nw113 = Array((new BigNumber(5)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw113.length); _i0_17++) {
            _nw113[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _658_v37 = _nw113;
          let _index106 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_658_v37).length));
          (_658_v37)[_index106] = p2;
          let _index107 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_658_v37).length));
          let _index108 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_656_v35).length));
          let _rhs97 = p1;
          let _rhs98 = _dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("b"));
          let _rhs99 = new _dafny.CodePoint('d'.codePointAt(0));
          let _rhs100 = p1;
          let _lhs57 = _658_v37;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_658_v37).length));
          let _lhs59 = _656_v35;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_656_v35).length));
          r0 = _rhs97;
          _lhs57[_lhs58] = _rhs98;
          _652_v31 = _rhs99;
          _lhs59[_lhs60] = _rhs100;
          let _662_v40;
          let _nw114 = new _module.C0();
          _nw114.__ctor((_651_v30.f16).Merge(function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of (function () {
              let _coll17 = new _dafny.Set();
              for (const _compr_17 of _dafny.IntegerRange(new BigNumber(307), new BigNumber(123))) {
                let _663_v39 = _compr_17;
                if (((new BigNumber(307)).isLessThanOrEqualTo(_663_v39)) && ((_663_v39).isLessThan(new BigNumber(123)))) {
                  _coll17.add(_module.__default.safeDivisionInt(_663_v39, _this.f12));
                }
              }
              return _coll17;
            }()).Elements) {
              let _664_v38 = _compr_16;
              if ((function () {
                let _coll18 = new _dafny.Set();
                for (const _compr_18 of _dafny.IntegerRange(new BigNumber(307), new BigNumber(123))) {
                  let _665_v39 = _compr_18;
                  if (((new BigNumber(307)).isLessThanOrEqualTo(_665_v39)) && ((_665_v39).isLessThan(new BigNumber(123)))) {
                    _coll18.add(_module.__default.safeDivisionInt(_665_v39, _this.f12));
                  }
                }
                return _coll18;
              }()).contains(_664_v38)) {
                _coll16.push([(_664_v38).minus(p0),(_656_v35)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_656_v35).length))]]);
              }
            }
            return _coll16;
          }()), _this.f12);
          _662_v40 = _nw114;
          _657_v36 = _dafny.Seq.of(_this.f12, _module.__default.safeModuloInt(_this.f12, new BigNumber(476)), p0);
        } else {
          let _666_v41;
          _666_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
          let _667_v42;
          _667_v42 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(653)), function (_668_i3) {
            return (_this).f13;
          })).length), new BigNumber((_666_v41).length));
          let _669_v43;
          _669_v43 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.fromElements(_651_v30.f16)).update(_651_v30.f16, _module.__default.abs((_dafny.ZERO).minus(p0)))).cardinality())), new BigNumber((_dafny.Seq.of(p1, p1)).length));
          let _670_v44;
          _670_v44 = _dafny.Set.fromElements(_669_v43);
          let _671_v45;
          _671_v45 = _dafny.Map.Empty.slice().updateUnsafe(_667_v42,new BigNumber((_670_v44).length));
          _671_v45 = _dafny.Map.Empty.slice().updateUnsafe(_667_v42,_this.f12);
          (globalState).f5 = p0;
          _650_v29 = _dafny.Seq.update(p2, _module.__default.safeIndex(p0, new BigNumber((p2).length)), (_this).f13);
          let _672_v46;
          let _nw115 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _672_v46 = _nw115;
          let _index109 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_672_v46).length));
          (_672_v46)[_index109] = p0;
          let _index110 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_672_v46).length));
          (_672_v46)[_index110] = (_dafny.ZERO).minus(new BigNumber((_650_v29).length));
          let _673_v47;
          _673_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f13);
          let _674_v48;
          let _nw116 = Array((new BigNumber(29)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _673_v47;
          _nw116[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,new _dafny.CodePoint('v'.codePointAt(0)));
          _nw116[(new BigNumber(2)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f13);
          _nw116[(new BigNumber(4)).toNumber()] = ((_673_v47).update(!(p1), (_this).f13)).Merge(_673_v47);
          _nw116[(new BigNumber(5)).toNumber()] = (_673_v47).Merge(_673_v47);
          _nw116[(new BigNumber(6)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(7)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(8)).toNumber()] = (_673_v47).Merge(_673_v47);
          _nw116[(new BigNumber(9)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(10)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(11)).toNumber()] = (_673_v47).Merge(_673_v47);
          _nw116[(new BigNumber(12)).toNumber()] = ((p1) ? (_673_v47) : (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f13)));
          _nw116[(new BigNumber(13)).toNumber()] = (_673_v47).Merge(_673_v47);
          _nw116[(new BigNumber(14)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f13);
          _nw116[(new BigNumber(16)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(17)).toNumber()] = (((_673_v47).update(p1, (_this).f13)).update(p1, new _dafny.CodePoint('t'.codePointAt(0)))).Merge(_673_v47);
          _nw116[(new BigNumber(18)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(19)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(20)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(21)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(22)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(23)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(24)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(25)).toNumber()] = (_673_v47).update(false, (_this).f13);
          _nw116[(new BigNumber(26)).toNumber()] = _673_v47;
          _nw116[(new BigNumber(27)).toNumber()] = (_673_v47).Merge(_673_v47);
          _nw116[(new BigNumber(28)).toNumber()] = _673_v47;
          _674_v48 = _nw116;
          let _index111 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_674_v48).length));
          (_674_v48)[_index111] = (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f13)).update(p1, (_this).f13);
          let _index112 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_674_v48).length));
          (_674_v48)[_index112] = _673_v47;
        }
        let _675_v49;
        let _nw117 = Array((new BigNumber(6)).toNumber()).fill(_module.D6.Default());
        _675_v49 = _nw117;
        let _676_v50;
        _676_v50 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _677_v51;
        let _nw118 = Array((new BigNumber(28)).toNumber());
        _nw118[(_dafny.ZERO).toNumber()] = p1;
        _nw118[(_dafny.ONE).toNumber()] = p1;
        _nw118[(new BigNumber(2)).toNumber()] = p1;
        _nw118[(new BigNumber(3)).toNumber()] = false;
        _nw118[(new BigNumber(4)).toNumber()] = false;
        _nw118[(new BigNumber(5)).toNumber()] = p1;
        _nw118[(new BigNumber(6)).toNumber()] = false;
        _nw118[(new BigNumber(7)).toNumber()] = p1;
        _nw118[(new BigNumber(8)).toNumber()] = p1;
        _nw118[(new BigNumber(9)).toNumber()] = p1;
        _nw118[(new BigNumber(10)).toNumber()] = p1;
        _nw118[(new BigNumber(11)).toNumber()] = p1;
        _nw118[(new BigNumber(12)).toNumber()] = true;
        _nw118[(new BigNumber(13)).toNumber()] = p1;
        _nw118[(new BigNumber(14)).toNumber()] = p1;
        _nw118[(new BigNumber(15)).toNumber()] = p1;
        _nw118[(new BigNumber(16)).toNumber()] = p1;
        _nw118[(new BigNumber(17)).toNumber()] = p1;
        _nw118[(new BigNumber(18)).toNumber()] = p1;
        _nw118[(new BigNumber(19)).toNumber()] = p1;
        _nw118[(new BigNumber(20)).toNumber()] = (((_676_v50).contains(p1)) ? ((_676_v50).get(p1)) : (p1));
        _nw118[(new BigNumber(21)).toNumber()] = p1;
        _nw118[(new BigNumber(22)).toNumber()] = p1;
        _nw118[(new BigNumber(23)).toNumber()] = p1;
        _nw118[(new BigNumber(24)).toNumber()] = p1;
        _nw118[(new BigNumber(25)).toNumber()] = false;
        _nw118[(new BigNumber(26)).toNumber()] = !(false);
        _nw118[(new BigNumber(27)).toNumber()] = p1;
        _677_v51 = _nw118;
        let _678_v52;
        _678_v52 = _module.D6.create_DC16(_677_v51, _this.f12);
        let _679_v53;
        _679_v53 = _module.D6.create_DC19(_678_v52);
        let _index113 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_675_v49).length));
        (_675_v49)[_index113] = _679_v53;
        let _680_v54;
        let _init18 = function (_681_i4) {
          return (_this).f13;
        };
        let _nw119 = Array((new BigNumber(10)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw119.length); _i0_18++) {
          _nw119[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _680_v54 = _nw119;
        let _index114 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_680_v54).length));
        (_680_v54)[_index114] = new _dafny.CodePoint('l'.codePointAt(0));
        let _682_v55;
        _682_v55 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(864));
        let _index115 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_675_v49).length));
        let _index116 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_680_v54).length));
        let _rhs101 = p1;
        let _rhs102 = _679_v53;
        let _rhs103 = new _dafny.CodePoint('k'.codePointAt(0));
        let _rhs104 = (_682_v55).Merge(_682_v55);
        let _lhs61 = _675_v49;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_675_v49).length));
        let _lhs63 = _680_v54;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_680_v54).length));
        r1 = _rhs101;
        _lhs61[_lhs62] = _rhs102;
        _lhs63[_lhs64] = _rhs103;
        _682_v55 = _rhs104;
        let _683_v56;
        _683_v56 = _dafny.Seq.of(p1);
        let _684_v57;
        _684_v57 = _dafny.Seq.of(_683_v56);
        let _685_v58;
        let _nw120 = Array((new BigNumber(9)).toNumber());
        _nw120[(_dafny.ZERO).toNumber()] = _683_v56;
        _nw120[(_dafny.ONE).toNumber()] = _683_v56;
        _nw120[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(p1);
        _nw120[(new BigNumber(3)).toNumber()] = _683_v56;
        _nw120[(new BigNumber(4)).toNumber()] = _683_v56;
        _nw120[(new BigNumber(5)).toNumber()] = _683_v56;
        _nw120[(new BigNumber(6)).toNumber()] = (_684_v57)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_684_v57, _module.__default.safeIndex(p0, new BigNumber((_684_v57).length)), _module.__default.fm9(p1, globalState))).length), new BigNumber((_684_v57).length))];
        _nw120[(new BigNumber(7)).toNumber()] = _683_v56;
        _nw120[(new BigNumber(8)).toNumber()] = _683_v56;
        _685_v58 = _nw120;
        let _686_v59;
        _686_v59 = _dafny.Seq.of(_685_v58);
        let _687_v60;
        let _nw121 = Array((new BigNumber(6)).toNumber());
        _nw121[(_dafny.ZERO).toNumber()] = _685_v58;
        _nw121[(_dafny.ONE).toNumber()] = _685_v58;
        _nw121[(new BigNumber(2)).toNumber()] = _685_v58;
        _nw121[(new BigNumber(3)).toNumber()] = _685_v58;
        _nw121[(new BigNumber(4)).toNumber()] = _685_v58;
        _nw121[(new BigNumber(5)).toNumber()] = (_686_v59)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f12), new BigNumber((_686_v59).length))];
        _687_v60 = _nw121;
        let _index117 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_687_v60).length));
        (_687_v60)[_index117] = _685_v58;
        let _index118 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_687_v60).length));
        (_687_v60)[_index118] = _685_v58;
      }
      let _688_v61;
      _688_v61 = _dafny.MultiSet.fromElements(p0);
      let _689_i5;
      _689_i5 = _dafny.ZERO;
      L5: {
        while (!(_688_v61).contains(p0)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_689_i5)) {
              break L5;
            }
            _689_i5 = (_689_i5).plus(_dafny.ONE);
            let _rhs105 = (_dafny.ZERO).minus(_this.f12);
            let _rhs106 = _this.f12;
            let _lhs65 = globalState;
            _lhs65.f5 = _rhs105;
            r2 = _rhs106;
            let _690_v62;
            _690_v62 = _module.D5.create_DC14(p1);
            let _source8 = _690_v62;
            if (_source8.is_DC13) {
              let _691___mcc_h0 = (_source8).cf18;
              let _692___mcc_h1 = (_source8).cf19;
              let _693_cf19 = _692___mcc_h1;
              let _694_cf18 = _691___mcc_h0;
              let _695_v63;
              _695_v63 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,!(p1));
              _694_cf18 = (((_695_v63).contains(_this.f12)) ? ((_695_v63).get(_this.f12)) : (!(_693_cf19)));
              r1 = !(!(p1));
              (globalState).f5 = ((_693_cf19) ? (new BigNumber(507)) : (_this.f12));
              let _696_v64;
              let _nw122 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
              _696_v64 = _nw122;
              let _index119 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_696_v64).length));
              (_696_v64)[_index119] = _module.__default.safeModuloInt(p0, new BigNumber((p2).length));
              let _index120 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_696_v64).length));
              (_696_v64)[_index120] = (_dafny.ZERO).minus(p0);
            } else if (_source8.is_DC14) {
              let _697___mcc_h2 = (_source8).cf20;
              let _698_cf20 = _697___mcc_h2;
              let _699_v65;
              _699_v65 = _dafny.Seq.of(p1);
              let _700_v66;
              let _nw123 = Array((new BigNumber(21)).toNumber());
              _nw123[(_dafny.ZERO).toNumber()] = !(_698_cf20);
              _nw123[(_dafny.ONE).toNumber()] = p1;
              _nw123[(new BigNumber(2)).toNumber()] = _698_cf20;
              _nw123[(new BigNumber(3)).toNumber()] = p1;
              _nw123[(new BigNumber(4)).toNumber()] = !(p1);
              _nw123[(new BigNumber(5)).toNumber()] = _698_cf20;
              _nw123[(new BigNumber(6)).toNumber()] = p1;
              _nw123[(new BigNumber(7)).toNumber()] = !((p0).isLessThan(p0));
              _nw123[(new BigNumber(8)).toNumber()] = p1;
              _nw123[(new BigNumber(9)).toNumber()] = !(_698_cf20);
              _nw123[(new BigNumber(10)).toNumber()] = p1;
              _nw123[(new BigNumber(11)).toNumber()] = _698_cf20;
              _nw123[(new BigNumber(12)).toNumber()] = (_699_v65)[_module.__default.safeIndex(p0, new BigNumber((_699_v65).length))];
              _nw123[(new BigNumber(13)).toNumber()] = _698_cf20;
              _nw123[(new BigNumber(14)).toNumber()] = !(p1) || (p1);
              _nw123[(new BigNumber(15)).toNumber()] = ((p1) ? (false) : (_698_cf20));
              _nw123[(new BigNumber(16)).toNumber()] = _698_cf20;
              _nw123[(new BigNumber(17)).toNumber()] = !(((_698_cf20) ? (_698_cf20) : (p1)));
              _nw123[(new BigNumber(18)).toNumber()] = (_this.f12).isLessThanOrEqualTo(_this.f12);
              _nw123[(new BigNumber(19)).toNumber()] = (p1) === (_698_cf20);
              _nw123[(new BigNumber(20)).toNumber()] = _698_cf20;
              _700_v66 = _nw123;
              let _index121 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_700_v66).length));
              (_700_v66)[_index121] = (p0).isEqualTo(p0);
              let _701_v67;
              _701_v67 = _dafny.Seq.of(new BigNumber(850));
              let _index122 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_700_v66).length));
              (_700_v66)[_index122] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_this.f12), ((p1) ? (_701_v67) : ((_this).fm5(p1, _698_cf20, _698_cf20, !(p1), globalState))));
              (globalState).f5 = p0;
              let _702_v68;
              _702_v68 = _module.D2.create_DC8(_698_cf20, p0, _698_cf20);
              let _index123 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_700_v66).length));
              (_700_v66)[_index123] = !(_module.__default.fm1(new BigNumber(757), _module.__default.fm1(_this.f12, (_700_v66)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_700_v66).length))], _this.f12, (_700_v66)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_700_v66).length))], globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-788)),p1)).length), (_702_v68).dtor_cf13, globalState));
              let _703_v69;
              _703_v69 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_698_cf20,!(_module.__default.fm1(new BigNumber(-745), (_700_v66)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_700_v66).length))], p0, _698_cf20, globalState))));
              let _704_v71;
              _704_v71 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)));
              let _705_v72;
              let _nw124 = Array((new BigNumber(17)).toNumber());
              _nw124[(_dafny.ZERO).toNumber()] = new BigNumber(176);
              _nw124[(_dafny.ONE).toNumber()] = _this.f12;
              _nw124[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(_this.f12));
              _nw124[(new BigNumber(3)).toNumber()] = ((false) ? ((_dafny.ZERO).minus(new BigNumber((_703_v69).length))) : (_this.f12));
              _nw124[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber(440));
              _nw124[(new BigNumber(5)).toNumber()] = p0;
              _nw124[(new BigNumber(6)).toNumber()] = _this.f12;
              _nw124[(new BigNumber(7)).toNumber()] = p0;
              _nw124[(new BigNumber(8)).toNumber()] = _this.f12;
              _nw124[(new BigNumber(9)).toNumber()] = new BigNumber((function () {
                let _coll19 = new _dafny.Map();
                for (const _compr_19 of (_704_v71).Elements) {
                  let _706_v70 = _compr_19;
                  if ((_704_v71).contains(_706_v70)) {
                    _coll19.push([_706_v70,_698_cf20]);
                  }
                }
                return _coll19;
              }()).length);
              _nw124[(new BigNumber(10)).toNumber()] = p0;
              _nw124[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-239));
              _nw124[(new BigNumber(12)).toNumber()] = (_this.f12).plus(new BigNumber((_dafny.MultiSet.FromArray(_699_v65)).cardinality()));
              _nw124[(new BigNumber(13)).toNumber()] = (_this.f12).plus(p0);
              _nw124[(new BigNumber(14)).toNumber()] = _this.f12;
              _nw124[(new BigNumber(15)).toNumber()] = new BigNumber(536);
              _nw124[(new BigNumber(16)).toNumber()] = _this.f12;
              _705_v72 = _nw124;
              _705_v72 = _705_v72;
            } else {
              let _707___mcc_h3 = (_source8).cf17;
              let _708_cf17 = _707___mcc_h3;
              let _709_v73;
              let _nw125 = new _module.C1();
              _nw125.__ctor((_this).f13, (_this).f13, _this.f12);
              _709_v73 = _nw125;
              let _710_v74;
              let _init19 = ((_711_p2) => function (_712_i6) {
                return _module.__default.safeDivisionInt(_712_i6, new BigNumber((_dafny.Set.fromElements(new BigNumber((_711_p2).length))).length));
              })(p2);
              let _nw126 = Array((new BigNumber(28)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw126.length); _i0_19++) {
                _nw126[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _710_v74 = _nw126;
              let _index124 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length));
              (_710_v74)[_index124] = new BigNumber(299);
              let _713_v75;
              _713_v75 = new _dafny.CodePoint('m'.codePointAt(0));
              let _index125 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length));
              let _rhs107 = _709_v73;
              let _rhs108 = p0;
              let _rhs109 = (_709_v73).f15;
              let _lhs66 = _710_v74;
              let _lhs67 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length));
              _709_v73 = _rhs107;
              _lhs66[_lhs67] = _rhs108;
              _713_v75 = _rhs109;
              _713_v75 = (_709_v73).f15;
              (globalState).f5 = (_709_v73).fm4(globalState);
              let _714_v76;
              let _715_v77;
              let _716_v78;
              let _out42;
              let _out43;
              let _out44;
              let _outcollector12 = _module.__default.m0((_710_v74)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length))], ((_710_v74)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length))]).multipliedBy(_this.f12), ((_710_v74)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length))]).plus((_710_v74)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_710_v74).length))]), true, globalState);
              _out42 = _outcollector12[0];
              _out43 = _outcollector12[1];
              _out44 = _outcollector12[2];
              _714_v76 = _out42;
              _715_v77 = _out43;
              _716_v78 = _out44;
            }
            let _717_v79;
            _717_v79 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f12),p2);
            let _718_v80;
            _718_v80 = _dafny.Seq.of(true);
            _717_v79 = (_717_v79).update(new BigNumber((_718_v80).length), _dafny.Seq.Concat(p2, p2));
            let _719_v81;
            let _nw127 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
            _719_v81 = _nw127;
            let _rhs110 = _719_v81;
            let _rhs111 = (_dafny.ZERO).minus(_module.__default.fm0(p0, _this.f12, globalState));
            let _lhs68 = _this;
            _719_v81 = _rhs110;
            _lhs68.f12 = _rhs111;
          }
        }
      }
      let _hi3 = _this.f12;
      for (let _720_i7 = (new BigNumber(-779)).multipliedBy(new BigNumber(357)); _720_i7.isLessThan(_hi3); _720_i7 = _720_i7.plus(_dafny.ONE)) {
        r0 = (_dafny.MultiSet.fromElements(p0)).IsSubsetOf(_688_v61);
        let _721_v82;
        let _nw128 = new _module.C0();
        _nw128.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_720_i7,p1), _720_i7);
        _721_v82 = _nw128;
        let _722_v83;
        _722_v83 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _723_v84;
        _723_v84 = _dafny.Seq.of(p1, (((_722_v83).contains(!(p1))) ? ((_722_v83).get(!(p1))) : (p1)), p1, p1, p1);
        let _724_v85;
        _724_v85 = _dafny.Seq.of(_723_v84);
        if ((_this.f12).isLessThanOrEqualTo(new BigNumber((_724_v85).length))) {
          let _725_v86;
          let _nw129 = Array((new BigNumber(22)).toNumber()).fill(false);
          _725_v86 = _nw129;
          let _index126 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_725_v86).length));
          (_725_v86)[_index126] = p1;
          let _index127 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_725_v86).length));
          (_725_v86)[_index127] = !_dafny.areEqual(p2, p2);
          let _726_v87;
          let _nw130 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _726_v87 = _nw130;
          let _index128 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_726_v87).length));
          (_726_v87)[_index128] = _720_i7;
          let _727_v88;
          _727_v88 = _dafny.Set.fromElements(false, (_725_v86)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_725_v86).length))]);
          let _728_v89;
          _728_v89 = _dafny.Map.Empty.slice().updateUnsafe(p1,_727_v88);
          let _index129 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_726_v87).length));
          (_726_v87)[_index129] = new BigNumber(((((_728_v89).contains((_this.f12).isEqualTo(_720_i7))) ? ((_728_v89).get((_this.f12).isEqualTo(_720_i7))) : (_727_v88))).length);
          let _index130 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_725_v86).length));
          (_725_v86)[_index130] = ((_module.D0.create_DC0(p1)).dtor_cf0) === ((_725_v86)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_725_v86).length))]);
          r1 = (_725_v86)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_725_v86).length))];
          let _729_v90;
          _729_v90 = _dafny.Seq.of(new BigNumber(796));
          let _730_v91;
          _730_v91 = _dafny.MultiSet.fromElements(_729_v90);
          let _index131 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_726_v87).length));
          let _index132 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_726_v87).length));
          let _rhs112 = _720_i7;
          let _rhs113 = (((_730_v91).contains(_dafny.Seq.Concat(_729_v90, _729_v90))) ? ((_730_v91).get(_dafny.Seq.Concat(_729_v90, _729_v90))) : (p0));
          let _rhs114 = _720_i7;
          let _lhs69 = _726_v87;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_726_v87).length));
          let _lhs71 = globalState;
          let _lhs72 = _726_v87;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_726_v87).length));
          _lhs69[_lhs70] = _rhs112;
          _lhs71.f5 = _rhs113;
          _lhs72[_lhs73] = _rhs114;
        } else {
          (globalState).f5 = (_dafny.ZERO).minus(_720_i7);
          let _731_v92;
          _731_v92 = _dafny.MultiSet.fromElements(p1, p1);
          _731_v92 = _dafny.MultiSet.fromElements(p1);
          let _732_v93;
          let _nw131 = Array((new BigNumber(22)).toNumber()).fill(false);
          _732_v93 = _nw131;
          let _index133 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_732_v93).length));
          (_732_v93)[_index133] = _module.__default.fm1(_this.f12, p1, _720_i7, false, globalState);
          let _index134 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_732_v93).length));
          (_732_v93)[_index134] = p1;
          let _733_v94;
          _733_v94 = new _dafny.CodePoint('j'.codePointAt(0));
          _733_v94 = (_this).f13;
          let _734_v95;
          let _735_v96;
          let _736_v97;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector13 = _module.__default.m0(_this.f12, (_dafny.ZERO).minus(_module.__default.fm0(p0, p0, globalState)), new BigNumber(406), (_732_v93)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_732_v93).length))], globalState);
          _out45 = _outcollector13[0];
          _out46 = _outcollector13[1];
          _out47 = _outcollector13[2];
          _734_v95 = _out45;
          _735_v96 = _out46;
          _736_v97 = _out47;
        }
        let _737_v98;
        _737_v98 = _dafny.Seq.UnicodeFromString("b");
        _737_v98 = _dafny.Seq.Concat(_737_v98, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(831)), function (_738_i8) {
          return (_this).f13;
        }), _dafny.Seq.UnicodeFromString("hxb")));
      }
      let _739_v99;
      let _init20 = ((_740_p1) => function (_741_i9) {
        return _740_p1;
      })(p1);
      let _nw132 = Array((new BigNumber(16)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw132.length); _i0_20++) {
        _nw132[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _739_v99 = _nw132;
      let _742_v100;
      _742_v100 = _module.D0.create_DC1(_739_v99);
      let _source9 = _742_v100;
      if (_source9.is_DC1) {
        let _743___mcc_h4 = (_source9).cf1;
        let _744_cf1 = _743___mcc_h4;
        let _index135 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_739_v99).length));
        (_739_v99)[_index135] = !(!(p1));
        let _index136 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_739_v99).length));
        (_739_v99)[_index136] = p1;
        let _745_v101;
        _745_v101 = _dafny.Seq.UnicodeFromString("uadjarwl");
        _745_v101 = _745_v101;
        let _746_v102;
        _746_v102 = _module.D2.create_DC9(p0);
        let _747_v103;
        _747_v103 = _module.__default.fm14(new BigNumber(-662), _746_v102, globalState);
        let _source10 = _747_v103;
        let _748___mcc_h11 = _source10;
        let _749_cf16 = _748___mcc_h11;
        let _750_v104;
        _750_v104 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),true);
        _750_v104 = (_750_v104).update(_module.__default.fm1(p0, p1, p0, p1, globalState), (_739_v99)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_739_v99).length))]);
        let _751_v105;
        _751_v105 = _dafny.Seq.of(_this.f12, p0, p0, (_this.f12).multipliedBy(p0));
        _751_v105 = _751_v105;
        _688_v61 = _688_v61;
        let _752_v106;
        let _nw133 = new _module.C1();
        _nw133.__ctor(new _dafny.CodePoint('l'.codePointAt(0)), (_this).f13, new BigNumber(738));
        _752_v106 = _nw133;
        _752_v106 = _752_v106;
        let _index137 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_739_v99).length));
        let _rhs115 = (new BigNumber((_745_v101).length)).multipliedBy(p0);
        let _rhs116 = p1;
        let _rhs117 = _module.__default.safeDivisionInt(_this.f12, new BigNumber(979));
        let _lhs74 = _739_v99;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_739_v99).length));
        let _lhs76 = _this;
        r2 = _rhs115;
        _lhs74[_lhs75] = _rhs116;
        _lhs76.f12 = _rhs117;
      } else if (_source9.is_DC2) {
        let _753___mcc_h5 = (_source9).cf2;
        let _754___mcc_h6 = (_source9).cf3;
        let _755___mcc_h7 = (_source9).cf4;
        let _756___mcc_h8 = (_source9).cf5;
        let _757_cf5 = _756___mcc_h8;
        let _758_cf4 = _755___mcc_h7;
        let _759_cf3 = _754___mcc_h6;
        let _760_cf2 = _753___mcc_h5;
        r0 = (_688_v61).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f12, p0, _this.f12)).length), (_dafny.ZERO).minus(new BigNumber((p2).length)))));
        let _761_v107;
        let _nw134 = new _module.C2();
        _nw134.__ctor(_760_cf2, new BigNumber(369));
        _761_v107 = _nw134;
        let _rhs118 = _761_v107;
        let _rhs119 = _760_cf2;
        _761_v107 = _rhs118;
        _760_cf2 = _rhs119;
        let _762_v108;
        let _nw135 = new _module.C1();
        _nw135.__ctor((_this).f13, (_this).f13, _this.f12);
        _762_v108 = _nw135;
        if (!(true)) {
          let _index138 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length));
          (_739_v99)[_index138] = _760_cf2;
          let _index139 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length));
          (_739_v99)[_index139] = !(true);
          (globalState).f5 = new BigNumber((p2).length);
          let _763_v109;
          _763_v109 = _dafny.Set.fromElements(_760_cf2, _module.__default.fm1(p0, _757_cf5, p0, (_739_v99)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length))], globalState));
          let _764_v110;
          _764_v110 = _763_v109;
          let _rhs120 = p0;
          let _rhs121 = _764_v110;
          let _lhs77 = _761_v107;
          _lhs77.f12 = _rhs120;
          _764_v110 = _rhs121;
          let _765_v111;
          let _nw136 = new _module.C0();
          _nw136.__ctor((_759_cf3).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),(_739_v99)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length))])), p0);
          _765_v111 = _nw136;
          let _766_v112;
          let _init21 = ((_767_cf5) => function (_768_i10) {
            return _767_cf5;
          })(_757_cf5);
          let _nw137 = Array((new BigNumber(28)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw137.length); _i0_21++) {
            _nw137[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _766_v112 = _nw137;
          let _769_v113;
          _769_v113 = _module.D6.create_DC16(_766_v112, _761_v107.f12);
          let _770_v115;
          _770_v115 = _dafny.Seq.of(false);
          let _771_v116;
          _771_v116 = _dafny.Map.Empty.slice().updateUnsafe((_739_v99)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length))],p0);
          let _772_v117;
          _772_v117 = _dafny.Seq.of(_761_v107.f12, new BigNumber(471), _this.f12);
          let _index140 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length));
          let _index141 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length));
          let _rhs122 = (_769_v113).dtor_cf23;
          let _rhs123 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(945), new BigNumber(-119))) {
              let _773_v114 = _compr_20;
              if (((new BigNumber(945)).isLessThanOrEqualTo(_773_v114)) && ((_773_v114).isLessThan(new BigNumber(-119)))) {
                _coll20.push([_module.__default.safeModuloInt(_773_v114, _761_v107.f12),_761_v107.f12]);
              }
            }
            return _coll20;
          }()).length)))).isLessThan(_761_v107.f12);
          let _rhs124 = ((!((_770_v115)[_module.__default.safeIndex(new BigNumber((_771_v116).length), new BigNumber((_770_v115).length))]) || (false)) ? (_760_cf2) : (_module.__default.fm1(new BigNumber((_772_v117).length), p1, _this.f12, _module.__default.fm1(new BigNumber(610), (_739_v99)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length))], _761_v107.f12, _760_cf2, globalState), globalState)));
          let _lhs78 = _761_v107;
          let _lhs79 = _739_v99;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length));
          let _lhs81 = _739_v99;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_739_v99).length));
          _lhs78.f12 = _rhs122;
          _lhs79[_lhs80] = _rhs123;
          _lhs81[_lhs82] = _rhs124;
        } else {
          let _774_v118;
          let _nw138 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _774_v118 = _nw138;
          let _775_v119;
          _775_v119 = _dafny.Set.fromElements(_759_cf3);
          let _rhs125 = _774_v118;
          let _rhs126 = _module.__default.safeModuloInt(new BigNumber((_775_v119).length), new BigNumber(808));
          let _lhs83 = _761_v107;
          _774_v118 = _rhs125;
          _lhs83.f12 = _rhs126;
          r2 = _761_v107.f12;
          r1 = !(!(_760_cf2));
          let _776_v120;
          _776_v120 = _dafny.Seq.UnicodeFromString("nbjki");
          _776_v120 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("c"), p2);
          let _777_v121;
          _777_v121 = _dafny.Seq.of(_761_v107.f12);
          let _778_v122;
          _778_v122 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_777_v121).length),(_762_v108).f15);
          let _779_v123;
          _779_v123 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),(_this).f13), _778_v122, (_778_v122).update(_this.f12, new _dafny.CodePoint('h'.codePointAt(0))), (_778_v122).Merge(_778_v122), (_778_v122).Merge(_778_v122));
          _779_v123 = _779_v123;
        }
      } else if (_source9.is_DC0) {
        let _780___mcc_h9 = (_source9).cf0;
        let _781_cf0 = _780___mcc_h9;
        let _782_v124;
        _782_v124 = _dafny.MultiSet.fromElements(_781_cf0);
        let _783_v125;
        _783_v125 = _dafny.Set.fromElements((_782_v124).IsSubsetOf(_782_v124));
        _783_v125 = _783_v125;
        let _index142 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
        (_739_v99)[_index142] = _781_cf0;
        let _index143 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
        (_739_v99)[_index143] = _781_cf0;
        let _784_v126;
        let _nw139 = Array((new BigNumber(19)).toNumber()).fill(false);
        _784_v126 = _nw139;
        let _785_v127;
        _785_v127 = _module.D6.create_DC16(_784_v126, p0);
        let _source11 = _785_v127;
        if (_source11.is_DC16) {
          let _786___mcc_h12 = (_source11).cf22;
          let _787___mcc_h13 = (_source11).cf23;
          let _788_cf23 = _787___mcc_h13;
          let _789_cf22 = _786___mcc_h12;
          let _790_v128;
          let _nw140 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _790_v128 = _nw140;
          let _index144 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_790_v128).length));
          (_790_v128)[_index144] = _788_cf23;
          let _791_v129;
          _791_v129 = _dafny.Seq.of(_789_cf22);
          let _index145 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_790_v128).length));
          (_790_v128)[_index145] = (new BigNumber((_dafny.Seq.Concat(_791_v129, _791_v129)).length)).multipliedBy(p0);
          let _792_v130;
          _792_v130 = _dafny.Map.Empty.slice().updateUnsafe((_790_v128)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_790_v128).length))],_module.__default.safeDivisionInt(p0, p0));
          let _index146 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_790_v128).length));
          (_790_v128)[_index146] = new BigNumber((_792_v130).length);
          let _793_v131;
          _793_v131 = _module.D2.create_DC8(_781_cf0, new BigNumber(288), _781_cf0);
          (_this).f12 = _module.__default.fm0(_module.__default.safeModuloInt((_790_v128)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_790_v128).length))], (_793_v131).dtor_cf12), new BigNumber(37), globalState);
          let _index147 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
          (_739_v99)[_index147] = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_789_cf22, _784_v126), _791_v129);
        } else if (_source11.is_DC17) {
          let _794___mcc_h14 = (_source11).cf24;
          let _795_cf24 = _794___mcc_h14;
          r0 = (_this.f12).isLessThan(_this.f12);
          let _796_v132;
          let _nw141 = new _module.C2();
          _nw141.__ctor(p1, _this.f12);
          _796_v132 = _nw141;
          let _797_v133;
          _797_v133 = _dafny.Map.Empty.slice().updateUnsafe(_796_v132,new BigNumber(919));
          let _798_v135;
          _798_v135 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13);
          let _799_v136;
          let _nw142 = new _module.C2();
          _nw142.__ctor(false, new BigNumber((function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of (_798_v135).Elements) {
              let _800_v134 = _compr_21;
              if ((_798_v135).contains(_800_v134)) {
                _coll21.push([_800_v134,p0]);
              }
            }
            return _coll21;
          }()).length));
          _799_v136 = _nw142;
          _797_v133 = (_797_v133).update(_796_v132, _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-507),_799_v136)).length), p0));
          let _801_v137;
          _801_v137 = _dafny.Seq.of(p1);
          let _802_v138;
          let _803_v139;
          let _804_v140;
          let _out48;
          let _out49;
          let _out50;
          let _outcollector14 = (_this).m3(((_781_cf0) ? (_dafny.MultiSet.FromArray(_801_v137)) : (_782_v124)), globalState);
          _out48 = _outcollector14[0];
          _out49 = _outcollector14[1];
          _out50 = _outcollector14[2];
          _802_v138 = _out48;
          _803_v139 = _out49;
          _804_v140 = _out50;
          let _805_v141;
          _805_v141 = _dafny.Seq.of(p0, _this.f12);
          let _806_v142;
          _806_v142 = _dafny.MultiSet.fromElements(_805_v141, _805_v141, _805_v141);
          let _index148 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
          (_739_v99)[_index148] = ((_806_v142).Intersect(_806_v142)).equals(_806_v142);
        } else if (_source11.is_DC18) {
          let _807___mcc_h15 = (_source11).cf25;
          let _808_cf25 = _807___mcc_h15;
          _808_cf25 = (_dafny.ZERO).minus(_808_cf25);
          let _809_v143;
          _809_v143 = _dafny.Seq.UnicodeFromString("r");
          _809_v143 = _dafny.Seq.UnicodeFromString("kj");
          let _index149 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
          (_739_v99)[_index149] = _781_cf0;
          let _810_v144;
          _810_v144 = new _dafny.CodePoint('l'.codePointAt(0));
          let _811_v145;
          _811_v145 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_781_cf0);
          let _812_v146;
          _812_v146 = _dafny.Map.Empty.slice().updateUnsafe(_781_cf0,_808_cf25);
          let _813_v147;
          let _nw143 = new _module.C0();
          _nw143.__ctor(_811_v145, (new BigNumber(95)).multipliedBy(new BigNumber((_812_v146).length)));
          _813_v147 = _nw143;
          let _814_v148;
          _814_v148 = _module.D6.create_DC18(_module.__default.safeModuloInt(_this.f12, (((_782_v124).contains((_739_v99)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length))])) ? ((_782_v124).get((_739_v99)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length))])) : (_808_cf25))));
          let _rhs127 = (((_688_v61).contains(new BigNumber(900))) ? ((_688_v61).get(new BigNumber(900))) : (_this.f12));
          let _rhs128 = _784_v126;
          let _rhs129 = _810_v144;
          let _rhs130 = _813_v147;
          let _rhs131 = _814_v148;
          let _lhs84 = _this;
          _lhs84.f12 = _rhs127;
          _784_v126 = _rhs128;
          _810_v144 = _rhs129;
          _813_v147 = _rhs130;
          _814_v148 = _rhs131;
        } else if (_source11.is_DC15) {
          let _815___mcc_h16 = (_source11).cf21;
          let _816_cf21 = _815___mcc_h16;
          let _817_v149;
          _817_v149 = _dafny.MultiSet.fromElements(p2, p2);
          _781_cf0 = (_817_v149).IsDisjointFrom((_817_v149).Difference(_817_v149));
          let _818_v150;
          let _nw144 = new _module.C1();
          _nw144.__ctor((_this).f13, (_this).f13, new BigNumber(-857));
          _818_v150 = _nw144;
          let _819_v151;
          _819_v151 = _module.D6.create_DC15(_dafny.Seq.Concat(_816_cf21, _816_cf21));
          _819_v151 = _module.D6.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(872)), ((_820_cf21) => function (_821_i11) {
  return new BigNumber((_820_cf21).length);
})(_816_cf21)));
          let _822_v152;
          let _init22 = ((_823_v99) => function (_824_i12) {
            return _dafny.Seq.of((_823_v99)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_823_v99).length))]);
          })(_739_v99);
          let _nw145 = Array((new BigNumber(11)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw145.length); _i0_22++) {
            _nw145[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _822_v152 = _nw145;
          let _825_v153;
          _825_v153 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f12),_this.f12);
          let _826_v154;
          _826_v154 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_825_v153).length)));
          let _index150 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
          let _rhs132 = (_826_v154).IsSubsetOf(_826_v154);
          let _rhs133 = (_739_v99)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length))];
          let _rhs134 = (_dafny.ZERO).minus(p0);
          let _rhs135 = _822_v152;
          let _rhs136 = false;
          let _lhs85 = _739_v99;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
          let _lhs87 = globalState;
          r0 = _rhs132;
          _lhs85[_lhs86] = _rhs133;
          _lhs87.f5 = _rhs134;
          _822_v152 = _rhs135;
          _781_cf0 = _rhs136;
        } else {
          let _827___mcc_h17 = (_source11).cf26;
          let _828_cf26 = _827___mcc_h17;
          (globalState).f5 = p0;
          let _829_v155;
          _829_v155 = _dafny.Map.Empty.slice().updateUnsafe(_782_v124,_this.f12);
          let _830_v156;
          _830_v156 = _module.D2.create_DC8(_781_cf0, (_dafny.ZERO).minus(p0), p1);
          _829_v155 = (_829_v155).update(_module.__default.fm19((_830_v156).dtor_cf12, globalState), p0);
          let _index151 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length));
          (_739_v99)[_index151] = _781_cf0;
          let _831_v157;
          let _nw146 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _831_v157 = _nw146;
          let _832_v158;
          _832_v158 = _dafny.Map.Empty.slice().updateUnsafe((_739_v99)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_739_v99).length))],_831_v157);
          let _833_v159;
          let _nw147 = Array((new BigNumber(5)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = _831_v157;
          _nw147[(_dafny.ONE).toNumber()] = _831_v157;
          _nw147[(new BigNumber(2)).toNumber()] = _831_v157;
          _nw147[(new BigNumber(3)).toNumber()] = (((_832_v158).contains(_781_cf0)) ? ((_832_v158).get(_781_cf0)) : (_831_v157));
          _nw147[(new BigNumber(4)).toNumber()] = _831_v157;
          _833_v159 = _nw147;
          let _index152 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_833_v159).length));
          (_833_v159)[_index152] = _831_v157;
          let _index153 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_833_v159).length));
          (_833_v159)[_index153] = _831_v157;
        }
        r1 = !(p1);
      } else {
        let _834___mcc_h10 = (_source9).cf6;
        let _835_cf6 = _834___mcc_h10;
        let _836_v160;
        let _nw148 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _836_v160 = _nw148;
        let _837_v161;
        _837_v161 = _dafny.Set.fromElements(_836_v160);
        let _838_v162;
        _838_v162 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_837_v161).length));
        let _839_v163;
        _839_v163 = _dafny.Map.Empty.slice().updateUnsafe(p0,_838_v162);
        _839_v163 = (_839_v163).update(new BigNumber((p2).length), _838_v162);
        if (p1) {
          let _840_v164;
          _840_v164 = _dafny.Set.fromElements(p1);
          let _841_v165;
          _841_v165 = _module.D6.create_DC17(p1);
          let _842_v166;
          _842_v166 = _dafny.Seq.of((_840_v164).IsSubsetOf(_840_v164), ((p1) ? (p1) : (p1)), (new BigNumber(188)).isLessThan(p0), true, (_841_v165).dtor_cf24);
          r1 = (_842_v166)[_module.__default.safeIndex(_this.f12, new BigNumber((_842_v166).length))];
          _836_v160 = _836_v160;
          let _843_v167;
          let _nw149 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _843_v167 = _nw149;
          let _844_v168;
          _844_v168 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _index154 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_843_v167).length));
          (_843_v167)[_index154] = _844_v168;
          let _index155 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_843_v167).length));
          (_843_v167)[_index155] = (_844_v168).Merge((_module.__default.fm18(_this.f12, _this.f12, p1, globalState)).Merge(_844_v168));
          r2 = ((_dafny.ZERO).minus(p0)).multipliedBy(_this.f12);
          let _845_v169;
          _845_v169 = _dafny.MultiSet.fromElements(p1, true);
          let _846_v170;
          let _847_v171;
          let _848_v172;
          let _out51;
          let _out52;
          let _out53;
          let _outcollector15 = (_this).m3((_845_v169).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(p1), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(p1)).length)), p1))), globalState);
          _out51 = _outcollector15[0];
          _out52 = _outcollector15[1];
          _out53 = _outcollector15[2];
          _846_v170 = _out51;
          _847_v171 = _out52;
          _848_v172 = _out53;
        } else {
          let _index156 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_836_v160).length));
          (_836_v160)[_index156] = new BigNumber((_dafny.Seq.UnicodeFromString("va")).length);
          let _index157 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_836_v160).length));
          (_836_v160)[_index157] = new BigNumber((p2).length);
          let _index158 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_836_v160).length));
          (_836_v160)[_index158] = _this.f12;
          let _849_v173;
          _849_v173 = _dafny.Seq.of(p0);
          let _850_v174;
          _850_v174 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((((_688_v61).contains(_this.f12)) ? ((_688_v61).get(_this.f12)) : ((_849_v173)[_module.__default.safeIndex(new BigNumber((p2).length), new BigNumber((_849_v173).length))])),new BigNumber((p2).length)),(_836_v160)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_836_v160).length))]);
          _850_v174 = ((p1) ? (_850_v174) : (_850_v174));
          let _851_v175;
          let _nw150 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _851_v175 = _nw150;
          let _index159 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_851_v175).length));
          (_851_v175)[_index159] = _dafny.Seq.Concat(p2, p2);
          let _index160 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_851_v175).length));
          (_851_v175)[_index160] = _dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("kemlgdi"));
          let _852_v176;
          _852_v176 = _dafny.Set.fromElements(p1);
          let _853_v177;
          _853_v177 = _dafny.Seq.of(_dafny.Set.fromElements(p1, !(true)), _852_v176, _852_v176);
          let _854_v178;
          _854_v178 = _dafny.Set.fromElements((_853_v177)[_module.__default.safeIndex((_836_v160)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_836_v160).length))], new BigNumber((_853_v177).length))], _852_v176, _852_v176);
          let _855_v179;
          _855_v179 = _module.D6.create_DC17(true);
          let _856_v180;
          _856_v180 = _dafny.Seq.of(_739_v99, _739_v99, _739_v99);
          let _rhs137 = _854_v178;
          let _rhs138 = _855_v179;
          let _rhs139 = _dafny.Seq.update(_dafny.Seq.Concat(_856_v180, _856_v180), _module.__default.safeIndex((_849_v173)[_module.__default.safeIndex(_this.f12, new BigNumber((_849_v173).length))], new BigNumber((_dafny.Seq.Concat(_856_v180, _856_v180)).length)), _739_v99);
          let _rhs140 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f12)), _this.f12);
          let _rhs141 = new BigNumber(-338);
          let _lhs88 = globalState;
          let _lhs89 = globalState;
          _854_v178 = _rhs137;
          _855_v179 = _rhs138;
          _856_v180 = _rhs139;
          _lhs88.f5 = _rhs140;
          _lhs89.f5 = _rhs141;
        }
        let _857_v181;
        let _init23 = ((_858_p2) => function (_859_i13) {
          return _858_p2;
        })(p2);
        let _nw151 = Array((new BigNumber(3)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw151.length); _i0_23++) {
          _nw151[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _857_v181 = _nw151;
        _857_v181 = ((p1) ? (_857_v181) : (_857_v181));
        let _860_v182;
        _860_v182 = _dafny.Seq.of(p1, p1);
        let _861_v183;
        _861_v183 = _dafny.Set.fromElements(p1);
        let _862_v184;
        _862_v184 = _dafny.Seq.of(p0, p0);
        _860_v182 = _module.__default.fm9(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(new BigNumber(((_838_v162).update(_this.f12, new BigNumber(105))).length), p0, (_dafny.ZERO).minus(new BigNumber((_861_v183).length))), _862_v184), globalState);
      }
      let _863_v185;
      _863_v185 = _dafny.MultiSet.fromElements(p1);
      let _864_v186;
      _864_v186 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),_863_v185);
      let _865_v187;
      _865_v187 = _dafny.Set.fromElements(true, p1);
      r0 = !(_module.__default.fm1(_this.f12, (_863_v185).IsSubsetOf((((_864_v186).contains(p1)) ? ((_864_v186).get(p1)) : (_863_v185))), new BigNumber((_865_v187).length), false, globalState));
      r0 = (p0).isLessThan(new BigNumber((_688_v61).cardinality()));
      r1 = !((p1) === (p1)) || (!(p1) || (p1));
      r2 = _module.__default.safeDivisionInt(p0, _this.f12);
      return [r0, r1, r2];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this.f11 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f10, f11) {
      let _this = this;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return (_this).f10;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = _dafny.Set.Empty;
      let r3 = [];
      let _866_v1;
      _866_v1 = _dafny.Set.fromElements(new BigNumber(999), (_this).f10, new BigNumber((_dafny.Seq.UnicodeFromString("stclb")).length), (_dafny.ZERO).minus((_this).f10), _this.f11);
      let _867_v2;
      _867_v2 = _module.D0.create_DC2(p3, function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of (_866_v1).Elements) {
    let _868_v0 = _compr_22;
    if ((_866_v1).contains(_868_v0)) {
      _coll22.push([(_868_v0).multipliedBy((_this).f10),true]);
    }
  }
  return _coll22;
}(), (_dafny.Map.Empty.slice().updateUnsafe(p2,p1)).update(p3, p1), p2);
      let _source12 = _867_v2;
      if (_source12.is_DC1) {
        let _869___mcc_h0 = (_source12).cf1;
        let _870_cf1 = _869___mcc_h0;
        r1 = _870_cf1;
        if (((p3) ? (p3) : (p2))) {
          let _871_v3;
          _871_v3 = _dafny.MultiSet.fromElements(_this.f11);
          let _872_v4;
          _872_v4 = _module.D2.create_DC8(p3, new BigNumber((_871_v3).cardinality()), p2);
          let _873_v5;
          _873_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p3);
          let _874_v6;
          _874_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f11);
          let _875_v7;
          _875_v7 = _dafny.Seq.of(_this.f11);
          let _876_v8;
          _876_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,p1);
          let _877_v9;
          let _nw152 = Array((new BigNumber(27)).toNumber());
          _nw152[(_dafny.ZERO).toNumber()] = _this.f11;
          _nw152[(_dafny.ONE).toNumber()] = _module.__default.fm0((_this).f10, (_this).f10, globalState);
          _nw152[(new BigNumber(2)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((p0).length), new BigNumber((p0).length));
          _nw152[(new BigNumber(4)).toNumber()] = new BigNumber(528);
          _nw152[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_872_v4).dtor_cf12);
          _nw152[(new BigNumber(7)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(8)).toNumber()] = new BigNumber((((p2) ? (_873_v5) : (_module.__default.fm3(p2, p3, new BigNumber((_874_v6).length), globalState)))).length);
          _nw152[(new BigNumber(9)).toNumber()] = _this.f11;
          _nw152[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(11)).toNumber()] = new BigNumber(171);
          _nw152[(new BigNumber(12)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(13)).toNumber()] = new BigNumber(-520);
          _nw152[(new BigNumber(14)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(15)).toNumber()] = (_875_v7)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_875_v7).length))];
          _nw152[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("tup"))).length);
          _nw152[(new BigNumber(17)).toNumber()] = new BigNumber((_876_v8).length);
          _nw152[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(106), _this.f11);
          _nw152[(new BigNumber(19)).toNumber()] = ((_this).f10).minus(_this.f11);
          _nw152[(new BigNumber(20)).toNumber()] = _this.f11;
          _nw152[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(_this.f11, _this.f11);
          _nw152[(new BigNumber(22)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(23)).toNumber()] = ((p3) ? (_this.f11) : (_this.f11));
          _nw152[(new BigNumber(24)).toNumber()] = (_this).f10;
          _nw152[(new BigNumber(25)).toNumber()] = _this.f11;
          _nw152[(new BigNumber(26)).toNumber()] = _this.f11;
          _877_v9 = _nw152;
          let _878_v10;
          _878_v10 = _module.D2.create_DC9(new BigNumber(779));
          let _index161 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_877_v9).length));
          (_877_v9)[_index161] = (_878_v10).dtor_cf14;
          let _index162 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_877_v9).length));
          (_877_v9)[_index162] = _this.f11;
          let _879_v11;
          _879_v11 = true;
          _879_v11 = false;
          let _880_v12;
          _880_v12 = _dafny.Set.fromElements(_871_v3, _871_v3, _dafny.MultiSet.fromElements((_this).f10));
          let _rhs142 = _880_v12;
          let _rhs143 = false;
          let _rhs144 = _module.__default.safeModuloInt((_this).f10, _this.f11);
          let _rhs145 = !(p2) || (_879_v11);
          _880_v12 = _rhs142;
          _879_v11 = _rhs143;
          r0 = _rhs144;
          _879_v11 = _rhs145;
          let _881_v13;
          _881_v13 = _dafny.Set.fromElements(p2);
          let _882_v14;
          let _nw153 = new _module.C2();
          _nw153.__ctor((_881_v13).IsSubsetOf(_881_v13), (_877_v9)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_877_v9).length))]);
          _882_v14 = _nw153;
          (globalState).f5 = (_877_v9)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_877_v9).length))];
        } else {
          let _883_v15;
          _883_v15 = false;
          let _884_v16;
          _884_v16 = _dafny.MultiSet.fromElements(_870_cf1, _870_cf1);
          _883_v15 = !((_884_v16).IsDisjointFrom(_dafny.MultiSet.fromElements(_870_cf1, _870_cf1, _870_cf1))) || (p2);
          let _885_v17;
          let _nw154 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _885_v17 = _nw154;
          let _index163 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_885_v17).length));
          (_885_v17)[_index163] = _this.f11;
          let _index164 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_885_v17).length));
          (_885_v17)[_index164] = _this.f11;
          _883_v15 = (((_883_v15) ? (_this.f11) : ((_this).f10))).isLessThan(new BigNumber((p0).length));
          (globalState).f5 = _module.__default.safeModuloInt((_885_v17)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_885_v17).length))], _this.f11);
          _883_v15 = (p2) || (!_dafny.areEqual(p0, p0));
        }
        let _886_v18;
        _886_v18 = _module.D5.create_DC13(true, false);
        let _887_v19;
        _887_v19 = _dafny.MultiSet.fromElements(p3, (_886_v18).dtor_cf19);
        let _888_v20;
        _888_v20 = _dafny.Seq.of(new BigNumber(681), _this.f11);
        let _889_v21;
        _889_v21 = _dafny.Seq.of(_887_v19, _887_v19, (_887_v19).update(p2, _module.__default.abs(new BigNumber((_888_v20).length))), _887_v19, _887_v19);
        let _source13 = _module.D5.create_DC12(_889_v21);
        if (_source13.is_DC13) {
          let _890___mcc_h7 = (_source13).cf18;
          let _891___mcc_h8 = (_source13).cf19;
          let _892_cf19 = _891___mcc_h8;
          let _893_cf18 = _890___mcc_h7;
          let _894_v22;
          let _nw155 = new _module.C1();
          _nw155.__ctor(p1, p1, (_this).f10);
          _894_v22 = _nw155;
          let _rhs146 = _894_v22;
          let _rhs147 = _892_cf19;
          _894_v22 = _rhs146;
          _893_cf18 = _rhs147;
          (globalState).f5 = new BigNumber(115);
          let _895_v23;
          _895_v23 = _dafny.MultiSet.fromElements(_this.f11, new BigNumber((p0).length));
          let _896_v24;
          _896_v24 = _module.D5.create_DC12(_dafny.Seq.of(_887_v19, _887_v19));
          (_this).f11 = new BigNumber((_module.__default.fm20((_this).fm2(_893_cf18, globalState), _this.f11, new BigNumber((_895_v23).cardinality()), _896_v24, globalState)).length);
          (globalState).f5 = new BigNumber((_887_v19).cardinality());
        } else if (_source13.is_DC14) {
          let _897___mcc_h9 = (_source13).cf20;
          let _898_cf20 = _897___mcc_h9;
          let _899_v25;
          let _nw156 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _899_v25 = _nw156;
          let _index165 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_899_v25).length));
          (_899_v25)[_index165] = p0;
          let _index166 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_899_v25).length));
          (_899_v25)[_index166] = p0;
          let _900_v26;
          let _nw157 = new _module.C1();
          _nw157.__ctor(p1, p1, new BigNumber(((_899_v25)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_899_v25).length))]).length));
          _900_v26 = _nw157;
          _900_v26 = ((false) ? (_900_v26) : (_900_v26));
          let _901_v27;
          let _nw158 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _901_v27 = _nw158;
          r3 = _901_v27;
          let _902_v28;
          _902_v28 = _dafny.MultiSet.fromElements(p0, _dafny.Seq.UnicodeFromString("ed"));
          let _903_v29;
          _903_v29 = _dafny.Map.Empty.slice().updateUnsafe(_870_cf1,!(_902_v28).contains(p0));
          _903_v29 = (_903_v29).update(_870_cf1, _module.__default.fm1((_this).f10, !(_898_cf20), (_this).f10, p3, globalState));
        } else {
          let _904___mcc_h10 = (_source13).cf17;
          let _905_cf17 = _904___mcc_h10;
          let _906_v30;
          let _nw159 = new _module.C1();
          _nw159.__ctor(p1, p1, _this.f11);
          _906_v30 = _nw159;
          r0 = _this.f11;
          let _907_v31;
          _907_v31 = true;
          _907_v31 = p2;
          let _908_v32;
          let _nw160 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _908_v32 = _nw160;
          let _index167 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_908_v32).length));
          (_908_v32)[_index167] = p0;
          let _index168 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_908_v32).length));
          (_908_v32)[_index168] = _dafny.Seq.Concat(p0, p0);
        }
        let _909_v33;
        let _nw161 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _909_v33 = _nw161;
        let _index169 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_909_v33).length));
        (_909_v33)[_index169] = (_this).f10;
        let _index170 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_909_v33).length));
        (_909_v33)[_index170] = _this.f11;
      } else if (_source12.is_DC2) {
        let _910___mcc_h1 = (_source12).cf2;
        let _911___mcc_h2 = (_source12).cf3;
        let _912___mcc_h3 = (_source12).cf4;
        let _913___mcc_h4 = (_source12).cf5;
        let _914_cf5 = _913___mcc_h4;
        let _915_cf4 = _912___mcc_h3;
        let _916_cf3 = _911___mcc_h2;
        let _917_cf2 = _910___mcc_h1;
        let _918_v34;
        _918_v34 = _module.D7.create_DC21();
        _918_v34 = _module.D7.create_DC21();
        if (false) {
          _914_cf5 = _914_cf5;
          let _919_v35;
          _919_v35 = _dafny.Seq.of(_this.f11, _this.f11, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,(_module.D2.create_DC7((_this).f10)).dtor_cf10)).length), new BigNumber(533));
          let _920_v36;
          let _nw162 = new _module.C2();
          _nw162.__ctor(p3, (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_919_v35).length), (_dafny.ZERO).minus(_this.f11))));
          _920_v36 = _nw162;
          let _921_v37;
          _921_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_920_v36).f17);
          let _922_v38;
          _922_v38 = _dafny.MultiSet.fromElements((((_921_v37).contains(new _dafny.CodePoint('a'.codePointAt(0)))) ? ((_921_v37).get(new _dafny.CodePoint('a'.codePointAt(0)))) : (!(p2))), _914_cf5);
          (globalState).f5 = (((_922_v38).contains(((_this).f10).isLessThan(_this.f11))) ? ((_922_v38).get(((_this).f10).isLessThan(_this.f11))) : (new BigNumber((_916_cf3).length)));
          let _923_v39;
          let _nw163 = Array((new BigNumber(28)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = _914_cf5;
          _nw163[(_dafny.ONE).toNumber()] = p3;
          _nw163[(new BigNumber(2)).toNumber()] = _917_cf2;
          _nw163[(new BigNumber(3)).toNumber()] = (_920_v36).f17;
          _nw163[(new BigNumber(4)).toNumber()] = p3;
          _nw163[(new BigNumber(5)).toNumber()] = _914_cf5;
          _nw163[(new BigNumber(6)).toNumber()] = (_920_v36).f17;
          _nw163[(new BigNumber(7)).toNumber()] = _914_cf5;
          _nw163[(new BigNumber(8)).toNumber()] = p3;
          _nw163[(new BigNumber(9)).toNumber()] = (_920_v36).f17;
          _nw163[(new BigNumber(10)).toNumber()] = (_920_v36).f17;
          _nw163[(new BigNumber(11)).toNumber()] = _917_cf2;
          _nw163[(new BigNumber(12)).toNumber()] = _914_cf5;
          _nw163[(new BigNumber(13)).toNumber()] = p2;
          _nw163[(new BigNumber(14)).toNumber()] = true;
          _nw163[(new BigNumber(15)).toNumber()] = p2;
          _nw163[(new BigNumber(16)).toNumber()] = _917_cf2;
          _nw163[(new BigNumber(17)).toNumber()] = _914_cf5;
          _nw163[(new BigNumber(18)).toNumber()] = _914_cf5;
          _nw163[(new BigNumber(19)).toNumber()] = p3;
          _nw163[(new BigNumber(20)).toNumber()] = true;
          _nw163[(new BigNumber(21)).toNumber()] = p2;
          _nw163[(new BigNumber(22)).toNumber()] = p3;
          _nw163[(new BigNumber(23)).toNumber()] = false;
          _nw163[(new BigNumber(24)).toNumber()] = p3;
          _nw163[(new BigNumber(25)).toNumber()] = p2;
          _nw163[(new BigNumber(26)).toNumber()] = p2;
          _nw163[(new BigNumber(27)).toNumber()] = p2;
          _923_v39 = _nw163;
          let _924_v40;
          let _nw164 = new _module.C1();
          _nw164.__ctor(p1, new _dafny.CodePoint('l'.codePointAt(0)), (_module.D6.create_DC16(_923_v39, (_this).f10)).dtor_cf23);
          _924_v40 = _nw164;
          let _925_v41;
          let _init24 = ((_926_p0) => function (_927_i0) {
            return _926_p0;
          })(p0);
          let _nw165 = Array((new BigNumber(6)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw165.length); _i0_24++) {
            _nw165[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _925_v41 = _nw165;
          let _928_v42;
          _928_v42 = _dafny.Seq.of(_924_v40, _924_v40);
          let _rhs148 = (_928_v42)[_module.__default.safeIndex(new BigNumber(-414), new BigNumber((_928_v42).length))];
          let _rhs149 = !_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), ((_929_p0) => function (_930_i2) {
            return _929_p0;
          })(p0)), _dafny.Seq.of(p0, p0)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_931_v40) => function (_932_i1) {
            return (_931_v40).f13;
          })(_924_v40)), p0));
          let _rhs150 = _925_v41;
          let _rhs151 = (_dafny.ZERO).minus((_this).f10);
          let _lhs90 = _924_v40;
          _924_v40 = _rhs148;
          _914_cf5 = _rhs149;
          _925_v41 = _rhs150;
          _lhs90.f12 = _rhs151;
          let _933_v43;
          _933_v43 = _dafny.MultiSet.fromElements(_module.__default.fm11(_924_v40.f12, _924_v40.f12, p2, globalState));
          let _934_v44;
          let _nw166 = Array((new BigNumber(23)).toNumber()).fill(_module.D2.Default());
          _934_v44 = _nw166;
          let _935_v45;
          _935_v45 = _dafny.Set.fromElements((_924_v40).f13);
          let _936_v46;
          _936_v46 = _dafny.Map.Empty.slice().updateUnsafe(_917_cf2,_module.__default.fm0(new BigNumber((_935_v45).length), _this.f11, globalState));
          let _937_v47;
          let _nw167 = new _module.C3();
          _nw167.__ctor(_934_v44, p1, (((_936_v46).contains(p2)) ? ((_936_v46).get(p2)) : (_this.f11)));
          _937_v47 = _nw167;
          let _938_v48;
          _938_v48 = _dafny.Map.Empty.slice().updateUnsafe((_module.D8.create_DC22(_933_v43)).dtor_cf28,_937_v47);
          _938_v48 = (_938_v48).update(_dafny.MultiSet.fromElements(p0, p0, p0, p0, p0), _937_v47);
        } else {
          let _939_v49;
          let _nw168 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
          _939_v49 = _nw168;
          let _940_v50;
          _940_v50 = _dafny.Set.fromElements(p2, _914_cf5);
          let _941_v51;
          _941_v51 = _dafny.MultiSet.fromElements(new BigNumber((_940_v50).length));
          let _index171 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_939_v49).length));
          (_939_v49)[_index171] = _941_v51;
          let _index172 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_939_v49).length));
          (_939_v49)[_index172] = _941_v51;
          let _942_v52;
          let _nw169 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _942_v52 = _nw169;
          let _943_v53;
          _943_v53 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f11);
          let _index173 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_942_v52).length));
          (_942_v52)[_index173] = new BigNumber(((_943_v53).update(p2, (_dafny.ZERO).minus(_this.f11))).length);
          let _944_v54;
          _944_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),(_this).f10);
          let _index174 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_942_v52).length));
          (_942_v52)[_index174] = ((_this.f11).minus(new BigNumber((_944_v54).length))).multipliedBy(_this.f11);
          _917_cf2 = _917_cf2;
          _917_cf2 = _module.__default.fm1(_this.f11, p3, (_this).f10, _914_cf5, globalState);
          let _945_v55;
          _945_v55 = _dafny.Set.fromElements(_942_v52, _942_v52, _942_v52);
          let _946_v56;
          _946_v56 = _module.D5.create_DC13(true, !((_dafny.ZERO).minus(_this.f11)).isEqualTo(new BigNumber((_945_v55).length)));
          _946_v56 = _946_v56;
        }
        (globalState).f5 = new BigNumber((_module.__default.fm10(_module.__default.safeModuloInt(new BigNumber(-926), _this.f11), globalState)).length);
        let _947_v57;
        _947_v57 = _dafny.Seq.of(_917_cf2);
        let _948_v58;
        _948_v58 = _dafny.Map.Empty.slice().updateUnsafe(_947_v57,_dafny.Set.fromElements(p3));
        let _949_v59;
        _949_v59 = _dafny.Set.fromElements(p2, p3);
        _914_cf5 = (_949_v59).IsProperSubsetOf(((((_948_v58).contains(_947_v57)) ? ((_948_v58).get(_947_v57)) : (_949_v59))).Difference(_dafny.Set.fromElements(p3, p3)));
      } else if (_source12.is_DC0) {
        let _950___mcc_h5 = (_source12).cf0;
        let _951_cf0 = _950___mcc_h5;
        let _952_v60;
        _952_v60 = _dafny.MultiSet.fromElements(_866_v1);
        _951_cf0 = (_952_v60).IsProperSubsetOf(_952_v60);
        let _953_v61;
        let _nw170 = Array((new BigNumber(8)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = _this.f11;
        _nw170[(_dafny.ONE).toNumber()] = _this.f11;
        _nw170[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_this.f11);
        _nw170[(new BigNumber(3)).toNumber()] = new BigNumber((_866_v1).length);
        _nw170[(new BigNumber(4)).toNumber()] = (_this).f10;
        _nw170[(new BigNumber(5)).toNumber()] = _this.f11;
        _nw170[(new BigNumber(6)).toNumber()] = (_this).f10;
        _nw170[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11)).length);
        _953_v61 = _nw170;
        let _954_v62;
        let _nw171 = Array((new BigNumber(16)).toNumber());
        _nw171[(_dafny.ZERO).toNumber()] = _951_cf0;
        _nw171[(_dafny.ONE).toNumber()] = false;
        _nw171[(new BigNumber(2)).toNumber()] = p3;
        _nw171[(new BigNumber(3)).toNumber()] = p3;
        _nw171[(new BigNumber(4)).toNumber()] = p2;
        _nw171[(new BigNumber(5)).toNumber()] = p2;
        _nw171[(new BigNumber(6)).toNumber()] = p2;
        _nw171[(new BigNumber(7)).toNumber()] = p3;
        _nw171[(new BigNumber(8)).toNumber()] = p2;
        _nw171[(new BigNumber(9)).toNumber()] = _module.__default.fm1((_this).f10, p3, (_this).f10, p2, globalState);
        _nw171[(new BigNumber(10)).toNumber()] = p3;
        _nw171[(new BigNumber(11)).toNumber()] = p2;
        _nw171[(new BigNumber(12)).toNumber()] = _951_cf0;
        _nw171[(new BigNumber(13)).toNumber()] = _951_cf0;
        _nw171[(new BigNumber(14)).toNumber()] = p3;
        _nw171[(new BigNumber(15)).toNumber()] = _951_cf0;
        _954_v62 = _nw171;
        let _955_v63;
        _955_v63 = _dafny.Map.Empty.slice().updateUnsafe(_953_v61,_954_v62);
        let _index175 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
        (_954_v62)[_index175] = p2;
        let _index176 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
        let _rhs152 = ((_955_v63).update(_953_v61, _954_v62)).Merge(_955_v63);
        let _rhs153 = p2;
        let _lhs91 = _954_v62;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
        _955_v63 = _rhs152;
        _lhs91[_lhs92] = _rhs153;
        let _956_v64;
        _956_v64 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hs"), _dafny.Seq.UnicodeFromString("yn"));
        let _957_v65;
        _957_v65 = _module.D8.create_DC22(_956_v64);
        let _source14 = _957_v65;
        if (_source14.is_DC23) {
          let _958___mcc_h11 = (_source14).cf29;
          let _959_cf29 = _958___mcc_h11;
          let _960_v66;
          _960_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,_953_v61);
          let _961_v67;
          let _nw172 = new _module.C1();
          _nw172.__ctor(p1, p1, _959_cf29);
          _961_v67 = _nw172;
          let _pat_let_tv6 = _960_v66;
          let _962_v68;
          _962_v68 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let13_0) {
            return function (_963_dt__update__tmp_h0) {
              return function (_pat_let14_0) {
                return function (_964_dt__update_hcf27_h0) {
                  return _module.D7.create_DC20(_964_dt__update_hcf27_h0);
                }(_pat_let14_0);
              }(_pat_let_tv6);
            }(_pat_let13_0);
          }(_module.D7.create_DC20(_dafny.Map.Empty.slice().updateUnsafe((_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))],_953_v61))),_961_v67);
          let _965_v69;
          _965_v69 = _dafny.Map.Empty.slice().updateUnsafe(_962_v68,!(p3));
          let _966_v70;
          _966_v70 = _dafny.Map.Empty.slice().updateUnsafe(_959_cf29,_965_v69);
          let _967_v71;
          _967_v71 = _module.D2.create_DC7(_961_v67.f12);
          let _968_v72;
          _968_v72 = _dafny.Map.Empty.slice().updateUnsafe((_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))],(_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))]);
          let _pat_let_tv7 = _959_cf29;
          let _969_v73;
          let _nw173 = Array((new BigNumber(15)).toNumber());
          _nw173[(_dafny.ZERO).toNumber()] = _967_v71;
          _nw173[(_dafny.ONE).toNumber()] = function (_pat_let15_0) {
            return function (_970_dt__update__tmp_h1) {
              return function (_pat_let16_0) {
                return function (_971_dt__update_hcf10_h0) {
                  return _module.D2.create_DC7(_971_dt__update_hcf10_h0);
                }(_pat_let16_0);
              }(_pat_let_tv7);
            }(_pat_let15_0);
          }(_967_v71);
          _nw173[(new BigNumber(2)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(3)).toNumber()] = _module.D2.create_DC7(new BigNumber((_968_v72).length));
          _nw173[(new BigNumber(4)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(5)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(6)).toNumber()] = _module.D2.create_DC7((_this).f10);
          _nw173[(new BigNumber(7)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(8)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(9)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(10)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(11)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(12)).toNumber()] = _module.D2.create_DC7(new BigNumber(539));
          _nw173[(new BigNumber(13)).toNumber()] = _967_v71;
          _nw173[(new BigNumber(14)).toNumber()] = _967_v71;
          _969_v73 = _nw173;
          let _972_v74;
          let _nw174 = new _module.C3();
          _nw174.__ctor(_969_v73, p1, new BigNumber(-596));
          _972_v74 = _nw174;
          let _973_v75;
          _973_v75 = _dafny.Seq.of(_972_v74);
          _966_v70 = (_966_v70).update(_module.__default.safeDivisionInt(_959_cf29, new BigNumber((_973_v75).length)), _965_v69);
          let _974_v76;
          _974_v76 = _module.D7.create_DC20(_960_v66);
          let _975_v77;
          _975_v77 = _dafny.Seq.of(_974_v76);
          _975_v77 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_975_v77, _975_v77), _975_v77), _module.__default.safeIndex(_961_v67.f12, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_975_v77, _975_v77), _975_v77)).length)), _974_v76);
          let _976_v78;
          let _init25 = ((_977_p3) => function (_978_i3) {
            return _dafny.Set.fromElements(_977_p3);
          })(p3);
          let _nw175 = Array((new BigNumber(29)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw175.length); _i0_25++) {
            _nw175[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _976_v78 = _nw175;
          let _979_v79;
          let _nw176 = new _module.C3();
          _nw176.__ctor((_972_v74).f14, p1, new BigNumber((_dafny.MultiSet.fromElements(_976_v78, _976_v78)).cardinality()));
          _979_v79 = _nw176;
          let _980_v80;
          _980_v80 = new _dafny.CodePoint('u'.codePointAt(0));
          let _981_v81;
          _981_v81 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_961_v67).f13);
          _980_v80 = (((_981_v81).contains((((_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))]) ? (p2) : (p2)))) ? ((_981_v81).get((((_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))]) ? (p2) : (p2)))) : (p1));
        } else if (_source14.is_DC24) {
          let _982___mcc_h12 = (_source14).cf30;
          let _983___mcc_h13 = (_source14).cf31;
          let _984_cf31 = _983___mcc_h13;
          let _985_cf30 = _982___mcc_h12;
          let _986_v82;
          _986_v82 = _module.D2.create_DC7(_this.f11);
          let _987_v83;
          let _nw177 = Array((new BigNumber(13)).toNumber());
          _nw177[(_dafny.ZERO).toNumber()] = _986_v82;
          _nw177[(_dafny.ONE).toNumber()] = _986_v82;
          _nw177[(new BigNumber(2)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(3)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(4)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(5)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(6)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(7)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(8)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(9)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(10)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(11)).toNumber()] = _986_v82;
          _nw177[(new BigNumber(12)).toNumber()] = _986_v82;
          _987_v83 = _nw177;
          let _988_v84;
          _988_v84 = _module.D9.create_DC25(p1);
          let _989_v85;
          let _nw178 = new _module.C3();
          _nw178.__ctor(_987_v83, (_988_v84).dtor_cf32, new BigNumber((_dafny.Set.fromElements(_951_cf0)).length));
          _989_v85 = _nw178;
          let _index177 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          (_954_v62)[_index177] = true;
          let _990_v86;
          _990_v86 = _dafny.Map.Empty.slice().updateUnsafe(true,!(!(p3)));
          let _991_v87;
          _991_v87 = _dafny.Seq.of(((!((_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))])) ? (new BigNumber((_990_v86).length)) : (_this.f11)));
          (globalState).f5 = (_991_v87)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f10)), new BigNumber((_991_v87).length))];
          let _992_v88;
          _992_v88 = _module.D6.create_DC18(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_951_cf0),p2)).length));
          let _index178 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          let _index179 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          let _index180 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          let _rhs154 = (_954_v62)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length))];
          let _rhs155 = (_992_v88).dtor_cf25;
          let _rhs156 = _module.__default.fm1((_this).f10, _951_cf0, (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f11, new BigNumber(250))), p2, globalState);
          let _rhs157 = p2;
          let _lhs93 = _954_v62;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          let _lhs95 = _this;
          let _lhs96 = _954_v62;
          let _lhs97 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          let _lhs98 = _954_v62;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_954_v62).length));
          _lhs93[_lhs94] = _rhs154;
          _lhs95.f11 = _rhs155;
          _lhs96[_lhs97] = _rhs156;
          _lhs98[_lhs99] = _rhs157;
        } else {
          let _993___mcc_h14 = (_source14).cf28;
          let _994_cf28 = _993___mcc_h14;
          let _995_v89;
          _995_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_this.f11);
          let _996_v90;
          _996_v90 = _dafny.Seq.of(new BigNumber(858));
          _995_v89 = (_995_v89).update(_this.f11, (_996_v90)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_996_v90).length))]);
          let _997_v91;
          let _nw179 = new _module.C1();
          _nw179.__ctor((p0)[_module.__default.safeIndex(_this.f11, new BigNumber((p0).length))], p1, (_this).f10);
          _997_v91 = _nw179;
          let _998_v92;
          let _init26 = ((_999_v91, _1000_p0) => function (_1001_i4) {
            return _module.D2.create_DC7(new BigNumber((_dafny.Set.fromElements(_999_v91.f12, (_this).f10, _999_v91.f12, new BigNumber((_1000_p0).length))).length));
          })(_997_v91, p0);
          let _nw180 = Array((new BigNumber(12)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw180.length); _i0_26++) {
            _nw180[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _998_v92 = _nw180;
          let _nw181 = new _module.C3();
          _nw181.__ctor(_998_v92, (_997_v91).f13, (_996_v90)[_module.__default.safeIndex((_this).f10, new BigNumber((_996_v90).length))]);
          _997_v91 = _nw181;
          let _1002_v93;
          _1002_v93 = new _dafny.CodePoint('n'.codePointAt(0));
          _1002_v93 = new _dafny.CodePoint('a'.codePointAt(0));
          _866_v1 = _866_v1;
        }
        _954_v62 = _954_v62;
      } else {
        let _1003___mcc_h6 = (_source12).cf6;
        let _1004_cf6 = _1003___mcc_h6;
        let _1005_v94;
        let _nw182 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1005_v94 = _nw182;
        let _1006_v95;
        _1006_v95 = _module.D0.create_DC1(_1005_v94);
        let _1007_v97;
        _1007_v97 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of (p0).Elements) {
            let _1008_v96 = _compr_23;
            if (_dafny.Seq.contains(p0, _1008_v96)) {
              _coll23.push([_1008_v96,p2]);
            }
          }
          return _coll23;
        }()).length))).length), new BigNumber(257));
        let _1009_v98;
        _1009_v98 = _dafny.Map.Empty.slice().updateUnsafe((_1006_v95).dtor_cf1,new BigNumber((_dafny.Seq.Concat(_1007_v97, _dafny.Seq.of(_this.f11))).length));
        let _index181 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_1005_v94).length));
        (_1005_v94)[_index181] = p3;
        let _1010_v99;
        _1010_v99 = _dafny.MultiSet.fromElements(p0, _dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(447)), ((_1011_p1) => function (_1012_i5) {
          return _1011_p1;
        })(p1))));
        let _1013_v100;
        _1013_v100 = _dafny.MultiSet.fromElements(_this.f11);
        let _1014_v101;
        _1014_v101 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_1007_v97), _1013_v100, _1013_v100);
        let _index182 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_1005_v94).length));
        let _rhs158 = _1009_v98;
        let _rhs159 = _this.f11;
        let _rhs160 = !(_1014_v101).equals((_1014_v101).Intersect(_1014_v101));
        let _rhs161 = _1010_v99;
        let _lhs100 = globalState;
        let _lhs101 = _1005_v94;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_1005_v94).length));
        _1009_v98 = _rhs158;
        _lhs100.f5 = _rhs159;
        _lhs101[_lhs102] = _rhs160;
        _1010_v99 = _rhs161;
        let _1015_v102;
        _1015_v102 = new _dafny.CodePoint('u'.codePointAt(0));
        _1015_v102 = _1015_v102;
        let _1016_v103;
        _1016_v103 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _1017_v104;
        _1017_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1016_v103,_this.f11);
        _1017_v104 = (_1017_v104).update(_1016_v103, (_this).f10);
        let _index183 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_1005_v94).length));
        (_1005_v94)[_index183] = p2;
      }
      let _1018_v105;
      let _nw183 = new _module.C1();
      _nw183.__ctor(p1, p1, new BigNumber(554));
      _1018_v105 = _nw183;
      let _1019_v106;
      _1019_v106 = _dafny.MultiSet.fromElements(new BigNumber(-902));
      let _1020_v107;
      _1020_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(146),_1019_v106);
      if (((_1019_v106).update(_this.f11, _module.__default.abs((_this).f10))).IsSubsetOf(((p3) ? (_1019_v106) : ((((_1020_v107).contains(_this.f11)) ? ((_1020_v107).get(_this.f11)) : (_1019_v106)))))) {
        let _1021_v108;
        _1021_v108 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f10);
        let _1022_v109;
        _1022_v109 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1021_v108).length));
        let _1023_v110;
        let _nw184 = new _module.C2();
        _nw184.__ctor(p3, (((_1022_v109).contains(p2)) ? ((_1022_v109).get(p2)) : (_this.f11)));
        _1023_v110 = _nw184;
        let _1024_v111;
        let _init27 = ((_1025_v110) => function (_1026_i6) {
          return (_1026_i6).minus(_1025_v110.f12);
        })(_1023_v110);
        let _nw185 = Array((new BigNumber(9)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw185.length); _i0_27++) {
          _nw185[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1024_v111 = _nw185;
        let _1027_v112;
        _1027_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1023_v110.f12,_1024_v111);
        let _1028_v113;
        _1028_v113 = _dafny.Seq.of(_1024_v111);
        let _rhs162 = new BigNumber((((_1019_v106).Intersect(_1019_v106)).Difference((_1019_v106).update((_this).f10, _module.__default.abs((_this).f10)))).cardinality());
        let _rhs163 = new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1024_v111)).Merge((_1027_v112).update(_this.f11, _1024_v111))).update(_this.f11, (_1028_v113)[_module.__default.safeIndex(_this.f11, new BigNumber((_1028_v113).length))])).length);
        let _rhs164 = _1023_v110;
        let _rhs165 = _this.f11;
        let _lhs103 = _this;
        let _lhs104 = globalState;
        let _lhs105 = _1023_v110;
        _lhs103.f11 = _rhs162;
        _lhs104.f5 = _rhs163;
        _1023_v110 = _rhs164;
        _lhs105.f12 = _rhs165;
        let _1029_v114;
        let _nw186 = Array((new BigNumber(25)).toNumber());
        _nw186[(_dafny.ZERO).toNumber()] = _1024_v111;
        _nw186[(_dafny.ONE).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(2)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(3)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(4)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(5)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(6)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(7)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(8)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(9)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(10)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(11)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(12)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(13)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(14)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(15)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(16)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(17)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(18)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(19)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(20)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(21)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(22)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(23)).toNumber()] = _1024_v111;
        _nw186[(new BigNumber(24)).toNumber()] = _1024_v111;
        _1029_v114 = _nw186;
        let _index184 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1029_v114).length));
        (_1029_v114)[_index184] = _1024_v111;
        let _1030_v115;
        _1030_v115 = _dafny.Seq.of(p3, !(true));
        let _1031_v116;
        let _init28 = ((_1032_p3) => function (_1033_i7) {
          return _1032_p3;
        })(p3);
        let _nw187 = Array((new BigNumber(17)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw187.length); _i0_28++) {
          _nw187[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1031_v116 = _nw187;
        let _1034_v117;
        _1034_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1023_v110.f12,_1031_v116);
        let _index185 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1029_v114).length));
        let _nw188 = Array((new BigNumber(5)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = new BigNumber((_1030_v115).length);
        _nw188[(_dafny.ONE).toNumber()] = _this.f11;
        _nw188[(new BigNumber(2)).toNumber()] = new BigNumber(((_1034_v117).update(_this.f11, _1031_v116)).length);
        _nw188[(new BigNumber(3)).toNumber()] = _1023_v110.f12;
        _nw188[(new BigNumber(4)).toNumber()] = (_this).f10;
        (_1029_v114)[_index185] = _nw188;
        (_1023_v110).f12 = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Concat(p0, p0))).length);
        let _1035_v118;
        let _init29 = ((_1036_v105) => function (_1037_i8) {
          return (_1036_v105).f15;
        })(_1018_v105);
        let _nw189 = Array((new BigNumber(10)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw189.length); _i0_29++) {
          _nw189[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1035_v118 = _nw189;
        let _index186 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_1035_v118).length));
        (_1035_v118)[_index186] = (_1018_v105).f15;
        let _index187 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_1035_v118).length));
        (_1035_v118)[_index187] = p1;
        let _index188 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1031_v116).length));
        (_1031_v116)[_index188] = !(p2);
        let _1038_v119;
        _1038_v119 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        let _index189 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1031_v116).length));
        (_1031_v116)[_index189] = ((p3) ? ((((_1038_v119).contains(_module.__default.fm1((_1018_v105).fm4(globalState), p3, _this.f11, !(false), globalState))) ? ((_1038_v119).get(_module.__default.fm1((_1018_v105).fm4(globalState), p3, _this.f11, !(false), globalState))) : (p2))) : (p3));
      } else {
        let _1039_v120;
        let _init30 = function (_1040_i9) {
          return _module.__default.safeModuloInt(_1040_i9, _this.f11);
        };
        let _nw190 = Array((_dafny.ONE).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw190.length); _i0_30++) {
          _nw190[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1039_v120 = _nw190;
        let _1041_v121;
        _1041_v121 = _module.D1.create_DC4(_1039_v120);
        let _1042_v122;
        _1042_v122 = _dafny.Set.fromElements(_module.D1.create_DC4(_1039_v120), _1041_v121, _1041_v121);
        let _1043_v123;
        _1043_v123 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v105).f15,_1042_v122);
        let _1044_v124;
        _1044_v124 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(p3,_1039_v120)),_this.f11);
        let _1045_v125;
        _1045_v125 = _dafny.Seq.of((_this).f10);
        let _1046_v126;
        _1046_v126 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1045_v125);
        let _1047_v127;
        _1047_v127 = _dafny.Seq.of((_dafny.Set.fromElements(_1041_v121)).equals((((_1043_v123).contains((_1018_v105).f15)) ? ((_1043_v123).get((_1018_v105).f15)) : (_1042_v122))), ((_dafny.ZERO).minus((((_1044_v124).contains(_module.D7.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(p3,_1039_v120)))) ? ((_1044_v124).get(_module.D7.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(p3,_1039_v120)))) : (new BigNumber((_1046_v126).length))))).isEqualTo(_this.f11), p3, (_this.f11).isLessThan((_this).f10));
        _1047_v127 = _module.__default.fm9(p2, globalState);
        let _1048_v128;
        _1048_v128 = false;
        _1048_v128 = p2;
        let _1049_v129;
        _1049_v129 = _module.D1.create_DC5(p2);
        let _1050_v130;
        _1050_v130 = _module.D1.create_DC6(_1049_v129);
        let _source15 = _1050_v130;
        if (_source15.is_DC5) {
          let _1051___mcc_h15 = (_source15).cf8;
          let _1052_cf8 = _1051___mcc_h15;
          let _1053_v131;
          _1053_v131 = _dafny.Map.Empty.slice().updateUnsafe(!(_1048_v128),p0);
          let _1054_v132;
          _1054_v132 = _dafny.Seq.of((((_1053_v131).contains(p3)) ? ((_1053_v131).get(p3)) : (p0)));
          let _1055_v133;
          _1055_v133 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_1054_v132);
          let _1056_v134;
          _1056_v134 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(((_1055_v133).contains(_this.f11)) ? ((_1055_v133).get(_this.f11)) : (_1054_v132)));
          _1054_v132 = (((_1056_v134).contains(((_this).f10).plus(new BigNumber((_1047_v127).length)))) ? ((_1056_v134).get(((_this).f10).plus(new BigNumber((_1047_v127).length)))) : (_1054_v132));
          _1052_cf8 = p2;
          let _1057_v136;
          let _nw191 = new _module.C0();
          _nw191.__ctor(function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(-138), new BigNumber(227))) {
              let _1058_v135 = _compr_24;
              if (((new BigNumber(-138)).isLessThanOrEqualTo(_1058_v135)) && ((_1058_v135).isLessThan(new BigNumber(227)))) {
                _coll24.push([(_1058_v135).plus(_this.f11),_1052_cf8]);
              }
            }
            return _coll24;
          }(), _this.f11);
          _1057_v136 = _nw191;
          let _1059_v137;
          _1059_v137 = _module.D1.create_DC5(_1048_v128);
          let _1060_v138;
          let _nw192 = new _module.C2();
          _nw192.__ctor((_1059_v137).dtor_cf8, (_this).f10);
          _1060_v138 = _nw192;
        } else if (_source15.is_DC4) {
          let _1061___mcc_h16 = (_source15).cf7;
          let _1062_cf7 = _1061___mcc_h16;
          let _index190 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1062_cf7).length));
          (_1062_cf7)[_index190] = _this.f11;
          let _index191 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1062_cf7).length));
          (_1062_cf7)[_index191] = _this.f11;
          let _1063_v140;
          _1063_v140 = _dafny.Map.Empty.slice().updateUnsafe((_1062_cf7)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_1062_cf7).length))],(_1018_v105).f15);
          let _1064_v141;
          _1064_v141 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,true);
          let _1065_v142;
          let _nw193 = new _module.C0();
          _nw193.__ctor((function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of (_1063_v140).Keys.Elements) {
              let _1066_v139 = _compr_25;
              if ((_1063_v140).contains(_1066_v139)) {
                _coll25.push([(_1066_v139).plus(_this.f11),p2]);
              }
            }
            return _coll25;
          }()).Merge(_1064_v141), (_this).f10);
          _1065_v142 = _nw193;
          let _1067_v143;
          _1067_v143 = _dafny.Seq.UnicodeFromString("puri");
          _1067_v143 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("s"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), ((_1068_v105) => function (_1069_i10) {
            return (_1068_v105).f15;
          })(_1018_v105))), _1067_v143);
          let _1070_v144;
          _1070_v144 = _1067_v143;
          _1067_v143 = _dafny.Seq.Concat((_1070_v144), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-971)), ((_1071_v105) => function (_1072_i11) {
            return (_1071_v105).f15;
          })(_1018_v105)));
        } else {
          let _1073___mcc_h17 = (_source15).cf9;
          let _1074_cf9 = _1073___mcc_h17;
          _1048_v128 = p3;
          let _1075_v145;
          _1075_v145 = _module.D5.create_DC14(p2);
          let _1076_v146;
          let _nw194 = Array((new BigNumber(11)).toNumber());
          _nw194[(_dafny.ZERO).toNumber()] = p3;
          _nw194[(_dafny.ONE).toNumber()] = _1048_v128;
          _nw194[(new BigNumber(2)).toNumber()] = _dafny.areEqual(_dafny.Seq.of((_this).f10), _1045_v125);
          _nw194[(new BigNumber(3)).toNumber()] = _1048_v128;
          _nw194[(new BigNumber(4)).toNumber()] = false;
          _nw194[(new BigNumber(5)).toNumber()] = _1048_v128;
          _nw194[(new BigNumber(6)).toNumber()] = _1048_v128;
          _nw194[(new BigNumber(7)).toNumber()] = p2;
          _nw194[(new BigNumber(8)).toNumber()] = p3;
          _nw194[(new BigNumber(9)).toNumber()] = !(_module.__default.fm1((_this).f10, !(_1048_v128), (_this).f10, _1048_v128, globalState));
          _nw194[(new BigNumber(10)).toNumber()] = (_1075_v145).dtor_cf20;
          _1076_v146 = _nw194;
          let _1077_v147;
          _1077_v147 = _dafny.Set.fromElements(p3, !(p3));
          let _index192 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1076_v146).length));
          (_1076_v146)[_index192] = _module.__default.fm1(new BigNumber((_1077_v147).length), p3, (_this).f10, p2, globalState);
          let _index193 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1076_v146).length));
          (_1076_v146)[_index193] = true;
          let _1078_v148;
          let _init31 = ((_1079_p0) => function (_1080_i12) {
            return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1079_p0).length)), _dafny.Seq.of(_this.f11));
          })(p0);
          let _nw195 = Array((new BigNumber(3)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw195.length); _i0_31++) {
            _nw195[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1078_v148 = _nw195;
          let _index194 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_1078_v148).length));
          (_1078_v148)[_index194] = _1045_v125;
          let _index195 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_1078_v148).length));
          (_1078_v148)[_index195] = _dafny.Seq.update(_1045_v125, _module.__default.safeIndex(_this.f11, new BigNumber((_1045_v125).length)), (_this).f10);
          _1048_v128 = _1048_v128;
        }
        _1048_v128 = !(_module.__default.fm1(_this.f11, _1048_v128, (_this).f10, (p3) === (p3), globalState));
        let _1081_v149;
        _1081_v149 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1048_v128);
        (globalState).f5 = ((_module.__default.fm1(_this.f11, _1048_v128, (_this).f10, _1048_v128, globalState)) ? (new BigNumber((_1081_v149).length)) : ((_this).f10));
      }
      let _1082_i13;
      _1082_i13 = _dafny.ZERO;
      L6: {
        while (!(!(!(p3)) || (true))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1082_i13)) {
              break L6;
            }
            _1082_i13 = (_1082_i13).plus(_dafny.ONE);
            let _1083_v150;
            _1083_v150 = false;
            let _1084_v151;
            let _init32 = ((_1085_v150) => function (_1086_i14) {
              return _1085_v150;
            })(_1083_v150);
            let _nw196 = Array((new BigNumber(22)).toNumber());
            for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw196.length); _i0_32++) {
              _nw196[_i0_32] = _init32(new BigNumber(_i0_32));
            }
            _1084_v151 = _nw196;
            let _1087_v152;
            _1087_v152 = _dafny.Set.fromElements(_1084_v151);
            _1083_v150 = (_1087_v152).IsProperSubsetOf((_1087_v152).Difference(_1087_v152));
            let _1088_v154;
            _1088_v154 = _dafny.MultiSet.fromElements(false, false);
            let _1089_v155;
            _1089_v155 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1088_v154).cardinality()),true);
            let _1090_v156;
            _1090_v156 = _module.D9.create_DC27(_1018_v105, _this.f11, _this.f11, (_this).f10, function () {
  let _coll26 = new _dafny.Map();
  for (const _compr_26 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_1089_v155))).Elements) {
    let _1091_v153 = _compr_26;
    if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_1089_v155))).contains(_1091_v153)) {
      _coll26.push([_1091_v153,(_1018_v105).f15]);
    }
  }
  return _coll26;
}());
            _1083_v150 = _module.__default.fm1((_this).f10, p3, new BigNumber(((_dafny.Set.fromElements((_this).f10, _this.f11, (_1090_v156).dtor_cf34, (_this).f10)).Union(_866_v1)).length), !(!(((_module.__default.fm1((_dafny.ZERO).minus(_this.f11), p3, (_this).f10, false, globalState)) ? (!(p3)) : (_1083_v150)))), globalState);
            let _1092_v157;
            _1092_v157 = _dafny.Seq.of(_this.f11);
            _1092_v157 = _1092_v157;
            if (_1083_v150) {
              let _1093_v158;
              let _nw197 = Array((new BigNumber(5)).toNumber());
              _nw197[(_dafny.ZERO).toNumber()] = _1092_v157;
              _nw197[(_dafny.ONE).toNumber()] = _1092_v157;
              _nw197[(new BigNumber(2)).toNumber()] = _1092_v157;
              _nw197[(new BigNumber(3)).toNumber()] = _1092_v157;
              _nw197[(new BigNumber(4)).toNumber()] = _1092_v157;
              _1093_v158 = _nw197;
              let _index196 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_1093_v158).length));
              (_1093_v158)[_index196] = _dafny.Seq.of((_1092_v157)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1092_v157).length))], (_this).f10);
              let _1094_v159;
              let _init33 = function (_1095_i15) {
                return _module.D2.create_DC7((_this).f10);
              };
              let _nw198 = Array((new BigNumber(12)).toNumber());
              for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw198.length); _i0_33++) {
                _nw198[_i0_33] = _init33(new BigNumber(_i0_33));
              }
              _1094_v159 = _nw198;
              let _1096_v160;
              let _nw199 = new _module.C3();
              _nw199.__ctor(_1094_v159, (p0)[_module.__default.safeIndex((_this).f10, new BigNumber((p0).length))], (_this).f10);
              _1096_v160 = _nw199;
              let _index197 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_1093_v158).length));
              let _rhs166 = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-973), _1096_v160.f12, new BigNumber(-185)), _1092_v157);
              let _rhs167 = new BigNumber(-939);
              let _rhs168 = ((_this).fm2(false, globalState)).multipliedBy(_1096_v160.f12);
              let _rhs169 = _1096_v160;
              let _rhs170 = _module.__default.fm1(new BigNumber(599), _1083_v150, _this.f11, _1083_v150, globalState);
              let _lhs106 = _1093_v158;
              let _lhs107 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_1093_v158).length));
              let _lhs108 = globalState;
              _lhs106[_lhs107] = _rhs166;
              _lhs108.f5 = _rhs167;
              r0 = _rhs168;
              _1096_v160 = _rhs169;
              _1083_v150 = _rhs170;
              let _1097_v161;
              _1097_v161 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v150,(_this).f10);
              let _1098_v162;
              _1098_v162 = _dafny.Seq.of(true);
              let _rhs171 = (_1096_v160.f12).isLessThanOrEqualTo(new BigNumber(((_1097_v161).Merge(_1097_v161)).length));
              let _rhs172 = _module.__default.fm1(new BigNumber((_1089_v155).length), _dafny.Seq.contains(_1098_v162, false), ((_1083_v150) ? ((_this).f10) : ((_dafny.ZERO).minus(new BigNumber((p0).length)))), p3, globalState);
              let _rhs173 = (_1096_v160.f12).isEqualTo((_this).f10);
              _1083_v150 = _rhs171;
              _1083_v150 = _rhs172;
              _1083_v150 = _rhs173;
              let _index198 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1084_v151).length));
              (_1084_v151)[_index198] = !(p2);
              let _index199 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1084_v151).length));
              (_1084_v151)[_index199] = p3;
              let _1099_v163;
              _1099_v163 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f10));
              let _1100_v164;
              let _nw200 = new _module.C2();
              _nw200.__ctor(p3, (_dafny.ZERO).minus(_module.__default.fm0(new BigNumber(((_1099_v163)[_module.__default.safeIndex((_this).f10, new BigNumber((_1099_v163).length))]).length), new BigNumber((_1097_v161).length), globalState)));
              _1100_v164 = _nw200;
              _1100_v164 = _1100_v164;
              let _index200 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1084_v151).length));
              (_1084_v151)[_index200] = p3;
            } else {
              _1083_v150 = p3;
              let _1101_v165;
              let _init34 = ((_1102_v150, _1103_p2) => function (_1104_i16) {
                return _module.D5.create_DC13(!(_1102_v150), !(_1103_p2));
              })(_1083_v150, p2);
              let _nw201 = Array((new BigNumber(5)).toNumber());
              for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw201.length); _i0_34++) {
                _nw201[_i0_34] = _init34(new BigNumber(_i0_34));
              }
              _1101_v165 = _nw201;
              let _1105_v166;
              _1105_v166 = _module.D5.create_DC13(!(p3), p3);
              let _index201 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1101_v165).length));
              (_1101_v165)[_index201] = _1105_v166;
              let _index202 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1101_v165).length));
              (_1101_v165)[_index202] = _1105_v166;
              let _1106_v167;
              let _nw202 = new _module.C1();
              _nw202.__ctor((_1018_v105).f15, (_1018_v105).f15, _this.f11);
              _1106_v167 = _nw202;
              let _1107_v168;
              let _nw203 = new _module.C1();
              _nw203.__ctor((_1106_v167).f15, (_1106_v167).f15, (_this).f10);
              _1107_v168 = _nw203;
              let _1108_v169;
              let _nw204 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
              _1108_v169 = _nw204;
              let _1109_v170;
              _1109_v170 = _dafny.Set.fromElements(p2, true, _1083_v150);
              let _index203 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1108_v169).length));
              (_1108_v169)[_index203] = new BigNumber((_1109_v170).length);
              let _index204 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1108_v169).length));
              (_1108_v169)[_index204] = (_this).f10;
            }
          }
        }
      }
      if ((_module.D2.create_DC8(p3, _this.f11, p2)).dtor_cf11) {
        let _1110_v171;
        _1110_v171 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p2);
        let _1111_v172;
        let _nw205 = new _module.C0();
        _nw205.__ctor(_1110_v171, (_this).f10);
        _1111_v172 = _nw205;
        let _1112_v173;
        _1112_v173 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v105).f15,_1111_v172);
        _1112_v173 = _1112_v173;
        if (((_this).f10).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("eubmv")).length))) {
          let _1113_v174;
          let _1114_v175;
          let _1115_v176;
          let _out54;
          let _out55;
          let _out56;
          let _outcollector16 = (_1018_v105).m5(globalState);
          _out54 = _outcollector16[0];
          _out55 = _outcollector16[1];
          _out56 = _outcollector16[2];
          _1113_v174 = _out54;
          _1114_v175 = _out55;
          _1115_v176 = _out56;
          let _1116_v177;
          _1116_v177 = _dafny.Set.fromElements(_1114_v175);
          let _1117_v178;
          _1117_v178 = _module.D0.create_DC0(true);
          let _1118_v179;
          let _nw206 = Array((new BigNumber(28)).toNumber());
          _nw206[(_dafny.ZERO).toNumber()] = _1116_v177;
          _nw206[(_dafny.ONE).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(p2);
          _nw206[(new BigNumber(3)).toNumber()] = (_1116_v177).Union(_1116_v177);
          _nw206[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(true);
          _nw206[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_1114_v175);
          _nw206[(new BigNumber(6)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(7)).toNumber()] = (_dafny.Set.fromElements(p3, p3)).Difference(_1116_v177);
          _nw206[(new BigNumber(8)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements((_1117_v178).dtor_cf0, p3, !(_1114_v175), _module.__default.fm1(new BigNumber(818), p3, _1113_v174, p2, globalState));
          _nw206[(new BigNumber(10)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(11)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(p2, p2);
          _nw206[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(p3);
          _nw206[(new BigNumber(14)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(15)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(16)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(17)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements(_1114_v175, _1114_v175, p3);
          _nw206[(new BigNumber(19)).toNumber()] = (_1116_v177).Difference(_1116_v177);
          _nw206[(new BigNumber(20)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(21)).toNumber()] = (_1116_v177).Union(_dafny.Set.fromElements(p3));
          _nw206[(new BigNumber(22)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(23)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(24)).toNumber()] = (_1116_v177).Difference(_1116_v177);
          _nw206[(new BigNumber(25)).toNumber()] = _1116_v177;
          _nw206[(new BigNumber(26)).toNumber()] = (_1116_v177).Intersect(_dafny.Set.fromElements(p3));
          _nw206[(new BigNumber(27)).toNumber()] = _1116_v177;
          _1118_v179 = _nw206;
          let _1119_v180;
          _1119_v180 = _1116_v177;
          let _index205 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1118_v179).length));
          (_1118_v179)[_index205] = ((_1119_v180)).Intersect((_1119_v180));
          let _index206 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1118_v179).length));
          let _rhs174 = false;
          let _rhs175 = (_1116_v177).Difference((_1116_v177).Difference(_1116_v177));
          let _lhs109 = _1118_v179;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1118_v179).length));
          _1114_v175 = _rhs174;
          _lhs109[_lhs110] = _rhs175;
          let _1120_v181;
          let _init35 = ((_1121_p0) => function (_1122_i17) {
            return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hhjkitey"), _1121_p0);
          })(p0);
          let _nw207 = Array((new BigNumber(8)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw207.length); _i0_35++) {
            _nw207[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1120_v181 = _nw207;
          let _index207 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1120_v181).length));
          (_1120_v181)[_index207] = _dafny.Seq.Concat(p0, p0);
          let _1123_v182;
          _1123_v182 = _dafny.Seq.of(_module.__default.fm1(_1113_v174, false, (_this).f10, p3, globalState));
          let _1124_v183;
          _1124_v183 = _module.D8.create_DC24(_1116_v177, _1123_v182);
          let _pat_let_tv8 = _1118_v179;
          let _pat_let_tv9 = _1118_v179;
          let _1125_v184;
          _1125_v184 = _dafny.MultiSet.fromElements(function (_pat_let17_0) {
            return function (_1126_dt__update__tmp_h2) {
              return function (_pat_let18_0) {
                return function (_1127_dt__update_hcf30_h0) {
                  return _module.D8.create_DC24(_1127_dt__update_hcf30_h0, (_1126_dt__update__tmp_h2).dtor_cf31);
                }(_pat_let18_0);
              }((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_pat_let_tv8).length))]);
            }(_pat_let17_0);
          }(_1124_v183));
          let _index208 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1120_v181).length));
          (_1120_v181)[_index208] = (((_1125_v184).IsSubsetOf(_1125_v184)) ? (p0) : (p0));
          let _1128_v185;
          let _init36 = ((_1129_v106) => function (_1130_i18) {
            return (_1130_i18).plus(new BigNumber((_1129_v106).cardinality()));
          })(_1019_v106);
          let _nw208 = Array((new BigNumber(4)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw208.length); _i0_36++) {
            _nw208[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1128_v185 = _nw208;
          let _1131_v186;
          let _init37 = ((_1132_v1) => function (_1133_i19) {
            return _module.D2.create_DC7(new BigNumber((_1132_v1).length));
          })(_866_v1);
          let _nw209 = Array((new BigNumber(2)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw209.length); _i0_37++) {
            _nw209[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1131_v186 = _nw209;
          let _1134_v187;
          let _nw210 = new _module.C3();
          _nw210.__ctor(_1131_v186, (_1018_v105).f15, (_this).f10);
          _1134_v187 = _nw210;
          let _1135_v188;
          _1135_v188 = _dafny.Seq.of(_1134_v187);
          let _1136_v189;
          _1136_v189 = _dafny.Map.Empty.slice().updateUnsafe(_1128_v185,_1135_v188);
          _1136_v189 = (_1136_v189).update(((false) ? (_1128_v185) : (_1128_v185)), _dafny.Seq.Concat(_1135_v188, _1135_v188));
          (globalState).f5 = _module.__default.safeDivisionInt((_this).f10, (_this).f10);
        } else {
          let _1137_v190;
          _1137_v190 = false;
          _1137_v190 = p3;
          let _1138_v191;
          _1138_v191 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v105).f15,_dafny.MultiSet.FromArray((_1018_v105).fm5(p2, !(p2), false, !(p3), globalState)));
          let _1139_v192;
          _1139_v192 = _dafny.Seq.of((_this).f10);
          let _1140_v193;
          _1140_v193 = _module.D0.create_DC0(p2);
          let _1141_v194;
          _1141_v194 = _dafny.Seq.of(_1137_v190);
          let _1142_v196;
          let _nw211 = Array((new BigNumber(20)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = false;
          _nw211[(_dafny.ONE).toNumber()] = (_module.__default.fm21((_this).f10, (_this).f10, new _dafny.CodePoint('h'.codePointAt(0)), (_this).f10, globalState)).IsSubsetOf(((((_1138_v191).contains((_1018_v105).f15)) ? ((_1138_v191).get((_1018_v105).f15)) : (_1019_v106))).update((_this).f10, _module.__default.abs(new BigNumber(-600))));
          _nw211[(new BigNumber(2)).toNumber()] = p2;
          _nw211[(new BigNumber(3)).toNumber()] = p3;
          _nw211[(new BigNumber(4)).toNumber()] = p2;
          _nw211[(new BigNumber(5)).toNumber()] = p2;
          _nw211[(new BigNumber(6)).toNumber()] = !(p2);
          _nw211[(new BigNumber(7)).toNumber()] = (new BigNumber((_1139_v192).length)).isLessThanOrEqualTo((((_1019_v106).contains(_this.f11)) ? ((_1019_v106).get(_this.f11)) : (_this.f11)));
          _nw211[(new BigNumber(8)).toNumber()] = (_1140_v193).dtor_cf0;
          _nw211[(new BigNumber(9)).toNumber()] = p3;
          _nw211[(new BigNumber(10)).toNumber()] = _module.__default.fm1(new BigNumber((_1139_v192).length), _module.__default.fm1(new BigNumber((_dafny.Set.fromElements((_this).f10)).length), p3, new BigNumber((_dafny.Seq.UnicodeFromString("wma")).length), (_1141_v194)[_module.__default.safeIndex((_this).f10, new BigNumber((_1141_v194).length))], globalState), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f10)), _1137_v190, globalState);
          _nw211[(new BigNumber(11)).toNumber()] = true;
          _nw211[(new BigNumber(12)).toNumber()] = true;
          _nw211[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), function (_1143_i20) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), p0);
          _nw211[(new BigNumber(14)).toNumber()] = p2;
          _nw211[(new BigNumber(15)).toNumber()] = !((function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(612), new BigNumber(-261))) {
              let _1144_v195 = _compr_27;
              if (((new BigNumber(612)).isLessThanOrEqualTo(_1144_v195)) && ((_1144_v195).isLessThan(new BigNumber(-261)))) {
                _coll27.add(_module.__default.safeDivisionInt(_1144_v195, (_this).f10));
              }
            }
            return _coll27;
          }()).IsSubsetOf(_866_v1));
          _nw211[(new BigNumber(16)).toNumber()] = _1137_v190;
          _nw211[(new BigNumber(17)).toNumber()] = _1137_v190;
          _nw211[(new BigNumber(18)).toNumber()] = (_this.f11).isLessThanOrEqualTo((_this).f10);
          _nw211[(new BigNumber(19)).toNumber()] = false;
          _1142_v196 = _nw211;
          let _index209 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1142_v196).length));
          (_1142_v196)[_index209] = p2;
          let _1145_v197;
          _1145_v197 = _dafny.Set.fromElements(false);
          let _1146_v198;
          _1146_v198 = _1145_v197;
          let _1147_v199;
          _1147_v199 = _dafny.Seq.of(_1111_v172);
          let _1148_v200;
          let _nw212 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _1148_v200 = _nw212;
          let _1149_v201;
          _1149_v201 = _dafny.Seq.of(_1148_v200, _1148_v200, _1148_v200);
          let _index210 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1142_v196).length));
          let _rhs176 = _module.__default.safeModuloInt(new BigNumber((_866_v1).length), new BigNumber((_1147_v199).length));
          let _rhs177 = ((_this).f10).multipliedBy((_this).f10);
          let _rhs178 = ((_1019_v106).Intersect(_1019_v106)).IsProperSubsetOf(_1019_v106);
          let _rhs179 = _1146_v198;
          let _rhs180 = (_1149_v201)[_module.__default.safeIndex(_this.f11, new BigNumber((_1149_v201).length))];
          let _lhs111 = globalState;
          let _lhs112 = globalState;
          let _lhs113 = _1142_v196;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1142_v196).length));
          _lhs111.f5 = _rhs176;
          _lhs112.f5 = _rhs177;
          _lhs113[_lhs114] = _rhs178;
          _1146_v198 = _rhs179;
          r3 = _rhs180;
          let _1150_v202;
          _1150_v202 = _module.D1.create_DC5(p3);
          _1137_v190 = (_1150_v202).dtor_cf8;
          let _nw213 = new _module.C0();
          _nw213.__ctor(function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(361), new BigNumber(634))) {
              let _1151_v203 = _compr_28;
              if (((new BigNumber(361)).isLessThanOrEqualTo(_1151_v203)) && ((_1151_v203).isLessThan(new BigNumber(634)))) {
                _coll28.push([_module.__default.safeDivisionInt(_1151_v203, (_this).f10),true]);
              }
            }
            return _coll28;
          }(), _this.f11);
          _1111_v172 = _nw213;
          let _1152_v204;
          _1152_v204 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1139_v192, _1139_v192),_1141_v194);
          _1141_v194 = (((_1152_v204).contains(_1139_v192)) ? ((_1152_v204).get(_1139_v192)) : (_dafny.Seq.of(_1137_v190)));
        }
        let _1153_v205;
        _1153_v205 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),(p2) && (p3));
        _1153_v205 = (_1153_v205).update(p3, p3);
        let _1154_v206;
        _1154_v206 = _dafny.Seq.of(_this.f11);
        (globalState).f5 = (_1154_v206)[_module.__default.safeIndex(new BigNumber((_1019_v106).cardinality()), new BigNumber((_1154_v206).length))];
        let _1155_v207;
        let _nw214 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1155_v207 = _nw214;
        let _index211 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1155_v207).length));
        (_1155_v207)[_index211] = ((p3) ? (p2) : (p3));
        let _index212 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1155_v207).length));
        (_1155_v207)[_index212] = p3;
      } else {
        let _1156_v208;
        _1156_v208 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
        _1156_v208 = (_1156_v208).update(p3, _module.__default.fm1((_this).f10, p2, _this.f11, !(p2), globalState));
        let _1157_v209;
        let _init38 = ((_1158_p3, _1159_p2) => function (_1160_i21) {
          return _dafny.Set.fromElements(_1158_p3, _1159_p2, _1159_p2, !(true), _1159_p2);
        })(p3, p2);
        let _nw215 = Array((new BigNumber(22)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw215.length); _i0_38++) {
          _nw215[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1157_v209 = _nw215;
        _1157_v209 = _1157_v209;
        let _1161_v210;
        _1161_v210 = _module.D2.create_DC7(_this.f11);
        let _1162_v211;
        let _nw216 = Array((new BigNumber(3)).toNumber());
        _nw216[(_dafny.ZERO).toNumber()] = _1161_v210;
        _nw216[(_dafny.ONE).toNumber()] = _module.D2.create_DC7(_this.f11);
        _nw216[(new BigNumber(2)).toNumber()] = _1161_v210;
        _1162_v211 = _nw216;
        let _1163_v212;
        let _nw217 = new _module.C3();
        _nw217.__ctor(_1162_v211, (_1018_v105).f15, _this.f11);
        _1163_v212 = _nw217;
        let _1164_v213;
        _1164_v213 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v212,_this.f11);
        let _1165_v214;
        _1165_v214 = _dafny.MultiSet.fromElements(_1164_v213, _1164_v213);
        let _1166_v215;
        _1166_v215 = _dafny.MultiSet.fromElements((_1165_v214).IsDisjointFrom(_1165_v214));
        _1166_v215 = _1166_v215;
        let _1167_v216;
        let _init39 = function (_1168_i22) {
          return (_1168_i22).minus((_this).f10);
        };
        let _nw218 = Array((new BigNumber(10)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw218.length); _i0_39++) {
          _nw218[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1167_v216 = _nw218;
        let _index213 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1167_v216).length));
        (_1167_v216)[_index213] = (_dafny.ZERO).minus((_this).f10);
        let _index214 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1167_v216).length));
        (_1167_v216)[_index214] = new BigNumber((p0).length);
        let _1169_v217;
        let _nw219 = Array((new BigNumber(25)).toNumber()).fill(false);
        _1169_v217 = _nw219;
        let _1170_v218;
        _1170_v218 = _module.D0.create_DC1(_1169_v217);
        _1170_v218 = _1170_v218;
      }
      _866_v1 = _866_v1;
      r0 = new BigNumber(-907);
      let _1171_v219;
      _1171_v219 = _dafny.MultiSet.fromElements(p2);
      let _1172_v220;
      let _nw220 = Array((new BigNumber(5)).toNumber());
      _nw220[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsProperPrefixOf(p0, p0);
      _nw220[(_dafny.ONE).toNumber()] = (_this.f11).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), ((_1173_v105) => function (_1174_i23) {
        return (_1173_v105).f15;
      })(_1018_v105))).length));
      _nw220[(new BigNumber(2)).toNumber()] = !(p3);
      _nw220[(new BigNumber(3)).toNumber()] = !(_module.__default.fm1((_this).f10, p2, new BigNumber((_1171_v219).cardinality()), p2, globalState));
      _nw220[(new BigNumber(4)).toNumber()] = p3;
      _1172_v220 = _nw220;
      r1 = _1172_v220;
      let _1175_v221;
      let _init40 = ((_1176_v2) => function (_1177_i24) {
        return _1176_v2;
      })(_867_v2);
      let _nw221 = Array((new BigNumber(3)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw221.length); _i0_40++) {
        _nw221[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1175_v221 = _nw221;
      let _1178_v222;
      _1178_v222 = _dafny.Set.fromElements(_1175_v221);
      r2 = ((_1178_v222).Intersect(_1178_v222)).Union(_1178_v222);
      let _1179_v223;
      let _nw222 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _1179_v223 = _nw222;
      r3 = _1179_v223;
      return [r0, r1, r2, r3];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
