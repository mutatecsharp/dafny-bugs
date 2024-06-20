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
    static fm0(p0, globalState) {
      return _module.D0.create_DC1((new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).length), new BigNumber(216))).length)).multipliedBy(new BigNumber(906)), ((!(true)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), function (_0_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})) : (_dafny.Seq.UnicodeFromString("jyfwejrpi"))), true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(168))).length))).length)), true);
    };
    static fm1(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat((_module.D14.create_DC32(true, new BigNumber((_dafny.Set.fromElements(new BigNumber(-907))).length), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0))), true)).dtor_cf55, _dafny.Seq.UnicodeFromString("racyt")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("irsogxd"), _dafny.Seq.UnicodeFromString("maxquypl")));
    };
    static fm2(p0, p1, globalState) {
      return (new BigNumber((((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(932),new BigNumber(220))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wgfhyrgb")).length),new BigNumber((_dafny.Seq.of(false)).length))))).length)).minus(new BigNumber(417));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      let _source0 = ((!(true)) ? (_dafny.MultiSet.fromElements(new BigNumber(-389))) : (_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-550))))));
      let _1___mcc_h0 = _source0;
      let _2_cf18 = _1___mcc_h0;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_2_cf18).cardinality())), _dafny.Seq.of(new BigNumber(67)));
    };
    static fm4(p0, p1, p2, globalState) {
      if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(879), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),false)).length)))).contains(new BigNumber((_dafny.Set.fromElements(true)).length))) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }
    };
    static fm5(p0, p1, p2, globalState) {
      return false;
    };
    static fm10(p0, p1, globalState) {
      return _module.D1.create_DC3(((true) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (new _dafny.CodePoint('p'.codePointAt(0)))));
    };
    static fm12(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(-113))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(!(true))).length)))).Intersect((_module.D6.create_DC13(new BigNumber(7), new BigNumber((_dafny.Seq.UnicodeFromString("hwvbfw")).length), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-416)), new BigNumber(686)))).dtor_cf22);
    };
    static fm13(p0, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(!(true), true))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(true), (_dafny.Seq.of(true))))).Difference((_dafny.Set.fromElements(_dafny.Seq.of(false, false))).Difference(_dafny.Set.fromElements(_dafny.Seq.of(true, false))));
    };
    static fm21(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length))).Difference((_dafny.MultiSet.fromElements(new BigNumber(-786), new BigNumber(57))));
    };
    static fm22(p0, p1, p2, globalState) {
      return _module.D1.create_DC4(!(false) || (false), !_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(true), _dafny.Set.fromElements(false), _dafny.Set.fromElements(false)), _dafny.Set.fromElements(true, false)));
    };
    static fm23(globalState) {
      return _dafny.Seq.of(true, !(true), true, !(!(new BigNumber(-314)).isEqualTo(new BigNumber((_dafny.Set.fromElements(new BigNumber(653))).length))), !(_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(597)), function (_3_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)),new BigNumber(125))).length);
      }), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true)))).cardinality())))));
    };
    static fm24(p0, globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(595), new BigNumber(298))) {
          let _4_v0 = _compr_0;
          if (((new BigNumber(595)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(298)))) {
            _coll0.add(_module.__default.safeModuloInt(_4_v0, new BigNumber(642)));
          }
        }
        return _coll0;
      }()).Intersect((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("pejhxpg")).length), new BigNumber(161))).Elements) {
          let _5_v1 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("pejhxpg")).length), new BigNumber(161)), _5_v1)) {
            _coll1.add((_5_v1).multipliedBy(new BigNumber(-835)));
          }
        }
        return _coll1;
      }()).Union(_dafny.Set.fromElements(new BigNumber(24), new BigNumber(-476), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(262),new _dafny.CodePoint('v'.codePointAt(0)))).length))));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(true),new BigNumber(((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true, true, !(true), false)).cardinality())))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber(782), new BigNumber(-196), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ef")).length)), new BigNumber(464)))).cardinality()));
    };
    static fm26(p0, p1, p2, globalState) {
      if (!(false) || (true)) {
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('h'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('u'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('o'.codePointAt(0))));
      } else {
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('s'.codePointAt(0)))))).Intersect(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('e'.codePointAt(0)))));
      }
    };
    static fm27(p0, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(false))))).Difference(_dafny.MultiSet.fromElements(true, !(false)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)));
    };
    static fm28(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ufhvh"), new _dafny.CodePoint('e'.codePointAt(0))));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements(true, false))).Intersect(_dafny.Set.fromElements(true, true));
    };
    static fm30(globalState) {
      return ((false) ? (_dafny.MultiSet.fromElements(new BigNumber(-988), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(632))).length)))) : (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true, !(false))).length))).length))).length))));
    };
    static fm31(globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)));
    };
    static fm32(globalState) {
      let _source1 = _module.D0.create_DC2();
      if (_source1.is_DC1) {
        let _6___mcc_h0 = (_source1).cf1;
        let _7___mcc_h1 = (_source1).cf2;
        let _8___mcc_h2 = (_source1).cf3;
        let _9___mcc_h3 = (_source1).cf4;
        let _10___mcc_h4 = (_source1).cf5;
        let _11_cf5 = _10___mcc_h4;
        let _12_cf4 = _9___mcc_h3;
        let _13_cf3 = _8___mcc_h2;
        let _14_cf2 = _7___mcc_h1;
        let _15_cf1 = _6___mcc_h0;
        return _module.D4.create_DC10(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)));
      } else if (_source1.is_DC2) {
        if (true) {
          return _module.D4.create_DC10(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)));
        } else {
          return _module.D4.create_DC10(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)));
        }
      } else {
        let _16___mcc_h5 = (_source1).cf0;
        let _17_cf0 = _16___mcc_h5;
        return _module.D4.create_DC10(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)));
      }
    };
    static fm33(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true),!(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true, false),true)));
    };
    static fm34(p0, p1, globalState) {
      if ((!(!(true))) === (false)) {
        return _module.D11.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(true,true));
      } else {
        return _module.D11.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(false)));
      }
    };
    static fm35(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)));
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,!(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length))), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(687)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(566)))).length)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(459)))).contains(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(138))));
    };
    static fm37(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(true)).Union(_dafny.MultiSet.fromElements(!(false), true)),(_module.D4.create_DC10(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).dtor_cf16);
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC6(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(297),_module.D14.create_DC32(true, new BigNumber((_dafny.Seq.UnicodeFromString("bjwmcxnlo")).length), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))), false))).length)));
    };
    static fm39(p0, p1, globalState) {
      return _module.D4.create_DC9(_dafny.Set.fromElements(true));
    };
    static fm40(p0, globalState) {
      return _module.D14.create_DC32((_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).IsProperSubsetOf((_module.D22.create_DC50(_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0))))).dtor_cf80), (new BigNumber(290)).plus(new BigNumber(-436)), (_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)), _dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))), true);
    };
    static fm41(p0, p1, p2, globalState) {
      return _module.D6.create_DC14(new BigNumber(((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(715),new BigNumber(609))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(330),!(true))).length))).length))))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(260), new BigNumber(-500)))).cardinality()), new BigNumber(586), (_dafny.MultiSet.fromElements(!(false))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))));
    };
    static fm42(p0, p1, p2, globalState) {
      return function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(848)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)))).Elements) {
          let _18_v0 = _compr_2;
          if (((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(848)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)))).contains(_18_v0)) {
            _coll2.push([_module.__default.safeModuloInt(_18_v0, new BigNumber(-263)),true]);
          }
        }
        return _coll2;
      }();
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-464), new BigNumber(557))) {
          let _19_v0 = _compr_3;
          if (((new BigNumber(-464)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(557)))) {
            _coll3.push([_module.__default.safeModuloInt(_19_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()), new BigNumber((_dafny.Set.fromElements(false, true)).length))).length))),new BigNumber(200)]);
          }
        }
        return _coll3;
      }();
    };
    static fm44(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(282)), function (_20_i0) {
        return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), function (_21_i1) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(328))).length);
        }));
      }), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(28)), function (_22_i2) {
        return new BigNumber((_dafny.Set.fromElements(true)).length);
      })), _dafny.MultiSet.fromElements(new BigNumber(647), new BigNumber(-625)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),true)).length))), _dafny.MultiSet.fromElements(new BigNumber(645), new BigNumber(732), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_23_i3) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length)), new BigNumber((function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-771)), function (_24_i4) {
          return new BigNumber((_dafny.Seq.of(false)).length);
        })).Elements) {
          let _25_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-771)), function (_24_i4) {
            return new BigNumber((_dafny.Seq.of(false)).length);
          }), _25_v0)) {
            _coll4.add((_25_v0).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), function (_26_i5) {
              return new BigNumber((_dafny.Seq.of(new BigNumber(-688))).length);
            })).length))));
          }
        }
        return _coll4;
      }()).length)), _dafny.MultiSet.fromElements(new BigNumber(687), new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(-742))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(586))).length)))).length)))));
    };
    static m0(p0, globalState) {
      let r0 = _dafny.ZERO;
      let _27_v0;
      _27_v0 = new BigNumber(629);
      let _28_i0;
      _28_i0 = _dafny.ZERO;
      L0: {
        while ((_27_v0).isLessThan(_27_v0)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_28_i0)) {
              break L0;
            }
            _28_i0 = (_28_i0).plus(_dafny.ONE);
            if (p0) {
              let _29_v1;
              let _nw0 = new _module.C1();
              _nw0.__ctor(p0);
              _29_v1 = _nw0;
              (globalState).f2 = _27_v0;
              _27_v0 = (_dafny.ZERO).minus(_27_v0);
              let _30_v2;
              _30_v2 = _module.D0.create_DC0((_29_v1).f32);
              let _31_v3;
              let _nw1 = new _module.C2();
              _nw1.__ctor(_27_v0, p0, (_dafny.ZERO).minus(new BigNumber(-851)));
              _31_v3 = _nw1;
              let _32_v4;
              _32_v4 = _dafny.Map.Empty.slice().updateUnsafe(_30_v2,new BigNumber((_dafny.Seq.of(_31_v3)).length));
              let _33_v5;
              _33_v5 = _module.D10.create_DC23();
              let _34_v6;
              _34_v6 = _dafny.Seq.UnicodeFromString("ehxqgq");
              let _35_v7;
              _35_v7 = _dafny.MultiSet.fromElements((_31_v3).f31, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), function (_36_i1) {
                return new _dafny.CodePoint('k'.codePointAt(0));
              })).length), new BigNumber((_34_v6).length), new BigNumber(411));
              let _37_v8;
              _37_v8 = _dafny.Map.Empty.slice().updateUnsafe(!((_29_v1).f32),(_29_v1).f32);
              let _38_v9;
              _38_v9 = _dafny.Seq.of((_29_v1).f32, p0, p0);
              let _39_v10;
              let _nw2 = Array((new BigNumber(26)).toNumber());
              _nw2[(_dafny.ZERO).toNumber()] = _module.__default.fm5((_29_v1).f32, _32_v4, (_31_v3).f31, globalState);
              _nw2[(_dafny.ONE).toNumber()] = p0;
              _nw2[(new BigNumber(2)).toNumber()] = ((_31_v3).f31).isLessThanOrEqualTo(new BigNumber(-184));
              _nw2[(new BigNumber(3)).toNumber()] = p0;
              _nw2[(new BigNumber(4)).toNumber()] = p0;
              _nw2[(new BigNumber(5)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(6)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(7)).toNumber()] = p0;
              _nw2[(new BigNumber(8)).toNumber()] = p0;
              _nw2[(new BigNumber(9)).toNumber()] = (new BigNumber((_dafny.Set.fromElements(_33_v5, _33_v5)).length)).isLessThanOrEqualTo(_27_v0);
              _nw2[(new BigNumber(10)).toNumber()] = (_27_v0).isLessThan(new BigNumber(362));
              _nw2[(new BigNumber(11)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(12)).toNumber()] = (_27_v0).isLessThan(_27_v0);
              _nw2[(new BigNumber(13)).toNumber()] = p0;
              _nw2[(new BigNumber(14)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(15)).toNumber()] = (_35_v7).equals(_35_v7);
              _nw2[(new BigNumber(16)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(17)).toNumber()] = (new BigNumber((_34_v6).length)).isEqualTo(new BigNumber(144));
              _nw2[(new BigNumber(18)).toNumber()] = !(_37_v8).equals((_37_v8).update(p0, (_38_v9)[_module.__default.safeIndex(_27_v0, new BigNumber((_38_v9).length))]));
              _nw2[(new BigNumber(19)).toNumber()] = (new BigNumber((_34_v6).length)).isLessThan((_31_v3).f31);
              _nw2[(new BigNumber(20)).toNumber()] = p0;
              _nw2[(new BigNumber(21)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(22)).toNumber()] = p0;
              _nw2[(new BigNumber(23)).toNumber()] = (_29_v1).f32;
              _nw2[(new BigNumber(24)).toNumber()] = !(p0) || ((_29_v1).f32);
              _nw2[(new BigNumber(25)).toNumber()] = !(((_31_v3).f31).isLessThanOrEqualTo((_31_v3).f31));
              _39_v10 = _nw2;
              let _rhs0 = _39_v10;
              let _rhs1 = !((_38_v9)[_module.__default.safeIndex((_31_v3).f31, new BigNumber((_38_v9).length))]);
              let _lhs0 = globalState;
              _39_v10 = _rhs0;
              _lhs0.f1 = _rhs1;
              let _40_v11;
              let _41_v12;
              let _out0;
              let _out1;
              let _outcollector0 = (_29_v1).m7((_31_v3).fm7(globalState), _30_v2, _34_v6, globalState);
              _out0 = _outcollector0[0];
              _out1 = _outcollector0[1];
              _40_v11 = _out0;
              _41_v12 = _out1;
            } else {
              let _42_v13;
              let _nw3 = new _module.C2();
              _nw3.__ctor(_27_v0, !(p0), _module.__default.safeModuloInt(_27_v0, _27_v0));
              _42_v13 = _nw3;
              (globalState).f0 = (_dafny.ZERO).minus(_27_v0);
              let _43_v14;
              _43_v14 = _dafny.Seq.UnicodeFromString("myjeh");
              (globalState).f1 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iplftk"), _43_v14), _dafny.Seq.Concat(_43_v14, _43_v14));
              let _44_v15;
              _44_v15 = _dafny.Map.Empty.slice().updateUnsafe(_27_v0,p0);
              let _45_v16;
              _45_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hfaji")).length),_27_v0);
              (globalState).f1 = !_dafny.Seq.contains(_dafny.Seq.of(_45_v16), _module.__default.fm43(new BigNumber((_44_v15).length), p0, new BigNumber((_dafny.MultiSet.fromElements(_27_v0)).cardinality()), (_dafny.ZERO).minus((_42_v13).f31), globalState));
              let _46_v17;
              _46_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _47_v18;
              _47_v18 = _dafny.Map.Empty.slice().updateUnsafe(_27_v0,_46_v17);
              let _48_v19;
              _48_v19 = _dafny.Seq.of(new BigNumber(75), _27_v0, (_42_v13).f31);
              let _49_v20;
              _49_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_47_v18).contains(new BigNumber((_48_v19).length))) ? ((_47_v18).get(new BigNumber((_48_v19).length))) : (_46_v17)),(_42_v13).f31);
              _49_v20 = (_49_v20).update(_46_v17, _27_v0);
            }
            let _50_v21;
            _50_v21 = _module.D0.create_DC0(p0);
            let _pat_let_tv0 = p0;
            let _51_v22;
            _51_v22 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let0_0) {
              return function (_52_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_53_dt__update_hcf0_h0) {
                    return _module.D0.create_DC0(_53_dt__update_hcf0_h0);
                  }(_pat_let1_0);
                }(_pat_let_tv0);
              }(_pat_let0_0);
            }(_50_v21),_27_v0);
            let _54_v23;
            _54_v23 = _dafny.Set.fromElements(p0, false);
            let _55_v24;
            let _nw4 = new _module.C2();
            _nw4.__ctor((_dafny.ZERO).minus((_27_v0).minus(_27_v0)), _module.__default.fm5(p0, _51_v22, _27_v0, globalState), new BigNumber((_54_v23).length));
            _55_v24 = _nw4;
            let _56_v25;
            _56_v25 = _dafny.Seq.UnicodeFromString("ww");
            let _57_v26;
            _57_v26 = _dafny.Seq.of((_dafny.ZERO).minus(((_55_v24).f31).minus(new BigNumber((_56_v25).length))), _27_v0, _module.__default.safeModuloInt((_55_v24).f31, (_55_v24).f31));
            (globalState).f0 = (_dafny.ZERO).minus(new BigNumber((_57_v26).length));
            let _58_v27;
            _58_v27 = _dafny.MultiSet.fromElements(p0);
            _58_v27 = _58_v27;
          }
        }
      }
      let _59_v28;
      _59_v28 = _dafny.MultiSet.fromElements(_27_v0, _27_v0, _module.__default.fm2(_27_v0, _27_v0, globalState));
      let _60_v29;
      _60_v29 = _dafny.Seq.UnicodeFromString("hfda");
      let _61_v30;
      let _nw5 = Array((new BigNumber(11)).toNumber());
      _nw5[(_dafny.ZERO).toNumber()] = _27_v0;
      _nw5[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_27_v0, _27_v0);
      _nw5[(new BigNumber(2)).toNumber()] = _27_v0;
      _nw5[(new BigNumber(3)).toNumber()] = ((((_59_v28).contains(_27_v0)) ? ((_59_v28).get(_27_v0)) : (_27_v0))).plus(_27_v0);
      _nw5[(new BigNumber(4)).toNumber()] = _27_v0;
      _nw5[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_27_v0, _27_v0);
      _nw5[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_module.__default.fm44(globalState))).cardinality());
      _nw5[(new BigNumber(7)).toNumber()] = _27_v0;
      _nw5[(new BigNumber(8)).toNumber()] = _27_v0;
      _nw5[(new BigNumber(9)).toNumber()] = new BigNumber((_60_v29).length);
      _nw5[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("ewsm")).length), _27_v0);
      _61_v30 = _nw5;
      let _index0 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length));
      (_61_v30)[_index0] = _27_v0;
      let _index1 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length));
      (_61_v30)[_index1] = _27_v0;
      (globalState).f0 = _27_v0;
      _60_v29 = _dafny.Seq.UnicodeFromString("fba");
      (globalState).f0 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), function (_62_i2) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length);
      if (!(!(p0) || (!(p0)))) {
        (globalState).f0 = ((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))]).plus(_module.__default.safeDivisionInt(_27_v0, _27_v0));
        let _63_v31;
        _63_v31 = _dafny.Map.Empty.slice().updateUnsafe(_60_v29,(_dafny.ZERO).minus(_27_v0));
        _63_v31 = (_63_v31).update(_60_v29, new BigNumber(970));
        (globalState).f2 = (_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))];
        let _64_v32;
        _64_v32 = _dafny.Map.Empty.slice().updateUnsafe(_27_v0,p0);
        _64_v32 = (_64_v32).update(new BigNumber((_dafny.Set.fromElements(p0, p0)).length), p0);
        let _65_v33;
        _65_v33 = _module.D0.create_DC0(true);
        let _66_v34;
        _66_v34 = _module.D0.create_DC0(_module.__default.fm5(false, _dafny.Map.Empty.slice().updateUnsafe(_65_v33,new BigNumber((_60_v29).length)), _27_v0, globalState));
        let _67_v35;
        _67_v35 = _dafny.Map.Empty.slice().updateUnsafe(_65_v33,(_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))]);
        let _68_v36;
        let _nw6 = new _module.C3();
        _nw6.__ctor(_dafny.Seq.UnicodeFromString("nj"), !(p0), p0, (_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))]);
        _68_v36 = _nw6;
        let _69_v37;
        _69_v37 = _dafny.Set.fromElements(_68_v36);
        let _70_v38;
        _70_v38 = _dafny.Seq.of(false);
        let _71_v39;
        let _nw7 = Array((new BigNumber(23)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = (_68_v36).f30;
        _nw7[(_dafny.ONE).toNumber()] = (_70_v38)[_module.__default.safeIndex((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))], new BigNumber((_70_v38).length))];
        _nw7[(new BigNumber(2)).toNumber()] = p0;
        _nw7[(new BigNumber(3)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(4)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(5)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(6)).toNumber()] = p0;
        _nw7[(new BigNumber(7)).toNumber()] = p0;
        _nw7[(new BigNumber(8)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(9)).toNumber()] = p0;
        _nw7[(new BigNumber(10)).toNumber()] = false;
        _nw7[(new BigNumber(11)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(12)).toNumber()] = p0;
        _nw7[(new BigNumber(13)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(14)).toNumber()] = !((_68_v36).f30);
        _nw7[(new BigNumber(15)).toNumber()] = true;
        _nw7[(new BigNumber(16)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(17)).toNumber()] = !(p0);
        _nw7[(new BigNumber(18)).toNumber()] = p0;
        _nw7[(new BigNumber(19)).toNumber()] = (_68_v36).f30;
        _nw7[(new BigNumber(20)).toNumber()] = p0;
        _nw7[(new BigNumber(21)).toNumber()] = p0;
        _nw7[(new BigNumber(22)).toNumber()] = p0;
        _71_v39 = _nw7;
        let _72_v40;
        let _nw8 = new _module.C0();
        _nw8.__ctor(_60_v29, _module.__default.fm5(!(p0), _67_v35, (_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))], globalState), new BigNumber((_69_v37).length), _71_v39, (_68_v36).f30);
        _72_v40 = _nw8;
        let _73_v41;
        _73_v41 = _module.D17.create_DC39(_module.D17.create_DC38());
        let _74_v42;
        let _nw9 = new _module.C2();
        _nw9.__ctor((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))], false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(p0, _67_v35, new BigNumber((_70_v38).length), globalState),(_68_v36).f30)).length));
        _74_v42 = _nw9;
        let _75_v43;
        _75_v43 = _module.D17.create_DC37(_74_v42);
        let _76_v44;
        _76_v44 = _dafny.Seq.of((_73_v41).dtor_cf64, _75_v43, _75_v43, _75_v43);
        let _77_v45;
        _77_v45 = _dafny.Seq.of(_72_v40.f35);
        let _pat_let_tv1 = _76_v44;
        let _pat_let_tv2 = _61_v30;
        let _pat_let_tv3 = _61_v30;
        let _pat_let_tv4 = _76_v44;
        let _rhs2 = _72_v40;
        let _rhs3 = function (_pat_let2_0) {
          return function (_78_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_79_dt__update_hcf64_h0) {
                return _module.D17.create_DC39(_79_dt__update_hcf64_h0);
              }(_pat_let3_0);
            }((_pat_let_tv1)[_module.__default.safeIndex((_pat_let_tv3)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_pat_let_tv2).length))], new BigNumber((_pat_let_tv4).length))]);
          }(_pat_let2_0);
        }(_73_v41);
        let _rhs4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_70_v38, _70_v38), _70_v38);
        let _rhs5 = (_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))];
        let _rhs6 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))], _27_v0))).minus((new BigNumber((_77_v45).length)).multipliedBy((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))]));
        let _lhs1 = globalState;
        let _lhs2 = globalState;
        _72_v40 = _rhs2;
        _73_v41 = _rhs3;
        _70_v38 = _rhs4;
        _lhs1.f2 = _rhs5;
        _lhs2.f0 = _rhs6;
      } else {
        let _80_v46;
        _80_v46 = new _dafny.CodePoint('r'.codePointAt(0));
        _80_v46 = _80_v46;
        let _81_v47;
        let _nw10 = Array((new BigNumber(3)).toNumber());
        _81_v47 = _nw10;
        _81_v47 = _81_v47;
        let _82_v48;
        _82_v48 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), ((_83_v29, _84_v30) => function (_85_i3) {
          return (_83_v29)[_module.__default.safeIndex((_84_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_84_v30).length))], new BigNumber((_83_v29).length))];
        })(_60_v29, _61_v30)), _60_v29);
        let _86_v49;
        let _nw11 = Array((new BigNumber(12)).toNumber());
        _86_v49 = _nw11;
        let _87_v50;
        _87_v50 = _dafny.Map.Empty.slice().updateUnsafe(_82_v48,_86_v49);
        let _88_v51;
        _88_v51 = _module.D21.create_DC48(_86_v49);
        _87_v50 = (_87_v50).update(_dafny.Seq.Concat(_82_v48, _82_v48), (_88_v51).dtor_cf77);
        let _89_v52;
        let _init0 = ((_90_p0) => function (_91_i4) {
          return _90_p0;
        })(p0);
        let _nw12 = Array((new BigNumber(29)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw12.length); _i0_0++) {
          _nw12[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _89_v52 = _nw12;
        let _index2 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_89_v52).length));
        (_89_v52)[_index2] = (_59_v28).IsSubsetOf(_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))])));
        let _index3 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_89_v52).length));
        (_89_v52)[_index3] = p0;
        let _92_v53;
        let _nw13 = new _module.C1();
        _nw13.__ctor((_89_v52)[_module.__default.safeIndex(new BigNumber(70), new BigNumber((_89_v52).length))]);
        _92_v53 = _nw13;
        let _93_v54;
        _93_v54 = _module.D10.create_DC24(_27_v0, _92_v53, _80_v46);
        let _94_v55;
        _94_v55 = _module.D10.create_DC24((_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))], (_93_v54).dtor_cf37, _80_v46);
        let _pat_let_tv5 = _92_v53;
        let _pat_let_tv6 = _80_v46;
        _94_v55 = function (_pat_let4_0) {
          return function (_95_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_96_dt__update_hcf37_h0) {
                return function (_pat_let6_0) {
                  return function (_97_dt__update_hcf38_h0) {
                    return _module.D10.create_DC24((_95_dt__update__tmp_h2).dtor_cf36, _96_dt__update_hcf37_h0, _97_dt__update_hcf38_h0);
                  }(_pat_let6_0);
                }(_pat_let_tv6);
              }(_pat_let5_0);
            }(_pat_let_tv5);
          }(_pat_let4_0);
        }(_94_v55);
      }
      r0 = (_61_v30)[_module.__default.safeIndex(new BigNumber(194), new BigNumber((_61_v30).length))];
      return r0;
    }
    static Main(__noArgsParameter) {
      let _98_v0;
      let _nw14 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Set.Empty);
      _98_v0 = _nw14;
      let _99_v1;
      _99_v1 = _dafny.Seq.UnicodeFromString("oduxsblwp");
      let _100_v2;
      _100_v2 = false;
      let _101_v3;
      _101_v3 = new BigNumber(189);
      let _102_v4;
      _102_v4 = _dafny.Map.Empty.slice().updateUnsafe(_100_v2,_101_v3);
      let _103_v5;
      _103_v5 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_100_v2,_101_v3), _102_v4);
      let _104_v6;
      _104_v6 = _dafny.Set.fromElements(_101_v3);
      let _105_globalState;
      let _nw15 = new _module.GlobalState();
      _nw15.__ctor(new BigNumber(281), false, new BigNumber(-82), _98_v0, true, new BigNumber(-974), new BigNumber(777), false, _dafny.Seq.Concat(_99_v1, _dafny.Seq.UnicodeFromString("n")), _103_v5, new BigNumber(-345), true, _dafny.Seq.UnicodeFromString("yrxdsa"), _dafny.Map.Empty.slice().updateUnsafe(_100_v2,_dafny.Seq.UnicodeFromString("jxiyve")), new BigNumber(362), new BigNumber(72), new BigNumber(327), (_104_v6).Union(_104_v6), true, new BigNumber(-736), true, true, false, false, true);
      _105_globalState = _nw15;
      _100_v2 = _100_v2;
      let _106_v7;
      let _init1 = ((_107_v2, _108_v1) => function (_109_i0) {
        return !((_107_v2) && ((_module.D0.create_DC1(new BigNumber(-205), _108_v1, _107_v2, new BigNumber((_108_v1).length), _107_v2)).dtor_cf5));
      })(_100_v2, _99_v1);
      let _nw16 = Array((new BigNumber(26)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw16.length); _i0_1++) {
        _nw16[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _106_v7 = _nw16;
      let _index4 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
      (_106_v7)[_index4] = !(_100_v2) || (_100_v2);
      let _index5 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
      (_106_v7)[_index5] = (_101_v3).isLessThanOrEqualTo(_101_v3);
      let _110_i1;
      _110_i1 = _dafny.ZERO;
      L1: {
        while (!((new BigNumber((_102_v4).length)).isLessThanOrEqualTo(_101_v3)) || (!(false))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_110_i1)) {
              break L1;
            }
            _110_i1 = (_110_i1).plus(_dafny.ONE);
            (_105_globalState).f0 = _101_v3;
            let _111_v8;
            _111_v8 = _dafny.MultiSet.fromElements(_100_v2);
            let _112_v9;
            _112_v9 = new _dafny.CodePoint('w'.codePointAt(0));
            let _113_v10;
            _113_v10 = _dafny.MultiSet.fromElements(_112_v9);
            let _114_v11;
            _114_v11 = _module.D0.create_DC1(new BigNumber(315), _module.__default.fm1(_105_globalState), (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _module.__default.fm2((((_111_v8).contains(_100_v2)) ? ((_111_v8).get(_100_v2)) : (new BigNumber((_113_v10).cardinality()))), _101_v3, _105_globalState), _100_v2);
            let _115_v12;
            _115_v12 = _dafny.Seq.of(_101_v3);
            let _116_v13;
            _116_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_115_v12).length),_101_v3);
            let _117_v14;
            let _nw17 = Array((new BigNumber(9)).toNumber());
            _nw17[(_dafny.ZERO).toNumber()] = _module.__default.fm0(_101_v3, _105_globalState);
            _nw17[(_dafny.ONE).toNumber()] = _114_v11;
            _nw17[(new BigNumber(2)).toNumber()] = _114_v11;
            _nw17[(new BigNumber(3)).toNumber()] = _114_v11;
            _nw17[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_101_v3, _105_globalState);
            _nw17[(new BigNumber(5)).toNumber()] = _114_v11;
            _nw17[(new BigNumber(6)).toNumber()] = _114_v11;
            _nw17[(new BigNumber(7)).toNumber()] = _114_v11;
            _nw17[(new BigNumber(8)).toNumber()] = _module.D0.create_DC1((((_116_v13).contains(_101_v3)) ? ((_116_v13).get(_101_v3)) : (_101_v3)), _99_v1, false, _101_v3, _100_v2);
            _117_v14 = _nw17;
            let _index6 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_117_v14).length));
            (_117_v14)[_index6] = _114_v11;
            let _index7 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_117_v14).length));
            (_117_v14)[_index7] = _114_v11;
            let _118_v15;
            _118_v15 = _dafny.Map.Empty.slice().updateUnsafe(_101_v3,_100_v2);
            (_105_globalState).f0 = (_115_v12)[_module.__default.safeIndex(new BigNumber((_118_v15).length), new BigNumber((_115_v12).length))];
            _111_v8 = (_111_v8).Union(_111_v8);
          }
        }
      }
      let _hi0 = _101_v3;
      for (let _119_i2 = _101_v3; _119_i2.isLessThan(_hi0); _119_i2 = _119_i2.plus(_dafny.ONE)) {
        let _120_v16;
        _120_v16 = new _dafny.CodePoint('o'.codePointAt(0));
        _120_v16 = _120_v16;
        let _121_v17;
        _121_v17 = _dafny.Seq.of(_101_v3);
        let _122_v18;
        _122_v18 = _dafny.MultiSet.fromElements((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _100_v2, _100_v2, _100_v2);
        let _123_v19;
        let _nw18 = Array((new BigNumber(9)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = new BigNumber(-393);
        _nw18[(_dafny.ONE).toNumber()] = ((false) ? (_101_v3) : (_101_v3));
        _nw18[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(_119_i2, _119_i2);
        _nw18[(new BigNumber(3)).toNumber()] = _101_v3;
        _nw18[(new BigNumber(4)).toNumber()] = _101_v3;
        _nw18[(new BigNumber(5)).toNumber()] = _119_i2;
        _nw18[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_121_v17, _module.__default.safeIndex(_119_i2, new BigNumber((_121_v17).length)), _101_v3), _121_v17)).length);
        _nw18[(new BigNumber(7)).toNumber()] = (((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) ? (new BigNumber((_122_v18).cardinality())) : (_119_i2));
        _nw18[(new BigNumber(8)).toNumber()] = (_121_v17)[_module.__default.safeIndex(_101_v3, new BigNumber((_121_v17).length))];
        _123_v19 = _nw18;
        let _index8 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length));
        (_123_v19)[_index8] = _module.__default.safeModuloInt(_101_v3, new BigNumber((_122_v18).cardinality()));
        let _index9 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length));
        (_123_v19)[_index9] = _module.__default.safeDivisionInt(_101_v3, _101_v3);
        let _124_v20;
        _124_v20 = _module.D0.create_DC0((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
        let _source2 = _124_v20;
        if (_source2.is_DC1) {
          let _125___mcc_h0 = (_source2).cf1;
          let _126___mcc_h1 = (_source2).cf2;
          let _127___mcc_h2 = (_source2).cf3;
          let _128___mcc_h3 = (_source2).cf4;
          let _129___mcc_h4 = (_source2).cf5;
          let _130_cf5 = _129___mcc_h4;
          let _131_cf4 = _128___mcc_h3;
          let _132_cf3 = _127___mcc_h2;
          let _133_cf2 = _126___mcc_h1;
          let _134_cf1 = _125___mcc_h0;
          let _index10 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          (_106_v7)[_index10] = _130_cf5;
          (_105_globalState).f0 = _119_i2;
          let _135_v21;
          _135_v21 = _dafny.Seq.of(_132_cf3, _100_v2);
          (_105_globalState).f1 = (_135_v21)[_module.__default.safeIndex(_134_cf1, new BigNumber((_135_v21).length))];
          let _136_v22;
          _136_v22 = _module.D0.create_DC2();
          let _137_v23;
          let _nw19 = Array((new BigNumber(15)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = _136_v22;
          _nw19[(_dafny.ONE).toNumber()] = _136_v22;
          _nw19[(new BigNumber(2)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(3)).toNumber()] = _module.D0.create_DC2();
          _nw19[(new BigNumber(4)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(5)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(6)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(7)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(8)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(9)).toNumber()] = ((_100_v2) ? (_136_v22) : (_136_v22));
          _nw19[(new BigNumber(10)).toNumber()] = _module.D0.create_DC2();
          _nw19[(new BigNumber(11)).toNumber()] = _module.D0.create_DC2();
          _nw19[(new BigNumber(12)).toNumber()] = _136_v22;
          _nw19[(new BigNumber(13)).toNumber()] = _module.D0.create_DC2();
          _nw19[(new BigNumber(14)).toNumber()] = _136_v22;
          _137_v23 = _nw19;
          let _index11 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_137_v23).length));
          (_137_v23)[_index11] = _136_v22;
          let _138_v24;
          _138_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(_130_cf5), _132_cf3, _100_v2)).length),(_119_i2).multipliedBy((_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))]));
          let _139_v25;
          let _init2 = ((_140_cf2) => function (_141_i3) {
            return _140_cf2;
          })(_133_cf2);
          let _nw20 = Array((new BigNumber(18)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
            _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _139_v25 = _nw20;
          let _index12 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_139_v25).length));
          (_139_v25)[_index12] = ((_132_cf3) ? (_133_cf2) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), ((_142_v16) => function (_143_i4) {
            return _142_v16;
          })(_120_v16))));
          let _144_v26;
          let _nw21 = Array((new BigNumber(15)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _121_v17;
          _nw21[(_dafny.ONE).toNumber()] = _121_v17;
          _nw21[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_121_v17, _121_v17);
          _nw21[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_121_v17, _121_v17);
          _nw21[(new BigNumber(4)).toNumber()] = _121_v17;
          _nw21[(new BigNumber(5)).toNumber()] = ((_100_v2) ? (_121_v17) : (_dafny.Seq.of((_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))], _101_v3, _101_v3, _134_cf1, new BigNumber((_133_cf2).length))));
          _nw21[(new BigNumber(6)).toNumber()] = _121_v17;
          _nw21[(new BigNumber(7)).toNumber()] = _121_v17;
          _nw21[(new BigNumber(8)).toNumber()] = _121_v17;
          _nw21[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_121_v17, _121_v17);
          _nw21[(new BigNumber(10)).toNumber()] = _121_v17;
          _nw21[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_119_i2, _134_cf1, (_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))], _101_v3, _101_v3);
          _nw21[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))], _134_cf1), _121_v17);
          _nw21[(new BigNumber(13)).toNumber()] = _module.__default.fm3(_134_cf1, _132_cf3, _101_v3, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _105_globalState);
          _nw21[(new BigNumber(14)).toNumber()] = _121_v17;
          _144_v26 = _nw21;
          let _index13 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_144_v26).length));
          (_144_v26)[_index13] = _121_v17;
          let _index14 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_137_v23).length));
          let _index15 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_139_v25).length));
          let _index16 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_144_v26).length));
          let _rhs7 = _136_v22;
          let _rhs8 = _138_v24;
          let _rhs9 = _99_v1;
          let _rhs10 = _dafny.Seq.update(_dafny.Seq.Concat(_121_v17, _dafny.Seq.of(_101_v3, _module.__default.fm2(_119_i2, _134_cf1, _105_globalState))), _module.__default.safeIndex(_module.__default.fm2(_101_v3, new BigNumber((function () {
            let _coll5 = new _dafny.Set();
            for (const _compr_5 of _dafny.IntegerRange(new BigNumber(388), new BigNumber(253))) {
              let _145_v27 = _compr_5;
              if (((new BigNumber(388)).isLessThanOrEqualTo(_145_v27)) && ((_145_v27).isLessThan(new BigNumber(253)))) {
                _coll5.add((_145_v27).plus(new BigNumber((_121_v17).length)));
              }
            }
            return _coll5;
          }()).length), _105_globalState), new BigNumber((_dafny.Seq.Concat(_121_v17, _dafny.Seq.of(_101_v3, _module.__default.fm2(_119_i2, _134_cf1, _105_globalState)))).length)), (new BigNumber((_135_v21).length)).multipliedBy(_131_cf4));
          let _lhs3 = _137_v23;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_137_v23).length));
          let _lhs5 = _139_v25;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_139_v25).length));
          let _lhs7 = _144_v26;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_144_v26).length));
          _lhs3[_lhs4] = _rhs7;
          _138_v24 = _rhs8;
          _lhs5[_lhs6] = _rhs9;
          _lhs7[_lhs8] = _rhs10;
        } else if (_source2.is_DC2) {
          let _146_v28;
          _146_v28 = _dafny.Seq.of(_121_v17, _121_v17);
          let _index17 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _rhs11 = !(_100_v2);
          let _rhs12 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-947)), ((_147_v17) => function (_148_i5) {
            return _147_v17;
          })(_121_v17));
          let _rhs13 = (_dafny.ZERO).minus(_119_i2);
          let _lhs9 = _106_v7;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _lhs11 = _105_globalState;
          _lhs9[_lhs10] = _rhs11;
          _146_v28 = _rhs12;
          _lhs11.f0 = _rhs13;
          let _149_v29;
          _149_v29 = _dafny.MultiSet.fromElements(_120_v16, _120_v16, _module.__default.fm4(_module.D0.create_DC1(_101_v3, _99_v1, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _119_i2, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), !(_100_v2), (_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))], _105_globalState));
          let _index18 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length));
          let _rhs14 = _149_v29;
          let _rhs15 = ((_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))]).isLessThanOrEqualTo((_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))]);
          let _rhs16 = _100_v2;
          let _rhs17 = _119_i2;
          let _lhs12 = _106_v7;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _lhs14 = _123_v19;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length));
          _149_v29 = _rhs14;
          _lhs12[_lhs13] = _rhs15;
          _100_v2 = _rhs16;
          _lhs14[_lhs15] = _rhs17;
          (_105_globalState).f2 = _119_i2;
          _100_v2 = (_101_v3).isLessThan(_119_i2);
        } else {
          let _150___mcc_h5 = (_source2).cf0;
          let _151_cf0 = _150___mcc_h5;
          (_105_globalState).f1 = _100_v2;
          let _152_v30;
          let _out2;
          _out2 = _module.__default.m0(!(_151_cf0), _105_globalState);
          _152_v30 = _out2;
          let _153_v31;
          _153_v31 = _dafny.Map.Empty.slice().updateUnsafe(_124_v20,_119_i2);
          _151_cf0 = !(_module.__default.fm5(_151_cf0, (_153_v31).Merge(_153_v31), _152_v30, _105_globalState));
          let _154_v32;
          let _out3;
          _out3 = _module.__default.m0((_100_v2) === (_100_v2), _105_globalState);
          _154_v32 = _out3;
        }
        let _155_v33;
        let _nw22 = new _module.C1();
        _nw22.__ctor((_dafny.Set.fromElements(false)).equals((_module.__default.fm39((_123_v19)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_123_v19).length))], !((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), _105_globalState)).dtor_cf15));
        _155_v33 = _nw22;
      }
      let _156_v34;
      _156_v34 = _dafny.Seq.of(!(_100_v2), (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
      let _157_v35;
      _157_v35 = _dafny.Seq.of(_101_v3);
      let _158_v36;
      _158_v36 = new _dafny.CodePoint('r'.codePointAt(0));
      let _159_v38;
      _159_v38 = _module.D0.create_DC0((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
      let _160_v39;
      _160_v39 = _dafny.Seq.of(_module.D0.create_DC0((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), _module.D0.create_DC0(_100_v2), _159_v38);
      let _161_v40;
      _161_v40 = _dafny.Map.Empty.slice().updateUnsafe(_158_v36,function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_160_v39).Elements) {
          let _162_v37 = _compr_6;
          if (_dafny.Seq.contains(_160_v39, _162_v37)) {
            _coll6.push([_162_v37,new BigNumber((_102_v4).length)]);
          }
        }
        return _coll6;
      }());
      let _163_v41;
      let _nw23 = new _module.C3();
      _nw23.__ctor(((!((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))])) ? (_99_v1) : (_99_v1)), _100_v2, _module.__default.fm5((_156_v34)[_module.__default.safeIndex(new BigNumber((_157_v35).length), new BigNumber((_156_v34).length))], (((_161_v40).contains(_158_v36)) ? ((_161_v40).get(_158_v36)) : (_dafny.Map.Empty.slice().updateUnsafe(_159_v38,new BigNumber((_99_v1).length)))), _101_v3, _105_globalState), _module.__default.safeDivisionInt(_101_v3, _101_v3));
      _163_v41 = _nw23;
      let _164_v42;
      let _nw24 = new _module.C2();
      _nw24.__ctor((_dafny.ZERO).minus(_101_v3), !((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), _101_v3);
      _164_v42 = _nw24;
      _164_v42 = _164_v42;
      let _165_v43;
      let _nw25 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
      _165_v43 = _nw25;
      let _166_v44;
      let _nw26 = Array((new BigNumber(21)).toNumber());
      _166_v44 = _nw26;
      let _index20 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_165_v43).length));
      (_165_v43)[_index20] = _dafny.Map.Empty.slice().updateUnsafe((_164_v42).f31,_166_v44);
      let _167_v45;
      _167_v45 = _dafny.Map.Empty.slice().updateUnsafe((_164_v42).f31,_166_v44);
      let _index21 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_165_v43).length));
      (_165_v43)[_index21] = _167_v45;
      let _hi1 = (_164_v42).f31;
      for (let _168_i6 = (_163_v41).fm6(_100_v2, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _101_v3, _105_globalState); _168_i6.isLessThan(_hi1); _168_i6 = _168_i6.plus(_dafny.ONE)) {
        if ((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) {
          let _169_v46;
          _169_v46 = _dafny.Map.Empty.slice().updateUnsafe(_168_i6,_156_v34);
          let _170_v47;
          _170_v47 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_101_v3,_156_v34), _169_v46, _169_v46);
          _102_v4 = (_102_v4).update((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], new BigNumber(((_170_v47)[_module.__default.safeIndex((_164_v42).f31, new BigNumber((_170_v47).length))]).length));
          _158_v36 = new _dafny.CodePoint('f'.codePointAt(0));
          let _171_v48;
          _171_v48 = _dafny.Map.Empty.slice().updateUnsafe(_100_v2,_99_v1);
          (_105_globalState).f1 = (((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) ? (!((_163_v41).f30)) : (!((_dafny.ZERO).minus(_101_v3)).isEqualTo(new BigNumber(((((_171_v48).contains(_100_v2)) ? ((_171_v48).get(_100_v2)) : ((_163_v41).f29))).length))));
          let _172_v49;
          _172_v49 = _dafny.Map.Empty.slice().updateUnsafe((_164_v42).fm11(_dafny.Map.Empty.slice().updateUnsafe(_99_v1,_99_v1), (_164_v42).f31, new BigNumber(((_163_v41).f29).length), _105_globalState),_163_v41);
          _172_v49 = (_172_v49).update(_168_i6, _163_v41);
          let _173_v50;
          let _nw27 = Array((new BigNumber(19)).toNumber()).fill([]);
          _173_v50 = _nw27;
          let _174_v51;
          let _init3 = ((_175_v42) => function (_176_i7) {
            return (_176_i7).minus((_175_v42).f31);
          })(_164_v42);
          let _nw28 = Array((new BigNumber(17)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw28.length); _i0_3++) {
            _nw28[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _174_v51 = _nw28;
          let _index22 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_173_v50).length));
          (_173_v50)[_index22] = _174_v51;
          let _177_v52;
          _177_v52 = _module.D12.create_DC28(_174_v51);
          let _index23 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_173_v50).length));
          (_173_v50)[_index23] = (_177_v52).dtor_cf44;
        } else {
          (_105_globalState).f2 = _module.__default.safeModuloInt(new BigNumber(805), _168_i6);
          let _rhs18 = (_163_v41).f29;
          let _rhs19 = (_164_v42).fm6((_163_v41).f30, (_163_v41).f30, (_dafny.ZERO).minus((_164_v42).f31), _105_globalState);
          let _rhs20 = _100_v2;
          let _lhs16 = _105_globalState;
          _99_v1 = _rhs18;
          _lhs16.f2 = _rhs19;
          _100_v2 = _rhs20;
          let _178_v53;
          _178_v53 = _dafny.Set.fromElements(_106_v7, _106_v7, _106_v7, _106_v7);
          (_105_globalState).f1 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_99_v1, _module.__default.safeIndex((_164_v42).f31, new BigNumber((_99_v1).length)), _158_v36), _99_v1)).length)).isEqualTo(((_164_v42).f31).multipliedBy(new BigNumber((_178_v53).length)));
          _106_v7 = (((_163_v41).f30) ? (_106_v7) : (_106_v7));
          let _179_v54;
          let _init4 = ((_180_v41, _181_v3) => function (_182_i8) {
            return _module.D9.create_DC19(_module.D6.create_DC14(new BigNumber(((_180_v41).f29).length), _181_v3, _dafny.MultiSet.fromElements((_180_v41).f30)));
          })(_163_v41, _101_v3);
          let _nw29 = Array((new BigNumber(29)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw29.length); _i0_4++) {
            _nw29[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _179_v54 = _nw29;
          _179_v54 = _179_v54;
        }
        let _183_v55;
        let _nw30 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
        _183_v55 = _nw30;
        let _184_v56;
        _184_v56 = _dafny.Map.Empty.slice().updateUnsafe(((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) && ((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]),_183_v55);
        _183_v55 = (((_184_v56).contains((_163_v41).f30)) ? ((_184_v56).get((_163_v41).f30)) : (_183_v55));
        (_105_globalState).f2 = (_dafny.ZERO).minus((_module.__default.safeModuloInt((_164_v42).f31, (_164_v42).f31)).plus(new BigNumber(377)));
        let _185_v57;
        let _nw31 = new _module.C1();
        _nw31.__ctor((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
        _185_v57 = _nw31;
      }
      if (!(!((_163_v41).f30) || (!(!_dafny.areEqual((_163_v41).f29, _99_v1))))) {
        if ((_163_v41).f30) {
          let _186_v58;
          let _nw32 = new _module.C1();
          _nw32.__ctor((_163_v41).f30);
          _186_v58 = _nw32;
          _186_v58 = _186_v58;
          let _187_v59;
          _187_v59 = _dafny.MultiSet.fromElements((_163_v41).f30, (_186_v58).f32, (_186_v58).f32);
          let _188_v60;
          _188_v60 = _dafny.Map.Empty.slice().updateUnsafe((_164_v42).f31,!((_163_v41).f30));
          let _rhs21 = (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))];
          let _rhs22 = ((false) ? ((_164_v42).f31) : ((new BigNumber(((_163_v41).f29).length)).plus((_164_v42).f31)));
          let _rhs23 = (_163_v41).f30;
          let _rhs24 = (_164_v42).f31;
          let _rhs25 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((((_187_v59).contains((_163_v41).f30)) ? ((_187_v59).get((_163_v41).f30)) : ((_164_v42).f31)),(_163_v41).f30)).Merge(_188_v60)).length);
          let _lhs17 = _105_globalState;
          let _lhs18 = _105_globalState;
          let _lhs19 = _105_globalState;
          let _lhs20 = _105_globalState;
          let _lhs21 = _105_globalState;
          _lhs17.f1 = _rhs21;
          _lhs18.f0 = _rhs22;
          _lhs19.f1 = _rhs23;
          _lhs20.f2 = _rhs24;
          _lhs21.f0 = _rhs25;
          let _189_v61;
          _189_v61 = _module.D17.create_DC37(_164_v42);
          let _190_v62;
          _190_v62 = _dafny.Seq.of(_164_v42, _164_v42, (_189_v61).dtor_cf63, (((_186_v58).f32) ? (_164_v42) : (_164_v42)), _164_v42);
          _190_v62 = _dafny.Seq.Concat(_dafny.Seq.update(_190_v62, _module.__default.safeIndex(new BigNumber((_156_v34).length), new BigNumber((_190_v62).length)), _164_v42), _dafny.Seq.of(_164_v42, _164_v42, _164_v42, _164_v42, _164_v42));
          (_105_globalState).f2 = (((_156_v34)[_module.__default.safeIndex((_164_v42).f31, new BigNumber((_156_v34).length))]) ? ((_164_v42).f31) : (_101_v3));
          let _191_v63;
          let _nw33 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _191_v63 = _nw33;
          let _192_v64;
          _192_v64 = _dafny.Seq.of(_191_v63, (((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) ? (_191_v63) : (_191_v63)), _191_v63, _191_v63);
          (_105_globalState).f0 = new BigNumber((_192_v64).length);
        } else {
          let _193_v65;
          let _nw34 = new _module.C4();
          _nw34.__ctor(_101_v3, (_163_v41).f29, (_163_v41).f30, (_164_v42).f31);
          _193_v65 = _nw34;
          let _194_v66;
          let _nw35 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _194_v66 = _nw35;
          _194_v66 = _194_v66;
          let _195_v67;
          _195_v67 = _module.D16.create_DC36((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _157_v35, _106_v7);
          let _196_v68;
          let _197_v69;
          let _198_v70;
          let _199_v71;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = (_193_v65).m2((_195_v67).dtor_cf60, _105_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _196_v68 = _out4;
          _197_v69 = _out5;
          _198_v70 = _out6;
          _199_v71 = _out7;
          _99_v1 = (_163_v41).f29;
          let _200_v72;
          let _nw36 = Array((new BigNumber(6)).toNumber()).fill([]);
          _200_v72 = _nw36;
          let _index24 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_200_v72).length));
          (_200_v72)[_index24] = _194_v66;
          let _201_v73;
          let _nw37 = Array((new BigNumber(7)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _157_v35;
          _nw37[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_157_v35, _157_v35);
          _nw37[(new BigNumber(2)).toNumber()] = _157_v35;
          _nw37[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-357)), ((_202_v42) => function (_203_i9) {
            return (_202_v42).f31;
          })(_164_v42));
          _nw37[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(430)), ((_204_v65) => function (_205_i10) {
            return (_204_v65).f27;
          })(_193_v65));
          _nw37[(new BigNumber(5)).toNumber()] = _157_v35;
          _nw37[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_206_v42) => function (_207_i11) {
            return (_206_v42).f31;
          })(_164_v42));
          _201_v73 = _nw37;
          let _index25 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_201_v73).length));
          (_201_v73)[_index25] = _dafny.Seq.Concat(_157_v35, _157_v35);
          let _index26 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_200_v72).length));
          let _index27 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_201_v73).length));
          let _rhs26 = _194_v66;
          let _rhs27 = (_164_v42).f31;
          let _rhs28 = (_module.__default.fm40(_100_v2, _105_globalState)).dtor_cf55;
          let _rhs29 = _dafny.Seq.of(_101_v3, new BigNumber(((((_163_v41).f30) ? (_module.__default.fm1(_105_globalState)) : ((_163_v41).f29))).length), (_164_v42).f31);
          let _lhs22 = _200_v72;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_200_v72).length));
          let _lhs24 = _105_globalState;
          let _lhs25 = _193_v65;
          let _lhs26 = _201_v73;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_201_v73).length));
          _lhs22[_lhs23] = _rhs26;
          _lhs24.f2 = _rhs27;
          _lhs25.f28 = _rhs28;
          _lhs26[_lhs27] = _rhs29;
        }
        (_105_globalState).f1 = (_163_v41).f30;
        _156_v34 = _dafny.Seq.Concat(_156_v34, _156_v34);
        let _208_v74;
        _208_v74 = _dafny.Map.Empty.slice().updateUnsafe((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))],_100_v2);
        let _209_v75;
        let _nw38 = new _module.C4();
        _nw38.__ctor(new BigNumber((_208_v74).length), _dafny.Seq.update(_99_v1, _module.__default.safeIndex(new BigNumber(((_163_v41).f29).length), new BigNumber((_99_v1).length)), _158_v36), _100_v2, (_164_v42).f31);
        _209_v75 = _nw38;
        let _210_v76;
        _210_v76 = _dafny.MultiSet.fromElements(_209_v75);
        let _211_v77;
        _211_v77 = _dafny.Seq.of(_209_v75);
        let _212_v78;
        _212_v78 = _dafny.MultiSet.FromArray(_211_v77);
        let _213_v79;
        _213_v79 = _dafny.Set.fromElements(_210_v76, (_210_v76).Difference(_210_v76), _210_v76, (_212_v78), _210_v76);
        let _rhs30 = _101_v3;
        let _rhs31 = _module.__default.fm2(((_209_v75).f27).minus(_101_v3), (_209_v75).fm6(_100_v2, (_163_v41).f30, _101_v3, _105_globalState), _105_globalState);
        let _rhs32 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(_209_v75, _209_v75, _209_v75));
        let _rhs33 = _module.__default.safeModuloInt((_163_v41).fm9(_105_globalState), new BigNumber(915));
        let _lhs28 = _105_globalState;
        let _lhs29 = _105_globalState;
        _101_v3 = _rhs30;
        _lhs28.f0 = _rhs31;
        _213_v79 = _rhs32;
        _lhs29.f0 = _rhs33;
        let _214_v80;
        _214_v80 = _dafny.Set.fromElements(_156_v34, _156_v34, _156_v34, _156_v34);
        let _215_v81;
        _215_v81 = _dafny.Seq.of(_156_v34, _dafny.Seq.of((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), _156_v34, _156_v34, _156_v34);
        (_105_globalState).f0 = (_dafny.ZERO).minus(new BigNumber(((_214_v80).Union(function () {
          let _coll7 = new _dafny.Set();
          for (const _compr_7 of (_215_v81).Elements) {
            let _216_v82 = _compr_7;
            if (_dafny.Seq.contains(_215_v81, _216_v82)) {
              _coll7.add(_216_v82);
            }
          }
          return _coll7;
        }())).length));
      } else {
        _99_v1 = _99_v1;
        if (((!((_163_v41).f30)) ? (_100_v2) : (!(true)))) {
          (_105_globalState).f0 = _101_v3;
          (_105_globalState).f0 = new BigNumber(338);
          let _217_v83;
          let _nw39 = new _module.C1();
          _nw39.__ctor(false);
          _217_v83 = _nw39;
          let _218_v84;
          let _nw40 = new _module.C4();
          _nw40.__ctor(new BigNumber(903), (_163_v41).f29, (_217_v83).f32, (_164_v42).f31);
          _218_v84 = _nw40;
          let _219_v85;
          _219_v85 = _dafny.MultiSet.fromElements(_218_v84, _218_v84);
          let _220_v86;
          _220_v86 = _module.D19.create_DC41(_218_v84);
          let _221_v87;
          _221_v87 = _dafny.Seq.of(_dafny.Set.fromElements((_218_v84).f27));
          let _222_v88;
          _222_v88 = _dafny.Seq.of(_157_v35);
          let _223_v89;
          _223_v89 = _dafny.Seq.of(_106_v7);
          let _224_v90;
          _224_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_218_v84).f27)).length),_dafny.Seq.update(_99_v1, _module.__default.safeIndex((_218_v84).f27, new BigNumber((_99_v1).length)), _158_v36));
          let _225_v91;
          _225_v91 = _dafny.Map.Empty.slice().updateUnsafe(!((_163_v41).f30),_106_v7);
          let _226_v92;
          _226_v92 = _module.D20.create_DC44(_225_v91);
          let _227_v93;
          _227_v93 = _dafny.Map.Empty.slice().updateUnsafe((_226_v92).dtor_cf73,_dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), ((_228_v36) => function (_229_i13) {
            return _228_v36;
          })(_158_v36)));
          let _230_v94;
          _230_v94 = _dafny.Set.fromElements(!((_164_v42).fm7(_105_globalState)), (_217_v83).f32);
          let _231_v95;
          _231_v95 = _dafny.Map.Empty.slice().updateUnsafe(_230_v94,false);
          let _232_v96;
          _232_v96 = _dafny.Map.Empty.slice().updateUnsafe(_231_v95,(_164_v42).f31);
          let _233_v97;
          let _nw41 = Array((new BigNumber(22)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = (((_219_v85).contains((_220_v86).dtor_cf66)) ? ((_219_v85).get((_220_v86).dtor_cf66)) : ((_218_v84).f27));
          _nw41[(_dafny.ONE).toNumber()] = (new BigNumber(((_163_v41).f29).length)).minus((_218_v84).f27);
          _nw41[(new BigNumber(2)).toNumber()] = (_164_v42).fm11(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_234_v36) => function (_235_i12) {
            return _234_v36;
          })(_158_v36)),_99_v1), _101_v3, (_164_v42).f31, _105_globalState);
          _nw41[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_101_v3, (_164_v42).f31);
          _nw41[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_221_v87).length), (_218_v84).f27);
          _nw41[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_218_v84).f27);
          _nw41[(new BigNumber(6)).toNumber()] = new BigNumber((_222_v88).length);
          _nw41[(new BigNumber(7)).toNumber()] = (_164_v42).f31;
          _nw41[(new BigNumber(8)).toNumber()] = (new BigNumber(671)).plus((_218_v84).f27);
          _nw41[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_164_v42).f31);
          _nw41[(new BigNumber(10)).toNumber()] = _101_v3;
          _nw41[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_223_v89)[_module.__default.safeIndex(new BigNumber((_224_v90).length), new BigNumber((_223_v89).length))],(_164_v42).f31)).length);
          _nw41[(new BigNumber(12)).toNumber()] = (_164_v42).f31;
          _nw41[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt((_164_v42).f31, (_164_v42).f31);
          _nw41[(new BigNumber(14)).toNumber()] = (_218_v84).f27;
          _nw41[(new BigNumber(15)).toNumber()] = new BigNumber((_module.__default.fm1(_105_globalState)).length);
          _nw41[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_164_v42).f31);
          _nw41[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((new BigNumber(((((_227_v93).contains(_225_v91)) ? ((_227_v93).get(_225_v91)) : ((_163_v41).f29))).length)).multipliedBy(new BigNumber(890)));
          _nw41[(new BigNumber(18)).toNumber()] = _101_v3;
          _nw41[(new BigNumber(19)).toNumber()] = new BigNumber(-639);
          _nw41[(new BigNumber(20)).toNumber()] = (_164_v42).f31;
          _nw41[(new BigNumber(21)).toNumber()] = (((_232_v96).contains(_231_v95)) ? ((_232_v96).get(_231_v95)) : ((_218_v84).f27));
          _233_v97 = _nw41;
          let _index28 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_233_v97).length));
          (_233_v97)[_index28] = (_164_v42).f31;
          let _index29 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_233_v97).length));
          (_233_v97)[_index29] = (((_218_v84).f27).multipliedBy((_164_v42).f31)).multipliedBy(new BigNumber(((_163_v41).f29).length));
          let _index30 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          (_106_v7)[_index30] = (_164_v42).fm7(_105_globalState);
        } else {
          (_105_globalState).f1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("j"), (_163_v41).f29);
          _157_v35 = _dafny.Seq.of(new BigNumber(((_163_v41).f29).length));
          let _rhs34 = (_163_v41).f29;
          let _rhs35 = (_163_v41).f30;
          let _lhs30 = _105_globalState;
          _99_v1 = _rhs34;
          _lhs30.f1 = _rhs35;
          let _236_v98;
          _236_v98 = _dafny.Set.fromElements(_99_v1);
          (_105_globalState).f1 = (_236_v98).IsSubsetOf(_236_v98);
          let _237_v99;
          _237_v99 = _dafny.Map.Empty.slice().updateUnsafe((_163_v41).f30,_100_v2);
          let _238_v100;
          _238_v100 = _dafny.Set.fromElements((_163_v41).f30);
          let _239_v101;
          let _nw42 = new _module.C4();
          _nw42.__ctor((_164_v42).f31, (_163_v41).f29, (_163_v41).f30, new BigNumber((_238_v100).length));
          _239_v101 = _nw42;
          let _240_v102;
          _240_v102 = _dafny.Set.fromElements(_239_v101, _239_v101);
          let _241_v103;
          _241_v103 = _module.D10.create_DC23();
          let _242_v104;
          let _nw43 = new _module.C0();
          _nw43.__ctor(_239_v101.f28, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _101_v3, _106_v7, (_163_v41).f30);
          _242_v104 = _nw43;
          let _243_v105;
          _243_v105 = _dafny.Map.Empty.slice().updateUnsafe(_242_v104,_157_v35);
          let _244_v107;
          _244_v107 = _dafny.MultiSet.fromElements((((_102_v4).contains((_242_v104).fm17(!((_163_v41).f30), (_163_v41).f30, _105_globalState))) ? ((_102_v4).get((_242_v104).fm17(!((_163_v41).f30), (_163_v41).f30, _105_globalState))) : ((_164_v42).f31)), (_164_v42).f31, (_164_v42).f31);
          let _245_v108;
          let _nw44 = new _module.C3();
          _nw44.__ctor(_239_v101.f28, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], new BigNumber(637));
          _245_v108 = _nw44;
          let _246_v109;
          _246_v109 = _dafny.Map.Empty.slice().updateUnsafe((_164_v42).f31,_245_v108);
          let _247_v110;
          let _init5 = ((_248_v3) => function (_249_i14) {
            return _module.__default.safeDivisionInt(_249_i14, _248_v3);
          })(_101_v3);
          let _nw45 = Array((new BigNumber(10)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw45.length); _i0_5++) {
            _nw45[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _247_v110 = _nw45;
          let _250_v111;
          _250_v111 = _dafny.Seq.of(_247_v110);
          let _251_v112;
          _251_v112 = _dafny.Seq.of(_module.__default.fm3(new BigNumber((_246_v109).length), (_163_v41).f30, new BigNumber((_250_v111).length), !((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), _105_globalState));
          let _252_v113;
          let _nw46 = Array((new BigNumber(28)).toNumber());
          _nw46[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_163_v41).fm6((_163_v41).f30, (_163_v41).f30, (_164_v42).f31, _105_globalState));
          _nw46[(_dafny.ONE).toNumber()] = _157_v35;
          _nw46[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_157_v35, _dafny.Seq.of(new BigNumber(125), new BigNumber((_237_v99).length), (_164_v42).f31, (_dafny.ZERO).minus((_164_v42).f31)));
          _nw46[(new BigNumber(3)).toNumber()] = (((_163_v41).f30) ? (_dafny.Seq.update(_157_v35, _module.__default.safeIndex((_dafny.ZERO).minus((_164_v42).f31), new BigNumber((_157_v35).length)), (_164_v42).f31)) : (_157_v35));
          _nw46[(new BigNumber(4)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_157_v35, _157_v35);
          _nw46[(new BigNumber(6)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(7)).toNumber()] = _dafny.Seq.of((_164_v42).f31);
          _nw46[(new BigNumber(8)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(9)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_240_v102)).cardinality()));
          _nw46[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_241_v103, _module.D10.create_DC23()),_106_v7)).length), new BigNumber((_157_v35).length), (_164_v42).f31, _101_v3, (_dafny.ZERO).minus((_164_v42).f31));
          _nw46[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(new BigNumber(811));
          _nw46[(new BigNumber(13)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(14)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(15)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(16)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(new BigNumber(307), (_164_v42).f31);
          _nw46[(new BigNumber(18)).toNumber()] = (((_243_v105).contains(_242_v104)) ? ((_243_v105).get(_242_v104)) : (_157_v35));
          _nw46[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_157_v35, _157_v35);
          _nw46[(new BigNumber(20)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_157_v35, _module.__default.safeIndex(_101_v3, new BigNumber((_157_v35).length)), _101_v3);
          _nw46[(new BigNumber(22)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(23)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_157_v35, _157_v35);
          _nw46[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_157_v35, _dafny.Seq.update(_157_v35, _module.__default.safeIndex(new BigNumber((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_244_v107).Elements) {
              let _253_v106 = _compr_8;
              if ((_244_v107).contains(_253_v106)) {
                _coll8.push([(_253_v106).minus((_164_v42).f31),new BigNumber(166)]);
              }
            }
            return _coll8;
          }()).length), new BigNumber((_157_v35).length)), (_164_v42).f31));
          _nw46[(new BigNumber(26)).toNumber()] = _157_v35;
          _nw46[(new BigNumber(27)).toNumber()] = (_251_v112)[_module.__default.safeIndex((_dafny.ZERO).minus((_245_v108).f26), new BigNumber((_251_v112).length))];
          _252_v113 = _nw46;
          let _254_v114;
          _254_v114 = _dafny.Seq.of((((_163_v41).f30) ? (_252_v113) : (_252_v113)), _252_v113);
          _252_v113 = (_254_v114)[_module.__default.safeIndex((_245_v108).f26, new BigNumber((_254_v114).length))];
        }
        let _255_v115;
        let _nw47 = new _module.C0();
        _nw47.__ctor(_dafny.Seq.UnicodeFromString("gkpbjugb"), (_163_v41).f30, (_164_v42).f31, _106_v7, ((_163_v41).f30) || (!((_163_v41).fm7(_105_globalState))));
        _255_v115 = _nw47;
        _255_v115 = _255_v115;
        let _256_v116;
        _256_v116 = _dafny.Map.Empty.slice().updateUnsafe(!(true) || (_100_v2),(_163_v41).f30);
        let _257_v117;
        _257_v117 = _dafny.MultiSet.fromElements(_100_v2, (_163_v41).f30, _100_v2, (_163_v41).f30);
        _256_v116 = (_256_v116).update((_module.__default.fm27(_101_v3, _105_globalState)).IsProperSubsetOf(_257_v117), _100_v2);
        (_105_globalState).f1 = (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))];
      }
      let _258_v118;
      _258_v118 = _dafny.Map.Empty.slice().updateUnsafe((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))],_158_v36);
      if (_dafny.Seq.contains(_156_v34, !(_258_v118).equals(_258_v118))) {
        let _259_v119;
        _259_v119 = _dafny.Set.fromElements((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
        let _index31 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_98_v0).length));
        (_98_v0)[_index31] = _259_v119;
        let _index32 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_98_v0).length));
        (_98_v0)[_index32] = _259_v119;
        let _260_v120;
        _260_v120 = _module.D1.create_DC4(_100_v2, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
        let _261_v121;
        _261_v121 = _module.D1.create_DC5(_260_v120);
        let _262_v122;
        _262_v122 = _module.D0.create_DC1(_101_v3, (_163_v41).f29, (_163_v41).f30, (_164_v42).f31, (_163_v41).f30);
        let _263_v123;
        let _264_v124;
        let _out8;
        let _out9;
        let _outcollector2 = (_164_v42).m5(_100_v2, _261_v121, _module.__default.fm4(_262_v122, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], new BigNumber(((_163_v41).f29).length), _105_globalState), _105_globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _263_v123 = _out8;
        _264_v124 = _out9;
        _158_v36 = _158_v36;
        _101_v3 = (((_dafny.ZERO).minus(_101_v3)).minus(new BigNumber(554))).minus((new BigNumber((_157_v35).length)).multipliedBy(_101_v3));
        let _265_v125;
        let _nw48 = Array((new BigNumber(26)).toNumber()).fill([]);
        _265_v125 = _nw48;
        let _266_v126;
        _266_v126 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(172),_156_v34);
        let _267_v127;
        let _nw49 = new _module.C4();
        _nw49.__ctor(new BigNumber(-205), (_163_v41).f29, (_163_v41).f30, new BigNumber((_259_v119).length));
        _267_v127 = _nw49;
        let _268_v128;
        _268_v128 = _dafny.Set.fromElements(_267_v127, _267_v127);
        let _269_v129;
        let _nw50 = Array((new BigNumber(25)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("eblm")).length);
        _nw50[(_dafny.ONE).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(2)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(3)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(4)).toNumber()] = _101_v3;
        _nw50[(new BigNumber(5)).toNumber()] = new BigNumber(-139);
        _nw50[(new BigNumber(6)).toNumber()] = new BigNumber(420);
        _nw50[(new BigNumber(7)).toNumber()] = _101_v3;
        _nw50[(new BigNumber(8)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(9)).toNumber()] = _101_v3;
        _nw50[(new BigNumber(10)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(11)).toNumber()] = new BigNumber((_266_v126).length);
        _nw50[(new BigNumber(12)).toNumber()] = new BigNumber(952);
        _nw50[(new BigNumber(13)).toNumber()] = new BigNumber(((_98_v0)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_98_v0).length))]).length);
        _nw50[(new BigNumber(14)).toNumber()] = _101_v3;
        _nw50[(new BigNumber(15)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(16)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(17)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(18)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(19)).toNumber()] = _101_v3;
        _nw50[(new BigNumber(20)).toNumber()] = new BigNumber(576);
        _nw50[(new BigNumber(21)).toNumber()] = (_164_v42).f31;
        _nw50[(new BigNumber(22)).toNumber()] = new BigNumber((_268_v128).length);
        _nw50[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.of(_263_v123, _263_v123, false)).length);
        _nw50[(new BigNumber(24)).toNumber()] = (_267_v127).f27;
        _269_v129 = _nw50;
        let _index33 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_265_v125).length));
        (_265_v125)[_index33] = _269_v129;
        let _index34 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_265_v125).length));
        let _index35 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
        let _rhs36 = _269_v129;
        let _rhs37 = _module.__default.safeModuloInt(new BigNumber(89), (_dafny.ZERO).minus(((_164_v42).f31).multipliedBy((_164_v42).f31)));
        let _rhs38 = (_267_v127).fm7(_105_globalState);
        let _lhs31 = _265_v125;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_265_v125).length));
        let _lhs33 = _105_globalState;
        let _lhs34 = _106_v7;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
        _lhs31[_lhs32] = _rhs36;
        _lhs33.f0 = _rhs37;
        _lhs34[_lhs35] = _rhs38;
      } else {
        let _270_v130;
        _270_v130 = _dafny.Seq.of(_163_v41, _163_v41);
        let _271_v131;
        let _out10;
        _out10 = (_164_v42).m6(_dafny.Map.Empty.slice().updateUnsafe((_164_v42).f31,(_164_v42).f31), new BigNumber((_270_v130).length), new BigNumber(257), _105_globalState);
        _271_v131 = _out10;
        if (!((((_163_v41).f30) ? (_dafny.Seq.IsProperPrefixOf(_99_v1, _99_v1)) : ((_163_v41).f30)))) {
          let _272_v132;
          let _nw51 = Array((new BigNumber(10)).toNumber()).fill([]);
          _272_v132 = _nw51;
          let _273_v133;
          let _init6 = ((_274_v42) => function (_275_i15) {
            return (_275_i15).plus((_274_v42).f31);
          })(_164_v42);
          let _nw52 = Array((_dafny.ONE).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw52.length); _i0_6++) {
            _nw52[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _273_v133 = _nw52;
          let _index36 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_272_v132).length));
          (_272_v132)[_index36] = _273_v133;
          let _index37 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_272_v132).length));
          (_272_v132)[_index37] = _273_v133;
          let _276_v134;
          let _nw53 = new _module.C4();
          _nw53.__ctor((_dafny.ZERO).minus((_164_v42).f31), (_163_v41).f29, (_163_v41).f30, (_164_v42).f31);
          _276_v134 = _nw53;
          let _277_v135;
          _277_v135 = _dafny.Set.fromElements(_276_v134);
          let _278_v136;
          _278_v136 = _dafny.Seq.of(_277_v135);
          let _279_v137;
          _279_v137 = _module.D20.create_DC46(true, (_164_v42).f31);
          let _280_v138;
          let _out11;
          _out11 = (_164_v42).m6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm35((_dafny.ZERO).minus(_101_v3), new BigNumber((_dafny.MultiSet.FromArray(_278_v136)).cardinality()), _105_globalState)).cardinality()),(_279_v137).dtor_cf76), _101_v3, (_276_v134).fm6(_100_v2, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _101_v3, _105_globalState), _105_globalState);
          _280_v138 = _out11;
          let _281_v139;
          _281_v139 = _module.D2.create_DC6(_104_v6);
          let _282_v140;
          _282_v140 = _dafny.Seq.of(_module.D2.create_DC6(_104_v6), _281_v139, _281_v139);
          _282_v140 = _dafny.Seq.update(_dafny.Seq.Concat(_282_v140, _dafny.Seq.Create(_module.__default.abs(new BigNumber(512)), ((_283_v139) => function (_284_i16) {
            return _283_v139;
          })(_281_v139))), _module.__default.safeIndex((((_163_v41).f30) ? (_101_v3) : ((_164_v42).f31)), new BigNumber((_dafny.Seq.Concat(_282_v140, _dafny.Seq.Create(_module.__default.abs(new BigNumber(512)), ((_285_v139) => function (_286_i16) {
            return _285_v139;
          })(_281_v139)))).length)), _module.D2.create_DC6(_104_v6));
          let _287_v141;
          _287_v141 = _dafny.MultiSet.fromElements((_164_v42).f31);
          let _288_v142;
          let _out12;
          _out12 = (_163_v41).m1(_287_v141, (_163_v41).f30, _100_v2, _105_globalState);
          _288_v142 = _out12;
          let _pat_let_tv7 = _164_v42;
          let _pat_let_tv8 = _164_v42;
          _101_v3 = (function (_pat_let7_0) {
            return function (_289_dt__update__tmp_h0) {
              return function (_pat_let8_0) {
                return function (_290_dt__update_hcf24_h0) {
                  return function (_pat_let9_0) {
                    return function (_291_dt__update_hcf23_h0) {
                      return _module.D6.create_DC14(_291_dt__update_hcf23_h0, _290_dt__update_hcf24_h0, (_289_dt__update__tmp_h0).dtor_cf25);
                    }(_pat_let9_0);
                  }((_pat_let_tv8).f31);
                }(_pat_let8_0);
              }((_pat_let_tv7).f31);
            }(_pat_let7_0);
          }(_module.__default.fm41(_288_v142, true, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _105_globalState))).dtor_cf23;
        } else {
          _104_v6 = (_module.D2.create_DC6(_104_v6)).dtor_cf10;
          let _index38 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          (_106_v7)[_index38] = _100_v2;
          let _292_v143;
          let _nw54 = new _module.C1();
          _nw54.__ctor((_163_v41).f30);
          _292_v143 = _nw54;
          let _293_v144;
          _293_v144 = _module.D10.create_DC24(new BigNumber(547), _292_v143, _158_v36);
          let _294_v145;
          _294_v145 = _dafny.MultiSet.fromElements(new BigNumber(((_163_v41).f29).length));
          let _295_v146;
          _295_v146 = _dafny.MultiSet.fromElements(_158_v36);
          let _296_v147;
          _296_v147 = _dafny.Seq.of(_164_v42);
          let _297_v148;
          _297_v148 = _module.D19.create_DC43((_296_v147)[_module.__default.safeIndex((_164_v42).f31, new BigNumber((_296_v147).length))], (_164_v42).f31, (_292_v143).f32);
          let _298_v149;
          _298_v149 = _dafny.MultiSet.fromElements(true, _100_v2);
          let _299_v150;
          _299_v150 = _module.D6.create_DC13((_164_v42).f31, (_164_v42).f31, _dafny.Set.fromElements(new BigNumber((_298_v149).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("ygqeqe")).length), (_164_v42).f31));
          let _300_v151;
          let _nw55 = Array((new BigNumber(10)).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = (_163_v41).fm6((_163_v41).f30, _100_v2, (_164_v42).f31, _105_globalState);
          _nw55[(_dafny.ONE).toNumber()] = new BigNumber(969);
          _nw55[(new BigNumber(2)).toNumber()] = new BigNumber((_99_v1).length);
          _nw55[(new BigNumber(3)).toNumber()] = (_164_v42).f31;
          _nw55[(new BigNumber(4)).toNumber()] = (_164_v42).f31;
          _nw55[(new BigNumber(5)).toNumber()] = (_299_v150).dtor_cf20;
          _nw55[(new BigNumber(6)).toNumber()] = _module.__default.fm2((_164_v42).f31, new BigNumber(421), _105_globalState);
          _nw55[(new BigNumber(7)).toNumber()] = (_164_v42).f31;
          _nw55[(new BigNumber(8)).toNumber()] = (_164_v42).f31;
          _nw55[(new BigNumber(9)).toNumber()] = (_164_v42).f31;
          _300_v151 = _nw55;
          let _301_v152;
          _301_v152 = _dafny.Set.fromElements(_300_v151);
          let _302_v153;
          let _nw56 = new _module.C0();
          _nw56.__ctor(_dafny.Seq.UnicodeFromString("yjpy"), false, _101_v3, _106_v7, (_292_v143).f32);
          _302_v153 = _nw56;
          let _303_v154;
          let _nw57 = Array((new BigNumber(24)).toNumber());
          _nw57[(_dafny.ZERO).toNumber()] = (_164_v42).f31;
          _nw57[(_dafny.ONE).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(2)).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(776)), ((_304_v36) => function (_305_i17) {
            return _304_v36;
          })(_158_v36))).length);
          _nw57[(new BigNumber(4)).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(5)).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(6)).toNumber()] = (_293_v144).dtor_cf36;
          _nw57[(new BigNumber(7)).toNumber()] = new BigNumber(187);
          _nw57[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_164_v42).f31);
          _nw57[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt((_164_v42).f31, (_164_v42).f31);
          _nw57[(new BigNumber(10)).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(11)).toNumber()] = _101_v3;
          _nw57[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt((_164_v42).f31, (_164_v42).f31);
          _nw57[(new BigNumber(13)).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(14)).toNumber()] = _101_v3;
          _nw57[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_294_v145).cardinality()));
          _nw57[(new BigNumber(16)).toNumber()] = ((false) ? (_101_v3) : (new BigNumber((_156_v34).length)));
          _nw57[(new BigNumber(17)).toNumber()] = (((_292_v143).f32) ? (new BigNumber((_295_v146).cardinality())) : ((_164_v42).f31));
          _nw57[(new BigNumber(18)).toNumber()] = new BigNumber(979);
          _nw57[(new BigNumber(19)).toNumber()] = (_164_v42).f31;
          _nw57[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((((_292_v143).f32) ? (_101_v3) : ((_164_v42).f31)));
          _nw57[(new BigNumber(21)).toNumber()] = (_297_v148).dtor_cf71;
          _nw57[(new BigNumber(22)).toNumber()] = new BigNumber(((_301_v152).Union(_301_v152)).length);
          _nw57[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_302_v153, _302_v153)).length);
          _303_v154 = _nw57;
          let _index39 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_300_v151).length));
          (_300_v151)[_index39] = (_164_v42).f31;
          let _306_v155;
          let _nw58 = new _module.C3();
          _nw58.__ctor(_module.__default.fm1(_105_globalState), _271_v131, _271_v131, new BigNumber(56));
          _306_v155 = _nw58;
          let _307_v156;
          _307_v156 = _dafny.Map.Empty.slice().updateUnsafe(_101_v3,(_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
          let _index40 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_300_v151).length));
          let _rhs39 = (new BigNumber((_307_v156).length)).minus((_164_v42).f31);
          let _rhs40 = _306_v155;
          let _rhs41 = (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))];
          let _lhs36 = _300_v151;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_300_v151).length));
          let _lhs38 = _105_globalState;
          _lhs36[_lhs37] = _rhs39;
          _306_v155 = _rhs40;
          _lhs38.f1 = _rhs41;
          let _308_v157;
          _308_v157 = _dafny.Set.fromElements((_292_v143).f32, (_292_v143).f32);
          let _309_v158;
          let _nw59 = new _module.C2();
          _nw59.__ctor((_300_v151)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_300_v151).length))], (_292_v143).fm14(_308_v157, (_300_v151)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_300_v151).length))], _105_globalState), (_300_v151)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_300_v151).length))]);
          _309_v158 = _nw59;
          _102_v4 = _102_v4;
        }
        let _310_v159;
        let _init7 = ((_311_v36) => function (_312_i18) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-591)), ((_313_v36) => function (_314_i19) {
            return _313_v36;
          })(_311_v36));
        })(_158_v36);
        let _nw60 = Array((new BigNumber(26)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw60.length); _i0_7++) {
          _nw60[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _310_v159 = _nw60;
        let _315_v160;
        _315_v160 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-774)), ((_316_v36) => function (_317_i20) {
          return _316_v36;
        })(_158_v36)));
        let _318_v161;
        _318_v161 = _dafny.MultiSet.fromElements((_164_v42).f31);
        let _319_v162;
        let _nw61 = new _module.C0();
        _nw61.__ctor(_99_v1, _271_v131, new BigNumber((_318_v161).cardinality()), _106_v7, (_163_v41).f30);
        _319_v162 = _nw61;
        let _320_v163;
        _320_v163 = _dafny.Seq.of(_319_v162);
        let _index41 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_310_v159).length));
        (_310_v159)[_index41] = _dafny.Seq.Concat(_99_v1, (_315_v160)[_module.__default.safeIndex(new BigNumber((_320_v163).length), new BigNumber((_315_v160).length))]);
        let _index42 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_310_v159).length));
        (_310_v159)[_index42] = (_163_v41).f29;
        _101_v3 = (_164_v42).f31;
        let _321_v164;
        _321_v164 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vrd"),_100_v2);
        _321_v164 = (_321_v164).update(_dafny.Seq.UnicodeFromString("tdm"), _271_v131);
      }
      let _322_v165;
      _322_v165 = _dafny.Set.fromElements((_163_v41).f30);
      if ((_322_v165).IsSubsetOf(_322_v165)) {
        let _323_v166;
        let _nw62 = Array((_dafny.ONE).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_164_v42).f31, (_164_v42).f31);
        _323_v166 = _nw62;
        let _index43 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
        (_323_v166)[_index43] = new BigNumber(-575);
        let _324_v167;
        _324_v167 = _dafny.MultiSet.fromElements(_101_v3);
        let _index44 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
        (_323_v166)[_index44] = (new BigNumber(-461)).minus(new BigNumber((_324_v167).cardinality()));
        let _325_v168;
        _325_v168 = _dafny.Seq.of(_104_v6, _104_v6, _104_v6, _104_v6, _104_v6);
        let _326_v169;
        _326_v169 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(80),_100_v2);
        let _327_v170;
        _327_v170 = _module.D12.create_DC29(new BigNumber(381), _100_v2, _325_v168, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_100_v2,new BigNumber((_326_v169).length))).length), new BigNumber(-170));
        let _source3 = ((_100_v2) ? (_327_v170) : (_327_v170));
        if (_source3.is_DC29) {
          let _328___mcc_h6 = (_source3).cf45;
          let _329___mcc_h7 = (_source3).cf46;
          let _330___mcc_h8 = (_source3).cf47;
          let _331___mcc_h9 = (_source3).cf48;
          let _332___mcc_h10 = (_source3).cf49;
          let _333_cf49 = _332___mcc_h10;
          let _334_cf48 = _331___mcc_h9;
          let _335_cf47 = _330___mcc_h8;
          let _336_cf46 = _329___mcc_h7;
          let _337_cf45 = _328___mcc_h6;
          let _338_v171;
          _338_v171 = _dafny.Seq.of(_156_v34, _156_v34, _156_v34);
          _156_v34 = _dafny.Seq.Concat(_156_v34, (_338_v171)[_module.__default.safeIndex(_337_cf45, new BigNumber((_338_v171).length))]);
          let _339_v172;
          let _nw63 = new _module.C3();
          _nw63.__ctor((_163_v41).f29, _336_cf46, (_163_v41).f30, _337_cf45);
          _339_v172 = _nw63;
          _99_v1 = _dafny.Seq.UnicodeFromString("kh");
          let _index45 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          (_106_v7)[_index45] = (_163_v41).f30;
        } else {
          let _340___mcc_h11 = (_source3).cf44;
          let _341_cf44 = _340___mcc_h11;
          let _342_v173;
          _342_v173 = _dafny.MultiSet.fromElements(!((_163_v41).f30));
          let _index46 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _index47 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _rhs42 = (_100_v2) && ((_163_v41).f30);
          let _rhs43 = _module.__default.safeModuloInt((_164_v42).fm6(true, true, (_dafny.ZERO).minus((_164_v42).f31), _105_globalState), (_157_v35)[_module.__default.safeIndex(new BigNumber((_99_v1).length), new BigNumber((_157_v35).length))]);
          let _rhs44 = !(_100_v2);
          let _rhs45 = (_342_v173).IsDisjointFrom((_342_v173).Intersect(_342_v173));
          let _rhs46 = _166_v44;
          let _lhs39 = _106_v7;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _lhs41 = _105_globalState;
          let _lhs42 = _106_v7;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _lhs44 = _105_globalState;
          _lhs39[_lhs40] = _rhs42;
          _lhs41.f0 = _rhs43;
          _lhs42[_lhs43] = _rhs44;
          _lhs44.f1 = _rhs45;
          _166_v44 = _rhs46;
          let _343_v174;
          let _nw64 = Array((new BigNumber(4)).toNumber()).fill([]);
          _343_v174 = _nw64;
          let _344_v175;
          _344_v175 = _dafny.Map.Empty.slice().updateUnsafe(_106_v7,_343_v174);
          let _index48 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
          let _index49 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          let _rhs47 = (_164_v42).f31;
          let _rhs48 = ((_164_v42).f31).isLessThanOrEqualTo(_101_v3);
          let _rhs49 = (((_344_v175).contains(_106_v7)) ? ((_344_v175).get(_106_v7)) : (_343_v174));
          let _rhs50 = _module.__default.fm21(_module.__default.fm1(_105_globalState), (_164_v42).f31, new BigNumber(82), _105_globalState);
          let _rhs51 = (_163_v41).f30;
          let _lhs45 = _323_v166;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
          let _lhs47 = _106_v7;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
          _lhs45[_lhs46] = _rhs47;
          _lhs47[_lhs48] = _rhs48;
          _343_v174 = _rhs49;
          _324_v167 = _rhs50;
          _100_v2 = _rhs51;
          let _345_v176;
          let _out13;
          _out13 = (_164_v42).m1(_324_v167, ((_164_v42).f31).isLessThanOrEqualTo(new BigNumber((_342_v173).cardinality())), (_163_v41).f30, _105_globalState);
          _345_v176 = _out13;
          _100_v2 = (_163_v41).f30;
        }
        let _346_v177;
        let _nw65 = new _module.C1();
        _nw65.__ctor(false);
        _346_v177 = _nw65;
        let _347_v178;
        _347_v178 = _dafny.Map.Empty.slice().updateUnsafe(_346_v177,(((_346_v177).f32) ? (_163_v41) : (_163_v41)));
        let _348_v179;
        _348_v179 = _dafny.Seq.of(_347_v178);
        let _349_v180;
        let _nw66 = new _module.C0();
        _nw66.__ctor((_163_v41).f29, (_163_v41).f30, (_dafny.ZERO).minus(new BigNumber(((_163_v41).f29).length)), _106_v7, (_163_v41).f30);
        _349_v180 = _nw66;
        let _rhs52 = (_348_v179)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_349_v180, _349_v180, _349_v180, _349_v180)).length), new BigNumber((_348_v179).length))];
        let _rhs53 = _106_v7;
        let _rhs54 = (_164_v42).fm6(!(((_164_v42).f31).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("jhjne")).length))), (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], _module.__default.safeModuloInt((((_324_v167).contains((_164_v42).f31)) ? ((_324_v167).get((_164_v42).f31)) : ((_323_v166)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length))])), (_323_v166)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length))]), _105_globalState);
        let _rhs55 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber(283), new BigNumber(161), new BigNumber(353))).length), (_164_v42).f31);
        let _lhs49 = _105_globalState;
        let _lhs50 = _105_globalState;
        _347_v178 = _rhs52;
        _106_v7 = _rhs53;
        _lhs49.f2 = _rhs54;
        _lhs50.f0 = _rhs55;
        let _index50 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
        let _rhs56 = !(_dafny.MultiSet.fromElements(_349_v180.f34, (_163_v41).f30)).contains(_100_v2);
        let _rhs57 = (_346_v177).f32;
        let _rhs58 = _323_v166;
        let _rhs59 = (_164_v42).f31;
        let _rhs60 = (_163_v41).f29;
        let _lhs51 = _349_v180;
        let _lhs52 = _106_v7;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
        _lhs51.f34 = _rhs56;
        _lhs52[_lhs53] = _rhs57;
        _323_v166 = _rhs58;
        _101_v3 = _rhs59;
        _99_v1 = _rhs60;
        let _350_v181;
        _350_v181 = _dafny.MultiSet.fromElements((_163_v41).f30, _100_v2);
        if (!(((((_350_v181).contains(_100_v2)) ? ((_350_v181).get(_100_v2)) : (_101_v3))).isEqualTo((_164_v42).f31))) {
          let _351_v182;
          let _nw67 = new _module.C4();
          _nw67.__ctor((_164_v42).f31, _99_v1, (_346_v177).f32, (_164_v42).f31);
          _351_v182 = _nw67;
          let _352_v183;
          _352_v183 = _dafny.Map.Empty.slice().updateUnsafe(_351_v182,(_164_v42).f31);
          let _353_v184;
          _353_v184 = _dafny.Map.Empty.slice().updateUnsafe((_163_v41).f30,(((_326_v169).contains((_351_v182).f27)) ? ((_326_v169).get((_351_v182).f27)) : ((_163_v41).f30)));
          let _354_v185;
          _354_v185 = _module.D14.create_DC32((_163_v41).f30, _101_v3, _353_v184, (_163_v41).f29, (_346_v177).f32);
          let _355_v186;
          let _nw68 = new _module.C4();
          _nw68.__ctor((((_352_v183).contains(_351_v182)) ? ((_352_v183).get(_351_v182)) : ((_164_v42).f31)), (_354_v185).dtor_cf55, (_346_v177).f32, new BigNumber((_102_v4).length));
          _355_v186 = _nw68;
          let _356_v187;
          let _nw69 = Array((new BigNumber(27)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _323_v166;
          _nw69[(_dafny.ONE).toNumber()] = _323_v166;
          _nw69[(new BigNumber(2)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(3)).toNumber()] = (((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) ? (_323_v166) : (_323_v166));
          _nw69[(new BigNumber(4)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(5)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(6)).toNumber()] = ((true) ? (_323_v166) : (_323_v166));
          _nw69[(new BigNumber(7)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(8)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(9)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(10)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(11)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(12)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(13)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(14)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(15)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(16)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(17)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(18)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(19)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(20)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(21)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(22)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(23)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(24)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(25)).toNumber()] = _323_v166;
          _nw69[(new BigNumber(26)).toNumber()] = _323_v166;
          _356_v187 = _nw69;
          let _index51 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
          let _rhs61 = (_156_v34)[_module.__default.safeIndex((_164_v42).f31, new BigNumber((_156_v34).length))];
          let _rhs62 = _101_v3;
          let _rhs63 = _356_v187;
          let _rhs64 = new BigNumber((_dafny.Seq.update(_157_v35, _module.__default.safeIndex((_323_v166)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length))], new BigNumber((_157_v35).length)), (_164_v42).f31)).length);
          let _lhs54 = _349_v180;
          let _lhs55 = _323_v166;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
          let _lhs57 = _105_globalState;
          _lhs54.f34 = _rhs61;
          _lhs55[_lhs56] = _rhs62;
          _356_v187 = _rhs63;
          _lhs57.f0 = _rhs64;
          _102_v4 = (_102_v4).update((_163_v41).f30, new BigNumber(-43));
          (_105_globalState).f2 = (_164_v42).f31;
          _104_v6 = _104_v6;
        } else {
          let _357_v188;
          _357_v188 = _dafny.Map.Empty.slice().updateUnsafe(_158_v36,_106_v7);
          let _358_v189;
          _358_v189 = _dafny.MultiSet.fromElements((_357_v188).Merge(_dafny.Map.Empty.slice().updateUnsafe(_158_v36,(_349_v180).f33)));
          (_105_globalState).f0 = (((_358_v189).contains(_357_v188)) ? ((_358_v189).get(_357_v188)) : ((_164_v42).f31));
          _99_v1 = _dafny.Seq.Concat((_163_v41).f29, _dafny.Seq.Concat((_163_v41).f29, (_163_v41).f29));
          (_105_globalState).f0 = (_164_v42).f31;
          let _index52 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
          (_323_v166)[_index52] = _module.__default.safeModuloInt(((_164_v42).f31).multipliedBy((_164_v42).f31), _101_v3);
          let _index53 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length));
          (_323_v166)[_index53] = (_323_v166)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_323_v166).length))];
        }
      } else {
        _157_v35 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm3((_164_v42).f31, (_163_v41).f30, (_164_v42).f31, _100_v2, _105_globalState), _157_v35), _157_v35);
        let _359_v190;
        _359_v190 = _dafny.Map.Empty.slice().updateUnsafe((_156_v34)[_module.__default.safeIndex(new BigNumber((_156_v34).length), new BigNumber((_156_v34).length))],true);
        let _360_v191;
        _360_v191 = _dafny.MultiSet.fromElements(_359_v190, _359_v190, (_359_v190).Merge(_359_v190), (_359_v190).update(!((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]), true), (_359_v190).update((_163_v41).f30, (_163_v41).f30));
        _360_v191 = ((_dafny.MultiSet.fromElements(_359_v190)).Difference(_dafny.MultiSet.fromElements(_359_v190, _359_v190, _359_v190))).Difference(_dafny.MultiSet.fromElements(_359_v190));
        let _361_v192;
        _361_v192 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("bqwbfped"), (_163_v41).f29, _99_v1, _99_v1);
        _361_v192 = _dafny.Seq.Concat(_361_v192, _361_v192);
        let _362_v193;
        _362_v193 = _dafny.Map.Empty.slice().updateUnsafe(_157_v35,_106_v7);
        _362_v193 = _362_v193;
        let _363_v194;
        let _nw70 = Array((new BigNumber(16)).toNumber()).fill([]);
        _363_v194 = _nw70;
        let _364_v195;
        let _nw71 = Array((new BigNumber(2)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = new BigNumber((_359_v190).length);
        _nw71[(_dafny.ONE).toNumber()] = _101_v3;
        _364_v195 = _nw71;
        let _index54 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_363_v194).length));
        (_363_v194)[_index54] = _364_v195;
        let _365_v196;
        _365_v196 = _dafny.Map.Empty.slice().updateUnsafe(((_164_v42).f31).plus(_101_v3),(_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]);
        let _366_v197;
        _366_v197 = _dafny.Seq.of(_156_v34);
        let _index55 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_363_v194).length));
        let _rhs65 = _364_v195;
        let _rhs66 = (((((_163_v41).f30) ? ((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))]) : ((_163_v41).f30))) ? (_364_v195) : (_364_v195));
        let _rhs67 = _module.__default.fm42(_dafny.Seq.Concat(_module.__default.fm3(new BigNumber((_dafny.Seq.of((_164_v42).f31)).length), (_163_v41).f30, (_164_v42).f31, (_163_v41).f30, _105_globalState), _dafny.Seq.update(_157_v35, _module.__default.safeIndex((_164_v42).f31, new BigNumber((_157_v35).length)), (_164_v42).f31)), _100_v2, _366_v197, _105_globalState);
        let _lhs58 = _363_v194;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_363_v194).length));
        _lhs58[_lhs59] = _rhs65;
        _364_v195 = _rhs66;
        _365_v196 = _rhs67;
      }
      let _367_v198;
      _367_v198 = _dafny.Map.Empty.slice().updateUnsafe(_99_v1,(_163_v41).f29);
      let _368_v199;
      _368_v199 = _dafny.MultiSet.fromElements((_164_v42).fm11(_367_v198, new BigNumber(((_163_v41).f29).length), (_164_v42).f31, _105_globalState), _101_v3, (_164_v42).f31, (_164_v42).f31);
      _368_v199 = (_368_v199).Union(_368_v199);
      let _369_v200;
      _369_v200 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(false),_101_v3);
      let _index56 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length));
      (_106_v7)[_index56] = !(_module.__default.fm5((_dafny.MultiSet.fromElements((_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))], !((_163_v41).f30), (_163_v41).f30, (_106_v7)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_106_v7).length))])).IsDisjointFrom(_dafny.MultiSet.fromElements((_163_v41).f30)), _369_v200, _101_v3, _105_globalState));
      let _370_v201;
      _370_v201 = _dafny.Map.Empty.slice().updateUnsafe(_106_v7,_dafny.Seq.Concat(_156_v34, _156_v34));
      _370_v201 = (_370_v201).update(_106_v7, _dafny.Seq.Concat(_dafny.Seq.of((_163_v41).f30), _156_v34));
      let _371_v202;
      let _init8 = ((_372_v3) => function (_373_i21) {
        return _module.__default.safeModuloInt(_373_i21, _372_v3);
      })(_101_v3);
      let _nw72 = Array((new BigNumber(21)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw72.length); _i0_8++) {
        _nw72[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _371_v202 = _nw72;
      let _index57 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_371_v202).length));
      (_371_v202)[_index57] = new BigNumber(923);
      let _index58 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_371_v202).length));
      let _rhs68 = _101_v3;
      let _rhs69 = (_99_v1)[_module.__default.safeIndex(_module.__default.safeModuloInt((_164_v42).f31, (_164_v42).f31), new BigNumber((_99_v1).length))];
      let _rhs70 = (_101_v3).isLessThan(new BigNumber(224));
      let _rhs71 = _156_v34;
      let _rhs72 = (_164_v42).f31;
      let _lhs60 = _371_v202;
      let _lhs61 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_371_v202).length));
      _101_v3 = _rhs68;
      _158_v36 = _rhs69;
      _100_v2 = _rhs70;
      _156_v34 = _rhs71;
      _lhs60[_lhs61] = _rhs72;
      let _374_v203;
      let _nw73 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _374_v203 = _nw73;
      let _index59 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_374_v203).length));
      (_374_v203)[_index59] = _dafny.Seq.Concat((_163_v41).f29, _dafny.Seq.UnicodeFromString("cd"));
      let _index60 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_374_v203).length));
      (_374_v203)[_index60] = _99_v1;
      process.stdout.write((_99_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_100_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_101_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-43)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_103_v5, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(189)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(189))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v6).equals(_dafny.Set.fromElements(new BigNumber(189)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_105_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_105_globalState).f9, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(189)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(189))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_105_globalState).f12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f13).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("jxiyve")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f17).equals(_dafny.Set.fromElements(new BigNumber(189)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v7)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_110_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_156_v34, _dafny.Seq.of(true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_157_v35, _dafny.Seq.of(new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_158_v36));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v38).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_160_v39, _dafny.Seq.of(_module.D0.create_DC0(true), _module.D0.create_DC0(false), _module.D0.create_DC0(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v40).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(true),_dafny.ONE).updateUnsafe(_module.D0.create_DC0(false),_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_163_v41).f29).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v41).f30));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v42).f31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_165_v43)[new BigNumber(10)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_167_v45).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v118).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('r'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v165).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_367_v198).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("oduxsblwp"),_dafny.Seq.UnicodeFromString("oduxsblwp")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_368_v199).equals(_dafny.MultiSet.fromElements(new BigNumber(991), new BigNumber(991), new BigNumber(-189), new BigNumber(-189), new BigNumber(-189), new BigNumber(-189), new BigNumber(-189), new BigNumber(-189)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_369_v200).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(false),new BigNumber(-189)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_370_v201).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v202)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_374_v203)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4, cf5) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2() {
      let $dt = new D0(1);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + this.cf2.toVerbatimString(true) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO, false);
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
    static create_DC4(cf7, cf8) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D1(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, false);
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
    static create_DC7(cf11, cf12, cf13) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, false, _dafny.ZERO);
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
    static create_DC8(cf14) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
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
      return _dafny.MultiSet.Empty;
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
    static create_DC10(cf16, cf17) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(new _dafny.CodePoint('D'.codePointAt(0)), new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC11(cf18) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf20, cf21, cf22) {
      let $dt = new D6(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf23, cf24, cf25) {
      let $dt = new D6(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D6(2);
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
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC13(_dafny.ZERO, _dafny.ZERO, _dafny.Set.Empty);
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
    static create_DC15(cf26) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf28, cf29) {
      let $dt = new D8(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC16(cf27) {
      let $dt = new D8(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf27 === other.cf27;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC17(_module.D0.Default(), false);
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
    static create_DC19(cf31) {
      let $dt = new D9(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC18(cf30) {
      let $dt = new D9(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC20(cf32) {
      let $dt = new D9(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC19(_module.D6.Default());
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
    static create_DC22(cf34, cf35) {
      let $dt = new D10(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC23() {
      let $dt = new D10(1);
      return $dt;
    }
    static create_DC24(cf36, cf37, cf38) {
      let $dt = new D10(2);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D10(3);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf34) + ", " + this.cf35.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC23";
      } else if (this.$tag === 2) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 3) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf33 === other.cf33;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22([], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC26(cf40, cf41, cf42) {
      let $dt = new D11(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC25(cf39) {
      let $dt = new D11(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC27(cf43) {
      let $dt = new D11(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC26(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC29(cf45, cf46, cf47, cf48, cf49) {
      let $dt = new D12(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC28(cf44) {
      let $dt = new D12(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48) && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf44 === other.cf44;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC29(_dafny.ZERO, false, _dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC30(cf50) {
      let $dt = new D13(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf52, cf53, cf54, cf55, cf56) {
      let $dt = new D14(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC31(cf51) {
      let $dt = new D14(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + this.cf55.toVerbatimString(true) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54) && _dafny.areEqual(this.cf55, other.cf55) && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC32(false, _dafny.ZERO, _dafny.Map.Empty, _dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC34(cf58) {
      let $dt = new D15(0);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC33(cf57) {
      let $dt = new D15(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf57 === other.cf57;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC34(_dafny.ZERO);
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
    static create_DC36(cf60, cf61, cf62) {
      let $dt = new D16(0);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC35(cf59) {
      let $dt = new D16(1);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC36" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf60 === other.cf60 && _dafny.areEqual(this.cf61, other.cf61) && this.cf62 === other.cf62;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC36(false, _dafny.Seq.of(), []);
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
    static create_DC38() {
      let $dt = new D17(0);
      return $dt;
    }
    static create_DC37(cf63) {
      let $dt = new D17(1);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC39(cf64) {
      let $dt = new D17(2);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC38";
      } else if (this.$tag === 1) {
        return "D17.DC37" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC39" + "(" + _dafny.toString(this.cf64) + ")";
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
        return other.$tag === 1 && this.cf63 === other.cf63;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC38();
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
    static create_DC40(cf65) {
      let $dt = new D18(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC40" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65);
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
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf67, cf68, cf69) {
      let $dt = new D19(0);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC43(cf70, cf71, cf72) {
      let $dt = new D19(1);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC41(cf66) {
      let $dt = new D19(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC42" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC43" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC41" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67) && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71) && this.cf72 === other.cf72;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf66 === other.cf66;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC42(_dafny.MultiSet.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC45(cf74) {
      let $dt = new D20(0);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC46(cf75, cf76) {
      let $dt = new D20(1);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC47() {
      let $dt = new D20(2);
      return $dt;
    }
    static create_DC44(cf73) {
      let $dt = new D20(3);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get is_DC44() { return this.$tag === 3; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC45" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC46" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC47";
      } else if (this.$tag === 3) {
        return "D20.DC44" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf75 === other.cf75 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC45(_dafny.ZERO);
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
    static create_DC49(cf78, cf79) {
      let $dt = new D21(0);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC48(cf77) {
      let $dt = new D21(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC49" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC48" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf78 === other.cf78 && this.cf79 === other.cf79;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf77 === other.cf77;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC49(false, []);
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
    static create_DC51(cf81, cf82) {
      let $dt = new D22(0);
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC50(cf80) {
      let $dt = new D22(1);
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC52(cf83) {
      let $dt = new D22(2);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC51" + "(" + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC50" + "(" + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC52" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf81, other.cf81) && this.cf82 === other.cf82;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC51(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
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
      this.f0 = _dafny.ZERO;
      this.f1 = false;
      this.f2 = _dafny.ZERO;
      this._f3 = [];
      this._f4 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = _dafny.Seq.of();
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = _dafny.Seq.UnicodeFromString("");
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.ZERO;
      this._f17 = _dafny.Set.Empty;
      this._f18 = false;
      this._f19 = _dafny.ZERO;
      this._f20 = false;
      this._f21 = false;
      this._f22 = false;
      this._f23 = false;
      this._f24 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
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
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f34 = false;
      this._f25 = false;
      this._f26 = _dafny.ZERO;
      this._f33 = [];
      this.f35 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0, _module.T1, _module.T2];
    }
    get f34() {
      let _this = this;
      return _this._f34;
    };
    set f34(value) {
      let _this = this;
      _this._f34 = value;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    __ctor(f35, f25, f26, f33, f34) {
      let _this = this;
      (_this).f35 = f35;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this)._f33 = f33;
      (_this)._f34 = f34;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f25, !(_this.f34)))).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f25, _this.f34)))) {
        return (_this).f26;
      } else {
        return (_this).f26;
      }
    };
    fm7(globalState) {
      let _this = this;
      return (_this).f25;
    };
    fm16(globalState) {
      let _this = this;
      return (new BigNumber(162)).plus((_this).f26);
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      return !((_this).f26).isEqualTo(((_this).f26).minus((_this).f26));
    };
    fm18(p0, p1, p2, globalState) {
      let _this = this;
      return function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of ((_dafny.MultiSet.fromElements(_this.f35, _this.f35, _this.f35, _this.f35, _this.f35)).Difference(_dafny.MultiSet.fromElements(_this.f35, _this.f35, _this.f35))).Elements) {
          let _375_v0 = _compr_9;
          if (((_dafny.MultiSet.fromElements(_this.f35, _this.f35, _this.f35, _this.f35, _this.f35)).Difference(_dafny.MultiSet.fromElements(_this.f35, _this.f35, _this.f35))).contains(_375_v0)) {
            _coll9.push([_375_v0,(_this).f25]);
          }
        }
        return _coll9;
      }();
    };
    fm19(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_this.f35).Elements) {
          let _376_v0 = _compr_10;
          if (_dafny.Seq.contains(_this.f35, _376_v0)) {
            _coll10.push([_376_v0,_this.f34]);
          }
        }
        return _coll10;
      }()).length))).Intersect((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f26))))).cardinality());
    };
    fm20(p0, globalState) {
      let _this = this;
      let _source4 = _module.D2.create_DC7((_this).f25, true, (_this).f26);
      if (_source4.is_DC7) {
        let _377___mcc_h0 = (_source4).cf11;
        let _378___mcc_h1 = (_source4).cf12;
        let _379___mcc_h2 = (_source4).cf13;
        let _380_cf13 = _379___mcc_h2;
        let _381_cf12 = _378___mcc_h1;
        let _382_cf11 = _377___mcc_h0;
        return new _dafny.CodePoint('m'.codePointAt(0));
      } else {
        let _383___mcc_h3 = (_source4).cf10;
        let _384_cf10 = _383___mcc_h3;
        return new _dafny.CodePoint('w'.codePointAt(0));
      }
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f2 = (_this).f26;
      let _385_v0;
      _385_v0 = _dafny.MultiSet.fromElements(_this.f35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(449)), function (_386_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }));
      (globalState).f2 = (((_385_v0).contains(_dafny.Seq.Concat(_this.f35, _dafny.Seq.UnicodeFromString("hfhbn")))) ? ((_385_v0).get(_dafny.Seq.Concat(_this.f35, _dafny.Seq.UnicodeFromString("hfhbn")))) : (((p1) ? ((_this).f26) : (new BigNumber((p0).cardinality())))));
      let _387_v1;
      let _nw74 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
      _387_v1 = _nw74;
      let _388_v2;
      _388_v2 = _dafny.Seq.of((_this).f33);
      let _index61 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_387_v1).length));
      (_387_v1)[_index61] = _388_v2;
      let _389_v3;
      let _nw75 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _389_v3 = _nw75;
      let _index62 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_389_v3).length));
      (_389_v3)[_index62] = _this.f35;
      let _390_v4;
      let _nw76 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _390_v4 = _nw76;
      let _391_v5;
      _391_v5 = _dafny.Map.Empty.slice().updateUnsafe(_390_v4,_module.__default.safeModuloInt((_this).f26, new BigNumber(574)));
      let _392_v6;
      _392_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Seq.Concat(_388_v2, _dafny.Seq.update(_dafny.Seq.update(_388_v2, _module.__default.safeIndex(new BigNumber(396), new BigNumber((_388_v2).length)), (_this).f33), _module.__default.safeIndex(new BigNumber((_this.f35).length), new BigNumber((_dafny.Seq.update(_388_v2, _module.__default.safeIndex(new BigNumber(396), new BigNumber((_388_v2).length)), (_this).f33)).length)), (_this).f33)));
      let _393_v7;
      _393_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
      let _394_v8;
      _394_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_393_v7).length),p1);
      let _index63 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_387_v1).length));
      let _index64 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_389_v3).length));
      let _rhs73 = (((_392_v6).contains((((_394_v8).contains((_this).f26)) ? ((_394_v8).get((_this).f26)) : (_this.f34)))) ? ((_392_v6).get((((_394_v8).contains((_this).f26)) ? ((_394_v8).get((_this).f26)) : (_this.f34)))) : (_dafny.Seq.Concat(_388_v2, _dafny.Seq.of((_this).f33, (_this).f33))));
      let _rhs74 = (_this).f25;
      let _rhs75 = _this.f35;
      let _rhs76 = _391_v5;
      let _lhs62 = _387_v1;
      let _lhs63 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_387_v1).length));
      let _lhs64 = globalState;
      let _lhs65 = _389_v3;
      let _lhs66 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_389_v3).length));
      _lhs62[_lhs63] = _rhs73;
      _lhs64.f1 = _rhs74;
      _lhs65[_lhs66] = _rhs75;
      _391_v5 = _rhs76;
      let _index65 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_390_v4).length));
      (_390_v4)[_index65] = _module.__default.safeModuloInt((_this).f26, new BigNumber(((_389_v3)[_module.__default.safeIndex(new BigNumber(289), new BigNumber((_389_v3).length))]).length));
      let _index66 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_390_v4).length));
      let _rhs77 = p1;
      let _rhs78 = (_dafny.ZERO).minus((_this).f26);
      let _rhs79 = p2;
      let _lhs67 = _390_v4;
      let _lhs68 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_390_v4).length));
      r0 = _rhs77;
      _lhs67[_lhs68] = _rhs78;
      r0 = _rhs79;
      (globalState).f2 = _module.__default.safeDivisionInt((_this).f26, ((_390_v4)[_module.__default.safeIndex(new BigNumber(197), new BigNumber((_390_v4).length))]).multipliedBy((_390_v4)[_module.__default.safeIndex(new BigNumber(197), new BigNumber((_390_v4).length))]));
      let _395_v9;
      _395_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f34,(_this).f26);
      (globalState).f2 = ((!(_395_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(_this.f34,new BigNumber(((_389_v3)[_module.__default.safeIndex(new BigNumber(289), new BigNumber((_389_v3).length))]).length)))) ? ((_this).f26) : ((new BigNumber((_dafny.Seq.of(!(_this.f34))).length)).minus((_390_v4)[_module.__default.safeIndex(new BigNumber(197), new BigNumber((_390_v4).length))])));
      r0 = p2;
      return r0;
    }
    m9(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _396_i0;
      _396_i0 = _dafny.ZERO;
      L2: {
        while ((true) && ((_this).f25)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_396_i0)) {
              break L2;
            }
            _396_i0 = (_396_i0).plus(_dafny.ONE);
            (globalState).f1 = ((_this).f26).isEqualTo((_this).f26);
            let _index67 = _module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length));
            ((_this).f33)[_index67] = _this.f34;
            let _index68 = _module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length));
            let _rhs80 = _module.__default.safeModuloInt(new BigNumber(-188), (_dafny.ZERO).minus((_this).f26));
            let _rhs81 = _this.f34;
            let _rhs82 = _this.f34;
            let _rhs83 = ((((_this).f26).isLessThanOrEqualTo((_this).f26)) ? ((_this).f26) : ((_this).f26));
            let _lhs69 = globalState;
            let _lhs70 = (_this).f33;
            let _lhs71 = _module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length));
            let _lhs72 = globalState;
            let _lhs73 = globalState;
            _lhs69.f0 = _rhs80;
            _lhs70[_lhs71] = _rhs81;
            _lhs72.f1 = _rhs82;
            _lhs73.f2 = _rhs83;
            let _397_v0;
            _397_v0 = _dafny.Seq.of((_this).f26);
            (globalState).f2 = _module.__default.fm2(((_this.f34) ? ((_397_v0)[_module.__default.safeIndex((_this).fm16(globalState), new BigNumber((_397_v0).length))]) : (new BigNumber(-33))), (_dafny.ZERO).minus(new BigNumber((_397_v0).length)), globalState);
            let _398_v1;
            _398_v1 = _dafny.Set.fromElements(((_this).f33)[_module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length))], ((_this).f33)[_module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length))]);
            let _399_v2;
            _399_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_398_v1).length));
            let _400_v3;
            _400_v3 = _module.D2.create_DC7((_this).f25, _this.f34, (((_399_v2).contains((_this).f26)) ? ((_399_v2).get((_this).f26)) : ((_397_v0)[_module.__default.safeIndex((_this).f26, new BigNumber((_397_v0).length))])));
            let _401_v4;
            let _nw77 = Array((new BigNumber(21)).toNumber());
            _nw77[(_dafny.ZERO).toNumber()] = _400_v3;
            _nw77[(_dafny.ONE).toNumber()] = _400_v3;
            _nw77[(new BigNumber(2)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(3)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(4)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(5)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(6)).toNumber()] = _module.D2.create_DC7(_this.f34, ((_this).f33)[_module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length))], (_this).f26);
            _nw77[(new BigNumber(7)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(8)).toNumber()] = _module.D2.create_DC7((_this).f25, ((_this).f33)[_module.__default.safeIndex(new BigNumber(435), new BigNumber(((_this).f33).length))], new BigNumber(442));
            _nw77[(new BigNumber(9)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(10)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(11)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(12)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(13)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(14)).toNumber()] = (((_this).f25) ? (_400_v3) : (_400_v3));
            _nw77[(new BigNumber(15)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(16)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(17)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(18)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(19)).toNumber()] = _400_v3;
            _nw77[(new BigNumber(20)).toNumber()] = _400_v3;
            _401_v4 = _nw77;
            let _index69 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_401_v4).length));
            (_401_v4)[_index69] = _400_v3;
            let _index70 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_401_v4).length));
            (_401_v4)[_index70] = _400_v3;
          }
        }
      }
      let _402_v5;
      _402_v5 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f26, (_this).f26),p0);
      _402_v5 = (_402_v5).update(new BigNumber(((p0).Intersect(p0)).cardinality()), (p0).update((_this).f25, _module.__default.abs((_this).f26)));
      (_this).f34 = _this.f34;
      (_this).f34 = _this.f34;
      let _403_v6;
      _403_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
      let _404_v7;
      _404_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_403_v6);
      let _hi2 = _module.__default.safeModuloInt((((p0).contains((_this).f25)) ? ((p0).get((_this).f25)) : (new BigNumber((_404_v7).length))), (_this).f26);
      for (let _405_i1 = (_dafny.ZERO).minus(((_this).fm6((_this).f25, (_this).f25, (_this).f26, globalState)).plus((_this).f26)); _405_i1.isLessThan(_hi2); _405_i1 = _405_i1.plus(_dafny.ONE)) {
        (globalState).f2 = new BigNumber(480);
        (globalState).f1 = (_405_i1).isLessThanOrEqualTo((_this).f26);
        (globalState).f1 = (_this).f25;
        (globalState).f1 = true;
      }
      (globalState).f0 = (_this).f26;
      let _406_v8;
      _406_v8 = new _dafny.CodePoint('g'.codePointAt(0));
      r0 = _406_v8;
      r1 = (_this).f33;
      return [r0, r1];
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f32 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f32) {
      let _this = this;
      (_this)._f32 = f32;
      return;
    }
    fm14(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ehss"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(556)), function (_407_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("kbxgttvgb"));
    };
    fm15(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(967)), function (_408_i0) {
        return new BigNumber(644);
      })).length)).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(727)), function (_409_i1) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length), new BigNumber(-971), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-428)), function (_410_i2) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(907)), function (_411_i3) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length), new BigNumber(943))).length));
    };
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = _dafny.Map.Empty;
      let _412_v0;
      let _init9 = function (_413_i0) {
        return (_this).f32;
      };
      let _nw78 = Array((new BigNumber(7)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw78.length); _i0_9++) {
        _nw78[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _412_v0 = _nw78;
      let _414_v1;
      _414_v1 = new _dafny.CodePoint('o'.codePointAt(0));
      let _415_v2;
      _415_v2 = _dafny.Map.Empty.slice().updateUnsafe(_414_v1,(_this).f32);
      let _index71 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
      (_412_v0)[_index71] = (((_415_v2).contains(_414_v1)) ? ((_415_v2).get(_414_v1)) : ((_this).f32));
      let _index72 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
      (_412_v0)[_index72] = (_this).f32;
      let _416_v3;
      _416_v3 = _dafny.Seq.UnicodeFromString("qby");
      _416_v3 = p2;
      let _417_v4;
      _417_v4 = new BigNumber(238);
      let _418_v5;
      _418_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,_417_v4);
      if (_module.__default.fm5(!((_this).f32), _418_v5, _417_v4, globalState)) {
        let _419_v6;
        _419_v6 = _dafny.Map.Empty.slice().updateUnsafe(_417_v4,p0);
        let _index73 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
        (_412_v0)[_index73] = ((_419_v6).update(_417_v4, (_this).f32)).equals((_419_v6).Merge(_419_v6));
        let _index74 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
        let _rhs84 = _417_v4;
        let _rhs85 = !((_module.D0.create_DC1(_417_v4, p2, (_this).f32, (_dafny.ZERO).minus(_417_v4), p0)).dtor_cf5);
        let _lhs74 = globalState;
        let _lhs75 = _412_v0;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
        _lhs74.f2 = _rhs84;
        _lhs75[_lhs76] = _rhs85;
        (globalState).f1 = (_this).f32;
        let _420_v7;
        _420_v7 = _dafny.Set.fromElements(false);
        let _421_v8;
        _421_v8 = _module.D4.create_DC9(_420_v7);
        let _rhs86 = _dafny.Seq.Concat(_416_v3, _416_v3);
        let _rhs87 = _module.__default.fm2(_417_v4, new BigNumber((((_421_v8).dtor_cf15).Difference(_420_v7)).length), globalState);
        let _rhs88 = (_dafny.ZERO).minus((((_this).f32) ? ((_dafny.ZERO).minus(_417_v4)) : (new BigNumber(((_420_v7).Difference(_dafny.Set.fromElements((_this).f32))).length))));
        let _lhs77 = globalState;
        let _lhs78 = globalState;
        _416_v3 = _rhs86;
        _lhs77.f2 = _rhs87;
        _lhs78.f0 = _rhs88;
        (globalState).f0 = _417_v4;
      } else {
        let _422_v9;
        _422_v9 = _dafny.Map.Empty.slice().updateUnsafe((_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))],new BigNumber(347));
        let _423_v10;
        _423_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f32,_417_v4), _422_v9);
        let _424_v11;
        _424_v11 = _dafny.Map.Empty.slice().updateUnsafe(_417_v4,_dafny.MultiSet.FromArray(_423_v10));
        let _425_v12;
        _425_v12 = _dafny.Set.fromElements((_this).f32, true, p0);
        let _426_v13;
        _426_v13 = _dafny.MultiSet.fromElements((_422_v9).update((_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))], _417_v4));
        let _427_v14;
        _427_v14 = (((_424_v11).contains(new BigNumber((_425_v12).length))) ? ((_424_v11).get(new BigNumber((_425_v12).length))) : (_426_v13));
        r0 = _427_v14;
        let _428_v15;
        _428_v15 = _module.D0.create_DC2();
        let _429_v16;
        _429_v16 = _dafny.Map.Empty.slice().updateUnsafe(_428_v15,(_this).f32);
        if ((_429_v16).contains(_428_v15)) {
          let _430_v17;
          let _init10 = ((_431_v4) => function (_432_i1) {
            return (_432_i1).multipliedBy(_431_v4);
          })(_417_v4);
          let _nw79 = Array((new BigNumber(29)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw79.length); _i0_10++) {
            _nw79[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _430_v17 = _nw79;
          let _rhs89 = (_417_v4).isLessThan((_417_v4).plus(_417_v4));
          let _rhs90 = _430_v17;
          let _lhs79 = globalState;
          _lhs79.f1 = _rhs89;
          _430_v17 = _rhs90;
          _416_v3 = _dafny.Seq.Concat(_416_v3, _416_v3);
          (globalState).f2 = _417_v4;
          let _433_v18;
          _433_v18 = _module.D0.create_DC1(_417_v4, _416_v3, p0, _417_v4, (_this).f32);
          let _434_v19;
          let _nw80 = new _module.C0();
          _nw80.__ctor((function (_pat_let10_0) {
            return function (_435_dt__update__tmp_h0) {
              return function (_pat_let11_0) {
                return function (_436_dt__update_hcf5_h0) {
                  return _module.D0.create_DC1((_435_dt__update__tmp_h0).dtor_cf1, (_435_dt__update__tmp_h0).dtor_cf2, (_435_dt__update__tmp_h0).dtor_cf3, (_435_dt__update__tmp_h0).dtor_cf4, _436_dt__update_hcf5_h0);
                }(_pat_let11_0);
              }((_this).f32);
            }(_pat_let10_0);
          }(_433_v18)).dtor_cf2, !_dafny.areEqual(_dafny.Seq.UnicodeFromString("yggum"), _416_v3), _417_v4, _412_v0, (_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))]);
          _434_v19 = _nw80;
          (globalState).f1 = p0;
        } else {
          (globalState).f0 = _417_v4;
          _414_v1 = _414_v1;
          let _437_v20;
          _437_v20 = _dafny.Seq.of((_this).f32);
          let _438_v21;
          let _nw81 = Array((new BigNumber(10)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = (new BigNumber(224)).minus(_417_v4);
          _nw81[(_dafny.ONE).toNumber()] = _417_v4;
          _nw81[(new BigNumber(2)).toNumber()] = new BigNumber(931);
          _nw81[(new BigNumber(3)).toNumber()] = _417_v4;
          _nw81[(new BigNumber(4)).toNumber()] = _417_v4;
          _nw81[(new BigNumber(5)).toNumber()] = (_417_v4).minus(new BigNumber(-173));
          _nw81[(new BigNumber(6)).toNumber()] = new BigNumber((p2).length);
          _nw81[(new BigNumber(7)).toNumber()] = _417_v4;
          _nw81[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_437_v20).length), _417_v4);
          _nw81[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_417_v4);
          _438_v21 = _nw81;
          let _439_v22;
          _439_v22 = _module.D0.create_DC1(new BigNumber(791), _416_v3, (_this).f32, _417_v4, true);
          let _index75 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_438_v21).length));
          (_438_v21)[_index75] = _module.__default.fm2(_417_v4, (_439_v22).dtor_cf1, globalState);
          let _440_v23;
          _440_v23 = _dafny.Map.Empty.slice().updateUnsafe(_417_v4,_417_v4);
          let _441_v24;
          _441_v24 = _dafny.Map.Empty.slice().updateUnsafe(_438_v21,_440_v23);
          let _442_v25;
          _442_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_441_v24);
          let _443_v27;
          _443_v27 = _dafny.Seq.of(_417_v4);
          let _444_v28;
          _444_v28 = _dafny.MultiSet.fromElements(new BigNumber((_443_v27).length), _417_v4);
          let _445_v29;
          _445_v29 = _dafny.Seq.of(p1);
          let _index76 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_438_v21).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
          let _rhs91 = _417_v4;
          let _rhs92 = ((((_442_v25).contains((_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))])) ? ((_442_v25).get((_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))])) : (_441_v24))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_438_v21,(function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_444_v28).Elements) {
              let _446_v26 = _compr_11;
              if ((_444_v28).contains(_446_v26)) {
                _coll11.push([(_446_v26).minus(_417_v4),_417_v4]);
              }
            }
            return _coll11;
          }()).update(new BigNumber(617), _417_v4)));
          let _rhs93 = (_417_v4).isLessThan(new BigNumber((_445_v29).length));
          let _rhs94 = (_417_v4).plus(_417_v4);
          let _lhs80 = _438_v21;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_438_v21).length));
          let _lhs82 = _412_v0;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
          let _lhs84 = globalState;
          _lhs80[_lhs81] = _rhs91;
          _441_v24 = _rhs92;
          _lhs82[_lhs83] = _rhs93;
          _lhs84.f2 = _rhs94;
          _422_v9 = (_422_v9).Merge(_422_v9);
          _416_v3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-795)), ((_447_v1) => function (_448_i2) {
            return _447_v1;
          })(_414_v1));
        }
        let _449_v30;
        _449_v30 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(_dafny.Seq.UnicodeFromString("adicjv"), _417_v4, _417_v4, globalState),(_this).f32);
        _449_v30 = _449_v30;
        _414_v1 = _414_v1;
        let _450_v31;
        _450_v31 = _dafny.Map.Empty.slice().updateUnsafe((_417_v4).plus(_417_v4),p0);
        let _451_v32;
        _451_v32 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), _414_v1);
        _450_v31 = (_450_v31).update(_417_v4, (_451_v32).IsSubsetOf(_451_v32));
      }
      let _452_v33;
      _452_v33 = _dafny.Set.fromElements((_this).f32);
      let _453_v34;
      _453_v34 = _module.D1.create_DC4((_this).fm14(_452_v33, new BigNumber(614), globalState), (_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))]);
      let _454_v35;
      _454_v35 = _dafny.MultiSet.fromElements(_453_v34);
      let _index78 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length));
      (_412_v0)[_index78] = (_454_v35).IsDisjointFrom(_454_v35);
      let _pat_let_tv9 = _417_v4;
      let _pat_let_tv10 = _417_v4;
      let _pat_let_tv11 = _417_v4;
      let _rhs95 = _module.__default.fm22((new BigNumber((_452_v33).length)).plus(_417_v4), _417_v4, _414_v1, globalState);
      let _rhs96 = true;
      let _rhs97 = _dafny.Seq.Concat(_dafny.Seq.Concat(_416_v3, _416_v3), p2);
      let _rhs98 = function (_source5) {
        if (_source5.is_DC4) {
          let _455___mcc_h0 = (_source5).cf7;
          let _456___mcc_h1 = (_source5).cf8;
          let _457_cf8 = _456___mcc_h1;
          let _458_cf7 = _455___mcc_h0;
          return new BigNumber(-400);
        } else if (_source5.is_DC3) {
          let _459___mcc_h2 = (_source5).cf6;
          let _460_cf6 = _459___mcc_h2;
          return _module.__default.safeDivisionInt(_pat_let_tv9, _pat_let_tv10);
        } else {
          let _461___mcc_h3 = (_source5).cf9;
          let _462_cf9 = _461___mcc_h3;
          return _pat_let_tv11;
        }
      }(_453_v34);
      let _rhs99 = new BigNumber(230);
      let _lhs85 = globalState;
      let _lhs86 = globalState;
      let _lhs87 = globalState;
      _453_v34 = _rhs95;
      _lhs85.f1 = _rhs96;
      _416_v3 = _rhs97;
      _lhs86.f0 = _rhs98;
      _lhs87.f0 = _rhs99;
      let _463_v36;
      _463_v36 = _dafny.Seq.of(p0, _dafny.Seq.contains(_416_v3, _414_v1), (_this).f32, (_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))]);
      _463_v36 = _module.__default.fm23(globalState);
      let _464_v37;
      _464_v37 = _dafny.Map.Empty.slice().updateUnsafe((_412_v0)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_412_v0).length))],_417_v4);
      let _465_v38;
      _465_v38 = _dafny.MultiSet.fromElements(_464_v37, _464_v37);
      let _466_v39;
      _466_v39 = _465_v38;
      r0 = _466_v39;
      let _467_v40;
      _467_v40 = _dafny.Map.Empty.slice().updateUnsafe((_417_v4).plus(_417_v4),(_this).fm14(_452_v33, new BigNumber((_463_v36).length), globalState));
      r1 = _467_v40;
      return [r0, r1];
    }
    m8(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _468_v0;
      _468_v0 = new _dafny.CodePoint('o'.codePointAt(0));
      let _469_v1;
      _469_v1 = _dafny.Map.Empty.slice().updateUnsafe(_468_v0,p1);
      (globalState).f1 = ((_this).f32) === (!((((_469_v1).contains(_468_v0)) ? ((_469_v1).get(_468_v0)) : ((_this).f32))));
      let _hi3 = p2;
      for (let _470_i0 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("vfgmwfj")).length), p2); _470_i0.isLessThan(_hi3); _470_i0 = _470_i0.plus(_dafny.ONE)) {
        let _471_v2;
        _471_v2 = _dafny.Seq.of((_this).f32, (_this).f32, p1, p1, p1);
        let _472_v3;
        _472_v3 = _dafny.Set.fromElements(p2);
        (globalState).f1 = (_471_v2)[_module.__default.safeIndex(new BigNumber(((_472_v3).Difference(_472_v3)).length), new BigNumber((_471_v2).length))];
        (globalState).f1 = !(p2).isEqualTo(_470_i0);
        r1 = (_this).f32;
        (globalState).f1 = ((_module.__default.fm24((_this).f32, globalState)).IsSubsetOf(_472_v3)) || (p1);
      }
      if (p1) {
        if (p1) {
          let _473_v4;
          _473_v4 = _dafny.Seq.UnicodeFromString("qh");
          let _474_v5;
          _474_v5 = _dafny.Map.Empty.slice().updateUnsafe((p2).minus(p2),_dafny.Seq.Concat(_473_v4, _473_v4));
          let _475_v6;
          _475_v6 = _dafny.Set.fromElements((_this).f32, p1);
          let _rhs100 = new BigNumber((_module.__default.fm1(globalState)).length);
          let _rhs101 = (p1) === (p1);
          let _rhs102 = ((p1) ? (_module.__default.safeModuloInt(p2, (_dafny.ZERO).minus(p2))) : (p2));
          let _rhs103 = (((_474_v5).contains(p2)) ? ((_474_v5).get(p2)) : (_473_v4));
          let _rhs104 = (new BigNumber((_475_v6).length)).isEqualTo(new BigNumber(113));
          let _lhs88 = globalState;
          let _lhs89 = globalState;
          let _lhs90 = globalState;
          let _lhs91 = globalState;
          _lhs88.f0 = _rhs100;
          _lhs89.f1 = _rhs101;
          _lhs90.f2 = _rhs102;
          r0 = _rhs103;
          _lhs91.f1 = _rhs104;
          (globalState).f2 = (p2).minus(p2);
          let _476_v7;
          let _nw82 = new _module.C0();
          _nw82.__ctor(_dafny.Seq.Concat(_473_v4, _dafny.Seq.UnicodeFromString("tgruw")), (_this).f32, p2, p0, !_dafny.areEqual(_473_v4, _473_v4));
          _476_v7 = _nw82;
          let _477_v8;
          _477_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _478_v9;
          _478_v9 = _dafny.MultiSet.fromElements((_this).f32, (((_477_v8).contains(true)) ? ((_477_v8).get(true)) : ((_this).f32)), (_this).f32, (_this).f32, !(!(!(p1))));
          let _rhs105 = (((_478_v9).IsDisjointFrom(_478_v9)) ? (((_this).f32) || (p1)) : (true));
          let _rhs106 = _476_v7;
          let _rhs107 = _468_v0;
          let _rhs108 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(19)), ((_479_v0) => function (_480_i1) {
            return _479_v0;
          })(_468_v0)), _dafny.Seq.UnicodeFromString("h"));
          let _lhs92 = globalState;
          let _lhs93 = globalState;
          _lhs92.f1 = _rhs105;
          _476_v7 = _rhs106;
          _468_v0 = _rhs107;
          _lhs93.f1 = _rhs108;
          (globalState).f0 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-248)), ((_481_v0) => function (_482_i2) {
            return _481_v0;
          })(_468_v0)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-248)), ((_483_v0) => function (_484_i2) {
            return _483_v0;
          })(_468_v0))).length)), (_476_v7.f35)[_module.__default.safeIndex(p2, new BigNumber((_476_v7.f35).length))]), _module.__default.safeIndex((p2).multipliedBy(p2), new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-248)), ((_485_v0) => function (_486_i2) {
            return _485_v0;
          })(_468_v0)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-248)), ((_487_v0) => function (_488_i2) {
            return _487_v0;
          })(_468_v0))).length)), (_476_v7.f35)[_module.__default.safeIndex(p2, new BigNumber((_476_v7.f35).length))])).length)), (_476_v7.f35)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(p2)), new BigNumber((_476_v7.f35).length))])).length);
          let _489_v10;
          let _nw83 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _489_v10 = _nw83;
          let _index79 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_489_v10).length));
          (_489_v10)[_index79] = p2;
          let _index80 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_489_v10).length));
          (_489_v10)[_index80] = p2;
        } else {
          let _490_v11;
          let _nw84 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _490_v11 = _nw84;
          let _491_v12;
          _491_v12 = _dafny.Set.fromElements(_490_v11);
          let _492_v13;
          _492_v13 = _dafny.MultiSet.fromElements(p2);
          let _493_v14;
          _493_v14 = _dafny.Seq.UnicodeFromString("lpv");
          let _494_v15;
          _494_v15 = _dafny.Seq.of(p1);
          let _495_v16;
          let _nw85 = Array((new BigNumber(8)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _492_v13;
          _nw85[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(p2, (_dafny.ZERO).minus(new BigNumber((_493_v14).length)), p2);
          _nw85[(new BigNumber(2)).toNumber()] = (_492_v13).Difference(_492_v13);
          _nw85[(new BigNumber(3)).toNumber()] = (_module.__default.fm21(_493_v14, new BigNumber((_494_v15).length), p2, globalState)).Union(_492_v13);
          _nw85[(new BigNumber(4)).toNumber()] = _492_v13;
          _nw85[(new BigNumber(5)).toNumber()] = (_492_v13).Intersect(_492_v13);
          _nw85[(new BigNumber(6)).toNumber()] = _492_v13;
          _nw85[(new BigNumber(7)).toNumber()] = _492_v13;
          _495_v16 = _nw85;
          let _index81 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_495_v16).length));
          (_495_v16)[_index81] = _492_v13;
          let _496_v17;
          _496_v17 = _492_v13;
          let _497_v18;
          _497_v18 = _dafny.Set.fromElements(_496_v17, _496_v17, _496_v17, _492_v13);
          let _498_v19;
          _498_v19 = _module.D6.create_DC12(_491_v12);
          let _index82 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_495_v16).length));
          let _rhs109 = (_dafny.ZERO).minus(p2);
          let _rhs110 = ((_498_v19).dtor_cf19).Difference(_491_v12);
          let _rhs111 = _492_v13;
          let _rhs112 = ((_497_v18).Union(_497_v18)).Intersect(_497_v18);
          let _lhs94 = globalState;
          let _lhs95 = _495_v16;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_495_v16).length));
          _lhs94.f2 = _rhs109;
          _491_v12 = _rhs110;
          _lhs95[_lhs96] = _rhs111;
          _497_v18 = _rhs112;
          let _499_v20;
          let _init11 = ((_500_v14) => function (_501_i3) {
            return _500_v14;
          })(_493_v14);
          let _nw86 = Array((new BigNumber(24)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw86.length); _i0_11++) {
            _nw86[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _499_v20 = _nw86;
          _499_v20 = _499_v20;
          (globalState).f0 = (new BigNumber(553)).plus(new BigNumber(740));
          (globalState).f0 = p2;
          (globalState).f1 = (_this).f32;
        }
        let _502_v21;
        _502_v21 = _module.D4.create_DC10(_468_v0, new _dafny.CodePoint('k'.codePointAt(0)));
        let _503_v22;
        _503_v22 = _module.D1.create_DC3(_468_v0);
        let _pat_let_tv12 = _503_v22;
        _502_v21 = function (_pat_let12_0) {
          return function (_504_dt__update__tmp_h1) {
            return function (_pat_let13_0) {
              return function (_505_dt__update_hcf16_h0) {
                return _module.D4.create_DC10(_505_dt__update_hcf16_h0, (_504_dt__update__tmp_h1).dtor_cf17);
              }(_pat_let13_0);
            }((_pat_let_tv12).dtor_cf6);
          }(_pat_let12_0);
        }(_502_v21);
        let _506_v23;
        _506_v23 = _dafny.Seq.UnicodeFromString("u");
        let _507_v24;
        let _nw87 = new _module.C0();
        _nw87.__ctor(_506_v23, ((_this).f32) && (p1), new BigNumber((_dafny.Seq.UnicodeFromString("norpc")).length), p0, p1);
        _507_v24 = _nw87;
        (globalState).f0 = (_dafny.ZERO).minus(p2);
        let _index83 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
        (p0)[_index83] = (p2).isLessThanOrEqualTo(p2);
        let _508_v25;
        let _nw88 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _508_v25 = _nw88;
        let _index84 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
        let _rhs113 = (p1) === (!(p2).isEqualTo(new BigNumber((_506_v23).length)));
        let _rhs114 = _508_v25;
        let _lhs97 = p0;
        let _lhs98 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
        _lhs97[_lhs98] = _rhs113;
        _508_v25 = _rhs114;
      } else {
        let _509_v26;
        let _nw89 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _509_v26 = _nw89;
        let _index85 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length));
        (_509_v26)[_index85] = p2;
        let _510_v27;
        _510_v27 = _dafny.Seq.UnicodeFromString("scemdf");
        let _index86 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length));
        let _rhs115 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_510_v27, _dafny.Seq.Concat(_510_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(195)), ((_511_v0) => function (_512_i4) {
          return _511_v0;
        })(_468_v0))))).length));
        let _rhs116 = p2;
        let _lhs99 = _509_v26;
        let _lhs100 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length));
        let _lhs101 = globalState;
        _lhs99[_lhs100] = _rhs115;
        _lhs101.f0 = _rhs116;
        (globalState).f2 = p2;
        if ((_module.__default.safeModuloInt(p2, new BigNumber(660))).isLessThanOrEqualTo((_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))])) {
          r1 = false;
          let _513_v28;
          let _init12 = ((_514_v0) => function (_515_i5) {
            return _514_v0;
          })(_468_v0);
          let _nw90 = Array((new BigNumber(5)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw90.length); _i0_12++) {
            _nw90[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _513_v28 = _nw90;
          let _index87 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_513_v28).length));
          (_513_v28)[_index87] = new _dafny.CodePoint('h'.codePointAt(0));
          let _index88 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_513_v28).length));
          (_513_v28)[_index88] = _468_v0;
          let _516_v29;
          _516_v29 = _dafny.Set.fromElements((_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))]);
          (globalState).f1 = ((_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))]).isLessThan(new BigNumber(((_516_v29).Union(_dafny.Set.fromElements((_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))]))).length));
          let _517_v30;
          let _init13 = ((_518_p1, _519_p2) => function (_520_i6) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_518_p1), _dafny.Seq.of((_module.D2.create_DC7(_518_p1, (_this).f32, _519_p2)).dtor_cf11));
          })(p1, p2);
          let _nw91 = Array((new BigNumber(5)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw91.length); _i0_13++) {
            _nw91[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _517_v30 = _nw91;
          let _521_v31;
          _521_v31 = _dafny.Seq.of(p1);
          let _522_v32;
          _522_v32 = _521_v31;
          let _index89 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_517_v30).length));
          (_517_v30)[_index89] = (_522_v32);
          let _523_v33;
          _523_v33 = _module.D6.create_DC12(_dafny.Set.fromElements(_509_v26));
          let _524_v34;
          _524_v34 = _dafny.Set.fromElements(_509_v26);
          let _index90 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_517_v30).length));
          (_517_v30)[_index90] = (((_dafny.MultiSet.fromElements(_523_v33)).equals(_dafny.MultiSet.fromElements(_module.D6.create_DC12(_524_v34), _523_v33))) ? (_dafny.Seq.Concat(_521_v31, _521_v31)) : (_dafny.Seq.Concat(_521_v31, _dafny.Seq.of((_this).f32, (_this).f32))));
          (globalState).f1 = (_this).f32;
        } else {
          let _525_v35;
          _525_v35 = _dafny.Set.fromElements(p1);
          let _526_v36;
          let _nw92 = new _module.C0();
          _nw92.__ctor(_510_v27, (_this).fm14(_525_v35, (_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))], globalState), (_dafny.ZERO).minus(p2), p0, p1);
          _526_v36 = _nw92;
          let _527_v37;
          _527_v37 = _dafny.Set.fromElements(_526_v36);
          _527_v37 = _527_v37;
          (globalState).f2 = (_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))];
          _525_v35 = _525_v35;
          let _index91 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length));
          (_509_v26)[_index91] = (p2).multipliedBy((_dafny.ZERO).minus((_509_v26)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_509_v26).length))]));
          let _528_v38;
          let _nw93 = new _module.C0();
          _nw93.__ctor(_526_v36.f35, (_this).f32, new BigNumber((_510_v27).length), p0, (_526_v36).fm7(globalState));
          _528_v38 = _nw93;
        }
        (globalState).f1 = (_this).f32;
        _509_v26 = _509_v26;
      }
      let _529_v39;
      _529_v39 = _module.D0.create_DC0((_this).f32);
      _529_v39 = _529_v39;
      let _530_v40;
      let _init14 = ((_531_p1) => function (_532_i7) {
        return _dafny.Map.Empty.slice().updateUnsafe(_531_p1,(_this).f32);
      })(p1);
      let _nw94 = Array((new BigNumber(15)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw94.length); _i0_14++) {
        _nw94[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _530_v40 = _nw94;
      let _533_v41;
      _533_v41 = _module.D8.create_DC16(_530_v40);
      let _534_v42;
      _534_v42 = _dafny.Map.Empty.slice().updateUnsafe((_533_v41).dtor_cf27,!(p1));
      let _535_v43;
      _535_v43 = _dafny.Seq.of(p1, (_this).f32, !(false));
      let _536_v44;
      _536_v44 = _dafny.Seq.of(p2);
      _534_v42 = (_534_v42).update(_530_v40, (_535_v43)[_module.__default.safeIndex(new BigNumber((_536_v44).length), new BigNumber((_535_v43).length))]);
      let _537_v45;
      _537_v45 = _module.D1.create_DC3(_468_v0);
      let _538_v46;
      _538_v46 = _module.D1.create_DC5(_537_v45);
      let _source6 = _538_v46;
      if (_source6.is_DC4) {
        let _539___mcc_h0 = (_source6).cf7;
        let _540___mcc_h1 = (_source6).cf8;
        let _541_cf8 = _540___mcc_h1;
        let _542_cf7 = _539___mcc_h0;
        if ((_535_v43)[_module.__default.safeIndex(p2, new BigNumber((_535_v43).length))]) {
          let _pat_let_tv13 = p1;
          let _543_v47;
          _543_v47 = _dafny.Seq.of(_529_v39, function (_pat_let14_0) {
            return function (_544_dt__update__tmp_h2) {
              return function (_pat_let15_0) {
                return function (_545_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_545_dt__update_hcf0_h0);
                }(_pat_let15_0);
              }(_pat_let_tv13);
            }(_pat_let14_0);
          }(_529_v39), _module.D0.create_DC0(_542_cf7));
          _543_v47 = _543_v47;
          (globalState).f1 = (_this).f32;
          let _index92 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((p0).length));
          (p0)[_index92] = (_this).f32;
          let _546_v48;
          _546_v48 = _dafny.Set.fromElements(p2);
          let _547_v49;
          _547_v49 = _dafny.MultiSet.fromElements(p1);
          let _548_v50;
          _548_v50 = _module.D6.create_DC14(new BigNumber(444), new BigNumber((_546_v48).length), _547_v49);
          let _549_v51;
          _549_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(p2, _module.__default.fm2(p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_548_v50).dtor_cf24)).length), globalState), globalState),p2);
          let _index93 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((p0).length));
          (p0)[_index93] = !(!(_549_v51).equals((_549_v51).Merge(_549_v51)));
          let _550_v52;
          let _nw95 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _550_v52 = _nw95;
          let _index94 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length));
          (_550_v52)[_index94] = p2;
          let _551_v53;
          _551_v53 = _dafny.Set.fromElements((p0)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((p0).length))], _541_cf8, (_this).f32);
          let _552_v54;
          _552_v54 = _dafny.Set.fromElements((_this).f32, (_this).fm14(_551_v53, p2, globalState));
          let _553_v55;
          _553_v55 = _module.D0.create_DC1(p2, _dafny.Seq.UnicodeFromString("jcg"), !((p0)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((p0).length))]), p2, (_this).fm14(_551_v53, p2, globalState));
          let _index95 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length));
          let _rhs117 = _module.__default.fm4(_553_v55, (p0)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((p0).length))], p2, globalState);
          let _rhs118 = p2;
          let _lhs102 = _550_v52;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length));
          _468_v0 = _rhs117;
          _lhs102[_lhs103] = _rhs118;
          let _554_v56;
          _554_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_dafny.Seq.IsPrefixOf(_536_v44, _536_v44));
          let _555_v57;
          let _nw96 = Array((new BigNumber(20)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _535_v43;
          _nw96[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_535_v43, _dafny.Seq.update(_535_v43, _module.__default.safeIndex((_550_v52)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length))], new BigNumber((_535_v43).length)), p1));
          _nw96[(new BigNumber(2)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(!((_this).f32));
          _nw96[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_542_cf7);
          _nw96[(new BigNumber(5)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(6)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(7)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of((_this).f32, (_535_v43)[_module.__default.safeIndex((_550_v52)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length))], new BigNumber((_535_v43).length))]), _module.__default.safeIndex((_550_v52)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length))], new BigNumber((_dafny.Seq.of((_this).f32, (_535_v43)[_module.__default.safeIndex((_550_v52)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length))], new BigNumber((_535_v43).length))])).length)), (p0)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((p0).length))]);
          _nw96[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_535_v43, _dafny.Seq.of(_542_cf7));
          _nw96[(new BigNumber(10)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(11)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(12)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(13)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(true);
          _nw96[(new BigNumber(15)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(16)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(17)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(18)).toNumber()] = _535_v43;
          _nw96[(new BigNumber(19)).toNumber()] = _535_v43;
          _555_v57 = _nw96;
          let _index96 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_555_v57).length));
          (_555_v57)[_index96] = _535_v43;
          let _index97 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length));
          let _index98 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_555_v57).length));
          let _rhs119 = p2;
          let _rhs120 = _542_cf7;
          let _rhs121 = _554_v56;
          let _rhs122 = _535_v43;
          let _lhs104 = _550_v52;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_550_v52).length));
          let _lhs106 = _555_v57;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_555_v57).length));
          _lhs104[_lhs105] = _rhs119;
          _542_cf7 = _rhs120;
          _554_v56 = _rhs121;
          _lhs106[_lhs107] = _rhs122;
        } else {
          let _556_v58;
          _556_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(609));
          (globalState).f0 = _module.__default.safeDivisionInt(p2, (((_556_v58).contains(!((_this).f32))) ? ((_556_v58).get(!((_this).f32))) : ((_dafny.ZERO).minus(p2))));
          let _557_v59;
          let _nw97 = Array((new BigNumber(3)).toNumber()).fill([]);
          _557_v59 = _nw97;
          let _index99 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length));
          (_557_v59)[_index99] = p0;
          let _index100 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length));
          (_557_v59)[_index100] = p0;
          (globalState).f0 = (p2).plus(p2);
          (globalState).f0 = _module.__default.safeModuloInt(new BigNumber(-529), (_module.__default.fm2(p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(625),false)).length), globalState)).multipliedBy(p2));
          let _arr0 = (_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))];
          let _index101 = _module.__default.safeIndex(new BigNumber(478), new BigNumber(((_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))]).length));
          _arr0[_index101] = p1;
          let _558_v60;
          _558_v60 = _dafny.Seq.UnicodeFromString("d");
          let _559_v61;
          let _nw98 = new _module.C0();
          _nw98.__ctor(_dafny.Seq.Concat(_558_v60, _558_v60), (_this).f32, _module.__default.safeModuloInt(p2, p2), (_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))], true);
          _559_v61 = _nw98;
          let _560_v62;
          let _nw99 = new _module.C0();
          _nw99.__ctor(_558_v60, false, p2, p0, (_this).f32);
          _560_v62 = _nw99;
          let _561_v63;
          _561_v63 = _module.D6.create_DC14((_559_v61).f26, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,_535_v43)).length), _dafny.MultiSet.fromElements(!(_542_cf7)));
          let _arr1 = (_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))];
          let _index102 = _module.__default.safeIndex(new BigNumber(478), new BigNumber(((_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))]).length));
          let _rhs123 = !((_561_v63).dtor_cf23).isEqualTo(p2);
          let _rhs124 = ((_559_v61).f26).minus((_559_v61).f26);
          let _rhs125 = _559_v61;
          let _rhs126 = _560_v62;
          let _lhs108 = (_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))];
          let _lhs109 = _module.__default.safeIndex(new BigNumber(478), new BigNumber(((_557_v59)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_557_v59).length))]).length));
          let _lhs110 = globalState;
          _lhs108[_lhs109] = _rhs123;
          _lhs110.f0 = _rhs124;
          _559_v61 = _rhs125;
          _560_v62 = _rhs126;
        }
        let _index103 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((p0).length));
        (p0)[_index103] = _541_cf8;
        let _index104 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((p0).length));
        (p0)[_index104] = true;
        let _562_v64;
        let _init15 = ((_563_v46) => function (_564_i8) {
          return _563_v46;
        })(_538_v46);
        let _nw100 = Array((new BigNumber(15)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw100.length); _i0_15++) {
          _nw100[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _562_v64 = _nw100;
        _562_v64 = _562_v64;
        (globalState).f1 = (_this).f32;
      } else if (_source6.is_DC3) {
        let _565___mcc_h2 = (_source6).cf6;
        let _566_cf6 = _565___mcc_h2;
        let _567_v65;
        let _568_v66;
        let _out14;
        let _out15;
        let _outcollector3 = (_this).m7(p1, _529_v39, _dafny.Seq.Create(_module.__default.abs(new BigNumber(666)), function (_569_i9) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), globalState);
        _out14 = _outcollector3[0];
        _out15 = _outcollector3[1];
        _567_v65 = _out14;
        _568_v66 = _out15;
        let _570_v67;
        _570_v67 = _dafny.Set.fromElements(p1);
        let _571_v68;
        _571_v68 = _module.D4.create_DC9((_570_v67).Union(_570_v67));
        let _source7 = _571_v68;
        if (_source7.is_DC10) {
          let _572___mcc_h4 = (_source7).cf16;
          let _573___mcc_h5 = (_source7).cf17;
          let _574_cf17 = _573___mcc_h5;
          let _575_cf16 = _572___mcc_h4;
          let _576_v69;
          let _nw101 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _576_v69 = _nw101;
          let _577_v70;
          _577_v70 = _dafny.Seq.of(_576_v69, _576_v69, _576_v69);
          let _578_v71;
          _578_v71 = _dafny.Seq.UnicodeFromString("cc");
          _576_v69 = (_577_v70)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_578_v71, _578_v71)).length), new BigNumber((_577_v70).length))];
          let _index105 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((p0).length));
          (p0)[_index105] = p1;
          let _index106 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((p0).length));
          (p0)[_index106] = true;
          let _index107 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((p0).length));
          (p0)[_index107] = !(!_dafny.Seq.contains(_578_v71, _575_cf16));
          let _579_v72;
          _579_v72 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ukvt"),p1);
          _579_v72 = _579_v72;
        } else {
          let _580___mcc_h6 = (_source7).cf15;
          let _581_cf15 = _580___mcc_h6;
          let _582_v73;
          _582_v73 = _dafny.Seq.UnicodeFromString("qten");
          let _583_v74;
          let _nw102 = new _module.C0();
          _nw102.__ctor(_582_v73, !((_this).f32), p2, p0, true);
          _583_v74 = _nw102;
          let _584_v75;
          _584_v75 = _dafny.Seq.of(_583_v74, _583_v74);
          let _585_v76;
          _585_v76 = _dafny.Map.Empty.slice().updateUnsafe((_535_v43)[_module.__default.safeIndex(p2, new BigNumber((_535_v43).length))],_dafny.Seq.Concat(_dafny.Seq.update(_584_v75, _module.__default.safeIndex(p2, new BigNumber((_584_v75).length)), _583_v74), _584_v75));
          _585_v76 = (_585_v76).update(p1, _584_v75);
          let _586_v77;
          let _init16 = ((_587_p2) => function (_588_i10) {
            return _module.__default.safeDivisionInt(_588_i10, _587_p2);
          })(p2);
          let _nw103 = Array((new BigNumber(29)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw103.length); _i0_16++) {
            _nw103[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _586_v77 = _nw103;
          let _589_v78;
          let _nw104 = Array((new BigNumber(4)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = _586_v77;
          _nw104[(_dafny.ONE).toNumber()] = _586_v77;
          _nw104[(new BigNumber(2)).toNumber()] = _586_v77;
          _nw104[(new BigNumber(3)).toNumber()] = _586_v77;
          _589_v78 = _nw104;
          _589_v78 = _589_v78;
          let _590_v79;
          _590_v79 = _dafny.MultiSet.fromElements((_this).f32);
          let _591_v80;
          _591_v80 = _module.D6.create_DC14(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(63)), ((_592_cf6) => function (_593_i11) {
  return _592_cf6;
})(_566_cf6)))).cardinality()), p2, _590_v79);
          let _594_v81;
          _594_v81 = _dafny.Set.fromElements(p2);
          let _595_v82;
          _595_v82 = _module.D6.create_DC13((_591_v80).dtor_cf23, new BigNumber(-119), _594_v81);
          let _596_v83;
          _596_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,p2);
          _595_v82 = _module.D6.create_DC13((((_596_v83).contains(!(!(true)))) ? ((_596_v83).get(!(!(true)))) : (new BigNumber(682))), _module.__default.safeModuloInt(new BigNumber((_535_v43).length), p2), _594_v81);
          (globalState).f2 = p2;
        }
        let _597_v84;
        _597_v84 = _dafny.MultiSet.fromElements((_this).f32, (_this).f32, !(p1));
        _597_v84 = (_dafny.MultiSet.fromElements(p1)).Difference(_597_v84);
        if ((_this).f32) {
          _536_v44 = _dafny.Seq.Concat(_536_v44, _536_v44);
          r0 = _module.__default.fm1(globalState);
          let _598_v85;
          let _nw105 = Array((new BigNumber(23)).toNumber()).fill(_module.D0.Default());
          _598_v85 = _nw105;
          let _599_v86;
          _599_v86 = _module.D0.create_DC2();
          let _index108 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_598_v85).length));
          (_598_v85)[_index108] = _599_v86;
          let _index109 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_598_v85).length));
          (_598_v85)[_index109] = _module.D0.create_DC2();
          let _600_v87;
          _600_v87 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(p2));
          let _601_v88;
          _601_v88 = _dafny.MultiSet.fromElements(_600_v87);
          let _rhs127 = p1;
          let _rhs128 = (_601_v88).Difference(_601_v88);
          let _lhs111 = globalState;
          _lhs111.f1 = _rhs127;
          _567_v65 = _rhs128;
          let _602_v89;
          _602_v89 = _dafny.Seq.UnicodeFromString("kbsu");
          let _603_v90;
          let _nw106 = new _module.C0();
          _nw106.__ctor(_602_v89, p1, p2, p0, !(p1));
          _603_v90 = _nw106;
        } else {
          let _604_v91;
          _604_v91 = _dafny.Seq.UnicodeFromString("doq");
          let _605_v92;
          let _nw107 = Array((new BigNumber(9)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = _604_v91;
          _nw107[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xqocpbbe"), _604_v91);
          _nw107[(new BigNumber(2)).toNumber()] = _604_v91;
          _nw107[(new BigNumber(3)).toNumber()] = _604_v91;
          _nw107[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_604_v91, _604_v91);
          _nw107[(new BigNumber(5)).toNumber()] = _604_v91;
          _nw107[(new BigNumber(6)).toNumber()] = _604_v91;
          _nw107[(new BigNumber(7)).toNumber()] = _604_v91;
          _nw107[(new BigNumber(8)).toNumber()] = _604_v91;
          _605_v92 = _nw107;
          _605_v92 = _605_v92;
          let _index110 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_605_v92).length));
          (_605_v92)[_index110] = _604_v91;
          let _index111 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_605_v92).length));
          let _rhs129 = !((p2).plus(new BigNumber(724))).isEqualTo(p2);
          let _rhs130 = _604_v91;
          let _lhs112 = globalState;
          let _lhs113 = _605_v92;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_605_v92).length));
          _lhs112.f1 = _rhs129;
          _lhs113[_lhs114] = _rhs130;
          r1 = ((p2).isLessThan(new BigNumber((_570_v67).length))) === (p1);
          let _606_v93;
          let _init17 = function (_607_i12) {
            return _module.__default.safeModuloInt(_607_i12, new BigNumber(-966));
          };
          let _nw108 = Array((new BigNumber(17)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw108.length); _i0_17++) {
            _nw108[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _606_v93 = _nw108;
          let _index112 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_606_v93).length));
          (_606_v93)[_index112] = (_dafny.ZERO).minus(p2);
          let _index113 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_606_v93).length));
          (_606_v93)[_index113] = (p2).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), ((_608_cf6) => function (_609_i13) {
            return _608_cf6;
          })(_566_cf6))).length));
          let _index114 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_606_v93).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_606_v93).length));
          let _rhs131 = new BigNumber(-988);
          let _rhs132 = p2;
          let _lhs115 = _606_v93;
          let _lhs116 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_606_v93).length));
          let _lhs117 = _606_v93;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_606_v93).length));
          _lhs115[_lhs116] = _rhs131;
          _lhs117[_lhs118] = _rhs132;
        }
      } else {
        let _610___mcc_h3 = (_source6).cf9;
        let _611_cf9 = _610___mcc_h3;
        let _612_v94;
        _612_v94 = _dafny.Seq.UnicodeFromString("rmndt");
        let _613_v95;
        let _nw109 = new _module.C0();
        _nw109.__ctor(_612_v94, (_this).f32, p2, p0, p1);
        _613_v95 = _nw109;
        let _614_v96;
        let _nw110 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
        _614_v96 = _nw110;
        let _615_v97;
        _615_v97 = _module.D4.create_DC10(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)));
        let _616_v98;
        let _nw111 = Array((new BigNumber(26)).toNumber());
        _nw111[(_dafny.ZERO).toNumber()] = _468_v0;
        _nw111[(_dafny.ONE).toNumber()] = _468_v0;
        _nw111[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
        _nw111[(new BigNumber(3)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(4)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(5)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(6)).toNumber()] = (_615_v97).dtor_cf16;
        _nw111[(new BigNumber(7)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(8)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(9)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw111[(new BigNumber(11)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw111[(new BigNumber(13)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw111[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw111[(new BigNumber(16)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(17)).toNumber()] = (_613_v95.f35)[_module.__default.safeIndex(p2, new BigNumber((_613_v95.f35).length))];
        _nw111[(new BigNumber(18)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(19)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(20)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(21)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(22)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(23)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(24)).toNumber()] = _468_v0;
        _nw111[(new BigNumber(25)).toNumber()] = _468_v0;
        _616_v98 = _nw111;
        let _617_v99;
        _617_v99 = _dafny.Map.Empty.slice().updateUnsafe(p2,_616_v98);
        let _618_v100;
        _618_v100 = _dafny.Map.Empty.slice().updateUnsafe(_617_v99,_614_v96);
        _614_v96 = (((_618_v100).contains((_617_v99).Merge(_617_v99))) ? ((_618_v100).get((_617_v99).Merge(_617_v99))) : (_614_v96));
        let _619_v101;
        let _nw112 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _619_v101 = _nw112;
        let _index116 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_619_v101).length));
        (_619_v101)[_index116] = p2;
        let _index117 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((p0).length));
        (p0)[_index117] = _dafny.Seq.IsProperPrefixOf(_535_v43, _dafny.Seq.of(p1, p1));
        let _index118 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_619_v101).length));
        let _index119 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((p0).length));
        let _rhs133 = (p2).multipliedBy((new BigNumber((_dafny.Seq.update(_613_v95.f35, _module.__default.safeIndex(new BigNumber(439), new BigNumber((_613_v95.f35).length)), _468_v0)).length)).multipliedBy(p2));
        let _rhs134 = false;
        let _lhs119 = _619_v101;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_619_v101).length));
        let _lhs121 = p0;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((p0).length));
        _lhs119[_lhs120] = _rhs133;
        _lhs121[_lhs122] = _rhs134;
        (globalState).f1 = ((p0)[_module.__default.safeIndex(new BigNumber(796), new BigNumber((p0).length))]) === (p1);
      }
      let _620_v102;
      _620_v102 = _dafny.Seq.UnicodeFromString("km");
      r0 = _620_v102;
      r1 = p1;
      return [r0, r1];
    }
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f25 = false;
      this._f26 = _dafny.ZERO;
      this._f31 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f31, f25, f26) {
      let _this = this;
      (_this)._f31 = f31;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f31;
    };
    fm7(globalState) {
      let _this = this;
      return (_module.D0.create_DC0((_this).f25)).dtor_cf0;
    };
    fm11(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(function (_source8) {
        if (_source8.is_DC1) {
          let _621___mcc_h0 = (_source8).cf1;
          let _622___mcc_h1 = (_source8).cf2;
          let _623___mcc_h2 = (_source8).cf3;
          let _624___mcc_h3 = (_source8).cf4;
          let _625___mcc_h4 = (_source8).cf5;
          let _626_cf5 = _625___mcc_h4;
          let _627_cf4 = _624___mcc_h3;
          let _628_cf3 = _623___mcc_h2;
          let _629_cf2 = _622___mcc_h1;
          let _630_cf1 = _621___mcc_h0;
          return _627_cf4;
        } else if (_source8.is_DC2) {
          return _module.__default.safeModuloInt((_this).f26, (_this).f31);
        } else {
          let _631___mcc_h5 = (_source8).cf0;
          let _632_cf0 = _631___mcc_h5;
          return _module.__default.safeDivisionInt(new BigNumber(523), (_this).f26);
        }
      }(_module.D0.create_DC1((_this).f31, _dafny.Seq.UnicodeFromString("lyagi"), (_this).f25, new BigNumber(-991), (_this).f25)));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _hi4 = (_this).f31;
      for (let _633_i0 = (_this).f31; _633_i0.isLessThan(_hi4); _633_i0 = _633_i0.plus(_dafny.ONE)) {
        let _634_v0;
        _634_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        _634_v0 = (_634_v0).update((_this).f25, p2);
        let _635_v1;
        _635_v1 = _module.D0.create_DC2();
        let _636_v2;
        _636_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(64),p1);
        let _637_v3;
        _637_v3 = _dafny.Seq.UnicodeFromString("dxrsfndo");
        let _638_v4;
        _638_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(154),new BigNumber((_637_v3).length));
        let _639_v5;
        _639_v5 = _dafny.Set.fromElements(new BigNumber((p0).cardinality()));
        let _640_v6;
        _640_v6 = _module.D2.create_DC6(_dafny.Set.fromElements((_this).f26));
        let _rhs135 = (_this).f25;
        let _rhs136 = _635_v1;
        let _rhs137 = (_this).fm6(p2, (((_636_v2).contains((((_638_v4).contains((_this).f26)) ? ((_638_v4).get((_this).f26)) : (_633_i0)))) ? ((_636_v2).get((((_638_v4).contains((_this).f26)) ? ((_638_v4).get((_this).f26)) : (_633_i0)))) : (p2)), _module.__default.safeDivisionInt((_this).f26, _633_i0), globalState);
        let _rhs138 = (_639_v5).equals(((_640_v6).dtor_cf10).Difference(_639_v5));
        let _lhs123 = globalState;
        let _lhs124 = globalState;
        let _lhs125 = globalState;
        _lhs123.f1 = _rhs135;
        _635_v1 = _rhs136;
        _lhs124.f0 = _rhs137;
        _lhs125.f1 = _rhs138;
        let _641_v7;
        _641_v7 = _module.D0.create_DC0(true);
        let _642_v8;
        _642_v8 = _dafny.Map.Empty.slice().updateUnsafe(_641_v7,(_this).f26);
        (globalState).f2 = ((!(!(_module.__default.fm5(p1, _642_v8, new BigNumber((function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-118)), function (_643_i1) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })).Elements) {
            let _644_v9 = _compr_12;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-118)), function (_643_i1) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            }), _644_v9)) {
              _coll12.add(_644_v9);
            }
          }
          return _coll12;
        }()).length), globalState)))) ? (_633_i0) : ((_this).f26));
        let _645_v10;
        _645_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,_637_v3);
        _645_v10 = (_645_v10).update(_module.__default.fm5(p1, _642_v8, (_this).f31, globalState), _637_v3);
      }
      (globalState).f0 = (_this).f31;
      let _646_v11;
      _646_v11 = _dafny.Seq.UnicodeFromString("mt");
      let _647_v12;
      _647_v12 = _module.D1.create_DC4(p1, p1);
      let _pat_let_tv14 = _647_v12;
      let _pat_let_tv15 = p1;
      let _pat_let_tv16 = p2;
      let _pat_let_tv17 = p2;
      let _pat_let_tv18 = p1;
      let _pat_let_tv19 = p1;
      let _rhs139 = p1;
      let _rhs140 = _module.__default.fm1(globalState);
      let _rhs141 = function (_source9) {
        if (_source9.is_DC4) {
          let _648___mcc_h0 = (_source9).cf7;
          let _649___mcc_h1 = (_source9).cf8;
          let _650_cf8 = _649___mcc_h1;
          let _651_cf7 = _648___mcc_h0;
          return (_pat_let_tv14).dtor_cf8;
        } else if (_source9.is_DC3) {
          let _652___mcc_h2 = (_source9).cf6;
          let _653_cf6 = _652___mcc_h2;
          return (_this).f25;
        } else {
          let _654___mcc_h3 = (_source9).cf9;
          let _655_cf9 = _654___mcc_h3;
          return (_dafny.MultiSet.fromElements((_this).f25)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_pat_let_tv15, _pat_let_tv16, _pat_let_tv17, _pat_let_tv18, _pat_let_tv19));
        }
      }(_647_v12);
      let _rhs142 = !(!(!(p1)));
      let _rhs143 = ((_this).f31).isLessThan((_this).f26);
      let _lhs126 = globalState;
      let _lhs127 = globalState;
      r0 = _rhs139;
      _646_v11 = _rhs140;
      _lhs126.f1 = _rhs141;
      _lhs127.f1 = _rhs142;
      r0 = _rhs143;
      let _hi5 = ((_this).f31).multipliedBy((_this).f26);
      for (let _656_i2 = _module.__default.safeModuloInt((_this).f26, (_this).f31); _656_i2.isLessThan(_hi5); _656_i2 = _656_i2.plus(_dafny.ONE)) {
        let _657_v13;
        _657_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,p2);
        let _658_v14;
        _658_v14 = _dafny.Map.Empty.slice().updateUnsafe(_657_v13,true);
        let _659_v17;
        _659_v17 = _module.D0.create_DC0(p1);
        let _660_v18;
        _660_v18 = _module.D0.create_DC1(_656_i2, _646_v11, p1, (_this).f26, (((_657_v13).contains(_656_i2)) ? ((_657_v13).get(_656_i2)) : (_module.__default.fm5((_this).f25, function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of (function () {
    let _coll14 = new _dafny.Map();
    for (const _compr_14 of (_dafny.Seq.of(_659_v17, _659_v17)).Elements) {
      let _661_v16 = _compr_14;
      if (_dafny.Seq.contains(_dafny.Seq.of(_659_v17, _659_v17), _661_v16)) {
        _coll14.push([_661_v16,(_this).f25]);
      }
    }
    return _coll14;
  }()).Keys.Elements) {
    let _662_v15 = _compr_13;
    if ((function () {
      let _coll15 = new _dafny.Map();
      for (const _compr_15 of (_dafny.Seq.of(_659_v17, _659_v17)).Elements) {
        let _661_v16 = _compr_15;
        if (_dafny.Seq.contains(_dafny.Seq.of(_659_v17, _659_v17), _661_v16)) {
          _coll15.push([_661_v16,(_this).f25]);
        }
      }
      return _coll15;
    }()).contains(_662_v15)) {
      _coll13.push([_662_v15,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_656_i2,(_this).f31)).length)]);
    }
  }
  return _coll13;
}(), new BigNumber((_646_v11).length), globalState))));
        let _663_v19;
        _663_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_658_v14).length),(_dafny.ZERO).minus((_660_v18).dtor_cf4));
        let _664_v20;
        _664_v20 = _dafny.Map.Empty.slice().updateUnsafe(_646_v11,_646_v11);
        _663_v19 = (_663_v19).update((_this).f26, (_this).fm11(_664_v20, (_this).f31, (_this).f31, globalState));
        (globalState).f0 = ((new BigNumber(281)).multipliedBy((_this).f31)).multipliedBy((_this).f31);
        let _665_v21;
        _665_v21 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(new BigNumber((p0).cardinality())).isEqualTo(new BigNumber((_646_v11).length)));
        _665_v21 = (_665_v21).update(p2, !(p1));
        _663_v19 = (_663_v19).update((new BigNumber((_665_v21).length)).minus((_this).f31), (_656_i2).multipliedBy((_this).f31));
      }
      let _666_v22;
      _666_v22 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f31), new BigNumber(537));
      let _667_v23;
      _667_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_666_v22).length));
      let _668_v24;
      _668_v24 = _dafny.Set.fromElements(_module.D2.create_DC6(_module.__default.fm12(new BigNumber((_667_v23).length), (_this).f31, globalState)));
      let _669_v25;
      _669_v25 = _dafny.Set.fromElements((_this).f31);
      let _670_v26;
      _670_v26 = _module.D2.create_DC6(_669_v25);
      _668_v24 = (_668_v24).Difference(_dafny.Set.fromElements(_670_v26, _module.D2.create_DC6(_dafny.Set.fromElements((_this).f26, new BigNumber(969), (_this).f31, (_this).f31)), _670_v26));
      let _hi6 = ((_this).f26).plus((_this).f26);
      for (let _671_i3 = (_this).f26; _671_i3.isLessThan(_hi6); _671_i3 = _671_i3.plus(_dafny.ONE)) {
        (globalState).f0 = new BigNumber(-20);
        let _672_v27;
        let _nw113 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
        _672_v27 = _nw113;
        let _673_v28;
        _673_v28 = _dafny.Seq.of((_this).f25);
        let _674_v29;
        _674_v29 = _dafny.Set.fromElements(_673_v28, _673_v28);
        let _index120 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_672_v27).length));
        (_672_v27)[_index120] = _674_v29;
        let _index121 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_672_v27).length));
        (_672_v27)[_index121] = _module.__default.fm13(p2, globalState);
        let _675_v30;
        _675_v30 = _module.D0.create_DC2();
        let _source10 = _675_v30;
        if (_source10.is_DC1) {
          let _676___mcc_h4 = (_source10).cf1;
          let _677___mcc_h5 = (_source10).cf2;
          let _678___mcc_h6 = (_source10).cf3;
          let _679___mcc_h7 = (_source10).cf4;
          let _680___mcc_h8 = (_source10).cf5;
          let _681_cf5 = _680___mcc_h8;
          let _682_cf4 = _679___mcc_h7;
          let _683_cf3 = _678___mcc_h6;
          let _684_cf2 = _677___mcc_h5;
          let _685_cf1 = _676___mcc_h4;
          let _686_v31;
          let _nw114 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
          _686_v31 = _nw114;
          let _687_v32;
          _687_v32 = _dafny.MultiSet.fromElements(_667_v23);
          let _index122 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_686_v31).length));
          (_686_v31)[_index122] = (_dafny.MultiSet.fromElements(_667_v23)).Union(_687_v32);
          let _688_v33;
          _688_v33 = _687_v32;
          let _index123 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_686_v31).length));
          (_686_v31)[_index123] = (_688_v33);
          let _689_v34;
          _689_v34 = _module.D2.create_DC7((_this).f25, p2, (_this).f31);
          let _690_v35;
          let _nw115 = Array((new BigNumber(7)).toNumber());
          _nw115[(_dafny.ZERO).toNumber()] = _681_cf5;
          _nw115[(_dafny.ONE).toNumber()] = _681_cf5;
          _nw115[(new BigNumber(2)).toNumber()] = true;
          _nw115[(new BigNumber(3)).toNumber()] = (_681_cf5) && (p2);
          _nw115[(new BigNumber(4)).toNumber()] = (_689_v34).dtor_cf12;
          _nw115[(new BigNumber(5)).toNumber()] = _681_cf5;
          _nw115[(new BigNumber(6)).toNumber()] = p2;
          _690_v35 = _nw115;
          let _index124 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_690_v35).length));
          (_690_v35)[_index124] = p1;
          let _index125 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_690_v35).length));
          (_690_v35)[_index125] = p1;
          let _691_v36;
          _691_v36 = new _dafny.CodePoint('c'.codePointAt(0));
          _691_v36 = _691_v36;
          let _692_v37;
          _692_v37 = _dafny.Map.Empty.slice().updateUnsafe(_682_cf4,_685_cf1);
          let _693_v38;
          _693_v38 = _dafny.Map.Empty.slice().updateUnsafe(_692_v37,_690_v35);
          let _rhs144 = !(_693_v38).equals((_693_v38).Merge(_693_v38));
          let _rhs145 = _dafny.Set.fromElements(_685_cf1, (_this).f26, _module.__default.safeDivisionInt(new BigNumber(960), _682_cf4));
          let _rhs146 = (((_690_v35)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_690_v35).length))]) ? ((_this).f26) : (((_this).f31).minus(_682_cf4)));
          let _rhs147 = (_this).f31;
          let _rhs148 = false;
          let _lhs128 = globalState;
          let _lhs129 = globalState;
          let _lhs130 = globalState;
          _lhs128.f1 = _rhs144;
          _669_v25 = _rhs145;
          _lhs129.f2 = _rhs146;
          _lhs130.f0 = _rhs147;
          _683_cf3 = _rhs148;
        } else if (_source10.is_DC2) {
          let _694_v39;
          _694_v39 = _dafny.Seq.of(_646_v11);
          let _695_v40;
          _695_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f31);
          _666_v22 = _module.__default.fm3(_671_i3, !(((_dafny.ZERO).minus((_this).f31)).isEqualTo(new BigNumber((_694_v39).length))), new BigNumber(766), !(_695_v40).equals(function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(914), new BigNumber(-407))) {
              let _696_v41 = _compr_16;
              if (((new BigNumber(914)).isLessThanOrEqualTo(_696_v41)) && ((_696_v41).isLessThan(new BigNumber(-407)))) {
                _coll16.push([(_696_v41).plus((_this).f31),(_dafny.ZERO).minus((_this).f31)]);
              }
            }
            return _coll16;
          }()), globalState);
          (globalState).f1 = (_this).fm7(globalState);
          let _697_v42;
          let _nw116 = new _module.C1();
          _nw116.__ctor(p2);
          _697_v42 = _nw116;
          (globalState).f1 = (_673_v28)[_module.__default.safeIndex((_this).f26, new BigNumber((_673_v28).length))];
        } else {
          let _698___mcc_h9 = (_source10).cf0;
          let _699_cf0 = _698___mcc_h9;
          let _700_v43;
          _700_v43 = _module.D0.create_DC0(_699_cf0);
          let _701_v44;
          let _nw117 = Array((new BigNumber(6)).toNumber()).fill(false);
          _701_v44 = _nw117;
          let _702_v45;
          let _nw118 = new _module.C0();
          _nw118.__ctor(_dafny.Seq.UnicodeFromString("me"), (_700_v43).dtor_cf0, _671_i3, _701_v44, false);
          _702_v45 = _nw118;
          let _703_v46;
          let _nw119 = new _module.C1();
          _nw119.__ctor(((p1) ? (p1) : (!(p2))));
          _703_v46 = _nw119;
          _673_v28 = _673_v28;
          let _704_v47;
          _704_v47 = _673_v28;
          let _705_v48;
          _705_v48 = _dafny.Seq.of(_673_v28, _704_v47, _704_v47, _704_v47, _704_v47);
          let _706_v49;
          _706_v49 = _dafny.Seq.of(_701_v44);
          (globalState).f1 = _dafny.areEqual(_705_v48, _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_704_v47, _704_v47, _704_v47, _704_v47), _module.__default.safeIndex(new BigNumber((_706_v49).length), new BigNumber((_dafny.Seq.of(_704_v47, _704_v47, _704_v47, _704_v47)).length)), _dafny.Seq.of(p1)), _705_v48));
        }
        let _707_v50;
        _707_v50 = _dafny.MultiSet.fromElements(!(p1), true);
        let _708_v51;
        _708_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(p2),new BigNumber((_707_v50).cardinality()));
        let _709_v52;
        _709_v52 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), ((_710_v11) => function (_711_i4) {
          return (_710_v11)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ferpotds")).length), new BigNumber((_710_v11).length))];
        })(_646_v11)));
        if ((_module.__default.fm5(p1, _708_v51, (_this).f31, globalState)) && ((new BigNumber(((_709_v52)[_module.__default.safeIndex((_this).f26, new BigNumber((_709_v52).length))]).length)).isLessThan(new BigNumber(-875)))) {
          let _712_v53;
          _712_v53 = _dafny.Map.Empty.slice().updateUnsafe(_671_i3,(_673_v28)[_module.__default.safeIndex(_671_i3, new BigNumber((_673_v28).length))]);
          let _713_v54;
          _713_v54 = new _dafny.CodePoint('r'.codePointAt(0));
          let _714_v55;
          let _nw120 = Array((new BigNumber(12)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f26);
          _nw120[(_dafny.ONE).toNumber()] = (_this).f31;
          _nw120[(new BigNumber(2)).toNumber()] = (_this).f26;
          _nw120[(new BigNumber(3)).toNumber()] = new BigNumber((((_712_v53).update((_this).f31, (_this).f25)).Merge(_712_v53)).length);
          _nw120[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f26), (_this).f26);
          _nw120[(new BigNumber(5)).toNumber()] = (((p0).contains(new BigNumber(926))) ? ((p0).get(new BigNumber(926))) : ((_this).f31));
          _nw120[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("b"), _module.__default.safeIndex((_this).f31, new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)), _713_v54), _module.__default.safeIndex((_this).f26, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("b"), _module.__default.safeIndex((_this).f31, new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)), _713_v54)).length)), new _dafny.CodePoint('s'.codePointAt(0)))).length);
          _nw120[(new BigNumber(7)).toNumber()] = _671_i3;
          _nw120[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt((_this).f26, (_this).f31);
          _nw120[(new BigNumber(9)).toNumber()] = _671_i3;
          _nw120[(new BigNumber(10)).toNumber()] = (_this).f31;
          _nw120[(new BigNumber(11)).toNumber()] = _module.__default.fm2(_671_i3, new BigNumber(-150), globalState);
          _714_v55 = _nw120;
          let _index126 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_714_v55).length));
          (_714_v55)[_index126] = ((p1) ? ((_this).f31) : (new BigNumber((_dafny.Seq.of((_this).f26, _671_i3)).length)));
          let _index127 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_714_v55).length));
          (_714_v55)[_index127] = _671_i3;
          let _715_v56;
          let _init18 = ((_716_p2) => function (_717_i5) {
            return _716_p2;
          })(p2);
          let _nw121 = Array((new BigNumber(23)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw121.length); _i0_18++) {
            _nw121[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _715_v56 = _nw121;
          let _718_v57;
          let _nw122 = new _module.C0();
          _nw122.__ctor(_646_v11, (_this).f25, new BigNumber(296), _715_v56, p2);
          _718_v57 = _nw122;
          let _index128 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_715_v56).length));
          (_715_v56)[_index128] = p2;
          let _719_v58;
          _719_v58 = _module.D0.create_DC0(p2);
          let _pat_let_tv20 = p1;
          let _index129 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_715_v56).length));
          (_715_v56)[_index129] = ((_this).f25) && (_module.__default.fm5((_this).f25, (_708_v51).update(function (_pat_let16_0) {
            return function (_720_dt__update__tmp_h0) {
              return function (_pat_let17_0) {
                return function (_721_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_721_dt__update_hcf0_h0);
                }(_pat_let17_0);
              }(_pat_let_tv20);
            }(_pat_let16_0);
          }(_719_v58), (_714_v55)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_714_v55).length))]), (_this).f26, globalState));
          let _722_v59;
          _722_v59 = _dafny.Seq.of(_715_v56);
          let _723_v60;
          _723_v60 = _dafny.Seq.of((_722_v59)[_module.__default.safeIndex((_this).f31, new BigNumber((_722_v59).length))]);
          _723_v60 = (((p0).IsProperSubsetOf(p0)) ? (_dafny.Seq.Concat(_722_v59, _723_v60)) : (_723_v60));
          (globalState).f2 = _671_i3;
        } else {
          let _724_v61;
          let _nw123 = Array((new BigNumber(5)).toNumber());
          _nw123[(_dafny.ZERO).toNumber()] = (_this).f25;
          _nw123[(_dafny.ONE).toNumber()] = (_this).f25;
          _nw123[(new BigNumber(2)).toNumber()] = p1;
          _nw123[(new BigNumber(3)).toNumber()] = !(p1);
          _nw123[(new BigNumber(4)).toNumber()] = p1;
          _724_v61 = _nw123;
          let _725_v62;
          let _nw124 = new _module.C0();
          _nw124.__ctor(_646_v11, p1, new BigNumber((_dafny.Seq.UnicodeFromString("kiivsm")).length), _724_v61, p2);
          _725_v62 = _nw124;
          let _726_v63;
          _726_v63 = _dafny.Seq.of(_725_v62);
          let _727_v64;
          _727_v64 = _dafny.Map.Empty.slice().updateUnsafe(true,_726_v63);
          let _728_v65;
          _728_v65 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_727_v64).contains((_725_v62).f25)) ? ((_727_v64).get((_725_v62).f25)) : (_726_v63)));
          let _729_v66;
          _729_v66 = _module.D9.create_DC18(_727_v64);
          _728_v65 = (((_729_v66).dtor_cf30).update((_725_v62).f25, _726_v63)).Merge(_727_v64);
          (globalState).f1 = (((_725_v62).f25) ? (true) : (_dafny.Seq.IsProperPrefixOf(_646_v11, _dafny.Seq.UnicodeFromString("wmubls"))));
          (globalState).f2 = new BigNumber((_673_v28).length);
          let _730_v67;
          let _out16;
          _out16 = _module.__default.m0(p1, globalState);
          _730_v67 = _out16;
          let _731_v68;
          _731_v68 = _dafny.Map.Empty.slice().updateUnsafe(_646_v11,_730_v67);
          (globalState).f1 = (_dafny.Map.Empty.slice().updateUnsafe(_646_v11,(_this).fm6((_this).f25, (_this).f25, _671_i3, globalState))).equals(_731_v68);
        }
      }
      r0 = (_this).fm7(globalState);
      return r0;
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _732_v0;
      _732_v0 = _dafny.Set.fromElements(p0, (_this).f25, (_this).f25, (_this).f25, !((_this).f25) || (!(p0)));
      let _733_v1;
      let _nw125 = Array((new BigNumber(4)).toNumber()).fill(false);
      _733_v1 = _nw125;
      let _index130 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
      (_733_v1)[_index130] = p0;
      let _734_v2;
      _734_v2 = _module.D0.create_DC0(p0);
      let _735_v3;
      _735_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5((_this).f25, _module.__default.fm25((_this).f26, (_734_v2).dtor_cf0, p0, globalState), new BigNumber(345), globalState),_732_v0);
      let _736_v4;
      _736_v4 = _dafny.Seq.UnicodeFromString("nf");
      let _737_v5;
      _737_v5 = _dafny.Seq.of(_module.__default.fm26((_this).f26, p2, p0, globalState));
      let _738_v6;
      _738_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2);
      let _739_v7;
      _739_v7 = _dafny.MultiSet.fromElements(_738_v6, _738_v6);
      let _740_v8;
      _740_v8 = _dafny.Map.Empty.slice().updateUnsafe(_734_v2,(_this).f31);
      let _index131 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
      let _rhs149 = (((_735_v3).contains(!(p0) || ((_this).f25))) ? ((_735_v3).get(!(p0) || ((_this).f25))) : (_732_v0));
      let _rhs150 = (_dafny.ZERO).minus((((_this).f25) ? (new BigNumber((_dafny.Seq.Concat(_736_v4, _736_v4)).length)) : ((_this).f31)));
      let _rhs151 = ((_739_v7).update(_738_v6, _module.__default.abs((_this).f31))).IsProperSubsetOf((_737_v5)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f31), new BigNumber((_737_v5).length))]);
      let _rhs152 = ((_module.__default.fm5((_this).f25, _740_v8, (_this).f31, globalState)) ? (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), ((_741_p2) => function (_742_i0) {
        return _741_p2;
      })(p2)), p2)) : ((_this).f25));
      let _lhs131 = globalState;
      let _lhs132 = globalState;
      let _lhs133 = _733_v1;
      let _lhs134 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
      _732_v0 = _rhs149;
      _lhs131.f0 = _rhs150;
      _lhs132.f1 = _rhs151;
      _lhs133[_lhs134] = _rhs152;
      let _hi7 = (_this).f31;
      for (let _743_i1 = (_this).f31; _743_i1.isLessThan(_hi7); _743_i1 = _743_i1.plus(_dafny.ONE)) {
        let _744_v9;
        let _init19 = ((_745_v4) => function (_746_i2) {
          return _745_v4;
        })(_736_v4);
        let _nw126 = Array((new BigNumber(19)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw126.length); _i0_19++) {
          _nw126[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _744_v9 = _nw126;
        let _747_v10;
        _747_v10 = _dafny.Seq.of((_this).f31);
        let _748_v11;
        _748_v11 = _dafny.Map.Empty.slice().updateUnsafe(_744_v9,_747_v10);
        _748_v11 = (_748_v11).update(_744_v9, _747_v10);
        let _749_v12;
        _749_v12 = _dafny.Map.Empty.slice().updateUnsafe(_747_v10,p2);
        let _750_v13;
        let _nw127 = new _module.C0();
        _nw127.__ctor(_736_v4, !(((_this).f31).isLessThan((_this).f26)), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_736_v4, _module.__default.safeIndex(new BigNumber(438), new BigNumber((_736_v4).length)), (((_749_v12).contains(_747_v10)) ? ((_749_v12).get(_747_v10)) : (p2))), _module.__default.safeIndex(_743_i1, new BigNumber((_dafny.Seq.update(_736_v4, _module.__default.safeIndex(new BigNumber(438), new BigNumber((_736_v4).length)), (((_749_v12).contains(_747_v10)) ? ((_749_v12).get(_747_v10)) : (p2)))).length)), p2)).length), _733_v1, !((_this).f25));
        _750_v13 = _nw127;
        let _index132 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
        (_733_v1)[_index132] = !((_this).f25);
        let _751_v14;
        _751_v14 = _dafny.Map.Empty.slice().updateUnsafe((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))],_743_i1);
        let _752_v15;
        _752_v15 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))],(_this).f31), (_751_v14).update((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))], new BigNumber(-691)));
        let _753_v16;
        _753_v16 = _752_v15;
        let _source11 = _dafny.MultiSet.fromElements(_751_v14);
        let _754___mcc_h0 = _source11;
        let _755_cf14 = _754___mcc_h0;
        let _756_v17;
        _756_v17 = _dafny.MultiSet.fromElements(p0);
        let _757_v18;
        let _nw128 = new _module.C0();
        _nw128.__ctor(_750_v13.f35, (_756_v17).IsDisjointFrom(_756_v17), new BigNumber(((_732_v0).Intersect(_dafny.Set.fromElements(p0))).length), _733_v1, (_module.__default.fm27((_this).f26, globalState)).IsDisjointFrom(_756_v17));
        _757_v18 = _nw128;
        let _758_v19;
        _758_v19 = _dafny.Seq.of(!(p0));
        let _index133 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
        let _rhs153 = _758_v19;
        let _rhs154 = _743_i1;
        let _rhs155 = ((_756_v17).Intersect(_756_v17)).IsProperSubsetOf(_756_v17);
        let _lhs135 = globalState;
        let _lhs136 = _733_v1;
        let _lhs137 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
        _758_v19 = _rhs153;
        _lhs135.f0 = _rhs154;
        _lhs136[_lhs137] = _rhs155;
        (globalState).f2 = _module.__default.safeDivisionInt(((_this).f26).minus((_this).f26), (_this).f31);
        let _759_v20;
        let _nw129 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _759_v20 = _nw129;
        let _index134 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_759_v20).length));
        (_759_v20)[_index134] = (_this).f26;
        let _index135 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_759_v20).length));
        (_759_v20)[_index135] = (_this).f26;
      }
      let _760_v21;
      let _nw130 = Array((new BigNumber(12)).toNumber()).fill([]);
      _760_v21 = _nw130;
      let _761_v22;
      _761_v22 = _dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))]);
      let _762_v23;
      _762_v23 = _dafny.Seq.of((_this).f25);
      let _763_v24;
      _763_v24 = _dafny.Set.fromElements((_this).f31, (_this).f31);
      let _764_v25;
      _764_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_762_v23)).cardinality()),_763_v24);
      let _765_v26;
      _765_v26 = _module.D6.create_DC13((_this).f31, (_this).f31, (((_764_v25).contains((_this).f26)) ? ((_764_v25).get((_this).f26)) : (_dafny.Set.fromElements((_this).f26))));
      let _766_v27;
      let _nw131 = Array((new BigNumber(2)).toNumber());
      _766_v27 = _nw131;
      let _767_v28;
      _767_v28 = _dafny.Seq.of(_766_v27, _766_v27);
      let _768_v29;
      let _nw132 = Array((new BigNumber(10)).toNumber());
      _nw132[(_dafny.ZERO).toNumber()] = (_this).f26;
      _nw132[(_dafny.ONE).toNumber()] = (((_761_v22).contains(p0)) ? ((_761_v22).get(p0)) : ((_this).f26));
      _nw132[(new BigNumber(2)).toNumber()] = new BigNumber((_736_v4).length);
      _nw132[(new BigNumber(3)).toNumber()] = new BigNumber(880);
      _nw132[(new BigNumber(4)).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(5)).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(6)).toNumber()] = new BigNumber((_736_v4).length);
      _nw132[(new BigNumber(7)).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(8)).toNumber()] = (_765_v26).dtor_cf20;
      _nw132[(new BigNumber(9)).toNumber()] = new BigNumber((_767_v28).length);
      _768_v29 = _nw132;
      let _index136 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length));
      (_760_v21)[_index136] = _768_v29;
      let _index137 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length));
      (_760_v21)[_index137] = _768_v29;
      let _hi8 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_this).f26)).length);
      for (let _769_i3 = (_this).f31; _769_i3.isLessThan(_hi8); _769_i3 = _769_i3.plus(_dafny.ONE)) {
        (globalState).f0 = _769_i3;
        let _arr2 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index138 = _module.__default.safeIndex(new BigNumber(170), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        _arr2[_index138] = _769_i3;
        let _arr3 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index139 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        _arr3[_index139] = (_this).f26;
        let _770_v30;
        _770_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm5((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))], _740_v8, (_this).f26, globalState));
        let _771_v31;
        _771_v31 = _dafny.Set.fromElements(_770_v30, _770_v30, _770_v30);
        let _772_v32;
        _772_v32 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_771_v31).length)), (_this).f26, new BigNumber(-489));
        let _arr4 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index140 = _module.__default.safeIndex(new BigNumber(170), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        let _arr5 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index141 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        let _rhs156 = (_this).f26;
        let _rhs157 = (((_761_v22).contains(p0)) ? ((_761_v22).get(p0)) : ((_this).f31));
        let _rhs158 = ((_this).f31).multipliedBy(_769_i3);
        let _rhs159 = _module.__default.safeModuloInt(_769_i3, new BigNumber((_772_v32).cardinality()));
        let _lhs138 = globalState;
        let _lhs139 = globalState;
        let _lhs140 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _lhs141 = _module.__default.safeIndex(new BigNumber(170), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        let _lhs142 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _lhs143 = _module.__default.safeIndex(new BigNumber(77), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        _lhs138.f2 = _rhs156;
        _lhs139.f0 = _rhs157;
        _lhs140[_lhs141] = _rhs158;
        _lhs142[_lhs143] = _rhs159;
        let _773_v33;
        let _nw133 = new _module.C1();
        _nw133.__ctor((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))]);
        _773_v33 = _nw133;
        let _774_v35;
        _774_v35 = _dafny.Map.Empty.slice().updateUnsafe(_736_v4,_736_v4);
        let _775_v36;
        _775_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm11(_774_v35, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_776_i3, _777_v22) => function (_778_i4) {
          return _module.D9.create_DC19(_module.D6.create_DC14(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_776_i3,(_this).f31)).length), (_this).f31, _777_v22));
        })(_769_i3, _761_v22))).length), _769_i3, globalState),new BigNumber((_772_v32).cardinality()));
        let _779_v37;
        let _out17;
        _out17 = (_this).m6((function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of ((_772_v32).update((_this).f26, _module.__default.abs((_this).f31))).Elements) {
            let _780_v34 = _compr_17;
            if (((_772_v32).update((_this).f26, _module.__default.abs((_this).f31))).contains(_780_v34)) {
              _coll17.push([(_780_v34).plus(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))])[_module.__default.safeIndex(new BigNumber(170), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length))]),_769_i3]);
            }
          }
          return _coll17;
        }()).Merge(_775_v36), ((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))])[_module.__default.safeIndex(new BigNumber(170), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length))], (_this).f31, globalState);
        _779_v37 = _out17;
      }
      (globalState).f0 = (_dafny.ZERO).minus((_this).f26);
      let _781_v39;
      _781_v39 = _dafny.Seq.of(new BigNumber(801));
      let _782_v40;
      _782_v40 = _module.D0.create_DC1((_this).f26, _dafny.Seq.UnicodeFromString("ac"), (_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))], new BigNumber((_762_v23).length), _module.__default.fm5((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))], _740_v8, new BigNumber((function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (_781_v39).Elements) {
    let _783_v38 = _compr_18;
    if (_dafny.Seq.contains(_781_v39, _783_v38)) {
      _coll18.push([(_783_v38).minus(new BigNumber(619)),p0]);
    }
  }
  return _coll18;
}()).length), globalState));
      let _784_v41;
      _784_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(new BigNumber((_736_v4).length), new BigNumber((_762_v23).length), globalState),p0);
      let _pat_let_tv21 = _784_v41;
      let _pat_let_tv22 = p0;
      if (((_this).f26).isLessThanOrEqualTo(new BigNumber(((function (_pat_let18_0) {
        return function (_785_dt__update__tmp_h1) {
          return function (_pat_let19_0) {
            return function (_786_dt__update_hcf1_h0) {
              return function (_pat_let20_0) {
                return function (_787_dt__update_hcf5_h0) {
                  return function (_pat_let21_0) {
                    return function (_788_dt__update_hcf4_h0) {
                      return _module.D0.create_DC1(_786_dt__update_hcf1_h0, (_785_dt__update__tmp_h1).dtor_cf2, (_785_dt__update__tmp_h1).dtor_cf3, _788_dt__update_hcf4_h0, _787_dt__update_hcf5_h0);
                    }(_pat_let21_0);
                  }((_this).f31);
                }(_pat_let20_0);
              }(_pat_let_tv22);
            }(_pat_let19_0);
          }(new BigNumber((_pat_let_tv21).length));
        }(_pat_let18_0);
      }(_782_v40)).dtor_cf2).length))) {
        let _index142 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length));
        (_733_v1)[_index142] = !((_dafny.ZERO).minus((_this).f31)).isEqualTo((_this).f31);
        (globalState).f1 = (_this).f25;
        (globalState).f0 = _module.__default.safeModuloInt((_this).f26, ((_dafny.ZERO).minus((_this).f26)).plus((_this).f31));
        let _arr6 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index143 = _module.__default.safeIndex(new BigNumber(189), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        _arr6[_index143] = (_this).f31;
        let _arr7 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index144 = _module.__default.safeIndex(new BigNumber(189), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        _arr7[_index144] = (_this).f26;
        let _789_v42;
        _789_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber((_dafny.MultiSet.FromArray(_module.__default.fm23(globalState))).cardinality()));
        _789_v42 = (_789_v42).update(p0, (_this).f31);
      } else {
        (globalState).f0 = new BigNumber(432);
        (globalState).f0 = ((new BigNumber((_736_v4).length)).multipliedBy((_this).f26)).plus(new BigNumber(800));
        let _790_v43;
        _790_v43 = _module.D6.create_DC14(new BigNumber((_736_v4).length), (_this).f26, _761_v22);
        let _791_v44;
        _791_v44 = _module.D9.create_DC19(_790_v43);
        let _792_v45;
        _792_v45 = _module.D9.create_DC20(_791_v44);
        let _793_v46;
        _793_v46 = _module.D9.create_DC20(_792_v45);
        let _794_v47;
        _794_v47 = _dafny.Map.Empty.slice().updateUnsafe(_768_v29,_793_v46);
        let _795_v48;
        _795_v48 = _dafny.Map.Empty.slice().updateUnsafe((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))],_794_v47);
        _795_v48 = (_795_v48).update((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))], _794_v47);
        let _796_v49;
        let _nw134 = new _module.C1();
        _nw134.__ctor((_this).f25);
        _796_v49 = _nw134;
        let _arr8 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index145 = _module.__default.safeIndex(new BigNumber(294), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        _arr8[_index145] = new BigNumber((_761_v22).cardinality());
        let _arr9 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _index146 = _module.__default.safeIndex(new BigNumber(294), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        let _index147 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length));
        let _rhs160 = _796_v49;
        let _rhs161 = !(p0);
        let _rhs162 = (_this).f26;
        let _rhs163 = _768_v29;
        let _lhs144 = globalState;
        let _lhs145 = (_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))];
        let _lhs146 = _module.__default.safeIndex(new BigNumber(294), new BigNumber(((_760_v21)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length))]).length));
        let _lhs147 = _760_v21;
        let _lhs148 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_760_v21).length));
        _796_v49 = _rhs160;
        _lhs144.f1 = _rhs161;
        _lhs145[_lhs146] = _rhs162;
        _lhs147[_lhs148] = _rhs163;
        let _797_v50;
        _797_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7(globalState),(_796_v49).f32);
        let _798_v51;
        _798_v51 = _dafny.Map.Empty.slice().updateUnsafe((_781_v39)[_module.__default.safeIndex((_this).f31, new BigNumber((_781_v39).length))],_797_v50);
        let _799_v52;
        _799_v52 = _module.D9.create_DC19(_790_v43);
        let _800_v53;
        _800_v53 = _dafny.Map.Empty.slice().updateUnsafe(((!((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))])) ? (_797_v50) : ((((_798_v51).contains((_this).f31)) ? ((_798_v51).get((_this).f31)) : (_797_v50)))),_799_v52);
        _800_v53 = _800_v53;
      }
      r0 = !(_dafny.areEqual(_762_v23, _dafny.Seq.of((_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))], (_733_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_733_v1).length))])));
      let _801_v54;
      _801_v54 = _dafny.MultiSet.fromElements((_this).fm6(false, false, (_this).f31, globalState), (_this).f26, (_this).f31);
      let _pat_let_tv23 = p0;
      r1 = function (_source12) {
        let _802___mcc_h1 = _source12;
        let _803_cf18 = _802___mcc_h1;
        return _pat_let_tv23;
      }(_801_v54);
      return [r0, r1];
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _804_v0;
      _804_v0 = new _dafny.CodePoint('y'.codePointAt(0));
      _804_v0 = _804_v0;
      if (((false) ? ((_this).f25) : (_module.__default.fm5((_this).f25, _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_this).f25),p2), new BigNumber(726), globalState)))) {
        let _805_v1;
        _805_v1 = _dafny.MultiSet.fromElements(new BigNumber(-40));
        let _806_v2;
        _806_v2 = _dafny.Seq.of(p2, p2, p2, (_dafny.ZERO).minus((((_805_v1).contains(new BigNumber(7))) ? ((_805_v1).get(new BigNumber(7))) : (p2))));
        let _807_v3;
        _807_v3 = _dafny.Set.fromElements((_this).f26, (_this).f31);
        let _808_v4;
        _808_v4 = _module.D6.create_DC13((_this).f26, (_this).f26, _807_v3);
        let _809_v5;
        _809_v5 = _dafny.Seq.of(_806_v2, _806_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_810_i1) {
          return (_this).f26;
        }));
        let _811_v6;
        _811_v6 = _dafny.Seq.of((_this).f25);
        let _812_v7;
        let _nw135 = Array((new BigNumber(26)).toNumber());
        _nw135[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_806_v2, _module.__default.safeIndex((_this).f26, new BigNumber((_806_v2).length)), (_this).f26);
        _nw135[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_806_v2, _806_v2);
        _nw135[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_806_v2, _806_v2);
        _nw135[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_806_v2, _806_v2);
        _nw135[(new BigNumber(4)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(5)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(6)).toNumber()] = _module.__default.fm3(p1, false, p2, (_this).f25, globalState);
        _nw135[(new BigNumber(7)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_806_v2, _806_v2);
        _nw135[(new BigNumber(9)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(10)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(p2, p2, p1);
        _nw135[(new BigNumber(12)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(13)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(14)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(15)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_806_v2, _806_v2), _module.__default.safeIndex((_808_v4).dtor_cf20, new BigNumber((_dafny.Seq.Concat(_806_v2, _806_v2)).length)), new BigNumber(383));
        _nw135[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_806_v2, _806_v2);
        _nw135[(new BigNumber(18)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(327), (_this).f31, p2, (_this).f26), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-679)), function (_813_i0) {
          return (_this).f26;
        }));
        _nw135[(new BigNumber(20)).toNumber()] = _dafny.Seq.of((_this).f26, p1, p1);
        _nw135[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_806_v2, _806_v2), _module.__default.safeIndex(new BigNumber(296), new BigNumber((_dafny.Seq.Concat(_806_v2, _806_v2)).length)), p1);
        _nw135[(new BigNumber(22)).toNumber()] = (_809_v5)[_module.__default.safeIndex(new BigNumber((_811_v6).length), new BigNumber((_809_v5).length))];
        _nw135[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_806_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), function (_814_i2) {
          return (_this).f26;
        }));
        _nw135[(new BigNumber(24)).toNumber()] = _806_v2;
        _nw135[(new BigNumber(25)).toNumber()] = _806_v2;
        _812_v7 = _nw135;
        let _index148 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_812_v7).length));
        (_812_v7)[_index148] = _806_v2;
        let _815_v8;
        _815_v8 = _dafny.MultiSet.fromElements(true);
        let _index149 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_812_v7).length));
        (_812_v7)[_index149] = _dafny.Seq.update((_809_v5)[_module.__default.safeIndex(((_this).f31).multipliedBy((_this).f26), new BigNumber((_809_v5).length))], _module.__default.safeIndex(new BigNumber(((_815_v8).update((_this).f25, _module.__default.abs((_dafny.ZERO).minus((_this).f26)))).cardinality()), new BigNumber(((_809_v5)[_module.__default.safeIndex(((_this).f31).multipliedBy((_this).f26), new BigNumber((_809_v5).length))]).length)), (_this).f31);
        let _816_v9;
        let _out18;
        _out18 = _module.__default.m0(!((_this).f25), globalState);
        _816_v9 = _out18;
        let _817_v10;
        let _out19;
        _out19 = _module.__default.m0((_this).f25, globalState);
        _817_v10 = _out19;
        if (!((_this).f25)) {
          let _818_v11;
          _818_v11 = _dafny.Seq.UnicodeFromString("sxl");
          let _819_v12;
          let _nw136 = Array((new BigNumber(28)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = (_this).f25;
          _nw136[(_dafny.ONE).toNumber()] = !((_this).f25);
          _nw136[(new BigNumber(2)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(3)).toNumber()] = !((_this).f25);
          _nw136[(new BigNumber(4)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(5)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(6)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(7)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(8)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(9)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(10)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(11)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(12)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(13)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(14)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(15)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(16)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(17)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(18)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(19)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(20)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(21)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(22)).toNumber()] = false;
          _nw136[(new BigNumber(23)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(24)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(25)).toNumber()] = true;
          _nw136[(new BigNumber(26)).toNumber()] = (_this).f25;
          _nw136[(new BigNumber(27)).toNumber()] = (_this).f25;
          _819_v12 = _nw136;
          let _820_v13;
          let _nw137 = new _module.C0();
          _nw137.__ctor(_818_v11, (_this).f25, (_this).f26, _819_v12, !((_this).f25));
          _820_v13 = _nw137;
          let _821_v14;
          let _nw138 = new _module.C1();
          _nw138.__ctor((((_this).f25) ? (false) : ((_this).f25)));
          _821_v14 = _nw138;
          let _822_v15;
          let _nw139 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _822_v15 = _nw139;
          let _index150 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_822_v15).length));
          (_822_v15)[_index150] = (_this).f31;
          let _index151 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_822_v15).length));
          (_822_v15)[_index151] = p1;
          let _823_v16;
          _823_v16 = _dafny.Set.fromElements(_822_v15, _822_v15);
          let _824_v17;
          _824_v17 = _module.D6.create_DC12(_823_v16);
          let _825_v18;
          _825_v18 = _module.D6.create_DC12((_823_v16).Union((_824_v17).dtor_cf19));
          _824_v17 = _825_v18;
          let _index152 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_819_v12).length));
          (_819_v12)[_index152] = (((_this).f25) ? ((_821_v14).f32) : ((_this).f25));
          let _826_v19;
          _826_v19 = _dafny.Set.fromElements(_820_v13.f35);
          let _index153 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_822_v15).length));
          let _index154 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_819_v12).length));
          let _rhs164 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f31, _816_v9), (_this).f26);
          let _rhs165 = (_821_v14).f32;
          let _rhs166 = new BigNumber(727);
          let _rhs167 = ((_826_v19).Intersect(_826_v19)).IsProperSubsetOf(_826_v19);
          let _lhs149 = globalState;
          let _lhs150 = _822_v15;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_822_v15).length));
          let _lhs152 = _819_v12;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_819_v12).length));
          _817_v10 = _rhs164;
          _lhs149.f1 = _rhs165;
          _lhs150[_lhs151] = _rhs166;
          _lhs152[_lhs153] = _rhs167;
        } else {
          let _827_v20;
          _827_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_805_v1).contains(_816_v9)) ? ((_805_v1).get(_816_v9)) : (new BigNumber(-600))),(((_this).f25) ? (_811_v6) : (_811_v6)));
          _827_v20 = (function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of ((_812_v7)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_812_v7).length))]).Elements) {
              let _828_v21 = _compr_19;
              if (_dafny.Seq.contains((_812_v7)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_812_v7).length))], _828_v21)) {
                _coll19.push([(_828_v21).minus((_dafny.ZERO).minus(p2)),_811_v6]);
              }
            }
            return _coll19;
          }()).update(_module.__default.safeDivisionInt(new BigNumber(((_812_v7)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_812_v7).length))]).length), p1), _811_v6);
          (globalState).f1 = (_this).f25;
          _812_v7 = _812_v7;
          let _829_v22;
          let _init20 = function (_830_i3) {
            return (_830_i3).minus(new BigNumber(948));
          };
          let _nw140 = Array((new BigNumber(28)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw140.length); _i0_20++) {
            _nw140[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _829_v22 = _nw140;
          _829_v22 = _829_v22;
          r0 = (p1).isLessThan(p1);
        }
        r0 = !((_this).f25);
      } else {
        let _831_v23;
        let _init21 = ((_832_p2) => function (_833_i4) {
          return _module.__default.safeModuloInt(_833_i4, (_dafny.ZERO).minus(_832_p2));
        })(p2);
        let _nw141 = Array((new BigNumber(13)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw141.length); _i0_21++) {
          _nw141[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _831_v23 = _nw141;
        let _index155 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length));
        (_831_v23)[_index155] = ((_this).f31).minus(p2);
        let _index156 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length));
        (_831_v23)[_index156] = (((_this).f25) ? ((_dafny.ZERO).minus(p2)) : ((_this).f26));
        let _834_v24;
        _834_v24 = _dafny.Seq.UnicodeFromString("pwperytjt");
        let _835_v25;
        _835_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2);
        let _source13 = _module.__default.fm28(_834_v24, (_834_v24)[_module.__default.safeIndex((_this).f31, new BigNumber((_834_v24).length))], _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2)).update((_this).f25, (((_835_v25).contains((_this).f25)) ? ((_835_v25).get((_this).f25)) : ((_this).f31))))).length), _module.__default.fm2(p1, (_dafny.ZERO).minus(new BigNumber(((_835_v25).update((_this).f25, (_831_v23)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length))])).length)), globalState)), globalState);
        if (_source13.is_DC1) {
          let _836___mcc_h0 = (_source13).cf1;
          let _837___mcc_h1 = (_source13).cf2;
          let _838___mcc_h2 = (_source13).cf3;
          let _839___mcc_h3 = (_source13).cf4;
          let _840___mcc_h4 = (_source13).cf5;
          let _841_cf5 = _840___mcc_h4;
          let _842_cf4 = _839___mcc_h3;
          let _843_cf3 = _838___mcc_h2;
          let _844_cf2 = _837___mcc_h1;
          let _845_cf1 = _836___mcc_h0;
          let _846_v26;
          let _init22 = ((_847_cf5) => function (_848_i6) {
            return _847_cf5;
          })(_841_cf5);
          let _nw142 = Array((new BigNumber(27)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw142.length); _i0_22++) {
            _nw142[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _846_v26 = _nw142;
          let _849_v27;
          let _nw143 = new _module.C0();
          _nw143.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), ((_850_v0) => function (_851_i5) {
            return _850_v0;
          })(_804_v0)), _843_cf3, p2, _846_v26, (_this).f25);
          _849_v27 = _nw143;
          let _852_v28;
          _852_v28 = _dafny.Seq.of(_849_v27);
          let _853_v29;
          _853_v29 = _dafny.Set.fromElements(_843_cf3);
          let _854_v30;
          _854_v30 = _dafny.Seq.of(_853_v29);
          let _855_v31;
          _855_v31 = _dafny.Map.Empty.slice().updateUnsafe(_852_v28,_dafny.Seq.Concat(_854_v30, _dafny.Seq.of(_853_v29)));
          _855_v31 = (_855_v31).update(_852_v28, _854_v30);
          let _856_v32;
          _856_v32 = _module.D2.create_DC7(_843_cf3, false, _842_cf4);
          let _857_v33;
          _857_v33 = _dafny.Seq.of(_841_cf5);
          let _858_v34;
          _858_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_843_cf3);
          let _859_v35;
          _859_v35 = _dafny.MultiSet.fromElements(_845_cf1, (_this).f31);
          let _860_v36;
          _860_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_857_v33).length),!((((_858_v34).contains((((_859_v35).contains(new BigNumber((_859_v35).cardinality()))) ? ((_859_v35).get(new BigNumber((_859_v35).cardinality()))) : ((_this).f26)))) ? ((_858_v34).get((((_859_v35).contains(new BigNumber((_859_v35).cardinality()))) ? ((_859_v35).get(new BigNumber((_859_v35).cardinality()))) : ((_this).f26)))) : (_843_cf3))));
          let _861_v37;
          _861_v37 = _dafny.MultiSet.fromElements(_860_v36);
          let _862_v38;
          _862_v38 = _dafny.Map.Empty.slice().updateUnsafe(_856_v32,((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-923),_841_cf5))).update(_858_v34, _module.__default.abs(_845_cf1))).Difference(_861_v37));
          let _863_v39;
          _863_v39 = _dafny.Seq.of(_861_v37);
          _862_v38 = (_862_v38).update(_856_v32, (_863_v39)[_module.__default.safeIndex((_this).f26, new BigNumber((_863_v39).length))]);
          let _864_v40;
          _864_v40 = _dafny.Set.fromElements(_804_v0);
          let _865_v41;
          let _out20;
          _out20 = _module.__default.m0(!(!(_864_v40).contains(_804_v0)), globalState);
          _865_v41 = _out20;
          let _index157 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length));
          (_831_v23)[_index157] = ((_this).f26).multipliedBy((_dafny.ZERO).minus(_842_cf4));
        } else if (_source13.is_DC2) {
          let _866_v42;
          _866_v42 = _dafny.Map.Empty.slice().updateUnsafe((_831_v23)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length))],(p1).isEqualTo((_831_v23)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length))]));
          _866_v42 = (_866_v42).update((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f26, (_this).f31)), !((_this).f25));
          let _867_v43;
          _867_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_834_v24);
          let _868_v44;
          let _init23 = function (_869_i7) {
            return (_this).f25;
          };
          let _nw144 = Array((new BigNumber(26)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw144.length); _i0_23++) {
            _nw144[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _868_v44 = _nw144;
          let _870_v45;
          _870_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_868_v44);
          let _871_v46;
          let _nw145 = new _module.C0();
          _nw145.__ctor(_834_v24, !(false), new BigNumber((_867_v43).length), (((_870_v45).contains((_this).f25)) ? ((_870_v45).get((_this).f25)) : (_868_v44)), (_this).f25);
          _871_v46 = _nw145;
          let _872_v47;
          _872_v47 = _dafny.Seq.of(_871_v46, _871_v46);
          let _873_v48;
          _873_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_872_v47);
          let _874_v49;
          _874_v49 = _module.D9.create_DC18(_873_v48);
          let _pat_let_tv24 = _873_v48;
          let _875_v50;
          let _nw146 = Array((new BigNumber(29)).toNumber());
          _nw146[(_dafny.ZERO).toNumber()] = _874_v49;
          _nw146[(_dafny.ONE).toNumber()] = _874_v49;
          _nw146[(new BigNumber(2)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(3)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(4)).toNumber()] = _module.D9.create_DC18(_873_v48);
          _nw146[(new BigNumber(5)).toNumber()] = _module.D9.create_DC18(_873_v48);
          _nw146[(new BigNumber(6)).toNumber()] = function (_pat_let22_0) {
            return function (_876_dt__update__tmp_h0) {
              return function (_pat_let23_0) {
                return function (_877_dt__update_hcf30_h0) {
                  return _module.D9.create_DC18(_877_dt__update_hcf30_h0);
                }(_pat_let23_0);
              }(_pat_let_tv24);
            }(_pat_let22_0);
          }(_874_v49);
          _nw146[(new BigNumber(7)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(8)).toNumber()] = _module.D9.create_DC18(_873_v48);
          _nw146[(new BigNumber(9)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(10)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(11)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(12)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(13)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(14)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(15)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(16)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(17)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(18)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(19)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(20)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(21)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(22)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(23)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(24)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(25)).toNumber()] = _module.D9.create_DC18(_873_v48);
          _nw146[(new BigNumber(26)).toNumber()] = _module.D9.create_DC18(_873_v48);
          _nw146[(new BigNumber(27)).toNumber()] = _874_v49;
          _nw146[(new BigNumber(28)).toNumber()] = _874_v49;
          _875_v50 = _nw146;
          let _878_v51;
          let _nw147 = Array((new BigNumber(25)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = _875_v50;
          _nw147[(_dafny.ONE).toNumber()] = _875_v50;
          _nw147[(new BigNumber(2)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(3)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(4)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(5)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(6)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(7)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(8)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(9)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(10)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(11)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(12)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(13)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(14)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(15)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(16)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(17)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(18)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(19)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(20)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(21)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(22)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(23)).toNumber()] = _875_v50;
          _nw147[(new BigNumber(24)).toNumber()] = _875_v50;
          _878_v51 = _nw147;
          let _index158 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_878_v51).length));
          (_878_v51)[_index158] = _875_v50;
          let _index159 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_878_v51).length));
          let _nw148 = Array((new BigNumber(3)).toNumber()).fill(_module.D9.Default());
          (_878_v51)[_index159] = _nw148;
          let _879_v52;
          let _nw149 = new _module.C1();
          _nw149.__ctor((_this).f25);
          _879_v52 = _nw149;
          _879_v52 = _879_v52;
          let _880_v53;
          _880_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(956),(_834_v24)[_module.__default.safeIndex((_this).f31, new BigNumber((_834_v24).length))]);
          let _881_v54;
          _881_v54 = _module.D4.create_DC10(_804_v0, _804_v0);
          _880_v53 = (_880_v53).update(p1, (_881_v54).dtor_cf16);
        } else {
          let _882___mcc_h5 = (_source13).cf0;
          let _883_cf0 = _882___mcc_h5;
          let _884_v55;
          let _out21;
          _out21 = _module.__default.m0(!(true), globalState);
          _884_v55 = _out21;
          (globalState).f0 = _module.__default.safeModuloInt((_this).f26, (_this).f26);
          let _885_v56;
          _885_v56 = _dafny.Seq.of(_dafny.Seq.of((_this).f25, false, (_this).f25));
          let _886_v57;
          let _nw150 = Array((new BigNumber(23)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _883_cf0;
          _nw150[(_dafny.ONE).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(2)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(3)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(4)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(5)).toNumber()] = !(_883_cf0);
          _nw150[(new BigNumber(6)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(7)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(8)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(9)).toNumber()] = !(p1).isEqualTo((_831_v23)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length))]);
          _nw150[(new BigNumber(10)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(11)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(12)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(13)).toNumber()] = false;
          _nw150[(new BigNumber(14)).toNumber()] = !(_883_cf0);
          _nw150[(new BigNumber(15)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(16)).toNumber()] = (_this).f25;
          _nw150[(new BigNumber(17)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(18)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(19)).toNumber()] = true;
          _nw150[(new BigNumber(20)).toNumber()] = (p0).contains(new BigNumber((_dafny.Set.fromElements(new BigNumber((_885_v56).length))).length));
          _nw150[(new BigNumber(21)).toNumber()] = _883_cf0;
          _nw150[(new BigNumber(22)).toNumber()] = _883_cf0;
          _886_v57 = _nw150;
          let _index160 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_886_v57).length));
          (_886_v57)[_index160] = (_this).f25;
          let _index161 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_886_v57).length));
          (_886_v57)[_index161] = _dafny.areEqual(_834_v24, _834_v24);
          let _887_v58;
          _887_v58 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f25);
          let _888_v59;
          let _nw151 = new _module.C1();
          _nw151.__ctor(_883_cf0);
          _888_v59 = _nw151;
          let _889_v60;
          _889_v60 = _dafny.Seq.of(_888_v59, _888_v59, _888_v59);
          let _890_v61;
          _890_v61 = _dafny.Seq.of(_889_v60);
          _887_v58 = ((_887_v58).update((_this).f25, true)).update(((_this).f26).isEqualTo(new BigNumber((_890_v61).length)), !_dafny.areEqual(_834_v24, _834_v24));
        }
        (globalState).f1 = true;
        (globalState).f1 = (((_this).f25) ? ((_this).f25) : (false));
        let _index162 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_831_v23).length));
        (_831_v23)[_index162] = (_this).f26;
      }
      let _891_i8;
      _891_i8 = _dafny.ZERO;
      L3: {
        while (!((_this).f25)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_891_i8)) {
              break L3;
            }
            _891_i8 = (_891_i8).plus(_dafny.ONE);
            let _892_v62;
            _892_v62 = _module.D0.create_DC2();
            let _893_v63;
            _893_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_892_v62);
            let _894_v64;
            _894_v64 = _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber((_893_v63).length), (_this).f31));
            (globalState).f0 = (_894_v64)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_894_v64).length))];
            let _895_v65;
            _895_v65 = _dafny.Set.fromElements(_module.__default.fm2(new BigNumber(-334), (_this).f31, globalState), p1);
            if (!(!((_895_v65).IsDisjointFrom(_895_v65)))) {
              let _896_v66;
              _896_v66 = _dafny.Seq.of((_this).f25, (_this).f25, (_this).f25, (_this).f25);
              let _897_v67;
              _897_v67 = _dafny.Seq.UnicodeFromString("ynxpxku");
              let _898_v68;
              _898_v68 = _dafny.Map.Empty.slice().updateUnsafe(p2,_897_v67);
              let _899_v69;
              _899_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_898_v68).contains(new BigNumber(-3))) ? ((_898_v68).get(new BigNumber(-3))) : (_dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), function (_900_i9) {
                return new _dafny.CodePoint('n'.codePointAt(0));
              })))).length),(_this).f25);
              let _901_v70;
              _901_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2);
              let _902_v71;
              _902_v71 = _dafny.Seq.of(_901_v70, _901_v70);
              let _903_v72;
              _903_v72 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_896_v66, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_899_v69).length)), new BigNumber((_896_v66).length)), (_this).f25), _module.__default.safeIndex(new BigNumber(807), new BigNumber((_dafny.Seq.update(_896_v66, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_899_v69).length)), new BigNumber((_896_v66).length)), (_this).f25)).length)), (_this).f25), _dafny.Seq.update(_896_v66, _module.__default.safeIndex(new BigNumber((_902_v71).length), new BigNumber((_896_v66).length)), (_this).f25));
              _903_v72 = _903_v72;
              let _904_v73;
              _904_v73 = _dafny.MultiSet.fromElements(_804_v0, _804_v0, (((_this).f25) ? (_804_v0) : (_804_v0)));
              let _905_v74;
              _905_v74 = _module.D0.create_DC1((_this).f31, _897_v67, (_this).f25, (_this).f31, (_this).f25);
              let _906_v76;
              _906_v76 = _dafny.Set.fromElements((_this).f25, true);
              let _907_v77;
              let _nw152 = Array((new BigNumber(13)).toNumber());
              _nw152[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(p2, p1);
              _nw152[(_dafny.ONE).toNumber()] = (((_this).f25) ? ((_this).f31) : (p1));
              _nw152[(new BigNumber(2)).toNumber()] = p1;
              _nw152[(new BigNumber(3)).toNumber()] = (_this).f26;
              _nw152[(new BigNumber(4)).toNumber()] = (_905_v74).dtor_cf4;
              _nw152[(new BigNumber(5)).toNumber()] = new BigNumber(22);
              _nw152[(new BigNumber(6)).toNumber()] = ((_this).f26).multipliedBy((_this).f31);
              _nw152[(new BigNumber(7)).toNumber()] = p2;
              _nw152[(new BigNumber(8)).toNumber()] = new BigNumber((function () {
                let _coll20 = new _dafny.Map();
                for (const _compr_20 of (p0).Keys.Elements) {
                  let _908_v75 = _compr_20;
                  if ((p0).contains(_908_v75)) {
                    _coll20.push([(_908_v75).multipliedBy((_this).f26),!((_this).f25)]);
                  }
                }
                return _coll20;
              }()).length);
              _nw152[(new BigNumber(9)).toNumber()] = (new BigNumber((_894_v64).length)).multipliedBy(p2);
              _nw152[(new BigNumber(10)).toNumber()] = (_this).f31;
              _nw152[(new BigNumber(11)).toNumber()] = (new BigNumber((_906_v76).length)).minus((_this).f31);
              _nw152[(new BigNumber(12)).toNumber()] = (_this).f31;
              _907_v77 = _nw152;
              let _index163 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_907_v77).length));
              (_907_v77)[_index163] = p1;
              let _909_v78;
              _909_v78 = _module.D2.create_DC7((_this).f25, (_this).f25, p1);
              let _index164 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_907_v77).length));
              let _rhs168 = _904_v73;
              let _rhs169 = ((_909_v78).dtor_cf13).minus(p1);
              let _rhs170 = (_this).f26;
              let _rhs171 = !(!(((!((_this).f25)) ? ((_this).f25) : ((_this).f25)))) || ((_this).f25);
              let _rhs172 = (_this).fm7(globalState);
              let _lhs154 = globalState;
              let _lhs155 = _907_v77;
              let _lhs156 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_907_v77).length));
              let _lhs157 = globalState;
              _904_v73 = _rhs168;
              _lhs154.f2 = _rhs169;
              _lhs155[_lhs156] = _rhs170;
              _lhs157.f1 = _rhs171;
              r0 = _rhs172;
              let _910_v79;
              _910_v79 = _dafny.Seq.of(_897_v67, _897_v67);
              _897_v67 = _dafny.Seq.Concat(_897_v67, _dafny.Seq.Concat(_897_v67, _dafny.Seq.update((_910_v79)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-414)), new BigNumber((_910_v79).length))], _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_897_v67, _module.__default.safeIndex(new BigNumber((_897_v67).length), new BigNumber((_897_v67).length)), _804_v0)).length), new BigNumber(((_910_v79)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-414)), new BigNumber((_910_v79).length))]).length)), _804_v0)));
              let _911_v80;
              _911_v80 = _dafny.MultiSet.fromElements((_this).f25);
              let _rhs173 = ((_911_v80).update((_this).f25, _module.__default.abs(p2))).Difference(_911_v80);
              let _rhs174 = !(!((_this).f25));
              let _rhs175 = (_906_v76).Intersect(_906_v76);
              let _lhs158 = globalState;
              _911_v80 = _rhs173;
              _lhs158.f1 = _rhs174;
              _906_v76 = _rhs175;
              let _912_v81;
              _912_v81 = _module.D6.create_DC13(new BigNumber(990), p1, _dafny.Set.fromElements((_this).f26, (_907_v77)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_907_v77).length))], new BigNumber((_897_v67).length)));
              let _913_v82;
              _913_v82 = _dafny.Map.Empty.slice().updateUnsafe(_804_v0,(_dafny.ZERO).minus((_912_v81).dtor_cf21));
              _913_v82 = _913_v82;
            } else {
              let _914_v83;
              _914_v83 = _dafny.Seq.UnicodeFromString("kg");
              let _915_v84;
              _915_v84 = _module.D0.create_DC0((_this).f25);
              let _916_v85;
              _916_v85 = _dafny.Map.Empty.slice().updateUnsafe(_915_v84,(_dafny.ZERO).minus((_this).f31));
              let _917_v86;
              let _nw153 = Array((new BigNumber(20)).toNumber()).fill(false);
              _917_v86 = _nw153;
              let _918_v87;
              let _nw154 = new _module.C0();
              _nw154.__ctor(_914_v83, _module.__default.fm5(true, _916_v85, new BigNumber(814), globalState), (_this).f26, _917_v86, false);
              _918_v87 = _nw154;
              let _919_v88;
              _919_v88 = _dafny.Seq.of(false);
              let _index165 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_917_v86).length));
              (_917_v86)[_index165] = (_919_v88)[_module.__default.safeIndex(p2, new BigNumber((_919_v88).length))];
              let _920_v89;
              let _init24 = ((_921_v64) => function (_922_i10) {
                return _dafny.Map.Empty.slice().updateUnsafe(_921_v64,true);
              })(_894_v64);
              let _nw155 = Array((new BigNumber(10)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw155.length); _i0_24++) {
                _nw155[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _920_v89 = _nw155;
              let _923_v91;
              _923_v91 = _dafny.Seq.of(_894_v64);
              let _index166 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_920_v89).length));
              (_920_v89)[_index166] = function () {
                let _coll21 = new _dafny.Map();
                for (const _compr_21 of (_923_v91).Elements) {
                  let _924_v90 = _compr_21;
                  if (_dafny.Seq.contains(_923_v91, _924_v90)) {
                    _coll21.push([_924_v90,(_this).f25]);
                  }
                }
                return _coll21;
              }();
              let _925_v92;
              _925_v92 = _dafny.Set.fromElements(_918_v87.f34);
              let _926_v93;
              _926_v93 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of(p2, new BigNumber((_914_v83).length)));
              let _index167 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_917_v86).length));
              let _index168 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_920_v89).length));
              let _rhs176 = (_this).f31;
              let _rhs177 = _918_v87;
              let _rhs178 = (_this).f25;
              let _rhs179 = ((((_this).f25) ? (p2) : ((_module.D2.create_DC7(true, _918_v87.f34, new BigNumber((_925_v92).length))).dtor_cf13))).multipliedBy(p2);
              let _rhs180 = (_dafny.Map.Empty.slice().updateUnsafe((((_926_v93).contains(new BigNumber(-64))) ? ((_926_v93).get(new BigNumber(-64))) : (_894_v64)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_894_v64,(_this).f25));
              let _lhs159 = globalState;
              let _lhs160 = _917_v86;
              let _lhs161 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_917_v86).length));
              let _lhs162 = globalState;
              let _lhs163 = _920_v89;
              let _lhs164 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_920_v89).length));
              _lhs159.f0 = _rhs176;
              _918_v87 = _rhs177;
              _lhs160[_lhs161] = _rhs178;
              _lhs162.f2 = _rhs179;
              _lhs163[_lhs164] = _rhs180;
              let _927_v94;
              _927_v94 = _dafny.Map.Empty.slice().updateUnsafe(_914_v83,_914_v83);
              _927_v94 = _927_v94;
              let _928_v95;
              let _nw156 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
              _928_v95 = _nw156;
              let _index169 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_928_v95).length));
              (_928_v95)[_index169] = p2;
              let _index170 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_928_v95).length));
              (_928_v95)[_index170] = (_this).f31;
              let _929_v96;
              _929_v96 = _module.D1.create_DC4((_917_v86)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_917_v86).length))], true);
              (globalState).f0 = _module.__default.fm2((((_929_v96).dtor_cf7) ? ((_928_v95)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_928_v95).length))]) : (p1)), (_928_v95)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_928_v95).length))], globalState);
              let _930_v97;
              _930_v97 = _dafny.Set.fromElements(_895_v65);
              let _index171 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_917_v86).length));
              (_917_v86)[_index171] = (_930_v97).IsDisjointFrom(_930_v97);
            }
            let _931_v98;
            _931_v98 = _dafny.Seq.UnicodeFromString("kbigaspv");
            let _932_v99;
            _932_v99 = _dafny.Set.fromElements((_this).f25);
            let _933_v100;
            _933_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_931_v98).length),_932_v99);
            let _934_v101;
            _934_v101 = _dafny.Seq.of((_this).f25);
            _933_v100 = (_933_v100).update(new BigNumber((_dafny.MultiSet.fromElements(_931_v98)).cardinality()), (_dafny.Set.fromElements((_934_v101)[_module.__default.safeIndex(p1, new BigNumber((_934_v101).length))], (_this).f25)).Difference(_module.__default.fm29((_this).fm7(globalState), p0, _804_v0, (_this).f26, globalState)));
            _932_v99 = _932_v99;
          }
        }
      }
      let _935_v102;
      let _out22;
      _out22 = _module.__default.m0((_this).f25, globalState);
      _935_v102 = _out22;
      let _936_v103;
      _936_v103 = _dafny.MultiSet.fromElements(_module.__default.fm2(new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality()), (_this).f31, globalState));
      let _937_v104;
      _937_v104 = _dafny.Seq.of((_this).f26);
      let _938_v105;
      _938_v105 = _dafny.MultiSet.fromElements((_937_v104)[_module.__default.safeIndex((_this).f26, new BigNumber((_937_v104).length))]);
      let _939_v106;
      let _nw157 = Array((new BigNumber(17)).toNumber());
      _nw157[(_dafny.ZERO).toNumber()] = _936_v103;
      _nw157[(_dafny.ONE).toNumber()] = _936_v103;
      _nw157[(new BigNumber(2)).toNumber()] = _938_v105;
      _nw157[(new BigNumber(3)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(4)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(5)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(6)).toNumber()] = _module.__default.fm30(globalState);
      _nw157[(new BigNumber(7)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(8)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(9)).toNumber()] = _module.__default.fm30(globalState);
      _nw157[(new BigNumber(10)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(11)).toNumber()] = _938_v105;
      _nw157[(new BigNumber(12)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(13)).toNumber()] = (_938_v105).update(p1, _module.__default.abs((_this).f26));
      _nw157[(new BigNumber(14)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(15)).toNumber()] = _936_v103;
      _nw157[(new BigNumber(16)).toNumber()] = _936_v103;
      _939_v106 = _nw157;
      let _940_v107;
      _940_v107 = _dafny.Seq.UnicodeFromString("nnxuic");
      let _941_v108;
      _941_v108 = _module.D0.create_DC1(_935_v102, _940_v107, true, p1, (_this).f25);
      let _942_v109;
      _942_v109 = _dafny.MultiSet.fromElements(true, (_this).f25, false);
      let _rhs181 = _939_v106;
      let _rhs182 = (_941_v108).dtor_cf4;
      let _rhs183 = (_942_v109).IsSubsetOf(_942_v109);
      let _rhs184 = _dafny.Seq.Concat(_937_v104, _dafny.Seq.Concat(_937_v104, _937_v104));
      let _lhs165 = globalState;
      let _lhs166 = globalState;
      _939_v106 = _rhs181;
      _lhs165.f2 = _rhs182;
      _lhs166.f1 = _rhs183;
      _937_v104 = _rhs184;
      let _943_v110;
      let _nw158 = Array((new BigNumber(15)).toNumber()).fill(false);
      _943_v110 = _nw158;
      let _index172 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_943_v110).length));
      (_943_v110)[_index172] = (_this).f25;
      let _944_v111;
      _944_v111 = _module.D0.create_DC0((_this).f25);
      let _index173 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_943_v110).length));
      (_943_v110)[_index173] = function (_source14) {
        if (_source14.is_DC1) {
          let _945___mcc_h6 = (_source14).cf1;
          let _946___mcc_h7 = (_source14).cf2;
          let _947___mcc_h8 = (_source14).cf3;
          let _948___mcc_h9 = (_source14).cf4;
          let _949___mcc_h10 = (_source14).cf5;
          let _950_cf5 = _949___mcc_h10;
          let _951_cf4 = _948___mcc_h9;
          let _952_cf3 = _947___mcc_h8;
          let _953_cf2 = _946___mcc_h7;
          let _954_cf1 = _945___mcc_h6;
          return (_this).f25;
        } else if (_source14.is_DC2) {
          return (_this).f25;
        } else {
          let _955___mcc_h11 = (_source14).cf0;
          let _956_cf0 = _955___mcc_h11;
          return _956_cf0;
        }
      }(_944_v111);
      r0 = (_943_v110)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_943_v110).length))];
      return r0;
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f25 = false;
      this._f26 = _dafny.ZERO;
      this._f29 = _dafny.Seq.UnicodeFromString("");
      this._f30 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f29, f30, f25, f26) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f26).multipliedBy(((_this).f26).minus((_this).f26));
    };
    fm7(globalState) {
      let _this = this;
      return !(new BigNumber(926)).isEqualTo((_this).f26);
    };
    fm9(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this).f26);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _957_v0;
      _957_v0 = new _dafny.CodePoint('u'.codePointAt(0));
      _957_v0 = (_module.__default.fm10((_this).f25, p2, globalState)).dtor_cf6;
      let _958_v1;
      _958_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,p1);
      if ((((((((_958_v1).contains(true)) ? ((_958_v1).get(true)) : ((_this).f25))) ? (true) : ((_this).f30))) ? ((_this).f25) : (p2))) {
        let _959_v2;
        _959_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,!((_this).f30) || ((_this).f30));
        _959_v2 = (_959_v2).update((_this).f26, (_this).f25);
        let _960_v3;
        let _nw159 = Array((new BigNumber(6)).toNumber()).fill(false);
        _960_v3 = _nw159;
        let _961_v4;
        _961_v4 = _dafny.Seq.of(_960_v3);
        let _962_v5;
        let _nw160 = Array((new BigNumber(26)).toNumber());
        _nw160[(_dafny.ZERO).toNumber()] = _960_v3;
        _nw160[(_dafny.ONE).toNumber()] = _960_v3;
        _nw160[(new BigNumber(2)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(3)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(4)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(5)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(6)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(7)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(8)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(9)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(10)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(11)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(12)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(13)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(14)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(15)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(16)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(17)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(18)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(19)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(20)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(21)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(22)).toNumber()] = _960_v3;
        _nw160[(new BigNumber(23)).toNumber()] = (((_this).f30) ? (_960_v3) : (_960_v3));
        _nw160[(new BigNumber(24)).toNumber()] = (_961_v4)[_module.__default.safeIndex((_this).f26, new BigNumber((_961_v4).length))];
        _nw160[(new BigNumber(25)).toNumber()] = (_961_v4)[_module.__default.safeIndex((_this).f26, new BigNumber((_961_v4).length))];
        _962_v5 = _nw160;
        let _963_v6;
        _963_v6 = _dafny.Set.fromElements((_this).f30, (_this).f25);
        let _964_v7;
        _964_v7 = _module.D1.create_DC5(_module.D1.create_DC3(_957_v0));
        let _965_v8;
        _965_v8 = _module.D1.create_DC5(_964_v7);
        let _pat_let_tv25 = _964_v7;
        let _966_v9;
        _966_v9 = _dafny.MultiSet.fromElements(_module.D1.create_DC5(_964_v7), function (_pat_let24_0) {
          return function (_967_dt__update__tmp_h0) {
            return function (_pat_let25_0) {
              return function (_968_dt__update_hcf9_h0) {
                return _module.D1.create_DC5(_968_dt__update_hcf9_h0);
              }(_pat_let25_0);
            }(_pat_let_tv25);
          }(_pat_let24_0);
        }(_965_v8), _965_v8);
        let _rhs185 = !((_963_v6).IsDisjointFrom(_963_v6)) || (p2);
        let _rhs186 = (_this).f26;
        let _rhs187 = _962_v5;
        let _rhs188 = (((_959_v2).contains((((_966_v9).contains(_965_v8)) ? ((_966_v9).get(_965_v8)) : ((_this).f26)))) ? ((_959_v2).get((((_966_v9).contains(_965_v8)) ? ((_966_v9).get(_965_v8)) : ((_this).f26)))) : (((_this).f25) || (p1)));
        let _lhs167 = globalState;
        let _lhs168 = globalState;
        r0 = _rhs185;
        _lhs167.f2 = _rhs186;
        _962_v5 = _rhs187;
        _lhs168.f1 = _rhs188;
        let _969_v10;
        let _init25 = function (_970_i0) {
          return _dafny.Seq.of(true, true);
        };
        let _nw161 = Array((new BigNumber(8)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw161.length); _i0_25++) {
          _nw161[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _969_v10 = _nw161;
        let _971_v11;
        _971_v11 = _dafny.Seq.of(p1);
        let _972_v12;
        _972_v12 = _module.D0.create_DC0((_this).f30);
        let _973_v13;
        _973_v13 = _dafny.Map.Empty.slice().updateUnsafe(_972_v12,(_this).f26);
        let _index174 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_969_v10).length));
        (_969_v10)[_index174] = _dafny.Seq.update(_971_v11, _module.__default.safeIndex((_this).f26, new BigNumber((_971_v11).length)), _module.__default.fm5((_this).f30, _973_v13, (_this).f26, globalState));
        let _index175 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_969_v10).length));
        (_969_v10)[_index175] = _dafny.Seq.of(p1);
        let _974_v14;
        _974_v14 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(689)), ((_975_v0) => function (_976_i1) {
  return _975_v0;
})(_957_v0))).length), (_this).f29, (_this).f30, (_this).f26, p2);
        _957_v0 = _module.__default.fm4(_974_v14, (_this).f25, (_dafny.ZERO).minus((_this).f26), globalState);
        (globalState).f0 = (_this).f26;
      } else {
        let _977_v15;
        _977_v15 = _dafny.Seq.of(p2);
        let _978_v16;
        let _nw162 = new _module.C2();
        _nw162.__ctor((_dafny.ZERO).minus((_this).f26), !(!((_this).f25)), ((p2) ? ((_this).f26) : (new BigNumber((_dafny.Seq.of(new BigNumber((_977_v15).length), new BigNumber((p0).cardinality()))).length))));
        _978_v16 = _nw162;
        _978_v16 = ((p1) ? (_978_v16) : (_978_v16));
        _977_v15 = _977_v15;
        (globalState).f1 = (_module.D0.create_DC1((_dafny.ZERO).minus((_978_v16).f26), (_this).f29, true, (_978_v16).f26, !((((_958_v1).contains(true)) ? ((_958_v1).get(true)) : (p1))))).dtor_cf5;
        let _rhs189 = (_this).f30;
        let _rhs190 = (_dafny.ZERO).minus((_this).f26);
        let _lhs169 = globalState;
        let _lhs170 = globalState;
        _lhs169.f1 = _rhs189;
        _lhs170.f2 = _rhs190;
        let _979_v17;
        _979_v17 = _dafny.Set.fromElements(_module.__default.fm2((_this).f26, (_this).f26, globalState), (_this).f26);
        let _980_v18;
        _980_v18 = _module.D6.create_DC13((_dafny.ZERO).minus((_this).f26), (_978_v16).f26, _979_v17);
        let _981_v19;
        _981_v19 = _dafny.Map.Empty.slice().updateUnsafe((_980_v18).dtor_cf21,(_this).f26);
        _981_v19 = _981_v19;
      }
      let _982_i2;
      _982_i2 = _dafny.ZERO;
      L4: {
        while (!(((_this).f26).isEqualTo(((_this).f26).minus((_this).f26)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_982_i2)) {
              break L4;
            }
            _982_i2 = (_982_i2).plus(_dafny.ONE);
            let _983_v20;
            let _nw163 = Array((new BigNumber(28)).toNumber()).fill(false);
            _983_v20 = _nw163;
            let _index176 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length));
            (_983_v20)[_index176] = (_this).f25;
            let _984_v21;
            _984_v21 = _dafny.Seq.of(false);
            let _985_v22;
            _985_v22 = _dafny.Seq.of(_984_v21, _984_v21);
            let _986_v23;
            _986_v23 = _dafny.MultiSet.fromElements((_985_v22)[_module.__default.safeIndex((_this).f26, new BigNumber((_985_v22).length))], _984_v21, _984_v21);
            let _index177 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length));
            (_983_v20)[_index177] = !(!(_986_v23).contains(_984_v21));
            if (!(p1)) {
              let _987_v24;
              _987_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,((_this).f26).plus((_this).f26));
              let _988_v25;
              let _nw164 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
              _988_v25 = _nw164;
              let _index178 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_988_v25).length));
              (_988_v25)[_index178] = _module.__default.fm3(new BigNumber(973), (_983_v20)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length))], (_this).f26, (_this).f30, globalState);
              let _989_v26;
              _989_v26 = _dafny.Seq.of((_this).f26, new BigNumber((_dafny.Seq.of((_this).f26)).length));
              let _index179 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_988_v25).length));
              let _rhs191 = _987_v24;
              let _rhs192 = _989_v26;
              let _lhs171 = _988_v25;
              let _lhs172 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_988_v25).length));
              _987_v24 = _rhs191;
              _lhs171[_lhs172] = _rhs192;
              let _990_v27;
              let _nw165 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _990_v27 = _nw165;
              let _index180 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_990_v27).length));
              (_990_v27)[_index180] = _957_v0;
              let _991_v28;
              _991_v28 = _dafny.Set.fromElements(((_this).f26).isLessThanOrEqualTo((_this).f26));
              let _pat_let_tv26 = _957_v0;
              let _index181 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_990_v27).length));
              let _rhs193 = _983_v20;
              let _rhs194 = _dafny.Seq.Concat(_984_v21, _984_v21);
              let _rhs195 = (function (_pat_let26_0) {
                return function (_992_dt__update__tmp_h1) {
                  return function (_pat_let27_0) {
                    return function (_993_dt__update_hcf17_h0) {
                      return _module.D4.create_DC10((_992_dt__update__tmp_h1).dtor_cf16, _993_dt__update_hcf17_h0);
                    }(_pat_let27_0);
                  }(_pat_let_tv26);
                }(_pat_let26_0);
              }(_module.D4.create_DC10(_957_v0, _957_v0))).dtor_cf17;
              let _rhs196 = (_this).f26;
              let _rhs197 = _991_v28;
              let _lhs173 = _990_v27;
              let _lhs174 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_990_v27).length));
              let _lhs175 = globalState;
              _983_v20 = _rhs193;
              _984_v21 = _rhs194;
              _lhs173[_lhs174] = _rhs195;
              _lhs175.f0 = _rhs196;
              _991_v28 = _rhs197;
              let _994_v29;
              _994_v29 = _dafny.Map.Empty.slice().updateUnsafe(_983_v20,(_983_v20)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length))]);
              let _995_v30;
              _995_v30 = _dafny.Seq.of(_994_v29, _994_v29, _994_v29, _994_v29, _994_v29);
              let _996_v31;
              let _nw166 = new _module.C2();
              _nw166.__ctor((_this).f26, !((_this).f25) || (false), (new BigNumber(((_995_v30)[_module.__default.safeIndex((_this).f26, new BigNumber((_995_v30).length))]).length)).multipliedBy(new BigNumber((_984_v21).length)));
              _996_v31 = _nw166;
              _957_v0 = (_990_v27)[_module.__default.safeIndex(new BigNumber(840), new BigNumber((_990_v27).length))];
              let _index182 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_988_v25).length));
              (_988_v25)[_index182] = _dafny.Seq.Concat(((p1) ? (_989_v26) : ((_988_v25)[_module.__default.safeIndex(new BigNumber(587), new BigNumber((_988_v25).length))])), (_988_v25)[_module.__default.safeIndex(new BigNumber(587), new BigNumber((_988_v25).length))]);
            } else {
              let _997_v32;
              _997_v32 = _module.D0.create_DC0(p2);
              let _998_v33;
              _998_v33 = _dafny.Map.Empty.slice().updateUnsafe(_997_v32,new BigNumber(649));
              let _999_v34;
              _999_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(326),(_dafny.ZERO).minus(new BigNumber(((_this).f29).length)));
              let _1000_v35;
              _1000_v35 = _dafny.MultiSet.fromElements((_983_v20)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length))], _module.__default.fm5((_this).f25, _998_v33, (((_999_v34).contains((_this).f26)) ? ((_999_v34).get((_this).f26)) : ((_this).f26)), globalState));
              (globalState).f1 = !((_1000_v35).contains(true));
              let _1001_v36;
              let _nw167 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
              _1001_v36 = _nw167;
              let _1002_v37;
              _1002_v37 = _dafny.Seq.of(_1001_v36);
              _1001_v36 = ((_module.__default.fm5(p1, _998_v33, (_this).f26, globalState)) ? (_1001_v36) : ((_1002_v37)[_module.__default.safeIndex((_this).f26, new BigNumber((_1002_v37).length))]));
              let _1003_v38;
              let _nw168 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
              _1003_v38 = _nw168;
              let _1004_v39;
              _1004_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_1000_v35);
              let _1005_v40;
              _1005_v40 = _module.D6.create_DC14((_this).f26, (_this).f26, (((_1004_v39).contains(new BigNumber(796))) ? ((_1004_v39).get(new BigNumber(796))) : ((_1000_v35).update((_this).f25, _module.__default.abs((_this).f26)))));
              let _1006_v41;
              _1006_v41 = _module.D9.create_DC19(_1005_v40);
              let _1007_v42;
              _1007_v42 = _module.D9.create_DC20(_1006_v41);
              let _1008_v43;
              _1008_v43 = _module.D0.create_DC1(new BigNumber(-977), (_this).f29, p1, (_this).f26, (_984_v21)[_module.__default.safeIndex((_this).f26, new BigNumber((_984_v21).length))]);
              let _pat_let_tv27 = p1;
              let _pat_let_tv28 = p1;
              let _1009_v44;
              _1009_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1007_v42,_dafny.Set.fromElements(_957_v0, _module.__default.fm4(function (_pat_let28_0) {
                return function (_1010_dt__update__tmp_h2) {
                  return function (_pat_let29_0) {
                    return function (_1011_dt__update_hcf3_h0) {
                      return function (_pat_let30_0) {
                        return function (_1012_dt__update_hcf4_h0) {
                          return function (_pat_let31_0) {
                            return function (_1013_dt__update_hcf5_h0) {
                              return _module.D0.create_DC1((_1010_dt__update__tmp_h2).dtor_cf1, (_1010_dt__update__tmp_h2).dtor_cf2, _1011_dt__update_hcf3_h0, _1012_dt__update_hcf4_h0, _1013_dt__update_hcf5_h0);
                            }(_pat_let31_0);
                          }(_pat_let_tv28);
                        }(_pat_let30_0);
                      }((_this).f26);
                    }(_pat_let29_0);
                  }(_pat_let_tv27);
                }(_pat_let28_0);
              }(_1008_v43), (_983_v20)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length))], (_this).f26, globalState), _957_v0, _957_v0));
              let _1014_v45;
              _1014_v45 = _dafny.Set.fromElements(_957_v0, _957_v0);
              let _pat_let_tv29 = _1007_v42;
              let _pat_let_tv30 = _1007_v42;
              let _index183 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1003_v38).length));
              (_1003_v38)[_index183] = (((_1009_v44).contains(function (_pat_let34_0) {
                return function (_1017_dt__update__tmp_h3) {
                  return function (_pat_let35_0) {
                    return function (_1018_dt__update_hcf32_h0) {
                      return _module.D9.create_DC20(_1018_dt__update_hcf32_h0);
                    }(_pat_let35_0);
                  }((_pat_let_tv30).dtor_cf32);
                }(_pat_let34_0);
              }(_1007_v42))) ? ((_1009_v44).get(function (_pat_let32_0) {
                return function (_1015_dt__update__tmp_h4) {
                  return function (_pat_let33_0) {
                    return function (_1016_dt__update_hcf32_h1) {
                      return _module.D9.create_DC20(_1016_dt__update_hcf32_h1);
                    }(_pat_let33_0);
                  }((_pat_let_tv29).dtor_cf32);
                }(_pat_let32_0);
              }(_1007_v42))) : (_1014_v45));
              let _index184 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1003_v38).length));
              (_1003_v38)[_index184] = _module.__default.fm31(globalState);
              let _index185 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1001_v36).length));
              (_1001_v36)[_index185] = (_this).f26;
              let _index186 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1001_v36).length));
              (_1001_v36)[_index186] = new BigNumber(625);
              let _1019_v46;
              _1019_v46 = _dafny.Seq.of((_this).f26);
              let _1020_v47;
              _1020_v47 = _dafny.Set.fromElements((_this).fm7(globalState), p1);
              let _1021_v48;
              let _nw169 = Array((new BigNumber(21)).toNumber());
              _nw169[(_dafny.ZERO).toNumber()] = _1019_v46;
              _nw169[(_dafny.ONE).toNumber()] = _module.__default.fm3((_1001_v36)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1001_v36).length))], false, (_this).f26, true, globalState);
              _nw169[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1019_v46, _1019_v46);
              _nw169[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(new BigNumber(352));
              _nw169[(new BigNumber(4)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(5)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(6)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(7)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(8)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(9)).toNumber()] = _module.__default.fm3((_this).f26, p2, new BigNumber((_1020_v47).length), false, globalState);
              _nw169[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f26, (_1001_v36)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1001_v36).length))], new BigNumber(784)), _1019_v46);
              _nw169[(new BigNumber(11)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(12)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f26, (_this).f26), _1019_v46);
              _nw169[(new BigNumber(14)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_1019_v46, _module.__default.safeIndex((_this).f26, new BigNumber((_1019_v46).length)), (_this).f26);
              _nw169[(new BigNumber(16)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(17)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(18)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus((_this).f26));
              _nw169[(new BigNumber(19)).toNumber()] = _1019_v46;
              _nw169[(new BigNumber(20)).toNumber()] = _dafny.Seq.of((_1001_v36)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1001_v36).length))]);
              _1021_v48 = _nw169;
              let _index187 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1021_v48).length));
              (_1021_v48)[_index187] = _1019_v46;
              let _index188 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1021_v48).length));
              (_1021_v48)[_index188] = _1019_v46;
            }
            let _1022_v49;
            let _nw170 = new _module.C2();
            _nw170.__ctor(new BigNumber(144), (p2) === ((_this).f25), ((_this).f26).multipliedBy((_this).f26));
            _1022_v49 = _nw170;
            let _1023_v50;
            _1023_v50 = _dafny.Map.Empty.slice().updateUnsafe((_983_v20)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_983_v20).length))],(_this).f26);
            _1023_v50 = (_1023_v50).update(p1, (_dafny.ZERO).minus((_1022_v49).f31));
          }
        }
      }
      (globalState).f0 = (_this).f26;
      let _1024_i3;
      _1024_i3 = _dafny.ZERO;
      L5: {
        while ((_module.__default.safeDivisionInt((_this).f26, _module.__default.fm2((_this).f26, (_this).f26, globalState))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of(true)).length))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1024_i3)) {
              break L5;
            }
            _1024_i3 = (_1024_i3).plus(_dafny.ONE);
            let _1025_v51;
            let _out23;
            _out23 = _module.__default.m0((_this).f25, globalState);
            _1025_v51 = _out23;
            (globalState).f0 = new BigNumber(-983);
            _958_v1 = (_958_v1).update((_this).f25, p2);
            _957_v0 = _957_v0;
          }
        }
      }
      let _hi9 = (_this).f26;
      for (let _1026_i4 = ((_this).f26).multipliedBy((_dafny.ZERO).minus((_this).f26)); _1026_i4.isLessThan(_hi9); _1026_i4 = _1026_i4.plus(_dafny.ONE)) {
        let _1027_v52;
        let _init26 = ((_1028_p1) => function (_1029_i5) {
          return _1028_p1;
        })(p1);
        let _nw171 = Array((new BigNumber(2)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw171.length); _i0_26++) {
          _nw171[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1027_v52 = _nw171;
        let _1030_v53;
        _1030_v53 = _dafny.Seq.of(_1027_v52);
        let _1031_v54;
        _1031_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1027_v52);
        _1030_v53 = _dafny.Seq.of((((_1031_v54).contains(!(!(p2)))) ? ((_1031_v54).get(!(!(p2)))) : (_1027_v52)), _1027_v52, _1027_v52, _1027_v52);
        (globalState).f0 = (_1026_i4).plus((_this).f26);
        let _1032_v55;
        let _nw172 = Array((new BigNumber(17)).toNumber()).fill(_module.D2.Default());
        _1032_v55 = _nw172;
        let _1033_v56;
        _1033_v56 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f26);
        let _1034_v57;
        _1034_v57 = _dafny.MultiSet.fromElements(_1033_v56, _1033_v56);
        let _1035_v58;
        _1035_v58 = _1034_v57;
        let _rhs198 = _1032_v55;
        let _rhs199 = _1035_v58;
        let _rhs200 = (_this).f26;
        let _lhs176 = globalState;
        _1032_v55 = _rhs198;
        _1035_v58 = _rhs199;
        _lhs176.f2 = _rhs200;
        let _1036_v59;
        let _init27 = function (_1037_i6) {
          return _dafny.Seq.of((_this).f26);
        };
        let _nw173 = Array((new BigNumber(5)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw173.length); _i0_27++) {
          _nw173[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1036_v59 = _nw173;
        let _1038_v60;
        _1038_v60 = _dafny.Seq.of(_1026_i4);
        let _index189 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1036_v59).length));
        (_1036_v59)[_index189] = _1038_v60;
        let _index190 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1036_v59).length));
        (_1036_v59)[_index190] = _1038_v60;
      }
      let _1039_v61;
      _1039_v61 = _module.D1.create_DC3(_957_v0);
      let _1040_v62;
      _1040_v62 = _dafny.Seq.of(_1039_v61);
      let _1041_v63;
      _1041_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_1040_v62);
      r0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((((_1041_v63).contains((_this).f26)) ? ((_1041_v63).get((_this).f26)) : (_1040_v62)), _1040_v62), _dafny.Seq.Create(_module.__default.abs(new BigNumber(304)), ((_1042_v61) => function (_1043_i7) {
        return _1042_v61;
      })(_1039_v61)));
      return r0;
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _source15 = _module.__default.fm32(globalState);
      if (_source15.is_DC10) {
        let _1044___mcc_h0 = (_source15).cf16;
        let _1045___mcc_h1 = (_source15).cf17;
        let _1046_cf17 = _1045___mcc_h1;
        let _1047_cf16 = _1044___mcc_h0;
        let _1048_v0;
        _1048_v0 = _dafny.MultiSet.fromElements((_this).f26, (_this).f26);
        let _1049_v1;
        _1049_v1 = _1048_v0;
        let _1050_v2;
        _1050_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _1051_v3;
        _1051_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),!(p1));
        let _1052_v4;
        let _nw174 = Array((new BigNumber(21)).toNumber());
        _nw174[(_dafny.ZERO).toNumber()] = true;
        _nw174[(_dafny.ONE).toNumber()] = (((_1050_v2).contains(p1)) ? ((_1050_v2).get(p1)) : (false));
        _nw174[(new BigNumber(2)).toNumber()] = (_this).f25;
        _nw174[(new BigNumber(3)).toNumber()] = p1;
        _nw174[(new BigNumber(4)).toNumber()] = p2;
        _nw174[(new BigNumber(5)).toNumber()] = p1;
        _nw174[(new BigNumber(6)).toNumber()] = (_this).f25;
        _nw174[(new BigNumber(7)).toNumber()] = p1;
        _nw174[(new BigNumber(8)).toNumber()] = (((_1051_v3).contains(new BigNumber(-745))) ? ((_1051_v3).get(new BigNumber(-745))) : (p2));
        _nw174[(new BigNumber(9)).toNumber()] = p2;
        _nw174[(new BigNumber(10)).toNumber()] = !(p1);
        _nw174[(new BigNumber(11)).toNumber()] = p2;
        _nw174[(new BigNumber(12)).toNumber()] = (_this).f25;
        _nw174[(new BigNumber(13)).toNumber()] = true;
        _nw174[(new BigNumber(14)).toNumber()] = (_this).f25;
        _nw174[(new BigNumber(15)).toNumber()] = p1;
        _nw174[(new BigNumber(16)).toNumber()] = (_this).f30;
        _nw174[(new BigNumber(17)).toNumber()] = p2;
        _nw174[(new BigNumber(18)).toNumber()] = p2;
        _nw174[(new BigNumber(19)).toNumber()] = p2;
        _nw174[(new BigNumber(20)).toNumber()] = p1;
        _1052_v4 = _nw174;
        let _1053_v5;
        let _nw175 = new _module.C0();
        _nw175.__ctor((_this).f29, true, p0, _1052_v4, (_this).f30);
        _1053_v5 = _nw175;
        let _1054_v6;
        _1054_v6 = _dafny.Set.fromElements((_this).f25);
        let _1055_v7;
        _1055_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1053_v5,new BigNumber((_1054_v6).length));
        let _1056_v8;
        _1056_v8 = _dafny.Seq.of((_this).f26, (_this).f26, (((_1055_v7).contains(_1053_v5)) ? ((_1055_v7).get(_1053_v5)) : ((_this).f26)));
        let _1057_v9;
        _1057_v9 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_1056_v8));
        let _1058_v10;
        let _init28 = ((_1059_v2) => function (_1060_i0) {
          return _1059_v2;
        })(_1050_v2);
        let _nw176 = Array((new BigNumber(6)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw176.length); _i0_28++) {
          _nw176[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1058_v10 = _nw176;
        let _1061_v11;
        _1061_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1053_v5).f26,_module.D8.create_DC16(_1058_v10));
        let _1062_v12;
        _1062_v12 = _dafny.MultiSet.fromElements(_1061_v11, _1061_v11);
        let _1063_v13;
        _1063_v13 = _dafny.MultiSet.fromElements(_module.__default.fm1(globalState));
        let _1064_v14;
        let _nw177 = Array((new BigNumber(21)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = (_this).f26;
        _nw177[(_dafny.ONE).toNumber()] = new BigNumber((_1057_v9).cardinality());
        _nw177[(new BigNumber(2)).toNumber()] = (_1053_v5).f26;
        _nw177[(new BigNumber(3)).toNumber()] = (_this).f26;
        _nw177[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus((((_1062_v12).contains(_1061_v11)) ? ((_1062_v12).get(_1061_v11)) : ((_1053_v5).f26)))).plus((_this).f26);
        _nw177[(new BigNumber(5)).toNumber()] = (((_1063_v13).contains(_dafny.Seq.UnicodeFromString("gesulblvl"))) ? ((_1063_v13).get(_dafny.Seq.UnicodeFromString("gesulblvl"))) : ((_dafny.ZERO).minus((_1053_v5).f26)));
        _nw177[(new BigNumber(6)).toNumber()] = (_this).f26;
        _nw177[(new BigNumber(7)).toNumber()] = (((_this).f30) ? (new BigNumber((_1051_v3).length)) : (p0));
        _nw177[(new BigNumber(8)).toNumber()] = (_1053_v5).f26;
        _nw177[(new BigNumber(9)).toNumber()] = p0;
        _nw177[(new BigNumber(10)).toNumber()] = (_1056_v8)[_module.__default.safeIndex((_1053_v5).f26, new BigNumber((_1056_v8).length))];
        _nw177[(new BigNumber(11)).toNumber()] = (_1053_v5).f26;
        _nw177[(new BigNumber(12)).toNumber()] = new BigNumber(-274);
        _nw177[(new BigNumber(13)).toNumber()] = new BigNumber((_1056_v8).length);
        _nw177[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_this).f29, (_this).f29)).length);
        _nw177[(new BigNumber(15)).toNumber()] = (_this).f26;
        _nw177[(new BigNumber(16)).toNumber()] = (new BigNumber(312)).minus((_1053_v5).f26);
        _nw177[(new BigNumber(17)).toNumber()] = p0;
        _nw177[(new BigNumber(18)).toNumber()] = (_this).fm6((_this).f25, (_this).f25, new BigNumber(132), globalState);
        _nw177[(new BigNumber(19)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1)).length)).minus((_1053_v5).f26);
        _nw177[(new BigNumber(20)).toNumber()] = (_this).f26;
        _1064_v14 = _nw177;
        let _index191 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length));
        (_1064_v14)[_index191] = p0;
        let _index192 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length));
        (_1064_v14)[_index192] = new BigNumber((_dafny.Seq.Concat((_this).f29, _dafny.Seq.UnicodeFromString("ix"))).length);
        let _1065_v15;
        _1065_v15 = _dafny.Set.fromElements(((p2) ? ((_1064_v14)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length))]) : ((_1053_v5).f26)));
        let _1066_v16;
        _1066_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1053_v5).f26,_module.__default.fm1(globalState));
        let _1067_v17;
        _1067_v17 = _dafny.Seq.of(p1, p2, p1, (new BigNumber(((_1066_v16).update((_1064_v14)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length))], (_this).f29)).length)).isLessThan((_1053_v5).f26));
        let _1068_v18;
        let _nw178 = Array((new BigNumber(9)).toNumber()).fill(_module.D6.Default());
        _1068_v18 = _nw178;
        let _1069_v19;
        _1069_v19 = _module.D6.create_DC13(_module.__default.fm2(_module.__default.fm2((_this).f26, new BigNumber(102), globalState), (_1064_v14)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length))], globalState), (_1053_v5).f26, _1065_v15);
        let _index193 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1068_v18).length));
        (_1068_v18)[_index193] = _1069_v19;
        let _pat_let_tv31 = _1064_v14;
        let _pat_let_tv32 = _1064_v14;
        let _index194 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1068_v18).length));
        let _rhs201 = _1065_v15;
        let _rhs202 = _1050_v2;
        let _rhs203 = _dafny.Seq.of((_1053_v5).f25);
        let _rhs204 = function (_pat_let36_0) {
          return function (_1070_dt__update__tmp_h1) {
            return function (_pat_let37_0) {
              return function (_1071_dt__update_hcf21_h0) {
                return _module.D6.create_DC13((_1070_dt__update__tmp_h1).dtor_cf20, _1071_dt__update_hcf21_h0, (_1070_dt__update__tmp_h1).dtor_cf22);
              }(_pat_let37_0);
            }((_pat_let_tv32)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_pat_let_tv31).length))]);
          }(_pat_let36_0);
        }(_1069_v19);
        let _rhs205 = _1047_cf16;
        let _lhs177 = _1068_v18;
        let _lhs178 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1068_v18).length));
        _1065_v15 = _rhs201;
        _1050_v2 = _rhs202;
        _1067_v17 = _rhs203;
        _lhs177[_lhs178] = _rhs204;
        _1047_cf16 = _rhs205;
        let _index195 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length));
        (_1064_v14)[_index195] = (_1064_v14)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_1064_v14).length))];
        (globalState).f1 = p2;
      } else {
        let _1072___mcc_h2 = (_source15).cf15;
        let _1073_cf15 = _1072___mcc_h2;
        let _1074_v20;
        _1074_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,p2);
        r1 = ((!((((_1074_v20).contains((_this).f26)) ? ((_1074_v20).get((_this).f26)) : ((_this).f30)))) ? ((_this).f26) : ((_this).f26));
        let _1075_v21;
        let _init29 = ((_1076_p1) => function (_1077_i1) {
          return _1076_p1;
        })(p1);
        let _nw179 = Array((new BigNumber(2)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw179.length); _i0_29++) {
          _nw179[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1075_v21 = _nw179;
        let _index196 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1075_v21).length));
        (_1075_v21)[_index196] = _dafny.areEqual((_this).f29, (_this).f29);
        let _1078_v22;
        _1078_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1074_v20).length),p2),p1);
        let _1079_v23;
        _1079_v23 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1080_v24;
        _1080_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1079_v23,(_this).f26);
        let _index197 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1075_v21).length));
        (_1075_v21)[_index197] = !(_dafny.MultiSet.fromElements(new BigNumber((_1080_v24).length))).contains((new BigNumber((_1078_v22).length)).plus(p0));
        let _1081_v25;
        let _nw180 = new _module.C1();
        _nw180.__ctor(p1);
        _1081_v25 = _nw180;
        let _1082_v26;
        let _nw181 = new _module.C2();
        _nw181.__ctor(p0, (_1075_v21)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_1075_v21).length))], new BigNumber(-12));
        _1082_v26 = _nw181;
      }
      let _hi10 = p0;
      for (let _1083_i2 = ((p2) ? (new BigNumber(606)) : ((_this).f26)); _1083_i2.isLessThan(_hi10); _1083_i2 = _1083_i2.plus(_dafny.ONE)) {
        let _1084_v27;
        let _nw182 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1084_v27 = _nw182;
        let _index198 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length));
        (_1084_v27)[_index198] = _1083_i2;
        let _1085_v28;
        _1085_v28 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm3(_1083_i2, p2, new BigNumber(-651), p1, globalState)).length));
        let _1086_v29;
        _1086_v29 = _dafny.MultiSet.fromElements(_1085_v28);
        let _index199 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length));
        (_1084_v27)[_index199] = (((_1086_v29).contains(_module.__default.fm21((_this).f29, (_this).f26, p0, globalState))) ? ((_1086_v29).get(_module.__default.fm21((_this).f29, (_this).f26, p0, globalState))) : (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), ((_1087_i2) => function (_1088_i3) {
          return _1087_i2;
        })(_1083_i2))).length), new BigNumber(((_this).f29).length))));
        let _1089_v30;
        _1089_v30 = _dafny.MultiSet.fromElements((_this).f30);
        let _1090_v31;
        let _nw183 = new _module.C2();
        _nw183.__ctor(new BigNumber((_1089_v30).cardinality()), false, (_dafny.ZERO).minus((_this).f26));
        _1090_v31 = _nw183;
        let _1091_v32;
        _1091_v32 = _dafny.MultiSet.fromElements(_1090_v31);
        if ((_dafny.MultiSet.fromElements(_1090_v31, _1090_v31)).IsProperSubsetOf(_1091_v32)) {
          r1 = p0;
          let _1092_v33;
          let _nw184 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
          _1092_v33 = _nw184;
          let _1093_v34;
          _1093_v34 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,(_1084_v27)[_module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length))]));
          let _index200 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1092_v33).length));
          (_1092_v33)[_index200] = _1093_v34;
          let _index201 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1092_v33).length));
          (_1092_v33)[_index201] = _1093_v34;
          let _1094_v35;
          _1094_v35 = _module.D0.create_DC0((_this).f30);
          (globalState).f1 = _module.__default.fm5((_this).f30, _dafny.Map.Empty.slice().updateUnsafe(_1094_v35,(_this).f26), (_this).f26, globalState);
          let _1095_v36;
          _1095_v36 = _dafny.Set.fromElements((_1084_v27)[_module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length))]);
          let _1096_v37;
          _1096_v37 = _dafny.Seq.of((_1090_v31).f31);
          r0 = (new BigNumber((_1095_v36).length)).minus(new BigNumber((_1096_v37).length));
          let _1097_v38;
          _1097_v38 = _module.D0.create_DC1((_dafny.ZERO).minus(_1083_i2), (_this).f29, p1, new BigNumber(145), p1);
          (globalState).f0 = (_1097_v38).dtor_cf1;
        } else {
          let _1098_v39;
          _1098_v39 = _module.D1.create_DC4((_this).f30, (_this).f25);
          let _1099_v40;
          _1099_v40 = _dafny.Seq.of(p1, (_1098_v39).dtor_cf7, (_this).f30);
          (globalState).f2 = (new BigNumber((_1099_v40).length)).multipliedBy(((p2) ? ((_1084_v27)[_module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length))]) : (new BigNumber(720))));
          let _1100_v41;
          _1100_v41 = _dafny.Seq.UnicodeFromString("pq");
          let _1101_v42;
          _1101_v42 = _dafny.MultiSet.fromElements(_1100_v41, (_this).f29, (_this).f29);
          let _1102_v43;
          _1102_v43 = _dafny.Set.fromElements((_1084_v27)[_module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length))]);
          let _1103_v44;
          let _nw185 = Array((new BigNumber(23)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = ((_1090_v31).f31).isLessThanOrEqualTo(new BigNumber(-210));
          _nw185[(_dafny.ONE).toNumber()] = (_1101_v42).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f29, _1100_v41, _1100_v41));
          _nw185[(new BigNumber(2)).toNumber()] = p2;
          _nw185[(new BigNumber(3)).toNumber()] = (_this).f30;
          _nw185[(new BigNumber(4)).toNumber()] = !((p1) && ((_this).f25));
          _nw185[(new BigNumber(5)).toNumber()] = ((true) ? ((_this).f25) : (p2));
          _nw185[(new BigNumber(6)).toNumber()] = (_this).f30;
          _nw185[(new BigNumber(7)).toNumber()] = (_1085_v28).IsSubsetOf(_1085_v28);
          _nw185[(new BigNumber(8)).toNumber()] = (_this).f30;
          _nw185[(new BigNumber(9)).toNumber()] = (_this).f25;
          _nw185[(new BigNumber(10)).toNumber()] = !((_this).f30) || ((_this).f25);
          _nw185[(new BigNumber(11)).toNumber()] = p2;
          _nw185[(new BigNumber(12)).toNumber()] = (_this).f25;
          _nw185[(new BigNumber(13)).toNumber()] = p1;
          _nw185[(new BigNumber(14)).toNumber()] = (_this).f25;
          _nw185[(new BigNumber(15)).toNumber()] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("nwtpj"), _1100_v41);
          _nw185[(new BigNumber(16)).toNumber()] = !((_this).f30);
          _nw185[(new BigNumber(17)).toNumber()] = (_this).f25;
          _nw185[(new BigNumber(18)).toNumber()] = p1;
          _nw185[(new BigNumber(19)).toNumber()] = (_dafny.Set.fromElements((_1090_v31).f31)).IsProperSubsetOf(_1102_v43);
          _nw185[(new BigNumber(20)).toNumber()] = (_this).f25;
          _nw185[(new BigNumber(21)).toNumber()] = !((_this).f30);
          _nw185[(new BigNumber(22)).toNumber()] = p1;
          _1103_v44 = _nw185;
          let _index202 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1103_v44).length));
          (_1103_v44)[_index202] = (new BigNumber((_1089_v30).cardinality())).isLessThan(_1083_i2);
          let _1104_v45;
          let _init30 = ((_1105_v28) => function (_1106_i4) {
            return _1105_v28;
          })(_1085_v28);
          let _nw186 = Array((new BigNumber(25)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw186.length); _i0_30++) {
            _nw186[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1104_v45 = _nw186;
          let _1107_v46;
          _1107_v46 = _module.D10.create_DC21(_1104_v45);
          let _index203 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1103_v44).length));
          let _index204 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length));
          let _rhs206 = _dafny.Seq.Concat((_this).f29, (_this).f29);
          let _rhs207 = true;
          let _rhs208 = _module.__default.safeDivisionInt(new BigNumber(-169), ((_dafny.ZERO).minus((_1084_v27)[_module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length))])).multipliedBy(_1083_i2));
          let _rhs209 = (_1107_v46).dtor_cf33;
          let _rhs210 = (_this).f25;
          let _lhs179 = _1103_v44;
          let _lhs180 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1103_v44).length));
          let _lhs181 = _1084_v27;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1084_v27).length));
          let _lhs183 = globalState;
          _1100_v41 = _rhs206;
          _lhs179[_lhs180] = _rhs207;
          _lhs181[_lhs182] = _rhs208;
          _1104_v45 = _rhs209;
          _lhs183.f1 = _rhs210;
          let _1108_v47;
          _1108_v47 = _dafny.Seq.of(_1084_v27);
          _1108_v47 = _dafny.Seq.Concat(_dafny.Seq.update(_1108_v47, _module.__default.safeIndex((_1090_v31).f31, new BigNumber((_1108_v47).length)), _1084_v27), _dafny.Seq.Concat(_1108_v47, _1108_v47));
          (globalState).f1 = !((_this).f25);
          let _1109_v48;
          _1109_v48 = _module.D0.create_DC0((_1103_v44)[_module.__default.safeIndex(new BigNumber(540), new BigNumber((_1103_v44).length))]);
          (globalState).f1 = (_1109_v48).dtor_cf0;
        }
        let _1110_v49;
        let _init31 = function (_1111_i5) {
          return (_this).f25;
        };
        let _nw187 = Array((new BigNumber(10)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw187.length); _i0_31++) {
          _nw187[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1110_v49 = _nw187;
        let _index205 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_1110_v49).length));
        (_1110_v49)[_index205] = p2;
        let _1112_v50;
        let _init32 = ((_1113_p2) => function (_1114_i6) {
          return _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1113_p2, (_this).f30),new _dafny.CodePoint('p'.codePointAt(0)))).length), (_this).f26);
        })(p2);
        let _nw188 = Array((new BigNumber(23)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw188.length); _i0_32++) {
          _nw188[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1112_v50 = _nw188;
        let _1115_v51;
        _1115_v51 = _dafny.Seq.of(p1);
        let _index206 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_1110_v49).length));
        let _rhs211 = _dafny.Seq.IsProperPrefixOf(_1115_v51, _dafny.Seq.of((_this).f30, (_this).f30));
        let _rhs212 = (_1090_v31).fm7(globalState);
        let _rhs213 = _1112_v50;
        let _lhs184 = _1110_v49;
        let _lhs185 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_1110_v49).length));
        r2 = _rhs211;
        _lhs184[_lhs185] = _rhs212;
        _1112_v50 = _rhs213;
      }
      if ((_this).f25) {
        let _1116_v52;
        _1116_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm7(globalState));
        let _1117_v53;
        _1117_v53 = _dafny.MultiSet.fromElements((_this).f26);
        let _1118_v54;
        _1118_v54 = _dafny.Seq.of((((_1116_v52).contains(new BigNumber((_1117_v53).cardinality()))) ? ((_1116_v52).get(new BigNumber((_1117_v53).cardinality()))) : ((_this).f30)));
        let _1119_v55;
        _1119_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v54,(_this).f30);
        let _1120_v56;
        _1120_v56 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1119_v55);
        let _1121_v57;
        _1121_v57 = _dafny.MultiSet.fromElements(true, (_this).f25, p1);
        let _1122_v58;
        _1122_v58 = _module.D6.create_DC14((_this).f26, new BigNumber(-684), _1121_v57);
        r0 = new BigNumber(((_1120_v56).Merge((_module.__default.fm33(_1122_v58, p0, globalState)).Merge(_1120_v56))).length);
        (globalState).f1 = ((_this).f26).isLessThan((_this).f26);
        let _1123_v59;
        let _init33 = ((_1124_p0) => function (_1125_i7) {
          return _module.__default.safeModuloInt(_1125_i7, _1124_p0);
        })(p0);
        let _nw189 = Array((_dafny.ONE).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw189.length); _i0_33++) {
          _nw189[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1123_v59 = _nw189;
        let _index207 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1123_v59).length));
        (_1123_v59)[_index207] = (_this).f26;
        let _index208 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1123_v59).length));
        (_1123_v59)[_index208] = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber(221), p0), (_this).f26);
        r0 = (_this).f26;
        let _1126_v60;
        let _init34 = ((_1127_p2) => function (_1128_i8) {
          return (_module.D11.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(_1127_p2,_1127_p2))).dtor_cf39;
        })(p2);
        let _nw190 = Array((new BigNumber(26)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw190.length); _i0_34++) {
          _nw190[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1126_v60 = _nw190;
        let _1129_v61;
        _1129_v61 = _module.D8.create_DC16(_1126_v60);
        let _1130_v62;
        _1130_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v54,_1129_v61);
        let _1131_v63;
        let _nw191 = new _module.C2();
        _nw191.__ctor(new BigNumber((_1130_v62).length), p1, (_dafny.ZERO).minus(new BigNumber((_1118_v54).length)));
        _1131_v63 = _nw191;
      } else {
        let _1132_v64;
        _1132_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber(545));
        _1132_v64 = _1132_v64;
        let _1133_v65;
        let _init35 = function (_1134_i9) {
          return _module.__default.safeModuloInt(_1134_i9, new BigNumber(843));
        };
        let _nw192 = Array((new BigNumber(7)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw192.length); _i0_35++) {
          _nw192[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1133_v65 = _nw192;
        let _1135_v66;
        let _init36 = function (_1136_i10) {
          return _dafny.Seq.of((_this).f29, (_this).f29, _dafny.Seq.Create(_module.__default.abs(new BigNumber(128)), function (_1137_i11) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }));
        };
        let _nw193 = Array((new BigNumber(5)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw193.length); _i0_36++) {
          _nw193[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1135_v66 = _nw193;
        let _1138_v67;
        _1138_v67 = _dafny.Seq.of((_this).f29);
        let _index209 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1135_v66).length));
        (_1135_v66)[_index209] = _dafny.Seq.Concat(_1138_v67, _1138_v67);
        let _1139_v68;
        _1139_v68 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1140_v69;
        _1140_v69 = _dafny.MultiSet.fromElements(new BigNumber(397), (_this).f26, (_this).f26);
        let _1141_v70;
        _1141_v70 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm28((_this).f29, _1139_v68, (_this).f26, globalState),new BigNumber(((_this).f29).length));
        let _1142_v73;
        _1142_v73 = _module.D0.create_DC0((_this).f30);
        let _1143_v74;
        _1143_v74 = _dafny.Seq.of(_1142_v73, _1142_v73);
        let _index210 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1135_v66).length));
        let _rhs214 = _1133_v65;
        let _rhs215 = (((_this).fm7(globalState)) ? (_1138_v67) : (_1138_v67));
        let _rhs216 = (_this).f25;
        let _rhs217 = _module.__default.fm5((_1140_v69).IsDisjointFrom(_module.__default.fm21(_dafny.Seq.UnicodeFromString("ncdfeu"), (_this).f26, (_this).f26, globalState)), ((false) ? (_1141_v70) : (function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of (function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_1143_v74).Elements) {
              let _1144_v72 = _compr_23;
              if (_dafny.Seq.contains(_1143_v74, _1144_v72)) {
                _coll23.push([_1144_v72,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f30)).length)]);
              }
            }
            return _coll23;
          }()).Keys.Elements) {
            let _1145_v71 = _compr_22;
            if ((function () {
              let _coll24 = new _dafny.Map();
              for (const _compr_24 of (_1143_v74).Elements) {
                let _1144_v72 = _compr_24;
                if (_dafny.Seq.contains(_1143_v74, _1144_v72)) {
                  _coll24.push([_1144_v72,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f30)).length)]);
                }
              }
              return _coll24;
            }()).contains(_1145_v71)) {
              _coll22.push([_1145_v71,new BigNumber(((_this).f29).length)]);
            }
          }
          return _coll22;
        }())), ((_this).f26).multipliedBy((((_1132_v64).contains((_this).f26)) ? ((_1132_v64).get((_this).f26)) : ((_dafny.ZERO).minus((_this).f26)))), globalState);
        let _rhs218 = _1139_v68;
        let _lhs186 = _1135_v66;
        let _lhs187 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1135_v66).length));
        _1133_v65 = _rhs214;
        _lhs186[_lhs187] = _rhs215;
        r2 = _rhs216;
        r2 = _rhs217;
        _1139_v68 = _rhs218;
        let _1146_v75;
        let _nw194 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
        _1146_v75 = _nw194;
        _1146_v75 = _1146_v75;
        let _1147_v76;
        let _nw195 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1147_v76 = _nw195;
        let _index211 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1147_v76).length));
        (_1147_v76)[_index211] = (_this).f25;
        let _index212 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1147_v76).length));
        (_1147_v76)[_index212] = ((_this).f26).isLessThanOrEqualTo(_module.__default.fm2(p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_1138_v67)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_1138_v67).length))]).length))), globalState));
        let _1148_v77;
        _1148_v77 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false, (_this).f25, p1)).length), p0);
        let _1149_v78;
        _1149_v78 = _dafny.Seq.of((_dafny.Set.fromElements((_this).f26)).IsDisjointFrom(_1148_v77), true);
        let _1150_v79;
        let _nw196 = new _module.C0();
        _nw196.__ctor((_this).f29, (_1147_v76)[_module.__default.safeIndex(new BigNumber(734), new BigNumber((_1147_v76).length))], (_this).f26, _1147_v76, false);
        _1150_v79 = _nw196;
        let _1151_v80;
        _1151_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v79,_1149_v78);
        _1149_v78 = (((_1151_v80).contains(_1150_v79)) ? ((_1151_v80).get(_1150_v79)) : (_module.__default.fm23(globalState)));
      }
      let _1152_v81;
      _1152_v81 = _module.D11.create_DC26((_this).f26, (_this).f26, (_this).f25);
      let _1153_v82;
      _1153_v82 = _dafny.Map.Empty.slice().updateUnsafe((_1152_v81).dtor_cf41,(_this).fm6(!(p2), (_1152_v81).dtor_cf42, (_dafny.ZERO).minus(p0), globalState));
      let _1154_v83;
      _1154_v83 = _dafny.Seq.of(p0, (_this).f26);
      _1153_v82 = (_1153_v82).update(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-976)), function (_1155_i12) {
        return (_this).f26;
      }), _1154_v83)).length), p0);
      let _1156_v84;
      let _init37 = function (_1157_i13) {
        return _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), function (_1158_i14) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), (_this).f29);
      };
      let _nw197 = Array((new BigNumber(3)).toNumber());
      for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw197.length); _i0_37++) {
        _nw197[_i0_37] = _init37(new BigNumber(_i0_37));
      }
      _1156_v84 = _nw197;
      let _index213 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_1156_v84).length));
      (_1156_v84)[_index213] = p1;
      let _index214 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_1156_v84).length));
      (_1156_v84)[_index214] = !((_this).f30);
      let _hi11 = _module.__default.safeModuloInt((_this).f26, (_dafny.ZERO).minus((_this).f26));
      for (let _1159_i15 = p0; _1159_i15.isLessThan(_hi11); _1159_i15 = _1159_i15.plus(_dafny.ONE)) {
        let _1160_v85;
        let _init38 = ((_1161_v84) => function (_1162_i16) {
          return (_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements((_1161_v84)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_1161_v84).length))]));
        })(_1156_v84);
        let _nw198 = Array((new BigNumber(24)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw198.length); _i0_38++) {
          _nw198[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1160_v85 = _nw198;
        let _1163_v86;
        _1163_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1154_v83).length),(_this).f25);
        let _1164_v87;
        _1164_v87 = _dafny.Set.fromElements((((_1163_v86).contains(_1159_i15)) ? ((_1163_v86).get(_1159_i15)) : ((_this).f30)));
        let _index215 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1160_v85).length));
        (_1160_v85)[_index215] = (_1164_v87).Union(_dafny.Set.fromElements((_this).f30));
        let _index216 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1160_v85).length));
        (_1160_v85)[_index216] = (_1164_v87).Intersect(_1164_v87);
        let _1165_v88;
        _1165_v88 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1156_v84)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_1156_v84).length))]);
        r2 = (((_1165_v88).contains((_this).f25)) ? ((_1165_v88).get((_this).f25)) : ((_1156_v84)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_1156_v84).length))]));
        let _1166_v89;
        _1166_v89 = _dafny.Seq.of((_1156_v84)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_1156_v84).length))]);
        let _1167_v90;
        _1167_v90 = _dafny.Set.fromElements(p0, _1159_i15, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_1166_v89, _module.__default.safeIndex(p0, new BigNumber((_1166_v89).length)), true))).cardinality()));
        let _1168_v91;
        let _nw199 = new _module.C2();
        _nw199.__ctor(new BigNumber(-927), (_this).f25, new BigNumber((_1167_v90).length));
        _1168_v91 = _nw199;
        let _1169_v92;
        _1169_v92 = _dafny.Seq.UnicodeFromString("wxlotbcg");
        let _rhs219 = (_this).f26;
        let _rhs220 = p1;
        let _rhs221 = _dafny.Seq.UnicodeFromString("c");
        let _lhs188 = globalState;
        let _lhs189 = globalState;
        _lhs188.f0 = _rhs219;
        _lhs189.f1 = _rhs220;
        _1169_v92 = _rhs221;
      }
      let _1170_v93;
      let _nw200 = new _module.C2();
      _nw200.__ctor((_this).f26, true, p0);
      _1170_v93 = _nw200;
      let _1171_v94;
      _1171_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1170_v93);
      r0 = new BigNumber(((_1171_v94).Merge((_1171_v94).Merge(_1171_v94))).length);
      r1 = (_this).f26;
      r2 = (_this).f30;
      return [r0, r1, r2];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f25 = false;
      this._f26 = _dafny.ZERO;
      this.f28 = _dafny.Seq.UnicodeFromString("");
      this._f27 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f27, f28, f25, f26) {
      let _this = this;
      (_this)._f27 = f27;
      (_this).f28 = f28;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f27;
    };
    fm7(globalState) {
      let _this = this;
      return !(!(!((_module.D0.create_DC0(true)).dtor_cf0)));
    };
    fm8(p0, globalState) {
      let _this = this;
      let _source16 = _module.D0.create_DC2();
      if (_source16.is_DC1) {
        let _1172___mcc_h0 = (_source16).cf1;
        let _1173___mcc_h1 = (_source16).cf2;
        let _1174___mcc_h2 = (_source16).cf3;
        let _1175___mcc_h3 = (_source16).cf4;
        let _1176___mcc_h4 = (_source16).cf5;
        let _1177_cf5 = _1176___mcc_h4;
        let _1178_cf4 = _1175___mcc_h3;
        let _1179_cf3 = _1174___mcc_h2;
        let _1180_cf2 = _1173___mcc_h1;
        let _1181_cf1 = _1172___mcc_h0;
        return _module.D0.create_DC2();
      } else if (_source16.is_DC2) {
        return _module.D0.create_DC2();
      } else {
        let _1182___mcc_h5 = (_source16).cf0;
        let _1183_cf0 = _1182___mcc_h5;
        return _module.D0.create_DC2();
      }
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f1 = (_this).f25;
      let _1184_v0;
      _1184_v0 = _dafny.Seq.of((_this).f25, (_this).f25);
      if ((_1184_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_1184_v0).length))]) {
        let _1185_v1;
        _1185_v1 = _dafny.Set.fromElements(_this.f28);
        let _1186_v2;
        _1186_v2 = _module.D0.create_DC0(p2);
        let _1187_v3;
        _1187_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1186_v2,(_this).f26);
        let _1188_v4;
        _1188_v4 = _dafny.Seq.of((_this).f26, (_this).f27);
        let _1189_v5;
        let _nw201 = Array((new BigNumber(25)).toNumber());
        _nw201[(_dafny.ZERO).toNumber()] = (true) || (p2);
        _nw201[(_dafny.ONE).toNumber()] = p2;
        _nw201[(new BigNumber(2)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(3)).toNumber()] = p2;
        _nw201[(new BigNumber(4)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(5)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(6)).toNumber()] = !(p1);
        _nw201[(new BigNumber(7)).toNumber()] = !(p1);
        _nw201[(new BigNumber(8)).toNumber()] = true;
        _nw201[(new BigNumber(9)).toNumber()] = true;
        _nw201[(new BigNumber(10)).toNumber()] = p1;
        _nw201[(new BigNumber(11)).toNumber()] = (_1185_v1).IsSubsetOf(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wrebsokg"), _this.f28));
        _nw201[(new BigNumber(12)).toNumber()] = (_1184_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_1184_v0).length))];
        _nw201[(new BigNumber(13)).toNumber()] = _module.__default.fm5((_1186_v2).dtor_cf0, (_1187_v3).update(_module.D0.create_DC0((_this).f25), (_1188_v4)[_module.__default.safeIndex((_this).f27, new BigNumber((_1188_v4).length))]), (_this).f26, globalState);
        _nw201[(new BigNumber(14)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(15)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(16)).toNumber()] = false;
        _nw201[(new BigNumber(17)).toNumber()] = !(p2);
        _nw201[(new BigNumber(18)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(19)).toNumber()] = p1;
        _nw201[(new BigNumber(20)).toNumber()] = !(p2) || ((_this).f25);
        _nw201[(new BigNumber(21)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(22)).toNumber()] = (_this).f25;
        _nw201[(new BigNumber(23)).toNumber()] = p2;
        _nw201[(new BigNumber(24)).toNumber()] = p1;
        _1189_v5 = _nw201;
        let _index217 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length));
        (_1189_v5)[_index217] = (new BigNumber((_this.f28).length)).isLessThanOrEqualTo((_this).f27);
        let _index218 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length));
        (_1189_v5)[_index218] = p2;
        if ((_this).f25) {
          let _1190_v6;
          let _nw202 = Array((_dafny.ONE).toNumber());
          _nw202[(_dafny.ZERO).toNumber()] = _this.f28;
          _1190_v6 = _nw202;
          let _index219 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_1190_v6).length));
          (_1190_v6)[_index219] = _dafny.Seq.Concat(_this.f28, _this.f28);
          let _index220 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_1190_v6).length));
          (_1190_v6)[_index220] = _module.__default.fm1(globalState);
          let _1191_v7;
          _1191_v7 = _dafny.Map.Empty.slice().updateUnsafe((_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))],(_this).f26);
          let _1192_v8;
          _1192_v8 = _module.D0.create_DC2();
          let _1193_v9;
          _1193_v9 = _dafny.MultiSet.fromElements(_1192_v8, _1192_v8);
          let _1194_v10;
          _1194_v10 = _dafny.Seq.of(_1192_v8);
          let _1195_v11;
          _1195_v11 = _module.D0.create_DC1(new BigNumber(535), (_1190_v6)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_1190_v6).length))], p2, (_this).f26, (_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))]);
          let _rhs222 = (_dafny.MultiSet.fromElements(_1192_v8, _1192_v8, _1192_v8, _module.D0.create_DC2())).IsSubsetOf((_1193_v9).Intersect(_dafny.MultiSet.FromArray(_1194_v10)));
          let _rhs223 = (_this).f25;
          let _rhs224 = (_dafny.Set.fromElements((_this).f26, (_this).f27, (_this).f27, new BigNumber((_dafny.Set.fromElements(false)).length))).contains(_module.__default.safeModuloInt((_this).f26, new BigNumber(50)));
          let _rhs225 = (_1195_v11).dtor_cf4;
          let _rhs226 = (_1191_v7).Merge(_1191_v7);
          let _lhs190 = globalState;
          let _lhs191 = globalState;
          let _lhs192 = globalState;
          let _lhs193 = globalState;
          _lhs190.f1 = _rhs222;
          _lhs191.f1 = _rhs223;
          _lhs192.f1 = _rhs224;
          _lhs193.f2 = _rhs225;
          _1191_v7 = _rhs226;
          let _index221 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length));
          (_1189_v5)[_index221] = (_this).f25;
          let _1196_v12;
          _1196_v12 = new _dafny.CodePoint('b'.codePointAt(0));
          _1196_v12 = _1196_v12;
          let _1197_v13;
          let _nw203 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
          _1197_v13 = _nw203;
          let _1198_v14;
          let _nw204 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1198_v14 = _nw204;
          let _1199_v15;
          _1199_v15 = _dafny.Set.fromElements(new BigNumber((_1188_v4).length), (_this).f27, (_this).f26, (_this).f27);
          let _index222 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1197_v13).length));
          (_1197_v13)[_index222] = (_dafny.Set.fromElements(new BigNumber(696), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1198_v14)).length))).Intersect(_1199_v15);
          let _index223 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1197_v13).length));
          let _rhs227 = (new BigNumber((_module.__default.fm1(globalState)).length)).minus(_module.__default.safeModuloInt(new BigNumber(261), (_this).f26));
          let _rhs228 = _1199_v15;
          let _rhs229 = _1196_v12;
          let _lhs194 = globalState;
          let _lhs195 = _1197_v13;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1197_v13).length));
          _lhs194.f0 = _rhs227;
          _lhs195[_lhs196] = _rhs228;
          _1196_v12 = _rhs229;
        } else {
          (globalState).f1 = (_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))];
          let _1200_v16;
          _1200_v16 = _dafny.Set.fromElements((_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))], (_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))]);
          let _1201_v17;
          _1201_v17 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1202_v18;
          _1202_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1200_v16,_1201_v17);
          let _1203_v19;
          _1203_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1202_v18).update(_1200_v16, _1201_v17),(_dafny.ZERO).minus(((_this).f26).minus((_this).f27)));
          _1203_v19 = (_1203_v19).update((_1202_v18).Merge(_1202_v18), (_this).f26);
          _1186_v2 = _1186_v2;
          let _1204_v20;
          let _nw205 = new _module.C3();
          _nw205.__ctor(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-475)), function (_1205_i0) {
            return (_this.f28)[_module.__default.safeIndex((_this).f27, new BigNumber((_this.f28).length))];
          }), _this.f28), p1, !((_this).f26).isEqualTo((_dafny.ZERO).minus((_this).f27)), (_this).f27);
          _1204_v20 = _nw205;
          let _1206_v21;
          let _nw206 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1206_v21 = _nw206;
          let _1207_v22;
          _1207_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1206_v21,(_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))]);
          let _1208_v23;
          _1208_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,((_dafny.ZERO).minus(new BigNumber((_1207_v22).length))).minus((_1204_v20).f26));
          let _1209_v24;
          _1209_v24 = _module.D0.create_DC1(new BigNumber((_this.f28).length), _dafny.Seq.UnicodeFromString("cmdoif"), p2, (_this).f26, (_this).f25);
          let _rhs230 = (_this).f25;
          let _rhs231 = _1204_v20;
          let _rhs232 = !((_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))]) || ((_1204_v20).fm7(globalState));
          let _rhs233 = (((_1204_v20).f25) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f26)) : (_1208_v23));
          let _rhs234 = _1209_v24;
          let _lhs197 = globalState;
          let _lhs198 = globalState;
          _lhs197.f1 = _rhs230;
          _1204_v20 = _rhs231;
          _lhs198.f1 = _rhs232;
          _1208_v23 = _rhs233;
          _1209_v24 = _rhs234;
          let _1210_v25;
          let _nw207 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1210_v25 = _nw207;
          let _1211_v26;
          _1211_v26 = _dafny.Seq.of(_1210_v25, _1210_v25);
          (globalState).f1 = ((_1204_v20).f26).isLessThanOrEqualTo(new BigNumber((_1211_v26).length));
        }
        let _index224 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length));
        (_1189_v5)[_index224] = (_1184_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_1184_v0).length))];
        let _1212_v27;
        let _init39 = function (_1213_i1) {
          return _this.f28;
        };
        let _nw208 = Array((new BigNumber(19)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw208.length); _i0_39++) {
          _nw208[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1212_v27 = _nw208;
        let _index225 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1212_v27).length));
        (_1212_v27)[_index225] = _dafny.Seq.Concat(_this.f28, _this.f28);
        let _index226 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1212_v27).length));
        (_1212_v27)[_index226] = _this.f28;
        let _1214_v28;
        _1214_v28 = _dafny.Set.fromElements(_1184_v0);
        if ((_1214_v28).IsProperSubsetOf(_module.__default.fm13((_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))], globalState))) {
          let _index227 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length));
          (_1189_v5)[_index227] = p2;
          let _1215_v29;
          _1215_v29 = _dafny.MultiSet.fromElements(_1189_v5);
          let _1216_v30;
          _1216_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))]);
          _1215_v29 = ((_1215_v29).Union((_1215_v29).update(_1189_v5, _module.__default.abs(new BigNumber((_1216_v30).length))))).Difference((_1215_v29).Union(_1215_v29));
          let _1217_v31;
          let _1218_v32;
          let _1219_v33;
          let _1220_v34;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector4 = (_this).m2(p1, globalState);
          _out24 = _outcollector4[0];
          _out25 = _outcollector4[1];
          _out26 = _outcollector4[2];
          _out27 = _outcollector4[3];
          _1217_v31 = _out24;
          _1218_v32 = _out25;
          _1219_v33 = _out26;
          _1220_v34 = _out27;
          let _1221_v35;
          _1221_v35 = new _dafny.CodePoint('b'.codePointAt(0));
          _1221_v35 = _1221_v35;
          let _1222_v36;
          _1222_v36 = _dafny.Seq.of(((p2) ? (_1189_v5) : (_1189_v5)), _1189_v5, _1189_v5);
          let _rhs235 = ((p2) ? (_dafny.Seq.IsProperPrefixOf(_this.f28, (_1212_v27)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_1212_v27).length))])) : ((_this).f25));
          let _rhs236 = _dafny.Seq.of(_1189_v5, _1189_v5, _1189_v5, _1189_v5, _1189_v5);
          _1218_v32 = _rhs235;
          _1222_v36 = _rhs236;
        } else {
          (_this).f28 = _this.f28;
          let _1223_v37;
          _1223_v37 = _dafny.MultiSet.fromElements((_this).f25, p1);
          (globalState).f1 = ((_1223_v37).Difference(_dafny.MultiSet.fromElements(p1, p1))).IsProperSubsetOf((_dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))], p1, p1)).Union(_1223_v37));
          let _1224_v38;
          _1224_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1212_v27)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_1212_v27).length))]);
          let _1225_v39;
          _1225_v39 = _dafny.Seq.of((_1212_v27)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_1212_v27).length))], _module.__default.fm1(globalState), _this.f28, _dafny.Seq.UnicodeFromString("wddepf"), (((_1224_v38).contains((_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))])) ? ((_1224_v38).get((_1189_v5)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_1189_v5).length))])) : ((_1212_v27)[_module.__default.safeIndex(new BigNumber(720), new BigNumber((_1212_v27).length))])));
          let _1226_v40;
          _1226_v40 = _dafny.Set.fromElements(_module.__default.fm5((_this).f25, _1187_v3, new BigNumber((_1184_v0).length), globalState));
          let _1227_v41;
          let _nw209 = new _module.C3();
          _nw209.__ctor((_1225_v39)[_module.__default.safeIndex(new BigNumber((_1226_v40).length), new BigNumber((_1225_v39).length))], !(_1226_v40).equals(_1226_v40), p1, (_dafny.ZERO).minus((_this).f26));
          _1227_v41 = _nw209;
          (globalState).f1 = !(p2) || (!(true) || (p1));
          let _1228_v42;
          _1228_v42 = _dafny.Map.Empty.slice().updateUnsafe((_1227_v41).f30,true);
          let _1229_v43;
          _1229_v43 = _module.D11.create_DC25(_1228_v42);
          let _1230_v44;
          _1230_v44 = new _dafny.CodePoint('j'.codePointAt(0));
          _1229_v43 = _module.__default.fm34(_dafny.Seq.update((_1227_v41).f29, _module.__default.safeIndex((_this).f26, new BigNumber(((_1227_v41).f29).length)), _1230_v44), (_this).f27, globalState);
        }
      } else {
        let _1231_v45;
        _1231_v45 = _dafny.Seq.of((_this).f26, (_this).f27, (_this).f27, new BigNumber(844));
        let _1232_v46;
        _1232_v46 = _dafny.Set.fromElements(p2);
        let _1233_v47;
        _1233_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_1232_v46);
        (globalState).f1 = (_1233_v47).contains((((_this).f25) ? ((_1231_v45)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_1231_v45).length))]) : ((_this).f27)));
        r0 = p1;
        let _1234_v48;
        _1234_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,((_this).f26).isLessThan(new BigNumber(649)));
        let _1235_v49;
        _1235_v49 = _dafny.MultiSet.fromElements(p0);
        let _1236_v50;
        _1236_v50 = _module.D0.create_DC1(new BigNumber((_1235_v49).cardinality()), _this.f28, p1, (_this).f26, p2);
        _1234_v48 = (_1234_v48).update(_module.__default.safeModuloInt(new BigNumber(-915), (_this).f27), (_1236_v50).dtor_cf3);
        let _1237_v51;
        _1237_v51 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1238_v52;
        _1238_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v51,p2);
        let _1239_v54;
        _1239_v54 = _dafny.Set.fromElements(_1237_v51);
        _1238_v52 = (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),p2)).Merge(function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of (_1239_v54).Elements) {
            let _1240_v53 = _compr_25;
            if ((_1239_v54).contains(_1240_v53)) {
              _coll25.push([_1240_v53,true]);
            }
          }
          return _coll25;
        }());
        (_this).f28 = _module.__default.fm1(globalState);
      }
      let _1241_v55;
      _1241_v55 = _dafny.Seq.of((_this).f26);
      if (!(_module.__default.safeDivisionInt((_this).f27, (_this).f27)).isEqualTo((_1241_v55)[_module.__default.safeIndex((_1241_v55)[_module.__default.safeIndex((_this).f27, new BigNumber((_1241_v55).length))], new BigNumber((_1241_v55).length))])) {
        let _1242_v56;
        _1242_v56 = _module.D1.create_DC4(false, p1);
        let _source17 = _1242_v56;
        if (_source17.is_DC4) {
          let _1243___mcc_h0 = (_source17).cf7;
          let _1244___mcc_h1 = (_source17).cf8;
          let _1245_cf8 = _1244___mcc_h1;
          let _1246_cf7 = _1243___mcc_h0;
          let _1247_v57;
          _1247_v57 = _module.D6.create_DC14(new BigNumber(-775), (_this).f26, _dafny.MultiSet.fromElements(false, !(_1246_cf7)));
          let _1248_v58;
          _1248_v58 = _dafny.Set.fromElements((_1247_v57).dtor_cf23);
          let _1249_v59;
          _1249_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f26);
          let _rhs237 = _dafny.Seq.Concat(_dafny.Seq.Concat(_this.f28, _this.f28), _this.f28);
          let _rhs238 = ((_this).f26).isLessThan(new BigNumber((_1248_v58).length));
          let _rhs239 = ((function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (_1249_v59).Keys.Elements) {
              let _1250_v60 = _compr_26;
              if ((_1249_v59).contains(_1250_v60)) {
                _coll26.add((_1250_v60).minus(new BigNumber(746)));
              }
            }
            return _coll26;
          }()).Difference(_1248_v58)).Intersect(_dafny.Set.fromElements((_this).f27));
          let _lhs199 = _this;
          _lhs199.f28 = _rhs237;
          _1245_cf8 = _rhs238;
          _1248_v58 = _rhs239;
          _1184_v0 = _1184_v0;
          let _1251_v61;
          _1251_v61 = _module.D2.create_DC7(_1246_cf7, _1245_cf8, (_this).f27);
          let _1252_v62;
          let _nw210 = Array((new BigNumber(20)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = new BigNumber((_this.f28).length);
          _nw210[(_dafny.ONE).toNumber()] = (_this).f26;
          _nw210[(new BigNumber(2)).toNumber()] = ((_1241_v55)[_module.__default.safeIndex(new BigNumber(-480), new BigNumber((_1241_v55).length))]).minus((_this).f27);
          _nw210[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(901), new BigNumber((p0).cardinality()));
          _nw210[(new BigNumber(4)).toNumber()] = (_this).f26;
          _nw210[(new BigNumber(5)).toNumber()] = _module.__default.fm2((_this).f27, new BigNumber((p0).cardinality()), globalState);
          _nw210[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_1248_v58).length))).minus((_this).f26);
          _nw210[(new BigNumber(7)).toNumber()] = (_this).f26;
          _nw210[(new BigNumber(8)).toNumber()] = _module.__default.fm2((_this).f26, _module.__default.fm2((_this).f27, new BigNumber((_1241_v55).length), globalState), globalState);
          _nw210[(new BigNumber(9)).toNumber()] = new BigNumber((_1184_v0).length);
          _nw210[(new BigNumber(10)).toNumber()] = ((_this).f27).minus((_1251_v61).dtor_cf13);
          _nw210[(new BigNumber(11)).toNumber()] = (_this).f26;
          _nw210[(new BigNumber(12)).toNumber()] = (_this).f26;
          _nw210[(new BigNumber(13)).toNumber()] = (_1247_v57).dtor_cf24;
          _nw210[(new BigNumber(14)).toNumber()] = (_this).f27;
          _nw210[(new BigNumber(15)).toNumber()] = ((_this).f26).multipliedBy((_this).f26);
          _nw210[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f27, (_this).f27));
          _nw210[(new BigNumber(17)).toNumber()] = (_this).f27;
          _nw210[(new BigNumber(18)).toNumber()] = (_this).f27;
          _nw210[(new BigNumber(19)).toNumber()] = (_this).f27;
          _1252_v62 = _nw210;
          let _nw211 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1252_v62 = _nw211;
          let _1253_v63;
          _1253_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v56,(_this).f27);
          let _index228 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_1252_v62).length));
          (_1252_v62)[_index228] = (_dafny.ZERO).minus((_this).f26);
          let _index229 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_1252_v62).length));
          let _rhs240 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(p1, true),(new BigNumber(8)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-569)), function (_1254_i2) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          })).length)));
          let _rhs241 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f27), (_this).f26);
          let _lhs200 = _1252_v62;
          let _lhs201 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_1252_v62).length));
          _1253_v63 = _rhs240;
          _lhs200[_lhs201] = _rhs241;
        } else if (_source17.is_DC3) {
          let _1255___mcc_h2 = (_source17).cf6;
          let _1256_cf6 = _1255___mcc_h2;
          let _1257_v64;
          _1257_v64 = _module.D2.create_DC6(_dafny.Set.fromElements((_this).f26));
          _1257_v64 = _1257_v64;
          let _1258_v65;
          _1258_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
          let _1259_v66;
          _1259_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f27);
          let _1260_v67;
          _1260_v67 = _module.D0.create_DC1((_this).f27, _dafny.Seq.UnicodeFromString("ssmnmk"), false, (_this).f27, (_this).f25);
          let _1261_v68;
          let _nw212 = Array((new BigNumber(27)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = (((_1258_v65).contains((_this).f27)) ? ((_1258_v65).get((_this).f27)) : ((_this).f27));
          _nw212[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber(((_1259_v66).update(p1, (_this).f26)).length), (_this).f27);
          _nw212[(new BigNumber(2)).toNumber()] = (_1260_v67).dtor_cf4;
          _nw212[(new BigNumber(3)).toNumber()] = ((_this).f26).multipliedBy((_this).f26);
          _nw212[(new BigNumber(4)).toNumber()] = (((_1259_v66).contains(!(p2))) ? ((_1259_v66).get(!(p2))) : (new BigNumber(-159)));
          _nw212[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_this.f28)).cardinality());
          _nw212[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.of(p2)).length);
          _nw212[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_this).f27, (_this).f27);
          _nw212[(new BigNumber(8)).toNumber()] = (_this).f27;
          _nw212[(new BigNumber(9)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(10)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(11)).toNumber()] = new BigNumber(619);
          _nw212[(new BigNumber(12)).toNumber()] = _module.__default.fm2((_this).f26, new BigNumber(731), globalState);
          _nw212[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-263));
          _nw212[(new BigNumber(14)).toNumber()] = ((_this).f26).multipliedBy((_this).f26);
          _nw212[(new BigNumber(15)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(16)).toNumber()] = (_this).f27;
          _nw212[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt((_this).f27, (_this).f26);
          _nw212[(new BigNumber(18)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(19)).toNumber()] = (_this).f27;
          _nw212[(new BigNumber(20)).toNumber()] = ((_this).f26).multipliedBy((_this).f26);
          _nw212[(new BigNumber(21)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(22)).toNumber()] = (((p0).contains((_this).f26)) ? ((p0).get((_this).f26)) : ((_this).f26));
          _nw212[(new BigNumber(23)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(24)).toNumber()] = (_this).f27;
          _nw212[(new BigNumber(25)).toNumber()] = (_this).f26;
          _nw212[(new BigNumber(26)).toNumber()] = new BigNumber((_this.f28).length);
          _1261_v68 = _nw212;
          let _nw213 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1261_v68 = _nw213;
          (globalState).f0 = (_this).f26;
          (globalState).f1 = ((!(p2) || (p2)) ? (p1) : (!((_this).f25) || (p1)));
        } else {
          let _1262___mcc_h3 = (_source17).cf9;
          let _1263_cf9 = _1262___mcc_h3;
          let _1264_v69;
          let _nw214 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1264_v69 = _nw214;
          let _init40 = function (_1265_i3) {
            return _module.__default.safeModuloInt(_1265_i3, (_this).f26);
          };
          let _nw215 = Array((new BigNumber(8)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw215.length); _i0_40++) {
            _nw215[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1264_v69 = _nw215;
          let _1266_v70;
          _1266_v70 = _module.D0.create_DC1(new BigNumber((_this.f28).length), _dafny.Seq.UnicodeFromString("dl"), true, (_this).f26, !(false));
          let _1267_v71;
          _1267_v71 = new _dafny.CodePoint('q'.codePointAt(0));
          (_this).f28 = _dafny.Seq.update(_dafny.Seq.Concat((_1266_v70).dtor_cf2, (((_this).f25) ? (_this.f28) : (_this.f28))), _module.__default.safeIndex((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).length)).plus((_this).f26), new BigNumber((_dafny.Seq.Concat((_1266_v70).dtor_cf2, (((_this).f25) ? (_this.f28) : (_this.f28)))).length)), _1267_v71);
          _1242_v56 = _module.D1.create_DC4(((_this).f26).isLessThan((_this).f26), (_this).f25);
          let _1268_v72;
          let _nw216 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1268_v72 = _nw216;
          let _1269_v73;
          let _init41 = ((_1270_p1) => function (_1271_i4) {
            return _dafny.Map.Empty.slice().updateUnsafe(!((_this).f25),_1270_p1);
          })(p1);
          let _nw217 = Array((new BigNumber(4)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw217.length); _i0_41++) {
            _nw217[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1269_v73 = _nw217;
          let _1272_v74;
          _1272_v74 = _module.D8.create_DC16(_1269_v73);
          let _1273_v75;
          _1273_v75 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
          let _1274_v76;
          _1274_v76 = _module.D12.create_DC28(_1264_v69);
          let _rhs242 = new BigNumber(576);
          let _rhs243 = ((!(_module.__default.fm5((_this).f25, _module.__default.fm25((_this).f26, (_this).f25, (_this).f25, globalState), new BigNumber((_1273_v75).length), globalState))) ? ((_1274_v76).dtor_cf44) : (_1264_v69));
          let _rhs244 = _1268_v72;
          let _rhs245 = _1272_v74;
          let _rhs246 = (_this).f25;
          let _lhs202 = globalState;
          let _lhs203 = globalState;
          _lhs202.f2 = _rhs242;
          _1264_v69 = _rhs243;
          _1268_v72 = _rhs244;
          _1272_v74 = _rhs245;
          _lhs203.f1 = _rhs246;
        }
        let _1275_v77;
        let _out28;
        _out28 = _module.__default.m0((((_this).f25) ? (true) : ((_this).f25)), globalState);
        _1275_v77 = _out28;
        let _1276_v78;
        let _nw218 = Array((new BigNumber(19)).toNumber());
        _nw218[(_dafny.ZERO).toNumber()] = (_this).f25;
        _nw218[(_dafny.ONE).toNumber()] = (_this).f25;
        _nw218[(new BigNumber(2)).toNumber()] = p1;
        _nw218[(new BigNumber(3)).toNumber()] = p1;
        _nw218[(new BigNumber(4)).toNumber()] = p2;
        _nw218[(new BigNumber(5)).toNumber()] = !((_this).f25);
        _nw218[(new BigNumber(6)).toNumber()] = (_this).fm7(globalState);
        _nw218[(new BigNumber(7)).toNumber()] = (_this).f25;
        _nw218[(new BigNumber(8)).toNumber()] = p2;
        _nw218[(new BigNumber(9)).toNumber()] = (_this).f25;
        _nw218[(new BigNumber(10)).toNumber()] = p2;
        _nw218[(new BigNumber(11)).toNumber()] = p2;
        _nw218[(new BigNumber(12)).toNumber()] = (_this).f25;
        _nw218[(new BigNumber(13)).toNumber()] = p2;
        _nw218[(new BigNumber(14)).toNumber()] = p1;
        _nw218[(new BigNumber(15)).toNumber()] = p2;
        _nw218[(new BigNumber(16)).toNumber()] = false;
        _nw218[(new BigNumber(17)).toNumber()] = (_this).f25;
        _nw218[(new BigNumber(18)).toNumber()] = p2;
        _1276_v78 = _nw218;
        let _1277_v79;
        _1277_v79 = new _dafny.CodePoint('p'.codePointAt(0));
        let _1278_v80;
        _1278_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1276_v78,_1277_v79);
        let _1279_v81;
        _1279_v81 = _dafny.Set.fromElements(_1275_v77, _1275_v77, (_this).f27, new BigNumber((_1278_v80).length), new BigNumber((_this.f28).length));
        let _1280_v82;
        _1280_v82 = _dafny.MultiSet.fromElements(_1279_v81);
        _1280_v82 = (_1280_v82).Difference((((_this).f25) ? (_1280_v82) : (_1280_v82)));
        let _1281_v83;
        let _nw219 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _1281_v83 = _nw219;
        _1281_v83 = _1281_v83;
        let _1282_v84;
        let _nw220 = new _module.C1();
        _nw220.__ctor((_this).f25);
        _1282_v84 = _nw220;
        let _1283_v85;
        _1283_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,((false) ? (_1282_v84) : (_1282_v84)));
        _1283_v85 = (_1283_v85).update(!(p2), _1282_v84);
      } else {
        if ((_this).f25) {
          let _1284_v86;
          _1284_v86 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1285_v87;
          _1285_v87 = _module.D6.create_DC14((_this).f26, (_this).f26, _dafny.MultiSet.fromElements(p2, (_this).f25));
          let _1286_v88;
          _1286_v88 = _module.D9.create_DC19(_1285_v87);
          let _1287_v89;
          _1287_v89 = _module.D9.create_DC20(((p1) ? (_1286_v88) : (_1286_v88)));
          let _1288_v90;
          _1288_v90 = _dafny.Set.fromElements(!((_1184_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), function (_1289_i5) {
            return _this.f28;
          })).length), new BigNumber((_1184_v0).length))]));
          let _pat_let_tv33 = _1286_v88;
          let _rhs247 = _1284_v86;
          let _rhs248 = !(p2) || (!((_1288_v90).equals(_1288_v90)));
          let _rhs249 = function (_pat_let38_0) {
            return function (_1290_dt__update__tmp_h0) {
              return function (_pat_let39_0) {
                return function (_1291_dt__update_hcf32_h0) {
                  return _module.D9.create_DC20(_1291_dt__update_hcf32_h0);
                }(_pat_let39_0);
              }(_module.D9.create_DC20(_pat_let_tv33));
            }(_pat_let38_0);
          }(_1287_v89);
          let _rhs250 = (new BigNumber(((((_this).f25) ? (_this.f28) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), ((_1292_v86) => function (_1293_i6) {
            return _1292_v86;
          })(_1284_v86))))).length)).isLessThan((new BigNumber((_1241_v55).length)).multipliedBy((_this).f26));
          let _lhs204 = globalState;
          _1284_v86 = _rhs247;
          r0 = _rhs248;
          _1287_v89 = _rhs249;
          _lhs204.f1 = _rhs250;
          let _1294_v91;
          let _nw221 = new _module.C3();
          _nw221.__ctor(_this.f28, p2, p2, (_this).f26);
          _1294_v91 = _nw221;
          let _1295_v92;
          _1295_v92 = _dafny.Seq.of(_module.__default.fm1(globalState));
          _1295_v92 = _dafny.Seq.of(_this.f28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(150)), ((_1296_v86) => function (_1297_i7) {
            return _1296_v86;
          })(_1284_v86)), _this.f28, _this.f28);
          let _1298_v93;
          let _nw222 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1298_v93 = _nw222;
          let _1299_v94;
          _1299_v94 = _module.D10.create_DC22(_1298_v93, _this.f28);
          let _1300_v95;
          _1300_v95 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_1299_v94).dtor_cf35);
          (globalState).f1 = _dafny.Seq.IsPrefixOf((((_1300_v95).contains(_this.f28)) ? ((_1300_v95).get(_this.f28)) : (_dafny.Seq.UnicodeFromString("rj"))), _this.f28);
          let _1301_v96;
          let _nw223 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1301_v96 = _nw223;
          let _index230 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_1301_v96).length));
          (_1301_v96)[_index230] = (_this).f27;
          let _index231 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_1301_v96).length));
          (_1301_v96)[_index231] = (_this).f27;
        } else {
          (globalState).f1 = p2;
          (globalState).f1 = p1;
          let _1302_v97;
          _1302_v97 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),false);
          let _1303_v98;
          let _nw224 = new _module.C1();
          _nw224.__ctor(p2);
          _1303_v98 = _nw224;
          let _1304_v99;
          _1304_v99 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1305_v100;
          _1305_v100 = _module.D10.create_DC24((_this).f27, _1303_v98, _1304_v99);
          let _1306_v101;
          _1306_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1305_v100,_dafny.MultiSet.fromElements((_this).f27));
          _1302_v97 = (_1302_v97).update(!(false), (new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality())).isEqualTo(new BigNumber(((((_1306_v101).contains(_1305_v100)) ? ((_1306_v101).get(_1305_v100)) : (p0))).cardinality())));
          let _1307_v102;
          let _nw225 = new _module.C1();
          _nw225.__ctor(p1);
          _1307_v102 = _nw225;
          (globalState).f1 = (_this).f25;
        }
        let _1308_v103;
        let _nw226 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1308_v103 = _nw226;
        let _index232 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length));
        (_1308_v103)[_index232] = (_this).f25;
        let _index233 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length));
        (_1308_v103)[_index233] = true;
        let _1309_v104;
        let _nw227 = Array((new BigNumber(15)).toNumber());
        _nw227[(_dafny.ZERO).toNumber()] = _1308_v103;
        _nw227[(_dafny.ONE).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(2)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(3)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(4)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(5)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(6)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(7)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(8)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(9)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(10)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(11)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(12)).toNumber()] = ((p2) ? (_1308_v103) : (_1308_v103));
        _nw227[(new BigNumber(13)).toNumber()] = _1308_v103;
        _nw227[(new BigNumber(14)).toNumber()] = _1308_v103;
        _1309_v104 = _nw227;
        let _index234 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1309_v104).length));
        (_1309_v104)[_index234] = _1308_v103;
        let _1310_v105;
        _1310_v105 = _dafny.Seq.of(_1308_v103, _1308_v103, _1308_v103, _1308_v103);
        let _index235 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1309_v104).length));
        (_1309_v104)[_index235] = (_1310_v105)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_1310_v105).length))];
        let _1311_v106;
        let _nw228 = new _module.C1();
        _nw228.__ctor(((_1308_v103)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length))]) && (p2));
        _1311_v106 = _nw228;
        if (((_this).f27).isLessThanOrEqualTo(_module.__default.safeModuloInt((_this).f27, (_this).f26))) {
          let _1312_v107;
          _1312_v107 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_1311_v106).fm14(_dafny.Set.fromElements(p2), (_this).f26, globalState));
          let _1313_v108;
          _1313_v108 = _dafny.Set.fromElements(new BigNumber((_1312_v107).length), (_this).f26);
          _1313_v108 = (_1313_v108).Difference(function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-465), new BigNumber(-266))) {
              let _1314_v109 = _compr_27;
              if (((new BigNumber(-465)).isLessThanOrEqualTo(_1314_v109)) && ((_1314_v109).isLessThan(new BigNumber(-266)))) {
                _coll27.add((_1314_v109).multipliedBy((_this).f27));
              }
            }
            return _coll27;
          }());
          let _1315_v110;
          _1315_v110 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1316_v111;
          _1316_v111 = _dafny.MultiSet.fromElements(_1315_v110);
          let _1317_v112;
          _1317_v112 = _dafny.Map.Empty.slice().updateUnsafe((_1311_v106).f32,new BigNumber((_dafny.Seq.UnicodeFromString("qtwjddafo")).length));
          let _1318_v113;
          _1318_v113 = _dafny.Seq.of(_1316_v111);
          let _1319_v114;
          let _nw229 = Array((new BigNumber(22)).toNumber());
          _nw229[(_dafny.ZERO).toNumber()] = _1316_v111;
          _nw229[(_dafny.ONE).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(2)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(3)).toNumber()] = (_1316_v111).Intersect(_1316_v111);
          _nw229[(new BigNumber(4)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(5)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(6)).toNumber()] = ((p1) ? (_1316_v111) : (_module.__default.fm35((_this).f27, new BigNumber((_1317_v112).length), globalState)));
          _nw229[(new BigNumber(7)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(8)).toNumber()] = (_1316_v111).Difference(_dafny.MultiSet.fromElements(_1315_v110));
          _nw229[(new BigNumber(9)).toNumber()] = (_1316_v111).Difference(_dafny.MultiSet.fromElements(_1315_v110));
          _nw229[(new BigNumber(10)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_1315_v110);
          _nw229[(new BigNumber(12)).toNumber()] = (((_1308_v103)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length))]) ? (_1316_v111) : (_dafny.MultiSet.fromElements(_1315_v110, _1315_v110, _1315_v110)));
          _nw229[(new BigNumber(13)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(14)).toNumber()] = (_1316_v111).update(_1315_v110, _module.__default.abs((_dafny.ZERO).minus((_this).f26)));
          _nw229[(new BigNumber(15)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(_1315_v110, (_this.f28)[_module.__default.safeIndex(new BigNumber((_1184_v0).length), new BigNumber((_this.f28).length))]);
          _nw229[(new BigNumber(17)).toNumber()] = (_1318_v113)[_module.__default.safeIndex(new BigNumber(((_dafny.MultiSet.fromElements((_1311_v106).f32, true)).update((_this).f25, _module.__default.abs((_this).f27))).cardinality()), new BigNumber((_1318_v113).length))];
          _nw229[(new BigNumber(18)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(19)).toNumber()] = _1316_v111;
          _nw229[(new BigNumber(20)).toNumber()] = (_1316_v111).Difference(_dafny.MultiSet.fromElements(_1315_v110, _1315_v110));
          _nw229[(new BigNumber(21)).toNumber()] = _1316_v111;
          _1319_v114 = _nw229;
          let _index236 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1319_v114).length));
          (_1319_v114)[_index236] = _1316_v111;
          let _index237 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1319_v114).length));
          (_1319_v114)[_index237] = _1316_v111;
          (globalState).f0 = (_this).f27;
          let _1320_v115;
          _1320_v115 = _dafny.Map.Empty.slice().updateUnsafe((_1309_v104)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_1309_v104).length))],(_this).f26);
          let _1321_v116;
          let _1322_v117;
          let _1323_v118;
          let _1324_v119;
          let _out29;
          let _out30;
          let _out31;
          let _out32;
          let _outcollector5 = (_this).m2(((_this).f26).isEqualTo(new BigNumber((_1320_v115).length)), globalState);
          _out29 = _outcollector5[0];
          _out30 = _outcollector5[1];
          _out31 = _outcollector5[2];
          _out32 = _outcollector5[3];
          _1321_v116 = _out29;
          _1322_v117 = _out30;
          _1323_v118 = _out31;
          _1324_v119 = _out32;
          r0 = p1;
        } else {
          let _1325_v120;
          _1325_v120 = _dafny.Set.fromElements((_this).f27);
          let _1326_v121;
          let _nw230 = Array((new BigNumber(10)).toNumber());
          _nw230[(_dafny.ZERO).toNumber()] = _module.__default.fm21(_this.f28, (_this).f27, new BigNumber((_1325_v120).length), globalState);
          _nw230[(_dafny.ONE).toNumber()] = p0;
          _nw230[(new BigNumber(2)).toNumber()] = (p0).update((_this).f26, _module.__default.abs(new BigNumber((_this.f28).length)));
          _nw230[(new BigNumber(3)).toNumber()] = p0;
          _nw230[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements((_this).f26);
          _nw230[(new BigNumber(5)).toNumber()] = p0;
          _nw230[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_1241_v55);
          _nw230[(new BigNumber(7)).toNumber()] = p0;
          _nw230[(new BigNumber(8)).toNumber()] = p0;
          _nw230[(new BigNumber(9)).toNumber()] = p0;
          _1326_v121 = _nw230;
          let _1327_v122;
          _1327_v122 = _module.D10.create_DC21(_1326_v121);
          let _index238 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length));
          let _rhs251 = _module.D10.create_DC21(_1326_v121);
          let _rhs252 = (_this).f25;
          let _lhs205 = _1308_v103;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length));
          _1327_v122 = _rhs251;
          _lhs205[_lhs206] = _rhs252;
          let _1328_v123;
          _1328_v123 = _dafny.Set.fromElements((_this).f25);
          let _1329_v124;
          _1329_v124 = _dafny.Seq.of(_1328_v123);
          let _1330_v125;
          _1330_v125 = _dafny.MultiSet.fromElements((_1311_v106).f32, (_this).f25, (_1311_v106).f32, (_1311_v106).fm14(_1328_v123, (_this).f27, globalState), (_1308_v103)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length))]);
          let _rhs253 = _1329_v124;
          let _rhs254 = p1;
          let _rhs255 = (((_1330_v125).equals(_dafny.MultiSet.fromElements((_1311_v106).f32, !(p1)))) ? (false) : (p1));
          let _rhs256 = (_dafny.ZERO).minus((_this).f26);
          let _lhs207 = globalState;
          let _lhs208 = globalState;
          let _lhs209 = globalState;
          _1329_v124 = _rhs253;
          _lhs207.f1 = _rhs254;
          _lhs208.f1 = _rhs255;
          _lhs209.f0 = _rhs256;
          let _1331_v126;
          let _nw231 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _1331_v126 = _nw231;
          let _index239 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1331_v126).length));
          (_1331_v126)[_index239] = _1184_v0;
          let _1332_v127;
          _1332_v127 = _1184_v0;
          let _index240 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1331_v126).length));
          let _rhs257 = _1332_v127;
          let _rhs258 = new BigNumber((_dafny.Seq.Concat(_this.f28, _this.f28)).length);
          let _rhs259 = ((p2) ? ((_dafny.MultiSet.fromElements((_1311_v106).f32, !((_1311_v106).f32))).Union(_module.__default.fm27(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_1309_v104)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_1309_v104).length))],(_1308_v103)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1308_v103).length))])).update((_1309_v104)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_1309_v104).length))], true)).length), globalState))) : (_dafny.MultiSet.FromArray(_1184_v0)));
          let _lhs210 = _1331_v126;
          let _lhs211 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1331_v126).length));
          let _lhs212 = globalState;
          _lhs210[_lhs211] = _rhs257;
          _lhs212.f0 = _rhs258;
          _1330_v125 = _rhs259;
          (globalState).f2 = (_this).f27;
          let _1333_v129;
          _1333_v129 = _dafny.Seq.of((_1325_v120).Union(_1325_v120), (_1325_v120).Difference(function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(141), new BigNumber(841))) {
              let _1334_v128 = _compr_28;
              if (((new BigNumber(141)).isLessThanOrEqualTo(_1334_v128)) && ((_1334_v128).isLessThan(new BigNumber(841)))) {
                _coll28.add(_module.__default.safeDivisionInt(_1334_v128, (_this).f27));
              }
            }
            return _coll28;
          }()), _1325_v120);
          let _rhs260 = new BigNumber(((_1333_v129)[_module.__default.safeIndex((_this).f26, new BigNumber((_1333_v129).length))]).length);
          let _rhs261 = (_1311_v106).f32;
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          _lhs213.f0 = _rhs260;
          _lhs214.f1 = _rhs261;
        }
      }
      let _1335_v130;
      let _nw232 = Array((new BigNumber(17)).toNumber()).fill(false);
      _1335_v130 = _nw232;
      let _index241 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1335_v130).length));
      (_1335_v130)[_index241] = p2;
      let _index242 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1335_v130).length));
      (_1335_v130)[_index242] = p1;
      let _1336_v131;
      _1336_v131 = _module.D1.create_DC4(p2, (_1335_v130)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1335_v130).length))]);
      let _1337_v132;
      _1337_v132 = _module.D1.create_DC5(_1336_v131);
      (globalState).f0 = function (_source18) {
        if (_source18.is_DC4) {
          let _1338___mcc_h4 = (_source18).cf7;
          let _1339___mcc_h5 = (_source18).cf8;
          let _1340_cf8 = _1339___mcc_h5;
          let _1341_cf7 = _1338___mcc_h4;
          return new BigNumber(((_dafny.Set.fromElements(new BigNumber((_this.f28).length), (_this).f27)).Intersect(_dafny.Set.fromElements((_this).f26))).length);
        } else if (_source18.is_DC3) {
          let _1342___mcc_h6 = (_source18).cf6;
          let _1343_cf6 = _1342___mcc_h6;
          return (_this).f27;
        } else {
          let _1344___mcc_h7 = (_source18).cf9;
          let _1345_cf9 = _1344___mcc_h7;
          return (_this).f27;
        }
      }(_1337_v132);
      let _source19 = _1184_v0;
      let _1346___mcc_h8 = _source19;
      let _1347_cf26 = _1346___mcc_h8;
      (globalState).f0 = (_this).f27;
      (globalState).f1 = ((_this).f26).isLessThanOrEqualTo((_this).f26);
      let _1348_v133;
      _1348_v133 = new _dafny.CodePoint('q'.codePointAt(0));
      _1348_v133 = ((p2) ? (_1348_v133) : (_1348_v133));
      (globalState).f0 = (_this).f27;
      r0 = (((_1335_v130)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1335_v130).length))]) ? (p2) : ((_1335_v130)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1335_v130).length))]));
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.Map.Empty;
      let _1349_v0;
      _1349_v0 = _dafny.Seq.of((_this).f25, (_this).f25, p0, p0, !((_this).f25));
      let _1350_v1;
      _1350_v1 = _dafny.Seq.of(_1349_v0);
      let _hi12 = ((_this).f27).minus((_dafny.ZERO).minus(new BigNumber((_1350_v1).length)));
      for (let _1351_i0 = (_this).f27; _1351_i0.isLessThan(_hi12); _1351_i0 = _1351_i0.plus(_dafny.ONE)) {
        let _1352_v2;
        _1352_v2 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1353_v3;
        _1353_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1352_v2,(_this).f25);
        let _1354_v4;
        let _nw233 = Array((new BigNumber(16)).toNumber());
        _nw233[(_dafny.ZERO).toNumber()] = (_this).f25;
        _nw233[(_dafny.ONE).toNumber()] = p0;
        _nw233[(new BigNumber(2)).toNumber()] = false;
        _nw233[(new BigNumber(3)).toNumber()] = !(p0);
        _nw233[(new BigNumber(4)).toNumber()] = p0;
        _nw233[(new BigNumber(5)).toNumber()] = (_this).f25;
        _nw233[(new BigNumber(6)).toNumber()] = p0;
        _nw233[(new BigNumber(7)).toNumber()] = p0;
        _nw233[(new BigNumber(8)).toNumber()] = (_this).f25;
        _nw233[(new BigNumber(9)).toNumber()] = p0;
        _nw233[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw233[(new BigNumber(11)).toNumber()] = (_this).f25;
        _nw233[(new BigNumber(12)).toNumber()] = (((_1353_v3).contains(_1352_v2)) ? ((_1353_v3).get(_1352_v2)) : (false));
        _nw233[(new BigNumber(13)).toNumber()] = p0;
        _nw233[(new BigNumber(14)).toNumber()] = (_this).f25;
        _nw233[(new BigNumber(15)).toNumber()] = (_this).f25;
        _1354_v4 = _nw233;
        let _1355_v5;
        _1355_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_1354_v4);
        let _1356_v6;
        _1356_v6 = _dafny.Map.Empty.slice().updateUnsafe((((_1355_v5).contains(_this.f28)) ? ((_1355_v5).get(_this.f28)) : (_1354_v4)),p0);
        _1356_v6 = (_1356_v6).update(_1354_v4, true);
        let _rhs262 = (_this).f25;
        let _rhs263 = true;
        let _lhs215 = globalState;
        let _lhs216 = globalState;
        _lhs215.f1 = _rhs262;
        _lhs216.f1 = _rhs263;
        let _1357_v7;
        let _nw234 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _1357_v7 = _nw234;
        let _1358_v8;
        let _nw235 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _1358_v8 = _nw235;
        let _1359_v9;
        _1359_v9 = _dafny.Seq.of(_1358_v8);
        let _index243 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1357_v7).length));
        (_1357_v7)[_index243] = _1359_v9;
        let _index244 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_1354_v4).length));
        (_1354_v4)[_index244] = (_this).f25;
        let _1360_v10;
        _1360_v10 = _module.D0.create_DC1(_1351_i0, _this.f28, (_this).f25, (_this).f26, (_this).f25);
        let _index245 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1357_v7).length));
        let _index246 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_1354_v4).length));
        let _rhs264 = true;
        let _rhs265 = _dafny.Seq.Concat(_1359_v9, _dafny.Seq.of(_1358_v8, _1358_v8, _1358_v8));
        let _rhs266 = p0;
        let _rhs267 = _module.__default.fm4(_1360_v10, (_this).f25, _module.__default.safeDivisionInt((_this).f26, (_this).f27), globalState);
        let _rhs268 = (_1349_v0)[_module.__default.safeIndex(((_this).f27).minus((_this).f27), new BigNumber((_1349_v0).length))];
        let _lhs217 = _1357_v7;
        let _lhs218 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1357_v7).length));
        let _lhs219 = _1354_v4;
        let _lhs220 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_1354_v4).length));
        let _lhs221 = globalState;
        r1 = _rhs264;
        _lhs217[_lhs218] = _rhs265;
        _lhs219[_lhs220] = _rhs266;
        _1352_v2 = _rhs267;
        _lhs221.f1 = _rhs268;
        _1354_v4 = _1354_v4;
      }
      let _1361_v11;
      _1361_v11 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1362_v12;
      _1362_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1361_v11,p0);
      let _1363_v13;
      _1363_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1362_v12,(_this).f26);
      _1363_v13 = _1363_v13;
      let _1364_v14;
      _1364_v14 = _module.D1.create_DC4((_1349_v0)[_module.__default.safeIndex(_module.__default.fm2((_this).f26, (_this).f26, globalState), new BigNumber((_1349_v0).length))], p0);
      let _1365_v15;
      _1365_v15 = _dafny.Seq.of(_1364_v14);
      let _1366_v16;
      _1366_v16 = _module.D0.create_DC2();
      let _1367_v17;
      _1367_v17 = _module.D8.create_DC17(_1366_v16, p0);
      let _pat_let_tv34 = p0;
      _1365_v15 = _dafny.Seq.of((((_1367_v17).dtor_cf29) ? (_1364_v14) : (function (_pat_let40_0) {
        return function (_1368_dt__update__tmp_h0) {
          return function (_pat_let41_0) {
            return function (_1369_dt__update_hcf7_h0) {
              return _module.D1.create_DC4(_1369_dt__update_hcf7_h0, (_1368_dt__update__tmp_h0).dtor_cf8);
            }(_pat_let41_0);
          }(_pat_let_tv34);
        }(_pat_let40_0);
      }(_1364_v14))));
      let _1370_v18;
      _1370_v18 = _dafny.MultiSet.fromElements((_this).f26, new BigNumber((_this.f28).length));
      let _1371_v19;
      _1371_v19 = _dafny.Set.fromElements(p0);
      let _1372_v20;
      _1372_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1370_v18,(_1371_v19).IsDisjointFrom(_1371_v19));
      _1372_v20 = function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f27)), _1370_v18)).Elements) {
          let _1373_v21 = _compr_29;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f27)), _1370_v18), _1373_v21)) {
            _coll29.push([_1373_v21,p0]);
          }
        }
        return _coll29;
      }();
      r0 = (_this).f26;
      r1 = p0;
      r0 = (_dafny.ZERO).minus(_module.__default.fm2((_this).f27, new BigNumber((((p0) ? (_this.f28) : (_this.f28))).length), globalState));
      r1 = (_this).f25;
      let _1374_v22;
      _1374_v22 = _dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_this).f25);
      let _1375_v23;
      _1375_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v22,(_this).f27);
      r2 = (_1375_v23).Merge((((_this).f25) ? (_1375_v23) : (_1375_v23)));
      let _1376_v24;
      _1376_v24 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f27).plus((_this).f26),((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), function (_1377_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })) : (_this.f28)));
      r3 = _1376_v24;
      return [r0, r1, r2, r3];
    }
    m3(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let r2 = _dafny.ZERO;
      if ((_this).f25) {
        r2 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(398)), function (_1378_i0) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length);
        let _1379_v0;
        let _init42 = function (_1380_i1) {
          return (_1380_i1).minus(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-692)), function (_1381_i2) {
            return (_this).f27;
          })), _dafny.MultiSet.fromElements(new BigNumber(659)), _dafny.MultiSet.fromElements(new BigNumber(-667), (_this).f27), _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f27)))).length));
        };
        let _nw236 = Array((new BigNumber(20)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw236.length); _i0_42++) {
          _nw236[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1379_v0 = _nw236;
        let _index247 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
        (_1379_v0)[_index247] = (_dafny.ZERO).minus((_this).f27);
        let _1382_v1;
        _1382_v1 = _module.D12.create_DC28(_1379_v0);
        let _1383_v2;
        let _nw237 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1383_v2 = _nw237;
        let _index248 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length));
        (_1383_v2)[_index248] = true;
        let _1384_v3;
        _1384_v3 = _dafny.Set.fromElements((_this).f25, (_this).f25, (_this).f25);
        let _index249 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
        let _index250 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length));
        let _rhs269 = (((_this).f25) ? ((_this).f26) : (new BigNumber(229)));
        let _rhs270 = (_this).f27;
        let _rhs271 = _this.f28;
        let _rhs272 = _1382_v1;
        let _rhs273 = (_1384_v3).IsSubsetOf(_dafny.Set.fromElements(!((_this).f25)));
        let _lhs222 = globalState;
        let _lhs223 = _1379_v0;
        let _lhs224 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
        let _lhs225 = _this;
        let _lhs226 = _1383_v2;
        let _lhs227 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length));
        _lhs222.f0 = _rhs269;
        _lhs223[_lhs224] = _rhs270;
        _lhs225.f28 = _rhs271;
        _1382_v1 = _rhs272;
        _lhs226[_lhs227] = _rhs273;
        (_this).f28 = _this.f28;
        let _index251 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
        let _rhs274 = (_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))];
        let _rhs275 = (_this).fm6((_this).f25, (_this).f25, (_this).f26, globalState);
        let _lhs228 = _1379_v0;
        let _lhs229 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
        let _lhs230 = globalState;
        _lhs228[_lhs229] = _rhs274;
        _lhs230.f2 = _rhs275;
        if (((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))]).isLessThanOrEqualTo((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))])) {
          let _1385_v4;
          _1385_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_1383_v2)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length))]);
          let _1386_v5;
          let _nw238 = new _module.C1();
          _nw238.__ctor((_this).f25);
          _1386_v5 = _nw238;
          let _1387_v6;
          _1387_v6 = _dafny.MultiSet.fromElements(_1386_v5);
          let _1388_v8;
          _1388_v8 = _module.D0.create_DC0((_1383_v2)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length))]);
          let _1389_v9;
          _1389_v9 = _dafny.Set.fromElements(_1388_v8);
          let _1390_v10;
          _1390_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(((_1385_v4).contains((((_1387_v6).contains(_1386_v5)) ? ((_1387_v6).get(_1386_v5)) : ((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))])))) ? ((_1385_v4).get((((_1387_v6).contains(_1386_v5)) ? ((_1387_v6).get(_1386_v5)) : ((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))])))) : (_module.__default.fm5((_module.D1.create_DC4((_this).f25, (_this).f25)).dtor_cf8, function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of (_1389_v9).Elements) {
              let _1391_v7 = _compr_30;
              if ((_1389_v9).contains(_1391_v7)) {
                _coll30.push([_1391_v7,(_this).f26]);
              }
            }
            return _coll30;
          }(), (_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))], globalState))));
          let _1392_v11;
          _1392_v11 = _dafny.Seq.of((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))], (_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))]);
          let _1393_v12;
          _1393_v12 = _dafny.Seq.of((_1386_v5).f32);
          _1390_v10 = (_1390_v10).update((new BigNumber((_1390_v10).length)).plus(new BigNumber((_dafny.Seq.update(_1392_v11, _module.__default.safeIndex((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))], new BigNumber((_1392_v11).length)), (_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))])).length)), (((_1383_v2)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length))]) ? ((_1393_v12)[_module.__default.safeIndex((_this).f27, new BigNumber((_1393_v12).length))]) : ((_1383_v2)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length))])));
          let _index252 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length));
          (_1383_v2)[_index252] = !((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))]).isEqualTo((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))]);
          let _1394_v13;
          _1394_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1379_v0,(_this).f25);
          let _index253 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
          (_1379_v0)[_index253] = new BigNumber((_1394_v13).length);
          let _index254 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length));
          (_1383_v2)[_index254] = !((_this).f25);
          _1383_v2 = _1383_v2;
        } else {
          let _index255 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length));
          (_1379_v0)[_index255] = (_dafny.ZERO).minus((_this).f26);
          let _1395_v14;
          _1395_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f27);
          let _1396_v15;
          _1396_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_1395_v14).contains((_this).f25)) ? ((_1395_v14).get((_this).f25)) : (new BigNumber((_this.f28).length))),_this.f28);
          _1396_v15 = (_1396_v15).update((_1379_v0)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1379_v0).length))], _dafny.Seq.Concat(_this.f28, _this.f28));
          let _1397_v16;
          _1397_v16 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1398_v17;
          _1398_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(((_1383_v2)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length))]) ? ((_this.f28)[_module.__default.safeIndex((_this).f26, new BigNumber((_this.f28).length))]) : (_1397_v16)));
          let _1399_v18;
          _1399_v18 = _dafny.Seq.of(false);
          _1398_v17 = (_1398_v17).update((_1383_v2)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1383_v2).length))], (((_1399_v18)[_module.__default.safeIndex((_this).f27, new BigNumber((_1399_v18).length))]) ? (_1397_v16) : (_1397_v16)));
          let _1400_v19;
          let _nw239 = new _module.C1();
          _nw239.__ctor(false);
          _1400_v19 = _nw239;
          let _1401_v20;
          _1401_v20 = _dafny.Seq.of((_this).f26, (_this).f26);
          let _1402_v21;
          _1402_v21 = _module.D11.create_DC26((_this).f26, new BigNumber(797), (_this).f25);
          _1401_v20 = _module.__default.fm3((_1402_v21).dtor_cf41, !(true), (_this).f26, (_this).f25, globalState);
        }
      } else {
        let _1403_v22;
        _1403_v22 = _module.D0.create_DC0((_this).f25);
        let _1404_v23;
        _1404_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(true, _dafny.Map.Empty.slice().updateUnsafe(_1403_v22,(_this).f26), (_this).f27, globalState),(_this).f25);
        _1404_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,!((_this).f25));
        (globalState).f0 = (_this).f27;
        let _1405_v24;
        _1405_v24 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1406_v25;
        _1406_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v24,(_this).f27);
        r1 = (((_this).f27).minus((((_1406_v25).contains(new _dafny.CodePoint('f'.codePointAt(0)))) ? ((_1406_v25).get(new _dafny.CodePoint('f'.codePointAt(0)))) : ((_this).f27)))).isLessThanOrEqualTo((_this).f26);
        let _1407_v26;
        let _nw240 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1407_v26 = _nw240;
        let _index256 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1407_v26).length));
        (_1407_v26)[_index256] = (_this).f26;
        let _index257 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1407_v26).length));
        (_1407_v26)[_index257] = (_this).f26;
        if ((_this).f25) {
          let _rhs276 = _dafny.Seq.Concat(_this.f28, _dafny.Seq.UnicodeFromString("y"));
          let _rhs277 = (_this).f26;
          let _lhs231 = _this;
          let _lhs232 = globalState;
          _lhs231.f28 = _rhs276;
          _lhs232.f0 = _rhs277;
          let _1408_v27;
          let _nw241 = new _module.C2();
          _nw241.__ctor((_this).f27, false, (_this).f26);
          _1408_v27 = _nw241;
          let _1409_v28;
          let _nw242 = new _module.C1();
          _nw242.__ctor(!(false));
          _1409_v28 = _nw242;
          let _1410_v29;
          let _init43 = function (_1411_i3) {
            return _dafny.Seq.Concat(_this.f28, _this.f28);
          };
          let _nw243 = Array((new BigNumber(12)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw243.length); _i0_43++) {
            _nw243[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1410_v29 = _nw243;
          let _index258 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1410_v29).length));
          (_1410_v29)[_index258] = _this.f28;
          let _index259 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1410_v29).length));
          (_1410_v29)[_index259] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rgnv"), _dafny.Seq.UnicodeFromString("cqotdmdy")), _dafny.Seq.Concat(_this.f28, _this.f28));
          let _1412_v30;
          _1412_v30 = _dafny.Seq.of((_this).f25);
          let _1413_v31;
          let _nw244 = new _module.C3();
          _nw244.__ctor(_this.f28, (_this).f25, (_this).f25, new BigNumber((_1412_v30).length));
          _1413_v31 = _nw244;
          let _1414_v32;
          _1414_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1413_v31,true);
          let _1415_v33;
          _1415_v33 = _dafny.Seq.of(_1413_v31);
          let _1416_v34;
          _1416_v34 = _dafny.MultiSet.fromElements((_this).f25, true, false);
          let _1417_v35;
          _1417_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1410_v29)[_module.__default.safeIndex(new BigNumber(838), new BigNumber((_1410_v29).length))]).length),_1416_v34);
          _1414_v32 = (_1414_v32).update((_1415_v33)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((((_1417_v35).contains((_this).f26)) ? ((_1417_v35).get((_this).f26)) : (_1416_v34))).cardinality())), new BigNumber((_1415_v33).length))], (_1409_v28).f32);
        } else {
          let _1418_v36;
          _1418_v36 = _module.D1.create_DC4((_this).f25, (_this).f25);
          (globalState).f1 = (_1418_v36).dtor_cf7;
          (globalState).f1 = (_this).f25;
          (globalState).f2 = _module.__default.safeModuloInt(new BigNumber((_this.f28).length), new BigNumber(-450));
          let _1419_v37;
          _1419_v37 = _dafny.Seq.of((_this).f26);
          _1419_v37 = _1419_v37;
          let _1420_v38;
          _1420_v38 = _dafny.Set.fromElements((_this).f25);
          (globalState).f1 = ((_dafny.Set.fromElements((_this).f25)).Difference(_1420_v38)).IsSubsetOf(_1420_v38);
        }
      }
      let _1421_v39;
      _1421_v39 = new _dafny.CodePoint('o'.codePointAt(0));
      let _1422_v40;
      _1422_v40 = _module.D1.create_DC3(_1421_v39);
      if (function (_source20) {
        if (_source20.is_DC4) {
          let _1423___mcc_h0 = (_source20).cf7;
          let _1424___mcc_h1 = (_source20).cf8;
          let _1425_cf8 = _1424___mcc_h1;
          let _1426_cf7 = _1423___mcc_h0;
          return (function () {
            let _coll31 = new _dafny.Set();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(-349), new BigNumber(15))) {
              let _1427_v41 = _compr_31;
              if (((new BigNumber(-349)).isLessThanOrEqualTo(_1427_v41)) && ((_1427_v41).isLessThan(new BigNumber(15)))) {
                _coll31.add((_1427_v41).minus(new BigNumber((_dafny.Seq.of((_this).f27, (_this).f27)).length)));
              }
            }
            return _coll31;
          }()).IsProperSubsetOf(_dafny.Set.fromElements((_this).f27));
        } else if (_source20.is_DC3) {
          let _1428___mcc_h2 = (_source20).cf6;
          let _1429_cf6 = _1428___mcc_h2;
          return (_this).f25;
        } else {
          let _1430___mcc_h3 = (_source20).cf9;
          let _1431_cf9 = _1430___mcc_h3;
          return (_this).f25;
        }
      }(_1422_v40)) {
        _1421_v39 = _1421_v39;
        (globalState).f0 = (_this).f26;
        let _1432_v42;
        let _init44 = function (_1433_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
        };
        let _nw245 = Array((new BigNumber(19)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw245.length); _i0_44++) {
          _nw245[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1432_v42 = _nw245;
        let _1434_v43;
        _1434_v43 = _module.D8.create_DC16(_1432_v42);
        let _1435_v44;
        let _nw246 = Array((new BigNumber(7)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = _module.D8.create_DC16(_1432_v42);
        _nw246[(_dafny.ONE).toNumber()] = _1434_v43;
        _nw246[(new BigNumber(2)).toNumber()] = _1434_v43;
        _nw246[(new BigNumber(3)).toNumber()] = _1434_v43;
        _nw246[(new BigNumber(4)).toNumber()] = _1434_v43;
        _nw246[(new BigNumber(5)).toNumber()] = _1434_v43;
        _nw246[(new BigNumber(6)).toNumber()] = _1434_v43;
        _1435_v44 = _nw246;
        let _1436_v45;
        _1436_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1435_v44);
        let _1437_v46;
        _1437_v46 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1435_v44));
        _1436_v45 = (_1437_v46)[_module.__default.safeIndex((_this).f27, new BigNumber((_1437_v46).length))];
        let _1438_v47;
        _1438_v47 = _dafny.Seq.of((_this).f27, (_this).f27);
        (globalState).f2 = (new BigNumber(703)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1438_v47,(_this).f25)).length));
        let _1439_v48;
        let _nw247 = new _module.C3();
        _nw247.__ctor(_this.f28, (_this).f25, (_this).f25, new BigNumber(-752));
        _1439_v48 = _nw247;
      } else {
        let _1440_v49;
        _1440_v49 = _dafny.Set.fromElements((_this).f25);
        if (((_1440_v49).Union(_1440_v49)).IsDisjointFrom(_dafny.Set.fromElements((_this).f25, (_this).f25))) {
          let _1441_v50;
          let _nw248 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1441_v50 = _nw248;
          _1441_v50 = _1441_v50;
          let _1442_v51;
          _1442_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
          _1442_v51 = (_1442_v51).update((_this).f25, !((_this).f25));
          let _1443_v52;
          let _init45 = function (_1444_i5) {
            return false;
          };
          let _nw249 = Array((new BigNumber(3)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw249.length); _i0_45++) {
            _nw249[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1443_v52 = _nw249;
          let _1445_v53;
          let _nw250 = new _module.C0();
          _nw250.__ctor(_this.f28, (_this).f25, (_this).f26, _1443_v52, (_this).f25);
          _1445_v53 = _nw250;
          let _1446_v54;
          let _nw251 = new _module.C3();
          _nw251.__ctor(_this.f28, (_this).f25, (_this).f25, (_this).f26);
          _1446_v54 = _nw251;
          let _1447_v55;
          let _nw252 = Array((new BigNumber(12)).toNumber());
          _nw252[(_dafny.ZERO).toNumber()] = _1446_v54;
          _nw252[(_dafny.ONE).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(2)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(3)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(4)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(5)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(6)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(7)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(8)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(9)).toNumber()] = ((false) ? (_1446_v54) : (_1446_v54));
          _nw252[(new BigNumber(10)).toNumber()] = _1446_v54;
          _nw252[(new BigNumber(11)).toNumber()] = _1446_v54;
          _1447_v55 = _nw252;
          let _1448_v56;
          _1448_v56 = _dafny.Seq.of(_1443_v52);
          let _1449_v57;
          _1449_v57 = _module.D0.create_DC0((_1446_v54).f30);
          let _1450_v58;
          _1450_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1449_v57,(_dafny.ZERO).minus((_this).f26));
          let _1451_v59;
          _1451_v59 = _dafny.MultiSet.fromElements((_this).f25);
          let _1452_v60;
          _1452_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1451_v59).cardinality()),(_this).f26);
          let _rhs278 = (_this).f26;
          let _rhs279 = ((_dafny.Seq.contains(_1448_v56, _1443_v52)) ? (_1447_v55) : (_1447_v55));
          let _rhs280 = _module.__default.fm5((_1446_v54).f30, _1450_v58, (_this).f27, globalState);
          let _rhs281 = (_dafny.ZERO).minus(((!(_1452_v60).contains((_1446_v54).fm9(globalState))) ? ((((_1446_v54).f30) ? ((_this).f26) : ((_this).f26))) : (_module.__default.safeModuloInt((_this).f27, (_this).f26))));
          let _lhs233 = globalState;
          let _lhs234 = globalState;
          let _lhs235 = globalState;
          _lhs233.f0 = _rhs278;
          _1447_v55 = _rhs279;
          _lhs234.f1 = _rhs280;
          _lhs235.f2 = _rhs281;
          let _1453_v61;
          let _init46 = ((_1454_v39) => function (_1455_i6) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(496)), ((_1456_v39) => function (_1457_i7) {
              return _1456_v39;
            })(_1454_v39));
          })(_1421_v39);
          let _nw253 = Array((new BigNumber(23)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw253.length); _i0_46++) {
            _nw253[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1453_v61 = _nw253;
          _1453_v61 = _1453_v61;
        } else {
          let _1458_v62;
          _1458_v62 = _dafny.MultiSet.fromElements((_this).f27);
          let _1459_v63;
          let _nw254 = new _module.C2();
          _nw254.__ctor(((((_1458_v62).contains(new BigNumber(824))) ? ((_1458_v62).get(new BigNumber(824))) : ((_this).f26))).plus((_this).f27), (_this).f25, (_this).f27);
          _1459_v63 = _nw254;
          _1459_v63 = _1459_v63;
          (globalState).f0 = (_module.__default.safeDivisionInt((_1459_v63).f31, (_this).f27)).multipliedBy((_1459_v63).f31);
          let _1460_v64;
          let _nw255 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _1460_v64 = _nw255;
          let _1461_v65;
          let _nw256 = Array((new BigNumber(9)).toNumber());
          _nw256[(_dafny.ZERO).toNumber()] = (_this).f27;
          _nw256[(_dafny.ONE).toNumber()] = (_this).f26;
          _nw256[(new BigNumber(2)).toNumber()] = (_1459_v63).f31;
          _nw256[(new BigNumber(3)).toNumber()] = (_1459_v63).f31;
          _nw256[(new BigNumber(4)).toNumber()] = (_1459_v63).f31;
          _nw256[(new BigNumber(5)).toNumber()] = (_this).f27;
          _nw256[(new BigNumber(6)).toNumber()] = (_this).f27;
          _nw256[(new BigNumber(7)).toNumber()] = (_1459_v63).f31;
          _nw256[(new BigNumber(8)).toNumber()] = new BigNumber(514);
          _1461_v65 = _nw256;
          let _1462_v66;
          _1462_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v65,(_this).f25);
          let _index260 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1460_v64).length));
          (_1460_v64)[_index260] = _1462_v66;
          let _index261 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1460_v64).length));
          (_1460_v64)[_index261] = _1462_v66;
          (globalState).f0 = (((_1458_v62).contains(new BigNumber(271))) ? ((_1458_v62).get(new BigNumber(271))) : ((_this).f27));
          let _1463_v67;
          _1463_v67 = _dafny.MultiSet.fromElements(_this.f28, _this.f28);
          _1463_v67 = (_1463_v67).Difference((_1463_v67).Intersect(_1463_v67));
        }
        let _1464_v68;
        _1464_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f25);
        _1464_v68 = (_1464_v68).update(new BigNumber((_this.f28).length), !((_this).f25));
        let _1465_v69;
        _1465_v69 = _dafny.Seq.of(_this.f28);
        let _1466_v70;
        let _nw257 = new _module.C3();
        _nw257.__ctor(_dafny.Seq.Concat(_this.f28, (_1465_v69)[_module.__default.safeIndex((_this).f26, new BigNumber((_1465_v69).length))]), (_this).f25, (_this).f25, (_this).f27);
        _1466_v70 = _nw257;
        let _1467_v71;
        let _nw258 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _1467_v71 = _nw258;
        let _1468_v72;
        _1468_v72 = _dafny.Seq.of((_1466_v70).f30);
        let _index262 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1467_v71).length));
        (_1467_v71)[_index262] = _dafny.Seq.Concat(_1468_v72, _1468_v72);
        let _index263 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1467_v71).length));
        (_1467_v71)[_index263] = _1468_v72;
        let _1469_v73;
        _1469_v73 = _dafny.MultiSet.fromElements(new BigNumber((_this.f28).length), (_this).f26, (_this).fm6((_this).f25, (_1466_v70).f30, (_dafny.ZERO).minus((_this).f27), globalState));
        (globalState).f0 = (((_1469_v73).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1440_v49,(_this).fm7(globalState))).length))) ? ((_1469_v73).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1440_v49,(_this).fm7(globalState))).length))) : (new BigNumber((_dafny.Seq.Concat(_this.f28, _this.f28)).length)));
      }
      let _hi13 = new BigNumber(801);
      for (let _1470_i8 = ((_this).f26).plus((_this).f27); _1470_i8.isLessThan(_hi13); _1470_i8 = _1470_i8.plus(_dafny.ONE)) {
        let _1471_v74;
        let _init47 = function (_1472_i9) {
          return (_this).f25;
        };
        let _nw259 = Array((new BigNumber(4)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw259.length); _i0_47++) {
          _nw259[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1471_v74 = _nw259;
        let _index264 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length));
        (_1471_v74)[_index264] = (_this).f25;
        let _1473_v75;
        let _nw260 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1473_v75 = _nw260;
        let _1474_v76;
        _1474_v76 = _dafny.MultiSet.fromElements(_1471_v74, _1471_v74, _1471_v74);
        let _index265 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1473_v75).length));
        (_1473_v75)[_index265] = _1474_v76;
        let _1475_v77;
        let _nw261 = Array((new BigNumber(13)).toNumber()).fill([]);
        _1475_v77 = _nw261;
        let _1476_v78;
        let _nw262 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1476_v78 = _nw262;
        let _index266 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1475_v77).length));
        (_1475_v77)[_index266] = _1476_v78;
        let _1477_v79;
        _1477_v79 = _dafny.Seq.of((_this).f25, (((_this).f25) ? (!((_this).f25)) : ((_this).f25)), ((_this).f25) || ((_this).f25));
        let _1478_v80;
        _1478_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1476_v78);
        let _1479_v81;
        _1479_v81 = _dafny.Set.fromElements((_this).f25);
        let _index267 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length));
        let _index268 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1473_v75).length));
        let _index269 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1475_v77).length));
        let _rhs282 = !((_this).f25);
        let _rhs283 = (_1477_v79)[_module.__default.safeIndex((_this).f27, new BigNumber((_1477_v79).length))];
        let _rhs284 = _1474_v76;
        let _rhs285 = (((_1478_v80).contains(new BigNumber((_module.__default.fm36((_this).f27, new BigNumber((_this.f28).length), _this.f28, (_this).f27, globalState)).length))) ? ((_1478_v80).get(new BigNumber((_module.__default.fm36((_this).f27, new BigNumber((_this.f28).length), _this.f28, (_this).f27, globalState)).length))) : ((((_this).f25) ? (_1476_v78) : (_1476_v78))));
        let _rhs286 = new BigNumber((_1479_v81).length);
        let _lhs236 = _1471_v74;
        let _lhs237 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length));
        let _lhs238 = globalState;
        let _lhs239 = _1473_v75;
        let _lhs240 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1473_v75).length));
        let _lhs241 = _1475_v77;
        let _lhs242 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1475_v77).length));
        let _lhs243 = globalState;
        _lhs236[_lhs237] = _rhs282;
        _lhs238.f1 = _rhs283;
        _lhs239[_lhs240] = _rhs284;
        _lhs241[_lhs242] = _rhs285;
        _lhs243.f0 = _rhs286;
        (globalState).f2 = _module.__default.fm2((_dafny.ZERO).minus(_module.__default.fm2((_this).f26, (_this).f27, globalState)), (_this).f27, globalState);
        let _1480_v82;
        let _nw263 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
        _1480_v82 = _nw263;
        let _1481_v83;
        _1481_v83 = _module.D0.create_DC0((_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))]);
        let _1482_v84;
        _1482_v84 = _dafny.Map.Empty.slice().updateUnsafe((_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))],(_this).f25);
        let _1483_v85;
        _1483_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1481_v83,new BigNumber(((_1482_v84).update(false, (_this).f25)).length));
        let _index270 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1480_v82).length));
        (_1480_v82)[_index270] = (_1479_v81).Intersect(_dafny.Set.fromElements((_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))], (_this).f25, (_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))], (_this).f25, _module.__default.fm5((_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))], _1483_v85, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-270)), function (_1484_i10) {
          return (_this).f27;
        })).length), globalState)));
        let _index271 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1480_v82).length));
        (_1480_v82)[_index271] = _1479_v81;
        let _1485_v86;
        _1485_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(605),(_this).f27);
        let _1486_v87;
        _1486_v87 = _dafny.MultiSet.fromElements((_this).f26);
        let _1487_v88;
        _1487_v88 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1480_v82)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1480_v82).length))]).length),_1486_v87);
        let _1488_v89;
        let _nw264 = new _module.C2();
        _nw264.__ctor((_this).f27, (_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))], (_dafny.ZERO).minus(new BigNumber((_1487_v88).length)));
        _1488_v89 = _nw264;
        let _1489_v90;
        _1489_v90 = _dafny.Set.fromElements(_1488_v89);
        let _1490_v91;
        _1490_v91 = _dafny.Set.fromElements((_this).f26, _1470_i8, new BigNumber((_1489_v90).length));
        let _1491_v92;
        _1491_v92 = _dafny.Seq.of(new BigNumber((_1490_v91).length), _1470_i8, (_this).f26, (_1488_v89).f31, (_1488_v89).f31);
        (globalState).f2 = _module.__default.fm2(new BigNumber((_dafny.Seq.Concat(_module.__default.fm3((_this).f27, (_1471_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1471_v74).length))], new BigNumber((_1485_v86).length), !(!((_this).f25)), globalState), _1491_v92)).length), (_this).f27, globalState);
      }
      if (!(!((_this).f25) || (!((_this).f25))) || ((((_this).f25) ? (false) : ((_this).f25)))) {
        (globalState).f1 = (_this).f25;
        let _1492_v93;
        _1492_v93 = _dafny.Seq.of(false);
        let _1493_v94;
        _1493_v94 = _dafny.Map.Empty.slice().updateUnsafe(true,(_1492_v93)[_module.__default.safeIndex((_this).f27, new BigNumber((_1492_v93).length))]);
        _1493_v94 = (_1493_v94).update((_this).f25, (_this).f25);
        let _1494_v95;
        let _nw265 = new _module.C2();
        _nw265.__ctor((_this).f26, (_this).f25, (_this).f26);
        _1494_v95 = _nw265;
        let _1495_v96;
        _1495_v96 = _dafny.Seq.of(_1494_v95, _1494_v95);
        let _1496_v97;
        _1496_v97 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f25),(_1495_v96)[_module.__default.safeIndex(new BigNumber(-67), new BigNumber((_1495_v96).length))]);
        let _1497_v98;
        _1497_v98 = _dafny.Seq.of((((_1496_v97).contains((_this).f25)) ? ((_1496_v97).get((_this).f25)) : (_1494_v95)));
        let _1498_v99;
        _1498_v99 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_1497_v98, _1495_v96));
        let _1499_v100;
        _1499_v100 = (_1498_v99).update(_1497_v98, _module.__default.abs((_this).f26));
        let _1500_v101;
        _1500_v101 = _dafny.Seq.of(_1497_v98);
        _1498_v99 = ((_1498_v99).Union(_dafny.MultiSet.fromElements(_1495_v96))).Union((_dafny.MultiSet.FromArray(_1500_v101)));
        let _1501_v102;
        _1501_v102 = _dafny.Map.Empty.slice().updateUnsafe((_1494_v95).f31,(_this).f25);
        _1501_v102 = (_1501_v102).update((_1494_v95).f31, true);
        (globalState).f1 = (_this).f25;
      } else {
        let _1502_v103;
        _1502_v103 = _dafny.MultiSet.fromElements((_this).f26);
        let _1503_v104;
        _1503_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(541),new BigNumber((_1502_v103).cardinality()));
        let _1504_v105;
        _1504_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1503_v104,(_this).f25);
        let _1505_v106;
        _1505_v106 = _dafny.Seq.of((_this).f25, (_this).f25, (((_1504_v105).contains(_1503_v104)) ? ((_1504_v105).get(_1503_v104)) : (!((_this).f25))), (_this).f25);
        if (_dafny.Seq.IsPrefixOf(_1505_v106, _1505_v106)) {
          let _1506_v107;
          let _nw266 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _1506_v107 = _nw266;
          let _1507_v108;
          _1507_v108 = _module.D12.create_DC28(_1506_v107);
          let _1508_v109;
          let _nw267 = Array((new BigNumber(27)).toNumber());
          _nw267[(_dafny.ZERO).toNumber()] = _1506_v107;
          _nw267[(_dafny.ONE).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(2)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(3)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(4)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(5)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(6)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(7)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(8)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(9)).toNumber()] = ((false) ? (_1506_v107) : (_1506_v107));
          _nw267[(new BigNumber(10)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(11)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(12)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(13)).toNumber()] = (_1507_v108).dtor_cf44;
          _nw267[(new BigNumber(14)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(15)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(16)).toNumber()] = (((_this).f25) ? (_1506_v107) : (_1506_v107));
          _nw267[(new BigNumber(17)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(18)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(19)).toNumber()] = (((_this).f25) ? (_1506_v107) : (_1506_v107));
          _nw267[(new BigNumber(20)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(21)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(22)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(23)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(24)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(25)).toNumber()] = _1506_v107;
          _nw267[(new BigNumber(26)).toNumber()] = _1506_v107;
          _1508_v109 = _nw267;
          let _index272 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1508_v109).length));
          (_1508_v109)[_index272] = _1506_v107;
          let _index273 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1508_v109).length));
          (_1508_v109)[_index273] = _1506_v107;
          let _1509_v110;
          let _init48 = function (_1510_i11) {
            return (_this).f25;
          };
          let _nw268 = Array((new BigNumber(25)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw268.length); _i0_48++) {
            _nw268[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1509_v110 = _nw268;
          let _index274 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1509_v110).length));
          (_1509_v110)[_index274] = !((_this).f25) || (true);
          let _index275 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1509_v110).length));
          (_1509_v110)[_index275] = (_this).f25;
          (globalState).f0 = _module.__default.safeModuloInt((_this).f26, (_this).f27);
          let _1511_v111;
          let _nw269 = new _module.C3();
          _nw269.__ctor(_this.f28, (_this).f25, false, (_this).f27);
          _1511_v111 = _nw269;
          let _1512_v112;
          _1512_v112 = _dafny.Map.Empty.slice().updateUnsafe((((_1502_v103).contains((_this).f27)) ? ((_1502_v103).get((_this).f27)) : (new BigNumber(205))),_1511_v111);
          _1503_v104 = (_1503_v104).update((new BigNumber(-314)).minus(new BigNumber((_1512_v112).length)), (_this).f27);
          let _index276 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1509_v110).length));
          (_1509_v110)[_index276] = (_1502_v103).equals(_1502_v103);
        } else {
          let _1513_v113;
          _1513_v113 = _module.D1.create_DC4((_this).f25, (_this).f25);
          let _1514_v114;
          _1514_v114 = _module.D1.create_DC5(_module.D1.create_DC5(_1513_v113));
          let _1515_v115;
          _1515_v115 = _dafny.Seq.of(_1514_v114, _1514_v114);
          let _1516_v116;
          _1516_v116 = _dafny.Set.fromElements(_1515_v115);
          let _1517_v117;
          _1517_v117 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(839)), ((_1518_v113) => function (_1519_i12) {
            return _module.D1.create_DC5(_1518_v113);
          })(_1513_v113)));
          let _1520_v119;
          _1520_v119 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_1516_v116).Difference(function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of (_1517_v117).Elements) {
              let _1521_v118 = _compr_32;
              if (_dafny.Seq.contains(_1517_v117, _1521_v118)) {
                _coll32.add(_1521_v118);
              }
            }
            return _coll32;
          }()));
          _1520_v119 = (_1520_v119).update((_this).f25, _1516_v116);
          let _1522_v120;
          _1522_v120 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f27);
          let _1523_v121;
          let _nw270 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1523_v121 = _nw270;
          r0 = ((!(_1522_v120).equals(_1522_v120)) ? (_1523_v121) : (_1523_v121));
          r1 = _dafny.Seq.contains(_module.__default.fm1(globalState), _1421_v39);
          (globalState).f0 = ((((_1503_v104).contains((_this).f26)) ? ((_1503_v104).get((_this).f26)) : ((_this).f26))).minus((_this).f27);
          r1 = (_this).f25;
        }
        let _1524_v122;
        _1524_v122 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f26),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(295)), function (_1525_i13) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25)).length);
        })).length),(_this).f25));
        r2 = _module.__default.fm2(new BigNumber((_1524_v122).length), (_this).f26, globalState);
        let _1526_v123;
        let _out33;
        _out33 = _module.__default.m0((_1505_v106)[_module.__default.safeIndex((_this).f26, new BigNumber((_1505_v106).length))], globalState);
        _1526_v123 = _out33;
        r2 = (_this).f27;
        (globalState).f1 = (_this).f25;
      }
      (globalState).f2 = (_dafny.ZERO).minus(((new BigNumber((_dafny.Seq.of(new BigNumber(782))).length)).plus((_this).f26)).plus((_dafny.ZERO).minus((_this).f27)));
      let _1527_v124;
      _1527_v124 = _module.D0.create_DC0((_this).f25);
      let _1528_v125;
      _1528_v125 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm28(_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), ((_1529_v39) => function (_1530_i14) {
        return _1529_v39;
      })(_1421_v39)), _1421_v39, new BigNumber(-187), globalState),(_this).f26);
      let _1531_v126;
      _1531_v126 = _dafny.MultiSet.fromElements(true, (_this).f25);
      if (_module.__default.fm5((_1527_v124).dtor_cf0, _1528_v125, (((_1531_v126).contains((_this).f25)) ? ((_1531_v126).get((_this).f25)) : ((_this).f27)), globalState)) {
        let _1532_v127;
        _1532_v127 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f27);
        let _1533_v128;
        _1533_v128 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,!(new BigNumber((_dafny.Seq.of((_this).f25)).length)).isEqualTo(new BigNumber((_1532_v127).length)));
        let _1534_v129;
        _1534_v129 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("umwt")).length));
        let _1535_v130;
        _1535_v130 = _module.D11.create_DC26((_dafny.ZERO).minus((_this).f26), (_this).f27, (_this).f25);
        let _1536_v131;
        _1536_v131 = _dafny.Set.fromElements((_this).f27, new BigNumber(649));
        let _1537_v132;
        _1537_v132 = _dafny.MultiSet.fromElements((((_1534_v129).contains((_1535_v130).dtor_cf41)) ? ((_1534_v129).get((_1535_v130).dtor_cf41)) : (_module.__default.fm2((_this).f27, new BigNumber((_1536_v131).length), globalState))), (_this).f26, (_this).f27);
        _1533_v128 = (_1533_v128).update(new BigNumber((_1537_v132).cardinality()), (_this).f25);
        let _1538_v133;
        _1538_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1531_v126,_1421_v39);
        let _1539_v134;
        _1539_v134 = _module.D14.create_DC31((_1538_v133).update(_dafny.MultiSet.fromElements((_this).f25, (_this).f25), _1421_v39));
        let _1540_v135;
        let _nw271 = Array((new BigNumber(22)).toNumber());
        _nw271[(_dafny.ZERO).toNumber()] = (_1538_v133).Merge(_1538_v133);
        _nw271[(_dafny.ONE).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(2)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(3)).toNumber()] = ((_1539_v134).dtor_cf51).update(_1531_v126, _1421_v39);
        _nw271[(new BigNumber(4)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm27((_this).f26, globalState),new _dafny.CodePoint('n'.codePointAt(0)));
        _nw271[(new BigNumber(6)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1531_v126,_1421_v39);
        _nw271[(new BigNumber(8)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(9)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(10)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1531_v126,_1421_v39);
        _nw271[(new BigNumber(12)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(13)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(14)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1531_v126,_1421_v39);
        _nw271[(new BigNumber(16)).toNumber()] = (_1538_v133).Merge(_1538_v133);
        _nw271[(new BigNumber(17)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(18)).toNumber()] = (_1538_v133).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1531_v126,_1421_v39));
        _nw271[(new BigNumber(19)).toNumber()] = _module.__default.fm37((_this).f25, globalState);
        _nw271[(new BigNumber(20)).toNumber()] = _1538_v133;
        _nw271[(new BigNumber(21)).toNumber()] = _1538_v133;
        _1540_v135 = _nw271;
        let _1541_v136;
        _1541_v136 = _dafny.Seq.of(_1536_v131, _dafny.Set.fromElements((_this).f27));
        let _index277 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1540_v135).length));
        (_1540_v135)[_index277] = _module.__default.fm37((_module.D12.create_DC29((_this).f27, (_this).f25, _1541_v136, (_dafny.ZERO).minus((_this).f26), new BigNumber((_1532_v127).length))).dtor_cf46, globalState);
        let _index278 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1540_v135).length));
        (_1540_v135)[_index278] = _1538_v133;
        let _source21 = _module.__default.fm34(_this.f28, new BigNumber(277), globalState);
        if (_source21.is_DC26) {
          let _1542___mcc_h4 = (_source21).cf40;
          let _1543___mcc_h5 = (_source21).cf41;
          let _1544___mcc_h6 = (_source21).cf42;
          let _1545_cf42 = _1544___mcc_h6;
          let _1546_cf41 = _1543___mcc_h5;
          let _1547_cf40 = _1542___mcc_h4;
          let _1548_v137;
          _1548_v137 = _dafny.Seq.of(_1547_cf40, (_this).f27, _1546_cf41, (_this).f26);
          _1548_v137 = _1548_v137;
          let _1549_v138;
          let _init49 = function (_1550_i15) {
            return (_this).f25;
          };
          let _nw272 = Array((new BigNumber(24)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw272.length); _i0_49++) {
            _nw272[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1549_v138 = _nw272;
          let _1551_v139;
          let _nw273 = new _module.C0();
          _nw273.__ctor(_this.f28, _1545_cf42, (_this).f27, _1549_v138, _1545_cf42);
          _1551_v139 = _nw273;
          let _1552_v140;
          let _nw274 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1552_v140 = _nw274;
          _1552_v140 = _1552_v140;
          let _1553_v141;
          _1553_v141 = _dafny.Map.Empty.slice().updateUnsafe(false,_1545_cf42);
          let _1554_v142;
          _1554_v142 = _module.D14.create_DC32((_this).f25, (_this).f26, _1553_v141, _dafny.Seq.of(_1421_v39), (_this).f25);
          let _1555_v143;
          let _nw275 = new _module.C1();
          _nw275.__ctor((_this).f25);
          _1555_v143 = _nw275;
          let _1556_v144;
          _1556_v144 = _module.D10.create_DC24((_1554_v142).dtor_cf53, _1555_v143, _1421_v39);
          (globalState).f2 = (_1556_v144).dtor_cf36;
        } else if (_source21.is_DC25) {
          let _1557___mcc_h7 = (_source21).cf39;
          let _1558_cf39 = _1557___mcc_h7;
          let _1559_v145;
          let _nw276 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1559_v145 = _nw276;
          let _rhs287 = (_this).f27;
          let _rhs288 = _1559_v145;
          let _rhs289 = (_this).f25;
          let _lhs244 = globalState;
          let _lhs245 = globalState;
          _lhs244.f2 = _rhs287;
          _1559_v145 = _rhs288;
          _lhs245.f1 = _rhs289;
          let _1560_v146;
          let _init50 = function (_1561_i16) {
            return !((_this).f25);
          };
          let _nw277 = Array((new BigNumber(10)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw277.length); _i0_50++) {
            _nw277[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1560_v146 = _nw277;
          let _index279 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_1560_v146).length));
          (_1560_v146)[_index279] = !((_this).f25);
          let _index280 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_1560_v146).length));
          (_1560_v146)[_index280] = (_this).f25;
          let _1562_v147;
          let _nw278 = Array((new BigNumber(16)).toNumber()).fill([]);
          _1562_v147 = _nw278;
          let _1563_v148;
          _1563_v148 = _module.D15.create_DC33(_1562_v147);
          _1562_v147 = ((((_1560_v146)[_module.__default.safeIndex(new BigNumber(374), new BigNumber((_1560_v146).length))]) ? (_1563_v148) : (_module.D15.create_DC33(_1562_v147)))).dtor_cf57;
          (globalState).f2 = (_this).f27;
        } else {
          let _1564___mcc_h8 = (_source21).cf43;
          let _1565_cf43 = _1564___mcc_h8;
          (globalState).f1 = (_this).f25;
          let _1566_v149;
          let _nw279 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1566_v149 = _nw279;
          let _index281 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1566_v149).length));
          (_1566_v149)[_index281] = _module.__default.safeModuloInt(new BigNumber((_1531_v126).cardinality()), _module.__default.fm2((_this).f27, (_this).f27, globalState));
          let _1567_v150;
          _1567_v150 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_1568_v39) => function (_1569_i17) {
            return _1568_v39;
          })(_1421_v39))).length),(_this).f27);
          let _1570_v151;
          _1570_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1567_v150,new BigNumber((_1534_v129).cardinality()));
          let _index282 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1566_v149).length));
          (_1566_v149)[_index282] = _module.__default.safeDivisionInt(new BigNumber(((_1570_v151).Merge(_1570_v151)).length), _module.__default.fm2((_this).f27, (_this).f26, globalState));
          r2 = ((_this).f27).multipliedBy((_this).f27);
          (globalState).f2 = _module.__default.safeModuloInt((_this).f27, (_this).f26);
        }
        let _1571_v153;
        _1571_v153 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,function () {
          let _coll33 = new _dafny.Set();
          for (const _compr_33 of _dafny.IntegerRange(new BigNumber(881), new BigNumber(80))) {
            let _1572_v152 = _compr_33;
            if (((new BigNumber(881)).isLessThanOrEqualTo(_1572_v152)) && ((_1572_v152).isLessThan(new BigNumber(80)))) {
              _coll33.add(_module.__default.safeModuloInt(_1572_v152, new BigNumber(328)));
            }
          }
          return _coll33;
        }());
        let _1573_v154;
        _1573_v154 = _module.D2.create_DC6((((_1571_v153).contains((_this).f27)) ? ((_1571_v153).get((_this).f27)) : (_1536_v131)));
        let _1574_v155;
        _1574_v155 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1536_v131).length),(_this).f26);
        let _1575_v156;
        _1575_v156 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
        let _1576_v157;
        _1576_v157 = _module.D14.create_DC32((_this).f25, (_this).f26, _1575_v156, _this.f28, false);
        let _pat_let_tv35 = _1575_v156;
        let _rhs290 = (_this).f25;
        let _rhs291 = function (_pat_let42_0) {
          return function (_1577_dt__update__tmp_h1) {
            return function (_pat_let43_0) {
              return function (_1578_dt__update_hcf10_h0) {
                return _module.D2.create_DC6(_1578_dt__update_hcf10_h0);
              }(_pat_let43_0);
            }((_dafny.Set.fromElements((_this).f26, (_this).f27, new BigNumber(700))).Union(_dafny.Set.fromElements((_this).f27)));
          }(_pat_let42_0);
        }(_module.__default.fm38((((_1533_v128).contains((_this).f27)) ? ((_1533_v128).get((_this).f27)) : ((_this).f25)), (_this).f26, (_dafny.ZERO).minus(new BigNumber((_1574_v155).length)), _module.__default.fm5((_this).f25, _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_this).f25),(_this).f27), (_this).fm6((_this).f25, !((_this).f25), new BigNumber(220), globalState), globalState), globalState));
        let _rhs292 = _module.__default.safeModuloInt((((_1532_v127).contains(true)) ? ((_1532_v127).get(true)) : ((function (_pat_let44_0) {
          return function (_1579_dt__update__tmp_h2) {
            return function (_pat_let45_0) {
              return function (_1580_dt__update_hcf54_h0) {
                return _module.D14.create_DC32((_1579_dt__update__tmp_h2).dtor_cf52, (_1579_dt__update__tmp_h2).dtor_cf53, _1580_dt__update_hcf54_h0, (_1579_dt__update__tmp_h2).dtor_cf55, (_1579_dt__update__tmp_h2).dtor_cf56);
              }(_pat_let45_0);
            }(_pat_let_tv35);
          }(_pat_let44_0);
        }(_1576_v157)).dtor_cf53)), (_this).f26);
        let _rhs293 = (_this).f27;
        let _lhs246 = globalState;
        let _lhs247 = globalState;
        let _lhs248 = globalState;
        _lhs246.f1 = _rhs290;
        _1573_v154 = _rhs291;
        _lhs247.f2 = _rhs292;
        _lhs248.f2 = _rhs293;
        let _1581_v158;
        let _init51 = function (_1582_i18) {
          return true;
        };
        let _nw280 = Array((new BigNumber(17)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw280.length); _i0_51++) {
          _nw280[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1581_v158 = _nw280;
        let _1583_v159;
        let _nw281 = new _module.C0();
        _nw281.__ctor(_this.f28, ((_this).f25) && ((_this).f25), (_this).f27, _1581_v158, (_this).f25);
        _1583_v159 = _nw281;
      } else {
        let _1584_v160;
        _1584_v160 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_dafny.Seq.of((_this).f26)).length));
        let _1585_v161;
        _1585_v161 = _dafny.MultiSet.fromElements(_this.f28);
        let _1586_v162;
        let _init52 = function (_1587_i19) {
          return (_this).f25;
        };
        let _nw282 = Array((new BigNumber(14)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw282.length); _i0_52++) {
          _nw282[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1586_v162 = _nw282;
        let _1588_v163;
        let _nw283 = new _module.C0();
        _nw283.__ctor(_this.f28, (_this).f25, (_this).f27, _1586_v162, (_this).f25);
        _1588_v163 = _nw283;
        let _1589_v164;
        _1589_v164 = _dafny.Map.Empty.slice().updateUnsafe(_1588_v163,(_this).f25);
        let _1590_v165;
        _1590_v165 = _module.D0.create_DC1(new BigNumber(39), _1588_v163.f35, false, new BigNumber((_dafny.MultiSet.fromElements((_this).f27, new BigNumber(434), (_this).f26)).cardinality()), (_this).f25);
        let _1591_v166;
        let _nw284 = new _module.C3();
        _nw284.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_1592_v39) => function (_1593_i20) {
          return _1592_v39;
        })(_1421_v39)), (_this).f25, (_this).f25, (_this).f26);
        _1591_v166 = _nw284;
        let _1594_v169;
        _1594_v169 = _module.D16.create_DC35(_1584_v160);
        let _1595_v172;
        _1595_v172 = _dafny.MultiSet.fromElements((_this).f26);
        let _1596_v174;
        _1596_v174 = _dafny.Seq.of(_1584_v160, _1584_v160);
        let _1597_v175;
        _1597_v175 = _dafny.Map.Empty.slice().updateUnsafe(true,_1421_v39);
        let _1598_v176;
        _1598_v176 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1597_v175).length),_this.f28);
        let _1599_v177;
        let _nw285 = Array((new BigNumber(29)).toNumber());
        _nw285[(_dafny.ZERO).toNumber()] = _1584_v160;
        _nw285[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(((_1585_v161).contains(_dafny.Seq.update(_this.f28, _module.__default.safeIndex(new BigNumber((_1589_v164).length), new BigNumber((_this.f28).length)), _module.__default.fm4(_1590_v165, (_this).f25, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1591_v166,(_1591_v166).f30)).length), globalState)))) ? ((_1585_v161).get(_dafny.Seq.update(_this.f28, _module.__default.safeIndex(new BigNumber((_1589_v164).length), new BigNumber((_this.f28).length)), _module.__default.fm4(_1590_v165, (_this).f25, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1591_v166,(_1591_v166).f30)).length), globalState)))) : ((_this).f26)));
        _nw285[(new BigNumber(2)).toNumber()] = (function () {
          let _coll34 = new _dafny.Map();
          for (const _compr_34 of _dafny.IntegerRange(new BigNumber(83), new BigNumber(340))) {
            let _1600_v167 = _compr_34;
            if (((new BigNumber(83)).isLessThanOrEqualTo(_1600_v167)) && ((_1600_v167).isLessThan(new BigNumber(340)))) {
              _coll34.push([(_1600_v167).minus((_this).f27),(_this).f27]);
            }
          }
          return _coll34;
        }()).Merge(_1584_v160);
        _nw285[(new BigNumber(3)).toNumber()] = (_1584_v160).Merge(_1584_v160);
        _nw285[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27)).update((_this).f27, (_this).f26);
        _nw285[(new BigNumber(5)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f26),(_this).fm6((_1591_v166).f30, !((_1591_v166).f30), (_this).f26, globalState))).Merge(_1584_v160);
        _nw285[(new BigNumber(7)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
        _nw285[(new BigNumber(9)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(10)).toNumber()] = (_1584_v160).Merge(_1584_v160);
        _nw285[(new BigNumber(11)).toNumber()] = function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of _dafny.IntegerRange(new BigNumber(491), new BigNumber(-882))) {
            let _1601_v168 = _compr_35;
            if (((new BigNumber(491)).isLessThanOrEqualTo(_1601_v168)) && ((_1601_v168).isLessThan(new BigNumber(-882)))) {
              _coll35.push([(_1601_v168).plus(new BigNumber((_dafny.MultiSet.FromArray((_1591_v166).f29)).cardinality())),new BigNumber(343)]);
            }
          }
          return _coll35;
        }();
        _nw285[(new BigNumber(12)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(13)).toNumber()] = (_1594_v169).dtor_cf59;
        _nw285[(new BigNumber(14)).toNumber()] = (function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of _dafny.IntegerRange(new BigNumber(580), new BigNumber(30))) {
            let _1602_v170 = _compr_36;
            if (((new BigNumber(580)).isLessThanOrEqualTo(_1602_v170)) && ((_1602_v170).isLessThan(new BigNumber(30)))) {
              _coll36.push([(_1602_v170).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_1591_v166).f30)).length)),(_this).f26]);
            }
          }
          return _coll36;
        }()).Merge(_1584_v160);
        _nw285[(new BigNumber(15)).toNumber()] = (function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_1595_v172).Elements) {
            let _1603_v171 = _compr_37;
            if ((_1595_v172).contains(_1603_v171)) {
              _coll37.push([_module.__default.safeModuloInt(_1603_v171, (_this).f27),(_this).f26]);
            }
          }
          return _coll37;
        }()).Merge(_1584_v160);
        _nw285[(new BigNumber(16)).toNumber()] = function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of _dafny.IntegerRange(new BigNumber(77), new BigNumber(-552))) {
            let _1604_v173 = _compr_38;
            if (((new BigNumber(77)).isLessThanOrEqualTo(_1604_v173)) && ((_1604_v173).isLessThan(new BigNumber(-552)))) {
              _coll38.push([(_1604_v173).plus(new BigNumber(-918)),(_this).f26]);
            }
          }
          return _coll38;
        }();
        _nw285[(new BigNumber(17)).toNumber()] = (_1596_v174)[_module.__default.safeIndex((_this).f27, new BigNumber((_1596_v174).length))];
        _nw285[(new BigNumber(18)).toNumber()] = (((_this).f25) ? (_1584_v160) : (_1584_v160));
        _nw285[(new BigNumber(19)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f27);
        _nw285[(new BigNumber(21)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(22)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(23)).toNumber()] = (_1584_v160).Merge(_1584_v160);
        _nw285[(new BigNumber(24)).toNumber()] = (_1584_v160).Merge(_1584_v160);
        _nw285[(new BigNumber(25)).toNumber()] = _1584_v160;
        _nw285[(new BigNumber(26)).toNumber()] = (_1584_v160).update(new BigNumber(159), new BigNumber((_1598_v176).length));
        _nw285[(new BigNumber(27)).toNumber()] = (_1584_v160).Merge(_1584_v160);
        _nw285[(new BigNumber(28)).toNumber()] = _1584_v160;
        _1599_v177 = _nw285;
        let _nw286 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
        _1599_v177 = _nw286;
        let _1605_v178;
        let _init53 = function (_1606_i21) {
          return _module.__default.safeModuloInt(_1606_i21, (_this).f26);
        };
        let _nw287 = Array((new BigNumber(13)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw287.length); _i0_53++) {
          _nw287[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1605_v178 = _nw287;
        _1605_v178 = (((_1591_v166).f30) ? (_1605_v178) : (_1605_v178));
        r1 = (_1591_v166).f30;
        _1421_v39 = _1421_v39;
        let _1607_v179;
        let _nw288 = Array((new BigNumber(8)).toNumber()).fill(_module.D0.Default());
        _1607_v179 = _nw288;
        let _1608_v180;
        _1608_v180 = _dafny.Set.fromElements(_1607_v179);
        (globalState).f1 = (_1608_v180).IsProperSubsetOf(_dafny.Set.fromElements(_1607_v179, _1607_v179));
      }
      let _1609_v181;
      let _nw289 = Array((new BigNumber(29)).toNumber()).fill(false);
      _1609_v181 = _nw289;
      r0 = _1609_v181;
      let _1610_v182;
      _1610_v182 = _dafny.Seq.of((_this).f25);
      let _1611_v183;
      _1611_v183 = _dafny.MultiSet.fromElements((_this).f27);
      let _1612_v184;
      _1612_v184 = _dafny.MultiSet.fromElements(_1610_v182, _dafny.Seq.of((_this).f25));
      r1 = !((_1612_v184).contains(_dafny.Seq.update(_1610_v182, _module.__default.safeIndex(_module.__default.fm2((_this).f27, (((_1611_v183).contains(new BigNumber((_1610_v182).length))) ? ((_1611_v183).get(new BigNumber((_1610_v182).length))) : ((_this).f26)), globalState), new BigNumber((_1610_v182).length)), (_this).f25)));
      r2 = _module.__default.safeDivisionInt((((_this).f25) ? ((_this).f26) : ((_this).f26)), (_this).f27);
      return [r0, r1, r2];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
