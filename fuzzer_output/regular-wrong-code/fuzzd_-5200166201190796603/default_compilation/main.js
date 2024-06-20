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
      return (_dafny.ZERO).minus((_dafny.ZERO).minus(((true) ? (((!(!(!(false)))) ? (new BigNumber(581)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-3),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xhl")).length))).length))).length)))) : (new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(981), new BigNumber(-818))).Difference(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-229)), new BigNumber((_dafny.Seq.of(true)).length))).Elements) {
          let _0_v0 = _compr_0;
          if ((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-229)), new BigNumber((_dafny.Seq.of(true)).length))).contains(_0_v0)) {
            _coll0.push([_module.__default.safeModuloInt(_0_v0, new BigNumber(-693)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-902),new _dafny.CodePoint('d'.codePointAt(0)))).length)]);
          }
        }
        return _coll0;
      }()).length), new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(430), new BigNumber(54))) {
          let _1_v1 = _compr_1;
          if (((new BigNumber(430)).isLessThanOrEqualTo(_1_v1)) && ((_1_v1).isLessThan(new BigNumber(54)))) {
            _coll1.add(_module.__default.safeModuloInt(_1_v1, new BigNumber(738)));
          }
        }
        return _coll1;
      }()).length)))).cardinality())))).cardinality())))));
    };
    static fm1(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), function (_2_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), function (_3_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tjfujd"), _dafny.Seq.UnicodeFromString("cnhbwhio")));
    };
    static fm2(p0, globalState) {
      return _module.D0.create_DC1(new BigNumber(977), false, new BigNumber(-930));
    };
    static fm3(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(true, !(false))).Union((_dafny.MultiSet.fromElements(false, true, false)).Difference(_dafny.MultiSet.fromElements(true)));
    };
    static fm4(p0, p1, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(953), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.of(true, true)).length))).length), new BigNumber(-771)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)))));
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("txsljff")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), function (_4_i0) {
        return _dafny.Seq.UnicodeFromString("cuu");
      }));
    };
    static fm6(p0, p1, globalState) {
      let _source0 = _module.D0.create_DC1(new BigNumber(369), false, new BigNumber(448));
      if (_source0.is_DC1) {
        let _5___mcc_h0 = (_source0).cf1;
        let _6___mcc_h1 = (_source0).cf2;
        let _7___mcc_h2 = (_source0).cf3;
        let _8_cf3 = _7___mcc_h2;
        let _9_cf2 = _6___mcc_h1;
        let _10_cf1 = _5___mcc_h0;
        return (_8_cf3).isLessThan(_10_cf1);
      } else if (_source0.is_DC2) {
        return (false) === (true);
      } else if (_source0.is_DC3) {
        return !(!(_dafny.Map.Empty.slice().updateUnsafe(!(false),!(false))).contains(true));
      } else {
        let _11___mcc_h3 = (_source0).cf0;
        let _12_cf0 = _11___mcc_h3;
        return (new BigNumber(917)).isLessThan((_dafny.ZERO).minus(_12_cf0));
      }
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))).Union(_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0))))).Union((_module.D11.create_DC29(_dafny.Set.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))))).dtor_cf39);
    };
    static fm8(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(636)), function (_13_i0) {
        return new BigNumber(437);
      }), _dafny.Seq.of(new BigNumber(-774)));
    };
    static fm9(p0, p1, globalState) {
      return (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality()), new BigNumber(768), new BigNumber(863))).Difference(_dafny.Set.fromElements(new BigNumber((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
          let _14_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0))), _14_v0)) {
            _coll2.push([_14_v0,_module.D9.create_DC23(new BigNumber((_dafny.Seq.UnicodeFromString("jtgcdb")).length))]);
          }
        }
        return _coll2;
      }()).length)));
    };
    static fm10(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.of(true, !(false), true), _dafny.Seq.of(true, !(true))));
    };
    static m0(p0, globalState) {
      let _15_v0;
      let _init0 = function (_16_i0) {
        return _module.__default.safeDivisionInt(_16_i0, new BigNumber(-52));
      };
      let _nw0 = Array((new BigNumber(28)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _15_v0 = _nw0;
      let _17_v1;
      _17_v1 = new BigNumber(473);
      let _index0 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
      (_15_v0)[_index0] = _module.__default.fm0(_17_v1, _17_v1, globalState);
      let _18_v2;
      _18_v2 = _dafny.Seq.of(p0);
      let _19_v3;
      _19_v3 = _dafny.Seq.of(_17_v1, new BigNumber((_18_v2).length));
      let _20_v4;
      _20_v4 = _dafny.Seq.of(_19_v3);
      let _index1 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
      (_15_v0)[_index1] = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_17_v1, _17_v1))).plus(new BigNumber((_20_v4).length));
      let _21_i1;
      _21_i1 = _dafny.ZERO;
      L0: {
        while ((_18_v2)[_module.__default.safeIndex((_17_v1).minus(_17_v1), new BigNumber((_18_v2).length))]) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_21_i1)) {
              break L0;
            }
            _21_i1 = (_21_i1).plus(_dafny.ONE);
            let _22_v5;
            let _init1 = function (_23_i2) {
              return _dafny.Seq.UnicodeFromString("kvewhxik");
            };
            let _nw1 = Array((new BigNumber(26)).toNumber());
            for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw1.length); _i0_1++) {
              _nw1[_i0_1] = _init1(new BigNumber(_i0_1));
            }
            _22_v5 = _nw1;
            let _24_v6;
            _24_v6 = _dafny.Seq.UnicodeFromString("mqwfqaq");
            let _index2 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_22_v5).length));
            (_22_v5)[_index2] = _dafny.Seq.Concat(_24_v6, _24_v6);
            let _25_v7;
            let _nw2 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
            _25_v7 = _nw2;
            let _26_v8;
            _26_v8 = _module.D0.create_DC0(_17_v1);
            let _27_v9;
            _27_v9 = _dafny.MultiSet.fromElements(p0, p0, p0, true, p0);
            let _index3 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_25_v7).length));
            (_25_v7)[_index3] = _dafny.Map.Empty.slice().updateUnsafe(_26_v8,_27_v9);
            let _28_v10;
            _28_v10 = _dafny.MultiSet.fromElements(new BigNumber(607));
            let _29_v11;
            _29_v11 = _dafny.Map.Empty.slice().updateUnsafe(_26_v8,_module.__default.fm3(p0, _28_v10, p0, globalState));
            let _index4 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_22_v5).length));
            let _index5 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_25_v7).length));
            let _rhs0 = _24_v6;
            let _rhs1 = (new BigNumber(106)).plus(_module.__default.safeDivisionInt((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]));
            let _rhs2 = (_29_v11).Merge(_29_v11);
            let _rhs3 = !((_27_v9).IsProperSubsetOf(_27_v9));
            let _lhs0 = _22_v5;
            let _lhs1 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_22_v5).length));
            let _lhs2 = globalState;
            let _lhs3 = _25_v7;
            let _lhs4 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_25_v7).length));
            let _lhs5 = globalState;
            _lhs0[_lhs1] = _rhs0;
            _lhs2.f4 = _rhs1;
            _lhs3[_lhs4] = _rhs2;
            _lhs5.f2 = _rhs3;
            let _index6 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
            (_15_v0)[_index6] = _17_v1;
            let _30_v13;
            let _init2 = ((_31_v1, _32_v10, _33_p0, _34_v0) => function (_35_i3) {
              return (function () {
                let _coll3 = new _dafny.Set();
                for (const _compr_3 of _dafny.IntegerRange(new BigNumber(671), new BigNumber(-131))) {
                  let _36_v12 = _compr_3;
                  if (((new BigNumber(671)).isLessThanOrEqualTo(_36_v12)) && ((_36_v12).isLessThan(new BigNumber(-131)))) {
                    _coll3.add((_36_v12).multipliedBy((_module.D0.create_DC1(new BigNumber((_32_v10).cardinality()), _33_p0, (_34_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_34_v0).length))])).dtor_cf3));
                  }
                }
                return _coll3;
              }()).IsDisjointFrom(_dafny.Set.fromElements(_31_v1));
            })(_17_v1, _28_v10, p0, _15_v0);
            let _nw3 = Array((new BigNumber(23)).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw3.length); _i0_2++) {
              _nw3[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _30_v13 = _nw3;
            let _index7 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_30_v13).length));
            (_30_v13)[_index7] = true;
            let _index8 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_30_v13).length));
            (_30_v13)[_index8] = !(_module.__default.fm6(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), function (_37_i4) {
              return new _dafny.CodePoint('k'.codePointAt(0));
            }), (_22_v5)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_22_v5).length))]), (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], globalState));
            let _38_v14;
            _38_v14 = _module.D0.create_DC1(_17_v1, p0, (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
            let _39_v15;
            _39_v15 = _dafny.Map.Empty.slice().updateUnsafe(_38_v14,(_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
            let _40_v16;
            let _nw4 = new _module.C0();
            _nw4.__ctor(_39_v15);
            _40_v16 = _nw4;
          }
        }
      }
      let _hi0 = _module.__default.safeDivisionInt((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
      for (let _41_i5 = (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]; _41_i5.isLessThan(_hi0); _41_i5 = _41_i5.plus(_dafny.ONE)) {
        let _42_v17;
        _42_v17 = _dafny.Set.fromElements(p0, p0);
        let _43_v18;
        _43_v18 = new _dafny.CodePoint('y'.codePointAt(0));
        let _44_v19;
        _44_v19 = _dafny.Map.Empty.slice().updateUnsafe(_42_v17,_43_v18);
        let _45_v20;
        _45_v20 = _module.D9.create_DC22(_42_v17);
        _44_v19 = (_44_v19).update((_45_v20).dtor_cf25, _43_v18);
        let _46_v21;
        _46_v21 = _module.D0.create_DC1(_41_i5, p0, _17_v1);
        let _47_v22;
        _47_v22 = _dafny.Map.Empty.slice().updateUnsafe(_46_v21,new BigNumber(-29));
        let _48_v23;
        let _nw5 = new _module.C0();
        _nw5.__ctor(_47_v22);
        _48_v23 = _nw5;
        (globalState).f3 = _41_i5;
        let _49_v24;
        _49_v24 = _dafny.Seq.UnicodeFromString("lfxeijna");
        let _50_v25;
        _50_v25 = _dafny.MultiSet.fromElements(_49_v24, _49_v24, _dafny.Seq.UnicodeFromString("mv"), _49_v24);
        let _51_v26;
        _51_v26 = _dafny.MultiSet.fromElements(_41_i5, (_19_v3)[_module.__default.safeIndex(new BigNumber((_50_v25).cardinality()), new BigNumber((_19_v3).length))]);
        let _52_v27;
        _52_v27 = _dafny.Map.Empty.slice().updateUnsafe((_17_v1).isEqualTo((((_51_v26).contains(_17_v1)) ? ((_51_v26).get(_17_v1)) : (_17_v1))),_module.__default.safeModuloInt(_41_i5, _17_v1));
        _52_v27 = (_52_v27).update(p0, _17_v1);
      }
      if (p0) {
        let _53_v29;
        _53_v29 = _dafny.Map.Empty.slice().updateUnsafe(_17_v1,p0);
        let _54_v30;
        _54_v30 = _module.D0.create_DC1((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], (((_53_v29).contains((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))])) ? ((_53_v29).get((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))])) : (p0)), _17_v1);
        let _55_v31;
        _55_v31 = _dafny.Map.Empty.slice().updateUnsafe(_54_v30,(_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
        let _56_v32;
        let _nw6 = new _module.C0();
        _nw6.__ctor(function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of (_55_v31).Keys.Elements) {
            let _57_v28 = _compr_4;
            if ((_55_v31).contains(_57_v28)) {
              _coll4.push([_57_v28,_17_v1]);
            }
          }
          return _coll4;
        }());
        _56_v32 = _nw6;
        let _58_v33;
        _58_v33 = new _dafny.CodePoint('l'.codePointAt(0));
        let _59_v34;
        _59_v34 = _module.D9.create_DC24(_module.__default.fm0((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], globalState), _56_v32, _58_v33, _17_v1);
        let _60_v35;
        _60_v35 = _dafny.Map.Empty.slice().updateUnsafe(_59_v34,(_dafny.ZERO).minus(_17_v1));
        let _61_v36;
        _61_v36 = _dafny.MultiSet.fromElements(p0);
        let _62_v37;
        _62_v37 = _dafny.Seq.UnicodeFromString("qaoscxjb");
        _60_v35 = (_60_v35).update(_59_v34, _module.__default.safeModuloInt((((_61_v36).contains(p0)) ? ((_61_v36).get(p0)) : (new BigNumber((_62_v37).length))), (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]));
        (globalState).f2 = p0;
        (globalState).f2 = p0;
        let _63_v38;
        _63_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_58_v33);
        let _64_v39;
        _64_v39 = _module.D8.create_DC20(_58_v33);
        let _65_v40;
        let _nw7 = Array((new BigNumber(26)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
        _nw7[(_dafny.ONE).toNumber()] = _58_v33;
        _nw7[(new BigNumber(2)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(3)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(4)).toNumber()] = (_module.D9.create_DC24((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], _56_v32, new _dafny.CodePoint('g'.codePointAt(0)), _17_v1)).dtor_cf29;
        _nw7[(new BigNumber(5)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(6)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(7)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw7[(new BigNumber(9)).toNumber()] = (((_63_v38).contains(p0)) ? ((_63_v38).get(p0)) : (_58_v33));
        _nw7[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw7[(new BigNumber(11)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(12)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(13)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(14)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
        _nw7[(new BigNumber(16)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(17)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(18)).toNumber()] = (_64_v39).dtor_cf22;
        _nw7[(new BigNumber(19)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(20)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(21)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(22)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw7[(new BigNumber(24)).toNumber()] = _58_v33;
        _nw7[(new BigNumber(25)).toNumber()] = _58_v33;
        _65_v40 = _nw7;
        let _index9 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_65_v40).length));
        (_65_v40)[_index9] = _58_v33;
        let _66_v41;
        _66_v41 = _dafny.Set.fromElements((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
        let _index10 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
        let _index11 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_65_v40).length));
        let _rhs4 = ((p0) ? (new BigNumber(-33)) : (new BigNumber((_66_v41).length)));
        let _rhs5 = _58_v33;
        let _rhs6 = _17_v1;
        let _lhs6 = _15_v0;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
        let _lhs8 = _65_v40;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_65_v40).length));
        let _lhs10 = globalState;
        _lhs6[_lhs7] = _rhs4;
        _lhs8[_lhs9] = _rhs5;
        _lhs10.f3 = _rhs6;
        let _67_v42;
        _67_v42 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_62_v37) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-146)), ((_68_v40) => function (_69_i6) {
          return (_68_v40)[_module.__default.safeIndex(new BigNumber(167), new BigNumber((_68_v40).length))];
        })(_65_v40)))),((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]).minus(new BigNumber(476)));
        _67_v42 = _67_v42;
      } else {
        if (p0) {
          let _70_v43;
          _70_v43 = _dafny.Seq.UnicodeFromString("cnu");
          let _71_v44;
          _71_v44 = new _dafny.CodePoint('n'.codePointAt(0));
          let _72_v45;
          _72_v45 = _module.D0.create_DC1((new BigNumber(343)).minus(_17_v1), _dafny.Seq.IsPrefixOf(_70_v43, _dafny.Seq.update(_70_v43, _module.__default.safeIndex(new BigNumber(335), new BigNumber((_70_v43).length)), _71_v44)), new BigNumber((_18_v2).length));
          _72_v45 = _72_v45;
          (globalState).f4 = (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))];
          let _73_v46;
          _73_v46 = _module.D0.create_DC3();
          let _74_v47;
          _74_v47 = _dafny.Map.Empty.slice().updateUnsafe(_73_v46,_15_v0);
          _74_v47 = (_74_v47).update(_73_v46, _15_v0);
          let _75_v48;
          _75_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _76_v50;
          _76_v50 = _dafny.Map.Empty.slice().updateUnsafe(_72_v45,_19_v3);
          let _77_v51;
          let _nw8 = new _module.C0();
          _nw8.__ctor(function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_76_v50).Keys.Elements) {
              let _78_v49 = _compr_5;
              if ((_76_v50).contains(_78_v49)) {
                _coll5.push([_78_v49,new BigNumber((_70_v43).length)]);
              }
            }
            return _coll5;
          }());
          _77_v51 = _nw8;
          let _79_v52;
          _79_v52 = _module.D10.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(_75_v48,_77_v51));
          _17_v1 = (new BigNumber(((_79_v52).dtor_cf31).length)).multipliedBy(_17_v1);
          _71_v44 = _71_v44;
        } else {
          let _80_v53;
          _80_v53 = _dafny.Seq.UnicodeFromString("kqlo");
          let _81_v54;
          let _init3 = ((_82_v2) => function (_83_i7) {
            return _dafny.Seq.Concat(_82_v2, _82_v2);
          })(_18_v2);
          let _nw9 = Array((_dafny.ONE).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw9.length); _i0_3++) {
            _nw9[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _81_v54 = _nw9;
          let _index12 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_81_v54).length));
          (_81_v54)[_index12] = _dafny.Seq.of(!(true), p0, p0, p0, p0);
          let _index13 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_81_v54).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
          let _rhs7 = ((p0) ? (_80_v53) : (_80_v53));
          let _rhs8 = (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))];
          let _rhs9 = _18_v2;
          let _rhs10 = (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))];
          let _lhs11 = globalState;
          let _lhs12 = _81_v54;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_81_v54).length));
          let _lhs14 = _15_v0;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
          _80_v53 = _rhs7;
          _lhs11.f4 = _rhs8;
          _lhs12[_lhs13] = _rhs9;
          _lhs14[_lhs15] = _rhs10;
          let _84_v55;
          _84_v55 = _dafny.Map.Empty.slice().updateUnsafe(_17_v1,(_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
          let _85_v57;
          _85_v57 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], new BigNumber(923), globalState),new BigNumber((_dafny.MultiSet.fromElements(_17_v1, new BigNumber(-135))).cardinality())));
          let _86_v58;
          _86_v58 = _dafny.MultiSet.fromElements(_84_v55, _84_v55, function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(686), new BigNumber(758))) {
              let _87_v56 = _compr_6;
              if (((new BigNumber(686)).isLessThanOrEqualTo(_87_v56)) && ((_87_v56).isLessThan(new BigNumber(758)))) {
                _coll6.push([(_87_v56).plus(new BigNumber(387)),_17_v1]);
              }
            }
            return _coll6;
          }(), (_85_v57)[_module.__default.safeIndex(_17_v1, new BigNumber((_85_v57).length))], _84_v55);
          (globalState).f2 = ((_86_v58).Union(_86_v58)).IsProperSubsetOf(_86_v58);
          let _88_v59;
          let _nw10 = Array((new BigNumber(18)).toNumber()).fill(false);
          _88_v59 = _nw10;
          let _89_v60;
          _89_v60 = _module.D7.create_DC17(_15_v0);
          let _90_v61;
          _90_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(451),(_89_v60).dtor_cf19);
          let _index15 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_88_v59).length));
          (_88_v59)[_index15] = !(_90_v61).equals(_90_v61);
          let _index16 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_88_v59).length));
          (_88_v59)[_index16] = p0;
          (globalState).f2 = (_88_v59)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_88_v59).length))];
          (globalState).f2 = !((_dafny.ZERO).minus((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))])).isEqualTo(new BigNumber(621));
        }
        _19_v3 = _dafny.Seq.of(new BigNumber(494));
        _15_v0 = _15_v0;
        let _91_v62;
        _91_v62 = _module.D0.create_DC1((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))], (_18_v2)[_module.__default.safeIndex(_17_v1, new BigNumber((_18_v2).length))], (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
        let _92_v63;
        _92_v63 = _dafny.Map.Empty.slice().updateUnsafe(_91_v62,_17_v1);
        let _93_v64;
        let _nw11 = new _module.C0();
        _nw11.__ctor(_92_v63);
        _93_v64 = _nw11;
        let _94_v65;
        _94_v65 = new _dafny.CodePoint('k'.codePointAt(0));
        let _95_v66;
        _95_v66 = _module.D9.create_DC24(_17_v1, _93_v64, _94_v65, (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]);
        let _96_v67;
        _96_v67 = _dafny.Map.Empty.slice().updateUnsafe(_95_v66,p0);
        let _97_v68;
        let _nw12 = Array((new BigNumber(9)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _18_v2;
        _nw12[(_dafny.ONE).toNumber()] = _18_v2;
        _nw12[(new BigNumber(2)).toNumber()] = _18_v2;
        _nw12[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(p0, p0);
        _nw12[(new BigNumber(4)).toNumber()] = ((p0) ? (_18_v2) : (_18_v2));
        _nw12[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(!(false), (((_96_v67).contains(_95_v66)) ? ((_96_v67).get(_95_v66)) : (p0)));
        _nw12[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(p0, p0);
        _nw12[(new BigNumber(7)).toNumber()] = _18_v2;
        _nw12[(new BigNumber(8)).toNumber()] = _18_v2;
        _97_v68 = _nw12;
        let _98_v69;
        let _nw13 = Array((new BigNumber(21)).toNumber()).fill(_module.D6.Default());
        _98_v69 = _nw13;
        let _99_v70;
        _99_v70 = _module.D6.create_DC15(_17_v1, _17_v1);
        let _index17 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_98_v69).length));
        (_98_v69)[_index17] = _99_v70;
        let _pat_let_tv0 = _17_v1;
        let _index18 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_98_v69).length));
        let _rhs11 = _97_v68;
        let _rhs12 = function (_pat_let0_0) {
          return function (_100_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_101_dt__update_hcf17_h0) {
                return _module.D6.create_DC15((_100_dt__update__tmp_h0).dtor_cf16, _101_dt__update_hcf17_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_99_v70);
        let _lhs16 = _98_v69;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_98_v69).length));
        _97_v68 = _rhs11;
        _lhs16[_lhs17] = _rhs12;
        let _102_v71;
        let _nw14 = Array((new BigNumber(26)).toNumber());
        _102_v71 = _nw14;
        let _103_v72;
        _103_v72 = _dafny.Map.Empty.slice().updateUnsafe(_102_v71,p0);
        let _index19 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
        let _index20 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
        let _rhs13 = (_dafny.ZERO).minus((((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]).multipliedBy((_dafny.ZERO).minus(_17_v1))).minus((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]));
        let _rhs14 = (p0) === (!(((_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))]).isLessThanOrEqualTo(new BigNumber((_103_v72).length))));
        let _rhs15 = (_15_v0)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length))];
        let _rhs16 = _module.__default.fm0(_17_v1, _17_v1, globalState);
        let _lhs18 = globalState;
        let _lhs19 = globalState;
        let _lhs20 = _15_v0;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
        let _lhs22 = _15_v0;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_15_v0).length));
        _lhs18.f4 = _rhs13;
        _lhs19.f2 = _rhs14;
        _lhs20[_lhs21] = _rhs15;
        _lhs22[_lhs23] = _rhs16;
      }
      let _104_v73;
      _104_v73 = new _dafny.CodePoint('f'.codePointAt(0));
      let _105_v74;
      _105_v74 = _dafny.Seq.of(_104_v73);
      let _hi1 = _17_v1;
      for (let _106_i8 = ((p0) ? (_17_v1) : (new BigNumber((_105_v74).length))); _106_i8.isLessThan(_hi1); _106_i8 = _106_i8.plus(_dafny.ONE)) {
        (globalState).f2 = !(p0) || (p0);
        let _107_v75;
        _107_v75 = _dafny.Set.fromElements(((p0) ? (p0) : (p0)));
        _107_v75 = _107_v75;
        let _108_v76;
        _108_v76 = _dafny.MultiSet.fromElements(false);
        _108_v76 = (_dafny.MultiSet.fromElements(p0, p0, p0, p0, true)).Union((_108_v76).Intersect(_108_v76));
        let _109_v77;
        _109_v77 = _dafny.Set.fromElements(_17_v1, _106_i8);
        let _110_v78;
        _110_v78 = _module.D0.create_DC1(_106_i8, p0, new BigNumber((_109_v77).length));
        let _111_v79;
        _111_v79 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let2_0) {
          return function (_112_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_113_dt__update_hcf3_h0) {
                return _module.D0.create_DC1((_112_dt__update__tmp_h1).dtor_cf1, (_112_dt__update__tmp_h1).dtor_cf2, _113_dt__update_hcf3_h0);
              }(_pat_let3_0);
            }(_106_i8);
          }(_pat_let2_0);
        }(_110_v78),_17_v1);
        let _114_v80;
        let _nw15 = new _module.C0();
        _nw15.__ctor(_111_v79);
        _114_v80 = _nw15;
        let _115_v81;
        let _nw16 = Array((new BigNumber(2)).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = _114_v80;
        _nw16[(_dafny.ONE).toNumber()] = _114_v80;
        _115_v81 = _nw16;
        _115_v81 = ((p0) ? (_115_v81) : (_115_v81));
      }
      let _116_v82;
      _116_v82 = _dafny.Set.fromElements(_19_v3);
      (globalState).f2 = (((_116_v82).IsProperSubsetOf(_116_v82)) ? (p0) : ((_17_v1).isLessThan(_17_v1)));
      return;
    }
    static Main(__noArgsParameter) {
      let _117_v0;
      _117_v0 = new BigNumber(877);
      let _118_v1;
      _118_v1 = _dafny.Set.fromElements((_dafny.ZERO).minus(_117_v0));
      let _119_v2;
      _119_v2 = _dafny.Seq.UnicodeFromString("yxsvjmbqp");
      let _120_v3;
      _120_v3 = _dafny.Set.fromElements(_119_v2, _119_v2);
      let _121_v4;
      let _nw17 = Array((new BigNumber(20)).toNumber());
      _nw17[(_dafny.ZERO).toNumber()] = _117_v0;
      _nw17[(_dafny.ONE).toNumber()] = _117_v0;
      _nw17[(new BigNumber(2)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(3)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(4)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(5)).toNumber()] = new BigNumber(782);
      _nw17[(new BigNumber(6)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(7)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(8)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(9)).toNumber()] = (_module.D0.create_DC0(_117_v0)).dtor_cf0;
      _nw17[(new BigNumber(10)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(11)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(12)).toNumber()] = new BigNumber((_118_v1).length);
      _nw17[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_120_v3).length));
      _nw17[(new BigNumber(14)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(15)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(16)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_117_v0);
      _nw17[(new BigNumber(18)).toNumber()] = _117_v0;
      _nw17[(new BigNumber(19)).toNumber()] = new BigNumber(217);
      _121_v4 = _nw17;
      let _122_v5;
      let _init4 = ((_123_v0) => function (_124_i0) {
        return (_module.D0.create_DC1(_123_v0, true, _123_v0)).dtor_cf2;
      })(_117_v0);
      let _nw18 = Array((new BigNumber(12)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw18.length); _i0_4++) {
        _nw18[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _122_v5 = _nw18;
      let _125_globalState;
      let _nw19 = new _module.GlobalState();
      _nw19.__ctor(false, true, true, new BigNumber(-701), new BigNumber(757), new BigNumber(943), _121_v4, true, _121_v4, _122_v5, new BigNumber(685), false, _119_v2, new BigNumber(826), true);
      _125_globalState = _nw19;
      let _126_i1;
      _126_i1 = _dafny.ZERO;
      L1: {
        while ((_117_v0).isEqualTo(_117_v0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_126_i1)) {
              break L1;
            }
            _126_i1 = (_126_i1).plus(_dafny.ONE);
            let _127_v6;
            _127_v6 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_117_v0);
            let _index21 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
            (_121_v4)[_index21] = new BigNumber(((_127_v6).Merge(_127_v6)).length);
            let _128_v7;
            _128_v7 = true;
            let _129_v8;
            _129_v8 = new _dafny.CodePoint('i'.codePointAt(0));
            let _index22 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
            (_121_v4)[_index22] = _module.__default.fm0(((_128_v7) ? (_117_v0) : (new BigNumber(-650))), new BigNumber((_dafny.Seq.update(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(438)), function (_130_i2) {
              return new _dafny.CodePoint('u'.codePointAt(0));
            })) : (_119_v2)), _module.__default.safeIndex(_117_v0, new BigNumber((((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(438)), function (_131_i2) {
              return new _dafny.CodePoint('u'.codePointAt(0));
            })) : (_119_v2))).length)), _129_v8)).length), _125_globalState);
            let _index23 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
            (_121_v4)[_index23] = _module.__default.fm0(_117_v0, _117_v0, _125_globalState);
            if ((_117_v0).isLessThan(_117_v0)) {
              let _132_v9;
              _132_v9 = _dafny.MultiSet.fromElements(_117_v0);
              let _133_v10;
              _133_v10 = _dafny.Seq.of(_119_v2);
              let _134_v11;
              _134_v11 = _dafny.MultiSet.fromElements(_module.__default.fm0((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], new BigNumber((_132_v9).cardinality()), _125_globalState), (_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], new BigNumber(((_133_v10)[_module.__default.safeIndex((((_132_v9).contains(_117_v0)) ? ((_132_v9).get(_117_v0)) : (_117_v0)), new BigNumber((_133_v10).length))]).length), _117_v0);
              let _135_v12;
              let _init5 = ((_136_v2) => function (_137_i3) {
                return _136_v2;
              })(_119_v2);
              let _nw20 = Array((new BigNumber(20)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw20.length); _i0_5++) {
                _nw20[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _135_v12 = _nw20;
              let _index24 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_135_v12).length));
              (_135_v12)[_index24] = _119_v2;
              let _138_v13;
              _138_v13 = _119_v2;
              let _index25 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
              let _index26 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_135_v12).length));
              let _rhs17 = _117_v0;
              let _rhs18 = _134_v11;
              let _rhs19 = _119_v2;
              let _rhs20 = _128_v7;
              let _rhs21 = (_138_v13);
              let _lhs24 = _121_v4;
              let _lhs25 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
              let _lhs26 = _125_globalState;
              let _lhs27 = _135_v12;
              let _lhs28 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_135_v12).length));
              _lhs24[_lhs25] = _rhs17;
              _132_v9 = _rhs18;
              _119_v2 = _rhs19;
              _lhs26.f2 = _rhs20;
              _lhs27[_lhs28] = _rhs21;
              let _139_v14;
              _139_v14 = _dafny.Seq.of((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], (_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], (_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], _117_v0);
              (_125_globalState).f2 = !_dafny.Seq.contains(_139_v14, _117_v0);
              let _140_v15;
              _140_v15 = _dafny.MultiSet.fromElements(_119_v2, (_135_v12)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_135_v12).length))]);
              let _index27 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_135_v12).length));
              let _rhs22 = _dafny.Seq.Concat(_dafny.Seq.update(_119_v2, _module.__default.safeIndex((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], new BigNumber((_119_v2).length)), _129_v8), (_138_v13));
              let _rhs23 = !(_117_v0).isEqualTo((_117_v0).multipliedBy((_139_v14)[_module.__default.safeIndex(_117_v0, new BigNumber((_139_v14).length))]));
              let _rhs24 = (_140_v15).IsProperSubsetOf(_140_v15);
              let _rhs25 = _117_v0;
              let _lhs29 = _135_v12;
              let _lhs30 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_135_v12).length));
              let _lhs31 = _125_globalState;
              let _lhs32 = _125_globalState;
              let _lhs33 = _125_globalState;
              _lhs29[_lhs30] = _rhs22;
              _lhs31.f2 = _rhs23;
              _lhs32.f2 = _rhs24;
              _lhs33.f4 = _rhs25;
              let _141_v16;
              _141_v16 = _module.D0.create_DC1((_dafny.ZERO).minus((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]), _128_v7, new BigNumber((_119_v2).length));
              let _142_v17;
              let _nw21 = new _module.C0();
              _nw21.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_141_v16,(_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]));
              _142_v17 = _nw21;
              (_125_globalState).f3 = ((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]).minus((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]);
            } else {
              _module.__default.m0(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("lkmnn"), _119_v2), _125_globalState);
              let _143_v18;
              let _nw22 = Array((new BigNumber(3)).toNumber());
              _nw22[(_dafny.ZERO).toNumber()] = ((_128_v7) ? (_122_v5) : (_122_v5));
              _nw22[(_dafny.ONE).toNumber()] = _122_v5;
              _nw22[(new BigNumber(2)).toNumber()] = _122_v5;
              _143_v18 = _nw22;
              let _index28 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_143_v18).length));
              (_143_v18)[_index28] = _122_v5;
              let _index29 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_143_v18).length));
              (_143_v18)[_index29] = _122_v5;
              let _144_v19;
              _144_v19 = _module.D0.create_DC2();
              _144_v19 = _144_v19;
              let _145_v20;
              let _init6 = ((_146_v7, _147_v8, _148_v2) => function (_149_i4) {
                return ((_146_v7) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), ((_150_v8) => function (_151_i5) {
                  return _150_v8;
                })(_147_v8))) : (_148_v2));
              })(_128_v7, _129_v8, _119_v2);
              let _nw23 = Array((new BigNumber(4)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw23.length); _i0_6++) {
                _nw23[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _145_v20 = _nw23;
              let _152_v21;
              _152_v21 = _dafny.Map.Empty.slice().updateUnsafe(!(_128_v7),_dafny.Seq.UnicodeFromString("hcnuuyv"));
              let _index30 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_145_v20).length));
              (_145_v20)[_index30] = (((_152_v21).contains(_128_v7)) ? ((_152_v21).get(_128_v7)) : (_module.__default.fm1(_125_globalState)));
              let _index31 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_145_v20).length));
              (_145_v20)[_index31] = _dafny.Seq.Concat(_119_v2, _119_v2);
              let _153_v22;
              _153_v22 = _dafny.MultiSet.fromElements((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], (_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]);
              _module.__default.m0(!(_153_v22).contains(new BigNumber(577)), _125_globalState);
            }
            if ((_117_v0).isEqualTo((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))])) {
              (_125_globalState).f3 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeModuloInt((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], _117_v0)).plus(_117_v0))));
              (_125_globalState).f4 = _117_v0;
              let _index32 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
              (_121_v4)[_index32] = (_dafny.ZERO).minus(_117_v0);
              let _154_v23;
              let _init7 = ((_155_v7, _156_v4) => function (_157_i6) {
                return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-2))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_155_v7,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_156_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_156_v4).length))],_155_v7)).length)));
              })(_128_v7, _121_v4);
              let _nw24 = Array((new BigNumber(13)).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw24.length); _i0_7++) {
                _nw24[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _154_v23 = _nw24;
              let _158_v24;
              _158_v24 = _dafny.Map.Empty.slice().updateUnsafe(_128_v7,_117_v0);
              let _index33 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_154_v23).length));
              (_154_v23)[_index33] = _158_v24;
              let _index34 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_154_v23).length));
              (_154_v23)[_index34] = (_158_v24).update(_128_v7, new BigNumber(410));
              let _159_v25;
              _159_v25 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_dafny.MultiSet.fromElements(new BigNumber(300), new BigNumber(58)));
              let _160_v26;
              _160_v26 = _dafny.Seq.of(_117_v0, _117_v0);
              let _161_v27;
              _161_v27 = _module.D2.create_DC5(_dafny.MultiSet.FromArray(_160_v26));
              _159_v25 = (_159_v25).update(_117_v0, (_161_v27).dtor_cf5);
            } else {
              let _162_v29;
              _162_v29 = _dafny.MultiSet.fromElements(_128_v7);
              let _163_v30;
              _163_v30 = _module.D0.create_DC1((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], _128_v7, (_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]);
              let _164_v31;
              _164_v31 = _dafny.Map.Empty.slice().updateUnsafe(_163_v30,_117_v0);
              let _165_v32;
              let _nw25 = new _module.C0();
              _nw25.__ctor(_164_v31);
              _165_v32 = _nw25;
              let _166_v33;
              _166_v33 = _dafny.MultiSet.fromElements(_165_v32, _165_v32);
              let _167_v34;
              _167_v34 = _module.D0.create_DC1(new BigNumber((_166_v33).cardinality()), _128_v7, _117_v0);
              let _168_v35;
              _168_v35 = _dafny.MultiSet.fromElements(_module.D0.create_DC1((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))], _128_v7, (((_162_v29).contains(_128_v7)) ? ((_162_v29).get(_128_v7)) : (_117_v0))), _167_v34);
              let _169_v36;
              let _nw26 = new _module.C0();
              _nw26.__ctor(function () {
                let _coll7 = new _dafny.Map();
                for (const _compr_7 of (_168_v35).Elements) {
                  let _170_v28 = _compr_7;
                  if ((_168_v35).contains(_170_v28)) {
                    _coll7.push([_170_v28,(_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]]);
                  }
                }
                return _coll7;
              }());
              _169_v36 = _nw26;
              let _171_v37;
              let _nw27 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
              _171_v37 = _nw27;
              let _172_v38;
              _172_v38 = _dafny.Map.Empty.slice().updateUnsafe(_165_v32,_171_v37);
              let _173_v39;
              let _nw28 = new _module.C0();
              _nw28.__ctor(_164_v31);
              _173_v39 = _nw28;
              let _174_v40;
              _174_v40 = _dafny.Set.fromElements(_173_v39);
              let _rhs26 = _128_v7;
              let _rhs27 = _128_v7;
              let _rhs28 = ((_174_v40).Difference(_174_v40)).IsDisjointFrom((_174_v40).Intersect(_dafny.Set.fromElements(_173_v39, _173_v39, _173_v39)));
              let _rhs29 = _128_v7;
              let _rhs30 = (_172_v38).Merge((_172_v38));
              let _lhs34 = _125_globalState;
              let _lhs35 = _125_globalState;
              _128_v7 = _rhs26;
              _lhs34.f2 = _rhs27;
              _128_v7 = _rhs28;
              _lhs35.f2 = _rhs29;
              _172_v38 = _rhs30;
              let _index35 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
              let _rhs31 = _128_v7;
              let _rhs32 = ((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]).plus((_dafny.ZERO).minus((_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]));
              let _lhs36 = _125_globalState;
              let _lhs37 = _121_v4;
              let _lhs38 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length));
              _lhs36.f2 = _rhs31;
              _lhs37[_lhs38] = _rhs32;
              let _175_v41;
              _175_v41 = _dafny.Set.fromElements(_128_v7, _128_v7, _128_v7, _128_v7, _128_v7);
              let _176_v42;
              _176_v42 = _dafny.Seq.of(_175_v41, _175_v41, _175_v41, _175_v41);
              let _177_v43;
              _177_v43 = _dafny.Map.Empty.slice().updateUnsafe(_128_v7,_dafny.MultiSet.FromArray(_176_v42));
              let _178_v44;
              _178_v44 = _dafny.MultiSet.fromElements(_175_v41, _175_v41, _175_v41);
              let _179_v45;
              _179_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_175_v41).length),_178_v44);
              _177_v43 = (_177_v43).update(_128_v7, (((_179_v45).contains(_117_v0)) ? ((_179_v45).get(_117_v0)) : (_178_v44)));
              let _180_v46;
              let _nw29 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
              _180_v46 = _nw29;
              _180_v46 = _180_v46;
              let _181_v47;
              let _nw30 = new _module.C0();
              _nw30.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_128_v7, _125_globalState),(_121_v4)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_121_v4).length))]));
              _181_v47 = _nw30;
            }
          }
        }
      }
      let _182_v48;
      _182_v48 = false;
      let _183_v49;
      _183_v49 = _dafny.Seq.of(_182_v48);
      if ((_182_v48) === (_dafny.Seq.IsProperPrefixOf(_183_v49, _dafny.Seq.of(true)))) {
        let _184_v50;
        _184_v50 = _dafny.MultiSet.fromElements(_119_v2, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eueosg"), _119_v2), _119_v2);
        _117_v0 = new BigNumber((_184_v50).cardinality());
        let _185_v51;
        _185_v51 = new _dafny.CodePoint('y'.codePointAt(0));
        _185_v51 = _185_v51;
        let _186_v52;
        _186_v52 = _module.D0.create_DC1(_117_v0, _182_v48, _module.__default.fm0(_117_v0, _117_v0, _125_globalState));
        let _187_v54;
        _187_v54 = _dafny.Map.Empty.slice().updateUnsafe(_186_v52,_117_v0);
        let _188_v55;
        let _nw31 = new _module.C0();
        _nw31.__ctor((_dafny.Map.Empty.slice().updateUnsafe(_186_v52,_117_v0)).Merge(function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_187_v54).Keys.Elements) {
            let _189_v53 = _compr_8;
            if ((_187_v54).contains(_189_v53)) {
              _coll8.push([_189_v53,(_186_v52).dtor_cf3]);
            }
          }
          return _coll8;
        }()));
        _188_v55 = _nw31;
        let _190_v56;
        _190_v56 = _dafny.Seq.of(_117_v0);
        if (!_dafny.Seq.contains(_190_v56, _117_v0)) {
          let _index36 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length));
          (_122_v5)[_index36] = _182_v48;
          let _191_v57;
          _191_v57 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_117_v0);
          let _192_v58;
          let _nw32 = new _module.C0();
          _nw32.__ctor(_187_v54);
          _192_v58 = _nw32;
          let _193_v59;
          _193_v59 = _dafny.Seq.of(_192_v58, _192_v58);
          let _194_v60;
          _194_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_191_v57).length), _117_v0),_dafny.Seq.Concat((_module.D4.create_DC9(_193_v59)).dtor_cf11, _193_v59));
          let _195_v61;
          _195_v61 = _dafny.Map.Empty.slice().updateUnsafe(_192_v58,_117_v0);
          let _196_v62;
          _196_v62 = _dafny.Map.Empty.slice().updateUnsafe(_192_v58,((_182_v48) ? (_195_v61) : (_195_v61)));
          let _index37 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length));
          let _rhs33 = _182_v48;
          let _rhs34 = (_dafny.Map.Empty.slice().updateUnsafe(_117_v0,_193_v59)).Merge(_194_v60);
          let _rhs35 = _196_v62;
          let _lhs39 = _122_v5;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length));
          _lhs39[_lhs40] = _rhs33;
          _194_v60 = _rhs34;
          _196_v62 = _rhs35;
          let _197_v63;
          _197_v63 = _187_v54;
          let _198_v64;
          let _nw33 = new _module.C0();
          _nw33.__ctor((_197_v63));
          _198_v64 = _nw33;
          _198_v64 = _198_v64;
          let _index38 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length));
          (_122_v5)[_index38] = (((_117_v0).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(_117_v0)).cardinality()))) ? ((_122_v5)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length))]) : (_182_v48));
          let _nw34 = new _module.C0();
          _nw34.__ctor((_188_v55).f15);
          _192_v58 = _nw34;
          let _index39 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length));
          (_122_v5)[_index39] = ((!_dafny.Seq.contains(_119_v2, _185_v51)) ? (false) : ((_122_v5)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_122_v5).length))]));
        } else {
          let _index40 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length));
          (_121_v4)[_index40] = new BigNumber((_190_v56).length);
          let _199_v65;
          _199_v65 = _dafny.MultiSet.fromElements(_182_v48);
          let _200_v66;
          _200_v66 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_dafny.MultiSet.fromElements(_182_v48));
          let _201_v67;
          _201_v67 = _dafny.MultiSet.fromElements(_117_v0);
          let _202_v68;
          let _nw35 = Array((new BigNumber(22)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(!(true), _182_v48);
          _nw35[(_dafny.ONE).toNumber()] = _199_v65;
          _nw35[(new BigNumber(2)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(3)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(4)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_183_v49);
          _nw35[(new BigNumber(6)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(7)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(8)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_183_v49);
          _nw35[(new BigNumber(10)).toNumber()] = (((_200_v66).contains(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_207_v0) => function (_208_i7) {
            return _207_v0;
          })(_117_v0)), _module.__default.safeIndex(_117_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_209_v0) => function (_210_i7) {
            return _209_v0;
          })(_117_v0))).length)), _117_v0)).length))) ? ((_200_v66).get(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_203_v0) => function (_204_i7) {
            return _203_v0;
          })(_117_v0)), _module.__default.safeIndex(_117_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_205_v0) => function (_206_i7) {
            return _205_v0;
          })(_117_v0))).length)), _117_v0)).length))) : (_199_v65));
          _nw35[(new BigNumber(11)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(12)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(13)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(14)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(15)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(16)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.FromArray(_183_v49);
          _nw35[(new BigNumber(18)).toNumber()] = _199_v65;
          _nw35[(new BigNumber(19)).toNumber()] = _module.__default.fm3(true, _201_v67, _182_v48, _125_globalState);
          _nw35[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.FromArray(_183_v49);
          _nw35[(new BigNumber(21)).toNumber()] = _199_v65;
          _202_v68 = _nw35;
          let _211_v69;
          _211_v69 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,_122_v5);
          let _212_v70;
          _212_v70 = _dafny.Map.Empty.slice().updateUnsafe(_202_v68,(((_211_v69).contains(_182_v48)) ? ((_211_v69).get(_182_v48)) : (_122_v5)));
          let _213_v71;
          _213_v71 = _dafny.Seq.of(_122_v5);
          let _214_v72;
          let _nw36 = Array((new BigNumber(25)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _122_v5;
          _nw36[(_dafny.ONE).toNumber()] = _122_v5;
          _nw36[(new BigNumber(2)).toNumber()] = (((_212_v70).contains(_202_v68)) ? ((_212_v70).get(_202_v68)) : (_122_v5));
          _nw36[(new BigNumber(3)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(4)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(5)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(6)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(7)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(8)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(9)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(10)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(11)).toNumber()] = (_213_v71)[_module.__default.safeIndex(_117_v0, new BigNumber((_213_v71).length))];
          _nw36[(new BigNumber(12)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(13)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(14)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(15)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(16)).toNumber()] = (_module.D6.create_DC14(_122_v5)).dtor_cf15;
          _nw36[(new BigNumber(17)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(18)).toNumber()] = ((_182_v48) ? (_122_v5) : (_122_v5));
          _nw36[(new BigNumber(19)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(20)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(21)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(22)).toNumber()] = ((_182_v48) ? (_122_v5) : (_122_v5));
          _nw36[(new BigNumber(23)).toNumber()] = _122_v5;
          _nw36[(new BigNumber(24)).toNumber()] = _122_v5;
          _214_v72 = _nw36;
          let _index41 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_214_v72).length));
          (_214_v72)[_index41] = _122_v5;
          let _215_v73;
          _215_v73 = _dafny.Map.Empty.slice().updateUnsafe(_119_v2,_117_v0);
          let _index42 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length));
          let _index43 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_214_v72).length));
          let _rhs36 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_117_v0), _117_v0), (((_215_v73).contains(_119_v2)) ? ((_215_v73).get(_119_v2)) : (_117_v0)));
          let _rhs37 = _119_v2;
          let _rhs38 = (new BigNumber(209)).minus(new BigNumber(602));
          let _rhs39 = _122_v5;
          let _lhs41 = _125_globalState;
          let _lhs42 = _121_v4;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length));
          let _lhs44 = _214_v72;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_214_v72).length));
          _lhs41.f3 = _rhs36;
          _119_v2 = _rhs37;
          _lhs42[_lhs43] = _rhs38;
          _lhs44[_lhs45] = _rhs39;
          let _216_v74;
          let _nw37 = new _module.C0();
          _nw37.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_186_v52,_117_v0));
          _216_v74 = _nw37;
          let _217_v75;
          _217_v75 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_119_v2);
          _217_v75 = (_217_v75).update((_121_v4)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_218_v51) => function (_219_i8) {
            return _218_v51;
          })(_185_v51)));
          (_125_globalState).f4 = (_dafny.ZERO).minus(((_121_v4)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length))]).plus(new BigNumber(820)));
          let _index44 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length));
          let _rhs40 = _190_v56;
          let _rhs41 = new BigNumber((_module.__default.fm4(!(_182_v48), _182_v48, _125_globalState)).cardinality());
          let _lhs46 = _121_v4;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_121_v4).length));
          _190_v56 = _rhs40;
          _lhs46[_lhs47] = _rhs41;
        }
        (_125_globalState).f2 = _182_v48;
      } else {
        let _220_v76;
        _220_v76 = new _dafny.CodePoint('x'.codePointAt(0));
        let _221_v77;
        _221_v77 = _dafny.Map.Empty.slice().updateUnsafe(_220_v76,_dafny.areEqual(_119_v2, _119_v2));
        let _222_v78;
        _222_v78 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_182_v48);
        let _223_v79;
        _223_v79 = _module.D0.create_DC1(new BigNumber(116), _182_v48, _117_v0);
        let _224_v80;
        _224_v80 = _dafny.Map.Empty.slice().updateUnsafe(_223_v79,_117_v0);
        let _225_v81;
        let _nw38 = new _module.C0();
        _nw38.__ctor((_224_v80).update(_223_v79, _117_v0));
        _225_v81 = _nw38;
        let _226_v82;
        _226_v82 = _dafny.Map.Empty.slice().updateUnsafe(_225_v81,_225_v81);
        let _227_v83;
        _227_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_225_v81,_225_v81),_117_v0);
        let _228_v84;
        _228_v84 = _dafny.Seq.of(new BigNumber((_183_v49).length));
        let _rhs42 = ((!(_227_v83).contains((_226_v82).update(_225_v81, _225_v81))) ? (_221_v77) : (_221_v77));
        let _rhs43 = (_222_v78).update(_117_v0, (function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(17), new BigNumber(599))) {
            let _229_v86 = _compr_9;
            if (((new BigNumber(17)).isLessThanOrEqualTo(_229_v86)) && ((_229_v86).isLessThan(new BigNumber(599)))) {
              _coll9.add((_229_v86).multipliedBy(new BigNumber(904)));
            }
          }
          return _coll9;
        }()).IsProperSubsetOf(function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_dafny.Set.fromElements((_228_v84)[_module.__default.safeIndex(_117_v0, new BigNumber((_228_v84).length))])).Elements) {
            let _230_v85 = _compr_10;
            if ((_dafny.Set.fromElements((_228_v84)[_module.__default.safeIndex(_117_v0, new BigNumber((_228_v84).length))])).contains(_230_v85)) {
              _coll10.add(_module.__default.safeModuloInt(_230_v85, new BigNumber(-93)));
            }
          }
          return _coll10;
        }()));
        _221_v77 = _rhs42;
        _222_v78 = _rhs43;
        let _231_v87;
        _231_v87 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_119_v2, _dafny.Seq.UnicodeFromString("jlvlgk")),_121_v4);
        _231_v87 = (_231_v87).update(_module.__default.fm5(_dafny.Map.Empty.slice().updateUnsafe(_182_v48,_117_v0), _182_v48, _117_v0, _117_v0, _125_globalState), _121_v4);
        _module.__default.m0(_182_v48, _125_globalState);
        let _232_v88;
        _232_v88 = _dafny.Map.Empty.slice().updateUnsafe(_121_v4,_117_v0);
        _232_v88 = _232_v88;
        let _index45 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_121_v4).length));
        (_121_v4)[_index45] = (_117_v0).multipliedBy(_117_v0);
        let _233_v89;
        let _init8 = ((_234_v84) => function (_235_i9) {
          return _234_v84;
        })(_228_v84);
        let _nw39 = Array((_dafny.ONE).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw39.length); _i0_8++) {
          _nw39[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _233_v89 = _nw39;
        let _236_v90;
        _236_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6(_module.__default.fm6(_182_v48, new BigNumber((_dafny.Set.fromElements(_121_v4, _121_v4, _121_v4, (_module.D7.create_DC17(_121_v4)).dtor_cf19, _121_v4)).length), _125_globalState), _117_v0, _125_globalState),_dafny.Seq.of((_dafny.ZERO).minus(_117_v0), _117_v0));
        let _237_v91;
        _237_v91 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,!(_182_v48));
        let _238_v92;
        _238_v92 = _module.D2.create_DC6(_237_v91, _182_v48);
        let _index46 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_233_v89).length));
        (_233_v89)[_index46] = (((_236_v90).contains(!((_238_v92).dtor_cf7))) ? ((_236_v90).get(!((_238_v92).dtor_cf7))) : (_dafny.Seq.of(_117_v0)));
        let _index47 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_121_v4).length));
        let _index48 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_233_v89).length));
        let _rhs44 = (_117_v0).multipliedBy(_117_v0);
        let _rhs45 = ((_117_v0).multipliedBy(_117_v0)).plus(_117_v0);
        let _rhs46 = _228_v84;
        let _lhs48 = _121_v4;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_121_v4).length));
        let _lhs50 = _125_globalState;
        let _lhs51 = _233_v89;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_233_v89).length));
        _lhs48[_lhs49] = _rhs44;
        _lhs50.f3 = _rhs45;
        _lhs51[_lhs52] = _rhs46;
      }
      let _239_i10;
      _239_i10 = _dafny.ZERO;
      L2: {
        while (false) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_239_i10)) {
              break L2;
            }
            _239_i10 = (_239_i10).plus(_dafny.ONE);
            let _240_v93;
            _240_v93 = _module.D0.create_DC1(_117_v0, _182_v48, _117_v0);
            let _241_v94;
            _241_v94 = _dafny.Map.Empty.slice().updateUnsafe(_240_v93,(_dafny.ZERO).minus(_117_v0));
            let _242_v95;
            _242_v95 = _241_v94;
            let _source1 = _242_v95;
            let _243___mcc_h0 = _source1;
            let _244_cf14 = _243___mcc_h0;
            let _245_v96;
            _245_v96 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,_182_v48);
            _182_v48 = ((new BigNumber((_245_v96).length)).isLessThanOrEqualTo(_117_v0)) === (_182_v48);
            _module.__default.m0(_182_v48, _125_globalState);
            _117_v0 = ((!(_182_v48)) ? (_117_v0) : (_117_v0));
            let _246_v97;
            _246_v97 = _module.D2.create_DC6(_245_v96, _182_v48);
            let _index49 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_122_v5).length));
            (_122_v5)[_index49] = ((_246_v97).dtor_cf7) === (_182_v48);
            let _index50 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_122_v5).length));
            (_122_v5)[_index50] = _182_v48;
            (_125_globalState).f2 = !(_182_v48);
            let _rhs47 = new BigNumber((_119_v2).length);
            let _rhs48 = _117_v0;
            let _lhs53 = _125_globalState;
            _117_v0 = _rhs47;
            _lhs53.f4 = _rhs48;
            if (_182_v48) {
              let _247_v98;
              _247_v98 = _module.D0.create_DC0(_117_v0);
              (_125_globalState).f2 = (_117_v0).isLessThan((_247_v98).dtor_cf0);
              _119_v2 = _dafny.Seq.UnicodeFromString("uuglv");
              let _248_v99;
              let _nw40 = new _module.C0();
              _nw40.__ctor(_241_v94);
              _248_v99 = _nw40;
              (_125_globalState).f4 = (new BigNumber((_module.__default.fm1(_125_globalState)).length)).plus(new BigNumber(759));
              let _249_v100;
              _249_v100 = _dafny.Seq.of(_117_v0, new BigNumber(609), _117_v0);
              let _index51 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_121_v4).length));
              (_121_v4)[_index51] = ((_182_v48) ? (_117_v0) : (new BigNumber((_249_v100).length)));
              let _index52 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_121_v4).length));
              (_121_v4)[_index52] = _117_v0;
            } else {
              let _250_v101;
              _250_v101 = _dafny.Map.Empty.slice().updateUnsafe((_117_v0).multipliedBy(_117_v0),(_117_v0).minus(_117_v0));
              _250_v101 = _250_v101;
              (_125_globalState).f3 = _117_v0;
              _module.__default.m0(_dafny.Seq.IsPrefixOf(_183_v49, _dafny.Seq.of(_182_v48)), _125_globalState);
              (_125_globalState).f3 = _117_v0;
              let _pat_let_tv1 = _117_v0;
              let _251_v102;
              let _nw41 = new _module.C0();
              _nw41.__ctor(_dafny.Map.Empty.slice().updateUnsafe(function (_pat_let4_0) {
                return function (_252_dt__update__tmp_h0) {
                  return function (_pat_let5_0) {
                    return function (_253_dt__update_hcf3_h0) {
                      return _module.D0.create_DC1((_252_dt__update__tmp_h0).dtor_cf1, (_252_dt__update__tmp_h0).dtor_cf2, _253_dt__update_hcf3_h0);
                    }(_pat_let5_0);
                  }(_pat_let_tv1);
                }(_pat_let4_0);
              }(_240_v93),_117_v0));
              _251_v102 = _nw41;
            }
          }
        }
      }
      let _source2 = _module.D0.create_DC0(_117_v0);
      if (_source2.is_DC1) {
        let _254___mcc_h1 = (_source2).cf1;
        let _255___mcc_h2 = (_source2).cf2;
        let _256___mcc_h3 = (_source2).cf3;
        let _257_cf3 = _256___mcc_h3;
        let _258_cf2 = _255___mcc_h2;
        let _259_cf1 = _254___mcc_h1;
        let _260_v103;
        _260_v103 = new _dafny.CodePoint('o'.codePointAt(0));
        _260_v103 = new _dafny.CodePoint('u'.codePointAt(0));
        let _261_v104;
        let _nw42 = Array((new BigNumber(27)).toNumber()).fill([]);
        _261_v104 = _nw42;
        let _index53 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_261_v104).length));
        (_261_v104)[_index53] = _121_v4;
        let _index54 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_261_v104).length));
        let _nw43 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        (_261_v104)[_index54] = _nw43;
        let _index55 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_122_v5).length));
        (_122_v5)[_index55] = _182_v48;
        let _index56 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_122_v5).length));
        (_122_v5)[_index56] = _182_v48;
        _259_cf1 = _module.__default.fm0(_module.__default.safeModuloInt(_257_cf3, _259_cf1), _module.__default.fm0(_257_cf3, _257_cf3, _125_globalState), _125_globalState);
      } else if (_source2.is_DC2) {
        (_125_globalState).f2 = _182_v48;
        let _index57 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_122_v5).length));
        (_122_v5)[_index57] = false;
        let _262_v105;
        _262_v105 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,_182_v48);
        let _263_v106;
        _263_v106 = _dafny.MultiSet.fromElements(new BigNumber((_262_v105).length));
        let _264_v107;
        _264_v107 = _module.D2.create_DC5(_263_v106);
        let _265_v108;
        _265_v108 = _dafny.Map.Empty.slice().updateUnsafe(((_182_v48) ? (_module.D2.create_DC5(_263_v106)) : (_264_v107)),(_117_v0).isEqualTo(new BigNumber(953)));
        let _index58 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_122_v5).length));
        let _rhs49 = !(!((new BigNumber(956)).isLessThan((_dafny.ZERO).minus(_117_v0))));
        let _rhs50 = true;
        let _rhs51 = _117_v0;
        let _rhs52 = (((_265_v108).contains(_264_v107)) ? ((_265_v108).get(_264_v107)) : (_182_v48));
        let _lhs54 = _125_globalState;
        let _lhs55 = _122_v5;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_122_v5).length));
        let _lhs57 = _125_globalState;
        let _lhs58 = _125_globalState;
        _lhs54.f2 = _rhs49;
        _lhs55[_lhs56] = _rhs50;
        _lhs57.f4 = _rhs51;
        _lhs58.f2 = _rhs52;
        _117_v0 = _117_v0;
        let _266_v109;
        _266_v109 = _module.D2.create_DC6(_262_v105, (_122_v5)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_122_v5).length))]);
        let _267_v110;
        _267_v110 = new _dafny.CodePoint('x'.codePointAt(0));
        let _268_v111;
        _268_v111 = _dafny.Map.Empty.slice().updateUnsafe(_266_v109,_267_v110);
        let _269_v112;
        _269_v112 = _dafny.Set.fromElements((((_268_v111).contains(_266_v109)) ? ((_268_v111).get(_266_v109)) : (_267_v110)));
        let _270_v113;
        _270_v113 = _dafny.Seq.of(_module.__default.fm0(_117_v0, new BigNumber((_dafny.Seq.UnicodeFromString("jfgpdshm")).length), _125_globalState), _117_v0);
        let _271_v114;
        _271_v114 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_270_v113)).cardinality()));
        let _rhs53 = _module.__default.fm7(_267_v110, new BigNumber((_dafny.Seq.Concat(_module.__default.fm8(_117_v0, _125_globalState), _271_v114)).length), _117_v0, _182_v48, _125_globalState);
        let _rhs54 = (_122_v5)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_122_v5).length))];
        let _rhs55 = _267_v110;
        let _lhs59 = _125_globalState;
        _269_v112 = _rhs53;
        _lhs59.f2 = _rhs54;
        _267_v110 = _rhs55;
      } else if (_source2.is_DC3) {
        (_125_globalState).f2 = _182_v48;
        let _272_v115;
        _272_v115 = _module.D0.create_DC0(_117_v0);
        _272_v115 = _272_v115;
        (_125_globalState).f2 = _182_v48;
        if (_182_v48) {
          let _273_v116;
          _273_v116 = _module.D0.create_DC1(new BigNumber((_119_v2).length), !(_182_v48), _117_v0);
          let _274_v117;
          _274_v117 = _dafny.Map.Empty.slice().updateUnsafe(_273_v116,new BigNumber((_183_v49).length));
          let _275_v118;
          let _nw44 = new _module.C0();
          _nw44.__ctor(_274_v117);
          _275_v118 = _nw44;
          _119_v2 = _dafny.Seq.Concat(_119_v2, _dafny.Seq.Concat(_119_v2, _119_v2));
          let _276_v119;
          let _nw45 = new _module.C0();
          _nw45.__ctor((_275_v118).f15);
          _276_v119 = _nw45;
          _276_v119 = _275_v118;
          let _index59 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_122_v5).length));
          (_122_v5)[_index59] = (_117_v0).isLessThan(_117_v0);
          let _index60 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_122_v5).length));
          (_122_v5)[_index60] = _module.__default.fm6(false, new BigNumber(603), _125_globalState);
        } else {
          _119_v2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_119_v2, _119_v2), _119_v2);
          let _277_v120;
          _277_v120 = _module.D0.create_DC1((_dafny.ZERO).minus((_dafny.ZERO).minus(_117_v0)), _182_v48, _117_v0);
          let _278_v121;
          _278_v121 = _dafny.Map.Empty.slice().updateUnsafe(_277_v120,_117_v0);
          let _279_v122;
          let _nw46 = new _module.C0();
          _nw46.__ctor(_278_v121);
          _279_v122 = _nw46;
          let _280_v124;
          _280_v124 = _dafny.Seq.of(_277_v120, _277_v120);
          let _281_v125;
          let _nw47 = new _module.C0();
          _nw47.__ctor(function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_280_v124).Elements) {
              let _282_v123 = _compr_11;
              if (_dafny.Seq.contains(_280_v124, _282_v123)) {
                _coll11.push([_282_v123,_117_v0]);
              }
            }
            return _coll11;
          }());
          _281_v125 = _nw47;
          let _index61 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_121_v4).length));
          (_121_v4)[_index61] = _117_v0;
          let _283_v126;
          _283_v126 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,_182_v48);
          let _284_v127;
          _284_v127 = _module.D2.create_DC6(_283_v126, (_182_v48) && (_182_v48));
          let _index62 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_121_v4).length));
          let _rhs56 = _281_v125;
          let _rhs57 = _module.__default.safeModuloInt(_117_v0, _117_v0);
          let _rhs58 = _182_v48;
          let _rhs59 = _117_v0;
          let _rhs60 = _284_v127;
          let _lhs60 = _121_v4;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_121_v4).length));
          let _lhs62 = _125_globalState;
          let _lhs63 = _125_globalState;
          _281_v125 = _rhs56;
          _lhs60[_lhs61] = _rhs57;
          _lhs62.f2 = _rhs58;
          _lhs63.f3 = _rhs59;
          _284_v127 = _rhs60;
          let _285_v128;
          _285_v128 = new _dafny.CodePoint('c'.codePointAt(0));
          let _286_v129;
          _286_v129 = _module.D8.create_DC20(_285_v128);
          _285_v128 = (_286_v129).dtor_cf22;
          let _287_v130;
          let _nw48 = Array((new BigNumber(6)).toNumber());
          _287_v130 = _nw48;
          let _index63 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_287_v130).length));
          (_287_v130)[_index63] = _281_v125;
          let _288_v131;
          _288_v131 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,_281_v125);
          let _index64 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_287_v130).length));
          (_287_v130)[_index64] = (((_288_v131).contains(_182_v48)) ? ((_288_v131).get(_182_v48)) : (_281_v125));
        }
      } else {
        let _289___mcc_h4 = (_source2).cf0;
        let _290_cf0 = _289___mcc_h4;
        let _291_v132;
        _291_v132 = new _dafny.CodePoint('i'.codePointAt(0));
        let _292_v133;
        _292_v133 = _dafny.MultiSet.fromElements(_291_v132, _291_v132, _291_v132, _291_v132);
        (_125_globalState).f2 = _module.__default.fm6(_182_v48, new BigNumber((_292_v133).cardinality()), _125_globalState);
        let _index65 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length));
        (_122_v5)[_index65] = (_183_v49)[_module.__default.safeIndex(_117_v0, new BigNumber((_183_v49).length))];
        let _293_v134;
        _293_v134 = _dafny.Seq.of(_117_v0);
        let _index66 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length));
        (_122_v5)[_index66] = (_module.__default.safeModuloInt(new BigNumber(694), new BigNumber((_dafny.Seq.of(_182_v48)).length))).isLessThanOrEqualTo((_293_v134)[_module.__default.safeIndex(_290_cf0, new BigNumber((_293_v134).length))]);
        let _294_v135;
        _294_v135 = _dafny.Set.fromElements(_182_v48);
        let _295_v136;
        _295_v136 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(new BigNumber((_294_v135).length), _182_v48, _117_v0),(_dafny.ZERO).minus(_290_cf0));
        let _296_v137;
        let _nw49 = new _module.C0();
        _nw49.__ctor(_295_v136);
        _296_v137 = _nw49;
        let _297_v138;
        _297_v138 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,_296_v137);
        if ((_118_v1).IsSubsetOf(_module.__default.fm9(new BigNumber((_297_v138).length), _module.__default.fm0(_290_cf0, _117_v0, _125_globalState), _125_globalState))) {
          (_125_globalState).f4 = (_module.__default.safeModuloInt(_290_cf0, _290_cf0)).minus(new BigNumber(-167));
          let _298_v139;
          _298_v139 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),(_122_v5)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length))]);
          _module.__default.m0((((_298_v139).contains(_291_v132)) ? ((_298_v139).get(_291_v132)) : (true)), _125_globalState);
          _296_v137 = _296_v137;
          let _299_v141;
          _299_v141 = _dafny.MultiSet.fromElements(_290_cf0);
          let _300_v142;
          _300_v142 = _module.D0.create_DC1((((_299_v141).contains(_290_cf0)) ? ((_299_v141).get(_290_cf0)) : (_117_v0)), (_122_v5)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length))], _290_cf0);
          let _301_v143;
          _301_v143 = _dafny.Map.Empty.slice().updateUnsafe(_300_v142,(_122_v5)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length))]);
          let _302_v144;
          let _nw50 = new _module.C0();
          _nw50.__ctor((function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_301_v143).Keys.Elements) {
              let _303_v140 = _compr_12;
              if ((_301_v143).contains(_303_v140)) {
                _coll12.push([_303_v140,_117_v0]);
              }
            }
            return _coll12;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_300_v142,_117_v0)));
          _302_v144 = _nw50;
          (_125_globalState).f2 = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), function (_304_i11) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }), _291_v132);
        } else {
          _291_v132 = (_module.D8.create_DC20(_291_v132)).dtor_cf22;
          _module.__default.m0((_122_v5)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length))], _125_globalState);
          (_125_globalState).f3 = _117_v0;
          let _305_v145;
          _305_v145 = _dafny.Map.Empty.slice().updateUnsafe((_122_v5)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_122_v5).length))],new BigNumber(189));
          let _306_v146;
          _306_v146 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,(((_305_v145).contains(true)) ? ((_305_v145).get(true)) : (_290_cf0)));
          let _307_v147;
          _307_v147 = _dafny.Map.Empty.slice().updateUnsafe((((_306_v146).contains(_290_cf0)) ? ((_306_v146).get(_290_cf0)) : (new BigNumber((_dafny.Seq.UnicodeFromString("yeoger")).length))),_290_cf0);
          let _308_v148;
          _308_v148 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(796),new BigNumber(480)));
          _307_v147 = (((_308_v148)[_module.__default.safeIndex(_290_cf0, new BigNumber((_308_v148).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_119_v2).length),new BigNumber(503)))).Merge(_307_v147);
          _119_v2 = _dafny.Seq.UnicodeFromString("vfpkjomp");
        }
        let _309_v149;
        _309_v149 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(_125_globalState),_296_v137);
        _290_cf0 = new BigNumber((_309_v149).length);
      }
      let _index67 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length));
      (_121_v4)[_index67] = _117_v0;
      let _310_v150;
      _310_v150 = _dafny.Set.fromElements(_122_v5, _122_v5, _122_v5, _122_v5);
      let _index68 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length));
      (_121_v4)[_index68] = _module.__default.safeModuloInt(new BigNumber((_310_v150).length), new BigNumber(328));
      let _311_v151;
      _311_v151 = _dafny.Map.Empty.slice().updateUnsafe(_117_v0,_module.__default.fm6(_182_v48, (_121_v4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length))], _125_globalState));
      let _index69 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
      (_122_v5)[_index69] = _182_v48;
      let _index70 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length));
      let _index71 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
      let _rhs61 = ((_121_v4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length))]).plus((_dafny.ZERO).minus(_117_v0));
      let _rhs62 = _182_v48;
      let _rhs63 = _311_v151;
      let _rhs64 = _121_v4;
      let _rhs65 = !(new BigNumber(956)).isEqualTo(_117_v0);
      let _lhs64 = _121_v4;
      let _lhs65 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length));
      let _lhs66 = _122_v5;
      let _lhs67 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
      _lhs64[_lhs65] = _rhs61;
      _182_v48 = _rhs62;
      _311_v151 = _rhs63;
      _121_v4 = _rhs64;
      _lhs66[_lhs67] = _rhs65;
      let _312_v152;
      _312_v152 = _dafny.MultiSet.fromElements(_118_v1);
      let _313_v153;
      _313_v153 = _dafny.Seq.of(_312_v152);
      let _314_i12;
      _314_i12 = _dafny.ZERO;
      L3: {
        while ((((_313_v153)[_module.__default.safeIndex(_117_v0, new BigNumber((_313_v153).length))]).Union(_312_v152)).IsProperSubsetOf((_312_v152).Union(_312_v152))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_314_i12)) {
              break L3;
            }
            _314_i12 = (_314_i12).plus(_dafny.ONE);
            let _315_v154;
            _315_v154 = _module.D0.create_DC1((_121_v4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length))], (_122_v5)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length))], (_121_v4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length))]);
            let _316_v155;
            _316_v155 = _dafny.Map.Empty.slice().updateUnsafe(_315_v154,_117_v0);
            let _317_v156;
            let _nw51 = new _module.C0();
            _nw51.__ctor(_316_v155);
            _317_v156 = _nw51;
            let _318_v157;
            _318_v157 = _dafny.Map.Empty.slice().updateUnsafe(_317_v156,_119_v2);
            let _index72 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
            let _index73 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
            let _rhs66 = (_122_v5)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length))];
            let _rhs67 = (((_318_v157).contains(_317_v156)) ? ((_318_v157).get(_317_v156)) : (_119_v2));
            let _rhs68 = (true) && (!(_182_v48));
            let _lhs68 = _122_v5;
            let _lhs69 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
            let _lhs70 = _122_v5;
            let _lhs71 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length));
            _lhs68[_lhs69] = _rhs66;
            _119_v2 = _rhs67;
            _lhs70[_lhs71] = _rhs68;
            let _319_v158;
            _319_v158 = new _dafny.CodePoint('d'.codePointAt(0));
            let _320_v159;
            _320_v159 = _module.D8.create_DC20(_319_v158);
            _182_v48 = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), ((_321_v159) => function (_322_i13) {
              return _321_v159;
            })(_320_v159)), _320_v159);
            _317_v156 = _317_v156;
            let _323_v160;
            _323_v160 = _dafny.Map.Empty.slice().updateUnsafe(!((_122_v5)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length))]),_317_v156);
            _323_v160 = _323_v160;
          }
        }
      }
      let _rhs69 = _182_v48;
      let _rhs70 = _dafny.Seq.of(_182_v48);
      let _lhs72 = _125_globalState;
      _lhs72.f2 = _rhs69;
      _183_v49 = _rhs70;
      (_125_globalState).f2 = _182_v48;
      let _324_v161;
      let _init9 = ((_325_v48, _326_v0) => function (_327_i15) {
        return _dafny.Map.Empty.slice().updateUnsafe(_325_v48,_326_v0);
      })(_182_v48, _117_v0);
      let _nw52 = Array((new BigNumber(4)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw52.length); _i0_9++) {
        _nw52[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _324_v161 = _nw52;
      let _ingredients0 = [];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_324_v161).length))) {
        let _328_i14 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_328_i14)) && ((_328_i14).isLessThan(new BigNumber((_324_v161).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_324_v161, (_328_i14).toNumber(), (_dafny.Map.Empty.slice().updateUnsafe(false,_117_v0)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_122_v5)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length))],(_121_v4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length))]))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      _module.__default.m0(!(_182_v48) || ((_122_v5)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length))]), _125_globalState);
      let _329_v162;
      _329_v162 = (_119_v2);
      let _pat_let_tv2 = _117_v0;
      let _pat_let_tv3 = _121_v4;
      let _pat_let_tv4 = _121_v4;
      let _index74 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length));
      (_121_v4)[_index74] = function (_source3) {
        let _330___mcc_h5 = _source3;
        let _331_cf4 = _330___mcc_h5;
        return (_dafny.ZERO).minus((_pat_let_tv2).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_pat_let_tv4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_pat_let_tv3).length))]))).cardinality())));
      }(_329_v162);
      _117_v0 = _117_v0;
      let _332_v163;
      _332_v163 = new _dafny.CodePoint('e'.codePointAt(0));
      let _333_v164;
      _333_v164 = _dafny.Map.Empty.slice().updateUnsafe(_332_v163,_182_v48);
      let _334_v165;
      _334_v165 = _dafny.Map.Empty.slice().updateUnsafe(_182_v48,new BigNumber((_333_v164).length));
      let _335_v166;
      _335_v166 = _module.D0.create_DC1(_117_v0, _182_v48, (_dafny.ZERO).minus(new BigNumber((_334_v165).length)));
      let _336_v167;
      _336_v167 = _dafny.Map.Empty.slice().updateUnsafe(_335_v166,_117_v0);
      let _337_v168;
      let _nw53 = new _module.C0();
      _nw53.__ctor(_336_v167);
      _337_v168 = _nw53;
      let _338_v169;
      _338_v169 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_337_v168,_337_v168));
      let _339_v170;
      _339_v170 = _module.D8.create_DC21(_117_v0, _337_v168);
      let _340_v171;
      _340_v171 = _dafny.Map.Empty.slice().updateUnsafe((_339_v170).dtor_cf24,_337_v168);
      _338_v169 = (_338_v169).update(_340_v171, _module.__default.abs(_117_v0));
      let _341_i16;
      _341_i16 = _dafny.ZERO;
      L4: {
        while (((_122_v5)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_122_v5).length))]) || (_182_v48)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_341_i16)) {
              break L4;
            }
            _341_i16 = (_341_i16).plus(_dafny.ONE);
            _334_v165 = (_334_v165).update(_182_v48, _117_v0);
            _117_v0 = (new BigNumber(-967)).plus(new BigNumber(((_118_v1).Union(_118_v1)).length));
            _119_v2 = _119_v2;
            _334_v165 = (_334_v165).update(_182_v48, (_121_v4)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_121_v4).length))]);
          }
        }
      }
      let _342_v172;
      let _nw54 = Array((new BigNumber(3)).toNumber()).fill([]);
      _342_v172 = _nw54;
      let _343_v173;
      _343_v173 = _dafny.Map.Empty.slice().updateUnsafe(_342_v172,_182_v48);
      _343_v173 = (_343_v173).update(_342_v172, _182_v48);
      process.stdout.write(_dafny.toString(_117_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v1).equals(_dafny.Set.fromElements(new BigNumber(-877)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_119_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("yxsvjmbqp")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_125_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_125_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_125_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f6)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f8)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_globalState).f9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_125_globalState).f12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_126_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_v48));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_183_v49, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_239_i10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_310_v150).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_311_v151).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_312_v152).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-877))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_313_v153, _dafny.Seq.of(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-877)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_314_i12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_324_v161)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)).updateUnsafe(true,new BigNumber(-2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_324_v161)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)).updateUnsafe(true,new BigNumber(-2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_324_v161)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)).updateUnsafe(true,new BigNumber(-2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_324_v161)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)).updateUnsafe(true,new BigNumber(-2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_329_v162)).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_332_v163));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_333_v164).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v165).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v166).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v166).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v166).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_336_v167).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(new BigNumber(3), false, new BigNumber(-1)),new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_337_v168).f15).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(new BigNumber(3), false, new BigNumber(-1)),new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_338_v169).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_339_v170).dtor_cf23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_339_v170).dtor_cf24).f15).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(new BigNumber(3), false, new BigNumber(-1)),new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_340_v171).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_341_i16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_343_v173).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2() {
      let $dt = new D0(1);
      return $dt;
    }
    static create_DC3() {
      let $dt = new D0(2);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2";
      } else if (this.$tag === 2) {
        return "D0.DC3";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC4(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + this.cf4.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC7(cf8, cf9) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.Map.Empty, false);
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
    static create_DC8(cf10) {
      let $dt = new D3(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf10) + ")";
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
    static create_DC10() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC11(cf12) {
      let $dt = new D4(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC9(cf11) {
      let $dt = new D4(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC12(cf13) {
      let $dt = new D4(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf13) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
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
    static create_DC13(cf14) {
      let $dt = new D5(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf16, cf17) {
      let $dt = new D6(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC14(cf15) {
      let $dt = new D6(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC16(cf18) {
      let $dt = new D6(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC18(cf20) {
      let $dt = new D7(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC17(cf19) {
      let $dt = new D7(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC19(cf21) {
      let $dt = new D7(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18(_dafny.ZERO);
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
    static create_DC21(cf23, cf24) {
      let $dt = new D8(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC20(cf22) {
      let $dt = new D8(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(_dafny.ZERO, null);
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
    static create_DC23(cf26) {
      let $dt = new D9(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC24(cf27, cf28, cf29, cf30) {
      let $dt = new D9(1);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC25() {
      let $dt = new D9(2);
      return $dt;
    }
    static create_DC22(cf25) {
      let $dt = new D9(3);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC22() { return this.$tag === 3; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC25";
      } else if (this.$tag === 3) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(_dafny.ZERO);
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
    static create_DC27(cf32, cf33) {
      let $dt = new D10(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC28(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D10(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC26(cf31) {
      let $dt = new D10(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf32 === other.cf32 && this.cf33 === other.cf33;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC27(false, false);
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
    static create_DC30(cf40) {
      let $dt = new D11(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC29(cf39) {
      let $dt = new D11(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = false;
      this.f3 = _dafny.ZERO;
      this.f4 = _dafny.ZERO;
      this._f0 = false;
      this._f1 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = [];
      this._f7 = false;
      this._f8 = [];
      this._f9 = [];
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = _dafny.Seq.UnicodeFromString("");
      this._f13 = _dafny.ZERO;
      this._f14 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f15 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f15) {
      let _this = this;
      (_this)._f15 = f15;
      return;
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
