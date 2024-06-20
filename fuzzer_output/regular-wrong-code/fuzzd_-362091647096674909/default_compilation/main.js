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
      return _module.__default.safeModuloInt(new BigNumber(-739), ((false) ? (new BigNumber(121)) : (new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))));
    };
    static fm1(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cww"), _dafny.Seq.UnicodeFromString("g")), ((true) ? (_dafny.Seq.UnicodeFromString("iekbdyjkt")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-653)), function (_0_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }))));
    };
    static fm2(globalState) {
      return !(((new BigNumber(-923)).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(true)).length))) || (true));
    };
    static fm3(p0, globalState) {
      return _dafny.MultiSet.fromElements(!(_dafny.Map.Empty.slice().updateUnsafe(false,true)).contains(false), !(false), (new BigNumber((_dafny.Seq.UnicodeFromString("fvsmrmtk")).length)).isEqualTo(new BigNumber(-139)));
    };
    static fm12(p0, p1, p2, globalState) {
      return _module.D0.create_DC1((_dafny.ZERO).minus((new BigNumber(391)).plus(new BigNumber(820))));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D0.create_DC1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(553),new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rejvchf"))).length))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),_module.D0.create_DC1(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),_module.D0.create_DC1(new BigNumber(-482))));
    };
    static fm14(p0, globalState) {
      return _module.D0.create_DC1(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kd")).length))).minus(new BigNumber(74)));
    };
    static fm18(globalState) {
      return _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-377), new BigNumber(741))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(-377)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(741)))) {
            _coll0.push([(_1_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("levb")).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, false, false)).length))]);
          }
        }
        return _coll0;
      }()).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(393), new BigNumber(-101))) {
          let _2_v1 = _compr_1;
          if (((new BigNumber(393)).isLessThanOrEqualTo(_2_v1)) && ((_2_v1).isLessThan(new BigNumber(-101)))) {
            _coll1.push([(_2_v1).plus(new BigNumber((_dafny.Set.fromElements(true, true)).length)),new BigNumber(-736)]);
          }
        }
        return _coll1;
      }())).length)), (new BigNumber(-318)).multipliedBy(new BigNumber(841)));
    };
    static fm19(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(!(true))));
    };
    static fm20(p0, p1, globalState) {
      return _module.D0.create_DC0(_dafny.MultiSet.fromElements(false));
    };
    static fm21(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), function (_3_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("vtjpii")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("nnq")));
    };
    static fm27(p0, p1, globalState) {
      return (function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Set.fromElements(_dafny.Seq.of(_module.D5.create_DC13(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("epbtsmm")).length), new BigNumber((_dafny.Seq.UnicodeFromString("upj")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), function (_4_i0) {
  return new BigNumber(-734);
})).length))).cardinality()), false, false)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber((_dafny.Seq.UnicodeFromString("ju")).length), false, true)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber(-703), true, false), _module.D5.create_DC13(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))).cardinality()), true, false)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(194)), function (_5_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length), true, true)))).Elements) {
          let _6_v0 = _compr_2;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(_module.D5.create_DC13(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("epbtsmm")).length), new BigNumber((_dafny.Seq.UnicodeFromString("upj")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), function (_4_i0) {
  return new BigNumber(-734);
})).length))).cardinality()), false, false)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber((_dafny.Seq.UnicodeFromString("ju")).length), false, true)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber(-703), true, false), _module.D5.create_DC13(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))).cardinality()), true, false)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(194)), function (_5_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length), true, true)))).contains(_6_v0)) {
            _coll2.add(_6_v0);
          }
        }
        return _coll2;
      }()).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(_module.D5.create_DC13((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('n'.codePointAt(0)))).length)), true, false)), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber(-505), true, !(false))), _dafny.Seq.of(_module.D5.create_DC13(new BigNumber(953), true, false))));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC2(!((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("teghdwfgg")).length))).length)))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dhgmuljeq")).length)))), _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(false, true)).length))).length)), ((true) ? (true) : (!(true))));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D0.create_DC0(_dafny.MultiSet.fromElements(false, false, false));
      if (_source0.is_DC1) {
        let _7___mcc_h0 = (_source0).cf1;
        let _8_cf1 = _7___mcc_h0;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else if (_source0.is_DC2) {
        let _9___mcc_h1 = (_source0).cf2;
        let _10___mcc_h2 = (_source0).cf3;
        let _11___mcc_h3 = (_source0).cf4;
        let _12_cf4 = _11___mcc_h3;
        let _13_cf3 = _10___mcc_h2;
        let _14_cf2 = _9___mcc_h1;
        return new _dafny.CodePoint('h'.codePointAt(0));
      } else if (_source0.is_DC0) {
        let _15___mcc_h4 = (_source0).cf0;
        let _16_cf0 = _15___mcc_h4;
        return new _dafny.CodePoint('k'.codePointAt(0));
      } else {
        let _17___mcc_h5 = (_source0).cf5;
        let _18_cf5 = _17___mcc_h5;
        return new _dafny.CodePoint('w'.codePointAt(0));
      }
    };
    static fm30(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), function (_19_i0) {
        return new BigNumber(-681);
      })).length)), _dafny.MultiSet.fromElements(new BigNumber(-454)), _dafny.MultiSet.fromElements(new BigNumber(-972)), _dafny.MultiSet.fromElements(new BigNumber(756), new BigNumber(399), new BigNumber((function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_20_i1) {
          return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("umygtlqs")).length), new BigNumber(503))).length);
        })).length))).length),false))).Elements) {
          let _21_v0 = _compr_3;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_20_i1) {
            return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("umygtlqs")).length), new BigNumber(503))).length);
          })).length))).length),false))).contains(_21_v0)) {
            _coll3.add(_21_v0);
          }
        }
        return _coll3;
      }()).length)), _dafny.MultiSet.fromElements(new BigNumber(817), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true)).length)))).Intersect((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber(529), new BigNumber(926)))).Union(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Seq.of(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-43))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, false)).length)))).Elements) {
          let _22_v1 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-43))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, false)).length))), _22_v1)) {
            _coll4.add(_22_v1);
          }
        }
        return _coll4;
      }()));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(554), new BigNumber(634))) {
          let _23_v0 = _compr_5;
          if (((new BigNumber(554)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(634)))) {
            _coll5.push([(_23_v0).plus(new BigNumber(707)),!(false)]);
          }
        }
        return _coll5;
      }()).length)))).length), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true), true)).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), function (_24_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).length));
      })))).length))), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("iytbot")).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(593), new BigNumber(-150))) {
          let _25_v1 = _compr_6;
          if (((new BigNumber(593)).isLessThanOrEqualTo(_25_v1)) && ((_25_v1).isLessThan(new BigNumber(-150)))) {
            _coll6.push([(_25_v1).multipliedBy(new BigNumber(680)),new BigNumber(551)]);
          }
        }
        return _coll6;
      }()).length))), _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-732)), function (_26_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),!(false))).length);
      })));
    };
    static fm34(p0, p1, globalState) {
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    static fm35(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(745))).Difference((_dafny.Set.fromElements(new BigNumber(69))).Union(_dafny.Set.fromElements(new BigNumber(184), new BigNumber(786), new BigNumber(16))));
    };
    static fm36(globalState) {
      return _dafny.Seq.of((false) === (true), !((new BigNumber((_dafny.Seq.UnicodeFromString("wqt")).length)).isLessThanOrEqualTo(new BigNumber(695))), true, true);
    };
    static fm37(p0, p1, globalState) {
      let _source1 = _module.D21.create_DC50(_dafny.MultiSet.fromElements(_module.D9.create_DC20(_dafny.Seq.UnicodeFromString("ere"), true, new BigNumber(-948)), _module.D9.create_DC20(_dafny.Seq.UnicodeFromString("nmnxpj"), true, new BigNumber(189))));
      if (_source1.is_DC51) {
        let _27___mcc_h0 = (_source1).cf81;
        let _28_cf81 = _27___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("rlotprgxp")).length),_dafny.Seq.of(true))).Merge(function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-654), new BigNumber(-939))) {
            let _29_v0 = _compr_7;
            if (((new BigNumber(-654)).isLessThanOrEqualTo(_29_v0)) && ((_29_v0).isLessThan(new BigNumber(-939)))) {
              _coll7.push([(_29_v0).multipliedBy(new BigNumber(-920)),_dafny.Seq.of(_28_cf81)]);
            }
          }
          return _coll7;
        }());
      } else if (_source1.is_DC50) {
        let _30___mcc_h1 = (_source1).cf80;
        let _31_cf80 = _30___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-855),_dafny.Seq.of(!(false)));
      } else {
        let _32___mcc_h2 = (_source1).cf82;
        let _33_cf82 = _32___mcc_h2;
        return function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of _dafny.IntegerRange(new BigNumber(201), new BigNumber(952))) {
            let _34_v1 = _compr_8;
            if (((new BigNumber(201)).isLessThanOrEqualTo(_34_v1)) && ((_34_v1).isLessThan(new BigNumber(952)))) {
              _coll8.push([(_34_v1).plus(new BigNumber(637)),_dafny.Seq.of(true, false, true)]);
            }
          }
          return _coll8;
        }();
      }
    };
    static fm38(p0, p1, p2, globalState) {
      return _module.D2.create_DC7(!(false), !((_dafny.Set.fromElements(new BigNumber(664))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(false, false), _dafny.Set.fromElements(false))).length)))));
    };
    static fm39(p0, globalState) {
      let _source2 = _module.D5.create_DC14(!(false));
      if (_source2.is_DC13) {
        let _35___mcc_h0 = (_source2).cf20;
        let _36___mcc_h1 = (_source2).cf21;
        let _37___mcc_h2 = (_source2).cf22;
        let _38_cf22 = _37___mcc_h2;
        let _39_cf21 = _36___mcc_h1;
        let _40_cf20 = _35___mcc_h0;
        return _module.D6.create_DC15(function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(153), new BigNumber(-772))) {
    let _41_v0 = _compr_9;
    if (((new BigNumber(153)).isLessThanOrEqualTo(_41_v0)) && ((_41_v0).isLessThan(new BigNumber(-772)))) {
      _coll9.push([(_41_v0).plus(_40_cf20),_dafny.Seq.of(_39_cf21)]);
    }
  }
  return _coll9;
}());
      } else if (_source2.is_DC14) {
        let _42___mcc_h3 = (_source2).cf23;
        let _43_cf23 = _42___mcc_h3;
        return _module.D6.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),_dafny.Seq.of(_43_cf23)));
      } else {
        let _44___mcc_h4 = (_source2).cf19;
        let _45_cf19 = _44___mcc_h4;
        return _module.D6.create_DC15(function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(555), new BigNumber(490))) {
    let _46_v1 = _compr_10;
    if (((new BigNumber(555)).isLessThanOrEqualTo(_46_v1)) && ((_46_v1).isLessThan(new BigNumber(490)))) {
      _coll10.push([(_46_v1).plus(new BigNumber(355)),_dafny.Seq.of(true, false, true, false)]);
    }
  }
  return _coll10;
}());
      }
    };
    static fm40(p0, globalState) {
      return _module.D5.create_DC13(((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vsqstwekh")).length)))).plus(new BigNumber(-868)), false, true);
    };
    static fm41(globalState) {
      return (_dafny.Set.fromElements(false)).Union((_dafny.Set.fromElements(true, !(true), false, !(false), false)).Intersect(_dafny.Set.fromElements(!(true), true, true)));
    };
    static fm42(p0, p1, globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(new BigNumber(898), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-776)), function (_47_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length))) : (_dafny.MultiSet.fromElements(new BigNumber(827))))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(-34)));
    };
    static fm43(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(_dafny.MultiSet.fromElements(false, false, false, false));
    };
    static fm44(p0, globalState) {
      if (false) {
        return _module.D4.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), function (_48_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}));
      } else {
        return _module.D4.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), function (_49_i1) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}));
      }
    };
    static fm45(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), function (_50_i0) {
        return new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)))).length);
      }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(89)), function (_51_i1) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("kqu")).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ehybm")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(485)), function (_52_i2) {
        return new BigNumber(588);
      })).length))));
    };
    static fm46(p0, p1, globalState) {
      return ((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-665),new BigNumber(731))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(413),new BigNumber(65))));
    };
    static fm47(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(542),false))).length), new BigNumber(396)),_dafny.Seq.of(new BigNumber(-939), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length))));
    };
    static fm48(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("uarmdrl"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("lpsc")));
    };
    static fm49(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((!(!(!(true)))) ? (true) : (true)),new BigNumber(657));
    };
    static fm50(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(_module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('u'.codePointAt(0)))))).Union(_dafny.MultiSet.fromElements(_module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe((_module.D13.create_DC29(false, false)).dtor_cf49,new _dafny.CodePoint('i'.codePointAt(0)))), _module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('y'.codePointAt(0)))), _module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(true))),new _dafny.CodePoint('c'.codePointAt(0))))));
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of ((function () {
          let _coll12 = new _dafny.Map();
          for (const _compr_12 of (_dafny.Seq.of(function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of _dafny.IntegerRange(new BigNumber(297), new BigNumber(936))) {
              let _53_v1 = _compr_13;
              if (((new BigNumber(297)).isLessThanOrEqualTo(_53_v1)) && ((_53_v1).isLessThan(new BigNumber(936)))) {
                _coll13.push([_module.__default.safeModuloInt(_53_v1, new BigNumber(-866)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(292))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).cardinality())]);
              }
            }
            return _coll13;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),(_module.D11.create_DC23(!(true), new BigNumber(-234), _dafny.Map.Empty.slice().updateUnsafe(true,true), _module.D5.create_DC14(false))).dtor_cf37))).Elements) {
            let _54_v0 = _compr_12;
            if (_dafny.Seq.contains(_dafny.Seq.of(function () {
              let _coll14 = new _dafny.Map();
              for (const _compr_14 of _dafny.IntegerRange(new BigNumber(297), new BigNumber(936))) {
                let _53_v1 = _compr_14;
                if (((new BigNumber(297)).isLessThanOrEqualTo(_53_v1)) && ((_53_v1).isLessThan(new BigNumber(936)))) {
                  _coll14.push([_module.__default.safeModuloInt(_53_v1, new BigNumber(-866)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(292))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).cardinality())]);
                }
              }
              return _coll14;
            }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),(_module.D11.create_DC23(!(true), new BigNumber(-234), _dafny.Map.Empty.slice().updateUnsafe(true,true), _module.D5.create_DC14(false))).dtor_cf37)), _54_v0)) {
              _coll12.push([_54_v0,_dafny.Seq.UnicodeFromString("ccyc")]);
            }
          }
          return _coll12;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber(973)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-461)), function (_55_i0) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })))).Keys.Elements) {
          let _56_v2 = _compr_11;
          if (((function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of (_dafny.Seq.of(function () {
              let _coll16 = new _dafny.Map();
              for (const _compr_16 of _dafny.IntegerRange(new BigNumber(297), new BigNumber(936))) {
                let _53_v1 = _compr_16;
                if (((new BigNumber(297)).isLessThanOrEqualTo(_53_v1)) && ((_53_v1).isLessThan(new BigNumber(936)))) {
                  _coll16.push([_module.__default.safeModuloInt(_53_v1, new BigNumber(-866)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(292))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).cardinality())]);
                }
              }
              return _coll16;
            }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),(_module.D11.create_DC23(!(true), new BigNumber(-234), _dafny.Map.Empty.slice().updateUnsafe(true,true), _module.D5.create_DC14(false))).dtor_cf37))).Elements) {
              let _54_v0 = _compr_15;
              if (_dafny.Seq.contains(_dafny.Seq.of(function () {
                let _coll17 = new _dafny.Map();
                for (const _compr_17 of _dafny.IntegerRange(new BigNumber(297), new BigNumber(936))) {
                  let _53_v1 = _compr_17;
                  if (((new BigNumber(297)).isLessThanOrEqualTo(_53_v1)) && ((_53_v1).isLessThan(new BigNumber(936)))) {
                    _coll17.push([_module.__default.safeModuloInt(_53_v1, new BigNumber(-866)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(292))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).cardinality())]);
                  }
                }
                return _coll17;
              }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),(_module.D11.create_DC23(!(true), new BigNumber(-234), _dafny.Map.Empty.slice().updateUnsafe(true,true), _module.D5.create_DC14(false))).dtor_cf37)), _54_v0)) {
                _coll15.push([_54_v0,_dafny.Seq.UnicodeFromString("ccyc")]);
              }
            }
            return _coll15;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber(973)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-461)), function (_55_i0) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })))).contains(_56_v2)) {
            _coll11.add(_56_v2);
          }
        }
        return _coll11;
      }();
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('b'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('f'.codePointAt(0))));
    };
    static fm53(globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ttv"), _dafny.Seq.UnicodeFromString("ctxpl"), _dafny.Seq.UnicodeFromString("ljm"))).Union(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(505)), function (_57_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })));
    };
    static fm54(p0, globalState) {
      return function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of (function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Set.fromElements(new BigNumber((function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).Elements) {
              let _59_v1 = _compr_20;
              if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).contains(_59_v1)) {
                _coll20.push([(_59_v1).multipliedBy(new BigNumber(666)),!(true)]);
              }
            }
            return _coll20;
          }()).length)), _dafny.Set.fromElements(new BigNumber(296)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, !(true))).length)))).Elements) {
            let _60_v0 = _compr_19;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Set.fromElements(new BigNumber((function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).Elements) {
                let _59_v1 = _compr_21;
                if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).contains(_59_v1)) {
                  _coll21.push([(_59_v1).multipliedBy(new BigNumber(666)),!(true)]);
                }
              }
              return _coll21;
            }()).length)), _dafny.Set.fromElements(new BigNumber(296)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, !(true))).length))), _60_v0)) {
              _coll19.push([_60_v0,_dafny.Seq.UnicodeFromString("bgfnoiq")]);
            }
          }
          return _coll19;
        }()).Keys.Elements) {
          let _61_v2 = _compr_18;
          if ((function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of (_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Set.fromElements(new BigNumber((function () {
              let _coll23 = new _dafny.Map();
              for (const _compr_23 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).Elements) {
                let _59_v1 = _compr_23;
                if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).contains(_59_v1)) {
                  _coll23.push([(_59_v1).multipliedBy(new BigNumber(666)),!(true)]);
                }
              }
              return _coll23;
            }()).length)), _dafny.Set.fromElements(new BigNumber(296)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, !(true))).length)))).Elements) {
              let _60_v0 = _compr_22;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Set.fromElements(new BigNumber((function () {
                let _coll24 = new _dafny.Map();
                for (const _compr_24 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).Elements) {
                  let _59_v1 = _compr_24;
                  if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_58_i0) {
                    return new _dafny.CodePoint('p'.codePointAt(0));
                  })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(47))).length))).contains(_59_v1)) {
                    _coll24.push([(_59_v1).multipliedBy(new BigNumber(666)),!(true)]);
                  }
                }
                return _coll24;
              }()).length)), _dafny.Set.fromElements(new BigNumber(296)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, !(true))).length))), _60_v0)) {
                _coll22.push([_60_v0,_dafny.Seq.UnicodeFromString("bgfnoiq")]);
              }
            }
            return _coll22;
          }()).contains(_61_v2)) {
            _coll18.add(_61_v2);
          }
        }
        return _coll18;
      }();
    };
    static fm55(p0, p1, p2, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(8)), function (_62_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('f'.codePointAt(0)));
      });
    };
    static fm56(p0, p1, globalState) {
      return (function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(784),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(-19)), _dafny.Seq.of(new BigNumber(428)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, true, false)).length), new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(437)))).cardinality()))).Keys.Elements) {
          let _63_v0 = _compr_25;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(784),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(-19)), _dafny.Seq.of(new BigNumber(428)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, true, false)).length), new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(437)))).cardinality()))).contains(_63_v0)) {
            _coll25.push([_module.__default.safeDivisionInt(_63_v0, new BigNumber(581)),true]);
          }
        }
        return _coll25;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll26 = new _dafny.Map();
        for (const _compr_26 of (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),true)).Keys.Elements) {
            let _64_v2 = _compr_27;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),true)).contains(_64_v2)) {
              _coll27.add((_64_v2).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ss")).length))).length)))));
            }
          }
          return _coll27;
        }()).length), new BigNumber((_dafny.Seq.UnicodeFromString("wqn")).length), new BigNumber((_dafny.Seq.UnicodeFromString("dltidiom")).length), new BigNumber(399))).Elements) {
          let _65_v1 = _compr_26;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),true)).Keys.Elements) {
              let _66_v2 = _compr_28;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),true)).contains(_66_v2)) {
                _coll28.add((_66_v2).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ss")).length))).length)))));
              }
            }
            return _coll28;
          }()).length), new BigNumber((_dafny.Seq.UnicodeFromString("wqn")).length), new BigNumber((_dafny.Seq.UnicodeFromString("dltidiom")).length), new BigNumber(399))).contains(_65_v1)) {
            _coll26.push([(_65_v1).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(185)), function (_67_i0) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })).length)),_dafny.MultiSet.fromElements(new BigNumber(148), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber(10), new BigNumber(813), new BigNumber(-879))]);
          }
        }
        return _coll26;
      }()).length),true)));
    };
    static fm57(globalState) {
      return _module.D19.create_DC47((_dafny.Set.fromElements(!(false), false)).IsSubsetOf(_dafny.Set.fromElements(false)));
    };
    static fm58(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(true),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(true),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(false),false)));
    };
    static fm59(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true), _module.D5.create_DC14(true), _module.D5.create_DC14(false)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(655),new BigNumber(-349)))).Merge((function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(false)),true)).Keys.Elements) {
          let _68_v0 = _compr_29;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(false)),true)).contains(_68_v0)) {
            _coll29.push([_68_v0,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-697),new BigNumber(-69))]);
          }
        }
        return _coll29;
      }()).Merge((_module.D31.create_DC68(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),function () {
  let _coll30 = new _dafny.Map();
  for (const _compr_30 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(109),!(false))).length))).Elements) {
    let _69_v1 = _compr_30;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(109),!(false))).length)), _69_v1)) {
      _coll30.push([(_69_v1).multipliedBy(new BigNumber(-893)),new BigNumber((_dafny.Seq.UnicodeFromString("qubdos")).length)]);
    }
  }
  return _coll30;
}()))).dtor_cf112));
    };
    static fm60(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(219)), function (_70_i0) {
        return new BigNumber(486);
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(449)), function (_71_i1) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(171)), function (_72_i2) {
          return new BigNumber(-574);
        });
      })), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), function (_73_i3) {
        return new BigNumber(337);
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(974)), function (_74_i4) {
        return _dafny.Seq.of(new BigNumber(-185), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber(734), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(913)), function (_75_i5) {
          return new BigNumber(781);
        })).length), new BigNumber(-302));
      })));
    };
    static fm61(p0, p1, p2, p3, globalState) {
      return ((function () {
        let _coll31 = new _dafny.Set();
        for (const _compr_31 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-252)), function (_76_i0) {
          return new BigNumber((function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of (_dafny.Set.fromElements(new BigNumber(-448), new BigNumber(124))).Elements) {
              let _77_v0 = _compr_32;
              if ((_dafny.Set.fromElements(new BigNumber(-448), new BigNumber(124))).contains(_77_v0)) {
                _coll32.add((_77_v0).plus(new BigNumber(212)));
              }
            }
            return _coll32;
          }()).length);
        })).length)))).Elements) {
          let _78_v1 = _compr_31;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-252)), function (_76_i0) {
            return new BigNumber((function () {
              let _coll33 = new _dafny.Set();
              for (const _compr_33 of (_dafny.Set.fromElements(new BigNumber(-448), new BigNumber(124))).Elements) {
                let _79_v0 = _compr_33;
                if ((_dafny.Set.fromElements(new BigNumber(-448), new BigNumber(124))).contains(_79_v0)) {
                  _coll33.add((_79_v0).plus(new BigNumber(212)));
                }
              }
              return _coll33;
            }()).length);
          })).length)))).contains(_78_v1)) {
            _coll31.add(_78_v1);
          }
        }
        return _coll31;
      }()).Union(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("faysnx")).length), new BigNumber(-59), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-141)), function (_80_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber(494))).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-772), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-478)), function (_81_i2) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("qs")).length), new BigNumber(995))).length))).length))))).Difference((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(88)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, true, false)).length)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, true, false, false)).cardinality())), _dafny.Seq.of(new BigNumber(-711), new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("eavsgwtj")).length)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber((_dafny.Seq.of(true)).length)),new BigNumber(-41))).length)), new BigNumber(753)))).Union(function () {
        let _coll34 = new _dafny.Set();
        for (const _compr_34 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(837)),_module.D5.create_DC14(false))).Keys.Elements) {
          let _82_v2 = _compr_34;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(837)),_module.D5.create_DC14(false))).contains(_82_v2)) {
            _coll34.add(_82_v2);
          }
        }
        return _coll34;
      }()));
    };
    static fm62(p0, globalState) {
      let _source3 = _module.D4.create_DC11((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false))).length)));
      if (_source3.is_DC11) {
        let _83___mcc_h0 = (_source3).cf18;
        let _84_cf18 = _83___mcc_h0;
        return function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of (function () {
              let _coll37 = new _dafny.Map();
              for (const _compr_37 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).Keys.Elements) {
                let _85_v1 = _compr_37;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).contains(_85_v1)) {
                  _coll37.push([_85_v1,_84_cf18]);
                }
              }
              return _coll37;
            }()).Keys.Elements) {
              let _86_v2 = _compr_36;
              if ((function () {
                let _coll38 = new _dafny.Map();
                for (const _compr_38 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).Keys.Elements) {
                  let _85_v1 = _compr_38;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).contains(_85_v1)) {
                    _coll38.push([_85_v1,_84_cf18]);
                  }
                }
                return _coll38;
              }()).contains(_86_v2)) {
                _coll36.add(_86_v2);
              }
            }
            return _coll36;
          }()).Elements) {
            let _87_v0 = _compr_35;
            if ((function () {
              let _coll39 = new _dafny.Set();
              for (const _compr_39 of (function () {
                let _coll40 = new _dafny.Map();
                for (const _compr_40 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).Keys.Elements) {
                  let _85_v1 = _compr_40;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).contains(_85_v1)) {
                    _coll40.push([_85_v1,_84_cf18]);
                  }
                }
                return _coll40;
              }()).Keys.Elements) {
                let _88_v2 = _compr_39;
                if ((function () {
                  let _coll41 = new _dafny.Map();
                  for (const _compr_41 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).Keys.Elements) {
                    let _85_v1 = _compr_41;
                    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).contains(_85_v1)) {
                      _coll41.push([_85_v1,_84_cf18]);
                    }
                  }
                  return _coll41;
                }()).contains(_88_v2)) {
                  _coll39.add(_88_v2);
                }
              }
              return _coll39;
            }()).contains(_87_v0)) {
              _coll35.push([_87_v0,true]);
            }
          }
          return _coll35;
        }();
      } else {
        let _89___mcc_h1 = (_source3).cf17;
        let _90_cf17 = _89___mcc_h1;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true));
      }
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _91_v0;
      let _nw0 = Array((new BigNumber(16)).toNumber()).fill([]);
      _91_v0 = _nw0;
      let _92_v1;
      _92_v1 = _dafny.Map.Empty.slice().updateUnsafe(_91_v0,!(p0).isEqualTo(p3));
      if ((((_92_v1).contains(_91_v0)) ? ((_92_v1).get(_91_v0)) : (!(p2) || (!(p2))))) {
        let _93_v2;
        _93_v2 = _dafny.Set.fromElements(false, p2);
        let _94_v3;
        _94_v3 = _dafny.Map.Empty.slice().updateUnsafe(_93_v2,p0);
        let _95_v4;
        _95_v4 = _dafny.Seq.of(p3);
        let _96_v5;
        _96_v5 = _dafny.Seq.of(p2, p2);
        let _97_v6;
        _97_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        let _98_v7;
        _98_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_95_v4).length),new BigNumber((_dafny.Seq.update(_96_v5, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_97_v6).length)), new BigNumber((_96_v5).length)), false)).length));
        _94_v3 = (_94_v3).update(_dafny.Set.fromElements(p2), (((_97_v6).contains(new BigNumber(-908))) ? ((_97_v6).get(new BigNumber(-908))) : (p0)));
        let _99_v8;
        let _init0 = ((_100_p2) => function (_101_i0) {
          return _100_p2;
        })(p2);
        let _nw1 = Array((new BigNumber(17)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _99_v8 = _nw1;
        let _102_v9;
        _102_v9 = _dafny.Seq.of(_99_v8);
        let _103_v10;
        _103_v10 = _module.D27.create_DC63(new BigNumber((p1).cardinality()), p0, p0, p2);
        let _104_v11;
        _104_v11 = _dafny.Seq.of(_95_v4, _95_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), ((_105_p0) => function (_106_i2) {
          return _105_p0;
        })(p0)), _95_v4);
        let _107_v12;
        _107_v12 = _dafny.Map.Empty.slice().updateUnsafe((_103_v10).dtor_cf107,_104_v11);
        let _108_v13;
        _108_v13 = _module.D1.create_DC5(p3, p0, (((_107_v12).contains(true)) ? ((_107_v12).get(true)) : (_104_v11)), _99_v8, p0);
        let _109_v14;
        _109_v14 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p3)).length));
        let _110_v15;
        _110_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,_99_v8);
        let _pat_let_tv0 = _104_v11;
        let _pat_let_tv1 = p3;
        let _111_v16;
        let _nw2 = Array((new BigNumber(17)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _module.D1.create_DC5(p3, p3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(844)), ((_112_v4) => function (_113_i1) {
  return _112_v4;
})(_95_v4)), _99_v8, new BigNumber(685));
        _nw2[(_dafny.ONE).toNumber()] = _108_v13;
        _nw2[(new BigNumber(2)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(3)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(4)).toNumber()] = function (_pat_let0_0) {
          return function (_114_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_115_dt__update_hcf9_h0) {
                return _module.D1.create_DC5((_114_dt__update__tmp_h0).dtor_cf7, (_114_dt__update__tmp_h0).dtor_cf8, _115_dt__update_hcf9_h0, (_114_dt__update__tmp_h0).dtor_cf10, (_114_dt__update__tmp_h0).dtor_cf11);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_108_v13);
        _nw2[(new BigNumber(5)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(6)).toNumber()] = ((p2) ? (_108_v13) : (_108_v13));
        _nw2[(new BigNumber(7)).toNumber()] = ((p2) ? (_108_v13) : (_module.D1.create_DC5(p0, p3, _104_v11, _99_v8, (_dafny.ZERO).minus(p3))));
        _nw2[(new BigNumber(8)).toNumber()] = _module.D1.create_DC5(p3, new BigNumber(806), _104_v11, _99_v8, p0);
        _nw2[(new BigNumber(9)).toNumber()] = _module.D1.create_DC5(p0, p3, _dafny.Seq.of(_dafny.Seq.of(new BigNumber((p1).cardinality()), p3, p0, p3, new BigNumber((_109_v14).length))), (((_110_v15).contains(p2)) ? ((_110_v15).get(p2)) : (_99_v8)), p3);
        _nw2[(new BigNumber(10)).toNumber()] = _module.D1.create_DC5(p3, p3, _104_v11, _99_v8, p0);
        _nw2[(new BigNumber(11)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(12)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(13)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(14)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(15)).toNumber()] = _108_v13;
        _nw2[(new BigNumber(16)).toNumber()] = function (_pat_let2_0) {
          return function (_116_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_117_dt__update_hcf11_h0) {
                return _module.D1.create_DC5((_116_dt__update__tmp_h1).dtor_cf7, (_116_dt__update__tmp_h1).dtor_cf8, (_116_dt__update__tmp_h1).dtor_cf9, (_116_dt__update__tmp_h1).dtor_cf10, _117_dt__update_hcf11_h0);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_108_v13);
        _111_v16 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_111_v16).length));
        (_111_v16)[_index0] = _108_v13;
        let _118_v17;
        _118_v17 = _dafny.Seq.of(_dafny.Seq.of(_99_v8), _102_v9, _102_v9, _102_v9, _102_v9);
        let _119_v18;
        _119_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_118_v17)[_module.__default.safeIndex(p3, new BigNumber((_118_v17).length))]);
        let _120_v19;
        _120_v19 = _dafny.Seq.UnicodeFromString("tlsyfds");
        let _index1 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_111_v16).length));
        let _rhs0 = _dafny.Seq.Concat(_102_v9, (((_119_v18).contains(new BigNumber((_120_v19).length))) ? ((_119_v18).get(new BigNumber((_120_v19).length))) : (_102_v9)));
        let _rhs1 = _108_v13;
        let _lhs0 = _111_v16;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_111_v16).length));
        _102_v9 = _rhs0;
        _lhs0[_lhs1] = _rhs1;
        let _121_v20;
        _121_v20 = _module.D0.create_DC0(p1);
        let _122_v21;
        _122_v21 = _module.D0.create_DC0((p1).Union((_121_v20).dtor_cf0));
        let _source4 = _122_v21;
        if (_source4.is_DC1) {
          let _123___mcc_h0 = (_source4).cf1;
          let _124_cf1 = _123___mcc_h0;
          let _125_v22;
          _125_v22 = _module.D9.create_DC20(_dafny.Seq.Create(_module.__default.abs(new BigNumber(86)), function (_126_i3) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), p2, p3);
          let _127_v23;
          _127_v23 = _dafny.MultiSet.fromElements(_125_v22);
          let _128_v24;
          _128_v24 = _module.D21.create_DC50(_127_v23);
          let _129_v25;
          _129_v25 = _dafny.Map.Empty.slice().updateUnsafe((_96_v5)[_module.__default.safeIndex(p0, new BigNumber((_96_v5).length))],p2);
          let _130_v26;
          _130_v26 = _module.D5.create_DC14(p2);
          let _131_v27;
          _131_v27 = _module.D11.create_DC23(p2, _124_cf1, _dafny.Map.Empty.slice().updateUnsafe(p2,p2), _130_v26);
          let _132_v28;
          let _nw3 = Array((new BigNumber(15)).toNumber());
          _nw3[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          _nw3[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          _nw3[(new BigNumber(2)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(3)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(4)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(5)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(6)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(7)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(8)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(9)).toNumber()] = (_131_v27).dtor_cf38;
          _nw3[(new BigNumber(10)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(11)).toNumber()] = (_131_v27).dtor_cf38;
          _nw3[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_96_v5)[_module.__default.safeIndex(p3, new BigNumber((_96_v5).length))],p2);
          _nw3[(new BigNumber(13)).toNumber()] = _129_v25;
          _nw3[(new BigNumber(14)).toNumber()] = _129_v25;
          _132_v28 = _nw3;
          let _133_v29;
          _133_v29 = _dafny.MultiSet.fromElements(_120_v19);
          let _134_v30;
          _134_v30 = _133_v29;
          r0 = (_124_cf1).minus((_dafny.ZERO).minus((_module.D23.create_DC55(_128_v24, _132_v28, _134_v30, p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).length))).dtor_cf89));
          let _pat_let_tv2 = _93_v2;
          let _pat_let_tv3 = _93_v2;
          let _rhs2 = p2;
          let _rhs3 = function (_pat_let4_0) {
            return function (_135_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_136_dt__update_hcf23_h0) {
                  return _module.D5.create_DC14(_136_dt__update_hcf23_h0);
                }(_pat_let5_0);
              }((_pat_let_tv2).IsSubsetOf(_pat_let_tv3));
            }(_pat_let4_0);
          }(_130_v26);
          let _rhs4 = _120_v19;
          let _rhs5 = _module.__default.safeDivisionInt((p3).minus(new BigNumber((_96_v5).length)), p0);
          let _lhs2 = globalState;
          _lhs2.f13 = _rhs2;
          _130_v26 = _rhs3;
          _120_v19 = _rhs4;
          r0 = _rhs5;
          let _index2 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_99_v8).length));
          (_99_v8)[_index2] = (!(p2)) || (p2);
          let _137_v31;
          let _nw4 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _137_v31 = _nw4;
          let _index3 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_99_v8).length));
          let _rhs6 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-374)), ((_138_v19, _139_p0) => function (_140_i4) {
            return (_138_v19)[_module.__default.safeIndex(_139_p0, new BigNumber((_138_v19).length))];
          })(_120_v19, p0)), _120_v19), _dafny.Seq.Concat(_120_v19, _dafny.Seq.UnicodeFromString("cnqxtfmaj"))), _module.__default.safeIndex(_module.__default.safeDivisionInt(p3, _124_cf1), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-374)), ((_141_v19, _142_p0) => function (_143_i4) {
            return (_141_v19)[_module.__default.safeIndex(_142_p0, new BigNumber((_141_v19).length))];
          })(_120_v19, p0)), _120_v19), _dafny.Seq.Concat(_120_v19, _dafny.Seq.UnicodeFromString("cnqxtfmaj")))).length)), new _dafny.CodePoint('q'.codePointAt(0)));
          let _rhs7 = (_124_cf1).plus(p0);
          let _rhs8 = ((p2) ? (p2) : (p2));
          let _rhs9 = p2;
          let _rhs10 = _137_v31;
          let _lhs3 = globalState;
          let _lhs4 = _99_v8;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_99_v8).length));
          _120_v19 = _rhs6;
          _lhs3.f7 = _rhs7;
          r1 = _rhs8;
          _lhs4[_lhs5] = _rhs9;
          _137_v31 = _rhs10;
          let _144_v32;
          _144_v32 = new _dafny.CodePoint('q'.codePointAt(0));
          let _145_v33;
          let _nw5 = new _module.C6();
          _nw5.__ctor(_144_v32, (_96_v5)[_module.__default.safeIndex(p0, new BigNumber((_96_v5).length))]);
          _145_v33 = _nw5;
          let _146_v34;
          let _nw6 = Array((new BigNumber(4)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _145_v33.f23;
          _nw6[(_dafny.ONE).toNumber()] = _145_v33.f23;
          _nw6[(new BigNumber(2)).toNumber()] = _144_v32;
          _nw6[(new BigNumber(3)).toNumber()] = _145_v33.f23;
          _146_v34 = _nw6;
          let _nw7 = new _module.C10();
          _nw7.__ctor(p2, _module.__default.safeDivisionInt(new BigNumber(872), p0), _144_v32, p2, _146_v34, _module.__default.fm2(globalState));
          _145_v33 = _nw7;
        } else if (_source4.is_DC2) {
          let _147___mcc_h1 = (_source4).cf2;
          let _148___mcc_h2 = (_source4).cf3;
          let _149___mcc_h3 = (_source4).cf4;
          let _150_cf4 = _149___mcc_h3;
          let _151_cf3 = _148___mcc_h2;
          let _152_cf2 = _147___mcc_h1;
          let _153_v35;
          let _nw8 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _153_v35 = _nw8;
          let _154_v36;
          _154_v36 = _dafny.Seq.of(_153_v35, _153_v35);
          let _155_v37;
          _155_v37 = _module.D18.create_DC42((p3).minus(p3), _dafny.Seq.contains(_154_v36, _153_v35));
          _155_v37 = _155_v37;
          let _156_v38;
          _156_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_120_v19, _120_v19),p0);
          _156_v38 = (_156_v38).update((_152_cf2) === (false), p0);
          let _157_v39;
          let _nw9 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _157_v39 = _nw9;
          let _158_v40;
          _158_v40 = _dafny.Map.Empty.slice().updateUnsafe(_150_cf4,_152_cf2);
          let _index4 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_157_v39).length));
          (_157_v39)[_index4] = _158_v40;
          let _index5 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_157_v39).length));
          (_157_v39)[_index5] = _158_v40;
          let _159_v41;
          _159_v41 = _module.D19.create_DC47(true);
          let _160_v42;
          _160_v42 = _module.D19.create_DC48(_159_v41);
          let _161_v43;
          _161_v43 = _module.D19.create_DC48(_160_v42);
          let _pat_let_tv4 = _160_v42;
          _161_v43 = function (_pat_let6_0) {
            return function (_162_dt__update__tmp_h3) {
              return function (_pat_let7_0) {
                return function (_163_dt__update_hcf78_h0) {
                  return _module.D19.create_DC48(_163_dt__update_hcf78_h0);
                }(_pat_let7_0);
              }(_pat_let_tv4);
            }(_pat_let6_0);
          }(_161_v43);
        } else if (_source4.is_DC0) {
          let _164___mcc_h4 = (_source4).cf0;
          let _165_cf0 = _164___mcc_h4;
          let _166_v44;
          _166_v44 = new _dafny.CodePoint('f'.codePointAt(0));
          let _167_v45;
          _167_v45 = _dafny.MultiSet.fromElements(_166_v44);
          _167_v45 = _167_v45;
          let _168_v46;
          let _nw10 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _168_v46 = _nw10;
          let _index6 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_168_v46).length));
          (_168_v46)[_index6] = _166_v44;
          let _index7 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_168_v46).length));
          (_168_v46)[_index7] = _166_v44;
          let _169_v47;
          _169_v47 = _dafny.MultiSet.fromElements(_120_v19, _120_v19);
          let _170_v48;
          _170_v48 = _dafny.Map.Empty.slice().updateUnsafe(p3,_169_v47);
          _169_v47 = (((_170_v48).contains((p3).plus(p0))) ? ((_170_v48).get((p3).plus(p0))) : (_dafny.MultiSet.fromElements(_120_v19)));
          let _171_v49;
          _171_v49 = _dafny.Seq.of(_120_v19, _120_v19);
          _120_v19 = _dafny.Seq.Concat(_dafny.Seq.Concat(_120_v19, _120_v19), (_171_v49)[_module.__default.safeIndex(p3, new BigNumber((_171_v49).length))]);
        } else {
          let _172___mcc_h5 = (_source4).cf5;
          let _173_cf5 = _172___mcc_h5;
          let _174_v50;
          _174_v50 = new _dafny.CodePoint('v'.codePointAt(0));
          let _175_v51;
          let _nw11 = new _module.C0();
          _nw11.__ctor(_module.__default.fm48(globalState), p2, _174_v50, p2);
          _175_v51 = _nw11;
          let _176_v52;
          let _nw12 = Array((new BigNumber(16)).toNumber());
          _nw12[(_dafny.ZERO).toNumber()] = _175_v51;
          _nw12[(_dafny.ONE).toNumber()] = _175_v51;
          _nw12[(new BigNumber(2)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(3)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(4)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(5)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(6)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(7)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(8)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(9)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(10)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(11)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(12)).toNumber()] = (_175_v51);
          _nw12[(new BigNumber(13)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(14)).toNumber()] = _175_v51;
          _nw12[(new BigNumber(15)).toNumber()] = _175_v51;
          _176_v52 = _nw12;
          let _index8 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_176_v52).length));
          (_176_v52)[_index8] = _175_v51;
          let _index9 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_176_v52).length));
          (_176_v52)[_index9] = _175_v51;
          let _177_v53;
          let _nw13 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
          _177_v53 = _nw13;
          let _178_v55;
          _178_v55 = _dafny.Set.fromElements(_175_v51.f23);
          let _index10 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_177_v53).length));
          (_177_v53)[_index10] = (function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of (function () {
              let _coll43 = new _dafny.Map();
              for (const _compr_43 of (_178_v55).Elements) {
                let _179_v54 = _compr_43;
                if ((_178_v55).contains(_179_v54)) {
                  _coll43.push([_179_v54,!((_175_v51).f24)]);
                }
              }
              return _coll43;
            }()).Keys.Elements) {
              let _180_v56 = _compr_42;
              if ((function () {
                let _coll44 = new _dafny.Map();
                for (const _compr_44 of (_178_v55).Elements) {
                  let _179_v54 = _compr_44;
                  if ((_178_v55).contains(_179_v54)) {
                    _coll44.push([_179_v54,!((_175_v51).f24)]);
                  }
                }
                return _coll44;
              }()).contains(_180_v56)) {
                _coll42.add(_180_v56);
              }
            }
            return _coll42;
          }()).Intersect(_178_v55);
          let _index11 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_177_v53).length));
          (_177_v53)[_index11] = _178_v55;
          let _181_v57;
          let _init1 = ((_182_v51) => function (_183_i5) {
            return _182_v51.f23;
          })(_175_v51);
          let _nw14 = Array((new BigNumber(5)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
            _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _181_v57 = _nw14;
          let _184_v58;
          let _nw15 = new _module.C7();
          _nw15.__ctor((_175_v51).f24, true, _181_v57, _175_v51.f23, false, !((_dafny.MultiSet.fromElements(p0, new BigNumber(448), p0)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_95_v4))));
          _184_v58 = _nw15;
          _184_v58 = _184_v58;
          (globalState).f13 = (_184_v58).f31;
        }
        let _185_v59;
        _185_v59 = _97_v6;
        let _186_v60;
        let _init2 = function (_187_i6) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        };
        let _nw16 = Array((new BigNumber(29)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw16.length); _i0_2++) {
          _nw16[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _186_v60 = _nw16;
        let _188_v61;
        _188_v61 = _dafny.Map.Empty.slice().updateUnsafe(p2,_186_v60);
        let _189_v62;
        _189_v62 = new _dafny.CodePoint('s'.codePointAt(0));
        let _190_v63;
        let _nw17 = new _module.C4();
        _nw17.__ctor(new BigNumber((_120_v19).length), _185_v59, (((_188_v61).contains(false)) ? ((_188_v61).get(false)) : (_186_v60)), _189_v62, p2);
        _190_v63 = _nw17;
        (globalState).f13 = (((_190_v63).fm22(p0, p2, globalState)) ? (p2) : (p2));
      } else {
        r2 = _module.__default.safeDivisionInt(new BigNumber(-599), new BigNumber((_module.__default.fm62(p0, globalState)).length));
        r0 = p0;
        let _191_v64;
        let _nw18 = Array((new BigNumber(23)).toNumber()).fill(false);
        _191_v64 = _nw18;
        let _index12 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length));
        (_191_v64)[_index12] = false;
        let _index13 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length));
        (_191_v64)[_index13] = p2;
        let _192_v65;
        _192_v65 = _dafny.Seq.of(!((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))]));
        let _193_v66;
        _193_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,_192_v65);
        let _194_v67;
        _194_v67 = _dafny.Set.fromElements(new BigNumber(425));
        let _195_v68;
        _195_v68 = _dafny.Seq.UnicodeFromString("k");
        let _196_v69;
        _196_v69 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Seq.UnicodeFromString("pfiap")).length));
        let _197_v70;
        _197_v70 = _dafny.Seq.of(p3, p3, p0, new BigNumber(141), p3);
        let _198_v71;
        _198_v71 = _dafny.Set.fromElements((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))], (_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))]);
        let _199_v72;
        _199_v72 = _dafny.MultiSet.fromElements(_198_v71);
        let _200_v74;
        let _nw19 = Array((new BigNumber(29)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = (p3).multipliedBy(new BigNumber((_193_v66).length));
        _nw19[(_dafny.ONE).toNumber()] = new BigNumber((_194_v67).length);
        _nw19[(new BigNumber(2)).toNumber()] = p3;
        _nw19[(new BigNumber(3)).toNumber()] = new BigNumber(-354);
        _nw19[(new BigNumber(4)).toNumber()] = new BigNumber(-37);
        _nw19[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(p3, p3);
        _nw19[(new BigNumber(6)).toNumber()] = p3;
        _nw19[(new BigNumber(7)).toNumber()] = new BigNumber((_195_v68).length);
        _nw19[(new BigNumber(8)).toNumber()] = new BigNumber(641);
        _nw19[(new BigNumber(9)).toNumber()] = (new BigNumber((_196_v69).length)).minus((((_196_v69).contains(!((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))]))) ? ((_196_v69).get(!((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))]))) : (p3)));
        _nw19[(new BigNumber(10)).toNumber()] = p3;
        _nw19[(new BigNumber(11)).toNumber()] = new BigNumber((_197_v70).length);
        _nw19[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
        _nw19[(new BigNumber(13)).toNumber()] = ((true) ? (p3) : ((_dafny.ZERO).minus(p3)));
        _nw19[(new BigNumber(14)).toNumber()] = p3;
        _nw19[(new BigNumber(15)).toNumber()] = (((_199_v72).contains(_198_v71)) ? ((_199_v72).get(_198_v71)) : (new BigNumber(184)));
        _nw19[(new BigNumber(16)).toNumber()] = p3;
        _nw19[(new BigNumber(17)).toNumber()] = p0;
        _nw19[(new BigNumber(18)).toNumber()] = p0;
        _nw19[(new BigNumber(19)).toNumber()] = p3;
        _nw19[(new BigNumber(20)).toNumber()] = p0;
        _nw19[(new BigNumber(21)).toNumber()] = p3;
        _nw19[(new BigNumber(22)).toNumber()] = p0;
        _nw19[(new BigNumber(23)).toNumber()] = p3;
        _nw19[(new BigNumber(24)).toNumber()] = p3;
        _nw19[(new BigNumber(25)).toNumber()] = new BigNumber(-424);
        _nw19[(new BigNumber(26)).toNumber()] = p3;
        _nw19[(new BigNumber(27)).toNumber()] = new BigNumber(((function () {
          let _coll45 = new _dafny.Set();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(649), new BigNumber(62))) {
            let _201_v73 = _compr_45;
            if (((new BigNumber(649)).isLessThanOrEqualTo(_201_v73)) && ((_201_v73).isLessThan(new BigNumber(62)))) {
              _coll45.add(_module.__default.safeDivisionInt(_201_v73, new BigNumber(234)));
            }
          }
          return _coll45;
        }()).Intersect(_194_v67)).length);
        _nw19[(new BigNumber(28)).toNumber()] = p0;
        _200_v74 = _nw19;
        let _index14 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_200_v74).length));
        (_200_v74)[_index14] = p0;
        let _index15 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_200_v74).length));
        let _rhs11 = p0;
        let _rhs12 = (p0).multipliedBy(p3);
        let _lhs6 = _200_v74;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_200_v74).length));
        let _lhs8 = globalState;
        _lhs6[_lhs7] = _rhs11;
        _lhs8.f7 = _rhs12;
        let _202_v75;
        let _nw20 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
        _202_v75 = _nw20;
        let _index16 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_202_v75).length));
        (_202_v75)[_index16] = _198_v71;
        let _203_v76;
        _203_v76 = _dafny.Seq.of(_198_v71);
        let _204_v77;
        _204_v77 = new _dafny.CodePoint('h'.codePointAt(0));
        let _205_v78;
        _205_v78 = _dafny.Map.Empty.slice().updateUnsafe((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))],p2);
        let _index17 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_202_v75).length));
        let _rhs13 = _module.__default.fm2(globalState);
        let _rhs14 = ((_198_v71).Union(_198_v71)).Intersect(((_203_v76)[_module.__default.safeIndex(_module.__default.fm0(p2, new BigNumber((_dafny.Set.fromElements(_204_v77)).length), globalState), new BigNumber((_203_v76).length))]).Difference(_dafny.Set.fromElements((((_205_v78).contains((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))])) ? ((_205_v78).get((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))])) : ((_191_v64)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_191_v64).length))])), p2)));
        let _lhs9 = _202_v75;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_202_v75).length));
        r1 = _rhs13;
        _lhs9[_lhs10] = _rhs14;
      }
      let _206_v79;
      _206_v79 = new _dafny.CodePoint('a'.codePointAt(0));
      let _207_v80;
      let _nw21 = new _module.C6();
      _nw21.__ctor(_206_v79, true);
      _207_v80 = _nw21;
      let _208_v81;
      _208_v81 = _module.D23.create_DC54(_207_v80);
      _207_v80 = (_208_v81).dtor_cf84;
      let _209_v82;
      _209_v82 = _dafny.Seq.UnicodeFromString("yj");
      let _210_v83;
      _210_v83 = _module.D9.create_DC20(_209_v82, p2, p0);
      let _211_v84;
      _211_v84 = _dafny.MultiSet.fromElements(_210_v83, _210_v83, _210_v83);
      let _212_v85;
      _212_v85 = _module.D21.create_DC50(_211_v84);
      let _213_v86;
      _213_v86 = _dafny.Seq.of(_212_v85, _212_v85, _module.D21.create_DC50(_211_v84));
      _213_v86 = _213_v86;
      let _214_v87;
      let _init3 = ((_215_p2) => function (_216_i7) {
        return _215_p2;
      })(p2);
      let _nw22 = Array((new BigNumber(29)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw22.length); _i0_3++) {
        _nw22[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _214_v87 = _nw22;
      _214_v87 = _214_v87;
      let _index18 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_214_v87).length));
      (_214_v87)[_index18] = p2;
      let _217_v88;
      _217_v88 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_209_v82).length)),p2);
      let _218_v89;
      let _nw23 = Array((new BigNumber(7)).toNumber());
      _nw23[(_dafny.ZERO).toNumber()] = new BigNumber((_217_v88).length);
      _nw23[(_dafny.ONE).toNumber()] = p0;
      _nw23[(new BigNumber(2)).toNumber()] = p0;
      _nw23[(new BigNumber(3)).toNumber()] = p3;
      _nw23[(new BigNumber(4)).toNumber()] = new BigNumber((_209_v82).length);
      _nw23[(new BigNumber(5)).toNumber()] = p3;
      _nw23[(new BigNumber(6)).toNumber()] = p3;
      _218_v89 = _nw23;
      let _219_v90;
      _219_v90 = _dafny.Map.Empty.slice().updateUnsafe(_218_v89,p1);
      let _220_v91;
      _220_v91 = _dafny.Seq.of(p0, p0);
      let _221_v92;
      _221_v92 = _dafny.Seq.of(_220_v91, _220_v91);
      let _index19 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_214_v87).length));
      (_214_v87)[_index19] = ((((_219_v90).contains(_218_v89)) ? ((_219_v90).get(_218_v89)) : (p1))).IsDisjointFrom((p1).update(p2, _module.__default.abs(new BigNumber((_221_v92).length))));
      let _222_v93;
      _222_v93 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("nrtq"),_218_v89);
      r2 = new BigNumber(((_222_v93).Merge(_222_v93)).length);
      r0 = p0;
      r1 = (_214_v87)[_module.__default.safeIndex(new BigNumber(776), new BigNumber((_214_v87).length))];
      r2 = p0;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _223_v0;
      _223_v0 = true;
      let _224_v1;
      let _nw24 = Array((_dafny.ONE).toNumber());
      _nw24[(_dafny.ZERO).toNumber()] = _223_v0;
      _224_v1 = _nw24;
      let _225_v2;
      _225_v2 = new BigNumber(230);
      let _226_v3;
      _226_v3 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_224_v1,(_dafny.ZERO).minus(_225_v2)));
      let _227_v4;
      _227_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(298),_223_v0);
      let _228_v5;
      _228_v5 = _dafny.Map.Empty.slice().updateUnsafe(_225_v2,_225_v2);
      let _229_v6;
      _229_v6 = _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)));
      let _230_v7;
      _230_v7 = _dafny.Seq.of(_223_v0);
      let _231_v8;
      _231_v8 = _dafny.Map.Empty.slice().updateUnsafe(_223_v0,_223_v0);
      let _232_v9;
      _232_v9 = _dafny.Map.Empty.slice().updateUnsafe(_230_v7,true);
      let _233_v10;
      _233_v10 = _dafny.Seq.UnicodeFromString("wjoqku");
      let _234_v11;
      let _nw25 = Array((new BigNumber(26)).toNumber());
      _nw25[(_dafny.ZERO).toNumber()] = _225_v2;
      _nw25[(_dafny.ONE).toNumber()] = new BigNumber((_227_v4).length);
      _nw25[(new BigNumber(2)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(3)).toNumber()] = (((_228_v5).contains(_225_v2)) ? ((_228_v5).get(_225_v2)) : (_225_v2));
      _nw25[(new BigNumber(4)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(5)).toNumber()] = new BigNumber((_229_v6).length);
      _nw25[(new BigNumber(6)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(7)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(8)).toNumber()] = new BigNumber(-846);
      _nw25[(new BigNumber(9)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(10)).toNumber()] = new BigNumber(379);
      _nw25[(new BigNumber(11)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(12)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(13)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(14)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_225_v2);
      _nw25[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.update(_230_v7, _module.__default.safeIndex(new BigNumber((_230_v7).length), new BigNumber((_230_v7).length)), _223_v0)).length);
      _nw25[(new BigNumber(17)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_225_v2);
      _nw25[(new BigNumber(19)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_231_v8).length));
      _nw25[(new BigNumber(21)).toNumber()] = new BigNumber((_232_v9).length);
      _nw25[(new BigNumber(22)).toNumber()] = new BigNumber((_233_v10).length);
      _nw25[(new BigNumber(23)).toNumber()] = _225_v2;
      _nw25[(new BigNumber(24)).toNumber()] = new BigNumber((_227_v4).length);
      _nw25[(new BigNumber(25)).toNumber()] = _225_v2;
      _234_v11 = _nw25;
      let _235_v12;
      _235_v12 = _dafny.Set.fromElements(_224_v1);
      let _236_v13;
      _236_v13 = new _dafny.CodePoint('a'.codePointAt(0));
      let _237_v14;
      _237_v14 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_225_v2),_231_v8);
      let _238_globalState;
      let _nw26 = new _module.GlobalState();
      _nw26.__ctor(_224_v1, new BigNumber(986), (_226_v3)[_module.__default.safeIndex(_225_v2, new BigNumber((_226_v3).length))], true, _234_v11, false, _224_v1, new BigNumber(986), _233_v10, true, _235_v12, new BigNumber(737), _dafny.Seq.update(_233_v10, _module.__default.safeIndex(_225_v2, new BigNumber((_233_v10).length)), _236_v13), false, new _dafny.CodePoint('a'.codePointAt(0)), (_231_v8).Merge((((_237_v14).contains(_225_v2)) ? ((_237_v14).get(_225_v2)) : (_231_v8))), new BigNumber(108), _dafny.Seq.Concat(_233_v10, _233_v10), _dafny.Seq.UnicodeFromString("n"), false, new BigNumber(-137), _234_v11, _234_v11);
      _238_globalState = _nw26;
      let _hi0 = (_dafny.ZERO).minus(((true) ? ((_dafny.ZERO).minus(_225_v2)) : (_module.__default.fm0(_223_v0, _225_v2, _238_globalState))));
      for (let _239_i0 = new BigNumber(425); _239_i0.isLessThan(_hi0); _239_i0 = _239_i0.plus(_dafny.ONE)) {
        if (true) {
          (_238_globalState).f7 = _225_v2;
          (_238_globalState).f7 = new BigNumber((_dafny.Seq.UnicodeFromString("sulow")).length);
          _233_v10 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), function (_240_i1) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          });
          (_238_globalState).f7 = _225_v2;
          let _241_v15;
          _241_v15 = _dafny.MultiSet.fromElements(_223_v0);
          let _242_v16;
          let _243_v17;
          let _244_v18;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = _module.__default.m0((_dafny.ZERO).minus(_239_i0), (_241_v15).Intersect(_241_v15), _223_v0, new BigNumber((_230_v7).length), _238_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _242_v16 = _out0;
          _243_v17 = _out1;
          _244_v18 = _out2;
        } else {
          let _index20 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_224_v1).length));
          (_224_v1)[_index20] = _223_v0;
          let _245_v19;
          _245_v19 = _dafny.Seq.of(new BigNumber(313));
          let _246_v20;
          _246_v20 = _dafny.Seq.of(_245_v19, _245_v19, _245_v19);
          let _index21 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_224_v1).length));
          (_224_v1)[_index21] = _dafny.areEqual((_246_v20)[_module.__default.safeIndex((_dafny.ZERO).minus(_225_v2), new BigNumber((_246_v20).length))], _245_v19);
          let _index22 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_224_v1).length));
          (_224_v1)[_index22] = ((new BigNumber((_module.__default.fm1(_module.__default.fm2(_238_globalState), _238_globalState)).length)).plus(_225_v2)).isLessThanOrEqualTo(new BigNumber(591));
          let _247_v21;
          _247_v21 = _dafny.MultiSet.fromElements(_223_v0, !(_223_v0), false, _223_v0, _223_v0);
          let _248_v22;
          let _249_v23;
          let _250_v24;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = _module.__default.m0(_225_v2, _247_v21, (_224_v1)[_module.__default.safeIndex(new BigNumber(240), new BigNumber((_224_v1).length))], _239_i0, _238_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _248_v22 = _out3;
          _249_v23 = _out4;
          _250_v24 = _out5;
          let _251_v25;
          _251_v25 = _module.D0.create_DC0(_module.__default.fm3(_223_v0, _238_globalState));
          let _252_v26;
          _252_v26 = _dafny.Set.fromElements(_250_v24);
          let _253_v27;
          let _254_v28;
          let _255_v29;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = _module.__default.m0(new BigNumber((_dafny.Seq.UnicodeFromString("lrbqh")).length), ((_251_v25).dtor_cf0).Difference(_247_v21), _223_v0, (_250_v24).multipliedBy(new BigNumber((_252_v26).length)), _238_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _253_v27 = _out6;
          _254_v28 = _out7;
          _255_v29 = _out8;
          (_238_globalState).f7 = _225_v2;
        }
        let _256_v30;
        let _init4 = function (_257_i2) {
          return _dafny.Seq.UnicodeFromString("feg");
        };
        let _nw27 = Array((new BigNumber(27)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw27.length); _i0_4++) {
          _nw27[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _256_v30 = _nw27;
        let _index23 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_256_v30).length));
        (_256_v30)[_index23] = _dafny.Seq.Concat(_233_v10, _233_v10);
        let _258_v31;
        _258_v31 = _dafny.MultiSet.fromElements(_223_v0, _223_v0, _223_v0);
        let _index24 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_256_v30).length));
        (_256_v30)[_index24] = _module.__default.fm1((_225_v2).isLessThanOrEqualTo((((_258_v31).contains(!(_223_v0))) ? ((_258_v31).get(!(_223_v0))) : (_239_i0))), _238_globalState);
        let _259_v32;
        let _nw28 = Array((new BigNumber(5)).toNumber()).fill(_module.D0.Default());
        _259_v32 = _nw28;
        let _index25 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_259_v32).length));
        (_259_v32)[_index25] = _module.D0.create_DC2(_223_v0, _239_i0, _223_v0);
        let _260_v34;
        _260_v34 = _dafny.Set.fromElements(_225_v2);
        let _261_v35;
        _261_v35 = _module.D0.create_DC2(!(_223_v0), new BigNumber((function () {
  let _coll46 = new _dafny.Map();
  for (const _compr_46 of (_260_v34).Elements) {
    let _262_v33 = _compr_46;
    if ((_260_v34).contains(_262_v33)) {
      _coll46.push([_module.__default.safeDivisionInt(_262_v33, _239_i0),_236_v13]);
    }
  }
  return _coll46;
}()).length), _223_v0);
        let _pat_let_tv5 = _223_v0;
        let _index26 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_259_v32).length));
        (_259_v32)[_index26] = function (_pat_let8_0) {
          return function (_263_dt__update__tmp_h0) {
            return function (_pat_let9_0) {
              return function (_264_dt__update_hcf2_h0) {
                return _module.D0.create_DC2(_264_dt__update_hcf2_h0, (_263_dt__update__tmp_h0).dtor_cf3, (_263_dt__update__tmp_h0).dtor_cf4);
              }(_pat_let9_0);
            }(!(_pat_let_tv5));
          }(_pat_let8_0);
        }(_261_v35);
        let _rhs15 = _239_i0;
        let _rhs16 = _module.__default.fm0(false, _239_i0, _238_globalState);
        let _rhs17 = ((_dafny.MultiSet.fromElements(_223_v0, _223_v0)).Union(_dafny.MultiSet.fromElements(_223_v0, _223_v0))).Intersect(_dafny.MultiSet.fromElements(_223_v0, !(_223_v0), true));
        let _rhs18 = _223_v0;
        let _lhs11 = _238_globalState;
        let _lhs12 = _238_globalState;
        _lhs11.f7 = _rhs15;
        _225_v2 = _rhs16;
        _258_v31 = _rhs17;
        _lhs12.f13 = _rhs18;
      }
      (_238_globalState).f13 = !((_225_v2).isLessThanOrEqualTo(_225_v2)) || (!(_dafny.Seq.IsPrefixOf(_233_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(215)), ((_265_v13) => function (_266_i3) {
        return _265_v13;
      })(_236_v13)))));
      let _267_v36;
      _267_v36 = _dafny.Set.fromElements(_223_v0);
      let _268_v37;
      _268_v37 = _dafny.Set.fromElements(_module.__default.fm0(_223_v0, _225_v2, _238_globalState), _225_v2);
      let _269_v38;
      _269_v38 = _dafny.Set.fromElements(new BigNumber(462), new BigNumber((_268_v37).length), _225_v2, _225_v2, _225_v2);
      _225_v2 = (_225_v2).minus(_module.__default.safeModuloInt(new BigNumber(((_228_v5).update(new BigNumber((_267_v36).length), new BigNumber((_268_v37).length))).length), _module.__default.fm0(_223_v0, new BigNumber((_dafny.Seq.of(_223_v0)).length), _238_globalState)));
      let _270_v39;
      _270_v39 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_module.__default.safeModuloInt(_225_v2, (_dafny.ZERO).minus(_225_v2))));
      _270_v39 = (_270_v39).Intersect(_270_v39);
      let _index27 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
      (_234_v11)[_index27] = _225_v2;
      let _index28 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
      (_234_v11)[_index28] = _225_v2;
      let _271_v40;
      let _nw29 = Array((new BigNumber(22)).toNumber()).fill(_dafny.MultiSet.Empty);
      _271_v40 = _nw29;
      let _272_v41;
      _272_v41 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lpbndwsy"), _module.__default.fm1(_223_v0, _238_globalState), _233_v10, _233_v10);
      let _index29 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_271_v40).length));
      (_271_v40)[_index29] = _272_v41;
      let _273_v42;
      _273_v42 = _module.D0.create_DC1(_225_v2);
      let _pat_let_tv6 = _272_v41;
      let _pat_let_tv7 = _233_v10;
      let _pat_let_tv8 = _230_v7;
      let _pat_let_tv9 = _272_v41;
      let _pat_let_tv10 = _272_v41;
      let _index30 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_271_v40).length));
      let _rhs19 = function (_source5) {
        if (_source5.is_DC1) {
          let _274___mcc_h0 = (_source5).cf1;
          let _275_cf1 = _274___mcc_h0;
          return (_pat_let_tv6).update(_pat_let_tv7, _module.__default.abs(new BigNumber((_pat_let_tv8).length)));
        } else if (_source5.is_DC2) {
          let _276___mcc_h1 = (_source5).cf2;
          let _277___mcc_h2 = (_source5).cf3;
          let _278___mcc_h3 = (_source5).cf4;
          let _279_cf4 = _278___mcc_h3;
          let _280_cf3 = _277___mcc_h2;
          let _281_cf2 = _276___mcc_h1;
          return _pat_let_tv9;
        } else if (_source5.is_DC0) {
          let _282___mcc_h4 = (_source5).cf0;
          let _283_cf0 = _282___mcc_h4;
          return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nt"));
        } else {
          let _284___mcc_h5 = (_source5).cf5;
          let _285_cf5 = _284___mcc_h5;
          return _pat_let_tv10;
        }
      }(_273_v42);
      let _rhs20 = ((_dafny.ZERO).minus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])).isEqualTo(_module.__default.safeModuloInt((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _225_v2));
      let _lhs13 = _271_v40;
      let _lhs14 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_271_v40).length));
      let _lhs15 = _238_globalState;
      _lhs13[_lhs14] = _rhs19;
      _lhs15.f13 = _rhs20;
      let _286_v43;
      _286_v43 = _dafny.Map.Empty.slice().updateUnsafe(_224_v1,(_233_v10)[_module.__default.safeIndex(_225_v2, new BigNumber((_233_v10).length))]);
      let _287_v44;
      let _init5 = ((_288_v13) => function (_289_i4) {
        return _288_v13;
      })(_236_v13);
      let _nw30 = Array((new BigNumber(5)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw30.length); _i0_5++) {
        _nw30[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _287_v44 = _nw30;
      let _290_v45;
      let _nw31 = new _module.C3();
      _nw31.__ctor(_286_v43, _287_v44, _module.__default.fm29((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _223_v0, _223_v0, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _238_globalState), !((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]).isEqualTo((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]));
      _290_v45 = _nw31;
      let _291_v46;
      _291_v46 = _module.D5.create_DC14(_223_v0);
      let _292_v47;
      _292_v47 = _dafny.Set.fromElements(_291_v46, _291_v46, _module.D5.create_DC14(_223_v0), function (_pat_let10_0) {
        return function (_293_dt__update__tmp_h1) {
          return function (_pat_let11_0) {
            return function (_294_dt__update_hcf23_h0) {
              return _module.D5.create_DC14(_294_dt__update_hcf23_h0);
            }(_pat_let11_0);
          }(true);
        }(_pat_let10_0);
      }(_module.D5.create_DC14(_223_v0)), _291_v46);
      let _295_v49;
      _295_v49 = _dafny.Map.Empty.slice().updateUnsafe(_292_v47,function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of _dafny.IntegerRange(new BigNumber(568), new BigNumber(81))) {
          let _296_v48 = _compr_47;
          if (((new BigNumber(568)).isLessThanOrEqualTo(_296_v48)) && ((_296_v48).isLessThan(new BigNumber(81)))) {
            _coll47.push([_module.__default.safeModuloInt(_296_v48, new BigNumber((_dafny.Seq.UnicodeFromString("vlhwkla")).length)),(((_270_v39).contains(new BigNumber(395))) ? ((_270_v39).get(new BigNumber(395))) : ((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]))]);
          }
        }
        return _coll47;
      }());
      let _297_v52;
      _297_v52 = _dafny.Seq.of(_295_v49, function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of (_dafny.Seq.of(_dafny.Set.fromElements(_291_v46), _292_v47, function () {
          let _coll49 = new _dafny.Set();
          for (const _compr_49 of (_module.__default.fm58(_236_v13, new BigNumber((_233_v10).length), _223_v0, _238_globalState)).Keys.Elements) {
            let _298_v51 = _compr_49;
            if ((_module.__default.fm58(_236_v13, new BigNumber((_233_v10).length), _223_v0, _238_globalState)).contains(_298_v51)) {
              _coll49.add(_298_v51);
            }
          }
          return _coll49;
        }())).Elements) {
          let _299_v50 = _compr_48;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(_291_v46), _292_v47, function () {
            let _coll50 = new _dafny.Set();
            for (const _compr_50 of (_module.__default.fm58(_236_v13, new BigNumber((_233_v10).length), _223_v0, _238_globalState)).Keys.Elements) {
              let _300_v51 = _compr_50;
              if ((_module.__default.fm58(_236_v13, new BigNumber((_233_v10).length), _223_v0, _238_globalState)).contains(_300_v51)) {
                _coll50.add(_300_v51);
              }
            }
            return _coll50;
          }()), _299_v50)) {
            _coll48.push([_299_v50,_228_v5]);
          }
        }
        return _coll48;
      }(), (_295_v49).Merge(_295_v49), _module.__default.fm59(_233_v10, _238_globalState), _295_v49);
      _295_v49 = (_297_v52)[_module.__default.safeIndex(new BigNumber(-874), new BigNumber((_297_v52).length))];
      let _301_v53;
      let _nw32 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
      _301_v53 = _nw32;
      let _index31 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_301_v53).length));
      (_301_v53)[_index31] = _dafny.Seq.of(_224_v1);
      let _302_v54;
      _302_v54 = _dafny.Seq.of(_224_v1, _224_v1, _224_v1, _224_v1, _224_v1);
      let _303_v55;
      _303_v55 = _dafny.Seq.of(_302_v54, _dafny.Seq.of(_224_v1));
      let _index32 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_301_v53).length));
      (_301_v53)[_index32] = (_303_v55)[_module.__default.safeIndex((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new BigNumber((_303_v55).length))];
      let _hi1 = _225_v2;
      for (let _304_i5 = _module.__default.safeDivisionInt(_225_v2, _225_v2); _304_i5.isLessThan(_hi1); _304_i5 = _304_i5.plus(_dafny.ONE)) {
        let _305_v56;
        let _init6 = ((_306_v0, _307_v2) => function (_308_i6) {
          return _dafny.Map.Empty.slice().updateUnsafe(_306_v0,_307_v2);
        })(_223_v0, _225_v2);
        let _nw33 = Array((new BigNumber(7)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw33.length); _i0_6++) {
          _nw33[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _305_v56 = _nw33;
        let _309_v57;
        let _nw34 = new _module.C1();
        _nw34.__ctor(true, false);
        _309_v57 = _nw34;
        let _310_v58;
        _310_v58 = _dafny.Map.Empty.slice().updateUnsafe(_305_v56,_309_v57);
        let _311_v59;
        _311_v59 = _dafny.Seq.of((((_310_v58).contains(_305_v56)) ? ((_310_v58).get(_305_v56)) : (_309_v57)), (((_309_v57).f39) ? (_309_v57) : (_309_v57)));
        let _rhs21 = _234_v11;
        let _rhs22 = _dafny.Seq.Concat(_311_v59, _311_v59);
        let _rhs23 = _224_v1;
        let _lhs16 = _238_globalState;
        _lhs16.f22 = _rhs21;
        _311_v59 = _rhs22;
        _224_v1 = _rhs23;
        (_238_globalState).f13 = (_230_v7)[_module.__default.safeIndex(_module.__default.fm0((_309_v57).f39, new BigNumber((_228_v5).length), _238_globalState), new BigNumber((_230_v7).length))];
        if ((new BigNumber(62)).isEqualTo((_304_i5).minus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]))) {
          let _312_v60;
          let _nw35 = new _module.C10();
          _nw35.__ctor((true) === ((_309_v57).f39), (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new _dafny.CodePoint('j'.codePointAt(0)), (_225_v2).isLessThanOrEqualTo(_225_v2), _287_v44, _223_v0);
          _312_v60 = _nw35;
          let _313_v61;
          _313_v61 = _dafny.Seq.of((((_309_v57).f39) ? (_312_v60.f29) : (_312_v60.f29)));
          let _rhs24 = _312_v60;
          let _rhs25 = _dafny.Seq.update(_dafny.Seq.Concat(((_223_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), ((_314_v60) => function (_315_i7) {
            return _314_v60.f29;
          })(_312_v60))) : (_313_v61)), _313_v61), _module.__default.safeIndex(new BigNumber((_233_v10).length), new BigNumber((_dafny.Seq.Concat(((_223_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), ((_316_v60) => function (_317_i7) {
            return _316_v60.f29;
          })(_312_v60))) : (_313_v61)), _313_v61)).length)), _304_i5);
          let _rhs26 = _dafny.Seq.Concat(_313_v61, _313_v61);
          let _rhs27 = _225_v2;
          let _rhs28 = _304_i5;
          let _lhs17 = _238_globalState;
          let _lhs18 = _238_globalState;
          _312_v60 = _rhs24;
          _313_v61 = _rhs25;
          _313_v61 = _rhs26;
          _lhs17.f7 = _rhs27;
          _lhs18.f7 = _rhs28;
          let _318_v62;
          _318_v62 = _dafny.MultiSet.fromElements(_223_v0, (_309_v57).f39, (_309_v57).f39);
          let _319_v63;
          let _nw36 = new _module.C2();
          _nw36.__ctor(_318_v62, _287_v44, _236_v13, (_312_v60).f28);
          _319_v63 = _nw36;
          (_238_globalState).f13 = _223_v0;
          let _320_v64;
          let _nw37 = new _module.C0();
          _nw37.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_312_v60).f28,_dafny.Seq.UnicodeFromString("w")), ((_309_v57).f39) && (_223_v0), new _dafny.CodePoint('l'.codePointAt(0)), (_312_v60).f28);
          _320_v64 = _nw37;
          let _index33 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          let _rhs29 = _320_v64;
          let _rhs30 = _318_v62;
          let _rhs31 = _module.__default.safeModuloInt(_225_v2, _312_v60.f29);
          let _rhs32 = _290_v45;
          let _rhs33 = ((_223_v0) ? (new BigNumber((_module.__default.fm3(!((_309_v57).f39), _238_globalState)).cardinality())) : ((_312_v60).fm9(_313_v61, _312_v60.f29, (_309_v57).f39, _238_globalState)));
          let _lhs19 = _238_globalState;
          let _lhs20 = _234_v11;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          _320_v64 = _rhs29;
          _318_v62 = _rhs30;
          _lhs19.f7 = _rhs31;
          _290_v45 = _rhs32;
          _lhs20[_lhs21] = _rhs33;
          _223_v0 = (_320_v64).f38;
        } else {
          let _321_v65;
          _321_v65 = _dafny.Map.Empty.slice().updateUnsafe((_309_v57).f39,_module.__default.fm1((_309_v57).f39, _238_globalState));
          let _322_v66;
          let _nw38 = new _module.C2();
          _nw38.__ctor(_dafny.MultiSet.FromArray(_230_v7), _287_v44, _236_v13, (_309_v57).f39);
          _322_v66 = _nw38;
          let _323_v67;
          _323_v67 = _dafny.Map.Empty.slice().updateUnsafe((((_321_v65).contains((_309_v57).f39)) ? ((_321_v65).get((_309_v57).f39)) : (_233_v10)),((_223_v0) ? (_322_v66) : (_322_v66)));
          _323_v67 = (_323_v67).update(_233_v10, _322_v66);
          let _324_v68;
          _324_v68 = _dafny.Seq.of(_304_i5, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
          let _325_v69;
          _325_v69 = _dafny.Map.Empty.slice().updateUnsafe((_322_v66).f24,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_228_v5).length),(_322_v66).f24));
          let _326_v71;
          _326_v71 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_327_i5) => function (_328_i8) {
            return _327_i5;
          })(_304_i5)), _module.__default.safeIndex(new BigNumber(((((_325_v69).contains((_309_v57).f39)) ? ((_325_v69).get((_309_v57).f39)) : (function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(-632), new BigNumber(635))) {
              let _329_v70 = _compr_51;
              if (((new BigNumber(-632)).isLessThanOrEqualTo(_329_v70)) && ((_329_v70).isLessThan(new BigNumber(635)))) {
                _coll51.push([(_329_v70).plus(new BigNumber((_227_v4).length)),(_module.D6.create_DC16(_231_v8, (_309_v57).f39, false)).dtor_cf27]);
              }
            }
            return _coll51;
          }()))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_330_i5) => function (_331_i8) {
            return _330_i5;
          })(_304_i5))).length)), new BigNumber((_233_v10).length)), _324_v68), _324_v68, _dafny.Seq.Create(_module.__default.abs(new BigNumber(930)), ((_332_v11) => function (_333_i9) {
            return (_332_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_332_v11).length))];
          })(_234_v11)), _324_v68, _324_v68);
          _324_v68 = (_326_v71)[_module.__default.safeIndex(((_309_v57).fm33(_238_globalState)).minus(_304_i5), new BigNumber((_326_v71).length))];
          let _334_v72;
          let _335_v73;
          let _336_v74;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = (_290_v45).m11(_223_v0, _238_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _334_v72 = _out9;
          _335_v73 = _out10;
          _336_v74 = _out11;
          let _337_v75;
          let _nw39 = new _module.C9();
          _nw39.__ctor(_236_v13, _336_v74);
          _337_v75 = _nw39;
          _337_v75 = _337_v75;
          _326_v71 = ((_334_v72) ? (_module.__default.fm60((_309_v57).f39, _238_globalState)) : (_326_v71));
        }
        _236_v13 = _236_v13;
      }
      let _338_v76;
      _338_v76 = _module.D21.create_DC51((_223_v0) || (_223_v0));
      let _source6 = _338_v76;
      if (_source6.is_DC51) {
        let _339___mcc_h6 = (_source6).cf81;
        let _340_cf81 = _339___mcc_h6;
        (_238_globalState).f13 = !(_223_v0);
        if ((new BigNumber(-409)).isLessThanOrEqualTo((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])) {
          let _341_v77;
          let _nw40 = new _module.C1();
          _nw40.__ctor(_223_v0, _340_cf81);
          _341_v77 = _nw40;
          let _342_v78;
          _342_v78 = _dafny.MultiSet.fromElements((_341_v77).fm32(_340_cf81, _225_v2, _238_globalState), (_230_v7)[_module.__default.safeIndex((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new BigNumber((_230_v7).length))]);
          let _343_v79;
          let _344_v80;
          let _345_v81;
          let _out12;
          let _out13;
          let _out14;
          let _outcollector4 = _module.__default.m0((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_341_v77,new _dafny.CodePoint('p'.codePointAt(0)))).length)).plus(new BigNumber((_270_v39).cardinality())), _342_v78, _223_v0, _225_v2, _238_globalState);
          _out12 = _outcollector4[0];
          _out13 = _outcollector4[1];
          _out14 = _outcollector4[2];
          _343_v79 = _out12;
          _344_v80 = _out13;
          _345_v81 = _out14;
          let _346_v82;
          let _nw41 = Array((new BigNumber(7)).toNumber()).fill([]);
          _346_v82 = _nw41;
          let _index34 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_346_v82).length));
          (_346_v82)[_index34] = _234_v11;
          let _index35 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_346_v82).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          let _rhs34 = (_227_v4).Merge(_227_v4);
          let _rhs35 = _344_v80;
          let _rhs36 = _234_v11;
          let _rhs37 = _module.__default.safeDivisionInt(_225_v2, _225_v2);
          let _rhs38 = _344_v80;
          let _lhs22 = _346_v82;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_346_v82).length));
          let _lhs24 = _234_v11;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          _227_v4 = _rhs34;
          _344_v80 = _rhs35;
          _lhs22[_lhs23] = _rhs36;
          _lhs24[_lhs25] = _rhs37;
          _223_v0 = _rhs38;
          let _347_v83;
          let _nw42 = new _module.C2();
          _nw42.__ctor(_342_v78, _287_v44, (((_341_v77).f39) ? (_236_v13) : (_236_v13)), _344_v80);
          _347_v83 = _nw42;
          _340_cf81 = _344_v80;
          let _348_v84;
          let _nw43 = new _module.C4();
          _nw43.__ctor((_347_v83).fm24((_347_v83).fm24(_345_v81, _340_cf81, _238_globalState), (_341_v77).f39, _238_globalState), _228_v5, _287_v44, _236_v13, (_343_v79).isLessThanOrEqualTo((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]));
          _348_v84 = _nw43;
          let _nw44 = new _module.C4();
          _nw44.__ctor(_343_v79, _348_v84.f34, _287_v44, _236_v13, _223_v0);
          _348_v84 = _nw44;
        } else {
          let _349_v85;
          let _350_v86;
          let _351_v87;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector5 = (_290_v45).m11(false, _238_globalState);
          _out15 = _outcollector5[0];
          _out16 = _outcollector5[1];
          _out17 = _outcollector5[2];
          _349_v85 = _out15;
          _350_v86 = _out16;
          _351_v87 = _out17;
          let _352_v88;
          let _nw45 = new _module.C5();
          _nw45.__ctor(_224_v1);
          _352_v88 = _nw45;
          let _353_v89;
          let _nw46 = Array((new BigNumber(12)).toNumber());
          _nw46[(_dafny.ZERO).toNumber()] = _352_v88;
          _nw46[(_dafny.ONE).toNumber()] = _352_v88;
          _nw46[(new BigNumber(2)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(3)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(4)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(5)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(6)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(7)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(8)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(9)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(10)).toNumber()] = _352_v88;
          _nw46[(new BigNumber(11)).toNumber()] = _352_v88;
          _353_v89 = _nw46;
          let _354_v90;
          _354_v90 = _module.D16.create_DC37(_352_v88);
          let _index37 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_353_v89).length));
          (_353_v89)[_index37] = (_354_v90).dtor_cf61;
          let _index38 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_353_v89).length));
          let _nw47 = new _module.C5();
          _nw47.__ctor(_224_v1);
          (_353_v89)[_index38] = _nw47;
          _233_v10 = _233_v10;
          let _355_v91;
          _355_v91 = _dafny.MultiSet.fromElements((new BigNumber(-238)).isEqualTo(_225_v2), _340_cf81);
          let _356_v94;
          let _nw48 = new _module.C4();
          _nw48.__ctor(_225_v2, _228_v5, _287_v44, _350_v86, _340_cf81);
          _356_v94 = _nw48;
          let _357_v95;
          _357_v95 = _module.D24.create_DC56(_356_v94);
          let _358_v96;
          _358_v96 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of _dafny.IntegerRange(new BigNumber(-204), new BigNumber(551))) {
              let _359_v93 = _compr_52;
              if (((new BigNumber(-204)).isLessThanOrEqualTo(_359_v93)) && ((_359_v93).isLessThan(new BigNumber(551)))) {
                _coll52.add((_359_v93).multipliedBy((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]));
              }
            }
            return _coll52;
          }(),(_357_v95).dtor_cf90);
          let _rhs39 = _module.__default.safeModuloInt(_225_v2, ((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]).minus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]));
          let _rhs40 = _module.__default.fm3(_351_v87, _238_globalState);
          let _rhs41 = (_358_v96).contains(function () {
            let _coll53 = new _dafny.Set();
            for (const _compr_53 of _dafny.IntegerRange(new BigNumber(609), new BigNumber(-929))) {
              let _360_v92 = _compr_53;
              if (((new BigNumber(609)).isLessThanOrEqualTo(_360_v92)) && ((_360_v92).isLessThan(new BigNumber(-929)))) {
                _coll53.add((_360_v92).minus((_dafny.ZERO).minus(_225_v2)));
              }
            }
            return _coll53;
          }());
          let _lhs26 = _238_globalState;
          _lhs26.f7 = _rhs39;
          _355_v91 = _rhs40;
          _340_cf81 = _rhs41;
          let _361_v97;
          _361_v97 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("b"),_234_v11);
          let _362_v98;
          let _nw49 = Array((new BigNumber(4)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = _234_v11;
          _nw49[(_dafny.ONE).toNumber()] = ((_351_v87) ? (_234_v11) : (_234_v11));
          _nw49[(new BigNumber(2)).toNumber()] = (((_361_v97).contains(_dafny.Seq.update(_233_v10, _module.__default.safeIndex(_module.__default.fm0(_351_v87, (_356_v94).f33, _238_globalState), new BigNumber((_233_v10).length)), _236_v13))) ? ((_361_v97).get(_dafny.Seq.update(_233_v10, _module.__default.safeIndex(_module.__default.fm0(_351_v87, (_356_v94).f33, _238_globalState), new BigNumber((_233_v10).length)), _236_v13))) : (_234_v11));
          _nw49[(new BigNumber(3)).toNumber()] = _234_v11;
          _362_v98 = _nw49;
          let _363_v99;
          _363_v99 = _module.D25.create_DC58(_362_v98);
          _362_v98 = (_363_v99).dtor_cf96;
        }
        let _364_v100;
        _364_v100 = _dafny.Set.fromElements(_287_v44, _287_v44);
        let _365_v101;
        _365_v101 = _dafny.Map.Empty.slice().updateUnsafe(_223_v0,_364_v100);
        if ((_module.__default.safeDivisionInt((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])).isLessThanOrEqualTo(new BigNumber(((((_365_v101).contains(_340_cf81)) ? ((_365_v101).get(_340_cf81)) : (_364_v100))).length))) {
          let _index39 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_224_v1).length));
          (_224_v1)[_index39] = _223_v0;
          let _index40 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_224_v1).length));
          let _index41 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          let _rhs42 = _236_v13;
          let _rhs43 = _223_v0;
          let _rhs44 = !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_366_v13) => function (_367_i10) {
            return _366_v13;
          })(_236_v13)), _233_v10);
          let _rhs45 = !(_340_cf81);
          let _rhs46 = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          let _lhs27 = _238_globalState;
          let _lhs28 = _224_v1;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_224_v1).length));
          let _lhs30 = _238_globalState;
          let _lhs31 = _234_v11;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          _236_v13 = _rhs42;
          _lhs27.f13 = _rhs43;
          _lhs28[_lhs29] = _rhs44;
          _lhs30.f13 = _rhs45;
          _lhs31[_lhs32] = _rhs46;
          let _368_v102;
          _368_v102 = _270_v39;
          let _369_v103;
          _369_v103 = _dafny.Seq.of((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
          let _nw50 = Array((new BigNumber(27)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = _225_v2;
          _nw50[(_dafny.ONE).toNumber()] = _225_v2;
          _nw50[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_225_v2, new BigNumber((_dafny.Seq.UnicodeFromString("fghnvwo")).length));
          _nw50[(new BigNumber(3)).toNumber()] = (_273_v42).dtor_cf1;
          _nw50[(new BigNumber(4)).toNumber()] = ((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]).multipliedBy((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
          _nw50[(new BigNumber(5)).toNumber()] = new BigNumber(946);
          _nw50[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(_225_v2, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
          _nw50[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_230_v7, _module.__default.safeIndex(new BigNumber(480), new BigNumber((_230_v7).length)), (_224_v1)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_224_v1).length))])).length));
          _nw50[(new BigNumber(8)).toNumber()] = (_225_v2).plus(_225_v2);
          _nw50[(new BigNumber(9)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(10)).toNumber()] = _225_v2;
          _nw50[(new BigNumber(11)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(12)).toNumber()] = _module.__default.fm0((_224_v1)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_224_v1).length))], _225_v2, _238_globalState);
          _nw50[(new BigNumber(13)).toNumber()] = new BigNumber((_module.__default.fm21(_267_v36, (_270_v39), _238_globalState)).length);
          _nw50[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(_225_v2, (((_228_v5).contains((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])) ? ((_228_v5).get((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])) : (_225_v2)));
          _nw50[(new BigNumber(15)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(16)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(17)).toNumber()] = new BigNumber(86);
          _nw50[(new BigNumber(18)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(19)).toNumber()] = _225_v2;
          _nw50[(new BigNumber(20)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(21)).toNumber()] = _225_v2;
          _nw50[(new BigNumber(22)).toNumber()] = _225_v2;
          _nw50[(new BigNumber(23)).toNumber()] = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          _nw50[(new BigNumber(24)).toNumber()] = _module.__default.fm0(_223_v0, new BigNumber((_227_v4).length), _238_globalState);
          _nw50[(new BigNumber(25)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), ((_370_v13) => function (_371_i11) {
            return _370_v13;
          })(_236_v13))).length), (_369_v103)[_module.__default.safeIndex((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new BigNumber((_369_v103).length))]);
          _nw50[(new BigNumber(26)).toNumber()] = _225_v2;
          _234_v11 = _nw50;
          (_238_globalState).f7 = _module.__default.safeModuloInt(((_369_v103)[_module.__default.safeIndex(_225_v2, new BigNumber((_369_v103).length))]).plus((_dafny.ZERO).minus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])), new BigNumber(-101));
          let _index42 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          (_234_v11)[_index42] = _module.__default.safeDivisionInt((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], (_dafny.ZERO).minus(((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]).minus(_225_v2)));
          (_238_globalState).f13 = _340_cf81;
        } else {
          let _372_v104;
          _372_v104 = _dafny.Seq.of(_module.__default.safeDivisionInt((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _225_v2));
          let _373_v105;
          _373_v105 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), ((_374_v2) => function (_375_i12) {
            return _374_v2;
          })(_225_v2)));
          let _376_v107;
          _376_v107 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), function (_377_i13) {
            return new BigNumber(672);
          }),_340_cf81);
          let _378_v109;
          _378_v109 = _dafny.Seq.of(_373_v105);
          let _379_v110;
          _379_v110 = _dafny.Map.Empty.slice().updateUnsafe(_372_v104,_236_v13);
          let _380_v112;
          _380_v112 = _module.D13.create_DC28(_340_cf81, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _225_v2);
          let _381_v114;
          _381_v114 = _dafny.Map.Empty.slice().updateUnsafe(_223_v0,_233_v10);
          let _382_v115;
          _382_v115 = _module.D2.create_DC6(_381_v114);
          let _383_v116;
          let _nw51 = Array((new BigNumber(22)).toNumber());
          _nw51[(_dafny.ZERO).toNumber()] = function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_373_v105).Elements) {
              let _384_v106 = _compr_54;
              if ((_373_v105).contains(_384_v106)) {
                _coll54.add(_384_v106);
              }
            }
            return _coll54;
          }();
          _nw51[(_dafny.ONE).toNumber()] = (_373_v105).Intersect(_373_v105);
          _nw51[(new BigNumber(2)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_372_v104);
          _nw51[(new BigNumber(4)).toNumber()] = ((_340_cf81) ? (_373_v105) : (_373_v105));
          _nw51[(new BigNumber(5)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(6)).toNumber()] = (function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (_376_v107).Keys.Elements) {
              let _385_v108 = _compr_55;
              if ((_376_v107).contains(_385_v108)) {
                _coll55.add(_385_v108);
              }
            }
            return _coll55;
          }()).Intersect(_373_v105);
          _nw51[(new BigNumber(7)).toNumber()] = (_378_v109)[_module.__default.safeIndex(_225_v2, new BigNumber((_378_v109).length))];
          _nw51[(new BigNumber(8)).toNumber()] = (_378_v109)[_module.__default.safeIndex(_225_v2, new BigNumber((_378_v109).length))];
          _nw51[(new BigNumber(9)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(10)).toNumber()] = ((_378_v109)[_module.__default.safeIndex((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new BigNumber((_378_v109).length))]).Intersect(_373_v105);
          _nw51[(new BigNumber(11)).toNumber()] = function () {
            let _coll56 = new _dafny.Set();
            for (const _compr_56 of (_379_v110).Keys.Elements) {
              let _386_v111 = _compr_56;
              if ((_379_v110).contains(_386_v111)) {
                _coll56.add(_386_v111);
              }
            }
            return _coll56;
          }();
          _nw51[(new BigNumber(12)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(13)).toNumber()] = (_373_v105).Difference(_373_v105);
          _nw51[(new BigNumber(14)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(15)).toNumber()] = (((_380_v112).dtor_cf45) ? (function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of (_dafny.Seq.of((_290_v45).fm7(_238_globalState))).Elements) {
              let _387_v113 = _compr_57;
              if (_dafny.Seq.contains(_dafny.Seq.of((_290_v45).fm7(_238_globalState)), _387_v113)) {
                _coll57.add(_387_v113);
              }
            }
            return _coll57;
          }()) : (_373_v105));
          _nw51[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements(_372_v104, _module.__default.fm45(_223_v0, _382_v115, (((_228_v5).contains(new BigNumber(591))) ? ((_228_v5).get(new BigNumber(591))) : (_225_v2)), _238_globalState));
          _nw51[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(_372_v104);
          _nw51[(new BigNumber(18)).toNumber()] = (_378_v109)[_module.__default.safeIndex(_225_v2, new BigNumber((_378_v109).length))];
          _nw51[(new BigNumber(19)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(20)).toNumber()] = _373_v105;
          _nw51[(new BigNumber(21)).toNumber()] = _373_v105;
          _383_v116 = _nw51;
          let _index43 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_383_v116).length));
          (_383_v116)[_index43] = _dafny.Set.fromElements(_372_v104);
          let _388_v117;
          let _nw52 = new _module.C10();
          _nw52.__ctor(_223_v0, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _236_v13, _223_v0, _287_v44, _340_cf81);
          _388_v117 = _nw52;
          let _389_v118;
          _389_v118 = _dafny.Map.Empty.slice().updateUnsafe(_340_cf81,_388_v117);
          let _390_v119;
          let _nw53 = Array((new BigNumber(18)).toNumber());
          _390_v119 = _nw53;
          let _391_v120;
          _391_v120 = _module.D18.create_DC44(_module.D18.create_DC41(_390_v119));
          let _392_v121;
          _392_v121 = _module.D11.create_DC23((_388_v117).f24, _225_v2, _231_v8, _291_v46);
          let _393_v122;
          _393_v122 = _module.D0.create_DC2(_340_cf81, _225_v2, _340_cf81);
          let _394_v123;
          _394_v123 = _module.D27.create_DC61(_389_v118);
          let _395_v124;
          _395_v124 = _dafny.Seq.of(((_340_cf81) ? ((_394_v123).dtor_cf99) : (_389_v118)), _389_v118);
          let _index44 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_383_v116).length));
          let _rhs47 = true;
          let _rhs48 = _dafny.Seq.Concat(_372_v104, _dafny.Seq.Concat(_372_v104, _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_391_v120)).cardinality())), _module.__default.safeIndex(_225_v2, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_391_v120)).cardinality()))).length)), _225_v2)));
          let _rhs49 = _230_v7;
          let _rhs50 = _module.__default.fm61((_392_v121).dtor_cf37, _module.__default.fm1(_223_v0, _238_globalState), ((!(false)) ? (_393_v122) : (_393_v122)), _225_v2, _238_globalState);
          let _rhs51 = (_395_v124)[_module.__default.safeIndex((_225_v2).plus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]), new BigNumber((_395_v124).length))];
          let _lhs33 = _238_globalState;
          let _lhs34 = _383_v116;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_383_v116).length));
          _lhs33.f13 = _rhs47;
          _372_v104 = _rhs48;
          _230_v7 = _rhs49;
          _lhs34[_lhs35] = _rhs50;
          _389_v118 = _rhs51;
          (_238_globalState).f13 = (_388_v117).f24;
          let _index45 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          (_234_v11)[_index45] = _225_v2;
          (_238_globalState).f13 = _223_v0;
          let _396_v125;
          let _nw54 = new _module.C4();
          _nw54.__ctor(_225_v2, _228_v5, _287_v44, _388_v117.f23, _340_cf81);
          _396_v125 = _nw54;
          let _397_v126;
          _397_v126 = _dafny.Map.Empty.slice().updateUnsafe(_396_v125,_340_cf81);
          let _398_v127;
          let _399_v128;
          let _400_v129;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector6 = (_290_v45).m3((((_397_v126).contains(_396_v125)) ? ((_397_v126).get(_396_v125)) : ((_396_v125).f24)), _236_v13, _238_globalState);
          _out18 = _outcollector6[0];
          _out19 = _outcollector6[1];
          _out20 = _outcollector6[2];
          _398_v127 = _out18;
          _399_v128 = _out19;
          _400_v129 = _out20;
        }
        _223_v0 = _223_v0;
      } else if (_source6.is_DC50) {
        let _401___mcc_h7 = (_source6).cf80;
        let _402_cf80 = _401___mcc_h7;
        let _403_v130;
        let _nw55 = new _module.C10();
        _nw55.__ctor(((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]).isLessThanOrEqualTo(new BigNumber(41)), _225_v2, _236_v13, false, _287_v44, _223_v0);
        _403_v130 = _nw55;
        let _404_v131;
        _404_v131 = _dafny.Seq.of(new BigNumber(685), new BigNumber(402), _225_v2);
        _404_v131 = _404_v131;
        let _index46 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
        (_234_v11)[_index46] = _403_v130.f29;
        _233_v10 = _dafny.Seq.UnicodeFromString("ueed");
      } else {
        let _405___mcc_h8 = (_source6).cf82;
        let _406_cf82 = _405___mcc_h8;
        let _407_v132;
        let _nw56 = Array((new BigNumber(29)).toNumber()).fill([]);
        _407_v132 = _nw56;
        let _index47 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_407_v132).length));
        (_407_v132)[_index47] = _234_v11;
        let _408_v133;
        _408_v133 = _dafny.Map.Empty.slice().updateUnsafe((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))],_234_v11);
        let _index48 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_407_v132).length));
        let _index49 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
        let _rhs52 = _234_v11;
        let _rhs53 = new BigNumber((_dafny.Seq.UnicodeFromString("xy")).length);
        let _rhs54 = _225_v2;
        let _rhs55 = _408_v133;
        let _lhs36 = _407_v132;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_407_v132).length));
        let _lhs38 = _234_v11;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
        _lhs36[_lhs37] = _rhs52;
        _225_v2 = _rhs53;
        _lhs38[_lhs39] = _rhs54;
        _408_v133 = _rhs55;
        let _409_v134;
        let _nw57 = new _module.C7();
        _nw57.__ctor(!(((_223_v0) ? (_223_v0) : (_223_v0))), false, _287_v44, _236_v13, _dafny.Seq.IsProperPrefixOf(_230_v7, _230_v7), (_290_v45).fm4(_225_v2, _230_v7, _238_globalState));
        _409_v134 = _nw57;
        _233_v10 = _dafny.Seq.UnicodeFromString("f");
        _233_v10 = _dafny.Seq.Concat(_dafny.Seq.update(_233_v10, _module.__default.safeIndex((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new BigNumber((_233_v10).length)), _236_v13), _233_v10);
      }
      let _rhs56 = (_dafny.ZERO).minus(_225_v2);
      let _rhs57 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kbtyeu"), _233_v10), _dafny.Seq.update(_module.__default.fm1(_223_v0, _238_globalState), _module.__default.safeIndex(_225_v2, new BigNumber((_module.__default.fm1(_223_v0, _238_globalState)).length)), _236_v13));
      let _lhs40 = _238_globalState;
      _lhs40.f7 = _rhs56;
      _223_v0 = _rhs57;
      let _index50 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
      (_224_v1)[_index50] = ((_223_v0) ? (_223_v0) : (_223_v0));
      let _410_v135;
      _410_v135 = _module.D18.create_DC42((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _223_v0);
      let _411_v136;
      _411_v136 = _dafny.Set.fromElements(_410_v135);
      let _412_v137;
      _412_v137 = _dafny.Map.Empty.slice().updateUnsafe(_410_v135,new BigNumber((_267_v36).length));
      let _index51 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
      (_224_v1)[_index51] = !(_411_v136).equals(function () {
        let _coll58 = new _dafny.Set();
        for (const _compr_58 of (_412_v137).Keys.Elements) {
          let _413_v138 = _compr_58;
          if ((_412_v137).contains(_413_v138)) {
            _coll58.add(_413_v138);
          }
        }
        return _coll58;
      }());
      let _ingredients0 = [];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_224_v1).length))) {
        let _414_i14 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_414_i14)) && ((_414_i14).isLessThan(new BigNumber((_224_v1).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_224_v1, (_414_i14).toNumber(), (_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))]));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _415_v139;
      _415_v139 = _module.D9.create_DC20(_233_v10, _223_v0, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
      let _416_v140;
      _416_v140 = _module.D4.create_DC10(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jnm"), _module.__default.safeIndex((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], new BigNumber((_dafny.Seq.UnicodeFromString("jnm")).length)), _236_v13), (_415_v139).dtor_cf31));
      let _source7 = _416_v140;
      if (_source7.is_DC11) {
        let _417___mcc_h9 = (_source7).cf18;
        let _418_cf18 = _417___mcc_h9;
        let _419_v141;
        _419_v141 = _dafny.Seq.of((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _418_cf18, (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
        let _420_v142;
        _420_v142 = _228_v5;
        let _421_v143;
        let _nw58 = new _module.C4();
        _nw58.__ctor(new BigNumber((_419_v141).length), _420_v142, _287_v44, _236_v13, (_290_v45).fm4((((_228_v5).contains(_225_v2)) ? ((_228_v5).get(_225_v2)) : ((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])), _230_v7, _238_globalState));
        _421_v143 = _nw58;
        let _422_v144;
        _422_v144 = _module.D24.create_DC56(_421_v143);
        let _index52 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
        let _index53 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
        let _rhs58 = (_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))];
        let _rhs59 = _422_v144;
        let _rhs60 = _223_v0;
        let _rhs61 = (_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))];
        let _lhs41 = _224_v1;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
        let _lhs43 = _224_v1;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
        let _lhs45 = _238_globalState;
        _lhs41[_lhs42] = _rhs58;
        _422_v144 = _rhs59;
        _lhs43[_lhs44] = _rhs60;
        _lhs45.f13 = _rhs61;
        (_238_globalState).f7 = (_dafny.ZERO).minus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
        let _423_v145;
        let _424_v146;
        let _425_v147;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector7 = (_290_v45).m11((((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))]) ? (_223_v0) : (_223_v0)), _238_globalState);
        _out21 = _outcollector7[0];
        _out22 = _outcollector7[1];
        _out23 = _outcollector7[2];
        _423_v145 = _out21;
        _424_v146 = _out22;
        _425_v147 = _out23;
        let _index54 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
        (_224_v1)[_index54] = !(_dafny.areEqual(_dafny.Seq.Concat(_230_v7, _dafny.Seq.of((_230_v7)[_module.__default.safeIndex(_418_cf18, new BigNumber((_230_v7).length))])), _dafny.Seq.Concat(_230_v7, _230_v7)));
      } else {
        let _426___mcc_h10 = (_source7).cf17;
        let _427_cf17 = _426___mcc_h10;
        let _428_v148;
        let _init7 = ((_429_v0, _430_v13) => function (_431_i15) {
          return _dafny.Map.Empty.slice().updateUnsafe(_429_v0,_430_v13);
        })(_223_v0, _236_v13);
        let _nw59 = Array((new BigNumber(5)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw59.length); _i0_7++) {
          _nw59[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _428_v148 = _nw59;
        let _432_v149;
        _432_v149 = _module.D14.create_DC30(_428_v148);
        let _433_v150;
        _433_v150 = _module.D14.create_DC32(_432_v149);
        let _434_v151;
        _434_v151 = _module.D14.create_DC32(((_223_v0) ? (_433_v150) : (_module.D14.create_DC31(_223_v0))));
        let _435_v152;
        _435_v152 = _dafny.Map.Empty.slice().updateUnsafe((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))],_225_v2);
        let _436_v153;
        let _nw60 = new _module.C4();
        _nw60.__ctor(_225_v2, _435_v152, _287_v44, _236_v13, true);
        _436_v153 = _nw60;
        let _437_v155;
        _437_v155 = _dafny.Map.Empty.slice().updateUnsafe(_225_v2,_410_v135);
        let _438_v156;
        _438_v156 = _dafny.Seq.of(_module.__default.fm0((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))], (_436_v153).f33, _238_globalState));
        let _rhs62 = (_290_v45).fm4((_225_v2).plus(new BigNumber((_dafny.MultiSet.fromElements(_436_v153, _436_v153)).cardinality())), _230_v7, _238_globalState);
        let _rhs63 = _434_v151;
        let _rhs64 = (_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))];
        let _rhs65 = new BigNumber((function () {
          let _coll59 = new _dafny.Map();
          for (const _compr_59 of ((_437_v155).update((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _410_v135)).Keys.Elements) {
            let _439_v154 = _compr_59;
            if (((_437_v155).update((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _410_v135)).contains(_439_v154)) {
              _coll59.push([_module.__default.safeModuloInt(_439_v154, (_436_v153).f33),_236_v13]);
            }
          }
          return _coll59;
        }()).length);
        let _rhs66 = _dafny.Seq.IsProperPrefixOf(((!((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))])) ? (_438_v156) : (_438_v156)), _dafny.Seq.update(_438_v156, _module.__default.safeIndex((_436_v153).f33, new BigNumber((_438_v156).length)), (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]));
        let _lhs46 = _238_globalState;
        let _lhs47 = _238_globalState;
        let _lhs48 = _238_globalState;
        let _lhs49 = _238_globalState;
        _lhs46.f13 = _rhs62;
        _434_v151 = _rhs63;
        _lhs47.f13 = _rhs64;
        _lhs48.f7 = _rhs65;
        _lhs49.f13 = _rhs66;
        if (!_dafny.Seq.contains(_dafny.Seq.Concat(_427_cf17, _233_v10), new _dafny.CodePoint('a'.codePointAt(0)))) {
          let _index55 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length));
          (_224_v1)[_index55] = (_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))];
          (_238_globalState).f7 = _module.__default.safeModuloInt((_225_v2).multipliedBy(new BigNumber((_dafny.Seq.update(_230_v7, _module.__default.safeIndex(new BigNumber((_427_cf17).length), new BigNumber((_230_v7).length)), _223_v0)).length)), (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]);
          let _440_v157;
          let _nw61 = new _module.C8();
          _nw61.__ctor(_223_v0);
          _440_v157 = _nw61;
          let _441_v158;
          _441_v158 = _440_v157;
          let _442_v159;
          _442_v159 = _dafny.Map.Empty.slice().updateUnsafe((_436_v153).f33,(_441_v158));
          _442_v159 = _442_v159;
          _225_v2 = _225_v2;
          let _443_v160;
          _443_v160 = _290_v45;
          _290_v45 = (_443_v160);
        } else {
          let _index56 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length));
          (_234_v11)[_index56] = ((_dafny.ZERO).minus((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])).multipliedBy(_225_v2);
          (_238_globalState).f7 = (_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))];
          let _444_v161;
          _444_v161 = _module.D14.create_DC30(_428_v148);
          _444_v161 = _444_v161;
          let _445_v162;
          let _446_v163;
          let _447_v164;
          let _out24;
          let _out25;
          let _out26;
          let _outcollector8 = (_436_v153).m3(((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))]).isLessThan(_225_v2), _236_v13, _238_globalState);
          _out24 = _outcollector8[0];
          _out25 = _outcollector8[1];
          _out26 = _outcollector8[2];
          _445_v162 = _out24;
          _446_v163 = _out25;
          _447_v164 = _out26;
          (_238_globalState).f13 = (_230_v7)[_module.__default.safeIndex(_225_v2, new BigNumber((_230_v7).length))];
        }
        let _448_v165;
        _448_v165 = _dafny.MultiSet.fromElements(_223_v0);
        let _449_v166;
        let _nw62 = new _module.C2();
        _nw62.__ctor(_448_v165, (((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))]) ? (_287_v44) : (_287_v44)), _236_v13, _223_v0);
        _449_v166 = _nw62;
        (_238_globalState).f7 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _225_v2), _dafny.Seq.of((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))])),!(_223_v0))).length);
      }
      let _450_v167;
      _450_v167 = _228_v5;
      let _source8 = _450_v167;
      let _451___mcc_h11 = _source8;
      let _452_cf16 = _451___mcc_h11;
      let _453_v168;
      let _454_v169;
      let _455_v170;
      let _out27;
      let _out28;
      let _out29;
      let _outcollector9 = _module.__default.m0((_234_v11)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_234_v11).length))], _dafny.MultiSet.fromElements((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))]), ((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))]) === (true), _225_v2, _238_globalState);
      _out27 = _outcollector9[0];
      _out28 = _outcollector9[1];
      _out29 = _outcollector9[2];
      _453_v168 = _out27;
      _454_v169 = _out28;
      _455_v170 = _out29;
      let _456_v171;
      _456_v171 = _dafny.MultiSet.fromElements(_234_v11, _234_v11);
      _456_v171 = _456_v171;
      (_238_globalState).f7 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_238_globalState),_223_v0)).length);
      let _457_v172;
      _457_v172 = _dafny.MultiSet.fromElements((_224_v1)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_224_v1).length))], _454_v169, _223_v0);
      let _458_v173;
      let _nw63 = new _module.C2();
      _nw63.__ctor(_457_v172, _287_v44, (((_290_v45).fm4(_455_v170, _230_v7, _238_globalState)) ? (_236_v13) : (new _dafny.CodePoint('a'.codePointAt(0)))), (false) === (!(_223_v0)));
      _458_v173 = _nw63;
      process.stdout.write(_dafny.toString(_223_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_225_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_226_v3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(298),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),new BigNumber(230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_230_v7, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_231_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_232_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_233_v10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v11)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_235_v12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_236_v13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-230),_dafny.Map.Empty.slice().updateUnsafe(true,true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_238_globalState).f2).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f4)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_238_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_238_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_238_globalState.f10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_238_globalState).f12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_238_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f15).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_238_globalState).f17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_238_globalState).f18).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f21)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f22)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v36).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v37).equals(_dafny.Set.fromElements(_dafny.ZERO, new BigNumber(230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v38).equals(_dafny.Set.fromElements(new BigNumber(462), new BigNumber(2), new BigNumber(230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v39).equals(_dafny.MultiSet.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_271_v40)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lpbndwsy"), _dafny.Seq.UnicodeFromString("cwwgiekbdyjkt"), _dafny.Seq.UnicodeFromString("wjoqku")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v41).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lpbndwsy"), _dafny.Seq.UnicodeFromString("cwwgiekbdyjkt"), _dafny.Seq.UnicodeFromString("wjoqku"), _dafny.Seq.UnicodeFromString("wjoqku")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_273_v42).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_286_v43).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v44)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v44)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v44)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v44)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v44)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_290_v45).f35).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_291_v46).dtor_cf23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_292_v47).equals(_dafny.Set.fromElements(_module.D5.create_DC14(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v49).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),_dafny.Map.Empty.slice()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_297_v52, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),_dafny.Map.Empty.slice()), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),new BigNumber(230))).updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true), _module.D5.create_DC14(false)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),new BigNumber(230))), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),_dafny.Map.Empty.slice()), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true), _module.D5.create_DC14(false)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(655),new BigNumber(-349))).updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(false)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-697),new BigNumber(-69))).updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-893),new BigNumber(6))), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_module.D5.create_DC14(true)),_dafny.Map.Empty.slice())))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_301_v53)[_dafny.ONE]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_302_v54).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_303_v55).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_338_v76).dtor_cf81));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_410_v135).dtor_cf68));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_410_v135).dtor_cf69));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_411_v136).equals(_dafny.Set.fromElements(_module.D18.create_DC42(_dafny.ZERO, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_412_v137).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D18.create_DC42(_dafny.ZERO, true),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_415_v139).dtor_cf31).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_415_v139).dtor_cf32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_415_v139).dtor_cf33));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_416_v140).dtor_cf17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_450_v167)).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),new BigNumber(230)))));
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
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D0(3);
      $dt.cf5 = cf5;
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
    get dtor_cf0() { return this.cf0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO);
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
    static create_DC5(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC4(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(_dafny.ZERO, _dafny.ZERO, _dafny.Seq.of(), [], _dafny.ZERO);
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
    static create_DC7(cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC8(cf15) {
      let $dt = new D2(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13 && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, false);
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
    static create_DC9(cf16) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf16) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC11(cf18) {
      let $dt = new D4(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC10(cf17) {
      let $dt = new D4(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + this.cf17.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.ZERO);
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
    static create_DC13(cf20, cf21, cf22) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf23) {
      let $dt = new D5(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D5(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(_dafny.ZERO, false, false);
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
    static create_DC16(cf25, cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC15(cf24) {
      let $dt = new D6(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(_dafny.Map.Empty, false, false);
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
    static create_DC17(cf28) {
      let $dt = new D7(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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
    static create_DC18(cf29) {
      let $dt = new D8(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf31, cf32, cf33) {
      let $dt = new D9(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC19(cf30) {
      let $dt = new D9(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + this.cf31.toVerbatimString(true) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC20(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC21(cf34) {
      let $dt = new D10(0);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf36, cf37, cf38, cf39) {
      let $dt = new D11(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC22(cf35) {
      let $dt = new D11(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC22" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf35 === other.cf35;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC23(false, _dafny.ZERO, _dafny.Map.Empty, _module.D5.Default());
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
    static create_DC25(cf41, cf42) {
      let $dt = new D12(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D12(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf43) {
      let $dt = new D12(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC25" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC25(false, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf45, cf46, cf47, cf48) {
      let $dt = new D13(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC29(cf49, cf50) {
      let $dt = new D13(1);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC27(cf44) {
      let $dt = new D13(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC28" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC27" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf49 === other.cf49 && this.cf50 === other.cf50;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC28(false, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31(cf52) {
      let $dt = new D14(0);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC30(cf51) {
      let $dt = new D14(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC32(cf53) {
      let $dt = new D14(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC31(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf55) {
      let $dt = new D15(0);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC35(cf56, cf57, cf58) {
      let $dt = new D15(1);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC36(cf59, cf60) {
      let $dt = new D15(2);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC33(cf54) {
      let $dt = new D15(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC35" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf55 === other.cf55;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf56, other.cf56) && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC34(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf62, cf63, cf64) {
      let $dt = new D16(0);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC37(cf61) {
      let $dt = new D16(1);
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC39(cf65) {
      let $dt = new D16(2);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC38" + "(" + this.cf62.toVerbatimString(true) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC37" + "(" + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC38(_dafny.Seq.UnicodeFromString(""), [], _dafny.Set.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC40(cf66) {
      let $dt = new D17(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf66 === other.cf66;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf68, cf69) {
      let $dt = new D18(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC43(cf70, cf71, cf72) {
      let $dt = new D18(1);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC41(cf67) {
      let $dt = new D18(2);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC44(cf73) {
      let $dt = new D18(3);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get is_DC44() { return this.$tag === 3; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC42" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + this.cf72.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC41" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 3) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71) && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf67 === other.cf67;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC42(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf75, cf76) {
      let $dt = new D19(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC47(cf77) {
      let $dt = new D19(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC45(cf74) {
      let $dt = new D19(2);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC48(cf78) {
      let $dt = new D19(3);
      $dt.cf78 = cf78;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get is_DC48() { return this.$tag === 3; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf78() { return this.cf78; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf75) + ", " + this.cf76.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC47" + "(" + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf78) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf77 === other.cf77;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf78, other.cf78);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC46(_dafny.Set.Empty, _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC49(cf79) {
      let $dt = new D20(0);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC51(cf81) {
      let $dt = new D21(0);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC50(cf80) {
      let $dt = new D21(1);
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC52(cf82) {
      let $dt = new D21(2);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC50" + "(" + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf81 === other.cf81;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC51(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf83) {
      let $dt = new D22(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC53" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf83 === other.cf83;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf85, cf86, cf87, cf88, cf89) {
      let $dt = new D23(0);
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC54(cf84) {
      let $dt = new D23(1);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC55" + "(" + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC54" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf85, other.cf85) && this.cf86 === other.cf86 && _dafny.areEqual(this.cf87, other.cf87) && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf84 === other.cf84;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC55(_module.D21.Default(), [], _dafny.MultiSet.Empty, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57(cf91, cf92, cf93, cf94, cf95) {
      let $dt = new D24(0);
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      return $dt;
    }
    static create_DC56(cf90) {
      let $dt = new D24(1);
      $dt.cf90 = cf90;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf90() { return this.cf90; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC57" + "(" + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC56" + "(" + _dafny.toString(this.cf90) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf91 === other.cf91 && this.cf92 === other.cf92 && _dafny.areEqual(this.cf93, other.cf93) && _dafny.areEqual(this.cf94, other.cf94) && this.cf95 === other.cf95;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf90 === other.cf90;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC57(null, false, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC59(cf97) {
      let $dt = new D25(0);
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC58(cf96) {
      let $dt = new D25(1);
      $dt.cf96 = cf96;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf96() { return this.cf96; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC59" + "(" + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC58" + "(" + _dafny.toString(this.cf96) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf97 === other.cf97;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf96 === other.cf96;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC59(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC60(cf98) {
      let $dt = new D26(0);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC60" + "(" + _dafny.toString(this.cf98) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf98, other.cf98);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62(cf100, cf101, cf102, cf103) {
      let $dt = new D27(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC63(cf104, cf105, cf106, cf107) {
      let $dt = new D27(1);
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC61(cf99) {
      let $dt = new D27(2);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC64(cf108) {
      let $dt = new D27(3);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get is_DC64() { return this.$tag === 3; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC62" + "(" + this.cf100.toVerbatimString(true) + ", " + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC63" + "(" + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 2) {
        return "D27.DC61" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 3) {
        return "D27.DC64" + "(" + _dafny.toString(this.cf108) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf100, other.cf100) && _dafny.areEqual(this.cf101, other.cf101) && this.cf102 === other.cf102 && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf104, other.cf104) && _dafny.areEqual(this.cf105, other.cf105) && _dafny.areEqual(this.cf106, other.cf106) && this.cf107 === other.cf107;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf99, other.cf99);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf108, other.cf108);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC62(_dafny.Seq.UnicodeFromString(""), new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC65(cf109) {
      let $dt = new D28(0);
      $dt.cf109 = cf109;
      return $dt;
    }
    get is_DC65() { return this.$tag === 0; }
    get dtor_cf109() { return this.cf109; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC65" + "(" + _dafny.toString(this.cf109) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf109 === other.cf109;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D28.Default();
        }
      };
    }
  }

  $module.D29 = class D29 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC66(cf110) {
      let $dt = new D29(0);
      $dt.cf110 = cf110;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get dtor_cf110() { return this.cf110; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC66" + "(" + _dafny.toString(this.cf110) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf110 === other.cf110;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D29.Default();
        }
      };
    }
  }

  $module.D30 = class D30 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC67(cf111) {
      let $dt = new D30(0);
      $dt.cf111 = cf111;
      return $dt;
    }
    get is_DC67() { return this.$tag === 0; }
    get dtor_cf111() { return this.cf111; }
    toString() {
      if (this.$tag === 0) {
        return "D30.DC67" + "(" + _dafny.toString(this.cf111) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf111, other.cf111);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D30.Default();
        }
      };
    }
  }

  $module.D31 = class D31 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC69(cf113) {
      let $dt = new D31(0);
      $dt.cf113 = cf113;
      return $dt;
    }
    static create_DC70(cf114) {
      let $dt = new D31(1);
      $dt.cf114 = cf114;
      return $dt;
    }
    static create_DC68(cf112) {
      let $dt = new D31(2);
      $dt.cf112 = cf112;
      return $dt;
    }
    get is_DC69() { return this.$tag === 0; }
    get is_DC70() { return this.$tag === 1; }
    get is_DC68() { return this.$tag === 2; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf112() { return this.cf112; }
    toString() {
      if (this.$tag === 0) {
        return "D31.DC69" + "(" + _dafny.toString(this.cf113) + ")";
      } else if (this.$tag === 1) {
        return "D31.DC70" + "(" + _dafny.toString(this.cf114) + ")";
      } else if (this.$tag === 2) {
        return "D31.DC68" + "(" + _dafny.toString(this.cf112) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf113, other.cf113);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf114, other.cf114);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf112, other.cf112);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D31.create_DC69(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D31.Default();
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

  $module.T3 = class T3 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f7 = _dafny.ZERO;
      this.f10 = _dafny.Set.Empty;
      this.f13 = false;
      this.f22 = [];
      this._f0 = [];
      this._f1 = _dafny.ZERO;
      this._f2 = _dafny.Map.Empty;
      this._f3 = false;
      this._f4 = [];
      this._f5 = false;
      this._f6 = [];
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = false;
      this._f11 = _dafny.ZERO;
      this._f12 = _dafny.Seq.UnicodeFromString("");
      this._f14 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f15 = _dafny.Map.Empty;
      this._f16 = _dafny.ZERO;
      this._f17 = _dafny.Seq.UnicodeFromString("");
      this._f18 = _dafny.Seq.UnicodeFromString("");
      this._f19 = false;
      this._f20 = _dafny.ZERO;
      this._f21 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this).f22 = f22;
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
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
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
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
      this.f37 = _dafny.Map.Empty;
      this._f38 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f37, f38, f23, f24) {
      let _this = this;
      (_this).f37 = f37;
      (_this)._f38 = f38;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    fm4(p0, p1, globalState) {
      let _this = this;
      return !(!(((_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f38,(_module.D0.create_DC2((_this).f24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(624),true)).length), (_this).f24)).dtor_cf4)).length)).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f24)).cardinality()))))).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f24))).length)))));
    };
    fm26(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_this).f24, (_this).f24)).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f38)).length))).length)), (new BigNumber(403)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bdnuosq")).length))));
    };
    get f38() {
      let _this = this;
      return _this._f38;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f27 = false;
      this._f39 = false;
    }
    _parentTraits() {
      return [_module.T3];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f39, f27) {
      let _this = this;
      (_this)._f39 = f39;
      (_this)._f27 = f27;
      return;
    }
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("snfxoe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), function (_459_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-692)), function (_460_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }))).cardinality())).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements((_this).f39, (_this).f27, (_this).f27)).cardinality()))), new BigNumber(934));
    };
    fm32(p0, p1, globalState) {
      let _this = this;
      return !(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("afadxteqa"),new BigNumber(915))).length),(_this).f39)).length)), _dafny.Seq.of(new BigNumber(155), new BigNumber((_dafny.Seq.of((_this).f39)).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(590), new BigNumber((_dafny.Seq.of((_this).f39)).length)), _dafny.Seq.of(new BigNumber(671), new BigNumber((_dafny.Seq.UnicodeFromString("yahpsabt")).length)))));
    };
    fm33(globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(506), new BigNumber(((function () {
        let _coll60 = new _dafny.Set();
        for (const _compr_60 of (_dafny.Seq.of(_module.D5.create_DC14((_this).f39))).Elements) {
          let _461_v0 = _compr_60;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D5.create_DC14((_this).f39)), _461_v0)) {
            _coll60.add(_461_v0);
          }
        }
        return _coll60;
      }()).Union(_dafny.Set.fromElements(_module.D5.create_DC14((_this).f39), _module.D5.create_DC14((_this).f27)))).length));
    };
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _462_i0;
      _462_i0 = _dafny.ZERO;
      L0: {
        while (true) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_462_i0)) {
              break L0;
            }
            _462_i0 = (_462_i0).plus(_dafny.ONE);
            let _463_v0;
            _463_v0 = new _dafny.CodePoint('h'.codePointAt(0));
            _463_v0 = _463_v0;
            if ((_this).f39) {
              let _464_v1;
              _464_v1 = new BigNumber(852);
              let _465_v2;
              _465_v2 = _dafny.Seq.of(_464_v1, _464_v1, _464_v1);
              let _466_v3;
              _466_v3 = _dafny.MultiSet.fromElements(_464_v1, _464_v1, new BigNumber((_465_v2).length));
              let _467_v4;
              _467_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f39,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("dr"), _module.__default.safeIndex(_464_v1, new BigNumber((_dafny.Seq.UnicodeFromString("dr")).length)), new _dafny.CodePoint('h'.codePointAt(0))));
              let _468_v5;
              _468_v5 = _module.D2.create_DC6(_467_v4);
              let _469_v6;
              _469_v6 = _dafny.Map.Empty.slice().updateUnsafe(_468_v5,p0);
              let _470_v7;
              _470_v7 = _dafny.Map.Empty.slice().updateUnsafe(_464_v1,true);
              let _471_v8;
              _471_v8 = _dafny.Seq.UnicodeFromString("nubq");
              let _472_v9;
              _472_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f39,p0);
              let _473_v10;
              let _nw64 = Array((new BigNumber(18)).toNumber());
              _nw64[(_dafny.ZERO).toNumber()] = _464_v1;
              _nw64[(_dafny.ONE).toNumber()] = _464_v1;
              _nw64[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_465_v2).length), (_this).fm33(globalState));
              _nw64[(new BigNumber(3)).toNumber()] = new BigNumber((_466_v3).cardinality());
              _nw64[(new BigNumber(4)).toNumber()] = new BigNumber((_469_v6).length);
              _nw64[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((((_470_v7).contains(new BigNumber(483))) ? ((_470_v7).get(new BigNumber(483))) : (p0)))).cardinality()));
              _nw64[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("jovube")).length);
              _nw64[(new BigNumber(7)).toNumber()] = (_this).fm9(_465_v2, _464_v1, (_this).f27, globalState);
              _nw64[(new BigNumber(8)).toNumber()] = (_this).fm33(globalState);
              _nw64[(new BigNumber(9)).toNumber()] = _464_v1;
              _nw64[(new BigNumber(10)).toNumber()] = _464_v1;
              _nw64[(new BigNumber(11)).toNumber()] = _464_v1;
              _nw64[(new BigNumber(12)).toNumber()] = _464_v1;
              _nw64[(new BigNumber(13)).toNumber()] = new BigNumber(-64);
              _nw64[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_471_v8, _471_v8)).length);
              _nw64[(new BigNumber(15)).toNumber()] = (((_this).f39) ? (_464_v1) : (_464_v1));
              _nw64[(new BigNumber(16)).toNumber()] = new BigNumber((_466_v3).cardinality());
              _nw64[(new BigNumber(17)).toNumber()] = new BigNumber((_472_v9).length);
              _473_v10 = _nw64;
              (globalState).f22 = _473_v10;
              _463_v0 = _module.__default.fm34((_464_v1).minus(_464_v1), p0, globalState);
              (globalState).f7 = _464_v1;
              let _474_v11;
              _474_v11 = _dafny.Set.fromElements(_464_v1);
              let _475_v12;
              _475_v12 = _dafny.Seq.of(_474_v11);
              let _476_v13;
              _476_v13 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_464_v1), (_475_v12)[_module.__default.safeIndex(_464_v1, new BigNumber((_475_v12).length))], _dafny.Set.fromElements(_464_v1), _474_v11, _module.__default.fm35((_this).f27, globalState));
              _465_v2 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("hrwmpm")).length), _464_v1), _465_v2), _dafny.Seq.Concat(_dafny.Seq.of((_this).fm9(_465_v2, _464_v1, (_this).f27, globalState), _464_v1, new BigNumber((_470_v7).length)), _dafny.Seq.of(new BigNumber((_476_v13).cardinality()), _464_v1))), _module.__default.safeIndex((_464_v1).minus((_this).fm9(_465_v2, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ulr"), _module.__default.safeIndex(_464_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ulr")).length)), _463_v0)).length), p0, globalState)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("hrwmpm")).length), _464_v1), _465_v2), _dafny.Seq.Concat(_dafny.Seq.of((_this).fm9(_465_v2, _464_v1, (_this).f27, globalState), _464_v1, new BigNumber((_470_v7).length)), _dafny.Seq.of(new BigNumber((_476_v13).cardinality()), _464_v1)))).length)), new BigNumber((_465_v2).length));
              _464_v1 = (_dafny.ZERO).minus(_464_v1);
            } else {
              let _477_v14;
              _477_v14 = new BigNumber(699);
              let _478_v15;
              _478_v15 = _dafny.Seq.of(_477_v14);
              let _479_v16;
              _479_v16 = _dafny.Seq.of(_478_v15, _478_v15);
              let _480_v17;
              let _init8 = function (_481_i1) {
                return (_this).f39;
              };
              let _nw65 = Array((new BigNumber(19)).toNumber());
              for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw65.length); _i0_8++) {
                _nw65[_i0_8] = _init8(new BigNumber(_i0_8));
              }
              _480_v17 = _nw65;
              let _482_v18;
              _482_v18 = _module.D1.create_DC5(_477_v14, _477_v14, _479_v16, _480_v17, _477_v14);
              let _rhs67 = (_482_v18).dtor_cf8;
              let _rhs68 = !((_this).f27) || (p0);
              let _rhs69 = (_this).f39;
              let _lhs50 = globalState;
              _lhs50.f7 = _rhs67;
              r0 = _rhs68;
              r0 = _rhs69;
              (globalState).f13 = p0;
              let _483_v19;
              _483_v19 = _dafny.Seq.of(p0);
              _477_v14 = new BigNumber((_483_v19).length);
              (globalState).f13 = (_this).f39;
              let _484_v20;
              let _nw66 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
              _484_v20 = _nw66;
              let _485_v21;
              let _nw67 = Array((new BigNumber(9)).toNumber());
              _nw67[(_dafny.ZERO).toNumber()] = _484_v20;
              _nw67[(_dafny.ONE).toNumber()] = _484_v20;
              _nw67[(new BigNumber(2)).toNumber()] = _484_v20;
              _nw67[(new BigNumber(3)).toNumber()] = _484_v20;
              _nw67[(new BigNumber(4)).toNumber()] = _484_v20;
              _nw67[(new BigNumber(5)).toNumber()] = _484_v20;
              _nw67[(new BigNumber(6)).toNumber()] = _484_v20;
              _nw67[(new BigNumber(7)).toNumber()] = _484_v20;
              _nw67[(new BigNumber(8)).toNumber()] = _484_v20;
              _485_v21 = _nw67;
              let _index57 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_485_v21).length));
              (_485_v21)[_index57] = _484_v20;
              let _index58 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_485_v21).length));
              (_485_v21)[_index58] = _484_v20;
            }
            let _486_v22;
            _486_v22 = new BigNumber(683);
            (globalState).f7 = (_dafny.ZERO).minus(_486_v22);
            let _487_v23;
            _487_v23 = _dafny.Map.Empty.slice().updateUnsafe(_486_v22,_module.__default.fm36(globalState));
            let _488_v24;
            let _init9 = ((_489_p0) => function (_490_i2) {
              return _dafny.Map.Empty.slice().updateUnsafe(_489_p0,_489_p0);
            })(p0);
            let _nw68 = Array((new BigNumber(27)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw68.length); _i0_9++) {
              _nw68[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _488_v24 = _nw68;
            let _491_v25;
            _491_v25 = _dafny.Set.fromElements(_486_v22);
            let _492_v26;
            _492_v26 = _dafny.Seq.of(false);
            let _493_v27;
            _493_v27 = _dafny.Map.Empty.slice().updateUnsafe(_488_v24,(_487_v23).Merge(((_module.D6.create_DC15(_487_v23)).dtor_cf24).update(new BigNumber((_491_v25).length), _492_v26)));
            _487_v23 = (((_493_v27).contains(_488_v24)) ? ((_493_v27).get(_488_v24)) : (_module.__default.fm37(_486_v22, p0, globalState)));
          }
        }
      }
      let _494_v28;
      _494_v28 = _dafny.MultiSet.fromElements((_this).f27);
      let _495_v29;
      _495_v29 = new BigNumber(-13);
      let _496_v30;
      _496_v30 = _dafny.Set.fromElements(new BigNumber(-220), _495_v29);
      let _497_v31;
      _497_v31 = _dafny.Map.Empty.slice().updateUnsafe(_494_v28,(_496_v30).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_496_v30).length), (((_494_v28).contains((_this).f27)) ? ((_494_v28).get((_this).f27)) : (_495_v29)))));
      _497_v31 = _497_v31;
      let _498_v32;
      _498_v32 = _dafny.Seq.UnicodeFromString("pkn");
      let _499_v33;
      let _nw69 = new _module.C0();
      _nw69.__ctor((_dafny.Map.Empty.slice().updateUnsafe((_this).f39,_498_v32)).update(p0, _498_v32), p0, new _dafny.CodePoint('m'.codePointAt(0)), p0);
      _499_v33 = _nw69;
      (globalState).f13 = p0;
      let _500_v34;
      _500_v34 = _module.D5.create_DC14((_this).f27);
      _500_v34 = _module.D5.create_DC14((_this).f39);
      let _501_v35;
      _501_v35 = new _dafny.CodePoint('u'.codePointAt(0));
      let _502_v36;
      _502_v36 = _dafny.Map.Empty.slice().updateUnsafe(_501_v35,(_500_v34).dtor_cf23);
      _502_v36 = ((_502_v36).Merge(_502_v36)).Merge(_502_v36);
      r0 = (_this).f27;
      return r0;
    }
    get f39() {
      let _this = this;
      return _this._f39;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
      this._f25 = [];
      this._f36 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f36, f25, f23, f24) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f25 = f25;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_this).f24)).length), new BigNumber((_dafny.Seq.of((_this).f24)).length)), new BigNumber(((_dafny.Set.fromElements(!((_this).f24), (_this).f24)).Union(_dafny.Set.fromElements((_this).f24))).length), (new BigNumber(998)).plus(new BigNumber(870)));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("ptgbrq"))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("sxchetkux")));
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_this).f24;
    };
    fm24(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)))).Union(function () {
        let _coll61 = new _dafny.Set();
        for (const _compr_61 of _dafny.IntegerRange(new BigNumber(584), new BigNumber(-980))) {
          let _503_v0 = _compr_61;
          if (((new BigNumber(584)).isLessThanOrEqualTo(_503_v0)) && ((_503_v0).isLessThan(new BigNumber(-980)))) {
            _coll61.add(_module.__default.safeModuloInt(_503_v0, new BigNumber(982)));
          }
        }
        return _coll61;
      }())).Union((_module.D5.create_DC12(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.UnicodeFromString("jayht")).length))).length)))).dtor_cf19)).length));
    };
    fm25(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of((_this).f24)), _dafny.Seq.Concat(_dafny.Seq.of((_this).f24), _dafny.Seq.of(false)));
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _504_v0;
      _504_v0 = new BigNumber(268);
      let _505_v1;
      _505_v1 = _dafny.Set.fromElements(_504_v0, _504_v0);
      let _506_v2;
      _506_v2 = _module.D5.create_DC12(_505_v1);
      let _source9 = _506_v2;
      if (_source9.is_DC13) {
        let _507___mcc_h0 = (_source9).cf20;
        let _508___mcc_h1 = (_source9).cf21;
        let _509___mcc_h2 = (_source9).cf22;
        let _510_cf22 = _509___mcc_h2;
        let _511_cf21 = _508___mcc_h1;
        let _512_cf20 = _507___mcc_h0;
        let _index59 = _module.__default.safeIndex(new BigNumber(332), new BigNumber(((_this).f25).length));
        ((_this).f25)[_index59] = p1;
        let _index60 = _module.__default.safeIndex(new BigNumber(332), new BigNumber(((_this).f25).length));
        ((_this).f25)[_index60] = new _dafny.CodePoint('d'.codePointAt(0));
        let _513_v3;
        _513_v3 = _module.D0.create_DC1(_504_v0);
        let _514_v4;
        _514_v4 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_dafny.Seq.of(_513_v3), _513_v3),_504_v0);
        let _515_v5;
        _515_v5 = _dafny.Seq.of(p0, (_this).f24);
        _514_v4 = (_514_v4).update((((_515_v5)[_module.__default.safeIndex(_module.__default.fm0((_this).f24, _504_v0, globalState), new BigNumber((_515_v5).length))]) ? (_511_cf21) : (_510_cf22)), (_dafny.ZERO).minus(new BigNumber((_515_v5).length)));
        let _516_v6;
        _516_v6 = _dafny.Map.Empty.slice().updateUnsafe((_512_cf20).multipliedBy(_504_v0),p0);
        _516_v6 = _dafny.Map.Empty.slice().updateUnsafe(_504_v0,(_510_cf22) === ((_this).f24));
        let _517_v7;
        let _nw70 = Array((new BigNumber(22)).toNumber()).fill(false);
        _517_v7 = _nw70;
        let _index61 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_517_v7).length));
        (_517_v7)[_index61] = false;
        let _index62 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_517_v7).length));
        (_517_v7)[_index62] = p0;
      } else if (_source9.is_DC14) {
        let _518___mcc_h3 = (_source9).cf23;
        let _519_cf23 = _518___mcc_h3;
        (globalState).f7 = (_dafny.ZERO).minus(((_519_cf23) ? (_504_v0) : (_504_v0)));
        let _520_v8;
        _520_v8 = _dafny.Seq.of((_this).f24);
        let _521_v9;
        let _nw71 = Array((new BigNumber(12)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = p0;
        _nw71[(_dafny.ONE).toNumber()] = p0;
        _nw71[(new BigNumber(2)).toNumber()] = _519_cf23;
        _nw71[(new BigNumber(3)).toNumber()] = p0;
        _nw71[(new BigNumber(4)).toNumber()] = _519_cf23;
        _nw71[(new BigNumber(5)).toNumber()] = p0;
        _nw71[(new BigNumber(6)).toNumber()] = !(_504_v0).isEqualTo((_dafny.ZERO).minus(_504_v0));
        _nw71[(new BigNumber(7)).toNumber()] = (_this).f24;
        _nw71[(new BigNumber(8)).toNumber()] = _module.__default.fm2(globalState);
        _nw71[(new BigNumber(9)).toNumber()] = ((true) ? ((_this).fm4(_504_v0, _dafny.Seq.update(_520_v8, _module.__default.safeIndex(new BigNumber(-415), new BigNumber((_520_v8).length)), p0), globalState)) : (p0));
        _nw71[(new BigNumber(10)).toNumber()] = p0;
        _nw71[(new BigNumber(11)).toNumber()] = ((_519_cf23) ? ((_this).f24) : (p0));
        _521_v9 = _nw71;
        let _index63 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_521_v9).length));
        (_521_v9)[_index63] = true;
        let _index64 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_521_v9).length));
        (_521_v9)[_index64] = (_this).f24;
        let _522_v10;
        _522_v10 = _dafny.Seq.UnicodeFromString("yf");
        let _523_v11;
        _523_v11 = _dafny.Seq.of(_504_v0);
        let _524_v12;
        _524_v12 = _dafny.Seq.of(_523_v11);
        let _525_v13;
        _525_v13 = _module.D1.create_DC5(new BigNumber((_522_v10).length), _504_v0, _524_v12, _521_v9, _504_v0);
        let _pat_let_tv11 = _504_v0;
        let _pat_let_tv12 = _521_v9;
        let _pat_let_tv13 = _524_v12;
        let _pat_let_tv14 = _504_v0;
        let _pat_let_tv15 = _523_v11;
        let _pat_let_tv16 = _504_v0;
        let _pat_let_tv17 = _523_v11;
        _521_v9 = (function (_pat_let12_0) {
          return function (_530_dt__update__tmp_h1) {
            return function (_pat_let17_0) {
              return function (_531_dt__update_hcf7_h1) {
                return function (_pat_let18_0) {
                  return function (_534_dt__update_hcf8_h0) {
                    return _module.D1.create_DC5(_531_dt__update_hcf7_h1, _534_dt__update_hcf8_h0, (_530_dt__update__tmp_h1).dtor_cf9, (_530_dt__update__tmp_h1).dtor_cf10, (_530_dt__update__tmp_h1).dtor_cf11);
                  }(_pat_let18_0);
                }((_pat_let_tv15)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-263)), ((_532_v0) => function (_533_i0) {
                  return _532_v0;
                })(_pat_let_tv16))).length), new BigNumber((_pat_let_tv17).length))]);
              }(_pat_let17_0);
            }(_pat_let_tv14);
          }(_pat_let12_0);
        }(function (_pat_let13_0) {
          return function (_526_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_527_dt__update_hcf7_h0) {
                return function (_pat_let15_0) {
                  return function (_528_dt__update_hcf10_h0) {
                    return function (_pat_let16_0) {
                      return function (_529_dt__update_hcf9_h0) {
                        return _module.D1.create_DC5(_527_dt__update_hcf7_h0, (_526_dt__update__tmp_h0).dtor_cf8, _529_dt__update_hcf9_h0, _528_dt__update_hcf10_h0, (_526_dt__update__tmp_h0).dtor_cf11);
                      }(_pat_let16_0);
                    }(_pat_let_tv13);
                  }(_pat_let15_0);
                }(_pat_let_tv12);
              }(_pat_let14_0);
            }(_pat_let_tv11);
          }(_pat_let13_0);
        }(_525_v13))).dtor_cf10;
        _522_v10 = _dafny.Seq.update(_522_v10, _module.__default.safeIndex(_504_v0, new BigNumber((_522_v10).length)), _this.f23);
      } else {
        let _535___mcc_h4 = (_source9).cf19;
        let _536_cf19 = _535___mcc_h4;
        (globalState).f13 = (_this).f24;
        let _537_v14;
        _537_v14 = _dafny.Seq.of((_this).f24, !(p0));
        let _538_v15;
        let _nw72 = Array((new BigNumber(21)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = p0;
        _nw72[(_dafny.ONE).toNumber()] = p0;
        _nw72[(new BigNumber(2)).toNumber()] = _module.__default.fm2(globalState);
        _nw72[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(4)).toNumber()] = p0;
        _nw72[(new BigNumber(5)).toNumber()] = p0;
        _nw72[(new BigNumber(6)).toNumber()] = p0;
        _nw72[(new BigNumber(7)).toNumber()] = (_537_v14)[_module.__default.safeIndex(_504_v0, new BigNumber((_537_v14).length))];
        _nw72[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(9)).toNumber()] = p0;
        _nw72[(new BigNumber(10)).toNumber()] = !((_this).f24);
        _nw72[(new BigNumber(11)).toNumber()] = false;
        _nw72[(new BigNumber(12)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(13)).toNumber()] = false;
        _nw72[(new BigNumber(14)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(15)).toNumber()] = p0;
        _nw72[(new BigNumber(16)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(17)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(18)).toNumber()] = !(p0);
        _nw72[(new BigNumber(19)).toNumber()] = (_this).f24;
        _nw72[(new BigNumber(20)).toNumber()] = !(false);
        _538_v15 = _nw72;
        let _539_v16;
        let _nw73 = Array((_dafny.ONE).toNumber());
        _nw73[(_dafny.ZERO).toNumber()] = _538_v15;
        _539_v16 = _nw73;
        let _index65 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length));
        (_539_v16)[_index65] = _538_v15;
        let _index66 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length));
        (_539_v16)[_index66] = _538_v15;
        let _540_v17;
        let _nw74 = new _module.C0();
        _nw74.__ctor(_dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm1(p0, globalState)), false, p1, false);
        _540_v17 = _nw74;
        let _541_v18;
        _541_v18 = _dafny.Seq.UnicodeFromString("t");
        let _arr0 = (_539_v16)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length))];
        let _index67 = _module.__default.safeIndex(new BigNumber(146), new BigNumber(((_539_v16)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length))]).length));
        _arr0[_index67] = (p0) && (p0);
        let _542_v19;
        _542_v19 = _dafny.Set.fromElements(_dafny.Seq.update(_537_v14, _module.__default.safeIndex(_504_v0, new BigNumber((_537_v14).length)), (_this).f24));
        let _arr1 = (_539_v16)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length))];
        let _index68 = _module.__default.safeIndex(new BigNumber(146), new BigNumber(((_539_v16)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length))]).length));
        let _rhs70 = _dafny.Seq.UnicodeFromString("eeheeq");
        let _rhs71 = !(_dafny.Set.fromElements(_537_v14)).equals(_542_v19);
        let _lhs51 = (_539_v16)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length))];
        let _lhs52 = _module.__default.safeIndex(new BigNumber(146), new BigNumber(((_539_v16)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_539_v16).length))]).length));
        _541_v18 = _rhs70;
        _lhs51[_lhs52] = _rhs71;
      }
      let _543_v20;
      _543_v20 = _dafny.Seq.UnicodeFromString("mcs");
      let _544_v21;
      _544_v21 = _dafny.Map.Empty.slice().updateUnsafe(_504_v0,false);
      let _545_v22;
      _545_v22 = _dafny.Seq.of(new BigNumber((_543_v20).length), new BigNumber((_544_v21).length));
      let _546_v23;
      _546_v23 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-728)), ((_547_v0) => function (_548_i2) {
        return (_module.D4.create_DC11(_547_v0)).dtor_cf18;
      })(_504_v0)), _545_v22);
      let _549_v24;
      _549_v24 = _dafny.Seq.of((_this).f24);
      let _550_v25;
      let _nw75 = Array((new BigNumber(29)).toNumber());
      _nw75[(_dafny.ZERO).toNumber()] = p0;
      _nw75[(_dafny.ONE).toNumber()] = !(false);
      _nw75[(new BigNumber(2)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(3)).toNumber()] = p0;
      _nw75[(new BigNumber(4)).toNumber()] = true;
      _nw75[(new BigNumber(5)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(6)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(7)).toNumber()] = _module.__default.fm2(globalState);
      _nw75[(new BigNumber(8)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(9)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(10)).toNumber()] = p0;
      _nw75[(new BigNumber(11)).toNumber()] = p0;
      _nw75[(new BigNumber(12)).toNumber()] = !((_this).f24);
      _nw75[(new BigNumber(13)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(14)).toNumber()] = p0;
      _nw75[(new BigNumber(15)).toNumber()] = p0;
      _nw75[(new BigNumber(16)).toNumber()] = true;
      _nw75[(new BigNumber(17)).toNumber()] = p0;
      _nw75[(new BigNumber(18)).toNumber()] = p0;
      _nw75[(new BigNumber(19)).toNumber()] = (_this).fm4(new BigNumber(-538), _549_v24, globalState);
      _nw75[(new BigNumber(20)).toNumber()] = p0;
      _nw75[(new BigNumber(21)).toNumber()] = p0;
      _nw75[(new BigNumber(22)).toNumber()] = p0;
      _nw75[(new BigNumber(23)).toNumber()] = p0;
      _nw75[(new BigNumber(24)).toNumber()] = p0;
      _nw75[(new BigNumber(25)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(26)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(27)).toNumber()] = (_this).f24;
      _nw75[(new BigNumber(28)).toNumber()] = true;
      _550_v25 = _nw75;
      let _551_v26;
      _551_v26 = _module.D1.create_DC5(_504_v0, _504_v0, _546_v23, _550_v25, (_dafny.ZERO).minus(new BigNumber((_505_v1).length)));
      let _552_v27;
      _552_v27 = _dafny.MultiSet.fromElements((_551_v26).dtor_cf10, _550_v25, _550_v25, _550_v25);
      let _553_i1;
      _553_i1 = _dafny.ZERO;
      L1: {
        while (!(((_552_v27).Union(_552_v27)).IsSubsetOf(_552_v27))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_553_i1)) {
              break L1;
            }
            _553_i1 = (_553_i1).plus(_dafny.ONE);
            let _554_v28;
            _554_v28 = _dafny.Seq.of(_dafny.Seq.update(_543_v20, _module.__default.safeIndex(_504_v0, new BigNumber((_543_v20).length)), _this.f23), _543_v20);
            if (_dafny.areEqual((_554_v28)[_module.__default.safeIndex(_504_v0, new BigNumber((_554_v28).length))], _543_v20)) {
              (globalState).f13 = (_this).f24;
              _545_v22 = _545_v22;
              let _555_v29;
              let _init10 = ((_556_v0) => function (_557_i3) {
                return (_557_i3).multipliedBy(_556_v0);
              })(_504_v0);
              let _nw76 = Array((new BigNumber(2)).toNumber());
              for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw76.length); _i0_10++) {
                _nw76[_i0_10] = _init10(new BigNumber(_i0_10));
              }
              _555_v29 = _nw76;
              let _index69 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_555_v29).length));
              (_555_v29)[_index69] = _504_v0;
              let _index70 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_555_v29).length));
              (_555_v29)[_index70] = new BigNumber(553);
              let _558_v30;
              _558_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_555_v29)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_555_v29).length))]);
              let _559_v31;
              _559_v31 = _dafny.Map.Empty.slice().updateUnsafe(_558_v30,p0);
              _559_v31 = (_559_v31).update(_558_v30, !(_504_v0).isEqualTo(new BigNumber((_545_v22).length)));
              r0 = (((_555_v29)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_555_v29).length))]).minus((_this).fm24(new BigNumber(975), (_this).f24, globalState))).isLessThan((_504_v0).plus(_504_v0));
            } else {
              (globalState).f13 = (_504_v0).isLessThan(_504_v0);
              r1 = _module.__default.fm0(_dafny.Seq.IsProperPrefixOf(_543_v20, _543_v20), _504_v0, globalState);
              let _560_v32;
              _560_v32 = _module.D0.create_DC1(_504_v0);
              _560_v32 = _module.D0.create_DC1((_504_v0).multipliedBy(_504_v0));
              let _561_v33;
              let _nw77 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
              _561_v33 = _nw77;
              let _index71 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_561_v33).length));
              (_561_v33)[_index71] = (_504_v0).plus(new BigNumber(604));
              let _562_v34;
              _562_v34 = _dafny.Map.Empty.slice().updateUnsafe(_561_v33,_dafny.Seq.Concat(_545_v22, _545_v22));
              let _index72 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_561_v33).length));
              (_561_v33)[_index72] = new BigNumber((_module.__default.fm1(p0, globalState)).length);
              let _index73 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_561_v33).length));
              let _index74 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_561_v33).length));
              let _rhs72 = (_504_v0).plus((_504_v0).minus(_504_v0));
              let _rhs73 = (_this).f24;
              let _rhs74 = _562_v34;
              let _rhs75 = new BigNumber((_543_v20).length);
              let _rhs76 = _504_v0;
              let _lhs53 = _561_v33;
              let _lhs54 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_561_v33).length));
              let _lhs55 = globalState;
              let _lhs56 = _561_v33;
              let _lhs57 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_561_v33).length));
              _lhs53[_lhs54] = _rhs72;
              r0 = _rhs73;
              _562_v34 = _rhs74;
              _lhs55.f7 = _rhs75;
              _lhs56[_lhs57] = _rhs76;
              let _563_v35;
              _563_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(389)), function (_564_i4) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              }),new BigNumber(472));
              _563_v35 = (_563_v35).update(_dafny.Seq.UnicodeFromString("cjl"), (_561_v33)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_561_v33).length))]);
            }
            if (!(!(p0))) {
              let _565_v36;
              _565_v36 = _dafny.Map.Empty.slice().updateUnsafe(_550_v25,_543_v20);
              let _566_v37;
              _566_v37 = _module.D4.create_DC10((((_565_v36).contains(_550_v25)) ? ((_565_v36).get(_550_v25)) : (_543_v20)));
              let _rhs77 = (_504_v0).minus(_504_v0);
              let _rhs78 = _566_v37;
              let _rhs79 = (new BigNumber((_dafny.Seq.Concat(_543_v20, _543_v20)).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_504_v0));
              let _rhs80 = (_dafny.ZERO).minus(_504_v0);
              let _rhs81 = p0;
              let _lhs58 = globalState;
              let _lhs59 = globalState;
              let _lhs60 = globalState;
              let _lhs61 = globalState;
              _lhs58.f7 = _rhs77;
              _566_v37 = _rhs78;
              _lhs59.f13 = _rhs79;
              _lhs60.f7 = _rhs80;
              _lhs61.f13 = _rhs81;
              (globalState).f7 = (_504_v0).minus(_504_v0);
              (_this).f23 = _this.f23;
              r0 = (_module.D2.create_DC7(!((_this).f24), p0)).dtor_cf14;
              _543_v20 = _543_v20;
            } else {
              (globalState).f13 = (_this).f24;
              let _567_v38;
              _567_v38 = _dafny.Set.fromElements(!((_this).f24));
              let _568_v39;
              _568_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
              let _569_v40;
              _569_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),(_this).f24)).Merge(_568_v39)).length));
              let _570_v41;
              _570_v41 = _module.D0.create_DC1(_504_v0);
              let _571_v42;
              _571_v42 = _module.D0.create_DC3(_570_v41);
              let _572_v43;
              _572_v43 = _module.D0.create_DC3(_571_v42);
              let _573_v44;
              let _nw78 = Array((new BigNumber(9)).toNumber()).fill(_module.D4.Default());
              _573_v44 = _nw78;
              let _574_v45;
              _574_v45 = _module.D4.create_DC10(_dafny.Seq.UnicodeFromString("jgil"));
              let _index75 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_573_v44).length));
              (_573_v44)[_index75] = _574_v45;
              let _575_v46;
              _575_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
              let _index76 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_573_v44).length));
              let _rhs82 = _567_v38;
              let _rhs83 = _dafny.Map.Empty.slice().updateUnsafe((_504_v0).isLessThan(new BigNumber((_575_v46).length)),(_504_v0).minus(_504_v0));
              let _rhs84 = _572_v43;
              let _rhs85 = _574_v45;
              let _lhs62 = _573_v44;
              let _lhs63 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_573_v44).length));
              _567_v38 = _rhs82;
              _569_v40 = _rhs83;
              _572_v43 = _rhs84;
              _lhs62[_lhs63] = _rhs85;
              _504_v0 = _504_v0;
              (globalState).f7 = _504_v0;
              let _576_v47;
              let _nw79 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
              _576_v47 = _nw79;
              let _577_v48;
              let _init11 = function (_578_i5) {
                return (_578_i5).plus(new BigNumber(291));
              };
              let _nw80 = Array((new BigNumber(7)).toNumber());
              for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw80.length); _i0_11++) {
                _nw80[_i0_11] = _init11(new BigNumber(_i0_11));
              }
              _577_v48 = _nw80;
              let _579_v49;
              _579_v49 = _dafny.Map.Empty.slice().updateUnsafe(_577_v48,_504_v0);
              let _index77 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_576_v47).length));
              (_576_v47)[_index77] = _579_v49;
              let _index78 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_576_v47).length));
              let _rhs86 = p0;
              let _rhs87 = (((_module.__default.fm2(globalState)) ? (((_579_v49).update(_577_v48, new BigNumber((_575_v46).length))).update(_577_v48, _504_v0)) : (_579_v49))).Merge(_579_v49);
              let _lhs64 = _576_v47;
              let _lhs65 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_576_v47).length));
              r2 = _rhs86;
              _lhs64[_lhs65] = _rhs87;
            }
            let _580_v50;
            _580_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
            let _581_v51;
            _581_v51 = _module.D0.create_DC2(p0, _504_v0, (((_580_v50).contains(p0)) ? ((_580_v50).get(p0)) : (!(p0))));
            let _index79 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_550_v25).length));
            (_550_v25)[_index79] = (_581_v51).dtor_cf4;
            let _index80 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_550_v25).length));
            (_550_v25)[_index80] = (_this).f24;
            r2 = p0;
          }
        }
      }
      let _582_v52;
      let _init12 = ((_583_v0) => function (_584_i6) {
        return _dafny.MultiSet.fromElements(_583_v0);
      })(_504_v0);
      let _nw81 = Array((new BigNumber(3)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw81.length); _i0_12++) {
        _nw81[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _582_v52 = _nw81;
      let _585_v53;
      _585_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
      let _586_v54;
      _586_v54 = _dafny.MultiSet.fromElements(_504_v0, new BigNumber((_585_v53).length));
      let _index81 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_582_v52).length));
      (_582_v52)[_index81] = (_dafny.MultiSet.fromElements(_504_v0)).Intersect(_586_v54);
      let _index82 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_582_v52).length));
      (_582_v52)[_index82] = (_dafny.MultiSet.FromArray(_545_v22)).Intersect(_586_v54);
      let _hi2 = (((_this).f24) ? ((_dafny.ZERO).minus(_504_v0)) : ((_dafny.ZERO).minus(_504_v0)));
      for (let _587_i7 = _504_v0; _587_i7.isLessThan(_hi2); _587_i7 = _587_i7.plus(_dafny.ONE)) {
        (globalState).f13 = (_505_v1).IsDisjointFrom(_505_v1);
        r0 = (_this).f24;
        let _588_v55;
        let _nw82 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
        _588_v55 = _nw82;
        let _589_v56;
        _589_v56 = _module.D5.create_DC13(_504_v0, (_this).f24, p0);
        let _590_v57;
        _590_v57 = _dafny.Set.fromElements(_dafny.Seq.of(_589_v56, _589_v56));
        let _591_v58;
        _591_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_590_v57);
        let _rhs88 = (_module.__default.fm27((_this).f24, p0, globalState)).equals((((_591_v58).contains(p0)) ? ((_591_v58).get(p0)) : (_590_v57)));
        let _rhs89 = _this.f23;
        let _rhs90 = (((_this).f24) ? (_588_v55) : (_588_v55));
        let _lhs66 = globalState;
        let _lhs67 = _this;
        _lhs66.f13 = _rhs88;
        _lhs67.f23 = _rhs89;
        _588_v55 = _rhs90;
        let _592_v59;
        let _nw83 = new _module.C0();
        _nw83.__ctor((_dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("owqe"))).update((_this).f24, _dafny.Seq.UnicodeFromString("wdpkd")), (_this).f24, _this.f23, true);
        _592_v59 = _nw83;
      }
      let _593_v60;
      let _nw84 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _593_v60 = _nw84;
      let _index83 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_593_v60).length));
      (_593_v60)[_index83] = _504_v0;
      let _index84 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_593_v60).length));
      (_593_v60)[_index84] = _module.__default.safeDivisionInt(new BigNumber((_586_v54).cardinality()), _504_v0);
      r1 = new BigNumber((_543_v20).length);
      r0 = (_dafny.MultiSet.fromElements(false)).IsProperSubsetOf((_this).f36);
      r1 = _504_v0;
      r2 = (_this).f24;
      return [r0, r1, r2];
    }
    m12(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.ZERO;
      let _594_v0;
      _594_v0 = _dafny.Seq.UnicodeFromString("ulpqjjbk");
      r1 = _dafny.Seq.Concat(_594_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), function (_595_i0) {
        return _this.f23;
      }));
      let _596_v1;
      _596_v1 = _dafny.Seq.of((_this).f24);
      let _rhs91 = (p1).plus(p1);
      let _rhs92 = new BigNumber((_596_v1).length);
      let _lhs68 = globalState;
      let _lhs69 = globalState;
      _lhs68.f7 = _rhs91;
      _lhs69.f7 = _rhs92;
      let _597_v2;
      _597_v2 = _module.D1.create_DC4((_594_v0)[_module.__default.safeIndex(p1, new BigNumber((_594_v0).length))]);
      let _source10 = _597_v2;
      if (_source10.is_DC5) {
        let _598___mcc_h0 = (_source10).cf7;
        let _599___mcc_h1 = (_source10).cf8;
        let _600___mcc_h2 = (_source10).cf9;
        let _601___mcc_h3 = (_source10).cf10;
        let _602___mcc_h4 = (_source10).cf11;
        let _603_cf11 = _602___mcc_h4;
        let _604_cf10 = _601___mcc_h3;
        let _605_cf9 = _600___mcc_h2;
        let _606_cf8 = _599___mcc_h1;
        let _607_cf7 = _598___mcc_h0;
        let _608_v3;
        _608_v3 = _dafny.Seq.of(_603_cf11);
        let _609_v4;
        _609_v4 = _dafny.Set.fromElements(_607_cf7, (_dafny.ZERO).minus(p1), (_608_v3)[_module.__default.safeIndex(_603_cf11, new BigNumber((_608_v3).length))]);
        let _rhs93 = _module.__default.fm0((_this).f24, _607_cf7, globalState);
        let _rhs94 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_609_v4).length)), new BigNumber((_596_v1).length));
        let _rhs95 = (p1).isLessThan((_dafny.ZERO).minus(_606_cf8));
        let _lhs70 = globalState;
        let _lhs71 = globalState;
        _603_cf11 = _rhs93;
        _lhs70.f7 = _rhs94;
        _lhs71.f13 = _rhs95;
        let _610_v5;
        _610_v5 = _module.D0.create_DC2((_this).f24, _607_cf7, (_this).f24);
        let _611_v6;
        _611_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_607_cf7);
        let _pat_let_tv18 = _611_v6;
        _610_v5 = function (_pat_let19_0) {
          return function (_612_dt__update__tmp_h0) {
            return function (_pat_let20_0) {
              return function (_613_dt__update_hcf3_h0) {
                return function (_pat_let21_0) {
                  return function (_614_dt__update_hcf4_h0) {
                    return _module.D0.create_DC2((_612_dt__update__tmp_h0).dtor_cf2, _613_dt__update_hcf3_h0, _614_dt__update_hcf4_h0);
                  }(_pat_let21_0);
                }(true);
              }(_pat_let20_0);
            }(new BigNumber((_pat_let_tv18).length));
          }(_pat_let19_0);
        }(_610_v5);
        let _615_v7;
        _615_v7 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_594_v0);
        let _616_v8;
        _616_v8 = _module.D2.create_DC6(_615_v7);
        _616_v8 = (((_this).f24) ? (_616_v8) : (_module.D2.create_DC6(_615_v7)));
        let _617_v9;
        _617_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(-372));
        let _618_v10;
        _618_v10 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_606_cf8), _617_v9, _617_v9, _617_v9);
        let _source11 = _module.__default.fm28(_594_v0, _608_v3, (_596_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_594_v0, _module.__default.safeIndex(new BigNumber((_594_v0).length), new BigNumber((_594_v0).length)), _this.f23)).length), new BigNumber((_596_v1).length))], _618_v10, globalState);
        if (_source11.is_DC1) {
          let _619___mcc_h6 = (_source11).cf1;
          let _620_cf1 = _619___mcc_h6;
          let _621_v11;
          let _nw85 = new _module.C0();
          _nw85.__ctor(_615_v7, (_module.D0.create_DC2(true, _607_cf7, (_this).f24)).dtor_cf4, _this.f23, (_this).f24);
          _621_v11 = _nw85;
          let _622_v12;
          _622_v12 = _dafny.Seq.of(_604_cf10);
          let _623_v13;
          _623_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_621_v11).f38, _606_cf8, globalState),(_622_v12)[_module.__default.safeIndex(new BigNumber((_609_v4).length), new BigNumber((_622_v12).length))]);
          _623_v13 = (_623_v13).update((_608_v3)[_module.__default.safeIndex(_603_cf11, new BigNumber((_608_v3).length))], _604_cf10);
          let _624_v14;
          _624_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_594_v0).length),_603_cf11);
          _624_v14 = (_624_v14).update((new BigNumber(843)).minus(p1), new BigNumber(474));
          let _625_v15;
          _625_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_617_v9).length),_596_v1);
          _625_v15 = (_625_v15).update((_dafny.ZERO).minus(_606_cf8), _dafny.Seq.Concat(_596_v1, _596_v1));
        } else if (_source11.is_DC2) {
          let _626___mcc_h7 = (_source11).cf2;
          let _627___mcc_h8 = (_source11).cf3;
          let _628___mcc_h9 = (_source11).cf4;
          let _629_cf4 = _628___mcc_h9;
          let _630_cf3 = _627___mcc_h8;
          let _631_cf2 = _626___mcc_h7;
          (globalState).f7 = _module.__default.safeModuloInt((_607_cf7).plus(_603_cf11), _607_cf7);
          (globalState).f7 = _630_cf3;
          let _632_v16;
          _632_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_603_cf11);
          let _pat_let_tv19 = p1;
          let _pat_let_tv20 = _631_cf2;
          let _rhs96 = (((function (_pat_let22_0) {
            return function (_633_dt__update__tmp_h1) {
              return function (_pat_let23_0) {
                return function (_634_dt__update_hcf3_h1) {
                  return function (_pat_let24_0) {
                    return function (_635_dt__update_hcf4_h1) {
                      return _module.D0.create_DC2((_633_dt__update__tmp_h1).dtor_cf2, _634_dt__update_hcf3_h1, _635_dt__update_hcf4_h1);
                    }(_pat_let24_0);
                  }(_pat_let_tv20);
                }(_pat_let23_0);
              }(_pat_let_tv19);
            }(_pat_let22_0);
          }(_610_v5)).dtor_cf4) ? (_this.f23) : (_module.__default.fm29(_630_cf3, _629_cf4, _631_cf2, _603_cf11, globalState)));
          let _rhs97 = (function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of (_608_v3).Elements) {
              let _636_v17 = _compr_62;
              if (_dafny.Seq.contains(_608_v3, _636_v17)) {
                _coll62.push([_module.__default.safeModuloInt(_636_v17, _630_cf3),_603_cf11]);
              }
            }
            return _coll62;
          }()).Merge((_632_v16).Merge(_dafny.Map.Empty.slice().updateUnsafe(_606_cf8,_603_cf11)));
          let _lhs72 = _this;
          _lhs72.f23 = _rhs96;
          _632_v16 = _rhs97;
          let _rhs98 = _dafny.Seq.Concat(_594_v0, _dafny.Seq.update(_dafny.Seq.Concat(_594_v0, _594_v0), _module.__default.safeIndex(_606_cf8, new BigNumber((_dafny.Seq.Concat(_594_v0, _594_v0)).length)), new _dafny.CodePoint('v'.codePointAt(0))));
          let _rhs99 = _this.f23;
          let _lhs73 = _this;
          r1 = _rhs98;
          _lhs73.f23 = _rhs99;
        } else if (_source11.is_DC0) {
          let _637___mcc_h10 = (_source11).cf0;
          let _638_cf0 = _637___mcc_h10;
          _594_v0 = _dafny.Seq.update(_594_v0, _module.__default.safeIndex(_606_cf8, new BigNumber((_594_v0).length)), new _dafny.CodePoint('a'.codePointAt(0)));
          let _639_v18;
          let _init13 = ((_640_v2) => function (_641_i1) {
            return (_641_i1).multipliedBy(new BigNumber((_dafny.Set.fromElements((_640_v2).dtor_cf6, _this.f23)).length));
          })(_597_v2);
          let _nw86 = Array((new BigNumber(11)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw86.length); _i0_13++) {
            _nw86[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _639_v18 = _nw86;
          let _rhs100 = ((_dafny.ZERO).minus(_606_cf8)).multipliedBy((_dafny.ZERO).minus(_module.__default.safeModuloInt(_603_cf11, _603_cf11)));
          let _rhs101 = _dafny.Seq.contains(_dafny.Seq.Concat(_596_v1, _596_v1), false);
          let _rhs102 = (_607_cf7).minus((_dafny.ZERO).minus(_607_cf7));
          let _rhs103 = _639_v18;
          let _lhs74 = globalState;
          let _lhs75 = globalState;
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          _lhs74.f7 = _rhs100;
          _lhs75.f13 = _rhs101;
          _lhs76.f7 = _rhs102;
          _lhs77.f22 = _rhs103;
          let _642_v19;
          _642_v19 = _module.D2.create_DC7((_this).f24, (_this).f24);
          let _643_v20;
          let _nw87 = new _module.C0();
          _nw87.__ctor(_615_v7, (_642_v19).dtor_cf14, _this.f23, (_this).f24);
          _643_v20 = _nw87;
          let _644_v21;
          _644_v21 = _dafny.MultiSet.fromElements(new BigNumber((_594_v0).length));
          let _645_v22;
          _645_v22 = _dafny.Set.fromElements((_644_v21).Difference(_644_v21));
          let _646_v23;
          _646_v23 = _dafny.Seq.of(_594_v0, _594_v0);
          let _rhs104 = _module.__default.fm30((((_this).f24) ? (_616_v8) : (_616_v8)), _646_v23, globalState);
          let _rhs105 = (_dafny.ZERO).minus(_607_cf7);
          _645_v22 = _rhs104;
          r0 = _rhs105;
        } else {
          let _647___mcc_h11 = (_source11).cf5;
          let _648_cf5 = _647___mcc_h11;
          _603_cf11 = _607_cf7;
          let _649_v24;
          _649_v24 = _dafny.MultiSet.fromElements(new BigNumber((_596_v1).length), _607_cf7, p1);
          let _650_v25;
          _650_v25 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_607_cf7), _649_v24, _649_v24);
          let _651_v26;
          _651_v26 = _module.D5.create_DC12(_dafny.Set.fromElements(p1));
          let _652_v27;
          let _nw88 = Array((new BigNumber(15)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ogv"), _594_v0)).length);
          _nw88[(_dafny.ONE).toNumber()] = _606_cf8;
          _nw88[(new BigNumber(2)).toNumber()] = new BigNumber((_609_v4).length);
          _nw88[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("mcybk")).length);
          _nw88[(new BigNumber(4)).toNumber()] = (new BigNumber(((_651_v26).dtor_cf19).length)).plus((((_649_v24).contains(_607_cf7)) ? ((_649_v24).get(_607_cf7)) : (new BigNumber((_594_v0).length))));
          _nw88[(new BigNumber(5)).toNumber()] = new BigNumber(276);
          _nw88[(new BigNumber(6)).toNumber()] = _606_cf8;
          _nw88[(new BigNumber(7)).toNumber()] = (new BigNumber((_608_v3).length)).plus((((_617_v9).contains((_596_v1)[_module.__default.safeIndex(new BigNumber((_608_v3).length), new BigNumber((_596_v1).length))])) ? ((_617_v9).get((_596_v1)[_module.__default.safeIndex(new BigNumber((_608_v3).length), new BigNumber((_596_v1).length))])) : (p1)));
          _nw88[(new BigNumber(8)).toNumber()] = p1;
          _nw88[(new BigNumber(9)).toNumber()] = _607_cf7;
          _nw88[(new BigNumber(10)).toNumber()] = p1;
          _nw88[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((p1).plus(p1));
          _nw88[(new BigNumber(12)).toNumber()] = _module.__default.fm0(false, _606_cf8, globalState);
          _nw88[(new BigNumber(13)).toNumber()] = _603_cf11;
          _nw88[(new BigNumber(14)).toNumber()] = new BigNumber(((_module.D4.create_DC10(_594_v0)).dtor_cf17).length);
          _652_v27 = _nw88;
          let _index85 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_652_v27).length));
          (_652_v27)[_index85] = p1;
          let _653_v28;
          _653_v28 = _module.D5.create_DC13(_603_cf11, true, (_this).f24);
          let _index86 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_604_cf10).length));
          (_604_cf10)[_index86] = (function (_pat_let25_0) {
            return function (_654_dt__update__tmp_h2) {
              return function (_pat_let26_0) {
                return function (_655_dt__update_hcf21_h0) {
                  return _module.D5.create_DC13((_654_dt__update__tmp_h2).dtor_cf20, _655_dt__update_hcf21_h0, (_654_dt__update__tmp_h2).dtor_cf22);
                }(_pat_let26_0);
              }((_this).f24);
            }(_pat_let25_0);
          }(_653_v28)).dtor_cf21;
          let _656_v29;
          _656_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_604_cf10);
          let _657_v30;
          _657_v30 = _dafny.Set.fromElements((_this).f24);
          let _index87 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_652_v27).length));
          let _index88 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_604_cf10).length));
          let _rhs106 = _module.__default.fm31((_this).f24, !((_this).f24), (_this).f24, _610_v5, globalState);
          let _rhs107 = (_module.__default.safeDivisionInt(_607_cf7, _607_cf7)).multipliedBy((new BigNumber((_656_v29).length)).minus(_603_cf11));
          let _rhs108 = (_657_v30).IsDisjointFrom((_657_v30).Intersect(_dafny.Set.fromElements((_596_v1)[_module.__default.safeIndex(_606_cf8, new BigNumber((_596_v1).length))])));
          let _lhs78 = _652_v27;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_652_v27).length));
          let _lhs80 = _604_cf10;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_604_cf10).length));
          _650_v25 = _rhs106;
          _lhs78[_lhs79] = _rhs107;
          _lhs80[_lhs81] = _rhs108;
          let _658_v31;
          let _nw89 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _658_v31 = _nw89;
          let _index89 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_658_v31).length));
          (_658_v31)[_index89] = _594_v0;
          let _index90 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_658_v31).length));
          (_658_v31)[_index90] = _dafny.Seq.UnicodeFromString("qxxltvc");
          let _659_v32;
          _659_v32 = _dafny.Map.Empty.slice().updateUnsafe(_649_v24,_603_cf11);
          let _660_v33;
          _660_v33 = _dafny.Map.Empty.slice().updateUnsafe(_603_cf11,_603_cf11);
          _659_v32 = (_659_v32).update(_649_v24, _module.__default.fm0(!((_604_cf10)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_604_cf10).length))]), new BigNumber((_660_v33).length), globalState));
        }
      } else {
        let _661___mcc_h5 = (_source10).cf6;
        let _662_cf6 = _661___mcc_h5;
        let _source12 = _module.D2.create_DC7((_this).f24, (_this).f24);
        if (_source12.is_DC7) {
          let _663___mcc_h12 = (_source12).cf13;
          let _664___mcc_h13 = (_source12).cf14;
          let _665_cf14 = _664___mcc_h13;
          let _666_cf13 = _663___mcc_h12;
          let _667_v34;
          _667_v34 = _module.D4.create_DC11(p1);
          let _pat_let_tv21 = p1;
          let _pat_let_tv22 = p1;
          _667_v34 = function (_pat_let27_0) {
            return function (_668_dt__update__tmp_h3) {
              return function (_pat_let28_0) {
                return function (_669_dt__update_hcf18_h0) {
                  return _module.D4.create_DC11(_669_dt__update_hcf18_h0);
                }(_pat_let28_0);
              }((_pat_let_tv21).plus(_pat_let_tv22));
            }(_pat_let27_0);
          }(_module.D4.create_DC11(p1));
          let _670_v35;
          _670_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p1);
          _670_v35 = (_670_v35).update(!(!(true)) || (_665_cf14), p1);
          let _671_v36;
          _671_v36 = _dafny.Seq.of(p1);
          let _672_v37;
          let _nw90 = new _module.C1();
          _nw90.__ctor(true, _666_cf13);
          _672_v37 = _nw90;
          let _673_v38;
          let _nw91 = Array((new BigNumber(13)).toNumber());
          _nw91[(_dafny.ZERO).toNumber()] = _672_v37;
          _nw91[(_dafny.ONE).toNumber()] = _672_v37;
          _nw91[(new BigNumber(2)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(3)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(4)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(5)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(6)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(7)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(8)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(9)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(10)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(11)).toNumber()] = _672_v37;
          _nw91[(new BigNumber(12)).toNumber()] = _672_v37;
          _673_v38 = _nw91;
          let _674_v39;
          _674_v39 = _dafny.Map.Empty.slice().updateUnsafe(_673_v38,_this.f23);
          let _rhs109 = ((_665_cf14) ? (_dafny.Seq.of(new BigNumber(-18))) : (_dafny.Seq.Concat(_671_v36, _671_v36)));
          let _rhs110 = _674_v39;
          _671_v36 = _rhs109;
          _674_v39 = _rhs110;
          let _675_v40;
          let _nw92 = Array((new BigNumber(18)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _596_v1;
          _nw92[(_dafny.ONE).toNumber()] = ((_666_cf13) ? (_596_v1) : (_dafny.Seq.of((_672_v37).f27)));
          _nw92[(new BigNumber(2)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(3)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(4)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(5)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(6)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(7)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(8)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_672_v37).f27, (_this).f24), _dafny.Seq.of(_665_cf14));
          _nw92[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of((_this).f24), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of((_this).f24)).length)), (_672_v37).f27);
          _nw92[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm2(globalState)), _596_v1);
          _nw92[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(true, _666_cf13), _596_v1);
          _nw92[(new BigNumber(13)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_596_v1, _596_v1);
          _nw92[(new BigNumber(15)).toNumber()] = _module.__default.fm36(globalState);
          _nw92[(new BigNumber(16)).toNumber()] = _596_v1;
          _nw92[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_596_v1, _dafny.Seq.of((_672_v37).f27));
          _675_v40 = _nw92;
          let _676_v41;
          _676_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(440),_596_v1);
          let _677_v42;
          _677_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_672_v37).f27);
          let _678_v43;
          _678_v43 = _module.D6.create_DC16(_677_v42, (_672_v37).f27, (_672_v37).f27);
          let _679_v44;
          _679_v44 = _dafny.Set.fromElements((_this).f25);
          let _680_v45;
          let _nw93 = Array((new BigNumber(27)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _665_cf14;
          _nw93[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_594_v0, _662_cf6);
          _nw93[(new BigNumber(2)).toNumber()] = _665_cf14;
          _nw93[(new BigNumber(3)).toNumber()] = (_this).f24;
          _nw93[(new BigNumber(4)).toNumber()] = _665_cf14;
          _nw93[(new BigNumber(5)).toNumber()] = ((_this).f36).IsProperSubsetOf(_dafny.MultiSet.fromElements(false, (_this).f24));
          _nw93[(new BigNumber(6)).toNumber()] = !(_676_v41).equals(_676_v41);
          _nw93[(new BigNumber(7)).toNumber()] = (_665_cf14) === (_666_cf13);
          _nw93[(new BigNumber(8)).toNumber()] = (_678_v43).dtor_cf27;
          _nw93[(new BigNumber(9)).toNumber()] = false;
          _nw93[(new BigNumber(10)).toNumber()] = !(_665_cf14);
          _nw93[(new BigNumber(11)).toNumber()] = (_module.D2.create_DC7((_672_v37).f27, (_672_v37).f27)).dtor_cf13;
          _nw93[(new BigNumber(12)).toNumber()] = _666_cf13;
          _nw93[(new BigNumber(13)).toNumber()] = !((_679_v44).IsProperSubsetOf(_679_v44));
          _nw93[(new BigNumber(14)).toNumber()] = (_this).f24;
          _nw93[(new BigNumber(15)).toNumber()] = (_672_v37).f27;
          _nw93[(new BigNumber(16)).toNumber()] = (_this).f24;
          _nw93[(new BigNumber(17)).toNumber()] = (_672_v37).f27;
          _nw93[(new BigNumber(18)).toNumber()] = (_this).f24;
          _nw93[(new BigNumber(19)).toNumber()] = true;
          _nw93[(new BigNumber(20)).toNumber()] = (_672_v37).f27;
          _nw93[(new BigNumber(21)).toNumber()] = !(false);
          _nw93[(new BigNumber(22)).toNumber()] = (_this).f24;
          _nw93[(new BigNumber(23)).toNumber()] = (_this).f24;
          _nw93[(new BigNumber(24)).toNumber()] = !((p1).isLessThan(p1));
          _nw93[(new BigNumber(25)).toNumber()] = true;
          _nw93[(new BigNumber(26)).toNumber()] = (_this).f24;
          _680_v45 = _nw93;
          let _681_v46;
          _681_v46 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_675_v40);
          let _rhs111 = (((_681_v46).contains(_module.__default.safeDivisionInt(new BigNumber(675), p1))) ? ((_681_v46).get(_module.__default.safeDivisionInt(new BigNumber(675), p1))) : (_675_v40));
          let _rhs112 = true;
          let _rhs113 = _680_v45;
          let _lhs82 = globalState;
          _675_v40 = _rhs111;
          _lhs82.f13 = _rhs112;
          _680_v45 = _rhs113;
        } else if (_source12.is_DC6) {
          let _682___mcc_h14 = (_source12).cf12;
          let _683_cf12 = _682___mcc_h14;
          let _684_v47;
          _684_v47 = _dafny.Map.Empty.slice().updateUnsafe(_594_v0,new BigNumber((_594_v0).length));
          r0 = (((_684_v47).contains(_594_v0)) ? ((_684_v47).get(_594_v0)) : (p1));
          (globalState).f13 = !((_this).f24);
          let _685_v48;
          let _nw94 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _685_v48 = _nw94;
          let _index91 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_685_v48).length));
          (_685_v48)[_index91] = p1;
          let _index92 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_685_v48).length));
          (_685_v48)[_index92] = _module.__default.safeDivisionInt(p1, p1);
          let _index93 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_685_v48).length));
          (_685_v48)[_index93] = (_685_v48)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_685_v48).length))];
        } else {
          let _686___mcc_h15 = (_source12).cf15;
          let _687_cf15 = _686___mcc_h15;
          let _688_v49;
          _688_v49 = _dafny.MultiSet.fromElements(new BigNumber(535));
          let _689_v50;
          _689_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,true);
          let _690_v51;
          _690_v51 = _dafny.Map.Empty.slice().updateUnsafe(_689_v50,new BigNumber(-793));
          let _691_v52;
          let _nw95 = Array((new BigNumber(16)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = p1;
          _nw95[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw95[(new BigNumber(2)).toNumber()] = p1;
          _nw95[(new BigNumber(3)).toNumber()] = p1;
          _nw95[(new BigNumber(4)).toNumber()] = ((((_this).f36).contains((_this).f24)) ? (((_this).f36).get((_this).f24)) : (new BigNumber(-489)));
          _nw95[(new BigNumber(5)).toNumber()] = new BigNumber((_688_v49).cardinality());
          _nw95[(new BigNumber(6)).toNumber()] = p1;
          _nw95[(new BigNumber(7)).toNumber()] = p1;
          _nw95[(new BigNumber(8)).toNumber()] = _module.__default.fm0(true, new BigNumber((_594_v0).length), globalState);
          _nw95[(new BigNumber(9)).toNumber()] = new BigNumber(529);
          _nw95[(new BigNumber(10)).toNumber()] = (((_690_v51).contains(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,true))) ? ((_690_v51).get(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,true))) : (new BigNumber(444)));
          _nw95[(new BigNumber(11)).toNumber()] = p1;
          _nw95[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_596_v1)[_module.__default.safeIndex(p1, new BigNumber((_596_v1).length))])).cardinality()));
          _nw95[(new BigNumber(13)).toNumber()] = p1;
          _nw95[(new BigNumber(14)).toNumber()] = p1;
          _nw95[(new BigNumber(15)).toNumber()] = p1;
          _691_v52 = _nw95;
          let _692_v53;
          _692_v53 = _dafny.Map.Empty.slice().updateUnsafe(_691_v52,(_this).f24);
          _692_v53 = _692_v53;
          let _693_v54;
          _693_v54 = _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f24), (_this).f24, (_this).f24);
          let _694_v55;
          _694_v55 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
          let _695_v56;
          _695_v56 = _dafny.Set.fromElements(p1, p1, p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_691_v52)).length), p1);
          let _pat_let_tv23 = _689_v50;
          let _696_v57;
          let _nw96 = Array((new BigNumber(20)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = (_this).f24;
          _nw96[(_dafny.ONE).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(2)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(3)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(4)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(5)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(6)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(7)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(8)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(9)).toNumber()] = !(_688_v49).contains((_dafny.ZERO).minus(p1));
          _nw96[(new BigNumber(10)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(11)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(12)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(13)).toNumber()] = !((function (_pat_let29_0) {
            return function (_697_dt__update__tmp_h4) {
              return function (_pat_let30_0) {
                return function (_698_dt__update_hcf25_h0) {
                  return _module.D6.create_DC16(_698_dt__update_hcf25_h0, (_697_dt__update__tmp_h4).dtor_cf26, (_697_dt__update__tmp_h4).dtor_cf27);
                }(_pat_let30_0);
              }(_pat_let_tv23);
            }(_pat_let29_0);
          }(_693_v54)).dtor_cf27) || ((_this).f24);
          _nw96[(new BigNumber(14)).toNumber()] = ((_dafny.ZERO).minus((((_694_v55).contains((_this).f24)) ? ((_694_v55).get((_this).f24)) : (p1)))).isLessThan(p1);
          _nw96[(new BigNumber(15)).toNumber()] = (_this).f24;
          _nw96[(new BigNumber(16)).toNumber()] = !(p1).isEqualTo(new BigNumber((_695_v56).length));
          _nw96[(new BigNumber(17)).toNumber()] = !((_this).f24);
          _nw96[(new BigNumber(18)).toNumber()] = ((_this).f36).IsProperSubsetOf((_this).f36);
          _nw96[(new BigNumber(19)).toNumber()] = (_this).f24;
          _696_v57 = _nw96;
          let _index94 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_696_v57).length));
          (_696_v57)[_index94] = (_this).f24;
          let _index95 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_696_v57).length));
          (_696_v57)[_index95] = (false) === (false);
          let _699_v58;
          _699_v58 = _dafny.Seq.of(_691_v52, _691_v52, _691_v52);
          let _700_v59;
          _700_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f24);
          _699_v58 = (((function () {
            let _coll63 = new _dafny.Set();
            for (const _compr_63 of (_700_v59).Keys.Elements) {
              let _701_v60 = _compr_63;
              if ((_700_v59).contains(_701_v60)) {
                _coll63.add((_701_v60).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-807),_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("mo")).length)))).length)));
              }
            }
            return _coll63;
          }()).IsSubsetOf(_695_v56)) ? (_699_v58) : (_dafny.Seq.Concat(_699_v58, _699_v58)));
          (globalState).f13 = (_696_v57)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_696_v57).length))];
        }
        let _702_v61;
        let _nw97 = new _module.C1();
        _nw97.__ctor((_this).f24, (_this).f24);
        _702_v61 = _nw97;
        let _703_v62;
        let _init14 = ((_704_v0) => function (_705_i2) {
          return _704_v0;
        })(_594_v0);
        let _nw98 = Array((new BigNumber(25)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw98.length); _i0_14++) {
          _nw98[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _703_v62 = _nw98;
        let _index96 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length));
        (_703_v62)[_index96] = _594_v0;
        let _706_v63;
        _706_v63 = _module.D4.create_DC10(_594_v0);
        let _index97 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length));
        (_703_v62)[_index97] = ((((_this).f24) ? (_706_v63) : (_706_v63))).dtor_cf17;
        if (_module.__default.fm2(globalState)) {
          let _707_v64;
          let _init15 = ((_708_p1) => function (_709_i3) {
            return _module.__default.safeDivisionInt(_709_i3, _708_p1);
          })(p1);
          let _nw99 = Array((new BigNumber(23)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw99.length); _i0_15++) {
            _nw99[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _707_v64 = _nw99;
          let _710_v65;
          _710_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,_707_v64);
          _710_v65 = (_710_v65).update(p1, _707_v64);
          (globalState).f13 = !((_this).f36).contains((_702_v61).f39);
          let _711_v66;
          let _nw100 = Array((new BigNumber(24)).toNumber()).fill(false);
          _711_v66 = _nw100;
          let _712_v67;
          _712_v67 = _dafny.Set.fromElements(p1);
          let _index98 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_711_v66).length));
          (_711_v66)[_index98] = !((_702_v61).fm32((_702_v61).f39, new BigNumber((_712_v67).length), globalState));
          let _713_v68;
          _713_v68 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("sn"), (_703_v62)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length))]);
          let _index99 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_711_v66).length));
          let _rhs114 = (_702_v61).f39;
          let _rhs115 = _module.__default.safeModuloInt(p1, p1);
          let _rhs116 = (_713_v68)[_module.__default.safeIndex(new BigNumber((_594_v0).length), new BigNumber((_713_v68).length))];
          let _rhs117 = (p1).isLessThanOrEqualTo(p1);
          let _rhs118 = p1;
          let _lhs83 = globalState;
          let _lhs84 = _711_v66;
          let _lhs85 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_711_v66).length));
          let _lhs86 = globalState;
          _lhs83.f13 = _rhs114;
          r2 = _rhs115;
          r1 = _rhs116;
          _lhs84[_lhs85] = _rhs117;
          _lhs86.f7 = _rhs118;
          let _714_v69;
          let _init16 = ((_715_p1, _716_v61) => function (_717_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-632),_715_p1)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_716_v61).f39,_715_p1));
          })(p1, _702_v61);
          let _nw101 = Array((new BigNumber(2)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw101.length); _i0_16++) {
            _nw101[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _714_v69 = _nw101;
          let _nw102 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _714_v69 = _nw102;
          let _718_v70;
          _718_v70 = _dafny.Set.fromElements((_this).f24);
          let _index100 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_711_v66).length));
          (_711_v66)[_index100] = (((_711_v66)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_711_v66).length))]) ? ((_this).f24) : ((_718_v70).IsSubsetOf(_718_v70)));
        } else {
          (globalState).f13 = ((_dafny.areEqual((_703_v62)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length))], (_703_v62)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length))])) ? ((_596_v1)[_module.__default.safeIndex(p1, new BigNumber((_596_v1).length))]) : (!((_this).f24)));
          (globalState).f13 = (_702_v61).f39;
          let _index101 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length));
          (_703_v62)[_index101] = (_703_v62)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_703_v62).length))];
          (globalState).f13 = (_702_v61).f39;
          let _719_v71;
          _719_v71 = _module.D2.create_DC7(true, (_702_v61).f39);
          let _720_v72;
          let _nw103 = Array((new BigNumber(23)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _719_v71;
          _nw103[(_dafny.ONE).toNumber()] = _719_v71;
          _nw103[(new BigNumber(2)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(3)).toNumber()] = _module.D2.create_DC7((_this).fm4(p1, _596_v1, globalState), (_702_v61).f39);
          _nw103[(new BigNumber(4)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(5)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(6)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(7)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(8)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(9)).toNumber()] = _module.D2.create_DC7((_702_v61).f39, (_702_v61).f39);
          _nw103[(new BigNumber(10)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(11)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(12)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(13)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(14)).toNumber()] = function (_pat_let31_0) {
            return function (_721_dt__update__tmp_h5) {
              return function (_pat_let32_0) {
                return function (_722_dt__update_hcf13_h0) {
                  return _module.D2.create_DC7(_722_dt__update_hcf13_h0, (_721_dt__update__tmp_h5).dtor_cf14);
                }(_pat_let32_0);
              }((_this).f24);
            }(_pat_let31_0);
          }(_719_v71);
          _nw103[(new BigNumber(15)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(16)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(17)).toNumber()] = _module.__default.fm38(true, (_702_v61).f39, new BigNumber((_596_v1).length), globalState);
          _nw103[(new BigNumber(18)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(19)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(20)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(21)).toNumber()] = _719_v71;
          _nw103[(new BigNumber(22)).toNumber()] = _719_v71;
          _720_v72 = _nw103;
          let _723_v73;
          _723_v73 = _dafny.Seq.of(_720_v72, _720_v72, _720_v72, _720_v72, _720_v72);
          let _724_v74;
          _724_v74 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_702_v61).f39);
          let _725_v75;
          _725_v75 = _dafny.Seq.of(p1);
          let _726_v76;
          _726_v76 = _dafny.Seq.of((_dafny.ZERO).minus((_725_v75)[_module.__default.safeIndex(p1, new BigNumber((_725_v75).length))]));
          let _rhs119 = _723_v73;
          let _rhs120 = (_702_v61).f39;
          let _rhs121 = (_this).f24;
          let _rhs122 = (_dafny.ZERO).minus((p1).multipliedBy(_module.__default.safeModuloInt(new BigNumber((_724_v74).length), (_726_v76)[_module.__default.safeIndex(p1, new BigNumber((_726_v76).length))])));
          let _rhs123 = true;
          let _lhs87 = globalState;
          let _lhs88 = globalState;
          let _lhs89 = globalState;
          _723_v73 = _rhs119;
          _lhs87.f13 = _rhs120;
          _lhs88.f13 = _rhs121;
          r2 = _rhs122;
          _lhs89.f13 = _rhs123;
        }
      }
      let _727_v77;
      let _init17 = function (_728_i6) {
        return _this.f23;
      };
      let _nw104 = Array((new BigNumber(9)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw104.length); _i0_17++) {
        _nw104[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _727_v77 = _nw104;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_727_v77).length))) {
        let _729_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_729_i5)) && ((_729_i5).isLessThan(new BigNumber((_727_v77).length))))) {
          (_727_v77)[(_729_i5)] = _this.f23;
        }
      }
      let _730_v78;
      let _nw105 = new _module.C1();
      _nw105.__ctor((_this).f24, (_this).f24);
      _730_v78 = _nw105;
      let _731_v79;
      let _init18 = function (_732_i7) {
        return (_this).f24;
      };
      let _nw106 = Array((new BigNumber(13)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw106.length); _i0_18++) {
        _nw106[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _731_v79 = _nw106;
      let _733_v80;
      _733_v80 = _dafny.Map.Empty.slice().updateUnsafe(_730_v78,(((_730_v78).f27) ? (_731_v79) : (_731_v79)));
      _733_v80 = ((_733_v80).Merge(_733_v80)).Merge((_733_v80).update(_730_v78, _731_v79));
      if (_dafny.Seq.contains(_594_v0, _this.f23)) {
        let _734_v81;
        _734_v81 = _module.D2.create_DC7((_730_v78).f27, (_730_v78).f27);
        (globalState).f13 = (_734_v81).dtor_cf14;
        let _735_v82;
        _735_v82 = _dafny.Map.Empty.slice().updateUnsafe((_730_v78).f27,_594_v0);
        let _736_v83;
        let _nw107 = new _module.C0();
        _nw107.__ctor(_735_v82, !((_this).f24), (_594_v0)[_module.__default.safeIndex(new BigNumber(-439), new BigNumber((_594_v0).length))], (new BigNumber((_594_v0).length)).isLessThanOrEqualTo(p1));
        _736_v83 = _nw107;
        let _737_v84;
        _737_v84 = _dafny.Map.Empty.slice().updateUnsafe(p1,_596_v1);
        let _738_v85;
        _738_v85 = _module.D6.create_DC15(_737_v84);
        let _pat_let_tv24 = _596_v1;
        let _739_v87;
        let _nw108 = Array((new BigNumber(8)).toNumber());
        _nw108[(_dafny.ZERO).toNumber()] = _738_v85;
        _nw108[(_dafny.ONE).toNumber()] = _module.__default.fm39(p1, globalState);
        _nw108[(new BigNumber(2)).toNumber()] = (((_736_v83).f38) ? (_738_v85) : (_738_v85));
        _nw108[(new BigNumber(3)).toNumber()] = function (_pat_let33_0) {
          return function (_740_dt__update__tmp_h6) {
            return function (_pat_let34_0) {
              return function (_742_dt__update_hcf24_h0) {
                return _module.D6.create_DC15(_742_dt__update_hcf24_h0);
              }(_pat_let34_0);
            }(function () {
              let _coll64 = new _dafny.Map();
              for (const _compr_64 of _dafny.IntegerRange(new BigNumber(36), new BigNumber(-788))) {
                let _741_v86 = _compr_64;
                if (((new BigNumber(36)).isLessThanOrEqualTo(_741_v86)) && ((_741_v86).isLessThan(new BigNumber(-788)))) {
                  _coll64.push([(_741_v86).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("exgcvpsfo")).length))).length)),_pat_let_tv24]);
                }
              }
              return _coll64;
            }());
          }(_pat_let33_0);
        }(_module.__default.fm39(new BigNumber(278), globalState));
        _nw108[(new BigNumber(4)).toNumber()] = _738_v85;
        _nw108[(new BigNumber(5)).toNumber()] = _738_v85;
        _nw108[(new BigNumber(6)).toNumber()] = _module.D6.create_DC15(_737_v84);
        _nw108[(new BigNumber(7)).toNumber()] = _738_v85;
        _739_v87 = _nw108;
        let _index102 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_739_v87).length));
        (_739_v87)[_index102] = _738_v85;
        let _index103 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_739_v87).length));
        (_739_v87)[_index103] = _738_v85;
        let _743_v88;
        _743_v88 = _module.D2.create_DC6((_735_v82).Merge(_735_v82));
        _743_v88 = _743_v88;
        let _744_v89;
        _744_v89 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-678)), function (_745_i8) {
          return new BigNumber(556);
        })).length)));
        (globalState).f13 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(54), p1, p1), _dafny.Seq.of(p1)), _dafny.Seq.Concat(_744_v89, _744_v89));
      } else {
        let _746_v90;
        _746_v90 = _dafny.Map.Empty.slice().updateUnsafe((_730_v78).f27,(_730_v78).f27);
        let _747_v91;
        _747_v91 = _dafny.Seq.of(_746_v90, _746_v90, _746_v90, _746_v90);
        let _748_v92;
        _748_v92 = _dafny.Map.Empty.slice().updateUnsafe(_597_v2,_747_v91);
        r0 = (_dafny.ZERO).minus(new BigNumber(((((_748_v92).contains(_597_v2)) ? ((_748_v92).get(_597_v2)) : (_dafny.Seq.Concat(_747_v91, _dafny.Seq.of(_746_v90))))).length));
        let _749_v93;
        _749_v93 = _module.D0.create_DC2((_this).fm4(p1, _596_v1, globalState), p1, (_730_v78).f27);
        let _750_v94;
        _750_v94 = _dafny.Map.Empty.slice().updateUnsafe(_749_v93,(_this).f24);
        if ((((_750_v94).contains(_749_v93)) ? ((_750_v94).get(_749_v93)) : ((_730_v78).f27))) {
          (_this).f23 = (((_this).f24) ? (_module.__default.fm29(p1, (_730_v78).f27, (_730_v78).f27, p1, globalState)) : (_this.f23));
          (globalState).f7 = (new BigNumber((_594_v0).length)).plus((_dafny.ZERO).minus(p1));
          let _751_v95;
          _751_v95 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm40((_730_v78).f27, globalState)).dtor_cf20,!((_730_v78).f27));
          let _752_v96;
          _752_v96 = _dafny.Seq.of(p1);
          _751_v95 = (_751_v95).update(((_730_v78).fm9(_752_v96, p1, (_730_v78).f27, globalState)).plus(p1), (_730_v78).f27);
          (globalState).f7 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((p1).multipliedBy(new BigNumber((_752_v96).length))), p1);
          let _753_v97;
          _753_v97 = _dafny.Set.fromElements((_730_v78).f27, (_730_v78).f27);
          _751_v95 = (_751_v95).update(new BigNumber((_753_v97).length), _dafny.areEqual(_752_v96, _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vbav")).length), p1, new BigNumber((_753_v97).length))));
        } else {
          let _754_v98;
          _754_v98 = _dafny.Seq.of(_594_v0, _594_v0, _dafny.Seq.update(_dafny.Seq.Concat(_594_v0, _594_v0), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_594_v0, _594_v0)).length)), _this.f23), _dafny.Seq.Concat(_594_v0, _594_v0));
          r1 = (_754_v98)[_module.__default.safeIndex(p1, new BigNumber((_754_v98).length))];
          (globalState).f13 = (_730_v78).f27;
          (globalState).f7 = p1;
          (globalState).f7 = p1;
          let _755_v99;
          _755_v99 = _dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), _this.f23, _this.f23);
          _755_v99 = _dafny.Set.fromElements((function (_pat_let35_0) {
            return function (_756_dt__update__tmp_h7) {
              return function (_pat_let36_0) {
                return function (_757_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4(_757_dt__update_hcf6_h0);
                }(_pat_let36_0);
              }(_this.f23);
            }(_pat_let35_0);
          }(_597_v2)).dtor_cf6, _this.f23);
        }
        if ((_730_v78).f27) {
          let _rhs124 = _this.f23;
          let _rhs125 = (_730_v78).f27;
          let _rhs126 = (_this).f24;
          let _lhs90 = _this;
          let _lhs91 = globalState;
          let _lhs92 = globalState;
          _lhs90.f23 = _rhs124;
          _lhs91.f13 = _rhs125;
          _lhs92.f13 = _rhs126;
          (globalState).f13 = (_730_v78).f27;
          let _758_v100;
          _758_v100 = _module.D2.create_DC7((_730_v78).f27, (_730_v78).f27);
          let _759_v101;
          _759_v101 = _module.D5.create_DC13(new BigNumber(377), !_dafny.areEqual(_596_v1, _596_v1), !((_758_v100).dtor_cf13) || (_module.__default.fm2(globalState)));
          _759_v101 = _759_v101;
          let _nw109 = Array((new BigNumber(16)).toNumber()).fill(false);
          _731_v79 = _nw109;
          let _nw110 = Array((new BigNumber(18)).toNumber()).fill(false);
          _731_v79 = _nw110;
        } else {
          let _760_v102;
          _760_v102 = _dafny.Map.Empty.slice().updateUnsafe(true,_594_v0);
          let _761_v103;
          _761_v103 = _dafny.MultiSet.fromElements(new BigNumber(975));
          let _762_v104;
          let _nw111 = new _module.C0();
          _nw111.__ctor(_760_v102, (_this).fm4(new BigNumber((_761_v103).cardinality()), _596_v1, globalState), _this.f23, (_730_v78).f27);
          _762_v104 = _nw111;
          let _763_v105;
          let _nw112 = Array((new BigNumber(26)).toNumber());
          _763_v105 = _nw112;
          let _764_v106;
          let _nw113 = new _module.C1();
          _nw113.__ctor((_730_v78).f27, (_730_v78).f27);
          _764_v106 = _nw113;
          let _index104 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_763_v105).length));
          (_763_v105)[_index104] = _764_v106;
          let _765_v107;
          _765_v107 = _764_v106;
          let _index105 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_763_v105).length));
          (_763_v105)[_index105] = (_765_v107);
          let _766_v108;
          let _out30;
          _out30 = (_764_v106).m4((_this).f24, globalState);
          _766_v108 = _out30;
          _766_v108 = !((_this).f24);
          (globalState).f13 = ((_this).f24) || (((_this).f24) || ((_730_v78).f27));
        }
        r0 = (new BigNumber((_594_v0).length)).multipliedBy((_this).fm24(p1, (_730_v78).f27, globalState));
        let _767_v109;
        _767_v109 = _dafny.Seq.of(p1, p1, _module.__default.fm0((_730_v78).f27, new BigNumber(395), globalState));
        let _768_v110;
        _768_v110 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_this).f24);
        (globalState).f7 = (new BigNumber(571)).minus((_dafny.ZERO).minus((_767_v109)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_768_v110)).length), new BigNumber((_767_v109).length))]));
      }
      r0 = (((_this).f24) ? ((((_730_v78).f27) ? ((_730_v78).fm9(_dafny.Seq.of(p1, p1), p1, (_this).f24, globalState)) : (p1))) : ((_dafny.ZERO).minus((p1).minus(p1))));
      r1 = _594_v0;
      r2 = _module.__default.safeModuloInt(new BigNumber(406), (((_this).f24) ? (p1) : (p1)));
      return [r0, r1, r2];
    }
    get f36() {
      let _this = this;
      return _this._f36;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
      this._f25 = [];
      this._f35 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f35, f25, f23, f24) {
      let _this = this;
      (_this)._f35 = f35;
      (_this)._f25 = f25;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    fm7(globalState) {
      let _this = this;
      if (((_this).f24) || (false)) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_769_i0) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("er")).length);
        });
      } else {
        return _dafny.Seq.of(new BigNumber(278));
      }
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("n"))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("f")));
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return ((new BigNumber(538)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), function (_770_i0) {
        return _this.f23;
      })).length))).isLessThanOrEqualTo(new BigNumber(((((_this).f24) ? (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("qkrmexiw"))) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("aaucpfakq"), _dafny.Seq.UnicodeFromString("hrrodm"), _dafny.Seq.UnicodeFromString("qgsjiaqub"))))).length));
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _771_v0;
      _771_v0 = new BigNumber(331);
      (globalState).f7 = (_771_v0).minus(_771_v0);
      let _772_v1;
      _772_v1 = _dafny.Seq.UnicodeFromString("gbplvd");
      let _773_v2;
      _773_v2 = _module.D4.create_DC10(_772_v1);
      _772_v1 = _dafny.Seq.Concat(_772_v1, (_773_v2).dtor_cf17);
      let _774_v3;
      _774_v3 = _dafny.MultiSet.fromElements(_771_v0, _771_v0, _771_v0, _771_v0, _771_v0);
      let _775_i0;
      _775_i0 = _dafny.ZERO;
      L2: {
        while (!((_774_v3).IsDisjointFrom((((_this).f24) ? (_774_v3) : (_774_v3))))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_775_i0)) {
              break L2;
            }
            _775_i0 = (_775_i0).plus(_dafny.ONE);
            let _776_v4;
            let _777_v5;
            let _778_v6;
            let _out31;
            let _out32;
            let _out33;
            let _outcollector10 = (_this).m11(p0, globalState);
            _out31 = _outcollector10[0];
            _out32 = _outcollector10[1];
            _out33 = _outcollector10[2];
            _776_v4 = _out31;
            _777_v5 = _out32;
            _778_v6 = _out33;
            let _779_v7;
            _779_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_772_v1);
            let _780_v8;
            let _nw114 = new _module.C0();
            _nw114.__ctor(_779_v7, _778_v6, p1, true);
            _780_v8 = _nw114;
            r1 = _771_v0;
            (globalState).f13 = (_this).f24;
          }
        }
      }
      let _781_v9;
      _781_v9 = _dafny.Map.Empty.slice().updateUnsafe(_772_v1,_771_v0);
      (globalState).f7 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_772_v1,_771_v0)).Merge(_781_v9)).length);
      r0 = p0;
      (globalState).f13 = (_this).f24;
      r0 = p0;
      r1 = _771_v0;
      let _782_v10;
      _782_v10 = _dafny.MultiSet.fromElements((_this).f24, p0);
      let _783_v11;
      _783_v11 = _dafny.Map.Empty.slice().updateUnsafe(_772_v1,_782_v10);
      r2 = !((((_783_v11).contains(_772_v1)) ? ((_783_v11).get(_772_v1)) : (_782_v10))).contains(p0);
      return [r0, r1, r2];
    }
    m11(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let r2 = false;
      let _784_v0;
      _784_v0 = _module.D5.create_DC14(p0);
      r0 = (_784_v0).dtor_cf23;
      let _785_i0;
      _785_i0 = _dafny.ZERO;
      L3: {
        while ((_this).f24) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_785_i0)) {
              break L3;
            }
            _785_i0 = (_785_i0).plus(_dafny.ONE);
            let _786_v1;
            _786_v1 = new BigNumber(49);
            (globalState).f7 = _786_v1;
            let _source13 = _784_v0;
            if (_source13.is_DC13) {
              let _787___mcc_h0 = (_source13).cf20;
              let _788___mcc_h1 = (_source13).cf21;
              let _789___mcc_h2 = (_source13).cf22;
              let _790_cf22 = _789___mcc_h2;
              let _791_cf21 = _788___mcc_h1;
              let _792_cf20 = _787___mcc_h0;
              let _793_v2;
              _793_v2 = _dafny.Seq.of(_786_v1, (_dafny.ZERO).minus(_786_v1));
              let _794_v3;
              _794_v3 = _dafny.Seq.of(_dafny.Seq.of(_786_v1, new BigNumber(-258), _786_v1, _786_v1, _792_cf20));
              (globalState).f7 = new BigNumber((_dafny.Seq.Concat(_793_v2, (_794_v3)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(690)), function (_795_i1) {
                return _this.f23;
              })).length), new BigNumber((_794_v3).length))])).length);
              (_this).f23 = _this.f23;
              let _796_v4;
              let _nw115 = new _module.C1();
              _nw115.__ctor(_790_cf22, p0);
              _796_v4 = _nw115;
              (globalState).f7 = _786_v1;
            } else if (_source13.is_DC14) {
              let _797___mcc_h3 = (_source13).cf23;
              let _798_cf23 = _797___mcc_h3;
              let _799_v5;
              let _nw116 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
              _799_v5 = _nw116;
              (globalState).f22 = _799_v5;
              let _800_v6;
              _800_v6 = _dafny.Set.fromElements(p0);
              _800_v6 = _module.__default.fm41(globalState);
              let _801_v7;
              _801_v7 = _dafny.Seq.of((_this).f24, p0, false, !(_798_cf23));
              let _802_v8;
              _802_v8 = _dafny.Map.Empty.slice().updateUnsafe(_801_v7,_786_v1);
              let _803_v9;
              _803_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm36(globalState));
              let _804_v10;
              let _init19 = function (_805_i2) {
                return false;
              };
              let _nw117 = Array((new BigNumber(4)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw117.length); _i0_19++) {
                _nw117[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _804_v10 = _nw117;
              let _806_v11;
              _806_v11 = _dafny.Map.Empty.slice().updateUnsafe(_804_v10,_786_v1);
              _802_v8 = (_802_v8).update(_dafny.Seq.Concat(_801_v7, (((_803_v9).contains(p0)) ? ((_803_v9).get(p0)) : (_801_v7))), (((_806_v11).contains(_804_v10)) ? ((_806_v11).get(_804_v10)) : (_786_v1)));
              _798_cf23 = p0;
            } else {
              let _807___mcc_h4 = (_source13).cf19;
              let _808_cf19 = _807___mcc_h4;
              let _809_v12;
              _809_v12 = _dafny.Set.fromElements(_this.f23, _this.f23, _this.f23);
              let _810_v13;
              _810_v13 = _dafny.Seq.of(_786_v1);
              let _811_v14;
              _811_v14 = _dafny.Seq.UnicodeFromString("cemw");
              let _812_v15;
              _812_v15 = _dafny.Map.Empty.slice().updateUnsafe(_786_v1,(_dafny.ZERO).minus((_810_v13)[_module.__default.safeIndex(_module.__default.fm0(p0, new BigNumber((_811_v14).length), globalState), new BigNumber((_810_v13).length))]));
              let _813_v16;
              _813_v16 = _dafny.Map.Empty.slice().updateUnsafe((_809_v12).IsProperSubsetOf(_809_v12),new BigNumber((_812_v15).length));
              _813_v16 = (_813_v16).update(p0, _786_v1);
              let _814_v18;
              _814_v18 = (_812_v15).Merge(function () {
                let _coll65 = new _dafny.Map();
                for (const _compr_65 of (_810_v13).Elements) {
                  let _815_v17 = _compr_65;
                  if (_dafny.Seq.contains(_810_v13, _815_v17)) {
                    _coll65.push([_module.__default.safeDivisionInt(_815_v17, _786_v1),_786_v1]);
                  }
                }
                return _coll65;
              }());
              _814_v18 = _812_v15;
              let _816_v19;
              _816_v19 = _dafny.Seq.of(p0);
              let _817_v20;
              _817_v20 = _module.D0.create_DC0(_dafny.MultiSet.FromArray(_816_v19));
              let _818_v21;
              _818_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
              let _819_v22;
              let _nw118 = Array((new BigNumber(26)).toNumber());
              _nw118[(_dafny.ZERO).toNumber()] = p0;
              _nw118[(_dafny.ONE).toNumber()] = (_786_v1).isLessThan(_module.__default.fm0(!((_this).f24), _786_v1, globalState));
              _nw118[(new BigNumber(2)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_816_v19, _816_v19);
              _nw118[(new BigNumber(4)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(5)).toNumber()] = p0;
              _nw118[(new BigNumber(6)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(7)).toNumber()] = !(false);
              _nw118[(new BigNumber(8)).toNumber()] = (p0) || (p0);
              _nw118[(new BigNumber(9)).toNumber()] = p0;
              _nw118[(new BigNumber(10)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(11)).toNumber()] = p0;
              _nw118[(new BigNumber(12)).toNumber()] = ((_817_v20).dtor_cf0).IsDisjointFrom(_dafny.MultiSet.fromElements(p0, (_this).f24));
              _nw118[(new BigNumber(13)).toNumber()] = (new BigNumber((_818_v21).length)).isLessThanOrEqualTo(new BigNumber((_818_v21).length));
              _nw118[(new BigNumber(14)).toNumber()] = p0;
              _nw118[(new BigNumber(15)).toNumber()] = _module.__default.fm2(globalState);
              _nw118[(new BigNumber(16)).toNumber()] = p0;
              _nw118[(new BigNumber(17)).toNumber()] = p0;
              _nw118[(new BigNumber(18)).toNumber()] = (_816_v19)[_module.__default.safeIndex(_786_v1, new BigNumber((_816_v19).length))];
              _nw118[(new BigNumber(19)).toNumber()] = p0;
              _nw118[(new BigNumber(20)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(21)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(22)).toNumber()] = (new BigNumber((_811_v14).length)).isEqualTo(_786_v1);
              _nw118[(new BigNumber(23)).toNumber()] = p0;
              _nw118[(new BigNumber(24)).toNumber()] = (_this).f24;
              _nw118[(new BigNumber(25)).toNumber()] = (_this).f24;
              _819_v22 = _nw118;
              let _820_v23;
              _820_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-580),(_this).f24);
              let _index106 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_819_v22).length));
              (_819_v22)[_index106] = (((_820_v23).contains(new BigNumber((_810_v13).length))) ? ((_820_v23).get(new BigNumber((_810_v13).length))) : (!((_this).f24)));
              let _index107 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_819_v22).length));
              (_819_v22)[_index107] = (_this).f24;
              let _821_v24;
              let _nw119 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
              _821_v24 = _nw119;
              let _index108 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_821_v24).length));
              (_821_v24)[_index108] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_786_v1, _786_v1));
              let _index109 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_821_v24).length));
              (_821_v24)[_index109] = _786_v1;
            }
            let _822_v25;
            let _init20 = ((_823_v1) => function (_824_i3) {
              return _module.D4.create_DC11(_823_v1);
            })(_786_v1);
            let _nw120 = Array((_dafny.ONE).toNumber());
            for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw120.length); _i0_20++) {
              _nw120[_i0_20] = _init20(new BigNumber(_i0_20));
            }
            _822_v25 = _nw120;
            let _825_v26;
            _825_v26 = _module.D4.create_DC11(new BigNumber(911));
            let _index110 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_822_v25).length));
            (_822_v25)[_index110] = _825_v26;
            let _index111 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_822_v25).length));
            (_822_v25)[_index111] = _825_v26;
            let _826_v27;
            let _nw121 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
            _826_v27 = _nw121;
            let _index112 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_826_v27).length));
            (_826_v27)[_index112] = (_dafny.ZERO).minus(_786_v1);
            let _index113 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_826_v27).length));
            (_826_v27)[_index113] = (_786_v1).multipliedBy(_786_v1);
          }
        }
      }
      (globalState).f13 = _module.__default.fm2(globalState);
      let _827_v28;
      _827_v28 = _dafny.Seq.UnicodeFromString("plwsf");
      let _828_v29;
      _828_v29 = new BigNumber(669);
      let _829_v30;
      let _nw122 = Array((new BigNumber(24)).toNumber());
      _nw122[(_dafny.ZERO).toNumber()] = _827_v28;
      _nw122[(_dafny.ONE).toNumber()] = _827_v28;
      _nw122[(new BigNumber(2)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(3)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(4)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(5)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(6)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(7)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(8)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(774)), function (_830_i4) {
        return _this.f23;
      });
      _nw122[(new BigNumber(10)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(11)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(12)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_827_v28, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), function (_831_i5) {
        return _this.f23;
      }), _module.__default.safeIndex(_828_v29, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), function (_832_i5) {
        return _this.f23;
      })).length)), _this.f23));
      _nw122[(new BigNumber(14)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_827_v28, _827_v28);
      _nw122[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(439)), function (_833_i6) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      });
      _nw122[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(769)), function (_834_i7) {
        return _this.f23;
      });
      _nw122[(new BigNumber(18)).toNumber()] = (((_this).f24) ? (_dafny.Seq.UnicodeFromString("lti")) : (_827_v28));
      _nw122[(new BigNumber(19)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(20)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(21)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(22)).toNumber()] = _827_v28;
      _nw122[(new BigNumber(23)).toNumber()] = _827_v28;
      _829_v30 = _nw122;
      let _835_v31;
      _835_v31 = _dafny.Seq.of(p0);
      let _836_v32;
      _836_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_835_v31).length),p0);
      let _837_v33;
      _837_v33 = _dafny.Set.fromElements((((_836_v32).contains(_module.__default.fm0((_this).f24, _828_v29, globalState))) ? ((_836_v32).get(_module.__default.fm0((_this).f24, _828_v29, globalState))) : ((_this).f24)), true, (_this).f24);
      let _pat_let_tv25 = _828_v29;
      let _pat_let_tv26 = _828_v29;
      let _rhs127 = _829_v30;
      let _rhs128 = !(p0);
      let _rhs129 = function (_source14) {
        if (_source14.is_DC13) {
          let _838___mcc_h5 = (_source14).cf20;
          let _839___mcc_h6 = (_source14).cf21;
          let _840___mcc_h7 = (_source14).cf22;
          let _841_cf22 = _840___mcc_h7;
          let _842_cf21 = _839___mcc_h6;
          let _843_cf20 = _838___mcc_h5;
          return _843_cf20;
        } else if (_source14.is_DC14) {
          let _844___mcc_h8 = (_source14).cf23;
          let _845_cf23 = _844___mcc_h8;
          return _pat_let_tv25;
        } else {
          let _846___mcc_h9 = (_source14).cf19;
          let _847_cf19 = _846___mcc_h9;
          return _pat_let_tv26;
        }
      }(_module.D5.create_DC13((_dafny.ZERO).minus(new BigNumber((_837_v33).length)), p0, false));
      let _rhs130 = p0;
      let _rhs131 = (_this).f24;
      let _lhs93 = globalState;
      let _lhs94 = globalState;
      let _lhs95 = globalState;
      _829_v30 = _rhs127;
      _lhs93.f13 = _rhs128;
      _lhs94.f7 = _rhs129;
      _lhs95.f13 = _rhs130;
      r0 = _rhs131;
      let _848_v34;
      let _init21 = ((_849_p0) => function (_850_i9) {
        return _849_p0;
      })(p0);
      let _nw123 = Array((new BigNumber(28)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw123.length); _i0_21++) {
        _nw123[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _848_v34 = _nw123;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_848_v34).length))) {
        let _851_i8 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_851_i8)) && ((_851_i8).isLessThan(new BigNumber((_848_v34).length))))) {
          (_848_v34)[(_851_i8)] = (_dafny.MultiSet.fromElements((_this).f24, p0)).IsSubsetOf(_dafny.MultiSet.fromElements(false, (_this).f24));
        }
      }
      (globalState).f13 = (_this).f24;
      r0 = p0;
      r1 = _module.__default.fm34(((true) ? (_828_v29) : (_828_v29)), !(false), globalState);
      r2 = (_this).f24;
      return [r0, r1, r2];
    }
    get f35() {
      let _this = this;
      return _this._f35;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
      this._f25 = [];
      this.f34 = _dafny.Map.Empty;
      this._f33 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f33, f34, f25, f23, f24) {
      let _this = this;
      (_this)._f33 = f33;
      (_this).f34 = f34;
      (_this)._f25 = f25;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return _dafny.Seq.of(new BigNumber(207), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(-927))).length)).multipliedBy(new BigNumber((function () {
        let _coll66 = new _dafny.Map();
        for (const _compr_66 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,false), _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24))).Elements) {
          let _852_v0 = _compr_66;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,false), _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24))).contains(_852_v0)) {
            _coll66.push([_852_v0,(_this).f33]);
          }
        }
        return _coll66;
      }()).length)));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return ((((_this).f24) ? (_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("h")))) : (_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("d")))))).dtor_cf12;
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_this).f24;
    };
    fm22(p0, p1, globalState) {
      let _this = this;
      return (_this).f24;
    };
    fm23(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f24;
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      if (p0) {
        let _853_v0;
        _853_v0 = _dafny.Seq.of(!(p0), p0);
        let _854_v1;
        _854_v1 = _dafny.MultiSet.fromElements(!((_this).f24), p0, (_853_v0)[_module.__default.safeIndex((_this).f33, new BigNumber((_853_v0).length))]);
        let _855_v2;
        let _nw124 = new _module.C2();
        _nw124.__ctor((_854_v1).Intersect(_dafny.MultiSet.fromElements((_this).f24, (_this).f24)), (_this).f25, new _dafny.CodePoint('c'.codePointAt(0)), _module.__default.fm2(globalState));
        _855_v2 = _nw124;
        let _856_v3;
        _856_v3 = _dafny.Set.fromElements(new BigNumber(315));
        (globalState).f13 = (_856_v3).IsSubsetOf(_856_v3);
        let _857_v4;
        let _nw125 = new _module.C1();
        _nw125.__ctor(true, (_this).f24);
        _857_v4 = _nw125;
        let _858_v5;
        _858_v5 = _dafny.Seq.UnicodeFromString("c");
        let _859_v6;
        _859_v6 = _dafny.Map.Empty.slice().updateUnsafe((_857_v4).f39,_858_v5);
        let _source15 = _module.D2.create_DC6(_859_v6);
        if (_source15.is_DC7) {
          let _860___mcc_h0 = (_source15).cf13;
          let _861___mcc_h1 = (_source15).cf14;
          let _862_cf14 = _861___mcc_h1;
          let _863_cf13 = _860___mcc_h0;
          let _864_v7;
          _864_v7 = _dafny.Seq.of((_this).f33);
          let _865_v8;
          _865_v8 = _dafny.Seq.of(_864_v7);
          let _866_v9;
          _866_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(43),new BigNumber((_865_v8).length));
          let _867_v11;
          _867_v11 = _dafny.Map.Empty.slice().updateUnsafe((_857_v4).f39,p0);
          let _868_v12;
          _868_v12 = _dafny.Set.fromElements(_module.D6.create_DC16(_867_v11, (_857_v4).f39, (_857_v4).f39));
          let _rhs132 = _862_cf14;
          let _rhs133 = _dafny.Seq.of(_862_cf14, (_857_v4).f39);
          let _rhs134 = function () {
            let _coll67 = new _dafny.Map();
            for (const _compr_67 of _dafny.IntegerRange(new BigNumber(452), new BigNumber(718))) {
              let _869_v10 = _compr_67;
              if (((new BigNumber(452)).isLessThanOrEqualTo(_869_v10)) && ((_869_v10).isLessThan(new BigNumber(718)))) {
                _coll67.push([_module.__default.safeDivisionInt(_869_v10, (_this).f33),(new BigNumber((_dafny.Set.fromElements(_863_cf13, (_857_v4).f39)).length)).plus(new BigNumber((_856_v3).length))]);
              }
            }
            return _coll67;
          }();
          let _rhs135 = ((_this).f33).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_868_v12)).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_870_i0) {
            return (_this).f33;
          })).length)));
          let _lhs96 = globalState;
          _863_cf13 = _rhs132;
          _853_v0 = _rhs133;
          _866_v9 = _rhs134;
          _lhs96.f7 = _rhs135;
          let _871_v13;
          _871_v13 = _module.D5.create_DC13((_this).f33, (_857_v4).f39, (_857_v4).f39);
          let _872_v14;
          _872_v14 = _dafny.MultiSet.fromElements(_871_v13, _871_v13);
          let _873_v15;
          _873_v15 = _dafny.MultiSet.fromElements(_858_v5);
          let _874_v16;
          let _init22 = ((_875_v0) => function (_876_i1) {
            return _875_v0;
          })(_853_v0);
          let _nw126 = Array((new BigNumber(5)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw126.length); _i0_22++) {
            _nw126[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _874_v16 = _nw126;
          let _index114 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_874_v16).length));
          (_874_v16)[_index114] = (_855_v2).fm25(_863_cf13, globalState);
          let _877_v17;
          _877_v17 = _dafny.MultiSet.fromElements((_this).f33);
          let _878_v18;
          _878_v18 = _dafny.Set.fromElements(_877_v17, _877_v17, _module.__default.fm42((_this).f24, new BigNumber((_858_v5).length), globalState));
          let _index115 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_874_v16).length));
          let _rhs136 = (_872_v14).Union(_872_v14);
          let _rhs137 = _873_v15;
          let _rhs138 = (_857_v4).fm32(_863_cf13, (_this).f33, globalState);
          let _rhs139 = !((_878_v18).IsSubsetOf(function () {
            let _coll68 = new _dafny.Set();
            for (const _compr_68 of (_dafny.Map.Empty.slice().updateUnsafe(_877_v17,(_this).f33)).Keys.Elements) {
              let _879_v19 = _compr_68;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_877_v17,(_this).f33)).contains(_879_v19)) {
                _coll68.add(_879_v19);
              }
            }
            return _coll68;
          }()));
          let _rhs140 = _853_v0;
          let _lhs97 = globalState;
          let _lhs98 = globalState;
          let _lhs99 = _874_v16;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_874_v16).length));
          _872_v14 = _rhs136;
          _873_v15 = _rhs137;
          _lhs97.f13 = _rhs138;
          _lhs98.f13 = _rhs139;
          _lhs99[_lhs100] = _rhs140;
          let _880_v20;
          let _nw127 = Array((new BigNumber(20)).toNumber()).fill(false);
          _880_v20 = _nw127;
          let _index116 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_880_v20).length));
          (_880_v20)[_index116] = (_857_v4).f39;
          let _index117 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_880_v20).length));
          let _rhs141 = _module.__default.safeDivisionInt((_this).f33, (_dafny.ZERO).minus(((_dafny.ZERO).minus((_this).f33)).minus((_this).f33)));
          let _rhs142 = (_857_v4).f39;
          let _lhs101 = globalState;
          let _lhs102 = _880_v20;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_880_v20).length));
          _lhs101.f7 = _rhs141;
          _lhs102[_lhs103] = _rhs142;
          r0 = true;
        } else if (_source15.is_DC6) {
          let _881___mcc_h2 = (_source15).cf12;
          let _882_cf12 = _881___mcc_h2;
          r2 = p0;
          let _883_v21;
          let _nw128 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _883_v21 = _nw128;
          let _884_v23;
          _884_v23 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
            let _coll69 = new _dafny.Set();
            for (const _compr_69 of _dafny.IntegerRange(new BigNumber(643), new BigNumber(65))) {
              let _885_v22 = _compr_69;
              if (((new BigNumber(643)).isLessThanOrEqualTo(_885_v22)) && ((_885_v22).isLessThan(new BigNumber(65)))) {
                _coll69.add((_885_v22).minus((_this).f33));
              }
            }
            return _coll69;
          }()).length));
          let _886_v24;
          _886_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_884_v23).length),_854_v1);
          let _index118 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_883_v21).length));
          (_883_v21)[_index118] = _886_v24;
          let _index119 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_883_v21).length));
          (_883_v21)[_index119] = _886_v24;
          let _887_v25;
          _887_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(148),p0);
          (globalState).f13 = (_853_v0)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f33, new BigNumber((_887_v25).length)), new BigNumber((_853_v0).length))];
          r2 = !(_module.__default.fm2(globalState));
        } else {
          let _888___mcc_h3 = (_source15).cf15;
          let _889_cf15 = _888___mcc_h3;
          r2 = (_857_v4).f39;
          let _890_v27;
          _890_v27 = _module.D4.create_DC11(new BigNumber((function () {
  let _coll70 = new _dafny.Map();
  for (const _compr_70 of (_dafny.Seq.of(_this.f23, p1)).Elements) {
    let _891_v26 = _compr_70;
    if (_dafny.Seq.contains(_dafny.Seq.of(_this.f23, p1), _891_v26)) {
      _coll70.push([_891_v26,(_857_v4).f39]);
    }
  }
  return _coll70;
}()).length));
          let _892_v28;
          _892_v28 = _module.D0.create_DC2((_857_v4).f39, (_890_v27).dtor_cf18, (_857_v4).f39);
          let _pat_let_tv27 = p0;
          let _pat_let_tv28 = _857_v4;
          (globalState).f7 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_858_v5, _module.__default.fm1(true, globalState)),function (_pat_let37_0) {
            return function (_896_dt__update__tmp_h1) {
              return function (_pat_let41_0) {
                return function (_897_dt__update_hcf4_h1) {
                  return function (_pat_let42_0) {
                    return function (_898_dt__update_hcf2_h0) {
                      return _module.D0.create_DC2(_898_dt__update_hcf2_h0, (_896_dt__update__tmp_h1).dtor_cf3, _897_dt__update_hcf4_h1);
                    }(_pat_let42_0);
                  }((_pat_let_tv28).f39);
                }(_pat_let41_0);
              }((_this).f24);
            }(_pat_let37_0);
          }(function (_pat_let38_0) {
            return function (_893_dt__update__tmp_h0) {
              return function (_pat_let39_0) {
                return function (_894_dt__update_hcf4_h0) {
                  return function (_pat_let40_0) {
                    return function (_895_dt__update_hcf3_h0) {
                      return _module.D0.create_DC2((_893_dt__update__tmp_h0).dtor_cf2, _895_dt__update_hcf3_h0, _894_dt__update_hcf4_h0);
                    }(_pat_let40_0);
                  }(new BigNumber(795));
                }(_pat_let39_0);
              }(_pat_let_tv27);
            }(_pat_let38_0);
          }(_892_v28)))).length);
          r0 = (_857_v4).f39;
          let _899_v29;
          let _nw129 = new _module.C2();
          _nw129.__ctor((_854_v1).update((_857_v4).f39, _module.__default.abs((_this).f33)), (_this).f25, new _dafny.CodePoint('f'.codePointAt(0)), _dafny.Seq.IsProperPrefixOf(_858_v5, _dafny.Seq.UnicodeFromString("i")));
          _899_v29 = _nw129;
        }
        r0 = (_857_v4).f39;
      } else {
        let _900_v30;
        let _nw130 = Array((new BigNumber(16)).toNumber()).fill(false);
        _900_v30 = _nw130;
        let _901_v31;
        let _nw131 = Array((new BigNumber(9)).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = _900_v30;
        _nw131[(_dafny.ONE).toNumber()] = _900_v30;
        _nw131[(new BigNumber(2)).toNumber()] = _900_v30;
        _nw131[(new BigNumber(3)).toNumber()] = _900_v30;
        _nw131[(new BigNumber(4)).toNumber()] = _900_v30;
        _nw131[(new BigNumber(5)).toNumber()] = _900_v30;
        _nw131[(new BigNumber(6)).toNumber()] = _900_v30;
        _nw131[(new BigNumber(7)).toNumber()] = _900_v30;
        _nw131[(new BigNumber(8)).toNumber()] = _900_v30;
        _901_v31 = _nw131;
        _901_v31 = _901_v31;
        let _902_v32;
        let _nw132 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _902_v32 = _nw132;
        let _903_v33;
        _903_v33 = _dafny.Seq.UnicodeFromString("hhp");
        let _904_v34;
        _904_v34 = _dafny.MultiSet.fromElements((_this).f33, new BigNumber((_903_v33).length));
        let _905_v35;
        _905_v35 = _dafny.MultiSet.fromElements(!(p0), (_this).f24, (_this).f24, (_this).f24);
        let _rhs143 = _902_v32;
        let _rhs144 = (_this).f33;
        let _rhs145 = _900_v30;
        let _rhs146 = _module.__default.safeModuloInt(((((_904_v34).contains((_this).f33)) ? ((_904_v34).get((_this).f33)) : (new BigNumber(-963)))).multipliedBy(new BigNumber((_905_v35).cardinality())), (_this).f33);
        let _lhs104 = globalState;
        _lhs104.f22 = _rhs143;
        r1 = _rhs144;
        _900_v30 = _rhs145;
        r1 = _rhs146;
        let _906_v36;
        _906_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f33);
        _906_v36 = (_906_v36).update(p1, new BigNumber(-924));
        let _index120 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_902_v32).length));
        (_902_v32)[_index120] = (_this).f33;
        let _index121 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_902_v32).length));
        (_902_v32)[_index121] = new BigNumber((_dafny.Seq.Concat(_903_v33, _dafny.Seq.UnicodeFromString("fb"))).length);
        let _907_v37;
        _907_v37 = _dafny.Map.Empty.slice().updateUnsafe(((!(false)) ? (p0) : ((_this).f24)),(_this).f33);
        let _rhs147 = p1;
        let _rhs148 = _907_v37;
        let _rhs149 = ((_this).f33).plus((_this).f33);
        let _rhs150 = p0;
        let _rhs151 = (_this).f24;
        let _lhs105 = _this;
        let _lhs106 = globalState;
        let _lhs107 = globalState;
        _lhs105.f23 = _rhs147;
        _907_v37 = _rhs148;
        _lhs106.f7 = _rhs149;
        r2 = _rhs150;
        _lhs107.f13 = _rhs151;
      }
      if ((p0) || (p0)) {
        let _908_v38;
        _908_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f24);
        let _909_v39;
        _909_v39 = _dafny.Seq.of(_908_v38, _908_v38);
        r0 = _dafny.areEqual(_dafny.Seq.Concat(_909_v39, _909_v39), _909_v39);
        let _910_v40;
        _910_v40 = _dafny.Seq.of((_this).f33);
        let _911_v41;
        _911_v41 = _dafny.MultiSet.fromElements(_910_v40);
        let _912_v42;
        let _nw133 = Array((new BigNumber(2)).toNumber());
        _nw133[(_dafny.ZERO).toNumber()] = (_this).f33;
        _nw133[(_dafny.ONE).toNumber()] = (_this).f33;
        _912_v42 = _nw133;
        let _913_v43;
        _913_v43 = _dafny.Map.Empty.slice().updateUnsafe(_911_v41,_912_v42);
        _913_v43 = (_913_v43).update(_911_v41, (((_this).f24) ? (_912_v42) : (_912_v42)));
        let _914_v44;
        _914_v44 = _dafny.Seq.UnicodeFromString("jp");
        let _915_v45;
        _915_v45 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_916_i2) {
          return _this.f23;
        }));
        let _917_v46;
        _917_v46 = _dafny.Map.Empty.slice().updateUnsafe(_914_v44,_915_v45);
        _917_v46 = (_917_v46).update(_914_v44, _915_v45);
        let _918_v47;
        _918_v47 = _dafny.Seq.of(false, p0, !((_this).f24), (_this).f24);
        let _919_v48;
        _919_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
        let _920_v49;
        let _nw134 = Array((new BigNumber(28)).toNumber());
        _nw134[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw134[(_dafny.ONE).toNumber()] = p0;
        _nw134[(new BigNumber(2)).toNumber()] = p0;
        _nw134[(new BigNumber(3)).toNumber()] = p0;
        _nw134[(new BigNumber(4)).toNumber()] = p0;
        _nw134[(new BigNumber(5)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(6)).toNumber()] = p0;
        _nw134[(new BigNumber(7)).toNumber()] = p0;
        _nw134[(new BigNumber(8)).toNumber()] = p0;
        _nw134[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(10)).toNumber()] = p0;
        _nw134[(new BigNumber(11)).toNumber()] = !((_918_v47)[_module.__default.safeIndex((_this).f33, new BigNumber((_918_v47).length))]);
        _nw134[(new BigNumber(12)).toNumber()] = p0;
        _nw134[(new BigNumber(13)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(14)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(15)).toNumber()] = false;
        _nw134[(new BigNumber(16)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(17)).toNumber()] = p0;
        _nw134[(new BigNumber(18)).toNumber()] = true;
        _nw134[(new BigNumber(19)).toNumber()] = (((_919_v48).contains(true)) ? ((_919_v48).get(true)) : (true));
        _nw134[(new BigNumber(20)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(21)).toNumber()] = (_918_v47)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f33), new BigNumber((_918_v47).length))];
        _nw134[(new BigNumber(22)).toNumber()] = (_this).fm4((_this).f33, _918_v47, globalState);
        _nw134[(new BigNumber(23)).toNumber()] = p0;
        _nw134[(new BigNumber(24)).toNumber()] = false;
        _nw134[(new BigNumber(25)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(26)).toNumber()] = (_this).f24;
        _nw134[(new BigNumber(27)).toNumber()] = true;
        _920_v49 = _nw134;
        let _921_v50;
        _921_v50 = _dafny.Map.Empty.slice().updateUnsafe(_920_v49,_this.f23);
        let _922_v51;
        _922_v51 = _dafny.Set.fromElements((_914_v44)[_module.__default.safeIndex((_this).f33, new BigNumber((_914_v44).length))], _this.f23);
        let _923_v52;
        _923_v52 = _dafny.MultiSet.fromElements((_this).f24, p0);
        let _924_v53;
        _924_v53 = _module.D0.create_DC0(_923_v52);
        let _925_v54;
        _925_v54 = _module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC3(_924_v53)));
        let _926_v55;
        let _nw135 = Array((new BigNumber(8)).toNumber());
        _nw135[(_dafny.ZERO).toNumber()] = _module.D0.create_DC3(_924_v53);
        _nw135[(_dafny.ONE).toNumber()] = _925_v54;
        _nw135[(new BigNumber(2)).toNumber()] = _925_v54;
        _nw135[(new BigNumber(3)).toNumber()] = _925_v54;
        _nw135[(new BigNumber(4)).toNumber()] = _925_v54;
        _nw135[(new BigNumber(5)).toNumber()] = _module.D0.create_DC3(_924_v53);
        _nw135[(new BigNumber(6)).toNumber()] = _925_v54;
        _nw135[(new BigNumber(7)).toNumber()] = _925_v54;
        _926_v55 = _nw135;
        let _927_v56;
        _927_v56 = _dafny.Map.Empty.slice().updateUnsafe(_926_v55,_dafny.Set.fromElements(_this.f23, p1));
        let _928_v57;
        let _nw136 = new _module.C3();
        _nw136.__ctor(_921_v50, (_this).f25, _this.f23, !(((((_927_v56).contains(_926_v55)) ? ((_927_v56).get(_926_v55)) : (_922_v51))).IsProperSubsetOf(_922_v51)));
        _928_v57 = _nw136;
        let _nw137 = new _module.C3();
        _nw137.__ctor((_928_v57).f35, (_this).f25, _this.f23, (new BigNumber((_dafny.Set.fromElements(_912_v42, _912_v42)).length)).isEqualTo((_dafny.ZERO).minus((_this).f33)));
        _928_v57 = _nw137;
        let _929_v58;
        _929_v58 = _dafny.Set.fromElements((_this).f33, new BigNumber(((_928_v57).fm7(globalState)).length));
        let _930_v59;
        _930_v59 = _module.D5.create_DC12(_929_v58);
        let _931_v60;
        _931_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,_this.f23);
        r0 = !(((_929_v58).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_919_v48).length)), (_this).f33, (_this).f33, (_this).f33, new BigNumber((_931_v60).length)))).IsSubsetOf((_930_v59).dtor_cf19));
      } else {
        (globalState).f7 = (_this).f33;
        let _932_v61;
        _932_v61 = _dafny.MultiSet.fromElements(!((_this).f24));
        let _933_v62;
        let _nw138 = new _module.C2();
        _nw138.__ctor((_dafny.MultiSet.fromElements((_this).f24, p0, (_this).f24)).Intersect(_932_v61), (_this).f25, ((p0) ? (_this.f23) : (_this.f23)), (_932_v61).IsDisjointFrom(_932_v61));
        _933_v62 = _nw138;
        let _934_v63;
        let _nw139 = Array((new BigNumber(22)).toNumber()).fill(false);
        _934_v63 = _nw139;
        let _index122 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_934_v63).length));
        (_934_v63)[_index122] = p0;
        let _index123 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_934_v63).length));
        (_934_v63)[_index123] = ((_this).f33).isEqualTo((_dafny.ZERO).minus((_this).f33));
        let _935_v64;
        _935_v64 = _dafny.Seq.of(p0);
        (globalState).f13 = (_935_v64)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f33, (_this).f33), new BigNumber((_935_v64).length))];
        let _936_v65;
        let _out34;
        _out34 = (_this).m10((_934_v63)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_934_v63).length))], globalState);
        _936_v65 = _out34;
      }
      (globalState).f13 = _module.__default.fm2(globalState);
      let _937_v66;
      _937_v66 = _dafny.Seq.of((_this).f24, (_this).f24, p0, !(p0), p0);
      let _938_v67;
      _938_v67 = _dafny.Map.Empty.slice().updateUnsafe(p0,_937_v66);
      (globalState).f13 = !((_938_v67).contains((_this).f24));
      let _hi3 = _module.__default.fm0((_this).f24, (_this).f33, globalState);
      for (let _939_i3 = (_this).f33; _939_i3.isLessThan(_hi3); _939_i3 = _939_i3.plus(_dafny.ONE)) {
        if ((_939_i3).isLessThan(_939_i3)) {
          (globalState).f7 = _939_i3;
          let _940_v68;
          _940_v68 = _dafny.Set.fromElements((_this).f24, (_this).f24);
          let _941_v69;
          _941_v69 = _dafny.Seq.of((_940_v68).Difference(_940_v68));
          let _942_v70;
          _942_v70 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_dafny.Seq.Concat(_dafny.Seq.of(_940_v68, _940_v68), _941_v69));
          let _943_v71;
          _943_v71 = _dafny.Set.fromElements(new BigNumber(999), _939_i3);
          let _944_v72;
          _944_v72 = _dafny.MultiSet.fromElements(_943_v71);
          let _rhs152 = !(!((_this).f24));
          let _rhs153 = (((_942_v70).contains(false)) ? ((_942_v70).get(false)) : (_941_v69));
          let _rhs154 = (_dafny.ZERO).minus(_939_i3);
          let _rhs155 = (_this).f24;
          let _rhs156 = (new BigNumber((_944_v72).cardinality())).multipliedBy((_this).f33);
          let _lhs108 = globalState;
          let _lhs109 = globalState;
          let _lhs110 = globalState;
          let _lhs111 = globalState;
          _lhs108.f13 = _rhs152;
          _941_v69 = _rhs153;
          _lhs109.f7 = _rhs154;
          _lhs110.f13 = _rhs155;
          _lhs111.f7 = _rhs156;
          (_this).f34 = _this.f34;
          let _945_v73;
          let _nw140 = Array((new BigNumber(4)).toNumber()).fill(false);
          _945_v73 = _nw140;
          let _946_v74;
          _946_v74 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f24);
          let _index124 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_945_v73).length));
          (_945_v73)[_index124] = (_this).fm23((_this).f33, _946_v74, _937_v66, globalState);
          let _index125 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_945_v73).length));
          (_945_v73)[_index125] = (_this).fm23(_939_i3, _946_v74, _dafny.Seq.Concat(_937_v66, _937_v66), globalState);
          let _947_v75;
          _947_v75 = _dafny.Seq.of((_this).f33, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f24)).length)));
          let _948_v76;
          _948_v76 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24, false);
          let _949_v77;
          _949_v77 = _dafny.Map.Empty.slice().updateUnsafe(_947_v75,((((_948_v76).contains(false)) ? ((_948_v76).get(false)) : ((_this).f33))).minus(_939_i3));
          _949_v77 = (_949_v77).update(_947_v75, _939_i3);
        } else {
          let _950_v78;
          _950_v78 = _dafny.Seq.UnicodeFromString("jctwmpljq");
          _950_v78 = _dafny.Seq.Concat(_950_v78, _dafny.Seq.update(_950_v78, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p0, p0, (_this).f24, true, p0)).length)), new BigNumber((_950_v78).length)), _this.f23));
          let _951_v79;
          let _nw141 = Array((new BigNumber(9)).toNumber()).fill(false);
          _951_v79 = _nw141;
          _951_v79 = ((p0) ? (_951_v79) : (_951_v79));
          let _952_v80;
          let _init23 = ((_953_p0) => function (_954_i4) {
            return _module.D0.create_DC0(_dafny.MultiSet.fromElements(!(_953_p0)));
          })(p0);
          let _nw142 = Array((new BigNumber(27)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw142.length); _i0_23++) {
            _nw142[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _952_v80 = _nw142;
          let _955_v81;
          _955_v81 = _dafny.MultiSet.fromElements((_this).f24);
          let _956_v82;
          _956_v82 = _module.D0.create_DC0(_955_v81);
          let _index126 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_952_v80).length));
          (_952_v80)[_index126] = _956_v82;
          let _index127 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_952_v80).length));
          (_952_v80)[_index127] = _module.__default.fm43(p0, (_this).f24, p0, globalState);
          let _957_v83;
          _957_v83 = _dafny.Seq.of(_950_v78, _950_v78);
          let _rhs157 = (_this).f24;
          let _rhs158 = (_939_i3).isLessThan(_939_i3);
          let _rhs159 = _957_v83;
          let _lhs112 = globalState;
          _lhs112.f13 = _rhs157;
          r0 = _rhs158;
          _957_v83 = _rhs159;
          let _958_v84;
          _958_v84 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(43));
          let _959_v85;
          _959_v85 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_958_v84);
          let _960_v86;
          _960_v86 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,((((_959_v85).contains(_this.f23)) ? ((_959_v85).get(_this.f23)) : (_958_v84))).contains(!(p0)));
          r1 = new BigNumber((_960_v86).length);
        }
        let _961_v87;
        _961_v87 = _module.D5.create_DC14(p0);
        let _962_v88;
        _962_v88 = _dafny.Seq.of(_961_v87, _961_v87);
        let _963_v89;
        _963_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_962_v88);
        _963_v89 = (_963_v89).update((_this).f24, _962_v88);
        let _964_v90;
        let _nw143 = new _module.C1();
        _nw143.__ctor((_this).f24, true);
        _964_v90 = _nw143;
        let _965_v91;
        _965_v91 = _964_v90;
        let _966_v92;
        _966_v92 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f24) ? (_939_i3) : ((_this).f33)),_965_v91);
        _966_v92 = (_966_v92).update(_module.__default.safeDivisionInt((_this).f33, (_this).f33), _965_v91);
        let _967_v93;
        _967_v93 = _dafny.Seq.UnicodeFromString("nvxmib");
        (globalState).f7 = _module.__default.safeModuloInt(_939_i3, new BigNumber((_967_v93).length));
      }
      let _968_v94;
      let _init24 = function (_969_i5) {
        return false;
      };
      let _nw144 = Array((new BigNumber(10)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw144.length); _i0_24++) {
        _nw144[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _968_v94 = _nw144;
      let _970_v95;
      _970_v95 = _dafny.Map.Empty.slice().updateUnsafe(_968_v94,_this.f23);
      let _971_v96;
      let _nw145 = new _module.C3();
      _nw145.__ctor(_970_v95, (_this).f25, p1, p0);
      _971_v96 = _nw145;
      r0 = p0;
      r1 = (_this).f33;
      let _972_v98;
      _972_v98 = _dafny.MultiSet.fromElements((_this).f33);
      let _973_v99;
      _973_v99 = _dafny.Set.fromElements(_968_v94);
      let _974_v100;
      _974_v100 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll71 = new _dafny.Map();
        for (const _compr_71 of (_dafny.Map.Empty.slice().updateUnsafe((((_972_v98).contains((_this).f33)) ? ((_972_v98).get((_this).f33)) : ((_this).f33)),new BigNumber(240))).Keys.Elements) {
          let _975_v97 = _compr_71;
          if ((_dafny.Map.Empty.slice().updateUnsafe((((_972_v98).contains((_this).f33)) ? ((_972_v98).get((_this).f33)) : ((_this).f33)),new BigNumber(240))).contains(_975_v97)) {
            _coll71.push([(_975_v97).multipliedBy((_this).f33),(_this).f33]);
          }
        }
        return _coll71;
      }(),_973_v99);
      let _976_v101;
      _976_v101 = _dafny.Seq.UnicodeFromString("ktaasd");
      let _977_v102;
      _977_v102 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_976_v101).length),(_this).f33);
      r2 = (_973_v99).IsSubsetOf((((_974_v100).contains(_977_v102)) ? ((_974_v100).get(_977_v102)) : (_973_v99)));
      return [r0, r1, r2];
    }
    m10(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _978_v0;
      _978_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(new BigNumber(237)).isLessThanOrEqualTo((_this).f33));
      _978_v0 = (_978_v0).update((_this).f33, ((p0) ? ((_this).f24) : (p0)));
      let _979_v1;
      _979_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f24);
      _979_v1 = (_979_v1).update(!((_this).f24), (_this).f24);
      (globalState).f13 = (((false) ? (p0) : (p0))) && (p0);
      let _hi4 = (_this).f33;
      for (let _980_i0 = (_this).f33; _980_i0.isLessThan(_hi4); _980_i0 = _980_i0.plus(_dafny.ONE)) {
        _979_v1 = _979_v1;
        let _981_v2;
        let _nw146 = Array((new BigNumber(2)).toNumber()).fill(false);
        _981_v2 = _nw146;
        let _index128 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_981_v2).length));
        (_981_v2)[_index128] = true;
        let _index129 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_981_v2).length));
        (_981_v2)[_index129] = (_this).f24;
        (globalState).f13 = !(p0);
        let _982_v3;
        _982_v3 = _dafny.Seq.of((_this).f24);
        let _983_v4;
        _983_v4 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.of((_981_v2)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_981_v2).length))]), _982_v3));
        _983_v4 = _983_v4;
      }
      let _984_i1;
      _984_i1 = _dafny.ZERO;
      L4: {
        while (((_this).f33).isEqualTo((_this).f33)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_984_i1)) {
              break L4;
            }
            _984_i1 = (_984_i1).plus(_dafny.ONE);
            (globalState).f13 = _module.__default.fm2(globalState);
            (globalState).f7 = (_this).f33;
            let _985_v5;
            _985_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),(_this).f33);
            let _986_v6;
            _986_v6 = _dafny.Seq.of((_this).f33);
            let _987_v7;
            _987_v7 = _986_v6;
            let _988_v8;
            _988_v8 = _module.D0.create_DC2(!(!((_this).f24)), (((_985_v5).contains((_this).f33)) ? ((_985_v5).get((_this).f33)) : ((_dafny.ZERO).minus(new BigNumber(((_987_v7)).length)))), p0);
            let _989_v9;
            let _init25 = function (_990_i2) {
              return (_990_i2).multipliedBy((_this).f33);
            };
            let _nw147 = Array((new BigNumber(24)).toNumber());
            for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw147.length); _i0_25++) {
              _nw147[_i0_25] = _init25(new BigNumber(_i0_25));
            }
            _989_v9 = _nw147;
            let _991_v10;
            _991_v10 = _dafny.Map.Empty.slice().updateUnsafe(_988_v8,_989_v9);
            let _992_v11;
            _992_v11 = _dafny.Seq.of(p0);
            _991_v10 = (_991_v10).update(_module.D0.create_DC2(!(p0), new BigNumber((_992_v11).length), (_this).f24), _989_v9);
            let _993_v12;
            let _init26 = ((_994_v6) => function (_995_i3) {
              return _994_v6;
            })(_986_v6);
            let _nw148 = Array((new BigNumber(25)).toNumber());
            for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw148.length); _i0_26++) {
              _nw148[_i0_26] = _init26(new BigNumber(_i0_26));
            }
            _993_v12 = _nw148;
            let _index130 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_993_v12).length));
            (_993_v12)[_index130] = (_this).fm7(globalState);
            let _index131 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_989_v9).length));
            (_989_v9)[_index131] = ((_this).f33).plus((_this).f33);
            let _996_v13;
            _996_v13 = _dafny.Seq.UnicodeFromString("xfjecqwws");
            let _index132 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_993_v12).length));
            let _index133 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_989_v9).length));
            let _rhs160 = _986_v6;
            let _rhs161 = _module.__default.safeModuloInt((new BigNumber(498)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_this).f33, new BigNumber((_996_v13).length))).length))), (new BigNumber(-59)).plus((_this).f33));
            let _rhs162 = _module.__default.safeModuloInt((_this).f33, (_this).f33);
            let _lhs113 = _993_v12;
            let _lhs114 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_993_v12).length));
            let _lhs115 = globalState;
            let _lhs116 = _989_v9;
            let _lhs117 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_989_v9).length));
            _lhs113[_lhs114] = _rhs160;
            _lhs115.f7 = _rhs161;
            _lhs116[_lhs117] = _rhs162;
          }
        }
      }
      r0 = (_this).f33;
      let _997_v14;
      _997_v14 = _dafny.MultiSet.fromElements(p0);
      r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_997_v14).contains((_this).f24)) ? ((_997_v14).get((_this).f24)) : ((_this).f33)), (_this).f33)));
      return r0;
    }
    get f33() {
      let _this = this;
      return _this._f33;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f32 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f32) {
      let _this = this;
      (_this)._f32 = f32;
      return;
    }
    fm17(p0, p1, p2, globalState) {
      let _this = this;
      return ((new BigNumber(744)).multipliedBy(new BigNumber(-870))).plus(new BigNumber(((_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).Intersect(_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0))))).cardinality()));
    };
    m8(globalState) {
      let _this = this;
      let r0 = [];
      let _998_v0;
      _998_v0 = true;
      let _999_v1;
      _999_v1 = _dafny.Seq.of(_998_v0, _998_v0, _998_v0);
      let _1000_v2;
      _1000_v2 = new BigNumber(987);
      let _1001_v3;
      _1001_v3 = _dafny.MultiSet.fromElements(new BigNumber((_999_v1).length), _1000_v2);
      let _1002_v4;
      _1002_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_998_v0, _998_v0, _998_v0)).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-624)), function (_1003_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }));
      (globalState).f13 = ((_module.__default.fm18(globalState)).Difference(_1001_v3)).IsProperSubsetOf((_1001_v3).update(new BigNumber(((((_1002_v4).contains(_1000_v2)) ? ((_1002_v4).get(_1000_v2)) : (_dafny.Seq.UnicodeFromString("rnfuhyqq")))).length), _module.__default.abs(_1000_v2)));
      let _1004_v5;
      _1004_v5 = _dafny.Set.fromElements(new BigNumber(164));
      _1004_v5 = (_1004_v5).Union(function () {
        let _coll72 = new _dafny.Set();
        for (const _compr_72 of _dafny.IntegerRange(new BigNumber(413), new BigNumber(860))) {
          let _1005_v6 = _compr_72;
          if (((new BigNumber(413)).isLessThanOrEqualTo(_1005_v6)) && ((_1005_v6).isLessThan(new BigNumber(860)))) {
            _coll72.add((_1005_v6).minus(_1000_v2));
          }
        }
        return _coll72;
      }());
      let _hi5 = _1000_v2;
      for (let _1006_i1 = _1000_v2; _1006_i1.isLessThan(_hi5); _1006_i1 = _1006_i1.plus(_dafny.ONE)) {
        if (_998_v0) {
          let _1007_v7;
          _1007_v7 = _dafny.Seq.UnicodeFromString("qwdbajgpg");
          let _1008_v8;
          _1008_v8 = _dafny.Map.Empty.slice().updateUnsafe(_998_v0,(((_1001_v3).contains(new BigNumber((_1004_v5).length))) ? ((_1001_v3).get(new BigNumber((_1004_v5).length))) : (_1000_v2)));
          let _1009_v9;
          _1009_v9 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1010_v10;
          let _nw149 = Array((new BigNumber(6)).toNumber());
          _nw149[(_dafny.ZERO).toNumber()] = _1000_v2;
          _nw149[(_dafny.ONE).toNumber()] = _module.__default.fm0(_998_v0, _1000_v2, globalState);
          _nw149[(new BigNumber(2)).toNumber()] = (_this).fm17(new BigNumber((_dafny.Seq.update(_1007_v7, _module.__default.safeIndex(new BigNumber((_1008_v8).length), new BigNumber((_1007_v7).length)), _1009_v9)).length), new _dafny.CodePoint('g'.codePointAt(0)), _module.__default.fm2(globalState), globalState);
          _nw149[(new BigNumber(3)).toNumber()] = _1006_i1;
          _nw149[(new BigNumber(4)).toNumber()] = (_1006_i1).minus(_1006_i1);
          _nw149[(new BigNumber(5)).toNumber()] = _module.__default.fm0(_998_v0, _1006_i1, globalState);
          _1010_v10 = _nw149;
          let _index134 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1010_v10).length));
          (_1010_v10)[_index134] = new BigNumber(454);
          let _index135 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1010_v10).length));
          (_1010_v10)[_index135] = _module.__default.safeDivisionInt((_this).fm17(_1006_i1, _1009_v9, _998_v0, globalState), _1006_i1);
          (globalState).f7 = _1006_i1;
          let _1011_v11;
          _1011_v11 = _dafny.Map.Empty.slice().updateUnsafe(_998_v0,_998_v0);
          let _1012_v12;
          _1012_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_998_v0,_dafny.Seq.update(_999_v1, _module.__default.safeIndex(_1006_i1, new BigNumber((_999_v1).length)), _998_v0)),(_1011_v11).Merge(_module.__default.fm19(globalState)));
          let _1013_v13;
          _1013_v13 = _dafny.Map.Empty.slice().updateUnsafe(_998_v0,_999_v1);
          _1012_v12 = (_1012_v12).update(((_998_v0) ? (_1013_v13) : (_1013_v13)), _1011_v11);
          let _index136 = _module.__default.safeIndex(new BigNumber(295), new BigNumber(((_this).f32).length));
          ((_this).f32)[_index136] = _998_v0;
          let _1014_v14;
          _1014_v14 = _dafny.Seq.of((_1010_v10)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1010_v10).length))], new BigNumber(-537));
          let _index137 = _module.__default.safeIndex(new BigNumber(295), new BigNumber(((_this).f32).length));
          ((_this).f32)[_index137] = (_dafny.MultiSet.fromElements(_1014_v14, _1014_v14)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1014_v14));
          (globalState).f7 = (_1010_v10)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1010_v10).length))];
        } else {
          let _1015_v15;
          let _out35;
          _out35 = (_this).m9((new BigNumber(-814)).isLessThanOrEqualTo(_1000_v2), _998_v0, globalState);
          _1015_v15 = _out35;
          (globalState).f13 = _998_v0;
          (globalState).f7 = _1000_v2;
          let _1016_v16;
          _1016_v16 = _dafny.Seq.of(_1000_v2, _1000_v2);
          (globalState).f7 = _module.__default.fm0(false, _module.__default.safeDivisionInt(_1006_i1, new BigNumber((_1016_v16).length)), globalState);
          let _1017_v17;
          let _nw150 = Array((new BigNumber(29)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-368)), ((_1018_i1) => function (_1019_i2) {
            return _1018_i1;
          })(_1006_i1));
          _nw150[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_1006_i1, _1006_i1, _1000_v2);
          _nw150[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(182)), ((_1020_v2) => function (_1021_i3) {
            return _1020_v2;
          })(_1000_v2));
          _nw150[(new BigNumber(3)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(4)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(5)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(6)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(7)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_1000_v2, _1006_i1);
          _nw150[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-645)), ((_1022_v2) => function (_1023_i4) {
            return (_dafny.ZERO).minus(_1022_v2);
          })(_1000_v2));
          _nw150[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), ((_1024_i1) => function (_1025_i5) {
            return _1024_i1;
          })(_1006_i1));
          _nw150[(new BigNumber(11)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(12)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(13)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(14)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(15)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(16)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(17)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(18)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_1006_i1, _1000_v2);
          _nw150[(new BigNumber(20)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(21)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(22)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(23)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(24)).toNumber()] = _dafny.Seq.of(_1006_i1);
          _nw150[(new BigNumber(25)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(26)).toNumber()] = _dafny.Seq.of(new BigNumber((_999_v1).length), _1000_v2);
          _nw150[(new BigNumber(27)).toNumber()] = _1016_v16;
          _nw150[(new BigNumber(28)).toNumber()] = _1016_v16;
          _1017_v17 = _nw150;
          let _1026_v18;
          _1026_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1017_v17,_998_v0);
          let _1027_v19;
          _1027_v19 = _dafny.MultiSet.fromElements((((_1026_v18).contains(_1017_v17)) ? ((_1026_v18).get(_1017_v17)) : (_998_v0)));
          let _1028_v20;
          _1028_v20 = _module.D0.create_DC0(_1027_v19);
          let _1029_v21;
          let _nw151 = Array((new BigNumber(19)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = _1028_v20;
          _nw151[(_dafny.ONE).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(2)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(3)).toNumber()] = ((!(_998_v0)) ? (_1028_v20) : (_1028_v20));
          _nw151[(new BigNumber(4)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(5)).toNumber()] = _module.D0.create_DC0(_dafny.MultiSet.fromElements(!(true), _998_v0));
          _nw151[(new BigNumber(6)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(7)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(8)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(9)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(10)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(11)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(12)).toNumber()] = _module.D0.create_DC0(_1027_v19);
          _nw151[(new BigNumber(13)).toNumber()] = _module.__default.fm20(new BigNumber((_1027_v19).cardinality()), _module.__default.fm20(_1000_v2, _1028_v20, globalState), globalState);
          _nw151[(new BigNumber(14)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(15)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(16)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(17)).toNumber()] = _1028_v20;
          _nw151[(new BigNumber(18)).toNumber()] = _1028_v20;
          _1029_v21 = _nw151;
          let _1030_v22;
          _1030_v22 = _dafny.Seq.UnicodeFromString("dgjysbxnh");
          let _1031_v23;
          let _init27 = ((_1032_v2) => function (_1033_i6) {
            return (_1033_i6).multipliedBy(_1032_v2);
          })(_1000_v2);
          let _nw152 = Array((new BigNumber(4)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw152.length); _i0_27++) {
            _nw152[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1031_v23 = _nw152;
          let _rhs163 = _1029_v21;
          let _rhs164 = _1031_v23;
          let _rhs165 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(_1000_v2, _1006_i1), _1006_i1));
          let _rhs166 = _998_v0;
          let _rhs167 = _1030_v22;
          let _lhs118 = globalState;
          let _lhs119 = globalState;
          let _lhs120 = globalState;
          _1029_v21 = _rhs163;
          _lhs118.f22 = _rhs164;
          _lhs119.f7 = _rhs165;
          _lhs120.f13 = _rhs166;
          _1030_v22 = _rhs167;
        }
        _1000_v2 = _1006_i1;
        (globalState).f13 = _998_v0;
        let _1034_v24;
        _1034_v24 = new _dafny.CodePoint('i'.codePointAt(0));
        _1034_v24 = _1034_v24;
      }
      let _1035_v25;
      _1035_v25 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1000_v2,_998_v0)).length)).isLessThanOrEqualTo(_1000_v2),_998_v0);
      _1035_v25 = (_1035_v25).update(_998_v0, _998_v0);
      let _1036_v26;
      let _nw153 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _1036_v26 = _nw153;
      let _1037_v27;
      _1037_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1000_v2,_1000_v2);
      let _1038_v28;
      _1038_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1037_v27,_998_v0);
      let _index138 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1036_v26).length));
      (_1036_v26)[_index138] = (new BigNumber((_1038_v28).length)).plus(_1000_v2);
      let _1039_v29;
      _1039_v29 = _dafny.Seq.UnicodeFromString("ejirnglkd");
      let _1040_v30;
      _1040_v30 = _module.D2.create_DC7(_998_v0, _998_v0);
      let _1041_v31;
      _1041_v31 = _module.D2.create_DC7(_dafny.Seq.IsProperPrefixOf(_1039_v29, _1039_v29), (_1040_v30).dtor_cf13);
      let _index139 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1036_v26).length));
      let _rhs168 = _1000_v2;
      let _rhs169 = _1040_v30;
      let _rhs170 = (_1000_v2).plus(new BigNumber((_dafny.Seq.update(((_998_v0) ? (_999_v1) : (_dafny.Seq.of(_998_v0))), _module.__default.safeIndex((_dafny.ZERO).minus(_1000_v2), new BigNumber((((_998_v0) ? (_999_v1) : (_dafny.Seq.of(_998_v0)))).length)), _998_v0)).length));
      let _rhs171 = _998_v0;
      let _rhs172 = _module.__default.safeDivisionInt(_1000_v2, (_1000_v2).multipliedBy(new BigNumber(304)));
      let _lhs121 = _1036_v26;
      let _lhs122 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1036_v26).length));
      let _lhs123 = globalState;
      let _lhs124 = globalState;
      let _lhs125 = globalState;
      _lhs121[_lhs122] = _rhs168;
      _1041_v31 = _rhs169;
      _lhs123.f7 = _rhs170;
      _lhs124.f13 = _rhs171;
      _lhs125.f7 = _rhs172;
      (globalState).f13 = _998_v0;
      r0 = (_this).f32;
      return r0;
    }
    m9(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let _1042_v0;
      _1042_v0 = _dafny.Seq.UnicodeFromString("ri");
      let _1043_v1;
      _1043_v1 = _dafny.Seq.of(_1042_v0, _dafny.Seq.UnicodeFromString("gwjxmn"));
      let _1044_v2;
      _1044_v2 = _dafny.Set.fromElements(p1, p0);
      let _1045_v3;
      _1045_v3 = new BigNumber(858);
      let _1046_v4;
      _1046_v4 = _dafny.MultiSet.fromElements(_1045_v3, _1045_v3);
      _1043_v1 = _dafny.Seq.Concat(_dafny.Seq.of(_1042_v0, _1042_v0, _1042_v0), _module.__default.fm21(_1044_v2, _1046_v4, globalState));
      if (p0) {
        let _1047_v5;
        _1047_v5 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1048_v6;
        _1048_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("av"), _module.__default.safeIndex(_1045_v3, new BigNumber((_dafny.Seq.UnicodeFromString("av")).length)), _1047_v5));
        let _1049_v7;
        let _nw154 = new _module.C0();
        _nw154.__ctor(((p0) ? (_1048_v6) : (_1048_v6)), (!(p0)) && (p1), _1047_v5, p0);
        _1049_v7 = _nw154;
        (globalState).f7 = _1045_v3;
        let _1050_v8;
        _1050_v8 = _dafny.Seq.of((_1049_v7).f38, !((_1049_v7).f38));
        let _1051_v9;
        let _nw155 = Array((new BigNumber(13)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = _1050_v8;
        _nw155[(_dafny.ONE).toNumber()] = _1050_v8;
        _nw155[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(p1, p1, (_1049_v7).f38);
        _nw155[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1050_v8, _1050_v8);
        _nw155[(new BigNumber(4)).toNumber()] = _1050_v8;
        _nw155[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_1049_v7).f38, p0), _1050_v8);
        _nw155[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1050_v8, _1050_v8);
        _nw155[(new BigNumber(7)).toNumber()] = _1050_v8;
        _nw155[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(!((_1049_v7).f38), p1, !(p1));
        _nw155[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(p0);
        _nw155[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1050_v8, _1050_v8);
        _nw155[(new BigNumber(11)).toNumber()] = _1050_v8;
        _nw155[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p1), _1050_v8);
        _1051_v9 = _nw155;
        let _index140 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1051_v9).length));
        (_1051_v9)[_index140] = _1050_v8;
        let _1052_v10;
        _1052_v10 = _dafny.MultiSet.fromElements(p0);
        let _1053_v11;
        _1053_v11 = _dafny.Set.fromElements(new BigNumber(829), _1045_v3, _1045_v3, _1045_v3);
        let _index141 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1051_v9).length));
        let _rhs173 = _1047_v5;
        let _rhs174 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(p0), _1050_v8), _dafny.Seq.update(_1050_v8, _module.__default.safeIndex((((_1052_v10).contains(true)) ? ((_1052_v10).get(true)) : (_1045_v3)), new BigNumber((_1050_v8).length)), p1));
        let _rhs175 = false;
        let _rhs176 = (_1053_v11).IsDisjointFrom(_1053_v11);
        let _lhs126 = _1051_v9;
        let _lhs127 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1051_v9).length));
        let _lhs128 = globalState;
        let _lhs129 = globalState;
        _1047_v5 = _rhs173;
        _lhs126[_lhs127] = _rhs174;
        _lhs128.f13 = _rhs175;
        _lhs129.f13 = _rhs176;
        (globalState).f7 = _1045_v3;
        if ((_1049_v7).f38) {
          let _1054_v12;
          _1054_v12 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1045_v3, _1045_v3)));
          _1054_v12 = (_1049_v7).fm26(new BigNumber((_1042_v0).length), ((p0) ? ((_1051_v9)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_1051_v9).length))]) : ((_1051_v9)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_1051_v9).length))])), function () {
            let _coll73 = new _dafny.Set();
            for (const _compr_73 of _dafny.IntegerRange(new BigNumber(885), new BigNumber(965))) {
              let _1055_v13 = _compr_73;
              if (((new BigNumber(885)).isLessThanOrEqualTo(_1055_v13)) && ((_1055_v13).isLessThan(new BigNumber(965)))) {
                _coll73.add(_module.__default.safeModuloInt(_1055_v13, _1045_v3));
              }
            }
            return _coll73;
          }(), (_1049_v7).f38, globalState);
          (globalState).f13 = p1;
          (globalState).f13 = (_1049_v7).f38;
          (globalState).f7 = _1045_v3;
          let _1056_v14;
          let _nw156 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1056_v14 = _nw156;
          let _index142 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1056_v14).length));
          (_1056_v14)[_index142] = _1047_v5;
          let _index143 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1056_v14).length));
          let _rhs177 = (_1043_v1)[_module.__default.safeIndex(_1045_v3, new BigNumber((_1043_v1).length))];
          let _rhs178 = new _dafny.CodePoint('l'.codePointAt(0));
          let _lhs130 = _1056_v14;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1056_v14).length));
          _1042_v0 = _rhs177;
          _lhs130[_lhs131] = _rhs178;
        } else {
          let _1057_v15;
          let _nw157 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1057_v15 = _nw157;
          let _1058_v16;
          let _nw158 = new _module.C2();
          _nw158.__ctor(_1052_v10, _1057_v15, _1047_v5, (_1049_v7).f38);
          _1058_v16 = _nw158;
          let _1059_v17;
          _1059_v17 = _dafny.Seq.of(_module.__default.safeDivisionInt((_this).fm17(_1045_v3, _1047_v5, (_1049_v7).f38, globalState), _1045_v3));
          _1059_v17 = _dafny.Seq.Concat((_1058_v16).fm7(globalState), _1059_v17);
          let _1060_v18;
          _1060_v18 = _dafny.MultiSet.fromElements(_1053_v11, (_1053_v11).Difference(_1053_v11), _1053_v11, _1053_v11, (_1053_v11).Union(_1053_v11));
          _1060_v18 = ((_1060_v18).Difference(_1060_v18)).Intersect((_1060_v18).Union(_dafny.MultiSet.fromElements(_1053_v11, _1053_v11)));
          let _index144 = _module.__default.safeIndex(new BigNumber(648), new BigNumber(((_this).f32).length));
          ((_this).f32)[_index144] = (_1049_v7).f38;
          let _index145 = _module.__default.safeIndex(new BigNumber(648), new BigNumber(((_this).f32).length));
          ((_this).f32)[_index145] = (_1049_v7).f38;
          (globalState).f7 = _1045_v3;
        }
      } else {
        let _1061_v19;
        let _nw159 = new _module.C1();
        _nw159.__ctor(p1, p1);
        _1061_v19 = _nw159;
        let _index146 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length));
        ((_this).f32)[_index146] = p1;
        let _1062_v20;
        _1062_v20 = _dafny.Seq.of((_1061_v19).f39);
        let _index147 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length));
        ((_this).f32)[_index147] = !(((false) ? ((_1045_v3).isLessThan(new BigNumber((_1062_v20).length))) : (true)));
        let _index148 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length));
        ((_this).f32)[_index148] = ((_this).f32)[_module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length))];
        let _1063_v21;
        let _init28 = ((_1064_v20) => function (_1065_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,_1064_v20);
        })(_1062_v20);
        let _nw160 = Array((_dafny.ONE).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw160.length); _i0_28++) {
          _nw160[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1063_v21 = _nw160;
        let _1066_v22;
        _1066_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1061_v19).f39,_module.__default.fm36(globalState));
        let _1067_v23;
        _1067_v23 = _module.D9.create_DC19(_1066_v22);
        let _index149 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1063_v21).length));
        (_1063_v21)[_index149] = (_1067_v23).dtor_cf30;
        let _index150 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1063_v21).length));
        let _index151 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length));
        let _rhs179 = _1066_v22;
        let _rhs180 = p0;
        let _rhs181 = (_1061_v19).f39;
        let _rhs182 = p0;
        let _lhs132 = _1063_v21;
        let _lhs133 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1063_v21).length));
        let _lhs134 = globalState;
        let _lhs135 = globalState;
        let _lhs136 = (_this).f32;
        let _lhs137 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length));
        _lhs132[_lhs133] = _rhs179;
        _lhs134.f13 = _rhs180;
        _lhs135.f13 = _rhs181;
        _lhs136[_lhs137] = _rhs182;
        if (_dafny.Seq.IsPrefixOf(_1062_v20, _1062_v20)) {
          let _1068_v24;
          let _nw161 = new _module.C1();
          _nw161.__ctor(((p1) ? (((_this).f32)[_module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length))]) : (p1)), p1);
          _1068_v24 = _nw161;
          (globalState).f13 = (((_this).f32)[_module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length))]) === ((_1068_v24).f39);
          (globalState).f7 = _1045_v3;
          (globalState).f7 = (_1045_v3).plus((_dafny.ZERO).minus(_1045_v3));
          let _1069_v25;
          let _init29 = ((_1070_v19, _1071_p1, _1072_p0, _1073_v24) => function (_1074_i1) {
            return (_dafny.MultiSet.fromElements(_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(((_this).f32)[_module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length))],((_this).f32)[_module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length))]), (_1070_v19).f39, _1071_p1), _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(!(_1072_p0),_1072_p0), (_1070_v19).f39, (_1073_v24).f39))).Intersect(_dafny.MultiSet.fromElements(_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(((_this).f32)[_module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f32).length))],(_1070_v19).f39), (_1073_v24).f39, !((_1070_v19).f39))));
          })(_1061_v19, p1, p0, _1068_v24);
          let _nw162 = Array((new BigNumber(25)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw162.length); _i0_29++) {
            _nw162[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1069_v25 = _nw162;
          let _1075_v26;
          _1075_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _1076_v27;
          _1076_v27 = _module.D6.create_DC16(_1075_v26, true, p1);
          let _1077_v28;
          _1077_v28 = _dafny.MultiSet.fromElements(_1076_v27, _1076_v27);
          let _index152 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1069_v25).length));
          (_1069_v25)[_index152] = (_1077_v28).Union((_1077_v28).update(_1076_v27, _module.__default.abs(_1045_v3)));
          let _index153 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1069_v25).length));
          (_1069_v25)[_index153] = _1077_v28;
        } else {
          let _1078_v29;
          _1078_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1061_v19).f39,_1042_v0);
          let _1079_v30;
          _1079_v30 = _module.D2.create_DC6(_1078_v29);
          _1079_v30 = _1079_v30;
          let _1080_v31;
          _1080_v31 = new _dafny.CodePoint('c'.codePointAt(0));
          _1080_v31 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1081_v32;
          _1081_v32 = _dafny.MultiSet.fromElements(!((_1061_v19).f39));
          let _1082_v33;
          _1082_v33 = _module.D1.create_DC4(_1080_v31);
          let _1083_v34;
          let _nw163 = Array((new BigNumber(28)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = _1080_v31;
          _nw163[(_dafny.ONE).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
          _nw163[(new BigNumber(3)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(4)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(5)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
          _nw163[(new BigNumber(7)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(8)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(9)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(10)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(11)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(12)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(13)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(14)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(15)).toNumber()] = (_1082_v33).dtor_cf6;
          _nw163[(new BigNumber(16)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(17)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
          _nw163[(new BigNumber(19)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(20)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(21)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(22)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(23)).toNumber()] = (_1082_v33).dtor_cf6;
          _nw163[(new BigNumber(24)).toNumber()] = _1080_v31;
          _nw163[(new BigNumber(25)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
          _nw163[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw163[(new BigNumber(27)).toNumber()] = _1080_v31;
          _1083_v34 = _nw163;
          let _1084_v35;
          let _nw164 = new _module.C2();
          _nw164.__ctor(_1081_v32, _1083_v34, _1080_v31, (_1061_v19).f39);
          _1084_v35 = _nw164;
          let _1085_v36;
          _1085_v36 = _module.D4.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(198)), ((_1086_v31) => function (_1087_i2) {
  return _1086_v31;
})(_1080_v31)));
          let _1088_v37;
          _1088_v37 = _dafny.Set.fromElements(new BigNumber(-793));
          _1085_v36 = function (_pat_let43_0) {
            return function (_1089_dt__update__tmp_h0) {
              return function (_pat_let44_0) {
                return function (_1090_dt__update_hcf17_h0) {
                  return _module.D4.create_DC10(_1090_dt__update_hcf17_h0);
                }(_pat_let44_0);
              }(_dafny.Seq.UnicodeFromString("qrk"));
            }(_pat_let43_0);
          }(_module.__default.fm44(new BigNumber((_1088_v37).length), globalState));
          let _1091_v38;
          _1091_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1061_v19).f39,new BigNumber(899));
          let _1092_v39;
          _1092_v39 = _dafny.Seq.of(new BigNumber((_1091_v38).length));
          let _1093_v40;
          let _init30 = ((_1094_v3) => function (_1095_i3) {
            return _module.__default.safeDivisionInt(_1095_i3, _1094_v3);
          })(_1045_v3);
          let _nw165 = Array((new BigNumber(15)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw165.length); _i0_30++) {
            _nw165[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1093_v40 = _nw165;
          let _1096_v41;
          _1096_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1093_v40,_dafny.Seq.update(_dafny.Seq.update(_1042_v0, _module.__default.safeIndex(_1045_v3, new BigNumber((_1042_v0).length)), _1080_v31), _module.__default.safeIndex(_1045_v3, new BigNumber((_dafny.Seq.update(_1042_v0, _module.__default.safeIndex(_1045_v3, new BigNumber((_1042_v0).length)), _1080_v31)).length)), _1080_v31));
          let _1097_v42;
          _1097_v42 = _dafny.Map.Empty.slice().updateUnsafe((_1061_v19).fm9(_1092_v39, _1045_v3, (_1061_v19).f39, globalState),_1096_v41);
          _1097_v42 = (_1097_v42).update((_1061_v19).fm9(_1092_v39, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-662)), ((_1098_v3) => function (_1099_i4) {
            return _1098_v3;
          })(_1045_v3))).length), (_1061_v19).f39, globalState), _1096_v41);
        }
      }
      (globalState).f7 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1045_v3), _1045_v3);
      let _1100_v43;
      _1100_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1045_v3,(_dafny.ZERO).minus(_1045_v3));
      let _hi6 = (((_1100_v43).contains(_1045_v3)) ? ((_1100_v43).get(_1045_v3)) : (_1045_v3));
      for (let _1101_i5 = (new BigNumber(56)).plus(new BigNumber(431)); _1101_i5.isLessThan(_hi6); _1101_i5 = _1101_i5.plus(_dafny.ONE)) {
        let _1102_v44;
        let _nw166 = Array((new BigNumber(9)).toNumber()).fill([]);
        _1102_v44 = _nw166;
        _1102_v44 = _1102_v44;
        (globalState).f13 = !(_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), function (_1103_i6) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _1042_v0), _1042_v0));
        (globalState).f13 = p0;
        _1046_v4 = (_1046_v4).Intersect(_1046_v4);
      }
      let _1104_v45;
      let _nw167 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
      _1104_v45 = _nw167;
      let _1105_v46;
      let _init31 = function (_1106_i7) {
        return (_1106_i7).minus(new BigNumber((_dafny.Seq.of(false)).length));
      };
      let _nw168 = Array((new BigNumber(5)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw168.length); _i0_31++) {
        _nw168[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1105_v46 = _nw168;
      let _1107_v47;
      _1107_v47 = _dafny.Set.fromElements(_1105_v46);
      let _index154 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1104_v45).length));
      (_1104_v45)[_index154] = _1107_v47;
      let _index155 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1104_v45).length));
      (_1104_v45)[_index155] = ((!(p1)) ? (_1107_v47) : (_1107_v47));
      let _1108_v48;
      let _init32 = function (_1109_i8) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      };
      let _nw169 = Array((new BigNumber(23)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw169.length); _i0_32++) {
        _nw169[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1108_v48 = _nw169;
      let _1110_v49;
      _1110_v49 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1111_v50;
      let _nw170 = new _module.C4();
      _nw170.__ctor(new BigNumber(-768), _1100_v43, _1108_v48, _1110_v49, p0);
      _1111_v50 = _nw170;
      let _1112_v51;
      let _nw171 = Array((new BigNumber(6)).toNumber()).fill([]);
      _1112_v51 = _nw171;
      let _1113_v52;
      _1113_v52 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1112_v51);
      r0 = (((_1113_v52).contains((_1045_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_1043_v1, _module.__default.safeIndex(new BigNumber(-248), new BigNumber((_1043_v1).length)), _1042_v0)).length)))) ? ((_1113_v52).get((_1045_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_1043_v1, _module.__default.safeIndex(new BigNumber(-248), new BigNumber((_1043_v1).length)), _1042_v0)).length)))) : (_1112_v51));
      return r0;
    }
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f23, f24) {
      let _this = this;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return _module.D0.create_DC3(_module.D0.create_DC1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_this.f23)).length),_dafny.Seq.UnicodeFromString("lwt"))).length)));
    };
    fm6(globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), function (_1114_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("gt")))).length);
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f24), (_module.D0.create_DC0(_dafny.MultiSet.fromElements((_this).f24, (_this).f24, (_this).f24, (_this).f24, (_this).f24))).dtor_cf0)).IsDisjointFrom((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f24))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.fromElements(false, (_this).f24)))));
    };
    fm16(globalState) {
      let _this = this;
      return !(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(298))).contains((_this).f24);
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D0.Default();
      let r2 = false;
      let _1115_v0;
      _1115_v0 = new BigNumber(914);
      let _hi7 = _1115_v0;
      for (let _1116_i0 = _1115_v0; _1116_i0.isLessThan(_hi7); _1116_i0 = _1116_i0.plus(_dafny.ONE)) {
        let _1117_v1;
        _1117_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("hblxvj"));
        let _1118_v2;
        let _nw172 = new _module.C0();
        _nw172.__ctor(_1117_v1, (_this).f24, _this.f23, (new BigNumber((_dafny.Seq.UnicodeFromString("xu")).length)).isEqualTo(_1116_i0));
        _1118_v2 = _nw172;
        let _1119_v3;
        _1119_v3 = _dafny.Seq.UnicodeFromString("c");
        (globalState).f7 = new BigNumber((_1119_v3).length);
        let _rhs183 = _module.__default.safeDivisionInt(_1116_i0, _1115_v0);
        let _rhs184 = _1115_v0;
        let _lhs138 = globalState;
        _lhs138.f7 = _rhs183;
        _1115_v0 = _rhs184;
        let _1120_v4;
        _1120_v4 = _dafny.Seq.of((_this).f24, (_this).f24);
        if ((((_1118_v2).f38) ? ((((_this).f24) ? (!(p1)) : (true))) : ((_1120_v4)[_module.__default.safeIndex(_1115_v0, new BigNumber((_1120_v4).length))]))) {
          (globalState).f7 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1115_v0), (_1115_v0).minus(new BigNumber((_1120_v4).length)));
          let _1121_v5;
          let _nw173 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1121_v5 = _nw173;
          let _1122_v6;
          _1122_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1121_v5,_1115_v0);
          let _1123_v7;
          _1123_v7 = _dafny.Seq.of(_1122_v6, _1122_v6);
          let _1124_v8;
          _1124_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(((_1123_v7)[_module.__default.safeIndex((_dafny.ZERO).minus(_1116_i0), new BigNumber((_1123_v7).length))]).length));
          let _1125_v9;
          _1125_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1116_i0,(_1124_v8).Merge(_1124_v8));
          _1125_v9 = (_1125_v9).update((_dafny.ZERO).minus(_1116_i0), _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.of(_1115_v0)).length)));
          r0 = _dafny.Seq.Concat(_module.__default.fm1((_1118_v2).f38, globalState), _1119_v3);
          (_this).f23 = _this.f23;
          _1115_v0 = _1115_v0;
        } else {
          let _1126_v10;
          _1126_v10 = _dafny.Set.fromElements(_1115_v0, _1115_v0);
          let _1127_v11;
          _1127_v11 = _dafny.Seq.of(_1126_v10, _1126_v10);
          (globalState).f7 = new BigNumber((_dafny.Seq.of((new BigNumber((_1127_v11).length)).plus(_1116_i0), _1115_v0, new BigNumber(-974))).length);
          let _1128_v12;
          let _init33 = ((_1129_v0) => function (_1130_i1) {
            return (_1130_i1).minus(_1129_v0);
          })(_1115_v0);
          let _nw174 = Array((new BigNumber(12)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw174.length); _i0_33++) {
            _nw174[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1128_v12 = _nw174;
          (globalState).f7 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1115_v0, _1116_i0)), _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(_1128_v12)).cardinality()), _1115_v0));
          (globalState).f7 = _1116_i0;
          (globalState).f13 = (_this).f24;
          let _1131_v13;
          _1131_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p1);
          let _rhs185 = _1116_i0;
          let _rhs186 = _module.__default.fm19(globalState);
          _1115_v0 = _rhs185;
          _1131_v13 = _rhs186;
        }
      }
      let _1132_v14;
      let _init34 = ((_1133_v0) => function (_1134_i2) {
        return _module.__default.safeModuloInt(_1134_i2, _1133_v0);
      })(_1115_v0);
      let _nw175 = Array((new BigNumber(23)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw175.length); _i0_34++) {
        _nw175[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1132_v14 = _nw175;
      let _1135_v15;
      _1135_v15 = _dafny.Seq.of(_1132_v14);
      let _1136_v16;
      _1136_v16 = _dafny.Seq.of((_1135_v15)[_module.__default.safeIndex(_1115_v0, new BigNumber((_1135_v15).length))], _1132_v14);
      let _1137_v17;
      _1137_v17 = _dafny.Seq.UnicodeFromString("bniwatogj");
      let _1138_v18;
      _1138_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1137_v17);
      let _1139_v19;
      _1139_v19 = _module.D2.create_DC6(_1138_v18);
      _1115_v0 = (new BigNumber((((p1) ? (_1136_v16) : (_1136_v16))).length)).minus(new BigNumber((_module.__default.fm45((_this).f24, _1139_v19, _1115_v0, globalState)).length));
      r2 = (_this).f24;
      let _1140_v20;
      let _nw176 = new _module.C1();
      _nw176.__ctor(p1, p1);
      _1140_v20 = _nw176;
      let _1141_v21;
      _1141_v21 = _1140_v20;
      let _1142_v22;
      _1142_v22 = _dafny.Seq.of(_1141_v21, _1141_v21, _1141_v21, _1141_v21, _1141_v21);
      let _1143_v23;
      _1143_v23 = _1142_v22;
      let _1144_v24;
      _1144_v24 = _dafny.MultiSet.fromElements(_1142_v22, _dafny.Seq.of(_1141_v21, _1141_v21, _1141_v21), _1142_v22, (_1143_v23), _1142_v22);
      _1144_v24 = ((_1144_v24).update(_1142_v22, _module.__default.abs(_1115_v0))).Difference(_1144_v24);
      let _1145_i3;
      _1145_i3 = _dafny.ZERO;
      L5: {
        while (p1) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1145_i3)) {
              break L5;
            }
            _1145_i3 = (_1145_i3).plus(_dafny.ONE);
            (globalState).f7 = _1115_v0;
            (globalState).f22 = _1132_v14;
            let _1146_v25;
            _1146_v25 = _module.D5.create_DC13((_dafny.ZERO).minus(_1115_v0), (_1140_v20).f39, (_this).f24);
            let _1147_v26;
            _1147_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(402),(_1146_v25).dtor_cf21);
            _1147_v26 = (_1147_v26).update(new BigNumber(431), p1);
            let _1148_v27;
            _1148_v27 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(153)), ((_1149_v0) => function (_1150_i4) {
              return (_dafny.ZERO).minus(_1149_v0);
            })(_1115_v0)));
            let _1151_v28;
            let _nw177 = Array((new BigNumber(10)).toNumber());
            _nw177[(_dafny.ZERO).toNumber()] = p1;
            _nw177[(_dafny.ONE).toNumber()] = p1;
            _nw177[(new BigNumber(2)).toNumber()] = (_this).f24;
            _nw177[(new BigNumber(3)).toNumber()] = p1;
            _nw177[(new BigNumber(4)).toNumber()] = (_this).f24;
            _nw177[(new BigNumber(5)).toNumber()] = (_this).f24;
            _nw177[(new BigNumber(6)).toNumber()] = (_this).f24;
            _nw177[(new BigNumber(7)).toNumber()] = p1;
            _nw177[(new BigNumber(8)).toNumber()] = (_this).f24;
            _nw177[(new BigNumber(9)).toNumber()] = (_1140_v20).f39;
            _1151_v28 = _nw177;
            let _1152_v29;
            _1152_v29 = _module.D1.create_DC5((_1115_v0).plus(_1115_v0), new BigNumber((_1137_v17).length), _1148_v27, ((true) ? (_1151_v28) : (_1151_v28)), _1115_v0);
            let _source16 = _1152_v29;
            if (_source16.is_DC5) {
              let _1153___mcc_h0 = (_source16).cf7;
              let _1154___mcc_h1 = (_source16).cf8;
              let _1155___mcc_h2 = (_source16).cf9;
              let _1156___mcc_h3 = (_source16).cf10;
              let _1157___mcc_h4 = (_source16).cf11;
              let _1158_cf11 = _1157___mcc_h4;
              let _1159_cf10 = _1156___mcc_h3;
              let _1160_cf9 = _1155___mcc_h2;
              let _1161_cf8 = _1154___mcc_h1;
              let _1162_cf7 = _1153___mcc_h0;
              let _index156 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1159_cf10).length));
              (_1159_cf10)[_index156] = (_1140_v20).f39;
              let _1163_v30;
              _1163_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1140_v20).f39,(_1140_v20).f39);
              let _1164_v31;
              _1164_v31 = _dafny.MultiSet.fromElements((_1140_v20).f39);
              let _1165_v32;
              _1165_v32 = _dafny.Seq.of((_1140_v20).f39, p1);
              let _index157 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1159_cf10).length));
              (_1159_cf10)[_index157] = (((_1163_v30).contains((_1140_v20).f39)) ? ((_1163_v30).get((_1140_v20).f39)) : ((_dafny.MultiSet.FromArray(_1165_v32)).IsProperSubsetOf(_1164_v31)));
              let _1166_v33;
              _1166_v33 = _module.D11.create_DC22(_1132_v14);
              let _1167_v34;
              _1167_v34 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (true) : ((_this).f24)),(_1166_v33).dtor_cf35);
              let _1168_v35;
              _1168_v35 = _module.D5.create_DC14((_1140_v20).f39);
              (globalState).f22 = (((_1167_v34).contains(((_1140_v20).fm32((_this).f24, _1161_cf8, globalState)) === ((_1168_v35).dtor_cf23))) ? ((_1167_v34).get(((_1140_v20).fm32((_this).f24, _1161_cf8, globalState)) === ((_1168_v35).dtor_cf23))) : (_1132_v14));
              let _index158 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1159_cf10).length));
              (_1159_cf10)[_index158] = !((_this).f24);
              let _1169_v36;
              _1169_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1137_v17);
              _1161_cf8 = new BigNumber((_1169_v36).length);
            } else {
              let _1170___mcc_h5 = (_source16).cf6;
              let _1171_cf6 = _1170___mcc_h5;
              let _1172_v37;
              let _out36;
              _out36 = (_1140_v20).m4(false, globalState);
              _1172_v37 = _out36;
              _1115_v0 = _1115_v0;
              let _1173_v38;
              _1173_v38 = _dafny.Seq.of(_1115_v0, _1115_v0, _1115_v0, new BigNumber((_1137_v17).length));
              let _1174_v39;
              let _init35 = function (_1175_i5) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              };
              let _nw178 = Array((new BigNumber(7)).toNumber());
              for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw178.length); _i0_35++) {
                _nw178[_i0_35] = _init35(new BigNumber(_i0_35));
              }
              _1174_v39 = _nw178;
              let _1176_v40;
              _1176_v40 = _module.D4.create_DC10(_1137_v17);
              let _1177_v41;
              let _nw179 = new _module.C4();
              _nw179.__ctor(_1115_v0, _module.__default.fm46(_1115_v0, _1173_v38, globalState), _1174_v39, _module.__default.fm34(new BigNumber(((_1176_v40).dtor_cf17).length), (_1140_v20).f39, globalState), p1);
              _1177_v41 = _nw179;
              let _rhs187 = new BigNumber(239);
              let _rhs188 = (_dafny.ZERO).minus((_1177_v41).f33);
              let _lhs139 = globalState;
              _lhs139.f7 = _rhs187;
              _1115_v0 = _rhs188;
            }
          }
        }
      }
      let _1178_v42;
      _1178_v42 = _dafny.Seq.of((_this).f24, false, ((!((_this).f24)) ? (p1) : ((_this).f24)));
      if ((_1178_v42)[_module.__default.safeIndex(new BigNumber(350), new BigNumber((_1178_v42).length))]) {
        let _1179_v43;
        _1179_v43 = _dafny.Set.fromElements((_1140_v20).f39);
        (globalState).f7 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-258)), ((_1180_p0) => function (_1181_i6) {
          return _1180_p0;
        })(p0)), _1137_v17)).length), (_dafny.ZERO).minus(new BigNumber((_1179_v43).length)));
        _1143_v23 = _1143_v23;
        _1135_v15 = _1136_v16;
        _1115_v0 = _1115_v0;
        let _1182_v44;
        let _nw180 = Array((new BigNumber(20)).toNumber()).fill(false);
        _1182_v44 = _nw180;
        let _index159 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1182_v44).length));
        (_1182_v44)[_index159] = (new BigNumber(805)).isLessThan(_1115_v0);
        let _1183_v45;
        _1183_v45 = _dafny.Seq.of(_1115_v0);
        let _1184_v46;
        _1184_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1178_v42).length),(_1183_v45));
        let _1185_v47;
        _1185_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1184_v46,((true) ? ((_this).f24) : ((_this).f24)));
        let _1186_v48;
        _1186_v48 = _dafny.Set.fromElements(_1115_v0, _1115_v0, new BigNumber(50));
        let _index160 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1182_v44).length));
        let _rhs189 = (((_1185_v47).contains(_module.__default.fm47(new BigNumber((_module.__default.fm18(globalState)).cardinality()), globalState))) ? ((_1185_v47).get(_module.__default.fm47(new BigNumber((_module.__default.fm18(globalState)).cardinality()), globalState))) : ((_dafny.Set.fromElements(_module.__default.fm0(p1, _1115_v0, globalState))).IsProperSubsetOf(_1186_v48)));
        let _rhs190 = new _dafny.CodePoint('i'.codePointAt(0));
        let _rhs191 = _1115_v0;
        let _lhs140 = _1182_v44;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1182_v44).length));
        let _lhs142 = _this;
        let _lhs143 = globalState;
        _lhs140[_lhs141] = _rhs189;
        _lhs142.f23 = _rhs190;
        _lhs143.f7 = _rhs191;
      } else {
        _1115_v0 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), ((_1187_v0) => function (_1188_i7) {
          return _dafny.Seq.of(_1187_v0);
        })(_1115_v0))).length);
        let _rhs192 = (_1140_v20).f39;
        let _rhs193 = (_1115_v0).multipliedBy((((_this).f24) ? (_1115_v0) : (new BigNumber(-78))));
        let _rhs194 = _1178_v42;
        let _lhs144 = globalState;
        r2 = _rhs192;
        _lhs144.f7 = _rhs193;
        _1178_v42 = _rhs194;
        let _1189_v49;
        let _nw181 = Array((new BigNumber(18)).toNumber()).fill(_module.D6.Default());
        _1189_v49 = _nw181;
        let _1190_v51;
        _1190_v51 = _module.D6.create_DC15(function () {
  let _coll74 = new _dafny.Map();
  for (const _compr_74 of _dafny.IntegerRange(new BigNumber(327), new BigNumber(336))) {
    let _1191_v50 = _compr_74;
    if (((new BigNumber(327)).isLessThanOrEqualTo(_1191_v50)) && ((_1191_v50).isLessThan(new BigNumber(336)))) {
      _coll74.push([_module.__default.safeDivisionInt(_1191_v50, _1115_v0),_1178_v42]);
    }
  }
  return _coll74;
}());
        let _index161 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1189_v49).length));
        (_1189_v49)[_index161] = _1190_v51;
        let _index162 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1189_v49).length));
        (_1189_v49)[_index162] = _module.__default.fm39(_1115_v0, globalState);
        let _1192_v52;
        _1192_v52 = _dafny.MultiSet.fromElements(_1115_v0);
        let _1193_v53;
        _1193_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1137_v17,_1192_v52);
        let _1194_v54;
        _1194_v54 = _module.D5.create_DC14((_1140_v20).f39);
        let _1195_v55;
        _1195_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(410),!(p1));
        let _1196_v56;
        _1196_v56 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(_1140_v20).f39);
        let _1197_v57;
        _1197_v57 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_1196_v56).contains((_1140_v20).f39)) ? ((_1196_v56).get((_1140_v20).f39)) : ((_1140_v20).f39)));
        let _1198_v58;
        _1198_v58 = _dafny.Set.fromElements(_1192_v52);
        let _1199_v59;
        _1199_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1140_v20).f39,_1198_v58);
        let _1200_v60;
        let _nw182 = Array((new BigNumber(27)).toNumber());
        _nw182[(_dafny.ZERO).toNumber()] = p1;
        _nw182[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_1137_v17, p0);
        _nw182[(new BigNumber(2)).toNumber()] = (((_this).f24) ? ((_1140_v20).f39) : ((_this).f24));
        _nw182[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw182[(new BigNumber(4)).toNumber()] = p1;
        _nw182[(new BigNumber(5)).toNumber()] = (_1192_v52).IsDisjointFrom(_1192_v52);
        _nw182[(new BigNumber(6)).toNumber()] = (_1193_v53).contains(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), ((_1201_p0) => function (_1202_i8) {
          return _1201_p0;
        })(p0)), _module.__default.safeIndex((_dafny.ZERO).minus(_1115_v0), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), ((_1203_p0) => function (_1204_i8) {
          return _1203_p0;
        })(p0))).length)), p0));
        _nw182[(new BigNumber(7)).toNumber()] = p1;
        _nw182[(new BigNumber(8)).toNumber()] = !(_1115_v0).isEqualTo(_1115_v0);
        _nw182[(new BigNumber(9)).toNumber()] = (_1194_v54).dtor_cf23;
        _nw182[(new BigNumber(10)).toNumber()] = (false) === ((_this).f24);
        _nw182[(new BigNumber(11)).toNumber()] = (_1140_v20).f39;
        _nw182[(new BigNumber(12)).toNumber()] = (_this).f24;
        _nw182[(new BigNumber(13)).toNumber()] = (_this).f24;
        _nw182[(new BigNumber(14)).toNumber()] = (_this).f24;
        _nw182[(new BigNumber(15)).toNumber()] = p1;
        _nw182[(new BigNumber(16)).toNumber()] = (_1140_v20).f39;
        _nw182[(new BigNumber(17)).toNumber()] = (_this).f24;
        _nw182[(new BigNumber(18)).toNumber()] = !((_this).f24);
        _nw182[(new BigNumber(19)).toNumber()] = p1;
        _nw182[(new BigNumber(20)).toNumber()] = p1;
        _nw182[(new BigNumber(21)).toNumber()] = (((_this).f24) ? ((_1140_v20).f39) : ((((_1195_v55).contains(_1115_v0)) ? ((_1195_v55).get(_1115_v0)) : ((_1140_v20).f39))));
        _nw182[(new BigNumber(22)).toNumber()] = (((_1197_v57).contains(p1)) ? ((_1197_v57).get(p1)) : (!(false)));
        _nw182[(new BigNumber(23)).toNumber()] = p1;
        _nw182[(new BigNumber(24)).toNumber()] = (_1140_v20).f39;
        _nw182[(new BigNumber(25)).toNumber()] = (new BigNumber(((((_1199_v59).contains((_1140_v20).f39)) ? ((_1199_v59).get((_1140_v20).f39)) : (_1198_v58))).length)).isLessThan(_1115_v0);
        _nw182[(new BigNumber(26)).toNumber()] = (_1140_v20).f39;
        _1200_v60 = _nw182;
        let _index163 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1200_v60).length));
        (_1200_v60)[_index163] = (_1140_v20).f39;
        let _index164 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1200_v60).length));
        (_1200_v60)[_index164] = p1;
        let _1205_v61;
        let _nw183 = Array((new BigNumber(27)).toNumber()).fill([]);
        _1205_v61 = _nw183;
        let _index165 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1205_v61).length));
        (_1205_v61)[_index165] = _1200_v60;
        let _1206_v62;
        _1206_v62 = _dafny.Map.Empty.slice().updateUnsafe(!((_1200_v60)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_1200_v60).length))]),_1200_v60);
        let _index166 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1205_v61).length));
        (_1205_v61)[_index166] = (((_1206_v62).contains(p1)) ? ((_1206_v62).get(p1)) : (_1200_v60));
      }
      r0 = _dafny.Seq.Concat(_1137_v17, _1137_v17);
      let _1207_v63;
      _1207_v63 = _module.D0.create_DC1(new BigNumber(((((_1138_v18).contains(p1)) ? ((_1138_v18).get(p1)) : (_1137_v17))).length));
      r1 = (((_this).f24) ? (_1207_v63) : (_1207_v63));
      r2 = !((_this).f24);
      return [r0, r1, r2];
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _source17 = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("hin")));
      if (_source17.is_DC7) {
        let _1208___mcc_h0 = (_source17).cf13;
        let _1209___mcc_h1 = (_source17).cf14;
        let _1210_cf14 = _1209___mcc_h1;
        let _1211_cf13 = _1208___mcc_h0;
        let _1212_v0;
        let _nw184 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1212_v0 = _nw184;
        _1212_v0 = _1212_v0;
        let _1213_v1;
        let _nw185 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1213_v1 = _nw185;
        _1213_v1 = _1213_v1;
        let _1214_v2;
        _1214_v2 = new BigNumber(486);
        r2 = (_1214_v2).isLessThanOrEqualTo((_dafny.ZERO).minus(_1214_v2));
        (globalState).f7 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(711)), function (_1215_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }), _module.__default.safeIndex(new BigNumber(-356), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(711)), function (_1216_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length)), _this.f23)).length)).multipliedBy((_dafny.ZERO).minus(_1214_v2));
      } else if (_source17.is_DC6) {
        let _1217___mcc_h2 = (_source17).cf12;
        let _1218_cf12 = _1217___mcc_h2;
        let _1219_v3;
        _1219_v3 = _dafny.Seq.of(false);
        let _1220_v4;
        _1220_v4 = new BigNumber(158);
        let _1221_v5;
        _1221_v5 = _dafny.Seq.of(new BigNumber((_1219_v3).length), _1220_v4, new BigNumber(430));
        let _1222_v6;
        _1222_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1221_v5);
        _1222_v6 = (_1222_v6).Merge(_1222_v6);
        let _1223_v7;
        let _nw186 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _1223_v7 = _nw186;
        let _1224_v8;
        _1224_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1220_v4);
        let _index167 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1223_v7).length));
        (_1223_v7)[_index167] = (_1224_v8).Merge(_1224_v8);
        let _index168 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1223_v7).length));
        (_1223_v7)[_index168] = _1224_v8;
        let _1225_v9;
        _1225_v9 = _dafny.Seq.UnicodeFromString("urkqv");
        let _1226_v10;
        _1226_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1220_v4,_1220_v4);
        let _1227_v12;
        let _nw187 = new _module.C0();
        _nw187.__ctor((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1225_v9)).Merge(_1218_cf12), (new BigNumber((function () {
          let _coll75 = new _dafny.Set();
          for (const _compr_75 of (_1226_v10).Keys.Elements) {
            let _1228_v11 = _compr_75;
            if ((_1226_v10).contains(_1228_v11)) {
              _coll75.add((_1228_v11).plus(new BigNumber(476)));
            }
          }
          return _coll75;
        }()).length)).isLessThan(_1220_v4), _this.f23, true);
        _1227_v12 = _nw187;
        _1219_v3 = _1219_v3;
      } else {
        let _1229___mcc_h3 = (_source17).cf15;
        let _1230_cf15 = _1229___mcc_h3;
        let _1231_v13;
        let _nw188 = new _module.C1();
        _nw188.__ctor((_this).f24, (_this).f24);
        _1231_v13 = _nw188;
        let _1232_v14;
        _1232_v14 = _1231_v13;
        let _1233_v15;
        _1233_v15 = _dafny.Seq.of(_1231_v13, _1231_v13, _1231_v13);
        let _1234_v16;
        _1234_v16 = new BigNumber(-984);
        let _source18 = (_1233_v15)[_module.__default.safeIndex(_1234_v16, new BigNumber((_1233_v15).length))];
        let _1235___mcc_h4 = _source18;
        let _1236_cf28 = _1235___mcc_h4;
        let _1237_v17;
        let _nw189 = Array((new BigNumber(22)).toNumber()).fill([]);
        _1237_v17 = _nw189;
        let _1238_v18;
        let _init36 = ((_1239_v16) => function (_1240_i1) {
          return (_1240_i1).minus(new BigNumber((_dafny.Seq.of(_1239_v16)).length));
        })(_1234_v16);
        let _nw190 = Array((new BigNumber(8)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw190.length); _i0_36++) {
          _nw190[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1238_v18 = _nw190;
        let _index169 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1237_v17).length));
        (_1237_v17)[_index169] = _1238_v18;
        let _index170 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1237_v17).length));
        (_1237_v17)[_index170] = _1238_v18;
        let _1241_v19;
        _1241_v19 = _dafny.Seq.of(!((_1236_cf28).f39), (_1231_v13).f39);
        let _1242_v20;
        _1242_v20 = _dafny.Seq.of(_1241_v19);
        let _1243_v21;
        let _1244_v22;
        let _1245_v23;
        let _out37;
        let _out38;
        let _out39;
        let _outcollector11 = _module.__default.m0(_1234_v16, _dafny.MultiSet.fromElements((_this).f24), ((_this).f24) === ((_1236_cf28).f39), new BigNumber((_1242_v20).length), globalState);
        _out37 = _outcollector11[0];
        _out38 = _outcollector11[1];
        _out39 = _outcollector11[2];
        _1243_v21 = _out37;
        _1244_v22 = _out38;
        _1245_v23 = _out39;
        let _1246_v24;
        let _nw191 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1246_v24 = _nw191;
        let _index171 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_1246_v24).length));
        (_1246_v24)[_index171] = _1244_v22;
        let _index172 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1246_v24).length));
        (_1246_v24)[_index172] = _1244_v22;
        let _1247_v25;
        _1247_v25 = _dafny.Seq.of(_1243_v21);
        let _index173 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_1246_v24).length));
        let _index174 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1246_v24).length));
        let _rhs195 = (_this).f24;
        let _rhs196 = _dafny.Seq.Concat(_1242_v20, _dafny.Seq.of(_1241_v19));
        let _rhs197 = ((_1244_v22) ? (!_dafny.areEqual(_1247_v25, _1247_v25)) : (true));
        let _lhs145 = _1246_v24;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_1246_v24).length));
        let _lhs147 = _1246_v24;
        let _lhs148 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1246_v24).length));
        _lhs145[_lhs146] = _rhs195;
        _1242_v20 = _rhs196;
        _lhs147[_lhs148] = _rhs197;
        let _1248_v26;
        _1248_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1244_v22,_1244_v22);
        let _1249_v27;
        _1249_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v19,(_1236_cf28).f39);
        let _1250_v28;
        _1250_v28 = _dafny.Map.Empty.slice().updateUnsafe((((_1248_v26).contains((_1231_v13).f39)) ? ((_1248_v26).get((_1231_v13).f39)) : ((((_1249_v27).contains(_1241_v19)) ? ((_1249_v27).get(_1241_v19)) : (!((_1231_v13).f39))))),_1243_v21);
        let _1251_v29;
        _1251_v29 = _dafny.Seq.of(_1250_v28, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),(_1231_v13).fm33(globalState)));
        _1250_v28 = (_1251_v29)[_module.__default.safeIndex(_1243_v21, new BigNumber((_1251_v29).length))];
        let _1252_v30;
        _1252_v30 = _dafny.Seq.UnicodeFromString("ycfwea");
        let _1253_v31;
        _1253_v31 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f24),_1252_v30);
        let _1254_v32;
        _1254_v32 = _module.D2.create_DC6(_1253_v31);
        let _source19 = _1254_v32;
        if (_source19.is_DC7) {
          let _1255___mcc_h5 = (_source19).cf13;
          let _1256___mcc_h6 = (_source19).cf14;
          let _1257_cf14 = _1256___mcc_h6;
          let _1258_cf13 = _1255___mcc_h5;
          let _1259_v33;
          _1259_v33 = _dafny.Seq.of(_1258_cf13);
          let _1260_v34;
          _1260_v34 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_1259_v33);
          _1260_v34 = (_1260_v34).update(_1258_cf13, _1259_v33);
          let _1261_v35;
          _1261_v35 = _dafny.Map.Empty.slice().updateUnsafe((_1231_v13).f39,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1252_v30).length),(_1231_v13).f39));
          let _1262_v36;
          _1262_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1234_v16,_module.__default.fm2(globalState));
          _1261_v35 = (_1261_v35).update((_this).f24, _1262_v36);
          let _1263_v37;
          _1263_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1260_v34).length),_1234_v16);
          let _1264_v38;
          _1264_v38 = _dafny.Seq.of(_1234_v16);
          let _1265_v39;
          _1265_v39 = _dafny.Set.fromElements((_this).f24, (_module.D5.create_DC13(_1234_v16, (_1231_v13).f39, (_1259_v33)[_module.__default.safeIndex(_1234_v16, new BigNumber((_1259_v33).length))])).dtor_cf21);
          let _1266_v40;
          _1266_v40 = _dafny.Seq.of(_dafny.Set.fromElements((_1231_v13).f39, (_1231_v13).f39), _1265_v39);
          let _1267_v41;
          let _nw192 = Array((new BigNumber(23)).toNumber());
          _nw192[(_dafny.ZERO).toNumber()] = _1234_v16;
          _nw192[(_dafny.ONE).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1234_v16), _1234_v16);
          _nw192[(new BigNumber(3)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(4)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(5)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(6)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(7)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(_1234_v16, (((_1263_v37).contains(new BigNumber((_1264_v38).length))) ? ((_1263_v37).get(new BigNumber((_1264_v38).length))) : (_1234_v16)));
          _nw192[(new BigNumber(9)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(10)).toNumber()] = (_this).fm6(globalState);
          _nw192[(new BigNumber(11)).toNumber()] = (_1234_v16).multipliedBy(new BigNumber(((_1266_v40)[_module.__default.safeIndex(_1234_v16, new BigNumber((_1266_v40).length))]).length));
          _nw192[(new BigNumber(12)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_1234_v16);
          _nw192[(new BigNumber(14)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(15)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(16)).toNumber()] = (_1264_v38)[_module.__default.safeIndex(_1234_v16, new BigNumber((_1264_v38).length))];
          _nw192[(new BigNumber(17)).toNumber()] = new BigNumber((_1262_v36).length);
          _nw192[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1234_v16));
          _nw192[(new BigNumber(19)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(20)).toNumber()] = _1234_v16;
          _nw192[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1258_cf13,_1234_v16)).length), _1234_v16));
          _nw192[(new BigNumber(22)).toNumber()] = (_1234_v16).plus(new BigNumber((_1252_v30).length));
          _1267_v41 = _nw192;
          let _index175 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1267_v41).length));
          (_1267_v41)[_index175] = new BigNumber((_1263_v37).length);
          let _1268_v42;
          _1268_v42 = _dafny.MultiSet.fromElements(_this.f23);
          let _1269_v43;
          _1269_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1258_cf13,false);
          let _1270_v44;
          _1270_v44 = _dafny.MultiSet.fromElements(_1269_v43, _1269_v43, _1269_v43);
          let _1271_v45;
          _1271_v45 = _dafny.MultiSet.fromElements(_1270_v44, _1270_v44, (_dafny.MultiSet.fromElements(_1269_v43)).update(_1269_v43, _module.__default.abs(_1234_v16)), _1270_v44, _1270_v44);
          let _index176 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1267_v41).length));
          let _rhs198 = _1234_v16;
          let _rhs199 = (((_1268_v42).contains(_this.f23)) ? ((_1268_v42).get(_this.f23)) : ((((_1271_v45).contains(_1270_v44)) ? ((_1271_v45).get(_1270_v44)) : (_1234_v16))));
          let _rhs200 = ((new BigNumber(821)).multipliedBy(_module.__default.fm0((_1231_v13).f39, _1234_v16, globalState))).multipliedBy(_1234_v16);
          let _lhs149 = _1267_v41;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1267_v41).length));
          let _lhs151 = globalState;
          let _lhs152 = globalState;
          _lhs149[_lhs150] = _rhs198;
          _lhs151.f7 = _rhs199;
          _lhs152.f7 = _rhs200;
          (globalState).f7 = _module.__default.fm0((_1231_v13).f39, _1234_v16, globalState);
        } else if (_source19.is_DC6) {
          let _1272___mcc_h7 = (_source19).cf12;
          let _1273_cf12 = _1272___mcc_h7;
          let _1274_v46;
          _1274_v46 = _dafny.MultiSet.fromElements((_1231_v13).f39, (_this).f24);
          _1274_v46 = _1274_v46;
          let _1275_v48;
          _1275_v48 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of _dafny.IntegerRange(new BigNumber(704), new BigNumber(468))) {
              let _1276_v47 = _compr_76;
              if (((new BigNumber(704)).isLessThanOrEqualTo(_1276_v47)) && ((_1276_v47).isLessThan(new BigNumber(468)))) {
                _coll76.push([(_1276_v47).multipliedBy(_1234_v16),_1234_v16]);
              }
            }
            return _coll76;
          }(),new BigNumber(39));
          let _1277_v49;
          _1277_v49 = _dafny.MultiSet.fromElements(_1234_v16, _1234_v16, _1234_v16);
          let _1278_v50;
          _1278_v50 = _dafny.Set.fromElements((_1231_v13).f39);
          let _nw193 = Array((new BigNumber(19)).toNumber());
          _nw193[(_dafny.ZERO).toNumber()] = _1234_v16;
          _nw193[(_dafny.ONE).toNumber()] = new BigNumber(31);
          _nw193[(new BigNumber(2)).toNumber()] = (_1234_v16).plus(_1234_v16);
          _nw193[(new BigNumber(3)).toNumber()] = new BigNumber((_1275_v48).length);
          _nw193[(new BigNumber(4)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(5)).toNumber()] = (((_1277_v49).contains(_1234_v16)) ? ((_1277_v49).get(_1234_v16)) : (_1234_v16));
          _nw193[(new BigNumber(6)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(7)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(8)).toNumber()] = (_1234_v16).minus(_1234_v16);
          _nw193[(new BigNumber(9)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_1234_v16, _1234_v16);
          _nw193[(new BigNumber(11)).toNumber()] = (_1234_v16).plus(_module.__default.fm0((_1231_v13).f39, _1234_v16, globalState));
          _nw193[(new BigNumber(12)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(13)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(14)).toNumber()] = (_1234_v16).plus((_dafny.ZERO).minus(_1234_v16));
          _nw193[(new BigNumber(15)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(16)).toNumber()] = _1234_v16;
          _nw193[(new BigNumber(17)).toNumber()] = _module.__default.fm0((_this).f24, new BigNumber((_1278_v50).length), globalState);
          _nw193[(new BigNumber(18)).toNumber()] = _1234_v16;
          (globalState).f22 = _nw193;
          (globalState).f7 = _module.__default.safeModuloInt(_1234_v16, _1234_v16);
          let _1279_v51;
          _1279_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(105),_1234_v16);
          let _1280_v52;
          let _nw194 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1280_v52 = _nw194;
          let _1281_v53;
          _1281_v53 = _module.D12.create_DC24(_1280_v52);
          let _1282_v54;
          _1282_v54 = _module.D0.create_DC1(new BigNumber((_1274_v46).cardinality()));
          let _1283_v55;
          let _nw195 = new _module.C4();
          _nw195.__ctor((_1234_v16).minus(_1234_v16), _1279_v51, (_1281_v53).dtor_cf40, _module.__default.fm34((_1282_v54).dtor_cf1, false, globalState), (_1231_v13).f39);
          _1283_v55 = _nw195;
        } else {
          let _1284___mcc_h8 = (_source19).cf15;
          let _1285_cf15 = _1284___mcc_h8;
          let _1286_v56;
          _1286_v56 = _dafny.Map.Empty.slice().updateUnsafe((_1231_v13).f39,(_1231_v13).f39);
          _1286_v56 = (_1286_v56).update((_1231_v13).f39, (_1231_v13).f39);
          let _1287_v59;
          _1287_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of _dafny.IntegerRange(new BigNumber(183), new BigNumber(311))) {
              let _1288_v57 = _compr_77;
              if (((new BigNumber(183)).isLessThanOrEqualTo(_1288_v57)) && ((_1288_v57).isLessThan(new BigNumber(311)))) {
                _coll77.push([(_1288_v57).multipliedBy(new BigNumber((function () {
                  let _coll78 = new _dafny.Map();
                  for (const _compr_78 of _dafny.IntegerRange(new BigNumber(-155), new BigNumber(293))) {
                    let _1289_v58 = _compr_78;
                    if (((new BigNumber(-155)).isLessThanOrEqualTo(_1289_v58)) && ((_1289_v58).isLessThan(new BigNumber(293)))) {
                      _coll78.push([_module.__default.safeDivisionInt(_1289_v58, _1234_v16),_1286_v56]);
                    }
                  }
                  return _coll78;
                }()).length)),_1234_v16]);
              }
            }
            return _coll77;
          }()).length),_dafny.Seq.update(_module.__default.fm1((_this).f24, globalState), _module.__default.safeIndex(_1234_v16, new BigNumber((_module.__default.fm1((_this).f24, globalState)).length)), _this.f23));
          _1252_v30 = (((_1287_v59).contains(_1234_v16)) ? ((_1287_v59).get(_1234_v16)) : (_1252_v30));
          _1252_v30 = _1252_v30;
          let _1290_v60;
          let _nw196 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1290_v60 = _nw196;
          let _1291_v61;
          let _nw197 = new _module.C5();
          _nw197.__ctor(_1290_v60);
          _1291_v61 = _nw197;
        }
        (globalState).f13 = (_this).f24;
        let _1292_v62;
        _1292_v62 = _dafny.MultiSet.fromElements(_1234_v16, _1234_v16);
        let _1293_v63;
        _1293_v63 = _dafny.Map.Empty.slice().updateUnsafe((_1234_v16).isLessThan((_dafny.ZERO).minus(new BigNumber((_1292_v62).cardinality()))),_1232_v14);
        _1293_v63 = (_1293_v63).update(((_1231_v13).f39) === ((_this).f24), _1232_v14);
      }
      let _1294_v64;
      _1294_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
      let _1295_v65;
      let _nw198 = Array((new BigNumber(24)).toNumber());
      _nw198[(_dafny.ZERO).toNumber()] = (_1294_v64).Merge(_1294_v64);
      _nw198[(_dafny.ONE).toNumber()] = (_1294_v64).Merge(_1294_v64);
      _nw198[(new BigNumber(2)).toNumber()] = (_module.__default.fm19(globalState)).Merge(_1294_v64);
      _nw198[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f24),(_this).f24);
      _nw198[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).fm16(globalState),(_this).f24);
      _nw198[(new BigNumber(5)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(6)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
      _nw198[(new BigNumber(8)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(9)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(10)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(11)).toNumber()] = (_1294_v64).Merge(_1294_v64);
      _nw198[(new BigNumber(12)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
      _nw198[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,false);
      _nw198[(new BigNumber(15)).toNumber()] = (_1294_v64).Merge(_1294_v64);
      _nw198[(new BigNumber(16)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,!((_this).f24))).update((_this).f24, (_this).f24);
      _nw198[(new BigNumber(17)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(18)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
      _nw198[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
      _nw198[(new BigNumber(21)).toNumber()] = (_1294_v64).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24));
      _nw198[(new BigNumber(22)).toNumber()] = _1294_v64;
      _nw198[(new BigNumber(23)).toNumber()] = (_1294_v64).Merge(_1294_v64);
      _1295_v65 = _nw198;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1295_v65).length))) {
        let _1296_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1296_i2)) && ((_1296_i2).isLessThan(new BigNumber((_1295_v65).length))))) {
          (_1295_v65)[(_1296_i2)] = (_1294_v64).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24));
        }
      }
      let _1297_v66;
      _1297_v66 = new BigNumber(590);
      let _1298_v67;
      _1298_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1297_v66,_1297_v66);
      let _1299_v68;
      _1299_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1297_v66,_1298_v67);
      let _1300_v69;
      _1300_v69 = _dafny.Seq.of(new BigNumber(((((_1299_v68).contains(_1297_v66)) ? ((_1299_v68).get(_1297_v66)) : (_1298_v67))).length), _1297_v66, _1297_v66);
      _1300_v69 = _dafny.Seq.Concat(_1300_v69, _1300_v69);
      r1 = _1297_v66;
      let _1301_v70;
      _1301_v70 = _dafny.Seq.of(!(true));
      let _1302_v71;
      let _nw199 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1302_v71 = _nw199;
      let _1303_v72;
      _1303_v72 = _dafny.MultiSet.fromElements(_this.f23);
      let _1304_v73;
      _1304_v73 = _dafny.Set.fromElements((((_1303_v72).contains(_this.f23)) ? ((_1303_v72).get(_this.f23)) : ((_this).fm6(globalState))), new BigNumber(-668));
      let _1305_v74;
      let _nw200 = new _module.C2();
      _nw200.__ctor(_dafny.MultiSet.FromArray(_1301_v70), (((_this).f24) ? (_1302_v71) : (_1302_v71)), _this.f23, (new BigNumber((_1304_v73).length)).isLessThan(_1297_v66));
      _1305_v74 = _nw200;
      r0 = true;
      r0 = (_this).f24;
      r1 = _1297_v66;
      r2 = (_this).f24;
      return [r0, r1, r2];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
      this._f25 = [];
      this._f27 = false;
      this._f30 = false;
      this._f31 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T3, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f30, f31, f25, f23, f24, f27) {
      let _this = this;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f25 = f25;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f27 = f27;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(564)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), function (_1306_i0) {
        return _this.f23;
      })).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f24))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!((_this).f30))).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), function (_1307_i1) {
        return new BigNumber(802);
      })));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return ((_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_this).f30,_dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), function (_1308_i0) {
  return _this.f23;
})))).dtor_cf12).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("y"))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("sufkeaph"))));
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.of(!(!((_this).f30)))).length)).isLessThan(new BigNumber(252));
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus((((_this).f30) ? (new BigNumber(-200)) : (new BigNumber(-98))))), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(891))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ych")).length),new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-903),new BigNumber(-298)))).length))).length)));
    };
    fm15(p0, globalState) {
      let _this = this;
      return _dafny.Seq.of((((_this).f31) ? (_dafny.Seq.of((_this).f27)) : (_dafny.Seq.of(true, (_this).f24, (_this).f24))));
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1309_v0;
      let _init37 = function (_1310_i0) {
        return (_this).f30;
      };
      let _nw201 = Array((new BigNumber(11)).toNumber());
      for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw201.length); _i0_37++) {
        _nw201[_i0_37] = _init37(new BigNumber(_i0_37));
      }
      _1309_v0 = _nw201;
      let _1311_v1;
      _1311_v1 = new BigNumber(235);
      let _1312_v2;
      _1312_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1313_v3;
      let _nw202 = Array((new BigNumber(10)).toNumber());
      _nw202[(_dafny.ZERO).toNumber()] = _1311_v1;
      _nw202[(_dafny.ONE).toNumber()] = new BigNumber(441);
      _nw202[(new BigNumber(2)).toNumber()] = new BigNumber(-229);
      _nw202[(new BigNumber(3)).toNumber()] = _1311_v1;
      _nw202[(new BigNumber(4)).toNumber()] = _1311_v1;
      _nw202[(new BigNumber(5)).toNumber()] = _1311_v1;
      _nw202[(new BigNumber(6)).toNumber()] = _1311_v1;
      _nw202[(new BigNumber(7)).toNumber()] = new BigNumber(459);
      _nw202[(new BigNumber(8)).toNumber()] = new BigNumber(855);
      _nw202[(new BigNumber(9)).toNumber()] = new BigNumber((_1312_v2).length);
      _1313_v3 = _nw202;
      let _1314_v4;
      _1314_v4 = _dafny.Seq.of(_1313_v3);
      let _index177 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length));
      (_1309_v0)[_index177] = _dafny.areEqual(_1314_v4, _1314_v4);
      let _1315_v5;
      _1315_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _1316_v6;
      _1316_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1315_v5).length),_1311_v1);
      let _1317_v7;
      _1317_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1316_v6);
      let _1318_v8;
      _1318_v8 = (((_1317_v7).contains((_this).f31)) ? ((_1317_v7).get((_this).f31)) : (_1316_v6));
      let _index178 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length));
      (_1309_v0)[_index178] = function (_source20) {
        let _1319___mcc_h0 = _source20;
        let _1320_cf16 = _1319___mcc_h0;
        return !((_this).f24) || ((_this).f30);
      }(_1318_v8);
      let _1321_v9;
      let _nw203 = Array((new BigNumber(5)).toNumber()).fill([]);
      _1321_v9 = _nw203;
      let _index179 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_1321_v9).length));
      (_1321_v9)[_index179] = _1309_v0;
      let _1322_v10;
      _1322_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v1,false);
      let _1323_v11;
      _1323_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1322_v10,_1309_v0);
      let _index180 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_1321_v9).length));
      (_1321_v9)[_index180] = (((_1323_v11).contains(_1322_v10)) ? ((_1323_v11).get(_1322_v10)) : (_1309_v0));
      let _1324_v12;
      let _nw204 = Array((new BigNumber(12)).toNumber()).fill([]);
      _1324_v12 = _nw204;
      let _1325_v13;
      let _nw205 = Array((new BigNumber(14)).toNumber());
      _nw205[(_dafny.ZERO).toNumber()] = _1313_v3;
      _nw205[(_dafny.ONE).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(2)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(3)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(4)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(5)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(6)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(7)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(8)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(9)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(10)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(11)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(12)).toNumber()] = _1313_v3;
      _nw205[(new BigNumber(13)).toNumber()] = _1313_v3;
      _1325_v13 = _nw205;
      let _index181 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1324_v12).length));
      (_1324_v12)[_index181] = _1325_v13;
      let _1326_v14;
      let _nw206 = new _module.C6();
      _nw206.__ctor(_this.f23, true);
      _1326_v14 = _nw206;
      let _index182 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1324_v12).length));
      let _index183 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length));
      let _rhs201 = (((_this).f30) ? ((((_this).f30) ? (_1325_v13) : (_1325_v13))) : (_1325_v13));
      let _rhs202 = (_1311_v1).isLessThan(new BigNumber(72));
      let _rhs203 = _1326_v14;
      let _lhs153 = _1324_v12;
      let _lhs154 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1324_v12).length));
      let _lhs155 = _1309_v0;
      let _lhs156 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length));
      _lhs153[_lhs154] = _rhs201;
      _lhs155[_lhs156] = _rhs202;
      _1326_v14 = _rhs203;
      let _1327_v15;
      _1327_v15 = _dafny.Set.fromElements(_1311_v1);
      let _1328_v16;
      _1328_v16 = _module.D5.create_DC12((_1327_v15).Union(_1327_v15));
      let _source21 = _1328_v16;
      if (_source21.is_DC13) {
        let _1329___mcc_h1 = (_source21).cf20;
        let _1330___mcc_h2 = (_source21).cf21;
        let _1331___mcc_h3 = (_source21).cf22;
        let _1332_cf22 = _1331___mcc_h3;
        let _1333_cf21 = _1330___mcc_h2;
        let _1334_cf20 = _1329___mcc_h1;
        let _1335_v17;
        _1335_v17 = _dafny.MultiSet.fromElements(p0);
        r1 = _module.__default.safeModuloInt(_1311_v1, (new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_1336_i1) {
          return new BigNumber(48);
        }), _module.__default.safeIndex(_1334_cf20, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_1337_i1) {
          return new BigNumber(48);
        })).length)), new BigNumber((_dafny.Set.fromElements(_1334_cf20, new BigNumber((_1335_v17).cardinality()), _1334_cf20)).length))).length)).multipliedBy((_1326_v14).fm6(globalState)));
        if (!((_this).f24)) {
          _1313_v3 = _1313_v3;
          let _1338_v18;
          let _nw207 = new _module.C1();
          _nw207.__ctor(!((((_1326_v14).f24) ? (_1333_cf21) : (_1333_cf21))), p0);
          _1338_v18 = _nw207;
          let _1339_v19;
          _1339_v19 = _module.D12.create_DC25(((_this).f24) && ((_1326_v14).f24), _1326_v14.f23);
          _1339_v19 = _1339_v19;
          let _1340_v20;
          _1340_v20 = _dafny.Seq.UnicodeFromString("rwgqydm");
          let _1341_v21;
          _1341_v21 = _dafny.Set.fromElements((_1326_v14).f24);
          let _1342_v22;
          _1342_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1340_v20,_1341_v21);
          _1342_v22 = (_1342_v22).update(_1340_v20, _1341_v21);
          _1332_cf22 = (_this).f31;
        } else {
          (globalState).f22 = _1313_v3;
          let _1343_v23;
          _1343_v23 = _dafny.Seq.UnicodeFromString("jlmrjcjpx");
          let _1344_v24;
          _1344_v24 = _module.D5.create_DC14(_module.__default.fm2(globalState));
          let _1345_v25;
          _1345_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1344_v24,_module.__default.fm0(_1333_cf21, _1334_cf20, globalState));
          let _1346_v26;
          _1346_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_1343_v23, _1326_v14.f23),_1345_v25);
          let _1347_v28;
          _1347_v28 = _dafny.Set.fromElements(_1344_v24);
          _1346_v26 = (_1346_v26).update(_1333_cf21, function () {
            let _coll79 = new _dafny.Map();
            for (const _compr_79 of (_1347_v28).Elements) {
              let _1348_v27 = _compr_79;
              if ((_1347_v28).contains(_1348_v27)) {
                _coll79.push([_1348_v27,_1334_cf20]);
              }
            }
            return _coll79;
          }());
          _1334_cf20 = _1334_cf20;
          let _1349_v29;
          _1349_v29 = _dafny.Seq.of(_1311_v1, _1334_cf20, _1334_cf20, _1311_v1);
          let _index184 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1313_v3).length));
          (_1313_v3)[_index184] = new BigNumber((_1349_v29).length);
          let _1350_v30;
          _1350_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v1,_1313_v3);
          let _1351_v31;
          _1351_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1312_v2,_1333_cf21);
          let _index185 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1313_v3).length));
          let _rhs204 = (_1311_v1).isEqualTo((_dafny.ZERO).minus(_1334_cf20));
          let _rhs205 = (_1334_cf20).plus(_module.__default.fm0(_1333_cf21, _1311_v1, globalState));
          let _rhs206 = (((_1350_v30).contains(new BigNumber((_1351_v31).length))) ? ((_1350_v30).get(new BigNumber((_1351_v31).length))) : (((false) ? (_1313_v3) : (_1313_v3))));
          let _rhs207 = (_dafny.ZERO).minus(_1311_v1);
          let _lhs157 = globalState;
          let _lhs158 = globalState;
          let _lhs159 = globalState;
          let _lhs160 = _1313_v3;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1313_v3).length));
          _lhs157.f13 = _rhs204;
          _lhs158.f7 = _rhs205;
          _lhs159.f22 = _rhs206;
          _lhs160[_lhs161] = _rhs207;
          _1333_cf21 = !(_1334_cf20).isEqualTo((_1313_v3)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_1313_v3).length))]);
        }
        _1333_cf21 = (_this).f30;
        let _index186 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length));
        (_1309_v0)[_index186] = (_1309_v0)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length))];
      } else if (_source21.is_DC14) {
        let _1352___mcc_h4 = (_source21).cf23;
        let _1353_cf23 = _1352___mcc_h4;
        let _index187 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1313_v3).length));
        (_1313_v3)[_index187] = _module.__default.safeModuloInt(_1311_v1, _module.__default.fm0((_this).f24, new BigNumber(262), globalState));
        let _1354_v32;
        _1354_v32 = _dafny.MultiSet.fromElements((_this).f24);
        let _index188 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1313_v3).length));
        (_1313_v3)[_index188] = (((_1354_v32).contains(false)) ? ((_1354_v32).get(false)) : (_1311_v1));
        let _1355_v33;
        _1355_v33 = _dafny.Seq.UnicodeFromString("dqbli");
        _1355_v33 = _dafny.Seq.Concat(_1355_v33, _1355_v33);
        let _1356_v34;
        let _nw208 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
        _1356_v34 = _nw208;
        let _1357_v35;
        _1357_v35 = _dafny.Map.Empty.slice().updateUnsafe((_1313_v3)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1313_v3).length))],(_this).f25);
        let _index189 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1356_v34).length));
        (_1356_v34)[_index189] = _1357_v35;
        let _index190 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1356_v34).length));
        (_1356_v34)[_index190] = _1357_v35;
        _1355_v33 = _dafny.Seq.update(_module.__default.fm1((_1309_v0)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length))], globalState), _module.__default.safeIndex(_module.__default.safeModuloInt((_1313_v3)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1313_v3).length))], _1311_v1), new BigNumber((_module.__default.fm1((_1309_v0)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length))], globalState)).length)), _1326_v14.f23);
      } else {
        let _1358___mcc_h5 = (_source21).cf19;
        let _1359_cf19 = _1358___mcc_h5;
        let _1360_v36;
        _1360_v36 = _dafny.Seq.of((_1321_v9)[_module.__default.safeIndex(new BigNumber(97), new BigNumber((_1321_v9).length))], _1309_v0, (_1321_v9)[_module.__default.safeIndex(new BigNumber(97), new BigNumber((_1321_v9).length))], (_1321_v9)[_module.__default.safeIndex(new BigNumber(97), new BigNumber((_1321_v9).length))]);
        _1360_v36 = _dafny.Seq.of(_1309_v0, (((_1326_v14).f24) ? ((_1360_v36)[_module.__default.safeIndex(_1311_v1, new BigNumber((_1360_v36).length))]) : (_1309_v0)), (_1321_v9)[_module.__default.safeIndex(new BigNumber(97), new BigNumber((_1321_v9).length))]);
        let _index191 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length));
        (_1309_v0)[_index191] = (_1326_v14).f24;
        let _1361_v37;
        _1361_v37 = _dafny.Seq.of(_1311_v1);
        _1361_v37 = _1361_v37;
        let _1362_v38;
        _1362_v38 = _dafny.Seq.of((_this).f24, !((_this).f31), (_1326_v14).f24, (_this).f24);
        let _1363_v39;
        _1363_v39 = _module.D4.create_DC11(_1311_v1);
        let _rhs208 = (_this).fm4((new BigNumber((_dafny.Seq.UnicodeFromString("dupgk")).length)).multipliedBy(_1311_v1), _1362_v38, globalState);
        let _rhs209 = (_1363_v39).dtor_cf18;
        let _rhs210 = (_1362_v38)[_module.__default.safeIndex(_1311_v1, new BigNumber((_1362_v38).length))];
        let _rhs211 = p0;
        let _lhs162 = globalState;
        let _lhs163 = globalState;
        _lhs162.f13 = _rhs208;
        _1311_v1 = _rhs209;
        _lhs163.f13 = _rhs210;
        r2 = _rhs211;
      }
      _1316_v6 = (_1316_v6).update(_1311_v1, _1311_v1);
      if ((((_1309_v0)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length))]) ? ((((_1322_v10).contains(new BigNumber(-678))) ? ((_1322_v10).get(new BigNumber(-678))) : ((_1326_v14).f24))) : (p0))) {
        (globalState).f7 = _1311_v1;
        (_this).f23 = (((_this).f31) ? (_1326_v14.f23) : (_1326_v14.f23));
        let _1364_v40;
        _1364_v40 = _dafny.Map.Empty.slice().updateUnsafe((_1326_v14).f24,_1311_v1);
        let _1365_v41;
        _1365_v41 = _dafny.Seq.of(new BigNumber((_1364_v40).length));
        r1 = new BigNumber((_dafny.Seq.Concat(_1365_v41, _1365_v41)).length);
        let _1366_v42;
        _1366_v42 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState));
        r1 = (new BigNumber(((_1366_v42).Union((_1366_v42).update(p0, _module.__default.abs(_1311_v1)))).cardinality())).plus((_dafny.ZERO).minus(_1311_v1));
        let _index192 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1313_v3).length));
        (_1313_v3)[_index192] = new BigNumber(342);
        let _index193 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1313_v3).length));
        (_1313_v3)[_index193] = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_1311_v1, _1311_v1), _1311_v1);
      } else {
        let _1367_v43;
        _1367_v43 = _module.D5.create_DC14((_1326_v14).f24);
        _1367_v43 = _1367_v43;
        let _1368_v44;
        _1368_v44 = _dafny.Seq.UnicodeFromString("qy");
        let _1369_v45;
        _1369_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1368_v44);
        let _1370_v46;
        _1370_v46 = _dafny.Seq.of((_1309_v0)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1309_v0).length))], (_this).f30);
        let _1371_v47;
        _1371_v47 = _dafny.Set.fromElements(_1368_v44, (((_1369_v45).contains((_1326_v14).fm4(_1311_v1, _1370_v46, globalState))) ? ((_1369_v45).get((_1326_v14).fm4(_1311_v1, _1370_v46, globalState))) : (_1368_v44)));
        let _1372_v48;
        _1372_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1368_v44,(_this).f24);
        _1371_v47 = (function () {
          let _coll80 = new _dafny.Set();
          for (const _compr_80 of (_1372_v48).Keys.Elements) {
            let _1373_v49 = _compr_80;
            if ((_1372_v48).contains(_1373_v49)) {
              _coll80.add(_1373_v49);
            }
          }
          return _coll80;
        }()).Difference(_1371_v47);
        _1368_v44 = _1368_v44;
        let _1374_v50;
        _1374_v50 = _dafny.Set.fromElements((_1370_v46)[_module.__default.safeIndex(_1311_v1, new BigNumber((_1370_v46).length))], (_this).f30);
        let _1375_v51;
        _1375_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_1374_v50);
        _1375_v51 = (_1375_v51).update(false, _1374_v50);
        (globalState).f7 = (_1311_v1).plus((_dafny.ZERO).minus(_1311_v1));
      }
      let _1376_v52;
      _1376_v52 = _dafny.MultiSet.fromElements(false);
      let _1377_v53;
      _1377_v53 = _dafny.Seq.of(_1376_v52);
      let _1378_v54;
      _1378_v54 = _dafny.Seq.of(_1311_v1, _1311_v1);
      r0 = (_dafny.Set.fromElements((_1377_v53)[_module.__default.safeIndex(new BigNumber((_1378_v54).length), new BigNumber((_1377_v53).length))], _dafny.MultiSet.fromElements((_1326_v14).f24))).IsDisjointFrom(function () {
        let _coll81 = new _dafny.Set();
        for (const _compr_81 of (_dafny.MultiSet.FromArray(_1377_v53)).Elements) {
          let _1379_v55 = _compr_81;
          if ((_dafny.MultiSet.FromArray(_1377_v53)).contains(_1379_v55)) {
            _coll81.add(_1379_v55);
          }
        }
        return _coll81;
      }());
      let _1380_v57;
      _1380_v57 = _dafny.Set.fromElements(_module.D5.create_DC12(function () {
  let _coll82 = new _dafny.Set();
  for (const _compr_82 of _dafny.IntegerRange(new BigNumber(-668), new BigNumber(-506))) {
    let _1381_v56 = _compr_82;
    if (((new BigNumber(-668)).isLessThanOrEqualTo(_1381_v56)) && ((_1381_v56).isLessThan(new BigNumber(-506)))) {
      _coll82.add((_1381_v56).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), function (_1382_i2) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)));
    }
  }
  return _coll82;
}()), _1328_v16);
      let _1383_v58;
      _1383_v58 = _dafny.Set.fromElements(_1380_v57);
      r1 = new BigNumber((_1383_v58).length);
      r2 = p0;
      return [r0, r1, r2];
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f13 = (_this).f24;
      let _1384_v0;
      _1384_v0 = _dafny.MultiSet.fromElements((_this).f30, (_this).f24);
      let _1385_v1;
      _1385_v1 = _module.D0.create_DC0(_1384_v0);
      let _1386_v2;
      let _nw209 = new _module.C2();
      _nw209.__ctor((((_this).f31) ? (_1384_v0) : ((_1385_v1).dtor_cf0)), (_this).f25, _this.f23, (_this).f24);
      _1386_v2 = _nw209;
      let _1387_i0;
      _1387_i0 = _dafny.ZERO;
      L6: {
        while (p0) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1387_i0)) {
              break L6;
            }
            _1387_i0 = (_1387_i0).plus(_dafny.ONE);
            let _1388_v3;
            let _nw210 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _1388_v3 = _nw210;
            let _1389_v4;
            _1389_v4 = _dafny.Seq.UnicodeFromString("gtxxn");
            let _index194 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_1388_v3).length));
            (_1388_v3)[_index194] = _1389_v4;
            let _1390_v5;
            _1390_v5 = new BigNumber(924);
            let _1391_v6;
            _1391_v6 = _dafny.Seq.of((_this).f24);
            let _index195 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_1388_v3).length));
            (_1388_v3)[_index195] = _dafny.Seq.update(_1389_v4, _module.__default.safeIndex((_module.__default.fm0((_this).f27, _1390_v5, globalState)).multipliedBy(new BigNumber((_1391_v6).length)), new BigNumber((_1389_v4).length)), _this.f23);
            let _1392_v7;
            _1392_v7 = _dafny.Seq.of(_module.__default.fm0(true, _1390_v5, globalState));
            let _1393_v8;
            _1393_v8 = _dafny.Seq.of((_1392_v7)[_module.__default.safeIndex(new BigNumber((_1392_v7).length), new BigNumber((_1392_v7).length))]);
            _1393_v8 = ((p0) ? (_1393_v8) : (_1393_v8));
            let _1394_v9;
            _1394_v9 = _module.D5.create_DC13(_1390_v5, true, (_this).f30);
            let _rhs212 = _1390_v5;
            let _rhs213 = _module.D5.create_DC13(_1390_v5, (_this).f24, !_dafny.Seq.contains(_1391_v6, (_this).f31));
            let _lhs164 = globalState;
            _lhs164.f7 = _rhs212;
            _1394_v9 = _rhs213;
            let _1395_v10;
            let _nw211 = Array((new BigNumber(24)).toNumber()).fill(false);
            _1395_v10 = _nw211;
            let _1396_v11;
            _1396_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1395_v10,_this.f23);
            let _1397_v12;
            let _nw212 = new _module.C3();
            _nw212.__ctor((_1396_v11).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1395_v10,_this.f23)), (_this).f25, _this.f23, true);
            _1397_v12 = _nw212;
          }
        }
      }
      if ((_this).f31) {
        (globalState).f13 = !((_this).f31);
        let _1398_v13;
        let _nw213 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _1398_v13 = _nw213;
        let _index196 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length));
        (_1398_v13)[_index196] = new BigNumber(-212);
        let _1399_v14;
        _1399_v14 = new BigNumber(17);
        let _index197 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length));
        (_1398_v13)[_index197] = _1399_v14;
        let _1400_v15;
        let _nw214 = new _module.C1();
        _nw214.__ctor((_this).f27, (_this).f30);
        _1400_v15 = _nw214;
        let _1401_v16;
        _1401_v16 = _1400_v15;
        let _1402_v17;
        _1402_v17 = _dafny.Seq.of(_1401_v16);
        let _1403_v18;
        _1403_v18 = (((_this).f27) ? (_dafny.Seq.of(_1401_v16)) : (_1402_v17));
        let _source22 = _1403_v18;
        let _1404___mcc_h0 = _source22;
        let _1405_cf34 = _1404___mcc_h0;
        (globalState).f13 = (_this).f31;
        let _1406_v19;
        _1406_v19 = _dafny.Seq.UnicodeFromString("eqrnaakjo");
        let _1407_v20;
        _1407_v20 = _dafny.Set.fromElements(new BigNumber(883), _1399_v14);
        let _1408_v21;
        _1408_v21 = _dafny.Seq.of(_1399_v14, new BigNumber((_1406_v19).length), new BigNumber((_1407_v20).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(_1399_v14)));
        let _1409_v22;
        _1409_v22 = _dafny.Seq.of(_1400_v15);
        let _1410_v23;
        _1410_v23 = _dafny.MultiSet.fromElements(_1400_v15, _1400_v15, _1400_v15);
        let _1411_v24;
        _1411_v24 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1400_v15, _1400_v15), _dafny.MultiSet.FromArray(_1409_v22), _1410_v23, _1410_v23);
        let _1412_v25;
        _1412_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1399_v14,(_1398_v13)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length))]);
        let _rhs214 = (_this).fm9(_1408_v21, (_1398_v13)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length))], (_1400_v15).f39, globalState);
        let _rhs215 = (new BigNumber((_1411_v24).length)).minus(_1399_v14);
        let _rhs216 = ((_this).f31) || (!(_1412_v25).contains((_dafny.ZERO).minus((_1398_v13)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length))])));
        let _lhs165 = globalState;
        _1399_v14 = _rhs214;
        _1399_v14 = _rhs215;
        _lhs165.f13 = _rhs216;
        let _1413_v26;
        _1413_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_1398_v13)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length))]);
        (globalState).f7 = (new BigNumber((_1413_v26).length)).multipliedBy(_module.__default.safeDivisionInt(_1399_v14, _1399_v14));
        (globalState).f7 = (_1398_v13)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1398_v13).length))];
        (globalState).f7 = _1399_v14;
        let _1414_v27;
        _1414_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(_1399_v14)).multipliedBy(_1399_v14),p0);
        _1414_v27 = (_1414_v27).update(_1399_v14, (_1400_v15).f39);
      } else {
        let _1415_v28;
        _1415_v28 = new BigNumber(782);
        let _1416_v29;
        _1416_v29 = _dafny.Seq.UnicodeFromString("byo");
        if ((new BigNumber((_1416_v29).length)).isLessThan(_1415_v28)) {
          (globalState).f13 = !(_dafny.Seq.IsPrefixOf(_1416_v29, _1416_v29));
          let _1417_v30;
          _1417_v30 = _dafny.Seq.of(_module.__default.fm2(globalState), (((_this).f30) ? (p0) : (p0)), (_this).f24);
          let _1418_v31;
          _1418_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1415_v28,_1415_v28);
          let _rhs217 = (_dafny.ZERO).minus(((((_this).f30) ? (new BigNumber(386)) : (_1415_v28))).multipliedBy(_module.__default.fm0(!((_this).f27), (((_1384_v0).contains(true)) ? ((_1384_v0).get(true)) : (_1415_v28)), globalState)));
          let _rhs218 = _1417_v30;
          let _rhs219 = (_dafny.ZERO).minus((((_1418_v31).contains(_1415_v28)) ? ((_1418_v31).get(_1415_v28)) : ((_dafny.ZERO).minus(_1415_v28))));
          let _lhs166 = globalState;
          let _lhs167 = globalState;
          _lhs166.f7 = _rhs217;
          _1417_v30 = _rhs218;
          _lhs167.f7 = _rhs219;
          let _1419_v32;
          _1419_v32 = _dafny.Set.fromElements((_this).f24);
          let _1420_v33;
          _1420_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1419_v32).length));
          let _1421_v34;
          _1421_v34 = _dafny.Set.fromElements(_1420_v33, _1420_v33);
          let _1422_v35;
          _1422_v35 = _module.D0.create_DC3(_module.__default.fm28(_1416_v29, _dafny.Seq.of(_1415_v28), _module.__default.fm2(globalState), _1421_v34, globalState));
          let _1423_v36;
          _1423_v36 = _dafny.Map.Empty.slice().updateUnsafe((_1419_v32).IsDisjointFrom(_1419_v32),_1422_v35);
          _1423_v36 = (_1423_v36).update((_this).f31, _1422_v35);
          let _1424_v37;
          let _init38 = ((_1425_v29) => function (_1426_i1) {
            return (_1426_i1).minus((_dafny.ZERO).minus(new BigNumber((_1425_v29).length)));
          })(_1416_v29);
          let _nw215 = Array((new BigNumber(20)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw215.length); _i0_38++) {
            _nw215[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1424_v37 = _nw215;
          let _1427_v38;
          _1427_v38 = _dafny.Set.fromElements(_1415_v28, _1415_v28, _1415_v28, _1415_v28);
          let _1428_v39;
          _1428_v39 = _dafny.MultiSet.fromElements(new BigNumber((_1417_v30).length));
          let _index198 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1424_v37).length));
          (_1424_v37)[_index198] = new BigNumber(((_1427_v38).Union(function () {
            let _coll83 = new _dafny.Set();
            for (const _compr_83 of (_1428_v39).Elements) {
              let _1429_v40 = _compr_83;
              if ((_1428_v39).contains(_1429_v40)) {
                _coll83.add(_module.__default.safeModuloInt(_1429_v40, new BigNumber(321)));
              }
            }
            return _coll83;
          }())).length);
          let _index199 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1424_v37).length));
          (_1424_v37)[_index199] = _1415_v28;
          let _1430_v41;
          let _nw216 = new _module.C1();
          _nw216.__ctor(true, (_this).f30);
          _1430_v41 = _nw216;
          let _1431_v42;
          _1431_v42 = _1430_v41;
          _1431_v42 = _1431_v42;
        } else {
          let _1432_v43;
          let _nw217 = new _module.C2();
          _nw217.__ctor(_1384_v0, (_this).f25, _this.f23, false);
          _1432_v43 = _nw217;
          let _1433_v44;
          _1433_v44 = _dafny.MultiSet.fromElements(_1415_v28);
          let _1434_v45;
          _1434_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1416_v29)[_module.__default.safeIndex(_1415_v28, new BigNumber((_1416_v29).length))],(_1433_v44).IsProperSubsetOf(_1433_v44));
          let _1435_v46;
          let _init39 = function (_1436_i2) {
            return (_this).f27;
          };
          let _nw218 = Array((new BigNumber(18)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw218.length); _i0_39++) {
            _nw218[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1435_v46 = _nw218;
          let _1437_v47;
          _1437_v47 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f24) ? (_1435_v46) : (_1435_v46)),_1434_v45);
          _1434_v45 = (((_1437_v47).contains(_1435_v46)) ? ((_1437_v47).get(_1435_v46)) : (((_1434_v45).update(_this.f23, p0)).Merge(_1434_v45)));
          (globalState).f13 = (_this).f30;
          _1415_v28 = _1415_v28;
          let _1438_v48;
          _1438_v48 = _dafny.Seq.of((_this).f31, (_this).f24, (_this).f31, (_this).f30, (_this).f31);
          let _1439_v49;
          _1439_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1415_v28,_1438_v48);
          let _1440_v50;
          _1440_v50 = _module.D6.create_DC15(_1439_v49);
          let _1441_v51;
          _1441_v51 = _dafny.Seq.of(_1440_v50);
          let _1442_v52;
          _1442_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_1441_v51, _1441_v51),_1415_v28);
          (globalState).f7 = new BigNumber((_1442_v52).length);
        }
        let _1443_v53;
        let _nw219 = Array((new BigNumber(25)).toNumber()).fill(false);
        _1443_v53 = _nw219;
        let _1444_v54;
        _1444_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v53,_module.__default.fm34(new BigNumber(123), (_this).f24, globalState));
        let _1445_v55;
        let _nw220 = new _module.C3();
        _nw220.__ctor(_1444_v54, (_this).f25, _this.f23, (_this).f31);
        _1445_v55 = _nw220;
        let _1446_v56;
        _1446_v56 = _dafny.MultiSet.fromElements(_this.f23);
        let _1447_v57;
        _1447_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1446_v56).cardinality()),_1445_v55);
        let _1448_v58;
        _1448_v58 = _dafny.Seq.of(_1445_v55, _1445_v55);
        let _1449_v59;
        let _nw221 = Array((new BigNumber(27)).toNumber());
        _nw221[(_dafny.ZERO).toNumber()] = _1445_v55;
        _nw221[(_dafny.ONE).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(2)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(3)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(4)).toNumber()] = (((_1447_v57).contains(_1415_v28)) ? ((_1447_v57).get(_1415_v28)) : (_1445_v55));
        _nw221[(new BigNumber(5)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(6)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(7)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(8)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(9)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(10)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(11)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(12)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(13)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(14)).toNumber()] = (_1448_v58)[_module.__default.safeIndex(_1415_v28, new BigNumber((_1448_v58).length))];
        _nw221[(new BigNumber(15)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(16)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(17)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(18)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(19)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(20)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(21)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(22)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(23)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(24)).toNumber()] = (_1448_v58)[_module.__default.safeIndex(_1415_v28, new BigNumber((_1448_v58).length))];
        _nw221[(new BigNumber(25)).toNumber()] = _1445_v55;
        _nw221[(new BigNumber(26)).toNumber()] = _1445_v55;
        _1449_v59 = _nw221;
        let _index200 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1449_v59).length));
        (_1449_v59)[_index200] = _1445_v55;
        let _index201 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1449_v59).length));
        (_1449_v59)[_index201] = _1445_v55;
        let _1450_v60;
        _1450_v60 = _module.D2.create_DC7(p0, p0);
        let _1451_v61;
        _1451_v61 = _dafny.Set.fromElements(_1450_v60, _1450_v60, _1450_v60, _1450_v60);
        let _pat_let_tv29 = p0;
        _1451_v61 = _dafny.Set.fromElements(function (_pat_let45_0) {
          return function (_1452_dt__update__tmp_h0) {
            return function (_pat_let46_0) {
              return function (_1453_dt__update_hcf14_h0) {
                return _module.D2.create_DC7((_1452_dt__update__tmp_h0).dtor_cf13, _1453_dt__update_hcf14_h0);
              }(_pat_let46_0);
            }(_pat_let_tv29);
          }(_pat_let45_0);
        }(_1450_v60));
        let _1454_v63;
        _1454_v63 = _dafny.Set.fromElements(_1416_v29);
        if (!((function () {
          let _coll84 = new _dafny.Map();
          for (const _compr_84 of (_1454_v63).Elements) {
            let _1455_v62 = _compr_84;
            if ((_1454_v63).contains(_1455_v62)) {
              _coll84.push([_1455_v62,(_module.D9.create_DC20(_1416_v29, p0, _1415_v28)).dtor_cf31]);
            }
          }
          return _coll84;
        }()).update(_1416_v29, _1416_v29)).contains((_module.D4.create_DC10(_module.__default.fm1(p0, globalState))).dtor_cf17)) {
          let _1456_v64;
          let _nw222 = new _module.C3();
          _nw222.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1443_v53,_this.f23), (((_this).f30) ? ((_this).f25) : ((_this).f25)), new _dafny.CodePoint('e'.codePointAt(0)), (_this).f31);
          _1456_v64 = _nw222;
          let _1457_v65;
          _1457_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v53,_module.__default.fm0((_this).f30, (_1386_v2).fm24((_dafny.ZERO).minus(_1415_v28), true, globalState), globalState));
          let _rhs220 = ((p0) ? (_this.f23) : (_this.f23));
          let _rhs221 = (((_1415_v28).isLessThan(_1415_v28)) ? (_1415_v28) : (new BigNumber((_1457_v65).length)));
          let _lhs168 = _this;
          let _lhs169 = globalState;
          _lhs168.f23 = _rhs220;
          _lhs169.f7 = _rhs221;
          _1415_v28 = new BigNumber((_1416_v29).length);
          let _1458_v66;
          _1458_v66 = _module.D12.create_DC25((_this).f31, _this.f23);
          _1458_v66 = _1458_v66;
          let _1459_v67;
          let _1460_v68;
          let _1461_v69;
          let _out40;
          let _out41;
          let _out42;
          let _outcollector12 = (_1445_v55).m3((_this).f27, _this.f23, globalState);
          _out40 = _outcollector12[0];
          _out41 = _outcollector12[1];
          _out42 = _outcollector12[2];
          _1459_v67 = _out40;
          _1460_v68 = _out41;
          _1461_v69 = _out42;
        } else {
          (globalState).f7 = _1415_v28;
          let _1462_v70;
          _1462_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new _dafny.CodePoint('q'.codePointAt(0)));
          let _1463_v71;
          _1463_v71 = _module.D13.create_DC27(_1462_v70);
          let _1464_v72;
          _1464_v72 = _module.D5.create_DC14((_this).f27);
          let _1465_v73;
          _1465_v73 = _module.D11.create_DC23(p0, (_dafny.ZERO).minus(_1415_v28), _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p0), _1464_v72);
          let _1466_v74;
          _1466_v74 = _dafny.Seq.of((_1465_v73).dtor_cf36);
          let _1467_v75;
          let _nw223 = Array((new BigNumber(8)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = (_1462_v70).update((_this).f24, _this.f23);
          _nw223[(_dafny.ONE).toNumber()] = (_1463_v71).dtor_cf44;
          _nw223[(new BigNumber(2)).toNumber()] = (_1462_v70).update((_this).f27, _this.f23);
          _nw223[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1466_v74)[_module.__default.safeIndex(_1415_v28, new BigNumber((_1466_v74).length))],_this.f23);
          _nw223[(new BigNumber(4)).toNumber()] = _1462_v70;
          _nw223[(new BigNumber(5)).toNumber()] = _1462_v70;
          _nw223[(new BigNumber(6)).toNumber()] = _1462_v70;
          _nw223[(new BigNumber(7)).toNumber()] = _1462_v70;
          _1467_v75 = _nw223;
          let _1468_v76;
          let _nw224 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1468_v76 = _nw224;
          let _1469_v77;
          _1469_v77 = _dafny.Seq.of(_1467_v75, _1467_v75, _1467_v75, _1467_v75, _1467_v75);
          let _1470_v78;
          _1470_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1468_v76,(_1469_v77)[_module.__default.safeIndex(_1415_v28, new BigNumber((_1469_v77).length))]);
          let _1471_v79;
          _1471_v79 = _module.D14.create_DC30(_1467_v75);
          let _1472_v80;
          _1472_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1467_v75);
          let _1473_v81;
          let _nw225 = Array((new BigNumber(25)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = _1467_v75;
          _nw225[(_dafny.ONE).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(2)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(3)).toNumber()] = (((_1470_v78).contains(_1468_v76)) ? ((_1470_v78).get(_1468_v76)) : (_1467_v75));
          _nw225[(new BigNumber(4)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(5)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(6)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(7)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(8)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(9)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(10)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(11)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(12)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(13)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(14)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(15)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(16)).toNumber()] = (_1471_v79).dtor_cf51;
          _nw225[(new BigNumber(17)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(18)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(19)).toNumber()] = (((_1472_v80).contains((_this).f24)) ? ((_1472_v80).get((_this).f24)) : (_1467_v75));
          _nw225[(new BigNumber(20)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(21)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(22)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(23)).toNumber()] = _1467_v75;
          _nw225[(new BigNumber(24)).toNumber()] = _1467_v75;
          _1473_v81 = _nw225;
          let _index202 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_1473_v81).length));
          (_1473_v81)[_index202] = _1467_v75;
          let _index203 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_1473_v81).length));
          (_1473_v81)[_index203] = _1467_v75;
          (_this).f23 = _this.f23;
          let _1474_v82;
          _1474_v82 = _dafny.Seq.of(_1415_v28);
          let _1475_v83;
          _1475_v83 = _dafny.MultiSet.fromElements(_1415_v28, _1415_v28, _1415_v28);
          (globalState).f13 = (_1475_v83).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1474_v82));
          let _index204 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1443_v53).length));
          (_1443_v53)[_index204] = ((_this).f24) && ((_this).f27);
          let _index205 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1443_v53).length));
          (_1443_v53)[_index205] = p0;
        }
        let _1476_v84;
        let _nw226 = new _module.C5();
        _nw226.__ctor(_1443_v53);
        _1476_v84 = _nw226;
      }
      let _1477_v85;
      let _init40 = ((_1478_p0) => function (_1479_i3) {
        return _1478_p0;
      })(p0);
      let _nw227 = Array((new BigNumber(2)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw227.length); _i0_40++) {
        _nw227[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1477_v85 = _nw227;
      let _1480_v86;
      _1480_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1477_v85,new BigNumber(550));
      let _1481_v87;
      _1481_v87 = new BigNumber(629);
      let _1482_v88;
      _1482_v88 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1481_v87), _1481_v87, _1481_v87, _1481_v87);
      let _1483_v89;
      _1483_v89 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f31),(_1386_v2).fm4(_1481_v87, _dafny.Seq.of((_this).f24, (_this).f31, (_this).f30), globalState));
      let _rhs222 = _1480_v86;
      let _rhs223 = ((_1482_v88).Difference(_module.__default.fm18(globalState))).Difference(_1482_v88);
      let _rhs224 = (_1481_v87).plus(new BigNumber((_1483_v89).length));
      let _lhs170 = globalState;
      _1480_v86 = _rhs222;
      _1482_v88 = _rhs223;
      _lhs170.f7 = _rhs224;
      let _1484_v90;
      let _nw228 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _1484_v90 = _nw228;
      let _index206 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1484_v90).length));
      (_1484_v90)[_index206] = _1481_v87;
      let _1485_v91;
      _1485_v91 = _dafny.Seq.UnicodeFromString("oaniuldr");
      let _1486_v92;
      _1486_v92 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_dafny.Seq.update(_1485_v91, _module.__default.safeIndex(_1481_v87, new BigNumber((_1485_v91).length)), _this.f23));
      let _index207 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1484_v90).length));
      let _rhs225 = (_1486_v92).equals((_1486_v92).Merge(_1486_v92));
      let _rhs226 = _1481_v87;
      let _lhs171 = globalState;
      let _lhs172 = _1484_v90;
      let _lhs173 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1484_v90).length));
      _lhs171.f13 = _rhs225;
      _lhs172[_lhs173] = _rhs226;
      r0 = ((_this).f27) && ((_this).f24);
      return r0;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let _1487_v0;
      let _1488_v1;
      let _1489_v2;
      let _out43;
      let _out44;
      let _out45;
      let _outcollector13 = _module.__default.m0((p0).plus(p0), _module.__default.fm3((_this).f30, globalState), true, (p0).plus(p0), globalState);
      _out43 = _outcollector13[0];
      _out44 = _outcollector13[1];
      _out45 = _outcollector13[2];
      _1487_v0 = _out43;
      _1488_v1 = _out44;
      _1489_v2 = _out45;
      (globalState).f7 = (p0).minus(p0);
      let _1490_v3;
      _1490_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p0);
      let _1491_v4;
      _1491_v4 = _dafny.Seq.UnicodeFromString("dfrewfj");
      let _1492_v5;
      _1492_v5 = _dafny.Set.fromElements(_1489_v2);
      let _1493_v6;
      _1493_v6 = _module.D5.create_DC12(_1492_v5);
      let _pat_let_tv30 = p1;
      let _pat_let_tv31 = _1487_v0;
      let _rhs227 = _1490_v3;
      let _rhs228 = p0;
      let _rhs229 = (_this).f24;
      let _rhs230 = !(function (_source23) {
        if (_source23.is_DC13) {
          let _1494___mcc_h0 = (_source23).cf20;
          let _1495___mcc_h1 = (_source23).cf21;
          let _1496___mcc_h2 = (_source23).cf22;
          let _1497_cf22 = _1496___mcc_h2;
          let _1498_cf21 = _1495___mcc_h1;
          let _1499_cf20 = _1494___mcc_h0;
          return !(_pat_let_tv30) || (_1498_cf21);
        } else if (_source23.is_DC14) {
          let _1500___mcc_h3 = (_source23).cf23;
          let _1501_cf23 = _1500___mcc_h3;
          return _1501_cf23;
        } else {
          let _1502___mcc_h4 = (_source23).cf19;
          let _1503_cf19 = _1502___mcc_h4;
          return (new BigNumber((_dafny.Seq.UnicodeFromString("wpj")).length)).isLessThan(_pat_let_tv31);
        }
      }(_1493_v6));
      let _rhs231 = ((_1488_v1) ? (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), function (_1504_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _module.__default.safeIndex(_1489_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), function (_1505_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length)), _this.f23)) : (_dafny.Seq.Concat(_1491_v4, _1491_v4)));
      _1490_v3 = _rhs227;
      _1489_v2 = _rhs228;
      _1488_v1 = _rhs229;
      _1488_v1 = _rhs230;
      _1491_v4 = _rhs231;
      let _hi8 = (p0).plus(_1489_v2);
      for (let _1506_i1 = p0; _1506_i1.isLessThan(_hi8); _1506_i1 = _1506_i1.plus(_dafny.ONE)) {
        let _1507_v7;
        _1507_v7 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (_1487_v0) : (p0)),!((_this).f27));
        let _1508_v8;
        _1508_v8 = _dafny.Seq.of(p1, (_this).f27);
        let _1509_v9;
        _1509_v9 = _dafny.Seq.of(false, (_this).fm4(p0, _dafny.Seq.of((_this).f30), globalState), (_1508_v8)[_module.__default.safeIndex(_1489_v2, new BigNumber((_1508_v8).length))], (_this).f30, (_this).f27);
        _1507_v7 = (_1507_v7).update(_module.__default.safeDivisionInt(_1489_v2, _1489_v2), (_1487_v0).isLessThanOrEqualTo(new BigNumber((_1509_v9).length)));
        _1488_v1 = p1;
        (globalState).f13 = p1;
        (globalState).f13 = (_this).fm4(_1487_v0, _1508_v8, globalState);
      }
      let _1510_v10;
      _1510_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(p0, _1487_v0));
      let _1511_v11;
      _1511_v11 = _module.D4.create_DC10(_1491_v4);
      let _1512_v12;
      _1512_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1510_v10,_1511_v11);
      let _1513_v14;
      let _nw229 = new _module.C0();
      _nw229.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this).f31,_1491_v4), (_this).f31, _this.f23, (_this).f31);
      _1513_v14 = _nw229;
      let _1514_v15;
      _1514_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1488_v1,_1513_v14);
      let _1515_v16;
      _1515_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1489_v2,_1514_v15);
      let _1516_v17;
      _1516_v17 = _dafny.Seq.of(_1487_v0, new BigNumber((_1515_v16).length), p0, _1487_v0);
      _1512_v12 = (_1512_v12).update(((!(!((_this).f31))) ? (function () {
        let _coll85 = new _dafny.Map();
        for (const _compr_85 of _dafny.IntegerRange(new BigNumber(-590), new BigNumber(391))) {
          let _1517_v13 = _compr_85;
          if (((new BigNumber(-590)).isLessThanOrEqualTo(_1517_v13)) && ((_1517_v13).isLessThan(new BigNumber(391)))) {
            _coll85.push([(_1517_v13).plus(new BigNumber(202)),_dafny.Seq.of(p0)]);
          }
        }
        return _coll85;
      }()) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1491_v4).length),_1516_v17))), _module.D4.create_DC10(_1491_v4));
      let _hi9 = _module.__default.safeDivisionInt((_module.D0.create_DC2(!(_1488_v1), _1487_v0, _1488_v1)).dtor_cf3, _1489_v2);
      for (let _1518_i2 = new BigNumber(743); _1518_i2.isLessThan(_hi9); _1518_i2 = _1518_i2.plus(_dafny.ONE)) {
        if (!((_1492_v5).IsProperSubsetOf(_1492_v5))) {
          let _rhs232 = (((_this).f27) ? (((false) ? ((_1513_v14).f24) : ((_1513_v14).f24))) : ((_this).f31));
          let _rhs233 = _1491_v4;
          let _lhs174 = globalState;
          _lhs174.f13 = _rhs232;
          _1491_v4 = _rhs233;
          let _1519_v18;
          _1519_v18 = _dafny.Seq.of((_this).f24);
          let _1520_v19;
          _1520_v19 = _module.D14.create_DC31((_1513_v14).f24);
          _1519_v18 = _dafny.Seq.Concat(_1519_v18, _dafny.Seq.of((_1520_v19).dtor_cf52));
          let _1521_v20;
          _1521_v20 = _dafny.Seq.of(_module.__default.fm36(globalState));
          let _1522_v21;
          _1522_v21 = _dafny.MultiSet.fromElements(new BigNumber((_1521_v20).length), _1489_v2, _1518_i2);
          let _1523_v22;
          _1523_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1522_v21).IsDisjointFrom(_1522_v21),_1492_v5);
          _1489_v2 = new BigNumber((_1523_v22).length);
          let _1524_v23;
          _1524_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1489_v2).isLessThanOrEqualTo(_1518_i2),!((_this).f27) || (!(_1488_v1)));
          _1524_v23 = (_1524_v23).update(true, !((_dafny.ZERO).minus(_1487_v0)).isEqualTo(p0));
          _1519_v18 = _1519_v18;
        } else {
          let _1525_v24;
          _1525_v24 = _dafny.MultiSet.fromElements((_1513_v14).f24);
          let _1526_v25;
          _1526_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1518_i2,_1489_v2);
          let _1527_v26;
          let _nw230 = Array((new BigNumber(17)).toNumber());
          _nw230[(_dafny.ZERO).toNumber()] = new BigNumber((_1491_v4).length);
          _nw230[(_dafny.ONE).toNumber()] = p0;
          _nw230[(new BigNumber(2)).toNumber()] = ((_dafny.ZERO).minus(_1518_i2)).minus(new BigNumber(862));
          _nw230[(new BigNumber(3)).toNumber()] = _1489_v2;
          _nw230[(new BigNumber(4)).toNumber()] = (new BigNumber(788)).multipliedBy(new BigNumber((_1525_v24).cardinality()));
          _nw230[(new BigNumber(5)).toNumber()] = (((_1525_v24).contains((_this).f24)) ? ((_1525_v24).get((_this).f24)) : (p0));
          _nw230[(new BigNumber(6)).toNumber()] = new BigNumber(-929);
          _nw230[(new BigNumber(7)).toNumber()] = p0;
          _nw230[(new BigNumber(8)).toNumber()] = _1518_i2;
          _nw230[(new BigNumber(9)).toNumber()] = new BigNumber(-59);
          _nw230[(new BigNumber(10)).toNumber()] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_1528_i3) {
            return _this.f23;
          })).length)).plus(_1487_v0);
          _nw230[(new BigNumber(11)).toNumber()] = _1489_v2;
          _nw230[(new BigNumber(12)).toNumber()] = (p0).minus((_dafny.ZERO).minus(_1518_i2));
          _nw230[(new BigNumber(13)).toNumber()] = p0;
          _nw230[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1489_v2));
          _nw230[(new BigNumber(15)).toNumber()] = p0;
          _nw230[(new BigNumber(16)).toNumber()] = (((_1526_v25).contains(_1489_v2)) ? ((_1526_v25).get(_1489_v2)) : (_1487_v0));
          _1527_v26 = _nw230;
          let _index208 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1527_v26).length));
          (_1527_v26)[_index208] = _1487_v0;
          let _index209 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1527_v26).length));
          (_1527_v26)[_index209] = (((_1525_v24).contains(!((_1513_v14).f24))) ? ((_1525_v24).get(!((_1513_v14).f24))) : ((_dafny.ZERO).minus((_1516_v17)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_1516_v17).length))])));
          let _1529_v27;
          _1529_v27 = _dafny.Seq.of(_dafny.Seq.Concat(_1491_v4, _1491_v4), _dafny.Seq.UnicodeFromString("ifjf"), _1491_v4);
          _1529_v27 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1529_v27, _1529_v27), _dafny.Seq.Concat(_1529_v27, _1529_v27));
          let _1530_v28;
          _1530_v28 = _module.D11.create_DC22(_1527_v26);
          (globalState).f22 = (_1530_v28).dtor_cf35;
          _1488_v1 = (_this).f24;
          let _1531_v29;
          _1531_v29 = _dafny.Map.Empty.slice().updateUnsafe(false,_1491_v4);
          let _1532_v30;
          _1532_v30 = _dafny.Seq.of(_1488_v1, (_this).f27, (_this).f27);
          let _1533_v31;
          let _nw231 = new _module.C0();
          _nw231.__ctor(((_1531_v29).update((_this).f30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), ((_1534_v14) => function (_1535_i4) {
            return _1534_v14.f23;
          })(_1513_v14)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_1491_v4)), (new BigNumber((_1532_v30).length)).isLessThanOrEqualTo(_1487_v0), _1513_v14.f23, (_this).f24);
          _1533_v31 = _nw231;
        }
        _1488_v1 = (_1513_v14).f24;
        let _1536_v32;
        _1536_v32 = _module.D12.create_DC25(!(p1), _1513_v14.f23);
        let _1537_v33;
        _1537_v33 = _module.D12.create_DC26(_1536_v32);
        _1537_v33 = _1537_v33;
        (globalState).f13 = (_this).f31;
      }
      return;
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f27 = false;
    }
    _parentTraits() {
      return [_module.T3];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f27) {
      let _this = this;
      (_this)._f27 = f27;
      return;
    }
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(-598);
    };
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      if ((_this).f27) {
        let _1538_v0;
        _1538_v0 = _dafny.MultiSet.fromElements(false);
        let _1539_v1;
        _1539_v1 = _dafny.Set.fromElements(_module.D0.create_DC0(_1538_v0));
        (globalState).f13 = (_1539_v1).IsDisjointFrom(_1539_v1);
        let _1540_v2;
        _1540_v2 = _dafny.Seq.UnicodeFromString("ktgrmbaw");
        let _1541_v3;
        _1541_v3 = _dafny.Set.fromElements(new BigNumber((_1540_v2).length));
        let _1542_v4;
        _1542_v4 = _dafny.Seq.of(_1541_v3);
        let _1543_v5;
        _1543_v5 = new BigNumber(57);
        (globalState).f13 = !(_1541_v3).equals(((_1542_v4)[_module.__default.safeIndex(_1543_v5, new BigNumber((_1542_v4).length))]).Union(_dafny.Set.fromElements(new BigNumber(639), _1543_v5)));
        let _1544_v6;
        _1544_v6 = new _dafny.CodePoint('y'.codePointAt(0));
        _1544_v6 = _1544_v6;
        let _1545_v7;
        _1545_v7 = _dafny.Seq.of(_1543_v5, _1543_v5, _1543_v5);
        let _1546_v8;
        _1546_v8 = _dafny.Seq.of(_1545_v7);
        let _1547_v9;
        let _nw232 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1547_v9 = _nw232;
        let _1548_v10;
        _1548_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1543_v5,_1547_v9);
        let _1549_v11;
        _1549_v11 = _dafny.Set.fromElements((((_1548_v10).contains(_1543_v5)) ? ((_1548_v10).get(_1543_v5)) : (_1547_v9)), _1547_v9, _1547_v9, _1547_v9);
        let _1550_v12;
        _1550_v12 = _dafny.Seq.of((_this).f27, (_this).f27);
        let _1551_v13;
        _1551_v13 = _dafny.Seq.of(_dafny.Seq.update(_1550_v12, _module.__default.safeIndex(new BigNumber((_1540_v2).length), new BigNumber((_1550_v12).length)), (_this).f27));
        let _rhs234 = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.of(_1543_v5, new BigNumber((_1549_v11).length), _1543_v5), _module.__default.safeIndex(_1543_v5, new BigNumber((_dafny.Seq.of(_1543_v5, new BigNumber((_1549_v11).length), _1543_v5)).length)), new BigNumber((_1540_v2).length)), _1545_v7), _1546_v8);
        let _rhs235 = _1543_v5;
        let _rhs236 = _dafny.Seq.IsProperPrefixOf(((p0) ? (_1551_v13) : (_1551_v13)), _dafny.Seq.of(_dafny.Seq.of(p0, !(!((_this).f27)))));
        let _rhs237 = _1545_v7;
        let _lhs175 = globalState;
        let _lhs176 = globalState;
        _1546_v8 = _rhs234;
        _lhs175.f7 = _rhs235;
        _lhs176.f13 = _rhs236;
        _1545_v7 = _rhs237;
        let _1552_v14;
        let _nw233 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1552_v14 = _nw233;
        let _1553_v15;
        _1553_v15 = _module.D1.create_DC4(_1544_v6);
        let _rhs238 = (_1553_v15).dtor_cf6;
        let _rhs239 = _1552_v14;
        let _rhs240 = (_1543_v5).plus((_1543_v5).plus(_module.__default.fm0(p0, _1543_v5, globalState)));
        let _lhs177 = globalState;
        _1544_v6 = _rhs238;
        _1552_v14 = _rhs239;
        _lhs177.f7 = _rhs240;
      } else {
        let _1554_v16;
        _1554_v16 = new BigNumber(253);
        let _1555_v17;
        _1555_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f27, new BigNumber(570), globalState),!((_module.D0.create_DC2((_this).f27, _1554_v16, (_this).f27)).dtor_cf4));
        if ((((_1555_v17).contains(new BigNumber(166))) ? ((_1555_v17).get(new BigNumber(166))) : (p0))) {
          let _1556_v18;
          let _init41 = ((_1557_v16) => function (_1558_i0) {
            return _module.__default.safeDivisionInt(_1558_i0, _1557_v16);
          })(_1554_v16);
          let _nw234 = Array((new BigNumber(24)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw234.length); _i0_41++) {
            _nw234[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1556_v18 = _nw234;
          let _1559_v19;
          _1559_v19 = _dafny.Set.fromElements(new BigNumber(409), _1554_v16, _1554_v16);
          let _rhs241 = _1556_v18;
          let _rhs242 = _1554_v16;
          let _rhs243 = (_1559_v19).IsDisjointFrom(_1559_v19);
          let _lhs178 = globalState;
          let _lhs179 = globalState;
          let _lhs180 = globalState;
          _lhs178.f22 = _rhs241;
          _lhs179.f7 = _rhs242;
          _lhs180.f13 = _rhs243;
          let _1560_v20;
          let _init42 = ((_1561_p0) => function (_1562_i1) {
            return _1561_p0;
          })(p0);
          let _nw235 = Array((new BigNumber(29)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw235.length); _i0_42++) {
            _nw235[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1560_v20 = _nw235;
          let _1563_v21;
          let _init43 = ((_1564_v16) => function (_1565_i2) {
            return _module.D0.create_DC1(_1564_v16);
          })(_1554_v16);
          let _nw236 = Array((new BigNumber(7)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw236.length); _i0_43++) {
            _nw236[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1563_v21 = _nw236;
          let _1566_v22;
          _1566_v22 = _dafny.Set.fromElements(_1563_v21);
          let _rhs244 = (_dafny.ZERO).minus(new BigNumber((((_1566_v22).Union(_1566_v22)).Union(_1566_v22)).length));
          let _rhs245 = _1560_v20;
          let _rhs246 = (_1554_v16).plus(_1554_v16);
          let _lhs181 = globalState;
          _1554_v16 = _rhs244;
          _1560_v20 = _rhs245;
          _lhs181.f7 = _rhs246;
          let _1567_v23;
          _1567_v23 = new _dafny.CodePoint('o'.codePointAt(0));
          let _1568_v24;
          let _nw237 = Array((new BigNumber(7)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = _1567_v23;
          _nw237[(_dafny.ONE).toNumber()] = _1567_v23;
          _nw237[(new BigNumber(2)).toNumber()] = _1567_v23;
          _nw237[(new BigNumber(3)).toNumber()] = _1567_v23;
          _nw237[(new BigNumber(4)).toNumber()] = _1567_v23;
          _nw237[(new BigNumber(5)).toNumber()] = _1567_v23;
          _nw237[(new BigNumber(6)).toNumber()] = _1567_v23;
          _1568_v24 = _nw237;
          let _index210 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_1568_v24).length));
          (_1568_v24)[_index210] = _1567_v23;
          let _index211 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_1568_v24).length));
          (_1568_v24)[_index211] = _1567_v23;
          (globalState).f7 = _module.__default.fm0((_this).f27, _1554_v16, globalState);
          let _1569_v25;
          _1569_v25 = _module.D0.create_DC2(p0, _1554_v16, p0);
          let _1570_v26;
          _1570_v26 = _dafny.Set.fromElements((((_this).f27) ? (p0) : (!(p0))), (_1569_v25).dtor_cf2);
          _1570_v26 = (_1570_v26).Difference((_1570_v26).Union(_1570_v26));
        } else {
          let _1571_v27;
          _1571_v27 = _dafny.Seq.of(_1554_v16);
          let _1572_v28;
          _1572_v28 = _dafny.Seq.UnicodeFromString("lwpgw");
          let _1573_v29;
          let _nw238 = Array((new BigNumber(10)).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = _1554_v16;
          _nw238[(_dafny.ONE).toNumber()] = _1554_v16;
          _nw238[(new BigNumber(2)).toNumber()] = (((_this).f27) ? (new BigNumber(137)) : (_1554_v16));
          _nw238[(new BigNumber(3)).toNumber()] = _1554_v16;
          _nw238[(new BigNumber(4)).toNumber()] = _1554_v16;
          _nw238[(new BigNumber(5)).toNumber()] = _1554_v16;
          _nw238[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1571_v27).length)), _1554_v16);
          _nw238[(new BigNumber(7)).toNumber()] = _1554_v16;
          _nw238[(new BigNumber(8)).toNumber()] = _module.__default.fm0(p0, _1554_v16, globalState);
          _nw238[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1572_v28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(2)), function (_1574_i3) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }))).length);
          _1573_v29 = _nw238;
          (globalState).f22 = _1573_v29;
          let _1575_v30;
          _1575_v30 = _dafny.Seq.of((_this).f27);
          let _1576_v31;
          _1576_v31 = _module.D0.create_DC1(_1554_v16);
          let _1577_v32;
          _1577_v32 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_1575_v30, p0),_1576_v31);
          let _1578_v33;
          _1578_v33 = _dafny.Seq.of(_1571_v27, _dafny.Seq.of(_1554_v16));
          let _1579_v34;
          _1579_v34 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1580_v35;
          let _nw239 = Array((new BigNumber(19)).toNumber());
          _nw239[(_dafny.ZERO).toNumber()] = _1579_v34;
          _nw239[(_dafny.ONE).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(2)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(3)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(4)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(5)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(6)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(7)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(8)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(9)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(10)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(11)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(12)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(13)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(14)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(15)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(16)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(17)).toNumber()] = _1579_v34;
          _nw239[(new BigNumber(18)).toNumber()] = _1579_v34;
          _1580_v35 = _nw239;
          let _1581_v36;
          _1581_v36 = _module.D12.create_DC24(_1580_v35);
          let _1582_v37;
          let _nw240 = new _module.C7();
          _nw240.__ctor((_this).f27, !(true), (_1581_v36).dtor_cf40, _1579_v34, (_this).f27, p0);
          _1582_v37 = _nw240;
          let _1583_v38;
          _1583_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1582_v37,(_this).f27);
          let _rhs247 = _module.__default.fm13(_module.__default.fm14((_this).f27, globalState), !(_dafny.Seq.IsProperPrefixOf(_1578_v33, _dafny.Seq.update(_1578_v33, _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1578_v33).length)), _dafny.Seq.of(_1554_v16)))), _1554_v16, (((_1583_v38).contains(_1582_v37)) ? ((_1583_v38).get(_1582_v37)) : ((_1582_v37).f27)), globalState);
          let _rhs248 = _module.__default.safeDivisionInt((_1571_v27)[_module.__default.safeIndex(_1554_v16, new BigNumber((_1571_v27).length))], _1554_v16);
          let _lhs182 = globalState;
          _1577_v32 = _rhs247;
          _lhs182.f7 = _rhs248;
          _1572_v28 = _dafny.Seq.Concat(_1572_v28, _1572_v28);
          let _nw241 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1573_v29 = _nw241;
          let _1584_v39;
          let _nw242 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
          _1584_v39 = _nw242;
          let _1585_v40;
          _1585_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1575_v30).length),_module.D14.create_DC30(_1584_v39));
          let _1586_v41;
          _1586_v41 = _dafny.Seq.of(_1585_v40, _1585_v40, _1585_v40);
          (globalState).f7 = (_this).fm9(_1571_v27, new BigNumber((_1586_v41).length), !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("g"), new _dafny.CodePoint('f'.codePointAt(0))), globalState);
        }
        let _1587_v42;
        _1587_v42 = _dafny.Set.fromElements(new BigNumber(-926));
        let _1588_v43;
        _1588_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1587_v42,_1587_v42);
        let _1589_v44;
        _1589_v44 = _dafny.Map.Empty.slice().updateUnsafe((((_1588_v43).contains(_1587_v42)) ? ((_1588_v43).get(_1587_v42)) : (_1587_v42)),_1554_v16);
        let _1590_v45;
        _1590_v45 = _dafny.MultiSet.fromElements(_1554_v16);
        _1554_v16 = _module.__default.safeModuloInt(new BigNumber(((_1589_v44).update(_1587_v42, new BigNumber(101))).length), ((p0) ? (new BigNumber((_1590_v45).cardinality())) : (_1554_v16)));
        (globalState).f7 = (new BigNumber(363)).minus(_module.__default.safeDivisionInt(_1554_v16, _1554_v16));
        let _1591_v46;
        _1591_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
        let _source24 = _module.D11.create_DC23((_this).f27, (_1554_v16).multipliedBy(_1554_v16), _1591_v46, _module.D5.create_DC14(p0));
        if (_source24.is_DC23) {
          let _1592___mcc_h0 = (_source24).cf36;
          let _1593___mcc_h1 = (_source24).cf37;
          let _1594___mcc_h2 = (_source24).cf38;
          let _1595___mcc_h3 = (_source24).cf39;
          let _1596_cf39 = _1595___mcc_h3;
          let _1597_cf38 = _1594___mcc_h2;
          let _1598_cf37 = _1593___mcc_h1;
          let _1599_cf36 = _1592___mcc_h0;
          _1598_cf37 = _1554_v16;
          let _1600_v47;
          let _init44 = ((_1601_cf37) => function (_1602_i4) {
            return (_1602_i4).plus(_1601_cf37);
          })(_1598_cf37);
          let _nw243 = Array((new BigNumber(23)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw243.length); _i0_44++) {
            _nw243[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1600_v47 = _nw243;
          let _1603_v48;
          _1603_v48 = _dafny.MultiSet.fromElements(_1600_v47);
          let _1604_v49;
          _1604_v49 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1605_v50;
          _1605_v50 = _dafny.Seq.UnicodeFromString("o");
          let _1606_v51;
          _1606_v51 = _dafny.Seq.of(true, (_this).f27);
          let _1607_v52;
          _1607_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1598_cf37,_1606_v51);
          let _1608_v53;
          _1608_v53 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_1605_v50, _1604_v49),new BigNumber((_1607_v52).length));
          let _rhs249 = _1603_v48;
          let _rhs250 = _1608_v53;
          _1603_v48 = _rhs249;
          _1608_v53 = _rhs250;
          let _1609_v54;
          let _init45 = ((_1610_p0) => function (_1611_i5) {
            return _1610_p0;
          })(p0);
          let _nw244 = Array((new BigNumber(15)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw244.length); _i0_45++) {
            _nw244[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1609_v54 = _nw244;
          let _1612_v55;
          _1612_v55 = _dafny.Set.fromElements(_1609_v54);
          (globalState).f10 = _1612_v55;
          let _index212 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1600_v47).length));
          (_1600_v47)[_index212] = (new BigNumber(59)).multipliedBy(_module.__default.fm0(p0, (_dafny.ZERO).minus(_1598_cf37), globalState));
          let _index213 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1600_v47).length));
          (_1600_v47)[_index213] = _1554_v16;
          let _1613_v56;
          _1613_v56 = _dafny.Seq.of(_1554_v16, _1554_v16, _1554_v16);
          let _1614_v57;
          _1614_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1598_cf37,_1598_cf37);
          let _1615_v58;
          _1615_v58 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_1613_v56).length), new BigNumber((_1606_v51).length)),_1614_v57);
          let _1616_v59;
          _1616_v59 = _module.D9.create_DC20(_1605_v50, (_this).f27, _1598_cf37);
          let _1617_v60;
          _1617_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1616_v59,_module.__default.fm0(_1599_cf36, _1598_cf37, globalState));
          let _1618_v61;
          _1618_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1617_v60,_1605_v50);
          let _1619_v63;
          _1619_v63 = _dafny.MultiSet.fromElements(_module.D9.create_DC20(_1605_v50, p0, _1554_v16));
          let _1620_v65;
          _1620_v65 = _module.D15.create_DC33(function () {
  let _coll86 = new _dafny.Map();
  for (const _compr_86 of (function () {
    let _coll87 = new _dafny.Set();
    for (const _compr_87 of (_1619_v63).Elements) {
      let _1621_v64 = _compr_87;
      if ((_1619_v63).contains(_1621_v64)) {
        _coll87.add(_1621_v64);
      }
    }
    return _coll87;
  }()).Elements) {
    let _1622_v62 = _compr_86;
    if ((function () {
      let _coll88 = new _dafny.Set();
      for (const _compr_88 of (_1619_v63).Elements) {
        let _1623_v64 = _compr_88;
        if ((_1619_v63).contains(_1623_v64)) {
          _coll88.add(_1623_v64);
        }
      }
      return _coll88;
    }()).contains(_1622_v62)) {
      _coll86.push([_1622_v62,_1554_v16]);
    }
  }
  return _coll86;
}());
          let _index214 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1600_v47).length));
          let _index215 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1600_v47).length));
          let _rhs251 = new BigNumber((_1615_v58).length);
          let _rhs252 = (_1598_cf37).minus((_dafny.ZERO).minus(_module.__default.fm0((_this).f27, _1554_v16, globalState)));
          let _rhs253 = (((_1618_v61).contains((_1617_v60).Merge((_1620_v65).dtor_cf54))) ? ((_1618_v61).get((_1617_v60).Merge((_1620_v65).dtor_cf54))) : (_1605_v50));
          let _rhs254 = ((!(_1599_cf36)) ? (_dafny.Seq.update(_module.__default.fm36(globalState), _module.__default.safeIndex(_1554_v16, new BigNumber((_module.__default.fm36(globalState)).length)), (_this).f27)) : (_dafny.Seq.Concat(_1606_v51, _1606_v51)));
          let _lhs183 = _1600_v47;
          let _lhs184 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1600_v47).length));
          let _lhs185 = _1600_v47;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1600_v47).length));
          _lhs183[_lhs184] = _rhs251;
          _lhs185[_lhs186] = _rhs252;
          _1605_v50 = _rhs253;
          _1606_v51 = _rhs254;
        } else {
          let _1624___mcc_h4 = (_source24).cf35;
          let _1625_cf35 = _1624___mcc_h4;
          let _1626_v66;
          let _nw245 = Array((new BigNumber(7)).toNumber()).fill([]);
          _1626_v66 = _nw245;
          let _1627_v67;
          let _nw246 = Array((new BigNumber(16)).toNumber()).fill(false);
          _1627_v67 = _nw246;
          let _index216 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_1626_v66).length));
          (_1626_v66)[_index216] = _1627_v67;
          let _1628_v68;
          let _nw247 = Array((new BigNumber(3)).toNumber());
          _1628_v68 = _nw247;
          let _1629_v69;
          _1629_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1628_v68,_1627_v67);
          let _index217 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_1626_v66).length));
          (_1626_v66)[_index217] = (((_1629_v69).contains(_1628_v68)) ? ((_1629_v69).get(_1628_v68)) : (_1627_v67));
          let _1630_v70;
          _1630_v70 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1631_v71;
          _1631_v71 = _dafny.Map.Empty.slice().updateUnsafe((_1554_v16).minus(_1554_v16),_1630_v70);
          _1631_v71 = (_1631_v71).update(_1554_v16, _1630_v70);
          let _1632_v72;
          let _nw248 = Array((new BigNumber(25)).toNumber()).fill(_module.D12.Default());
          _1632_v72 = _nw248;
          let _1633_v73;
          _1633_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1554_v16,_1632_v72);
          let _1634_v74;
          _1634_v74 = _dafny.Seq.of(_1554_v16, _1554_v16, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), ((_1635_v70) => function (_1636_i6) {
            return _1635_v70;
          })(_1630_v70))).length));
          _1633_v73 = (_1633_v73).update(new BigNumber((_1634_v74).length), _1632_v72);
          let _1637_v75;
          _1637_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1554_v16,_1590_v45);
          _1637_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1554_v16,_1590_v45);
        }
        (globalState).f13 = !(_module.__default.fm2(globalState)) || (!(p0));
      }
      let _1638_v76;
      _1638_v76 = new BigNumber(549);
      let _hi10 = _1638_v76;
      for (let _1639_i7 = _1638_v76; _1639_i7.isLessThan(_hi10); _1639_i7 = _1639_i7.plus(_dafny.ONE)) {
        (globalState).f13 = (p0) === (false);
        let _1640_v77;
        let _nw249 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1640_v77 = _nw249;
        let _1641_v78;
        _1641_v78 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1642_v79;
        _1642_v79 = _dafny.Seq.UnicodeFromString("bjjnna");
        let _1643_v80;
        _1643_v80 = _module.D9.create_DC20(_1642_v79, p0, new BigNumber((_dafny.Set.fromElements(_1639_i7)).length));
        let _1644_v81;
        _1644_v81 = _dafny.Seq.of(_1639_i7, (_1643_v80).dtor_cf33);
        let _1645_v82;
        _1645_v82 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1642_v79).length)),(_this).f27)).length), new BigNumber(720), new BigNumber((_1644_v81).length));
        let _1646_v83;
        let _nw250 = new _module.C7();
        _nw250.__ctor(true, !(p0), _1640_v77, _1641_v78, !(!(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), ((_1647_i7) => function (_1648_i8) {
          return _1647_i7;
        })(_1639_i7)), _1644_v81))), !(_module.__default.fm2(globalState)) || (p0));
        _1646_v83 = _nw250;
        let _1649_v84;
        _1649_v84 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1644_v81);
        let _1650_v85;
        _1650_v85 = _dafny.MultiSet.fromElements(true);
        let _1651_v86;
        let _1652_v87;
        let _1653_v88;
        let _out46;
        let _out47;
        let _out48;
        let _outcollector14 = _module.__default.m0(new BigNumber(((((_1646_v83).f31) ? (_1649_v84) : (_1649_v84))).length), _1650_v85, (_1646_v83).f30, _1639_i7, globalState);
        _out46 = _outcollector14[0];
        _out47 = _outcollector14[1];
        _out48 = _outcollector14[2];
        _1651_v86 = _out46;
        _1652_v87 = _out47;
        _1653_v88 = _out48;
        _1642_v79 = _dafny.Seq.UnicodeFromString("k");
      }
      let _1654_v89;
      _1654_v89 = _dafny.Seq.UnicodeFromString("vgc");
      _1654_v89 = _1654_v89;
      _1638_v76 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1654_v89, _1654_v89), _dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_1655_i9) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }))).length);
      r0 = _module.__default.fm2(globalState);
      _1638_v76 = _module.__default.safeDivisionInt(_1638_v76, (_1638_v76).plus(_1638_v76));
      r0 = (_module.D5.create_DC13((_dafny.ZERO).minus(new BigNumber((_1654_v89).length)), p0, p0)).dtor_cf22;
      return r0;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f23, f24) {
      let _this = this;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return _module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2((_this).f24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f24, !((_this).f24))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC3(_module.D0.create_DC0(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f24)))),(_this).f24)).length))).length), (_this).f24)));
    };
    fm6(globalState) {
      let _this = this;
      return new BigNumber((((_dafny.Set.fromElements(_module.D0.create_DC1(new BigNumber(219)), _module.D0.create_DC1(new BigNumber(367)))).Difference(_dafny.Set.fromElements(_module.D0.create_DC1(new BigNumber((_dafny.Seq.of(_this.f23, new _dafny.CodePoint('u'.codePointAt(0)))).length)), _module.D0.create_DC1(new BigNumber(715)), _module.D0.create_DC1(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(380)),(_this).f24)).length))).cardinality()))))).Intersect((_dafny.Set.fromElements(_module.D0.create_DC1(new BigNumber(890)))).Union(_dafny.Set.fromElements(_module.D0.create_DC1(new BigNumber((_dafny.Seq.UnicodeFromString("obaunck")).length)), _module.D0.create_DC1(new BigNumber(-720)), _module.D0.create_DC1(new BigNumber((_dafny.Seq.of((_this).f24)).length)), _module.D0.create_DC1(new BigNumber((_dafny.Seq.of(new BigNumber(83), new BigNumber(946))).length)))))).length);
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of((_this).f24)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), function (_1656_i0) {
        return _module.D0.create_DC3(_module.D0.create_DC1(new BigNumber(987)));
      })).length))).isLessThan((_dafny.ZERO).minus(new BigNumber(((function () {
        let _coll89 = new _dafny.Map();
        for (const _compr_89 of _dafny.IntegerRange(new BigNumber(456), new BigNumber(762))) {
          let _1657_v0 = _compr_89;
          if (((new BigNumber(456)).isLessThanOrEqualTo(_1657_v0)) && ((_1657_v0).isLessThan(new BigNumber(762)))) {
            _coll89.push([(_1657_v0).plus(new BigNumber(36)),false]);
          }
        }
        return _coll89;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(613),(_this).f24))).length)));
    };
    fm11(p0, p1, globalState) {
      let _this = this;
      return (_this).f24;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D0.Default();
      let r2 = false;
      let _1658_v0;
      _1658_v0 = new BigNumber(135);
      let _1659_v1;
      _1659_v1 = _module.D0.create_DC2(p1, _1658_v0, false);
      if ((_1659_v1).dtor_cf2) {
        let _1660_v2;
        _1660_v2 = _dafny.Seq.of(new BigNumber(357));
        let _1661_v3;
        _1661_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p1, globalState),new BigNumber((_1660_v2).length));
        let _1662_v4;
        _1662_v4 = _dafny.Seq.UnicodeFromString("rptgus");
        _1661_v3 = (_1661_v3).update(_1662_v4, _1658_v0);
        let _1663_v5;
        let _nw251 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1663_v5 = _nw251;
        let _1664_v6;
        let _nw252 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _1664_v6 = _nw252;
        let _index218 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1664_v6).length));
        (_1664_v6)[_index218] = _1658_v0;
        let _1665_v7;
        let _nw253 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
        _1665_v7 = _nw253;
        let _index219 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1664_v6).length));
        let _rhs255 = ((p1) ? (_1663_v5) : (_1663_v5));
        let _rhs256 = _1658_v0;
        let _rhs257 = _1665_v7;
        let _lhs187 = _1664_v6;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1664_v6).length));
        _1663_v5 = _rhs255;
        _lhs187[_lhs188] = _rhs256;
        _1665_v7 = _rhs257;
        _1658_v0 = (_module.__default.fm12(_this.f23, p1, _1658_v0, globalState)).dtor_cf1;
        _1662_v4 = _1662_v4;
        let _index220 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1663_v5).length));
        (_1663_v5)[_index220] = (_this).f24;
        let _1666_v8;
        _1666_v8 = _dafny.Set.fromElements((_1664_v6)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_1664_v6).length))]);
        let _index221 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1663_v5).length));
        (_1663_v5)[_index221] = (_1666_v8).IsDisjointFrom(_1666_v8);
      } else {
        let _1667_v9;
        _1667_v9 = _dafny.MultiSet.fromElements(p1);
        let _1668_v10;
        _1668_v10 = _dafny.Map.Empty.slice().updateUnsafe((p1) || ((_this).f24),(_1667_v9).contains((_this).f24));
        let _1669_v11;
        let _nw254 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1669_v11 = _nw254;
        let _1670_v12;
        let _nw255 = new _module.C7();
        _nw255.__ctor(p1, p1, _1669_v11, new _dafny.CodePoint('j'.codePointAt(0)), p1, true);
        _1670_v12 = _nw255;
        let _1671_v13;
        _1671_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1670_v12,_this.f23);
        let _1672_v14;
        _1672_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1658_v0,_1671_v13);
        _1668_v10 = (_1668_v10).update((_dafny.Map.Empty.slice().updateUnsafe(_1658_v0,_1671_v13)).equals(_1672_v14), (_this).f24);
        (globalState).f7 = _1658_v0;
        if (p1) {
          let _1673_v15;
          _1673_v15 = _dafny.Seq.UnicodeFromString("hapquyumu");
          let _1674_v16;
          _1674_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1673_v15);
          let _1675_v17;
          let _nw256 = new _module.C0();
          _nw256.__ctor((_1674_v16).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_1673_v15)), p1, _this.f23, p1);
          _1675_v17 = _nw256;
          _1675_v17 = _1675_v17;
          let _pat_let_tv32 = _1670_v12;
          r2 = (function (_pat_let47_0) {
            return function (_1676_dt__update__tmp_h0) {
              return function (_pat_let48_0) {
                return function (_1677_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2(_1677_dt__update_hcf2_h0, (_1676_dt__update__tmp_h0).dtor_cf3, (_1676_dt__update__tmp_h0).dtor_cf4);
                }(_pat_let48_0);
              }((_pat_let_tv32).f27);
            }(_pat_let47_0);
          }(_1659_v1)).dtor_cf4;
          let _1678_v18;
          let _init46 = ((_1679_v0) => function (_1680_i0) {
            return (_1680_i0).plus(_1679_v0);
          })(_1658_v0);
          let _nw257 = Array((new BigNumber(19)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw257.length); _i0_46++) {
            _nw257[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1678_v18 = _nw257;
          let _index222 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1678_v18).length));
          (_1678_v18)[_index222] = _1658_v0;
          let _1681_v19;
          _1681_v19 = _dafny.Seq.of((_dafny.ZERO).minus((_1658_v0).plus(_1658_v0)));
          let _index223 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1678_v18).length));
          let _rhs258 = new BigNumber((_1681_v19).length);
          let _rhs259 = _1658_v0;
          let _lhs189 = _1678_v18;
          let _lhs190 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1678_v18).length));
          _lhs189[_lhs190] = _rhs258;
          _1658_v0 = _rhs259;
          let _1682_v20;
          let _nw258 = new _module.C8();
          _nw258.__ctor((_this).f24);
          _1682_v20 = _nw258;
          _1658_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_1678_v18)[_module.__default.safeIndex(new BigNumber(135), new BigNumber((_1678_v18).length))], _1658_v0));
        } else {
          let _1683_v21;
          let _nw259 = new _module.C6();
          _nw259.__ctor(p0, !((_1670_v12).f27));
          _1683_v21 = _nw259;
          let _1684_v22;
          _1684_v22 = _dafny.Seq.of(_1669_v11);
          let _1685_v23;
          _1685_v23 = _module.D14.create_DC31((_1670_v12).f27);
          let _1686_v24;
          let _nw260 = new _module.C7();
          _nw260.__ctor(false, (_1670_v12).f27, (_1684_v22)[_module.__default.safeIndex((_dafny.ZERO).minus(_1658_v0), new BigNumber((_1684_v22).length))], new _dafny.CodePoint('w'.codePointAt(0)), (_this).f24, (_1685_v23).dtor_cf52);
          _1686_v24 = _nw260;
          let _1687_v25;
          let _nw261 = Array((new BigNumber(2)).toNumber()).fill(false);
          _1687_v25 = _nw261;
          let _1688_v26;
          let _nw262 = new _module.C5();
          _nw262.__ctor(_1687_v25);
          _1688_v26 = _nw262;
          let _1689_v27;
          let _init47 = function (_1690_i1) {
            return _module.D15.create_DC34(false);
          };
          let _nw263 = Array((new BigNumber(6)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw263.length); _i0_47++) {
            _nw263[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1689_v27 = _nw263;
          let _1691_v28;
          _1691_v28 = _dafny.Seq.of(_1689_v27);
          let _1692_v29;
          _1692_v29 = _module.D16.create_DC37(_1688_v26);
          let _1693_v30;
          _1693_v30 = _module.D16.create_DC37((_1692_v29).dtor_cf61);
          let _rhs260 = new BigNumber((_1691_v28).length);
          let _rhs261 = _1658_v0;
          let _rhs262 = (_1693_v30).dtor_cf61;
          let _rhs263 = (_this).fm6(globalState);
          let _rhs264 = true;
          let _lhs191 = globalState;
          let _lhs192 = globalState;
          let _lhs193 = globalState;
          _1658_v0 = _rhs260;
          _lhs191.f7 = _rhs261;
          _1688_v26 = _rhs262;
          _lhs192.f7 = _rhs263;
          _lhs193.f13 = _rhs264;
          let _1694_v31;
          _1694_v31 = _dafny.Seq.UnicodeFromString("dtuyeu");
          r0 = _1694_v31;
          r2 = ((_1670_v12).f27) || (false);
        }
        let _1695_v32;
        let _nw264 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _1695_v32 = _nw264;
        let _1696_v33;
        _1696_v33 = _dafny.Seq.of(_1695_v32);
        let _1697_v34;
        let _nw265 = Array((new BigNumber(17)).toNumber());
        _nw265[(_dafny.ZERO).toNumber()] = _1695_v32;
        _nw265[(_dafny.ONE).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(2)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(3)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(4)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(5)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(6)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(7)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(8)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(9)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(10)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(11)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(12)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(13)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(14)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(15)).toNumber()] = _1695_v32;
        _nw265[(new BigNumber(16)).toNumber()] = (_1696_v33)[_module.__default.safeIndex(_1658_v0, new BigNumber((_1696_v33).length))];
        _1697_v34 = _nw265;
        let _1698_v35;
        _1698_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1697_v34);
        _1698_v35 = (_1698_v35).update(new _dafny.CodePoint('y'.codePointAt(0)), _1697_v34);
        let _1699_v36;
        _1699_v36 = _dafny.MultiSet.fromElements(p0);
        let _1700_v37;
        _1700_v37 = _dafny.MultiSet.fromElements(_1699_v36);
        let _1701_v38;
        let _nw266 = new _module.C0();
        _nw266.__ctor((_module.__default.fm48(globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1670_v12).f27,_dafny.Seq.UnicodeFromString("ufkfma"))), (_this).f24, _this.f23, (_1700_v37).IsProperSubsetOf(_1700_v37));
        _1701_v38 = _nw266;
      }
      let _1702_v39;
      _1702_v39 = _module.D5.create_DC13(_1658_v0, !((_this).f24), (_this).f24);
      let _source25 = _1702_v39;
      if (_source25.is_DC13) {
        let _1703___mcc_h0 = (_source25).cf20;
        let _1704___mcc_h1 = (_source25).cf21;
        let _1705___mcc_h2 = (_source25).cf22;
        let _1706_cf22 = _1705___mcc_h2;
        let _1707_cf21 = _1704___mcc_h1;
        let _1708_cf20 = _1703___mcc_h0;
        let _1709_v40;
        _1709_v40 = _dafny.Seq.UnicodeFromString("ntvyfg");
        (globalState).f13 = _dafny.areEqual(_1709_v40, _dafny.Seq.Concat(_1709_v40, _1709_v40));
        let _1710_v41;
        let _nw267 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1710_v41 = _nw267;
        let _index224 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1710_v41).length));
        (_1710_v41)[_index224] = (_this).fm6(globalState);
        let _1711_v42;
        _1711_v42 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p1)),_1658_v0);
        let _1712_v43;
        _1712_v43 = _dafny.Seq.of(_1711_v42);
        let _1713_v44;
        _1713_v44 = _dafny.Set.fromElements(p0, _this.f23, _this.f23, _this.f23, _this.f23);
        let _index225 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1710_v41).length));
        (_1710_v41)[_index225] = new BigNumber((_dafny.Seq.Concat(_1712_v43, _dafny.Seq.update(_1712_v43, _module.__default.safeIndex(new BigNumber((_1713_v44).length), new BigNumber((_1712_v43).length)), _1711_v42))).length);
        let _1714_v45;
        _1714_v45 = _dafny.MultiSet.fromElements(p1);
        let _1715_v46;
        _1715_v46 = _dafny.Seq.of(_1714_v45);
        let _1716_v47;
        let _nw268 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1716_v47 = _nw268;
        let _1717_v48;
        let _nw269 = new _module.C2();
        _nw269.__ctor((_1715_v46)[_module.__default.safeIndex((_1710_v41)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_1710_v41).length))], new BigNumber((_1715_v46).length))], _1716_v47, new _dafny.CodePoint('s'.codePointAt(0)), p1);
        _1717_v48 = _nw269;
        _1717_v48 = _1717_v48;
        let _1718_v49;
        _1718_v49 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1710_v41)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_1710_v41).length))]),_1707_cf21);
        _1718_v49 = (_1718_v49).update((_1710_v41)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_1710_v41).length))], _1706_cf22);
      } else if (_source25.is_DC14) {
        let _1719___mcc_h3 = (_source25).cf23;
        let _1720_cf23 = _1719___mcc_h3;
        (globalState).f13 = _1720_cf23;
        _1659_v1 = _1659_v1;
        if ((((p1) ? ((_this).fm11(_1658_v0, (_this).f24, globalState)) : (p1))) && ((_this).f24)) {
          let _1721_v50;
          _1721_v50 = _dafny.Seq.of(p1, false, (_this).f24, _1720_cf23, _1720_cf23);
          let _1722_v51;
          _1722_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1658_v0,_this.f23);
          let _1723_v52;
          _1723_v52 = _dafny.Map.Empty.slice().updateUnsafe((_1721_v50)[_module.__default.safeIndex(new BigNumber((_1722_v51).length), new BigNumber((_1721_v50).length))],_1720_cf23);
          r2 = ((false) ? ((_this).f24) : ((_this).fm4(new BigNumber((_1723_v52).length), _module.__default.fm36(globalState), globalState)));
          let _1724_v53;
          _1724_v53 = _dafny.Seq.UnicodeFromString("hj");
          r0 = _1724_v53;
          (globalState).f13 = _1720_cf23;
          (globalState).f7 = _1658_v0;
          let _1725_v54;
          _1725_v54 = _dafny.MultiSet.fromElements(!(false));
          let _1726_v55;
          let _1727_v56;
          let _1728_v57;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector15 = _module.__default.m0(_1658_v0, (_dafny.MultiSet.fromElements(p1, (_this).f24)).Intersect(_1725_v54), p1, _module.__default.safeDivisionInt(_1658_v0, new BigNumber(878)), globalState);
          _out49 = _outcollector15[0];
          _out50 = _outcollector15[1];
          _out51 = _outcollector15[2];
          _1726_v55 = _out49;
          _1727_v56 = _out50;
          _1728_v57 = _out51;
        } else {
          let _1729_v58;
          let _init48 = function (_1730_i2) {
            return false;
          };
          let _nw270 = Array((new BigNumber(4)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw270.length); _i0_48++) {
            _nw270[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1729_v58 = _nw270;
          let _1731_v59;
          _1731_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1729_v58,p0);
          let _1732_v60;
          _1732_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_1658_v0, _1658_v0),_1731_v59);
          _1732_v60 = (_1732_v60).update(_1658_v0, _1731_v59);
          let _1733_v61;
          let _init49 = ((_1734_v0) => function (_1735_i3) {
            return _module.__default.safeModuloInt(_1735_i3, _1734_v0);
          })(_1658_v0);
          let _nw271 = Array((new BigNumber(8)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw271.length); _i0_49++) {
            _nw271[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1733_v61 = _nw271;
          let _1736_v62;
          _1736_v62 = _dafny.Seq.of(p1);
          let _index226 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1733_v61).length));
          (_1733_v61)[_index226] = (_1658_v0).multipliedBy(_module.__default.fm0(_1720_cf23, new BigNumber((_1736_v62).length), globalState));
          let _index227 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1733_v61).length));
          (_1733_v61)[_index227] = _1658_v0;
          (globalState).f7 = _1658_v0;
          let _1737_v63;
          _1737_v63 = _dafny.Seq.of(p0, p0);
          let _1738_v64;
          let _nw272 = Array((new BigNumber(27)).toNumber());
          _nw272[(_dafny.ZERO).toNumber()] = p0;
          _nw272[(_dafny.ONE).toNumber()] = _this.f23;
          _nw272[(new BigNumber(2)).toNumber()] = (_1737_v63)[_module.__default.safeIndex(new BigNumber((_1737_v63).length), new BigNumber((_1737_v63).length))];
          _nw272[(new BigNumber(3)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(4)).toNumber()] = p0;
          _nw272[(new BigNumber(5)).toNumber()] = p0;
          _nw272[(new BigNumber(6)).toNumber()] = (_1737_v63)[_module.__default.safeIndex(_1658_v0, new BigNumber((_1737_v63).length))];
          _nw272[(new BigNumber(7)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(8)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(9)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(10)).toNumber()] = p0;
          _nw272[(new BigNumber(11)).toNumber()] = p0;
          _nw272[(new BigNumber(12)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
          _nw272[(new BigNumber(14)).toNumber()] = p0;
          _nw272[(new BigNumber(15)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(16)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(17)).toNumber()] = p0;
          _nw272[(new BigNumber(18)).toNumber()] = _this.f23;
          _nw272[(new BigNumber(19)).toNumber()] = p0;
          _nw272[(new BigNumber(20)).toNumber()] = p0;
          _nw272[(new BigNumber(21)).toNumber()] = (_1737_v63)[_module.__default.safeIndex((_1733_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1733_v61).length))], new BigNumber((_1737_v63).length))];
          _nw272[(new BigNumber(22)).toNumber()] = p0;
          _nw272[(new BigNumber(23)).toNumber()] = p0;
          _nw272[(new BigNumber(24)).toNumber()] = _module.__default.fm29((_1733_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1733_v61).length))], p1, _1720_cf23, _1658_v0, globalState);
          _nw272[(new BigNumber(25)).toNumber()] = p0;
          _nw272[(new BigNumber(26)).toNumber()] = p0;
          _1738_v64 = _nw272;
          let _1739_v65;
          _1739_v65 = _module.D12.create_DC24(_1738_v64);
          let _1740_v66;
          let _nw273 = new _module.C7();
          _nw273.__ctor(((true) ? ((_this).f24) : (true)), (_this).f24, (_1739_v65).dtor_cf40, _this.f23, !((_1658_v0).isLessThan((_1733_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1733_v61).length))])), p1);
          _1740_v66 = _nw273;
          let _1741_v67;
          let _nw274 = Array((new BigNumber(6)).toNumber()).fill(_module.D13.Default());
          _1741_v67 = _nw274;
          let _1742_v68;
          _1742_v68 = _module.D13.create_DC29((_1740_v66).f30, (_1740_v66).f31);
          let _index228 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1741_v67).length));
          (_1741_v67)[_index228] = _1742_v68;
          let _index229 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1741_v67).length));
          (_1741_v67)[_index229] = _1742_v68;
        }
        _1720_cf23 = (_this).f24;
      } else {
        let _1743___mcc_h4 = (_source25).cf19;
        let _1744_cf19 = _1743___mcc_h4;
        let _1745_v69;
        let _init50 = ((_1746_v0) => function (_1747_i4) {
          return (_1747_i4).minus(_1746_v0);
        })(_1658_v0);
        let _nw275 = Array((new BigNumber(16)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw275.length); _i0_50++) {
          _nw275[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1745_v69 = _nw275;
        let _index230 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1745_v69).length));
        (_1745_v69)[_index230] = _1658_v0;
        let _index231 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1745_v69).length));
        (_1745_v69)[_index231] = _1658_v0;
        let _1748_v70;
        _1748_v70 = _dafny.Seq.of(_1658_v0);
        let _rhs265 = _dafny.Seq.Concat(_1748_v70, _1748_v70);
        let _rhs266 = !(p1);
        let _lhs194 = globalState;
        _1748_v70 = _rhs265;
        _lhs194.f13 = _rhs266;
        let _1749_v71;
        let _init51 = ((_1750_p1) => function (_1751_i5) {
          return (_module.D2.create_DC7(_1750_p1, (_this).f24)).dtor_cf14;
        })(p1);
        let _nw276 = Array((new BigNumber(5)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw276.length); _i0_51++) {
          _nw276[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1749_v71 = _nw276;
        let _1752_v72;
        _1752_v72 = _dafny.Seq.UnicodeFromString("cbxbjje");
        let _1753_v73;
        _1753_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1752_v72);
        let _1754_v74;
        _1754_v74 = _dafny.Seq.of((((_1753_v73).contains(false)) ? ((_1753_v73).get(false)) : (_1752_v72)));
        let _1755_v75;
        _1755_v75 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),p1);
        let _1756_v76;
        _1756_v76 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_1754_v74).length), new BigNumber((_1755_v75).length)),_1749_v71);
        _1749_v71 = (((_1756_v76).contains(_1658_v0)) ? ((_1756_v76).get(_1658_v0)) : (_1749_v71));
        let _1757_v77;
        _1757_v77 = _module.D0.create_DC1(new BigNumber(-87));
        r1 = _1757_v77;
      }
      if (p1) {
        let _1758_v78;
        let _init52 = ((_1759_v0) => function (_1760_i6) {
          return (_1760_i6).minus(_1759_v0);
        })(_1658_v0);
        let _nw277 = Array((new BigNumber(26)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw277.length); _i0_52++) {
          _nw277[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1758_v78 = _nw277;
        (globalState).f22 = _1758_v78;
        let _1761_v79;
        _1761_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
        (globalState).f7 = _module.__default.fm0(p1, new BigNumber((_1761_v79).length), globalState);
        (globalState).f13 = (_this).f24;
        let _1762_v80;
        let _nw278 = new _module.C6();
        _nw278.__ctor(_this.f23, (_this).f24);
        _1762_v80 = _nw278;
        _1762_v80 = _1762_v80;
        (globalState).f13 = ((false) ? ((_1762_v80).f24) : ((_this).f24));
      } else {
        let _rhs267 = new BigNumber(86);
        let _rhs268 = ((p1) ? ((_this).fm11(_1658_v0, p1, globalState)) : (!(p1) || (p1)));
        let _rhs269 = (_this).fm11(new BigNumber(-862), (_this).f24, globalState);
        let _rhs270 = (_this).fm11(_1658_v0, p1, globalState);
        let _rhs271 = new BigNumber(9);
        let _lhs195 = globalState;
        let _lhs196 = globalState;
        let _lhs197 = globalState;
        let _lhs198 = globalState;
        _lhs195.f7 = _rhs267;
        r2 = _rhs268;
        _lhs196.f13 = _rhs269;
        _lhs197.f13 = _rhs270;
        _lhs198.f7 = _rhs271;
        let _1763_v81;
        _1763_v81 = _dafny.Seq.UnicodeFromString("c");
        let _1764_v82;
        let _init53 = function (_1765_i7) {
          return (_this).f24;
        };
        let _nw279 = Array((new BigNumber(14)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw279.length); _i0_53++) {
          _nw279[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1764_v82 = _nw279;
        let _1766_v83;
        _1766_v83 = _dafny.Set.fromElements((_this).f24, (_this).f24);
        let _1767_v84;
        _1767_v84 = _module.D16.create_DC38(_dafny.Seq.Concat(_1763_v81, _1763_v81), _1764_v82, _1766_v83);
        let _source26 = _1767_v84;
        if (_source26.is_DC38) {
          let _1768___mcc_h5 = (_source26).cf62;
          let _1769___mcc_h6 = (_source26).cf63;
          let _1770___mcc_h7 = (_source26).cf64;
          let _1771_cf64 = _1770___mcc_h7;
          let _1772_cf63 = _1769___mcc_h6;
          let _1773_cf62 = _1768___mcc_h5;
          let _1774_v85;
          _1774_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(524),_1658_v0);
          let _1775_v86;
          _1775_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1658_v0,new BigNumber((_1774_v85).length));
          let _1776_v87;
          _1776_v87 = _dafny.Map.Empty.slice().updateUnsafe((_1658_v0).minus(new BigNumber((_1775_v86).length)),_1658_v0);
          _1774_v85 = _1774_v85;
          let _1777_v88;
          let _nw280 = new _module.C5();
          _nw280.__ctor(_1772_cf63);
          _1777_v88 = _nw280;
          let _1778_v89;
          let _nw281 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1778_v89 = _nw281;
          let _index232 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1778_v89).length));
          (_1778_v89)[_index232] = (_dafny.ZERO).minus(_1658_v0);
          let _1779_v90;
          _1779_v90 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p1);
          let _1780_v91;
          _1780_v91 = _module.D5.create_DC14(p1);
          let _1781_v92;
          _1781_v92 = _module.D11.create_DC23(p1, new BigNumber((_dafny.Seq.of(p1, !(p1))).length), (_1779_v90).update((_this).f24, !((_this).f24)), _1780_v91);
          let _index233 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1778_v89).length));
          let _rhs272 = _1658_v0;
          let _rhs273 = (_1781_v92).dtor_cf37;
          let _rhs274 = (_this).f24;
          let _lhs199 = _1778_v89;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1778_v89).length));
          let _lhs201 = globalState;
          let _lhs202 = globalState;
          _lhs199[_lhs200] = _rhs272;
          _lhs201.f7 = _rhs273;
          _lhs202.f13 = _rhs274;
          _1779_v90 = (_1779_v90).update((_this).f24, (_this).f24);
        } else if (_source26.is_DC37) {
          let _1782___mcc_h8 = (_source26).cf61;
          let _1783_cf61 = _1782___mcc_h8;
          let _1784_v93;
          let _nw282 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1784_v93 = _nw282;
          let _index234 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_1784_v93).length));
          (_1784_v93)[_index234] = _1658_v0;
          let _index235 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_1784_v93).length));
          (_1784_v93)[_index235] = (_1658_v0).multipliedBy(_1658_v0);
          (_this).f23 = p0;
          _1658_v0 = _module.__default.fm0(!((_module.__default.fm0(true, new BigNumber(364), globalState)).isLessThanOrEqualTo(_1658_v0)), (_1784_v93)[_module.__default.safeIndex(new BigNumber(374), new BigNumber((_1784_v93).length))], globalState);
          let _1785_v94;
          let _nw283 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1785_v94 = _nw283;
          let _1786_v95;
          _1786_v95 = _dafny.MultiSet.fromElements(_1658_v0);
          let _index236 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1785_v94).length));
          (_1785_v94)[_index236] = (_1786_v95).Intersect(_1786_v95);
          let _index237 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1785_v94).length));
          (_1785_v94)[_index237] = _1786_v95;
        } else {
          let _1787___mcc_h9 = (_source26).cf65;
          let _1788_cf65 = _1787___mcc_h9;
          (globalState).f13 = !(p1) || (false);
          let _1789_v96;
          _1789_v96 = _dafny.Seq.of(_1658_v0);
          let _1790_v97;
          let _nw284 = Array((new BigNumber(25)).toNumber());
          _nw284[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1763_v81).length));
          _nw284[(_dafny.ONE).toNumber()] = (_1789_v96)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_1789_v96).length))];
          _nw284[(new BigNumber(2)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(3)).toNumber()] = new BigNumber(521);
          _nw284[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_1658_v0);
          _nw284[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_1789_v96)).cardinality());
          _nw284[(new BigNumber(6)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_1658_v0);
          _nw284[(new BigNumber(8)).toNumber()] = (_1659_v1).dtor_cf3;
          _nw284[(new BigNumber(9)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(10)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(11)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(12)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(13)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(14)).toNumber()] = _module.__default.fm0(!(p1), _1658_v0, globalState);
          _nw284[(new BigNumber(15)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(16)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(17)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(18)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(19)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(20)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(21)).toNumber()] = (_1658_v0).multipliedBy(_1658_v0);
          _nw284[(new BigNumber(22)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(23)).toNumber()] = _1658_v0;
          _nw284[(new BigNumber(24)).toNumber()] = (_1658_v0).plus(new BigNumber((_1763_v81).length));
          _1790_v97 = _nw284;
          let _1791_v98;
          _1791_v98 = _module.D5.create_DC14(p1);
          let _1792_v99;
          _1792_v99 = _module.D11.create_DC23(p1, _1658_v0, _dafny.Map.Empty.slice().updateUnsafe(p1,p1), _1791_v98);
          let _1793_v100;
          _1793_v100 = _dafny.Seq.of(_1792_v99, _1792_v99, _1792_v99, _1792_v99, _1792_v99);
          let _index238 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1790_v97).length));
          (_1790_v97)[_index238] = new BigNumber((((_dafny.MultiSet.FromArray(_1793_v100)).update(_1792_v99, _module.__default.abs(_1658_v0))).Intersect(_dafny.MultiSet.fromElements(_1792_v99))).cardinality());
          let _index239 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1790_v97).length));
          (_1790_v97)[_index239] = _module.__default.safeDivisionInt(_1658_v0, ((_dafny.ZERO).minus(_1658_v0)).minus(_1658_v0));
          _1658_v0 = new BigNumber(375);
          let _index240 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1790_v97).length));
          (_1790_v97)[_index240] = _1658_v0;
        }
        let _1794_v101;
        let _nw285 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _1794_v101 = _nw285;
        let _index241 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1794_v101).length));
        (_1794_v101)[_index241] = _1658_v0;
        let _1795_v102;
        _1795_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1658_v0);
        let _1796_v103;
        _1796_v103 = _dafny.Seq.of(p1);
        let _index242 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1794_v101).length));
        let _rhs275 = (((_1795_v102).contains((_this).f24)) ? ((_1795_v102).get((_this).f24)) : (new BigNumber((_1796_v103).length)));
        let _rhs276 = _1658_v0;
        let _lhs203 = globalState;
        let _lhs204 = _1794_v101;
        let _lhs205 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1794_v101).length));
        _lhs203.f7 = _rhs275;
        _lhs204[_lhs205] = _rhs276;
        let _1797_v104;
        let _nw286 = new _module.C8();
        _nw286.__ctor((_this).f24);
        _1797_v104 = _nw286;
        let _1798_v105;
        _1798_v105 = _dafny.Seq.of(_1658_v0, (_1794_v101)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1794_v101).length))]);
        let _1799_v106;
        _1799_v106 = _dafny.Seq.of(_1658_v0, (_1798_v105)[_module.__default.safeIndex((_1794_v101)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1794_v101).length))], new BigNumber((_1798_v105).length))]);
        let _1800_v107;
        _1800_v107 = _dafny.MultiSet.fromElements(_1658_v0, new BigNumber((_1799_v106).length), _1658_v0, _1658_v0);
        r2 = ((!((_this).f24)) ? ((_dafny.MultiSet.fromElements((_1794_v101)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1794_v101).length))])).IsDisjointFrom(_1800_v107)) : ((_this).f24));
      }
      let _1801_v108;
      _1801_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1702_v39,(_this).f24);
      _1801_v108 = (_1801_v108).update(_1702_v39, (_this).f24);
      if ((_this).f24) {
        let _1802_v109;
        let _nw287 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1802_v109 = _nw287;
        let _1803_v110;
        _1803_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1802_v109,new _dafny.CodePoint('y'.codePointAt(0)));
        let _1804_v111;
        _1804_v111 = _dafny.Seq.UnicodeFromString("hu");
        let _1805_v112;
        let _nw288 = Array((new BigNumber(11)).toNumber());
        _nw288[(_dafny.ZERO).toNumber()] = _this.f23;
        _nw288[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw288[(new BigNumber(2)).toNumber()] = p0;
        _nw288[(new BigNumber(3)).toNumber()] = p0;
        _nw288[(new BigNumber(4)).toNumber()] = _this.f23;
        _nw288[(new BigNumber(5)).toNumber()] = p0;
        _nw288[(new BigNumber(6)).toNumber()] = _this.f23;
        _nw288[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw288[(new BigNumber(8)).toNumber()] = p0;
        _nw288[(new BigNumber(9)).toNumber()] = p0;
        _nw288[(new BigNumber(10)).toNumber()] = (_1804_v111)[_module.__default.safeIndex(_1658_v0, new BigNumber((_1804_v111).length))];
        _1805_v112 = _nw288;
        let _1806_v113;
        let _nw289 = new _module.C3();
        _nw289.__ctor(_1803_v110, _1805_v112, p0, (_this).f24);
        _1806_v113 = _nw289;
        let _1807_v114;
        _1807_v114 = _dafny.Seq.of(_1658_v0, _1658_v0, new BigNumber(129), _1658_v0);
        let _1808_v115;
        _1808_v115 = _dafny.Map.Empty.slice().updateUnsafe(_1806_v113,_module.__default.safeDivisionInt((_1807_v114)[_module.__default.safeIndex(_1658_v0, new BigNumber((_1807_v114).length))], _1658_v0));
        _1808_v115 = _1808_v115;
        let _1809_v116;
        let _nw290 = new _module.C7();
        _nw290.__ctor(false, (_this).f24, _1805_v112, p0, p1, !(true));
        _1809_v116 = _nw290;
        let _1810_v117;
        let _nw291 = Array((new BigNumber(7)).toNumber());
        _nw291[(_dafny.ZERO).toNumber()] = _1809_v116;
        _nw291[(_dafny.ONE).toNumber()] = _1809_v116;
        _nw291[(new BigNumber(2)).toNumber()] = _1809_v116;
        _nw291[(new BigNumber(3)).toNumber()] = _1809_v116;
        _nw291[(new BigNumber(4)).toNumber()] = _1809_v116;
        _nw291[(new BigNumber(5)).toNumber()] = _1809_v116;
        _nw291[(new BigNumber(6)).toNumber()] = _1809_v116;
        _1810_v117 = _nw291;
        let _index243 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1810_v117).length));
        (_1810_v117)[_index243] = _1809_v116;
        let _index244 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1810_v117).length));
        (_1810_v117)[_index244] = (((_1809_v116).f27) ? (_1809_v116) : (_1809_v116));
        let _1811_v118;
        _1811_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1804_v111,((!((_1809_v116).f27)) ? (_1802_v109) : (_1802_v109)));
        let _1812_v119;
        _1812_v119 = _dafny.Seq.of((_this).f24, (_this).f24);
        let _1813_v120;
        _1813_v120 = _module.D4.create_DC10(_1804_v111);
        let _1814_v121;
        _1814_v121 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1804_v111);
        let _1815_v122;
        _1815_v122 = _dafny.Seq.of((((_1814_v121).contains(p1)) ? ((_1814_v121).get(p1)) : (_dafny.Seq.UnicodeFromString("fpa"))), _1804_v111, _dafny.Seq.UnicodeFromString("bfygf"));
        let _1816_v123;
        _1816_v123 = _dafny.MultiSet.fromElements((_1809_v116).f27);
        let _1817_v124;
        let _nw292 = new _module.C2();
        _nw292.__ctor(_1816_v123, _1805_v112, _this.f23, (_this).f24);
        _1817_v124 = _nw292;
        let _1818_v125;
        _1818_v125 = _dafny.Seq.of(_1817_v124, _1817_v124);
        let _rhs277 = ((_this).fm4((_this).fm6(globalState), _1812_v119, globalState)) || (p1);
        let _rhs278 = _1811_v118;
        let _rhs279 = (_1658_v0).minus(new BigNumber((_dafny.Seq.of(_1813_v120)).length));
        let _rhs280 = (_this).f24;
        let _rhs281 = (_1815_v122)[_module.__default.safeIndex(_module.__default.safeModuloInt(_1658_v0, new BigNumber((_1818_v125).length)), new BigNumber((_1815_v122).length))];
        let _lhs206 = globalState;
        let _lhs207 = globalState;
        r2 = _rhs277;
        _1811_v118 = _rhs278;
        _lhs206.f7 = _rhs279;
        _lhs207.f13 = _rhs280;
        _1804_v111 = _rhs281;
        let _1819_v126;
        let _nw293 = new _module.C1();
        _nw293.__ctor(p1, _dafny.areEqual(_dafny.Seq.of(_module.__default.fm2(globalState), (_1809_v116).f27), _1812_v119));
        _1819_v126 = _nw293;
        r2 = (_this).f24;
      } else {
        let _1820_v127;
        _1820_v127 = _dafny.Map.Empty.slice().updateUnsafe((_1658_v0).multipliedBy(_1658_v0),_1658_v0);
        _1820_v127 = _1820_v127;
        (globalState).f7 = (_dafny.ZERO).minus(_1658_v0);
        _1659_v1 = _1659_v1;
        let _1821_v128;
        let _init54 = ((_1822_v0) => function (_1823_i8) {
          return (_1823_i8).multipliedBy(_1822_v0);
        })(_1658_v0);
        let _nw294 = Array((new BigNumber(12)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw294.length); _i0_54++) {
          _nw294[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1821_v128 = _nw294;
        let _1824_v129;
        _1824_v129 = _dafny.Set.fromElements(_1658_v0);
        let _index245 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length));
        (_1821_v128)[_index245] = new BigNumber((_1824_v129).length);
        let _index246 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length));
        (_1821_v128)[_index246] = ((((_this).f24) ? (_1658_v0) : (_1658_v0))).plus(_1658_v0);
        let _1825_v130;
        _1825_v130 = _dafny.MultiSet.fromElements(_1658_v0);
        if (((_1825_v130).Difference(_dafny.MultiSet.fromElements((_this).fm6(globalState), _1658_v0))).IsSubsetOf(_1825_v130)) {
          _1820_v127 = (_1820_v127).update((_dafny.ZERO).minus(((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))]).plus((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))])), (_dafny.ZERO).minus(_1658_v0));
          let _1826_v131;
          let _nw295 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1826_v131 = _nw295;
          let _1827_v132;
          let _nw296 = new _module.C5();
          _nw296.__ctor(_1826_v131);
          _1827_v132 = _nw296;
          r2 = (_this).f24;
          (globalState).f13 = (_module.D5.create_DC13(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_1658_v0))).length), !((_this).f24), p1)).dtor_cf22;
          let _1828_v133;
          _1828_v133 = _dafny.MultiSet.fromElements((_this).f24);
          let _1829_v134;
          _1829_v134 = _module.D0.create_DC0(_1828_v133);
          _1829_v134 = _1829_v134;
        } else {
          (globalState).f7 = _1658_v0;
          let _1830_v135;
          _1830_v135 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f24, (_this).f24, p1),p1);
          let _1831_v136;
          _1831_v136 = _dafny.Seq.of(p1);
          let _1832_v137;
          _1832_v137 = _dafny.Seq.UnicodeFromString("blrabrw");
          let _1833_v138;
          _1833_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1658_v0,_1832_v137);
          let _1834_v139;
          _1834_v139 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(new BigNumber((_1830_v135).length), _1831_v136, globalState),(((_1833_v138).contains(_1658_v0)) ? ((_1833_v138).get(_1658_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-330)), function (_1835_i9) {
            return _this.f23;
          }))));
          let _1836_v140;
          let _nw297 = new _module.C1();
          _nw297.__ctor((_this).f24, false);
          _1836_v140 = _nw297;
          let _1837_v141;
          _1837_v141 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v140,_1658_v0);
          let _1838_v142;
          _1838_v142 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v141,(_this).f24);
          let _1839_v143;
          _1839_v143 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f24),(_dafny.ZERO).minus(new BigNumber((_1838_v142).length)));
          let _1840_v144;
          _1840_v144 = _dafny.Map.Empty.slice().updateUnsafe(_1839_v143,p1);
          let _1841_v145;
          let _nw298 = new _module.C0();
          _nw298.__ctor(_1834_v139, (_this).f24, _module.__default.fm29((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))], p1, p1, (_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))], globalState), (_1840_v144).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1831_v136).length)),(_1836_v140).f39)));
          _1841_v145 = _nw298;
          (_this).f23 = (_1832_v137)[_module.__default.safeIndex(_1658_v0, new BigNumber((_1832_v137).length))];
          let _1842_v146;
          _1842_v146 = _dafny.Seq.of((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))], ((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))]).minus((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))]), _1658_v0, (_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))]);
          (globalState).f7 = (_1842_v146)[_module.__default.safeIndex((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))], new BigNumber((_1842_v146).length))];
          let _1843_v147;
          _1843_v147 = _dafny.MultiSet.fromElements((_1841_v145).f38);
          (globalState).f7 = ((!(_1843_v147).contains(!(!((_1836_v140).f39)))) ? ((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))]) : (_module.__default.safeModuloInt((_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))], (_1821_v128)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1821_v128).length))])));
        }
      }
      let _1844_v148;
      let _nw299 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1844_v148 = _nw299;
      let _1845_v149;
      _1845_v149 = _dafny.Seq.of(_1844_v148, _1844_v148);
      let _hi11 = (((_this).f24) ? (new BigNumber((_1845_v149).length)) : (_1658_v0));
      for (let _1846_i10 = _1658_v0; _1846_i10.isLessThan(_hi11); _1846_i10 = _1846_i10.plus(_dafny.ONE)) {
        let _1847_v150;
        _1847_v150 = _dafny.Seq.UnicodeFromString("enqsgx");
        (globalState).f13 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_module.__default.fm1((_this).f24, globalState), _1847_v150), _1847_v150);
        let _1848_v151;
        let _nw300 = Array((new BigNumber(16)).toNumber()).fill(false);
        _1848_v151 = _nw300;
        let _index247 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1848_v151).length));
        (_1848_v151)[_index247] = (_this).f24;
        let _1849_v152;
        _1849_v152 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1847_v150);
        let _1850_v153;
        _1850_v153 = _dafny.Seq.of(p1);
        let _1851_v154;
        let _nw301 = new _module.C0();
        _nw301.__ctor(_1849_v152, (_1850_v153)[_module.__default.safeIndex(_1846_i10, new BigNumber((_1850_v153).length))], _this.f23, !((_this).f24));
        _1851_v154 = _nw301;
        let _1852_v155;
        _1852_v155 = _dafny.MultiSet.fromElements(_1846_i10);
        let _index248 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1848_v151).length));
        let _rhs282 = false;
        let _rhs283 = (_dafny.ZERO).minus(_1846_i10);
        let _rhs284 = _1851_v154;
        let _rhs285 = (_1852_v155).Intersect(_dafny.MultiSet.fromElements(_1846_i10, _1846_i10, _1658_v0));
        let _rhs286 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1850_v153, _dafny.Seq.of(true)), _dafny.Seq.Concat(_1850_v153, _1850_v153));
        let _lhs208 = _1848_v151;
        let _lhs209 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1848_v151).length));
        _lhs208[_lhs209] = _rhs282;
        _1658_v0 = _rhs283;
        _1851_v154 = _rhs284;
        _1852_v155 = _rhs285;
        _1850_v153 = _rhs286;
        let _1853_v156;
        let _nw302 = new _module.C8();
        _nw302.__ctor(p1);
        _1853_v156 = _nw302;
        let _1854_v157;
        _1854_v157 = _dafny.MultiSet.fromElements(_1853_v156);
        let _1855_v158;
        _1855_v158 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("lufwkuo")).length), _1846_i10, new BigNumber(((_1854_v157).update(_1853_v156, _module.__default.abs(_1846_i10))).cardinality()));
        let _1856_v159;
        _1856_v159 = _dafny.Set.fromElements(new BigNumber((_1855_v158).length), _1846_i10, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(114)), ((_1857_v154) => function (_1858_i11) {
          return _1857_v154.f23;
        })(_1851_v154))).length));
        (globalState).f7 = ((p1) ? (_module.__default.safeDivisionInt(_1658_v0, new BigNumber((_1856_v159).length))) : (_1846_i10));
        let _1859_v160;
        _1859_v160 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("y"),_1851_v154.f23);
        (_1851_v154).f23 = (((_1859_v160).contains(_dafny.Seq.Concat(_1847_v150, _1847_v150))) ? ((_1859_v160).get(_dafny.Seq.Concat(_1847_v150, _1847_v150))) : (_1851_v154.f23));
      }
      r0 = _dafny.Seq.UnicodeFromString("oyguroyys");
      let _1860_v161;
      _1860_v161 = _module.D0.create_DC1(_1658_v0);
      r1 = _1860_v161;
      let _1861_v162;
      let _nw303 = new _module.C1();
      _nw303.__ctor(p1, p1);
      _1861_v162 = _nw303;
      let _1862_v163;
      _1862_v163 = _dafny.Set.fromElements(_1861_v162, _1861_v162);
      r2 = (_1862_v163).IsProperSubsetOf(_1862_v163);
      return [r0, r1, r2];
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _source27 = _module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_this.f23));
      if (_source27.is_DC28) {
        let _1863___mcc_h0 = (_source27).cf45;
        let _1864___mcc_h1 = (_source27).cf46;
        let _1865___mcc_h2 = (_source27).cf47;
        let _1866___mcc_h3 = (_source27).cf48;
        let _1867_cf48 = _1866___mcc_h3;
        let _1868_cf47 = _1865___mcc_h2;
        let _1869_cf46 = _1864___mcc_h1;
        let _1870_cf45 = _1863___mcc_h0;
        let _1871_v0;
        let _init55 = ((_1872_cf45) => function (_1873_i0) {
          return _1872_cf45;
        })(_1870_cf45);
        let _nw304 = Array((new BigNumber(27)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw304.length); _i0_55++) {
          _nw304[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _1871_v0 = _nw304;
        let _1874_v1;
        _1874_v1 = _dafny.Seq.UnicodeFromString("ig");
        let _1875_v2;
        _1875_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1871_v0,(_1874_v1)[_module.__default.safeIndex(_1868_cf47, new BigNumber((_1874_v1).length))]);
        let _1876_v3;
        let _nw305 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1876_v3 = _nw305;
        let _1877_v4;
        let _nw306 = new _module.C3();
        _nw306.__ctor(_1875_v2, _1876_v3, _this.f23, _1870_cf45);
        _1877_v4 = _nw306;
        let _1878_v5;
        _1878_v5 = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(true,_1874_v1));
        let _1879_v6;
        _1879_v6 = _module.__default.fm45((_this).f24, _1878_v5, new BigNumber((_1874_v1).length), globalState);
        let _1880_v7;
        _1880_v7 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f24) || (true),(_1879_v6));
        let _1881_v8;
        _1881_v8 = _dafny.Seq.of(new BigNumber(-557), _1868_cf47);
        _1880_v7 = (_1880_v7).update((_this).f24, _1881_v8);
        let _1882_v9;
        _1882_v9 = _module.D12.create_DC25((_this).f24, (_module.D1.create_DC4(_this.f23)).dtor_cf6);
        if (((_1870_cf45) ? ((_this).f24) : ((_1882_v9).dtor_cf41))) {
          _1882_v9 = _1882_v9;
          let _1883_v10;
          _1883_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1870_cf45,_dafny.Seq.UnicodeFromString("sworc"));
          let _1884_v11;
          let _nw307 = new _module.C0();
          _nw307.__ctor(_1883_v10, _1870_cf45, _this.f23, _1870_cf45);
          _1884_v11 = _nw307;
          let _1885_v12;
          _1885_v12 = _1884_v11;
          let _1886_v13;
          let _nw308 = Array((new BigNumber(14)).toNumber());
          _nw308[(_dafny.ZERO).toNumber()] = (_1885_v12);
          _nw308[(_dafny.ONE).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(2)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(3)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(4)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(5)).toNumber()] = (((_1884_v11).f24) ? (_1884_v11) : (_1884_v11));
          _nw308[(new BigNumber(6)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(7)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(8)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(9)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(10)).toNumber()] = (((_1884_v11).f24) ? (_1884_v11) : (_1884_v11));
          _nw308[(new BigNumber(11)).toNumber()] = _1884_v11;
          _nw308[(new BigNumber(12)).toNumber()] = (((_1884_v11).f24) ? (_1884_v11) : (_1884_v11));
          _nw308[(new BigNumber(13)).toNumber()] = _1884_v11;
          _1886_v13 = _nw308;
          let _1887_v14;
          _1887_v14 = _dafny.MultiSet.fromElements(_1867_cf48, (_dafny.ZERO).minus(new BigNumber((_1874_v1).length)));
          let _1888_v15;
          _1888_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1870_cf45,(_module.D5.create_DC13((_dafny.ZERO).minus(_1868_cf47), !((_this).f24), (_1884_v11).f24)).dtor_cf20);
          let _1889_v16;
          _1889_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1888_v15).length),(_this).f24);
          let _1890_v17;
          let _nw309 = Array((new BigNumber(19)).toNumber());
          _nw309[(_dafny.ZERO).toNumber()] = _1869_cf46;
          _nw309[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1874_v1).length), _1867_cf48);
          _nw309[(new BigNumber(2)).toNumber()] = _1867_cf48;
          _nw309[(new BigNumber(3)).toNumber()] = (_this).fm6(globalState);
          _nw309[(new BigNumber(4)).toNumber()] = _1867_cf48;
          _nw309[(new BigNumber(5)).toNumber()] = _1868_cf47;
          _nw309[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_1867_cf48, new BigNumber((_1874_v1).length));
          _nw309[(new BigNumber(7)).toNumber()] = new BigNumber(310);
          _nw309[(new BigNumber(8)).toNumber()] = new BigNumber(592);
          _nw309[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt((((_1887_v14).contains(_1869_cf46)) ? ((_1887_v14).get(_1869_cf46)) : (new BigNumber((_1889_v16).length))), _1868_cf47);
          _nw309[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_1869_cf46, _1868_cf47);
          _nw309[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_1867_cf48, new BigNumber((_1874_v1).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(_1867_cf48)), _1869_cf46, _1868_cf47)).length);
          _nw309[(new BigNumber(12)).toNumber()] = _1869_cf46;
          _nw309[(new BigNumber(13)).toNumber()] = _1868_cf47;
          _nw309[(new BigNumber(14)).toNumber()] = _1867_cf48;
          _nw309[(new BigNumber(15)).toNumber()] = new BigNumber(132);
          _nw309[(new BigNumber(16)).toNumber()] = _1868_cf47;
          _nw309[(new BigNumber(17)).toNumber()] = _1869_cf46;
          _nw309[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(_1867_cf48, _1868_cf47);
          _1890_v17 = _nw309;
          let _index249 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1890_v17).length));
          (_1890_v17)[_index249] = new BigNumber(-225);
          let _1891_v18;
          let _nw310 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _1891_v18 = _nw310;
          let _index250 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1891_v18).length));
          (_1891_v18)[_index250] = _module.__default.fm49(globalState);
          let _1892_v19;
          _1892_v19 = _module.D18.create_DC41(_1886_v13);
          let _1893_v20;
          _1893_v20 = _dafny.Seq.of((_1884_v11).f24);
          let _index251 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1890_v17).length));
          let _index252 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1891_v18).length));
          let _rhs287 = (_dafny.ZERO).minus(_1868_cf47);
          let _rhs288 = (_1892_v19).dtor_cf67;
          let _rhs289 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("adnjjpgt"), _1874_v1)).length), _1867_cf48);
          let _rhs290 = ((_this).fm4(_1868_cf47, _1893_v20, globalState)) && ((_this).f24);
          let _rhs291 = _1888_v15;
          let _lhs210 = globalState;
          let _lhs211 = _1890_v17;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1890_v17).length));
          let _lhs213 = globalState;
          let _lhs214 = _1891_v18;
          let _lhs215 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1891_v18).length));
          _lhs210.f7 = _rhs287;
          _1886_v13 = _rhs288;
          _lhs211[_lhs212] = _rhs289;
          _lhs213.f13 = _rhs290;
          _lhs214[_lhs215] = _rhs291;
          let _index253 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1890_v17).length));
          (_1890_v17)[_index253] = (_dafny.ZERO).minus(_1869_cf46);
          let _1894_v21;
          let _nw311 = new _module.C1();
          _nw311.__ctor((_this).f24, (_this).f24);
          _1894_v21 = _nw311;
          let _1895_v22;
          _1895_v22 = _1894_v21;
          _1895_v22 = _1895_v22;
          let _index254 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_1871_v0).length));
          (_1871_v0)[_index254] = ((_this).f24) && ((_this).f24);
          let _index255 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_1871_v0).length));
          (_1871_v0)[_index255] = (_this).f24;
        } else {
          let _1896_v23;
          _1896_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.UnicodeFromString("wjgyr")).length));
          let _1897_v24;
          _1897_v24 = _dafny.Set.fromElements((_this).f24);
          let _1898_v25;
          _1898_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1897_v24).length),new BigNumber(151));
          _1896_v23 = (_1896_v23).update((_this).f24, (((_1898_v25).contains(_1868_cf47)) ? ((_1898_v25).get(_1868_cf47)) : (_1868_cf47)));
          let _1899_v26;
          _1899_v26 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1867_cf48));
          let _1900_v27;
          _1900_v27 = _dafny.Set.fromElements(_1868_cf47, new BigNumber((_1899_v26).length));
          let _1901_v28;
          _1901_v28 = _dafny.Seq.of(_1870_cf45);
          let _1902_v29;
          _1902_v29 = _dafny.Set.fromElements(_1900_v27);
          let _1903_v30;
          _1903_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1901_v28,new BigNumber((_1902_v29).length));
          let _1904_v31;
          _1904_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1874_v1, _module.__default.safeIndex(new BigNumber((_1899_v26).length), new BigNumber((_1874_v1).length)), _this.f23),_1903_v30);
          _1904_v31 = (_1904_v31).update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uso"), _1874_v1), (function () {
            let _coll90 = new _dafny.Map();
            for (const _compr_90 of (_dafny.Seq.of(_1901_v28, _1901_v28, _1901_v28)).Elements) {
              let _1905_v32 = _compr_90;
              if (_dafny.Seq.contains(_dafny.Seq.of(_1901_v28, _1901_v28, _1901_v28), _1905_v32)) {
                _coll90.push([_1905_v32,new BigNumber(66)]);
              }
            }
            return _coll90;
          }()).Merge(_1903_v30));
          let _1906_v33;
          _1906_v33 = _dafny.MultiSet.fromElements(_1870_cf45, false);
          let _1907_v34;
          _1907_v34 = _dafny.Seq.of(_1876_v3);
          let _1908_v35;
          let _nw312 = new _module.C2();
          _nw312.__ctor(_1906_v33, (_1907_v34)[_module.__default.safeIndex(_1867_cf48, new BigNumber((_1907_v34).length))], _this.f23, _1870_cf45);
          _1908_v35 = _nw312;
          _1870_cf45 = (_this).f24;
          (globalState).f13 = (_this).f24;
        }
        let _1909_v36;
        _1909_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1871_v0);
        (globalState).f7 = new BigNumber((_1909_v36).length);
      } else if (_source27.is_DC29) {
        let _1910___mcc_h4 = (_source27).cf49;
        let _1911___mcc_h5 = (_source27).cf50;
        let _1912_cf50 = _1911___mcc_h5;
        let _1913_cf49 = _1910___mcc_h4;
        let _1914_v37;
        _1914_v37 = new BigNumber(-698);
        let _1915_v38;
        let _nw313 = Array((new BigNumber(7)).toNumber());
        _nw313[(_dafny.ZERO).toNumber()] = _1914_v37;
        _nw313[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus(_1914_v37)).multipliedBy(new BigNumber(944));
        _nw313[(new BigNumber(2)).toNumber()] = _1914_v37;
        _nw313[(new BigNumber(3)).toNumber()] = _1914_v37;
        _nw313[(new BigNumber(4)).toNumber()] = _1914_v37;
        _nw313[(new BigNumber(5)).toNumber()] = _1914_v37;
        _nw313[(new BigNumber(6)).toNumber()] = _1914_v37;
        _1915_v38 = _nw313;
        (globalState).f22 = _1915_v38;
        let _1916_v39;
        let _init56 = function (_1917_i1) {
          return false;
        };
        let _nw314 = Array((new BigNumber(16)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw314.length); _i0_56++) {
          _nw314[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _1916_v39 = _nw314;
        let _index256 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1916_v39).length));
        (_1916_v39)[_index256] = (_this).f24;
        let _index257 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1916_v39).length));
        (_1916_v39)[_index257] = _1913_cf49;
        let _1918_v40;
        let _init57 = ((_1919_cf50) => function (_1920_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1919_cf50,new BigNumber(736));
        })(_1912_cf50);
        let _nw315 = Array((new BigNumber(7)).toNumber());
        for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw315.length); _i0_57++) {
          _nw315[_i0_57] = _init57(new BigNumber(_i0_57));
        }
        _1918_v40 = _nw315;
        _1918_v40 = _1918_v40;
        (globalState).f13 = (_this).f24;
      } else {
        let _1921___mcc_h6 = (_source27).cf44;
        let _1922_cf44 = _1921___mcc_h6;
        let _1923_v41;
        _1923_v41 = _dafny.Seq.of(!((_this).f24));
        let _1924_v42;
        _1924_v42 = _dafny.Set.fromElements((_this).f24);
        let _1925_v43;
        _1925_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1923_v41,_1924_v42);
        let _1926_v44;
        _1926_v44 = new BigNumber(35);
        let _1927_v45;
        _1927_v45 = _dafny.Seq.UnicodeFromString("tghfl");
        let _rhs292 = _1925_v43;
        let _rhs293 = _1926_v44;
        let _rhs294 = _dafny.Seq.contains(_1927_v45, _this.f23);
        let _rhs295 = !((_this).f24);
        let _lhs216 = globalState;
        let _lhs217 = globalState;
        _1925_v43 = _rhs292;
        r1 = _rhs293;
        _lhs216.f13 = _rhs294;
        _lhs217.f13 = _rhs295;
        let _1928_v46;
        let _nw316 = new _module.C6();
        _nw316.__ctor(_this.f23, (_this).f24);
        _1928_v46 = _nw316;
        let _1929_v47;
        _1929_v47 = _dafny.Seq.of(_1928_v46, _1928_v46);
        _1929_v47 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1929_v47, _dafny.Seq.of(_1928_v46)), _1929_v47);
        if (((_1928_v46).f24) || ((_this).f24)) {
          (globalState).f7 = _1926_v44;
          let _1930_v48;
          let _nw317 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1930_v48 = _nw317;
          let _index258 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1930_v48).length));
          (_1930_v48)[_index258] = (_dafny.ZERO).minus(new BigNumber((_1923_v41).length));
          let _1931_v49;
          _1931_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1926_v44,(_this).f24);
          let _index259 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1930_v48).length));
          (_1930_v48)[_index259] = new BigNumber(((_1931_v49).Merge((((_this).f24) ? (_1931_v49) : (_1931_v49)))).length);
          let _1932_v50;
          _1932_v50 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_1928_v46).f24,(_1928_v46).f24));
          let _1933_v51;
          _1933_v51 = _dafny.MultiSet.fromElements((_this).f24);
          let _1934_v52;
          _1934_v52 = _dafny.Set.fromElements(_1926_v44);
          let _1935_v53;
          let _1936_v54;
          let _1937_v55;
          let _out52;
          let _out53;
          let _out54;
          let _outcollector16 = _module.__default.m0(new BigNumber((_1932_v50).cardinality()), (_1933_v51).update(true, _module.__default.abs(new BigNumber((_1934_v52).length))), (_1928_v46).f24, _1926_v44, globalState);
          _out52 = _outcollector16[0];
          _out53 = _outcollector16[1];
          _out54 = _outcollector16[2];
          _1935_v53 = _out52;
          _1936_v54 = _out53;
          _1937_v55 = _out54;
          let _1938_v56;
          let _nw318 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1938_v56 = _nw318;
          let _nw319 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1938_v56 = _nw319;
          _1923_v41 = _1923_v41;
        } else {
          _1926_v44 = _module.__default.safeDivisionInt(new BigNumber(726), _1926_v44);
          let _1939_v57;
          _1939_v57 = _dafny.Map.Empty.slice().updateUnsafe((_1928_v46).f24,_1923_v41);
          _1939_v57 = (_1939_v57).update((_1928_v46).f24, _1923_v41);
          (globalState).f13 = (_1928_v46).f24;
          let _1940_v58;
          _1940_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1926_v44,(_1928_v46).f24);
          let _1941_v59;
          _1941_v59 = _dafny.Seq.of(_module.__default.safeDivisionInt(_1926_v44, _1926_v44), new BigNumber((_1940_v58).length));
          let _rhs296 = (_1926_v44).isLessThanOrEqualTo((_dafny.ZERO).minus(_1926_v44));
          let _rhs297 = (_1926_v44).multipliedBy((new BigNumber((_1923_v41).length)).minus(_1926_v44));
          let _rhs298 = _dafny.Seq.Concat(_dafny.Seq.of(_1926_v44, _1926_v44), _1941_v59);
          let _rhs299 = (_this).f24;
          let _lhs218 = globalState;
          let _lhs219 = globalState;
          r2 = _rhs296;
          _lhs218.f7 = _rhs297;
          _1941_v59 = _rhs298;
          _lhs219.f13 = _rhs299;
          let _1942_v60;
          _1942_v60 = _module.D14.create_DC31((_1928_v46).f24);
          r2 = (_1942_v60).dtor_cf52;
        }
        let _1943_v61;
        let _nw320 = Array((new BigNumber(25)).toNumber()).fill(false);
        _1943_v61 = _nw320;
        let _1944_v62;
        _1944_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1943_v61,_1928_v46.f23);
        let _1945_v63;
        let _init58 = function (_1946_i3) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        };
        let _nw321 = Array((new BigNumber(17)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw321.length); _i0_58++) {
          _nw321[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _1945_v63 = _nw321;
        let _1947_v65;
        let _nw322 = new _module.C3();
        _nw322.__ctor(_1944_v62, _1945_v63, (_1927_v45)[_module.__default.safeIndex(_1926_v44, new BigNumber((_1927_v45).length))], (_dafny.Set.fromElements(_1926_v44)).equals(function () {
          let _coll91 = new _dafny.Set();
          for (const _compr_91 of _dafny.IntegerRange(new BigNumber(164), new BigNumber(617))) {
            let _1948_v64 = _compr_91;
            if (((new BigNumber(164)).isLessThanOrEqualTo(_1948_v64)) && ((_1948_v64).isLessThan(new BigNumber(617)))) {
              _coll91.add((_1948_v64).minus(_1926_v44));
            }
          }
          return _coll91;
        }()));
        _1947_v65 = _nw322;
      }
      let _1949_v66;
      _1949_v66 = new BigNumber(-734);
      let _1950_v67;
      _1950_v67 = _dafny.Seq.UnicodeFromString("yfjo");
      let _1951_v68;
      _1951_v68 = _dafny.Set.fromElements((_this).fm6(globalState), _1949_v66, _1949_v66);
      let _1952_v69;
      _1952_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1949_v66,_this.f23);
      let _1953_v70;
      let _nw323 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _1953_v70 = _nw323;
      let _1954_v71;
      _1954_v71 = _dafny.Seq.of(_1953_v70);
      let _1955_v72;
      _1955_v72 = _dafny.Seq.of(_1952_v69, _1952_v69, _1952_v69, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1954_v71).length),_this.f23));
      let _1956_v73;
      _1956_v73 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1949_v66), _1949_v66);
      let _1957_v74;
      let _nw324 = Array((new BigNumber(21)).toNumber());
      _nw324[(_dafny.ZERO).toNumber()] = _1949_v66;
      _nw324[(_dafny.ONE).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(2)).toNumber()] = (_1949_v66).multipliedBy(new BigNumber((_1950_v67).length));
      _nw324[(new BigNumber(3)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(4)).toNumber()] = (_1949_v66).minus(new BigNumber(434));
      _nw324[(new BigNumber(5)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(6)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(7)).toNumber()] = new BigNumber((_1950_v67).length);
      _nw324[(new BigNumber(8)).toNumber()] = new BigNumber((_1951_v68).length);
      _nw324[(new BigNumber(9)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(10)).toNumber()] = (_1949_v66).minus(_1949_v66);
      _nw324[(new BigNumber(11)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(12)).toNumber()] = (_1949_v66).minus(new BigNumber((_1955_v72).length));
      _nw324[(new BigNumber(13)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(14)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(15)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(16)).toNumber()] = (_1949_v66).plus(_1949_v66);
      _nw324[(new BigNumber(17)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1956_v73).cardinality()), _1949_v66);
      _nw324[(new BigNumber(19)).toNumber()] = _1949_v66;
      _nw324[(new BigNumber(20)).toNumber()] = new BigNumber(139);
      _1957_v74 = _nw324;
      let _1958_v75;
      _1958_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1949_v66,new BigNumber(904));
      let _1959_v76;
      _1959_v76 = _1958_v75;
      let _1960_v77;
      let _init59 = function (_1961_i4) {
        return _this.f23;
      };
      let _nw325 = Array((new BigNumber(23)).toNumber());
      for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw325.length); _i0_59++) {
        _nw325[_i0_59] = _init59(new BigNumber(_i0_59));
      }
      _1960_v77 = _nw325;
      let _1962_v78;
      let _nw326 = new _module.C4();
      _nw326.__ctor(new BigNumber(872), _1959_v76, _1960_v77, _this.f23, !((_this).f24));
      _1962_v78 = _nw326;
      let _1963_v79;
      _1963_v79 = _dafny.Seq.of(_1962_v78, _1962_v78, _1962_v78, _1962_v78, _1962_v78);
      let _1964_v80;
      _1964_v80 = _dafny.Seq.of((_1963_v79)[_module.__default.safeIndex(_1949_v66, new BigNumber((_1963_v79).length))], _1962_v78);
      let _1965_v81;
      _1965_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1964_v80);
      let _index260 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1957_v74).length));
      (_1957_v74)[_index260] = new BigNumber(((((_1965_v81).contains((_1962_v78).f24)) ? ((_1965_v81).get((_1962_v78).f24)) : (_1964_v80))).length);
      let _1966_v82;
      _1966_v82 = _dafny.Map.Empty.slice().updateUnsafe((_1962_v78).f24,new BigNumber((_dafny.Seq.of((_this).f24, (_this).f24)).length));
      let _index261 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1957_v74).length));
      (_1957_v74)[_index261] = _module.__default.safeDivisionInt(_1949_v66, (((_1966_v82).contains((_1962_v78).f24)) ? ((_1966_v82).get((_1962_v78).f24)) : (new BigNumber(-751))));
      r1 = (_1957_v74)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_1957_v74).length))];
      let _1967_v83;
      _1967_v83 = _dafny.Map.Empty.slice().updateUnsafe((_1962_v78).f25,((_1957_v74)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_1957_v74).length))]).minus((_1957_v74)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_1957_v74).length))]));
      _1967_v83 = (_1967_v83).update(_1960_v77, _1949_v66);
      let _rhs300 = _1962_v78.f23;
      let _rhs301 = _module.__default.fm34(_module.__default.safeModuloInt(new BigNumber((_1951_v68).length), _1949_v66), (_1962_v78).f24, globalState);
      let _lhs220 = _1962_v78;
      let _lhs221 = _this;
      _lhs220.f23 = _rhs300;
      _lhs221.f23 = _rhs301;
      let _1968_v84;
      let _init60 = ((_1969_v68) => function (_1970_i5) {
        return _1969_v68;
      })(_1951_v68);
      let _nw327 = Array((new BigNumber(18)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw327.length); _i0_60++) {
        _nw327[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _1968_v84 = _nw327;
      let _index262 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_1968_v84).length));
      (_1968_v84)[_index262] = (_module.D5.create_DC12(_1951_v68)).dtor_cf19;
      let _index263 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_1968_v84).length));
      (_1968_v84)[_index263] = _1951_v68;
      r0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1950_v67, _dafny.Seq.UnicodeFromString("ebhil")), _1950_v67);
      r1 = _1949_v66;
      r2 = (_module.D15.create_DC34(true)).dtor_cf55;
      return [r0, r1, r2];
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let _1971_v0;
      _1971_v0 = _dafny.Seq.UnicodeFromString("bnuau");
      let _hi12 = _module.__default.safeDivisionInt(p3, (_dafny.ZERO).minus(new BigNumber((_1971_v0).length)));
      for (let _1972_i0 = p3; _1972_i0.isLessThan(_hi12); _1972_i0 = _1972_i0.plus(_dafny.ONE)) {
        if (false) {
          _1971_v0 = _1971_v0;
          (globalState).f13 = p2;
          (globalState).f7 = ((p2) ? (p3) : ((((_this).f24) ? (_1972_i0) : (_1972_i0))));
          let _1973_v1;
          let _init61 = function (_1974_i1) {
            return (_1974_i1).plus(new BigNumber(930));
          };
          let _nw328 = Array((new BigNumber(16)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw328.length); _i0_61++) {
            _nw328[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _1973_v1 = _nw328;
          let _index264 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1973_v1).length));
          (_1973_v1)[_index264] = _1972_i0;
          let _index265 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1973_v1).length));
          (_1973_v1)[_index265] = _module.__default.safeDivisionInt(p3, p3);
          let _1975_v2;
          _1975_v2 = _dafny.Seq.of(p1);
          let _1976_v3;
          let _nw329 = new _module.C1();
          _nw329.__ctor((_this).fm4(new BigNumber(475), _1975_v2, globalState), p2);
          _1976_v3 = _nw329;
        } else {
          let _rhs302 = (_this).f24;
          let _rhs303 = !(p2) || (((p2) ? (p1) : (true)));
          let _lhs222 = globalState;
          let _lhs223 = globalState;
          _lhs222.f13 = _rhs302;
          _lhs223.f13 = _rhs303;
          let _1977_v4;
          _1977_v4 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1972_i0), p3);
          _1977_v4 = _1977_v4;
          let _1978_v5;
          _1978_v5 = _dafny.MultiSet.fromElements(_1972_i0, _module.__default.fm0((_this).f24, p3, globalState));
          let _1979_v6;
          _1979_v6 = _dafny.Seq.of(p2, (_this).f24);
          let _1980_v7;
          _1980_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1971_v0,(_this).f24);
          let _1981_v8;
          _1981_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1978_v5).IsDisjointFrom((_1978_v5).update(new BigNumber((_1979_v6).length), _module.__default.abs(p3))),new BigNumber((_1980_v7).length));
          _1981_v8 = _1981_v8;
          (globalState).f13 = false;
          let _1982_v9;
          let _init62 = ((_1983_v6) => function (_1984_i2) {
            return (_1984_i2).multipliedBy(new BigNumber((_1983_v6).length));
          })(_1979_v6);
          let _nw330 = Array((new BigNumber(10)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw330.length); _i0_62++) {
            _nw330[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _1982_v9 = _nw330;
          let _index266 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1982_v9).length));
          (_1982_v9)[_index266] = p3;
          let _index267 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1982_v9).length));
          (_1982_v9)[_index267] = _module.__default.safeModuloInt(_1972_i0, p3);
        }
        let _1985_v10;
        _1985_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(445)), function (_1986_i3) {
          return new BigNumber(635);
        })).length));
        let _1987_v11;
        _1987_v11 = _dafny.MultiSet.fromElements(_1972_i0, (_dafny.ZERO).minus((((_1985_v10).contains(p2)) ? ((_1985_v10).get(p2)) : (_1972_i0))), _1972_i0);
        let _1988_v12;
        _1988_v12 = _dafny.MultiSet.fromElements(_1971_v0);
        (globalState).f13 = !((_1987_v11).IsSubsetOf(_1987_v11)) || ((_dafny.MultiSet.fromElements(_1971_v0)).IsSubsetOf(_1988_v12));
        let _1989_v13;
        let _nw331 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _1989_v13 = _nw331;
        (globalState).f22 = _1989_v13;
        let _1990_v14;
        let _nw332 = Array((new BigNumber(25)).toNumber()).fill([]);
        _1990_v14 = _nw332;
        let _index268 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_1990_v14).length));
        (_1990_v14)[_index268] = _1989_v13;
        let _index269 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_1990_v14).length));
        (_1990_v14)[_index269] = ((p1) ? (_1989_v13) : (_1989_v13));
      }
      let _1991_v15;
      let _init63 = ((_1992_p1) => function (_1993_i4) {
        return _1992_p1;
      })(p1);
      let _nw333 = Array((new BigNumber(18)).toNumber());
      for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw333.length); _i0_63++) {
        _nw333[_i0_63] = _init63(new BigNumber(_i0_63));
      }
      _1991_v15 = _nw333;
      let _index270 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
      (_1991_v15)[_index270] = p2;
      let _index271 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
      (_1991_v15)[_index271] = (_this).f24;
      let _1994_v16;
      let _nw334 = Array((new BigNumber(5)).toNumber());
      _nw334[(_dafny.ZERO).toNumber()] = _this.f23;
      _nw334[(_dafny.ONE).toNumber()] = (_1971_v0)[_module.__default.safeIndex(p3, new BigNumber((_1971_v0).length))];
      _nw334[(new BigNumber(2)).toNumber()] = _this.f23;
      _nw334[(new BigNumber(3)).toNumber()] = _module.__default.fm34(p3, (_this).f24, globalState);
      _nw334[(new BigNumber(4)).toNumber()] = _this.f23;
      _1994_v16 = _nw334;
      let _source28 = _module.D12.create_DC24(_1994_v16);
      if (_source28.is_DC25) {
        let _1995___mcc_h0 = (_source28).cf41;
        let _1996___mcc_h1 = (_source28).cf42;
        let _1997_cf42 = _1996___mcc_h1;
        let _1998_cf41 = _1995___mcc_h0;
        (globalState).f13 = p1;
        r0 = _1998_cf41;
        let _1999_v17;
        _1999_v17 = _dafny.Set.fromElements(p1);
        if ((_dafny.Set.fromElements(_1998_cf41)).equals(_1999_v17)) {
          let _2000_v18;
          let _nw335 = Array((new BigNumber(5)).toNumber()).fill(_module.D18.Default());
          _2000_v18 = _nw335;
          let _2001_v19;
          let _nw336 = new _module.C0();
          _nw336.__ctor(_dafny.Map.Empty.slice().updateUnsafe(p2,_1971_v0), _1998_cf41, _this.f23, _1998_cf41);
          _2001_v19 = _nw336;
          let _2002_v20;
          _2002_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,_2001_v19);
          let _2003_v21;
          _2003_v21 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2001_v19);
          let _2004_v22;
          _2004_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2001_v19.f23,_2001_v19);
          let _2005_v23;
          let _nw337 = Array((new BigNumber(13)).toNumber());
          _nw337[(_dafny.ZERO).toNumber()] = _2001_v19;
          _nw337[(_dafny.ONE).toNumber()] = (((_2002_v20).contains((_this).f24)) ? ((_2002_v20).get((_this).f24)) : (_2001_v19));
          _nw337[(new BigNumber(2)).toNumber()] = (((_2003_v21).contains(p3)) ? ((_2003_v21).get(p3)) : (_2001_v19));
          _nw337[(new BigNumber(3)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(4)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(5)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(6)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(7)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(8)).toNumber()] = (((_2004_v22).contains(_2001_v19.f23)) ? ((_2004_v22).get(_2001_v19.f23)) : (_2001_v19));
          _nw337[(new BigNumber(9)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(10)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(11)).toNumber()] = _2001_v19;
          _nw337[(new BigNumber(12)).toNumber()] = _2001_v19;
          _2005_v23 = _nw337;
          let _pat_let_tv33 = _2005_v23;
          let _index272 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2000_v18).length));
          (_2000_v18)[_index272] = function (_pat_let49_0) {
            return function (_2006_dt__update__tmp_h0) {
              return function (_pat_let50_0) {
                return function (_2007_dt__update_hcf67_h0) {
                  return _module.D18.create_DC41(_2007_dt__update_hcf67_h0);
                }(_pat_let50_0);
              }(_pat_let_tv33);
            }(_pat_let49_0);
          }(_module.D18.create_DC41(_2005_v23));
          let _2008_v24;
          _2008_v24 = _module.D18.create_DC41(_2005_v23);
          let _pat_let_tv34 = _2005_v23;
          let _index273 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2000_v18).length));
          (_2000_v18)[_index273] = function (_pat_let51_0) {
            return function (_2009_dt__update__tmp_h1) {
              return function (_pat_let52_0) {
                return function (_2010_dt__update_hcf67_h1) {
                  return _module.D18.create_DC41(_2010_dt__update_hcf67_h1);
                }(_pat_let52_0);
              }(_pat_let_tv34);
            }(_pat_let51_0);
          }(_2008_v24);
          let _2011_v25;
          _2011_v25 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1998_cf41);
          let _2012_v26;
          _2012_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2011_v25).length),p1);
          let _2013_v27;
          let _nw338 = new _module.C8();
          _nw338.__ctor((_2011_v25).equals(_2011_v25));
          _2013_v27 = _nw338;
          _1971_v0 = ((!(_1998_cf41)) ? (_1971_v0) : (_1971_v0));
          (globalState).f13 = (_this).f24;
          let _2014_v28;
          _2014_v28 = _dafny.MultiSet.fromElements((_2001_v19).f24, false, (p3).isLessThan(p3), (_2001_v19).f24);
          _2014_v28 = _dafny.MultiSet.fromElements((function () {
            let _coll92 = new _dafny.Set();
            for (const _compr_92 of _dafny.IntegerRange(new BigNumber(767), new BigNumber(828))) {
              let _2015_v29 = _compr_92;
              if (((new BigNumber(767)).isLessThanOrEqualTo(_2015_v29)) && ((_2015_v29).isLessThan(new BigNumber(828)))) {
                _coll92.add(_module.__default.safeDivisionInt(_2015_v29, p3));
              }
            }
            return _coll92;
          }()).IsDisjointFrom(function () {
            let _coll93 = new _dafny.Set();
            for (const _compr_93 of _dafny.IntegerRange(new BigNumber(529), new BigNumber(499))) {
              let _2016_v30 = _compr_93;
              if (((new BigNumber(529)).isLessThanOrEqualTo(_2016_v30)) && ((_2016_v30).isLessThan(new BigNumber(499)))) {
                _coll93.add((_2016_v30).minus(p3));
              }
            }
            return _coll93;
          }()));
        } else {
          let _2017_v31;
          _2017_v31 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
          let _2018_v32;
          let _nw339 = new _module.C4();
          _nw339.__ctor(p3, _2017_v31, _1994_v16, _this.f23, !(p1));
          _2018_v32 = _nw339;
          let _2019_v33;
          let _nw340 = Array((new BigNumber(10)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2019_v33 = _nw340;
          let _2020_v34;
          let _nw341 = new _module.C7();
          _nw341.__ctor((_this).f24, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))], _1994_v16, _1997_cf42, p1, p1);
          _2020_v34 = _nw341;
          let _2021_v35;
          _2021_v35 = _dafny.MultiSet.fromElements(_2020_v34);
          let _index274 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_2019_v33).length));
          (_2019_v33)[_index274] = _2021_v35;
          let _2022_v36;
          _2022_v36 = _dafny.Seq.of(p3);
          let _2023_v37;
          _2023_v37 = _dafny.MultiSet.fromElements(new BigNumber(905), new BigNumber((_dafny.Seq.of(_2022_v36)).length), new BigNumber((_1971_v0).length));
          let _2024_v38;
          _2024_v38 = _module.D15.create_DC34(p1);
          let _2025_v39;
          _2025_v39 = _dafny.Seq.of((_2020_v34).f31, p2);
          let _index275 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_2019_v33).length));
          let _rhs304 = _1999_v17;
          let _rhs305 = ((_dafny.MultiSet.fromElements(_2020_v34)).Intersect(_2021_v35)).Difference(_2021_v35);
          let _rhs306 = _dafny.MultiSet.fromElements((_2018_v32).f33, ((_dafny.ZERO).minus(p3)).plus((_2018_v32).f33), (new BigNumber((_dafny.MultiSet.fromElements(_2024_v38)).cardinality())).minus(new BigNumber((_2025_v39).length)));
          let _lhs224 = _2019_v33;
          let _lhs225 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_2019_v33).length));
          _1999_v17 = _rhs304;
          _lhs224[_lhs225] = _rhs305;
          _2023_v37 = _rhs306;
          let _2026_v40;
          _2026_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1998_cf41,(_2020_v34).f31);
          (globalState).f13 = (((_2026_v40).contains(_1998_cf41)) ? ((_2026_v40).get(_1998_cf41)) : (_1998_cf41));
          let _2027_v41;
          _2027_v41 = _dafny.Set.fromElements((_2018_v32).f33);
          (globalState).f13 = ((_dafny.Set.fromElements((_2018_v32).f33)).Union(_2027_v41)).IsSubsetOf(_2027_v41);
          let _2028_v42;
          let _init64 = ((_2029_p2) => function (_2030_i5) {
            return _dafny.Map.Empty.slice().updateUnsafe(_2029_p2,_this.f23);
          })(p2);
          let _nw342 = Array((new BigNumber(20)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw342.length); _i0_64++) {
            _nw342[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2028_v42 = _nw342;
          let _2031_v43;
          _2031_v43 = _module.D14.create_DC30(_2028_v42);
          let _2032_v44;
          _2032_v44 = _dafny.Map.Empty.slice().updateUnsafe((_2018_v32).f33,_2031_v43);
          _2032_v44 = _dafny.Map.Empty.slice().updateUnsafe((p3).plus(p3),_module.D14.create_DC30(_2028_v42));
        }
        let _2033_v45;
        _2033_v45 = _module.D12.create_DC25((p3).isLessThanOrEqualTo(p3), _1997_cf42);
        _2033_v45 = _2033_v45;
      } else if (_source28.is_DC24) {
        let _2034___mcc_h2 = (_source28).cf40;
        let _2035_cf40 = _2034___mcc_h2;
        let _2036_v46;
        _2036_v46 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm2(globalState));
        let _2037_v47;
        _2037_v47 = _dafny.Set.fromElements((((_2036_v46).contains(p3)) ? ((_2036_v46).get(p3)) : ((_this).f24)), (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
        let _2038_v48;
        _2038_v48 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_2037_v47).Union(_2037_v47));
        _2038_v48 = (_2038_v48).update(p3, _dafny.Set.fromElements(p1, false));
        if (p1) {
          let _2039_v49;
          _2039_v49 = _module.D12.create_DC25((_this).f24, _this.f23);
          let _index276 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_1994_v16).length));
          (_1994_v16)[_index276] = (_2039_v49).dtor_cf42;
          let _index277 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_1994_v16).length));
          let _rhs307 = _this.f23;
          let _rhs308 = !((_this).f24) || (!((_this).f24));
          let _rhs309 = p2;
          let _lhs226 = _1994_v16;
          let _lhs227 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_1994_v16).length));
          let _lhs228 = globalState;
          let _lhs229 = globalState;
          _lhs226[_lhs227] = _rhs307;
          _lhs228.f13 = _rhs308;
          _lhs229.f13 = _rhs309;
          (globalState).f7 = _module.__default.safeModuloInt((p3).plus((_dafny.ZERO).minus(p3)), (((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]) ? (p3) : (p3)));
          (globalState).f13 = p2;
          let _2040_v50;
          let _nw343 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _2040_v50 = _nw343;
          let _index278 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2040_v50).length));
          (_2040_v50)[_index278] = p3;
          let _index279 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2040_v50).length));
          (_2040_v50)[_index279] = (_dafny.ZERO).minus(p3);
          (globalState).f7 = p3;
        } else {
          let _2041_v51;
          _2041_v51 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
          (globalState).f7 = _module.__default.fm0((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))], (((_2041_v51).contains(p2)) ? ((_2041_v51).get(p2)) : (new BigNumber((_2036_v46).length))), globalState);
          let _2042_v52;
          let _nw344 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _2042_v52 = _nw344;
          let _2043_v53;
          _2043_v53 = _module.D15.create_DC34((_this).f24);
          let _2044_v54;
          _2044_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_2043_v53);
          let _index280 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_2042_v52).length));
          (_2042_v52)[_index280] = _2044_v54;
          let _index281 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_2042_v52).length));
          (_2042_v52)[_index281] = _2044_v54;
          (globalState).f7 = new BigNumber(263);
          _1971_v0 = _dafny.Seq.Concat(_1971_v0, _dafny.Seq.update(_1971_v0, _module.__default.safeIndex(p3, new BigNumber((_1971_v0).length)), new _dafny.CodePoint('o'.codePointAt(0))));
          let _2045_v55;
          _2045_v55 = _dafny.Seq.of(_dafny.Set.fromElements(!((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]), p2, p1, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]));
          (globalState).f13 = !(((p3).plus(p3)).isLessThan(new BigNumber((_dafny.Seq.Concat(_2045_v55, _dafny.Seq.of(_module.__default.fm41(globalState)))).length)));
        }
        (_this).f23 = _this.f23;
        (_this).f23 = _this.f23;
      } else {
        let _2046___mcc_h3 = (_source28).cf43;
        let _2047_cf43 = _2046___mcc_h3;
        _1971_v0 = _1971_v0;
        let _2048_v56;
        _2048_v56 = _dafny.Seq.of((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
        let _2049_v57;
        _2049_v57 = _dafny.MultiSet.fromElements((_this).fm11(new BigNumber(-498), p1, globalState), (_this).f24);
        let _2050_v58;
        _2050_v58 = _dafny.Seq.of(p3);
        let _2051_v59;
        _2051_v59 = _dafny.Set.fromElements(_2050_v58);
        let _2052_v60;
        let _2053_v61;
        let _2054_v62;
        let _out55;
        let _out56;
        let _out57;
        let _outcollector17 = _module.__default.m0(new BigNumber((_2048_v56).length), _2049_v57, !(!((_2051_v59).IsProperSubsetOf(_2051_v59))), (p3).multipliedBy((_dafny.ZERO).minus(p3)), globalState);
        _out55 = _outcollector17[0];
        _out56 = _outcollector17[1];
        _out57 = _outcollector17[2];
        _2052_v60 = _out55;
        _2053_v61 = _out56;
        _2054_v62 = _out57;
        let _2055_v63;
        _2055_v63 = _dafny.Set.fromElements(p1, p2, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
        let _2056_v64;
        let _nw345 = new _module.C8();
        _nw345.__ctor(!(_2055_v63).equals(_2055_v63));
        _2056_v64 = _nw345;
        _2056_v64 = _2056_v64;
        let _2057_v65;
        _2057_v65 = _dafny.Set.fromElements((_2052_v60).multipliedBy((_2050_v58)[_module.__default.safeIndex(new BigNumber((_2055_v63).length), new BigNumber((_2050_v58).length))]), _2054_v62, p3, _2052_v60, _2052_v60);
        _2057_v65 = _dafny.Set.fromElements(p3);
      }
      let _hi13 = p3;
      for (let _2058_i6 = (_dafny.ZERO).minus(p3); _2058_i6.isLessThan(_hi13); _2058_i6 = _2058_i6.plus(_dafny.ONE)) {
        if ((((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]) ? (p1) : (p1))) {
          let _2059_v66;
          _2059_v66 = _dafny.Seq.of(p1, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))], p1);
          (globalState).f7 = (new BigNumber((_2059_v66).length)).multipliedBy(p3);
          let _2060_v67;
          _2060_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1991_v15,p3);
          let _2061_v68;
          _2061_v68 = _dafny.MultiSet.fromElements(p2);
          _2060_v67 = (_2060_v67).update(_1991_v15, _module.__default.safeModuloInt(_2058_i6, new BigNumber((_2061_v68).cardinality())));
          let _index282 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
          (_1991_v15)[_index282] = (_2061_v68).equals((_2061_v68).Union(_2061_v68));
          (globalState).f13 = !(p1) || ((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
          let _2062_v69;
          _2062_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1991_v15,_this.f23);
          let _2063_v70;
          let _nw346 = new _module.C3();
          _nw346.__ctor(_2062_v69, _1994_v16, _this.f23, !((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]) || (p2));
          _2063_v70 = _nw346;
        } else {
          let _2064_v71;
          let _nw347 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _2064_v71 = _nw347;
          let _index283 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_2064_v71).length));
          (_2064_v71)[_index283] = p3;
          let _index284 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_2064_v71).length));
          (_2064_v71)[_index284] = new BigNumber((_dafny.Seq.Concat(_1971_v0, _1971_v0)).length);
          let _2065_v72;
          _2065_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1991_v15,_this.f23);
          let _2066_v73;
          let _nw348 = new _module.C3();
          _nw348.__ctor(_2065_v72, _1994_v16, _this.f23, p1);
          _2066_v73 = _nw348;
          let _index285 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_2064_v71).length));
          (_2064_v71)[_index285] = (_2064_v71)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_2064_v71).length))];
          let _2067_v74;
          _2067_v74 = _dafny.MultiSet.fromElements((_2064_v71)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_2064_v71).length))]);
          let _2068_v75;
          let _nw349 = new _module.C1();
          _nw349.__ctor(true, (_dafny.MultiSet.fromElements(new BigNumber(-913))).IsProperSubsetOf(_2067_v74));
          _2068_v75 = _nw349;
          let _2069_v76;
          _2069_v76 = _module.D5.create_DC12(_dafny.Set.fromElements(p3));
          let _2070_v77;
          _2070_v77 = _dafny.MultiSet.fromElements(_2069_v76);
          _2070_v77 = _2070_v77;
        }
        let _index286 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1994_v16).length));
        (_1994_v16)[_index286] = _this.f23;
        let _index287 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1994_v16).length));
        let _rhs310 = _module.__default.safeDivisionInt(_2058_i6, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(341)), ((_2071_v0, _2072_p3) => function (_2073_i7) {
          return (_2071_v0)[_module.__default.safeIndex(_2072_p3, new BigNumber((_2071_v0).length))];
        })(_1971_v0, p3)), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(341)), ((_2074_v0, _2075_p3) => function (_2076_i7) {
          return (_2074_v0)[_module.__default.safeIndex(_2075_p3, new BigNumber((_2074_v0).length))];
        })(_1971_v0, p3))).length)), _this.f23), _1971_v0)).length));
        let _rhs311 = new _dafny.CodePoint('n'.codePointAt(0));
        let _rhs312 = _1991_v15;
        let _rhs313 = _this.f23;
        let _lhs230 = globalState;
        let _lhs231 = _1994_v16;
        let _lhs232 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1994_v16).length));
        let _lhs233 = _this;
        _lhs230.f7 = _rhs310;
        _lhs231[_lhs232] = _rhs311;
        _1991_v15 = _rhs312;
        _lhs233.f23 = _rhs313;
        r0 = !(true);
        let _2077_v78;
        _2077_v78 = _dafny.Seq.of(p1);
        let _2078_v79;
        _2078_v79 = _dafny.Set.fromElements((_this).f24, !((_this).f24), (_this).fm4(new BigNumber(-465), _2077_v78, globalState));
        (globalState).f7 = new BigNumber((_2078_v79).length);
      }
      if (p2) {
        let _2079_v80;
        _2079_v80 = _dafny.MultiSet.fromElements(!(p1));
        let _2080_v81;
        let _2081_v82;
        let _2082_v83;
        let _out58;
        let _out59;
        let _out60;
        let _outcollector18 = _module.__default.m0((_this).fm6(globalState), (_module.__default.fm3(p2, globalState)).Union(_2079_v80), p2, p3, globalState);
        _out58 = _outcollector18[0];
        _out59 = _outcollector18[1];
        _out60 = _outcollector18[2];
        _2080_v81 = _out58;
        _2081_v82 = _out59;
        _2082_v83 = _out60;
        let _2083_v84;
        _2083_v84 = _module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(true,_this.f23));
        let _pat_let_tv35 = _2081_v82;
        let _2084_v85;
        _2084_v85 = _dafny.MultiSet.fromElements(_2083_v84, (((_this).f24) ? (function (_pat_let53_0) {
          return function (_2085_dt__update__tmp_h2) {
            return function (_pat_let54_0) {
              return function (_2086_dt__update_hcf44_h0) {
                return _module.D13.create_DC27(_2086_dt__update_hcf44_h0);
              }(_pat_let54_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv35,_this.f23));
          }(_pat_let53_0);
        }(_2083_v84)) : (_2083_v84)), _2083_v84);
        _2084_v85 = _module.__default.fm50((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))], _2080_v81, globalState);
        let _2087_v86;
        _2087_v86 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(599)), ((_2088_v0) => function (_2089_i8) {
          return new BigNumber((_2088_v0).length);
        })(_1971_v0))).length));
        let _2090_v87;
        _2090_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_2087_v86, _module.__default.safeIndex(_2080_v81, new BigNumber((_2087_v86).length)), _2082_v83)).length),_this.f23);
        (_this).f23 = ((_2081_v82) ? (_this.f23) : ((((_2090_v87).contains(new BigNumber(-794))) ? ((_2090_v87).get(new BigNumber(-794))) : (new _dafny.CodePoint('b'.codePointAt(0))))));
        _2080_v81 = p3;
        let _2091_v88;
        _2091_v88 = _dafny.Set.fromElements(_2081_v82);
        (globalState).f7 = (_dafny.ZERO).minus(new BigNumber(((_2091_v88).Union(_2091_v88)).length));
      } else {
        (globalState).f7 = new BigNumber(650);
        let _index288 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
        (_1991_v15)[_index288] = (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))];
        let _2092_v89;
        _2092_v89 = _module.D2.create_DC7(!(false), p2);
        let _2093_v90;
        _2093_v90 = _dafny.Seq.of(_2092_v89);
        _2093_v90 = _dafny.Seq.Concat(_2093_v90, _dafny.Seq.Concat(_2093_v90, _dafny.Seq.of(_2092_v89)));
        let _2094_v91;
        _2094_v91 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("aynidalob"), _1971_v0, _1971_v0, _1971_v0, _1971_v0);
        let _2095_v92;
        _2095_v92 = _dafny.Set.fromElements(new BigNumber(271), (((_2094_v91).contains(_1971_v0)) ? ((_2094_v91).get(_1971_v0)) : (p3)));
        _2095_v92 = _2095_v92;
        let _2096_v93;
        _2096_v93 = _dafny.Set.fromElements((_this).f24);
        let _2097_v94;
        _2097_v94 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_2096_v93).length));
        _2097_v94 = _2097_v94;
      }
      if (((true) ? ((_this).f24) : ((_this).f24))) {
        let _2098_v95;
        _2098_v95 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))], new BigNumber((_1971_v0).length), globalState),p3);
        let _2099_v96;
        _2099_v96 = _dafny.Set.fromElements(_2098_v95);
        let _2100_v97;
        _2100_v97 = _dafny.Map.Empty.slice().updateUnsafe((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))],(_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
        let _2101_v98;
        _2101_v98 = _dafny.Seq.of((_this).f24, (_this).f24);
        let _2102_v99;
        let _nw350 = new _module.C8();
        _nw350.__ctor(p2);
        _2102_v99 = _nw350;
        let _2103_v100;
        _2103_v100 = _dafny.Map.Empty.slice().updateUnsafe((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))],_2102_v99);
        _2099_v96 = _module.__default.fm51(_dafny.Map.Empty.slice().updateUnsafe(_2100_v97,_2101_v98), p3, p3, ((p1) ? ((_2101_v98)[_module.__default.safeIndex(new BigNumber((_2103_v100).length), new BigNumber((_2101_v98).length))]) : (p1)), globalState);
        let _2104_v101;
        _2104_v101 = _dafny.MultiSet.fromElements(p2, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))], (_2101_v98)[_module.__default.safeIndex(p3, new BigNumber((_2101_v98).length))]);
        if (!((_dafny.MultiSet.fromElements(p3)).update(new BigNumber((_2104_v101).cardinality()), _module.__default.abs(p3))).contains(_module.__default.fm0(!((_this).f24), p3, globalState))) {
          let _2105_v102;
          let _2106_v103;
          let _2107_v104;
          let _out61;
          let _out62;
          let _out63;
          let _outcollector19 = _module.__default.m0(p3, (_2104_v101).Union(_2104_v101), p1, p3, globalState);
          _out61 = _outcollector19[0];
          _out62 = _outcollector19[1];
          _out63 = _outcollector19[2];
          _2105_v102 = _out61;
          _2106_v103 = _out62;
          _2107_v104 = _out63;
          (globalState).f7 = (_dafny.ZERO).minus(new BigNumber((_1971_v0).length));
          let _2108_v105;
          let _nw351 = Array((new BigNumber(23)).toNumber());
          _2108_v105 = _nw351;
          let _2109_v106;
          let _nw352 = new _module.C6();
          _nw352.__ctor(_this.f23, !(_module.__default.fm2(globalState)));
          _2109_v106 = _nw352;
          let _index289 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_2108_v105).length));
          (_2108_v105)[_index289] = _2109_v106;
          let _index290 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_2108_v105).length));
          (_2108_v105)[_index290] = _2109_v106;
          let _2110_v107;
          _2110_v107 = _module.D13.create_DC28(_dafny.Seq.IsProperPrefixOf(_1971_v0, _1971_v0), (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1971_v0, _module.__default.safeIndex(new BigNumber((_2101_v98).length), new BigNumber((_1971_v0).length)), _this.f23), _module.__default.safeIndex((_dafny.ZERO).minus(_2105_v102), new BigNumber((_dafny.Seq.update(_1971_v0, _module.__default.safeIndex(new BigNumber((_2101_v98).length), new BigNumber((_1971_v0).length)), _this.f23)).length)), _this.f23)).length)).multipliedBy(p3)), (_2107_v104).plus(new BigNumber(920)), (_2107_v104).multipliedBy(new BigNumber(592)));
          let _rhs314 = _2110_v107;
          let _rhs315 = !(_2106_v103) || ((_this).f24);
          let _lhs234 = globalState;
          _2110_v107 = _rhs314;
          _lhs234.f13 = _rhs315;
          _2107_v104 = (_module.D15.create_DC36(_module.__default.fm2(globalState), new BigNumber((_2100_v97).length))).dtor_cf60;
        } else {
          let _2111_v108;
          _2111_v108 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new _dafny.CodePoint('t'.codePointAt(0)));
          let _2112_v109;
          let _nw353 = new _module.C5();
          _nw353.__ctor(_1991_v15);
          _2112_v109 = _nw353;
          let _2113_v110;
          _2113_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2112_v109,(_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
          let _2114_v111;
          _2114_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2113_v110).length),p1);
          _2111_v108 = _module.__default.fm52(new BigNumber((_2114_v111).length), !_dafny.areEqual(_dafny.Seq.of(_this.f23), _1971_v0), p3, p3, globalState);
          let _2115_v112;
          let _nw354 = new _module.C6();
          _nw354.__ctor(_this.f23, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
          _2115_v112 = _nw354;
          let _2116_v113;
          _2116_v113 = _module.D5.create_DC13((_dafny.ZERO).minus(_module.__default.safeModuloInt(p3, (((_2104_v101).contains((_this).f24)) ? ((_2104_v101).get((_this).f24)) : (new BigNumber((_module.__default.fm1((_2115_v112).f24, globalState)).length))))), (_2115_v112).f24, (_2115_v112).f24);
          let _2117_v114;
          _2117_v114 = _dafny.Map.Empty.slice().updateUnsafe((_2115_v112).f24,new BigNumber(522));
          let _2118_v115;
          _2118_v115 = _dafny.Map.Empty.slice().updateUnsafe(_1991_v15,(_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
          let _2119_v116;
          _2119_v116 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber((_2118_v115).length));
          let _2120_v117;
          let _nw355 = new _module.C6();
          _nw355.__ctor(_2115_v112.f23, p2);
          _2120_v117 = _nw355;
          let _2121_v118;
          _2121_v118 = _dafny.Map.Empty.slice().updateUnsafe(_2120_v117,p3);
          let _2122_v119;
          let _nw356 = Array((new BigNumber(29)).toNumber());
          _nw356[(_dafny.ZERO).toNumber()] = p3;
          _nw356[(_dafny.ONE).toNumber()] = new BigNumber(565);
          _nw356[(new BigNumber(2)).toNumber()] = p3;
          _nw356[(new BigNumber(3)).toNumber()] = p3;
          _nw356[(new BigNumber(4)).toNumber()] = p3;
          _nw356[(new BigNumber(5)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,_this.f23)).Merge(_2111_v108)).length);
          _nw356[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(216), new BigNumber((_2117_v114).length));
          _nw356[(new BigNumber(7)).toNumber()] = p3;
          _nw356[(new BigNumber(8)).toNumber()] = p3;
          _nw356[(new BigNumber(9)).toNumber()] = (p3).multipliedBy(new BigNumber((_1971_v0).length));
          _nw356[(new BigNumber(10)).toNumber()] = p3;
          _nw356[(new BigNumber(11)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_this).fm6(globalState))).Merge(_2119_v116)).length);
          _nw356[(new BigNumber(12)).toNumber()] = p3;
          _nw356[(new BigNumber(13)).toNumber()] = p3;
          _nw356[(new BigNumber(14)).toNumber()] = (((_2104_v101).contains((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))])) ? ((_2104_v101).get((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))])) : (new BigNumber(218)));
          _nw356[(new BigNumber(15)).toNumber()] = (p3).minus(p3);
          _nw356[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2121_v118).length), p3);
          _nw356[(new BigNumber(17)).toNumber()] = p3;
          _nw356[(new BigNumber(18)).toNumber()] = p3;
          _nw356[(new BigNumber(19)).toNumber()] = new BigNumber(463);
          _nw356[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt(p3, p3);
          _nw356[(new BigNumber(21)).toNumber()] = new BigNumber((_1971_v0).length);
          _nw356[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(287), p3);
          _nw356[(new BigNumber(23)).toNumber()] = new BigNumber(-90);
          _nw356[(new BigNumber(24)).toNumber()] = new BigNumber((_2101_v98).length);
          _nw356[(new BigNumber(25)).toNumber()] = p3;
          _nw356[(new BigNumber(26)).toNumber()] = p3;
          _nw356[(new BigNumber(27)).toNumber()] = new BigNumber(420);
          _nw356[(new BigNumber(28)).toNumber()] = p3;
          _2122_v119 = _nw356;
          let _index291 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2122_v119).length));
          (_2122_v119)[_index291] = new BigNumber((_dafny.Seq.of(p3, p3, p3)).length);
          let _2123_v120;
          _2123_v120 = _dafny.MultiSet.fromElements(p3);
          let _2124_v121;
          _2124_v121 = _module.D11.create_DC23((_this).f24, new BigNumber((_2123_v120).cardinality()), _2100_v97, _module.D5.create_DC14((_2115_v112).f24));
          let _pat_let_tv36 = p3;
          let _index292 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2122_v119).length));
          let _rhs316 = _2115_v112;
          let _rhs317 = !((((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]) ? ((_this).f24) : (p1)));
          let _rhs318 = function (_pat_let55_0) {
            return function (_2125_dt__update__tmp_h3) {
              return function (_pat_let56_0) {
                return function (_2126_dt__update_hcf20_h0) {
                  return _module.D5.create_DC13(_2126_dt__update_hcf20_h0, (_2125_dt__update__tmp_h3).dtor_cf21, (_2125_dt__update__tmp_h3).dtor_cf22);
                }(_pat_let56_0);
              }(_pat_let_tv36);
            }(_pat_let55_0);
          }(_2116_v113);
          let _rhs319 = _module.__default.fm0(p2, new BigNumber((_2104_v101).cardinality()), globalState);
          let _rhs320 = (((_2124_v121).dtor_cf36) ? (new BigNumber(390)) : ((((_2123_v120).contains(p3)) ? ((_2123_v120).get(p3)) : (new BigNumber((_2117_v114).length)))));
          let _lhs235 = globalState;
          let _lhs236 = _2122_v119;
          let _lhs237 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2122_v119).length));
          let _lhs238 = globalState;
          _2115_v112 = _rhs316;
          _lhs235.f13 = _rhs317;
          _2116_v113 = _rhs318;
          _lhs236[_lhs237] = _rhs319;
          _lhs238.f7 = _rhs320;
          (globalState).f13 = p1;
          let _index293 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2122_v119).length));
          (_2122_v119)[_index293] = (_2115_v112).fm6(globalState);
          let _index294 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
          (_1991_v15)[_index294] = !(!(true) || (p1)) || (!((_2115_v112).f24) || ((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]));
        }
        if ((_this).f24) {
          let _2127_v122;
          let _nw357 = new _module.C6();
          _nw357.__ctor(_this.f23, p1);
          _2127_v122 = _nw357;
          let _2128_v123;
          _2128_v123 = _module.D5.create_DC14(p1);
          let _2129_v124;
          _2129_v124 = _module.D11.create_DC23((_this).f24, new BigNumber((_1971_v0).length), _2100_v97, _2128_v123);
          _1991_v15 = (((_2129_v124).dtor_cf36) ? (_1991_v15) : (_1991_v15));
          let _2130_v125;
          let _init65 = ((_2131_v0) => function (_2132_i9) {
            return _2131_v0;
          })(_1971_v0);
          let _nw358 = Array((new BigNumber(18)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw358.length); _i0_65++) {
            _nw358[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2130_v125 = _nw358;
          _2130_v125 = _2130_v125;
          _2098_v95 = (_2098_v95).update(p3, (((_2104_v101).contains(false)) ? ((_2104_v101).get(false)) : (p3)));
          let _2133_v126;
          _2133_v126 = _dafny.Set.fromElements(((_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]) === ((_this).f24));
          _2133_v126 = (_2133_v126).Union(_2133_v126);
        } else {
          r0 = p1;
          (globalState).f13 = p1;
          (globalState).f13 = !_dafny.Seq.contains(_1971_v0, _this.f23);
          r0 = p2;
          let _2134_v127;
          _2134_v127 = _dafny.Set.fromElements(_dafny.Seq.Concat(_1971_v0, _dafny.Seq.UnicodeFromString("ykpfixig")));
          _2134_v127 = _module.__default.fm53(globalState);
        }
        _2100_v97 = (_2100_v97).update(p2, (_1991_v15)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length))]);
        (globalState).f7 = p3;
      } else {
        let _2135_v128;
        _2135_v128 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        let _2136_v129;
        let _nw359 = new _module.C4();
        _nw359.__ctor(p3, _2135_v128, _1994_v16, _this.f23, (_this).f24);
        _2136_v129 = _nw359;
        let _2137_v130;
        _2137_v130 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2136_v129);
        let _2138_v131;
        _2138_v131 = _dafny.Seq.of(p2);
        _2137_v130 = (_2137_v130).update((new BigNumber((_2138_v131).length)).plus(p3), _2136_v129);
        let _2139_v132;
        _2139_v132 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState));
        let _index295 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
        (_1991_v15)[_index295] = (!(p1) || (true)) || ((_2139_v132).IsSubsetOf(_2139_v132));
        let _index296 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1991_v15).length));
        (_1991_v15)[_index296] = !((_this).f24);
        (globalState).f7 = new BigNumber(57);
        (globalState).f13 = p1;
      }
      let _pat_let_tv37 = p1;
      let _pat_let_tv38 = p1;
      let _pat_let_tv39 = _1991_v15;
      let _pat_let_tv40 = _1991_v15;
      let _pat_let_tv41 = p3;
      let _pat_let_tv42 = p3;
      let _pat_let_tv43 = p2;
      let _pat_let_tv44 = _1971_v0;
      r0 = function (_source29) {
        if (_source29.is_DC13) {
          let _2140___mcc_h4 = (_source29).cf20;
          let _2141___mcc_h5 = (_source29).cf21;
          let _2142___mcc_h6 = (_source29).cf22;
          let _2143_cf22 = _2142___mcc_h6;
          let _2144_cf21 = _2141___mcc_h5;
          let _2145_cf20 = _2140___mcc_h4;
          return (_pat_let_tv37) === (_pat_let_tv38);
        } else if (_source29.is_DC14) {
          let _2146___mcc_h7 = (_source29).cf23;
          let _2147_cf23 = _2146___mcc_h7;
          return (_pat_let_tv40)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_pat_let_tv39).length))];
        } else {
          let _2148___mcc_h8 = (_source29).cf19;
          let _2149_cf19 = _2148___mcc_h8;
          return (_dafny.MultiSet.fromElements(_pat_let_tv41, _pat_let_tv42)).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_pat_let_tv43)).length), new BigNumber((_pat_let_tv44).length), new BigNumber(622)));
        }
      }(_module.D5.create_DC13(new BigNumber((_1971_v0).length), true, p1));
      return r0;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      if (p3) {
        let _2150_v0;
        _2150_v0 = _dafny.MultiSet.fromElements(true);
        let _2151_v1;
        _2151_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2150_v0,_this.f23);
        _2151_v1 = _2151_v1;
        let _2152_v2;
        _2152_v2 = new BigNumber(641);
        let _2153_v3;
        _2153_v3 = _dafny.Seq.of(true);
        let _2154_v4;
        _2154_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2152_v2,_2152_v2);
        let _2155_v5;
        _2155_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2152_v2);
        let _2156_v6;
        let _nw360 = Array((new BigNumber(10)).toNumber());
        _nw360[(_dafny.ZERO).toNumber()] = _2154_v4;
        _nw360[(_dafny.ONE).toNumber()] = _2154_v4;
        _nw360[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),_2152_v2);
        _nw360[(new BigNumber(3)).toNumber()] = _2154_v4;
        _nw360[(new BigNumber(4)).toNumber()] = _2154_v4;
        _nw360[(new BigNumber(5)).toNumber()] = _2154_v4;
        _nw360[(new BigNumber(6)).toNumber()] = _2154_v4;
        _nw360[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2150_v0).cardinality()),(_dafny.ZERO).minus(new BigNumber((_2155_v5).length)));
        _nw360[(new BigNumber(8)).toNumber()] = _2154_v4;
        _nw360[(new BigNumber(9)).toNumber()] = _2154_v4;
        _2156_v6 = _nw360;
        let _2157_v7;
        _2157_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2156_v6,(_this).fm6(globalState));
        let _2158_v8;
        _2158_v8 = _module.D9.create_DC20(p2, p0, _2152_v2);
        let _2159_v9;
        _2159_v9 = _dafny.Seq.of(new BigNumber(-691));
        let _2160_v10;
        _2160_v10 = _dafny.MultiSet.fromElements(new BigNumber((_2159_v9).length));
        let _2161_v11;
        let _nw361 = Array((new BigNumber(22)).toNumber());
        _nw361[(_dafny.ZERO).toNumber()] = _2152_v2;
        _nw361[(_dafny.ONE).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(2)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(3)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(4)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(5)).toNumber()] = (_2152_v2).plus(_2152_v2);
        _nw361[(new BigNumber(6)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(7)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(8)).toNumber()] = (((_2150_v0).contains((_this).f24)) ? ((_2150_v0).get((_this).f24)) : (_2152_v2));
        _nw361[(new BigNumber(9)).toNumber()] = new BigNumber(648);
        _nw361[(new BigNumber(10)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(11)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(12)).toNumber()] = (((_2153_v3)[_module.__default.safeIndex(_2152_v2, new BigNumber((_2153_v3).length))]) ? (_2152_v2) : ((((_2157_v7).contains(_2156_v6)) ? ((_2157_v7).get(_2156_v6)) : (_2152_v2))));
        _nw361[(new BigNumber(13)).toNumber()] = (_2152_v2).plus(new BigNumber(((_dafny.MultiSet.fromElements(_2158_v8)).update(_module.D9.create_DC20(p2, (_this).f24, new BigNumber((_2160_v10).cardinality())), _module.__default.abs(_2152_v2))).cardinality()));
        _nw361[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(_2152_v2, new BigNumber((p2).length));
        _nw361[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_2152_v2);
        _nw361[(new BigNumber(16)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(17)).toNumber()] = (_2152_v2).plus(new BigNumber((_dafny.Seq.UnicodeFromString("xhna")).length));
        _nw361[(new BigNumber(18)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(19)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(20)).toNumber()] = _2152_v2;
        _nw361[(new BigNumber(21)).toNumber()] = (((_this).f24) ? (_2152_v2) : (_2152_v2));
        _2161_v11 = _nw361;
        let _index297 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_2161_v11).length));
        (_2161_v11)[_index297] = _2152_v2;
        let _index298 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_2161_v11).length));
        (_2161_v11)[_index298] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((p1).dtor_cf3), _2152_v2);
        let _2162_v12;
        _2162_v12 = _dafny.Set.fromElements(_2150_v0, _2150_v0, _2150_v0, (_dafny.MultiSet.fromElements((_this).f24)).Intersect(_dafny.MultiSet.FromArray(_2153_v3)));
        _2162_v12 = _dafny.Set.fromElements(_module.__default.fm3(p3, globalState), (_2150_v0).update((_this).f24, _module.__default.abs((_2161_v11)[_module.__default.safeIndex(new BigNumber(583), new BigNumber((_2161_v11).length))])));
        let _2163_v13;
        _2163_v13 = _dafny.Seq.UnicodeFromString("d");
        let _2164_v14;
        _2164_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2152_v2,_this.f23);
        _2163_v13 = _dafny.Seq.update(_module.__default.fm1((((_this).f24) ? (p3) : ((_this).f24)), globalState), _module.__default.safeIndex((_2161_v11)[_module.__default.safeIndex(new BigNumber(583), new BigNumber((_2161_v11).length))], new BigNumber((_module.__default.fm1((((_this).f24) ? (p3) : ((_this).f24)), globalState)).length)), (((_2164_v14).contains(new BigNumber(702))) ? ((_2164_v14).get(new BigNumber(702))) : (_this.f23)));
        let _2165_v15;
        let _nw362 = Array((new BigNumber(9)).toNumber()).fill(false);
        _2165_v15 = _nw362;
        let _index299 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_2165_v15).length));
        (_2165_v15)[_index299] = !((_this).f24) || (p3);
        let _2166_v16;
        _2166_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,false);
        let _2167_v17;
        _2167_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,_2163_v13);
        let _2168_v18;
        let _nw363 = new _module.C0();
        _nw363.__ctor(_2167_v17, true, _this.f23, (_this).f24);
        _2168_v18 = _nw363;
        let _2169_v19;
        _2169_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2168_v18,(_this).f24);
        let _2170_v20;
        let _nw364 = new _module.C1();
        _nw364.__ctor(p3, (((_2166_v16).contains((((_2169_v19).contains(_2168_v18)) ? ((_2169_v19).get(_2168_v18)) : (true)))) ? ((_2166_v16).get((((_2169_v19).contains(_2168_v18)) ? ((_2169_v19).get(_2168_v18)) : (true)))) : (true)));
        _2170_v20 = _nw364;
        let _2171_v21;
        _2171_v21 = _dafny.Seq.of(_2170_v20);
        let _2172_v22;
        _2172_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2163_v13,_2171_v21);
        let _index300 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_2165_v15).length));
        let _rhs321 = new BigNumber(((_2172_v22).update(_2163_v13, _2171_v21)).length);
        let _rhs322 = ((_dafny.ZERO).minus(new BigNumber(-745))).minus(new BigNumber((_2163_v13).length));
        let _rhs323 = (_2168_v18).f24;
        let _rhs324 = (_2170_v20).f27;
        let _lhs239 = globalState;
        let _lhs240 = globalState;
        let _lhs241 = _2165_v15;
        let _lhs242 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_2165_v15).length));
        _lhs239.f7 = _rhs321;
        _lhs240.f7 = _rhs322;
        r0 = _rhs323;
        _lhs241[_lhs242] = _rhs324;
      } else {
        let _2173_v23;
        _2173_v23 = new BigNumber(861);
        let _2174_v24;
        _2174_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p2);
        let _2175_v25;
        _2175_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2173_v23,_2174_v24);
        let _2176_v26;
        let _nw365 = new _module.C0();
        _nw365.__ctor((_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC7(!((_this).f24), (_this).f24)).dtor_cf14,p2)).Merge((((_2175_v25).contains(_2173_v23)) ? ((_2175_v25).get(_2173_v23)) : (_2174_v24))), p3, _this.f23, p0);
        _2176_v26 = _nw365;
        _2176_v26 = _2176_v26;
        let _2177_v27;
        _2177_v27 = _dafny.Set.fromElements((_this).f24);
        let _2178_v28;
        _2178_v28 = _dafny.Set.fromElements((_dafny.ZERO).minus(_2173_v23), _2173_v23);
        let _2179_v29;
        _2179_v29 = _dafny.Set.fromElements(_2178_v28, _2178_v28);
        (globalState).f7 = _module.__default.fm0((_2179_v29).IsSubsetOf(_module.__default.fm54(_2177_v27, globalState)), (_2173_v23).minus(_2173_v23), globalState);
        let _2180_v30;
        let _nw366 = Array((new BigNumber(23)).toNumber());
        _nw366[(_dafny.ZERO).toNumber()] = p0;
        _nw366[(_dafny.ONE).toNumber()] = (_this).fm11(_2173_v23, (_2176_v26).f24, globalState);
        _nw366[(new BigNumber(2)).toNumber()] = p3;
        _nw366[(new BigNumber(3)).toNumber()] = (_2176_v26).f24;
        _nw366[(new BigNumber(4)).toNumber()] = p3;
        _nw366[(new BigNumber(5)).toNumber()] = p0;
        _nw366[(new BigNumber(6)).toNumber()] = p3;
        _nw366[(new BigNumber(7)).toNumber()] = (_this).f24;
        _nw366[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw366[(new BigNumber(9)).toNumber()] = p3;
        _nw366[(new BigNumber(10)).toNumber()] = (_2176_v26).f24;
        _nw366[(new BigNumber(11)).toNumber()] = (_this).f24;
        _nw366[(new BigNumber(12)).toNumber()] = p0;
        _nw366[(new BigNumber(13)).toNumber()] = p3;
        _nw366[(new BigNumber(14)).toNumber()] = (_this).f24;
        _nw366[(new BigNumber(15)).toNumber()] = p0;
        _nw366[(new BigNumber(16)).toNumber()] = p3;
        _nw366[(new BigNumber(17)).toNumber()] = false;
        _nw366[(new BigNumber(18)).toNumber()] = (_this).f24;
        _nw366[(new BigNumber(19)).toNumber()] = (_this).f24;
        _nw366[(new BigNumber(20)).toNumber()] = (_2176_v26).f24;
        _nw366[(new BigNumber(21)).toNumber()] = p3;
        _nw366[(new BigNumber(22)).toNumber()] = p0;
        _2180_v30 = _nw366;
        let _2181_v31;
        _2181_v31 = _dafny.Seq.of(_2180_v30, _2180_v30);
        let _2182_v32;
        let _nw367 = Array((_dafny.ONE).toNumber());
        _nw367[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2181_v31, _2181_v31);
        _2182_v32 = _nw367;
        let _index301 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_2182_v32).length));
        (_2182_v32)[_index301] = _2181_v31;
        let _index302 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_2182_v32).length));
        (_2182_v32)[_index302] = _dafny.Seq.of(_2180_v30);
        (globalState).f7 = (_2173_v23).multipliedBy(_2173_v23);
        let _index303 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_2180_v30).length));
        (_2180_v30)[_index303] = (_this).f24;
        let _2183_v33;
        _2183_v33 = _dafny.Seq.of(p3, (_this).f24, p0, p3, (_2176_v26).f24);
        let _2184_v34;
        _2184_v34 = _dafny.Seq.of(_2183_v33);
        let _2185_v35;
        _2185_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2173_v23,_2173_v23);
        let _2186_v36;
        _2186_v36 = _dafny.Seq.of(_2173_v23, (_this).fm6(globalState));
        let _2187_v37;
        let _nw368 = new _module.C1();
        _nw368.__ctor(p3, p0);
        _2187_v37 = _nw368;
        let _2188_v38;
        _2188_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2187_v37,_2173_v23);
        let _2189_v39;
        _2189_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2173_v23,_2188_v38);
        let _2190_v40;
        _2190_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2173_v23,false);
        let _2191_v41;
        _2191_v41 = _module.D19.create_DC45(_dafny.Map.Empty.slice().updateUnsafe(_2187_v37,_2173_v23));
        let _index304 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_2180_v30).length));
        let _rhs325 = (_2185_v35).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_this).f24), _2183_v33), _module.__default.safeIndex(new BigNumber(912), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f24), _2183_v33)).length)), (_this).f24), _module.__default.safeIndex(new BigNumber((_2184_v34).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_this).f24), _2183_v33), _module.__default.safeIndex(new BigNumber(912), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f24), _2183_v33)).length)), (_this).f24)).length)), p0)).length)));
        let _rhs326 = ((!(p0) || ((_this).f24)) ? (new BigNumber((_2186_v36).length)) : (_2173_v23));
        let _rhs327 = (new BigNumber(((_2189_v39).update(new BigNumber((_2190_v40).length), (_2191_v41).dtor_cf74)).length)).isLessThan(new BigNumber(-491));
        let _rhs328 = (((_this).f24) ? ((_2187_v37).f39) : ((_this).fm4(_2173_v23, _2183_v33, globalState)));
        let _lhs243 = globalState;
        let _lhs244 = _2180_v30;
        let _lhs245 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_2180_v30).length));
        r0 = _rhs325;
        _2173_v23 = _rhs326;
        _lhs243.f13 = _rhs327;
        _lhs244[_lhs245] = _rhs328;
      }
      let _2192_v42;
      _2192_v42 = new BigNumber(458);
      let _2193_v43;
      _2193_v43 = _dafny.Map.Empty.slice().updateUnsafe(_2192_v42,p0);
      if (!((((_2193_v43).contains(_2192_v42)) ? ((_2193_v43).get(_2192_v42)) : (p3)))) {
        let _2194_v44;
        _2194_v44 = _dafny.Seq.of((_2192_v42).multipliedBy(_2192_v42));
        let _rhs329 = _this.f23;
        let _rhs330 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2194_v44, _2194_v44), _2194_v44);
        let _lhs246 = _this;
        _lhs246.f23 = _rhs329;
        _2194_v44 = _rhs330;
        let _2195_v45;
        let _nw369 = new _module.C8();
        _nw369.__ctor(p3);
        _2195_v45 = _nw369;
        let _2196_v46;
        _2196_v46 = _dafny.Seq.of(p0, (_this).f24);
        let _2197_v47;
        _2197_v47 = _dafny.Set.fromElements((_this).f24, !((_2196_v46)[_module.__default.safeIndex(_2192_v42, new BigNumber((_2196_v46).length))]), (_this).f24, p3, (_this).f24);
        let _2198_v48;
        _2198_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2197_v47,_dafny.MultiSet.FromArray(_2194_v44));
        let _2199_v49;
        _2199_v49 = _module.D15.create_DC35(_2192_v42, _2192_v42, new BigNumber((_2197_v47).length));
        let _2200_v50;
        _2200_v50 = _dafny.Seq.of(_2199_v49, _2199_v49);
        _2198_v48 = (_2198_v48).Merge(((p3) ? (_2198_v48) : (_dafny.Map.Empty.slice().updateUnsafe(_2197_v47,_module.__default.fm42(p3, new BigNumber((_2200_v50).length), globalState)))));
        let _2201_v51;
        _2201_v51 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        let _2202_v52;
        _2202_v52 = _dafny.Seq.of(_2194_v44);
        let _2203_v53;
        _2203_v53 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(399));
        let _2204_v54;
        _2204_v54 = _dafny.Set.fromElements(_2192_v42, new BigNumber((_2202_v52).length), new BigNumber((_2203_v53).length), _2192_v42, new BigNumber((p2).length));
        let _2205_v55;
        let _nw370 = new _module.C0();
        _nw370.__ctor(((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,p2)).update(_module.__default.fm2(globalState), _dafny.Seq.UnicodeFromString("b"))).Merge(_2201_v51), (_2204_v54).IsSubsetOf(_dafny.Set.fromElements((_this).fm6(globalState))), _this.f23, true);
        _2205_v55 = _nw370;
        let _2206_v56;
        _2206_v56 = _module.D15.create_DC34((_this).f24);
        let _source30 = _2206_v56;
        if (_source30.is_DC34) {
          let _2207___mcc_h0 = (_source30).cf55;
          let _2208_cf55 = _2207___mcc_h0;
          (globalState).f7 = _2192_v42;
          (globalState).f7 = _2192_v42;
          let _2209_v57;
          let _nw371 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _2209_v57 = _nw371;
          let _2210_v58;
          _2210_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2192_v42,new BigNumber(-824));
          let _2211_v59;
          _2211_v59 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f23);
          let _2212_v60;
          _2212_v60 = _dafny.Seq.of(_2211_v59);
          let _index305 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_2209_v57).length));
          (_2209_v57)[_index305] = _dafny.Seq.Concat(_module.__default.fm55(_2208_cf55, _2210_v58, _dafny.Seq.Create(_module.__default.abs(new BigNumber(404)), ((_2213_v42) => function (_2214_i0) {
            return _2213_v42;
          })(_2192_v42)), globalState), _2212_v60);
          let _2215_v61;
          let _nw372 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2215_v61 = _nw372;
          let _index306 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_2215_v61).length));
          (_2215_v61)[_index306] = p2;
          let _2216_v62;
          let _init66 = ((_2217_v42) => function (_2218_i1) {
            return (_2218_i1).multipliedBy(_2217_v42);
          })(_2192_v42);
          let _nw373 = Array((new BigNumber(6)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw373.length); _i0_66++) {
            _nw373[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2216_v62 = _nw373;
          let _index307 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_2216_v62).length));
          (_2216_v62)[_index307] = ((_this).fm6(globalState)).plus(new BigNumber((_2194_v44).length));
          let _2219_v63;
          _2219_v63 = _dafny.MultiSet.fromElements(p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(887)), function (_2220_i2) {
            return _this.f23;
          }), p2, p2, p2);
          let _2221_v64;
          _2221_v64 = _2219_v63;
          let _2222_v65;
          _2222_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_dafny.MultiSet.fromElements(p2));
          let _index308 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_2209_v57).length));
          let _index309 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_2215_v61).length));
          let _index310 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_2216_v62).length));
          let _rhs331 = _2212_v60;
          let _rhs332 = !(p3) || (((_2221_v64)).IsProperSubsetOf((((_2222_v65).contains(new _dafny.CodePoint('j'.codePointAt(0)))) ? ((_2222_v65).get(new _dafny.CodePoint('j'.codePointAt(0)))) : (_2219_v63))));
          let _rhs333 = _dafny.Seq.Concat(_dafny.Seq.Concat(p2, p2), _dafny.Seq.Concat(p2, p2));
          let _rhs334 = _2192_v42;
          let _lhs247 = _2209_v57;
          let _lhs248 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_2209_v57).length));
          let _lhs249 = _2215_v61;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_2215_v61).length));
          let _lhs251 = _2216_v62;
          let _lhs252 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_2216_v62).length));
          _lhs247[_lhs248] = _rhs331;
          _2208_cf55 = _rhs332;
          _lhs249[_lhs250] = _rhs333;
          _lhs251[_lhs252] = _rhs334;
          let _2223_v66;
          _2223_v66 = _dafny.Map.Empty.slice().updateUnsafe((_2205_v55).fm26((_2216_v62)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_2216_v62).length))], _2196_v46, _2204_v54, (_2205_v55).f38, globalState),(_2216_v62)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_2216_v62).length))]);
          let _2224_v67;
          _2224_v67 = _dafny.MultiSet.fromElements(new BigNumber((_2223_v66).length), _2192_v42);
          (globalState).f13 = (_2224_v67).IsSubsetOf(_2224_v67);
        } else if (_source30.is_DC35) {
          let _2225___mcc_h1 = (_source30).cf56;
          let _2226___mcc_h2 = (_source30).cf57;
          let _2227___mcc_h3 = (_source30).cf58;
          let _2228_cf58 = _2227___mcc_h3;
          let _2229_cf57 = _2226___mcc_h2;
          let _2230_cf56 = _2225___mcc_h1;
          let _2231_v68;
          _2231_v68 = _dafny.Seq.UnicodeFromString("uiky");
          let _2232_v69;
          _2232_v69 = _module.D19.create_DC46(_2197_v47, p2);
          let _2233_v70;
          _2233_v70 = _dafny.Seq.of(_2231_v68, _dafny.Seq.Concat(_module.__default.fm1((_this).f24, globalState), p2), p2, (_2232_v69).dtor_cf76);
          _2231_v68 = (_2233_v70)[_module.__default.safeIndex(_2229_cf57, new BigNumber((_2233_v70).length))];
          let _2234_v71;
          _2234_v71 = _dafny.MultiSet.fromElements((_2205_v55).f38, (_this).f24);
          let _2235_v72;
          _2235_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2229_cf57,_2234_v71);
          _2193_v43 = (_2193_v43).update(_module.__default.safeDivisionInt(_2192_v42, new BigNumber(((((_2235_v72).contains(_2228_cf58)) ? ((_2235_v72).get(_2228_cf58)) : (_dafny.MultiSet.fromElements((_2205_v55).f38, p3)))).cardinality())), !((_2205_v55).f38) || (p3));
          let _2236_v73;
          _2236_v73 = _dafny.Map.Empty.slice().updateUnsafe(_2192_v42,_dafny.Seq.Create(_module.__default.abs(new BigNumber(972)), ((_2237_cf57) => function (_2238_i3) {
            return _2237_cf57;
          })(_2229_cf57)));
          _2236_v73 = (_2236_v73).update((new BigNumber((_2231_v68).length)).minus(_2192_v42), _2194_v44);
          let _2239_v74;
          _2239_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2194_v44,_this.f23);
          let _2240_v75;
          _2240_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2231_v68,_2239_v74);
          _2240_v75 = (_2240_v75).update(p2, _2239_v74);
        } else if (_source30.is_DC36) {
          let _2241___mcc_h4 = (_source30).cf59;
          let _2242___mcc_h5 = (_source30).cf60;
          let _2243_cf60 = _2242___mcc_h5;
          let _2244_cf59 = _2241___mcc_h4;
          let _2245_v76;
          let _init67 = function (_2246_i4) {
            return _this.f23;
          };
          let _nw374 = Array((new BigNumber(28)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw374.length); _i0_67++) {
            _nw374[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2245_v76 = _nw374;
          let _index311 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2245_v76).length));
          (_2245_v76)[_index311] = _this.f23;
          let _2247_v77;
          let _nw375 = Array((new BigNumber(8)).toNumber()).fill(false);
          _2247_v77 = _nw375;
          let _2248_v78;
          _2248_v78 = _dafny.MultiSet.fromElements(_2247_v77, _2247_v77);
          let _index312 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2245_v76).length));
          let _rhs335 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm1(_2244_cf59, globalState), _dafny.Seq.Concat(p2, _dafny.Seq.UnicodeFromString("qyif")))).length);
          let _rhs336 = _this.f23;
          let _rhs337 = (((_2248_v78).contains(_2247_v77)) ? ((_2248_v78).get(_2247_v77)) : (_2243_cf60));
          let _rhs338 = p3;
          let _lhs253 = globalState;
          let _lhs254 = _2245_v76;
          let _lhs255 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2245_v76).length));
          let _lhs256 = globalState;
          let _lhs257 = globalState;
          _lhs253.f7 = _rhs335;
          _lhs254[_lhs255] = _rhs336;
          _lhs256.f7 = _rhs337;
          _lhs257.f13 = _rhs338;
          let _index313 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2245_v76).length));
          (_2245_v76)[_index313] = _this.f23;
          _2243_cf60 = _2192_v42;
          let _2249_v79;
          let _nw376 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _2249_v79 = _nw376;
          let _2250_v80;
          _2250_v80 = _dafny.Set.fromElements(_2249_v79, _2249_v79);
          (globalState).f7 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_2250_v80).length), _2192_v42))));
        } else {
          let _2251___mcc_h6 = (_source30).cf54;
          let _2252_cf54 = _2251___mcc_h6;
          let _2253_v81;
          let _nw377 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _2253_v81 = _nw377;
          let _index314 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2253_v81).length));
          (_2253_v81)[_index314] = new BigNumber(-378);
          let _2254_v82;
          _2254_v82 = _dafny.MultiSet.fromElements(_2203_v53);
          let _index315 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2253_v81).length));
          (_2253_v81)[_index315] = new BigNumber(((_2254_v82).Intersect((_2254_v82).Difference((_2254_v82).update(_2203_v53, _module.__default.abs(_2192_v42))))).cardinality());
          let _2255_v83;
          let _nw378 = new _module.C0();
          _nw378.__ctor(_2201_v51, (_2205_v55).f38, _this.f23, (_2205_v55).f38);
          _2255_v83 = _nw378;
          let _2256_v84;
          _2256_v84 = _2255_v83;
          _2256_v84 = _2256_v84;
          let _2257_v85;
          let _nw379 = Array((new BigNumber(13)).toNumber());
          _2257_v85 = _nw379;
          let _2258_v86;
          _2258_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2257_v85,p2);
          _2258_v86 = (_2258_v86).update(_2257_v85, p2);
          (globalState).f13 = p0;
        }
      } else {
        (globalState).f7 = (_2192_v42).multipliedBy((_2192_v42).plus(_2192_v42));
        let _2259_v87;
        _2259_v87 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f23);
        (globalState).f13 = !(!(!(_2259_v87).equals(_2259_v87)));
        let _2260_v88;
        _2260_v88 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        let _2261_v89;
        let _nw380 = new _module.C0();
        _nw380.__ctor(_2260_v88, p0, _this.f23, p0);
        _2261_v89 = _nw380;
        _2192_v42 = ((_this).fm6(globalState)).minus(_2192_v42);
        let _2262_v90;
        _2262_v90 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),(_this).f24)).length)).minus(_2192_v42)),_module.__default.fm56(p3, _2192_v42, globalState));
        _2262_v90 = (function () {
          let _coll94 = new _dafny.Map();
          for (const _compr_94 of (_dafny.MultiSet.fromElements(_2192_v42)).Elements) {
            let _2263_v91 = _compr_94;
            if ((_dafny.MultiSet.fromElements(_2192_v42)).contains(_2263_v91)) {
              _coll94.push([_module.__default.safeModuloInt(_2263_v91, _2192_v42),_2193_v43]);
            }
          }
          return _coll94;
        }()).Merge(_2262_v90);
      }
      let _2264_v92;
      _2264_v92 = _dafny.Seq.of(new BigNumber((p2).length));
      let _2265_v93;
      _2265_v93 = _dafny.Set.fromElements((_this).fm6(globalState), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_2192_v42), _2264_v92)).length));
      _2265_v93 = _2265_v93;
      let _2266_v94;
      _2266_v94 = _dafny.Set.fromElements(true);
      let _2267_v95;
      _2267_v95 = _dafny.MultiSet.fromElements(_2192_v42, (_dafny.ZERO).minus(_2192_v42));
      let _2268_v96;
      _2268_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_2266_v94, _dafny.Set.fromElements(p0, p3, true))).length),new BigNumber((_2267_v95).cardinality()));
      _2268_v96 = (_2268_v96).Merge(_2268_v96);
      let _2269_v97;
      _2269_v97 = _module.D9.create_DC20(_dafny.Seq.UnicodeFromString("of"), (_this).f24, _2192_v42);
      let _2270_v98;
      _2270_v98 = _dafny.MultiSet.fromElements(_2269_v97);
      let _2271_v99;
      _2271_v99 = _module.D21.create_DC50(_2270_v98);
      _2270_v98 = (_2271_v99).dtor_cf80;
      let _2272_i5;
      _2272_i5 = _dafny.ZERO;
      L7: {
        while ((_this).f24) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2272_i5)) {
              break L7;
            }
            _2272_i5 = (_2272_i5).plus(_dafny.ONE);
            _2193_v43 = (_2193_v43).update((_dafny.ZERO).minus((_2192_v42).plus(_2192_v42)), (_2192_v42).isEqualTo(_2192_v42));
            let _2273_v100;
            _2273_v100 = _dafny.MultiSet.fromElements(_2265_v93);
            if ((_2273_v100).IsProperSubsetOf((((_this).f24) ? (_2273_v100) : (_2273_v100)))) {
              let _2274_v101;
              let _nw381 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _2274_v101 = _nw381;
              let _2275_v102;
              let _nw382 = Array((new BigNumber(7)).toNumber());
              _nw382[(_dafny.ZERO).toNumber()] = p0;
              _nw382[(_dafny.ONE).toNumber()] = _module.__default.fm2(globalState);
              _nw382[(new BigNumber(2)).toNumber()] = p0;
              _nw382[(new BigNumber(3)).toNumber()] = (_this).f24;
              _nw382[(new BigNumber(4)).toNumber()] = true;
              _nw382[(new BigNumber(5)).toNumber()] = p0;
              _nw382[(new BigNumber(6)).toNumber()] = (_this).f24;
              _2275_v102 = _nw382;
              let _2276_v103;
              let _nw383 = Array((new BigNumber(3)).toNumber());
              _nw383[(_dafny.ZERO).toNumber()] = _2275_v102;
              _nw383[(_dafny.ONE).toNumber()] = _2275_v102;
              _nw383[(new BigNumber(2)).toNumber()] = _2275_v102;
              _2276_v103 = _nw383;
              let _index316 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2276_v103).length));
              (_2276_v103)[_index316] = _2275_v102;
              let _index317 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2276_v103).length));
              let _rhs339 = _2274_v101;
              let _rhs340 = (_this).f24;
              let _rhs341 = _dafny.Seq.IsPrefixOf(p2, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), ((_2277_p2, _2278_v42) => function (_2279_i6) {
                return (_2277_p2)[_module.__default.safeIndex(_2278_v42, new BigNumber((_2277_p2).length))];
              })(p2, _2192_v42)), p2));
              let _rhs342 = _2192_v42;
              let _rhs343 = _2275_v102;
              let _lhs258 = globalState;
              let _lhs259 = globalState;
              let _lhs260 = _2276_v103;
              let _lhs261 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2276_v103).length));
              _2274_v101 = _rhs339;
              _lhs258.f13 = _rhs340;
              r0 = _rhs341;
              _lhs259.f7 = _rhs342;
              _lhs260[_lhs261] = _rhs343;
              let _2280_v104;
              let _init68 = function (_2281_i7) {
                return _this.f23;
              };
              let _nw384 = Array((new BigNumber(2)).toNumber());
              for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw384.length); _i0_68++) {
                _nw384[_i0_68] = _init68(new BigNumber(_i0_68));
              }
              _2280_v104 = _nw384;
              _2280_v104 = _2280_v104;
              let _2282_v105;
              _2282_v105 = _dafny.MultiSet.fromElements(p0);
              let _2283_v106;
              let _nw385 = new _module.C2();
              _nw385.__ctor(_2282_v105, _2280_v104, _this.f23, p3);
              _2283_v106 = _nw385;
              let _2284_v107;
              _2284_v107 = _module.D5.create_DC13(_2192_v42, p3, p3);
              let _rhs344 = (_2284_v107).dtor_cf20;
              let _rhs345 = _2192_v42;
              let _rhs346 = !_dafny.Seq.contains(p2, _this.f23);
              let _rhs347 = (_dafny.ZERO).minus(_2192_v42);
              let _lhs262 = globalState;
              let _lhs263 = globalState;
              let _lhs264 = globalState;
              _lhs262.f7 = _rhs344;
              _2192_v42 = _rhs345;
              _lhs263.f13 = _rhs346;
              _lhs264.f7 = _rhs347;
              r0 = !((_this).f24);
            } else {
              let _2285_v108;
              let _init69 = ((_2286_v95) => function (_2287_i8) {
                return _2286_v95;
              })(_2267_v95);
              let _nw386 = Array((new BigNumber(28)).toNumber());
              for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw386.length); _i0_69++) {
                _nw386[_i0_69] = _init69(new BigNumber(_i0_69));
              }
              _2285_v108 = _nw386;
              _2285_v108 = _2285_v108;
              (globalState).f13 = !(_2192_v42).isEqualTo(_2192_v42);
              let _2288_v109;
              _2288_v109 = _dafny.Seq.UnicodeFromString("cjiuamnmc");
              _2288_v109 = _2288_v109;
              let _2289_v110;
              _2289_v110 = _module.D0.create_DC1(new BigNumber((_2268_v96).length));
              let _2290_v111;
              _2290_v111 = _dafny.Seq.of(_module.D0.create_DC1(new BigNumber(38)), _2289_v110, _2289_v110);
              let _2291_v112;
              _2291_v112 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(339),_dafny.Set.fromElements(_2192_v42));
              let _2292_v113;
              _2292_v113 = _dafny.Seq.of(true, p3);
              let _2293_v114;
              let _nw387 = Array((new BigNumber(28)).toNumber());
              _nw387[(_dafny.ZERO).toNumber()] = p3;
              _nw387[(_dafny.ONE).toNumber()] = p3;
              _nw387[(new BigNumber(2)).toNumber()] = p3;
              _nw387[(new BigNumber(3)).toNumber()] = (_this).f24;
              _nw387[(new BigNumber(4)).toNumber()] = false;
              _nw387[(new BigNumber(5)).toNumber()] = p0;
              _nw387[(new BigNumber(6)).toNumber()] = p0;
              _nw387[(new BigNumber(7)).toNumber()] = p3;
              _nw387[(new BigNumber(8)).toNumber()] = (_this).fm4(_2192_v42, _2292_v113, globalState);
              _nw387[(new BigNumber(9)).toNumber()] = _module.__default.fm2(globalState);
              _nw387[(new BigNumber(10)).toNumber()] = p0;
              _nw387[(new BigNumber(11)).toNumber()] = false;
              _nw387[(new BigNumber(12)).toNumber()] = p0;
              _nw387[(new BigNumber(13)).toNumber()] = p3;
              _nw387[(new BigNumber(14)).toNumber()] = p0;
              _nw387[(new BigNumber(15)).toNumber()] = p3;
              _nw387[(new BigNumber(16)).toNumber()] = false;
              _nw387[(new BigNumber(17)).toNumber()] = p0;
              _nw387[(new BigNumber(18)).toNumber()] = p0;
              _nw387[(new BigNumber(19)).toNumber()] = !((_this).f24);
              _nw387[(new BigNumber(20)).toNumber()] = p0;
              _nw387[(new BigNumber(21)).toNumber()] = p0;
              _nw387[(new BigNumber(22)).toNumber()] = p3;
              _nw387[(new BigNumber(23)).toNumber()] = p0;
              _nw387[(new BigNumber(24)).toNumber()] = true;
              _nw387[(new BigNumber(25)).toNumber()] = p3;
              _nw387[(new BigNumber(26)).toNumber()] = p3;
              _nw387[(new BigNumber(27)).toNumber()] = p0;
              _2293_v114 = _nw387;
              let _2294_v115;
              let _nw388 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _2294_v115 = _nw388;
              let _2295_v116;
              let _nw389 = new _module.C3();
              _nw389.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_2293_v114,_this.f23), _2294_v115, new _dafny.CodePoint('v'.codePointAt(0)), false);
              _2295_v116 = _nw389;
              let _2296_v117;
              _2296_v117 = _dafny.Seq.of(_2295_v116, _2295_v116);
              let _2297_v118;
              let _out64;
              _out64 = (_this).m5(_2290_v111, (_2265_v93).IsDisjointFrom((((_2291_v112).contains(_2192_v42)) ? ((_2291_v112).get(_2192_v42)) : (_2265_v93))), !(false), new BigNumber((_2296_v117).length), globalState);
              _2297_v118 = _out64;
              (globalState).f7 = (_dafny.ZERO).minus(new BigNumber((_2288_v109).length));
            }
            let _2298_v119;
            _2298_v119 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _2299_v120;
            _2299_v120 = _dafny.Seq.of(p0);
            let _2300_v121;
            _2300_v121 = _dafny.Map.Empty.slice().updateUnsafe(_2299_v120,_dafny.Seq.UnicodeFromString("ri"));
            let _2301_v122;
            _2301_v122 = _dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC7(p3, p0)).dtor_cf13,(((_this).fm11(new BigNumber((_2298_v119).length), (_this).f24, globalState)) ? (_2300_v121) : (_2300_v121)));
            _2301_v122 = (_2301_v122).update(!_dafny.Seq.contains(p2, _this.f23), _2300_v121);
            let _2302_v123;
            let _nw390 = Array((new BigNumber(16)).toNumber()).fill(_module.D4.Default());
            _2302_v123 = _nw390;
            let _2303_v124;
            _2303_v124 = _module.D4.create_DC11(_2192_v42);
            let _index318 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_2302_v123).length));
            (_2302_v123)[_index318] = _2303_v124;
            let _index319 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_2302_v123).length));
            (_2302_v123)[_index319] = _module.D4.create_DC11(_2192_v42);
          }
        }
      }
      let _2304_v125;
      _2304_v125 = _dafny.Seq.of((_2265_v93).contains(_2192_v42));
      r0 = (_2304_v125)[_module.__default.safeIndex((_dafny.ZERO).minus((((_this).f24) ? (_2192_v42) : (_2192_v42))), new BigNumber((_2304_v125).length))];
      return r0;
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f24 = false;
      this._f25 = [];
      this._f27 = false;
      this.f29 = _dafny.ZERO;
      this._f28 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T3, _module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f28, f29, f23, f24, f25, f27) {
      let _this = this;
      (_this)._f28 = f28;
      (_this).f29 = f29;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this)._f27 = f27;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return _module.D0.create_DC3((((_this).f27) ? (_module.D0.create_DC0(_dafny.MultiSet.fromElements((_this).f28))) : (_module.D0.create_DC3(_module.D0.create_DC2(true, new BigNumber((_dafny.Seq.of(new BigNumber(997))).length), (_this).f28)))));
    };
    fm6(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this.f29).plus(_this.f29));
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return !((_this).f28);
    };
    fm7(globalState) {
      let _this = this;
      if ((new BigNumber((_dafny.Seq.of((_this).f27)).length)).isLessThanOrEqualTo(_this.f29)) {
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f29))).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(874)), function (_2305_i0) {
          return _this.f29;
        }));
      } else {
        return _dafny.Seq.of(_this.f29);
      }
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_dafny.Seq.UnicodeFromString("th"))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("jedxluog")))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,_dafny.Seq.UnicodeFromString("afbr")));
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f29)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bfg")).length))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f23)).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f29)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f29)))).length);
    };
    fm10(p0, p1, globalState) {
      let _this = this;
      return _this.f23;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D0.Default();
      let r2 = false;
      let _2306_v0;
      let _init70 = ((_2307_p0) => function (_2308_i0) {
        return _dafny.Seq.update(_dafny.Seq.UnicodeFromString("qccj"), _module.__default.safeIndex(_this.f29, new BigNumber((_dafny.Seq.UnicodeFromString("qccj")).length)), _2307_p0);
      })(p0);
      let _nw391 = Array((new BigNumber(13)).toNumber());
      for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw391.length); _i0_70++) {
        _nw391[_i0_70] = _init70(new BigNumber(_i0_70));
      }
      _2306_v0 = _nw391;
      let _2309_v1;
      _2309_v1 = _dafny.Seq.UnicodeFromString("nbcmfn");
      let _index320 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length));
      (_2306_v0)[_index320] = _2309_v1;
      let _index321 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length));
      (_2306_v0)[_index321] = _dafny.Seq.update(_2309_v1, _module.__default.safeIndex(_module.__default.safeDivisionInt(_this.f29, (_dafny.ZERO).minus(_this.f29)), new BigNumber((_2309_v1).length)), p0);
      let _2310_v2;
      let _nw392 = new _module.C1();
      _nw392.__ctor((!((_this).f24)) && (p1), (_this).f28);
      _2310_v2 = _nw392;
      let _2311_v3;
      _2311_v3 = _dafny.Seq.of(_dafny.Seq.Concat(_module.__default.fm1((_this).f28, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), function (_2312_i1) {
        return _this.f23;
      })), _2309_v1, _2309_v1, _dafny.Seq.Concat(_2309_v1, _2309_v1), (_2306_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length))]);
      _2311_v3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm21(_dafny.Set.fromElements((_this).f27), _dafny.MultiSet.FromArray((_this).fm7(globalState)), globalState), _2311_v3), _2311_v3);
      let _2313_v4;
      let _nw393 = Array((new BigNumber(2)).toNumber()).fill([]);
      _2313_v4 = _nw393;
      let _2314_v5;
      let _nw394 = new _module.C7();
      _nw394.__ctor((_2310_v2).f39, (_2310_v2).f39, (_this).f25, _this.f23, !((_this).f27), (_this).f24);
      _2314_v5 = _nw394;
      let _2315_v6;
      let _nw395 = Array((new BigNumber(14)).toNumber());
      _nw395[(_dafny.ZERO).toNumber()] = _2314_v5;
      _nw395[(_dafny.ONE).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(2)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(3)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(4)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(5)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(6)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(7)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(8)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(9)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(10)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(11)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(12)).toNumber()] = _2314_v5;
      _nw395[(new BigNumber(13)).toNumber()] = _2314_v5;
      _2315_v6 = _nw395;
      let _2316_v7;
      _2316_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_2315_v6);
      let _index322 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_2313_v4).length));
      (_2313_v4)[_index322] = (((_2316_v7).contains(_this.f23)) ? ((_2316_v7).get(_this.f23)) : (_2315_v6));
      let _index323 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_2313_v4).length));
      let _nw396 = Array((new BigNumber(4)).toNumber());
      _nw396[(_dafny.ZERO).toNumber()] = _2314_v5;
      _nw396[(_dafny.ONE).toNumber()] = _2314_v5;
      _nw396[(new BigNumber(2)).toNumber()] = (((_2314_v5).f30) ? (_2314_v5) : (_2314_v5));
      _nw396[(new BigNumber(3)).toNumber()] = _2314_v5;
      (_2313_v4)[_index323] = _nw396;
      let _ingredients1 = [];
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2306_v0).length))) {
        let _2317_i2 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2317_i2)) && ((_2317_i2).isLessThan(new BigNumber((_2306_v0).length))))) {
          _ingredients1.push(_dafny.Tuple.of(_2306_v0, (_2317_i2).toNumber(), _dafny.Seq.Concat(_dafny.Seq.Concat((_2306_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length))], _dafny.Seq.UnicodeFromString("rmtvniulv")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(310)), function (_2318_i3) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }))));
        }
      }
      for (const _tup1 of _ingredients1) {
        _tup1[0][_tup1[1]] = _tup1[2];
      }
      let _2319_v8;
      _2319_v8 = _module.D9.create_DC20((_2306_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length))], _module.__default.fm2(globalState), new BigNumber((_2309_v1).length));
      let _2320_v9;
      _2320_v9 = _dafny.MultiSet.fromElements(_2319_v8);
      let _2321_v10;
      _2321_v10 = _module.D21.create_DC50((_2320_v9).Difference(_2320_v9));
      let _source31 = _2321_v10;
      if (_source31.is_DC51) {
        let _2322___mcc_h0 = (_source31).cf81;
        let _2323_cf81 = _2322___mcc_h0;
        _2309_v1 = ((!(_2323_cf81) || ((_2314_v5).f30)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), ((_2324_p0) => function (_2325_i4) {
          return _2324_p0;
        })(p0))) : (_dafny.Seq.Concat(_module.__default.fm1((_this).f24, globalState), _2309_v1)));
        let _2326_v11;
        let _out65;
        _out65 = (_2314_v5).m4((_2314_v5).f30, globalState);
        _2326_v11 = _out65;
        let _2327_v12;
        _2327_v12 = _dafny.Seq.of((_2314_v5).f30, !((_2314_v5).f30), (_2319_v8).dtor_cf32);
        _2326_v11 = (new BigNumber((_2327_v12).length)).isLessThan(_this.f29);
        let _2328_v13;
        _2328_v13 = _dafny.Seq.of(_this.f29);
        let _2329_v14;
        _2329_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(843),(_2328_v13)[_module.__default.safeIndex(_this.f29, new BigNumber((_2328_v13).length))]);
        _2327_v12 = _dafny.Seq.update(_2327_v12, _module.__default.safeIndex((((_2329_v14).contains(new BigNumber(157))) ? ((_2329_v14).get(new BigNumber(157))) : (_this.f29)), new BigNumber((_2327_v12).length)), _2326_v11);
      } else if (_source31.is_DC50) {
        let _2330___mcc_h1 = (_source31).cf80;
        let _2331_cf80 = _2330___mcc_h1;
        let _2332_v15;
        let _nw397 = new _module.C9();
        _nw397.__ctor(_this.f23, (_this).f28);
        _2332_v15 = _nw397;
        let _2333_v16;
        _2333_v16 = _dafny.Seq.of((_2314_v5).f30, false);
        let _2334_v17;
        let _nw398 = Array((new BigNumber(26)).toNumber());
        _nw398[(_dafny.ZERO).toNumber()] = (_this).f27;
        _nw398[(_dafny.ONE).toNumber()] = (_2314_v5).f30;
        _nw398[(new BigNumber(2)).toNumber()] = (_this).fm4(_this.f29, _2333_v16, globalState);
        _nw398[(new BigNumber(3)).toNumber()] = true;
        _nw398[(new BigNumber(4)).toNumber()] = (_this).f28;
        _nw398[(new BigNumber(5)).toNumber()] = p1;
        _nw398[(new BigNumber(6)).toNumber()] = (_2314_v5).f30;
        _nw398[(new BigNumber(7)).toNumber()] = true;
        _nw398[(new BigNumber(8)).toNumber()] = (_this).f28;
        _nw398[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw398[(new BigNumber(10)).toNumber()] = false;
        _nw398[(new BigNumber(11)).toNumber()] = (_this).f28;
        _nw398[(new BigNumber(12)).toNumber()] = (_2310_v2).f39;
        _nw398[(new BigNumber(13)).toNumber()] = (_2310_v2).f39;
        _nw398[(new BigNumber(14)).toNumber()] = !(!((_2310_v2).f39));
        _nw398[(new BigNumber(15)).toNumber()] = (_2314_v5).f30;
        _nw398[(new BigNumber(16)).toNumber()] = p1;
        _nw398[(new BigNumber(17)).toNumber()] = !((_2310_v2).f39);
        _nw398[(new BigNumber(18)).toNumber()] = true;
        _nw398[(new BigNumber(19)).toNumber()] = (_2314_v5).f31;
        _nw398[(new BigNumber(20)).toNumber()] = (_this).f27;
        _nw398[(new BigNumber(21)).toNumber()] = p1;
        _nw398[(new BigNumber(22)).toNumber()] = (_2314_v5).f31;
        _nw398[(new BigNumber(23)).toNumber()] = (_2314_v5).fm4(_this.f29, _dafny.Seq.update(_dafny.Seq.of(false, (_this).f24, (_2310_v2).f39), _module.__default.safeIndex(_this.f29, new BigNumber((_dafny.Seq.of(false, (_this).f24, (_2310_v2).f39)).length)), (_2310_v2).f39), globalState);
        _nw398[(new BigNumber(24)).toNumber()] = (_2310_v2).f39;
        _nw398[(new BigNumber(25)).toNumber()] = (_2314_v5).f31;
        _2334_v17 = _nw398;
        let _2335_v18;
        let _nw399 = new _module.C5();
        _nw399.__ctor(_2334_v17);
        _2335_v18 = _nw399;
        if ((_this.f29).isLessThanOrEqualTo(_this.f29)) {
          (globalState).f7 = _this.f29;
          (globalState).f7 = _this.f29;
          let _2336_v19;
          _2336_v19 = _dafny.MultiSet.fromElements(new BigNumber(((_2306_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length))]).length), _this.f29, _this.f29, (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f29)), _this.f29);
          let _2337_v20;
          _2337_v20 = _module.D5.create_DC13((((_2336_v19).contains(_this.f29)) ? ((_2336_v19).get(_this.f29)) : (_this.f29)), true, false);
          _2337_v20 = _2337_v20;
          let _2338_v21;
          let _nw400 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _2338_v21 = _nw400;
          (globalState).f22 = _2338_v21;
          let _2339_v22;
          _2339_v22 = _dafny.Set.fromElements(true, (_2314_v5).f31);
          let _2340_v23;
          _2340_v23 = _dafny.Map.Empty.slice().updateUnsafe((_2339_v22).Intersect(_2339_v22),_2311_v3);
          _2340_v23 = (_2340_v23).update((_2339_v22).Difference(_2339_v22), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("e")));
        } else {
          (globalState).f7 = (_module.D5.create_DC13(_this.f29, (_2310_v2).f39, (_this).f24)).dtor_cf20;
          (_this).f29 = _this.f29;
          let _2341_v24;
          let _init71 = function (_2342_i5) {
            return (_2342_i5).multipliedBy(new BigNumber((_dafny.Seq.of(_this.f29, _this.f29)).length));
          };
          let _nw401 = Array((new BigNumber(2)).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw401.length); _i0_71++) {
            _nw401[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _2341_v24 = _nw401;
          let _index324 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_2341_v24).length));
          (_2341_v24)[_index324] = _this.f29;
          let _2343_v25;
          _2343_v25 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_2333_v16, _2333_v16)).length), _this.f29);
          let _index325 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_2341_v24).length));
          let _rhs348 = (_dafny.ZERO).minus(_this.f29);
          let _rhs349 = _this.f23;
          let _rhs350 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_this.f29), _module.__default.safeIndex(_this.f29, new BigNumber((_dafny.Seq.of(_this.f29)).length)), _this.f29), _2343_v25), _dafny.Seq.Concat(_2343_v25, _2343_v25));
          let _rhs351 = (_dafny.ZERO).minus(_this.f29);
          let _lhs265 = _2341_v24;
          let _lhs266 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_2341_v24).length));
          let _lhs267 = _this;
          let _lhs268 = globalState;
          _lhs265[_lhs266] = _rhs348;
          _lhs267.f23 = _rhs349;
          _2343_v25 = _rhs350;
          _lhs268.f7 = _rhs351;
          let _2344_v26;
          _2344_v26 = _module.D21.create_DC51(true);
          let _2345_v27;
          _2345_v27 = _dafny.Map.Empty.slice().updateUnsafe((_2344_v26).dtor_cf81,(_2314_v5).f31);
          _2345_v27 = (_2345_v27).update((_2314_v5).f30, (_2333_v16)[_module.__default.safeIndex(new BigNumber((_2309_v1).length), new BigNumber((_2333_v16).length))]);
          (globalState).f13 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_2309_v1, (_2306_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length))]), _2309_v1);
        }
        let _2346_v28;
        let _nw402 = new _module.C1();
        _nw402.__ctor((_2310_v2).f39, (_2310_v2).f39);
        _2346_v28 = _nw402;
      } else {
        let _2347___mcc_h2 = (_source31).cf82;
        let _2348_cf82 = _2347___mcc_h2;
        (_this).f29 = _this.f29;
        r2 = (_2314_v5).f31;
        let _2349_v29;
        let _nw403 = Array((new BigNumber(26)).toNumber()).fill(false);
        _2349_v29 = _nw403;
        let _2350_v30;
        let _nw404 = new _module.C5();
        _nw404.__ctor(_2349_v29);
        _2350_v30 = _nw404;
        let _2351_v31;
        _2351_v31 = _module.D16.create_DC37(_2350_v30);
        let _2352_v32;
        _2352_v32 = _module.D16.create_DC39(_2351_v31);
        let _pat_let_tv45 = _2351_v31;
        let _2353_v33;
        let _nw405 = Array((new BigNumber(17)).toNumber());
        _nw405[(_dafny.ZERO).toNumber()] = _2352_v32;
        _nw405[(_dafny.ONE).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(2)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(3)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(4)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(5)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(6)).toNumber()] = _module.D16.create_DC39(_2351_v31);
        _nw405[(new BigNumber(7)).toNumber()] = (((_2314_v5).f31) ? (_2352_v32) : (_2352_v32));
        _nw405[(new BigNumber(8)).toNumber()] = function (_pat_let57_0) {
          return function (_2354_dt__update__tmp_h0) {
            return function (_pat_let58_0) {
              return function (_2355_dt__update_hcf65_h0) {
                return _module.D16.create_DC39(_2355_dt__update_hcf65_h0);
              }(_pat_let58_0);
            }(_pat_let_tv45);
          }(_pat_let57_0);
        }(_2352_v32);
        _nw405[(new BigNumber(9)).toNumber()] = _module.D16.create_DC39(_2351_v31);
        _nw405[(new BigNumber(10)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(11)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(12)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(13)).toNumber()] = _module.D16.create_DC39(_2351_v31);
        _nw405[(new BigNumber(14)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(15)).toNumber()] = _2352_v32;
        _nw405[(new BigNumber(16)).toNumber()] = _2352_v32;
        _2353_v33 = _nw405;
        let _index326 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_2353_v33).length));
        (_2353_v33)[_index326] = _module.D16.create_DC39(_2351_v31);
        let _index327 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_2353_v33).length));
        (_2353_v33)[_index327] = _2352_v32;
        let _2356_v34;
        _2356_v34 = _module.D21.create_DC51((_2314_v5).f30);
        _2356_v34 = _2356_v34;
      }
      r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-100)), ((_2357_p0) => function (_2358_i6) {
        return _2357_p0;
      })(p0));
      let _2359_v35;
      _2359_v35 = _module.D0.create_DC1(_module.__default.safeDivisionInt(new BigNumber(((_2306_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_2306_v0).length))]).length), _this.f29));
      r1 = _2359_v35;
      let _2360_v36;
      _2360_v36 = _dafny.MultiSet.fromElements(_this.f29);
      r2 = (_this.f29).isLessThan(new BigNumber(((((_this).f24) ? (_2360_v36) : (_2360_v36))).cardinality()));
      return [r0, r1, r2];
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      (globalState).f13 = (_this).f24;
      (globalState).f7 = _module.__default.safeModuloInt(_this.f29, _this.f29);
      let _2361_v0;
      let _nw406 = Array((new BigNumber(22)).toNumber()).fill(_module.D0.Default());
      _2361_v0 = _nw406;
      let _2362_v1;
      _2362_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,new BigNumber(-956));
      let _2363_v2;
      _2363_v2 = _module.D0.create_DC1(new BigNumber((_2362_v1).length));
      let _2364_v3;
      _2364_v3 = _module.D0.create_DC3(_2363_v2);
      let _index328 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_2361_v0).length));
      (_2361_v0)[_index328] = _2364_v3;
      let _index329 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_2361_v0).length));
      (_2361_v0)[_index329] = _2364_v3;
      let _2365_v4;
      let _nw407 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _2365_v4 = _nw407;
      let _2366_v5;
      _2366_v5 = _dafny.MultiSet.fromElements(_2365_v4);
      (globalState).f7 = (_this.f29).minus(new BigNumber((_2366_v5).cardinality()));
      let _2367_v6;
      _2367_v6 = _dafny.Seq.UnicodeFromString("i");
      let _2368_v7;
      _2368_v7 = _dafny.MultiSet.fromElements(_2367_v6, _2367_v6);
      let _2369_v8;
      _2369_v8 = (_2368_v7).Difference(_dafny.MultiSet.fromElements(_2367_v6, _2367_v6));
      let _source32 = _2369_v8;
      let _2370___mcc_h0 = _source32;
      let _2371_cf79 = _2370___mcc_h0;
      let _2372_v9;
      _2372_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(865));
      let _2373_v10;
      _2373_v10 = _dafny.MultiSet.fromElements(new BigNumber((_2372_v9).length), _this.f29);
      let _2374_v11;
      _2374_v11 = _dafny.Map.Empty.slice().updateUnsafe(false,(_2373_v10).IsProperSubsetOf(_2373_v10));
      let _2375_v12;
      let _nw408 = Array((new BigNumber(23)).toNumber()).fill(false);
      _2375_v12 = _nw408;
      let _2376_v13;
      let _nw409 = new _module.C5();
      _nw409.__ctor(_2375_v12);
      _2376_v13 = _nw409;
      let _2377_v14;
      _2377_v14 = _dafny.Set.fromElements(_2376_v13, _2376_v13, _2376_v13, _2376_v13, _2376_v13);
      _2374_v11 = (_2374_v11).update(_module.__default.fm2(globalState), (_2377_v14).IsDisjointFrom(_dafny.Set.fromElements(_2376_v13)));
      let _2378_v15;
      _2378_v15 = _dafny.Set.fromElements(_2375_v12, (_2376_v13).f32, _2375_v12);
      (globalState).f13 = (_2378_v15).IsProperSubsetOf(_2378_v15);
      (globalState).f13 = (_this).f28;
      let _2379_v16;
      _2379_v16 = _dafny.MultiSet.fromElements((_this).f27);
      let _2380_v17;
      _2380_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f25);
      let _2381_v18;
      let _nw410 = new _module.C2();
      _nw410.__ctor(_2379_v16, (((_2380_v17).contains((_this).f27)) ? ((_2380_v17).get((_this).f27)) : ((_this).f25)), _this.f23, (_this).f24);
      _2381_v18 = _nw410;
      let _hi14 = _this.f29;
      for (let _2382_i0 = (_dafny.ZERO).minus(_this.f29); _2382_i0.isLessThan(_hi14); _2382_i0 = _2382_i0.plus(_dafny.ONE)) {
        let _index330 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length));
        (_2365_v4)[_index330] = _2382_i0;
        let _2383_v19;
        _2383_v19 = _dafny.Seq.of(_this.f29);
        let _index331 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length));
        (_2365_v4)[_index331] = (_module.__default.safeDivisionInt(_this.f29, (_2383_v19)[_module.__default.safeIndex(_this.f29, new BigNumber((_2383_v19).length))])).multipliedBy((_2382_i0).minus(_2382_i0));
        let _2384_v20;
        _2384_v20 = _module.D12.create_DC25((_this).f24, (_2367_v6)[_module.__default.safeIndex(_this.f29, new BigNumber((_2367_v6).length))]);
        let _2385_v21;
        _2385_v21 = _module.D12.create_DC26(_2384_v20);
        let _2386_v22;
        _2386_v22 = _module.D12.create_DC26(_2384_v20);
        let _source33 = _2386_v22;
        if (_source33.is_DC25) {
          let _2387___mcc_h1 = (_source33).cf41;
          let _2388___mcc_h2 = (_source33).cf42;
          let _2389_cf42 = _2388___mcc_h2;
          let _2390_cf41 = _2387___mcc_h1;
          let _2391_v23;
          _2391_v23 = _dafny.MultiSet.fromElements((_this).f28, (_this).f24, (_this).f27);
          let _2392_v24;
          let _2393_v25;
          let _2394_v26;
          let _out66;
          let _out67;
          let _out68;
          let _outcollector20 = _module.__default.m0(new BigNumber((_2367_v6).length), _2391_v23, _2390_cf41, ((_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))]).plus((_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))]), globalState);
          _out66 = _outcollector20[0];
          _out67 = _outcollector20[1];
          _out68 = _outcollector20[2];
          _2392_v24 = _out66;
          _2393_v25 = _out67;
          _2394_v26 = _out68;
          r2 = (_this).f28;
          let _2395_v27;
          _2395_v27 = _module.D0.create_DC2((_this).f24, (_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))], (_this).f24);
          _2390_cf41 = (_2395_v27).dtor_cf4;
          let _index332 = _module.__default.safeIndex(new BigNumber(654), new BigNumber(((_this).f25).length));
          ((_this).f25)[_index332] = _2389_cf42;
          let _index333 = _module.__default.safeIndex(new BigNumber(654), new BigNumber(((_this).f25).length));
          ((_this).f25)[_index333] = _this.f23;
        } else if (_source33.is_DC24) {
          let _2396___mcc_h3 = (_source33).cf40;
          let _2397_cf40 = _2396___mcc_h3;
          let _2398_v28;
          let _nw411 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _2398_v28 = _nw411;
          let _2399_v29;
          let _nw412 = Array((new BigNumber(4)).toNumber()).fill(false);
          _2399_v29 = _nw412;
          let _2400_v30;
          _2400_v30 = _module.D12.create_DC25((_this).f24, _this.f23);
          let _2401_v31;
          let _nw413 = new _module.C3();
          _nw413.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_2399_v29,(_2400_v30).dtor_cf42), _2397_cf40, _this.f23, (_this).f24);
          _2401_v31 = _nw413;
          let _2402_v32;
          _2402_v32 = _dafny.Map.Empty.slice().updateUnsafe((_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))],_2401_v31);
          let _index334 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_2398_v28).length));
          (_2398_v28)[_index334] = _2402_v32;
          let _2403_v34;
          _2403_v34 = _dafny.Set.fromElements((_this).fm9(_2383_v19, _2382_i0, (_this).f24, globalState), _2382_i0);
          let _2404_v35;
          _2404_v35 = _dafny.Set.fromElements(!((_this).f24) || (_module.__default.fm2(globalState)), (_2403_v34).IsSubsetOf(function () {
            let _coll95 = new _dafny.Set();
            for (const _compr_95 of _dafny.IntegerRange(new BigNumber(-657), new BigNumber(218))) {
              let _2405_v33 = _compr_95;
              if (((new BigNumber(-657)).isLessThanOrEqualTo(_2405_v33)) && ((_2405_v33).isLessThan(new BigNumber(218)))) {
                _coll95.add((_2405_v33).minus(_this.f29));
              }
            }
            return _coll95;
          }()));
          let _index335 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length));
          let _index336 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_2398_v28).length));
          let _rhs352 = new BigNumber((_2404_v35).length);
          let _rhs353 = _2402_v32;
          let _lhs269 = _2365_v4;
          let _lhs270 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length));
          let _lhs271 = _2398_v28;
          let _lhs272 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_2398_v28).length));
          _lhs269[_lhs270] = _rhs352;
          _lhs271[_lhs272] = _rhs353;
          (globalState).f7 = new BigNumber((_2367_v6).length);
          let _2406_v36;
          _2406_v36 = _module.D15.create_DC34(_module.__default.fm2(globalState));
          let _2407_v37;
          _2407_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f24);
          let _2408_v38;
          _2408_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2406_v36,new BigNumber((_2407_v37).length));
          let _2409_v39;
          let _nw414 = new _module.C6();
          _nw414.__ctor(_this.f23, (_this).f24);
          _2409_v39 = _nw414;
          let _2410_v40;
          _2410_v40 = _dafny.Map.Empty.slice().updateUnsafe(false,_2409_v39);
          let _2411_v41;
          _2411_v41 = _dafny.MultiSet.fromElements(_2382_i0);
          let _2412_v42;
          _2412_v42 = _dafny.Map.Empty.slice().updateUnsafe((_2408_v38).update(_2406_v36, new BigNumber((_dafny.MultiSet.fromElements(_2410_v40, _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_2409_v39), _2410_v40, _2410_v40, _2410_v40)).cardinality())),_2411_v41);
          let _2413_v43;
          _2413_v43 = _dafny.Set.fromElements(_2367_v6, _dafny.Seq.update(_2367_v6, _module.__default.safeIndex(_this.f29, new BigNumber((_2367_v6).length)), new _dafny.CodePoint('u'.codePointAt(0))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aeorvcm"), _2367_v6), _2367_v6);
          let _2414_v44;
          _2414_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_2399_v29);
          let _2415_v45;
          let _nw415 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _2415_v45 = _nw415;
          let _2416_v46;
          _2416_v46 = _dafny.Seq.of(_2367_v6);
          let _rhs354 = _2412_v42;
          let _rhs355 = (_2414_v44).contains(true);
          let _rhs356 = _2415_v45;
          let _rhs357 = function () {
            let _coll96 = new _dafny.Set();
            for (const _compr_96 of (_2416_v46).Elements) {
              let _2417_v47 = _compr_96;
              if (_dafny.Seq.contains(_2416_v46, _2417_v47)) {
                _coll96.add(_2417_v47);
              }
            }
            return _coll96;
          }();
          let _lhs273 = globalState;
          let _lhs274 = globalState;
          _2412_v42 = _rhs354;
          _lhs273.f13 = _rhs355;
          _lhs274.f22 = _rhs356;
          _2413_v43 = _rhs357;
          let _index337 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_2399_v29).length));
          (_2399_v29)[_index337] = (_this.f29).isLessThanOrEqualTo(_module.__default.fm0((_this).f28, _this.f29, globalState));
          let _2418_v48;
          _2418_v48 = _dafny.Seq.of(!(!(new BigNumber(-418)).isEqualTo(new BigNumber((_2367_v6).length))), (_this).f28, true);
          let _index338 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_2399_v29).length));
          (_2399_v29)[_index338] = !((_2418_v48)[_module.__default.safeIndex(new BigNumber(-95), new BigNumber((_2418_v48).length))]);
        } else {
          let _2419___mcc_h4 = (_source33).cf43;
          let _2420_cf43 = _2419___mcc_h4;
          let _2421_v49;
          let _nw416 = Array((new BigNumber(23)).toNumber()).fill(false);
          _2421_v49 = _nw416;
          _2421_v49 = _2421_v49;
          r2 = (_this).f24;
          let _2422_v50;
          _2422_v50 = _dafny.MultiSet.fromElements(false, (_this).f24);
          let _2423_v51;
          let _2424_v52;
          let _2425_v53;
          let _out69;
          let _out70;
          let _out71;
          let _outcollector21 = _module.__default.m0(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_2382_i0), _2382_i0), _2422_v50, (((_this).f28) ? ((_this).f28) : ((_this).f28)), _this.f29, globalState);
          _out69 = _outcollector21[0];
          _out70 = _outcollector21[1];
          _out71 = _outcollector21[2];
          _2423_v51 = _out69;
          _2424_v52 = _out70;
          _2425_v53 = _out71;
          let _2426_v54;
          let _nw417 = new _module.C8();
          _nw417.__ctor(_2424_v52);
          _2426_v54 = _nw417;
          let _2427_v55;
          _2427_v55 = _2426_v54;
          let _2428_v56;
          _2428_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(298),(_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))])).length),_2426_v54);
          let _2429_v57;
          _2429_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2367_v6);
          let _2430_v58;
          _2430_v58 = _module.D2.create_DC6(_2429_v57);
          let _2431_v59;
          _2431_v59 = _dafny.MultiSet.fromElements(_2423_v51, new BigNumber((_module.__default.fm45((_this).f27, _2430_v58, (_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))], globalState)).length));
          let _2432_v60;
          let _nw418 = Array((new BigNumber(19)).toNumber());
          _nw418[(_dafny.ZERO).toNumber()] = _2426_v54;
          _nw418[(_dafny.ONE).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(2)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(3)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(4)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(5)).toNumber()] = (_2427_v55);
          _nw418[(new BigNumber(6)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(7)).toNumber()] = (_2427_v55);
          _nw418[(new BigNumber(8)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(9)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(10)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(11)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(12)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(13)).toNumber()] = (((_2428_v56).contains((((_2431_v59).contains(_2425_v53)) ? ((_2431_v59).get(_2425_v53)) : ((_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))])))) ? ((_2428_v56).get((((_2431_v59).contains(_2425_v53)) ? ((_2431_v59).get(_2425_v53)) : ((_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))])))) : (_2426_v54));
          _nw418[(new BigNumber(14)).toNumber()] = (_2427_v55);
          _nw418[(new BigNumber(15)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(16)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(17)).toNumber()] = _2426_v54;
          _nw418[(new BigNumber(18)).toNumber()] = _2426_v54;
          _2432_v60 = _nw418;
          let _index339 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_2432_v60).length));
          (_2432_v60)[_index339] = _2426_v54;
          let _2433_v61;
          _2433_v61 = _dafny.Map.Empty.slice().updateUnsafe((_2426_v54).f27,false);
          let _2434_v62;
          _2434_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2433_v61,_2367_v6);
          let _2435_v63;
          _2435_v63 = _dafny.Seq.of(_2367_v6, (((_2434_v62).contains(_2433_v61)) ? ((_2434_v62).get(_2433_v61)) : (_2367_v6)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_2436_i1) {
            return _this.f23;
          }));
          let _2437_v64;
          _2437_v64 = _dafny.MultiSet.fromElements(_2383_v19);
          let _2438_v65;
          _2438_v65 = _dafny.Seq.of(_2424_v52, (_this).f24, (_this).f27, false);
          let _2439_v67;
          _2439_v67 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_this.f23);
          let _2440_v68;
          _2440_v68 = _dafny.Map.Empty.slice().updateUnsafe((_2365_v4)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_2365_v4).length))],new BigNumber((_2439_v67).length));
          let _2441_v69;
          _2441_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2440_v68).length),(_this).f27);
          let _2442_v70;
          _2442_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2438_v65,_2440_v68);
          let _2443_v71;
          _2443_v71 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_2438_v65,function () {
            let _coll97 = new _dafny.Map();
            for (const _compr_97 of (_2441_v69).Keys.Elements) {
              let _2444_v66 = _compr_97;
              if ((_2441_v69).contains(_2444_v66)) {
                _coll97.push([_module.__default.safeModuloInt(_2444_v66, _this.f29),new BigNumber((_2438_v65).length)]);
              }
            }
            return _coll97;
          }())).Merge(_2442_v70),_2426_v54);
          let _index340 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_2432_v60).length));
          let _rhs358 = (new BigNumber((_2435_v63).length)).multipliedBy((((_2437_v64).contains(_2383_v19)) ? ((_2437_v64).get(_2383_v19)) : (_2423_v51)));
          let _rhs359 = (((_2443_v71).contains(_2442_v70)) ? ((_2443_v71).get(_2442_v70)) : (_2426_v54));
          let _lhs275 = globalState;
          let _lhs276 = _2432_v60;
          let _lhs277 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_2432_v60).length));
          _lhs275.f7 = _rhs358;
          _lhs276[_lhs277] = _rhs359;
        }
        (_this).f29 = _this.f29;
        let _2445_v72;
        let _nw419 = new _module.C1();
        _nw419.__ctor((_this).f28, (((_this).f24) ? ((_this).f28) : ((_this).f28)));
        _2445_v72 = _nw419;
      }
      let _2446_v73;
      _2446_v73 = _dafny.Seq.of(new BigNumber(718), _this.f29);
      r0 = ((_dafny.ZERO).minus(new BigNumber((_2446_v73).length))).isLessThan(_this.f29);
      r1 = _this.f29;
      let _2447_v74;
      _2447_v74 = _dafny.Set.fromElements((_this).f24);
      let _2448_v75;
      _2448_v75 = _dafny.Seq.of((_this).f24);
      r2 = !((_this).f28) || ((_2447_v74).IsDisjointFrom(_dafny.Set.fromElements((_2448_v75)[_module.__default.safeIndex(_this.f29, new BigNumber((_2448_v75).length))])));
      return [r0, r1, r2];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _hi15 = _this.f29;
      for (let _2449_i0 = _this.f29; _2449_i0.isLessThan(_hi15); _2449_i0 = _2449_i0.plus(_dafny.ONE)) {
        (globalState).f13 = (_this).f28;
        let _2450_v0;
        _2450_v0 = _dafny.MultiSet.fromElements((_this).f24);
        let _2451_v1;
        _2451_v1 = _dafny.Seq.of(_2449_i0, _this.f29, new BigNumber(-260));
        let _2452_v2;
        let _2453_v3;
        let _2454_v4;
        let _out72;
        let _out73;
        let _out74;
        let _outcollector22 = _module.__default.m0(_this.f29, (_2450_v0).Difference((_2450_v0).update(p0, _module.__default.abs(_2449_i0))), (_this).f28, _module.__default.fm0((_this).f27, (_2451_v1)[_module.__default.safeIndex(_this.f29, new BigNumber((_2451_v1).length))], globalState), globalState);
        _out72 = _outcollector22[0];
        _out73 = _outcollector22[1];
        _out74 = _outcollector22[2];
        _2452_v2 = _out72;
        _2453_v3 = _out73;
        _2454_v4 = _out74;
        let _2455_v5;
        let _nw420 = new _module.C6();
        _nw420.__ctor((_this).fm10((_this).f27, _2452_v2, globalState), _module.__default.fm2(globalState));
        _2455_v5 = _nw420;
        let _2456_v6;
        _2456_v6 = _module.D23.create_DC54(_2455_v5);
        _2455_v5 = (_2456_v6).dtor_cf84;
        let _2457_v7;
        _2457_v7 = _dafny.Set.fromElements(_2452_v2);
        let _2458_v8;
        _2458_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,((!(p0)) ? (_2457_v7) : (_2457_v7)));
        let _2459_v10;
        _2459_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2449_i0,_2452_v2);
        let _2460_v11;
        let _nw421 = new _module.C4();
        _nw421.__ctor(new BigNumber((function () {
          let _coll98 = new _dafny.Set();
          for (const _compr_98 of _dafny.IntegerRange(new BigNumber(272), new BigNumber(866))) {
            let _2461_v9 = _compr_98;
            if (((new BigNumber(272)).isLessThanOrEqualTo(_2461_v9)) && ((_2461_v9).isLessThan(new BigNumber(866)))) {
              _coll98.add((_2461_v9).minus(_2454_v4));
            }
          }
          return _coll98;
        }()).length), _2459_v10, (_this).f25, _this.f23, _2453_v3);
        _2460_v11 = _nw421;
        let _2462_v12;
        _2462_v12 = _dafny.Seq.of(_2460_v11);
        let _2463_v13;
        _2463_v13 = _dafny.Seq.of(p0);
        let _rhs360 = (_dafny.Map.Empty.slice().updateUnsafe(_this.f23,_dafny.Set.fromElements(new BigNumber(687)))).Merge((_2458_v8).Merge(_2458_v8));
        let _rhs361 = new BigNumber((_dafny.MultiSet.FromArray(_2451_v1)).cardinality());
        let _rhs362 = _module.__default.safeModuloInt(new BigNumber((_2462_v12).length), new BigNumber((_dafny.Seq.Concat(_2463_v13, _2463_v13)).length));
        let _rhs363 = _2453_v3;
        let _lhs278 = _this;
        let _lhs279 = globalState;
        _2458_v8 = _rhs360;
        _lhs278.f29 = _rhs361;
        r1 = _rhs362;
        _lhs279.f13 = _rhs363;
      }
      let _hi16 = _this.f29;
      for (let _2464_i1 = new BigNumber(186); _2464_i1.isLessThan(_hi16); _2464_i1 = _2464_i1.plus(_dafny.ONE)) {
        let _2465_v14;
        _2465_v14 = _dafny.Set.fromElements(!((_this).f27), p0, p0);
        let _2466_v15;
        _2466_v15 = _dafny.Seq.UnicodeFromString("eeajmfk");
        let _2467_v16;
        _2467_v16 = _module.D19.create_DC46((_2465_v14).Union(_2465_v14), _2466_v15);
        _2467_v16 = _2467_v16;
        (_this).f29 = _this.f29;
        (globalState).f7 = (_2464_i1).plus(_this.f29);
        (_this).f23 = _this.f23;
      }
      let _2468_v17;
      let _nw422 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
      _2468_v17 = _nw422;
      let _2469_v18;
      let _init72 = function (_2470_i2) {
        return (_this).f28;
      };
      let _nw423 = Array((new BigNumber(9)).toNumber());
      for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw423.length); _i0_72++) {
        _nw423[_i0_72] = _init72(new BigNumber(_i0_72));
      }
      _2469_v18 = _nw423;
      let _2471_v19;
      _2471_v19 = _dafny.Set.fromElements(_2469_v18, _2469_v18, _2469_v18, _2469_v18, _2469_v18);
      let _index341 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2468_v17).length));
      (_2468_v17)[_index341] = _2471_v19;
      let _index342 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2468_v17).length));
      (_2468_v17)[_index342] = ((((_this).f27) ? (_2471_v19) : (_dafny.Set.fromElements(_2469_v18, _2469_v18)))).Intersect(_2471_v19);
      r2 = (_this).f24;
      let _2472_v20;
      let _nw424 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
      _2472_v20 = _nw424;
      let _2473_v21;
      let _nw425 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _2473_v21 = _nw425;
      let _2474_v22;
      _2474_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2473_v21);
      let _index343 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_2472_v20).length));
      (_2472_v20)[_index343] = (_2474_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2473_v21));
      let _index344 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_2472_v20).length));
      (_2472_v20)[_index344] = ((_dafny.Map.Empty.slice().updateUnsafe(!(p0),_2473_v21)).Merge(_2474_v22)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2473_v21));
      let _hi17 = _this.f29;
      for (let _2475_i3 = _this.f29; _2475_i3.isLessThan(_hi17); _2475_i3 = _2475_i3.plus(_dafny.ONE)) {
        let _2476_v23;
        _2476_v23 = _dafny.Set.fromElements((_this).f24);
        let _2477_v24;
        let _nw426 = new _module.C5();
        _nw426.__ctor(_2469_v18);
        _2477_v24 = _nw426;
        let _2478_v25;
        _2478_v25 = _dafny.Set.fromElements(_2477_v24, _2477_v24);
        let _rhs364 = _2473_v21;
        let _rhs365 = true;
        let _rhs366 = (_dafny.ZERO).minus(new BigNumber((((((_this).f24) ? (_2476_v23) : (_2476_v23))).Union(_2476_v23)).length));
        let _rhs367 = (((_this).f27) ? (((false) ? (_module.__default.fm2(globalState)) : (p0))) : ((_2478_v25).IsProperSubsetOf(_2478_v25)));
        let _lhs280 = globalState;
        let _lhs281 = _this;
        _lhs280.f22 = _rhs364;
        r0 = _rhs365;
        _lhs281.f29 = _rhs366;
        r2 = _rhs367;
        let _2479_v26;
        _2479_v26 = _dafny.Seq.of(_this.f29);
        let _2480_v27;
        let _nw427 = new _module.C6();
        _nw427.__ctor(p1, p0);
        _2480_v27 = _nw427;
        let _2481_v28;
        let _nw428 = Array((new BigNumber(17)).toNumber());
        _nw428[(_dafny.ZERO).toNumber()] = _2480_v27;
        _nw428[(_dafny.ONE).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(2)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(3)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(4)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(5)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(6)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(7)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(8)).toNumber()] = ((p0) ? (_2480_v27) : (_2480_v27));
        _nw428[(new BigNumber(9)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(10)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(11)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(12)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(13)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(14)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(15)).toNumber()] = _2480_v27;
        _nw428[(new BigNumber(16)).toNumber()] = _2480_v27;
        _2481_v28 = _nw428;
        let _index345 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_2481_v28).length));
        (_2481_v28)[_index345] = _2480_v27;
        let _2482_v29;
        _2482_v29 = _dafny.Seq.UnicodeFromString("vfbyqdoi");
        let _2483_v30;
        _2483_v30 = _dafny.Seq.of(p0);
        let _2484_v32;
        _2484_v32 = _module.D5.create_DC12(function () {
  let _coll99 = new _dafny.Set();
  for (const _compr_99 of (_dafny.Seq.of(new BigNumber((_2482_v29).length), new BigNumber((_2483_v30).length))).Elements) {
    let _2485_v31 = _compr_99;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_2482_v29).length), new BigNumber((_2483_v30).length)), _2485_v31)) {
      _coll99.add((_2485_v31).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ru")).length)));
    }
  }
  return _coll99;
}());
        let _2486_v33;
        _2486_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2482_v29);
        let _2487_v34;
        _2487_v34 = _module.D2.create_DC6(_2486_v33);
        let _index346 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_2481_v28).length));
        let _rhs368 = _dafny.Seq.Concat(_dafny.Seq.of(_2475_i3), _module.__default.fm45((_this).f28, _2487_v34, new BigNumber(633), globalState));
        let _rhs369 = _2480_v27.f23;
        let _rhs370 = _2480_v27;
        let _rhs371 = _dafny.Seq.Concat(_2482_v29, _2482_v29);
        let _rhs372 = ((_module.__default.fm2(globalState)) ? (_2484_v32) : (_2484_v32));
        let _lhs282 = _this;
        let _lhs283 = _2481_v28;
        let _lhs284 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_2481_v28).length));
        _2479_v26 = _rhs368;
        _lhs282.f23 = _rhs369;
        _lhs283[_lhs284] = _rhs370;
        _2482_v29 = _rhs371;
        _2484_v32 = _rhs372;
        if (!(false) || ((_this).f27)) {
          (_this).f23 = _2480_v27.f23;
          (_this).f29 = (_dafny.ZERO).minus(_2475_i3);
          let _2488_v35;
          _2488_v35 = _module.D1.create_DC4(_2480_v27.f23);
          (_this).f23 = (_2488_v35).dtor_cf6;
          let _2489_v36;
          let _init73 = ((_2490_v29) => function (_2491_i4) {
            return _2490_v29;
          })(_2482_v29);
          let _nw429 = Array((new BigNumber(12)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw429.length); _i0_73++) {
            _nw429[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2489_v36 = _nw429;
          let _2492_v37;
          _2492_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2482_v29,_2482_v29);
          let _index347 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_2489_v36).length));
          (_2489_v36)[_index347] = (((_2492_v37).contains(_2482_v29)) ? ((_2492_v37).get(_2482_v29)) : (_2482_v29));
          let _index348 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_2489_v36).length));
          (_2489_v36)[_index348] = _dafny.Seq.Concat(_2482_v29, _dafny.Seq.Concat(_2482_v29, _dafny.Seq.UnicodeFromString("lydhu")));
          (globalState).f7 = _this.f29;
        } else {
          let _2493_v38;
          _2493_v38 = _dafny.MultiSet.fromElements((_this).f24, (_2480_v27).f24, true);
          let _2494_v39;
          _2494_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_2493_v38),_dafny.Seq.Concat(_2482_v29, _2482_v29));
          let _2495_v40;
          _2495_v40 = _module.D0.create_DC0(_2493_v38);
          _2494_v39 = (_2494_v39).update(_2495_v40, _dafny.Seq.Concat(_2482_v29, _dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), ((_2496_v27) => function (_2497_i5) {
            return _2496_v27.f23;
          })(_2480_v27))));
          r1 = (_this.f29).multipliedBy(new BigNumber(968));
          let _2498_v41;
          let _nw430 = Array((new BigNumber(7)).toNumber());
          _nw430[(_dafny.ZERO).toNumber()] = _2480_v27.f23;
          _nw430[(_dafny.ONE).toNumber()] = _this.f23;
          _nw430[(new BigNumber(2)).toNumber()] = _2480_v27.f23;
          _nw430[(new BigNumber(3)).toNumber()] = _module.__default.fm29(_2475_i3, (_this).f28, (_this).f27, (_dafny.ZERO).minus(_2475_i3), globalState);
          _nw430[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
          _nw430[(new BigNumber(5)).toNumber()] = _2480_v27.f23;
          _nw430[(new BigNumber(6)).toNumber()] = _this.f23;
          _2498_v41 = _nw430;
          _2498_v41 = _2498_v41;
          (globalState).f13 = !(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_2479_v26)[_module.__default.safeIndex(_this.f29, new BigNumber((_2479_v26).length))]), new BigNumber(471))).isEqualTo(new BigNumber(552));
          let _2499_v42;
          _2499_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2475_i3,_2482_v29);
          _2499_v42 = (_2499_v42).update(_module.__default.safeModuloInt(_this.f29, new BigNumber((_2476_v23).length)), _2482_v29);
        }
        let _index349 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_2469_v18).length));
        (_2469_v18)[_index349] = ((_this).f28) && (p0);
        let _2500_v43;
        _2500_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f27);
        let _2501_v44;
        _2501_v44 = _dafny.Seq.of(_2500_v43, _2500_v43, _2500_v43, _2500_v43);
        let _2502_v45;
        _2502_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2473_v21,(_2501_v44)[_module.__default.safeIndex(new BigNumber((_2476_v23).length), new BigNumber((_2501_v44).length))]);
        let _index350 = _module.__default.safeIndex(new BigNumber(619), new BigNumber(((_2477_v24).f32).length));
        ((_2477_v24).f32)[_index350] = (_this).f28;
        let _2503_v46;
        _2503_v46 = _dafny.Map.Empty.slice().updateUnsafe((_2477_v24).f32,p1);
        let _2504_v47;
        let _nw431 = new _module.C3();
        _nw431.__ctor(_2503_v46, (_this).f25, _this.f23, (_this).f27);
        _2504_v47 = _nw431;
        let _2505_v48;
        _2505_v48 = _dafny.Set.fromElements(_2504_v47, _2504_v47, _2504_v47, _2504_v47, _2504_v47);
        let _index351 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_2469_v18).length));
        let _index352 = _module.__default.safeIndex(new BigNumber(619), new BigNumber(((_2477_v24).f32).length));
        let _rhs373 = !(!(new BigNumber(916)).isEqualTo((_this.f29).plus(_2475_i3)));
        let _rhs374 = _2502_v45;
        let _rhs375 = (_2480_v27).fm4(new BigNumber((_2505_v48).length), _module.__default.fm36(globalState), globalState);
        let _lhs285 = _2469_v18;
        let _lhs286 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_2469_v18).length));
        let _lhs287 = (_2477_v24).f32;
        let _lhs288 = _module.__default.safeIndex(new BigNumber(619), new BigNumber(((_2477_v24).f32).length));
        _lhs285[_lhs286] = _rhs373;
        _2502_v45 = _rhs374;
        _lhs287[_lhs288] = _rhs375;
      }
      r0 = (_this).f28;
      r1 = _this.f29;
      let _2506_v49;
      _2506_v49 = _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(false,p0), (_this).f27, p0);
      r2 = (_2506_v49).dtor_cf26;
      return [r0, r1, r2];
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _2507_v0;
      _2507_v0 = _dafny.Seq.of(_this.f29, _this.f29);
      let _2508_v1;
      _2508_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2507_v0,_this.f23);
      let _2509_v2;
      _2509_v2 = _dafny.MultiSet.fromElements((((_2508_v1).contains(_2507_v0)) ? ((_2508_v1).get(_2507_v0)) : (_this.f23)), _this.f23);
      _2509_v2 = _2509_v2;
      let _2510_v3;
      _2510_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p0);
      let _2511_v4;
      _2511_v4 = _dafny.Map.Empty.slice().updateUnsafe((((_2510_v3).contains((_this).f24)) ? ((_2510_v3).get((_this).f24)) : ((_this).f28)),(_this.f29).minus(_this.f29));
      let _2512_v5;
      _2512_v5 = _dafny.MultiSet.fromElements((_this).f28, (_this).f28, !((_this).f27), (_this).f28);
      let _2513_v6;
      _2513_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2512_v5).cardinality()),(_this).f24);
      _2511_v4 = (_2511_v4).update((new BigNumber((_module.__default.fm3((_this).f24, globalState)).cardinality())).isLessThanOrEqualTo(new BigNumber((_2513_v6).length)), _this.f29);
      (globalState).f7 = _this.f29;
      (_this).f29 = (_this).fm6(globalState);
      let _2514_v7;
      let _nw432 = Array((new BigNumber(17)).toNumber());
      _2514_v7 = _nw432;
      let _2515_v8;
      _2515_v8 = _dafny.Seq.of(_2514_v7);
      let _2516_v9;
      let _nw433 = Array((new BigNumber(13)).toNumber());
      _nw433[(_dafny.ZERO).toNumber()] = _2514_v7;
      _nw433[(_dafny.ONE).toNumber()] = (_2515_v8)[_module.__default.safeIndex(_this.f29, new BigNumber((_2515_v8).length))];
      _nw433[(new BigNumber(2)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(3)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(4)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(5)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(6)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(7)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(8)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(9)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(10)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(11)).toNumber()] = _2514_v7;
      _nw433[(new BigNumber(12)).toNumber()] = _2514_v7;
      _2516_v9 = _nw433;
      let _index353 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_2516_v9).length));
      (_2516_v9)[_index353] = _2514_v7;
      let _index354 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_2516_v9).length));
      let _rhs376 = (_this).f24;
      let _rhs377 = _2514_v7;
      let _lhs289 = _2516_v9;
      let _lhs290 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_2516_v9).length));
      r0 = _rhs376;
      _lhs289[_lhs290] = _rhs377;
      let _2517_v10;
      _2517_v10 = _dafny.MultiSet.fromElements(_this.f29, new BigNumber((_2512_v5).cardinality()));
      let _2518_v11;
      _2518_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber(((_2517_v10).Difference(_2517_v10)).cardinality()));
      let _2519_v12;
      _2519_v12 = _module.D19.create_DC47((_this).f24);
      let _2520_v13;
      let _nw434 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2520_v13 = _nw434;
      let _2521_v14;
      _2521_v14 = _dafny.Seq.UnicodeFromString("myrcehyl");
      let _index355 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_2520_v13).length));
      (_2520_v13)[_index355] = _2521_v14;
      let _2522_v15;
      _2522_v15 = _dafny.Seq.of((_this).f27, (_this).f28);
      let _2523_v16;
      _2523_v16 = _dafny.Set.fromElements((_2522_v15)[_module.__default.safeIndex(_this.f29, new BigNumber((_2522_v15).length))]);
      let _2524_v17;
      _2524_v17 = _module.D0.create_DC2((_this).f28, new BigNumber((_2510_v3).length), !(_module.__default.fm2(globalState)));
      let _index356 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_2520_v13).length));
      let _rhs378 = ((_2518_v11).update((_dafny.ZERO).minus(_this.f29), _this.f29)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f29));
      let _rhs379 = _module.__default.fm57(globalState);
      let _rhs380 = (_2524_v17).dtor_cf3;
      let _rhs381 = _2521_v14;
      let _rhs382 = (_2523_v16).Intersect(_2523_v16);
      let _lhs291 = globalState;
      let _lhs292 = _2520_v13;
      let _lhs293 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_2520_v13).length));
      _2518_v11 = _rhs378;
      _2519_v12 = _rhs379;
      _lhs291.f7 = _rhs380;
      _lhs292[_lhs293] = _rhs381;
      _2523_v16 = _rhs382;
      r0 = (_this).f24;
      return r0;
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
