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
    static fm2(globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(910)).plus(new BigNumber(260)));
    };
    static fm3(p0, p1, p2, globalState) {
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uirae"), _dafny.Seq.UnicodeFromString("r"))).length)).multipliedBy(new BigNumber(21));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true, true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-93)));
    };
    static fm5(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,(_module.D9.create_DC26(_dafny.MultiSet.fromElements(new BigNumber(-657), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), function (_0_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length)))).dtor_cf38);
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(56)).minus(new BigNumber((_dafny.Seq.of(true, true)).length)),((true) ? (new BigNumber(((_module.D10.create_DC28(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(384)), function (_1_i0) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})).length)))).dtor_cf39).length)) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("y"))).length))))));
    };
    static fm7(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("qaoo"))).Intersect((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), function (_2_i0) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }))).Elements) {
          let _3_v0 = _compr_0;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), function (_2_i0) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }))).contains(_3_v0)) {
            _coll0.add(_3_v0);
          }
        }
        return _coll0;
      }()).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-523)), function (_4_i1) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("flq"))));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return !(!((_dafny.ZERO).minus((new BigNumber(478)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length)))).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(128), new BigNumber(805))).cardinality())));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC2(new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(216), _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of (_dafny.Seq.of(new BigNumber(989), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).cardinality()))).Elements) {
    let _5_v0 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(989), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).cardinality())), _5_v0)) {
      _coll1.add(_module.__default.safeDivisionInt(_5_v0, new BigNumber(-206)));
    }
  }
  return _coll1;
}()).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),_module.D10.create_DC29((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length))), new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(832), new BigNumber(101))) {
    let _6_v1 = _compr_2;
    if (((new BigNumber(832)).isLessThanOrEqualTo(_6_v1)) && ((_6_v1).isLessThan(new BigNumber(101)))) {
      _coll2.push([_module.__default.safeDivisionInt(_6_v1, new BigNumber(228)),!(true)]);
    }
  }
  return _coll2;
}()).length)))).length)));
    };
    static fm10(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sh"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pcvumw"), _dafny.Seq.UnicodeFromString("yl")));
    };
    static m1(p0, p1, globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = undefined;
      let r2 = _dafny.Seq.of();
      let r3 = false;
      if (p1) {
        let _7_v0;
        _7_v0 = new BigNumber(845);
        _7_v0 = _7_v0;
        let _8_v1;
        let _nw0 = new _module.C1();
        _nw0.__ctor();
        _8_v1 = _nw0;
        let _9_v3;
        _9_v3 = _dafny.Seq.of(p1);
        let _10_v4;
        _10_v4 = _dafny.Map.Empty.slice().updateUnsafe(_7_v0,_9_v3);
        let _11_v5;
        _11_v5 = _dafny.Map.Empty.slice().updateUnsafe(_7_v0,_module.__default.fm8(_7_v0, function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_10_v4).Keys.Elements) {
            let _12_v2 = _compr_3;
            if ((_10_v4).contains(_12_v2)) {
              _coll3.push([(_12_v2).minus(_7_v0),_7_v0]);
            }
          }
          return _coll3;
        }(), _7_v0, true, globalState));
        let _13_v6;
        _13_v6 = _dafny.Map.Empty.slice().updateUnsafe(_8_v1,_11_v5);
        let _14_v7;
        _14_v7 = _module.D8.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(_8_v1,_dafny.Map.Empty.slice().updateUnsafe(_7_v0,p1)));
        let _15_v8;
        _15_v8 = _module.D8.create_DC22((_14_v7).dtor_cf29);
        _13_v6 = ((_13_v6).Merge((_14_v7).dtor_cf29)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_8_v1,_11_v5)).update(_8_v1, _11_v5));
        let _16_v9;
        let _nw1 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _16_v9 = _nw1;
        let _17_v10;
        _17_v10 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_7_v0));
        let _index0 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_16_v9).length));
        (_16_v9)[_index0] = _17_v10;
        let _18_v11;
        _18_v11 = _dafny.Map.Empty.slice().updateUnsafe(_7_v0,_7_v0);
        let _19_v12;
        _19_v12 = _dafny.Set.fromElements(p1, p1);
        let _20_v13;
        _20_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm2(globalState));
        let _21_v14;
        _21_v14 = false;
        let _22_v15;
        let _nw2 = new _module.C0();
        _nw2.__ctor(_20_v13, _7_v0, _21_v14);
        _22_v15 = _nw2;
        let _23_v16;
        _23_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(303));
        let _24_v17;
        _24_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,_23_v16);
        let _25_v18;
        _25_v18 = _dafny.MultiSet.fromElements(_7_v0);
        let _index1 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_16_v9).length));
        let _rhs0 = _module.__default.fm8(_7_v0, (_18_v11).update(new BigNumber((_19_v12).length), _7_v0), _7_v0, p1, globalState);
        let _rhs1 = (_dafny.ZERO).minus((((p1) && (true)) ? (_7_v0) : (_7_v0)));
        let _rhs2 = _22_v15;
        let _rhs3 = _module.__default.safeDivisionInt(_module.__default.fm3(p0, p1, p1, globalState), (_dafny.ZERO).minus((_7_v0).multipliedBy(new BigNumber(547))));
        let _rhs4 = _dafny.Seq.Concat(_17_v10, _dafny.Seq.Concat(_17_v10, _dafny.Seq.update(_17_v10, _module.__default.safeIndex(new BigNumber((_24_v17).length), new BigNumber((_17_v10).length)), (_25_v18).update(_7_v0, _module.__default.abs(_7_v0)))));
        let _lhs0 = _16_v9;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_16_v9).length));
        r3 = _rhs0;
        _7_v0 = _rhs1;
        r1 = _rhs2;
        _7_v0 = _rhs3;
        _lhs0[_lhs1] = _rhs4;
        let _26_v19;
        _26_v19 = _dafny.Seq.UnicodeFromString("pbwt");
        let _rhs5 = (_dafny.ZERO).minus(((p1) ? (_7_v0) : (_7_v0)));
        let _rhs6 = _7_v0;
        let _rhs7 = ((p1) ? (_7_v0) : ((_22_v15).fm0(_7_v0, p0, p1, _7_v0, globalState)));
        let _rhs8 = _26_v19;
        _7_v0 = _rhs5;
        _7_v0 = _rhs6;
        _7_v0 = _rhs7;
        r0 = _rhs8;
        _23_v16 = (_23_v16).update(true, new BigNumber((_26_v19).length));
      } else {
        let _27_v20;
        let _nw3 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _27_v20 = _nw3;
        let _28_v21;
        _28_v21 = _dafny.Seq.of(_27_v20, _27_v20, _27_v20);
        _28_v21 = _28_v21;
        let _29_v22;
        _29_v22 = new BigNumber(736);
        _29_v22 = (_29_v22).minus(_module.__default.safeModuloInt(new BigNumber(-887), _29_v22));
        let _30_v23;
        _30_v23 = _dafny.MultiSet.fromElements(_29_v22);
        let _31_v24;
        _31_v24 = _dafny.Map.Empty.slice().updateUnsafe(false,_30_v23);
        let _32_v25;
        _32_v25 = p1;
        let _33_v26;
        let _nw4 = new _module.C0();
        _nw4.__ctor(_31_v24, new BigNumber((_dafny.Seq.of(!(p1))).length), _32_v25);
        _33_v26 = _nw4;
        let _34_v27;
        _34_v27 = _dafny.MultiSet.fromElements(_33_v26);
        let _35_v28;
        _35_v28 = _dafny.MultiSet.fromElements(_34_v27, _34_v27);
        _35_v28 = (_dafny.MultiSet.fromElements(_34_v27)).Difference(_35_v28);
        let _36_v29;
        _36_v29 = _dafny.Map.Empty.slice().updateUnsafe((_33_v26).f6,(_33_v26).f6);
        if (!(_module.__default.fm8((_dafny.ZERO).minus(_29_v22), _36_v29, new BigNumber((function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(454), new BigNumber(-189))) {
            let _37_v30 = _compr_4;
            if (((new BigNumber(454)).isLessThanOrEqualTo(_37_v30)) && ((_37_v30).isLessThan(new BigNumber(-189)))) {
              _coll4.push([(_37_v30).plus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(p1))).cardinality())),(_dafny.ZERO).minus(_29_v22)]);
            }
          }
          return _coll4;
        }()).length), p1, globalState)) || (p1)) {
          let _38_v31;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_dafny.Map.Empty.slice().updateUnsafe(!(true),_30_v23), _29_v22, p1);
          _38_v31 = _nw5;
          _29_v22 = _29_v22;
          let _39_v32;
          _39_v32 = _dafny.Set.fromElements(p1);
          let _40_v33;
          _40_v33 = _dafny.Map.Empty.slice().updateUnsafe(_39_v32,false);
          _29_v22 = (_module.__default.safeModuloInt((_33_v26).f6, (_33_v26).f6)).minus(new BigNumber((_40_v33).length));
          let _41_v34;
          _41_v34 = _module.D6.create_DC15(_38_v31);
          let _42_v35;
          _42_v35 = _dafny.Seq.of(_38_v31, _38_v31, _38_v31, _38_v31);
          let _43_v36;
          _43_v36 = _dafny.Seq.of(p1);
          let _44_v37;
          _44_v37 = _dafny.Map.Empty.slice().updateUnsafe((_43_v36)[_module.__default.safeIndex(new BigNumber(-721), new BigNumber((_43_v36).length))],_29_v22);
          let _45_v38;
          _45_v38 = _dafny.Seq.UnicodeFromString("njpmb");
          let _46_v39;
          let _nw6 = Array((new BigNumber(10)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _38_v31;
          _nw6[(_dafny.ONE).toNumber()] = _38_v31;
          _nw6[(new BigNumber(2)).toNumber()] = _38_v31;
          _nw6[(new BigNumber(3)).toNumber()] = _38_v31;
          _nw6[(new BigNumber(4)).toNumber()] = (_41_v34).dtor_cf21;
          _nw6[(new BigNumber(5)).toNumber()] = _38_v31;
          _nw6[(new BigNumber(6)).toNumber()] = _38_v31;
          _nw6[(new BigNumber(7)).toNumber()] = (_42_v35)[_module.__default.safeIndex(new BigNumber(((_44_v37).update(p1, new BigNumber((_45_v38).length))).length), new BigNumber((_42_v35).length))];
          _nw6[(new BigNumber(8)).toNumber()] = _38_v31;
          _nw6[(new BigNumber(9)).toNumber()] = _38_v31;
          _46_v39 = _nw6;
          _46_v39 = _46_v39;
          let _47_v40;
          _47_v40 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          _47_v40 = (_47_v40).update(p1, !(_dafny.Seq.IsProperPrefixOf(_45_v38, _45_v38)));
        } else {
          let _48_v41;
          let _nw7 = Array((new BigNumber(14)).toNumber()).fill(false);
          _48_v41 = _nw7;
          let _index2 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_48_v41).length));
          (_48_v41)[_index2] = p1;
          let _49_v42;
          _49_v42 = _dafny.Map.Empty.slice().updateUnsafe(_29_v22,p1);
          let _50_v43;
          _50_v43 = _dafny.Seq.of((((_49_v42).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_29_v22)))) ? ((_49_v42).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_29_v22)))) : (p1)));
          let _index3 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_48_v41).length));
          (_48_v41)[_index3] = ((_33_v26).f6).isLessThan(new BigNumber((_50_v43).length));
          _29_v22 = (_dafny.ZERO).minus(_29_v22);
          let _51_v44;
          let _nw8 = new _module.C1();
          _nw8.__ctor();
          _51_v44 = _nw8;
          let _52_v45;
          _52_v45 = _dafny.Map.Empty.slice().updateUnsafe(_29_v22,_51_v44);
          let _53_v46;
          _53_v46 = _module.D1.create_DC2(new _dafny.CodePoint('u'.codePointAt(0)), _29_v22, new BigNumber(-318));
          let _54_v47;
          _54_v47 = _dafny.Set.fromElements(new BigNumber((_50_v43).length));
          let _55_v48;
          _55_v48 = _dafny.Set.fromElements((_module.D1.create_DC2((_53_v46).dtor_cf2, (_33_v26).f6, (_dafny.ZERO).minus(new BigNumber((_54_v47).length)))).dtor_cf4);
          _52_v45 = (_52_v45).update((_33_v26).f6, ((!(_module.__default.fm8((_33_v26).f6, _dafny.Map.Empty.slice().updateUnsafe((_33_v26).f6,(_33_v26).f6), new BigNumber((_54_v47).length), p1, globalState))) ? (_51_v44) : (_51_v44)));
          let _index4 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_48_v41).length));
          (_48_v41)[_index4] = p1;
          let _56_v49;
          let _nw9 = new _module.C1();
          _nw9.__ctor();
          _56_v49 = _nw9;
        }
        let _57_v50;
        let _nw10 = new _module.C1();
        _nw10.__ctor();
        _57_v50 = _nw10;
      }
      let _58_v51;
      let _init0 = ((_59_p1) => function (_60_i0) {
        return _module.__default.safeModuloInt(_60_i0, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ovbwu"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_59_p1)).length), new BigNumber((_dafny.Seq.UnicodeFromString("ovbwu")).length)), new _dafny.CodePoint('n'.codePointAt(0)))).length));
      })(p1);
      let _nw11 = Array((new BigNumber(19)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw11.length); _i0_0++) {
        _nw11[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _58_v51 = _nw11;
      let _61_v52;
      _61_v52 = new BigNumber(926);
      let _index5 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length));
      (_58_v51)[_index5] = (_dafny.ZERO).minus(_61_v52);
      let _index6 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length));
      (_58_v51)[_index6] = _61_v52;
      let _62_v53;
      _62_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.MultiSet.fromElements(_61_v52));
      let _63_v54;
      _63_v54 = p1;
      let _64_v55;
      let _nw12 = new _module.C0();
      _nw12.__ctor(_62_v53, (_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))], _63_v54);
      _64_v55 = _nw12;
      let _65_v56;
      _65_v56 = _module.D7.create_DC21(p1, _64_v55);
      if ((_65_v56).dtor_cf27) {
        _61_v52 = (_dafny.ZERO).minus(((_64_v55).f6).minus((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]));
        let _66_v57;
        _66_v57 = _dafny.Seq.UnicodeFromString("fqqxnh");
        let _67_v58;
        _67_v58 = _dafny.MultiSet.fromElements((_64_v55).f6);
        let _68_v59;
        _68_v59 = _dafny.MultiSet.fromElements(_58_v51, _58_v51);
        let _69_v60;
        _69_v60 = _module.D3.create_DC6(_68_v59);
        let _70_v61;
        let _nw13 = new _module.C1();
        _nw13.__ctor();
        _70_v61 = _nw13;
        let _71_v62;
        _71_v62 = _dafny.Map.Empty.slice().updateUnsafe(_69_v60,_70_v61);
        let _72_v63;
        _72_v63 = _dafny.Seq.of((_64_v55).f6, ((_64_v55).f6).multipliedBy((_64_v55).f6), _module.__default.safeDivisionInt((_64_v55).f6, new BigNumber((_66_v57).length)), new BigNumber(((_67_v58).Difference(_67_v58)).cardinality()), ((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]).multipliedBy(new BigNumber((_71_v62).length)));
        let _index7 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length));
        (_58_v51)[_index7] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_72_v63)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_61_v52)), new BigNumber((_72_v63).length))]));
        r3 = p1;
        let _73_v64;
        let _nw14 = new _module.C0();
        _nw14.__ctor(_module.__default.fm5(globalState), _61_v52, _63_v54);
        _73_v64 = _nw14;
        let _74_v65;
        _74_v65 = _dafny.Map.Empty.slice().updateUnsafe(_73_v64,!(false));
        _74_v65 = (_74_v65).update(_73_v64, p1);
        let _75_v66;
        _75_v66 = _dafny.Set.fromElements(_72_v63);
        let _index8 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length));
        let _rhs9 = new BigNumber(-252);
        let _rhs10 = _dafny.Seq.Concat(_66_v57, _dafny.Seq.Concat(_66_v57, _dafny.Seq.UnicodeFromString("jnqvr")));
        let _rhs11 = _61_v52;
        let _rhs12 = _dafny.Set.fromElements(_72_v63, _72_v63);
        let _lhs2 = _58_v51;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length));
        _lhs2[_lhs3] = _rhs9;
        _66_v57 = _rhs10;
        _61_v52 = _rhs11;
        _75_v66 = _rhs12;
      } else {
        let _76_v67;
        _76_v67 = _module.D5.create_DC12(((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]).isLessThanOrEqualTo(_61_v52));
        _76_v67 = _76_v67;
        let _77_v68;
        _77_v68 = _dafny.Seq.UnicodeFromString("ibjngh");
        r3 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_77_v68, _module.__default.fm10(globalState)), _77_v68);
        let _78_v69;
        _78_v69 = _dafny.Seq.of(new BigNumber((_77_v68).length), new BigNumber((_77_v68).length), new BigNumber(751), new BigNumber(427));
        let _79_v70;
        _79_v70 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_64_v55);
        let _80_v71;
        _80_v71 = _dafny.Map.Empty.slice().updateUnsafe(_78_v69,_79_v70);
        let _81_v72;
        _81_v72 = _dafny.Map.Empty.slice().updateUnsafe(_61_v52,(_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]);
        r3 = !_dafny.Seq.contains(_dafny.Seq.of(_79_v70), ((((_80_v71).contains(_78_v69)) ? ((_80_v71).get(_78_v69)) : (_79_v70))).update(_module.__default.fm8((_64_v55).f6, _81_v72, new BigNumber(902), p1, globalState), _64_v55));
        r3 = (((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]).minus((_64_v55).f6)).isEqualTo((_64_v55).f6);
        let _82_v73;
        _82_v73 = _module.D1.create_DC2(p0, _61_v52, (_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]);
        r3 = ((!_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC2(p0, _61_v52, (_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]), _82_v73), _82_v73)) ? (p1) : ((new BigNumber(557)).isEqualTo((_64_v55).f6)));
      }
      let _83_v74;
      let _nw15 = new _module.C0();
      _nw15.__ctor((_62_v53).Merge(_62_v53), _61_v52, _63_v54);
      _83_v74 = _nw15;
      if (p1) {
        let _84_v75;
        _84_v75 = _dafny.Seq.UnicodeFromString("b");
        r0 = _84_v75;
        let _85_v76;
        let _nw16 = new _module.C1();
        _nw16.__ctor();
        _85_v76 = _nw16;
        let _86_v77;
        let _out0;
        _out0 = (_85_v76).m0(p1, p0, p1, globalState);
        _86_v77 = _out0;
        let _87_v78;
        _87_v78 = _dafny.Seq.of(_84_v75);
        let _rhs13 = _87_v78;
        let _rhs14 = _83_v74;
        let _rhs15 = ((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]).plus(((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))]).multipliedBy(_86_v77));
        let _rhs16 = ((new BigNumber(405)).plus(new BigNumber(375))).isLessThanOrEqualTo((_64_v55).f6);
        _87_v78 = _rhs13;
        _83_v74 = _rhs14;
        _86_v77 = _rhs15;
        r3 = _rhs16;
        let _88_v79;
        _88_v79 = _dafny.Map.Empty.slice().updateUnsafe(_64_v55,_64_v55);
        _64_v55 = ((p1) ? ((((_88_v79).contains(_64_v55)) ? ((_88_v79).get(_64_v55)) : (_64_v55))) : (_64_v55));
      } else {
        let _89_v80;
        let _init1 = ((_90_p1) => function (_91_i1) {
          return _90_p1;
        })(p1);
        let _nw17 = Array((new BigNumber(4)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw17.length); _i0_1++) {
          _nw17[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _89_v80 = _nw17;
        let _index9 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
        (_89_v80)[_index9] = p1;
        let _index10 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
        (_89_v80)[_index10] = !(p1);
        let _92_v81;
        _92_v81 = _dafny.Set.fromElements(_89_v80);
        _92_v81 = (_92_v81).Intersect(_dafny.Set.fromElements(_89_v80, _89_v80, _89_v80));
        let _93_v82;
        _93_v82 = new _dafny.CodePoint('a'.codePointAt(0));
        _93_v82 = p0;
        _61_v52 = (_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))];
        if ((_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))]) {
          let _94_v83;
          _94_v83 = _dafny.Set.fromElements(_93_v82);
          let _index11 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          let _index12 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          let _rhs17 = (_94_v83).IsSubsetOf(_94_v83);
          let _rhs18 = (((_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))]) ? (false) : ((_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))]));
          let _lhs4 = _89_v80;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          let _lhs6 = _89_v80;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          _lhs4[_lhs5] = _rhs17;
          _lhs6[_lhs7] = _rhs18;
          let _95_v84;
          _95_v84 = _dafny.Seq.of(p1, true, p1, p1);
          let _96_v85;
          _96_v85 = _dafny.MultiSet.fromElements((_64_v55).f6);
          let _97_v86;
          _97_v86 = _dafny.Seq.of(new BigNumber((_96_v85).cardinality()));
          let _98_v87;
          _98_v87 = _dafny.Set.fromElements(!(p1), _dafny.Seq.contains(_97_v86, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))],(_95_v84)[_module.__default.safeIndex((_58_v51)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_58_v51).length))], new BigNumber((_95_v84).length))])).length)), p1);
          (globalState).f5 = _98_v87;
          let _99_v88;
          let _nw18 = new _module.C1();
          _nw18.__ctor();
          _99_v88 = _nw18;
          let _100_v89;
          let _init2 = ((_101_v86) => function (_102_i2) {
            return _101_v86;
          })(_97_v86);
          let _nw19 = Array((new BigNumber(27)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw19.length); _i0_2++) {
            _nw19[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _100_v89 = _nw19;
          let _index13 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_100_v89).length));
          (_100_v89)[_index13] = _97_v86;
          let _index14 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_100_v89).length));
          (_100_v89)[_index14] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(937)), ((_103_v52) => function (_104_i3) {
            return new BigNumber((_dafny.Set.fromElements(_103_v52)).length);
          })(_61_v52));
          let _105_v90;
          _105_v90 = _module.D4.create_DC9(_89_v80);
          let _106_v91;
          _106_v91 = _dafny.Map.Empty.slice().updateUnsafe(_105_v90,_93_v82);
          _93_v82 = (((_106_v91).contains(_module.D4.create_DC9(_89_v80))) ? ((_106_v91).get(_module.D4.create_DC9(_89_v80))) : (p0));
        } else {
          let _107_v92;
          let _nw20 = new _module.C0();
          _nw20.__ctor((_83_v74.f8).Merge(_83_v74.f8), _61_v52, p1);
          _107_v92 = _nw20;
          let _108_v93;
          let _nw21 = new _module.C1();
          _nw21.__ctor();
          _108_v93 = _nw21;
          let _109_v94;
          _109_v94 = _dafny.Map.Empty.slice().updateUnsafe((_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))],_108_v93);
          let _110_v95;
          _110_v95 = _dafny.MultiSet.fromElements(p1, (_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))], p1);
          let _index15 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          (_89_v80)[_index15] = !((_83_v74).fm1((_64_v55).f6, false, new BigNumber((_109_v94).length), _110_v95, globalState)) || (p1);
          let _111_v96;
          _111_v96 = _dafny.Seq.of((_dafny.ZERO).minus((_64_v55).f6));
          let _112_v97;
          _112_v97 = _dafny.Seq.of((_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))], _dafny.areEqual(_111_v96, _111_v96), p1);
          let _index16 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          let _rhs19 = _108_v93;
          let _rhs20 = !((_112_v97)[_module.__default.safeIndex((_64_v55).f6, new BigNumber((_112_v97).length))]);
          let _rhs21 = (_89_v80)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length))];
          let _lhs8 = _89_v80;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_89_v80).length));
          _108_v93 = _rhs19;
          _lhs8[_lhs9] = _rhs20;
          r3 = _rhs21;
          let _113_v98;
          let _out1;
          _out1 = (_108_v93).m0(!(true), _93_v82, (_64_v55).f7, globalState);
          _113_v98 = _out1;
          let _114_v99;
          _114_v99 = _dafny.Seq.UnicodeFromString("aqwo");
          r0 = _114_v99;
        }
      }
      let _115_v100;
      let _init3 = ((_116_p1) => function (_117_i4) {
        return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_116_p1,_116_p1), _dafny.Map.Empty.slice().updateUnsafe(_116_p1,!(_116_p1)), _dafny.Map.Empty.slice().updateUnsafe(_116_p1,_116_p1))).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(_116_p1),_116_p1)));
      })(p1);
      let _nw22 = Array((new BigNumber(5)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw22.length); _i0_3++) {
        _nw22[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _115_v100 = _nw22;
      let _118_v101;
      _118_v101 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p1,false));
      let _index17 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_115_v100).length));
      (_115_v100)[_index17] = ((p1) ? (_118_v101) : (_118_v101));
      let _index18 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_115_v100).length));
      (_115_v100)[_index18] = _118_v101;
      let _119_v102;
      _119_v102 = _dafny.Seq.UnicodeFromString("pkeaema");
      r0 = _119_v102;
      r1 = _83_v74;
      r2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-574)), ((_120_p1, _121_v52) => function (_122_i5) {
        return function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(146), new BigNumber(220))) {
            let _123_v103 = _compr_5;
            if (((new BigNumber(146)).isLessThanOrEqualTo(_123_v103)) && ((_123_v103).isLessThan(new BigNumber(220)))) {
              _coll5.push([(_123_v103).minus(new BigNumber((_dafny.Set.fromElements(_120_p1)).length)),_121_v52]);
            }
          }
          return _coll5;
        }();
      })(p1, _61_v52));
      r3 = p1;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _124_v0;
      _124_v0 = new BigNumber(197);
      let _125_v1;
      _125_v1 = _dafny.MultiSet.fromElements(_124_v0);
      let _126_v2;
      let _nw23 = Array((new BigNumber(13)).toNumber());
      _nw23[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_124_v0);
      _nw23[(_dafny.ONE).toNumber()] = _125_v1;
      _nw23[(new BigNumber(2)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(3)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(4)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_124_v0, _124_v0, new BigNumber(265));
      _nw23[(new BigNumber(6)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(7)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(8)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(9)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(10)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(11)).toNumber()] = _125_v1;
      _nw23[(new BigNumber(12)).toNumber()] = _125_v1;
      _126_v2 = _nw23;
      let _127_v3;
      let _nw24 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _127_v3 = _nw24;
      let _128_v4;
      _128_v4 = _dafny.Map.Empty.slice().updateUnsafe(_124_v0,_127_v3);
      let _129_v5;
      _129_v5 = false;
      let _130_v6;
      _130_v6 = _dafny.Set.fromElements(_129_v5, _129_v5, _129_v5, _129_v5);
      let _131_globalState;
      let _nw25 = new _module.GlobalState();
      _nw25.__ctor(_126_v2, _126_v2, true, _128_v4, new _dafny.CodePoint('x'.codePointAt(0)), (_130_v6).Union(_130_v6));
      _131_globalState = _nw25;
      let _hi0 = _124_v0;
      for (let _132_i0 = (_124_v0).plus(new BigNumber(302)); _132_i0.isLessThan(_hi0); _132_i0 = _132_i0.plus(_dafny.ONE)) {
        let _133_v7;
        _133_v7 = _dafny.Seq.UnicodeFromString("a");
        let _134_v8;
        _134_v8 = _dafny.Seq.of(_129_v5, _129_v5);
        let _135_v9;
        _135_v9 = new _dafny.CodePoint('b'.codePointAt(0));
        let _136_v10;
        _136_v10 = _dafny.Set.fromElements(_135_v9);
        let _137_v11;
        _137_v11 = _129_v5;
        let _138_v12;
        _138_v12 = _dafny.Map.Empty.slice().updateUnsafe(_135_v9,false);
        let _139_v14;
        let _nw26 = Array((new BigNumber(21)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _129_v5;
        _nw26[(_dafny.ONE).toNumber()] = _129_v5;
        _nw26[(new BigNumber(2)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(3)).toNumber()] = !_dafny.areEqual(_133_v7, _dafny.Seq.update(_133_v7, _module.__default.safeIndex(new BigNumber((_134_v8).length), new BigNumber((_133_v7).length)), _135_v9));
        _nw26[(new BigNumber(4)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(5)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(6)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(7)).toNumber()] = (_124_v0).isEqualTo(_132_i0);
        _nw26[(new BigNumber(8)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(9)).toNumber()] = (_132_i0).isLessThanOrEqualTo(new BigNumber((_136_v10).length));
        _nw26[(new BigNumber(10)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(11)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(12)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(13)).toNumber()] = (_137_v11);
        _nw26[(new BigNumber(14)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(15)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(16)).toNumber()] = !(_dafny.MultiSet.fromElements(_129_v5)).contains(_129_v5);
        _nw26[(new BigNumber(17)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(18)).toNumber()] = (_136_v10).IsProperSubsetOf(function () {
          let _coll6 = new _dafny.Set();
          for (const _compr_6 of (_138_v12).Keys.Elements) {
            let _140_v13 = _compr_6;
            if ((_138_v12).contains(_140_v13)) {
              _coll6.add(_140_v13);
            }
          }
          return _coll6;
        }());
        _nw26[(new BigNumber(19)).toNumber()] = _129_v5;
        _nw26[(new BigNumber(20)).toNumber()] = _129_v5;
        _139_v14 = _nw26;
        let _141_v15;
        _141_v15 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("b"));
        let _142_v16;
        _142_v16 = _dafny.Seq.of(new BigNumber(-894), (_dafny.ZERO).minus(_124_v0));
        let _143_v17;
        _143_v17 = _dafny.Set.fromElements((((_137_v11)) ? (_132_i0) : (new BigNumber((_134_v8).length))), _module.__default.safeDivisionInt(_124_v0, new BigNumber(((_125_v1).update(new BigNumber(320), _module.__default.abs((_142_v16)[_module.__default.safeIndex(_124_v0, new BigNumber((_142_v16).length))]))).cardinality())));
        let _144_v18;
        _144_v18 = _dafny.Map.Empty.slice().updateUnsafe(_142_v16,_139_v14);
        let _145_v19;
        _145_v19 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_130_v6).length),_124_v0));
        let _146_v20;
        _146_v20 = _dafny.Seq.of(_143_v17);
        let _rhs22 = _129_v5;
        let _rhs23 = (((_144_v18).contains(_dafny.Seq.Concat(_dafny.Seq.update(_142_v16, _module.__default.safeIndex(new BigNumber(-122), new BigNumber((_142_v16).length)), new BigNumber(((_145_v19)[_module.__default.safeIndex(_124_v0, new BigNumber((_145_v19).length))]).length)), _142_v16))) ? ((_144_v18).get(_dafny.Seq.Concat(_dafny.Seq.update(_142_v16, _module.__default.safeIndex(new BigNumber(-122), new BigNumber((_142_v16).length)), new BigNumber(((_145_v19)[_module.__default.safeIndex(_124_v0, new BigNumber((_145_v19).length))]).length)), _142_v16))) : (_139_v14));
        let _rhs24 = _141_v15;
        let _rhs25 = (_143_v17).Intersect((_146_v20)[_module.__default.safeIndex(_132_i0, new BigNumber((_146_v20).length))]);
        _129_v5 = _rhs22;
        _139_v14 = _rhs23;
        _141_v15 = _rhs24;
        _143_v17 = _rhs25;
        let _147_v21;
        let _nw27 = new _module.C1();
        _nw27.__ctor();
        _147_v21 = _nw27;
        _129_v5 = _129_v5;
        _124_v0 = _124_v0;
      }
      let _hi1 = _124_v0;
      for (let _148_i1 = new BigNumber(71); _148_i1.isLessThan(_hi1); _148_i1 = _148_i1.plus(_dafny.ONE)) {
        let _149_v22;
        _149_v22 = new _dafny.CodePoint('r'.codePointAt(0));
        let _150_v23;
        _150_v23 = _dafny.Seq.of(_148_i1);
        let _151_v24;
        _151_v24 = _module.D1.create_DC2(_149_v22, _124_v0, new BigNumber((_150_v23).length));
        let _152_v25;
        _152_v25 = _module.D3.create_DC7(_151_v24, _129_v5);
        _129_v5 = (_152_v25).dtor_cf15;
        let _153_v26;
        _153_v26 = _dafny.Seq.UnicodeFromString("dbh");
        let _154_v27;
        _154_v27 = _dafny.MultiSet.fromElements(_153_v26, _dafny.Seq.update(_153_v26, _module.__default.safeIndex(new BigNumber((_150_v23).length), new BigNumber((_153_v26).length)), _149_v22), _153_v26);
        _154_v27 = _154_v27;
        let _155_v28;
        let _nw28 = new _module.C1();
        _nw28.__ctor();
        _155_v28 = _nw28;
        _155_v28 = _155_v28;
        let _156_v29;
        _156_v29 = !(_129_v5);
        let _157_v30;
        let _nw29 = new _module.C0();
        _nw29.__ctor(_module.__default.fm5(_131_globalState), (_148_i1).plus(_124_v0), _156_v29);
        _157_v30 = _nw29;
      }
      let _158_v31;
      _158_v31 = _dafny.Seq.UnicodeFromString("nt");
      let _hi2 = _124_v0;
      for (let _159_i2 = new BigNumber((_158_v31).length); _159_i2.isLessThan(_hi2); _159_i2 = _159_i2.plus(_dafny.ONE)) {
        let _160_v32;
        _160_v32 = _module.D2.create_DC4(_129_v5, _124_v0);
        let _161_v33;
        let _nw30 = Array((new BigNumber(8)).toNumber());
        _nw30[(_dafny.ZERO).toNumber()] = _129_v5;
        _nw30[(_dafny.ONE).toNumber()] = _129_v5;
        _nw30[(new BigNumber(2)).toNumber()] = _129_v5;
        _nw30[(new BigNumber(3)).toNumber()] = _129_v5;
        _nw30[(new BigNumber(4)).toNumber()] = _129_v5;
        _nw30[(new BigNumber(5)).toNumber()] = _129_v5;
        _nw30[(new BigNumber(6)).toNumber()] = (_160_v32).dtor_cf6;
        _nw30[(new BigNumber(7)).toNumber()] = _129_v5;
        _161_v33 = _nw30;
        let _index19 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length));
        (_161_v33)[_index19] = (_159_i2).isLessThan(_159_i2);
        let _162_v34;
        _162_v34 = new _dafny.CodePoint('u'.codePointAt(0));
        let _163_v35;
        _163_v35 = _dafny.Map.Empty.slice().updateUnsafe(_159_i2,_159_i2);
        let _164_v36;
        _164_v36 = _dafny.Set.fromElements(_158_v31);
        let _165_v37;
        _165_v37 = _module.D2.create_DC5(((_129_v5) ? (_module.__default.fm6(_module.__default.fm3(_162_v34, _129_v5, _129_v5, _131_globalState), _129_v5, _159_i2, _131_globalState)) : (_163_v35)), (_module.__default.fm7(_129_v5, _129_v5, _131_globalState)).IsDisjointFrom(_164_v36), _129_v5, _129_v5, _129_v5);
        let _166_v38;
        _166_v38 = _dafny.Seq.of(_129_v5);
        let _167_v39;
        _167_v39 = _dafny.Map.Empty.slice().updateUnsafe((_166_v38)[_module.__default.safeIndex(new BigNumber((_166_v38).length), new BigNumber((_166_v38).length))],_129_v5);
        let _index20 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length));
        let _rhs26 = !(_129_v5);
        let _rhs27 = ((((_167_v39).contains(false)) ? ((_167_v39).get(false)) : (!(_129_v5)))) === ((_module.__default.fm3(_162_v34, _129_v5, _129_v5, _131_globalState)).isLessThanOrEqualTo(_159_i2));
        let _rhs28 = _module.D2.create_DC5(_163_v35, _129_v5, _129_v5, _129_v5, _129_v5);
        let _lhs10 = _161_v33;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length));
        _129_v5 = _rhs26;
        _lhs10[_lhs11] = _rhs27;
        _165_v37 = _rhs28;
        let _168_v40;
        _168_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,_127_v3);
        let _pat_let_tv0 = _162_v34;
        let _source0 = function (_pat_let0_0) {
          return function (_169_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_170_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_170_dt__update_hcf2_h0, (_169_dt__update__tmp_h0).dtor_cf3, (_169_dt__update__tmp_h0).dtor_cf4);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }((((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))]) ? (_module.D1.create_DC2(_162_v34, _124_v0, new BigNumber((_168_v40).length))) : (_module.D1.create_DC2(_162_v34, new BigNumber((_158_v31).length), (_dafny.ZERO).minus(_124_v0)))));
        if (_source0.is_DC2) {
          let _171___mcc_h0 = (_source0).cf2;
          let _172___mcc_h1 = (_source0).cf3;
          let _173___mcc_h2 = (_source0).cf4;
          let _174_cf4 = _173___mcc_h2;
          let _175_cf3 = _172___mcc_h1;
          let _176_cf2 = _171___mcc_h0;
          let _177_v41;
          _177_v41 = _dafny.Seq.of(_124_v0, _159_i2);
          let _178_v42;
          _178_v42 = _dafny.Map.Empty.slice().updateUnsafe((_166_v38)[_module.__default.safeIndex(_124_v0, new BigNumber((_166_v38).length))],_dafny.MultiSet.FromArray(_177_v41));
          let _179_v43;
          _179_v43 = _129_v5;
          let _180_v44;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_178_v42, new BigNumber((_158_v31).length), _179_v43);
          _180_v44 = _nw31;
          _180_v44 = _180_v44;
          _127_v3 = _127_v3;
          let _index21 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length));
          (_161_v33)[_index21] = _dafny.areEqual(_158_v31, _dafny.Seq.UnicodeFromString("nwf"));
          let _181_v45;
          _181_v45 = _dafny.Map.Empty.slice().updateUnsafe(_162_v34,_129_v5);
          _175_cf3 = new BigNumber(((_181_v45).Merge(_181_v45)).length);
        } else {
          let _182___mcc_h3 = (_source0).cf1;
          let _183_cf1 = _182___mcc_h3;
          _129_v5 = (_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))];
          let _184_v46;
          let _nw32 = new _module.C1();
          _nw32.__ctor();
          _184_v46 = _nw32;
          let _185_v47;
          let _nw33 = Array((new BigNumber(15)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = _184_v46;
          _nw33[(_dafny.ONE).toNumber()] = _184_v46;
          _nw33[(new BigNumber(2)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(3)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(4)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(5)).toNumber()] = (((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))]) ? (_184_v46) : (_184_v46));
          _nw33[(new BigNumber(6)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(7)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(8)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(9)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(10)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(11)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(12)).toNumber()] = _184_v46;
          _nw33[(new BigNumber(13)).toNumber()] = ((_129_v5) ? (_184_v46) : (_184_v46));
          _nw33[(new BigNumber(14)).toNumber()] = _184_v46;
          _185_v47 = _nw33;
          let _index22 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_185_v47).length));
          (_185_v47)[_index22] = _184_v46;
          let _index23 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_185_v47).length));
          (_185_v47)[_index23] = _184_v46;
          let _186_v48;
          _186_v48 = _module.D4.create_DC9(_161_v33);
          _161_v33 = (_186_v48).dtor_cf17;
          _127_v3 = _127_v3;
        }
        if (!(_129_v5)) {
          let _187_v49;
          let _nw34 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _187_v49 = _nw34;
          let _188_v50;
          _188_v50 = _dafny.Seq.of(_159_i2, _124_v0, _159_i2);
          let _index24 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_187_v49).length));
          (_187_v49)[_index24] = _dafny.Seq.of((_188_v50)[_module.__default.safeIndex(_124_v0, new BigNumber((_188_v50).length))], new BigNumber((_163_v35).length), _159_i2, new BigNumber((_125_v1).cardinality()));
          let _index25 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_187_v49).length));
          (_187_v49)[_index25] = _188_v50;
          _129_v5 = !(!(_module.__default.fm8(_124_v0, _163_v35, (((_125_v1).contains(new BigNumber(254))) ? ((_125_v1).get(new BigNumber(254))) : (new BigNumber((_166_v38).length))), _129_v5, _131_globalState)) || (false));
          let _index26 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length));
          (_161_v33)[_index26] = (((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))]) ? (_dafny.Seq.IsPrefixOf(_dafny.Seq.of((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))]), _dafny.Seq.of(_129_v5))) : ((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))]));
          let _189_v51;
          _189_v51 = _dafny.Map.Empty.slice().updateUnsafe((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))],(_125_v1).Difference(_dafny.MultiSet.fromElements(_124_v0)));
          _189_v51 = (_189_v51).update((_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))], _module.__default.fm2(_131_globalState));
          let _190_v52;
          _190_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("lpuwc"),_module.__default.fm3(_162_v34, (_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))], _129_v5, _131_globalState));
          let _191_v53;
          _191_v53 = (_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))];
          let _192_v54;
          let _nw35 = new _module.C0();
          _nw35.__ctor(_189_v51, new BigNumber((_125_v1).cardinality()), _191_v53);
          _192_v54 = _nw35;
          let _193_v55;
          _193_v55 = _dafny.Map.Empty.slice().updateUnsafe(_192_v54,_129_v5);
          _190_v52 = (_190_v52).update(_dafny.Seq.UnicodeFromString("bo"), new BigNumber(((_193_v55).update(_192_v54, (_161_v33)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length))])).length));
        } else {
          _124_v0 = _124_v0;
          let _194_v56;
          _194_v56 = _dafny.Seq.of(_127_v3);
          let _195_v57;
          _195_v57 = _module.D5.create_DC11(_194_v56);
          let _index27 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_161_v33).length));
          (_161_v33)[_index27] = _module.__default.fm8(_159_i2, _163_v35, _124_v0, _module.__default.fm8(new BigNumber(((_195_v57).dtor_cf18).length), (_163_v35).update(_159_i2, _124_v0), _159_i2, _129_v5, _131_globalState), _131_globalState);
          _194_v56 = _dafny.Seq.of(_127_v3, _127_v3, _127_v3, _127_v3, _127_v3);
          _124_v0 = _159_i2;
          let _196_v58;
          let _nw36 = new _module.C1();
          _nw36.__ctor();
          _196_v58 = _nw36;
        }
        _129_v5 = _module.__default.fm8(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_197_v34) => function (_198_i3) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_199_v34) => function (_200_i4) {
            return _199_v34;
          })(_197_v34))).length);
        })(_162_v34))).length), _163_v35, _159_i2, _129_v5, _131_globalState);
      }
      if (_129_v5) {
        let _201_v59;
        let _init4 = ((_202_v5) => function (_203_i5) {
          return _202_v5;
        })(_129_v5);
        let _nw37 = Array((new BigNumber(19)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw37.length); _i0_4++) {
          _nw37[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _201_v59 = _nw37;
        let _204_v60;
        _204_v60 = new _dafny.CodePoint('m'.codePointAt(0));
        let _205_v61;
        _205_v61 = _dafny.Map.Empty.slice().updateUnsafe(_204_v60,_124_v0);
        let _index28 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_201_v59).length));
        (_201_v59)[_index28] = !((_125_v1).IsDisjointFrom(_dafny.MultiSet.fromElements(_124_v0, _124_v0, _124_v0, new BigNumber((_205_v61).length))));
        let _index29 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_201_v59).length));
        (_201_v59)[_index29] = _129_v5;
        let _206_v62;
        _206_v62 = _dafny.Set.fromElements(_201_v59);
        let _index30 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_201_v59).length));
        (_201_v59)[_index30] = (_dafny.Set.fromElements(_201_v59, _201_v59)).IsSubsetOf((_206_v62).Intersect(_206_v62));
        _124_v0 = new BigNumber(705);
        let _207_v63;
        _207_v63 = _dafny.Map.Empty.slice().updateUnsafe(true,_125_v1);
        let _208_v64;
        _208_v64 = (_201_v59)[_module.__default.safeIndex(new BigNumber(594), new BigNumber((_201_v59).length))];
        let _209_v65;
        let _nw38 = new _module.C0();
        _nw38.__ctor(_207_v63, _124_v0, _208_v64);
        _209_v65 = _nw38;
        let _210_v66;
        _210_v66 = _dafny.Map.Empty.slice().updateUnsafe((_201_v59)[_module.__default.safeIndex(new BigNumber(594), new BigNumber((_201_v59).length))],_124_v0);
        let _211_v67;
        _211_v67 = _dafny.Set.fromElements(new BigNumber((_130_v6).length), new BigNumber(238), new BigNumber((_210_v66).length));
        let _212_v68;
        _212_v68 = _dafny.Seq.of(_211_v67);
        let _213_v69;
        _213_v69 = _module.D6.create_DC15(_209_v65);
        let _214_v70;
        _214_v70 = _module.D7.create_DC19(_212_v68);
        let _rhs29 = (_213_v69).dtor_cf21;
        let _rhs30 = _124_v0;
        let _rhs31 = (_214_v70).dtor_cf25;
        _209_v65 = _rhs29;
        _124_v0 = _rhs30;
        _212_v68 = _rhs31;
        let _215_v71;
        let _nw39 = new _module.C1();
        _nw39.__ctor();
        _215_v71 = _nw39;
        _215_v71 = _215_v71;
      } else {
        let _216_v72;
        _216_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_124_v0),_124_v0);
        let _217_v73;
        _217_v73 = _dafny.MultiSet.fromElements(_129_v5);
        _216_v72 = (_216_v72).update((new BigNumber((_217_v73).cardinality())).plus(_124_v0), (_dafny.ZERO).minus(_124_v0));
        _124_v0 = (_124_v0).minus(_module.__default.safeModuloInt(_124_v0, _124_v0));
        let _218_v74;
        _218_v74 = new _dafny.CodePoint('n'.codePointAt(0));
        _158_v31 = _dafny.Seq.update(_158_v31, _module.__default.safeIndex(new BigNumber(791), new BigNumber((_158_v31).length)), _218_v74);
        let _219_v75;
        let _init5 = ((_220_v0) => function (_221_i6) {
          return _dafny.Seq.of(_220_v0, new BigNumber(535));
        })(_124_v0);
        let _nw40 = Array((new BigNumber(16)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw40.length); _i0_5++) {
          _nw40[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _219_v75 = _nw40;
        let _222_v76;
        _222_v76 = _dafny.Seq.of(new BigNumber(117));
        let _index31 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_219_v75).length));
        (_219_v75)[_index31] = _dafny.Seq.update(_222_v76, _module.__default.safeIndex(new BigNumber(973), new BigNumber((_222_v76).length)), new BigNumber(462));
        let _index32 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_219_v75).length));
        (_219_v75)[_index32] = _222_v76;
        let _223_v77;
        let _224_v78;
        let _225_v79;
        let _226_v80;
        let _out2;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector0 = _module.__default.m1(_218_v74, _129_v5, _131_globalState);
        _out2 = _outcollector0[0];
        _out3 = _outcollector0[1];
        _out4 = _outcollector0[2];
        _out5 = _outcollector0[3];
        _223_v77 = _out2;
        _224_v78 = _out3;
        _225_v79 = _out4;
        _226_v80 = _out5;
      }
      let _227_v81;
      _227_v81 = _dafny.Map.Empty.slice().updateUnsafe(_124_v0,_129_v5);
      _227_v81 = (_227_v81).update(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_124_v0)).length), new BigNumber(623)), _129_v5);
      if (true) {
        let _228_v82;
        _228_v82 = new _dafny.CodePoint('c'.codePointAt(0));
        let _229_v83;
        let _230_v84;
        let _231_v85;
        let _232_v86;
        let _out6;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector1 = _module.__default.m1(((true) ? (_228_v82) : (_228_v82)), _129_v5, _131_globalState);
        _out6 = _outcollector1[0];
        _out7 = _outcollector1[1];
        _out8 = _outcollector1[2];
        _out9 = _outcollector1[3];
        _229_v83 = _out6;
        _230_v84 = _out7;
        _231_v85 = _out8;
        _232_v86 = _out9;
        _129_v5 = _129_v5;
        _124_v0 = _124_v0;
        let _233_v87;
        let _234_v88;
        let _235_v89;
        let _236_v90;
        let _out10;
        let _out11;
        let _out12;
        let _out13;
        let _outcollector2 = _module.__default.m1(_228_v82, _129_v5, _131_globalState);
        _out10 = _outcollector2[0];
        _out11 = _outcollector2[1];
        _out12 = _outcollector2[2];
        _out13 = _outcollector2[3];
        _233_v87 = _out10;
        _234_v88 = _out11;
        _235_v89 = _out12;
        _236_v90 = _out13;
        _236_v90 = true;
      } else {
        let _237_v91;
        _237_v91 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_238_i7) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length));
        let _239_v92;
        _239_v92 = _dafny.Map.Empty.slice().updateUnsafe(_129_v5,_dafny.MultiSet.FromArray(_237_v91));
        let _240_v93;
        _240_v93 = new _dafny.CodePoint('x'.codePointAt(0));
        let _241_v94;
        let _nw41 = new _module.C0();
        _nw41.__ctor(_239_v92, _module.__default.fm3(_240_v93, _129_v5, _129_v5, _131_globalState), false);
        _241_v94 = _nw41;
        let _242_v95;
        _242_v95 = _dafny.Seq.of(_241_v94);
        let _243_v96;
        _243_v96 = _module.D6.create_DC15(_241_v94);
        let _244_v97;
        _244_v97 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Concat(_242_v95, _242_v95), _module.__default.safeIndex(new BigNumber((_158_v31).length), new BigNumber((_dafny.Seq.Concat(_242_v95, _242_v95)).length)), (_243_v96).dtor_cf21), _242_v95, _242_v95);
        _244_v97 = _dafny.Seq.of(_242_v95);
        let _index33 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_127_v3).length));
        (_127_v3)[_index33] = (_124_v0).minus(_124_v0);
        let _index34 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_127_v3).length));
        (_127_v3)[_index34] = (_124_v0).multipliedBy(_124_v0);
        let _245_v98;
        _245_v98 = _dafny.Seq.of(false, _129_v5, _129_v5);
        let _246_v99;
        _246_v99 = _dafny.Seq.of(_245_v98, _245_v98, _dafny.Seq.update(_245_v98, _module.__default.safeIndex((_127_v3)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_127_v3).length))], new BigNumber((_245_v98).length)), _129_v5), _245_v98, _245_v98);
        let _247_v100;
        _247_v100 = false;
        let _248_v101;
        let _nw42 = new _module.C0();
        _nw42.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_129_v5,_dafny.MultiSet.fromElements((_127_v3)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_127_v3).length))], new BigNumber((_dafny.Seq.of((_127_v3)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_127_v3).length))], _124_v0)).length))), new BigNumber((_246_v99).length), !(_129_v5));
        _248_v101 = _nw42;
        let _249_v103;
        _249_v103 = _dafny.Set.fromElements(_227_v81, function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(930), new BigNumber(440))) {
            let _250_v102 = _compr_7;
            if (((new BigNumber(930)).isLessThanOrEqualTo(_250_v102)) && ((_250_v102).isLessThan(new BigNumber(440)))) {
              _coll7.push([(_250_v102).multipliedBy((_127_v3)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_127_v3).length))]),true]);
            }
          }
          return _coll7;
        }());
        let _251_v104;
        _251_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_245_v98).length),_249_v103);
        _251_v104 = (_251_v104).update(new BigNumber(-48), _249_v103);
        let _252_v105;
        let _init6 = ((_253_v5) => function (_254_i8) {
          return (_dafny.Map.Empty.slice().updateUnsafe(!(_253_v5),_253_v5)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(_253_v5)));
        })(_129_v5);
        let _nw43 = Array((new BigNumber(27)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw43.length); _i0_6++) {
          _nw43[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _252_v105 = _nw43;
        let _255_v106;
        _255_v106 = _dafny.Map.Empty.slice().updateUnsafe(!(_129_v5),_129_v5);
        let _index35 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_252_v105).length));
        (_252_v105)[_index35] = (_255_v106).Merge(_dafny.Map.Empty.slice().updateUnsafe(_129_v5,_129_v5));
        let _index36 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_252_v105).length));
        (_252_v105)[_index36] = _dafny.Map.Empty.slice().updateUnsafe(false,_129_v5);
      }
      let _hi3 = _124_v0;
      for (let _256_i9 = _124_v0; _256_i9.isLessThan(_hi3); _256_i9 = _256_i9.plus(_dafny.ONE)) {
        let _257_v107;
        _257_v107 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_125_v1);
        let _258_v108;
        _258_v108 = _129_v5;
        let _259_v109;
        let _nw44 = new _module.C0();
        _nw44.__ctor(_257_v107, (_124_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_158_v31).length))), _258_v108);
        _259_v109 = _nw44;
        _259_v109 = _259_v109;
        let _260_v110;
        let _nw45 = Array((new BigNumber(13)).toNumber()).fill(false);
        _260_v110 = _nw45;
        let _index37 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_260_v110).length));
        (_260_v110)[_index37] = _129_v5;
        let _261_v111;
        _261_v111 = _dafny.MultiSet.fromElements(!(false) || (_129_v5), _129_v5);
        let _index38 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_260_v110).length));
        let _rhs32 = _124_v0;
        let _rhs33 = (_125_v1).Intersect(_125_v1);
        let _rhs34 = _129_v5;
        let _rhs35 = (((_261_v111).contains(false)) ? ((_261_v111).get(false)) : (_124_v0));
        let _lhs12 = _260_v110;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_260_v110).length));
        _124_v0 = _rhs32;
        _125_v1 = _rhs33;
        _lhs12[_lhs13] = _rhs34;
        _124_v0 = _rhs35;
        let _index39 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_260_v110).length));
        (_260_v110)[_index39] = !(new BigNumber(((_261_v111).Intersect(_261_v111)).cardinality())).isEqualTo(_256_i9);
        let _262_v112;
        let _nw46 = new _module.C1();
        _nw46.__ctor();
        _262_v112 = _nw46;
      }
      let _263_v113;
      let _nw47 = Array((new BigNumber(28)).toNumber()).fill(false);
      _263_v113 = _nw47;
      let _index40 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length));
      (_263_v113)[_index40] = true;
      let _index41 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length));
      (_263_v113)[_index41] = _129_v5;
      let _264_v114;
      _264_v114 = _dafny.Map.Empty.slice().updateUnsafe((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))],(_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))]);
      let _265_v115;
      _265_v115 = _dafny.Seq.of(_124_v0);
      _264_v114 = (_264_v114).update((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))], _dafny.Seq.contains(_265_v115, _124_v0));
      let _266_v116;
      _266_v116 = new _dafny.CodePoint('i'.codePointAt(0));
      _266_v116 = _266_v116;
      _266_v116 = new _dafny.CodePoint('h'.codePointAt(0));
      let _267_v117;
      _267_v117 = _dafny.Map.Empty.slice().updateUnsafe(_124_v0,_124_v0);
      let _268_v118;
      _268_v118 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_267_v117).length),_124_v0);
      let _269_v119;
      _269_v119 = _dafny.Seq.of(_129_v5);
      let _hi4 = ((((_268_v118).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), ((_272_v116) => function (_273_i11) {
        return _272_v116;
      })(_266_v116))).length))) ? ((_268_v118).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), ((_270_v116) => function (_271_i11) {
        return _270_v116;
      })(_266_v116))).length))) : (_124_v0))).minus(new BigNumber((_269_v119).length));
      for (let _274_i10 = (_124_v0).multipliedBy(_124_v0); _274_i10.isLessThan(_hi4); _274_i10 = _274_i10.plus(_dafny.ONE)) {
        _129_v5 = !(_129_v5);
        if (_129_v5) {
          let _275_v120;
          _275_v120 = _dafny.Set.fromElements(_266_v116);
          let _index42 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length));
          (_263_v113)[_index42] = (new BigNumber((_275_v120).length)).isLessThanOrEqualTo((new BigNumber(-528)).minus(_274_i10));
          let _276_v121;
          let _277_v122;
          let _278_v123;
          let _279_v124;
          let _out14;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector3 = _module.__default.m1((_158_v31)[_module.__default.safeIndex((_dafny.ZERO).minus(_124_v0), new BigNumber((_158_v31).length))], (_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))], _131_globalState);
          _out14 = _outcollector3[0];
          _out15 = _outcollector3[1];
          _out16 = _outcollector3[2];
          _out17 = _outcollector3[3];
          _276_v121 = _out14;
          _277_v122 = _out15;
          _278_v123 = _out16;
          _279_v124 = _out17;
          let _280_v125;
          _280_v125 = _129_v5;
          let _281_v126;
          let _nw48 = new _module.C0();
          _nw48.__ctor(_module.__default.fm5(_131_globalState), _module.__default.safeDivisionInt(_274_i10, new BigNumber(-74)), _280_v125);
          _281_v126 = _nw48;
          _281_v126 = _281_v126;
          _124_v0 = new BigNumber(201);
          _124_v0 = (_module.__default.fm9(_274_i10, _279_v124, (_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))], new BigNumber((_276_v121).length), _131_globalState)).dtor_cf4;
        } else {
          let _282_v127;
          _282_v127 = _dafny.Map.Empty.slice().updateUnsafe((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))],_module.__default.fm3(_266_v116, _129_v5, (_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))], _131_globalState));
          _282_v127 = (_282_v127).update(!(!(_124_v0).isEqualTo(_274_i10)), new BigNumber((_269_v119).length));
          let _283_v128;
          _283_v128 = _dafny.Map.Empty.slice().updateUnsafe(_130_v6,_129_v5);
          _283_v128 = (_283_v128).update(_130_v6, _129_v5);
          _265_v115 = _dafny.Seq.Concat(_265_v115, _265_v115);
          _129_v5 = ((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))]) || ((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))]);
          let _284_v129;
          let _285_v130;
          let _286_v131;
          let _287_v132;
          let _out18;
          let _out19;
          let _out20;
          let _out21;
          let _outcollector4 = _module.__default.m1(((_129_v5) ? (_266_v116) : (_266_v116)), (_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))], _131_globalState);
          _out18 = _outcollector4[0];
          _out19 = _outcollector4[1];
          _out20 = _outcollector4[2];
          _out21 = _outcollector4[3];
          _284_v129 = _out18;
          _285_v130 = _out19;
          _286_v131 = _out20;
          _287_v132 = _out21;
        }
        let _288_v133;
        _288_v133 = _dafny.Map.Empty.slice().updateUnsafe((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))],_124_v0);
        let _289_v134;
        _289_v134 = _dafny.Map.Empty.slice().updateUnsafe(_129_v5,_125_v1);
        let _290_v135;
        _290_v135 = (_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))];
        let _291_v136;
        let _nw49 = new _module.C0();
        _nw49.__ctor(_289_v134, new BigNumber((_158_v31).length), _290_v135);
        _291_v136 = _nw49;
        let _292_v137;
        _292_v137 = _dafny.Map.Empty.slice().updateUnsafe(_291_v136,_158_v31);
        _268_v118 = (_268_v118).update(_124_v0, (((_288_v133).contains((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))])) ? ((_288_v133).get((_263_v113)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_263_v113).length))])) : (new BigNumber((_292_v137).length))));
        _124_v0 = _274_i10;
      }
      _129_v5 = _129_v5;
      let _293_v138;
      _293_v138 = _dafny.Seq.of(_127_v3, _127_v3, _127_v3);
      let _294_v139;
      _294_v139 = _module.D5.create_DC11(_293_v138);
      let _rhs36 = ((_129_v5) ? (_263_v113) : (_263_v113));
      let _rhs37 = _294_v139;
      _263_v113 = _rhs36;
      _294_v139 = _rhs37;
      _129_v5 = _129_v5;
      let _295_v140;
      let _nw50 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
      _295_v140 = _nw50;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_295_v140).length))) {
        let _296_i12 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_296_i12)) && ((_296_i12).isLessThan(new BigNumber((_295_v140).length))))) {
          (_295_v140)[(_296_i12)] = _265_v115;
        }
      }
      process.stdout.write(_dafny.toString(_124_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v1).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197), new BigNumber(197), new BigNumber(265)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v2)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_128_v4).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_129_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v6).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197), new BigNumber(197), new BigNumber(265)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f0)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197), new BigNumber(197), new BigNumber(265)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_131_globalState).f1)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(new BigNumber(197)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_131_globalState).f3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState.f5).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_158_v31).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v81).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(196),false).updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v113)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_264_v114).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_265_v115, _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_266_v116));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v117).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(196),new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v118).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_269_v119, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_293_v138).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_294_v139).dtor_cf18).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[_dafny.ONE], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(7)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(8)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(9)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_295_v140)[new BigNumber(10)], _dafny.Seq.of(new BigNumber(196)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
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
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC4(cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf8, cf9, cf10, cf11, cf12) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && this.cf10 === other.cf10 && this.cf11 === other.cf11 && this.cf12 === other.cf12;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(false, _dafny.ZERO);
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
    static create_DC7(cf14, cf15) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC6(cf13) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC8(cf16) {
      let $dt = new D3(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(_module.D1.Default(), false);
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
    static create_DC10() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D4(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf17) + ")";
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
        return other.$tag === 1 && this.cf17 === other.cf17;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10();
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
    static create_DC12(cf19) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC13() {
      let $dt = new D5(1);
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D5(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC14(cf20) {
      let $dt = new D5(3);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(false);
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
    static create_DC16(cf22) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC17(cf23) {
      let $dt = new D6(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC18(cf24) {
      let $dt = new D6(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf21) {
      let $dt = new D6(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf24 === other.cf24;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf21 === other.cf21;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(false);
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
    static create_DC20(cf26) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC21(cf27, cf28) {
      let $dt = new D7(1);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC19(cf25) {
      let $dt = new D7(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf27 === other.cf27 && this.cf28 === other.cf28;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20([]);
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
    static create_DC23(cf30, cf31, cf32, cf33) {
      let $dt = new D8(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC24(cf34, cf35, cf36) {
      let $dt = new D8(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC22(cf29) {
      let $dt = new D8(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC25(cf37) {
      let $dt = new D8(3);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(false, new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.Seq.of());
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
    static create_DC27() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC26(cf38) {
      let $dt = new D9(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC27";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf38) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC27();
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
    static create_DC29(cf40, cf41) {
      let $dt = new D10(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC28(cf39) {
      let $dt = new D10(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f5 = _dafny.Set.Empty;
      this._f0 = [];
      this._f1 = [];
      this._f2 = false;
      this._f3 = _dafny.Map.Empty;
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this.f8 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    __ctor(f8, f6, f7) {
      let _this = this;
      (_this).f8 = f8;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f6), (_this).f6), (_this).f6);
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((((true) ? ((_this).f6) : ((_this).f6))).isLessThan((_this).f6));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _297_v0;
      _297_v0 = _dafny.Seq.UnicodeFromString("mdcxtl");
      let _298_v1;
      _298_v1 = new BigNumber(362);
      let _299_v2;
      _299_v2 = _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), p1, (_297_v0)[_module.__default.safeIndex(_298_v1, new BigNumber((_297_v0).length))]);
      _299_v2 = ((_299_v2).Difference(_dafny.Set.fromElements(p1))).Union(_299_v2);
      let _300_i0;
      _300_i0 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_300_i0)) {
              break L0;
            }
            _300_i0 = (_300_i0).plus(_dafny.ONE);
            let _301_v3;
            _301_v3 = false;
            _301_v3 = (p2);
            let _302_v4;
            _302_v4 = _dafny.Map.Empty.slice().updateUnsafe(_301_v3,_module.__default.fm2(globalState));
            let _303_v5;
            _303_v5 = _module.D1.create_DC1(_302_v4);
            let _304_v6;
            let _nw51 = new _module.C0();
            _nw51.__ctor((_303_v5).dtor_cf1, _298_v1, p2);
            _304_v6 = _nw51;
            let _305_v7;
            _305_v7 = _dafny.MultiSet.fromElements(p0);
            let _306_v8;
            _306_v8 = _dafny.MultiSet.fromElements(_304_v6, _304_v6);
            if ((_304_v6).fm1(_298_v1, !((_304_v6).fm1(_298_v1, _301_v3, _298_v1, _305_v7, globalState)), _module.__default.safeModuloInt(new BigNumber((_306_v8).cardinality()), new BigNumber((_305_v7).cardinality())), (_305_v7).Intersect(_305_v7), globalState)) {
              let _307_v9;
              _307_v9 = _dafny.Seq.of((_305_v7).IsProperSubsetOf(_305_v7));
              _307_v9 = _307_v9;
              r0 = _298_v1;
              let _308_v10;
              _308_v10 = _dafny.Set.fromElements(p0);
              let _309_v11;
              let _nw52 = new _module.C0();
              _nw52.__ctor(_304_v6.f8, new BigNumber((_308_v10).length), p2);
              _309_v11 = _nw52;
              let _310_v12;
              let _nw53 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
              _310_v12 = _nw53;
              let _index43 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_310_v12).length));
              (_310_v12)[_index43] = _module.__default.safeDivisionInt(_298_v1, _298_v1);
              let _index44 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_310_v12).length));
              (_310_v12)[_index44] = _298_v1;
              let _311_v13;
              let _nw54 = Array((new BigNumber(5)).toNumber());
              _nw54[(_dafny.ZERO).toNumber()] = false;
              _nw54[(_dafny.ONE).toNumber()] = _301_v3;
              _nw54[(new BigNumber(2)).toNumber()] = false;
              _nw54[(new BigNumber(3)).toNumber()] = _301_v3;
              _nw54[(new BigNumber(4)).toNumber()] = (_dafny.Set.fromElements(!(_301_v3))).IsDisjointFrom(_308_v10);
              _311_v13 = _nw54;
              let _index45 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_311_v13).length));
              (_311_v13)[_index45] = (new BigNumber(341)).isLessThanOrEqualTo(_298_v1);
              let _index46 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_311_v13).length));
              (_311_v13)[_index46] = p0;
            } else {
              r0 = _298_v1;
              let _312_v14;
              let _nw55 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
              _312_v14 = _nw55;
              let _index47 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_312_v14).length));
              (_312_v14)[_index47] = _module.__default.safeModuloInt(_298_v1, _298_v1);
              let _index48 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_312_v14).length));
              (_312_v14)[_index48] = new BigNumber(47);
              _298_v1 = (_312_v14)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_312_v14).length))];
              let _313_v15;
              let _nw56 = new _module.C0();
              _nw56.__ctor(_304_v6.f8, _298_v1, true);
              _313_v15 = _nw56;
              let _314_v16;
              _314_v16 = _dafny.Seq.of((_312_v14)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_312_v14).length))]);
              let _315_v17;
              _315_v17 = _dafny.MultiSet.fromElements(new BigNumber((_314_v16).length));
              let _316_v18;
              _316_v18 = _dafny.Map.Empty.slice().updateUnsafe(_315_v17,(_312_v14)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_312_v14).length))]);
              let _317_v19;
              _317_v19 = _dafny.Seq.of(_301_v3);
              _316_v18 = (_316_v18).update(_dafny.MultiSet.fromElements(_298_v1, new BigNumber((_317_v19).length)), _298_v1);
            }
            let _318_v20;
            _318_v20 = _dafny.MultiSet.fromElements(_298_v1);
            let _319_v21;
            let _nw57 = new _module.C0();
            _nw57.__ctor((_dafny.Map.Empty.slice().updateUnsafe(p0,_318_v20)).Merge(_302_v4), _module.__default.safeModuloInt(_298_v1, _298_v1), p2);
            _319_v21 = _nw57;
          }
        }
      }
      let _320_v22;
      _320_v22 = false;
      _320_v22 = !(!(_module.__default.fm3(p1, _320_v22, _320_v22, globalState)).isEqualTo(_298_v1));
      let _321_i1;
      _321_i1 = _dafny.ZERO;
      L1: {
        while (p0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_321_i1)) {
              break L1;
            }
            _321_i1 = (_321_i1).plus(_dafny.ONE);
            let _322_v23;
            let _init7 = ((_323_p0) => function (_324_i2) {
              return _323_p0;
            })(p0);
            let _nw58 = Array((new BigNumber(18)).toNumber());
            for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw58.length); _i0_7++) {
              _nw58[_i0_7] = _init7(new BigNumber(_i0_7));
            }
            _322_v23 = _nw58;
            let _325_v24;
            _325_v24 = _dafny.Set.fromElements(_322_v23);
            _325_v24 = _325_v24;
            let _326_v25;
            let _nw59 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _326_v25 = _nw59;
            let _327_v26;
            _327_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_326_v25);
            let _328_v27;
            let _nw60 = Array((new BigNumber(21)).toNumber());
            _nw60[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("rymhitw")).length);
            _nw60[(_dafny.ONE).toNumber()] = _298_v1;
            _nw60[(new BigNumber(2)).toNumber()] = new BigNumber(687);
            _nw60[(new BigNumber(3)).toNumber()] = _298_v1;
            _nw60[(new BigNumber(4)).toNumber()] = new BigNumber(77);
            _nw60[(new BigNumber(5)).toNumber()] = new BigNumber((_297_v0).length);
            _nw60[(new BigNumber(6)).toNumber()] = (_298_v1).plus(_298_v1);
            _nw60[(new BigNumber(7)).toNumber()] = (_298_v1).plus(_298_v1);
            _nw60[(new BigNumber(8)).toNumber()] = _298_v1;
            _nw60[(new BigNumber(9)).toNumber()] = new BigNumber(899);
            _nw60[(new BigNumber(10)).toNumber()] = new BigNumber(397);
            _nw60[(new BigNumber(11)).toNumber()] = (_298_v1).plus(_298_v1);
            _nw60[(new BigNumber(12)).toNumber()] = _298_v1;
            _nw60[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-807)), ((_329_p1) => function (_330_i3) {
              return _329_p1;
            })(p1)), _297_v0)).length);
            _nw60[(new BigNumber(14)).toNumber()] = new BigNumber(((_327_v26).update(_320_v22, _326_v25)).length);
            _nw60[(new BigNumber(15)).toNumber()] = _298_v1;
            _nw60[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_297_v0).length), _298_v1));
            _nw60[(new BigNumber(17)).toNumber()] = new BigNumber(140);
            _nw60[(new BigNumber(18)).toNumber()] = (_298_v1).plus(_298_v1);
            _nw60[(new BigNumber(19)).toNumber()] = (_298_v1).minus(_298_v1);
            _nw60[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_298_v1);
            _328_v27 = _nw60;
            let _index49 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length));
            (_326_v25)[_index49] = (_298_v1).multipliedBy(new BigNumber(56));
            let _index50 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length));
            (_326_v25)[_index50] = (_298_v1).multipliedBy((new BigNumber(-528)).plus(_298_v1));
            if ((_298_v1).isLessThan((new BigNumber((_dafny.Set.fromElements((_326_v25)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length))], _298_v1)).length)).multipliedBy(new BigNumber(((_module.D2.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_320_v22,_298_v1))).dtor_cf5).length)))) {
              _320_v22 = p0;
              let _index51 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length));
              (_326_v25)[_index51] = (_326_v25)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length))];
              let _331_v28;
              _331_v28 = _module.D1.create_DC2(p1, new BigNumber(-281), _298_v1);
              let _332_v29;
              _332_v29 = _dafny.MultiSet.fromElements(_331_v28, _331_v28, _331_v28, _331_v28);
              _320_v22 = (_332_v29).IsSubsetOf(_dafny.MultiSet.fromElements(_331_v28));
              let _333_v30;
              _333_v30 = _dafny.MultiSet.fromElements((_326_v25)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length))]);
              let _334_v31;
              _334_v31 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_320_v22,_333_v30));
              let _335_v32;
              _335_v32 = _dafny.Seq.of(_320_v22);
              let _336_v33;
              _336_v33 = _dafny.Seq.of(_335_v32);
              let _337_v34;
              let _nw61 = new _module.C0();
              _nw61.__ctor((_334_v31)[_module.__default.safeIndex(new BigNumber((_336_v33).length), new BigNumber((_334_v31).length))], (_326_v25)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length))], p2);
              _337_v34 = _nw61;
              let _338_v35;
              let _nw62 = new _module.C0();
              _nw62.__ctor(_dafny.Map.Empty.slice().updateUnsafe(p0,_333_v30), (_326_v25)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length))], p2);
              _338_v35 = _nw62;
            } else {
              let _339_v36;
              _339_v36 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_298_v1)),(_dafny.ZERO).minus(new BigNumber(-582)));
              let _340_v37;
              _340_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_module.__default.fm4(_339_v36, p1, p0, _320_v22, globalState)).length));
              let _341_v38;
              _341_v38 = _module.D2.create_DC3(_340_v37);
              _341_v38 = _341_v38;
              let _342_v39;
              _342_v39 = _dafny.Seq.of(_320_v22, p0, _320_v22, _320_v22, p0);
              _340_v37 = (_340_v37).update(p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_342_v39).length), (_dafny.ZERO).minus(new BigNumber((_297_v0).length))))));
              _320_v22 = _320_v22;
              let _343_v40;
              _343_v40 = _dafny.Set.fromElements(p0, p0);
              let _index52 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_322_v23).length));
              (_322_v23)[_index52] = (_343_v40).IsProperSubsetOf(_343_v40);
              let _index53 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_322_v23).length));
              (_322_v23)[_index53] = (_320_v22) || (p0);
              let _index54 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length));
              (_326_v25)[_index54] = (_326_v25)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_326_v25).length))];
            }
            r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_298_v1));
          }
        }
      }
      _298_v1 = (_module.D2.create_DC4(p0, _298_v1)).dtor_cf7;
      let _344_v41;
      _344_v41 = _dafny.Set.fromElements(new BigNumber(980));
      let _345_v42;
      _345_v42 = _dafny.Seq.of(_298_v1, _298_v1);
      let _346_v43;
      let _nw63 = Array((new BigNumber(23)).toNumber());
      _nw63[(_dafny.ZERO).toNumber()] = _298_v1;
      _nw63[(_dafny.ONE).toNumber()] = _298_v1;
      _nw63[(new BigNumber(2)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(3)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(4)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(5)).toNumber()] = new BigNumber(107);
      _nw63[(new BigNumber(6)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(7)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(8)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(9)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(10)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(11)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(12)).toNumber()] = new BigNumber((_344_v41).length);
      _nw63[(new BigNumber(13)).toNumber()] = (_345_v42)[_module.__default.safeIndex(_298_v1, new BigNumber((_345_v42).length))];
      _nw63[(new BigNumber(14)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(15)).toNumber()] = new BigNumber(696);
      _nw63[(new BigNumber(16)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(17)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(18)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(19)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(20)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(21)).toNumber()] = _298_v1;
      _nw63[(new BigNumber(22)).toNumber()] = _298_v1;
      _346_v43 = _nw63;
      let _347_v44;
      _347_v44 = _dafny.MultiSet.fromElements(_346_v43);
      let _rhs38 = _297_v0;
      let _rhs39 = ((_module.D3.create_DC6(_347_v44)).dtor_cf13).Intersect(_347_v44);
      let _rhs40 = !(_320_v22);
      _297_v0 = _rhs38;
      _347_v44 = _rhs39;
      _320_v22 = _rhs40;
      r0 = _module.__default.safeDivisionInt(_298_v1, (_298_v1).minus(_298_v1));
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
