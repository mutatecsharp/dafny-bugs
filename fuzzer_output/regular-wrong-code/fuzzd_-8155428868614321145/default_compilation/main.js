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
      return (_module.D21.create_DC62(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),true), false)).dtor_cf105;
    };
    static fm1(globalState) {
      return new BigNumber(969);
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_0_i0) {
          return _dafny.Seq.of(false);
        })).Elements) {
          let _1_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_0_i0) {
            return _dafny.Seq.of(false);
          }), _1_v0)) {
            _coll0.push([_1_v0,_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))).length))).length),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,true), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,false))).cardinality()))).length))]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, false),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(760))));
    };
    static fm3(p0, p1, globalState) {
      return _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_2_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })).length), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).length),!(false))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kfo")).length),new BigNumber(-676))).Keys.Elements) {
          let _3_v0 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kfo")).length),new BigNumber(-676))).contains(_3_v0)) {
            _coll1.push([(_3_v0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_4_i1) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            })).length)),true]);
          }
        }
        return _coll1;
      }())).length), (_dafny.ZERO).minus((new BigNumber(765)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(129))).length))), _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ny")).length), new BigNumber(382)), new BigNumber(-840));
    };
    static fm4(p0, p1, p2, globalState) {
      if ((_dafny.MultiSet.fromElements(_module.D7.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(449))).length),new BigNumber(478))))).IsSubsetOf(_dafny.MultiSet.fromElements(_module.D7.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D27.create_DC72(_dafny.Seq.of(new BigNumber(405)))).dtor_cf123).length),new BigNumber(445))), _module.D7.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, true, false, false, false), _dafny.Seq.of(false, true))).length),false)).length),new BigNumber((_dafny.Seq.of(new BigNumber(52))).length))), _module.D7.create_DC22(function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(33), new BigNumber(380))) {
    let _5_v0 = _compr_2;
    if (((new BigNumber(33)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(380)))) {
      _coll2.push([(_5_v0).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), true, true, true))).cardinality())),new BigNumber((_dafny.Set.fromElements(new BigNumber(839), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length))).length)]);
    }
  }
  return _coll2;
}()), _module.D7.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(257), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false, true)).length), new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((function () {
  let _coll3 = new _dafny.Set();
  for (const _compr_3 of (_dafny.Seq.of(_dafny.Seq.of(true, !(true)))).Elements) {
    let _6_v1 = _compr_3;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(true, !(true))), _6_v1)) {
      _coll3.add(_6_v1);
    }
  }
  return _coll3;
}()).length))).length))).length))))))) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }
    };
    static fm11(p0, p1, p2, globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(true,false)) : (_dafny.Map.Empty.slice().updateUnsafe(true,false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(((false) ? (_dafny.Seq.UnicodeFromString("txrnv")) : (_dafny.Seq.UnicodeFromString("wvkpmu"))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pwry"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-302)), function (_7_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("y"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), function (_8_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })));
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yuwq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_9_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("s"));
    };
    static fm18(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.of(!(false))).length))).length)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("taessl")).length)), new BigNumber(129), (new BigNumber(125)).plus(new BigNumber(-426)), (new BigNumber(736)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true)).length))).length)));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D18.create_DC56(_dafny.Seq.of(true));
      if (_source0.is_DC54) {
        let _10___mcc_h0 = (_source0).cf91;
        let _11___mcc_h1 = (_source0).cf92;
        let _12___mcc_h2 = (_source0).cf93;
        let _13___mcc_h3 = (_source0).cf94;
        let _14_cf94 = _13___mcc_h3;
        let _15_cf93 = _12___mcc_h2;
        let _16_cf92 = _11___mcc_h1;
        let _17_cf91 = _10___mcc_h0;
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(true)))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.of((_module.D27.create_DC73(_15_cf93, _16_cf92, _dafny.Set.fromElements(_15_cf93, true))).dtor_cf124, _15_cf93)));
      } else if (_source0.is_DC55) {
        let _18___mcc_h4 = (_source0).cf95;
        let _19___mcc_h5 = (_source0).cf96;
        let _20_cf96 = _19___mcc_h5;
        let _21_cf95 = _18___mcc_h4;
        if (false) {
          return _dafny.MultiSet.fromElements(_dafny.Seq.of(true, false));
        } else {
          return _dafny.MultiSet.fromElements(_dafny.Seq.of(true, true, true), _dafny.Seq.of(true));
        }
      } else if (_source0.is_DC56) {
        let _22___mcc_h6 = (_source0).cf97;
        let _23_cf97 = _22___mcc_h6;
        return _dafny.MultiSet.fromElements(_dafny.Seq.of(false), _23_cf97, _23_cf97);
      } else {
        let _24___mcc_h7 = (_source0).cf90;
        let _25_cf90 = _24___mcc_h7;
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(false)));
      }
    };
    static fm24(globalState) {
      return (_dafny.MultiSet.fromElements(true)).Union((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.UnicodeFromString("b");
    };
    static fm26(globalState) {
      return ((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()))).length)))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-77), new BigNumber(196))).length)))).Difference((_dafny.Set.fromElements(new BigNumber(382))).Difference(_dafny.Set.fromElements(new BigNumber(-244))));
    };
    static fm28(p0, p1, globalState) {
      return function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("dk"), _dafny.Seq.UnicodeFromString("yxusns"))).Elements) {
          let _26_v0 = _compr_4;
          if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("dk"), _dafny.Seq.UnicodeFromString("yxusns"))).contains(_26_v0)) {
            _coll4.add(_26_v0);
          }
        }
        return _coll4;
      }();
    };
    static fm29(p0, p1, p2, globalState) {
      return _module.D3.create_DC11((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), function (_27_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("iqsfuxsch")).length);
})).length)).minus(new BigNumber((_dafny.Set.fromElements(!(!(true)))).length)), (_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm30(globalState) {
      return _dafny.Set.fromElements(!(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("gtved"), _dafny.Seq.UnicodeFromString("vhtg")))), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("g"), _dafny.Seq.UnicodeFromString("hd")));
    };
    static fm31(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_module.D19.create_DC58(new BigNumber((_dafny.Seq.UnicodeFromString("hekts")).length), !(true))).dtor_cf100)).length), new BigNumber(275))).Intersect(_dafny.Set.fromElements(new BigNumber(((_module.D28.create_DC76(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(344))).length)))).dtor_cf134).length)));
    };
    static fm32(p0, p1, p2, globalState) {
      return _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("mijxeu"));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm34(p0, globalState) {
      return _module.D2.create_DC8(_module.D2.create_DC7(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), function (_28_i0) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length), true, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gjqrabqx")).length)), new BigNumber((_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(789),!(true))).length))), _dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length)), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(111), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))).length))).length))).length))))).length), new BigNumber(-233), new BigNumber(-262))).length)));
    };
    static fm35(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new BigNumber(451), (new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, !(!(false)))).cardinality())))).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-630)), function (_29_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length)));
    };
    static fm38(p0, p1, globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), function (_30_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })) : (_dafny.Seq.UnicodeFromString("xrwuvm"))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dbhrr"), _dafny.Seq.UnicodeFromString("ctmp")));
    };
    static fm39(p0, p1, p2, globalState) {
      return _module.D2.create_DC6(new _dafny.CodePoint('k'.codePointAt(0)));
    };
    static fm40(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(false));
    };
    static fm41(globalState) {
      return (_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements(false, true, true, false, false));
    };
    static fm42(globalState) {
      let _source1 = _module.D7.create_DC24(new BigNumber(-767));
      if (_source1.is_DC23) {
        return _module.D11.create_DC32(function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-528), new BigNumber(-312))) {
    let _31_v0 = _compr_5;
    if (((new BigNumber(-528)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(-312)))) {
      _coll5.push([_module.__default.safeModuloInt(_31_v0, new BigNumber(66)),!(true)]);
    }
  }
  return _coll5;
}());
      } else if (_source1.is_DC24) {
        let _32___mcc_h0 = (_source1).cf41;
        let _33_cf41 = _32___mcc_h0;
        return _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe(_33_cf41,true));
      } else {
        let _34___mcc_h1 = (_source1).cf40;
        let _35_cf40 = _34___mcc_h1;
        return _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),!(true)));
      }
    };
    static fm43(p0, p1, p2, globalState) {
      let _source2 = _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),true));
      if (_source2.is_DC14) {
        let _36___mcc_h0 = (_source2).cf21;
        let _37___mcc_h1 = (_source2).cf22;
        let _38_cf22 = _37___mcc_h1;
        let _39_cf21 = _36___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_module.D5.create_DC18(_module.D5.create_DC17(false, new _dafny.CodePoint('v'.codePointAt(0)))), _module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC17(false, new _dafny.CodePoint('k'.codePointAt(0)))))), _dafny.Seq.of(_module.D5.create_DC18(_module.D5.create_DC17(true, new _dafny.CodePoint('e'.codePointAt(0)))), _module.D5.create_DC18(_module.D5.create_DC16(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)))), _module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC16(_dafny.MultiSet.fromElements(true))))));
      } else if (_source2.is_DC15) {
        let _40___mcc_h2 = (_source2).cf23;
        let _41___mcc_h3 = (_source2).cf24;
        let _42___mcc_h4 = (_source2).cf25;
        let _43___mcc_h5 = (_source2).cf26;
        let _44_cf26 = _43___mcc_h5;
        let _45_cf25 = _42___mcc_h4;
        let _46_cf24 = _41___mcc_h3;
        let _47_cf23 = _40___mcc_h2;
        return _dafny.Seq.of(_module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC16(_dafny.MultiSet.fromElements(true)))));
      } else {
        let _48___mcc_h6 = (_source2).cf20;
        let _49_cf20 = _48___mcc_h6;
        return _dafny.Seq.Concat(_dafny.Seq.of(_module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC16(_dafny.MultiSet.fromElements(false))))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), function (_50_i0) {
          return _module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC16(_dafny.MultiSet.fromElements(false, !(false))))));
        }));
      }
    };
    static fm44(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(459)), function (_51_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }),new BigNumber(153))).Merge(function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ajwlwtmnx"))).Elements) {
          let _52_v0 = _compr_6;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ajwlwtmnx")), _52_v0)) {
            _coll6.push([_52_v0,new BigNumber(264)]);
          }
        }
        return _coll6;
      }());
    };
    static fm45(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tp")).length),new BigNumber(-214));
    };
    static fm46(p0, p1, p2, globalState) {
      if ((_module.D2.create_DC7(new BigNumber((_dafny.Seq.of(false, true)).length), true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true)).length))).dtor_cf10) {
        return _module.D10.create_DC31(!(false), false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length));
      } else {
        return _module.D10.create_DC31(true, true, new BigNumber(-559));
      }
    };
    static fm47(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))), _dafny.Seq.of(new BigNumber(884)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, false)).length)), _dafny.Seq.of((_module.D1.create_DC3(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(770))).cardinality()))).dtor_cf4))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), function (_53_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length), new BigNumber(900))))).Union((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false))).length)))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(143)), function (_54_i1) {
        return new BigNumber(535);
      }))));
    };
    static fm48(globalState) {
      return (_dafny.MultiSet.fromElements(_module.D2.create_DC8(_module.D2.create_DC8(_module.D2.create_DC7(new BigNumber((_dafny.Seq.of(true)).length), false, new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(false, false))).length)))), _module.D2.create_DC8(_module.D2.create_DC6(new _dafny.CodePoint('a'.codePointAt(0)))))).Difference(_dafny.MultiSet.fromElements(_module.D2.create_DC8(_module.D2.create_DC6(new _dafny.CodePoint('o'.codePointAt(0))))));
    };
    static fm49(p0, p1, globalState) {
      return _module.D3.create_DC9((_dafny.Set.fromElements(_dafny.Set.fromElements(!(!(!(true)))), _dafny.Set.fromElements(true))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(true), _dafny.Set.fromElements(true))));
    };
    static fm50(p0, p1, globalState) {
      if (true) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(419)), function (_55_i0) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length),new _dafny.CodePoint('a'.codePointAt(0)))).Merge(function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_dafny.Seq.of(new BigNumber(631))).Elements) {
            let _56_v0 = _compr_7;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(631)), _56_v0)) {
              _coll7.push([(_56_v0).plus(new BigNumber(-154)),new _dafny.CodePoint('n'.codePointAt(0))]);
            }
          }
          return _coll7;
        }());
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(827),new _dafny.CodePoint('q'.codePointAt(0)));
      }
    };
    static fm51(globalState) {
      let _source3 = _module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC16(_dafny.MultiSet.fromElements(true))))));
      if (_source3.is_DC17) {
        let _57___mcc_h0 = (_source3).cf28;
        let _58___mcc_h1 = (_source3).cf29;
        let _59_cf29 = _58___mcc_h1;
        let _60_cf28 = _57___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(120))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_60_cf28,new BigNumber(-616)));
      } else if (_source3.is_DC16) {
        let _61___mcc_h2 = (_source3).cf27;
        let _62_cf27 = _61___mcc_h2;
        return (_module.D28.create_DC76(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Set.fromElements(new BigNumber(812))).length)))).dtor_cf134;
      } else {
        let _63___mcc_h3 = (_source3).cf30;
        let _64_cf30 = _63___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber(-316)));
      }
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(false, true), _dafny.Seq.of(false), _dafny.Seq.of(true, true), _dafny.Seq.of(!(false)), _dafny.Seq.of(false, true)), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(true, false, false), _dafny.Seq.of(true, false), _dafny.Seq.of(true), _dafny.Seq.of(!(!(!(true))), !(true)), _dafny.Seq.of(false, true)), _dafny.Seq.of(_dafny.Seq.of(true))));
    };
    static fm53(p0, p1, globalState) {
      return _module.D14.create_DC43(new _dafny.CodePoint('t'.codePointAt(0)), new BigNumber(487), new BigNumber(304), (new BigNumber(-702)).multipliedBy(new BigNumber(-457)));
    };
    static fm54(globalState) {
      let _source4 = _module.D10.create_DC31(false, false, new BigNumber((_dafny.Seq.UnicodeFromString("yiqmqaeh")).length));
      if (_source4.is_DC31) {
        let _65___mcc_h0 = (_source4).cf50;
        let _66___mcc_h1 = (_source4).cf51;
        let _67___mcc_h2 = (_source4).cf52;
        let _68_cf52 = _67___mcc_h2;
        let _69_cf51 = _66___mcc_h1;
        let _70_cf50 = _65___mcc_h0;
        return _module.D4.create_DC15(_68_cf52, _dafny.MultiSet.fromElements(_68_cf52), new _dafny.CodePoint('r'.codePointAt(0)), _68_cf52);
      } else {
        let _71___mcc_h3 = (_source4).cf49;
        let _72_cf49 = _71___mcc_h3;
        return _module.D4.create_DC15(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).cardinality()))).length), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, false)).length)), new _dafny.CodePoint('f'.codePointAt(0)), new BigNumber(677));
      }
    };
    static fm55(p0, p1, p2, globalState) {
      let _source5 = ((false) ? (_module.D20.create_DC60(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))) : (_module.D20.create_DC60(new BigNumber(589))));
      if (_source5.is_DC60) {
        let _73___mcc_h0 = (_source5).cf102;
        let _74_cf102 = _73___mcc_h0;
        return _module.D18.create_DC53(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),!(false)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false)));
      } else {
        let _75___mcc_h1 = (_source5).cf101;
        let _76_cf101 = _75___mcc_h1;
        return _module.D18.create_DC53(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),true)));
      }
    };
    static fm56(globalState) {
      return _module.D16.create_DC49(new BigNumber(394));
    };
    static fm57(globalState) {
      return _module.D15.create_DC45(new BigNumber(-864), (_module.D3.create_DC10(_dafny.Seq.UnicodeFromString("fsgxfjmq"), false, true)).dtor_cf16, false);
    };
    static fm58(globalState) {
      return _module.D19.create_DC58((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(!(false))).length), new BigNumber(814))).cardinality())).plus(new BigNumber((_dafny.Seq.of(false)).length)), !((new BigNumber(757)).isLessThanOrEqualTo(new BigNumber(805))));
    };
    static fm59(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat((_module.D29.create_DC79((_module.D29.create_DC79(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), function (_77_i0) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(990),true)).length);
}), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length)), _dafny.Seq.of(new BigNumber((function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of _dafny.IntegerRange(new BigNumber(464), new BigNumber(352))) {
    let _78_v0 = _compr_8;
    if (((new BigNumber(464)).isLessThanOrEqualTo(_78_v0)) && ((_78_v0).isLessThan(new BigNumber(352)))) {
      _coll8.push([_module.__default.safeModuloInt(_78_v0, new BigNumber((_dafny.Seq.of(false)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)]);
    }
  }
  return _coll8;
}()).length))))).dtor_cf141)).dtor_cf141, _dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_79_i1) {
        return _dafny.Seq.of(new BigNumber(952), new BigNumber((_dafny.Seq.of(_module.D18.create_DC53(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),false))), _module.D18.create_DC53(_dafny.Seq.of(function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("doi")).length)))).length))).Keys.Elements) {
    let _80_v1 = _compr_9;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("doi")).length)))).length))).contains(_80_v1)) {
      _coll9.push([_80_v1,true]);
    }
  }
  return _coll9;
}(), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),false))))).length));
      })), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(485), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-386)), function (_81_i2) {
        return (_module.D6.create_DC21(true, true, true, new BigNumber(400))).dtor_cf39;
      }))).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_82_i3) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(925),false)).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(250))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(345),false)).length))));
    };
    static fm60(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(!(true)), false)).cardinality())), new BigNumber(141)),function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).Elements) {
          let _83_v0 = _compr_10;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))), _83_v0)) {
            _coll10.push([_83_v0,true]);
          }
        }
        return _coll10;
      }());
    };
    static fm61(p0, p1, p2, p3, globalState) {
      return _module.D7.create_DC22((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-6),new BigNumber(-792))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),new BigNumber((_dafny.Seq.UnicodeFromString("redxim")).length))));
    };
    static fm62(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(false, true)).Difference(_dafny.MultiSet.fromElements(!(false), !(!(false)))),(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(129),false)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(236)), function (_84_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(true, !(true))).cardinality());
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length)))).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),new BigNumber(483))).length))))));
    };
    static fm63(globalState) {
      let _source6 = _module.D18.create_DC56(_dafny.Seq.of(false));
      if (_source6.is_DC54) {
        let _85___mcc_h0 = (_source6).cf91;
        let _86___mcc_h1 = (_source6).cf92;
        let _87___mcc_h2 = (_source6).cf93;
        let _88___mcc_h3 = (_source6).cf94;
        let _89_cf94 = _88___mcc_h3;
        let _90_cf93 = _87___mcc_h2;
        let _91_cf92 = _86___mcc_h1;
        let _92_cf91 = _85___mcc_h0;
        return _module.D5.create_DC16(_dafny.MultiSet.fromElements(_90_cf93, _90_cf93, _90_cf93));
      } else if (_source6.is_DC55) {
        let _93___mcc_h4 = (_source6).cf95;
        let _94___mcc_h5 = (_source6).cf96;
        let _95_cf96 = _94___mcc_h5;
        let _96_cf95 = _93___mcc_h4;
        return _module.D5.create_DC16(_dafny.MultiSet.fromElements(false));
      } else if (_source6.is_DC56) {
        let _97___mcc_h6 = (_source6).cf97;
        let _98_cf97 = _97___mcc_h6;
        return _module.D5.create_DC16((_module.D5.create_DC16(_dafny.MultiSet.fromElements(false))).dtor_cf27);
      } else {
        let _99___mcc_h7 = (_source6).cf90;
        let _100_cf90 = _99___mcc_h7;
        return _module.D5.create_DC16(_dafny.MultiSet.fromElements(true, !(true)));
      }
    };
    static fm64(globalState) {
      let _source7 = _module.D5.create_DC16(_dafny.MultiSet.fromElements(!(false), false, false, true));
      if (_source7.is_DC17) {
        let _101___mcc_h0 = (_source7).cf28;
        let _102___mcc_h1 = (_source7).cf29;
        let _103_cf29 = _102___mcc_h1;
        let _104_cf28 = _101___mcc_h0;
        return _module.D10.create_DC30(_dafny.Set.fromElements(_104_cf28, !(_104_cf28), _104_cf28));
      } else if (_source7.is_DC16) {
        let _105___mcc_h2 = (_source7).cf27;
        let _106_cf27 = _105___mcc_h2;
        return _module.D10.create_DC30(_dafny.Set.fromElements(!(true)));
      } else {
        let _107___mcc_h3 = (_source7).cf30;
        let _108_cf30 = _107___mcc_h3;
        return _module.D10.create_DC30(_dafny.Set.fromElements(!(true)));
      }
    };
    static fm65(p0, p1, p2, globalState) {
      let _source8 = _module.D28.create_DC77((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),_dafny.Set.fromElements(false))).length)), _dafny.Set.fromElements(true), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_109_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})).length), true, _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(448), new BigNumber((_dafny.Set.fromElements(true)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-690))).cardinality()))),true));
      if (_source8.is_DC77) {
        let _110___mcc_h0 = (_source8).cf135;
        let _111___mcc_h1 = (_source8).cf136;
        let _112___mcc_h2 = (_source8).cf137;
        let _113___mcc_h3 = (_source8).cf138;
        let _114___mcc_h4 = (_source8).cf139;
        let _115_cf139 = _114___mcc_h4;
        let _116_cf138 = _113___mcc_h3;
        let _117_cf137 = _112___mcc_h2;
        let _118_cf136 = _111___mcc_h1;
        let _119_cf135 = _110___mcc_h0;
        return _module.D3.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-969)), function (_120_i1) {
  return new _dafny.CodePoint('e'.codePointAt(0));
}), _116_cf138, !(_116_cf138));
      } else if (_source8.is_DC76) {
        let _121___mcc_h5 = (_source8).cf134;
        let _122_cf134 = _121___mcc_h5;
        return _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("eicihmi"), !(false), false);
      } else {
        let _123___mcc_h6 = (_source8).cf140;
        let _124_cf140 = _123___mcc_h6;
        return _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("gmdp"), !(false), false);
      }
    };
    static fm66(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),(_module.D27.create_DC73(true, new BigNumber(-640), _dafny.Set.fromElements(true))).dtor_cf124));
    };
    static fm67(p0, p1, p2, globalState) {
      return _module.D21.create_DC63(new BigNumber((_dafny.Seq.UnicodeFromString("swsbxe")).length), (true) && (false), !((_dafny.Set.fromElements(new BigNumber(862), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Seq.of(!(true), !(false), false)).length), new BigNumber((_dafny.Set.fromElements(true)).length))).IsProperSubsetOf(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()))))), (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-712), new BigNumber(727), new BigNumber(993))).length))).length)).isEqualTo(new BigNumber(658)));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.Seq.UnicodeFromString("");
      let _125_v0;
      _125_v0 = false;
      let _126_v1;
      _126_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,_125_v0);
      let _127_v2;
      let _init0 = ((_128_p0) => function (_129_i0) {
        return (_129_i0).multipliedBy(_128_p0);
      })(p0);
      let _nw0 = Array((new BigNumber(21)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _127_v2 = _nw0;
      let _index0 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
      (_127_v2)[_index0] = p0;
      let _130_v3;
      _130_v3 = _dafny.Seq.of(_125_v0);
      let _131_v4;
      _131_v4 = _dafny.Seq.of(new BigNumber(114), p0);
      let _132_v5;
      _132_v5 = _dafny.MultiSet.fromElements(_125_v0, _125_v0);
      let _133_v6;
      _133_v6 = new _dafny.CodePoint('h'.codePointAt(0));
      let _134_v7;
      _134_v7 = _dafny.Seq.of(_133_v6, _133_v6);
      let _135_v8;
      _135_v8 = _module.D11.create_DC34((_dafny.ZERO).minus(p0), !(_module.__default.fm0(new BigNumber((_132_v5).cardinality()), globalState)), new BigNumber((_134_v7).length));
      let _pat_let_tv0 = _133_v6;
      let _pat_let_tv1 = _134_v7;
      let _pat_let_tv2 = _133_v6;
      let _pat_let_tv3 = _134_v7;
      let _index1 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
      let _rhs0 = _module.__default.fm0(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(((_125_v0) ? (_130_v3) : (_130_v3)), _module.__default.safeIndex(p0, new BigNumber((((_125_v0) ? (_130_v3) : (_130_v3))).length)), _125_v0))).cardinality()), globalState);
      let _rhs1 = true;
      let _rhs2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.areEqual(_131_v4, _131_v4));
      let _rhs3 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(function (_source9) {
        if (_source9.is_DC33) {
          let _136___mcc_h0 = (_source9).cf54;
          let _137___mcc_h1 = (_source9).cf55;
          let _138_cf55 = _137___mcc_h1;
          let _139_cf54 = _136___mcc_h0;
          return _dafny.Seq.UnicodeFromString("dwc");
        } else if (_source9.is_DC34) {
          let _140___mcc_h2 = (_source9).cf56;
          let _141___mcc_h3 = (_source9).cf57;
          let _142___mcc_h4 = (_source9).cf58;
          let _143_cf58 = _142___mcc_h4;
          let _144_cf57 = _141___mcc_h3;
          let _145_cf56 = _140___mcc_h2;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-397)), ((_146_v6) => function (_147_i1) {
            return _146_v6;
          })(_pat_let_tv0));
        } else {
          let _148___mcc_h5 = (_source9).cf53;
          let _149_cf53 = _148___mcc_h5;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("duyjggi"), _pat_let_tv1);
        }
      }(_135_v8), _module.__default.safeIndex(p0, new BigNumber((function (_source10) {
        if (_source10.is_DC33) {
          let _150___mcc_h6 = (_source10).cf54;
          let _151___mcc_h7 = (_source10).cf55;
          let _152_cf55 = _151___mcc_h7;
          let _153_cf54 = _150___mcc_h6;
          return _dafny.Seq.UnicodeFromString("dwc");
        } else if (_source10.is_DC34) {
          let _154___mcc_h8 = (_source10).cf56;
          let _155___mcc_h9 = (_source10).cf57;
          let _156___mcc_h10 = (_source10).cf58;
          let _157_cf58 = _156___mcc_h10;
          let _158_cf57 = _155___mcc_h9;
          let _159_cf56 = _154___mcc_h8;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-397)), ((_160_v6) => function (_161_i1) {
            return _160_v6;
          })(_pat_let_tv2));
        } else {
          let _162___mcc_h11 = (_source10).cf53;
          let _163_cf53 = _162___mcc_h11;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("duyjggi"), _pat_let_tv3);
        }
      }(_135_v8)).length)), _133_v6)).length));
      let _lhs0 = _127_v2;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
      _125_v0 = _rhs0;
      _125_v0 = _rhs1;
      _126_v1 = _rhs2;
      _lhs0[_lhs1] = _rhs3;
      let _164_i2;
      _164_i2 = _dafny.ZERO;
      L0: {
        while (((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]).isEqualTo(_module.__default.safeDivisionInt(p0, new BigNumber(-473)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_164_i2)) {
              break L0;
            }
            _164_i2 = (_164_i2).plus(_dafny.ONE);
            let _165_v9;
            let _nw1 = Array((new BigNumber(5)).toNumber()).fill([]);
            _165_v9 = _nw1;
            let _index2 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_165_v9).length));
            (_165_v9)[_index2] = p1;
            let _index3 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_165_v9).length));
            (_165_v9)[_index3] = ((((_125_v0) ? (!(_125_v0)) : (_125_v0))) ? (p1) : (p1));
            r1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_134_v7, _134_v7), _134_v7);
            (globalState).f7 = (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))];
            if (_125_v0) {
              let _166_v10;
              _166_v10 = _dafny.Map.Empty.slice().updateUnsafe(_125_v0,_125_v0);
              let _167_v11;
              let _nw2 = new _module.C12();
              _nw2.__ctor(_125_v0, _125_v0, _125_v0);
              _167_v11 = _nw2;
              let _168_v12;
              _168_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_166_v10).length),_167_v11);
              let _169_v13;
              _169_v13 = _dafny.Seq.of((((_168_v12).contains((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))])) ? ((_168_v12).get((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))])) : (_167_v11)), _167_v11);
              let _170_v14;
              _170_v14 = _dafny.Seq.of(_dafny.Seq.update(_130_v3, _module.__default.safeIndex(p0, new BigNumber((_130_v3).length)), (_167_v11).f10), _130_v3, _130_v3, _130_v3, _dafny.Seq.of((_167_v11).f10));
              _169_v13 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_169_v13, _169_v13), _module.__default.safeIndex(new BigNumber((_170_v14).length), new BigNumber((_dafny.Seq.Concat(_169_v13, _169_v13)).length)), _167_v11), _169_v13);
              let _171_v15;
              _171_v15 = _dafny.Set.fromElements(p0);
              let _172_v16;
              _172_v16 = _dafny.Set.fromElements((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))], p0, (_dafny.ZERO).minus(new BigNumber((_171_v15).length)), p0);
              _125_v0 = ((_171_v15).Intersect(_171_v15)).IsSubsetOf(_dafny.Set.fromElements(new BigNumber(-157)));
              _127_v2 = _127_v2;
              _131_v4 = _module.__default.fm3(new BigNumber(142), (((_167_v11).f10) ? (p0) : (new BigNumber(973))), globalState);
              r0 = _134_v7;
            } else {
              (globalState).f0 = _module.__default.safeDivisionInt(p0, _module.__default.safeDivisionInt(p0, new BigNumber(-689)));
              let _173_v17;
              _173_v17 = _dafny.Map.Empty.slice().updateUnsafe((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))],p0);
              _173_v17 = (_173_v17).update((_dafny.ZERO).minus(p0), _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Set.fromElements(p0, new BigNumber((_134_v7).length))).length)));
              _125_v0 = _module.__default.fm0(_module.__default.fm1(globalState), globalState);
              let _174_v18;
              let _nw3 = Array((new BigNumber(2)).toNumber()).fill(_module.D1.Default());
              _174_v18 = _nw3;
              let _175_v19;
              _175_v19 = _module.D8.create_DC27(_125_v0, _dafny.Seq.of(_174_v18, _174_v18), _131_v4, (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]);
              let _176_v20;
              _176_v20 = _module.D21.create_DC63((_175_v19).dtor_cf46, _125_v0, _125_v0, _125_v0);
              _125_v0 = !(((_dafny.ZERO).minus((_176_v20).dtor_cf106)).isLessThan(((_dafny.ZERO).minus(p0)).multipliedBy((((_132_v5).contains(_125_v0)) ? ((_132_v5).get(_125_v0)) : (new BigNumber(938))))));
              (globalState).f1 = _module.__default.safeDivisionInt(new BigNumber(355), p0);
            }
          }
        }
      }
      if (!(!(_125_v0) || (!(_125_v0))) || (_125_v0)) {
        let _177_v21;
        _177_v21 = _dafny.Seq.of(_134_v7);
        let _178_v22;
        _178_v22 = _dafny.Map.Empty.slice().updateUnsafe(_177_v21,!(_125_v0));
        _178_v22 = (_178_v22).update(_dafny.Seq.of(_134_v7), _125_v0);
        let _179_v23;
        _179_v23 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]));
        let _180_v24;
        let _nw4 = Array((new BigNumber(2)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _179_v23;
        _nw4[(_dafny.ONE).toNumber()] = _179_v23;
        _180_v24 = _nw4;
        _180_v24 = _180_v24;
        (globalState).f1 = p0;
        _125_v0 = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("nblvw"), _133_v6);
        let _181_v25;
        let _nw5 = new _module.C9();
        _nw5.__ctor(_125_v0, _module.__default.fm0(new BigNumber((_module.__default.fm33(_125_v0, _125_v0, _134_v7, _125_v0, globalState)).length), globalState));
        _181_v25 = _nw5;
      } else {
        (globalState).f0 = p0;
        let _182_v26;
        _182_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(801));
        if ((((((_132_v5).contains(_125_v0)) ? ((_132_v5).get(_125_v0)) : (new BigNumber((_182_v26).length)))).minus((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))])).isLessThan(p0)) {
          let _183_v27;
          _183_v27 = _dafny.Seq.of(_127_v2, _127_v2, _127_v2, _127_v2, _127_v2);
          let _184_v28;
          let _nw6 = new _module.C2();
          _nw6.__ctor(_125_v0, (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]);
          _184_v28 = _nw6;
          let _185_v29;
          _185_v29 = _dafny.MultiSet.fromElements(_184_v28);
          let _186_v30;
          _186_v30 = _dafny.Map.Empty.slice().updateUnsafe((_183_v27)[_module.__default.safeIndex((_dafny.ZERO).minus((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]), new BigNumber((_183_v27).length))],(_185_v29).equals(_185_v29));
          _186_v30 = (_186_v30).Merge(_186_v30);
          let _187_v31;
          let _nw7 = new _module.C3();
          _nw7.__ctor(_125_v0, _125_v0);
          _187_v31 = _nw7;
          _187_v31 = _187_v31;
          let _index4 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
          (_127_v2)[_index4] = (_dafny.ZERO).minus(p0);
          let _nw8 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _127_v2 = _nw8;
          _125_v0 = (_187_v31).f10;
        } else {
          let _188_v32;
          let _nw9 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _188_v32 = _nw9;
          let _index5 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_188_v32).length));
          (_188_v32)[_index5] = _134_v7;
          let _index6 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_188_v32).length));
          (_188_v32)[_index6] = _dafny.Seq.Concat(_dafny.Seq.update(_134_v7, _module.__default.safeIndex(p0, new BigNumber((_134_v7).length)), _133_v6), _134_v7);
          let _189_v33;
          _189_v33 = _dafny.MultiSet.fromElements(new BigNumber(((_188_v32)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_188_v32).length))]).length), (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]);
          let _index7 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length));
          (p1)[_index7] = (_189_v33).IsDisjointFrom(_189_v33);
          let _index8 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length));
          let _index9 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
          let _rhs4 = p0;
          let _rhs5 = _dafny.Seq.IsProperPrefixOf(_134_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), ((_190_v6) => function (_191_i3) {
            return _190_v6;
          })(_133_v6)));
          let _rhs6 = p0;
          let _lhs2 = globalState;
          let _lhs3 = p1;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length));
          let _lhs5 = _127_v2;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
          _lhs2.f7 = _rhs4;
          _lhs3[_lhs4] = _rhs5;
          _lhs5[_lhs6] = _rhs6;
          let _192_v34;
          let _nw10 = Array((new BigNumber(27)).toNumber()).fill([]);
          _192_v34 = _nw10;
          let _index10 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_192_v34).length));
          (_192_v34)[_index10] = _127_v2;
          let _index11 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_192_v34).length));
          let _nw11 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          (_192_v34)[_index11] = _nw11;
          let _193_v35;
          _193_v35 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.of((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))]), _130_v3));
          let _194_v36;
          _194_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_125_v0);
          _193_v35 = _module.__default.fm19((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))], ((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]).isLessThan((_dafny.ZERO).minus(new BigNumber((_194_v36).length))), (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))], (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))], globalState);
          let _195_v37;
          _195_v37 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))],(p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))]);
          let _196_v38;
          let _nw12 = Array((new BigNumber(21)).toNumber());
          _nw12[(_dafny.ZERO).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(_dafny.ONE).toNumber()] = false;
          _nw12[(new BigNumber(2)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(3)).toNumber()] = _module.__default.fm0((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))], globalState);
          _nw12[(new BigNumber(4)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(5)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(6)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(7)).toNumber()] = (((_195_v37).contains((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))])) ? ((_195_v37).get((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))])) : ((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))]));
          _nw12[(new BigNumber(8)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(9)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(10)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(11)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(12)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(13)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(14)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(15)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))];
          _nw12[(new BigNumber(16)).toNumber()] = false;
          _nw12[(new BigNumber(17)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(18)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(19)).toNumber()] = _125_v0;
          _nw12[(new BigNumber(20)).toNumber()] = !((p1)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p1).length))]);
          _196_v38 = _nw12;
          let _197_v39;
          let _nw13 = Array((new BigNumber(2)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _196_v38;
          _nw13[(_dafny.ONE).toNumber()] = _196_v38;
          _197_v39 = _nw13;
          let _index12 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_197_v39).length));
          (_197_v39)[_index12] = _196_v38;
          let _index13 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_197_v39).length));
          (_197_v39)[_index13] = _196_v38;
        }
        let _198_v40;
        let _nw14 = new _module.C8();
        _nw14.__ctor(_module.__default.fm4(_125_v0, true, _125_v0, globalState));
        _198_v40 = _nw14;
        let _199_v41;
        _199_v41 = _198_v40;
        let _200_v42;
        let _nw15 = Array((new BigNumber(2)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = ((true) ? ((_199_v41)) : (_198_v40));
        _nw15[(_dafny.ONE).toNumber()] = _198_v40;
        _200_v42 = _nw15;
        _200_v42 = _200_v42;
        _125_v0 = _125_v0;
        let _201_v43;
        _201_v43 = _dafny.MultiSet.fromElements(new BigNumber((_131_v4).length));
        let _202_v44;
        _202_v44 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_201_v43);
        (globalState).f1 = ((_dafny.ZERO).minus(new BigNumber(((((_202_v44).contains(_125_v0)) ? ((_202_v44).get(_125_v0)) : (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,true)).length), p0, p0, new BigNumber((_dafny.MultiSet.FromArray(_131_v4)).cardinality()), p0)))).cardinality()))).multipliedBy(p0);
      }
      let _hi0 = new BigNumber((_130_v3).length);
      for (let _203_i4 = (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]; _203_i4.isLessThan(_hi0); _203_i4 = _203_i4.plus(_dafny.ONE)) {
        let _204_v45;
        _204_v45 = _dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), _133_v6, _133_v6);
        let _index14 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
        (_127_v2)[_index14] = (((_204_v45).IsSubsetOf(function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_dafny.Seq.update(_134_v7, _module.__default.safeIndex(p0, new BigNumber((_134_v7).length)), _133_v6)).Elements) {
            let _205_v46 = _compr_11;
            if (_dafny.Seq.contains(_dafny.Seq.update(_134_v7, _module.__default.safeIndex(p0, new BigNumber((_134_v7).length)), _133_v6), _205_v46)) {
              _coll11.add(_205_v46);
            }
          }
          return _coll11;
        }())) ? ((_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))]) : (_203_i4));
        let _206_v47;
        let _nw16 = Array((new BigNumber(6)).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = _134_v7;
        _nw16[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_134_v7, _134_v7);
        _nw16[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-211)), ((_207_v6) => function (_208_i5) {
          return _207_v6;
        })(_133_v6));
        _nw16[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_134_v7, _134_v7);
        _nw16[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_134_v7, _134_v7);
        _nw16[(new BigNumber(5)).toNumber()] = _134_v7;
        _206_v47 = _nw16;
        let _index15 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_206_v47).length));
        (_206_v47)[_index15] = _134_v7;
        let _index16 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_206_v47).length));
        (_206_v47)[_index16] = _dafny.Seq.UnicodeFromString("gwrlipv");
        if (_125_v0) {
          _134_v7 = _dafny.Seq.Concat((_206_v47)[_module.__default.safeIndex(new BigNumber(927), new BigNumber((_206_v47).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(657)), ((_209_v6) => function (_210_i6) {
            return _209_v6;
          })(_133_v6)));
          let _index17 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
          (_127_v2)[_index17] = new BigNumber((_dafny.Set.fromElements(!(_125_v0) || (_125_v0))).length);
          _125_v0 = _125_v0;
          let _211_v48;
          _211_v48 = _dafny.Map.Empty.slice().updateUnsafe(_125_v0,_125_v0);
          let _212_v49;
          _212_v49 = _dafny.Map.Empty.slice().updateUnsafe(_125_v0,_211_v48);
          _212_v49 = (_212_v49).update(!_dafny.areEqual(_130_v3, _dafny.Seq.of(false)), _211_v48);
          let _213_v50;
          let _nw17 = new _module.C9();
          _nw17.__ctor(_125_v0, _125_v0);
          _213_v50 = _nw17;
          let _214_v51;
          _214_v51 = _dafny.Map.Empty.slice().updateUnsafe(_133_v6,_213_v50);
          _214_v51 = ((_214_v51).update(_133_v6, _213_v50)).Merge((_214_v51).update(_133_v6, _213_v50));
        } else {
          _125_v0 = _module.__default.fm0(new BigNumber(825), globalState);
          (globalState).f1 = (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))];
          let _215_v52;
          let _nw18 = new _module.C10();
          _nw18.__ctor(_203_i4, _125_v0);
          _215_v52 = _nw18;
          _125_v0 = (_215_v52).f16;
          let _index18 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p1).length));
          (p1)[_index18] = _125_v0;
          let _index19 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p1).length));
          (p1)[_index19] = _module.__default.fm0(new BigNumber(185), globalState);
        }
        let _index20 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
        (_127_v2)[_index20] = (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))];
      }
      let _index21 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
      let _rhs7 = new BigNumber(-290);
      let _rhs8 = (_127_v2)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length))];
      let _rhs9 = _125_v0;
      let _lhs7 = _127_v2;
      let _lhs8 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_127_v2).length));
      let _lhs9 = globalState;
      _lhs7[_lhs8] = _rhs7;
      _lhs9.f7 = _rhs8;
      _125_v0 = _rhs9;
      let _216_v53;
      let _nw19 = new _module.C2();
      _nw19.__ctor(_125_v0, new BigNumber(936));
      _216_v53 = _nw19;
      let _217_v54;
      _217_v54 = _dafny.MultiSet.fromElements(_216_v53);
      (globalState).f1 = (new BigNumber(-876)).plus(new BigNumber((_217_v54).cardinality()));
      r0 = _134_v7;
      r1 = _134_v7;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _218_v0;
      _218_v0 = true;
      let _219_v1;
      _219_v1 = _dafny.Set.fromElements(_218_v0);
      let _220_v2;
      _220_v2 = new _dafny.CodePoint('c'.codePointAt(0));
      let _221_v3;
      _221_v3 = _dafny.MultiSet.fromElements(!(false));
      let _222_v4;
      _222_v4 = _dafny.Seq.UnicodeFromString("hxfbexr");
      let _223_globalState;
      let _nw20 = new _module.GlobalState();
      _nw20.__ctor(new BigNumber(546), new BigNumber(521), true, _219_v1, _dafny.Set.fromElements(_220_v2), _221_v3, _dafny.Seq.Concat(_222_v4, _222_v4), new BigNumber(359));
      _223_globalState = _nw20;
      _218_v0 = _218_v0;
      let _224_v5;
      let _init1 = ((_225_v0) => function (_226_i1) {
        return _225_v0;
      })(_218_v0);
      let _nw21 = Array((new BigNumber(25)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw21.length); _i0_1++) {
        _nw21[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _224_v5 = _nw21;
      let _227_v6;
      _227_v6 = _dafny.Seq.of(_224_v5);
      let _228_i0;
      _228_i0 = _dafny.ZERO;
      L1: {
        while (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(_224_v5), _227_v6), _dafny.Seq.Concat(_227_v6, _dafny.Seq.of(_224_v5, _224_v5)))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_228_i0)) {
              break L1;
            }
            _228_i0 = (_228_i0).plus(_dafny.ONE);
            _220_v2 = _220_v2;
            let _229_v7;
            _229_v7 = new BigNumber(52);
            let _230_v8;
            let _231_v9;
            let _out0;
            let _out1;
            let _outcollector0 = _module.__default.m0((_dafny.ZERO).minus(_229_v7), _224_v5, _223_globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _230_v8 = _out0;
            _231_v9 = _out1;
            let _232_v10;
            _232_v10 = _dafny.Seq.of(_218_v0, _218_v0);
            let _233_v11;
            _233_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_232_v10, _232_v10),_229_v7);
            _233_v11 = (_233_v11).update(_218_v0, new BigNumber((_230_v8).length));
            let _234_v12;
            _234_v12 = _dafny.Map.Empty.slice().updateUnsafe(_220_v2,_218_v0);
            _234_v12 = (_234_v12).update(_220_v2, _218_v0);
          }
        }
      }
      let _235_v13;
      _235_v13 = _dafny.Map.Empty.slice().updateUnsafe(_218_v0,_218_v0);
      let _236_v14;
      _236_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_235_v13).length),(_218_v0) || (_218_v0));
      let _index22 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      (_224_v5)[_index22] = _218_v0;
      let _index23 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      let _rhs10 = _236_v14;
      let _rhs11 = _dafny.Seq.IsPrefixOf(_222_v4, _222_v4);
      let _lhs10 = _224_v5;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      _236_v14 = _rhs10;
      _lhs10[_lhs11] = _rhs11;
      let _237_v15;
      _237_v15 = _dafny.Seq.of((new BigNumber((_dafny.Seq.UnicodeFromString("nsrrqfxka")).length)).minus(new BigNumber(435)));
      (_223_globalState).f1 = (_dafny.ZERO).minus(new BigNumber((_237_v15).length));
      let _index24 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      (_224_v5)[_index24] = _218_v0;
      let _238_v16;
      _238_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_222_v4).length),new BigNumber(901));
      let _239_v17;
      _239_v17 = new BigNumber(470);
      let _240_v18;
      _240_v18 = _module.D0.create_DC2((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_227_v6)[_module.__default.safeIndex((((_238_v16).contains(new BigNumber(803))) ? ((_238_v16).get(new BigNumber(803))) : (_239_v17)), new BigNumber((_227_v6).length))]);
      let _source11 = _240_v18;
      if (_source11.is_DC0) {
        let _241___mcc_h0 = (_source11).cf0;
        let _242_cf0 = _241___mcc_h0;
        let _243_v19;
        _243_v19 = _dafny.Seq.of(_218_v0, _218_v0);
        let _244_v20;
        let _init2 = ((_245_v17) => function (_246_i2) {
          return (_246_i2).plus(_245_v17);
        })(_239_v17);
        let _nw22 = Array((new BigNumber(23)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw22.length); _i0_2++) {
          _nw22[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _244_v20 = _nw22;
        let _247_v21;
        _247_v21 = _module.D1.create_DC4(new BigNumber((_243_v19).length), _244_v20);
        let _index25 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        let _rhs12 = _239_v17;
        let _rhs13 = !(((_dafny.ZERO).minus(_239_v17)).isLessThan((_dafny.ZERO).minus(_239_v17)));
        let _rhs14 = (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))];
        let _rhs15 = !(_239_v17).isEqualTo((_247_v21).dtor_cf5);
        let _rhs16 = (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))];
        let _lhs12 = _223_globalState;
        let _lhs13 = _224_v5;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        _lhs12.f1 = _rhs12;
        _lhs13[_lhs14] = _rhs13;
        _218_v0 = _rhs14;
        _218_v0 = _rhs15;
        _218_v0 = _rhs16;
        let _index26 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        (_224_v5)[_index26] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_218_v0, !(false), (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_module.__default.fm0(_239_v17, _223_globalState)), _module.__default.safeIndex(_239_v17, new BigNumber((_dafny.Seq.of(_module.__default.fm0(_239_v17, _223_globalState))).length)), (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]), _dafny.Seq.of(true)));
        let _rhs17 = _module.__default.fm1(_223_globalState);
        let _rhs18 = (_240_v18).dtor_cf2;
        _239_v17 = _rhs17;
        _218_v0 = _rhs18;
        let _rhs19 = _module.__default.fm1(_223_globalState);
        let _rhs20 = _218_v0;
        let _rhs21 = _222_v4;
        _239_v17 = _rhs19;
        _218_v0 = _rhs20;
        _222_v4 = _rhs21;
      } else if (_source11.is_DC1) {
        let _248___mcc_h1 = (_source11).cf1;
        let _249_cf1 = _248___mcc_h1;
        let _250_v22;
        _250_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ncsyqp"),false);
        let _index27 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        (_224_v5)[_index27] = (((_250_v22).contains(_249_cf1)) ? ((_250_v22).get(_249_cf1)) : (!((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))])));
        let _251_v23;
        _251_v23 = _dafny.Seq.of(true);
        let _252_v24;
        _252_v24 = _module.D1.create_DC3(new BigNumber((_dafny.Set.fromElements(_240_v18)).length));
        (_223_globalState).f1 = (new BigNumber((_251_v23).length)).minus((_252_v24).dtor_cf4);
        let _253_v25;
        _253_v25 = _dafny.Map.Empty.slice().updateUnsafe(_224_v5,_239_v17);
        (_223_globalState).f7 = (((_253_v25).contains(_224_v5)) ? ((_253_v25).get(_224_v5)) : (_239_v17));
        let _254_v26;
        _254_v26 = _dafny.Map.Empty.slice().updateUnsafe((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))],new BigNumber((_221_v3).cardinality()));
        let _255_v27;
        _255_v27 = _dafny.Map.Empty.slice().updateUnsafe(_251_v23,_254_v26);
        let _256_v28;
        _256_v28 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_251_v23,_254_v26)).update(_251_v23, _254_v26));
        let _257_v29;
        let _nw23 = Array((new BigNumber(7)).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = ((false) ? (_255_v27) : (_dafny.Map.Empty.slice().updateUnsafe(_251_v23,_254_v26)));
        _nw23[(_dafny.ONE).toNumber()] = ((_218_v0) ? (_255_v27) : (_255_v27));
        _nw23[(new BigNumber(2)).toNumber()] = _255_v27;
        _nw23[(new BigNumber(3)).toNumber()] = _module.__default.fm2(_239_v17, new _dafny.CodePoint('o'.codePointAt(0)), _239_v17, _239_v17, _223_globalState);
        _nw23[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_251_v23,_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_239_v17, _223_globalState),_239_v17));
        _nw23[(new BigNumber(5)).toNumber()] = (_256_v28)[_module.__default.safeIndex(_239_v17, new BigNumber((_256_v28).length))];
        _nw23[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_251_v23,_254_v26);
        _257_v29 = _nw23;
        let _index28 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_257_v29).length));
        (_257_v29)[_index28] = _255_v27;
        let _index29 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_257_v29).length));
        (_257_v29)[_index29] = _255_v27;
      } else {
        let _258___mcc_h2 = (_source11).cf2;
        let _259___mcc_h3 = (_source11).cf3;
        let _260_cf3 = _259___mcc_h3;
        let _261_cf2 = _258___mcc_h2;
        let _262_v30;
        let _263_v31;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0(_239_v17, _224_v5, _223_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _262_v30 = _out2;
        _263_v31 = _out3;
        _239_v17 = new BigNumber(544);
        if (true) {
          let _264_v32;
          _264_v32 = _module.D0.create_DC0(_224_v5);
          let _index30 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          let _rhs22 = ((_module.__default.fm0(_239_v17, _223_globalState)) ? ((_240_v18).dtor_cf2) : (false));
          let _rhs23 = _module.__default.fm1(_223_globalState);
          let _rhs24 = _261_cf2;
          let _rhs25 = _264_v32;
          let _lhs15 = _224_v5;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          let _lhs17 = _223_globalState;
          _lhs15[_lhs16] = _rhs22;
          _lhs17.f1 = _rhs23;
          _261_cf2 = _rhs24;
          _264_v32 = _rhs25;
          let _265_v33;
          let _266_v34;
          let _out4;
          let _out5;
          let _outcollector2 = _module.__default.m0(_module.__default.fm1(_223_globalState), _224_v5, _223_globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _265_v33 = _out4;
          _266_v34 = _out5;
          let _267_v35;
          let _268_v36;
          let _out6;
          let _out7;
          let _outcollector3 = _module.__default.m0((new BigNumber((_263_v31).length)).minus(_239_v17), _260_cf3, _223_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _267_v35 = _out6;
          _268_v36 = _out7;
          let _index31 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          (_224_v5)[_index31] = (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))];
          (_223_globalState).f0 = (new BigNumber((_267_v35).length)).minus(_239_v17);
        } else {
          _218_v0 = (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))];
          let _269_v37;
          _269_v37 = _dafny.Seq.of(false);
          let _270_v38;
          let _271_v39;
          let _out8;
          let _out9;
          let _outcollector4 = _module.__default.m0(new BigNumber((_269_v37).length), _224_v5, _223_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _270_v38 = _out8;
          _271_v39 = _out9;
          let _272_v40;
          _272_v40 = _module.D0.create_DC0(_260_cf3);
          _272_v40 = _module.D0.create_DC0(_260_cf3);
          let _273_v41;
          let _274_v42;
          let _out10;
          let _out11;
          let _outcollector5 = _module.__default.m0(_239_v17, _224_v5, _223_globalState);
          _out10 = _outcollector5[0];
          _out11 = _outcollector5[1];
          _273_v41 = _out10;
          _274_v42 = _out11;
          let _275_v43;
          let _nw24 = Array((new BigNumber(10)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(917)), ((_276_v2) => function (_277_i3) {
            return _276_v2;
          })(_220_v2))).length);
          _nw24[(_dafny.ONE).toNumber()] = _239_v17;
          _nw24[(new BigNumber(2)).toNumber()] = _239_v17;
          _nw24[(new BigNumber(3)).toNumber()] = _239_v17;
          _nw24[(new BigNumber(4)).toNumber()] = _239_v17;
          _nw24[(new BigNumber(5)).toNumber()] = _239_v17;
          _nw24[(new BigNumber(6)).toNumber()] = new BigNumber((_227_v6).length);
          _nw24[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_261_cf2)).cardinality());
          _nw24[(new BigNumber(8)).toNumber()] = new BigNumber((_238_v16).length);
          _nw24[(new BigNumber(9)).toNumber()] = _239_v17;
          _275_v43 = _nw24;
          (_223_globalState).f1 = (_module.D1.create_DC4(new BigNumber(620), _275_v43)).dtor_cf5;
        }
        let _278_v44;
        _278_v44 = _module.D0.create_DC1(_dafny.Seq.Concat(_262_v30, _263_v31));
        let _source12 = _278_v44;
        if (_source12.is_DC0) {
          let _279___mcc_h4 = (_source12).cf0;
          let _280_cf0 = _279___mcc_h4;
          let _281_v45;
          let _282_v46;
          let _out12;
          let _out13;
          let _outcollector6 = _module.__default.m0(_module.__default.safeModuloInt((_dafny.ZERO).minus(_239_v17), _239_v17), _260_cf3, _223_globalState);
          _out12 = _outcollector6[0];
          _out13 = _outcollector6[1];
          _281_v45 = _out12;
          _282_v46 = _out13;
          _261_cf2 = !(_module.__default.fm0(_239_v17, _223_globalState));
          let _283_v47;
          let _nw25 = Array((new BigNumber(7)).toNumber()).fill([]);
          _283_v47 = _nw25;
          let _index32 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_283_v47).length));
          (_283_v47)[_index32] = _224_v5;
          let _284_v48;
          _284_v48 = _dafny.MultiSet.fromElements(_239_v17, _239_v17, _239_v17);
          let _285_v49;
          _285_v49 = _dafny.Set.fromElements(_281_v45, _dafny.Seq.Create(_module.__default.abs(new BigNumber(476)), ((_286_v2) => function (_287_i4) {
            return _286_v2;
          })(_220_v2)), _dafny.Seq.UnicodeFromString("ig"));
          let _index33 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          let _index34 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_283_v47).length));
          let _rhs26 = ((_261_cf2) ? (_dafny.Seq.IsProperPrefixOf(_281_v45, _222_v4)) : ((_284_v48).IsDisjointFrom(_284_v48)));
          let _rhs27 = _222_v4;
          let _rhs28 = (!(_285_v49).contains(_263_v31)) || (((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) === (_261_cf2));
          let _rhs29 = _239_v17;
          let _rhs30 = _260_cf3;
          let _lhs18 = _224_v5;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          let _lhs20 = _223_globalState;
          let _lhs21 = _283_v47;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_283_v47).length));
          _261_cf2 = _rhs26;
          _263_v31 = _rhs27;
          _lhs18[_lhs19] = _rhs28;
          _lhs20.f7 = _rhs29;
          _lhs21[_lhs22] = _rhs30;
          let _288_v50;
          let _289_v51;
          let _out14;
          let _out15;
          let _outcollector7 = _module.__default.m0(_239_v17, (_227_v6)[_module.__default.safeIndex(_239_v17, new BigNumber((_227_v6).length))], _223_globalState);
          _out14 = _outcollector7[0];
          _out15 = _outcollector7[1];
          _288_v50 = _out14;
          _289_v51 = _out15;
        } else if (_source12.is_DC1) {
          let _290___mcc_h5 = (_source12).cf1;
          let _291_cf1 = _290___mcc_h5;
          let _292_v52;
          let _nw26 = Array((new BigNumber(5)).toNumber()).fill([]);
          _292_v52 = _nw26;
          let _index35 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_292_v52).length));
          (_292_v52)[_index35] = _260_cf3;
          let _293_v53;
          _293_v53 = _dafny.Map.Empty.slice().updateUnsafe(false,_222_v4);
          let _294_v54;
          _294_v54 = _dafny.Seq.of(_218_v0);
          let _295_v55;
          _295_v55 = _dafny.MultiSet.fromElements(new BigNumber(752), _239_v17, new BigNumber(-36), _239_v17, new BigNumber(2));
          let _296_v56;
          _296_v56 = _module.D0.create_DC0(_260_cf3);
          let _297_v57;
          _297_v57 = _dafny.MultiSet.fromElements(_module.D0.create_DC0(_260_cf3), _296_v56, _296_v56, _module.D0.create_DC0(_224_v5));
          let _298_v58;
          _298_v58 = _dafny.Map.Empty.slice().updateUnsafe(_239_v17,_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), ((_299_v17) => function (_300_i5) {
            return _299_v17;
          })(_239_v17)));
          let _index36 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_292_v52).length));
          let _rhs31 = (((_293_v53).contains(!((((_235_v13).contains(!(_261_cf2))) ? ((_235_v13).get(!(_261_cf2))) : ((_294_v54)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_294_v54).length))]))))) ? ((_293_v53).get(!((((_235_v13).contains(!(_261_cf2))) ? ((_235_v13).get(!(_261_cf2))) : ((_294_v54)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_294_v54).length))]))))) : (_291_cf1));
          let _rhs32 = !((_295_v55).IsProperSubsetOf(_295_v55));
          let _rhs33 = (((_dafny.MultiSet.fromElements(_296_v56)).IsSubsetOf(_297_v57)) ? (_dafny.Seq.Concat(_module.__default.fm3(_239_v17, _239_v17, _223_globalState), (((_298_v58).contains((((_221_v3).contains(_218_v0)) ? ((_221_v3).get(_218_v0)) : (_239_v17)))) ? ((_298_v58).get((((_221_v3).contains(_218_v0)) ? ((_221_v3).get(_218_v0)) : (_239_v17)))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(75)), ((_301_v17) => function (_302_i6) {
            return _301_v17;
          })(_239_v17)))))) : (_237_v15));
          let _rhs34 = _260_cf3;
          let _lhs23 = _292_v52;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_292_v52).length));
          _262_v30 = _rhs31;
          _218_v0 = _rhs32;
          _237_v15 = _rhs33;
          _lhs23[_lhs24] = _rhs34;
          let _303_v59;
          let _304_v60;
          let _out16;
          let _out17;
          let _outcollector8 = _module.__default.m0(new BigNumber(340), _224_v5, _223_globalState);
          _out16 = _outcollector8[0];
          _out17 = _outcollector8[1];
          _303_v59 = _out16;
          _304_v60 = _out17;
          let _305_v61;
          let _nw27 = Array((new BigNumber(20)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _220_v2;
          _nw27[(_dafny.ONE).toNumber()] = _220_v2;
          _nw27[(new BigNumber(2)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(3)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(4)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(5)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(6)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(7)).toNumber()] = (_module.D2.create_DC6(_module.__default.fm4(_218_v0, true, (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _223_globalState))).dtor_cf8;
          _nw27[(new BigNumber(8)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(9)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(10)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(11)).toNumber()] = _module.__default.fm4(_261_cf2, _218_v0, (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _223_globalState);
          _nw27[(new BigNumber(12)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(13)).toNumber()] = ((_218_v0) ? (_220_v2) : (new _dafny.CodePoint('y'.codePointAt(0))));
          _nw27[(new BigNumber(14)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(15)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(16)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(17)).toNumber()] = ((!(_261_cf2)) ? (_220_v2) : (_220_v2));
          _nw27[(new BigNumber(18)).toNumber()] = _220_v2;
          _nw27[(new BigNumber(19)).toNumber()] = _220_v2;
          _305_v61 = _nw27;
          let _index37 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_305_v61).length));
          (_305_v61)[_index37] = _220_v2;
          let _index38 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_305_v61).length));
          (_305_v61)[_index38] = _module.__default.fm4((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], !(_module.__default.fm0((_dafny.ZERO).minus(new BigNumber(-842)), _223_globalState)), (_dafny.MultiSet.FromArray(_294_v54)).IsSubsetOf((_221_v3).update(_218_v0, _module.__default.abs(_239_v17))), _223_globalState);
          (_223_globalState).f3 = ((!((_module.D0.create_DC2(_261_cf2, _224_v5)).dtor_cf2)) ? ((_219_v1).Union(_219_v1)) : (_219_v1));
        } else {
          let _306___mcc_h6 = (_source12).cf2;
          let _307___mcc_h7 = (_source12).cf3;
          let _308_cf3 = _307___mcc_h7;
          let _309_cf2 = _306___mcc_h6;
          let _310_v62;
          _310_v62 = _module.D2.create_DC6(_220_v2);
          let _pat_let_tv4 = _220_v2;
          let _311_v63;
          _311_v63 = _dafny.MultiSet.fromElements((_262_v30)[_module.__default.safeIndex(_239_v17, new BigNumber((_262_v30).length))], (function (_pat_let0_0) {
            return function (_312_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_313_dt__update_hcf8_h0) {
                  return _module.D2.create_DC6(_313_dt__update_hcf8_h0);
                }(_pat_let1_0);
              }(_pat_let_tv4);
            }(_pat_let0_0);
          }(_310_v62)).dtor_cf8);
          let _index39 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          (_224_v5)[_index39] = !(new BigNumber(925)).isEqualTo((((_311_v63).contains(_220_v2)) ? ((_311_v63).get(_220_v2)) : (_module.__default.fm1(_223_globalState))));
          (_223_globalState).f1 = _module.__default.fm1(_223_globalState);
          let _314_v64;
          _314_v64 = _dafny.Seq.of((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]);
          let _315_v65;
          _315_v65 = _dafny.Set.fromElements(_314_v64);
          let _index40 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          (_224_v5)[_index40] = (_315_v65).IsProperSubsetOf((_315_v65).Union(_dafny.Set.fromElements(_314_v64)));
          let _316_v66;
          let _nw28 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _316_v66 = _nw28;
          let _317_v67;
          _317_v67 = _module.D1.create_DC4(_239_v17, _316_v66);
          let _index41 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_316_v66).length));
          (_316_v66)[_index41] = (_239_v17).multipliedBy((_317_v67).dtor_cf5);
          let _index42 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_316_v66).length));
          let _rhs35 = new BigNumber(984);
          let _rhs36 = _239_v17;
          let _lhs25 = _316_v66;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_316_v66).length));
          let _lhs27 = _223_globalState;
          _lhs25[_lhs26] = _rhs35;
          _lhs27.f0 = _rhs36;
        }
      }
      let _318_v68;
      _318_v68 = _dafny.Seq.of(_219_v1);
      let _319_v69;
      _319_v69 = _dafny.Set.fromElements((_318_v68)[_module.__default.safeIndex(_239_v17, new BigNumber((_318_v68).length))]);
      let _320_v70;
      _320_v70 = _module.D3.create_DC9(_319_v69);
      let _index43 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      (_224_v5)[_index43] = !(((_320_v70).dtor_cf13).IsSubsetOf(_319_v69));
      let _hi1 = _239_v17;
      for (let _321_i7 = _239_v17; _321_i7.isLessThan(_hi1); _321_i7 = _321_i7.plus(_dafny.ONE)) {
        (_223_globalState).f7 = _239_v17;
        let _322_v71;
        _322_v71 = _dafny.Seq.of(!(_218_v0));
        let _323_v72;
        _323_v72 = _dafny.MultiSet.fromElements(_239_v17);
        let _324_v73;
        _324_v73 = _dafny.Map.Empty.slice().updateUnsafe(_323_v72,new BigNumber(664));
        let _325_v74;
        let _nw29 = Array((new BigNumber(19)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = _321_i7;
        _nw29[(_dafny.ONE).toNumber()] = _321_i7;
        _nw29[(new BigNumber(2)).toNumber()] = (_239_v17).plus(_239_v17);
        _nw29[(new BigNumber(3)).toNumber()] = new BigNumber(368);
        _nw29[(new BigNumber(4)).toNumber()] = _321_i7;
        _nw29[(new BigNumber(5)).toNumber()] = new BigNumber((_322_v71).length);
        _nw29[(new BigNumber(6)).toNumber()] = (_module.__default.fm1(_223_globalState)).multipliedBy(new BigNumber(409));
        _nw29[(new BigNumber(7)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))],new BigNumber((_219_v1).length))).length)).minus(new BigNumber(344));
        _nw29[(new BigNumber(8)).toNumber()] = new BigNumber((_322_v71).length);
        _nw29[(new BigNumber(9)).toNumber()] = (new BigNumber((_323_v72).cardinality())).plus(_321_i7);
        _nw29[(new BigNumber(10)).toNumber()] = (((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) ? (new BigNumber((_322_v71).length)) : (_321_i7));
        _nw29[(new BigNumber(11)).toNumber()] = _239_v17;
        _nw29[(new BigNumber(12)).toNumber()] = _239_v17;
        _nw29[(new BigNumber(13)).toNumber()] = _321_i7;
        _nw29[(new BigNumber(14)).toNumber()] = _module.__default.fm1(_223_globalState);
        _nw29[(new BigNumber(15)).toNumber()] = new BigNumber((_324_v73).length);
        _nw29[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_239_v17, _239_v17);
        _nw29[(new BigNumber(17)).toNumber()] = _321_i7;
        _nw29[(new BigNumber(18)).toNumber()] = (_239_v17).multipliedBy(_321_i7);
        _325_v74 = _nw29;
        _325_v74 = _325_v74;
        (_223_globalState).f1 = _239_v17;
        let _index44 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        (_224_v5)[_index44] = _218_v0;
      }
      if ((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) {
        _218_v0 = !(_218_v0);
        let _326_v75;
        let _nw30 = new _module.C1();
        _nw30.__ctor((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]);
        _326_v75 = _nw30;
        let _327_v76;
        let _nw31 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _327_v76 = _nw31;
        let _index45 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_327_v76).length));
        (_327_v76)[_index45] = _239_v17;
        let _index46 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_327_v76).length));
        (_327_v76)[_index46] = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_239_v17, _239_v17), (new BigNumber(438)).minus((_326_v75).fm9(_222_v4, _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_218_v0,(_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))])), _223_globalState)));
        let _328_v77;
        let _nw32 = new _module.C5();
        _nw32.__ctor();
        _328_v77 = _nw32;
        let _index47 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        (_224_v5)[_index47] = !((_326_v75).fm6(new BigNumber(445), (_327_v76)[_module.__default.safeIndex(new BigNumber(895), new BigNumber((_327_v76).length))], (_dafny.ZERO).minus((_239_v17).minus((_327_v76)[_module.__default.safeIndex(new BigNumber(895), new BigNumber((_327_v76).length))])), (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _223_globalState));
      } else {
        let _index48 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        let _rhs37 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-10)), ((_329_v2) => function (_330_i8) {
          return _329_v2;
        })(_220_v2));
        let _rhs38 = _218_v0;
        let _lhs28 = _224_v5;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        _222_v4 = _rhs37;
        _lhs28[_lhs29] = _rhs38;
        (_223_globalState).f7 = (new BigNumber((_222_v4).length)).minus((new BigNumber(372)).multipliedBy(_239_v17));
        let _331_v78;
        _331_v78 = _dafny.Map.Empty.slice().updateUnsafe((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))],_222_v4);
        let _332_v79;
        _332_v79 = _dafny.Seq.of(_218_v0, false);
        _331_v78 = (_331_v78).update((_dafny.MultiSet.fromElements((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_module.__default.fm65(_218_v0, _332_v79, _239_v17, _223_globalState)).dtor_cf16)).contains((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]), _222_v4);
        let _333_v80;
        let _nw33 = Array((new BigNumber(19)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _220_v2;
        _nw33[(_dafny.ONE).toNumber()] = _220_v2;
        _nw33[(new BigNumber(2)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(3)).toNumber()] = _module.__default.fm4((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _218_v0, _218_v0, _223_globalState);
        _nw33[(new BigNumber(4)).toNumber()] = ((!((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))])) ? (_220_v2) : (new _dafny.CodePoint('v'.codePointAt(0))));
        _nw33[(new BigNumber(5)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(6)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(7)).toNumber()] = (((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) ? (_module.__default.fm4(!(!(_218_v0)), (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _218_v0, _223_globalState)) : (_220_v2));
        _nw33[(new BigNumber(8)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw33[(new BigNumber(10)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(11)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(12)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(13)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
        _nw33[(new BigNumber(15)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(16)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(17)).toNumber()] = _220_v2;
        _nw33[(new BigNumber(18)).toNumber()] = _220_v2;
        _333_v80 = _nw33;
        let _334_v81;
        _334_v81 = _module.D12.create_DC37(_239_v17, _333_v80, _220_v2);
        let _335_v82;
        _335_v82 = _dafny.Map.Empty.slice().updateUnsafe(_239_v17,(_334_v81).dtor_cf62);
        let _336_v83;
        _336_v83 = _dafny.MultiSet.fromElements(_239_v17, _module.__default.fm1(_223_globalState), new BigNumber(705), _239_v17, _239_v17);
        let _rhs39 = (((_335_v82).contains(new BigNumber((_336_v83).cardinality()))) ? ((_335_v82).get(new BigNumber((_336_v83).cardinality()))) : ((((_335_v82).contains(_239_v17)) ? ((_335_v82).get(_239_v17)) : (_333_v80))));
        let _rhs40 = _220_v2;
        let _rhs41 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_239_v17).minus(_239_v17)));
        let _lhs30 = _223_globalState;
        _333_v80 = _rhs39;
        _220_v2 = _rhs40;
        _lhs30.f7 = _rhs41;
        (_223_globalState).f1 = _239_v17;
      }
      _218_v0 = _218_v0;
      let _337_v84;
      _337_v84 = _module.D13.create_DC39(_239_v17, true);
      let _pat_let_tv5 = _238_v16;
      let _pat_let_tv6 = _238_v16;
      let _pat_let_tv7 = _238_v16;
      let _source13 = function (_source14) {
        if (_source14.is_DC39) {
          let _338___mcc_h10 = (_source14).cf65;
          let _339___mcc_h11 = (_source14).cf66;
          let _340_cf66 = _339___mcc_h11;
          let _341_cf65 = _338___mcc_h10;
          return _module.D7.create_DC22(_pat_let_tv5);
        } else if (_source14.is_DC40) {
          let _342___mcc_h12 = (_source14).cf67;
          let _343___mcc_h13 = (_source14).cf68;
          let _344___mcc_h14 = (_source14).cf69;
          let _345___mcc_h15 = (_source14).cf70;
          let _346_cf70 = _345___mcc_h15;
          let _347_cf69 = _344___mcc_h14;
          let _348_cf68 = _343___mcc_h13;
          let _349_cf67 = _342___mcc_h12;
          return _module.D7.create_DC22(_pat_let_tv6);
        } else {
          let _350___mcc_h16 = (_source14).cf64;
          let _351_cf64 = _350___mcc_h16;
          return _module.D7.create_DC22(_pat_let_tv7);
        }
      }(_337_v84);
      if (_source13.is_DC23) {
        let _index49 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
        (_224_v5)[_index49] = (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))];
        if (_218_v0) {
          (_223_globalState).f7 = _239_v17;
          let _352_v85;
          let _nw34 = Array((new BigNumber(3)).toNumber()).fill(_module.D8.Default());
          _352_v85 = _nw34;
          let _353_v86;
          let _nw35 = Array((new BigNumber(25)).toNumber()).fill(_module.D1.Default());
          _353_v86 = _nw35;
          let _354_v87;
          _354_v87 = _dafny.Seq.of(_353_v86, _353_v86);
          let _355_v88;
          _355_v88 = _module.D8.create_DC27(true, _354_v87, _237_v15, _239_v17);
          let _index50 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_352_v85).length));
          (_352_v85)[_index50] = _355_v88;
          let _356_v89;
          let _nw36 = new _module.C7();
          _nw36.__ctor((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _218_v0);
          _356_v89 = _nw36;
          let _357_v90;
          _357_v90 = _dafny.Map.Empty.slice().updateUnsafe(_356_v89,_220_v2);
          let _index51 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_352_v85).length));
          let _rhs42 = _355_v88;
          let _rhs43 = (_dafny.Map.Empty.slice().updateUnsafe(_356_v89,_220_v2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_356_v89,_220_v2));
          let _rhs44 = (_dafny.ZERO).minus(_239_v17);
          let _lhs31 = _352_v85;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_352_v85).length));
          let _lhs33 = _223_globalState;
          _lhs31[_lhs32] = _rhs42;
          _357_v90 = _rhs43;
          _lhs33.f0 = _rhs44;
          (_223_globalState).f1 = _239_v17;
          let _358_v91;
          let _nw37 = new _module.C6();
          _nw37.__ctor();
          _358_v91 = _nw37;
          _358_v91 = _358_v91;
          let _359_v92;
          let _nw38 = Array((new BigNumber(27)).toNumber());
          _359_v92 = _nw38;
          let _360_v93;
          let _init3 = ((_361_v4) => function (_362_i9) {
            return _dafny.Seq.Concat(_361_v4, _361_v4);
          })(_222_v4);
          let _nw39 = Array((new BigNumber(25)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw39.length); _i0_3++) {
            _nw39[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _360_v93 = _nw39;
          let _index52 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_360_v93).length));
          (_360_v93)[_index52] = _dafny.Seq.UnicodeFromString("jqicmgfh");
          let _index53 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_360_v93).length));
          let _rhs45 = _359_v92;
          let _rhs46 = (_dafny.ZERO).minus((new BigNumber(-265)).plus(_239_v17));
          let _rhs47 = _dafny.Seq.Concat(_222_v4, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qsyncjph"), _222_v4));
          let _lhs34 = _223_globalState;
          let _lhs35 = _360_v93;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_360_v93).length));
          _359_v92 = _rhs45;
          _lhs34.f0 = _rhs46;
          _lhs35[_lhs36] = _rhs47;
        } else {
          let _363_v94;
          let _nw40 = new _module.C2();
          _nw40.__ctor((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_239_v17).minus(_239_v17));
          _363_v94 = _nw40;
          let _364_v95;
          let _365_v96;
          let _out18;
          let _out19;
          let _outcollector9 = _module.__default.m0(_239_v17, _224_v5, _223_globalState);
          _out18 = _outcollector9[0];
          _out19 = _outcollector9[1];
          _364_v95 = _out18;
          _365_v96 = _out19;
          _220_v2 = _220_v2;
          let _366_v97;
          let _nw41 = new _module.C8();
          _nw41.__ctor((_364_v95)[_module.__default.safeIndex((_363_v94).f22, new BigNumber((_364_v95).length))]);
          _366_v97 = _nw41;
          let _367_v98;
          let _init4 = ((_368_v17) => function (_369_i10) {
            return _module.__default.safeModuloInt(_369_i10, _368_v17);
          })(_239_v17);
          let _nw42 = Array((new BigNumber(6)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw42.length); _i0_4++) {
            _nw42[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _367_v98 = _nw42;
          let _370_v99;
          _370_v99 = _dafny.Seq.of(_367_v98);
          let _371_v100;
          _371_v100 = _module.D24.create_DC68(_dafny.Seq.update(_370_v99, _module.__default.safeIndex(new BigNumber((_236_v14).length), new BigNumber((_370_v99).length)), _367_v98));
          let _372_v101;
          _372_v101 = _dafny.Map.Empty.slice().updateUnsafe((_371_v100).dtor_cf117,(((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) ? ((_363_v94).f21) : ((_363_v94).f21)));
          let _373_v102;
          _373_v102 = _dafny.Seq.of(_370_v99);
          let _374_v103;
          _374_v103 = _module.D4.create_DC13(_module.__default.fm66((_363_v94).f21, _220_v2, _239_v17, (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _223_globalState));
          let _375_v104;
          _375_v104 = _dafny.Seq.of(_374_v103);
          _372_v101 = (_372_v101).update((_373_v102)[_module.__default.safeIndex(_239_v17, new BigNumber((_373_v102).length))], (_366_v97).fm17((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _239_v17, _375_v104, (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _223_globalState));
        }
        if (!((_239_v17).isLessThan(_239_v17))) {
          let _376_v105;
          let _377_v106;
          let _out20;
          let _out21;
          let _outcollector10 = _module.__default.m0((new BigNumber(-56)).multipliedBy(new BigNumber(445)), _224_v5, _223_globalState);
          _out20 = _outcollector10[0];
          _out21 = _outcollector10[1];
          _376_v105 = _out20;
          _377_v106 = _out21;
          let _378_v107;
          _378_v107 = _dafny.Map.Empty.slice().updateUnsafe(_224_v5,_239_v17);
          let _379_v108;
          let _380_v109;
          let _out22;
          let _out23;
          let _outcollector11 = _module.__default.m0((((_378_v107).contains(_224_v5)) ? ((_378_v107).get(_224_v5)) : (new BigNumber(425))), _224_v5, _223_globalState);
          _out22 = _outcollector11[0];
          _out23 = _outcollector11[1];
          _379_v108 = _out22;
          _380_v109 = _out23;
          _224_v5 = ((_218_v0) ? (((_218_v0) ? (_224_v5) : (_224_v5))) : (_224_v5));
          let _381_v110;
          let _382_v111;
          let _out24;
          let _out25;
          let _outcollector12 = _module.__default.m0(_239_v17, _224_v5, _223_globalState);
          _out24 = _outcollector12[0];
          _out25 = _outcollector12[1];
          _381_v110 = _out24;
          _382_v111 = _out25;
          _218_v0 = _module.__default.fm0(_239_v17, _223_globalState);
        } else {
          let _383_v112;
          let _384_v113;
          let _out26;
          let _out27;
          let _outcollector13 = _module.__default.m0(_239_v17, (_module.D0.create_DC2((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _224_v5)).dtor_cf3, _223_globalState);
          _out26 = _outcollector13[0];
          _out27 = _outcollector13[1];
          _383_v112 = _out26;
          _384_v113 = _out27;
          let _385_v114;
          let _nw43 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _385_v114 = _nw43;
          let _index54 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_385_v114).length));
          (_385_v114)[_index54] = _239_v17;
          let _index55 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_385_v114).length));
          (_385_v114)[_index55] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_219_v1).length)));
          let _386_v115;
          _386_v115 = _module.D16.create_DC49(_239_v17);
          _386_v115 = _386_v115;
          let _387_v116;
          let _nw44 = new _module.C11();
          _nw44.__ctor((_385_v114)[_module.__default.safeIndex(new BigNumber(955), new BigNumber((_385_v114).length))], _383_v112, (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]);
          _387_v116 = _nw44;
          _387_v116 = _387_v116;
          let _388_v117;
          let _nw45 = new _module.C2();
          _nw45.__ctor((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _239_v17);
          _388_v117 = _nw45;
        }
        let _389_v118;
        let _nw46 = new _module.C11();
        _nw46.__ctor(_module.__default.fm1(_223_globalState), ((_218_v0) ? (_222_v4) : (_222_v4)), (_module.__default.fm67((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _239_v17, _239_v17, _223_globalState)).dtor_cf109);
        _389_v118 = _nw46;
      } else if (_source13.is_DC24) {
        let _390___mcc_h8 = (_source13).cf41;
        let _391_cf41 = _390___mcc_h8;
        if ((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) {
          let _392_v119;
          let _nw47 = new _module.C12();
          _nw47.__ctor(_218_v0, _218_v0, _218_v0);
          _392_v119 = _nw47;
          let _393_v120;
          _393_v120 = _dafny.Seq.of(_392_v119);
          let _index56 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
          (_224_v5)[_index56] = (new BigNumber((_393_v120).length)).isLessThanOrEqualTo(_239_v17);
          let _394_v121;
          let _nw48 = new _module.C0();
          _nw48.__ctor();
          _394_v121 = _nw48;
          _220_v2 = ((_392_v119.f12) ? (new _dafny.CodePoint('e'.codePointAt(0))) : (_220_v2));
          let _395_v122;
          _395_v122 = _dafny.Seq.of(_222_v4, _222_v4);
          let _396_v123;
          _396_v123 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), _220_v2),true);
          let _397_v124;
          let _nw49 = new _module.C1();
          _nw49.__ctor((((_396_v123).contains(_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0))))) ? ((_396_v123).get(_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0))))) : ((_394_v121).fm37(_222_v4, _391_cf41, (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], _392_v119.f12, _223_globalState))));
          _397_v124 = _nw49;
          let _398_v125;
          let _nw50 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _398_v125 = _nw50;
          let _index57 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_398_v125).length));
          (_398_v125)[_index57] = _391_cf41;
          let _399_v126;
          let _nw51 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _399_v126 = _nw51;
          let _400_v127;
          _400_v127 = _module.D12.create_DC37(new BigNumber(-799), _399_v126, _220_v2);
          let _401_v128;
          _401_v128 = _dafny.MultiSet.fromElements(_391_cf41, _239_v17);
          let _402_v131;
          _402_v131 = _module.D7.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(_391_cf41,_391_cf41));
          let _403_v132;
          _403_v132 = _dafny.Seq.of(_238_v16, (_238_v16).update(_239_v17, _391_cf41), function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of _dafny.IntegerRange(new BigNumber(547), new BigNumber(470))) {
              let _404_v129 = _compr_12;
              if (((new BigNumber(547)).isLessThanOrEqualTo(_404_v129)) && ((_404_v129).isLessThan(new BigNumber(470)))) {
                _coll12.push([_module.__default.safeModuloInt(_404_v129, _391_cf41),_391_cf41]);
              }
            }
            return _coll12;
          }(), function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of _dafny.IntegerRange(new BigNumber(975), new BigNumber(-989))) {
              let _405_v130 = _compr_13;
              if (((new BigNumber(975)).isLessThanOrEqualTo(_405_v130)) && ((_405_v130).isLessThan(new BigNumber(-989)))) {
                _coll13.push([(_405_v130).minus((_module.D4.create_DC15((((_401_v128).contains(new BigNumber((_236_v14).length))) ? ((_401_v128).get(new BigNumber((_236_v14).length))) : (new BigNumber(341))), _401_v128, _220_v2, _391_cf41)).dtor_cf23),new BigNumber((_dafny.Seq.UnicodeFromString("tolfmsbf")).length)]);
              }
            }
            return _coll13;
          }(), (_402_v131).dtor_cf40);
          let _406_v133;
          _406_v133 = _module.D13.create_DC40(_400_v127, true, (((_401_v128).contains(new BigNumber((_219_v1).length))) ? ((_401_v128).get(new BigNumber((_219_v1).length))) : (_239_v17)), new BigNumber(((_403_v132)[_module.__default.safeIndex(_391_cf41, new BigNumber((_403_v132).length))]).length));
          let _407_v134;
          _407_v134 = _dafny.Seq.of(_406_v133);
          let _408_v135;
          _408_v135 = _dafny.MultiSet.fromElements(_407_v134);
          let _index58 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_398_v125).length));
          let _rhs48 = _395_v122;
          let _rhs49 = _391_cf41;
          let _rhs50 = _397_v124;
          let _rhs51 = (((_408_v135).contains(_dafny.Seq.Concat(_407_v134, _407_v134))) ? ((_408_v135).get(_dafny.Seq.Concat(_407_v134, _407_v134))) : (_391_cf41));
          let _lhs37 = _223_globalState;
          let _lhs38 = _398_v125;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_398_v125).length));
          _395_v122 = _rhs48;
          _lhs37.f0 = _rhs49;
          _397_v124 = _rhs50;
          _lhs38[_lhs39] = _rhs51;
          let _409_v136;
          let _nw52 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _409_v136 = _nw52;
          _409_v136 = _409_v136;
        } else {
          let _410_v137;
          _410_v137 = _module.D11.create_DC32((_236_v14).Merge(_236_v14));
          _410_v137 = _410_v137;
          let _411_v138;
          _411_v138 = _module.D19.create_DC58(new BigNumber((_222_v4).length), (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]);
          let _412_v139;
          _412_v139 = _dafny.Map.Empty.slice().updateUnsafe(_411_v138,_dafny.Seq.Concat(_222_v4, _222_v4));
          let _pat_let_tv8 = _224_v5;
          let _pat_let_tv9 = _224_v5;
          _412_v139 = (_412_v139).update(function (_pat_let2_0) {
            return function (_413_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_414_dt__update_hcf100_h0) {
                  return _module.D19.create_DC58((_413_dt__update__tmp_h1).dtor_cf99, _414_dt__update_hcf100_h0);
                }(_pat_let3_0);
              }((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_pat_let_tv8).length))]);
            }(_pat_let2_0);
          }(_411_v138), _dafny.Seq.Concat(_222_v4, _222_v4));
          let _415_v140;
          _415_v140 = _module.D20.create_DC60(_239_v17);
          let _pat_let_tv10 = _391_cf41;
          _415_v140 = function (_pat_let4_0) {
            return function (_416_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_417_dt__update_hcf102_h0) {
                  return _module.D20.create_DC60(_417_dt__update_hcf102_h0);
                }(_pat_let5_0);
              }(_pat_let_tv10);
            }(_pat_let4_0);
          }(_module.D20.create_DC60((_237_v15)[_module.__default.safeIndex(new BigNumber((_219_v1).length), new BigNumber((_237_v15).length))]));
          let _418_v141;
          _418_v141 = _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("hrso"), _218_v0, _218_v0);
          let _419_v142;
          _419_v142 = _dafny.Map.Empty.slice().updateUnsafe((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))],(_dafny.ZERO).minus(_239_v17));
          let _420_v143;
          _420_v143 = _dafny.Set.fromElements(_220_v2);
          let _421_v144;
          _421_v144 = _dafny.Map.Empty.slice().updateUnsafe(_418_v141,(_module.__default.fm18(_239_v17, _419_v142, _218_v0, _223_globalState)).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_420_v143).length), _391_cf41, _239_v17)));
          let _422_v145;
          _422_v145 = _dafny.MultiSet.fromElements(_391_cf41, _391_cf41);
          _421_v144 = (_421_v144).update(_418_v141, _422_v145);
          let _423_v146;
          let _nw53 = Array((new BigNumber(9)).toNumber());
          _nw53[(_dafny.ZERO).toNumber()] = _220_v2;
          _nw53[(_dafny.ONE).toNumber()] = _220_v2;
          _nw53[(new BigNumber(2)).toNumber()] = _220_v2;
          _nw53[(new BigNumber(3)).toNumber()] = _220_v2;
          _nw53[(new BigNumber(4)).toNumber()] = _220_v2;
          _nw53[(new BigNumber(5)).toNumber()] = _220_v2;
          _nw53[(new BigNumber(6)).toNumber()] = _220_v2;
          _nw53[(new BigNumber(7)).toNumber()] = _220_v2;
          _nw53[(new BigNumber(8)).toNumber()] = (((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]) ? (_220_v2) : (_220_v2));
          _423_v146 = _nw53;
          let _index59 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_423_v146).length));
          (_423_v146)[_index59] = _220_v2;
          let _index60 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_423_v146).length));
          (_423_v146)[_index60] = _220_v2;
        }
        let _424_v147;
        _424_v147 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("mpdogu"), _222_v4);
        let _425_v148;
        _425_v148 = _dafny.Map.Empty.slice().updateUnsafe(_239_v17,new _dafny.CodePoint('m'.codePointAt(0)));
        let _rhs52 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_222_v4, _222_v4), (_424_v147)[_module.__default.safeIndex((_dafny.ZERO).minus(_239_v17), new BigNumber((_424_v147).length))])).length);
        let _rhs53 = (((_425_v148).contains(_391_cf41)) ? ((_425_v148).get(_391_cf41)) : ((((((_236_v14).contains(_391_cf41)) ? ((_236_v14).get(_391_cf41)) : ((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]))) ? (_220_v2) : (_220_v2))));
        let _lhs40 = _223_globalState;
        _lhs40.f1 = _rhs52;
        _220_v2 = _rhs53;
        _222_v4 = ((_218_v0) ? (_dafny.Seq.UnicodeFromString("lfu")) : (_dafny.Seq.UnicodeFromString("o")));
        let _426_v149;
        let _nw54 = Array((new BigNumber(18)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
        _nw54[(_dafny.ONE).toNumber()] = _220_v2;
        _nw54[(new BigNumber(2)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(3)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(4)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(5)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(6)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(7)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw54[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _nw54[(new BigNumber(10)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(11)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(12)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(13)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(14)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(15)).toNumber()] = _220_v2;
        _nw54[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw54[(new BigNumber(17)).toNumber()] = _220_v2;
        _426_v149 = _nw54;
        let _427_v150;
        _427_v150 = _module.D12.create_DC37(_239_v17, _426_v149, _220_v2);
        let _428_v151;
        let _nw55 = Array((new BigNumber(28)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = _220_v2;
        _nw55[(_dafny.ONE).toNumber()] = _220_v2;
        _nw55[(new BigNumber(2)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(3)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(4)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(5)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(6)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(7)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(8)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(9)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(10)).toNumber()] = (_222_v4)[_module.__default.safeIndex(_391_cf41, new BigNumber((_222_v4).length))];
        _nw55[(new BigNumber(11)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(12)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(13)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw55[(new BigNumber(15)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(16)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(17)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(18)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw55[(new BigNumber(20)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(21)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(22)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(23)).toNumber()] = (_427_v150).dtor_cf63;
        _nw55[(new BigNumber(24)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(25)).toNumber()] = ((!((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))])) ? (_220_v2) : (_220_v2));
        _nw55[(new BigNumber(26)).toNumber()] = _220_v2;
        _nw55[(new BigNumber(27)).toNumber()] = _220_v2;
        _428_v151 = _nw55;
        let _index61 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_426_v149).length));
        (_426_v149)[_index61] = _220_v2;
        let _index62 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_426_v149).length));
        (_426_v149)[_index62] = ((true) ? (_220_v2) : (new _dafny.CodePoint('k'.codePointAt(0))));
      } else {
        let _429___mcc_h9 = (_source13).cf40;
        let _430_cf40 = _429___mcc_h9;
        _218_v0 = (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))];
        (_223_globalState).f7 = _239_v17;
        (_223_globalState).f7 = _239_v17;
        let _431_v152;
        let _nw56 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _431_v152 = _nw56;
        let _432_v153;
        _432_v153 = _dafny.Map.Empty.slice().updateUnsafe(_218_v0,_220_v2);
        let _433_v154;
        _433_v154 = _dafny.Map.Empty.slice().updateUnsafe(_432_v153,(_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]);
        let _434_v155;
        _434_v155 = _433_v154;
        let _index63 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_431_v152).length));
        (_431_v152)[_index63] = _434_v155;
        let _435_v156;
        let _nw57 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _435_v156 = _nw57;
        let _436_v157;
        _436_v157 = _module.D2.create_DC6(_220_v2);
        let _437_v158;
        _437_v158 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_236_v14).length),(_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]));
        let _438_v159;
        _438_v159 = _dafny.Seq.of((_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))], (_224_v5)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length))]);
        let _439_v161;
        _439_v161 = _dafny.Map.Empty.slice().updateUnsafe(_239_v17,function () {
          let _coll14 = new _dafny.Set();
          for (const _compr_14 of (_dafny.MultiSet.fromElements(_236_v14, _236_v14, _236_v14, _236_v14)).Elements) {
            let _440_v160 = _compr_14;
            if ((_dafny.MultiSet.fromElements(_236_v14, _236_v14, _236_v14, _236_v14)).contains(_440_v160)) {
              _coll14.add(_440_v160);
            }
          }
          return _coll14;
        }());
        let _441_v162;
        _441_v162 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_239_v17,_218_v0));
        let _442_v164;
        _442_v164 = _dafny.Seq.of(_437_v158, _437_v158);
        let _443_v165;
        let _nw58 = Array((new BigNumber(13)).toNumber());
        _nw58[(_dafny.ZERO).toNumber()] = (_437_v158).Intersect(_437_v158);
        _nw58[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_236_v14, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_222_v4).length),false), _236_v14, _236_v14, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_238_v16).length),(_438_v159)[_module.__default.safeIndex(_239_v17, new BigNumber((_438_v159).length))]));
        _nw58[(new BigNumber(2)).toNumber()] = _437_v158;
        _nw58[(new BigNumber(3)).toNumber()] = _437_v158;
        _nw58[(new BigNumber(4)).toNumber()] = (((_439_v161).contains(_239_v17)) ? ((_439_v161).get(_239_v17)) : (_437_v158));
        _nw58[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_239_v17,_218_v0), _236_v14, _236_v14);
        _nw58[(new BigNumber(6)).toNumber()] = (function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of (_441_v162).Elements) {
            let _444_v163 = _compr_15;
            if (_dafny.Seq.contains(_441_v162, _444_v163)) {
              _coll15.add(_444_v163);
            }
          }
          return _coll15;
        }()).Intersect((_442_v164)[_module.__default.safeIndex(_239_v17, new BigNumber((_442_v164).length))]);
        _nw58[(new BigNumber(7)).toNumber()] = _437_v158;
        _nw58[(new BigNumber(8)).toNumber()] = _437_v158;
        _nw58[(new BigNumber(9)).toNumber()] = ((_218_v0) ? (_437_v158) : ((_437_v158)));
        _nw58[(new BigNumber(10)).toNumber()] = _437_v158;
        _nw58[(new BigNumber(11)).toNumber()] = (_437_v158).Difference(_437_v158);
        _nw58[(new BigNumber(12)).toNumber()] = (_dafny.Set.fromElements(_236_v14)).Difference(_437_v158);
        _443_v165 = _nw58;
        let _index64 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_443_v165).length));
        (_443_v165)[_index64] = _437_v158;
        let _index65 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_431_v152).length));
        let _index66 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_443_v165).length));
        let _rhs54 = _dafny.Seq.Concat(_222_v4, _module.__default.fm38(_218_v0, _239_v17, _223_globalState));
        let _rhs55 = _434_v155;
        let _rhs56 = _435_v156;
        let _rhs57 = _436_v157;
        let _rhs58 = ((_module.__default.fm0(_239_v17, _223_globalState)) ? (_437_v158) : (_437_v158));
        let _lhs41 = _431_v152;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_431_v152).length));
        let _lhs43 = _443_v165;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_443_v165).length));
        _222_v4 = _rhs54;
        _lhs41[_lhs42] = _rhs55;
        _435_v156 = _rhs56;
        _436_v157 = _rhs57;
        _lhs43[_lhs44] = _rhs58;
      }
      let _445_v166;
      let _nw59 = new _module.C6();
      _nw59.__ctor();
      _445_v166 = _nw59;
      let _446_v167;
      _446_v167 = _dafny.Map.Empty.slice().updateUnsafe(_239_v17,_445_v166);
      let _447_v168;
      _447_v168 = _dafny.Set.fromElements(new BigNumber((_237_v15).length));
      _446_v167 = (_446_v167).update(new BigNumber((_447_v168).length), _445_v166);
      let _index67 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      let _rhs59 = _239_v17;
      let _rhs60 = ((_239_v17).minus(_239_v17)).isLessThanOrEqualTo(_239_v17);
      let _lhs45 = _223_globalState;
      let _lhs46 = _224_v5;
      let _lhs47 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_224_v5).length));
      _lhs45.f7 = _rhs59;
      _lhs46[_lhs47] = _rhs60;
      let _448_v169;
      let _nw60 = Array((new BigNumber(20)).toNumber()).fill([]);
      _448_v169 = _nw60;
      let _449_v170;
      let _init5 = ((_450_v2) => function (_451_i11) {
        return _450_v2;
      })(_220_v2);
      let _nw61 = Array((_dafny.ONE).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw61.length); _i0_5++) {
        _nw61[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _449_v170 = _nw61;
      let _index68 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_448_v169).length));
      (_448_v169)[_index68] = _449_v170;
      let _index69 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_448_v169).length));
      (_448_v169)[_index69] = _449_v170;
      let _452_v171;
      _452_v171 = _module.D24.create_DC69(_445_v166, new BigNumber(-344), _239_v17);
      _445_v166 = (_452_v171).dtor_cf118;
      let _453_v172;
      let _nw62 = new _module.C11();
      _nw62.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("qwudkt")).length), _222_v4, !(_239_v17).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), function (_454_i12) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)));
      _453_v172 = _nw62;
      process.stdout.write(_dafny.toString(_218_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_220_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_221_v3).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_222_v4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_223_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_223_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_globalState.f3).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_223_globalState).f4).equals(_dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_223_globalState).f5).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_globalState).f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_223_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v5)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_227_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_237_v15, _dafny.Seq.of(new BigNumber(409), new BigNumber(2), new BigNumber(-766), _dafny.ZERO, new BigNumber(-840), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544), new BigNumber(544)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(7),new BigNumber(901)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_239_v17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_v18).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_v18).dtor_cf3)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_318_v68, _dafny.Seq.of(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_319_v69).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v70).dtor_cf13).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_337_v84).dtor_cf65));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_337_v84).dtor_cf66));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_446_v167).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_447_v168).equals(_dafny.Set.fromElements(new BigNumber(80)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_448_v169)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_449_v170)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_452_v171).dtor_cf119));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_452_v171).dtor_cf120));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_453_v172).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_453_v172).f14).toVerbatimString(false));
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
    static create_DC1(cf1) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC2(cf2, cf3) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + this.cf1.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf2 === other.cf2 && this.cf3 === other.cf3;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0([]);
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
    static create_DC4(cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf7) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, []);
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
    static create_DC7(cf9, cf10, cf11) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC8(cf12) {
      let $dt = new D2(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC10(cf14, cf15, cf16) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC11(cf17, cf18) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC9(cf13) {
      let $dt = new D3(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D3(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + this.cf14.toVerbatimString(true) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15 && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.Seq.UnicodeFromString(""), false, false);
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
    static create_DC14(cf21, cf22) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC15(cf23, cf24, cf25, cf26) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D4(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf20) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC17(cf28, cf29) {
      let $dt = new D5(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC16(cf27) {
      let $dt = new D5(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC18(cf30) {
      let $dt = new D5(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC20(cf32, cf33, cf34, cf35) {
      let $dt = new D6(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC21(cf36, cf37, cf38, cf39) {
      let $dt = new D6(1);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC19(cf31) {
      let $dt = new D6(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf32 === other.cf32 && this.cf33 === other.cf33 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf36 === other.cf36 && this.cf37 === other.cf37 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC20(false, false, false, _dafny.ZERO);
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
    static create_DC23() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC24(cf41) {
      let $dt = new D7(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC22(cf40) {
      let $dt = new D7(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC23";
      } else if (this.$tag === 1) {
        return "D7.DC24" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf40) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC23();
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
    static create_DC26() {
      let $dt = new D8(0);
      return $dt;
    }
    static create_DC27(cf43, cf44, cf45, cf46) {
      let $dt = new D8(1);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC25(cf42) {
      let $dt = new D8(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC28(cf47) {
      let $dt = new D8(3);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26";
      } else if (this.$tag === 1) {
        return "D8.DC27" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC28" + "(" + _dafny.toString(this.cf47) + ")";
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
        return other.$tag === 1 && this.cf43 === other.cf43 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf42 === other.cf42;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf47, other.cf47);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC26();
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
    static create_DC29(cf48) {
      let $dt = new D9(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC29" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf48, other.cf48);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31(cf50, cf51, cf52) {
      let $dt = new D10(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC30(cf49) {
      let $dt = new D10(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC31" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC31(false, false, _dafny.ZERO);
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
    static create_DC33(cf54, cf55) {
      let $dt = new D11(0);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC34(cf56, cf57, cf58) {
      let $dt = new D11(1);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC32(cf53) {
      let $dt = new D11(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC34" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC33(false, _dafny.ZERO);
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
    static create_DC36(cf60) {
      let $dt = new D12(0);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC37(cf61, cf62, cf63) {
      let $dt = new D12(1);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC35(cf59) {
      let $dt = new D12(2);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf61, other.cf61) && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC36(_dafny.MultiSet.Empty);
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
    static create_DC39(cf65, cf66) {
      let $dt = new D13(0);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC40(cf67, cf68, cf69, cf70) {
      let $dt = new D13(1);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC38(cf64) {
      let $dt = new D13(2);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC39" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC40" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67) && this.cf68 === other.cf68 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC39(_dafny.ZERO, false);
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
    static create_DC42(cf72, cf73, cf74) {
      let $dt = new D14(0);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC43(cf75, cf76, cf77, cf78) {
      let $dt = new D14(1);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC41(cf71) {
      let $dt = new D14(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC42" + "(" + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC43" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf72, other.cf72) && _dafny.areEqual(this.cf73, other.cf73) && this.cf74 === other.cf74;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC42(_module.D1.Default(), _dafny.Seq.of(), false);
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
    static create_DC45(cf80, cf81, cf82) {
      let $dt = new D15(0);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC46() {
      let $dt = new D15(1);
      return $dt;
    }
    static create_DC44(cf79) {
      let $dt = new D15(2);
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC47(cf83) {
      let $dt = new D15(3);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC45" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC46";
      } else if (this.$tag === 2) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC47" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf80, other.cf80) && this.cf81 === other.cf81 && this.cf82 === other.cf82;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC45(_dafny.ZERO, false, false);
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
    static create_DC49(cf85) {
      let $dt = new D16(0);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC48(cf84) {
      let $dt = new D16(1);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC49" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC48" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf84 === other.cf84;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC49(_dafny.ZERO);
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
    static create_DC51(cf87, cf88) {
      let $dt = new D17(0);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC50(cf86) {
      let $dt = new D17(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC52(cf89) {
      let $dt = new D17(2);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC51" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC50" + "(" + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC52" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf87, other.cf87) && this.cf88 === other.cf88;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf86 === other.cf86;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC51(new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC54(cf91, cf92, cf93, cf94) {
      let $dt = new D18(0);
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC55(cf95, cf96) {
      let $dt = new D18(1);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC56(cf97) {
      let $dt = new D18(2);
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC53(cf90) {
      let $dt = new D18(3);
      $dt.cf90 = cf90;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get is_DC56() { return this.$tag === 2; }
    get is_DC53() { return this.$tag === 3; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf90() { return this.cf90; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC54" + "(" + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC55" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC56" + "(" + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 3) {
        return "D18.DC53" + "(" + _dafny.toString(this.cf90) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf91 === other.cf91 && _dafny.areEqual(this.cf92, other.cf92) && this.cf93 === other.cf93 && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf95 === other.cf95 && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf97, other.cf97);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf90, other.cf90);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC54(false, _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC58(cf99, cf100) {
      let $dt = new D19(0);
      $dt.cf99 = cf99;
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC57(cf98) {
      let $dt = new D19(1);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC58" + "(" + _dafny.toString(this.cf99) + ", " + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC57" + "(" + _dafny.toString(this.cf98) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf99, other.cf99) && this.cf100 === other.cf100;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf98, other.cf98);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC58(_dafny.ZERO, false);
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
    static create_DC60(cf102) {
      let $dt = new D20(0);
      $dt.cf102 = cf102;
      return $dt;
    }
    static create_DC59(cf101) {
      let $dt = new D20(1);
      $dt.cf101 = cf101;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC59() { return this.$tag === 1; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf101() { return this.cf101; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC60" + "(" + _dafny.toString(this.cf102) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC59" + "(" + _dafny.toString(this.cf101) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf102, other.cf102);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf101, other.cf101);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC60(_dafny.ZERO);
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
    static create_DC62(cf104, cf105) {
      let $dt = new D21(0);
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC63(cf106, cf107, cf108, cf109) {
      let $dt = new D21(1);
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC61(cf103) {
      let $dt = new D21(2);
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC64(cf110) {
      let $dt = new D21(3);
      $dt.cf110 = cf110;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get is_DC64() { return this.$tag === 3; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf110() { return this.cf110; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC62" + "(" + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC63" + "(" + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC61" + "(" + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC64" + "(" + _dafny.toString(this.cf110) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf104, other.cf104) && this.cf105 === other.cf105;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf106, other.cf106) && this.cf107 === other.cf107 && this.cf108 === other.cf108 && this.cf109 === other.cf109;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf103 === other.cf103;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf110, other.cf110);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC62(_dafny.Map.Empty, false);
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
    static create_DC66(cf112, cf113, cf114, cf115) {
      let $dt = new D22(0);
      $dt.cf112 = cf112;
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      $dt.cf115 = cf115;
      return $dt;
    }
    static create_DC65(cf111) {
      let $dt = new D22(1);
      $dt.cf111 = cf111;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf111() { return this.cf111; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC66" + "(" + _dafny.toString(this.cf112) + ", " + _dafny.toString(this.cf113) + ", " + _dafny.toString(this.cf114) + ", " + this.cf115.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC65" + "(" + _dafny.toString(this.cf111) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf112, other.cf112) && this.cf113 === other.cf113 && this.cf114 === other.cf114 && _dafny.areEqual(this.cf115, other.cf115);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf111 === other.cf111;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC66(new _dafny.CodePoint('D'.codePointAt(0)), false, false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC67(cf116) {
      let $dt = new D23(0);
      $dt.cf116 = cf116;
      return $dt;
    }
    get is_DC67() { return this.$tag === 0; }
    get dtor_cf116() { return this.cf116; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC67" + "(" + _dafny.toString(this.cf116) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf116 === other.cf116;
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
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC69(cf118, cf119, cf120) {
      let $dt = new D24(0);
      $dt.cf118 = cf118;
      $dt.cf119 = cf119;
      $dt.cf120 = cf120;
      return $dt;
    }
    static create_DC68(cf117) {
      let $dt = new D24(1);
      $dt.cf117 = cf117;
      return $dt;
    }
    get is_DC69() { return this.$tag === 0; }
    get is_DC68() { return this.$tag === 1; }
    get dtor_cf118() { return this.cf118; }
    get dtor_cf119() { return this.cf119; }
    get dtor_cf120() { return this.cf120; }
    get dtor_cf117() { return this.cf117; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC69" + "(" + _dafny.toString(this.cf118) + ", " + _dafny.toString(this.cf119) + ", " + _dafny.toString(this.cf120) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC68" + "(" + _dafny.toString(this.cf117) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf118 === other.cf118 && _dafny.areEqual(this.cf119, other.cf119) && _dafny.areEqual(this.cf120, other.cf120);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf117, other.cf117);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC69(null, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC70(cf121) {
      let $dt = new D25(0);
      $dt.cf121 = cf121;
      return $dt;
    }
    get is_DC70() { return this.$tag === 0; }
    get dtor_cf121() { return this.cf121; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC70" + "(" + _dafny.toString(this.cf121) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf121, other.cf121);
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
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC71(cf122) {
      let $dt = new D26(0);
      $dt.cf122 = cf122;
      return $dt;
    }
    get is_DC71() { return this.$tag === 0; }
    get dtor_cf122() { return this.cf122; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC71" + "(" + _dafny.toString(this.cf122) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf122 === other.cf122;
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
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC73(cf124, cf125, cf126) {
      let $dt = new D27(0);
      $dt.cf124 = cf124;
      $dt.cf125 = cf125;
      $dt.cf126 = cf126;
      return $dt;
    }
    static create_DC74(cf127, cf128, cf129, cf130, cf131) {
      let $dt = new D27(1);
      $dt.cf127 = cf127;
      $dt.cf128 = cf128;
      $dt.cf129 = cf129;
      $dt.cf130 = cf130;
      $dt.cf131 = cf131;
      return $dt;
    }
    static create_DC75(cf132, cf133) {
      let $dt = new D27(2);
      $dt.cf132 = cf132;
      $dt.cf133 = cf133;
      return $dt;
    }
    static create_DC72(cf123) {
      let $dt = new D27(3);
      $dt.cf123 = cf123;
      return $dt;
    }
    get is_DC73() { return this.$tag === 0; }
    get is_DC74() { return this.$tag === 1; }
    get is_DC75() { return this.$tag === 2; }
    get is_DC72() { return this.$tag === 3; }
    get dtor_cf124() { return this.cf124; }
    get dtor_cf125() { return this.cf125; }
    get dtor_cf126() { return this.cf126; }
    get dtor_cf127() { return this.cf127; }
    get dtor_cf128() { return this.cf128; }
    get dtor_cf129() { return this.cf129; }
    get dtor_cf130() { return this.cf130; }
    get dtor_cf131() { return this.cf131; }
    get dtor_cf132() { return this.cf132; }
    get dtor_cf133() { return this.cf133; }
    get dtor_cf123() { return this.cf123; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC73" + "(" + _dafny.toString(this.cf124) + ", " + _dafny.toString(this.cf125) + ", " + _dafny.toString(this.cf126) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC74" + "(" + _dafny.toString(this.cf127) + ", " + _dafny.toString(this.cf128) + ", " + _dafny.toString(this.cf129) + ", " + _dafny.toString(this.cf130) + ", " + _dafny.toString(this.cf131) + ")";
      } else if (this.$tag === 2) {
        return "D27.DC75" + "(" + _dafny.toString(this.cf132) + ", " + _dafny.toString(this.cf133) + ")";
      } else if (this.$tag === 3) {
        return "D27.DC72" + "(" + _dafny.toString(this.cf123) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf124 === other.cf124 && _dafny.areEqual(this.cf125, other.cf125) && _dafny.areEqual(this.cf126, other.cf126);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf127, other.cf127) && this.cf128 === other.cf128 && _dafny.areEqual(this.cf129, other.cf129) && this.cf130 === other.cf130 && _dafny.areEqual(this.cf131, other.cf131);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf132 === other.cf132 && this.cf133 === other.cf133;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf123, other.cf123);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC73(false, _dafny.ZERO, _dafny.Set.Empty);
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
    static create_DC77(cf135, cf136, cf137, cf138, cf139) {
      let $dt = new D28(0);
      $dt.cf135 = cf135;
      $dt.cf136 = cf136;
      $dt.cf137 = cf137;
      $dt.cf138 = cf138;
      $dt.cf139 = cf139;
      return $dt;
    }
    static create_DC76(cf134) {
      let $dt = new D28(1);
      $dt.cf134 = cf134;
      return $dt;
    }
    static create_DC78(cf140) {
      let $dt = new D28(2);
      $dt.cf140 = cf140;
      return $dt;
    }
    get is_DC77() { return this.$tag === 0; }
    get is_DC76() { return this.$tag === 1; }
    get is_DC78() { return this.$tag === 2; }
    get dtor_cf135() { return this.cf135; }
    get dtor_cf136() { return this.cf136; }
    get dtor_cf137() { return this.cf137; }
    get dtor_cf138() { return this.cf138; }
    get dtor_cf139() { return this.cf139; }
    get dtor_cf134() { return this.cf134; }
    get dtor_cf140() { return this.cf140; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC77" + "(" + _dafny.toString(this.cf135) + ", " + _dafny.toString(this.cf136) + ", " + _dafny.toString(this.cf137) + ", " + _dafny.toString(this.cf138) + ", " + _dafny.toString(this.cf139) + ")";
      } else if (this.$tag === 1) {
        return "D28.DC76" + "(" + _dafny.toString(this.cf134) + ")";
      } else if (this.$tag === 2) {
        return "D28.DC78" + "(" + _dafny.toString(this.cf140) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf135, other.cf135) && _dafny.areEqual(this.cf136, other.cf136) && _dafny.areEqual(this.cf137, other.cf137) && this.cf138 === other.cf138 && _dafny.areEqual(this.cf139, other.cf139);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf134, other.cf134);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf140, other.cf140);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC77(_dafny.ZERO, _dafny.Set.Empty, _dafny.ZERO, false, _dafny.Map.Empty);
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
    static create_DC80(cf142, cf143) {
      let $dt = new D29(0);
      $dt.cf142 = cf142;
      $dt.cf143 = cf143;
      return $dt;
    }
    static create_DC81(cf144, cf145, cf146) {
      let $dt = new D29(1);
      $dt.cf144 = cf144;
      $dt.cf145 = cf145;
      $dt.cf146 = cf146;
      return $dt;
    }
    static create_DC79(cf141) {
      let $dt = new D29(2);
      $dt.cf141 = cf141;
      return $dt;
    }
    get is_DC80() { return this.$tag === 0; }
    get is_DC81() { return this.$tag === 1; }
    get is_DC79() { return this.$tag === 2; }
    get dtor_cf142() { return this.cf142; }
    get dtor_cf143() { return this.cf143; }
    get dtor_cf144() { return this.cf144; }
    get dtor_cf145() { return this.cf145; }
    get dtor_cf146() { return this.cf146; }
    get dtor_cf141() { return this.cf141; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC80" + "(" + _dafny.toString(this.cf142) + ", " + _dafny.toString(this.cf143) + ")";
      } else if (this.$tag === 1) {
        return "D29.DC81" + "(" + _dafny.toString(this.cf144) + ", " + _dafny.toString(this.cf145) + ", " + _dafny.toString(this.cf146) + ")";
      } else if (this.$tag === 2) {
        return "D29.DC79" + "(" + _dafny.toString(this.cf141) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf142, other.cf142) && this.cf143 === other.cf143;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf144 === other.cf144 && _dafny.areEqual(this.cf145, other.cf145) && _dafny.areEqual(this.cf146, other.cf146);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf141, other.cf141);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D29.create_DC80(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D29.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f1 = _dafny.ZERO;
      this.f3 = _dafny.Set.Empty;
      this.f7 = _dafny.ZERO;
      this._f2 = false;
      this._f4 = _dafny.Set.Empty;
      this._f5 = _dafny.MultiSet.Empty;
      this._f6 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      return;
    }
    get f2() {
      let _this = this;
      return _this._f2;
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
    fm36(globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber(744))).length), (_dafny.ZERO).minus(new BigNumber(-290)))).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(306))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-286), new BigNumber((_dafny.Seq.of(new BigNumber(304), new BigNumber(-160))).length))))).length));
    };
    fm37(p0, p1, p2, p3, globalState) {
      let _this = this;
      return false;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f10 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f10) {
      let _this = this;
      (_this)._f10 = f10;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_this).f10, (_this).f10, false)).IsProperSubsetOf(_dafny.MultiSet.fromElements(!((_this).f10)));
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return (((_dafny.ZERO).minus(new BigNumber(-283))).minus(new BigNumber(-605))).multipliedBy(new BigNumber(((((_this).f10) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(622)), function (_455_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })) : (_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))))).length));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(((((_this).f10) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.UnicodeFromString("mljbys"))) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.UnicodeFromString("gekebnpd"))))).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.UnicodeFromString("vup"))));
    };
    fm7(p0, globalState) {
      let _this = this;
      return !((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-457)))).equals((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_this).f10)).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f10)).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.Seq.UnicodeFromString("xgpysmdex")).length))).length)))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mifruf")).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(55)), function (_456_i0) {
        return new BigNumber(125);
      })).length))))));
    };
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _457_v0;
      _457_v0 = new BigNumber(747);
      let _pat_let_tv11 = _457_v0;
      let _pat_let_tv12 = _457_v0;
      (globalState).f7 = function (_source15) {
        if (_source15.is_DC4) {
          let _458___mcc_h0 = (_source15).cf5;
          let _459___mcc_h1 = (_source15).cf6;
          let _460_cf6 = _459___mcc_h1;
          let _461_cf5 = _458___mcc_h0;
          return new BigNumber((_dafny.Seq.of((_this).f10, !((_this).f10))).length);
        } else if (_source15.is_DC3) {
          let _462___mcc_h2 = (_source15).cf4;
          let _463_cf4 = _462___mcc_h2;
          return new BigNumber(163);
        } else {
          let _464___mcc_h3 = (_source15).cf7;
          let _465_cf7 = _464___mcc_h3;
          return _module.__default.safeDivisionInt(_pat_let_tv11, _pat_let_tv12);
        }
      }(_module.D1.create_DC3(_457_v0));
      let _466_v1;
      let _init6 = function (_467_i0) {
        return _module.__default.safeDivisionInt(_467_i0, new BigNumber(287));
      };
      let _nw63 = Array((new BigNumber(23)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw63.length); _i0_6++) {
        _nw63[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _466_v1 = _nw63;
      let _index70 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length));
      (_466_v1)[_index70] = _457_v0;
      let _468_v2;
      _468_v2 = new _dafny.CodePoint('r'.codePointAt(0));
      let _index71 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length));
      (_466_v1)[_index71] = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(_468_v2)).cardinality()), _457_v0), (_457_v0).multipliedBy(_457_v0));
      let _469_v3;
      let _init7 = function (_470_i1) {
        return (_this).f10;
      };
      let _nw64 = Array((new BigNumber(25)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw64.length); _i0_7++) {
        _nw64[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _469_v3 = _nw64;
      let _471_v4;
      _471_v4 = _dafny.Seq.of(_469_v3);
      let _472_v5;
      _472_v5 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),_471_v4);
      _471_v4 = (((_472_v5).contains((_this).f10)) ? ((_472_v5).get((_this).f10)) : (_dafny.Seq.of(_469_v3)));
      r1 = (_this).f10;
      let _473_v6;
      _473_v6 = _dafny.Seq.of((_this).f10, (_this).f10, false);
      let _474_v7;
      _474_v7 = _dafny.Seq.UnicodeFromString("raqvjn");
      let _475_v8;
      _475_v8 = _dafny.Map.Empty.slice().updateUnsafe(_474_v7,_468_v2);
      let _index72 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length));
      (_466_v1)[_index72] = ((_dafny.Seq.IsPrefixOf(_473_v6, _473_v6)) ? (((_466_v1)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length))]).plus(new BigNumber((_475_v8).length))) : (_457_v0));
      let _index73 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length));
      (_466_v1)[_index73] = _457_v0;
      let _476_v9;
      _476_v9 = _dafny.Map.Empty.slice().updateUnsafe(_474_v7,(_this).f10);
      let _477_v11;
      _477_v11 = _dafny.Set.fromElements(_474_v7, _474_v7);
      r0 = ((function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_476_v9).Keys.Elements) {
          let _478_v10 = _compr_16;
          if ((_476_v9).contains(_478_v10)) {
            _coll16.add(_478_v10);
          }
        }
        return _coll16;
      }()).Intersect(_477_v11)).Intersect(_477_v11);
      let _479_v12;
      _479_v12 = _dafny.Set.fromElements((_466_v1)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length))], (_466_v1)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length))], (_466_v1)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_466_v1).length))], new BigNumber(796), _457_v0);
      r1 = ((((_this).f10) || (!((_this).f10))) ? (!((_this).f10)) : ((_479_v12).IsDisjointFrom(_479_v12)));
      r2 = _457_v0;
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      let _480_v0;
      _480_v0 = _dafny.Seq.of(!((_this).f10));
      _480_v0 = (((_this).f10) ? (_dafny.Seq.Concat(_480_v0, _dafny.Seq.of((_this).f10, (_this).f10, (_this).f10))) : (_480_v0));
      let _481_v1;
      _481_v1 = _dafny.Set.fromElements((_this).f10, (_this).f10);
      let _482_v2;
      _482_v2 = _dafny.Seq.of(p0, p0, new BigNumber((_481_v1).length));
      let _483_v3;
      let _nw65 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _483_v3 = _nw65;
      let _484_v4;
      _484_v4 = _dafny.MultiSet.fromElements(_483_v3);
      let _485_v5;
      _485_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10);
      let _486_v6;
      let _nw66 = Array((new BigNumber(14)).toNumber());
      _nw66[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsProperPrefixOf(_482_v2, _482_v2);
      _nw66[(_dafny.ONE).toNumber()] = (_this).f10;
      _nw66[(new BigNumber(2)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.of(_480_v0, _480_v0, _480_v0, _dafny.Seq.update(_480_v0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(676)), ((_487_p0) => function (_488_i0) {
        return _487_p0;
      })(p0))).length), new BigNumber((_480_v0).length)), (_this).f10), _480_v0), _480_v0);
      _nw66[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.fromElements(_483_v3)).IsDisjointFrom(_484_v4);
      _nw66[(new BigNumber(4)).toNumber()] = (p0).isEqualTo(p0);
      _nw66[(new BigNumber(5)).toNumber()] = (_481_v1).IsDisjointFrom(_481_v1);
      _nw66[(new BigNumber(6)).toNumber()] = !(!((_this).fm6((_dafny.ZERO).minus(p0), _module.__default.fm1(globalState), p0, (_this).f10, globalState)));
      _nw66[(new BigNumber(7)).toNumber()] = ((((_485_v5).contains((_480_v0)[_module.__default.safeIndex(p0, new BigNumber((_480_v0).length))])) ? ((_485_v5).get((_480_v0)[_module.__default.safeIndex(p0, new BigNumber((_480_v0).length))])) : ((_this).f10))) && ((_this).f10);
      _nw66[(new BigNumber(8)).toNumber()] = (_this).f10;
      _nw66[(new BigNumber(9)).toNumber()] = !((_this).f10);
      _nw66[(new BigNumber(10)).toNumber()] = (_this).f10;
      _nw66[(new BigNumber(11)).toNumber()] = (_this).f10;
      _nw66[(new BigNumber(12)).toNumber()] = !((_this).f10) || (!((_this).f10));
      _nw66[(new BigNumber(13)).toNumber()] = (_this).f10;
      _486_v6 = _nw66;
      let _489_v7;
      _489_v7 = _dafny.Seq.of(_486_v6, _486_v6, _486_v6, _486_v6, _486_v6);
      _486_v6 = (_489_v7)[_module.__default.safeIndex(p0, new BigNumber((_489_v7).length))];
      let _490_v8;
      let _nw67 = Array((new BigNumber(7)).toNumber()).fill(_module.D2.Default());
      _490_v8 = _nw67;
      let _491_v9;
      _491_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,_490_v8);
      let _492_v10;
      _492_v10 = _module.D6.create_DC19(_491_v9);
      let _493_v11;
      _493_v11 = _dafny.Map.Empty.slice().updateUnsafe(_492_v10,!((_this).f10));
      let _494_v12;
      _494_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
      let _pat_let_tv13 = _491_v9;
      _493_v11 = (_493_v11).update(function (_pat_let6_0) {
        return function (_495_dt__update__tmp_h0) {
          return function (_pat_let7_0) {
            return function (_496_dt__update_hcf31_h0) {
              return _module.D6.create_DC19(_496_dt__update_hcf31_h0);
            }(_pat_let7_0);
          }(_pat_let_tv13);
        }(_pat_let6_0);
      }(_492_v10), (new BigNumber(-462)).isLessThan(new BigNumber(((_494_v12).update((_this).f10, p0)).length)));
      let _497_v13;
      _497_v13 = true;
      let _498_v14;
      let _nw68 = Array((new BigNumber(22)).toNumber()).fill([]);
      _498_v14 = _nw68;
      let _index74 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_498_v14).length));
      (_498_v14)[_index74] = _486_v6;
      let _499_v15;
      _499_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f10);
      let _500_v16;
      _500_v16 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),(((_499_v15).contains(new BigNumber(240))) ? ((_499_v15).get(new BigNumber(240))) : (false)));
      let _501_v17;
      _501_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, globalState),_483_v3);
      let _index75 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_498_v14).length));
      let _rhs61 = (((_500_v16).contains(new BigNumber(((_501_v17).Merge(_501_v17)).length))) ? ((_500_v16).get(new BigNumber(((_501_v17).Merge(_501_v17)).length))) : ((_this).f10));
      let _rhs62 = ((_497_v13) ? (_486_v6) : (_486_v6));
      let _lhs48 = _498_v14;
      let _lhs49 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_498_v14).length));
      _497_v13 = _rhs61;
      _lhs48[_lhs49] = _rhs62;
      let _502_v18;
      _502_v18 = _dafny.Set.fromElements(p0);
      let _503_v19;
      _503_v19 = _dafny.Seq.UnicodeFromString("xmpiyma");
      let _504_v20;
      _504_v20 = _dafny.Map.Empty.slice().updateUnsafe((_503_v19)[_module.__default.safeIndex(p0, new BigNumber((_503_v19).length))],(_this).f10);
      let _505_v21;
      _505_v21 = new _dafny.CodePoint('t'.codePointAt(0));
      let _506_v22;
      _506_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_505_v21);
      _497_v13 = (_502_v18).IsDisjointFrom((_module.__default.fm35(_482_v2, _dafny.Seq.of(p0, new BigNumber((_504_v20).length), new BigNumber(195)), _497_v13, globalState)).Difference(_dafny.Set.fromElements(new BigNumber((_506_v22).length), p0)));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_486_v6).length))) {
        let _507_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_507_i1)) && ((_507_i1).isLessThan(new BigNumber((_486_v6).length))))) {
          (_486_v6)[(_507_i1)] = !((!(_497_v13)) && (_497_v13)) || (!(new BigNumber(205)).isEqualTo(new BigNumber(-815)));
        }
      }
      return;
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Map.Empty;
      let _508_i0;
      _508_i0 = _dafny.ZERO;
      L2: {
        while (!((_this).f10)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_508_i0)) {
              break L2;
            }
            _508_i0 = (_508_i0).plus(_dafny.ONE);
            let _509_v0;
            let _nw69 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
            _509_v0 = _nw69;
            let _510_v2;
            _510_v2 = new _dafny.CodePoint('m'.codePointAt(0));
            let _511_v3;
            _511_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,_510_v2);
            let _512_v4;
            _512_v4 = _dafny.Seq.of(_511_v3, _511_v3, _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_510_v2));
            let _index76 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_509_v0).length));
            (_509_v0)[_index76] = function () {
              let _coll17 = new _dafny.Map();
              for (const _compr_17 of (_512_v4).Elements) {
                let _513_v1 = _compr_17;
                if (_dafny.Seq.contains(_512_v4, _513_v1)) {
                  _coll17.push([_513_v1,!(!((_this).f10))]);
                }
              }
              return _coll17;
            }();
            let _514_v5;
            _514_v5 = _dafny.Map.Empty.slice().updateUnsafe(_511_v3,(_this).f10);
            let _index77 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_509_v0).length));
            (_509_v0)[_index77] = (_514_v5);
            let _515_v6;
            _515_v6 = new BigNumber(585);
            r0 = (_515_v6).isEqualTo(_515_v6);
            let _516_v7;
            _516_v7 = _module.D4.create_DC14(_515_v6, new BigNumber(699));
            let _517_v8;
            _517_v8 = _dafny.Map.Empty.slice().updateUnsafe(_515_v6,_516_v7);
            _517_v8 = (_dafny.Map.Empty.slice().updateUnsafe(_515_v6,_module.D4.create_DC14(_515_v6, (_dafny.ZERO).minus(_515_v6)))).Merge(_517_v8);
            let _518_v9;
            _518_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_515_v6);
            if (_module.__default.fm0((((_518_v9).contains(p1)) ? ((_518_v9).get(p1)) : (_515_v6)), globalState)) {
              let _519_v10;
              let _nw70 = Array((new BigNumber(28)).toNumber()).fill(false);
              _519_v10 = _nw70;
              let _index78 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_519_v10).length));
              (_519_v10)[_index78] = (_this).fm8(_515_v6, (_this).f10, p1, globalState);
              let _520_v11;
              _520_v11 = _dafny.Seq.UnicodeFromString("ojm");
              let _521_v12;
              _521_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10);
              let _522_v13;
              _522_v13 = _dafny.MultiSet.fromElements(_518_v9);
              let _index79 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_519_v10).length));
              let _rhs63 = (_this).fm9(_520_v11, _dafny.Set.fromElements(_521_v12), globalState);
              let _rhs64 = (_this).f10;
              let _rhs65 = !(_518_v9).contains((_522_v13).IsProperSubsetOf(_522_v13));
              let _rhs66 = _module.__default.safeDivisionInt(_515_v6, _515_v6);
              let _lhs50 = globalState;
              let _lhs51 = _519_v10;
              let _lhs52 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_519_v10).length));
              let _lhs53 = globalState;
              _lhs50.f1 = _rhs63;
              r0 = _rhs64;
              _lhs51[_lhs52] = _rhs65;
              _lhs53.f1 = _rhs66;
              let _523_v14;
              let _init8 = ((_524_v6, _525_v12) => function (_526_i1) {
                return _module.__default.safeModuloInt(_526_i1, (_module.D3.create_DC11(_524_v6, _525_v12)).dtor_cf17);
              })(_515_v6, _521_v12);
              let _nw71 = Array((_dafny.ONE).toNumber());
              for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw71.length); _i0_8++) {
                _nw71[_i0_8] = _init8(new BigNumber(_i0_8));
              }
              _523_v14 = _nw71;
              let _527_v15;
              _527_v15 = _dafny.MultiSet.fromElements(_515_v6);
              let _index80 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_523_v14).length));
              (_523_v14)[_index80] = (((_527_v15).contains(_515_v6)) ? ((_527_v15).get(_515_v6)) : (_515_v6));
              let _index81 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_523_v14).length));
              (_523_v14)[_index81] = _515_v6;
              let _528_v16;
              let _nw72 = new _module.C0();
              _nw72.__ctor();
              _528_v16 = _nw72;
              let _529_v17;
              _529_v17 = _dafny.Seq.of((_523_v14)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_523_v14).length))]);
              _529_v17 = _529_v17;
              let _530_v18;
              _530_v18 = _dafny.MultiSet.fromElements(p1, (_this).f10, false, (_this).f10, (_519_v10)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_519_v10).length))]);
              let _531_v19;
              _531_v19 = _dafny.Seq.of((_519_v10)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_519_v10).length))]);
              r0 = (_dafny.MultiSet.fromElements(p1)).IsDisjointFrom((_530_v18).update((_531_v19)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_529_v17).length)), new BigNumber((_531_v19).length))], _module.__default.abs(new BigNumber(-324))));
            } else {
              let _532_v20;
              let _nw73 = new _module.C0();
              _nw73.__ctor();
              _532_v20 = _nw73;
              r0 = (_this).f10;
              let _533_v21;
              _533_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_515_v6, globalState),p1);
              let _534_v22;
              _534_v22 = _dafny.Set.fromElements(!((((_533_v21).contains(p1)) ? ((_533_v21).get(p1)) : (p1))), p1, (_this).f10, p1);
              let _535_v23;
              _535_v23 = _module.D10.create_DC30(_dafny.Set.fromElements((_this).f10, (_this).f10, (_this).f10, p1));
              r0 = ((_535_v23).dtor_cf49).IsSubsetOf((_534_v22).Difference(_534_v22));
              let _536_v24;
              _536_v24 = _dafny.Seq.of(new BigNumber(523));
              let _537_v25;
              _537_v25 = _dafny.Seq.UnicodeFromString("c");
              let _rhs67 = ((_515_v6).multipliedBy(_515_v6)).minus(_515_v6);
              let _rhs68 = _536_v24;
              let _rhs69 = _dafny.Seq.Concat(_537_v25, _537_v25);
              let _lhs54 = globalState;
              _lhs54.f1 = _rhs67;
              _536_v24 = _rhs68;
              r1 = _rhs69;
              r0 = (_this).f10;
            }
          }
        }
      }
      let _538_v26;
      _538_v26 = new _dafny.CodePoint('p'.codePointAt(0));
      _538_v26 = _538_v26;
      let _539_v27;
      _539_v27 = new BigNumber(380);
      let _540_v28;
      _540_v28 = _dafny.Seq.of(_539_v27);
      let _541_v29;
      _541_v29 = _dafny.Seq.of((_module.D6.create_DC20((_this).f10, p1, (_this).f10, new BigNumber((_540_v28).length))).dtor_cf33, p1);
      let _542_v30;
      _542_v30 = _dafny.Seq.UnicodeFromString("obshd");
      let _543_v31;
      _543_v31 = _dafny.Set.fromElements(_542_v30);
      let _544_v32;
      _544_v32 = _dafny.MultiSet.fromElements(p1);
      let _545_v33;
      _545_v33 = _module.D5.create_DC17((_this).fm8(_539_v27, (_this).f10, (_this).f10, globalState), _538_v26);
      let _546_v34;
      _546_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_544_v32).cardinality()),(_545_v33).dtor_cf28);
      let _547_v35;
      _547_v35 = _module.D6.create_DC20(p1, p1, (_this).f10, _539_v27);
      let _548_v36;
      let _nw74 = Array((new BigNumber(24)).toNumber());
      _nw74[(_dafny.ZERO).toNumber()] = (_this).f10;
      _nw74[(_dafny.ONE).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(2)).toNumber()] = (_541_v29)[_module.__default.safeIndex(_539_v27, new BigNumber((_541_v29).length))];
      _nw74[(new BigNumber(3)).toNumber()] = (_539_v27).isEqualTo(new BigNumber(728));
      _nw74[(new BigNumber(4)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(5)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(6)).toNumber()] = !(_539_v27).isEqualTo(_539_v27);
      _nw74[(new BigNumber(7)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(8)).toNumber()] = (_543_v31).IsSubsetOf(_543_v31);
      _nw74[(new BigNumber(9)).toNumber()] = (((_546_v34).contains(_539_v27)) ? ((_546_v34).get(_539_v27)) : ((_this).f10));
      _nw74[(new BigNumber(10)).toNumber()] = (((_this).f10) ? (!(!((_this).f10))) : (false));
      _nw74[(new BigNumber(11)).toNumber()] = (_547_v35).dtor_cf32;
      _nw74[(new BigNumber(12)).toNumber()] = (((_546_v34).contains((_dafny.ZERO).minus(_module.__default.fm1(globalState)))) ? ((_546_v34).get((_dafny.ZERO).minus(_module.__default.fm1(globalState)))) : (p1));
      _nw74[(new BigNumber(13)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(14)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(15)).toNumber()] = (_539_v27).isLessThan(_539_v27);
      _nw74[(new BigNumber(16)).toNumber()] = p1;
      _nw74[(new BigNumber(17)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(18)).toNumber()] = p1;
      _nw74[(new BigNumber(19)).toNumber()] = p1;
      _nw74[(new BigNumber(20)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(21)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_539_v27,true)).length)).isLessThan(_539_v27);
      _nw74[(new BigNumber(22)).toNumber()] = (_this).f10;
      _nw74[(new BigNumber(23)).toNumber()] = (_this).f10;
      _548_v36 = _nw74;
      let _index82 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length));
      (_548_v36)[_index82] = ((true) ? (!(!(p1))) : (p1));
      let _549_v37;
      let _init9 = ((_550_v27) => function (_551_i2) {
        return _module.__default.safeDivisionInt(_551_i2, _550_v27);
      })(_539_v27);
      let _nw75 = Array((new BigNumber(28)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw75.length); _i0_9++) {
        _nw75[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _549_v37 = _nw75;
      let _552_v38;
      _552_v38 = _module.D1.create_DC4(_539_v27, _549_v37);
      let _553_v39;
      _553_v39 = _module.D2.create_DC6(_538_v26);
      let _554_v40;
      _554_v40 = _dafny.Map.Empty.slice().updateUnsafe(_553_v39,_549_v37);
      let _555_v41;
      _555_v41 = _dafny.Map.Empty.slice().updateUnsafe(p1,_549_v37);
      let _pat_let_tv14 = _549_v37;
      let _556_v42;
      let _nw76 = Array((new BigNumber(29)).toNumber());
      _nw76[(_dafny.ZERO).toNumber()] = _549_v37;
      _nw76[(_dafny.ONE).toNumber()] = _549_v37;
      _nw76[(new BigNumber(2)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(3)).toNumber()] = (((_this).f10) ? (_549_v37) : (_549_v37));
      _nw76[(new BigNumber(4)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(5)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(6)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(7)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(8)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(9)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(10)).toNumber()] = (_552_v38).dtor_cf6;
      _nw76[(new BigNumber(11)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(12)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(13)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(14)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(15)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(16)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(17)).toNumber()] = (_module.D1.create_DC4((_dafny.ZERO).minus(_539_v27), _549_v37)).dtor_cf6;
      _nw76[(new BigNumber(18)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(19)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(20)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(21)).toNumber()] = (function (_pat_let8_0) {
        return function (_557_dt__update__tmp_h0) {
          return function (_pat_let9_0) {
            return function (_558_dt__update_hcf6_h0) {
              return _module.D1.create_DC4((_557_dt__update__tmp_h0).dtor_cf5, _558_dt__update_hcf6_h0);
            }(_pat_let9_0);
          }(_pat_let_tv14);
        }(_pat_let8_0);
      }(_552_v38)).dtor_cf6;
      _nw76[(new BigNumber(22)).toNumber()] = (((_554_v40).contains(_553_v39)) ? ((_554_v40).get(_553_v39)) : (_549_v37));
      _nw76[(new BigNumber(23)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(24)).toNumber()] = (((_555_v41).contains((_this).f10)) ? ((_555_v41).get((_this).f10)) : (_549_v37));
      _nw76[(new BigNumber(25)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(26)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(27)).toNumber()] = _549_v37;
      _nw76[(new BigNumber(28)).toNumber()] = _549_v37;
      _556_v42 = _nw76;
      let _559_v43;
      _559_v43 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
      let _index83 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length));
      let _rhs70 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_this).f10, p1), _541_v29);
      let _rhs71 = ((p1) ? (_556_v42) : (_556_v42));
      let _rhs72 = _dafny.Seq.UnicodeFromString("xk");
      let _rhs73 = !(p1) || ((((_559_v43).contains(p1)) ? ((_559_v43).get(p1)) : ((_this).f10)));
      let _lhs55 = _548_v36;
      let _lhs56 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length));
      _lhs55[_lhs56] = _rhs70;
      _556_v42 = _rhs71;
      r1 = _rhs72;
      r0 = _rhs73;
      let _index84 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length));
      let _rhs74 = ((true) ? ((true) === ((_this).fm7((_548_v36)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length))], globalState))) : ((_548_v36)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length))]));
      let _rhs75 = (_dafny.ZERO).minus(_539_v27);
      let _rhs76 = _539_v27;
      let _lhs57 = _548_v36;
      let _lhs58 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length));
      let _lhs59 = globalState;
      _lhs57[_lhs58] = _rhs74;
      _539_v27 = _rhs75;
      _lhs59.f7 = _rhs76;
      let _560_v44;
      _560_v44 = _dafny.Seq.of(_542_v30, _module.__default.fm38((_this).f10, _539_v27, globalState));
      let _561_v45;
      _561_v45 = _dafny.Map.Empty.slice().updateUnsafe(_539_v27,_560_v44);
      _561_v45 = (_561_v45).update(_539_v27, _560_v44);
      let _index85 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_548_v36).length));
      (_548_v36)[_index85] = (_this).f10;
      r0 = ((_this).f10) && (p1);
      r1 = _542_v30;
      let _562_v46;
      _562_v46 = _dafny.Map.Empty.slice().updateUnsafe(_539_v27,new BigNumber(715));
      r2 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_546_v34).length),_539_v27)).Merge((_562_v46).Merge(_dafny.Map.Empty.slice().updateUnsafe(_539_v27,_539_v27)));
      return [r0, r1, r2];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f21 = false;
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor(f21, f22) {
      let _this = this;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f21;
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return (_this).f22;
    };
    fm27(p0, p1, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),(_this).f22);
    };
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _563_v0;
      _563_v0 = new _dafny.CodePoint('i'.codePointAt(0));
      let _564_v1;
      _564_v1 = _dafny.Map.Empty.slice().updateUnsafe(_563_v0,_dafny.Seq.UnicodeFromString("rvocmt"));
      let _565_v2;
      _565_v2 = _dafny.Seq.UnicodeFromString("a");
      _564_v1 = ((_564_v1).Merge(_564_v1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_563_v0,_565_v2));
      if ((_this).f21) {
        let _566_v3;
        _566_v3 = _dafny.Set.fromElements(_565_v2, _565_v2, _565_v2, _565_v2);
        if ((((false) ? (_566_v3) : (_566_v3))).IsSubsetOf((_566_v3).Union(_module.__default.fm28((_this).f21, (_this).f21, globalState)))) {
          let _567_v4;
          let _init10 = function (_568_i0) {
            return _module.__default.safeModuloInt(_568_i0, new BigNumber(201));
          };
          let _nw77 = Array((new BigNumber(22)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw77.length); _i0_10++) {
            _nw77[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _567_v4 = _nw77;
          _567_v4 = _567_v4;
          let _569_v5;
          _569_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f21);
          _569_v5 = (_569_v5).update((_this).f22, (_this).f21);
          let _570_v6;
          _570_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
          let _571_v7;
          _571_v7 = _module.D3.create_DC11((_this).f22, (_570_v6).Merge(_570_v6));
          let _572_v8;
          _572_v8 = _module.D6.create_DC20((_this).f21, (_this).f21, (_this).f21, (_this).f22);
          let _rhs77 = _571_v7;
          let _rhs78 = (_572_v8).dtor_cf33;
          let _rhs79 = (_this).f22;
          let _rhs80 = (_this).f22;
          let _lhs60 = globalState;
          let _lhs61 = globalState;
          _571_v7 = _rhs77;
          r1 = _rhs78;
          _lhs60.f7 = _rhs79;
          _lhs61.f0 = _rhs80;
          (globalState).f7 = ((_this).f22).multipliedBy(((_this).f22).multipliedBy((_this).f22));
          _563_v0 = _563_v0;
        } else {
          let _573_v9;
          _573_v9 = _module.D7.create_DC23();
          _573_v9 = _573_v9;
          let _574_v10;
          let _nw78 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _574_v10 = _nw78;
          let _575_v11;
          _575_v11 = _dafny.Seq.of((_this).f21, (_this).f21);
          let _nw79 = Array((_dafny.ONE).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_575_v11, _dafny.Seq.of((_this).f21));
          _574_v10 = _nw79;
          let _576_v12;
          _576_v12 = _dafny.Seq.of((_this).f22, new BigNumber((_dafny.Seq.of((_this).f21)).length));
          let _577_v13;
          let _nw80 = Array((new BigNumber(8)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = true;
          _nw80[(_dafny.ONE).toNumber()] = (_this).f21;
          _nw80[(new BigNumber(2)).toNumber()] = true;
          _nw80[(new BigNumber(3)).toNumber()] = (_this).f21;
          _nw80[(new BigNumber(4)).toNumber()] = !((_this).f21);
          _nw80[(new BigNumber(5)).toNumber()] = (_this).f21;
          _nw80[(new BigNumber(6)).toNumber()] = (_this).f21;
          _nw80[(new BigNumber(7)).toNumber()] = (_this).f21;
          _577_v13 = _nw80;
          let _578_v14;
          _578_v14 = _dafny.Map.Empty.slice().updateUnsafe(_576_v12,_577_v13);
          let _579_v15;
          let _out28;
          _out28 = (_this).m15(_578_v14, (_this).f21, false, _563_v0, globalState);
          _579_v15 = _out28;
          let _580_v16;
          _580_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_581_i1) {
            return (_this).f22;
          }), _576_v12),_module.__default.safeDivisionInt(new BigNumber(670), _module.__default.fm1(globalState)));
          _580_v16 = (_580_v16).update((_this).f21, (_this).f22);
          r1 = (_this).f21;
        }
        (globalState).f7 = ((_this).f22).minus(new BigNumber((_565_v2).length));
        let _582_v17;
        _582_v17 = _dafny.MultiSet.fromElements(false, (_this).f21);
        let _583_v18;
        _583_v18 = _dafny.Seq.of((_this).f21, (_this).f21, true, (_this).f21);
        let _584_v19;
        _584_v19 = _dafny.Seq.of((_this).f22, new BigNumber((_583_v18).length));
        let _585_v20;
        _585_v20 = _dafny.Seq.of(_584_v19);
        let _586_v21;
        _586_v21 = _module.D2.create_DC7((((_582_v17).contains((_this).f21)) ? ((_582_v17).get((_this).f21)) : ((_this).f22)), _module.__default.fm0(new BigNumber((_585_v20).length), globalState), new BigNumber(-225));
        let _587_v22;
        _587_v22 = _dafny.Seq.of((_586_v21).dtor_cf10);
        _587_v22 = _dafny.Seq.Concat(_dafny.Seq.Concat(_583_v18, _587_v22), _dafny.Seq.Concat(_587_v22, _587_v22));
        let _rhs81 = ((((_582_v17).contains((_this).f21)) ? ((_582_v17).get((_this).f21)) : ((_584_v19)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_584_v19).length))]))).minus((_dafny.ZERO).minus(((_this).f22).plus(new BigNumber(49))));
        let _rhs82 = _582_v17;
        let _lhs62 = globalState;
        _lhs62.f0 = _rhs81;
        _582_v17 = _rhs82;
        let _588_v23;
        _588_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f22);
        let _589_v24;
        _589_v24 = _dafny.Map.Empty.slice().updateUnsafe(_588_v23,(_this).f22);
        let _source16 = _module.__default.fm29(new BigNumber((_589_v24).length), (_this).f21, _dafny.Seq.of((_this).f21, (_this).f21, (_this).fm8((_this).f22, (_this).f21, (_this).fm8((_this).f22, (_this).f21, (_this).f21, globalState), globalState), (_this).f21, true), globalState);
        if (_source16.is_DC10) {
          let _590___mcc_h0 = (_source16).cf14;
          let _591___mcc_h1 = (_source16).cf15;
          let _592___mcc_h2 = (_source16).cf16;
          let _593_cf16 = _592___mcc_h2;
          let _594_cf15 = _591___mcc_h1;
          let _595_cf14 = _590___mcc_h0;
          let _596_v25;
          let _nw81 = Array((new BigNumber(13)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = _563_v0;
          _nw81[(_dafny.ONE).toNumber()] = _563_v0;
          _nw81[(new BigNumber(2)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(3)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(4)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(5)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(6)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(7)).toNumber()] = _module.__default.fm4((_this).f21, _594_cf15, _594_cf15, globalState);
          _nw81[(new BigNumber(8)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(9)).toNumber()] = (_565_v2)[_module.__default.safeIndex((_this).f22, new BigNumber((_565_v2).length))];
          _nw81[(new BigNumber(10)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(11)).toNumber()] = _563_v0;
          _nw81[(new BigNumber(12)).toNumber()] = _563_v0;
          _596_v25 = _nw81;
          let _597_v26;
          _597_v26 = _module.D5.create_DC17((_this).f21, _563_v0);
          let _index86 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_596_v25).length));
          (_596_v25)[_index86] = (_597_v26).dtor_cf29;
          let _598_v27;
          _598_v27 = _module.D2.create_DC6(new _dafny.CodePoint('b'.codePointAt(0)));
          let _index87 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_596_v25).length));
          (_596_v25)[_index87] = (_598_v27).dtor_cf8;
          r1 = true;
          let _599_v28;
          _599_v28 = _dafny.Seq.of(_module.__default.fm30(globalState));
          let _600_v29;
          _600_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_dafny.Set.fromElements(_594_cf15));
          let _601_v30;
          _601_v30 = _dafny.Set.fromElements((_this).f21, false, true, !((_this).f21));
          _594_cf15 = _dafny.areEqual(_dafny.Seq.update(_599_v28, _module.__default.safeIndex(new BigNumber(((((_600_v29).contains(new BigNumber(220))) ? ((_600_v29).get(new BigNumber(220))) : (_601_v30))).length), new BigNumber((_599_v28).length)), _601_v30), _599_v28);
          let _602_v31;
          let _nw82 = Array((new BigNumber(22)).toNumber()).fill([]);
          _602_v31 = _nw82;
          let _603_v32;
          let _init11 = function (_604_i2) {
            return (_604_i2).minus((_this).f22);
          };
          let _nw83 = Array((new BigNumber(4)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw83.length); _i0_11++) {
            _nw83[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _603_v32 = _nw83;
          let _605_v33;
          _605_v33 = _dafny.Seq.of(_603_v32, _603_v32, _603_v32, _603_v32);
          let _index88 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_602_v31).length));
          (_602_v31)[_index88] = (_605_v33)[_module.__default.safeIndex((_this).f22, new BigNumber((_605_v33).length))];
          let _index89 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_602_v31).length));
          let _rhs83 = _dafny.Seq.Concat(_565_v2, _dafny.Seq.UnicodeFromString("j"));
          let _rhs84 = _603_v32;
          let _lhs63 = _602_v31;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_602_v31).length));
          _565_v2 = _rhs83;
          _lhs63[_lhs64] = _rhs84;
        } else if (_source16.is_DC11) {
          let _606___mcc_h3 = (_source16).cf17;
          let _607___mcc_h4 = (_source16).cf18;
          let _608_cf18 = _607___mcc_h4;
          let _609_cf17 = _606___mcc_h3;
          let _610_v34;
          let _nw84 = Array((new BigNumber(26)).toNumber()).fill(false);
          _610_v34 = _nw84;
          let _index90 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_610_v34).length));
          (_610_v34)[_index90] = !((_this).f21) || (_module.__default.fm0(_609_cf17, globalState));
          let _index91 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_610_v34).length));
          (_610_v34)[_index91] = (_this).f21;
          let _611_v35;
          let _nw85 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _611_v35 = _nw85;
          let _612_v36;
          _612_v36 = _dafny.Set.fromElements(!((_this).f21), true);
          let _613_v37;
          _613_v37 = _dafny.Set.fromElements(new BigNumber((_612_v36).length), (_dafny.ZERO).minus(new BigNumber((_584_v19).length)), _609_cf17, new BigNumber(595));
          let _index92 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_611_v35).length));
          (_611_v35)[_index92] = new BigNumber(((_613_v37).Intersect(_613_v37)).length);
          let _614_v38;
          _614_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f22),(_this).f21);
          let _index93 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_611_v35).length));
          (_611_v35)[_index93] = (_584_v19)[_module.__default.safeIndex(new BigNumber(((_614_v38).Merge(_614_v38)).length), new BigNumber((_584_v19).length))];
          _614_v38 = (_614_v38).update(new BigNumber(10), (_this).f21);
          let _index94 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_610_v34).length));
          (_610_v34)[_index94] = (_610_v34)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_610_v34).length))];
        } else if (_source16.is_DC9) {
          let _615___mcc_h5 = (_source16).cf13;
          let _616_cf13 = _615___mcc_h5;
          let _617_v39;
          let _nw86 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _617_v39 = _nw86;
          let _618_v40;
          _618_v40 = _dafny.Set.fromElements(_617_v39);
          _565_v2 = (((_618_v40).IsProperSubsetOf(_618_v40)) ? (_565_v2) : (_dafny.Seq.UnicodeFromString("uuqiag")));
          let _619_v41;
          _619_v41 = _dafny.Set.fromElements(!((_this).f21));
          let _rhs85 = false;
          let _rhs86 = (((_this).f21) ? ((false) || ((_this).f21)) : ((_this).f21));
          let _rhs87 = (_this).f22;
          let _rhs88 = !((((_619_v41).IsProperSubsetOf(_619_v41)) ? (((_this).f21) || ((_this).f21)) : ((_this).f21)));
          let _lhs65 = globalState;
          r1 = _rhs85;
          r1 = _rhs86;
          _lhs65.f7 = _rhs87;
          r1 = _rhs88;
          let _620_v42;
          _620_v42 = _dafny.Set.fromElements((_this).f22);
          let _621_v43;
          _621_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f21);
          let _622_v44;
          let _nw87 = Array((new BigNumber(27)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsPrefixOf(_565_v2, _565_v2);
          _nw87[(_dafny.ONE).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(2)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(3)).toNumber()] = (true) && ((_this).f21);
          _nw87[(new BigNumber(4)).toNumber()] = !((_this).fm8((_this).f22, (_this).f21, (_this).f21, globalState));
          _nw87[(new BigNumber(5)).toNumber()] = _module.__default.fm0(new BigNumber((_620_v42).length), globalState);
          _nw87[(new BigNumber(6)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(7)).toNumber()] = _dafny.Seq.contains(_584_v19, (_this).f22);
          _nw87[(new BigNumber(8)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(9)).toNumber()] = !(((_dafny.ZERO).minus((_this).f22)).isLessThan(new BigNumber((_588_v23).length)));
          _nw87[(new BigNumber(10)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(11)).toNumber()] = false;
          _nw87[(new BigNumber(12)).toNumber()] = !((_this).f21);
          _nw87[(new BigNumber(13)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(14)).toNumber()] = (((_621_v43).contains((_this).f22)) ? ((_621_v43).get((_this).f22)) : ((_this).f21));
          _nw87[(new BigNumber(15)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(16)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(17)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(18)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(19)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(20)).toNumber()] = !((_this).f21);
          _nw87[(new BigNumber(21)).toNumber()] = (_this).fm8((_this).f22, (_this).f21, (_this).f21, globalState);
          _nw87[(new BigNumber(22)).toNumber()] = !(!(new BigNumber(931)).isEqualTo((_this).f22));
          _nw87[(new BigNumber(23)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(24)).toNumber()] = (_this).f21;
          _nw87[(new BigNumber(25)).toNumber()] = (_620_v42).IsSubsetOf(_module.__default.fm31(new BigNumber(82), globalState));
          _nw87[(new BigNumber(26)).toNumber()] = (_this).f21;
          _622_v44 = _nw87;
          let _index95 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_622_v44).length));
          (_622_v44)[_index95] = (_this).f21;
          let _623_v45;
          _623_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
          let _index96 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_622_v44).length));
          let _rhs89 = (_module.D3.create_DC11((_this).f22, _623_v45)).dtor_cf17;
          let _rhs90 = !(!(new BigNumber(134)).isEqualTo(new BigNumber(-995))) || ((_this).f21);
          let _rhs91 = (_this).f21;
          let _lhs66 = globalState;
          let _lhs67 = _622_v44;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_622_v44).length));
          _lhs66.f1 = _rhs89;
          r1 = _rhs90;
          _lhs67[_lhs68] = _rhs91;
          let _index97 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_617_v39).length));
          (_617_v39)[_index97] = ((_this).f22).multipliedBy((_this).f22);
          let _index98 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_617_v39).length));
          (_617_v39)[_index98] = (_dafny.ZERO).minus(new BigNumber(-859));
        } else {
          let _624___mcc_h6 = (_source16).cf19;
          let _625_cf19 = _624___mcc_h6;
          let _626_v46;
          _626_v46 = _dafny.Set.fromElements((_this).f21);
          (globalState).f3 = _626_v46;
          (globalState).f7 = ((_this).f22).plus((_this).f22);
          r1 = _dafny.Seq.contains(_583_v18, ((_this).f22).isLessThan((_this).f22));
          let _627_v47;
          let _init12 = function (_628_i3) {
            return true;
          };
          let _nw88 = Array((new BigNumber(25)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw88.length); _i0_12++) {
            _nw88[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _627_v47 = _nw88;
          let _629_v48;
          _629_v48 = _dafny.Set.fromElements((_this).f22);
          let _index99 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_627_v47).length));
          (_627_v47)[_index99] = (_629_v48).IsSubsetOf(_629_v48);
          let _index100 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_627_v47).length));
          (_627_v47)[_index100] = (_this).f21;
        }
      } else {
        let _630_v49;
        _630_v49 = _dafny.Seq.of(((_this).f21) && ((_this).f21));
        _630_v49 = _dafny.Seq.of((_this).f21, (false) === ((_this).f21), (_this).f21);
        let _rhs92 = (_module.__default.safeDivisionInt((_this).f22, (_this).f22)).isLessThanOrEqualTo((_this).f22);
        let _rhs93 = (_this).f21;
        r1 = _rhs92;
        r1 = _rhs93;
        let _631_v50;
        let _init13 = ((_632_v0, _633_v2) => function (_634_i4) {
          return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("esg"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(537)), ((_635_v0) => function (_636_i5) {
            return _635_v0;
          })(_632_v0)), _633_v2)).IsDisjointFrom(_dafny.Set.fromElements(_633_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(710)), ((_637_v0) => function (_638_i6) {
            return _637_v0;
          })(_632_v0))));
        })(_563_v0, _565_v2);
        let _nw89 = Array((new BigNumber(18)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw89.length); _i0_13++) {
          _nw89[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _631_v50 = _nw89;
        let _639_v51;
        _639_v51 = _dafny.MultiSet.fromElements((_this).f22, new BigNumber(-467));
        let _index101 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length));
        (_631_v50)[_index101] = !(_module.__default.fm0(new BigNumber((_639_v51).cardinality()), globalState));
        let _640_v52;
        _640_v52 = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("imswxpkga"));
        let _pat_let_tv15 = _565_v2;
        let _641_v53;
        let _nw90 = Array((new BigNumber(28)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = _640_v52;
        _nw90[(_dafny.ONE).toNumber()] = _module.D0.create_DC1(_565_v2);
        _nw90[(new BigNumber(2)).toNumber()] = _module.__default.fm32((_this).f21, (_this).f22, false, globalState);
        _nw90[(new BigNumber(3)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(4)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(5)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(6)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(7)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(8)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(9)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(10)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(11)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(12)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(13)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(14)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(15)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(16)).toNumber()] = _module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(823)), ((_642_v0) => function (_643_i7) {
  return _642_v0;
})(_563_v0)));
        _nw90[(new BigNumber(17)).toNumber()] = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("uuajlyoaj"));
        _nw90[(new BigNumber(18)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(19)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(20)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(21)).toNumber()] = _module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(6)), ((_644_v0) => function (_645_i8) {
  return _644_v0;
})(_563_v0)));
        _nw90[(new BigNumber(22)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(23)).toNumber()] = _640_v52;
        _nw90[(new BigNumber(24)).toNumber()] = function (_pat_let10_0) {
          return function (_646_dt__update__tmp_h0) {
            return function (_pat_let11_0) {
              return function (_647_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_647_dt__update_hcf1_h0);
              }(_pat_let11_0);
            }(_pat_let_tv15);
          }(_pat_let10_0);
        }(_640_v52);
        _nw90[(new BigNumber(25)).toNumber()] = ((!((_this).f21)) ? (_640_v52) : (_640_v52));
        _nw90[(new BigNumber(26)).toNumber()] = _module.D0.create_DC1(_565_v2);
        _nw90[(new BigNumber(27)).toNumber()] = _module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_648_v0) => function (_649_i9) {
  return _648_v0;
})(_563_v0)));
        _641_v53 = _nw90;
        let _index102 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_641_v53).length));
        (_641_v53)[_index102] = _640_v52;
        let _index103 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length));
        let _index104 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_641_v53).length));
        let _rhs94 = (_this).f21;
        let _rhs95 = new BigNumber(180);
        let _rhs96 = _640_v52;
        let _lhs69 = _631_v50;
        let _lhs70 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length));
        let _lhs71 = globalState;
        let _lhs72 = _641_v53;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_641_v53).length));
        _lhs69[_lhs70] = _rhs94;
        _lhs71.f1 = _rhs95;
        _lhs72[_lhs73] = _rhs96;
        let _650_v54;
        _650_v54 = _dafny.MultiSet.fromElements(_563_v0, _563_v0, _563_v0);
        let _651_v55;
        _651_v55 = _dafny.Set.fromElements((_631_v50)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length))]);
        _650_v54 = (_650_v54).Union((_650_v54).update(_563_v0, _module.__default.abs((((_639_v51).contains((_this).f22)) ? ((_639_v51).get((_this).f22)) : (new BigNumber((_651_v55).length))))));
        let _652_v56;
        _652_v56 = _dafny.Map.Empty.slice().updateUnsafe((_630_v49)[_module.__default.safeIndex((_this).f22, new BigNumber((_630_v49).length))],(_this).f21);
        let _653_v57;
        _653_v57 = _dafny.Seq.of(((_652_v56).update((_631_v50)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length))], (_631_v50)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length))])).Merge(_module.__default.fm33((_631_v50)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length))], (_this).f21, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-744)), ((_654_v0) => function (_655_i10) {
          return _654_v0;
        })(_563_v0)), (_631_v50)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_631_v50).length))], globalState)));
        (globalState).f0 = new BigNumber((_653_v57).length);
      }
      let _source17 = _module.__default.fm34((_this).f22, globalState);
      if (_source17.is_DC7) {
        let _656___mcc_h7 = (_source17).cf9;
        let _657___mcc_h8 = (_source17).cf10;
        let _658___mcc_h9 = (_source17).cf11;
        let _659_cf11 = _658___mcc_h9;
        let _660_cf10 = _657___mcc_h8;
        let _661_cf9 = _656___mcc_h7;
        _565_v2 = _565_v2;
        let _662_v59;
        _662_v59 = _dafny.Seq.of(_660_cf10);
        let _663_v60;
        _663_v60 = _dafny.Set.fromElements((_this).f22, new BigNumber(4));
        let _664_v61;
        _664_v61 = _dafny.Map.Empty.slice().updateUnsafe(_662_v59,new BigNumber((_663_v60).length));
        let _665_v62;
        _665_v62 = _module.D4.create_DC13((function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (_664_v61).Keys.Elements) {
    let _666_v58 = _compr_18;
    if ((_664_v61).contains(_666_v58)) {
      _coll18.push([_666_v58,_660_cf10]);
    }
  }
  return _coll18;
}()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_662_v59,_660_cf10)));
        let _source18 = _665_v62;
        if (_source18.is_DC14) {
          let _667___mcc_h12 = (_source18).cf21;
          let _668___mcc_h13 = (_source18).cf22;
          let _669_cf22 = _668___mcc_h13;
          let _670_cf21 = _667___mcc_h12;
          let _671_v63;
          let _nw91 = Array((new BigNumber(7)).toNumber());
          _nw91[(_dafny.ZERO).toNumber()] = _660_cf10;
          _nw91[(_dafny.ONE).toNumber()] = (_this).f21;
          _nw91[(new BigNumber(2)).toNumber()] = (_this).f21;
          _nw91[(new BigNumber(3)).toNumber()] = (_this).f21;
          _nw91[(new BigNumber(4)).toNumber()] = _660_cf10;
          _nw91[(new BigNumber(5)).toNumber()] = (_this).f21;
          _nw91[(new BigNumber(6)).toNumber()] = (_this).f21;
          _671_v63 = _nw91;
          let _672_v64;
          _672_v64 = _dafny.MultiSet.fromElements(_671_v63);
          let _673_v65;
          _673_v65 = _dafny.Set.fromElements((_this).f21, !(true));
          let _rhs97 = (((_this).f21) ? (_dafny.MultiSet.fromElements(_671_v63)) : (_672_v64));
          let _rhs98 = (_673_v65).IsProperSubsetOf(_673_v65);
          _672_v64 = _rhs97;
          r1 = _rhs98;
          let _674_v66;
          _674_v66 = _dafny.Map.Empty.slice().updateUnsafe(_659_cf11,_dafny.Seq.Concat(_662_v59, _662_v59));
          _674_v66 = (_674_v66).update(_661_cf9, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_660_cf10), _dafny.Seq.of((_this).f21)), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_660_cf10, (_this).f21)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_660_cf10), _dafny.Seq.of((_this).f21))).length)), _660_cf10));
          let _675_v67;
          let _nw92 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _675_v67 = _nw92;
          let _index105 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_675_v67).length));
          (_675_v67)[_index105] = _669_cf22;
          let _676_v68;
          _676_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_661_cf9);
          let _index106 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_675_v67).length));
          (_675_v67)[_index106] = ((_660_cf10) ? (new BigNumber((_676_v68).length)) : (_661_cf9));
          r1 = (_this).fm8((_dafny.ZERO).minus(_659_cf11), _660_cf10, (_this).f21, globalState);
        } else if (_source18.is_DC15) {
          let _677___mcc_h14 = (_source18).cf23;
          let _678___mcc_h15 = (_source18).cf24;
          let _679___mcc_h16 = (_source18).cf25;
          let _680___mcc_h17 = (_source18).cf26;
          let _681_cf26 = _680___mcc_h17;
          let _682_cf25 = _679___mcc_h16;
          let _683_cf24 = _678___mcc_h15;
          let _684_cf23 = _677___mcc_h14;
          let _685_v69;
          let _nw93 = Array((new BigNumber(7)).toNumber()).fill([]);
          _685_v69 = _nw93;
          _685_v69 = (_module.D8.create_DC25(_685_v69)).dtor_cf42;
          _565_v2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-601)), ((_686_v2, _687_cf11, _688_cf9) => function (_689_i11) {
            return (_686_v2)[_module.__default.safeIndex((_module.D4.create_DC14(_687_cf11, _688_cf9)).dtor_cf22, new BigNumber((_686_v2).length))];
          })(_565_v2, _659_cf11, _661_cf9));
          r1 = !((_660_cf10) && ((_this).f21));
          let _rhs99 = true;
          let _rhs100 = _660_cf10;
          r1 = _rhs99;
          r1 = _rhs100;
        } else {
          let _690___mcc_h18 = (_source18).cf20;
          let _691_cf20 = _690___mcc_h18;
          let _692_v70;
          _692_v70 = _dafny.Set.fromElements(_660_cf10);
          (globalState).f3 = (_692_v70).Difference(_692_v70);
          let _693_v71;
          let _init14 = ((_694_v2, _695_v0) => function (_696_i12) {
            return !_dafny.areEqual(_694_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_697_v0) => function (_698_i13) {
              return _697_v0;
            })(_695_v0)));
          })(_565_v2, _563_v0);
          let _nw94 = Array((new BigNumber(9)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw94.length); _i0_14++) {
            _nw94[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _693_v71 = _nw94;
          _693_v71 = _693_v71;
          r1 = !(_dafny.Seq.IsProperPrefixOf(_565_v2, _565_v2));
          _693_v71 = _693_v71;
        }
        let _699_v72;
        _699_v72 = _dafny.Map.Empty.slice().updateUnsafe(!(_660_cf10),_661_cf9);
        r1 = (((_663_v60).IsSubsetOf(_663_v60)) ? (_660_cf10) : (((((_699_v72).contains((_this).f21)) ? ((_699_v72).get((_this).f21)) : ((_this).f22))).isLessThan(_659_cf11)));
        let _700_v73;
        let _nw95 = new _module.C1();
        _nw95.__ctor((_this).f21);
        _700_v73 = _nw95;
      } else if (_source17.is_DC6) {
        let _701___mcc_h10 = (_source17).cf8;
        let _702_cf8 = _701___mcc_h10;
        if ((_this).f21) {
          r1 = (_this).f21;
          let _703_v74;
          _703_v74 = _dafny.Seq.of(true, (_this).f21);
          let _704_v75;
          _704_v75 = _dafny.MultiSet.fromElements(_703_v74);
          let _705_v76;
          _705_v76 = _dafny.Map.Empty.slice().updateUnsafe(_704_v75,(_this).f22);
          let _706_v77;
          _706_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f21);
          let _707_v78;
          _707_v78 = _dafny.Map.Empty.slice().updateUnsafe((((_706_v77).contains((_this).f22)) ? ((_706_v77).get((_this).f22)) : ((_this).f21)),new BigNumber(61));
          let _708_v79;
          _708_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_707_v78);
          (globalState).f0 = (((_705_v76).contains((_704_v75).Difference(_dafny.MultiSet.fromElements(_703_v74)))) ? ((_705_v76).get((_704_v75).Difference(_dafny.MultiSet.fromElements(_703_v74)))) : (new BigNumber((_708_v79).length)));
          (globalState).f1 = (_this).f22;
          let _709_v80;
          let _nw96 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _709_v80 = _nw96;
          let _710_v81;
          _710_v81 = _dafny.Map.Empty.slice().updateUnsafe(_709_v80,_565_v2);
          let _711_v82;
          _711_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
          let _712_v83;
          _712_v83 = _dafny.Set.fromElements((_this).f21);
          _710_v81 = (((((_711_v82).contains((_this).f21)) ? ((_711_v82).get((_this).f21)) : ((_this).f21))) ? (_dafny.Map.Empty.slice().updateUnsafe(_709_v80,_565_v2)) : (_dafny.Map.Empty.slice().updateUnsafe(_709_v80,_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_565_v2, _module.__default.safeIndex(new BigNumber((_712_v83).length), new BigNumber((_565_v2).length)), _702_cf8), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.update(_565_v2, _module.__default.safeIndex(new BigNumber((_712_v83).length), new BigNumber((_565_v2).length)), _702_cf8)).length)), _563_v0), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_565_v2, _module.__default.safeIndex(new BigNumber((_712_v83).length), new BigNumber((_565_v2).length)), _702_cf8), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.update(_565_v2, _module.__default.safeIndex(new BigNumber((_712_v83).length), new BigNumber((_565_v2).length)), _702_cf8)).length)), _563_v0)).length)), _702_cf8))));
          let _713_v84;
          _713_v84 = _dafny.MultiSet.fromElements(_702_cf8);
          let _714_v85;
          _714_v85 = _module.D6.create_DC21((_this).f21, (_this).f21, (_this).f21, (((_713_v84).contains(_702_cf8)) ? ((_713_v84).get(_702_cf8)) : (new BigNumber(918))));
          r1 = (_714_v85).dtor_cf38;
        } else {
          let _715_v86;
          let _nw97 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _715_v86 = _nw97;
          let _index107 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_715_v86).length));
          (_715_v86)[_index107] = new BigNumber((_dafny.Seq.Concat(_565_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), ((_716_cf8) => function (_717_i14) {
            return _716_cf8;
          })(_702_cf8)))).length);
          let _index108 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_715_v86).length));
          (_715_v86)[_index108] = _module.__default.safeModuloInt((_this).f22, (_this).f22);
          let _718_v87;
          _718_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
          _718_v87 = (_718_v87).Merge((_718_v87).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,false)));
          let _719_v89;
          _719_v89 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-580)), ((_720_v0) => function (_721_i15) {
            return _720_v0;
          })(_563_v0)), _565_v2);
          let _722_v90;
          _722_v90 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_719_v89).Elements) {
              let _723_v88 = _compr_19;
              if (_dafny.Seq.contains(_719_v89, _723_v88)) {
                _coll19.push([_723_v88,(_this).f21]);
              }
            }
            return _coll19;
          }(),(((_this).f21) ? ((_this).f21) : ((_this).f21)));
          let _724_v91;
          _724_v91 = _dafny.Map.Empty.slice().updateUnsafe(_565_v2,(_this).fm8((_this).f22, (_this).f21, (_this).f21, globalState));
          _722_v90 = (_722_v90).update(_724_v91, (_this).f21);
          let _725_v92;
          let _init15 = function (_726_i16) {
            return _module.D2.create_DC6(new _dafny.CodePoint('t'.codePointAt(0)));
          };
          let _nw98 = Array((new BigNumber(19)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw98.length); _i0_15++) {
            _nw98[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _725_v92 = _nw98;
          let _727_v93;
          _727_v93 = _module.D2.create_DC6(_563_v0);
          let _index109 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_725_v92).length));
          (_725_v92)[_index109] = _727_v93;
          let _index110 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_725_v92).length));
          (_725_v92)[_index110] = _module.__default.fm39(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(71)), function (_728_i17) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_729_v0) => function (_730_i18) {
            return _729_v0;
          })(_563_v0)), false, globalState);
          r1 = _module.__default.fm0((((_this).f21) ? ((_this).f22) : ((_this).f22)), globalState);
        }
        let _731_v94;
        _731_v94 = _dafny.MultiSet.fromElements((_this).f21);
        let _732_v95;
        _732_v95 = _module.D5.create_DC16((_731_v94).update((_this).f21, _module.__default.abs((_this).f22)));
        r1 = !(((_732_v95).dtor_cf27).IsDisjointFrom(_731_v94)) || (!(true) || ((_this).f21));
        r1 = _module.__default.fm0(_module.__default.safeDivisionInt((_this).f22, (_this).f22), globalState);
        let _733_v96;
        _733_v96 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).fm8((_this).f22, (_this).f21, (_this).f21, globalState));
        let _734_v97;
        _734_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_733_v96).length),(_this).f21);
        let _735_v98;
        _735_v98 = _dafny.Set.fromElements((_this).f22, ((_this).f22).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f22)).length)), new BigNumber(((_734_v97).Merge(_734_v97)).length));
        _735_v98 = _module.__default.fm31((_this).f22, globalState);
      } else {
        let _736___mcc_h11 = (_source17).cf12;
        let _737_cf12 = _736___mcc_h11;
        r1 = (_this).f21;
        let _738_v99;
        let _nw99 = Array((new BigNumber(23)).toNumber()).fill(false);
        _738_v99 = _nw99;
        let _739_v100;
        _739_v100 = _module.D0.create_DC0(_738_v99);
        let _740_v101;
        let _nw100 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _740_v101 = _nw100;
        let _741_v102;
        let _nw101 = Array((new BigNumber(12)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = _740_v101;
        _nw101[(_dafny.ONE).toNumber()] = _740_v101;
        _nw101[(new BigNumber(2)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(3)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(4)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(5)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(6)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(7)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(8)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(9)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(10)).toNumber()] = _740_v101;
        _nw101[(new BigNumber(11)).toNumber()] = _740_v101;
        _741_v102 = _nw101;
        let _index111 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length));
        (_741_v102)[_index111] = _740_v101;
        let _742_v103;
        _742_v103 = _dafny.Map.Empty.slice().updateUnsafe(_563_v0,((_this).f21) && ((_this).f21));
        let _743_v105;
        _743_v105 = _dafny.Seq.of(false, (_this).f21);
        let _744_v106;
        _744_v106 = _dafny.Map.Empty.slice().updateUnsafe(_743_v105,(_this).f21);
        let _745_v107;
        _745_v107 = _module.D4.create_DC13(function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of (_744_v106).Keys.Elements) {
    let _746_v104 = _compr_20;
    if ((_744_v106).contains(_746_v104)) {
      _coll20.push([_746_v104,(_this).f21]);
    }
  }
  return _coll20;
}());
        let _index112 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length));
        let _rhs101 = _739_v100;
        let _rhs102 = _740_v101;
        let _rhs103 = _742_v103;
        let _rhs104 = (((_this).fm8((_this).f22, (_this).f21, (_this).f21, globalState)) ? (_565_v2) : (_dafny.Seq.UnicodeFromString("hgl")));
        let _rhs105 = _745_v107;
        let _lhs74 = _741_v102;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length));
        _739_v100 = _rhs101;
        _lhs74[_lhs75] = _rhs102;
        _742_v103 = _rhs103;
        _565_v2 = _rhs104;
        _745_v107 = _rhs105;
        r1 = !(false);
        let _747_v108;
        _747_v108 = _module.D8.create_DC26();
        let _source19 = _747_v108;
        if (_source19.is_DC26) {
          let _748_v109;
          _748_v109 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21);
          let _749_v110;
          _749_v110 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_748_v109).length),(_this).f21);
          let _rhs106 = _565_v2;
          let _rhs107 = _dafny.Seq.Concat(_743_v105, _dafny.Seq.Concat(_dafny.Seq.of((_this).f21, (_this).f21, (_this).f21, (_743_v105)[_module.__default.safeIndex(new BigNumber((_749_v110).length), new BigNumber((_743_v105).length))]), _module.__default.fm40((_this).f21, false, globalState)));
          let _rhs108 = (_this).f22;
          let _rhs109 = _module.__default.safeModuloInt((_this).f22, _module.__default.safeDivisionInt((_this).f22, (_this).f22));
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          _565_v2 = _rhs106;
          _743_v105 = _rhs107;
          _lhs76.f0 = _rhs108;
          _lhs77.f0 = _rhs109;
          r1 = !((_this).f21);
          let _750_v111;
          _750_v111 = _module.D3.create_DC11((_this).f22, _748_v109);
          let _751_v112;
          let _752_v113;
          let _out29;
          let _out30;
          let _outcollector14 = _module.__default.m0((_750_v111).dtor_cf17, ((!((_this).f21)) ? (_738_v99) : (_738_v99)), globalState);
          _out29 = _outcollector14[0];
          _out30 = _outcollector14[1];
          _751_v112 = _out29;
          _752_v113 = _out30;
          let _753_v114;
          let _init16 = function (_754_i19) {
            return _dafny.MultiSet.fromElements((_this).f21);
          };
          let _nw102 = Array((new BigNumber(17)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw102.length); _i0_16++) {
            _nw102[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _753_v114 = _nw102;
          let _755_v115;
          _755_v115 = _dafny.MultiSet.fromElements((_this).f21, (_this).f21, (_this).f21);
          let _756_v116;
          _756_v116 = _dafny.Seq.of(_755_v115);
          let _757_v117;
          _757_v117 = _module.D2.create_DC7((_this).f22, (_this).f21, (_dafny.ZERO).minus(_module.__default.fm1(globalState)));
          let _758_v118;
          _758_v118 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_dafny.MultiSet.FromArray(_743_v105));
          let _nw103 = Array((new BigNumber(23)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _755_v115;
          _nw103[(_dafny.ONE).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm40((_this).f21, (_this).f21, globalState));
          _nw103[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements((_this).f21, false);
          _nw103[(new BigNumber(3)).toNumber()] = (_756_v116)[_module.__default.safeIndex((_this).f22, new BigNumber((_756_v116).length))];
          _nw103[(new BigNumber(4)).toNumber()] = (_755_v115).Difference(_755_v115);
          _nw103[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_743_v105, _dafny.Seq.of((_this).f21, (_this).f21)));
          _nw103[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.FromArray(_743_v105)).update((_this).f21, _module.__default.abs((_this).f22));
          _nw103[(new BigNumber(7)).toNumber()] = (((_this).f21) ? (_dafny.MultiSet.fromElements(true)) : (_755_v115));
          _nw103[(new BigNumber(8)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(9)).toNumber()] = (_755_v115).Difference(_755_v115);
          _nw103[(new BigNumber(10)).toNumber()] = _module.__default.fm41(globalState);
          _nw103[(new BigNumber(11)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(12)).toNumber()] = (_755_v115).update((_757_v117).dtor_cf10, _module.__default.abs((_this).f22));
          _nw103[(new BigNumber(13)).toNumber()] = ((((_758_v118).contains((_this).f21)) ? ((_758_v118).get((_this).f21)) : (_755_v115))).Intersect(_dafny.MultiSet.FromArray(_743_v105));
          _nw103[(new BigNumber(14)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(15)).toNumber()] = (_755_v115).Intersect(_755_v115);
          _nw103[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_743_v105, _743_v105));
          _nw103[(new BigNumber(17)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements((_this).f21);
          _nw103[(new BigNumber(19)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(20)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(21)).toNumber()] = _755_v115;
          _nw103[(new BigNumber(22)).toNumber()] = _755_v115;
          _753_v114 = _nw103;
        } else if (_source19.is_DC27) {
          let _759___mcc_h19 = (_source19).cf43;
          let _760___mcc_h20 = (_source19).cf44;
          let _761___mcc_h21 = (_source19).cf45;
          let _762___mcc_h22 = (_source19).cf46;
          let _763_cf46 = _762___mcc_h22;
          let _764_cf45 = _761___mcc_h21;
          let _765_cf44 = _760___mcc_h20;
          let _766_cf43 = _759___mcc_h19;
          let _index113 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_738_v99).length));
          (_738_v99)[_index113] = _766_cf43;
          let _index114 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_738_v99).length));
          let _rhs110 = _766_cf43;
          let _rhs111 = new BigNumber((_dafny.Seq.UnicodeFromString("nungwwqg")).length);
          let _lhs78 = _738_v99;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_738_v99).length));
          let _lhs80 = globalState;
          _lhs78[_lhs79] = _rhs110;
          _lhs80.f0 = _rhs111;
          _766_cf43 = (_766_cf43) && (false);
          let _arr0 = (_741_v102)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length))];
          let _index115 = _module.__default.safeIndex(new BigNumber(972), new BigNumber(((_741_v102)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length))]).length));
          _arr0[_index115] = _module.__default.safeModuloInt(new BigNumber((_565_v2).length), _763_cf46);
          let _arr1 = (_741_v102)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length))];
          let _index116 = _module.__default.safeIndex(new BigNumber(972), new BigNumber(((_741_v102)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_741_v102).length))]).length));
          _arr1[_index116] = _module.__default.safeModuloInt(new BigNumber(-332), _763_cf46);
          (globalState).f7 = (_dafny.ZERO).minus((_this).f22);
        } else if (_source19.is_DC25) {
          let _767___mcc_h23 = (_source19).cf42;
          let _768_cf42 = _767___mcc_h23;
          let _769_v120;
          _769_v120 = _dafny.MultiSet.fromElements((_this).f22, (_this).f22, (_this).f22);
          let _770_v121;
          _770_v121 = _dafny.Seq.of((((_769_v120).contains((_this).f22)) ? ((_769_v120).get((_this).f22)) : ((_this).f22)), new BigNumber((_743_v105).length));
          let _771_v122;
          _771_v122 = _dafny.MultiSet.fromElements((_this).f22, new BigNumber((_770_v121).length));
          let _772_v123;
          _772_v123 = _dafny.Seq.of(_771_v122, _771_v122);
          (globalState).f0 = (((_this).f22).multipliedBy(new BigNumber((function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of (_772_v123).Elements) {
              let _773_v119 = _compr_21;
              if (_dafny.Seq.contains(_772_v123, _773_v119)) {
                _coll21.push([_773_v119,_770_v121]);
              }
            }
            return _coll21;
          }()).length))).plus((_this).f22);
          let _774_v124;
          _774_v124 = _dafny.Seq.of(_770_v121);
          let _775_v125;
          let _nw104 = Array((new BigNumber(11)).toNumber()).fill(_module.D1.Default());
          _775_v125 = _nw104;
          let _776_v126;
          _776_v126 = _dafny.Seq.of(_775_v125);
          let _777_v127;
          _777_v127 = _module.D8.create_DC27((_this).f21, _776_v126, _770_v121, (_this).f22);
          _770_v121 = _dafny.Seq.Concat((_774_v124)[_module.__default.safeIndex(new BigNumber((_565_v2).length), new BigNumber((_774_v124).length))], (_777_v127).dtor_cf45);
          let _778_v128;
          let _nw105 = new _module.C1();
          _nw105.__ctor((_this).f21);
          _778_v128 = _nw105;
          let _index117 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_738_v99).length));
          (_738_v99)[_index117] = (_this).f21;
          let _index118 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_738_v99).length));
          let _rhs112 = (_this).f21;
          let _rhs113 = ((_this).f22).minus(new BigNumber(-525));
          let _rhs114 = (_this).f21;
          let _lhs81 = globalState;
          let _lhs82 = _738_v99;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_738_v99).length));
          r1 = _rhs112;
          _lhs81.f0 = _rhs113;
          _lhs82[_lhs83] = _rhs114;
        } else {
          let _779___mcc_h24 = (_source19).cf47;
          let _780_cf47 = _779___mcc_h24;
          let _index119 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length));
          (_740_v101)[_index119] = (_this).f22;
          let _index120 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length));
          (_740_v101)[_index120] = (_this).f22;
          let _781_v129;
          _781_v129 = _dafny.Map.Empty.slice().updateUnsafe((_740_v101)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length))],(_this).fm8((_740_v101)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length))], (_this).f21, (_this).f21, globalState));
          let _782_v130;
          _782_v130 = _dafny.Map.Empty.slice().updateUnsafe((_781_v129).Merge(_781_v129),(_this).f22);
          let _pat_let_tv16 = _781_v129;
          let _pat_let_tv17 = _781_v129;
          let _index121 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length));
          (_740_v101)[_index121] = (((_782_v130).contains((function (_pat_let14_0) {
            return function (_785_dt__update__tmp_h1) {
              return function (_pat_let15_0) {
                return function (_786_dt__update_hcf53_h0) {
                  return _module.D11.create_DC32(_786_dt__update_hcf53_h0);
                }(_pat_let15_0);
              }(_pat_let_tv17);
            }(_pat_let14_0);
          }(_module.__default.fm42(globalState))).dtor_cf53)) ? ((_782_v130).get((function (_pat_let12_0) {
            return function (_783_dt__update__tmp_h2) {
              return function (_pat_let13_0) {
                return function (_784_dt__update_hcf53_h1) {
                  return _module.D11.create_DC32(_784_dt__update_hcf53_h1);
                }(_pat_let13_0);
              }(_pat_let_tv16);
            }(_pat_let12_0);
          }(_module.__default.fm42(globalState))).dtor_cf53)) : (((_740_v101)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length))]).plus((_this).f22)));
          let _787_v131;
          _787_v131 = _dafny.Seq.of((_this).f22);
          let _788_v132;
          _788_v132 = _dafny.MultiSet.fromElements((_this).f21);
          let _789_v133;
          _789_v133 = _dafny.MultiSet.fromElements(new BigNumber((_788_v132).cardinality()), (_this).f22, (_this).f22, (_740_v101)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length))]);
          let _790_v134;
          _790_v134 = _dafny.Set.fromElements(_789_v133);
          let _rhs115 = _738_v99;
          let _rhs116 = _563_v0;
          let _rhs117 = !(!((_this).f21));
          let _rhs118 = _dafny.Seq.Concat(_787_v131, _787_v131);
          let _rhs119 = _790_v134;
          _738_v99 = _rhs115;
          _563_v0 = _rhs116;
          r1 = _rhs117;
          _787_v131 = _rhs118;
          _790_v134 = _rhs119;
          let _791_v136;
          let _init17 = ((_792_v0) => function (_793_i20) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_792_v0,function () {
              let _coll22 = new _dafny.Map();
              for (const _compr_22 of _dafny.IntegerRange(new BigNumber(712), new BigNumber(334))) {
                let _794_v135 = _compr_22;
                if (((new BigNumber(712)).isLessThanOrEqualTo(_794_v135)) && ((_794_v135).isLessThan(new BigNumber(334)))) {
                  _coll22.push([(_794_v135).minus((_this).f22),_792_v0]);
                }
              }
              return _coll22;
            }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(_792_v0,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-560),_792_v0)));
          })(_563_v0);
          let _nw106 = Array((new BigNumber(24)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw106.length); _i0_17++) {
            _nw106[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _791_v136 = _nw106;
          let _795_v137;
          _795_v137 = _dafny.Map.Empty.slice().updateUnsafe((_740_v101)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_740_v101).length))],_563_v0);
          let _796_v138;
          _796_v138 = _dafny.Map.Empty.slice().updateUnsafe(_563_v0,_795_v137);
          let _797_v139;
          _797_v139 = _module.D5.create_DC17((_this).f21, _563_v0);
          let _index122 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_791_v136).length));
          (_791_v136)[_index122] = (_796_v138).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4((_this).f21, (_this).f21, (_797_v139).dtor_cf28, globalState),_795_v137));
          let _index123 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_791_v136).length));
          (_791_v136)[_index123] = _dafny.Map.Empty.slice().updateUnsafe(_563_v0,_795_v137);
        }
      }
      let _798_v140;
      let _nw107 = Array((new BigNumber(26)).toNumber()).fill(false);
      _798_v140 = _nw107;
      let _index124 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
      (_798_v140)[_index124] = (_this).f21;
      let _799_v141;
      _799_v141 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
      let _800_v142;
      _800_v142 = _dafny.MultiSet.fromElements((_this).f21);
      let _801_v143;
      _801_v143 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((((_799_v141).contains((_this).f21)) ? ((_799_v141).get((_this).f21)) : ((_this).f21)))).Intersect(_800_v142),_565_v2);
      let _802_v144;
      _802_v144 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f21);
      let _803_v145;
      _803_v145 = _dafny.MultiSet.fromElements(_802_v144);
      let _804_v146;
      _804_v146 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f22);
      let _805_v147;
      _805_v147 = _dafny.Map.Empty.slice().updateUnsafe(_804_v146,(_this).f22);
      let _806_v148;
      _806_v148 = _module.D6.create_DC20((_this).f21, (_this).f21, (_this).f21, (_this).f22);
      let _807_v149;
      _807_v149 = _dafny.Map.Empty.slice().updateUnsafe(_806_v148,(_this).f21);
      let _808_v150;
      _808_v150 = _dafny.Seq.of(true);
      let _809_v152;
      _809_v152 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_801_v143).Merge(function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), ((_810_v142) => function (_811_i21) {
          return _810_v142;
        })(_800_v142))).Elements) {
          let _812_v151 = _compr_23;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), ((_813_v142) => function (_811_i21) {
            return _813_v142;
          })(_800_v142)), _812_v151)) {
            _coll23.push([_812_v151,_565_v2]);
          }
        }
        return _coll23;
      }()));
      let _index125 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
      let _rhs120 = (_this).fm8(((_module.__default.fm0((_this).f22, globalState)) ? ((((_803_v145).contains(_802_v144)) ? ((_803_v145).get(_802_v144)) : ((_dafny.ZERO).minus((((_805_v147).contains(_804_v146)) ? ((_805_v147).get(_804_v146)) : (new BigNumber(-78))))))) : ((_this).f22)), (((((_807_v149).contains(_806_v148)) ? ((_807_v149).get(_806_v148)) : ((_this).f21))) ? ((_this).f21) : ((_this).f21)), (((_this).f21) ? ((_this).f21) : ((_this).f21)), globalState);
      let _rhs121 = ((_this).f22).isLessThan((_module.D11.create_DC34((_this).f22, (_this).f21, new BigNumber((_808_v150).length))).dtor_cf58);
      let _rhs122 = _563_v0;
      let _rhs123 = (((_809_v152).contains((_this).f22)) ? ((_809_v152).get((_this).f22)) : (_801_v143));
      let _lhs84 = _798_v140;
      let _lhs85 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
      _lhs84[_lhs85] = _rhs120;
      r1 = _rhs121;
      _563_v0 = _rhs122;
      _801_v143 = _rhs123;
      let _814_v153;
      _814_v153 = _module.D0.create_DC0(_798_v140);
      let _815_v154;
      _815_v154 = _dafny.Map.Empty.slice().updateUnsafe(_814_v153,(_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))]);
      r1 = _module.__default.fm0(new BigNumber((_815_v154).length), globalState);
      let _816_v155;
      _816_v155 = _module.D7.create_DC24((_this).f22);
      let _source20 = _816_v155;
      if (_source20.is_DC23) {
        let _817_v157;
        let _init18 = function (_818_i22) {
          return _module.D1.create_DC3(new BigNumber((function () {
  let _coll24 = new _dafny.Map();
  for (const _compr_24 of _dafny.IntegerRange(new BigNumber(128), new BigNumber(543))) {
    let _819_v156 = _compr_24;
    if (((new BigNumber(128)).isLessThanOrEqualTo(_819_v156)) && ((_819_v156).isLessThan(new BigNumber(543)))) {
      _coll24.push([(_819_v156).plus((_this).f22),(_this).f22]);
    }
  }
  return _coll24;
}()).length));
        };
        let _nw108 = Array((new BigNumber(16)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw108.length); _i0_18++) {
          _nw108[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _817_v157 = _nw108;
        let _820_v158;
        _820_v158 = _module.D8.create_DC28(_module.D8.create_DC27((_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))], _dafny.Seq.of(_817_v157), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), function (_821_i23) {
  return (_this).f22;
}), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), function (_822_i23) {
  return (_this).f22;
})).length)), new BigNumber(-2)), new BigNumber(791)));
        let _823_v159;
        _823_v159 = _module.D8.create_DC28(_820_v158);
        let _rhs124 = _823_v159;
        let _rhs125 = (_this).f22;
        _823_v159 = _rhs124;
        r2 = _rhs125;
        _802_v144 = (_802_v144).update((_dafny.ZERO).minus((_this).f22), (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))]);
        let _824_v160;
        _824_v160 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_dafny.MultiSet.fromElements((_this).f22));
        if ((new BigNumber((_824_v160).length)).isLessThan((_this).f22)) {
          let _825_v161;
          let _nw109 = new _module.C1();
          _nw109.__ctor((_this).f21);
          _825_v161 = _nw109;
          let _826_v162;
          _826_v162 = _dafny.Map.Empty.slice().updateUnsafe(_825_v161,_module.__default.fm1(globalState));
          let _827_v163;
          _827_v163 = _dafny.Seq.of(new BigNumber((_565_v2).length), new BigNumber((_826_v162).length), (_this).f22, (_this).f22, (_this).f22);
          let _828_v164;
          _828_v164 = _dafny.Map.Empty.slice().updateUnsafe(_827_v163,_798_v140);
          let _829_v165;
          let _out31;
          _out31 = (_this).m15(_828_v164, (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))], (_this).f21, _563_v0, globalState);
          _829_v165 = _out31;
          (globalState).f7 = (_this).f22;
          _808_v150 = _808_v150;
          _829_v165 = _829_v165;
          let _830_v166;
          _830_v166 = _module.D5.create_DC16(_dafny.MultiSet.fromElements((_this).f21));
          let _831_v167;
          _831_v167 = _module.D5.create_DC18(_830_v166);
          let _832_v168;
          _832_v168 = _module.D5.create_DC18(_830_v166);
          let _833_v169;
          _833_v169 = _dafny.Seq.of(_832_v168);
          let _834_v170;
          _834_v170 = _dafny.Map.Empty.slice().updateUnsafe(_833_v169,(_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))]);
          let _835_v171;
          _835_v171 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f22),_833_v169);
          _834_v170 = (_834_v170).update((((_835_v171).contains((_this).f22)) ? ((_835_v171).get((_this).f22)) : (_module.__default.fm43((_this).f22, (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))], (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))], globalState))), (_808_v150)[_module.__default.safeIndex((_this).f22, new BigNumber((_808_v150).length))]);
        } else {
          let _836_v172;
          _836_v172 = _module.D11.create_DC34((_this).f22, (_this).f21, new BigNumber((_808_v150).length));
          let _837_v173;
          _837_v173 = _dafny.Set.fromElements((_this).f22, ((_this).f22).minus((_836_v172).dtor_cf58));
          _837_v173 = ((_837_v173).Union(_837_v173)).Difference((_837_v173).Difference(_837_v173));
          let _838_v174;
          _838_v174 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_798_v140);
          let _839_v175;
          _839_v175 = _module.D0.create_DC2(((_this).f21) && ((_this).f21), (((_838_v174).contains((_this).f21)) ? ((_838_v174).get((_this).f21)) : (_798_v140)));
          _839_v175 = _module.D0.create_DC2(false, _798_v140);
          let _index126 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
          (_798_v140)[_index126] = (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))];
          r1 = (_this).f21;
          let _840_v176;
          let _nw110 = new _module.C0();
          _nw110.__ctor();
          _840_v176 = _nw110;
        }
        let _841_v177;
        let _nw111 = new _module.C0();
        _nw111.__ctor();
        _841_v177 = _nw111;
      } else if (_source20.is_DC24) {
        let _842___mcc_h25 = (_source20).cf41;
        let _843_cf41 = _842___mcc_h25;
        let _844_v178;
        let _nw112 = Array((new BigNumber(24)).toNumber()).fill([]);
        _844_v178 = _nw112;
        let _845_v179;
        _845_v179 = _dafny.Seq.of(_800_v142);
        let _846_v180;
        let _nw113 = Array((new BigNumber(9)).toNumber());
        _nw113[(_dafny.ZERO).toNumber()] = (_this).f22;
        _nw113[(_dafny.ONE).toNumber()] = (_this).f22;
        _nw113[(new BigNumber(2)).toNumber()] = _843_cf41;
        _nw113[(new BigNumber(3)).toNumber()] = (_this).f22;
        _nw113[(new BigNumber(4)).toNumber()] = (_this).f22;
        _nw113[(new BigNumber(5)).toNumber()] = _843_cf41;
        _nw113[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_dafny.Seq.of((_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))])).length), globalState), (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))], (_this).f21, (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))], (_this).f21)).length);
        _nw113[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_843_cf41);
        _nw113[(new BigNumber(8)).toNumber()] = new BigNumber((_845_v179).length);
        _846_v180 = _nw113;
        let _index127 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_844_v178).length));
        (_844_v178)[_index127] = _846_v180;
        let _index128 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_844_v178).length));
        (_844_v178)[_index128] = _846_v180;
        let _847_v181;
        _847_v181 = _module.D3.create_DC10(_565_v2, !((_this).f21), (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))]);
        let _848_v182;
        _848_v182 = _dafny.Set.fromElements(_dafny.Seq.of(_563_v0), (_847_v181).dtor_cf14);
        r0 = ((_848_v182).Intersect(_dafny.Set.fromElements(_565_v2, _dafny.Seq.update(_565_v2, _module.__default.safeIndex((_this).f22, new BigNumber((_565_v2).length)), _563_v0)))).Union(_848_v182);
        _843_cf41 = new BigNumber((_module.__default.fm3(new BigNumber((_dafny.Seq.Concat(_565_v2, _565_v2)).length), _843_cf41, globalState)).length);
        let _index129 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
        (_798_v140)[_index129] = (_this).f21;
      } else {
        let _849___mcc_h26 = (_source20).cf40;
        let _850_cf40 = _849___mcc_h26;
        let _index130 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
        (_798_v140)[_index130] = !(((_this).f22).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of(_798_v140)).length))) || (!((_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))]) || ((_this).f21));
        let _851_v183;
        _851_v183 = _dafny.Set.fromElements((_this).f22, (_this).f22, (_this).f22);
        let _index131 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
        let _rhs126 = (_808_v150)[_module.__default.safeIndex(new BigNumber((_851_v183).length), new BigNumber((_808_v150).length))];
        let _rhs127 = new BigNumber(169);
        let _rhs128 = (_this).f21;
        let _lhs86 = globalState;
        let _lhs87 = _798_v140;
        let _lhs88 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length));
        r1 = _rhs126;
        _lhs86.f0 = _rhs127;
        _lhs87[_lhs88] = _rhs128;
        let _852_v184;
        let _853_v185;
        let _out32;
        let _out33;
        let _outcollector15 = _module.__default.m0((_this).f22, _798_v140, globalState);
        _out32 = _outcollector15[0];
        _out33 = _outcollector15[1];
        _852_v184 = _out32;
        _853_v185 = _out33;
        let _854_v186;
        let _nw114 = new _module.C1();
        _nw114.__ctor(!((_this).f21));
        _854_v186 = _nw114;
      }
      let _855_v187;
      _855_v187 = _dafny.Set.fromElements(_565_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(206)), function (_856_i24) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }));
      r0 = (_dafny.Set.fromElements(_565_v2, _565_v2, _565_v2)).Difference(_855_v187);
      r1 = (_798_v140)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_798_v140).length))];
      r2 = (_dafny.ZERO).minus(((_this).f22).multipliedBy((new BigNumber(-75)).minus((_this).f22)));
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      let _857_v0;
      _857_v0 = _dafny.Seq.of((_this).f21, (_this).f21, (_this).f21);
      (globalState).f1 = (new BigNumber((_857_v0).length)).minus(_module.__default.fm1(globalState));
      let _858_v1;
      let _nw115 = Array((new BigNumber(24)).toNumber()).fill(false);
      _858_v1 = _nw115;
      _858_v1 = _858_v1;
      let _859_v2;
      _859_v2 = false;
      let _860_v3;
      let _nw116 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _860_v3 = _nw116;
      let _861_v4;
      _861_v4 = _dafny.Set.fromElements(p0);
      let _862_v5;
      _862_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_861_v4).length),(_this).f22);
      let _index132 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_860_v3).length));
      (_860_v3)[_index132] = (((_862_v5).contains(new BigNumber(874))) ? ((_862_v5).get(new BigNumber(874))) : (p0));
      let _index133 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_860_v3).length));
      (_860_v3)[_index133] = p0;
      let _863_v6;
      _863_v6 = _module.D10.create_DC31((_this).f21, _859_v2, (_this).f22);
      let _864_v7;
      _864_v7 = _dafny.Seq.of(_863_v6);
      let _865_v8;
      _865_v8 = _dafny.Seq.UnicodeFromString("j");
      let _866_v9;
      _866_v9 = new _dafny.CodePoint('p'.codePointAt(0));
      let _867_v10;
      _867_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_859_v2);
      let _868_v11;
      _868_v11 = _dafny.Set.fromElements(_867_v10, _867_v10);
      let _869_v12;
      _869_v12 = _module.D12.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(529)), ((_870_v6) => function (_871_i1) {
  return _870_v6;
})(_863_v6)));
      let _872_v13;
      _872_v13 = _dafny.Seq.of(_dafny.Seq.Concat((_869_v12).dtor_cf59, _864_v7), _dafny.Seq.Concat(_864_v7, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), ((_873_v6) => function (_874_i2) {
        return _873_v6;
      })(_863_v6)), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), ((_875_v6) => function (_876_i2) {
        return _875_v6;
      })(_863_v6))).length)), _863_v6)));
      let _877_v14;
      _877_v14 = _dafny.MultiSet.fromElements((_this).f22);
      let _index134 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_860_v3).length));
      let _index135 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_860_v3).length));
      let _rhs129 = _859_v2;
      let _rhs130 = (_this).f22;
      let _rhs131 = (_this).fm9(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_865_v8, _module.__default.safeIndex(new BigNumber(465), new BigNumber((_865_v8).length)), new _dafny.CodePoint('m'.codePointAt(0))), _865_v8), _module.__default.safeIndex(new BigNumber((_module.__default.fm38(false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-606)), function (_878_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length), globalState)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_865_v8, _module.__default.safeIndex(new BigNumber(465), new BigNumber((_865_v8).length)), new _dafny.CodePoint('m'.codePointAt(0))), _865_v8)).length)), _866_v9), (_868_v11).Difference(_dafny.Set.fromElements(_867_v10)), globalState);
      let _rhs132 = (_872_v13)[_module.__default.safeIndex(_module.__default.safeModuloInt(_module.__default.fm1(globalState), (((_877_v14).contains((_this).f22)) ? ((_877_v14).get((_this).f22)) : (p0))), new BigNumber((_872_v13).length))];
      let _lhs89 = _860_v3;
      let _lhs90 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_860_v3).length));
      let _lhs91 = _860_v3;
      let _lhs92 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_860_v3).length));
      _859_v2 = _rhs129;
      _lhs89[_lhs90] = _rhs130;
      _lhs91[_lhs92] = _rhs131;
      _864_v7 = _rhs132;
      let _879_v15;
      _879_v15 = _dafny.MultiSet.fromElements(_865_v8);
      let _rhs133 = !((_879_v15).IsProperSubsetOf(_879_v15));
      let _rhs134 = _module.__default.fm4((((_this).f21) ? (_859_v2) : ((_this).f21)), (_this).f21, (_this).f21, globalState);
      _859_v2 = _rhs133;
      _866_v9 = _rhs134;
      _858_v1 = _858_v1;
      let _880_v16;
      _880_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_859_v2));
      let _881_v17;
      _881_v17 = _dafny.MultiSet.fromElements(_880_v16);
      (globalState).f1 = (((_862_v5).contains(new BigNumber((_881_v17).cardinality()))) ? ((_862_v5).get(new BigNumber((_881_v17).cardinality()))) : ((_860_v3)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_860_v3).length))]));
      return;
    }
    m15(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let _882_v0;
      _882_v0 = _dafny.Seq.of(p1);
      let _hi2 = (_this).f22;
      for (let _883_i0 = new BigNumber((_882_v0).length); _883_i0.isLessThan(_hi2); _883_i0 = _883_i0.plus(_dafny.ONE)) {
        let _884_v1;
        _884_v1 = _dafny.Seq.of((_dafny.ZERO).minus(_883_i0));
        let _885_v2;
        _885_v2 = _dafny.Set.fromElements(p2);
        let _886_v3;
        let _nw117 = Array((new BigNumber(10)).toNumber()).fill(false);
        _886_v3 = _nw117;
        let _887_v4;
        let _888_v5;
        let _out34;
        let _out35;
        let _outcollector16 = _module.__default.m0(new BigNumber((_dafny.Seq.Concat(_884_v1, _module.__default.fm3((_this).f22, new BigNumber((_885_v2).length), globalState))).length), _886_v3, globalState);
        _out34 = _outcollector16[0];
        _out35 = _outcollector16[1];
        _887_v4 = _out34;
        _888_v5 = _out35;
        let _index136 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_886_v3).length));
        (_886_v3)[_index136] = p1;
        let _index137 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_886_v3).length));
        let _rhs135 = (_882_v0)[_module.__default.safeIndex((_883_i0).plus(_883_i0), new BigNumber((_882_v0).length))];
        let _rhs136 = (_this).f22;
        let _lhs93 = _886_v3;
        let _lhs94 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_886_v3).length));
        let _lhs95 = globalState;
        _lhs93[_lhs94] = _rhs135;
        _lhs95.f7 = _rhs136;
        let _889_v6;
        _889_v6 = _dafny.MultiSet.fromElements(_883_i0, new BigNumber((_dafny.MultiSet.fromElements(p2, p2, (_this).f21, p1)).cardinality()), (_dafny.ZERO).minus(_883_i0));
        _889_v6 = (_889_v6).Intersect((_889_v6).Union(_889_v6));
        let _890_v7;
        _890_v7 = _dafny.Map.Empty.slice().updateUnsafe(_883_i0,_883_i0);
        _890_v7 = _890_v7;
      }
      let _891_v8;
      _891_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm40(p2, (_this).f21, globalState),p1);
      let _892_v9;
      _892_v9 = _module.D4.create_DC13((_891_v8).Merge(_891_v8));
      _892_v9 = _892_v9;
      let _893_v10;
      let _init19 = ((_894_p1) => function (_895_i1) {
        return _894_p1;
      })(p1);
      let _nw118 = Array((new BigNumber(23)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw118.length); _i0_19++) {
        _nw118[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _893_v10 = _nw118;
      let _896_v11;
      _896_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f22);
      let _897_v12;
      _897_v12 = _dafny.Seq.UnicodeFromString("n");
      let _898_v13;
      _898_v13 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("bv"));
      let _899_v14;
      _899_v14 = _dafny.Set.fromElements(_897_v12, (_898_v13)[_module.__default.safeIndex((_this).f22, new BigNumber((_898_v13).length))]);
      let _900_v15;
      _900_v15 = _dafny.MultiSet.fromElements((_this).f22);
      let _901_v16;
      let _nw119 = Array((new BigNumber(11)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = (_this).f22;
      _nw119[(_dafny.ONE).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(2)).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(3)).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(4)).toNumber()] = (new BigNumber(806)).minus((((_896_v11).contains(p1)) ? ((_896_v11).get(p1)) : (new BigNumber((_899_v14).length))));
      _nw119[(new BigNumber(5)).toNumber()] = new BigNumber((_900_v15).cardinality());
      _nw119[(new BigNumber(6)).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(7)).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(8)).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(9)).toNumber()] = (_this).f22;
      _nw119[(new BigNumber(10)).toNumber()] = (_this).f22;
      _901_v16 = _nw119;
      let _902_v17;
      _902_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
      let _903_v18;
      _903_v18 = _dafny.Set.fromElements((_902_v17).update(p2, (_this).f21), _902_v17, (_902_v17).update(p2, p1), _902_v17);
      let _index138 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length));
      (_901_v16)[_index138] = ((_this).f22).plus((_this).fm9(_897_v12, _903_v18, globalState));
      let _904_v19;
      _904_v19 = false;
      let _905_v20;
      _905_v20 = _dafny.MultiSet.fromElements(p3, _module.__default.fm4(p2, (_this).f21, (_this).f21, globalState), p3);
      let _906_v21;
      _906_v21 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_897_v12), _905_v20, _905_v20);
      let _907_v22;
      _907_v22 = _dafny.MultiSet.fromElements(_901_v16, _901_v16, _901_v16);
      let _index139 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length));
      let _rhs137 = _893_v10;
      let _rhs138 = (_module.__default.fm1(globalState)).multipliedBy((_this).f22);
      let _rhs139 = !(((_906_v21)[_module.__default.safeIndex((_this).f22, new BigNumber((_906_v21).length))]).IsSubsetOf(_905_v20));
      let _rhs140 = (_907_v22).IsSubsetOf(_907_v22);
      let _lhs96 = _901_v16;
      let _lhs97 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length));
      _893_v10 = _rhs137;
      _lhs96[_lhs97] = _rhs138;
      _904_v19 = _rhs139;
      _904_v19 = _rhs140;
      let _908_v23;
      _908_v23 = _dafny.MultiSet.fromElements(_893_v10);
      let _909_v24;
      _909_v24 = _module.D13.create_DC38(_908_v23);
      let _910_v25;
      _910_v25 = _dafny.Map.Empty.slice().updateUnsafe(_904_v19,_908_v23);
      let _911_v26;
      let _nw120 = Array((new BigNumber(22)).toNumber());
      _nw120[(_dafny.ZERO).toNumber()] = ((_909_v24).dtor_cf64).Union((((_910_v25).contains(_904_v19)) ? ((_910_v25).get(_904_v19)) : (_908_v23)));
      _nw120[(_dafny.ONE).toNumber()] = _908_v23;
      _nw120[(new BigNumber(2)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_893_v10);
      _nw120[(new BigNumber(4)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(5)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(6)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(7)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(8)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(9)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(10)).toNumber()] = ((_909_v24).dtor_cf64).Intersect((_dafny.MultiSet.fromElements(_893_v10)).update(_893_v10, _module.__default.abs((_dafny.ZERO).minus((_this).f22))));
      _nw120[(new BigNumber(11)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(12)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_893_v10, _893_v10, _893_v10);
      _nw120[(new BigNumber(14)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(15)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(16)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(17)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(18)).toNumber()] = (_908_v23).update(_893_v10, _module.__default.abs((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))]));
      _nw120[(new BigNumber(19)).toNumber()] = (_908_v23).Difference(_dafny.MultiSet.fromElements(_893_v10));
      _nw120[(new BigNumber(20)).toNumber()] = _908_v23;
      _nw120[(new BigNumber(21)).toNumber()] = _908_v23;
      _911_v26 = _nw120;
      let _index140 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_911_v26).length));
      (_911_v26)[_index140] = _dafny.MultiSet.fromElements(_893_v10);
      let _index141 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_911_v26).length));
      (_911_v26)[_index141] = _908_v23;
      let _912_v27;
      _912_v27 = _dafny.Seq.of((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))], new BigNumber(-928));
      let _913_v28;
      _913_v28 = _dafny.Seq.of(new BigNumber((_902_v17).length), new BigNumber((_912_v27).length));
      _904_v19 = ((_913_v28)[_module.__default.safeIndex((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))], new BigNumber((_913_v28).length))]).isLessThan(new BigNumber(65));
      let _hi3 = (_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))];
      for (let _914_i2 = _module.__default.safeModuloInt(new BigNumber(25), (_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))]); _914_i2.isLessThan(_hi3); _914_i2 = _914_i2.plus(_dafny.ONE)) {
        let _915_v29;
        _915_v29 = _dafny.Map.Empty.slice().updateUnsafe(_914_i2,_899_v14);
        let _rhs141 = _module.__default.fm40((_this).f21, (_899_v14).IsSubsetOf((((_915_v29).contains((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))])) ? ((_915_v29).get((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))])) : (_dafny.Set.fromElements(_897_v12, _897_v12)))), globalState);
        let _rhs142 = !(_dafny.Seq.contains(((!(p1)) ? (_897_v12) : (_dafny.Seq.UnicodeFromString("co"))), p3));
        _882_v0 = _rhs141;
        _904_v19 = _rhs142;
        let _916_v30;
        _916_v30 = _module.D6.create_DC21((_this).f21, _904_v19, (_this).f21, (((_900_v15).contains(new BigNumber(262))) ? ((_900_v15).get(new BigNumber(262))) : (_914_i2)));
        _904_v19 = (((_916_v30).dtor_cf39).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(417)), function (_917_i3) {
          return (_this).f22;
        })).length))).isLessThan((_this).f22);
        let _918_v31;
        _918_v31 = _module.D3.create_DC10(_dafny.Seq.update(_897_v12, _module.__default.safeIndex((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))], new BigNumber((_897_v12).length)), p3), p2, true);
        let _919_v32;
        _919_v32 = _module.D13.create_DC39((((_896_v11).contains(!(true))) ? ((_896_v11).get(!(true))) : ((_this).f22)), _904_v19);
        let _920_v33;
        let _nw121 = Array((new BigNumber(28)).toNumber());
        _nw121[(_dafny.ZERO).toNumber()] = _897_v12;
        _nw121[(_dafny.ONE).toNumber()] = _897_v12;
        _nw121[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_897_v12, _897_v12);
        _nw121[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_897_v12, _897_v12);
        _nw121[(new BigNumber(4)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(5)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(6)).toNumber()] = (_898_v13)[_module.__default.safeIndex((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))], new BigNumber((_898_v13).length))];
        _nw121[(new BigNumber(7)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(8)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm38((_this).f21, new BigNumber((_897_v12).length), globalState), _dafny.Seq.UnicodeFromString("gotocugl"));
        _nw121[(new BigNumber(10)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(11)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_897_v12, _897_v12);
        _nw121[(new BigNumber(13)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(14)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(15)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(16)).toNumber()] = (_918_v31).dtor_cf14;
        _nw121[(new BigNumber(17)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(77)), ((_921_p3) => function (_922_i4) {
          return _921_p3;
        })(p3)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), ((_923_p3) => function (_924_i5) {
          return _923_p3;
        })(p3)));
        _nw121[(new BigNumber(19)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(20)).toNumber()] = ((true) ? (_897_v12) : (_897_v12));
        _nw121[(new BigNumber(21)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(22)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(23)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_897_v12, (_898_v13)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_898_v13).length))]);
        _nw121[(new BigNumber(25)).toNumber()] = _897_v12;
        _nw121[(new BigNumber(26)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_897_v12, _897_v12), _module.__default.safeIndex((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))], new BigNumber((_dafny.Seq.Concat(_897_v12, _897_v12)).length)), _module.__default.fm4(p1, (_this).f21, (_this).f21, globalState));
        _nw121[(new BigNumber(27)).toNumber()] = _dafny.Seq.update(_897_v12, _module.__default.safeIndex((_919_v32).dtor_cf65, new BigNumber((_897_v12).length)), p3);
        _920_v33 = _nw121;
        let _index142 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_920_v33).length));
        (_920_v33)[_index142] = _dafny.Seq.UnicodeFromString("mu");
        let _index143 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_920_v33).length));
        (_920_v33)[_index143] = ((p1) ? (_dafny.Seq.Concat(_897_v12, _dafny.Seq.UnicodeFromString("ussllkr"))) : (_897_v12));
        let _925_v34;
        _925_v34 = _dafny.Set.fromElements((_this).f22);
        let _926_v35;
        _926_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_925_v34).length),(_this).f22);
        let _927_v36;
        _927_v36 = _module.D11.create_DC34((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))], p1, (((_926_v35).contains((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))])) ? ((_926_v35).get((_901_v16)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_901_v16).length))])) : (new BigNumber((_925_v34).length))));
        _927_v36 = _927_v36;
      }
      r0 = _901_v16;
      return r0;
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f10 = false;
      this._f20 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f20, f10) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f10 = f10;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f20;
    };
    fm7(p0, globalState) {
      let _this = this;
      return !((_this).f10);
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f20;
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pmomwxle"), _dafny.Seq.UnicodeFromString("nd")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lltikvrj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), function (_928_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })))).length);
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Map.Empty;
      let _929_v0;
      _929_v0 = _dafny.Set.fromElements(false, true);
      let _930_v1;
      _930_v1 = new _dafny.CodePoint('l'.codePointAt(0));
      let _931_v2;
      _931_v2 = _module.D5.create_DC17((_929_v0).IsSubsetOf(_dafny.Set.fromElements((_this).f10)), _930_v1);
      let _source21 = _931_v2;
      if (_source21.is_DC17) {
        let _932___mcc_h0 = (_source21).cf28;
        let _933___mcc_h1 = (_source21).cf29;
        let _934_cf29 = _933___mcc_h1;
        let _935_cf28 = _932___mcc_h0;
        let _936_v3;
        _936_v3 = new BigNumber(525);
        (globalState).f7 = _module.__default.safeDivisionInt(_936_v3, _936_v3);
        let _937_v4;
        let _nw122 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _937_v4 = _nw122;
        _937_v4 = _937_v4;
        r0 = p1;
        let _938_v5;
        _938_v5 = _dafny.Seq.of((_this).f10);
        let _index144 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_937_v4).length));
        (_937_v4)[_index144] = (_dafny.ZERO).minus((((_938_v5)[_module.__default.safeIndex(_936_v3, new BigNumber((_938_v5).length))]) ? (_936_v3) : (_936_v3)));
        let _index145 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_937_v4).length));
        (_937_v4)[_index145] = _936_v3;
      } else if (_source21.is_DC16) {
        let _939___mcc_h2 = (_source21).cf27;
        let _940_cf27 = _939___mcc_h2;
        let _941_v6;
        let _nw123 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _941_v6 = _nw123;
        let _942_v7;
        _942_v7 = _dafny.Seq.UnicodeFromString("ifnhqam");
        let _index146 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_941_v6).length));
        (_941_v6)[_index146] = _942_v7;
        let _943_v8;
        _943_v8 = new BigNumber(967);
        let _index147 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_941_v6).length));
        let _rhs143 = new BigNumber(892);
        let _rhs144 = _dafny.Seq.update(_942_v7, _module.__default.safeIndex(_943_v8, new BigNumber((_942_v7).length)), _930_v1);
        let _lhs98 = globalState;
        let _lhs99 = _941_v6;
        let _lhs100 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_941_v6).length));
        _lhs98.f1 = _rhs143;
        _lhs99[_lhs100] = _rhs144;
        let _944_v9;
        _944_v9 = _dafny.Seq.of((_this).f10, (_this).f10, false);
        let _945_v10;
        let _nw124 = Array((new BigNumber(19)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = _944_v9;
        _nw124[(_dafny.ONE).toNumber()] = _944_v9;
        _nw124[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_this).f20, (_this).f20, (_this).f20, p1, true);
        _nw124[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_944_v9, _module.__default.safeIndex(_943_v8, new BigNumber((_944_v9).length)), (_this).f10);
        _nw124[(new BigNumber(4)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(p1);
        _nw124[(new BigNumber(6)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(7)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_944_v9, _944_v9);
        _nw124[(new BigNumber(9)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(10)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(11)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(12)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_944_v9, _dafny.Seq.of((_this).f10));
        _nw124[(new BigNumber(14)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(15)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(16)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(17)).toNumber()] = _944_v9;
        _nw124[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_944_v9, _944_v9);
        _945_v10 = _nw124;
        _945_v10 = _945_v10;
        let _946_v11;
        let _init20 = function (_947_i0) {
          return !((_this).f10);
        };
        let _nw125 = Array((new BigNumber(24)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw125.length); _i0_20++) {
          _nw125[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _946_v11 = _nw125;
        let _index148 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_946_v11).length));
        (_946_v11)[_index148] = (_this).f20;
        let _948_v12;
        _948_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(23),!((_this).f20));
        let _index149 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_946_v11).length));
        (_946_v11)[_index149] = (_this).fm6(_943_v8, _module.__default.fm1(globalState), _943_v8, (((_948_v12).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(480)), ((_951_v1) => function (_952_i1) {
          return _951_v1;
        })(_930_v1))).length))) ? ((_948_v12).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(480)), ((_949_v1) => function (_950_i1) {
          return _949_v1;
        })(_930_v1))).length))) : ((_this).f10)), globalState);
        let _953_v13;
        _953_v13 = _dafny.MultiSet.fromElements(_943_v8);
        let _954_v14;
        _954_v14 = _module.D4.create_DC15(_943_v8, _953_v13, _930_v1, _943_v8);
        let _955_v15;
        let _init21 = ((_956_v8) => function (_957_i2) {
          return _module.__default.safeDivisionInt(_957_i2, (_dafny.ZERO).minus(_956_v8));
        })(_943_v8);
        let _nw126 = Array((new BigNumber(21)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw126.length); _i0_21++) {
          _nw126[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _955_v15 = _nw126;
        let _index150 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_955_v15).length));
        (_955_v15)[_index150] = _943_v8;
        let _958_v16;
        _958_v16 = _dafny.Set.fromElements(new BigNumber(508));
        let _pat_let_tv18 = _953_v13;
        let _pat_let_tv19 = _953_v13;
        let _pat_let_tv20 = _943_v8;
        let _index151 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_955_v15).length));
        let _rhs145 = new BigNumber((_958_v16).length);
        let _rhs146 = function (_pat_let16_0) {
          return function (_959_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_960_dt__update_hcf24_h0) {
                return function (_pat_let18_0) {
                  return function (_961_dt__update_hcf23_h0) {
                    return _module.D4.create_DC15(_961_dt__update_hcf23_h0, _960_dt__update_hcf24_h0, (_959_dt__update__tmp_h0).dtor_cf25, (_959_dt__update__tmp_h0).dtor_cf26);
                  }(_pat_let18_0);
                }(_pat_let_tv20);
              }(_pat_let17_0);
            }((_pat_let_tv18).Union(_pat_let_tv19));
          }(_pat_let16_0);
        }(_module.D4.create_DC15(new BigNumber((_958_v16).length), _dafny.MultiSet.fromElements(_943_v8), _930_v1, _943_v8));
        let _rhs147 = new BigNumber(((_941_v6)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_941_v6).length))]).length);
        let _rhs148 = _943_v8;
        let _lhs101 = _955_v15;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_955_v15).length));
        let _lhs103 = globalState;
        _943_v8 = _rhs145;
        _954_v14 = _rhs146;
        _lhs101[_lhs102] = _rhs147;
        _lhs103.f1 = _rhs148;
      } else {
        let _962___mcc_h3 = (_source21).cf30;
        let _963_cf30 = _962___mcc_h3;
        let _964_v17;
        let _nw127 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _964_v17 = _nw127;
        let _965_v18;
        _965_v18 = new BigNumber(590);
        let _index152 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_964_v17).length));
        (_964_v17)[_index152] = _module.__default.safeDivisionInt(_965_v18, _965_v18);
        let _index153 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_964_v17).length));
        let _rhs149 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_965_v18, (_965_v18).minus(_965_v18)));
        let _rhs150 = !((_this).f20) || ((_this).f10);
        let _rhs151 = !(false) || ((_this).f20);
        let _rhs152 = _930_v1;
        let _rhs153 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(_965_v18))).plus(_965_v18);
        let _lhs104 = _964_v17;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_964_v17).length));
        let _lhs106 = globalState;
        _lhs104[_lhs105] = _rhs149;
        r0 = _rhs150;
        r0 = _rhs151;
        _930_v1 = _rhs152;
        _lhs106.f0 = _rhs153;
        _930_v1 = _930_v1;
        r0 = false;
        let _966_v19;
        _966_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,_965_v18);
        let _967_v20;
        _967_v20 = _dafny.MultiSet.fromElements((((_966_v19).contains((_this).f20)) ? ((_966_v19).get((_this).f20)) : (_965_v18)), _965_v18);
        r0 = ((((_dafny.ZERO).minus((_964_v17)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_964_v17).length))])).isLessThan(new BigNumber((_967_v20).cardinality()))) ? ((_this).f20) : (false));
      }
      let _968_v21;
      _968_v21 = new BigNumber(-136);
      let _969_v22;
      _969_v22 = _dafny.Map.Empty.slice().updateUnsafe((_968_v21).isLessThanOrEqualTo(_968_v21),(new BigNumber(699)).minus(_module.__default.fm1(globalState)));
      let _970_v23;
      _970_v23 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.Seq.UnicodeFromString("jjtarsvry")).length)), _969_v22);
      let _971_v24;
      _971_v24 = _dafny.Seq.UnicodeFromString("wjgi");
      _969_v22 = (_969_v22).update((_this).f10, new BigNumber(((_970_v23)[_module.__default.safeIndex(new BigNumber((_971_v24).length), new BigNumber((_970_v23).length))]).length));
      let _972_v25;
      let _nw128 = Array((new BigNumber(29)).toNumber()).fill([]);
      _972_v25 = _nw128;
      let _973_v26;
      _973_v26 = _module.D6.create_DC20((_this).f20, false, (_this).f20, (_dafny.ZERO).minus(_968_v21));
      let _974_v27;
      let _nw129 = Array((new BigNumber(21)).toNumber());
      _nw129[(_dafny.ZERO).toNumber()] = p1;
      _nw129[(_dafny.ONE).toNumber()] = p1;
      _nw129[(new BigNumber(2)).toNumber()] = p1;
      _nw129[(new BigNumber(3)).toNumber()] = false;
      _nw129[(new BigNumber(4)).toNumber()] = false;
      _nw129[(new BigNumber(5)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(6)).toNumber()] = (_this).f10;
      _nw129[(new BigNumber(7)).toNumber()] = false;
      _nw129[(new BigNumber(8)).toNumber()] = (_973_v26).dtor_cf32;
      _nw129[(new BigNumber(9)).toNumber()] = true;
      _nw129[(new BigNumber(10)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(11)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(12)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(13)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(14)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(15)).toNumber()] = (_this).f10;
      _nw129[(new BigNumber(16)).toNumber()] = (_this).f10;
      _nw129[(new BigNumber(17)).toNumber()] = (_this).f20;
      _nw129[(new BigNumber(18)).toNumber()] = true;
      _nw129[(new BigNumber(19)).toNumber()] = (_this).f10;
      _nw129[(new BigNumber(20)).toNumber()] = true;
      _974_v27 = _nw129;
      let _index154 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_972_v25).length));
      (_972_v25)[_index154] = _974_v27;
      let _index155 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_972_v25).length));
      (_972_v25)[_index155] = _974_v27;
      r0 = !(p1) || (p1);
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber(((_972_v25)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_972_v25).length))]).length))) {
        let _975_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_975_i3)) && ((_975_i3).isLessThan(new BigNumber(((_972_v25)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_972_v25).length))]).length))))) {
          _ingredients0.push(_dafny.Tuple.of((_972_v25)[_module.__default.safeIndex(new BigNumber(810), new BigNumber((_972_v25).length))], (_975_i3).toNumber(), (_this).f20));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _976_v28;
      let _init22 = ((_977_v21) => function (_978_i4) {
        return _module.__default.safeDivisionInt(_978_i4, _977_v21);
      })(_968_v21);
      let _nw130 = Array((new BigNumber(5)).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw130.length); _i0_22++) {
        _nw130[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _976_v28 = _nw130;
      let _979_v29;
      _979_v29 = _dafny.Set.fromElements(_976_v28, ((p1) ? (_976_v28) : (_976_v28)));
      _979_v29 = _979_v29;
      r0 = !(_dafny.areEqual(_dafny.Seq.UnicodeFromString("metlvgmn"), _dafny.Seq.UnicodeFromString("pbiudyvt")));
      let _980_v30;
      _980_v30 = _dafny.Seq.of(_971_v24, _971_v24, _971_v24);
      r1 = (_980_v30)[_module.__default.safeIndex(_968_v21, new BigNumber((_980_v30).length))];
      let _981_v31;
      _981_v31 = _dafny.Map.Empty.slice().updateUnsafe(_968_v21,_968_v21);
      let _982_v32;
      _982_v32 = _module.D7.create_DC22(_981_v31);
      r2 = (_982_v32).dtor_cf40;
      return [r0, r1, r2];
    }
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _983_v0;
      _983_v0 = _dafny.Seq.UnicodeFromString("vetxc");
      let _984_v1;
      _984_v1 = new BigNumber(208);
      let _985_v2;
      _985_v2 = new _dafny.CodePoint('l'.codePointAt(0));
      let _986_v3;
      _986_v3 = _dafny.MultiSet.fromElements((_this).f10, true);
      let _987_v4;
      _987_v4 = _dafny.Seq.of((_this).f10);
      let _988_v5;
      _988_v5 = _dafny.Map.Empty.slice().updateUnsafe(!((_987_v4)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_987_v4).length))]),_986_v3);
      if ((new BigNumber(((_986_v3).Union((((_988_v5).contains((_this).f20)) ? ((_988_v5).get((_this).f20)) : (_986_v3)))).cardinality())).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_983_v0, _983_v0), _module.__default.safeIndex(_984_v1, new BigNumber((_dafny.Seq.Concat(_983_v0, _983_v0)).length)), _985_v2)).length))) {
        let _989_v6;
        _989_v6 = _dafny.Map.Empty.slice().updateUnsafe(false,_984_v1);
        _984_v1 = (((_989_v6).contains((_this).f10)) ? ((_989_v6).get((_this).f10)) : (new BigNumber((_dafny.Seq.Concat(_983_v0, _dafny.Seq.UnicodeFromString("qssp"))).length)));
        r1 = (_this).fm6(_984_v1, _984_v1, (_984_v1).multipliedBy(_984_v1), (_this).f20, globalState);
        let _990_v7;
        let _init23 = function (_991_i0) {
          return (_this).f10;
        };
        let _nw131 = Array((new BigNumber(26)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw131.length); _i0_23++) {
          _nw131[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _990_v7 = _nw131;
        _990_v7 = _990_v7;
        if (!((_this).f10)) {
          let _992_v9;
          let _init24 = ((_993_v1) => function (_994_i1) {
            return function () {
              let _coll25 = new _dafny.Map();
              for (const _compr_25 of _dafny.IntegerRange(new BigNumber(763), new BigNumber(-353))) {
                let _995_v8 = _compr_25;
                if (((new BigNumber(763)).isLessThanOrEqualTo(_995_v8)) && ((_995_v8).isLessThan(new BigNumber(-353)))) {
                  _coll25.push([(_995_v8).multipliedBy(_993_v1),(_this).f20]);
                }
              }
              return _coll25;
            }();
          })(_984_v1);
          let _nw132 = Array((new BigNumber(19)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw132.length); _i0_24++) {
            _nw132[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _992_v9 = _nw132;
          let _996_v10;
          let _nw133 = Array((new BigNumber(24)).toNumber());
          _nw133[(_dafny.ZERO).toNumber()] = _992_v9;
          _nw133[(_dafny.ONE).toNumber()] = _992_v9;
          _nw133[(new BigNumber(2)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(3)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(4)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(5)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(6)).toNumber()] = (((_this).f20) ? (_992_v9) : (_992_v9));
          _nw133[(new BigNumber(7)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(8)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(9)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(10)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(11)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(12)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(13)).toNumber()] = (((_this).f10) ? (_992_v9) : (_992_v9));
          _nw133[(new BigNumber(14)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(15)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(16)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(17)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(18)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(19)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(20)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(21)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(22)).toNumber()] = _992_v9;
          _nw133[(new BigNumber(23)).toNumber()] = _992_v9;
          _996_v10 = _nw133;
          let _index156 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_996_v10).length));
          (_996_v10)[_index156] = _992_v9;
          let _index157 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_996_v10).length));
          (_996_v10)[_index157] = _992_v9;
          let _997_v11;
          let _nw134 = new _module.C2();
          _nw134.__ctor((_this).f20, _984_v1);
          _997_v11 = _nw134;
          let _998_v12;
          _998_v12 = _dafny.Map.Empty.slice().updateUnsafe(_997_v11,_983_v0);
          let _999_v13;
          _999_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_998_v12).update(_997_v11, _dafny.Seq.UnicodeFromString("sywu")));
          let _1000_v14;
          let _nw135 = Array((new BigNumber(4)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = (((_999_v13).contains((_this).f10)) ? ((_999_v13).get((_this).f10)) : (_998_v12));
          _nw135[(_dafny.ONE).toNumber()] = _998_v12;
          _nw135[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_997_v11,_983_v0)).Merge(_998_v12);
          _nw135[(new BigNumber(3)).toNumber()] = _998_v12;
          _1000_v14 = _nw135;
          let _index158 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_1000_v14).length));
          (_1000_v14)[_index158] = _998_v12;
          let _1001_v15;
          _1001_v15 = _dafny.Seq.of(_998_v12, _998_v12);
          let _index159 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_1000_v14).length));
          (_1000_v14)[_index159] = ((_998_v12).Merge(_998_v12)).Merge((_1001_v15)[_module.__default.safeIndex(_984_v1, new BigNumber((_1001_v15).length))]);
          (globalState).f1 = _984_v1;
          let _1002_v16;
          _1002_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_983_v0);
          let _1003_v17;
          _1003_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1002_v16).length),_984_v1);
          let _1004_v18;
          _1004_v18 = _dafny.Seq.of((_1003_v17).Merge((_1003_v17).update(_984_v1, _984_v1)));
          _1004_v18 = _dafny.Seq.of(_1003_v17);
          let _1005_v19;
          let _nw136 = new _module.C2();
          _nw136.__ctor((_987_v4)[_module.__default.safeIndex(_984_v1, new BigNumber((_987_v4).length))], _984_v1);
          _1005_v19 = _nw136;
          _1005_v19 = _1005_v19;
        } else {
          let _1006_v20;
          let _nw137 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _1006_v20 = _nw137;
          let _1007_v21;
          _1007_v21 = _module.D4.create_DC14(new BigNumber(-209), new BigNumber(678));
          let _1008_v22;
          _1008_v22 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,_1007_v21);
          let _1009_v23;
          _1009_v23 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,_1008_v22);
          let _index160 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_1006_v20).length));
          (_1006_v20)[_index160] = (((_1009_v23).contains(_984_v1)) ? ((_1009_v23).get(_984_v1)) : (_1008_v22));
          let _index161 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_1006_v20).length));
          (_1006_v20)[_index161] = _1008_v22;
          let _1010_v24;
          let _init25 = ((_1011_v1) => function (_1012_i2) {
            return _module.D7.create_DC24(_1011_v1);
          })(_984_v1);
          let _nw138 = Array((new BigNumber(26)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw138.length); _i0_25++) {
            _nw138[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1010_v24 = _nw138;
          let _1013_v25;
          _1013_v25 = _module.D7.create_DC24(new BigNumber((_983_v0).length));
          let _index162 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_1010_v24).length));
          (_1010_v24)[_index162] = _1013_v25;
          let _pat_let_tv21 = _984_v1;
          let _index163 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_1010_v24).length));
          (_1010_v24)[_index163] = function (_pat_let19_0) {
            return function (_1014_dt__update__tmp_h0) {
              return function (_pat_let20_0) {
                return function (_1015_dt__update_hcf41_h0) {
                  return _module.D7.create_DC24(_1015_dt__update_hcf41_h0);
                }(_pat_let20_0);
              }(_pat_let_tv21);
            }(_pat_let19_0);
          }(_1013_v25);
          let _1016_v26;
          let _nw139 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1016_v26 = _nw139;
          let _1017_v27;
          _1017_v27 = _dafny.MultiSet.fromElements(_983_v0, _983_v0);
          let _index164 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1016_v26).length));
          (_1016_v26)[_index164] = _1017_v27;
          let _1018_v28;
          _1018_v28 = _module.D5.create_DC16(_dafny.MultiSet.fromElements((_this).f10, (_this).f20, (_this).f20, (_this).f10));
          let _1019_v29;
          _1019_v29 = _dafny.Set.fromElements(_1018_v28);
          let _index165 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1016_v26).length));
          let _rhs154 = new BigNumber(544);
          let _rhs155 = _1017_v27;
          let _rhs156 = _985_v2;
          let _rhs157 = _984_v1;
          let _rhs158 = _module.__default.fm4((_1019_v29).IsDisjointFrom(_1019_v29), (_this).f20, (_this).f10, globalState);
          let _lhs107 = _1016_v26;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1016_v26).length));
          let _lhs109 = globalState;
          _984_v1 = _rhs154;
          _lhs107[_lhs108] = _rhs155;
          _985_v2 = _rhs156;
          _lhs109.f0 = _rhs157;
          _985_v2 = _rhs158;
          let _1020_v30;
          let _nw140 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1020_v30 = _nw140;
          let _index166 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1020_v30).length));
          (_1020_v30)[_index166] = _984_v1;
          let _1021_v31;
          _1021_v31 = _dafny.Seq.of(_984_v1);
          let _1022_v32;
          _1022_v32 = _dafny.Map.Empty.slice().updateUnsafe(_983_v0,new BigNumber((_1021_v31).length));
          let _index167 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1020_v30).length));
          (_1020_v30)[_index167] = new BigNumber(((_1022_v32).Merge((_module.__default.fm44(_984_v1, globalState)).Merge(_1022_v32))).length);
          r1 = (_this).fm6((_1020_v30)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1020_v30).length))], (_1020_v30)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1020_v30).length))], (((_1017_v27).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), ((_1025_v2) => function (_1026_i3) {
            return _1025_v2;
          })(_985_v2)))) ? ((_1017_v27).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), ((_1023_v2) => function (_1024_i3) {
            return _1023_v2;
          })(_985_v2)))) : (_984_v1)), (_this).f10, globalState);
        }
        r1 = (_this).f20;
      } else {
        let _1027_v33;
        _1027_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_983_v0);
        _1027_v33 = (_1027_v33).update((_this).f10, _dafny.Seq.UnicodeFromString("qsvlui"));
        let _1028_v34;
        _1028_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-146),_984_v1);
        let _1029_v35;
        _1029_v35 = _dafny.Seq.of(_984_v1);
        let _1030_v36;
        _1030_v36 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,(_this).f20);
        let _rhs159 = (_this).f20;
        let _rhs160 = (_module.__default.fm45(_dafny.Seq.of(_1029_v35), _984_v1, globalState)).update(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_984_v1,(_this).f20)).Merge(_1030_v36)).length), (_984_v1).multipliedBy(_984_v1));
        r1 = _rhs159;
        _1028_v34 = _rhs160;
        r2 = ((_984_v1).multipliedBy(_984_v1)).multipliedBy(_module.__default.fm1(globalState));
        _1029_v35 = _1029_v35;
        let _1031_v37;
        _1031_v37 = _dafny.Seq.of(_983_v0);
        let _1032_v38;
        _1032_v38 = _dafny.MultiSet.fromElements(new BigNumber(864), (_dafny.ZERO).minus((_1029_v35)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(359)), ((_1033_v2) => function (_1034_i4) {
          return _1033_v2;
        })(_985_v2))).length), new BigNumber((_1029_v35).length))]), _984_v1, new BigNumber((_1031_v37).length), _984_v1);
        let _1035_v39;
        let _nw141 = Array((new BigNumber(12)).toNumber());
        _nw141[(_dafny.ZERO).toNumber()] = _984_v1;
        _nw141[(_dafny.ONE).toNumber()] = _984_v1;
        _nw141[(new BigNumber(2)).toNumber()] = (new BigNumber((_1032_v38).cardinality())).multipliedBy(_984_v1);
        _nw141[(new BigNumber(3)).toNumber()] = _984_v1;
        _nw141[(new BigNumber(4)).toNumber()] = (_984_v1).multipliedBy(_984_v1);
        _nw141[(new BigNumber(5)).toNumber()] = (new BigNumber(293)).multipliedBy(_984_v1);
        _nw141[(new BigNumber(6)).toNumber()] = _984_v1;
        _nw141[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm1(globalState));
        _nw141[(new BigNumber(8)).toNumber()] = _984_v1;
        _nw141[(new BigNumber(9)).toNumber()] = _module.__default.fm1(globalState);
        _nw141[(new BigNumber(10)).toNumber()] = _984_v1;
        _nw141[(new BigNumber(11)).toNumber()] = (_module.__default.fm46((_this).f10, (_this).f20, (_this).f10, globalState)).dtor_cf52;
        _1035_v39 = _nw141;
        _1035_v39 = (((_this).f20) ? (_1035_v39) : (_1035_v39));
      }
      let _1036_v40;
      _1036_v40 = _module.D11.create_DC33((_this).f20, _984_v1);
      _1036_v40 = _1036_v40;
      let _1037_v41;
      _1037_v41 = _module.D7.create_DC23();
      let _1038_v42;
      _1038_v42 = _dafny.Seq.of(_1037_v41, _1037_v41);
      let _1039_v43;
      _1039_v43 = _module.D12.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_1040_i5) {
  return _module.D10.create_DC31((_this).f20, (_this).f10, new BigNumber(-733));
}));
      let _1041_v44;
      _1041_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1039_v43,_983_v0);
      _984_v1 = new BigNumber((_dafny.Seq.update(_1038_v42, _module.__default.safeIndex(new BigNumber(((_1041_v44).update(_1039_v43, _983_v0)).length), new BigNumber((_1038_v42).length)), _1037_v41)).length);
      r1 = true;
      let _1042_v45;
      let _init26 = ((_1043_v1, _1044_v0) => function (_1045_i6) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_1043_v1, _1043_v1), _dafny.Seq.of(_1043_v1, new BigNumber((_1044_v0).length)));
      })(_984_v1, _983_v0);
      let _nw142 = Array((new BigNumber(16)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw142.length); _i0_26++) {
        _nw142[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _1042_v45 = _nw142;
      _1042_v45 = (((_this).f20) ? (_1042_v45) : (_1042_v45));
      let _1046_v46;
      let _nw143 = Array((new BigNumber(2)).toNumber()).fill(false);
      _1046_v46 = _nw143;
      let _1047_v47;
      _1047_v47 = _module.D0.create_DC2((_this).f10, _1046_v46);
      let _source22 = _1047_v47;
      if (_source22.is_DC0) {
        let _1048___mcc_h0 = (_source22).cf0;
        let _1049_cf0 = _1048___mcc_h0;
        let _1050_v48;
        _1050_v48 = _dafny.Set.fromElements(_984_v1, _984_v1, _984_v1, _984_v1, _984_v1);
        let _1051_v49;
        _1051_v49 = _dafny.Seq.of(new BigNumber((_983_v0).length));
        let _1052_v50;
        _1052_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(new BigNumber((_983_v0).length), (_this).f10, (_this).f10, globalState),(_dafny.Set.fromElements(_1050_v48, _module.__default.fm35(_1051_v49, _1051_v49, (_this).f10, globalState))).contains(_1050_v48));
        let _1053_v51;
        let _init27 = ((_1054_v1) => function (_1055_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f20,!(!((_module.D6.create_DC20((_this).f10, (_this).f10, (_this).f10, _1054_v1)).dtor_cf33)));
        })(_984_v1);
        let _nw144 = Array((new BigNumber(22)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw144.length); _i0_27++) {
          _nw144[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1053_v51 = _nw144;
        let _rhs161 = _1042_v45;
        let _rhs162 = (_this).f10;
        let _rhs163 = (_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge((_1052_v50).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)));
        let _rhs164 = _1053_v51;
        _1042_v45 = _rhs161;
        r1 = _rhs162;
        _1052_v50 = _rhs163;
        _1053_v51 = _rhs164;
        let _1056_v52;
        _1056_v52 = _dafny.Set.fromElements((_this).f20, !((_this).f20));
        let _1057_v53;
        _1057_v53 = _dafny.Set.fromElements(_1056_v52);
        let _1058_v54;
        _1058_v54 = _module.D3.create_DC9(_1057_v53);
        let _1059_v55;
        _1059_v55 = _module.D3.create_DC12(_1058_v54);
        let _1060_v56;
        _1060_v56 = _module.D3.create_DC12(_1059_v55);
        _1060_v56 = _1060_v56;
        let _1061_v57;
        let _nw145 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1061_v57 = _nw145;
        let _index168 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length));
        (_1061_v57)[_index168] = _984_v1;
        let _index169 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length));
        (_1061_v57)[_index169] = _984_v1;
        if (_dafny.Seq.contains(_987_v4, (_this).f10)) {
          let _1062_v58;
          _1062_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), ((_1063_v2) => function (_1064_i8) {
            return _1063_v2;
          })(_985_v2)),(_this).f10);
          let _index170 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1049_cf0).length));
          (_1049_cf0)[_index170] = !((((_1062_v58).contains(_dafny.Seq.update(_983_v0, _module.__default.safeIndex(_984_v1, new BigNumber((_983_v0).length)), _985_v2))) ? ((_1062_v58).get(_dafny.Seq.update(_983_v0, _module.__default.safeIndex(_984_v1, new BigNumber((_983_v0).length)), _985_v2))) : ((_this).f20)));
          let _index171 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1049_cf0).length));
          (_1049_cf0)[_index171] = !((_this).f20) || ((_this).f10);
          let _1065_v59;
          let _nw146 = Array((_dafny.ONE).toNumber());
          _1065_v59 = _nw146;
          let _1066_v60;
          let _nw147 = new _module.C1();
          _nw147.__ctor((_this).f10);
          _1066_v60 = _nw147;
          let _index172 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1065_v59).length));
          (_1065_v59)[_index172] = _1066_v60;
          let _index173 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1065_v59).length));
          let _nw148 = new _module.C1();
          _nw148.__ctor((_1056_v52).IsProperSubsetOf(_1056_v52));
          (_1065_v59)[_index173] = _nw148;
          (globalState).f7 = (_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))];
          (_1066_v60).m5(new BigNumber((_1051_v49).length), globalState);
          r1 = _module.__default.fm0(_984_v1, globalState);
        } else {
          let _1067_v61;
          let _nw149 = new _module.C0();
          _nw149.__ctor();
          _1067_v61 = _nw149;
          r1 = (((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))]).plus((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))])).isLessThan((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))]);
          let _1068_v62;
          _1068_v62 = _dafny.Map.Empty.slice().updateUnsafe((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))],new BigNumber(143));
          let _1069_v63;
          _1069_v63 = _dafny.Map.Empty.slice().updateUnsafe(!((new BigNumber(908)).isEqualTo((((_1068_v62).contains((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))])) ? ((_1068_v62).get((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))])) : (_984_v1)))),_984_v1);
          let _1070_v64;
          _1070_v64 = _module.D3.create_DC10(_983_v0, (_this).f20, (_this).f20);
          let _index174 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length));
          let _rhs165 = (((_1069_v63).contains((_this).f20)) ? ((_1069_v63).get((_this).f20)) : ((_module.D2.create_DC7((_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))], (_this).f10, (_1061_v57)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length))])).dtor_cf11));
          let _rhs166 = !((_this).f10);
          let _rhs167 = _dafny.Seq.Concat((_1070_v64).dtor_cf14, _983_v0);
          let _lhs110 = _1061_v57;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1061_v57).length));
          _lhs110[_lhs111] = _rhs165;
          r1 = _rhs166;
          _983_v0 = _rhs167;
          r1 = (_this).f10;
          r1 = (_this).f10;
        }
      } else if (_source22.is_DC1) {
        let _1071___mcc_h1 = (_source22).cf1;
        let _1072_cf1 = _1071___mcc_h1;
        _983_v0 = _dafny.Seq.UnicodeFromString("bvlqw");
        let _1073_v65;
        _1073_v65 = _dafny.Seq.of(_984_v1);
        let _1074_v66;
        let _nw150 = new _module.C2();
        _nw150.__ctor((_984_v1).isLessThan(new BigNumber(-760)), new BigNumber((_1073_v65).length));
        _1074_v66 = _nw150;
        _1074_v66 = _1074_v66;
        let _1075_v67;
        _1075_v67 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f20);
        r1 = (((_1075_v67).contains((_1074_v66).f21)) ? ((_1075_v67).get((_1074_v66).f21)) : ((_1074_v66).f21));
        (_1074_v66).m5(_module.__default.safeDivisionInt((_1074_v66).f22, (_1074_v66).f22), globalState);
      } else {
        let _1076___mcc_h2 = (_source22).cf2;
        let _1077___mcc_h3 = (_source22).cf3;
        let _1078_cf3 = _1077___mcc_h3;
        let _1079_cf2 = _1076___mcc_h2;
        let _1080_v68;
        _1080_v68 = _dafny.Seq.of(new BigNumber(324), _984_v1);
        let _1081_v69;
        _1081_v69 = _dafny.Seq.of(_1080_v68, _1080_v68, _1080_v68, _dafny.Seq.Concat(_1080_v68, _1080_v68), _dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), ((_1082_v1) => function (_1083_i9) {
          return _1082_v1;
        })(_984_v1)));
        let _1084_v70;
        _1084_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1081_v69);
        _1081_v69 = _dafny.Seq.Concat((((_1084_v70).contains(false)) ? ((_1084_v70).get(false)) : (_1081_v69)), _1081_v69);
        let _1085_v71;
        let _nw151 = Array((new BigNumber(28)).toNumber()).fill([]);
        _1085_v71 = _nw151;
        let _1086_v72;
        let _nw152 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1086_v72 = _nw152;
        let _index175 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1085_v71).length));
        (_1085_v71)[_index175] = _1086_v72;
        let _index176 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1085_v71).length));
        let _rhs168 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kwtb"), _module.__default.fm38((_this).f10, _984_v1, globalState));
        let _rhs169 = _1086_v72;
        let _lhs112 = _1085_v71;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1085_v71).length));
        _983_v0 = _rhs168;
        _lhs112[_lhs113] = _rhs169;
        let _1087_v73;
        _1087_v73 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_983_v0, _983_v0),false);
        _1087_v73 = (_1087_v73).update(_983_v0, (_984_v1).isLessThanOrEqualTo(_984_v1));
        r1 = !(_1079_cf2) || ((_this).f10);
      }
      let _1088_v74;
      _1088_v74 = _dafny.Set.fromElements(_dafny.Seq.Concat(_983_v0, _983_v0), _dafny.Seq.Concat(_dafny.Seq.update(_983_v0, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_984_v1)).cardinality()), new BigNumber((_983_v0).length)), _985_v2), _983_v0), _983_v0);
      r0 = _1088_v74;
      r1 = (_this).f10;
      r2 = (((_this).f20) ? (_984_v1) : (_984_v1));
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      let _1089_v0;
      _1089_v0 = false;
      let _1090_v1;
      let _nw153 = Array((new BigNumber(18)).toNumber()).fill(false);
      _1090_v1 = _nw153;
      let _1091_v2;
      _1091_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(113),_1090_v1);
      let _1092_v3;
      _1092_v3 = _dafny.MultiSet.fromElements(new BigNumber(381));
      let _1093_v4;
      _1093_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v0,_1090_v1);
      let _1094_v5;
      let _nw154 = Array((new BigNumber(15)).toNumber());
      _nw154[(_dafny.ZERO).toNumber()] = _1090_v1;
      _nw154[(_dafny.ONE).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(2)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(3)).toNumber()] = (((_1091_v2).contains(new BigNumber((_1092_v3).cardinality()))) ? ((_1091_v2).get(new BigNumber((_1092_v3).cardinality()))) : (_1090_v1));
      _nw154[(new BigNumber(4)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(5)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(6)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(7)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(8)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(9)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(10)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(11)).toNumber()] = _1090_v1;
      _nw154[(new BigNumber(12)).toNumber()] = (((_1093_v4).contains((_this).f20)) ? ((_1093_v4).get((_this).f20)) : (_1090_v1));
      _nw154[(new BigNumber(13)).toNumber()] = (((_this).f20) ? (_1090_v1) : (_1090_v1));
      _nw154[(new BigNumber(14)).toNumber()] = _1090_v1;
      _1094_v5 = _nw154;
      let _index177 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_1094_v5).length));
      (_1094_v5)[_index177] = _1090_v1;
      let _1095_v6;
      _1095_v6 = _dafny.Seq.of((_this).f20);
      let _1096_v7;
      _1096_v7 = _dafny.Seq.of(p0, p0, p0, p0, new BigNumber(-966));
      let _1097_v8;
      _1097_v8 = _dafny.Set.fromElements(new BigNumber((_1095_v6).length), new BigNumber(908), (_1096_v7)[_module.__default.safeIndex(p0, new BigNumber((_1096_v7).length))]);
      let _1098_v9;
      _1098_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p0).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_1097_v8).length))));
      let _1099_v10;
      _1099_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v0,(((_1098_v9).contains(p0)) ? ((_1098_v9).get(p0)) : ((_this).f20)));
      let _1100_v11;
      _1100_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ftlkutv")).length),_1099_v10);
      let _1101_v12;
      _1101_v12 = _module.D14.create_DC41(_1100_v11);
      let _index178 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_1094_v5).length));
      let _rhs170 = (((_1098_v9).contains(new BigNumber(((_1099_v10).Merge(_1099_v10)).length))) ? ((_1098_v9).get(new BigNumber(((_1099_v10).Merge(_1099_v10)).length))) : ((_this).fm8(new BigNumber(-556), (_this).f10, true, globalState)));
      let _rhs171 = _1090_v1;
      let _rhs172 = (new BigNumber((_dafny.Set.fromElements(_1098_v9)).length)).isEqualTo((p0).plus(p0));
      let _rhs173 = new BigNumber(((_1101_v12).dtor_cf71).length);
      let _lhs114 = _1094_v5;
      let _lhs115 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_1094_v5).length));
      let _lhs116 = globalState;
      _1089_v0 = _rhs170;
      _lhs114[_lhs115] = _rhs171;
      _1089_v0 = _rhs172;
      _lhs116.f7 = _rhs173;
      _1089_v0 = (new BigNumber((function () {
        let _coll26 = new _dafny.Map();
        for (const _compr_26 of _dafny.IntegerRange(new BigNumber(65), new BigNumber(828))) {
          let _1102_v13 = _compr_26;
          if (((new BigNumber(65)).isLessThanOrEqualTo(_1102_v13)) && ((_1102_v13).isLessThan(new BigNumber(828)))) {
            _coll26.push([(_1102_v13).multipliedBy(p0),_1089_v0]);
          }
        }
        return _coll26;
      }()).length)).isLessThanOrEqualTo(p0);
      let _1103_v14;
      _1103_v14 = _dafny.Seq.UnicodeFromString("gip");
      let _1104_v15;
      _1104_v15 = _dafny.MultiSet.fromElements((_1095_v6)[_module.__default.safeIndex(p0, new BigNumber((_1095_v6).length))], false);
      let _1105_v16;
      let _nw155 = Array((new BigNumber(10)).toNumber());
      _nw155[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1103_v14).length), p0);
      _nw155[(_dafny.ONE).toNumber()] = (p0).multipliedBy(new BigNumber((_1104_v15).cardinality()));
      _nw155[(new BigNumber(2)).toNumber()] = new BigNumber(((_1104_v15).Intersect(_dafny.MultiSet.fromElements(_1089_v0))).cardinality());
      _nw155[(new BigNumber(3)).toNumber()] = p0;
      _nw155[(new BigNumber(4)).toNumber()] = p0;
      _nw155[(new BigNumber(5)).toNumber()] = (new BigNumber(231)).minus(p0);
      _nw155[(new BigNumber(6)).toNumber()] = p0;
      _nw155[(new BigNumber(7)).toNumber()] = _module.__default.fm1(globalState);
      _nw155[(new BigNumber(8)).toNumber()] = p0;
      _nw155[(new BigNumber(9)).toNumber()] = p0;
      _1105_v16 = _nw155;
      let _1106_v17;
      _1106_v17 = new _dafny.CodePoint('v'.codePointAt(0));
      let _rhs174 = _1105_v16;
      let _rhs175 = (_dafny.ZERO).minus((p0).multipliedBy(_module.__default.safeModuloInt(_module.__default.fm1(globalState), p0)));
      let _rhs176 = ((_module.D4.create_DC15(p0, _dafny.MultiSet.FromArray(_1096_v7), _1106_v17, p0)).dtor_cf26).minus(new BigNumber((_1095_v6).length));
      let _rhs177 = new BigNumber((_1095_v6).length);
      let _lhs117 = globalState;
      let _lhs118 = globalState;
      let _lhs119 = globalState;
      _1105_v16 = _rhs174;
      _lhs117.f0 = _rhs175;
      _lhs118.f1 = _rhs176;
      _lhs119.f0 = _rhs177;
      let _1107_i0;
      _1107_i0 = _dafny.ZERO;
      L3: {
        while (!(!((_this).f10))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1107_i0)) {
              break L3;
            }
            _1107_i0 = (_1107_i0).plus(_dafny.ONE);
            _1089_v0 = (_this).f10;
            _1093_v4 = ((_1093_v4).Merge(_1093_v4)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1089_v0,_1090_v1));
            _1089_v0 = true;
            let _1108_v18;
            let _init28 = ((_1109_v15, _1110_v14, _1111_v9, _1112_v0) => function (_1113_i1) {
              return (_1109_v15).update((((_1111_v9).contains(new BigNumber((_1110_v14).length))) ? ((_1111_v9).get(new BigNumber((_1110_v14).length))) : (_1112_v0)), _module.__default.abs(new BigNumber((_1110_v14).length)));
            })(_1104_v15, _1103_v14, _1098_v9, _1089_v0);
            let _nw156 = Array((new BigNumber(18)).toNumber());
            for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw156.length); _i0_28++) {
              _nw156[_i0_28] = _init28(new BigNumber(_i0_28));
            }
            _1108_v18 = _nw156;
            let _index179 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_1108_v18).length));
            (_1108_v18)[_index179] = (_dafny.MultiSet.fromElements(_1089_v0)).Difference(_1104_v15);
            let _index180 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_1108_v18).length));
            (_1108_v18)[_index180] = _1104_v15;
          }
        }
      }
      let _1114_v19;
      _1114_v19 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0), _dafny.Seq.Concat(_1096_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), ((_1115_p0) => function (_1116_i2) {
        return _1115_p0;
      })(p0))), _1096_v7);
      _1114_v19 = _module.__default.fm47(p0, (_this).f10, globalState);
      let _index181 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1105_v16).length));
      (_1105_v16)[_index181] = p0;
      let _index182 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1105_v16).length));
      (_1105_v16)[_index182] = p0;
      return;
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return !((!(false)) && (false)) || (!(false));
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber(-916));
    };
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1117_v0;
      let _nw157 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
      _1117_v0 = _nw157;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1117_v0).length))) {
        let _1118_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1118_i0)) && ((_1118_i0).isLessThan(new BigNumber((_1117_v0).length))))) {
          (_1117_v0)[(_1118_i0)] = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (!(false)) : (true)),_module.D6.create_DC21(true, false, true, new BigNumber(970)));
        }
      }
      let _1119_v1;
      _1119_v1 = false;
      let _1120_v2;
      _1120_v2 = new BigNumber(-545);
      let _1121_v3;
      let _nw158 = new _module.C2();
      _nw158.__ctor(_1119_v1, _1120_v2);
      _1121_v3 = _nw158;
      let _1122_v4;
      _1122_v4 = _dafny.MultiSet.fromElements(_1119_v1, (_1121_v3).f21);
      _1122_v4 = _1122_v4;
      let _1123_v5;
      _1123_v5 = _dafny.Seq.UnicodeFromString("tiojau");
      let _1124_v6;
      _1124_v6 = _module.D10.create_DC30(_dafny.Set.fromElements(_1119_v1, _1119_v1, (_1121_v3).f21, (_1121_v3).f21, (_1121_v3).f21));
      let _pat_let_tv22 = _1123_v5;
      _1123_v5 = function (_source23) {
        if (_source23.is_DC31) {
          let _1125___mcc_h0 = (_source23).cf50;
          let _1126___mcc_h1 = (_source23).cf51;
          let _1127___mcc_h2 = (_source23).cf52;
          let _1128_cf52 = _1127___mcc_h2;
          let _1129_cf51 = _1126___mcc_h1;
          let _1130_cf50 = _1125___mcc_h0;
          return _dafny.Seq.UnicodeFromString("ydrrviwd");
        } else {
          let _1131___mcc_h3 = (_source23).cf49;
          let _1132_cf49 = _1131___mcc_h3;
          return _pat_let_tv22;
        }
      }(_1124_v6);
      r1 = (_1119_v1) === (_1119_v1);
      let _1133_v7;
      _1133_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1119_v1,(_1121_v3).f21);
      let _1134_v8;
      _1134_v8 = _dafny.Set.fromElements(_1133_v7, _1133_v7);
      (globalState).f7 = _module.__default.safeModuloInt((new BigNumber(113)).plus(_1120_v2), (_this).fm9(_dafny.Seq.update(_1123_v5, _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1123_v5).length)), new _dafny.CodePoint('m'.codePointAt(0))), _1134_v8, globalState));
      let _1135_v9;
      _1135_v9 = _dafny.Map.Empty.slice().updateUnsafe(!((_1121_v3).f21),_1123_v5);
      let _1136_v10;
      _1136_v10 = _dafny.Set.fromElements(_1123_v5);
      r0 = ((_dafny.Set.fromElements((((_1135_v9).contains((_1121_v3).f21)) ? ((_1135_v9).get((_1121_v3).f21)) : (_1123_v5)))).Intersect(_1136_v10)).Union(_1136_v10);
      r1 = (_1121_v3).f21;
      r2 = (_1121_v3).f22;
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      let _1137_v0;
      _1137_v0 = true;
      let _rhs178 = _1137_v0;
      let _rhs179 = new BigNumber((_module.__default.fm48(globalState)).cardinality());
      let _rhs180 = p0;
      let _lhs120 = globalState;
      let _lhs121 = globalState;
      _1137_v0 = _rhs178;
      _lhs120.f7 = _rhs179;
      _lhs121.f7 = _rhs180;
      let _1138_v1;
      let _nw159 = Array((new BigNumber(6)).toNumber()).fill(false);
      _1138_v1 = _nw159;
      let _index183 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
      (_1138_v1)[_index183] = true;
      let _index184 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
      (_1138_v1)[_index184] = (p0).isLessThan(p0);
      let _1139_v2;
      _1139_v2 = _dafny.MultiSet.fromElements(p0);
      let _1140_v3;
      _1140_v3 = _dafny.Seq.of(p0, p0, p0);
      let _1141_i0;
      _1141_i0 = _dafny.ZERO;
      L4: {
        let _pat_let_tv23 = _1137_v0;
        let _pat_let_tv24 = _1137_v0;
        let _pat_let_tv25 = _1137_v0;
        let _pat_let_tv26 = _1138_v1;
        let _pat_let_tv27 = _1138_v1;
        while (function (_source24) {
          if (_source24.is_DC10) {
            let _1150___mcc_h0 = (_source24).cf14;
            let _1151___mcc_h1 = (_source24).cf15;
            let _1152___mcc_h2 = (_source24).cf16;
            let _1153_cf16 = _1152___mcc_h2;
            let _1154_cf15 = _1151___mcc_h1;
            let _1155_cf14 = _1150___mcc_h0;
            return _1154_cf15;
          } else if (_source24.is_DC11) {
            let _1156___mcc_h3 = (_source24).cf17;
            let _1157___mcc_h4 = (_source24).cf18;
            let _1158_cf18 = _1157___mcc_h4;
            let _1159_cf17 = _1156___mcc_h3;
            return (_pat_let_tv23) && (_pat_let_tv24);
          } else if (_source24.is_DC9) {
            let _1160___mcc_h5 = (_source24).cf13;
            let _1161_cf13 = _1160___mcc_h5;
            return _pat_let_tv25;
          } else {
            let _1162___mcc_h6 = (_source24).cf19;
            let _1163_cf19 = _1162___mcc_h6;
            return (_pat_let_tv27)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_pat_let_tv26).length))];
          }
        }(_module.__default.fm49((((_1139_v2).contains(p0)) ? ((_1139_v2).get(p0)) : ((_1140_v3)[_module.__default.safeIndex(p0, new BigNumber((_1140_v3).length))])), p0, globalState))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1141_i0)) {
              break L4;
            }
            _1141_i0 = (_1141_i0).plus(_dafny.ONE);
            let _1142_v4;
            _1142_v4 = _dafny.Seq.of(_module.__default.fm0(p0, globalState));
            let _index185 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
            let _rhs181 = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,false)).update((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))], (_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))])).length)).isLessThanOrEqualTo(new BigNumber((_1142_v4).length));
            let _rhs182 = new BigNumber((_dafny.Seq.update(_1142_v4, _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1142_v4).length)), (_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))])).length);
            let _lhs122 = _1138_v1;
            let _lhs123 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
            let _lhs124 = globalState;
            _lhs122[_lhs123] = _rhs181;
            _lhs124.f1 = _rhs182;
            let _1143_v5;
            let _nw160 = Array((new BigNumber(18)).toNumber()).fill([]);
            _1143_v5 = _nw160;
            let _1144_v6;
            let _init29 = function (_1145_i1) {
              return _dafny.Seq.UnicodeFromString("wmyxcurb");
            };
            let _nw161 = Array((new BigNumber(7)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw161.length); _i0_29++) {
              _nw161[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _1144_v6 = _nw161;
            let _index186 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1143_v5).length));
            (_1143_v5)[_index186] = _1144_v6;
            let _index187 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
            let _index188 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1143_v5).length));
            let _rhs183 = _1137_v0;
            let _rhs184 = p0;
            let _rhs185 = _1144_v6;
            let _rhs186 = _1142_v4;
            let _lhs125 = _1138_v1;
            let _lhs126 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
            let _lhs127 = globalState;
            let _lhs128 = _1143_v5;
            let _lhs129 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1143_v5).length));
            _lhs125[_lhs126] = _rhs183;
            _lhs127.f7 = _rhs184;
            _lhs128[_lhs129] = _rhs185;
            _1142_v4 = _rhs186;
            let _1146_v7;
            let _nw162 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
            _1146_v7 = _nw162;
            let _1147_v8;
            _1147_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))])).length));
            let _index189 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1146_v7).length));
            (_1146_v7)[_index189] = _1147_v8;
            let _index190 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1146_v7).length));
            (_1146_v7)[_index190] = _1147_v8;
            let _1148_v9;
            _1148_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1137_v0,_1138_v1);
            let _1149_v10;
            _1149_v10 = _dafny.Set.fromElements(p0, new BigNumber((_1148_v9).length));
            _1137_v0 = (new BigNumber(-52)).isLessThan((p0).multipliedBy(new BigNumber((_1149_v10).length)));
          }
        }
      }
      let _hi4 = p0;
      for (let _1164_i2 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), function (_1175_i3) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length); _1164_i2.isLessThan(_hi4); _1164_i2 = _1164_i2.plus(_dafny.ONE)) {
        let _1165_v11;
        _1165_v11 = _dafny.Seq.of((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]);
        let _index191 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
        (_1138_v1)[_index191] = !((_1165_v11)[_module.__default.safeIndex(_1164_i2, new BigNumber((_1165_v11).length))]);
        let _1166_v12;
        _1166_v12 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1167_v13;
        _1167_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1165_v11,p0);
        let _index192 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
        let _index193 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
        let _rhs187 = (_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))];
        let _rhs188 = _1137_v0;
        let _rhs189 = (((p0).isLessThanOrEqualTo(p0)) ? (_1166_v12) : (_1166_v12));
        let _rhs190 = (((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]) ? (((true) ? (new BigNumber((_1167_v13).length)) : (p0))) : (_1164_i2));
        let _lhs130 = _1138_v1;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
        let _lhs132 = _1138_v1;
        let _lhs133 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
        let _lhs134 = globalState;
        _lhs130[_lhs131] = _rhs187;
        _lhs132[_lhs133] = _rhs188;
        _1166_v12 = _rhs189;
        _lhs134.f0 = _rhs190;
        let _1168_v14;
        let _init30 = ((_1169_p0) => function (_1170_i4) {
          return (_1170_i4).plus(_1169_p0);
        })(p0);
        let _nw163 = Array((new BigNumber(24)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw163.length); _i0_30++) {
          _nw163[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1168_v14 = _nw163;
        _1168_v14 = _1168_v14;
        let _1171_v15;
        _1171_v15 = _dafny.Seq.UnicodeFromString("evnnhnsnk");
        let _1172_v16;
        _1172_v16 = _module.D1.create_DC4((_dafny.ZERO).minus(p0), _1168_v14);
        let _1173_v17;
        _1173_v17 = _module.D14.create_DC42(_1172_v16, _1165_v11, false);
        let _1174_v18;
        _1174_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1173_v17,(_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]);
        let _rhs191 = (_module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.update(_1165_v11, _module.__default.safeIndex(p0, new BigNumber((_1165_v11).length)), (_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))])).length))).isEqualTo((new BigNumber((_1171_v15).length)).multipliedBy(p0));
        let _rhs192 = (p0).minus(_module.__default.safeModuloInt(new BigNumber((_1174_v18).length), _1164_i2));
        let _lhs135 = globalState;
        _1137_v0 = _rhs191;
        _lhs135.f7 = _rhs192;
      }
      let _1176_v19;
      _1176_v19 = _module.D11.create_DC33((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))], p0);
      let _1177_i5;
      _1177_i5 = _dafny.ZERO;
      L5: {
        while (((_1176_v19).dtor_cf54) || ((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))])) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1177_i5)) {
              break L5;
            }
            _1177_i5 = (_1177_i5).plus(_dafny.ONE);
            let _1178_v20;
            _1178_v20 = _module.D13.create_DC38(_dafny.MultiSet.fromElements(_1138_v1, _1138_v1, _1138_v1));
            let _1179_v21;
            _1179_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,(_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]),_1178_v20);
            _1179_v21 = _1179_v21;
            let _1180_v22;
            _1180_v22 = _module.D7.create_DC23();
            _1180_v22 = _1180_v22;
            let _1181_v23;
            _1181_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _1182_v24;
            _1182_v24 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
            let _1183_v25;
            _1183_v25 = _dafny.Seq.UnicodeFromString("ovsrjxkbl");
            let _1184_v26;
            _1184_v26 = new _dafny.CodePoint('w'.codePointAt(0));
            let _1185_v27;
            _1185_v27 = _dafny.Map.Empty.slice().updateUnsafe(!(_1137_v0),_1137_v0);
            let _1186_v28;
            _1186_v28 = _dafny.MultiSet.fromElements((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]);
            let _1187_v29;
            _1187_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1184_v26,_1137_v0);
            let _1188_v30;
            _1188_v30 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!((((_1187_v29).contains(_1184_v26)) ? ((_1187_v29).get(_1184_v26)) : (true))),_1137_v0), _1185_v27);
            let _1189_v31;
            _1189_v31 = _dafny.Seq.of(_module.__default.fm3(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(840)), ((_1190_v3) => function (_1191_i6) {
              return _1190_v3;
            })(_1140_v3))).length), new BigNumber(-277), globalState));
            let _1192_v32;
            let _nw164 = Array((new BigNumber(27)).toNumber());
            _nw164[(_dafny.ZERO).toNumber()] = (((_1181_v23).contains(new BigNumber(698))) ? ((_1181_v23).get(new BigNumber(698))) : (new BigNumber(485)));
            _nw164[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_1182_v24).Merge(_1182_v24)).length));
            _nw164[(new BigNumber(2)).toNumber()] = p0;
            _nw164[(new BigNumber(3)).toNumber()] = (_this).fm9(_dafny.Seq.update(_dafny.Seq.update(_1183_v25, _module.__default.safeIndex(p0, new BigNumber((_1183_v25).length)), _1184_v26), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_1183_v25, _module.__default.safeIndex(p0, new BigNumber((_1183_v25).length)), _1184_v26)).length)), _1184_v26), _dafny.Set.fromElements(_1185_v27, _1185_v27), globalState);
            _nw164[(new BigNumber(4)).toNumber()] = new BigNumber((_1186_v28).cardinality());
            _nw164[(new BigNumber(5)).toNumber()] = new BigNumber(802);
            _nw164[(new BigNumber(6)).toNumber()] = p0;
            _nw164[(new BigNumber(7)).toNumber()] = p0;
            _nw164[(new BigNumber(8)).toNumber()] = p0;
            _nw164[(new BigNumber(9)).toNumber()] = (_this).fm9(_1183_v25, _1188_v30, globalState);
            _nw164[(new BigNumber(10)).toNumber()] = p0;
            _nw164[(new BigNumber(11)).toNumber()] = new BigNumber((_1182_v24).length);
            _nw164[(new BigNumber(12)).toNumber()] = ((_this).fm9(_1183_v25, _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))],_1137_v0)), globalState)).minus(p0);
            _nw164[(new BigNumber(13)).toNumber()] = p0;
            _nw164[(new BigNumber(14)).toNumber()] = p0;
            _nw164[(new BigNumber(15)).toNumber()] = new BigNumber(204);
            _nw164[(new BigNumber(16)).toNumber()] = ((!(_1137_v0)) ? (p0) : (p0));
            _nw164[(new BigNumber(17)).toNumber()] = new BigNumber(867);
            _nw164[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
            _nw164[(new BigNumber(19)).toNumber()] = p0;
            _nw164[(new BigNumber(20)).toNumber()] = p0;
            _nw164[(new BigNumber(21)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
            _nw164[(new BigNumber(22)).toNumber()] = (new BigNumber(14)).plus(p0);
            _nw164[(new BigNumber(23)).toNumber()] = p0;
            _nw164[(new BigNumber(24)).toNumber()] = _module.__default.safeDivisionInt(p0, (_this).fm9(_1183_v25, _1188_v30, globalState));
            _nw164[(new BigNumber(25)).toNumber()] = (p0).minus(new BigNumber(((_1189_v31)[_module.__default.safeIndex(p0, new BigNumber((_1189_v31).length))]).length));
            _nw164[(new BigNumber(26)).toNumber()] = (_dafny.ZERO).minus(p0);
            _1192_v32 = _nw164;
            let _index194 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1192_v32).length));
            (_1192_v32)[_index194] = new BigNumber(((_1182_v24).Merge(_1182_v24)).length);
            let _1193_v33;
            _1193_v33 = _module.D4.create_DC15(p0, _1139_v2, _1184_v26, new BigNumber((_1185_v27).length));
            let _index195 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
            let _index196 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1192_v32).length));
            let _rhs193 = (_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))];
            let _rhs194 = p0;
            let _rhs195 = _1193_v33;
            let _rhs196 = p0;
            let _lhs136 = _1138_v1;
            let _lhs137 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
            let _lhs138 = _1192_v32;
            let _lhs139 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1192_v32).length));
            let _lhs140 = globalState;
            _lhs136[_lhs137] = _rhs193;
            _lhs138[_lhs139] = _rhs194;
            _1193_v33 = _rhs195;
            _lhs140.f7 = _rhs196;
            let _1194_v34;
            let _init31 = ((_1195_v25) => function (_1196_i7) {
              return _dafny.MultiSet.FromArray(_1195_v25);
            })(_1183_v25);
            let _nw165 = Array((new BigNumber(12)).toNumber());
            for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw165.length); _i0_31++) {
              _nw165[_i0_31] = _init31(new BigNumber(_i0_31));
            }
            _1194_v34 = _nw165;
            let _1197_v35;
            _1197_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1194_v34,(_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]);
            _1197_v35 = (_1197_v35).update(_1194_v34, _1137_v0);
          }
        }
      }
      let _hi5 = ((_module.__default.fm0(new BigNumber(216), globalState)) ? (p0) : (p0));
      for (let _1198_i8 = p0; _1198_i8.isLessThan(_hi5); _1198_i8 = _1198_i8.plus(_dafny.ONE)) {
        let _1199_v36;
        _1199_v36 = _dafny.Set.fromElements(p0, _module.__default.fm1(globalState));
        (globalState).f1 = (_module.__default.safeModuloInt(new BigNumber((_1199_v36).length), new BigNumber((_1140_v3).length))).plus(new BigNumber(701));
        let _1200_v37;
        _1200_v37 = _dafny.Seq.of(_1137_v0, _module.__default.fm0((_dafny.ZERO).minus(p0), globalState), (((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]) ? ((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))]) : ((_1138_v1)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length))])), (new BigNumber((_dafny.MultiSet.FromArray(_1140_v3)).cardinality())).isLessThan(_1198_i8));
        let _index197 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1138_v1).length));
        (_1138_v1)[_index197] = (_1200_v37)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1200_v37)).cardinality()), new BigNumber((_1200_v37).length))];
        let _1201_v38;
        _1201_v38 = _dafny.Seq.UnicodeFromString("tleyvejo");
        let _1202_v39;
        _1202_v39 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1203_v40;
        let _nw166 = new _module.C1();
        _nw166.__ctor(_1137_v0);
        _1203_v40 = _nw166;
        let _1204_v41;
        _1204_v41 = _dafny.Set.fromElements(_1203_v40);
        (globalState).f7 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1201_v38, _module.__default.safeIndex(_1198_i8, new BigNumber((_1201_v38).length)), _1202_v39), _module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_dafny.Seq.update(_1201_v38, _module.__default.safeIndex(_1198_i8, new BigNumber((_1201_v38).length)), _1202_v39)).length)), _1202_v39)).length)).plus(new BigNumber((_1204_v41).length));
        let _1205_v42;
        let _1206_v43;
        let _out36;
        let _out37;
        let _outcollector17 = (_this).m14(globalState);
        _out36 = _outcollector17[0];
        _out37 = _outcollector17[1];
        _1205_v42 = _out36;
        _1206_v43 = _out37;
      }
      return;
    }
    m14(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _1207_v0;
      _1207_v0 = new BigNumber(154);
      let _1208_v1;
      let _nw167 = Array((new BigNumber(24)).toNumber()).fill(false);
      _1208_v1 = _nw167;
      let _1209_v2;
      let _1210_v3;
      let _out38;
      let _out39;
      let _outcollector18 = _module.__default.m0(_1207_v0, _1208_v1, globalState);
      _out38 = _outcollector18[0];
      _out39 = _outcollector18[1];
      _1209_v2 = _out38;
      _1210_v3 = _out39;
      let _1211_v4;
      _1211_v4 = true;
      let _1212_v5;
      _1212_v5 = _dafny.Set.fromElements(_1211_v4, _1211_v4);
      let _1213_v6;
      _1213_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1212_v5).length),_1207_v0);
      if (!(_1213_v6).contains(_1207_v0)) {
        _1211_v4 = true;
        let _index198 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length));
        (_1208_v1)[_index198] = _1211_v4;
        let _index199 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length));
        (_1208_v1)[_index199] = _1211_v4;
        let _1214_v7;
        let _nw168 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1214_v7 = _nw168;
        let _1215_v8;
        let _1216_v9;
        let _out40;
        let _out41;
        let _outcollector19 = _module.__default.m0(_module.__default.safeModuloInt(_1207_v0, (_dafny.ZERO).minus(new BigNumber((_1209_v2).length))), _1214_v7, globalState);
        _out40 = _outcollector19[0];
        _out41 = _outcollector19[1];
        _1215_v8 = _out40;
        _1216_v9 = _out41;
        let _1217_v10;
        let _nw169 = new _module.C3();
        _nw169.__ctor(((_1208_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length))]) === ((_1208_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length))]), _module.__default.fm0(new BigNumber(366), globalState));
        _1217_v10 = _nw169;
        let _1218_v11;
        _1218_v11 = _module.D7.create_DC22(_1213_v6);
        let _source25 = _1218_v11;
        if (_source25.is_DC23) {
          let _1219_v12;
          let _nw170 = new _module.C2();
          _nw170.__ctor(_1211_v4, (_dafny.ZERO).minus(_1207_v0));
          _1219_v12 = _nw170;
          _1211_v4 = _1211_v4;
          _1211_v4 = (_1212_v5).IsSubsetOf((_1212_v5).Difference(_module.__default.fm30(globalState)));
          let _1220_v13;
          _1220_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1212_v5).IsProperSubsetOf(_1212_v5),_1214_v7);
          _1214_v7 = (((_1220_v13).contains((_1219_v12).f21)) ? ((_1220_v13).get((_1219_v12).f21)) : (_1214_v7));
        } else if (_source25.is_DC24) {
          let _1221___mcc_h0 = (_source25).cf41;
          let _1222_cf41 = _1221___mcc_h0;
          let _1223_v14;
          let _nw171 = new _module.C1();
          _nw171.__ctor((_1208_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length))]);
          _1223_v14 = _nw171;
          let _index200 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length));
          let _rhs197 = _1223_v14;
          let _rhs198 = (_1217_v10).f20;
          let _lhs141 = _1208_v1;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length));
          _1223_v14 = _rhs197;
          _lhs141[_lhs142] = _rhs198;
          let _1224_v15;
          _1224_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(962),(_1217_v10).f20);
          let _1225_v16;
          _1225_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1217_v10).f20,(_1208_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length))]);
          let _1226_v17;
          _1226_v17 = _dafny.MultiSet.fromElements(_1225_v16);
          (globalState).f1 = new BigNumber(((_1224_v15).update(new BigNumber((_1226_v17).cardinality()), (_1223_v14).fm6(new BigNumber(242), _1207_v0, _1222_cf41, _1211_v4, globalState))).length);
          let _index201 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length));
          (_1208_v1)[_index201] = _1211_v4;
          _1211_v4 = (_1222_cf41).isLessThanOrEqualTo(new BigNumber(62));
        } else {
          let _1227___mcc_h1 = (_source25).cf40;
          let _1228_cf40 = _1227___mcc_h1;
          _1211_v4 = !(true);
          let _1229_v18;
          let _nw172 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _1229_v18 = _nw172;
          let _1230_v19;
          _1230_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1208_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length))],!(_1211_v4));
          let _index202 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1229_v18).length));
          (_1229_v18)[_index202] = _1230_v19;
          let _index203 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1229_v18).length));
          (_1229_v18)[_index203] = _dafny.Map.Empty.slice().updateUnsafe((_1208_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1208_v1).length))],(_1217_v10).f20);
          let _1231_v20;
          _1231_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(_1211_v4),_module.__default.safeModuloInt(_1207_v0, new BigNumber((_1216_v9).length)));
          _1231_v20 = (_1231_v20).update((_1217_v10).f20, _1207_v0);
          let _1232_v21;
          _1232_v21 = new _dafny.CodePoint('l'.codePointAt(0));
          _1232_v21 = _1232_v21;
        }
      } else {
        let _1233_v22;
        let _nw173 = new _module.C0();
        _nw173.__ctor();
        _1233_v22 = _nw173;
        _1209_v2 = _1210_v3;
        let _1234_v23;
        _1234_v23 = _dafny.Set.fromElements(_1208_v1, _1208_v1);
        let _1235_v24;
        _1235_v24 = _dafny.Seq.of(true, (_1234_v23).IsProperSubsetOf(_1234_v23), !(_1212_v5).equals(_1212_v5), _1211_v4, (_1207_v0).isLessThan(_1207_v0));
        let _1236_v25;
        _1236_v25 = _dafny.MultiSet.fromElements(_1211_v4);
        let _rhs199 = (_1235_v24)[_module.__default.safeIndex(new BigNumber(523), new BigNumber((_1235_v24).length))];
        let _rhs200 = new BigNumber((((_1236_v25).update(true, _module.__default.abs(_1207_v0))).Union(_1236_v25)).cardinality());
        let _lhs143 = globalState;
        _1211_v4 = _rhs199;
        _lhs143.f1 = _rhs200;
        let _index204 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1208_v1).length));
        (_1208_v1)[_index204] = _module.__default.fm0(new BigNumber((_1210_v3).length), globalState);
        let _index205 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1208_v1).length));
        (_1208_v1)[_index205] = (_1233_v22).fm37(_1209_v2, (_dafny.ZERO).minus(_1207_v0), _1211_v4, _1211_v4, globalState);
        let _1237_v26;
        let _init32 = ((_1238_v4) => function (_1239_i0) {
          return _1238_v4;
        })(_1211_v4);
        let _nw174 = Array((new BigNumber(17)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw174.length); _i0_32++) {
          _nw174[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1237_v26 = _nw174;
        let _1240_v27;
        let _1241_v28;
        let _out42;
        let _out43;
        let _outcollector20 = _module.__default.m0(_1207_v0, _1237_v26, globalState);
        _out42 = _outcollector20[0];
        _out43 = _outcollector20[1];
        _1240_v27 = _out42;
        _1241_v28 = _out43;
      }
      r0 = (new BigNumber(-298)).minus(_1207_v0);
      let _1242_v29;
      let _nw175 = new _module.C3();
      _nw175.__ctor(_1211_v4, !(_1211_v4));
      _1242_v29 = _nw175;
      if ((_1242_v29).f20) {
        let _1243_v30;
        _1243_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1242_v29).f20,!((_1207_v0).isLessThan(_1207_v0)));
        let _1244_v31;
        _1244_v31 = _dafny.MultiSet.fromElements(_1207_v0, _1207_v0);
        let _1245_v32;
        _1245_v32 = _module.D3.create_DC11(_1207_v0, _1243_v30);
        let _1246_v33;
        _1246_v33 = _dafny.Map.Empty.slice().updateUnsafe((((_1244_v31).contains((_1242_v29).fm9(_1209_v2, _dafny.Set.fromElements(_1243_v30, _1243_v30, _1243_v30, (_1245_v32).dtor_cf18), globalState))) ? ((_1244_v31).get((_1242_v29).fm9(_1209_v2, _dafny.Set.fromElements(_1243_v30, _1243_v30, _1243_v30, (_1245_v32).dtor_cf18), globalState))) : (_1207_v0)),_dafny.Map.Empty.slice().updateUnsafe((_1242_v29).f20,!((_1242_v29).f20)));
        let _rhs201 = _1208_v1;
        let _rhs202 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_1209_v2, _1210_v3)).length), _1207_v0);
        let _rhs203 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), function (_1247_i1) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length), _1207_v0);
        let _rhs204 = (_1242_v29).f20;
        let _rhs205 = ((((_1246_v33).contains(_1207_v0)) ? ((_1246_v33).get(_1207_v0)) : (_1243_v30))).update((_1242_v29).f20, !_dafny.areEqual(_dafny.Seq.UnicodeFromString("pkc"), _1209_v2));
        let _lhs144 = globalState;
        let _lhs145 = globalState;
        _1208_v1 = _rhs201;
        _lhs144.f7 = _rhs202;
        _lhs145.f1 = _rhs203;
        _1211_v4 = _rhs204;
        _1243_v30 = _rhs205;
        let _1248_v34;
        let _init33 = ((_1249_v0) => function (_1250_i2) {
          return (_1250_i2).multipliedBy(_1249_v0);
        })(_1207_v0);
        let _nw176 = Array((new BigNumber(6)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw176.length); _i0_33++) {
          _nw176[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1248_v34 = _nw176;
        let _index206 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1248_v34).length));
        (_1248_v34)[_index206] = new BigNumber(715);
        let _1251_v35;
        _1251_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1211_v4,_1207_v0);
        let _1252_v36;
        _1252_v36 = _dafny.Set.fromElements(_1251_v35);
        let _1253_v37;
        _1253_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v0,_1252_v36);
        let _index207 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1248_v34).length));
        (_1248_v34)[_index207] = (_dafny.ZERO).minus(new BigNumber(((((_1253_v37).contains(new BigNumber((_dafny.Seq.UnicodeFromString("fqjtt")).length))) ? ((_1253_v37).get(new BigNumber((_dafny.Seq.UnicodeFromString("fqjtt")).length))) : (_1252_v36))).length));
        _1211_v4 = !(false);
        let _1254_v38;
        _1254_v38 = _module.D10.create_DC30(_1212_v5);
        _1254_v38 = function (_pat_let21_0) {
          return function (_1255_dt__update__tmp_h0) {
            return function (_pat_let22_0) {
              return function (_1256_dt__update_hcf49_h0) {
                return _module.D10.create_DC30(_1256_dt__update_hcf49_h0);
              }(_pat_let22_0);
            }(_dafny.Set.fromElements(false));
          }(_pat_let21_0);
        }(_module.D10.create_DC30(_1212_v5));
        let _index208 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1208_v1).length));
        (_1208_v1)[_index208] = (_1207_v0).isEqualTo(_1207_v0);
        let _index209 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_1208_v1).length));
        (_1208_v1)[_index209] = (_1242_v29).f20;
        let _1257_v39;
        let _init34 = function (_1258_i3) {
          return _dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))));
        };
        let _nw177 = Array((new BigNumber(20)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw177.length); _i0_34++) {
          _nw177[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1257_v39 = _nw177;
        let _1259_v40;
        _1259_v40 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1260_v41;
        _1260_v41 = _dafny.Set.fromElements(_1259_v40, _1259_v40, _1259_v40, _1259_v40);
        let _1261_v42;
        _1261_v42 = _dafny.MultiSet.fromElements(_1260_v41, _1260_v41);
        let _index210 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1257_v39).length));
        (_1257_v39)[_index210] = _1261_v42;
        let _index211 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1208_v1).length));
        let _index212 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_1208_v1).length));
        let _index213 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1257_v39).length));
        let _rhs206 = !(((_1207_v0).plus(new BigNumber(273))).isLessThan((_1248_v34)[_module.__default.safeIndex(new BigNumber(11), new BigNumber((_1248_v34).length))]));
        let _rhs207 = _1207_v0;
        let _rhs208 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.update(_1209_v2, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1242_v29).f20,_1211_v4)).length), new BigNumber((_1209_v2).length)), new _dafny.CodePoint('v'.codePointAt(0))), _1210_v3), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nhmbab"), _1209_v2));
        let _rhs209 = _1261_v42;
        let _lhs146 = _1208_v1;
        let _lhs147 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1208_v1).length));
        let _lhs148 = globalState;
        let _lhs149 = _1208_v1;
        let _lhs150 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_1208_v1).length));
        let _lhs151 = _1257_v39;
        let _lhs152 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1257_v39).length));
        _lhs146[_lhs147] = _rhs206;
        _lhs148.f1 = _rhs207;
        _lhs149[_lhs150] = _rhs208;
        _lhs151[_lhs152] = _rhs209;
      } else {
        let _1262_v43;
        _1262_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1211_v4,_1207_v0);
        let _1263_v44;
        _1263_v44 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1264_v45;
        _1264_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1263_v44,false);
        let _1265_v46;
        _1265_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1262_v43).length),_1264_v45);
        let _1266_v47;
        let _1267_v48;
        let _1268_v49;
        let _out44;
        let _out45;
        let _out46;
        let _outcollector21 = (_1242_v29).m3(_1265_v46, _1211_v4, globalState);
        _out44 = _outcollector21[0];
        _out45 = _outcollector21[1];
        _out46 = _outcollector21[2];
        _1266_v47 = _out44;
        _1267_v48 = _out45;
        _1268_v49 = _out46;
        let _1269_v50;
        _1269_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v0,(_1242_v29).f20);
        let _1270_v51;
        _1270_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1269_v50,_1211_v4);
        (globalState).f1 = (new BigNumber(((_1270_v51).Merge(_1270_v51)).length)).minus(_1207_v0);
        let _1271_v52;
        _1271_v52 = _dafny.Seq.of((_1242_v29).f20);
        let _1272_v53;
        _1272_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v0,_1271_v52);
        _1266_v47 = (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_1271_v52, (((_1272_v53).contains(_1207_v0)) ? ((_1272_v53).get(_1207_v0)) : (_1271_v52))), _module.__default.safeIndex(_1207_v0, new BigNumber((_dafny.Seq.Concat(_1271_v52, (((_1272_v53).contains(_1207_v0)) ? ((_1272_v53).get(_1207_v0)) : (_1271_v52)))).length)), (_1271_v52)[_module.__default.safeIndex(new BigNumber(-258), new BigNumber((_1271_v52).length))]))).cardinality())).isLessThan(_1207_v0);
        let _1273_v54;
        _1273_v54 = _dafny.Set.fromElements(_1262_v43);
        _1266_v47 = !((_1273_v54).contains(_1262_v43)) || ((_1207_v0).isLessThanOrEqualTo(_1207_v0));
        let _1274_v55;
        let _nw178 = Array((new BigNumber(28)).toNumber());
        _nw178[(_dafny.ZERO).toNumber()] = _1263_v44;
        _nw178[(_dafny.ONE).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(2)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(3)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(4)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(5)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(6)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(7)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _nw178[(new BigNumber(9)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(10)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(11)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(12)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(13)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(14)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(15)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw178[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw178[(new BigNumber(18)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(19)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(20)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(21)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(22)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(23)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(24)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(25)).toNumber()] = _1263_v44;
        _nw178[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw178[(new BigNumber(27)).toNumber()] = _1263_v44;
        _1274_v55 = _nw178;
        let _1275_v56;
        _1275_v56 = _module.D12.create_DC37((_dafny.ZERO).minus(_1207_v0), _1274_v55, _1263_v44);
        let _1276_v57;
        _1276_v57 = _module.D13.create_DC40(_1275_v56, !(_1211_v4), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), function (_1277_i4) {
  return new BigNumber(-227);
})).length), _1207_v0);
        (globalState).f1 = (_1276_v57).dtor_cf69;
      }
      let _1278_v58;
      let _nw179 = Array((new BigNumber(19)).toNumber()).fill([]);
      _1278_v58 = _nw179;
      _1278_v58 = _1278_v58;
      r0 = _1207_v0;
      r1 = (_dafny.Set.fromElements(_module.__default.fm0((_dafny.ZERO).minus(_1207_v0), globalState))).Intersect(_module.__default.fm30(globalState));
      return [r0, r1];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m13(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _1279_v0;
      _1279_v0 = true;
      if (_1279_v0) {
        let _1280_v1;
        let _init35 = ((_1281_v0) => function (_1282_i0) {
          return _module.D3.create_DC12(_module.D3.create_DC10(_dafny.Seq.UnicodeFromString("k"), _1281_v0, _1281_v0));
        })(_1279_v0);
        let _nw180 = Array((new BigNumber(8)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw180.length); _i0_35++) {
          _nw180[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1280_v1 = _nw180;
        let _1283_v2;
        _1283_v2 = _dafny.Seq.of(_1280_v1);
        let _1284_v3;
        _1284_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(897),_1283_v2);
        let _1285_v4;
        _1285_v4 = new BigNumber(303);
        _1284_v3 = (_1284_v3).update((_1285_v4).plus(_1285_v4), _dafny.Seq.update(_1283_v2, _module.__default.safeIndex((_dafny.ZERO).minus(_1285_v4), new BigNumber((_1283_v2).length)), _1280_v1));
        let _1286_v5;
        let _nw181 = new _module.C0();
        _nw181.__ctor();
        _1286_v5 = _nw181;
        let _1287_v6;
        let _nw182 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1287_v6 = _nw182;
        _1287_v6 = _1287_v6;
        let _1288_v7;
        let _nw183 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1288_v7 = _nw183;
        r0 = _1288_v7;
        let _1289_v8;
        let _nw184 = new _module.C0();
        _nw184.__ctor();
        _1289_v8 = _nw184;
      } else {
        let _1290_v9;
        _1290_v9 = new BigNumber(-401);
        let _1291_v10;
        _1291_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1290_v9,_1279_v0);
        let _1292_v11;
        _1292_v11 = _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("em")).length),_1279_v0));
        let _1293_v12;
        _1293_v12 = _dafny.Set.fromElements(new BigNumber(((_1291_v10).Merge((_1292_v11).dtor_cf53)).length));
        (globalState).f7 = new BigNumber((_1293_v12).length);
        let _1294_v13;
        _1294_v13 = new _dafny.CodePoint('j'.codePointAt(0));
        _1294_v13 = _1294_v13;
        let _1295_v14;
        _1295_v14 = _dafny.Seq.of(new BigNumber((_module.__default.fm40(_1279_v0, true, globalState)).length));
        let _1296_v15;
        _1296_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1279_v0,true);
        let _1297_v16;
        _1297_v16 = _module.D3.create_DC11(new BigNumber((_1295_v14).length), _1296_v15);
        let _1298_v17;
        _1298_v17 = _module.D3.create_DC12(_1297_v16);
        let _1299_v18;
        _1299_v18 = _module.D3.create_DC12(_1298_v17);
        _1299_v18 = _module.D3.create_DC12(_1297_v16);
        let _1300_v19;
        _1300_v19 = _dafny.Seq.UnicodeFromString("vawvwkwoa");
        _1300_v19 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1300_v19, _1300_v19), _1300_v19);
        let _1301_v20;
        _1301_v20 = _dafny.MultiSet.fromElements(_1279_v0, _1279_v0);
        _1279_v0 = ((_module.__default.fm41(globalState)).update(!(_1279_v0), _module.__default.abs(_1290_v9))).IsProperSubsetOf(_1301_v20);
      }
      let _1302_v21;
      _1302_v21 = _dafny.Seq.UnicodeFromString("m");
      let _1303_v22;
      _1303_v22 = _dafny.Seq.of(_1279_v0, _1279_v0);
      let _1304_v23;
      _1304_v23 = new BigNumber(201);
      let _1305_v24;
      let _nw185 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1305_v24 = _nw185;
      let _1306_v25;
      _1306_v25 = _module.D1.create_DC4(_1304_v23, _1305_v24);
      let _1307_v26;
      _1307_v26 = _module.D14.create_DC42(_1306_v25, _1303_v22, _1279_v0);
      let _1308_v27;
      _1308_v27 = _dafny.MultiSet.fromElements(_1303_v22, _dafny.Seq.Concat(_1303_v22, _1303_v22), _module.__default.fm40(_1279_v0, _1279_v0, globalState), _dafny.Seq.Concat(_1303_v22, _1303_v22), _dafny.Seq.of((_1307_v26).dtor_cf74));
      let _index214 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1305_v24).length));
      (_1305_v24)[_index214] = _1304_v23;
      let _1309_v28;
      _1309_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1279_v0,_module.__default.fm0(_1304_v23, globalState));
      let _index215 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1305_v24).length));
      let _rhs210 = _dafny.Seq.Concat(_1302_v21, _dafny.Seq.UnicodeFromString("ys"));
      let _rhs211 = _1308_v27;
      let _rhs212 = _1304_v23;
      let _rhs213 = _dafny.Seq.IsProperPrefixOf(_1303_v22, _dafny.Seq.update(_1303_v22, _module.__default.safeIndex((_dafny.ZERO).minus(_1304_v23), new BigNumber((_1303_v22).length)), (((_1309_v28).contains(_1279_v0)) ? ((_1309_v28).get(_1279_v0)) : (_1279_v0))));
      let _rhs214 = _module.__default.fm0(_1304_v23, globalState);
      let _lhs153 = _1305_v24;
      let _lhs154 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1305_v24).length));
      _1302_v21 = _rhs210;
      _1308_v27 = _rhs211;
      _lhs153[_lhs154] = _rhs212;
      _1279_v0 = _rhs213;
      _1279_v0 = _rhs214;
      let _1310_v29;
      _1310_v29 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1279_v0,true));
      let _1311_v30;
      _1311_v30 = _dafny.Map.Empty.slice().updateUnsafe((p0).fm9(_1302_v21, _1310_v29, globalState),_1279_v0);
      _1311_v30 = (_1311_v30).update(_1304_v23, _1279_v0);
      let _hi6 = _1304_v23;
      for (let _1312_i1 = _1304_v23; _1312_i1.isLessThan(_hi6); _1312_i1 = _1312_i1.plus(_dafny.ONE)) {
        let _1313_v31;
        let _nw186 = Array((new BigNumber(12)).toNumber()).fill(_module.D1.Default());
        _1313_v31 = _nw186;
        let _1314_v32;
        _1314_v32 = _dafny.Seq.of(_1313_v31, _1313_v31, _1313_v31);
        let _1315_v33;
        _1315_v33 = _dafny.Seq.of(_1312_i1, _1304_v23, _1312_i1, _1304_v23);
        let _1316_v34;
        _1316_v34 = _module.D8.create_DC27(_1279_v0, _1314_v32, _1315_v33, _1312_i1);
        let _index216 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1305_v24).length));
        (_1305_v24)[_index216] = (_1316_v34).dtor_cf46;
        (globalState).f1 = ((_1315_v33)[_module.__default.safeIndex(new BigNumber(-856), new BigNumber((_1315_v33).length))]).minus(new BigNumber(96));
        let _1317_v35;
        let _nw187 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1317_v35 = _nw187;
        let _index217 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length));
        (_1317_v35)[_index217] = _1279_v0;
        let _index218 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length));
        (_1317_v35)[_index218] = _1279_v0;
        let _1318_v36;
        _1318_v36 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1319_v37;
        let _nw188 = Array((new BigNumber(18)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
        _nw188[(_dafny.ONE).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(2)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw188[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw188[(new BigNumber(5)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(6)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(7)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(8)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(9)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(10)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(11)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(12)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(13)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(14)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(15)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(16)).toNumber()] = _1318_v36;
        _nw188[(new BigNumber(17)).toNumber()] = _1318_v36;
        _1319_v37 = _nw188;
        let _1320_v38;
        _1320_v38 = _dafny.Seq.of(_1319_v37, _1319_v37, _1319_v37, _1319_v37, _1319_v37);
        let _1321_v39;
        _1321_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1279_v0,_module.__default.fm50(_1279_v0, _1312_i1, globalState));
        let _1322_v40;
        _1322_v40 = _dafny.Set.fromElements((_1317_v35)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length))], (_1317_v35)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length))]);
        let _index219 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length));
        let _rhs215 = !(!(!(_1321_v39).contains(_1279_v0))) || ((((_1317_v35)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length))]) ? ((p0).fm8(new BigNumber((_1320_v38).length), !(_1279_v0), _1279_v0, globalState)) : ((_1317_v35)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length))])));
        let _rhs216 = (_1322_v40).Union(_dafny.Set.fromElements((_1317_v35)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length))], _1279_v0));
        let _lhs155 = _1317_v35;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1317_v35).length));
        let _lhs157 = globalState;
        _lhs155[_lhs156] = _rhs215;
        _lhs157.f3 = _rhs216;
      }
      let _1323_v41;
      _1323_v41 = new _dafny.CodePoint('x'.codePointAt(0));
      let _1324_v42;
      let _nw189 = Array((new BigNumber(9)).toNumber());
      _nw189[(_dafny.ZERO).toNumber()] = _1323_v41;
      _nw189[(_dafny.ONE).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(2)).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(3)).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(4)).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(5)).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(6)).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(7)).toNumber()] = _1323_v41;
      _nw189[(new BigNumber(8)).toNumber()] = _1323_v41;
      _1324_v42 = _nw189;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1324_v42).length))) {
        let _1325_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1325_i2)) && ((_1325_i2).isLessThan(new BigNumber((_1324_v42).length))))) {
          (_1324_v42)[(_1325_i2)] = _1323_v41;
        }
      }
      let _1326_v43;
      let _nw190 = new _module.C1();
      _nw190.__ctor(_1279_v0);
      _1326_v43 = _nw190;
      let _1327_v44;
      _1327_v44 = _dafny.Seq.of(_1326_v43, _1326_v43);
      let _1328_v45;
      let _nw191 = Array((new BigNumber(18)).toNumber());
      _nw191[(_dafny.ZERO).toNumber()] = _1326_v43;
      _nw191[(_dafny.ONE).toNumber()] = (_1327_v44)[_module.__default.safeIndex(_1304_v23, new BigNumber((_1327_v44).length))];
      _nw191[(new BigNumber(2)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(3)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(4)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(5)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(6)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(7)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(8)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(9)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(10)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(11)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(12)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(13)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(14)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(15)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(16)).toNumber()] = _1326_v43;
      _nw191[(new BigNumber(17)).toNumber()] = _1326_v43;
      _1328_v45 = _nw191;
      let _1329_v46;
      _1329_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_1305_v24)[_module.__default.safeIndex(new BigNumber(587), new BigNumber((_1305_v24).length))], new BigNumber(884)),_1328_v45);
      _1329_v46 = (_1329_v46).update(_1304_v23, ((_1279_v0) ? (_1328_v45) : (_1328_v45)));
      r0 = _1324_v42;
      r1 = _1279_v0;
      return [r0, r1];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm22(p0, p1, globalState) {
      let _this = this;
      return !(((true) ? (new BigNumber(202)) : (new BigNumber((_dafny.Seq.UnicodeFromString("n")).length)))).isEqualTo(new BigNumber(133));
    };
    fm23(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (true) && ((_dafny.MultiSet.fromElements(new BigNumber(520), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-300),false)).length))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,true), _dafny.Map.Empty.slice().updateUnsafe(false,true))).length))));
    };
    m11(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _1330_v0;
      _1330_v0 = new BigNumber(-465);
      let _1331_v1;
      _1331_v1 = _dafny.Seq.of(_1330_v0);
      let _1332_v2;
      _1332_v2 = _dafny.Seq.of(_1330_v0, new BigNumber((_1331_v1).length));
      let _1333_v3;
      _1333_v3 = _dafny.MultiSet.fromElements(new BigNumber((_1331_v1).length), _1330_v0, new BigNumber(962));
      r0 = (_1333_v3).IsProperSubsetOf((_1333_v3).Union((_dafny.MultiSet.fromElements(_1330_v0)).update((_1332_v2)[_module.__default.safeIndex(_1330_v0, new BigNumber((_1332_v2).length))], _module.__default.abs(_1330_v0))));
      let _1334_v4;
      _1334_v4 = _dafny.Seq.UnicodeFromString("rcvrbcdrn");
      let _1335_v5;
      _1335_v5 = _dafny.Seq.of(_module.D3.create_DC10(_1334_v4, false, p0));
      let _1336_v6;
      _1336_v6 = _module.D3.create_DC12((_1335_v5)[_module.__default.safeIndex(_1330_v0, new BigNumber((_1335_v5).length))]);
      let _1337_i0;
      _1337_i0 = _dafny.ZERO;
      L6: {
        let _pat_let_tv28 = p0;
        let _pat_let_tv29 = p0;
        let _pat_let_tv30 = p0;
        let _pat_let_tv31 = p0;
        while (function (_source26) {
          if (_source26.is_DC10) {
            let _1343___mcc_h0 = (_source26).cf14;
            let _1344___mcc_h1 = (_source26).cf15;
            let _1345___mcc_h2 = (_source26).cf16;
            let _1346_cf16 = _1345___mcc_h2;
            let _1347_cf15 = _1344___mcc_h1;
            let _1348_cf14 = _1343___mcc_h0;
            return false;
          } else if (_source26.is_DC11) {
            let _1349___mcc_h3 = (_source26).cf17;
            let _1350___mcc_h4 = (_source26).cf18;
            let _1351_cf18 = _1350___mcc_h4;
            let _1352_cf17 = _1349___mcc_h3;
            return (_pat_let_tv28) === (true);
          } else if (_source26.is_DC9) {
            let _1353___mcc_h5 = (_source26).cf13;
            let _1354_cf13 = _1353___mcc_h5;
            return _pat_let_tv29;
          } else {
            let _1355___mcc_h6 = (_source26).cf19;
            let _1356_cf19 = _1355___mcc_h6;
            return !(_pat_let_tv30) || (_pat_let_tv31);
          }
        }(_1336_v6)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1337_i0)) {
              break L6;
            }
            _1337_i0 = (_1337_i0).plus(_dafny.ONE);
            let _1338_v7;
            let _nw192 = Array((new BigNumber(18)).toNumber()).fill(false);
            _1338_v7 = _nw192;
            let _index220 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length));
            (_1338_v7)[_index220] = _dafny.Seq.IsProperPrefixOf(_1334_v4, _1334_v4);
            let _index221 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length));
            (_1338_v7)[_index221] = p0;
            let _1339_v8;
            _1339_v8 = _dafny.Seq.of((_1338_v7)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length))]);
            let _1340_v9;
            let _nw193 = Array((new BigNumber(11)).toNumber()).fill(_module.D2.Default());
            _1340_v9 = _nw193;
            let _1341_v10;
            _1341_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1339_v8, _1339_v8)).length),_1340_v9);
            let _1342_v11;
            _1342_v11 = _module.D6.create_DC19(_1341_v10);
            _1341_v10 = (_1341_v10).Merge((_1342_v11).dtor_cf31);
            let _index222 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length));
            (_1338_v7)[_index222] = (_1338_v7)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length))];
            r0 = !(((_1338_v7)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length))]) && ((_1338_v7)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_1338_v7).length))])) || (p0);
          }
        }
      }
      let _1357_v12;
      _1357_v12 = _module.D6.create_DC21((_this).fm23(p0, _1330_v0, p0, _1330_v0, globalState), p0, p0, new BigNumber((_1331_v1).length));
      let _source27 = _1357_v12;
      if (_source27.is_DC20) {
        let _1358___mcc_h7 = (_source27).cf32;
        let _1359___mcc_h8 = (_source27).cf33;
        let _1360___mcc_h9 = (_source27).cf34;
        let _1361___mcc_h10 = (_source27).cf35;
        let _1362_cf35 = _1361___mcc_h10;
        let _1363_cf34 = _1360___mcc_h9;
        let _1364_cf33 = _1359___mcc_h8;
        let _1365_cf32 = _1358___mcc_h7;
        let _1366_v13;
        let _init36 = ((_1367_cf33, _1368_p0, _1369_cf32) => function (_1370_i1) {
          return _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1367_cf33, _1368_p0),_1369_cf32));
        })(_1364_cf33, p0, _1365_cf32);
        let _nw194 = Array((new BigNumber(9)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw194.length); _i0_36++) {
          _nw194[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1366_v13 = _nw194;
        let _1371_v14;
        _1371_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1364_cf33),_1363_cf34);
        let _index223 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1366_v13).length));
        (_1366_v13)[_index223] = _module.D4.create_DC13(_1371_v14);
        let _1372_v15;
        _1372_v15 = _module.D4.create_DC13(_1371_v14);
        let _1373_v16;
        let _nw195 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1373_v16 = _nw195;
        let _1374_v17;
        _1374_v17 = _module.D0.create_DC2(_1365_cf32, _1373_v16);
        let _index224 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1366_v13).length));
        let _rhs217 = _1372_v15;
        let _rhs218 = _1330_v0;
        let _rhs219 = (_1331_v1)[_module.__default.safeIndex(_1330_v0, new BigNumber((_1331_v1).length))];
        let _rhs220 = (_1374_v17).dtor_cf2;
        let _lhs158 = _1366_v13;
        let _lhs159 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1366_v13).length));
        let _lhs160 = globalState;
        let _lhs161 = globalState;
        _lhs158[_lhs159] = _rhs217;
        _lhs160.f1 = _rhs218;
        _lhs161.f7 = _rhs219;
        _1363_cf34 = _rhs220;
        let _1375_v18;
        _1375_v18 = _dafny.Seq.of(_1332_v2, _1331_v1);
        let _1376_v19;
        let _init37 = ((_1377_cf35) => function (_1378_i2) {
          return (_1378_i2).multipliedBy(_1377_cf35);
        })(_1362_cf35);
        let _nw196 = Array((new BigNumber(29)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw196.length); _i0_37++) {
          _nw196[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1376_v19 = _nw196;
        let _1379_v20;
        _1379_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1375_v18).length),_1376_v19);
        let _1380_v21;
        _1380_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(193),_1379_v20);
        let _1381_v22;
        _1381_v22 = _dafny.Seq.of(_1379_v20, _1379_v20, _1379_v20, _1379_v20);
        _1380_v21 = (_1380_v21).update(_1362_cf35, ((_1365_cf32) ? (_1379_v20) : ((_1381_v22)[_module.__default.safeIndex((_dafny.ZERO).minus(_1330_v0), new BigNumber((_1381_v22).length))])));
        let _1382_v23;
        _1382_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1373_v16,_1336_v6);
        _1382_v23 = (_1382_v23).update(_1373_v16, _1336_v6);
        let _1383_v24;
        _1383_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1364_cf33,(_1362_cf35).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1330_v0,_1364_cf33)).length)));
        let _1384_v25;
        _1384_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1362_cf35,p0);
        _1383_v24 = (_1383_v24).update((((_1384_v25).contains(new BigNumber(993))) ? ((_1384_v25).get(new BigNumber(993))) : (_module.__default.fm0(_1330_v0, globalState))), _1362_cf35);
      } else if (_source27.is_DC21) {
        let _1385___mcc_h11 = (_source27).cf36;
        let _1386___mcc_h12 = (_source27).cf37;
        let _1387___mcc_h13 = (_source27).cf38;
        let _1388___mcc_h14 = (_source27).cf39;
        let _1389_cf39 = _1388___mcc_h14;
        let _1390_cf38 = _1387___mcc_h13;
        let _1391_cf37 = _1386___mcc_h12;
        let _1392_cf36 = _1385___mcc_h11;
        let _1393_v26;
        _1393_v26 = _dafny.Seq.of(!(_1390_cf38));
        _1392_cf36 = (((_this).fm23(p0, _1330_v0, _1390_cf38, new BigNumber((_module.__default.fm24(globalState)).cardinality()), globalState)) ? (((_dafny.ZERO).minus((_dafny.ZERO).minus(_1389_cf39))).isLessThanOrEqualTo(_1330_v0)) : ((_1393_v26)[_module.__default.safeIndex(_1389_cf39, new BigNumber((_1393_v26).length))]));
        let _1394_v27;
        _1394_v27 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1395_v28;
        let _nw197 = Array((new BigNumber(3)).toNumber());
        _nw197[(_dafny.ZERO).toNumber()] = _module.__default.fm4(_1391_cf37, true, _1392_cf36, globalState);
        _nw197[(_dafny.ONE).toNumber()] = _1394_v27;
        _nw197[(new BigNumber(2)).toNumber()] = _1394_v27;
        _1395_v28 = _nw197;
        let _index225 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1395_v28).length));
        (_1395_v28)[_index225] = _1394_v27;
        let _index226 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1395_v28).length));
        let _rhs221 = _1390_cf38;
        let _rhs222 = (_1334_v4)[_module.__default.safeIndex(new BigNumber(-412), new BigNumber((_1334_v4).length))];
        let _lhs162 = _1395_v28;
        let _lhs163 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1395_v28).length));
        _1391_cf37 = _rhs221;
        _lhs162[_lhs163] = _rhs222;
        let _1396_v29;
        let _nw198 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _1396_v29 = _nw198;
        let _1397_v30;
        _1397_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1390_cf38,_1390_cf38);
        let _index227 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1396_v29).length));
        (_1396_v29)[_index227] = _1397_v30;
        let _index228 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1396_v29).length));
        let _rhs223 = (_1397_v30).Merge((_1397_v30).update(false, p0));
        let _rhs224 = !(false) || (_1392_cf36);
        let _rhs225 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("echwct"), _dafny.Seq.UnicodeFromString("fckgecd"));
        let _lhs164 = _1396_v29;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1396_v29).length));
        _lhs164[_lhs165] = _rhs223;
        _1391_cf37 = _rhs224;
        r0 = _rhs225;
        let _1398_v31;
        let _nw199 = Array((new BigNumber(4)).toNumber());
        _nw199[(_dafny.ZERO).toNumber()] = _1390_cf38;
        _nw199[(_dafny.ONE).toNumber()] = _1391_cf37;
        _nw199[(new BigNumber(2)).toNumber()] = p0;
        _nw199[(new BigNumber(3)).toNumber()] = p0;
        _1398_v31 = _nw199;
        let _1399_v32;
        _1399_v32 = _dafny.Seq.of(_1398_v31, _1398_v31);
        let _1400_v33;
        let _nw200 = Array((new BigNumber(27)).toNumber());
        _nw200[(_dafny.ZERO).toNumber()] = _1398_v31;
        _nw200[(_dafny.ONE).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(2)).toNumber()] = ((_1391_cf37) ? (_1398_v31) : (_1398_v31));
        _nw200[(new BigNumber(3)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(4)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(5)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(6)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(7)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(8)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(9)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(10)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(11)).toNumber()] = (_1399_v32)[_module.__default.safeIndex(_1389_cf39, new BigNumber((_1399_v32).length))];
        _nw200[(new BigNumber(12)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(13)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(14)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(15)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(16)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(17)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(18)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(19)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(20)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(21)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(22)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(23)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(24)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(25)).toNumber()] = _1398_v31;
        _nw200[(new BigNumber(26)).toNumber()] = _1398_v31;
        _1400_v33 = _nw200;
        let _1401_v34;
        _1401_v34 = _dafny.Seq.of(_1400_v33);
        let _1402_v35;
        _1402_v35 = _dafny.MultiSet.fromElements(!(p0));
        let _rhs226 = (_1401_v34)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1330_v0, _1389_cf39)), new BigNumber((_1401_v34).length))];
        let _rhs227 = p0;
        let _rhs228 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_1334_v4, ((_1391_cf37) ? (_1334_v4) : (_1334_v4))), _module.__default.safeIndex(new BigNumber(((_1402_v35).Intersect(_dafny.MultiSet.fromElements(_1391_cf37))).cardinality()), new BigNumber((_dafny.Seq.Concat(_1334_v4, ((_1391_cf37) ? (_1334_v4) : (_1334_v4)))).length)), new _dafny.CodePoint('m'.codePointAt(0))), _module.__default.safeIndex(_1330_v0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1334_v4, ((_1391_cf37) ? (_1334_v4) : (_1334_v4))), _module.__default.safeIndex(new BigNumber(((_1402_v35).Intersect(_dafny.MultiSet.fromElements(_1391_cf37))).cardinality()), new BigNumber((_dafny.Seq.Concat(_1334_v4, ((_1391_cf37) ? (_1334_v4) : (_1334_v4)))).length)), new _dafny.CodePoint('m'.codePointAt(0)))).length)), (_1395_v28)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_1395_v28).length))])).length);
        let _lhs166 = globalState;
        _1400_v33 = _rhs226;
        _1391_cf37 = _rhs227;
        _lhs166.f1 = _rhs228;
      } else {
        let _1403___mcc_h15 = (_source27).cf31;
        let _1404_cf31 = _1403___mcc_h15;
        let _1405_v36;
        _1405_v36 = _module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-533)), function (_1406_i3) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}));
        let _1407_v37;
        _1407_v37 = _dafny.Map.Empty.slice().updateUnsafe(false,_1405_v36);
        let _1408_v38;
        _1408_v38 = _dafny.Map.Empty.slice().updateUnsafe(false,_1407_v37);
        let _pat_let_tv32 = _1334_v4;
        if (!((_dafny.Map.Empty.slice().updateUnsafe(p0,function (_pat_let23_0) {
          return function (_1409_dt__update__tmp_h0) {
            return function (_pat_let24_0) {
              return function (_1410_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_1410_dt__update_hcf1_h0);
              }(_pat_let24_0);
            }(_pat_let_tv32);
          }(_pat_let23_0);
        }(_1405_v36))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1405_v36))).equals(((p0) ? (((((_1408_v38).contains(p0)) ? ((_1408_v38).get(p0)) : (_1407_v37))).update(p0, _1405_v36)) : (_1407_v37)))) {
          let _1411_v39;
          _1411_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), function (_1412_i4) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }), _1334_v4),_1330_v0);
          _1411_v39 = (_1411_v39).update(_1334_v4, _module.__default.fm1(globalState));
          let _1413_v40;
          let _nw201 = Array((new BigNumber(8)).toNumber()).fill([]);
          _1413_v40 = _nw201;
          let _1414_v41;
          let _nw202 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1414_v41 = _nw202;
          let _index229 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1413_v40).length));
          (_1413_v40)[_index229] = _1414_v41;
          let _1415_v42;
          _1415_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(987),_1414_v41);
          let _1416_v43;
          _1416_v43 = _dafny.Seq.of(_1414_v41);
          let _index230 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1413_v40).length));
          (_1413_v40)[_index230] = (((_1415_v42).contains((_1330_v0).multipliedBy(_1330_v0))) ? ((_1415_v42).get((_1330_v0).multipliedBy(_1330_v0))) : ((_1416_v43)[_module.__default.safeIndex(new BigNumber(-802), new BigNumber((_1416_v43).length))]));
          let _1417_v44;
          let _nw203 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1417_v44 = _nw203;
          _1417_v44 = _1417_v44;
          let _1418_v45;
          _1418_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1330_v0);
          let _1419_v46;
          _1419_v46 = _module.D2.create_DC7(new BigNumber((_1418_v45).length), p0, _1330_v0);
          r0 = !(p0) || ((_1419_v46).dtor_cf10);
          let _1420_v47;
          _1420_v47 = _dafny.Seq.of(true);
          let _arr2 = (_1413_v40)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1413_v40).length))];
          let _index231 = _module.__default.safeIndex(new BigNumber(870), new BigNumber(((_1413_v40)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1413_v40).length))]).length));
          _arr2[_index231] = new BigNumber((_1420_v47).length);
          let _1421_v48;
          _1421_v48 = _dafny.Set.fromElements(_1420_v47, _1420_v47, _dafny.Seq.Concat(_1420_v47, _dafny.Seq.of(p0)));
          let _arr3 = (_1413_v40)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1413_v40).length))];
          let _index232 = _module.__default.safeIndex(new BigNumber(870), new BigNumber(((_1413_v40)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1413_v40).length))]).length));
          _arr3[_index232] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1421_v48).length)));
        } else {
          let _1422_v49;
          let _nw204 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _1422_v49 = _nw204;
          let _1423_v50;
          _1423_v50 = _dafny.Seq.of(p0);
          let _index233 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_1422_v49).length));
          (_1422_v49)[_index233] = _dafny.Seq.Concat(_1423_v50, _dafny.Seq.of(p0, p0));
          let _index234 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_1422_v49).length));
          (_1422_v49)[_index234] = _1423_v50;
          let _1424_v51;
          _1424_v51 = _dafny.Seq.of(_1334_v4, _module.__default.fm25(new BigNumber(410), true, new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()), globalState));
          let _1425_v52;
          _1425_v52 = _dafny.Set.fromElements(_1330_v0, _1330_v0, _1330_v0, _1330_v0);
          let _rhs229 = p0;
          let _rhs230 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1330_v0), _1330_v0);
          let _rhs231 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), ((_1426_v4) => function (_1427_i5) {
            return _1426_v4;
          })(_1334_v4));
          let _rhs232 = (!_dafny.areEqual(_1334_v4, _1334_v4)) === (!(p0));
          let _rhs233 = (_dafny.Set.fromElements(_1330_v0, new BigNumber(648), new BigNumber(510))).IsDisjointFrom((_1425_v52).Difference(_module.__default.fm26(globalState)));
          let _lhs167 = globalState;
          r0 = _rhs229;
          _lhs167.f0 = _rhs230;
          _1424_v51 = _rhs231;
          r0 = _rhs232;
          r0 = _rhs233;
          let _1428_v53;
          _1428_v53 = _module.D4.create_DC14(_1330_v0, _1330_v0);
          (_this).m12(_module.__default.safeModuloInt(_1330_v0, (_1428_v53).dtor_cf21), globalState);
          let _1429_v54;
          _1429_v54 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(p0, p0));
          _1423_v50 = _dafny.Seq.Concat((((_1429_v54).contains(p0)) ? ((_1429_v54).get(p0)) : ((_1422_v49)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_1422_v49).length))])), _dafny.Seq.of(false));
          r0 = !(p0) || (p0);
        }
        (_this).m12(_1330_v0, globalState);
        let _1430_v55;
        _1430_v55 = _module.D6.create_DC20(p0, p0, p0, ((((_1333_v3).contains(_1330_v0)) ? ((_1333_v3).get(_1330_v0)) : (_1330_v0))).minus(_1330_v0));
        let _source28 = _1430_v55;
        if (_source28.is_DC20) {
          let _1431___mcc_h16 = (_source28).cf32;
          let _1432___mcc_h17 = (_source28).cf33;
          let _1433___mcc_h18 = (_source28).cf34;
          let _1434___mcc_h19 = (_source28).cf35;
          let _1435_cf35 = _1434___mcc_h19;
          let _1436_cf34 = _1433___mcc_h18;
          let _1437_cf33 = _1432___mcc_h17;
          let _1438_cf32 = _1431___mcc_h16;
          let _1439_v56;
          _1439_v56 = new _dafny.CodePoint('v'.codePointAt(0));
          let _1440_v57;
          _1440_v57 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_1331_v1, _1332_v2),_1439_v56);
          _1440_v57 = (_1440_v57).update(true, new _dafny.CodePoint('b'.codePointAt(0)));
          let _1441_v58;
          let _nw205 = new _module.C3();
          _nw205.__ctor(p0, p0);
          _1441_v58 = _nw205;
          let _1442_v59;
          let _nw206 = Array((new BigNumber(21)).toNumber()).fill(_module.D4.Default());
          _1442_v59 = _nw206;
          let _1443_v60;
          _1443_v60 = _module.D4.create_DC15((_1331_v1)[_module.__default.safeIndex(_1435_cf35, new BigNumber((_1331_v1).length))], _1333_v3, _1439_v56, _1435_cf35);
          let _index235 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_1442_v59).length));
          (_1442_v59)[_index235] = _1443_v60;
          let _index236 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_1442_v59).length));
          (_1442_v59)[_index236] = _1443_v60;
          _1334_v4 = _1334_v4;
        } else if (_source28.is_DC21) {
          let _1444___mcc_h20 = (_source28).cf36;
          let _1445___mcc_h21 = (_source28).cf37;
          let _1446___mcc_h22 = (_source28).cf38;
          let _1447___mcc_h23 = (_source28).cf39;
          let _1448_cf39 = _1447___mcc_h23;
          let _1449_cf38 = _1446___mcc_h22;
          let _1450_cf37 = _1445___mcc_h21;
          let _1451_cf36 = _1444___mcc_h20;
          _1451_cf36 = p0;
          let _1452_v61;
          _1452_v61 = _dafny.Seq.of(_1332_v2);
          let _1453_v62;
          _1453_v62 = _module.D13.create_DC39(_1330_v0, p0);
          let _1454_v63;
          let _nw207 = Array((new BigNumber(14)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = _1332_v2;
          _nw207[(_dafny.ONE).toNumber()] = ((p0) ? (_dafny.Seq.of(_1448_cf39, _1448_cf39, (_dafny.ZERO).minus(_1330_v0))) : (_1332_v2));
          _nw207[(new BigNumber(2)).toNumber()] = (_1452_v61)[_module.__default.safeIndex(new BigNumber((_1334_v4).length), new BigNumber((_1452_v61).length))];
          _nw207[(new BigNumber(3)).toNumber()] = _1331_v1;
          _nw207[(new BigNumber(4)).toNumber()] = _1331_v1;
          _nw207[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-211)), ((_1455_cf39) => function (_1456_i6) {
            return _1455_cf39;
          })(_1448_cf39)), _module.__default.safeIndex(_1330_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-211)), ((_1457_cf39) => function (_1458_i6) {
            return _1457_cf39;
          })(_1448_cf39))).length)), _1448_cf39), _1332_v2);
          _nw207[(new BigNumber(6)).toNumber()] = _1332_v2;
          _nw207[(new BigNumber(7)).toNumber()] = _1331_v1;
          _nw207[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1332_v2, _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1332_v2).length)), _1330_v0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-687)), function (_1459_i7) {
            return new BigNumber(774);
          }));
          _nw207[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_module.__default.fm1(globalState));
          _nw207[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-811)))), _1332_v2);
          _nw207[(new BigNumber(11)).toNumber()] = _1332_v2;
          _nw207[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), ((_1460_v4) => function (_1461_i8) {
            return (_dafny.ZERO).minus(new BigNumber((_1460_v4).length));
          })(_1334_v4));
          _nw207[(new BigNumber(13)).toNumber()] = _dafny.Seq.of((_1453_v62).dtor_cf65);
          _1454_v63 = _nw207;
          _1454_v63 = _1454_v63;
          let _1462_v64;
          _1462_v64 = _dafny.Seq.of(_1450_cf37);
          let _1463_v65;
          _1463_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1333_v3,_dafny.Seq.Concat(_1462_v64, _1462_v64));
          _1463_v65 = _1463_v65;
          _1450_cf37 = !(_1451_cf36);
        } else {
          let _1464___mcc_h24 = (_source28).cf31;
          let _1465_cf31 = _1464___mcc_h24;
          r0 = p0;
          let _1466_v66;
          let _init38 = function (_1467_i9) {
            return _module.__default.safeModuloInt(_1467_i9, new BigNumber(22));
          };
          let _nw208 = Array((new BigNumber(21)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw208.length); _i0_38++) {
            _nw208[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1466_v66 = _nw208;
          let _1468_v67;
          _1468_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1333_v3,_1466_v66);
          _1468_v67 = (_1468_v67).update(_dafny.MultiSet.FromArray(_1331_v1), _1466_v66);
          let _1469_v68;
          _1469_v68 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_module.__default.fm1(globalState));
          _1469_v68 = (_1469_v68).update(p0, _1330_v0);
          r0 = p0;
        }
        r0 = p0;
      }
      let _1470_v69;
      let _init39 = ((_1471_p0) => function (_1472_i10) {
        return _1471_p0;
      })(p0);
      let _nw209 = Array((new BigNumber(27)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw209.length); _i0_39++) {
        _nw209[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1470_v69 = _nw209;
      let _index237 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
      (_1470_v69)[_index237] = p0;
      let _1473_v70;
      _1473_v70 = _module.D2.create_DC7(_1330_v0, p0, _1330_v0);
      let _1474_v71;
      _1474_v71 = _module.D2.create_DC8(_1473_v70);
      let _pat_let_tv33 = p0;
      let _pat_let_tv34 = p0;
      let _pat_let_tv35 = p0;
      let _pat_let_tv36 = p0;
      let _index238 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
      (_1470_v69)[_index238] = function (_source29) {
        if (_source29.is_DC7) {
          let _1475___mcc_h25 = (_source29).cf9;
          let _1476___mcc_h26 = (_source29).cf10;
          let _1477___mcc_h27 = (_source29).cf11;
          let _1478_cf11 = _1477___mcc_h27;
          let _1479_cf10 = _1476___mcc_h26;
          let _1480_cf9 = _1475___mcc_h25;
          return _pat_let_tv33;
        } else if (_source29.is_DC6) {
          let _1481___mcc_h28 = (_source29).cf8;
          let _1482_cf8 = _1481___mcc_h28;
          return (_pat_let_tv34) || (_pat_let_tv35);
        } else {
          let _1483___mcc_h29 = (_source29).cf12;
          let _1484_cf12 = _1483___mcc_h29;
          return _pat_let_tv36;
        }
      }(_1474_v71);
      let _1485_i11;
      _1485_i11 = _dafny.ZERO;
      L7: {
        while ((_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))]) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1485_i11)) {
              break L7;
            }
            _1485_i11 = (_1485_i11).plus(_dafny.ONE);
            let _1486_v72;
            _1486_v72 = _dafny.Set.fromElements(_1330_v0, _1330_v0);
            let _1487_v73;
            _1487_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1330_v0,p0);
            let _1488_v75;
            _1488_v75 = _dafny.Seq.of(_1486_v72, _1486_v72, (function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of (_1487_v73).Keys.Elements) {
                let _1489_v74 = _compr_27;
                if ((_1487_v73).contains(_1489_v74)) {
                  _coll27.add((_1489_v74).plus(new BigNumber(335)));
                }
              }
              return _coll27;
            }()).Intersect(_1486_v72));
            let _1490_v76;
            _1490_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
            _1488_v75 = (((((_1490_v76).contains((_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))])) ? ((_1490_v76).get((_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))])) : ((_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))]))) ? (_1488_v75) : (_1488_v75));
            r0 = !_dafny.areEqual(_1334_v4, _1334_v4);
            let _1491_v77;
            let _nw210 = new _module.C1();
            _nw210.__ctor(((_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))]) === ((_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))]));
            _1491_v77 = _nw210;
            let _1492_v78;
            _1492_v78 = _dafny.Seq.of(false);
            let _1493_v79;
            _1493_v79 = _dafny.MultiSet.fromElements((_1492_v78)[_module.__default.safeIndex(_1330_v0, new BigNumber((_1492_v78).length))], (_1487_v73).contains(new BigNumber((_1334_v4).length)));
            let _index239 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
            let _rhs234 = _module.__default.fm41(globalState);
            let _rhs235 = new BigNumber(((_1487_v73).Merge(_1487_v73)).length);
            let _rhs236 = (_1470_v69)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length))];
            let _rhs237 = (_1486_v72).IsProperSubsetOf(_1486_v72);
            let _lhs168 = _1470_v69;
            let _lhs169 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
            _1493_v79 = _rhs234;
            _1330_v0 = _rhs235;
            _lhs168[_lhs169] = _rhs236;
            r0 = _rhs237;
          }
        }
      }
      let _index240 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
      let _index241 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
      let _rhs238 = _module.__default.fm0(_1330_v0, globalState);
      let _rhs239 = p0;
      let _lhs170 = _1470_v69;
      let _lhs171 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
      let _lhs172 = _1470_v69;
      let _lhs173 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1470_v69).length));
      _lhs170[_lhs171] = _rhs238;
      _lhs172[_lhs173] = _rhs239;
      r0 = !((_this).fm22(_module.__default.safeDivisionInt(_1330_v0, new BigNumber(739)), p0, globalState));
      r1 = _1330_v0;
      let _1494_v80;
      let _nw211 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1494_v80 = _nw211;
      r2 = _1494_v80;
      return [r0, r1, r2];
    }
    m12(p0, globalState) {
      let _this = this;
      let _1495_v0;
      let _nw212 = Array((new BigNumber(22)).toNumber()).fill(false);
      _1495_v0 = _nw212;
      let _1496_v1;
      _1496_v1 = false;
      let _index242 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
      (_1495_v0)[_index242] = _1496_v1;
      let _1497_v2;
      let _nw213 = new _module.C0();
      _nw213.__ctor();
      _1497_v2 = _nw213;
      let _1498_v3;
      _1498_v3 = _dafny.MultiSet.fromElements(_1497_v2, _1497_v2);
      let _1499_v4;
      _1499_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1496_v1,_dafny.MultiSet.fromElements(_1497_v2));
      let _index243 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
      (_1495_v0)[_index243] = !(_1498_v3).equals((((_1499_v4).contains(_1496_v1)) ? ((_1499_v4).get(_1496_v1)) : (_1498_v3)));
      let _1500_v5;
      let _1501_v6;
      let _out47;
      let _out48;
      let _outcollector22 = _module.__default.m0((p0).minus(p0), _1495_v0, globalState);
      _out47 = _outcollector22[0];
      _out48 = _outcollector22[1];
      _1500_v5 = _out47;
      _1501_v6 = _out48;
      let _1502_v7;
      _1502_v7 = _module.D1.create_DC3(p0);
      let _source30 = _1502_v7;
      if (_source30.is_DC4) {
        let _1503___mcc_h0 = (_source30).cf5;
        let _1504___mcc_h1 = (_source30).cf6;
        let _1505_cf6 = _1504___mcc_h1;
        let _1506_cf5 = _1503___mcc_h0;
        let _1507_v8;
        _1507_v8 = _dafny.Seq.of(_1495_v0);
        let _1508_v9;
        _1508_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1506_cf5,(_1507_v8)[_module.__default.safeIndex(p0, new BigNumber((_1507_v8).length))]);
        _1508_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1506_cf5,_1495_v0);
        let _1509_v10;
        _1509_v10 = _module.D5.create_DC16(_module.__default.fm41(globalState));
        _1509_v10 = _1509_v10;
        let _1510_v11;
        _1510_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))],(_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))]);
        let _1511_v12;
        _1511_v12 = _dafny.Set.fromElements(p0);
        let _1512_v13;
        _1512_v13 = _dafny.Seq.of(_1511_v12);
        let _1513_v14;
        _1513_v14 = _dafny.Set.fromElements(new BigNumber((_1512_v13).length));
        let _1514_v15;
        _1514_v15 = _dafny.MultiSet.fromElements(_1511_v12);
        let _1515_v16;
        _1515_v16 = _dafny.Seq.of(new BigNumber((_1510_v11).length), p0, p0, new BigNumber((_1514_v15).cardinality()));
        (globalState).f7 = ((_1515_v16)[_module.__default.safeIndex(p0, new BigNumber((_1515_v16).length))]).multipliedBy(_1506_cf5);
        (globalState).f7 = new BigNumber((((false) ? (function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of _dafny.IntegerRange(new BigNumber(334), new BigNumber(780))) {
            let _1516_v17 = _compr_28;
            if (((new BigNumber(334)).isLessThanOrEqualTo(_1516_v17)) && ((_1516_v17).isLessThan(new BigNumber(780)))) {
              _coll28.add((_1516_v17).multipliedBy(new BigNumber(84)));
            }
          }
          return _coll28;
        }()) : (_dafny.Set.fromElements(p0)))).length);
      } else if (_source30.is_DC3) {
        let _1517___mcc_h2 = (_source30).cf4;
        let _1518_cf4 = _1517___mcc_h2;
        let _1519_v18;
        let _nw214 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1519_v18 = _nw214;
        let _1520_v19;
        let _nw215 = new _module.C4();
        _nw215.__ctor();
        _1520_v19 = _nw215;
        let _index244 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_1519_v18).length));
        (_1519_v18)[_index244] = _dafny.MultiSet.fromElements(_1520_v19);
        let _1521_v20;
        _1521_v20 = _dafny.MultiSet.fromElements(_1520_v19, _1520_v19, ((_1496_v1) ? (_1520_v19) : (_1520_v19)));
        let _index245 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_1519_v18).length));
        let _rhs240 = new BigNumber((_1500_v5).length);
        let _rhs241 = _1521_v20;
        let _lhs174 = _1519_v18;
        let _lhs175 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_1519_v18).length));
        _1518_cf4 = _rhs240;
        _lhs174[_lhs175] = _rhs241;
        let _index246 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
        (_1495_v0)[_index246] = !(_1496_v1);
        let _1522_v21;
        _1522_v21 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1523_v22;
        _1523_v22 = _module.D14.create_DC43(_1522_v21, p0, p0, _1518_cf4);
        let _1524_v23;
        _1524_v23 = _dafny.Set.fromElements(_module.D14.create_DC43(_1522_v21, _1518_cf4, new BigNumber((_1501_v6).length), _1518_cf4), _1523_v22);
        let _1525_v24;
        _1525_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))],_1496_v1);
        let _1526_v25;
        _1526_v25 = _module.D3.create_DC11(_1518_cf4, _1525_v24);
        let _1527_v26;
        let _nw216 = Array((new BigNumber(4)).toNumber());
        _nw216[(_dafny.ZERO).toNumber()] = (_1497_v2).fm36(globalState);
        _nw216[(_dafny.ONE).toNumber()] = new BigNumber((_1524_v23).length);
        _nw216[(new BigNumber(2)).toNumber()] = (_1526_v25).dtor_cf17;
        _nw216[(new BigNumber(3)).toNumber()] = _1518_cf4;
        _1527_v26 = _nw216;
        let _1528_v27;
        _1528_v27 = _dafny.Seq.of(new BigNumber((_1501_v6).length));
        let _1529_v28;
        _1529_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1522_v21,_1518_cf4);
        let _1530_v29;
        _1530_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1518_cf4);
        let _nw217 = Array((new BigNumber(13)).toNumber());
        _nw217[(_dafny.ZERO).toNumber()] = (_1518_cf4).minus(_1518_cf4);
        _nw217[(_dafny.ONE).toNumber()] = new BigNumber(320);
        _nw217[(new BigNumber(2)).toNumber()] = p0;
        _nw217[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_1518_cf4, _1518_cf4);
        _nw217[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1528_v27).length), p0);
        _nw217[(new BigNumber(5)).toNumber()] = p0;
        _nw217[(new BigNumber(6)).toNumber()] = _1518_cf4;
        _nw217[(new BigNumber(7)).toNumber()] = new BigNumber(937);
        _nw217[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1501_v6).length));
        _nw217[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber((_1529_v28).length));
        _nw217[(new BigNumber(10)).toNumber()] = _1518_cf4;
        _nw217[(new BigNumber(11)).toNumber()] = _1518_cf4;
        _nw217[(new BigNumber(12)).toNumber()] = new BigNumber((_1530_v29).length);
        _1527_v26 = _nw217;
        let _1531_v30;
        let _nw218 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
        _1531_v30 = _nw218;
        let _1532_v31;
        _1532_v31 = _dafny.Seq.of(_1496_v1);
        let _index247 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1531_v30).length));
        (_1531_v30)[_index247] = _dafny.Seq.Concat(_1532_v31, _dafny.Seq.of(_1496_v1));
        let _index248 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1531_v30).length));
        (_1531_v30)[_index248] = _dafny.Seq.Concat(_1532_v31, _dafny.Seq.Concat(_1532_v31, _dafny.Seq.of(_1496_v1)));
      } else {
        let _1533___mcc_h3 = (_source30).cf7;
        let _1534_cf7 = _1533___mcc_h3;
        if ((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))]) {
          (globalState).f1 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), function (_1535_i0) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length))).multipliedBy(_module.__default.safeModuloInt(p0, p0)));
          (globalState).f7 = new BigNumber(740);
          _1496_v1 = (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))];
          let _1536_v32;
          let _nw219 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _1536_v32 = _nw219;
          let _1537_v33;
          _1537_v33 = _dafny.Seq.of(_1496_v1, (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))]);
          let _index249 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1536_v32).length));
          (_1536_v32)[_index249] = new BigNumber((_1537_v33).length);
          let _1538_v34;
          _1538_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1539_v35;
          _1539_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1538_v34,new BigNumber(888));
          let _1540_v36;
          _1540_v36 = _dafny.Map.Empty.slice().updateUnsafe((((_1539_v35).contains(_1538_v34)) ? ((_1539_v35).get(_1538_v34)) : (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1501_v6).length))).cardinality()))),p0);
          let _1541_v37;
          _1541_v37 = _dafny.MultiSet.fromElements(true, _1496_v1);
          let _index250 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1536_v32).length));
          let _index251 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
          let _rhs242 = (_1497_v2).fm37(_dafny.Seq.Concat(_1501_v6, _1500_v5), (((_1540_v36).contains(p0)) ? ((_1540_v36).get(p0)) : (p0)), (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))], (_1541_v37).contains(_1496_v1), globalState);
          let _rhs243 = p0;
          let _rhs244 = _1500_v5;
          let _rhs245 = (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))];
          let _lhs176 = _1536_v32;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1536_v32).length));
          let _lhs178 = _1495_v0;
          let _lhs179 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
          _1496_v1 = _rhs242;
          _lhs176[_lhs177] = _rhs243;
          _1501_v6 = _rhs244;
          _lhs178[_lhs179] = _rhs245;
          let _1542_v38;
          _1542_v38 = new _dafny.CodePoint('q'.codePointAt(0));
          _1542_v38 = _1542_v38;
        } else {
          let _1543_v40;
          _1543_v40 = _dafny.Set.fromElements(function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of _dafny.IntegerRange(new BigNumber(-89), new BigNumber(883))) {
              let _1544_v39 = _compr_29;
              if (((new BigNumber(-89)).isLessThanOrEqualTo(_1544_v39)) && ((_1544_v39).isLessThan(new BigNumber(883)))) {
                _coll29.push([(_1544_v39).minus(p0),p0]);
              }
            }
            return _coll29;
          }(), _dafny.Map.Empty.slice().updateUnsafe(p0,p0));
          let _1545_v41;
          _1545_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1546_v43;
          _1546_v43 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), ((_1547_p0) => function (_1548_i1) {
            return _1547_p0;
          })(p0)));
          let _1549_v44;
          _1549_v44 = _dafny.Seq.of(new BigNumber((_1546_v43).length));
          _1543_v40 = _dafny.Set.fromElements((_1545_v41).Merge((function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of (_1549_v44).Elements) {
              let _1550_v42 = _compr_30;
              if (_dafny.Seq.contains(_1549_v44, _1550_v42)) {
                _coll30.push([(_1550_v42).plus(p0),new BigNumber(325)]);
              }
            }
            return _coll30;
          }()).update(p0, p0)), _1545_v41);
          let _1551_v45;
          let _init40 = ((_1552_p0) => function (_1553_i2) {
            return _module.__default.safeDivisionInt(_1553_i2, _1552_p0);
          })(p0);
          let _nw220 = Array((new BigNumber(9)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw220.length); _i0_40++) {
            _nw220[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1551_v45 = _nw220;
          let _1554_v46;
          let _nw221 = Array((new BigNumber(15)).toNumber());
          _nw221[(_dafny.ZERO).toNumber()] = _1551_v45;
          _nw221[(_dafny.ONE).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(2)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(3)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(4)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(5)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(6)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(7)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(8)).toNumber()] = ((!((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))])) ? (_1551_v45) : (_1551_v45));
          _nw221[(new BigNumber(9)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(10)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(11)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(12)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(13)).toNumber()] = _1551_v45;
          _nw221[(new BigNumber(14)).toNumber()] = _1551_v45;
          _1554_v46 = _nw221;
          let _index252 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1554_v46).length));
          (_1554_v46)[_index252] = _1551_v45;
          let _1555_v47;
          _1555_v47 = _dafny.MultiSet.fromElements(p0);
          let _1556_v48;
          _1556_v48 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0), p0);
          let _index253 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
          let _index254 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1554_v46).length));
          let _rhs246 = new BigNumber((((_1555_v47).Difference(_dafny.MultiSet.fromElements(p0))).Union(_1555_v47)).cardinality());
          let _rhs247 = (p0).isLessThan((_dafny.ZERO).minus(new BigNumber((_1556_v48).length)));
          let _rhs248 = p0;
          let _rhs249 = _1551_v45;
          let _lhs180 = globalState;
          let _lhs181 = _1495_v0;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
          let _lhs183 = globalState;
          let _lhs184 = _1554_v46;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1554_v46).length));
          _lhs180.f7 = _rhs246;
          _lhs181[_lhs182] = _rhs247;
          _lhs183.f0 = _rhs248;
          _lhs184[_lhs185] = _rhs249;
          let _index255 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
          (_1495_v0)[_index255] = (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))];
          _1496_v1 = _1496_v1;
          let _index256 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1551_v45).length));
          (_1551_v45)[_index256] = _module.__default.safeModuloInt(p0, p0);
          let _index257 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1551_v45).length));
          (_1551_v45)[_index257] = p0;
        }
        _1500_v5 = _dafny.Seq.UnicodeFromString("vhffhwbq");
        let _1557_v49;
        _1557_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(838),_dafny.Seq.of(p0));
        let _1558_v50;
        _1558_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1557_v49).length),(_1497_v2).fm36(globalState));
        let _1559_v51;
        _1559_v51 = _dafny.MultiSet.fromElements(p0);
        let _1560_v52;
        _1560_v52 = _dafny.Map.Empty.slice().updateUnsafe((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))],new BigNumber((_1559_v51).cardinality()));
        let _1561_v53;
        _1561_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1560_v52,new BigNumber(647));
        let _1562_v54;
        _1562_v54 = _dafny.Map.Empty.slice().updateUnsafe((((_1558_v50).contains(p0)) ? ((_1558_v50).get(p0)) : (p0)),new BigNumber((_1561_v53).length));
        let _1563_v55;
        let _nw222 = Array((new BigNumber(14)).toNumber());
        _1563_v55 = _nw222;
        let _1564_v56;
        _1564_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1496_v1,_1563_v55);
        let _1565_v57;
        _1565_v57 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1566_v58;
        _1566_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1496_v1,_1565_v57);
        (globalState).f0 = (new BigNumber(((_1562_v54).update(new BigNumber((_1564_v56).length), (_dafny.ZERO).minus(p0))).length)).minus(new BigNumber((_1566_v58).length));
        let _1567_v59;
        let _nw223 = new _module.C0();
        _nw223.__ctor();
        _1567_v59 = _nw223;
      }
      let _1568_v60;
      _1568_v60 = _dafny.Seq.of(p0, new BigNumber(-978));
      let _1569_v61;
      _1569_v61 = _dafny.Seq.of((_1568_v60)[_module.__default.safeIndex(new BigNumber(81), new BigNumber((_1568_v60).length))]);
      let _1570_v62;
      _1570_v62 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0),_1501_v6);
      let _1571_i3;
      _1571_i3 = _dafny.ZERO;
      L8: {
        while (!(((_1570_v62).update(_dafny.Seq.of(p0, p0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(477)), function (_1591_i4) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }))).Merge(_1570_v62)).contains(_1569_v61)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1571_i3)) {
              break L8;
            }
            _1571_i3 = (_1571_i3).plus(_dafny.ONE);
            let _1572_v63;
            _1572_v63 = new _dafny.CodePoint('y'.codePointAt(0));
            let _1573_v64;
            _1573_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))]);
            let _1574_v65;
            _1574_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()),new BigNumber((_1573_v64).length));
            let _1575_v66;
            let _nw224 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1575_v66 = _nw224;
            let _1576_v67;
            _1576_v67 = _module.D12.create_DC37(p0, _1575_v66, new _dafny.CodePoint('j'.codePointAt(0)));
            let _1577_v68;
            _1577_v68 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_1501_v6, _1500_v5), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1501_v6, _1500_v5)).length)), _1572_v63),((_dafny.ZERO).minus((((_1574_v65).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_1580_v63) => function (_1581_i5) {
              return _1580_v63;
            })(_1572_v63))).length))) ? ((_1574_v65).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_1578_v63) => function (_1579_i5) {
              return _1578_v63;
            })(_1572_v63))).length))) : ((_1576_v67).dtor_cf61)))).multipliedBy(p0));
            let _1582_v70;
            _1582_v70 = _dafny.Seq.of(_1500_v5);
            _1577_v68 = (function () {
              let _coll31 = new _dafny.Map();
              for (const _compr_31 of (_1582_v70).Elements) {
                let _1583_v69 = _compr_31;
                if (_dafny.Seq.contains(_1582_v70, _1583_v69)) {
                  _coll31.push([_1583_v69,(_dafny.ZERO).minus(p0)]);
                }
              }
              return _coll31;
            }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("budcfif"),p0));
            (globalState).f1 = _module.__default.safeModuloInt(p0, p0);
            let _1584_v71;
            _1584_v71 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, globalState),_1572_v63);
            let _1585_v72;
            _1585_v72 = _dafny.Set.fromElements(_1496_v1, _1496_v1, (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))]);
            let _1586_v73;
            _1586_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))],_1585_v72);
            _1584_v71 = (_1584_v71).update((_1585_v72).IsSubsetOf((((_1586_v73).contains((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))])) ? ((_1586_v73).get((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))])) : (_1585_v72))), _1572_v63);
            let _1587_v74;
            _1587_v74 = _module.D10.create_DC31((_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))], false, new BigNumber(930));
            let _1588_v75;
            _1588_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1587_v74,(_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))]);
            let _1589_v76;
            _1589_v76 = _dafny.MultiSet.fromElements(_1588_v75, _1588_v75, _1588_v75, _1588_v75);
            let _1590_v77;
            _1590_v77 = _module.D15.create_DC44(_1589_v76);
            _1496_v1 = ((_1590_v77).dtor_cf79).contains(_1588_v75);
          }
        }
      }
      let _index258 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length));
      (_1495_v0)[_index258] = (_1495_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1495_v0).length))];
      let _1592_v78;
      let _nw225 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _1592_v78 = _nw225;
      let _index259 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1592_v78).length));
      (_1592_v78)[_index259] = p0;
      let _index260 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1592_v78).length));
      (_1592_v78)[_index260] = (p0).multipliedBy((new BigNumber((_1501_v6).length)).plus(p0));
      return;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f10 = false;
      this._f19 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f19, f10) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f10 = f10;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((new BigNumber(-377)).isLessThan(new BigNumber(-864))) || ((_this).f10);
    };
    fm7(p0, globalState) {
      let _this = this;
      return (_this).f19;
    };
    fm20(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_this).f19)).Intersect((_module.D5.create_DC16(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f19)))).dtor_cf27);
    };
    fm21(p0, globalState) {
      let _this = this;
      return (_this).f19;
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Map.Empty;
      let _1593_v0;
      _1593_v0 = new BigNumber(733);
      (globalState).f0 = _1593_v0;
      let _1594_v1;
      let _nw226 = new _module.C5();
      _nw226.__ctor();
      _1594_v1 = _nw226;
      if (p1) {
        let _1595_v2;
        _1595_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10);
        let _1596_v3;
        _1596_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1595_v2).update((_this).f19, p1)).length),_dafny.Seq.of(false, (_this).f10));
        (globalState).f0 = (_dafny.ZERO).minus(((p1) ? (_1593_v0) : ((_1593_v0).minus(new BigNumber((_1596_v3).length)))));
        let _1597_v4;
        let _nw227 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1597_v4 = _nw227;
        let _1598_v5;
        _1598_v5 = _dafny.Seq.of(_1597_v4, _1597_v4, _1597_v4);
        let _1599_v6;
        _1599_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1593_v0,(_this).f19);
        let _1600_v7;
        _1600_v7 = _dafny.Seq.of(_1593_v0, new BigNumber((_1599_v6).length), _1593_v0);
        let _1601_v8;
        _1601_v8 = _dafny.Seq.of(_1600_v7);
        _1597_v4 = (_1598_v5)[_module.__default.safeIndex(new BigNumber(((_1601_v8)[_module.__default.safeIndex((_1600_v7)[_module.__default.safeIndex(_1593_v0, new BigNumber((_1600_v7).length))], new BigNumber((_1601_v8).length))]).length), new BigNumber((_1598_v5).length))];
        (globalState).f1 = _module.__default.safeModuloInt(_1593_v0, _1593_v0);
        let _1602_v10;
        _1602_v10 = _dafny.Set.fromElements(new BigNumber((function () {
          let _coll32 = new _dafny.Set();
          for (const _compr_32 of _dafny.IntegerRange(new BigNumber(930), new BigNumber(386))) {
            let _1603_v9 = _compr_32;
            if (((new BigNumber(930)).isLessThanOrEqualTo(_1603_v9)) && ((_1603_v9).isLessThan(new BigNumber(386)))) {
              _coll32.add((_1603_v9).plus(_1593_v0));
            }
          }
          return _coll32;
        }()).length));
        let _1604_v11;
        _1604_v11 = _dafny.Set.fromElements(new BigNumber((_1602_v10).length));
        let _rhs250 = ((_1604_v11).Difference(_1604_v11)).IsSubsetOf((_1602_v10).Union(_1604_v11));
        let _rhs251 = _1593_v0;
        let _lhs186 = globalState;
        r0 = _rhs250;
        _lhs186.f1 = _rhs251;
        let _1605_v12;
        _1605_v12 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1606_v13;
        _1606_v13 = _dafny.Seq.of((_this).f10, !(p1), p1, p1);
        let _1607_v14;
        _1607_v14 = _dafny.Seq.UnicodeFromString("tmux");
        _1605_v12 = _module.__default.fm4((_1606_v13)[_module.__default.safeIndex(new BigNumber((_1607_v14).length), new BigNumber((_1606_v13).length))], (_this).f10, (_this).fm7(false, globalState), globalState);
      } else {
        (globalState).f1 = _1593_v0;
        let _1608_v15;
        _1608_v15 = _dafny.Seq.UnicodeFromString("lsp");
        let _1609_v16;
        _1609_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1593_v0).isLessThan(_1593_v0),_1608_v15);
        let _1610_v17;
        _1610_v17 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1611_v18;
        let _nw228 = Array((new BigNumber(2)).toNumber());
        _nw228[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1608_v15, _module.__default.safeIndex(_1593_v0, new BigNumber((_1608_v15).length)), _1610_v17);
        _nw228[(_dafny.ONE).toNumber()] = _1608_v15;
        _1611_v18 = _nw228;
        let _index261 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_1611_v18).length));
        (_1611_v18)[_index261] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), ((_1612_v17) => function (_1613_i0) {
          return _1612_v17;
        })(_1610_v17));
        let _1614_v19;
        let _nw229 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _1614_v19 = _nw229;
        let _1615_v20;
        _1615_v20 = _dafny.Seq.of(p1);
        let _1616_v21;
        _1616_v21 = _dafny.Set.fromElements(_1615_v20);
        let _1617_v22;
        _1617_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1616_v21).length),_module.__default.fm51(globalState));
        let _1618_v23;
        _1618_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1615_v20).length),_1609_v16);
        let _index262 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_1611_v18).length));
        let _rhs252 = (_1609_v16).Merge((((_1618_v23).contains(new BigNumber(-866))) ? ((_1618_v23).get(new BigNumber(-866))) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1608_v15))));
        let _rhs253 = _1608_v15;
        let _rhs254 = _1614_v19;
        let _rhs255 = _1617_v22;
        let _rhs256 = p1;
        let _lhs187 = _1611_v18;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_1611_v18).length));
        _1609_v16 = _rhs252;
        _lhs187[_lhs188] = _rhs253;
        _1614_v19 = _rhs254;
        _1617_v22 = _rhs255;
        r0 = _rhs256;
        let _1619_v24;
        _1619_v24 = _module.D5.create_DC17(!(false), _1610_v17);
        let _1620_v25;
        _1620_v25 = _module.D5.create_DC18(_1619_v24);
        let _1621_v26;
        _1621_v26 = _module.D5.create_DC18(_1620_v25);
        let _source31 = _1621_v26;
        if (_source31.is_DC17) {
          let _1622___mcc_h0 = (_source31).cf28;
          let _1623___mcc_h1 = (_source31).cf29;
          let _1624_cf29 = _1623___mcc_h1;
          let _1625_cf28 = _1622___mcc_h0;
          let _1626_v27;
          _1626_v27 = _dafny.MultiSet.fromElements(true, _1625_cf28, _1625_cf28);
          let _1627_v28;
          _1627_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1593_v0,(((_1626_v27).contains((_this).f19)) ? ((_1626_v27).get((_this).f19)) : ((_dafny.ZERO).minus(_1593_v0))));
          let _1628_v29;
          let _nw230 = new _module.C2();
          _nw230.__ctor(_1625_cf28, new BigNumber(((_1627_v28).Merge(_1627_v28)).length));
          _1628_v29 = _nw230;
          let _1629_v30;
          let _nw231 = new _module.C3();
          _nw231.__ctor(_1625_cf28, _1625_cf28);
          _1629_v30 = _nw231;
          (globalState).f0 = (_1628_v29).f22;
          let _1630_v31;
          _1630_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1625_cf28,(_this).f19);
          let _1631_v32;
          _1631_v32 = _module.D3.create_DC11(new BigNumber(788), _1630_v31);
          let _1632_v33;
          _1632_v33 = _dafny.Set.fromElements(p1);
          r0 = ((!(_1630_v31).equals(((_1631_v32).dtor_cf18).update((_this).f10, (_this).f10))) ? ((_1628_v29).f21) : ((_1632_v33).IsSubsetOf(_1632_v33)));
        } else if (_source31.is_DC16) {
          let _1633___mcc_h2 = (_source31).cf27;
          let _1634_cf27 = _1633___mcc_h2;
          r0 = (p1) === ((_this).f19);
          (globalState).f0 = (_1593_v0).minus((_1593_v0).multipliedBy(_1593_v0));
          let _1635_v34;
          _1635_v34 = _module.D11.create_DC34(_1593_v0, (_this).fm21(_1593_v0, globalState), _1593_v0);
          (globalState).f1 = ((_1635_v34).dtor_cf56).plus(_1593_v0);
          let _1636_v35;
          let _nw232 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1636_v35 = _nw232;
          let _1637_v36;
          _1637_v36 = _dafny.Set.fromElements((_this).f19);
          let _1638_v37;
          _1638_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1636_v35,(_1637_v36).IsProperSubsetOf(_1637_v36));
          _1638_v37 = (_1638_v37).update(_1636_v35, (_this).f10);
        } else {
          let _1639___mcc_h3 = (_source31).cf30;
          let _1640_cf30 = _1639___mcc_h3;
          let _1641_v38;
          _1641_v38 = _dafny.Set.fromElements((_this).f10);
          let _1642_v39;
          _1642_v39 = _dafny.Set.fromElements(_1641_v38, _1641_v38);
          let _1643_v40;
          _1643_v40 = _module.D3.create_DC9((_1642_v39).Intersect(_1642_v39));
          _1643_v40 = ((p1) ? (_1643_v40) : (_1643_v40));
          let _1644_v41;
          let _nw233 = new _module.C0();
          _nw233.__ctor();
          _1644_v41 = _nw233;
          let _1645_v42;
          let _nw234 = new _module.C0();
          _nw234.__ctor();
          _1645_v42 = _nw234;
          let _1646_v43;
          let _nw235 = new _module.C3();
          _nw235.__ctor(p1, true);
          _1646_v43 = _nw235;
          let _1647_v44;
          let _1648_v45;
          let _out49;
          let _out50;
          let _outcollector23 = (_1594_v1).m13((((_this).f10) ? (_1646_v43) : (_1646_v43)), globalState);
          _out49 = _outcollector23[0];
          _out50 = _outcollector23[1];
          _1647_v44 = _out49;
          _1648_v45 = _out50;
        }
        let _rhs257 = !((_1593_v0).isLessThan((_dafny.ZERO).minus(_1593_v0)));
        let _rhs258 = _1593_v0;
        let _lhs189 = globalState;
        r0 = _rhs257;
        _lhs189.f0 = _rhs258;
        let _rhs259 = p1;
        let _rhs260 = _1593_v0;
        let _rhs261 = (_this).f10;
        let _rhs262 = _1610_v17;
        let _rhs263 = _1611_v18;
        let _lhs190 = globalState;
        r0 = _rhs259;
        _lhs190.f1 = _rhs260;
        r0 = _rhs261;
        _1610_v17 = _rhs262;
        _1611_v18 = _rhs263;
      }
      let _1649_v46;
      _1649_v46 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_1593_v0)).length));
      let _1650_v47;
      _1650_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1649_v46);
      _1650_v47 = (_1650_v47).update((_this).f10, _1649_v46);
      let _1651_v48;
      _1651_v48 = _module.D10.create_DC31((_this).f19, p1, _1593_v0);
      let _1652_v49;
      _1652_v49 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1653_v50;
      _1653_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1593_v0);
      let _1654_v51;
      _1654_v51 = _dafny.Seq.of((_this).f19, p1);
      let _1655_v52;
      _1655_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1652_v49,!((_this).f19));
      let _1656_v53;
      let _nw236 = Array((new BigNumber(29)).toNumber());
      _nw236[(_dafny.ZERO).toNumber()] = (_1651_v48).dtor_cf50;
      _nw236[(_dafny.ONE).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(_1652_v49,(_this).f10)).contains(_1652_v49);
      _nw236[(new BigNumber(2)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(3)).toNumber()] = (_1653_v50).contains((_this).fm6(_1593_v0, _1593_v0, _1593_v0, (_this).f19, globalState));
      _nw236[(new BigNumber(4)).toNumber()] = p1;
      _nw236[(new BigNumber(5)).toNumber()] = (_this).f10;
      _nw236[(new BigNumber(6)).toNumber()] = (_this).f10;
      _nw236[(new BigNumber(7)).toNumber()] = (_1593_v0).isLessThan(_1593_v0);
      _nw236[(new BigNumber(8)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(9)).toNumber()] = (_this).f10;
      _nw236[(new BigNumber(10)).toNumber()] = (_1593_v0).isLessThanOrEqualTo(_1593_v0);
      _nw236[(new BigNumber(11)).toNumber()] = !(true);
      _nw236[(new BigNumber(12)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(13)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(14)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(15)).toNumber()] = false;
      _nw236[(new BigNumber(16)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(17)).toNumber()] = p1;
      _nw236[(new BigNumber(18)).toNumber()] = (_1653_v50).contains(p1);
      _nw236[(new BigNumber(19)).toNumber()] = (_this).f10;
      _nw236[(new BigNumber(20)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(21)).toNumber()] = (_this).f10;
      _nw236[(new BigNumber(22)).toNumber()] = p1;
      _nw236[(new BigNumber(23)).toNumber()] = !_dafny.areEqual(_dafny.Seq.of(true), _1654_v51);
      _nw236[(new BigNumber(24)).toNumber()] = (((_1655_v52).contains(_1652_v49)) ? ((_1655_v52).get(_1652_v49)) : ((_this).f19));
      _nw236[(new BigNumber(25)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(26)).toNumber()] = !(p1);
      _nw236[(new BigNumber(27)).toNumber()] = (_this).f19;
      _nw236[(new BigNumber(28)).toNumber()] = (_this).f19;
      _1656_v53 = _nw236;
      let _index263 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1656_v53).length));
      (_1656_v53)[_index263] = (_this).f10;
      let _index264 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1656_v53).length));
      let _rhs264 = (_this).f10;
      let _rhs265 = (_1593_v0).plus(new BigNumber(-186));
      let _rhs266 = ((_this).f10) === (false);
      let _rhs267 = !((_this).f19);
      let _rhs268 = _1593_v0;
      let _lhs191 = globalState;
      let _lhs192 = _1656_v53;
      let _lhs193 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1656_v53).length));
      let _lhs194 = globalState;
      r0 = _rhs264;
      _lhs191.f0 = _rhs265;
      _lhs192[_lhs193] = _rhs266;
      r0 = _rhs267;
      _lhs194.f1 = _rhs268;
      let _1657_v54;
      _1657_v54 = _dafny.MultiSet.fromElements(_1652_v49, _1652_v49, _1652_v49, _1652_v49, _1652_v49);
      let _1658_v55;
      _1658_v55 = _dafny.Seq.UnicodeFromString("ebbclmx");
      let _1659_v56;
      _1659_v56 = _dafny.Set.fromElements(_1658_v55);
      let _1660_v57;
      _1660_v57 = _dafny.Seq.of(_1654_v51, _dafny.Seq.of((_this).f10, (_this).f19));
      let _1661_v58;
      _1661_v58 = _dafny.MultiSet.fromElements(new BigNumber(592));
      let _1662_v59;
      _1662_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1660_v57)[_module.__default.safeIndex((((_1661_v58).contains(new BigNumber(-847))) ? ((_1661_v58).get(new BigNumber(-847))) : (_1593_v0)), new BigNumber((_1660_v57).length))]).length),new BigNumber((_1653_v50).length));
      let _1663_v60;
      _1663_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1649_v46).length),true);
      let _1664_v61;
      let _nw237 = Array((new BigNumber(22)).toNumber());
      _nw237[(_dafny.ZERO).toNumber()] = new BigNumber(190);
      _nw237[(_dafny.ONE).toNumber()] = new BigNumber(((_1657_v54).Intersect(_1657_v54)).cardinality());
      _nw237[(new BigNumber(2)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(3)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(4)).toNumber()] = new BigNumber((_1659_v56).length);
      _nw237[(new BigNumber(5)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_1593_v0);
      _nw237[(new BigNumber(7)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(8)).toNumber()] = ((((_1662_v59).contains(_1593_v0)) ? ((_1662_v59).get(_1593_v0)) : (new BigNumber(-693)))).multipliedBy(new BigNumber(684));
      _nw237[(new BigNumber(9)).toNumber()] = (new BigNumber((_dafny.Set.fromElements((_1656_v53)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_1656_v53).length))])).length)).minus(new BigNumber((_1658_v55).length));
      _nw237[(new BigNumber(10)).toNumber()] = new BigNumber((_1663_v60).length);
      _nw237[(new BigNumber(11)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(12)).toNumber()] = new BigNumber((_1658_v55).length);
      _nw237[(new BigNumber(13)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(14)).toNumber()] = _module.__default.fm1(globalState);
      _nw237[(new BigNumber(15)).toNumber()] = (new BigNumber(-342)).minus(new BigNumber((_1658_v55).length));
      _nw237[(new BigNumber(16)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(17)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(18)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(19)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(20)).toNumber()] = _1593_v0;
      _nw237[(new BigNumber(21)).toNumber()] = _1593_v0;
      _1664_v61 = _nw237;
      let _index265 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_1664_v61).length));
      (_1664_v61)[_index265] = _module.__default.fm1(globalState);
      let _index266 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_1664_v61).length));
      (_1664_v61)[_index266] = _module.__default.fm1(globalState);
      r0 = (_this).f10;
      r1 = _1658_v55;
      let _1665_v62;
      _1665_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1654_v51).length),_1662_v59);
      r2 = (_1662_v59).Merge((((_1665_v62).contains(new BigNumber(852))) ? ((_1665_v62).get(new BigNumber(852))) : (_1662_v59)));
      return [r0, r1, r2];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this.f18 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f18) {
      let _this = this;
      (_this).f18 = f18;
      return;
    }
    fm16(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.D4.create_DC13(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),false)) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),false))));
    };
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(new BigNumber(-426))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(-269)))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(207), new BigNumber(-989), new BigNumber((_dafny.Seq.of(false, false)).length)));
    };
    m10(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _1666_v0;
      _1666_v0 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_this.f18)).length));
      let _1667_v1;
      _1667_v1 = true;
      let _1668_v2;
      _1668_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1667_v1,p0);
      r0 = ((_1666_v0).Intersect(_1666_v0)).IsDisjointFrom(_module.__default.fm18((_dafny.ZERO).minus(p0), _1668_v2, !(!(_1667_v1)), globalState));
      let _1669_v3;
      _1669_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_1667_v1) ? (p0) : (p0)));
      _1669_v3 = (_1669_v3).update((p0).plus(p0), (p0).minus(new BigNumber((_module.__default.fm19(_1667_v1, _1667_v1, _1667_v1, p0, globalState)).cardinality())));
      let _1670_v4;
      let _nw238 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _1670_v4 = _nw238;
      let _index267 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length));
      (_1670_v4)[_index267] = (_dafny.ZERO).minus(p0);
      let _index268 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length));
      (_1670_v4)[_index268] = (_module.D1.create_DC3(_module.__default.fm1(globalState))).dtor_cf4;
      let _1671_v5;
      let _nw239 = new _module.C6();
      _nw239.__ctor();
      _1671_v5 = _nw239;
      (globalState).f7 = _module.__default.safeDivisionInt(((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]).multipliedBy((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]), (_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]);
      let _1672_i0;
      _1672_i0 = _dafny.ZERO;
      L9: {
        while (_1667_v1) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1672_i0)) {
              break L9;
            }
            _1672_i0 = (_1672_i0).plus(_dafny.ONE);
            _1667_v1 = (_dafny.MultiSet.fromElements(new BigNumber(316), (_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))])).IsDisjointFrom(_1666_v0);
            let _1673_v6;
            _1673_v6 = _module.D13.create_DC39(new BigNumber(247), _1667_v1);
            _1667_v1 = (_1673_v6).dtor_cf66;
            let _1674_v7;
            _1674_v7 = _dafny.Seq.UnicodeFromString("elrwfcc");
            let _1675_v9;
            _1675_v9 = _dafny.Seq.of(false, _1667_v1);
            let _1676_v10;
            _1676_v10 = _dafny.Seq.of(_1668_v2);
            let _1677_v11;
            _1677_v11 = _module.D11.create_DC34(new BigNumber((_1675_v9).length), _1667_v1, new BigNumber((_dafny.Seq.update(_1676_v10, _module.__default.safeIndex(new BigNumber((_1666_v0).cardinality()), new BigNumber((_1676_v10).length)), _module.__default.fm51(globalState))).length));
            let _1678_v12;
            _1678_v12 = _module.D4.create_DC13((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_1677_v11).dtor_cf57, _1667_v1),_1667_v1)).update(_1675_v9, _1667_v1));
            let _1679_v13;
            _1679_v13 = _dafny.Seq.of(_1678_v12);
            let _1680_v14;
            _1680_v14 = _module.D0.create_DC1(_1674_v7);
            let _1681_v15;
            _1681_v15 = _dafny.Seq.of(_1680_v14, _1680_v14, _1680_v14);
            if ((_this).fm17(_dafny.Seq.IsPrefixOf(_1674_v7, _dafny.Seq.UnicodeFromString("js")), new BigNumber(448), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(75)), ((_1682_v1) => function (_1683_i1) {
              return _module.D4.create_DC13(function () {
  let _coll33 = new _dafny.Map();
  for (const _compr_33 of (_dafny.Seq.of(_dafny.Seq.of(_1682_v1))).Elements) {
    let _1684_v8 = _compr_33;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(_1682_v1)), _1684_v8)) {
      _coll33.push([_1684_v8,_1682_v1]);
    }
  }
  return _coll33;
}());
            })(_1667_v1)), _1679_v13), _dafny.Seq.contains(_1681_v15, _1680_v14), globalState)) {
              let _1685_v16;
              let _init41 = ((_1686_v1) => function (_1687_i2) {
                return !(_dafny.Set.fromElements(_1686_v1, _1686_v1, _1686_v1, _1686_v1)).equals(_dafny.Set.fromElements(true));
              })(_1667_v1);
              let _nw240 = Array((new BigNumber(15)).toNumber());
              for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw240.length); _i0_41++) {
                _nw240[_i0_41] = _init41(new BigNumber(_i0_41));
              }
              _1685_v16 = _nw240;
              _1685_v16 = p1;
              let _1688_v17;
              _1688_v17 = _dafny.Seq.of((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]);
              let _1689_v18;
              let _nw241 = new _module.C4();
              _nw241.__ctor();
              _1689_v18 = _nw241;
              let _rhs269 = _dafny.Seq.of((p0).minus((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]), (_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))], (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1689_v18,_1667_v1)).length)).plus(p0));
              let _rhs270 = _dafny.Seq.contains(_1674_v7, new _dafny.CodePoint('a'.codePointAt(0)));
              _1688_v17 = _rhs269;
              r0 = _rhs270;
              let _1690_v19;
              _1690_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1671_v5).fm23(false, new BigNumber((_1666_v0).cardinality()), _1667_v1, new BigNumber(406), globalState),_1667_v1);
              let _1691_v20;
              _1691_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1690_v19).length),_1667_v1);
              let _index269 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_1685_v16).length));
              (_1685_v16)[_index269] = (((_1691_v20).contains(new BigNumber((_1666_v0).cardinality()))) ? ((_1691_v20).get(new BigNumber((_1666_v0).cardinality()))) : (_module.__default.fm0((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))], globalState)));
              let _index270 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_1685_v16).length));
              (_1685_v16)[_index270] = _1667_v1;
              (globalState).f1 = new BigNumber(370);
              (globalState).f7 = ((((_1666_v0).contains((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))])) ? ((_1666_v0).get((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))])) : ((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]))).plus(((_1667_v1) ? (p0) : ((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))])));
            } else {
              let _1692_v21;
              _1692_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, globalState),_1667_v1);
              let _1693_v22;
              _1693_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.D3.create_DC11((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))], _1692_v21));
              let _1694_v23;
              _1694_v23 = _module.D3.create_DC10(_1674_v7, true, _1667_v1);
              let _1695_v24;
              _1695_v24 = _module.D3.create_DC12((((_1693_v22).contains(p1)) ? ((_1693_v22).get(p1)) : (_1694_v23)));
              _1695_v24 = _1695_v24;
              let _index271 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length));
              (_1670_v4)[_index271] = (((true) ? (new BigNumber(294)) : (new BigNumber((_1675_v9).length)))).plus((_1670_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1670_v4).length))]);
              let _1696_v26;
              let _init42 = ((_1697_p0, _1698_v1, _1699_v4) => function (_1700_i3) {
                return function () {
                  let _coll34 = new _dafny.Map();
                  for (const _compr_34 of _dafny.IntegerRange(new BigNumber(195), new BigNumber(-108))) {
                    let _1701_v25 = _compr_34;
                    if (((new BigNumber(195)).isLessThanOrEqualTo(_1701_v25)) && ((_1701_v25).isLessThan(new BigNumber(-108)))) {
                      _coll34.push([(_1701_v25).multipliedBy(_1697_p0),_module.D11.create_DC33(_1698_v1, (_1699_v4)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1699_v4).length))])]);
                    }
                  }
                  return _coll34;
                }();
              })(p0, _1667_v1, _1670_v4);
              let _nw242 = Array((new BigNumber(8)).toNumber());
              for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw242.length); _i0_42++) {
                _nw242[_i0_42] = _init42(new BigNumber(_i0_42));
              }
              _1696_v26 = _nw242;
              let _1702_v27;
              _1702_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1696_v26,_1667_v1);
              _1702_v27 = (_1702_v27).update(_1696_v26, true);
              let _1703_v28;
              let _nw243 = new _module.C4();
              _nw243.__ctor();
              _1703_v28 = _nw243;
              let _rhs271 = _1667_v1;
              let _rhs272 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qe"), _1674_v7), _1674_v7);
              let _rhs273 = p0;
              let _rhs274 = _1667_v1;
              let _lhs195 = globalState;
              _1667_v1 = _rhs271;
              _1674_v7 = _rhs272;
              _lhs195.f1 = _rhs273;
              _1667_v1 = _rhs274;
            }
            (globalState).f0 = (_dafny.ZERO).minus(_module.__default.fm1(globalState));
          }
        }
      }
      r0 = !((!(_1667_v1)) || (_1667_v1));
      return r0;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f10 = false;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f17, f10) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f10 = f10;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cepp"), _dafny.Seq.UnicodeFromString("onfrv")), _dafny.Seq.UnicodeFromString("d"));
    };
    fm7(p0, globalState) {
      let _this = this;
      let _source32 = _module.D2.create_DC6(new _dafny.CodePoint('b'.codePointAt(0)));
      if (_source32.is_DC7) {
        let _1704___mcc_h0 = (_source32).cf9;
        let _1705___mcc_h1 = (_source32).cf10;
        let _1706___mcc_h2 = (_source32).cf11;
        let _1707_cf11 = _1706___mcc_h2;
        let _1708_cf10 = _1705___mcc_h1;
        let _1709_cf9 = _1704___mcc_h0;
        return _1708_cf10;
      } else if (_source32.is_DC6) {
        let _1710___mcc_h3 = (_source32).cf8;
        let _1711_cf8 = _1710___mcc_h3;
        return (_this).f10;
      } else {
        let _1712___mcc_h4 = (_source32).cf12;
        let _1713_cf12 = _1712___mcc_h4;
        return false;
      }
    };
    fm14(p0, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jn"))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lbvekb")));
    };
    fm15(globalState) {
      let _this = this;
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_this).f10) ? (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f17, (_this).f10, (_this).f17))).cardinality())))) : (_dafny.Set.fromElements(new BigNumber(529)))),(_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber(774))).length)).multipliedBy(new BigNumber(660))))).length);
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Map.Empty;
      let _1714_v0;
      let _nw244 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
      _1714_v0 = _nw244;
      let _1715_v1;
      _1715_v1 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1716_v2;
      _1716_v2 = new BigNumber(260);
      let _index272 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_1714_v0).length));
      (_1714_v0)[_index272] = _dafny.Map.Empty.slice().updateUnsafe(_1715_v1,(_dafny.ZERO).minus(_1716_v2));
      let _1717_v3;
      _1717_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1715_v1,_1716_v2);
      let _index273 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_1714_v0).length));
      (_1714_v0)[_index273] = (_1717_v3).Merge(_1717_v3);
      let _1718_v4;
      let _init43 = function (_1719_i0) {
        return (_1719_i0).plus(new BigNumber(-728));
      };
      let _nw245 = Array((new BigNumber(11)).toNumber());
      for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw245.length); _i0_43++) {
        _nw245[_i0_43] = _init43(new BigNumber(_i0_43));
      }
      _1718_v4 = _nw245;
      let _1720_v5;
      _1720_v5 = _dafny.Set.fromElements(_1718_v4, _1718_v4, _1718_v4);
      if ((_1720_v5).IsDisjointFrom(_1720_v5)) {
        let _1721_v6;
        let _nw246 = Array((new BigNumber(12)).toNumber()).fill(false);
        _1721_v6 = _nw246;
        _1721_v6 = _1721_v6;
        _1721_v6 = (((_this).fm7((_this).f10, globalState)) ? (_1721_v6) : (_1721_v6));
        let _1722_v7;
        let _1723_v8;
        let _out51;
        let _out52;
        let _outcollector24 = _module.__default.m0(new BigNumber((_dafny.Seq.of(_1715_v1, _1715_v1, _1715_v1)).length), _1721_v6, globalState);
        _out51 = _outcollector24[0];
        _out52 = _outcollector24[1];
        _1722_v7 = _out51;
        _1723_v8 = _out52;
        _1715_v1 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1724_v9;
        let _nw247 = new _module.C1();
        _nw247.__ctor((_this).f17);
        _1724_v9 = _nw247;
      } else {
        let _1725_v10;
        _1725_v10 = _dafny.Seq.UnicodeFromString("rbye");
        let _1726_v11;
        _1726_v11 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_1716_v2).minus((_this).fm15(globalState))), _1716_v2, _1716_v2, _1716_v2, new BigNumber((_1725_v10).length));
        let _1727_v12;
        _1727_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1726_v11,(_this).f10);
        (globalState).f1 = (((_1726_v11).contains(_1716_v2)) ? ((_1726_v11).get(_1716_v2)) : (new BigNumber((_1727_v12).length)));
        let _1728_v13;
        let _nw248 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1728_v13 = _nw248;
        _1728_v13 = _1728_v13;
        let _1729_v14;
        _1729_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1716_v2);
        r0 = !(_1729_v14).contains(((p1) ? ((_this).f10) : (p1)));
        let _1730_v15;
        let _nw249 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1730_v15 = _nw249;
        let _1731_v16;
        _1731_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1725_v10);
        let _1732_v17;
        _1732_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1731_v16).length),_1716_v2);
        let _1733_v18;
        _1733_v18 = _dafny.Seq.of(_1732_v17, _1732_v17);
        let _1734_v19;
        _1734_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1732_v17);
        let _1735_v20;
        let _nw250 = Array((new BigNumber(10)).toNumber());
        _nw250[(_dafny.ZERO).toNumber()] = (_1733_v18)[_module.__default.safeIndex(_1716_v2, new BigNumber((_1733_v18).length))];
        _nw250[(_dafny.ONE).toNumber()] = (_1732_v17).Merge(_1732_v17);
        _nw250[(new BigNumber(2)).toNumber()] = (((_1734_v19).contains((_this).f17)) ? ((_1734_v19).get((_this).f17)) : (_1732_v17));
        _nw250[(new BigNumber(3)).toNumber()] = _1732_v17;
        _nw250[(new BigNumber(4)).toNumber()] = _1732_v17;
        _nw250[(new BigNumber(5)).toNumber()] = _1732_v17;
        _nw250[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1716_v2,(_dafny.ZERO).minus(_1716_v2));
        _nw250[(new BigNumber(7)).toNumber()] = (_1732_v17).Merge(_1732_v17);
        _nw250[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1716_v2,_1716_v2);
        _nw250[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1716_v2,_1716_v2);
        _1735_v20 = _nw250;
        let _index274 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1735_v20).length));
        (_1735_v20)[_index274] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(109),_module.__default.fm1(globalState));
        let _1736_v21;
        _1736_v21 = _dafny.Seq.of(_1725_v10, _1725_v10);
        let _index275 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1735_v20).length));
        let _rhs275 = ((_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_1736_v21, _module.__default.safeIndex(_1716_v2, new BigNumber((_1736_v21).length)), _1725_v10), _1736_v21)) ? (_1730_v15) : (_1730_v15));
        let _rhs276 = _1732_v17;
        let _lhs196 = _1735_v20;
        let _lhs197 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1735_v20).length));
        _1730_v15 = _rhs275;
        _lhs196[_lhs197] = _rhs276;
        _1716_v2 = _1716_v2;
      }
      let _1737_v22;
      _1737_v22 = _dafny.Seq.of(_module.__default.safeDivisionInt(_1716_v2, new BigNumber((_1720_v5).length)));
      let _rhs277 = new BigNumber(964);
      let _rhs278 = _1716_v2;
      let _rhs279 = _1737_v22;
      let _rhs280 = (_dafny.ZERO).minus(_1716_v2);
      let _lhs198 = globalState;
      let _lhs199 = globalState;
      _lhs198.f0 = _rhs277;
      _1716_v2 = _rhs278;
      _1737_v22 = _rhs279;
      _lhs199.f0 = _rhs280;
      let _1738_v23;
      _1738_v23 = _dafny.Seq.UnicodeFromString("lfanlsw");
      let _1739_v24;
      _1739_v24 = _module.D3.create_DC10(_1738_v23, p1, p1);
      let _1740_i1;
      _1740_i1 = _dafny.ZERO;
      L10: {
        while ((_1739_v24).dtor_cf15) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1740_i1)) {
              break L10;
            }
            _1740_i1 = (_1740_i1).plus(_dafny.ONE);
            let _1741_v25;
            _1741_v25 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber(477), globalState));
            (globalState).f1 = new BigNumber((_1741_v25).length);
            let _1742_v26;
            _1742_v26 = _module.D10.create_DC31(p1, p1, _1716_v2);
            let _1743_v27;
            _1743_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v26,false);
            let _1744_v29;
            _1744_v29 = _dafny.MultiSet.fromElements(_1742_v26);
            let _1745_v30;
            _1745_v30 = _dafny.MultiSet.fromElements(_1743_v27, function () {
              let _coll35 = new _dafny.Map();
              for (const _compr_35 of (_1744_v29).Elements) {
                let _1746_v28 = _compr_35;
                if ((_1744_v29).contains(_1746_v28)) {
                  _coll35.push([_1746_v28,p1]);
                }
              }
              return _coll35;
            }());
            let _1747_v31;
            _1747_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.D15.create_DC44(_1745_v30));
            let _1748_v32;
            _1748_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1737_v22,_1747_v31);
            let _1749_v33;
            _1749_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1748_v32,(_this).fm6(_1716_v2, _1716_v2, _1716_v2, false, globalState));
            if (!((((_1749_v33).contains(_1748_v32)) ? ((_1749_v33).get(_1748_v32)) : (p1)))) {
              let _1750_v34;
              _1750_v34 = _module.D1.create_DC4(_module.__default.safeDivisionInt(_1716_v2, _1716_v2), _1718_v4);
              let _rhs281 = !(!(p1));
              let _rhs282 = _1715_v1;
              let _rhs283 = _module.__default.safeModuloInt(_1716_v2, (_1716_v2).multipliedBy(_1716_v2));
              let _rhs284 = _1750_v34;
              let _rhs285 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nn"), _1738_v23), _module.__default.safeIndex((_dafny.ZERO).minus(_1716_v2), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nn"), _1738_v23)).length)), _1715_v1);
              let _lhs200 = globalState;
              r0 = _rhs281;
              _1715_v1 = _rhs282;
              _lhs200.f1 = _rhs283;
              _1750_v34 = _rhs284;
              r1 = _rhs285;
              let _1751_v35;
              _1751_v35 = _dafny.Seq.of(_1741_v25, _1741_v25, _1741_v25, _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_1738_v23).length), globalState), p1, (_this).f10), _1741_v25);
              r0 = _dafny.Seq.contains(_dafny.Seq.Concat(_1751_v35, _1751_v35), _1741_v25);
              (globalState).f0 = _1716_v2;
              let _1752_v36;
              let _init44 = ((_1753_p1) => function (_1754_i2) {
                return _1753_p1;
              })(p1);
              let _nw251 = Array((new BigNumber(4)).toNumber());
              for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw251.length); _i0_44++) {
                _nw251[_i0_44] = _init44(new BigNumber(_i0_44));
              }
              _1752_v36 = _nw251;
              let _index276 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_1752_v36).length));
              (_1752_v36)[_index276] = (_this).f17;
              let _index277 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_1752_v36).length));
              (_1752_v36)[_index277] = !(!((_1741_v25).IsSubsetOf(_1741_v25)));
              (globalState).f0 = (_dafny.ZERO).minus(_1716_v2);
            } else {
              let _1755_v37;
              _1755_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f17);
              let _1756_v38;
              let _nw252 = Array((new BigNumber(24)).toNumber());
              _nw252[(_dafny.ZERO).toNumber()] = _1718_v4;
              _nw252[(_dafny.ONE).toNumber()] = (((((_1755_v37).contains((_this).f17)) ? ((_1755_v37).get((_this).f17)) : ((_this).f17))) ? (_1718_v4) : (_1718_v4));
              _nw252[(new BigNumber(2)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(3)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(4)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(5)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(6)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(7)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(8)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(9)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(10)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(11)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(12)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(13)).toNumber()] = (((_this).f10) ? (_1718_v4) : (_1718_v4));
              _nw252[(new BigNumber(14)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(15)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(16)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(17)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(18)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(19)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(20)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(21)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(22)).toNumber()] = _1718_v4;
              _nw252[(new BigNumber(23)).toNumber()] = _1718_v4;
              _1756_v38 = _nw252;
              let _index278 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1756_v38).length));
              (_1756_v38)[_index278] = _1718_v4;
              let _index279 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1756_v38).length));
              (_1756_v38)[_index279] = _1718_v4;
              let _1757_v39;
              _1757_v39 = _dafny.MultiSet.fromElements(_1716_v2);
              let _1758_v40;
              _1758_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v26,_1716_v2);
              let _rhs286 = !(!(_1758_v40).equals(_1758_v40)) || ((_this).f17);
              let _rhs287 = (_1757_v39).Intersect(_1757_v39);
              let _rhs288 = new BigNumber((_1738_v23).length);
              let _lhs201 = globalState;
              r0 = _rhs286;
              _1757_v39 = _rhs287;
              _lhs201.f7 = _rhs288;
              let _1759_v41;
              let _nw253 = Array((new BigNumber(7)).toNumber());
              _nw253[(_dafny.ZERO).toNumber()] = (_this).f10;
              _nw253[(_dafny.ONE).toNumber()] = (_this).f10;
              _nw253[(new BigNumber(2)).toNumber()] = (_this).f10;
              _nw253[(new BigNumber(3)).toNumber()] = true;
              _nw253[(new BigNumber(4)).toNumber()] = p1;
              _nw253[(new BigNumber(5)).toNumber()] = false;
              _nw253[(new BigNumber(6)).toNumber()] = (_this).f10;
              _1759_v41 = _nw253;
              let _1760_v42;
              _1760_v42 = _module.D0.create_DC2((_this).f17, _1759_v41);
              let _1761_v43;
              _1761_v43 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1716_v2),_1760_v42);
              _1761_v43 = _1761_v43;
              let _1762_v44;
              let _init45 = ((_1763_v2) => function (_1764_i3) {
                return _module.D2.create_DC8(_module.D2.create_DC7(_1763_v2, (_this).f17, new BigNumber((_dafny.Set.fromElements(_1763_v2)).length)));
              })(_1716_v2);
              let _nw254 = Array((new BigNumber(16)).toNumber());
              for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw254.length); _i0_45++) {
                _nw254[_i0_45] = _init45(new BigNumber(_i0_45));
              }
              _1762_v44 = _nw254;
              let _1765_v45;
              _1765_v45 = _module.D2.create_DC6(new _dafny.CodePoint('m'.codePointAt(0)));
              let _1766_v46;
              _1766_v46 = _module.D2.create_DC8(_1765_v45);
              let _index280 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_1762_v44).length));
              (_1762_v44)[_index280] = _1766_v46;
              let _index281 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_1762_v44).length));
              (_1762_v44)[_index281] = _module.__default.fm34(_1716_v2, globalState);
              r0 = (_1716_v2).isLessThanOrEqualTo(_module.__default.safeModuloInt(_1716_v2, _1716_v2));
            }
            (globalState).f1 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1716_v2).plus(_1716_v2)));
            let _1767_v47;
            let _init46 = ((_1768_v2, _1769_v23) => function (_1770_i4) {
              return (_dafny.Set.fromElements(_module.D1.create_DC3(new BigNumber(644)), _module.D1.create_DC3(_1768_v2), _module.D1.create_DC3(new BigNumber((_dafny.MultiSet.fromElements(_1768_v2)).cardinality())), _module.D1.create_DC3(new BigNumber((_1769_v23).length)), _module.D1.create_DC3(_1768_v2))).contains(_module.D1.create_DC3(_1768_v2));
            })(_1716_v2, _1738_v23);
            let _nw255 = Array((new BigNumber(27)).toNumber());
            for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw255.length); _i0_46++) {
              _nw255[_i0_46] = _init46(new BigNumber(_i0_46));
            }
            _1767_v47 = _nw255;
            let _1771_v48;
            _1771_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1716_v2,(_this).f17);
            let _1772_v49;
            _1772_v49 = _dafny.Map.Empty.slice().updateUnsafe((((_1771_v48).contains(_1716_v2)) ? ((_1771_v48).get(_1716_v2)) : ((_this).f10)),(_this).f17);
            let _1773_v50;
            _1773_v50 = _dafny.Seq.of((_this).f17, p1);
            let _index282 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1767_v47).length));
            (_1767_v47)[_index282] = (((_1772_v49).contains(!((_this).f17))) ? ((_1772_v49).get(!((_this).f17))) : ((_1773_v50)[_module.__default.safeIndex(new BigNumber((_1738_v23).length), new BigNumber((_1773_v50).length))]));
            let _index283 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1767_v47).length));
            (_1767_v47)[_index283] = (((_1772_v49).contains(!(p1))) ? ((_1772_v49).get(!(p1))) : ((_this).fm6(_1716_v2, _module.__default.fm1(globalState), _1716_v2, p1, globalState)));
          }
        }
      }
      let _1774_v51;
      let _nw256 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1774_v51 = _nw256;
      let _index284 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1774_v51).length));
      (_1774_v51)[_index284] = _1738_v23;
      let _index285 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1774_v51).length));
      (_1774_v51)[_index285] = _dafny.Seq.Concat(_dafny.Seq.update(_1738_v23, _module.__default.safeIndex(_1716_v2, new BigNumber((_1738_v23).length)), _1715_v1), _dafny.Seq.Concat(_1738_v23, _1738_v23));
      let _1775_v52;
      _1775_v52 = _dafny.Seq.of((_this).f17, p1);
      r0 = (_1775_v52)[_module.__default.safeIndex(_1716_v2, new BigNumber((_1775_v52).length))];
      r0 = (!(p1) || (true)) || (!(_1716_v2).isEqualTo((_dafny.ZERO).minus(_1716_v2)));
      r1 = (_1774_v51)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_1774_v51).length))];
      let _1776_v53;
      _1776_v53 = _dafny.Seq.of(_1737_v22, _dafny.Seq.of(_1716_v2));
      r2 = _module.__default.fm45(_1776_v53, (_dafny.ZERO).minus((_1737_v22)[_module.__default.safeIndex(_1716_v2, new BigNumber((_1737_v22).length))]), globalState);
      return [r0, r1, r2];
    }
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let _1777_v0;
      _1777_v0 = new BigNumber(690);
      let _1778_v1;
      _1778_v1 = _dafny.Seq.UnicodeFromString("hshhrwt");
      let _1779_i0;
      _1779_i0 = _dafny.ZERO;
      L11: {
        while (_dafny.areEqual(_module.__default.fm25(_1777_v0, (_this).f17, _1777_v0, globalState), _1778_v1)) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1779_i0)) {
              break L11;
            }
            _1779_i0 = (_1779_i0).plus(_dafny.ONE);
            let _1780_v2;
            let _nw257 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
            _1780_v2 = _nw257;
            _1780_v2 = _1780_v2;
            let _1781_v3;
            let _nw258 = Array((new BigNumber(28)).toNumber()).fill(false);
            _1781_v3 = _nw258;
            let _1782_v4;
            let _1783_v5;
            let _out53;
            let _out54;
            let _outcollector25 = _module.__default.m0((_this).fm15(globalState), _1781_v3, globalState);
            _out53 = _outcollector25[0];
            _out54 = _outcollector25[1];
            _1782_v4 = _out53;
            _1783_v5 = _out54;
            let _1784_v6;
            _1784_v6 = false;
            _1784_v6 = _module.__default.fm0(_1777_v0, globalState);
            let _1785_v7;
            _1785_v7 = _module.D10.create_DC30(p3);
            let _1786_v8;
            _1786_v8 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (_1785_v7) : (_1785_v7)),_1781_v3);
            _1786_v8 = (_1786_v8).update(_1785_v7, _1781_v3);
          }
        }
      }
      let _1787_v9;
      _1787_v9 = _module.D7.create_DC23();
      let _1788_v10;
      _1788_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1777_v0,(_this).f10);
      let _1789_v11;
      _1789_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1787_v9,new BigNumber((_1788_v10).length));
      let _1790_v12;
      _1790_v12 = _dafny.Seq.of(_1777_v0, (((_1789_v11).contains(_1787_v9)) ? ((_1789_v11).get(_1787_v9)) : (_1777_v0)), _1777_v0);
      let _1791_v13;
      _1791_v13 = _module.D1.create_DC3(new BigNumber((_1790_v12).length));
      _1791_v13 = ((p1) ? (_1791_v13) : (_module.D1.create_DC3(_1777_v0)));
      let _1792_v14;
      _1792_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1778_v1).length),_1777_v0);
      _1792_v14 = (_1792_v14).update(_1777_v0, _1777_v0);
      let _1793_v15;
      let _nw259 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1793_v15 = _nw259;
      let _1794_v16;
      _1794_v16 = _dafny.Seq.of(_1793_v15, _1793_v15, _1793_v15);
      _1793_v15 = (_1794_v16)[_module.__default.safeIndex(_1777_v0, new BigNumber((_1794_v16).length))];
      let _hi7 = _1777_v0;
      for (let _1795_i1 = _1777_v0; _1795_i1.isLessThan(_hi7); _1795_i1 = _1795_i1.plus(_dafny.ONE)) {
        let _1796_v17;
        let _nw260 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1796_v17 = _nw260;
        _1796_v17 = _1796_v17;
        let _1797_v19;
        let _init47 = ((_1798_p2, _1799_p1, _1800_p0) => function (_1801_i2) {
          return (_dafny.MultiSet.fromElements(_module.D4.create_DC13((_module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1798_p2),_1799_p1))).dtor_cf20), _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1800_p0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("neft")).length), new BigNumber((_1800_p0).length)), (_this).f10),(_module.D6.create_DC21(true, (_this).f10, (_this).f17, new BigNumber((_dafny.Seq.UnicodeFromString("cjrt")).length))).dtor_cf36)), _module.D4.create_DC13(function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of (_dafny.Seq.of(_1800_p0)).Elements) {
    let _1802_v18 = _compr_36;
    if (_dafny.Seq.contains(_dafny.Seq.of(_1800_p0), _1802_v18)) {
      _coll36.push([_1802_v18,_1798_p2]);
    }
  }
  return _coll36;
}()), _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f10, (_this).f10, true),(_this).f17)))).Difference(_dafny.MultiSet.fromElements(_module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_1800_p0,(_this).f10)), _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_1800_p0,(_this).f17))));
        })(p2, p1, p0);
        let _nw261 = Array((new BigNumber(22)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw261.length); _i0_47++) {
          _nw261[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1797_v19 = _nw261;
        let _1803_v20;
        _1803_v20 = _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(p0,p1));
        let _index286 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1797_v19).length));
        (_1797_v19)[_index286] = (_dafny.MultiSet.fromElements(_1803_v20)).Intersect(_dafny.MultiSet.fromElements(_1803_v20));
        let _1804_v21;
        _1804_v21 = _dafny.MultiSet.fromElements(_1803_v20, _1803_v20);
        let _index287 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1797_v19).length));
        (_1797_v19)[_index287] = ((_1804_v21).Union(_1804_v21)).update(_1803_v20, _module.__default.abs(_module.__default.safeModuloInt((_1790_v12)[_module.__default.safeIndex(_1795_i1, new BigNumber((_1790_v12).length))], _1795_i1)));
        let _1805_v22;
        _1805_v22 = true;
        let _1806_v23;
        _1806_v23 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1807_v24;
        _1807_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1806_v23,(_this).f17);
        _1805_v22 = (((((_1807_v24).contains(_1806_v23)) ? ((_1807_v24).get(_1806_v23)) : (false))) ? (!(_1805_v22)) : ((_this).f17));
        let _1808_v25;
        _1808_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1777_v0).isLessThan(_1795_i1),(_dafny.ZERO).minus((_dafny.ZERO).minus(_1795_i1)));
        _1808_v25 = (_1808_v25).update((_this).fm7((_this).f10, globalState), _1795_i1);
      }
      let _pat_let_tv37 = _1777_v0;
      let _pat_let_tv38 = _1777_v0;
      (globalState).f7 = function (_source33) {
        if (_source33.is_DC23) {
          return _pat_let_tv37;
        } else if (_source33.is_DC24) {
          let _1809___mcc_h0 = (_source33).cf41;
          let _1810_cf41 = _1809___mcc_h0;
          return new BigNumber(-122);
        } else {
          let _1811___mcc_h1 = (_source33).cf40;
          let _1812_cf40 = _1811___mcc_h1;
          return (_pat_let_tv38).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), function (_1813_i3) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(114)), function (_1814_i4) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }))).length));
        }
      }(_module.D7.create_DC24(_1777_v0));
      let _1815_v26;
      _1815_v26 = _dafny.Set.fromElements(_1777_v0);
      r0 = _1815_v26;
      return r0;
    }
    m9(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _1816_v0;
      _1816_v0 = _dafny.Map.Empty.slice().updateUnsafe(!(!(!((_this).f17))),(_this).f17);
      let _1817_v1;
      _1817_v1 = _module.D10.create_DC31(p1, (_this).f10, p0);
      let _pat_let_tv39 = p1;
      let _pat_let_tv40 = p0;
      let _pat_let_tv41 = p0;
      let _pat_let_tv42 = p0;
      let _pat_let_tv43 = p0;
      let _pat_let_tv44 = p1;
      let _rhs289 = function (_source34) {
        if (_source34.is_DC36) {
          let _1818___mcc_h0 = (_source34).cf60;
          let _1819_cf60 = _1818___mcc_h0;
          if (_pat_let_tv39) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(812)), ((_1820_p0) => function (_1821_i0) {
              return _1820_p0;
            })(_pat_let_tv40))).length);
          } else {
            return _pat_let_tv41;
          }
        } else if (_source34.is_DC37) {
          let _1822___mcc_h1 = (_source34).cf61;
          let _1823___mcc_h2 = (_source34).cf62;
          let _1824___mcc_h3 = (_source34).cf63;
          let _1825_cf63 = _1824___mcc_h3;
          let _1826_cf62 = _1823___mcc_h2;
          let _1827_cf61 = _1822___mcc_h1;
          return _pat_let_tv42;
        } else {
          let _1828___mcc_h4 = (_source34).cf59;
          let _1829_cf59 = _1828___mcc_h4;
          return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_pat_let_tv43)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv44,new BigNumber(-149)))).length);
        }
      }(_module.D12.create_DC35(_dafny.Seq.of(_1817_v1, _module.__default.fm46((_this).f17, (_this).f17, false, globalState), _1817_v1, _1817_v1, _1817_v1)));
      let _rhs290 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10);
      let _lhs202 = globalState;
      _lhs202.f7 = _rhs289;
      _1816_v0 = _rhs290;
      let _1830_v2;
      _1830_v2 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1831_v3;
      _1831_v3 = _dafny.Seq.of(_module.__default.fm4((_this).f10, (_this).f17, p1, globalState));
      let _rhs291 = (_this).f17;
      let _rhs292 = _1830_v2;
      let _rhs293 = _dafny.Seq.IsProperPrefixOf(_1831_v3, _1831_v3);
      let _rhs294 = p0;
      let _lhs203 = globalState;
      r0 = _rhs291;
      _1830_v2 = _rhs292;
      r0 = _rhs293;
      _lhs203.f7 = _rhs294;
      let _1832_v4;
      _1832_v4 = _dafny.MultiSet.fromElements(p0, p0, (p0).plus(p0));
      _1832_v4 = _1832_v4;
      r0 = !((((p2)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((p2).length))]) ? (((_this).f17) || (!((_this).f17))) : ((_this).f10)));
      if (p1) {
        let _1833_v5;
        _1833_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1816_v0);
        _1816_v0 = ((((_1833_v5).contains((_this).f10)) ? ((_1833_v5).get((_this).f10)) : (_1816_v0))).update((_this).f17, p1);
        _1831_v3 = _1831_v3;
        let _1834_v6;
        let _init48 = ((_1835_p1, _1836_p0) => function (_1837_i1) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_1835_p1,_1836_p0)).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.Seq.UnicodeFromString("okyayl")).length)));
        })(p1, p0);
        let _nw262 = Array((new BigNumber(19)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw262.length); _i0_48++) {
          _nw262[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1834_v6 = _nw262;
        let _index288 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length));
        (_1834_v6)[_index288] = p1;
        let _1838_v8;
        _1838_v8 = _dafny.MultiSet.fromElements(p2);
        let _1839_v9;
        _1839_v9 = _dafny.Seq.of(p0, new BigNumber((_1831_v3).length), new BigNumber((function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_1838_v8).Elements) {
            let _1840_v7 = _compr_37;
            if ((_1838_v8).contains(_1840_v7)) {
              _coll37.push([_1840_v7,(_this).f10]);
            }
          }
          return _coll37;
        }()).length));
        let _index289 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length));
        (_1834_v6)[_index289] = !_dafny.Seq.contains(_1839_v9, p0);
        let _1841_v10;
        _1841_v10 = _dafny.Seq.of(_module.__default.fm40((_this).f17, p1, globalState));
        let _index290 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length));
        let _rhs295 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm52(p0, p1, p0, new BigNumber((_1839_v9).length), globalState), _1841_v10), _dafny.Seq.Concat(_dafny.Seq.of(p2, p2, p2), _1841_v10));
        let _rhs296 = (_this).f10;
        let _lhs204 = _1834_v6;
        let _lhs205 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length));
        _1841_v10 = _rhs295;
        _lhs204[_lhs205] = _rhs296;
        let _1842_v11;
        _1842_v11 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f17) ? ((_1834_v6)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length))]) : ((_1834_v6)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length))])),_dafny.Seq.Concat(_1831_v3, _dafny.Seq.UnicodeFromString("meguffaj")));
        let _rhs297 = (((_1842_v11).contains(false)) ? ((_1842_v11).get(false)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(160)), ((_1843_v2) => function (_1844_i2) {
          return _1843_v2;
        })(_1830_v2))));
        let _rhs298 = (((_1834_v6)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length))]) ? (!(p0).isEqualTo(p0)) : ((((_1834_v6)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_1834_v6).length))]) ? ((_this).f17) : (p1))));
        let _rhs299 = _module.__default.fm1(globalState);
        let _lhs206 = globalState;
        _1831_v3 = _rhs297;
        r0 = _rhs298;
        _lhs206.f7 = _rhs299;
      } else {
        let _1845_v12;
        let _nw263 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1845_v12 = _nw263;
        let _1846_v13;
        _1846_v13 = _module.D13.create_DC38(_dafny.MultiSet.fromElements(_1845_v12));
        let _1847_v14;
        _1847_v14 = _dafny.Seq.of(_1831_v3);
        let _1848_v15;
        _1848_v15 = _dafny.Seq.of(p0, new BigNumber(((_1847_v14)[_module.__default.safeIndex(new BigNumber(-690), new BigNumber((_1847_v14).length))]).length), new BigNumber((_1816_v0).length), new BigNumber(638), new BigNumber((_1831_v3).length));
        let _pat_let_tv45 = _1846_v13;
        let _pat_let_tv46 = _1845_v12;
        let _pat_let_tv47 = _1848_v15;
        let _pat_let_tv48 = p0;
        let _pat_let_tv49 = _1848_v15;
        _1846_v13 = function (_pat_let25_0) {
          return function (_1849_dt__update__tmp_h0) {
            return function (_pat_let26_0) {
              return function (_1850_dt__update_hcf64_h0) {
                return _module.D13.create_DC38(_1850_dt__update_hcf64_h0);
              }(_pat_let26_0);
            }(((_pat_let_tv45).dtor_cf64).update(_pat_let_tv46, _module.__default.abs((_pat_let_tv47)[_module.__default.safeIndex((_dafny.ZERO).minus(_pat_let_tv48), new BigNumber((_pat_let_tv49).length))])));
          }(_pat_let25_0);
        }(_1846_v13);
        let _1851_v16;
        let _nw264 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1851_v16 = _nw264;
        _1851_v16 = _1851_v16;
        let _index291 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_1851_v16).length));
        (_1851_v16)[_index291] = _module.__default.safeModuloInt(new BigNumber((_1831_v3).length), p0);
        let _index292 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_1851_v16).length));
        (_1851_v16)[_index292] = p0;
        (globalState).f7 = _module.__default.safeModuloInt((((_1832_v4).contains(p0)) ? ((_1832_v4).get(p0)) : (new BigNumber(-131))), p0);
        let _index293 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_1851_v16).length));
        (_1851_v16)[_index293] = (_dafny.ZERO).minus((_1851_v16)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_1851_v16).length))]);
      }
      let _1852_v17;
      _1852_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_dafny.ZERO).minus(p0));
      let _1853_v18;
      _1853_v18 = _dafny.Set.fromElements(((((_1852_v17).contains(true)) ? ((_1852_v17).get(true)) : (new BigNumber((_1831_v3).length)))).multipliedBy(p0), p0, new BigNumber(565), new BigNumber(((_1832_v4).Intersect(_1832_v4)).cardinality()));
      let _1854_v19;
      _1854_v19 = _dafny.Set.fromElements(_1830_v2, _1830_v2);
      let _1855_v21;
      _1855_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-207));
      let _1856_v22;
      let _nw265 = new _module.C8();
      _nw265.__ctor(_1830_v2);
      _1856_v22 = _nw265;
      let _1857_v23;
      _1857_v23 = _dafny.Seq.of(_module.__default.fm1(globalState));
      let _1858_v24;
      _1858_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1856_v22,_dafny.Seq.update(_1857_v23, _module.__default.safeIndex(p0, new BigNumber((_1857_v23).length)), p0));
      let _1859_v25;
      _1859_v25 = _dafny.Seq.of(_1857_v23, _1857_v23);
      let _1860_v26;
      _1860_v26 = _dafny.MultiSet.fromElements((((_1858_v24).contains(_1856_v22)) ? ((_1858_v24).get(_1856_v22)) : ((_1859_v25)[_module.__default.safeIndex(p0, new BigNumber((_1859_v25).length))])), _1857_v23, _dafny.Seq.of(new BigNumber(587), p0));
      let _1861_v28;
      _1861_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(globalState),_1831_v3);
      let _1862_v29;
      _1862_v29 = _dafny.MultiSet.fromElements(p1);
      let _1863_v30;
      let _nw266 = Array((new BigNumber(27)).toNumber());
      _nw266[(_dafny.ZERO).toNumber()] = (_this).f10;
      _nw266[(_dafny.ONE).toNumber()] = (_1854_v19).IsProperSubsetOf(function () {
        let _coll38 = new _dafny.Set();
        for (const _compr_38 of (_1854_v19).Elements) {
          let _1864_v20 = _compr_38;
          if ((_1854_v19).contains(_1864_v20)) {
            _coll38.add(_1864_v20);
          }
        }
        return _coll38;
      }());
      _nw266[(new BigNumber(2)).toNumber()] = (new BigNumber((_1855_v21).length)).isLessThan((((_1860_v26).contains(_dafny.Seq.of(p0))) ? ((_1860_v26).get(_dafny.Seq.of(p0))) : (p0)));
      _nw266[(new BigNumber(3)).toNumber()] = false;
      _nw266[(new BigNumber(4)).toNumber()] = (((_this).f17) ? ((_this).f17) : ((_this).f10));
      _nw266[(new BigNumber(5)).toNumber()] = p1;
      _nw266[(new BigNumber(6)).toNumber()] = (function () {
        let _coll39 = new _dafny.Map();
        for (const _compr_39 of (_1861_v28).Keys.Elements) {
          let _1865_v27 = _compr_39;
          if ((_1861_v28).contains(_1865_v27)) {
            _coll39.push([(_1865_v27).plus(p0),p0]);
          }
        }
        return _coll39;
      }()).contains(p0);
      _nw266[(new BigNumber(7)).toNumber()] = (p0).isEqualTo(p0);
      _nw266[(new BigNumber(8)).toNumber()] = !(p0).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(430)), ((_1866_p1) => function (_1867_i3) {
        return _dafny.Set.fromElements(_1866_p1, _1866_p1, (_this).f17, !((_this).f10));
      })(p1))).length));
      _nw266[(new BigNumber(9)).toNumber()] = false;
      _nw266[(new BigNumber(10)).toNumber()] = true;
      _nw266[(new BigNumber(11)).toNumber()] = (p0).isEqualTo(p0);
      _nw266[(new BigNumber(12)).toNumber()] = !(true);
      _nw266[(new BigNumber(13)).toNumber()] = ((_this).f17) === ((_this).f10);
      _nw266[(new BigNumber(14)).toNumber()] = (_1862_v29).IsDisjointFrom(_1862_v29);
      _nw266[(new BigNumber(15)).toNumber()] = false;
      _nw266[(new BigNumber(16)).toNumber()] = (((_1816_v0).contains((_this).f17)) ? ((_1816_v0).get((_this).f17)) : ((_this).f10));
      _nw266[(new BigNumber(17)).toNumber()] = (_this).f17;
      _nw266[(new BigNumber(18)).toNumber()] = (true) === ((_this).f10);
      _nw266[(new BigNumber(19)).toNumber()] = p1;
      _nw266[(new BigNumber(20)).toNumber()] = (new BigNumber((_module.__default.fm38((_this).f17, p0, globalState)).length)).isLessThanOrEqualTo(p0);
      _nw266[(new BigNumber(21)).toNumber()] = (p0).isEqualTo(p0);
      _nw266[(new BigNumber(22)).toNumber()] = (_this).f17;
      _nw266[(new BigNumber(23)).toNumber()] = !(new BigNumber((_1831_v3).length)).isEqualTo(p0);
      _nw266[(new BigNumber(24)).toNumber()] = p1;
      _nw266[(new BigNumber(25)).toNumber()] = false;
      _nw266[(new BigNumber(26)).toNumber()] = (_this).f17;
      _1863_v30 = _nw266;
      let _1868_v31;
      _1868_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1863_v30,_1863_v30);
      let _rhs300 = ((_1853_v18).Union(_1853_v18)).Union(_1853_v18);
      let _rhs301 = (_this).f10;
      let _rhs302 = (((_1868_v31).contains(_1863_v30)) ? ((_1868_v31).get(_1863_v30)) : (_1863_v30));
      _1853_v18 = _rhs300;
      r0 = _rhs301;
      _1863_v30 = _rhs302;
      r0 = !(p1) || ((_this).f10);
      return r0;
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f15 = _dafny.ZERO;
      this._f16 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor(f15, f16) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return !((_module.__default.safeDivisionInt((_this).f15, new BigNumber((_dafny.Seq.of((_this).f15)).length))).isEqualTo(_module.__default.safeDivisionInt((_this).f15, (_this).f15)));
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Set.fromElements(new BigNumber(516))).Difference((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_this).f16)).length))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()), new BigNumber(352))))).length);
    };
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _hi8 = _module.__default.safeModuloInt(new BigNumber(339), (_this).f15);
      for (let _1869_i0 = (_this).f15; _1869_i0.isLessThan(_hi8); _1869_i0 = _1869_i0.plus(_dafny.ONE)) {
        if ((_this).f16) {
          let _1870_v0;
          _1870_v0 = _dafny.Seq.UnicodeFromString("ktduqgy");
          let _1871_v1;
          _1871_v1 = _dafny.Set.fromElements((_this).f16);
          let _1872_v2;
          _1872_v2 = _dafny.Set.fromElements(_1871_v1);
          let _1873_v3;
          _1873_v3 = _module.D3.create_DC9((_1872_v2).Union(_1872_v2));
          let _1874_v4;
          _1874_v4 = _module.D0.create_DC1(_module.__default.fm13((_this).f16, !((_this).f16), globalState));
          let _1875_v5;
          _1875_v5 = _dafny.Seq.of(_1872_v2);
          let _pat_let_tv50 = _1875_v5;
          let _pat_let_tv51 = _1875_v5;
          let _pat_let_tv52 = _1872_v2;
          let _rhs303 = (((_this).f16) ? (_1870_v0) : (_dafny.Seq.Concat((_1874_v4).dtor_cf1, _1870_v0)));
          let _rhs304 = function (_pat_let27_0) {
            return function (_1876_dt__update__tmp_h0) {
              return function (_pat_let28_0) {
                return function (_1877_dt__update_hcf13_h0) {
                  return _module.D3.create_DC9(_1877_dt__update_hcf13_h0);
                }(_pat_let28_0);
              }(((_pat_let_tv50)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f15), new BigNumber((_pat_let_tv51).length))]).Difference(_pat_let_tv52));
            }(_pat_let27_0);
          }(_1873_v3);
          let _rhs305 = _module.__default.safeDivisionInt((_this).f15, (_1869_i0).multipliedBy(_1869_i0));
          let _lhs207 = globalState;
          _1870_v0 = _rhs303;
          _1873_v3 = _rhs304;
          _lhs207.f7 = _rhs305;
          let _1878_v6;
          _1878_v6 = _dafny.MultiSet.fromElements((_this).f16);
          let _1879_v7;
          _1879_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber((_1878_v6).cardinality()));
          let _1880_v8;
          _1880_v8 = _module.D4.create_DC14((_this).f15, (_this).f15);
          let _1881_v9;
          _1881_v9 = _dafny.Seq.of(false);
          let _1882_v10;
          let _nw267 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1882_v10 = _nw267;
          let _1883_v11;
          _1883_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1882_v10,(_this).f15);
          let _1884_v12;
          _1884_v12 = _dafny.Seq.of((_this).f15);
          let _1885_v13;
          let _nw268 = Array((new BigNumber(26)).toNumber());
          _nw268[(_dafny.ZERO).toNumber()] = (_this).f15;
          _nw268[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f15);
          _nw268[(new BigNumber(2)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(3)).toNumber()] = new BigNumber((_1871_v1).length);
          _nw268[(new BigNumber(4)).toNumber()] = (_1880_v8).dtor_cf22;
          _nw268[(new BigNumber(5)).toNumber()] = new BigNumber((_1881_v9).length);
          _nw268[(new BigNumber(6)).toNumber()] = (_this).f15;
          _nw268[(new BigNumber(7)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_1869_i0);
          _nw268[(new BigNumber(9)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(10)).toNumber()] = new BigNumber(244);
          _nw268[(new BigNumber(11)).toNumber()] = (_this).f15;
          _nw268[(new BigNumber(12)).toNumber()] = (_this).f15;
          _nw268[(new BigNumber(13)).toNumber()] = (_this).f15;
          _nw268[(new BigNumber(14)).toNumber()] = new BigNumber(435);
          _nw268[(new BigNumber(15)).toNumber()] = new BigNumber(526);
          _nw268[(new BigNumber(16)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(17)).toNumber()] = new BigNumber(484);
          _nw268[(new BigNumber(18)).toNumber()] = new BigNumber((_1883_v11).length);
          _nw268[(new BigNumber(19)).toNumber()] = new BigNumber(-955);
          _nw268[(new BigNumber(20)).toNumber()] = new BigNumber(496);
          _nw268[(new BigNumber(21)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(22)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(23)).toNumber()] = _1869_i0;
          _nw268[(new BigNumber(24)).toNumber()] = (_1884_v12)[_module.__default.safeIndex(new BigNumber((_1870_v0).length), new BigNumber((_1884_v12).length))];
          _nw268[(new BigNumber(25)).toNumber()] = (_this).f15;
          _1885_v13 = _nw268;
          let _1886_v14;
          _1886_v14 = _module.D1.create_DC4((((_1879_v7).contains(!((_this).fm8(new BigNumber((_1870_v0).length), (_this).f16, (_this).f16, globalState)))) ? ((_1879_v7).get(!((_this).fm8(new BigNumber((_1870_v0).length), (_this).f16, (_this).f16, globalState)))) : (_1869_i0)), _1882_v10);
          let _1887_v15;
          _1887_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
          let _1888_v16;
          let _nw269 = new _module.C1();
          _nw269.__ctor((_this).f16);
          _1888_v16 = _nw269;
          let _pat_let_tv53 = _1882_v10;
          let _1889_v17;
          let _nw270 = Array((new BigNumber(12)).toNumber());
          _nw270[(_dafny.ZERO).toNumber()] = _1886_v14;
          _nw270[(_dafny.ONE).toNumber()] = _1886_v14;
          _nw270[(new BigNumber(2)).toNumber()] = _1886_v14;
          _nw270[(new BigNumber(3)).toNumber()] = _1886_v14;
          _nw270[(new BigNumber(4)).toNumber()] = _1886_v14;
          _nw270[(new BigNumber(5)).toNumber()] = function (_pat_let29_0) {
            return function (_1890_dt__update__tmp_h1) {
              return function (_pat_let30_0) {
                return function (_1891_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4((_1890_dt__update__tmp_h1).dtor_cf5, _1891_dt__update_hcf6_h0);
                }(_pat_let30_0);
              }(_pat_let_tv53);
            }(_pat_let29_0);
          }(_1886_v14);
          _nw270[(new BigNumber(6)).toNumber()] = _module.D1.create_DC4((_this).f15, _1882_v10);
          _nw270[(new BigNumber(7)).toNumber()] = _module.D1.create_DC4((((_1887_v15).contains(new BigNumber((_dafny.Seq.of(new BigNumber(652), (_dafny.ZERO).minus((_this).f15), _1869_i0)).length))) ? ((_1887_v15).get(new BigNumber((_dafny.Seq.of(new BigNumber(652), (_dafny.ZERO).minus((_this).f15), _1869_i0)).length))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1870_v0).length),_1888_v16)).length))), _1885_v13);
          _nw270[(new BigNumber(8)).toNumber()] = _module.D1.create_DC4(_1869_i0, _1882_v10);
          _nw270[(new BigNumber(9)).toNumber()] = _1886_v14;
          _nw270[(new BigNumber(10)).toNumber()] = _1886_v14;
          _nw270[(new BigNumber(11)).toNumber()] = _1886_v14;
          _1889_v17 = _nw270;
          let _1892_v18;
          let _nw271 = Array((new BigNumber(2)).toNumber());
          _nw271[(_dafny.ZERO).toNumber()] = _1889_v17;
          _nw271[(_dafny.ONE).toNumber()] = _1889_v17;
          _1892_v18 = _nw271;
          let _index294 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1892_v18).length));
          (_1892_v18)[_index294] = _1889_v17;
          let _index295 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1892_v18).length));
          (_1892_v18)[_index295] = _1889_v17;
          r1 = (_1888_v16).f10;
          let _1893_v19;
          _1893_v19 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1894_v20;
          _1894_v20 = _module.D14.create_DC43(_1893_v19, (_dafny.ZERO).minus(new BigNumber((_1870_v0).length)), _module.__default.fm1(globalState), (_this).f15);
          let _1895_v21;
          _1895_v21 = _dafny.MultiSet.fromElements(_1885_v13, _1885_v13, _1885_v13, _1885_v13, _1882_v10);
          let _1896_v22;
          _1896_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f16)).length),_1895_v21);
          let _1897_v23;
          _1897_v23 = _dafny.Seq.of(_1895_v21);
          let _1898_v24;
          _1898_v24 = _module.D12.create_DC36(_1895_v21);
          let _1899_v25;
          let _nw272 = Array((new BigNumber(18)).toNumber());
          _nw272[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_1882_v10);
          _nw272[(_dafny.ONE).toNumber()] = _1895_v21;
          _nw272[(new BigNumber(2)).toNumber()] = (_1895_v21).Union(_1895_v21);
          _nw272[(new BigNumber(3)).toNumber()] = (((_1896_v22).contains(_1869_i0)) ? ((_1896_v22).get(_1869_i0)) : (_dafny.MultiSet.fromElements(_1882_v10)));
          _nw272[(new BigNumber(4)).toNumber()] = _1895_v21;
          _nw272[(new BigNumber(5)).toNumber()] = _1895_v21;
          _nw272[(new BigNumber(6)).toNumber()] = _1895_v21;
          _nw272[(new BigNumber(7)).toNumber()] = (((_this).f16) ? (_1895_v21) : ((((_1896_v22).contains((_this).f15)) ? ((_1896_v22).get((_this).f15)) : (_1895_v21))));
          _nw272[(new BigNumber(8)).toNumber()] = _1895_v21;
          _nw272[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_1882_v10, _1885_v13, _1882_v10);
          _nw272[(new BigNumber(10)).toNumber()] = (_1897_v23)[_module.__default.safeIndex(_1869_i0, new BigNumber((_1897_v23).length))];
          _nw272[(new BigNumber(11)).toNumber()] = ((_1895_v21).update((_module.D1.create_DC4((_this).f15, _1882_v10)).dtor_cf6, _module.__default.abs(_1869_i0))).Difference(_1895_v21);
          _nw272[(new BigNumber(12)).toNumber()] = (_1895_v21).Intersect(_1895_v21);
          _nw272[(new BigNumber(13)).toNumber()] = _1895_v21;
          _nw272[(new BigNumber(14)).toNumber()] = (_1895_v21).update(_1885_v13, _module.__default.abs(_1869_i0));
          _nw272[(new BigNumber(15)).toNumber()] = (_1895_v21).Union(_1895_v21);
          _nw272[(new BigNumber(16)).toNumber()] = (_1898_v24).dtor_cf60;
          _nw272[(new BigNumber(17)).toNumber()] = _1895_v21;
          _1899_v25 = _nw272;
          let _index296 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1899_v25).length));
          (_1899_v25)[_index296] = (_1895_v21).Union(_1895_v21);
          let _index297 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1899_v25).length));
          let _rhs306 = _1888_v16;
          let _rhs307 = _module.__default.fm53(_1893_v19, (_this).f15, globalState);
          let _rhs308 = (((_this).f16) ? (new BigNumber((_module.__default.fm13((_this).f16, (_this).f16, globalState)).length)) : ((_this).f15));
          let _rhs309 = (_1895_v21).Intersect(_dafny.MultiSet.fromElements(_1882_v10, _1882_v10, _1882_v10, _1882_v10));
          let _lhs208 = globalState;
          let _lhs209 = _1899_v25;
          let _lhs210 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1899_v25).length));
          _1888_v16 = _rhs306;
          _1894_v20 = _rhs307;
          _lhs208.f0 = _rhs308;
          _lhs209[_lhs210] = _rhs309;
          let _1900_v26;
          _1900_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_this).f16);
          let _1901_v27;
          _1901_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1884_v12);
          let _1902_v28;
          _1902_v28 = _module.D13.create_DC39(new BigNumber(((((_1901_v27).contains((_this).f15)) ? ((_1901_v27).get((_this).f15)) : (_1884_v12))).length), (_1888_v16).f10);
          let _nw273 = Array((new BigNumber(29)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = (_this).f15;
          _nw273[(_dafny.ONE).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(2)).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(3)).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1871_v1).length));
          _nw273[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_1869_i0, new BigNumber((_1900_v26).length));
          _nw273[(new BigNumber(6)).toNumber()] = (_module.__default.fm54(globalState)).dtor_cf23;
          _nw273[(new BigNumber(7)).toNumber()] = (new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1870_v0, _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1870_v0).length)), _1893_v19), _module.__default.safeIndex((_this).f15, new BigNumber((_dafny.Seq.update(_1870_v0, _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1870_v0).length)), _1893_v19)).length)), _1893_v19)).length)).multipliedBy(_1869_i0);
          _nw273[(new BigNumber(8)).toNumber()] = _1869_i0;
          _nw273[(new BigNumber(9)).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(10)).toNumber()] = _1869_i0;
          _nw273[(new BigNumber(11)).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(12)).toNumber()] = (((_1878_v6).contains((_1888_v16).f10)) ? ((_1878_v6).get((_1888_v16).f10)) : (new BigNumber((_1870_v0).length)));
          _nw273[(new BigNumber(13)).toNumber()] = ((_this).f15).multipliedBy(new BigNumber((_1900_v26).length));
          _nw273[(new BigNumber(14)).toNumber()] = _1869_i0;
          _nw273[(new BigNumber(15)).toNumber()] = _1869_i0;
          _nw273[(new BigNumber(16)).toNumber()] = (_1902_v28).dtor_cf65;
          _nw273[(new BigNumber(17)).toNumber()] = (_1869_i0).minus((_this).f15);
          _nw273[(new BigNumber(18)).toNumber()] = _module.__default.fm1(globalState);
          _nw273[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.of((_1888_v16).f10)).length);
          _nw273[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f15), _1869_i0);
          _nw273[(new BigNumber(21)).toNumber()] = new BigNumber(64);
          _nw273[(new BigNumber(22)).toNumber()] = new BigNumber(579);
          _nw273[(new BigNumber(23)).toNumber()] = ((_this).f15).multipliedBy(new BigNumber((_1870_v0).length));
          _nw273[(new BigNumber(24)).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(25)).toNumber()] = (_this).f15;
          _nw273[(new BigNumber(26)).toNumber()] = _1869_i0;
          _nw273[(new BigNumber(27)).toNumber()] = _1869_i0;
          _nw273[(new BigNumber(28)).toNumber()] = (_this).f15;
          _1885_v13 = _nw273;
        } else {
          (globalState).f0 = _module.__default.fm1(globalState);
          let _1903_v29;
          let _nw274 = Array((new BigNumber(14)).toNumber());
          _nw274[(_dafny.ZERO).toNumber()] = !(true);
          _nw274[(_dafny.ONE).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(2)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(3)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(4)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(5)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(6)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(7)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(8)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(9)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(10)).toNumber()] = (_this).f16;
          _nw274[(new BigNumber(11)).toNumber()] = false;
          _nw274[(new BigNumber(12)).toNumber()] = false;
          _nw274[(new BigNumber(13)).toNumber()] = (_this).f16;
          _1903_v29 = _nw274;
          let _1904_v30;
          _1904_v30 = _dafny.Set.fromElements(_1903_v29, _1903_v29);
          let _1905_v31;
          let _nw275 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1905_v31 = _nw275;
          let _1906_v32;
          _1906_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1905_v31);
          let _1907_v33;
          _1907_v33 = _dafny.Seq.UnicodeFromString("uwfviquy");
          let _1908_v34;
          _1908_v34 = _dafny.Seq.of(_1905_v31);
          let _1909_v35;
          _1909_v35 = _dafny.Seq.of((((_1906_v32).contains(new BigNumber((_1907_v33).length))) ? ((_1906_v32).get(new BigNumber((_1907_v33).length))) : ((_1908_v34)[_module.__default.safeIndex((_this).f15, new BigNumber((_1908_v34).length))])), _1905_v31);
          let _1910_v36;
          _1910_v36 = _dafny.Set.fromElements(_1869_i0, new BigNumber((_dafny.Set.fromElements((_1909_v35)[_module.__default.safeIndex(_1869_i0, new BigNumber((_1909_v35).length))], _1905_v31)).length));
          let _1911_v37;
          _1911_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1904_v30,_1910_v36);
          _1911_v37 = (_1911_v37).update(_1904_v30, _1910_v36);
          let _1912_v38;
          let _nw276 = new _module.C1();
          _nw276.__ctor((_this).f16);
          _1912_v38 = _nw276;
          r1 = (_this).f16;
          let _1913_v39;
          _1913_v39 = _dafny.Seq.of(new BigNumber(38));
          let _1914_v40;
          _1914_v40 = _dafny.Seq.of((_this).f15, (_1913_v39)[_module.__default.safeIndex(new BigNumber((_1913_v39).length), new BigNumber((_1913_v39).length))]);
          let _index298 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1905_v31).length));
          (_1905_v31)[_index298] = (_dafny.ZERO).minus((_1913_v39)[_module.__default.safeIndex(_1869_i0, new BigNumber((_1913_v39).length))]);
          let _1915_v41;
          _1915_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4((_this).f16, (_this).f16, (_this).f16, globalState),(_this).f16);
          let _1916_v42;
          let _nw277 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _1916_v42 = _nw277;
          let _1917_v43;
          _1917_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1869_i0,_1916_v42);
          let _1918_v44;
          let _nw278 = Array((new BigNumber(23)).toNumber());
          _nw278[(_dafny.ZERO).toNumber()] = _1916_v42;
          _nw278[(_dafny.ONE).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(2)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(3)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(4)).toNumber()] = (((_1917_v43).contains((_this).f15)) ? ((_1917_v43).get((_this).f15)) : (_1916_v42));
          _nw278[(new BigNumber(5)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(6)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(7)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(8)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(9)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(10)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(11)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(12)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(13)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(14)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(15)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(16)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(17)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(18)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(19)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(20)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(21)).toNumber()] = _1916_v42;
          _nw278[(new BigNumber(22)).toNumber()] = _1916_v42;
          _1918_v44 = _nw278;
          let _index299 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1918_v44).length));
          (_1918_v44)[_index299] = _1916_v42;
          let _index300 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1905_v31).length));
          let _index301 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1918_v44).length));
          let _rhs310 = (_module.__default.safeModuloInt(_1869_i0, _module.__default.fm1(globalState))).plus(_1869_i0);
          let _rhs311 = (((_this).f16) ? (_1915_v41) : (_1915_v41));
          let _rhs312 = _1916_v42;
          let _lhs211 = _1905_v31;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1905_v31).length));
          let _lhs213 = _1918_v44;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1918_v44).length));
          _lhs211[_lhs212] = _rhs310;
          _1915_v41 = _rhs311;
          _lhs213[_lhs214] = _rhs312;
        }
        r1 = (_this).f16;
        (globalState).f0 = _1869_i0;
        let _1919_v45;
        let _nw279 = new _module.C1();
        _nw279.__ctor(false);
        _1919_v45 = _nw279;
      }
      let _hi9 = (_this).f15;
      for (let _1920_i1 = new BigNumber(9); _1920_i1.isLessThan(_hi9); _1920_i1 = _1920_i1.plus(_dafny.ONE)) {
        let _1921_v46;
        _1921_v46 = _dafny.MultiSet.fromElements((_this).f16);
        let _1922_v47;
        _1922_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1921_v46).Difference(_1921_v46),(_module.__default.fm42(globalState)).dtor_cf53);
        let _1923_v48;
        _1923_v48 = _dafny.Seq.UnicodeFromString("qbiish");
        let _1924_v49;
        _1924_v49 = _dafny.Seq.of(_module.__default.fm0(_1920_i1, globalState));
        let _1925_v50;
        _1925_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1923_v48).length),(_1924_v49)[_module.__default.safeIndex((_this).f15, new BigNumber((_1924_v49).length))]);
        _1922_v47 = (_1922_v47).update(_1921_v46, _1925_v50);
        let _1926_v51;
        let _nw280 = new _module.C1();
        _nw280.__ctor((_this).f16);
        _1926_v51 = _nw280;
        r1 = (_this).f16;
        let _1927_v52;
        _1927_v52 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1928_v53;
        _1928_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1927_v52);
        let _1929_v54;
        _1929_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1920_i1,(_1928_v53).Merge(_1928_v53));
        _1929_v54 = (_1929_v54).update(_1920_i1, _1928_v53);
      }
      let _1930_v55;
      let _nw281 = Array((new BigNumber(3)).toNumber()).fill([]);
      _1930_v55 = _nw281;
      let _1931_v56;
      let _nw282 = Array((new BigNumber(26)).toNumber()).fill(false);
      _1931_v56 = _nw282;
      let _1932_v57;
      let _nw283 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _1932_v57 = _nw283;
      let _1933_v58;
      _1933_v58 = _module.D1.create_DC4((_this).f15, _1932_v57);
      let _1934_v59;
      _1934_v59 = _dafny.Seq.of((_this).f16);
      let _1935_v60;
      _1935_v60 = _module.D14.create_DC42(_1933_v58, _1934_v59, (_this).f16);
      let _pat_let_tv54 = _1933_v58;
      let _1936_v61;
      let _nw284 = Array((new BigNumber(19)).toNumber());
      _nw284[(_dafny.ZERO).toNumber()] = _1935_v60;
      _nw284[(_dafny.ONE).toNumber()] = function (_pat_let31_0) {
        return function (_1937_dt__update__tmp_h2) {
          return function (_pat_let32_0) {
            return function (_1938_dt__update_hcf74_h0) {
              return function (_pat_let33_0) {
                return function (_1939_dt__update_hcf72_h0) {
                  return _module.D14.create_DC42(_1939_dt__update_hcf72_h0, (_1937_dt__update__tmp_h2).dtor_cf73, _1938_dt__update_hcf74_h0);
                }(_pat_let33_0);
              }(_pat_let_tv54);
            }(_pat_let32_0);
          }((_this).f16);
        }(_pat_let31_0);
      }(_1935_v60);
      _nw284[(new BigNumber(2)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(3)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(4)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(5)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(6)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(7)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(8)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(9)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(10)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(11)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(12)).toNumber()] = _module.D14.create_DC42(_1933_v58, _dafny.Seq.of((_this).f16, (_this).f16), true);
      _nw284[(new BigNumber(13)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(14)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(15)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(16)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(17)).toNumber()] = _1935_v60;
      _nw284[(new BigNumber(18)).toNumber()] = _1935_v60;
      _1936_v61 = _nw284;
      let _1940_v62;
      _1940_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1931_v56,(_module.D16.create_DC48(_1936_v61)).dtor_cf84);
      let _1941_v63;
      _1941_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
      let _1942_v64;
      _1942_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1930_v55);
      let _1943_v65;
      _1943_v65 = _module.D17.create_DC50((((_1942_v64).contains((_this).f16)) ? ((_1942_v64).get((_this).f16)) : (_1930_v55)));
      let _1944_v66;
      _1944_v66 = _dafny.Seq.UnicodeFromString("cdcyl");
      let _rhs313 = (_1943_v65).dtor_cf86;
      let _rhs314 = _1940_v62;
      let _rhs315 = (_this).f15;
      let _rhs316 = _1941_v63;
      let _rhs317 = !_dafny.areEqual(_1944_v66, _1944_v66);
      let _lhs215 = globalState;
      _1930_v55 = _rhs313;
      _1940_v62 = _rhs314;
      _lhs215.f0 = _rhs315;
      _1941_v63 = _rhs316;
      r1 = _rhs317;
      let _1945_v67;
      _1945_v67 = new _dafny.CodePoint('r'.codePointAt(0));
      _1945_v67 = _1945_v67;
      let _index302 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1930_v55).length));
      (_1930_v55)[_index302] = _1931_v56;
      let _1946_v68;
      _1946_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f15, globalState),_1931_v56);
      let _index303 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1930_v55).length));
      (_1930_v55)[_index303] = (((_this).f16) ? ((((_1946_v68).contains((_this).f16)) ? ((_1946_v68).get((_this).f16)) : (_1931_v56))) : (_1931_v56));
      let _arr4 = (_1930_v55)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1930_v55).length))];
      let _index304 = _module.__default.safeIndex(new BigNumber(495), new BigNumber(((_1930_v55)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1930_v55).length))]).length));
      _arr4[_index304] = (_this).f16;
      let _1947_v69;
      _1947_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1944_v66);
      let _arr5 = (_1930_v55)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1930_v55).length))];
      let _index305 = _module.__default.safeIndex(new BigNumber(495), new BigNumber(((_1930_v55)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1930_v55).length))]).length));
      _arr5[_index305] = !(!_dafny.areEqual((((_1947_v69).contains((_this).f16)) ? ((_1947_v69).get((_this).f16)) : (_1944_v66)), _dafny.Seq.Concat(_1944_v66, _1944_v66)));
      let _1948_v70;
      _1948_v70 = _dafny.Set.fromElements(_1944_v66);
      r0 = (_1948_v70).Intersect(_1948_v70);
      r1 = (_this).f16;
      r2 = new BigNumber(299);
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      let _1949_v0;
      _1949_v0 = new _dafny.CodePoint('r'.codePointAt(0));
      let _1950_v1;
      _1950_v1 = _dafny.Seq.UnicodeFromString("sbhekjmp");
      _1949_v0 = (_1950_v1)[_module.__default.safeIndex(p0, new BigNumber((_1950_v1).length))];
      let _1951_v2;
      let _nw285 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
      _1951_v2 = _nw285;
      _1951_v2 = _1951_v2;
      let _1952_v3;
      let _nw286 = Array((new BigNumber(3)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1952_v3 = _nw286;
      let _1953_v4;
      _1953_v4 = _dafny.Seq.of((_this).f16, (_this).f16, (_this).f16);
      let _index306 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_1952_v3).length));
      (_1952_v3)[_index306] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1953_v4, _1953_v4));
      let _1954_v5;
      let _nw287 = new _module.C9();
      _nw287.__ctor((_this).f16, (_this).f16);
      _1954_v5 = _nw287;
      let _1955_v6;
      _1955_v6 = _dafny.Set.fromElements(_1954_v5);
      let _1956_v7;
      _1956_v7 = _dafny.MultiSet.fromElements((_this).f16, (_1955_v6).IsProperSubsetOf(_1955_v6));
      let _index307 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_1952_v3).length));
      (_1952_v3)[_index307] = _1956_v7;
      let _1957_v8;
      let _nw288 = Array((_dafny.ONE).toNumber()).fill(false);
      _1957_v8 = _nw288;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1957_v8).length))) {
        let _1958_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1958_i0)) && ((_1958_i0).isLessThan(new BigNumber((_1957_v8).length))))) {
          (_1957_v8)[(_1958_i0)] = !(!(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,(_this).f16),p0)).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,(_1954_v5).f17),p0)));
        }
      }
      let _1959_v9;
      _1959_v9 = _dafny.Set.fromElements(!(false), (_this).f16);
      let _1960_v10;
      _1960_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1959_v9).length),(_1954_v5).f17);
      let _1961_v12;
      _1961_v12 = _dafny.Map.Empty.slice().updateUnsafe(true,(_1954_v5).f17);
      let _1962_i1;
      _1962_i1 = _dafny.ZERO;
      L12: {
        while (((new BigNumber((_1961_v12).length)).minus(p0)).isLessThan(new BigNumber((function () {
          let _coll49 = new _dafny.Set();
          for (const _compr_49 of (_1960_v10).Keys.Elements) {
            let _2033_v11 = _compr_49;
            if ((_1960_v10).contains(_2033_v11)) {
              _coll49.add(_module.__default.safeDivisionInt(_2033_v11, new BigNumber(340)));
            }
          }
          return _coll49;
        }()).length))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1962_i1)) {
              break L12;
            }
            _1962_i1 = (_1962_i1).plus(_dafny.ONE);
            if ((_this).f16) {
              let _1963_v13;
              let _nw289 = new _module.C7();
              _nw289.__ctor(false, !(!((_1954_v5).f17)));
              _1963_v13 = _nw289;
              let _1964_v14;
              _1964_v14 = false;
              _1964_v14 = (_1963_v13).fm7(((_this).f15).isLessThan((_this).f15), globalState);
              let _1965_v15;
              _1965_v15 = _dafny.Set.fromElements((_this).f15, (_this).f15);
              let _1966_v16;
              _1966_v16 = _dafny.Seq.of(_1965_v15, _1965_v15);
              let _1967_v19;
              _1967_v19 = _dafny.Seq.of(new BigNumber(737));
              let _1968_v20;
              _1968_v20 = _dafny.Seq.of((_this).f15, (_1967_v19)[_module.__default.safeIndex(p0, new BigNumber((_1967_v19).length))]);
              let _1969_v23;
              _1969_v23 = _dafny.MultiSet.fromElements(_module.D15.create_DC46());
              let _1970_v26;
              let _nw290 = Array((new BigNumber(19)).toNumber());
              _nw290[(_dafny.ZERO).toNumber()] = _1965_v15;
              _nw290[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements((_this).f15);
              _nw290[(new BigNumber(2)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(3)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(4)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(5)).toNumber()] = (_1966_v16)[_module.__default.safeIndex(new BigNumber(-357), new BigNumber((_1966_v16).length))];
              _nw290[(new BigNumber(6)).toNumber()] = function () {
                let _coll40 = new _dafny.Set();
                for (const _compr_40 of _dafny.IntegerRange(new BigNumber(993), new BigNumber(733))) {
                  let _1971_v17 = _compr_40;
                  if (((new BigNumber(993)).isLessThanOrEqualTo(_1971_v17)) && ((_1971_v17).isLessThan(new BigNumber(733)))) {
                    _coll40.add((_1971_v17).multipliedBy((_this).f15));
                  }
                }
                return _coll40;
              }();
              _nw290[(new BigNumber(7)).toNumber()] = function () {
                let _coll41 = new _dafny.Set();
                for (const _compr_41 of _dafny.IntegerRange(new BigNumber(113), new BigNumber(960))) {
                  let _1972_v18 = _compr_41;
                  if (((new BigNumber(113)).isLessThanOrEqualTo(_1972_v18)) && ((_1972_v18).isLessThan(new BigNumber(960)))) {
                    _coll41.add((_1972_v18).multipliedBy(new BigNumber((_1953_v4).length)));
                  }
                }
                return _coll41;
              }();
              _nw290[(new BigNumber(8)).toNumber()] = function () {
                let _coll42 = new _dafny.Set();
                for (const _compr_42 of (_1968_v20).Elements) {
                  let _1973_v21 = _compr_42;
                  if (_dafny.Seq.contains(_1968_v20, _1973_v21)) {
                    _coll42.add((_1973_v21).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).cardinality()))));
                  }
                }
                return _coll42;
              }();
              _nw290[(new BigNumber(9)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(10)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(11)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(12)).toNumber()] = function () {
                let _coll43 = new _dafny.Set();
                for (const _compr_43 of _dafny.IntegerRange(new BigNumber(155), new BigNumber(775))) {
                  let _1974_v22 = _compr_43;
                  if (((new BigNumber(155)).isLessThanOrEqualTo(_1974_v22)) && ((_1974_v22).isLessThan(new BigNumber(775)))) {
                    _coll43.add((_1974_v22).multipliedBy((_this).f15));
                  }
                }
                return _coll43;
              }();
              _nw290[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements((_this).f15, p0, new BigNumber((_1969_v23).cardinality()));
              _nw290[(new BigNumber(14)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(15)).toNumber()] = _1965_v15;
              _nw290[(new BigNumber(16)).toNumber()] = function () {
                let _coll44 = new _dafny.Set();
                for (const _compr_44 of (_1960_v10).Keys.Elements) {
                  let _1975_v24 = _compr_44;
                  if ((_1960_v10).contains(_1975_v24)) {
                    _coll44.add(_module.__default.safeDivisionInt(_1975_v24, new BigNumber(-723)));
                  }
                }
                return _coll44;
              }();
              _nw290[(new BigNumber(17)).toNumber()] = function () {
                let _coll45 = new _dafny.Set();
                for (const _compr_45 of _dafny.IntegerRange(new BigNumber(405), new BigNumber(863))) {
                  let _1976_v25 = _compr_45;
                  if (((new BigNumber(405)).isLessThanOrEqualTo(_1976_v25)) && ((_1976_v25).isLessThan(new BigNumber(863)))) {
                    _coll45.add(_module.__default.safeModuloInt(_1976_v25, new BigNumber((_dafny.Set.fromElements(_module.D16.create_DC49((_dafny.ZERO).minus(p0)), _module.D16.create_DC49((_this).f15))).length)));
                  }
                }
                return _coll45;
              }();
              _nw290[(new BigNumber(18)).toNumber()] = _1965_v15;
              _1970_v26 = _nw290;
              let _1977_v27;
              _1977_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1960_v10,_1970_v26);
              _1977_v27 = (_1977_v27).update(_1960_v10, _1970_v26);
              _1953_v4 = _dafny.Seq.Concat(_1953_v4, _1953_v4);
              _1964_v14 = true;
            } else {
              let _1978_v28;
              _1978_v28 = false;
              let _1979_v29;
              _1979_v29 = _dafny.MultiSet.fromElements(_1959_v9);
              _1978_v28 = (_dafny.MultiSet.fromElements(_1959_v9)).IsDisjointFrom((_1979_v29).Intersect(_1979_v29));
              let _1980_v30;
              _1980_v30 = _dafny.Seq.of(_1957_v8, _1957_v8, _1957_v8, _1957_v8, _1957_v8);
              let _1981_v31;
              _1981_v31 = _dafny.Seq.of(_1957_v8, _1957_v8, (_1980_v30)[_module.__default.safeIndex(p0, new BigNumber((_1980_v30).length))], _1957_v8, _1957_v8);
              let _1982_v32;
              _1982_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1981_v31, _module.__default.safeIndex((_this).f15, new BigNumber((_1981_v31).length)), _1957_v8),(_this).f16);
              _1982_v32 = (_1982_v32).update(_1980_v30, ((_this).f15).isLessThanOrEqualTo((_this).f15));
              _1950_v1 = _dafny.Seq.Concat(_1950_v1, _1950_v1);
              let _1983_v33;
              _1983_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,p0);
              let _1984_v34;
              let _nw291 = Array((new BigNumber(24)).toNumber());
              _nw291[(_dafny.ZERO).toNumber()] = _1983_v33;
              _nw291[(_dafny.ONE).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(2)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(3)).toNumber()] = (_1983_v33).Merge(_1983_v33);
              _nw291[(new BigNumber(4)).toNumber()] = (_1983_v33).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1978_v28,(_this).f15));
              _nw291[(new BigNumber(5)).toNumber()] = ((_1983_v33).update(true, new BigNumber(57))).Merge(_1983_v33);
              _nw291[(new BigNumber(6)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(7)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,(_this).f15);
              _nw291[(new BigNumber(9)).toNumber()] = (_1983_v33).update(_1978_v28, p0);
              _nw291[(new BigNumber(10)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(11)).toNumber()] = (_1983_v33).Merge(_1983_v33);
              _nw291[(new BigNumber(12)).toNumber()] = (_1983_v33).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,p0));
              _nw291[(new BigNumber(13)).toNumber()] = (_1983_v33).Merge((_1983_v33).update((_1954_v5).f17, (_this).f15));
              _nw291[(new BigNumber(14)).toNumber()] = (_1983_v33).Merge(_1983_v33);
              _nw291[(new BigNumber(15)).toNumber()] = (_1983_v33).update((_this).f16, (_this).f15);
              _nw291[(new BigNumber(16)).toNumber()] = (((_1954_v5).f17) ? (_1983_v33) : (_1983_v33));
              _nw291[(new BigNumber(17)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,_module.__default.fm1(globalState));
              _nw291[(new BigNumber(19)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(20)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(21)).toNumber()] = (_1983_v33).Merge(_1983_v33);
              _nw291[(new BigNumber(22)).toNumber()] = _1983_v33;
              _nw291[(new BigNumber(23)).toNumber()] = _1983_v33;
              _1984_v34 = _nw291;
              _1984_v34 = _1984_v34;
              _1949_v0 = _1949_v0;
            }
            let _1985_v35;
            _1985_v35 = _module.D17.create_DC51(_1949_v0, (_this).f16);
            let _source35 = _1985_v35;
            if (_source35.is_DC51) {
              let _1986___mcc_h0 = (_source35).cf87;
              let _1987___mcc_h1 = (_source35).cf88;
              let _1988_cf88 = _1987___mcc_h1;
              let _1989_cf87 = _1986___mcc_h0;
              (globalState).f0 = (_this).f15;
              _1988_cf88 = !(((_this).f15).plus(p0)).isEqualTo(p0);
              let _1990_v36;
              _1990_v36 = _module.D6.create_DC21((_1954_v5).f17, _1988_cf88, (_this).f16, (_this).f15);
              let _1991_v37;
              let _out55;
              _out55 = (_1954_v5).m8(_1953_v4, _1988_cf88, (((_1990_v36).dtor_cf36) ? (false) : ((_1954_v5).fm6(p0, (_this).f15, (_this).f15, (_1954_v5).f17, globalState))), _1959_v9, globalState);
              _1991_v37 = _out55;
              (globalState).f0 = _module.__default.safeDivisionInt(p0, (_this).f15);
            } else if (_source35.is_DC50) {
              let _1992___mcc_h2 = (_source35).cf86;
              let _1993_cf86 = _1992___mcc_h2;
              let _index308 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1957_v8).length));
              (_1957_v8)[_index308] = (_1954_v5).f17;
              let _index309 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1957_v8).length));
              (_1957_v8)[_index309] = !(!((_this).f16));
              let _1994_v38;
              _1994_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,_1949_v0);
              _1994_v38 = (_1994_v38).update(!((_this).f16) || (!((_1954_v5).f17)), (_1950_v1)[_module.__default.safeIndex(p0, new BigNumber((_1950_v1).length))]);
              let _1995_v39;
              let _nw292 = new _module.C8();
              _nw292.__ctor(_1949_v0);
              _1995_v39 = _nw292;
              let _1996_v40;
              _1996_v40 = _module.D14.create_DC43(_1949_v0, new BigNumber((_dafny.Set.fromElements((_this).f15, (_this).f15)).length), (_this).f15, new BigNumber(994));
              let _pat_let_tv55 = p0;
              let _pat_let_tv56 = p0;
              let _pat_let_tv57 = _1949_v0;
              let _pat_let_tv58 = p0;
              let _1997_v41;
              let _nw293 = Array((new BigNumber(29)).toNumber());
              _nw293[(_dafny.ZERO).toNumber()] = _1996_v40;
              _nw293[(_dafny.ONE).toNumber()] = (((_1954_v5).f17) ? (_1996_v40) : (_1996_v40));
              _nw293[(new BigNumber(2)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(3)).toNumber()] = function (_pat_let34_0) {
                return function (_2001_dt__update__tmp_h1) {
                  return function (_pat_let38_0) {
                    return function (_2002_dt__update_hcf75_h0) {
                      return function (_pat_let39_0) {
                        return function (_2003_dt__update_hcf77_h0) {
                          return _module.D14.create_DC43(_2002_dt__update_hcf75_h0, (_2001_dt__update__tmp_h1).dtor_cf76, _2003_dt__update_hcf77_h0, (_2001_dt__update__tmp_h1).dtor_cf78);
                        }(_pat_let39_0);
                      }((_this).f15);
                    }(_pat_let38_0);
                  }(_pat_let_tv57);
                }(_pat_let34_0);
              }(function (_pat_let35_0) {
                return function (_1998_dt__update__tmp_h0) {
                  return function (_pat_let36_0) {
                    return function (_1999_dt__update_hcf76_h0) {
                      return function (_pat_let37_0) {
                        return function (_2000_dt__update_hcf78_h0) {
                          return _module.D14.create_DC43((_1998_dt__update__tmp_h0).dtor_cf75, _1999_dt__update_hcf76_h0, (_1998_dt__update__tmp_h0).dtor_cf77, _2000_dt__update_hcf78_h0);
                        }(_pat_let37_0);
                      }(_pat_let_tv56);
                    }(_pat_let36_0);
                  }(_pat_let_tv55);
                }(_pat_let35_0);
              }(_1996_v40));
              _nw293[(new BigNumber(4)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(5)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(6)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(7)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(8)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(9)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(10)).toNumber()] = _module.__default.fm53(_1995_v39.f18, p0, globalState);
              _nw293[(new BigNumber(11)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(12)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(13)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(14)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(15)).toNumber()] = _module.D14.create_DC43(_1949_v0, new BigNumber(14), p0, (_dafny.ZERO).minus((_this).f15));
              _nw293[(new BigNumber(16)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(17)).toNumber()] = _module.D14.create_DC43(_1949_v0, p0, p0, p0);
              _nw293[(new BigNumber(18)).toNumber()] = _module.__default.fm53(_1995_v39.f18, (_this).f15, globalState);
              _nw293[(new BigNumber(19)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(20)).toNumber()] = _module.D14.create_DC43(_1995_v39.f18, (_this).f15, new BigNumber((_1950_v1).length), (_this).f15);
              _nw293[(new BigNumber(21)).toNumber()] = _module.D14.create_DC43(new _dafny.CodePoint('d'.codePointAt(0)), (_this).f15, p0, p0);
              _nw293[(new BigNumber(22)).toNumber()] = _module.D14.create_DC43(_1995_v39.f18, new BigNumber(380), (_this).f15, new BigNumber((_1950_v1).length));
              _nw293[(new BigNumber(23)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(24)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(25)).toNumber()] = function (_pat_let40_0) {
                return function (_2004_dt__update__tmp_h2) {
                  return function (_pat_let41_0) {
                    return function (_2005_dt__update_hcf76_h1) {
                      return _module.D14.create_DC43((_2004_dt__update__tmp_h2).dtor_cf75, _2005_dt__update_hcf76_h1, (_2004_dt__update__tmp_h2).dtor_cf77, (_2004_dt__update__tmp_h2).dtor_cf78);
                    }(_pat_let41_0);
                  }(_pat_let_tv58);
                }(_pat_let40_0);
              }(_1996_v40);
              _nw293[(new BigNumber(26)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(27)).toNumber()] = _1996_v40;
              _nw293[(new BigNumber(28)).toNumber()] = _1996_v40;
              _1997_v41 = _nw293;
              let _index310 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1997_v41).length));
              (_1997_v41)[_index310] = _1996_v40;
              let _index311 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1997_v41).length));
              let _rhs318 = _module.__default.safeModuloInt((_this).f15, (_this).f15);
              let _rhs319 = _1996_v40;
              let _lhs216 = globalState;
              let _lhs217 = _1997_v41;
              let _lhs218 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1997_v41).length));
              _lhs216.f1 = _rhs318;
              _lhs217[_lhs218] = _rhs319;
            } else {
              let _2006___mcc_h3 = (_source35).cf89;
              let _2007_cf89 = _2006___mcc_h3;
              let _2008_v42;
              let _nw294 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
              _2008_v42 = _nw294;
              let _2009_v43;
              let _init49 = ((_2010_v4, _2011_v5) => function (_2012_i2) {
                return _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_2010_v4,(_2011_v5).f17));
              })(_1953_v4, _1954_v5);
              let _nw295 = Array((new BigNumber(13)).toNumber());
              for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw295.length); _i0_49++) {
                _nw295[_i0_49] = _init49(new BigNumber(_i0_49));
              }
              _2009_v43 = _nw295;
              let _2013_v44;
              _2013_v44 = _module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1953_v4, _module.__default.safeIndex((_this).f15, new BigNumber((_1953_v4).length)), (_1954_v5).f17),(_1954_v5).f17));
              let _index312 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_2009_v43).length));
              (_2009_v43)[_index312] = _2013_v44;
              let _index313 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_2009_v43).length));
              let _rhs320 = _2008_v42;
              let _rhs321 = _2013_v44;
              let _lhs219 = _2009_v43;
              let _lhs220 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_2009_v43).length));
              _2008_v42 = _rhs320;
              _lhs219[_lhs220] = _rhs321;
              let _2014_v45;
              _2014_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1961_v12).update((_1954_v5).f17, (_1954_v5).f17),_1949_v0);
              (globalState).f7 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).fm9(_dafny.Seq.UnicodeFromString("vxqtm"), function () {
                let _coll46 = new _dafny.Set();
                for (const _compr_46 of (_2014_v45).Keys.Elements) {
                  let _2015_v46 = _compr_46;
                  if ((_2014_v45).contains(_2015_v46)) {
                    _coll46.add(_2015_v46);
                  }
                }
                return _coll46;
              }(), globalState)));
              let _2016_v47;
              _2016_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1950_v1,p0);
              let _2017_v48;
              _2017_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber((_2016_v47).length));
              let _2018_v49;
              _2018_v49 = _dafny.MultiSet.fromElements(new BigNumber((_2017_v48).length), (_this).f15);
              let _2019_v50;
              let _nw296 = new _module.C5();
              _nw296.__ctor();
              _2019_v50 = _nw296;
              let _2020_v51;
              _2020_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
              let _2021_v52;
              _2021_v52 = _dafny.Seq.of((((_2020_v51).contains(new BigNumber((_2020_v51).length))) ? ((_2020_v51).get(new BigNumber((_2020_v51).length))) : (p0)));
              let _2022_v53;
              _2022_v53 = _dafny.Map.Empty.slice().updateUnsafe((_2018_v49).Difference(_dafny.MultiSet.fromElements(p0, new BigNumber((_dafny.Set.fromElements(_2019_v50)).length), new BigNumber((_1950_v1).length))),_dafny.Seq.Concat(_2021_v52, _2021_v52));
              _2022_v53 = (_2022_v53).Merge(_2022_v53);
              let _2023_v54;
              let _nw297 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
              _2023_v54 = _nw297;
              let _index314 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_2023_v54).length));
              (_2023_v54)[_index314] = _dafny.Map.Empty.slice().updateUnsafe(_1961_v12,(_1954_v5).f17);
              let _2024_v55;
              _2024_v55 = _dafny.Map.Empty.slice().updateUnsafe((_1961_v12).Merge(_1961_v12),true);
              let _2025_v56;
              _2025_v56 = _dafny.Seq.of(_1960_v10, _dafny.Map.Empty.slice().updateUnsafe((_this).fm9(_1950_v1, _dafny.Set.fromElements(_1961_v12), globalState),(_1954_v5).f17), _1960_v10, _1960_v10);
              let _index315 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_2023_v54).length));
              let _rhs322 = _1949_v0;
              let _rhs323 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2021_v52, _2021_v52), _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_1959_v9).length), (_this).f15), new BigNumber((_dafny.Seq.Concat(_2021_v52, _2021_v52)).length)), new BigNumber((_2025_v56).length))).length);
              let _rhs324 = _2024_v55;
              let _rhs325 = (((_1954_v5).f17) ? (_2024_v55) : ((function () {
                let _coll47 = new _dafny.Map();
                for (const _compr_47 of (_2024_v55).Keys.Elements) {
                  let _2026_v57 = _compr_47;
                  if ((_2024_v55).contains(_2026_v57)) {
                    _coll47.push([_2026_v57,(_1954_v5).f17]);
                  }
                }
                return _coll47;
              }()).update(_1961_v12, (_1954_v5).f17)));
              let _lhs221 = globalState;
              let _lhs222 = _2023_v54;
              let _lhs223 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_2023_v54).length));
              _1949_v0 = _rhs322;
              _lhs221.f1 = _rhs323;
              _lhs222[_lhs223] = _rhs324;
              _2024_v55 = _rhs325;
            }
            let _2027_v58;
            let _nw298 = new _module.C6();
            _nw298.__ctor();
            _2027_v58 = _nw298;
            let _2028_v59;
            _2028_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_dafny.Set.fromElements(_1949_v0));
            let _2029_v61;
            _2029_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber(((((_2028_v59).contains(true)) ? ((_2028_v59).get(true)) : (function () {
              let _coll48 = new _dafny.Set();
              for (const _compr_48 of (_dafny.Map.Empty.slice().updateUnsafe(_1949_v0,(_this).f15)).Keys.Elements) {
                let _2030_v60 = _compr_48;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_1949_v0,(_this).f15)).contains(_2030_v60)) {
                  _coll48.add(_2030_v60);
                }
              }
              return _coll48;
            }()))).length));
            let _2031_v62;
            _2031_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_2029_v61);
            let _2032_v63;
            _2032_v63 = _dafny.Map.Empty.slice().updateUnsafe((((_2031_v62).contains(true)) ? ((_2031_v62).get(true)) : (_2029_v61)),(_dafny.Set.fromElements(true, (_1954_v5).f17, (_this).f16)).Union(_1959_v9));
            _2032_v63 = (_2032_v63).update(_2029_v61, _dafny.Set.fromElements((_1954_v5).f17, (_this).f16, (_this).f16, (_this).f16));
          }
        }
      }
      if ((_this).f16) {
        (globalState).f7 = new BigNumber(449);
        let _2034_v64;
        let _nw299 = Array((new BigNumber(2)).toNumber());
        _nw299[(_dafny.ZERO).toNumber()] = _1949_v0;
        _nw299[(_dafny.ONE).toNumber()] = _1949_v0;
        _2034_v64 = _nw299;
        let _2035_v65;
        _2035_v65 = _module.D17.create_DC51(_1949_v0, (_1954_v5).f17);
        let _index316 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_2034_v64).length));
        (_2034_v64)[_index316] = (((_1954_v5).f17) ? (new _dafny.CodePoint('i'.codePointAt(0))) : ((_2035_v65).dtor_cf87));
        let _2036_v66;
        _2036_v66 = false;
        let _2037_v67;
        _2037_v67 = _module.D5.create_DC17((_1954_v5).f17, new _dafny.CodePoint('o'.codePointAt(0)));
        let _2038_v68;
        _2038_v68 = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,(_2037_v67).dtor_cf29);
        let _2039_v69;
        _2039_v69 = _dafny.Map.Empty.slice().updateUnsafe((((_2038_v68).contains((_1954_v5).f17)) ? ((_2038_v68).get((_1954_v5).f17)) : (_1949_v0)),_2036_v66);
        let _2040_v70;
        _2040_v70 = _dafny.Seq.of(_2039_v69, _2039_v69, _2039_v69);
        let _2041_v72;
        _2041_v72 = _module.D18.create_DC53(_dafny.Seq.of(function () {
  let _coll50 = new _dafny.Map();
  for (const _compr_50 of (_1950_v1).Elements) {
    let _2042_v71 = _compr_50;
    if (_dafny.Seq.contains(_1950_v1, _2042_v71)) {
      _coll50.push([_2042_v71,true]);
    }
  }
  return _coll50;
}(), _2039_v69));
        let _index317 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_2034_v64).length));
        let _rhs326 = _1949_v0;
        let _rhs327 = (_this).f16;
        let _rhs328 = _dafny.Seq.update((_2041_v72).dtor_cf90, _module.__default.safeIndex(p0, new BigNumber(((_2041_v72).dtor_cf90).length)), _2039_v69);
        let _lhs224 = _2034_v64;
        let _lhs225 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_2034_v64).length));
        _lhs224[_lhs225] = _rhs326;
        _2036_v66 = _rhs327;
        _2040_v70 = _rhs328;
        if ((_1954_v5).f17) {
          (globalState).f1 = p0;
          let _2043_v73;
          let _nw300 = Array((_dafny.ONE).toNumber()).fill([]);
          _2043_v73 = _nw300;
          _2043_v73 = _2043_v73;
          let _2044_v74;
          _2044_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(852),_1957_v8);
          _2044_v74 = (_2044_v74).update((_this).f15, _1957_v8);
          (globalState).f1 = new BigNumber(239);
          _1950_v1 = _dafny.Seq.Concat(_1950_v1, _dafny.Seq.Concat(_1950_v1, _1950_v1));
        } else {
          let _2045_v75;
          _2045_v75 = _dafny.MultiSet.fromElements((_this).f15);
          let _2046_v76;
          _2046_v76 = _dafny.Set.fromElements(((((_2045_v75).contains(new BigNumber(419))) ? ((_2045_v75).get(new BigNumber(419))) : ((_this).f15))).minus((_this).f15));
          _2046_v76 = ((_2046_v76).Intersect(function () {
            let _coll51 = new _dafny.Set();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(614), new BigNumber(831))) {
              let _2047_v77 = _compr_51;
              if (((new BigNumber(614)).isLessThanOrEqualTo(_2047_v77)) && ((_2047_v77).isLessThan(new BigNumber(831)))) {
                _coll51.add(_module.__default.safeDivisionInt(_2047_v77, (_this).f15));
              }
            }
            return _coll51;
          }())).Intersect((_2046_v76).Union(function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of (_1960_v10).Keys.Elements) {
              let _2048_v78 = _compr_52;
              if ((_1960_v10).contains(_2048_v78)) {
                _coll52.add(_module.__default.safeModuloInt(_2048_v78, new BigNumber(186)));
              }
            }
            return _coll52;
          }()));
          let _index318 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_1957_v8).length));
          (_1957_v8)[_index318] = (_1954_v5).f17;
          let _index319 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_1957_v8).length));
          let _rhs329 = _2036_v66;
          let _rhs330 = (_1956_v7).IsSubsetOf(_dafny.MultiSet.fromElements((_1954_v5).f17, (_this).f16));
          let _rhs331 = p0;
          let _rhs332 = _1950_v1;
          let _rhs333 = (_1954_v5).f17;
          let _lhs226 = globalState;
          let _lhs227 = _1957_v8;
          let _lhs228 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_1957_v8).length));
          _2036_v66 = _rhs329;
          _2036_v66 = _rhs330;
          _lhs226.f0 = _rhs331;
          _1950_v1 = _rhs332;
          _lhs227[_lhs228] = _rhs333;
          let _2049_v79;
          _2049_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1950_v1,(p0).minus(p0));
          _2049_v79 = _2049_v79;
          let _2050_v80;
          _2050_v80 = _module.D3.create_DC11((_this).f15, _1961_v12);
          _1961_v12 = (_2050_v80).dtor_cf18;
          let _2051_v81;
          let _2052_v82;
          let _2053_v83;
          let _2054_v84;
          let _out56;
          let _out57;
          let _out58;
          let _out59;
          let _outcollector26 = (_this).m7((_1954_v5).f17, true, globalState);
          _out56 = _outcollector26[0];
          _out57 = _outcollector26[1];
          _out58 = _outcollector26[2];
          _out59 = _outcollector26[3];
          _2051_v81 = _out56;
          _2052_v82 = _out57;
          _2053_v83 = _out58;
          _2054_v84 = _out59;
        }
        _2036_v66 = !((_1954_v5).f17) || (_dafny.Seq.IsProperPrefixOf(_1950_v1, _1950_v1));
        let _2055_v85;
        _2055_v85 = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,_1950_v1);
        (globalState).f7 = _module.__default.safeDivisionInt(new BigNumber((_2055_v85).length), p0);
      } else {
        let _2056_v86;
        let _nw301 = new _module.C8();
        _nw301.__ctor(_1949_v0);
        _2056_v86 = _nw301;
        let _2057_v87;
        let _nw302 = new _module.C0();
        _nw302.__ctor();
        _2057_v87 = _nw302;
        if ((_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements((_1954_v5).f17, (_1954_v5).f17)).cardinality()), p0)).isLessThanOrEqualTo((_dafny.ZERO).minus(p0))) {
          let _2058_v88;
          _2058_v88 = true;
          let _2059_v89;
          _2059_v89 = _dafny.Seq.of(new BigNumber(164), (_dafny.ZERO).minus(new BigNumber((_1950_v1).length)));
          let _2060_v90;
          _2060_v90 = _dafny.Seq.of(new BigNumber((_1950_v1).length), new BigNumber((_2059_v89).length));
          let _2061_v91;
          _2061_v91 = _dafny.Seq.of(_2059_v89);
          _2058_v88 = _dafny.Seq.IsProperPrefixOf(_2061_v91, _2061_v91);
          let _2062_v92;
          let _nw303 = Array((new BigNumber(20)).toNumber()).fill([]);
          _2062_v92 = _nw303;
          let _2063_v93;
          _2063_v93 = _dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,(((_1954_v5).f17) ? (_2062_v92) : (_2062_v92)));
          let _2064_v94;
          let _nw304 = new _module.C7();
          _nw304.__ctor((_1954_v5).f17, !((_this).f16));
          _2064_v94 = _nw304;
          let _2065_v95;
          _2065_v95 = _dafny.Map.Empty.slice().updateUnsafe(_2064_v94,(_dafny.ZERO).minus((_this).f15));
          _2063_v93 = (_2063_v93).update((p0).isEqualTo(new BigNumber((_2065_v95).length)), _2062_v92);
          let _2066_v96;
          let _nw305 = new _module.C4();
          _nw305.__ctor();
          _2066_v96 = _nw305;
          let _2067_v98;
          _2067_v98 = _module.D19.create_DC57(function () {
  let _coll53 = new _dafny.Set();
  for (const _compr_53 of _dafny.IntegerRange(new BigNumber(74), new BigNumber(436))) {
    let _2068_v97 = _compr_53;
    if (((new BigNumber(74)).isLessThanOrEqualTo(_2068_v97)) && ((_2068_v97).isLessThan(new BigNumber(436)))) {
      _coll53.add((_2068_v97).multipliedBy(p0));
    }
  }
  return _coll53;
}());
          let _2069_v99;
          _2069_v99 = _dafny.Set.fromElements((_2067_v98).dtor_cf98);
          let _2070_v100;
          _2070_v100 = _dafny.MultiSet.fromElements((_this).f15, p0, (new BigNumber((_2069_v99).length)).multipliedBy(p0));
          _2070_v100 = ((((_this).f16) ? (_2070_v100) : (_2070_v100))).Union(_2070_v100);
          (globalState).f7 = _module.__default.safeModuloInt((_2059_v89)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_2064_v94).f10)).length), new BigNumber((_2059_v89).length))], _module.__default.safeDivisionInt((_this).f15, new BigNumber((_1959_v9).length)));
        } else {
          let _2071_v101;
          _2071_v101 = _dafny.Map.Empty.slice().updateUnsafe(_2056_v86.f18,(_this).f16);
          let _2072_v102;
          let _nw306 = new _module.C2();
          _nw306.__ctor((_1954_v5).f17, (_this).f15);
          _2072_v102 = _nw306;
          let _2073_v103;
          _2073_v103 = _dafny.Map.Empty.slice().updateUnsafe((((_2071_v101).contains(_2056_v86.f18)) ? ((_2071_v101).get(_2056_v86.f18)) : ((_1954_v5).f17)),_2072_v102);
          _2073_v103 = (_2073_v103).update((_1954_v5).f17, _2072_v102);
          let _rhs334 = ((_this).f15).minus((_this).f15);
          let _rhs335 = _2056_v86.f18;
          let _rhs336 = (_dafny.ZERO).minus((p0).minus((((_1954_v5).f17) ? (p0) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(334)), ((_2074_v86) => function (_2075_i3) {
            return _2074_v86.f18;
          })(_2056_v86))).length)))));
          let _lhs229 = globalState;
          let _lhs230 = _2056_v86;
          let _lhs231 = globalState;
          _lhs229.f1 = _rhs334;
          _lhs230.f18 = _rhs335;
          _lhs231.f0 = _rhs336;
          let _2076_v104;
          let _init50 = ((_2077_v1) => function (_2078_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_2077_v1,_2077_v1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pabfkukml"),_dafny.Seq.Create(_module.__default.abs(new BigNumber(102)), function (_2079_i5) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            })));
          })(_1950_v1);
          let _nw307 = Array((new BigNumber(20)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw307.length); _i0_50++) {
            _nw307[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _2076_v104 = _nw307;
          let _2080_v105;
          _2080_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1950_v1,_1950_v1);
          let _index320 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2076_v104).length));
          (_2076_v104)[_index320] = (_2080_v105).Merge(_2080_v105);
          let _2081_v106;
          _2081_v106 = _dafny.Seq.of((_2080_v105).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1950_v1,_1950_v1)));
          let _2082_v107;
          _2082_v107 = _dafny.Set.fromElements(_1961_v12, _1961_v12);
          let _2083_v108;
          _2083_v108 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm55((_this).f15, (_this).f15, (_1954_v5).f17, globalState),(_this).fm9(_1950_v1, _2082_v107, globalState));
          let _index321 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2076_v104).length));
          (_2076_v104)[_index321] = (_2081_v106)[_module.__default.safeIndex((((_1954_v5).f17) ? ((_this).f15) : (new BigNumber((_2083_v108).length))), new BigNumber((_2081_v106).length))];
          let _2084_v109;
          let _2085_v110;
          let _2086_v111;
          let _2087_v112;
          let _out60;
          let _out61;
          let _out62;
          let _out63;
          let _outcollector27 = (_this).m7((_1954_v5).f17, (_1954_v5).f17, globalState);
          _out60 = _outcollector27[0];
          _out61 = _outcollector27[1];
          _out62 = _outcollector27[2];
          _out63 = _outcollector27[3];
          _2084_v109 = _out60;
          _2085_v110 = _out61;
          _2086_v111 = _out62;
          _2087_v112 = _out63;
          let _2088_v113;
          let _nw308 = Array((_dafny.ONE).toNumber()).fill([]);
          _2088_v113 = _nw308;
          let _2089_v114;
          _2089_v114 = _module.D17.create_DC50(_2088_v113);
          let _pat_let_tv59 = _2088_v113;
          let _2090_v115;
          _2090_v115 = _dafny.Seq.of(_2089_v114, _2089_v114, function (_pat_let42_0) {
            return function (_2091_dt__update__tmp_h3) {
              return function (_pat_let43_0) {
                return function (_2092_dt__update_hcf86_h0) {
                  return _module.D17.create_DC50(_2092_dt__update_hcf86_h0);
                }(_pat_let43_0);
              }(_pat_let_tv59);
            }(_pat_let42_0);
          }(_2089_v114));
          let _2093_v116;
          _2093_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wnd")).length),((!((_1954_v5).f17)) ? (_dafny.Seq.of(_2089_v114, _2089_v114, _2089_v114, _module.D17.create_DC50(_2088_v113))) : (_2090_v115)));
          let _index322 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1957_v8).length));
          (_1957_v8)[_index322] = (_1954_v5).f17;
          let _index323 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1957_v8).length));
          (_1957_v8)[_index323] = (_1954_v5).f17;
          let _2094_v117;
          _2094_v117 = _module.D3.create_DC9(_dafny.Set.fromElements(_dafny.Set.fromElements((_1954_v5).f17), _1959_v9, _1959_v9, _1959_v9));
          let _2095_v118;
          _2095_v118 = _module.D20.create_DC59(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(388),_2090_v115));
          let _2096_v119;
          let _nw309 = Array((new BigNumber(6)).toNumber()).fill(_module.D13.Default());
          _2096_v119 = _nw309;
          let _2097_v120;
          _2097_v120 = _dafny.Map.Empty.slice().updateUnsafe(_2084_v109,_2096_v119);
          let _index324 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1957_v8).length));
          let _index325 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1957_v8).length));
          let _rhs337 = (((_2095_v118).dtor_cf101).Merge(_2093_v116)).Merge(_2093_v116);
          let _rhs338 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_1954_v5).f17,_2096_v119)).Merge(_2097_v120)).length);
          let _rhs339 = _2085_v110;
          let _rhs340 = false;
          let _rhs341 = _2094_v117;
          let _lhs232 = globalState;
          let _lhs233 = _1957_v8;
          let _lhs234 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_1957_v8).length));
          let _lhs235 = _1957_v8;
          let _lhs236 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1957_v8).length));
          _2093_v116 = _rhs337;
          _lhs232.f1 = _rhs338;
          _lhs233[_lhs234] = _rhs339;
          _lhs235[_lhs236] = _rhs340;
          _2094_v117 = _rhs341;
        }
        let _index326 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1957_v8).length));
        (_1957_v8)[_index326] = (_1954_v5).f17;
        let _index327 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1957_v8).length));
        (_1957_v8)[_index327] = (((_1954_v5).f17) ? ((_1954_v5).f17) : (!((_this).f16)));
        (globalState).f1 = (_this).f15;
      }
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _module.D0.Default();
      let r3 = _dafny.ZERO;
      let _hi10 = (_this).f15;
      for (let _2098_i0 = ((_this).f15).plus((_this).f15); _2098_i0.isLessThan(_hi10); _2098_i0 = _2098_i0.plus(_dafny.ONE)) {
        let _2099_v0;
        _2099_v0 = _dafny.Map.Empty.slice().updateUnsafe(_2098_i0,(_module.__default.fm1(globalState)).isLessThanOrEqualTo((_this).f15));
        _2099_v0 = (_2099_v0).update((_this).f15, true);
        let _2100_v1;
        let _nw310 = new _module.C5();
        _nw310.__ctor();
        _2100_v1 = _nw310;
        _2100_v1 = _2100_v1;
        let _2101_v2;
        let _nw311 = Array((new BigNumber(28)).toNumber()).fill(false);
        _2101_v2 = _nw311;
        let _index328 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2101_v2).length));
        (_2101_v2)[_index328] = p0;
        let _2102_v3;
        let _init51 = function (_2103_i1) {
          return _module.D1.create_DC3((_this).f15);
        };
        let _nw312 = Array((new BigNumber(14)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw312.length); _i0_51++) {
          _nw312[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _2102_v3 = _nw312;
        let _2104_v4;
        _2104_v4 = _dafny.Seq.of(_2102_v3);
        let _2105_v5;
        _2105_v5 = _dafny.Seq.of((_this).f15);
        let _2106_v6;
        _2106_v6 = _module.D8.create_DC27(p1, _2104_v4, _2105_v5, _2098_i0);
        let _2107_v7;
        _2107_v7 = _dafny.Seq.of((_2106_v6).dtor_cf43, (_this).f16, p0);
        let _index329 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2101_v2).length));
        (_2101_v2)[_index329] = ((!(p1)) ? (p1) : ((_2107_v7)[_module.__default.safeIndex((_this).f15, new BigNumber((_2107_v7).length))]));
        let _index330 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2101_v2).length));
        (_2101_v2)[_index330] = (_2107_v7)[_module.__default.safeIndex((_this).f15, new BigNumber((_2107_v7).length))];
      }
      let _2108_v8;
      let _nw313 = Array((new BigNumber(5)).toNumber()).fill([]);
      _2108_v8 = _nw313;
      _2108_v8 = _2108_v8;
      let _2109_v9;
      _2109_v9 = _dafny.Set.fromElements((_this).f15);
      if ((_2109_v9).IsSubsetOf(_2109_v9)) {
        let _2110_v10;
        _2110_v10 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2111_v11;
        _2111_v11 = _dafny.Set.fromElements(p1, (_this).f16);
        let _2112_v13;
        _2112_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2110_v10);
        let _2113_v14;
        _2113_v14 = _module.D5.create_DC17((_this).f16, new _dafny.CodePoint('y'.codePointAt(0)));
        let _2114_v15;
        _2114_v15 = _dafny.Seq.of(_2113_v14, _module.D5.create_DC17(p0, _2110_v10));
        let _2115_v17;
        let _nw314 = Array((new BigNumber(15)).toNumber());
        _nw314[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2110_v10)).update(new BigNumber((_2111_v11).length), _2110_v10);
        _nw314[(_dafny.ONE).toNumber()] = function () {
          let _coll54 = new _dafny.Map();
          for (const _compr_54 of _dafny.IntegerRange(new BigNumber(145), new BigNumber(430))) {
            let _2116_v12 = _compr_54;
            if (((new BigNumber(145)).isLessThanOrEqualTo(_2116_v12)) && ((_2116_v12).isLessThan(new BigNumber(430)))) {
              _coll54.push([_module.__default.safeModuloInt(_2116_v12, (_this).f15),_2110_v10]);
            }
          }
          return _coll54;
        }();
        _nw314[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2110_v10);
        _nw314[(new BigNumber(3)).toNumber()] = _2112_v13;
        _nw314[(new BigNumber(4)).toNumber()] = _2112_v13;
        _nw314[(new BigNumber(5)).toNumber()] = (_2112_v13).update((_this).f15, _2110_v10);
        _nw314[(new BigNumber(6)).toNumber()] = _2112_v13;
        _nw314[(new BigNumber(7)).toNumber()] = _2112_v13;
        _nw314[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2110_v10);
        _nw314[(new BigNumber(9)).toNumber()] = _module.__default.fm50(p0, new BigNumber((_2114_v15).length), globalState);
        _nw314[(new BigNumber(10)).toNumber()] = _2112_v13;
        _nw314[(new BigNumber(11)).toNumber()] = function () {
          let _coll55 = new _dafny.Map();
          for (const _compr_55 of _dafny.IntegerRange(new BigNumber(12), new BigNumber(-997))) {
            let _2117_v16 = _compr_55;
            if (((new BigNumber(12)).isLessThanOrEqualTo(_2117_v16)) && ((_2117_v16).isLessThan(new BigNumber(-997)))) {
              _coll55.push([(_2117_v16).plus((_this).f15),_2110_v10]);
            }
          }
          return _coll55;
        }();
        _nw314[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2110_v10);
        _nw314[(new BigNumber(13)).toNumber()] = _2112_v13;
        _nw314[(new BigNumber(14)).toNumber()] = _module.__default.fm50((_module.D5.create_DC17((_this).f16, _2110_v10)).dtor_cf28, new BigNumber(909), globalState);
        _2115_v17 = _nw314;
        let _2118_v18;
        _2118_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_2115_v17);
        let _2119_v19;
        let _nw315 = Array((new BigNumber(15)).toNumber());
        _nw315[(_dafny.ZERO).toNumber()] = _2115_v17;
        _nw315[(_dafny.ONE).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(2)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(3)).toNumber()] = (((_this).f16) ? (_2115_v17) : (_2115_v17));
        _nw315[(new BigNumber(4)).toNumber()] = ((p0) ? (_2115_v17) : (_2115_v17));
        _nw315[(new BigNumber(5)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(6)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(7)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(8)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(9)).toNumber()] = (((_2118_v18).contains(p0)) ? ((_2118_v18).get(p0)) : (_2115_v17));
        _nw315[(new BigNumber(10)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(11)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(12)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(13)).toNumber()] = _2115_v17;
        _nw315[(new BigNumber(14)).toNumber()] = _2115_v17;
        _2119_v19 = _nw315;
        let _index331 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_2119_v19).length));
        (_2119_v19)[_index331] = _2115_v17;
        let _index332 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_2119_v19).length));
        (_2119_v19)[_index332] = _2115_v17;
        let _2120_v20;
        _2120_v20 = _module.D17.create_DC50(_2108_v8);
        let _2121_v21;
        _2121_v21 = _module.D17.create_DC52(_2120_v20);
        let _2122_v22;
        _2122_v22 = _module.D17.create_DC52(((p1) ? (_2120_v20) : (_2120_v20)));
        _2122_v22 = _2122_v22;
        let _2123_v23;
        let _nw316 = Array((new BigNumber(18)).toNumber());
        _nw316[(_dafny.ZERO).toNumber()] = _2111_v11;
        _nw316[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(p1);
        _nw316[(new BigNumber(2)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(3)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(4)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(5)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(6)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(7)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(8)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(9)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(10)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(11)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(12)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(13)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements((_this).f16);
        _nw316[(new BigNumber(15)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(16)).toNumber()] = _2111_v11;
        _nw316[(new BigNumber(17)).toNumber()] = _2111_v11;
        _2123_v23 = _nw316;
        let _2124_v24;
        let _nw317 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _2124_v24 = _nw317;
        let _2125_v25;
        _2125_v25 = _dafny.Set.fromElements(_2124_v24, _2124_v24, _2124_v24);
        let _2126_v26;
        _2126_v26 = _dafny.Seq.of(_2125_v25, _dafny.Set.fromElements(_2124_v24, _2124_v24), _2125_v25);
        let _2127_v27;
        _2127_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2123_v23,(_2126_v26)[_module.__default.safeIndex((_this).f15, new BigNumber((_2126_v26).length))]);
        r0 = (_2125_v25).IsProperSubsetOf((((_2127_v27).contains(_2123_v23)) ? ((_2127_v27).get(_2123_v23)) : (_2125_v25)));
        let _2128_v28;
        _2128_v28 = _dafny.Seq.UnicodeFromString("vtpb");
        let _2129_v29;
        _2129_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f16);
        let _2130_v30;
        _2130_v30 = _dafny.Set.fromElements(_2129_v29);
        (globalState).f7 = (_this).fm9(_2128_v28, _2130_v30, globalState);
        (globalState).f1 = (_this).f15;
      } else {
        let _2131_v31;
        _2131_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(479),(_this).f15);
        _2131_v31 = (_2131_v31).update((_this).f15, _module.__default.fm1(globalState));
        r1 = p1;
        r3 = (_this).f15;
        let _2132_v32;
        let _nw318 = Array((new BigNumber(27)).toNumber()).fill(_module.D1.Default());
        _2132_v32 = _nw318;
        let _2133_v33;
        _2133_v33 = _dafny.Seq.of(_2132_v32);
        let _2134_v34;
        _2134_v34 = _dafny.Seq.of((_this).f15, new BigNumber((_2109_v9).length), (_this).f15);
        let _2135_v35;
        _2135_v35 = _module.D8.create_DC27(p1, _2133_v33, _2134_v34, new BigNumber(726));
        let _2136_v36;
        _2136_v36 = _dafny.Seq.of((_this).f16, _module.__default.fm0((_this).f15, globalState), ((_2135_v35).dtor_cf43) && (p0), ((_this).f15).isEqualTo((_this).f15), p1);
        r0 = (_2136_v36)[_module.__default.safeIndex((_this).f15, new BigNumber((_2136_v36).length))];
        let _2137_v37;
        let _init52 = ((_2138_p1) => function (_2139_i2) {
          return _2138_p1;
        })(p1);
        let _nw319 = Array((new BigNumber(4)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw319.length); _i0_52++) {
          _nw319[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _2137_v37 = _nw319;
        let _2140_v38;
        _2140_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D0.create_DC0(_2137_v37)).dtor_cf0);
        _2140_v38 = (_2140_v38).update(p1, _2137_v37);
      }
      let _2141_v39;
      let _nw320 = new _module.C0();
      _nw320.__ctor();
      _2141_v39 = _nw320;
      r1 = (_this).f16;
      let _2142_v40;
      _2142_v40 = _dafny.Map.Empty.slice().updateUnsafe(!((_2141_v39).fm37(_dafny.Seq.UnicodeFromString("v"), (_this).f15, p1, (_this).f16, globalState)),_dafny.Seq.UnicodeFromString("icicg"));
      _2142_v40 = (_2142_v40).update((_this).f16, _dafny.Seq.UnicodeFromString("gfbbxc"));
      r0 = (_this).f16;
      let _2143_v41;
      _2143_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,!(!(p1)));
      r1 = (((_2143_v41).contains((_this).f15)) ? ((_2143_v41).get((_this).f15)) : (false));
      let _2144_v42;
      _2144_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
      let _2145_v43;
      _2145_v43 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2146_v44;
      _2146_v44 = _module.D0.create_DC1(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nwthpjsf"), _module.__default.safeIndex(new BigNumber((_2144_v42).length), new BigNumber((_dafny.Seq.UnicodeFromString("nwthpjsf")).length)), _2145_v43));
      r2 = _2146_v44;
      r3 = ((_this).f15).minus((_this).f15);
      return [r0, r1, r2, r3];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
      this._f10 = false;
      this._f13 = _dafny.ZERO;
      this._f14 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f13, f14, f10) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f10 = f10;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements((_this).f10, (_this).f10)).Difference(_dafny.Set.fromElements((_this).f10, (_this).f10))).IsDisjointFrom((_dafny.Set.fromElements((_this).f10)).Difference(_dafny.Set.fromElements(!(true))));
    };
    fm7(p0, globalState) {
      let _this = this;
      return (_this).f10;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      let _source36 = _module.D3.create_DC10((_this).f14, (_this).f10, (_this).f10);
      if (_source36.is_DC10) {
        let _2147___mcc_h0 = (_source36).cf14;
        let _2148___mcc_h1 = (_source36).cf15;
        let _2149___mcc_h2 = (_source36).cf16;
        let _2150_cf16 = _2149___mcc_h2;
        let _2151_cf15 = _2148___mcc_h1;
        let _2152_cf14 = _2147___mcc_h0;
        return (_this).f10;
      } else if (_source36.is_DC11) {
        let _2153___mcc_h3 = (_source36).cf17;
        let _2154___mcc_h4 = (_source36).cf18;
        let _2155_cf18 = _2154___mcc_h4;
        let _2156_cf17 = _2153___mcc_h3;
        return (_this).f10;
      } else if (_source36.is_DC9) {
        let _2157___mcc_h5 = (_source36).cf13;
        let _2158_cf13 = _2157___mcc_h5;
        return true;
      } else {
        let _2159___mcc_h6 = (_source36).cf19;
        let _2160_cf19 = _2159___mcc_h6;
        return (_this).f10;
      }
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      let _source37 = _module.D2.create_DC6(new _dafny.CodePoint('b'.codePointAt(0)));
      if (_source37.is_DC7) {
        let _2161___mcc_h0 = (_source37).cf9;
        let _2162___mcc_h1 = (_source37).cf10;
        let _2163___mcc_h2 = (_source37).cf11;
        let _2164_cf11 = _2163___mcc_h2;
        let _2165_cf10 = _2162___mcc_h1;
        let _2166_cf9 = _2161___mcc_h0;
        return (new BigNumber((_dafny.Seq.of((_this).f10)).length)).minus(_2164_cf11);
      } else if (_source37.is_DC6) {
        let _2167___mcc_h3 = (_source37).cf8;
        let _2168_cf8 = _2167___mcc_h3;
        return (_this).f13;
      } else {
        let _2169___mcc_h4 = (_source37).cf12;
        let _2170_cf12 = _2169___mcc_h4;
        return _module.__default.safeModuloInt(new BigNumber(627), (_this).f13);
      }
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Map.Empty;
      r0 = p1;
      let _2171_v0;
      let _init53 = function (_2172_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1((_this).f14),(_this).f14);
      };
      let _nw321 = Array((new BigNumber(14)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw321.length); _i0_53++) {
        _nw321[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _2171_v0 = _nw321;
      let _2173_v1;
      _2173_v1 = _module.D0.create_DC1((_this).f14);
      let _2174_v2;
      _2174_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2173_v1,(_this).f14);
      let _index333 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_2171_v0).length));
      (_2171_v0)[_index333] = (_2174_v2).Merge(_2174_v2);
      let _index334 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_2171_v0).length));
      (_2171_v0)[_index334] = _2174_v2;
      let _2175_v3;
      let _nw322 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _2175_v3 = _nw322;
      let _2176_v4;
      _2176_v4 = _module.D3.create_DC10((_this).f14, p1, (_this).f10);
      let _2177_v5;
      _2177_v5 = _module.D3.create_DC12(_2176_v4);
      let _2178_v6;
      _2178_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p1);
      let _2179_v7;
      _2179_v7 = _dafny.Set.fromElements(_2178_v6, _2178_v6);
      let _2180_v8;
      _2180_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm9((_this).f14, _2179_v7, globalState),(_this).f10);
      let _2181_v9;
      _2181_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC12(_2176_v4),(_2180_v8).update((_dafny.ZERO).minus((_this).f13), p1));
      let _2182_v10;
      _2182_v10 = _module.D3.create_DC12(_2177_v5);
      let _index335 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length));
      (_2175_v3)[_index335] = ((p1) ? (new BigNumber(((_2181_v9).update(_2182_v10, _2180_v8)).length)) : ((_this).f13));
      let _index336 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length));
      (_2175_v3)[_index336] = (_this).f13;
      let _hi11 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_this).f13).plus(new BigNumber(135))));
      for (let _2183_i1 = ((_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]).multipliedBy((_dafny.ZERO).minus((_this).f13)); _2183_i1.isLessThan(_hi11); _2183_i1 = _2183_i1.plus(_dafny.ONE)) {
        let _2184_v11;
        _2184_v11 = _dafny.Set.fromElements((_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]);
        _2184_v11 = ((_2184_v11).Difference(_dafny.Set.fromElements((_module.D2.create_DC7(_2183_i1, true, (_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))])).dtor_cf11, _dafny.ZERO, new BigNumber(241), _2183_i1, (_this).f13))).Difference(_2184_v11);
        let _2185_v12;
        let _nw323 = new _module.C2();
        _nw323.__ctor(p1, _2183_i1);
        _2185_v12 = _nw323;
        r0 = !(p1) || (p1);
        if ((_this).f10) {
          let _2186_v13;
          _2186_v13 = _dafny.MultiSet.fromElements((_2185_v12).f21);
          let _2187_v14;
          _2187_v14 = _module.D5.create_DC16(_2186_v13);
          let _2188_v15;
          let _nw324 = Array((new BigNumber(11)).toNumber());
          _nw324[(_dafny.ZERO).toNumber()] = _2187_v14;
          _nw324[(_dafny.ONE).toNumber()] = _2187_v14;
          _nw324[(new BigNumber(2)).toNumber()] = _2187_v14;
          _nw324[(new BigNumber(3)).toNumber()] = _module.D5.create_DC16(_2186_v13);
          _nw324[(new BigNumber(4)).toNumber()] = _module.D5.create_DC16(_2186_v13);
          _nw324[(new BigNumber(5)).toNumber()] = _2187_v14;
          _nw324[(new BigNumber(6)).toNumber()] = _2187_v14;
          _nw324[(new BigNumber(7)).toNumber()] = _2187_v14;
          _nw324[(new BigNumber(8)).toNumber()] = _module.D5.create_DC16(_2186_v13);
          _nw324[(new BigNumber(9)).toNumber()] = _2187_v14;
          _nw324[(new BigNumber(10)).toNumber()] = _2187_v14;
          _2188_v15 = _nw324;
          let _index337 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_2188_v15).length));
          (_2188_v15)[_index337] = _2187_v14;
          let _index338 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_2188_v15).length));
          (_2188_v15)[_index338] = _2187_v14;
          let _2189_v16;
          let _nw325 = Array((new BigNumber(9)).toNumber()).fill(_module.D16.Default());
          _2189_v16 = _nw325;
          let _2190_v17;
          let _nw326 = Array((new BigNumber(23)).toNumber()).fill(_module.D14.Default());
          _2190_v17 = _nw326;
          let _2191_v18;
          _2191_v18 = _module.D16.create_DC48(_2190_v17);
          let _index339 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_2189_v16).length));
          (_2189_v16)[_index339] = _2191_v18;
          let _index340 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_2189_v16).length));
          (_2189_v16)[_index340] = _2191_v18;
          let _2192_v19;
          _2192_v19 = _dafny.Seq.of((_2185_v12).f21, (((_2180_v8).contains((_2185_v12).f22)) ? ((_2180_v8).get((_2185_v12).f22)) : ((_2185_v12).f21)));
          let _2193_v20;
          _2193_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2192_v19,_2184_v11);
          _2193_v20 = (_2193_v20).update(_dafny.Seq.Concat(_dafny.Seq.of(p1, true, (_this).f10, (_2185_v12).f21), _dafny.Seq.of((_2185_v12).f21)), (_2184_v11).Intersect(_dafny.Set.fromElements(_2183_i1, (_this).f13, (_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))])));
          let _2194_v21;
          _2194_v21 = _module.D19.create_DC58((_2183_i1).minus((_this).f13), (_2185_v12).f21);
          _2194_v21 = _2194_v21;
          (globalState).f1 = (_this).f13;
        } else {
          let _2195_v22;
          _2195_v22 = _module.D19.create_DC57(_2184_v11);
          let _2196_v23;
          _2196_v23 = _dafny.Seq.of(_2195_v22);
          let _2197_v24;
          _2197_v24 = _module.D2.create_DC7((_this).f13, (_2185_v12).f21, new BigNumber(245));
          let _2198_v25;
          _2198_v25 = _module.D2.create_DC8(_2197_v24);
          let _2199_v26;
          _2199_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2196_v23).length),_2198_v25);
          let _2200_v27;
          _2200_v27 = new _dafny.CodePoint('r'.codePointAt(0));
          r1 = _dafny.Seq.update((_this).f14, _module.__default.safeIndex((new BigNumber((_2199_v26).length)).multipliedBy((_this).f13), new BigNumber(((_this).f14).length)), _2200_v27);
          let _2201_v28;
          let _init54 = function (_2202_i2) {
            return (_this).f10;
          };
          let _nw327 = Array((new BigNumber(22)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw327.length); _i0_54++) {
            _nw327[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _2201_v28 = _nw327;
          let _index341 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2201_v28).length));
          (_2201_v28)[_index341] = p1;
          let _index342 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2201_v28).length));
          (_2201_v28)[_index342] = p1;
          (globalState).f0 = _2183_i1;
          (globalState).f1 = new BigNumber((function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of (function () {
              let _coll57 = new _dafny.Map();
              for (const _compr_57 of (_2184_v11).Elements) {
                let _2203_v30 = _compr_57;
                if ((_2184_v11).contains(_2203_v30)) {
                  _coll57.push([(_2203_v30).plus((_2185_v12).f22),_dafny.Seq.of((_this).f10)]);
                }
              }
              return _coll57;
            }()).Keys.Elements) {
              let _2204_v29 = _compr_56;
              if ((function () {
                let _coll58 = new _dafny.Map();
                for (const _compr_58 of (_2184_v11).Elements) {
                  let _2203_v30 = _compr_58;
                  if ((_2184_v11).contains(_2203_v30)) {
                    _coll58.push([(_2203_v30).plus((_2185_v12).f22),_dafny.Seq.of((_this).f10)]);
                  }
                }
                return _coll58;
              }()).contains(_2204_v29)) {
                _coll56.push([_module.__default.safeModuloInt(_2204_v29, (_dafny.ZERO).minus(new BigNumber(((_this).f14).length))),new BigNumber(578)]);
              }
            }
            return _coll56;
          }()).length);
          r0 = !(p1) || (!(false));
        }
      }
      r0 = p1;
      let _hi12 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]), (_dafny.ZERO).minus(new BigNumber(((_this).f14).length)));
      for (let _2205_i3 = (_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]; _2205_i3.isLessThan(_hi12); _2205_i3 = _2205_i3.plus(_dafny.ONE)) {
        let _index343 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length));
        (_2175_v3)[_index343] = ((_this).f13).multipliedBy(_module.__default.safeDivisionInt(new BigNumber((_module.__default.fm41(globalState)).cardinality()), (_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]));
        let _2206_v31;
        let _init55 = ((_2207_p1) => function (_2208_i4) {
          return _dafny.Set.fromElements((_this).f10, _2207_p1);
        })(p1);
        let _nw328 = Array((new BigNumber(16)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw328.length); _i0_55++) {
          _nw328[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2206_v31 = _nw328;
        let _2209_v32;
        _2209_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2206_v31,_2205_i3);
        _2209_v32 = (_2209_v32).update(_2206_v31, (_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]);
        let _2210_v33;
        _2210_v33 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(p1), _dafny.Seq.update(_dafny.Seq.of((_this).f10, (_this).f10), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.of((_this).f10, (_this).f10)).length)), (_this).f10)));
        _2210_v33 = _2210_v33;
        _2180_v8 = (_2180_v8).update(_2205_i3, p1);
      }
      r0 = p1;
      r1 = _dafny.Seq.Concat((_this).f14, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_2211_i5) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), (_this).f14));
      let _2212_v34;
      _2212_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(212),(_2175_v3)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_2175_v3).length))]);
      r2 = _2212_v34;
      return [r0, r1, r2];
    }
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _hi13 = (((_this).f10) ? ((_this).f13) : ((_this).f13));
      for (let _2213_i0 = (_this).f13; _2213_i0.isLessThan(_hi13); _2213_i0 = _2213_i0.plus(_dafny.ONE)) {
        let _2214_v0;
        _2214_v0 = _dafny.Seq.of((_this).f13, _2213_i0);
        let _2215_v1;
        _2215_v1 = _dafny.Set.fromElements(new BigNumber((_2214_v0).length), (_this).f13);
        let _2216_v2;
        _2216_v2 = _dafny.Set.fromElements((_this).f10);
        let _2217_v3;
        _2217_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_2216_v2).length));
        let _2218_v4;
        let _nw329 = Array((new BigNumber(4)).toNumber()).fill(_module.D2.Default());
        _2218_v4 = _nw329;
        let _2219_v5;
        _2219_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f13),_2218_v4);
        let _2220_v6;
        _2220_v6 = _module.D6.create_DC19(_2219_v5);
        let _2221_v7;
        _2221_v7 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((_this).f13, new BigNumber((_2215_v1).length))).IsSubsetOf(_module.__default.fm18((_this).f13, _2217_v3, (_this).f10, globalState)),_2220_v6);
        _2221_v7 = (_2221_v7).update((_2213_i0).isLessThan((_this).f13), _2220_v6);
        let _2222_v8;
        let _init56 = ((_2223_i0) => function (_2224_i1) {
          return (_2224_i1).plus(_2223_i0);
        })(_2213_i0);
        let _nw330 = Array((new BigNumber(13)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw330.length); _i0_56++) {
          _nw330[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _2222_v8 = _nw330;
        let _index344 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2222_v8).length));
        (_2222_v8)[_index344] = _2213_i0;
        let _index345 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2222_v8).length));
        (_2222_v8)[_index345] = ((_this).f13).minus(_2213_i0);
        r1 = (_this).f10;
        let _2225_v9;
        _2225_v9 = _dafny.MultiSet.fromElements((_this).f13);
        let _2226_v10;
        let _nw331 = Array((new BigNumber(11)).toNumber());
        _nw331[(_dafny.ZERO).toNumber()] = false;
        _nw331[(_dafny.ONE).toNumber()] = (_this).f10;
        _nw331[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw331[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw331[(new BigNumber(4)).toNumber()] = !((_2213_i0).isLessThan(_2213_i0));
        _nw331[(new BigNumber(5)).toNumber()] = false;
        _nw331[(new BigNumber(6)).toNumber()] = (_2225_v9).IsSubsetOf(_2225_v9);
        _nw331[(new BigNumber(7)).toNumber()] = false;
        _nw331[(new BigNumber(8)).toNumber()] = (_this).f10;
        _nw331[(new BigNumber(9)).toNumber()] = (_this).f10;
        _nw331[(new BigNumber(10)).toNumber()] = (_this).f10;
        _2226_v10 = _nw331;
        let _index346 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2226_v10).length));
        (_2226_v10)[_index346] = (_this).f10;
        let _2227_v11;
        _2227_v11 = _module.D13.create_DC39((((_2217_v3).contains((_this).f10)) ? ((_2217_v3).get((_this).f10)) : ((_2222_v8)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_2222_v8).length))])), !((_this).f10));
        let _index347 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2226_v10).length));
        (_2226_v10)[_index347] = !(((((_this).f10) ? (_2227_v11) : (_2227_v11))).dtor_cf66);
      }
      let _2228_v12;
      let _nw332 = Array((new BigNumber(16)).toNumber()).fill([]);
      _2228_v12 = _nw332;
      let _2229_v13;
      let _nw333 = Array((new BigNumber(15)).toNumber());
      _nw333[(_dafny.ZERO).toNumber()] = (_this).f10;
      _nw333[(_dafny.ONE).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(2)).toNumber()] = true;
      _nw333[(new BigNumber(3)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(4)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(5)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(6)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(7)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(8)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(9)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(10)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(11)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(12)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(13)).toNumber()] = (_this).f10;
      _nw333[(new BigNumber(14)).toNumber()] = (_this).f10;
      _2229_v13 = _nw333;
      let _index348 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length));
      (_2228_v12)[_index348] = _2229_v13;
      let _index349 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length));
      (_2228_v12)[_index349] = _2229_v13;
      let _2230_v14;
      _2230_v14 = _module.D10.create_DC30(_dafny.Set.fromElements((_this).f10, false));
      let _source38 = _2230_v14;
      if (_source38.is_DC31) {
        let _2231___mcc_h0 = (_source38).cf50;
        let _2232___mcc_h1 = (_source38).cf51;
        let _2233___mcc_h2 = (_source38).cf52;
        let _2234_cf52 = _2233___mcc_h2;
        let _2235_cf51 = _2232___mcc_h1;
        let _2236_cf50 = _2231___mcc_h0;
        let _2237_v15;
        let _nw334 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _2237_v15 = _nw334;
        let _index350 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2237_v15).length));
        (_2237_v15)[_index350] = (_this).f13;
        let _index351 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2237_v15).length));
        (_2237_v15)[_index351] = _2234_cf52;
        let _2238_v16;
        _2238_v16 = _module.D0.create_DC0((_module.D0.create_DC2(false, (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))])).dtor_cf3);
        let _2239_v17;
        _2239_v17 = _dafny.Seq.of((_2238_v16).dtor_cf0, _2229_v13, _2229_v13);
        let _2240_v18;
        let _nw335 = new _module.C10();
        _nw335.__ctor(new BigNumber((_2239_v17).length), _2236_cf50);
        _2240_v18 = _nw335;
        (globalState).f7 = _module.__default.safeModuloInt((_2240_v18).f15, _module.__default.safeModuloInt((_2240_v18).f15, (_this).f13));
        if ((_2240_v18).f16) {
          let _2241_v19;
          let _nw336 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2241_v19 = _nw336;
          let _2242_v20;
          _2242_v20 = new _dafny.CodePoint('k'.codePointAt(0));
          let _index352 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_2241_v19).length));
          (_2241_v19)[_index352] = _2242_v20;
          let _index353 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_2241_v19).length));
          (_2241_v19)[_index353] = _2242_v20;
          let _2243_v21;
          let _nw337 = new _module.C5();
          _nw337.__ctor();
          _2243_v21 = _nw337;
          let _index354 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_2229_v13).length));
          (_2229_v13)[_index354] = true;
          let _index355 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_2229_v13).length));
          (_2229_v13)[_index355] = _2235_cf51;
          let _2244_v22;
          _2244_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2242_v20,(_2240_v18).f15);
          let _2245_v23;
          _2245_v23 = _dafny.Set.fromElements((_2240_v18).f16, _2235_cf51);
          _2244_v22 = (_2244_v22).update((_2241_v19)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_2241_v19).length))], new BigNumber((_2245_v23).length));
          let _index356 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_2241_v19).length));
          (_2241_v19)[_index356] = _2242_v20;
        } else {
          r1 = (_2240_v18).f16;
          let _2246_v24;
          let _init57 = ((_2247_cf50, _2248_v18) => function (_2249_i2) {
            return (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_2247_cf50,new _dafny.CodePoint('h'.codePointAt(0))),(_2248_v18).f16)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2248_v18).f16,_dafny.Map.Empty.slice().updateUnsafe((_2248_v18).f16,_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('c'.codePointAt(0))),(_2248_v18).f16))));
          })(_2236_cf50, _2240_v18);
          let _nw338 = Array((new BigNumber(12)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw338.length); _i0_57++) {
            _nw338[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2246_v24 = _nw338;
          let _2250_v25;
          _2250_v25 = _dafny.Set.fromElements((_this).f10, (_this).f10);
          let _2251_v26;
          _2251_v26 = new _dafny.CodePoint('p'.codePointAt(0));
          let _2252_v27;
          _2252_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2250_v25).length),_2251_v26);
          let _2253_v28;
          _2253_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(((_2252_v27).contains((_2240_v18).f15)) ? ((_2252_v27).get((_2240_v18).f15)) : (_2251_v26)));
          let _2254_v29;
          _2254_v29 = _dafny.Map.Empty.slice().updateUnsafe(_2253_v28,_2235_cf51);
          let _2255_v30;
          _2255_v30 = _2254_v29;
          let _2256_v31;
          _2256_v31 = _dafny.Map.Empty.slice().updateUnsafe((_2240_v18).f16,_2255_v30);
          let _2257_v32;
          _2257_v32 = _dafny.Map.Empty.slice().updateUnsafe((_2240_v18).f16,_2256_v31);
          let _index357 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_2246_v24).length));
          (_2246_v24)[_index357] = _2257_v32;
          let _2258_v33;
          _2258_v33 = _dafny.MultiSet.fromElements((_2237_v15)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_2237_v15).length))]);
          let _2259_v34;
          _2259_v34 = _dafny.MultiSet.fromElements(new BigNumber((_2258_v33).cardinality()), (_2240_v18).f15);
          let _2260_v35;
          _2260_v35 = _module.D4.create_DC15((_2240_v18).f15, _2258_v33, new _dafny.CodePoint('u'.codePointAt(0)), (_2240_v18).f15);
          let _2261_v36;
          let _nw339 = Array((new BigNumber(7)).toNumber());
          _nw339[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw339[(_dafny.ONE).toNumber()] = _2251_v26;
          _nw339[(new BigNumber(2)).toNumber()] = _2251_v26;
          _nw339[(new BigNumber(3)).toNumber()] = (_2260_v35).dtor_cf25;
          _nw339[(new BigNumber(4)).toNumber()] = _2251_v26;
          _nw339[(new BigNumber(5)).toNumber()] = _2251_v26;
          _nw339[(new BigNumber(6)).toNumber()] = ((true) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (_2251_v26));
          _2261_v36 = _nw339;
          let _2262_v37;
          _2262_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2237_v15,(_2240_v18).f16);
          let _index358 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_2246_v24).length));
          let _rhs342 = (((_2262_v37).contains(_2237_v15)) ? ((_2262_v37).get(_2237_v15)) : (_2235_cf51));
          let _rhs343 = (_this).fm8((_2240_v18).f15, (_2240_v18).f16, !(_2235_cf51) || (_2235_cf51), globalState);
          let _rhs344 = (_2257_v32).Merge(_2257_v32);
          let _rhs345 = _2261_v36;
          let _lhs237 = _2246_v24;
          let _lhs238 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_2246_v24).length));
          _2236_cf50 = _rhs342;
          _2236_cf50 = _rhs343;
          _lhs237[_lhs238] = _rhs344;
          _2261_v36 = _rhs345;
          (globalState).f1 = _2234_cf52;
          let _2263_v38;
          _2263_v38 = _dafny.Seq.of((_this).f14);
          let _2264_v39;
          _2264_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2236_cf50,_2236_cf50);
          let _2265_v40;
          _2265_v40 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_2235_cf51,_2235_cf51), _dafny.Map.Empty.slice().updateUnsafe(true,(_2240_v18).f16), _2264_v39, _dafny.Map.Empty.slice().updateUnsafe((_2240_v18).f16,_2235_cf51), _2264_v39);
          let _index359 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2237_v15).length));
          (_2237_v15)[_index359] = (_this).fm9((_2263_v38)[_module.__default.safeIndex(_2234_cf52, new BigNumber((_2263_v38).length))], (_2265_v40).Intersect(_dafny.Set.fromElements(_2264_v39, (_2264_v39).update(_2236_cf50, (_2240_v18).f16), _dafny.Map.Empty.slice().updateUnsafe(_2235_cf51,(_this).f10), _2264_v39)), globalState);
          let _arr6 = (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))];
          let _index360 = _module.__default.safeIndex(new BigNumber(590), new BigNumber(((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))]).length));
          _arr6[_index360] = _2235_cf51;
          let _2266_v41;
          let _nw340 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _2266_v41 = _nw340;
          let _2267_v42;
          _2267_v42 = _dafny.Seq.of((_2240_v18).f16);
          let _index361 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_2266_v41).length));
          (_2266_v41)[_index361] = _dafny.Seq.Concat(_2267_v42, _2267_v42);
          let _arr7 = (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))];
          let _index362 = _module.__default.safeIndex(new BigNumber(220), new BigNumber(((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))]).length));
          _arr7[_index362] = _2235_cf51;
          let _2268_v43;
          _2268_v43 = _module.D11.create_DC34((_2240_v18).f15, true, _2234_cf52);
          let _2269_v44;
          _2269_v44 = _dafny.Map.Empty.slice().updateUnsafe((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))],_2236_cf50);
          let _2270_v45;
          let _nw341 = new _module.C2();
          _nw341.__ctor(_2235_cf51, (_2240_v18).f15);
          _2270_v45 = _nw341;
          let _2271_v46;
          _2271_v46 = _dafny.Seq.of(_2270_v45);
          let _pat_let_tv60 = _2228_v12;
          let _pat_let_tv61 = _2228_v12;
          let _pat_let_tv62 = _2229_v13;
          let _pat_let_tv63 = _2240_v18;
          let _arr8 = (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))];
          let _index363 = _module.__default.safeIndex(new BigNumber(590), new BigNumber(((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))]).length));
          let _index364 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_2266_v41).length));
          let _arr9 = (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))];
          let _index365 = _module.__default.safeIndex(new BigNumber(220), new BigNumber(((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))]).length));
          let _rhs346 = (_this).f10;
          let _rhs347 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2267_v42, _2267_v42), _dafny.Seq.of((function (_pat_let44_0) {
            return function (_2274_dt__update__tmp_h1) {
              return function (_pat_let47_0) {
                return function (_2275_dt__update_hcf57_h0) {
                  return function (_pat_let48_0) {
                    return function (_2276_dt__update_hcf56_h0) {
                      return _module.D11.create_DC34(_2276_dt__update_hcf56_h0, _2275_dt__update_hcf57_h0, (_2274_dt__update__tmp_h1).dtor_cf58);
                    }(_pat_let48_0);
                  }((_this).f13);
                }(_pat_let47_0);
              }((_pat_let_tv63).f16);
            }(_pat_let44_0);
          }(function (_pat_let45_0) {
            return function (_2272_dt__update__tmp_h0) {
              return function (_pat_let46_0) {
                return function (_2273_dt__update_hcf58_h0) {
                  return _module.D11.create_DC34((_2272_dt__update__tmp_h0).dtor_cf56, (_2272_dt__update__tmp_h0).dtor_cf57, _2273_dt__update_hcf58_h0);
                }(_pat_let46_0);
              }(new BigNumber((_dafny.Set.fromElements((_pat_let_tv61)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_pat_let_tv60).length))], _pat_let_tv62)).length));
            }(_pat_let45_0);
          }(_2268_v43))).dtor_cf57, (((_2269_v44).contains(_2229_v13)) ? ((_2269_v44).get(_2229_v13)) : ((_this).f10))));
          let _rhs348 = (_2240_v18).f16;
          let _rhs349 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2271_v46, _dafny.Seq.of(_2270_v45)), _2271_v46)).length);
          let _rhs350 = false;
          let _lhs239 = (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))];
          let _lhs240 = _module.__default.safeIndex(new BigNumber(590), new BigNumber(((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))]).length));
          let _lhs241 = _2266_v41;
          let _lhs242 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_2266_v41).length));
          let _lhs243 = (_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))];
          let _lhs244 = _module.__default.safeIndex(new BigNumber(220), new BigNumber(((_2228_v12)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_2228_v12).length))]).length));
          let _lhs245 = globalState;
          _lhs239[_lhs240] = _rhs346;
          _lhs241[_lhs242] = _rhs347;
          _lhs243[_lhs244] = _rhs348;
          _lhs245.f1 = _rhs349;
          r1 = _rhs350;
        }
      } else {
        let _2277___mcc_h3 = (_source38).cf49;
        let _2278_cf49 = _2277___mcc_h3;
        let _2279_v47;
        _2279_v47 = _dafny.MultiSet.fromElements((_this).f10);
        let _2280_v48;
        _2280_v48 = _module.D5.create_DC16(_2279_v47);
        let _source39 = _2280_v48;
        if (_source39.is_DC17) {
          let _2281___mcc_h4 = (_source39).cf28;
          let _2282___mcc_h5 = (_source39).cf29;
          let _2283_cf29 = _2282___mcc_h5;
          let _2284_cf28 = _2281___mcc_h4;
          let _index366 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2229_v13).length));
          (_2229_v13)[_index366] = _2284_cf28;
          let _index367 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2229_v13).length));
          let _rhs351 = (_this).f10;
          let _rhs352 = _2283_cf29;
          let _rhs353 = (_this).f13;
          let _lhs246 = _2229_v13;
          let _lhs247 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2229_v13).length));
          _lhs246[_lhs247] = _rhs351;
          _2283_cf29 = _rhs352;
          r2 = _rhs353;
          let _2285_v49;
          let _nw342 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
          _2285_v49 = _nw342;
          let _2286_v50;
          _2286_v50 = _dafny.Set.fromElements((_this).f13);
          let _index368 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2285_v49).length));
          (_2285_v49)[_index368] = _2286_v50;
          let _index369 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2285_v49).length));
          (_2285_v49)[_index369] = (_dafny.Set.fromElements((_this).f13, (_this).f13)).Union((_2286_v50).Difference(_2286_v50));
          let _2287_v51;
          let _nw343 = Array((new BigNumber(12)).toNumber()).fill([]);
          _2287_v51 = _nw343;
          let _2288_v52;
          let _nw344 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _2288_v52 = _nw344;
          let _index370 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_2287_v51).length));
          (_2287_v51)[_index370] = _2288_v52;
          let _index371 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_2287_v51).length));
          (_2287_v51)[_index371] = _2288_v52;
          r1 = !((_2229_v13)[_module.__default.safeIndex(new BigNumber(170), new BigNumber((_2229_v13).length))]);
        } else if (_source39.is_DC16) {
          let _2289___mcc_h6 = (_source39).cf27;
          let _2290_cf27 = _2289___mcc_h6;
          let _2291_v53;
          _2291_v53 = _dafny.Seq.of(new BigNumber(-535));
          let _2292_v55;
          _2292_v55 = _dafny.MultiSet.fromElements((((_2290_cf27).contains((_this).f10)) ? ((_2290_cf27).get((_this).f10)) : ((_this).f13)), new BigNumber((function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_2291_v53).Elements) {
              let _2293_v54 = _compr_59;
              if (_dafny.Seq.contains(_2291_v53, _2293_v54)) {
                _coll59.push([_module.__default.safeDivisionInt(_2293_v54, (_this).f13),(_this).f13]);
              }
            }
            return _coll59;
          }()).length));
          let _2294_v56;
          _2294_v56 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f10);
          let _2295_v57;
          let _nw345 = Array((new BigNumber(29)).toNumber());
          _nw345[(_dafny.ZERO).toNumber()] = (_this).f13;
          _nw345[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(((_this).f13).multipliedBy((_this).f13));
          _nw345[(new BigNumber(2)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(3)).toNumber()] = new BigNumber(((_this).f14).length);
          _nw345[(new BigNumber(4)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(5)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus((_2291_v53)[_module.__default.safeIndex((_this).f13, new BigNumber((_2291_v53).length))])).plus((_this).f13);
          _nw345[(new BigNumber(7)).toNumber()] = new BigNumber(658);
          _nw345[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt((_this).f13, (_this).f13);
          _nw345[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-834));
          _nw345[(new BigNumber(10)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(11)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(12)).toNumber()] = new BigNumber(((_2292_v55).Intersect(_2292_v55)).cardinality());
          _nw345[(new BigNumber(13)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(14)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(15)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)).Merge(_2294_v56)).length);
          _nw345[(new BigNumber(16)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(17)).toNumber()] = ((_this).f13).minus(_module.__default.fm1(globalState));
          _nw345[(new BigNumber(18)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt((_this).f13, new BigNumber(((_this).f14).length));
          _nw345[(new BigNumber(20)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("ffvvncjhg")).length)).plus((_this).f13));
          _nw345[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm3((_dafny.ZERO).minus((_this).f13), (_this).f13, globalState)).length));
          _nw345[(new BigNumber(23)).toNumber()] = _module.__default.safeDivisionInt((_this).f13, new BigNumber(626));
          _nw345[(new BigNumber(24)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(25)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(26)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(27)).toNumber()] = (_this).f13;
          _nw345[(new BigNumber(28)).toNumber()] = (_this).f13;
          _2295_v57 = _nw345;
          let _index372 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_2295_v57).length));
          (_2295_v57)[_index372] = new BigNumber((_module.__default.fm41(globalState)).cardinality());
          let _index373 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_2295_v57).length));
          (_2295_v57)[_index373] = (_this).f13;
          let _2296_v58;
          let _nw346 = new _module.C6();
          _nw346.__ctor();
          _2296_v58 = _nw346;
          r1 = (_this).f10;
          let _2297_v59;
          _2297_v59 = _module.D19.create_DC58((_this).f13, (_this).f10);
          let _2298_v60;
          let _nw347 = new _module.C9();
          _nw347.__ctor((_this).f10, (_2297_v59).dtor_cf100);
          _2298_v60 = _nw347;
          let _2299_v61;
          _2299_v61 = _dafny.Seq.of(_2298_v60, _2298_v60);
          r1 = _dafny.areEqual((((_this).f10) ? (_2299_v61) : (_dafny.Seq.update(_2299_v61, _module.__default.safeIndex((_2295_v57)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((_2295_v57).length))], new BigNumber((_2299_v61).length)), _2298_v60))), _2299_v61);
        } else {
          let _2300___mcc_h7 = (_source39).cf30;
          let _2301_cf30 = _2300___mcc_h7;
          let _2302_v62;
          _2302_v62 = _dafny.Seq.of((_this).f10, (_this).f10);
          let _2303_v63;
          _2303_v63 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2302_v62, _2302_v62),(new BigNumber(668)).minus((_dafny.ZERO).minus((_this).f13)));
          _2303_v63 = (_2303_v63).update(_2302_v62, (_dafny.ZERO).minus((_this).f13));
          let _2304_v64;
          _2304_v64 = new _dafny.CodePoint('k'.codePointAt(0));
          let _2305_v65;
          _2305_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10);
          let _2306_v66;
          _2306_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2304_v64,_2305_v65);
          _2306_v66 = (_2306_v66).Merge(_2306_v66);
          let _2307_v67;
          _2307_v67 = _dafny.Seq.UnicodeFromString("rmkyku");
          _2307_v67 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-662)), ((_2308_v64) => function (_2309_i3) {
            return _2308_v64;
          })(_2304_v64));
          (globalState).f7 = ((_this).f13).plus(new BigNumber((_dafny.Seq.UnicodeFromString("hpyfyh")).length));
        }
        let _2310_v68;
        let _nw348 = new _module.C7();
        _nw348.__ctor((_this).f10, ((!((_this).f10)) ? ((_this).f10) : ((_this).f10)));
        _2310_v68 = _nw348;
        let _2311_v69;
        let _nw349 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _2311_v69 = _nw349;
        let _index374 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_2311_v69).length));
        (_2311_v69)[_index374] = new BigNumber((_2278_cf49).length);
        let _2312_v70;
        let _nw350 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2312_v70 = _nw350;
        let _2313_v71;
        _2313_v71 = new _dafny.CodePoint('t'.codePointAt(0));
        let _index375 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_2312_v70).length));
        (_2312_v70)[_index375] = _2313_v71;
        let _index376 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_2311_v69).length));
        let _index377 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_2312_v70).length));
        let _rhs354 = _2279_v47;
        let _rhs355 = (_2278_cf49).IsProperSubsetOf((_2278_cf49).Union(_2278_cf49));
        let _rhs356 = ((_this).f13).minus((_this).f13);
        let _rhs357 = _2313_v71;
        let _lhs248 = _2311_v69;
        let _lhs249 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_2311_v69).length));
        let _lhs250 = _2312_v70;
        let _lhs251 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_2312_v70).length));
        _2279_v47 = _rhs354;
        r1 = _rhs355;
        _lhs248[_lhs249] = _rhs356;
        _lhs250[_lhs251] = _rhs357;
        let _2314_v72;
        _2314_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
        _2314_v72 = (_2314_v72).update(new BigNumber((_module.__default.fm18((_this).f13, _dafny.Map.Empty.slice().updateUnsafe((_2310_v68).f19,(_this).f13), (_2310_v68).f19, globalState)).cardinality()), (_this).f13);
      }
      let _2315_v73;
      let _nw351 = new _module.C9();
      _nw351.__ctor((_this).f10, (_this).f10);
      _2315_v73 = _nw351;
      let _2316_v74;
      _2316_v74 = _dafny.Set.fromElements(_2315_v73, _2315_v73);
      let _hi14 = new BigNumber((_2316_v74).length);
      for (let _2317_i4 = (_this).f13; _2317_i4.isLessThan(_hi14); _2317_i4 = _2317_i4.plus(_dafny.ONE)) {
        let _2318_v75;
        _2318_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2317_i4,((_this).f13).plus((_this).f13));
        _2318_v75 = (_2318_v75).update(_2317_i4, new BigNumber(((_2318_v75).Merge(_2318_v75)).length));
        let _2319_v76;
        _2319_v76 = _dafny.MultiSet.fromElements(!((_this).f10), (_2315_v73).f17, (_2315_v73).f17);
        let _2320_v77;
        _2320_v77 = _module.D11.create_DC34(new BigNumber((_2319_v76).cardinality()), (_this).f10, new BigNumber(((_this).f14).length));
        let _2321_v78;
        _2321_v78 = _dafny.Map.Empty.slice().updateUnsafe((_2315_v73).f17,_module.__default.safeModuloInt((_2320_v77).dtor_cf56, _2317_i4));
        let _2322_v79;
        _2322_v79 = _dafny.Seq.UnicodeFromString("gijrn");
        let _2323_v80;
        _2323_v80 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2324_v81;
        _2324_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2323_v80,_module.__default.safeDivisionInt((_this).f13, _2317_i4));
        let _2325_v82;
        _2325_v82 = _dafny.Seq.of((_this).f13);
        let _2326_v83;
        _2326_v83 = _dafny.Seq.of((_2315_v73).f17, false);
        let _rhs358 = (((_2321_v78).update((_this).f10, (_this).f13)).update((_2315_v73).f17, _2317_i4)).update(!((_2315_v73).f17) || ((_2315_v73).f17), (_2325_v82)[_module.__default.safeIndex((_this).f13, new BigNumber((_2325_v82).length))]);
        let _rhs359 = _dafny.Seq.Concat((_this).f14, (_this).f14);
        let _rhs360 = ((_2324_v81).update(_module.__default.fm4((_2315_v73).f17, (_2315_v73).f17, (_2326_v83)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_2326_v83).length))], globalState), _2317_i4)).update(new _dafny.CodePoint('i'.codePointAt(0)), (_dafny.ZERO).minus(_2317_i4));
        _2321_v78 = _rhs358;
        _2322_v79 = _rhs359;
        _2324_v81 = _rhs360;
        let _2327_v84;
        let _nw352 = Array((new BigNumber(11)).toNumber()).fill([]);
        _2327_v84 = _nw352;
        _2327_v84 = _2327_v84;
        _2326_v83 = _dafny.Seq.Concat(_2326_v83, (_module.D18.create_DC56(_2326_v83)).dtor_cf97);
      }
      (globalState).f0 = _module.__default.safeDivisionInt((_this).f13, (_this).f13);
      let _2328_v85;
      _2328_v85 = _dafny.Map.Empty.slice().updateUnsafe(!((_2315_v73).f17),_2229_v13);
      _2328_v85 = (_2328_v85).Merge(_2328_v85);
      let _2329_v86;
      _2329_v86 = _dafny.Set.fromElements((_this).f14);
      r0 = (_2329_v86).Difference((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("k"))).Intersect(_2329_v86));
      r1 = (new BigNumber(720)).isLessThanOrEqualTo((_this).f13);
      r2 = _module.__default.safeModuloInt((_this).f13, (_dafny.ZERO).minus((_this).f13));
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      let _2330_v0;
      _2330_v0 = false;
      let _2331_v1;
      _2331_v1 = _dafny.Seq.of(new BigNumber(((_this).f14).length));
      let _2332_v2;
      let _nw353 = Array((new BigNumber(3)).toNumber());
      _nw353[(_dafny.ZERO).toNumber()] = _2331_v1;
      _nw353[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2331_v1, _2331_v1);
      _nw353[(new BigNumber(2)).toNumber()] = _2331_v1;
      _2332_v2 = _nw353;
      let _2333_v4;
      _2333_v4 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(517)), function (_2334_i0) {
        return new BigNumber((function () {
          let _coll60 = new _dafny.Map();
          for (const _compr_60 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13)).Keys.Elements) {
            let _2335_v3 = _compr_60;
            if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13)).contains(_2335_v3)) {
              _coll60.push([(_2335_v3).multipliedBy(new BigNumber(((_this).f14).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(239)), function (_2336_i1) {
                return new _dafny.CodePoint('n'.codePointAt(0));
              })]);
            }
          }
          return _coll60;
        }()).length);
      }), _dafny.Seq.of((_this).f13));
      let _2337_v5;
      _2337_v5 = _dafny.Seq.of((_this).fm7((_this).f10, globalState), _2330_v0);
      let _index378 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_2332_v2).length));
      (_2332_v2)[_index378] = (_2333_v4)[_module.__default.safeIndex(new BigNumber((_2337_v5).length), new BigNumber((_2333_v4).length))];
      let _2338_v6;
      _2338_v6 = _module.D16.create_DC49((_module.D20.create_DC60(new BigNumber(627))).dtor_cf102);
      let _pat_let_tv64 = p0;
      let _2339_v7;
      let _nw354 = Array((new BigNumber(27)).toNumber());
      _nw354[(_dafny.ZERO).toNumber()] = _2338_v6;
      _nw354[(_dafny.ONE).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(2)).toNumber()] = function (_pat_let49_0) {
        return function (_2340_dt__update__tmp_h0) {
          return function (_pat_let50_0) {
            return function (_2341_dt__update_hcf85_h0) {
              return _module.D16.create_DC49(_2341_dt__update_hcf85_h0);
            }(_pat_let50_0);
          }((_dafny.ZERO).minus((_this).f13));
        }(_pat_let49_0);
      }(_2338_v6);
      _nw354[(new BigNumber(3)).toNumber()] = ((_2330_v0) ? (function (_pat_let51_0) {
        return function (_2342_dt__update__tmp_h1) {
          return function (_pat_let52_0) {
            return function (_2343_dt__update_hcf85_h1) {
              return _module.D16.create_DC49(_2343_dt__update_hcf85_h1);
            }(_pat_let52_0);
          }((_this).f13);
        }(_pat_let51_0);
      }(_2338_v6)) : (_2338_v6));
      _nw354[(new BigNumber(4)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(5)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(6)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(7)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(8)).toNumber()] = ((!(_module.__default.fm0(p0, globalState))) ? (_2338_v6) : (_2338_v6));
      _nw354[(new BigNumber(9)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(10)).toNumber()] = _module.__default.fm56(globalState);
      _nw354[(new BigNumber(11)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(12)).toNumber()] = _module.D16.create_DC49((_this).f13);
      _nw354[(new BigNumber(13)).toNumber()] = _module.D16.create_DC49(new BigNumber(711));
      _nw354[(new BigNumber(14)).toNumber()] = _module.__default.fm56(globalState);
      _nw354[(new BigNumber(15)).toNumber()] = _module.D16.create_DC49(p0);
      _nw354[(new BigNumber(16)).toNumber()] = function (_pat_let53_0) {
        return function (_2344_dt__update__tmp_h2) {
          return function (_pat_let54_0) {
            return function (_2345_dt__update_hcf85_h2) {
              return _module.D16.create_DC49(_2345_dt__update_hcf85_h2);
            }(_pat_let54_0);
          }((_dafny.ZERO).minus(_pat_let_tv64));
        }(_pat_let53_0);
      }(_2338_v6);
      _nw354[(new BigNumber(17)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(18)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(19)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(20)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(21)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(22)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(23)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(24)).toNumber()] = _2338_v6;
      _nw354[(new BigNumber(25)).toNumber()] = _module.D16.create_DC49(new BigNumber((_module.__default.fm13((_this).f10, _2330_v0, globalState)).length));
      _nw354[(new BigNumber(26)).toNumber()] = _2338_v6;
      _2339_v7 = _nw354;
      let _index379 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2339_v7).length));
      (_2339_v7)[_index379] = _2338_v6;
      let _2346_v8;
      _2346_v8 = _dafny.Set.fromElements((_this).f13, p0);
      let _pat_let_tv65 = p0;
      let _index380 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_2332_v2).length));
      let _index381 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2339_v7).length));
      let _rhs361 = !(_2346_v8).equals((_2346_v8).Union(_2346_v8));
      let _rhs362 = _dafny.Seq.Concat(_2331_v1, _2331_v1);
      let _rhs363 = function (_pat_let55_0) {
        return function (_2347_dt__update__tmp_h3) {
          return function (_pat_let56_0) {
            return function (_2348_dt__update_hcf85_h3) {
              return _module.D16.create_DC49(_2348_dt__update_hcf85_h3);
            }(_pat_let56_0);
          }(_pat_let_tv65);
        }(_pat_let55_0);
      }(_2338_v6);
      let _lhs252 = _2332_v2;
      let _lhs253 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_2332_v2).length));
      let _lhs254 = _2339_v7;
      let _lhs255 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_2339_v7).length));
      _2330_v0 = _rhs361;
      _lhs252[_lhs253] = _rhs362;
      _lhs254[_lhs255] = _rhs363;
      let _2349_v9;
      _2349_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4((_this).f10, _2330_v0, (_this).f10, globalState),_module.__default.safeDivisionInt((_2331_v1)[_module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_2331_v1).length))], new BigNumber(916)));
      let _2350_v10;
      _2350_v10 = new _dafny.CodePoint('x'.codePointAt(0));
      _2349_v9 = (_2349_v9).update(_2350_v10, p0);
      _2349_v9 = (_2349_v9).update(_2350_v10, new BigNumber(((_this).f14).length));
      let _2351_v11;
      _2351_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2350_v10,_2337_v5);
      _2351_v11 = (_2351_v11).update(_2350_v10, _2337_v5);
      let _hi15 = _module.__default.fm1(globalState);
      for (let _2352_i2 = p0; _2352_i2.isLessThan(_hi15); _2352_i2 = _2352_i2.plus(_dafny.ONE)) {
        if ((_this).f10) {
          let _2353_v12;
          _2353_v12 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("hqev"));
          _2353_v12 = _2353_v12;
          let _2354_v13;
          _2354_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2330_v0,_2350_v10);
          _2354_v13 = (_2354_v13).update(_2330_v0, _2350_v10);
          let _2355_v14;
          _2355_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2330_v0,p0);
          let _2356_v15;
          _2356_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p0),_2355_v14);
          _2356_v15 = _2356_v15;
          let _2357_v16;
          _2357_v16 = _dafny.Set.fromElements(!((_2337_v5)[_module.__default.safeIndex(new BigNumber(-719), new BigNumber((_2337_v5).length))]), _2330_v0, _2330_v0);
          let _2358_v17;
          _2358_v17 = _module.D10.create_DC31(_module.__default.fm0((_this).f13, globalState), (_2357_v16).equals(_2357_v16), ((_this).f13).plus(_2352_i2));
          let _pat_let_tv66 = p0;
          _2358_v17 = function (_pat_let57_0) {
            return function (_2359_dt__update__tmp_h4) {
              return function (_pat_let58_0) {
                return function (_2360_dt__update_hcf52_h0) {
                  return _module.D10.create_DC31((_2359_dt__update__tmp_h4).dtor_cf50, (_2359_dt__update__tmp_h4).dtor_cf51, _2360_dt__update_hcf52_h0);
                }(_pat_let58_0);
              }(_pat_let_tv66);
            }(_pat_let57_0);
          }(_2358_v17);
          let _2361_v18;
          let _nw355 = new _module.C0();
          _nw355.__ctor();
          _2361_v18 = _nw355;
        } else {
          (globalState).f1 = _module.__default.fm1(globalState);
          let _2362_v19;
          _2362_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2352_i2,_2352_i2);
          let _2363_v20;
          let _nw356 = new _module.C4();
          _nw356.__ctor();
          _2363_v20 = _nw356;
          let _2364_v21;
          _2364_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2363_v20);
          let _2365_v22;
          _2365_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.update((_this).f14, _module.__default.safeIndex((((_2362_v19).contains((_this).f13)) ? ((_2362_v19).get((_this).f13)) : (new BigNumber((_module.__default.fm26(globalState)).length))), new BigNumber(((_this).f14).length)), _2350_v10)).length), _2352_i2)).length),new BigNumber(((_2364_v21).Merge(_2364_v21)).length));
          let _2366_v23;
          _2366_v23 = _dafny.MultiSet.fromElements(_2350_v10);
          let _2367_v24;
          _2367_v24 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13, new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_2363_v20).fm8((_this).f13, (_this).f10, (_this).f10, globalState), _2330_v0, (_this).f10, (_this).f10, (_this).f10), _module.__default.safeIndex((((_2349_v9).contains(_2350_v10)) ? ((_2349_v9).get(_2350_v10)) : ((((_2366_v23).contains(_2350_v10)) ? ((_2366_v23).get(_2350_v10)) : (new BigNumber(804))))), new BigNumber((_dafny.Seq.of((_2363_v20).fm8((_this).f13, (_this).f10, (_this).f10, globalState), _2330_v0, (_this).f10, (_this).f10, (_this).f10)).length)), _2330_v0)).length), p0, _2352_i2);
          _2362_v19 = (_2362_v19).update((p0).minus((_dafny.ZERO).minus(new BigNumber((_2367_v24).cardinality()))), _2352_i2);
          let _2368_v25;
          let _nw357 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _2368_v25 = _nw357;
          _2368_v25 = _2368_v25;
          let _2369_v26;
          _2369_v26 = _dafny.Seq.of(_2363_v20, _2363_v20, _2363_v20, _2363_v20, _2363_v20);
          let _2370_v27;
          let _nw358 = Array((_dafny.ONE).toNumber());
          _nw358[(_dafny.ZERO).toNumber()] = (_2369_v26)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_2369_v26).length))];
          _2370_v27 = _nw358;
          _2370_v27 = _2370_v27;
          let _2371_v28;
          let _nw359 = Array((new BigNumber(6)).toNumber());
          _nw359[(_dafny.ZERO).toNumber()] = _2330_v0;
          _nw359[(_dafny.ONE).toNumber()] = _2330_v0;
          _nw359[(new BigNumber(2)).toNumber()] = _2330_v0;
          _nw359[(new BigNumber(3)).toNumber()] = !((_this).f10);
          _nw359[(new BigNumber(4)).toNumber()] = (_this).f10;
          _nw359[(new BigNumber(5)).toNumber()] = _2330_v0;
          _2371_v28 = _nw359;
          let _2372_v29;
          _2372_v29 = _dafny.Seq.of(_2371_v28);
          let _2373_v30;
          _2373_v30 = _module.D0.create_DC0((_2372_v29)[_module.__default.safeIndex(_2352_i2, new BigNumber((_2372_v29).length))]);
          _2373_v30 = _2373_v30;
        }
        let _2374_v31;
        _2374_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2330_v0,_2350_v10);
        let _2375_v32;
        _2375_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2350_v10,true);
        let _2376_v33;
        _2376_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2374_v31,(((_2375_v32).contains(_2350_v10)) ? ((_2375_v32).get(_2350_v10)) : (_2330_v0)));
        let _2377_v34;
        let _nw360 = Array((new BigNumber(5)).toNumber());
        _nw360[(_dafny.ZERO).toNumber()] = !(_2330_v0) || (false);
        _nw360[(_dafny.ONE).toNumber()] = !(_2330_v0);
        _nw360[(new BigNumber(2)).toNumber()] = _2330_v0;
        _nw360[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw360[(new BigNumber(4)).toNumber()] = !((new BigNumber(((_2376_v33).update((_2374_v31).update((_this).f10, _2350_v10), true)).length)).isLessThan(new BigNumber((_2331_v1).length)));
        _2377_v34 = _nw360;
        let _index382 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_2377_v34).length));
        (_2377_v34)[_index382] = _2330_v0;
        let _index383 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_2377_v34).length));
        (_2377_v34)[_index383] = _2330_v0;
        _2330_v0 = _2330_v0;
        let _2378_v35;
        let _2379_v36;
        let _out64;
        let _out65;
        let _outcollector28 = _module.__default.m0((_this).f13, _2377_v34, globalState);
        _out64 = _outcollector28[0];
        _out65 = _outcollector28[1];
        _2378_v35 = _out64;
        _2379_v36 = _out65;
      }
      let _2380_v37;
      _2380_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2330_v0,_2330_v0);
      let _2381_v38;
      _2381_v38 = _dafny.Set.fromElements(_2380_v37);
      let _2382_v39;
      _2382_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2350_v10,(_this).f10);
      let _2383_v40;
      _2383_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2382_v39,(_this).f10);
      let _2384_v41;
      _2384_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2330_v0,(_this).f13);
      let _2385_v42;
      _2385_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm9((_this).f14, _2381_v38, globalState),(_module.__default.fm18((_this).f13, _2384_v41, true, globalState)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f13, new BigNumber(((_2383_v40).update(_2382_v39, false)).length), (_this).f13)));
      let _2386_v43;
      _2386_v43 = _dafny.Seq.of((_this).f14);
      let _2387_i3;
      _2387_i3 = _dafny.ZERO;
      L13: {
        while ((((_2385_v42).contains(p0)) ? ((_2385_v42).get(p0)) : (_dafny.Seq.contains((_2386_v43)[_module.__default.safeIndex(p0, new BigNumber((_2386_v43).length))], _2350_v10)))) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2387_i3)) {
              break L13;
            }
            _2387_i3 = (_2387_i3).plus(_dafny.ONE);
            if (_module.__default.fm0((_this).f13, globalState)) {
              _2380_v37 = (_2380_v37).Merge((_2380_v37).Merge(_2380_v37));
              let _2388_v44;
              let _nw361 = Array((new BigNumber(8)).toNumber());
              _2388_v44 = _nw361;
              _2388_v44 = _2388_v44;
              _2330_v0 = (_this).fm6((_this).fm9(_dafny.Seq.UnicodeFromString("ywsbusnf"), _2381_v38, globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), ((_2389_v10) => function (_2390_i4) {
                return _2389_v10;
              })(_2350_v10)), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), ((_2391_v10) => function (_2392_i4) {
                return _2391_v10;
              })(_2350_v10))).length)), _2350_v10)).length))).length), (p0).minus(_module.__default.fm1(globalState)), (_this).f10, globalState);
              let _2393_v45;
              let _nw362 = new _module.C4();
              _nw362.__ctor();
              _2393_v45 = _nw362;
              let _2394_v46;
              let _nw363 = Array((new BigNumber(18)).toNumber()).fill(false);
              _2394_v46 = _nw363;
              let _index384 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_2394_v46).length));
              (_2394_v46)[_index384] = ((_2330_v0) ? ((_this).f10) : (true));
              let _index385 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_2394_v46).length));
              (_2394_v46)[_index385] = _2330_v0;
              let _index386 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_2394_v46).length));
              let _index387 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_2394_v46).length));
              let _rhs364 = _2393_v45;
              let _rhs365 = (_2393_v45).fm9(_module.__default.fm38(_2330_v0, p0, globalState), _2381_v38, globalState);
              let _rhs366 = _2330_v0;
              let _rhs367 = (_2337_v5)[_module.__default.safeIndex(p0, new BigNumber((_2337_v5).length))];
              let _lhs256 = globalState;
              let _lhs257 = _2394_v46;
              let _lhs258 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_2394_v46).length));
              let _lhs259 = _2394_v46;
              let _lhs260 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_2394_v46).length));
              _2393_v45 = _rhs364;
              _lhs256.f7 = _rhs365;
              _lhs257[_lhs258] = _rhs366;
              _lhs259[_lhs260] = _rhs367;
              let _index388 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_2394_v46).length));
              (_2394_v46)[_index388] = (_2337_v5)[_module.__default.safeIndex((_this).f13, new BigNumber((_2337_v5).length))];
            } else {
              (globalState).f7 = (((_this).f10) ? ((_module.__default.fm1(globalState)).multipliedBy(new BigNumber(207))) : ((_this).f13));
              let _2395_v47;
              _2395_v47 = _dafny.Seq.UnicodeFromString("rbiusl");
              _2395_v47 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_2396_v10) => function (_2397_i5) {
                return (((_this).f10) ? (_2396_v10) : (_2396_v10));
              })(_2350_v10));
              _2330_v0 = (_this).f10;
              let _2398_v48;
              let _nw364 = new _module.C8();
              _nw364.__ctor(_2350_v10);
              _2398_v48 = _nw364;
              let _2399_v49;
              let _nw365 = Array((new BigNumber(6)).toNumber()).fill(false);
              _2399_v49 = _nw365;
              let _index389 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2399_v49).length));
              (_2399_v49)[_index389] = _2330_v0;
              let _index390 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2399_v49).length));
              (_2399_v49)[_index390] = (((_this).f10) ? ((_this).f10) : ((_2337_v5)[_module.__default.safeIndex((_this).f13, new BigNumber((_2337_v5).length))]));
            }
            _2330_v0 = ((_this).f13).isLessThanOrEqualTo(((_this).f13).minus(_dafny.ZERO));
            (globalState).f1 = (_dafny.ZERO).minus((new BigNumber(((_this).f14).length)).minus((_this).f13));
            let _2400_v50;
            let _init58 = function (_2401_i6) {
              return (_this).f10;
            };
            let _nw366 = Array((new BigNumber(29)).toNumber());
            for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw366.length); _i0_58++) {
              _nw366[_i0_58] = _init58(new BigNumber(_i0_58));
            }
            _2400_v50 = _nw366;
            let _index391 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_2400_v50).length));
            (_2400_v50)[_index391] = (_this).f10;
            let _index392 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_2400_v50).length));
            (_2400_v50)[_index392] = true;
          }
        }
      }
      return;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C12 = class C12 {
    constructor () {
      this._tname = "_module.C12";
      this._f10 = false;
      this.f12 = false;
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f11, f12, f10) {
      let _this = this;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f10 = f10;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f11;
    };
    fm7(p0, globalState) {
      let _this = this;
      return !(new BigNumber(-237)).isEqualTo((new BigNumber(-48)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_module.D2.create_DC6(new _dafny.CodePoint('r'.codePointAt(0))))).cardinality())));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return !(((_dafny.Set.fromElements(_this.f12)).Difference(_dafny.Set.fromElements(_this.f12))).IsProperSubsetOf((_dafny.Set.fromElements((_this).f11)).Union(_dafny.Set.fromElements((_this).f11, (_this).f11))));
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll61 = new _dafny.Set();
        for (const _compr_61 of _dafny.IntegerRange(new BigNumber(930), new BigNumber(542))) {
          let _2402_v0 = _compr_61;
          if (((new BigNumber(930)).isLessThanOrEqualTo(_2402_v0)) && ((_2402_v0).isLessThan(new BigNumber(542)))) {
            _coll61.add((_2402_v0).plus(new BigNumber(-560)));
          }
        }
        return _coll61;
      }()).length))).plus(new BigNumber(498))).plus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(new BigNumber(405))).length), new BigNumber(-418))));
    };
    fm10(globalState) {
      let _this = this;
      return !(((new BigNumber((_dafny.Set.fromElements(_this.f12)).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), function (_2403_i0) {
        return new BigNumber(210);
      })).length))).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("rhyfgw")).length), new BigNumber((_dafny.Seq.UnicodeFromString("nhdnniyk")).length))));
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Map.Empty;
      if (!((_this).f11)) {
        let _2404_v0;
        _2404_v0 = new _dafny.CodePoint('r'.codePointAt(0));
        let _2405_v1;
        _2405_v1 = _module.D2.create_DC6(_2404_v0);
        let _2406_v2;
        _2406_v2 = _dafny.Seq.of(_2404_v0, (_2405_v1).dtor_cf8);
        let _2407_v3;
        _2407_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,!(p1));
        let _2408_v4;
        _2408_v4 = new BigNumber(926);
        let _2409_v5;
        _2409_v5 = _dafny.Set.fromElements(_2407_v3, _module.__default.fm11((_this).f11, _2408_v4, _2408_v4, globalState), _2407_v3);
        let _2410_v6;
        _2410_v6 = _dafny.Set.fromElements((_2406_v2)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).fm9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), ((_2411_v0) => function (_2412_i0) {
          return _2411_v0;
        })(_2404_v0)), _2409_v5, globalState)), new BigNumber((_2406_v2).length))]);
        _2410_v6 = ((!(p1)) ? (_2410_v6) : (_2410_v6));
        if (p1) {
          let _2413_v7;
          _2413_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2408_v4,(_this).f11);
          _2413_v7 = (_2413_v7).update(_2408_v4, true);
          let _2414_v8;
          _2414_v8 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), ((_2415_v0) => function (_2416_i1) {
            return _2415_v0;
          })(_2404_v0)), _2406_v2);
          let _2417_v9;
          let _2418_v10;
          let _out66;
          let _out67;
          let _outcollector29 = (_this).m6((_2408_v4).plus(_2408_v4), _this.f12, (_2414_v8).Union(_2414_v8), globalState);
          _out66 = _outcollector29[0];
          _out67 = _outcollector29[1];
          _2417_v9 = _out66;
          _2418_v10 = _out67;
          r0 = _this.f12;
          let _2419_v11;
          _2419_v11 = _dafny.MultiSet.fromElements(true, !(p1), p1);
          let _2420_v12;
          _2420_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_2406_v2);
          (_this).f12 = !((new BigNumber(((((_2420_v12).contains(_2417_v9)) ? ((_2420_v12).get(_2417_v9)) : (_2406_v2))).length)).isLessThan((((_2419_v11).contains((_this).f11)) ? ((_2419_v11).get((_this).f11)) : (new BigNumber(239)))));
          _2417_v9 = _2417_v9;
        } else {
          let _2421_v13;
          let _nw367 = Array((new BigNumber(14)).toNumber()).fill(false);
          _2421_v13 = _nw367;
          let _2422_v14;
          _2422_v14 = _module.D0.create_DC2((_this).f11, _2421_v13);
          let _2423_v15;
          _2423_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_2407_v3);
          let _2424_v16;
          _2424_v16 = _dafny.Seq.of(((p1) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_2422_v14).dtor_cf2)) : (_module.__default.fm11(_this.f12, _2408_v4, _2408_v4, globalState))), _2407_v3, _2407_v3, _2407_v3, (((_2423_v15).contains((_this).f11)) ? ((_2423_v15).get((_this).f11)) : (_2407_v3)));
          let _2425_v17;
          _2425_v17 = _dafny.Seq.of(_this.f12);
          let _rhs368 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2424_v16, _2424_v16), _dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), ((_2426_v4, _2427_v3) => function (_2428_i2) {
            return (_module.D3.create_DC11(_2426_v4, _2427_v3)).dtor_cf18;
          })(_2408_v4, _2407_v3)));
          let _rhs369 = _2406_v2;
          let _rhs370 = (_2425_v17)[_module.__default.safeIndex(_2408_v4, new BigNumber((_2425_v17).length))];
          let _lhs261 = _this;
          _2424_v16 = _rhs368;
          r1 = _rhs369;
          _lhs261.f12 = _rhs370;
          let _2429_v18;
          let _init59 = ((_2430_v4) => function (_2431_i3) {
            return (_2431_i3).multipliedBy(_2430_v4);
          })(_2408_v4);
          let _nw368 = Array((new BigNumber(5)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw368.length); _i0_59++) {
            _nw368[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _2429_v18 = _nw368;
          let _2432_v19;
          _2432_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_2429_v18),_2408_v4);
          let _2433_v20;
          _2433_v20 = _dafny.MultiSet.fromElements((((_2432_v19).contains(_dafny.MultiSet.fromElements(_2429_v18))) ? ((_2432_v19).get(_dafny.MultiSet.fromElements(_2429_v18))) : (_2408_v4)), _2408_v4, new BigNumber((_module.__default.fm12((_this).f11, (_this).f10, new BigNumber((_2406_v2).length), _2408_v4, globalState)).length));
          let _2434_v21;
          _2434_v21 = _module.D4.create_DC15(_2408_v4, _2433_v20, _2404_v0, _2408_v4);
          let _2435_v22;
          _2435_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2404_v0,((false) ? (_2434_v21) : (_2434_v21)));
          let _pat_let_tv67 = _2404_v0;
          let _pat_let_tv68 = _2408_v4;
          _2435_v22 = (_2435_v22).update(_2404_v0, function (_pat_let59_0) {
            return function (_2436_dt__update__tmp_h0) {
              return function (_pat_let60_0) {
                return function (_2437_dt__update_hcf25_h0) {
                  return function (_pat_let61_0) {
                    return function (_2438_dt__update_hcf24_h0) {
                      return _module.D4.create_DC15((_2436_dt__update__tmp_h0).dtor_cf23, _2438_dt__update_hcf24_h0, _2437_dt__update_hcf25_h0, (_2436_dt__update__tmp_h0).dtor_cf26);
                    }(_pat_let61_0);
                  }(_dafny.MultiSet.fromElements(_pat_let_tv68));
                }(_pat_let60_0);
              }(_pat_let_tv67);
            }(_pat_let59_0);
          }(_2434_v21));
          let _index393 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_2421_v13).length));
          (_2421_v13)[_index393] = (((_2407_v3).contains((_this).f10)) ? ((_2407_v3).get((_this).f10)) : (!(true)));
          let _2439_v23;
          _2439_v23 = _dafny.Seq.of(new BigNumber(8), _2408_v4);
          let _2440_v24;
          _2440_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_module.D4.create_DC14((_dafny.ZERO).minus(_2408_v4), _2408_v4));
          let _2441_v25;
          _2441_v25 = _dafny.MultiSet.fromElements(_2440_v24, _2440_v24);
          let _2442_v26;
          _2442_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v13,p1);
          let _2443_v27;
          _2443_v27 = _dafny.Set.fromElements(_2408_v4, new BigNumber(11), _2408_v4, new BigNumber((_2433_v20).cardinality()));
          let _2444_v28;
          _2444_v28 = _dafny.Set.fromElements(_2408_v4, new BigNumber((_2443_v27).length));
          let _2445_v29;
          _2445_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2443_v27).length),_2442_v26);
          let _index394 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_2421_v13).length));
          let _rhs371 = (new BigNumber((_2439_v23).length)).plus(_2408_v4);
          let _rhs372 = p1;
          let _rhs373 = (_2441_v25).IsSubsetOf(_dafny.MultiSet.fromElements(_2440_v24));
          let _rhs374 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2421_v13,(_2425_v17)[_module.__default.safeIndex(_2408_v4, new BigNumber((_2425_v17).length))])).Merge((_2442_v26).Merge((((_2445_v29).contains(new BigNumber((_2425_v17).length))) ? ((_2445_v29).get(new BigNumber((_2425_v17).length))) : (_dafny.Map.Empty.slice().updateUnsafe(_2421_v13,_this.f12)))))).length);
          let _lhs262 = globalState;
          let _lhs263 = _2421_v13;
          let _lhs264 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_2421_v13).length));
          let _lhs265 = globalState;
          _lhs262.f1 = _rhs371;
          r0 = _rhs372;
          _lhs263[_lhs264] = _rhs373;
          _lhs265.f0 = _rhs374;
          let _2446_v30;
          let _2447_v31;
          let _out68;
          let _out69;
          let _outcollector30 = _module.__default.m0(_2408_v4, _2421_v13, globalState);
          _out68 = _outcollector30[0];
          _out69 = _outcollector30[1];
          _2446_v30 = _out68;
          _2447_v31 = _out69;
          (globalState).f0 = (_2408_v4).minus(_2408_v4);
        }
        let _2448_v32;
        _2448_v32 = _dafny.Seq.of(!(p1));
        let _2449_v33;
        _2449_v33 = _dafny.MultiSet.fromElements(_2406_v2);
        let _2450_v35;
        _2450_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2408_v4,_this.f12);
        let _2451_v36;
        _2451_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2450_v35).length),p1);
        let _rhs375 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2448_v32, _dafny.Seq.update(_dafny.Seq.of(true), _module.__default.safeIndex(_2408_v4, new BigNumber((_dafny.Seq.of(true)).length)), (_this).fm8(new BigNumber((function () {
          let _coll62 = new _dafny.Map();
          for (const _compr_62 of (_2451_v36).Keys.Elements) {
            let _2452_v34 = _compr_62;
            if ((_2451_v36).contains(_2452_v34)) {
              _coll62.push([_module.__default.safeModuloInt(_2452_v34, _2408_v4),_2407_v3]);
            }
          }
          return _coll62;
        }()).length), (_this).f11, false, globalState))), _2448_v32);
        let _rhs376 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2406_v2, _2406_v2), _2406_v2);
        let _rhs377 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), ((_2453_v0) => function (_2454_i4) {
          return _2453_v0;
        })(_2404_v0)), _2406_v2), _dafny.Seq.UnicodeFromString("y"));
        let _rhs378 = _2449_v33;
        _2448_v32 = _rhs375;
        r1 = _rhs376;
        r0 = _rhs377;
        _2449_v33 = _rhs378;
        if (!(false) || ((_this).f10)) {
          let _2455_v37;
          let _nw369 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _2455_v37 = _nw369;
          let _2456_v38;
          _2456_v38 = _dafny.Seq.of(new BigNumber((_2448_v32).length), _2408_v4);
          let _index395 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2455_v37).length));
          (_2455_v37)[_index395] = new BigNumber((_2456_v38).length);
          let _index396 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2455_v37).length));
          (_2455_v37)[_index396] = (_2408_v4).multipliedBy(_2408_v4);
          let _2457_v39;
          let _nw370 = Array((new BigNumber(9)).toNumber()).fill(false);
          _2457_v39 = _nw370;
          let _index397 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2457_v39).length));
          (_2457_v39)[_index397] = _this.f12;
          let _2458_v40;
          _2458_v40 = _module.D4.create_DC14((_2455_v37)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2455_v37).length))], _2408_v4);
          let _index398 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2457_v39).length));
          let _rhs379 = (_this).f10;
          let _rhs380 = _2408_v4;
          let _rhs381 = (_2458_v40).dtor_cf21;
          let _lhs266 = _2457_v39;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2457_v39).length));
          let _lhs268 = globalState;
          let _lhs269 = globalState;
          _lhs266[_lhs267] = _rhs379;
          _lhs268.f0 = _rhs380;
          _lhs269.f7 = _rhs381;
          let _2459_v41;
          _2459_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2408_v4,new BigNumber((_2448_v32).length));
          (globalState).f0 = _module.__default.safeModuloInt(new BigNumber(579), (((_2459_v41).contains((_2455_v37)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2455_v37).length))])) ? ((_2459_v41).get((_2455_v37)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2455_v37).length))])) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f11)).length)))));
          _2459_v41 = (_2459_v41).update(_2408_v4, ((_2455_v37)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2455_v37).length))]).minus(_2408_v4));
          (_this).f12 = (_2448_v32)[_module.__default.safeIndex(_2408_v4, new BigNumber((_2448_v32).length))];
        } else {
          let _2460_v42;
          _2460_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2408_v4);
          let _2461_v43;
          let _nw371 = new _module.C3();
          _nw371.__ctor(p1, !(_2460_v42).contains((_2448_v32)[_module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_2448_v32).length))]));
          _2461_v43 = _nw371;
          let _2462_v44;
          let _nw372 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _2462_v44 = _nw372;
          let _index399 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2462_v44).length));
          (_2462_v44)[_index399] = new BigNumber(((_2451_v36).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2408_v4,true))).length);
          let _2463_v45;
          _2463_v45 = _dafny.Seq.of(new BigNumber((_2448_v32).length), _2408_v4, new BigNumber(963));
          let _index400 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2462_v44).length));
          (_2462_v44)[_index400] = (_2463_v45)[_module.__default.safeIndex(new BigNumber((_2406_v2).length), new BigNumber((_2463_v45).length))];
          let _2464_v46;
          _2464_v46 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p1),(_this).f10);
          (globalState).f7 = ((_dafny.ZERO).minus((_2462_v44)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2462_v44).length))])).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2448_v32, _2448_v32), _module.__default.safeIndex(new BigNumber(((_2464_v46).update(_dafny.Set.fromElements(false), _this.f12)).length), new BigNumber((_dafny.Seq.Concat(_2448_v32, _2448_v32)).length)), (_this).f11)).length));
          r0 = _this.f12;
          let _2465_v47;
          let _nw373 = new _module.C9();
          _nw373.__ctor((_2461_v43).f20, p1);
          _2465_v47 = _nw373;
        }
        r0 = _module.__default.fm0(_2408_v4, globalState);
      } else {
        let _2466_v48;
        _2466_v48 = _dafny.Seq.UnicodeFromString("cnhqf");
        let _2467_v49;
        _2467_v49 = new BigNumber(633);
        let _2468_v50;
        _2468_v50 = new _dafny.CodePoint('u'.codePointAt(0));
        let _2469_v51;
        _2469_v51 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.IsPrefixOf(_2466_v48, _dafny.Seq.update(_2466_v48, _module.__default.safeIndex(_2467_v49, new BigNumber((_2466_v48).length)), _2468_v50)));
        _2469_v51 = ((_2469_v51).update(_this.f12, p1)).Merge(_2469_v51);
        r0 = _dafny.Seq.contains(_2466_v48, _2468_v50);
        let _2470_v52;
        _2470_v52 = _dafny.Seq.of(p1);
        let _2471_v53;
        _2471_v53 = _module.D3.create_DC11(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2470_v52, _2470_v52), _module.__default.safeIndex(_2467_v49, new BigNumber((_dafny.Seq.Concat(_2470_v52, _2470_v52)).length)), true)).length), _2469_v51);
        let _source40 = _2471_v53;
        if (_source40.is_DC10) {
          let _2472___mcc_h0 = (_source40).cf14;
          let _2473___mcc_h1 = (_source40).cf15;
          let _2474___mcc_h2 = (_source40).cf16;
          let _2475_cf16 = _2474___mcc_h2;
          let _2476_cf15 = _2473___mcc_h1;
          let _2477_cf14 = _2472___mcc_h0;
          let _2478_v54;
          _2478_v54 = _dafny.Set.fromElements((_this).fm8(_2467_v49, _this.f12, _this.f12, globalState));
          let _2479_v55;
          _2479_v55 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2467_v49));
          let _2480_v57;
          _2480_v57 = _module.D13.create_DC39(new BigNumber((function () {
  let _coll63 = new _dafny.Map();
  for (const _compr_63 of _dafny.IntegerRange(new BigNumber(-537), new BigNumber(695))) {
    let _2481_v56 = _compr_63;
    if (((new BigNumber(-537)).isLessThanOrEqualTo(_2481_v56)) && ((_2481_v56).isLessThan(new BigNumber(695)))) {
      _coll63.push([(_2481_v56).multipliedBy(_2467_v49),_module.D3.create_DC10(_2477_cf14, _2476_cf15, true)]);
    }
  }
  return _coll63;
}()).length), _2476_cf15);
          let _2482_v58;
          _2482_v58 = _dafny.MultiSet.fromElements(_this.f12);
          let _2483_v59;
          _2483_v59 = _dafny.Set.fromElements(_2469_v51, _2469_v51);
          let _2484_v60;
          _2484_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v49,(_this).fm9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-607)), ((_2485_v50) => function (_2486_i5) {
            return _2485_v50;
          })(_2468_v50)), _2483_v59, globalState));
          let _2487_v61;
          _2487_v61 = _dafny.MultiSet.fromElements(_2484_v60);
          let _2488_v62;
          let _nw374 = Array((new BigNumber(27)).toNumber());
          _nw374[(_dafny.ZERO).toNumber()] = ((_module.D2.create_DC7(new BigNumber((_2478_v54).length), _2476_cf15, _2467_v49)).dtor_cf10) || ((_this).f10);
          _nw374[(_dafny.ONE).toNumber()] = !((_this).f10);
          _nw374[(new BigNumber(2)).toNumber()] = !(_this.f12);
          _nw374[(new BigNumber(3)).toNumber()] = !((_2479_v55).IsProperSubsetOf(_2479_v55));
          _nw374[(new BigNumber(4)).toNumber()] = (_this).f10;
          _nw374[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw374[(new BigNumber(6)).toNumber()] = (_2470_v52)[_module.__default.safeIndex(_2467_v49, new BigNumber((_2470_v52).length))];
          _nw374[(new BigNumber(7)).toNumber()] = (_2467_v49).isLessThanOrEqualTo(_2467_v49);
          _nw374[(new BigNumber(8)).toNumber()] = _this.f12;
          _nw374[(new BigNumber(9)).toNumber()] = !(_2476_cf15);
          _nw374[(new BigNumber(10)).toNumber()] = _2476_cf15;
          _nw374[(new BigNumber(11)).toNumber()] = (_2476_cf15) === (!((_this).f10));
          _nw374[(new BigNumber(12)).toNumber()] = (_2480_v57).dtor_cf66;
          _nw374[(new BigNumber(13)).toNumber()] = false;
          _nw374[(new BigNumber(14)).toNumber()] = (_2467_v49).isLessThan(new BigNumber((_2482_v58).cardinality()));
          _nw374[(new BigNumber(15)).toNumber()] = (_this).f11;
          _nw374[(new BigNumber(16)).toNumber()] = _this.f12;
          _nw374[(new BigNumber(17)).toNumber()] = _2475_cf16;
          _nw374[(new BigNumber(18)).toNumber()] = _this.f12;
          _nw374[(new BigNumber(19)).toNumber()] = !(!(!((_this).fm10(globalState))));
          _nw374[(new BigNumber(20)).toNumber()] = false;
          _nw374[(new BigNumber(21)).toNumber()] = _2475_cf16;
          _nw374[(new BigNumber(22)).toNumber()] = (_2479_v55).IsProperSubsetOf(_2479_v55);
          _nw374[(new BigNumber(23)).toNumber()] = _dafny.areEqual(_2477_cf14, _2477_cf14);
          _nw374[(new BigNumber(24)).toNumber()] = p1;
          _nw374[(new BigNumber(25)).toNumber()] = !(_this.f12);
          _nw374[(new BigNumber(26)).toNumber()] = !(_2487_v61).contains(_2484_v60);
          _2488_v62 = _nw374;
          let _index401 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2488_v62).length));
          (_2488_v62)[_index401] = (_this).f10;
          let _2489_v63;
          let _nw375 = Array((new BigNumber(21)).toNumber()).fill([]);
          _2489_v63 = _nw375;
          let _2490_v64;
          _2490_v64 = _module.D17.create_DC50(_2489_v63);
          let _pat_let_tv69 = _2489_v63;
          let _2491_v65;
          let _nw376 = Array((new BigNumber(15)).toNumber());
          _nw376[(_dafny.ZERO).toNumber()] = _2490_v64;
          _nw376[(_dafny.ONE).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(2)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(3)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(4)).toNumber()] = function (_pat_let62_0) {
            return function (_2492_dt__update__tmp_h1) {
              return function (_pat_let63_0) {
                return function (_2493_dt__update_hcf86_h0) {
                  return _module.D17.create_DC50(_2493_dt__update_hcf86_h0);
                }(_pat_let63_0);
              }(_pat_let_tv69);
            }(_pat_let62_0);
          }(_2490_v64);
          _nw376[(new BigNumber(5)).toNumber()] = _module.D17.create_DC50(_2489_v63);
          _nw376[(new BigNumber(6)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(7)).toNumber()] = _module.D17.create_DC50(_2489_v63);
          _nw376[(new BigNumber(8)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(9)).toNumber()] = _module.D17.create_DC50(_2489_v63);
          _nw376[(new BigNumber(10)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(11)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(12)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(13)).toNumber()] = _2490_v64;
          _nw376[(new BigNumber(14)).toNumber()] = _module.D17.create_DC50(_2489_v63);
          _2491_v65 = _nw376;
          let _index402 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2491_v65).length));
          (_2491_v65)[_index402] = _2490_v64;
          let _index403 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2488_v62).length));
          let _index404 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2491_v65).length));
          let _rhs382 = !(p1);
          let _rhs383 = _module.D17.create_DC50(_2489_v63);
          let _lhs270 = _2488_v62;
          let _lhs271 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2488_v62).length));
          let _lhs272 = _2491_v65;
          let _lhs273 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2491_v65).length));
          _lhs270[_lhs271] = _rhs382;
          _lhs272[_lhs273] = _rhs383;
          (globalState).f7 = _2467_v49;
          let _index405 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2488_v62).length));
          (_2488_v62)[_index405] = _2476_cf15;
          let _index406 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2488_v62).length));
          (_2488_v62)[_index406] = !(_this.f12);
        } else if (_source40.is_DC11) {
          let _2494___mcc_h3 = (_source40).cf17;
          let _2495___mcc_h4 = (_source40).cf18;
          let _2496_cf18 = _2495___mcc_h4;
          let _2497_cf17 = _2494___mcc_h3;
          (_this).f12 = (_this).fm10(globalState);
          (globalState).f0 = (_2497_cf17).minus(new BigNumber(131));
          (_this).f12 = (_this).fm10(globalState);
          let _2498_v66;
          let _init60 = function (_2499_i6) {
            return !((_this).f11);
          };
          let _nw377 = Array((new BigNumber(9)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw377.length); _i0_60++) {
            _nw377[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2498_v66 = _nw377;
          let _2500_v67;
          _2500_v67 = _dafny.Set.fromElements(_2496_cf18);
          let _2501_v68;
          _2501_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2498_v66,_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).fm9(_2466_v48, _2500_v67, globalState)), new BigNumber(98)));
          let _2502_v69;
          _2502_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2468_v50,_dafny.Map.Empty.slice().updateUnsafe(_2498_v66,_2467_v49));
          let _rhs384 = new BigNumber(-101);
          let _rhs385 = (((_2501_v68).update(_2498_v66, _2467_v49)).Merge(_2501_v68)).Merge((((_2502_v69).contains(_2468_v50)) ? ((_2502_v69).get(_2468_v50)) : (_2501_v68)));
          let _rhs386 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_2497_cf17));
          let _rhs387 = _this.f12;
          let _lhs274 = globalState;
          let _lhs275 = _this;
          _2497_cf17 = _rhs384;
          _2501_v68 = _rhs385;
          _lhs274.f0 = _rhs386;
          _lhs275.f12 = _rhs387;
        } else if (_source40.is_DC9) {
          let _2503___mcc_h5 = (_source40).cf13;
          let _2504_cf13 = _2503___mcc_h5;
          let _2505_v70;
          _2505_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v49,_2466_v48);
          let _2506_v71;
          _2506_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v49,_this.f12);
          let _2507_v72;
          _2507_v72 = _dafny.MultiSet.fromElements(_2467_v49);
          let _2508_v73;
          let _nw378 = Array((new BigNumber(23)).toNumber());
          _nw378[(_dafny.ZERO).toNumber()] = (_2467_v49).isLessThan(_2467_v49);
          _nw378[(_dafny.ONE).toNumber()] = !(p1);
          _nw378[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsProperPrefixOf((((_2505_v70).contains(new BigNumber((_2506_v71).length))) ? ((_2505_v70).get(new BigNumber((_2506_v71).length))) : (_2466_v48)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), ((_2509_v50) => function (_2510_i7) {
            return _2509_v50;
          })(_2468_v50)));
          _nw378[(new BigNumber(3)).toNumber()] = (_this).f10;
          _nw378[(new BigNumber(4)).toNumber()] = (_this).f11;
          _nw378[(new BigNumber(5)).toNumber()] = _this.f12;
          _nw378[(new BigNumber(6)).toNumber()] = (_this).f10;
          _nw378[(new BigNumber(7)).toNumber()] = p1;
          _nw378[(new BigNumber(8)).toNumber()] = (_this).f11;
          _nw378[(new BigNumber(9)).toNumber()] = false;
          _nw378[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw378[(new BigNumber(11)).toNumber()] = _this.f12;
          _nw378[(new BigNumber(12)).toNumber()] = p1;
          _nw378[(new BigNumber(13)).toNumber()] = (_this).f10;
          _nw378[(new BigNumber(14)).toNumber()] = p1;
          _nw378[(new BigNumber(15)).toNumber()] = !((_this).fm7((_this).f10, globalState));
          _nw378[(new BigNumber(16)).toNumber()] = !(!(true)) || (p1);
          _nw378[(new BigNumber(17)).toNumber()] = p1;
          _nw378[(new BigNumber(18)).toNumber()] = !(!((_2467_v49).isLessThanOrEqualTo(_2467_v49)));
          _nw378[(new BigNumber(19)).toNumber()] = (_this).f11;
          _nw378[(new BigNumber(20)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber(222))).IsSubsetOf(_2507_v72);
          _nw378[(new BigNumber(21)).toNumber()] = (_this).f10;
          _nw378[(new BigNumber(22)).toNumber()] = _this.f12;
          _2508_v73 = _nw378;
          let _index407 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2508_v73).length));
          (_2508_v73)[_index407] = !(((_this).f11) || (p1));
          let _index408 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2508_v73).length));
          (_2508_v73)[_index408] = _dafny.areEqual(_2470_v52, _dafny.Seq.Concat(_module.__default.fm40(_this.f12, (_this).f11, globalState), _2470_v52));
          let _2511_v74;
          let _init61 = function (_2512_i8) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          };
          let _nw379 = Array((new BigNumber(6)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw379.length); _i0_61++) {
            _nw379[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2511_v74 = _nw379;
          let _2513_v75;
          _2513_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2511_v74,_module.__default.fm1(globalState));
          _2513_v75 = _2513_v75;
          let _2514_v76;
          _2514_v76 = _dafny.Set.fromElements(new BigNumber((_2466_v48).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2470_v52,_2467_v49)).length));
          let _index409 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2508_v73).length));
          (_2508_v73)[_index409] = !(_2514_v76).contains(_2467_v49);
          let _2515_v77;
          _2515_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v49,_2508_v73);
          let _2516_v78;
          let _nw380 = new _module.C11();
          _nw380.__ctor(new BigNumber((_2515_v77).length), _dafny.Seq.Concat(_2466_v48, _2466_v48), (_2507_v72).IsSubsetOf(_2507_v72));
          _2516_v78 = _nw380;
        } else {
          let _2517___mcc_h6 = (_source40).cf19;
          let _2518_cf19 = _2517___mcc_h6;
          let _2519_v79;
          let _nw381 = Array((new BigNumber(16)).toNumber()).fill(false);
          _2519_v79 = _nw381;
          let _index410 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_2519_v79).length));
          (_2519_v79)[_index410] = (!((_this).f10)) === ((_this).f10);
          let _2520_v80;
          _2520_v80 = _dafny.MultiSet.fromElements(_2467_v49);
          let _index411 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_2519_v79).length));
          let _rhs388 = ((((_this).f11) ? (_2520_v80) : (_module.__default.fm18(_2467_v49, _module.__default.fm51(globalState), (_this).f10, globalState)))).IsProperSubsetOf(_2520_v80);
          let _rhs389 = _this.f12;
          let _lhs276 = _this;
          let _lhs277 = _2519_v79;
          let _lhs278 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_2519_v79).length));
          _lhs276.f12 = _rhs388;
          _lhs277[_lhs278] = _rhs389;
          let _2521_v81;
          _2521_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(631),_2469_v51);
          let _2522_v82;
          _2522_v82 = _module.D14.create_DC41(_2521_v81);
          _2522_v82 = _2522_v82;
          let _2523_v83;
          _2523_v83 = _module.D20.create_DC60(new BigNumber((_2466_v48).length));
          let _2524_v84;
          _2524_v84 = _dafny.Set.fromElements((_2519_v79)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_2519_v79).length))]);
          let _2525_v85;
          _2525_v85 = _dafny.Seq.of(new BigNumber((_2520_v80).cardinality()));
          let _2526_v86;
          let _nw382 = Array((new BigNumber(7)).toNumber());
          _nw382[(_dafny.ZERO).toNumber()] = ((p1) ? (_2467_v49) : (_2467_v49));
          _nw382[(_dafny.ONE).toNumber()] = _2467_v49;
          _nw382[(new BigNumber(2)).toNumber()] = _2467_v49;
          _nw382[(new BigNumber(3)).toNumber()] = _2467_v49;
          _nw382[(new BigNumber(4)).toNumber()] = (_2523_v83).dtor_cf102;
          _nw382[(new BigNumber(5)).toNumber()] = (_2467_v49).plus((_dafny.ZERO).minus(new BigNumber((_2524_v84).length)));
          _nw382[(new BigNumber(6)).toNumber()] = (_2467_v49).minus(new BigNumber((_2525_v85).length));
          _2526_v86 = _nw382;
          let _index412 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_2526_v86).length));
          (_2526_v86)[_index412] = (_2467_v49).plus(_2467_v49);
          let _2527_v87;
          _2527_v87 = _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe(_2467_v49,(_this).f10));
          let _index413 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_2526_v86).length));
          let _rhs390 = (_2467_v49).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_2520_v80).Difference(_2520_v80)).cardinality()))));
          let _rhs391 = _2467_v49;
          let _rhs392 = (((_2520_v80).contains(((p1) ? (new BigNumber(255)) : (new BigNumber(((_2527_v87).dtor_cf53).length))))) ? ((_2520_v80).get(((p1) ? (new BigNumber(255)) : (new BigNumber(((_2527_v87).dtor_cf53).length))))) : (new BigNumber((_2524_v84).length)));
          let _lhs279 = globalState;
          let _lhs280 = globalState;
          let _lhs281 = _2526_v86;
          let _lhs282 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_2526_v86).length));
          _lhs279.f0 = _rhs390;
          _lhs280.f0 = _rhs391;
          _lhs281[_lhs282] = _rhs392;
          let _2528_v88;
          _2528_v88 = _dafny.Map.Empty.slice().updateUnsafe((_2526_v86)[_module.__default.safeIndex(new BigNumber(516), new BigNumber((_2526_v86).length))],_this.f12);
          let _2529_v89;
          let _nw383 = new _module.C9();
          _nw383.__ctor((_this).f11, (((_2528_v88).contains((_2526_v86)[_module.__default.safeIndex(new BigNumber(516), new BigNumber((_2526_v86).length))])) ? ((_2528_v88).get((_2526_v86)[_module.__default.safeIndex(new BigNumber(516), new BigNumber((_2526_v86).length))])) : (true)));
          _2529_v89 = _nw383;
        }
        let _2530_v90;
        let _nw384 = Array((new BigNumber(27)).toNumber()).fill([]);
        _2530_v90 = _nw384;
        let _2531_v91;
        _2531_v91 = _dafny.MultiSet.fromElements((_this).f11, (_this).f10);
        let _2532_v92;
        _2532_v92 = _dafny.Seq.of(_2467_v49, new BigNumber((_2531_v91).cardinality()));
        let _2533_v93;
        let _nw385 = new _module.C9();
        _nw385.__ctor(_this.f12, !(true));
        _2533_v93 = _nw385;
        let _2534_v94;
        _2534_v94 = _dafny.Seq.of(_2533_v93);
        let _2535_v95;
        _2535_v95 = _module.D19.create_DC58(new BigNumber((_2534_v94).length), (_this).f10);
        let _2536_v96;
        let _nw386 = Array((new BigNumber(16)).toNumber());
        _nw386[(_dafny.ZERO).toNumber()] = _2467_v49;
        _nw386[(_dafny.ONE).toNumber()] = _module.__default.fm1(globalState);
        _nw386[(new BigNumber(2)).toNumber()] = new BigNumber(629);
        _nw386[(new BigNumber(3)).toNumber()] = new BigNumber((_2532_v92).length);
        _nw386[(new BigNumber(4)).toNumber()] = _2467_v49;
        _nw386[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-399)), ((_2537_v50) => function (_2538_i9) {
          return _2537_v50;
        })(_2468_v50))).length);
        _nw386[(new BigNumber(6)).toNumber()] = new BigNumber((_2466_v48).length);
        _nw386[(new BigNumber(7)).toNumber()] = (_2535_v95).dtor_cf99;
        _nw386[(new BigNumber(8)).toNumber()] = new BigNumber(83);
        _nw386[(new BigNumber(9)).toNumber()] = new BigNumber(-750);
        _nw386[(new BigNumber(10)).toNumber()] = _2467_v49;
        _nw386[(new BigNumber(11)).toNumber()] = _2467_v49;
        _nw386[(new BigNumber(12)).toNumber()] = new BigNumber(208);
        _nw386[(new BigNumber(13)).toNumber()] = _2467_v49;
        _nw386[(new BigNumber(14)).toNumber()] = _module.__default.fm1(globalState);
        _nw386[(new BigNumber(15)).toNumber()] = _2467_v49;
        _2536_v96 = _nw386;
        let _index414 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_2530_v90).length));
        (_2530_v90)[_index414] = _2536_v96;
        let _2539_v97;
        _2539_v97 = _module.D3.create_DC12(_module.D3.create_DC11((_dafny.ZERO).minus(_2467_v49), _2469_v51));
        let _2540_v98;
        _2540_v98 = _dafny.Set.fromElements(_this.f12, _this.f12);
        let _2541_v99;
        _2541_v99 = _dafny.Set.fromElements(_2540_v98);
        let _2542_v100;
        _2542_v100 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC12(_module.D3.create_DC9(_2541_v99)),true);
        let _2543_v101;
        _2543_v101 = _dafny.Map.Empty.slice().updateUnsafe((_2533_v93).f17,new BigNumber(698));
        let _2544_v102;
        _2544_v102 = _dafny.MultiSet.fromElements(_2467_v49);
        let _2545_v103;
        let _nw387 = Array((new BigNumber(19)).toNumber());
        _nw387[(_dafny.ZERO).toNumber()] = (_2533_v93).f17;
        _nw387[(_dafny.ONE).toNumber()] = _this.f12;
        _nw387[(new BigNumber(2)).toNumber()] = !(_2542_v100).contains(_2539_v97);
        _nw387[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2466_v48, _2466_v48);
        _nw387[(new BigNumber(4)).toNumber()] = (_2467_v49).isEqualTo(_2467_v49);
        _nw387[(new BigNumber(5)).toNumber()] = (new BigNumber((_2469_v51).length)).isLessThanOrEqualTo(_2467_v49);
        _nw387[(new BigNumber(6)).toNumber()] = (_this).f10;
        _nw387[(new BigNumber(7)).toNumber()] = (_this).f10;
        _nw387[(new BigNumber(8)).toNumber()] = (_2531_v91).IsProperSubsetOf(_dafny.MultiSet.fromElements((_2533_v93).f17));
        _nw387[(new BigNumber(9)).toNumber()] = (_2543_v101).contains(!((_this).f10));
        _nw387[(new BigNumber(10)).toNumber()] = _this.f12;
        _nw387[(new BigNumber(11)).toNumber()] = (_2544_v102).IsProperSubsetOf(_dafny.MultiSet.fromElements(_2467_v49, new BigNumber((_2466_v48).length)));
        _nw387[(new BigNumber(12)).toNumber()] = !(true);
        _nw387[(new BigNumber(13)).toNumber()] = p1;
        _nw387[(new BigNumber(14)).toNumber()] = (_2533_v93).f17;
        _nw387[(new BigNumber(15)).toNumber()] = (_this).f11;
        _nw387[(new BigNumber(16)).toNumber()] = (_2533_v93).f17;
        _nw387[(new BigNumber(17)).toNumber()] = (_2533_v93).f17;
        _nw387[(new BigNumber(18)).toNumber()] = (_this).fm6(_2467_v49, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_2466_v48, _module.__default.safeIndex(new BigNumber((_2470_v52).length), new BigNumber((_2466_v48).length)), _2468_v50)).length)), _2467_v49, _this.f12, globalState);
        _2545_v103 = _nw387;
        let _index415 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2545_v103).length));
        (_2545_v103)[_index415] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-354)), ((_2546_v50) => function (_2547_i10) {
          return _2546_v50;
        })(_2468_v50)), _2466_v48);
        let _2548_v104;
        _2548_v104 = _module.D15.create_DC46();
        let _index416 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_2530_v90).length));
        let _index417 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2545_v103).length));
        let _rhs393 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-732)), ((_2549_v50) => function (_2550_i11) {
          return _2549_v50;
        })(_2468_v50)), _2466_v48);
        let _rhs394 = _2536_v96;
        let _rhs395 = (_2533_v93).f17;
        let _rhs396 = _module.D15.create_DC46();
        let _lhs283 = _2530_v90;
        let _lhs284 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_2530_v90).length));
        let _lhs285 = _2545_v103;
        let _lhs286 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2545_v103).length));
        r0 = _rhs393;
        _lhs283[_lhs284] = _rhs394;
        _lhs285[_lhs286] = _rhs395;
        _2548_v104 = _rhs396;
        (globalState).f7 = (_2467_v49).plus(_2467_v49);
      }
      let _2551_v105;
      _2551_v105 = _dafny.Seq.UnicodeFromString("kaqhnb");
      let _2552_v106;
      _2552_v106 = new BigNumber(557);
      let _2553_v107;
      let _nw388 = new _module.C11();
      _nw388.__ctor(_2552_v106, _2551_v105, p1);
      _2553_v107 = _nw388;
      let _2554_v108;
      _2554_v108 = _dafny.Seq.of(_2553_v107, _2553_v107, _2553_v107);
      let _2555_v109;
      _2555_v109 = _dafny.Map.Empty.slice().updateUnsafe(_2554_v108,_this.f12);
      let _2556_v110;
      _2556_v110 = _dafny.Seq.of((_2553_v107).f10);
      let _2557_v111;
      _2557_v111 = _module.D0.create_DC1(_2551_v105);
      let _2558_v112;
      _2558_v112 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_2551_v105);
      let _2559_v113;
      let _nw389 = Array((new BigNumber(26)).toNumber());
      _nw389[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("prqda");
      _nw389[(_dafny.ONE).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wx"), _2551_v105);
      _nw389[(new BigNumber(3)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(4)).toNumber()] = (((((_2555_v109).contains(_2554_v108)) ? ((_2555_v109).get(_2554_v108)) : ((_this).f10))) ? (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("wjuvsa"), _module.__default.safeIndex(_2552_v106, new BigNumber((_dafny.Seq.UnicodeFromString("wjuvsa")).length)), new _dafny.CodePoint('h'.codePointAt(0)))) : (_2551_v105));
      _nw389[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-557)), function (_2560_i12) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _2551_v105);
      _nw389[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_2551_v105, _2551_v105);
      _nw389[(new BigNumber(7)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(8)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jnfj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), function (_2561_i13) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }));
      _nw389[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm13((_2553_v107).f10, (_2556_v110)[_module.__default.safeIndex(_2552_v106, new BigNumber((_2556_v110).length))], globalState), _2551_v105);
      _nw389[(new BigNumber(11)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(12)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(13)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(14)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(15)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(16)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(17)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(18)).toNumber()] = ((!((_this).f10)) ? (_2551_v105) : (_2551_v105));
      _nw389[(new BigNumber(19)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(20)).toNumber()] = (_2557_v111).dtor_cf1;
      _nw389[(new BigNumber(21)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(22)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(23)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(24)).toNumber()] = _2551_v105;
      _nw389[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat((((_2558_v112).contains((_2553_v107).f10)) ? ((_2558_v112).get((_2553_v107).f10)) : (_2551_v105)), _2551_v105);
      _2559_v113 = _nw389;
      let _index418 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2559_v113).length));
      (_2559_v113)[_index418] = _2551_v105;
      let _index419 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2559_v113).length));
      (_2559_v113)[_index419] = _2551_v105;
      let _2562_v114;
      _2562_v114 = new _dafny.CodePoint('a'.codePointAt(0));
      (_this).f12 = _dafny.Seq.contains(_2551_v105, _2562_v114);
      let _2563_v115;
      _2563_v115 = _module.D15.create_DC45(_2552_v106, p1, p1);
      (globalState).f0 = (_2563_v115).dtor_cf80;
      (_this).f12 = (_2552_v106).isLessThanOrEqualTo(_module.__default.fm1(globalState));
      let _2564_v116;
      _2564_v116 = _dafny.Map.Empty.slice().updateUnsafe(_2552_v106,new BigNumber(330));
      let _2565_v117;
      _2565_v117 = _module.D7.create_DC22((_2564_v116).Merge(_2564_v116));
      _2565_v117 = _2565_v117;
      r0 = !(_2552_v106).isEqualTo((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-380), _2552_v106)));
      r1 = _2551_v105;
      r2 = _2564_v116;
      return [r0, r1, r2];
    }
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _2566_v0;
      _2566_v0 = new BigNumber(-364);
      let _2567_v1;
      let _init62 = ((_2568_v0) => function (_2569_i1) {
        return (_2569_i1).multipliedBy(_2568_v0);
      })(_2566_v0);
      let _nw390 = Array((new BigNumber(10)).toNumber());
      for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw390.length); _i0_62++) {
        _nw390[_i0_62] = _init62(new BigNumber(_i0_62));
      }
      _2567_v1 = _nw390;
      let _2570_v2;
      _2570_v2 = _module.D10.create_DC31((_this).f10, _this.f12, _2566_v0);
      let _2571_v3;
      _2571_v3 = _dafny.Seq.of((_2570_v2).dtor_cf50);
      let _2572_v4;
      _2572_v4 = _module.D14.create_DC42(_module.D1.create_DC4(_2566_v0, _2567_v1), _2571_v3, (_this).f11);
      let _2573_i0;
      _2573_i0 = _dafny.ZERO;
      L14: {
        while ((_2572_v4).dtor_cf74) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2573_i0)) {
              break L14;
            }
            _2573_i0 = (_2573_i0).plus(_dafny.ONE);
            let _2574_v5;
            let _nw391 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2574_v5 = _nw391;
            let _2575_v6;
            _2575_v6 = _dafny.Seq.UnicodeFromString("rbbgv");
            let _index420 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length));
            (_2574_v5)[_index420] = _dafny.Seq.Concat(_2575_v6, _2575_v6);
            let _index421 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length));
            (_2574_v5)[_index421] = _2575_v6;
            let _2576_v7;
            _2576_v7 = new _dafny.CodePoint('u'.codePointAt(0));
            _2576_v7 = _2576_v7;
            let _2577_v8;
            _2577_v8 = _module.D6.create_DC20(_this.f12, false, (_this).f11, _2566_v0);
            let _2578_v9;
            _2578_v9 = _module.D4.create_DC14(new BigNumber((_dafny.Set.fromElements((_this).f11, (_this).f11, (_2577_v8).dtor_cf34)).length), _2566_v0);
            let _source41 = _2578_v9;
            if (_source41.is_DC14) {
              let _2579___mcc_h0 = (_source41).cf21;
              let _2580___mcc_h1 = (_source41).cf22;
              let _2581_cf22 = _2580___mcc_h1;
              let _2582_cf21 = _2579___mcc_h0;
              let _2583_v10;
              let _nw392 = new _module.C8();
              _nw392.__ctor(_2576_v7);
              _2583_v10 = _nw392;
              let _2584_v11;
              let _nw393 = Array((new BigNumber(10)).toNumber()).fill(false);
              _2584_v11 = _nw393;
              let _index422 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_2584_v11).length));
              (_2584_v11)[_index422] = _this.f12;
              let _index423 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_2584_v11).length));
              (_2584_v11)[_index423] = (_this).f10;
              let _2585_v12;
              let _nw394 = new _module.C3();
              _nw394.__ctor((_2582_cf21).isEqualTo(_2566_v0), !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_2586_v5) => function (_2587_i2) {
                return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-525)), ((_2588_v5) => function (_2589_i3) {
                  return new BigNumber(((_2588_v5)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_2588_v5).length))]).length);
                })(_2586_v5))).length));
              })(_2574_v5)), new BigNumber(((_2574_v5)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length))]).length)));
              _2585_v12 = _nw394;
              let _2590_v13;
              _2590_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2581_cf22,(_2584_v11)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_2584_v11).length))]);
              let _index424 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_2567_v1).length));
              (_2567_v1)[_index424] = new BigNumber(319);
              let _index425 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length));
              let _index426 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_2567_v1).length));
              let _rhs397 = (_dafny.Map.Empty.slice().updateUnsafe(_2582_cf21,(_this).f10)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2582_cf21,(_this).f11));
              let _rhs398 = (_2585_v12).fm8(_2581_cf22, true, (_2581_cf22).isEqualTo(new BigNumber(304)), globalState);
              let _rhs399 = (_2574_v5)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length))];
              let _rhs400 = false;
              let _rhs401 = _2581_cf22;
              let _lhs287 = _this;
              let _lhs288 = _2574_v5;
              let _lhs289 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length));
              let _lhs290 = _2567_v1;
              let _lhs291 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_2567_v1).length));
              _2590_v13 = _rhs397;
              _lhs287.f12 = _rhs398;
              _lhs288[_lhs289] = _rhs399;
              r1 = _rhs400;
              _lhs290[_lhs291] = _rhs401;
            } else if (_source41.is_DC15) {
              let _2591___mcc_h2 = (_source41).cf23;
              let _2592___mcc_h3 = (_source41).cf24;
              let _2593___mcc_h4 = (_source41).cf25;
              let _2594___mcc_h5 = (_source41).cf26;
              let _2595_cf26 = _2594___mcc_h5;
              let _2596_cf25 = _2593___mcc_h4;
              let _2597_cf24 = _2592___mcc_h3;
              let _2598_cf23 = _2591___mcc_h2;
              (globalState).f0 = _module.__default.fm1(globalState);
              let _2599_v14;
              let _nw395 = Array((new BigNumber(7)).toNumber()).fill(false);
              _2599_v14 = _nw395;
              let _2600_v15;
              let _nw396 = Array((new BigNumber(2)).toNumber());
              _nw396[(_dafny.ZERO).toNumber()] = _2599_v14;
              _nw396[(_dafny.ONE).toNumber()] = _2599_v14;
              _2600_v15 = _nw396;
              let _index427 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2600_v15).length));
              (_2600_v15)[_index427] = _2599_v14;
              let _index428 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2600_v15).length));
              (_2600_v15)[_index428] = _2599_v14;
              r1 = !(false);
              let _rhs402 = _2567_v1;
              let _rhs403 = _2598_cf23;
              let _rhs404 = (_this).f11;
              let _rhs405 = (_this).fm7((_this).f10, globalState);
              let _lhs292 = globalState;
              _2567_v1 = _rhs402;
              _lhs292.f7 = _rhs403;
              r1 = _rhs404;
              r1 = _rhs405;
            } else {
              let _2601___mcc_h6 = (_source41).cf20;
              let _2602_cf20 = _2601___mcc_h6;
              let _2603_v16;
              let _init63 = function (_2604_i4) {
                return (_this).f11;
              };
              let _nw397 = Array((new BigNumber(3)).toNumber());
              for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw397.length); _i0_63++) {
                _nw397[_i0_63] = _init63(new BigNumber(_i0_63));
              }
              _2603_v16 = _nw397;
              let _index429 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_2603_v16).length));
              (_2603_v16)[_index429] = false;
              let _2605_v17;
              _2605_v17 = _dafny.Set.fromElements((((_this).f10) ? (_2603_v16) : (_2603_v16)));
              let _index430 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_2603_v16).length));
              let _rhs406 = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), ((_2606_v7) => function (_2607_i5) {
                return _2606_v7;
              })(_2576_v7)), _2576_v7);
              let _rhs407 = _2566_v0;
              let _rhs408 = _module.__default.fm0(_2566_v0, globalState);
              let _rhs409 = (_2605_v17).Union((_2605_v17).Union(_2605_v17));
              let _rhs410 = _2566_v0;
              let _lhs293 = globalState;
              let _lhs294 = _2603_v16;
              let _lhs295 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_2603_v16).length));
              let _lhs296 = globalState;
              r1 = _rhs406;
              _lhs293.f7 = _rhs407;
              _lhs294[_lhs295] = _rhs408;
              _2605_v17 = _rhs409;
              _lhs296.f7 = _rhs410;
              let _index431 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2567_v1).length));
              (_2567_v1)[_index431] = new BigNumber(768);
              let _2608_v18;
              _2608_v18 = _dafny.Map.Empty.slice().updateUnsafe((_2603_v16)[_module.__default.safeIndex(new BigNumber(989), new BigNumber((_2603_v16).length))],_2566_v0);
              let _2609_v19;
              _2609_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2566_v0,new BigNumber(653));
              let _2610_v20;
              _2610_v20 = _dafny.Seq.of(_2609_v19);
              let _2611_v21;
              _2611_v21 = _dafny.MultiSet.fromElements(_module.__default.fm0(new BigNumber(11), globalState));
              let _index432 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2567_v1).length));
              let _rhs411 = (_2608_v18).equals((_2608_v18).Merge(_2608_v18));
              let _rhs412 = (_2566_v0).minus(new BigNumber(911));
              let _rhs413 = (_this.f12) && (_this.f12);
              let _rhs414 = (new BigNumber(((_2610_v20)[_module.__default.safeIndex(new BigNumber((_2611_v21).cardinality()), new BigNumber((_2610_v20).length))]).length)).minus(_2566_v0);
              let _lhs297 = globalState;
              let _lhs298 = _this;
              let _lhs299 = _2567_v1;
              let _lhs300 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2567_v1).length));
              r1 = _rhs411;
              _lhs297.f0 = _rhs412;
              _lhs298.f12 = _rhs413;
              _lhs299[_lhs300] = _rhs414;
              _2576_v7 = _2576_v7;
              let _2612_v22;
              _2612_v22 = _dafny.MultiSet.fromElements((_2567_v1)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_2567_v1).length))]);
              let _2613_v23;
              let _nw398 = new _module.C11();
              _nw398.__ctor(new BigNumber((_2612_v22).cardinality()), _2575_v6, (_this).f10);
              _2613_v23 = _nw398;
              let _2614_v24;
              _2614_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2613_v23,(_this).f10);
              _2614_v24 = _2614_v24;
            }
            if ((_this).fm8(new BigNumber(((_2574_v5)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_2574_v5).length))]).length), (_this).f10, _dafny.Seq.IsPrefixOf(_module.__default.fm40((_this).f10, (_this).f10, globalState), _2571_v3), globalState)) {
              let _index433 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length));
              (_2567_v1)[_index433] = (_2566_v0).multipliedBy(_2566_v0);
              let _index434 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2567_v1).length));
              (_2567_v1)[_index434] = _2566_v0;
              let _index435 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length));
              let _index436 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2567_v1).length));
              let _rhs415 = _2566_v0;
              let _rhs416 = _2566_v0;
              let _lhs301 = _2567_v1;
              let _lhs302 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length));
              let _lhs303 = _2567_v1;
              let _lhs304 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2567_v1).length));
              _lhs301[_lhs302] = _rhs415;
              _lhs303[_lhs304] = _rhs416;
              let _2615_v25;
              let _nw399 = Array((new BigNumber(24)).toNumber()).fill([]);
              _2615_v25 = _nw399;
              let _rhs417 = _module.__default.safeDivisionInt((_2567_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length))], (_2567_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length))]);
              let _rhs418 = !((_this).f10);
              let _rhs419 = (((_this).f10) ? (_2615_v25) : (_2615_v25));
              let _lhs305 = _this;
              r2 = _rhs417;
              _lhs305.f12 = _rhs418;
              _2615_v25 = _rhs419;
              let _2616_v26;
              _2616_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-881),_2566_v0);
              (_this).f12 = ((_2567_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length))]).isLessThan((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_2566_v0, new BigNumber((_2616_v26).length))));
              let _2617_v27;
              _2617_v27 = _dafny.MultiSet.fromElements((_this).f10);
              let _2618_v28;
              _2618_v28 = _module.D15.create_DC45((_2567_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length))], (_this).f10, (_this).f11);
              let _2619_v29;
              let _nw400 = Array((new BigNumber(14)).toNumber());
              _nw400[(_dafny.ZERO).toNumber()] = _2618_v28;
              _nw400[(_dafny.ONE).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(2)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(3)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(4)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(5)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(6)).toNumber()] = _module.D15.create_DC45(new BigNumber(434), (_this).f10, _this.f12);
              _nw400[(new BigNumber(7)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(8)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(9)).toNumber()] = _module.__default.fm57(globalState);
              _nw400[(new BigNumber(10)).toNumber()] = _module.D15.create_DC45(new BigNumber(942), !((_this).f11), (_this).f11);
              _nw400[(new BigNumber(11)).toNumber()] = _module.__default.fm57(globalState);
              _nw400[(new BigNumber(12)).toNumber()] = _2618_v28;
              _nw400[(new BigNumber(13)).toNumber()] = _2618_v28;
              _2619_v29 = _nw400;
              let _index437 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_2619_v29).length));
              (_2619_v29)[_index437] = _2618_v28;
              let _2620_v30;
              _2620_v30 = _dafny.MultiSet.fromElements((_2567_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length))]);
              let _index438 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_2619_v29).length));
              let _index439 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length));
              let _rhs420 = _2617_v27;
              let _rhs421 = _2618_v28;
              let _rhs422 = (_2567_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length))];
              let _rhs423 = _2575_v6;
              let _rhs424 = _module.__default.fm0(new BigNumber((_2620_v30).cardinality()), globalState);
              let _lhs306 = _2619_v29;
              let _lhs307 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_2619_v29).length));
              let _lhs308 = _2567_v1;
              let _lhs309 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2567_v1).length));
              _2617_v27 = _rhs420;
              _lhs306[_lhs307] = _rhs421;
              _lhs308[_lhs309] = _rhs422;
              _2575_v6 = _rhs423;
              r1 = _rhs424;
              (_this).f12 = _this.f12;
            } else {
              let _rhs425 = _2566_v0;
              let _rhs426 = _2566_v0;
              let _lhs310 = globalState;
              let _lhs311 = globalState;
              _lhs310.f7 = _rhs425;
              _lhs311.f7 = _rhs426;
              (globalState).f1 = (_2566_v0).minus((_dafny.ZERO).minus(_2566_v0));
              r1 = (_this).f10;
              let _2621_v31;
              _2621_v31 = _dafny.MultiSet.fromElements(_2566_v0, _2566_v0, _2566_v0);
              let _2622_v32;
              let _nw401 = new _module.C6();
              _nw401.__ctor();
              _2622_v32 = _nw401;
              let _2623_v33;
              _2623_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2622_v32,(_this).f11);
              let _2624_v34;
              _2624_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2621_v31,_2623_v33);
              r2 = new BigNumber((_2624_v34).length);
              (_this).f12 = _this.f12;
            }
          }
        }
      }
      let _hi16 = _2566_v0;
      for (let _2625_i6 = _module.__default.fm1(globalState); _2625_i6.isLessThan(_hi16); _2625_i6 = _2625_i6.plus(_dafny.ONE)) {
        let _index440 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length));
        (_2567_v1)[_index440] = _2566_v0;
        let _index441 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length));
        (_2567_v1)[_index441] = _module.__default.fm1(globalState);
        if ((_this).f11) {
          (globalState).f7 = new BigNumber(-961);
          (globalState).f7 = _2566_v0;
          (_this).f12 = (_this.f12) || ((_this).f10);
          let _2626_v35;
          let _init64 = function (_2627_i7) {
            return (_dafny.MultiSet.fromElements(_module.D5.create_DC16(_dafny.MultiSet.fromElements((_this).f11)))).IsDisjointFrom(_dafny.MultiSet.fromElements(_module.D5.create_DC16(_dafny.MultiSet.fromElements((_this).f11, (_this).f10)), _module.D5.create_DC16(_dafny.MultiSet.fromElements((_this).f10))));
          };
          let _nw402 = Array((new BigNumber(7)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw402.length); _i0_64++) {
            _nw402[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2626_v35 = _nw402;
          let _index442 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2626_v35).length));
          (_2626_v35)[_index442] = (_this).f11;
          let _index443 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2626_v35).length));
          (_2626_v35)[_index443] = !(!(_2566_v0).isEqualTo(_2625_i6)) || ((_this).f10);
          let _2628_v36;
          _2628_v36 = _dafny.Set.fromElements((_2626_v35)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_2626_v35).length))], (_this).f11);
          (globalState).f7 = _module.__default.safeModuloInt((_2567_v1)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length))], new BigNumber((_2628_v36).length));
        } else {
          let _2629_v37;
          _2629_v37 = new _dafny.CodePoint('i'.codePointAt(0));
          let _2630_v38;
          let _nw403 = Array((new BigNumber(3)).toNumber());
          _nw403[(_dafny.ZERO).toNumber()] = _2629_v37;
          _nw403[(_dafny.ONE).toNumber()] = _2629_v37;
          _nw403[(new BigNumber(2)).toNumber()] = _2629_v37;
          _2630_v38 = _nw403;
          let _2631_v39;
          _2631_v39 = _module.D12.create_DC37(new BigNumber(-726), _2630_v38, _2629_v37);
          let _2632_v40;
          _2632_v40 = _module.D13.create_DC40(_2631_v39, (_this).f10, _2625_i6, (_2567_v1)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length))]);
          (globalState).f0 = (_2632_v40).dtor_cf70;
          let _2633_v41;
          _2633_v41 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f12),(_this).f11);
          let _2634_v42;
          let _nw404 = new _module.C7();
          _nw404.__ctor((((_2633_v41).contains(_dafny.Seq.of((_this).f11))) ? ((_2633_v41).get(_dafny.Seq.of((_this).f11))) : ((_this).f10)), false);
          _2634_v42 = _nw404;
          let _2635_v43;
          let _nw405 = new _module.C10();
          _nw405.__ctor(_2566_v0, (_this).f10);
          _2635_v43 = _nw405;
          let _2636_v44;
          _2636_v44 = _dafny.Set.fromElements((((_2634_v42).f19) ? ((_dafny.ZERO).minus(_2566_v0)) : (_2566_v0)), (_2635_v43).f15, _2625_i6, _2566_v0, (_2566_v0).minus((_dafny.ZERO).minus(_2566_v0)));
          _2636_v44 = ((((_this).f10) ? (_module.__default.fm26(globalState)) : (_2636_v44))).Union(_dafny.Set.fromElements((_2635_v43).f15));
          let _2637_v45;
          let _init65 = function (_2638_i8) {
            return (_this).f10;
          };
          let _nw406 = Array((new BigNumber(15)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw406.length); _i0_65++) {
            _nw406[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2637_v45 = _nw406;
          let _2639_v46;
          _2639_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(752),_2637_v45);
          let _2640_v47;
          let _nw407 = Array((new BigNumber(18)).toNumber());
          _nw407[(_dafny.ZERO).toNumber()] = _2637_v45;
          _nw407[(_dafny.ONE).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(2)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(3)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(4)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(5)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(6)).toNumber()] = (((_2639_v46).contains((_2567_v1)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length))])) ? ((_2639_v46).get((_2567_v1)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length))])) : (_2637_v45));
          _nw407[(new BigNumber(7)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(8)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(9)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(10)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(11)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(12)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(13)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(14)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(15)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(16)).toNumber()] = _2637_v45;
          _nw407[(new BigNumber(17)).toNumber()] = _2637_v45;
          _2640_v47 = _nw407;
          let _2641_v48;
          _2641_v48 = _module.D17.create_DC50(_2640_v47);
          let _2642_v49;
          let _nw408 = Array((new BigNumber(2)).toNumber());
          _nw408[(_dafny.ZERO).toNumber()] = _2641_v48;
          _nw408[(_dafny.ONE).toNumber()] = _2641_v48;
          _2642_v49 = _nw408;
          _2642_v49 = _2642_v49;
        }
        let _2643_v50;
        _2643_v50 = _dafny.Seq.UnicodeFromString("idhb");
        let _2644_v51;
        _2644_v51 = _dafny.MultiSet.fromElements(_2625_i6, (_dafny.ZERO).minus(new BigNumber((_2643_v50).length)), _2625_i6);
        let _2645_v52;
        let _nw409 = new _module.C6();
        _nw409.__ctor();
        _2645_v52 = _nw409;
        let _2646_v53;
        _2646_v53 = _dafny.Seq.of(((((_2644_v51).contains(_2566_v0)) ? ((_2644_v51).get(_2566_v0)) : ((_2567_v1)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length))]))).minus(new BigNumber((_dafny.Seq.of(_2645_v52)).length)), (_2567_v1)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_2567_v1).length))], _2625_i6);
        (globalState).f7 = (_2646_v53)[_module.__default.safeIndex(_2625_i6, new BigNumber((_2646_v53).length))];
        (_this).f12 = (_2571_v3)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_2643_v50, _2643_v50)).length), new BigNumber((_2571_v3).length))];
      }
      let _2647_v54;
      _2647_v54 = _dafny.Seq.UnicodeFromString("gjbsmqq");
      let _2648_v55;
      _2648_v55 = new _dafny.CodePoint('i'.codePointAt(0));
      let _2649_v56;
      _2649_v56 = _dafny.Seq.of(_dafny.Seq.update(_2647_v54, _module.__default.safeIndex(new BigNumber((_2647_v54).length), new BigNumber((_2647_v54).length)), _2648_v55));
      let _2650_v57;
      _2650_v57 = _dafny.MultiSet.fromElements(_2647_v54, (_2649_v56)[_module.__default.safeIndex(_2566_v0, new BigNumber((_2649_v56).length))]);
      let _2651_v58;
      let _2652_v59;
      let _out70;
      let _out71;
      let _outcollector31 = (_this).m6(_2566_v0, _this.f12, _2650_v57, globalState);
      _out70 = _outcollector31[0];
      _out71 = _outcollector31[1];
      _2651_v58 = _out70;
      _2652_v59 = _out71;
      let _hi17 = _module.__default.safeModuloInt(_2652_v59, _2652_v59);
      for (let _2653_i9 = new BigNumber(-382); _2653_i9.isLessThan(_hi17); _2653_i9 = _2653_i9.plus(_dafny.ONE)) {
        (_this).f12 = (_dafny.MultiSet.FromArray(((_2651_v58) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(730)), ((_2654_v0) => function (_2655_i10) {
          return _2654_v0;
        })(_2566_v0))) : (_dafny.Seq.of(_2652_v59))))).contains(_2653_i9);
        let _2656_v60;
        _2656_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_this.f12);
        _2656_v60 = (_2656_v60).update(_2651_v58, !(!(!((_this).f11))));
        if (_module.__default.fm0(_2652_v59, globalState)) {
          (_this).f12 = (_this).f11;
          let _2657_v61;
          _2657_v61 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-361)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2566_v0,(_this).f11)).length)).multipliedBy(_2653_i9));
          (globalState).f7 = (_2657_v61)[_module.__default.safeIndex(new BigNumber(375), new BigNumber((_2657_v61).length))];
          let _2658_v62;
          let _nw410 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
          _2658_v62 = _nw410;
          _2658_v62 = _2658_v62;
          let _2659_v63;
          let _init66 = ((_2660_v61, _2661_v0) => function (_2662_i11) {
            return !(new BigNumber((_2660_v61).length)).isEqualTo(_2661_v0);
          })(_2657_v61, _2566_v0);
          let _nw411 = Array((new BigNumber(2)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw411.length); _i0_66++) {
            _nw411[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2659_v63 = _nw411;
          let _index444 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2659_v63).length));
          (_2659_v63)[_index444] = !((_2566_v0).isLessThan(new BigNumber(875)));
          let _2663_v64;
          _2663_v64 = _dafny.Seq.of(_2659_v63);
          let _index445 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2659_v63).length));
          (_2659_v63)[_index445] = !_dafny.Seq.contains(_2663_v64, _2659_v63);
          _2651_v58 = (_2652_v59).isLessThan(new BigNumber(233));
        } else {
          let _2664_v65;
          _2664_v65 = _dafny.Seq.of(_2571_v3, _dafny.Seq.update(_dafny.Seq.of(_this.f12), _module.__default.safeIndex(_2566_v0, new BigNumber((_dafny.Seq.of(_this.f12)).length)), (_this).f11), _2571_v3);
          _2651_v58 = !_dafny.Seq.contains(_2664_v65, _dafny.Seq.Concat(_2571_v3, _2571_v3));
          let _2665_v66;
          let _nw412 = Array((new BigNumber(19)).toNumber()).fill(false);
          _2665_v66 = _nw412;
          let _index446 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length));
          (_2665_v66)[_index446] = (_this).f10;
          let _index447 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length));
          (_2665_v66)[_index447] = _2651_v58;
          let _2666_v67;
          _2666_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2648_v55,_dafny.Seq.update(_2647_v54, _module.__default.safeIndex(_2653_i9, new BigNumber((_2647_v54).length)), (_2647_v54)[_module.__default.safeIndex(_2653_i9, new BigNumber((_2647_v54).length))]));
          let _2667_v68;
          _2667_v68 = _dafny.Set.fromElements(_2651_v58, _2651_v58, false, (_this).f11);
          let _2668_v69;
          _2668_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2667_v68,_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), ((_2669_v55) => function (_2670_i16) {
            return _2669_v55;
          })(_2648_v55)));
          let _2671_v70;
          let _nw413 = Array((new BigNumber(24)).toNumber());
          _nw413[(_dafny.ZERO).toNumber()] = (((_2666_v67).contains(_2648_v55)) ? ((_2666_v67).get(_2648_v55)) : (_2647_v54));
          _nw413[(_dafny.ONE).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(2)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_2647_v54, _2647_v54);
          _nw413[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(354)), function (_2672_i12) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          });
          _nw413[(new BigNumber(5)).toNumber()] = (((_this).f10) ? (_dafny.Seq.UnicodeFromString("jcg")) : (_module.__default.fm25(_2566_v0, (_2665_v66)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length))], _2653_i9, globalState)));
          _nw413[(new BigNumber(6)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(7)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(8)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(9)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), ((_2673_v55) => function (_2674_i13) {
            return _2673_v55;
          })(_2648_v55));
          _nw413[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), ((_2675_v55) => function (_2676_i14) {
            return _2675_v55;
          })(_2648_v55));
          _nw413[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_2647_v54, _2647_v54);
          _nw413[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("whgxodd");
          _nw413[(new BigNumber(14)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_2647_v54, _2647_v54), _module.__default.safeIndex(_2566_v0, new BigNumber((_dafny.Seq.Concat(_2647_v54, _2647_v54)).length)), _2648_v55);
          _nw413[(new BigNumber(16)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(17)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_2647_v54, _2647_v54);
          _nw413[(new BigNumber(19)).toNumber()] = _2647_v54;
          _nw413[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), ((_2677_v55) => function (_2678_i15) {
            return _2677_v55;
          })(_2648_v55));
          _nw413[(new BigNumber(21)).toNumber()] = (((_2668_v69).contains(_2667_v68)) ? ((_2668_v69).get(_2667_v68)) : (_2647_v54));
          _nw413[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_2647_v54, _2647_v54);
          _nw413[(new BigNumber(23)).toNumber()] = _2647_v54;
          _2671_v70 = _nw413;
          let _2679_v71;
          let _nw414 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _2679_v71 = _nw414;
          let _2680_v73;
          _2680_v73 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ofs")).length));
          let _2681_v74;
          _2681_v74 = _dafny.Seq.of(_2680_v73);
          let _2682_v75;
          _2682_v75 = _dafny.Seq.of(_2681_v74);
          let _index448 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_2679_v71).length));
          (_2679_v71)[_index448] = function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of ((_2682_v75)[_module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_2682_v75).length))]).Elements) {
              let _2683_v72 = _compr_64;
              if (_dafny.Seq.contains((_2682_v75)[_module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_2682_v75).length))], _2683_v72)) {
                _coll64.push([_2683_v72,_dafny.Map.Empty.slice().updateUnsafe(_2570_v2,_module.D5.create_DC18(_module.D5.create_DC18(_module.D5.create_DC17((_2665_v66)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length))], _2648_v55))))]);
              }
            }
            return _coll64;
          }();
          let _index449 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2567_v1).length));
          (_2567_v1)[_index449] = _2566_v0;
          let _index450 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_2567_v1).length));
          (_2567_v1)[_index450] = new BigNumber((_2667_v68).length);
          let _2684_v76;
          _2684_v76 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2651_v58, (_this).f11),_2653_i9);
          let _2685_v77;
          _2685_v77 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),_2653_i9);
          let _2686_v78;
          _2686_v78 = _dafny.MultiSet.fromElements(_2651_v58);
          let _2687_v79;
          _2687_v79 = _module.D5.create_DC16((_2686_v78).update((_2665_v66)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length))], _module.__default.abs(_2652_v59)));
          let _2688_v80;
          _2688_v80 = _module.D5.create_DC18(_2687_v79);
          let _2689_v81;
          _2689_v81 = _module.D5.create_DC18(_2688_v80);
          let _2690_v82;
          _2690_v82 = _dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC31(_2651_v58, _2651_v58, _2652_v59),_2689_v81);
          let _2691_v83;
          _2691_v83 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(_2653_i9, _2685_v77, (_2665_v66)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length))], globalState),_2690_v82);
          let _2692_v84;
          _2692_v84 = _dafny.Set.fromElements(_module.__default.fm58(globalState));
          let _2693_v85;
          _2693_v85 = _module.D19.create_DC58(new BigNumber((_2680_v73).cardinality()), (_2665_v66)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length))]);
          let _2694_v86;
          _2694_v86 = _dafny.Seq.of(_2692_v84, _dafny.Set.fromElements(_2693_v85, _2693_v85));
          let _index451 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_2679_v71).length));
          let _index452 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2567_v1).length));
          let _index453 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_2567_v1).length));
          let _rhs427 = !(_2684_v76).contains(_2571_v3);
          let _rhs428 = _2671_v70;
          let _rhs429 = _2691_v83;
          let _rhs430 = new BigNumber((_2694_v86).length);
          let _rhs431 = ((_dafny.ZERO).minus(_2653_i9)).plus((_2566_v0).multipliedBy(_2652_v59));
          let _lhs312 = _2679_v71;
          let _lhs313 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_2679_v71).length));
          let _lhs314 = _2567_v1;
          let _lhs315 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2567_v1).length));
          let _lhs316 = _2567_v1;
          let _lhs317 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_2567_v1).length));
          _2651_v58 = _rhs427;
          _2671_v70 = _rhs428;
          _lhs312[_lhs313] = _rhs429;
          _lhs314[_lhs315] = _rhs430;
          _lhs316[_lhs317] = _rhs431;
          _2685_v77 = (_2685_v77).update(_this.f12, _2652_v59);
          let _2695_v87;
          _2695_v87 = _dafny.Seq.of(_2665_v66, _2665_v66, _2665_v66);
          let _2696_v88;
          let _nw415 = Array((new BigNumber(20)).toNumber());
          _nw415[(_dafny.ZERO).toNumber()] = _2648_v55;
          _nw415[(_dafny.ONE).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(2)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
          _nw415[(new BigNumber(4)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(5)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(6)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(7)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(8)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(9)).toNumber()] = _module.__default.fm4(_this.f12, (_2665_v66)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2665_v66).length))], false, globalState);
          _nw415[(new BigNumber(10)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(11)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(12)).toNumber()] = ((_2651_v58) ? (_2648_v55) : (_2648_v55));
          _nw415[(new BigNumber(13)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(14)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(15)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(16)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(17)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(18)).toNumber()] = _2648_v55;
          _nw415[(new BigNumber(19)).toNumber()] = _2648_v55;
          _2696_v88 = _nw415;
          let _index454 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_2696_v88).length));
          (_2696_v88)[_index454] = _2648_v55;
          let _index455 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_2696_v88).length));
          let _rhs432 = _2695_v87;
          let _rhs433 = _2651_v58;
          let _rhs434 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gim"), _module.__default.fm13((_this).f11, (_this).f10, globalState)), _2647_v54);
          let _rhs435 = _2648_v55;
          let _lhs318 = _2696_v88;
          let _lhs319 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_2696_v88).length));
          _2695_v87 = _rhs432;
          r1 = _rhs433;
          _2647_v54 = _rhs434;
          _lhs318[_lhs319] = _rhs435;
        }
        let _2697_v89;
        let _nw416 = new _module.C0();
        _nw416.__ctor();
        _2697_v89 = _nw416;
      }
      let _pat_let_tv70 = _2566_v0;
      let _pat_let_tv71 = _2566_v0;
      let _pat_let_tv72 = _2566_v0;
      (globalState).f0 = function (_source42) {
        if (_source42.is_DC31) {
          let _2698___mcc_h7 = (_source42).cf50;
          let _2699___mcc_h8 = (_source42).cf51;
          let _2700___mcc_h9 = (_source42).cf52;
          let _2701_cf52 = _2700___mcc_h9;
          let _2702_cf51 = _2699___mcc_h8;
          let _2703_cf50 = _2698___mcc_h7;
          return (_dafny.ZERO).minus((new BigNumber(886)).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(353), _pat_let_tv70, _pat_let_tv71, _pat_let_tv72)).length)));
        } else {
          let _2704___mcc_h10 = (_source42).cf49;
          let _2705_cf49 = _2704___mcc_h10;
          return new BigNumber(154);
        }
      }(function (_pat_let64_0) {
        return function (_2706_dt__update__tmp_h0) {
          return function (_pat_let65_0) {
            return function (_2707_dt__update_hcf50_h0) {
              return _module.D10.create_DC31(_2707_dt__update_hcf50_h0, (_2706_dt__update__tmp_h0).dtor_cf51, (_2706_dt__update__tmp_h0).dtor_cf52);
            }(_pat_let65_0);
          }(_this.f12);
        }(_pat_let64_0);
      }(_2570_v2));
      let _2708_v90;
      let _nw417 = new _module.C1();
      _nw417.__ctor(_2651_v58);
      _2708_v90 = _nw417;
      let _2709_v91;
      _2709_v91 = _dafny.MultiSet.fromElements((_this).f10);
      let _rhs436 = _2708_v90;
      let _rhs437 = (_2709_v91).IsSubsetOf(_2709_v91);
      _2708_v90 = _rhs436;
      _2651_v58 = _rhs437;
      let _2710_v92;
      _2710_v92 = _dafny.Set.fromElements(_2647_v54);
      r0 = ((_2710_v92).Union(_2710_v92)).Difference((_2710_v92).Difference(_2710_v92));
      let _2711_v93;
      _2711_v93 = _dafny.MultiSet.fromElements(new BigNumber((_2571_v3).length));
      let _2712_v94;
      let _nw418 = Array((new BigNumber(2)).toNumber());
      _nw418[(_dafny.ZERO).toNumber()] = !(_this.f12);
      _nw418[(_dafny.ONE).toNumber()] = _2651_v58;
      _2712_v94 = _nw418;
      let _2713_v95;
      _2713_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2571_v3).length),_2712_v94);
      r1 = !(_module.__default.fm0(_module.__default.safeDivisionInt((((_2711_v93).contains(new BigNumber((_2713_v95).length))) ? ((_2711_v93).get(new BigNumber((_2713_v95).length))) : (new BigNumber((_2571_v3).length))), _2652_v59), globalState));
      r2 = _module.__default.fm1(globalState);
      return [r0, r1, r2];
    }
    m5(p0, globalState) {
      let _this = this;
      if ((_this).f10) {
        (globalState).f0 = p0;
        let _2714_v0;
        let _nw419 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _2714_v0 = _nw419;
        let _index456 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length));
        (_2714_v0)[_index456] = (p0).minus(p0);
        let _2715_v1;
        _2715_v1 = _dafny.Set.fromElements(_this.f12, (_this).f11);
        let _2716_v2;
        _2716_v2 = _module.D1.create_DC3(p0);
        let _2717_v3;
        _2717_v3 = _dafny.Set.fromElements(p0, (_2716_v2).dtor_cf4);
        let _index457 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length));
        (_2714_v0)[_index457] = (new BigNumber((_2715_v1).length)).multipliedBy(new BigNumber(((_2717_v3).Intersect(_module.__default.fm31(p0, globalState))).length));
        let _2718_v4;
        _2718_v4 = _module.D6.create_DC20((_this).f11, _this.f12, (_this).f11, _module.__default.safeDivisionInt(p0, (_2714_v0)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length))]));
        let _2719_v5;
        _2719_v5 = new _dafny.CodePoint('a'.codePointAt(0));
        let _2720_v6;
        _2720_v6 = _dafny.MultiSet.fromElements(p0);
        let _pat_let_tv73 = p0;
        let _index458 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length));
        let _rhs438 = p0;
        let _rhs439 = function (_pat_let66_0) {
          return function (_2721_dt__update__tmp_h0) {
            return function (_pat_let67_0) {
              return function (_2722_dt__update_hcf32_h0) {
                return function (_pat_let68_0) {
                  return function (_2723_dt__update_hcf35_h0) {
                    return function (_pat_let69_0) {
                      return function (_2724_dt__update_hcf34_h0) {
                        return _module.D6.create_DC20(_2722_dt__update_hcf32_h0, (_2721_dt__update__tmp_h0).dtor_cf33, _2724_dt__update_hcf34_h0, _2723_dt__update_hcf35_h0);
                      }(_pat_let69_0);
                    }(true);
                  }(_pat_let68_0);
                }(_pat_let_tv73);
              }(_pat_let67_0);
            }(_this.f12);
          }(_pat_let66_0);
        }(_module.D6.create_DC20((_this).f11, (_this).f11, (_module.D5.create_DC17(_this.f12, _2719_v5)).dtor_cf28, new BigNumber((_2720_v6).cardinality())));
        let _rhs440 = _this.f12;
        let _rhs441 = (((_this.f12) ? ((_2714_v0)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length))]) : ((_2714_v0)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length))]))).minus((_2714_v0)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length))]);
        let _lhs320 = globalState;
        let _lhs321 = _this;
        let _lhs322 = _2714_v0;
        let _lhs323 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length));
        _lhs320.f7 = _rhs438;
        _2718_v4 = _rhs439;
        _lhs321.f12 = _rhs440;
        _lhs322[_lhs323] = _rhs441;
        let _2725_v7;
        let _nw420 = Array((new BigNumber(11)).toNumber()).fill([]);
        _2725_v7 = _nw420;
        let _2726_v8;
        let _nw421 = Array((new BigNumber(3)).toNumber()).fill(false);
        _2726_v8 = _nw421;
        let _index459 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2725_v7).length));
        (_2725_v7)[_index459] = _2726_v8;
        let _index460 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2725_v7).length));
        (_2725_v7)[_index460] = _2726_v8;
        let _2727_v9;
        let _nw422 = new _module.C10();
        _nw422.__ctor(((_2714_v0)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_2714_v0).length))]).multipliedBy(p0), _this.f12);
        _2727_v9 = _nw422;
      } else {
        let _2728_v10;
        _2728_v10 = _dafny.Seq.UnicodeFromString("pbgo");
        let _2729_v11;
        let _nw423 = new _module.C3();
        _nw423.__ctor(false, _this.f12);
        _2729_v11 = _nw423;
        let _2730_v12;
        _2730_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2729_v11,p0);
        let _2731_v13;
        let _nw424 = Array((new BigNumber(18)).toNumber());
        _nw424[(_dafny.ZERO).toNumber()] = new BigNumber(965);
        _nw424[(_dafny.ONE).toNumber()] = p0;
        _nw424[(new BigNumber(2)).toNumber()] = new BigNumber((_2728_v10).length);
        _nw424[(new BigNumber(3)).toNumber()] = _module.__default.fm1(globalState);
        _nw424[(new BigNumber(4)).toNumber()] = p0;
        _nw424[(new BigNumber(5)).toNumber()] = new BigNumber((_2728_v10).length);
        _nw424[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw424[(new BigNumber(7)).toNumber()] = p0;
        _nw424[(new BigNumber(8)).toNumber()] = p0;
        _nw424[(new BigNumber(9)).toNumber()] = p0;
        _nw424[(new BigNumber(10)).toNumber()] = new BigNumber(768);
        _nw424[(new BigNumber(11)).toNumber()] = p0;
        _nw424[(new BigNumber(12)).toNumber()] = p0;
        _nw424[(new BigNumber(13)).toNumber()] = new BigNumber((_2728_v10).length);
        _nw424[(new BigNumber(14)).toNumber()] = new BigNumber((_2730_v12).length);
        _nw424[(new BigNumber(15)).toNumber()] = p0;
        _nw424[(new BigNumber(16)).toNumber()] = new BigNumber(-624);
        _nw424[(new BigNumber(17)).toNumber()] = p0;
        _2731_v13 = _nw424;
        let _2732_v14;
        _2732_v14 = _dafny.Seq.of(_2731_v13, _2731_v13);
        let _2733_v15;
        _2733_v15 = _dafny.Map.Empty.slice().updateUnsafe((_2732_v14)[_module.__default.safeIndex(p0, new BigNumber((_2732_v14).length))],_module.__default.fm0(p0, globalState));
        _2733_v15 = (_2733_v15).update(_2731_v13, (_this).f11);
        (globalState).f1 = (((_2729_v11).f20) ? (_module.__default.safeDivisionInt(new BigNumber(-29), p0)) : (p0));
        (_this).f12 = !((p0).isLessThan(p0)) || (!((_2729_v11).f20));
        let _2734_v16;
        let _init67 = function (_2735_i0) {
          return _this.f12;
        };
        let _nw425 = Array((new BigNumber(10)).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw425.length); _i0_67++) {
          _nw425[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2734_v16 = _nw425;
        let _2736_v17;
        _2736_v17 = _dafny.MultiSet.fromElements(_2734_v16);
        let _rhs442 = _2736_v17;
        let _rhs443 = p0;
        let _rhs444 = (_this).f10;
        let _lhs324 = globalState;
        let _lhs325 = _this;
        _2736_v17 = _rhs442;
        _lhs324.f0 = _rhs443;
        _lhs325.f12 = _rhs444;
        let _2737_v18;
        let _init68 = ((_2738_p0, _2739_v10) => function (_2740_i1) {
          return _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2738_p0), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_2739_v10).length))).length));
        })(p0, _2728_v10);
        let _nw426 = Array((new BigNumber(21)).toNumber());
        for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw426.length); _i0_68++) {
          _nw426[_i0_68] = _init68(new BigNumber(_i0_68));
        }
        _2737_v18 = _nw426;
        let _2741_v19;
        _2741_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_2729_v11).f20) ? (_2737_v18) : (_2737_v18)),_2731_v13);
        _2741_v19 = (_2741_v19).update(_2737_v18, _2731_v13);
      }
      let _2742_v20;
      _2742_v20 = _dafny.Seq.of((_this).f11, _this.f12);
      let _2743_v21;
      _2743_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2742_v20,p0);
      let _hi18 = new BigNumber((_2743_v21).length);
      for (let _2744_i2 = p0; _2744_i2.isLessThan(_hi18); _2744_i2 = _2744_i2.plus(_dafny.ONE)) {
        let _2745_v22;
        let _init69 = ((_2746_p0) => function (_2747_i3) {
          return (_2747_i3).plus(_2746_p0);
        })(p0);
        let _nw427 = Array((new BigNumber(20)).toNumber());
        for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw427.length); _i0_69++) {
          _nw427[_i0_69] = _init69(new BigNumber(_i0_69));
        }
        _2745_v22 = _nw427;
        let _2748_v23;
        _2748_v23 = _dafny.MultiSet.fromElements(_2744_i2);
        let _2749_v24;
        _2749_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_2748_v23);
        let _index461 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_2745_v22).length));
        (_2745_v22)[_index461] = new BigNumber((_2749_v24).length);
        let _index462 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_2745_v22).length));
        (_2745_v22)[_index462] = _2744_i2;
        let _2750_v25;
        let _init70 = function (_2751_i4) {
          return (_this).f10;
        };
        let _nw428 = Array((new BigNumber(10)).toNumber());
        for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw428.length); _i0_70++) {
          _nw428[_i0_70] = _init70(new BigNumber(_i0_70));
        }
        _2750_v25 = _nw428;
        let _rhs445 = (_module.__default.safeModuloInt(_2744_i2, p0)).minus(p0);
        let _rhs446 = _2750_v25;
        let _lhs326 = globalState;
        _lhs326.f0 = _rhs445;
        _2750_v25 = _rhs446;
        let _2752_v26;
        let _nw429 = new _module.C0();
        _nw429.__ctor();
        _2752_v26 = _nw429;
        let _2753_v27;
        let _nw430 = new _module.C1();
        _nw430.__ctor(!(_this.f12) || ((_this).f10));
        _2753_v27 = _nw430;
      }
      if (_this.f12) {
        let _2754_v28;
        _2754_v28 = _dafny.Seq.UnicodeFromString("vgpnm");
        let _2755_v29;
        _2755_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(_2754_v28, (_this).f11, (_this).f10),p0);
        let _2756_v30;
        _2756_v30 = _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("qlowkigp"), (_this).f11, true);
        _2755_v29 = (_2755_v29).update(_2756_v30, new BigNumber((_2754_v28).length));
        _2754_v28 = _2754_v28;
        let _2757_v31;
        let _init71 = function (_2758_i5) {
          return (_this).f11;
        };
        let _nw431 = Array((new BigNumber(29)).toNumber());
        for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw431.length); _i0_71++) {
          _nw431[_i0_71] = _init71(new BigNumber(_i0_71));
        }
        _2757_v31 = _nw431;
        let _index463 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_2757_v31).length));
        (_2757_v31)[_index463] = _this.f12;
        let _index464 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_2757_v31).length));
        (_2757_v31)[_index464] = _this.f12;
        let _2759_v32;
        _2759_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_module.__default.safeModuloInt(p0, new BigNumber((_2754_v28).length)));
        _2759_v32 = _2759_v32;
        let _index465 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_2757_v31).length));
        (_2757_v31)[_index465] = (p0).isLessThan(p0);
      } else {
        let _2760_v33;
        _2760_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-724),p0);
        let _2761_v34;
        _2761_v34 = _dafny.Seq.of(new BigNumber(955));
        let _2762_v35;
        _2762_v35 = _dafny.Set.fromElements((_this).f11, (_this).f11);
        let _2763_v36;
        _2763_v36 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(419)), function (_2764_i7) {
          return new BigNumber(541);
        }));
        let _2765_v37;
        _2765_v37 = new _dafny.CodePoint('m'.codePointAt(0));
        let _2766_v38;
        _2766_v38 = _dafny.Seq.of(_2765_v37, _2765_v37);
        let _2767_v39;
        let _nw432 = Array((new BigNumber(19)).toNumber());
        _nw432[(_dafny.ZERO).toNumber()] = new BigNumber(648);
        _nw432[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,p0)).length);
        _nw432[(new BigNumber(2)).toNumber()] = p0;
        _nw432[(new BigNumber(3)).toNumber()] = p0;
        _nw432[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw432[(new BigNumber(5)).toNumber()] = p0;
        _nw432[(new BigNumber(6)).toNumber()] = p0;
        _nw432[(new BigNumber(7)).toNumber()] = p0;
        _nw432[(new BigNumber(8)).toNumber()] = p0;
        _nw432[(new BigNumber(9)).toNumber()] = p0;
        _nw432[(new BigNumber(10)).toNumber()] = new BigNumber((_2760_v33).length);
        _nw432[(new BigNumber(11)).toNumber()] = p0;
        _nw432[(new BigNumber(12)).toNumber()] = new BigNumber(208);
        _nw432[(new BigNumber(13)).toNumber()] = (((_2760_v33).contains(new BigNumber((_module.__default.fm45(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), ((_2772_p0) => function (_2773_i6) {
          return _dafny.Seq.of(_2772_p0, _2772_p0);
        })(p0)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), ((_2774_p0) => function (_2775_i6) {
          return _dafny.Seq.of(_2774_p0, _2774_p0);
        })(p0))).length)), _2761_v34), p0, globalState)).length))) ? ((_2760_v33).get(new BigNumber((_module.__default.fm45(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), ((_2768_p0) => function (_2769_i6) {
          return _dafny.Seq.of(_2768_p0, _2768_p0);
        })(p0)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), ((_2770_p0) => function (_2771_i6) {
          return _dafny.Seq.of(_2770_p0, _2770_p0);
        })(p0))).length)), _2761_v34), p0, globalState)).length))) : (new BigNumber(541)));
        _nw432[(new BigNumber(14)).toNumber()] = p0;
        _nw432[(new BigNumber(15)).toNumber()] = p0;
        _nw432[(new BigNumber(16)).toNumber()] = p0;
        _nw432[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2762_v35).length), p0);
        _nw432[(new BigNumber(18)).toNumber()] = (new BigNumber((_2763_v36).cardinality())).plus(new BigNumber((_2766_v38).length));
        _2767_v39 = _nw432;
        let _index466 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_2767_v39).length));
        (_2767_v39)[_index466] = new BigNumber(571);
        let _index467 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_2767_v39).length));
        (_2767_v39)[_index467] = p0;
        (_this).f12 = false;
        (_this).f12 = (_this).f11;
        let _2776_v40;
        let _nw433 = new _module.C3();
        _nw433.__ctor(_this.f12, _this.f12);
        _2776_v40 = _nw433;
        let _2777_v41;
        let _nw434 = Array((new BigNumber(2)).toNumber()).fill(false);
        _2777_v41 = _nw434;
        let _index468 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_2777_v41).length));
        (_2777_v41)[_index468] = (_2776_v40).f20;
        let _2778_v42;
        _2778_v42 = _module.D17.create_DC51(_2765_v37, (_2776_v40).f20);
        let _2779_v43;
        _2779_v43 = _dafny.MultiSet.fromElements(_2778_v42, _2778_v42, _2778_v42);
        let _index469 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_2777_v41).length));
        (_2777_v41)[_index469] = ((_2779_v43).Union(_2779_v43)).IsSubsetOf(_2779_v43);
      }
      let _2780_v44;
      _2780_v44 = _module.D15.create_DC46();
      let _source43 = _2780_v44;
      if (_source43.is_DC45) {
        let _2781___mcc_h0 = (_source43).cf80;
        let _2782___mcc_h1 = (_source43).cf81;
        let _2783___mcc_h2 = (_source43).cf82;
        let _2784_cf82 = _2783___mcc_h2;
        let _2785_cf81 = _2782___mcc_h1;
        let _2786_cf80 = _2781___mcc_h0;
        _2784_cf82 = _2785_cf81;
        let _2787_v45;
        _2787_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2785_cf81,_2786_cf80);
        let _2788_v46;
        let _nw435 = new _module.C7();
        _nw435.__ctor(_this.f12, !(_2787_v45).contains(false));
        _2788_v46 = _nw435;
        (globalState).f1 = new BigNumber(746);
        _2785_cf81 = _this.f12;
      } else if (_source43.is_DC46) {
        let _2789_v47;
        let _nw436 = Array((new BigNumber(9)).toNumber()).fill(false);
        _2789_v47 = _nw436;
        let _2790_v48;
        let _2791_v49;
        let _out72;
        let _out73;
        let _outcollector32 = _module.__default.m0(p0, _2789_v47, globalState);
        _out72 = _outcollector32[0];
        _out73 = _outcollector32[1];
        _2790_v48 = _out72;
        _2791_v49 = _out73;
        if ((((_this).f10) ? ((_this).f10) : ((_2742_v20)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_2742_v20).length))]))) {
          let _2792_v50;
          let _nw437 = new _module.C2();
          _nw437.__ctor((_this).f11, p0);
          _2792_v50 = _nw437;
          let _2793_v51;
          _2793_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f11, (_this).f11)).length),_2792_v50);
          let _2794_v52;
          _2794_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2792_v50).f22);
          _2793_v51 = (_2793_v51).update(new BigNumber((_2794_v52).length), _2792_v50);
          let _2795_v53;
          let _nw438 = new _module.C3();
          _nw438.__ctor(true, (_this).fm10(globalState));
          _2795_v53 = _nw438;
          let _2796_v54;
          let _nw439 = Array((new BigNumber(13)).toNumber());
          _nw439[(_dafny.ZERO).toNumber()] = _2795_v53;
          _nw439[(_dafny.ONE).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(2)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(3)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(4)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(5)).toNumber()] = ((_this.f12) ? (_2795_v53) : (_2795_v53));
          _nw439[(new BigNumber(6)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(7)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(8)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(9)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(10)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(11)).toNumber()] = _2795_v53;
          _nw439[(new BigNumber(12)).toNumber()] = _2795_v53;
          _2796_v54 = _nw439;
          let _index470 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2796_v54).length));
          (_2796_v54)[_index470] = _2795_v53;
          let _2797_v55;
          _2797_v55 = _module.D21.create_DC61(_2795_v53);
          let _index471 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2796_v54).length));
          (_2796_v54)[_index471] = ((!(new BigNumber(423)).isEqualTo((_dafny.ZERO).minus((_2792_v50).f22))) ? ((((_2792_v50).f21) ? (_2795_v53) : ((_2797_v55).dtor_cf103))) : (_2795_v53));
          let _index472 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2789_v47).length));
          (_2789_v47)[_index472] = (_this).f10;
          let _index473 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2789_v47).length));
          (_2789_v47)[_index473] = (((_2792_v50).f21) ? ((_2742_v20)[_module.__default.safeIndex(p0, new BigNumber((_2742_v20).length))]) : ((_2792_v50).f21));
          let _2798_v56;
          let _nw440 = new _module.C9();
          _nw440.__ctor((_this).f10, (_this).f11);
          _2798_v56 = _nw440;
          let _2799_v57;
          _2799_v57 = _dafny.Map.Empty.slice().updateUnsafe(_2790_v48,(_2798_v56).f17);
          _2799_v57 = (_2799_v57).update(_2791_v49, (_2798_v56).f17);
        } else {
          (globalState).f1 = _module.__default.safeModuloInt(p0, p0);
          let _2800_v58;
          let _nw441 = new _module.C2();
          _nw441.__ctor(!((_2742_v20)[_module.__default.safeIndex(p0, new BigNumber((_2742_v20).length))]), p0);
          _2800_v58 = _nw441;
          let _2801_v59;
          let _nw442 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _2801_v59 = _nw442;
          let _nw443 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2801_v59 = _nw443;
          (globalState).f0 = new BigNumber((_dafny.Seq.UnicodeFromString("qe")).length);
          let _2802_v60;
          let _nw444 = new _module.C9();
          _nw444.__ctor((_2742_v20)[_module.__default.safeIndex((_dafny.ZERO).minus((_2800_v58).f22), new BigNumber((_2742_v20).length))], (_this).f10);
          _2802_v60 = _nw444;
        }
        let _2803_v61;
        _2803_v61 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2804_v62;
        _2804_v62 = _module.D2.create_DC6(_2803_v61);
        let _2805_v63;
        let _nw445 = Array((new BigNumber(17)).toNumber());
        _nw445[(_dafny.ZERO).toNumber()] = _2803_v61;
        _nw445[(_dafny.ONE).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw445[(new BigNumber(3)).toNumber()] = _module.__default.fm4(_this.f12, (_this).f10, _this.f12, globalState);
        _nw445[(new BigNumber(4)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(5)).toNumber()] = (_2791_v49)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2791_v49).length))];
        _nw445[(new BigNumber(6)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(7)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(8)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(9)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(10)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(11)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
        _nw445[(new BigNumber(13)).toNumber()] = (_2804_v62).dtor_cf8;
        _nw445[(new BigNumber(14)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(15)).toNumber()] = _2803_v61;
        _nw445[(new BigNumber(16)).toNumber()] = _module.__default.fm4((_this).f10, (_this).f11, (_this).f11, globalState);
        _2805_v63 = _nw445;
        let _index474 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2805_v63).length));
        (_2805_v63)[_index474] = _2803_v61;
        let _index475 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2805_v63).length));
        (_2805_v63)[_index475] = _2803_v61;
        let _index476 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_2789_v47).length));
        (_2789_v47)[_index476] = (_this).f10;
        let _index477 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_2789_v47).length));
        (_2789_v47)[_index477] = !(!(true));
      } else if (_source43.is_DC44) {
        let _2806___mcc_h3 = (_source43).cf79;
        let _2807_cf79 = _2806___mcc_h3;
        let _2808_v64;
        let _nw446 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _2808_v64 = _nw446;
        let _2809_v65;
        _2809_v65 = _module.D21.create_DC63(p0, (_this).f10, (_this).f10, (_this).f11);
        let _rhs447 = ((_2809_v65).dtor_cf108) === ((_this).f11);
        let _rhs448 = _2808_v64;
        let _lhs327 = _this;
        _lhs327.f12 = _rhs447;
        _2808_v64 = _rhs448;
        let _2810_v66;
        let _nw447 = new _module.C5();
        _nw447.__ctor();
        _2810_v66 = _nw447;
        let _2811_v67;
        _2811_v67 = _dafny.Seq.of((_dafny.ZERO).minus(p0));
        let _2812_v68;
        _2812_v68 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f12);
        let _2813_v69;
        _2813_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2812_v68,p0);
        let _2814_v70;
        _2814_v70 = new _dafny.CodePoint('i'.codePointAt(0));
        let _2815_v71;
        _2815_v71 = _dafny.Seq.of(_2814_v70);
        let _2816_v72;
        _2816_v72 = _dafny.MultiSet.fromElements((((_2813_v69).contains(_2812_v68)) ? ((_2813_v69).get(_2812_v68)) : (p0)), p0, p0, new BigNumber((_2815_v71).length), p0);
        let _2817_v73;
        _2817_v73 = _module.D4.create_DC15(((!((_this).f11)) ? (_module.__default.fm1(globalState)) : (new BigNumber((_2811_v67).length))), (_2816_v72).update(p0, _module.__default.abs(p0)), (_2815_v71)[_module.__default.safeIndex(p0, new BigNumber((_2815_v71).length))], (_dafny.ZERO).minus(p0));
        let _source44 = _2817_v73;
        if (_source44.is_DC14) {
          let _2818___mcc_h5 = (_source44).cf21;
          let _2819___mcc_h6 = (_source44).cf22;
          let _2820_cf22 = _2819___mcc_h6;
          let _2821_cf21 = _2818___mcc_h5;
          let _2822_v74;
          let _nw448 = new _module.C2();
          _nw448.__ctor(_this.f12, new BigNumber((_2812_v68).length));
          _2822_v74 = _nw448;
          let _2823_v75;
          _2823_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2820_cf22,!(_this.f12) || (_this.f12));
          let _rhs449 = _2822_v74;
          let _rhs450 = _2815_v71;
          let _rhs451 = _2823_v75;
          let _rhs452 = _dafny.Seq.of(_module.__default.fm0(new BigNumber((_2815_v71).length), globalState), (_2822_v74).f21, !((_2822_v74).f21) || (!(_this.f12)));
          let _rhs453 = _dafny.Seq.Concat(((!((_2822_v74).f21)) ? (_2811_v67) : (_2811_v67)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-929)), ((_2824_v74) => function (_2825_i8) {
            return (_2824_v74).f22;
          })(_2822_v74)));
          _2822_v74 = _rhs449;
          _2815_v71 = _rhs450;
          _2823_v75 = _rhs451;
          _2742_v20 = _rhs452;
          _2811_v67 = _rhs453;
          let _2826_v76;
          _2826_v76 = _dafny.Set.fromElements(_2812_v68, _2812_v68);
          _2812_v68 = (_2812_v68).update((_this).fm10(globalState), (_dafny.MultiSet.fromElements(new BigNumber((_2742_v20).length))).IsSubsetOf((_dafny.MultiSet.FromArray(_module.__default.fm3(_2821_cf21, new BigNumber(469), globalState))).update((_dafny.ZERO).minus((_2822_v74).fm9(_2815_v71, _2826_v76, globalState)), _module.__default.abs((_this).fm9(_2815_v71, _2826_v76, globalState)))));
          let _2827_v77;
          _2827_v77 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2742_v20);
          _2827_v77 = (_2827_v77).update(_2820_cf22, (((_2827_v77).contains(_module.__default.fm1(globalState))) ? ((_2827_v77).get(_module.__default.fm1(globalState))) : (_module.__default.fm40(!(_this.f12), (_this).f10, globalState))));
          let _2828_v78;
          _2828_v78 = _module.D4.create_DC14(_2820_cf22, _2820_cf22);
          (globalState).f7 = (_2828_v78).dtor_cf22;
        } else if (_source44.is_DC15) {
          let _2829___mcc_h7 = (_source44).cf23;
          let _2830___mcc_h8 = (_source44).cf24;
          let _2831___mcc_h9 = (_source44).cf25;
          let _2832___mcc_h10 = (_source44).cf26;
          let _2833_cf26 = _2832___mcc_h10;
          let _2834_cf25 = _2831___mcc_h9;
          let _2835_cf24 = _2830___mcc_h8;
          let _2836_cf23 = _2829___mcc_h7;
          let _2837_v79;
          let _nw449 = new _module.C5();
          _nw449.__ctor();
          _2837_v79 = _nw449;
          (globalState).f1 = _2836_cf23;
          (globalState).f7 = _2833_cf26;
          let _2838_v80;
          let _nw450 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2838_v80 = _nw450;
          let _index478 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_2838_v80).length));
          (_2838_v80)[_index478] = _2808_v64;
          let _index479 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_2838_v80).length));
          (_2838_v80)[_index479] = _2808_v64;
        } else {
          let _2839___mcc_h11 = (_source44).cf20;
          let _2840_cf20 = _2839___mcc_h11;
          (globalState).f1 = (_dafny.ZERO).minus(p0);
          _2808_v64 = _2808_v64;
          let _2841_v81;
          let _nw451 = Array((new BigNumber(29)).toNumber()).fill([]);
          _2841_v81 = _nw451;
          let _2842_v82;
          _2842_v82 = _module.D8.create_DC25(_2841_v81);
          let _2843_v83;
          _2843_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(((_this).f11) ? (_2842_v82) : (_2842_v82)));
          let _2844_v84;
          _2844_v84 = _dafny.MultiSet.fromElements(false);
          let _2845_v85;
          _2845_v85 = _dafny.MultiSet.fromElements(_2844_v84);
          _2843_v83 = (_2843_v83).update((_2845_v85).IsSubsetOf(_2845_v85), _module.D8.create_DC25(_2841_v81));
          let _rhs454 = (_this).fm8(p0, (_this).f11, _this.f12, globalState);
          let _rhs455 = false;
          let _rhs456 = p0;
          let _lhs328 = _this;
          let _lhs329 = _this;
          let _lhs330 = globalState;
          _lhs328.f12 = _rhs454;
          _lhs329.f12 = _rhs455;
          _lhs330.f1 = _rhs456;
        }
        let _2846_v86;
        _2846_v86 = _dafny.MultiSet.fromElements(_2814_v70, _2814_v70, _2814_v70, new _dafny.CodePoint('v'.codePointAt(0)));
        let _2847_v87;
        _2847_v87 = _dafny.Map.Empty.slice().updateUnsafe(_2846_v86,_2812_v68);
        let _2848_v88;
        _2848_v88 = _dafny.Set.fromElements(_2812_v68, _2812_v68);
        let _2849_v90;
        _2849_v90 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("able"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("able")).length)), _module.__default.fm4((_this).f11, _this.f12, true, globalState)),_2808_v64);
        let _rhs457 = (_this).fm9(_2815_v71, function () {
          let _coll65 = new _dafny.Set();
          for (const _compr_65 of (_2848_v88).Elements) {
            let _2850_v89 = _compr_65;
            if ((_2848_v88).contains(_2850_v89)) {
              _coll65.add(_2850_v89);
            }
          }
          return _coll65;
        }(), globalState);
        let _rhs458 = p0;
        let _rhs459 = (((_2849_v90).contains(_2815_v71)) ? ((_2849_v90).get(_2815_v71)) : (_2808_v64));
        let _rhs460 = (_this).f11;
        let _rhs461 = _2847_v87;
        let _lhs331 = globalState;
        let _lhs332 = globalState;
        let _lhs333 = _this;
        _lhs331.f0 = _rhs457;
        _lhs332.f7 = _rhs458;
        _2808_v64 = _rhs459;
        _lhs333.f12 = _rhs460;
        _2847_v87 = _rhs461;
      } else {
        let _2851___mcc_h4 = (_source43).cf83;
        let _2852_cf83 = _2851___mcc_h4;
        (_this).f12 = !((_this).f10);
        (_this).f12 = (_this).f11;
        if (!(new BigNumber(656)).isEqualTo(p0)) {
          let _2853_v91;
          _2853_v91 = new _dafny.CodePoint('n'.codePointAt(0));
          let _2854_v92;
          _2854_v92 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_2853_v91);
          let _2855_v93;
          _2855_v93 = _module.D6.create_DC21((_this).f10, (_this).f10, (_this).f11, p0);
          let _2856_v94;
          let _nw452 = Array((new BigNumber(22)).toNumber());
          _nw452[(_dafny.ZERO).toNumber()] = _2853_v91;
          _nw452[(_dafny.ONE).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(2)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(3)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(4)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(5)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(6)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(7)).toNumber()] = (((_this).f11) ? (_2853_v91) : (new _dafny.CodePoint('i'.codePointAt(0))));
          _nw452[(new BigNumber(8)).toNumber()] = _module.__default.fm4((_this).f11, _this.f12, _this.f12, globalState);
          _nw452[(new BigNumber(9)).toNumber()] = (((_this).f10) ? (_2853_v91) : (_2853_v91));
          _nw452[(new BigNumber(10)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw452[(new BigNumber(12)).toNumber()] = (((_this).f10) ? (_2853_v91) : ((((_2854_v92).contains((_this).fm7(_this.f12, globalState))) ? ((_2854_v92).get((_this).fm7(_this.f12, globalState))) : (_2853_v91))));
          _nw452[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
          _nw452[(new BigNumber(14)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(15)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
          _nw452[(new BigNumber(17)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(18)).toNumber()] = _module.__default.fm4(!(_this.f12), _this.f12, (_2855_v93).dtor_cf37, globalState);
          _nw452[(new BigNumber(19)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(20)).toNumber()] = _2853_v91;
          _nw452[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _2856_v94 = _nw452;
          let _index480 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_2856_v94).length));
          (_2856_v94)[_index480] = _2853_v91;
          let _2857_v95;
          let _nw453 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2857_v95 = _nw453;
          let _2858_v96;
          _2858_v96 = _dafny.Seq.UnicodeFromString("tjsrprqpl");
          let _index481 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2857_v95).length));
          (_2857_v95)[_index481] = _2858_v96;
          let _2859_v97;
          _2859_v97 = _module.D3.create_DC10(_2858_v96, (_this).f11, (_this).f10);
          let _index482 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_2856_v94).length));
          let _index483 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2857_v95).length));
          let _rhs462 = (_this).f11;
          let _rhs463 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2858_v96, _2858_v96)).length)), p0);
          let _rhs464 = _2853_v91;
          let _rhs465 = ((_this.f12) ? (_2858_v96) : ((_2859_v97).dtor_cf14));
          let _lhs334 = _this;
          let _lhs335 = globalState;
          let _lhs336 = _2856_v94;
          let _lhs337 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_2856_v94).length));
          let _lhs338 = _2857_v95;
          let _lhs339 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2857_v95).length));
          _lhs334.f12 = _rhs462;
          _lhs335.f0 = _rhs463;
          _lhs336[_lhs337] = _rhs464;
          _lhs338[_lhs339] = _rhs465;
          let _2860_v98;
          _2860_v98 = _dafny.Set.fromElements(_this.f12);
          let _2861_v99;
          let _nw454 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _2861_v99 = _nw454;
          let _2862_v100;
          _2862_v100 = _dafny.Seq.of(_2861_v99);
          let _2863_v101;
          _2863_v101 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f11);
          let _2864_v102;
          _2864_v102 = _dafny.Set.fromElements(_2863_v101);
          let _2865_v103;
          let _nw455 = Array((new BigNumber(13)).toNumber());
          _nw455[(_dafny.ZERO).toNumber()] = p0;
          _nw455[(_dafny.ONE).toNumber()] = p0;
          _nw455[(new BigNumber(2)).toNumber()] = p0;
          _nw455[(new BigNumber(3)).toNumber()] = new BigNumber((_2860_v98).length);
          _nw455[(new BigNumber(4)).toNumber()] = p0;
          _nw455[(new BigNumber(5)).toNumber()] = p0;
          _nw455[(new BigNumber(6)).toNumber()] = p0;
          _nw455[(new BigNumber(7)).toNumber()] = new BigNumber((_2862_v100).length);
          _nw455[(new BigNumber(8)).toNumber()] = p0;
          _nw455[(new BigNumber(9)).toNumber()] = p0;
          _nw455[(new BigNumber(10)).toNumber()] = new BigNumber(473);
          _nw455[(new BigNumber(11)).toNumber()] = (_this).fm9(_2858_v96, _2864_v102, globalState);
          _nw455[(new BigNumber(12)).toNumber()] = p0;
          _2865_v103 = _nw455;
          let _2866_v104;
          _2866_v104 = _dafny.Seq.of(_2861_v99, _2865_v103);
          let _2867_v105;
          _2867_v105 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v100,(_this).f11);
          _2867_v105 = (_2867_v105).update(_dafny.Seq.of(_2865_v103), (_this).f11);
          let _2868_v106;
          _2868_v106 = _dafny.Set.fromElements(_dafny.Seq.Concat((_2857_v95)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2857_v95).length))], _2858_v96), (_2857_v95)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2857_v95).length))], (_2857_v95)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2857_v95).length))]);
          _2868_v106 = _2868_v106;
          (globalState).f1 = (p0).multipliedBy(p0);
          _2858_v96 = _2858_v96;
        } else {
          (globalState).f0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
          let _rhs466 = (_this).f10;
          let _rhs467 = p0;
          let _lhs340 = _this;
          let _lhs341 = globalState;
          _lhs340.f12 = _rhs466;
          _lhs341.f0 = _rhs467;
          let _2869_v107;
          let _nw456 = Array((new BigNumber(29)).toNumber());
          _nw456[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw456[(_dafny.ONE).toNumber()] = !(false);
          _nw456[(new BigNumber(2)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(4)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(6)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(7)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(8)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(9)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(10)).toNumber()] = false;
          _nw456[(new BigNumber(11)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(12)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(13)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(14)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(15)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(16)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(18)).toNumber()] = _module.__default.fm0(p0, globalState);
          _nw456[(new BigNumber(19)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(20)).toNumber()] = true;
          _nw456[(new BigNumber(21)).toNumber()] = _this.f12;
          _nw456[(new BigNumber(22)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(23)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(24)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(25)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(26)).toNumber()] = (_this).f11;
          _nw456[(new BigNumber(27)).toNumber()] = (_this).f10;
          _nw456[(new BigNumber(28)).toNumber()] = _this.f12;
          _2869_v107 = _nw456;
          let _2870_v108;
          _2870_v108 = _module.D0.create_DC0(_2869_v107);
          let _2871_v109;
          _2871_v109 = _module.D18.create_DC56(_dafny.Seq.Concat(_2742_v20, _dafny.Seq.of((_this).f10)));
          let _2872_v110;
          _2872_v110 = new _dafny.CodePoint('x'.codePointAt(0));
          let _2873_v111;
          _2873_v111 = _dafny.Seq.UnicodeFromString("wknhrdj");
          let _rhs468 = _module.D0.create_DC0(_2869_v107);
          let _rhs469 = _2871_v109;
          let _rhs470 = p0;
          let _rhs471 = _2872_v110;
          let _rhs472 = (_2873_v111)[_module.__default.safeIndex(p0, new BigNumber((_2873_v111).length))];
          let _lhs342 = globalState;
          _2870_v108 = _rhs468;
          _2871_v109 = _rhs469;
          _lhs342.f1 = _rhs470;
          _2872_v110 = _rhs471;
          _2872_v110 = _rhs472;
          (_this).f12 = _this.f12;
          let _2874_v112;
          let _2875_v113;
          let _out74;
          let _out75;
          let _outcollector33 = _module.__default.m0((_dafny.ZERO).minus((p0).minus(new BigNumber(507))), _2869_v107, globalState);
          _out74 = _outcollector33[0];
          _out75 = _outcollector33[1];
          _2874_v112 = _out74;
          _2875_v113 = _out75;
        }
        let _2876_v114;
        _2876_v114 = new _dafny.CodePoint('o'.codePointAt(0));
        _2876_v114 = _2876_v114;
      }
      let _2877_v115;
      let _nw457 = Array((new BigNumber(25)).toNumber()).fill(_module.D3.Default());
      _2877_v115 = _nw457;
      let _rhs473 = !(_this.f12) || (_this.f12);
      let _rhs474 = _2877_v115;
      let _lhs343 = _this;
      _lhs343.f12 = _rhs473;
      _2877_v115 = _rhs474;
      let _2878_v116;
      let _init72 = function (_2879_i9) {
        return false;
      };
      let _nw458 = Array((new BigNumber(4)).toNumber());
      for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw458.length); _i0_72++) {
        _nw458[_i0_72] = _init72(new BigNumber(_i0_72));
      }
      _2878_v116 = _nw458;
      let _2880_v117;
      _2880_v117 = _dafny.Map.Empty.slice().updateUnsafe(_2878_v116,_this.f12);
      _2880_v117 = (_2880_v117).update(((_this.f12) ? (_2878_v116) : (_2878_v116)), _this.f12);
      return;
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _2881_i0;
      _2881_i0 = _dafny.ZERO;
      L15: {
        while ((_this).f10) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2881_i0)) {
              break L15;
            }
            _2881_i0 = (_2881_i0).plus(_dafny.ONE);
            (_this).f12 = (_this).f10;
            let _2882_v0;
            _2882_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
            r0 = !((((_2882_v0).contains(p1)) ? ((_2882_v0).get(p1)) : (p0))).isEqualTo((_dafny.ZERO).minus(p0));
            let _2883_v1;
            let _nw459 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
            _2883_v1 = _nw459;
            let _2884_v2;
            _2884_v2 = _dafny.Seq.of(((_this.f12) ? (_2883_v1) : (_2883_v1)));
            _2884_v2 = _2884_v2;
            let _2885_v3;
            let _nw460 = Array((new BigNumber(7)).toNumber()).fill(false);
            _2885_v3 = _nw460;
            let _2886_v4;
            _2886_v4 = _dafny.Seq.of(false);
            let _2887_v5;
            let _nw461 = Array((new BigNumber(22)).toNumber());
            _nw461[(_dafny.ZERO).toNumber()] = _2885_v3;
            _nw461[(_dafny.ONE).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(2)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(3)).toNumber()] = (((_2886_v4)[_module.__default.safeIndex(p0, new BigNumber((_2886_v4).length))]) ? (_2885_v3) : (_2885_v3));
            _nw461[(new BigNumber(4)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(5)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(6)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(7)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(8)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(9)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(10)).toNumber()] = ((p1) ? (_2885_v3) : (_2885_v3));
            _nw461[(new BigNumber(11)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(12)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(13)).toNumber()] = ((!(!(true))) ? (_2885_v3) : (_2885_v3));
            _nw461[(new BigNumber(14)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(15)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(16)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(17)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(18)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(19)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(20)).toNumber()] = _2885_v3;
            _nw461[(new BigNumber(21)).toNumber()] = _2885_v3;
            _2887_v5 = _nw461;
            let _index484 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2887_v5).length));
            (_2887_v5)[_index484] = ((p1) ? (_2885_v3) : (_2885_v3));
            let _index485 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2887_v5).length));
            (_2887_v5)[_index485] = _2885_v3;
          }
        }
      }
      (globalState).f7 = p0;
      let _2888_v6;
      _2888_v6 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),(_this).f11);
      let _2889_v7;
      _2889_v7 = _module.D11.create_DC32(_2888_v6);
      let _pat_let_tv74 = p0;
      (globalState).f7 = function (_source45) {
        if (_source45.is_DC33) {
          let _2890___mcc_h0 = (_source45).cf54;
          let _2891___mcc_h1 = (_source45).cf55;
          let _2892_cf55 = _2891___mcc_h1;
          let _2893_cf54 = _2890___mcc_h0;
          return _pat_let_tv74;
        } else if (_source45.is_DC34) {
          let _2894___mcc_h2 = (_source45).cf56;
          let _2895___mcc_h3 = (_source45).cf57;
          let _2896___mcc_h4 = (_source45).cf58;
          let _2897_cf58 = _2896___mcc_h4;
          let _2898_cf57 = _2895___mcc_h3;
          let _2899_cf56 = _2894___mcc_h2;
          return new BigNumber((_dafny.Seq.of(_this.f12)).length);
        } else {
          let _2900___mcc_h5 = (_source45).cf53;
          let _2901_cf53 = _2900___mcc_h5;
          return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f11, _this.f12, _this.f12), _dafny.Seq.of((_this).f11))).length);
        }
      }(_2889_v7);
      let _2902_v8;
      _2902_v8 = _dafny.Seq.UnicodeFromString("julndsliu");
      let _2903_v9;
      _2903_v9 = new _dafny.CodePoint('f'.codePointAt(0));
      _2902_v8 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_2902_v8, _module.__default.safeIndex(p0, new BigNumber((_2902_v8).length)), _2903_v9), _2902_v8), _dafny.Seq.Concat(_2902_v8, _2902_v8));
      let _2904_v10;
      let _nw462 = Array((new BigNumber(10)).toNumber()).fill([]);
      _2904_v10 = _nw462;
      let _2905_v11;
      _2905_v11 = _module.D17.create_DC50(_2904_v10);
      let _2906_v12;
      _2906_v12 = _module.D17.create_DC52(_2905_v11);
      let _2907_v13;
      _2907_v13 = _module.D17.create_DC52(_2906_v12);
      let _2908_v14;
      _2908_v14 = _dafny.Seq.of(_2907_v13);
      let _2909_v15;
      _2909_v15 = _dafny.MultiSet.fromElements(_2908_v14, _2908_v14);
      let _2910_v16;
      _2910_v16 = _dafny.Seq.of(_2908_v14);
      _2909_v15 = ((_2909_v15).Difference(_dafny.MultiSet.FromArray(_2910_v16))).Union(_2909_v15);
      let _2911_v17;
      let _init73 = function (_2912_i1) {
        return (_this).f10;
      };
      let _nw463 = Array((new BigNumber(25)).toNumber());
      for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw463.length); _i0_73++) {
        _nw463[_i0_73] = _init73(new BigNumber(_i0_73));
      }
      _2911_v17 = _nw463;
      let _index486 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_2911_v17).length));
      (_2911_v17)[_index486] = (_this).fm8(p0, _this.f12, (_this).f11, globalState);
      let _2913_v18;
      _2913_v18 = _module.D14.create_DC43(_2903_v9, p0, (_dafny.ZERO).minus(p0), p0);
      let _index487 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_2911_v17).length));
      let _rhs475 = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus((_2913_v18).dtor_cf78))).length);
      let _rhs476 = !(_this.f12) || (false);
      let _rhs477 = new BigNumber((_module.__default.fm11(((true) ? ((_this).f11) : (_this.f12)), p0, p0, globalState)).length);
      let _lhs344 = globalState;
      let _lhs345 = _2911_v17;
      let _lhs346 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_2911_v17).length));
      let _lhs347 = globalState;
      _lhs344.f1 = _rhs475;
      _lhs345[_lhs346] = _rhs476;
      _lhs347.f7 = _rhs477;
      r0 = !(_this.f12) || ((_this).f11);
      r1 = new BigNumber(-942);
      return [r0, r1];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C13 = class C13 {
    constructor () {
      this._tname = "_module.C13";
      this._f8 = false;
      this._f9 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f8, f9) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(((_module.D4.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f8, (_this).f8),false))).dtor_cf20).length)))).length))).length)).plus(new BigNumber(-400))).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_this).f9)).length)));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _2914_v0;
      _2914_v0 = _dafny.Seq.of(p0, p0, p0);
      let _2915_v1;
      let _nw464 = new _module.C11();
      _nw464.__ctor(new BigNumber((_2914_v0).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-595)), function (_2916_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }), (_this).f9);
      _2915_v1 = _nw464;
      let _2917_i1;
      _2917_i1 = _dafny.ZERO;
      L16: {
        while ((_this).f8) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2917_i1)) {
              break L16;
            }
            _2917_i1 = (_2917_i1).plus(_dafny.ONE);
            let _2918_v2;
            let _nw465 = new _module.C7();
            _nw465.__ctor(!((_this).f9) || ((_this).f8), (_this).f8);
            _2918_v2 = _nw465;
            let _2919_v3;
            let _nw466 = new _module.C5();
            _nw466.__ctor();
            _2919_v3 = _nw466;
            let _2920_v4;
            _2920_v4 = _dafny.MultiSet.fromElements(_2919_v3);
            _2920_v4 = _dafny.MultiSet.fromElements(_2919_v3, _2919_v3);
            let _2921_v5;
            _2921_v5 = _module.D7.create_DC23();
            let _source46 = _2921_v5;
            if (_source46.is_DC23) {
              (globalState).f1 = (_2915_v1).f13;
              let _2922_v6;
              _2922_v6 = _dafny.Seq.of((_2918_v2).f19);
              let _2923_v7;
              _2923_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2918_v2).f19,_2922_v6);
              let _2924_v8;
              let _init74 = ((_2925_v1) => function (_2926_i2) {
                return (_2925_v1).f14;
              })(_2915_v1);
              let _nw467 = Array((new BigNumber(20)).toNumber());
              for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw467.length); _i0_74++) {
                _nw467[_i0_74] = _init74(new BigNumber(_i0_74));
              }
              _2924_v8 = _nw467;
              let _2927_v9;
              _2927_v9 = _dafny.Set.fromElements((_2915_v1).f14);
              let _2928_v10;
              _2928_v10 = new _dafny.CodePoint('k'.codePointAt(0));
              let _rhs478 = _2923_v7;
              let _rhs479 = (((function () {
                let _coll66 = new _dafny.Set();
                for (const _compr_66 of ((_dafny.MultiSet.fromElements((_2915_v1).f14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), function (_2929_i3) {
                  return new _dafny.CodePoint('t'.codePointAt(0));
                }), (_2915_v1).f14, (_2915_v1).f14, _dafny.Seq.UnicodeFromString("cwulthy"))).update(_dafny.Seq.update((_2915_v1).f14, _module.__default.safeIndex((_2915_v1).f13, new BigNumber(((_2915_v1).f14).length)), _2928_v10), _module.__default.abs(new BigNumber(293)))).Elements) {
                  let _2930_v11 = _compr_66;
                  if (((_dafny.MultiSet.fromElements((_2915_v1).f14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), function (_2929_i3) {
                    return new _dafny.CodePoint('t'.codePointAt(0));
                  }), (_2915_v1).f14, (_2915_v1).f14, _dafny.Seq.UnicodeFromString("cwulthy"))).update(_dafny.Seq.update((_2915_v1).f14, _module.__default.safeIndex((_2915_v1).f13, new BigNumber(((_2915_v1).f14).length)), _2928_v10), _module.__default.abs(new BigNumber(293)))).contains(_2930_v11)) {
                    _coll66.add(_2930_v11);
                  }
                }
                return _coll66;
              }()).IsSubsetOf(_2927_v9)) ? (_2924_v8) : (_2924_v8));
              _2923_v7 = _rhs478;
              _2924_v8 = _rhs479;
              let _2931_v12;
              let _nw468 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
              _2931_v12 = _nw468;
              let _index488 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_2931_v12).length));
              (_2931_v12)[_index488] = _module.__default.safeDivisionInt(p0, new BigNumber(116));
              let _2932_v13;
              _2932_v13 = _dafny.MultiSet.fromElements(false, (_2918_v2).f19);
              let _2933_v14;
              _2933_v14 = _dafny.Map.Empty.slice().updateUnsafe((_2915_v1).f13,new BigNumber(441));
              let _index489 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_2931_v12).length));
              (_2931_v12)[_index489] = new BigNumber((((_2932_v13).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.update(_2922_v6, _module.__default.safeIndex((((_2933_v14).contains((_dafny.ZERO).minus(p0))) ? ((_2933_v14).get((_dafny.ZERO).minus(p0))) : ((_dafny.ZERO).minus(p0))), new BigNumber((_2922_v6).length)), (_2918_v2).f19)))).Difference((_2932_v13).Difference(_2932_v13))).cardinality());
              let _2934_v15;
              _2934_v15 = _dafny.Map.Empty.slice().updateUnsafe((_2918_v2).f19,new BigNumber(((_dafny.MultiSet.fromElements(_2933_v14)).update(_2933_v14, _module.__default.abs((_2915_v1).f13))).cardinality()));
              let _2935_v16;
              _2935_v16 = _dafny.Map.Empty.slice().updateUnsafe((_2931_v12)[_module.__default.safeIndex(new BigNumber(898), new BigNumber((_2931_v12).length))],(_this).f9);
              let _2936_v17;
              _2936_v17 = _dafny.Seq.of((_2933_v14).update(new BigNumber((_2935_v16).length), p0));
              _2934_v15 = (_2934_v15).update((_2918_v2).f19, new BigNumber(((_2936_v17)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("hioaqwmh")).length), new BigNumber((_2936_v17).length))]).length));
            } else if (_source46.is_DC24) {
              let _2937___mcc_h0 = (_source46).cf41;
              let _2938_cf41 = _2937___mcc_h0;
              let _2939_v18;
              _2939_v18 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber(859), (_2915_v1).f13), _2914_v0);
              let _2940_v19;
              let _nw469 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _2940_v19 = _nw469;
              let _2941_v20;
              _2941_v20 = new _dafny.CodePoint('t'.codePointAt(0));
              let _2942_v21;
              _2942_v21 = _dafny.Seq.of((_2939_v18)[_module.__default.safeIndex((_module.D12.create_DC37(new BigNumber((_dafny.MultiSet.FromArray(_2914_v0)).cardinality()), _2940_v19, _2941_v20)).dtor_cf61, new BigNumber((_2939_v18).length))]);
              let _2943_v22;
              _2943_v22 = _dafny.Set.fromElements(true);
              let _2944_v23;
              _2944_v23 = _module.D10.create_DC30(_2943_v22);
              _2939_v18 = _module.__default.fm59(true, _dafny.Seq.Concat((_2915_v1).f14, (_2915_v1).f14), _2944_v23, (_2915_v1).f13, globalState);
              let _2945_v24;
              _2945_v24 = _dafny.Map.Empty.slice().updateUnsafe((_2918_v2).f19,(_this).f8);
              let _2946_v25;
              _2946_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_2914_v0),new _dafny.CodePoint('t'.codePointAt(0)));
              let _2947_v27;
              _2947_v27 = _dafny.MultiSet.fromElements((_2915_v1).f13, new BigNumber(((function () {
                let _coll67 = new _dafny.Map();
                for (const _compr_67 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_2914_v0, _2914_v0))).Elements) {
                  let _2948_v26 = _compr_67;
                  if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_2914_v0, _2914_v0))).contains(_2948_v26)) {
                    _coll67.push([_2948_v26,new BigNumber(-199)]);
                  }
                }
                return _coll67;
              }()).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(572)), ((_2949_cf41) => function (_2950_i4) {
                return _2949_cf41;
              })(_2938_cf41)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2938_cf41,(_2915_v1).f13)).length))).length));
              let _2951_v28;
              _2951_v28 = _module.D15.create_DC46();
              let _2952_v29;
              let _2953_v30;
              let _2954_v31;
              let _out76;
              let _out77;
              let _out78;
              let _outcollector34 = (_2918_v2).m3(_module.__default.fm60(_module.__default.fm61(!((_this).f9), _2945_v24, (_2915_v1).f13, (_this).f9, globalState), (((_2946_v25).contains(_2947_v27)) ? ((_2946_v25).get(_2947_v27)) : (new _dafny.CodePoint('e'.codePointAt(0)))), (_this).f9, _2951_v28, globalState), (_this).f8, globalState);
              _out76 = _outcollector34[0];
              _out77 = _outcollector34[1];
              _out78 = _outcollector34[2];
              _2952_v29 = _out76;
              _2953_v30 = _out77;
              _2954_v31 = _out78;
              r0 = !((_2938_cf41).plus(p0)).isEqualTo(_module.__default.safeModuloInt(new BigNumber((_2947_v27).cardinality()), new BigNumber(818)));
              let _2955_v32;
              _2955_v32 = _dafny.Set.fromElements(_2941_v20);
              let _2956_v33;
              _2956_v33 = _dafny.MultiSet.fromElements((_2947_v27).IsProperSubsetOf(_dafny.MultiSet.fromElements((_2915_v1).f13, new BigNumber((_2955_v32).length), (_2915_v1).f13)), (((_2918_v2).f19) ? ((_2918_v2).f19) : ((_2918_v2).f19)), (_2918_v2).f19);
              let _rhs480 = _module.__default.safeDivisionInt(new BigNumber(((_2915_v1).f14).length), (new BigNumber(305)).minus(_2938_cf41));
              let _rhs481 = (_2915_v1).f13;
              let _rhs482 = _2956_v33;
              let _lhs348 = globalState;
              _lhs348.f0 = _rhs480;
              _2938_cf41 = _rhs481;
              _2956_v33 = _rhs482;
            } else {
              let _2957___mcc_h1 = (_source46).cf40;
              let _2958_cf40 = _2957___mcc_h1;
              let _2959_v34;
              let _nw470 = Array((_dafny.ONE).toNumber()).fill([]);
              _2959_v34 = _nw470;
              let _2960_v35;
              let _nw471 = new _module.C12();
              _nw471.__ctor((_2918_v2).f19, (_2918_v2).f19, (_2918_v2).f19);
              _2960_v35 = _nw471;
              let _2961_v36;
              _2961_v36 = _dafny.Map.Empty.slice().updateUnsafe(((_2915_v1).f13).plus((_2915_v1).f13),_dafny.Map.Empty.slice().updateUnsafe((_this).f8,_module.__default.fm0(new BigNumber((_dafny.Set.fromElements(_2960_v35)).length), globalState)));
              let _rhs483 = _2959_v34;
              let _rhs484 = _2961_v36;
              _2959_v34 = _rhs483;
              _2961_v36 = _rhs484;
              r0 = (_2960_v35).f10;
              let _2962_v37;
              _2962_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2918_v2).fm7(!((_this).f8), globalState));
              let _2963_v38;
              let _nw472 = new _module.C2();
              _nw472.__ctor((((_2962_v37).contains(p0)) ? ((_2962_v37).get(p0)) : ((_2960_v35).f10)), new BigNumber(251));
              _2963_v38 = _nw472;
              let _2964_v39;
              _2964_v39 = _dafny.Map.Empty.slice().updateUnsafe((_2960_v35).f10,(_this).f8);
              let _2965_v40;
              _2965_v40 = _dafny.MultiSet.fromElements((_this).f9, (_2963_v38).f21);
              let _2966_v41;
              _2966_v41 = _dafny.Seq.of(_2965_v40);
              let _2967_v42;
              _2967_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2915_v1,_2966_v41);
              let _2968_v43;
              _2968_v43 = _dafny.Seq.of((_this).f8);
              let _2969_v44;
              _2969_v44 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_2964_v39).length), (_2915_v1).f13),(((_2967_v42).contains(_2915_v1)) ? ((_2967_v42).get(_2915_v1)) : (_dafny.Seq.of(_dafny.MultiSet.FromArray(_2968_v43)))));
              let _2970_v45;
              let _nw473 = Array((new BigNumber(28)).toNumber()).fill([]);
              _2970_v45 = _nw473;
              let _2971_v46;
              _2971_v46 = _module.D8.create_DC25(_2970_v45);
              let _2972_v47;
              _2972_v47 = _dafny.MultiSet.fromElements(_2971_v46, _module.D8.create_DC25(_2970_v45));
              _2969_v44 = (_2969_v44).update((((_2965_v40).contains((_2918_v2).f19)) ? ((_2965_v40).get((_2918_v2).f19)) : ((_dafny.ZERO).minus(new BigNumber((_2972_v47).cardinality())))), _2966_v41);
            }
            r0 = false;
          }
        }
      }
      let _2973_v48;
      _2973_v48 = _dafny.Seq.of((_this).f9, (_this).f8, (_this).f8, !((_this).f9), (_this).f9);
      let _2974_v49;
      _2974_v49 = _dafny.Map.Empty.slice().updateUnsafe((_2915_v1).fm8(p0, (_this).f9, (_this).f8, globalState),(_2973_v48)[_module.__default.safeIndex((_2914_v0)[_module.__default.safeIndex(p0, new BigNumber((_2914_v0).length))], new BigNumber((_2973_v48).length))]);
      let _2975_v50;
      let _nw474 = Array((new BigNumber(3)).toNumber());
      _nw474[(_dafny.ZERO).toNumber()] = false;
      _nw474[(_dafny.ONE).toNumber()] = (_this).f9;
      _nw474[(new BigNumber(2)).toNumber()] = (((_2974_v49).contains((_this).f9)) ? ((_2974_v49).get((_this).f9)) : ((_this).f9));
      _2975_v50 = _nw474;
      let _index490 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2975_v50).length));
      (_2975_v50)[_index490] = (_this).f9;
      let _index491 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2975_v50).length));
      (_2975_v50)[_index491] = (_this).f9;
      let _2976_v51;
      _2976_v51 = _dafny.Seq.of((_2915_v1).f14, (_2915_v1).f14, (_2915_v1).f14, (_2915_v1).f14, (_2915_v1).f14);
      _2976_v51 = _2976_v51;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2975_v50).length))) {
        let _2977_i5 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2977_i5)) && ((_2977_i5).isLessThan(new BigNumber((_2975_v50).length))))) {
          (_2975_v50)[(_2977_i5)] = true;
        }
      }
      let _2978_v52;
      let _nw475 = new _module.C3();
      _nw475.__ctor((_this).f9, (_this).f9);
      _2978_v52 = _nw475;
      let _rhs485 = new BigNumber(838);
      let _rhs486 = _2978_v52;
      let _lhs349 = globalState;
      _lhs349.f1 = _rhs485;
      _2978_v52 = _rhs486;
      r0 = (((_this).f8) ? (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("bnr"), (_2915_v1).f14)) : ((_this).f8));
      return r0;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f0 = new BigNumber(20);
      let _2979_v0;
      let _nw476 = Array((new BigNumber(27)).toNumber()).fill([]);
      _2979_v0 = _nw476;
      let _2980_v1;
      _2980_v1 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2981_v2;
      _2981_v2 = _dafny.Seq.of(_2980_v1, _2980_v1);
      let _2982_v3;
      _2982_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2980_v1,_2980_v1);
      let _2983_v4;
      _2983_v4 = _module.D3.create_DC10(_2981_v2, (_this).f9, (_this).f9);
      let _2984_v5;
      let _nw477 = Array((new BigNumber(29)).toNumber());
      _nw477[(_dafny.ZERO).toNumber()] = _2981_v2;
      _nw477[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_2980_v1, (((_2982_v3).contains(_2980_v1)) ? ((_2982_v3).get(_2980_v1)) : (_2980_v1)));
      _nw477[(new BigNumber(2)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(3)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(4)).toNumber()] = (_2983_v4).dtor_cf14;
      _nw477[(new BigNumber(5)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(6)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(7)).toNumber()] = _module.__default.fm25(p2, p0, p2, globalState);
      _nw477[(new BigNumber(8)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(9)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(10)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(11)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(12)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(13)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(14)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(15)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(16)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(17)).toNumber()] = (_2983_v4).dtor_cf14;
      _nw477[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_2980_v1, _module.__default.fm4(p1, p1, (_this).f8, globalState));
      _nw477[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), ((_2985_v1) => function (_2986_i0) {
        return _2985_v1;
      })(_2980_v1));
      _nw477[(new BigNumber(20)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(21)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(22)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(23)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(24)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(25)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(26)).toNumber()] = _2981_v2;
      _nw477[(new BigNumber(27)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_2987_i1) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      });
      _nw477[(new BigNumber(28)).toNumber()] = _2981_v2;
      _2984_v5 = _nw477;
      let _index492 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length));
      (_2979_v0)[_index492] = _2984_v5;
      let _2988_v6;
      _2988_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm0(new BigNumber(488), globalState)),(_dafny.ZERO).minus(p2));
      let _2989_v7;
      _2989_v7 = _dafny.Set.fromElements(new BigNumber(10));
      let _index493 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length));
      let _rhs487 = _2984_v5;
      let _rhs488 = _2984_v5;
      let _rhs489 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? ((_this).f8) : ((_this).f9)),new BigNumber(((_dafny.Set.fromElements((_dafny.ZERO).minus(p2))).Union(_2989_v7)).length));
      let _lhs350 = _2979_v0;
      let _lhs351 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length));
      _lhs350[_lhs351] = _rhs487;
      _2984_v5 = _rhs488;
      _2988_v6 = _rhs489;
      let _2990_v8;
      _2990_v8 = _dafny.Seq.of(p1);
      let _2991_v9;
      _2991_v9 = _module.D4.create_DC14(p2, p2);
      let _hi19 = (_2991_v9).dtor_cf22;
      for (let _2992_i2 = _module.__default.safeDivisionInt(p2, new BigNumber((_2990_v8).length)); _2992_i2.isLessThan(_hi19); _2992_i2 = _2992_i2.plus(_dafny.ONE)) {
        let _2993_v10;
        _2993_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2992_i2,p0);
        let _2994_v11;
        _2994_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8);
        let _2995_v12;
        _2995_v12 = _dafny.Set.fromElements((((_2994_v11).contains(p0)) ? ((_2994_v11).get(p0)) : (false)));
        let _2996_v13;
        _2996_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2993_v10,_module.__default.safeModuloInt(new BigNumber((_2995_v12).length), p2));
        _2996_v13 = (_2996_v13).update(_2993_v10, p2);
        _2981_v2 = _module.__default.fm13(p1, !((_dafny.ZERO).minus(p2)).isEqualTo(_2992_i2), globalState);
        if (p0) {
          let _2997_v14;
          let _nw478 = Array((new BigNumber(7)).toNumber()).fill(false);
          _2997_v14 = _nw478;
          let _index494 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2997_v14).length));
          (_2997_v14)[_index494] = (p0) === (p0);
          let _index495 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2997_v14).length));
          (_2997_v14)[_index495] = p0;
          let _2998_v15;
          let _2999_v16;
          let _out79;
          let _out80;
          let _outcollector35 = _module.__default.m0(new BigNumber(782), _2997_v14, globalState);
          _out79 = _outcollector35[0];
          _out80 = _outcollector35[1];
          _2998_v15 = _out79;
          _2999_v16 = _out80;
          let _3000_v17;
          let _nw479 = new _module.C6();
          _nw479.__ctor();
          _3000_v17 = _nw479;
          let _3001_v18;
          let _nw480 = new _module.C3();
          _nw480.__ctor((p2).isLessThanOrEqualTo(p2), p1);
          _3001_v18 = _nw480;
          let _3002_v19;
          let _nw481 = new _module.C2();
          _nw481.__ctor(p0, (((_this).f8) ? (p2) : (p2)));
          _3002_v19 = _nw481;
        } else {
          let _3003_v20;
          _3003_v20 = _dafny.Seq.of(new BigNumber(646), p2, new BigNumber(154));
          let _3004_v21;
          _3004_v21 = _module.D11.create_DC33((_this).f9, p2);
          let _3005_v22;
          let _nw482 = Array((new BigNumber(14)).toNumber());
          _nw482[(_dafny.ZERO).toNumber()] = false;
          _nw482[(_dafny.ONE).toNumber()] = p0;
          _nw482[(new BigNumber(2)).toNumber()] = (_3004_v21).dtor_cf54;
          _nw482[(new BigNumber(3)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(4)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(6)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(7)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(8)).toNumber()] = p0;
          _nw482[(new BigNumber(9)).toNumber()] = p1;
          _nw482[(new BigNumber(10)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(11)).toNumber()] = p0;
          _nw482[(new BigNumber(12)).toNumber()] = (_this).f9;
          _nw482[(new BigNumber(13)).toNumber()] = (_this).f9;
          _3005_v22 = _nw482;
          let _3006_v23;
          let _3007_v24;
          let _out81;
          let _out82;
          let _outcollector36 = _module.__default.m0(new BigNumber((_3003_v20).length), _3005_v22, globalState);
          _out81 = _outcollector36[0];
          _out82 = _outcollector36[1];
          _3006_v23 = _out81;
          _3007_v24 = _out82;
          let _arr10 = (_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))];
          let _index496 = _module.__default.safeIndex(new BigNumber(699), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length));
          _arr10[_index496] = _dafny.Seq.Concat(_dafny.Seq.update(_3006_v23, _module.__default.safeIndex(_2992_i2, new BigNumber((_3006_v23).length)), _2980_v1), _3006_v23);
          let _3008_v25;
          _3008_v25 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("vxw"));
          let _arr11 = (_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))];
          let _index497 = _module.__default.safeIndex(new BigNumber(699), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length));
          _arr11[_index497] = (_3008_v25)[_module.__default.safeIndex(_2992_i2, new BigNumber((_3008_v25).length))];
          let _index498 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_3005_v22).length));
          (_3005_v22)[_index498] = (new BigNumber((((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))])[_module.__default.safeIndex(new BigNumber(699), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length))]).length)).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements(p0, (_2990_v8)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll68 = new _dafny.Map();
            for (const _compr_68 of (_3006_v23).Elements) {
              let _3009_v26 = _compr_68;
              if (_dafny.Seq.contains(_3006_v23, _3009_v26)) {
                _coll68.push([_3009_v26,p0]);
              }
            }
            return _coll68;
          }()).length), new BigNumber((_2990_v8).length))])).cardinality()));
          let _index499 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_3005_v22).length));
          (_3005_v22)[_index499] = p0;
          let _3010_v28;
          _3010_v28 = _module.D19.create_DC57(function () {
  let _coll69 = new _dafny.Set();
  for (const _compr_69 of _dafny.IntegerRange(new BigNumber(219), new BigNumber(-661))) {
    let _3011_v27 = _compr_69;
    if (((new BigNumber(219)).isLessThanOrEqualTo(_3011_v27)) && ((_3011_v27).isLessThan(new BigNumber(-661)))) {
      _coll69.add((_3011_v27).plus(new BigNumber((_3007_v24).length)));
    }
  }
  return _coll69;
}());
          let _3012_v29;
          _3012_v29 = _dafny.Seq.of(_2989_v7, _2989_v7, _2989_v7, (_3010_v28).dtor_cf98, _2989_v7);
          _2989_v7 = (_3012_v29)[_module.__default.safeIndex(new BigNumber((((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))])[_module.__default.safeIndex(new BigNumber(699), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length))]).length), new BigNumber((_3012_v29).length))];
          let _3013_v30;
          _3013_v30 = _dafny.Seq.of(_2993_v10, _2993_v10, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("aldu")).length),(_3005_v22)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_3005_v22).length))]));
          let _3014_v31;
          let _nw483 = Array((new BigNumber(24)).toNumber());
          _nw483[(_dafny.ZERO).toNumber()] = _2980_v1;
          _nw483[(_dafny.ONE).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(2)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(3)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(4)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(5)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw483[(new BigNumber(7)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(8)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(9)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(10)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(11)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(12)).toNumber()] = _module.__default.fm4((_this).f8, p0, (_3005_v22)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_3005_v22).length))], globalState);
          _nw483[(new BigNumber(13)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(14)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(15)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(16)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(17)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(18)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(19)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(20)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(21)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(22)).toNumber()] = _2980_v1;
          _nw483[(new BigNumber(23)).toNumber()] = _2980_v1;
          _3014_v31 = _nw483;
          let _3015_v32;
          _3015_v32 = _module.D12.create_DC37(new BigNumber((_3013_v30).length), _3014_v31, _2980_v1);
          let _3016_v33;
          _3016_v33 = _module.D13.create_DC40(_3015_v32, p0, _2992_i2, _2992_i2);
          (globalState).f7 = (_module.__default.safeModuloInt(new BigNumber((_2989_v7).length), (_3016_v33).dtor_cf69)).minus(_2992_i2);
        }
        let _3017_v34;
        _3017_v34 = false;
        _3017_v34 = (_this).f8;
      }
      if (((p2).multipliedBy(new BigNumber(-824))).isEqualTo((_dafny.ZERO).minus((p2).minus(p2)))) {
        let _3018_v35;
        _3018_v35 = true;
        _3018_v35 = p0;
        let _3019_v36;
        let _nw484 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _3019_v36 = _nw484;
        let _index500 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_3019_v36).length));
        (_3019_v36)[_index500] = _2980_v1;
        let _3020_v37;
        _3020_v37 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm25(p2, p1, p2, globalState));
        let _arr12 = (_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))];
        let _index501 = _module.__default.safeIndex(new BigNumber(715), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length));
        _arr12[_index501] = (((_3020_v37).contains(p0)) ? ((_3020_v37).get(p0)) : (_2981_v2));
        let _3021_v38;
        let _init75 = ((_3022_p2) => function (_3023_i3) {
          return _module.__default.safeDivisionInt(_3023_i3, _3022_p2);
        })(p2);
        let _nw485 = Array((new BigNumber(17)).toNumber());
        for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw485.length); _i0_75++) {
          _nw485[_i0_75] = _init75(new BigNumber(_i0_75));
        }
        _3021_v38 = _nw485;
        let _index502 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length));
        (_3021_v38)[_index502] = (_dafny.ZERO).minus((p2).minus(p2));
        let _3024_v39;
        _3024_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _3025_v40;
        let _nw486 = new _module.C9();
        _nw486.__ctor(p0, (((_3024_v39).contains(_3018_v35)) ? ((_3024_v39).get(_3018_v35)) : (!(p1))));
        _3025_v40 = _nw486;
        let _3026_v41;
        _3026_v41 = _dafny.Map.Empty.slice().updateUnsafe(p2,_3025_v40);
        let _3027_v42;
        let _nw487 = Array((new BigNumber(13)).toNumber());
        _nw487[(_dafny.ZERO).toNumber()] = _3025_v40;
        _nw487[(_dafny.ONE).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(2)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(3)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(4)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(5)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(6)).toNumber()] = (((_3026_v41).contains(p2)) ? ((_3026_v41).get(p2)) : (_3025_v40));
        _nw487[(new BigNumber(7)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(8)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(9)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(10)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(11)).toNumber()] = _3025_v40;
        _nw487[(new BigNumber(12)).toNumber()] = _3025_v40;
        _3027_v42 = _nw487;
        let _index503 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_3019_v36).length));
        let _arr13 = (_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))];
        let _index504 = _module.__default.safeIndex(new BigNumber(715), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length));
        let _index505 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length));
        let _rhs490 = _2980_v1;
        let _rhs491 = _dafny.Seq.UnicodeFromString("adr");
        let _rhs492 = (_dafny.ZERO).minus(p2);
        let _rhs493 = (_module.D22.create_DC65(_3027_v42)).dtor_cf111;
        let _lhs352 = _3019_v36;
        let _lhs353 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_3019_v36).length));
        let _lhs354 = (_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))];
        let _lhs355 = _module.__default.safeIndex(new BigNumber(715), new BigNumber(((_2979_v0)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_2979_v0).length))]).length));
        let _lhs356 = _3021_v38;
        let _lhs357 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length));
        _lhs352[_lhs353] = _rhs490;
        _lhs354[_lhs355] = _rhs491;
        _lhs356[_lhs357] = _rhs492;
        _3027_v42 = _rhs493;
        (globalState).f7 = new BigNumber((_module.__default.fm62(_2989_v7, globalState)).length);
        let _source47 = _module.__default.fm63(globalState);
        if (_source47.is_DC17) {
          let _3028___mcc_h0 = (_source47).cf28;
          let _3029___mcc_h1 = (_source47).cf29;
          let _3030_cf29 = _3029___mcc_h1;
          let _3031_cf28 = _3028___mcc_h0;
          _2989_v7 = _2989_v7;
          (globalState).f7 = new BigNumber(582);
          _3018_v35 = _3018_v35;
          let _3032_v43;
          _3032_v43 = _module.D2.create_DC6(_2980_v1);
          _3032_v43 = _3032_v43;
        } else if (_source47.is_DC16) {
          let _3033___mcc_h2 = (_source47).cf27;
          let _3034_cf27 = _3033___mcc_h2;
          let _3035_v44;
          let _nw488 = Array((new BigNumber(8)).toNumber());
          _nw488[(_dafny.ZERO).toNumber()] = p1;
          _nw488[(_dafny.ONE).toNumber()] = (_this).f8;
          _nw488[(new BigNumber(2)).toNumber()] = !((_3025_v40).fm6(p2, new BigNumber((_2988_v6).length), p2, p0, globalState));
          _nw488[(new BigNumber(3)).toNumber()] = p0;
          _nw488[(new BigNumber(4)).toNumber()] = false;
          _nw488[(new BigNumber(5)).toNumber()] = (_this).f8;
          _nw488[(new BigNumber(6)).toNumber()] = (_3025_v40).f17;
          _nw488[(new BigNumber(7)).toNumber()] = (_this).f9;
          _3035_v44 = _nw488;
          let _3036_v45;
          _3036_v45 = _module.D0.create_DC0(_3035_v44);
          let _3037_v46;
          let _nw489 = Array((new BigNumber(29)).toNumber());
          _nw489[(_dafny.ZERO).toNumber()] = _3035_v44;
          _nw489[(_dafny.ONE).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(2)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(3)).toNumber()] = (_3036_v45).dtor_cf0;
          _nw489[(new BigNumber(4)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(5)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(6)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(7)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(8)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(9)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(10)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(11)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(12)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(13)).toNumber()] = (_module.D0.create_DC2(_3018_v35, _3035_v44)).dtor_cf3;
          _nw489[(new BigNumber(14)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(15)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(16)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(17)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(18)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(19)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(20)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(21)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(22)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(23)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(24)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(25)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(26)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(27)).toNumber()] = _3035_v44;
          _nw489[(new BigNumber(28)).toNumber()] = _3035_v44;
          _3037_v46 = _nw489;
          let _3038_v47;
          _3038_v47 = _module.D17.create_DC50(_3037_v46);
          let _pat_let_tv75 = _3037_v46;
          let _pat_let_tv76 = _3037_v46;
          let _pat_let_tv77 = _3037_v46;
          let _3039_v48;
          let _nw490 = Array((new BigNumber(19)).toNumber());
          _nw490[(_dafny.ZERO).toNumber()] = _3038_v47;
          _nw490[(_dafny.ONE).toNumber()] = function (_pat_let70_0) {
            return function (_3040_dt__update__tmp_h0) {
              return function (_pat_let71_0) {
                return function (_3041_dt__update_hcf86_h0) {
                  return _module.D17.create_DC50(_3041_dt__update_hcf86_h0);
                }(_pat_let71_0);
              }(_pat_let_tv75);
            }(_pat_let70_0);
          }(_3038_v47);
          _nw490[(new BigNumber(2)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(3)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(4)).toNumber()] = ((p0) ? (_module.D17.create_DC50(_3037_v46)) : (_module.D17.create_DC50(_3037_v46)));
          _nw490[(new BigNumber(5)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(6)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(7)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(8)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(9)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(10)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(11)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(12)).toNumber()] = _module.D17.create_DC50(_3037_v46);
          _nw490[(new BigNumber(13)).toNumber()] = _module.D17.create_DC50(_3037_v46);
          _nw490[(new BigNumber(14)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(15)).toNumber()] = _3038_v47;
          _nw490[(new BigNumber(16)).toNumber()] = function (_pat_let72_0) {
            return function (_3042_dt__update__tmp_h1) {
              return function (_pat_let73_0) {
                return function (_3043_dt__update_hcf86_h1) {
                  return _module.D17.create_DC50(_3043_dt__update_hcf86_h1);
                }(_pat_let73_0);
              }(_pat_let_tv76);
            }(_pat_let72_0);
          }(_3038_v47);
          _nw490[(new BigNumber(17)).toNumber()] = function (_pat_let74_0) {
            return function (_3044_dt__update__tmp_h2) {
              return function (_pat_let75_0) {
                return function (_3045_dt__update_hcf86_h2) {
                  return _module.D17.create_DC50(_3045_dt__update_hcf86_h2);
                }(_pat_let75_0);
              }(_pat_let_tv77);
            }(_pat_let74_0);
          }(_3038_v47);
          _nw490[(new BigNumber(18)).toNumber()] = _module.D17.create_DC50(_3037_v46);
          _3039_v48 = _nw490;
          let _index506 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_3039_v48).length));
          (_3039_v48)[_index506] = _3038_v47;
          let _index507 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_3039_v48).length));
          let _rhs494 = (_3019_v36)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_3019_v36).length))];
          let _rhs495 = _dafny.Seq.Concat(_2990_v8, _2990_v8);
          let _rhs496 = _3038_v47;
          let _lhs358 = _3039_v48;
          let _lhs359 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_3039_v48).length));
          _2980_v1 = _rhs494;
          _2990_v8 = _rhs495;
          _lhs358[_lhs359] = _rhs496;
          let _3046_v49;
          _3046_v49 = _dafny.MultiSet.fromElements(p2, p2);
          let _index508 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length));
          let _rhs497 = _module.__default.fm18(p2, _2988_v6, !_dafny.areEqual(_2981_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-547)), ((_3047_v36) => function (_3048_i4) {
            return (_3047_v36)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_3047_v36).length))];
          })(_3019_v36))), globalState);
          let _rhs498 = (_3025_v40).fm15(globalState);
          let _lhs360 = _3021_v38;
          let _lhs361 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length));
          _3046_v49 = _rhs497;
          _lhs360[_lhs361] = _rhs498;
          let _3049_v50;
          _3049_v50 = _module.D5.create_DC17(p1, _2980_v1);
          _3018_v35 = ((!(p0)) ? (!((_3049_v50).dtor_cf28)) : ((_3025_v40).f17));
          (globalState).f0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p2));
        } else {
          let _3050___mcc_h3 = (_source47).cf30;
          let _3051_cf30 = _3050___mcc_h3;
          let _3052_v51;
          _3052_v51 = _module.D1.create_DC3(p2);
          let _3053_v52;
          _3053_v52 = _module.D1.create_DC5(_3052_v51);
          let _3054_v53;
          _3054_v53 = _dafny.Set.fromElements(_3053_v52, _3053_v52, _3053_v52);
          _3054_v53 = (_3054_v53).Difference(_3054_v53);
          let _3055_v54;
          let _nw491 = Array((new BigNumber(3)).toNumber()).fill(false);
          _3055_v54 = _nw491;
          let _index509 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_3055_v54).length));
          (_3055_v54)[_index509] = ((_dafny.ZERO).minus(p2)).isEqualTo((_3021_v38)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length))]);
          let _index510 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_3055_v54).length));
          (_3055_v54)[_index510] = !((_3025_v40).f17);
          let _3056_v55;
          let _nw492 = new _module.C3();
          _nw492.__ctor(p0, false);
          _3056_v55 = _nw492;
          let _3057_v56;
          let _nw493 = new _module.C9();
          _nw493.__ctor(((_3021_v38)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length))]).isLessThan((_3021_v38)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length))]), (_3055_v54)[_module.__default.safeIndex(new BigNumber(822), new BigNumber((_3055_v54).length))]);
          _3057_v56 = _nw493;
        }
        let _3058_v57;
        _3058_v57 = _dafny.Seq.of(p2);
        let _3059_v58;
        _3059_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3058_v57).length),(_3021_v38)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_3021_v38).length))]);
        _3059_v58 = (_3059_v58).update((p2).plus(p2), p2);
      } else {
        let _3060_v59;
        _3060_v59 = _dafny.Seq.of(new BigNumber(-513), p2);
        let _3061_v60;
        _3061_v60 = _dafny.MultiSet.fromElements((_this).f9, (_this).f9);
        let _3062_v61;
        let _nw494 = Array((new BigNumber(4)).toNumber());
        _nw494[(_dafny.ZERO).toNumber()] = (_3060_v59)[_module.__default.safeIndex(p2, new BigNumber((_3060_v59).length))];
        _nw494[(_dafny.ONE).toNumber()] = (((_3061_v60).contains(p0)) ? ((_3061_v60).get(p0)) : (p2));
        _nw494[(new BigNumber(2)).toNumber()] = p2;
        _nw494[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(942), (_dafny.ZERO).minus(p2));
        _3062_v61 = _nw494;
        let _3063_v62;
        _3063_v62 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f8);
        let _3064_v63;
        _3064_v63 = _dafny.MultiSet.fromElements(_2980_v1, _2980_v1);
        let _3065_v64;
        let _nw495 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _3065_v64 = _nw495;
        let _3066_v65;
        _3066_v65 = _module.D12.create_DC37((_dafny.ZERO).minus(p2), _3065_v64, _2980_v1);
        let _3067_v66;
        let _nw496 = Array((new BigNumber(28)).toNumber());
        _nw496[(_dafny.ZERO).toNumber()] = (_this).f8;
        _nw496[(_dafny.ONE).toNumber()] = (_this).f8;
        _nw496[(new BigNumber(2)).toNumber()] = (_this).f8;
        _nw496[(new BigNumber(3)).toNumber()] = true;
        _nw496[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw496[(new BigNumber(5)).toNumber()] = p1;
        _nw496[(new BigNumber(6)).toNumber()] = false;
        _nw496[(new BigNumber(7)).toNumber()] = p1;
        _nw496[(new BigNumber(8)).toNumber()] = p0;
        _nw496[(new BigNumber(9)).toNumber()] = p0;
        _nw496[(new BigNumber(10)).toNumber()] = p1;
        _nw496[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw496[(new BigNumber(12)).toNumber()] = (_this).f9;
        _nw496[(new BigNumber(13)).toNumber()] = p0;
        _nw496[(new BigNumber(14)).toNumber()] = (_this).f9;
        _nw496[(new BigNumber(15)).toNumber()] = (_this).f8;
        _nw496[(new BigNumber(16)).toNumber()] = (_this).f8;
        _nw496[(new BigNumber(17)).toNumber()] = p1;
        _nw496[(new BigNumber(18)).toNumber()] = false;
        _nw496[(new BigNumber(19)).toNumber()] = (_this).f8;
        _nw496[(new BigNumber(20)).toNumber()] = p1;
        _nw496[(new BigNumber(21)).toNumber()] = p0;
        _nw496[(new BigNumber(22)).toNumber()] = p1;
        _nw496[(new BigNumber(23)).toNumber()] = false;
        _nw496[(new BigNumber(24)).toNumber()] = (_this).f9;
        _nw496[(new BigNumber(25)).toNumber()] = p0;
        _nw496[(new BigNumber(26)).toNumber()] = p1;
        _nw496[(new BigNumber(27)).toNumber()] = (_this).f8;
        _3067_v66 = _nw496;
        let _3068_v67;
        let _nw497 = new _module.C10();
        _nw497.__ctor(new BigNumber((_3061_v60).cardinality()), p1);
        _3068_v67 = _nw497;
        let _3069_v68;
        _3069_v68 = _dafny.Map.Empty.slice().updateUnsafe(_3067_v66,_3068_v67);
        let _3070_v69;
        _3070_v69 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_3069_v68).length)));
        let _3071_v70;
        _3071_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(160)), ((_3072_v67) => function (_3073_i5) {
          return (_3072_v67).f15;
        })(_3068_v67)),new BigNumber((_3063_v62).length));
        let _3074_v71;
        _3074_v71 = _dafny.Set.fromElements(_3063_v62);
        let _3075_v72;
        _3075_v72 = _dafny.Map.Empty.slice().updateUnsafe((_3068_v67).fm9(_2981_v2, _3074_v71, globalState),p2);
        let _nw498 = Array((new BigNumber(29)).toNumber());
        _nw498[(_dafny.ZERO).toNumber()] = new BigNumber(600);
        _nw498[(_dafny.ONE).toNumber()] = p2;
        _nw498[(new BigNumber(2)).toNumber()] = (p2).multipliedBy((_dafny.ZERO).minus(p2));
        _nw498[(new BigNumber(3)).toNumber()] = p2;
        _nw498[(new BigNumber(4)).toNumber()] = p2;
        _nw498[(new BigNumber(5)).toNumber()] = p2;
        _nw498[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((((_3061_v60).contains(p0)) ? ((_3061_v60).get(p0)) : (_module.__default.fm1(globalState))));
        _nw498[(new BigNumber(7)).toNumber()] = p2;
        _nw498[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(p2), p2));
        _nw498[(new BigNumber(9)).toNumber()] = p2;
        _nw498[(new BigNumber(10)).toNumber()] = p2;
        _nw498[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(p2, (_this).fm5(p2, globalState));
        _nw498[(new BigNumber(12)).toNumber()] = p2;
        _nw498[(new BigNumber(13)).toNumber()] = (p2).multipliedBy(p2);
        _nw498[(new BigNumber(14)).toNumber()] = p2;
        _nw498[(new BigNumber(15)).toNumber()] = p2;
        _nw498[(new BigNumber(16)).toNumber()] = p2;
        _nw498[(new BigNumber(17)).toNumber()] = p2;
        _nw498[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.of(_3063_v62, _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8), _3063_v62)).length);
        _nw498[(new BigNumber(19)).toNumber()] = (((_3064_v63).contains(_2980_v1)) ? ((_3064_v63).get(_2980_v1)) : ((_3060_v59)[_module.__default.safeIndex(new BigNumber((_2981_v2).length), new BigNumber((_3060_v59).length))]));
        _nw498[(new BigNumber(20)).toNumber()] = (p2).plus((_dafny.ZERO).minus(p2));
        _nw498[(new BigNumber(21)).toNumber()] = p2;
        _nw498[(new BigNumber(22)).toNumber()] = (_module.D13.create_DC40(_3066_v65, p1, (((_2988_v6).contains((_this).f8)) ? ((_2988_v6).get((_this).f8)) : (p2)), p2)).dtor_cf70;
        _nw498[(new BigNumber(23)).toNumber()] = p2;
        _nw498[(new BigNumber(24)).toNumber()] = (p2).plus(new BigNumber((_3070_v69).cardinality()));
        _nw498[(new BigNumber(25)).toNumber()] = (p2).multipliedBy(p2);
        _nw498[(new BigNumber(26)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_3071_v70).length));
        _nw498[(new BigNumber(27)).toNumber()] = (p2).multipliedBy(p2);
        _nw498[(new BigNumber(28)).toNumber()] = new BigNumber((_3075_v72).length);
        _3062_v61 = _nw498;
        let _index511 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length));
        (_3067_v66)[_index511] = p0;
        let _index512 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length));
        (_3067_v66)[_index512] = (_2990_v8)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(p2)), new BigNumber((_2990_v8).length))];
        let _index513 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length));
        (_3067_v66)[_index513] = !((_3068_v67).f16);
        let _3076_v73;
        _3076_v73 = _dafny.Map.Empty.slice().updateUnsafe((_3067_v66)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length))],new _dafny.CodePoint('s'.codePointAt(0)));
        let _3077_v74;
        _3077_v74 = _module.D2.create_DC7((_3068_v67).f15, (_3068_v67).f16, new BigNumber((_3076_v73).length));
        (globalState).f7 = (_3077_v74).dtor_cf11;
        if (!((_3068_v67).f16)) {
          let _index514 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_3062_v61).length));
          (_3062_v61)[_index514] = (_3068_v67).f15;
          let _3078_v75;
          _3078_v75 = _dafny.Seq.of(_2988_v6);
          let _index515 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_3062_v61).length));
          (_3062_v61)[_index515] = (_dafny.ZERO).minus(new BigNumber(((((_this).f9) ? (((_3078_v75)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), ((_3079_v67) => function (_3080_i6) {
            return (_3079_v67).f15;
          })(_3068_v67))).length), new BigNumber((_3078_v75).length))]).update((_3067_v66)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length))], (_dafny.ZERO).minus((_3068_v67).f15))) : ((_2988_v6).Merge(_module.__default.fm51(globalState))))).length));
          (globalState).f0 = new BigNumber(389);
          let _index516 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_3062_v61).length));
          (_3062_v61)[_index516] = ((p2).minus(p2)).minus((new BigNumber((_3070_v69).cardinality())).minus((_this).fm5((_3068_v67).f15, globalState)));
          let _3081_v76;
          let _nw499 = new _module.C6();
          _nw499.__ctor();
          _3081_v76 = _nw499;
          _3067_v66 = _3067_v66;
        } else {
          let _3082_v77;
          _3082_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2981_v2,_2980_v1);
          _3082_v77 = (_3082_v77).update(_2981_v2, _2980_v1);
          let _3083_v78;
          _3083_v78 = _dafny.Set.fromElements(_3065_v64);
          let _3084_v79;
          _3084_v79 = _dafny.MultiSet.fromElements(_3083_v78, _3083_v78, _dafny.Set.fromElements(_3065_v64, _3065_v64), _3083_v78, _dafny.Set.fromElements(_3065_v64, _3065_v64));
          let _3085_v80;
          _3085_v80 = _dafny.Map.Empty.slice().updateUnsafe(_2990_v8,(_3067_v66)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length))]);
          let _3086_v82;
          let _nw500 = Array((new BigNumber(27)).toNumber()).fill(_module.D1.Default());
          _3086_v82 = _nw500;
          let _3087_v83;
          _3087_v83 = _dafny.Seq.of(_3086_v82, _3086_v82, _3086_v82);
          let _3088_v84;
          _3088_v84 = _module.D8.create_DC27(p0, _3087_v83, _3060_v59, (_3068_v67).f15);
          let _3089_v85;
          _3089_v85 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_3063_v62).length)),(_3088_v84).dtor_cf45);
          (globalState).f0 = (((_3084_v79).contains(_3083_v78)) ? ((_3084_v79).get(_3083_v78)) : ((new BigNumber((function () {
            let _coll70 = new _dafny.Set();
            for (const _compr_70 of (_3085_v80).Keys.Elements) {
              let _3090_v81 = _compr_70;
              if ((_3085_v80).contains(_3090_v81)) {
                _coll70.add(_3090_v81);
              }
            }
            return _coll70;
          }()).length)).plus((_dafny.ZERO).minus(new BigNumber((_3089_v85).length)))));
          let _index517 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length));
          (_3067_v66)[_index517] = false;
          let _index518 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_3067_v66).length));
          (_3067_v66)[_index518] = (_this).f9;
          let _3091_v86;
          _3091_v86 = _module.D0.create_DC1(_2981_v2);
          let _index519 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_3062_v61).length));
          (_3062_v61)[_index519] = (_dafny.ZERO).minus((_3060_v59)[_module.__default.safeIndex(new BigNumber(((_3091_v86).dtor_cf1).length), new BigNumber((_3060_v59).length))]);
          let _index520 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_3062_v61).length));
          (_3062_v61)[_index520] = _module.__default.fm1(globalState);
        }
      }
      let _source48 = _module.__default.fm64(globalState);
      if (_source48.is_DC31) {
        let _3092___mcc_h4 = (_source48).cf50;
        let _3093___mcc_h5 = (_source48).cf51;
        let _3094___mcc_h6 = (_source48).cf52;
        let _3095_cf52 = _3094___mcc_h6;
        let _3096_cf51 = _3093___mcc_h5;
        let _3097_cf50 = _3092___mcc_h4;
        let _3098_v87;
        let _out83;
        _out83 = (_this).m1(p2, globalState);
        _3098_v87 = _out83;
        let _3099_v88;
        _3099_v88 = _dafny.Seq.of(_3095_cf52, _3095_cf52, p2);
        let _3100_v89;
        _3100_v89 = _dafny.Map.Empty.slice().updateUnsafe(_3097_cf50,_2988_v6);
        let _3101_v90;
        _3101_v90 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ihsbrh"),_2988_v6);
        let _3102_v91;
        let _nw501 = Array((new BigNumber(17)).toNumber());
        _nw501[(_dafny.ZERO).toNumber()] = (_2988_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_3098_v87,(_dafny.ZERO).minus(new BigNumber((_3099_v88).length))));
        _nw501[(_dafny.ONE).toNumber()] = ((((_3100_v89).contains((_this).f9)) ? ((_3100_v89).get((_this).f9)) : (_2988_v6))).Merge(_2988_v6);
        _nw501[(new BigNumber(2)).toNumber()] = _2988_v6;
        _nw501[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_2981_v2).length))).Merge(_2988_v6);
        _nw501[(new BigNumber(4)).toNumber()] = _2988_v6;
        _nw501[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(541));
        _nw501[(new BigNumber(6)).toNumber()] = _2988_v6;
        _nw501[(new BigNumber(7)).toNumber()] = (_2988_v6).update(_3098_v87, p2);
        _nw501[(new BigNumber(8)).toNumber()] = (_2988_v6).Merge(_2988_v6);
        _nw501[(new BigNumber(9)).toNumber()] = _2988_v6;
        _nw501[(new BigNumber(10)).toNumber()] = ((_2988_v6).update(!(_3098_v87), p2)).Merge(_2988_v6);
        _nw501[(new BigNumber(11)).toNumber()] = ((((_3101_v90).contains(_2981_v2)) ? ((_3101_v90).get(_2981_v2)) : (_2988_v6))).Merge(_2988_v6);
        _nw501[(new BigNumber(12)).toNumber()] = _2988_v6;
        _nw501[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_3096_cf51),_3095_cf52);
        _nw501[(new BigNumber(14)).toNumber()] = _2988_v6;
        _nw501[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(false),p2);
        _nw501[(new BigNumber(16)).toNumber()] = _2988_v6;
        _3102_v91 = _nw501;
        let _3103_v92;
        let _init76 = function (_3104_i7) {
          return (_this).f8;
        };
        let _nw502 = Array((new BigNumber(7)).toNumber());
        for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw502.length); _i0_76++) {
          _nw502[_i0_76] = _init76(new BigNumber(_i0_76));
        }
        _3103_v92 = _nw502;
        let _3105_v93;
        let _nw503 = new _module.C9();
        _nw503.__ctor(true, p1);
        _3105_v93 = _nw503;
        let _3106_v94;
        _3106_v94 = _dafny.MultiSet.fromElements(_3105_v93, _3105_v93);
        let _index521 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
        (_3103_v92)[_index521] = _module.__default.fm0(new BigNumber((_3106_v94).cardinality()), globalState);
        let _3107_v95;
        _3107_v95 = _dafny.MultiSet.fromElements(_2989_v7, _2989_v7);
        let _index522 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
        let _rhs499 = _3102_v91;
        let _rhs500 = (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_3095_cf52))).IsSubsetOf(((p1) ? (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_3095_cf52))) : (_3107_v95)));
        let _lhs362 = _3103_v92;
        let _lhs363 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
        _3102_v91 = _rhs499;
        _lhs362[_lhs363] = _rhs500;
        (globalState).f0 = p2;
        if ((p2).isLessThan(_3095_cf52)) {
          let _3108_v96;
          let _nw504 = new _module.C5();
          _nw504.__ctor();
          _3108_v96 = _nw504;
          _3108_v96 = ((_dafny.Seq.IsProperPrefixOf(_2981_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(431)), ((_3109_v1) => function (_3110_i8) {
            return _3109_v1;
          })(_2980_v1)))) ? (_3108_v96) : (((!((_this).f8)) ? (_3108_v96) : (_3108_v96))));
          let _3111_v97;
          let _nw505 = new _module.C4();
          _nw505.__ctor();
          _3111_v97 = _nw505;
          let _3112_v98;
          _3112_v98 = _3111_v97;
          _3111_v97 = (_3112_v98);
          let _3113_v99;
          _3113_v99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,p0);
          let _3114_v100;
          _3114_v100 = _dafny.Set.fromElements(p0, !(!((_3105_v93).f10)), (((_3113_v99).contains(false)) ? ((_3113_v99).get(false)) : ((_this).f8)), (_3103_v92)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length))]);
          let _3115_v101;
          let _nw506 = new _module.C12();
          _nw506.__ctor((_2990_v8)[_module.__default.safeIndex(p2, new BigNumber((_2990_v8).length))], (_3103_v92)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length))], (new BigNumber((_3114_v100).length)).isLessThan(_3095_cf52));
          _3115_v101 = _nw506;
          let _3116_v102;
          let _init77 = ((_3117_v1) => function (_3118_i9) {
            return _3117_v1;
          })(_2980_v1);
          let _nw507 = Array((new BigNumber(2)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw507.length); _i0_77++) {
            _nw507[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _3116_v102 = _nw507;
          let _index523 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_3116_v102).length));
          (_3116_v102)[_index523] = _module.__default.fm4(_3098_v87, (_3103_v92)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length))], _3098_v87, globalState);
          let _3119_v103;
          let _nw508 = Array((new BigNumber(2)).toNumber());
          _nw508[(_dafny.ZERO).toNumber()] = new BigNumber((_2988_v6).length);
          _nw508[(_dafny.ONE).toNumber()] = (_3095_cf52).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3095_cf52,_3095_cf52)).length));
          _3119_v103 = _nw508;
          let _index524 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_3119_v103).length));
          (_3119_v103)[_index524] = p2;
          let _index525 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
          let _index526 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_3116_v102).length));
          let _index527 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_3119_v103).length));
          let _rhs501 = (_3115_v101).f11;
          let _rhs502 = _2980_v1;
          let _rhs503 = _module.__default.safeModuloInt((_dafny.ZERO).minus((p2).multipliedBy(new BigNumber((_2981_v2).length))), (_3095_cf52).minus(new BigNumber((_3099_v88).length)));
          let _rhs504 = (_2989_v7).Intersect(_2989_v7);
          let _lhs364 = _3103_v92;
          let _lhs365 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
          let _lhs366 = _3116_v102;
          let _lhs367 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_3116_v102).length));
          let _lhs368 = _3119_v103;
          let _lhs369 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_3119_v103).length));
          _lhs364[_lhs365] = _rhs501;
          _lhs366[_lhs367] = _rhs502;
          _lhs368[_lhs369] = _rhs503;
          _2989_v7 = _rhs504;
          _3097_cf50 = p0;
        } else {
          let _3120_v104;
          _3120_v104 = _dafny.MultiSet.fromElements(_3098_v87);
          let _3121_v105;
          _3121_v105 = _dafny.Seq.of(_3120_v104);
          let _3122_v106;
          let _nw509 = Array((new BigNumber(17)).toNumber());
          _nw509[(_dafny.ZERO).toNumber()] = _3120_v104;
          _nw509[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.FromArray(_2990_v8)).Difference(_3120_v104);
          _nw509[(new BigNumber(2)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_2990_v8);
          _nw509[(new BigNumber(4)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(5)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(6)).toNumber()] = _module.__default.fm41(globalState);
          _nw509[(new BigNumber(7)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(8)).toNumber()] = (_3121_v105)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_3121_v105).length))];
          _nw509[(new BigNumber(9)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(10)).toNumber()] = (_module.D5.create_DC16(_3120_v104)).dtor_cf27;
          _nw509[(new BigNumber(11)).toNumber()] = (_3120_v104).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_3097_cf50)));
          _nw509[(new BigNumber(12)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements((_3105_v93).f10, (_this).f8);
          _nw509[(new BigNumber(14)).toNumber()] = ((_dafny.MultiSet.fromElements(!(p1))).update(_3097_cf50, _module.__default.abs(p2))).Union(_3120_v104);
          _nw509[(new BigNumber(15)).toNumber()] = _3120_v104;
          _nw509[(new BigNumber(16)).toNumber()] = _3120_v104;
          _3122_v106 = _nw509;
          let _index528 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_3122_v106).length));
          (_3122_v106)[_index528] = _dafny.MultiSet.fromElements(_3096_cf51);
          let _3123_v107;
          _3123_v107 = _module.D0.create_DC0(_3103_v92);
          let _index529 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_3122_v106).length));
          let _index530 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
          let _rhs505 = _3120_v104;
          let _rhs506 = (_this).f8;
          let _rhs507 = (_3105_v93).fm7(!(_3098_v87), globalState);
          let _rhs508 = _3123_v107;
          let _lhs370 = _3122_v106;
          let _lhs371 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_3122_v106).length));
          let _lhs372 = _3103_v92;
          let _lhs373 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_3103_v92).length));
          _lhs370[_lhs371] = _rhs505;
          _3098_v87 = _rhs506;
          _lhs372[_lhs373] = _rhs507;
          _3123_v107 = _rhs508;
          let _3124_v108;
          _3124_v108 = _dafny.Map.Empty.slice().updateUnsafe((_3095_cf52).plus(p2),_2980_v1);
          _3124_v108 = (_3124_v108).update(_3095_cf52, _2980_v1);
          let _3125_v109;
          let _nw510 = new _module.C5();
          _nw510.__ctor();
          _3125_v109 = _nw510;
          (globalState).f0 = (_3095_cf52).multipliedBy(new BigNumber((_dafny.Seq.Concat(_2981_v2, _dafny.Seq.UnicodeFromString("a"))).length));
          (globalState).f1 = p2;
        }
      } else {
        let _3126___mcc_h7 = (_source48).cf49;
        let _3127_cf49 = _3126___mcc_h7;
        let _3128_v110;
        let _nw511 = Array((new BigNumber(12)).toNumber()).fill(false);
        _3128_v110 = _nw511;
        let _nw512 = Array((new BigNumber(3)).toNumber()).fill(false);
        _3128_v110 = _nw512;
        let _3129_v111;
        _3129_v111 = true;
        let _index531 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        (_3128_v110)[_index531] = false;
        let _3130_v112;
        _3130_v112 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p2));
        let _index532 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        let _rhs509 = !((_this).f8);
        let _rhs510 = (_3130_v112).IsDisjointFrom(_3130_v112);
        let _rhs511 = (_2990_v8)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_2981_v2, _2981_v2)).length), new BigNumber((_2990_v8).length))];
        let _rhs512 = p2;
        let _lhs374 = _3128_v110;
        let _lhs375 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        let _lhs376 = globalState;
        _3129_v111 = _rhs509;
        _lhs374[_lhs375] = _rhs510;
        _3129_v111 = _rhs511;
        _lhs376.f0 = _rhs512;
        let _index533 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        (_3128_v110)[_index533] = _3129_v111;
        let _3131_v113;
        _3131_v113 = _dafny.Set.fromElements(_3128_v110);
        let _3132_v115;
        _3132_v115 = _module.D10.create_DC30(_3127_cf49);
        let _3133_v116;
        _3133_v116 = _dafny.Seq.of(p2);
        let _pat_let_tv78 = p0;
        let _index534 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        let _index535 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        let _rhs513 = p1;
        let _rhs514 = (_dafny.Set.fromElements(p0, (function (_pat_let76_0) {
          return function (_3135_dt__update__tmp_h3) {
            return function (_pat_let77_0) {
              return function (_3136_dt__update_hcf54_h0) {
                return _module.D11.create_DC33(_3136_dt__update_hcf54_h0, (_3135_dt__update__tmp_h3).dtor_cf55);
              }(_pat_let77_0);
            }(_pat_let_tv78);
          }(_pat_let76_0);
        }(_module.D11.create_DC33((_this).f9, (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll71 = new _dafny.Map();
  for (const _compr_71 of _dafny.IntegerRange(new BigNumber(744), new BigNumber(853))) {
    let _3134_v114 = _compr_71;
    if (((new BigNumber(744)).isLessThanOrEqualTo(_3134_v114)) && ((_3134_v114).isLessThan(new BigNumber(853)))) {
      _coll71.push([(_3134_v114).minus(new BigNumber((_2990_v8).length)),new BigNumber((_2981_v2).length)]);
    }
  }
  return _coll71;
}()).length))))).dtor_cf54)).Union((_dafny.Set.fromElements(true, p1, (_3128_v110)[_module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length))])).Intersect((_3132_v115).dtor_cf49));
        let _rhs515 = (_3131_v113).Difference(_3131_v113);
        let _rhs516 = p2;
        let _rhs517 = !(new BigNumber((_3133_v116).length)).isEqualTo(p2);
        let _lhs377 = _3128_v110;
        let _lhs378 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        let _lhs379 = globalState;
        let _lhs380 = _3128_v110;
        let _lhs381 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_3128_v110).length));
        _lhs377[_lhs378] = _rhs513;
        _3127_cf49 = _rhs514;
        _3131_v113 = _rhs515;
        _lhs379.f0 = _rhs516;
        _lhs380[_lhs381] = _rhs517;
      }
      let _3137_v117;
      _3137_v117 = false;
      let _3138_v118;
      _3138_v118 = _dafny.Set.fromElements((_this).f8, (_this).f8);
      _3137_v117 = (_3138_v118).IsSubsetOf(_3138_v118);
      return;
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
