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
      return ((new BigNumber(916)).plus(new BigNumber(948))).plus(_module.__default.safeDivisionInt(new BigNumber(405), new BigNumber(577)));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sgstncg"),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), function (_0_i0) {
          return new BigNumber((function () {
            let _coll1 = new _dafny.Set();
            for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), function (_1_i1) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(922))).length),false)).length);
            })).Elements) {
              let _2_v1 = _compr_1;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), function (_1_i1) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(922))).length),false)).length);
              }), _2_v1)) {
                _coll1.add(_module.__default.safeDivisionInt(_2_v1, new BigNumber(866)));
              }
            }
            return _coll1;
          }()).length);
        })).length))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(700),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(true, false))).length))))).Keys.Elements) {
          let _3_v0 = _compr_0;
          if (((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sgstncg"),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), function (_0_i0) {
            return new BigNumber((function () {
              let _coll2 = new _dafny.Set();
              for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), function (_1_i1) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(922))).length),false)).length);
              })).Elements) {
                let _4_v1 = _compr_2;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), function (_1_i1) {
                  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(922))).length),false)).length);
                }), _4_v1)) {
                  _coll2.add(_module.__default.safeDivisionInt(_4_v1, new BigNumber(866)));
                }
              }
              return _coll2;
            }()).length);
          })).length))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(700),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(true, false))).length))))).contains(_3_v0)) {
            _coll0.push([(_3_v0).minus((_dafny.ZERO).minus(new BigNumber(-128))),(new BigNumber(10)).plus(new BigNumber(62))]);
          }
        }
        return _coll0;
      }();
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jwdka"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(877)), function (_5_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("mgogecpi"));
    };
    static fm5(globalState) {
      if (true) {
        return function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-923), new BigNumber(295))) {
            let _6_v0 = _compr_3;
            if (((new BigNumber(-923)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(295)))) {
              _coll3.add((_6_v0).multipliedBy(new BigNumber(595)));
            }
          }
          return _coll3;
        }();
      } else {
        return (_dafny.Set.fromElements(new BigNumber(-878), new BigNumber(28))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("yycduo")).length)));
      }
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(!(true))).length),false)).length)))).length), new BigNumber(183), (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-701), new BigNumber((_dafny.Seq.of(new BigNumber(951))).length), new BigNumber(-146))).cardinality())), new BigNumber(872)))));
    };
    static fm7(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-490),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("q")).length)),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC6(false, new BigNumber(813), new _dafny.CodePoint('s'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("nyaleobm")).length), new BigNumber(-473))).dtor_cf7,true)).Merge(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Set.fromElements(new BigNumber(35), new BigNumber(453))).Elements) {
          let _7_v0 = _compr_4;
          if ((_dafny.Set.fromElements(new BigNumber(35), new BigNumber(453))).contains(_7_v0)) {
            _coll4.push([_module.__default.safeDivisionInt(_7_v0, (_module.D5.create_DC12(_dafny.Set.fromElements(true), new BigNumber(149), true)).dtor_cf20),false]);
          }
        }
        return _coll4;
      }()));
    };
    static fm8(p0, p1, globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm9(p0, p1, globalState) {
      return ((new BigNumber(962)).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), function (_8_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length))).length))).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),!(true))).length));
    };
    static fm10(p0, p1, globalState) {
      return _dafny.Set.fromElements(true);
    };
    static fm11(globalState) {
      if (false) {
        return _module.D1.create_DC2(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(false))), _dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.FromArray(_dafny.Seq.of(false)), _dafny.MultiSet.fromElements(false, false)));
      } else {
        return _module.D1.create_DC2(_dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), function (_9_i0) {
  return _dafny.MultiSet.fromElements(false);
}));
      }
    };
    static fm12(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(978))).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
          let _10_v0 = _compr_5;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).contains(_10_v0)) {
            _coll5.push([_10_v0,new BigNumber(188)]);
          }
        }
        return _coll5;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),new BigNumber(149)));
    };
    static fm13(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.FromArray((_module.D7.create_DC16(_dafny.Seq.of(false))).dtor_cf26))).Difference((_dafny.MultiSet.fromElements(false, false, false)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(false))));
    };
    static Main(__noArgsParameter) {
      let _11_v0;
      _11_v0 = new BigNumber(141);
      let _12_v1;
      _12_v1 = _dafny.MultiSet.fromElements(_11_v0, _11_v0);
      let _13_v2;
      let _nw0 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _13_v2 = _nw0;
      let _14_v3;
      _14_v3 = false;
      let _15_v4;
      _15_v4 = _dafny.Map.Empty.slice().updateUnsafe(_11_v0,_11_v0);
      let _16_v5;
      _16_v5 = _dafny.Map.Empty.slice().updateUnsafe(_14_v3,new BigNumber((_15_v4).length));
      let _17_v6;
      let _init0 = ((_18_v3) => function (_19_i0) {
        return _18_v3;
      })(_14_v3);
      let _nw1 = Array((new BigNumber(22)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
        _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _17_v6 = _nw1;
      let _20_v7;
      _20_v7 = _dafny.Map.Empty.slice().updateUnsafe(_11_v0,_17_v6);
      let _21_v8;
      _21_v8 = _dafny.Map.Empty.slice().updateUnsafe(_14_v3,_17_v6);
      let _22_v9;
      _22_v9 = _module.D0.create_DC0(_17_v6);
      let _23_v10;
      let _nw2 = Array((new BigNumber(8)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = (((_20_v7).contains(_11_v0)) ? ((_20_v7).get(_11_v0)) : (_17_v6));
      _nw2[(_dafny.ONE).toNumber()] = _17_v6;
      _nw2[(new BigNumber(2)).toNumber()] = (((_21_v8).contains(_14_v3)) ? ((_21_v8).get(_14_v3)) : ((_22_v9).dtor_cf0));
      _nw2[(new BigNumber(3)).toNumber()] = _17_v6;
      _nw2[(new BigNumber(4)).toNumber()] = _17_v6;
      _nw2[(new BigNumber(5)).toNumber()] = (((_20_v7).contains(_11_v0)) ? ((_20_v7).get(_11_v0)) : (_17_v6));
      _nw2[(new BigNumber(6)).toNumber()] = _17_v6;
      _nw2[(new BigNumber(7)).toNumber()] = _17_v6;
      _23_v10 = _nw2;
      let _24_v11;
      let _init1 = ((_25_v0) => function (_26_i1) {
        return (_26_i1).multipliedBy(_25_v0);
      })(_11_v0);
      let _nw3 = Array((new BigNumber(20)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
        _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _24_v11 = _nw3;
      let _27_v12;
      _27_v12 = _dafny.Seq.UnicodeFromString("h");
      let _28_v13;
      _28_v13 = new _dafny.CodePoint('y'.codePointAt(0));
      let _29_globalState;
      let _nw4 = new _module.GlobalState();
      _nw4.__ctor((_12_v1).Union(_12_v1), false, false, new BigNumber(736), _13_v2, true, new BigNumber(128), new BigNumber(130), new _dafny.CodePoint('l'.codePointAt(0)), _16_v5, new BigNumber(-233), false, _23_v10, _24_v11, new BigNumber(84), _15_v4, _dafny.Seq.update(_27_v12, _module.__default.safeIndex(_11_v0, new BigNumber((_27_v12).length)), _28_v13), new BigNumber(795), new BigNumber(109), new BigNumber(-292));
      _29_globalState = _nw4;
      (_29_globalState).f2 = !(!((_14_v3) || (_14_v3)));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_17_v6).length))) {
        let _30_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_30_i2)) && ((_30_i2).isLessThan(new BigNumber((_17_v6).length))))) {
          (_17_v6)[(_30_i2)] = _14_v3;
        }
      }
      let _31_i3;
      _31_i3 = _dafny.ZERO;
      L0: {
        while (_14_v3) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_31_i3)) {
              break L0;
            }
            _31_i3 = (_31_i3).plus(_dafny.ONE);
            let _32_v14;
            _32_v14 = _module.D0.create_DC1();
            let _33_v15;
            _33_v15 = _dafny.Map.Empty.slice().updateUnsafe(_14_v3,_14_v3);
            let _34_v16;
            _34_v16 = _dafny.Map.Empty.slice().updateUnsafe(_16_v5,new BigNumber(((_33_v15).Merge(_33_v15)).length));
            let _pat_let_tv0 = _17_v6;
            let _rhs0 = function (_pat_let0_0) {
              return function (_35_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_36_dt__update_hcf0_h0) {
                    return _module.D0.create_DC0(_36_dt__update_hcf0_h0);
                  }(_pat_let1_0);
                }(_pat_let_tv0);
              }(_pat_let0_0);
            }(_module.D0.create_DC0(_17_v6));
            let _rhs1 = _32_v14;
            let _rhs2 = _34_v16;
            let _rhs3 = _module.__default.safeDivisionInt(new BigNumber(778), _11_v0);
            let _lhs0 = _29_globalState;
            _22_v9 = _rhs0;
            _32_v14 = _rhs1;
            _34_v16 = _rhs2;
            _lhs0.f6 = _rhs3;
            let _37_v17;
            let _nw5 = new _module.C1();
            _nw5.__ctor(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("wdbsbbie")).length), _module.__default.fm2(_29_globalState)), _24_v11);
            _37_v17 = _nw5;
            (_29_globalState).f17 = new BigNumber((_27_v12).length);
            let _38_v18;
            let _out0;
            _out0 = (_37_v17).m0(_29_globalState);
            _38_v18 = _out0;
          }
        }
      }
      let _hi0 = _11_v0;
      for (let _39_i4 = _11_v0; _39_i4.isLessThan(_hi0); _39_i4 = _39_i4.plus(_dafny.ONE)) {
        let _40_v19;
        let _nw6 = new _module.C0();
        _nw6.__ctor();
        _40_v19 = _nw6;
        let _41_v20;
        _41_v20 = _dafny.Set.fromElements(_28_v13);
        _15_v4 = ((_15_v4).Merge(_15_v4)).Merge((_15_v4).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_41_v20).length),_11_v0)));
        (_29_globalState).f2 = (_12_v1).IsProperSubsetOf((_12_v1).Difference((_module.D6.create_DC14(_dafny.MultiSet.fromElements(_11_v0, _11_v0))).dtor_cf22));
        let _42_v21;
        let _nw7 = new _module.C1();
        _nw7.__ctor(_11_v0, _24_v11);
        _42_v21 = _nw7;
        (_29_globalState).f17 = (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_42_v21,_39_i4)).length)).plus(_module.__default.safeModuloInt((_42_v21).f20, _11_v0)));
      }
      let _43_v22;
      _43_v22 = _dafny.Seq.of(_11_v0);
      _43_v22 = _43_v22;
      let _index0 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
      (_17_v6)[_index0] = _14_v3;
      let _index1 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
      (_17_v6)[_index1] = ((_14_v3) ? (!(true) || (_14_v3)) : (_14_v3));
      let _44_v23;
      let _nw8 = new _module.C0();
      _nw8.__ctor();
      _44_v23 = _nw8;
      let _45_v24;
      _45_v24 = _dafny.Seq.of(_44_v23);
      let _46_v25;
      _46_v25 = _dafny.Map.Empty.slice().updateUnsafe(_45_v24,new _dafny.CodePoint('t'.codePointAt(0)));
      let _47_v26;
      _47_v26 = _dafny.Map.Empty.slice().updateUnsafe(_14_v3,_44_v23);
      let _48_v27;
      _48_v27 = _dafny.Seq.of(_45_v24, _45_v24, _45_v24, _45_v24, _dafny.Seq.of(_44_v23, _44_v23, _44_v23, _44_v23, (((_47_v26).contains((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) ? ((_47_v26).get((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) : (_44_v23))));
      let _49_v28;
      _49_v28 = _dafny.Set.fromElements(_11_v0);
      _46_v25 = (_46_v25).update((_48_v27)[_module.__default.safeIndex(new BigNumber((_49_v28).length), new BigNumber((_48_v27).length))], _28_v13);
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_17_v6).length))) {
        let _50_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_50_i5)) && ((_50_i5).isLessThan(new BigNumber((_17_v6).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_17_v6, (_50_i5).toNumber(), (((((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]) ? (_14_v3) : ((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]))) ? (_14_v3) : ((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _51_v29;
      _51_v29 = _dafny.Set.fromElements(_14_v3);
      let _52_v30;
      _52_v30 = _dafny.Map.Empty.slice().updateUnsafe((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))],_51_v29);
      let _53_v31;
      _53_v31 = _dafny.Seq.of(true);
      let _54_v32;
      _54_v32 = _dafny.Seq.of(_53_v31, _53_v31);
      let _index2 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
      (_17_v6)[_index2] = (_44_v23).fm1(_module.D0.create_DC1(), ((((_52_v30).contains(_14_v3)) ? ((_52_v30).get(_14_v3)) : (_module.__default.fm10(_11_v0, _11_v0, _29_globalState)))).IsProperSubsetOf(_module.__default.fm10(_11_v0, _11_v0, _29_globalState)), _dafny.MultiSet.fromElements((_53_v31)[_module.__default.safeIndex(new BigNumber((_54_v32).length), new BigNumber((_53_v31).length))]), _29_globalState);
      (_44_v23).m1(_29_globalState);
      if (true) {
        let _index3 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_13_v2).length));
        (_13_v2)[_index3] = _dafny.Seq.UnicodeFromString("eaalrjgr");
        let _index4 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_13_v2).length));
        (_13_v2)[_index4] = _dafny.Seq.Concat(_27_v12, _dafny.Seq.UnicodeFromString("mewdink"));
        (_29_globalState).f17 = new BigNumber(719);
        let _55_v33;
        _55_v33 = _dafny.MultiSet.fromElements((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], false);
        let _56_v34;
        _56_v34 = _dafny.Seq.of(_55_v33);
        let _57_v35;
        _57_v35 = _module.D1.create_DC2(_56_v34);
        let _58_v36;
        _58_v36 = _dafny.Map.Empty.slice().updateUnsafe((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))],_57_v35);
        let _59_v37;
        let _nw9 = Array((new BigNumber(2)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = _58_v36;
        _nw9[(_dafny.ONE).toNumber()] = _58_v36;
        _59_v37 = _nw9;
        let _index5 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_59_v37).length));
        (_59_v37)[_index5] = (((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]) ? (_58_v36) : (_dafny.Map.Empty.slice().updateUnsafe((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))],_57_v35)));
        let _60_v38;
        _60_v38 = _dafny.Seq.of(_58_v36, _58_v36, (_58_v36).update(_14_v3, _57_v35));
        let _index6 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_59_v37).length));
        (_59_v37)[_index6] = ((_60_v38)[_module.__default.safeIndex(_11_v0, new BigNumber((_60_v38).length))]).update((_49_v28).IsSubsetOf(_dafny.Set.fromElements(_11_v0)), _module.__default.fm11(_29_globalState));
        _53_v31 = _dafny.Seq.of(((true) ? (_14_v3) : (_module.__default.fm9(_11_v0, (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], _29_globalState))), (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]);
        let _61_v39;
        let _nw10 = new _module.C1();
        _nw10.__ctor(new BigNumber(-422), _24_v11);
        _61_v39 = _nw10;
        let _62_v40;
        _62_v40 = _dafny.Map.Empty.slice().updateUnsafe(_61_v39,(_61_v39).f21);
        (_29_globalState).f2 = !(!((new BigNumber((_62_v40).length)).multipliedBy(new BigNumber(939))).isEqualTo(((_61_v39).f20).minus(_11_v0)));
      } else {
        let _63_v41;
        _63_v41 = _module.D0.create_DC1();
        let _64_v42;
        _64_v42 = _dafny.MultiSet.fromElements(true, (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]);
        let _65_v43;
        _65_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_11_v0),(_44_v23).fm1(_63_v41, !((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]), _64_v42, _29_globalState));
        if ((((_65_v43).contains(_module.__default.fm12((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], _11_v0, _14_v3, _29_globalState))) ? ((_65_v43).get(_module.__default.fm12((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], _11_v0, _14_v3, _29_globalState))) : ((_64_v42).IsProperSubsetOf(_64_v42)))) {
          let _66_v44;
          let _nw11 = new _module.C1();
          _nw11.__ctor(_11_v0, ((false) ? (_24_v11) : (_24_v11)));
          _66_v44 = _nw11;
          _27_v12 = _dafny.Seq.Concat(_27_v12, _dafny.Seq.UnicodeFromString("jasdwxryi"));
          let _67_v45;
          let _out1;
          _out1 = (_66_v44).m0(_29_globalState);
          _67_v45 = _out1;
          let _68_v46;
          _68_v46 = _dafny.Map.Empty.slice().updateUnsafe(_11_v0,_24_v11);
          let _69_v47;
          _69_v47 = _module.D4.create_DC10(_11_v0, (_68_v46).Merge(_68_v46));
          let _70_v48;
          _70_v48 = _dafny.Map.Empty.slice().updateUnsafe(_17_v6,_67_v45);
          let _rhs4 = (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))];
          let _rhs5 = ((!((_14_v3) === (_14_v3))) ? ((_66_v44).f20) : ((_66_v44).f20));
          let _rhs6 = _69_v47;
          let _rhs7 = (_70_v48).update(_17_v6, _67_v45);
          let _lhs1 = _29_globalState;
          let _lhs2 = _29_globalState;
          _lhs1.f2 = _rhs4;
          _lhs2.f6 = _rhs5;
          _69_v47 = _rhs6;
          _70_v48 = _rhs7;
          (_44_v23).m1(_29_globalState);
        } else {
          (_29_globalState).f17 = (new BigNumber((_53_v31).length)).multipliedBy(_11_v0);
          _51_v29 = ((_dafny.Set.fromElements(_14_v3)).Intersect(_51_v29)).Difference(_51_v29);
          let _71_v49;
          let _nw12 = new _module.C1();
          _nw12.__ctor(_11_v0, _24_v11);
          _71_v49 = _nw12;
          let _index7 = _module.__default.safeIndex(new BigNumber(141), new BigNumber(((_71_v49).f21).length));
          ((_71_v49).f21)[_index7] = _11_v0;
          let _index8 = _module.__default.safeIndex(new BigNumber(141), new BigNumber(((_71_v49).f21).length));
          let _rhs8 = (((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]) ? (_16_v5) : ((_16_v5).Merge(_16_v5)));
          let _rhs9 = (_71_v49).f20;
          let _lhs3 = (_71_v49).f21;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(141), new BigNumber(((_71_v49).f21).length));
          _16_v5 = _rhs8;
          _lhs3[_lhs4] = _rhs9;
          (_29_globalState).f2 = !_dafny.areEqual(_27_v12, _dafny.Seq.update(_dafny.Seq.Concat(_27_v12, _27_v12), _module.__default.safeIndex(((_71_v49).f21)[_module.__default.safeIndex(new BigNumber(141), new BigNumber(((_71_v49).f21).length))], new BigNumber((_dafny.Seq.Concat(_27_v12, _27_v12)).length)), _28_v13));
        }
        let _72_v50;
        _72_v50 = _dafny.Map.Empty.slice().updateUnsafe((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))],new _dafny.CodePoint('h'.codePointAt(0)));
        let _73_v51;
        let _nw13 = Array((new BigNumber(20)).toNumber());
        _nw13[(_dafny.ZERO).toNumber()] = _28_v13;
        _nw13[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw13[(new BigNumber(2)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(3)).toNumber()] = (((_72_v50).contains((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) ? ((_72_v50).get((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) : (_28_v13));
        _nw13[(new BigNumber(4)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(5)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(6)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(7)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(8)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(9)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(10)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(11)).toNumber()] = (_44_v23).fm0(_27_v12, _29_globalState);
        _nw13[(new BigNumber(12)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(13)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(14)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(15)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(16)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(17)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(18)).toNumber()] = _28_v13;
        _nw13[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
        _73_v51 = _nw13;
        let _index9 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_73_v51).length));
        (_73_v51)[_index9] = _28_v13;
        let _index10 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_73_v51).length));
        (_73_v51)[_index10] = _28_v13;
        let _74_v52;
        _74_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_27_v12).length),_49_v28);
        if ((_44_v23).fm1(_module.D0.create_DC1(), ((((_74_v52).contains(new BigNumber((_54_v32).length))) ? ((_74_v52).get(new BigNumber((_54_v32).length))) : (_module.__default.fm5(_29_globalState)))).IsDisjointFrom(_49_v28), _64_v42, _29_globalState)) {
          (_44_v23).m1(_29_globalState);
          let _75_v53;
          _75_v53 = _dafny.Map.Empty.slice().updateUnsafe(_14_v3,_53_v31);
          _53_v31 = (((_75_v53).contains((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) ? ((_75_v53).get((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) : (_53_v31));
          (_29_globalState).f2 = !(_11_v0).isEqualTo(_11_v0);
          let _76_v54;
          _76_v54 = _module.D1.create_DC3();
          _76_v54 = _module.D1.create_DC3();
          (_29_globalState).f17 = (_11_v0).plus(_11_v0);
        } else {
          let _77_v55;
          _77_v55 = _dafny.Map.Empty.slice().updateUnsafe((((_15_v4).contains(_11_v0)) ? ((_15_v4).get(_11_v0)) : (new BigNumber((_27_v12).length))),_28_v13);
          let _rhs10 = _11_v0;
          let _rhs11 = (_49_v28).Intersect((_49_v28).Union(function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (_77_v55).Keys.Elements) {
              let _78_v56 = _compr_6;
              if ((_77_v55).contains(_78_v56)) {
                _coll6.add((_78_v56).multipliedBy(new BigNumber((_dafny.Set.fromElements(!(false))).length)));
              }
            }
            return _coll6;
          }()));
          let _rhs12 = new BigNumber(722);
          let _lhs5 = _29_globalState;
          _11_v0 = _rhs10;
          _49_v28 = _rhs11;
          _lhs5.f6 = _rhs12;
          let _index11 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_24_v11).length));
          (_24_v11)[_index11] = _11_v0;
          let _79_v57;
          _79_v57 = _dafny.Map.Empty.slice().updateUnsafe(_16_v5,_11_v0);
          let _index12 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_24_v11).length));
          (_24_v11)[_index12] = (_11_v0).plus(_module.__default.safeDivisionInt((((_79_v57).contains(_16_v5)) ? ((_79_v57).get(_16_v5)) : (_11_v0)), _11_v0));
          let _80_v58;
          _80_v58 = _dafny.Seq.of(_17_v6, _17_v6, _17_v6, _17_v6);
          let _81_v59;
          _81_v59 = _dafny.Map.Empty.slice().updateUnsafe(_80_v58,!((_module.D3.create_DC6(_14_v3, _11_v0, _28_v13, (_24_v11)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_24_v11).length))], new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_24_v11)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_24_v11).length))]))).cardinality()))).dtor_cf4));
          _81_v59 = (_81_v59).update(_80_v58, (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]);
          let _index13 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_73_v51).length));
          (_73_v51)[_index13] = new _dafny.CodePoint('c'.codePointAt(0));
          let _82_v60;
          let _nw14 = new _module.C0();
          _nw14.__ctor();
          _82_v60 = _nw14;
        }
        _11_v0 = new BigNumber(569);
        (_29_globalState).f6 = _11_v0;
      }
      let _83_i6;
      _83_i6 = _dafny.ZERO;
      L1: {
        while (!(_49_v28).contains(_11_v0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_83_i6)) {
              break L1;
            }
            _83_i6 = (_83_i6).plus(_dafny.ONE);
            _49_v28 = _module.__default.fm5(_29_globalState);
            (_44_v23).m1(_29_globalState);
            (_29_globalState).f2 = !(!(_14_v3));
            _11_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_11_v0));
          }
        }
      }
      let _84_i7;
      _84_i7 = _dafny.ZERO;
      L2: {
        while ((_14_v3) === ((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))])) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_84_i7)) {
              break L2;
            }
            _84_i7 = (_84_i7).plus(_dafny.ONE);
            let _85_v61;
            let _nw15 = new _module.C1();
            _nw15.__ctor((_11_v0).multipliedBy(_module.__default.fm2(_29_globalState)), _24_v11);
            _85_v61 = _nw15;
            let _86_v62;
            _86_v62 = _dafny.Map.Empty.slice().updateUnsafe((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))],(_27_v12)[_module.__default.safeIndex(_11_v0, new BigNumber((_27_v12).length))]);
            let _87_v63;
            _87_v63 = _module.D3.create_DC6(_14_v3, _11_v0, _28_v13, _11_v0, new BigNumber((_27_v12).length));
            _86_v62 = (_86_v62).update((_87_v63).dtor_cf4, _28_v13);
            let _88_v65;
            let _init2 = ((_89_v31, _90_v3) => function (_91_i8) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(758)), ((_92_v31, _93_v3) => function (_94_i9) {
                return new BigNumber((function () {
                  let _coll7 = new _dafny.Map();
                  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(533), new BigNumber(290))) {
                    let _95_v64 = _compr_7;
                    if (((new BigNumber(533)).isLessThanOrEqualTo(_95_v64)) && ((_95_v64).isLessThan(new BigNumber(290)))) {
                      _coll7.push([_module.__default.safeModuloInt(_95_v64, new BigNumber((_92_v31).length)),!(_93_v3)]);
                    }
                  }
                  return _coll7;
                }()).length);
              })(_89_v31, _90_v3));
            })(_53_v31, _14_v3);
            let _nw16 = Array((new BigNumber(2)).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw16.length); _i0_2++) {
              _nw16[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _88_v65 = _nw16;
            let _index14 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
            let _rhs13 = !((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]);
            let _rhs14 = _dafny.Set.fromElements(_11_v0, _11_v0, (_85_v61).f20, _11_v0);
            let _rhs15 = _88_v65;
            let _lhs6 = _17_v6;
            let _lhs7 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
            _lhs6[_lhs7] = _rhs13;
            _49_v28 = _rhs14;
            _88_v65 = _rhs15;
            let _96_v66;
            _96_v66 = _dafny.MultiSet.fromElements(_85_v61, _85_v61, _85_v61, _85_v61, _85_v61);
            if ((((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]) ? ((_96_v66).equals(_96_v66)) : (_dafny.Seq.IsPrefixOf(_27_v12, _dafny.Seq.UnicodeFromString("cm"))))) {
              (_29_globalState).f17 = _11_v0;
              let _97_v67;
              _97_v67 = _dafny.MultiSet.fromElements(_28_v13, new _dafny.CodePoint('j'.codePointAt(0)));
              let _98_v68;
              _98_v68 = _dafny.Map.Empty.slice().updateUnsafe((_85_v61).f20,(_module.D4.create_DC9(_97_v67, !((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]), _97_v67, _11_v0, _17_v6)).dtor_cf12);
              _98_v68 = _98_v68;
              (_29_globalState).f17 = (_dafny.ZERO).minus(new BigNumber((_27_v12).length));
              let _index15 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
              (_17_v6)[_index15] = _14_v3;
              _44_v23 = _44_v23;
            } else {
              let _99_v69;
              _99_v69 = _dafny.Seq.of(_27_v12, _27_v12);
              let _100_v70;
              _100_v70 = _module.D3.create_DC5(_module.__default.fm4(_99_v69, _14_v3, _29_globalState));
              let _101_v71;
              _101_v71 = _dafny.Map.Empty.slice().updateUnsafe(_100_v70,_14_v3);
              let _index16 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
              (_17_v6)[_index16] = (((_101_v71).contains(_100_v70)) ? ((_101_v71).get(_100_v70)) : ((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]));
              (_29_globalState).f6 = new BigNumber(-303);
              let _index17 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_24_v11).length));
              (_24_v11)[_index17] = new BigNumber((_16_v5).length);
              let _index18 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_24_v11).length));
              let _rhs16 = (_dafny.ZERO).minus((_85_v61).f20);
              let _rhs17 = new BigNumber((_27_v12).length);
              let _rhs18 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), function (_102_i10) {
                return new _dafny.CodePoint('c'.codePointAt(0));
              }), _dafny.Seq.update(_27_v12, _module.__default.safeIndex(new BigNumber((_27_v12).length), new BigNumber((_27_v12).length)), _28_v13))).length));
              let _rhs19 = ((_85_v61).f20).minus(new BigNumber((_dafny.Seq.Concat(_27_v12, _dafny.Seq.UnicodeFromString("kur"))).length));
              let _lhs8 = _24_v11;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_24_v11).length));
              let _lhs10 = _29_globalState;
              let _lhs11 = _29_globalState;
              let _lhs12 = _29_globalState;
              _lhs8[_lhs9] = _rhs16;
              _lhs10.f17 = _rhs17;
              _lhs11.f17 = _rhs18;
              _lhs12.f17 = _rhs19;
              (_29_globalState).f13 = (_85_v61).f21;
              (_29_globalState).f6 = (_85_v61).f20;
            }
          }
        }
      }
      (_44_v23).m1(_29_globalState);
      let _source0 = _module.D3.create_DC5(_dafny.Seq.Concat(_27_v12, _27_v12));
      if (_source0.is_DC6) {
        let _103___mcc_h0 = (_source0).cf4;
        let _104___mcc_h1 = (_source0).cf5;
        let _105___mcc_h2 = (_source0).cf6;
        let _106___mcc_h3 = (_source0).cf7;
        let _107___mcc_h4 = (_source0).cf8;
        let _108_cf8 = _107___mcc_h4;
        let _109_cf7 = _106___mcc_h3;
        let _110_cf6 = _105___mcc_h2;
        let _111_cf5 = _104___mcc_h1;
        let _112_cf4 = _103___mcc_h0;
        _14_v3 = _14_v3;
        let _113_v72;
        _113_v72 = _dafny.Seq.of(_24_v11);
        let _114_v73;
        _114_v73 = _dafny.Seq.of(_dafny.Seq.Concat(_27_v12, _27_v12), _27_v12, _27_v12);
        let _115_v74;
        _115_v74 = _dafny.Map.Empty.slice().updateUnsafe(_111_cf5,_112_cf4);
        let _index19 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
        let _rhs20 = new BigNumber(((_115_v74).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm2(_29_globalState)),!(_14_v3)))).length);
        let _rhs21 = new BigNumber((_27_v12).length);
        let _rhs22 = _113_v72;
        let _rhs23 = _114_v73;
        let _rhs24 = (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))];
        let _lhs13 = _29_globalState;
        let _lhs14 = _17_v6;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
        _lhs13.f6 = _rhs20;
        _109_cf7 = _rhs21;
        _113_v72 = _rhs22;
        _114_v73 = _rhs23;
        _lhs14[_lhs15] = _rhs24;
        let _index20 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_24_v11).length));
        (_24_v11)[_index20] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("oqich")).length), _109_cf7));
        let _index21 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_24_v11).length));
        (_24_v11)[_index21] = _108_cf8;
        _108_cf8 = ((_24_v11)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_24_v11).length))]).multipliedBy(new BigNumber((_module.__default.fm10(_111_cf5, new BigNumber(635), _29_globalState)).length));
      } else if (_source0.is_DC5) {
        let _116___mcc_h5 = (_source0).cf3;
        let _117_cf3 = _116___mcc_h5;
        let _118_v75;
        _118_v75 = _dafny.Map.Empty.slice().updateUnsafe(_12_v1,_53_v31);
        let _119_v76;
        let _nw17 = Array((new BigNumber(9)).toNumber());
        _nw17[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(false), _53_v31);
        _nw17[(_dafny.ONE).toNumber()] = _53_v31;
        _nw17[(new BigNumber(2)).toNumber()] = _53_v31;
        _nw17[(new BigNumber(3)).toNumber()] = _53_v31;
        _nw17[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_53_v31, _53_v31);
        _nw17[(new BigNumber(5)).toNumber()] = _53_v31;
        _nw17[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((((_118_v75).contains(_12_v1)) ? ((_118_v75).get(_12_v1)) : (_53_v31)), _53_v31);
        _nw17[(new BigNumber(7)).toNumber()] = _53_v31;
        _nw17[(new BigNumber(8)).toNumber()] = _53_v31;
        _119_v76 = _nw17;
        let _120_v77;
        _120_v77 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_11_v0, (_dafny.ZERO).minus(_11_v0)),_119_v76);
        _119_v76 = (((_120_v77).contains(_module.__default.safeModuloInt(_11_v0, _module.__default.fm2(_29_globalState)))) ? ((_120_v77).get(_module.__default.safeModuloInt(_11_v0, _module.__default.fm2(_29_globalState)))) : (_119_v76));
        if (_14_v3) {
          (_44_v23).m1(_29_globalState);
          let _121_v78;
          let _init3 = ((_122_v13) => function (_123_i11) {
            return _122_v13;
          })(_28_v13);
          let _nw18 = Array((new BigNumber(28)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw18.length); _i0_3++) {
            _nw18[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _121_v78 = _nw18;
          let _index22 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_121_v78).length));
          (_121_v78)[_index22] = _28_v13;
          let _124_v79;
          _124_v79 = _dafny.MultiSet.fromElements(_117_cf3);
          let _index23 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_121_v78).length));
          let _rhs25 = _28_v13;
          let _rhs26 = (_dafny.MultiSet.fromElements(_117_cf3, _dafny.Seq.UnicodeFromString("qjnnuhas"))).IsDisjointFrom(_124_v79);
          let _lhs16 = _121_v78;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_121_v78).length));
          let _lhs18 = _29_globalState;
          _lhs16[_lhs17] = _rhs25;
          _lhs18.f2 = _rhs26;
          (_29_globalState).f6 = (_dafny.ZERO).minus((_11_v0).minus(new BigNumber((_49_v28).length)));
          _14_v3 = _14_v3;
          (_29_globalState).f17 = _11_v0;
        } else {
          let _125_v80;
          _125_v80 = _module.D1.create_DC3();
          let _126_v81;
          _126_v81 = _dafny.MultiSet.fromElements(_125_v80, _module.D1.create_DC3());
          let _127_v82;
          _127_v82 = _dafny.Map.Empty.slice().updateUnsafe(_126_v81,_24_v11);
          _127_v82 = (_127_v82).update(_126_v81, _24_v11);
          let _128_v83;
          let _nw19 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _128_v83 = _nw19;
          let _index24 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_128_v83).length));
          (_128_v83)[_index24] = new _dafny.CodePoint('d'.codePointAt(0));
          let _index25 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
          let _index26 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_128_v83).length));
          let _rhs27 = _14_v3;
          let _rhs28 = _28_v13;
          let _rhs29 = _44_v23;
          let _rhs30 = new BigNumber((_module.__default.fm3(_11_v0, _51_v29, _dafny.Map.Empty.slice().updateUnsafe((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))],_14_v3), new BigNumber(166), _29_globalState)).length);
          let _rhs31 = _module.__default.fm9(_11_v0, _14_v3, _29_globalState);
          let _lhs19 = _17_v6;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
          let _lhs21 = _128_v83;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_128_v83).length));
          let _lhs23 = _29_globalState;
          let _lhs24 = _29_globalState;
          _lhs19[_lhs20] = _rhs27;
          _lhs21[_lhs22] = _rhs28;
          _44_v23 = _rhs29;
          _lhs23.f17 = _rhs30;
          _lhs24.f2 = _rhs31;
          (_44_v23).m1(_29_globalState);
          (_44_v23).m1(_29_globalState);
          _44_v23 = _44_v23;
        }
        _28_v13 = new _dafny.CodePoint('j'.codePointAt(0));
        (_44_v23).m1(_29_globalState);
      } else {
        let _129___mcc_h6 = (_source0).cf9;
        let _130_cf9 = _129___mcc_h6;
        _17_v6 = _17_v6;
        let _index27 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
        (_17_v6)[_index27] = (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))];
        if (!(_12_v1).equals((_dafny.MultiSet.fromElements(new BigNumber(914))).update((_dafny.ZERO).minus(_11_v0), _module.__default.abs(_module.__default.fm2(_29_globalState))))) {
          (_29_globalState).f17 = (_11_v0).minus(_11_v0);
          (_44_v23).m1(_29_globalState);
          let _index28 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length));
          (_24_v11)[_index28] = new BigNumber(600);
          let _131_v84;
          _131_v84 = _dafny.MultiSet.fromElements(_28_v13);
          let _132_v85;
          _132_v85 = _module.D4.create_DC9(_131_v84, (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], _dafny.MultiSet.fromElements(_28_v13, _28_v13, _28_v13), _11_v0, _17_v6);
          let _index29 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length));
          let _rhs32 = _module.__default.safeDivisionInt(_11_v0, (_11_v0).multipliedBy(new BigNumber(((_15_v4).update(new BigNumber((_27_v12).length), _11_v0)).length)));
          let _rhs33 = (_dafny.ZERO).minus((new BigNumber(851)).plus(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(_11_v0)).cardinality()), _11_v0)));
          let _rhs34 = new BigNumber((_dafny.Seq.of((_132_v85).dtor_cf12, _14_v3, _14_v3, (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], (_11_v0).isLessThan(_11_v0))).length);
          let _lhs25 = _29_globalState;
          let _lhs26 = _24_v11;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length));
          let _lhs28 = _29_globalState;
          _lhs25.f6 = _rhs32;
          _lhs26[_lhs27] = _rhs33;
          _lhs28.f6 = _rhs34;
          let _133_v86;
          _133_v86 = _module.D0.create_DC1();
          let _134_v87;
          _134_v87 = _dafny.Map.Empty.slice().updateUnsafe((_24_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length))],_49_v28);
          let _135_v88;
          _135_v88 = _module.D5.create_DC12(_dafny.Set.fromElements(_14_v3, (_44_v23).fm1(_133_v86, (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))], _module.__default.fm13(_28_v13, _29_globalState), _29_globalState)), new BigNumber(((((_134_v87).contains((_24_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length))])) ? ((_134_v87).get((_24_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length))])) : (_49_v28))).length), true);
          let _index30 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
          (_17_v6)[_index30] = (_135_v88).dtor_cf21;
          (_29_globalState).f6 = (((_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))]) ? (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_11_v0), _11_v0)) : (new BigNumber((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(977), new BigNumber(442))) {
              let _136_v89 = _compr_8;
              if (((new BigNumber(977)).isLessThanOrEqualTo(_136_v89)) && ((_136_v89).isLessThan(new BigNumber(442)))) {
                _coll8.push([(_136_v89).minus((_24_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_24_v11).length))]),_14_v3]);
              }
            }
            return _coll8;
          }()).length)));
        } else {
          (_44_v23).m1(_29_globalState);
          let _index31 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_24_v11).length));
          (_24_v11)[_index31] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(555)), ((_137_v13) => function (_138_i12) {
            return _137_v13;
          })(_28_v13))).length);
          let _index32 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_24_v11).length));
          (_24_v11)[_index32] = _11_v0;
          let _index33 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
          let _rhs35 = _dafny.MultiSet.fromElements((_24_v11)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_24_v11).length))]);
          let _rhs36 = _14_v3;
          let _rhs37 = (_17_v6)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length))];
          let _lhs29 = _17_v6;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_17_v6).length));
          _12_v1 = _rhs35;
          _lhs29[_lhs30] = _rhs36;
          _14_v3 = _rhs37;
          let _index34 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_13_v2).length));
          (_13_v2)[_index34] = _27_v12;
          let _index35 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_13_v2).length));
          (_13_v2)[_index35] = _27_v12;
          (_29_globalState).f17 = (_24_v11)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_24_v11).length))];
        }
        (_44_v23).m1(_29_globalState);
      }
      (_44_v23).m1(_29_globalState);
      process.stdout.write(_dafny.toString(_11_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_12_v1).equals(_dafny.MultiSet.fromElements(new BigNumber(141), new BigNumber(141)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_13_v2)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_14_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_15_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(141),new BigNumber(141)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_16_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_20_v7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_21_v8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_22_v9).dtor_cf0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ZERO])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[_dafny.ONE])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(2)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(3)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(4)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(5)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(6)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_23_v10)[new BigNumber(7)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_24_v11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_27_v12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_v13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_29_globalState).f0).equals(_dafny.MultiSet.fromElements(new BigNumber(141), new BigNumber(141), new BigNumber(141), new BigNumber(141)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_29_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_29_globalState).f4)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_29_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_29_globalState).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ZERO])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[_dafny.ONE])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(2)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(3)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(4)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(5)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(6)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_29_globalState).f12)[new BigNumber(7)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f13)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState.f15).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(141),new BigNumber(141)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_29_globalState).f16).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_29_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_31_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_43_v22, _dafny.Seq.of(new BigNumber(141)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_45_v24).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_46_v25).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_47_v26).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_48_v27).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_49_v28).equals(_dafny.Set.fromElements(new BigNumber(141), new BigNumber(262824)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v29).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_52_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_53_v31, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_54_v32, _dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_83_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_84_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
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
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC3() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC2(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3();
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
    static create_DC4(cf2) {
      let $dt = new D2(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.ZERO;
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
    static create_DC6(cf4, cf5, cf6, cf7, cf8) {
      let $dt = new D3(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC5(cf3) {
      let $dt = new D3(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC7(cf9) {
      let $dt = new D3(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC5" + "(" + this.cf3.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(false, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC9(cf11, cf12, cf13, cf14, cf15) {
      let $dt = new D4(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10(cf16, cf17) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8(cf10) {
      let $dt = new D4(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9(_dafny.MultiSet.Empty, false, _dafny.MultiSet.Empty, _dafny.ZERO, []);
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
    static create_DC12(cf19, cf20, cf21) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
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
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf18 === other.cf18;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(_dafny.Set.Empty, _dafny.ZERO, false);
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
    static create_DC15(cf23, cf24, cf25) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf22) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15([], _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC17(cf27, cf28) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D7(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = false;
      this.f6 = _dafny.ZERO;
      this.f13 = [];
      this.f15 = _dafny.Map.Empty;
      this.f17 = _dafny.ZERO;
      this._f0 = _dafny.MultiSet.Empty;
      this._f1 = false;
      this._f3 = _dafny.ZERO;
      this._f4 = [];
      this._f5 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f9 = _dafny.Map.Empty;
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = [];
      this._f14 = _dafny.ZERO;
      this._f16 = _dafny.Seq.UnicodeFromString("");
      this._f18 = _dafny.ZERO;
      this._f19 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (((false) ? (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-808)))) : (_dafny.MultiSet.fromElements(new BigNumber(537), new BigNumber(342), new BigNumber(-761), new BigNumber((_dafny.Seq.of(new BigNumber(-545))).length))))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(598))));
    };
    m1(globalState) {
      let _this = this;
      let _139_v0;
      _139_v0 = new BigNumber(-925);
      let _140_v1;
      _140_v1 = _dafny.Seq.of(_139_v0, _module.__default.fm2(globalState), _139_v0);
      let _141_v2;
      _141_v2 = true;
      let _142_v3;
      let _nw20 = Array((new BigNumber(6)).toNumber());
      _nw20[(_dafny.ZERO).toNumber()] = (_140_v1)[_module.__default.safeIndex(new BigNumber((_140_v1).length), new BigNumber((_140_v1).length))];
      _nw20[(_dafny.ONE).toNumber()] = ((_141_v2) ? (_139_v0) : (_139_v0));
      _nw20[(new BigNumber(2)).toNumber()] = _139_v0;
      _nw20[(new BigNumber(3)).toNumber()] = _139_v0;
      _nw20[(new BigNumber(4)).toNumber()] = _139_v0;
      _nw20[(new BigNumber(5)).toNumber()] = _module.__default.fm2(globalState);
      _142_v3 = _nw20;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_142_v3).length))) {
        let _143_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_143_i0)) && ((_143_i0).isLessThan(new BigNumber((_142_v3).length))))) {
          (_142_v3)[(_143_i0)] = (_143_i0).plus(new BigNumber((((_141_v2) ? (_dafny.Seq.UnicodeFromString("gpbyetby")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(393)), function (_144_i1) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })))).length));
        }
      }
      if (!(!(_141_v2) || (_141_v2))) {
        if (_141_v2) {
          let _145_v4;
          _145_v4 = _dafny.Seq.UnicodeFromString("m");
          (globalState).f2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_145_v4, _145_v4), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wdrdnpx"), _145_v4));
          let _146_v5;
          _146_v5 = _dafny.Set.fromElements(_141_v2, _141_v2);
          let _147_v6;
          _147_v6 = _dafny.Map.Empty.slice().updateUnsafe(_141_v2,_141_v2);
          let _148_v7;
          _148_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(23),_139_v0);
          let _rhs38 = ((_141_v2) ? (_module.__default.fm3(_139_v0, _146_v5, _147_v6, _module.__default.fm2(globalState), globalState)) : (_148_v7));
          let _rhs39 = _139_v0;
          let _lhs31 = globalState;
          let _lhs32 = globalState;
          _lhs31.f15 = _rhs38;
          _lhs32.f17 = _rhs39;
          _141_v2 = _141_v2;
          let _149_v8;
          let _init4 = ((_150_v2) => function (_151_i2) {
            return _150_v2;
          })(_141_v2);
          let _nw21 = Array((new BigNumber(5)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw21.length); _i0_4++) {
            _nw21[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _149_v8 = _nw21;
          let _index36 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_149_v8).length));
          (_149_v8)[_index36] = !(_141_v2);
          let _index37 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_149_v8).length));
          (_149_v8)[_index37] = _141_v2;
          let _152_v9;
          _152_v9 = _module.D0.create_DC0(_149_v8);
          let _153_v10;
          _153_v10 = _dafny.MultiSet.fromElements(_152_v9, _152_v9, _152_v9, _module.D0.create_DC0(_149_v8));
          _153_v10 = _153_v10;
        } else {
          (globalState).f17 = (_140_v1)[_module.__default.safeIndex(_139_v0, new BigNumber((_140_v1).length))];
          (globalState).f17 = _139_v0;
          let _154_v11;
          let _nw22 = Array((new BigNumber(13)).toNumber()).fill(_module.D0.Default());
          _154_v11 = _nw22;
          let _155_v12;
          let _nw23 = Array((new BigNumber(7)).toNumber()).fill(false);
          _155_v12 = _nw23;
          let _156_v13;
          _156_v13 = _module.D0.create_DC0(_155_v12);
          let _index38 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_154_v11).length));
          (_154_v11)[_index38] = _156_v13;
          let _index39 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_154_v11).length));
          (_154_v11)[_index39] = _156_v13;
          (globalState).f17 = _139_v0;
          let _157_v14;
          _157_v14 = _module.D0.create_DC1();
          _157_v14 = _157_v14;
        }
        let _158_v15;
        _158_v15 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_142_v3);
        if ((_158_v15).contains(_142_v3)) {
          let _index40 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          (_142_v3)[_index40] = _139_v0;
          let _159_v16;
          let _nw24 = Array((new BigNumber(10)).toNumber()).fill(false);
          _159_v16 = _nw24;
          let _index41 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_159_v16).length));
          (_159_v16)[_index41] = _141_v2;
          let _160_v17;
          _160_v17 = _module.D0.create_DC1();
          let _index42 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          let _index43 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_159_v16).length));
          let _rhs40 = new BigNumber(-30);
          let _rhs41 = _141_v2;
          let _rhs42 = new BigNumber((_dafny.MultiSet.fromElements(_139_v0, (_139_v0).plus(_139_v0), (_dafny.ZERO).minus(new BigNumber(-973)))).cardinality());
          let _rhs43 = _module.D0.create_DC1();
          let _lhs33 = _142_v3;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          let _lhs35 = _159_v16;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_159_v16).length));
          let _lhs37 = globalState;
          _lhs33[_lhs34] = _rhs40;
          _lhs35[_lhs36] = _rhs41;
          _lhs37.f17 = _rhs42;
          _160_v17 = _rhs43;
          _140_v1 = _dafny.Seq.Concat(_140_v1, _140_v1);
          let _161_v18;
          _161_v18 = _dafny.Seq.UnicodeFromString("g");
          let _162_v19;
          _162_v19 = _dafny.Seq.of(_141_v2, (_159_v16)[_module.__default.safeIndex(new BigNumber(763), new BigNumber((_159_v16).length))]);
          let _163_v20;
          _163_v20 = new _dafny.CodePoint('u'.codePointAt(0));
          let _164_v21;
          _164_v21 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("csfuxo"), _module.__default.safeIndex(_139_v0, new BigNumber((_dafny.Seq.UnicodeFromString("csfuxo")).length)), _163_v20));
          let _165_v22;
          _165_v22 = _dafny.Map.Empty.slice().updateUnsafe(_159_v16,_160_v17);
          let _index44 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          let _index45 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          let _rhs44 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_139_v0, ((_141_v2) ? ((_142_v3)[_module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length))]) : (new BigNumber((_162_v19).length))))));
          let _rhs45 = _module.__default.fm4(_164_v21, _141_v2, globalState);
          let _rhs46 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_165_v22).length), _139_v0), (_139_v0).minus(_139_v0));
          let _lhs38 = _142_v3;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          let _lhs40 = _142_v3;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_142_v3).length));
          _lhs38[_lhs39] = _rhs44;
          _161_v18 = _rhs45;
          _lhs40[_lhs41] = _rhs46;
          let _166_v23;
          _166_v23 = _module.D0.create_DC0(_159_v16);
          _166_v23 = _166_v23;
          (globalState).f2 = !((_159_v16)[_module.__default.safeIndex(new BigNumber(763), new BigNumber((_159_v16).length))]) || (!(_141_v2));
        } else {
          let _167_v24;
          _167_v24 = _dafny.Set.fromElements(_141_v2, _141_v2);
          _167_v24 = _167_v24;
          (globalState).f17 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_139_v0, (_139_v0).minus(_139_v0)));
          (globalState).f2 = !(_139_v0).isEqualTo(_139_v0);
          let _168_v25;
          let _nw25 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _168_v25 = _nw25;
          let _index46 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_168_v25).length));
          (_168_v25)[_index46] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(716)), function (_169_i3) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          });
          let _170_v26;
          _170_v26 = _dafny.Seq.UnicodeFromString("yxrjk");
          let _index47 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_168_v25).length));
          (_168_v25)[_index47] = _module.__default.fm4(_dafny.Seq.of(_170_v26), true, globalState);
          let _171_v27;
          let _nw26 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _171_v27 = _nw26;
          let _172_v28;
          _172_v28 = new _dafny.CodePoint('k'.codePointAt(0));
          let _173_v29;
          _173_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(763),_172_v28);
          let _174_v30;
          _174_v30 = _dafny.Map.Empty.slice().updateUnsafe(_171_v27,((_141_v2) ? (new BigNumber((_173_v29).length)) : (_139_v0)));
          _174_v30 = (_174_v30).update(_171_v27, _139_v0);
        }
        let _175_v31;
        let _nw27 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _175_v31 = _nw27;
        let _176_v32;
        _176_v32 = _dafny.Seq.UnicodeFromString("xl");
        let _177_v33;
        _177_v33 = _dafny.Set.fromElements(_139_v0, _139_v0);
        let _178_v34;
        _178_v34 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_177_v33);
        let _rhs47 = _139_v0;
        let _rhs48 = ((((_178_v34).contains(_142_v3)) ? ((_178_v34).get(_142_v3)) : (_177_v33))).IsProperSubsetOf(_module.__default.fm5(globalState));
        let _rhs49 = _175_v31;
        let _rhs50 = _dafny.Seq.Concat(_dafny.Seq.Concat(_176_v32, _176_v32), _176_v32);
        let _rhs51 = _141_v2;
        let _lhs42 = globalState;
        let _lhs43 = globalState;
        _139_v0 = _rhs47;
        _lhs42.f2 = _rhs48;
        _175_v31 = _rhs49;
        _176_v32 = _rhs50;
        _lhs43.f2 = _rhs51;
        (globalState).f17 = ((_139_v0).minus(new BigNumber((_dafny.Seq.of(_141_v2)).length))).plus(new BigNumber(-438));
        let _179_v35;
        _179_v35 = _dafny.MultiSet.fromElements(new BigNumber(65), new BigNumber((_140_v1).length));
        let _180_v36;
        _180_v36 = _dafny.MultiSet.fromElements((_140_v1)[_module.__default.safeIndex(_139_v0, new BigNumber((_140_v1).length))], new BigNumber((_179_v35).cardinality()), _139_v0, new BigNumber(-332));
        let _181_v37;
        _181_v37 = _module.D0.create_DC1();
        let _182_v38;
        _182_v38 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_180_v36).contains(_139_v0)),_181_v37);
        let _183_v39;
        _183_v39 = _dafny.Seq.of(true, _141_v2);
        let _184_v41;
        _184_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),_139_v0);
        let _185_v42;
        _185_v42 = _dafny.MultiSet.fromElements(_184_v41);
        let _186_v44;
        _186_v44 = _dafny.Set.fromElements(_184_v41);
        let _rhs52 = _module.__default.fm2(globalState);
        let _rhs53 = ((_177_v33).Intersect(_177_v33)).Difference((((_183_v39)[_module.__default.safeIndex(new BigNumber((_183_v39).length), new BigNumber((_183_v39).length))]) ? (function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(808), new BigNumber(773))) {
            let _187_v40 = _compr_9;
            if (((new BigNumber(808)).isLessThanOrEqualTo(_187_v40)) && ((_187_v40).isLessThan(new BigNumber(773)))) {
              _coll9.add((_187_v40).multipliedBy(_139_v0));
            }
          }
          return _coll9;
        }()) : (_177_v33)));
        let _rhs54 = (_186_v44).IsProperSubsetOf(function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_185_v42).Elements) {
            let _188_v43 = _compr_10;
            if ((_185_v42).contains(_188_v43)) {
              _coll10.add(_188_v43);
            }
          }
          return _coll10;
        }());
        let _rhs55 = false;
        let _rhs56 = _182_v38;
        let _lhs44 = globalState;
        let _lhs45 = globalState;
        let _lhs46 = globalState;
        _lhs44.f17 = _rhs52;
        _177_v33 = _rhs53;
        _lhs45.f2 = _rhs54;
        _lhs46.f2 = _rhs55;
        _182_v38 = _rhs56;
      } else {
        let _189_v45;
        _189_v45 = _dafny.Seq.of(!(_141_v2));
        let _190_v46;
        _190_v46 = _dafny.Map.Empty.slice().updateUnsafe((_189_v45)[_module.__default.safeIndex(_139_v0, new BigNumber((_189_v45).length))],(_139_v0).minus(_139_v0));
        (globalState).f17 = (((_190_v46).contains(!(_139_v0).isEqualTo(_139_v0))) ? ((_190_v46).get(!(_139_v0).isEqualTo(_139_v0))) : (_139_v0));
        if (!_dafny.areEqual(_189_v45, _189_v45)) {
          let _191_v47;
          _191_v47 = _module.D0.create_DC1();
          _191_v47 = _191_v47;
          let _192_v48;
          let _init5 = function (_193_i4) {
            return false;
          };
          let _nw28 = Array((new BigNumber(26)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw28.length); _i0_5++) {
            _nw28[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _192_v48 = _nw28;
          _192_v48 = _192_v48;
          (globalState).f2 = _141_v2;
          let _194_v49;
          _194_v49 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_141_v2)), _dafny.MultiSet.fromElements(_141_v2));
          let _195_v50;
          _195_v50 = _dafny.MultiSet.fromElements(_141_v2, _141_v2);
          let _196_v51;
          _196_v51 = _module.D1.create_DC2(_194_v49);
          _194_v49 = _dafny.Seq.Concat(_dafny.Seq.of(_195_v50, _195_v50, _195_v50), (_196_v51).dtor_cf1);
          (globalState).f13 = _142_v3;
        } else {
          (globalState).f2 = !(_141_v2);
          _189_v45 = _dafny.Seq.Concat(_dafny.Seq.Concat(_189_v45, _189_v45), _189_v45);
          _141_v2 = (_139_v0).isLessThanOrEqualTo(_139_v0);
          _141_v2 = (_139_v0).isEqualTo((_dafny.ZERO).minus(_139_v0));
          let _197_v52;
          _197_v52 = _module.D1.create_DC3();
          _197_v52 = _197_v52;
        }
        let _198_v53;
        _198_v53 = _dafny.MultiSet.fromElements(true);
        let _199_v55;
        _199_v55 = _dafny.MultiSet.fromElements(_139_v0);
        _141_v2 = !(_dafny.Set.fromElements(new BigNumber((function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (_199_v55).Elements) {
            let _200_v54 = _compr_11;
            if ((_199_v55).contains(_200_v54)) {
              _coll11.push([(_200_v54).minus(_139_v0),_139_v0]);
            }
          }
          return _coll11;
        }()).length))).contains(((((_198_v53).contains(_141_v2)) ? ((_198_v53).get(_141_v2)) : (_139_v0))).minus(new BigNumber(546)));
        _139_v0 = _139_v0;
        (globalState).f6 = (_dafny.ZERO).minus(_139_v0);
      }
      let _index48 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_142_v3).length));
      (_142_v3)[_index48] = (_139_v0).plus(_139_v0);
      let _index49 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_142_v3).length));
      (_142_v3)[_index49] = _module.__default.safeModuloInt(_139_v0, _139_v0);
      (globalState).f2 = !(_141_v2);
      let _201_v56;
      _201_v56 = _dafny.Set.fromElements(_141_v2);
      let _202_v57;
      _202_v57 = _dafny.Seq.of(_201_v56, (_201_v56).Union(_201_v56), _201_v56, _201_v56);
      _202_v57 = _dafny.Seq.update(_202_v57, _module.__default.safeIndex((_dafny.ZERO).minus((_142_v3)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_142_v3).length))]), new BigNumber((_202_v57).length)), _201_v56);
      _139_v0 = (_dafny.ZERO).minus(((_142_v3)[_module.__default.safeIndex(new BigNumber(146), new BigNumber((_142_v3).length))]));
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f20 = _dafny.ZERO;
      this._f21 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f20, f21) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    m0(globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f6 = (new BigNumber((_dafny.Seq.of((_this).f20)).length)).minus((_this).f20);
      let _203_v0;
      _203_v0 = true;
      let _204_i0;
      _204_i0 = _dafny.ZERO;
      L3: {
        while (_203_v0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_204_i0)) {
              break L3;
            }
            _204_i0 = (_204_i0).plus(_dafny.ONE);
            let _index50 = _module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length));
            ((_this).f21)[_index50] = (_this).f20;
            let _index51 = _module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length));
            ((_this).f21)[_index51] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f20));
            if ((_203_v0) && (((_203_v0) ? (_203_v0) : (false)))) {
              let _205_v1;
              _205_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,new _dafny.CodePoint('i'.codePointAt(0)));
              let _206_v2;
              _206_v2 = _dafny.Seq.of(_205_v1);
              (globalState).f6 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_206_v2, _206_v2), _206_v2)).length));
              let _207_v3;
              let _nw29 = new _module.C0();
              _nw29.__ctor();
              _207_v3 = _nw29;
              let _208_v4;
              let _nw30 = Array((new BigNumber(7)).toNumber());
              _nw30[(_dafny.ZERO).toNumber()] = _203_v0;
              _nw30[(_dafny.ONE).toNumber()] = _203_v0;
              _nw30[(new BigNumber(2)).toNumber()] = _203_v0;
              _nw30[(new BigNumber(3)).toNumber()] = _203_v0;
              _nw30[(new BigNumber(4)).toNumber()] = _203_v0;
              _nw30[(new BigNumber(5)).toNumber()] = _203_v0;
              _nw30[(new BigNumber(6)).toNumber()] = _203_v0;
              _208_v4 = _nw30;
              let _209_v5;
              _209_v5 = _dafny.Seq.of(_208_v4);
              let _210_v6;
              let _nw31 = Array((new BigNumber(18)).toNumber());
              _nw31[(_dafny.ZERO).toNumber()] = _208_v4;
              _nw31[(_dafny.ONE).toNumber()] = _208_v4;
              _nw31[(new BigNumber(2)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(3)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(4)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(5)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(6)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(7)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(8)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(9)).toNumber()] = (_209_v5)[_module.__default.safeIndex((_this).f20, new BigNumber((_209_v5).length))];
              _nw31[(new BigNumber(10)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(11)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(12)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(13)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(14)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(15)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(16)).toNumber()] = _208_v4;
              _nw31[(new BigNumber(17)).toNumber()] = _208_v4;
              _210_v6 = _nw31;
              let _211_v7;
              _211_v7 = _dafny.Map.Empty.slice().updateUnsafe(_210_v6,_203_v0);
              let _212_v8;
              _212_v8 = new _dafny.CodePoint('x'.codePointAt(0));
              let _213_v9;
              _213_v9 = _dafny.Set.fromElements(_212_v8, new _dafny.CodePoint('x'.codePointAt(0)), _212_v8);
              let _214_v10;
              _214_v10 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], new BigNumber((_213_v9).length))).length));
              let _215_v11;
              _215_v11 = _dafny.Map.Empty.slice().updateUnsafe(_212_v8,_214_v10);
              _211_v7 = (_211_v7).update(_210_v6, (_215_v11).contains(_212_v8));
              let _216_v12;
              _216_v12 = _dafny.Seq.UnicodeFromString("gjo");
              (globalState).f17 = _module.__default.safeModuloInt(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], new BigNumber((_216_v12).length));
              (globalState).f2 = _203_v0;
            } else {
              (globalState).f2 = _203_v0;
              let _217_v13;
              let _nw32 = new _module.C0();
              _nw32.__ctor();
              _217_v13 = _nw32;
              let _218_v14;
              _218_v14 = _dafny.Seq.UnicodeFromString("irpy");
              let _219_v15;
              _219_v15 = _module.D3.create_DC5(_218_v14);
              let _220_v16;
              _220_v16 = _dafny.Seq.of(new BigNumber(((_219_v15).dtor_cf3).length));
              let _221_v17;
              _221_v17 = _module.D4.create_DC8(_220_v16);
              let _222_v18;
              _222_v18 = _dafny.Set.fromElements((_this).f20);
              let _223_v19;
              _223_v19 = _dafny.MultiSet.fromElements(_222_v18);
              let _224_v20;
              _224_v20 = _dafny.Seq.of(_217_v13);
              let _225_v21;
              _225_v21 = _dafny.MultiSet.fromElements(new BigNumber(607));
              let _226_v22;
              _226_v22 = _dafny.Seq.of(_dafny.Seq.of((_this).f20, new BigNumber((_225_v21).cardinality())));
              let _227_v23;
              _227_v23 = new _dafny.CodePoint('c'.codePointAt(0));
              let _228_v24;
              _228_v24 = _dafny.Seq.of(_203_v0, false);
              let _229_v25;
              _229_v25 = _dafny.Map.Empty.slice().updateUnsafe(_227_v23,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_228_v24,((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))])).length));
              let _230_v27;
              _230_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))],(_this).f20);
              let _231_v28;
              _231_v28 = _dafny.Set.fromElements(_203_v0);
              let _pat_let_tv1 = _220_v16;
              let _232_v29;
              let _nw33 = Array((new BigNumber(22)).toNumber());
              _nw33[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))]);
              _nw33[(_dafny.ONE).toNumber()] = _220_v16;
              _nw33[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_220_v16, _220_v16);
              _nw33[(new BigNumber(3)).toNumber()] = (function (_pat_let2_0) {
                return function (_233_dt__update__tmp_h0) {
                  return function (_pat_let3_0) {
                    return function (_234_dt__update_hcf10_h0) {
                      return _module.D4.create_DC8(_234_dt__update_hcf10_h0);
                    }(_pat_let3_0);
                  }(_pat_let_tv1);
                }(_pat_let2_0);
              }(_221_v17)).dtor_cf10;
              _nw33[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_220_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(159)), function (_235_i1) {
                return ((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))];
              }));
              _nw33[(new BigNumber(5)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(6)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(7)).toNumber()] = _module.__default.fm6(_223_v19, (_this).f20, _203_v0, _222_v18, globalState);
              _nw33[(new BigNumber(8)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_220_v16, _220_v16);
              _nw33[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_220_v16, _220_v16);
              _nw33[(new BigNumber(11)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), function (_236_i2) {
                return ((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))];
              });
              _nw33[(new BigNumber(13)).toNumber()] = ((_203_v0) ? (_module.__default.fm6(_223_v19, (_this).f20, true, _222_v18, globalState)) : (_dafny.Seq.of(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], _module.__default.fm2(globalState), new BigNumber((_224_v20).length), ((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], ((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))])));
              _nw33[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), ((_237_v0) => function (_238_i3) {
                return new BigNumber((_dafny.Seq.of(_237_v0, false, !(_237_v0), _237_v0, _237_v0)).length);
              })(_203_v0));
              _nw33[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_220_v16, _module.__default.safeIndex(new BigNumber(865), new BigNumber((_220_v16).length)), new BigNumber((_226_v22).length)), _dafny.Seq.update(_220_v16, _module.__default.safeIndex((_this).f20, new BigNumber((_220_v16).length)), new BigNumber((_module.__default.fm7(_203_v0, _203_v0, globalState)).length)));
              _nw33[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f20, (_this).f20, new BigNumber((function () {
                let _coll12 = new _dafny.Set();
                for (const _compr_12 of (_229_v25).Keys.Elements) {
                  let _239_v26 = _compr_12;
                  if ((_229_v25).contains(_239_v26)) {
                    _coll12.add(_239_v26);
                  }
                }
                return _coll12;
              }()).length), (_this).f20), _dafny.Seq.of(new BigNumber((_230_v27).length), (_this).f20));
              _nw33[(new BigNumber(17)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(18)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_220_v16, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_203_v0)).length)), new BigNumber((_220_v16).length)), new BigNumber((_218_v14).length));
              _nw33[(new BigNumber(20)).toNumber()] = _220_v16;
              _nw33[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_220_v16, _module.__default.safeIndex(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], new BigNumber((_220_v16).length)), new BigNumber((_231_v28).length));
              _232_v29 = _nw33;
              let _index52 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_232_v29).length));
              (_232_v29)[_index52] = _dafny.Seq.of(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], (_this).f20, new BigNumber(414), new BigNumber(683), new BigNumber((_dafny.Seq.update(_218_v14, _module.__default.safeIndex(new BigNumber((_231_v28).length), new BigNumber((_218_v14).length)), _227_v23)).length));
              let _240_v30;
              _240_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-732),_217_v13);
              let _241_v31;
              _241_v31 = _dafny.MultiSet.fromElements(_227_v23);
              let _242_v32;
              _242_v32 = _dafny.Map.Empty.slice().updateUnsafe((((_240_v30).contains(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))])) ? ((_240_v30).get(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))])) : (_217_v13)),new BigNumber((_241_v31).cardinality()));
              let _243_v33;
              _243_v33 = _dafny.Seq.of(_222_v18, _222_v18);
              let _index53 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_232_v29).length));
              (_232_v29)[_index53] = _module.__default.fm6(_223_v19, (((_242_v32).contains(_217_v13)) ? ((_242_v32).get(_217_v13)) : ((_this).f20)), (_dafny.MultiSet.FromArray(_243_v33)).IsSubsetOf(_223_v19), _222_v18, globalState);
              let _244_v34;
              _244_v34 = _dafny.MultiSet.fromElements(_203_v0, _203_v0);
              let _245_v35;
              _245_v35 = _dafny.Seq.of(_244_v34);
              let _246_v36;
              let _init6 = ((_247_v24) => function (_248_i4) {
                return _247_v24;
              })(_228_v24);
              let _nw34 = Array((new BigNumber(21)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw34.length); _i0_6++) {
                _nw34[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _246_v36 = _nw34;
              let _249_v37;
              _249_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_245_v35, _245_v35),_246_v36);
              _249_v37 = (_249_v37).update(_245_v35, _246_v36);
              (globalState).f6 = new BigNumber(768);
            }
            let _250_v38;
            _250_v38 = _dafny.Map.Empty.slice().updateUnsafe(true,_203_v0);
            let _251_v39;
            _251_v39 = _dafny.Map.Empty.slice().updateUnsafe(_250_v38,!(_203_v0));
            _251_v39 = (_251_v39).update(_250_v38, _203_v0);
            let _252_v40;
            let _nw35 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _252_v40 = _nw35;
            let _253_v41;
            let _nw36 = new _module.C0();
            _nw36.__ctor();
            _253_v41 = _nw36;
            let _254_v42;
            _254_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_253_v41);
            let _255_v43;
            _255_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,_203_v0),_253_v41);
            let _256_v44;
            _256_v44 = _dafny.Map.Empty.slice().updateUnsafe(_253_v41,_203_v0);
            let _257_v45;
            _257_v45 = _dafny.Seq.of(_253_v41);
            let _258_v46;
            _258_v46 = _module.D5.create_DC11((_257_v45)[_module.__default.safeIndex(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))], new BigNumber((_257_v45).length))]);
            let _259_v47;
            let _nw37 = Array((new BigNumber(21)).toNumber());
            _nw37[(_dafny.ZERO).toNumber()] = false;
            _nw37[(_dafny.ONE).toNumber()] = _203_v0;
            _nw37[(new BigNumber(2)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(3)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(4)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(5)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(6)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(7)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(8)).toNumber()] = false;
            _nw37[(new BigNumber(9)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(10)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(11)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(12)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(13)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(14)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(15)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(16)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(17)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(18)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(19)).toNumber()] = _203_v0;
            _nw37[(new BigNumber(20)).toNumber()] = _203_v0;
            _259_v47 = _nw37;
            let _260_v48;
            _260_v48 = _dafny.Map.Empty.slice().updateUnsafe(_259_v47,_203_v0);
            let _261_v49;
            _261_v49 = _dafny.Map.Empty.slice().updateUnsafe(_260_v48,_253_v41);
            let _262_v50;
            let _nw38 = Array((new BigNumber(29)).toNumber());
            _nw38[(_dafny.ZERO).toNumber()] = _253_v41;
            _nw38[(_dafny.ONE).toNumber()] = _253_v41;
            _nw38[(new BigNumber(2)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(3)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(4)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(5)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(6)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(7)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(8)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(9)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(10)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(11)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(12)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(13)).toNumber()] = (((_254_v42).contains((_dafny.ZERO).minus((_this).f20))) ? ((_254_v42).get((_dafny.ZERO).minus((_this).f20))) : (_253_v41));
            _nw38[(new BigNumber(14)).toNumber()] = (((_255_v43).contains(_dafny.Map.Empty.slice().updateUnsafe((((_256_v44).contains(_253_v41)) ? ((_256_v44).get(_253_v41)) : (_203_v0)),_203_v0))) ? ((_255_v43).get(_dafny.Map.Empty.slice().updateUnsafe((((_256_v44).contains(_253_v41)) ? ((_256_v44).get(_253_v41)) : (_203_v0)),_203_v0))) : (_253_v41));
            _nw38[(new BigNumber(15)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(16)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(17)).toNumber()] = (_258_v46).dtor_cf18;
            _nw38[(new BigNumber(18)).toNumber()] = (((_261_v49).contains(_260_v48)) ? ((_261_v49).get(_260_v48)) : (_253_v41));
            _nw38[(new BigNumber(19)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(20)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(21)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(22)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(23)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(24)).toNumber()] = (((_254_v42).contains(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))])) ? ((_254_v42).get(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))])) : (_253_v41));
            _nw38[(new BigNumber(25)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(26)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(27)).toNumber()] = _253_v41;
            _nw38[(new BigNumber(28)).toNumber()] = _253_v41;
            _262_v50 = _nw38;
            let _index54 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_262_v50).length));
            (_262_v50)[_index54] = _253_v41;
            let _263_v51;
            _263_v51 = _module.D0.create_DC0(_259_v47);
            let _264_v52;
            _264_v52 = _dafny.MultiSet.fromElements(_263_v51, _module.D0.create_DC0(_259_v47), _263_v51, _263_v51, _263_v51);
            let _index55 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_262_v50).length));
            let _rhs57 = _252_v40;
            let _rhs58 = _203_v0;
            let _rhs59 = new BigNumber(((_264_v52).Difference((_264_v52).Intersect((_264_v52).update(_263_v51, _module.__default.abs(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))]))))).cardinality());
            let _rhs60 = ((_this).f20).plus(((_this).f21)[_module.__default.safeIndex(new BigNumber(565), new BigNumber(((_this).f21).length))]);
            let _rhs61 = _253_v41;
            let _lhs47 = globalState;
            let _lhs48 = globalState;
            let _lhs49 = _262_v50;
            let _lhs50 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_262_v50).length));
            _252_v40 = _rhs57;
            r0 = _rhs58;
            _lhs47.f17 = _rhs59;
            _lhs48.f6 = _rhs60;
            _lhs49[_lhs50] = _rhs61;
          }
        }
      }
      let _265_v53;
      let _init7 = function (_266_i6) {
        return (_266_i6).plus((_this).f20);
      };
      let _nw39 = Array((new BigNumber(28)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw39.length); _i0_7++) {
        _nw39[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _265_v53 = _nw39;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_265_v53).length))) {
        let _267_i5 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_267_i5)) && ((_267_i5).isLessThan(new BigNumber((_265_v53).length))))) {
          (_265_v53)[(_267_i5)] = (_267_i5).plus(new BigNumber((_dafny.MultiSet.fromElements(_203_v0)).cardinality()));
        }
      }
      let _268_v54;
      let _nw40 = Array((new BigNumber(6)).toNumber()).fill(_module.D5.Default());
      _268_v54 = _nw40;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_268_v54).length))) {
        let _269_i7 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_269_i7)) && ((_269_i7).isLessThan(new BigNumber((_268_v54).length))))) {
          (_268_v54)[(_269_i7)] = _module.D5.create_DC13();
        }
      }
      let _270_v55;
      _270_v55 = _dafny.Set.fromElements(false);
      let _271_v56;
      _271_v56 = new _dafny.CodePoint('i'.codePointAt(0));
      let _272_v57;
      _272_v57 = _dafny.MultiSet.fromElements(_module.__default.fm8(new BigNumber(-157), _203_v0, globalState));
      let _273_v58;
      _273_v58 = _module.D3.create_DC6(_203_v0, new BigNumber((_270_v55).length), _271_v56, (_this).f20, (((_272_v57).contains(_271_v56)) ? ((_272_v57).get(_271_v56)) : ((_this).f20)));
      _203_v0 = (_273_v58).dtor_cf4;
      let _274_i8;
      _274_i8 = _dafny.ZERO;
      L4: {
        while (!(_203_v0)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_274_i8)) {
              break L4;
            }
            _274_i8 = (_274_i8).plus(_dafny.ONE);
            _271_v56 = _271_v56;
            (globalState).f17 = (_this).f20;
            let _275_v59;
            _275_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_270_v55);
            _275_v59 = (_dafny.Map.Empty.slice().updateUnsafe((_this).f20,_270_v55)).update(((_203_v0) ? ((_this).f20) : ((_this).f20)), (_270_v55).Difference(_270_v55));
            let _276_v60;
            _276_v60 = _module.D4.create_DC8(_dafny.Seq.of((_this).f20));
            let _277_v61;
            _277_v61 = _dafny.MultiSet.fromElements(_276_v60, _276_v60);
            let _278_v62;
            let _nw41 = Array((new BigNumber(13)).toNumber());
            _nw41[(_dafny.ZERO).toNumber()] = _203_v0;
            _nw41[(_dafny.ONE).toNumber()] = _203_v0;
            _nw41[(new BigNumber(2)).toNumber()] = _203_v0;
            _nw41[(new BigNumber(3)).toNumber()] = _module.__default.fm9(new BigNumber(((_277_v61).update(_276_v60, _module.__default.abs(new BigNumber(827)))).cardinality()), _203_v0, globalState);
            _nw41[(new BigNumber(4)).toNumber()] = _203_v0;
            _nw41[(new BigNumber(5)).toNumber()] = _203_v0;
            _nw41[(new BigNumber(6)).toNumber()] = _203_v0;
            _nw41[(new BigNumber(7)).toNumber()] = _203_v0;
            _nw41[(new BigNumber(8)).toNumber()] = true;
            _nw41[(new BigNumber(9)).toNumber()] = false;
            _nw41[(new BigNumber(10)).toNumber()] = false;
            _nw41[(new BigNumber(11)).toNumber()] = !(_203_v0);
            _nw41[(new BigNumber(12)).toNumber()] = _203_v0;
            _278_v62 = _nw41;
            let _279_v63;
            _279_v63 = _module.D4.create_DC9(_dafny.MultiSet.fromElements(_module.__default.fm8((_this).f20, _203_v0, globalState)), _203_v0, (_272_v57).update(new _dafny.CodePoint('t'.codePointAt(0)), _module.__default.abs((_this).f20)), (_this).f20, _278_v62);
            let _280_v64;
            let _nw42 = Array((new BigNumber(25)).toNumber());
            _nw42[(_dafny.ZERO).toNumber()] = _278_v62;
            _nw42[(_dafny.ONE).toNumber()] = _278_v62;
            _nw42[(new BigNumber(2)).toNumber()] = (_279_v63).dtor_cf15;
            _nw42[(new BigNumber(3)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(4)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(5)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(6)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(7)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(8)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(9)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(10)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(11)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(12)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(13)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(14)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(15)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(16)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(17)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(18)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(19)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(20)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(21)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(22)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(23)).toNumber()] = _278_v62;
            _nw42[(new BigNumber(24)).toNumber()] = _278_v62;
            _280_v64 = _nw42;
            let _281_v65;
            _281_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f20), (_this).f20),_280_v64);
            _281_v65 = (_281_v65).update((_this).f20, _280_v64);
          }
        }
      }
      r0 = _203_v0;
      return r0;
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
