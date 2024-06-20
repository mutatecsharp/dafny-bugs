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
      return (_dafny.ZERO).minus(new BigNumber(-194));
    };
    static fm1(p0, p1, globalState) {
      return !(((false) ? (((true) ? (false) : (true))) : (_dafny.Seq.IsPrefixOf(_dafny.Seq.of(new BigNumber(506), new BigNumber(281)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, true)).length), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))).length))))));
    };
    static fm2(p0, p1, p2, globalState) {
      if (((true) ? (false) : (false))) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), function (_0_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        });
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sgxcrrbck"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_1_i1) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }));
      }
    };
    static fm3(p0, p1, globalState) {
      return new _dafny.CodePoint('c'.codePointAt(0));
    };
    static fm4(p0, p1, globalState) {
      return _module.D0.create_DC2((_module.D0.create_DC2(_dafny.MultiSet.fromElements(!(false), true), false, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)), false)).dtor_cf4, true, new BigNumber(-858), _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_2_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-240)), function (_3_i1) {
  return new _dafny.CodePoint('p'.codePointAt(0));
})));
    };
    static fm7(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(897)), function (_4_i0) {
        return new BigNumber(720);
      })).length),((false) ? ((_module.D0.create_DC2(_dafny.MultiSet.fromElements(!(true)), true, new BigNumber(611), false)).dtor_cf6) : (new BigNumber(-491))));
    };
    static fm14(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-621)), function (_5_i0) {
        return function () {
          let _coll0 = new _dafny.Set();
          for (const _compr_0 of _dafny.IntegerRange(new BigNumber(869), new BigNumber(-384))) {
            let _6_v0 = _compr_0;
            if (((new BigNumber(869)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(-384)))) {
              _coll0.add((_6_v0).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-635)), function (_7_i1) {
                return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(421), new BigNumber(-384), new BigNumber((_dafny.Seq.UnicodeFromString("jfdgtb")).length), new BigNumber(389))).cardinality()))).length);
              }))).cardinality())));
            }
          }
          return _coll0;
        }();
      })));
    };
    static fm15(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(286)), function (_8_i0) {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false));
      });
    };
    static fm16(globalState) {
      return _dafny.Set.fromElements(((!(!(false))) ? (false) : (!(!(false)))), false, true);
    };
    static fm17(globalState) {
      return _module.D2.create_DC7(!(!(new BigNumber(380)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-846), new BigNumber((_dafny.Seq.of(true)).length))).length))))), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-811), new BigNumber(752))), ((true) ? (!(true)) : (false)), _dafny.Seq.UnicodeFromString("uyiex"));
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length)), (_dafny.MultiSet.fromElements(new BigNumber(329))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.FromArray(_dafny.Seq.of(true)))).length))))).Union((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(671), new BigNumber((_dafny.Seq.of(false)).length))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("yxlacl")).length)), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(511))).length), new BigNumber(604), new BigNumber((_dafny.Set.fromElements(new BigNumber(-660), new BigNumber(685))).length), new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of(new BigNumber(78))).Elements) {
          let _9_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(78)), _9_v0)) {
            _coll1.push([(_9_v0).plus((_module.D20.create_DC40(false, new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).dtor_cf68),new BigNumber(823)]);
          }
        }
        return _coll1;
      }()).length))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), function (_10_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length))))).Union(function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), function (_11_i1) {
          return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(176))).length));
        })).Elements) {
          let _12_v1 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), function (_11_i1) {
            return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(176))).length));
          }), _12_v1)) {
            _coll2.add(_12_v1);
          }
        }
        return _coll2;
      }()));
    };
    static fm21(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(793), new BigNumber(442))) {
          let _13_v0 = _compr_3;
          if (((new BigNumber(793)).isLessThanOrEqualTo(_13_v0)) && ((_13_v0).isLessThan(new BigNumber(442)))) {
            _coll3.add((_13_v0).minus(new BigNumber((_dafny.Seq.of(!(false))).length)));
          }
        }
        return _coll3;
      }(),new BigNumber((_dafny.Set.fromElements(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length),false)).length)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-545)))),new BigNumber(51)));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),true)).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),false)) : (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true))));
    };
    static fm23(p0, globalState) {
      if (false) {
        return (function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(244), new BigNumber(762))) {
            let _14_v0 = _compr_4;
            if (((new BigNumber(244)).isLessThanOrEqualTo(_14_v0)) && ((_14_v0).isLessThan(new BigNumber(762)))) {
              _coll4.add(_module.__default.safeDivisionInt(_14_v0, new BigNumber((_dafny.Set.fromElements(false)).length)));
            }
          }
          return _coll4;
        }()).Intersect(_dafny.Set.fromElements(new BigNumber(-694), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-387)), function (_15_i0) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length))).length)));
      } else {
        return (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(!(true), false, false, true)).cardinality()))).Intersect(_dafny.Set.fromElements(new BigNumber(348)));
      }
    };
    static fm24(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber(91), new BigNumber(768)),new _dafny.CodePoint('o'.codePointAt(0)));
    };
    static fm25(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-954))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false)).length)));
    };
    static fm26(p0, p1, p2, globalState) {
      return (_module.D5.create_DC13(_dafny.Seq.of(new BigNumber(-784)))).dtor_cf25;
    };
    static fm27(p0, globalState) {
      return ((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(284), new BigNumber(491))) {
          let _16_v0 = _compr_5;
          if (((new BigNumber(284)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(491)))) {
            _coll5.push([_module.__default.safeDivisionInt(_16_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-314))).length)),!(false)]);
          }
        }
        return _coll5;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(846), new BigNumber(273)))).cardinality()),(_module.D2.create_DC7(false, new BigNumber(544), false, _dafny.Seq.UnicodeFromString("xoq"))).dtor_cf11))).Merge((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-785),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), function (_17_i0) {
          return new BigNumber(277);
        })).length))).Keys.Elements) {
          let _18_v1 = _compr_6;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-785),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), function (_17_i0) {
            return new BigNumber(277);
          })).length))).contains(_18_v1)) {
            _coll6.push([(_18_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(245))).length)),true]);
          }
        }
        return _coll6;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(39),false)));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-791)), function (_19_i0) {
        return _dafny.Seq.UnicodeFromString("ieinao");
      });
    };
    static fm31(globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-281)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(-172))).length)))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(386), new BigNumber(340), new BigNumber((_dafny.Seq.UnicodeFromString("ropewiho")).length), new BigNumber((_dafny.Seq.UnicodeFromString("feeuekkg")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(373),new BigNumber(155))).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber((_dafny.Set.fromElements(true)).length))).length))))).Union((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(859)), function (_20_i0) {
        return new BigNumber((_dafny.Set.fromElements(new BigNumber(-885), new BigNumber(23))).length);
      })).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), function (_21_i1) {
        return new BigNumber(998);
      }))).length))).length)))).Union(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-50)), function (_22_i2) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("irwxgif")).length);
      }))));
    };
    static fm32(globalState) {
      return _dafny.Seq.of((_dafny.Set.fromElements(new BigNumber(34))).Intersect(function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-128),new BigNumber(460))).Keys.Elements) {
            let _23_v0 = _compr_8;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-128),new BigNumber(460))).contains(_23_v0)) {
              _coll8.push([(_23_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(629))).length))),_dafny.Seq.of(new BigNumber(363))]);
            }
          }
          return _coll8;
        }()).Keys.Elements) {
          let _24_v1 = _compr_7;
          if ((function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-128),new BigNumber(460))).Keys.Elements) {
              let _23_v0 = _compr_9;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-128),new BigNumber(460))).contains(_23_v0)) {
                _coll9.push([(_23_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(629))).length))),_dafny.Seq.of(new BigNumber(363))]);
              }
            }
            return _coll9;
          }()).contains(_24_v1)) {
            _coll7.add(_module.__default.safeModuloInt(_24_v1, new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)));
          }
        }
        return _coll7;
      }()), function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!(true), false, false, false, !(true))).length))).Elements) {
          let _25_v2 = _compr_10;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!(true), false, false, false, !(true))).length))).contains(_25_v2)) {
            _coll10.add(_module.__default.safeModuloInt(_25_v2, new BigNumber(168)));
          }
        }
        return _coll10;
      }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ikr")).length)), function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber((_dafny.Seq.UnicodeFromString("eneorssfn")).length))).Keys.Elements) {
          let _26_v3 = _compr_11;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber((_dafny.Seq.UnicodeFromString("eneorssfn")).length))).contains(_26_v3)) {
            _coll11.add((_26_v3).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))));
          }
        }
        return _coll11;
      }(), function () {
        let _coll12 = new _dafny.Set();
        for (const _compr_12 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),!(false))).length))).Elements) {
          let _27_v4 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),!(false))).length)), _27_v4)) {
            _coll12.add(_module.__default.safeModuloInt(_27_v4, new BigNumber(-749)));
          }
        }
        return _coll12;
      }());
    };
    static fm33(p0, globalState) {
      return (_dafny.MultiSet.fromElements(!(false), false)).Difference(_dafny.MultiSet.fromElements(true));
    };
    static fm34(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,!(!(true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(true)));
    };
    static fm35(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(351)), function (_28_i0) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).cardinality())),new BigNumber(280))).length))).length);
})), _module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), function (_29_i1) {
  return new BigNumber(-705);
})), _module.D5.create_DC13(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xs")).length)),false)).length)))), _dafny.Seq.of(_module.D5.create_DC13(_dafny.Seq.of(new BigNumber(872))))), _dafny.Seq.of(_module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(675)), function (_30_i2) {
  return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
    let _coll13 = new _dafny.Map();
    for (const _compr_13 of (_dafny.Set.fromElements(new BigNumber(304))).Elements) {
      let _31_v0 = _compr_13;
      if ((_dafny.Set.fromElements(new BigNumber(304))).contains(_31_v0)) {
        _coll13.push([_module.__default.safeModuloInt(_31_v0, new BigNumber((_dafny.Seq.UnicodeFromString("lbekjpmmt")).length)),new BigNumber((function () {
          let _coll14 = new _dafny.Set();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(361), new BigNumber(168))) {
            let _32_v1 = _compr_14;
            if (((new BigNumber(361)).isLessThanOrEqualTo(_32_v1)) && ((_32_v1).isLessThan(new BigNumber(168)))) {
              _coll14.add((_32_v1).minus(new BigNumber(833)));
            }
          }
          return _coll14;
        }()).length)]);
      }
    }
    return _coll13;
  }()).length))).cardinality());
}))));
    };
    static fm36(p0, globalState) {
      return _dafny.Seq.of((new BigNumber((_dafny.Seq.UnicodeFromString("nurarvw")).length)).plus(new BigNumber(-32)), (new BigNumber(219)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("pybfvhx")).length)));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      if (!(false)) {
        return _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))).cardinality()));
      } else {
        return function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-57), new BigNumber(768))) {
            let _33_v0 = _compr_15;
            if (((new BigNumber(-57)).isLessThanOrEqualTo(_33_v0)) && ((_33_v0).isLessThan(new BigNumber(768)))) {
              _coll15.add((_33_v0).multipliedBy(new BigNumber(307)));
            }
          }
          return _coll15;
        }();
      }
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("teicle")).length))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("levf")).length))));
    };
    static fm39(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!(false), false), _dafny.Seq.of((_module.D22.create_DC46(false)).dtor_cf76)), _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(false)));
    };
    static fm40(p0, globalState) {
      return _module.D0.create_DC0(((!(true)) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(913),new BigNumber(-795))).length)) : (new BigNumber(764))));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC10(_dafny.Set.fromElements(true, true));
    };
    static fm44(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, !(true)), _dafny.Seq.of(false)), _dafny.Seq.of(false));
    };
    static fm45(p0, p1, globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))) : (_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))));
    };
    static fm46(globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(622))).cardinality())).isLessThan((_module.D8.create_DC19(_dafny.Seq.UnicodeFromString("vapfcanir"), new BigNumber(939))).dtor_cf35));
    };
    static fm49(p0, globalState) {
      return _module.D2.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-885)), function (_34_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
}));
    };
    static fm50(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(!(false))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("utwya")).length)));
    };
    static fm51(globalState) {
      return ((function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Seq.of(_module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("o")).length))), _module.D0.create_DC0(new BigNumber(444)), _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("tqobdcs")).length)))).Elements) {
          let _35_v0 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("o")).length))), _module.D0.create_DC0(new BigNumber(444)), _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("tqobdcs")).length))), _35_v0)) {
            _coll16.push([_35_v0,new _dafny.CodePoint('p'.codePointAt(0))]);
          }
        }
        return _coll16;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-783)))).length)),new _dafny.CodePoint('l'.codePointAt(0))))).Merge(function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jminm")).length)), new BigNumber(48))).length)),new BigNumber((_dafny.Seq.of(!(!(false)))).length))).Keys.Elements) {
          let _36_v1 = _compr_17;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jminm")).length)), new BigNumber(48))).length)),new BigNumber((_dafny.Seq.of(!(!(false)))).length))).contains(_36_v1)) {
            _coll17.push([_36_v1,new _dafny.CodePoint('v'.codePointAt(0))]);
          }
        }
        return _coll17;
      }());
    };
    static fm52(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality())))).Merge(function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-508), new BigNumber(187))) {
          let _37_v0 = _compr_18;
          if (((new BigNumber(-508)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(187)))) {
            _coll18.push([(_37_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-834), new BigNumber(291), new BigNumber(447), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('f'.codePointAt(0)))).length))).length)),new BigNumber(58)]);
          }
        }
        return _coll18;
      }())).Merge(function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-31),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).cardinality())))).length)), function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of _dafny.IntegerRange(new BigNumber(405), new BigNumber(737))) {
            let _38_v2 = _compr_20;
            if (((new BigNumber(405)).isLessThanOrEqualTo(_38_v2)) && ((_38_v2).isLessThan(new BigNumber(737)))) {
              _coll20.push([_module.__default.safeDivisionInt(_38_v2, new BigNumber(754)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(666))).length)]);
            }
          }
          return _coll20;
        }())).length),new BigNumber(988))).Keys.Elements) {
          let _39_v1 = _compr_19;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-31),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).cardinality())))).length)), function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(405), new BigNumber(737))) {
              let _38_v2 = _compr_21;
              if (((new BigNumber(405)).isLessThanOrEqualTo(_38_v2)) && ((_38_v2).isLessThan(new BigNumber(737)))) {
                _coll21.push([_module.__default.safeDivisionInt(_38_v2, new BigNumber(754)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(666))).length)]);
              }
            }
            return _coll21;
          }())).length),new BigNumber(988))).contains(_39_v1)) {
            _coll19.push([_module.__default.safeModuloInt(_39_v1, new BigNumber((_dafny.Seq.of(new BigNumber(37), new BigNumber(64))).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), function (_40_i0) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            })).length)]);
          }
        }
        return _coll19;
      }());
    };
    static fm53(globalState) {
      return (_dafny.Set.fromElements(false, !(!(!(false))), true, true)).Union(((_module.D4.create_DC10(_dafny.Set.fromElements(true, !(true)))).dtor_cf18).Union(_dafny.Set.fromElements(false, true)));
    };
    static fm54(p0, p1, p2, globalState) {
      if ((!(true)) === (false)) {
        return _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-476)), new BigNumber(634), new BigNumber(444));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(990), new BigNumber(676))).length)), new BigNumber(891), (_dafny.ZERO).minus((_module.D9.create_DC22(new BigNumber(693), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(317),_dafny.Seq.UnicodeFromString("wmcxvlhm")))).dtor_cf40))).length))), _dafny.Seq.of(new BigNumber(-25), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(855))).length)));
      }
    };
    static fm55(p0, p1, p2, globalState) {
      return _module.D5.create_DC14((false) && (false));
    };
    static fm56(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(850))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length)))).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),new _dafny.CodePoint('v'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.of(true, false, true, true)).length), new BigNumber(591)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-448), new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,function () {
        let _coll22 = new _dafny.Set();
        for (const _compr_22 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qac")).length),true)).Keys.Elements) {
          let _41_v0 = _compr_22;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qac")).length),true)).contains(_41_v0)) {
            _coll22.add(_module.__default.safeModuloInt(_41_v0, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ylqio")).length)))).cardinality())));
          }
        }
        return _coll22;
      }())).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(436)))).length)))).length), new BigNumber(507), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),new BigNumber((_dafny.Set.fromElements(true)).length))).length), new BigNumber(96)))));
    };
    static fm57(globalState) {
      return _module.D1.create_DC4(new _dafny.CodePoint('j'.codePointAt(0)));
    };
    static fm58(p0, p1, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), function (_42_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("kbab")).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.MultiSet.FromArray(_dafny.Seq.of(true)))).length))));
    };
    static fm59(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(true, true)),(new BigNumber(-846)).plus(new BigNumber(141)));
    };
    static fm60(p0, p1, globalState) {
      return _module.D8.create_DC19(_dafny.Seq.UnicodeFromString("scicnjy"), new BigNumber((_dafny.Seq.of(new BigNumber(-492), new BigNumber(-464), new BigNumber(441))).length));
    };
    static fm61(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),_dafny.Seq.of(true, true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(939),_dafny.Seq.of(false, false))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll23 = new _dafny.Set();
        for (const _compr_23 of _dafny.IntegerRange(new BigNumber(-130), new BigNumber(889))) {
          let _43_v0 = _compr_23;
          if (((new BigNumber(-130)).isLessThanOrEqualTo(_43_v0)) && ((_43_v0).isLessThan(new BigNumber(889)))) {
            _coll23.add((_43_v0).multipliedBy(new BigNumber(19)));
          }
        }
        return _coll23;
      }()).length)),_dafny.Seq.of(true, true, true))));
    };
    static fm62(p0, globalState) {
      return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length),false));
    };
    static fm63(globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kici"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), function (_44_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })));
    };
    static m0(globalState) {
      let _45_v0;
      _45_v0 = false;
      let _46_i0;
      _46_i0 = _dafny.ZERO;
      L0: {
        while (_45_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_46_i0)) {
              break L0;
            }
            _46_i0 = (_46_i0).plus(_dafny.ONE);
            if (_45_v0) {
              let _47_v1;
              _47_v1 = _dafny.Seq.of(_45_v0, _45_v0);
              let _48_v2;
              _48_v2 = new BigNumber(-142);
              (globalState).f9 = _module.__default.safeModuloInt(new BigNumber(772), (new BigNumber((_47_v1).length)).plus(_48_v2));
              let _49_v3;
              _49_v3 = new _dafny.CodePoint('m'.codePointAt(0));
              let _50_v4;
              _50_v4 = _dafny.Set.fromElements(_module.__default.fm1(_45_v0, _49_v3, globalState), _45_v0, _45_v0);
              (globalState).f8 = (_50_v4).IsDisjointFrom(_50_v4);
              let _51_v5;
              let _nw0 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _51_v5 = _nw0;
              let _52_v6;
              let _nw1 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
              _52_v6 = _nw1;
              let _53_v7;
              let _nw2 = new _module.C5();
              _nw2.__ctor(_52_v6);
              _53_v7 = _nw2;
              let _54_v8;
              _54_v8 = _dafny.MultiSet.fromElements(_53_v7, _53_v7);
              let _55_v9;
              _55_v9 = _dafny.Seq.of(_49_v3);
              let _56_v10;
              _56_v10 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_54_v8).cardinality())).isLessThanOrEqualTo(_48_v2),_55_v9);
              let _57_v11;
              _57_v11 = _dafny.Set.fromElements(new BigNumber((_55_v9).length), new BigNumber((_47_v1).length));
              let _58_v12;
              _58_v12 = _module.D21.create_DC42(_51_v5);
              let _59_v13;
              _59_v13 = _dafny.Seq.of(_56_v10, _56_v10);
              let _rhs0 = _45_v0;
              let _rhs1 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_57_v11).length)), _48_v2)).multipliedBy(_48_v2);
              let _rhs2 = (_58_v12).dtor_cf70;
              let _rhs3 = (_59_v13)[_module.__default.safeIndex(_48_v2, new BigNumber((_59_v13).length))];
              let _lhs0 = globalState;
              let _lhs1 = globalState;
              _lhs0.f8 = _rhs0;
              _lhs1.f9 = _rhs1;
              _51_v5 = _rhs2;
              _56_v10 = _rhs3;
              (globalState).f9 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_47_v1, _module.__default.safeIndex(_48_v2, new BigNumber((_47_v1).length)), _45_v0), _dafny.Seq.of(_45_v0))).length)).minus(_48_v2);
              let _60_v14;
              let _nw3 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
              _60_v14 = _nw3;
              let _index0 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_60_v14).length));
              (_60_v14)[_index0] = _48_v2;
              let _61_v15;
              _61_v15 = _dafny.Seq.of(_48_v2);
              let _index1 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_60_v14).length));
              (_60_v14)[_index1] = (_61_v15)[_module.__default.safeIndex(new BigNumber(94), new BigNumber((_61_v15).length))];
            } else {
              let _62_v16;
              _62_v16 = _dafny.Seq.of(_45_v0, _45_v0);
              let _63_v17;
              _63_v17 = new BigNumber(55);
              let _64_v18;
              _64_v18 = _module.D8.create_DC20(_module.D0.create_DC2(_dafny.MultiSet.FromArray(_62_v16), _45_v0, _63_v17, _45_v0), _45_v0, _45_v0);
              _45_v0 = (_64_v18).dtor_cf37;
              let _65_v19;
              let _init0 = ((_66_v0) => function (_67_i1) {
                return (_dafny.Set.fromElements(_66_v0, _66_v0, _66_v0, true)).Intersect(_dafny.Set.fromElements(!(_66_v0), _66_v0, !(_66_v0), !(_66_v0), _66_v0));
              })(_45_v0);
              let _nw4 = Array((new BigNumber(10)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw4.length); _i0_0++) {
                _nw4[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _65_v19 = _nw4;
              let _68_v20;
              _68_v20 = _dafny.Set.fromElements(_45_v0, _45_v0, _45_v0, _45_v0, _45_v0);
              let _index2 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_65_v19).length));
              (_65_v19)[_index2] = _68_v20;
              let _index3 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_65_v19).length));
              (_65_v19)[_index3] = _68_v20;
              let _69_v21;
              _69_v21 = _dafny.Seq.UnicodeFromString("vtof");
              let _70_v22;
              let _nw5 = new _module.C6();
              _nw5.__ctor(_63_v17, _69_v21);
              _70_v22 = _nw5;
              _63_v17 = (_70_v22).f15;
              (globalState).f9 = (_70_v22).f15;
            }
            let _71_v23;
            _71_v23 = new BigNumber(885);
            let _72_v24;
            _72_v24 = _dafny.Set.fromElements((_dafny.ZERO).minus((_71_v23).plus(_71_v23)), (_dafny.ZERO).minus(new BigNumber(-673)));
            _72_v24 = (_72_v24).Difference((_dafny.Set.fromElements(_71_v23)).Intersect(_72_v24));
            if (!(!(true) || (_45_v0))) {
              let _73_v25;
              let _init1 = ((_74_v0) => function (_75_i2) {
                return _74_v0;
              })(_45_v0);
              let _nw6 = Array((new BigNumber(8)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
                _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _73_v25 = _nw6;
              let _index4 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_73_v25).length));
              (_73_v25)[_index4] = false;
              let _76_v26;
              _76_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_73_v25);
              let _77_v27;
              _77_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-216),_45_v0);
              let _78_v28;
              _78_v28 = _dafny.Set.fromElements((((_76_v26).contains((((_77_v27).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_71_v23)))) ? ((_77_v27).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_71_v23)))) : (_45_v0)))) ? ((_76_v26).get((((_77_v27).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_71_v23)))) ? ((_77_v27).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_71_v23)))) : (_45_v0)))) : (_73_v25)), _73_v25, _73_v25, _73_v25, _73_v25);
              let _79_v29;
              _79_v29 = _dafny.MultiSet.fromElements(new BigNumber(672));
              let _80_v30;
              _80_v30 = new _dafny.CodePoint('h'.codePointAt(0));
              let _81_v31;
              _81_v31 = _dafny.MultiSet.fromElements(_45_v0);
              let _index5 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_73_v25).length));
              let _rhs4 = new BigNumber((((_45_v0) ? (_78_v28) : ((_78_v28).Union(_78_v28)))).length);
              let _rhs5 = (_79_v29).IsSubsetOf(_79_v29);
              let _rhs6 = _module.__default.fm1(_45_v0, _80_v30, globalState);
              let _rhs7 = !(_81_v31).contains(true);
              let _lhs2 = globalState;
              let _lhs3 = _73_v25;
              let _lhs4 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_73_v25).length));
              let _lhs5 = globalState;
              _lhs2.f9 = _rhs4;
              _45_v0 = _rhs5;
              _lhs3[_lhs4] = _rhs6;
              _lhs5.f8 = _rhs7;
              _45_v0 = _45_v0;
              (globalState).f9 = _71_v23;
              let _82_v32;
              _82_v32 = _module.D0.create_DC0(_71_v23);
              let _83_v33;
              let _nw7 = new _module.C1();
              _nw7.__ctor(_45_v0, _45_v0, _82_v32, _45_v0);
              _83_v33 = _nw7;
              let _84_v34;
              _84_v34 = _module.D1.create_DC4(_80_v30);
              let _85_v35;
              _85_v35 = _dafny.Seq.UnicodeFromString("enequbdi");
              let _86_v36;
              _86_v36 = _dafny.Seq.of(_85_v35);
              let _87_v37;
              let _nw8 = new _module.C6();
              _nw8.__ctor((_83_v33).fm9(_84_v34, _72_v24, globalState), (_86_v36)[_module.__default.safeIndex(_71_v23, new BigNumber((_86_v36).length))]);
              _87_v37 = _nw8;
            } else {
              let _88_v38;
              _88_v38 = _dafny.Map.Empty.slice().updateUnsafe(_71_v23,!(_45_v0));
              let _89_v40;
              let _nw9 = Array((new BigNumber(13)).toNumber());
              _nw9[(_dafny.ZERO).toNumber()] = _88_v38;
              _nw9[(_dafny.ONE).toNumber()] = _88_v38;
              _nw9[(new BigNumber(2)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(3)).toNumber()] = function () {
                let _coll24 = new _dafny.Map();
                for (const _compr_24 of _dafny.IntegerRange(new BigNumber(647), new BigNumber(676))) {
                  let _90_v39 = _compr_24;
                  if (((new BigNumber(647)).isLessThanOrEqualTo(_90_v39)) && ((_90_v39).isLessThan(new BigNumber(676)))) {
                    _coll24.push([(_90_v39).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(174)), function (_91_i3) {
                      return new _dafny.CodePoint('o'.codePointAt(0));
                    })).length)),true]);
                  }
                }
                return _coll24;
              }();
              _nw9[(new BigNumber(4)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(5)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(6)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(7)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(8)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(9)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(10)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(11)).toNumber()] = _88_v38;
              _nw9[(new BigNumber(12)).toNumber()] = _88_v38;
              _89_v40 = _nw9;
              let _92_v41;
              let _nw10 = new _module.C2();
              _nw10.__ctor(_89_v40);
              _92_v41 = _nw10;
              _92_v41 = _92_v41;
              _71_v23 = ((_45_v0) ? (_71_v23) : ((_71_v23).multipliedBy(_module.__default.fm0(_45_v0, globalState))));
              let _93_v42;
              let _nw11 = new _module.C3();
              _nw11.__ctor(_45_v0, !((_45_v0) && (_45_v0)), _89_v40);
              _93_v42 = _nw11;
              _93_v42 = _93_v42;
              let _94_v43;
              _94_v43 = _dafny.Seq.UnicodeFromString("jxobqbdo");
              let _95_v44;
              _95_v44 = new _dafny.CodePoint('m'.codePointAt(0));
              let _96_v45;
              _96_v45 = _dafny.MultiSet.fromElements(_71_v23);
              let _97_v46;
              _97_v46 = _dafny.Map.Empty.slice().updateUnsafe(_71_v23,new BigNumber((_88_v38).length));
              let _98_v47;
              let _nw12 = Array((new BigNumber(23)).toNumber());
              _nw12[(_dafny.ZERO).toNumber()] = (_92_v41).fm19(globalState);
              _nw12[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_94_v43, _module.__default.safeIndex(_71_v23, new BigNumber((_94_v43).length)), _95_v44), _94_v43);
              _nw12[(new BigNumber(2)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(3)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_94_v43, _94_v43);
              _nw12[(new BigNumber(5)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(6)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("aamc");
              _nw12[(new BigNumber(8)).toNumber()] = _module.__default.fm2(_71_v23, _96_v45, _97_v46, globalState);
              _nw12[(new BigNumber(9)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("s");
              _nw12[(new BigNumber(11)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("mrwlfrsf");
              _nw12[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_94_v43, _94_v43), _module.__default.safeIndex(_71_v23, new BigNumber((_dafny.Seq.Concat(_94_v43, _94_v43)).length)), _95_v44);
              _nw12[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_94_v43, _94_v43);
              _nw12[(new BigNumber(15)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), function (_99_i4) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              });
              _nw12[(new BigNumber(17)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(18)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(19)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(20)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(21)).toNumber()] = _94_v43;
              _nw12[(new BigNumber(22)).toNumber()] = _94_v43;
              _98_v47 = _nw12;
              let _index6 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_98_v47).length));
              (_98_v47)[_index6] = _dafny.Seq.Concat(_94_v43, _94_v43);
              let _100_v48;
              _100_v48 = _module.D2.create_DC5(_94_v43);
              let _index7 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_98_v47).length));
              let _rhs8 = _dafny.Seq.Concat(_94_v43, _94_v43);
              let _rhs9 = _module.__default.fm1(_module.__default.fm1(_45_v0, _95_v44, globalState), _95_v44, globalState);
              let _rhs10 = _dafny.Seq.IsPrefixOf(_94_v43, (_100_v48).dtor_cf10);
              let _lhs6 = _98_v47;
              let _lhs7 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_98_v47).length));
              let _lhs8 = globalState;
              let _lhs9 = globalState;
              _lhs6[_lhs7] = _rhs8;
              _lhs8.f8 = _rhs9;
              _lhs9.f8 = _rhs10;
              let _101_v49;
              let _nw13 = Array((new BigNumber(5)).toNumber());
              _101_v49 = _nw13;
              let _102_v50;
              _102_v50 = _dafny.Map.Empty.slice().updateUnsafe(_101_v49,(_dafny.ZERO).minus(_71_v23));
              let _103_v51;
              _103_v51 = _module.D22.create_DC45(_101_v49);
              _102_v50 = (_102_v50).update((_103_v51).dtor_cf75, new BigNumber(-627));
            }
            (globalState).f8 = _45_v0;
          }
        }
      }
      let _104_v52;
      let _nw14 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
      _104_v52 = _nw14;
      let _105_v53;
      let _nw15 = new _module.C3();
      _nw15.__ctor(_45_v0, false, _104_v52);
      _105_v53 = _nw15;
      let _106_v54;
      _106_v54 = new BigNumber(546);
      let _hi0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("nklwdo")).length), new BigNumber(124)));
      for (let _107_i5 = _106_v54; _107_i5.isLessThan(_hi0); _107_i5 = _107_i5.plus(_dafny.ONE)) {
        let _108_v55;
        let _nw16 = new _module.C2();
        _nw16.__ctor(_104_v52);
        _108_v55 = _nw16;
        let _109_v56;
        _109_v56 = _dafny.Seq.UnicodeFromString("rilg");
        let _110_v57;
        _110_v57 = _dafny.Seq.of(_105_v53.f21, _105_v53.f21);
        let _111_v58;
        _111_v58 = _dafny.MultiSet.fromElements(_105_v53.f21);
        let _112_v59;
        _112_v59 = _module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_111_v58).cardinality())));
        let _113_v60;
        _113_v60 = new _dafny.CodePoint('x'.codePointAt(0));
        let _114_v61;
        let _nw17 = new _module.C1();
        _nw17.__ctor((_106_v54).isLessThan((_dafny.ZERO).minus(new BigNumber((_109_v56).length))), (_110_v57)[_module.__default.safeIndex(_106_v54, new BigNumber((_110_v57).length))], _112_v59, _dafny.Seq.contains(_109_v56, _113_v60));
        _114_v61 = _nw17;
        _114_v61 = _114_v61;
        if (((((_114_v61).f13) ? (new BigNumber(939)) : (_106_v54))).isLessThanOrEqualTo(_107_i5)) {
          let _115_v62;
          _115_v62 = _dafny.Map.Empty.slice().updateUnsafe(_107_i5,(_105_v53).f22);
          (globalState).f9 = new BigNumber(((_115_v62).Merge(_115_v62)).length);
          let _116_v63;
          let _init2 = ((_117_v54) => function (_118_i6) {
            return (_118_i6).multipliedBy(_117_v54);
          })(_106_v54);
          let _nw18 = Array((new BigNumber(17)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw18.length); _i0_2++) {
            _nw18[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _116_v63 = _nw18;
          let _119_v64;
          _119_v64 = _dafny.Seq.of(_115_v62, _115_v62);
          let _120_v65;
          _120_v65 = _dafny.Set.fromElements(new BigNumber((_119_v64).length), _106_v54);
          let _121_v66;
          let _nw19 = new _module.C5();
          _nw19.__ctor(_104_v52);
          _121_v66 = _nw19;
          let _122_v67;
          _122_v67 = _dafny.Map.Empty.slice().updateUnsafe(_121_v66,_107_i5);
          let _123_v68;
          _123_v68 = _dafny.MultiSet.fromElements(_122_v67, _122_v67, _122_v67);
          let _rhs11 = _116_v63;
          let _rhs12 = _109_v56;
          let _rhs13 = (_120_v65).IsSubsetOf(_120_v65);
          let _rhs14 = _module.__default.safeDivisionInt((new BigNumber((_109_v56).length)).minus(new BigNumber((_123_v68).cardinality())), _106_v54);
          let _rhs15 = _dafny.Seq.Concat(_110_v57, _dafny.Seq.Concat(_module.__default.fm44(_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_124_v57) => function (_125_i7) {
            return _124_v57;
          })(_110_v57)), _113_v60, globalState), _110_v57));
          let _lhs10 = globalState;
          let _lhs11 = globalState;
          _116_v63 = _rhs11;
          _109_v56 = _rhs12;
          _lhs10.f8 = _rhs13;
          _lhs11.f9 = _rhs14;
          _110_v57 = _rhs15;
          let _126_v69;
          let _nw20 = Array((new BigNumber(19)).toNumber()).fill(false);
          _126_v69 = _nw20;
          let _127_v70;
          let _nw21 = new _module.C4();
          _nw21.__ctor(_104_v52);
          _127_v70 = _nw21;
          let _128_v71;
          _128_v71 = _dafny.Map.Empty.slice().updateUnsafe((_114_v61).f13,_127_v70);
          let _129_v72;
          _129_v72 = _dafny.MultiSet.fromElements(_116_v63, _116_v63);
          let _130_v73;
          _130_v73 = _module.D0.create_DC1(_129_v72, new BigNumber((_110_v57).length), _106_v54);
          let _index8 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_126_v69).length));
          (_126_v69)[_index8] = !(new BigNumber((_128_v71).length)).isEqualTo((_130_v73).dtor_cf3);
          let _index9 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_126_v69).length));
          (_126_v69)[_index9] = _module.__default.fm1((_114_v61).f13, _113_v60, globalState);
          let _rhs16 = _107_i5;
          let _rhs17 = (_105_v53).f22;
          let _lhs12 = globalState;
          let _lhs13 = _105_v53;
          _lhs12.f9 = _rhs16;
          _lhs13.f21 = _rhs17;
          let _131_v74;
          let _nw22 = new _module.C6();
          _nw22.__ctor(_106_v54, _109_v56);
          _131_v74 = _nw22;
        } else {
          let _132_v75;
          _132_v75 = _dafny.Map.Empty.slice().updateUnsafe((_114_v61).f13,(_105_v53).f22);
          let _133_v76;
          _133_v76 = _dafny.MultiSet.fromElements(_132_v75);
          _133_v76 = (_133_v76).Intersect((((_105_v53).f22) ? ((_133_v76).update(_dafny.Map.Empty.slice().updateUnsafe(_105_v53.f21,(_105_v53).f22), _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_109_v56).length))))) : (_133_v76)));
          (globalState).f9 = _module.__default.safeDivisionInt(_106_v54, _106_v54);
          let _134_v77;
          _134_v77 = _dafny.MultiSet.fromElements(_106_v54);
          let _135_v78;
          _135_v78 = _dafny.Map.Empty.slice().updateUnsafe(_107_i5,_106_v54);
          let _136_v79;
          _136_v79 = _dafny.Map.Empty.slice().updateUnsafe(_106_v54,_135_v78);
          let _137_v80;
          _137_v80 = _135_v78;
          _109_v56 = _module.__default.fm2(_106_v54, _134_v77, (((_136_v79).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_105_v53.f21,_113_v60)).length))) ? ((_136_v79).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_105_v53.f21,_113_v60)).length))) : ((_137_v80))), globalState);
          let _138_v81;
          let _init3 = ((_139_v53) => function (_140_i8) {
            return _139_v53.f21;
          })(_105_v53);
          let _nw23 = Array((new BigNumber(10)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw23.length); _i0_3++) {
            _nw23[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _138_v81 = _nw23;
          let _index10 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_138_v81).length));
          (_138_v81)[_index10] = _105_v53.f21;
          let _index11 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_138_v81).length));
          (_138_v81)[_index11] = (_105_v53).f22;
          let _141_v82;
          _141_v82 = _dafny.Set.fromElements((_105_v53).f22);
          let _142_v83;
          _142_v83 = _dafny.Map.Empty.slice().updateUnsafe(!(_141_v82).contains((_114_v61).f13),_109_v56);
          _142_v83 = (_142_v83).update((_114_v61).f13, _109_v56);
        }
        let _143_v84;
        _143_v84 = _dafny.Map.Empty.slice().updateUnsafe(!(_45_v0) || ((_105_v53).f22),_105_v53.f21);
        let _144_v85;
        _144_v85 = _143_v84;
        _143_v84 = (_module.__default.fm34(!(_105_v53.f21), _dafny.MultiSet.fromElements((_105_v53).f22), globalState)).Merge((_143_v84).Merge((_144_v85)));
      }
      let _145_v86;
      _145_v86 = _dafny.Set.fromElements(_106_v54);
      let _146_v87;
      _146_v87 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_106_v54, new BigNumber((_145_v86).length)),_45_v0);
      if ((((_146_v87).contains(_106_v54)) ? ((_146_v87).get(_106_v54)) : ((_45_v0) === (false)))) {
        _106_v54 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(366)), function (_147_i9) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("sp"))).length)).minus(_module.__default.safeModuloInt(_106_v54, _106_v54));
        (_105_v53).f21 = _45_v0;
        (_105_v53).f21 = ((_45_v0) ? (!(_105_v53.f21) || (_105_v53.f21)) : (_105_v53.f21));
        _146_v87 = (_146_v87).update(_106_v54, _105_v53.f21);
        let _148_v88;
        let _init4 = function (_149_i10) {
          return _module.__default.safeModuloInt(_149_i10, new BigNumber(650));
        };
        let _nw24 = Array((new BigNumber(24)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw24.length); _i0_4++) {
          _nw24[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _148_v88 = _nw24;
        let _150_v89;
        _150_v89 = _dafny.Map.Empty.slice().updateUnsafe(_106_v54,_148_v88);
        let _151_v90;
        _151_v90 = _module.D9.create_DC21((((_150_v89).contains(_106_v54)) ? ((_150_v89).get(_106_v54)) : (_148_v88)));
        let _pat_let_tv0 = _148_v88;
        let _source0 = function (_pat_let0_0) {
          return function (_152_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_153_dt__update_hcf39_h0) {
                return _module.D9.create_DC21(_153_dt__update_hcf39_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_151_v90);
        if (_source0.is_DC22) {
          let _154___mcc_h0 = (_source0).cf40;
          let _155___mcc_h1 = (_source0).cf41;
          let _156_cf41 = _155___mcc_h1;
          let _157_cf40 = _154___mcc_h0;
          let _158_v91;
          _158_v91 = _dafny.Seq.of((_105_v53).f22, _105_v53.f21);
          let _index12 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_148_v88).length));
          (_148_v88)[_index12] = _module.__default.fm0((_158_v91)[_module.__default.safeIndex(_106_v54, new BigNumber((_158_v91).length))], globalState);
          let _index13 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_148_v88).length));
          (_148_v88)[_index13] = (_dafny.ZERO).minus(_106_v54);
          let _159_v92;
          _159_v92 = _dafny.Map.Empty.slice().updateUnsafe(_146_v87,false);
          let _160_v93;
          _160_v93 = _dafny.Seq.of(_106_v54);
          let _161_v94;
          _161_v94 = _dafny.Map.Empty.slice().updateUnsafe((_105_v53).f22,_160_v93);
          let _162_v95;
          _162_v95 = _module.D7.create_DC16(_158_v91);
          let _163_v96;
          _163_v96 = new _dafny.CodePoint('l'.codePointAt(0));
          let _164_v97;
          _164_v97 = _dafny.Map.Empty.slice().updateUnsafe(_163_v96,_105_v53.f21);
          let _165_v98;
          _165_v98 = _dafny.MultiSet.fromElements(new BigNumber((_164_v97).length), _106_v54);
          let _166_v99;
          _166_v99 = _165_v98;
          let _167_v100;
          _167_v100 = _dafny.Set.fromElements(_166_v99);
          let _168_v101;
          _168_v101 = _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber((_161_v94).length), _106_v54), new BigNumber(-234), new BigNumber(((_162_v95).dtor_cf28).length), _157_cf40, new BigNumber((_167_v100).length));
          let _169_v102;
          _169_v102 = _dafny.Seq.UnicodeFromString("e");
          let _170_v103;
          _170_v103 = _dafny.Seq.of(_dafny.Seq.of(_105_v53.f21));
          let _index14 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_148_v88).length));
          let _rhs18 = _159_v92;
          let _rhs19 = _160_v93;
          let _rhs20 = _dafny.Seq.IsProperPrefixOf(_169_v102, _169_v102);
          let _rhs21 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_158_v91, _module.__default.safeIndex(_106_v54, new BigNumber((_158_v91).length)), false), _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_105_v53).f22, (_105_v53).f22, false, _module.__default.fm1((_105_v53).f22, _163_v96, globalState)), _module.__default.fm44(_170_v103, _163_v96, globalState)), _module.__default.safeIndex(_106_v54, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_105_v53).f22, (_105_v53).f22, false, _module.__default.fm1((_105_v53).f22, _163_v96, globalState)), _module.__default.fm44(_170_v103, _163_v96, globalState))).length)), _105_v53.f21))).length);
          let _lhs14 = _105_v53;
          let _lhs15 = _148_v88;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_148_v88).length));
          _159_v92 = _rhs18;
          _160_v93 = _rhs19;
          _lhs14.f21 = _rhs20;
          _lhs15[_lhs16] = _rhs21;
          let _171_v104;
          _171_v104 = _module.D0.create_DC0(_157_cf40);
          let _172_v105;
          let _nw25 = new _module.C7();
          _nw25.__ctor((_105_v53).f22, _104_v52, _171_v104, (_105_v53).f22);
          _172_v105 = _nw25;
          let _173_v106;
          _173_v106 = _dafny.Seq.of(_172_v105);
          let _174_v107;
          let _nw26 = Array((new BigNumber(27)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = ((_105_v53.f21) ? (_172_v105) : (_172_v105));
          _nw26[(_dafny.ONE).toNumber()] = _172_v105;
          _nw26[(new BigNumber(2)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(3)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(4)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(5)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(6)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(7)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(8)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(9)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(10)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(11)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(12)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(13)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(14)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(15)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(16)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(17)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(18)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(19)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(20)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(21)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(22)).toNumber()] = (_173_v106)[_module.__default.safeIndex(_106_v54, new BigNumber((_173_v106).length))];
          _nw26[(new BigNumber(23)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(24)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(25)).toNumber()] = _172_v105;
          _nw26[(new BigNumber(26)).toNumber()] = _172_v105;
          _174_v107 = _nw26;
          let _index15 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_174_v107).length));
          (_174_v107)[_index15] = _172_v105;
          let _index16 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_174_v107).length));
          (_174_v107)[_index16] = _172_v105;
          (globalState).f9 = new BigNumber(-780);
        } else {
          let _175___mcc_h2 = (_source0).cf39;
          let _176_cf39 = _175___mcc_h2;
          (_105_v53).f21 = _module.__default.fm1((_105_v53).f22, new _dafny.CodePoint('q'.codePointAt(0)), globalState);
          let _177_v108;
          _177_v108 = new _dafny.CodePoint('e'.codePointAt(0));
          let _178_v109;
          _178_v109 = _dafny.Seq.of(_177_v108, _177_v108, _module.__default.fm3(_106_v54, _105_v53.f21, globalState));
          let _179_v110;
          _179_v110 = _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_178_v109, _178_v109)).length)));
          let _180_v111;
          _180_v111 = _dafny.Seq.of(_178_v109, _178_v109);
          let _rhs22 = (_180_v111)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber(591), _106_v54), new BigNumber((_180_v111).length))];
          let _rhs23 = (_179_v110).Merge(_dafny.Map.Empty.slice().updateUnsafe((_105_v53).f22,_106_v54));
          let _rhs24 = _106_v54;
          _178_v109 = _rhs22;
          _179_v110 = _rhs23;
          _106_v54 = _rhs24;
          (globalState).f9 = _106_v54;
          let _181_v112;
          let _nw27 = new _module.C3();
          _nw27.__ctor(!(_45_v0), (_105_v53).f22, _104_v52);
          _181_v112 = _nw27;
          let _182_v113;
          _182_v113 = _dafny.Map.Empty.slice().updateUnsafe(_105_v53.f21,_181_v112);
          _182_v113 = (_182_v113).update(_45_v0, _181_v112);
        }
      } else {
        let _183_v114;
        let _nw28 = Array((new BigNumber(21)).toNumber());
        _183_v114 = _nw28;
        let _184_v115;
        let _nw29 = new _module.C4();
        _nw29.__ctor(_104_v52);
        _184_v115 = _nw29;
        let _index17 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_183_v114).length));
        (_183_v114)[_index17] = _184_v115;
        let _index18 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_183_v114).length));
        (_183_v114)[_index18] = _184_v115;
        _106_v54 = _106_v54;
        let _185_v116;
        let _nw30 = new _module.C2();
        _nw30.__ctor(_104_v52);
        _185_v116 = _nw30;
        let _186_v117;
        _186_v117 = _module.D8.create_DC19(_dafny.Seq.Create(_module.__default.abs(new BigNumber(647)), function (_187_i11) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}), _106_v54);
        let _188_v118;
        _188_v118 = _dafny.Seq.UnicodeFromString("jp");
        let _189_v119;
        _189_v119 = new _dafny.CodePoint('j'.codePointAt(0));
        _186_v117 = _module.D8.create_DC19(((_105_v53.f21) ? (_dafny.Seq.update(_188_v118, _module.__default.safeIndex(new BigNumber(762), new BigNumber((_188_v118).length)), _189_v119)) : (_dafny.Seq.UnicodeFromString("bpmojegsf"))), _module.__default.safeDivisionInt(_106_v54, _106_v54));
        let _190_v120;
        let _nw31 = new _module.C8();
        _nw31.__ctor(!(_105_v53.f21));
        _190_v120 = _nw31;
      }
      let _191_v121;
      _191_v121 = _module.D8.create_DC18(_145_v86);
      let _192_v122;
      _192_v122 = _dafny.Seq.of(_145_v86, _145_v86, _145_v86, _145_v86);
      let _pat_let_tv1 = _192_v122;
      let _pat_let_tv2 = _106_v54;
      let _pat_let_tv3 = _192_v122;
      _191_v121 = function (_pat_let2_0) {
        return function (_193_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_194_dt__update_hcf33_h0) {
              return _module.D8.create_DC18(_194_dt__update_hcf33_h0);
            }(_pat_let3_0);
          }((_pat_let_tv1)[_module.__default.safeIndex(_pat_let_tv2, new BigNumber((_pat_let_tv3).length))]);
        }(_pat_let2_0);
      }(_191_v121);
      let _195_v123;
      let _nw32 = Array((_dafny.ONE).toNumber());
      _nw32[(_dafny.ZERO).toNumber()] = _45_v0;
      _195_v123 = _nw32;
      let _index19 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_195_v123).length));
      (_195_v123)[_index19] = _45_v0;
      let _index20 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_195_v123).length));
      (_195_v123)[_index20] = !(_105_v53.f21);
      return;
    }
    static Main(__noArgsParameter) {
      let _196_v0;
      let _init5 = function (_197_i0) {
        return (_197_i0).multipliedBy(new BigNumber(-343));
      };
      let _nw33 = Array((new BigNumber(17)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
        _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _196_v0 = _nw33;
      let _198_v1;
      let _nw34 = Array((_dafny.ONE).toNumber());
      _nw34[(_dafny.ZERO).toNumber()] = _196_v0;
      _198_v1 = _nw34;
      let _199_globalState;
      let _nw35 = new _module.GlobalState();
      _nw35.__ctor(false, _196_v0, _dafny.Seq.UnicodeFromString("k"), _198_v1, new BigNumber(-63), new BigNumber(660), new BigNumber(763), new BigNumber(917), false, new BigNumber(720));
      _199_globalState = _nw35;
      let _200_v2;
      _200_v2 = new BigNumber(-797);
      let _hi1 = new BigNumber(901);
      for (let _201_i1 = _200_v2; _201_i1.isLessThan(_hi1); _201_i1 = _201_i1.plus(_dafny.ONE)) {
        let _202_v3;
        _202_v3 = false;
        let _203_v4;
        _203_v4 = _dafny.Set.fromElements(!(_202_v3));
        let _204_v5;
        _204_v5 = _dafny.Seq.of(_203_v4, _203_v4, _203_v4);
        let _205_v6;
        _205_v6 = _dafny.Map.Empty.slice().updateUnsafe(_202_v3,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_202_v3,new BigNumber(((_204_v5)[_module.__default.safeIndex(_module.__default.fm0(_202_v3, _199_globalState), new BigNumber((_204_v5).length))]).length))).length));
        let _206_v7;
        _206_v7 = new _dafny.CodePoint('s'.codePointAt(0));
        let _207_v8;
        _207_v8 = _dafny.MultiSet.fromElements(_206_v7, _206_v7, _206_v7);
        (_199_globalState).f9 = (((_205_v6).contains((_207_v8).IsProperSubsetOf(_207_v8))) ? ((_205_v6).get((_207_v8).IsProperSubsetOf(_207_v8))) : (_200_v2));
        let _208_v9;
        _208_v9 = _dafny.Set.fromElements(_201_i1);
        (_199_globalState).f8 = ((_208_v9).Difference(_208_v9)).IsProperSubsetOf(_208_v9);
        let _209_v10;
        _209_v10 = _dafny.MultiSet.fromElements(_202_v3);
        _209_v10 = ((_209_v10).Union(_209_v10)).update(false, _module.__default.abs((_201_i1).plus(_201_i1)));
        (_199_globalState).f9 = (_200_v2).multipliedBy((_dafny.ZERO).minus(_200_v2));
      }
      _module.__default.m0(_199_globalState);
      let _210_v11;
      _210_v11 = true;
      let _211_i2;
      _211_i2 = _dafny.ZERO;
      L1: {
        while (_210_v11) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_211_i2)) {
              break L1;
            }
            _211_i2 = (_211_i2).plus(_dafny.ONE);
            if (_210_v11) {
              _200_v2 = (_200_v2).plus(_200_v2);
              let _212_v12;
              _212_v12 = _dafny.Seq.UnicodeFromString("wwabxmyaf");
              let _213_v13;
              _213_v13 = _dafny.Map.Empty.slice().updateUnsafe(_212_v12,_200_v2);
              let _index21 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length));
              (_196_v0)[_index21] = new BigNumber((_213_v13).length);
              let _index22 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length));
              (_196_v0)[_index22] = (_200_v2).multipliedBy(_200_v2);
              let _214_v14;
              _214_v14 = _dafny.MultiSet.fromElements((_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))]);
              let _215_v15;
              _215_v15 = new _dafny.CodePoint('r'.codePointAt(0));
              let _216_v16;
              _216_v16 = _dafny.Seq.of(_210_v11, _210_v11, _module.__default.fm1(_210_v11, _215_v15, _199_globalState), _210_v11);
              (_199_globalState).f9 = (_dafny.ZERO).minus(((((_214_v14).contains(_200_v2)) ? ((_214_v14).get(_200_v2)) : ((_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))]))).plus((new BigNumber((_216_v16).length)).minus(_200_v2)));
              let _217_v17;
              let _nw36 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _217_v17 = _nw36;
              let _index23 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_217_v17).length));
              (_217_v17)[_index23] = _212_v12;
              let _218_v18;
              let _nw37 = Array((new BigNumber(4)).toNumber()).fill([]);
              _218_v18 = _nw37;
              let _219_v19;
              let _nw38 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Set.Empty);
              _219_v19 = _nw38;
              let _index24 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_218_v18).length));
              (_218_v18)[_index24] = _219_v19;
              let _index25 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_217_v17).length));
              let _index26 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_218_v18).length));
              let _rhs25 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm2((_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))], _dafny.MultiSet.fromElements((_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))]), _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_200_v2), _199_globalState), _dafny.Seq.update(_212_v12, _module.__default.safeIndex((_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))], new BigNumber((_212_v12).length)), _215_v15))).length);
              let _rhs26 = _212_v12;
              let _rhs27 = _210_v11;
              let _rhs28 = _219_v19;
              let _rhs29 = _200_v2;
              let _lhs17 = _199_globalState;
              let _lhs18 = _217_v17;
              let _lhs19 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_217_v17).length));
              let _lhs20 = _199_globalState;
              let _lhs21 = _218_v18;
              let _lhs22 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_218_v18).length));
              let _lhs23 = _199_globalState;
              _lhs17.f9 = _rhs25;
              _lhs18[_lhs19] = _rhs26;
              _lhs20.f8 = _rhs27;
              _lhs21[_lhs22] = _rhs28;
              _lhs23.f9 = _rhs29;
              _210_v11 = (((_210_v11) ? (_200_v2) : (_200_v2))).isEqualTo(_module.__default.safeDivisionInt((_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))], (_196_v0)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_196_v0).length))]));
            } else {
              (_199_globalState).f8 = _210_v11;
              (_199_globalState).f8 = _210_v11;
              (_199_globalState).f9 = _module.__default.fm0(true, _199_globalState);
              let _220_v20;
              let _nw39 = Array((new BigNumber(26)).toNumber()).fill(false);
              _220_v20 = _nw39;
              let _221_v21;
              _221_v21 = new _dafny.CodePoint('f'.codePointAt(0));
              let _222_v22;
              _222_v22 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jgl"));
              let _223_v23;
              _223_v23 = _dafny.MultiSet.fromElements(_200_v2);
              let _224_v24;
              _224_v24 = _dafny.Seq.UnicodeFromString("bwtvm");
              let _225_v25;
              _225_v25 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_220_v20);
              let _rhs30 = new BigNumber((_dafny.Seq.Concat(_222_v22, _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm2(_200_v2, _223_v23, _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_200_v2), _199_globalState)), _dafny.Seq.of(_224_v24)))).length);
              let _rhs31 = (((_225_v25).contains((_dafny.ZERO).minus(_200_v2))) ? ((_225_v25).get((_dafny.ZERO).minus(_200_v2))) : (_220_v20));
              let _rhs32 = _221_v21;
              let _lhs24 = _199_globalState;
              _lhs24.f9 = _rhs30;
              _220_v20 = _rhs31;
              _221_v21 = _rhs32;
              let _226_v26;
              _226_v26 = _dafny.MultiSet.fromElements(_210_v11, _210_v11);
              let _227_v27;
              _227_v27 = _module.D0.create_DC2(_226_v26, true, _200_v2, _210_v11);
              let _rhs33 = (_200_v2).minus(_200_v2);
              let _rhs34 = !((_227_v27).dtor_cf6).isEqualTo(_200_v2);
              let _rhs35 = ((_dafny.areEqual(_224_v24, _224_v24)) ? ((_dafny.ZERO).minus(_module.__default.fm0(_210_v11, _199_globalState))) : (_200_v2));
              let _lhs25 = _199_globalState;
              let _lhs26 = _199_globalState;
              let _lhs27 = _199_globalState;
              _lhs25.f9 = _rhs33;
              _lhs26.f8 = _rhs34;
              _lhs27.f9 = _rhs35;
            }
            (_199_globalState).f8 = true;
            let _228_v28;
            _228_v28 = _dafny.Map.Empty.slice().updateUnsafe(_210_v11,!(_210_v11));
            let _229_v29;
            _229_v29 = _dafny.Map.Empty.slice().updateUnsafe((((_228_v28).contains(_210_v11)) ? ((_228_v28).get(_210_v11)) : (_210_v11)),_210_v11);
            let _rhs36 = _210_v11;
            let _rhs37 = (((_210_v11) ? (_229_v29) : (_228_v28))).Merge((_228_v28).Merge(_229_v29));
            let _rhs38 = (_200_v2).isEqualTo(_200_v2);
            let _lhs28 = _199_globalState;
            _lhs28.f8 = _rhs36;
            _229_v29 = _rhs37;
            _210_v11 = _rhs38;
            _module.__default.m0(_199_globalState);
          }
        }
      }
      _module.__default.m0(_199_globalState);
      if (_210_v11) {
        if (_210_v11) {
          let _230_v30;
          let _nw40 = Array((new BigNumber(14)).toNumber()).fill(false);
          _230_v30 = _nw40;
          let _index27 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_230_v30).length));
          (_230_v30)[_index27] = _210_v11;
          let _231_v31;
          _231_v31 = _dafny.Map.Empty.slice().updateUnsafe(_210_v11,_200_v2);
          let _232_v32;
          _232_v32 = _dafny.Map.Empty.slice().updateUnsafe((((_231_v31).contains(_210_v11)) ? ((_231_v31).get(_210_v11)) : (_200_v2)),!(_210_v11));
          let _233_v33;
          _233_v33 = _dafny.Map.Empty.slice().updateUnsafe(_210_v11,_dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), function (_234_i3) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          }));
          let _index28 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_230_v30).length));
          (_230_v30)[_index28] = (((_232_v32).contains((new BigNumber((_233_v33).length)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_210_v11)).cardinality())))) ? ((_232_v32).get((new BigNumber((_233_v33).length)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_210_v11)).cardinality())))) : (!(_210_v11)));
          let _235_v34;
          _235_v34 = new _dafny.CodePoint('q'.codePointAt(0));
          (_199_globalState).f8 = _module.__default.fm1((_230_v30)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_230_v30).length))], _235_v34, _199_globalState);
          _231_v31 = _231_v31;
          _235_v34 = _module.__default.fm3(new BigNumber((_dafny.Seq.of(_210_v11)).length), (_230_v30)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_230_v30).length))], _199_globalState);
          _230_v30 = _230_v30;
        } else {
          let _236_v35;
          _236_v35 = _dafny.Seq.UnicodeFromString("nliacvc");
          _236_v35 = _236_v35;
          let _237_v36;
          let _nw41 = Array((new BigNumber(17)).toNumber()).fill(false);
          _237_v36 = _nw41;
          let _index29 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length));
          (_237_v36)[_index29] = _210_v11;
          let _index30 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length));
          (_237_v36)[_index30] = ((_200_v2).plus(_200_v2)).isLessThanOrEqualTo(_module.__default.fm0(_210_v11, _199_globalState));
          let _238_v37;
          _238_v37 = _dafny.MultiSet.fromElements(_200_v2, _200_v2, (_dafny.ZERO).minus(_200_v2));
          _238_v37 = _238_v37;
          let _239_v38;
          _239_v38 = _dafny.MultiSet.fromElements(_210_v11, _210_v11);
          let _240_v39;
          _240_v39 = new _dafny.CodePoint('e'.codePointAt(0));
          let _241_v40;
          _241_v40 = _module.D0.create_DC2((_239_v38).Union(_239_v38), (_200_v2).isLessThan(_200_v2), _200_v2, _module.__default.fm1((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))], _240_v39, _199_globalState));
          let _242_v41;
          _242_v41 = _dafny.Seq.of((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))], (_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))]);
          let _243_v42;
          let _nw42 = Array((new BigNumber(21)).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _242_v41;
          _nw42[(_dafny.ONE).toNumber()] = _242_v41;
          _nw42[(new BigNumber(2)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(3)).toNumber()] = _dafny.Seq.of((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))]);
          _nw42[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_242_v41, _242_v41);
          _nw42[(new BigNumber(5)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_242_v41, _dafny.Seq.update(_242_v41, _module.__default.safeIndex(_200_v2, new BigNumber((_242_v41).length)), (_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))]));
          _nw42[(new BigNumber(7)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(8)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(9)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(10)).toNumber()] = _dafny.Seq.of((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))], true);
          _nw42[(new BigNumber(11)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(12)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_242_v41, _242_v41);
          _nw42[(new BigNumber(14)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(((true) ? (_dafny.Seq.update(_242_v41, _module.__default.safeIndex(new BigNumber((_242_v41).length), new BigNumber((_242_v41).length)), _210_v11)) : (_242_v41)), _module.__default.safeIndex(_200_v2, new BigNumber((((true) ? (_dafny.Seq.update(_242_v41, _module.__default.safeIndex(new BigNumber((_242_v41).length), new BigNumber((_242_v41).length)), _210_v11)) : (_242_v41))).length)), false);
          _nw42[(new BigNumber(16)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(17)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(18)).toNumber()] = _242_v41;
          _nw42[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_242_v41, _242_v41);
          _nw42[(new BigNumber(20)).toNumber()] = _242_v41;
          _243_v42 = _nw42;
          let _index31 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_243_v42).length));
          (_243_v42)[_index31] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))], (_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))]), _module.__default.safeIndex(_200_v2, new BigNumber((_dafny.Seq.of((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))], (_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))])).length)), _210_v11), _242_v41);
          let _244_v43;
          _244_v43 = _dafny.Set.fromElements((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))], !((_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))]), _210_v11, true, (_237_v36)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length))]);
          let _245_v44;
          _245_v44 = _dafny.Seq.of((_dafny.ZERO).minus(_200_v2));
          let _246_v45;
          _246_v45 = _dafny.Map.Empty.slice().updateUnsafe(_236_v35,_200_v2);
          let _index32 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length));
          let _index33 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_243_v42).length));
          let _rhs39 = _module.__default.fm4(new BigNumber((_244_v43).length), _200_v2, _199_globalState);
          let _rhs40 = _237_v36;
          let _rhs41 = !((new BigNumber((_245_v44).length)).isLessThanOrEqualTo((_200_v2).plus(new BigNumber((_246_v45).length))));
          let _rhs42 = _242_v41;
          let _rhs43 = _module.__default.fm3(((_dafny.ZERO).minus(_200_v2)).multipliedBy(_200_v2), false, _199_globalState);
          let _lhs29 = _237_v36;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_237_v36).length));
          let _lhs31 = _243_v42;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_243_v42).length));
          _241_v40 = _rhs39;
          _237_v36 = _rhs40;
          _lhs29[_lhs30] = _rhs41;
          _lhs31[_lhs32] = _rhs42;
          _240_v39 = _rhs43;
          _module.__default.m0(_199_globalState);
        }
        let _247_v46;
        _247_v46 = _dafny.Seq.UnicodeFromString("cyv");
        let _248_v47;
        _248_v47 = new _dafny.CodePoint('k'.codePointAt(0));
        _247_v46 = _dafny.Seq.update(_dafny.Seq.Concat(_247_v46, _247_v46), _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_200_v2, _200_v2)), new BigNumber((_dafny.Seq.Concat(_247_v46, _247_v46)).length)), _248_v47);
        let _249_v48;
        _249_v48 = _module.D5.create_DC14(_210_v11);
        let _250_v49;
        let _init6 = ((_251_v2, _252_v11) => function (_253_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe(_251_v2,_252_v11);
        })(_200_v2, _210_v11);
        let _nw43 = Array((new BigNumber(2)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw43.length); _i0_6++) {
          _nw43[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _250_v49 = _nw43;
        let _254_v50;
        let _nw44 = new _module.C3();
        _nw44.__ctor((_249_v48).dtor_cf26, _210_v11, ((_210_v11) ? (_250_v49) : (_250_v49)));
        _254_v50 = _nw44;
        let _255_v51;
        _255_v51 = _module.D0.create_DC0(_200_v2);
        let _256_v52;
        let _nw45 = new _module.C1();
        _nw45.__ctor(_254_v50.f21, _254_v50.f21, _255_v51, _210_v11);
        _256_v52 = _nw45;
        let _257_v53;
        _257_v53 = _dafny.Seq.of(_256_v52, _256_v52);
        let _258_v54;
        _258_v54 = _dafny.MultiSet.fromElements(_257_v53);
        let _259_v55;
        _259_v55 = _dafny.Map.Empty.slice().updateUnsafe(_196_v0,_257_v53);
        let _260_v56;
        let _261_v57;
        let _out0;
        let _out1;
        let _outcollector0 = (_254_v50).m13(_200_v2, (_254_v50).f22, new BigNumber(641), (_dafny.MultiSet.FromArray(_dafny.Seq.of((((_259_v55).contains(_196_v0)) ? ((_259_v55).get(_196_v0)) : (_257_v53))))).IsSubsetOf(_258_v54), _199_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _260_v56 = _out0;
        _261_v57 = _out1;
        if ((_254_v50).f22) {
          let _index34 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length));
          (_196_v0)[_index34] = _200_v2;
          let _index35 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length));
          (_196_v0)[_index35] = _200_v2;
          let _262_v58;
          let _nw46 = new _module.C2();
          _nw46.__ctor(_250_v49);
          _262_v58 = _nw46;
          _262_v58 = _262_v58;
          let _263_v59;
          let _nw47 = Array((new BigNumber(5)).toNumber()).fill(false);
          _263_v59 = _nw47;
          let _264_v60;
          _264_v60 = _dafny.Seq.of(new BigNumber((_247_v46).length));
          let _265_v61;
          _265_v61 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_254_v50.f21);
          let _266_v62;
          _266_v62 = _dafny.Seq.of(((_196_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length))]).plus(_200_v2), _200_v2, new BigNumber((_264_v60).length), (_256_v52).fm30(_265_v61, _199_globalState), _200_v2);
          let _267_v63;
          _267_v63 = _dafny.Seq.of((_254_v50).f22);
          let _rhs44 = _200_v2;
          let _rhs45 = (_264_v60)[_module.__default.safeIndex(new BigNumber((_267_v63).length), new BigNumber((_264_v60).length))];
          let _rhs46 = _module.__default.safeDivisionInt(_200_v2, (_dafny.ZERO).minus((_200_v2).plus(_200_v2)));
          let _rhs47 = _263_v59;
          let _rhs48 = ((((_254_v50).f22) ? ((_196_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length))]) : ((_196_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length))]))).plus(_200_v2);
          let _lhs33 = _199_globalState;
          let _lhs34 = _199_globalState;
          let _lhs35 = _199_globalState;
          _lhs33.f9 = _rhs44;
          _lhs34.f9 = _rhs45;
          _200_v2 = _rhs46;
          _263_v59 = _rhs47;
          _lhs35.f9 = _rhs48;
          let _268_v64;
          _268_v64 = _module.D2.create_DC6();
          let _269_v65;
          _269_v65 = _dafny.Map.Empty.slice().updateUnsafe(_268_v64,true);
          let _270_v66;
          _270_v66 = _dafny.Map.Empty.slice().updateUnsafe(true,(_256_v52).f20);
          let _271_v67;
          _271_v67 = _dafny.Map.Empty.slice().updateUnsafe(_270_v66,_200_v2);
          _269_v65 = (_269_v65).update(_268_v64, ((((_271_v67).contains(_270_v66)) ? ((_271_v67).get(_270_v66)) : (new BigNumber((_267_v63).length)))).isLessThanOrEqualTo(new BigNumber(977)));
          let _272_v68;
          let _init7 = ((_273_v2, _274_v0) => function (_275_i5) {
            return (_dafny.MultiSet.fromElements(_273_v2)).Intersect(_dafny.MultiSet.fromElements((_274_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_274_v0).length))], (_274_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_274_v0).length))], (_274_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_274_v0).length))]));
          })(_200_v2, _196_v0);
          let _nw48 = Array((new BigNumber(27)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw48.length); _i0_7++) {
            _nw48[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _272_v68 = _nw48;
          let _276_v69;
          _276_v69 = _dafny.MultiSet.fromElements((_196_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length))]);
          let _index36 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_272_v68).length));
          (_272_v68)[_index36] = (_276_v69).Difference(_276_v69);
          let _277_v70;
          let _nw49 = new _module.C6();
          _nw49.__ctor(_200_v2, _247_v46);
          _277_v70 = _nw49;
          let _278_v71;
          _278_v71 = _dafny.Map.Empty.slice().updateUnsafe(_277_v70,(_256_v52).f19);
          let _279_v72;
          _279_v72 = _dafny.Map.Empty.slice().updateUnsafe((_277_v70).f15,new BigNumber((_277_v70.f16).length));
          let _index37 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_272_v68).length));
          (_272_v68)[_index37] = (_276_v69).Union((_276_v69).Intersect(((_276_v69).update(new BigNumber((_278_v71).length), _module.__default.abs(new BigNumber(571)))).update(new BigNumber((_279_v72).length), _module.__default.abs((_196_v0)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_196_v0).length))]))));
        } else {
          (_199_globalState).f9 = _200_v2;
          let _280_v73;
          let _nw50 = new _module.C5();
          _nw50.__ctor(_250_v49);
          _280_v73 = _nw50;
          let _281_v74;
          _281_v74 = _dafny.Set.fromElements(_280_v73, _280_v73);
          let _rhs49 = ((_281_v74).Intersect(_281_v74)).IsProperSubsetOf(((_210_v11) ? (_dafny.Set.fromElements(_280_v73, _280_v73, _280_v73, _280_v73)) : (_281_v74)));
          let _rhs50 = (_254_v50).f22;
          _260_v56 = _rhs49;
          _210_v11 = _rhs50;
          (_199_globalState).f9 = (_256_v52).fm30(_dafny.Map.Empty.slice().updateUnsafe(_200_v2,(_256_v52).f20), _199_globalState);
          let _282_v75;
          _282_v75 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_dafny.Map.Empty.slice().updateUnsafe((_256_v52).f19,_200_v2));
          let _283_v76;
          _283_v76 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_254_v50.f21);
          let _284_v77;
          _284_v77 = _dafny.Map.Empty.slice().updateUnsafe((_256_v52).f19,(_256_v52).fm30(_283_v76, _199_globalState));
          let _285_v78;
          let _out2;
          _out2 = (_254_v50).m12(_260_v56, _260_v56, (((_282_v75).contains(_200_v2)) ? ((_282_v75).get(_200_v2)) : (_284_v77)), _199_globalState);
          _285_v78 = _out2;
          let _286_v79;
          let _nw51 = new _module.C0();
          _nw51.__ctor(_285_v78);
          _286_v79 = _nw51;
          let _287_v80;
          let _nw52 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _287_v80 = _nw52;
          let _288_v81;
          let _nw53 = new _module.C5();
          _nw53.__ctor(_250_v49);
          _288_v81 = _nw53;
          let _289_v82;
          _289_v82 = _dafny.Map.Empty.slice().updateUnsafe(_288_v81,(_254_v50).f22);
          let _index38 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_287_v80).length));
          (_287_v80)[_index38] = _289_v82;
          let _290_v83;
          _290_v83 = _module.D4.create_DC11(_200_v2, _200_v2, (_254_v50).f22);
          let _pat_let_tv4 = _200_v2;
          let _index39 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_287_v80).length));
          let _rhs51 = _286_v79;
          let _rhs52 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_200_v2, _200_v2), _200_v2);
          let _rhs53 = _dafny.Map.Empty.slice().updateUnsafe(_288_v81,!(_260_v56));
          let _rhs54 = (_200_v2).multipliedBy((((_254_v50).f22) ? (_200_v2) : (_200_v2)));
          let _rhs55 = (function (_pat_let4_0) {
            return function (_291_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_292_dt__update_hcf19_h0) {
                  return _module.D4.create_DC11(_292_dt__update_hcf19_h0, (_291_dt__update__tmp_h0).dtor_cf20, (_291_dt__update__tmp_h0).dtor_cf21);
                }(_pat_let5_0);
              }(_pat_let_tv4);
            }(_pat_let4_0);
          }(_290_v83)).dtor_cf21;
          let _lhs36 = _287_v80;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_287_v80).length));
          let _lhs38 = _199_globalState;
          let _lhs39 = _254_v50;
          _286_v79 = _rhs51;
          _200_v2 = _rhs52;
          _lhs36[_lhs37] = _rhs53;
          _lhs38.f9 = _rhs54;
          _lhs39.f21 = _rhs55;
        }
      } else {
        let _293_v84;
        _293_v84 = _dafny.Seq.of(_210_v11, false);
        let _294_v85;
        _294_v85 = new _dafny.CodePoint('y'.codePointAt(0));
        if (_dafny.areEqual(_293_v84, _module.__default.fm44(_dafny.Seq.of(_dafny.Seq.of(_210_v11)), _294_v85, _199_globalState))) {
          (_199_globalState).f8 = _module.__default.fm1(_210_v11, _294_v85, _199_globalState);
          let _295_v86;
          _295_v86 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
          let _296_v87;
          _296_v87 = _dafny.MultiSet.fromElements(_210_v11);
          let _297_v88;
          _297_v88 = _dafny.Seq.of(_200_v2, new BigNumber((_295_v86).length), new BigNumber((_296_v87).cardinality()));
          let _298_v89;
          _298_v89 = _dafny.Seq.of(_297_v88);
          _200_v2 = ((_210_v11) ? (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_299_v2) => function (_300_i6) {
            return _dafny.Seq.of(_299_v2, _299_v2, _299_v2, new BigNumber((_dafny.Seq.UnicodeFromString("bojpia")).length));
          })(_200_v2)), _module.__default.safeIndex(_200_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_301_v2) => function (_302_i6) {
            return _dafny.Seq.of(_301_v2, _301_v2, _301_v2, new BigNumber((_dafny.Seq.UnicodeFromString("bojpia")).length));
          })(_200_v2))).length)), _297_v88), _298_v89)).length)) : (_200_v2));
          let _303_v90;
          _303_v90 = _dafny.Seq.UnicodeFromString("nq");
          (_199_globalState).f8 = _module.__default.fm1(_210_v11, (_303_v90)[_module.__default.safeIndex((_dafny.ZERO).minus(_200_v2), new BigNumber((_303_v90).length))], _199_globalState);
          let _304_v91;
          _304_v91 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_295_v86);
          _304_v91 = (_304_v91).update((new BigNumber((_293_v84).length)).plus(_200_v2), function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(943), new BigNumber(750))) {
              let _305_v92 = _compr_25;
              if (((new BigNumber(943)).isLessThanOrEqualTo(_305_v92)) && ((_305_v92).isLessThan(new BigNumber(750)))) {
                _coll25.push([(_305_v92).multipliedBy(_200_v2),(_293_v84)[_module.__default.safeIndex(new BigNumber(-98), new BigNumber((_293_v84).length))]]);
              }
            }
            return _coll25;
          }());
          let _306_v93;
          _306_v93 = _dafny.Set.fromElements(_303_v90, _303_v90, _303_v90, _303_v90, _303_v90);
          let _307_v94;
          let _nw54 = Array((new BigNumber(12)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _306_v93;
          _nw54[(_dafny.ONE).toNumber()] = _306_v93;
          _nw54[(new BigNumber(2)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(3)).toNumber()] = _module.__default.fm63(_199_globalState);
          _nw54[(new BigNumber(4)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(5)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(6)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(7)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(8)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(9)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(10)).toNumber()] = _306_v93;
          _nw54[(new BigNumber(11)).toNumber()] = _306_v93;
          _307_v94 = _nw54;
          let _index40 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_307_v94).length));
          (_307_v94)[_index40] = _306_v93;
          let _index41 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_307_v94).length));
          (_307_v94)[_index41] = _module.__default.fm63(_199_globalState);
        } else {
          let _308_v95;
          let _nw55 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _308_v95 = _nw55;
          let _309_v96;
          let _init8 = ((_310_v2, _311_v11) => function (_312_i7) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rvf")).length), _310_v2)).length),_311_v11);
          })(_200_v2, _210_v11);
          let _nw56 = Array((new BigNumber(6)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw56.length); _i0_8++) {
            _nw56[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _309_v96 = _nw56;
          let _313_v97;
          let _nw57 = new _module.C2();
          _nw57.__ctor(_309_v96);
          _313_v97 = _nw57;
          let _314_v98;
          _314_v98 = _dafny.Map.Empty.slice().updateUnsafe(_210_v11,_313_v97);
          let _index42 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_308_v95).length));
          (_308_v95)[_index42] = (_314_v98).Merge(_314_v98);
          let _index43 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_308_v95).length));
          (_308_v95)[_index43] = ((_dafny.Map.Empty.slice().updateUnsafe(_210_v11,_313_v97)).Merge(_314_v98)).Merge(_314_v98);
          let _315_v99;
          _315_v99 = _dafny.Set.fromElements(_210_v11);
          let _index44 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_196_v0).length));
          (_196_v0)[_index44] = new BigNumber((_315_v99).length);
          let _index45 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_196_v0).length));
          (_196_v0)[_index45] = _module.__default.safeDivisionInt(_200_v2, _200_v2);
          let _316_v100;
          _316_v100 = _module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber(-905)));
          let _317_v101;
          let _nw58 = new _module.C7();
          _nw58.__ctor(_210_v11, (_313_v97).f11, _316_v100, _210_v11);
          _317_v101 = _nw58;
          _317_v101 = _317_v101;
          (_199_globalState).f8 = true;
          let _318_v102;
          let _nw59 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _318_v102 = _nw59;
          let _319_v103;
          let _320_v104;
          let _321_v105;
          let _322_v106;
          let _out3;
          let _out4;
          let _out5;
          let _out6;
          let _outcollector1 = (_317_v101).m4(_200_v2, _318_v102, _199_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _out6 = _outcollector1[3];
          _319_v103 = _out3;
          _320_v104 = _out4;
          _321_v105 = _out5;
          _322_v106 = _out6;
        }
        _210_v11 = _210_v11;
        let _323_v107;
        _323_v107 = _dafny.Set.fromElements(_210_v11, !(_210_v11));
        if (((_323_v107).Union(_dafny.Set.fromElements(_210_v11))).IsProperSubsetOf(_323_v107)) {
          let _324_v108;
          let _nw60 = Array((new BigNumber(17)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = (_293_v84)[_module.__default.safeIndex(_200_v2, new BigNumber((_293_v84).length))];
          _nw60[(_dafny.ONE).toNumber()] = _210_v11;
          _nw60[(new BigNumber(2)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(3)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(4)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(5)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(6)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(7)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(8)).toNumber()] = true;
          _nw60[(new BigNumber(9)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(10)).toNumber()] = !(true);
          _nw60[(new BigNumber(11)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(12)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(13)).toNumber()] = !(false);
          _nw60[(new BigNumber(14)).toNumber()] = _210_v11;
          _nw60[(new BigNumber(15)).toNumber()] = true;
          _nw60[(new BigNumber(16)).toNumber()] = _210_v11;
          _324_v108 = _nw60;
          let _325_v109;
          _325_v109 = _module.D10.create_DC24(_324_v108, _196_v0, _dafny.Seq.UnicodeFromString("rsxa"));
          _196_v0 = (_325_v109).dtor_cf44;
          (_199_globalState).f8 = _210_v11;
          let _326_v110;
          let _nw61 = Array((new BigNumber(19)).toNumber());
          _326_v110 = _nw61;
          let _327_v112;
          let _init9 = ((_328_v2) => function (_329_i8) {
            return function () {
              let _coll26 = new _dafny.Map();
              for (const _compr_26 of (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_328_v2),new BigNumber(229))).Keys.Elements) {
                let _330_v111 = _compr_26;
                if ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_328_v2),new BigNumber(229))).contains(_330_v111)) {
                  _coll26.push([_module.__default.safeModuloInt(_330_v111, _328_v2),true]);
                }
              }
              return _coll26;
            }();
          })(_200_v2);
          let _nw62 = Array((new BigNumber(22)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw62.length); _i0_9++) {
            _nw62[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _327_v112 = _nw62;
          let _331_v113;
          let _nw63 = new _module.C3();
          _nw63.__ctor(!(_210_v11), _210_v11, _327_v112);
          _331_v113 = _nw63;
          let _index46 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_326_v110).length));
          (_326_v110)[_index46] = _331_v113;
          let _332_v114;
          _332_v114 = _331_v113;
          let _index47 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_326_v110).length));
          (_326_v110)[_index47] = (_332_v114);
          let _333_v115;
          _333_v115 = _dafny.Seq.of(_200_v2);
          let _334_v116;
          _334_v116 = _module.D7.create_DC17(_dafny.Seq.UnicodeFromString("rvj"), _324_v108, (_331_v113).f22, new BigNumber((_333_v115).length));
          let _335_v117;
          let _nw64 = Array((new BigNumber(17)).toNumber());
          _nw64[(_dafny.ZERO).toNumber()] = _324_v108;
          _nw64[(_dafny.ONE).toNumber()] = _324_v108;
          _nw64[(new BigNumber(2)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(3)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(4)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(5)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(6)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(7)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(8)).toNumber()] = (_334_v116).dtor_cf30;
          _nw64[(new BigNumber(9)).toNumber()] = (((_331_v113).f22) ? (_324_v108) : (_324_v108));
          _nw64[(new BigNumber(10)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(11)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(12)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(13)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(14)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(15)).toNumber()] = _324_v108;
          _nw64[(new BigNumber(16)).toNumber()] = _324_v108;
          _335_v117 = _nw64;
          let _index48 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_335_v117).length));
          (_335_v117)[_index48] = _324_v108;
          let _index49 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_335_v117).length));
          (_335_v117)[_index49] = _324_v108;
          (_199_globalState).f8 = _331_v113.f21;
        } else {
          let _336_v118;
          _336_v118 = _dafny.Seq.of(_200_v2);
          _336_v118 = _336_v118;
          let _337_v119;
          _337_v119 = _dafny.MultiSet.fromElements(_210_v11, _210_v11, !(!(_210_v11)));
          let _338_v120;
          _338_v120 = _dafny.Map.Empty.slice().updateUnsafe(_337_v119,(_210_v11) && (_210_v11));
          _338_v120 = (_338_v120).update(_337_v119, _210_v11);
          let _339_v121;
          _339_v121 = _dafny.MultiSet.fromElements(_200_v2, _200_v2, _module.__default.fm0(_210_v11, _199_globalState), new BigNumber(79), _200_v2);
          let _340_v122;
          let _init10 = ((_341_v2, _342_v11) => function (_343_i9) {
            return _dafny.Map.Empty.slice().updateUnsafe(_341_v2,_342_v11);
          })(_200_v2, _210_v11);
          let _nw65 = Array((new BigNumber(23)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw65.length); _i0_10++) {
            _nw65[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _340_v122 = _nw65;
          let _344_v123;
          let _nw66 = new _module.C5();
          _nw66.__ctor(_340_v122);
          _344_v123 = _nw66;
          let _345_v124;
          _345_v124 = _dafny.Map.Empty.slice().updateUnsafe((_339_v121).Union(_339_v121),_344_v123);
          _345_v124 = (_345_v124).update(_339_v121, _344_v123);
          _210_v11 = _210_v11;
          let _346_v125;
          _346_v125 = _dafny.Seq.UnicodeFromString("khkrqxbg");
          let _347_v126;
          _347_v126 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(579),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_210_v11,_200_v2)).length));
          let _348_v127;
          let _nw67 = Array((new BigNumber(13)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,new BigNumber((_337_v119).cardinality()));
          _nw67[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(702), _200_v2)).length),new BigNumber((_346_v125).length));
          _nw67[(new BigNumber(2)).toNumber()] = _module.__default.fm7(_199_globalState);
          _nw67[(new BigNumber(3)).toNumber()] = (_347_v126).update(_200_v2, (_dafny.ZERO).minus(_200_v2));
          _nw67[(new BigNumber(4)).toNumber()] = _module.__default.fm7(_199_globalState);
          _nw67[(new BigNumber(5)).toNumber()] = _347_v126;
          _nw67[(new BigNumber(6)).toNumber()] = _347_v126;
          _nw67[(new BigNumber(7)).toNumber()] = _347_v126;
          _nw67[(new BigNumber(8)).toNumber()] = _347_v126;
          _nw67[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,new BigNumber(204));
          _nw67[(new BigNumber(10)).toNumber()] = _347_v126;
          _nw67[(new BigNumber(11)).toNumber()] = _347_v126;
          _nw67[(new BigNumber(12)).toNumber()] = _347_v126;
          _348_v127 = _nw67;
          let _349_v128;
          _349_v128 = _dafny.Seq.of(_347_v126);
          let _index50 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_348_v127).length));
          (_348_v127)[_index50] = (_347_v126).Merge((_349_v128)[_module.__default.safeIndex((_dafny.ZERO).minus(_200_v2), new BigNumber((_349_v128).length))]);
          let _index51 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_348_v127).length));
          (_348_v127)[_index51] = (_347_v126).Merge(_347_v126);
        }
        let _350_v129;
        _350_v129 = _dafny.Map.Empty.slice().updateUnsafe(_210_v11,(_dafny.ZERO).minus(_200_v2));
        let _351_v130;
        _351_v130 = _dafny.Seq.UnicodeFromString("cbvy");
        let _352_v131;
        _352_v131 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_351_v130).length),!(_210_v11));
        let _353_v132;
        _353_v132 = _dafny.Seq.of((((_350_v129).contains(_210_v11)) ? ((_350_v129).get(_210_v11)) : (new BigNumber((_352_v131).length))));
        _353_v132 = _353_v132;
        _200_v2 = _200_v2;
      }
      let _354_v133;
      _354_v133 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
      let _355_v135;
      _355_v135 = _dafny.MultiSet.fromElements(_200_v2, _200_v2);
      let _356_v136;
      let _nw68 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
      _356_v136 = _nw68;
      let _357_v137;
      let _nw69 = new _module.C2();
      _nw69.__ctor(_356_v136);
      _357_v137 = _nw69;
      let _358_v138;
      _358_v138 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_357_v137);
      let _359_v141;
      _359_v141 = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_200_v2);
      let _360_v142;
      let _nw70 = Array((new BigNumber(23)).toNumber());
      _nw70[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
      _nw70[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
      _nw70[(new BigNumber(2)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(3)).toNumber()] = function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of (_355_v135).Elements) {
          let _361_v134 = _compr_27;
          if ((_355_v135).contains(_361_v134)) {
            _coll27.push([(_361_v134).minus(_200_v2),true]);
          }
        }
        return _coll27;
      }();
      _nw70[(new BigNumber(4)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(5)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(6)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(7)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
      _nw70[(new BigNumber(9)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(10)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(11)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(12)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_358_v138).length),_210_v11);
      _nw70[(new BigNumber(14)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(15)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(16)).toNumber()] = function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of _dafny.IntegerRange(new BigNumber(793), new BigNumber(340))) {
          let _362_v139 = _compr_28;
          if (((new BigNumber(793)).isLessThanOrEqualTo(_362_v139)) && ((_362_v139).isLessThan(new BigNumber(340)))) {
            _coll28.push([_module.__default.safeModuloInt(_362_v139, new BigNumber((_354_v133).length)),_210_v11]);
          }
        }
        return _coll28;
      }();
      _nw70[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
      _nw70[(new BigNumber(18)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(19)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(20)).toNumber()] = _354_v133;
      _nw70[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_200_v2,_210_v11);
      _nw70[(new BigNumber(22)).toNumber()] = function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_359_v141).Keys.Elements) {
          let _363_v140 = _compr_29;
          if ((_359_v141).contains(_363_v140)) {
            _coll29.push([_module.__default.safeModuloInt(_363_v140, _200_v2),_210_v11]);
          }
        }
        return _coll29;
      }();
      _360_v142 = _nw70;
      let _364_v143;
      let _nw71 = new _module.C2();
      _nw71.__ctor(_360_v142);
      _364_v143 = _nw71;
      let _365_v144;
      _365_v144 = _dafny.Set.fromElements(_200_v2);
      let _366_v145;
      _366_v145 = _dafny.Map.Empty.slice().updateUnsafe(_357_v137,new BigNumber((_365_v144).length));
      (_199_globalState).f9 = (new BigNumber((_366_v145).length)).multipliedBy((_200_v2).plus(new BigNumber(894)));
      let _367_v146;
      let _nw72 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
      _367_v146 = _nw72;
      let _368_v147;
      let _nw73 = new _module.C8();
      _nw73.__ctor(_210_v11);
      _368_v147 = _nw73;
      let _369_v148;
      _369_v148 = _dafny.MultiSet.fromElements(_368_v147, _368_v147);
      let _index52 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_367_v146).length));
      (_367_v146)[_index52] = (_369_v148).Difference((_369_v148).update(_368_v147, _module.__default.abs(_200_v2)));
      let _370_v149;
      _370_v149 = _dafny.Seq.of(_368_v147, _368_v147, _368_v147, _368_v147, _368_v147);
      let _index53 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_367_v146).length));
      (_367_v146)[_index53] = _dafny.MultiSet.FromArray(_370_v149);
      let _371_v150;
      _371_v150 = _dafny.Seq.of((_368_v147).f10, false, _210_v11, (_368_v147).f10, false);
      let _372_i10;
      _372_i10 = _dafny.ZERO;
      L2: {
        while (((((!(_210_v11)) ? ((_368_v147).f10) : ((_368_v147).f10))) ? (_210_v11) : ((_371_v150)[_module.__default.safeIndex(_200_v2, new BigNumber((_371_v150).length))]))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_372_i10)) {
              break L2;
            }
            _372_i10 = (_372_i10).plus(_dafny.ONE);
            let _373_v151;
            let _374_v152;
            let _out7;
            let _out8;
            let _outcollector2 = (_357_v137).m3((_368_v147).f10, _199_globalState);
            _out7 = _outcollector2[0];
            _out8 = _outcollector2[1];
            _373_v151 = _out7;
            _374_v152 = _out8;
            let _375_v153;
            let _nw74 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
            _375_v153 = _nw74;
            let _index54 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_375_v153).length));
            (_375_v153)[_index54] = _dafny.Set.fromElements(_196_v0, _196_v0);
            let _376_v154;
            let _nw75 = new _module.C4();
            _nw75.__ctor(_356_v136);
            _376_v154 = _nw75;
            let _377_v155;
            _377_v155 = _dafny.Map.Empty.slice().updateUnsafe(_376_v154,_196_v0);
            let _378_v156;
            _378_v156 = _dafny.Seq.of(_376_v154);
            let _379_v157;
            _379_v157 = _dafny.Set.fromElements(_196_v0, _196_v0, (((_377_v155).contains((_378_v156)[_module.__default.safeIndex(_373_v151, new BigNumber((_378_v156).length))])) ? ((_377_v155).get((_378_v156)[_module.__default.safeIndex(_373_v151, new BigNumber((_378_v156).length))])) : (_196_v0)));
            let _index55 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_375_v153).length));
            (_375_v153)[_index55] = _379_v157;
            let _380_v158;
            _380_v158 = _dafny.Seq.UnicodeFromString("go");
            let _381_v159;
            _381_v159 = new _dafny.CodePoint('c'.codePointAt(0));
            let _382_v160;
            let _nw76 = Array((new BigNumber(13)).toNumber());
            _nw76[(_dafny.ZERO).toNumber()] = _380_v158;
            _nw76[(_dafny.ONE).toNumber()] = _380_v158;
            _nw76[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_380_v158, _dafny.Seq.UnicodeFromString("gbsy"));
            _nw76[(new BigNumber(3)).toNumber()] = _380_v158;
            _nw76[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-187)), function (_383_i11) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            }), _380_v158);
            _nw76[(new BigNumber(5)).toNumber()] = _380_v158;
            _nw76[(new BigNumber(6)).toNumber()] = _380_v158;
            _nw76[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(909)), function (_384_i12) {
              return new _dafny.CodePoint('c'.codePointAt(0));
            });
            _nw76[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_380_v158, _dafny.Seq.UnicodeFromString("idw"));
            _nw76[(new BigNumber(9)).toNumber()] = _380_v158;
            _nw76[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_380_v158, _dafny.Seq.UnicodeFromString("xnjnwfbu"));
            _nw76[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_380_v158, _module.__default.safeIndex(new BigNumber((_366_v145).length), new BigNumber((_380_v158).length)), _381_v159);
            _nw76[(new BigNumber(12)).toNumber()] = _380_v158;
            _382_v160 = _nw76;
            let _rhs56 = (_373_v151).minus(_200_v2);
            let _rhs57 = _382_v160;
            let _rhs58 = _368_v147;
            _373_v151 = _rhs56;
            _382_v160 = _rhs57;
            _368_v147 = _rhs58;
            let _385_v161;
            let _nw77 = new _module.C8();
            _nw77.__ctor(true);
            _385_v161 = _nw77;
          }
        }
      }
      if (_210_v11) {
        let _386_v162;
        _386_v162 = _dafny.Seq.UnicodeFromString("ultd");
        let _387_v163;
        _387_v163 = _dafny.Seq.of(_386_v162);
        let _388_v164;
        _388_v164 = new _dafny.CodePoint('w'.codePointAt(0));
        let _389_v165;
        _389_v165 = _dafny.Seq.of(_387_v163);
        let _390_v166;
        _390_v166 = _dafny.Map.Empty.slice().updateUnsafe((_368_v147).f10,(_368_v147).f10);
        let _391_v167;
        let _nw78 = Array((new BigNumber(19)).toNumber()).fill(false);
        _391_v167 = _nw78;
        let _392_v168;
        _392_v168 = _module.D7.create_DC17(_386_v162, _391_v167, (_368_v147).f10, _200_v2);
        let _393_v169;
        _393_v169 = _dafny.Map.Empty.slice().updateUnsafe((_368_v147).f10,_200_v2);
        let _394_v170;
        _394_v170 = _dafny.Map.Empty.slice().updateUnsafe(_388_v164,new BigNumber((_393_v169).length));
        let _395_v172;
        _395_v172 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), ((_396_v162) => function (_397_i14) {
          return _396_v162;
        })(_386_v162));
        let _398_v173;
        let _nw79 = Array((new BigNumber(16)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = _387_v163;
        _nw79[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_387_v163, _dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), ((_399_v162) => function (_400_i13) {
          return _399_v162;
        })(_386_v162)));
        _nw79[(new BigNumber(2)).toNumber()] = (((_368_v147).f10) ? (_387_v163) : (_dafny.Seq.of(_dafny.Seq.update(_386_v162, _module.__default.safeIndex(_200_v2, new BigNumber((_386_v162).length)), _388_v164), _386_v162, _dafny.Seq.update(_386_v162, _module.__default.safeIndex((_dafny.ZERO).minus(_200_v2), new BigNumber((_386_v162).length)), _388_v164))));
        _nw79[(new BigNumber(3)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(4)).toNumber()] = (_389_v165)[_module.__default.safeIndex(_module.__default.fm0((_368_v147).f10, _199_globalState), new BigNumber((_389_v165).length))];
        _nw79[(new BigNumber(5)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(6)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(7)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(8)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(9)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(10)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(11)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(12)).toNumber()] = _module.__default.fm28(_388_v164, new BigNumber((_390_v166).length), _200_v2, (_368_v147).f10, _199_globalState);
        _nw79[(new BigNumber(13)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(14)).toNumber()] = _387_v163;
        _nw79[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_386_v162, _386_v162, (_392_v168).dtor_cf29, _module.__default.fm2(_200_v2, _dafny.MultiSet.fromElements((((_394_v170).contains(_388_v164)) ? ((_394_v170).get(_388_v164)) : (new BigNumber((_386_v162).length))), _200_v2), function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of (_365_v144).Elements) {
            let _401_v171 = _compr_30;
            if ((_365_v144).contains(_401_v171)) {
              _coll30.push([(_401_v171).multipliedBy(new BigNumber((_386_v162).length)),_200_v2]);
            }
          }
          return _coll30;
        }(), _199_globalState)), (_395_v172));
        _398_v173 = _nw79;
        let _index56 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_398_v173).length));
        (_398_v173)[_index56] = _dafny.Seq.Concat(_387_v163, _387_v163);
        let _index57 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_398_v173).length));
        (_398_v173)[_index57] = _387_v163;
        let _index58 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_391_v167).length));
        (_391_v167)[_index58] = _210_v11;
        let _index59 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_391_v167).length));
        (_391_v167)[_index59] = (_368_v147).f10;
        let _402_v174;
        let _nw80 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _402_v174 = _nw80;
        let _403_v175;
        _403_v175 = _dafny.Seq.of(_200_v2);
        let _index60 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_402_v174).length));
        (_402_v174)[_index60] = _403_v175;
        let _index61 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_196_v0).length));
        (_196_v0)[_index61] = _200_v2;
        let _404_v176;
        _404_v176 = _dafny.MultiSet.fromElements(!(false), !(!(new BigNumber((_390_v166).length)).isEqualTo(new BigNumber((_386_v162).length))), (_368_v147).f10, (_391_v167)[_module.__default.safeIndex(new BigNumber(70), new BigNumber((_391_v167).length))], !_dafny.Seq.contains(_403_v175, new BigNumber(775)));
        let _index62 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_402_v174).length));
        let _index63 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_196_v0).length));
        let _rhs59 = _module.__default.fm54(false, (_371_v150)[_module.__default.safeIndex(_200_v2, new BigNumber((_371_v150).length))], (_368_v147).f10, _199_globalState);
        let _rhs60 = (((_404_v176).contains(_module.__default.fm1((_368_v147).f10, new _dafny.CodePoint('i'.codePointAt(0)), _199_globalState))) ? ((_404_v176).get(_module.__default.fm1((_368_v147).f10, new _dafny.CodePoint('i'.codePointAt(0)), _199_globalState))) : (_200_v2));
        let _rhs61 = _module.__default.safeDivisionInt(new BigNumber((_403_v175).length), (_dafny.ZERO).minus(_200_v2));
        let _lhs40 = _402_v174;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_402_v174).length));
        let _lhs42 = _199_globalState;
        let _lhs43 = _196_v0;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_196_v0).length));
        _lhs40[_lhs41] = _rhs59;
        _lhs42.f9 = _rhs60;
        _lhs43[_lhs44] = _rhs61;
        let _405_v177;
        _405_v177 = _module.D0.create_DC0((_196_v0)[_module.__default.safeIndex(new BigNumber(814), new BigNumber((_196_v0).length))]);
        let _406_v178;
        _406_v178 = _dafny.Map.Empty.slice().updateUnsafe(_386_v162,_405_v177);
        _406_v178 = (_406_v178).update(_386_v162, _405_v177);
        _393_v169 = _dafny.Map.Empty.slice().updateUnsafe(!((_196_v0)[_module.__default.safeIndex(new BigNumber(814), new BigNumber((_196_v0).length))]).isEqualTo(new BigNumber(560)),_200_v2);
      } else {
        _210_v11 = _210_v11;
        let _407_v179;
        let _nw81 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Set.Empty);
        _407_v179 = _nw81;
        let _408_v180;
        _408_v180 = _dafny.Seq.UnicodeFromString("nvdnk");
        let _409_v181;
        _409_v181 = _dafny.Set.fromElements(_408_v180);
        let _index64 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_407_v179).length));
        (_407_v179)[_index64] = _409_v181;
        let _index65 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_407_v179).length));
        let _rhs62 = _409_v181;
        let _rhs63 = !((_368_v147).f10);
        let _rhs64 = _210_v11;
        let _rhs65 = (_371_v150)[_module.__default.safeIndex(_200_v2, new BigNumber((_371_v150).length))];
        let _rhs66 = _408_v180;
        let _lhs45 = _407_v179;
        let _lhs46 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_407_v179).length));
        let _lhs47 = _199_globalState;
        let _lhs48 = _199_globalState;
        _lhs45[_lhs46] = _rhs62;
        _210_v11 = _rhs63;
        _lhs47.f8 = _rhs64;
        _lhs48.f8 = _rhs65;
        _408_v180 = _rhs66;
        (_199_globalState).f8 = (_368_v147).f10;
        let _410_v182;
        _410_v182 = new _dafny.CodePoint('t'.codePointAt(0));
        let _411_v183;
        _411_v183 = _dafny.Set.fromElements(_410_v182);
        let _412_v184;
        _412_v184 = _411_v183;
        let _source1 = _412_v184;
        let _413___mcc_h0 = _source1;
        let _414_cf27 = _413___mcc_h0;
        let _index66 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_196_v0).length));
        (_196_v0)[_index66] = _200_v2;
        let _415_v185;
        _415_v185 = _dafny.MultiSet.fromElements(_196_v0, _196_v0);
        let _416_v186;
        _416_v186 = _module.D0.create_DC1(_415_v185, _200_v2, _200_v2);
        let _417_v187;
        _417_v187 = _dafny.MultiSet.fromElements((_368_v147).f10);
        let _index67 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_196_v0).length));
        (_196_v0)[_index67] = ((_416_v186).dtor_cf2).minus(new BigNumber(((_417_v187).Difference((_dafny.MultiSet.FromArray(_371_v150)).update((_368_v147).f10, _module.__default.abs(_200_v2)))).cardinality()));
        _408_v180 = _408_v180;
        let _418_v188;
        let _nw82 = new _module.C8();
        _nw82.__ctor(!((_368_v147).f10));
        _418_v188 = _nw82;
        let _419_v189;
        _419_v189 = _module.D2.create_DC8(_200_v2, _410_v182);
        let _420_v190;
        let _nw83 = new _module.C7();
        _nw83.__ctor(true, (_357_v137).f11, _module.__default.fm40((_419_v189).dtor_cf16, _199_globalState), (_368_v147).f10);
        _420_v190 = _nw83;
        let _421_v191;
        _421_v191 = _dafny.Seq.of(_420_v190);
        let _422_v192;
        let _nw84 = new _module.C3();
        _nw84.__ctor(_210_v11, (_368_v147).f10, (_357_v137).f11);
        _422_v192 = _nw84;
        let _423_v193;
        _423_v193 = _dafny.Map.Empty.slice().updateUnsafe(_421_v191,_422_v192);
        (_199_globalState).f8 = !((_423_v193).Merge(_423_v193)).contains(_421_v191);
        let _424_v194;
        let _nw85 = Array((new BigNumber(21)).toNumber());
        _nw85[(_dafny.ZERO).toNumber()] = (_368_v147).f10;
        _nw85[(_dafny.ONE).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(2)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(3)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(4)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(5)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(6)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(7)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(8)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(9)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(10)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(11)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(12)).toNumber()] = _210_v11;
        _nw85[(new BigNumber(13)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(14)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(15)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(16)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(17)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(18)).toNumber()] = (_368_v147).f10;
        _nw85[(new BigNumber(19)).toNumber()] = _210_v11;
        _nw85[(new BigNumber(20)).toNumber()] = _210_v11;
        _424_v194 = _nw85;
        let _425_v195;
        _425_v195 = _module.D7.create_DC17(_dafny.Seq.UnicodeFromString("h"), _424_v194, _210_v11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(929)), ((_426_v2) => function (_427_i15) {
  return _426_v2;
})(_200_v2))).length));
        let _428_v196;
        let _nw86 = Array((new BigNumber(6)).toNumber());
        _nw86[(_dafny.ZERO).toNumber()] = (_368_v147).f10;
        _nw86[(_dafny.ONE).toNumber()] = (_368_v147).f10;
        _nw86[(new BigNumber(2)).toNumber()] = !((_368_v147).f10);
        _nw86[(new BigNumber(3)).toNumber()] = !((((_354_v133).contains(_200_v2)) ? ((_354_v133).get(_200_v2)) : (_210_v11))) || (_210_v11);
        _nw86[(new BigNumber(4)).toNumber()] = (_368_v147).f10;
        _nw86[(new BigNumber(5)).toNumber()] = (_371_v150)[_module.__default.safeIndex((_425_v195).dtor_cf32, new BigNumber((_371_v150).length))];
        _428_v196 = _nw86;
        let _index68 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_428_v196).length));
        (_428_v196)[_index68] = _210_v11;
        let _index69 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_428_v196).length));
        (_428_v196)[_index69] = (_368_v147).f10;
      }
      let _429_v197;
      _429_v197 = _dafny.Seq.UnicodeFromString("negiseois");
      let _430_v198;
      let _nw87 = new _module.C6();
      _nw87.__ctor(_200_v2, _dafny.Seq.Concat(_429_v197, _dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_431_i16) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })));
      _430_v198 = _nw87;
      let _hi2 = _module.__default.safeDivisionInt(_200_v2, (_430_v198).f15);
      for (let _432_i17 = new BigNumber((_430_v198.f16).length); _432_i17.isLessThan(_hi2); _432_i17 = _432_i17.plus(_dafny.ONE)) {
        let _433_v199;
        _433_v199 = new _dafny.CodePoint('j'.codePointAt(0));
        (_199_globalState).f9 = ((_module.__default.fm1((_368_v147).f10, _433_v199, _199_globalState)) ? ((((_368_v147).f10) ? ((_430_v198).f15) : (_200_v2))) : ((_200_v2).plus(new BigNumber((_430_v198.f16).length))));
        let _434_v200;
        _434_v200 = _module.D2.create_DC5(_430_v198.f16);
        _434_v200 = _module.D2.create_DC5(_430_v198.f16);
        (_199_globalState).f9 = (_dafny.ZERO).minus(new BigNumber((_355_v135).cardinality()));
        let _435_v201;
        _435_v201 = _dafny.Seq.of((_430_v198).f15);
        let _436_v202;
        let _nw88 = Array((new BigNumber(5)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = _435_v201;
        _nw88[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_430_v198).f15), _435_v201);
        _nw88[(new BigNumber(2)).toNumber()] = (_module.D5.create_DC13(_435_v201)).dtor_cf25;
        _nw88[(new BigNumber(3)).toNumber()] = _435_v201;
        _nw88[(new BigNumber(4)).toNumber()] = ((_210_v11) ? (_435_v201) : (_435_v201));
        _436_v202 = _nw88;
        let _index70 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_436_v202).length));
        (_436_v202)[_index70] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(159)), ((_437_v198) => function (_438_i18) {
          return (_437_v198).f15;
        })(_430_v198));
        let _index71 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_436_v202).length));
        (_436_v202)[_index71] = _435_v201;
      }
      let _439_v203;
      _439_v203 = _module.D7.create_DC16(_371_v150);
      _210_v11 = (_371_v150)[_module.__default.safeIndex((new BigNumber((_dafny.Seq.update((_439_v203).dtor_cf28, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_429_v197).length), _200_v2, new BigNumber((_430_v198.f16).length))).length), new BigNumber(((_439_v203).dtor_cf28).length)), _210_v11)).length)).multipliedBy(_200_v2), new BigNumber((_371_v150).length))];
      let _440_v204;
      let _nw89 = Array((new BigNumber(10)).toNumber()).fill(false);
      _440_v204 = _nw89;
      let _441_v205;
      _441_v205 = _dafny.Seq.of(_196_v0);
      let _index72 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
      (_440_v204)[_index72] = (new BigNumber((_441_v205).length)).isLessThanOrEqualTo((_430_v198).f15);
      let _442_v206;
      _442_v206 = _dafny.Set.fromElements((_368_v147).f10);
      let _443_v207;
      let _nw90 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
      _443_v207 = _nw90;
      let _444_v208;
      _444_v208 = _dafny.Seq.of(new BigNumber(510), _200_v2);
      let _index73 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_443_v207).length));
      (_443_v207)[_index73] = _444_v208;
      let _445_v209;
      _445_v209 = _module.D7.create_DC17(_dafny.Seq.UnicodeFromString("ervymfer"), _440_v204, _210_v11, _200_v2);
      let _446_v210;
      _446_v210 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_430_v198).f15),_430_v198.f16);
      let _447_v211;
      _447_v211 = _module.D9.create_DC22((_dafny.ZERO).minus((_430_v198).f15), _446_v210);
      let _448_v212;
      _448_v212 = _dafny.Seq.of(_444_v208);
      let _index74 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
      let _index75 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_443_v207).length));
      let _rhs67 = (_445_v209).dtor_cf30;
      let _rhs68 = (((_210_v11) ? ((_368_v147).f10) : (_210_v11))) === (true);
      let _rhs69 = _module.__default.fm53(_199_globalState);
      let _rhs70 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_444_v208, _444_v208), _444_v208), _module.__default.safeIndex(_200_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_444_v208, _444_v208), _444_v208)).length)), (((_355_v135).contains((_447_v211).dtor_cf40)) ? ((_355_v135).get((_447_v211).dtor_cf40)) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber(((_448_v212)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_448_v212).length))]).length)), _module.__default.safeIndex((_430_v198).f15, new BigNumber((_dafny.Seq.of(new BigNumber(((_448_v212)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_448_v212).length))]).length))).length)), _200_v2)).length))));
      let _lhs49 = _440_v204;
      let _lhs50 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
      let _lhs51 = _443_v207;
      let _lhs52 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_443_v207).length));
      _440_v204 = _rhs67;
      _lhs49[_lhs50] = _rhs68;
      _442_v206 = _rhs69;
      _lhs51[_lhs52] = _rhs70;
      let _hi3 = ((_dafny.ZERO).minus(new BigNumber(-60))).multipliedBy(new BigNumber((_429_v197).length));
      for (let _449_i19 = (_430_v198).f15; _449_i19.isLessThan(_hi3); _449_i19 = _449_i19.plus(_dafny.ONE)) {
        _200_v2 = new BigNumber((_dafny.MultiSet.fromElements(_445_v209, _445_v209, _445_v209)).cardinality());
        let _index76 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
        (_440_v204)[_index76] = !(_210_v11);
        let _450_v213;
        let _nw91 = Array((new BigNumber(7)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(672),new BigNumber(((_355_v135).update((_430_v198).f15, _module.__default.abs(new BigNumber(619)))).cardinality()));
        _nw91[(_dafny.ONE).toNumber()] = _359_v141;
        _nw91[(new BigNumber(2)).toNumber()] = _359_v141;
        _nw91[(new BigNumber(3)).toNumber()] = _359_v141;
        _nw91[(new BigNumber(4)).toNumber()] = _359_v141;
        _nw91[(new BigNumber(5)).toNumber()] = _359_v141;
        _nw91[(new BigNumber(6)).toNumber()] = _359_v141;
        _450_v213 = _nw91;
        let _451_v214;
        _451_v214 = new _dafny.CodePoint('a'.codePointAt(0));
        let _452_v215;
        _452_v215 = _dafny.Map.Empty.slice().updateUnsafe(_450_v213,!(false) || (_module.__default.fm1((_368_v147).f10, _451_v214, _199_globalState)));
        _452_v215 = (_452_v215).update(_450_v213, (_368_v147).fm6(_199_globalState));
        (_199_globalState).f9 = (_dafny.ZERO).minus((_430_v198).f15);
      }
      let _453_v216;
      _453_v216 = _module.D13.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(_200_v2,(_368_v147).f10));
      let _pat_let_tv5 = _371_v150;
      let _pat_let_tv6 = _371_v150;
      let _pat_let_tv7 = _430_v198;
      let _pat_let_tv8 = _430_v198;
      let _pat_let_tv9 = _430_v198;
      (_199_globalState).f8 = function (_source2) {
        if (_source2.is_DC30) {
          let _454___mcc_h1 = (_source2).cf51;
          let _455_cf51 = _454___mcc_h1;
          return (_dafny.MultiSet.fromElements(_pat_let_tv5)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_pat_let_tv6));
        } else if (_source2.is_DC31) {
          let _456___mcc_h2 = (_source2).cf52;
          let _457_cf52 = _456___mcc_h2;
          return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("jfaxiys"), (_pat_let_tv7.f16)[_module.__default.safeIndex((_pat_let_tv8).f15, new BigNumber((_pat_let_tv9.f16).length))]);
        } else {
          let _458___mcc_h3 = (_source2).cf50;
          let _459_cf50 = _458___mcc_h3;
          return false;
        }
      }(_453_v216);
      if ((_368_v147).f10) {
        let _460_v217;
        _460_v217 = new _dafny.CodePoint('g'.codePointAt(0));
        (_199_globalState).f8 = _module.__default.fm1((_442_v206).IsProperSubsetOf(_442_v206), _460_v217, _199_globalState);
        (_199_globalState).f8 = (_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))];
        let _461_v218;
        let _nw92 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
        _461_v218 = _nw92;
        let _index77 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_461_v218).length));
        (_461_v218)[_index77] = _371_v150;
        let _462_v219;
        let _init11 = ((_463_v217) => function (_464_i20) {
          return _463_v217;
        })(_460_v217);
        let _nw93 = Array((new BigNumber(24)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw93.length); _i0_11++) {
          _nw93[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _462_v219 = _nw93;
        let _465_v220;
        _465_v220 = _dafny.Map.Empty.slice().updateUnsafe((_368_v147).f10,_462_v219);
        let _466_v221;
        _466_v221 = _dafny.Map.Empty.slice().updateUnsafe((((_465_v220).contains((_368_v147).f10)) ? ((_465_v220).get((_368_v147).f10)) : (_462_v219)),(_dafny.MultiSet.fromElements((_368_v147).f10)).update((_368_v147).f10, _module.__default.abs((_430_v198).f15)));
        let _467_v222;
        _467_v222 = _dafny.Map.Empty.slice().updateUnsafe(_364_v143,false);
        let _468_v223;
        let _nw94 = new _module.C3();
        _nw94.__ctor(_210_v11, (_368_v147).f10, (_357_v137).f11);
        _468_v223 = _nw94;
        let _469_v224;
        _469_v224 = _dafny.Map.Empty.slice().updateUnsafe(_468_v223,_364_v143);
        let _470_v225;
        _470_v225 = _dafny.MultiSet.fromElements((((_467_v222).contains((((_469_v224).contains(_468_v223)) ? ((_469_v224).get(_468_v223)) : (_364_v143)))) ? ((_467_v222).get((((_469_v224).contains(_468_v223)) ? ((_469_v224).get(_468_v223)) : (_364_v143)))) : ((_368_v147).f10)));
        let _471_v226;
        _471_v226 = _dafny.Map.Empty.slice().updateUnsafe(_440_v204,new BigNumber(((_466_v221).update(_462_v219, _470_v225)).length));
        let _472_v227;
        let _nw95 = Array((new BigNumber(17)).toNumber());
        _472_v227 = _nw95;
        let _473_v228;
        _473_v228 = _dafny.Seq.of(_472_v227, _472_v227);
        let _index78 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
        let _index79 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_461_v218).length));
        let _index80 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
        let _rhs71 = !((((_359_v141).contains((_430_v198).f15)) ? ((_359_v141).get((_430_v198).f15)) : ((((_471_v226).contains(_440_v204)) ? ((_471_v226).get(_440_v204)) : (new BigNumber((_dafny.Seq.of((_368_v147).f10)).length)))))).isEqualTo(((_430_v198).f15).plus((_430_v198).f15));
        let _rhs72 = _module.__default.fm0((_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))], _199_globalState);
        let _rhs73 = ((_dafny.Map.Empty.slice().updateUnsafe(_357_v137,(_430_v198).f15)).update(_357_v137, _200_v2)).Merge(_366_v145);
        let _rhs74 = _dafny.Seq.Concat(_dafny.Seq.Concat(_371_v150, _371_v150), _371_v150);
        let _rhs75 = !((_368_v147).f10) || (!((_dafny.ZERO).minus(new BigNumber((_473_v228).length))).isEqualTo((_430_v198).f15));
        let _lhs53 = _440_v204;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
        let _lhs55 = _461_v218;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_461_v218).length));
        let _lhs57 = _440_v204;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length));
        _lhs53[_lhs54] = _rhs71;
        _200_v2 = _rhs72;
        _366_v145 = _rhs73;
        _lhs55[_lhs56] = _rhs74;
        _lhs57[_lhs58] = _rhs75;
        let _474_v229;
        let _nw96 = new _module.C5();
        _nw96.__ctor((_364_v143).f11);
        _474_v229 = _nw96;
        let _475_v230;
        let _nw97 = Array((new BigNumber(8)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _474_v229;
        _nw97[(_dafny.ONE).toNumber()] = _474_v229;
        _nw97[(new BigNumber(2)).toNumber()] = _474_v229;
        _nw97[(new BigNumber(3)).toNumber()] = _474_v229;
        _nw97[(new BigNumber(4)).toNumber()] = _474_v229;
        _nw97[(new BigNumber(5)).toNumber()] = _474_v229;
        _nw97[(new BigNumber(6)).toNumber()] = _474_v229;
        _nw97[(new BigNumber(7)).toNumber()] = _474_v229;
        _475_v230 = _nw97;
        let _476_v231;
        _476_v231 = _475_v230;
        let _rhs76 = (_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))];
        let _rhs77 = _476_v231;
        let _lhs59 = _199_globalState;
        _lhs59.f8 = _rhs76;
        _476_v231 = _rhs77;
        let _index81 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_196_v0).length));
        (_196_v0)[_index81] = (_430_v198).f15;
        let _477_v232;
        let _nw98 = Array((new BigNumber(2)).toNumber());
        _477_v232 = _nw98;
        let _478_v233;
        _478_v233 = _module.D0.create_DC0(_200_v2);
        let _479_v234;
        let _nw99 = new _module.C1();
        _nw99.__ctor((_368_v147).f10, (_368_v147).f10, _478_v233, (_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))]);
        _479_v234 = _nw99;
        let _480_v235;
        _480_v235 = _module.D20.create_DC38(_479_v234);
        let _481_v236;
        _481_v236 = _dafny.Map.Empty.slice().updateUnsafe((_430_v198).f15,(_480_v235).dtor_cf59);
        let _index82 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_477_v232).length));
        (_477_v232)[_index82] = (((_481_v236).contains(_dafny.ONE)) ? ((_481_v236).get(_dafny.ONE)) : (_479_v234));
        let _482_v237;
        _482_v237 = _dafny.MultiSet.fromElements(_468_v223);
        let _index83 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_196_v0).length));
        let _index84 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_477_v232).length));
        let _rhs78 = !(!((_442_v206).IsSubsetOf(_442_v206))) || ((_482_v237).IsProperSubsetOf(_482_v237));
        let _rhs79 = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm39((_430_v198).f15, _199_globalState)).length), (_430_v198).f15);
        let _rhs80 = (_442_v206).Intersect(_dafny.Set.fromElements((_479_v234).f13));
        let _rhs81 = (((_368_v147).fm6(_199_globalState)) ? (_430_v198.f16) : (_430_v198.f16));
        let _rhs82 = _479_v234;
        let _lhs60 = _199_globalState;
        let _lhs61 = _196_v0;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_196_v0).length));
        let _lhs63 = _477_v232;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_477_v232).length));
        _lhs60.f8 = _rhs78;
        _lhs61[_lhs62] = _rhs79;
        _442_v206 = _rhs80;
        _429_v197 = _rhs81;
        _lhs63[_lhs64] = _rhs82;
      } else {
        (_199_globalState).f9 = new BigNumber((_430_v198.f16).length);
        let _483_v238;
        _483_v238 = _dafny.MultiSet.fromElements(!(_359_v141).equals(_359_v141), false);
        _483_v238 = _483_v238;
        _200_v2 = _200_v2;
        let _484_v239;
        _484_v239 = _dafny.Map.Empty.slice().updateUnsafe((_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))],(_430_v198).f15);
        let _485_v240;
        _485_v240 = new _dafny.CodePoint('i'.codePointAt(0));
        let _486_v241;
        _486_v241 = _dafny.Seq.of(_484_v239, _module.__default.fm50(_485_v240, (_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))], (_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))], _199_globalState));
        _486_v241 = _dafny.Seq.Concat(_486_v241, _dafny.Seq.Concat(_486_v241, _486_v241));
        let _487_v242;
        _487_v242 = _dafny.Map.Empty.slice().updateUnsafe(_429_v197,(_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))]);
        let _488_v243;
        _488_v243 = _module.D0.create_DC2(_483_v238, (((_487_v242).contains(_430_v198.f16)) ? ((_487_v242).get(_430_v198.f16)) : ((_368_v147).f10)), (_430_v198).f15, (_368_v147).f10);
        let _489_v244;
        _489_v244 = _dafny.Map.Empty.slice().updateUnsafe(true,_488_v243);
        _489_v244 = (_489_v244).update((_440_v204)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_440_v204).length))], _488_v243);
      }
      process.stdout.write(_dafny.toString((_196_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v1)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState).f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_199_globalState).f2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_199_globalState.f3)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_199_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_199_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_200_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_210_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_211_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_354_v133).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v135).equals(_dafny.MultiSet.fromElements(new BigNumber(-1), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_358_v138).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v141).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v142)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_364_v143).f11)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_365_v144).equals(_dafny.Set.fromElements(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_366_v145).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_367_v146)[new BigNumber(4)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_368_v147).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_369_v148).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_370_v149).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_371_v150, _dafny.Seq.of(true, false, true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_372_i10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_429_v197).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_430_v198).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_430_v198.f16).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_439_v203).dtor_cf28, _dafny.Seq.of(true, false, true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_440_v204)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_441_v205).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_442_v206).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_443_v207)[new BigNumber(4)], _dafny.Seq.of(_dafny.ONE, new BigNumber(-1), new BigNumber(510), new BigNumber(-1), new BigNumber(510), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_444_v208, _dafny.Seq.of(new BigNumber(510), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_445_v209).dtor_cf29).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_445_v209).dtor_cf30)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_445_v209).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_445_v209).dtor_cf32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_446_v210).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.UnicodeFromString("negiseoisqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_447_v211).dtor_cf40));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_447_v211).dtor_cf41).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.UnicodeFromString("negiseoisqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_448_v212, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(510), new BigNumber(-1))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_453_v216).dtor_cf50).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),true))));
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
    static create_DC2(cf4, cf5, cf6, cf7) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
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
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.MultiSet.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC4(cf9) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3(cf8) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC6() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC7(cf11, cf12, cf13, cf14) {
      let $dt = new D2(1);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC8(cf15, cf16) {
      let $dt = new D2(2);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC5(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + this.cf14.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC5" + "(" + this.cf10.toVerbatimString(true) + ")";
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
        return other.$tag === 1 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6();
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
    static create_DC9(cf17) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
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
    static create_DC11(cf19, cf20, cf21) {
      let $dt = new D4(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC12(cf22, cf23, cf24) {
      let $dt = new D4(1);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC10(cf18) {
      let $dt = new D4(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf18) + ")";
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
        return other.$tag === 1 && this.cf22 === other.cf22 && this.cf23 === other.cf23 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC14(cf26) {
      let $dt = new D5(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC13(cf25) {
      let $dt = new D5(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf25) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(false);
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
    static create_DC15(cf27) {
      let $dt = new D6(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf29, cf30, cf31, cf32) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC16(cf28) {
      let $dt = new D7(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + this.cf29.toVerbatimString(true) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(_dafny.Seq.UnicodeFromString(""), [], false, _dafny.ZERO);
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
    static create_DC19(cf34, cf35) {
      let $dt = new D8(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC20(cf36, cf37, cf38) {
      let $dt = new D8(1);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D8(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + this.cf34.toVerbatimString(true) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && this.cf38 === other.cf38;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC19(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC22(cf40, cf41) {
      let $dt = new D9(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC21(cf39) {
      let $dt = new D9(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf39) + ")";
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
        return other.$tag === 1 && this.cf39 === other.cf39;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC22(_dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC24(cf43, cf44, cf45) {
      let $dt = new D10(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC23(cf42) {
      let $dt = new D10(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + this.cf45.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43 && this.cf44 === other.cf44 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf42 === other.cf42;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC24([], [], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC26(cf47) {
      let $dt = new D11(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC25(cf46) {
      let $dt = new D11(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC27(cf48) {
      let $dt = new D11(2);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf46 === other.cf46;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC26(_dafny.ZERO);
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
    static create_DC28(cf49) {
      let $dt = new D12(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf51) {
      let $dt = new D13(0);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC31(cf52) {
      let $dt = new D13(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC29(cf50) {
      let $dt = new D13(2);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC30(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC32(cf53) {
      let $dt = new D14(0);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf53, other.cf53);
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC33(cf54) {
      let $dt = new D15(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf54 === other.cf54;
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
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf55) {
      let $dt = new D16(0);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC34" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf55, other.cf55);
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
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf56) {
      let $dt = new D17(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC35" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56;
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
    static create_DC36(cf57) {
      let $dt = new D18(0);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC36" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf57 === other.cf57;
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
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf58) {
      let $dt = new D19(0);
      $dt.cf58 = cf58;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf58() { return this.cf58; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC37" + "(" + _dafny.toString(this.cf58) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58);
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC39(cf60, cf61, cf62, cf63, cf64) {
      let $dt = new D20(0);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC40(cf65, cf66, cf67, cf68) {
      let $dt = new D20(1);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC38(cf59) {
      let $dt = new D20(2);
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC41(cf69) {
      let $dt = new D20(3);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC39" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC40" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC38" + "(" + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 3) {
        return "D20.DC41" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61 && this.cf62 === other.cf62 && this.cf63 === other.cf63 && this.cf64 === other.cf64;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66) && _dafny.areEqual(this.cf67, other.cf67) && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf59 === other.cf59;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf69, other.cf69);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC39(_dafny.ZERO, false, null, null, false);
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
    static create_DC43(cf71, cf72, cf73) {
      let $dt = new D21(0);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC42(cf70) {
      let $dt = new D21(1);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC44(cf74) {
      let $dt = new D21(2);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC43" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC42" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC44" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf71, other.cf71) && this.cf72 === other.cf72 && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf70 === other.cf70;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC43(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC46(cf76) {
      let $dt = new D22(0);
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC47(cf77) {
      let $dt = new D22(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC45(cf75) {
      let $dt = new D22(2);
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC48(cf78) {
      let $dt = new D22(3);
      $dt.cf78 = cf78;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get is_DC48() { return this.$tag === 3; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf78() { return this.cf78; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC46" + "(" + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC47" + "(" + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC45" + "(" + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC48" + "(" + _dafny.toString(this.cf78) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf76 === other.cf76;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf75 === other.cf75;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf78, other.cf78);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC46(false);
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
      this.f3 = [];
      this.f8 = false;
      this.f9 = _dafny.ZERO;
      this._f0 = false;
      this._f1 = [];
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f17 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f17) {
      let _this = this;
      (_this)._f17 = f17;
      return;
    }
    fm12(p0, p1, globalState) {
      let _this = this;
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(6)), _dafny.Seq.of(new BigNumber(70), new BigNumber(-717), new BigNumber((_dafny.Set.fromElements(new BigNumber(77))).length))));
    };
    fm13(p0, globalState) {
      let _this = this;
      let _source3 = _dafny.MultiSet.fromElements(new BigNumber(570), new BigNumber(969));
      let _490___mcc_h0 = _source3;
      let _491_cf17 = _490___mcc_h0;
      return (_dafny.Set.fromElements((_this).f17, (_this).f17)).Intersect((_module.D4.create_DC10(_dafny.Set.fromElements(true, (_this).f17))).dtor_cf18);
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f12 = _module.D0.Default();
      this._f13 = false;
      this._f19 = false;
      this._f20 = false;
    }
    _parentTraits() {
      return [_module.T2];
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
    __ctor(f19, f20, f12, f13) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm8(globalState) {
      let _this = this;
      let _source4 = _module.D2.create_DC5(_dafny.Seq.UnicodeFromString("fwsgch"));
      if (_source4.is_DC6) {
        return _module.D2.create_DC6();
      } else if (_source4.is_DC7) {
        let _492___mcc_h0 = (_source4).cf11;
        let _493___mcc_h1 = (_source4).cf12;
        let _494___mcc_h2 = (_source4).cf13;
        let _495___mcc_h3 = (_source4).cf14;
        let _496_cf14 = _495___mcc_h3;
        let _497_cf13 = _494___mcc_h2;
        let _498_cf12 = _493___mcc_h1;
        let _499_cf11 = _492___mcc_h0;
        return _module.D2.create_DC6();
      } else if (_source4.is_DC8) {
        let _500___mcc_h4 = (_source4).cf15;
        let _501___mcc_h5 = (_source4).cf16;
        let _502_cf16 = _501___mcc_h5;
        let _503_cf15 = _500___mcc_h4;
        return _module.D2.create_DC6();
      } else {
        let _504___mcc_h6 = (_source4).cf10;
        let _505_cf10 = _504___mcc_h6;
        if ((_this).f19) {
          return _module.D2.create_DC6();
        } else {
          return _module.D2.create_DC6();
        }
      }
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f19),new BigNumber(420))).length))).minus(new BigNumber(111));
    };
    fm29(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber(-454), new BigNumber(983)), (_dafny.ZERO).minus(new BigNumber(((function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber(451), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("anmfxxj")).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-920)), function (_506_i0) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })).length))).length)))).Elements) {
          let _507_v0 = _compr_31;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(451), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("anmfxxj")).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-920)), function (_506_i0) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length))).length))), _507_v0)) {
            _coll31.push([_507_v0,(_this).f13]);
          }
        }
        return _coll31;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-466)), function (_508_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f19)).cardinality())),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(778)), function (_509_i2) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }),(_this).f19)).length))).length);
      })).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f13)).cardinality()), new BigNumber(542), new BigNumber(131), new BigNumber(-34)),!((_this).f13)))).length)));
    };
    fm30(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f13, (_this).f20, (_this).f13)).length),(((_this).f13) ? (new BigNumber((_dafny.Seq.of(new BigNumber(164))).length)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qsxjgq")).length)))))).length);
    };
    m9(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.MultiSet.Empty;
      let _510_v0;
      _510_v0 = new BigNumber(-644);
      if ((_510_v0).isLessThanOrEqualTo(_510_v0)) {
        if ((_this).f20) {
          let _511_v1;
          let _nw100 = Array((new BigNumber(4)).toNumber()).fill(false);
          _511_v1 = _nw100;
          let _512_v2;
          _512_v2 = _dafny.Set.fromElements(_511_v1, _511_v1);
          (globalState).f8 = (new BigNumber(((_dafny.Set.fromElements(_511_v1)).Difference(_512_v2)).length)).isLessThan(new BigNumber(791));
          let _513_v3;
          _513_v3 = new _dafny.CodePoint('w'.codePointAt(0));
          _513_v3 = _513_v3;
          let _514_v4;
          _514_v4 = _module.D1.create_DC4(_513_v3);
          let _515_v5;
          _515_v5 = _dafny.Map.Empty.slice().updateUnsafe(_513_v3,_514_v4);
          _515_v5 = (_515_v5).update(new _dafny.CodePoint('y'.codePointAt(0)), _514_v4);
          let _516_v6;
          _516_v6 = _dafny.Seq.of(_510_v0);
          _516_v6 = _516_v6;
          let _517_v7;
          let _nw101 = new _module.C0();
          _nw101.__ctor((_this).f19);
          _517_v7 = _nw101;
          let _518_v8;
          _518_v8 = _dafny.Seq.of(_517_v7);
          (globalState).f8 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_518_v8, _dafny.Seq.of(_517_v7, _517_v7)), _518_v8);
        } else {
          let _519_v9;
          _519_v9 = _dafny.Set.fromElements(true, (_this).f19);
          let _rhs83 = (((_dafny.Set.fromElements((_this).f20)).IsDisjointFrom(_519_v9)) ? ((_this).f20) : ((_this).f20));
          let _rhs84 = (_this).f19;
          let _rhs85 = _510_v0;
          let _lhs65 = globalState;
          let _lhs66 = globalState;
          let _lhs67 = globalState;
          _lhs65.f8 = _rhs83;
          _lhs66.f8 = _rhs84;
          _lhs67.f9 = _rhs85;
          let _520_v10;
          let _nw102 = new _module.C0();
          _nw102.__ctor(!((new BigNumber((_519_v9).length)).isLessThanOrEqualTo(new BigNumber((_519_v9).length))));
          _520_v10 = _nw102;
          let _521_v11;
          _521_v11 = new _dafny.CodePoint('y'.codePointAt(0));
          let _522_v12;
          _522_v12 = _dafny.Seq.of(_module.__default.fm3(new BigNumber((_519_v9).length), (_520_v10).f17, globalState), _521_v11);
          let _523_v13;
          _523_v13 = _dafny.Seq.of(!((_this).f20), !((_this).f19));
          let _524_v14;
          _524_v14 = _dafny.Map.Empty.slice().updateUnsafe(_510_v0,_510_v0);
          let _525_v15;
          _525_v15 = _dafny.Seq.of(_524_v14);
          let _526_v16;
          _526_v16 = _dafny.Seq.of(_525_v15, _525_v15, _dafny.Seq.of(_524_v14));
          let _527_v18;
          _527_v18 = _dafny.Set.fromElements(_510_v0);
          let _528_v19;
          _528_v19 = _dafny.Map.Empty.slice().updateUnsafe(_510_v0,(_520_v10).f17);
          let _rhs86 = ((new BigNumber((_dafny.Seq.update(_522_v12, _module.__default.safeIndex(_510_v0, new BigNumber((_522_v12).length)), _521_v11)).length)).minus(_510_v0)).minus(_module.__default.safeDivisionInt(_510_v0, new BigNumber((_523_v13).length)));
          let _rhs87 = (_this).f13;
          let _rhs88 = (_510_v0).plus(_510_v0);
          let _rhs89 = _module.__default.safeDivisionInt((_510_v0).minus(new BigNumber(((_526_v16)[_module.__default.safeIndex(_510_v0, new BigNumber((_526_v16).length))]).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm30(function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_dafny.Set.fromElements(new BigNumber((_527_v18).length))).Elements) {
              let _529_v17 = _compr_32;
              if ((_dafny.Set.fromElements(new BigNumber((_527_v18).length))).contains(_529_v17)) {
                _coll32.push([_module.__default.safeDivisionInt(_529_v17, _510_v0),(_this).f20]);
              }
            }
            return _coll32;
          }(), globalState),_528_v19)).length));
          let _lhs68 = globalState;
          let _lhs69 = globalState;
          _lhs68.f9 = _rhs86;
          _lhs69.f8 = _rhs87;
          r0 = _rhs88;
          r0 = _rhs89;
          let _530_v20;
          _530_v20 = _dafny.Map.Empty.slice().updateUnsafe(_522_v12,_module.D4.create_DC10(_519_v9));
          let _531_v21;
          _531_v21 = _module.D4.create_DC10(_519_v9);
          _530_v20 = (_530_v20).update(_522_v12, _531_v21);
          (globalState).f9 = new BigNumber(857);
        }
        let _532_v22;
        _532_v22 = _dafny.Seq.of(_510_v0);
        _532_v22 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), ((_533_v0) => function (_534_i0) {
          return _533_v0;
        })(_510_v0));
        let _535_v23;
        let _nw103 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _535_v23 = _nw103;
        let _536_v24;
        _536_v24 = _dafny.MultiSet.fromElements(_535_v23, _535_v23, _535_v23);
        let _537_v25;
        _537_v25 = _dafny.MultiSet.fromElements((_this).f13);
        let _538_v26;
        _538_v26 = _dafny.Seq.of(_537_v25);
        let _539_v27;
        _539_v27 = _dafny.Seq.of((_538_v26)[_module.__default.safeIndex(_510_v0, new BigNumber((_538_v26).length))]);
        let _540_v28;
        _540_v28 = _module.D0.create_DC1((_536_v24).update(_535_v23, _module.__default.abs(_510_v0)), new BigNumber((_539_v27).length), (new BigNumber((_532_v22).length)).multipliedBy(_510_v0));
        _540_v28 = _540_v28;
        let _541_v29;
        _541_v29 = _dafny.Seq.UnicodeFromString("vfyu");
        _541_v29 = _dafny.Seq.Concat(_dafny.Seq.Concat(_541_v29, _541_v29), _dafny.Seq.Concat(_541_v29, _541_v29));
        let _542_v30;
        _542_v30 = _dafny.Seq.of((_this).f19, (_this).f20);
        let _543_v31;
        _543_v31 = _dafny.MultiSet.fromElements(new BigNumber((_542_v30).length));
        _541_v29 = _module.__default.fm2((_540_v28).dtor_cf3, _543_v31, _dafny.Map.Empty.slice().updateUnsafe(_510_v0,_510_v0), globalState);
      } else {
        let _544_v32;
        let _nw104 = new _module.C0();
        _nw104.__ctor((_this).f19);
        _544_v32 = _nw104;
        let _545_v33;
        _545_v33 = _dafny.Set.fromElements((_dafny.ZERO).minus(_510_v0), new BigNumber((_module.__default.fm31(globalState)).length), _510_v0);
        let _546_v34;
        _546_v34 = _dafny.Seq.of(_545_v33, _dafny.Set.fromElements(_510_v0), _545_v33, (_module.D8.create_DC18(_545_v33)).dtor_cf33, _545_v33);
        _546_v34 = _module.__default.fm32(globalState);
        r0 = _510_v0;
        let _547_v35;
        _547_v35 = _dafny.MultiSet.fromElements((_544_v32).f17);
        _547_v35 = _dafny.MultiSet.fromElements((!(true)) && ((_this).f19));
        let _548_v36;
        _548_v36 = new _dafny.CodePoint('t'.codePointAt(0));
        _548_v36 = _548_v36;
      }
      let _549_v37;
      let _init12 = function (_550_i1) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ejs"), _dafny.Seq.UnicodeFromString("b"));
      };
      let _nw105 = Array((new BigNumber(26)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw105.length); _i0_12++) {
        _nw105[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _549_v37 = _nw105;
      _549_v37 = _549_v37;
      let _551_i2;
      _551_i2 = _dafny.ZERO;
      L3: {
        while (!(!((_this).f13)) || ((_this).f20)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_551_i2)) {
              break L3;
            }
            _551_i2 = (_551_i2).plus(_dafny.ONE);
            let _552_v38;
            _552_v38 = new _dafny.CodePoint('h'.codePointAt(0));
            (globalState).f8 = (_module.__default.fm1((_this).f19, _552_v38, globalState)) && (false);
            let _553_v39;
            _553_v39 = _dafny.Seq.of((_this).f19, (_this).f19);
            let _554_v40;
            _554_v40 = _dafny.Map.Empty.slice().updateUnsafe(_510_v0,_553_v39);
            let _555_v41;
            _555_v41 = _dafny.Seq.of(_510_v0, _510_v0);
            let _556_v42;
            _556_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_553_v39);
            let _557_v43;
            let _nw106 = Array((new BigNumber(17)).toNumber());
            _nw106[(_dafny.ZERO).toNumber()] = (((_554_v40).contains(_510_v0)) ? ((_554_v40).get(_510_v0)) : (_dafny.Seq.of((_this).f13)));
            _nw106[(_dafny.ONE).toNumber()] = _553_v39;
            _nw106[(new BigNumber(2)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(3)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(4)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(5)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(6)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(7)).toNumber()] = (((_this).f13) ? (_dafny.Seq.update(_553_v39, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_555_v41).length)), new BigNumber((_553_v39).length)), true)) : (_553_v39));
            _nw106[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_553_v39, _553_v39);
            _nw106[(new BigNumber(9)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(10)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(11)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_553_v39, _553_v39);
            _nw106[(new BigNumber(13)).toNumber()] = _dafny.Seq.update((((_556_v42).contains((_this).f20)) ? ((_556_v42).get((_this).f20)) : (_553_v39)), _module.__default.safeIndex(_510_v0, new BigNumber(((((_556_v42).contains((_this).f20)) ? ((_556_v42).get((_this).f20)) : (_553_v39))).length)), (_this).f20);
            _nw106[(new BigNumber(14)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(15)).toNumber()] = _553_v39;
            _nw106[(new BigNumber(16)).toNumber()] = _553_v39;
            _557_v43 = _nw106;
            let _index85 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_557_v43).length));
            (_557_v43)[_index85] = _553_v39;
            let _index86 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_557_v43).length));
            (_557_v43)[_index86] = _553_v39;
            let _558_v44;
            _558_v44 = _dafny.MultiSet.fromElements(_510_v0);
            let _559_v45;
            _559_v45 = _module.D1.create_DC4(new _dafny.CodePoint('v'.codePointAt(0)));
            let _560_v46;
            _560_v46 = _dafny.Map.Empty.slice().updateUnsafe(_552_v38,(_this).f20);
            let _561_v47;
            _561_v47 = _dafny.Set.fromElements(_510_v0);
            let _562_v48;
            _562_v48 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_553_v39, _module.__default.safeIndex(new BigNumber((_560_v46).length), new BigNumber((_553_v39).length)), (_this).f19),_561_v47);
            let _rhs90 = (_this).f13;
            let _rhs91 = (_dafny.MultiSet.fromElements((_this).fm9(_559_v45, (((_562_v48).contains(_553_v39)) ? ((_562_v48).get(_553_v39)) : (_561_v47)), globalState))).Union(_558_v44);
            let _rhs92 = _510_v0;
            let _lhs70 = globalState;
            _lhs70.f8 = _rhs90;
            _558_v44 = _rhs91;
            _510_v0 = _rhs92;
            let _563_v49;
            let _init13 = ((_564_v0) => function (_565_i3) {
              return (_564_v0).isEqualTo(_564_v0);
            })(_510_v0);
            let _nw107 = Array((new BigNumber(3)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw107.length); _i0_13++) {
              _nw107[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _563_v49 = _nw107;
            let _index87 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_563_v49).length));
            (_563_v49)[_index87] = (_this).f19;
            let _index88 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_563_v49).length));
            (_563_v49)[_index88] = _module.__default.fm1((_this).f13, new _dafny.CodePoint('c'.codePointAt(0)), globalState);
          }
        }
      }
      let _566_v50;
      let _nw108 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
      _566_v50 = _nw108;
      let _index89 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_566_v50).length));
      (_566_v50)[_index89] = new BigNumber(115);
      let _index90 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_566_v50).length));
      (_566_v50)[_index90] = new BigNumber(-973);
      let _567_v51;
      _567_v51 = _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)));
      let _source5 = (((_this).f20) ? ((((_this).f20) ? (_567_v51) : (_567_v51))) : (_567_v51));
      let _568___mcc_h0 = _source5;
      let _569_cf27 = _568___mcc_h0;
      (globalState).f8 = !((_this).f19);
      let _570_v52;
      _570_v52 = new _dafny.CodePoint('i'.codePointAt(0));
      let _571_v53;
      _571_v53 = _dafny.Seq.of((_this).f20, _module.__default.fm1((_this).f20, _570_v52, globalState), (_this).f20);
      _571_v53 = _dafny.Seq.Concat(_571_v53, _571_v53);
      (globalState).f8 = (_this).f13;
      let _572_v54;
      _572_v54 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f19) ? (_510_v0) : ((_dafny.ZERO).minus(_510_v0))),(_this).f13);
      _572_v54 = (_572_v54).update((_566_v50)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_566_v50).length))], !((_this).f19));
      let _573_v55;
      _573_v55 = _dafny.Seq.UnicodeFromString("q");
      let _rhs93 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-233)), function (_574_i4) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      });
      let _rhs94 = (_510_v0).minus(new BigNumber((_573_v55).length));
      let _lhs71 = globalState;
      _573_v55 = _rhs93;
      _lhs71.f9 = _rhs94;
      r0 = _510_v0;
      r1 = p0;
      return [r0, r1];
    }
    m10(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _575_i0;
      _575_i0 = _dafny.ZERO;
      L4: {
        while ((_this).f19) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_575_i0)) {
              break L4;
            }
            _575_i0 = (_575_i0).plus(_dafny.ONE);
            let _576_v0;
            _576_v0 = new BigNumber(362);
            let _577_v1;
            _577_v1 = _dafny.Seq.UnicodeFromString("gt");
            let _578_v2;
            _578_v2 = new _dafny.CodePoint('l'.codePointAt(0));
            (globalState).f9 = (_576_v0).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_577_v1, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-789)), function (_579_i1) {
              return _module.D4.create_DC12((_this).f20, (_this).f13, (_this).f19);
            })).length), new BigNumber((_577_v1).length)), _578_v2), _577_v1)).length));
            _577_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_577_v1, _577_v1), _577_v1);
            if ((_this).f20) {
              let _580_v3;
              _580_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f19);
              _580_v3 = (_580_v3).update((_this).f13, (_this).f19);
              let _581_v4;
              _581_v4 = _dafny.Set.fromElements((_this).f20);
              let _582_v5;
              _582_v5 = _dafny.Map.Empty.slice().updateUnsafe(_581_v4,(_this).f19);
              let _583_v6;
              _583_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_582_v5);
              let _584_v7;
              _584_v7 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f19) || ((_this).f13),(((_583_v6).contains((_this).f13)) ? ((_583_v6).get((_this).f13)) : (_582_v5)));
              _584_v7 = (_584_v7).update((_this).f20, _582_v5);
              let _585_v8;
              _585_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f12);
              let _586_v9;
              _586_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-747),true);
              let _587_v10;
              _587_v10 = _dafny.Set.fromElements(_576_v0);
              let _588_v11;
              _588_v11 = _dafny.Seq.of(new BigNumber((_587_v10).length), new BigNumber(-703));
              let _589_v12;
              let _nw109 = Array((new BigNumber(15)).toNumber());
              _nw109[(_dafny.ZERO).toNumber()] = _576_v0;
              _nw109[(_dafny.ONE).toNumber()] = new BigNumber(-270);
              _nw109[(new BigNumber(2)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(3)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(4)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(5)).toNumber()] = (((_this).f13) ? (_576_v0) : (new BigNumber((_585_v8).length)));
              _nw109[(new BigNumber(6)).toNumber()] = (_576_v0).multipliedBy(_576_v0);
              _nw109[(new BigNumber(7)).toNumber()] = new BigNumber((_586_v9).length);
              _nw109[(new BigNumber(8)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(9)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_588_v11)[_module.__default.safeIndex(_576_v0, new BigNumber((_588_v11).length))]);
              _nw109[(new BigNumber(11)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(12)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(13)).toNumber()] = _576_v0;
              _nw109[(new BigNumber(14)).toNumber()] = _576_v0;
              _589_v12 = _nw109;
              let _590_v13;
              _590_v13 = _module.D9.create_DC21(_589_v12);
              _589_v12 = (_590_v13).dtor_cf39;
              _588_v11 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_588_v11, _dafny.Seq.update(_588_v11, _module.__default.safeIndex(new BigNumber((_577_v1).length), new BigNumber((_588_v11).length)), _576_v0)), _588_v11), _module.__default.safeIndex((_dafny.ZERO).minus(_576_v0), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_588_v11, _dafny.Seq.update(_588_v11, _module.__default.safeIndex(new BigNumber((_577_v1).length), new BigNumber((_588_v11).length)), _576_v0)), _588_v11)).length)), (((_this).f19) ? (new BigNumber((_580_v3).length)) : (_576_v0)));
              let _591_v14;
              _591_v14 = _dafny.Map.Empty.slice().updateUnsafe(_587_v10,false);
              _591_v14 = (_591_v14).update(_587_v10, (_this).f20);
            } else {
              let _592_v15;
              _592_v15 = _module.D4.create_DC12((_this).f20, (_this).f13, (_this).f19);
              let _rhs95 = _592_v15;
              let _rhs96 = (_this).f19;
              let _rhs97 = _module.__default.safeModuloInt((_576_v0).plus(_576_v0), _576_v0);
              let _lhs72 = globalState;
              let _lhs73 = globalState;
              _592_v15 = _rhs95;
              _lhs72.f8 = _rhs96;
              _lhs73.f9 = _rhs97;
              let _593_v16;
              _593_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_577_v1, _577_v1)).length),_module.__default.safeDivisionInt(_576_v0, _module.__default.fm0((_this).f19, globalState)));
              let _594_v17;
              _594_v17 = _module.D1.create_DC4(_578_v2);
              let _595_v18;
              _595_v18 = _dafny.Set.fromElements(_576_v0, _576_v0);
              let _596_v19;
              _596_v19 = _module.D8.create_DC19(_577_v1, (_this).fm9(_594_v17, _595_v18, globalState));
              let _597_v20;
              let _nw110 = Array((new BigNumber(11)).toNumber());
              _nw110[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_577_v1, _module.__default.safeIndex(_576_v0, new BigNumber((_577_v1).length)), _578_v2);
              _nw110[(_dafny.ONE).toNumber()] = _577_v1;
              _nw110[(new BigNumber(2)).toNumber()] = _577_v1;
              _nw110[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_577_v1, _dafny.Seq.update((_596_v19).dtor_cf34, _module.__default.safeIndex(_576_v0, new BigNumber(((_596_v19).dtor_cf34).length)), _578_v2));
              _nw110[(new BigNumber(4)).toNumber()] = _577_v1;
              _nw110[(new BigNumber(5)).toNumber()] = _577_v1;
              _nw110[(new BigNumber(6)).toNumber()] = _577_v1;
              _nw110[(new BigNumber(7)).toNumber()] = _577_v1;
              _nw110[(new BigNumber(8)).toNumber()] = _577_v1;
              _nw110[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_577_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(421)), ((_598_v2) => function (_599_i2) {
                return _598_v2;
              })(_578_v2)));
              _nw110[(new BigNumber(10)).toNumber()] = _577_v1;
              _597_v20 = _nw110;
              let _600_v21;
              let _nw111 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
              _600_v21 = _nw111;
              let _601_v22;
              _601_v22 = _dafny.Map.Empty.slice().updateUnsafe(_600_v21,_576_v0);
              let _602_v23;
              _602_v23 = _dafny.Set.fromElements(false, (_this).f19, false);
              let _603_v24;
              _603_v24 = _dafny.Seq.of(_577_v1);
              let _604_v25;
              _604_v25 = _dafny.Map.Empty.slice().updateUnsafe(_603_v24,_577_v1);
              let _rhs98 = _dafny.Map.Empty.slice().updateUnsafe(_576_v0,_576_v0);
              let _rhs99 = (_602_v23).equals(_602_v23);
              let _rhs100 = _597_v20;
              let _rhs101 = _601_v22;
              let _rhs102 = (((_604_v25).contains(_dafny.Seq.Concat(_603_v24, _603_v24))) ? ((_604_v25).get(_dafny.Seq.Concat(_603_v24, _603_v24))) : (_dafny.Seq.UnicodeFromString("nor")));
              let _lhs74 = globalState;
              _593_v16 = _rhs98;
              _lhs74.f8 = _rhs99;
              _597_v20 = _rhs100;
              _601_v22 = _rhs101;
              r2 = _rhs102;
              let _605_v26;
              let _init14 = function (_606_i3) {
                return (_this).f13;
              };
              let _nw112 = Array((new BigNumber(22)).toNumber());
              for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw112.length); _i0_14++) {
                _nw112[_i0_14] = _init14(new BigNumber(_i0_14));
              }
              _605_v26 = _nw112;
              let _index91 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_605_v26).length));
              (_605_v26)[_index91] = !(true);
              let _index92 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_605_v26).length));
              (_605_v26)[_index92] = (_this).f13;
              let _607_v27;
              _607_v27 = _dafny.MultiSet.fromElements(_576_v0);
              let _608_v28;
              _608_v28 = _dafny.Map.Empty.slice().updateUnsafe(_576_v0,(new BigNumber((_module.__default.fm2(_576_v0, _607_v27, _593_v16, globalState)).length)).isLessThanOrEqualTo(new BigNumber(72)));
              _608_v28 = (_608_v28).update((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_576_v0, new BigNumber((_602_v23).length)))), !((_this).f13));
              let _609_v29;
              _609_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_607_v27).cardinality()),_578_v2);
              let _610_v30;
              _610_v30 = _dafny.Set.fromElements(_609_v29);
              let _611_v31;
              _611_v31 = _dafny.Seq.of(_609_v29);
              (globalState).f8 = !(_610_v30).equals(function () {
                let _coll33 = new _dafny.Set();
                for (const _compr_33 of (_611_v31).Elements) {
                  let _612_v32 = _compr_33;
                  if (_dafny.Seq.contains(_611_v31, _612_v32)) {
                    _coll33.add(_612_v32);
                  }
                }
                return _coll33;
              }());
            }
            (globalState).f8 = (_this).f20;
          }
        }
      }
      let _613_v33;
      _613_v33 = new BigNumber(-346);
      let _614_v34;
      _614_v34 = _dafny.Seq.UnicodeFromString("ot");
      let _615_v35;
      _615_v35 = _module.D2.create_DC7((((_this).f19) ? ((_this).f13) : ((_this).f13)), _613_v33, true, _614_v34);
      let _rhs103 = (_613_v33).minus(_613_v33);
      let _rhs104 = _615_v35;
      let _lhs75 = globalState;
      _lhs75.f9 = _rhs103;
      _615_v35 = _rhs104;
      let _hi4 = _613_v33;
      for (let _616_i4 = _613_v33; _616_i4.isLessThan(_hi4); _616_i4 = _616_i4.plus(_dafny.ONE)) {
        let _617_v36;
        let _nw113 = Array((new BigNumber(3)).toNumber());
        _nw113[(_dafny.ZERO).toNumber()] = true;
        _nw113[(_dafny.ONE).toNumber()] = (_this).f20;
        _nw113[(new BigNumber(2)).toNumber()] = (_this).f20;
        _617_v36 = _nw113;
        let _618_v37;
        _618_v37 = _dafny.MultiSet.fromElements(_617_v36);
        let _619_v38;
        _619_v38 = _dafny.Map.Empty.slice().updateUnsafe((_618_v37).Intersect(_618_v37),new _dafny.CodePoint('y'.codePointAt(0)));
        let _620_v39;
        _620_v39 = new _dafny.CodePoint('o'.codePointAt(0));
        _619_v38 = (_619_v38).update(_618_v37, _620_v39);
        let _621_v40;
        _621_v40 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13, (_this).f20, (_this).f20, (_this).f20);
        let _622_v41;
        _622_v41 = _dafny.Seq.of(true, (_this).f19, (_this).f19, (_this).f13, (_this).f13);
        (globalState).f9 = new BigNumber(((_621_v40).Intersect((_dafny.MultiSet.FromArray(_622_v41)).Union(_621_v40))).cardinality());
        (globalState).f8 = !(!((_this).f20));
        if ((_module.__default.fm33(_module.__default.fm34((_this).f13, _621_v40, globalState), globalState)).IsProperSubsetOf(_621_v40)) {
          r0 = _620_v39;
          let _623_v42;
          let _nw114 = Array((new BigNumber(5)).toNumber());
          _nw114[(_dafny.ZERO).toNumber()] = _613_v33;
          _nw114[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("fgwlycu"), _module.__default.safeIndex(_616_i4, new BigNumber((_dafny.Seq.UnicodeFromString("fgwlycu")).length)), new _dafny.CodePoint('x'.codePointAt(0)))).length);
          _nw114[(new BigNumber(2)).toNumber()] = _613_v33;
          _nw114[(new BigNumber(3)).toNumber()] = _616_i4;
          _nw114[(new BigNumber(4)).toNumber()] = _613_v33;
          _623_v42 = _nw114;
          let _index93 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_623_v42).length));
          (_623_v42)[_index93] = (new BigNumber(914)).minus(_616_i4);
          let _index94 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_623_v42).length));
          (_623_v42)[_index94] = _613_v33;
          let _624_v43;
          _624_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_614_v34);
          let _625_v44;
          _625_v44 = _dafny.Set.fromElements(_616_i4);
          let _626_v45;
          _626_v45 = _dafny.Set.fromElements(_625_v44);
          let _627_v46;
          _627_v46 = _dafny.Seq.of((_623_v42)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_623_v42).length))]);
          let _628_v47;
          _628_v47 = _dafny.MultiSet.fromElements(new BigNumber((_627_v46).length));
          let _629_v48;
          _629_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_625_v44).length),_616_i4);
          _624_v43 = (_624_v43).update((_626_v45).IsSubsetOf(_626_v45), _dafny.Seq.Concat(_module.__default.fm2(_613_v33, _628_v47, _629_v48, globalState), _614_v34));
          let _630_v49;
          let _nw115 = new _module.C0();
          _nw115.__ctor((_this).f19);
          _630_v49 = _nw115;
          let _631_v50;
          _631_v50 = _module.D5.create_DC13(_627_v46);
          let _632_v51;
          _632_v51 = _dafny.Seq.of(_631_v50);
          let _633_v52;
          _633_v52 = _dafny.MultiSet.fromElements(_module.__default.fm35((_this).f13, (_this).f19, globalState), _632_v51);
          _613_v33 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt((((_633_v52).contains(_632_v51)) ? ((_633_v52).get(_632_v51)) : (_613_v33)), (((_629_v48).contains((_623_v42)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_623_v42).length))])) ? ((_629_v48).get((_623_v42)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_623_v42).length))])) : (_616_i4))), _613_v33));
        } else {
          let _634_v53;
          let _nw116 = new _module.C0();
          _nw116.__ctor(!(new BigNumber(388)).isEqualTo(new BigNumber(827)));
          _634_v53 = _nw116;
          let _635_v54;
          _635_v54 = _dafny.Set.fromElements(_613_v33);
          let _636_v55;
          _636_v55 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-928), _616_i4)),_635_v54);
          _636_v55 = (_636_v55).update(_616_i4, _635_v54);
          let _637_v56;
          _637_v56 = _dafny.Seq.of(_616_i4);
          let _638_v57;
          _638_v57 = _dafny.Seq.of(new BigNumber(-527), _613_v33, new BigNumber((_637_v56).length));
          let _639_v58;
          _639_v58 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_638_v57, _637_v56));
          _639_v58 = ((_dafny.MultiSet.fromElements(_dafny.Seq.of(_616_i4), _module.__default.fm36((_this).f20, globalState))).update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(143)), ((_640_v33) => function (_641_i5) {
            return _640_v33;
          })(_613_v33)), _module.__default.safeIndex(_613_v33, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(143)), ((_642_v33) => function (_643_i5) {
            return _642_v33;
          })(_613_v33))).length)), _613_v33), _module.__default.abs(_616_i4))).Intersect((_639_v58).update(_637_v56, _module.__default.abs(_616_i4)));
          let _644_v59;
          _644_v59 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_622_v41).length), _613_v33),(_638_v57)[_module.__default.safeIndex((_this).fm9(_module.D1.create_DC4(_620_v39), _635_v54, globalState), new BigNumber((_638_v57).length))]);
          _644_v59 = (_644_v59).update(_613_v33, (_dafny.ZERO).minus((_616_i4).multipliedBy(_613_v33)));
          let _index95 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_617_v36).length));
          (_617_v36)[_index95] = (_this).f13;
          let _index96 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_617_v36).length));
          (_617_v36)[_index96] = _module.__default.fm1((_this).f20, _620_v39, globalState);
        }
      }
      let _hi5 = (_613_v33).plus(new BigNumber(-932));
      for (let _645_i6 = _613_v33; _645_i6.isLessThan(_hi5); _645_i6 = _645_i6.plus(_dafny.ONE)) {
        (globalState).f8 = (_this).f19;
        (globalState).f8 = (_613_v33).isLessThanOrEqualTo(new BigNumber(246));
        let _646_v60;
        let _nw117 = Array((new BigNumber(24)).toNumber()).fill(false);
        _646_v60 = _nw117;
        let _647_v61;
        _647_v61 = new _dafny.CodePoint('o'.codePointAt(0));
        let _index97 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_646_v60).length));
        (_646_v60)[_index97] = ((_this).fm9(_module.D1.create_DC4(_647_v61), _module.__default.fm37(false, _613_v33, _613_v33, _645_i6, globalState), globalState)).isLessThanOrEqualTo(_645_i6);
        let _648_v62;
        _648_v62 = _dafny.Map.Empty.slice().updateUnsafe(_645_i6,(_this).f19);
        let _index98 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_646_v60).length));
        (_646_v60)[_index98] = (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-842)), ((_649_v61) => function (_650_i7) {
          return _649_v61;
        })(_647_v61)), _module.__default.safeIndex(_613_v33, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-842)), ((_651_v61) => function (_652_i7) {
          return _651_v61;
        })(_647_v61))).length)), new _dafny.CodePoint('w'.codePointAt(0)))).length), _645_i6)).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_this).fm30(_648_v62, globalState), _645_i6));
        let _653_v63;
        _653_v63 = _dafny.MultiSet.fromElements(_613_v33);
        let _654_v64;
        _654_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_614_v34).length),_613_v33);
        let _655_v65;
        _655_v65 = _dafny.Map.Empty.slice().updateUnsafe(_654_v64,new BigNumber((_614_v34).length));
        let _656_v66;
        _656_v66 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(new BigNumber((_dafny.Seq.of((_646_v60)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_646_v60).length))], (_this).f20, (_this).f20, (_this).f19)).length), (_this).f13, globalState),_655_v65);
        (globalState).f9 = new BigNumber(((_653_v63).update(_645_i6, _module.__default.abs(new BigNumber(((_655_v65).Merge((((_656_v66).contains(_647_v61)) ? ((_656_v66).get(_647_v61)) : (_655_v65)))).length)))).cardinality());
      }
      let _657_v67;
      let _init15 = ((_658_v33) => function (_659_i8) {
        return (_659_i8).minus(_658_v33);
      })(_613_v33);
      let _nw118 = Array((new BigNumber(11)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw118.length); _i0_15++) {
        _nw118[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _657_v67 = _nw118;
      let _660_v68;
      _660_v68 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rnrfgxihe"),(_this).f13);
      let _index99 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length));
      (_657_v67)[_index99] = new BigNumber((_660_v68).length);
      let _661_v69;
      _661_v69 = _dafny.Seq.of(_613_v33);
      let _662_v70;
      _662_v70 = _dafny.MultiSet.fromElements(_661_v69, _661_v69, _661_v69, ((false) ? (_661_v69) : (_661_v69)));
      let _index100 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length));
      (_657_v67)[_index100] = (((_662_v70).contains(_661_v69)) ? ((_662_v70).get(_661_v69)) : (_613_v33));
      if ((_this).f13) {
        let _663_v71;
        _663_v71 = _dafny.Set.fromElements((_this).f20);
        (globalState).f8 = ((_dafny.ZERO).minus(new BigNumber((_663_v71).length))).isLessThan((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]);
        let _664_v72;
        _664_v72 = new _dafny.CodePoint('l'.codePointAt(0));
        let _665_v73;
        _665_v73 = _dafny.Set.fromElements(_664_v72);
        let _666_v74;
        _666_v74 = _dafny.MultiSet.fromElements((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))], _613_v33);
        let _667_v75;
        _667_v75 = _666_v74;
        let _668_v76;
        _668_v76 = _module.D4.create_DC12(true, _module.__default.fm1((_this).f19, _664_v72, globalState), (_this).f19);
        let _669_v77;
        _669_v77 = _dafny.Map.Empty.slice().updateUnsafe(_613_v33,_613_v33);
        let _670_v78;
        let _nw119 = new _module.C0();
        _nw119.__ctor((_this).f13);
        _670_v78 = _nw119;
        let _671_v79;
        _671_v79 = _dafny.Set.fromElements(_670_v78);
        let _672_v80;
        _672_v80 = _dafny.Seq.of(_663_v71);
        let _673_v81;
        _673_v81 = _dafny.Map.Empty.slice().updateUnsafe(!((_670_v78).f17),(_this).f20);
        let _674_v82;
        _674_v82 = _dafny.MultiSet.fromElements((_this).f19);
        let _675_v83;
        _675_v83 = _module.D0.create_DC2(_674_v82, (_670_v78).f17, new BigNumber((_663_v71).length), true);
        let _676_v84;
        _676_v84 = _module.D8.create_DC20(_675_v83, (_this).f20, _module.__default.fm1((_this).f19, _664_v72, globalState));
        let _677_v85;
        let _nw120 = Array((new BigNumber(29)).toNumber());
        _nw120[(_dafny.ZERO).toNumber()] = ((_667_v75)).IsDisjointFrom(_666_v74);
        _nw120[(_dafny.ONE).toNumber()] = true;
        _nw120[(new BigNumber(2)).toNumber()] = (_this).f19;
        _nw120[(new BigNumber(3)).toNumber()] = true;
        _nw120[(new BigNumber(4)).toNumber()] = !((_663_v71).IsSubsetOf(_663_v71));
        _nw120[(new BigNumber(5)).toNumber()] = (_this).f13;
        _nw120[(new BigNumber(6)).toNumber()] = (_this).f20;
        _nw120[(new BigNumber(7)).toNumber()] = (_this).f19;
        _nw120[(new BigNumber(8)).toNumber()] = (_this).f19;
        _nw120[(new BigNumber(9)).toNumber()] = (_this).f20;
        _nw120[(new BigNumber(10)).toNumber()] = (_this).f20;
        _nw120[(new BigNumber(11)).toNumber()] = (_this).f13;
        _nw120[(new BigNumber(12)).toNumber()] = !((_this).f13);
        _nw120[(new BigNumber(13)).toNumber()] = (!((_this).f13)) === ((_this).f19);
        _nw120[(new BigNumber(14)).toNumber()] = (function (_pat_let6_0) {
          return function (_678_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_679_dt__update_hcf22_h0) {
                return function (_pat_let8_0) {
                  return function (_680_dt__update_hcf23_h0) {
                    return _module.D4.create_DC12(_679_dt__update_hcf22_h0, _680_dt__update_hcf23_h0, (_678_dt__update__tmp_h0).dtor_cf24);
                  }(_pat_let8_0);
                }((_this).f13);
              }(_pat_let7_0);
            }((_this).f13);
          }(_pat_let6_0);
        }(_668_v76)).dtor_cf23;
        _nw120[(new BigNumber(15)).toNumber()] = _module.__default.fm1(true, _664_v72, globalState);
        _nw120[(new BigNumber(16)).toNumber()] = (_this).f19;
        _nw120[(new BigNumber(17)).toNumber()] = ((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]).isEqualTo(new BigNumber((_669_v77).length));
        _nw120[(new BigNumber(18)).toNumber()] = (_this).f20;
        _nw120[(new BigNumber(19)).toNumber()] = (_this).f19;
        _nw120[(new BigNumber(20)).toNumber()] = !(_671_v79).equals(_dafny.Set.fromElements(_670_v78));
        _nw120[(new BigNumber(21)).toNumber()] = !((_this).f20) || (true);
        _nw120[(new BigNumber(22)).toNumber()] = ((_672_v80)[_module.__default.safeIndex(new BigNumber(-231), new BigNumber((_672_v80).length))]).IsDisjointFrom(_dafny.Set.fromElements((_670_v78).f17));
        _nw120[(new BigNumber(23)).toNumber()] = !((_this).f20) || (_module.__default.fm1((_this).f20, new _dafny.CodePoint('u'.codePointAt(0)), globalState));
        _nw120[(new BigNumber(24)).toNumber()] = (((((_673_v81).contains((_this).f13)) ? ((_673_v81).get((_this).f13)) : ((_670_v78).f17))) ? ((_675_v83).dtor_cf7) : ((_670_v78).f17));
        _nw120[(new BigNumber(25)).toNumber()] = !((_670_v78).f17);
        _nw120[(new BigNumber(26)).toNumber()] = (_670_v78).f17;
        _nw120[(new BigNumber(27)).toNumber()] = (_this).f20;
        _nw120[(new BigNumber(28)).toNumber()] = (_676_v84).dtor_cf37;
        _677_v85 = _nw120;
        let _index101 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_677_v85).length));
        (_677_v85)[_index101] = !((_this).f19);
        let _index102 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_677_v85).length));
        let _index103 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length));
        let _rhs105 = _665_v73;
        let _rhs106 = (_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))];
        let _rhs107 = (((_this).f20) ? ((_670_v78).f17) : (true));
        let _rhs108 = new BigNumber((_673_v81).length);
        let _lhs76 = globalState;
        let _lhs77 = _677_v85;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_677_v85).length));
        let _lhs79 = _657_v67;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length));
        _665_v73 = _rhs105;
        _lhs76.f9 = _rhs106;
        _lhs77[_lhs78] = _rhs107;
        _lhs79[_lhs80] = _rhs108;
        _613_v33 = (_dafny.ZERO).minus((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]);
        let _index104 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_677_v85).length));
        (_677_v85)[_index104] = (_677_v85)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_677_v85).length))];
        let _681_v86;
        _681_v86 = _module.D9.create_DC21(_657_v67);
        let _rhs109 = _681_v86;
        let _rhs110 = (_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))];
        _681_v86 = _rhs109;
        _613_v33 = _rhs110;
      } else {
        _613_v33 = new BigNumber(193);
        (globalState).f8 = ((_this).f20) || ((_this).f13);
        let _index105 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length));
        (_657_v67)[_index105] = ((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]).plus((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]);
        let _682_v87;
        let _nw121 = Array((new BigNumber(10)).toNumber()).fill([]);
        _682_v87 = _nw121;
        let _683_v88;
        _683_v88 = new _dafny.CodePoint('i'.codePointAt(0));
        let _684_v89;
        _684_v89 = _module.D1.create_DC4(_683_v88);
        let _685_v90;
        _685_v90 = _dafny.Set.fromElements((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))], new BigNumber(415), (_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]);
        let _rhs111 = (_this).fm9(_684_v89, _685_v90, globalState);
        let _rhs112 = _682_v87;
        _613_v33 = _rhs111;
        _682_v87 = _rhs112;
        let _686_v91;
        _686_v91 = _dafny.MultiSet.fromElements(_614_v34, _614_v34);
        let _687_v92;
        _687_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_686_v91).cardinality()),_683_v88);
        let _688_v93;
        _688_v93 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_687_v92).Merge(_687_v92));
        let _689_v94;
        let _init16 = function (_690_i9) {
          return (_dafny.Set.fromElements(!((_this).f20))).IsProperSubsetOf(_dafny.Set.fromElements((_this).f20, true));
        };
        let _nw122 = Array((new BigNumber(4)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw122.length); _i0_16++) {
          _nw122[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _689_v94 = _nw122;
        let _691_v95;
        _691_v95 = _dafny.Seq.of((_this).f19);
        let _rhs113 = (_688_v93).update(!((_this).f19), _687_v92);
        let _rhs114 = (_module.D7.create_DC17(_dafny.Seq.UnicodeFromString("p"), _689_v94, (_this).f19, new BigNumber((_691_v95).length))).dtor_cf32;
        let _rhs115 = (((((_this).f13) ? ((_this).f13) : (_module.__default.fm1(true, _module.__default.fm3((_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))], false, globalState), globalState)))) ? (_657_v67) : (_657_v67));
        let _rhs116 = _689_v94;
        let _rhs117 = (!((_this).f19)) && ((new BigNumber((function () {
          let _coll34 = new _dafny.Map();
          for (const _compr_34 of _dafny.IntegerRange(new BigNumber(160), new BigNumber(441))) {
            let _692_v96 = _compr_34;
            if (((new BigNumber(160)).isLessThanOrEqualTo(_692_v96)) && ((_692_v96).isLessThan(new BigNumber(441)))) {
              _coll34.push([(_692_v96).minus(_613_v33),(_657_v67)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_657_v67).length))]]);
            }
          }
          return _coll34;
        }()).length)).isLessThanOrEqualTo(new BigNumber((_614_v34).length)));
        let _lhs81 = globalState;
        let _lhs82 = globalState;
        _688_v93 = _rhs113;
        _lhs81.f9 = _rhs114;
        _657_v67 = _rhs115;
        _689_v94 = _rhs116;
        _lhs82.f8 = _rhs117;
      }
      r0 = new _dafny.CodePoint('y'.codePointAt(0));
      let _693_v97;
      _693_v97 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("r"), _614_v34);
      r1 = (_661_v69)[_module.__default.safeIndex(new BigNumber((_693_v97).length), new BigNumber((_661_v69).length))];
      r2 = _614_v34;
      return [r0, r1, r2];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f11 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    fm19(globalState) {
      let _this = this;
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), function (_694_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      });
    };
    fm20(p0, p1, globalState) {
      let _this = this;
      if ((false) || (!(false))) {
        return (new BigNumber(-248)).multipliedBy(new BigNumber(-552));
      } else {
        return _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(525))).length))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-220),new BigNumber((_dafny.MultiSet.fromElements(true, !(false))).cardinality()))).length));
      }
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _module.D2.Default();
      let r3 = false;
      (globalState).f9 = p0;
      let _695_v0;
      _695_v0 = true;
      let _696_v1;
      _696_v1 = _dafny.MultiSet.fromElements(true, _695_v0, _695_v0);
      let _697_v2;
      _697_v2 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(962),new BigNumber((_696_v1).cardinality())));
      let _698_v3;
      _698_v3 = _dafny.Map.Empty.slice().updateUnsafe((_697_v2)[_module.__default.safeIndex(_module.__default.fm0(_695_v0, globalState), new BigNumber((_697_v2).length))],p0);
      let _699_v4;
      _699_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      _698_v3 = (_698_v3).update(_699_v4, new BigNumber(581));
      (globalState).f8 = _695_v0;
      let _700_v5;
      _700_v5 = _dafny.Seq.UnicodeFromString("jfnnndul");
      _700_v5 = _dafny.Seq.UnicodeFromString("bogafiid");
      let _701_v6;
      _701_v6 = _module.D0.create_DC0(new BigNumber(715));
      let _pat_let_tv10 = p0;
      _701_v6 = function (_pat_let9_0) {
        return function (_702_dt__update__tmp_h0) {
          return function (_pat_let10_0) {
            return function (_703_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_703_dt__update_hcf0_h0);
            }(_pat_let10_0);
          }(_pat_let_tv10);
        }(_pat_let9_0);
      }(_701_v6);
      _695_v0 = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("wnktrnh"), new _dafny.CodePoint('r'.codePointAt(0)));
      let _704_v7;
      _704_v7 = new _dafny.CodePoint('k'.codePointAt(0));
      r0 = _704_v7;
      r1 = p0;
      let _705_v8;
      _705_v8 = _module.D2.create_DC8(p0, _704_v7);
      r2 = _705_v8;
      r3 = _695_v0;
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = false;
      let r2 = false;
      let _706_v0;
      _706_v0 = true;
      let _707_v1;
      let _nw123 = new _module.C0();
      _nw123.__ctor(!(!(_706_v0)));
      _707_v1 = _nw123;
      let _708_v2;
      _708_v2 = new BigNumber(354);
      let _709_v3;
      _709_v3 = _dafny.Seq.of((_707_v1).f17, _706_v0, _706_v0);
      let _hi6 = _module.__default.safeModuloInt(new BigNumber((_709_v3).length), _708_v2);
      for (let _710_i0 = _708_v2; _710_i0.isLessThan(_hi6); _710_i0 = _710_i0.plus(_dafny.ONE)) {
        let _711_v4;
        let _nw124 = Array((new BigNumber(14)).toNumber()).fill(_module.D1.Default());
        _711_v4 = _nw124;
        let _712_v5;
        _712_v5 = _dafny.Seq.UnicodeFromString("lfrdl");
        let _713_v6;
        _713_v6 = new _dafny.CodePoint('c'.codePointAt(0));
        let _714_v7;
        _714_v7 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_712_v5, _module.__default.safeIndex(new BigNumber(592), new BigNumber((_712_v5).length)), _713_v6)).length), _710_i0, _710_i0);
        let _715_v8;
        _715_v8 = _dafny.Map.Empty.slice().updateUnsafe(_711_v4,(_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_714_v7).length), _710_i0))));
        _715_v8 = _dafny.Map.Empty.slice().updateUnsafe(_711_v4,_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("wnfg"), _module.__default.safeIndex(_708_v2, new BigNumber((_dafny.Seq.UnicodeFromString("wnfg")).length)), new _dafny.CodePoint('f'.codePointAt(0)))).length), _708_v2));
        let _716_v9;
        let _nw125 = Array((new BigNumber(8)).toNumber()).fill(false);
        _716_v9 = _nw125;
        let _index106 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_716_v9).length));
        (_716_v9)[_index106] = _706_v0;
        let _717_v10;
        _717_v10 = _dafny.MultiSet.fromElements((_707_v1).f17);
        let _718_v11;
        _718_v11 = _dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,_710_i0);
        let _719_v12;
        _719_v12 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_709_v3), (_717_v10).update((_707_v1).f17, _module.__default.abs((((_718_v11).contains((_707_v1).f17)) ? ((_718_v11).get((_707_v1).f17)) : (_708_v2)))));
        let _index107 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_716_v9).length));
        (_716_v9)[_index107] = (_717_v10).IsProperSubsetOf((_719_v12)[_module.__default.safeIndex(_710_i0, new BigNumber((_719_v12).length))]);
        let _720_v13;
        let _nw126 = new _module.C0();
        _nw126.__ctor(!(_708_v2).isEqualTo(_708_v2));
        _720_v13 = _nw126;
        let _721_v14;
        _721_v14 = _module.D2.create_DC7(_706_v0, new BigNumber((_717_v10).cardinality()), (_707_v1).f17, _712_v5);
        let _rhs118 = (_707_v1).f17;
        let _rhs119 = (_721_v14).dtor_cf13;
        let _lhs83 = globalState;
        r2 = _rhs118;
        _lhs83.f8 = _rhs119;
      }
      _708_v2 = new BigNumber(978);
      let _hi7 = (_dafny.ZERO).minus(_708_v2);
      for (let _722_i1 = (_708_v2).multipliedBy(new BigNumber(600)); _722_i1.isLessThan(_hi7); _722_i1 = _722_i1.plus(_dafny.ONE)) {
        let _723_v15;
        _723_v15 = new _dafny.CodePoint('o'.codePointAt(0));
        (globalState).f8 = (_module.D4.create_DC12(_706_v0, !((_707_v1).f17), !(_module.__default.fm1((_707_v1).f17, _723_v15, globalState)))).dtor_cf24;
        r2 = _706_v0;
        let _724_v16;
        _724_v16 = _dafny.Set.fromElements(_708_v2, _708_v2);
        let _725_v17;
        _725_v17 = _dafny.MultiSet.fromElements(_722_i1, new BigNumber((_724_v16).length));
        let _726_v18;
        _726_v18 = _module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-209)), function (_727_i2) {
  return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(135)), function (_728_i3) {
    return new _dafny.CodePoint('i'.codePointAt(0));
  })).length));
}));
        let _729_v20;
        _729_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_706_v0, globalState),_706_v0);
        let _730_v21;
        let _nw127 = Array((new BigNumber(9)).toNumber());
        _nw127[(_dafny.ZERO).toNumber()] = (_707_v1).f17;
        _nw127[(_dafny.ONE).toNumber()] = _706_v0;
        _nw127[(new BigNumber(2)).toNumber()] = (_707_v1).f17;
        _nw127[(new BigNumber(3)).toNumber()] = (_707_v1).f17;
        _nw127[(new BigNumber(4)).toNumber()] = _706_v0;
        _nw127[(new BigNumber(5)).toNumber()] = _module.__default.fm1((_707_v1).f17, _723_v15, globalState);
        _nw127[(new BigNumber(6)).toNumber()] = _706_v0;
        _nw127[(new BigNumber(7)).toNumber()] = (_725_v17).IsDisjointFrom(_dafny.MultiSet.FromArray((_726_v18).dtor_cf25));
        _nw127[(new BigNumber(8)).toNumber()] = !(function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of _dafny.IntegerRange(new BigNumber(714), new BigNumber(223))) {
            let _731_v19 = _compr_35;
            if (((new BigNumber(714)).isLessThanOrEqualTo(_731_v19)) && ((_731_v19).isLessThan(new BigNumber(223)))) {
              _coll35.push([_module.__default.safeDivisionInt(_731_v19, _722_i1),(_707_v1).f17]);
            }
          }
          return _coll35;
        }()).equals(_729_v20);
        _730_v21 = _nw127;
        _730_v21 = _730_v21;
        let _source6 = _module.D2.create_DC8(_708_v2, _723_v15);
        if (_source6.is_DC6) {
          let _732_v22;
          let _nw128 = Array((new BigNumber(23)).toNumber()).fill(_module.D0.Default());
          _732_v22 = _nw128;
          let _733_v23;
          let _init17 = ((_734_i1) => function (_735_i4) {
            return (_735_i4).minus(_734_i1);
          })(_722_i1);
          let _nw129 = Array((new BigNumber(17)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw129.length); _i0_17++) {
            _nw129[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _733_v23 = _nw129;
          let _736_v24;
          _736_v24 = _dafny.MultiSet.fromElements(_733_v23);
          let _737_v25;
          _737_v25 = _module.D0.create_DC1(_736_v24, (_this).fm20((_707_v1).f17, _722_i1, globalState), new BigNumber(-331));
          let _index108 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_732_v22).length));
          (_732_v22)[_index108] = _737_v25;
          let _index109 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_732_v22).length));
          (_732_v22)[_index109] = _737_v25;
          let _738_v26;
          _738_v26 = _dafny.Seq.UnicodeFromString("aolhi");
          _738_v26 = _dafny.Seq.UnicodeFromString("fr");
          let _739_v27;
          _739_v27 = _dafny.MultiSet.fromElements((_707_v1).f17);
          let _740_v28;
          _740_v28 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, (_707_v1).f17), _739_v27);
          let _741_v29;
          _741_v29 = _module.D0.create_DC2(_739_v27, (_707_v1).f17, _722_i1, (_707_v1).f17);
          let _742_v30;
          _742_v30 = _dafny.Map.Empty.slice().updateUnsafe(_708_v2,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-978)), ((_743_v15) => function (_744_i5) {
            return _743_v15;
          })(_723_v15)));
          let _rhs120 = (_dafny.Set.fromElements((_741_v29).dtor_cf4)).IsSubsetOf(_740_v28);
          let _rhs121 = (_module.__default.fm21((_dafny.ZERO).minus((_this).fm20((_707_v1).f17, new BigNumber(203), globalState)), globalState)).contains(function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of (_742_v30).Keys.Elements) {
              let _745_v31 = _compr_36;
              if ((_742_v30).contains(_745_v31)) {
                _coll36.add(_module.__default.safeDivisionInt(_745_v31, new BigNumber((_dafny.Set.fromElements(new BigNumber(938))).length)));
              }
            }
            return _coll36;
          }());
          let _lhs84 = globalState;
          _706_v0 = _rhs120;
          _lhs84.f8 = _rhs121;
          (globalState).f9 = _722_i1;
        } else if (_source6.is_DC7) {
          let _746___mcc_h0 = (_source6).cf11;
          let _747___mcc_h1 = (_source6).cf12;
          let _748___mcc_h2 = (_source6).cf13;
          let _749___mcc_h3 = (_source6).cf14;
          let _750_cf14 = _749___mcc_h3;
          let _751_cf13 = _748___mcc_h2;
          let _752_cf12 = _747___mcc_h1;
          let _753_cf11 = _746___mcc_h0;
          _750_cf14 = _750_cf14;
          r1 = _753_cf11;
          let _754_v32;
          let _nw130 = new _module.C0();
          _nw130.__ctor((_dafny.Set.fromElements(_752_cf12, new BigNumber(-941))).IsProperSubsetOf(_724_v16));
          _754_v32 = _nw130;
          (globalState).f9 = _module.__default.safeModuloInt(_708_v2, new BigNumber(271));
        } else if (_source6.is_DC8) {
          let _755___mcc_h4 = (_source6).cf15;
          let _756___mcc_h5 = (_source6).cf16;
          let _757_cf16 = _756___mcc_h5;
          let _758_cf15 = _755___mcc_h4;
          let _index110 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_730_v21).length));
          (_730_v21)[_index110] = _706_v0;
          let _index111 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_730_v21).length));
          (_730_v21)[_index111] = (((false) ? (_758_cf15) : (new BigNumber((_725_v17).cardinality())))).isLessThan(_708_v2);
          let _759_v33;
          let _nw131 = Array((new BigNumber(3)).toNumber());
          _nw131[(_dafny.ZERO).toNumber()] = _706_v0;
          _nw131[(_dafny.ONE).toNumber()] = (_707_v1).f17;
          _nw131[(new BigNumber(2)).toNumber()] = (_707_v1).f17;
          _759_v33 = _nw131;
          let _760_v34;
          _760_v34 = _dafny.Map.Empty.slice().updateUnsafe(_758_cf15,_759_v33);
          let _761_v35;
          _761_v35 = _dafny.Seq.of(_708_v2);
          let _762_v36;
          _762_v36 = _dafny.Seq.UnicodeFromString("ormmod");
          let _index112 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_730_v21).length));
          (_730_v21)[_index112] = ((!(_760_v34).equals(_760_v34)) ? (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.update(_761_v35, _module.__default.safeIndex(new BigNumber((_762_v36).length), new BigNumber((_761_v35).length)), _722_i1), _module.__default.safeIndex(_758_cf15, new BigNumber((_dafny.Seq.update(_761_v35, _module.__default.safeIndex(new BigNumber((_762_v36).length), new BigNumber((_761_v35).length)), _722_i1)).length)), _722_i1), _761_v35)) : ((_707_v1).f17));
          let _763_v37;
          let _nw132 = new _module.C0();
          _nw132.__ctor((_707_v1).f17);
          _763_v37 = _nw132;
          let _764_v38;
          _764_v38 = _module.D5.create_DC14(true);
          let _765_v39;
          _765_v39 = _dafny.MultiSet.fromElements((_764_v38).dtor_cf26);
          (globalState).f8 = (_module.D0.create_DC2(_765_v39, (_707_v1).f17, _722_i1, _706_v0)).dtor_cf5;
        } else {
          let _766___mcc_h6 = (_source6).cf10;
          let _767_cf10 = _766___mcc_h6;
          let _768_v40;
          let _nw133 = new _module.C0();
          _nw133.__ctor(false);
          _768_v40 = _nw133;
          _723_v15 = _723_v15;
          let _769_v41;
          let _nw134 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _769_v41 = _nw134;
          _769_v41 = _769_v41;
          let _rhs122 = (_707_v1).f17;
          let _rhs123 = _722_i1;
          let _rhs124 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_722_i1).multipliedBy(new BigNumber((_dafny.Seq.Concat(_767_cf10, _767_cf10)).length))));
          let _rhs125 = _708_v2;
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          let _lhs87 = globalState;
          _lhs85.f8 = _rhs122;
          _708_v2 = _rhs123;
          _lhs86.f9 = _rhs124;
          _lhs87.f9 = _rhs125;
        }
      }
      if (!(true)) {
        (globalState).f8 = !((_708_v2).isLessThanOrEqualTo((_dafny.ZERO).minus(((_706_v0) ? (_708_v2) : (_708_v2)))));
        let _770_v42;
        _770_v42 = _dafny.MultiSet.fromElements((_707_v1).f17, (_707_v1).f17, _706_v0, true);
        let _771_v43;
        _771_v43 = new _dafny.CodePoint('y'.codePointAt(0));
        let _772_v44;
        _772_v44 = _dafny.MultiSet.fromElements(_771_v43);
        let _773_v45;
        _773_v45 = _dafny.Seq.of((_module.D0.create_DC2(_770_v42, (_module.D5.create_DC14(false)).dtor_cf26, _708_v2, (_707_v1).f17)).dtor_cf6, (((_772_v44).contains(_module.__default.fm3(_708_v2, (_707_v1).f17, globalState))) ? ((_772_v44).get(_module.__default.fm3(_708_v2, (_707_v1).f17, globalState))) : (_708_v2)));
        let _774_v46;
        _774_v46 = _dafny.Set.fromElements((_dafny.ZERO).minus(_708_v2));
        let _775_v47;
        _775_v47 = _module.D0.create_DC2(_770_v42, _706_v0, _708_v2, _706_v0);
        let _776_v48;
        _776_v48 = _dafny.Map.Empty.slice().updateUnsafe((_775_v47).dtor_cf7,(_707_v1).f17);
        let _777_v49;
        _777_v49 = _dafny.MultiSet.fromElements(_708_v2);
        _706_v0 = !(!((_dafny.MultiSet.FromArray(_773_v45)).IsProperSubsetOf((_dafny.MultiSet.fromElements(new BigNumber((_774_v46).length), new BigNumber((_776_v48).length), _708_v2, new BigNumber(596), new BigNumber(650))).Union(_777_v49))));
        let _778_v50;
        _778_v50 = _dafny.Map.Empty.slice().updateUnsafe(_774_v46,(_707_v1).f17);
        _778_v50 = (_778_v50).update(_774_v46, (_707_v1).f17);
        let _779_v51;
        _779_v51 = _dafny.Set.fromElements(_706_v0);
        let _780_v52;
        _780_v52 = _module.D4.create_DC10(_779_v51);
        _780_v52 = _780_v52;
        _706_v0 = !(_708_v2).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-411)), ((_781_v43) => function (_782_i6) {
          return _781_v43;
        })(_771_v43))).length));
      } else {
        let _783_v53;
        _783_v53 = new _dafny.CodePoint('m'.codePointAt(0));
        let _784_v54;
        _784_v54 = _dafny.Seq.UnicodeFromString("ehbrla");
        if (!_dafny.Seq.contains(_784_v54, _783_v53)) {
          (globalState).f9 = _708_v2;
          (globalState).f9 = _708_v2;
          _706_v0 = (_707_v1).f17;
          let _785_v55;
          let _nw135 = Array((new BigNumber(23)).toNumber()).fill(false);
          _785_v55 = _nw135;
          let _786_v56;
          _786_v56 = _dafny.Seq.of(_785_v55);
          let _787_v57;
          _787_v57 = _dafny.MultiSet.fromElements(_783_v53, _783_v53, new _dafny.CodePoint('i'.codePointAt(0)), (_784_v54)[_module.__default.safeIndex(_708_v2, new BigNumber((_784_v54).length))]);
          _786_v56 = _dafny.Seq.update(_786_v56, _module.__default.safeIndex((_708_v2).multipliedBy((((_787_v57).contains(_783_v53)) ? ((_787_v57).get(_783_v53)) : (_708_v2))), new BigNumber((_786_v56).length)), _785_v55);
          let _788_v58;
          let _init18 = ((_789_v54) => function (_790_i7) {
            return _module.D2.create_DC5(_789_v54);
          })(_784_v54);
          let _nw136 = Array((new BigNumber(21)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw136.length); _i0_18++) {
            _nw136[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _788_v58 = _nw136;
          _788_v58 = _788_v58;
        } else {
          let _791_v59;
          _791_v59 = _dafny.MultiSet.fromElements(_708_v2);
          _791_v59 = _791_v59;
          let _792_v60;
          let _nw137 = Array((new BigNumber(26)).toNumber()).fill([]);
          _792_v60 = _nw137;
          let _793_v61;
          let _init19 = ((_794_v2) => function (_795_i8) {
            return (_795_i8).multipliedBy(_794_v2);
          })(_708_v2);
          let _nw138 = Array((new BigNumber(14)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw138.length); _i0_19++) {
            _nw138[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _793_v61 = _nw138;
          let _index113 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_792_v60).length));
          (_792_v60)[_index113] = _793_v61;
          let _796_v62;
          let _nw139 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
          _796_v62 = _nw139;
          let _797_v63;
          _797_v63 = _dafny.Map.Empty.slice().updateUnsafe(_796_v62,_706_v0);
          let _798_v64;
          _798_v64 = _dafny.Seq.of(_793_v61);
          let _799_v65;
          _799_v65 = _dafny.Map.Empty.slice().updateUnsafe(_793_v61,_709_v3);
          let _800_v66;
          _800_v66 = _dafny.Set.fromElements((_707_v1).f17);
          let _801_v67;
          _801_v67 = _dafny.Set.fromElements(_800_v66, _800_v66, _800_v66);
          let _index114 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_792_v60).length));
          let _nw140 = Array((new BigNumber(11)).toNumber());
          _nw140[(_dafny.ZERO).toNumber()] = _708_v2;
          _nw140[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_708_v2);
          _nw140[(new BigNumber(2)).toNumber()] = new BigNumber(-427);
          _nw140[(new BigNumber(3)).toNumber()] = _module.__default.fm0((((_797_v63).contains(_796_v62)) ? ((_797_v63).get(_796_v62)) : ((_707_v1).f17)), globalState);
          _nw140[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_709_v3).length), _708_v2);
          _nw140[(new BigNumber(5)).toNumber()] = new BigNumber(-64);
          _nw140[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_708_v2).multipliedBy(_708_v2));
          _nw140[(new BigNumber(7)).toNumber()] = _708_v2;
          _nw140[(new BigNumber(8)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_798_v64)[_module.__default.safeIndex(_708_v2, new BigNumber((_798_v64).length))],_709_v3)).Merge(((_799_v65).update(_793_v61, _dafny.Seq.of((_707_v1).f17))).update(_793_v61, _dafny.Seq.update(_709_v3, _module.__default.safeIndex(_708_v2, new BigNumber((_709_v3).length)), (_707_v1).f17)))).length);
          _nw140[(new BigNumber(9)).toNumber()] = new BigNumber(((_801_v67).Intersect(_801_v67)).length);
          _nw140[(new BigNumber(10)).toNumber()] = _708_v2;
          (_792_v60)[_index114] = _nw140;
          (globalState).f8 = !((_707_v1).f17);
          let _802_v68;
          _802_v68 = _module.D5.create_DC14((_707_v1).f17);
          let _803_v69;
          _803_v69 = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_802_v68).dtor_cf26);
          let _804_v70;
          _804_v70 = _dafny.Map.Empty.slice().updateUnsafe(_706_v0,_dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_707_v1).f17));
          let _805_v71;
          _805_v71 = _dafny.Seq.of(_803_v69, _803_v69);
          let _806_v75;
          _806_v75 = _dafny.MultiSet.fromElements(_783_v53, _783_v53);
          let _807_v77;
          _807_v77 = _module.D4.create_DC12((_707_v1).f17, (_707_v1).f17, (_707_v1).f17);
          let _808_v78;
          let _nw141 = Array((new BigNumber(27)).toNumber());
          _nw141[(_dafny.ZERO).toNumber()] = _803_v69;
          _nw141[(_dafny.ONE).toNumber()] = (((_804_v70).contains(_706_v0)) ? ((_804_v70).get(_706_v0)) : (_module.__default.fm22((_707_v1).f17, (_707_v1).f17, _783_v53, _module.__default.fm3(_708_v2, (_707_v1).f17, globalState), globalState)));
          _nw141[(new BigNumber(2)).toNumber()] = (((_707_v1).f17) ? (_803_v69) : (_803_v69));
          _nw141[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_707_v1).f17);
          _nw141[(new BigNumber(4)).toNumber()] = (_803_v69).Merge((_805_v71)[_module.__default.safeIndex(_708_v2, new BigNumber((_805_v71).length))]);
          _nw141[(new BigNumber(5)).toNumber()] = function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_dafny.Seq.of(_783_v53)).Elements) {
              let _809_v72 = _compr_37;
              if (_dafny.Seq.contains(_dafny.Seq.of(_783_v53), _809_v72)) {
                _coll37.push([_809_v72,true]);
              }
            }
            return _coll37;
          }();
          _nw141[(new BigNumber(6)).toNumber()] = (((_707_v1).f17) ? (_dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_707_v1).f17)) : (_803_v69));
          _nw141[(new BigNumber(7)).toNumber()] = ((_706_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_707_v1).f17)) : (_dafny.Map.Empty.slice().updateUnsafe((_784_v54)[_module.__default.safeIndex(_708_v2, new BigNumber((_784_v54).length))],(_707_v1).f17)));
          _nw141[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_707_v1).f17);
          _nw141[(new BigNumber(9)).toNumber()] = function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-618)), function (_810_i9) {
              return new _dafny.CodePoint('b'.codePointAt(0));
            })).Elements) {
              let _811_v73 = _compr_38;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-618)), function (_810_i9) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              }), _811_v73)) {
                _coll38.push([_811_v73,(_707_v1).f17]);
              }
            }
            return _coll38;
          }();
          _nw141[(new BigNumber(10)).toNumber()] = (_803_v69).Merge(_803_v69);
          _nw141[(new BigNumber(11)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_783_v53,!((_707_v1).f17))).Merge(_803_v69);
          _nw141[(new BigNumber(12)).toNumber()] = (_803_v69).update(_783_v53, (_707_v1).f17);
          _nw141[(new BigNumber(13)).toNumber()] = (_803_v69).update(_783_v53, !((_707_v1).f17));
          _nw141[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_707_v1).f17);
          _nw141[(new BigNumber(15)).toNumber()] = function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_806_v75).Elements) {
              let _812_v74 = _compr_39;
              if ((_806_v75).contains(_812_v74)) {
                _coll39.push([_812_v74,(_707_v1).f17]);
              }
            }
            return _coll39;
          }();
          _nw141[(new BigNumber(16)).toNumber()] = _803_v69;
          _nw141[(new BigNumber(17)).toNumber()] = function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of (_803_v69).Keys.Elements) {
              let _813_v76 = _compr_40;
              if ((_803_v69).contains(_813_v76)) {
                _coll40.push([_813_v76,_706_v0]);
              }
            }
            return _coll40;
          }();
          _nw141[(new BigNumber(18)).toNumber()] = _803_v69;
          _nw141[(new BigNumber(19)).toNumber()] = (_803_v69).Merge(_803_v69);
          _nw141[(new BigNumber(20)).toNumber()] = _803_v69;
          _nw141[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,(_807_v77).dtor_cf23);
          _nw141[(new BigNumber(22)).toNumber()] = _803_v69;
          _nw141[(new BigNumber(23)).toNumber()] = _803_v69;
          _nw141[(new BigNumber(24)).toNumber()] = _803_v69;
          _nw141[(new BigNumber(25)).toNumber()] = _module.__default.fm22(_706_v0, (_707_v1).f17, _783_v53, _783_v53, globalState);
          _nw141[(new BigNumber(26)).toNumber()] = _803_v69;
          _808_v78 = _nw141;
          _808_v78 = _808_v78;
          let _814_v79;
          _814_v79 = _dafny.Seq.of(_709_v3);
          _814_v79 = _dafny.Seq.of(_709_v3);
        }
        let _815_v80;
        let _nw142 = new _module.C0();
        _nw142.__ctor(_706_v0);
        _815_v80 = _nw142;
        let _816_v81;
        _816_v81 = _module.D5.create_DC14(!(_706_v0));
        let _source7 = (((_815_v80).f17) ? (_816_v81) : (_816_v81));
        if (_source7.is_DC14) {
          let _817___mcc_h7 = (_source7).cf26;
          let _818_cf26 = _817___mcc_h7;
          let _819_v82;
          _819_v82 = _dafny.Seq.of(_784_v54);
          let _820_v83;
          _820_v83 = _dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,_708_v2);
          let _821_v84;
          let _nw143 = Array((new BigNumber(26)).toNumber());
          _nw143[(_dafny.ZERO).toNumber()] = (_707_v1).f17;
          _nw143[(_dafny.ONE).toNumber()] = !(_818_cf26);
          _nw143[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_818_cf26, _783_v53, globalState);
          _nw143[(new BigNumber(3)).toNumber()] = (new BigNumber((_819_v82).length)).isLessThan(_708_v2);
          _nw143[(new BigNumber(4)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(5)).toNumber()] = false;
          _nw143[(new BigNumber(6)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(7)).toNumber()] = !(true);
          _nw143[(new BigNumber(8)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(9)).toNumber()] = (_708_v2).isLessThan(_708_v2);
          _nw143[(new BigNumber(10)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(11)).toNumber()] = _818_cf26;
          _nw143[(new BigNumber(12)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(13)).toNumber()] = !((_707_v1).f17) || (false);
          _nw143[(new BigNumber(14)).toNumber()] = ((_815_v80).f17) || ((_707_v1).f17);
          _nw143[(new BigNumber(15)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(16)).toNumber()] = _dafny.Seq.IsPrefixOf(_709_v3, _709_v3);
          _nw143[(new BigNumber(17)).toNumber()] = !((_708_v2).isLessThanOrEqualTo(_708_v2));
          _nw143[(new BigNumber(18)).toNumber()] = !(_820_v83).equals(_820_v83);
          _nw143[(new BigNumber(19)).toNumber()] = _module.__default.fm1((_815_v80).f17, _783_v53, globalState);
          _nw143[(new BigNumber(20)).toNumber()] = !((((_815_v80).f17) ? (!(_706_v0)) : ((_707_v1).f17)));
          _nw143[(new BigNumber(21)).toNumber()] = (_707_v1).f17;
          _nw143[(new BigNumber(22)).toNumber()] = _module.__default.fm1((_707_v1).f17, _783_v53, globalState);
          _nw143[(new BigNumber(23)).toNumber()] = !((_815_v80).f17);
          _nw143[(new BigNumber(24)).toNumber()] = ((_706_v0) ? ((_707_v1).f17) : ((_707_v1).f17));
          _nw143[(new BigNumber(25)).toNumber()] = (_818_cf26) && ((_815_v80).f17);
          _821_v84 = _nw143;
          let _index115 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_821_v84).length));
          (_821_v84)[_index115] = _818_cf26;
          let _index116 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_821_v84).length));
          (_821_v84)[_index116] = _706_v0;
          let _index117 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_821_v84).length));
          (_821_v84)[_index117] = _818_cf26;
          let _822_v85;
          _822_v85 = _dafny.Map.Empty.slice().updateUnsafe(_708_v2,(_815_v80).f17);
          let _823_v86;
          _823_v86 = _dafny.MultiSet.fromElements((_707_v1).f17);
          let _824_v87;
          _824_v87 = _module.D0.create_DC2(_823_v86, (_815_v80).f17, _708_v2, (_821_v84)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_821_v84).length))]);
          _822_v85 = (_822_v85).update(_708_v2, (_824_v87).dtor_cf7);
          let _825_v88;
          let _nw144 = new _module.C0();
          _nw144.__ctor((((_815_v80).f17) ? ((_821_v84)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_821_v84).length))]) : ((_821_v84)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_821_v84).length))])));
          _825_v88 = _nw144;
        } else {
          let _826___mcc_h8 = (_source7).cf25;
          let _827_cf25 = _826___mcc_h8;
          let _828_v89;
          let _init20 = ((_829_v53, _830_v2) => function (_831_i10) {
            return _dafny.Map.Empty.slice().updateUnsafe(_829_v53,_830_v2);
          })(_783_v53, _708_v2);
          let _nw145 = Array((new BigNumber(19)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw145.length); _i0_20++) {
            _nw145[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _828_v89 = _nw145;
          let _832_v90;
          _832_v90 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_707_v1,true),_783_v53);
          let _833_v91;
          _833_v91 = _dafny.Map.Empty.slice().updateUnsafe(_707_v1,_706_v0);
          let _834_v92;
          _834_v92 = _dafny.Set.fromElements(_783_v53, (((_832_v90).contains(_833_v91)) ? ((_832_v90).get(_833_v91)) : (_783_v53)));
          let _835_v93;
          _835_v93 = _834_v92;
          let _rhs126 = _708_v2;
          let _rhs127 = _828_v89;
          let _rhs128 = (_835_v93);
          let _rhs129 = (_707_v1).f17;
          let _lhs88 = globalState;
          _lhs88.f9 = _rhs126;
          _828_v89 = _rhs127;
          _834_v92 = _rhs128;
          _706_v0 = _rhs129;
          let _836_v94;
          _836_v94 = _dafny.Set.fromElements((_707_v1).f17);
          let _837_v95;
          _837_v95 = _module.D4.create_DC10(_836_v94);
          let _838_v96;
          let _out9;
          _out9 = (_this).m8(_837_v95, _708_v2, globalState);
          _838_v96 = _out9;
          let _839_v97;
          _839_v97 = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,new BigNumber(437));
          let _rhs130 = false;
          let _rhs131 = _839_v97;
          let _lhs89 = globalState;
          _lhs89.f8 = _rhs130;
          _839_v97 = _rhs131;
          let _840_v98;
          _840_v98 = _dafny.Set.fromElements(_module.__default.fm23(_706_v0, globalState));
          let _841_v99;
          _841_v99 = _dafny.Set.fromElements(_708_v2);
          _840_v98 = _dafny.Set.fromElements(_841_v99, _dafny.Set.fromElements(new BigNumber(385), new BigNumber((_module.__default.fm24(_784_v54, (_707_v1).f17, (_707_v1).f17, globalState)).length)));
        }
        let _source8 = _module.D2.create_DC6();
        if (_source8.is_DC6) {
          (globalState).f8 = !((_708_v2).isLessThanOrEqualTo(_708_v2));
          (globalState).f8 = _706_v0;
          let _842_v100;
          let _init21 = ((_843_v53) => function (_844_i11) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-151)), ((_845_v53) => function (_846_i12) {
              return _845_v53;
            })(_843_v53));
          })(_783_v53);
          let _nw146 = Array((_dafny.ONE).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw146.length); _i0_21++) {
            _nw146[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _842_v100 = _nw146;
          _842_v100 = _842_v100;
          _784_v54 = _dafny.Seq.update(_784_v54, _module.__default.safeIndex((_dafny.ZERO).minus(_708_v2), new BigNumber((_784_v54).length)), new _dafny.CodePoint('a'.codePointAt(0)));
        } else if (_source8.is_DC7) {
          let _847___mcc_h9 = (_source8).cf11;
          let _848___mcc_h10 = (_source8).cf12;
          let _849___mcc_h11 = (_source8).cf13;
          let _850___mcc_h12 = (_source8).cf14;
          let _851_cf14 = _850___mcc_h12;
          let _852_cf13 = _849___mcc_h11;
          let _853_cf12 = _848___mcc_h10;
          let _854_cf11 = _847___mcc_h9;
          let _855_v101;
          _855_v101 = _module.D2.create_DC8(_853_cf12, new _dafny.CodePoint('t'.codePointAt(0)));
          let _856_v102;
          _856_v102 = _dafny.Seq.of(_853_cf12);
          let _857_v103;
          _857_v103 = _dafny.Set.fromElements(_708_v2, _853_cf12, (_855_v101).dtor_cf15, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_853_cf12), _856_v102)).length)));
          let _858_v104;
          _858_v104 = _dafny.Map.Empty.slice().updateUnsafe(_706_v0,_851_cf14);
          let _859_v105;
          _859_v105 = _dafny.MultiSet.fromElements((((_858_v104).contains(!(true))) ? ((_858_v104).get(!(true))) : (_851_cf14)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(314)), ((_860_v53) => function (_861_i13) {
            return _860_v53;
          })(_783_v53)), _784_v54, _dafny.Seq.update(_851_cf14, _module.__default.safeIndex(_708_v2, new BigNumber((_851_cf14).length)), _783_v53));
          let _rhs132 = (((_859_v105).IsSubsetOf(_dafny.MultiSet.fromElements(_851_cf14, _851_cf14, _784_v54, _851_cf14, _784_v54))) ? (_852_cf13) : ((_707_v1).f17));
          let _rhs133 = (_module.__default.fm23(_706_v0, globalState)).Union(function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of (_dafny.MultiSet.FromArray(_856_v102)).Elements) {
              let _862_v106 = _compr_41;
              if ((_dafny.MultiSet.FromArray(_856_v102)).contains(_862_v106)) {
                _coll41.add((_862_v106).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-93)), function (_863_i14) {
                  return new _dafny.CodePoint('x'.codePointAt(0));
                })).length)));
              }
            }
            return _coll41;
          }());
          let _rhs134 = _853_cf12;
          _854_cf11 = _rhs132;
          _857_v103 = _rhs133;
          _853_cf12 = _rhs134;
          let _864_v107;
          let _nw147 = Array((new BigNumber(21)).toNumber());
          _864_v107 = _nw147;
          let _index118 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_864_v107).length));
          (_864_v107)[_index118] = _707_v1;
          let _index119 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_864_v107).length));
          (_864_v107)[_index119] = _815_v80;
          let _865_v108;
          let _nw148 = Array((new BigNumber(9)).toNumber()).fill(false);
          _865_v108 = _nw148;
          _865_v108 = (((_815_v80).f17) ? (_865_v108) : (_865_v108));
          let _866_v109;
          let _nw149 = Array((new BigNumber(7)).toNumber());
          _nw149[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_853_cf12);
          _nw149[(_dafny.ONE).toNumber()] = _708_v2;
          _nw149[(new BigNumber(2)).toNumber()] = _708_v2;
          _nw149[(new BigNumber(3)).toNumber()] = _853_cf12;
          _nw149[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_853_cf12), _853_cf12);
          _nw149[(new BigNumber(5)).toNumber()] = _708_v2;
          _nw149[(new BigNumber(6)).toNumber()] = (_this).fm20((_707_v1).f17, _853_cf12, globalState);
          _866_v109 = _nw149;
          let _867_v110;
          _867_v110 = _dafny.MultiSet.fromElements((_707_v1).f17);
          let _index120 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_866_v109).length));
          (_866_v109)[_index120] = new BigNumber(((_867_v110).update(_852_cf13, _module.__default.abs(_853_cf12))).cardinality());
          let _index121 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_866_v109).length));
          (_866_v109)[_index121] = (_dafny.ZERO).minus(_708_v2);
        } else if (_source8.is_DC8) {
          let _868___mcc_h13 = (_source8).cf15;
          let _869___mcc_h14 = (_source8).cf16;
          let _870_cf16 = _869___mcc_h14;
          let _871_cf15 = _868___mcc_h13;
          let _872_v111;
          let _nw150 = Array((new BigNumber(29)).toNumber()).fill(false);
          _872_v111 = _nw150;
          _872_v111 = _872_v111;
          let _873_v112;
          _873_v112 = _dafny.Seq.of(new BigNumber(363));
          let _874_v113;
          _874_v113 = _dafny.Map.Empty.slice().updateUnsafe(_784_v54,new BigNumber((_784_v54).length));
          (globalState).f8 = ((_871_cf15).plus((_dafny.ZERO).minus(_871_cf15))).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(_873_v112, _module.__default.safeIndex(_708_v2, new BigNumber((_873_v112).length)), new BigNumber(-597))).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_873_v112, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,(_dafny.ZERO).minus(_871_cf15)), _module.__default.fm25(_708_v2, globalState))).length), new BigNumber((_873_v112).length)), (((_874_v113).contains(_784_v54)) ? ((_874_v113).get(_784_v54)) : (_708_v2))), _module.__default.safeIndex(_871_cf15, new BigNumber((_dafny.Seq.update(_873_v112, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,(_dafny.ZERO).minus(_871_cf15)), _module.__default.fm25(_708_v2, globalState))).length), new BigNumber((_873_v112).length)), (((_874_v113).contains(_784_v54)) ? ((_874_v113).get(_784_v54)) : (_708_v2)))).length)), _708_v2)).length)));
          let _875_v114;
          let _nw151 = new _module.C0();
          _nw151.__ctor((_707_v1).f17);
          _875_v114 = _nw151;
          let _876_v115;
          let _nw152 = new _module.C0();
          _nw152.__ctor((_707_v1).f17);
          _876_v115 = _nw152;
        } else {
          let _877___mcc_h15 = (_source8).cf10;
          let _878_cf10 = _877___mcc_h15;
          let _879_v116;
          let _nw153 = Array((new BigNumber(15)).toNumber());
          _nw153[(_dafny.ZERO).toNumber()] = _784_v54;
          _nw153[(_dafny.ONE).toNumber()] = _784_v54;
          _nw153[(new BigNumber(2)).toNumber()] = _784_v54;
          _nw153[(new BigNumber(3)).toNumber()] = _878_cf10;
          _nw153[(new BigNumber(4)).toNumber()] = _878_cf10;
          _nw153[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_784_v54, _module.__default.safeIndex(_708_v2, new BigNumber((_784_v54).length)), _783_v53);
          _nw153[(new BigNumber(6)).toNumber()] = _878_cf10;
          _nw153[(new BigNumber(7)).toNumber()] = _878_cf10;
          _nw153[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("triwllfx");
          _nw153[(new BigNumber(9)).toNumber()] = _784_v54;
          _nw153[(new BigNumber(10)).toNumber()] = _878_cf10;
          _nw153[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("pnlwkn");
          _nw153[(new BigNumber(12)).toNumber()] = _784_v54;
          _nw153[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_784_v54, _878_cf10);
          _nw153[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("v");
          _879_v116 = _nw153;
          let _index122 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_879_v116).length));
          (_879_v116)[_index122] = _878_cf10;
          let _index123 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_879_v116).length));
          (_879_v116)[_index123] = _878_cf10;
          let _880_v117;
          let _nw154 = Array((new BigNumber(11)).toNumber()).fill([]);
          _880_v117 = _nw154;
          let _881_v118;
          let _nw155 = Array((new BigNumber(6)).toNumber());
          _nw155[(_dafny.ZERO).toNumber()] = (_707_v1).f17;
          _nw155[(_dafny.ONE).toNumber()] = (_815_v80).f17;
          _nw155[(new BigNumber(2)).toNumber()] = (_707_v1).f17;
          _nw155[(new BigNumber(3)).toNumber()] = (_707_v1).f17;
          _nw155[(new BigNumber(4)).toNumber()] = (_707_v1).f17;
          _nw155[(new BigNumber(5)).toNumber()] = (_815_v80).f17;
          _881_v118 = _nw155;
          let _index124 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_880_v117).length));
          (_880_v117)[_index124] = _881_v118;
          let _index125 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_880_v117).length));
          (_880_v117)[_index125] = _881_v118;
          (globalState).f8 = _module.__default.fm1((_707_v1).f17, _783_v53, globalState);
          let _882_v119;
          _882_v119 = _dafny.Seq.of((_879_v116)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_879_v116).length))], _784_v54);
          _706_v0 = _dafny.Seq.IsProperPrefixOf((_882_v119)[_module.__default.safeIndex(_708_v2, new BigNumber((_882_v119).length))], (_879_v116)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_879_v116).length))]);
        }
        let _883_v120;
        _883_v120 = _module.D2.create_DC6();
        let _884_v121;
        _884_v121 = _dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,_883_v120);
        _884_v121 = (_884_v121).update((_707_v1).f17, _883_v120);
      }
      if (_706_v0) {
        let _885_v122;
        let _init22 = ((_886_v1) => function (_887_i15) {
          return !((_886_v1).f17);
        })(_707_v1);
        let _nw156 = Array((new BigNumber(16)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw156.length); _i0_22++) {
          _nw156[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _885_v122 = _nw156;
        let _888_v123;
        let _nw157 = Array((new BigNumber(3)).toNumber());
        _nw157[(_dafny.ZERO).toNumber()] = _885_v122;
        _nw157[(_dafny.ONE).toNumber()] = _885_v122;
        _nw157[(new BigNumber(2)).toNumber()] = _885_v122;
        _888_v123 = _nw157;
        let _index126 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_888_v123).length));
        (_888_v123)[_index126] = _885_v122;
        let _index127 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_888_v123).length));
        (_888_v123)[_index127] = _885_v122;
        let _889_v124;
        _889_v124 = _dafny.MultiSet.fromElements(false);
        let _890_v125;
        _890_v125 = _dafny.Seq.of(new BigNumber((_889_v124).cardinality()));
        _890_v125 = _890_v125;
        _706_v0 = (_707_v1).f17;
        let _891_v126;
        _891_v126 = _dafny.Seq.UnicodeFromString("bldcpsr");
        let _892_v127;
        _892_v127 = _dafny.Seq.of(_891_v126, _891_v126, _dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_893_i16) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }));
        let _894_v128;
        _894_v128 = _dafny.Map.Empty.slice().updateUnsafe(((_890_v125)[_module.__default.safeIndex(_708_v2, new BigNumber((_890_v125).length))]).plus(new BigNumber((_892_v127).length)),(((_707_v1).f17) ? (!(_706_v0)) : (_706_v0)));
        let _895_v129;
        _895_v129 = _dafny.Set.fromElements(_module.D4.create_DC11(_708_v2, _708_v2, true));
        let _rhs135 = (_894_v128).Merge(_894_v128);
        let _rhs136 = ((_707_v1).f17) === ((_707_v1).f17);
        let _rhs137 = (_895_v129).IsProperSubsetOf(_895_v129);
        let _lhs90 = globalState;
        _894_v128 = _rhs135;
        _lhs90.f8 = _rhs136;
        _706_v0 = _rhs137;
        let _896_v130;
        _896_v130 = new _dafny.CodePoint('l'.codePointAt(0));
        (globalState).f9 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_708_v2), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_707_v1).f17, _896_v130, globalState),_708_v2)).length));
      } else {
        if (((((_707_v1).f17) ? (_706_v0) : (_706_v0))) === (((!(false)) ? ((_707_v1).f17) : ((_707_v1).f17)))) {
          _708_v2 = _708_v2;
          let _897_v131;
          _897_v131 = _dafny.Seq.UnicodeFromString("ccokdbo");
          let _898_v132;
          _898_v132 = _dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,_708_v2);
          let _899_v133;
          _899_v133 = _dafny.Seq.of(_708_v2, _708_v2, new BigNumber((_898_v132).length));
          let _900_v134;
          _900_v134 = new _dafny.CodePoint('w'.codePointAt(0));
          let _rhs138 = (!(_706_v0) || ((_707_v1).f17)) === (!_dafny.areEqual(_899_v133, _dafny.Seq.update(_899_v133, _module.__default.safeIndex(new BigNumber((_module.__default.fm26(!((_707_v1).f17), new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), function (_901_i17) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _module.__default.safeIndex(_708_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), function (_902_i17) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })).length)), _900_v134)).length), _706_v0, globalState)).length), new BigNumber((_899_v133).length)), _708_v2)));
          let _rhs139 = _708_v2;
          let _rhs140 = _897_v131;
          let _lhs91 = globalState;
          r2 = _rhs138;
          _lhs91.f9 = _rhs139;
          _897_v131 = _rhs140;
          (globalState).f9 = (_708_v2).plus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_708_v2, _708_v2)));
          let _903_v135;
          let _nw158 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _903_v135 = _nw158;
          let _index128 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_903_v135).length));
          (_903_v135)[_index128] = _708_v2;
          let _index129 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_903_v135).length));
          (_903_v135)[_index129] = new BigNumber((_dafny.Set.fromElements((_707_v1).f17, _706_v0, !(!((new BigNumber((_897_v131).length)).isLessThan(_708_v2))))).length);
          let _904_v136;
          _904_v136 = _dafny.Map.Empty.slice().updateUnsafe((_903_v135)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_903_v135).length))],_706_v0);
          let _905_v137;
          let _nw159 = new _module.C0();
          _nw159.__ctor((((_904_v136).contains((_903_v135)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_903_v135).length))])) ? ((_904_v136).get((_903_v135)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_903_v135).length))])) : (true)));
          _905_v137 = _nw159;
        } else {
          let _906_v138;
          _906_v138 = _dafny.Seq.UnicodeFromString("wqtu");
          let _907_v139;
          _907_v139 = _dafny.MultiSet.fromElements(_708_v2);
          let _908_v140;
          _908_v140 = _dafny.Map.Empty.slice().updateUnsafe(_708_v2,new BigNumber((_module.__default.fm27(_dafny.Seq.of(_709_v3), globalState)).length));
          _906_v138 = _dafny.Seq.Concat(_module.__default.fm2(_708_v2, _907_v139, _908_v140, globalState), _906_v138);
          let _909_v141;
          _909_v141 = new _dafny.CodePoint('u'.codePointAt(0));
          _906_v138 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_906_v138, _906_v138), _dafny.Seq.Concat(_906_v138, _dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), function (_910_i18) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }))), _module.__default.safeIndex(_module.__default.safeModuloInt(_708_v2, _708_v2), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_906_v138, _906_v138), _dafny.Seq.Concat(_906_v138, _dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), function (_911_i18) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })))).length)), _909_v141);
          let _912_v142;
          let _init23 = ((_913_v1) => function (_914_i19) {
            return (_913_v1).f17;
          })(_707_v1);
          let _nw160 = Array((new BigNumber(27)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw160.length); _i0_23++) {
            _nw160[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _912_v142 = _nw160;
          let _index130 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length));
          (_912_v142)[_index130] = (_707_v1).f17;
          let _915_v143;
          let _nw161 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _915_v143 = _nw161;
          let _index131 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_915_v143).length));
          (_915_v143)[_index131] = (((_707_v1).f17) ? (_708_v2) : (_708_v2));
          let _916_v144;
          _916_v144 = _dafny.Map.Empty.slice().updateUnsafe(_708_v2,_709_v3);
          let _index132 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length));
          let _index133 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_915_v143).length));
          let _rhs141 = _dafny.Seq.IsPrefixOf(_709_v3, (((_916_v144).contains(_708_v2)) ? ((_916_v144).get(_708_v2)) : (_709_v3)));
          let _rhs142 = (_707_v1).f17;
          let _rhs143 = _dafny.Seq.Concat(_906_v138, _906_v138);
          let _rhs144 = _708_v2;
          let _rhs145 = _912_v142;
          let _lhs92 = globalState;
          let _lhs93 = _912_v142;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length));
          let _lhs95 = _915_v143;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_915_v143).length));
          _lhs92.f8 = _rhs141;
          _lhs93[_lhs94] = _rhs142;
          _906_v138 = _rhs143;
          _lhs95[_lhs96] = _rhs144;
          _912_v142 = _rhs145;
          let _917_v145;
          _917_v145 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), ((_918_v141) => function (_919_i20) {
            return _918_v141;
          })(_909_v141)), _906_v138);
          let _920_v146;
          _920_v146 = _dafny.MultiSet.fromElements((_912_v142)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length))]);
          let _921_v147;
          _921_v147 = _module.D4.create_DC12(_dafny.Seq.IsProperPrefixOf(_917_v145, _module.__default.fm28(new _dafny.CodePoint('o'.codePointAt(0)), _708_v2, new BigNumber(-300), _706_v0, globalState)), (_920_v146).IsSubsetOf(_920_v146), (_912_v142)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length))]);
          let _922_v148;
          _922_v148 = _dafny.Seq.of(_708_v2);
          let _index134 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length));
          let _index135 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_915_v143).length));
          let _rhs146 = _module.D4.create_DC12(_706_v0, !(_module.__default.fm1(!((_912_v142)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length))]), _909_v141, globalState)), (_707_v1).f17);
          let _rhs147 = (_707_v1).f17;
          let _rhs148 = (_922_v148)[_module.__default.safeIndex(_708_v2, new BigNumber((_922_v148).length))];
          let _rhs149 = (((_915_v143)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_915_v143).length))]).minus(new BigNumber(44))).plus(_708_v2);
          let _lhs97 = _912_v142;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length));
          let _lhs99 = globalState;
          let _lhs100 = _915_v143;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_915_v143).length));
          _921_v147 = _rhs146;
          _lhs97[_lhs98] = _rhs147;
          _lhs99.f9 = _rhs148;
          _lhs100[_lhs101] = _rhs149;
          let _923_v149;
          _923_v149 = _module.D1.create_DC4(_909_v141);
          let _924_v150;
          _924_v150 = _dafny.Map.Empty.slice().updateUnsafe(_923_v149,_708_v2);
          let _925_v151;
          _925_v151 = _dafny.Map.Empty.slice().updateUnsafe((_707_v1).f17,(_924_v150).Merge(_924_v150));
          _925_v151 = (_925_v151).update((_912_v142)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_912_v142).length))], (_924_v150).Merge(_924_v150));
        }
        let _926_v152;
        _926_v152 = _module.D7.create_DC16(_709_v3);
        let _927_v153;
        _927_v153 = _dafny.Set.fromElements(_706_v0);
        let _928_v154;
        _928_v154 = _dafny.Map.Empty.slice().updateUnsafe((_709_v3)[_module.__default.safeIndex(new BigNumber((_927_v153).length), new BigNumber((_709_v3).length))],new BigNumber(-789));
        (globalState).f8 = !(_dafny.Seq.contains(_dafny.Seq.update((_926_v152).dtor_cf28, _module.__default.safeIndex((((_928_v154).contains(!((_707_v1).f17))) ? ((_928_v154).get(!((_707_v1).f17))) : (_708_v2)), new BigNumber(((_926_v152).dtor_cf28).length)), _706_v0), !((_707_v1).f17))) || ((_708_v2).isEqualTo(_708_v2));
        if (_706_v0) {
          _927_v153 = _927_v153;
          let _929_v155;
          _929_v155 = _dafny.Seq.UnicodeFromString("qaxdpco");
          let _930_v156;
          _930_v156 = _module.D2.create_DC5(_929_v155);
          let _931_v157;
          _931_v157 = _dafny.Map.Empty.slice().updateUnsafe(_929_v155,_930_v156);
          let _pat_let_tv11 = _929_v155;
          _931_v157 = (_931_v157).update(_929_v155, function (_pat_let11_0) {
            return function (_932_dt__update__tmp_h0) {
              return function (_pat_let12_0) {
                return function (_933_dt__update_hcf10_h0) {
                  return _module.D2.create_DC5(_933_dt__update_hcf10_h0);
                }(_pat_let12_0);
              }(_pat_let_tv11);
            }(_pat_let11_0);
          }(_930_v156));
          (globalState).f9 = ((_dafny.ZERO).minus(_708_v2)).minus(_708_v2);
          let _934_v158;
          _934_v158 = new _dafny.CodePoint('y'.codePointAt(0));
          r0 = _934_v158;
          let _935_v159;
          let _nw162 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _935_v159 = _nw162;
          let _index136 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_935_v159).length));
          (_935_v159)[_index136] = new BigNumber((_dafny.Seq.Concat(_709_v3, _709_v3)).length);
          let _936_v161;
          _936_v161 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of (_929_v155).Elements) {
              let _937_v160 = _compr_42;
              if (_dafny.Seq.contains(_929_v155, _937_v160)) {
                _coll42.add(_937_v160);
              }
            }
            return _coll42;
          }(),_707_v1);
          let _938_v162;
          _938_v162 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(932)), ((_939_v158) => function (_940_i21) {
            return _939_v158;
          })(_934_v158))).length),_706_v0);
          let _index137 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_935_v159).length));
          let _rhs150 = (((_928_v154).contains((((_938_v162).contains(new BigNumber((_dafny.MultiSet.fromElements(_708_v2)).cardinality()))) ? ((_938_v162).get(new BigNumber((_dafny.MultiSet.fromElements(_708_v2)).cardinality()))) : ((_707_v1).f17)))) ? ((_928_v154).get((((_938_v162).contains(new BigNumber((_dafny.MultiSet.fromElements(_708_v2)).cardinality()))) ? ((_938_v162).get(new BigNumber((_dafny.MultiSet.fromElements(_708_v2)).cardinality()))) : ((_707_v1).f17)))) : ((_this).fm20((_707_v1).f17, _708_v2, globalState)));
          let _rhs151 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(791)), function (_941_i22) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length);
          let _rhs152 = _936_v161;
          let _rhs153 = _929_v155;
          let _rhs154 = !(_706_v0);
          let _lhs102 = globalState;
          let _lhs103 = _935_v159;
          let _lhs104 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_935_v159).length));
          let _lhs105 = globalState;
          _lhs102.f9 = _rhs150;
          _lhs103[_lhs104] = _rhs151;
          _936_v161 = _rhs152;
          _929_v155 = _rhs153;
          _lhs105.f8 = _rhs154;
        } else {
          (globalState).f8 = (_707_v1).f17;
          let _942_v163;
          let _nw163 = new _module.C0();
          _nw163.__ctor((_707_v1).f17);
          _942_v163 = _nw163;
          let _943_v164;
          _943_v164 = _dafny.Seq.of(_708_v2);
          (globalState).f9 = new BigNumber((_943_v164).length);
          let _rhs155 = _708_v2;
          let _rhs156 = (((_707_v1).f17) && ((_707_v1).f17)) && (((true) ? ((_707_v1).f17) : ((_707_v1).f17)));
          _708_v2 = _rhs155;
          r1 = _rhs156;
          let _944_v165;
          _944_v165 = new _dafny.CodePoint('q'.codePointAt(0));
          (globalState).f8 = _module.__default.fm1((_707_v1).f17, _944_v165, globalState);
        }
        if (_706_v0) {
          let _945_v166;
          _945_v166 = _dafny.Seq.of(_708_v2, _708_v2);
          (globalState).f9 = (_945_v166)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_945_v166).length))];
          _module.__default.m0(globalState);
          let _946_v167;
          _946_v167 = _dafny.MultiSet.fromElements(_708_v2, _708_v2);
          _928_v154 = (_928_v154).update((_707_v1).f17, new BigNumber((_dafny.Seq.update(_945_v166, _module.__default.safeIndex((((_946_v167).contains(_708_v2)) ? ((_946_v167).get(_708_v2)) : (_708_v2)), new BigNumber((_945_v166).length)), _708_v2)).length));
          let _rhs157 = _708_v2;
          let _rhs158 = _dafny.Seq.Concat(_dafny.Seq.Concat(_709_v3, _709_v3), _709_v3);
          _708_v2 = _rhs157;
          _709_v3 = _rhs158;
          let _947_v168;
          let _nw164 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
          _947_v168 = _nw164;
          let _nw165 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Set.Empty);
          _947_v168 = _nw165;
        } else {
          let _948_v169;
          _948_v169 = _dafny.Set.fromElements(_708_v2, _708_v2);
          let _949_v170;
          let _init24 = ((_950_v3) => function (_951_i23) {
            return _dafny.Seq.Concat(_950_v3, _950_v3);
          })(_709_v3);
          let _nw166 = Array((new BigNumber(5)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw166.length); _i0_24++) {
            _nw166[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _949_v170 = _nw166;
          let _952_v171;
          _952_v171 = _dafny.Seq.UnicodeFromString("rg");
          let _953_v172;
          _953_v172 = _module.D0.create_DC0(_708_v2);
          let _954_v173;
          _954_v173 = new _dafny.CodePoint('h'.codePointAt(0));
          let _955_v174;
          let _nw167 = new _module.C1();
          _nw167.__ctor((_707_v1).f17, (_707_v1).f17, _953_v172, _module.__default.fm1((_707_v1).f17, _954_v173, globalState));
          _955_v174 = _nw167;
          let _956_v175;
          _956_v175 = _dafny.Set.fromElements(_955_v174, _955_v174);
          let _957_v176;
          _957_v176 = _dafny.MultiSet.fromElements(new BigNumber((_956_v175).length));
          let _958_v177;
          _958_v177 = _dafny.Seq.of(new BigNumber((_952_v171).length));
          let _959_v178;
          _959_v178 = _dafny.Seq.of(new BigNumber((_958_v177).length), _708_v2);
          let _960_v179;
          let _nw168 = Array((new BigNumber(8)).toNumber());
          _nw168[(_dafny.ZERO).toNumber()] = _706_v0;
          _nw168[(_dafny.ONE).toNumber()] = _706_v0;
          _nw168[(new BigNumber(2)).toNumber()] = _706_v0;
          _nw168[(new BigNumber(3)).toNumber()] = (_707_v1).f17;
          _nw168[(new BigNumber(4)).toNumber()] = (_707_v1).f17;
          _nw168[(new BigNumber(5)).toNumber()] = (_957_v176).IsSubsetOf(_dafny.MultiSet.FromArray(_958_v177));
          _nw168[(new BigNumber(6)).toNumber()] = _706_v0;
          _nw168[(new BigNumber(7)).toNumber()] = (_955_v174).f13;
          _960_v179 = _nw168;
          let _rhs159 = _948_v169;
          let _rhs160 = !(true);
          let _rhs161 = _949_v170;
          let _rhs162 = _952_v171;
          let _rhs163 = _960_v179;
          _948_v169 = _rhs159;
          _706_v0 = _rhs160;
          _949_v170 = _rhs161;
          _952_v171 = _rhs162;
          _960_v179 = _rhs163;
          _952_v171 = _952_v171;
          _948_v169 = _948_v169;
          r2 = (_707_v1).f17;
          let _index138 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_960_v179).length));
          (_960_v179)[_index138] = _706_v0;
          let _index139 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_960_v179).length));
          let _rhs164 = _954_v173;
          let _rhs165 = (_708_v2).isLessThanOrEqualTo((_dafny.ZERO).minus(_708_v2));
          let _rhs166 = (_709_v3)[_module.__default.safeIndex(_708_v2, new BigNumber((_709_v3).length))];
          let _lhs106 = globalState;
          let _lhs107 = _960_v179;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_960_v179).length));
          _954_v173 = _rhs164;
          _lhs106.f8 = _rhs165;
          _lhs107[_lhs108] = _rhs166;
        }
        let _961_v180;
        _961_v180 = _dafny.Map.Empty.slice().updateUnsafe(_708_v2,_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qrta"), _dafny.Seq.UnicodeFromString("uqxf")));
        _961_v180 = (_961_v180).update(_708_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(826)), function (_962_i24) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }));
      }
      let _963_v181;
      _963_v181 = new _dafny.CodePoint('b'.codePointAt(0));
      r0 = _963_v181;
      r1 = _706_v0;
      r2 = false;
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _964_v0;
      _964_v0 = _dafny.Seq.UnicodeFromString("rewlpo");
      let _965_v1;
      _965_v1 = new BigNumber(218);
      let _966_v2;
      _966_v2 = _module.D8.create_DC19(_964_v0, _965_v1);
      let _967_v3;
      _967_v3 = _dafny.MultiSet.fromElements(_966_v2);
      let _968_v4;
      _968_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(559),_965_v1);
      let _969_v5;
      _969_v5 = _dafny.Map.Empty.slice().updateUnsafe(_965_v1,(((_968_v4).contains(_965_v1)) ? ((_968_v4).get(_965_v1)) : (new BigNumber((_964_v0).length))));
      let _970_v6;
      _970_v6 = _module.D0.create_DC0(new BigNumber((_968_v4).length));
      let _971_v7;
      let _nw169 = new _module.C1();
      _nw169.__ctor(p0, p0, _970_v6, p0);
      _971_v7 = _nw169;
      let _972_v8;
      _972_v8 = _dafny.Map.Empty.slice().updateUnsafe((_967_v3).update(_966_v2, _module.__default.abs(new BigNumber(-943))),_971_v7);
      (globalState).f8 = !((_972_v8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_967_v3,_971_v7))).equals(_972_v8);
      (globalState).f8 = ((p0) ? ((_971_v7).f13) : ((_971_v7).f13));
      let _hi8 = _module.__default.safeModuloInt(_965_v1, (_dafny.ZERO).minus(_965_v1));
      for (let _973_i0 = (_965_v1).minus(_965_v1); _973_i0.isLessThan(_hi8); _973_i0 = _973_i0.plus(_dafny.ONE)) {
        let _974_v9;
        _974_v9 = new _dafny.CodePoint('n'.codePointAt(0));
        let _975_v10;
        _975_v10 = _dafny.Seq.of(new BigNumber((_module.__default.fm37((_971_v7).f13, (_dafny.ZERO).minus(_973_i0), _973_i0, _965_v1, globalState)).length), new BigNumber((_module.__default.fm38((_971_v7).f13, (_971_v7).f13, _974_v9, _973_i0, globalState)).cardinality()), _965_v1, _973_i0, new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), _974_v9)).length));
        _965_v1 = _module.__default.safeModuloInt(_973_i0, (_975_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kcbsmgk")).length), new BigNumber((_975_v10).length))]);
        let _976_v11;
        _976_v11 = _module.D1.create_DC4(_974_v9);
        let _977_v12;
        _977_v12 = _dafny.MultiSet.fromElements((_971_v7).f13);
        let _978_v13;
        _978_v13 = _dafny.Set.fromElements((_this).fm20((_971_v7).f13, (((_977_v12).contains(_module.__default.fm1((_971_v7).f13, _974_v9, globalState))) ? ((_977_v12).get(_module.__default.fm1((_971_v7).f13, _974_v9, globalState))) : ((_975_v10)[_module.__default.safeIndex(_973_i0, new BigNumber((_975_v10).length))])), globalState), _965_v1);
        let _979_v14;
        let _nw170 = new _module.C0();
        _nw170.__ctor(((_971_v7).fm9(_976_v11, _978_v13, globalState)).isEqualTo(_965_v1));
        _979_v14 = _nw170;
        let _980_v15;
        let _nw171 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _980_v15 = _nw171;
        let _981_v16;
        _981_v16 = _module.D9.create_DC21(_980_v15);
        let _982_v17;
        let _nw172 = Array((new BigNumber(24)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = _980_v15;
        _nw172[(_dafny.ONE).toNumber()] = _980_v15;
        _nw172[(new BigNumber(2)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(3)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(4)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(5)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(6)).toNumber()] = (_981_v16).dtor_cf39;
        _nw172[(new BigNumber(7)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(8)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(9)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(10)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(11)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(12)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(13)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(14)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(15)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(16)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(17)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(18)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(19)).toNumber()] = ((p0) ? (_980_v15) : (_980_v15));
        _nw172[(new BigNumber(20)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(21)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(22)).toNumber()] = _980_v15;
        _nw172[(new BigNumber(23)).toNumber()] = _980_v15;
        _982_v17 = _nw172;
        let _index140 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_982_v17).length));
        (_982_v17)[_index140] = ((!((_979_v14).f17)) ? (_980_v15) : (_980_v15));
        let _index141 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_982_v17).length));
        (_982_v17)[_index141] = _980_v15;
        let _983_v18;
        _983_v18 = _module.D2.create_DC8(_965_v1, _module.__default.fm3(_973_i0, (_971_v7).f13, globalState));
        let _984_v19;
        _984_v19 = _dafny.Set.fromElements(_974_v9, (_983_v18).dtor_cf16, _974_v9, _974_v9, _974_v9);
        let _985_v20;
        _985_v20 = _dafny.Seq.of(!(_973_i0).isEqualTo(new BigNumber((_984_v19).length)));
        let _986_v21;
        let _nw173 = Array((new BigNumber(27)).toNumber()).fill(false);
        _986_v21 = _nw173;
        let _987_v22;
        _987_v22 = _dafny.Map.Empty.slice().updateUnsafe(_973_i0,_986_v21);
        let _index142 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_986_v21).length));
        (_986_v21)[_index142] = (_979_v14).f17;
        let _988_v23;
        _988_v23 = _dafny.Map.Empty.slice().updateUnsafe((_971_v7).f13,(_979_v14).f17);
        let _index143 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_986_v21).length));
        let _rhs167 = (_dafny.ZERO).minus((_965_v1).plus(((_this).fm20((_971_v7).f13, (_dafny.ZERO).minus(new BigNumber(((_988_v23).update(true, (_971_v7).f13)).length)), globalState)).minus(_965_v1)));
        let _rhs168 = _dafny.Seq.Concat(_dafny.Seq.of(!(p0)), _985_v20);
        let _rhs169 = _987_v22;
        let _rhs170 = ((_979_v14).f17) && (!((_985_v20)[_module.__default.safeIndex(_973_i0, new BigNumber((_985_v20).length))]));
        let _rhs171 = _module.__default.fm1(_dafny.areEqual(_dafny.Seq.of(_965_v1), _dafny.Seq.of(_973_i0)), _974_v9, globalState);
        let _lhs109 = globalState;
        let _lhs110 = _986_v21;
        let _lhs111 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_986_v21).length));
        _965_v1 = _rhs167;
        _985_v20 = _rhs168;
        _987_v22 = _rhs169;
        _lhs109.f8 = _rhs170;
        _lhs110[_lhs111] = _rhs171;
      }
      let _989_v24;
      let _init25 = ((_990_v7) => function (_991_i1) {
        return (_990_v7).f13;
      })(_971_v7);
      let _nw174 = Array((new BigNumber(17)).toNumber());
      for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw174.length); _i0_25++) {
        _nw174[_i0_25] = _init25(new BigNumber(_i0_25));
      }
      _989_v24 = _nw174;
      let _992_v25;
      let _nw175 = Array((new BigNumber(2)).toNumber());
      _nw175[(_dafny.ZERO).toNumber()] = _989_v24;
      _nw175[(_dafny.ONE).toNumber()] = _989_v24;
      _992_v25 = _nw175;
      _992_v25 = _992_v25;
      let _index144 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length));
      (_989_v24)[_index144] = (_971_v7).f13;
      let _993_v26;
      _993_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _index145 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length));
      (_989_v24)[_index145] = !(p0) || ((((_993_v26).contains((_971_v7).f13)) ? ((_993_v26).get((_971_v7).f13)) : ((_971_v7).f13)));
      let _994_v27;
      _994_v27 = new _dafny.CodePoint('y'.codePointAt(0));
      let _995_v28;
      _995_v28 = _dafny.Set.fromElements(_994_v27);
      let _996_v29;
      _996_v29 = _995_v28;
      let _source9 = _996_v29;
      let _997___mcc_h0 = _source9;
      let _998_cf27 = _997___mcc_h0;
      let _source10 = (((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))]) ? (_996_v29) : (_996_v29));
      let _999___mcc_h1 = _source10;
      let _1000_cf27 = _999___mcc_h1;
      let _1001_v30;
      _1001_v30 = _dafny.MultiSet.fromElements(false);
      let _1002_v31;
      _1002_v31 = _dafny.Seq.of(_965_v1);
      let _1003_v32;
      _1003_v32 = _dafny.Map.Empty.slice().updateUnsafe(!((_971_v7).f13),(((_971_v7).f13) ? ((_dafny.ZERO).minus(_965_v1)) : ((_1002_v31)[_module.__default.safeIndex(_965_v1, new BigNumber((_1002_v31).length))])));
      let _index146 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length));
      let _rhs172 = false;
      let _rhs173 = _1001_v30;
      let _rhs174 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_965_v1, _965_v1), _965_v1);
      let _rhs175 = ((_965_v1).multipliedBy(new BigNumber((_module.__default.fm39(new BigNumber(249), globalState)).length))).minus(_module.__default.safeDivisionInt(_965_v1, new BigNumber(15)));
      let _rhs176 = (((_971_v7).f13) ? ((_1003_v32).Merge(_1003_v32)) : (((_1003_v32).update(false, _965_v1)).Merge(_1003_v32)));
      let _lhs112 = _989_v24;
      let _lhs113 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length));
      let _lhs114 = globalState;
      let _lhs115 = globalState;
      _lhs112[_lhs113] = _rhs172;
      _1001_v30 = _rhs173;
      _lhs114.f9 = _rhs174;
      _lhs115.f9 = _rhs175;
      _1003_v32 = _rhs176;
      _989_v24 = ((((_971_v7).f13) && ((_971_v7).f13)) ? (_989_v24) : (_989_v24));
      let _1004_v33;
      let _init26 = function (_1005_i2) {
        return (_1005_i2).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length));
      };
      let _nw176 = Array((new BigNumber(3)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw176.length); _i0_26++) {
        _nw176[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _1004_v33 = _nw176;
      let _index147 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1004_v33).length));
      (_1004_v33)[_index147] = _965_v1;
      let _index148 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1004_v33).length));
      (_1004_v33)[_index148] = _965_v1;
      let _index149 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length));
      (_989_v24)[_index149] = true;
      let _index150 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length));
      (_989_v24)[_index150] = (_971_v7).f13;
      let _1006_v34;
      _1006_v34 = _dafny.MultiSet.fromElements(_965_v1);
      let _1007_v35;
      _1007_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(192),_1006_v34);
      let _1008_v36;
      _1008_v36 = _dafny.MultiSet.fromElements((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))], ((((_1007_v35).contains(_965_v1)) ? ((_1007_v35).get(_965_v1)) : (_1006_v34))).IsProperSubsetOf(_1006_v34), (((_993_v26).contains((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))])) ? ((_993_v26).get((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))])) : (_module.__default.fm1(p0, _994_v27, globalState))), (_971_v7).f13);
      _1008_v36 = _1008_v36;
      if (true) {
        (globalState).f8 = !(p0);
        (globalState).f9 = _965_v1;
        let _1009_v37;
        _1009_v37 = _dafny.Map.Empty.slice().updateUnsafe(_994_v27,_989_v24);
        let _1010_v38;
        _1010_v38 = _dafny.Map.Empty.slice().updateUnsafe((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))],_1009_v37);
        _1010_v38 = (_1010_v38).update((_971_v7).f13, _1009_v37);
        let _1011_v39;
        let _nw177 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Set.Empty);
        _1011_v39 = _nw177;
        _1011_v39 = (((p0) === ((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))])) ? (_1011_v39) : (_1011_v39));
        _module.__default.m0(globalState);
      } else {
        let _1012_v40;
        _1012_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), ((_1013_v1) => function (_1014_i3) {
          return _1013_v1;
        })(_965_v1))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(952)), function (_1015_i4) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }));
        let _1016_v41;
        _1016_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v40,(_965_v1).isLessThan(_965_v1));
        let _1017_v43;
        _1017_v43 = _dafny.Seq.of(true, !(!((_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))])), (_971_v7).f13);
        let _1018_v44;
        _1018_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1006_v34,(_1017_v43)[_module.__default.safeIndex(_965_v1, new BigNumber((_1017_v43).length))]);
        let _1019_v45;
        _1019_v45 = _dafny.Seq.of(new BigNumber(575), _965_v1);
        let _1020_v46;
        _1020_v46 = _dafny.Seq.of((_dafny.ZERO).minus((_1019_v45)[_module.__default.safeIndex(new BigNumber((_968_v4).length), new BigNumber((_1019_v45).length))]));
        _1016_v41 = (_1016_v41).update(function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of (_1018_v44).Keys.Elements) {
            let _1021_v42 = _compr_43;
            if ((_1018_v44).contains(_1021_v42)) {
              _coll43.push([_1021_v42,_dafny.Seq.UnicodeFromString("j")]);
            }
          }
          return _coll43;
        }(), ((_dafny.ZERO).minus(_965_v1)).isLessThan(new BigNumber((_1019_v45).length)));
        (globalState).f9 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_964_v0, _964_v0)).length), _965_v1);
        let _1022_v47;
        let _nw178 = new _module.C1();
        _nw178.__ctor((_971_v7).f13, (_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))], _970_v6, (_989_v24)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_989_v24).length))]);
        _1022_v47 = _nw178;
        let _1023_v48;
        _1023_v48 = _dafny.MultiSet.fromElements(_1022_v47, _1022_v47, _1022_v47);
        let _1024_v49;
        _1024_v49 = _dafny.Map.Empty.slice().updateUnsafe(((_1023_v48).update(_1022_v47, _module.__default.abs(_965_v1))).IsSubsetOf(_1023_v48),_965_v1);
        _1024_v49 = ((true) ? (_module.__default.fm25(_965_v1, globalState)) : (_1024_v49));
        let _1025_v50;
        let _nw179 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1025_v50 = _nw179;
        let _rhs177 = _1008_v36;
        let _rhs178 = _1025_v50;
        let _rhs179 = _1025_v50;
        let _rhs180 = _1025_v50;
        _1008_v36 = _rhs177;
        _1025_v50 = _rhs178;
        _1025_v50 = _rhs179;
        _1025_v50 = _rhs180;
        let _1026_v52;
        let _init27 = ((_1027_v0) => function (_1028_i5) {
          return (function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_dafny.Set.fromElements(_1027_v0)).Elements) {
              let _1029_v51 = _compr_44;
              if ((_dafny.Set.fromElements(_1027_v0)).contains(_1029_v51)) {
                _coll44.add(_1029_v51);
              }
            }
            return _coll44;
          }()).Difference(_dafny.Set.fromElements(_1027_v0));
        })(_964_v0);
        let _nw180 = Array((new BigNumber(18)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw180.length); _i0_27++) {
          _nw180[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1026_v52 = _nw180;
        let _1030_v53;
        _1030_v53 = _dafny.Set.fromElements(_964_v0, _964_v0, _964_v0, _964_v0);
        let _index151 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_1026_v52).length));
        (_1026_v52)[_index151] = (_1030_v53).Intersect(_dafny.Set.fromElements(_module.__default.fm2(_965_v1, _1006_v34, _968_v4, globalState), _964_v0, (_this).fm19(globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(252)), ((_1031_v27) => function (_1032_i6) {
          return _1031_v27;
        })(_994_v27))));
        let _1033_v54;
        _1033_v54 = _dafny.Map.Empty.slice().updateUnsafe((_1030_v53).Union(_1030_v53),_1030_v53);
        let _index152 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_1026_v52).length));
        (_1026_v52)[_index152] = (((_1033_v54).contains(_1030_v53)) ? ((_1033_v54).get(_1030_v53)) : (_1030_v53));
      }
      r0 = _965_v1;
      let _1034_v55;
      _1034_v55 = _dafny.Set.fromElements(p0, p0);
      let _1035_v56;
      _1035_v56 = _dafny.Set.fromElements((_965_v1).isLessThanOrEqualTo(new BigNumber((_1034_v55).length)), _module.__default.fm1(true, new _dafny.CodePoint('w'.codePointAt(0)), globalState));
      r1 = _1034_v55;
      return [r0, r1];
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1036_v0;
      let _init28 = function (_1037_i0) {
        return true;
      };
      let _nw181 = Array((new BigNumber(18)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw181.length); _i0_28++) {
        _nw181[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _1036_v0 = _nw181;
      let _1038_v1;
      _1038_v1 = true;
      let _index153 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length));
      (_1036_v0)[_index153] = _1038_v1;
      let _index154 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length));
      (_1036_v0)[_index154] = _1038_v1;
      let _1039_v2;
      _1039_v2 = new _dafny.CodePoint('j'.codePointAt(0));
      let _index155 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length));
      (_1036_v0)[_index155] = _module.__default.fm1((_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))], _1039_v2, globalState);
      let _1040_i1;
      _1040_i1 = _dafny.ZERO;
      L5: {
        while (_1038_v1) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1040_i1)) {
              break L5;
            }
            _1040_i1 = (_1040_i1).plus(_dafny.ONE);
            let _1041_v3;
            _1041_v3 = _dafny.Seq.UnicodeFromString("wlm");
            let _index156 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length));
            (_1036_v0)[_index156] = (_module.__default.fm1((_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))], _1039_v2, globalState)) && (_dafny.Seq.IsPrefixOf(_1041_v3, _1041_v3));
            let _index157 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length));
            (_1036_v0)[_index157] = false;
            r0 = p1;
            let _1042_v4;
            let _init29 = ((_1043_p1) => function (_1044_i2) {
              return _module.__default.safeDivisionInt(_1044_i2, _1043_p1);
            })(p1);
            let _nw182 = Array((new BigNumber(18)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw182.length); _i0_29++) {
              _nw182[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _1042_v4 = _nw182;
            let _index158 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1042_v4).length));
            (_1042_v4)[_index158] = (p1).minus(p1);
            let _index159 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1042_v4).length));
            let _rhs181 = _module.__default.safeDivisionInt((p1).plus(p1), p1);
            let _rhs182 = p1;
            let _lhs116 = globalState;
            let _lhs117 = _1042_v4;
            let _lhs118 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1042_v4).length));
            _lhs116.f9 = _rhs181;
            _lhs117[_lhs118] = _rhs182;
          }
        }
      }
      let _1045_v5;
      _1045_v5 = _module.D5.create_DC14((_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))]);
      let _1046_v6;
      _1046_v6 = _module.D0.create_DC0(p1);
      let _1047_v7;
      _1047_v7 = _dafny.Seq.of(_dafny.Seq.of(_1038_v1));
      let _1048_v8;
      let _nw183 = new _module.C1();
      _nw183.__ctor((_1045_v5).dtor_cf26, (_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))], _1046_v6, (new BigNumber(((_1047_v7)[_module.__default.safeIndex(new BigNumber(803), new BigNumber((_1047_v7).length))]).length)).isEqualTo(new BigNumber(222)));
      _1048_v8 = _nw183;
      if ((_1048_v8).f20) {
        if ((_module.__default.safeDivisionInt(p1, p1)).isEqualTo(p1)) {
          let _1049_v9;
          _1049_v9 = _dafny.MultiSet.fromElements((_1048_v8).f19);
          let _1050_v10;
          _1050_v10 = _module.D0.create_DC2(_1049_v9, (_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))], p1, (_1048_v8).f20);
          let _1051_v11;
          _1051_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1050_v10);
          _1051_v11 = (_1051_v11).update(p1, _1050_v10);
          (globalState).f9 = (_dafny.ZERO).minus(p1);
          (globalState).f8 = !(!((_1048_v8).f20));
          let _index160 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length));
          (_1036_v0)[_index160] = (p1).isLessThan((p1).multipliedBy(p1));
          (globalState).f9 = new BigNumber(749);
        } else {
          let _1052_v12;
          _1052_v12 = _dafny.Seq.of(!((_1048_v8).f20));
          r0 = (new BigNumber((_dafny.Seq.Concat(_1052_v12, _1052_v12)).length)).multipliedBy((_dafny.ZERO).minus(p1));
          let _1053_v13;
          _1053_v13 = _dafny.Seq.UnicodeFromString("fddf");
          _1053_v13 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1053_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(441)), ((_1054_v2) => function (_1055_i3) {
            return _1054_v2;
          })(_1039_v2))), _dafny.Seq.Concat(_1053_v13, _1053_v13));
          _1053_v13 = _dafny.Seq.Concat((((_1048_v8).f20) ? (_dafny.Seq.UnicodeFromString("phdrqre")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-297)), ((_1056_v2) => function (_1057_i4) {
            return _1056_v2;
          })(_1039_v2)))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qlro"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), ((_1058_v2) => function (_1059_i5) {
            return _1058_v2;
          })(_1039_v2))));
          let _1060_v14;
          _1060_v14 = _dafny.MultiSet.fromElements(_1038_v1, true);
          let _1061_v15;
          let _nw184 = new _module.C1();
          _nw184.__ctor((_1048_v8).f20, !((_1060_v14).IsSubsetOf(_1060_v14)), _module.__default.fm40(new _dafny.CodePoint('u'.codePointAt(0)), globalState), (_1048_v8).f19);
          _1061_v15 = _nw184;
          let _rhs183 = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_1053_v13).length)).plus(p1)));
          let _rhs184 = _1036_v0;
          let _rhs185 = _1061_v15;
          let _rhs186 = (_1061_v15).f19;
          let _lhs119 = globalState;
          r0 = _rhs183;
          _1036_v0 = _rhs184;
          _1061_v15 = _rhs185;
          _lhs119.f8 = _rhs186;
          (globalState).f9 = ((_dafny.ZERO).minus(_module.__default.fm0((_1048_v8).f19, globalState))).minus((p1).plus(new BigNumber(860)));
        }
        let _rhs187 = _module.__default.fm1((_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))], _1039_v2, globalState);
        let _rhs188 = _1039_v2;
        let _rhs189 = (_1048_v8).f20;
        let _rhs190 = _module.__default.fm0((p1).isEqualTo(new BigNumber(-324)), globalState);
        let _lhs120 = globalState;
        let _lhs121 = globalState;
        _1038_v1 = _rhs187;
        _1039_v2 = _rhs188;
        _lhs120.f8 = _rhs189;
        _lhs121.f9 = _rhs190;
        r0 = (_dafny.ZERO).minus(p1);
        (globalState).f9 = p1;
        (globalState).f9 = p1;
      } else {
        _1047_v7 = _1047_v7;
        let _1062_v16;
        let _nw185 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1062_v16 = _nw185;
        let _1063_v17;
        _1063_v17 = _dafny.Seq.UnicodeFromString("ncq");
        let _index161 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1062_v16).length));
        (_1062_v16)[_index161] = _1063_v17;
        let _1064_v18;
        let _nw186 = Array((new BigNumber(22)).toNumber()).fill([]);
        _1064_v18 = _nw186;
        let _1065_v19;
        let _nw187 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1065_v19 = _nw187;
        let _1066_v20;
        let _nw188 = Array((new BigNumber(10)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = _1065_v19;
        _nw188[(_dafny.ONE).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(2)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(3)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(4)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(5)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(6)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(7)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(8)).toNumber()] = _1065_v19;
        _nw188[(new BigNumber(9)).toNumber()] = _1065_v19;
        _1066_v20 = _nw188;
        let _index162 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1064_v18).length));
        (_1064_v18)[_index162] = _1066_v20;
        let _index163 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1062_v16).length));
        let _index164 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1064_v18).length));
        let _rhs191 = _1063_v17;
        let _rhs192 = _1066_v20;
        let _lhs122 = _1062_v16;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1062_v16).length));
        let _lhs124 = _1064_v18;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1064_v18).length));
        _lhs122[_lhs123] = _rhs191;
        _lhs124[_lhs125] = _rhs192;
        let _1067_v21;
        _1067_v21 = _dafny.MultiSet.fromElements(true, true);
        let _1068_v22;
        _1068_v22 = _module.D0.create_DC2(_1067_v21, (_1048_v8).f19, p1, _1038_v1);
        let _index165 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1065_v19).length));
        (_1065_v19)[_index165] = (_1068_v22).dtor_cf6;
        let _index166 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1065_v19).length));
        (_1065_v19)[_index166] = (p1).plus((_dafny.ZERO).minus((new BigNumber((_module.__default.fm23(false, globalState)).length)).plus(p1)));
        (globalState).f9 = p1;
        _1039_v2 = _1039_v2;
      }
      let _rhs193 = ((!((_1036_v0)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_1036_v0).length))])) ? (_1038_v1) : ((_1048_v8).f20));
      let _rhs194 = _1038_v1;
      let _rhs195 = p1;
      let _rhs196 = p1;
      _1038_v1 = _rhs193;
      _1038_v1 = _rhs194;
      r0 = _rhs195;
      r0 = _rhs196;
      r0 = (((_1048_v8).f20) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus((p1).minus(p1)))) : (p1));
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f11 = [];
      this.f21 = false;
      this._f22 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f21, f22, f11) {
      let _this = this;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f11 = f11;
      return;
    }
    fm47(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vedhv")).length));
    };
    fm48(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(new BigNumber(416))).Difference(_dafny.Set.fromElements(new BigNumber((function () {
        let _coll45 = new _dafny.Map();
        for (const _compr_45 of (_dafny.Seq.of(_dafny.Seq.of(false, (_this).f22))).Elements) {
          let _1069_v0 = _compr_45;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(false, (_this).f22)), _1069_v0)) {
            _coll45.push([_1069_v0,new BigNumber((_dafny.Set.fromElements(false, false)).length)]);
          }
        }
        return _coll45;
      }()).length)))).Intersect((_dafny.Set.fromElements(new BigNumber(807), new BigNumber(539))).Difference(_dafny.Set.fromElements(new BigNumber(859))));
    };
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      if (p0) {
        let _1070_v0;
        let _init30 = function (_1071_i0) {
          return _this.f21;
        };
        let _nw189 = Array((_dafny.ONE).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw189.length); _i0_30++) {
          _nw189[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1070_v0 = _nw189;
        _1070_v0 = _1070_v0;
        (globalState).f8 = _this.f21;
        let _1072_v1;
        let _nw190 = new _module.C2();
        _nw190.__ctor(((_this.f21) ? ((_this).f11) : ((_this).f11)));
        _1072_v1 = _nw190;
        let _1073_v2;
        _1073_v2 = new BigNumber(708);
        let _1074_v3;
        _1074_v3 = _dafny.Seq.UnicodeFromString("vilgvrgs");
        let _rhs197 = _module.__default.safeDivisionInt(_1073_v2, (new BigNumber(873)).minus(_1073_v2));
        let _rhs198 = !(!((_module.__default.safeDivisionInt(new BigNumber((_1074_v3).length), _1073_v2)).isLessThan(_1073_v2)));
        let _lhs126 = globalState;
        r0 = _rhs197;
        _lhs126.f8 = _rhs198;
        (globalState).f9 = ((p0) ? (_1073_v2) : (_1073_v2));
      } else {
        (globalState).f8 = _this.f21;
        let _1075_v4;
        let _nw191 = new _module.C0();
        _nw191.__ctor(_this.f21);
        _1075_v4 = _nw191;
        let _1076_v5;
        _1076_v5 = new BigNumber(180);
        let _1077_v6;
        _1077_v6 = _dafny.Seq.UnicodeFromString("aiuask");
        let _1078_v7;
        _1078_v7 = _dafny.Map.Empty.slice().updateUnsafe((_1075_v4).f17,new BigNumber((_1077_v6).length));
        let _1079_v8;
        _1079_v8 = _module.D8.create_DC19(_dafny.Seq.UnicodeFromString("s"), _1076_v5);
        let _1080_v9;
        let _nw192 = new _module.C2();
        _nw192.__ctor((_this).f11);
        _1080_v9 = _nw192;
        let _1081_v10;
        _1081_v10 = _dafny.Set.fromElements(_1080_v9);
        let _1082_v11;
        _1082_v11 = _dafny.MultiSet.fromElements((_this).f22);
        let _1083_v12;
        let _nw193 = Array((new BigNumber(26)).toNumber());
        _nw193[(_dafny.ZERO).toNumber()] = _module.__default.fm0((_this).f22, globalState);
        _nw193[(_dafny.ONE).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(2)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(3)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(4)).toNumber()] = new BigNumber(470);
        _nw193[(new BigNumber(5)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(6)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(7)).toNumber()] = new BigNumber(-750);
        _nw193[(new BigNumber(8)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(9)).toNumber()] = (new BigNumber((_1077_v6).length)).minus(_1076_v5);
        _nw193[(new BigNumber(10)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(11)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(12)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(13)).toNumber()] = new BigNumber(680);
        _nw193[(new BigNumber(14)).toNumber()] = new BigNumber(991);
        _nw193[(new BigNumber(15)).toNumber()] = (((_1078_v7).contains(_this.f21)) ? ((_1078_v7).get(_this.f21)) : (_1076_v5));
        _nw193[(new BigNumber(16)).toNumber()] = ((_module.D4.create_DC11(_1076_v5, (_dafny.ZERO).minus((_1079_v8).dtor_cf35), !((_this).f22))).dtor_cf20).minus(new BigNumber((_1081_v10).length));
        _nw193[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-340)), function (_1084_i1) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        })).length);
        _nw193[(new BigNumber(18)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(19)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(20)).toNumber()] = (((_this).f22) ? (new BigNumber((_dafny.Set.fromElements(_1075_v4)).length)) : (new BigNumber((_1082_v11).cardinality())));
        _nw193[(new BigNumber(21)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(22)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(23)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(24)).toNumber()] = _1076_v5;
        _nw193[(new BigNumber(25)).toNumber()] = _1076_v5;
        _1083_v12 = _nw193;
        let _index167 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length));
        (_1083_v12)[_index167] = (_1080_v9).fm20(!(_this.f21), _1076_v5, globalState);
        let _index168 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length));
        let _rhs199 = (_this).f22;
        let _rhs200 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1076_v5, _dafny.ZERO));
        let _lhs127 = globalState;
        let _lhs128 = _1083_v12;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length));
        _lhs127.f8 = _rhs199;
        _lhs128[_lhs129] = _rhs200;
        let _1085_v13;
        let _init31 = ((_1086_p0) => function (_1087_i2) {
          return _1086_p0;
        })(p0);
        let _nw194 = Array((new BigNumber(5)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw194.length); _i0_31++) {
          _nw194[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1085_v13 = _nw194;
        let _1088_v14;
        let _nw195 = Array((new BigNumber(25)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = _1085_v13;
        _nw195[(_dafny.ONE).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(2)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(3)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(4)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(5)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(6)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(7)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(8)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(9)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(10)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(11)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(12)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(13)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(14)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(15)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(16)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(17)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(18)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(19)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(20)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(21)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(22)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(23)).toNumber()] = _1085_v13;
        _nw195[(new BigNumber(24)).toNumber()] = _1085_v13;
        _1088_v14 = _nw195;
        let _index169 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1088_v14).length));
        (_1088_v14)[_index169] = _1085_v13;
        let _index170 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1088_v14).length));
        (_1088_v14)[_index170] = _1085_v13;
        let _index171 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length));
        (_1083_v12)[_index171] = ((_1083_v12)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length))]).multipliedBy(((_1083_v12)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length))]).multipliedBy((_1083_v12)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1083_v12).length))]));
      }
      let _1089_v15;
      _1089_v15 = new BigNumber(989);
      let _1090_v16;
      _1090_v16 = _module.D0.create_DC0(_1089_v15);
      let _source11 = _1090_v16;
      if (_source11.is_DC1) {
        let _1091___mcc_h0 = (_source11).cf1;
        let _1092___mcc_h1 = (_source11).cf2;
        let _1093___mcc_h2 = (_source11).cf3;
        let _1094_cf3 = _1093___mcc_h2;
        let _1095_cf2 = _1092___mcc_h1;
        let _1096_cf1 = _1091___mcc_h0;
        let _1097_v17;
        _1097_v17 = _dafny.MultiSet.fromElements((_this).f22, (_this).f22, (_this).f22);
        _1097_v17 = _1097_v17;
        _1089_v15 = _1094_cf3;
        let _1098_v18;
        _1098_v18 = _dafny.Set.fromElements(p0);
        r1 = _1098_v18;
        let _1099_v19;
        _1099_v19 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1100_v20;
        _1100_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1099_v19)).cardinality()),(_dafny.ZERO).minus((_1089_v15).plus(_1089_v15)));
        let _1101_v22;
        _1101_v22 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),_1094_cf3));
        let _1102_v23;
        _1102_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1099_v19,_1101_v22);
        (globalState).f9 = (((_1100_v20).contains((new BigNumber((function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of ((((_1102_v23).contains(_1099_v19)) ? ((_1102_v23).get(_1099_v19)) : (_1101_v22))).Elements) {
            let _1104_v21 = _compr_47;
            if (((((_1102_v23).contains(_1099_v19)) ? ((_1102_v23).get(_1099_v19)) : (_1101_v22))).contains(_1104_v21)) {
              _coll47.push([_1104_v21,p0]);
            }
          }
          return _coll47;
        }()).length)).multipliedBy(_1089_v15))) ? ((_1100_v20).get((new BigNumber((function () {
          let _coll46 = new _dafny.Map();
          for (const _compr_46 of ((((_1102_v23).contains(_1099_v19)) ? ((_1102_v23).get(_1099_v19)) : (_1101_v22))).Elements) {
            let _1103_v21 = _compr_46;
            if (((((_1102_v23).contains(_1099_v19)) ? ((_1102_v23).get(_1099_v19)) : (_1101_v22))).contains(_1103_v21)) {
              _coll46.push([_1103_v21,p0]);
            }
          }
          return _coll46;
        }()).length)).multipliedBy(_1089_v15))) : (_1094_cf3));
      } else if (_source11.is_DC2) {
        let _1105___mcc_h3 = (_source11).cf4;
        let _1106___mcc_h4 = (_source11).cf5;
        let _1107___mcc_h5 = (_source11).cf6;
        let _1108___mcc_h6 = (_source11).cf7;
        let _1109_cf7 = _1108___mcc_h6;
        let _1110_cf6 = _1107___mcc_h5;
        let _1111_cf5 = _1106___mcc_h4;
        let _1112_cf4 = _1105___mcc_h3;
        let _1113_v24;
        _1113_v24 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1114_v25;
        let _1115_v26;
        let _out10;
        let _out11;
        let _outcollector3 = (_this).m13(((_this.f21) ? (_1110_cf6) : (_1089_v15)), (_this).f22, _1110_cf6, _module.__default.fm1(_1109_cf7, _1113_v24, globalState), globalState);
        _out10 = _outcollector3[0];
        _out11 = _outcollector3[1];
        _1114_v25 = _out10;
        _1115_v26 = _out11;
        if (!((_this).f22)) {
          let _1116_v27;
          _1116_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), ((_1117_v15) => function (_1118_i3) {
            return _1117_v15;
          })(_1089_v15))).length),_1110_cf6);
          let _1119_v28;
          _1119_v28 = _dafny.Seq.of(new BigNumber((_1116_v27).length));
          _1119_v28 = _1119_v28;
          let _1120_v29;
          _1120_v29 = _dafny.MultiSet.fromElements(_1110_cf6, _1089_v15);
          let _1121_v30;
          _1121_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1112_cf4,_1120_v29);
          let _1122_v31;
          _1122_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v15,_this.f21);
          let _1123_v32;
          _1123_v32 = _dafny.Seq.UnicodeFromString("bfiaytxu");
          let _1124_v33;
          _1124_v33 = _dafny.Seq.of((((_1122_v31).contains(new BigNumber((_1123_v32).length))) ? ((_1122_v31).get(new BigNumber((_1123_v32).length))) : (_1114_v25)));
          _1121_v30 = ((_1121_v30).update(_dafny.MultiSet.FromArray(_1124_v33), _1120_v29)).Merge(_1121_v30);
          let _1125_v34;
          _1125_v34 = _module.D2.create_DC5(_1123_v32);
          let _1126_v35;
          let _nw196 = Array((new BigNumber(23)).toNumber());
          _nw196[(_dafny.ZERO).toNumber()] = _1125_v34;
          _nw196[(_dafny.ONE).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(2)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(3)).toNumber()] = _module.__default.fm49(_1110_cf6, globalState);
          _nw196[(new BigNumber(4)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(5)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(6)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(7)).toNumber()] = (((_this).f22) ? (_module.D2.create_DC5(_dafny.Seq.UnicodeFromString("u"))) : (_1125_v34));
          _nw196[(new BigNumber(8)).toNumber()] = _module.D2.create_DC5(_1123_v32);
          _nw196[(new BigNumber(9)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(10)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(11)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(12)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(13)).toNumber()] = _module.D2.create_DC5(_dafny.Seq.UnicodeFromString("eywbyk"));
          _nw196[(new BigNumber(14)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(15)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(16)).toNumber()] = _module.D2.create_DC5(_1123_v32);
          _nw196[(new BigNumber(17)).toNumber()] = _module.__default.fm49(_1089_v15, globalState);
          _nw196[(new BigNumber(18)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(19)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(20)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(21)).toNumber()] = _1125_v34;
          _nw196[(new BigNumber(22)).toNumber()] = _1125_v34;
          _1126_v35 = _nw196;
          let _index172 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1126_v35).length));
          (_1126_v35)[_index172] = _1125_v34;
          let _index173 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1126_v35).length));
          (_1126_v35)[_index173] = function (_pat_let13_0) {
            return function (_1127_dt__update__tmp_h0) {
              return function (_pat_let14_0) {
                return function (_1128_dt__update_hcf10_h0) {
                  return _module.D2.create_DC5(_1128_dt__update_hcf10_h0);
                }(_pat_let14_0);
              }(_dafny.Seq.UnicodeFromString("shuvv"));
            }(_pat_let13_0);
          }(_module.D2.create_DC5(_1123_v32));
          let _1129_v36;
          let _nw197 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _1129_v36 = _nw197;
          let _1130_v37;
          _1130_v37 = _dafny.MultiSet.fromElements(_1129_v36);
          _1110_cf6 = (((_1130_v37).contains(_1129_v36)) ? ((_1130_v37).get(_1129_v36)) : (_1089_v15));
          let _1131_v38;
          let _nw198 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _1131_v38 = _nw198;
          _1131_v38 = _1131_v38;
        } else {
          let _1132_v39;
          let _init32 = function (_1133_i4) {
            return _dafny.Seq.UnicodeFromString("dqxxcuign");
          };
          let _nw199 = Array((new BigNumber(9)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw199.length); _i0_32++) {
            _nw199[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1132_v39 = _nw199;
          let _1134_v40;
          let _nw200 = Array((new BigNumber(12)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _1132_v39;
          _nw200[(_dafny.ONE).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(2)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(3)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(4)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(5)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(6)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(7)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(8)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(9)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(10)).toNumber()] = _1132_v39;
          _nw200[(new BigNumber(11)).toNumber()] = _1132_v39;
          _1134_v40 = _nw200;
          let _index174 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1134_v40).length));
          (_1134_v40)[_index174] = _1132_v39;
          let _index175 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1134_v40).length));
          (_1134_v40)[_index175] = _1132_v39;
          let _1135_v41;
          _1135_v41 = _dafny.Seq.of(_1113_v24);
          let _1136_v42;
          _1136_v42 = _dafny.Seq.of(_1089_v15, (_1089_v15).minus(_module.__default.fm0(_1109_cf7, globalState)), _1089_v15, _1089_v15, _module.__default.fm0(!(_1114_v25), globalState));
          let _rhs201 = _1135_v41;
          let _rhs202 = _1136_v42;
          let _rhs203 = _this.f21;
          let _lhs130 = globalState;
          _1135_v41 = _rhs201;
          _1136_v42 = _rhs202;
          _lhs130.f8 = _rhs203;
          (_this).f21 = !(_module.__default.fm1((_this).f22, _1113_v24, globalState));
          (globalState).f8 = !(!((_this).f22));
          let _1137_v43;
          _1137_v43 = _dafny.Set.fromElements(_1113_v24);
          (globalState).f8 = (new BigNumber((_1137_v43).length)).isLessThan(_1110_cf6);
        }
        let _1138_v44;
        _1138_v44 = _module.D2.create_DC8(_1110_cf6, _1113_v24);
        _1089_v15 = _module.__default.safeModuloInt(((_1138_v44).dtor_cf15).multipliedBy(new BigNumber(-533)), _1089_v15);
        let _1139_v45;
        let _nw201 = Array((new BigNumber(25)).toNumber()).fill(false);
        _1139_v45 = _nw201;
        let _1140_v46;
        _1140_v46 = _dafny.Seq.of(_1139_v45, _1139_v45, _1139_v45);
        _1139_v45 = (_1140_v46)[_module.__default.safeIndex(_1089_v15, new BigNumber((_1140_v46).length))];
      } else {
        let _1141___mcc_h7 = (_source11).cf0;
        let _1142_cf0 = _1141___mcc_h7;
        let _1143_v47;
        _1143_v47 = _dafny.Set.fromElements(_this.f21);
        let _1144_v48;
        _1144_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements((_this).f22));
        let _1145_v49;
        _1145_v49 = _dafny.MultiSet.fromElements(_1143_v47, (((_1144_v48).contains((_this).f22)) ? ((_1144_v48).get((_this).f22)) : (_dafny.Set.fromElements(!((_this).f22), p0))));
        _1145_v49 = (_1145_v49).Difference(_1145_v49);
        if (p0) {
          (globalState).f9 = _1142_cf0;
          let _1146_v50;
          let _init33 = ((_1147_p0) => function (_1148_i5) {
            return _1147_p0;
          })(p0);
          let _nw202 = Array((new BigNumber(12)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw202.length); _i0_33++) {
            _nw202[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1146_v50 = _nw202;
          let _index176 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1146_v50).length));
          (_1146_v50)[_index176] = (_this).f22;
          let _1149_v51;
          _1149_v51 = _module.D2.create_DC8((_dafny.ZERO).minus(_1089_v15), new _dafny.CodePoint('k'.codePointAt(0)));
          let _1150_v52;
          _1150_v52 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1151_v53;
          _1151_v53 = _dafny.MultiSet.fromElements(_1142_cf0, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_module.__default.fm1(p0, _1150_v52, globalState))).length)), (_this).fm47(_1145_v49, globalState));
          let _1152_v54;
          _1152_v54 = _dafny.Seq.of(_1089_v15, _1089_v15, new BigNumber((_1151_v53).cardinality()));
          let _index177 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1146_v50).length));
          (_1146_v50)[_index177] = ((_1149_v51).dtor_cf15).isEqualTo((_1152_v54)[_module.__default.safeIndex(_1142_cf0, new BigNumber((_1152_v54).length))]);
          _1089_v15 = (_dafny.ZERO).minus(_1142_cf0);
          let _1153_v55;
          _1153_v55 = _1151_v53;
          let _1154_v56;
          let _nw203 = new _module.C1();
          _nw203.__ctor(p0, !(_this.f21), _module.D0.create_DC0(_1089_v15), ((_1153_v55)).IsProperSubsetOf(_1151_v53));
          _1154_v56 = _nw203;
          _1089_v15 = (_1142_cf0).plus(_module.__default.safeModuloInt(_1089_v15, _1089_v15));
        } else {
          let _1155_v57;
          let _nw204 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1155_v57 = _nw204;
          let _1156_v58;
          _1156_v58 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f22) || (p0),_1155_v57);
          _1156_v58 = (_1156_v58).update(p0, _1155_v57);
          (_this).f21 = (_this).f22;
          let _1157_v59;
          let _init34 = function (_1158_i6) {
            return false;
          };
          let _nw205 = Array((new BigNumber(18)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw205.length); _i0_34++) {
            _nw205[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1157_v59 = _nw205;
          let _1159_v60;
          _1159_v60 = _dafny.Map.Empty.slice().updateUnsafe(_this.f21,_1157_v59);
          let _1160_v61;
          _1160_v61 = _dafny.Map.Empty.slice().updateUnsafe((((_1159_v60).contains(p0)) ? ((_1159_v60).get(p0)) : (_1157_v59)),new BigNumber(42));
          let _1161_v62;
          _1161_v62 = _dafny.Seq.of((_dafny.ZERO).minus(_1142_cf0));
          _1160_v61 = (_1160_v61).update(((!((_this).f22)) ? (_1157_v59) : ((((_1159_v60).contains((_this).f22)) ? ((_1159_v60).get((_this).f22)) : (_1157_v59)))), new BigNumber((_1161_v62).length));
          let _1162_v63;
          _1162_v63 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1163_v64;
          _1163_v64 = _dafny.Seq.of(_module.__default.fm1(_this.f21, _1162_v63, globalState), _this.f21);
          let _1164_v65;
          _1164_v65 = _dafny.MultiSet.fromElements(_1089_v15);
          let _1165_v66;
          _1165_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(216),_1142_cf0);
          let _1166_v67;
          _1166_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_module.__default.fm2(_1142_cf0, _1164_v65, _1165_v66, globalState));
          let _1167_v68;
          _1167_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_1166_v67);
          let _1168_v69;
          _1168_v69 = _dafny.Seq.UnicodeFromString("higtskjfd");
          let _1169_v70;
          _1169_v70 = _module.D9.create_DC22(new BigNumber((_dafny.Seq.Concat(_1163_v64, _1163_v64)).length), (((_1167_v68).contains(new BigNumber(-276))) ? ((_1167_v68).get(new BigNumber(-276))) : ((_1166_v67).update(new BigNumber((_1163_v64).length), _1168_v69))));
          _1169_v70 = _1169_v70;
          _1142_cf0 = (_dafny.ZERO).minus(_1089_v15);
        }
        let _1170_v71;
        _1170_v71 = _dafny.Seq.UnicodeFromString("bub");
        if ((new BigNumber((_dafny.Seq.Concat(_1170_v71, _dafny.Seq.UnicodeFromString("y"))).length)).isEqualTo((_dafny.ZERO).minus((_1142_cf0).plus(_module.__default.fm0((_this).f22, globalState))))) {
          let _1171_v72;
          _1171_v72 = _dafny.Seq.of(_1142_cf0);
          let _1172_v73;
          _1172_v73 = _dafny.Set.fromElements(_1171_v72);
          let _1173_v74;
          let _nw206 = new _module.C0();
          _nw206.__ctor(!((_1172_v73).IsSubsetOf(_1172_v73)));
          _1173_v74 = _nw206;
          _1171_v72 = (_module.D5.create_DC13(_dafny.Seq.of(_1142_cf0))).dtor_cf25;
          let _1174_v75;
          let _init35 = ((_1175_cf0) => function (_1176_i7) {
            return (_1176_i7).plus(_1175_cf0);
          })(_1142_cf0);
          let _nw207 = Array((new BigNumber(29)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw207.length); _i0_35++) {
            _nw207[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1174_v75 = _nw207;
          let _index178 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_1174_v75).length));
          (_1174_v75)[_index178] = (_dafny.ZERO).minus(_1089_v15);
          let _1177_v76;
          _1177_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_1089_v15);
          let _index179 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_1174_v75).length));
          let _rhs204 = _module.__default.safeDivisionInt(new BigNumber((_1177_v76).length), _1089_v15);
          let _rhs205 = !((_1173_v74).f17);
          let _lhs131 = _1174_v75;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_1174_v75).length));
          let _lhs133 = _this;
          _lhs131[_lhs132] = _rhs204;
          _lhs133.f21 = _rhs205;
          let _index180 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_1174_v75).length));
          (_1174_v75)[_index180] = (_1174_v75)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_1174_v75).length))];
          let _1178_v77;
          let _init36 = ((_1179_v71) => function (_1180_i8) {
            return _dafny.Seq.Concat(_1179_v71, _1179_v71);
          })(_1170_v71);
          let _nw208 = Array((new BigNumber(4)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw208.length); _i0_36++) {
            _nw208[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1178_v77 = _nw208;
          let _index181 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1178_v77).length));
          (_1178_v77)[_index181] = _1170_v71;
          let _1181_v78;
          _1181_v78 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), function (_1182_i9) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }));
          let _index182 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1178_v77).length));
          (_1178_v77)[_index182] = (((_1089_v15).isLessThanOrEqualTo(_1142_cf0)) ? ((_1181_v78)[_module.__default.safeIndex(_1142_cf0, new BigNumber((_1181_v78).length))]) : (_1170_v71));
        } else {
          let _1183_v79;
          _1183_v79 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
          let _1184_v80;
          _1184_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1183_v79,(_this).f22);
          let _1185_v82;
          let _nw209 = Array((new BigNumber(8)).toNumber());
          _nw209[(_dafny.ZERO).toNumber()] = _1184_v80;
          _nw209[(_dafny.ONE).toNumber()] = _1184_v80;
          _nw209[(new BigNumber(2)).toNumber()] = _1184_v80;
          _nw209[(new BigNumber(3)).toNumber()] = _1184_v80;
          _nw209[(new BigNumber(4)).toNumber()] = _1184_v80;
          _nw209[(new BigNumber(5)).toNumber()] = _1184_v80;
          _nw209[(new BigNumber(6)).toNumber()] = _1184_v80;
          _nw209[(new BigNumber(7)).toNumber()] = function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of ((_dafny.Map.Empty.slice().updateUnsafe(_1183_v79,p0)).update(_1183_v79, _this.f21)).Keys.Elements) {
              let _1186_v81 = _compr_48;
              if (((_dafny.Map.Empty.slice().updateUnsafe(_1183_v79,p0)).update(_1183_v79, _this.f21)).contains(_1186_v81)) {
                _coll48.push([_1186_v81,!((_this).f22)]);
              }
            }
            return _coll48;
          }();
          _1185_v82 = _nw209;
          let _1187_v84;
          _1187_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1183_v79,_1142_cf0);
          let _index183 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1185_v82).length));
          (_1185_v82)[_index183] = (_1184_v80).Merge(function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_1187_v84).Keys.Elements) {
              let _1188_v83 = _compr_49;
              if ((_1187_v84).contains(_1188_v83)) {
                _coll49.push([_1188_v83,p0]);
              }
            }
            return _coll49;
          }());
          let _1189_v85;
          let _nw210 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1189_v85 = _nw210;
          let _1190_v86;
          _1190_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_this.f21);
          let _index184 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1189_v85).length));
          (_1189_v85)[_index184] = (((_1190_v86).contains(new BigNumber(264))) ? ((_1190_v86).get(new BigNumber(264))) : (true));
          let _1191_v87;
          _1191_v87 = new _dafny.CodePoint('o'.codePointAt(0));
          let _index185 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1189_v85).length));
          (_1189_v85)[_index185] = _module.__default.fm1(p0, _1191_v87, globalState);
          let _1192_v89;
          let _init37 = ((_1193_v79) => function (_1194_i10) {
            return function () {
              let _coll50 = new _dafny.Map();
              for (const _compr_50 of _dafny.IntegerRange(new BigNumber(921), new BigNumber(224))) {
                let _1195_v88 = _compr_50;
                if (((new BigNumber(921)).isLessThanOrEqualTo(_1195_v88)) && ((_1195_v88).isLessThan(new BigNumber(224)))) {
                  _coll50.push([(_1195_v88).multipliedBy(new BigNumber(42)),_1193_v79]);
                }
              }
              return _coll50;
            }();
          })(_1183_v79);
          let _nw211 = Array((new BigNumber(6)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw211.length); _i0_37++) {
            _nw211[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1192_v89 = _nw211;
          let _1196_v90;
          _1196_v90 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm47(_dafny.MultiSet.fromElements(_1143_v47), globalState),_1183_v79);
          let _index186 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_1192_v89).length));
          (_1192_v89)[_index186] = _1196_v90;
          let _1197_v91;
          let _nw212 = Array((new BigNumber(12)).toNumber()).fill(_module.D2.Default());
          _1197_v91 = _nw212;
          let _1198_v92;
          _1198_v92 = _dafny.MultiSet.fromElements(_1197_v91);
          let _1199_v93;
          _1199_v93 = _dafny.Seq.of(new BigNumber((_1198_v92).cardinality()));
          let _index187 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1185_v82).length));
          let _index188 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1189_v85).length));
          let _index189 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1189_v85).length));
          let _index190 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_1192_v89).length));
          let _rhs206 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,true),(_1089_v15).isLessThanOrEqualTo(_module.__default.fm0(_this.f21, globalState)));
          let _rhs207 = _dafny.Seq.IsPrefixOf(_1199_v93, _1199_v93);
          let _rhs208 = (new BigNumber(542)).isEqualTo(new BigNumber((_1170_v71).length));
          let _rhs209 = !(!(_this.f21));
          let _rhs210 = ((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_dafny.Map.Empty.slice().updateUnsafe(p0,true))) : (_1196_v90));
          let _lhs134 = _1185_v82;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1185_v82).length));
          let _lhs136 = globalState;
          let _lhs137 = _1189_v85;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1189_v85).length));
          let _lhs139 = _1189_v85;
          let _lhs140 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1189_v85).length));
          let _lhs141 = _1192_v89;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_1192_v89).length));
          _lhs134[_lhs135] = _rhs206;
          _lhs136.f8 = _rhs207;
          _lhs137[_lhs138] = _rhs208;
          _lhs139[_lhs140] = _rhs209;
          _lhs141[_lhs142] = _rhs210;
          let _1200_v94;
          _1200_v94 = _dafny.Seq.of(_1143_v47, _1143_v47);
          let _1201_v95;
          _1201_v95 = _dafny.MultiSet.fromElements((_this).fm47(_dafny.MultiSet.FromArray(_1200_v94), globalState));
          let _1202_v97;
          _1202_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v15,new BigNumber((function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(559), new BigNumber(914))) {
              let _1203_v96 = _compr_51;
              if (((new BigNumber(559)).isLessThanOrEqualTo(_1203_v96)) && ((_1203_v96).isLessThan(new BigNumber(914)))) {
                _coll51.push([(_1203_v96).multipliedBy((_dafny.ZERO).minus(_1089_v15)),_1142_cf0]);
              }
            }
            return _coll51;
          }()).length));
          let _1204_v98;
          let _nw213 = new _module.C0();
          _nw213.__ctor(_this.f21);
          _1204_v98 = _nw213;
          let _1205_v99;
          _1205_v99 = _dafny.MultiSet.fromElements(_1204_v98);
          let _1206_v100;
          _1206_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1205_v99);
          let _1207_v101;
          let _nw214 = Array((new BigNumber(18)).toNumber());
          _nw214[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), ((_1208_v87) => function (_1209_i11) {
            return _1208_v87;
          })(_1191_v87));
          _nw214[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yygsnu"), _1170_v71);
          _nw214[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1170_v71, _1170_v71);
          _nw214[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1170_v71, _dafny.Seq.update(_module.__default.fm2((_dafny.ZERO).minus(_1142_cf0), _1201_v95, _1202_v97, globalState), _module.__default.safeIndex(_1089_v15, new BigNumber((_module.__default.fm2((_dafny.ZERO).minus(_1142_cf0), _1201_v95, _1202_v97, globalState)).length)), new _dafny.CodePoint('g'.codePointAt(0)))), _module.__default.safeIndex(new BigNumber((_1206_v100).length), new BigNumber((_dafny.Seq.Concat(_1170_v71, _dafny.Seq.update(_module.__default.fm2((_dafny.ZERO).minus(_1142_cf0), _1201_v95, _1202_v97, globalState), _module.__default.safeIndex(_1089_v15, new BigNumber((_module.__default.fm2((_dafny.ZERO).minus(_1142_cf0), _1201_v95, _1202_v97, globalState)).length)), new _dafny.CodePoint('g'.codePointAt(0))))).length)), _1191_v87);
          _nw214[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("wquijjw");
          _nw214[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1170_v71, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-538)), ((_1210_v87) => function (_1211_i12) {
            return _1210_v87;
          })(_1191_v87)));
          _nw214[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-476)), ((_1212_v87) => function (_1213_i13) {
            return _1212_v87;
          })(_1191_v87));
          _nw214[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1170_v71, _dafny.Seq.UnicodeFromString("gmheovam"));
          _nw214[(new BigNumber(8)).toNumber()] = _1170_v71;
          _nw214[(new BigNumber(9)).toNumber()] = _module.__default.fm2(new BigNumber(589), _1201_v95, _1202_v97, globalState);
          _nw214[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1170_v71, _1170_v71);
          _nw214[(new BigNumber(11)).toNumber()] = _1170_v71;
          _nw214[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-550)), ((_1214_v87) => function (_1215_i14) {
            return _1214_v87;
          })(_1191_v87)), _1170_v71);
          _nw214[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_1170_v71, _module.__default.safeIndex(new BigNumber((_1170_v71).length), new BigNumber((_1170_v71).length)), (_1170_v71)[_module.__default.safeIndex(_1089_v15, new BigNumber((_1170_v71).length))]);
          _nw214[(new BigNumber(14)).toNumber()] = _1170_v71;
          _nw214[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1170_v71, _dafny.Seq.update(_1170_v71, _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1170_v71).length)), (_1170_v71)[_module.__default.safeIndex(_1089_v15, new BigNumber((_1170_v71).length))]));
          _nw214[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("hrfp");
          _nw214[(new BigNumber(17)).toNumber()] = _1170_v71;
          _1207_v101 = _nw214;
          let _1216_v102;
          _1216_v102 = _dafny.Seq.of((_1189_v85)[_module.__default.safeIndex(new BigNumber(110), new BigNumber((_1189_v85).length))]);
          let _1217_v103;
          _1217_v103 = _module.D7.create_DC16(_1216_v102);
          let _1218_v104;
          _1218_v104 = _dafny.Seq.of(_1217_v103);
          let _1219_v105;
          _1219_v105 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_1142_cf0, new BigNumber((_1218_v104).length)),_1207_v101);
          _1207_v101 = (((_1219_v105).contains(_module.__default.safeModuloInt(_1142_cf0, _1142_cf0))) ? ((_1219_v105).get(_module.__default.safeModuloInt(_1142_cf0, _1142_cf0))) : (((p0) ? (_1207_v101) : (_1207_v101))));
          r0 = _module.__default.fm0(!(false), globalState);
          let _1220_v106;
          _1220_v106 = _module.D2.create_DC7(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(320)), ((_1221_v87) => function (_1222_i15) {
  return _1221_v87;
})(_1191_v87))).length), !((_this).f22), _1170_v71);
          let _1223_v107;
          _1223_v107 = _dafny.Map.Empty.slice().updateUnsafe((((_1189_v85)[_module.__default.safeIndex(new BigNumber(110), new BigNumber((_1189_v85).length))]) ? (_1202_v97) : ((_1202_v97).update(_1089_v15, new BigNumber((_1216_v102).length)))),(_1220_v106).dtor_cf12);
          _1223_v107 = ((_1223_v107).Merge(function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_dafny.Set.fromElements(_1202_v97, _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_1089_v15))).Elements) {
              let _1224_v108 = _compr_52;
              if ((_dafny.Set.fromElements(_1202_v97, _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_1089_v15))).contains(_1224_v108)) {
                _coll52.push([_1224_v108,new BigNumber((_1201_v95).cardinality())]);
              }
            }
            return _coll52;
          }())).Merge(_1223_v107);
          let _1225_v109;
          let _nw215 = new _module.C1();
          _nw215.__ctor(_this.f21, false, _1090_v16, _this.f21);
          _1225_v109 = _nw215;
        }
        let _1226_v110;
        _1226_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1142_cf0,_1089_v15);
        r0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(((false) ? (_1089_v15) : (new BigNumber((_1226_v110).length))), (_dafny.ZERO).minus(_1142_cf0)));
      }
      let _1227_v111;
      let _init38 = ((_1228_v15) => function (_1229_i17) {
        return (_1229_i17).minus(_1228_v15);
      })(_1089_v15);
      let _nw216 = Array((new BigNumber(15)).toNumber());
      for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw216.length); _i0_38++) {
        _nw216[_i0_38] = _init38(new BigNumber(_i0_38));
      }
      _1227_v111 = _nw216;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1227_v111).length))) {
        let _1230_i16 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1230_i16)) && ((_1230_i16).isLessThan(new BigNumber((_1227_v111).length))))) {
          (_1227_v111)[(_1230_i16)] = (_1230_i16).plus((_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(591)).minus(new BigNumber(-626)))));
        }
      }
      let _1231_v112;
      let _nw217 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
      _1231_v112 = _nw217;
      let _1232_v113;
      _1232_v113 = _module.D10.create_DC23(_1231_v112);
      _1231_v112 = (_1232_v113).dtor_cf42;
      let _1233_v114;
      _1233_v114 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
      let _1234_v115;
      _1234_v115 = _dafny.Seq.of(_1233_v114);
      let _1235_v116;
      _1235_v116 = new _dafny.CodePoint('y'.codePointAt(0));
      let _hi9 = new BigNumber((_module.__default.fm50(_1235_v116, true, (_this).f22, globalState)).length);
      for (let _1236_i18 = _module.__default.safeModuloInt(new BigNumber((_1234_v115).length), new BigNumber(-112)); _1236_i18.isLessThan(_hi9); _1236_i18 = _1236_i18.plus(_dafny.ONE)) {
        let _1237_v117;
        _1237_v117 = _dafny.Set.fromElements(new BigNumber(-535), new BigNumber(-722));
        let _1238_v118;
        _1238_v118 = _module.D8.create_DC18((_1237_v117).Difference(_1237_v117));
        let _source12 = _1238_v118;
        if (_source12.is_DC19) {
          let _1239___mcc_h8 = (_source12).cf34;
          let _1240___mcc_h9 = (_source12).cf35;
          let _1241_cf35 = _1240___mcc_h9;
          let _1242_cf34 = _1239___mcc_h8;
          _1241_cf35 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1241_cf35,new _dafny.CodePoint('s'.codePointAt(0)))).update(_1236_i18, _1235_v116)).length);
          let _1243_v119;
          _1243_v119 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("y"), _module.__default.safeIndex(_1236_i18, new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)), _1235_v116));
          _1243_v119 = (_1243_v119).update(_this.f21, _dafny.Seq.UnicodeFromString("lxhumbbwu"));
          let _1244_v120;
          _1244_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1090_v16,new _dafny.CodePoint('x'.codePointAt(0)));
          _1244_v120 = ((_module.__default.fm51(globalState)).Merge(_1244_v120)).update(_1090_v16, _module.__default.fm3(_1241_cf35, (_this).f22, globalState));
          (globalState).f9 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(_1236_i18))).plus(_1089_v15);
        } else if (_source12.is_DC20) {
          let _1245___mcc_h10 = (_source12).cf36;
          let _1246___mcc_h11 = (_source12).cf37;
          let _1247___mcc_h12 = (_source12).cf38;
          let _1248_cf38 = _1247___mcc_h12;
          let _1249_cf37 = _1246___mcc_h11;
          let _1250_cf36 = _1245___mcc_h10;
          (globalState).f9 = _1089_v15;
          r0 = _1236_i18;
          let _1251_v121;
          _1251_v121 = _dafny.MultiSet.fromElements(_1249_cf37);
          (globalState).f8 = !((((_1233_v114).contains((_1251_v121).IsSubsetOf(_1251_v121))) ? ((_1233_v114).get((_1251_v121).IsSubsetOf(_1251_v121))) : ((_1236_i18).isLessThanOrEqualTo(_1089_v15))));
          let _1252_v122;
          _1252_v122 = _module.D8.create_DC20(_1250_cf36, _1249_cf37, p0);
          let _rhs211 = _1252_v122;
          let _rhs212 = _1235_v116;
          _1252_v122 = _rhs211;
          _1235_v116 = _rhs212;
        } else {
          let _1253___mcc_h13 = (_source12).cf33;
          let _1254_cf33 = _1253___mcc_h13;
          _module.__default.m0(globalState);
          let _1255_v123;
          _1255_v123 = _dafny.Seq.UnicodeFromString("ses");
          (globalState).f9 = ((new BigNumber((_1255_v123).length)).multipliedBy(_1089_v15)).plus(new BigNumber(836));
          let _1256_v124;
          _1256_v124 = _dafny.Set.fromElements(_1255_v123, _1255_v123, _1255_v123);
          (_this).f21 = !(_module.__default.fm1((_1256_v124).IsDisjointFrom(_dafny.Set.fromElements(_1255_v123)), _1235_v116, globalState));
          let _1257_v125;
          _1257_v125 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v15,(_dafny.ZERO).minus((_dafny.ZERO).minus((_1236_i18).minus(_1089_v15))));
          r0 = new BigNumber((_1257_v125).length);
        }
        (_this).f21 = (_this).f22;
        (globalState).f9 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1236_i18), _1089_v15)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), ((_1258_v116) => function (_1259_i19) {
          return _1258_v116;
        })(_1235_v116))).length));
        (globalState).f8 = _this.f21;
      }
      let _1260_v126;
      _1260_v126 = _dafny.Seq.UnicodeFromString("taiis");
      let _1261_v127;
      _1261_v127 = _module.D2.create_DC7(_this.f21, _1089_v15, p0, _1260_v126);
      let _1262_v128;
      _1262_v128 = _dafny.Set.fromElements((_1261_v127).dtor_cf11);
      let _1263_v129;
      _1263_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1262_v128,_1089_v15);
      if ((_1263_v129).contains(_1262_v128)) {
        r0 = _1089_v15;
        let _1264_v130;
        _1264_v130 = _dafny.Set.fromElements(new BigNumber(-109));
        let _1265_v131;
        _1265_v131 = _module.D8.create_DC18(_1264_v130);
        let _1266_v132;
        _1266_v132 = _dafny.Map.Empty.slice().updateUnsafe(_1265_v131,(_this).f22);
        let _1267_v133;
        _1267_v133 = _dafny.Seq.of(_1266_v132, _dafny.Map.Empty.slice().updateUnsafe(_1265_v131,(_this).f22), _1266_v132, _1266_v132, _1266_v132);
        let _1268_v134;
        _1268_v134 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f22),_1089_v15);
        let _1269_v135;
        _1269_v135 = _dafny.MultiSet.fromElements(_1262_v128, _1262_v128, _1262_v128);
        let _nw218 = Array((new BigNumber(17)).toNumber());
        _nw218[(_dafny.ZERO).toNumber()] = _1089_v15;
        _nw218[(_dafny.ONE).toNumber()] = new BigNumber((_1267_v133).length);
        _nw218[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(722), _1089_v15);
        _nw218[(new BigNumber(3)).toNumber()] = _1089_v15;
        _nw218[(new BigNumber(4)).toNumber()] = new BigNumber((_1260_v126).length);
        _nw218[(new BigNumber(5)).toNumber()] = _1089_v15;
        _nw218[(new BigNumber(6)).toNumber()] = new BigNumber(241);
        _nw218[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1268_v134).length), _1089_v15);
        _nw218[(new BigNumber(8)).toNumber()] = (_1089_v15).minus(_1089_v15);
        _nw218[(new BigNumber(9)).toNumber()] = _1089_v15;
        _nw218[(new BigNumber(10)).toNumber()] = (_this).fm47(_1269_v135, globalState);
        _nw218[(new BigNumber(11)).toNumber()] = _1089_v15;
        _nw218[(new BigNumber(12)).toNumber()] = _1089_v15;
        _nw218[(new BigNumber(13)).toNumber()] = ((p0) ? (_1089_v15) : (_1089_v15));
        _nw218[(new BigNumber(14)).toNumber()] = _1089_v15;
        _nw218[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("b")).length);
        _nw218[(new BigNumber(16)).toNumber()] = _1089_v15;
        _1227_v111 = _nw218;
        let _1270_v136;
        _1270_v136 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-248),_1260_v126);
        let _1271_v137;
        _1271_v137 = _module.D9.create_DC22(_1089_v15, _1270_v136);
        let _1272_v138;
        _1272_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v15,_1271_v137);
        _1272_v138 = (_1272_v138).update(_1089_v15, _1271_v137);
        let _1273_v139;
        _1273_v139 = _dafny.MultiSet.fromElements(!(p0));
        let _1274_v140;
        _1274_v140 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v15,(_1273_v139).IsDisjointFrom(_1273_v139));
        _1274_v140 = (_1274_v140).update(((p0) ? (_1089_v15) : (new BigNumber((_module.__default.fm52(globalState)).length))), _this.f21);
        let _1275_v141;
        _1275_v141 = _module.D4.create_DC12((_this).f22, (_this).f22, p0);
        let _1276_v142;
        let _nw219 = Array((new BigNumber(11)).toNumber());
        _nw219[(_dafny.ZERO).toNumber()] = (_this).f22;
        _nw219[(_dafny.ONE).toNumber()] = p0;
        _nw219[(new BigNumber(2)).toNumber()] = false;
        _nw219[(new BigNumber(3)).toNumber()] = (_this).f22;
        _nw219[(new BigNumber(4)).toNumber()] = p0;
        _nw219[(new BigNumber(5)).toNumber()] = _this.f21;
        _nw219[(new BigNumber(6)).toNumber()] = (_1275_v141).dtor_cf22;
        _nw219[(new BigNumber(7)).toNumber()] = p0;
        _nw219[(new BigNumber(8)).toNumber()] = (_this).f22;
        _nw219[(new BigNumber(9)).toNumber()] = (_this).f22;
        _nw219[(new BigNumber(10)).toNumber()] = p0;
        _1276_v142 = _nw219;
        let _1277_v143;
        _1277_v143 = _module.D7.create_DC17(_1260_v126, _1276_v142, p0, new BigNumber((_1262_v128).length));
        let _1278_v144;
        let _nw220 = Array((new BigNumber(11)).toNumber());
        _nw220[(_dafny.ZERO).toNumber()] = (_1277_v143).dtor_cf30;
        _nw220[(_dafny.ONE).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(2)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(3)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(4)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(5)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(6)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(7)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(8)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(9)).toNumber()] = _1276_v142;
        _nw220[(new BigNumber(10)).toNumber()] = _1276_v142;
        _1278_v144 = _nw220;
        let _1279_v145;
        _1279_v145 = _dafny.Map.Empty.slice().updateUnsafe(_1278_v144,_module.__default.safeDivisionInt(_1089_v15, new BigNumber(512)));
        _1279_v145 = (_1279_v145).update(_1278_v144, _module.__default.fm0((_this).f22, globalState));
      } else {
        (_this).f21 = true;
        let _1280_v146;
        _1280_v146 = _dafny.MultiSet.fromElements(_this.f21);
        let _1281_v147;
        _1281_v147 = _dafny.Seq.of(_this.f21, !(p0), true);
        let _1282_v148;
        _1282_v148 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_1089_v15);
        if (!((new BigNumber((_1280_v146).cardinality())).isEqualTo((new BigNumber((_1281_v147).length)).minus(new BigNumber((_1282_v148).length))))) {
          (globalState).f9 = new BigNumber(367);
          let _1283_v149;
          _1283_v149 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1280_v146).Intersect(_1280_v146)).cardinality()),(_1089_v15).multipliedBy(_module.__default.fm0((_this).f22, globalState)));
          let _1284_v150;
          _1284_v150 = _dafny.Seq.of(_1227_v111);
          let _1285_v151;
          _1285_v151 = _module.D9.create_DC21((_1284_v150)[_module.__default.safeIndex(_1089_v15, new BigNumber((_1284_v150).length))]);
          let _1286_v152;
          _1286_v152 = _dafny.Seq.of(_1227_v111, (_1285_v151).dtor_cf39, _1227_v111, _1227_v111, _1227_v111);
          _1283_v149 = (_1283_v149).update(new BigNumber((_1284_v150).length), new BigNumber(((_1233_v114).update(false, _this.f21)).length));
          let _1287_v153;
          let _nw221 = new _module.C0();
          _nw221.__ctor(_this.f21);
          _1287_v153 = _nw221;
          let _1288_v154;
          let _nw222 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1288_v154 = _nw222;
          let _index191 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1288_v154).length));
          (_1288_v154)[_index191] = p0;
          let _index192 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1288_v154).length));
          let _rhs213 = new BigNumber((_dafny.Seq.UnicodeFromString("nbmrqg")).length);
          let _rhs214 = (_1287_v153).f17;
          let _rhs215 = (((_this).f22) ? (_1288_v154) : (_1288_v154));
          let _rhs216 = (_this).f22;
          let _lhs143 = globalState;
          let _lhs144 = _1288_v154;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1288_v154).length));
          let _lhs146 = globalState;
          _lhs143.f9 = _rhs213;
          _lhs144[_lhs145] = _rhs214;
          _1288_v154 = _rhs215;
          _lhs146.f8 = _rhs216;
          let _1289_v155;
          _1289_v155 = _module.D0.create_DC2(_dafny.MultiSet.FromArray(_1281_v147), false, new BigNumber(-678), (_this).f22);
          let _1290_v156;
          _1290_v156 = _module.D4.create_DC12((_1289_v155).dtor_cf5, (_this).f22, (_this).f22);
          (globalState).f8 = (_1290_v156).dtor_cf24;
        } else {
          let _index193 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_1227_v111).length));
          (_1227_v111)[_index193] = new BigNumber(((_1233_v114).update(!((_this).f22), (_this).f22)).length);
          let _index194 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_1227_v111).length));
          (_1227_v111)[_index194] = new BigNumber(711);
          (globalState).f8 = (_this).f22;
          let _1291_v157;
          let _nw223 = new _module.C0();
          _nw223.__ctor(!(p0));
          _1291_v157 = _nw223;
          _1089_v15 = _1089_v15;
          (_this).f21 = (_1291_v157).f17;
        }
        let _1292_v158;
        let _nw224 = new _module.C2();
        _nw224.__ctor((_this).f11);
        _1292_v158 = _nw224;
        r0 = _1089_v15;
        let _1293_v159;
        _1293_v159 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v15,_dafny.Seq.IsPrefixOf(_1260_v126, _1260_v126));
        _1293_v159 = _1293_v159;
      }
      r0 = (_dafny.ZERO).minus(_1089_v15);
      r1 = _module.__default.fm53(globalState);
      return [r0, r1];
    }
    m12(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _1294_v0;
      _1294_v0 = _dafny.Seq.of(!(false));
      let _1295_v1;
      let _nw225 = Array((new BigNumber(5)).toNumber());
      _nw225[(_dafny.ZERO).toNumber()] = (_this).f22;
      _nw225[(_dafny.ONE).toNumber()] = (true) && (p1);
      _nw225[(new BigNumber(2)).toNumber()] = (_this).f22;
      _nw225[(new BigNumber(3)).toNumber()] = _dafny.areEqual(_dafny.Seq.of(_this.f21), _1294_v0);
      _nw225[(new BigNumber(4)).toNumber()] = p0;
      _1295_v1 = _nw225;
      let _1296_v2;
      _1296_v2 = _dafny.Seq.UnicodeFromString("bhgv");
      let _1297_v3;
      _1297_v3 = new BigNumber(485);
      let _1298_v4;
      _1298_v4 = new _dafny.CodePoint('l'.codePointAt(0));
      let _index195 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length));
      (_1295_v1)[_index195] = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_1296_v2, _module.__default.safeIndex(_1297_v3, new BigNumber((_1296_v2).length)), _1298_v4), _1296_v2);
      let _index196 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length));
      (_1295_v1)[_index196] = _module.__default.fm1(p1, _1298_v4, globalState);
      let _1299_v5;
      _1299_v5 = _dafny.Set.fromElements(p1, _this.f21);
      let _1300_v6;
      _1300_v6 = _dafny.Seq.of(new BigNumber((_1299_v5).length));
      let _hi10 = _1297_v3;
      for (let _1301_i0 = (_1300_v6)[_module.__default.safeIndex(_1297_v3, new BigNumber((_1300_v6).length))]; _1301_i0.isLessThan(_hi10); _1301_i0 = _1301_i0.plus(_dafny.ONE)) {
        let _1302_v7;
        _1302_v7 = _dafny.Set.fromElements(_1295_v1);
        _1302_v7 = (_1302_v7).Intersect(_1302_v7);
        if (p1) {
          let _1303_v8;
          _1303_v8 = _dafny.Set.fromElements(_1301_i0);
          let _1304_v9;
          _1304_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1303_v8,_this.f21);
          _1297_v3 = new BigNumber(((_1304_v9).Merge(_1304_v9)).length);
          let _index197 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length));
          (_1295_v1)[_index197] = p0;
          let _1305_v10;
          let _init39 = ((_1306_v3) => function (_1307_i1) {
            return (_1307_i1).minus(_1306_v3);
          })(_1297_v3);
          let _nw226 = Array((new BigNumber(15)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw226.length); _i0_39++) {
            _nw226[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1305_v10 = _nw226;
          let _1308_v11;
          _1308_v11 = _dafny.MultiSet.fromElements(_1305_v10);
          let _1309_v12;
          let _nw227 = Array((new BigNumber(18)).toNumber());
          _nw227[(_dafny.ZERO).toNumber()] = _1301_i0;
          _nw227[(_dafny.ONE).toNumber()] = _1297_v3;
          _nw227[(new BigNumber(2)).toNumber()] = _1301_i0;
          _nw227[(new BigNumber(3)).toNumber()] = new BigNumber(-484);
          _nw227[(new BigNumber(4)).toNumber()] = _1297_v3;
          _nw227[(new BigNumber(5)).toNumber()] = _1297_v3;
          _nw227[(new BigNumber(6)).toNumber()] = new BigNumber((_1300_v6).length);
          _nw227[(new BigNumber(7)).toNumber()] = _1301_i0;
          _nw227[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.of(_1297_v3)).length);
          _nw227[(new BigNumber(9)).toNumber()] = _1297_v3;
          _nw227[(new BigNumber(10)).toNumber()] = _1301_i0;
          _nw227[(new BigNumber(11)).toNumber()] = _1301_i0;
          _nw227[(new BigNumber(12)).toNumber()] = _1301_i0;
          _nw227[(new BigNumber(13)).toNumber()] = _1301_i0;
          _nw227[(new BigNumber(14)).toNumber()] = _1297_v3;
          _nw227[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_1301_i0);
          _nw227[(new BigNumber(16)).toNumber()] = (_module.D0.create_DC1(_1308_v11, _1301_i0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(939)), ((_1310_v4) => function (_1311_i2) {
  return _1310_v4;
})(_1298_v4))).length))).dtor_cf3;
          _nw227[(new BigNumber(17)).toNumber()] = _1301_i0;
          _1309_v12 = _nw227;
          let _1312_v13;
          _1312_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1297_v3,_1305_v10);
          _1312_v13 = (_1312_v13).update(new BigNumber(963), _1309_v12);
          let _1313_v14;
          _1313_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1298_v4,p1);
          let _1314_v15;
          _1314_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(865),(_dafny.Map.Empty.slice().updateUnsafe(_1298_v4,p0)).Merge(_1313_v14));
          _1314_v15 = (_1314_v15).update(_1297_v3, _1313_v14);
          (globalState).f9 = _1297_v3;
        } else {
          let _1315_v16;
          let _nw228 = new _module.C2();
          _nw228.__ctor((_this).f11);
          _1315_v16 = _nw228;
          let _index198 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length));
          (_1295_v1)[_index198] = p0;
          let _1316_v17;
          let _nw229 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1316_v17 = _nw229;
          let _1317_v18;
          let _nw230 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1317_v18 = _nw230;
          let _rhs217 = _1316_v17;
          let _rhs218 = _module.__default.safeDivisionInt(_1301_i0, _1301_i0);
          let _rhs219 = ((_module.__default.fm1(!(true), _1298_v4, globalState)) ? (_1317_v18) : (_1317_v18));
          let _lhs147 = globalState;
          _1316_v17 = _rhs217;
          _lhs147.f9 = _rhs218;
          _1317_v18 = _rhs219;
          let _1318_v19;
          let _nw231 = Array((new BigNumber(9)).toNumber());
          _nw231[(_dafny.ZERO).toNumber()] = p2;
          _nw231[(_dafny.ONE).toNumber()] = p2;
          _nw231[(new BigNumber(2)).toNumber()] = p2;
          _nw231[(new BigNumber(3)).toNumber()] = p2;
          _nw231[(new BigNumber(4)).toNumber()] = p2;
          _nw231[(new BigNumber(5)).toNumber()] = (p2).Merge(p2);
          _nw231[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f21,new BigNumber((_1300_v6).length));
          _nw231[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1295_v1)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length))],_1301_i0);
          _nw231[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1297_v3);
          _1318_v19 = _nw231;
          let _index199 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length));
          let _rhs220 = _1298_v4;
          let _rhs221 = (_this).f22;
          let _rhs222 = _1318_v19;
          let _rhs223 = ((_dafny.ZERO).minus(_1297_v3)).isLessThan(new BigNumber(99));
          let _lhs148 = globalState;
          let _lhs149 = _1295_v1;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length));
          _1298_v4 = _rhs220;
          _lhs148.f8 = _rhs221;
          _1318_v19 = _rhs222;
          _lhs149[_lhs150] = _rhs223;
          let _1319_v20;
          _1319_v20 = _module.D0.create_DC0(_1301_i0);
          let _1320_v21;
          let _nw232 = new _module.C1();
          _nw232.__ctor((_1295_v1)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_1295_v1).length))], _this.f21, _1319_v20, !(_this.f21) || (_this.f21));
          _1320_v21 = _nw232;
        }
        _module.__default.m0(globalState);
        (globalState).f9 = _1297_v3;
      }
      let _1321_v22;
      _1321_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1297_v3,_this.f21);
      let _1322_v23;
      let _nw233 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _1322_v23 = _nw233;
      let _1323_v24;
      _1323_v24 = _dafny.Set.fromElements(_1322_v23, _1322_v23);
      (globalState).f9 = new BigNumber(((((false) ? (_1321_v22) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(960),p1)))).update(new BigNumber(((_1323_v24).Intersect(_1323_v24)).length), p1)).length);
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1295_v1).length))) {
        let _1324_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1324_i3)) && ((_1324_i3).isLessThan(new BigNumber((_1295_v1).length))))) {
          (_1295_v1)[(_1324_i3)] = !(!((_1297_v3).isLessThan(_1297_v3)));
        }
      }
      let _1325_v25;
      let _nw234 = new _module.C0();
      _nw234.__ctor(true);
      _1325_v25 = _nw234;
      let _1326_v26;
      let _init40 = ((_1327_v3) => function (_1328_i4) {
        return ((true) ? (_dafny.Set.fromElements(_1327_v3)) : (_dafny.Set.fromElements(_1327_v3, new BigNumber(67), _1327_v3, _1327_v3)));
      })(_1297_v3);
      let _nw235 = Array((new BigNumber(3)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw235.length); _i0_40++) {
        _nw235[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1326_v26 = _nw235;
      let _1329_v27;
      _1329_v27 = _dafny.Set.fromElements(new BigNumber(-662), new BigNumber((_dafny.Seq.UnicodeFromString("d")).length), _1297_v3, _1297_v3, (_dafny.ZERO).minus(_1297_v3));
      let _index200 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1326_v26).length));
      (_1326_v26)[_index200] = _1329_v27;
      let _1330_v28;
      let _nw236 = Array((new BigNumber(13)).toNumber()).fill(_module.D2.Default());
      _1330_v28 = _nw236;
      let _1331_v29;
      _1331_v29 = _module.D11.create_DC25(_1330_v28);
      let _1332_v30;
      let _nw237 = Array((new BigNumber(22)).toNumber());
      _nw237[(_dafny.ZERO).toNumber()] = _1330_v28;
      _nw237[(_dafny.ONE).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(2)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(3)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(4)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(5)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(6)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(7)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(8)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(9)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(10)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(11)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(12)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(13)).toNumber()] = (_1331_v29).dtor_cf46;
      _nw237[(new BigNumber(14)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(15)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(16)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(17)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(18)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(19)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(20)).toNumber()] = _1330_v28;
      _nw237[(new BigNumber(21)).toNumber()] = _1330_v28;
      _1332_v30 = _nw237;
      let _1333_v31;
      _1333_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p0, _1298_v4, globalState),p1);
      let _1334_v32;
      _1334_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1333_v31).length),_1330_v28);
      let _1335_v33;
      _1335_v33 = _dafny.Seq.of((((_1334_v32).contains(_1297_v3)) ? ((_1334_v32).get(_1297_v3)) : (_1330_v28)), _1330_v28, _1330_v28);
      let _1336_v34;
      _1336_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1297_v3,new BigNumber(862));
      let _index201 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1332_v30).length));
      (_1332_v30)[_index201] = (_1335_v33)[_module.__default.safeIndex((((_1336_v34).contains(_1297_v3)) ? ((_1336_v34).get(_1297_v3)) : (_1297_v3)), new BigNumber((_1335_v33).length))];
      let _index202 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1326_v26).length));
      let _index203 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1332_v30).length));
      let _rhs224 = _dafny.Set.fromElements(_1297_v3, _1297_v3, _1297_v3);
      let _rhs225 = _1330_v28;
      let _rhs226 = _1297_v3;
      let _rhs227 = (_1297_v3).plus(_1297_v3);
      let _lhs151 = _1326_v26;
      let _lhs152 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1326_v26).length));
      let _lhs153 = _1332_v30;
      let _lhs154 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1332_v30).length));
      let _lhs155 = globalState;
      _lhs151[_lhs152] = _rhs224;
      _lhs153[_lhs154] = _rhs225;
      _lhs155.f9 = _rhs226;
      _1297_v3 = _rhs227;
      r0 = (_1325_v25).f17;
      return r0;
    }
    m13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let _1337_v0;
      let _nw238 = Array((new BigNumber(21)).toNumber()).fill([]);
      _1337_v0 = _nw238;
      let _1338_v1;
      _1338_v1 = _dafny.Seq.UnicodeFromString("gadrckqd");
      let _1339_v3;
      let _nw239 = Array((new BigNumber(16)).toNumber());
      _nw239[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(348)), function (_1340_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      });
      _nw239[(_dafny.ONE).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-276)), function (_1341_i1) {
        return (_module.D2.create_DC8(new BigNumber((function () {
  let _coll53 = new _dafny.Map();
  for (const _compr_53 of _dafny.IntegerRange(new BigNumber(-908), new BigNumber(962))) {
    let _1342_v2 = _compr_53;
    if (((new BigNumber(-908)).isLessThanOrEqualTo(_1342_v2)) && ((_1342_v2).isLessThan(new BigNumber(962)))) {
      _coll53.push([(_1342_v2).minus(new BigNumber(116)),new _dafny.CodePoint('e'.codePointAt(0))]);
    }
  }
  return _coll53;
}()).length), new _dafny.CodePoint('d'.codePointAt(0)))).dtor_cf16;
      });
      _nw239[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("longt");
      _nw239[(new BigNumber(4)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-677)), function (_1343_i2) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      });
      _nw239[(new BigNumber(6)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(7)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(8)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(9)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(10)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("sp");
      _nw239[(new BigNumber(12)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(13)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(14)).toNumber()] = _1338_v1;
      _nw239[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("nk");
      _1339_v3 = _nw239;
      let _index204 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1337_v0).length));
      (_1337_v0)[_index204] = _1339_v3;
      let _index205 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1337_v0).length));
      (_1337_v0)[_index205] = _1339_v3;
      let _1344_v4;
      _1344_v4 = _dafny.Seq.of(p3);
      let _rhs228 = _module.__default.fm1((_this).f22, new _dafny.CodePoint('h'.codePointAt(0)), globalState);
      let _rhs229 = new BigNumber((_1338_v1).length);
      let _rhs230 = false;
      let _rhs231 = _1344_v4;
      let _lhs156 = globalState;
      let _lhs157 = globalState;
      let _lhs158 = globalState;
      _lhs156.f8 = _rhs228;
      _lhs157.f9 = _rhs229;
      _lhs158.f8 = _rhs230;
      _1344_v4 = _rhs231;
      let _1345_v5;
      let _init41 = ((_1346_p0, _1347_v1) => function (_1348_i3) {
        return _dafny.Map.Empty.slice().updateUnsafe((_this).f22,new BigNumber((_dafny.MultiSet.fromElements(_1346_p0, new BigNumber((_1347_v1).length), _1346_p0, _1346_p0, _1346_p0)).cardinality()));
      })(p0, _1338_v1);
      let _nw240 = Array((new BigNumber(7)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw240.length); _i0_41++) {
        _nw240[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _1345_v5 = _nw240;
      let _1349_v6;
      _1349_v6 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_1338_v1).length));
      let _nw241 = Array((_dafny.ONE).toNumber());
      _nw241[(_dafny.ZERO).toNumber()] = (_1349_v6).Merge(_1349_v6);
      _1345_v5 = _nw241;
      (_this).f21 = p1;
      let _hi11 = p0;
      for (let _1350_i4 = new BigNumber((_1338_v1).length); _1350_i4.isLessThan(_hi11); _1350_i4 = _1350_i4.plus(_dafny.ONE)) {
        (globalState).f8 = (_this).f22;
        let _1351_v7;
        _1351_v7 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xkcswj"));
        let _1352_v8;
        _1352_v8 = _module.D5.create_DC13(_module.__default.fm54(p1, !((_this).f22), _this.f21, globalState));
        let _source13 = (((_1351_v7).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1338_v1))) ? (_1352_v8) : (_1352_v8));
        if (_source13.is_DC14) {
          let _1353___mcc_h0 = (_source13).cf26;
          let _1354_cf26 = _1353___mcc_h0;
          (globalState).f8 = (_this).f22;
          let _1355_v9;
          let _init42 = ((_1356_p3) => function (_1357_i5) {
            return _1356_p3;
          })(p3);
          let _nw242 = Array((new BigNumber(6)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw242.length); _i0_42++) {
            _nw242[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1355_v9 = _nw242;
          let _1358_v10;
          _1358_v10 = _dafny.Seq.of(_1355_v9);
          _1358_v10 = _dafny.Seq.Concat(_1358_v10, _1358_v10);
          let _1359_v11;
          _1359_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.Create(_module.__default.abs(new BigNumber(642)), function (_1360_i6) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }));
          (globalState).f9 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1350_i4, _module.__default.safeDivisionInt(new BigNumber((_1359_v11).length), p2)));
          let _1361_v12;
          _1361_v12 = _module.D0.create_DC0((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1344_v4).length))));
          let _1362_v13;
          _1362_v13 = _dafny.MultiSet.fromElements(p2);
          let _1363_v14;
          _1363_v14 = _dafny.Seq.of(p0);
          let _1364_v15;
          _1364_v15 = _dafny.Set.fromElements(p2);
          let _1365_v16;
          _1365_v16 = _dafny.Seq.of(_1364_v15);
          let _1366_v17;
          _1366_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_1338_v1, _module.__default.safeIndex(new BigNumber((_1363_v14).length), new BigNumber((_1338_v1).length)), new _dafny.CodePoint('k'.codePointAt(0)))).length),_1355_v9);
          let _1367_v18;
          _1367_v18 = _module.D8.create_DC19(_1338_v1, _1350_i4);
          let _1368_v19;
          let _nw243 = Array((new BigNumber(26)).toNumber());
          _nw243[(_dafny.ZERO).toNumber()] = _1350_i4;
          _nw243[(_dafny.ONE).toNumber()] = p2;
          _nw243[(new BigNumber(2)).toNumber()] = (p0).multipliedBy((_1361_v12).dtor_cf0);
          _nw243[(new BigNumber(3)).toNumber()] = p2;
          _nw243[(new BigNumber(4)).toNumber()] = p0;
          _nw243[(new BigNumber(5)).toNumber()] = _1350_i4;
          _nw243[(new BigNumber(6)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(_1338_v1)).cardinality())).minus(p2);
          _nw243[(new BigNumber(7)).toNumber()] = _1350_i4;
          _nw243[(new BigNumber(8)).toNumber()] = ((_this.f21) ? (p0) : (_1350_i4));
          _nw243[(new BigNumber(9)).toNumber()] = p0;
          _nw243[(new BigNumber(10)).toNumber()] = _1350_i4;
          _nw243[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1362_v13).cardinality()));
          _nw243[(new BigNumber(12)).toNumber()] = new BigNumber((_1363_v14).length);
          _nw243[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), ((_1369_i4) => function (_1370_i7) {
            return _dafny.Set.fromElements(_1369_i4);
          })(_1350_i4)), _1365_v16)).length);
          _nw243[(new BigNumber(14)).toNumber()] = new BigNumber(((_1366_v17).Merge(_1366_v17)).length);
          _nw243[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(_1350_i4, (_1367_v18).dtor_cf35);
          _nw243[(new BigNumber(16)).toNumber()] = new BigNumber((_1362_v13).cardinality());
          _nw243[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(249), (_dafny.ZERO).minus(p2));
          _nw243[(new BigNumber(18)).toNumber()] = (_1350_i4).minus(new BigNumber((_1338_v1).length));
          _nw243[(new BigNumber(19)).toNumber()] = p2;
          _nw243[(new BigNumber(20)).toNumber()] = new BigNumber((_1338_v1).length);
          _nw243[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(108), new BigNumber((_1344_v4).length));
          _nw243[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw243[(new BigNumber(23)).toNumber()] = p0;
          _nw243[(new BigNumber(24)).toNumber()] = p2;
          _nw243[(new BigNumber(25)).toNumber()] = _module.__default.fm0((_this).f22, globalState);
          _1368_v19 = _nw243;
          let _index206 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1368_v19).length));
          (_1368_v19)[_index206] = p0;
          let _index207 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1368_v19).length));
          (_1368_v19)[_index207] = p2;
        } else {
          let _1371___mcc_h1 = (_source13).cf25;
          let _1372_cf25 = _1371___mcc_h1;
          let _1373_v20;
          let _nw244 = Array((new BigNumber(17)).toNumber()).fill(false);
          _1373_v20 = _nw244;
          let _1374_v21;
          _1374_v21 = _dafny.Set.fromElements(_1373_v20);
          let _rhs232 = (_dafny.ZERO).minus(p2);
          let _rhs233 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), new BigNumber((_1338_v1).length)));
          let _rhs234 = !(_1374_v21).contains(_1373_v20);
          let _rhs235 = p3;
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          let _lhs161 = globalState;
          _lhs159.f9 = _rhs232;
          _lhs160.f9 = _rhs233;
          _lhs161.f8 = _rhs234;
          r0 = _rhs235;
          let _1375_v23;
          _1375_v23 = _dafny.MultiSet.fromElements(_this.f21);
          let _1376_v24;
          _1376_v24 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f22);
          let _1377_v25;
          _1377_v25 = _module.D8.create_DC19(_1338_v1, p2);
          let _1378_v26;
          _1378_v26 = _dafny.MultiSet.fromElements(p0);
          let _1379_v27;
          let _out12;
          _out12 = (_this).m12(((function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of _dafny.IntegerRange(new BigNumber(373), new BigNumber(-484))) {
              let _1380_v22 = _compr_54;
              if (((new BigNumber(373)).isLessThanOrEqualTo(_1380_v22)) && ((_1380_v22).isLessThan(new BigNumber(-484)))) {
                _coll54.push([_module.__default.safeDivisionInt(_1380_v22, p2),_dafny.Map.Empty.slice().updateUnsafe(_this.f21,p3)]);
              }
            }
            return _coll54;
          }()).update(new BigNumber((_1375_v23).cardinality()), _1376_v24)).equals(_dafny.Map.Empty.slice().updateUnsafe((_1377_v25).dtor_cf35,_1376_v24)), (_1378_v26).IsSubsetOf(_1378_v26), _1349_v6, globalState);
          _1379_v27 = _out12;
          (globalState).f8 = !(((_this).f22) === (p1));
          let _1381_v28;
          _1381_v28 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f22) ? ((_dafny.ZERO).minus(_1350_i4)) : (_1350_i4)),_1373_v20);
          let _1382_v29;
          _1382_v29 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1383_v30;
          _1383_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1382_v29,new BigNumber((_1338_v1).length));
          _1381_v28 = (_dafny.Map.Empty.slice().updateUnsafe((((_1383_v30).contains(_1382_v29)) ? ((_1383_v30).get(_1382_v29)) : (p2)),_1373_v20)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1350_i4,_1373_v20));
        }
        (globalState).f8 = p3;
        let _1384_v31;
        _1384_v31 = _module.D4.create_DC12(true, false, (_this).f22);
        let _1385_v32;
        _1385_v32 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("tey"));
        let _1386_v33;
        _1386_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1338_v1,(_1385_v32)[_module.__default.safeIndex(p2, new BigNumber((_1385_v32).length))]);
        (globalState).f8 = ((_1384_v31).dtor_cf24) && (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("pxt"), (((_1386_v33).contains(_1338_v1)) ? ((_1386_v33).get(_1338_v1)) : (_dafny.Seq.UnicodeFromString("turdyhf")))));
      }
      let _1387_i8;
      _1387_i8 = _dafny.ZERO;
      L6: {
        while (p1) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1387_i8)) {
              break L6;
            }
            _1387_i8 = (_1387_i8).plus(_dafny.ONE);
            _1338_v1 = _1338_v1;
            let _1388_v34;
            let _init43 = function (_1389_i9) {
              return (_this).f22;
            };
            let _nw245 = Array((new BigNumber(2)).toNumber());
            for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw245.length); _i0_43++) {
              _nw245[_i0_43] = _init43(new BigNumber(_i0_43));
            }
            _1388_v34 = _nw245;
            let _1390_v35;
            _1390_v35 = _module.D7.create_DC17(_1338_v1, _1388_v34, (_this).f22, new BigNumber(629));
            let _1391_v36;
            _1391_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1390_v35,_1344_v4);
            _1391_v36 = (_1391_v36).update(_1390_v35, _1344_v4);
            let _1392_v37;
            _1392_v37 = new _dafny.CodePoint('e'.codePointAt(0));
            let _pat_let_tv12 = _1392_v37;
            let _source14 = function (_pat_let15_0) {
              return function (_1393_dt__update__tmp_h0) {
                return function (_pat_let16_0) {
                  return function (_1394_dt__update_hcf9_h0) {
                    return _module.D1.create_DC4(_1394_dt__update_hcf9_h0);
                  }(_pat_let16_0);
                }(_pat_let_tv12);
              }(_pat_let15_0);
            }(_module.D1.create_DC4(_1392_v37));
            if (_source14.is_DC4) {
              let _1395___mcc_h2 = (_source14).cf9;
              let _1396_cf9 = _1395___mcc_h2;
              let _1397_v38;
              _1397_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f21);
              let _1398_v39;
              _1398_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_1397_v38);
              let _1399_v40;
              let _init44 = ((_1400_p2) => function (_1401_i10) {
                return _module.__default.safeModuloInt(_1401_i10, _1400_p2);
              })(p2);
              let _nw246 = Array((new BigNumber(13)).toNumber());
              for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw246.length); _i0_44++) {
                _nw246[_i0_44] = _init44(new BigNumber(_i0_44));
              }
              _1399_v40 = _nw246;
              let _1402_v41;
              _1402_v41 = _dafny.MultiSet.fromElements(_1399_v40);
              let _1403_v42;
              _1403_v42 = _module.D0.create_DC1(_1402_v41, p2, new BigNumber(777));
              let _1404_v43;
              _1404_v43 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber((_1398_v39).length)).minus((_1403_v42).dtor_cf3)),_1388_v34);
              _1404_v43 = (_1404_v43).update(_module.__default.fm0(_this.f21, globalState), _1388_v34);
              let _1405_v44;
              _1405_v44 = _dafny.Seq.of(new BigNumber((_1338_v1).length), p0, p0, (_dafny.ZERO).minus(p2));
              (globalState).f9 = (_1405_v44)[_module.__default.safeIndex(p2, new BigNumber((_1405_v44).length))];
              (globalState).f9 = p2;
              let _1406_v45;
              let _nw247 = new _module.C2();
              _nw247.__ctor((_this).f11);
              _1406_v45 = _nw247;
              let _index208 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1399_v40).length));
              (_1399_v40)[_index208] = p2;
              let _index209 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1388_v34).length));
              (_1388_v34)[_index209] = (_this).f22;
              let _1407_v46;
              _1407_v46 = _dafny.Set.fromElements((_this).f22);
              let _1408_v47;
              _1408_v47 = _dafny.MultiSet.fromElements(_this.f21);
              let _index210 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1399_v40).length));
              let _index211 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1388_v34).length));
              let _rhs236 = new BigNumber((_dafny.Seq.Concat(_1338_v1, _1338_v1)).length);
              let _rhs237 = _1406_v45;
              let _rhs238 = _module.__default.safeModuloInt(((true) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), ((_1409_v37) => function (_1410_i11) {
                return _1409_v37;
              })(_1392_v37))).length))) : (p0)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1338_v1, _module.__default.safeIndex(p0, new BigNumber((_1338_v1).length)), _1396_cf9), _1338_v1)).length)));
              let _rhs239 = _module.__default.fm0(p3, globalState);
              let _rhs240 = ((!((p2).isLessThan((_dafny.ZERO).minus(new BigNumber((_1407_v46).length))))) ? (!_dafny.Seq.contains(_1344_v4, _this.f21)) : ((_1408_v47).IsSubsetOf(_1408_v47)));
              let _lhs162 = globalState;
              let _lhs163 = _1399_v40;
              let _lhs164 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1399_v40).length));
              let _lhs165 = globalState;
              let _lhs166 = _1388_v34;
              let _lhs167 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1388_v34).length));
              _lhs162.f9 = _rhs236;
              _1406_v45 = _rhs237;
              _lhs163[_lhs164] = _rhs238;
              _lhs165.f9 = _rhs239;
              _lhs166[_lhs167] = _rhs240;
            } else {
              let _1411___mcc_h3 = (_source14).cf8;
              let _1412_cf8 = _1411___mcc_h3;
              let _1413_v48;
              _1413_v48 = _dafny.Set.fromElements(p1, p3, _this.f21);
              let _1414_v49;
              _1414_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1413_v48,new BigNumber(822));
              _1414_v49 = (_1414_v49).update(_1413_v48, p2);
              _1392_v37 = _1392_v37;
              _module.__default.m0(globalState);
              let _1415_v50;
              _1415_v50 = _module.D5.create_DC14(p1);
              _1415_v50 = (((function () {
                let _coll55 = new _dafny.Map();
                for (const _compr_55 of _dafny.IntegerRange(new BigNumber(412), new BigNumber(572))) {
                  let _1416_v51 = _compr_55;
                  if (((new BigNumber(412)).isLessThanOrEqualTo(_1416_v51)) && ((_1416_v51).isLessThan(new BigNumber(572)))) {
                    _coll55.push([(_1416_v51).minus(p2),true]);
                  }
                }
                return _coll55;
              }()).contains(p0)) ? (_module.__default.fm55(p1, p1, new BigNumber(-483), globalState)) : (_1415_v50));
            }
            let _rhs241 = _this.f21;
            let _rhs242 = (p0).minus(new BigNumber(-712));
            let _rhs243 = p2;
            let _lhs168 = globalState;
            let _lhs169 = globalState;
            let _lhs170 = globalState;
            _lhs168.f8 = _rhs241;
            _lhs169.f9 = _rhs242;
            _lhs170.f9 = _rhs243;
          }
        }
      }
      r0 = p1;
      let _1417_v52;
      let _nw248 = Array((new BigNumber(28)).toNumber()).fill([]);
      _1417_v52 = _nw248;
      r1 = _1417_v52;
      return [r0, r1];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f11 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    fm41(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wmeejk"), _dafny.Seq.UnicodeFromString("guqpxy"))).length)).minus((new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(400),new BigNumber((function () {
        let _coll56 = new _dafny.Map();
        for (const _compr_56 of _dafny.IntegerRange(new BigNumber(674), new BigNumber(977))) {
          let _1418_v0 = _compr_56;
          if (((new BigNumber(674)).isLessThanOrEqualTo(_1418_v0)) && ((_1418_v0).isLessThan(new BigNumber(977)))) {
            _coll56.push([(_1418_v0).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0))))).cardinality())),(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rhmhc")).length)))).length))]);
          }
        }
        return _coll56;
      }()).length))).length)))).cardinality())).minus(new BigNumber(-830)));
    };
    fm42(p0, p1, globalState) {
      let _this = this;
      let _source15 = _module.D2.create_DC7(true, new BigNumber(8), false, _dafny.Seq.UnicodeFromString("qusyb"));
      if (_source15.is_DC6) {
        return true;
      } else if (_source15.is_DC7) {
        let _1419___mcc_h0 = (_source15).cf11;
        let _1420___mcc_h1 = (_source15).cf12;
        let _1421___mcc_h2 = (_source15).cf13;
        let _1422___mcc_h3 = (_source15).cf14;
        let _1423_cf14 = _1422___mcc_h3;
        let _1424_cf13 = _1421___mcc_h2;
        let _1425_cf12 = _1420___mcc_h1;
        let _1426_cf11 = _1419___mcc_h0;
        return true;
      } else if (_source15.is_DC8) {
        let _1427___mcc_h4 = (_source15).cf15;
        let _1428___mcc_h5 = (_source15).cf16;
        let _1429_cf16 = _1428___mcc_h5;
        let _1430_cf15 = _1427___mcc_h4;
        return !(false) || (false);
      } else {
        let _1431___mcc_h6 = (_source15).cf10;
        let _1432_cf10 = _1431___mcc_h6;
        if (false) {
          return false;
        } else {
          return true;
        }
      }
    };
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _1433_v0;
      _1433_v0 = _dafny.Seq.of(p0);
      (globalState).f8 = _dafny.Seq.contains(_dafny.Seq.Concat(_1433_v0, _1433_v0), p0);
      let _1434_v1;
      _1434_v1 = new BigNumber(-850);
      let _1435_v2;
      _1435_v2 = _dafny.Set.fromElements((_1433_v0)[_module.__default.safeIndex(_1434_v1, new BigNumber((_1433_v0).length))]);
      let _1436_v3;
      _1436_v3 = _dafny.MultiSet.fromElements(new BigNumber(256));
      let _1437_v4;
      _1437_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1436_v3).cardinality()),new BigNumber((_dafny.Seq.of(!(p0))).length));
      let _rhs244 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_1433_v0, _1433_v0)).length), (_1434_v1).multipliedBy(new BigNumber((_1435_v2).length)));
      let _rhs245 = (_1434_v1).plus(new BigNumber((_1437_v4).length));
      let _rhs246 = _1434_v1;
      let _rhs247 = _1434_v1;
      let _lhs171 = globalState;
      r0 = _rhs244;
      r0 = _rhs245;
      r0 = _rhs246;
      _lhs171.f9 = _rhs247;
      let _1438_v5;
      _1438_v5 = _module.D5.create_DC14(p0);
      let _1439_v6;
      _1439_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1438_v5,_module.__default.fm0(p0, globalState));
      (globalState).f9 = ((((_1439_v6).contains(_module.D5.create_DC14(p0))) ? ((_1439_v6).get(_module.D5.create_DC14(p0))) : (_1434_v1))).multipliedBy(new BigNumber(-842));
      let _1440_v7;
      _1440_v7 = _module.D0.create_DC0(_1434_v1);
      let _1441_v8;
      let _nw249 = new _module.C1();
      _nw249.__ctor(false, false, _1440_v7, p0);
      _1441_v8 = _nw249;
      let _1442_v9;
      _1442_v9 = _dafny.MultiSet.fromElements(_1441_v8);
      let _pat_let_tv13 = p0;
      let _pat_let_tv14 = _1441_v8;
      let _pat_let_tv15 = p0;
      let _pat_let_tv16 = p0;
      let _pat_let_tv17 = _1441_v8;
      let _pat_let_tv18 = _1441_v8;
      let _pat_let_tv19 = _1434_v1;
      let _pat_let_tv20 = _1441_v8;
      let _pat_let_tv21 = _1434_v1;
      let _pat_let_tv22 = _1441_v8;
      let _source16 = function (_source17) {
        if (_source17.is_DC11) {
          let _1443___mcc_h7 = (_source17).cf19;
          let _1444___mcc_h8 = (_source17).cf20;
          let _1445___mcc_h9 = (_source17).cf21;
          let _1446_cf21 = _1445___mcc_h9;
          let _1447_cf20 = _1444___mcc_h8;
          let _1448_cf19 = _1443___mcc_h7;
          if (_pat_let_tv13) {
            return _module.D2.create_DC7((_pat_let_tv14).f19, (_dafny.ZERO).minus(_1448_cf19), _pat_let_tv15, _dafny.Seq.UnicodeFromString("pvjyrtkae"));
          } else {
            return _module.D2.create_DC7(_pat_let_tv16, _1448_cf19, (_pat_let_tv17).f20, _dafny.Seq.UnicodeFromString("aqxjcm"));
          }
        } else if (_source17.is_DC12) {
          let _1449___mcc_h10 = (_source17).cf22;
          let _1450___mcc_h11 = (_source17).cf23;
          let _1451___mcc_h12 = (_source17).cf24;
          let _1452_cf24 = _1451___mcc_h12;
          let _1453_cf23 = _1450___mcc_h11;
          let _1454_cf22 = _1449___mcc_h10;
          return _module.D2.create_DC7((_pat_let_tv18).f19, _pat_let_tv19, _1453_cf23, _dafny.Seq.UnicodeFromString("qx"));
        } else {
          let _1455___mcc_h13 = (_source17).cf18;
          let _1456_cf18 = _1455___mcc_h13;
          return _module.D2.create_DC7((_pat_let_tv20).f19, _pat_let_tv21, (_pat_let_tv22).f19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), function (_1457_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}));
        }
      }(_module.__default.fm43(p0, _1433_v0, new BigNumber((_1442_v9).cardinality()), new BigNumber(-980), globalState));
      if (_source16.is_DC6) {
        let _1458_v10;
        let _init45 = ((_1459_v1) => function (_1460_i1) {
          return _module.__default.safeDivisionInt(_1460_i1, _1459_v1);
        })(_1434_v1);
        let _nw250 = Array((new BigNumber(22)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw250.length); _i0_45++) {
          _nw250[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1458_v10 = _nw250;
        let _index212 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length));
        (_1458_v10)[_index212] = (new BigNumber((_dafny.Seq.UnicodeFromString("qaxhnnkb")).length)).plus(_1434_v1);
        let _index213 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length));
        (_1458_v10)[_index213] = (_dafny.ZERO).minus(_1434_v1);
        (globalState).f8 = !((_1441_v8).f20) || ((_1441_v8).f19);
        if (!((_1441_v8).f19)) {
          (globalState).f9 = (_1434_v1).plus((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))]);
          let _1461_v11;
          _1461_v11 = _dafny.Set.fromElements(_1434_v1, _1434_v1);
          r0 = ((_module.D2.create_DC7((_1441_v8).f19, (_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))], false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-336)), function (_1462_i2) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}))).dtor_cf12).plus((new BigNumber((_1461_v11).length)).plus(new BigNumber(650)));
          let _1463_v12;
          _1463_v12 = _module.D8.create_DC18(_1461_v11);
          _1461_v11 = (function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of ((_1463_v12).dtor_cf33).Elements) {
              let _1464_v13 = _compr_57;
              if (((_1463_v12).dtor_cf33).contains(_1464_v13)) {
                _coll57.add((_1464_v13).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber(-435))).length)));
              }
            }
            return _coll57;
          }()).Difference(_1461_v11);
          let _1465_v14;
          let _init46 = ((_1466_v8) => function (_1467_i3) {
            return (_1466_v8).f19;
          })(_1441_v8);
          let _nw251 = Array((new BigNumber(10)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw251.length); _i0_46++) {
            _nw251[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1465_v14 = _nw251;
          let _1468_v15;
          _1468_v15 = _dafny.Seq.UnicodeFromString("e");
          _1465_v14 = (_module.D7.create_DC17(_1468_v15, _1465_v14, (_1441_v8).f19, _1434_v1)).dtor_cf30;
          let _1469_v16;
          _1469_v16 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1470_v17;
          _1470_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm42((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))], _1434_v1, globalState),_1434_v1);
          let _1471_v18;
          _1471_v18 = _dafny.Seq.of(new BigNumber((_1470_v17).length));
          let _1472_v19;
          _1472_v19 = _dafny.MultiSet.fromElements(_1465_v14, _1465_v14);
          let _rhs248 = _module.__default.safeDivisionInt(((!(!(false))) ? ((_this).fm41(_1468_v15, new BigNumber((_1471_v18).length), _1434_v1, (_1441_v8).f20, globalState)) : ((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))])), new BigNumber((_1472_v19).cardinality()));
          let _rhs249 = _1469_v16;
          let _lhs172 = globalState;
          _lhs172.f9 = _rhs248;
          _1469_v16 = _rhs249;
        } else {
          let _1473_v20;
          _1473_v20 = new _dafny.CodePoint('v'.codePointAt(0));
          _1473_v20 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1474_v21;
          _1474_v21 = _module.D2.create_DC5(_dafny.Seq.UnicodeFromString("kqmqprqxw"));
          let _1475_v22;
          _1475_v22 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1441_v8).f20);
          let _1476_v23;
          _1476_v23 = _dafny.Seq.UnicodeFromString("gx");
          let _1477_v24;
          _1477_v24 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), ((_1478_v20) => function (_1479_i4) {
            return _1478_v20;
          })(_1473_v20)), _1476_v23);
          let _1480_v25;
          _1480_v25 = _module.D1.create_DC4(_1473_v20);
          let _1481_v26;
          _1481_v26 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), ((_1482_v20) => function (_1483_i5) {
            return _1482_v20;
          })(_1473_v20))).length));
          let _1484_v27;
          _1484_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1477_v24)[_module.__default.safeIndex((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))], new BigNumber((_1477_v24).length))],(_1441_v8).fm9(_1480_v25, _1481_v26, globalState));
          let _1485_v28;
          let _nw252 = Array((new BigNumber(9)).toNumber());
          _nw252[(_dafny.ZERO).toNumber()] = (_1441_v8).f20;
          _nw252[(_dafny.ONE).toNumber()] = (_1441_v8).f19;
          _nw252[(new BigNumber(2)).toNumber()] = (_1441_v8).f20;
          _nw252[(new BigNumber(3)).toNumber()] = p0;
          _nw252[(new BigNumber(4)).toNumber()] = !((_1441_v8).f19);
          _nw252[(new BigNumber(5)).toNumber()] = (_1441_v8).f20;
          _nw252[(new BigNumber(6)).toNumber()] = (_1441_v8).f19;
          _nw252[(new BigNumber(7)).toNumber()] = (_1441_v8).f20;
          _nw252[(new BigNumber(8)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_1475_v22).length))).isLessThan((((_1484_v27).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_1488_v20) => function (_1489_i6) {
            return _1488_v20;
          })(_1473_v20)))) ? ((_1484_v27).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_1486_v20) => function (_1487_i6) {
            return _1486_v20;
          })(_1473_v20)))) : (_1434_v1)));
          _1485_v28 = _nw252;
          let _1490_v29;
          _1490_v29 = _dafny.Seq.of((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))]);
          let _index214 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1485_v28).length));
          (_1485_v28)[_index214] = (new BigNumber((_dafny.Seq.update(_1490_v29, _module.__default.safeIndex(_1434_v1, new BigNumber((_1490_v29).length)), new BigNumber((_1437_v4).length))).length)).isLessThan((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))]);
          let _index215 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1485_v28).length));
          let _rhs250 = _1474_v21;
          let _rhs251 = (_module.__default.safeDivisionInt(_1434_v1, new BigNumber((_dafny.Seq.of(_1434_v1)).length))).isLessThan(((false) ? (_1434_v1) : ((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))])));
          let _rhs252 = _1473_v20;
          let _rhs253 = (((_1441_v8).f19) ? (_1434_v1) : (new BigNumber(837)));
          let _rhs254 = _dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_1433_v0, _dafny.Seq.of(p0, (_1441_v8).f20)), _module.__default.safeIndex(new BigNumber(-896), new BigNumber((_dafny.Seq.Concat(_1433_v0, _dafny.Seq.of(p0, (_1441_v8).f20))).length)), (_1441_v8).f19), _1433_v0);
          let _lhs173 = _1485_v28;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1485_v28).length));
          let _lhs175 = globalState;
          _1474_v21 = _rhs250;
          _lhs173[_lhs174] = _rhs251;
          _1473_v20 = _rhs252;
          r0 = _rhs253;
          _lhs175.f8 = _rhs254;
          let _1491_v30;
          let _init47 = ((_1492_v10, _1493_v29) => function (_1494_i7) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(298)), ((_1495_v10) => function (_1496_i8) {
              return (_1495_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1495_v10).length))];
            })(_1492_v10)), _1493_v29);
          })(_1458_v10, _1490_v29);
          let _nw253 = Array((new BigNumber(16)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw253.length); _i0_47++) {
            _nw253[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1491_v30 = _nw253;
          let _index216 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1491_v30).length));
          (_1491_v30)[_index216] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_1434_v1), _module.__default.safeIndex(_1434_v1, new BigNumber((_dafny.Seq.of(_1434_v1)).length)), (_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(464)), ((_1497_v10) => function (_1498_i9) {
            return (_1497_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1497_v10).length))];
          })(_1458_v10)));
          let _index217 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1491_v30).length));
          (_1491_v30)[_index217] = _dafny.Seq.Concat(_1490_v29, _1490_v29);
          let _1499_v31;
          let _nw254 = new _module.C2();
          _nw254.__ctor((_this).f11);
          _1499_v31 = _nw254;
          let _1500_v32;
          let _1501_v33;
          let _out13;
          let _out14;
          let _outcollector4 = (_1441_v8).m9(_dafny.MultiSet.fromElements((_1491_v30)[_module.__default.safeIndex(new BigNumber(631), new BigNumber((_1491_v30).length))]), globalState);
          _out13 = _outcollector4[0];
          _out14 = _outcollector4[1];
          _1500_v32 = _out13;
          _1501_v33 = _out14;
        }
        if (!(!((_1441_v8).f20)) || (false)) {
          let _1502_v34;
          _1502_v34 = new _dafny.CodePoint('v'.codePointAt(0));
          let _1503_v35;
          let _nw255 = Array((new BigNumber(21)).toNumber());
          _nw255[(_dafny.ZERO).toNumber()] = _1502_v34;
          _nw255[(_dafny.ONE).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(2)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(3)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(4)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(5)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(6)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(7)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(8)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(9)).toNumber()] = _module.__default.fm3((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))], (_1441_v8).f19, globalState);
          _nw255[(new BigNumber(10)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw255[(new BigNumber(12)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(13)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(14)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(15)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(16)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(17)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(18)).toNumber()] = _1502_v34;
          _nw255[(new BigNumber(19)).toNumber()] = _module.__default.fm3((_dafny.ZERO).minus(new BigNumber(-858)), (_1441_v8).f20, globalState);
          _nw255[(new BigNumber(20)).toNumber()] = _1502_v34;
          _1503_v35 = _nw255;
          let _index218 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1503_v35).length));
          (_1503_v35)[_index218] = _1502_v34;
          let _1504_v36;
          _1504_v36 = _module.D1.create_DC4(_1502_v34);
          let _1505_v37;
          _1505_v37 = _dafny.Set.fromElements((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))]);
          let _index219 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1503_v35).length));
          let _rhs255 = _1502_v34;
          let _rhs256 = (_1441_v8).f20;
          let _rhs257 = (_1441_v8).fm9(_1504_v36, _1505_v37, globalState);
          let _rhs258 = (_1441_v8).f19;
          let _lhs176 = _1503_v35;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1503_v35).length));
          let _lhs178 = globalState;
          let _lhs179 = globalState;
          _lhs176[_lhs177] = _rhs255;
          _lhs178.f8 = _rhs256;
          r0 = _rhs257;
          _lhs179.f8 = _rhs258;
          _1502_v34 = (_1503_v35)[_module.__default.safeIndex(new BigNumber(84), new BigNumber((_1503_v35).length))];
          (globalState).f9 = (_1434_v1).plus(_1434_v1);
          _1434_v1 = (_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))];
          let _1506_v38;
          _1506_v38 = _dafny.Set.fromElements(_1458_v10, _1458_v10, _1458_v10, _1458_v10, _1458_v10);
          let _index220 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length));
          (_1458_v10)[_index220] = new BigNumber((_1506_v38).length);
        } else {
          (globalState).f8 = p0;
          let _1507_v39;
          _1507_v39 = _module.D8.create_DC20(_module.D0.create_DC2(_dafny.MultiSet.fromElements(false, (_1441_v8).f19), (_1441_v8).f19, _1434_v1, (_1441_v8).f19), true, (_1441_v8).f19);
          let _1508_v40;
          _1508_v40 = _dafny.Seq.of(_1507_v39);
          let _1509_v41;
          _1509_v41 = _dafny.Seq.of(new BigNumber((_1508_v40).length));
          _1509_v41 = _dafny.Seq.Concat(_1509_v41, _1509_v41);
          r0 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("psom"), _dafny.Seq.UnicodeFromString("jjhaifpp"))).length), _1434_v1);
          let _1510_v42;
          _1510_v42 = _dafny.Seq.UnicodeFromString("lfja");
          let _1511_v43;
          _1511_v43 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1512_v44;
          _1512_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1510_v42,_1511_v43);
          let _1513_v45;
          _1513_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1441_v8).f20,_dafny.MultiSet.FromArray(_1509_v41));
          let _1514_v46;
          _1514_v46 = _dafny.Map.Empty.slice().updateUnsafe((_1441_v8).f19,_module.__default.fm0((_1441_v8).f19, globalState));
          _1512_v44 = (_1512_v44).update(_module.__default.fm2(_1434_v1, (((_1513_v45).contains(p0)) ? ((_1513_v45).get(p0)) : (_dafny.MultiSet.FromArray(_dafny.Seq.update(_1509_v41, _module.__default.safeIndex((((_1514_v46).contains((_1441_v8).f19)) ? ((_1514_v46).get((_1441_v8).f19)) : (_1434_v1)), new BigNumber((_1509_v41).length)), _1434_v1)))), function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_1437_v4).Keys.Elements) {
              let _1515_v47 = _compr_58;
              if ((_1437_v4).contains(_1515_v47)) {
                _coll58.push([(_1515_v47).minus((_1458_v10)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1458_v10).length))]),_1434_v1]);
              }
            }
            return _coll58;
          }(), globalState), _1511_v43);
          let _1516_v48;
          _1516_v48 = _module.D8.create_DC19(_dafny.Seq.UnicodeFromString("g"), new BigNumber(297));
          let _1517_v49;
          _1517_v49 = _module.D2.create_DC8(_1434_v1, (_1510_v42)[_module.__default.safeIndex((((_1437_v4).contains(_1434_v1)) ? ((_1437_v4).get(_1434_v1)) : ((_1516_v48).dtor_cf35)), new BigNumber((_1510_v42).length))]);
          let _rhs259 = (_1434_v1).plus(_1434_v1);
          let _rhs260 = (_1517_v49).dtor_cf15;
          let _rhs261 = (((_1441_v8).f19) ? (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(460)), ((_1518_v0) => function (_1519_i10) {
            return _1518_v0;
          })(_1433_v0)), _dafny.Seq.of(_1433_v0))) : ((_1441_v8).f20));
          let _lhs180 = globalState;
          let _lhs181 = globalState;
          let _lhs182 = globalState;
          _lhs180.f9 = _rhs259;
          _lhs181.f9 = _rhs260;
          _lhs182.f8 = _rhs261;
        }
      } else if (_source16.is_DC7) {
        let _1520___mcc_h0 = (_source16).cf11;
        let _1521___mcc_h1 = (_source16).cf12;
        let _1522___mcc_h2 = (_source16).cf13;
        let _1523___mcc_h3 = (_source16).cf14;
        let _1524_cf14 = _1523___mcc_h3;
        let _1525_cf13 = _1522___mcc_h2;
        let _1526_cf12 = _1521___mcc_h1;
        let _1527_cf11 = _1520___mcc_h0;
        let _1528_v50;
        let _nw256 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1528_v50 = _nw256;
        let _1529_v51;
        _1529_v51 = new _dafny.CodePoint('p'.codePointAt(0));
        let _index221 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1528_v50).length));
        (_1528_v50)[_index221] = _module.__default.fm1(_1527_cf11, _1529_v51, globalState);
        let _index222 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1528_v50).length));
        (_1528_v50)[_index222] = !((_1441_v8).f19) || ((_1441_v8).f19);
        r0 = _1434_v1;
        let _1530_v52;
        _1530_v52 = _dafny.Seq.of(_1526_cf12);
        _1530_v52 = _dafny.Seq.Concat(_1530_v52, _dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), ((_1531_cf12) => function (_1532_i11) {
          return (_dafny.ZERO).minus((_dafny.ZERO).minus(_1531_cf12));
        })(_1526_cf12)));
        _1524_cf14 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), ((_1533_v51) => function (_1534_i12) {
          return _1533_v51;
        })(_1529_v51));
      } else if (_source16.is_DC8) {
        let _1535___mcc_h4 = (_source16).cf15;
        let _1536___mcc_h5 = (_source16).cf16;
        let _1537_cf16 = _1536___mcc_h5;
        let _1538_cf15 = _1535___mcc_h4;
        let _1539_v53;
        let _init48 = ((_1540_v8, _1541_cf15) => function (_1542_i13) {
          return (_module.D8.create_DC20(_module.D0.create_DC2(_dafny.MultiSet.fromElements((_1540_v8).f20, (_1540_v8).f20), (_1540_v8).f19, _1541_cf15, (_1540_v8).f20), (_1540_v8).f20, true)).dtor_cf37;
        })(_1441_v8, _1538_cf15);
        let _nw257 = Array((_dafny.ONE).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw257.length); _i0_48++) {
          _nw257[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1539_v53 = _nw257;
        let _1543_v54;
        let _nw258 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1543_v54 = _nw258;
        let _index223 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length));
        (_1543_v54)[_index223] = new BigNumber(455);
        let _1544_v55;
        _1544_v55 = _module.D1.create_DC4(_1537_cf16);
        let _1545_v56;
        _1545_v56 = _dafny.Set.fromElements(_1538_cf15);
        let _index224 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length));
        let _rhs262 = ((false) ? (_1539_v53) : (_1539_v53));
        let _rhs263 = (_1441_v8).fm9(_1544_v55, _1545_v56, globalState);
        let _lhs183 = _1543_v54;
        let _lhs184 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length));
        _1539_v53 = _rhs262;
        _lhs183[_lhs184] = _rhs263;
        (globalState).f9 = (_dafny.ZERO).minus(new BigNumber(((_1435_v2).Intersect(_1435_v2)).length));
        let _1546_v57;
        let _init49 = ((_1547_v8, _1548_cf15, _1549_p0, _1550_v54) => function (_1551_i14) {
          return (((_1547_v8).f20) ? (_module.D2.create_DC7((_1547_v8).f19, _1548_cf15, _1549_p0, _dafny.Seq.UnicodeFromString("wpk"))) : (_module.D2.create_DC7((_1547_v8).f19, (_1550_v54)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_1550_v54).length))], true, _dafny.Seq.UnicodeFromString("mynlae"))));
        })(_1441_v8, _1538_cf15, p0, _1543_v54);
        let _nw259 = Array((_dafny.ONE).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw259.length); _i0_49++) {
          _nw259[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1546_v57 = _nw259;
        let _1552_v58;
        _1552_v58 = _module.D2.create_DC7((_1441_v8).f20, new BigNumber(773), (_1441_v8).f19, _dafny.Seq.UnicodeFromString("hticdw"));
        let _index225 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1546_v57).length));
        (_1546_v57)[_index225] = _1552_v58;
        let _1553_v59;
        _1553_v59 = _dafny.Seq.UnicodeFromString("qye");
        let _pat_let_tv23 = _1543_v54;
        let _pat_let_tv24 = _1543_v54;
        let _pat_let_tv25 = _1441_v8;
        let _pat_let_tv26 = _1553_v59;
        let _index226 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1546_v57).length));
        (_1546_v57)[_index226] = function (_pat_let17_0) {
          return function (_1555_dt__update__tmp_h0) {
            return function (_pat_let18_0) {
              return function (_1556_dt__update_hcf12_h0) {
                return function (_pat_let19_0) {
                  return function (_1557_dt__update_hcf11_h0) {
                    return function (_pat_let20_0) {
                      return function (_1558_dt__update_hcf14_h0) {
                        return _module.D2.create_DC7(_1557_dt__update_hcf11_h0, _1556_dt__update_hcf12_h0, (_1555_dt__update__tmp_h0).dtor_cf13, _1558_dt__update_hcf14_h0);
                      }(_pat_let20_0);
                    }(_pat_let_tv26);
                  }(_pat_let19_0);
                }((_pat_let_tv25).f20);
              }(_pat_let18_0);
            }((_pat_let_tv24)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_pat_let_tv23).length))]);
          }(_pat_let17_0);
        }((((_1441_v8).f19) ? (_module.D2.create_DC7(!((_1441_v8).f19), (_1543_v54)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length))], (_1441_v8).f20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-908)), function (_1554_i15) {
  return new _dafny.CodePoint('b'.codePointAt(0));
}))) : (_1552_v58)));
        let _1559_v60;
        _1559_v60 = _dafny.MultiSet.fromElements((_1441_v8).f19, false);
        let _1560_v61;
        _1560_v61 = _module.D0.create_DC2(_1559_v60, true, _1538_cf15, p0);
        let _1561_v62;
        _1561_v62 = _dafny.Seq.of(_1434_v1);
        let _1562_v63;
        _1562_v63 = _module.D8.create_DC20(_1560_v61, !_dafny.areEqual(_1561_v62, _1561_v62), (_1441_v8).f20);
        let _source18 = _1562_v63;
        if (_source18.is_DC19) {
          let _1563___mcc_h14 = (_source18).cf34;
          let _1564___mcc_h15 = (_source18).cf35;
          let _1565_cf35 = _1564___mcc_h15;
          let _1566_cf34 = _1563___mcc_h14;
          _1538_cf15 = (_1538_cf15).plus(_1538_cf15);
          let _1567_v64;
          _1567_v64 = _dafny.Map.Empty.slice().updateUnsafe((_1441_v8).f20,!(false) || (true));
          _1567_v64 = (_1567_v64).update(false, (_1441_v8).f20);
          _1537_cf16 = _1537_cf16;
          _1441_v8 = _1441_v8;
        } else if (_source18.is_DC20) {
          let _1568___mcc_h16 = (_source18).cf36;
          let _1569___mcc_h17 = (_source18).cf37;
          let _1570___mcc_h18 = (_source18).cf38;
          let _1571_cf38 = _1570___mcc_h18;
          let _1572_cf37 = _1569___mcc_h17;
          let _1573_cf36 = _1568___mcc_h16;
          let _index227 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1539_v53).length));
          (_1539_v53)[_index227] = (_module.__default.fm0(!((_1441_v8).f19), globalState)).isEqualTo(_1434_v1);
          let _index228 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1539_v53).length));
          (_1539_v53)[_index228] = !_dafny.Seq.contains(_1561_v62, (_dafny.ZERO).minus(_1434_v1));
          let _1574_v65;
          let _nw260 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1574_v65 = _nw260;
          let _index229 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1574_v65).length));
          (_1574_v65)[_index229] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-863)), ((_1575_cf16) => function (_1576_i16) {
            return _1575_cf16;
          })(_1537_cf16));
          let _index230 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1574_v65).length));
          (_1574_v65)[_index230] = _1553_v59;
          _1434_v1 = (_1543_v54)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length))];
          let _1577_v66;
          let _nw261 = new _module.C1();
          _nw261.__ctor((_1441_v8).f20, (_this).fm42((_1543_v54)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length))], new BigNumber(((_1574_v65)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1574_v65).length))]).length), globalState), _module.D0.create_DC0(new BigNumber(136)), true);
          _1577_v66 = _nw261;
        } else {
          let _1578___mcc_h19 = (_source18).cf33;
          let _1579_cf33 = _1578___mcc_h19;
          (globalState).f8 = (_1441_v8).f19;
          (globalState).f8 = (_1441_v8).f19;
          let _1580_v67;
          let _init50 = ((_1581_v8, _1582_v62) => function (_1583_i17) {
            return _dafny.Map.Empty.slice().updateUnsafe(!((_1581_v8).f19),new BigNumber((_1582_v62).length));
          })(_1441_v8, _1561_v62);
          let _nw262 = Array((new BigNumber(24)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw262.length); _i0_50++) {
            _nw262[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1580_v67 = _nw262;
          _1580_v67 = _1580_v67;
          let _index231 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1543_v54).length));
          (_1543_v54)[_index231] = (_1434_v1).multipliedBy(new BigNumber((_1561_v62).length));
        }
      } else {
        let _1584___mcc_h6 = (_source16).cf10;
        let _1585_cf10 = _1584___mcc_h6;
        let _1586_v68;
        let _nw263 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1586_v68 = _nw263;
        let _index232 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1586_v68).length));
        (_1586_v68)[_index232] = (_1441_v8).f20;
        let _1587_v69;
        _1587_v69 = _dafny.MultiSet.fromElements(p0);
        let _1588_v70;
        _1588_v70 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1589_v71;
        _1589_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1434_v1,_1588_v70);
        let _1590_v72;
        _1590_v72 = _module.D2.create_DC8(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1434_v1,_1434_v1)).length), _1588_v70);
        let _1591_v73;
        _1591_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1434_v1,(_1441_v8).f19);
        let _1592_v74;
        let _nw264 = Array((new BigNumber(11)).toNumber());
        _nw264[(_dafny.ZERO).toNumber()] = _1434_v1;
        _nw264[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("abdf")).length));
        _nw264[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_1590_v72).dtor_cf15);
        _nw264[(new BigNumber(3)).toNumber()] = (new BigNumber(48)).plus(_1434_v1);
        _nw264[(new BigNumber(4)).toNumber()] = new BigNumber((_1587_v69).cardinality());
        _nw264[(new BigNumber(5)).toNumber()] = _1434_v1;
        _nw264[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_1436_v3).contains(_1434_v1)) ? ((_1436_v3).get(_1434_v1)) : (_1434_v1)), _1434_v1));
        _nw264[(new BigNumber(7)).toNumber()] = _1434_v1;
        _nw264[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("h"), _1585_cf10)).length);
        _nw264[(new BigNumber(9)).toNumber()] = (_1441_v8).fm30(_1591_v73, globalState);
        _nw264[(new BigNumber(10)).toNumber()] = (_1434_v1).plus(_1434_v1);
        _1592_v74 = _nw264;
        let _index233 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
        (_1592_v74)[_index233] = ((_dafny.ZERO).minus(_1434_v1)).multipliedBy(_1434_v1);
        let _1593_v75;
        _1593_v75 = _dafny.Seq.of(_1585_cf10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), ((_1594_v70) => function (_1595_i18) {
          return _1594_v70;
        })(_1588_v70)));
        let _index234 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1586_v68).length));
        let _index235 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
        let _rhs264 = p0;
        let _rhs265 = _1587_v69;
        let _rhs266 = _1586_v68;
        let _rhs267 = _1589_v71;
        let _rhs268 = _module.__default.safeDivisionInt(new BigNumber((_1593_v75).length), _1434_v1);
        let _lhs185 = _1586_v68;
        let _lhs186 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1586_v68).length));
        let _lhs187 = _1592_v74;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
        _lhs185[_lhs186] = _rhs264;
        _1587_v69 = _rhs265;
        _1586_v68 = _rhs266;
        _1589_v71 = _rhs267;
        _lhs187[_lhs188] = _rhs268;
        let _index236 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
        (_1592_v74)[_index236] = (((_1441_v8).f20) ? (new BigNumber(305)) : ((new BigNumber(825)).multipliedBy(new BigNumber(318))));
        let _1596_v76;
        _1596_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1434_v1).plus(_1434_v1),_1436_v3);
        let _1597_v77;
        let _init51 = ((_1598_v8) => function (_1599_i19) {
          return _dafny.Seq.of(true, (_1598_v8).f19, (_1598_v8).f19, (_1598_v8).f19);
        })(_1441_v8);
        let _nw265 = Array((new BigNumber(22)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw265.length); _i0_51++) {
          _nw265[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1597_v77 = _nw265;
        let _1600_v78;
        _1600_v78 = _dafny.Seq.of(_1433_v0, _1433_v0, _dafny.Seq.of((_1441_v8).f20));
        let _index237 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length));
        (_1597_v77)[_index237] = _module.__default.fm44(_1600_v78, _1588_v70, globalState);
        let _1601_v79;
        let _nw266 = Array((new BigNumber(6)).toNumber()).fill([]);
        _1601_v79 = _nw266;
        let _1602_v80;
        let _init52 = ((_1603_v2) => function (_1604_i20) {
          return _1603_v2;
        })(_1435_v2);
        let _nw267 = Array((new BigNumber(9)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw267.length); _i0_52++) {
          _nw267[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1602_v80 = _nw267;
        let _index238 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1601_v79).length));
        (_1601_v79)[_index238] = _1602_v80;
        let _1605_v81;
        _1605_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1592_v74,(_1441_v8).f20);
        let _index239 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length));
        let _index240 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1601_v79).length));
        let _rhs269 = _1596_v76;
        let _rhs270 = _1433_v0;
        let _rhs271 = (((_1605_v81).contains(_1592_v74)) ? ((_1605_v81).get(_1592_v74)) : ((_1441_v8).f19));
        let _rhs272 = _1602_v80;
        let _lhs189 = _1597_v77;
        let _lhs190 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length));
        let _lhs191 = globalState;
        let _lhs192 = _1601_v79;
        let _lhs193 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1601_v79).length));
        _1596_v76 = _rhs269;
        _lhs189[_lhs190] = _rhs270;
        _lhs191.f8 = _rhs271;
        _lhs192[_lhs193] = _rhs272;
        if (((_1597_v77)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length))])[_module.__default.safeIndex(((_1592_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length))]).multipliedBy((_1592_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length))]), new BigNumber(((_1597_v77)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length))]).length))]) {
          let _1606_v82;
          let _nw268 = new _module.C0();
          _nw268.__ctor((((_1441_v8).f19) ? ((_1441_v8).f20) : ((_1441_v8).f19)));
          _1606_v82 = _nw268;
          let _index241 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
          (_1592_v74)[_index241] = _1434_v1;
          let _1607_v83;
          let _nw269 = Array((new BigNumber(17)).toNumber()).fill(_module.D0.Default());
          _1607_v83 = _nw269;
          let _1608_v84;
          _1608_v84 = _module.D0.create_DC2(_dafny.MultiSet.fromElements((_1441_v8).f20, (_1606_v82).f17), (_1441_v8).f20, (_1592_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length))], (_1441_v8).f20);
          let _index242 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_1607_v83).length));
          (_1607_v83)[_index242] = (((_1606_v82).f17) ? (_1608_v84) : (_1608_v84));
          let _index243 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_1607_v83).length));
          (_1607_v83)[_index243] = _1608_v84;
          (globalState).f8 = p0;
          let _1609_v85;
          _1609_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qwbav")).length),(_1593_v75)[_module.__default.safeIndex((_1592_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length))], new BigNumber((_1593_v75).length))]);
          let _1610_v86;
          _1610_v86 = _module.D9.create_DC22(_1434_v1, _1609_v85);
          (globalState).f9 = (_1610_v86).dtor_cf40;
        } else {
          let _pat_let_tv27 = _1434_v1;
          let _index244 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
          let _index245 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length));
          let _rhs273 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), ((_1611_p0, _1612_v74) => function (_1613_i21) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1611_p0,(_1612_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1612_v74).length))]);
          })(p0, _1592_v74))).length);
          let _rhs274 = (new BigNumber((_1591_v73).length)).plus(new BigNumber(606));
          let _rhs275 = _dafny.Seq.of(false);
          let _rhs276 = (function (_pat_let21_0) {
            return function (_1614_dt__update__tmp_h1) {
              return function (_pat_let22_0) {
                return function (_1615_dt__update_hcf15_h0) {
                  return _module.D2.create_DC8(_1615_dt__update_hcf15_h0, (_1614_dt__update__tmp_h1).dtor_cf16);
                }(_pat_let22_0);
              }(_pat_let_tv27);
            }(_pat_let21_0);
          }(_1590_v72)).dtor_cf15;
          let _lhs194 = _1592_v74;
          let _lhs195 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
          let _lhs196 = globalState;
          let _lhs197 = _1597_v77;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length));
          let _lhs199 = globalState;
          _lhs194[_lhs195] = _rhs273;
          _lhs196.f9 = _rhs274;
          _lhs197[_lhs198] = _rhs275;
          _lhs199.f9 = _rhs276;
          let _1616_v87;
          _1616_v87 = _module.D4.create_DC12(true, (_1441_v8).f20, (_1441_v8).f20);
          let _1617_v88;
          _1617_v88 = _dafny.MultiSet.fromElements(_1616_v87, _1616_v87);
          let _1618_v89;
          _1618_v89 = _module.D0.create_DC2(_dafny.MultiSet.FromArray((_1597_v77)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_1597_v77).length))]), (_1586_v68)[_module.__default.safeIndex(new BigNumber(162), new BigNumber((_1586_v68).length))], new BigNumber((_1617_v88).cardinality()), p0);
          let _1619_v90;
          _1619_v90 = _dafny.MultiSet.fromElements(_1618_v89, _1618_v89, _1618_v89);
          (globalState).f9 = (new BigNumber((_1619_v90).cardinality())).plus(_1434_v1);
          let _1620_v91;
          _1620_v91 = _dafny.Seq.of(_1592_v74);
          let _1621_v92;
          let _nw270 = Array((new BigNumber(11)).toNumber());
          _nw270[(_dafny.ZERO).toNumber()] = _1592_v74;
          _nw270[(_dafny.ONE).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(2)).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(3)).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(4)).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(5)).toNumber()] = (_1620_v91)[_module.__default.safeIndex(_1434_v1, new BigNumber((_1620_v91).length))];
          _nw270[(new BigNumber(6)).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(7)).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(8)).toNumber()] = _1592_v74;
          _nw270[(new BigNumber(9)).toNumber()] = (((_1441_v8).f19) ? (_1592_v74) : (_1592_v74));
          _nw270[(new BigNumber(10)).toNumber()] = _1592_v74;
          _1621_v92 = _nw270;
          let _1622_v93;
          _1622_v93 = _module.D9.create_DC21((_1620_v91)[_module.__default.safeIndex((_1592_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length))], new BigNumber((_1620_v91).length))]);
          let _index246 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1621_v92).length));
          (_1621_v92)[_index246] = (_1622_v93).dtor_cf39;
          let _index247 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1621_v92).length));
          (_1621_v92)[_index247] = _1592_v74;
          (globalState).f8 = !_dafny.Seq.contains(_1433_v0, (_1441_v8).f20);
          let _index248 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length));
          (_1592_v74)[_index248] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1592_v74)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1592_v74).length))]));
        }
      }
      let _1623_v94;
      _1623_v94 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(_1436_v3).Union(_1436_v3));
      let _1624_v95;
      _1624_v95 = new _dafny.CodePoint('o'.codePointAt(0));
      _1623_v94 = (_1623_v94).update((((_1441_v8).f20) ? (_1624_v95) : (_1624_v95)), _1436_v3);
      let _1625_i22;
      _1625_i22 = _dafny.ZERO;
      L7: {
        while ((_1441_v8).f19) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1625_i22)) {
              break L7;
            }
            _1625_i22 = (_1625_i22).plus(_dafny.ONE);
            let _1626_v96;
            _1626_v96 = _dafny.Seq.of(_1434_v1, _1434_v1, _1434_v1);
            _1437_v4 = (_1437_v4).update(((p0) ? (_1434_v1) : ((_dafny.ZERO).minus(_1434_v1))), new BigNumber((_1626_v96).length));
            let _1627_v97;
            let _nw271 = Array((new BigNumber(6)).toNumber());
            _nw271[(_dafny.ZERO).toNumber()] = false;
            _nw271[(_dafny.ONE).toNumber()] = (_1441_v8).f19;
            _nw271[(new BigNumber(2)).toNumber()] = p0;
            _nw271[(new BigNumber(3)).toNumber()] = !(!((_1441_v8).f19));
            _nw271[(new BigNumber(4)).toNumber()] = (_1441_v8).f20;
            _nw271[(new BigNumber(5)).toNumber()] = (_1441_v8).f19;
            _1627_v97 = _nw271;
            let _1628_v98;
            _1628_v98 = _dafny.Seq.of(_1627_v97);
            let _1629_v99;
            _1629_v99 = _dafny.Seq.of(_dafny.Seq.of(_1627_v97, _1627_v97), _dafny.Seq.of(_1627_v97, _1627_v97, _1627_v97));
            let _1630_v100;
            _1630_v100 = _dafny.Seq.of(_1628_v98, (_1629_v99)[_module.__default.safeIndex(_1434_v1, new BigNumber((_1629_v99).length))]);
            _1628_v98 = _dafny.Seq.Concat((_1630_v100)[_module.__default.safeIndex(_1434_v1, new BigNumber((_1630_v100).length))], _1628_v98);
            let _1631_v101;
            _1631_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1434_v1,p0);
            _1631_v101 = _1631_v101;
            let _1632_v102;
            let _nw272 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
            _1632_v102 = _nw272;
            let _index249 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_1632_v102).length));
            (_1632_v102)[_index249] = (_1434_v1).multipliedBy(new BigNumber(15));
            let _1633_v103;
            _1633_v103 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(765), (_dafny.ZERO).minus(_1434_v1), new BigNumber((_1628_v98).length), _1434_v1, _1434_v1));
            let _index250 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_1632_v102).length));
            (_1632_v102)[_index250] = ((!((_this).fm42(_1434_v1, new BigNumber((_1633_v103).cardinality()), globalState))) ? ((_dafny.ZERO).minus(_1434_v1)) : (new BigNumber(878)));
          }
        }
      }
      let _1634_v104;
      _1634_v104 = _dafny.Seq.UnicodeFromString("yxm");
      r0 = (_dafny.ZERO).minus(((!(p0) || (p0)) ? (_1434_v1) : ((new BigNumber((_1634_v104).length)).plus(_1434_v1))));
      r1 = ((_1435_v2).Union(_1435_v2)).Difference(_1435_v2);
      return [r0, r1];
    }
    m11(globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let _1635_v0;
      _1635_v0 = new BigNumber(-315);
      (globalState).f9 = _module.__default.safeDivisionInt(_1635_v0, _1635_v0);
      let _1636_v1;
      _1636_v1 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), function (_1637_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length), _1635_v0);
      let _1638_v2;
      _1638_v2 = _module.D5.create_DC13(_1636_v1);
      let _1639_v3;
      _1639_v3 = false;
      let _1640_v4;
      _1640_v4 = _dafny.Seq.of(_1639_v3);
      let _1641_v5;
      _1641_v5 = new _dafny.CodePoint('r'.codePointAt(0));
      let _1642_v6;
      _1642_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1641_v5,_1635_v0);
      let _1643_v7;
      _1643_v7 = _dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_1642_v6).length));
      let _1644_v8;
      _1644_v8 = _dafny.Set.fromElements(_1639_v3, _1639_v3);
      let _1645_v9;
      _1645_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(916),(_this).fm42(new BigNumber((_1640_v4).length), new BigNumber((_1644_v8).length), globalState));
      let _1646_v10;
      let _nw273 = Array((new BigNumber(22)).toNumber());
      _nw273[(_dafny.ZERO).toNumber()] = _1635_v0;
      _nw273[(_dafny.ONE).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(2)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(3)).toNumber()] = new BigNumber((_module.__default.fm45(_1635_v0, (_1638_v2).dtor_cf25, globalState)).cardinality());
      _nw273[(new BigNumber(4)).toNumber()] = new BigNumber((_1640_v4).length);
      _nw273[(new BigNumber(5)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(6)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(7)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(8)).toNumber()] = new BigNumber(-209);
      _nw273[(new BigNumber(9)).toNumber()] = (((_1643_v7).contains(_1639_v3)) ? ((_1643_v7).get(_1639_v3)) : (_1635_v0));
      _nw273[(new BigNumber(10)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(11)).toNumber()] = new BigNumber((_1645_v9).length);
      _nw273[(new BigNumber(12)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(13)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(14)).toNumber()] = new BigNumber(118);
      _nw273[(new BigNumber(15)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(16)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(17)).toNumber()] = new BigNumber((_1642_v6).length);
      _nw273[(new BigNumber(18)).toNumber()] = new BigNumber((_1643_v7).length);
      _nw273[(new BigNumber(19)).toNumber()] = _1635_v0;
      _nw273[(new BigNumber(20)).toNumber()] = new BigNumber(-547);
      _nw273[(new BigNumber(21)).toNumber()] = _1635_v0;
      _1646_v10 = _nw273;
      let _pat_let_tv28 = _1646_v10;
      let _pat_let_tv29 = _1646_v10;
      let _source19 = function (_pat_let23_0) {
        return function (_1649_dt__update__tmp_h1) {
          return function (_pat_let26_0) {
            return function (_1650_dt__update_hcf39_h1) {
              return _module.D9.create_DC21(_1650_dt__update_hcf39_h1);
            }(_pat_let26_0);
          }(_pat_let_tv29);
        }(_pat_let23_0);
      }(function (_pat_let24_0) {
        return function (_1647_dt__update__tmp_h0) {
          return function (_pat_let25_0) {
            return function (_1648_dt__update_hcf39_h0) {
              return _module.D9.create_DC21(_1648_dt__update_hcf39_h0);
            }(_pat_let25_0);
          }(_pat_let_tv28);
        }(_pat_let24_0);
      }(_module.D9.create_DC21(_1646_v10)));
      if (_source19.is_DC22) {
        let _1651___mcc_h0 = (_source19).cf40;
        let _1652___mcc_h1 = (_source19).cf41;
        let _1653_cf41 = _1652___mcc_h1;
        let _1654_cf40 = _1651___mcc_h0;
        let _1655_v11;
        let _init53 = function (_1656_i1) {
          return true;
        };
        let _nw274 = Array((new BigNumber(22)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw274.length); _i0_53++) {
          _nw274[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1655_v11 = _nw274;
        let _1657_v12;
        _1657_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1636_v1, _1636_v1),_1655_v11);
        let _rhs277 = (_1635_v0).isLessThanOrEqualTo(new BigNumber(302));
        let _rhs278 = _1657_v12;
        let _lhs200 = globalState;
        _lhs200.f8 = _rhs277;
        _1657_v12 = _rhs278;
        (globalState).f8 = !((_this).fm42((new BigNumber((_module.__default.fm46(globalState)).cardinality())).multipliedBy(new BigNumber(18)), _1635_v0, globalState));
        let _1658_v13;
        _1658_v13 = _dafny.Seq.UnicodeFromString("orpv");
        _1654_cf40 = ((new BigNumber((_dafny.Seq.UnicodeFromString("eadjqkvyy")).length)).multipliedBy((_this).fm41(_1658_v13, _1635_v0, _1654_cf40, (_1640_v4)[_module.__default.safeIndex(_1654_cf40, new BigNumber((_1640_v4).length))], globalState))).minus(_module.__default.safeDivisionInt(_1654_cf40, _1654_cf40));
        let _1659_v14;
        _1659_v14 = _dafny.MultiSet.fromElements(false);
        let _1660_v15;
        let _nw275 = Array((new BigNumber(4)).toNumber());
        _nw275[(_dafny.ZERO).toNumber()] = _1659_v14;
        _nw275[(_dafny.ONE).toNumber()] = _1659_v14;
        _nw275[(new BigNumber(2)).toNumber()] = _1659_v14;
        _nw275[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(!(_1639_v3), _1639_v3);
        _1660_v15 = _nw275;
        let _1661_v16;
        _1661_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1660_v15,_1639_v3);
        _1661_v16 = (_1661_v16).update(_1660_v15, (_1635_v0).isEqualTo((_dafny.ZERO).minus(_1654_cf40)));
      } else {
        let _1662___mcc_h2 = (_source19).cf39;
        let _1663_cf39 = _1662___mcc_h2;
        let _1664_v17;
        let _nw276 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1664_v17 = _nw276;
        let _index251 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1664_v17).length));
        (_1664_v17)[_index251] = _1639_v3;
        let _1665_v18;
        let _nw277 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1665_v18 = _nw277;
        let _index252 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1665_v18).length));
        (_1665_v18)[_index252] = _1641_v5;
        let _1666_v19;
        _1666_v19 = _dafny.Seq.UnicodeFromString("u");
        let _1667_v20;
        let _nw278 = Array((new BigNumber(14)).toNumber());
        _nw278[(_dafny.ZERO).toNumber()] = _1636_v1;
        _nw278[(_dafny.ONE).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(2)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(3)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(4)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(5)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(6)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-140)), ((_1668_v1) => function (_1669_i2) {
          return new BigNumber((_1668_v1).length);
        })(_1636_v1));
        _nw278[(new BigNumber(8)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(9)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(10)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(11)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(12)).toNumber()] = _1636_v1;
        _nw278[(new BigNumber(13)).toNumber()] = _1636_v1;
        _1667_v20 = _nw278;
        let _1670_v21;
        _1670_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1667_v20,new BigNumber((_dafny.Seq.UnicodeFromString("erktiyai")).length));
        let _index253 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1664_v17).length));
        let _index254 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1665_v18).length));
        let _rhs279 = !_dafny.Seq.contains(_1666_v19, _1641_v5);
        let _rhs280 = new BigNumber(625);
        let _rhs281 = _module.__default.fm1(_1639_v3, _1641_v5, globalState);
        let _rhs282 = ((((_1670_v21).contains(_1667_v20)) ? ((_1670_v21).get(_1667_v20)) : (_1635_v0))).minus(new BigNumber(461));
        let _rhs283 = _1641_v5;
        let _lhs201 = globalState;
        let _lhs202 = _1664_v17;
        let _lhs203 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1664_v17).length));
        let _lhs204 = globalState;
        let _lhs205 = _1665_v18;
        let _lhs206 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1665_v18).length));
        _1639_v3 = _rhs279;
        _lhs201.f9 = _rhs280;
        _lhs202[_lhs203] = _rhs281;
        _lhs204.f9 = _rhs282;
        _lhs205[_lhs206] = _rhs283;
        let _1671_v22;
        _1671_v22 = _dafny.Set.fromElements(_1635_v0);
        let _1672_v23;
        _1672_v23 = _module.D0.create_DC0(_1635_v0);
        let _1673_v24;
        _1673_v24 = _module.D2.create_DC5(_1666_v19);
        let _1674_v25;
        let _nw279 = new _module.C1();
        _nw279.__ctor((_1671_v22).IsProperSubsetOf(_1671_v22), (_1664_v17)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_1664_v17).length))], _1672_v23, _dafny.Seq.IsPrefixOf((_1673_v24).dtor_cf10, _1666_v19));
        _1674_v25 = _nw279;
        (globalState).f9 = (_1635_v0).plus(_1635_v0);
        (globalState).f8 = _1639_v3;
      }
      let _1675_v26;
      let _nw280 = new _module.C3();
      _nw280.__ctor(!(_1635_v0).isEqualTo(new BigNumber((_1643_v7).length)), (_1639_v3) === (_1639_v3), (_this).f11);
      _1675_v26 = _nw280;
      _1675_v26 = _1675_v26;
      _module.__default.m0(globalState);
      _1646_v10 = ((_1639_v3) ? (_1646_v10) : (_1646_v10));
      let _hi12 = _module.__default.fm0(_1639_v3, globalState);
      for (let _1676_i3 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1635_v0,_1639_v3)).Merge(_1645_v9)).length); _1676_i3.isLessThan(_hi12); _1676_i3 = _1676_i3.plus(_dafny.ONE)) {
        let _1677_v27;
        let _init54 = ((_1678_v3, _1679_i3) => function (_1680_i4) {
          return _module.D2.create_DC7(_1678_v3, (_dafny.ZERO).minus(_1679_i3), _1678_v3, _dafny.Seq.UnicodeFromString("ygpysjnm"));
        })(_1639_v3, _1676_i3);
        let _nw281 = Array((new BigNumber(25)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw281.length); _i0_54++) {
          _nw281[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1677_v27 = _nw281;
        let _1681_v28;
        _1681_v28 = _module.D11.create_DC25(_1677_v27);
        let _source20 = _1681_v28;
        if (_source20.is_DC26) {
          let _1682___mcc_h3 = (_source20).cf47;
          let _1683_cf47 = _1682___mcc_h3;
          let _1684_v29;
          _1684_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1639_v3,_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(364)), ((_1685_cf47) => function (_1686_i5) {
            return _1685_cf47;
          })(_1683_cf47))));
          let _1687_v30;
          _1687_v30 = _dafny.MultiSet.fromElements(_1683_cf47);
          _1639_v3 = ((_1684_v29).update(_1639_v3, _1687_v30)).equals(_1684_v29);
          let _1688_v31;
          let _nw282 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
          _1688_v31 = _nw282;
          let _index255 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_1688_v31).length));
          (_1688_v31)[_index255] = _dafny.Set.fromElements(_1635_v0, (_dafny.ZERO).minus(_1635_v0));
          let _index256 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_1688_v31).length));
          (_1688_v31)[_index256] = function () {
            let _coll59 = new _dafny.Set();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(428), new BigNumber(619))) {
              let _1689_v32 = _compr_59;
              if (((new BigNumber(428)).isLessThanOrEqualTo(_1689_v32)) && ((_1689_v32).isLessThan(new BigNumber(619)))) {
                _coll59.add((_1689_v32).multipliedBy((_module.D9.create_DC22(_1676_i3, _dafny.Map.Empty.slice().updateUnsafe(_1676_i3,_dafny.Seq.UnicodeFromString("kamnjg")))).dtor_cf40));
              }
            }
            return _coll59;
          }();
          let _1690_v35;
          _1690_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(((_1688_v31)[_module.__default.safeIndex(new BigNumber(81), new BigNumber((_1688_v31).length))]).length), new BigNumber(834), _1635_v0)).length),(((_1643_v7).contains(_1639_v3)) ? ((_1643_v7).get(_1639_v3)) : (_1683_cf47)));
          let _1691_v36;
          _1691_v36 = _dafny.Seq.UnicodeFromString("vusjkfsmm");
          let _1692_v37;
          _1692_v37 = _module.D2.create_DC7(_1639_v3, _1676_i3, _1639_v3, _1691_v36);
          let _1693_v39;
          let _nw283 = Array((new BigNumber(20)).toNumber());
          _nw283[(_dafny.ZERO).toNumber()] = _1639_v3;
          _nw283[(_dafny.ONE).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(2)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(3)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(4)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(5)).toNumber()] = false;
          _nw283[(new BigNumber(6)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(7)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(8)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(9)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(10)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(11)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(12)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(13)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(14)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(15)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(16)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(17)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(18)).toNumber()] = _1639_v3;
          _nw283[(new BigNumber(19)).toNumber()] = _1639_v3;
          _1693_v39 = _nw283;
          let _1694_v40;
          _1694_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1693_v39,_1693_v39);
          let _1695_v41;
          let _nw284 = Array((new BigNumber(15)).toNumber());
          _nw284[(_dafny.ZERO).toNumber()] = function () {
            let _coll60 = new _dafny.Map();
            for (const _compr_60 of (_1645_v9).Keys.Elements) {
              let _1696_v33 = _compr_60;
              if ((_1645_v9).contains(_1696_v33)) {
                _coll60.push([(_1696_v33).plus(_1676_i3),_1683_cf47]);
              }
            }
            return _coll60;
          }();
          _nw284[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_1684_v29).contains(_1639_v3)) ? ((_1684_v29).get(_1639_v3)) : (_1687_v30))).cardinality()),new BigNumber((function () {
            let _coll61 = new _dafny.Map();
            for (const _compr_61 of (_1645_v9).Keys.Elements) {
              let _1697_v34 = _compr_61;
              if ((_1645_v9).contains(_1697_v34)) {
                _coll61.push([_module.__default.safeModuloInt(_1697_v34, _1683_cf47),_1676_i3]);
              }
            }
            return _coll61;
          }()).length));
          _nw284[(new BigNumber(2)).toNumber()] = _1690_v35;
          _nw284[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1691_v36).length),new BigNumber(((_1692_v37).dtor_cf14).length));
          _nw284[(new BigNumber(4)).toNumber()] = _1690_v35;
          _nw284[(new BigNumber(5)).toNumber()] = function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(979), new BigNumber(110))) {
              let _1698_v38 = _compr_62;
              if (((new BigNumber(979)).isLessThanOrEqualTo(_1698_v38)) && ((_1698_v38).isLessThan(new BigNumber(110)))) {
                _coll62.push([_module.__default.safeModuloInt(_1698_v38, new BigNumber(-347)),_1676_i3]);
              }
            }
            return _coll62;
          }();
          _nw284[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(558),_1635_v0);
          _nw284[(new BigNumber(7)).toNumber()] = (_1690_v35);
          _nw284[(new BigNumber(8)).toNumber()] = _1690_v35;
          _nw284[(new BigNumber(9)).toNumber()] = _1690_v35;
          _nw284[(new BigNumber(10)).toNumber()] = _1690_v35;
          _nw284[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1683_cf47,_1683_cf47);
          _nw284[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1694_v40).length),_1676_i3);
          _nw284[(new BigNumber(13)).toNumber()] = _1690_v35;
          _nw284[(new BigNumber(14)).toNumber()] = _1690_v35;
          _1695_v41 = _nw284;
          let _1699_v42;
          _1699_v42 = _module.D10.create_DC23(_1695_v41);
          let _1700_v43;
          _1700_v43 = _dafny.Seq.of(_1687_v30);
          let _pat_let_tv30 = _1695_v41;
          let _rhs284 = (((_1645_v9).contains(_1635_v0)) ? ((_1645_v9).get(_1635_v0)) : (true));
          let _rhs285 = (_1691_v36)[_module.__default.safeIndex(_1676_i3, new BigNumber((_1691_v36).length))];
          let _rhs286 = function (_pat_let27_0) {
            return function (_1701_dt__update__tmp_h2) {
              return function (_pat_let28_0) {
                return function (_1702_dt__update_hcf42_h0) {
                  return _module.D10.create_DC23(_1702_dt__update_hcf42_h0);
                }(_pat_let28_0);
              }(_pat_let_tv30);
            }(_pat_let27_0);
          }(_1699_v42);
          let _rhs287 = (_dafny.Set.fromElements(_1687_v30, (_1700_v43)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_1676_i3)), new BigNumber((_1700_v43).length))])).IsProperSubsetOf(_dafny.Set.fromElements(((_dafny.MultiSet.fromElements(new BigNumber((_1636_v1).length))).update(_1635_v0, _module.__default.abs(_1676_i3))).update((_dafny.ZERO).minus(new BigNumber((_1636_v1).length)), _module.__default.abs(new BigNumber((_1691_v36).length)))));
          let _lhs207 = globalState;
          _lhs207.f8 = _rhs284;
          _1641_v5 = _rhs285;
          _1699_v42 = _rhs286;
          _1639_v3 = _rhs287;
          let _1703_v44;
          _1703_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1683_cf47,_1646_v10);
          let _1704_v45;
          _1704_v45 = _dafny.Seq.of(_1703_v44);
          let _1705_v46;
          _1705_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1646_v10,(_1704_v45)[_module.__default.safeIndex(_1683_cf47, new BigNumber((_1704_v45).length))]);
          _1703_v44 = ((((_1705_v46).contains(_1646_v10)) ? ((_1705_v46).get(_1646_v10)) : (_1703_v44))).update(_1635_v0, _1646_v10);
        } else if (_source20.is_DC25) {
          let _1706___mcc_h4 = (_source20).cf46;
          let _1707_cf46 = _1706___mcc_h4;
          let _index257 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1646_v10).length));
          (_1646_v10)[_index257] = (_dafny.ZERO).minus(_1676_i3);
          let _index258 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1646_v10).length));
          (_1646_v10)[_index258] = _1676_i3;
          let _1708_v47;
          _1708_v47 = _dafny.Seq.UnicodeFromString("f");
          let _1709_v48;
          _1709_v48 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wj"), _dafny.Seq.Concat(_1708_v47, _1708_v47));
          _1709_v48 = _1709_v48;
          let _1710_v49;
          _1710_v49 = _dafny.MultiSet.fromElements(_1639_v3);
          (globalState).f9 = ((_1639_v3) ? ((((_1710_v49).contains(_1639_v3)) ? ((_1710_v49).get(_1639_v3)) : (_1676_i3))) : (_1635_v0));
          let _1711_v50;
          let _init55 = ((_1712_v47) => function (_1713_i6) {
            return _module.__default.safeDivisionInt(_1713_i6, new BigNumber((_1712_v47).length));
          })(_1708_v47);
          let _nw285 = Array((new BigNumber(19)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw285.length); _i0_55++) {
            _nw285[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _1711_v50 = _nw285;
          let _1714_v51;
          _1714_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1711_v50,_1639_v3);
          let _1715_v52;
          _1715_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1714_v51,_1639_v3);
          _1715_v52 = (_1715_v52).update((_dafny.Map.Empty.slice().updateUnsafe(_1711_v50,_1639_v3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1711_v50,_1639_v3)), false);
        } else {
          let _1716___mcc_h5 = (_source20).cf48;
          let _1717_cf48 = _1716___mcc_h5;
          let _1718_v53;
          _1718_v53 = _dafny.MultiSet.fromElements(new BigNumber(533), (_1636_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(_1635_v0), new BigNumber((_1636_v1).length))]);
          let _1719_v54;
          _1719_v54 = _1718_v53;
          let _1720_v55;
          _1720_v55 = _dafny.Seq.UnicodeFromString("offr");
          let _1721_v56;
          _1721_v56 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), ((_1722_v4) => function (_1723_i7) {
            return new BigNumber((_dafny.MultiSet.FromArray(_1722_v4)).cardinality());
          })(_1640_v4)));
          let _1724_v57;
          _1724_v57 = _dafny.Seq.of(_1718_v53, _1718_v53);
          let _1725_v58;
          _1725_v58 = _dafny.Seq.of(_1718_v53, _1718_v53, (_1724_v57)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1676_i3, new BigNumber(698))).length), new BigNumber((_1724_v57).length))], _1718_v53, _dafny.MultiSet.fromElements(_1676_i3, _1635_v0));
          let _1726_v59;
          let _nw286 = Array((new BigNumber(20)).toNumber());
          _nw286[(_dafny.ZERO).toNumber()] = (_1718_v53).Intersect(_1718_v53);
          _nw286[(_dafny.ONE).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(2)).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(3)).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(4)).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(5)).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(256));
          _nw286[(new BigNumber(7)).toNumber()] = (_1718_v53).Difference(_1718_v53);
          _nw286[(new BigNumber(8)).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(9)).toNumber()] = (_1719_v54);
          _nw286[(new BigNumber(10)).toNumber()] = (_1718_v53).Intersect(_1718_v53);
          _nw286[(new BigNumber(11)).toNumber()] = (_1718_v53).Difference(_dafny.MultiSet.fromElements(new BigNumber((_1720_v55).length), _1676_i3, _1676_i3, new BigNumber((_1643_v7).length), (_dafny.ZERO).minus(_1676_i3)));
          _nw286[(new BigNumber(12)).toNumber()] = _module.__default.fm56(_1639_v3, _1676_i3, globalState);
          _nw286[(new BigNumber(13)).toNumber()] = _1718_v53;
          _nw286[(new BigNumber(14)).toNumber()] = (_1718_v53).update(_module.__default.fm0(_1639_v3, globalState), _module.__default.abs(_module.__default.fm0(!(_1639_v3), globalState)));
          _nw286[(new BigNumber(15)).toNumber()] = ((_1639_v3) ? (_dafny.MultiSet.fromElements(_1635_v0, new BigNumber((_1721_v56).cardinality()), _1676_i3)) : (_1718_v53));
          _nw286[(new BigNumber(16)).toNumber()] = (_1718_v53).Union((_1724_v57)[_module.__default.safeIndex(_1635_v0, new BigNumber((_1724_v57).length))]);
          _nw286[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm54(_1639_v3, _1639_v3, !(_1639_v3), globalState));
          _nw286[(new BigNumber(18)).toNumber()] = (_1718_v53).Intersect(_1718_v53);
          _nw286[(new BigNumber(19)).toNumber()] = _1718_v53;
          _1726_v59 = _nw286;
          let _index259 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1726_v59).length));
          (_1726_v59)[_index259] = _1718_v53;
          let _index260 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1726_v59).length));
          (_1726_v59)[_index260] = _1718_v53;
          let _1727_v60;
          let _nw287 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1727_v60 = _nw287;
          let _1728_v61;
          _1728_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1646_v10,_1639_v3);
          let _index261 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1727_v60).length));
          (_1727_v60)[_index261] = (((_1728_v61).contains(_1646_v10)) ? ((_1728_v61).get(_1646_v10)) : (_1639_v3));
          let _index262 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1727_v60).length));
          (_1727_v60)[_index262] = (_1644_v8).IsProperSubsetOf(_1644_v8);
          _1639_v3 = _1639_v3;
          let _1729_v62;
          _1729_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1676_i3,_1676_i3);
          let _index263 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1646_v10).length));
          (_1646_v10)[_index263] = (((_1729_v62).contains(_1676_i3)) ? ((_1729_v62).get(_1676_i3)) : (new BigNumber(407)));
          let _1730_v63;
          _1730_v63 = _dafny.Seq.of(_1640_v4);
          let _1731_v64;
          _1731_v64 = _dafny.MultiSet.fromElements((_1727_v60)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_1727_v60).length))], (new BigNumber((_module.__default.fm2((_dafny.ZERO).minus(_1676_i3), ((_1726_v59)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1726_v59).length))]).update(_1635_v0, _module.__default.abs(_1676_i3)), _1729_v62, globalState)).length)).isLessThan(_1676_i3), (_1727_v60)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_1727_v60).length))]);
          let _1732_v65;
          _1732_v65 = _dafny.Set.fromElements(new BigNumber((_1720_v55).length));
          let _index264 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1646_v10).length));
          let _rhs288 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1640_v4, _module.__default.fm44(_1730_v63, new _dafny.CodePoint('e'.codePointAt(0)), globalState)), _1640_v4);
          let _rhs289 = (((_1731_v64).contains((_1732_v65).IsSubsetOf(_1732_v65))) ? ((_1731_v64).get((_1732_v65).IsSubsetOf(_1732_v65))) : (_module.__default.safeDivisionInt(_1635_v0, (_dafny.ZERO).minus(new BigNumber((_1720_v55).length)))));
          let _lhs208 = _1646_v10;
          let _lhs209 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1646_v10).length));
          _1640_v4 = _rhs288;
          _lhs208[_lhs209] = _rhs289;
        }
        _1639_v3 = _1639_v3;
        let _1733_v66;
        _1733_v66 = _dafny.MultiSet.fromElements(_1676_i3);
        let _1734_v67;
        _1734_v67 = _dafny.MultiSet.fromElements(_1733_v66, _1733_v66, (_1733_v66).update(_1635_v0, _module.__default.abs(_1676_i3)));
        let _rhs290 = _1639_v3;
        let _rhs291 = _1734_v67;
        let _lhs210 = globalState;
        _lhs210.f8 = _rhs290;
        _1734_v67 = _rhs291;
        let _1735_v68;
        _1735_v68 = _dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), _1641_v5);
        let _1736_v69;
        _1736_v69 = _1735_v68;
        _1643_v7 = (_1643_v7).update(_1639_v3, (_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(_1736_v69)).length)).minus(new BigNumber(504))));
      }
      let _1737_v70;
      _1737_v70 = _dafny.MultiSet.fromElements(_1639_v3, !(_1639_v3), _1639_v3);
      r0 = _1737_v70;
      return r0;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f11 = [];
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _1738_v0;
      _1738_v0 = new BigNumber(-540);
      let _1739_v1;
      _1739_v1 = _dafny.MultiSet.fromElements(_1738_v0, _1738_v0, _module.__default.fm0(p0, globalState));
      (globalState).f9 = _module.__default.fm0((_1739_v1).IsSubsetOf(_1739_v1), globalState);
      (globalState).f8 = p0;
      (globalState).f8 = p0;
      let _1740_v2;
      _1740_v2 = new _dafny.CodePoint('y'.codePointAt(0));
      let _1741_v3;
      _1741_v3 = _module.D2.create_DC8(_1738_v0, _1740_v2);
      let _1742_v4;
      _1742_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(817),(_1741_v3).dtor_cf16);
      _1742_v4 = (_1742_v4).update(_1738_v0, _1740_v2);
      _module.__default.m0(globalState);
      let _1743_v5;
      let _nw288 = new _module.C0();
      _nw288.__ctor(p0);
      _1743_v5 = _nw288;
      r0 = _1738_v0;
      let _pat_let_tv31 = _1743_v5;
      let _pat_let_tv32 = _1743_v5;
      let _pat_let_tv33 = _1743_v5;
      let _pat_let_tv34 = _1743_v5;
      let _pat_let_tv35 = _1743_v5;
      let _pat_let_tv36 = _1743_v5;
      r1 = function (_source21) {
        if (_source21.is_DC6) {
          return (_dafny.Set.fromElements(!((_pat_let_tv31).f17))).Union(_dafny.Set.fromElements((_pat_let_tv32).f17, false));
        } else if (_source21.is_DC7) {
          let _1744___mcc_h0 = (_source21).cf11;
          let _1745___mcc_h1 = (_source21).cf12;
          let _1746___mcc_h2 = (_source21).cf13;
          let _1747___mcc_h3 = (_source21).cf14;
          let _1748_cf14 = _1747___mcc_h3;
          let _1749_cf13 = _1746___mcc_h2;
          let _1750_cf12 = _1745___mcc_h1;
          let _1751_cf11 = _1744___mcc_h0;
          return _dafny.Set.fromElements(true, true);
        } else if (_source21.is_DC8) {
          let _1752___mcc_h4 = (_source21).cf15;
          let _1753___mcc_h5 = (_source21).cf16;
          let _1754_cf16 = _1753___mcc_h5;
          let _1755_cf15 = _1752___mcc_h4;
          return (_dafny.Set.fromElements((_pat_let_tv33).f17)).Intersect(_dafny.Set.fromElements((_pat_let_tv34).f17, (_pat_let_tv35).f17));
        } else {
          let _1756___mcc_h6 = (_source21).cf10;
          let _1757_cf10 = _1756___mcc_h6;
          return (_dafny.Set.fromElements((_pat_let_tv36).f17)).Intersect(_dafny.Set.fromElements(false));
        }
      }(_1741_v3);
      return [r0, r1];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _module.D2.Default();
      let r3 = false;
      let _1758_v0;
      _1758_v0 = true;
      if (_1758_v0) {
        let _1759_v2;
        let _init56 = ((_1760_v0) => function (_1761_i0) {
          return function () {
            let _coll63 = new _dafny.Set();
            for (const _compr_63 of _dafny.IntegerRange(new BigNumber(-866), new BigNumber(290))) {
              let _1762_v1 = _compr_63;
              if (((new BigNumber(-866)).isLessThanOrEqualTo(_1762_v1)) && ((_1762_v1).isLessThan(new BigNumber(290)))) {
                _coll63.add((_1762_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1760_v0,_1760_v0)).length))).length)));
              }
            }
            return _coll63;
          }();
        })(_1758_v0);
        let _nw289 = Array((new BigNumber(10)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw289.length); _i0_56++) {
          _nw289[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _1759_v2 = _nw289;
        let _index265 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_1759_v2).length));
        (_1759_v2)[_index265] = _dafny.Set.fromElements(p0, p0);
        let _1763_v3;
        _1763_v3 = _dafny.Set.fromElements(p0);
        let _index266 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_1759_v2).length));
        (_1759_v2)[_index266] = (_1763_v3).Difference((_1763_v3).Union(_1763_v3));
        let _1764_v4;
        let _init57 = ((_1765_v0) => function (_1766_i1) {
          return _1765_v0;
        })(_1758_v0);
        let _nw290 = Array((new BigNumber(4)).toNumber());
        for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw290.length); _i0_57++) {
          _nw290[_i0_57] = _init57(new BigNumber(_i0_57));
        }
        _1764_v4 = _nw290;
        let _index267 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1764_v4).length));
        (_1764_v4)[_index267] = true;
        let _index268 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1764_v4).length));
        (_1764_v4)[_index268] = (p0).isLessThan(p0);
        r1 = p0;
        let _1767_v5;
        _1767_v5 = new _dafny.CodePoint('h'.codePointAt(0));
        r0 = _1767_v5;
        let _1768_v6;
        _1768_v6 = _dafny.Seq.UnicodeFromString("jh");
        (globalState).f9 = new BigNumber((_dafny.Seq.update(_1768_v6, _module.__default.safeIndex(p0, new BigNumber((_1768_v6).length)), _1767_v5)).length);
      } else {
        _module.__default.m0(globalState);
        let _1769_v7;
        _1769_v7 = _dafny.Seq.UnicodeFromString("cyw");
        if (((!_dafny.areEqual(_1769_v7, _1769_v7)) ? (_1758_v0) : (_1758_v0))) {
          let _1770_v8;
          _1770_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v0,p0);
          (globalState).f8 = !((p0).plus(p0)).isEqualTo(new BigNumber(((_1770_v8).Merge(_1770_v8)).length));
          r1 = new BigNumber(-719);
          (globalState).f8 = _1758_v0;
          let _1771_v9;
          _1771_v9 = _module.D2.create_DC7(_1758_v0, p0, _1758_v0, _1769_v7);
          let _1772_v10;
          _1772_v10 = _dafny.MultiSet.fromElements(_1758_v0);
          let _rhs292 = !((_1771_v9).dtor_cf13) || (_1758_v0);
          let _rhs293 = _1758_v0;
          let _rhs294 = ((_dafny.ZERO).minus(p0)).multipliedBy((_dafny.ZERO).minus((p0).minus((_dafny.ZERO).minus(new BigNumber((_1772_v10).cardinality())))));
          let _rhs295 = p0;
          let _lhs211 = globalState;
          let _lhs212 = globalState;
          let _lhs213 = globalState;
          _lhs211.f8 = _rhs292;
          _lhs212.f8 = _rhs293;
          _lhs213.f9 = _rhs294;
          r1 = _rhs295;
          _1758_v0 = _dafny.Seq.contains(_1769_v7, new _dafny.CodePoint('u'.codePointAt(0)));
        } else {
          (globalState).f9 = p0;
          r1 = p0;
          let _1773_v11;
          let _nw291 = Array((new BigNumber(9)).toNumber()).fill([]);
          _1773_v11 = _nw291;
          let _index269 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1773_v11).length));
          (_1773_v11)[_index269] = p1;
          let _1774_v12;
          let _nw292 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1774_v12 = _nw292;
          let _index270 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1774_v12).length));
          (_1774_v12)[_index270] = !(_module.__default.fm0(_1758_v0, globalState)).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(p0, p0, p0, p0, p0)).cardinality()));
          let _index271 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1773_v11).length));
          let _index272 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1774_v12).length));
          let _rhs296 = p1;
          let _rhs297 = _1758_v0;
          let _lhs214 = _1773_v11;
          let _lhs215 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1773_v11).length));
          let _lhs216 = _1774_v12;
          let _lhs217 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1774_v12).length));
          _lhs214[_lhs215] = _rhs296;
          _lhs216[_lhs217] = _rhs297;
          let _1775_v13;
          _1775_v13 = _dafny.Seq.of(p0, p0);
          (globalState).f9 = (_1775_v13)[_module.__default.safeIndex(p0, new BigNumber((_1775_v13).length))];
          let _1776_v14;
          _1776_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v0,p0);
          _1776_v14 = (_1776_v14).update((_1758_v0) || ((_1774_v12)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_1774_v12).length))]), p0);
        }
        (globalState).f9 = _module.__default.safeModuloInt(p0, new BigNumber(879));
        let _1777_v15;
        _1777_v15 = _dafny.MultiSet.fromElements(_1758_v0, ((_1758_v0) ? (false) : (_1758_v0)), _1758_v0);
        _1777_v15 = _1777_v15;
        _1758_v0 = _1758_v0;
      }
      let _hi13 = _module.__default.fm0(_1758_v0, globalState);
      for (let _1778_i2 = new BigNumber((_dafny.Seq.UnicodeFromString("xpnbhaeoo")).length); _1778_i2.isLessThan(_hi13); _1778_i2 = _1778_i2.plus(_dafny.ONE)) {
        let _source22 = _module.D4.create_DC10(_module.__default.fm16(globalState));
        if (_source22.is_DC11) {
          let _1779___mcc_h0 = (_source22).cf19;
          let _1780___mcc_h1 = (_source22).cf20;
          let _1781___mcc_h2 = (_source22).cf21;
          let _1782_cf21 = _1781___mcc_h2;
          let _1783_cf20 = _1780___mcc_h1;
          let _1784_cf19 = _1779___mcc_h0;
          (globalState).f9 = p0;
          let _1785_v16;
          _1785_v16 = _dafny.Seq.of(_1782_cf21);
          let _1786_v17;
          _1786_v17 = _dafny.Seq.UnicodeFromString("fsyimuwr");
          let _1787_v18;
          _1787_v18 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1788_v19;
          _1788_v19 = _module.D2.create_DC7(_1758_v0, _1778_i2, (_1785_v16)[_module.__default.safeIndex(p0, new BigNumber((_1785_v16).length))], _dafny.Seq.Concat(_dafny.Seq.update(_1786_v17, _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1786_v17).length)), _1787_v18), _1786_v17));
          _1788_v19 = _module.__default.fm17(globalState);
          let _index273 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((p1).length));
          (p1)[_index273] = _1784_cf19;
          let _1789_v20;
          let _nw293 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1789_v20 = _nw293;
          let _1790_v21;
          _1790_v21 = _dafny.MultiSet.fromElements(_1789_v20, _1789_v20);
          let _index274 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((p1).length));
          (p1)[_index274] = (_dafny.ZERO).minus(new BigNumber((_1790_v21).cardinality()));
          (globalState).f9 = p0;
        } else if (_source22.is_DC12) {
          let _1791___mcc_h3 = (_source22).cf22;
          let _1792___mcc_h4 = (_source22).cf23;
          let _1793___mcc_h5 = (_source22).cf24;
          let _1794_cf24 = _1793___mcc_h5;
          let _1795_cf23 = _1792___mcc_h4;
          let _1796_cf22 = _1791___mcc_h3;
          let _1797_v22;
          _1797_v22 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1798_v23;
          _1798_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(),_module.__default.fm1(_1758_v0, _1797_v22, globalState));
          let _1799_v24;
          _1799_v24 = _module.D2.create_DC6();
          let _1800_v25;
          let _nw294 = new _module.C0();
          _nw294.__ctor(_1795_cf23);
          _1800_v25 = _nw294;
          let _1801_v26;
          _1801_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1797_v22,_1800_v25);
          let _1802_v27;
          _1802_v27 = _dafny.Seq.UnicodeFromString("nhbbrfs");
          _1798_v23 = (_1798_v23).update(_1799_v24, (new BigNumber((_1801_v26).length)).isLessThan(new BigNumber((_1802_v27).length)));
          let _1803_v28;
          let _nw295 = Array((new BigNumber(4)).toNumber());
          _nw295[(_dafny.ZERO).toNumber()] = p0;
          _nw295[(_dafny.ONE).toNumber()] = p0;
          _nw295[(new BigNumber(2)).toNumber()] = new BigNumber(588);
          _nw295[(new BigNumber(3)).toNumber()] = (new BigNumber(20)).multipliedBy(new BigNumber(633));
          _1803_v28 = _nw295;
          _1803_v28 = _1803_v28;
          (globalState).f9 = p0;
          let _1804_v29;
          _1804_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(782),_1778_i2);
          let _1805_v30;
          _1805_v30 = _dafny.MultiSet.fromElements(new BigNumber((_1804_v29).length));
          let _1806_v31;
          _1806_v31 = _dafny.Set.fromElements(_1805_v30);
          let _1807_v32;
          _1807_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1778_i2,false);
          let _1808_v33;
          _1808_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1796_cf22,(((_1807_v32).contains(_1778_i2)) ? ((_1807_v32).get(_1778_i2)) : (_1795_cf23)));
          let _1809_v34;
          _1809_v34 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), ((_1810_v22) => function (_1811_i3) {
            return _1810_v22;
          })(_1797_v22)));
          let _1812_v35;
          _1812_v35 = _dafny.Seq.of(_1778_i2);
          let _1813_v36;
          _1813_v36 = _dafny.Seq.of(_1796_cf22, (_1806_v31).IsDisjointFrom(_module.__default.fm18(_1808_v33, _1809_v34, _1797_v22, (_1812_v35)[_module.__default.safeIndex(_1778_i2, new BigNumber((_1812_v35).length))], globalState)), (_1805_v30).IsProperSubsetOf(_1805_v30));
          let _rhs298 = _1800_v25;
          let _rhs299 = _1758_v0;
          let _rhs300 = _module.__default.fm1((_1795_cf23) === (_1794_cf24), new _dafny.CodePoint('a'.codePointAt(0)), globalState);
          let _rhs301 = _1813_v36;
          let _rhs302 = (_1800_v25).f17;
          let _lhs218 = globalState;
          _1800_v25 = _rhs298;
          _lhs218.f8 = _rhs299;
          r3 = _rhs300;
          _1813_v36 = _rhs301;
          _1794_cf24 = _rhs302;
        } else {
          let _1814___mcc_h6 = (_source22).cf18;
          let _1815_cf18 = _1814___mcc_h6;
          let _1816_v37;
          _1816_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v0,_1778_i2);
          _1816_v37 = (_1816_v37).update(_1758_v0, _1778_i2);
          let _1817_v38;
          _1817_v38 = _dafny.Seq.of(new BigNumber(799), p0);
          _1817_v38 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-124)), ((_1818_i2) => function (_1819_i4) {
            return _1818_i2;
          })(_1778_i2)), _module.__default.safeIndex(_module.__default.safeModuloInt(p0, p0), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-124)), ((_1820_i2) => function (_1821_i4) {
            return _1820_i2;
          })(_1778_i2))).length)), (_1817_v38)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1817_v38).length))]);
          (globalState).f9 = _1778_i2;
          let _1822_v39;
          let _nw296 = new _module.C0();
          _nw296.__ctor(!(_dafny.Seq.IsProperPrefixOf(_1817_v38, _1817_v38)));
          _1822_v39 = _nw296;
        }
        let _1823_v40;
        let _nw297 = new _module.C2();
        _nw297.__ctor((_this).f11);
        _1823_v40 = _nw297;
        let _1824_v41;
        _1824_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1823_v40);
        r1 = new BigNumber(((_1824_v41).Merge(_1824_v41)).length);
        let _1825_v42;
        _1825_v42 = _dafny.MultiSet.fromElements(_1758_v0, _1758_v0);
        let _1826_v43;
        _1826_v43 = _dafny.MultiSet.fromElements(_1825_v42, _1825_v42);
        let _1827_v44;
        _1827_v44 = _dafny.Seq.of(_1758_v0);
        let _1828_v45;
        _1828_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1758_v0);
        let _1829_v46;
        _1829_v46 = _dafny.Seq.of(new BigNumber((_1828_v45).length), _1778_i2);
        let _1830_v47;
        _1830_v47 = _dafny.Seq.of(_1829_v46);
        let _1831_v48;
        _1831_v48 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_1830_v47, _module.__default.safeIndex(new BigNumber((_1827_v44).length), new BigNumber((_1830_v47).length)), _1829_v46)).length), _1778_i2);
        (globalState).f9 = (((_1826_v43).contains(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1827_v44, _1827_v44)))) ? ((_1826_v43).get(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1827_v44, _1827_v44)))) : (new BigNumber((_1831_v48).length)));
        r3 = _module.__default.fm1(_1758_v0, new _dafny.CodePoint('h'.codePointAt(0)), globalState);
      }
      let _1832_v49;
      _1832_v49 = _dafny.Seq.UnicodeFromString("iaxo");
      let _1833_v50;
      _1833_v50 = _dafny.MultiSet.fromElements(new BigNumber((_1832_v49).length));
      if (((_1833_v50).Intersect(_1833_v50)).IsProperSubsetOf(_1833_v50)) {
        let _1834_v51;
        _1834_v51 = _module.D0.create_DC0(p0);
        let _1835_v52;
        let _nw298 = new _module.C1();
        _nw298.__ctor(_1758_v0, !(!(_1758_v0)), _1834_v51, !(!(_1758_v0) || (_1758_v0)));
        _1835_v52 = _nw298;
        _1832_v49 = _dafny.Seq.UnicodeFromString("gkhk");
        let _1836_v53;
        let _nw299 = Array((new BigNumber(10)).toNumber());
        _nw299[(_dafny.ZERO).toNumber()] = p1;
        _nw299[(_dafny.ONE).toNumber()] = p1;
        _nw299[(new BigNumber(2)).toNumber()] = p1;
        _nw299[(new BigNumber(3)).toNumber()] = p1;
        _nw299[(new BigNumber(4)).toNumber()] = p1;
        _nw299[(new BigNumber(5)).toNumber()] = p1;
        _nw299[(new BigNumber(6)).toNumber()] = p1;
        _nw299[(new BigNumber(7)).toNumber()] = p1;
        _nw299[(new BigNumber(8)).toNumber()] = p1;
        _nw299[(new BigNumber(9)).toNumber()] = p1;
        _1836_v53 = _nw299;
        let _index275 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1836_v53).length));
        (_1836_v53)[_index275] = p1;
        let _1837_v54;
        _1837_v54 = _dafny.Seq.of((_1835_v52).f20, _1758_v0, (_1835_v52).f19, _1758_v0);
        let _index276 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((p1).length));
        (p1)[_index276] = new BigNumber((_1837_v54).length);
        let _1838_v55;
        _1838_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v0,p1);
        let _index277 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1836_v53).length));
        (_1836_v53)[_index277] = (((_1838_v55).contains((_1837_v54)[_module.__default.safeIndex(p0, new BigNumber((_1837_v54).length))])) ? ((_1838_v55).get((_1837_v54)[_module.__default.safeIndex(p0, new BigNumber((_1837_v54).length))])) : (p1));
        let _1839_v56;
        _1839_v56 = _dafny.MultiSet.fromElements(_1758_v0, (_1835_v52).f19, !((_1835_v52).f19));
        let _1840_v57;
        _1840_v57 = _module.D0.create_DC2(_1839_v56, (_1835_v52).f19, new BigNumber(325), false);
        let _1841_v58;
        _1841_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1840_v57);
        let _index278 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1836_v53).length));
        let _index279 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((p1).length));
        let _index280 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1836_v53).length));
        let _rhs303 = p1;
        let _rhs304 = (new BigNumber((_1841_v58).length)).multipliedBy(_module.__default.fm0((_1835_v52).f19, globalState));
        let _rhs305 = p1;
        let _lhs219 = _1836_v53;
        let _lhs220 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1836_v53).length));
        let _lhs221 = p1;
        let _lhs222 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((p1).length));
        let _lhs223 = _1836_v53;
        let _lhs224 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1836_v53).length));
        _lhs219[_lhs220] = _rhs303;
        _lhs221[_lhs222] = _rhs304;
        _lhs223[_lhs224] = _rhs305;
        let _1842_v59;
        _1842_v59 = new _dafny.CodePoint('j'.codePointAt(0));
        _1758_v0 = _dafny.Seq.contains(_1832_v49, _1842_v59);
        (globalState).f9 = _module.__default.safeDivisionInt(p0, new BigNumber(438));
      } else {
        let _1843_v60;
        _1843_v60 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1758_v0);
        let _rhs306 = (((_1843_v60).contains(p1)) ? ((_1843_v60).get(p1)) : (_1758_v0));
        let _rhs307 = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p0));
        let _lhs225 = globalState;
        r3 = _rhs306;
        _lhs225.f9 = _rhs307;
        let _1844_v61;
        _1844_v61 = new _dafny.CodePoint('t'.codePointAt(0));
        _1758_v0 = _module.__default.fm1(_1758_v0, _1844_v61, globalState);
        let _1845_v62;
        let _nw300 = new _module.C3();
        _nw300.__ctor(_1758_v0, _1758_v0, (_this).f11);
        _1845_v62 = _nw300;
        _1845_v62 = _1845_v62;
        let _1846_v63;
        let _nw301 = Array((new BigNumber(16)).toNumber()).fill(false);
        _1846_v63 = _nw301;
        let _index281 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length));
        (_1846_v63)[_index281] = _module.__default.fm1(_module.__default.fm1(_1758_v0, _1844_v61, globalState), _1844_v61, globalState);
        let _1847_v64;
        _1847_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v0,true);
        let _1848_v65;
        _1848_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1847_v64,_1758_v0);
        let _1849_v66;
        _1849_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1848_v65).length),p0);
        let _index282 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length));
        (_1846_v63)[_index282] = ((((_1849_v66).contains(new BigNumber((_1849_v66).length))) ? ((_1849_v66).get(new BigNumber((_1849_v66).length))) : (new BigNumber(-893)))).isLessThanOrEqualTo(p0);
        let _1850_v67;
        _1850_v67 = _dafny.MultiSet.fromElements(p1, p1, p1);
        if ((_1850_v67).IsSubsetOf(_1850_v67)) {
          let _1851_v68;
          _1851_v68 = _dafny.Seq.of(new BigNumber(790), new BigNumber((_1832_v49).length));
          (globalState).f8 = (new BigNumber((_1851_v68).length)).isLessThanOrEqualTo(new BigNumber((_1851_v68).length));
          let _index283 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length));
          (_1846_v63)[_index283] = (_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))];
          let _1852_v69;
          let _init58 = ((_1853_v64) => function (_1854_i5) {
            return _1853_v64;
          })(_1847_v64);
          let _nw302 = Array((new BigNumber(24)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw302.length); _i0_58++) {
            _nw302[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1852_v69 = _nw302;
          let _index284 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((p1).length));
          (p1)[_index284] = (_dafny.ZERO).minus(new BigNumber((_1851_v68).length));
          let _1855_v70;
          _1855_v70 = _dafny.MultiSet.fromElements((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], (_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], !(_1758_v0));
          let _index285 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length));
          let _index286 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((p1).length));
          let _rhs308 = _1852_v69;
          let _rhs309 = _1758_v0;
          let _rhs310 = ((!(!((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))]) || (false))) ? (_dafny.Seq.Concat(_1832_v49, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ouk"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("ouk")).length)), _1844_v61), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ouk"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("ouk")).length)), _1844_v61)).length)), _1844_v61))) : (_1832_v49));
          let _rhs311 = _module.__default.fm1((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], _1844_v61, globalState);
          let _rhs312 = ((((_1855_v70).contains((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))])) ? ((_1855_v70).get((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))])) : (new BigNumber(448)))).minus(p0);
          let _lhs226 = _1846_v63;
          let _lhs227 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length));
          let _lhs228 = globalState;
          let _lhs229 = p1;
          let _lhs230 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((p1).length));
          _1852_v69 = _rhs308;
          _lhs226[_lhs227] = _rhs309;
          _1832_v49 = _rhs310;
          _lhs228.f8 = _rhs311;
          _lhs229[_lhs230] = _rhs312;
          _1847_v64 = (_1847_v64).update((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], _1758_v0);
          (globalState).f9 = p0;
        } else {
          let _1856_v71;
          _1856_v71 = _dafny.Seq.of(_module.__default.fm1(_1758_v0, _1844_v61, globalState));
          let _1857_v72;
          _1857_v72 = _dafny.Set.fromElements(_1847_v64, (_1847_v64).Merge(_1847_v64), ((!(_1758_v0)) ? (_1847_v64) : (_1847_v64)), ((!(_module.__default.fm1((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], _1844_v61, globalState))) ? (_1847_v64) : (_module.__default.fm34(_1758_v0, (_dafny.MultiSet.FromArray(_1856_v71)).update(_1758_v0, _module.__default.abs(new BigNumber((_1832_v49).length))), globalState))), _1847_v64);
          let _1858_v73;
          _1858_v73 = _dafny.Seq.of(_1857_v72);
          _1857_v72 = (((_1858_v73)[_module.__default.safeIndex(p0, new BigNumber((_1858_v73).length))]).Difference(_1857_v72)).Intersect(_dafny.Set.fromElements(_1847_v64, _1847_v64, _1847_v64, _dafny.Map.Empty.slice().updateUnsafe((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))],(_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))])));
          let _1859_v74;
          _1859_v74 = _dafny.Set.fromElements(new BigNumber(50));
          _1859_v74 = (((_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))]) ? (_dafny.Set.fromElements(new BigNumber(806))) : (_1859_v74));
          let _1860_v75;
          _1860_v75 = _dafny.Seq.of(_1846_v63, _1846_v63);
          let _1861_v76;
          _1861_v76 = _dafny.Seq.of(((false) ? ((_1860_v75)[_module.__default.safeIndex(p0, new BigNumber((_1860_v75).length))]) : (_1846_v63)), _1846_v63);
          _1846_v63 = (_1860_v75)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber(588), new BigNumber((_1859_v74).length)), new BigNumber((_1860_v75).length))];
          let _1862_v77;
          _1862_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1844_v61,_1832_v49);
          _1832_v49 = _dafny.Seq.update((((_1862_v77).contains(_module.__default.fm3(p0, (_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], globalState))) ? ((_1862_v77).get(_module.__default.fm3(p0, (_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], globalState))) : (_dafny.Seq.Concat(_1832_v49, _dafny.Seq.update(_1832_v49, _module.__default.safeIndex(p0, new BigNumber((_1832_v49).length)), _1844_v61)))), _module.__default.safeIndex(_module.__default.safeDivisionInt(p0, p0), new BigNumber(((((_1862_v77).contains(_module.__default.fm3(p0, (_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], globalState))) ? ((_1862_v77).get(_module.__default.fm3(p0, (_1846_v63)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_1846_v63).length))], globalState))) : (_dafny.Seq.Concat(_1832_v49, _dafny.Seq.update(_1832_v49, _module.__default.safeIndex(p0, new BigNumber((_1832_v49).length)), _1844_v61))))).length)), _1844_v61);
          _1856_v71 = _1856_v71;
        }
      }
      if (((_dafny.ZERO).minus(p0)).isLessThan(p0)) {
        (globalState).f9 = (p0).minus(p0);
        let _1863_v78;
        _1863_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v0,_1758_v0);
        let _1864_v79;
        _1864_v79 = new _dafny.CodePoint('m'.codePointAt(0));
        if (!(!((((_1863_v78).contains(_1758_v0)) ? ((_1863_v78).get(_1758_v0)) : (_module.__default.fm1(_1758_v0, _1864_v79, globalState)))) || (_1758_v0)) || (_1758_v0)) {
          r3 = _1758_v0;
          (globalState).f9 = p0;
          let _1865_v80;
          _1865_v80 = _module.D0.create_DC0(new BigNumber((_module.__default.fm50(_1864_v79, _1758_v0, _1758_v0, globalState)).length));
          let _1866_v81;
          _1866_v81 = _dafny.Set.fromElements(_1758_v0, _1758_v0, _1758_v0);
          let _1867_v82;
          _1867_v82 = _module.D4.create_DC10(_1866_v81);
          let _1868_v83;
          let _nw303 = new _module.C1();
          _nw303.__ctor(_1758_v0, _1758_v0, _1865_v80, ((_1867_v82).dtor_cf18).IsProperSubsetOf(_1866_v81));
          _1868_v83 = _nw303;
          let _1869_v84;
          _1869_v84 = _dafny.MultiSet.fromElements(_1758_v0);
          _1869_v84 = _1869_v84;
          _1758_v0 = false;
        } else {
          let _rhs313 = p0;
          let _rhs314 = _module.__default.safeDivisionInt(p0, _module.__default.fm0(_1758_v0, globalState));
          let _rhs315 = _dafny.Seq.Concat(_1832_v49, _dafny.Seq.Concat(_1832_v49, _1832_v49));
          let _lhs231 = globalState;
          _lhs231.f9 = _rhs313;
          r1 = _rhs314;
          _1832_v49 = _rhs315;
          (globalState).f9 = (_dafny.ZERO).minus(p0);
          let _1870_v85;
          _1870_v85 = _dafny.Set.fromElements(_1758_v0);
          let _1871_v86;
          _1871_v86 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p0, new BigNumber(-685)),_1758_v0);
          let _1872_v87;
          let _nw304 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1872_v87 = _nw304;
          let _1873_v88;
          _1873_v88 = _dafny.MultiSet.fromElements(_1872_v87, _1872_v87, _1872_v87);
          let _1874_v89;
          _1874_v89 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(p0));
          let _rhs316 = _1758_v0;
          let _rhs317 = _1758_v0;
          let _rhs318 = (((_module.__default.fm1(_1758_v0, _1864_v79, globalState)) ? (_1870_v85) : (_1870_v85))).contains(_1758_v0);
          let _rhs319 = (((_1871_v86).contains(new BigNumber((_1832_v49).length))) ? ((_1871_v86).get(new BigNumber((_1832_v49).length))) : ((_1873_v88).equals(_1873_v88)));
          let _rhs320 = _dafny.Seq.IsPrefixOf(_1874_v89, _1874_v89);
          let _lhs232 = globalState;
          let _lhs233 = globalState;
          let _lhs234 = globalState;
          let _lhs235 = globalState;
          _lhs232.f8 = _rhs316;
          _lhs233.f8 = _rhs317;
          _lhs234.f8 = _rhs318;
          _lhs235.f8 = _rhs319;
          r3 = _rhs320;
          _1863_v78 = (_1863_v78).update(true, (p0).isLessThan(p0));
          (globalState).f9 = new BigNumber((_module.__default.fm39(p0, globalState)).length);
        }
        (globalState).f9 = p0;
        r1 = p0;
        _module.__default.m0(globalState);
      } else {
        let _index287 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length));
        (p1)[_index287] = new BigNumber(770);
        let _1875_v90;
        _1875_v90 = _dafny.Seq.of(new BigNumber(-29));
        let _index288 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length));
        (p1)[_index288] = ((_1758_v0) ? (p0) : (new BigNumber((_1875_v90).length)));
        let _1876_v91;
        _1876_v91 = _module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(643)), ((_1877_p1) => function (_1878_i6) {
  return (_1877_p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_1877_p1).length))];
})(p1)));
        let _1879_v92;
        _1879_v92 = _dafny.Seq.of(_1758_v0, false);
        let _1880_v93;
        _1880_v93 = _dafny.Set.fromElements(_1758_v0, _1758_v0, _1758_v0);
        let _1881_v94;
        let _nw305 = Array((new BigNumber(6)).toNumber());
        _nw305[(_dafny.ZERO).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length))];
        _nw305[(_dafny.ONE).toNumber()] = p0;
        _nw305[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length))]);
        _nw305[(new BigNumber(3)).toNumber()] = (new BigNumber((_1879_v92).length)).plus((p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length))]);
        _nw305[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw305[(new BigNumber(5)).toNumber()] = new BigNumber((_1880_v93).length);
        _1881_v94 = _nw305;
        let _rhs321 = _1876_v91;
        let _rhs322 = _1758_v0;
        let _rhs323 = _1881_v94;
        let _rhs324 = ((_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length))])).plus(_module.__default.safeDivisionInt(new BigNumber(418), (p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length))]));
        let _lhs236 = globalState;
        _1876_v91 = _rhs321;
        _lhs236.f8 = _rhs322;
        _1881_v94 = _rhs323;
        r1 = _rhs324;
        let _1882_v95;
        let _nw306 = new _module.C2();
        _nw306.__ctor((_this).f11);
        _1882_v95 = _nw306;
        let _1883_v96;
        _1883_v96 = _dafny.MultiSet.fromElements(_1758_v0);
        (globalState).f9 = _module.__default.safeDivisionInt((new BigNumber(20)).multipliedBy(p0), new BigNumber(((_1883_v96).Difference(_1883_v96)).cardinality()));
        let _1884_v97;
        _1884_v97 = new _dafny.CodePoint('o'.codePointAt(0));
        let _index289 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length));
        let _index290 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length));
        let _rhs325 = _1884_v97;
        let _rhs326 = (p0).plus(_module.__default.safeModuloInt((p1)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length))], (_dafny.ZERO).minus(p0)));
        let _rhs327 = (_dafny.ZERO).minus(new BigNumber(-409));
        let _lhs237 = p1;
        let _lhs238 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length));
        let _lhs239 = p1;
        let _lhs240 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((p1).length));
        r0 = _rhs325;
        _lhs237[_lhs238] = _rhs326;
        _lhs239[_lhs240] = _rhs327;
      }
      _module.__default.m0(globalState);
      let _1885_v98;
      _1885_v98 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1886_i7;
      _1886_i7 = _dafny.ZERO;
      L8: {
        while (_module.__default.fm1((p0).isLessThan(p0), _1885_v98, globalState)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1886_i7)) {
              break L8;
            }
            _1886_i7 = (_1886_i7).plus(_dafny.ONE);
            if (_1758_v0) {
              let _1887_v99;
              _1887_v99 = _module.D5.create_DC14(!_dafny.areEqual(_1832_v49, _1832_v49));
              let _pat_let_tv37 = _1758_v0;
              _1887_v99 = function (_pat_let29_0) {
                return function (_1888_dt__update__tmp_h0) {
                  return function (_pat_let30_0) {
                    return function (_1889_dt__update_hcf26_h0) {
                      return _module.D5.create_DC14(_1889_dt__update_hcf26_h0);
                    }(_pat_let30_0);
                  }(!(!(_pat_let_tv37)));
                }(_pat_let29_0);
              }(_1887_v99);
              let _1890_v100;
              let _nw307 = Array((new BigNumber(25)).toNumber());
              _nw307[(_dafny.ZERO).toNumber()] = _1885_v98;
              _nw307[(_dafny.ONE).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(2)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
              _nw307[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
              _nw307[(new BigNumber(5)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(6)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(7)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(8)).toNumber()] = (_1832_v49)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_1832_v49).length))];
              _nw307[(new BigNumber(9)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(10)).toNumber()] = (_1832_v49)[_module.__default.safeIndex(p0, new BigNumber((_1832_v49).length))];
              _nw307[(new BigNumber(11)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
              _nw307[(new BigNumber(13)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(14)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(15)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(16)).toNumber()] = (_1832_v49)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((_1832_v49).length))];
              _nw307[(new BigNumber(17)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(18)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(19)).toNumber()] = _module.__default.fm3(p0, false, globalState);
              _nw307[(new BigNumber(20)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
              _nw307[(new BigNumber(22)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(23)).toNumber()] = _1885_v98;
              _nw307[(new BigNumber(24)).toNumber()] = _1885_v98;
              _1890_v100 = _nw307;
              let _1891_v101;
              let _nw308 = Array((new BigNumber(2)).toNumber());
              _nw308[(_dafny.ZERO).toNumber()] = _1890_v100;
              _nw308[(_dafny.ONE).toNumber()] = _1890_v100;
              _1891_v101 = _nw308;
              let _1892_v102;
              _1892_v102 = _dafny.MultiSet.fromElements(_1885_v98, _1885_v98, _1885_v98);
              let _index291 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((p1).length));
              (p1)[_index291] = p0;
              let _1893_v103;
              _1893_v103 = _dafny.Seq.of(_1758_v0, _1758_v0, _1758_v0, _1758_v0);
              let _1894_v104;
              _1894_v104 = _dafny.Set.fromElements(p0, p0);
              let _index292 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((p1).length));
              let _rhs328 = _1891_v101;
              let _rhs329 = _1892_v102;
              let _rhs330 = _module.__default.fm0(((_1893_v103)[_module.__default.safeIndex(p0, new BigNumber((_1893_v103).length))]) || (_1758_v0), globalState);
              let _rhs331 = !(_1894_v104).contains(p0);
              let _lhs241 = p1;
              let _lhs242 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((p1).length));
              _1891_v101 = _rhs328;
              _1892_v102 = _rhs329;
              _lhs241[_lhs242] = _rhs330;
              r3 = _rhs331;
              let _1895_v105;
              _1895_v105 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_1758_v0, globalState),_1758_v0);
              let _1896_v106;
              _1896_v106 = _module.D4.create_DC11(new BigNumber((_1895_v105).length), (p1)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((p1).length))], _1758_v0);
              let _1897_v107;
              _1897_v107 = _module.D0.create_DC0((p1)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((p1).length))]);
              let _1898_v108;
              _1898_v108 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),(p1)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((p1).length))]);
              let _1899_v109;
              let _nw309 = new _module.C1();
              _nw309.__ctor((_1896_v106).dtor_cf21, _1758_v0, _1897_v107, !(new BigNumber((_1898_v108).length)).isEqualTo(p0));
              _1899_v109 = _nw309;
              (globalState).f9 = p0;
              (globalState).f8 = (_1899_v109).f19;
            } else {
              let _1900_v111;
              _1900_v111 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _1901_v112;
              let _nw310 = Array((new BigNumber(15)).toNumber());
              _nw310[(_dafny.ZERO).toNumber()] = _1832_v49;
              _nw310[(_dafny.ONE).toNumber()] = ((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-820)), ((_1902_v98) => function (_1903_i8) {
                return _1902_v98;
              })(_1885_v98))) : (_1832_v49));
              _nw310[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1832_v49, _1832_v49), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1832_v49, _1832_v49)).length)), _1885_v98);
              _nw310[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_1885_v98, _1885_v98, _1885_v98);
              _nw310[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1832_v49, _dafny.Seq.of(_1885_v98, _1885_v98));
              _nw310[(new BigNumber(5)).toNumber()] = ((_1758_v0) ? (_1832_v49) : (_1832_v49));
              _nw310[(new BigNumber(6)).toNumber()] = _1832_v49;
              _nw310[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1832_v49, _dafny.Seq.of(_1885_v98));
              _nw310[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_1885_v98);
              _nw310[(new BigNumber(9)).toNumber()] = _1832_v49;
              _nw310[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1885_v98, _1885_v98), _module.__default.fm2(p0, _1833_v50, function () {
                let _coll64 = new _dafny.Map();
                for (const _compr_64 of (_1833_v50).Elements) {
                  let _1904_v110 = _compr_64;
                  if ((_1833_v50).contains(_1904_v110)) {
                    _coll64.push([(_1904_v110).plus(p0),new BigNumber(885)]);
                  }
                }
                return _coll64;
              }(), globalState));
              _nw310[(new BigNumber(11)).toNumber()] = _module.__default.fm2(p0, _1833_v50, _1900_v111, globalState);
              _nw310[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), ((_1905_v98) => function (_1906_i9) {
                return _1905_v98;
              })(_1885_v98));
              _nw310[(new BigNumber(13)).toNumber()] = _1832_v49;
              _nw310[(new BigNumber(14)).toNumber()] = _1832_v49;
              _1901_v112 = _nw310;
              let _index293 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1901_v112).length));
              (_1901_v112)[_index293] = _1832_v49;
              let _index294 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1901_v112).length));
              (_1901_v112)[_index294] = (_module.D2.create_DC5(_1832_v49)).dtor_cf10;
              (globalState).f8 = _1758_v0;
              let _1907_v113;
              let _nw311 = Array((new BigNumber(26)).toNumber()).fill(_module.D2.Default());
              _1907_v113 = _nw311;
              _1907_v113 = _1907_v113;
              let _1908_v114;
              let _nw312 = new _module.C0();
              _nw312.__ctor(true);
              _1908_v114 = _nw312;
              (globalState).f9 = p0;
            }
            let _1909_v115;
            _1909_v115 = _dafny.Seq.of(_1758_v0);
            _1885_v98 = (((_1909_v115)[_module.__default.safeIndex(new BigNumber((_1833_v50).cardinality()), new BigNumber((_1909_v115).length))]) ? (_1885_v98) : (_1885_v98));
            let _index295 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length));
            (p1)[_index295] = p0;
            let _index296 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length));
            (p1)[_index296] = p0;
            let _1910_v116;
            let _nw313 = new _module.C3();
            _nw313.__ctor(_1758_v0, _module.__default.fm1(_1758_v0, _1885_v98, globalState), (_this).f11);
            _1910_v116 = _nw313;
            let _1911_v117;
            _1911_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1910_v116,(_1910_v116).f22);
            let _1912_v118;
            _1912_v118 = _dafny.Seq.of((p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))], new BigNumber(526));
            let _1913_v119;
            _1913_v119 = _dafny.MultiSet.fromElements(_1912_v118);
            _1911_v117 = (_1911_v117).update(_1910_v116, (_dafny.MultiSet.fromElements((_module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), function (_1914_i10) {
  return new BigNumber(421);
}))).dtor_cf25)).IsDisjointFrom(_1913_v119));
          }
        }
      }
      r0 = _1885_v98;
      r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((((false) ? (_1832_v49) : (_1832_v49))).length)), (p0).plus(p0));
      let _1915_v120;
      _1915_v120 = _module.D2.create_DC8(p0, _1885_v98);
      r2 = _1915_v120;
      r3 = _1758_v0;
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = false;
      let r2 = false;
      let _1916_v0;
      _1916_v0 = true;
      let _1917_i0;
      _1917_i0 = _dafny.ZERO;
      L9: {
        while (!(_1916_v0) || (!(false))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1917_i0)) {
              break L9;
            }
            _1917_i0 = (_1917_i0).plus(_dafny.ONE);
            let _1918_v1;
            _1918_v1 = new BigNumber(631);
            let _1919_v2;
            _1919_v2 = _dafny.MultiSet.fromElements(_1916_v0);
            let _1920_v3;
            _1920_v3 = _dafny.Seq.of(new BigNumber(-628));
            let _1921_v4;
            _1921_v4 = _dafny.Seq.of((((_1919_v2).contains(_1916_v0)) ? ((_1919_v2).get(_1916_v0)) : (_1918_v1)), _1918_v1, (_1920_v3)[_module.__default.safeIndex(_1918_v1, new BigNumber((_1920_v3).length))]);
            let _rhs332 = _1918_v1;
            let _rhs333 = new BigNumber((((_1916_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), ((_1922_v1) => function (_1923_i1) {
              return _1922_v1;
            })(_1918_v1))) : (_1921_v4))).length);
            let _lhs243 = globalState;
            let _lhs244 = globalState;
            _lhs243.f9 = _rhs332;
            _lhs244.f9 = _rhs333;
            _1916_v0 = true;
            let _1924_v5;
            _1924_v5 = new _dafny.CodePoint('c'.codePointAt(0));
            let _1925_v6;
            _1925_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1918_v1,new BigNumber(921));
            let _1926_v7;
            _1926_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1924_v5,_1925_v6);
            let _1927_v8;
            _1927_v8 = _module.D5.create_DC13(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1920_v3, _module.__default.safeIndex(_1918_v1, new BigNumber((_1920_v3).length)), _1918_v1), _1921_v4), _module.__default.safeIndex(new BigNumber((_1921_v4).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1920_v3, _module.__default.safeIndex(_1918_v1, new BigNumber((_1920_v3).length)), _1918_v1), _1921_v4)).length)), new BigNumber(((((_1926_v7).contains(_1924_v5)) ? ((_1926_v7).get(_1924_v5)) : (_1925_v6))).length)));
            _1927_v8 = _1927_v8;
            let _1928_v9;
            let _nw314 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
            _1928_v9 = _nw314;
            let _index297 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1928_v9).length));
            (_1928_v9)[_index297] = _1918_v1;
            let _index298 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1928_v9).length));
            (_1928_v9)[_index298] = (_dafny.ZERO).minus(_1918_v1);
          }
        }
      }
      _module.__default.m0(globalState);
      let _1929_v10;
      let _init59 = ((_1930_v0) => function (_1931_i3) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_1930_v0, _1930_v0), _dafny.Seq.of(_1930_v0, _1930_v0));
      })(_1916_v0);
      let _nw315 = Array((new BigNumber(29)).toNumber());
      for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw315.length); _i0_59++) {
        _nw315[_i0_59] = _init59(new BigNumber(_i0_59));
      }
      _1929_v10 = _nw315;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1929_v10).length))) {
        let _1932_i2 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1932_i2)) && ((_1932_i2).isLessThan(new BigNumber((_1929_v10).length))))) {
          (_1929_v10)[(_1932_i2)] = _dafny.Seq.Concat(_dafny.Seq.of(_1916_v0, !(_1916_v0), _1916_v0), _dafny.Seq.of(_1916_v0, _1916_v0));
        }
      }
      let _1933_v11;
      _1933_v11 = new BigNumber(985);
      (globalState).f9 = (_1933_v11).minus(_1933_v11);
      if ((_module.__default.fm0(_1916_v0, globalState)).isLessThan((new BigNumber(355)).multipliedBy(_1933_v11))) {
        let _1934_v12;
        let _init60 = ((_1935_v0) => function (_1936_i4) {
          return _1935_v0;
        })(_1916_v0);
        let _nw316 = Array((new BigNumber(16)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw316.length); _i0_60++) {
          _nw316[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _1934_v12 = _nw316;
        let _1937_v13;
        _1937_v13 = _module.D5.create_DC14(_1916_v0);
        let _index299 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length));
        (_1934_v12)[_index299] = (_1937_v13).dtor_cf26;
        let _index300 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length));
        (_1934_v12)[_index300] = ((_1916_v0) ? (_1916_v0) : (_1916_v0));
        let _1938_v14;
        _1938_v14 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_1933_v11, _1933_v11)).length));
        let _1939_v15;
        let _nw317 = new _module.C0();
        _nw317.__ctor((_1933_v11).isLessThanOrEqualTo((_1938_v14)[_module.__default.safeIndex(_1933_v11, new BigNumber((_1938_v14).length))]));
        _1939_v15 = _nw317;
        let _1940_v16;
        _1940_v16 = _dafny.Seq.UnicodeFromString("kba");
        let _1941_v17;
        _1941_v17 = _dafny.Set.fromElements((_1939_v15).f17, (_1939_v15).f17, (_1939_v15).f17);
        let _1942_v18;
        _1942_v18 = _dafny.Seq.of(_dafny.Set.fromElements((_1934_v12)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length))], true), (_1941_v17).Intersect(_1941_v17), _1941_v17);
        let _rhs334 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1940_v16, _1940_v16), _dafny.Seq.UnicodeFromString("fns"));
        let _rhs335 = (_dafny.ZERO).minus(_1933_v11);
        let _rhs336 = (_1942_v18)[_module.__default.safeIndex(_module.__default.fm0((_1939_v15).f17, globalState), new BigNumber((_1942_v18).length))];
        let _rhs337 = true;
        let _lhs245 = globalState;
        let _lhs246 = globalState;
        _1940_v16 = _rhs334;
        _lhs245.f9 = _rhs335;
        _1941_v17 = _rhs336;
        _lhs246.f8 = _rhs337;
        let _1943_v19;
        _1943_v19 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1944_v20;
        _1944_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_module.__default.fm1((_1939_v15).f17, _1943_v19, globalState), globalState),_module.__default.safeDivisionInt(_1933_v11, _1933_v11));
        _1944_v20 = (_1944_v20).update(_module.__default.fm0((_1934_v12)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length))], globalState), new BigNumber(518));
        let _1945_v21;
        _1945_v21 = _module.D7.create_DC17(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("y"), _module.__default.safeIndex(_1933_v11, new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)), _1943_v19), _dafny.Seq.update(_1940_v16, _module.__default.safeIndex(_1933_v11, new BigNumber((_1940_v16).length)), _module.__default.fm3(_1933_v11, true, globalState))), _1934_v12, _1916_v0, _1933_v11);
        let _source23 = _1945_v21;
        if (_source23.is_DC17) {
          let _1946___mcc_h0 = (_source23).cf29;
          let _1947___mcc_h1 = (_source23).cf30;
          let _1948___mcc_h2 = (_source23).cf31;
          let _1949___mcc_h3 = (_source23).cf32;
          let _1950_cf32 = _1949___mcc_h3;
          let _1951_cf31 = _1948___mcc_h2;
          let _1952_cf30 = _1947___mcc_h1;
          let _1953_cf29 = _1946___mcc_h0;
          let _1954_v22;
          let _init61 = ((_1955_cf29) => function (_1956_i5) {
            return _1955_cf29;
          })(_1953_cf29);
          let _nw318 = Array((new BigNumber(27)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw318.length); _i0_61++) {
            _nw318[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _1954_v22 = _nw318;
          let _index301 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1954_v22).length));
          (_1954_v22)[_index301] = _1940_v16;
          let _index302 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1954_v22).length));
          let _rhs338 = new _dafny.CodePoint('a'.codePointAt(0));
          let _rhs339 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-245)), ((_1957_v19) => function (_1958_i6) {
            return _1957_v19;
          })(_1943_v19)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-320)), ((_1959_v19) => function (_1960_i7) {
            return _1959_v19;
          })(_1943_v19)));
          let _lhs247 = _1954_v22;
          let _lhs248 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1954_v22).length));
          r0 = _rhs338;
          _lhs247[_lhs248] = _rhs339;
          (globalState).f9 = _1950_cf32;
          let _index303 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1954_v22).length));
          (_1954_v22)[_index303] = _1940_v16;
          (globalState).f9 = (((_1938_v14)[_module.__default.safeIndex(_1950_cf32, new BigNumber((_1938_v14).length))]).multipliedBy((_dafny.ZERO).minus(_1950_cf32))).multipliedBy(_1933_v11);
        } else {
          let _1961___mcc_h4 = (_source23).cf28;
          let _1962_cf28 = _1961___mcc_h4;
          let _1963_v23;
          _1963_v23 = _module.D0.create_DC0(_1933_v11);
          let _1964_v24;
          let _nw319 = new _module.C1();
          _nw319.__ctor((_1939_v15).f17, (_1939_v15).f17, _1963_v23, (_1934_v12)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length))]);
          _1964_v24 = _nw319;
          let _1965_v25;
          let _nw320 = new _module.C1();
          _nw320.__ctor((new BigNumber((_1944_v20).length)).isLessThanOrEqualTo(_1933_v11), (_1939_v15).f17, _module.D0.create_DC0(new BigNumber((_dafny.Seq.of(_1964_v24)).length)), (_1964_v24).f20);
          _1965_v25 = _nw320;
          let _1966_v26;
          _1966_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1933_v11,(_1965_v25).f20);
          let _1967_v27;
          _1967_v27 = _dafny.MultiSet.fromElements((_1964_v24).f20);
          let _1968_v28;
          _1968_v28 = _module.D0.create_DC2(_1967_v27, (_1965_v25).f19, _1933_v11, (_1964_v24).f19);
          let _1969_v29;
          _1969_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("pmk"), _1940_v16);
          let _rhs340 = (_1966_v26).update(new BigNumber(277), (_1939_v15).f17);
          let _rhs341 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1933_v11), (_dafny.ZERO).minus(_1933_v11));
          let _rhs342 = (_1968_v28).dtor_cf7;
          let _rhs343 = (_1969_v29).IsDisjointFrom(_1969_v29);
          let _lhs249 = globalState;
          _1966_v26 = _rhs340;
          _1933_v11 = _rhs341;
          r1 = _rhs342;
          _lhs249.f8 = _rhs343;
          let _1970_v30;
          _1970_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(59),new _dafny.CodePoint('s'.codePointAt(0)));
          let _1971_v31;
          _1971_v31 = _dafny.Set.fromElements(_1933_v11, new BigNumber((_1940_v16).length));
          let _1972_v33;
          let _nw321 = Array((new BigNumber(24)).toNumber());
          _nw321[(_dafny.ZERO).toNumber()] = new BigNumber(((((_1962_cf28)[_module.__default.safeIndex(new BigNumber((_1970_v30).length), new BigNumber((_1962_cf28).length))]) ? (_dafny.Seq.UnicodeFromString("gsddkqjn")) : (_dafny.Seq.UnicodeFromString("qbyry")))).length);
          _nw321[(_dafny.ONE).toNumber()] = new BigNumber((_1971_v31).length);
          _nw321[(new BigNumber(2)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(3)).toNumber()] = (_1964_v24).fm9(_module.__default.fm57(globalState), _dafny.Set.fromElements(_1933_v11), globalState);
          _nw321[(new BigNumber(4)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(5)).toNumber()] = (_1965_v25).fm30(_dafny.Map.Empty.slice().updateUnsafe(_1933_v11,(_1934_v12)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length))]), globalState);
          _nw321[(new BigNumber(6)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_1933_v11, new BigNumber((_1938_v14).length))).length);
          _nw321[(new BigNumber(8)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(9)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(10)).toNumber()] = (((_1939_v15).f17) ? (_1933_v11) : (_1933_v11));
          _nw321[(new BigNumber(11)).toNumber()] = (_1938_v14)[_module.__default.safeIndex(_1933_v11, new BigNumber((_1938_v14).length))];
          _nw321[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.update(_1940_v16, _module.__default.safeIndex(_1933_v11, new BigNumber((_1940_v16).length)), _1943_v19)).length);
          _nw321[(new BigNumber(13)).toNumber()] = ((_1965_v25).fm29(_1933_v11, globalState)).multipliedBy(new BigNumber((function () {
            let _coll65 = new _dafny.Set();
            for (const _compr_65 of _dafny.IntegerRange(new BigNumber(372), new BigNumber(-617))) {
              let _1973_v32 = _compr_65;
              if (((new BigNumber(372)).isLessThanOrEqualTo(_1973_v32)) && ((_1973_v32).isLessThan(new BigNumber(-617)))) {
                _coll65.add(_module.__default.safeModuloInt(_1973_v32, (_dafny.ZERO).minus(_1933_v11)));
              }
            }
            return _coll65;
          }()).length));
          _nw321[(new BigNumber(14)).toNumber()] = new BigNumber(332);
          _nw321[(new BigNumber(15)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_1933_v11);
          _nw321[(new BigNumber(17)).toNumber()] = (new BigNumber((_dafny.Seq.of((_1964_v24).f20, false)).length)).multipliedBy(_1933_v11);
          _nw321[(new BigNumber(18)).toNumber()] = _1933_v11;
          _nw321[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(_1933_v11, _1933_v11);
          _nw321[(new BigNumber(20)).toNumber()] = (_1933_v11).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1965_v25).f19,_dafny.Seq.UnicodeFromString("jrtnqutc"))).length));
          _nw321[(new BigNumber(21)).toNumber()] = (_1933_v11).plus(_1933_v11);
          _nw321[(new BigNumber(22)).toNumber()] = (_1938_v14)[_module.__default.safeIndex(_1933_v11, new BigNumber((_1938_v14).length))];
          _nw321[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(_1933_v11, _1933_v11);
          _1972_v33 = _nw321;
          let _init62 = ((_1974_v11) => function (_1975_i8) {
            return _module.__default.safeModuloInt(_1975_i8, _1974_v11);
          })(_1933_v11);
          let _nw322 = Array((new BigNumber(26)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw322.length); _i0_62++) {
            _nw322[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _1972_v33 = _nw322;
          let _1976_v34;
          let _nw323 = new _module.C0();
          _nw323.__ctor((_1934_v12)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1934_v12).length))]);
          _1976_v34 = _nw323;
        }
      } else {
        let _1977_v35;
        _1977_v35 = _dafny.Seq.of((_this).f11);
        let _1978_v36;
        let _nw324 = new _module.C4();
        _nw324.__ctor((_1977_v35)[_module.__default.safeIndex(_1933_v11, new BigNumber((_1977_v35).length))]);
        _1978_v36 = _nw324;
        _module.__default.m0(globalState);
        (globalState).f8 = !(_1916_v0);
        (globalState).f9 = new BigNumber(385);
        (globalState).f9 = (_dafny.ZERO).minus((_module.__default.fm0(false, globalState)).plus(_1933_v11));
      }
      let _1979_v37;
      _1979_v37 = _dafny.MultiSet.fromElements(_1933_v11);
      let _1980_v38;
      _1980_v38 = _dafny.Seq.UnicodeFromString("vmba");
      let _1981_v39;
      _1981_v39 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1982_v40;
      _1982_v40 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_1980_v38, _module.__default.safeIndex(_1933_v11, new BigNumber((_1980_v38).length)), _1981_v39)).length));
      _1979_v37 = (_1979_v37).update((new BigNumber((_1982_v40).length)).multipliedBy(_1933_v11), _module.__default.abs((_1933_v11).multipliedBy((_dafny.ZERO).minus(_1933_v11))));
      r0 = _module.__default.fm3(new BigNumber(((_dafny.MultiSet.fromElements(_1933_v11)).Difference(_dafny.MultiSet.fromElements(_1933_v11, new BigNumber(-496)))).cardinality()), (_1916_v0) || (!(_1916_v0)), globalState);
      let _1983_v41;
      _1983_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1933_v11,_1916_v0);
      r1 = !((((_1983_v41).contains(_module.__default.fm0(!(!(_1916_v0)), globalState))) ? ((_1983_v41).get(_module.__default.fm0(!(!(_1916_v0)), globalState))) : (_1916_v0))) || (!_dafny.areEqual(_dafny.Seq.UnicodeFromString("f"), _dafny.Seq.UnicodeFromString("tnwss")));
      r2 = _1916_v0;
      return [r0, r1, r2];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this.f16 = _dafny.Seq.UnicodeFromString("");
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f15, f16) {
      let _this = this;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      return;
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let _1984_v0;
      _1984_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f15);
      let _hi14 = (((_1984_v0).contains(false)) ? ((_1984_v0).get(false)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(647)), function (_1985_i1) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length)));
      for (let _1986_i0 = (_this).f15; _1986_i0.isLessThan(_hi14); _1986_i0 = _1986_i0.plus(_dafny.ONE)) {
        let _1987_v1;
        _1987_v1 = _dafny.Seq.of(p0, p0);
        let _1988_v2;
        let _nw325 = Array((new BigNumber(3)).toNumber());
        _nw325[(_dafny.ZERO).toNumber()] = (_1987_v1)[_module.__default.safeIndex(_1986_i0, new BigNumber((_1987_v1).length))];
        _nw325[(_dafny.ONE).toNumber()] = p0;
        _nw325[(new BigNumber(2)).toNumber()] = p0;
        _1988_v2 = _nw325;
        let _1989_v3;
        _1989_v3 = _dafny.Set.fromElements(_1988_v2);
        let _rhs344 = (_1989_v3).Difference(((p0) ? (_1989_v3) : (_1989_v3)));
        let _rhs345 = (_this).f15;
        let _rhs346 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1986_i0, new BigNumber((_this.f16).length)));
        let _lhs250 = globalState;
        let _lhs251 = globalState;
        _1989_v3 = _rhs344;
        _lhs250.f9 = _rhs345;
        _lhs251.f9 = _rhs346;
        let _1990_v4;
        _1990_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
        _1990_v4 = (_1990_v4).update(p0, _this.f16);
        let _1991_v5;
        _1991_v5 = _dafny.MultiSet.fromElements(p0);
        let _1992_v6;
        _1992_v6 = _module.D0.create_DC2(_1991_v5, p0, (_this).f15, p0);
        let _1993_v7;
        let _nw326 = new _module.C0();
        _nw326.__ctor((_1992_v6).dtor_cf7);
        _1993_v7 = _nw326;
        let _1994_v8;
        _1994_v8 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1995_v9;
        let _nw327 = Array((new BigNumber(19)).toNumber());
        _nw327[(_dafny.ZERO).toNumber()] = _1994_v8;
        _nw327[(_dafny.ONE).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(2)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(3)).toNumber()] = _module.__default.fm3(_1986_i0, p0, globalState);
        _nw327[(new BigNumber(4)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(5)).toNumber()] = (_module.D1.create_DC4((_this.f16)[_module.__default.safeIndex(new BigNumber((_1987_v1).length), new BigNumber((_this.f16).length))])).dtor_cf9;
        _nw327[(new BigNumber(6)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(7)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(8)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(9)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(10)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(11)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(12)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(13)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(14)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(15)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(16)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(17)).toNumber()] = _1994_v8;
        _nw327[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _1995_v9 = _nw327;
        let _1996_v10;
        _1996_v10 = _dafny.MultiSet.fromElements(_1995_v9, _1995_v9);
        let _1997_v11;
        let _nw328 = Array((_dafny.ONE).toNumber());
        _nw328[(_dafny.ZERO).toNumber()] = _1996_v10;
        _1997_v11 = _nw328;
        let _index304 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1997_v11).length));
        (_1997_v11)[_index304] = _dafny.MultiSet.fromElements(_1995_v9, _1995_v9);
        let _1998_v12;
        _1998_v12 = _dafny.Seq.of(_1996_v10);
        let _index305 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1997_v11).length));
        (_1997_v11)[_index305] = ((_1998_v12)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_1998_v12).length))]).Intersect(_1996_v10);
      }
      let _1999_v13;
      _1999_v13 = _dafny.Set.fromElements(_module.__default.fm0(p0, globalState));
      let _2000_v14;
      _2000_v14 = _dafny.MultiSet.fromElements(!(p0));
      let _2001_v15;
      let _nw329 = Array((new BigNumber(23)).toNumber());
      _nw329[(_dafny.ZERO).toNumber()] = true;
      _nw329[(_dafny.ONE).toNumber()] = ((p0) ? (p0) : (!(p0)));
      _nw329[(new BigNumber(2)).toNumber()] = p0;
      _nw329[(new BigNumber(3)).toNumber()] = p0;
      _nw329[(new BigNumber(4)).toNumber()] = p0;
      _nw329[(new BigNumber(5)).toNumber()] = ((_this).f15).isEqualTo(new BigNumber((_1984_v0).length));
      _nw329[(new BigNumber(6)).toNumber()] = (_module.__default.fm0(true, globalState)).isLessThan((_this).f15);
      _nw329[(new BigNumber(7)).toNumber()] = p0;
      _nw329[(new BigNumber(8)).toNumber()] = p0;
      _nw329[(new BigNumber(9)).toNumber()] = p0;
      _nw329[(new BigNumber(10)).toNumber()] = p0;
      _nw329[(new BigNumber(11)).toNumber()] = (_1999_v13).IsProperSubsetOf(_1999_v13);
      _nw329[(new BigNumber(12)).toNumber()] = p0;
      _nw329[(new BigNumber(13)).toNumber()] = p0;
      _nw329[(new BigNumber(14)).toNumber()] = p0;
      _nw329[(new BigNumber(15)).toNumber()] = p0;
      _nw329[(new BigNumber(16)).toNumber()] = p0;
      _nw329[(new BigNumber(17)).toNumber()] = p0;
      _nw329[(new BigNumber(18)).toNumber()] = false;
      _nw329[(new BigNumber(19)).toNumber()] = !(p0);
      _nw329[(new BigNumber(20)).toNumber()] = p0;
      _nw329[(new BigNumber(21)).toNumber()] = p0;
      _nw329[(new BigNumber(22)).toNumber()] = ((_2000_v14).update(!(p0), _module.__default.abs(new BigNumber((_this.f16).length)))).IsSubsetOf(_2000_v14);
      _2001_v15 = _nw329;
      let _index306 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      (_2001_v15)[_index306] = p0;
      let _2002_v16;
      _2002_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _index307 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      (_2001_v15)[_index307] = (((_2002_v16).contains(p0)) ? ((_2002_v16).get(p0)) : ((_dafny.Set.fromElements((_this).f15, (_this).f15)).IsProperSubsetOf(_1999_v13)));
      let _2003_v17;
      let _nw330 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2003_v17 = _nw330;
      let _2004_v18;
      _2004_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2003_v17,new BigNumber((_this.f16).length));
      _2004_v18 = (_2004_v18).update(_2003_v17, _module.__default.safeDivisionInt(new BigNumber(-627), new BigNumber(-656)));
      let _2005_v19;
      let _nw331 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _2005_v19 = _nw331;
      let _index308 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
      (_2005_v19)[_index308] = ((_this).f15).minus((_this).f15);
      let _2006_v20;
      _2006_v20 = _dafny.Seq.of((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]);
      let _2007_v21;
      _2007_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_2006_v20)[_module.__default.safeIndex((_this).f15, new BigNumber((_2006_v20).length))]);
      let _index309 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      let _index310 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      let _index311 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      let _index312 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
      let _rhs347 = (((_2007_v21).contains(_module.__default.fm0((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], globalState))) ? ((_2007_v21).get(_module.__default.fm0((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], globalState))) : ((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]));
      let _rhs348 = (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))];
      let _rhs349 = (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))];
      let _rhs350 = (_this).f15;
      let _lhs252 = _2001_v15;
      let _lhs253 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      let _lhs254 = _2001_v15;
      let _lhs255 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      let _lhs256 = _2001_v15;
      let _lhs257 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
      let _lhs258 = _2005_v19;
      let _lhs259 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
      _lhs252[_lhs253] = _rhs347;
      _lhs254[_lhs255] = _rhs348;
      _lhs256[_lhs257] = _rhs349;
      _lhs258[_lhs259] = _rhs350;
      let _2008_v22;
      _2008_v22 = _dafny.Set.fromElements((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]);
      _2008_v22 = _2008_v22;
      if (!((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) {
        let _2009_v23;
        _2009_v23 = _dafny.Map.Empty.slice().updateUnsafe((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))],_1999_v13);
        let _2010_v24;
        _2010_v24 = _dafny.Map.Empty.slice().updateUnsafe((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))],_1999_v13);
        let _2011_v25;
        _2011_v25 = _dafny.MultiSet.fromElements((((_2009_v23).contains(new BigNumber((_this.f16).length))) ? ((_2009_v23).get(new BigNumber((_this.f16).length))) : ((((_2010_v24).contains((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) ? ((_2010_v24).get((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) : (_1999_v13)))));
        let _2012_v26;
        _2012_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2011_v25);
        let _index313 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
        (_2001_v15)[_index313] = !((((_2012_v26).contains((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) ? ((_2012_v26).get((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) : (_2011_v25))).equals(((!(p0)) ? (_2011_v25) : (_module.__default.fm14(_2006_v20, (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], globalState))));
        let _2013_v27;
        _2013_v27 = _dafny.Seq.of(_2002_v16);
        let _2014_v28;
        _2014_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2005_v19,_2013_v27);
        let _2015_v29;
        let _nw332 = Array((new BigNumber(13)).toNumber());
        _nw332[(_dafny.ZERO).toNumber()] = (((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]) ? (_2013_v27) : (_2013_v27));
        _nw332[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2013_v27, _2013_v27);
        _nw332[(new BigNumber(2)).toNumber()] = _2013_v27;
        _nw332[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_2002_v16);
        _nw332[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-892)), ((_2016_v16) => function (_2017_i2) {
          return _2016_v16;
        })(_2002_v16));
        _nw332[(new BigNumber(5)).toNumber()] = _2013_v27;
        _nw332[(new BigNumber(6)).toNumber()] = (((_2014_v28).contains(_2005_v19)) ? ((_2014_v28).get(_2005_v19)) : (_2013_v27));
        _nw332[(new BigNumber(7)).toNumber()] = _2013_v27;
        _nw332[(new BigNumber(8)).toNumber()] = _module.__default.fm15((_2006_v20)[_module.__default.safeIndex(new BigNumber((_this.f16).length), new BigNumber((_2006_v20).length))], globalState);
        _nw332[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2013_v27, _2013_v27);
        _nw332[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(_2002_v16, _2002_v16);
        _nw332[(new BigNumber(11)).toNumber()] = _2013_v27;
        _nw332[(new BigNumber(12)).toNumber()] = ((true) ? (_2013_v27) : (_2013_v27));
        _2015_v29 = _nw332;
        let _index314 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2015_v29).length));
        (_2015_v29)[_index314] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-742)), ((_2018_v15, _2019_p0) => function (_2020_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe((_2018_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2018_v15).length))],_2019_p0);
        })(_2001_v15, p0));
        let _index315 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2015_v29).length));
        (_2015_v29)[_index315] = _dafny.Seq.Concat(_2013_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(960)), ((_2021_v16) => function (_2022_i4) {
          return _2021_v16;
        })(_2002_v16)));
        if (p0) {
          let _2023_v30;
          _2023_v30 = _module.D2.create_DC7(p0, (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), function (_2024_i5) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}));
          let _2025_v31;
          _2025_v31 = new _dafny.CodePoint('u'.codePointAt(0));
          let _index316 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          let _rhs351 = ((!((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) ? ((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]) : (new BigNumber((_this.f16).length)));
          let _rhs352 = new BigNumber((_this.f16).length);
          let _rhs353 = ((!(true) || (p0)) ? ((_module.__default.fm1((_2023_v30).dtor_cf13, _2025_v31, globalState)) || (false)) : (p0));
          let _lhs260 = globalState;
          let _lhs261 = globalState;
          let _lhs262 = _2001_v15;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          _lhs260.f9 = _rhs351;
          _lhs261.f9 = _rhs352;
          _lhs262[_lhs263] = _rhs353;
          let _2026_v33;
          _2026_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f16).length),_module.__default.fm0((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], globalState));
          let _2027_v34;
          let _nw333 = Array((new BigNumber(8)).toNumber());
          _nw333[(_dafny.ZERO).toNumber()] = _2007_v21;
          _nw333[(_dafny.ONE).toNumber()] = _2007_v21;
          _nw333[(new BigNumber(2)).toNumber()] = _2007_v21;
          _nw333[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,p0);
          _nw333[(new BigNumber(4)).toNumber()] = (_2007_v21).update(_module.__default.fm0((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], globalState), (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]);
          _nw333[(new BigNumber(5)).toNumber()] = _2007_v21;
          _nw333[(new BigNumber(6)).toNumber()] = function () {
            let _coll66 = new _dafny.Map();
            for (const _compr_66 of (_2026_v33).Keys.Elements) {
              let _2028_v32 = _compr_66;
              if ((_2026_v33).contains(_2028_v32)) {
                _coll66.push([_module.__default.safeDivisionInt(_2028_v32, (_this).f15),p0]);
              }
            }
            return _coll66;
          }();
          _nw333[(new BigNumber(7)).toNumber()] = _2007_v21;
          _2027_v34 = _nw333;
          let _2029_v35;
          let _nw334 = new _module.C2();
          _nw334.__ctor(_2027_v34);
          _2029_v35 = _nw334;
          _2029_v35 = _2029_v35;
          let _index317 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          (_2001_v15)[_index317] = (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))];
          let _index318 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          (_2005_v19)[_index318] = (_dafny.ZERO).minus((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _2030_v36;
          _2030_v36 = _dafny.MultiSet.fromElements(_2001_v15, _2001_v15);
          _2008_v22 = _dafny.Set.fromElements((((_2007_v21).contains((_this).f15)) ? ((_2007_v21).get((_this).f15)) : (!(p0))), p0, (_2030_v36).IsProperSubsetOf(_2030_v36));
        } else {
          let _2031_v37;
          _2031_v37 = new _dafny.CodePoint('n'.codePointAt(0));
          (globalState).f8 = _module.__default.fm1((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], (((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]) ? (_2031_v37) : (_2031_v37)), globalState);
          let _2032_v38;
          _2032_v38 = _dafny.MultiSet.fromElements((_this).f15, (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _2033_v39;
          _2033_v39 = _module.D2.create_DC7(p0, new BigNumber((_2032_v38).cardinality()), false, _this.f16);
          let _2034_v40;
          _2034_v40 = _dafny.MultiSet.fromElements(_2031_v37);
          let _2035_v41;
          _2035_v41 = _dafny.Map.Empty.slice().updateUnsafe((((_2033_v39).dtor_cf13) ? (_2034_v40) : (_2034_v40)),new BigNumber((_1999_v13).length));
          _2035_v41 = (_2035_v41).update((_2034_v40).update(_2031_v37, _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), ((_2036_v37) => function (_2037_i6) {
            return _2036_v37;
          })(_2031_v37))).length))), (_dafny.ZERO).minus((_this).f15));
          let _index319 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          (_2005_v19)[_index319] = (_this).f15;
          let _2038_v42;
          _2038_v42 = _module.D0.create_DC2((_2000_v14).update((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], _module.__default.abs((_this).f15)), (((_2007_v21).contains((_this).f15)) ? ((_2007_v21).get((_this).f15)) : ((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])), (_this).f15, (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]);
          let _2039_v43;
          _2039_v43 = _dafny.Seq.of(_module.__default.fm33(_2002_v16, globalState), (_2038_v42).dtor_cf4, _module.__default.fm33(_2002_v16, globalState));
          _2000_v14 = (_2039_v43)[_module.__default.safeIndex((_this).f15, new BigNumber((_2039_v43).length))];
          (_this).f16 = _this.f16;
        }
        if ((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]) {
          (globalState).f9 = new BigNumber(((((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]) ? (_this.f16) : (_this.f16))).length);
          let _2040_v44;
          _2040_v44 = _dafny.Seq.of(_dafny.Seq.of(!((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]), (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]));
          let _2041_v45;
          _2041_v45 = _dafny.Seq.of(_2040_v44);
          let _2042_v46;
          _2042_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
          let _2043_v47;
          let _nw335 = Array((new BigNumber(16)).toNumber());
          _nw335[(_dafny.ZERO).toNumber()] = _2007_v21;
          _nw335[(_dafny.ONE).toNumber()] = (_2007_v21).update((_this).f15, (_2006_v20)[_module.__default.safeIndex((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], new BigNumber((_2006_v20).length))]);
          _nw335[(new BigNumber(2)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(3)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(4)).toNumber()] = _module.__default.fm27((_2041_v45)[_module.__default.safeIndex((_this).f15, new BigNumber((_2041_v45).length))], globalState);
          _nw335[(new BigNumber(5)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(6)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(7)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(true, globalState),(_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]);
          _nw335[(new BigNumber(9)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(10)).toNumber()] = (_2007_v21).update(new BigNumber(((((_2042_v46).contains((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) ? ((_2042_v46).get((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) : (_this.f16))).length), !((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]));
          _nw335[(new BigNumber(11)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,false);
          _nw335[(new BigNumber(13)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(14)).toNumber()] = _2007_v21;
          _nw335[(new BigNumber(15)).toNumber()] = _module.__default.fm27(_2040_v44, globalState);
          _2043_v47 = _nw335;
          let _2044_v48;
          let _nw336 = new _module.C4();
          _nw336.__ctor(_2043_v47);
          _2044_v48 = _nw336;
          let _index320 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2003_v17).length));
          (_2003_v17)[_index320] = _this.f16;
          let _2045_v49;
          _2045_v49 = new _dafny.CodePoint('c'.codePointAt(0));
          let _2046_v50;
          _2046_v50 = _dafny.Set.fromElements(_this.f16);
          let _2047_v51;
          _2047_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2046_v50);
          let _index321 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2003_v17).length));
          let _rhs354 = new BigNumber(((_dafny.Set.fromElements(_this.f16, _dafny.Seq.update(_this.f16, _module.__default.safeIndex((_dafny.ZERO).minus((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]), new BigNumber((_this.f16).length)), _2045_v49), _dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), ((_2048_v49) => function (_2049_i7) {
            return _2048_v49;
          })(_2045_v49)))).Intersect((((_2047_v51).contains((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))])) ? ((_2047_v51).get((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))])) : (_dafny.Set.fromElements(_this.f16, _this.f16))))).length);
          let _rhs355 = _2044_v48;
          let _rhs356 = _this.f16;
          let _lhs264 = globalState;
          let _lhs265 = _2003_v17;
          let _lhs266 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2003_v17).length));
          _lhs264.f9 = _rhs354;
          _2044_v48 = _rhs355;
          _lhs265[_lhs266] = _rhs356;
          let _rhs357 = _module.__default.safeModuloInt(new BigNumber(785), (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _rhs358 = _module.__default.fm33(_2002_v16, globalState);
          let _lhs267 = globalState;
          _lhs267.f9 = _rhs357;
          _2000_v14 = _rhs358;
          let _2050_v52;
          _2050_v52 = _dafny.Seq.of((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _2051_v53;
          _2051_v53 = _module.D5.create_DC13(_2050_v52);
          let _2052_v54;
          _2052_v54 = _dafny.Map.Empty.slice().updateUnsafe((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))],_module.__default.fm0((_2044_v48).fm42((_this).f15, new BigNumber(((_2003_v17)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_2003_v17).length))]).length), globalState), globalState));
          let _index322 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          let _index323 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _index324 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _rhs359 = _2051_v53;
          let _rhs360 = !((new BigNumber((_2007_v21).length)).plus((_this).f15)).isEqualTo(new BigNumber((_2006_v20).length));
          let _rhs361 = (_this).f15;
          let _rhs362 = new BigNumber((((_2052_v54).Merge(_2052_v54)).Merge(_2052_v54)).length);
          let _lhs268 = _2001_v15;
          let _lhs269 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          let _lhs270 = _2005_v19;
          let _lhs271 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _lhs272 = _2005_v19;
          let _lhs273 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          _2051_v53 = _rhs359;
          _lhs268[_lhs269] = _rhs360;
          _lhs270[_lhs271] = _rhs361;
          _lhs272[_lhs273] = _rhs362;
          _2008_v22 = _2008_v22;
        } else {
          let _index325 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          (_2001_v15)[_index325] = (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))];
          let _index326 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          (_2005_v19)[_index326] = ((_this).f15).minus(new BigNumber((_dafny.Set.fromElements(_1999_v13)).length));
          let _index327 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          (_2001_v15)[_index327] = (_2006_v20)[_module.__default.safeIndex(_module.__default.fm0((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], globalState), new BigNumber((_2006_v20).length))];
          let _2053_v55;
          let _init63 = ((_2054_v14, _2055_v20) => function (_2056_i8) {
            return (_2054_v14).Union(_dafny.MultiSet.FromArray(_2055_v20));
          })(_2000_v14, _2006_v20);
          let _nw337 = Array((new BigNumber(23)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw337.length); _i0_63++) {
            _nw337[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2053_v55 = _nw337;
          let _2057_v56;
          _2057_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.fromElements((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]));
          let _index328 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_2053_v55).length));
          (_2053_v55)[_index328] = (((_2057_v56).contains(!(p0))) ? ((_2057_v56).get(!(p0))) : (_dafny.MultiSet.fromElements(p0)));
          let _index329 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_2053_v55).length));
          let _index330 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _index331 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _rhs363 = ((_2000_v14).Difference(_2000_v14)).Union((_dafny.MultiSet.fromElements(p0)).Intersect(_2000_v14));
          let _rhs364 = (_this).f15;
          let _rhs365 = (_this).f15;
          let _rhs366 = new BigNumber((_1999_v13).length);
          let _lhs274 = _2053_v55;
          let _lhs275 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_2053_v55).length));
          let _lhs276 = _2005_v19;
          let _lhs277 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _lhs278 = globalState;
          let _lhs279 = _2005_v19;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          _lhs274[_lhs275] = _rhs363;
          _lhs276[_lhs277] = _rhs364;
          _lhs278.f9 = _rhs365;
          _lhs279[_lhs280] = _rhs366;
          let _2058_v57;
          _2058_v57 = _dafny.MultiSet.fromElements((_this).f15);
          let _2059_v58;
          _2059_v58 = _module.D0.create_DC0((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _2060_v59;
          let _nw338 = new _module.C1();
          _nw338.__ctor((_2058_v57).IsDisjointFrom(_dafny.MultiSet.fromElements((((_2000_v14).contains((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) ? ((_2000_v14).get((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))])) : ((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))])))), (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))], _2059_v58, (_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]);
          _2060_v59 = _nw338;
        }
        let _2061_v60;
        let _init64 = ((_2062_v21) => function (_2063_i9) {
          return _2062_v21;
        })(_2007_v21);
        let _nw339 = Array((new BigNumber(16)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw339.length); _i0_64++) {
          _nw339[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _2061_v60 = _nw339;
        let _2064_v61;
        let _nw340 = new _module.C4();
        _nw340.__ctor(_2061_v60);
        _2064_v61 = _nw340;
        let _2065_v62;
        _2065_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
        let _2066_v63;
        _2066_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2064_v61,new BigNumber((_2065_v62).length));
        (globalState).f9 = new BigNumber((((!(p0)) ? (_2066_v63) : (_2066_v63))).length);
      } else {
        let _2067_v64;
        let _nw341 = new _module.C0();
        _nw341.__ctor(((_dafny.ZERO).minus(new BigNumber((_this.f16).length))).isLessThanOrEqualTo((_this).f15));
        _2067_v64 = _nw341;
        let _2068_v65;
        _2068_v65 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(289)), function (_2069_i10) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length));
        let _2070_v66;
        let _nw342 = new _module.C1();
        _nw342.__ctor(!(p0) || (false), p0, _2068_v65, ((_2006_v20)[_module.__default.safeIndex((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], new BigNumber((_2006_v20).length))]) === (!((_2067_v64).f17)));
        _2070_v66 = _nw342;
        if ((_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]) {
          (globalState).f8 = (_2070_v66).f19;
          let _2071_v67;
          _2071_v67 = new _dafny.CodePoint('b'.codePointAt(0));
          let _2072_v68;
          _2072_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2071_v67,(_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _2073_v69;
          _2073_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_this.f16).length), (_this).f15, new BigNumber((_module.__default.fm36((_2070_v66).f19, globalState)).length), (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]), _module.__default.safeIndex((((_2072_v68).contains(_2071_v67)) ? ((_2072_v68).get(_2071_v67)) : ((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))])), new BigNumber((_dafny.Seq.of(new BigNumber((_this.f16).length), (_this).f15, new BigNumber((_module.__default.fm36((_2070_v66).f19, globalState)).length), (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))])).length)), (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]),!(_2000_v14).contains((_2070_v66).f20));
          let _2074_v70;
          _2074_v70 = _dafny.Seq.of((_this).f15);
          let _rhs367 = _2001_v15;
          let _rhs368 = false;
          let _rhs369 = (_dafny.Map.Empty.slice().updateUnsafe(_2074_v70,true)).Merge(_2073_v69);
          let _lhs281 = globalState;
          _2001_v15 = _rhs367;
          _lhs281.f8 = _rhs368;
          _2073_v69 = _rhs369;
          let _2075_v71;
          _2075_v71 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_2070_v66).f20);
          _2075_v71 = (_2075_v71).update(_this.f16, (_2000_v14).IsSubsetOf(_2000_v14));
          let _index332 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          (_2005_v19)[_index332] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], (_dafny.ZERO).minus(_module.__default.fm0((_2070_v66).f19, globalState))));
          (globalState).f9 = (_dafny.ZERO).minus((_this).f15);
        } else {
          _2000_v14 = (_2000_v14).Union(_2000_v14);
          let _2076_v72;
          _2076_v72 = new _dafny.CodePoint('c'.codePointAt(0));
          let _2077_v73;
          _2077_v73 = _module.D5.create_DC14((_2070_v66).f19);
          let _2078_v74;
          _2078_v74 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm58(_2076_v72, (_2077_v73).dtor_cf26, globalState),p0);
          let _2079_v75;
          _2079_v75 = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-208)), function (_2080_i11) {
            return (_this).f15;
          }));
          let _2081_v76;
          _2081_v76 = _dafny.Seq.of(new BigNumber((_this.f16).length));
          _2078_v74 = (_2078_v74).update(_dafny.MultiSet.FromArray(_2081_v76), (_2070_v66).f19);
          let _2082_v77;
          let _init65 = ((_2083_v20, _2084_v19, _2085_v21, _2086_v64, _2087_v66) => function (_2088_i12) {
            return _dafny.Seq.Concat(_dafny.Seq.update(_2083_v20, _module.__default.safeIndex((_2084_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2084_v19).length))], new BigNumber((_2083_v20).length)), (((_2085_v21).contains((_this).f15)) ? ((_2085_v21).get((_this).f15)) : ((_2086_v64).f17))), _dafny.Seq.of((_2087_v66).f19, (_2086_v64).f17, (_2086_v64).f17));
          })(_2006_v20, _2005_v19, _2007_v21, _2067_v64, _2070_v66);
          let _nw343 = Array((new BigNumber(7)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw343.length); _i0_65++) {
            _nw343[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2082_v77 = _nw343;
          let _2089_v78;
          _2089_v78 = _dafny.MultiSet.fromElements((_this).f15);
          let _2090_v80;
          _2090_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2000_v14).cardinality()),_2089_v78);
          let _2091_v81;
          _2091_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2076_v72,_2089_v78);
          let _2092_v82;
          let _nw344 = Array((new BigNumber(18)).toNumber());
          _nw344[(_dafny.ZERO).toNumber()] = _2089_v78;
          _nw344[(_dafny.ONE).toNumber()] = (_2089_v78).Difference(_2089_v78);
          _nw344[(new BigNumber(2)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(3)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(4)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(5)).toNumber()] = (_2089_v78).update(new BigNumber((_dafny.Seq.of((_2067_v64).f17, (_2070_v66).f19)).length), _module.__default.abs(new BigNumber((_this.f16).length)));
          _nw344[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f15)).Difference(_dafny.MultiSet.fromElements((_this).f15, (_2070_v66).fm30(function () {
            let _coll67 = new _dafny.Map();
            for (const _compr_67 of _dafny.IntegerRange(new BigNumber(-154), new BigNumber(309))) {
              let _2093_v79 = _compr_67;
              if (((new BigNumber(-154)).isLessThanOrEqualTo(_2093_v79)) && ((_2093_v79).isLessThan(new BigNumber(309)))) {
                _coll67.push([_module.__default.safeModuloInt(_2093_v79, (_this).f15),(_2001_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length))]]);
              }
            }
            return _coll67;
          }(), globalState), (_this).f15));
          _nw344[(new BigNumber(7)).toNumber()] = _module.__default.fm38((_2070_v66).f19, (_2070_v66).f19, _2076_v72, (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], globalState);
          _nw344[(new BigNumber(8)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(9)).toNumber()] = (((_2090_v80).contains(new BigNumber(979))) ? ((_2090_v80).get(new BigNumber(979))) : (_2089_v78));
          _nw344[(new BigNumber(10)).toNumber()] = (((_2091_v81).contains(_2076_v72)) ? ((_2091_v81).get(_2076_v72)) : (_dafny.MultiSet.fromElements((_this).f15, (_this).f15)));
          _nw344[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.FromArray(_2081_v76)).Intersect(_dafny.MultiSet.FromArray(_2081_v76));
          _nw344[(new BigNumber(12)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], new BigNumber((_1999_v13).length));
          _nw344[(new BigNumber(14)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(15)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(16)).toNumber()] = _2089_v78;
          _nw344[(new BigNumber(17)).toNumber()] = _2089_v78;
          _2092_v82 = _nw344;
          let _index333 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_2092_v82).length));
          (_2092_v82)[_index333] = _2089_v78;
          let _index334 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          let _index335 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_2092_v82).length));
          let _rhs370 = _this.f16;
          let _rhs371 = _2082_v77;
          let _rhs372 = ((_this).f15).isLessThan((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          let _rhs373 = (((_2070_v66).f20) ? ((_this).f15) : ((_this).f15));
          let _rhs374 = ((_dafny.MultiSet.fromElements((_this).f15, (_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))])).Union(_2089_v78)).Union(_module.__default.fm56((_2070_v66).f19, (_this).f15, globalState));
          let _lhs282 = _this;
          let _lhs283 = _2001_v15;
          let _lhs284 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          let _lhs285 = globalState;
          let _lhs286 = _2092_v82;
          let _lhs287 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_2092_v82).length));
          _lhs282.f16 = _rhs370;
          _2082_v77 = _rhs371;
          _lhs283[_lhs284] = _rhs372;
          _lhs285.f9 = _rhs373;
          _lhs286[_lhs287] = _rhs374;
          let _2094_v83;
          let _nw345 = Array((new BigNumber(12)).toNumber()).fill(_module.D2.Default());
          _2094_v83 = _nw345;
          let _2095_v84;
          _2095_v84 = _dafny.Map.Empty.slice().updateUnsafe((_2067_v64).f17,(_2092_v82)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_2092_v82).length))]);
          let _index336 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          let _rhs375 = !(((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]).isLessThan((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]));
          let _rhs376 = _2094_v83;
          let _rhs377 = (((_2095_v84).contains((_2067_v64).f17)) ? ((_2095_v84).get((_2067_v64).f17)) : (((_2092_v82)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_2092_v82).length))]).update(new BigNumber(719), _module.__default.abs(new BigNumber((_this.f16).length)))));
          let _lhs288 = _2001_v15;
          let _lhs289 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_2001_v15).length));
          _lhs288[_lhs289] = _rhs375;
          _2094_v83 = _rhs376;
          _2089_v78 = _rhs377;
          let _2096_v87;
          let _init66 = ((_2097_v76, _2098_v66) => function (_2099_i13) {
            return function () {
              let _coll68 = new _dafny.Map();
              for (const _compr_68 of (_dafny.MultiSet.FromArray(_2097_v76)).Elements) {
                let _2100_v86 = _compr_68;
                if ((_dafny.MultiSet.FromArray(_2097_v76)).contains(_2100_v86)) {
                  _coll68.push([(_2100_v86).minus((_this).f15),(_2098_v66).f19]);
                }
              }
              return _coll68;
            }();
          })(_2081_v76, _2070_v66);
          let _nw346 = Array((new BigNumber(26)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw346.length); _i0_66++) {
            _nw346[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2096_v87 = _nw346;
          let _2101_v88;
          let _nw347 = new _module.C5();
          _nw347.__ctor(_2096_v87);
          _2101_v88 = _nw347;
          let _2102_v89;
          _2102_v89 = _dafny.MultiSet.fromElements(_2101_v88, _2101_v88);
          let _2103_v90;
          let _nw348 = Array((new BigNumber(22)).toNumber());
          _nw348[(_dafny.ZERO).toNumber()] = _2007_v21;
          _nw348[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_2070_v66).f19);
          _nw348[(new BigNumber(2)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(3)).toNumber()] = function () {
            let _coll69 = new _dafny.Map();
            for (const _compr_69 of _dafny.IntegerRange(new BigNumber(330), new BigNumber(75))) {
              let _2104_v85 = _compr_69;
              if (((new BigNumber(330)).isLessThanOrEqualTo(_2104_v85)) && ((_2104_v85).isLessThan(new BigNumber(75)))) {
                _coll69.push([_module.__default.safeModuloInt(_2104_v85, new BigNumber(((_module.D13.create_DC29(_2007_v21)).dtor_cf50).length)),(_2067_v64).f17]);
              }
            }
            return _coll69;
          }();
          _nw348[(new BigNumber(4)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(5)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(6)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(7)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_2070_v66).fm29((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))], globalState),(_2070_v66).f20);
          _nw348[(new BigNumber(9)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(10)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(11)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(12)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(13)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_2102_v89).contains(_2101_v88)) ? ((_2102_v89).get(_2101_v88)) : ((_this).f15)),(_2070_v66).f20);
          _nw348[(new BigNumber(15)).toNumber()] = (_2007_v21).update(new BigNumber(-153), (_2070_v66).f20);
          _nw348[(new BigNumber(16)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_2067_v64).f17);
          _nw348[(new BigNumber(18)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(19)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(20)).toNumber()] = _2007_v21;
          _nw348[(new BigNumber(21)).toNumber()] = _2007_v21;
          _2103_v90 = _nw348;
          let _2105_v91;
          let _nw349 = new _module.C4();
          _nw349.__ctor((_2101_v88).f11);
          _2105_v91 = _nw349;
        }
        let _2106_v92;
        let _init67 = function (_2107_i14) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        };
        let _nw350 = Array((new BigNumber(18)).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw350.length); _i0_67++) {
          _nw350[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2106_v92 = _nw350;
        let _2108_v93;
        _2108_v93 = new _dafny.CodePoint('d'.codePointAt(0));
        let _index337 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2106_v92).length));
        (_2106_v92)[_index337] = _2108_v93;
        let _index338 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2106_v92).length));
        (_2106_v92)[_index338] = _2108_v93;
        let _2109_v94;
        _2109_v94 = _module.D11.create_DC26((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
        let _source24 = _2109_v94;
        if (_source24.is_DC26) {
          let _2110___mcc_h0 = (_source24).cf47;
          let _2111_cf47 = _2110___mcc_h0;
          let _2112_v95;
          _2112_v95 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(new BigNumber(-545)).multipliedBy(_2111_cf47));
          let _2113_v96;
          _2113_v96 = _dafny.MultiSet.fromElements((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]);
          _2112_v95 = (_2112_v95).update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bj"), _module.__default.fm2(new BigNumber(-834), _2113_v96, function () {
            let _coll70 = new _dafny.Map();
            for (const _compr_70 of _dafny.IntegerRange(new BigNumber(609), new BigNumber(344))) {
              let _2114_v97 = _compr_70;
              if (((new BigNumber(609)).isLessThanOrEqualTo(_2114_v97)) && ((_2114_v97).isLessThan(new BigNumber(344)))) {
                _coll70.push([(_2114_v97).multipliedBy(_2111_cf47),(_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]]);
              }
            }
            return _coll70;
          }(), globalState)), _2111_cf47);
          (globalState).f9 = new BigNumber(601);
          let _2115_v98;
          _2115_v98 = _dafny.Map.Empty.slice().updateUnsafe(_2108_v93,p0);
          let _index339 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          let _rhs378 = !((_2070_v66).f20) || ((((_2115_v98).contains((_2106_v92)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_2106_v92).length))])) ? ((_2115_v98).get((_2106_v92)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_2106_v92).length))])) : ((_2070_v66).f20)));
          let _rhs379 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_this).f15, _2111_cf47)).plus((_2005_v19)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length))]));
          let _lhs290 = globalState;
          let _lhs291 = _2005_v19;
          let _lhs292 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2005_v19).length));
          _lhs290.f8 = _rhs378;
          _lhs291[_lhs292] = _rhs379;
          _2002_v16 = (_2002_v16).update((_2070_v66).f20, true);
        } else if (_source24.is_DC25) {
          let _2116___mcc_h1 = (_source24).cf46;
          let _2117_cf46 = _2116___mcc_h1;
          (globalState).f8 = _module.__default.fm1(_module.__default.fm1((_2067_v64).f17, _2108_v93, globalState), (_2106_v92)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_2106_v92).length))], globalState);
          let _2118_v99;
          let _init68 = ((_2119_v21) => function (_2120_i15) {
            return _2119_v21;
          })(_2007_v21);
          let _nw351 = Array((new BigNumber(2)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw351.length); _i0_68++) {
            _nw351[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2118_v99 = _nw351;
          let _2121_v100;
          let _nw352 = new _module.C2();
          _nw352.__ctor(_2118_v99);
          _2121_v100 = _nw352;
          (globalState).f8 = (_2000_v14).IsDisjointFrom(_dafny.MultiSet.FromArray(_2006_v20));
          let _2122_v101;
          let _nw353 = new _module.C3();
          _nw353.__ctor(!((_this).f15).isEqualTo(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-388)), function (_2123_i16) {
            return (_this).f15;
          }))).cardinality())), p0, _2118_v99);
          _2122_v101 = _nw353;
        } else {
          let _2124___mcc_h2 = (_source24).cf48;
          let _2125_cf48 = _2124___mcc_h2;
          (globalState).f8 = (((_2067_v64).f17) ? ((_2070_v66).f20) : ((_2070_v66).f20));
          let _2126_v102;
          _2126_v102 = _module.D1.create_DC4(new _dafny.CodePoint('f'.codePointAt(0)));
          _2126_v102 = _2126_v102;
          r0 = _2108_v93;
          let _2127_v103;
          let _nw354 = Array((_dafny.ONE).toNumber());
          _2127_v103 = _nw354;
          let _2128_v104;
          let _init69 = ((_2129_v21) => function (_2130_i17) {
            return _2129_v21;
          })(_2007_v21);
          let _nw355 = Array((new BigNumber(5)).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw355.length); _i0_69++) {
            _nw355[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _2128_v104 = _nw355;
          let _2131_v105;
          let _nw356 = new _module.C5();
          _nw356.__ctor(_2128_v104);
          _2131_v105 = _nw356;
          let _index340 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_2127_v103).length));
          (_2127_v103)[_index340] = _2131_v105;
          let _index341 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_2127_v103).length));
          (_2127_v103)[_index341] = _2131_v105;
        }
      }
      let _2132_v106;
      _2132_v106 = new _dafny.CodePoint('g'.codePointAt(0));
      r0 = _2132_v106;
      return r0;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.of();
      let r3 = false;
      let _2133_v0;
      _2133_v0 = _dafny.Seq.of(p1);
      let _2134_v1;
      _2134_v1 = _dafny.Set.fromElements(p0, (_this).f15, new BigNumber((_2133_v0).length), new BigNumber(854));
      let _hi15 = new BigNumber((_2134_v1).length);
      for (let _2135_i0 = new BigNumber(100); _2135_i0.isLessThan(_hi15); _2135_i0 = _2135_i0.plus(_dafny.ONE)) {
        let _2136_v3;
        let _nw357 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2136_v3 = _nw357;
        let _2137_v4;
        _2137_v4 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll71 = new _dafny.Map();
          for (const _compr_71 of _dafny.IntegerRange(new BigNumber(692), new BigNumber(698))) {
            let _2138_v2 = _compr_71;
            if (((new BigNumber(692)).isLessThanOrEqualTo(_2138_v2)) && ((_2138_v2).isLessThan(new BigNumber(698)))) {
              _coll71.push([(_2138_v2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length)),_dafny.Map.Empty.slice().updateUnsafe(p1,p0)]);
            }
          }
          return _coll71;
        }(),_2136_v3);
        let _2139_v5;
        _2139_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f15);
        _2137_v4 = (_2137_v4).update(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2139_v5), _2136_v3);
        let _2140_v6;
        _2140_v6 = _dafny.Seq.of(_2135_i0, _2135_i0);
        let _2141_v7;
        _2141_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2140_v6).length),_dafny.Seq.of(p0, new BigNumber((_this.f16).length), new BigNumber(165)));
        let _2142_v8;
        _2142_v8 = _dafny.MultiSet.fromElements(new BigNumber(-357));
        let _2143_v9;
        _2143_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _2144_v10;
        _2144_v10 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p1,p1), _2143_v9, _2143_v9, _2143_v9);
        let _2145_v11;
        _2145_v11 = _2144_v10;
        (globalState).f9 = ((_dafny.ZERO).minus(new BigNumber(((((_2141_v7).contains(new BigNumber((_2142_v8).cardinality()))) ? ((_2141_v7).get(new BigNumber((_2142_v8).cardinality()))) : (_dafny.Seq.of(new BigNumber((_2142_v8).cardinality()), p0)))).length))).plus(new BigNumber(((_dafny.Set.fromElements(_2143_v9))).length));
        _module.__default.m0(globalState);
        let _2146_v12;
        let _nw358 = Array((new BigNumber(24)).toNumber()).fill([]);
        _2146_v12 = _nw358;
        let _index342 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2146_v12).length));
        (_2146_v12)[_index342] = _2136_v3;
        let _index343 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2146_v12).length));
        let _rhs380 = _2135_i0;
        let _rhs381 = _this.f16;
        let _rhs382 = _this.f16;
        let _rhs383 = _module.__default.safeDivisionInt(p0, p0);
        let _rhs384 = (((p0).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("vrgni")).length))) ? (_2136_v3) : (_2136_v3));
        let _lhs293 = _this;
        let _lhs294 = _this;
        let _lhs295 = globalState;
        let _lhs296 = _2146_v12;
        let _lhs297 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2146_v12).length));
        r1 = _rhs380;
        _lhs293.f16 = _rhs381;
        _lhs294.f16 = _rhs382;
        _lhs295.f9 = _rhs383;
        _lhs296[_lhs297] = _rhs384;
      }
      r0 = (_dafny.ZERO).minus(((_this).f15).plus((_this).f15));
      let _2147_i1;
      _2147_i1 = _dafny.ZERO;
      L10: {
        while (p1) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2147_i1)) {
              break L10;
            }
            _2147_i1 = (_2147_i1).plus(_dafny.ONE);
            (globalState).f8 = p1;
            let _2148_v13;
            _2148_v13 = new _dafny.CodePoint('j'.codePointAt(0));
            let _2149_v14;
            _2149_v14 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm39(p0, globalState)).length), (_this).f15);
            let _2150_v15;
            _2150_v15 = _dafny.Seq.of(p0, (_this).f15, p0, p0, new BigNumber((_this.f16).length));
            let _2151_v16;
            let _nw359 = Array((new BigNumber(26)).toNumber());
            _nw359[(_dafny.ZERO).toNumber()] = p1;
            _nw359[(_dafny.ONE).toNumber()] = p1;
            _nw359[(new BigNumber(2)).toNumber()] = p1;
            _nw359[(new BigNumber(3)).toNumber()] = p1;
            _nw359[(new BigNumber(4)).toNumber()] = p1;
            _nw359[(new BigNumber(5)).toNumber()] = _module.__default.fm1(p1, new _dafny.CodePoint('f'.codePointAt(0)), globalState);
            _nw359[(new BigNumber(6)).toNumber()] = p1;
            _nw359[(new BigNumber(7)).toNumber()] = (new BigNumber((_this.f16).length)).isLessThan(p0);
            _nw359[(new BigNumber(8)).toNumber()] = !((_2134_v1).IsDisjointFrom(_2134_v1));
            _nw359[(new BigNumber(9)).toNumber()] = p1;
            _nw359[(new BigNumber(10)).toNumber()] = (_2133_v0)[_module.__default.safeIndex((_this).f15, new BigNumber((_2133_v0).length))];
            _nw359[(new BigNumber(11)).toNumber()] = p1;
            _nw359[(new BigNumber(12)).toNumber()] = false;
            _nw359[(new BigNumber(13)).toNumber()] = p1;
            _nw359[(new BigNumber(14)).toNumber()] = p1;
            _nw359[(new BigNumber(15)).toNumber()] = p1;
            _nw359[(new BigNumber(16)).toNumber()] = p1;
            _nw359[(new BigNumber(17)).toNumber()] = _module.__default.fm1(p1, _2148_v13, globalState);
            _nw359[(new BigNumber(18)).toNumber()] = ((_module.__default.fm38(p1, p1, (_this.f16)[_module.__default.safeIndex((_this).f15, new BigNumber((_this.f16).length))], p0, globalState)).update((_this).f15, _module.__default.abs(p0))).IsProperSubsetOf(_2149_v14);
            _nw359[(new BigNumber(19)).toNumber()] = p1;
            _nw359[(new BigNumber(20)).toNumber()] = (new BigNumber((_2150_v15).length)).isLessThanOrEqualTo(p0);
            _nw359[(new BigNumber(21)).toNumber()] = _dafny.Seq.contains(_this.f16, (_this.f16)[_module.__default.safeIndex(new BigNumber((_2149_v14).cardinality()), new BigNumber((_this.f16).length))]);
            _nw359[(new BigNumber(22)).toNumber()] = p1;
            _nw359[(new BigNumber(23)).toNumber()] = p1;
            _nw359[(new BigNumber(24)).toNumber()] = p1;
            _nw359[(new BigNumber(25)).toNumber()] = p1;
            _2151_v16 = _nw359;
            let _index344 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2151_v16).length));
            (_2151_v16)[_index344] = p1;
            let _index345 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2151_v16).length));
            (_2151_v16)[_index345] = ((p1) ? (p1) : (p1));
            let _index346 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2151_v16).length));
            (_2151_v16)[_index346] = p1;
            let _2152_v17;
            let _nw360 = Array((new BigNumber(6)).toNumber());
            _2152_v17 = _nw360;
            let _2153_v18;
            _2153_v18 = _2152_v17;
            let _2154_v19;
            _2154_v19 = _dafny.Seq.of(_2152_v17, _2152_v17);
            let _2155_v20;
            let _nw361 = Array((new BigNumber(29)).toNumber());
            _nw361[(_dafny.ZERO).toNumber()] = _2152_v17;
            _nw361[(_dafny.ONE).toNumber()] = (_2153_v18);
            _nw361[(new BigNumber(2)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(3)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(4)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(5)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(6)).toNumber()] = (_2154_v19)[_module.__default.safeIndex((_this).f15, new BigNumber((_2154_v19).length))];
            _nw361[(new BigNumber(7)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(8)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(9)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(10)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(11)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(12)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(13)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(14)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(15)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(16)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(17)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(18)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(19)).toNumber()] = ((p1) ? (_2152_v17) : (_2152_v17));
            _nw361[(new BigNumber(20)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(21)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(22)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(23)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(24)).toNumber()] = ((false) ? (_2152_v17) : (_2152_v17));
            _nw361[(new BigNumber(25)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(26)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(27)).toNumber()] = _2152_v17;
            _nw361[(new BigNumber(28)).toNumber()] = _2152_v17;
            _2155_v20 = _nw361;
            let _index347 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_2155_v20).length));
            (_2155_v20)[_index347] = _2152_v17;
            let _index348 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2151_v16).length));
            let _index349 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_2155_v20).length));
            let _rhs385 = (p0).isLessThan(p0);
            let _rhs386 = _2152_v17;
            let _rhs387 = _2151_v16;
            let _rhs388 = new BigNumber((_module.__default.fm59(globalState)).length);
            let _lhs298 = _2151_v16;
            let _lhs299 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2151_v16).length));
            let _lhs300 = _2155_v20;
            let _lhs301 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_2155_v20).length));
            _lhs298[_lhs299] = _rhs385;
            _lhs300[_lhs301] = _rhs386;
            _2151_v16 = _rhs387;
            r1 = _rhs388;
          }
        }
      }
      let _2156_v21;
      _2156_v21 = _dafny.MultiSet.fromElements((_this).f15);
      if (((_this).f15).isLessThan((((_2156_v21).contains(p0)) ? ((_2156_v21).get(p0)) : ((_this).f15)))) {
        let _2157_v22;
        let _nw362 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _2157_v22 = _nw362;
        let _index350 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length));
        (_2157_v22)[_index350] = p0;
        let _2158_v23;
        let _nw363 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2158_v23 = _nw363;
        let _index351 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length));
        let _rhs389 = _this.f16;
        let _rhs390 = (_this).f15;
        let _rhs391 = (_this).f15;
        let _rhs392 = _2158_v23;
        let _lhs302 = _this;
        let _lhs303 = _2157_v22;
        let _lhs304 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length));
        _lhs302.f16 = _rhs389;
        _lhs303[_lhs304] = _rhs390;
        r0 = _rhs391;
        _2158_v23 = _rhs392;
        let _2159_v24;
        _2159_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
        let _2160_v25;
        _2160_v25 = _module.D9.create_DC22((_this).f15, _2159_v24);
        r0 = (_2160_v25).dtor_cf40;
        let _2161_v26;
        _2161_v26 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _2162_v27;
        _2162_v27 = _dafny.Seq.of(_2161_v26);
        let _2163_v28;
        _2163_v28 = _module.D0.create_DC1(_dafny.MultiSet.fromElements(_2157_v22, _2157_v22), _module.__default.fm0(p1, globalState), new BigNumber((_2162_v27).length));
        if (!((((_2133_v0)[_module.__default.safeIndex((_2163_v28).dtor_cf3, new BigNumber((_2133_v0).length))]) ? (p1) : (p1)))) {
          let _2164_v29;
          let _init70 = ((_2165_p1) => function (_2166_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2165_p1);
          })(p1);
          let _nw364 = Array((new BigNumber(21)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw364.length); _i0_70++) {
            _nw364[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2164_v29 = _nw364;
          let _2167_v30;
          let _nw365 = new _module.C5();
          _nw365.__ctor(_2164_v29);
          _2167_v30 = _nw365;
          _2167_v30 = _2167_v30;
          r1 = ((_2157_v22)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length))]).multipliedBy(p0);
          let _2168_v31;
          _2168_v31 = new _dafny.CodePoint('s'.codePointAt(0));
          let _2169_v32;
          _2169_v32 = _module.D0.create_DC0(p0);
          let _2170_v33;
          _2170_v33 = _dafny.Map.Empty.slice().updateUnsafe((_2169_v32).dtor_cf0,_2168_v31);
          _2168_v31 = (((_2170_v33).contains(_module.__default.fm0(true, globalState))) ? ((_2170_v33).get(_module.__default.fm0(true, globalState))) : (new _dafny.CodePoint('p'.codePointAt(0))));
          let _2171_v34;
          let _nw366 = Array((new BigNumber(14)).toNumber()).fill([]);
          _2171_v34 = _nw366;
          let _2172_v35;
          let _init71 = function (_2173_i3) {
            return true;
          };
          let _nw367 = Array((new BigNumber(16)).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw367.length); _i0_71++) {
            _nw367[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _2172_v35 = _nw367;
          let _index352 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_2171_v34).length));
          (_2171_v34)[_index352] = _2172_v35;
          let _2174_v36;
          _2174_v36 = _module.D10.create_DC24(_2172_v35, _2157_v22, _this.f16);
          let _index353 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_2171_v34).length));
          (_2171_v34)[_index353] = (_2174_v36).dtor_cf43;
          let _2175_v37;
          _2175_v37 = _2161_v26;
          let _2176_v38;
          _2176_v38 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f15));
          let _2177_v39;
          let _nw368 = Array((new BigNumber(28)).toNumber());
          _nw368[(_dafny.ZERO).toNumber()] = _2161_v26;
          _nw368[(_dafny.ONE).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(2)).toNumber()] = (_2175_v37);
          _nw368[(new BigNumber(3)).toNumber()] = (_2161_v26).Merge(_2161_v26);
          _nw368[(new BigNumber(4)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(5)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(6)).toNumber()] = (_2162_v27)[_module.__default.safeIndex(new BigNumber((_2176_v38).length), new BigNumber((_2162_v27).length))];
          _nw368[(new BigNumber(7)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(8)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(9)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(10)).toNumber()] = (_2161_v26).Merge(_2161_v26);
          _nw368[(new BigNumber(11)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(12)).toNumber()] = (_2161_v26).update(p1, p1);
          _nw368[(new BigNumber(13)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(14)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(15)).toNumber()] = ((true) ? (_dafny.Map.Empty.slice().updateUnsafe(p1,p1)) : (_2161_v26));
          _nw368[(new BigNumber(16)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(17)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm1(p1, _2168_v31, globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,p1));
          _nw368[(new BigNumber(18)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(19)).toNumber()] = ((_2162_v27)[_module.__default.safeIndex((_this).f15, new BigNumber((_2162_v27).length))]).Merge((_2162_v27)[_module.__default.safeIndex(new BigNumber((_2133_v0).length), new BigNumber((_2162_v27).length))]);
          _nw368[(new BigNumber(20)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(21)).toNumber()] = ((p1) ? (_2161_v26) : (_2161_v26));
          _nw368[(new BigNumber(22)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(23)).toNumber()] = (_2161_v26).update(p1, !((((_2161_v26).contains(p1)) ? ((_2161_v26).get(p1)) : (p1))));
          _nw368[(new BigNumber(24)).toNumber()] = (_2175_v37);
          _nw368[(new BigNumber(25)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(26)).toNumber()] = _2161_v26;
          _nw368[(new BigNumber(27)).toNumber()] = _2161_v26;
          _2177_v39 = _nw368;
          let _index354 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_2177_v39).length));
          (_2177_v39)[_index354] = _2161_v26;
          let _index355 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_2177_v39).length));
          (_2177_v39)[_index355] = (_2161_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,p1));
        } else {
          let _2178_v40;
          _2178_v40 = _dafny.Set.fromElements(p1);
          (globalState).f8 = (_2178_v40).IsDisjointFrom(_2178_v40);
          let _2179_v41;
          _2179_v41 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          _2179_v41 = (_2179_v41).update((p1) && (p1), p0);
          let _2180_v42;
          _2180_v42 = _dafny.Map.Empty.slice().updateUnsafe((_2157_v22)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length))],p1);
          let _2181_v45;
          _2181_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f15);
          let _2182_v46;
          let _nw369 = Array((new BigNumber(27)).toNumber());
          _nw369[(_dafny.ZERO).toNumber()] = _2180_v42;
          _nw369[(_dafny.ONE).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(2)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(3)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(4)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(5)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(6)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((((_2161_v26).contains(p1)) ? ((_2161_v26).get(p1)) : (p1)), globalState),p1);
          _nw369[(new BigNumber(8)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(9)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(10)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(11)).toNumber()] = function () {
            let _coll72 = new _dafny.Map();
            for (const _compr_72 of _dafny.IntegerRange(new BigNumber(194), new BigNumber(972))) {
              let _2183_v43 = _compr_72;
              if (((new BigNumber(194)).isLessThanOrEqualTo(_2183_v43)) && ((_2183_v43).isLessThan(new BigNumber(972)))) {
                _coll72.push([_module.__default.safeDivisionInt(_2183_v43, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length)),p1]);
              }
            }
            return _coll72;
          }();
          _nw369[(new BigNumber(12)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(13)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(14)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(15)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(16)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(17)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(18)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(19)).toNumber()] = function () {
            let _coll73 = new _dafny.Map();
            for (const _compr_73 of (_2181_v45).Keys.Elements) {
              let _2184_v44 = _compr_73;
              if ((_2181_v45).contains(_2184_v44)) {
                _coll73.push([(_2184_v44).plus((_this).f15),p1]);
              }
            }
            return _coll73;
          }();
          _nw369[(new BigNumber(20)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(21)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(22)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(23)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(24)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(25)).toNumber()] = _2180_v42;
          _nw369[(new BigNumber(26)).toNumber()] = _2180_v42;
          _2182_v46 = _nw369;
          let _2185_v47;
          let _nw370 = new _module.C4();
          _nw370.__ctor(((p1) ? (_2182_v46) : (_2182_v46)));
          _2185_v47 = _nw370;
          let _2186_v48;
          let _nw371 = Array((new BigNumber(10)).toNumber()).fill(false);
          _2186_v48 = _nw371;
          let _2187_v49;
          _2187_v49 = _module.D7.create_DC17(_this.f16, _2186_v48, !(p1), new BigNumber(-549));
          let _2188_v50;
          _2188_v50 = _dafny.Set.fromElements(_2157_v22, _2157_v22);
          let _2189_v51;
          _2189_v51 = _module.D7.create_DC17(_dafny.Seq.UnicodeFromString("v"), (_2187_v49).dtor_cf30, p1, new BigNumber((_2188_v50).length));
          let _2190_v52;
          _2190_v52 = _dafny.Set.fromElements((_2187_v49).dtor_cf30);
          _2190_v52 = _2190_v52;
          let _2191_v53;
          let _nw372 = new _module.C5();
          _nw372.__ctor(_2182_v46);
          _2191_v53 = _nw372;
        }
        let _2192_v54;
        let _init72 = function (_2193_i4) {
          return false;
        };
        let _nw373 = Array((new BigNumber(4)).toNumber());
        for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw373.length); _i0_72++) {
          _nw373[_i0_72] = _init72(new BigNumber(_i0_72));
        }
        _2192_v54 = _nw373;
        let _2194_v55;
        let _nw374 = Array((new BigNumber(11)).toNumber());
        _nw374[(_dafny.ZERO).toNumber()] = _2192_v54;
        _nw374[(_dafny.ONE).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(2)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(3)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(4)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(5)).toNumber()] = ((p1) ? (_2192_v54) : (_2192_v54));
        _nw374[(new BigNumber(6)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(7)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(8)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(9)).toNumber()] = _2192_v54;
        _nw374[(new BigNumber(10)).toNumber()] = _2192_v54;
        _2194_v55 = _nw374;
        let _index356 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_2194_v55).length));
        (_2194_v55)[_index356] = _2192_v54;
        let _2195_v56;
        _2195_v56 = _dafny.Set.fromElements(p1, p1);
        let _2196_v57;
        _2196_v57 = _dafny.Map.Empty.slice().updateUnsafe(_2192_v54,(_2157_v22)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length))]);
        let _2197_v58;
        _2197_v58 = new _dafny.CodePoint('w'.codePointAt(0));
        let _2198_v59;
        _2198_v59 = _module.D2.create_DC8((_dafny.ZERO).minus(new BigNumber((_2196_v57).length)), _2197_v58);
        let _pat_let_tv38 = _2157_v22;
        let _pat_let_tv39 = _2157_v22;
        let _index357 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length));
        let _index358 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_2194_v55).length));
        let _rhs393 = (_module.__default.fm0(p1, globalState)).multipliedBy((new BigNumber((_2195_v56).length)).multipliedBy((_this).f15));
        let _rhs394 = _2192_v54;
        let _rhs395 = (function (_pat_let31_0) {
          return function (_2199_dt__update__tmp_h1) {
            return function (_pat_let32_0) {
              return function (_2200_dt__update_hcf15_h0) {
                return _module.D2.create_DC8(_2200_dt__update_hcf15_h0, (_2199_dt__update__tmp_h1).dtor_cf16);
              }(_pat_let32_0);
            }((_pat_let_tv39)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_pat_let_tv38).length))]);
          }(_pat_let31_0);
        }(_2198_v59)).dtor_cf15;
        let _lhs305 = _2157_v22;
        let _lhs306 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length));
        let _lhs307 = _2194_v55;
        let _lhs308 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_2194_v55).length));
        let _lhs309 = globalState;
        _lhs305[_lhs306] = _rhs393;
        _lhs307[_lhs308] = _rhs394;
        _lhs309.f9 = _rhs395;
        let _2201_v60;
        _2201_v60 = _dafny.Seq.of(p0);
        let _2202_v61;
        _2202_v61 = _dafny.Seq.of(new BigNumber((_2195_v56).length), (_2157_v22)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length))], (_dafny.ZERO).minus((_2201_v60)[_module.__default.safeIndex((_2157_v22)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_2157_v22).length))], new BigNumber((_2201_v60).length))]), new BigNumber(967));
        let _2203_v62;
        _2203_v62 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber(786)), _2201_v60, _dafny.Seq.of(new BigNumber(-650)));
        (globalState).f9 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2201_v60, (_2203_v62)[_module.__default.safeIndex((_this).f15, new BigNumber((_2203_v62).length))])).length));
      } else {
        let _2204_v63;
        _2204_v63 = new _dafny.CodePoint('k'.codePointAt(0));
        (globalState).f8 = _module.__default.fm1(p1, _2204_v63, globalState);
        let _2205_v64;
        let _init73 = ((_2206_p1) => function (_2207_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2206_p1);
        })(p1);
        let _nw375 = Array((new BigNumber(22)).toNumber());
        for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw375.length); _i0_73++) {
          _nw375[_i0_73] = _init73(new BigNumber(_i0_73));
        }
        _2205_v64 = _nw375;
        let _2208_v65;
        let _nw376 = new _module.C4();
        _nw376.__ctor(_2205_v64);
        _2208_v65 = _nw376;
        let _2209_v66;
        _2209_v66 = _dafny.Seq.of(p0);
        let _2210_v67;
        _2210_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2208_v65,new BigNumber((_2209_v66).length));
        let _2211_v68;
        _2211_v68 = _dafny.MultiSet.fromElements(_2210_v67);
        let _2212_v69;
        _2212_v69 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2208_v65,(_this).f15));
        _2211_v68 = _dafny.MultiSet.FromArray(_2212_v69);
        (globalState).f9 = p0;
        let _2213_v70;
        _2213_v70 = _dafny.Set.fromElements(p1);
        let _2214_v71;
        _2214_v71 = _module.D0.create_DC2(_dafny.MultiSet.fromElements(!(p1)), false, (_dafny.ZERO).minus(new BigNumber((_2213_v70).length)), p1);
        let _2215_v72;
        _2215_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2214_v71,p1);
        let _2216_v73;
        _2216_v73 = _dafny.Map.Empty.slice().updateUnsafe((_2215_v72).contains(_2214_v71),p1);
        _2216_v73 = (_2216_v73).update(p1, p1);
        r3 = p1;
      }
      let _hi16 = p0;
      for (let _2217_i6 = new BigNumber((_dafny.Seq.Concat(_2133_v0, _dafny.Seq.update(_2133_v0, _module.__default.safeIndex((_this).f15, new BigNumber((_2133_v0).length)), p1))).length); _2217_i6.isLessThan(_hi16); _2217_i6 = _2217_i6.plus(_dafny.ONE)) {
        _module.__default.m0(globalState);
        let _2218_v74;
        _2218_v74 = _dafny.Seq.of(p0);
        let _2219_v75;
        _2219_v75 = _module.D0.create_DC0(p0);
        let _2220_v76;
        let _nw377 = new _module.C1();
        _nw377.__ctor(p1, p1, _2219_v75, p1);
        _2220_v76 = _nw377;
        let _2221_v77;
        _2221_v77 = _dafny.Seq.of(_2220_v76);
        let _2222_v78;
        _2222_v78 = _dafny.Seq.of(_2217_i6, ((_this).f15).plus((_2218_v74)[_module.__default.safeIndex(new BigNumber((_2221_v77).length), new BigNumber((_2218_v74).length))]), _module.__default.safeModuloInt((_this).f15, p0), p0);
        let _2223_v79;
        let _nw378 = Array((new BigNumber(3)).toNumber()).fill(false);
        _2223_v79 = _nw378;
        let _index359 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2223_v79).length));
        (_2223_v79)[_index359] = _dafny.Seq.IsPrefixOf(_this.f16, _dafny.Seq.UnicodeFromString("hapbpk"));
        let _index360 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2223_v79).length));
        let _rhs396 = !(true);
        let _rhs397 = (_2220_v76).f20;
        let _rhs398 = _2218_v74;
        let _rhs399 = ((_2220_v76).f20) || (p1);
        let _lhs310 = globalState;
        let _lhs311 = _2223_v79;
        let _lhs312 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2223_v79).length));
        _lhs310.f8 = _rhs396;
        r3 = _rhs397;
        _2218_v74 = _rhs398;
        _lhs311[_lhs312] = _rhs399;
        let _2224_v80;
        _2224_v80 = _dafny.MultiSet.fromElements((_2220_v76).f20);
        let _2225_v81;
        _2225_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2223_v79,_2224_v80);
        _2225_v81 = (_2225_v81).update(_2223_v79, _2224_v80);
        r3 = (_2220_v76).f19;
      }
      let _hi17 = (_this).f15;
      for (let _2226_i7 = p0; _2226_i7.isLessThan(_hi17); _2226_i7 = _2226_i7.plus(_dafny.ONE)) {
        let _2227_v82;
        _2227_v82 = _dafny.MultiSet.fromElements(!(p1), p1);
        let _2228_v83;
        _2228_v83 = _module.D0.create_DC2(_2227_v82, p1, p0, p1);
        let _2229_v84;
        _2229_v84 = _module.D8.create_DC20(_2228_v83, p1, p1);
        r3 = (_2229_v84).dtor_cf38;
        (globalState).f8 = p1;
        let _2230_v85;
        _2230_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2226_i7);
        let _2231_v86;
        _2231_v86 = _dafny.Seq.of(new BigNumber(32));
        (globalState).f8 = ((new BigNumber((_2231_v86).length)).plus((_this).f15)).isLessThanOrEqualTo((new BigNumber((_2230_v85).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), function (_2232_i8) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        })).length))));
        let _2233_v87;
        _2233_v87 = new _dafny.CodePoint('e'.codePointAt(0));
        r3 = !(p1) || (_module.__default.fm1(p1, _2233_v87, globalState));
      }
      r0 = (_this).f15;
      r1 = p0;
      let _2234_v88;
      _2234_v88 = _module.D2.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), ((_2235_p0) => function (_2236_i9) {
  return (_this.f16)[_module.__default.safeIndex((_dafny.ZERO).minus(_2235_p0), new BigNumber((_this.f16).length))];
})(p0)));
      let _2237_v89;
      _2237_v89 = _dafny.Seq.of(_2234_v88, _2234_v88, function (_pat_let33_0) {
        return function (_2238_dt__update__tmp_h2) {
          return function (_pat_let34_0) {
            return function (_2239_dt__update_hcf10_h0) {
              return _module.D2.create_DC5(_2239_dt__update_hcf10_h0);
            }(_pat_let34_0);
          }(_this.f16);
        }(_pat_let33_0);
      }(_2234_v88), function (_pat_let35_0) {
        return function (_2240_dt__update__tmp_h3) {
          return function (_pat_let36_0) {
            return function (_2241_dt__update_hcf10_h1) {
              return _module.D2.create_DC5(_2241_dt__update_hcf10_h1);
            }(_pat_let36_0);
          }(_this.f16);
        }(_pat_let35_0);
      }(_module.D2.create_DC5(_this.f16)));
      r2 = _2237_v89;
      r3 = p1;
      return [r0, r1, r2, r3];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f12 = _module.D0.Default();
      this._f11 = [];
      this._f13 = false;
      this._f14 = false;
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
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f14, f11, f12, f13) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm8(globalState) {
      let _this = this;
      let _source25 = _module.D0.create_DC2(_dafny.MultiSet.fromElements(!((_this).f14), (_this).f13), (_this).f14, new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-158)))).cardinality()), (_this).f14);
      if (_source25.is_DC1) {
        let _2242___mcc_h0 = (_source25).cf1;
        let _2243___mcc_h1 = (_source25).cf2;
        let _2244___mcc_h2 = (_source25).cf3;
        let _2245_cf3 = _2244___mcc_h2;
        let _2246_cf2 = _2243___mcc_h1;
        let _2247_cf1 = _2242___mcc_h0;
        return _module.D2.create_DC6();
      } else if (_source25.is_DC2) {
        let _2248___mcc_h3 = (_source25).cf4;
        let _2249___mcc_h4 = (_source25).cf5;
        let _2250___mcc_h5 = (_source25).cf6;
        let _2251___mcc_h6 = (_source25).cf7;
        let _2252_cf7 = _2251___mcc_h6;
        let _2253_cf6 = _2250___mcc_h5;
        let _2254_cf5 = _2249___mcc_h4;
        let _2255_cf4 = _2248___mcc_h3;
        return _module.D2.create_DC6();
      } else {
        let _2256___mcc_h7 = (_source25).cf0;
        let _2257_cf0 = _2256___mcc_h7;
        return _module.D2.create_DC6();
      }
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("equcqbrac")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)))).cardinality()))), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bxaa"), _dafny.Seq.UnicodeFromString("lgrqsywy"))).length));
    };
    fm10(globalState) {
      let _this = this;
      return (_this).f14;
    };
    fm11(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-656)),_dafny.Seq.of((_this).f14))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(485),_dafny.Seq.of((_this).f14, (_this).f13)))).length), (new BigNumber((_dafny.Seq.UnicodeFromString("rvymo")).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("x")).length))));
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _module.D2.Default();
      let r3 = false;
      let _2258_v0;
      _2258_v0 = _dafny.Seq.UnicodeFromString("owbxdur");
      let _2259_v1;
      _2259_v1 = _module.D1.create_DC4((_2258_v0)[_module.__default.safeIndex(p0, new BigNumber((_2258_v0).length))]);
      let _2260_v3;
      _2260_v3 = new _dafny.CodePoint('v'.codePointAt(0));
      let _2261_v4;
      _2261_v4 = _dafny.MultiSet.fromElements((_this).f14);
      let _2262_v5;
      _2262_v5 = _dafny.Set.fromElements((_this).f13);
      let _2263_v6;
      _2263_v6 = _module.D0.create_DC2(_2261_v4, false, new BigNumber((_2262_v5).length), (_this).f14);
      let _2264_v7;
      _2264_v7 = _dafny.Seq.of(new BigNumber((_2258_v0).length));
      let _2265_v8;
      _2265_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2264_v7,p0);
      let _2266_v9;
      _2266_v9 = _dafny.MultiSet.fromElements(p0);
      let _2267_v10;
      _2267_v10 = _dafny.MultiSet.fromElements((((_2265_v8).contains(_2264_v7)) ? ((_2265_v8).get(_2264_v7)) : (new BigNumber(((_2266_v9)).cardinality()))));
      let _2268_v11;
      _2268_v11 = _dafny.Set.fromElements((_this).fm11(_module.__default.fm1((_this).f13, _2260_v3, globalState), false, p0, (_2263_v6).dtor_cf5, globalState), (((_2267_v10).contains(p0)) ? ((_2267_v10).get(p0)) : (new BigNumber((_2258_v0).length))), p0);
      let _2269_v12;
      _2269_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm9(_2259_v1, function () {
        let _coll74 = new _dafny.Set();
        for (const _compr_74 of _dafny.IntegerRange(new BigNumber(-48), new BigNumber(590))) {
          let _2270_v2 = _compr_74;
          if (((new BigNumber(-48)).isLessThanOrEqualTo(_2270_v2)) && ((_2270_v2).isLessThan(new BigNumber(590)))) {
            _coll74.add(_module.__default.safeModuloInt(_2270_v2, p0));
          }
        }
        return _coll74;
      }(), globalState),(_2268_v11).IsSubsetOf(_2268_v11));
      _2269_v12 = (_2269_v12).update(p0, (_this).f14);
      let _2271_v13;
      let _nw379 = Array((new BigNumber(25)).toNumber()).fill(false);
      _2271_v13 = _nw379;
      let _index361 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
      (_2271_v13)[_index361] = (_this).f13;
      let _2272_v14;
      _2272_v14 = _dafny.MultiSet.fromElements(p1, p1, p1);
      let _index362 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
      (_2271_v13)[_index362] = (_2272_v14).IsProperSubsetOf((_2272_v14).update(p1, _module.__default.abs(new BigNumber((_2258_v0).length))));
      let _2273_v15;
      let _nw380 = Array((new BigNumber(14)).toNumber()).fill(_module.D1.Default());
      _2273_v15 = _nw380;
      _2273_v15 = _2273_v15;
      let _2274_v16;
      _2274_v16 = _module.D2.create_DC6();
      let _source26 = _2274_v16;
      if (_source26.is_DC6) {
        let _2275_v17;
        _2275_v17 = _module.D2.create_DC7((_this).f14, p0, (_this).f13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), ((_2276_v3) => function (_2277_i0) {
  return _2276_v3;
})(_2260_v3)));
        let _rhs400 = ((_2261_v4).Difference(_dafny.MultiSet.fromElements((_this).f13, (_this).f14))).Difference((_2263_v6).dtor_cf4);
        let _rhs401 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.fm0(!((_this).f13), globalState)), (_2275_v17).dtor_cf12);
        let _rhs402 = _module.__default.fm0((_this).f13, globalState);
        let _lhs313 = globalState;
        let _lhs314 = globalState;
        _2261_v4 = _rhs400;
        _lhs313.f9 = _rhs401;
        _lhs314.f9 = _rhs402;
        let _index363 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
        let _rhs403 = (p0).multipliedBy(p0);
        let _rhs404 = _module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Set.fromElements((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))])).length));
        let _rhs405 = ((!((_this).f13) || (!((_this).f14))) ? ((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]) : ((_this).f14));
        let _rhs406 = p0;
        let _rhs407 = true;
        let _lhs315 = globalState;
        let _lhs316 = _2271_v13;
        let _lhs317 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
        _lhs315.f9 = _rhs403;
        r1 = _rhs404;
        r3 = _rhs405;
        r1 = _rhs406;
        _lhs316[_lhs317] = _rhs407;
        let _2278_v18;
        _2278_v18 = _dafny.MultiSet.fromElements(_2271_v13, _2271_v13);
        let _2279_v19;
        let _nw381 = new _module.C1();
        _nw381.__ctor(((_this).f14) && ((_this).f13), !(_2278_v18).equals(_dafny.MultiSet.fromElements(_2271_v13)), _module.__default.fm40(_2260_v3, globalState), (_this).f13);
        _2279_v19 = _nw381;
        let _index364 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((p1).length));
        (p1)[_index364] = p0;
        let _2280_v20;
        _2280_v20 = _dafny.Seq.of((_this).f14, (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
        let _2281_v21;
        _2281_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2280_v20,_module.__default.fm1((_this).f14, _2260_v3, globalState));
        let _index365 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((p1).length));
        let _rhs408 = (((((_2279_v19).f19) ? ((_this).f14) : ((_this).f13))) ? (_dafny.Seq.IsPrefixOf(_2258_v0, _2258_v0)) : ((_2280_v20)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_2280_v20).length))]));
        let _rhs409 = (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))];
        let _rhs410 = (p0).minus(new BigNumber((_2281_v21).length));
        let _rhs411 = ((_2268_v11).Intersect(_dafny.Set.fromElements(p0))).IsSubsetOf(_2268_v11);
        let _rhs412 = _2260_v3;
        let _lhs318 = globalState;
        let _lhs319 = p1;
        let _lhs320 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((p1).length));
        _lhs318.f8 = _rhs408;
        r3 = _rhs409;
        _lhs319[_lhs320] = _rhs410;
        r3 = _rhs411;
        _2260_v3 = _rhs412;
      } else if (_source26.is_DC7) {
        let _2282___mcc_h0 = (_source26).cf11;
        let _2283___mcc_h1 = (_source26).cf12;
        let _2284___mcc_h2 = (_source26).cf13;
        let _2285___mcc_h3 = (_source26).cf14;
        let _2286_cf14 = _2285___mcc_h3;
        let _2287_cf13 = _2284___mcc_h2;
        let _2288_cf12 = _2283___mcc_h1;
        let _2289_cf11 = _2282___mcc_h0;
        let _2290_v22;
        _2290_v22 = _dafny.Set.fromElements(_2271_v13);
        let _2291_v23;
        _2291_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2290_v22);
        _2291_v23 = (_2291_v23).update(p1, (_2290_v22).Intersect(_2290_v22));
        if ((_this).f13) {
          (globalState).f9 = ((_2288_cf12).minus(p0)).multipliedBy(new BigNumber((_2258_v0).length));
          let _2292_v24;
          _2292_v24 = _dafny.Seq.of(false);
          _2258_v0 = _dafny.Seq.Concat(_2286_cf14, (_module.__default.fm60(new BigNumber((_2292_v24).length), !(_2289_cf11), globalState)).dtor_cf34);
          let _2293_v25;
          let _init74 = ((_2294_v3) => function (_2295_i1) {
            return _2294_v3;
          })(_2260_v3);
          let _nw382 = Array((new BigNumber(29)).toNumber());
          for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw382.length); _i0_74++) {
            _nw382[_i0_74] = _init74(new BigNumber(_i0_74));
          }
          _2293_v25 = _nw382;
          let _index366 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_2293_v25).length));
          (_2293_v25)[_index366] = _2260_v3;
          let _index367 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_2293_v25).length));
          (_2293_v25)[_index367] = _2260_v3;
          _2292_v24 = _2292_v24;
          let _index368 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_2293_v25).length));
          (_2293_v25)[_index368] = _module.__default.fm3((_dafny.ZERO).minus(_2288_cf12), _2289_cf11, globalState);
        } else {
          let _2296_v26;
          let _nw383 = new _module.C0();
          _nw383.__ctor(_2289_cf11);
          _2296_v26 = _nw383;
          let _2297_v27;
          let _nw384 = new _module.C5();
          _nw384.__ctor((_this).f11);
          _2297_v27 = _nw384;
          let _index369 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index369] = ((p0).isLessThanOrEqualTo(_2288_cf12)) && ((_dafny.Set.fromElements((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], (_2296_v26).f17, _2289_cf11, _module.__default.fm1((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], _2260_v3, globalState), true)).IsSubsetOf(_dafny.Set.fromElements((_this).f14)));
          let _2298_v28;
          _2298_v28 = _dafny.Map.Empty.slice().updateUnsafe((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))],new BigNumber((_2258_v0).length));
          _2298_v28 = (_2298_v28).update((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], p0);
          let _index370 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index370] = (p0).isLessThan((p0).multipliedBy(new BigNumber(619)));
        }
        if (_module.__default.fm1(false, new _dafny.CodePoint('w'.codePointAt(0)), globalState)) {
          let _2299_v29;
          _2299_v29 = _dafny.Seq.of((((_2269_v12).contains((_dafny.ZERO).minus(p0))) ? ((_2269_v12).get((_dafny.ZERO).minus(p0))) : ((_this).f14)), false);
          _2299_v29 = _dafny.Seq.of((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], false);
          r0 = _2260_v3;
          let _2300_v30;
          let _nw385 = new _module.C6();
          _nw385.__ctor((p0).multipliedBy(p0), _dafny.Seq.Concat(_2258_v0, _2258_v0));
          _2300_v30 = _nw385;
          _2299_v29 = _2299_v29;
          (globalState).f9 = _2288_cf12;
        } else {
          _2259_v1 = _2259_v1;
          _module.__default.m0(globalState);
          let _2301_v31;
          _2301_v31 = _dafny.Seq.of(_2258_v0);
          let _index371 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index371] = _dafny.Seq.contains((_2301_v31)[_module.__default.safeIndex(p0, new BigNumber((_2301_v31).length))], _2260_v3);
          let _2302_v32;
          let _nw386 = new _module.C6();
          _nw386.__ctor(p0, _2258_v0);
          _2302_v32 = _nw386;
          let _2303_v33;
          _2303_v33 = _dafny.Seq.of(_2289_cf11);
          let _2304_v34;
          _2304_v34 = _module.D2.create_DC7(_2287_cf13, new BigNumber((_dafny.Seq.update((_2301_v31)[_module.__default.safeIndex(new BigNumber((_2303_v33).length), new BigNumber((_2301_v31).length))], _module.__default.safeIndex((_2302_v32).f15, new BigNumber(((_2301_v31)[_module.__default.safeIndex(new BigNumber((_2303_v33).length), new BigNumber((_2301_v31).length))]).length)), _2260_v3)).length), (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], _2302_v32.f16);
          let _2305_v35;
          _2305_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2266_v9,!((_this).f13));
          let _rhs413 = _module.__default.fm26((_2304_v34).dtor_cf13, (_dafny.ZERO).minus(new BigNumber((_2305_v35).length)), !_dafny.areEqual(_2303_v33, _2303_v33), globalState);
          let _rhs414 = _module.__default.safeModuloInt(_2288_cf12, _module.__default.safeModuloInt(_2288_cf12, new BigNumber((_dafny.Seq.UnicodeFromString("ilvvkfxmt")).length)));
          let _rhs415 = _2302_v32;
          let _rhs416 = ((_2287_cf13) ? ((_this).f14) : (true));
          let _lhs321 = globalState;
          let _lhs322 = globalState;
          _2264_v7 = _rhs413;
          _lhs321.f9 = _rhs414;
          _2302_v32 = _rhs415;
          _lhs322.f8 = _rhs416;
          let _index372 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((p1).length));
          (p1)[_index372] = (p0).minus(p0);
          let _index373 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((p1).length));
          let _index374 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          let _rhs417 = (new BigNumber(512)).multipliedBy(_module.__default.safeModuloInt((_2302_v32).f15, new BigNumber((_dafny.Seq.UnicodeFromString("aiwr")).length)));
          let _rhs418 = !((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(326), p0))).isEqualTo(_2288_cf12);
          let _lhs323 = p1;
          let _lhs324 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((p1).length));
          let _lhs325 = _2271_v13;
          let _lhs326 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          _lhs323[_lhs324] = _rhs417;
          _lhs325[_lhs326] = _rhs418;
        }
        let _index375 = _module.__default.safeIndex(new BigNumber(845), new BigNumber(((_this).f11).length));
        ((_this).f11)[_index375] = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
        let _2306_v36;
        _2306_v36 = _dafny.Seq.of((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], (_this).f13, (_this).f13, (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], (_this).f13);
        let _index376 = _module.__default.safeIndex(new BigNumber(845), new BigNumber(((_this).f11).length));
        let _rhs419 = _dafny.Seq.contains(_2258_v0, _2260_v3);
        let _rhs420 = (_2262_v5).IsProperSubsetOf(_2262_v5);
        let _rhs421 = _2269_v12;
        let _rhs422 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        let _rhs423 = _2306_v36;
        let _lhs327 = (_this).f11;
        let _lhs328 = _module.__default.safeIndex(new BigNumber(845), new BigNumber(((_this).f11).length));
        let _lhs329 = globalState;
        r3 = _rhs419;
        _2287_cf13 = _rhs420;
        _lhs327[_lhs328] = _rhs421;
        _lhs329.f9 = _rhs422;
        _2306_v36 = _rhs423;
      } else if (_source26.is_DC8) {
        let _2307___mcc_h4 = (_source26).cf15;
        let _2308___mcc_h5 = (_source26).cf16;
        let _2309_cf16 = _2308___mcc_h5;
        let _2310_cf15 = _2307___mcc_h4;
        (globalState).f9 = _2310_cf15;
        let _2311_v37;
        let _nw387 = new _module.C1();
        _nw387.__ctor((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], (_this).f14, _this.f12, (_this).f14);
        _2311_v37 = _nw387;
        _2310_cf15 = ((_2310_cf15).plus((_2311_v37).fm9(_module.D1.create_DC4(_2260_v3), function () {
          let _coll75 = new _dafny.Set();
          for (const _compr_75 of _dafny.IntegerRange(new BigNumber(653), new BigNumber(65))) {
            let _2312_v38 = _compr_75;
            if (((new BigNumber(653)).isLessThanOrEqualTo(_2312_v38)) && ((_2312_v38).isLessThan(new BigNumber(65)))) {
              _coll75.add((_2312_v38).plus(p0));
            }
          }
          return _coll75;
        }(), globalState))).plus((((_2267_v10).contains((_dafny.ZERO).minus(p0))) ? ((_2267_v10).get((_dafny.ZERO).minus(p0))) : (_2310_cf15)));
        (globalState).f8 = _module.__default.fm1((_this).f13, _2309_cf16, globalState);
      } else {
        let _2313___mcc_h6 = (_source26).cf10;
        let _2314_cf10 = _2313___mcc_h6;
        let _rhs424 = _2271_v13;
        let _rhs425 = (_this).f14;
        let _rhs426 = _module.__default.safeModuloInt(p0, new BigNumber((_2262_v5).length));
        _2271_v13 = _rhs424;
        r3 = _rhs425;
        r1 = _rhs426;
        (globalState).f9 = (_dafny.ZERO).minus(p0);
        let _2315_v39;
        _2315_v39 = _module.D2.create_DC7((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], p0, (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], _2314_cf10);
        if (!((_2315_v39).dtor_cf13)) {
          let _2316_v40;
          _2316_v40 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          _2316_v40 = (_2316_v40).update((((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]) ? (new BigNumber(615)) : (new BigNumber((_module.__default.fm53(globalState)).length))), p1);
          let _2317_v41;
          _2317_v41 = _module.D4.create_DC11(p0, p0, (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
          let _2318_v42;
          let _nw388 = new _module.C0();
          _nw388.__ctor((_2317_v41).dtor_cf21);
          _2318_v42 = _nw388;
          let _2319_v43;
          _2319_v43 = _dafny.Map.Empty.slice().updateUnsafe((_2268_v11).IsSubsetOf(_2268_v11),_2274_v16);
          _2319_v43 = (_2319_v43).update((_2318_v42).f17, _2274_v16);
          let _2320_v44;
          let _nw389 = new _module.C4();
          _nw389.__ctor((_this).f11);
          _2320_v44 = _nw389;
          _2320_v44 = _2320_v44;
          let _2321_v45;
          let _nw390 = new _module.C5();
          _nw390.__ctor((_this).f11);
          _2321_v45 = _nw390;
        } else {
          let _index377 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((p1).length));
          (p1)[_index377] = p0;
          let _2322_v47;
          _2322_v47 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of (_2264_v7).Elements) {
              let _2323_v46 = _compr_76;
              if (_dafny.Seq.contains(_2264_v7, _2323_v46)) {
                _coll76.push([(_2323_v46).multipliedBy(p0),new _dafny.CodePoint('i'.codePointAt(0))]);
              }
            }
            return _coll76;
          }()).length));
          let _2324_v48;
          _2324_v48 = _dafny.MultiSet.fromElements((_2322_v47).Merge(_2322_v47), _2322_v47, _2322_v47, (_module.__default.fm50(_2260_v3, (_this).f14, (_this).f13, globalState)).Merge(_2322_v47), _2322_v47);
          let _2325_v49;
          _2325_v49 = _module.D2.create_DC8(p0, _2260_v3);
          let _2326_v50;
          _2326_v50 = _dafny.MultiSet.fromElements(_module.__default.fm37((_this).f14, new BigNumber((_2258_v0).length), p0, p0, globalState), _2268_v11, _2268_v11);
          let _index378 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((p1).length));
          let _rhs427 = (((_2324_v48).contains(_2322_v47)) ? ((_2324_v48).get(_2322_v47)) : (p0));
          let _rhs428 = (_2264_v7)[_module.__default.safeIndex(new BigNumber((_2258_v0).length), new BigNumber((_2264_v7).length))];
          let _rhs429 = _module.__default.fm1((_this).f13, (_2325_v49).dtor_cf16, globalState);
          let _rhs430 = _module.__default.fm0(_module.__default.fm1((_this).f13, _2260_v3, globalState), globalState);
          let _rhs431 = (((_2326_v50).contains(_2268_v11)) ? ((_2326_v50).get(_2268_v11)) : (p0));
          let _lhs330 = globalState;
          let _lhs331 = p1;
          let _lhs332 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((p1).length));
          let _lhs333 = globalState;
          let _lhs334 = globalState;
          let _lhs335 = globalState;
          _lhs330.f9 = _rhs427;
          _lhs331[_lhs332] = _rhs428;
          _lhs333.f8 = _rhs429;
          _lhs334.f9 = _rhs430;
          _lhs335.f9 = _rhs431;
          let _2327_v51;
          let _nw391 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2327_v51 = _nw391;
          let _index379 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_2327_v51).length));
          (_2327_v51)[_index379] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(570)), ((_2328_v3) => function (_2329_i2) {
            return _2328_v3;
          })(_2260_v3)), _module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(42), new BigNumber((p1).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(570)), ((_2330_v3) => function (_2331_i2) {
            return _2330_v3;
          })(_2260_v3))).length)), new _dafny.CodePoint('a'.codePointAt(0)));
          let _index380 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_2327_v51).length));
          (_2327_v51)[_index380] = _dafny.Seq.Concat(_2258_v0, _2314_cf10);
          let _2332_v52;
          let _nw392 = Array((new BigNumber(23)).toNumber());
          _2332_v52 = _nw392;
          let _2333_v53;
          let _nw393 = new _module.C1();
          _nw393.__ctor((_this).f14, false, _this.f12, !(false));
          _2333_v53 = _nw393;
          let _index381 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2332_v52).length));
          (_2332_v52)[_index381] = _2333_v53;
          let _2334_v54;
          let _nw394 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _2334_v54 = _nw394;
          let _2335_v55;
          _2335_v55 = _dafny.Seq.of((_2333_v53).f19);
          let _2336_v56;
          _2336_v56 = _dafny.Seq.of((_dafny.MultiSet.FromArray(_2335_v55)).Intersect(_dafny.MultiSet.FromArray(_2335_v55)), _2261_v4);
          let _index382 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((p1).length));
          let _index383 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2332_v52).length));
          let _index384 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          let _rhs432 = new BigNumber((_2336_v56).length);
          let _rhs433 = _2333_v53;
          let _rhs434 = _2334_v54;
          let _rhs435 = (_this).f13;
          let _rhs436 = (_module.__default.fm1((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], _2260_v3, globalState)) || (_dafny.areEqual(_2314_cf10, _2258_v0));
          let _lhs336 = p1;
          let _lhs337 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((p1).length));
          let _lhs338 = _2332_v52;
          let _lhs339 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2332_v52).length));
          let _lhs340 = _2271_v13;
          let _lhs341 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          let _lhs342 = globalState;
          _lhs336[_lhs337] = _rhs432;
          _lhs338[_lhs339] = _rhs433;
          _2334_v54 = _rhs434;
          _lhs340[_lhs341] = _rhs435;
          _lhs342.f8 = _rhs436;
          let _2337_v57;
          let _init75 = ((_2338_v39) => function (_2339_i3) {
            return _2338_v39;
          })(_2315_v39);
          let _nw395 = Array((new BigNumber(23)).toNumber());
          for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw395.length); _i0_75++) {
            _nw395[_i0_75] = _init75(new BigNumber(_i0_75));
          }
          _2337_v57 = _nw395;
          let _2340_v58;
          _2340_v58 = _module.D11.create_DC25(_2337_v57);
          let _2341_v59;
          _2341_v59 = _module.D11.create_DC27(_2340_v58);
          let _2342_v60;
          _2342_v60 = _module.D11.create_DC27(_2341_v59);
          let _pat_let_tv40 = _2341_v59;
          _2342_v60 = function (_pat_let37_0) {
            return function (_2343_dt__update__tmp_h0) {
              return function (_pat_let38_0) {
                return function (_2344_dt__update_hcf48_h0) {
                  return _module.D11.create_DC27(_2344_dt__update_hcf48_h0);
                }(_pat_let38_0);
              }(_module.D11.create_DC27(_module.D11.create_DC27(_pat_let_tv40)));
            }(_pat_let37_0);
          }(_2342_v60);
          let _2345_v61;
          let _nw396 = new _module.C0();
          _nw396.__ctor(!((_2315_v39).dtor_cf13));
          _2345_v61 = _nw396;
        }
        r1 = p0;
      }
      let _2346_v62;
      _2346_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("erufxrt"));
      let _2347_v63;
      _2347_v63 = _module.D9.create_DC22(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(735)), ((_2348_v7, _2349_v0) => function (_2350_i4) {
  return (_2348_v7)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_2349_v0)).cardinality()), new BigNumber((_2348_v7).length))];
})(_2264_v7, _2258_v0))).length), _2346_v62);
      let _2351_v64;
      _2351_v64 = _dafny.Seq.of((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
      let _2352_v65;
      _2352_v65 = _dafny.Seq.of(_2347_v63, _module.D9.create_DC22(new BigNumber((_2351_v64).length), _2346_v62), _module.D9.create_DC22(p0, _2346_v62));
      let _pat_let_tv41 = _2260_v3;
      let _pat_let_tv42 = _2260_v3;
      let _rhs437 = ((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]) || ((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
      let _rhs438 = _2352_v65;
      let _rhs439 = function (_source27) {
        if (_source27.is_DC22) {
          let _2353___mcc_h7 = (_source27).cf40;
          let _2354___mcc_h8 = (_source27).cf41;
          let _2355_cf41 = _2354___mcc_h8;
          let _2356_cf40 = _2353___mcc_h7;
          return _pat_let_tv41;
        } else {
          let _2357___mcc_h9 = (_source27).cf39;
          let _2358_cf39 = _2357___mcc_h9;
          return _pat_let_tv42;
        }
      }(_2347_v63);
      let _rhs440 = _dafny.Seq.Concat(_2258_v0, _2258_v0);
      let _rhs441 = p0;
      let _lhs343 = globalState;
      r3 = _rhs437;
      _2352_v65 = _rhs438;
      r0 = _rhs439;
      _2258_v0 = _rhs440;
      _lhs343.f9 = _rhs441;
      if ((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]) {
        (globalState).f9 = p0;
        let _index385 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
        (_2271_v13)[_index385] = (_this).f14;
        (globalState).f9 = (p0).plus(new BigNumber((_2351_v64).length));
        if ((_this).f13) {
          let _index386 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((p1).length));
          (p1)[_index386] = _module.__default.safeModuloInt(p0, p0);
          let _2359_v66;
          _2359_v66 = _dafny.Set.fromElements(_2258_v0);
          let _index387 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((p1).length));
          (p1)[_index387] = new BigNumber((_2359_v66).length);
          let _2360_v67;
          let _nw397 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2360_v67 = _nw397;
          let _2361_v68;
          _2361_v68 = _dafny.Seq.of(_2262_v5, _2262_v5);
          let _2362_v69;
          _2362_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2360_v67,(p1)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p1).length))])).length),new BigNumber((_2361_v68).length));
          let _2363_v70;
          _2363_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2266_v9);
          let _2364_v71;
          _2364_v71 = _dafny.Seq.of(_2266_v9);
          _2362_v69 = (_2362_v69).update((new BigNumber((_2264_v7).length)).plus(new BigNumber(((((_2363_v70).contains(new BigNumber((_2261_v4).cardinality()))) ? ((_2363_v70).get(new BigNumber((_2261_v4).cardinality()))) : ((_2364_v71)[_module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p1).length))], new BigNumber((_2364_v71).length))]))).cardinality())), (p1)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p1).length))]);
          let _index388 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index388] = (_this).f13;
          (globalState).f8 = _module.__default.fm1((_2268_v11).IsProperSubsetOf(_2268_v11), _2260_v3, globalState);
          _2264_v7 = _2264_v7;
        } else {
          let _rhs442 = true;
          let _rhs443 = (_2351_v64)[_module.__default.safeIndex(p0, new BigNumber((_2351_v64).length))];
          let _rhs444 = ((_dafny.MultiSet.fromElements(true)).Intersect(_2261_v4)).contains(true);
          let _rhs445 = (((_this).f14) ? ((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]) : ((_this).f13));
          let _lhs344 = globalState;
          let _lhs345 = globalState;
          let _lhs346 = globalState;
          _lhs344.f8 = _rhs442;
          _lhs345.f8 = _rhs443;
          r3 = _rhs444;
          _lhs346.f8 = _rhs445;
          let _index389 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((p1).length));
          (p1)[_index389] = new BigNumber(((((_this).f14) ? (_2351_v64) : (_2351_v64))).length);
          let _2365_v72;
          _2365_v72 = _module.D2.create_DC7((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], (((_2267_v10).contains(new BigNumber((_dafny.Seq.update(_2258_v0, _module.__default.safeIndex(p0, new BigNumber((_2258_v0).length)), _2260_v3)).length))) ? ((_2267_v10).get(new BigNumber((_dafny.Seq.update(_2258_v0, _module.__default.safeIndex(p0, new BigNumber((_2258_v0).length)), _2260_v3)).length))) : (p0)), (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))], _2258_v0);
          let _index390 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((p1).length));
          (p1)[_index390] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((_2365_v72).dtor_cf14).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(p0))));
          let _index391 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index391] = (_this).f13;
          let _2366_v73;
          _2366_v73 = _module.D7.create_DC17(_2258_v0, _2271_v13, (_this).f13, (p1)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((p1).length))]);
          _2366_v73 = _2366_v73;
          (globalState).f8 = !((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
        }
        (globalState).f9 = (_dafny.ZERO).minus(new BigNumber(((_module.__default.fm53(globalState)).Intersect(_2262_v5)).length));
      } else {
        if ((_this).f13) {
          let _index392 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index392] = (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))];
          _2261_v4 = (_2261_v4).Difference(_module.__default.fm46(globalState));
          _2266_v9 = _2267_v10;
          let _index393 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((p1).length));
          (p1)[_index393] = p0;
          let _index394 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((p1).length));
          (p1)[_index394] = new BigNumber(-95);
          let _2367_v74;
          let _nw398 = new _module.C6();
          _nw398.__ctor((_2264_v7)[_module.__default.safeIndex(p0, new BigNumber((_2264_v7).length))], _dafny.Seq.update(_dafny.Seq.Concat(_2258_v0, _2258_v0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_2258_v0, _2258_v0)).length)), _2260_v3));
          _2367_v74 = _nw398;
        } else {
          let _2368_v75;
          let _nw399 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2368_v75 = _nw399;
          _2368_v75 = ((false) ? (_2368_v75) : (_2368_v75));
          (globalState).f9 = new BigNumber(-735);
          (globalState).f9 = (_this).fm11((_this).f13, false, new BigNumber(((_2261_v4).Intersect(_2261_v4)).cardinality()), (_this).f14, globalState);
          let _index395 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index395] = (_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))];
          let _2369_v76;
          let _init76 = ((_2370_v0) => function (_2371_i5) {
            return (_dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_2370_v0).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber((_2370_v0).length)));
          })(_2258_v0);
          let _nw400 = Array((new BigNumber(28)).toNumber());
          for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw400.length); _i0_76++) {
            _nw400[_i0_76] = _init76(new BigNumber(_i0_76));
          }
          _2369_v76 = _nw400;
          let _2372_v77;
          _2372_v77 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f14) ? (_2351_v64) : (_2351_v64)),(_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))]);
          let _index396 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          let _rhs446 = _2369_v76;
          let _rhs447 = _2271_v13;
          let _rhs448 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2351_v64, _2351_v64),(_this).f13);
          let _rhs449 = (_this).fm11((_this).f14, true, (_module.__default.fm0((_this).f13, globalState)).plus(p0), (_this).f13, globalState);
          let _rhs450 = (_this).f14;
          let _lhs347 = globalState;
          let _lhs348 = _2271_v13;
          let _lhs349 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          _2369_v76 = _rhs446;
          _2271_v13 = _rhs447;
          _2372_v77 = _rhs448;
          _lhs347.f9 = _rhs449;
          _lhs348[_lhs349] = _rhs450;
        }
        let _2373_v78;
        let _nw401 = new _module.C0();
        _nw401.__ctor((_this).f14);
        _2373_v78 = _nw401;
        let _index397 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length));
        (p1)[_index397] = p0;
        let _index398 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length));
        (p1)[_index398] = p0;
        let _2374_v79;
        _2374_v79 = _module.D8.create_DC20(_2263_v6, (_this).f14, (_this).f14);
        let _source28 = _2374_v79;
        if (_source28.is_DC19) {
          let _2375___mcc_h10 = (_source28).cf34;
          let _2376___mcc_h11 = (_source28).cf35;
          let _2377_cf35 = _2376___mcc_h11;
          let _2378_cf34 = _2375___mcc_h10;
          let _index399 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index399] = (_this).f13;
          let _2379_v80;
          let _nw402 = new _module.C4();
          _nw402.__ctor((_this).f11);
          _2379_v80 = _nw402;
          let _index400 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length));
          (p1)[_index400] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2264_v7, _dafny.Seq.Concat(_2264_v7, _2264_v7)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_2264_v7, _dafny.Seq.Concat(_2264_v7, _2264_v7))).length)), p0)).length);
          (globalState).f9 = (_2347_v63).dtor_cf40;
        } else if (_source28.is_DC20) {
          let _2380___mcc_h12 = (_source28).cf36;
          let _2381___mcc_h13 = (_source28).cf37;
          let _2382___mcc_h14 = (_source28).cf38;
          let _2383_cf38 = _2382___mcc_h14;
          let _2384_cf37 = _2381___mcc_h13;
          let _2385_cf36 = _2380___mcc_h12;
          (globalState).f9 = (p1)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length))];
          (globalState).f9 = new BigNumber(547);
          let _2386_v81;
          _2386_v81 = _dafny.Set.fromElements(_2271_v13);
          let _2387_v82;
          let _nw403 = Array((new BigNumber(3)).toNumber());
          _nw403[(_dafny.ZERO).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length))];
          _nw403[(_dafny.ONE).toNumber()] = new BigNumber((_2386_v81).length);
          _nw403[(new BigNumber(2)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length))];
          _2387_v82 = _nw403;
          let _2388_v83;
          _2388_v83 = _dafny.Seq.of(_2387_v82);
          _2388_v83 = _2388_v83;
          let _index401 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length));
          (p1)[_index401] = _module.__default.safeModuloInt(p0, (new BigNumber((_2261_v4).cardinality())).plus((p1)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length))]));
        } else {
          let _2389___mcc_h15 = (_source28).cf33;
          let _2390_cf33 = _2389___mcc_h15;
          let _2391_v84;
          _2391_v84 = _dafny.Map.Empty.slice().updateUnsafe(_2264_v7,true);
          let _2392_v85;
          _2392_v85 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((p1).length))]),new BigNumber((_module.__default.fm46(globalState)).cardinality()));
          let _rhs451 = (_this).fm9(_2259_v1, _2390_cf33, globalState);
          let _rhs452 = !(_2392_v85).contains((new BigNumber(((_2391_v84).update(_2264_v7, (_this).f13)).length)).multipliedBy(p0));
          let _lhs350 = globalState;
          let _lhs351 = globalState;
          _lhs350.f9 = _rhs451;
          _lhs351.f8 = _rhs452;
          (globalState).f8 = (_module.__default.fm0((_this).f13, globalState)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_2351_v64, _2351_v64)).length));
          _2260_v3 = _2260_v3;
          let _index402 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length));
          (_2271_v13)[_index402] = (_this).f14;
        }
        r3 = (_this).f13;
      }
      r0 = _2260_v3;
      let _2393_v86;
      _2393_v86 = _module.D5.create_DC14((_this).f13);
      let _pat_let_tv43 = _2258_v0;
      let _pat_let_tv44 = p0;
      r1 = function (_source29) {
        if (_source29.is_DC14) {
          let _2394___mcc_h16 = (_source29).cf26;
          let _2395_cf26 = _2394___mcc_h16;
          return new BigNumber((_pat_let_tv43).length);
        } else {
          let _2396___mcc_h17 = (_source29).cf25;
          let _2397_cf25 = _2396___mcc_h17;
          return (_dafny.ZERO).minus(_pat_let_tv44);
        }
      }(function (_pat_let39_0) {
        return function (_2398_dt__update__tmp_h1) {
          return function (_pat_let40_0) {
            return function (_2399_dt__update_hcf26_h0) {
              return _module.D5.create_DC14(_2399_dt__update_hcf26_h0);
            }(_pat_let40_0);
          }((_this).f13);
        }(_pat_let39_0);
      }(_2393_v86));
      let _2400_v87;
      _2400_v87 = _module.D4.create_DC11(p0, p0, (_this).f14);
      let _2401_v88;
      _2401_v88 = _module.D2.create_DC8(p0, _2260_v3);
      r2 = (((_2400_v87).dtor_cf21) ? (_2401_v88) : (_2401_v88));
      let _2402_v89;
      _2402_v89 = _dafny.Map.Empty.slice().updateUnsafe((_2271_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_2271_v13).length))],p0);
      r3 = (new BigNumber(-251)).isEqualTo((((_2402_v89).contains(!((_this).f13))) ? ((_2402_v89).get(!((_this).f13))) : (p0)));
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = false;
      let r2 = false;
      if (!(!((_this).f13)) || (((_this).f13) && ((_this).f13))) {
        _module.__default.m0(globalState);
        let _2403_v0;
        let _nw404 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _2403_v0 = _nw404;
        let _index403 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length));
        (_2403_v0)[_index403] = new BigNumber(740);
        let _2404_v1;
        _2404_v1 = new BigNumber(686);
        let _2405_v2;
        _2405_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2404_v1,_dafny.Seq.of((_this).f13, (_this).f13, (_this).f13));
        let _2406_v3;
        _2406_v3 = _dafny.Seq.UnicodeFromString("y");
        let _2407_v4;
        _2407_v4 = _dafny.MultiSet.fromElements(_2406_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), function (_2408_i0) {
          return (_module.D13.create_DC30(new _dafny.CodePoint('q'.codePointAt(0)))).dtor_cf51;
        }));
        let _2409_v5;
        _2409_v5 = _dafny.Seq.of((new BigNumber(462)).minus((_dafny.ZERO).minus(_2404_v1)));
        let _2410_v6;
        _2410_v6 = _dafny.Set.fromElements((_this).f13, (_this).f14, (_this).f13, (_this).f14);
        let _index404 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length));
        let _rhs453 = (_2409_v5)[_module.__default.safeIndex(_module.__default.safeModuloInt(_2404_v1, _2404_v1), new BigNumber((_2409_v5).length))];
        let _rhs454 = (_this).f14;
        let _rhs455 = _module.__default.fm61(globalState);
        let _rhs456 = ((_2410_v6).Intersect(_2410_v6)).IsSubsetOf(_2410_v6);
        let _rhs457 = _2407_v4;
        let _lhs352 = _2403_v0;
        let _lhs353 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length));
        _lhs352[_lhs353] = _rhs453;
        r1 = _rhs454;
        _2405_v2 = _rhs455;
        r2 = _rhs456;
        _2407_v4 = _rhs457;
        let _2411_v7;
        let _init77 = function (_2412_i1) {
          return !((_this).f13);
        };
        let _nw405 = Array((new BigNumber(25)).toNumber());
        for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw405.length); _i0_77++) {
          _nw405[_i0_77] = _init77(new BigNumber(_i0_77));
        }
        _2411_v7 = _nw405;
        let _2413_v8;
        _2413_v8 = _module.D7.create_DC17(_2406_v3, _2411_v7, (_this).f13, _2404_v1);
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((_2413_v8).dtor_cf29, _dafny.Seq.UnicodeFromString("jcasst")), _2406_v3)) {
          (globalState).f9 = _2404_v1;
          r1 = (_this).f14;
          let _index405 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length));
          (_2403_v0)[_index405] = _module.__default.fm0((_this).fm10(globalState), globalState);
          let _2414_v9;
          _2414_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f11);
          let _2415_v10;
          let _nw406 = new _module.C5();
          _nw406.__ctor((((_2414_v9).contains((_this).f14)) ? ((_2414_v9).get((_this).f14)) : ((_this).f11)));
          _2415_v10 = _nw406;
          let _2416_v11;
          _2416_v11 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_2404_v1));
          let _index406 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length));
          (_2403_v0)[_index406] = (((_2416_v11).contains(_dafny.Seq.of((_2403_v0)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length))], _2404_v1))) ? ((_2416_v11).get(_dafny.Seq.of((_2403_v0)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length))], _2404_v1))) : (_module.__default.fm0((_this).f14, globalState)));
        } else {
          let _2417_v12;
          _2417_v12 = new _dafny.CodePoint('s'.codePointAt(0));
          r0 = _2417_v12;
          let _2418_v13;
          _2418_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_2409_v5, _module.__default.safeIndex((_2403_v0)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length))], new BigNumber((_2409_v5).length)), (_2403_v0)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length))]),_2404_v1);
          _2418_v13 = (_2418_v13).update(_2409_v5, (_2403_v0)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_2403_v0).length))]);
          r0 = _2417_v12;
          let _rhs458 = (((_this).f14) ? ((_this).f13) : (false));
          let _rhs459 = (_this).f13;
          let _rhs460 = (_this).f13;
          let _rhs461 = _module.__default.safeModuloInt(new BigNumber(464), _2404_v1);
          let _lhs354 = globalState;
          r1 = _rhs458;
          r2 = _rhs459;
          r2 = _rhs460;
          _lhs354.f9 = _rhs461;
          let _2419_v14;
          _2419_v14 = _dafny.MultiSet.fromElements((_this).f13, (_this).f14, (_this).f13, (_this).f13, (_this).f14);
          let _2420_v15;
          _2420_v15 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_2419_v14).cardinality())));
          let _2421_v16;
          let _nw407 = new _module.C3();
          _nw407.__ctor((_this).f13, (_2420_v15).IsProperSubsetOf(_2420_v15), (_this).f11);
          _2421_v16 = _nw407;
        }
        let _2422_v17;
        _2422_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,!(!((_this).f13)));
        let _2423_v18;
        _2423_v18 = new _dafny.CodePoint('t'.codePointAt(0));
        let _2424_v19;
        let _nw408 = new _module.C1();
        _nw408.__ctor((((_2422_v17).contains((_this).f13)) ? ((_2422_v17).get((_this).f13)) : ((_this).f14)), (_this).f14, _module.__default.fm40(_2423_v18, globalState), (_this).f14);
        _2424_v19 = _nw408;
        _2424_v19 = _2424_v19;
        let _2425_v20;
        _2425_v20 = _dafny.MultiSet.fromElements(false, (_this).f14, (_2424_v19).f13, (_2424_v19).f13, (_this).f13);
        let _2426_v21;
        _2426_v21 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yphid"), _2406_v3);
        let _2427_v22;
        _2427_v22 = _module.D0.create_DC2(_2425_v20, (_this).f14, new BigNumber(((_2426_v21)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((_2426_v21).length))]).length), (_this).f13);
        r1 = (_2427_v22).dtor_cf5;
      } else {
        let _2428_v23;
        let _nw409 = new _module.C5();
        _nw409.__ctor((_this).f11);
        _2428_v23 = _nw409;
        let _2429_v24;
        _2429_v24 = new BigNumber(938);
        let _2430_v25;
        _2430_v25 = _dafny.Seq.UnicodeFromString("oleu");
        let _2431_v26;
        _2431_v26 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2432_v27;
        _2432_v27 = _module.D2.create_DC7((new BigNumber((_dafny.MultiSet.fromElements(_2428_v23)).cardinality())).isLessThanOrEqualTo(_2429_v24), (_2429_v24).plus(new BigNumber((_dafny.Seq.of((_this).f13)).length)), !((_this).f14), _dafny.Seq.update(_2430_v25, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2430_v25).length)), new BigNumber((_2430_v25).length)), _2431_v26));
        let _source30 = _2432_v27;
        if (_source30.is_DC6) {
          (globalState).f9 = _2429_v24;
          let _2433_v28;
          let _nw410 = Array((new BigNumber(16)).toNumber()).fill(false);
          _2433_v28 = _nw410;
          let _2434_v29;
          _2434_v29 = _dafny.Seq.of(_2429_v24);
          let _2435_v30;
          _2435_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2433_v28,_dafny.Seq.update(_2434_v29, _module.__default.safeIndex(new BigNumber(753), new BigNumber((_2434_v29).length)), new BigNumber((_2434_v29).length)));
          let _2436_v31;
          _2436_v31 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13);
          let _2437_v32;
          _2437_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2429_v24);
          let _2438_v33;
          _2438_v33 = _dafny.Seq.of(_2430_v25);
          let _2439_v34;
          let _nw411 = Array((new BigNumber(27)).toNumber());
          _nw411[(_dafny.ZERO).toNumber()] = _2429_v24;
          _nw411[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_2429_v24, _2429_v24);
          _nw411[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_2429_v24, _2429_v24);
          _nw411[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_2434_v29)[_module.__default.safeIndex(_2429_v24, new BigNumber((_2434_v29).length))], _2429_v24);
          _nw411[(new BigNumber(4)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(5)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(6)).toNumber()] = (new BigNumber((_dafny.Seq.of(_2429_v24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2429_v24)).length), new BigNumber((_2436_v31).cardinality()), new BigNumber(359), _2429_v24)).length)).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_2434_v29)).cardinality()));
          _nw411[(new BigNumber(7)).toNumber()] = new BigNumber((_2437_v32).length);
          _nw411[(new BigNumber(8)).toNumber()] = ((((_2437_v32).contains((_this).f14)) ? ((_2437_v32).get((_this).f14)) : (_2429_v24))).minus(_2429_v24);
          _nw411[(new BigNumber(9)).toNumber()] = (_2429_v24).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rh")).length));
          _nw411[(new BigNumber(10)).toNumber()] = new BigNumber(334);
          _nw411[(new BigNumber(11)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(12)).toNumber()] = (((_this).f14) ? (new BigNumber((_2434_v29).length)) : (_2429_v24));
          _nw411[(new BigNumber(13)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(14)).toNumber()] = new BigNumber((_2436_v31).cardinality());
          _nw411[(new BigNumber(15)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(16)).toNumber()] = (((_this).f14) ? (_2429_v24) : ((_dafny.ZERO).minus(_2429_v24)));
          _nw411[(new BigNumber(17)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(18)).toNumber()] = _module.__default.fm0((_this).f13, globalState);
          _nw411[(new BigNumber(19)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(20)).toNumber()] = new BigNumber(514);
          _nw411[(new BigNumber(21)).toNumber()] = new BigNumber(278);
          _nw411[(new BigNumber(22)).toNumber()] = (_2434_v29)[_module.__default.safeIndex((_dafny.ZERO).minus(_2429_v24), new BigNumber((_2434_v29).length))];
          _nw411[(new BigNumber(23)).toNumber()] = new BigNumber((_2438_v33).length);
          _nw411[(new BigNumber(24)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(25)).toNumber()] = _2429_v24;
          _nw411[(new BigNumber(26)).toNumber()] = new BigNumber((_2434_v29).length);
          _2439_v34 = _nw411;
          let _index407 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2439_v34).length));
          (_2439_v34)[_index407] = new BigNumber(909);
          let _2440_v35;
          _2440_v35 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber(962)));
          let _index408 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2439_v34).length));
          let _rhs462 = _2431_v26;
          let _rhs463 = (_2435_v30).Merge(_2435_v30);
          let _rhs464 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2440_v35, _module.__default.safeIndex(_2429_v24, new BigNumber((_2440_v35).length)), _2437_v32), _2440_v35)).length);
          let _rhs465 = _module.__default.fm1(!((_this).f14), _2431_v26, globalState);
          let _lhs355 = _2439_v34;
          let _lhs356 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2439_v34).length));
          let _lhs357 = globalState;
          r0 = _rhs462;
          _2435_v30 = _rhs463;
          _lhs355[_lhs356] = _rhs464;
          _lhs357.f8 = _rhs465;
          let _2441_v36;
          _2441_v36 = _dafny.Seq.of(!(_module.__default.fm1((_this).f13, _2431_v26, globalState)));
          _2441_v36 = _dafny.Seq.Concat(_2441_v36, _2441_v36);
          let _2442_v37;
          let _nw412 = new _module.C0();
          _nw412.__ctor(!_dafny.areEqual(_2430_v25, _2430_v25));
          _2442_v37 = _nw412;
        } else if (_source30.is_DC7) {
          let _2443___mcc_h0 = (_source30).cf11;
          let _2444___mcc_h1 = (_source30).cf12;
          let _2445___mcc_h2 = (_source30).cf13;
          let _2446___mcc_h3 = (_source30).cf14;
          let _2447_cf14 = _2446___mcc_h3;
          let _2448_cf13 = _2445___mcc_h2;
          let _2449_cf12 = _2444___mcc_h1;
          let _2450_cf11 = _2443___mcc_h0;
          let _2451_v38;
          let _nw413 = Array((new BigNumber(24)).toNumber());
          _nw413[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_2429_v24);
          _nw413[(_dafny.ONE).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(2)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(3)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(4)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(5)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(6)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(7)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(8)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(9)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(10)).toNumber()] = new BigNumber(-98);
          _nw413[(new BigNumber(11)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(12)).toNumber()] = new BigNumber(288);
          _nw413[(new BigNumber(13)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(14)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(15)).toNumber()] = _2429_v24;
          _nw413[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(937)), ((_2452_cf12) => function (_2453_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(true,_2452_cf12);
          })(_2449_cf12))).length);
          _nw413[(new BigNumber(17)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_2449_cf12);
          _nw413[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("tnrgo")).length);
          _nw413[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_2449_cf12);
          _nw413[(new BigNumber(21)).toNumber()] = _module.__default.fm0((_this).f14, globalState);
          _nw413[(new BigNumber(22)).toNumber()] = _2449_cf12;
          _nw413[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), ((_2454_v26) => function (_2455_i3) {
            return _2454_v26;
          })(_2431_v26))).length);
          _2451_v38 = _nw413;
          let _2456_v39;
          let _2457_v40;
          let _2458_v41;
          let _2459_v42;
          let _out15;
          let _out16;
          let _out17;
          let _out18;
          let _outcollector5 = (_2428_v23).m4(_2429_v24, _2451_v38, globalState);
          _out15 = _outcollector5[0];
          _out16 = _outcollector5[1];
          _out17 = _outcollector5[2];
          _out18 = _outcollector5[3];
          _2456_v39 = _out15;
          _2457_v40 = _out16;
          _2458_v41 = _out17;
          _2459_v42 = _out18;
          let _2460_v43;
          _2460_v43 = _dafny.Seq.of(true);
          let _2461_v44;
          _2461_v44 = _dafny.Seq.of(_module.__default.fm1(!((_2460_v43)[_module.__default.safeIndex(_2429_v24, new BigNumber((_2460_v43).length))]), _2431_v26, globalState));
          _2459_v42 = ((_this).f13) || (_dafny.Seq.IsProperPrefixOf(_2461_v44, _2461_v44));
          _2430_v25 = _2447_cf14;
          _2450_cf11 = ((new BigNumber(340)).plus(_2457_v40)).isLessThan((_dafny.ZERO).minus((_2449_cf12).plus(new BigNumber(897))));
        } else if (_source30.is_DC8) {
          let _2462___mcc_h4 = (_source30).cf15;
          let _2463___mcc_h5 = (_source30).cf16;
          let _2464_cf16 = _2463___mcc_h5;
          let _2465_cf15 = _2462___mcc_h4;
          let _2466_v45;
          let _nw414 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _2466_v45 = _nw414;
          let _2467_v46;
          _2467_v46 = _dafny.MultiSet.fromElements(_2465_cf15, _2465_cf15);
          let _2468_v47;
          _2468_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2467_v46);
          let _index409 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2466_v45).length));
          (_2466_v45)[_index409] = (_dafny.ZERO).minus(new BigNumber((_2468_v47).length));
          let _2469_v48;
          _2469_v48 = _dafny.Seq.of((_this).f14);
          let _2470_v49;
          _2470_v49 = _dafny.Seq.of(_dafny.Seq.of((_this).f14), _dafny.Seq.update(_2469_v48, _module.__default.safeIndex(_module.__default.fm0((_this).f14, globalState), new BigNumber((_2469_v48).length)), (_this).f13));
          let _index410 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2466_v45).length));
          (_2466_v45)[_index410] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm44(_2470_v49, _2431_v26, globalState), _dafny.Seq.Concat(_2469_v48, _2469_v48))).length);
          (globalState).f8 = (_this).f13;
          let _2471_v50;
          let _nw415 = new _module.C0();
          _nw415.__ctor((_this).f13);
          _2471_v50 = _nw415;
          let _2472_v51;
          let _nw416 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _2472_v51 = _nw416;
          let _index411 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2472_v51).length));
          (_2472_v51)[_index411] = _module.__default.fm50(_2464_cf16, (_this).f13, !((_this).f13), globalState);
          let _2473_v52;
          _2473_v52 = _module.D4.create_DC11(_2465_cf15, _2465_cf15, _module.__default.fm1((_2471_v50).f17, _2431_v26, globalState));
          let _2474_v53;
          _2474_v53 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.UnicodeFromString("onjpllpsl")).length)).isLessThan(new BigNumber(614)),(_2473_v52).dtor_cf19);
          let _index412 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2472_v51).length));
          (_2472_v51)[_index412] = _2474_v53;
        } else {
          let _2475___mcc_h6 = (_source30).cf10;
          let _2476_cf10 = _2475___mcc_h6;
          (globalState).f9 = _module.__default.safeDivisionInt(_2429_v24, _2429_v24);
          let _2477_v54;
          let _init78 = function (_2478_i4) {
            return (_this).f14;
          };
          let _nw417 = Array((new BigNumber(17)).toNumber());
          for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw417.length); _i0_78++) {
            _nw417[_i0_78] = _init78(new BigNumber(_i0_78));
          }
          _2477_v54 = _nw417;
          let _index413 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_2477_v54).length));
          (_2477_v54)[_index413] = (_this).f13;
          let _2479_v55;
          let _nw418 = Array((new BigNumber(19)).toNumber()).fill([]);
          _2479_v55 = _nw418;
          let _2480_v56;
          _2480_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f14),true);
          let _2481_v57;
          _2481_v57 = _module.D2.create_DC8(new BigNumber((_2480_v56).length), _2431_v26);
          let _2482_v58;
          let _nw419 = Array((new BigNumber(7)).toNumber());
          _nw419[(_dafny.ZERO).toNumber()] = _2431_v26;
          _nw419[(_dafny.ONE).toNumber()] = (_2481_v57).dtor_cf16;
          _nw419[(new BigNumber(2)).toNumber()] = _2431_v26;
          _nw419[(new BigNumber(3)).toNumber()] = _2431_v26;
          _nw419[(new BigNumber(4)).toNumber()] = _2431_v26;
          _nw419[(new BigNumber(5)).toNumber()] = _2431_v26;
          _nw419[(new BigNumber(6)).toNumber()] = _2431_v26;
          _2482_v58 = _nw419;
          let _index414 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_2479_v55).length));
          (_2479_v55)[_index414] = _2482_v58;
          let _index415 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_2477_v54).length));
          let _index416 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_2479_v55).length));
          let _rhs466 = (_this).fm10(globalState);
          let _rhs467 = _2482_v58;
          let _rhs468 = (_this).f13;
          let _lhs358 = _2477_v54;
          let _lhs359 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_2477_v54).length));
          let _lhs360 = _2479_v55;
          let _lhs361 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_2479_v55).length));
          let _lhs362 = globalState;
          _lhs358[_lhs359] = _rhs466;
          _lhs360[_lhs361] = _rhs467;
          _lhs362.f8 = _rhs468;
          let _2483_v59;
          _2483_v59 = _dafny.Map.Empty.slice().updateUnsafe(_2429_v24,_2429_v24);
          (globalState).f9 = new BigNumber((_2483_v59).length);
          (globalState).f9 = _module.__default.safeModuloInt(_2429_v24, new BigNumber((_module.__default.fm62((_this).f14, globalState)).length));
        }
        let _2484_v60;
        let _nw420 = new _module.C5();
        _nw420.__ctor((_this).f11);
        _2484_v60 = _nw420;
        let _2485_v61;
        let _2486_v62;
        let _out19;
        let _out20;
        let _outcollector6 = (_2428_v23).m3((_this).f14, globalState);
        _out19 = _outcollector6[0];
        _out20 = _outcollector6[1];
        _2485_v61 = _out19;
        _2486_v62 = _out20;
        _2485_v61 = _2429_v24;
        let _2487_v63;
        let _nw421 = Array((new BigNumber(8)).toNumber()).fill(false);
        _2487_v63 = _nw421;
        let _index417 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2487_v63).length));
        (_2487_v63)[_index417] = (_this).f13;
        let _index418 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2487_v63).length));
        (_2487_v63)[_index418] = _module.__default.fm1((_this).f14, _2431_v26, globalState);
      }
      let _2488_v64;
      _2488_v64 = new BigNumber(475);
      let _rhs469 = _2488_v64;
      let _rhs470 = _2488_v64;
      let _lhs363 = globalState;
      let _lhs364 = globalState;
      _lhs363.f9 = _rhs469;
      _lhs364.f9 = _rhs470;
      let _2489_v65;
      _2489_v65 = _dafny.Seq.of(_2488_v64);
      let _2490_v66;
      let _init79 = function (_2491_i6) {
        return false;
      };
      let _nw422 = Array((new BigNumber(6)).toNumber());
      for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw422.length); _i0_79++) {
        _nw422[_i0_79] = _init79(new BigNumber(_i0_79));
      }
      _2490_v66 = _nw422;
      let _2492_v67;
      _2492_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2489_v65,_2490_v66);
      let _hi18 = _2488_v64;
      for (let _2493_i5 = new BigNumber(((_2492_v67).update(_2489_v65, _2490_v66)).length); _2493_i5.isLessThan(_hi18); _2493_i5 = _2493_i5.plus(_dafny.ONE)) {
        let _index419 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2490_v66).length));
        (_2490_v66)[_index419] = (_2488_v64).isLessThan(_2493_i5);
        let _2494_v68;
        _2494_v68 = _dafny.Seq.of((_this).f14);
        let _2495_v69;
        _2495_v69 = _module.D0.create_DC2(_dafny.MultiSet.FromArray(_2494_v68), (_this).f13, _2488_v64, (_this).f14);
        let _2496_v70;
        _2496_v70 = _dafny.Set.fromElements(_2488_v64);
        let _2497_v71;
        _2497_v71 = _dafny.Seq.UnicodeFromString("vbrdtah");
        let _2498_v72;
        let _nw423 = Array((new BigNumber(23)).toNumber());
        _nw423[(_dafny.ZERO).toNumber()] = _2488_v64;
        _nw423[(_dafny.ONE).toNumber()] = _2488_v64;
        _nw423[(new BigNumber(2)).toNumber()] = _2488_v64;
        _nw423[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_2488_v64, _2493_i5);
        _nw423[(new BigNumber(4)).toNumber()] = _2488_v64;
        _nw423[(new BigNumber(5)).toNumber()] = (_2493_i5).multipliedBy(_2493_i5);
        _nw423[(new BigNumber(6)).toNumber()] = _2493_i5;
        _nw423[(new BigNumber(7)).toNumber()] = _2493_i5;
        _nw423[(new BigNumber(8)).toNumber()] = _2493_i5;
        _nw423[(new BigNumber(9)).toNumber()] = (new BigNumber(208)).minus(_2488_v64);
        _nw423[(new BigNumber(10)).toNumber()] = _2488_v64;
        _nw423[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((function (_pat_let41_0) {
          return function (_2499_dt__update__tmp_h0) {
            return function (_pat_let42_0) {
              return function (_2500_dt__update_hcf6_h0) {
                return _module.D0.create_DC2((_2499_dt__update__tmp_h0).dtor_cf4, (_2499_dt__update__tmp_h0).dtor_cf5, _2500_dt__update_hcf6_h0, (_2499_dt__update__tmp_h0).dtor_cf7);
              }(_pat_let42_0);
            }(_2493_i5);
          }(_pat_let41_0);
        }(_2495_v69)).dtor_cf6, new BigNumber(183));
        _nw423[(new BigNumber(12)).toNumber()] = new BigNumber(471);
        _nw423[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(_2493_i5, _2488_v64);
        _nw423[(new BigNumber(14)).toNumber()] = _2488_v64;
        _nw423[(new BigNumber(15)).toNumber()] = (_2493_i5).minus(_2488_v64);
        _nw423[(new BigNumber(16)).toNumber()] = (((_this).f14) ? (_2488_v64) : (_2493_i5));
        _nw423[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(_2488_v64, _2493_i5);
        _nw423[(new BigNumber(18)).toNumber()] = _2488_v64;
        _nw423[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(946), new BigNumber((_2494_v68).length));
        _nw423[(new BigNumber(20)).toNumber()] = _2493_i5;
        _nw423[(new BigNumber(21)).toNumber()] = new BigNumber((_2496_v70).length);
        _nw423[(new BigNumber(22)).toNumber()] = new BigNumber((_2497_v71).length);
        _2498_v72 = _nw423;
        let _index420 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_2498_v72).length));
        (_2498_v72)[_index420] = _2493_i5;
        let _2501_v73;
        _2501_v73 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), ((_2502_v71, _2503_v68) => function (_2504_i7) {
          return (_2502_v71)[_module.__default.safeIndex(new BigNumber((_2503_v68).length), new BigNumber((_2502_v71).length))];
        })(_2497_v71, _2494_v68)));
        let _index421 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2490_v66).length));
        let _index422 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_2498_v72).length));
        let _rhs471 = ((!((_dafny.ZERO).minus(_2493_i5)).isEqualTo(new BigNumber(-11))) ? (_2488_v64) : (_module.__default.fm0((_this).f14, globalState)));
        let _rhs472 = (_this).f13;
        let _rhs473 = (((_2496_v70).IsProperSubsetOf(_2496_v70)) ? (_2493_i5) : ((((_2501_v73).contains(_2497_v71)) ? ((_2501_v73).get(_2497_v71)) : (_module.__default.fm0((_this).f13, globalState)))));
        let _rhs474 = (new BigNumber(880)).minus(_2488_v64);
        let _lhs365 = globalState;
        let _lhs366 = _2490_v66;
        let _lhs367 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2490_v66).length));
        let _lhs368 = globalState;
        let _lhs369 = _2498_v72;
        let _lhs370 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_2498_v72).length));
        _lhs365.f9 = _rhs471;
        _lhs366[_lhs367] = _rhs472;
        _lhs368.f9 = _rhs473;
        _lhs369[_lhs370] = _rhs474;
        (globalState).f9 = _2493_i5;
        let _2505_v74;
        let _nw424 = new _module.C3();
        _nw424.__ctor((_this).f13, (_2490_v66)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2490_v66).length))], (_this).f11);
        _2505_v74 = _nw424;
        (globalState).f8 = (_2505_v74).f22;
      }
      (globalState).f8 = (_this).f13;
      let _2506_v75;
      _2506_v75 = _dafny.Seq.UnicodeFromString("ucbqax");
      let _2507_v76;
      _2507_v76 = _dafny.Seq.of(_2506_v75);
      let _2508_v77;
      _2508_v77 = _dafny.Map.Empty.slice().updateUnsafe((_2507_v76)[_module.__default.safeIndex(new BigNumber((_2506_v75).length), new BigNumber((_2507_v76).length))],(_this).f13);
      let _2509_v78;
      let _nw425 = new _module.C3();
      _nw425.__ctor((_this).f14, !((((_2508_v77).contains(_2506_v75)) ? ((_2508_v77).get(_2506_v75)) : ((_this).f13))), (_this).f11);
      _2509_v78 = _nw425;
      let _2510_v79;
      let _nw426 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _2510_v79 = _nw426;
      let _index423 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_2510_v79).length));
      (_2510_v79)[_index423] = new BigNumber(813);
      let _index424 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_2510_v79).length));
      (_2510_v79)[_index424] = (_2488_v64).multipliedBy(_2488_v64);
      r0 = new _dafny.CodePoint('e'.codePointAt(0));
      r1 = (_this).f14;
      r2 = (_2488_v64).isEqualTo((_2488_v64).multipliedBy(_2488_v64));
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _2511_v0;
      _2511_v0 = new BigNumber(989);
      let _2512_v1;
      _2512_v1 = _dafny.MultiSet.fromElements(new BigNumber(137));
      let _2513_v2;
      _2513_v2 = _dafny.MultiSet.fromElements(new BigNumber((_2512_v1).cardinality()));
      let _2514_v3;
      _2514_v3 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_2511_v0), _2513_v2, _2512_v1);
      let _2515_v4;
      let _nw427 = new _module.C3();
      _nw427.__ctor((_2514_v3).IsProperSubsetOf((_2514_v3).update(_2512_v1, _module.__default.abs(_2511_v0))), p0, (_this).f11);
      _2515_v4 = _nw427;
      let _hi19 = _2511_v0;
      for (let _2516_i0 = _module.__default.fm0((_this).f13, globalState); _2516_i0.isLessThan(_hi19); _2516_i0 = _2516_i0.plus(_dafny.ONE)) {
        (globalState).f9 = ((false) ? ((_dafny.ZERO).minus(_2511_v0)) : (_2511_v0));
        r1 = _dafny.Set.fromElements(false);
        (globalState).f8 = _2515_v4.f21;
        let _2517_v5;
        let _nw428 = new _module.C5();
        _nw428.__ctor((_this).f11);
        _2517_v5 = _nw428;
      }
      let _2518_v6;
      _2518_v6 = _dafny.Map.Empty.slice().updateUnsafe((_2515_v4).f22,(_2511_v0).minus(_2511_v0));
      _2518_v6 = (_2518_v6).update((_this).f13, new BigNumber(724));
      let _2519_v7;
      let _nw429 = new _module.C2();
      _nw429.__ctor((((_this).f13) ? ((_this).f11) : ((_this).f11)));
      _2519_v7 = _nw429;
      let _2520_i1;
      _2520_i1 = _dafny.ZERO;
      L11: {
        while ((_this).f13) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2520_i1)) {
              break L11;
            }
            _2520_i1 = (_2520_i1).plus(_dafny.ONE);
            let _2521_v8;
            let _init80 = ((_2522_v0) => function (_2523_i2) {
              return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_2522_v0), _dafny.Seq.of(_2522_v0, _2522_v0, new BigNumber(255), _2522_v0, _2522_v0));
            })(_2511_v0);
            let _nw430 = Array((new BigNumber(18)).toNumber());
            for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw430.length); _i0_80++) {
              _nw430[_i0_80] = _init80(new BigNumber(_i0_80));
            }
            _2521_v8 = _nw430;
            let _index425 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_2521_v8).length));
            (_2521_v8)[_index425] = !((_this).f13) || (_2515_v4.f21);
            let _index426 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_2521_v8).length));
            (_2521_v8)[_index426] = (_this).f14;
            let _2524_v9;
            _2524_v9 = _dafny.Map.Empty.slice().updateUnsafe(false,_2515_v4.f21);
            _2511_v0 = (((_2511_v0).isEqualTo(_2511_v0)) ? (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_2515_v4).f22,true)).Merge(_2524_v9)).length)) : (_2511_v0));
            let _2525_v10;
            let _nw431 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
            _2525_v10 = _nw431;
            _2525_v10 = _2525_v10;
            let _2526_v11;
            _2526_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_2527_v0) => function (_2528_i3) {
              return _2527_v0;
            })(_2511_v0)),(_2511_v0).minus(_2511_v0));
            let _2529_v12;
            _2529_v12 = _dafny.Seq.of(_2511_v0, _2511_v0);
            _2526_v11 = (_2526_v11).update(_dafny.Seq.Concat(_2529_v12, _2529_v12), _2511_v0);
          }
        }
      }
      let _2530_v13;
      let _nw432 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
      _2530_v13 = _nw432;
      let _nw433 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
      _2530_v13 = _nw433;
      let _2531_v14;
      _2531_v14 = new _dafny.CodePoint('l'.codePointAt(0));
      let _2532_v16;
      _2532_v16 = function () {
        let _coll77 = new _dafny.Set();
        for (const _compr_77 of ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),p0)).update(_2531_v14, (_this).f13)).Keys.Elements) {
          let _2533_v15 = _compr_77;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),p0)).update(_2531_v14, (_this).f13)).contains(_2533_v15)) {
            _coll77.add(_2533_v15);
          }
        }
        return _coll77;
      }();
      let _pat_let_tv45 = _2511_v0;
      r0 = function (_source31) {
        let _2534___mcc_h0 = _source31;
        let _2535_cf27 = _2534___mcc_h0;
        return (_module.D8.create_DC19(_dafny.Seq.UnicodeFromString("imiictkf"), _pat_let_tv45)).dtor_cf35;
      }(_2532_v16);
      let _2536_v17;
      _2536_v17 = _dafny.Set.fromElements((_this).f13, _module.__default.fm1(_2515_v4.f21, _2531_v14, globalState));
      r1 = _2536_v17;
      return [r0, r1];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f10 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f10) {
      let _this = this;
      (_this)._f10 = f10;
      return;
    }
    fm5(p0, p1, p2, globalState) {
      let _this = this;
      return _module.D0.create_DC2(_dafny.MultiSet.fromElements(false, (_this).f10), (_this).f10, (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length)).minus(new BigNumber(514)), (_this).f10);
    };
    fm6(globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_module.D0.create_DC0(new BigNumber((function () {
  let _coll78 = new _dafny.Map();
  for (const _compr_78 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of((_module.D0.create_DC0(new BigNumber(268))).dtor_cf0, new BigNumber(-311)))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(551))).length))).length))).Keys.Elements) {
    let _2537_v0 = _compr_78;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of((_module.D0.create_DC0(new BigNumber(268))).dtor_cf0, new BigNumber(-311)))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(551))).length))).length))).contains(_2537_v0)) {
      _coll78.push([_module.__default.safeDivisionInt(_2537_v0, new BigNumber((function () {
        let _coll79 = new _dafny.Map();
        for (const _compr_79 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-384),(_this).f10)).Keys.Elements) {
          let _2538_v1 = _compr_79;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-384),(_this).f10)).contains(_2538_v1)) {
            _coll79.push([_module.__default.safeDivisionInt(_2538_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(881),new BigNumber((_dafny.MultiSet.fromElements((_this).f10)).cardinality()))).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new _dafny.CodePoint('q'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-562), new BigNumber(-528))).length))).length))).length)),new BigNumber((_dafny.Seq.of((_this).f10)).length)]);
          }
        }
        return _coll79;
      }()).length)),new BigNumber(604)]);
    }
  }
  return _coll78;
}()).length)), _module.D0.create_DC0(new BigNumber(-987)), _module.D0.create_DC0(new BigNumber(-923)), _module.D0.create_DC0(new BigNumber(186)))).IsProperSubsetOf(_dafny.Set.fromElements(_module.D0.create_DC0(new BigNumber(801)), _module.D0.create_DC0(new BigNumber(239))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      (globalState).f9 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)), p1);
      if ((_this).fm6(globalState)) {
        let _2539_v1;
        _2539_v1 = _dafny.Seq.of(_dafny.Set.fromElements(p1), p0, function () {
          let _coll80 = new _dafny.Set();
          for (const _compr_80 of _dafny.IntegerRange(new BigNumber(163), new BigNumber(-320))) {
            let _2540_v0 = _compr_80;
            if (((new BigNumber(163)).isLessThanOrEqualTo(_2540_v0)) && ((_2540_v0).isLessThan(new BigNumber(-320)))) {
              _coll80.add((_2540_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("icn")).length)));
            }
          }
          return _coll80;
        }());
        _2539_v1 = _dafny.Seq.of(p0, function () {
          let _coll81 = new _dafny.Set();
          for (const _compr_81 of _dafny.IntegerRange(new BigNumber(844), new BigNumber(84))) {
            let _2541_v2 = _compr_81;
            if (((new BigNumber(844)).isLessThanOrEqualTo(_2541_v2)) && ((_2541_v2).isLessThan(new BigNumber(84)))) {
              _coll81.add(_module.__default.safeModuloInt(_2541_v2, new BigNumber(708)));
            }
          }
          return _coll81;
        }(), p0);
        let _2542_v3;
        _2542_v3 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_this).f10)),(_this).f10);
        r2 = !((new BigNumber(326)).isLessThan(((!((_this).f10)) ? (new BigNumber((_2542_v3).length)) : (p1))));
        let _2543_v4;
        let _init81 = function (_2544_i0) {
          return (_this).f10;
        };
        let _nw434 = Array((new BigNumber(19)).toNumber());
        for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw434.length); _i0_81++) {
          _nw434[_i0_81] = _init81(new BigNumber(_i0_81));
        }
        _2543_v4 = _nw434;
        let _2545_v5;
        let _nw435 = Array((new BigNumber(6)).toNumber());
        _nw435[(_dafny.ZERO).toNumber()] = p1;
        _nw435[(_dafny.ONE).toNumber()] = p1;
        _nw435[(new BigNumber(2)).toNumber()] = p1;
        _nw435[(new BigNumber(3)).toNumber()] = p1;
        _nw435[(new BigNumber(4)).toNumber()] = p1;
        _nw435[(new BigNumber(5)).toNumber()] = new BigNumber(192);
        _2545_v5 = _nw435;
        let _2546_v6;
        _2546_v6 = _module.D0.create_DC1(_dafny.MultiSet.fromElements(_2545_v5), p1, p1);
        let _2547_v7;
        _2547_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6);
        let _2548_v8;
        _2548_v8 = _dafny.Seq.of(_2547_v7);
        let _2549_v9;
        _2549_v9 = _dafny.MultiSet.fromElements(_2545_v5);
        let _2550_v10;
        _2550_v10 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2547_v7);
        let _2551_v11;
        _2551_v11 = _dafny.Seq.UnicodeFromString("jxkmnorw");
        let _2552_v12;
        _2552_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2551_v11,_2547_v7);
        let _pat_let_tv46 = _2549_v9;
        let _2553_v13;
        let _nw436 = Array((new BigNumber(12)).toNumber());
        _nw436[(_dafny.ZERO).toNumber()] = _2547_v7;
        _nw436[(_dafny.ONE).toNumber()] = (_2548_v8)[_module.__default.safeIndex(p1, new BigNumber((_2548_v8).length))];
        _nw436[(new BigNumber(2)).toNumber()] = _2547_v7;
        _nw436[(new BigNumber(3)).toNumber()] = ((_2547_v7).update(_2543_v4, _2546_v6)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2543_v4,function (_pat_let43_0) {
          return function (_2554_dt__update__tmp_h0) {
            return function (_pat_let44_0) {
              return function (_2555_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_2555_dt__update_hcf1_h0, (_2554_dt__update__tmp_h0).dtor_cf2, (_2554_dt__update__tmp_h0).dtor_cf3);
              }(_pat_let44_0);
            }(_pat_let_tv46);
          }(_pat_let43_0);
        }(_module.D0.create_DC1(_dafny.MultiSet.fromElements(_2545_v5), p1, p1))));
        _nw436[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6)).Merge(_2547_v7);
        _nw436[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6);
        _nw436[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6)).Merge((_2547_v7).update(_2543_v4, _2546_v6));
        _nw436[(new BigNumber(7)).toNumber()] = (((_2550_v10).contains(p1)) ? ((_2550_v10).get(p1)) : (_2547_v7));
        _nw436[(new BigNumber(8)).toNumber()] = _2547_v7;
        _nw436[(new BigNumber(9)).toNumber()] = _2547_v7;
        _nw436[(new BigNumber(10)).toNumber()] = ((((_2552_v12).contains(_2551_v11)) ? ((_2552_v12).get(_2551_v11)) : (_dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6)))).Merge(_2547_v7);
        _nw436[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6);
        _2553_v13 = _nw436;
        let _2556_v14;
        _2556_v14 = _module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_2543_v4,_2546_v6));
        let _index427 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_2553_v13).length));
        (_2553_v13)[_index427] = (_2556_v14).dtor_cf8;
        let _index428 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_2553_v13).length));
        (_2553_v13)[_index428] = _2547_v7;
        (globalState).f8 = (_this).f10;
        let _2557_v15;
        _2557_v15 = _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, p1)));
        let _2558_v16;
        let _nw437 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2558_v16 = _nw437;
        let _2559_v17;
        _2559_v17 = _dafny.MultiSet.fromElements(_2546_v6);
        let _index429 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_2558_v16).length));
        (_2558_v16)[_index429] = _2559_v17;
        let _2560_v18;
        _2560_v18 = _dafny.Seq.of(_2551_v11);
        let _2561_v19;
        _2561_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_2557_v15).update((_this).f10, p1));
        let _2562_v20;
        _2562_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(((_2561_v19).contains((_this).f10)) ? ((_2561_v19).get((_this).f10)) : (_2557_v15)));
        let _2563_v21;
        _2563_v21 = _dafny.Seq.of((_2557_v15).Merge((((_2562_v20).contains((_this).f10)) ? ((_2562_v20).get((_this).f10)) : (_2557_v15))));
        let _index430 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_2558_v16).length));
        let _rhs475 = (_2563_v21)[_module.__default.safeIndex(p1, new BigNumber((_2563_v21).length))];
        let _rhs476 = _2559_v17;
        let _rhs477 = _dafny.Seq.Concat(_dafny.Seq.of(_2551_v11), _2560_v18);
        let _rhs478 = (_this).f10;
        let _lhs371 = _2558_v16;
        let _lhs372 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_2558_v16).length));
        let _lhs373 = globalState;
        _2557_v15 = _rhs475;
        _lhs371[_lhs372] = _rhs476;
        _2560_v18 = _rhs477;
        _lhs373.f8 = _rhs478;
      } else {
        let _2564_v22;
        _2564_v22 = _dafny.Seq.UnicodeFromString("i");
        let _2565_v23;
        _2565_v23 = new _dafny.CodePoint('u'.codePointAt(0));
        let _2566_v24;
        _2566_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f10);
        let _2567_v25;
        _2567_v25 = _dafny.Seq.of((_this).f10, (_this).f10);
        let _2568_v26;
        _2568_v26 = _module.D0.create_DC2(_dafny.MultiSet.fromElements(true), false, p1, (_2567_v25)[_module.__default.safeIndex(p1, new BigNumber((_2567_v25).length))]);
        let _2569_v27;
        _2569_v27 = _dafny.MultiSet.fromElements((_2568_v26).dtor_cf6, p1);
        let _2570_v29;
        _2570_v29 = _module.D0.create_DC0(p1);
        let _2571_v30;
        _2571_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2570_v29);
        let _2572_v31;
        _2572_v31 = _dafny.MultiSet.fromElements(_2571_v30, _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2570_v29), _2571_v30);
        let _2573_v32;
        _2573_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2572_v31).cardinality()),_2565_v23);
        let _2574_v33;
        let _nw438 = Array((new BigNumber(19)).toNumber());
        _nw438[(_dafny.ZERO).toNumber()] = new BigNumber((_2566_v24).length);
        _nw438[(_dafny.ONE).toNumber()] = p1;
        _nw438[(new BigNumber(2)).toNumber()] = new BigNumber(472);
        _nw438[(new BigNumber(3)).toNumber()] = p1;
        _nw438[(new BigNumber(4)).toNumber()] = p1;
        _nw438[(new BigNumber(5)).toNumber()] = p1;
        _nw438[(new BigNumber(6)).toNumber()] = new BigNumber((_2564_v22).length);
        _nw438[(new BigNumber(7)).toNumber()] = p1;
        _nw438[(new BigNumber(8)).toNumber()] = p1;
        _nw438[(new BigNumber(9)).toNumber()] = new BigNumber((_2569_v27).cardinality());
        _nw438[(new BigNumber(10)).toNumber()] = p1;
        _nw438[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("cxhqskyew")).length);
        _nw438[(new BigNumber(12)).toNumber()] = new BigNumber((function () {
          let _coll82 = new _dafny.Map();
          for (const _compr_82 of _dafny.IntegerRange(new BigNumber(943), new BigNumber(-329))) {
            let _2575_v28 = _compr_82;
            if (((new BigNumber(943)).isLessThanOrEqualTo(_2575_v28)) && ((_2575_v28).isLessThan(new BigNumber(-329)))) {
              _coll82.push([(_2575_v28).plus(p1),!((_this).f10)]);
            }
          }
          return _coll82;
        }()).length);
        _nw438[(new BigNumber(13)).toNumber()] = p1;
        _nw438[(new BigNumber(14)).toNumber()] = p1;
        _nw438[(new BigNumber(15)).toNumber()] = new BigNumber(250);
        _nw438[(new BigNumber(16)).toNumber()] = new BigNumber((_2573_v32).length);
        _nw438[(new BigNumber(17)).toNumber()] = new BigNumber((_2569_v27).cardinality());
        _nw438[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(p1);
        _2574_v33 = _nw438;
        let _2576_v34;
        _2576_v34 = _dafny.MultiSet.fromElements(_2574_v33, _2574_v33);
        let _2577_v35;
        _2577_v35 = _module.D0.create_DC1(_2576_v34, new BigNumber((_2567_v25).length), p1);
        let _2578_v36;
        let _nw439 = Array((new BigNumber(15)).toNumber());
        _nw439[(_dafny.ZERO).toNumber()] = (((_this).f10) ? (p1) : (p1));
        _nw439[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.of((_this).f10, (_this).f10)).length));
        _nw439[(new BigNumber(2)).toNumber()] = new BigNumber((p0).length);
        _nw439[(new BigNumber(3)).toNumber()] = p1;
        _nw439[(new BigNumber(4)).toNumber()] = new BigNumber((_2564_v22).length);
        _nw439[(new BigNumber(5)).toNumber()] = (p1).multipliedBy(p1);
        _nw439[(new BigNumber(6)).toNumber()] = p1;
        _nw439[(new BigNumber(7)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2565_v23,_2574_v33)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2565_v23,_2574_v33))).length);
        _nw439[(new BigNumber(8)).toNumber()] = p1;
        _nw439[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber((_2564_v22).length));
        _nw439[(new BigNumber(10)).toNumber()] = p1;
        _nw439[(new BigNumber(11)).toNumber()] = p1;
        _nw439[(new BigNumber(12)).toNumber()] = (_2577_v35).dtor_cf3;
        _nw439[(new BigNumber(13)).toNumber()] = (p1).multipliedBy(p1);
        _nw439[(new BigNumber(14)).toNumber()] = p1;
        _2578_v36 = _nw439;
        _2578_v36 = _2578_v36;
        let _2579_v37;
        _2579_v37 = _module.D2.create_DC5(_2564_v22);
        let _2580_v38;
        let _nw440 = Array((new BigNumber(5)).toNumber());
        _nw440[(_dafny.ZERO).toNumber()] = _2564_v22;
        _nw440[(_dafny.ONE).toNumber()] = _2564_v22;
        _nw440[(new BigNumber(2)).toNumber()] = _module.__default.fm2(p1, _2569_v27, _module.__default.fm7(globalState), globalState);
        _nw440[(new BigNumber(3)).toNumber()] = (_2579_v37).dtor_cf10;
        _nw440[(new BigNumber(4)).toNumber()] = _2564_v22;
        _2580_v38 = _nw440;
        let _index431 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_2580_v38).length));
        (_2580_v38)[_index431] = _2564_v22;
        let _index432 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_2580_v38).length));
        (_2580_v38)[_index432] = _dafny.Seq.Concat(_2564_v22, _dafny.Seq.update(_2564_v22, _module.__default.safeIndex(p1, new BigNumber((_2564_v22).length)), _2565_v23));
        let _2581_v39;
        _2581_v39 = _module.D13.create_DC29(_2566_v24);
        let _2582_v40;
        let _nw441 = Array((new BigNumber(6)).toNumber());
        _nw441[(_dafny.ZERO).toNumber()] = _2566_v24;
        _nw441[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f10);
        _nw441[(new BigNumber(2)).toNumber()] = _2566_v24;
        _nw441[(new BigNumber(3)).toNumber()] = (_2581_v39).dtor_cf50;
        _nw441[(new BigNumber(4)).toNumber()] = _2566_v24;
        _nw441[(new BigNumber(5)).toNumber()] = (_2581_v39).dtor_cf50;
        _2582_v40 = _nw441;
        let _2583_v41;
        let _nw442 = new _module.C2();
        _nw442.__ctor(_2582_v40);
        _2583_v41 = _nw442;
        let _2584_v42;
        let _nw443 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2584_v42 = _nw443;
        let _2585_v43;
        _2585_v43 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,p1),p1)).length));
        let _index433 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2584_v42).length));
        (_2584_v42)[_index433] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_2585_v43, _2585_v43));
        let _2586_v44;
        _2586_v44 = _dafny.Seq.of(_2569_v27, _dafny.MultiSet.fromElements(p1));
        let _index434 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2584_v42).length));
        (_2584_v42)[_index434] = ((_2586_v44)[_module.__default.safeIndex(p1, new BigNumber((_2586_v44).length))]).Union(_2569_v27);
        let _2587_v45;
        _2587_v45 = _module.D4.create_DC10(_dafny.Set.fromElements((_this).f10));
        _2587_v45 = _2587_v45;
      }
      let _2588_v46;
      _2588_v46 = _dafny.Seq.of((_this).f10, !((_this).f10));
      r1 = new BigNumber((_2588_v46).length);
      let _2589_v47;
      _2589_v47 = new _dafny.CodePoint('o'.codePointAt(0));
      let _2590_v48;
      _2590_v48 = _dafny.Seq.of(_2589_v47, _2589_v47);
      let _2591_v49;
      _2591_v49 = _module.D1.create_DC4((_2590_v48)[_module.__default.safeIndex(p1, new BigNumber((_2590_v48).length))]);
      let _2592_v50;
      let _nw444 = Array((new BigNumber(8)).toNumber());
      _nw444[(_dafny.ZERO).toNumber()] = _module.__default.fm3(p1, !(!((_this).f10)), globalState);
      _nw444[(_dafny.ONE).toNumber()] = _2589_v47;
      _nw444[(new BigNumber(2)).toNumber()] = _2589_v47;
      _nw444[(new BigNumber(3)).toNumber()] = (_2591_v49).dtor_cf9;
      _nw444[(new BigNumber(4)).toNumber()] = _2589_v47;
      _nw444[(new BigNumber(5)).toNumber()] = _2589_v47;
      _nw444[(new BigNumber(6)).toNumber()] = _2589_v47;
      _nw444[(new BigNumber(7)).toNumber()] = (_2590_v48)[_module.__default.safeIndex(p1, new BigNumber((_2590_v48).length))];
      _2592_v50 = _nw444;
      let _index435 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_2592_v50).length));
      (_2592_v50)[_index435] = _2589_v47;
      let _index436 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_2592_v50).length));
      (_2592_v50)[_index436] = _2589_v47;
      if (((p1).plus(new BigNumber((p0).length))).isLessThan((_dafny.ZERO).minus(p1))) {
        r2 = (_this).f10;
        _module.__default.m0(globalState);
        (globalState).f9 = new BigNumber((_2590_v48).length);
        let _2593_v51;
        _2593_v51 = _dafny.Seq.of(new BigNumber(-738), p1);
        let _2594_v52;
        _2594_v52 = _dafny.MultiSet.fromElements((_this).f10);
        let _2595_v53;
        _2595_v53 = _module.D0.create_DC2((_2594_v52).update((_this).f10, _module.__default.abs(p1)), (_this).f10, p1, (_this).f10);
        let _2596_v54;
        _2596_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _2597_v55;
        _2597_v55 = _dafny.Set.fromElements(_dafny.Seq.of(p1), _2593_v51, _2593_v51, _2593_v51, _2593_v51);
        let _2598_v56;
        let _nw445 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _2598_v56 = _nw445;
        let _2599_v57;
        _2599_v57 = _dafny.Set.fromElements(_2598_v56, _2598_v56, _2598_v56);
        let _2600_v58;
        let _nw446 = Array((new BigNumber(24)).toNumber());
        _nw446[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(p1, new BigNumber((_2593_v51).length));
        _nw446[(_dafny.ONE).toNumber()] = (p1).minus(p1);
        _nw446[(new BigNumber(2)).toNumber()] = p1;
        _nw446[(new BigNumber(3)).toNumber()] = new BigNumber(309);
        _nw446[(new BigNumber(4)).toNumber()] = (_2595_v53).dtor_cf6;
        _nw446[(new BigNumber(5)).toNumber()] = new BigNumber((_2596_v54).length);
        _nw446[(new BigNumber(6)).toNumber()] = p1;
        _nw446[(new BigNumber(7)).toNumber()] = (p1).multipliedBy(p1);
        _nw446[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
        _nw446[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of((_this).f10), _2588_v46)).length), new BigNumber((_dafny.Seq.update(_2590_v48, _module.__default.safeIndex(p1, new BigNumber((_2590_v48).length)), (_2592_v50)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_2592_v50).length))])).length));
        _nw446[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(_module.__default.fm0((_this).f10, globalState), new BigNumber((_dafny.Seq.update(_2590_v48, _module.__default.safeIndex(p1, new BigNumber((_2590_v48).length)), new _dafny.CodePoint('a'.codePointAt(0)))).length));
        _nw446[(new BigNumber(11)).toNumber()] = new BigNumber((_2597_v55).length);
        _nw446[(new BigNumber(12)).toNumber()] = p1;
        _nw446[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw446[(new BigNumber(14)).toNumber()] = p1;
        _nw446[(new BigNumber(15)).toNumber()] = (_module.__default.fm0(true, globalState)).plus(p1);
        _nw446[(new BigNumber(16)).toNumber()] = p1;
        _nw446[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_2590_v48).length)).minus(p1)));
        _nw446[(new BigNumber(18)).toNumber()] = new BigNumber((_2599_v57).length);
        _nw446[(new BigNumber(19)).toNumber()] = new BigNumber((_2593_v51).length);
        _nw446[(new BigNumber(20)).toNumber()] = p1;
        _nw446[(new BigNumber(21)).toNumber()] = p1;
        _nw446[(new BigNumber(22)).toNumber()] = p1;
        _nw446[(new BigNumber(23)).toNumber()] = new BigNumber(-123);
        _2600_v58 = _nw446;
        let _index437 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_2600_v58).length));
        (_2600_v58)[_index437] = p1;
        let _index438 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_2600_v58).length));
        (_2600_v58)[_index438] = (p1).multipliedBy(p1);
        let _2601_v59;
        _2601_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2590_v48, _2590_v48),(_this).f10);
        if ((((_2601_v59).contains(_2590_v48)) ? ((_2601_v59).get(_2590_v48)) : ((_this).f10))) {
          (globalState).f8 = (_this).f10;
          _2590_v48 = _2590_v48;
          let _2602_v60;
          let _nw447 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _2602_v60 = _nw447;
          let _2603_v61;
          let _nw448 = new _module.C3();
          _nw448.__ctor((_this).f10, !((_this).f10), _2602_v60);
          _2603_v61 = _nw448;
          _2603_v61 = _2603_v61;
          let _2604_v62;
          let _nw449 = Array((new BigNumber(13)).toNumber()).fill(false);
          _2604_v62 = _nw449;
          let _index439 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2604_v62).length));
          (_2604_v62)[_index439] = (_this).f10;
          let _index440 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2604_v62).length));
          (_2604_v62)[_index440] = (_dafny.Seq.IsProperPrefixOf(_2588_v46, _2588_v46)) === ((_2588_v46)[_module.__default.safeIndex((_2600_v58)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_2600_v58).length))], new BigNumber((_2588_v46).length))]);
          let _2605_v63;
          _2605_v63 = _dafny.Map.Empty.slice().updateUnsafe((_2604_v62)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2604_v62).length))],(_2600_v58)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_2600_v58).length))]);
          let _2606_v64;
          _2606_v64 = _dafny.Map.Empty.slice().updateUnsafe((_2605_v63).Merge((_2605_v63).update(!((_2604_v62)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2604_v62).length))]), p1)),_2588_v46);
          _2606_v64 = (_2606_v64).update((_2605_v63).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,p1)), _dafny.Seq.Concat(_dafny.Seq.of((_2604_v62)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2604_v62).length))], (_2603_v61).f22), _dafny.Seq.of((_2604_v62)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2604_v62).length))], _2603_v61.f21)));
        } else {
          let _2607_v66;
          let _init82 = ((_2608_v58, _2609_p1) => function (_2610_i1) {
            return function () {
              let _coll83 = new _dafny.Map();
              for (const _compr_83 of _dafny.IntegerRange(new BigNumber(609), new BigNumber(188))) {
                let _2611_v65 = _compr_83;
                if (((new BigNumber(609)).isLessThanOrEqualTo(_2611_v65)) && ((_2611_v65).isLessThan(new BigNumber(188)))) {
                  _coll83.push([_module.__default.safeDivisionInt(_2611_v65, (_2608_v58)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_2608_v58).length))]),_2609_p1]);
                }
              }
              return _coll83;
            }();
          })(_2600_v58, p1);
          let _nw450 = Array((new BigNumber(16)).toNumber());
          for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw450.length); _i0_82++) {
            _nw450[_i0_82] = _init82(new BigNumber(_i0_82));
          }
          _2607_v66 = _nw450;
          let _2612_v67;
          _2612_v67 = _module.D10.create_DC23(_2607_v66);
          let _2613_v68;
          _2613_v68 = _dafny.Seq.of(_2612_v67, _module.D10.create_DC23(_2607_v66));
          _2613_v68 = _2613_v68;
          let _2614_v69;
          _2614_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(919),false);
          let _2615_v71;
          _2615_v71 = _module.D13.create_DC29(_2614_v69);
          let _2616_v73;
          let _nw451 = Array((new BigNumber(16)).toNumber());
          _nw451[(_dafny.ZERO).toNumber()] = _2614_v69;
          _nw451[(_dafny.ONE).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(2)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(3)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(4)).toNumber()] = function () {
            let _coll84 = new _dafny.Map();
            for (const _compr_84 of _dafny.IntegerRange(new BigNumber(706), new BigNumber(184))) {
              let _2617_v70 = _compr_84;
              if (((new BigNumber(706)).isLessThanOrEqualTo(_2617_v70)) && ((_2617_v70).isLessThan(new BigNumber(184)))) {
                _coll84.push([(_2617_v70).plus(p1),(_this).f10]);
              }
            }
            return _coll84;
          }();
          _nw451[(new BigNumber(5)).toNumber()] = (_2615_v71).dtor_cf50;
          _nw451[(new BigNumber(6)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(7)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(8)).toNumber()] = function () {
            let _coll85 = new _dafny.Map();
            for (const _compr_85 of _dafny.IntegerRange(new BigNumber(956), new BigNumber(115))) {
              let _2618_v72 = _compr_85;
              if (((new BigNumber(956)).isLessThanOrEqualTo(_2618_v72)) && ((_2618_v72).isLessThan(new BigNumber(115)))) {
                _coll85.push([(_2618_v72).multipliedBy((_2600_v58)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_2600_v58).length))]),(_this).f10]);
              }
            }
            return _coll85;
          }();
          _nw451[(new BigNumber(9)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(10)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(11)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(12)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(13)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(14)).toNumber()] = _2614_v69;
          _nw451[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f10)).length),(_this).f10);
          _2616_v73 = _nw451;
          let _2619_v74;
          let _nw452 = new _module.C2();
          _nw452.__ctor(_2616_v73);
          _2619_v74 = _nw452;
          let _2620_v75;
          let _nw453 = Array((new BigNumber(7)).toNumber()).fill(false);
          _2620_v75 = _nw453;
          _2620_v75 = _2620_v75;
          let _rhs479 = (_this).f10;
          let _rhs480 = p1;
          let _rhs481 = (((_this).f10) ? (_dafny.Seq.Concat(_2590_v48, _2590_v48)) : (_dafny.Seq.UnicodeFromString("cydhyslov")));
          let _lhs374 = globalState;
          let _lhs375 = globalState;
          _lhs374.f8 = _rhs479;
          _lhs375.f9 = _rhs480;
          _2590_v48 = _rhs481;
          let _2621_v76;
          _2621_v76 = _dafny.Set.fromElements((_this).f10, !((_this).f10));
          let _2622_v77;
          _2622_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2621_v76).length),_2592_v50);
          let _2623_v78;
          _2623_v78 = _dafny.Seq.of(_2620_v75, _2620_v75);
          _2622_v77 = (_2622_v77).update(new BigNumber((_dafny.Seq.Concat(_2623_v78, _2623_v78)).length), _2592_v50);
        }
      } else {
        let _2624_v79;
        _2624_v79 = _dafny.Seq.of(p1, p1);
        _2590_v48 = (((new BigNumber((_2624_v79).length)).isLessThan(new BigNumber((_dafny.Seq.of(false)).length))) ? (_2590_v48) : (_dafny.Seq.UnicodeFromString("g")));
        let _2625_v80;
        _2625_v80 = _dafny.MultiSet.fromElements(!((_this).f10), !(_module.__default.fm1((_this).f10, _2589_v47, globalState)));
        let _2626_v81;
        _2626_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_2625_v80).contains((_this).f10)) ? ((_2625_v80).get((_this).f10)) : (p1)),p1);
        let _2627_v82;
        _2627_v82 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17(globalState),(_dafny.ZERO).minus(new BigNumber((_2626_v81).length)));
        let _2628_v83;
        _2628_v83 = _module.D2.create_DC7(false, p1, true, _2590_v48);
        _2627_v82 = (_dafny.Map.Empty.slice().updateUnsafe(_2628_v83,p1)).Merge(_2627_v82);
        r0 = p1;
        let _2629_v84;
        let _nw454 = Array((new BigNumber(28)).toNumber()).fill([]);
        _2629_v84 = _nw454;
        let _2630_v85;
        let _nw455 = Array((new BigNumber(16)).toNumber());
        _2630_v85 = _nw455;
        let _2631_v86;
        let _init83 = ((_2632_p1) => function (_2633_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(_2632_p1,(_this).f10);
        })(p1);
        let _nw456 = Array((new BigNumber(17)).toNumber());
        for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw456.length); _i0_83++) {
          _nw456[_i0_83] = _init83(new BigNumber(_i0_83));
        }
        _2631_v86 = _nw456;
        let _2634_v87;
        let _nw457 = new _module.C4();
        _nw457.__ctor(_2631_v86);
        _2634_v87 = _nw457;
        let _2635_v88;
        _2635_v88 = _2634_v87;
        let _index441 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_2630_v85).length));
        (_2630_v85)[_index441] = (_2635_v88);
        let _2636_v89;
        _2636_v89 = _dafny.Set.fromElements((_this).f10, (_this).f10);
        let _2637_v90;
        _2637_v90 = _module.D4.create_DC10(_dafny.Set.fromElements(!((_this).f10)));
        let _index442 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_2630_v85).length));
        let _rhs482 = (_2636_v89).IsDisjointFrom(((_2637_v90).dtor_cf18).Intersect((_2637_v90).dtor_cf18));
        let _rhs483 = p1;
        let _rhs484 = (_module.__default.fm0((_this).f10, globalState)).plus((p1).plus(p1));
        let _rhs485 = _2629_v84;
        let _rhs486 = _2634_v87;
        let _lhs376 = globalState;
        let _lhs377 = globalState;
        let _lhs378 = _2630_v85;
        let _lhs379 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_2630_v85).length));
        _lhs376.f8 = _rhs482;
        _lhs377.f9 = _rhs483;
        r0 = _rhs484;
        _2629_v84 = _rhs485;
        _lhs378[_lhs379] = _rhs486;
        (globalState).f8 = (((_this).f10) ? ((_this).f10) : ((_this).f10));
      }
      let _2638_v91;
      _2638_v91 = _dafny.MultiSet.fromElements((_this).f10);
      let _2639_v92;
      _2639_v92 = _module.D0.create_DC2(_2638_v91, false, p1, false);
      let _2640_v93;
      _2640_v93 = _module.D8.create_DC20(_2639_v92, true, (_this).f10);
      let _source32 = _2640_v93;
      if (_source32.is_DC19) {
        let _2641___mcc_h0 = (_source32).cf34;
        let _2642___mcc_h1 = (_source32).cf35;
        let _2643_cf35 = _2642___mcc_h1;
        let _2644_cf34 = _2641___mcc_h0;
        let _2645_v94;
        _2645_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2643_cf35,(_this).f10);
        let _2646_v97;
        _2646_v97 = _dafny.Seq.of(_2588_v46);
        let _2647_v99;
        _2647_v99 = _module.D13.create_DC29(_2645_v94);
        let _pat_let_tv47 = _2645_v94;
        let _2648_v100;
        let _nw458 = Array((new BigNumber(17)).toNumber());
        _nw458[(_dafny.ZERO).toNumber()] = _2645_v94;
        _nw458[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(428),(_this).f10);
        _nw458[(new BigNumber(2)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(3)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(4)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(5)).toNumber()] = function () {
          let _coll86 = new _dafny.Map();
          for (const _compr_86 of _dafny.IntegerRange(new BigNumber(-281), new BigNumber(217))) {
            let _2649_v95 = _compr_86;
            if (((new BigNumber(-281)).isLessThanOrEqualTo(_2649_v95)) && ((_2649_v95).isLessThan(new BigNumber(217)))) {
              _coll86.push([(_2649_v95).multipliedBy(new BigNumber(322)),(_this).f10]);
            }
          }
          return _coll86;
        }();
        _nw458[(new BigNumber(6)).toNumber()] = function () {
          let _coll87 = new _dafny.Map();
          for (const _compr_87 of _dafny.IntegerRange(new BigNumber(185), new BigNumber(173))) {
            let _2650_v96 = _compr_87;
            if (((new BigNumber(185)).isLessThanOrEqualTo(_2650_v96)) && ((_2650_v96).isLessThan(new BigNumber(173)))) {
              _coll87.push([_module.__default.safeModuloInt(_2650_v96, new BigNumber(-70)),(_this).f10]);
            }
          }
          return _coll87;
        }();
        _nw458[(new BigNumber(7)).toNumber()] = _module.__default.fm27(_2646_v97, globalState);
        _nw458[(new BigNumber(8)).toNumber()] = function () {
          let _coll88 = new _dafny.Map();
          for (const _compr_88 of (p0).Elements) {
            let _2651_v98 = _compr_88;
            if ((p0).contains(_2651_v98)) {
              _coll88.push([_module.__default.safeDivisionInt(_2651_v98, p1),(_this).f10]);
            }
          }
          return _coll88;
        }();
        _nw458[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2643_cf35,true);
        _nw458[(new BigNumber(10)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(11)).toNumber()] = (function (_pat_let45_0) {
          return function (_2652_dt__update__tmp_h1) {
            return function (_pat_let46_0) {
              return function (_2653_dt__update_hcf50_h0) {
                return _module.D13.create_DC29(_2653_dt__update_hcf50_h0);
              }(_pat_let46_0);
            }(_pat_let_tv47);
          }(_pat_let45_0);
        }(_2647_v99)).dtor_cf50;
        _nw458[(new BigNumber(12)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(460),true);
        _nw458[(new BigNumber(14)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(15)).toNumber()] = _2645_v94;
        _nw458[(new BigNumber(16)).toNumber()] = _2645_v94;
        _2648_v100 = _nw458;
        let _2654_v101;
        let _nw459 = new _module.C7();
        _nw459.__ctor(!(_2643_cf35).isEqualTo(_module.__default.fm0((_this).f10, globalState)), _2648_v100, _module.D0.create_DC0(p1), !((_2643_cf35).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(_2643_cf35)).cardinality()))));
        _2654_v101 = _nw459;
        _2589_v47 = (_2592_v50)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_2592_v50).length))];
        _2588_v46 = _dafny.Seq.Concat(_2588_v46, _2588_v46);
        let _2655_v102;
        _2655_v102 = _dafny.Map.Empty.slice().updateUnsafe((_2654_v101).f14,(_2654_v101).f14);
        let _2656_v103;
        _2656_v103 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_2654_v101).f14), _2655_v102);
        let _2657_v104;
        _2657_v104 = _dafny.Seq.of((_dafny.ZERO).minus((_2654_v101).fm9(_2591_v49, _dafny.Set.fromElements(new BigNumber(((_2656_v103)[_module.__default.safeIndex(p1, new BigNumber((_2656_v103).length))]).length), p1), globalState)));
        r1 = _module.__default.safeModuloInt(new BigNumber((_2657_v104).length), new BigNumber((_2638_v91).cardinality()));
      } else if (_source32.is_DC20) {
        let _2658___mcc_h2 = (_source32).cf36;
        let _2659___mcc_h3 = (_source32).cf37;
        let _2660___mcc_h4 = (_source32).cf38;
        let _2661_cf38 = _2660___mcc_h4;
        let _2662_cf37 = _2659___mcc_h3;
        let _2663_cf36 = _2658___mcc_h2;
        let _2664_v105;
        let _init84 = function (_2665_i3) {
          return (_this).f10;
        };
        let _nw460 = Array((new BigNumber(26)).toNumber());
        for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw460.length); _i0_84++) {
          _nw460[_i0_84] = _init84(new BigNumber(_i0_84));
        }
        _2664_v105 = _nw460;
        let _2666_v106;
        _2666_v106 = _dafny.Seq.of((_dafny.ZERO).minus((_module.D8.create_DC19(_2590_v48, new BigNumber((_2590_v48).length))).dtor_cf35), p1, p1);
        let _index443 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2664_v105).length));
        (_2664_v105)[_index443] = ((_2666_v106)[_module.__default.safeIndex(p1, new BigNumber((_2666_v106).length))]).isLessThanOrEqualTo(p1);
        let _2667_v107;
        let _nw461 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _2667_v107 = _nw461;
        let _2668_v108;
        _2668_v108 = _dafny.MultiSet.fromElements(_2667_v107);
        let _2669_v109;
        _2669_v109 = _module.D0.create_DC1(_2668_v108, p1, new BigNumber(454));
        let _2670_v110;
        _2670_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2669_v109,_2664_v105);
        let _2671_v111;
        _2671_v111 = _dafny.Seq.of((((_2670_v110).contains(_2669_v109)) ? ((_2670_v110).get(_2669_v109)) : (_2664_v105)));
        let _index444 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2664_v105).length));
        let _rhs487 = (_2588_v46)[_module.__default.safeIndex(p1, new BigNumber((_2588_v46).length))];
        let _rhs488 = (_2671_v111)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_2671_v111).length))];
        let _rhs489 = (_2592_v50)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_2592_v50).length))];
        let _rhs490 = (_this).f10;
        let _lhs380 = _2664_v105;
        let _lhs381 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2664_v105).length));
        _lhs380[_lhs381] = _rhs487;
        _2664_v105 = _rhs488;
        _2589_v47 = _rhs489;
        r2 = _rhs490;
        let _rhs491 = ((p0).Union(p0)).equals((p0).Intersect(p0));
        let _rhs492 = _2662_cf37;
        let _rhs493 = p1;
        let _lhs382 = globalState;
        _2661_cf38 = _rhs491;
        _lhs382.f8 = _rhs492;
        r0 = _rhs493;
        let _2672_v112;
        let _nw462 = new _module.C6();
        _nw462.__ctor(p1, _2590_v48);
        _2672_v112 = _nw462;
        let _index445 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2664_v105).length));
        (_2664_v105)[_index445] = _2662_cf37;
      } else {
        let _2673___mcc_h5 = (_source32).cf33;
        let _2674_cf33 = _2673___mcc_h5;
        (globalState).f8 = _module.__default.fm1((_this).f10, _2589_v47, globalState);
        let _2675_v113;
        let _nw463 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _2675_v113 = _nw463;
        let _2676_v114;
        _2676_v114 = _dafny.Seq.of(_2675_v113, _2675_v113);
        let _2677_v115;
        let _nw464 = new _module.C3();
        _nw464.__ctor((_this).f10, (_this).f10, (_2676_v114)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_2676_v114).length))]);
        _2677_v115 = _nw464;
        _2588_v46 = _dafny.Seq.of((new BigNumber(-322)).isLessThanOrEqualTo(p1), _2677_v115.f21, _2677_v115.f21, !(_2677_v115.f21));
        let _2678_v116;
        let _nw465 = Array((new BigNumber(15)).toNumber());
        _nw465[(_dafny.ZERO).toNumber()] = (_2588_v46)[_module.__default.safeIndex(p1, new BigNumber((_2588_v46).length))];
        _nw465[(_dafny.ONE).toNumber()] = (_2588_v46)[_module.__default.safeIndex(p1, new BigNumber((_2588_v46).length))];
        _nw465[(new BigNumber(2)).toNumber()] = (_2677_v115).f22;
        _nw465[(new BigNumber(3)).toNumber()] = (_2677_v115).f22;
        _nw465[(new BigNumber(4)).toNumber()] = false;
        _nw465[(new BigNumber(5)).toNumber()] = (_2677_v115).f22;
        _nw465[(new BigNumber(6)).toNumber()] = (_this).f10;
        _nw465[(new BigNumber(7)).toNumber()] = false;
        _nw465[(new BigNumber(8)).toNumber()] = _2677_v115.f21;
        _nw465[(new BigNumber(9)).toNumber()] = !(_2677_v115.f21);
        _nw465[(new BigNumber(10)).toNumber()] = _2677_v115.f21;
        _nw465[(new BigNumber(11)).toNumber()] = (_2677_v115).f22;
        _nw465[(new BigNumber(12)).toNumber()] = _2677_v115.f21;
        _nw465[(new BigNumber(13)).toNumber()] = (_2677_v115).f22;
        _nw465[(new BigNumber(14)).toNumber()] = (_this).f10;
        _2678_v116 = _nw465;
        let _2679_v117;
        _2679_v117 = _module.D7.create_DC17(_2590_v48, _2678_v116, (_2677_v115).f22, p1);
        let _2680_v118;
        _2680_v118 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2678_v116);
        let _2681_v119;
        _2681_v119 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).fm6(globalState));
        let _2682_v120;
        _2682_v120 = _module.D13.create_DC29(_2681_v119);
        let _2683_v121;
        _2683_v121 = _dafny.Seq.of(_2682_v120, _2682_v120, _2682_v120, _2682_v120, _module.D13.create_DC29((_2682_v120).dtor_cf50));
        _2679_v117 = _module.D7.create_DC17(_dafny.Seq.UnicodeFromString("oggvmdojg"), (((_2680_v118).contains(_2677_v115.f21)) ? ((_2680_v118).get(_2677_v115.f21)) : (_2678_v116)), _module.__default.fm1((_this).f10, (_2592_v50)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_2592_v50).length))], globalState), new BigNumber((function () {
  let _coll89 = new _dafny.Set();
  for (const _compr_89 of (_2683_v121).Elements) {
    let _2684_v122 = _compr_89;
    if (_dafny.Seq.contains(_2683_v121, _2684_v122)) {
      _coll89.add(_2684_v122);
    }
  }
  return _coll89;
}()).length));
      }
      let _2685_v123;
      let _init85 = ((_2686_p1) => function (_2687_i4) {
        return _dafny.Map.Empty.slice().updateUnsafe(_2686_p1,(_this).f10);
      })(p1);
      let _nw466 = Array((new BigNumber(2)).toNumber());
      for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw466.length); _i0_85++) {
        _nw466[_i0_85] = _init85(new BigNumber(_i0_85));
      }
      _2685_v123 = _nw466;
      let _2688_v124;
      _2688_v124 = _module.D0.create_DC0(p1);
      let _2689_v125;
      let _nw467 = new _module.C7();
      _nw467.__ctor((_this).f10, _2685_v123, _2688_v124, (_this).f10);
      _2689_v125 = _nw467;
      let _2690_v126;
      _2690_v126 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2689_v125);
      let _2691_v127;
      _2691_v127 = _dafny.MultiSet.fromElements(p1, p1);
      let _2692_v128;
      _2692_v128 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber(-533), new BigNumber((_module.__default.fm23((_this).f10, globalState)).length)), p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length), new BigNumber((_2690_v126).length), new BigNumber((_2691_v127).cardinality()));
      r0 = new BigNumber((_2691_v127).cardinality());
      r1 = _module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus(p1));
      r2 = (_this).f10;
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _2693_v0;
      _2693_v0 = _dafny.Seq.UnicodeFromString("mdworlyct");
      let _2694_v1;
      _2694_v1 = _dafny.Seq.of(_2693_v0);
      _2694_v1 = _2694_v1;
      let _2695_v2;
      let _nw468 = Array((new BigNumber(20)).toNumber()).fill(false);
      _2695_v2 = _nw468;
      let _2696_v3;
      _2696_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2695_v2);
      _2696_v3 = _2696_v3;
      let _2697_v4;
      _2697_v4 = new BigNumber(-574);
      let _hi20 = _module.__default.safeDivisionInt(new BigNumber(456), _2697_v4);
      for (let _2698_i0 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-574)), function (_2752_i1) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length); _2698_i0.isLessThan(_hi20); _2698_i0 = _2698_i0.plus(_dafny.ONE)) {
        let _index446 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
        (_2695_v2)[_index446] = (_this).f10;
        let _index447 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
        (_2695_v2)[_index447] = (_this).f10;
        let _2699_v5;
        _2699_v5 = new _dafny.CodePoint('i'.codePointAt(0));
        let _2700_v6;
        _2700_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2699_v5,false);
        let _2701_v7;
        _2701_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(new BigNumber((_2700_v6).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_2698_i0)));
        _2701_v7 = (_2701_v7).Merge(_2701_v7);
        let _2702_v8;
        _2702_v8 = _module.D2.create_DC8(_2698_i0, _2699_v5);
        let _source33 = _2702_v8;
        if (_source33.is_DC6) {
          let _2703_v9;
          _2703_v9 = _dafny.Seq.of((_this).f10);
          _2703_v9 = _2703_v9;
          let _2704_v11;
          _2704_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2698_i0,_2698_i0);
          let _2705_v12;
          _2705_v12 = _2704_v11;
          let _2706_v13;
          _2706_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2705_v12,_2698_i0);
          let _2707_v14;
          _2707_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2705_v12,(_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]);
          let _2708_v15;
          _2708_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2697_v4),(function () {
            let _coll90 = new _dafny.Map();
            for (const _compr_90 of (_2706_v13).Keys.Elements) {
              let _2709_v10 = _compr_90;
              if ((_2706_v13).contains(_2709_v10)) {
                _coll90.push([_2709_v10,!((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))])]);
              }
            }
            return _coll90;
          }()).Merge(_2707_v14));
          _2708_v15 = (_2708_v15).update(_module.__default.fm0((_this).f10, globalState), _2707_v14);
          let _2710_v16;
          _2710_v16 = _dafny.MultiSet.fromElements(_2697_v4, _2697_v4, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_2694_v1, _module.__default.safeIndex(_2697_v4, new BigNumber((_2694_v1).length)), _2693_v0))).cardinality()));
          let _2711_v17;
          _2711_v17 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),_2698_i0);
          (globalState).f8 = (new BigNumber(561)).isLessThanOrEqualTo((((_2710_v16).contains(_2697_v4)) ? ((_2710_v16).get(_2697_v4)) : ((((_2711_v17).contains((_this).f10)) ? ((_2711_v17).get((_this).f10)) : (new BigNumber((_2704_v11).length))))));
          _2693_v0 = (((((_2701_v7).contains(true)) ? ((_2701_v7).get(true)) : (true))) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iqu"), _2693_v0)) : (_2693_v0));
        } else if (_source33.is_DC7) {
          let _2712___mcc_h0 = (_source33).cf11;
          let _2713___mcc_h1 = (_source33).cf12;
          let _2714___mcc_h2 = (_source33).cf13;
          let _2715___mcc_h3 = (_source33).cf14;
          let _2716_cf14 = _2715___mcc_h3;
          let _2717_cf13 = _2714___mcc_h2;
          let _2718_cf12 = _2713___mcc_h1;
          let _2719_cf11 = _2712___mcc_h0;
          let _2720_v18;
          _2720_v18 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(638)).plus((_dafny.ZERO).minus(_2698_i0)),_2693_v0);
          _2720_v18 = (_2720_v18).update(_module.__default.fm0(_2719_cf11, globalState), _2716_cf14);
          let _2721_v19;
          let _init86 = ((_2722_v4, _2723_cf13) => function (_2724_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(_2722_v4,_2723_cf13);
          })(_2697_v4, _2717_cf13);
          let _nw469 = Array((new BigNumber(14)).toNumber());
          for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw469.length); _i0_86++) {
            _nw469[_i0_86] = _init86(new BigNumber(_i0_86));
          }
          _2721_v19 = _nw469;
          let _2725_v20;
          let _nw470 = new _module.C3();
          _nw470.__ctor(_2717_cf13, !(_2719_cf11), _2721_v19);
          _2725_v20 = _nw470;
          let _2726_v21;
          _2726_v21 = _dafny.Seq.of(_2698_i0, new BigNumber(320));
          let _rhs494 = _module.__default.fm1((_this).f10, _2699_v5, globalState);
          let _rhs495 = _2725_v20;
          let _rhs496 = _module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2697_v4),_2698_i0)).update(_2726_v21, new BigNumber((_2716_cf14).length))).length), new BigNumber((_dafny.Seq.Concat(_2726_v21, _2726_v21)).length));
          r0 = _rhs494;
          _2725_v20 = _rhs495;
          _2718_cf12 = _rhs496;
          _2719_cf11 = !(_2719_cf11);
          let _index448 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
          let _index449 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
          let _rhs497 = !(_2717_cf13);
          let _rhs498 = !(!(_2719_cf11));
          let _lhs383 = _2695_v2;
          let _lhs384 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
          let _lhs385 = _2695_v2;
          let _lhs386 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
          _lhs383[_lhs384] = _rhs497;
          _lhs385[_lhs386] = _rhs498;
        } else if (_source33.is_DC8) {
          let _2727___mcc_h4 = (_source33).cf15;
          let _2728___mcc_h5 = (_source33).cf16;
          let _2729_cf16 = _2728___mcc_h5;
          let _2730_cf15 = _2727___mcc_h4;
          let _2731_v22;
          _2731_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]),!((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]));
          let _2732_v23;
          _2732_v23 = _dafny.Seq.of((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]);
          let _2733_v24;
          _2733_v24 = _dafny.Map.Empty.slice().updateUnsafe((_2731_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2732_v23,(_this).f10)),_dafny.Seq.of(_2697_v4, new BigNumber(622), new BigNumber((_2732_v23).length), _2697_v4));
          _2697_v4 = new BigNumber((_dafny.Seq.update((((_2733_v24).contains(_2731_v22)) ? ((_2733_v24).get(_2731_v22)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), function (_2734_i3) {
            return new BigNumber(145);
          }))), _module.__default.safeIndex(_2730_cf15, new BigNumber(((((_2733_v24).contains(_2731_v22)) ? ((_2733_v24).get(_2731_v22)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), function (_2735_i3) {
            return new BigNumber(145);
          })))).length)), _2697_v4)).length);
          (globalState).f9 = _2730_cf15;
          let _2736_v25;
          let _nw471 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _2736_v25 = _nw471;
          let _2737_v26;
          let _nw472 = new _module.C5();
          _nw472.__ctor(_2736_v25);
          _2737_v26 = _nw472;
          let _2738_v27;
          _2738_v27 = _module.D0.create_DC0(_module.__default.fm0((_this).f10, globalState));
          let _2739_v28;
          let _nw473 = new _module.C7();
          _nw473.__ctor((_this).f10, _2736_v25, _2738_v27, (_this).f10);
          _2739_v28 = _nw473;
          let _2740_v29;
          _2740_v29 = _dafny.Seq.of(_2739_v28);
          let _2741_v30;
          _2741_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2740_v29,false);
          let _index450 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length));
          (_2695_v2)[_index450] = !((((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]) ? (_module.__default.fm1((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))], _2699_v5, globalState)) : ((((_2741_v30).contains(_dafny.Seq.update(_2740_v29, _module.__default.safeIndex(_2730_cf15, new BigNumber((_2740_v29).length)), _2739_v28))) ? ((_2741_v30).get(_dafny.Seq.update(_2740_v29, _module.__default.safeIndex(_2730_cf15, new BigNumber((_2740_v29).length)), _2739_v28))) : ((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]))))) || (!((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]));
        } else {
          let _2742___mcc_h6 = (_source33).cf10;
          let _2743_cf10 = _2742___mcc_h6;
          _2701_v7 = (_2701_v7).update((_this).f10, (_this).f10);
          let _2744_v31;
          let _nw474 = new _module.C0();
          _nw474.__ctor(true);
          _2744_v31 = _nw474;
          let _2745_v32;
          let _nw475 = Array((new BigNumber(4)).toNumber());
          _nw475[(_dafny.ZERO).toNumber()] = _2744_v31;
          _nw475[(_dafny.ONE).toNumber()] = _2744_v31;
          _nw475[(new BigNumber(2)).toNumber()] = _2744_v31;
          _nw475[(new BigNumber(3)).toNumber()] = _2744_v31;
          _2745_v32 = _nw475;
          let _index451 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2745_v32).length));
          (_2745_v32)[_index451] = _2744_v31;
          let _index452 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2745_v32).length));
          (_2745_v32)[_index452] = _2744_v31;
          let _2746_v33;
          _2746_v33 = _module.D5.create_DC14((!((_2744_v31).f17)) === ((_this).f10));
          _2746_v33 = _module.__default.fm55((_2698_i0).isLessThan(_2697_v4), true, (_dafny.ZERO).minus(_2698_i0), globalState);
          let _2747_v34;
          _2747_v34 = _module.D8.create_DC20(_module.__default.fm4(_2697_v4, _2697_v4, globalState), (_2744_v31).f17, (_this).f10);
          _2747_v34 = _2747_v34;
        }
        let _2748_v35;
        _2748_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2698_i0,!((_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]));
        let _2749_v36;
        _2749_v36 = _module.D13.create_DC29(_2748_v35);
        let _2750_v37;
        let _nw476 = Array((_dafny.ONE).toNumber());
        _nw476[(_dafny.ZERO).toNumber()] = (_2749_v36).dtor_cf50;
        _2750_v37 = _nw476;
        let _index453 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2750_v37).length));
        (_2750_v37)[_index453] = _2748_v35;
        let _index454 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2750_v37).length));
        (_2750_v37)[_index454] = (_2748_v35).Merge((function () {
          let _coll91 = new _dafny.Map();
          for (const _compr_91 of _dafny.IntegerRange(new BigNumber(518), new BigNumber(-893))) {
            let _2751_v38 = _compr_91;
            if (((new BigNumber(518)).isLessThanOrEqualTo(_2751_v38)) && ((_2751_v38).isLessThan(new BigNumber(-893)))) {
              _coll91.push([(_2751_v38).plus(_2697_v4),false]);
            }
          }
          return _coll91;
        }()).update(new BigNumber(173), (_2695_v2)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2695_v2).length))]));
      }
      r0 = (_this).fm6(globalState);
      let _2753_v40;
      _2753_v40 = _dafny.Seq.of(_2697_v4, _2697_v4);
      let _2754_v41;
      _2754_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2753_v40).length),_2697_v4);
      let _2755_v42;
      let _nw477 = Array((new BigNumber(4)).toNumber());
      _nw477[(_dafny.ZERO).toNumber()] = function () {
        let _coll92 = new _dafny.Map();
        for (const _compr_92 of _dafny.IntegerRange(new BigNumber(403), new BigNumber(421))) {
          let _2756_v39 = _compr_92;
          if (((new BigNumber(403)).isLessThanOrEqualTo(_2756_v39)) && ((_2756_v39).isLessThan(new BigNumber(421)))) {
            _coll92.push([(_2756_v39).plus(_2697_v4),_2697_v4]);
          }
        }
        return _coll92;
      }();
      _nw477[(_dafny.ONE).toNumber()] = _2754_v41;
      _nw477[(new BigNumber(2)).toNumber()] = _2754_v41;
      _nw477[(new BigNumber(3)).toNumber()] = _2754_v41;
      _2755_v42 = _nw477;
      let _2757_v43;
      _2757_v43 = _module.D10.create_DC23(_2755_v42);
      let _pat_let_tv48 = _2755_v42;
      let _source34 = function (_pat_let47_0) {
        return function (_2758_dt__update__tmp_h0) {
          return function (_pat_let48_0) {
            return function (_2759_dt__update_hcf42_h0) {
              return _module.D10.create_DC23(_2759_dt__update_hcf42_h0);
            }(_pat_let48_0);
          }(_pat_let_tv48);
        }(_pat_let47_0);
      }(_2757_v43);
      if (_source34.is_DC24) {
        let _2760___mcc_h7 = (_source34).cf43;
        let _2761___mcc_h8 = (_source34).cf44;
        let _2762___mcc_h9 = (_source34).cf45;
        let _2763_cf45 = _2762___mcc_h9;
        let _2764_cf44 = _2761___mcc_h8;
        let _2765_cf43 = _2760___mcc_h7;
        let _index455 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_2764_cf44).length));
        (_2764_cf44)[_index455] = new BigNumber(-775);
        let _index456 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_2764_cf44).length));
        (_2764_cf44)[_index456] = _module.__default.fm0((_this).f10, globalState);
        let _2766_v44;
        let _nw478 = new _module.C6();
        _nw478.__ctor((_2764_cf44)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_2764_cf44).length))], _2693_v0);
        _2766_v44 = _nw478;
        let _2767_v45;
        let _nw479 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
        _2767_v45 = _nw479;
        let _2768_v46;
        let _nw480 = new _module.C5();
        _nw480.__ctor(_2767_v45);
        _2768_v46 = _nw480;
        let _2769_v47;
        _2769_v47 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2770_v48;
        _2770_v48 = _dafny.MultiSet.fromElements(_2769_v47, _2769_v47);
        (globalState).f8 = ((_2770_v48).Intersect(_2770_v48)).IsSubsetOf((_2770_v48).update(_2769_v47, _module.__default.abs((_2764_cf44)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_2764_cf44).length))])));
      } else {
        let _2771___mcc_h10 = (_source34).cf42;
        let _2772_cf42 = _2771___mcc_h10;
        let _2773_v49;
        _2773_v49 = _module.D11.create_DC26(_2697_v4);
        let _2774_v50;
        _2774_v50 = _module.D4.create_DC11(new BigNumber(602), _2697_v4, (_this).f10);
        let _2775_v51;
        _2775_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2774_v50,_2695_v2);
        let _2776_v52;
        _2776_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2773_v49,(((_2775_v51).contains(_2774_v50)) ? ((_2775_v51).get(_2774_v50)) : (_2695_v2)));
        let _2777_v53;
        _2777_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2697_v4,(((_this).f10) ? (_2695_v2) : ((((_2776_v52).contains(_2773_v49)) ? ((_2776_v52).get(_2773_v49)) : (_2695_v2)))));
        let _2778_v54;
        _2778_v54 = _dafny.MultiSet.fromElements(_2697_v4, new BigNumber((_dafny.Seq.UnicodeFromString("jcvn")).length), _2697_v4);
        let _2779_v55;
        _2779_v55 = _dafny.Seq.of(_2695_v2, _2695_v2);
        _2695_v2 = (((_2777_v53).contains(new BigNumber((_2778_v54).cardinality()))) ? ((_2777_v53).get(new BigNumber((_2778_v54).cardinality()))) : ((_2779_v55)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_2779_v55).length))]));
        let _2780_v56;
        let _nw481 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2780_v56 = _nw481;
        let _2781_v57;
        _2781_v57 = _dafny.Map.Empty.slice().updateUnsafe((_2697_v4).plus(_2697_v4),_2780_v56);
        _2781_v57 = (_2781_v57).update((_2697_v4).multipliedBy(_2697_v4), _2780_v56);
        _2754_v41 = (_2754_v41).update(new BigNumber(343), _2697_v4);
        let _2782_v58;
        _2782_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_2693_v0, _2693_v0)).length),(_this).f10);
        r0 = (((_2782_v58).contains((_2697_v4).minus(_2697_v4))) ? ((_2782_v58).get((_2697_v4).minus(_2697_v4))) : ((_this).f10));
      }
      let _2783_v59;
      let _nw482 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
      _2783_v59 = _nw482;
      let _2784_v60;
      _2784_v60 = _module.D0.create_DC0(_2697_v4);
      let _2785_v61;
      let _nw483 = new _module.C7();
      _nw483.__ctor(true, _2783_v59, _2784_v60, (_this).f10);
      _2785_v61 = _nw483;
      let _2786_v62;
      let _nw484 = new _module.C1();
      _nw484.__ctor((_2785_v61).f14, (_2785_v61).fm10(globalState), _2784_v60, (_this).f10);
      _2786_v62 = _nw484;
      let _2787_v63;
      _2787_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2785_v61,_2786_v62);
      let _2788_v64;
      _2788_v64 = _dafny.Seq.of(_2787_v63, _2787_v63);
      let _index457 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2695_v2).length));
      (_2695_v2)[_index457] = !((_2786_v62).f20) || ((_2786_v62).f20);
      let _2789_v65;
      let _nw485 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2789_v65 = _nw485;
      let _index458 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2789_v65).length));
      (_2789_v65)[_index458] = _2693_v0;
      let _2790_v66;
      _2790_v66 = _module.D7.create_DC17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), function (_2791_i4) {
  return new _dafny.CodePoint('j'.codePointAt(0));
}), _2695_v2, (_2786_v62).f19, _2697_v4);
      let _2792_v67;
      _2792_v67 = new _dafny.CodePoint('l'.codePointAt(0));
      let _index459 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2695_v2).length));
      let _index460 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2789_v65).length));
      let _rhs499 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2785_v61,_2786_v62), (_dafny.Map.Empty.slice().updateUnsafe(_2785_v61,_2786_v62)).Merge((_2787_v63).update(_2785_v61, _2786_v62)), (_dafny.Map.Empty.slice().updateUnsafe(_2785_v61,_2786_v62)).Merge(_2787_v63));
      let _rhs500 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("qldina"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(497)), function (_2793_i5) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })));
      let _rhs501 = _dafny.Seq.update(_2693_v0, _module.__default.safeIndex(_2697_v4, new BigNumber((_2693_v0).length)), _2792_v67);
      let _rhs502 = _2790_v66;
      let _lhs387 = _2695_v2;
      let _lhs388 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2695_v2).length));
      let _lhs389 = _2789_v65;
      let _lhs390 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2789_v65).length));
      _2788_v64 = _rhs499;
      _lhs387[_lhs388] = _rhs500;
      _lhs389[_lhs390] = _rhs501;
      _2790_v66 = _rhs502;
      let _2794_v68;
      _2794_v68 = _dafny.Map.Empty.slice().updateUnsafe((_2786_v62).f20,(_2785_v61).f14);
      let _2795_v69;
      _2795_v69 = _dafny.MultiSet.fromElements((_2786_v62).f20, (_2695_v2)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_2695_v2).length))]);
      r0 = (((_2794_v68).contains((_2795_v69).contains(!((_2695_v2)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_2695_v2).length))])))) ? ((_2794_v68).get((_2795_v69).contains(!((_2695_v2)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_2695_v2).length))])))) : ((_2785_v61).f14));
      return r0;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
