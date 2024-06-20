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
    static fm0(globalState) {
      return (_module.D7.create_DC16(new BigNumber(163), new BigNumber(125), true, false, new BigNumber(276))).dtor_cf26;
    };
    static fm1(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(387))).Intersect(_dafny.Set.fromElements(new BigNumber(-805)))).Union((_dafny.Set.fromElements(new BigNumber(624), (_module.D6.create_DC13(new BigNumber(510))).dtor_cf19)));
    };
    static fm2(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('f'.codePointAt(0));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)), _dafny.Seq.of(true, false));
    };
    static fm4(p0, p1, globalState) {
      return new BigNumber(549);
    };
    static fm5(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, true, !(true), !(false), false)).length),(_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(192)))));
    };
    static fm6(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true))));
    };
    static fm7(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(315),false)).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(525), new BigNumber(56))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(525)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(56)))) {
            _coll0.push([(_0_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())),!(!(false))]);
          }
        }
        return _coll0;
      }());
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("drfsgt"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.UnicodeFromString("wtfucath")));
    };
    static fm9(globalState) {
      if (false) {
        return _module.D3.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(864)), function (_1_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}));
      } else {
        return _module.D3.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-403)), function (_2_i1) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}));
      }
    };
    static fm10(p0, p1, p2, globalState) {
      let _source0 = ((false) ? (_module.D7.create_DC15((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(983),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(true))).length))) : (_module.D7.create_DC15(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(!(false)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ujipcla")).length))));
      if (_source0.is_DC15) {
        let _3___mcc_h0 = (_source0).cf21;
        let _4___mcc_h1 = (_source0).cf22;
        let _5_cf22 = _4___mcc_h1;
        let _6_cf21 = _3___mcc_h0;
        return function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("htgdju")).length),_6_cf21), _dafny.Map.Empty.slice().updateUnsafe(_6_cf21,_5_cf22), _dafny.Map.Empty.slice().updateUnsafe(_6_cf21,_6_cf21)))).Elements) {
            let _7_v0 = _compr_1;
            if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("htgdju")).length),_6_cf21), _dafny.Map.Empty.slice().updateUnsafe(_6_cf21,_5_cf22), _dafny.Map.Empty.slice().updateUnsafe(_6_cf21,_6_cf21)))).contains(_7_v0)) {
              _coll1.push([_7_v0,true]);
            }
          }
          return _coll1;
        }();
      } else if (_source0.is_DC16) {
        let _8___mcc_h2 = (_source0).cf23;
        let _9___mcc_h3 = (_source0).cf24;
        let _10___mcc_h4 = (_source0).cf25;
        let _11___mcc_h5 = (_source0).cf26;
        let _12___mcc_h6 = (_source0).cf27;
        let _13_cf27 = _12___mcc_h6;
        let _14_cf26 = _11___mcc_h5;
        let _15_cf25 = _10___mcc_h4;
        let _16_cf24 = _9___mcc_h3;
        let _17_cf23 = _8___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_17_cf23,_16_cf24),_15_cf25);
      } else if (_source0.is_DC14) {
        let _18___mcc_h7 = (_source0).cf20;
        let _19_cf20 = _18___mcc_h7;
        return (_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(51),true)).Keys.Elements) {
            let _20_v1 = _compr_2;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(51),true)).contains(_20_v1)) {
              _coll2.push([(_20_v1).multipliedBy(new BigNumber((_dafny.Set.fromElements(false, false)).length)),new BigNumber(420)]);
            }
          }
          return _coll2;
        }(),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).Elements) {
              let _21_v3 = _compr_4;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).contains(_21_v3)) {
                _coll4.push([_21_v3,new BigNumber(-542)]);
              }
            }
            return _coll4;
          }()).length), new BigNumber((function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_22_i0) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            })).length))).Elements) {
              let _23_v4 = _compr_5;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_22_i0) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              })).length)), _23_v4)) {
                _coll5.push([(_23_v4).plus(new BigNumber(405)),new BigNumber(190)]);
              }
            }
            return _coll5;
          }()).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), function (_24_i1) {
            return new BigNumber(83);
          })).length), new BigNumber(350))).Elements) {
            let _25_v2 = _compr_3;
            if ((_dafny.MultiSet.fromElements(new BigNumber((function () {
              let _coll6 = new _dafny.Map();
              for (const _compr_6 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).Elements) {
                let _21_v3 = _compr_6;
                if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).contains(_21_v3)) {
                  _coll6.push([_21_v3,new BigNumber(-542)]);
                }
              }
              return _coll6;
            }()).length), new BigNumber((function () {
              let _coll7 = new _dafny.Map();
              for (const _compr_7 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_22_i0) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              })).length))).Elements) {
                let _23_v4 = _compr_7;
                if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_22_i0) {
                  return new _dafny.CodePoint('a'.codePointAt(0));
                })).length)), _23_v4)) {
                  _coll7.push([(_23_v4).plus(new BigNumber(405)),new BigNumber(190)]);
                }
              }
              return _coll7;
            }()).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), function (_24_i1) {
              return new BigNumber(83);
            })).length), new BigNumber(350))).contains(_25_v2)) {
              _coll3.push([_module.__default.safeDivisionInt(_25_v2, new BigNumber((_dafny.Set.fromElements(!(true))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true, false, true)).length))).length))]);
            }
          }
          return _coll3;
        }(),!(true)));
      } else {
        let _26___mcc_h8 = (_source0).cf28;
        let _27_cf28 = _26___mcc_h8;
        return function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Set.fromElements(new BigNumber(658), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_28_i2) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length), new BigNumber(-688), new BigNumber(-524), new BigNumber(-555))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fvgbw")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(394))).length)))).cardinality()),new BigNumber(-341)))).Elements) {
            let _29_v5 = _compr_8;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Set.fromElements(new BigNumber(658), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_28_i2) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            })).length), new BigNumber(-688), new BigNumber(-524), new BigNumber(-555))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fvgbw")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(394))).length)))).cardinality()),new BigNumber(-341))), _29_v5)) {
              _coll8.push([_29_v5,false]);
            }
          }
          return _coll8;
        }();
      }
    };
    static fm11(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(492), new BigNumber(841))) {
          let _30_v0 = _compr_9;
          if (((new BigNumber(492)).isLessThanOrEqualTo(_30_v0)) && ((_30_v0).isLessThan(new BigNumber(841)))) {
            _coll9.add((_30_v0).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(256))).length))));
          }
        }
        return _coll9;
      }()).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("orpdwcbhu")).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), function (_31_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()))).Elements) {
          let _32_v1 = _compr_10;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality())), _32_v1)) {
            _coll10.push([(_32_v1).plus(new BigNumber(445)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(908))).length)]);
          }
        }
        return _coll10;
      }()).length))));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _33_i0;
      _33_i0 = _dafny.ZERO;
      L0: {
        while (p3) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_33_i0)) {
              break L0;
            }
            _33_i0 = (_33_i0).plus(_dafny.ONE);
            let _34_v0;
            _34_v0 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
            (globalState).f8 = new BigNumber(((_34_v0).update(p3, p0)).length);
            let _35_v1;
            let _nw0 = new _module.C0();
            _nw0.__ctor();
            _35_v1 = _nw0;
            let _36_v2;
            let _nw1 = new _module.C0();
            _nw1.__ctor();
            _36_v2 = _nw1;
            let _37_v3;
            _37_v3 = _module.D3.create_DC5(p3, p1);
            let _38_v4;
            _38_v4 = _dafny.Seq.UnicodeFromString("bkerwni");
            let _39_v5;
            _39_v5 = _dafny.Set.fromElements(false, p0);
            let _40_v6;
            _40_v6 = _dafny.Map.Empty.slice().updateUnsafe(_36_v2,p3);
            let _41_v7;
            _41_v7 = _dafny.Seq.of(_40_v6);
            let _42_v8;
            let _nw2 = Array((new BigNumber(18)).toNumber()).fill(false);
            _42_v8 = _nw2;
            let _43_v9;
            _43_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p0, p3)).length),_42_v8);
            let _44_v10;
            let _nw3 = Array((new BigNumber(21)).toNumber());
            _nw3[(_dafny.ZERO).toNumber()] = p1;
            _nw3[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
            _nw3[(new BigNumber(2)).toNumber()] = (_37_v3).dtor_cf7;
            _nw3[(new BigNumber(3)).toNumber()] = p1;
            _nw3[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p1);
            _nw3[(new BigNumber(5)).toNumber()] = p1;
            _nw3[(new BigNumber(6)).toNumber()] = p1;
            _nw3[(new BigNumber(7)).toNumber()] = new BigNumber((_38_v4).length);
            _nw3[(new BigNumber(8)).toNumber()] = new BigNumber((_39_v5).length);
            _nw3[(new BigNumber(9)).toNumber()] = p1;
            _nw3[(new BigNumber(10)).toNumber()] = new BigNumber((_38_v4).length);
            _nw3[(new BigNumber(11)).toNumber()] = p1;
            _nw3[(new BigNumber(12)).toNumber()] = new BigNumber(526);
            _nw3[(new BigNumber(13)).toNumber()] = p1;
            _nw3[(new BigNumber(14)).toNumber()] = p1;
            _nw3[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber((_dafny.Seq.update(_41_v7, _module.__default.safeIndex(p1, new BigNumber((_41_v7).length)), _40_v6)).length));
            _nw3[(new BigNumber(16)).toNumber()] = p1;
            _nw3[(new BigNumber(17)).toNumber()] = new BigNumber((_38_v4).length);
            _nw3[(new BigNumber(18)).toNumber()] = p1;
            _nw3[(new BigNumber(19)).toNumber()] = new BigNumber(((_43_v9).Merge(_43_v9)).length);
            _nw3[(new BigNumber(20)).toNumber()] = p1;
            _44_v10 = _nw3;
            let _index0 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_44_v10).length));
            (_44_v10)[_index0] = new BigNumber(((_34_v0).update(p3, true)).length);
            let _index1 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_44_v10).length));
            (_44_v10)[_index1] = p1;
          }
        }
      }
      let _45_v11;
      _45_v11 = _dafny.Seq.of((_dafny.ZERO).minus(p1));
      let _46_v12;
      _46_v12 = _dafny.Seq.UnicodeFromString("owdvbmnf");
      let _47_v13;
      let _nw4 = new _module.C0();
      _nw4.__ctor();
      _47_v13 = _nw4;
      let _48_v14;
      _48_v14 = _dafny.Seq.of(_47_v13, _47_v13);
      let _49_v15;
      _49_v15 = _dafny.Seq.of(_48_v14, _48_v14, _48_v14, _dafny.Seq.of(_47_v13, _47_v13, _47_v13));
      let _50_v16;
      _50_v16 = _dafny.Map.Empty.slice().updateUnsafe((_45_v11)[_module.__default.safeIndex(new BigNumber((_46_v12).length), new BigNumber((_45_v11).length))],_dafny.Seq.contains(_49_v15, _48_v14));
      let _51_i1;
      _51_i1 = _dafny.ZERO;
      L1: {
        while ((((_50_v16).contains(p1)) ? ((_50_v16).get(p1)) : (p3))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_51_i1)) {
              break L1;
            }
            _51_i1 = (_51_i1).plus(_dafny.ONE);
            if (!(!(!(p0)))) {
              let _52_v17;
              _52_v17 = _dafny.Seq.of(p0, p0);
              let _53_v18;
              _53_v18 = _dafny.Map.Empty.slice().updateUnsafe(_52_v17,_47_v13);
              let _54_v19;
              _54_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_53_v18).contains(_52_v17)) ? ((_53_v18).get(_52_v17)) : (_47_v13)),_dafny.Seq.IsProperPrefixOf(_45_v11, _45_v11));
              _54_v19 = _dafny.Map.Empty.slice().updateUnsafe(_47_v13,p3);
              (globalState).f8 = (_module.__default.fm4(p3, p1, globalState)).multipliedBy(p1);
              let _55_v20;
              let _init0 = ((_56_p3, _57_v11, _58_p1) => function (_59_i2) {
                return (new BigNumber((_dafny.MultiSet.fromElements(_56_p3)).cardinality())).isLessThan((_57_v11)[_module.__default.safeIndex(_58_p1, new BigNumber((_57_v11).length))]);
              })(p3, _45_v11, p1);
              let _nw5 = Array((new BigNumber(22)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw5.length); _i0_0++) {
                _nw5[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _55_v20 = _nw5;
              let _60_v21;
              _60_v21 = _dafny.MultiSet.fromElements(p3);
              let _index2 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_55_v20).length));
              (_55_v20)[_index2] = !(_60_v21).equals(_60_v21);
              let _61_v22;
              _61_v22 = _dafny.MultiSet.fromElements(p1);
              let _62_v23;
              _62_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-270),p1);
              let _63_v24;
              _63_v24 = _dafny.Set.fromElements(new BigNumber(357));
              let _64_v25;
              _64_v25 = _dafny.Map.Empty.slice().updateUnsafe(_47_v13,_63_v24);
              let _index3 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_55_v20).length));
              (_55_v20)[_index3] = (new BigNumber((_52_v17).length)).isLessThan((_dafny.ZERO).minus(((_45_v11)[_module.__default.safeIndex(new BigNumber((_61_v22).cardinality()), new BigNumber((_45_v11).length))]).minus((_dafny.ZERO).minus((((_62_v23).contains(p1)) ? ((_62_v23).get(p1)) : (new BigNumber((_64_v25).length)))))));
              (globalState).f8 = new BigNumber(182);
              let _65_v26;
              _65_v26 = _dafny.Set.fromElements(_47_v13, _47_v13);
              let _66_v27;
              _66_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,_65_v26);
              let _index4 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_55_v20).length));
              (_55_v20)[_index4] = (_65_v26).IsDisjointFrom((((_66_v27).contains(p1)) ? ((_66_v27).get(p1)) : (_65_v26)));
            } else {
              let _67_v28;
              _67_v28 = false;
              let _rhs0 = (_45_v11)[_module.__default.safeIndex((new BigNumber(136)).plus(p1), new BigNumber((_45_v11).length))];
              let _rhs1 = !(_67_v28);
              let _rhs2 = _46_v12;
              let _rhs3 = _dafny.Seq.contains(_46_v12, p2);
              let _lhs0 = globalState;
              _lhs0.f8 = _rhs0;
              _67_v28 = _rhs1;
              _46_v12 = _rhs2;
              _67_v28 = _rhs3;
              let _68_v29;
              _68_v29 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),p0);
              let _69_v30;
              _69_v30 = _dafny.Seq.of(_68_v29, _dafny.Map.Empty.slice().updateUnsafe(p3,_67_v28));
              _69_v30 = _dafny.Seq.Concat(_dafny.Seq.Concat(_69_v30, _69_v30), _69_v30);
              let _70_v31;
              _70_v31 = _dafny.Map.Empty.slice().updateUnsafe(_47_v13,p2);
              _70_v31 = (_70_v31).update(_47_v13, p2);
              let _71_v32;
              let _nw6 = new _module.C0();
              _nw6.__ctor();
              _71_v32 = _nw6;
              let _72_v33;
              _72_v33 = _dafny.MultiSet.fromElements(p2, new _dafny.CodePoint('q'.codePointAt(0)), p2, new _dafny.CodePoint('f'.codePointAt(0)), p2);
              let _73_v34;
              _73_v34 = _dafny.Set.fromElements(p1);
              let _74_v35;
              _74_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _75_v36;
              let _nw7 = Array((new BigNumber(19)).toNumber());
              _nw7[(_dafny.ZERO).toNumber()] = p1;
              _nw7[(_dafny.ONE).toNumber()] = new BigNumber(293);
              _nw7[(new BigNumber(2)).toNumber()] = p1;
              _nw7[(new BigNumber(3)).toNumber()] = (((_72_v33).contains(p2)) ? ((_72_v33).get(p2)) : (p1));
              _nw7[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_45_v11, _45_v11)).length);
              _nw7[(new BigNumber(5)).toNumber()] = p1;
              _nw7[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm4(p0, p1, globalState));
              _nw7[(new BigNumber(7)).toNumber()] = p1;
              _nw7[(new BigNumber(8)).toNumber()] = p1;
              _nw7[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_71_v32,new BigNumber((_45_v11).length))).length);
              _nw7[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_module.D3.create_DC5(false, new BigNumber(662))).dtor_cf7);
              _nw7[(new BigNumber(11)).toNumber()] = p1;
              _nw7[(new BigNumber(12)).toNumber()] = new BigNumber(185);
              _nw7[(new BigNumber(13)).toNumber()] = (new BigNumber((_73_v34).length)).plus(p1);
              _nw7[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(p1, p1);
              _nw7[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(862)), function (_76_i3) {
                return new BigNumber(326);
              })).length);
              _nw7[(new BigNumber(16)).toNumber()] = p1;
              _nw7[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-257), p1);
              _nw7[(new BigNumber(18)).toNumber()] = new BigNumber(((_74_v35).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,p1))).length);
              _75_v36 = _nw7;
              let _index5 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_75_v36).length));
              (_75_v36)[_index5] = new BigNumber(506);
              let _77_v37;
              _77_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,_46_v12);
              let _index6 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_75_v36).length));
              (_75_v36)[_index6] = (p1).plus(_module.__default.safeDivisionInt(new BigNumber(15), new BigNumber((_77_v37).length)));
            }
            let _78_v38;
            _78_v38 = false;
            _78_v38 = p0;
            let _79_v39;
            let _nw8 = new _module.C0();
            _nw8.__ctor();
            _79_v39 = _nw8;
            let _80_v40;
            let _nw9 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _80_v40 = _nw9;
            let _81_v41;
            _81_v41 = _dafny.Map.Empty.slice().updateUnsafe(_80_v40,_dafny.Seq.Concat(_46_v12, _dafny.Seq.update(_46_v12, _module.__default.safeIndex(p1, new BigNumber((_46_v12).length)), new _dafny.CodePoint('e'.codePointAt(0)))));
            let _82_v42;
            _82_v42 = _module.D3.create_DC5(true, p1);
            let _rhs4 = (_81_v41).Merge(_81_v41);
            let _rhs5 = p3;
            let _rhs6 = (_82_v42).dtor_cf7;
            let _rhs7 = _module.__default.fm4(_78_v38, (_dafny.ZERO).minus(p1), globalState);
            let _lhs1 = globalState;
            let _lhs2 = globalState;
            _81_v41 = _rhs4;
            _78_v38 = _rhs5;
            _lhs1.f0 = _rhs6;
            _lhs2.f0 = _rhs7;
          }
        }
      }
      let _source1 = _module.__default.fm9(globalState);
      if (_source1.is_DC5) {
        let _83___mcc_h0 = (_source1).cf6;
        let _84___mcc_h1 = (_source1).cf7;
        let _85_cf7 = _84___mcc_h1;
        let _86_cf6 = _83___mcc_h0;
        let _87_v43;
        let _nw10 = new _module.C0();
        _nw10.__ctor();
        _87_v43 = _nw10;
        let _88_v44;
        _88_v44 = new _dafny.CodePoint('g'.codePointAt(0));
        _88_v44 = ((p0) ? (_module.__default.fm2(_dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(p1)), !(!(_86_cf6)), _85_cf7, globalState)) : (_88_v44));
        _86_cf6 = p0;
        _86_cf6 = ((!(_86_cf6)) ? (p3) : (p0));
      } else {
        let _89___mcc_h2 = (_source1).cf5;
        let _90_cf5 = _89___mcc_h2;
        let _91_v45;
        _91_v45 = _dafny.Map.Empty.slice().updateUnsafe(_45_v11,_90_cf5);
        let _92_v46;
        _92_v46 = _module.D6.create_DC10(_dafny.Seq.of(p1));
        _91_v45 = (_91_v45).update(_dafny.Seq.update(_dafny.Seq.update(_45_v11, _module.__default.safeIndex(new BigNumber(((_92_v46).dtor_cf14).length), new BigNumber((_45_v11).length)), (_45_v11)[_module.__default.safeIndex(p1, new BigNumber((_45_v11).length))]), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.update(_45_v11, _module.__default.safeIndex(new BigNumber(((_92_v46).dtor_cf14).length), new BigNumber((_45_v11).length)), (_45_v11)[_module.__default.safeIndex(p1, new BigNumber((_45_v11).length))])).length)), p1), _dafny.Seq.UnicodeFromString("imvdwrgjx"));
        let _93_v47;
        let _nw11 = new _module.C0();
        _nw11.__ctor();
        _93_v47 = _nw11;
        (globalState).f0 = p1;
        let _94_v48;
        let _nw12 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _94_v48 = _nw12;
        let _index7 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_94_v48).length));
        (_94_v48)[_index7] = (new BigNumber((_90_cf5).length)).plus(new BigNumber((function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of _dafny.IntegerRange(new BigNumber(400), new BigNumber(602))) {
            let _95_v49 = _compr_11;
            if (((new BigNumber(400)).isLessThanOrEqualTo(_95_v49)) && ((_95_v49).isLessThan(new BigNumber(602)))) {
              _coll11.push([_module.__default.safeDivisionInt(_95_v49, p1),p2]);
            }
          }
          return _coll11;
        }()).length));
        let _index8 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_94_v48).length));
        (_94_v48)[_index8] = p1;
      }
      let _96_v50;
      _96_v50 = _dafny.MultiSet.fromElements(p1);
      let _97_v51;
      _97_v51 = _dafny.Seq.of(((!(p3)) ? (!(p3)) : (p0)), (_96_v50).IsProperSubsetOf(_96_v50));
      _97_v51 = _97_v51;
      let _98_v52;
      let _nw13 = new _module.C0();
      _nw13.__ctor();
      _98_v52 = _nw13;
      let _hi0 = (_dafny.ZERO).minus((p1).plus(p1));
      for (let _99_i4 = p1; _99_i4.isLessThan(_hi0); _99_i4 = _99_i4.plus(_dafny.ONE)) {
        let _100_v53;
        let _init1 = ((_101_p1) => function (_102_i5) {
          return (_102_i5).multipliedBy(_101_p1);
        })(p1);
        let _nw14 = Array((new BigNumber(3)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
          _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _100_v53 = _nw14;
        let _index9 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length));
        (_100_v53)[_index9] = (new BigNumber(751)).minus(p1);
        let _103_v54;
        let _nw15 = Array((new BigNumber(6)).toNumber()).fill(false);
        _103_v54 = _nw15;
        let _104_v55;
        _104_v55 = _dafny.Map.Empty.slice().updateUnsafe(_103_v54,(_97_v51)[_module.__default.safeIndex(p1, new BigNumber((_97_v51).length))]);
        let _index10 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length));
        let _rhs8 = (_dafny.ZERO).minus((((p0) ? (p1) : (_99_i4))).plus(_module.__default.safeModuloInt(_99_i4, _99_i4)));
        let _rhs9 = ((p3) ? (_46_v12) : (_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("br"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("br")).length)), p2), _46_v12)));
        let _rhs10 = _module.__default.safeModuloInt(_99_i4, p1);
        let _rhs11 = _104_v55;
        let _rhs12 = _module.__default.safeDivisionInt(_99_i4, (p1).minus(p1));
        let _lhs3 = globalState;
        let _lhs4 = globalState;
        let _lhs5 = _100_v53;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length));
        let _lhs7 = globalState;
        _lhs3.f0 = _rhs8;
        _lhs4.f2 = _rhs9;
        _lhs5[_lhs6] = _rhs10;
        _104_v55 = _rhs11;
        _lhs7.f8 = _rhs12;
        if (false) {
          let _index11 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_103_v54).length));
          (_103_v54)[_index11] = p3;
          let _index12 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_103_v54).length));
          (_103_v54)[_index12] = ((!(!(p3))) ? (!(!(p1).isEqualTo(_99_i4))) : ((false) && (p0)));
          let _105_v56;
          _105_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_103_v54)[_module.__default.safeIndex(new BigNumber(643), new BigNumber((_103_v54).length))]);
          let _106_v57;
          _106_v57 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_100_v53)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length))]),_105_v56);
          _106_v57 = (_106_v57).update((_dafny.ZERO).minus(_99_i4), _105_v56);
          let _index13 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_103_v54).length));
          (_103_v54)[_index13] = p3;
          let _107_v58;
          _107_v58 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(p0, _99_i4, new BigNumber(-914), globalState),p1);
          let _108_v60;
          _108_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(736),p1);
          let _109_v61;
          _109_v61 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("dxrmfw"), _46_v12);
          let _110_v62;
          _110_v62 = _module.D1.create_DC2(_109_v61, p2);
          let _111_v63;
          _111_v63 = _dafny.Map.Empty.slice().updateUnsafe(_108_v60,_110_v62);
          _107_v58 = (_107_v58).update(function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_111_v63).Keys.Elements) {
              let _112_v59 = _compr_12;
              if ((_111_v63).contains(_112_v59)) {
                _coll12.push([_112_v59,p0]);
              }
            }
            return _coll12;
          }(), p1);
          let _index14 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_103_v54).length));
          (_103_v54)[_index14] = p3;
        } else {
          let _113_v64;
          _113_v64 = true;
          _113_v64 = !((_100_v53)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length))]).isEqualTo(_99_i4);
          _113_v64 = _dafny.Seq.IsProperPrefixOf(_45_v11, _dafny.Seq.Concat(_45_v11, _dafny.Seq.update(_45_v11, _module.__default.safeIndex((_100_v53)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length))], new BigNumber((_45_v11).length)), p1)));
          let _114_v65;
          _114_v65 = _dafny.Seq.of((_dafny.MultiSet.FromArray(_48_v14)).Intersect(_dafny.MultiSet.fromElements(_98_v52)), _dafny.MultiSet.FromArray(((p0) ? (_48_v14) : (_48_v14))), _dafny.MultiSet.FromArray(_dafny.Seq.update(((p0) ? (_dafny.Seq.of(_98_v52)) : (_48_v14)), _module.__default.safeIndex(p1, new BigNumber((((p0) ? (_dafny.Seq.of(_98_v52)) : (_48_v14))).length)), _47_v13)));
          let _rhs13 = _47_v13;
          let _rhs14 = _47_v13;
          let _rhs15 = _114_v65;
          _47_v13 = _rhs13;
          _98_v52 = _rhs14;
          _114_v65 = _rhs15;
          let _115_v66;
          _115_v66 = _module.D4.create_DC8();
          _115_v66 = _115_v66;
          let _116_v67;
          let _nw16 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _116_v67 = _nw16;
          let _117_v68;
          _117_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.MultiSet.fromElements((_97_v51)[_module.__default.safeIndex(p1, new BigNumber((_97_v51).length))])).cardinality()));
          let _118_v69;
          _118_v69 = _dafny.Seq.of(_module.__default.fm11(new BigNumber((_45_v11).length), (_100_v53)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length))], _117_v68, globalState));
          let _index15 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_116_v67).length));
          (_116_v67)[_index15] = _118_v69;
          let _119_v70;
          _119_v70 = _dafny.Set.fromElements(p1);
          let _index16 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_116_v67).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length));
          let _rhs16 = ((_dafny.Seq.IsPrefixOf(_46_v12, _46_v12)) ? (_113_v64) : (p0));
          let _rhs17 = _118_v69;
          let _rhs18 = (_100_v53)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length))];
          let _rhs19 = _module.__default.safeDivisionInt((_100_v53)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length))], new BigNumber((_119_v70).length));
          let _lhs8 = _116_v67;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_116_v67).length));
          let _lhs10 = _100_v53;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_100_v53).length));
          let _lhs12 = globalState;
          _113_v64 = _rhs16;
          _lhs8[_lhs9] = _rhs17;
          _lhs10[_lhs11] = _rhs18;
          _lhs12.f0 = _rhs19;
        }
        let _index18 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_103_v54).length));
        (_103_v54)[_index18] = (_97_v51)[_module.__default.safeIndex(_99_i4, new BigNumber((_97_v51).length))];
        let _120_v71;
        _120_v71 = _dafny.MultiSet.fromElements(_100_v53, _100_v53, _100_v53);
        let _index19 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_103_v54).length));
        (_103_v54)[_index19] = (_120_v71).IsDisjointFrom(_120_v71);
        let _121_v72;
        _121_v72 = _dafny.Map.Empty.slice().updateUnsafe((p0) === (p3),(_96_v50).IsProperSubsetOf(_96_v50));
        let _122_v73;
        _122_v73 = _dafny.Map.Empty.slice().updateUnsafe(_47_v13,(_103_v54)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_103_v54).length))]);
        let _123_v74;
        _123_v74 = _module.D7.create_DC14(_122_v73);
        let _124_v75;
        _124_v75 = _dafny.Map.Empty.slice().updateUnsafe(_96_v50,new BigNumber((((_123_v74).dtor_cf20).update(_47_v13, (_103_v54)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_103_v54).length))])).length));
        _121_v72 = (_121_v72).update((new BigNumber((_124_v75).length)).isLessThan((_dafny.ZERO).minus(p1)), (_103_v54)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_103_v54).length))]);
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _125_v0;
      _125_v0 = _dafny.Seq.UnicodeFromString("si");
      let _126_v1;
      let _nw17 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _126_v1 = _nw17;
      let _127_v2;
      _127_v2 = false;
      let _128_v3;
      let _nw18 = Array((new BigNumber(3)).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = !(_127_v2);
      _nw18[(_dafny.ONE).toNumber()] = !(_127_v2);
      _nw18[(new BigNumber(2)).toNumber()] = _127_v2;
      _128_v3 = _nw18;
      let _129_v4;
      _129_v4 = new _dafny.CodePoint('j'.codePointAt(0));
      let _130_v5;
      _130_v5 = _dafny.Set.fromElements(_129_v4);
      let _131_globalState;
      let _nw19 = new _module.GlobalState();
      _nw19.__ctor(new BigNumber(-164), new BigNumber(378), _125_v0, _126_v1, true, true, _128_v3, new BigNumber(223), new BigNumber(583), _130_v5, new BigNumber(782));
      _131_globalState = _nw19;
      _127_v2 = !((!(_127_v2)) && (_127_v2));
      let _132_v6;
      _132_v6 = new BigNumber(520);
      let _133_v7;
      _133_v7 = _dafny.Map.Empty.slice().updateUnsafe(((_127_v2) ? (_126_v1) : (_126_v1)),((_module.__default.fm0(_131_globalState)) ? (_132_v6) : (_132_v6)));
      let _134_v8;
      _134_v8 = _dafny.Seq.of(_127_v2);
      let _135_v9;
      _135_v9 = _dafny.Set.fromElements(_132_v6, _132_v6, new BigNumber((_134_v8).length));
      _133_v7 = (_133_v7).update(_126_v1, (new BigNumber((_135_v9).length)).minus(_132_v6));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_128_v3).length))) {
        let _136_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_136_i0)) && ((_136_i0).isLessThan(new BigNumber((_128_v3).length))))) {
          (_128_v3)[(_136_i0)] = (_dafny.MultiSet.fromElements(false, _127_v2, false, _127_v2)).IsSubsetOf((_dafny.MultiSet.fromElements(false)));
        }
      }
      let _137_v10;
      _137_v10 = _dafny.MultiSet.fromElements(!(_127_v2), _127_v2);
      let _138_v11;
      _138_v11 = _137_v10;
      let _source2 = _138_v11;
      let _139___mcc_h0 = _source2;
      let _140_cf0 = _139___mcc_h0;
      if (!(!(_127_v2) || (!((_135_v9).IsProperSubsetOf(_135_v9))))) {
        _module.__default.m0(((_127_v2) ? (_module.__default.fm0(_131_globalState)) : ((_134_v8)[_module.__default.safeIndex(_132_v6, new BigNumber((_134_v8).length))])), new BigNumber((_dafny.MultiSet.FromArray(_134_v8)).cardinality()), _129_v4, _127_v2, _131_globalState);
        let _141_v12;
        _141_v12 = _dafny.Set.fromElements(false);
        let _142_v13;
        _142_v13 = _dafny.Seq.of(new BigNumber((_134_v8).length));
        let _143_v14;
        _143_v14 = _dafny.Map.Empty.slice().updateUnsafe((_135_v9).Difference(_dafny.Set.fromElements(new BigNumber((_141_v12).length), new BigNumber(-308))),((_142_v13)[_module.__default.safeIndex(_132_v6, new BigNumber((_142_v13).length))]).plus(_132_v6));
        let _rhs20 = _module.__default.fm0(_131_globalState);
        let _rhs21 = _132_v6;
        let _rhs22 = _132_v6;
        let _rhs23 = (_143_v14).update(_module.__default.fm1(new _dafny.CodePoint('d'.codePointAt(0)), _127_v2, _127_v2, _131_globalState), new BigNumber(80));
        let _lhs13 = _131_globalState;
        let _lhs14 = _131_globalState;
        _127_v2 = _rhs20;
        _lhs13.f8 = _rhs21;
        _lhs14.f8 = _rhs22;
        _143_v14 = _rhs23;
        _127_v2 = _127_v2;
        (_131_globalState).f0 = _132_v6;
        let _144_v15;
        _144_v15 = _dafny.MultiSet.fromElements(_134_v8, _dafny.Seq.of(_127_v2, _127_v2));
        let _145_v16;
        _145_v16 = _dafny.Map.Empty.slice().updateUnsafe(_144_v15,_125_v0);
        let _146_v17;
        _146_v17 = _dafny.Seq.of(_144_v15, _144_v15, _144_v15);
        _145_v16 = (_145_v16).update((_146_v17)[_module.__default.safeIndex(_132_v6, new BigNumber((_146_v17).length))], _125_v0);
      } else {
        let _index20 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_126_v1).length));
        (_126_v1)[_index20] = _129_v4;
        let _147_v18;
        _147_v18 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_132_v6);
        let _index21 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_126_v1).length));
        (_126_v1)[_index21] = _module.__default.fm2(_147_v18, _127_v2, _132_v6, _131_globalState);
        let _148_v19;
        _148_v19 = _dafny.Set.fromElements(_127_v2);
        let _149_v20;
        _149_v20 = _dafny.Map.Empty.slice().updateUnsafe(_129_v4,_dafny.Seq.Create(_module.__default.abs(new BigNumber(430)), ((_150_v4) => function (_151_i1) {
          return _150_v4;
        })(_129_v4)));
        let _152_v21;
        _152_v21 = _dafny.Map.Empty.slice().updateUnsafe(_148_v19,!(_149_v20).equals(_149_v20));
        _152_v21 = (_152_v21).update(_148_v19, _127_v2);
        (_131_globalState).f0 = _132_v6;
        (_131_globalState).f2 = _125_v0;
        _module.__default.m0((_134_v8)[_module.__default.safeIndex(_132_v6, new BigNumber((_134_v8).length))], _132_v6, _129_v4, _127_v2, _131_globalState);
      }
      (_131_globalState).f0 = _132_v6;
      let _153_v22;
      _153_v22 = _module.D1.create_DC1(true);
      if ((_153_v22).dtor_cf1) {
        let _154_v23;
        let _nw20 = Array((new BigNumber(18)).toNumber()).fill([]);
        _154_v23 = _nw20;
        let _155_v24;
        _155_v24 = _dafny.Map.Empty.slice().updateUnsafe(_128_v3,_127_v2);
        let _156_v25;
        _156_v25 = _dafny.Map.Empty.slice().updateUnsafe(_154_v23,(_132_v6).multipliedBy(new BigNumber((_155_v24).length)));
        _156_v25 = (_156_v25).update(_154_v23, _132_v6);
        _134_v8 = _dafny.Seq.Concat(_module.__default.fm3(_131_globalState), _134_v8);
        (_131_globalState).f0 = (_132_v6).minus(_132_v6);
        let _157_v26;
        let _nw21 = new _module.C0();
        _nw21.__ctor();
        _157_v26 = _nw21;
        let _158_v27;
        let _nw22 = Array((new BigNumber(11)).toNumber()).fill([]);
        _158_v27 = _nw22;
        let _159_v28;
        _159_v28 = _dafny.Map.Empty.slice().updateUnsafe(_127_v2,_125_v0);
        let _160_v29;
        _160_v29 = _dafny.Map.Empty.slice().updateUnsafe(_157_v26,(((_159_v28).contains(!(_127_v2))) ? ((_159_v28).get(!(_127_v2))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), ((_161_v4) => function (_162_i2) {
          return _161_v4;
        })(_129_v4)))));
        let _163_v30;
        let _nw23 = Array((new BigNumber(8)).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_129_v4, new _dafny.CodePoint('g'.codePointAt(0)), _129_v4, _129_v4);
        _nw23[(_dafny.ONE).toNumber()] = (((_160_v29).contains(_157_v26)) ? ((_160_v29).get(_157_v26)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-93)), ((_164_v4) => function (_165_i3) {
          return _164_v4;
        })(_129_v4))));
        _nw23[(new BigNumber(2)).toNumber()] = _125_v0;
        _nw23[(new BigNumber(3)).toNumber()] = _125_v0;
        _nw23[(new BigNumber(4)).toNumber()] = _125_v0;
        _nw23[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), function (_166_i4) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        });
        _nw23[(new BigNumber(6)).toNumber()] = _125_v0;
        _nw23[(new BigNumber(7)).toNumber()] = _125_v0;
        _163_v30 = _nw23;
        let _index22 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_158_v27).length));
        (_158_v27)[_index22] = _163_v30;
        let _index23 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_126_v1).length));
        (_126_v1)[_index23] = _129_v4;
        let _167_v31;
        let _nw24 = Array((new BigNumber(27)).toNumber());
        _167_v31 = _nw24;
        let _168_v32;
        _168_v32 = _dafny.Map.Empty.slice().updateUnsafe(_167_v31,_163_v30);
        let _index24 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_158_v27).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_126_v1).length));
        let _rhs24 = (((_168_v32).contains(_167_v31)) ? ((_168_v32).get(_167_v31)) : (_163_v30));
        let _rhs25 = _154_v23;
        let _rhs26 = _129_v4;
        let _rhs27 = _132_v6;
        let _lhs15 = _158_v27;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_158_v27).length));
        let _lhs17 = _126_v1;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_126_v1).length));
        let _lhs19 = _131_globalState;
        _lhs15[_lhs16] = _rhs24;
        _154_v23 = _rhs25;
        _lhs17[_lhs18] = _rhs26;
        _lhs19.f0 = _rhs27;
      } else {
        let _169_v33;
        let _nw25 = Array((new BigNumber(27)).toNumber()).fill([]);
        _169_v33 = _nw25;
        let _index26 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_169_v33).length));
        (_169_v33)[_index26] = _126_v1;
        let _index27 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_169_v33).length));
        let _rhs28 = _126_v1;
        let _rhs29 = (_dafny.ZERO).minus((_132_v6).plus(_132_v6));
        let _lhs20 = _169_v33;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_169_v33).length));
        let _lhs22 = _131_globalState;
        _lhs20[_lhs21] = _rhs28;
        _lhs22.f8 = _rhs29;
        let _170_v34;
        _170_v34 = _dafny.Map.Empty.slice().updateUnsafe(_127_v2,_132_v6);
        _170_v34 = (_170_v34).update(_127_v2, _132_v6);
        let _171_v35;
        _171_v35 = _dafny.Map.Empty.slice().updateUnsafe(_127_v2,_127_v2);
        let _172_v36;
        _172_v36 = _dafny.Seq.of(_135_v9);
        let _173_v37;
        _173_v37 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,new BigNumber((_125_v0).length));
        let _174_v38;
        _174_v38 = _dafny.MultiSet.fromElements(_129_v4);
        let _rhs30 = (_172_v36)[_module.__default.safeIndex((_132_v6).multipliedBy(_132_v6), new BigNumber((_172_v36).length))];
        let _rhs31 = _module.__default.fm4(_127_v2, _132_v6, _131_globalState);
        let _rhs32 = _module.__default.safeDivisionInt(_132_v6, _132_v6);
        let _rhs33 = !((new BigNumber((_173_v37).length)).isEqualTo(new BigNumber((_174_v38).cardinality()))) || ((_127_v2) && (_127_v2));
        let _rhs34 = ((_171_v35).Merge((_171_v35).update(_127_v2, true))).update(_module.__default.fm0(_131_globalState), _127_v2);
        let _lhs23 = _131_globalState;
        _135_v9 = _rhs30;
        _132_v6 = _rhs31;
        _lhs23.f0 = _rhs32;
        _127_v2 = _rhs33;
        _171_v35 = _rhs34;
        let _175_v39;
        _175_v39 = _dafny.Seq.of(_173_v37);
        let _176_v40;
        _176_v40 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_127_v2);
        let _177_v41;
        _177_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((((_176_v40).contains(_132_v6)) ? ((_176_v40).get(_132_v6)) : (_127_v2)))).length),_127_v2);
        let _178_v43;
        let _nw26 = Array((new BigNumber(10)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _173_v37;
        _nw26[(_dafny.ONE).toNumber()] = (_175_v39)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_175_v39).length))];
        _nw26[(new BigNumber(2)).toNumber()] = _module.__default.fm5(new BigNumber(183), _127_v2, _131_globalState);
        _nw26[(new BigNumber(3)).toNumber()] = (_173_v37).Merge(_173_v37);
        _nw26[(new BigNumber(4)).toNumber()] = _173_v37;
        _nw26[(new BigNumber(5)).toNumber()] = _173_v37;
        _nw26[(new BigNumber(6)).toNumber()] = (_module.__default.fm5(new BigNumber((_dafny.Seq.UnicodeFromString("aywqiw")).length), _127_v2, _131_globalState)).Merge(_173_v37);
        _nw26[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(377),new BigNumber((function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of (_177_v41).Keys.Elements) {
            let _179_v42 = _compr_13;
            if ((_177_v41).contains(_179_v42)) {
              _coll13.add((_179_v42).multipliedBy(new BigNumber(51)));
            }
          }
          return _coll13;
        }()).length));
        _nw26[(new BigNumber(8)).toNumber()] = (_173_v37).Merge(_module.__default.fm5(_132_v6, _127_v2, _131_globalState));
        _nw26[(new BigNumber(9)).toNumber()] = _173_v37;
        _178_v43 = _nw26;
        let _index28 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_178_v43).length));
        (_178_v43)[_index28] = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,new BigNumber((_134_v8).length));
        let _180_v44;
        _180_v44 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("nagnmac"));
        let _181_v45;
        _181_v45 = _module.D1.create_DC2(_180_v44, _129_v4);
        let _182_v46;
        let _nw27 = Array((new BigNumber(7)).toNumber());
        _182_v46 = _nw27;
        let _183_v47;
        _183_v47 = _dafny.Map.Empty.slice().updateUnsafe(_140_cf0,((_127_v2) ? (_182_v46) : (_182_v46)));
        let _184_v48;
        _184_v48 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_173_v37);
        let _index29 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_178_v43).length));
        let _rhs35 = ((_173_v37).update(_132_v6, _132_v6)).Merge((((_184_v48).contains(_132_v6)) ? ((_184_v48).get(_132_v6)) : (_173_v37)));
        let _rhs36 = _181_v45;
        let _rhs37 = (_183_v47).Merge((_183_v47).Merge(_dafny.Map.Empty.slice().updateUnsafe(_138_v11,_182_v46)));
        let _rhs38 = (_132_v6).isLessThanOrEqualTo(_132_v6);
        let _lhs24 = _178_v43;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_178_v43).length));
        _lhs24[_lhs25] = _rhs35;
        _181_v45 = _rhs36;
        _183_v47 = _rhs37;
        _127_v2 = _rhs38;
        _127_v2 = false;
      }
      _module.__default.m0(_127_v2, (_dafny.ZERO).minus(_132_v6), _129_v4, (_132_v6).isLessThanOrEqualTo(new BigNumber(-278)), _131_globalState);
      _127_v2 = _127_v2;
      let _index30 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
      (_128_v3)[_index30] = _127_v2;
      let _index31 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
      (_128_v3)[_index31] = !(_127_v2);
      let _185_v49;
      _185_v49 = _dafny.Seq.of(_module.__default.fm4((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))], _132_v6, _131_globalState));
      let _186_v50;
      _186_v50 = _dafny.Map.Empty.slice().updateUnsafe(_185_v49,_132_v6);
      let _187_v51;
      _187_v51 = _dafny.Map.Empty.slice().updateUnsafe((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))],new BigNumber((_dafny.Seq.update(_125_v0, _module.__default.safeIndex(new BigNumber(122), new BigNumber((_125_v0).length)), _129_v4)).length));
      let _188_v52;
      _188_v52 = _dafny.Map.Empty.slice().updateUnsafe(_186_v50,_187_v51);
      _188_v52 = (_188_v52).update((_dafny.Map.Empty.slice().updateUnsafe(_185_v49,_132_v6)).Merge(function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of (_186_v50).Keys.Elements) {
          let _189_v53 = _compr_14;
          if ((_186_v50).contains(_189_v53)) {
            _coll14.push([_189_v53,new BigNumber((_135_v9).length)]);
          }
        }
        return _coll14;
      }()), _dafny.Map.Empty.slice().updateUnsafe((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))],_132_v6));
      if ((_127_v2) || ((_132_v6).isLessThanOrEqualTo(_132_v6))) {
        (_131_globalState).f8 = _132_v6;
        let _190_v54;
        _190_v54 = _dafny.Map.Empty.slice().updateUnsafe(_128_v3,_127_v2);
        _190_v54 = ((_190_v54).Merge(_190_v54)).update(_128_v3, true);
        let _191_v55;
        _191_v55 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("pnjsvda"));
        let _192_v56;
        _192_v56 = _dafny.Seq.update(_dafny.Seq.update(_134_v8, _module.__default.safeIndex(new BigNumber((_191_v55).length), new BigNumber((_134_v8).length)), _127_v2), _module.__default.safeIndex(new BigNumber((_125_v0).length), new BigNumber((_dafny.Seq.update(_134_v8, _module.__default.safeIndex(new BigNumber((_191_v55).length), new BigNumber((_134_v8).length)), _127_v2)).length)), false);
        _134_v8 = (_192_v56);
        let _193_v57;
        let _nw28 = new _module.C0();
        _nw28.__ctor();
        _193_v57 = _nw28;
        let _194_v58;
        _194_v58 = _dafny.Seq.of(_193_v57);
        let _195_v59;
        _195_v59 = _dafny.Set.fromElements(_193_v57);
        let _196_v60;
        _196_v60 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_193_v57);
        let _197_v61;
        _197_v61 = _module.D3.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), ((_198_v4) => function (_199_i5) {
  return _198_v4;
})(_129_v4)));
        let _200_v62;
        _200_v62 = _dafny.Seq.of(_138_v11, _module.__default.fm6(new _dafny.CodePoint('i'.codePointAt(0)), (_197_v61).dtor_cf5, _127_v2, _131_globalState));
        _module.__default.m0((_dafny.Set.fromElements((_194_v58)[_module.__default.safeIndex(_132_v6, new BigNumber((_194_v58).length))], _193_v57, _193_v57, _193_v57, _193_v57)).IsDisjointFrom(_195_v59), new BigNumber(((_196_v60).Merge(_dafny.Map.Empty.slice().updateUnsafe(_132_v6,_193_v57))).length), _129_v4, !_dafny.areEqual(_200_v62, _dafny.Seq.of(_138_v11, _138_v11)), _131_globalState);
        let _201_v63;
        let _init2 = ((_202_v6) => function (_203_i6) {
          return (_203_i6).multipliedBy(_202_v6);
        })(_132_v6);
        let _nw29 = Array((new BigNumber(3)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw29.length); _i0_2++) {
          _nw29[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _201_v63 = _nw29;
        _201_v63 = _201_v63;
      } else {
        let _204_v64;
        _204_v64 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_132_v6);
        _204_v64 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_125_v0).length),_132_v6)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_132_v6,_132_v6));
        let _205_v65;
        _205_v65 = _module.D3.create_DC4(_125_v0);
        let _206_v66;
        _206_v66 = _dafny.Seq.of(_205_v65, _205_v65);
        let _207_v67;
        _207_v67 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_131_globalState),_206_v66);
        let _208_v68;
        _208_v68 = _dafny.Map.Empty.slice().updateUnsafe(!((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]),(((_207_v67).contains((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))])) ? ((_207_v67).get((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))])) : (_dafny.Seq.of(_205_v65))));
        let _209_v69;
        _209_v69 = _module.D4.create_DC6(_206_v66);
        _207_v67 = (_207_v67).update((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))], (_209_v69).dtor_cf8);
        let _210_v70;
        _210_v70 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_127_v2);
        let _index32 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        (_128_v3)[_index32] = !(_210_v70).equals(_module.__default.fm7(_131_globalState));
        let _index33 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        (_128_v3)[_index33] = !(_127_v2);
        _135_v9 = _135_v9;
      }
      _127_v2 = _127_v2;
      let _211_v71;
      let _nw30 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
      _211_v71 = _nw30;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_211_v71).length))) {
        let _212_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_212_i7)) && ((_212_i7).isLessThan(new BigNumber((_211_v71).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_211_v71, (_212_i7).toNumber(), ((_dafny.Map.Empty.slice().updateUnsafe(_134_v8,(_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))])).Merge(function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of (function () {
              let _coll16 = new _dafny.Map();
              for (const _compr_16 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(862)), function (_213_i8) {
                return _dafny.Seq.of(false);
              })).Elements) {
                let _214_v73 = _compr_16;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(862)), function (_213_i8) {
                  return _dafny.Seq.of(false);
                }), _214_v73)) {
                  _coll16.push([_214_v73,_dafny.Map.Empty.slice().updateUnsafe(!(true),_132_v6)]);
                }
              }
              return _coll16;
            }()).Keys.Elements) {
              let _215_v72 = _compr_15;
              if ((function () {
                let _coll17 = new _dafny.Map();
                for (const _compr_17 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(862)), function (_213_i8) {
                  return _dafny.Seq.of(false);
                })).Elements) {
                  let _214_v73 = _compr_17;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(862)), function (_213_i8) {
                    return _dafny.Seq.of(false);
                  }), _214_v73)) {
                    _coll17.push([_214_v73,_dafny.Map.Empty.slice().updateUnsafe(!(true),_132_v6)]);
                  }
                }
                return _coll17;
              }()).contains(_215_v72)) {
                _coll15.push([_215_v72,(_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]]);
              }
            }
            return _coll15;
          }())).Merge((((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]) ? (_dafny.Map.Empty.slice().updateUnsafe(_134_v8,!((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]))) : (_dafny.Map.Empty.slice().updateUnsafe(_134_v8,(_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]))))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _216_v74;
      let _nw31 = new _module.C0();
      _nw31.__ctor();
      _216_v74 = _nw31;
      let _217_v75;
      _217_v75 = _dafny.Map.Empty.slice().updateUnsafe(_132_v6,_132_v6);
      _125_v0 = _dafny.Seq.Concat(_module.__default.fm8(_132_v6, _217_v75, new BigNumber((_185_v49).length), _131_globalState), _dafny.Seq.Concat(_125_v0, _125_v0));
      let _218_v76;
      _218_v76 = _dafny.Map.Empty.slice().updateUnsafe(_216_v74,_132_v6);
      _218_v76 = (_218_v76).update(_216_v74, _132_v6);
      _127_v2 = (_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))];
      let _219_v77;
      let _nw32 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _219_v77 = _nw32;
      _219_v77 = _219_v77;
      let _220_v78;
      _220_v78 = _dafny.MultiSet.fromElements(_132_v6);
      let _hi1 = (new BigNumber((_220_v78).cardinality())).multipliedBy(_132_v6);
      for (let _221_i9 = _module.__default.fm4(false, new BigNumber((_dafny.Seq.update(_134_v8, _module.__default.safeIndex(new BigNumber(79), new BigNumber((_134_v8).length)), (_module.D1.create_DC1((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))])).dtor_cf1)).length), _131_globalState); _221_i9.isLessThan(_hi1); _221_i9 = _221_i9.plus(_dafny.ONE)) {
        let _222_v79;
        _222_v79 = _dafny.Seq.of(_128_v3, _128_v3, _128_v3);
        _222_v79 = _dafny.Seq.of(_128_v3);
        let _223_v80;
        _223_v80 = _219_v77;
        let _224_v81;
        _224_v81 = _dafny.Seq.of(_219_v77, (_223_v80), _219_v77, _219_v77);
        _224_v81 = _224_v81;
        let _225_v82;
        _225_v82 = _module.D3.create_DC5(_127_v2, (_185_v49)[_module.__default.safeIndex(_221_i9, new BigNumber((_185_v49).length))]);
        let _index34 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        (_128_v3)[_index34] = (((_225_v82).dtor_cf7).multipliedBy(new BigNumber(-792))).isLessThan(_132_v6);
        let _index35 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        let _index36 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        let _rhs39 = !(_127_v2);
        let _rhs40 = ((_137_v10).Intersect(_137_v10)).IsProperSubsetOf(_137_v10);
        let _rhs41 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]), _134_v8), _dafny.Seq.of(_127_v2, (_128_v3)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length))]))).length);
        let _rhs42 = (_dafny.Set.fromElements(_132_v6)).Difference(_135_v9);
        let _rhs43 = _216_v74;
        let _lhs26 = _128_v3;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        let _lhs28 = _128_v3;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_128_v3).length));
        let _lhs30 = _131_globalState;
        _lhs26[_lhs27] = _rhs39;
        _lhs28[_lhs29] = _rhs40;
        _lhs30.f8 = _rhs41;
        _135_v9 = _rhs42;
        _216_v74 = _rhs43;
      }
      process.stdout.write((_125_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_129_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v5).equals(_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_131_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_131_globalState.f2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_131_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f9).equals(_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_132_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_133_v7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_134_v8, _dafny.Seq.of(true, true, false, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v9).equals(_dafny.Set.fromElements(new BigNumber(520), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_v10).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_v11)).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_185_v49, _dafny.Seq.of(new BigNumber(549)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v50).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(549)),new BigNumber(520)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v51).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v52).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(549)),new BigNumber(520)),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(2))).updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(549)),new BigNumber(2)),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(520))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_211_v71)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true, false, false, true),false).updateUnsafe(_dafny.Seq.of(false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_211_v71)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true, false, false, true),false).updateUnsafe(_dafny.Seq.of(false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_211_v71)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true, false, false, true),false).updateUnsafe(_dafny.Seq.of(false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_211_v71)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true, false, false, true),false).updateUnsafe(_dafny.Seq.of(false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_217_v75).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(520),new BigNumber(520)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_218_v76).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v78).equals(_dafny.MultiSet.fromElements(new BigNumber(520)))));
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
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
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
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
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
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.Set.Empty, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC3(cf4) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf4) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC5(cf6, cf7) {
      let $dt = new D3(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC4(cf5) {
      let $dt = new D3(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC4" + "(" + this.cf5.toVerbatimString(true) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC5(false, _dafny.ZERO);
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
    static create_DC7(cf9, cf10, cf11, cf12) {
      let $dt = new D4(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC8() {
      let $dt = new D4(1);
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D4(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC8";
      } else if (this.$tag === 2) {
        return "D4.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9 && this.cf10 === other.cf10 && this.cf11 === other.cf11 && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC7(false, null, false, false);
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
    static create_DC9(cf13) {
      let $dt = new D5(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC9" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13;
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf15, cf16) {
      let $dt = new D6(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC12(cf17, cf18) {
      let $dt = new D6(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D6(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC10(cf14) {
      let $dt = new D6(3);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC10" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf17 === other.cf17 && this.cf18 === other.cf18;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC11(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC15(cf21, cf22) {
      let $dt = new D7(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC16(cf23, cf24, cf25, cf26, cf27) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC14(cf20) {
      let $dt = new D7(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC17(cf28) {
      let $dt = new D7(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get is_DC17() { return this.$tag === 3; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(_dafny.ZERO, _dafny.ZERO);
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
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f2 = _dafny.Seq.UnicodeFromString("");
      this.f8 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f3 = [];
      this._f4 = false;
      this._f5 = false;
      this._f6 = [];
      this._f7 = _dafny.ZERO;
      this._f9 = _dafny.Set.Empty;
      this._f10 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
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
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
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
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
