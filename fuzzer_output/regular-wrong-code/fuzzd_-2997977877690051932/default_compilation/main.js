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
    static fm0(p0, p1, p2, p3, globalState) {
      return ((_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(true, !(false))).length)).plus(new BigNumber(559)))).isLessThan(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(false, false, false), (_dafny.Seq.of(true)))).length));
    };
    static fm2(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(false, false, false, false, true))).Union((_dafny.Set.fromElements(!(false))).Union(_dafny.Set.fromElements(true, !(true))));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("uybruf")).length)), new BigNumber(695), new BigNumber(-92), new BigNumber((_dafny.Seq.UnicodeFromString("txqeg")).length), new BigNumber(559)), _dafny.Seq.of(new BigNumber(657), new BigNumber(-678)));
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-387)),(_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(93), new BigNumber((_dafny.Seq.UnicodeFromString("xuqnfvfjl")).length))));
    };
    static fm5(p0, globalState) {
      return _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),false));
    };
    static fm6(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("trfmv")).length))),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ataulj")).length)),true));
    };
    static fm7(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_0_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mclcxtfx"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), function (_1_i1) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _module.__default.safeDivisionInt(new BigNumber(161), new BigNumber((((true) ? (_dafny.Seq.UnicodeFromString("wvcutvx")) : (_dafny.Seq.UnicodeFromString("entn")))).length));
    };
    static fm9(p0, p1, p2, globalState) {
      return (((!(false)) ? (_dafny.Map.Empty.slice().updateUnsafe(!(false),true)) : (_dafny.Map.Empty.slice().updateUnsafe(false,!(false))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false)));
    };
    static fm10(p0, p1, globalState) {
      if (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("h"), _dafny.Seq.UnicodeFromString("ehe"))) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = false;
      let r1 = false;
      let _2_v0;
      _2_v0 = false;
      let _3_i0;
      _3_i0 = _dafny.ZERO;
      L0: {
        while (_2_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3_i0)) {
              break L0;
            }
            _3_i0 = (_3_i0).plus(_dafny.ONE);
            let _4_v1;
            let _nw0 = new _module.C0();
            _nw0.__ctor(p1);
            _4_v1 = _nw0;
            let _5_v2;
            _5_v2 = _module.D3.create_DC8(_module.D3.create_DC6(_4_v1));
            let _source0 = ((true) ? (_5_v2) : (_5_v2));
            if (_source0.is_DC7) {
              let _6_v3;
              _6_v3 = _dafny.Set.fromElements(_2_v0, _2_v0);
              let _7_v4;
              _7_v4 = _dafny.Set.fromElements(_2_v0);
              (globalState).f6 = ((_7_v4)).IsSubsetOf(_6_v3);
              let _8_v5;
              let _nw1 = new _module.C0();
              _nw1.__ctor((_4_v1).f22);
              _8_v5 = _nw1;
              let _9_v6;
              _9_v6 = _dafny.Seq.UnicodeFromString("ack");
              let _10_v7;
              _10_v7 = _module.D0.create_DC1(new BigNumber((_9_v6).length), _2_v0);
              let _11_v8;
              _11_v8 = _dafny.Seq.of((new BigNumber(693)).isLessThanOrEqualTo((_10_v7).dtor_cf1), (_6_v3).IsDisjointFrom(_6_v3));
              _11_v8 = _dafny.Seq.Concat(_11_v8, _dafny.Seq.of(false, _2_v0));
              (globalState).f6 = !(_2_v0);
            } else if (_source0.is_DC6) {
              let _12___mcc_h0 = (_source0).cf8;
              let _13_cf8 = _12___mcc_h0;
              let _14_v9;
              _14_v9 = _dafny.Seq.UnicodeFromString("jyy");
              _14_v9 = _14_v9;
              (globalState).f6 = !(_2_v0);
              let _15_v10;
              _15_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-222),_2_v0);
              let _16_v11;
              _16_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("y"),(_13_cf8).fm1((((_15_v10).contains(p2)) ? ((_15_v10).get(p2)) : (_2_v0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(984)), function (_17_i1) {
                return (_dafny.Seq.UnicodeFromString("kqfg"));
              }), globalState));
              let _18_v12;
              _18_v12 = new _dafny.CodePoint('m'.codePointAt(0));
              _16_v11 = (_16_v11).update(_14_v9, !_dafny.Seq.contains(_14_v9, _18_v12));
              let _19_v13;
              let _nw2 = new _module.C0();
              _nw2.__ctor(((_4_v1).f22).update(_2_v0, p0));
              _19_v13 = _nw2;
            } else {
              let _20___mcc_h1 = (_source0).cf9;
              let _21_cf9 = _20___mcc_h1;
              let _22_v14;
              let _nw3 = Array((new BigNumber(9)).toNumber()).fill([]);
              _22_v14 = _nw3;
              let _index0 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_22_v14).length));
              (_22_v14)[_index0] = p0;
              let _index1 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_22_v14).length));
              (_22_v14)[_index1] = p0;
              let _23_v15;
              _23_v15 = _dafny.Seq.UnicodeFromString("u");
              let _24_v16;
              _24_v16 = _dafny.Seq.of(_4_v1, _4_v1, _4_v1, _4_v1, _4_v1);
              let _25_v17;
              _25_v17 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_23_v15).length));
              let _26_v18;
              _26_v18 = _module.D3.create_DC7();
              let _27_v19;
              _27_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2_v0,_26_v18);
              let _28_v20;
              _28_v20 = _dafny.Set.fromElements(_2_v0, _2_v0);
              let _29_v21;
              _29_v21 = new _dafny.CodePoint('t'.codePointAt(0));
              let _30_v22;
              let _nw4 = Array((new BigNumber(26)).toNumber());
              _nw4[(_dafny.ZERO).toNumber()] = new BigNumber(740);
              _nw4[(_dafny.ONE).toNumber()] = p2;
              _nw4[(new BigNumber(2)).toNumber()] = p2;
              _nw4[(new BigNumber(3)).toNumber()] = p2;
              _nw4[(new BigNumber(4)).toNumber()] = (new BigNumber(992)).minus(p2);
              _nw4[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(p2, p2);
              _nw4[(new BigNumber(6)).toNumber()] = p2;
              _nw4[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(p2, p2);
              _nw4[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-561));
              _nw4[(new BigNumber(9)).toNumber()] = p2;
              _nw4[(new BigNumber(10)).toNumber()] = new BigNumber((_24_v16).length);
              _nw4[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(p2, p2);
              _nw4[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length));
              _nw4[(new BigNumber(13)).toNumber()] = p2;
              _nw4[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_25_v17).length));
              _nw4[(new BigNumber(15)).toNumber()] = (p2).minus(p2);
              _nw4[(new BigNumber(16)).toNumber()] = new BigNumber(((_27_v19).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2_v0,_module.D3.create_DC7()))).length);
              _nw4[(new BigNumber(17)).toNumber()] = p2;
              _nw4[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(p2);
              _nw4[(new BigNumber(19)).toNumber()] = p2;
              _nw4[(new BigNumber(20)).toNumber()] = p2;
              _nw4[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((p2).multipliedBy(p2));
              _nw4[(new BigNumber(22)).toNumber()] = (p2).multipliedBy(new BigNumber((_28_v20).length));
              _nw4[(new BigNumber(23)).toNumber()] = p2;
              _nw4[(new BigNumber(24)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eybkhsaop"), _dafny.Seq.UnicodeFromString("ddyvcpun")), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eybkhsaop"), _dafny.Seq.UnicodeFromString("ddyvcpun"))).length)), _29_v21)).length);
              _nw4[(new BigNumber(25)).toNumber()] = new BigNumber(-209);
              _30_v22 = _nw4;
              let _index2 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_30_v22).length));
              (_30_v22)[_index2] = p2;
              let _31_v23;
              _31_v23 = _dafny.MultiSet.fromElements(_4_v1);
              let _32_v24;
              _32_v24 = _module.D3.create_DC6(_4_v1);
              let _33_v25;
              _33_v25 = _dafny.Map.Empty.slice().updateUnsafe(_32_v24,_2_v0);
              let _index3 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_30_v22).length));
              let _rhs0 = ((_31_v23).Intersect(_31_v23)).equals(_dafny.MultiSet.fromElements(_4_v1));
              let _rhs1 = _23_v15;
              let _rhs2 = ((_2_v0) ? (new BigNumber((_33_v25).length)) : ((p2).multipliedBy(p2)));
              let _rhs3 = new BigNumber(650);
              let _lhs0 = _30_v22;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_30_v22).length));
              let _lhs2 = globalState;
              r0 = _rhs0;
              _23_v15 = _rhs1;
              _lhs0[_lhs1] = _rhs2;
              _lhs2.f7 = _rhs3;
              let _34_v26;
              _34_v26 = _dafny.Seq.of(_23_v15);
              let _35_v27;
              _35_v27 = _dafny.MultiSet.fromElements(_2_v0, !(_2_v0), _2_v0, !(_2_v0), ((_2_v0) ? ((_4_v1).fm1(_2_v0, _34_v26, globalState)) : (_2_v0)));
              let _rhs4 = _35_v27;
              let _rhs5 = (_30_v22)[_module.__default.safeIndex(new BigNumber(783), new BigNumber((_30_v22).length))];
              let _lhs3 = globalState;
              _35_v27 = _rhs4;
              _lhs3.f20 = _rhs5;
              let _36_v28;
              _36_v28 = _dafny.Seq.of(false, _2_v0, _2_v0, _2_v0);
              let _index4 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((p0).length));
              (p0)[_index4] = !_dafny.Seq.contains(_36_v28, _2_v0);
              let _index5 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((p0).length));
              (p0)[_index5] = _2_v0;
            }
            let _37_v29;
            let _nw5 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
            _37_v29 = _nw5;
            let _38_v30;
            _38_v30 = _dafny.MultiSet.fromElements(_2_v0);
            let _39_v31;
            _39_v31 = _dafny.Seq.of(_4_v1);
            let _40_v32;
            _40_v32 = _dafny.Seq.UnicodeFromString("nb");
            let _index6 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_37_v29).length));
            (_37_v29)[_index6] = ((((_38_v30).contains(_2_v0)) ? ((_38_v30).get(_2_v0)) : (new BigNumber((_39_v31).length)))).multipliedBy(_module.__default.fm8(p2, new BigNumber((_40_v32).length), _2_v0, _2_v0, globalState));
            let _41_v33;
            _41_v33 = _dafny.Map.Empty.slice().updateUnsafe(_4_v1,_2_v0);
            let _42_v34;
            _42_v34 = new _dafny.CodePoint('p'.codePointAt(0));
            let _index7 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_37_v29).length));
            let _rhs6 = _2_v0;
            let _rhs7 = (!(!(_2_v0)) || (_2_v0)) === (_2_v0);
            let _rhs8 = ((new BigNumber((_module.__default.fm9(p2, p2, _2_v0, globalState)).length)).multipliedBy(p2)).plus(p2);
            let _rhs9 = ((!_dafny.areEqual(_40_v32, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("necjelxnj"), _module.__default.safeIndex(new BigNumber((_41_v33).length), new BigNumber((_dafny.Seq.UnicodeFromString("necjelxnj")).length)), _42_v34))) ? ((_dafny.ZERO).minus(p2)) : (p2));
            let _rhs10 = (p2).minus((new BigNumber((_dafny.Seq.UnicodeFromString("j")).length)).multipliedBy(p2));
            let _lhs4 = globalState;
            let _lhs5 = globalState;
            let _lhs6 = globalState;
            let _lhs7 = _37_v29;
            let _lhs8 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_37_v29).length));
            let _lhs9 = globalState;
            _lhs4.f6 = _rhs6;
            _lhs5.f4 = _rhs7;
            _lhs6.f14 = _rhs8;
            _lhs7[_lhs8] = _rhs9;
            _lhs9.f7 = _rhs10;
            _2_v0 = _2_v0;
            let _43_v35;
            _43_v35 = _dafny.Seq.of(_2_v0, _2_v0);
            let _44_v36;
            _44_v36 = _dafny.Set.fromElements(_2_v0, !((_43_v35)[_module.__default.safeIndex((_dafny.ZERO).minus((_37_v29)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_37_v29).length))]), new BigNumber((_43_v35).length))]));
            let _45_v37;
            _45_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2_v0,(_44_v36).equals(_44_v36));
            let _46_v38;
            _46_v38 = _dafny.Seq.of(_module.__default.fm8(new BigNumber((_40_v32).length), p2, _2_v0, _2_v0, globalState));
            let _47_v39;
            _47_v39 = _module.D1.create_DC2(p1);
            let _48_v40;
            _48_v40 = _module.D1.create_DC4(_47_v39);
            let _49_v41;
            _49_v41 = _dafny.Seq.of(_48_v40);
            let _50_v42;
            _50_v42 = _dafny.Set.fromElements(new BigNumber((_49_v41).length));
            let _51_v43;
            _51_v43 = _dafny.Map.Empty.slice().updateUnsafe(!(_2_v0),new _dafny.CodePoint('n'.codePointAt(0)));
            let _52_v44;
            _52_v44 = _dafny.MultiSet.fromElements(_51_v43, _51_v43);
            let _rhs11 = _4_v1;
            let _rhs12 = (_45_v37).Merge((_45_v37).update(_2_v0, _2_v0));
            let _rhs13 = (_37_v29)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_37_v29).length))];
            let _rhs14 = _module.__default.fm10(_50_v42, (_52_v44).Difference(_52_v44), globalState);
            let _rhs15 = _46_v38;
            let _lhs10 = globalState;
            _4_v1 = _rhs11;
            _45_v37 = _rhs12;
            _lhs10.f14 = _rhs13;
            _42_v34 = _rhs14;
            _46_v38 = _rhs15;
          }
        }
      }
      r1 = _2_v0;
      (globalState).f20 = p2;
      let _53_v45;
      _53_v45 = _dafny.Seq.UnicodeFromString("eqgwm");
      let _54_v46;
      _54_v46 = _dafny.Seq.of(_dafny.areEqual(_53_v45, _53_v45), (p2).isLessThan((_dafny.ZERO).minus(p2)), (_2_v0) || (_2_v0), _2_v0, _2_v0);
      let _55_v48;
      _55_v48 = _dafny.Seq.of(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(209), new BigNumber(925))) {
          let _56_v47 = _compr_0;
          if (((new BigNumber(209)).isLessThanOrEqualTo(_56_v47)) && ((_56_v47).isLessThan(new BigNumber(925)))) {
            _coll0.push([_module.__default.safeModuloInt(_56_v47, p2),_2_v0]);
          }
        }
        return _coll0;
      }()).length));
      r0 = (_54_v46)[_module.__default.safeIndex(_module.__default.fm8(new BigNumber((_55_v48).length), p2, _2_v0, false, globalState), new BigNumber((_54_v46).length))];
      let _57_v49;
      let _nw6 = Array((_dafny.ONE).toNumber()).fill(false);
      _57_v49 = _nw6;
      _57_v49 = p0;
      let _58_v50;
      _58_v50 = _dafny.Map.Empty.slice().updateUnsafe(_57_v49,(_55_v48)[_module.__default.safeIndex(p2, new BigNumber((_55_v48).length))]);
      _58_v50 = (_58_v50).Merge(_58_v50);
      let _59_v51;
      _59_v51 = _module.D0.create_DC1(p2, _2_v0);
      let _pat_let_tv0 = _2_v0;
      r0 = function (_source1) {
        if (_source1.is_DC1) {
          let _60___mcc_h2 = (_source1).cf1;
          let _61___mcc_h3 = (_source1).cf2;
          let _62_cf2 = _61___mcc_h3;
          let _63_cf1 = _60___mcc_h2;
          return _pat_let_tv0;
        } else {
          let _64___mcc_h4 = (_source1).cf0;
          let _65_cf0 = _64___mcc_h4;
          return false;
        }
      }(_59_v51);
      let _66_v52;
      _66_v52 = _dafny.MultiSet.fromElements(p2, p2);
      r1 = (_66_v52).IsDisjointFrom(_66_v52);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _67_v0;
      _67_v0 = new _dafny.CodePoint('w'.codePointAt(0));
      let _68_v1;
      _68_v1 = new BigNumber(557);
      let _69_v2;
      let _nw7 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _69_v2 = _nw7;
      let _70_v3;
      _70_v3 = _dafny.Map.Empty.slice().updateUnsafe(_69_v2,_68_v1);
      let _71_v4;
      _71_v4 = _dafny.Set.fromElements(new BigNumber((_70_v3).length), _68_v1);
      let _72_v5;
      _72_v5 = _dafny.MultiSet.fromElements(_71_v4);
      let _73_v6;
      _73_v6 = true;
      let _74_v7;
      _74_v7 = _dafny.MultiSet.fromElements(_73_v6);
      let _75_v8;
      _75_v8 = _dafny.Map.Empty.slice().updateUnsafe(_73_v6,_68_v1);
      let _76_v9;
      let _nw8 = Array((new BigNumber(22)).toNumber());
      _nw8[(_dafny.ZERO).toNumber()] = _68_v1;
      _nw8[(_dafny.ONE).toNumber()] = new BigNumber((_72_v5).cardinality());
      _nw8[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-906)), ((_77_v0) => function (_78_i1) {
        return _77_v0;
      })(_67_v0))).length);
      _nw8[(new BigNumber(3)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(4)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(5)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(6)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(7)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_68_v1,_73_v6)).length);
      _nw8[(new BigNumber(9)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(10)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(11)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(12)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(13)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(14)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(15)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(16)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(17)).toNumber()] = (((_74_v7).contains(_73_v6)) ? ((_74_v7).get(_73_v6)) : (_68_v1));
      _nw8[(new BigNumber(18)).toNumber()] = new BigNumber((_75_v8).length);
      _nw8[(new BigNumber(19)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(20)).toNumber()] = _68_v1;
      _nw8[(new BigNumber(21)).toNumber()] = new BigNumber(875);
      _76_v9 = _nw8;
      let _79_v10;
      _79_v10 = _dafny.Map.Empty.slice().updateUnsafe(_73_v6,!(_73_v6));
      let _80_v11;
      _80_v11 = _dafny.Set.fromElements(_75_v8, _75_v8, _75_v8);
      let _81_v12;
      let _nw9 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _81_v12 = _nw9;
      let _82_v13;
      _82_v13 = _dafny.Set.fromElements(_73_v6);
      let _83_v14;
      _83_v14 = _dafny.Seq.of(_82_v13);
      let _84_v15;
      _84_v15 = _dafny.Seq.UnicodeFromString("je");
      let _85_globalState;
      let _nw10 = new _module.GlobalState();
      _nw10.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-417)), function (_86_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }),_67_v0), _76_v9, _79_v10, _80_v11, true, _81_v12, true, new BigNumber(929), _74_v7, new BigNumber(319), new BigNumber(-434), true, (_71_v4).Intersect(_71_v4), _76_v9, new BigNumber(-685), true, ((_83_v14)[_module.__default.safeIndex(_68_v1, new BigNumber((_83_v14).length))]).Union(_dafny.Set.fromElements(_73_v6)), _69_v2, (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_84_v15).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_68_v1)), new BigNumber(581), new BigNumber(37), _84_v15);
      _85_globalState = _nw10;
      if (true) {
        let _index8 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_69_v2).length));
        (_69_v2)[_index8] = (_dafny.ZERO).minus(_68_v1);
        let _index9 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_69_v2).length));
        (_69_v2)[_index9] = _68_v1;
        _75_v8 = (_75_v8).update((_74_v7).IsSubsetOf(_74_v7), ((_69_v2)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_69_v2).length))]).multipliedBy((_69_v2)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_69_v2).length))]));
        let _87_v16;
        let _nw11 = Array((new BigNumber(19)).toNumber()).fill(false);
        _87_v16 = _nw11;
        _87_v16 = _87_v16;
        _84_v15 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), ((_88_v0) => function (_89_i2) {
          return _88_v0;
        })(_67_v0));
        _73_v6 = !(true) || (_73_v6);
      } else {
        let _90_v17;
        let _init0 = ((_91_v15) => function (_92_i3) {
          return _dafny.areEqual(_91_v15, _91_v15);
        })(_84_v15);
        let _nw12 = Array((new BigNumber(2)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw12.length); _i0_0++) {
          _nw12[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _90_v17 = _nw12;
        let _index10 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_90_v17).length));
        (_90_v17)[_index10] = _73_v6;
        let _93_v18;
        _93_v18 = _dafny.MultiSet.fromElements(new BigNumber(805), _68_v1);
        let _94_v19;
        _94_v19 = _dafny.Seq.of(_68_v1);
        let _95_v20;
        _95_v20 = _dafny.MultiSet.fromElements(((((_93_v18).contains(_68_v1)) ? ((_93_v18).get(_68_v1)) : (_68_v1))).minus(_68_v1), _68_v1, (new BigNumber((_94_v19).length)).plus(_68_v1), _68_v1);
        let _index11 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_69_v2).length));
        (_69_v2)[_index11] = _module.__default.safeDivisionInt(_68_v1, _68_v1);
        let _96_v21;
        _96_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_84_v15).length), _68_v1),_dafny.MultiSet.fromElements(_68_v1));
        let _index12 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_90_v17).length));
        let _index13 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_69_v2).length));
        let _rhs16 = _73_v6;
        let _rhs17 = (((_96_v21).contains(_94_v19)) ? ((_96_v21).get(_94_v19)) : (_95_v20));
        let _rhs18 = (new BigNumber(-971)).minus(_68_v1);
        let _rhs19 = _68_v1;
        let _lhs11 = _90_v17;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_90_v17).length));
        let _lhs13 = _69_v2;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_69_v2).length));
        let _lhs15 = _85_globalState;
        _lhs11[_lhs12] = _rhs16;
        _95_v20 = _rhs17;
        _lhs13[_lhs14] = _rhs18;
        _lhs15.f20 = _rhs19;
        (_85_globalState).f6 = _module.__default.fm0(_dafny.Map.Empty.slice().updateUnsafe(_68_v1,(_69_v2)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_69_v2).length))]), _68_v1, (_90_v17)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_90_v17).length))], _73_v6, _85_globalState);
        _76_v9 = _76_v9;
        let _97_v22;
        _97_v22 = _dafny.Map.Empty.slice().updateUnsafe((_69_v2)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_69_v2).length))],_68_v1);
        let _98_v23;
        _98_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_97_v22).length),(_90_v17)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_90_v17).length))]);
        let _99_v24;
        _99_v24 = _dafny.Map.Empty.slice().updateUnsafe(_98_v23,new BigNumber(446));
        _73_v6 = !(_99_v24).equals(_99_v24);
        _82_v13 = _82_v13;
      }
      let _100_v25;
      let _nw13 = Array((new BigNumber(16)).toNumber()).fill([]);
      _100_v25 = _nw13;
      let _101_v26;
      let _nw14 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _101_v26 = _nw14;
      let _index14 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_100_v25).length));
      (_100_v25)[_index14] = _101_v26;
      let _index15 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_100_v25).length));
      (_100_v25)[_index15] = _101_v26;
      (_85_globalState).f4 = _73_v6;
      let _102_v27;
      _102_v27 = _dafny.Map.Empty.slice().updateUnsafe(_73_v6,_dafny.Seq.UnicodeFromString("vfucpd"));
      _84_v15 = (((_102_v27).contains(_73_v6)) ? ((_102_v27).get(_73_v6)) : (_84_v15));
      (_85_globalState).f4 = (_68_v1).isLessThan(_68_v1);
      if (!(_73_v6) || (!(_73_v6))) {
        (_85_globalState).f14 = new BigNumber(-584);
        let _103_v28;
        _103_v28 = _dafny.Seq.of(_73_v6, _73_v6, _73_v6);
        let _104_v29;
        _104_v29 = _dafny.Map.Empty.slice().updateUnsafe(_68_v1,new BigNumber((_71_v4).length));
        let _105_v30;
        let _nw15 = Array((new BigNumber(11)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _73_v6;
        _nw15[(_dafny.ONE).toNumber()] = _73_v6;
        _nw15[(new BigNumber(2)).toNumber()] = _73_v6;
        _nw15[(new BigNumber(3)).toNumber()] = _73_v6;
        _nw15[(new BigNumber(4)).toNumber()] = (_103_v28)[_module.__default.safeIndex(_68_v1, new BigNumber((_103_v28).length))];
        _nw15[(new BigNumber(5)).toNumber()] = _73_v6;
        _nw15[(new BigNumber(6)).toNumber()] = _73_v6;
        _nw15[(new BigNumber(7)).toNumber()] = _73_v6;
        _nw15[(new BigNumber(8)).toNumber()] = _module.__default.fm0(_104_v29, _68_v1, _73_v6, _73_v6, _85_globalState);
        _nw15[(new BigNumber(9)).toNumber()] = _73_v6;
        _nw15[(new BigNumber(10)).toNumber()] = !(!(_73_v6));
        _105_v30 = _nw15;
        let _106_v31;
        _106_v31 = _dafny.Map.Empty.slice().updateUnsafe(_73_v6,_105_v30);
        let _107_v32;
        _107_v32 = _dafny.Map.Empty.slice().updateUnsafe(false,_74_v7);
        let _108_v33;
        _108_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.__default.fm0(_104_v29, _68_v1, _73_v6, _73_v6, _85_globalState), false),(((_107_v32).contains(true)) ? ((_107_v32).get(true)) : (_dafny.MultiSet.fromElements(_73_v6))));
        let _109_v34;
        let _110_v35;
        let _out0;
        let _out1;
        let _outcollector0 = _module.__default.m0(_105_v30, (_106_v31).update(true, _105_v30), new BigNumber((((_73_v6) ? (_108_v33) : (_108_v33))).length), _85_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _109_v34 = _out0;
        _110_v35 = _out1;
        let _111_v36;
        _111_v36 = _dafny.Map.Empty.slice().updateUnsafe(_76_v9,false);
        let _112_v37;
        _112_v37 = _dafny.Seq.of((_68_v1).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(236))).cardinality())), _68_v1, _module.__default.safeModuloInt(new BigNumber(((_111_v36).update(_76_v9, _109_v34)).length), _68_v1));
        let _113_v38;
        _113_v38 = _module.D0.create_DC0(_79_v10);
        let _114_v39;
        _114_v39 = _dafny.Map.Empty.slice().updateUnsafe(_68_v1,_110_v35);
        let _115_v40;
        _115_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_113_v38).dtor_cf0).length),_114_v39);
        _112_v37 = _dafny.Seq.update(_112_v37, _module.__default.safeIndex(_68_v1, new BigNumber((_112_v37).length)), new BigNumber(((_115_v40).Merge(_dafny.Map.Empty.slice().updateUnsafe(_68_v1,_114_v39))).length));
        let _116_v41;
        let _nw16 = new _module.C0();
        _nw16.__ctor(_106_v31);
        _116_v41 = _nw16;
        let _117_v42;
        let _nw17 = new _module.C0();
        _nw17.__ctor(_106_v31);
        _117_v42 = _nw17;
      } else {
        (_85_globalState).f20 = _module.__default.safeDivisionInt((_68_v1).minus((_dafny.ZERO).minus(_68_v1)), _68_v1);
        let _118_v43;
        let _nw18 = Array((new BigNumber(7)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = _73_v6;
        _nw18[(_dafny.ONE).toNumber()] = _73_v6;
        _nw18[(new BigNumber(2)).toNumber()] = true;
        _nw18[(new BigNumber(3)).toNumber()] = _73_v6;
        _nw18[(new BigNumber(4)).toNumber()] = _73_v6;
        _nw18[(new BigNumber(5)).toNumber()] = true;
        _nw18[(new BigNumber(6)).toNumber()] = _73_v6;
        _118_v43 = _nw18;
        let _119_v44;
        _119_v44 = _dafny.Map.Empty.slice().updateUnsafe(_73_v6,_118_v43);
        let _120_v45;
        let _nw19 = new _module.C0();
        _nw19.__ctor(_119_v44);
        _120_v45 = _nw19;
        _120_v45 = _120_v45;
        let _121_v46;
        _121_v46 = _dafny.Map.Empty.slice().updateUnsafe(_68_v1,_68_v1);
        let _122_v47;
        let _123_v48;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0(_118_v43, (_120_v45).f22, (((_121_v46).contains(_68_v1)) ? ((_121_v46).get(_68_v1)) : (_68_v1)), _85_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _122_v47 = _out2;
        _123_v48 = _out3;
        let _124_v49;
        _124_v49 = _dafny.Map.Empty.slice().updateUnsafe(_79_v10,new BigNumber((_84_v15).length));
        (_85_globalState).f14 = (((_124_v49).contains(_79_v10)) ? ((_124_v49).get(_79_v10)) : (_68_v1));
        let _125_v50;
        _125_v50 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_68_v1),_123_v48);
        _125_v50 = (_125_v50).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-604),_122_v47));
      }
      (_85_globalState).f14 = _68_v1;
      let _126_v51;
      let _nw20 = Array((new BigNumber(23)).toNumber());
      _nw20[(_dafny.ZERO).toNumber()] = _73_v6;
      _nw20[(_dafny.ONE).toNumber()] = _73_v6;
      _nw20[(new BigNumber(2)).toNumber()] = false;
      _nw20[(new BigNumber(3)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(4)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(5)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(6)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(7)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(8)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(9)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(10)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(11)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(12)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(13)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(14)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(15)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(16)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(17)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(18)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(19)).toNumber()] = !((((_79_v10).contains(true)) ? ((_79_v10).get(true)) : (_73_v6)));
      _nw20[(new BigNumber(20)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(21)).toNumber()] = _73_v6;
      _nw20[(new BigNumber(22)).toNumber()] = _73_v6;
      _126_v51 = _nw20;
      let _127_v52;
      _127_v52 = _dafny.Map.Empty.slice().updateUnsafe(false,_126_v51);
      let _128_v53;
      let _129_v54;
      let _out4;
      let _out5;
      let _outcollector2 = _module.__default.m0(_126_v51, _127_v52, (_68_v1).multipliedBy(_68_v1), _85_globalState);
      _out4 = _outcollector2[0];
      _out5 = _outcollector2[1];
      _128_v53 = _out4;
      _129_v54 = _out5;
      let _130_v56;
      _130_v56 = _dafny.Seq.of(_79_v10, _79_v10);
      let _131_v57;
      _131_v57 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_130_v56).Elements) {
          let _132_v55 = _compr_1;
          if (_dafny.Seq.contains(_130_v56, _132_v55)) {
            _coll1.push([_132_v55,!(_128_v53)]);
          }
        }
        return _coll1;
      }(),_128_v53);
      let _133_v58;
      _133_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dyot")).length),_68_v1);
      let _134_v59;
      _134_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_129_v54,_129_v54),_module.__default.fm0(_133_v58, _68_v1, _128_v53, false, _85_globalState));
      (_85_globalState).f4 = (((_131_v57).contains(_134_v59)) ? ((_131_v57).get(_134_v59)) : (_module.__default.fm0(_133_v58, _68_v1, _129_v54, _129_v54, _85_globalState)));
      (_85_globalState).f4 = _129_v54;
      let _135_v60;
      _135_v60 = _module.D1.create_DC2(_127_v52);
      let _136_v61;
      let _nw21 = new _module.C0();
      _nw21.__ctor((_135_v60).dtor_cf3);
      _136_v61 = _nw21;
      _136_v61 = _136_v61;
      let _hi0 = (_dafny.ZERO).minus(_68_v1);
      for (let _137_i4 = _68_v1; _137_i4.isLessThan(_hi0); _137_i4 = _137_i4.plus(_dafny.ONE)) {
        let _138_v62;
        let _nw22 = Array((new BigNumber(16)).toNumber()).fill([]);
        _138_v62 = _nw22;
        let _139_v63;
        _139_v63 = _dafny.Map.Empty.slice().updateUnsafe(_128_v53,_82_v13);
        let _140_v64;
        _140_v64 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(282)), ((_141_v0) => function (_142_i5) {
          return _141_v0;
        })(_67_v0)));
        let _143_v65;
        let _nw23 = Array((new BigNumber(14)).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_129_v54);
        _nw23[(_dafny.ONE).toNumber()] = _82_v13;
        _nw23[(new BigNumber(2)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(3)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(4)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(5)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(6)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(7)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(8)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(9)).toNumber()] = (((_139_v63).contains(_129_v54)) ? ((_139_v63).get(_129_v54)) : (_module.__default.fm2(_129_v54, (_136_v61).fm1(_129_v54, _dafny.Seq.update(_140_v64, _module.__default.safeIndex(_68_v1, new BigNumber((_140_v64).length)), _84_v15), _85_globalState), new BigNumber((_84_v15).length), _85_globalState)));
        _nw23[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_128_v53);
        _nw23[(new BigNumber(11)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(12)).toNumber()] = _82_v13;
        _nw23[(new BigNumber(13)).toNumber()] = _82_v13;
        _143_v65 = _nw23;
        let _index16 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_138_v62).length));
        (_138_v62)[_index16] = _143_v65;
        let _index17 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_138_v62).length));
        (_138_v62)[_index17] = ((_129_v54) ? (_143_v65) : (_143_v65));
        let _144_v66;
        _144_v66 = _69_v2;
        let _145_v67;
        _145_v67 = _dafny.MultiSet.fromElements((_144_v66), _76_v9, _76_v9, _69_v2, _76_v9);
        let _146_v68;
        _146_v68 = _dafny.Seq.of(!(true), _129_v54);
        let _147_v69;
        _147_v69 = _dafny.Map.Empty.slice().updateUnsafe(_145_v67,_146_v68);
        _147_v69 = (_147_v69).update(_dafny.MultiSet.fromElements(_76_v9, _69_v2), _dafny.Seq.Concat(_146_v68, _146_v68));
        (_85_globalState).f4 = ((_73_v6) ? ((_128_v53) || ((_136_v61).fm1(true, _140_v64, _85_globalState))) : (_128_v53));
        let _148_v70;
        _148_v70 = _module.D0.create_DC0(_79_v10);
        let _pat_let_tv1 = _79_v10;
        let _149_v71;
        _149_v71 = _dafny.Seq.of(_module.D0.create_DC0(_79_v10), function (_pat_let0_0) {
          return function (_150_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_151_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_151_dt__update_hcf0_h0);
              }(_pat_let1_0);
            }(_pat_let_tv1);
          }(_pat_let0_0);
        }(_148_v70));
        let _rhs20 = new BigNumber(495);
        let _rhs21 = _137_i4;
        let _rhs22 = _dafny.Seq.Concat(_dafny.Seq.of(_148_v70), _149_v71);
        let _rhs23 = ((_68_v1).plus(_137_i4)).multipliedBy((_68_v1).multipliedBy(_137_i4));
        let _rhs24 = _68_v1;
        let _lhs16 = _85_globalState;
        let _lhs17 = _85_globalState;
        let _lhs18 = _85_globalState;
        _68_v1 = _rhs20;
        _lhs16.f14 = _rhs21;
        _149_v71 = _rhs22;
        _lhs17.f7 = _rhs23;
        _lhs18.f20 = _rhs24;
      }
      if (_module.__default.fm0(_133_v58, _68_v1, _129_v54, _73_v6, _85_globalState)) {
        let _index18 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_81_v12).length));
        (_81_v12)[_index18] = _84_v15;
        let _152_v72;
        _152_v72 = _dafny.MultiSet.fromElements(new BigNumber(746), _68_v1);
        let _153_v73;
        _153_v73 = _dafny.Seq.of(new BigNumber((_152_v72).cardinality()));
        let _154_v74;
        _154_v74 = _dafny.Seq.of(_153_v73, _153_v73, _153_v73);
        let _index19 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_81_v12).length));
        let _rhs25 = _dafny.Seq.IsProperPrefixOf((_154_v74)[_module.__default.safeIndex(_68_v1, new BigNumber((_154_v74).length))], _module.__default.fm3(_85_globalState));
        let _rhs26 = (_68_v1).minus(_68_v1);
        let _rhs27 = !(((((_152_v72).contains(_68_v1)) ? ((_152_v72).get(_68_v1)) : (_68_v1))).isEqualTo(_68_v1));
        let _rhs28 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))), _84_v15), _module.__default.safeIndex(_68_v1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))), _84_v15)).length)), _67_v0), _module.__default.safeIndex(_module.__default.safeDivisionInt(_68_v1, _68_v1), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))), _84_v15), _module.__default.safeIndex(_68_v1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))), _84_v15)).length)), _67_v0)).length)), ((_129_v54) ? (_67_v0) : (_67_v0)));
        let _lhs19 = _85_globalState;
        let _lhs20 = _81_v12;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_81_v12).length));
        _129_v54 = _rhs25;
        _lhs19.f7 = _rhs26;
        _129_v54 = _rhs27;
        _lhs20[_lhs21] = _rhs28;
        let _155_v75;
        _155_v75 = _dafny.Seq.of(_127_v52, _127_v52);
        let _156_v76;
        let _nw24 = new _module.C0();
        _nw24.__ctor((_155_v75)[_module.__default.safeIndex(_68_v1, new BigNumber((_155_v75).length))]);
        _156_v76 = _nw24;
        let _157_v77;
        let _nw25 = new _module.C0();
        _nw25.__ctor((_136_v61).f22);
        _157_v77 = _nw25;
        _133_v58 = _module.__default.fm4(((_128_v53) ? (_68_v1) : (_68_v1)), _68_v1, _85_globalState);
        let _158_v78;
        _158_v78 = _dafny.Set.fromElements(_67_v0, _67_v0);
        let _159_v79;
        _159_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_79_v10),new BigNumber((_158_v78).length));
        let _160_v80;
        _160_v80 = _dafny.Map.Empty.slice().updateUnsafe(_129_v54,_152_v72);
        let _161_v81;
        let _162_v82;
        let _out6;
        let _out7;
        let _outcollector3 = _module.__default.m0(_126_v51, (_156_v76).f22, ((((_159_v79).contains(_module.__default.fm5(_152_v72, _85_globalState))) ? ((_159_v79).get(_module.__default.fm5(_152_v72, _85_globalState))) : (new BigNumber(315)))).minus(new BigNumber(((((_160_v80).contains(_73_v6)) ? ((_160_v80).get(_73_v6)) : (_152_v72))).cardinality())), _85_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _161_v81 = _out6;
        _162_v82 = _out7;
      } else {
        if (((_136_v61).fm1(_128_v53, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-683)), function (_163_i6) {
          return _dafny.Seq.UnicodeFromString("u");
        }), _85_globalState)) && ((_68_v1).isLessThanOrEqualTo(_68_v1))) {
          let _164_v83;
          _164_v83 = _dafny.Map.Empty.slice().updateUnsafe(_68_v1,_84_v15);
          _84_v15 = (((_164_v83).contains(_68_v1)) ? ((_164_v83).get(_68_v1)) : (_84_v15));
          let _165_v84;
          let _nw26 = new _module.C0();
          _nw26.__ctor((_136_v61).f22);
          _165_v84 = _nw26;
          let _index20 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_76_v9).length));
          (_76_v9)[_index20] = _68_v1;
          let _166_v85;
          _166_v85 = _module.D0.create_DC1(_68_v1, _129_v54);
          let _167_v86;
          _167_v86 = _dafny.MultiSet.fromElements(_166_v85);
          let _168_v87;
          _168_v87 = _dafny.Map.Empty.slice().updateUnsafe(_167_v86,_129_v54);
          let _169_v88;
          let _nw27 = Array((new BigNumber(2)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _168_v87;
          _nw27[(_dafny.ONE).toNumber()] = _168_v87;
          _169_v88 = _nw27;
          let _index21 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_169_v88).length));
          (_169_v88)[_index21] = (_168_v87).Merge(_168_v87);
          let _170_v89;
          _170_v89 = _dafny.Seq.of(_73_v6, _129_v54, !(!(true)));
          let _171_v90;
          _171_v90 = _69_v2;
          let _172_v91;
          _172_v91 = (_171_v90);
          let _173_v92;
          _173_v92 = _dafny.Seq.of(new BigNumber(-376), new BigNumber(212));
          let _174_v93;
          _174_v93 = _dafny.Map.Empty.slice().updateUnsafe(_172_v91,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_173_v92, _module.__default.safeIndex(_68_v1, new BigNumber((_173_v92).length)), _68_v1)).length)));
          let _175_v94;
          _175_v94 = _module.D0.create_DC0(_79_v10);
          let _176_v95;
          _176_v95 = _dafny.Map.Empty.slice().updateUnsafe(_175_v94,_136_v61);
          let _177_v96;
          _177_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_174_v93).Merge(_174_v93)).length),(((_176_v95).contains(_175_v94)) ? ((_176_v95).get(_175_v94)) : (_165_v84)));
          let _178_v97;
          _178_v97 = _module.D3.create_DC6(_165_v84);
          let _179_v98;
          _179_v98 = _dafny.MultiSet.fromElements(_82_v13);
          let _index22 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_76_v9).length));
          let _index23 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_169_v88).length));
          let _rhs29 = _68_v1;
          let _rhs30 = (((_177_v96).contains(_68_v1)) ? ((_177_v96).get(_68_v1)) : ((_178_v97).dtor_cf8));
          let _rhs31 = new BigNumber((_179_v98).cardinality());
          let _rhs32 = ((_168_v87).Merge(_168_v87)).Merge(_168_v87);
          let _rhs33 = _170_v89;
          let _lhs22 = _76_v9;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_76_v9).length));
          let _lhs24 = _85_globalState;
          let _lhs25 = _169_v88;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_169_v88).length));
          _lhs22[_lhs23] = _rhs29;
          _136_v61 = _rhs30;
          _lhs24.f7 = _rhs31;
          _lhs25[_lhs26] = _rhs32;
          _170_v89 = _rhs33;
          (_85_globalState).f4 = (!(_129_v54)) || (_129_v54);
          _84_v15 = _84_v15;
        } else {
          let _180_v99;
          let _nw28 = new _module.C0();
          _nw28.__ctor((_136_v61).f22);
          _180_v99 = _nw28;
          _84_v15 = _84_v15;
          _129_v54 = _129_v54;
          let _181_v100;
          let _nw29 = new _module.C0();
          _nw29.__ctor((_136_v61).f22);
          _181_v100 = _nw29;
          let _182_v101;
          _182_v101 = _dafny.Seq.of(_181_v100);
          let _183_v102;
          let _nw30 = Array((new BigNumber(21)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _136_v61;
          _nw30[(_dafny.ONE).toNumber()] = (_182_v101)[_module.__default.safeIndex(_68_v1, new BigNumber((_182_v101).length))];
          _nw30[(new BigNumber(2)).toNumber()] = _136_v61;
          _nw30[(new BigNumber(3)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(4)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(5)).toNumber()] = _180_v99;
          _nw30[(new BigNumber(6)).toNumber()] = _136_v61;
          _nw30[(new BigNumber(7)).toNumber()] = _136_v61;
          _nw30[(new BigNumber(8)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(9)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(10)).toNumber()] = _136_v61;
          _nw30[(new BigNumber(11)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(12)).toNumber()] = _180_v99;
          _nw30[(new BigNumber(13)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(14)).toNumber()] = _180_v99;
          _nw30[(new BigNumber(15)).toNumber()] = _136_v61;
          _nw30[(new BigNumber(16)).toNumber()] = _136_v61;
          _nw30[(new BigNumber(17)).toNumber()] = _180_v99;
          _nw30[(new BigNumber(18)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(19)).toNumber()] = _181_v100;
          _nw30[(new BigNumber(20)).toNumber()] = _181_v100;
          _183_v102 = _nw30;
          let _184_v103;
          _184_v103 = _dafny.Set.fromElements(_183_v102);
          let _rhs34 = !(_128_v53);
          let _rhs35 = _73_v6;
          let _rhs36 = new BigNumber((((_184_v103).Union(_184_v103)).Union((_184_v103).Difference(_dafny.Set.fromElements(_183_v102)))).length);
          let _lhs27 = _85_globalState;
          let _lhs28 = _85_globalState;
          _73_v6 = _rhs34;
          _lhs27.f6 = _rhs35;
          _lhs28.f14 = _rhs36;
        }
        let _index24 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length));
        (_76_v9)[_index24] = _68_v1;
        let _185_v104;
        _185_v104 = _dafny.Map.Empty.slice().updateUnsafe(_129_v54,_76_v9);
        let _index25 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length));
        (_76_v9)[_index25] = new BigNumber((_185_v104).length);
        let _186_v105;
        _186_v105 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_84_v15).length),_129_v54);
        let _187_v106;
        _187_v106 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-955),!((((_186_v105).contains(_68_v1)) ? ((_186_v105).get(_68_v1)) : (_73_v6))));
        _187_v106 = _187_v106;
        (_85_globalState).f6 = true;
        let _188_v107;
        _188_v107 = _dafny.Seq.of(_129_v54, _128_v53, _129_v54);
        let _189_v108;
        _189_v108 = _dafny.Seq.of((_76_v9)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length))]);
        let _190_v109;
        _190_v109 = _dafny.Map.Empty.slice().updateUnsafe(_188_v107,new BigNumber((_module.__default.fm6(_189_v108, _128_v53, _73_v6, _85_globalState)).length));
        let _index26 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length));
        let _rhs37 = ((_129_v54) ? (((_76_v9)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length))]).minus((_76_v9)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length))])) : ((_76_v9)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length))]));
        let _rhs38 = !(_73_v6);
        let _rhs39 = (_dafny.ZERO).minus((((_190_v109).contains(_dafny.Seq.Concat(_188_v107, _188_v107))) ? ((_190_v109).get(_dafny.Seq.Concat(_188_v107, _188_v107))) : ((_76_v9)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length))])));
        let _rhs40 = (_76_v9)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length))];
        let _lhs29 = _76_v9;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_76_v9).length));
        let _lhs31 = _85_globalState;
        let _lhs32 = _85_globalState;
        let _lhs33 = _85_globalState;
        _lhs29[_lhs30] = _rhs37;
        _lhs31.f4 = _rhs38;
        _lhs32.f14 = _rhs39;
        _lhs33.f14 = _rhs40;
      }
      (_85_globalState).f14 = (_68_v1).multipliedBy(_68_v1);
      let _191_i7;
      _191_i7 = _dafny.ZERO;
      L1: {
        while (_dafny.Seq.IsProperPrefixOf(_84_v15, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pt"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(834)), ((_207_v0) => function (_208_i8) {
          return _207_v0;
        })(_67_v0))))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_191_i7)) {
              break L1;
            }
            _191_i7 = (_191_i7).plus(_dafny.ONE);
            let _192_v110;
            _192_v110 = _module.D3.create_DC6(_136_v61);
            _192_v110 = _module.D3.create_DC6((_192_v110).dtor_cf8);
            let _193_v111;
            let _nw31 = new _module.C0();
            _nw31.__ctor(_127_v52);
            _193_v111 = _nw31;
            let _194_v112;
            _194_v112 = (((_102_v27).contains(_73_v6)) ? ((_102_v27).get(_73_v6)) : (_84_v15));
            _84_v15 = _dafny.Seq.Concat(_dafny.Seq.Concat((_194_v112), _84_v15), _84_v15);
            let _195_v113;
            _195_v113 = _dafny.Seq.of(_84_v15, _dafny.Seq.Create(_module.__default.abs(new BigNumber(285)), function (_196_i10) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            }), _dafny.Seq.UnicodeFromString("ikudn"));
            if (_module.__default.fm0(_133_v58, ((_dafny.ZERO).minus((_dafny.ZERO).minus(_68_v1))).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(6)), function (_197_i9) {
              return new BigNumber(90);
            })).length)), _128_v53, (_193_v111).fm1(_73_v6, _195_v113, _85_globalState), _85_globalState)) {
              (_85_globalState).f14 = (_68_v1).minus(_module.__default.safeDivisionInt(new BigNumber(915), _68_v1));
              _84_v15 = _dafny.Seq.UnicodeFromString("sv");
              _84_v15 = _dafny.Seq.UnicodeFromString("jbjlpyup");
              let _198_v114;
              let _nw32 = new _module.C0();
              _nw32.__ctor(((_136_v61).f22).Merge(((_136_v61).f22).update(_73_v6, _126_v51)));
              _198_v114 = _nw32;
              let _199_v115;
              _199_v115 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_193_v111,_193_v111)).length), _68_v1);
              let _200_v117;
              _200_v117 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_84_v15, _module.__default.fm7(function () {
                let _coll2 = new _dafny.Set();
                for (const _compr_2 of (_199_v115).Elements) {
                  let _201_v116 = _compr_2;
                  if (_dafny.Seq.contains(_199_v115, _201_v116)) {
                    _coll2.add((_201_v116).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), function (_202_i11) {
                      return (_dafny.ZERO).minus(new BigNumber(-117));
                    })).length)));
                  }
                }
                return _coll2;
              }(), _85_globalState)),_dafny.Map.Empty.slice().updateUnsafe(_68_v1,_68_v1));
              let _203_v118;
              _203_v118 = _module.D3.create_DC8(_module.D3.create_DC7());
              let _rhs41 = _73_v6;
              let _rhs42 = _200_v117;
              let _rhs43 = _76_v9;
              let _rhs44 = _203_v118;
              let _lhs34 = _85_globalState;
              _lhs34.f4 = _rhs41;
              _200_v117 = _rhs42;
              _76_v9 = _rhs43;
              _203_v118 = _rhs44;
            } else {
              _84_v15 = _84_v15;
              let _204_v119;
              let _nw33 = Array((new BigNumber(3)).toNumber()).fill(_module.D3.Default());
              _204_v119 = _nw33;
              let _rhs45 = (_module.D3.create_DC6(_193_v111)).dtor_cf8;
              let _rhs46 = _204_v119;
              let _rhs47 = _84_v15;
              _136_v61 = _rhs45;
              _204_v119 = _rhs46;
              _84_v15 = _rhs47;
              let _205_v120;
              let _nw34 = new _module.C0();
              _nw34.__ctor(_127_v52);
              _205_v120 = _nw34;
              let _206_v121;
              let _nw35 = new _module.C0();
              _nw35.__ctor((_136_v61).f22);
              _206_v121 = _nw35;
              _75_v8 = (_75_v8).update(_129_v54, _68_v1);
            }
          }
        }
      }
      let _hi1 = new BigNumber(492);
      for (let _209_i12 = ((_73_v6) ? (new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(372), new BigNumber(739))) {
          let _225_v122 = _compr_3;
          if (((new BigNumber(372)).isLessThanOrEqualTo(_225_v122)) && ((_225_v122).isLessThan(new BigNumber(739)))) {
            _coll3.push([_module.__default.safeDivisionInt(_225_v122, _68_v1),_128_v53]);
          }
        }
        return _coll3;
      }()).length)) : (_68_v1)); _209_i12.isLessThan(_hi1); _209_i12 = _209_i12.plus(_dafny.ONE)) {
        let _210_v123;
        _210_v123 = _module.D3.create_DC6(_136_v61);
        let _211_v124;
        _211_v124 = _module.D3.create_DC8(_210_v123);
        let _pat_let_tv2 = _210_v123;
        let _pat_let_tv3 = _210_v123;
        let _212_v125;
        let _nw36 = Array((new BigNumber(6)).toNumber());
        _nw36[(_dafny.ZERO).toNumber()] = _211_v124;
        _nw36[(_dafny.ONE).toNumber()] = _211_v124;
        _nw36[(new BigNumber(2)).toNumber()] = _211_v124;
        _nw36[(new BigNumber(3)).toNumber()] = function (_pat_let2_0) {
          return function (_213_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_214_dt__update_hcf9_h0) {
                return _module.D3.create_DC8(_214_dt__update_hcf9_h0);
              }(_pat_let3_0);
            }(_pat_let_tv2);
          }(_pat_let2_0);
        }(_211_v124);
        _nw36[(new BigNumber(4)).toNumber()] = function (_pat_let4_0) {
          return function (_215_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_216_dt__update_hcf9_h1) {
                return _module.D3.create_DC8(_216_dt__update_hcf9_h1);
              }(_pat_let5_0);
            }(_pat_let_tv3);
          }(_pat_let4_0);
        }(_211_v124);
        _nw36[(new BigNumber(5)).toNumber()] = _211_v124;
        _212_v125 = _nw36;
        let _217_v126;
        _217_v126 = _dafny.Map.Empty.slice().updateUnsafe(_212_v125,_76_v9);
        let _index27 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        (_126_v51)[_index27] = _129_v54;
        let _218_v127;
        _218_v127 = _dafny.Map.Empty.slice().updateUnsafe(_67_v0,_136_v61);
        let _index28 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        let _rhs48 = ((_217_v126).update(_212_v125, _69_v2)).Merge((_217_v126).Merge(_217_v126));
        let _rhs49 = _126_v51;
        let _rhs50 = (((_218_v127).contains(new _dafny.CodePoint('t'.codePointAt(0)))) ? ((_218_v127).get(new _dafny.CodePoint('t'.codePointAt(0)))) : (_136_v61));
        let _rhs51 = (_129_v54) || (_73_v6);
        let _lhs35 = _126_v51;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        _217_v126 = _rhs48;
        _126_v51 = _rhs49;
        _136_v61 = _rhs50;
        _lhs35[_lhs36] = _rhs51;
        let _219_v128;
        let _nw37 = new _module.C0();
        _nw37.__ctor((_136_v61).f22);
        _219_v128 = _nw37;
        let _index29 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_69_v2).length));
        (_69_v2)[_index29] = _68_v1;
        let _220_v129;
        _220_v129 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(198)), ((_221_v0) => function (_222_i13) {
          return _221_v0;
        })(_67_v0))).length), _209_i12);
        let _223_v130;
        _223_v130 = _dafny.Map.Empty.slice().updateUnsafe((_220_v129).Difference(_220_v129),_dafny.Seq.of(new BigNumber(653)));
        let _index30 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_69_v2).length));
        let _index31 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        let _rhs52 = (new BigNumber((_dafny.Seq.of(_68_v1, _209_i12, _68_v1, _209_i12, new BigNumber(658))).length)).minus(_module.__default.safeModuloInt(_68_v1, _68_v1));
        let _rhs53 = _223_v130;
        let _rhs54 = ((!(_73_v6)) ? (_136_v61) : (_219_v128));
        let _rhs55 = !((_126_v51)[_module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length))]);
        let _lhs37 = _69_v2;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_69_v2).length));
        let _lhs39 = _126_v51;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        _lhs37[_lhs38] = _rhs52;
        _223_v130 = _rhs53;
        _219_v128 = _rhs54;
        _lhs39[_lhs40] = _rhs55;
        let _224_v131;
        _224_v131 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("uiiomhn"));
        let _index32 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        let _rhs56 = (((_126_v51)[_module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length))]) ? (_module.__default.fm0(_133_v58, _68_v1, (_219_v128).fm1(!(_128_v53), _dafny.Seq.of(_module.__default.fm7(_71_v4, _85_globalState), _dafny.Seq.UnicodeFromString("odeavklj")), _85_globalState), (_219_v128).fm1(true, _224_v131, _85_globalState), _85_globalState)) : (_128_v53));
        let _rhs57 = (_69_v2)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_69_v2).length))];
        let _lhs41 = _126_v51;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_126_v51).length));
        let _lhs43 = _85_globalState;
        _lhs41[_lhs42] = _rhs56;
        _lhs43.f20 = _rhs57;
      }
      process.stdout.write(_dafny.toString(_67_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_68_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_70_v3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v4).equals(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(557)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_72_v5).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(557))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_73_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v7).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(310249)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v9)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_79_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v11).equals(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(557))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_81_v12)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_82_v13).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_83_v14, _dafny.Seq.of(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_84_v15).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f0).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))),new _dafny.CodePoint('w'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f1)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f2).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f3).equals(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(557))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_85_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_85_globalState).f5)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_85_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_85_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f8).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f12).equals(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(557)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f13)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_85_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f16).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f17)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f17)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_globalState).f18).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(2)).updateUnsafe(true,new BigNumber(557)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_85_globalState.f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_85_globalState).f21).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("vfucpd")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v51)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_127_v52).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_128_v53));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_129_v54));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_130_v56, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v57).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,false),true),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v58).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(387),new BigNumber(-10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v59).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,false),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_135_v60).dtor_cf3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_136_v61).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_191_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, false);
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
    static create_DC3(cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf3) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC4(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4 && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false, false);
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
    static create_DC5(cf7) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC7() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D3(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC8(cf9) {
      let $dt = new D3(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf9) + ")";
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
        return other.$tag === 1 && this.cf8 === other.cf8;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7();
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
    static create_DC9(cf10) {
      let $dt = new D4(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf11) {
      let $dt = new D5(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf12) {
      let $dt = new D6(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f4 = false;
      this.f6 = false;
      this.f7 = _dafny.ZERO;
      this.f14 = _dafny.ZERO;
      this.f20 = _dafny.ZERO;
      this._f0 = _dafny.Map.Empty;
      this._f1 = [];
      this._f2 = _dafny.Map.Empty;
      this._f3 = _dafny.Set.Empty;
      this._f5 = [];
      this._f8 = _dafny.MultiSet.Empty;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = _dafny.Set.Empty;
      this._f13 = [];
      this._f15 = false;
      this._f16 = _dafny.Set.Empty;
      this._f17 = [];
      this._f18 = _dafny.Map.Empty;
      this._f19 = _dafny.ZERO;
      this._f21 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this).f20 = f20;
      (_this)._f21 = f21;
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
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f22 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22) {
      let _this = this;
      (_this)._f22 = f22;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      let _source2 = _module.D0.create_DC1(new BigNumber((_dafny.Set.fromElements(!(false), true, true)).length), !(false));
      if (_source2.is_DC1) {
        let _226___mcc_h0 = (_source2).cf1;
        let _227___mcc_h1 = (_source2).cf2;
        let _228_cf2 = _227___mcc_h1;
        let _229_cf1 = _226___mcc_h0;
        return (_module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mpsgaxsnn")).length)), _228_cf2)).dtor_cf2;
      } else {
        let _230___mcc_h2 = (_source2).cf0;
        let _231_cf0 = _230___mcc_h2;
        return !(!(false)) || (true);
      }
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
