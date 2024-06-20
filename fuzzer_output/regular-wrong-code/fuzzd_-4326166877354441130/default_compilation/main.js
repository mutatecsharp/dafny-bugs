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
      let _source0 = _module.D1.create_DC2(!(false), !(!(false)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rhjaet")).length)), new BigNumber(748));
      if (_source0.is_DC2) {
        let _0___mcc_h0 = (_source0).cf2;
        let _1___mcc_h1 = (_source0).cf3;
        let _2___mcc_h2 = (_source0).cf4;
        let _3___mcc_h3 = (_source0).cf5;
        let _4_cf5 = _3___mcc_h3;
        let _5_cf4 = _2___mcc_h2;
        let _6_cf3 = _1___mcc_h1;
        let _7_cf2 = _0___mcc_h0;
        return new _dafny.CodePoint('c'.codePointAt(0));
      } else {
        let _8___mcc_h4 = (_source0).cf1;
        let _9_cf1 = _8___mcc_h4;
        return (_module.D2.create_DC3(new _dafny.CodePoint('g'.codePointAt(0)))).dtor_cf6;
      }
    };
    static fm1(p0, globalState) {
      return !(!_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-795)), _dafny.Seq.of(new BigNumber(950))), ((false) ? (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-194), new BigNumber(-576))).cardinality()),!(true))).length))) : (_dafny.Seq.of(new BigNumber(43), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("koxnyli")).length)))))));
    };
    static fm2(globalState) {
      return _dafny.Set.fromElements(new BigNumber(806), (new BigNumber(-179)).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(31), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-355)),new BigNumber(655))).length)))).length)));
    };
    static fm5(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(153)), function (_10_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("kfihj")).length);
      });
    };
    static fm6(p0, p1, p2, globalState) {
      let _source1 = _module.D1.create_DC1(new BigNumber((_dafny.Seq.UnicodeFromString("ywoabtk")).length));
      if (_source1.is_DC2) {
        let _11___mcc_h0 = (_source1).cf2;
        let _12___mcc_h1 = (_source1).cf3;
        let _13___mcc_h2 = (_source1).cf4;
        let _14___mcc_h3 = (_source1).cf5;
        let _15_cf5 = _14___mcc_h3;
        let _16_cf4 = _13___mcc_h2;
        let _17_cf3 = _12___mcc_h1;
        let _18_cf2 = _11___mcc_h0;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), function (_19_i0) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        });
      } else {
        let _20___mcc_h4 = (_source1).cf1;
        let _21_cf1 = _20___mcc_h4;
        return _dafny.Seq.UnicodeFromString("cncea");
      }
    };
    static fm7(p0, p1, p2, globalState) {
      let _source2 = _module.D2.create_DC3(new _dafny.CodePoint('u'.codePointAt(0)));
      if (_source2.is_DC4) {
        let _22___mcc_h0 = (_source2).cf7;
        let _23___mcc_h1 = (_source2).cf8;
        let _24___mcc_h2 = (_source2).cf9;
        let _25___mcc_h3 = (_source2).cf10;
        let _26_cf10 = _25___mcc_h3;
        let _27_cf9 = _24___mcc_h2;
        let _28_cf8 = _23___mcc_h1;
        let _29_cf7 = _22___mcc_h0;
        return _module.D1.create_DC2(_27_cf9, (_module.D2.create_DC4(_29_cf7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_28_cf8)).cardinality()),_28_cf8)).length), _27_cf9, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_27_cf9)).length)))).dtor_cf9, _28_cf8, new BigNumber(354));
      } else if (_source2.is_DC3) {
        let _30___mcc_h4 = (_source2).cf6;
        let _31_cf6 = _30___mcc_h4;
        return _module.D1.create_DC2(!(true), true, new BigNumber(883), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xep"),false)).length));
      } else {
        let _32___mcc_h5 = (_source2).cf11;
        let _33_cf11 = _32___mcc_h5;
        return _module.D1.create_DC2(!(false), !(true), new BigNumber((_dafny.Seq.UnicodeFromString("cijfftd")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length));
      }
    };
    static fm8(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_module.D2.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-419)),new BigNumber(920)), new BigNumber((_dafny.Seq.of(false, true)).length), false, new BigNumber(-917))).dtor_cf9, !(false), !(false), true)).cardinality())),_module.D1.create_DC2(true, true, new BigNumber(-480), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ttkwy")).length),new BigNumber(911))).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), function (_34_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length),_module.D1.create_DC2(false, true, new BigNumber(41), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length),true)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(637))).Elements) {
    let _35_v0 = _compr_0;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(637)), _35_v0)) {
      _coll0.push([(_35_v0).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(334))).length))).length)),new BigNumber(140)]);
    }
  }
  return _coll0;
}()).length))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), function (_36_i1) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})).length))).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), function (_37_i2) {
  return _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(37));
})).length))).length))))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(667), new BigNumber(900))) {
          let _38_v1 = _compr_1;
          if (((new BigNumber(667)).isLessThanOrEqualTo(_38_v1)) && ((_38_v1).isLessThan(new BigNumber(900)))) {
            _coll1.push([(_38_v1).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), function (_39_i3) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            })).length)),_module.D1.create_DC2(false, !(false), new BigNumber(836), new BigNumber((_dafny.Seq.UnicodeFromString("msq")).length))]);
          }
        }
        return _coll1;
      }());
    };
    static fm9(p0, p1, p2, globalState) {
      return (_dafny.ZERO).minus(function (_source3) {
        if (_source3.is_DC4) {
          let _40___mcc_h0 = (_source3).cf7;
          let _41___mcc_h1 = (_source3).cf8;
          let _42___mcc_h2 = (_source3).cf9;
          let _43___mcc_h3 = (_source3).cf10;
          let _44_cf10 = _43___mcc_h3;
          let _45_cf9 = _42___mcc_h2;
          let _46_cf8 = _41___mcc_h1;
          let _47_cf7 = _40___mcc_h0;
          return _44_cf10;
        } else if (_source3.is_DC3) {
          let _48___mcc_h4 = (_source3).cf6;
          let _49_cf6 = _48___mcc_h4;
          return (_dafny.ZERO).minus(((false) ? (new BigNumber((_dafny.Seq.UnicodeFromString("sjnmc")).length)) : (new BigNumber(-3))));
        } else {
          let _50___mcc_h5 = (_source3).cf11;
          let _51_cf11 = _50___mcc_h5;
          return new BigNumber(301);
        }
      }(_module.D2.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bururqtaf")).length), new BigNumber((_dafny.Seq.of(true)).length))).Elements) {
    let _52_v0 = _compr_2;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bururqtaf")).length), new BigNumber((_dafny.Seq.of(true)).length)), _52_v0)) {
      _coll2.push([(_52_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("g")).length)),new BigNumber(472)]);
    }
  }
  return _coll2;
}()).length)),new BigNumber((_dafny.Seq.UnicodeFromString("kabs")).length)), new BigNumber((_dafny.Seq.of(false)).length), false, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(915), new BigNumber(104))).cardinality()))).length))))));
    };
    static m0(p0, p1, globalState) {
      (globalState).f25 = new BigNumber(122);
      let _53_v0;
      _53_v0 = new _dafny.CodePoint('u'.codePointAt(0));
      let _54_v1;
      _54_v1 = _dafny.Seq.of(_53_v0);
      let _55_v2;
      _55_v2 = _dafny.Seq.of(_54_v1);
      let _56_v3;
      let _nw0 = Array((new BigNumber(3)).toNumber()).fill(false);
      _56_v3 = _nw0;
      let _57_v4;
      _57_v4 = _56_v3;
      let _58_v5;
      _58_v5 = _dafny.Seq.of(new BigNumber(241), new BigNumber((_dafny.Set.fromElements(new BigNumber((_54_v1).length))).length));
      let _59_v6;
      _59_v6 = _dafny.Map.Empty.slice().updateUnsafe(_57_v4,_58_v5);
      let _60_v7;
      let _nw1 = new _module.C0();
      _nw1.__ctor(_59_v6, p1);
      _60_v7 = _nw1;
      let _61_v8;
      _61_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Set.fromElements(_60_v7)).length));
      let _62_v9;
      let _nw2 = Array((new BigNumber(5)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = new BigNumber(((_55_v2)[_module.__default.safeIndex(p1, new BigNumber((_55_v2).length))]).length);
      _nw2[(_dafny.ONE).toNumber()] = new BigNumber(827);
      _nw2[(new BigNumber(2)).toNumber()] = p1;
      _nw2[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(969), (((_61_v8).contains((_60_v7).f29)) ? ((_61_v8).get((_60_v7).f29)) : ((_60_v7).f29)));
      _nw2[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(p1, (_60_v7).f29);
      _62_v9 = _nw2;
      _62_v9 = _62_v9;
      let _63_i0;
      _63_i0 = _dafny.ZERO;
      L0: {
        while (!(true)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_63_i0)) {
              break L0;
            }
            _63_i0 = (_63_i0).plus(_dafny.ONE);
            let _64_v10;
            _64_v10 = _dafny.MultiSet.fromElements((_60_v7).f29);
            (globalState).f2 = (((_64_v10).contains(new BigNumber(-85))) ? ((_64_v10).get(new BigNumber(-85))) : (p1));
            let _65_v11;
            _65_v11 = _dafny.Map.Empty.slice().updateUnsafe(_60_v7,_54_v1);
            let _66_v12;
            _66_v12 = true;
            _53_v0 = _module.__default.fm0(new BigNumber((((_65_v11).update(_60_v7, _54_v1)).Merge(_65_v11)).length), new BigNumber(490), _66_v12, (_dafny.ZERO).minus((_60_v7).f29), globalState);
            let _67_v13;
            _67_v13 = _dafny.Map.Empty.slice().updateUnsafe(_54_v1,_66_v12);
            if (((new BigNumber((_67_v13).length)).minus(p1)).isLessThanOrEqualTo((_60_v7).f29)) {
              (globalState).f16 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), ((_68_v0) => function (_69_i1) {
                return _68_v0;
              })(_53_v0)), _module.__default.safeIndex((_60_v7).f29, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), ((_70_v0) => function (_71_i1) {
                return _70_v0;
              })(_53_v0))).length)), _53_v0);
              let _72_v14;
              let _nw3 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
              _72_v14 = _nw3;
              _72_v14 = _72_v14;
              (globalState).f25 = ((_60_v7).fm3(globalState)).multipliedBy((_60_v7).f29);
              (globalState).f11 = _66_v12;
              (globalState).f2 = _module.__default.safeModuloInt((_60_v7).f29, (_60_v7).f29);
            } else {
              (globalState).f2 = ((_60_v7).f29).multipliedBy(p1);
              (globalState).f11 = ((_64_v10).Union(_64_v10)).IsSubsetOf(_64_v10);
              let _index0 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_62_v9).length));
              (_62_v9)[_index0] = p1;
              let _index1 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_62_v9).length));
              (_62_v9)[_index1] = (_60_v7).f29;
              (globalState).f4 = _66_v12;
              let _index2 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_56_v3).length));
              (_56_v3)[_index2] = _66_v12;
              let _73_v15;
              _73_v15 = _dafny.Set.fromElements((_60_v7).f29);
              let _74_v16;
              _74_v16 = _dafny.Set.fromElements(_66_v12, _66_v12);
              let _75_v17;
              _75_v17 = _dafny.MultiSet.fromElements(_74_v16);
              let _index3 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_56_v3).length));
              let _rhs0 = (_75_v17).IsDisjointFrom(((_75_v17).update(_74_v16, _module.__default.abs((_60_v7).f29))).Intersect(_75_v17));
              let _rhs1 = _73_v15;
              let _lhs0 = _56_v3;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_56_v3).length));
              _lhs0[_lhs1] = _rhs0;
              _73_v15 = _rhs1;
            }
            let _76_v18;
            _76_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_66_v12),(_60_v7).f29);
            let _77_v19;
            _77_v19 = _dafny.MultiSet.fromElements(_66_v12);
            (globalState).f2 = (((_76_v18).contains(_77_v19)) ? ((_76_v18).get(_77_v19)) : (((_60_v7).f29).multipliedBy(p1)));
          }
        }
      }
      let _78_v20;
      _78_v20 = false;
      (globalState).f4 = _78_v20;
      (globalState).f25 = (_60_v7).f29;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_62_v9).length))) {
        let _79_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_79_i2)) && ((_79_i2).isLessThan(new BigNumber((_62_v9).length))))) {
          (_62_v9)[(_79_i2)] = (_79_i2).multipliedBy((_dafny.ZERO).minus(_module.__default.safeModuloInt((_60_v7).f29, (_60_v7).f29)));
        }
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _80_v0;
      _80_v0 = true;
      let _81_v1;
      _81_v1 = _dafny.Seq.of(!(_80_v0), _80_v0);
      let _82_v2;
      _82_v2 = _dafny.Set.fromElements(_81_v1, _81_v1);
      let _83_v4;
      _83_v4 = new BigNumber(530);
      let _84_v5;
      _84_v5 = _dafny.Map.Empty.slice().updateUnsafe(!(_80_v0),_80_v0);
      let _85_v6;
      _85_v6 = _dafny.Seq.of(_83_v4, _83_v4, new BigNumber((_84_v5).length));
      let _86_v7;
      let _init0 = function (_87_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      };
      let _nw4 = Array((new BigNumber(7)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw4.length); _i0_0++) {
        _nw4[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _86_v7 = _nw4;
      let _88_v9;
      _88_v9 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.of(_83_v4)).Elements) {
          let _89_v8 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_83_v4), _89_v8)) {
            _coll3.push([_module.__default.safeDivisionInt(_89_v8, new BigNumber((_85_v6).length)),_80_v0]);
          }
        }
        return _coll3;
      }(),_80_v0);
      let _90_v10;
      _90_v10 = _dafny.Map.Empty.slice().updateUnsafe((_85_v6)[_module.__default.safeIndex(_83_v4, new BigNumber((_85_v6).length))],_80_v0);
      let _91_v11;
      _91_v11 = _dafny.MultiSet.fromElements((((_88_v9).contains(_90_v10)) ? ((_88_v9).get(_90_v10)) : (_80_v0)), _80_v0, (_81_v1)[_module.__default.safeIndex(_83_v4, new BigNumber((_81_v1).length))], _80_v0, _80_v0);
      let _92_v12;
      _92_v12 = _dafny.Seq.UnicodeFromString("ic");
      let _93_v13;
      _93_v13 = _dafny.Set.fromElements(_83_v4, _83_v4, _83_v4, _83_v4);
      let _94_v14;
      let _nw5 = Array((new BigNumber(10)).toNumber());
      _nw5[(_dafny.ZERO).toNumber()] = new BigNumber(-78);
      _nw5[(_dafny.ONE).toNumber()] = _83_v4;
      _nw5[(new BigNumber(2)).toNumber()] = new BigNumber((_93_v13).length);
      _nw5[(new BigNumber(3)).toNumber()] = _83_v4;
      _nw5[(new BigNumber(4)).toNumber()] = _83_v4;
      _nw5[(new BigNumber(5)).toNumber()] = _83_v4;
      _nw5[(new BigNumber(6)).toNumber()] = new BigNumber(687);
      _nw5[(new BigNumber(7)).toNumber()] = _83_v4;
      _nw5[(new BigNumber(8)).toNumber()] = _83_v4;
      _nw5[(new BigNumber(9)).toNumber()] = _83_v4;
      _94_v14 = _nw5;
      let _95_v15;
      _95_v15 = _dafny.Map.Empty.slice().updateUnsafe(_83_v4,_92_v12);
      let _96_globalState;
      let _nw6 = new _module.GlobalState();
      _nw6.__ctor(new BigNumber(-129), _82_v2, new BigNumber(512), new BigNumber(90), false, false, new BigNumber(916), true, new BigNumber(913), _dafny.Seq.Create(_module.__default.abs(new BigNumber(473)), function (_97_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624));
      }), new BigNumber(-876), false, function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_85_v6).Elements) {
          let _98_v3 = _compr_4;
          if (_dafny.Seq.contains(_85_v6, _98_v3)) {
            _coll4.push([_module.__default.safeModuloInt(_98_v3, (_dafny.ZERO).minus(_83_v4)),false]);
          }
        }
        return _coll4;
      }(), _86_v7, new BigNumber(359), (_dafny.MultiSet.FromArray(_81_v1)).Union(_91_v11), _dafny.Seq.Concat(_92_v12, _92_v12), false, _93_v13, false, _86_v7, _dafny.MultiSet.fromElements(_93_v13), true, _94_v14, true, new BigNumber(990), new BigNumber(255), _95_v15);
      _96_globalState = _nw6;
      let _99_v16;
      _99_v16 = new _dafny.CodePoint('f'.codePointAt(0));
      let _100_v17;
      _100_v17 = _dafny.MultiSet.fromElements(((_80_v0) ? (_99_v16) : (_99_v16)), _module.__default.fm0(_83_v4, new BigNumber((_81_v1).length), !(_80_v0), _83_v4, _96_globalState));
      _100_v17 = (_100_v17).Intersect(_dafny.MultiSet.fromElements(_99_v16));
      (_96_globalState).f4 = _80_v0;
      (_96_globalState).f11 = _module.__default.fm1((_83_v4).multipliedBy(new BigNumber(282)), _96_globalState);
      let _101_v18;
      let _init1 = ((_102_v13) => function (_103_i2) {
        return !((_102_v13).IsSubsetOf(_102_v13));
      })(_93_v13);
      let _nw7 = Array((new BigNumber(5)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
        _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _101_v18 = _nw7;
      let _index4 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
      (_101_v18)[_index4] = (_81_v1)[_module.__default.safeIndex(_83_v4, new BigNumber((_81_v1).length))];
      let _index5 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
      (_101_v18)[_index5] = _80_v0;
      let _104_v19;
      _104_v19 = _dafny.Map.Empty.slice().updateUnsafe(_83_v4,_83_v4);
      (_96_globalState).f25 = new BigNumber((((_104_v19).Merge(_104_v19)).Merge(_104_v19)).length);
      _101_v18 = _101_v18;
      (_96_globalState).f2 = (_83_v4).plus((_83_v4).plus(_83_v4));
      let _rhs2 = false;
      let _rhs3 = _80_v0;
      let _rhs4 = _94_v14;
      let _rhs5 = ((_84_v5).update(_80_v0, _80_v0)).Merge(_84_v5);
      let _lhs2 = _96_globalState;
      _lhs2.f4 = _rhs2;
      _80_v0 = _rhs3;
      _94_v14 = _rhs4;
      _84_v5 = _rhs5;
      let _index6 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
      (_101_v18)[_index6] = ((_81_v1)[_module.__default.safeIndex(new BigNumber((_81_v1).length), new BigNumber((_81_v1).length))]) || ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_80_v0,(_dafny.ZERO).minus(_83_v4))).length)).isLessThan(new BigNumber((_90_v10).length)));
      let _index7 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
      (_94_v14)[_index7] = _83_v4;
      let _index8 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
      (_94_v14)[_index8] = (((_104_v19).contains(_83_v4)) ? ((_104_v19).get(_83_v4)) : (_83_v4));
      if ((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]) {
        let _105_v20;
        _105_v20 = _dafny.Map.Empty.slice().updateUnsafe((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))],_101_v18);
        let _106_v21;
        _106_v21 = _101_v18;
        let _107_v22;
        let _nw8 = Array((new BigNumber(26)).toNumber());
        _nw8[(_dafny.ZERO).toNumber()] = _101_v18;
        _nw8[(_dafny.ONE).toNumber()] = _101_v18;
        _nw8[(new BigNumber(2)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(3)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(4)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(5)).toNumber()] = (((_105_v20).contains((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))])) ? ((_105_v20).get((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))])) : (_101_v18));
        _nw8[(new BigNumber(6)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(7)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(8)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(9)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(10)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(11)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(12)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(13)).toNumber()] = (_101_v18);
        _nw8[(new BigNumber(14)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(15)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(16)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(17)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(18)).toNumber()] = (_106_v21);
        _nw8[(new BigNumber(19)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(20)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(21)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(22)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(23)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(24)).toNumber()] = _101_v18;
        _nw8[(new BigNumber(25)).toNumber()] = _101_v18;
        _107_v22 = _nw8;
        let _index9 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_107_v22).length));
        (_107_v22)[_index9] = _101_v18;
        let _index10 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_107_v22).length));
        (_107_v22)[_index10] = _101_v18;
        let _108_v24;
        _108_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_83_v4)),function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(820), new BigNumber(542))) {
            let _109_v23 = _compr_5;
            if (((new BigNumber(820)).isLessThanOrEqualTo(_109_v23)) && ((_109_v23).isLessThan(new BigNumber(542)))) {
              _coll5.add((_109_v23).multipliedBy(_83_v4));
            }
          }
          return _coll5;
        }());
        if (((((_dafny.ZERO).minus(new BigNumber(((((_108_v24).contains(_83_v4)) ? ((_108_v24).get(_83_v4)) : (_module.__default.fm2(_96_globalState)))).length))).isLessThanOrEqualTo((_dafny.ZERO).minus((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]))) ? (_80_v0) : (_module.__default.fm1(_83_v4, _96_globalState)))) {
          (_96_globalState).f11 = !((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]);
          _99_v16 = _99_v16;
          let _110_v25;
          _110_v25 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("tvycemat"));
          _module.__default.m0(_110_v25, _83_v4, _96_globalState);
          let _111_v26;
          _111_v26 = _dafny.MultiSet.fromElements((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
          _104_v19 = (_104_v19).update(new BigNumber((_111_v26).cardinality()), _module.__default.safeModuloInt(_83_v4, _83_v4));
          let _index11 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
          (_94_v14)[_index11] = _83_v4;
        } else {
          let _112_v27;
          _112_v27 = _dafny.Map.Empty.slice().updateUnsafe((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))],((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]).plus(new BigNumber((_81_v1).length)));
          _112_v27 = (_112_v27).update(_80_v0, (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
          let _113_v28;
          _113_v28 = _dafny.Map.Empty.slice().updateUnsafe(_106_v21,_85_v6);
          let _114_v29;
          _114_v29 = _module.D1.create_DC1(_83_v4);
          let _115_v30;
          let _nw9 = new _module.C0();
          _nw9.__ctor(_113_v28, (_114_v29).dtor_cf1);
          _115_v30 = _nw9;
          _99_v16 = _99_v16;
          let _116_v31;
          let _nw10 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
          _116_v31 = _nw10;
          _116_v31 = _116_v31;
          let _117_v32;
          _117_v32 = _dafny.Map.Empty.slice().updateUnsafe(_99_v16,_80_v0);
          _117_v32 = (_117_v32).update(_99_v16, _80_v0);
        }
        if ((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]) {
          (_96_globalState).f11 = (_80_v0) === (((_80_v0) ? ((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]) : (_80_v0)));
          let _118_v33;
          let _nw11 = Array((new BigNumber(29)).toNumber());
          _118_v33 = _nw11;
          let _119_v34;
          _119_v34 = _dafny.Map.Empty.slice().updateUnsafe(_106_v21,_dafny.Seq.update(_dafny.Seq.of(new BigNumber(681), _83_v4), _module.__default.safeIndex((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], new BigNumber((_dafny.Seq.of(new BigNumber(681), _83_v4)).length)), new BigNumber((_dafny.Seq.UnicodeFromString("am")).length)));
          let _120_v35;
          let _nw12 = new _module.C0();
          _nw12.__ctor(_119_v34, (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
          _120_v35 = _nw12;
          let _index12 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_118_v33).length));
          (_118_v33)[_index12] = _120_v35;
          let _index13 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_118_v33).length));
          let _rhs6 = _120_v35;
          let _rhs7 = _80_v0;
          let _lhs3 = _118_v33;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_118_v33).length));
          let _lhs5 = _96_globalState;
          _lhs3[_lhs4] = _rhs6;
          _lhs5.f4 = _rhs7;
          _99_v16 = _module.__default.fm0((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], (_120_v35).fm3(_96_globalState), (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], (_120_v35).f29, _96_globalState);
          let _121_v36;
          _121_v36 = _dafny.Seq.of(_92_v12, _dafny.Seq.UnicodeFromString("caxv"));
          let _122_v37;
          let _nw13 = new _module.C0();
          _nw13.__ctor((_120_v35).f28, (((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]) ? (new BigNumber((_81_v1).length)) : (new BigNumber(((_121_v36)[_module.__default.safeIndex(_83_v4, new BigNumber((_121_v36).length))]).length))));
          _122_v37 = _nw13;
          (_96_globalState).f25 = (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))];
        } else {
          let _index14 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
          (_94_v14)[_index14] = (_83_v4).multipliedBy((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(776), new BigNumber(827))) {
              let _123_v38 = _compr_6;
              if (((new BigNumber(776)).isLessThanOrEqualTo(_123_v38)) && ((_123_v38).isLessThan(new BigNumber(827)))) {
                _coll6.add(_module.__default.safeDivisionInt(_123_v38, new BigNumber((_85_v6).length)));
              }
            }
            return _coll6;
          }()).length)));
          _107_v22 = _107_v22;
          let _124_v39;
          _124_v39 = _dafny.Seq.of(_94_v14, _94_v14);
          let _125_v40;
          _125_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,_92_v12);
          let _126_v41;
          _126_v41 = _dafny.MultiSet.fromElements(_92_v12, _dafny.Seq.update((((_125_v40).contains(_80_v0)) ? ((_125_v40).get(_80_v0)) : (_92_v12)), _module.__default.safeIndex((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], new BigNumber(((((_125_v40).contains(_80_v0)) ? ((_125_v40).get(_80_v0)) : (_92_v12))).length)), _99_v16));
          let _rhs8 = (_124_v39)[_module.__default.safeIndex((_85_v6)[_module.__default.safeIndex((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], new BigNumber((_85_v6).length))], new BigNumber((_124_v39).length))];
          let _rhs9 = !((_126_v41).contains(_92_v12));
          let _rhs10 = _94_v14;
          let _rhs11 = (_85_v6)[_module.__default.safeIndex((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], new BigNumber((_85_v6).length))];
          let _rhs12 = _81_v1;
          let _lhs6 = _96_globalState;
          let _lhs7 = _96_globalState;
          _94_v14 = _rhs8;
          _lhs6.f11 = _rhs9;
          _94_v14 = _rhs10;
          _lhs7.f8 = _rhs11;
          _81_v1 = _rhs12;
          let _127_v42;
          _127_v42 = _dafny.Map.Empty.slice().updateUnsafe(_106_v21,_85_v6);
          let _128_v43;
          let _nw14 = new _module.C0();
          _nw14.__ctor(((_127_v42).update(_106_v21, _85_v6)).Merge(_127_v42), _83_v4);
          _128_v43 = _nw14;
          let _129_v44;
          _129_v44 = _module.D1.create_DC1((_85_v6)[_module.__default.safeIndex((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], new BigNumber((_85_v6).length))]);
          _129_v44 = _129_v44;
        }
        _99_v16 = _99_v16;
        let _130_v45;
        _130_v45 = _module.D1.create_DC2(_80_v0, !((((_84_v5).contains((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))])) ? ((_84_v5).get((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))])) : (_80_v0))), _83_v4, _83_v4);
        let _pat_let_tv0 = _80_v0;
        let _pat_let_tv1 = _85_v6;
        let _pat_let_tv2 = _83_v4;
        let _pat_let_tv3 = _85_v6;
        let _pat_let_tv4 = _101_v18;
        let _pat_let_tv5 = _101_v18;
        let _131_v46;
        let _nw15 = Array((new BigNumber(13)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _130_v45;
        _nw15[(_dafny.ONE).toNumber()] = _130_v45;
        _nw15[(new BigNumber(2)).toNumber()] = _130_v45;
        _nw15[(new BigNumber(3)).toNumber()] = _130_v45;
        _nw15[(new BigNumber(4)).toNumber()] = _module.D1.create_DC2(_80_v0, (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _83_v4, _83_v4);
        _nw15[(new BigNumber(5)).toNumber()] = _130_v45;
        _nw15[(new BigNumber(6)).toNumber()] = function (_pat_let0_0) {
          return function (_132_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_133_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_133_dt__update_hcf2_h0, (_132_dt__update__tmp_h0).dtor_cf3, (_132_dt__update__tmp_h0).dtor_cf4, (_132_dt__update__tmp_h0).dtor_cf5);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_130_v45);
        _nw15[(new BigNumber(7)).toNumber()] = _module.D1.create_DC2((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
        _nw15[(new BigNumber(8)).toNumber()] = _130_v45;
        _nw15[(new BigNumber(9)).toNumber()] = function (_pat_let2_0) {
          return function (_134_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_135_dt__update_hcf4_h0) {
                return function (_pat_let4_0) {
                  return function (_136_dt__update_hcf3_h0) {
                    return _module.D1.create_DC2((_134_dt__update__tmp_h1).dtor_cf2, _136_dt__update_hcf3_h0, _135_dt__update_hcf4_h0, (_134_dt__update__tmp_h1).dtor_cf5);
                  }(_pat_let4_0);
                }((_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_pat_let_tv4).length))]);
              }(_pat_let3_0);
            }((_pat_let_tv1)[_module.__default.safeIndex(_pat_let_tv2, new BigNumber((_pat_let_tv3).length))]);
          }(_pat_let2_0);
        }(_130_v45);
        _nw15[(new BigNumber(10)).toNumber()] = _130_v45;
        _nw15[(new BigNumber(11)).toNumber()] = _130_v45;
        _nw15[(new BigNumber(12)).toNumber()] = _130_v45;
        _131_v46 = _nw15;
        let _index15 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_131_v46).length));
        (_131_v46)[_index15] = _130_v45;
        let _137_v47;
        _137_v47 = _dafny.MultiSet.fromElements(_module.D1.create_DC2((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _80_v0, (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], new BigNumber((_85_v6).length)));
        let _index16 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
        let _index17 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_131_v46).length));
        let _index18 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
        let _rhs13 = (_137_v47).IsSubsetOf(_137_v47);
        let _rhs14 = _module.D1.create_DC2(!(_83_v4).isEqualTo((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]), _80_v0, _83_v4, (_83_v4).multipliedBy(_83_v4));
        let _rhs15 = (true) && (!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ruxv"), new _dafny.CodePoint('p'.codePointAt(0))));
        let _rhs16 = (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))];
        let _rhs17 = ((_80_v0) ? (false) : ((_80_v0) && (!((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]))));
        let _lhs8 = _101_v18;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
        let _lhs10 = _131_v46;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_131_v46).length));
        let _lhs12 = _96_globalState;
        let _lhs13 = _96_globalState;
        let _lhs14 = _101_v18;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
        _lhs8[_lhs9] = _rhs13;
        _lhs10[_lhs11] = _rhs14;
        _lhs12.f4 = _rhs15;
        _lhs13.f4 = _rhs16;
        _lhs14[_lhs15] = _rhs17;
      } else {
        let _index19 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
        let _rhs18 = (_dafny.ZERO).minus((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
        let _rhs19 = (_dafny.ZERO).minus((new BigNumber(16)).plus(_module.__default.safeDivisionInt(new BigNumber((_84_v5).length), (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))])));
        let _rhs20 = _81_v1;
        let _rhs21 = _module.__default.safeModuloInt(new BigNumber(((_90_v10).Merge(_90_v10)).length), (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
        let _lhs16 = _94_v14;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
        let _lhs18 = _96_globalState;
        let _lhs19 = _96_globalState;
        _lhs16[_lhs17] = _rhs18;
        _lhs18.f25 = _rhs19;
        _81_v1 = _rhs20;
        _lhs19.f0 = _rhs21;
        if (_dafny.Seq.IsProperPrefixOf(_92_v12, _92_v12)) {
          (_96_globalState).f11 = ((_dafny.MultiSet.FromArray(_81_v1)).update((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _module.__default.abs(new BigNumber(-193)))).IsSubsetOf(_91_v11);
          let _138_v48;
          _138_v48 = _101_v18;
          let _139_v49;
          _139_v49 = _dafny.Map.Empty.slice().updateUnsafe(_138_v48,_85_v6);
          let _140_v50;
          _140_v50 = _dafny.MultiSet.fromElements(new BigNumber((_92_v12).length), _83_v4, _83_v4);
          let _141_v51;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_139_v49, (((_104_v19).contains(new BigNumber((_140_v50).cardinality()))) ? ((_104_v19).get(new BigNumber((_140_v50).cardinality()))) : (_83_v4)));
          _141_v51 = _nw16;
          let _nw17 = new _module.C0();
          _nw17.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_138_v48,_85_v6), (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
          _141_v51 = _nw17;
          let _index20 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length));
          (_94_v14)[_index20] = _module.__default.safeDivisionInt((_141_v51).f29, new BigNumber((_93_v13).length));
          _104_v19 = (_104_v19).update(new BigNumber((((_80_v0) ? (_81_v1) : (_81_v1))).length), _83_v4);
          let _142_v52;
          let _init2 = function (_143_i3) {
            return _dafny.Seq.UnicodeFromString("ycg");
          };
          let _nw18 = Array((new BigNumber(28)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw18.length); _i0_2++) {
            _nw18[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _142_v52 = _nw18;
          let _144_v53;
          _144_v53 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,_142_v52);
          _144_v53 = (_144_v53).update((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _142_v52);
        } else {
          let _index21 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
          (_101_v18)[_index21] = _dafny.Seq.IsProperPrefixOf(_92_v12, _92_v12);
          _94_v14 = _94_v14;
          let _145_v54;
          _145_v54 = _dafny.Map.Empty.slice().updateUnsafe(_101_v18,_dafny.Seq.of(_83_v4));
          let _146_v55;
          let _nw19 = new _module.C0();
          _nw19.__ctor((_145_v54).Merge(_145_v54), (_83_v4).plus((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]));
          _146_v55 = _nw19;
          (_96_globalState).f4 = _module.__default.fm1(_83_v4, _96_globalState);
          (_96_globalState).f11 = (_81_v1)[_module.__default.safeIndex((new BigNumber((_92_v12).length)).multipliedBy((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]), new BigNumber((_81_v1).length))];
        }
        let _147_v56;
        let _nw20 = Array((new BigNumber(19)).toNumber());
        _147_v56 = _nw20;
        let _148_v57;
        _148_v57 = _101_v18;
        let _149_v58;
        _149_v58 = _dafny.Map.Empty.slice().updateUnsafe(_148_v57,_85_v6);
        let _150_v59;
        let _nw21 = new _module.C0();
        _nw21.__ctor((_149_v58).update(_148_v57, _module.__default.fm5(_80_v0, _99_v16, _96_globalState)), (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
        _150_v59 = _nw21;
        let _index22 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_147_v56).length));
        (_147_v56)[_index22] = _150_v59;
        let _index23 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_147_v56).length));
        let _nw22 = new _module.C0();
        _nw22.__ctor((_150_v59).f28, new BigNumber((_92_v12).length));
        (_147_v56)[_index23] = _nw22;
        let _151_v60;
        _151_v60 = _dafny.Map.Empty.slice().updateUnsafe(_99_v16,_80_v0);
        _151_v60 = (_151_v60).update(_99_v16, false);
        _94_v14 = _94_v14;
      }
      let _hi0 = (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))];
      for (let _152_i4 = ((false) ? ((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]) : (_83_v4)); _152_i4.isLessThan(_hi0); _152_i4 = _152_i4.plus(_dafny.ONE)) {
        (_96_globalState).f16 = _module.__default.fm6(_80_v0, _80_v0, new BigNumber(((_84_v5).update(_module.__default.fm1(new BigNumber((_81_v1).length), _96_globalState), _80_v0)).length), _96_globalState);
        _101_v18 = _101_v18;
        let _153_v61;
        _153_v61 = _dafny.Set.fromElements(_101_v18);
        _153_v61 = _153_v61;
        let _154_v62;
        _154_v62 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_155_v16) => function (_156_i5) {
          return _155_v16;
        })(_99_v16)), _92_v12);
        if ((_154_v62).IsSubsetOf(_154_v62)) {
          _module.__default.m0(_154_v62, _module.__default.safeModuloInt(_152_i4, _152_i4), _96_globalState);
          let _157_v63;
          _157_v63 = _101_v18;
          let _158_v64;
          _158_v64 = _dafny.Map.Empty.slice().updateUnsafe(_157_v63,_85_v6);
          let _159_v65;
          _159_v65 = _dafny.Seq.of(_158_v64);
          let _160_v66;
          _160_v66 = _dafny.Seq.of((_159_v65)[_module.__default.safeIndex(_83_v4, new BigNumber((_159_v65).length))]);
          let _161_v67;
          let _nw23 = new _module.C0();
          _nw23.__ctor((_160_v66)[_module.__default.safeIndex(_152_i4, new BigNumber((_160_v66).length))], _83_v4);
          _161_v67 = _nw23;
          _161_v67 = _161_v67;
          _module.__default.m0(_154_v62, (_161_v67).f29, _96_globalState);
          let _162_v68;
          _162_v68 = _dafny.Set.fromElements((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _80_v0);
          _162_v68 = (_162_v68).Union(_162_v68);
          let _163_v69;
          _163_v69 = _dafny.Map.Empty.slice().updateUnsafe(_152_i4,_module.__default.fm7((_161_v67).f29, (_161_v67).f29, _92_v12, _96_globalState));
          let _164_v70;
          _164_v70 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,_163_v69);
          let _165_v71;
          _165_v71 = _dafny.Seq.of((((_164_v70).contains(!((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]))) ? ((_164_v70).get(!((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]))) : (_163_v69)), _module.__default.fm8((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _83_v4, _96_globalState));
          let _rhs22 = !_dafny.areEqual(_dafny.Seq.Concat(_92_v12, _92_v12), _dafny.Seq.Concat(_92_v12, _92_v12));
          let _rhs23 = (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))];
          let _rhs24 = (_165_v71)[_module.__default.safeIndex((_dafny.ZERO).minus(_83_v4), new BigNumber((_165_v71).length))];
          let _rhs25 = _92_v12;
          let _lhs20 = _96_globalState;
          let _lhs21 = _96_globalState;
          let _lhs22 = _96_globalState;
          _lhs20.f4 = _rhs22;
          _lhs21.f4 = _rhs23;
          _163_v69 = _rhs24;
          _lhs22.f16 = _rhs25;
        } else {
          let _166_v72;
          _166_v72 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))],_80_v0));
          let _167_v73;
          _167_v73 = _dafny.MultiSet.fromElements(new BigNumber((_81_v1).length));
          _104_v19 = (_104_v19).update((((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]) ? (new BigNumber((_81_v1).length)) : (_module.__default.fm9((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], _166_v72, _152_i4, _96_globalState))), new BigNumber(((_dafny.MultiSet.fromElements(_83_v4, _83_v4, _83_v4)).Union(_167_v73)).cardinality()));
          let _168_v74;
          _168_v74 = _101_v18;
          let _169_v75;
          _169_v75 = _dafny.Map.Empty.slice().updateUnsafe(_168_v74,_dafny.Seq.of(_152_i4));
          let _170_v76;
          let _nw24 = new _module.C0();
          _nw24.__ctor((_169_v75).update(_168_v74, _85_v6), (new BigNumber(222)).plus(_83_v4));
          _170_v76 = _nw24;
          _170_v76 = _170_v76;
          _80_v0 = true;
          _83_v4 = _module.__default.safeModuloInt(_152_i4, (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))]);
          let _rhs26 = (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))];
          let _rhs27 = _152_i4;
          let _lhs23 = _96_globalState;
          let _lhs24 = _96_globalState;
          _lhs23.f11 = _rhs26;
          _lhs24.f25 = _rhs27;
        }
      }
      let _171_v77;
      _171_v77 = _101_v18;
      let _172_v78;
      _172_v78 = _dafny.Map.Empty.slice().updateUnsafe(_171_v77,_85_v6);
      let _173_v79;
      let _nw25 = new _module.C0();
      _nw25.__ctor(_172_v78, _83_v4);
      _173_v79 = _nw25;
      let _174_v80;
      _174_v80 = _dafny.Seq.of(_173_v79);
      let _175_v81;
      _175_v81 = _module.D1.create_DC2(_80_v0, (_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))], _83_v4, new BigNumber((_174_v80).length));
      let _176_v82;
      _176_v82 = _dafny.Seq.of(_module.__default.fm7((_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], (_94_v14)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_94_v14).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), ((_177_v16) => function (_178_i6) {
        return _177_v16;
      })(_99_v16)), _96_globalState), _175_v81, _175_v81);
      let _179_v83;
      let _nw26 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
      _179_v83 = _nw26;
      let _pat_let_tv6 = _83_v4;
      let _pat_let_tv7 = _80_v0;
      let _index24 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
      let _rhs28 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_92_v12, _module.__default.safeIndex(new BigNumber(638), new BigNumber((_92_v12).length)), _99_v16)).length));
      let _rhs29 = _dafny.Seq.of(function (_pat_let5_0) {
        return function (_180_dt__update__tmp_h2) {
          return function (_pat_let6_0) {
            return function (_181_dt__update_hcf5_h0) {
              return _module.D1.create_DC2((_180_dt__update__tmp_h2).dtor_cf2, (_180_dt__update__tmp_h2).dtor_cf3, (_180_dt__update__tmp_h2).dtor_cf4, _181_dt__update_hcf5_h0);
            }(_pat_let6_0);
          }(_pat_let_tv6);
        }(_pat_let5_0);
      }(_module.D1.create_DC2(_80_v0, !(_80_v0), (_173_v79).f29, _83_v4)));
      let _rhs30 = function (_source4) {
        if (_source4.is_DC2) {
          let _182___mcc_h0 = (_source4).cf2;
          let _183___mcc_h1 = (_source4).cf3;
          let _184___mcc_h2 = (_source4).cf4;
          let _185___mcc_h3 = (_source4).cf5;
          let _186_cf5 = _185___mcc_h3;
          let _187_cf4 = _184___mcc_h2;
          let _188_cf3 = _183___mcc_h1;
          let _189_cf2 = _182___mcc_h0;
          return _188_cf3;
        } else {
          let _190___mcc_h4 = (_source4).cf1;
          let _191_cf1 = _190___mcc_h4;
          return (_pat_let_tv7) && (false);
        }
      }(_175_v81);
      let _rhs31 = !((_101_v18)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length))]);
      let _rhs32 = _179_v83;
      let _lhs25 = _101_v18;
      let _lhs26 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_101_v18).length));
      let _lhs27 = _96_globalState;
      _93_v13 = _rhs28;
      _176_v82 = _rhs29;
      _lhs25[_lhs26] = _rhs30;
      _lhs27.f11 = _rhs31;
      _179_v83 = _rhs32;
      let _rhs33 = _80_v0;
      let _rhs34 = ((((_module.__default.fm1(new BigNumber((_92_v12).length), _96_globalState)) ? (_80_v0) : (_80_v0))) ? (_dafny.Seq.Concat(_92_v12, _92_v12)) : (_92_v12));
      let _rhs35 = (_dafny.ZERO).minus((_173_v79).f29);
      let _rhs36 = _99_v16;
      let _lhs28 = _96_globalState;
      let _lhs29 = _96_globalState;
      _lhs28.f4 = _rhs33;
      _92_v12 = _rhs34;
      _lhs29.f8 = _rhs35;
      _99_v16 = _rhs36;
      let _192_v84;
      _192_v84 = _dafny.Map.Empty.slice().updateUnsafe(_99_v16,_92_v12);
      (_96_globalState).f16 = (((_192_v84).contains(_99_v16)) ? ((_192_v84).get(_99_v16)) : (_92_v12));
      (_96_globalState).f4 = false;
      process.stdout.write(_dafny.toString(_80_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_81_v1, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_82_v2).equals(_dafny.Set.fromElements(_dafny.Seq.of(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_83_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_84_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true).updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_85_v6, _dafny.Seq.of(new BigNumber(530), new BigNumber(530), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(176),true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v11).equals(_dafny.MultiSet.fromElements(true, true, true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_92_v12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_v13).equals(_dafny.Set.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v14)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_95_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(530),_dafny.Seq.UnicodeFromString("ic")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f1).equals(_dafny.Set.fromElements(_dafny.Seq.of(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_96_globalState.f9, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(624))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false).updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f15).equals(_dafny.MultiSet.fromElements(false, false, true, true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_96_globalState.f16).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f18).equals(_dafny.Set.fromElements(new BigNumber(530)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f20)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f21).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(530))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_96_globalState).f23)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_96_globalState.f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_globalState.f27).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(530),_dafny.Seq.UnicodeFromString("ic")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_99_v16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v17).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v18)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v18)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v18)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v18)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v18)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(530),new BigNumber(530)).updateUnsafe(new BigNumber(2),new BigNumber(530)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_171_v77))[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_171_v77))[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_171_v77))[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_171_v77))[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_171_v77))[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_172_v78).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_173_v79).f28).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_173_v79).f29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_174_v80).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v81).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v81).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v81).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v81).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_176_v82, _dafny.Seq.of(_module.D1.create_DC2(true, false, new BigNumber(530), new BigNumber(530))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v84).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_dafny.Seq.UnicodeFromString("icic")))));
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
      return [];
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
    static create_DC2(cf2, cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
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
    get dtor_cf5() { return this.cf5; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
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
        return other.$tag === 0 && this.cf2 === other.cf2 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(false, false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC4(cf7, cf8, cf9, cf10) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf11) {
      let $dt = new D2(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.Map.Empty, _dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f2 = _dafny.ZERO;
      this.f4 = false;
      this.f8 = _dafny.ZERO;
      this.f9 = _dafny.Seq.of();
      this.f11 = false;
      this.f12 = _dafny.Map.Empty;
      this.f14 = _dafny.ZERO;
      this.f15 = _dafny.MultiSet.Empty;
      this.f16 = _dafny.Seq.UnicodeFromString("");
      this.f20 = [];
      this.f21 = _dafny.MultiSet.Empty;
      this.f25 = _dafny.ZERO;
      this.f27 = _dafny.Map.Empty;
      this._f1 = _dafny.Set.Empty;
      this._f3 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f10 = _dafny.ZERO;
      this._f13 = [];
      this._f17 = false;
      this._f18 = _dafny.Set.Empty;
      this._f19 = false;
      this._f22 = false;
      this._f23 = [];
      this._f24 = false;
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this).f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this).f25 = f25;
      (_this)._f26 = f26;
      (_this).f27 = f27;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f13() {
      let _this = this;
      return _this._f13;
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
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f28 = _dafny.Map.Empty;
      this._f29 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f28, f29) {
      let _this = this;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f29));
    };
    fm4(globalState) {
      let _this = this;
      return !_dafny.Seq.contains(_dafny.Seq.of(!(!(!(false))), false), !((function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(869), new BigNumber(196))) {
          let _193_v0 = _compr_7;
          if (((new BigNumber(869)).isLessThanOrEqualTo(_193_v0)) && ((_193_v0).isLessThan(new BigNumber(196)))) {
            _coll7.add(_module.__default.safeDivisionInt(_193_v0, (_this).f29));
          }
        }
        return _coll7;
      }()).IsProperSubsetOf(_dafny.Set.fromElements((_this).f29))));
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
