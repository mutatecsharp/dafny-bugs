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
      return _dafny.Seq.of(new BigNumber(616), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-434)), function (_0_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(26)), function (_1_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }))).length));
    };
    static fm1(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), function (_2_i0) {
        return new BigNumber(356);
      })).length))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(343)), _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-843))), _dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
            let _3_v0 = _compr_1;
            if ((_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).contains(_3_v0)) {
              _coll1.push([_3_v0,false]);
            }
          }
          return _coll1;
        }()).Keys.Elements) {
          let _4_v1 = _compr_0;
          if ((function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
              let _3_v0 = _compr_2;
              if ((_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).contains(_3_v0)) {
                _coll2.push([_3_v0,false]);
              }
            }
            return _coll2;
          }()).contains(_4_v1)) {
            _coll0.add(_4_v1);
          }
        }
        return _coll0;
      }()).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_5_i1) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length));
      })));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return !(!((!(false)) || (!(false)))) || ((_dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)))).IsProperSubsetOf(_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return (((false) ? (new BigNumber(884)) : (new BigNumber((_dafny.Seq.UnicodeFromString("xjmw")).length)))).multipliedBy(((true) ? (new BigNumber((_dafny.Seq.of(true, !(true))).length)) : (new BigNumber(-443))));
    };
    static fm4(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(!(true)),false);
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC5();
    };
    static fm9(globalState) {
      return _module.D2.create_DC6(_module.D2.create_DC6(_module.D2.create_DC6(_module.D2.create_DC4(_dafny.Seq.UnicodeFromString("bx"), true))));
    };
    static fm10(p0, p1, globalState) {
      return (_dafny.Set.fromElements(true)).Difference((_dafny.Set.fromElements(false)).Union(_dafny.Set.fromElements(false)));
    };
    static fm11(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(465),new BigNumber(64))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-896),new BigNumber(863)));
    };
    static fm12(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true, false, !(true))));
    };
    static fm14(p0, p1, p2, globalState) {
      if (!(false)) {
        return _module.D2.create_DC6(_module.D2.create_DC5());
      } else {
        return _module.D2.create_DC6(_module.D2.create_DC3(_dafny.Seq.of(_module.D0.create_DC1(true), _module.D0.create_DC1(false), _module.D0.create_DC1(false))));
      }
    };
    static fm17(p0, globalState) {
      return new _dafny.CodePoint('s'.codePointAt(0));
    };
    static fm18(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vn"), _dafny.Seq.UnicodeFromString("pyh")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_6_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("nt")));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("iql");
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.of(function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-443), new BigNumber(-121))) {
          let _7_v0 = _compr_3;
          if (((new BigNumber(-443)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(-121)))) {
            _coll3.add((_7_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length))).cardinality())));
          }
        }
        return _coll3;
      }(), _dafny.Set.fromElements(new BigNumber(917), new BigNumber(-461)), _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false, !(!(true)))).cardinality())), _dafny.Set.fromElements(new BigNumber(-841), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), function (_8_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), new BigNumber(36)), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length))))).length)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),!(false))).length)),_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("x"), new _dafny.CodePoint('j'.codePointAt(0))));
    };
    static fm21(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-222), new BigNumber(310))).cardinality()))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true, false, false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(false)))),new BigNumber((_dafny.Seq.UnicodeFromString("sp")).length))));
    };
    static fm22(globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(303)))).Union(_dafny.MultiSet.fromElements(new BigNumber(965)))).Union(_dafny.MultiSet.fromElements(new BigNumber(610)));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return _module.D5.create_DC13(((false) ? (new BigNumber(392)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ntyxfwen")).length))), _dafny.Set.fromElements(false), new _dafny.CodePoint('g'.codePointAt(0)), new BigNumber(-831), false);
    };
    static fm24(p0, p1, p2, p3, globalState) {
      if (false) {
        return (_dafny.MultiSet.fromElements(_module.D2.create_DC6(_module.D2.create_DC3(_dafny.Seq.of(_module.D0.create_DC1(false), _module.D0.create_DC1(true)))))).Intersect(_dafny.MultiSet.fromElements(_module.D2.create_DC6(_module.D2.create_DC5())));
      } else {
        return _dafny.MultiSet.fromElements(_module.D2.create_DC6(_module.D2.create_DC6(_module.D2.create_DC6(_module.D2.create_DC3(_dafny.Seq.of(_module.D0.create_DC1(true), _module.D0.create_DC1(!(false)), _module.D0.create_DC1(true), _module.D0.create_DC1(true)))))), _module.D2.create_DC6(_module.D2.create_DC3(_dafny.Seq.of(_module.D0.create_DC1(true), _module.D0.create_DC1(true)))));
      }
    };
    static fm25(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,false);
    };
    static fm26(globalState) {
      return _module.D7.create_DC18(_module.__default.safeDivisionInt(new BigNumber(698), new BigNumber((_dafny.Seq.UnicodeFromString("todmnh")).length)));
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(false, !(!(true)))).Difference(_dafny.Set.fromElements(false, !(false)))).Union(_dafny.Set.fromElements(false));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(!(false))).Union(_dafny.MultiSet.fromElements(false, !(true), false))).Intersect(((false) ? (_dafny.MultiSet.fromElements(!(!(true)))) : (_dafny.MultiSet.fromElements(false))));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(new BigNumber(121), new BigNumber(39))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(953),false)).length))),(new BigNumber(532)).isEqualTo(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("s")).length))).length))).length)));
    };
    static fm30(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.of(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, false, true, true, false))));
    };
    static fm31(globalState) {
      return (_dafny.Set.fromElements(new _dafny.CodePoint('l'.codePointAt(0)))).Union(((true) ? (_dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))) : (function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), function (_9_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })).Elements) {
          let _10_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), function (_9_i0) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), _10_v0)) {
            _coll4.add(_10_v0);
          }
        }
        return _coll4;
      }())));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.FromArray(((!(false)) ? (_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-625))), _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(454)), function (_11_i0) {
        return new BigNumber(580);
      })))) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(116), new BigNumber(990)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("sng")).length)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("iajxrdi")).length))), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pocfxd")).length),false)).length)))))))).Difference((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(973), new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false))).Elements) {
          let _12_v0 = _compr_5;
          if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false))).contains(_12_v0)) {
            _coll5.push([_12_v0,false]);
          }
        }
        return _coll5;
      }()).length), new BigNumber(59), new BigNumber(168)))).Intersect(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-520), new BigNumber(279)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(150),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("sq")).length))).length))).length), new BigNumber(789), new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-224), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Set.fromElements(false, true)).length))).cardinality())))));
    };
    static fm33(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-871),_module.D18.create_DC39(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sbpaa")).length),_module.D18.create_DC39(!(true))));
    };
    static fm34(globalState) {
      return function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(573),_module.D13.create_DC28(true, new _dafny.CodePoint('j'.codePointAt(0))))).Merge(function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), function (_13_i0) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_14_i1) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            })).length);
          })).Elements) {
            let _15_v0 = _compr_7;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), function (_13_i0) {
              return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_14_i1) {
                return new _dafny.CodePoint('w'.codePointAt(0));
              })).length);
            }), _15_v0)) {
              _coll7.push([(_15_v0).minus(new BigNumber(91)),_module.D13.create_DC28(false, new _dafny.CodePoint('o'.codePointAt(0)))]);
            }
          }
          return _coll7;
        }())).Keys.Elements) {
          let _16_v1 = _compr_6;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(573),_module.D13.create_DC28(true, new _dafny.CodePoint('j'.codePointAt(0))))).Merge(function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), function (_13_i0) {
              return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_14_i1) {
                return new _dafny.CodePoint('w'.codePointAt(0));
              })).length);
            })).Elements) {
              let _15_v0 = _compr_8;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), function (_13_i0) {
                return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_14_i1) {
                  return new _dafny.CodePoint('w'.codePointAt(0));
                })).length);
              }), _15_v0)) {
                _coll8.push([(_15_v0).minus(new BigNumber(91)),_module.D13.create_DC28(false, new _dafny.CodePoint('o'.codePointAt(0)))]);
              }
            }
            return _coll8;
          }())).contains(_16_v1)) {
            _coll6.add(_module.__default.safeDivisionInt(_16_v1, new BigNumber(316)));
          }
        }
        return _coll6;
      }();
    };
    static fm35(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))), new BigNumber((_dafny.Seq.of(true)).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(31), new BigNumber(724))).length)), function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(511), new BigNumber(620))) {
          let _17_v0 = _compr_9;
          if (((new BigNumber(511)).isLessThanOrEqualTo(_17_v0)) && ((_17_v0).isLessThan(new BigNumber(620)))) {
            _coll9.add(_module.__default.safeModuloInt(_17_v0, new BigNumber((_dafny.Seq.UnicodeFromString("oyikk")).length)));
          }
        }
        return _coll9;
      }());
    };
    static m0(p0, p1, globalState) {
      let r0 = false;
      let _18_v0;
      _18_v0 = new _dafny.CodePoint('q'.codePointAt(0));
      let _19_v1;
      _19_v1 = _module.D13.create_DC28(p0, _18_v0);
      let _source0 = _19_v1;
      if (_source0.is_DC27) {
        let _20___mcc_h0 = (_source0).cf43;
        let _21___mcc_h1 = (_source0).cf44;
        let _22___mcc_h2 = (_source0).cf45;
        let _23___mcc_h3 = (_source0).cf46;
        let _24_cf46 = _23___mcc_h3;
        let _25_cf45 = _22___mcc_h2;
        let _26_cf44 = _21___mcc_h1;
        let _27_cf43 = _20___mcc_h0;
        let _28_v2;
        let _nw0 = Array((new BigNumber(3)).toNumber()).fill(false);
        _28_v2 = _nw0;
        let _29_v3;
        _29_v3 = _module.D11.create_DC24(_28_v2, p0, p0);
        let _30_v4;
        _30_v4 = _dafny.Set.fromElements(_29_v3);
        let _31_v5;
        _31_v5 = _28_v2;
        let _32_v6;
        _32_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _33_v7;
        _33_v7 = _dafny.Seq.UnicodeFromString("gvt");
        let _34_v8;
        _34_v8 = _module.D18.create_DC38(_30_v4);
        let _35_v9;
        let _nw1 = Array((new BigNumber(20)).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = (_30_v4).Difference(_30_v4);
        _nw1[(_dafny.ONE).toNumber()] = ((p0) ? (_30_v4) : (_30_v4));
        _nw1[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements(_29_v3)).Intersect(_30_v4);
        _nw1[(new BigNumber(3)).toNumber()] = (_dafny.Set.fromElements(_module.D11.create_DC24((_31_v5), p0, p0))).Difference(_30_v4);
        _nw1[(new BigNumber(4)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(5)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(6)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(7)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_module.D11.create_DC24(_28_v2, p0, p0));
        _nw1[(new BigNumber(9)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(10)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(11)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(12)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(13)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(14)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(15)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(16)).toNumber()] = _30_v4;
        _nw1[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(_module.D11.create_DC24(_28_v2, (((_32_v6).contains((_26_cf44).f0)) ? ((_32_v6).get((_26_cf44).f0)) : (true)), p0), _29_v3, _module.D11.create_DC24(_28_v2, p0, (_26_cf44).fm6(p0, _33_v7, _33_v7, _24_cf46, globalState)));
        _nw1[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements(_29_v3);
        _nw1[(new BigNumber(19)).toNumber()] = ((_34_v8).dtor_cf58).Union(_30_v4);
        _35_v9 = _nw1;
        _35_v9 = _35_v9;
        let _36_v10;
        _36_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,p0),p1);
        let _37_v11;
        _37_v11 = _dafny.Map.Empty.slice().updateUnsafe((_26_cf44).f0,(_26_cf44).f0);
        let _38_v12;
        let _nw2 = new _module.C5();
        _nw2.__ctor(_25_cf45, (_dafny.ZERO).minus((_module.D5.create_DC12(p1, (((_36_v10).contains(_dafny.Map.Empty.slice().updateUnsafe(p0,p0))) ? ((_36_v10).get(_dafny.Map.Empty.slice().updateUnsafe(p0,p0))) : (p1)))).dtor_cf15), new BigNumber((_37_v11).length));
        _38_v12 = _nw2;
        let _39_v13;
        let _nw3 = new _module.C2();
        _nw3.__ctor(new BigNumber((_33_v7).length));
        _39_v13 = _nw3;
        let _rhs0 = _38_v12;
        let _rhs1 = _39_v13;
        let _rhs2 = (_38_v12).f2;
        let _rhs3 = p0;
        let _rhs4 = ((_38_v12).f2).isEqualTo(_module.__default.safeModuloInt(p1, p1));
        _38_v12 = _rhs0;
        _39_v13 = _rhs1;
        _24_cf46 = _rhs2;
        r0 = _rhs3;
        r0 = _rhs4;
        let _40_v14;
        _40_v14 = _dafny.Set.fromElements((_38_v12).f2);
        let _41_v15;
        _41_v15 = _module.D3.create_DC8(p0, (_38_v12).f2);
        let _42_v16;
        _42_v16 = _module.D7.create_DC17(p0, p0, p0);
        let _43_v17;
        _43_v17 = _dafny.Set.fromElements(_42_v16);
        let _rhs5 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_38_v12).f2), new BigNumber((_40_v14).length));
        let _rhs6 = !((_38_v12).fm6(p0, _33_v7, (_module.D2.create_DC4(_33_v7, p0)).dtor_cf4, (_41_v15).dtor_cf9, globalState));
        let _rhs7 = (_38_v12).f2;
        let _rhs8 = !((_dafny.MultiSet.fromElements(_43_v17, _43_v17)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_43_v17)));
        _24_cf46 = _rhs5;
        r0 = _rhs6;
        _24_cf46 = _rhs7;
        r0 = _rhs8;
        let _44_v18;
        let _nw4 = new _module.C4();
        _nw4.__ctor(_module.__default.fm3(_dafny.Set.fromElements(p0), (_38_v12).f1, new BigNumber(-553), new BigNumber(258), globalState));
        _44_v18 = _nw4;
      } else if (_source0.is_DC28) {
        let _45___mcc_h4 = (_source0).cf47;
        let _46___mcc_h5 = (_source0).cf48;
        let _47_cf48 = _46___mcc_h5;
        let _48_cf47 = _45___mcc_h4;
        if (((true) ? (_48_cf47) : (_48_cf47))) {
          let _49_v19;
          let _nw5 = new _module.C5();
          _nw5.__ctor(_47_cf48, new BigNumber(515), new BigNumber(237));
          _49_v19 = _nw5;
          let _50_v20;
          _50_v20 = _dafny.Map.Empty.slice().updateUnsafe(p1,_49_v19);
          let _51_v21;
          _51_v21 = _dafny.Map.Empty.slice().updateUnsafe(_18_v0,(_50_v20).Merge(_50_v20));
          _51_v21 = (_51_v21).Merge(_dafny.Map.Empty.slice().updateUnsafe(_47_cf48,_50_v20));
          let _52_v22;
          _52_v22 = new BigNumber(118);
          _52_v22 = p1;
          let _53_v23;
          _53_v23 = _dafny.MultiSet.fromElements(_52_v22, (_49_v19).f0, _52_v22, (_49_v19).f0);
          let _54_v24;
          _54_v24 = _dafny.MultiSet.fromElements((_53_v23).Difference((_53_v23).update(p1, _module.__default.abs(p1))), _53_v23, (_53_v23).Intersect(_53_v23));
          let _55_v25;
          _55_v25 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(p1));
          let _56_v26;
          _56_v26 = _dafny.Map.Empty.slice().updateUnsafe(_49_v19,p1);
          let _57_v27;
          _57_v27 = _dafny.Seq.of(p1);
          let _58_v28;
          _58_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_57_v27).length),_dafny.Seq.of(_18_v0, _47_cf48));
          _54_v24 = _module.__default.fm32((_55_v25).Merge(_55_v25), new BigNumber((_56_v26).length), false, _58_v28, globalState);
          let _59_v29;
          _59_v29 = _dafny.Seq.UnicodeFromString("wkht");
          let _60_v30;
          _60_v30 = _module.D2.create_DC4(_59_v29, _48_cf47);
          let _61_v31;
          _61_v31 = _dafny.Map.Empty.slice().updateUnsafe(_60_v30,p1);
          let _62_v32;
          _62_v32 = _module.D18.create_DC39(p0);
          let _pat_let_tv0 = _48_cf47;
          let _63_v33;
          _63_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qhqkkmfj")).length),function (_pat_let0_0) {
            return function (_64_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_65_dt__update_hcf59_h0) {
                  return _module.D18.create_DC39(_65_dt__update_hcf59_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_62_v32));
          let _66_v35;
          _66_v35 = _dafny.Seq.of(_module.__default.fm33((_49_v19).f0, _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC4(_59_v29, true),_52_v22), globalState), _module.__default.fm33(_52_v22, _61_v31, globalState), (_63_v33).update((_dafny.ZERO).minus((_49_v19).f0), _62_v32), _module.__default.fm33((_49_v19).f0, _61_v31, globalState), function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of _dafny.IntegerRange(new BigNumber(648), new BigNumber(747))) {
              let _67_v34 = _compr_10;
              if (((new BigNumber(648)).isLessThanOrEqualTo(_67_v34)) && ((_67_v34).isLessThan(new BigNumber(747)))) {
                _coll10.push([(_67_v34).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_59_v29)).cardinality())),_62_v32]);
              }
            }
            return _coll10;
          }());
          _52_v22 = new BigNumber((_66_v35).length);
          let _68_v36;
          let _nw6 = Array((new BigNumber(12)).toNumber());
          _68_v36 = _nw6;
          let _rhs9 = p0;
          let _rhs10 = _68_v36;
          _48_cf47 = _rhs9;
          _68_v36 = _rhs10;
        } else {
          let _69_v37;
          _69_v37 = new BigNumber(223);
          _69_v37 = _69_v37;
          let _70_v38;
          _70_v38 = _dafny.Map.Empty.slice().updateUnsafe(_48_cf47,_module.__default.fm2(p0, p1, p0, _48_cf47, globalState));
          let _71_v39;
          _71_v39 = _70_v38;
          let _72_v40;
          _72_v40 = _dafny.Map.Empty.slice().updateUnsafe(_71_v39,_48_cf47);
          _72_v40 = (_72_v40).update(_71_v39, p0);
          let _73_v41;
          let _nw7 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _73_v41 = _nw7;
          let _index0 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_73_v41).length));
          (_73_v41)[_index0] = _69_v37;
          let _index1 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_73_v41).length));
          (_73_v41)[_index1] = new BigNumber((_module.__default.fm22(globalState)).cardinality());
          _48_cf47 = p0;
          let _74_v42;
          _74_v42 = _dafny.Seq.of((_73_v41)[_module.__default.safeIndex(new BigNumber(367), new BigNumber((_73_v41).length))]);
          let _index2 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_73_v41).length));
          (_73_v41)[_index2] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_74_v42, _74_v42)).length), _module.__default.safeModuloInt(_69_v37, p1));
        }
        let _75_v43;
        _75_v43 = new BigNumber(-188);
        let _76_v44;
        _76_v44 = _dafny.Set.fromElements(!(_48_cf47));
        let _77_v45;
        let _init0 = ((_78_cf47) => function (_79_i0) {
          return _78_cf47;
        })(_48_cf47);
        let _nw8 = Array((new BigNumber(24)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw8.length); _i0_0++) {
          _nw8[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _77_v45 = _nw8;
        let _80_v46;
        _80_v46 = _module.D11.create_DC24(_77_v45, _48_cf47, p0);
        let _81_v47;
        _81_v47 = _dafny.Set.fromElements(_80_v46, function (_pat_let2_0) {
          return function (_82_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_83_dt__update_hcf39_h0) {
                return _module.D11.create_DC24((_82_dt__update__tmp_h1).dtor_cf38, _83_dt__update_hcf39_h0, (_82_dt__update__tmp_h1).dtor_cf40);
              }(_pat_let3_0);
            }(false);
          }(_pat_let2_0);
        }(_80_v46));
        let _84_v48;
        _84_v48 = _module.D18.create_DC38(_81_v47);
        let _85_v49;
        _85_v49 = _dafny.MultiSet.fromElements(_module.D18.create_DC38(_dafny.Set.fromElements(_80_v46)), _84_v48);
        let _86_v50;
        _86_v50 = _module.D19.create_DC41(_85_v49);
        _75_v43 = _module.__default.fm3(_76_v44, _47_cf48, ((_48_cf47) ? (new BigNumber(((_86_v50).dtor_cf62).cardinality())) : (_75_v43)), _75_v43, globalState);
        if (!(!((_75_v43).isEqualTo(new BigNumber(182))))) {
          _75_v43 = (new BigNumber(641)).minus(p1);
          _77_v45 = _77_v45;
          let _87_v51;
          _87_v51 = _dafny.Seq.of(p0);
          _48_cf47 = _module.__default.fm2((_76_v44).IsDisjointFrom(_76_v44), _module.__default.safeDivisionInt(_75_v43, (_dafny.ZERO).minus(new BigNumber((_87_v51).length))), p0, true, globalState);
          let _pat_let_tv1 = p0;
          _19_v1 = ((_48_cf47) ? (function (_pat_let4_0) {
            return function (_88_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_89_dt__update_hcf47_h0) {
                  return _module.D13.create_DC28(_89_dt__update_hcf47_h0, (_88_dt__update__tmp_h2).dtor_cf48);
                }(_pat_let5_0);
              }(_pat_let_tv1);
            }(_pat_let4_0);
          }(_19_v1)) : (_19_v1));
          let _90_v52;
          _90_v52 = _dafny.Set.fromElements(p1);
          let _91_v53;
          let _nw9 = new _module.C1();
          _nw9.__ctor(p1);
          _91_v53 = _nw9;
          let _92_v54;
          _92_v54 = _dafny.MultiSet.fromElements(_91_v53);
          let _93_v55;
          _93_v55 = _dafny.Seq.of(_90_v52, _dafny.Set.fromElements(_module.__default.fm3(_76_v44, new _dafny.CodePoint('d'.codePointAt(0)), new BigNumber(872), new BigNumber((_92_v54).cardinality()), globalState)));
          let _94_v56;
          _94_v56 = _dafny.Seq.of((_93_v55)[_module.__default.safeIndex(p1, new BigNumber((_93_v55).length))]);
          _90_v52 = ((_93_v55)[_module.__default.safeIndex(new BigNumber(536), new BigNumber((_93_v55).length))]).Intersect(_90_v52);
        } else {
          let _95_v57;
          let _nw10 = new _module.C1();
          _nw10.__ctor((p1).multipliedBy(p1));
          _95_v57 = _nw10;
          let _96_v58;
          _96_v58 = _dafny.Seq.UnicodeFromString("pmdqlx");
          let _97_v59;
          _97_v59 = _module.D2.create_DC4(_96_v58, p0);
          let _98_v60;
          _98_v60 = _dafny.Map.Empty.slice().updateUnsafe((_97_v59).dtor_cf4,p0);
          let _99_v61;
          _99_v61 = _dafny.Map.Empty.slice().updateUnsafe(_48_cf47,(_95_v57).fm15(_98_v60, _48_cf47, _48_cf47, globalState));
          let _rhs11 = _77_v45;
          let _rhs12 = _dafny.Seq.IsProperPrefixOf(_96_v58, _96_v58);
          let _rhs13 = _95_v57;
          let _rhs14 = ((p0) ? (((((_99_v61).contains(p0)) ? ((_99_v61).get(p0)) : (_75_v43))).multipliedBy(_75_v43)) : (new BigNumber(-427)));
          _77_v45 = _rhs11;
          r0 = _rhs12;
          _95_v57 = _rhs13;
          _75_v43 = _rhs14;
          let _100_v62;
          let _nw11 = new _module.C5();
          _nw11.__ctor((_96_v58)[_module.__default.safeIndex(p1, new BigNumber((_96_v58).length))], p1, _module.__default.safeModuloInt(_75_v43, _75_v43));
          _100_v62 = _nw11;
          _100_v62 = _100_v62;
          let _101_v63;
          let _nw12 = new _module.C1();
          _nw12.__ctor((_100_v62).f0);
          _101_v63 = _nw12;
          let _102_v64;
          _102_v64 = _dafny.Set.fromElements((_dafny.ZERO).minus(_75_v43), p1, p1, new BigNumber(308));
          let _103_v65;
          _103_v65 = _dafny.MultiSet.fromElements((((_99_v61).contains(_48_cf47)) ? ((_99_v61).get(_48_cf47)) : ((_100_v62).f0)));
          let _104_v66;
          let _nw13 = Array((new BigNumber(15)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = p1;
          _nw13[(_dafny.ONE).toNumber()] = _75_v43;
          _nw13[(new BigNumber(2)).toNumber()] = new BigNumber((_102_v64).length);
          _nw13[(new BigNumber(3)).toNumber()] = new BigNumber(132);
          _nw13[(new BigNumber(4)).toNumber()] = ((_100_v62).f0).plus((_100_v62).f0);
          _nw13[(new BigNumber(5)).toNumber()] = p1;
          _nw13[(new BigNumber(6)).toNumber()] = new BigNumber(854);
          _nw13[(new BigNumber(7)).toNumber()] = new BigNumber(842);
          _nw13[(new BigNumber(8)).toNumber()] = p1;
          _nw13[(new BigNumber(9)).toNumber()] = p1;
          _nw13[(new BigNumber(10)).toNumber()] = new BigNumber(686);
          _nw13[(new BigNumber(11)).toNumber()] = (_100_v62).f0;
          _nw13[(new BigNumber(12)).toNumber()] = p1;
          _nw13[(new BigNumber(13)).toNumber()] = new BigNumber((_103_v65).cardinality());
          _nw13[(new BigNumber(14)).toNumber()] = p1;
          _104_v66 = _nw13;
          let _105_v67;
          _105_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(667),p0);
          let _106_v68;
          _106_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_75_v43, new BigNumber((_105_v67).length)),_48_cf47);
          let _rhs15 = p1;
          let _rhs16 = (((_106_v68).contains(p1)) ? ((_106_v68).get(p1)) : (_48_cf47));
          let _rhs17 = (_75_v43).isLessThan(p1);
          let _rhs18 = _104_v66;
          _75_v43 = _rhs15;
          r0 = _rhs16;
          _48_cf47 = _rhs17;
          _104_v66 = _rhs18;
          let _107_v69;
          _107_v69 = _dafny.Seq.of(p0, p0, p0);
          let _108_v70;
          _108_v70 = _dafny.Map.Empty.slice().updateUnsafe(_107_v69,_48_cf47);
          _108_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_107_v69, _107_v69),p0);
        }
        let _index3 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_77_v45).length));
        (_77_v45)[_index3] = p0;
        let _109_v71;
        _109_v71 = _dafny.MultiSet.fromElements(_48_cf47, p0, _48_cf47);
        let _index4 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_77_v45).length));
        (_77_v45)[_index4] = ((((_109_v71).contains(false)) ? ((_109_v71).get(false)) : (_75_v43))).isLessThan(_module.__default.fm3(_dafny.Set.fromElements(_48_cf47, _48_cf47, _48_cf47), _18_v0, p1, p1, globalState));
      } else if (_source0.is_DC29) {
        let _110___mcc_h6 = (_source0).cf49;
        let _111___mcc_h7 = (_source0).cf50;
        let _112_cf50 = _111___mcc_h7;
        let _113_cf49 = _110___mcc_h6;
        let _114_v72;
        let _nw14 = Array((new BigNumber(23)).toNumber()).fill([]);
        _114_v72 = _nw14;
        let _nw15 = Array((new BigNumber(13)).toNumber()).fill([]);
        _114_v72 = _nw15;
        let _115_v73;
        _115_v73 = _dafny.Set.fromElements(_112_cf50, p0, p0, !(_112_cf50), _112_cf50);
        let _116_v74;
        let _nw16 = new _module.C5();
        _nw16.__ctor(new _dafny.CodePoint('c'.codePointAt(0)), _module.__default.fm3(_115_v73, _18_v0, p1, p1, globalState), (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, p1))));
        _116_v74 = _nw16;
        _116_v74 = ((_112_cf50) ? (_116_v74) : (_116_v74));
        let _117_v75;
        let _nw17 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _117_v75 = _nw17;
        let _118_v76;
        _118_v76 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_116_v74).f0);
        let _index5 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_117_v75).length));
        (_117_v75)[_index5] = (_118_v76).Merge(_118_v76);
        let _index6 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_117_v75).length));
        (_117_v75)[_index6] = (((_118_v76).update(p1, (_116_v74).f0)).Merge(_118_v76)).Merge((_118_v76).Merge(_118_v76));
        let _119_v77;
        let _init1 = ((_120_cf50) => function (_121_i1) {
          return _120_cf50;
        })(_112_cf50);
        let _nw18 = Array((new BigNumber(23)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw18.length); _i0_1++) {
          _nw18[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _119_v77 = _nw18;
        _119_v77 = _119_v77;
      } else if (_source0.is_DC26) {
        let _122___mcc_h8 = (_source0).cf42;
        let _123_cf42 = _122___mcc_h8;
        let _124_v78;
        let _nw19 = new _module.C1();
        _nw19.__ctor(p1);
        _124_v78 = _nw19;
        let _125_v79;
        _125_v79 = _dafny.Seq.of(_124_v78);
        let _126_v80;
        _126_v80 = _dafny.Seq.of(_125_v79);
        _125_v79 = (_126_v80)[_module.__default.safeIndex(p1, new BigNumber((_126_v80).length))];
        let _127_v81;
        _127_v81 = new BigNumber(70);
        let _128_v82;
        _128_v82 = _dafny.Set.fromElements(false, p0, p0);
        _127_v81 = (_127_v81).minus(_module.__default.fm3(_128_v82, _18_v0, _127_v81, p1, globalState));
        r0 = p0;
        _127_v81 = _127_v81;
      } else {
        let _129___mcc_h9 = (_source0).cf51;
        let _130_cf51 = _129___mcc_h9;
        r0 = !(p0);
        let _131_v83;
        let _init2 = ((_132_p0) => function (_133_i2) {
          return !(_132_p0);
        })(p0);
        let _nw20 = Array((new BigNumber(15)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
          _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _131_v83 = _nw20;
        let _134_v84;
        _134_v84 = _dafny.MultiSet.fromElements(p1);
        let _index7 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_131_v83).length));
        (_131_v83)[_index7] = (_134_v84).IsSubsetOf(_134_v84);
        let _index8 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_131_v83).length));
        (_131_v83)[_index8] = p0;
        let _135_v85;
        let _nw21 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _135_v85 = _nw21;
        let _136_v86;
        _136_v86 = _dafny.Set.fromElements(p0);
        let _137_v87;
        let _nw22 = new _module.C1();
        _nw22.__ctor(p1);
        _137_v87 = _nw22;
        let _138_v88;
        _138_v88 = _dafny.Seq.of(_137_v87, _137_v87);
        let _139_v89;
        _139_v89 = _dafny.Seq.of(_131_v83);
        let _140_v90;
        let _nw23 = Array((new BigNumber(17)).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw23[(_dafny.ONE).toNumber()] = p1;
        _nw23[(new BigNumber(2)).toNumber()] = p1;
        _nw23[(new BigNumber(3)).toNumber()] = p1;
        _nw23[(new BigNumber(4)).toNumber()] = p1;
        _nw23[(new BigNumber(5)).toNumber()] = p1;
        _nw23[(new BigNumber(6)).toNumber()] = p1;
        _nw23[(new BigNumber(7)).toNumber()] = new BigNumber(122);
        _nw23[(new BigNumber(8)).toNumber()] = p1;
        _nw23[(new BigNumber(9)).toNumber()] = p1;
        _nw23[(new BigNumber(10)).toNumber()] = p1;
        _nw23[(new BigNumber(11)).toNumber()] = new BigNumber((_136_v86).length);
        _nw23[(new BigNumber(12)).toNumber()] = new BigNumber((_138_v88).length);
        _nw23[(new BigNumber(13)).toNumber()] = p1;
        _nw23[(new BigNumber(14)).toNumber()] = p1;
        _nw23[(new BigNumber(15)).toNumber()] = p1;
        _nw23[(new BigNumber(16)).toNumber()] = new BigNumber((_139_v89).length);
        _140_v90 = _nw23;
        let _141_v91;
        _141_v91 = _dafny.Map.Empty.slice().updateUnsafe((_131_v83)[_module.__default.safeIndex(new BigNumber(64), new BigNumber((_131_v83).length))],_140_v90);
        let _index9 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_135_v85).length));
        (_135_v85)[_index9] = _141_v91;
        let _index10 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_135_v85).length));
        (_135_v85)[_index10] = _141_v91;
        let _142_v92;
        _142_v92 = new BigNumber(601);
        let _143_v93;
        _143_v93 = _dafny.Seq.of(p1, _142_v92);
        let _144_v94;
        _144_v94 = _dafny.Seq.of(_134_v84);
        _142_v92 = new BigNumber(((((_131_v83)[_module.__default.safeIndex(new BigNumber(64), new BigNumber((_131_v83).length))]) ? (_dafny.MultiSet.FromArray(_143_v93)) : ((_144_v94)[_module.__default.safeIndex(_142_v92, new BigNumber((_144_v94).length))]))).cardinality());
      }
      let _145_v95;
      _145_v95 = new BigNumber(898);
      let _rhs19 = p1;
      let _rhs20 = p0;
      _145_v95 = _rhs19;
      r0 = _rhs20;
      let _146_v96;
      _146_v96 = _dafny.Seq.UnicodeFromString("dwej");
      let _147_v97;
      _147_v97 = _dafny.Seq.of(_145_v95);
      let _148_v98;
      _148_v98 = _dafny.Seq.of(p0, _module.__default.fm2(p0, new BigNumber((_147_v97).length), p0, p0, globalState));
      let _149_v99;
      let _nw24 = new _module.C5();
      _nw24.__ctor(new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber((_148_v98).length), _145_v95);
      _149_v99 = _nw24;
      let _150_v100;
      _150_v100 = _dafny.Seq.of(_149_v99);
      let _151_v101;
      _151_v101 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_146_v96,!(p0))).length), p1, p1, new BigNumber((_150_v100).length));
      if ((_151_v101).IsProperSubsetOf(_module.__default.fm34(globalState))) {
        if (p0) {
          let _152_v102;
          _152_v102 = _151_v101;
          let _153_v103;
          _153_v103 = _dafny.Seq.of(_152_v102, _151_v101, _151_v101);
          let _154_v104;
          _154_v104 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          _153_v103 = _dafny.Seq.Concat(_153_v103, _module.__default.fm35((((_154_v104).contains(p0)) ? ((_154_v104).get(p0)) : (new BigNumber((_146_v96).length))), (_149_v99).f0, p0, globalState));
          r0 = p0;
          let _155_v105;
          let _nw25 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _155_v105 = _nw25;
          let _156_v106;
          let _nw26 = new _module.C0();
          _nw26.__ctor(_18_v0);
          _156_v106 = _nw26;
          let _157_v107;
          _157_v107 = _dafny.Map.Empty.slice().updateUnsafe(_156_v106,_dafny.Seq.Create(_module.__default.abs(new BigNumber(289)), ((_158_v99) => function (_159_i3) {
            return (_158_v99).f0;
          })(_149_v99)));
          let _160_v108;
          _160_v108 = _dafny.Map.Empty.slice().updateUnsafe(_157_v107,_145_v95);
          let _index11 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_155_v105).length));
          (_155_v105)[_index11] = (((_160_v108).contains(_157_v107)) ? ((_160_v108).get(_157_v107)) : (new BigNumber(376)));
          let _161_v109;
          let _init3 = ((_162_p0) => function (_163_i4) {
            return _162_p0;
          })(p0);
          let _nw27 = Array((new BigNumber(22)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw27.length); _i0_3++) {
            _nw27[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _161_v109 = _nw27;
          let _164_v110;
          _164_v110 = _dafny.Set.fromElements(_161_v109);
          let _index12 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_155_v105).length));
          (_155_v105)[_index12] = ((_149_v99).f0).minus(new BigNumber((_164_v110).length));
          let _165_v111;
          let _nw28 = new _module.C3();
          _nw28.__ctor((_149_v99).f0);
          _165_v111 = _nw28;
          let _166_v112;
          _166_v112 = _dafny.Map.Empty.slice().updateUnsafe((_155_v105)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_155_v105).length))],_dafny.Set.fromElements(_165_v111));
          _145_v95 = _module.__default.safeModuloInt(new BigNumber((_166_v112).length), new BigNumber(223));
          let _index13 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_155_v105).length));
          (_155_v105)[_index13] = (_149_v99).f0;
        } else {
          _149_v99 = _149_v99;
          let _167_v113;
          _167_v113 = _dafny.Set.fromElements(_149_v99);
          let _rhs21 = _146_v96;
          let _rhs22 = ((false) ? (p0) : (p0));
          let _rhs23 = (new BigNumber(((_167_v113).Union(_167_v113)).length)).minus((_149_v99).f0);
          _146_v96 = _rhs21;
          r0 = _rhs22;
          _145_v95 = _rhs23;
          let _168_v114;
          let _nw29 = Array((new BigNumber(3)).toNumber());
          _168_v114 = _nw29;
          let _index14 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_168_v114).length));
          (_168_v114)[_index14] = _149_v99;
          let _index15 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_168_v114).length));
          (_168_v114)[_index15] = _149_v99;
          _145_v95 = (_dafny.ZERO).minus(((_149_v99).f0).multipliedBy(((_149_v99).f0).minus((_dafny.ZERO).minus((_149_v99).f0))));
          _145_v95 = (new BigNumber(-90)).minus((_149_v99).f0);
        }
        let _169_v115;
        _169_v115 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _170_v116;
        _170_v116 = _dafny.Seq.of(_169_v115, _169_v115, _dafny.Map.Empty.slice().updateUnsafe(p0,p0), (_169_v115).Merge(_169_v115));
        _170_v116 = _170_v116;
        let _171_v117;
        let _nw30 = new _module.C3();
        _nw30.__ctor(new BigNumber((_148_v98).length));
        _171_v117 = _nw30;
        let _172_v118;
        _172_v118 = _dafny.Set.fromElements(p0, p0);
        _145_v95 = (new BigNumber(720)).minus((_module.__default.fm3(_172_v118, _18_v0, new BigNumber((function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of _dafny.IntegerRange(new BigNumber(766), new BigNumber(816))) {
            let _173_v119 = _compr_11;
            if (((new BigNumber(766)).isLessThanOrEqualTo(_173_v119)) && ((_173_v119).isLessThan(new BigNumber(816)))) {
              _coll11.add(_module.__default.safeModuloInt(_173_v119, new BigNumber(918)));
            }
          }
          return _coll11;
        }()).length), (_149_v99).f0, globalState)).minus((_149_v99).f0));
        let _174_v120;
        let _nw31 = new _module.C0();
        _nw31.__ctor(new _dafny.CodePoint('m'.codePointAt(0)));
        _174_v120 = _nw31;
      } else {
        let _175_v121;
        let _nw32 = Array((new BigNumber(16)).toNumber());
        _175_v121 = _nw32;
        let _176_v122;
        _176_v122 = _module.D3.create_DC9(new BigNumber(745), _175_v121, p0);
        let _177_v123;
        _177_v123 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _178_v124;
        let _nw33 = Array((new BigNumber(26)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = p0;
        _nw33[(_dafny.ONE).toNumber()] = p0;
        _nw33[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsPrefixOf(_146_v96, _146_v96);
        _nw33[(new BigNumber(3)).toNumber()] = false;
        _nw33[(new BigNumber(4)).toNumber()] = p0;
        _nw33[(new BigNumber(5)).toNumber()] = p0;
        _nw33[(new BigNumber(6)).toNumber()] = p0;
        _nw33[(new BigNumber(7)).toNumber()] = p0;
        _nw33[(new BigNumber(8)).toNumber()] = p0;
        _nw33[(new BigNumber(9)).toNumber()] = p0;
        _nw33[(new BigNumber(10)).toNumber()] = p0;
        _nw33[(new BigNumber(11)).toNumber()] = ((p0) ? (p0) : (true));
        _nw33[(new BigNumber(12)).toNumber()] = p0;
        _nw33[(new BigNumber(13)).toNumber()] = !_dafny.Seq.contains(_146_v96, _18_v0);
        _nw33[(new BigNumber(14)).toNumber()] = !(!(p0)) || (p0);
        _nw33[(new BigNumber(15)).toNumber()] = p0;
        _nw33[(new BigNumber(16)).toNumber()] = p0;
        _nw33[(new BigNumber(17)).toNumber()] = p0;
        _nw33[(new BigNumber(18)).toNumber()] = true;
        _nw33[(new BigNumber(19)).toNumber()] = (_176_v122).dtor_cf12;
        _nw33[(new BigNumber(20)).toNumber()] = ((p0) ? (true) : ((((_177_v123).contains(p0)) ? ((_177_v123).get(p0)) : (p0))));
        _nw33[(new BigNumber(21)).toNumber()] = p0;
        _nw33[(new BigNumber(22)).toNumber()] = ((p0) ? (p0) : (p0));
        _nw33[(new BigNumber(23)).toNumber()] = p0;
        _nw33[(new BigNumber(24)).toNumber()] = p0;
        _nw33[(new BigNumber(25)).toNumber()] = !(p0);
        _178_v124 = _nw33;
        _178_v124 = _178_v124;
        if (p0) {
          let _179_v125;
          _179_v125 = _dafny.Map.Empty.slice().updateUnsafe(_18_v0,_145_v95);
          let _180_v126;
          _180_v126 = _dafny.MultiSet.fromElements(_179_v125);
          _180_v126 = _180_v126;
          let _181_v127;
          let _nw34 = new _module.C4();
          _nw34.__ctor(new BigNumber((_148_v98).length));
          _181_v127 = _nw34;
          let _182_v128;
          _182_v128 = _dafny.Seq.of(_181_v127);
          let _183_v129;
          _183_v129 = _dafny.Seq.of(_182_v128);
          r0 = _dafny.Seq.IsProperPrefixOf(((p0) ? ((_183_v129)[_module.__default.safeIndex(p1, new BigNumber((_183_v129).length))]) : (_dafny.Seq.of(_181_v127))), _dafny.Seq.update(_182_v128, _module.__default.safeIndex((_149_v99).f0, new BigNumber((_182_v128).length)), _181_v127));
          let _184_v130;
          let _nw35 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _184_v130 = _nw35;
          let _185_v131;
          let _nw36 = new _module.C4();
          _nw36.__ctor((_149_v99).f0);
          _185_v131 = _nw36;
          let _186_v132;
          _186_v132 = _module.D10.create_DC21(_185_v131);
          let _187_v133;
          _187_v133 = _dafny.Map.Empty.slice().updateUnsafe(_181_v127,(_186_v132).dtor_cf33);
          let _188_v134;
          let _nw37 = Array((new BigNumber(18)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = ((p0) ? (_185_v131) : (_185_v131));
          _nw37[(_dafny.ONE).toNumber()] = _185_v131;
          _nw37[(new BigNumber(2)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(3)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(4)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(5)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(6)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(7)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(8)).toNumber()] = (((_187_v133).contains(_181_v127)) ? ((_187_v133).get(_181_v127)) : (_185_v131));
          _nw37[(new BigNumber(9)).toNumber()] = ((p0) ? (_185_v131) : (_185_v131));
          _nw37[(new BigNumber(10)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(11)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(12)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(13)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(14)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(15)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(16)).toNumber()] = _185_v131;
          _nw37[(new BigNumber(17)).toNumber()] = _185_v131;
          _188_v134 = _nw37;
          let _index16 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_188_v134).length));
          (_188_v134)[_index16] = _185_v131;
          let _189_v135;
          _189_v135 = _dafny.Map.Empty.slice().updateUnsafe(_175_v121,p0);
          let _index17 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_188_v134).length));
          let _rhs24 = _184_v130;
          let _rhs25 = _module.__default.fm17(new BigNumber(((_189_v135).Merge(_189_v135)).length), globalState);
          let _rhs26 = (_149_v99).f0;
          let _rhs27 = _185_v131;
          let _lhs0 = _188_v134;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_188_v134).length));
          _184_v130 = _rhs24;
          _18_v0 = _rhs25;
          _145_v95 = _rhs26;
          _lhs0[_lhs1] = _rhs27;
          let _index18 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_178_v124).length));
          (_178_v124)[_index18] = !(!(p0));
          let _190_v136;
          _190_v136 = _dafny.Map.Empty.slice().updateUnsafe(_149_v99,(_149_v99).f0);
          let _index19 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_178_v124).length));
          (_178_v124)[_index19] = (_190_v136).equals(_190_v136);
          r0 = !(!(((p0) ? (p0) : (p0))));
        } else {
          _18_v0 = _18_v0;
          let _191_v137;
          _191_v137 = _dafny.Set.fromElements((_149_v99).fm6(p0, _dafny.Seq.UnicodeFromString("cqrcjrmk"), _146_v96, (_149_v99).f0, globalState), p0, p0, p0, p0);
          _145_v95 = _module.__default.fm3(_191_v137, _18_v0, _module.__default.safeModuloInt(p1, (_149_v99).f0), p1, globalState);
          r0 = p0;
          _145_v95 = (_149_v99).f0;
          let _192_v139;
          let _init4 = ((_193_v99) => function (_194_i5) {
            return function () {
              let _coll12 = new _dafny.Set();
              for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe((_193_v99).f0,true)).Keys.Elements) {
                let _195_v138 = _compr_12;
                if ((_dafny.Map.Empty.slice().updateUnsafe((_193_v99).f0,true)).contains(_195_v138)) {
                  _coll12.add((_195_v138).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)));
                }
              }
              return _coll12;
            }();
          })(_149_v99);
          let _nw38 = Array((new BigNumber(21)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw38.length); _i0_4++) {
            _nw38[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _192_v139 = _nw38;
          let _index20 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_192_v139).length));
          (_192_v139)[_index20] = _151_v101;
          let _196_v140;
          _196_v140 = _module.D16.create_DC35((_149_v99).f0);
          let _index21 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_192_v139).length));
          let _rhs28 = _module.__default.fm34(globalState);
          let _rhs29 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_197_v0) => function (_198_i6) {
            return _197_v0;
          })(_18_v0)), _module.__default.safeIndex(new BigNumber(323), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_199_v0) => function (_200_i6) {
            return _199_v0;
          })(_18_v0))).length)), (_146_v96)[_module.__default.safeIndex((_dafny.ZERO).minus((_196_v140).dtor_cf55), new BigNumber((_146_v96).length))]);
          let _lhs2 = _192_v139;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_192_v139).length));
          _lhs2[_lhs3] = _rhs28;
          _146_v96 = _rhs29;
        }
        let _201_v141;
        _201_v141 = _module.D3.create_DC8(p0, (_149_v99).f0);
        let _202_v142;
        let _nw39 = new _module.C1();
        _nw39.__ctor((_149_v99).f0);
        _202_v142 = _nw39;
        let _203_v143;
        _203_v143 = _dafny.Seq.of(_202_v142);
        let _204_v144;
        _204_v144 = _dafny.MultiSet.fromElements(_202_v142, _202_v142, (_203_v143)[_module.__default.safeIndex(_145_v95, new BigNumber((_203_v143).length))]);
        let _205_v145;
        _205_v145 = _dafny.Map.Empty.slice().updateUnsafe(_201_v141,!((_dafny.MultiSet.fromElements(_202_v142, _202_v142, _202_v142)).IsSubsetOf(_204_v144)));
        _205_v145 = (_205_v145).update(_201_v141, p0);
        let _206_v146;
        let _nw40 = new _module.C4();
        _nw40.__ctor(((_149_v99).f0).multipliedBy((_149_v99).f0));
        _206_v146 = _nw40;
        let _207_v147;
        let _nw41 = new _module.C2();
        _nw41.__ctor(_145_v95);
        _207_v147 = _nw41;
      }
      let _208_v148;
      let _nw42 = new _module.C2();
      _nw42.__ctor(new BigNumber((_146_v96).length));
      _208_v148 = _nw42;
      let _209_v149;
      _209_v149 = _dafny.Map.Empty.slice().updateUnsafe(!(_145_v95).isEqualTo(p1),_208_v148);
      let _210_v150;
      _210_v150 = _module.D6.create_DC15((_dafny.ZERO).minus(_145_v95), p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(_149_v99).f0)).length));
      let _211_v151;
      _211_v151 = _dafny.Set.fromElements(p0);
      let _212_v152;
      _212_v152 = _module.D5.create_DC13((_149_v99).f0, _211_v151, _18_v0, _145_v95, p0);
      let _213_v153;
      _213_v153 = _dafny.Seq.of(((p0) ? (_210_v150) : (_210_v150)), _module.D6.create_DC15((_212_v152).dtor_cf20, p0, new BigNumber(898)), _210_v150, _210_v150, _210_v150);
      let _214_v154;
      let _init5 = ((_215_v95) => function (_216_i7) {
        return (_216_i7).multipliedBy(_215_v95);
      })(_145_v95);
      let _nw43 = Array((new BigNumber(29)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw43.length); _i0_5++) {
        _nw43[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _214_v154 = _nw43;
      let _217_v155;
      _217_v155 = _dafny.Seq.of(_214_v154, _214_v154);
      let _218_v156;
      _218_v156 = _dafny.Map.Empty.slice().updateUnsafe(p1,_214_v154);
      let _219_v157;
      let _nw44 = Array((new BigNumber(23)).toNumber());
      _nw44[(_dafny.ZERO).toNumber()] = (_217_v155)[_module.__default.safeIndex((_147_v97)[_module.__default.safeIndex((_149_v99).f0, new BigNumber((_147_v97).length))], new BigNumber((_217_v155).length))];
      _nw44[(_dafny.ONE).toNumber()] = ((p0) ? (_214_v154) : (_214_v154));
      _nw44[(new BigNumber(2)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(3)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(4)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(5)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(6)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(7)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(8)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(9)).toNumber()] = (_217_v155)[_module.__default.safeIndex(new BigNumber((_211_v151).length), new BigNumber((_217_v155).length))];
      _nw44[(new BigNumber(10)).toNumber()] = ((p0) ? (_214_v154) : (_214_v154));
      _nw44[(new BigNumber(11)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(12)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(13)).toNumber()] = (((_218_v156).contains((_dafny.ZERO).minus(new BigNumber((_147_v97).length)))) ? ((_218_v156).get((_dafny.ZERO).minus(new BigNumber((_147_v97).length)))) : (_214_v154));
      _nw44[(new BigNumber(14)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(15)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(16)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(17)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(18)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(19)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(20)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(21)).toNumber()] = _214_v154;
      _nw44[(new BigNumber(22)).toNumber()] = _214_v154;
      _219_v157 = _nw44;
      let _index22 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length));
      (_219_v157)[_index22] = _214_v154;
      let _220_v158;
      _220_v158 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_147_v97).length),_209_v149);
      let _221_v159;
      _221_v159 = _dafny.MultiSet.fromElements(p1);
      let _index23 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length));
      let _rhs30 = (_209_v149).Merge((_dafny.Map.Empty.slice().updateUnsafe(p0,_208_v148)).Merge((((_220_v158).contains(new BigNumber((_module.__default.fm19(_221_v159, new BigNumber((_211_v151).length), p1, _145_v95, globalState)).length))) ? ((_220_v158).get(new BigNumber((_module.__default.fm19(_221_v159, new BigNumber((_211_v151).length), p1, _145_v95, globalState)).length))) : (_209_v149))));
      let _rhs31 = _213_v153;
      let _rhs32 = _214_v154;
      let _lhs4 = _219_v157;
      let _lhs5 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length));
      _209_v149 = _rhs30;
      _213_v153 = _rhs31;
      _lhs4[_lhs5] = _rhs32;
      let _222_i8;
      _222_i8 = _dafny.ZERO;
      L0: {
        while ((p1).isLessThanOrEqualTo(((_dafny.ZERO).minus(_145_v95)).minus(_145_v95))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_222_i8)) {
              break L0;
            }
            _222_i8 = (_222_i8).plus(_dafny.ONE);
            let _223_v160;
            let _nw45 = new _module.C4();
            _nw45.__ctor(_145_v95);
            _223_v160 = _nw45;
            _145_v95 = (_module.D13.create_DC27(_18_v0, _223_v160, _18_v0, new BigNumber((_146_v96).length))).dtor_cf46;
            let _224_v161;
            _224_v161 = _dafny.MultiSet.fromElements(p0, true);
            let _225_v162;
            _225_v162 = _dafny.Map.Empty.slice().updateUnsafe((_224_v161).IsProperSubsetOf(_224_v161),_211_v151);
            let _226_v163;
            _226_v163 = _module.D7.create_DC18((_223_v160).f0);
            let _rhs33 = (_223_v160).f0;
            let _rhs34 = _module.__default.fm3(_211_v151, _18_v0, (_149_v99).f0, (_226_v163).dtor_cf30, globalState);
            let _rhs35 = _225_v162;
            let _rhs36 = _dafny.Seq.IsPrefixOf(_148_v98, _dafny.Seq.Concat(_148_v98, _dafny.Seq.of(p0)));
            _145_v95 = _rhs33;
            _145_v95 = _rhs34;
            _225_v162 = _rhs35;
            r0 = _rhs36;
            _145_v95 = new BigNumber((_dafny.Seq.Concat(_146_v96, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uvrdr"), _146_v96))).length);
            let _arr0 = (_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))];
            let _index24 = _module.__default.safeIndex(new BigNumber(937), new BigNumber(((_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))]).length));
            _arr0[_index24] = (_149_v99).f0;
            let _227_v164;
            _227_v164 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_211_v151, _18_v0, p1, (_149_v99).f0, globalState),p0);
            let _228_v165;
            let _nw46 = new _module.C1();
            _nw46.__ctor(new BigNumber((_227_v164).length));
            _228_v165 = _nw46;
            let _229_v166;
            _229_v166 = _dafny.MultiSet.fromElements(_228_v165);
            let _arr1 = (_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))];
            let _index25 = _module.__default.safeIndex(new BigNumber(937), new BigNumber(((_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))]).length));
            _arr1[_index25] = _module.__default.safeDivisionInt(((p0) ? ((((_229_v166).contains(_228_v165)) ? ((_229_v166).get(_228_v165)) : ((_149_v99).f0))) : ((_149_v99).f0)), (_dafny.ZERO).minus((new BigNumber((_148_v98).length)).plus(_145_v95)));
          }
        }
      }
      let _230_v167;
      _230_v167 = _dafny.Seq.of(_221_v159);
      let _hi0 = _145_v95;
      for (let _231_i9 = new BigNumber(((_230_v167)[_module.__default.safeIndex(p1, new BigNumber((_230_v167).length))]).cardinality()); _231_i9.isLessThan(_hi0); _231_i9 = _231_i9.plus(_dafny.ONE)) {
        let _232_v168;
        let _nw47 = Array((new BigNumber(2)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = p0;
        _nw47[(_dafny.ONE).toNumber()] = p0;
        _232_v168 = _nw47;
        let _index26 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length));
        (_232_v168)[_index26] = p0;
        let _index27 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length));
        (_232_v168)[_index27] = false;
        if ((_232_v168)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length))]) {
          let _233_v169;
          _233_v169 = _dafny.Seq.of(_146_v96);
          let _234_v170;
          _234_v170 = _dafny.Set.fromElements(_146_v96, ((false) ? (_146_v96) : (_146_v96)), _dafny.Seq.Concat(_146_v96, _146_v96), (_233_v169)[_module.__default.safeIndex(_231_i9, new BigNumber((_233_v169).length))], _146_v96);
          _234_v170 = _234_v170;
          _145_v95 = (((_221_v159).contains((new BigNumber((_dafny.Seq.update(_148_v98, _module.__default.safeIndex(_231_i9, new BigNumber((_148_v98).length)), (_232_v168)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length))])).length)).plus(_231_i9))) ? ((_221_v159).get((new BigNumber((_dafny.Seq.update(_148_v98, _module.__default.safeIndex(_231_i9, new BigNumber((_148_v98).length)), (_232_v168)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length))])).length)).plus(_231_i9))) : (((false) ? (new BigNumber((_151_v101).length)) : (_module.__default.fm3(_211_v151, _18_v0, new BigNumber((_146_v96).length), (_149_v99).f0, globalState)))));
          let _235_v171;
          _235_v171 = _dafny.Seq.of((((_232_v168)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length))]) ? (_232_v168) : (_232_v168)), _232_v168);
          _235_v171 = _dafny.Seq.of(_232_v168);
          let _236_v172;
          _236_v172 = _dafny.MultiSet.fromElements(_232_v168);
          _236_v172 = ((_236_v172).Union(_dafny.MultiSet.fromElements(_232_v168))).Union(_236_v172);
          let _arr2 = (_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))];
          let _index28 = _module.__default.safeIndex(new BigNumber(124), new BigNumber(((_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))]).length));
          _arr2[_index28] = ((_149_v99).f0).minus((_147_v97)[_module.__default.safeIndex(p1, new BigNumber((_147_v97).length))]);
          let _237_v173;
          let _nw48 = new _module.C1();
          _nw48.__ctor((_149_v99).f0);
          _237_v173 = _nw48;
          let _238_v174;
          _238_v174 = _module.D13.create_DC27(new _dafny.CodePoint('w'.codePointAt(0)), _237_v173, new _dafny.CodePoint('o'.codePointAt(0)), (_149_v99).f0);
          let _arr3 = (_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))];
          let _index29 = _module.__default.safeIndex(new BigNumber(124), new BigNumber(((_219_v157)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_219_v157).length))]).length));
          _arr3[_index29] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(p1, (_149_v99).f0), (_238_v174).dtor_cf46);
        } else {
          _145_v95 = (_149_v99).f0;
          let _239_v175;
          _239_v175 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_18_v0),(_232_v168)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length))]);
          let _240_v176;
          _240_v176 = _dafny.MultiSet.fromElements(_18_v0, _18_v0);
          let _241_v178;
          _241_v178 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(788)), ((_242_v0) => function (_243_i10) {
            return _242_v0;
          })(_18_v0)), _146_v96, _146_v96, _dafny.Seq.UnicodeFromString("hdjof"), _dafny.Seq.UnicodeFromString("edovto"));
          _239_v175 = (_239_v175).update((_240_v176).Difference(_240_v176), (_dafny.Map.Empty.slice().updateUnsafe(_146_v96,_148_v98)).equals(function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_241_v178).Elements) {
              let _244_v177 = _compr_13;
              if (_dafny.Seq.contains(_241_v178, _244_v177)) {
                _coll13.push([_244_v177,_dafny.Seq.of((_232_v168)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length))])]);
              }
            }
            return _coll13;
          }()));
          let _245_v179;
          _245_v179 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(71),_146_v96);
          _145_v95 = ((_149_v99).f0).multipliedBy(new BigNumber((((p0) ? ((((_245_v179).contains((_149_v99).f0)) ? ((_245_v179).get((_149_v99).f0)) : (_dafny.Seq.UnicodeFromString("lv")))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-667)), ((_246_v0) => function (_247_i11) {
            return _246_v0;
          })(_18_v0))))).length));
          let _248_v180;
          let _nw49 = new _module.C5();
          _nw49.__ctor(((false) ? (_18_v0) : (_18_v0)), new BigNumber((_146_v96).length), _145_v95);
          _248_v180 = _nw49;
          let _index30 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_232_v168).length));
          (_232_v168)[_index30] = p0;
        }
        let _249_v181;
        let _nw50 = Array((new BigNumber(21)).toNumber());
        _249_v181 = _nw50;
        let _index31 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_249_v181).length));
        (_249_v181)[_index31] = _208_v148;
        let _index32 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_249_v181).length));
        (_249_v181)[_index32] = ((((_149_v99).f0).isLessThan((_dafny.ZERO).minus(_231_i9))) ? (_208_v148) : (_208_v148));
        let _250_v182;
        _250_v182 = _dafny.Seq.of(_232_v168, _232_v168, _232_v168, _232_v168);
        _232_v168 = (_250_v182)[_module.__default.safeIndex((_dafny.ZERO).minus((_149_v99).f0), new BigNumber((_250_v182).length))];
      }
      r0 = p0;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _251_globalState;
      let _nw51 = new _module.GlobalState();
      _nw51.__ctor();
      _251_globalState = _nw51;
      let _252_v0;
      let _init6 = function (_253_i0) {
        return false;
      };
      let _nw52 = Array((new BigNumber(11)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw52.length); _i0_6++) {
        _nw52[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _252_v0 = _nw52;
      let _index33 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
      (_252_v0)[_index33] = (new BigNumber((_module.__default.fm0(_251_globalState)).length)).isLessThanOrEqualTo(new BigNumber(716));
      let _254_v1;
      _254_v1 = false;
      let _index34 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_252_v0).length));
      (_252_v0)[_index34] = _254_v1;
      let _255_v2;
      _255_v2 = new BigNumber(-806);
      let _256_v3;
      let _init7 = ((_257_v2) => function (_258_i1) {
        return (_258_i1).minus(_257_v2);
      })(_255_v2);
      let _nw53 = Array((new BigNumber(28)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw53.length); _i0_7++) {
        _nw53[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _256_v3 = _nw53;
      let _259_v4;
      _259_v4 = _dafny.MultiSet.fromElements(_256_v3, _256_v3, _256_v3, _256_v3);
      let _260_v5;
      _260_v5 = _dafny.MultiSet.fromElements(_255_v2);
      let _index35 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
      let _index36 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_252_v0).length));
      let _rhs37 = (true) || ((_dafny.MultiSet.fromElements(_255_v2)).IsProperSubsetOf(_260_v5));
      let _rhs38 = ((_254_v1) ? (_252_v0) : (_252_v0));
      let _rhs39 = _254_v1;
      let _rhs40 = _255_v2;
      let _rhs41 = _259_v4;
      let _lhs6 = _252_v0;
      let _lhs7 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
      let _lhs8 = _252_v0;
      let _lhs9 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_252_v0).length));
      _lhs6[_lhs7] = _rhs37;
      _252_v0 = _rhs38;
      _lhs8[_lhs9] = _rhs39;
      _255_v2 = _rhs40;
      _259_v4 = _rhs41;
      if ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) {
        _255_v2 = _255_v2;
        let _261_v6;
        _261_v6 = new _dafny.CodePoint('p'.codePointAt(0));
        _255_v2 = new BigNumber((_module.__default.fm1(_261_v6, _251_globalState)).length);
        if ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) {
          _255_v2 = _255_v2;
          let _262_v7;
          _262_v7 = _dafny.Map.Empty.slice().updateUnsafe(_254_v1,new BigNumber((_dafny.Set.fromElements(!(_254_v1))).length));
          _262_v7 = (_262_v7).update(_254_v1, new BigNumber((_260_v5).cardinality()));
          _261_v6 = _261_v6;
          let _263_v8;
          _263_v8 = _dafny.Seq.UnicodeFromString("qqyjceeg");
          let _264_v9;
          let _out0;
          _out0 = _module.__default.m0(_254_v1, new BigNumber((_263_v8).length), _251_globalState);
          _264_v9 = _out0;
          let _265_v10;
          _265_v10 = _dafny.Seq.of(false);
          let _266_v11;
          _266_v11 = _module.D0.create_DC1((_265_v10)[_module.__default.safeIndex(_255_v2, new BigNumber((_265_v10).length))]);
          let _267_v12;
          _267_v12 = _dafny.Map.Empty.slice().updateUnsafe(_254_v1,_264_v9);
          _264_v9 = (((!((_266_v11).dtor_cf1)) && ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))])) ? (((_254_v1) ? ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) : (_254_v1))) : (!(_module.__default.fm2(_264_v9, _255_v2, !(true), _module.__default.fm2((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], new BigNumber(888), _254_v1, (((_267_v12).contains((_265_v10)[_module.__default.safeIndex(_255_v2, new BigNumber((_265_v10).length))])) ? ((_267_v12).get((_265_v10)[_module.__default.safeIndex(_255_v2, new BigNumber((_265_v10).length))])) : (_264_v9)), _251_globalState), _251_globalState))));
        } else {
          let _268_v13;
          let _out1;
          _out1 = _module.__default.m0((((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) ? ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) : ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))])), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(725)), ((_269_v6) => function (_270_i2) {
            return _269_v6;
          })(_261_v6))).length), _251_globalState);
          _268_v13 = _out1;
          let _271_v14;
          let _nw54 = Array((new BigNumber(3)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _252_v0;
          _nw54[(_dafny.ONE).toNumber()] = _252_v0;
          _nw54[(new BigNumber(2)).toNumber()] = _252_v0;
          _271_v14 = _nw54;
          let _index37 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_271_v14).length));
          (_271_v14)[_index37] = _252_v0;
          let _index38 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_271_v14).length));
          (_271_v14)[_index38] = _252_v0;
          let _index39 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
          (_252_v0)[_index39] = (_268_v13) === (_268_v13);
          let _index40 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
          (_252_v0)[_index40] = _254_v1;
          _268_v13 = !((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
        }
        let _272_v15;
        _272_v15 = _dafny.Set.fromElements(_261_v6);
        let _273_v16;
        _273_v16 = _dafny.Seq.of(_255_v2, _255_v2);
        let _274_v17;
        _274_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), ((_275_v6) => function (_276_i3) {
          return _275_v6;
        })(_261_v6)),new BigNumber((_273_v16).length));
        _254_v1 = (new BigNumber((_272_v15).length)).isLessThanOrEqualTo(new BigNumber((_274_v17).length));
        let _277_v18;
        _277_v18 = _dafny.Map.Empty.slice().updateUnsafe(_255_v2,_255_v2);
        let _278_v19;
        _278_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_277_v18).contains(_255_v2)) ? ((_277_v18).get(_255_v2)) : (_255_v2)),!(_254_v1));
        let _index41 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        (_252_v0)[_index41] = ((_254_v1) ? (!((((_278_v19).contains(_255_v2)) ? ((_278_v19).get(_255_v2)) : (_254_v1))) || (_254_v1)) : (_254_v1));
      } else {
        let _279_v20;
        _279_v20 = new _dafny.CodePoint('o'.codePointAt(0));
        let _280_v21;
        _280_v21 = _dafny.Seq.of((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
        let _281_v22;
        let _out2;
        _out2 = _module.__default.m0((_module.__default.fm3(_dafny.Set.fromElements((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]), _279_v20, _255_v2, _255_v2, _251_globalState)).isLessThan(_255_v2), new BigNumber((_dafny.Set.fromElements(_255_v2, new BigNumber((_280_v21).length))).length), _251_globalState);
        _281_v22 = _out2;
        if (((_254_v1) ? ((new BigNumber((_280_v21).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_280_v21).length)))) : (_281_v22))) {
          let _282_v23;
          _282_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]),_281_v22);
          let _283_v24;
          _283_v24 = _dafny.Set.fromElements(_281_v22);
          let _rhs42 = ((_260_v5).Union(_dafny.MultiSet.fromElements(_module.__default.fm3(_283_v24, _279_v20, _255_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-155)), ((_284_v1) => function (_285_i4) {
            return _module.D0.create_DC0(_284_v1);
          })(_254_v1))).length), _251_globalState), _255_v2))).Union((_260_v5).update(_255_v2, _module.__default.abs(_255_v2)));
          let _rhs43 = _module.__default.fm4(_255_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_286_i5) {
            return new BigNumber(-995);
          })).length), new BigNumber(185), _251_globalState);
          _260_v5 = _rhs42;
          _282_v23 = _rhs43;
          _255_v2 = ((_281_v22) ? (_255_v2) : (_255_v2));
          _279_v20 = _279_v20;
          let _287_v25;
          _287_v25 = _dafny.MultiSet.fromElements(_283_v24);
          _255_v2 = _module.__default.safeModuloInt(new BigNumber(547), new BigNumber((_287_v25).cardinality()));
          let _index42 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
          (_252_v0)[_index42] = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
        } else {
          _255_v2 = (new BigNumber(((_260_v5).Difference(_260_v5)).cardinality())).multipliedBy(_255_v2);
          _281_v22 = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
          _255_v2 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), ((_288_v2) => function (_289_i6) {
            return _288_v2;
          })(_255_v2))).length), _255_v2);
          let _290_v26;
          _290_v26 = _dafny.Seq.of(_255_v2, _255_v2);
          let _291_v27;
          _291_v27 = _dafny.MultiSet.fromElements(_254_v1, (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _254_v1);
          let _292_v28;
          _292_v28 = _dafny.Set.fromElements(new BigNumber((_291_v27).cardinality()));
          let _293_v29;
          _293_v29 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_290_v26, _module.__default.safeIndex(_255_v2, new BigNumber((_290_v26).length)), new BigNumber((_292_v28).length)), _290_v26), _dafny.Seq.Concat(_290_v26, _dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), ((_294_v2) => function (_295_i7) {
            return _294_v2;
          })(_255_v2))));
          let _index43 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
          let _rhs44 = !(_254_v1);
          let _rhs45 = new BigNumber((_293_v29).length);
          let _lhs10 = _252_v0;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
          _lhs10[_lhs11] = _rhs44;
          _255_v2 = _rhs45;
          let _index44 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_256_v3).length));
          (_256_v3)[_index44] = (((_260_v5).contains(_255_v2)) ? ((_260_v5).get(_255_v2)) : ((_dafny.ZERO).minus(_255_v2)));
          let _index45 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_256_v3).length));
          (_256_v3)[_index45] = _255_v2;
        }
        let _296_v30;
        _296_v30 = _dafny.Seq.of(_255_v2);
        let _297_v31;
        let _out3;
        _out3 = _module.__default.m0((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], (_296_v30)[_module.__default.safeIndex(_255_v2, new BigNumber((_296_v30).length))], _251_globalState);
        _297_v31 = _out3;
        let _298_v32;
        let _nw55 = Array((new BigNumber(6)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = (_252_v0);
        _nw55[(_dafny.ONE).toNumber()] = _252_v0;
        _nw55[(new BigNumber(2)).toNumber()] = _252_v0;
        _nw55[(new BigNumber(3)).toNumber()] = _252_v0;
        _nw55[(new BigNumber(4)).toNumber()] = _252_v0;
        _nw55[(new BigNumber(5)).toNumber()] = _252_v0;
        _298_v32 = _nw55;
        let _index46 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_298_v32).length));
        (_298_v32)[_index46] = _252_v0;
        let _299_v33;
        let _nw56 = Array((new BigNumber(9)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = _256_v3;
        _nw56[(_dafny.ONE).toNumber()] = _256_v3;
        _nw56[(new BigNumber(2)).toNumber()] = _256_v3;
        _nw56[(new BigNumber(3)).toNumber()] = _256_v3;
        _nw56[(new BigNumber(4)).toNumber()] = _256_v3;
        _nw56[(new BigNumber(5)).toNumber()] = _256_v3;
        _nw56[(new BigNumber(6)).toNumber()] = _256_v3;
        _nw56[(new BigNumber(7)).toNumber()] = _256_v3;
        _nw56[(new BigNumber(8)).toNumber()] = _256_v3;
        _299_v33 = _nw56;
        let _300_v34;
        _300_v34 = _dafny.Seq.UnicodeFromString("w");
        let _index47 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_256_v3).length));
        (_256_v3)[_index47] = new BigNumber((_300_v34).length);
        let _301_v35;
        _301_v35 = _module.D0.create_DC1(!(_254_v1));
        let _302_v36;
        _302_v36 = _dafny.Seq.of(_301_v35);
        let _303_v37;
        _303_v37 = _dafny.Map.Empty.slice().updateUnsafe((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))],(_module.D2.create_DC3(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_304_v35) => function (_305_i8) {
  return _304_v35;
})(_301_v35)), _module.__default.safeIndex(_255_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_306_v35) => function (_307_i8) {
  return _306_v35;
})(_301_v35))).length)), _301_v35))).dtor_cf3);
        let _308_v38;
        _308_v38 = _dafny.Seq.of(_302_v36, (((_303_v37).contains(_281_v22)) ? ((_303_v37).get(_281_v22)) : (_302_v36)), _302_v36, _dafny.Seq.of(_301_v35, _301_v35));
        let _index48 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_298_v32).length));
        let _index49 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_256_v3).length));
        let _rhs46 = _252_v0;
        let _rhs47 = _299_v33;
        let _rhs48 = _dafny.Seq.IsPrefixOf(_302_v36, (_308_v38)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_308_v38).length))]);
        let _rhs49 = (_255_v2).minus(new BigNumber((_dafny.Set.fromElements(_281_v22)).length));
        let _lhs12 = _298_v32;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_298_v32).length));
        let _lhs14 = _256_v3;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_256_v3).length));
        _lhs12[_lhs13] = _rhs46;
        _299_v33 = _rhs47;
        _281_v22 = _rhs48;
        _lhs14[_lhs15] = _rhs49;
        let _index50 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        (_252_v0)[_index50] = !((_280_v21)[_module.__default.safeIndex((_256_v3)[_module.__default.safeIndex(new BigNumber(807), new BigNumber((_256_v3).length))], new BigNumber((_280_v21).length))]);
      }
      if (false) {
        let _index51 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        (_252_v0)[_index51] = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
        let _309_v39;
        let _out4;
        _out4 = _module.__default.m0(!((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]), _255_v2, _251_globalState);
        _309_v39 = _out4;
        _252_v0 = _252_v0;
        let _index52 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        let _rhs50 = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
        let _rhs51 = _309_v39;
        let _rhs52 = _255_v2;
        let _lhs16 = _252_v0;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        _309_v39 = _rhs50;
        _lhs16[_lhs17] = _rhs51;
        _255_v2 = _rhs52;
        let _310_v40;
        _310_v40 = _dafny.Seq.UnicodeFromString("qn");
        let _311_v41;
        _311_v41 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fylsrkxqe")).length), _255_v2, _255_v2);
        let _index53 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        (_252_v0)[_index53] = (((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) ? ((new BigNumber((_310_v40).length)).isEqualTo(_255_v2)) : ((_311_v41).IsProperSubsetOf(_311_v41)));
      } else {
        let _312_v42;
        _312_v42 = _dafny.MultiSet.fromElements(_254_v1, !(_254_v1), (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
        let _313_v43;
        _313_v43 = _dafny.Seq.of(_254_v1);
        _312_v42 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(((_254_v1) ? (_313_v43) : (_dafny.Seq.of(_254_v1, !(_254_v1)))), _313_v43));
        _255_v2 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_255_v2));
        let _314_v44;
        _314_v44 = _dafny.Map.Empty.slice().updateUnsafe((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))],(_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
        let _315_v45;
        let _out5;
        _out5 = _module.__default.m0(!((_dafny.Map.Empty.slice().updateUnsafe(false,_254_v1)).equals((_314_v44).update(_254_v1, _module.__default.fm2((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _255_v2, (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _254_v1, _251_globalState)))), _255_v2, _251_globalState);
        _315_v45 = _out5;
        if (!(_254_v1)) {
          let _index54 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_256_v3).length));
          (_256_v3)[_index54] = new BigNumber((_dafny.Set.fromElements(_315_v45, (_313_v43)[_module.__default.safeIndex(new BigNumber(930), new BigNumber((_313_v43).length))])).length);
          let _index55 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_256_v3).length));
          (_256_v3)[_index55] = _module.__default.safeModuloInt(_255_v2, new BigNumber(501));
          _255_v2 = _255_v2;
          _254_v1 = _254_v1;
          _314_v44 = (_314_v44).update(_315_v45, _315_v45);
          _315_v45 = false;
        } else {
          let _316_v46;
          let _out6;
          _out6 = _module.__default.m0(!(_315_v45), _255_v2, _251_globalState);
          _316_v46 = _out6;
          let _317_v47;
          let _out7;
          _out7 = _module.__default.m0(_254_v1, _255_v2, _251_globalState);
          _317_v47 = _out7;
          _255_v2 = (_dafny.ZERO).minus((_255_v2).plus(_255_v2));
          let _318_v48;
          let _nw57 = new _module.C1();
          _nw57.__ctor(_255_v2);
          _318_v48 = _nw57;
          let _319_v49;
          _319_v49 = new _dafny.CodePoint('g'.codePointAt(0));
          let _320_v50;
          _320_v50 = _dafny.Seq.of(_319_v49);
          let _321_v51;
          _321_v51 = _dafny.Map.Empty.slice().updateUnsafe(_320_v50,_315_v45);
          _321_v51 = (_321_v51).update(_dafny.Seq.of(_319_v49), _dafny.Seq.IsProperPrefixOf(_313_v43, _313_v43));
        }
        let _322_v52;
        _322_v52 = new _dafny.CodePoint('y'.codePointAt(0));
        let _323_v53;
        let _nw58 = new _module.C5();
        _nw58.__ctor(_322_v52, _255_v2, _255_v2);
        _323_v53 = _nw58;
      }
      let _324_v54;
      let _nw59 = new _module.C3();
      _nw59.__ctor(_255_v2);
      _324_v54 = _nw59;
      let _325_v55;
      _325_v55 = _module.D16.create_DC33(_324_v54);
      let _326_v56;
      let _nw60 = Array((new BigNumber(5)).toNumber());
      _nw60[(_dafny.ZERO).toNumber()] = _324_v54;
      _nw60[(_dafny.ONE).toNumber()] = _324_v54;
      _nw60[(new BigNumber(2)).toNumber()] = _324_v54;
      _nw60[(new BigNumber(3)).toNumber()] = _324_v54;
      _nw60[(new BigNumber(4)).toNumber()] = (_325_v55).dtor_cf54;
      _326_v56 = _nw60;
      let _327_v57;
      _327_v57 = _dafny.Seq.of(_326_v56);
      let _328_v58;
      _328_v58 = new _dafny.CodePoint('x'.codePointAt(0));
      let _329_v59;
      _329_v59 = _dafny.Map.Empty.slice().updateUnsafe((_327_v57)[_module.__default.safeIndex(_255_v2, new BigNumber((_327_v57).length))],_328_v58);
      _329_v59 = (_329_v59).Merge(_329_v59);
      _255_v2 = _255_v2;
      _328_v58 = _328_v58;
      let _330_v60;
      _330_v60 = _module.D2.create_DC4(_dafny.Seq.UnicodeFromString("ontdavgjc"), (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
      let _pat_let_tv2 = _330_v60;
      let _pat_let_tv3 = _330_v60;
      let _pat_let_tv4 = _330_v60;
      let _index56 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
      let _rhs53 = ((_255_v2).plus(_255_v2)).multipliedBy(new BigNumber(-502));
      let _rhs54 = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
      let _rhs55 = function (_source1) {
        if (_source1.is_DC12) {
          let _331___mcc_h0 = (_source1).cf15;
          let _332___mcc_h1 = (_source1).cf16;
          let _333_cf16 = _332___mcc_h1;
          let _334_cf15 = _331___mcc_h0;
          return _pat_let_tv2;
        } else if (_source1.is_DC13) {
          let _335___mcc_h2 = (_source1).cf17;
          let _336___mcc_h3 = (_source1).cf18;
          let _337___mcc_h4 = (_source1).cf19;
          let _338___mcc_h5 = (_source1).cf20;
          let _339___mcc_h6 = (_source1).cf21;
          let _340_cf21 = _339___mcc_h6;
          let _341_cf20 = _338___mcc_h5;
          let _342_cf19 = _337___mcc_h4;
          let _343_cf18 = _336___mcc_h3;
          let _344_cf17 = _335___mcc_h2;
          return _pat_let_tv3;
        } else {
          let _345___mcc_h7 = (_source1).cf14;
          let _346_cf14 = _345___mcc_h7;
          return _pat_let_tv4;
        }
      }(_module.D5.create_DC12(_255_v2, _255_v2));
      let _lhs18 = _252_v0;
      let _lhs19 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
      _255_v2 = _rhs53;
      _lhs18[_lhs19] = _rhs54;
      _330_v60 = _rhs55;
      let _347_v61;
      _347_v61 = _dafny.Set.fromElements((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
      let _348_i9;
      _348_i9 = _dafny.ZERO;
      L1: {
        while ((_347_v61).IsProperSubsetOf(_dafny.Set.fromElements(_254_v1, _254_v1, (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_348_i9)) {
              break L1;
            }
            _348_i9 = (_348_i9).plus(_dafny.ONE);
            let _349_v62;
            _349_v62 = _module.D3.create_DC8((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _255_v2);
            let _350_v63;
            let _out8;
            _out8 = _module.__default.m0(((_dafny.ZERO).minus(_module.__default.fm3(_dafny.Set.fromElements(_254_v1, _254_v1), _328_v58, _255_v2, _255_v2, _251_globalState))).isLessThan((_349_v62).dtor_cf9), _255_v2, _251_globalState);
            _350_v63 = _out8;
            let _index57 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_256_v3).length));
            (_256_v3)[_index57] = _module.__default.safeDivisionInt(_255_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), ((_351_v0) => function (_352_i10) {
              return (_module.D13.create_DC29(new _dafny.CodePoint('j'.codePointAt(0)), (_351_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_351_v0).length))])).dtor_cf49;
            })(_252_v0))).length));
            let _353_v64;
            let _nw61 = new _module.C5();
            _nw61.__ctor(_328_v58, _255_v2, _255_v2);
            _353_v64 = _nw61;
            let _354_v65;
            _354_v65 = _dafny.Map.Empty.slice().updateUnsafe(_353_v64,_255_v2);
            let _index58 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_256_v3).length));
            (_256_v3)[_index58] = (((_354_v65).contains(_353_v64)) ? ((_354_v65).get(_353_v64)) : (_255_v2));
            let _355_v66;
            let _nw62 = Array((new BigNumber(9)).toNumber()).fill([]);
            _355_v66 = _nw62;
            _355_v66 = _355_v66;
            let _index59 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_256_v3).length));
            (_256_v3)[_index59] = (_353_v64).f2;
          }
        }
      }
      _255_v2 = _module.__default.safeDivisionInt(_255_v2, _255_v2);
      let _356_i11;
      _356_i11 = _dafny.ZERO;
      L2: {
        while (!(_254_v1)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_356_i11)) {
              break L2;
            }
            _356_i11 = (_356_i11).plus(_dafny.ONE);
            _254_v1 = _254_v1;
            _255_v2 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_255_v2, new BigNumber(770)), _255_v2);
            if ((_255_v2).isEqualTo(_255_v2)) {
              let _357_v68;
              let _nw63 = new _module.C3();
              _nw63.__ctor(new BigNumber((function () {
                let _coll14 = new _dafny.Map();
                for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_255_v2,_254_v1)).Keys.Elements) {
                  let _358_v67 = _compr_14;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_255_v2,_254_v1)).contains(_358_v67)) {
                    _coll14.push([_module.__default.safeDivisionInt(_358_v67, new BigNumber((_347_v61).length)),_255_v2]);
                  }
                }
                return _coll14;
              }()).length));
              _357_v68 = _nw63;
              let _359_v69;
              _359_v69 = _module.D13.create_DC27(_328_v58, _357_v68, _328_v58, _255_v2);
              _255_v2 = (_359_v69).dtor_cf46;
              let _360_v70;
              _360_v70 = _dafny.Set.fromElements(_252_v0);
              _255_v2 = _module.__default.safeDivisionInt(new BigNumber((_360_v70).length), (_dafny.ZERO).minus((_359_v69).dtor_cf46));
              let _361_v71;
              _361_v71 = _dafny.Seq.UnicodeFromString("nmqirv");
              (_324_v54).m7(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("oaqldlvsi"), _361_v71), _254_v1, _256_v3, new BigNumber((_361_v71).length), _251_globalState);
              _255_v2 = (_357_v68).f0;
              let _362_v72;
              _362_v72 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), ((_363_v58) => function (_364_i12) {
                return _363_v58;
              })(_328_v58)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), ((_365_v58) => function (_366_i13) {
                return _365_v58;
              })(_328_v58)));
              let _367_v73;
              _367_v73 = _dafny.Seq.of((_255_v2).isLessThanOrEqualTo(new BigNumber(-839)), _dafny.Seq.IsPrefixOf(_361_v71, (_362_v72)[_module.__default.safeIndex((_357_v68).f0, new BigNumber((_362_v72).length))]), _254_v1);
              let _368_v74;
              _368_v74 = _dafny.Map.Empty.slice().updateUnsafe(_328_v58,new BigNumber((_dafny.Seq.UnicodeFromString("qmslbj")).length));
              let _369_v75;
              _369_v75 = _dafny.Set.fromElements((((_368_v74).contains(_328_v58)) ? ((_368_v74).get(_328_v58)) : (_255_v2)), (_357_v68).f0);
              let _370_v76;
              _370_v76 = _dafny.MultiSet.fromElements(_328_v58, _328_v58);
              let _371_v77;
              _371_v77 = _dafny.Map.Empty.slice().updateUnsafe((_357_v68).f0,_dafny.MultiSet.fromElements(_328_v58));
              let _372_v78;
              _372_v78 = _dafny.Seq.of(_370_v76, (((_371_v77).contains(_255_v2)) ? ((_371_v77).get(_255_v2)) : (_370_v76)));
              let _rhs56 = _module.__default.fm3(_347_v61, (((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) ? (_328_v58) : (_328_v58)), (_255_v2).multipliedBy((_357_v68).f0), _module.__default.safeDivisionInt(new BigNumber((_361_v71).length), new BigNumber((_369_v75).length)), _251_globalState);
              let _rhs57 = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
              let _rhs58 = _367_v73;
              let _rhs59 = _module.__default.fm10(((_372_v78)[_module.__default.safeIndex(new BigNumber(447), new BigNumber((_372_v78).length))]).Intersect(_370_v76), (_324_v54).fm13(_255_v2, (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _251_globalState), _251_globalState);
              let _rhs60 = _328_v58;
              _255_v2 = _rhs56;
              _254_v1 = _rhs57;
              _367_v73 = _rhs58;
              _347_v61 = _rhs59;
              _328_v58 = _rhs60;
            } else {
              _255_v2 = _255_v2;
              let _index60 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
              (_252_v0)[_index60] = (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))];
              let _373_v79;
              _373_v79 = _dafny.Map.Empty.slice().updateUnsafe(_255_v2,false);
              let _374_v80;
              _374_v80 = _dafny.MultiSet.fromElements(_328_v58, _328_v58);
              let _375_v81;
              _375_v81 = _dafny.Map.Empty.slice().updateUnsafe(_374_v80,_256_v3);
              _373_v79 = (_373_v79).update(_255_v2, ((_375_v81).update(_374_v80, _256_v3)).contains(_374_v80));
              let _376_v82;
              _376_v82 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_255_v2, _255_v2),(_dafny.Set.fromElements((_324_v54).fm13(new BigNumber(-924), (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _251_globalState))).Union(_347_v61));
              _376_v82 = (_376_v82).update(_255_v2, _347_v61);
              let _377_v83;
              _377_v83 = _module.D6.create_DC14(_256_v3);
              let _378_v84;
              _378_v84 = _dafny.Seq.of(_255_v2);
              let _379_v85;
              _379_v85 = _dafny.Seq.UnicodeFromString("ryvnthp");
              let _380_v86;
              let _nw64 = new _module.C4();
              _nw64.__ctor(_255_v2);
              _380_v86 = _nw64;
              let _381_v87;
              _381_v87 = _dafny.Seq.of(_380_v86);
              let _382_v88;
              _382_v88 = _dafny.Seq.of(_380_v86, (_381_v87)[_module.__default.safeIndex(_255_v2, new BigNumber((_381_v87).length))]);
              (_324_v54).m7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), function (_383_i14) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              }), _254_v1, (((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) ? (_256_v3) : ((_377_v83).dtor_cf22)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_378_v84, _module.__default.safeIndex(new BigNumber(180), new BigNumber((_378_v84).length)), _255_v2), _dafny.Seq.of(new BigNumber((_379_v85).length), new BigNumber((_382_v88).length), _255_v2))).length)), _251_globalState);
            }
            let _index61 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
            (_252_v0)[_index61] = _254_v1;
          }
        }
      }
      let _384_i15;
      _384_i15 = _dafny.ZERO;
      L3: {
        while ((_254_v1) || ((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))])) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_384_i15)) {
              break L3;
            }
            _384_i15 = (_384_i15).plus(_dafny.ONE);
            let _385_v89;
            let _nw65 = new _module.C5();
            _nw65.__ctor(_328_v58, _255_v2, _255_v2);
            _385_v89 = _nw65;
            let _386_v90;
            let _nw66 = Array((new BigNumber(13)).toNumber());
            _386_v90 = _nw66;
            let _387_v91;
            _387_v91 = _dafny.Map.Empty.slice().updateUnsafe(_385_v89,_386_v90);
            let _388_v92;
            _388_v92 = _module.D3.create_DC9(_255_v2, (((_387_v91).contains(_385_v89)) ? ((_387_v91).get(_385_v89)) : (_386_v90)), (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
            let _389_v93;
            _389_v93 = _dafny.Set.fromElements((_388_v92).dtor_cf10, (_385_v89).f0, new BigNumber(20));
            let _390_v94;
            _390_v94 = _module.D5.create_DC13((_255_v2).multipliedBy(_module.__default.fm3(_dafny.Set.fromElements(_254_v1), _328_v58, _255_v2, _255_v2, _251_globalState)), _347_v61, _328_v58, _module.__default.safeModuloInt(_255_v2, _255_v2), (_389_v93).IsDisjointFrom(_389_v93));
            _390_v94 = _390_v94;
            let _391_v95;
            _391_v95 = _dafny.Map.Empty.slice().updateUnsafe((_385_v89).f0,(_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]);
            _391_v95 = (_391_v95).update((_385_v89).f0, _254_v1);
            _256_v3 = _256_v3;
            let _index62 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_256_v3).length));
            (_256_v3)[_index62] = (_dafny.ZERO).minus((_385_v89).f0);
            let _index63 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_256_v3).length));
            (_256_v3)[_index63] = _255_v2;
          }
        }
      }
      let _392_v96;
      let _nw67 = new _module.C3();
      _nw67.__ctor(_module.__default.safeDivisionInt(new BigNumber(747), new BigNumber(988)));
      _392_v96 = _nw67;
      let _393_v97;
      _393_v97 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(811)), function (_394_i16) {
        return _dafny.Seq.UnicodeFromString("nsgqps");
      });
      let _395_v98;
      _395_v98 = _dafny.Map.Empty.slice().updateUnsafe((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))],_393_v97);
      _395_v98 = (_395_v98).update((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], _393_v97);
      let _396_v99;
      _396_v99 = _dafny.Seq.of(_328_v58);
      (_324_v54).m7(_dafny.Seq.UnicodeFromString("nujtypxb"), (_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))], ((true) ? (_256_v3) : (_256_v3)), new BigNumber((_396_v99).length), _251_globalState);
      let _397_v100;
      let _nw68 = Array((new BigNumber(22)).toNumber()).fill([]);
      _397_v100 = _nw68;
      let _398_v101;
      let _399_v102;
      let _400_v103;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector0 = (_324_v54).m1(_397_v100, ((_254_v1) ? ((_dafny.ZERO).minus(_255_v2)) : ((_dafny.ZERO).minus(_255_v2))), _251_globalState);
      _out9 = _outcollector0[0];
      _out10 = _outcollector0[1];
      _out11 = _outcollector0[2];
      _398_v101 = _out9;
      _399_v102 = _out10;
      _400_v103 = _out11;
      let _401_v104;
      _401_v104 = _dafny.Seq.of(_255_v2);
      let _hi1 = (((_252_v0)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length))]) ? (_255_v2) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-266)), ((_402_v58) => function (_403_i18) {
        return _402_v58;
      })(_328_v58))).length)));
      for (let _404_i17 = (_401_v104)[_module.__default.safeIndex(_398_v101, new BigNumber((_401_v104).length))]; _404_i17.isLessThan(_hi1); _404_i17 = _404_i17.plus(_dafny.ONE)) {
        let _405_v105;
        _405_v105 = _dafny.Map.Empty.slice().updateUnsafe((_260_v5).IsProperSubsetOf(_dafny.MultiSet.fromElements(_255_v2)),_252_v0);
        _405_v105 = (_405_v105).update(!_dafny.areEqual((_330_v60).dtor_cf4, _396_v99), ((_254_v1) ? (_252_v0) : (_252_v0)));
        let _index64 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_252_v0).length));
        (_252_v0)[_index64] = _254_v1;
        let _406_v106;
        _406_v106 = _dafny.Map.Empty.slice().updateUnsafe(_255_v2,_255_v2);
        let _rhs61 = (((_406_v106).contains(_module.__default.fm3(_347_v61, _328_v58, _404_i17, _404_i17, _251_globalState))) ? ((_406_v106).get(_module.__default.fm3(_347_v61, _328_v58, _404_i17, _404_i17, _251_globalState))) : (_399_v102));
        let _rhs62 = _dafny.Seq.Concat(_396_v99, _396_v99);
        _255_v2 = _rhs61;
        _396_v99 = _rhs62;
        let _407_v107;
        let _init8 = ((_408_v104) => function (_409_i19) {
          return (_dafny.Set.fromElements(new BigNumber((_408_v104).length)));
        })(_401_v104);
        let _nw69 = Array((new BigNumber(20)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw69.length); _i0_8++) {
          _nw69[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _407_v107 = _nw69;
        _407_v107 = _407_v107;
      }
      process.stdout.write(_dafny.toString((_252_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_254_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_255_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v3)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_259_v4).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v5).equals(_dafny.MultiSet.fromElements(new BigNumber(-806)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_325_v55).dtor_cf54).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_326_v56)[_dafny.ZERO]).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_326_v56)[_dafny.ONE]).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_326_v56)[new BigNumber(2)]).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_326_v56)[new BigNumber(3)]).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_326_v56)[new BigNumber(4)]).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_327_v57).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_328_v58));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_329_v59).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_330_v60).dtor_cf4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_330_v60).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_347_v61).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_348_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_356_i11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_384_i15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_393_v97), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_395_v98).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(_dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"), _dafny.Seq.UnicodeFromString("nsgqps"))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_396_v99, _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_398_v101));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_399_v102));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_400_v103));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_401_v104, _dafny.Seq.of(_dafny.ONE))));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false);
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
    static create_DC2(cf2) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf2 === other.cf2;
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf4, cf5) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC5() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D2(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC6(cf6) {
      let $dt = new D2(3);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + this.cf4.toVerbatimString(true) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC8(cf8, cf9) {
      let $dt = new D3(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC9(cf10, cf11, cf12) {
      let $dt = new D3(1);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7(cf7) {
      let $dt = new D3(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11 && this.cf12 === other.cf12;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(false, _dafny.ZERO);
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
    static create_DC10(cf13) {
      let $dt = new D4(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC12(cf15, cf16) {
      let $dt = new D5(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC13(cf17, cf18, cf19, cf20, cf21) {
      let $dt = new D5(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf14) {
      let $dt = new D5(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf14) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(_dafny.ZERO, _dafny.ZERO);
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
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf22 === other.cf22;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC17(cf27, cf28, cf29) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC18(cf30) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D7(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27 && this.cf28 === other.cf28 && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(false, false, false);
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
    static create_DC19(cf31) {
      let $dt = new D8(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf31 === other.cf31;
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf32) {
      let $dt = new D9(0);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32);
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
    static create_DC22(cf34, cf35, cf36) {
      let $dt = new D10(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D10(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf33 === other.cf33;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22(_dafny.ZERO, false, false);
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
    static create_DC24(cf38, cf39, cf40) {
      let $dt = new D11(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC23(cf37) {
      let $dt = new D11(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38 && this.cf39 === other.cf39 && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC24([], false, false);
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
    static create_DC25(cf41) {
      let $dt = new D12(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC25" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41;
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf43, cf44, cf45, cf46) {
      let $dt = new D13(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC28(cf47, cf48) {
      let $dt = new D13(1);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC29(cf49, cf50) {
      let $dt = new D13(2);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC26(cf42) {
      let $dt = new D13(3);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC30(cf51) {
      let $dt = new D13(4);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC26() { return this.$tag === 3; }
    get is_DC30() { return this.$tag === 4; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC27" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC28" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC26" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 4) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44 && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf47 === other.cf47 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf42 === other.cf42;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC27(new _dafny.CodePoint('D'.codePointAt(0)), null, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    get is_DC31() { return this.$tag === 0; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52);
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf53) {
      let $dt = new D15(0);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf53) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC34() {
      let $dt = new D16(0);
      return $dt;
    }
    static create_DC35(cf55) {
      let $dt = new D16(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC33(cf54) {
      let $dt = new D16(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC36(cf56) {
      let $dt = new D16(3);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC36() { return this.$tag === 3; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC34";
      } else if (this.$tag === 1) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC33" + "(" + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC36" + "(" + _dafny.toString(this.cf56) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf54 === other.cf54;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf56, other.cf56);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC34();
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
    static create_DC37(cf57) {
      let $dt = new D17(0);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC37" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf57, other.cf57);
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC39(cf59) {
      let $dt = new D18(0);
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC40(cf60, cf61) {
      let $dt = new D18(1);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC38(cf58) {
      let $dt = new D18(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf58() { return this.cf58; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC39" + "(" + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC40" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC38" + "(" + _dafny.toString(this.cf58) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf60 === other.cf60 && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf58, other.cf58);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC39(false);
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
    static create_DC42(cf63, cf64) {
      let $dt = new D19(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC41(cf62) {
      let $dt = new D19(1);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC43(cf65) {
      let $dt = new D19(2);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC42" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC41" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC43" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC42(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
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
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f3 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f3) {
      let _this = this;
      (_this)._f3 = f3;
      return;
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f0 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return !_dafny.areEqual(_dafny.Seq.UnicodeFromString("m"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(600)), function (_410_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("s")));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(false) || (true);
    };
    fm15(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f0;
    };
    fm16(p0, p1, globalState) {
      let _this = this;
      return !((_dafny.MultiSet.fromElements(false)).IsSubsetOf(_dafny.MultiSet.fromElements(false, !(false))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _411_v0;
      _411_v0 = true;
      let _412_i0;
      _412_i0 = _dafny.ZERO;
      L4: {
        while (((true) ? (_411_v0) : (((_411_v0) ? (!(_411_v0)) : (_411_v0))))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_412_i0)) {
              break L4;
            }
            _412_i0 = (_412_i0).plus(_dafny.ONE);
            let _413_v1;
            let _nw70 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
            _413_v1 = _nw70;
            let _index65 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_413_v1).length));
            (_413_v1)[_index65] = (_this).f0;
            let _index66 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_413_v1).length));
            (_413_v1)[_index66] = p1;
            let _414_v2;
            let _out12;
            _out12 = _module.__default.m0((p1).isLessThan((_dafny.ZERO).minus(new BigNumber(-237))), (_this).f0, globalState);
            _414_v2 = _out12;
            let _415_v3;
            _415_v3 = new _dafny.CodePoint('h'.codePointAt(0));
            let _416_v4;
            _416_v4 = _dafny.Map.Empty.slice().updateUnsafe(_415_v3,new BigNumber(-500));
            _416_v4 = (_416_v4).Merge(_416_v4);
            let _417_v5;
            let _nw71 = new _module.C0();
            _nw71.__ctor(_415_v3);
            _417_v5 = _nw71;
          }
        }
      }
      let _418_v6;
      _418_v6 = new _dafny.CodePoint('e'.codePointAt(0));
      let _419_v7;
      let _nw72 = new _module.C0();
      _nw72.__ctor(_418_v6);
      _419_v7 = _nw72;
      let _420_v8;
      _420_v8 = _dafny.Seq.UnicodeFromString("fce");
      let _421_v9;
      _421_v9 = _dafny.Map.Empty.slice().updateUnsafe(_420_v8,_411_v0);
      let _422_v10;
      _422_v10 = _dafny.Seq.of(new BigNumber(584));
      r0 = (_this).fm15(_421_v9, !(_411_v0), _dafny.areEqual(_422_v10, _422_v10), globalState);
      let _423_v11;
      let _nw73 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _423_v11 = _nw73;
      let _index67 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_423_v11).length));
      (_423_v11)[_index67] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f0, (_dafny.ZERO).minus(p1)));
      let _index68 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_423_v11).length));
      (_423_v11)[_index68] = ((_411_v0) ? ((_this).f0) : (new BigNumber((_dafny.Seq.UnicodeFromString("flbjqa")).length)));
      let _424_v12;
      _424_v12 = _dafny.MultiSet.fromElements(_411_v0);
      let _425_v13;
      _425_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC7(_422_v10),(_424_v12).equals(_424_v12));
      let _426_v14;
      _426_v14 = _module.D3.create_DC7(_422_v10);
      _425_v13 = (_425_v13).update(_426_v14, (_424_v12).IsSubsetOf(_424_v12));
      let _427_v15;
      _427_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_dafny.Seq.Create(_module.__default.abs(new BigNumber(718)), ((_428_v7) => function (_429_i2) {
        return (_428_v7).f3;
      })(_419_v7)));
      let _430_v16;
      let _nw74 = Array((new BigNumber(25)).toNumber());
      _nw74[(_dafny.ZERO).toNumber()] = ((_411_v0) ? (_420_v8) : (_dafny.Seq.update(_420_v8, _module.__default.safeIndex(new BigNumber((_420_v8).length), new BigNumber((_420_v8).length)), new _dafny.CodePoint('y'.codePointAt(0)))));
      _nw74[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("btdr"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), function (_431_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }));
      _nw74[(new BigNumber(2)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(3)).toNumber()] = (((_427_v15).contains((_this).f0)) ? ((_427_v15).get((_this).f0)) : (_420_v8));
      _nw74[(new BigNumber(4)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_420_v8, _module.__default.safeIndex((_423_v11)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_423_v11).length))], new BigNumber((_420_v8).length)), _module.__default.fm17(p1, globalState));
      _nw74[(new BigNumber(6)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(7)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(8)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("cx");
      _nw74[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_420_v8, _420_v8);
      _nw74[(new BigNumber(11)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(12)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_420_v8, _dafny.Seq.UnicodeFromString("u"));
      _nw74[(new BigNumber(14)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(15)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(16)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(730)), ((_432_v7) => function (_433_i3) {
        return (_432_v7).f3;
      })(_419_v7));
      _nw74[(new BigNumber(18)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(19)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_420_v8, _420_v8);
      _nw74[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("oxyrf");
      _nw74[(new BigNumber(22)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(23)).toNumber()] = _420_v8;
      _nw74[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-916)), ((_434_v6) => function (_435_i4) {
        return _434_v6;
      })(_418_v6));
      _430_v16 = _nw74;
      let _436_v17;
      _436_v17 = _module.D0.create_DC0(!(_dafny.Seq.IsProperPrefixOf(_420_v8, _420_v8)));
      let _437_v18;
      _437_v18 = _dafny.Map.Empty.slice().updateUnsafe(_411_v0,_411_v0);
      let _index69 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_423_v11).length));
      let _rhs63 = ((_411_v0) ? (_430_v16) : (_430_v16));
      let _rhs64 = p1;
      let _rhs65 = ((_411_v0) ? (_420_v8) : (_dafny.Seq.UnicodeFromString("jhobkl")));
      let _rhs66 = ((((_dafny.ZERO).minus((_423_v11)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_423_v11).length))])).isLessThanOrEqualTo((_this).f0)) ? (_module.D0.create_DC0((((_437_v18).contains(_411_v0)) ? ((_437_v18).get(_411_v0)) : (true)))) : (_436_v17));
      let _lhs20 = _423_v11;
      let _lhs21 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_423_v11).length));
      _430_v16 = _rhs63;
      _lhs20[_lhs21] = _rhs64;
      _420_v8 = _rhs65;
      _436_v17 = _rhs66;
      r0 = (_this).f0;
      r1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_420_v8, _420_v8), _420_v8)).length);
      r2 = p1;
      return [r0, r1, r2];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f0 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return (((true) ? (true) : (true))) && (!(!(true) || (false)));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((false) || (true));
    };
    m2(globalState) {
      let _this = this;
      let _438_v0;
      _438_v0 = true;
      let _439_v1;
      let _nw75 = Array((new BigNumber(2)).toNumber());
      _nw75[(_dafny.ZERO).toNumber()] = _438_v0;
      _nw75[(_dafny.ONE).toNumber()] = _438_v0;
      _439_v1 = _nw75;
      let _440_v2;
      _440_v2 = _dafny.MultiSet.fromElements(_439_v1);
      let _441_v3;
      _441_v3 = _dafny.Seq.of(new BigNumber((_440_v2).cardinality()));
      let _pat_let_tv5 = _441_v3;
      let _pat_let_tv6 = _441_v3;
      _441_v3 = function (_source2) {
        if (_source2.is_DC1) {
          let _442___mcc_h0 = (_source2).cf1;
          let _443_cf1 = _442___mcc_h0;
          return _pat_let_tv5;
        } else {
          let _444___mcc_h1 = (_source2).cf0;
          let _445_cf0 = _444___mcc_h1;
          return _pat_let_tv6;
        }
      }(_module.D0.create_DC0(_438_v0));
      let _446_v4;
      _446_v4 = _dafny.Seq.UnicodeFromString("bdfea");
      let _447_v5;
      _447_v5 = _module.D2.create_DC4(_446_v4, _438_v0);
      let _pat_let_tv7 = _438_v0;
      let _pat_let_tv8 = _438_v0;
      let _pat_let_tv9 = _438_v0;
      let _pat_let_tv10 = _438_v0;
      let _pat_let_tv11 = _438_v0;
      _438_v0 = function (_source3) {
        if (_source3.is_DC4) {
          let _448___mcc_h2 = (_source3).cf4;
          let _449___mcc_h3 = (_source3).cf5;
          let _450_cf5 = _449___mcc_h3;
          let _451_cf4 = _448___mcc_h2;
          return _pat_let_tv7;
        } else if (_source3.is_DC5) {
          return _pat_let_tv8;
        } else if (_source3.is_DC3) {
          let _452___mcc_h4 = (_source3).cf3;
          let _453_cf3 = _452___mcc_h4;
          return !((_pat_let_tv9) || (_pat_let_tv10));
        } else {
          let _454___mcc_h5 = (_source3).cf6;
          let _455_cf6 = _454___mcc_h5;
          return _pat_let_tv11;
        }
      }(_447_v5);
      let _hi2 = new BigNumber(221);
      for (let _456_i0 = (_this).f0; _456_i0.isLessThan(_hi2); _456_i0 = _456_i0.plus(_dafny.ONE)) {
        let _457_v6;
        let _nw76 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
        _457_v6 = _nw76;
        let _458_v7;
        _458_v7 = _dafny.Map.Empty.slice().updateUnsafe(_439_v1,(_this).f0);
        let _index70 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_457_v6).length));
        (_457_v6)[_index70] = _458_v7;
        let _459_v8;
        _459_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(globalState),_458_v7);
        let _index71 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_457_v6).length));
        (_457_v6)[_index71] = ((((_459_v8).contains(_438_v0)) ? ((_459_v8).get(_438_v0)) : (_458_v7))).update(_439_v1, _456_i0);
        let _460_v9;
        _460_v9 = new BigNumber(715);
        let _461_v10;
        _461_v10 = _dafny.MultiSet.fromElements(_438_v0, !(_438_v0), _438_v0, _438_v0, true);
        _460_v9 = (((_461_v10).contains(_dafny.Seq.IsPrefixOf(_446_v4, _446_v4))) ? ((_461_v10).get(_dafny.Seq.IsPrefixOf(_446_v4, _446_v4))) : (_456_i0));
        let _462_v11;
        _462_v11 = _dafny.Seq.of(_438_v0, _438_v0);
        let _463_v12;
        _463_v12 = new _dafny.CodePoint('t'.codePointAt(0));
        let _464_v13;
        let _out13;
        _out13 = _module.__default.m0(((_this).f0).isLessThan(_456_i0), _module.__default.fm3(_dafny.Set.fromElements(_module.__default.fm2(_438_v0, new BigNumber(-904), _438_v0, _438_v0, globalState), _438_v0, (_462_v11)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_462_v11).length))], _438_v0, _438_v0), _463_v12, _460_v9, _460_v9, globalState), globalState);
        _464_v13 = _out13;
        let _index72 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_439_v1).length));
        (_439_v1)[_index72] = false;
        let _465_v14;
        _465_v14 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f0).isLessThan(_456_i0),_460_v9);
        let _466_v15;
        _466_v15 = _dafny.Map.Empty.slice().updateUnsafe(_465_v14,_462_v11);
        let _467_v16;
        _467_v16 = _dafny.Set.fromElements(_438_v0, _438_v0, true, _464_v13);
        let _index73 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_439_v1).length));
        let _rhs67 = (((_465_v14).contains(true)) ? ((_465_v14).get(true)) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(_460_v9))));
        let _rhs68 = ((_dafny.ZERO).minus(_456_i0)).isLessThanOrEqualTo(_456_i0);
        let _rhs69 = _dafny.Seq.Concat(_462_v11, _dafny.Seq.Concat(_462_v11, (((_466_v15).contains(_465_v14)) ? ((_466_v15).get(_465_v14)) : (_462_v11))));
        let _rhs70 = !((_dafny.Set.fromElements(_464_v13)).equals(_467_v16));
        let _lhs22 = _439_v1;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_439_v1).length));
        _460_v9 = _rhs67;
        _lhs22[_lhs23] = _rhs68;
        _462_v11 = _rhs69;
        _438_v0 = _rhs70;
      }
      let _468_v17;
      _468_v17 = new BigNumber(653);
      let _469_v18;
      _469_v18 = _dafny.Set.fromElements(_438_v0);
      let _470_v19;
      _470_v19 = new _dafny.CodePoint('h'.codePointAt(0));
      let _471_v20;
      _471_v20 = _dafny.Map.Empty.slice().updateUnsafe(_468_v17,(_this).fm5(globalState));
      let _472_v21;
      _472_v21 = _dafny.MultiSet.fromElements(_471_v20, _471_v20);
      _468_v17 = _module.__default.fm3((_469_v18).Union(_469_v18), _470_v19, (_468_v17).plus((((_472_v21).contains(_471_v20)) ? ((_472_v21).get(_471_v20)) : (_468_v17))), (_this).f0, globalState);
      let _hi3 = _468_v17;
      for (let _473_i1 = _468_v17; _473_i1.isLessThan(_hi3); _473_i1 = _473_i1.plus(_dafny.ONE)) {
        let _474_v22;
        _474_v22 = _dafny.MultiSet.fromElements(_473_i1, (_this).f0);
        _446_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_446_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(585)), ((_475_v19) => function (_476_i2) {
          return _475_v19;
        })(_470_v19))), _module.__default.fm19(_474_v22, _473_i1, (_this).f0, _468_v17, globalState));
        let _477_v23;
        let _nw77 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _477_v23 = _nw77;
        let _478_v24;
        _478_v24 = _dafny.Set.fromElements(_477_v23);
        let _rhs71 = _441_v3;
        let _rhs72 = !(_438_v0) || (_438_v0);
        let _rhs73 = new BigNumber((_478_v24).length);
        let _rhs74 = (_dafny.ZERO).minus(_468_v17);
        _441_v3 = _rhs71;
        _438_v0 = _rhs72;
        _468_v17 = _rhs73;
        _468_v17 = _rhs74;
        let _479_v25;
        _479_v25 = _dafny.Map.Empty.slice().updateUnsafe(!(_438_v0),_439_v1);
        _439_v1 = (((_479_v25).contains(_438_v0)) ? ((_479_v25).get(_438_v0)) : (_439_v1));
        let _480_v26;
        let _nw78 = new _module.C0();
        _nw78.__ctor(new _dafny.CodePoint('n'.codePointAt(0)));
        _480_v26 = _nw78;
      }
      let _481_v27;
      let _nw79 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _481_v27 = _nw79;
      let _482_v28;
      _482_v28 = _dafny.MultiSet.fromElements(_481_v27, _481_v27, _481_v27, _481_v27, _481_v27);
      _468_v17 = _module.__default.safeModuloInt((_this).f0, _module.__default.safeModuloInt(new BigNumber((_482_v28).cardinality()), (_this).f0));
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _483_v0;
      _483_v0 = false;
      _483_v0 = _483_v0;
      let _484_v1;
      _484_v1 = _dafny.Set.fromElements((_this).f0, (_this).f0);
      _483_v0 = !((_484_v1).Intersect(_484_v1)).equals(_484_v1);
      let _485_v2;
      let _init9 = function (_486_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      };
      let _nw80 = Array((new BigNumber(20)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw80.length); _i0_9++) {
        _nw80[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _485_v2 = _nw80;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_485_v2).length))) {
        let _487_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_487_i0)) && ((_487_i0).isLessThan(new BigNumber((_485_v2).length))))) {
          (_485_v2)[(_487_i0)] = new _dafny.CodePoint('h'.codePointAt(0));
        }
      }
      let _488_v3;
      _488_v3 = _dafny.MultiSet.fromElements(_483_v0, (_this).fm5(globalState), (_this).fm5(globalState));
      let _489_v4;
      _489_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_488_v3).cardinality()),p1);
      let _490_v5;
      let _nw81 = Array((new BigNumber(20)).toNumber());
      _nw81[(_dafny.ZERO).toNumber()] = false;
      _nw81[(_dafny.ONE).toNumber()] = _483_v0;
      _nw81[(new BigNumber(2)).toNumber()] = _module.__default.fm2(_483_v0, new BigNumber(847), _483_v0, _483_v0, globalState);
      _nw81[(new BigNumber(3)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(4)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(5)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(6)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(7)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(8)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(9)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(10)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(11)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(12)).toNumber()] = true;
      _nw81[(new BigNumber(13)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(14)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(15)).toNumber()] = true;
      _nw81[(new BigNumber(16)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(17)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(18)).toNumber()] = _483_v0;
      _nw81[(new BigNumber(19)).toNumber()] = _483_v0;
      _490_v5 = _nw81;
      let _491_v6;
      _491_v6 = _490_v5;
      let _492_v7;
      _492_v7 = _dafny.Map.Empty.slice().updateUnsafe(_491_v6,new BigNumber(-32));
      let _493_v8;
      let _nw82 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _493_v8 = _nw82;
      let _494_v9;
      _494_v9 = _dafny.MultiSet.fromElements(_493_v8);
      let _495_v10;
      _495_v10 = _dafny.Set.fromElements(_483_v0, _483_v0);
      let _496_v11;
      _496_v11 = new _dafny.CodePoint('k'.codePointAt(0));
      let _497_v12;
      _497_v12 = _dafny.Seq.of(_dafny.Seq.of(_483_v0));
      let _498_v13;
      _498_v13 = _dafny.Seq.UnicodeFromString("ltx");
      let _499_v14;
      _499_v14 = _dafny.Seq.of(true, _483_v0);
      let _500_v15;
      _500_v15 = _module.D7.create_DC16(_499_v14);
      let _501_v16;
      let _nw83 = Array((new BigNumber(21)).toNumber());
      _nw83[(_dafny.ZERO).toNumber()] = new BigNumber(98);
      _nw83[(_dafny.ONE).toNumber()] = new BigNumber((_489_v4).length);
      _nw83[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_this).f0, (_this).f0);
      _nw83[(new BigNumber(3)).toNumber()] = (new BigNumber(124)).multipliedBy((_this).f0);
      _nw83[(new BigNumber(4)).toNumber()] = new BigNumber((_492_v7).length);
      _nw83[(new BigNumber(5)).toNumber()] = p1;
      _nw83[(new BigNumber(6)).toNumber()] = new BigNumber((_494_v9).cardinality());
      _nw83[(new BigNumber(7)).toNumber()] = (_this).f0;
      _nw83[(new BigNumber(8)).toNumber()] = p1;
      _nw83[(new BigNumber(9)).toNumber()] = new BigNumber(71);
      _nw83[(new BigNumber(10)).toNumber()] = (_this).f0;
      _nw83[(new BigNumber(11)).toNumber()] = p1;
      _nw83[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt((_this).f0, (_dafny.ZERO).minus(_module.__default.fm3(_495_v10, _496_v11, new BigNumber((_497_v12).length), p1, globalState)));
      _nw83[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_498_v13, _dafny.Seq.UnicodeFromString("hbu"))).length);
      _nw83[(new BigNumber(14)).toNumber()] = p1;
      _nw83[(new BigNumber(15)).toNumber()] = (((_488_v3).contains(_483_v0)) ? ((_488_v3).get(_483_v0)) : ((_this).f0));
      _nw83[(new BigNumber(16)).toNumber()] = ((_483_v0) ? (new BigNumber(((_500_v15).dtor_cf26).length)) : ((_this).f0));
      _nw83[(new BigNumber(17)).toNumber()] = p1;
      _nw83[(new BigNumber(18)).toNumber()] = (_this).f0;
      _nw83[(new BigNumber(19)).toNumber()] = p1;
      _nw83[(new BigNumber(20)).toNumber()] = (_this).f0;
      _501_v16 = _nw83;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_501_v16).length))) {
        let _502_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_502_i2)) && ((_502_i2).isLessThan(new BigNumber((_501_v16).length))))) {
          (_501_v16)[(_502_i2)] = (_502_i2).plus(p1);
        }
      }
      let _503_v17;
      _503_v17 = _module.D5.create_DC13(p1, _495_v10, _496_v11, new BigNumber(701), _483_v0);
      let _504_i3;
      _504_i3 = _dafny.ZERO;
      L5: {
        while ((_503_v17).dtor_cf21) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_504_i3)) {
              break L5;
            }
            _504_i3 = (_504_i3).plus(_dafny.ONE);
            let _505_v18;
            _505_v18 = _dafny.Seq.of(_493_v8);
            let _506_v19;
            _506_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_505_v18, _505_v18),_483_v0);
            let _507_v20;
            _507_v20 = _module.D2.create_DC4(_498_v13, !(_483_v0));
            _506_v19 = (_506_v19).update(_505_v18, (_507_v20).dtor_cf5);
            let _508_v21;
            _508_v21 = _dafny.Map.Empty.slice().updateUnsafe(_483_v0,_496_v11);
            let _509_v22;
            _509_v22 = _dafny.Seq.of((_this).f0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_483_v0,_483_v0)).length), new BigNumber((_module.__default.fm20(_483_v0, _508_v21, true, _496_v11, globalState)).length), (_this).f0, p1);
            let _rhs75 = !((_this).f0).isEqualTo(_module.__default.safeModuloInt((_this).f0, (_this).f0));
            let _rhs76 = _496_v11;
            let _rhs77 = p1;
            let _rhs78 = _dafny.areEqual(_509_v22, _509_v22);
            let _rhs79 = _module.__default.safeDivisionInt((_this).f0, (_this).f0);
            _483_v0 = _rhs75;
            _496_v11 = _rhs76;
            r1 = _rhs77;
            _483_v0 = _rhs78;
            r2 = _rhs79;
            _483_v0 = _module.__default.fm2((_483_v0) && (_483_v0), ((_this).f0).plus((_this).f0), !(_483_v0), _483_v0, globalState);
            _483_v0 = (new BigNumber(357)).isLessThanOrEqualTo(new BigNumber((_495_v10).length));
          }
        }
      }
      let _hi4 = p1;
      for (let _510_i4 = new BigNumber((_dafny.Seq.of((_this).f0)).length); _510_i4.isLessThan(_hi4); _510_i4 = _510_i4.plus(_dafny.ONE)) {
        _483_v0 = ((_dafny.ZERO).minus((_this).f0)).isEqualTo(p1);
        _483_v0 = _483_v0;
        if (_483_v0) {
          let _511_v23;
          _511_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_498_v13).length),_483_v0);
          let _512_v24;
          _512_v24 = _dafny.Map.Empty.slice().updateUnsafe(_483_v0,new BigNumber((_module.__default.fm19(_module.__default.fm22(globalState), _510_i4, (_this).f0, new BigNumber((_511_v23).length), globalState)).length));
          let _513_v25;
          _513_v25 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_483_v0,(_this).f0), _module.__default.fm21(globalState), _512_v24);
          _513_v25 = _513_v25;
          let _rhs80 = _510_i4;
          let _rhs81 = _496_v11;
          let _rhs82 = ((_510_i4).minus(p1)).isLessThan(p1);
          r1 = _rhs80;
          _496_v11 = _rhs81;
          _483_v0 = _rhs82;
          _483_v0 = _483_v0;
          r2 = _module.__default.fm3(_495_v10, _496_v11, p1, (p1).plus(new BigNumber(-34)), globalState);
          r0 = _510_i4;
        } else {
          r2 = _510_i4;
          _493_v8 = _493_v8;
          let _514_v26;
          _514_v26 = _dafny.MultiSet.fromElements(_488_v3);
          let _515_v27;
          _515_v27 = _dafny.MultiSet.fromElements(new BigNumber((_498_v13).length), p1, (((_514_v26).contains(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, _483_v0)))) ? ((_514_v26).get(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, _483_v0)))) : (new BigNumber((_498_v13).length))));
          let _516_v28;
          _516_v28 = _dafny.Map.Empty.slice().updateUnsafe(((_483_v0) ? (_490_v5) : (_490_v5)),_515_v27);
          _516_v28 = (_516_v28).Merge(_516_v28);
          _493_v8 = _501_v16;
          let _517_v29;
          let _nw84 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _517_v29 = _nw84;
          let _index74 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_517_v29).length));
          (_517_v29)[_index74] = _dafny.Seq.of(_510_i4);
          let _518_v30;
          _518_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_483_v0);
          let _519_v31;
          _519_v31 = _dafny.Seq.of(new BigNumber((_488_v3).cardinality()), _module.__default.safeDivisionInt(new BigNumber((_518_v30).length), (_this).f0));
          let _index75 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_517_v29).length));
          (_517_v29)[_index75] = _519_v31;
        }
        if (!(_510_i4).isEqualTo(_510_i4)) {
          let _520_v32;
          _520_v32 = _dafny.MultiSet.fromElements(_510_i4);
          _483_v0 = ((new BigNumber((_520_v32).cardinality())).multipliedBy((_this).f0)).isEqualTo(p1);
          r0 = ((false) ? ((_this).f0) : (p1));
          let _index76 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_490_v5).length));
          (_490_v5)[_index76] = !((_this).fm6(_483_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-317)), function (_521_i5) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _498_v13, _510_i4, globalState));
          let _index77 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_490_v5).length));
          (_490_v5)[_index77] = ((_483_v0) ? (!((_510_i4).isLessThanOrEqualTo(new BigNumber(-183)))) : (_483_v0));
          r1 = new BigNumber((_498_v13).length);
          r0 = _module.__default.fm3((_495_v10).Intersect(_495_v10), _496_v11, p1, ((_this).f0).minus(p1), globalState);
        } else {
          _496_v11 = _496_v11;
          let _522_v33;
          _522_v33 = _dafny.Map.Empty.slice().updateUnsafe(_483_v0,_493_v8);
          _522_v33 = (_522_v33).update(_483_v0, _501_v16);
          let _index78 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_490_v5).length));
          (_490_v5)[_index78] = _483_v0;
          let _index79 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_490_v5).length));
          (_490_v5)[_index79] = ((false) ? (_483_v0) : ((_499_v14)[_module.__default.safeIndex(p1, new BigNumber((_499_v14).length))]));
          let _523_v34;
          _523_v34 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC18(p1),(_490_v5)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_490_v5).length))]);
          let _524_v36;
          _524_v36 = _module.D7.create_DC18((_this).f0);
          let _525_v37;
          _525_v37 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, (_this).f0)),((_483_v0) ? (_523_v34) : (function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(_524_v36,_483_v0)).Keys.Elements) {
              let _526_v35 = _compr_15;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_524_v36,_483_v0)).contains(_526_v35)) {
                _coll15.push([_526_v35,(_490_v5)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_490_v5).length))]]);
              }
            }
            return _coll15;
          }())));
          _525_v37 = (_525_v37).Merge((_525_v37).Merge(_525_v37));
          _496_v11 = _496_v11;
        }
      }
      let _527_v38;
      _527_v38 = _dafny.Seq.of(new BigNumber((_498_v13).length));
      r0 = new BigNumber((_dafny.Seq.Concat(_527_v38, _527_v38)).length);
      r1 = p1;
      r2 = p1;
      return [r0, r1, r2];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f0 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return !(!(_dafny.areEqual(_dafny.Seq.UnicodeFromString("jdrpodwh"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_528_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("hehmjmhcu")))));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (true) {
        return !(((false) ? (false) : (true)));
      } else {
        return (true) && (true);
      }
    };
    fm13(p0, p1, globalState) {
      let _this = this;
      return true;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _529_v0;
      _529_v0 = _dafny.Set.fromElements(p1, (_this).f0);
      let _530_v1;
      _530_v1 = _dafny.Set.fromElements(_529_v0, _529_v0, _529_v0);
      let _531_v2;
      _531_v2 = true;
      let _532_v3;
      _532_v3 = _dafny.Seq.of(_530_v1);
      let _533_v4;
      _533_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,_531_v2);
      let _534_v5;
      _534_v5 = _dafny.Set.fromElements(_531_v2);
      let _535_v6;
      _535_v6 = new _dafny.CodePoint('j'.codePointAt(0));
      let _536_v7;
      _536_v7 = _dafny.Map.Empty.slice().updateUnsafe(_529_v0,_531_v2);
      let _537_v9;
      _537_v9 = _dafny.Seq.of(_529_v0, _529_v0);
      let _538_v11;
      let _nw85 = Array((new BigNumber(21)).toNumber());
      _nw85[(_dafny.ZERO).toNumber()] = _530_v1;
      _nw85[(_dafny.ONE).toNumber()] = _530_v1;
      _nw85[(new BigNumber(2)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(3)).toNumber()] = (_530_v1).Intersect(_530_v1);
      _nw85[(new BigNumber(4)).toNumber()] = (_530_v1).Union(_530_v1);
      _nw85[(new BigNumber(5)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(6)).toNumber()] = ((_531_v2) ? ((_532_v3)[_module.__default.safeIndex(new BigNumber((_533_v4).length), new BigNumber((_532_v3).length))]) : (_530_v1));
      _nw85[(new BigNumber(7)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(8)).toNumber()] = (_532_v3)[_module.__default.safeIndex(_module.__default.fm3(_534_v5, _535_v6, (_this).f0, p1, globalState), new BigNumber((_532_v3).length))];
      _nw85[(new BigNumber(9)).toNumber()] = function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_536_v7).Keys.Elements) {
          let _539_v8 = _compr_16;
          if ((_536_v7).contains(_539_v8)) {
            _coll16.add(_539_v8);
          }
        }
        return _coll16;
      }();
      _nw85[(new BigNumber(10)).toNumber()] = (_dafny.Set.fromElements(_529_v0)).Intersect(_dafny.Set.fromElements(_dafny.Set.fromElements(p1), _dafny.Set.fromElements(p1), _529_v0, _529_v0, _529_v0));
      _nw85[(new BigNumber(11)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(12)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(_529_v0, _529_v0, _529_v0, _529_v0);
      _nw85[(new BigNumber(14)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(_529_v0);
      _nw85[(new BigNumber(16)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(17)).toNumber()] = function () {
        let _coll17 = new _dafny.Set();
        for (const _compr_17 of (_537_v9).Elements) {
          let _540_v10 = _compr_17;
          if (_dafny.Seq.contains(_537_v9, _540_v10)) {
            _coll17.add(_540_v10);
          }
        }
        return _coll17;
      }();
      _nw85[(new BigNumber(18)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(19)).toNumber()] = _530_v1;
      _nw85[(new BigNumber(20)).toNumber()] = (_530_v1).Union(_530_v1);
      _538_v11 = _nw85;
      let _index80 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_538_v11).length));
      (_538_v11)[_index80] = _dafny.Set.fromElements(_529_v0);
      let _541_v12;
      _541_v12 = _module.D5.create_DC11(_dafny.Set.fromElements(_529_v0, _529_v0));
      let _index81 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_538_v11).length));
      (_538_v11)[_index81] = (_541_v12).dtor_cf14;
      let _source4 = _module.__default.fm14((_this).f0, _531_v2, _535_v6, globalState);
      if (_source4.is_DC4) {
        let _542___mcc_h0 = (_source4).cf4;
        let _543___mcc_h1 = (_source4).cf5;
        let _544_cf5 = _543___mcc_h1;
        let _545_cf4 = _542___mcc_h0;
        let _546_v13;
        let _nw86 = Array((new BigNumber(25)).toNumber()).fill(false);
        _546_v13 = _nw86;
        let _547_v14;
        _547_v14 = _dafny.Seq.of(false, true, _544_cf5, _544_cf5);
        let _index82 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_546_v13).length));
        (_546_v13)[_index82] = !(!(_531_v2) || ((_547_v14)[_module.__default.safeIndex(p1, new BigNumber((_547_v14).length))]));
        let _index83 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_546_v13).length));
        (_546_v13)[_index83] = _531_v2;
        let _548_v15;
        _548_v15 = _dafny.Map.Empty.slice().updateUnsafe(_531_v2,_545_cf4);
        _548_v15 = (_548_v15).update((_dafny.Set.fromElements(p1, (_this).f0)).IsProperSubsetOf(_529_v0), _545_cf4);
        r2 = new BigNumber((_545_cf4).length);
        _546_v13 = _546_v13;
      } else if (_source4.is_DC5) {
        let _549_v16;
        _549_v16 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        let _550_v17;
        _550_v17 = _dafny.Seq.of(_531_v2, _531_v2, _531_v2);
        let _551_v18;
        _551_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_550_v17).length));
        let _552_v19;
        _552_v19 = _dafny.Seq.of(_module.__default.fm3(_534_v5, _535_v6, new BigNumber((_551_v18).length), (_this).f0, globalState));
        let _553_v20;
        _553_v20 = _dafny.Map.Empty.slice().updateUnsafe(p1,_535_v6);
        let _554_v21;
        let _nw87 = Array((new BigNumber(26)).toNumber());
        _nw87[(_dafny.ZERO).toNumber()] = (((_549_v16).contains(_531_v2)) ? ((_549_v16).get(_531_v2)) : (new BigNumber(-553)));
        _nw87[(_dafny.ONE).toNumber()] = p1;
        _nw87[(new BigNumber(2)).toNumber()] = p1;
        _nw87[(new BigNumber(3)).toNumber()] = (((_this).fm5(globalState)) ? ((_this).f0) : ((_this).f0));
        _nw87[(new BigNumber(4)).toNumber()] = (_552_v19)[_module.__default.safeIndex((_this).f0, new BigNumber((_552_v19).length))];
        _nw87[(new BigNumber(5)).toNumber()] = p1;
        _nw87[(new BigNumber(6)).toNumber()] = ((false) ? ((_this).f0) : (new BigNumber(345)));
        _nw87[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw87[(new BigNumber(8)).toNumber()] = new BigNumber(610);
        _nw87[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw87[(new BigNumber(10)).toNumber()] = _module.__default.fm3(_534_v5, (((_553_v20).contains(p1)) ? ((_553_v20).get(p1)) : (_535_v6)), (_this).f0, (_this).f0, globalState);
        _nw87[(new BigNumber(11)).toNumber()] = (_this).f0;
        _nw87[(new BigNumber(12)).toNumber()] = (_this).f0;
        _nw87[(new BigNumber(13)).toNumber()] = new BigNumber((_529_v0).length);
        _nw87[(new BigNumber(14)).toNumber()] = (_this).f0;
        _nw87[(new BigNumber(15)).toNumber()] = new BigNumber(-41);
        _nw87[(new BigNumber(16)).toNumber()] = p1;
        _nw87[(new BigNumber(17)).toNumber()] = new BigNumber((_552_v19).length);
        _nw87[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_this).f0);
        _nw87[(new BigNumber(19)).toNumber()] = p1;
        _nw87[(new BigNumber(20)).toNumber()] = new BigNumber(808);
        _nw87[(new BigNumber(21)).toNumber()] = p1;
        _nw87[(new BigNumber(22)).toNumber()] = (_this).f0;
        _nw87[(new BigNumber(23)).toNumber()] = p1;
        _nw87[(new BigNumber(24)).toNumber()] = (_this).f0;
        _nw87[(new BigNumber(25)).toNumber()] = p1;
        _554_v21 = _nw87;
        let _index84 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length));
        (_554_v21)[_index84] = p1;
        let _index85 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length));
        (_554_v21)[_index85] = (p1).plus(_module.__default.safeDivisionInt(p1, p1));
        let _555_v22;
        _555_v22 = _dafny.MultiSet.fromElements(_531_v2, false);
        let _556_v23;
        _556_v23 = _dafny.MultiSet.fromElements(p1, p1, (_554_v21)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length))]);
        let _557_v24;
        let _nw88 = new _module.C0();
        _nw88.__ctor(new _dafny.CodePoint('v'.codePointAt(0)));
        _557_v24 = _nw88;
        let _558_v25;
        _558_v25 = _dafny.Map.Empty.slice().updateUnsafe(_557_v24,_531_v2);
        let _559_v26;
        _559_v26 = _dafny.Map.Empty.slice().updateUnsafe((((_558_v25).contains(_557_v24)) ? ((_558_v25).get(_557_v24)) : (_module.__default.fm2(_531_v2, p1, _531_v2, _531_v2, globalState))),true);
        let _560_v27;
        _560_v27 = _module.D0.create_DC0(_531_v2);
        let _561_v28;
        let _nw89 = Array((new BigNumber(13)).toNumber());
        _nw89[(_dafny.ZERO).toNumber()] = (_560_v27).dtor_cf0;
        _nw89[(_dafny.ONE).toNumber()] = true;
        _nw89[(new BigNumber(2)).toNumber()] = false;
        _nw89[(new BigNumber(3)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(4)).toNumber()] = (_this).fm5(globalState);
        _nw89[(new BigNumber(5)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(6)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(7)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(8)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(9)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(10)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(11)).toNumber()] = _531_v2;
        _nw89[(new BigNumber(12)).toNumber()] = _531_v2;
        _561_v28 = _nw89;
        let _562_v29;
        _562_v29 = _dafny.MultiSet.fromElements(_561_v28);
        let _563_v30;
        let _nw90 = Array((new BigNumber(26)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = (_555_v22).IsDisjointFrom(_dafny.MultiSet.FromArray(_550_v17));
        _nw90[(_dafny.ONE).toNumber()] = ((_531_v2) ? (_531_v2) : (_531_v2));
        _nw90[(new BigNumber(2)).toNumber()] = (false) === (_531_v2);
        _nw90[(new BigNumber(3)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(4)).toNumber()] = (_550_v17)[_module.__default.safeIndex((_554_v21)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length))], new BigNumber((_550_v17).length))];
        _nw90[(new BigNumber(5)).toNumber()] = !(_531_v2) || (_531_v2);
        _nw90[(new BigNumber(6)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(7)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(8)).toNumber()] = (_550_v17)[_module.__default.safeIndex(p1, new BigNumber((_550_v17).length))];
        _nw90[(new BigNumber(9)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(10)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(11)).toNumber()] = (_556_v23).IsDisjointFrom(_556_v23);
        _nw90[(new BigNumber(12)).toNumber()] = (_550_v17)[_module.__default.safeIndex((_this).f0, new BigNumber((_550_v17).length))];
        _nw90[(new BigNumber(13)).toNumber()] = ((_550_v17)[_module.__default.safeIndex((_554_v21)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length))], new BigNumber((_550_v17).length))]) || ((((_559_v26).contains(_531_v2)) ? ((_559_v26).get(_531_v2)) : (false)));
        _nw90[(new BigNumber(14)).toNumber()] = (_531_v2) === (_531_v2);
        _nw90[(new BigNumber(15)).toNumber()] = true;
        _nw90[(new BigNumber(16)).toNumber()] = (_dafny.MultiSet.fromElements(_561_v28)).IsProperSubsetOf(_562_v29);
        _nw90[(new BigNumber(17)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(18)).toNumber()] = (_this).fm5(globalState);
        _nw90[(new BigNumber(19)).toNumber()] = (_529_v0).IsProperSubsetOf(_529_v0);
        _nw90[(new BigNumber(20)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(21)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(22)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(23)).toNumber()] = _531_v2;
        _nw90[(new BigNumber(24)).toNumber()] = !((_this).fm5(globalState));
        _nw90[(new BigNumber(25)).toNumber()] = _531_v2;
        _563_v30 = _nw90;
        let _564_v31;
        _564_v31 = _dafny.Seq.UnicodeFromString("ykehlcy");
        let _index86 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_561_v28).length));
        (_561_v28)[_index86] = !_dafny.Seq.contains(_564_v31, _535_v6);
        let _565_v32;
        _565_v32 = _dafny.Seq.of(_554_v21, _554_v21, _554_v21);
        let _566_v33;
        _566_v33 = _module.D2.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(382)), function (_567_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}), _531_v2);
        let _index87 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_561_v28).length));
        let _index88 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length));
        let _rhs83 = !(!(_531_v2) || (!_dafny.Seq.contains(_565_v32, _554_v21)));
        let _rhs84 = _module.__default.fm3(_534_v5, ((_531_v2) ? ((_557_v24).f3) : ((_557_v24).f3)), p1, p1, globalState);
        let _rhs85 = (new BigNumber(((_566_v33).dtor_cf4).length)).isLessThanOrEqualTo((_this).f0);
        let _rhs86 = _531_v2;
        let _rhs87 = (_dafny.ZERO).minus(new BigNumber((_550_v17).length));
        let _lhs24 = _561_v28;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_561_v28).length));
        let _lhs26 = _554_v21;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_554_v21).length));
        _531_v2 = _rhs83;
        r2 = _rhs84;
        _lhs24[_lhs25] = _rhs85;
        _531_v2 = _rhs86;
        _lhs26[_lhs27] = _rhs87;
        let _nw91 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _554_v21 = _nw91;
        let _568_v34;
        let _nw92 = new _module.C0();
        _nw92.__ctor((_557_v24).f3);
        _568_v34 = _nw92;
      } else if (_source4.is_DC3) {
        let _569___mcc_h2 = (_source4).cf3;
        let _570_cf3 = _569___mcc_h2;
        let _571_v35;
        let _init10 = ((_572_p1) => function (_573_i1) {
          return (_573_i1).minus(_572_p1);
        })(p1);
        let _nw93 = Array((new BigNumber(5)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw93.length); _i0_10++) {
          _nw93[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _571_v35 = _nw93;
        let _index89 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
        (_571_v35)[_index89] = _module.__default.safeModuloInt(p1, p1);
        let _574_v36;
        _574_v36 = _dafny.MultiSet.fromElements(_571_v35);
        let _575_v37;
        _575_v37 = _module.D6.create_DC14(_571_v35);
        let _576_v38;
        let _nw94 = new _module.C0();
        _nw94.__ctor(new _dafny.CodePoint('a'.codePointAt(0)));
        _576_v38 = _nw94;
        let _577_v39;
        _577_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,_576_v38);
        let _578_v40;
        _578_v40 = _dafny.Map.Empty.slice().updateUnsafe(_531_v2,p1);
        let _579_v41;
        _579_v41 = _dafny.Seq.of(_576_v38, _576_v38);
        let _580_v42;
        let _nw95 = Array((new BigNumber(26)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = _576_v38;
        _nw95[(_dafny.ONE).toNumber()] = _576_v38;
        _nw95[(new BigNumber(2)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(3)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(4)).toNumber()] = (((_577_v39).contains((((_578_v40).contains(_531_v2)) ? ((_578_v40).get(_531_v2)) : (p1)))) ? ((_577_v39).get((((_578_v40).contains(_531_v2)) ? ((_578_v40).get(_531_v2)) : (p1)))) : (_576_v38));
        _nw95[(new BigNumber(5)).toNumber()] = (_579_v41)[_module.__default.safeIndex(p1, new BigNumber((_579_v41).length))];
        _nw95[(new BigNumber(6)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(7)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(8)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(9)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(10)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(11)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(12)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(13)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(14)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(15)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(16)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(17)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(18)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(19)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(20)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(21)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(22)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(23)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(24)).toNumber()] = _576_v38;
        _nw95[(new BigNumber(25)).toNumber()] = _576_v38;
        _580_v42 = _nw95;
        let _581_v43;
        _581_v43 = _module.D3.create_DC9((_this).f0, _580_v42, _531_v2);
        let _pat_let_tv12 = _571_v35;
        let _pat_let_tv13 = _531_v2;
        let _index90 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
        let _rhs88 = (_574_v36).IsProperSubsetOf((_574_v36).Intersect(_dafny.MultiSet.fromElements((function (_pat_let6_0) {
          return function (_582_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_583_dt__update_hcf22_h0) {
                return _module.D6.create_DC14(_583_dt__update_hcf22_h0);
              }(_pat_let7_0);
            }(_pat_let_tv12);
          }(_pat_let6_0);
        }(_575_v37)).dtor_cf22, _571_v35, _571_v35, _571_v35)));
        let _rhs89 = ((_this).f0).multipliedBy((function (_pat_let8_0) {
          return function (_584_dt__update__tmp_h1) {
            return function (_pat_let9_0) {
              return function (_585_dt__update_hcf10_h0) {
                return function (_pat_let10_0) {
                  return function (_586_dt__update_hcf12_h0) {
                    return _module.D3.create_DC9(_585_dt__update_hcf10_h0, (_584_dt__update__tmp_h1).dtor_cf11, _586_dt__update_hcf12_h0);
                  }(_pat_let10_0);
                }(_pat_let_tv13);
              }(_pat_let9_0);
            }((_this).f0);
          }(_pat_let8_0);
        }(_581_v43)).dtor_cf10);
        let _lhs28 = _571_v35;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
        _531_v2 = _rhs88;
        _lhs28[_lhs29] = _rhs89;
        let _587_v44;
        _587_v44 = _dafny.Seq.UnicodeFromString("arknopxms");
        let _588_v45;
        _588_v45 = _module.D2.create_DC4(_587_v44, _531_v2);
        let _589_v46;
        _589_v46 = _dafny.Set.fromElements(_588_v45, _588_v45, _588_v45);
        if ((_dafny.Set.fromElements(_588_v45)).IsDisjointFrom(_589_v46)) {
          _531_v2 = _531_v2;
          _531_v2 = ((!((_this).f0).isEqualTo((_this).f0)) ? (_531_v2) : (_531_v2));
          _531_v2 = !(_module.__default.fm2(_531_v2, ((_this).f0).multipliedBy((_this).f0), true, false, globalState));
          let _590_v47;
          _590_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm18((_571_v35)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length))], globalState)).length),(_this).f0);
          let _591_v48;
          let _nw96 = new _module.C1();
          _nw96.__ctor((((_590_v47).contains(p1)) ? ((_590_v47).get(p1)) : (_module.__default.fm3(_534_v5, _535_v6, (_this).f0, p1, globalState))));
          _591_v48 = _nw96;
          let _index91 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
          let _rhs90 = _591_v48;
          let _rhs91 = new BigNumber((_dafny.Seq.of((_dafny.Set.fromElements(_531_v2)).Intersect(_dafny.Set.fromElements((_this).fm6(_531_v2, _587_v44, _587_v44, (_591_v48).f0, globalState), _531_v2)), _534_v5, _534_v5)).length);
          let _lhs30 = _571_v35;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
          _591_v48 = _rhs90;
          _lhs30[_lhs31] = _rhs91;
          let _592_v49;
          let _nw97 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _592_v49 = _nw97;
          _592_v49 = _592_v49;
        } else {
          let _593_v50;
          let _nw98 = new _module.C0();
          _nw98.__ctor((_576_v38).f3);
          _593_v50 = _nw98;
          let _594_v51;
          let _nw99 = Array((new BigNumber(11)).toNumber()).fill(false);
          _594_v51 = _nw99;
          let _index92 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_594_v51).length));
          (_594_v51)[_index92] = _531_v2;
          let _index93 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_594_v51).length));
          (_594_v51)[_index93] = _531_v2;
          _531_v2 = (_594_v51)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_594_v51).length))];
          _588_v45 = _module.D2.create_DC4(_587_v44, _531_v2);
          r1 = ((_571_v35)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length))]).multipliedBy((_571_v35)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length))]);
        }
        let _595_v52;
        _595_v52 = _module.D0.create_DC1(_531_v2);
        let _596_v53;
        _596_v53 = _dafny.Seq.of((_this).f0);
        let _pat_let_tv14 = _531_v2;
        let _index94 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
        let _rhs92 = (_587_v44)[_module.__default.safeIndex((_571_v35)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length))], new BigNumber((_587_v44).length))];
        let _rhs93 = (function (_pat_let11_0) {
          return function (_597_dt__update__tmp_h2) {
            return function (_pat_let12_0) {
              return function (_598_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_598_dt__update_hcf1_h0);
              }(_pat_let12_0);
            }(_pat_let_tv14);
          }(_pat_let11_0);
        }(_595_v52)).dtor_cf1;
        let _rhs94 = new BigNumber((_596_v53).length);
        let _lhs32 = _571_v35;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
        _535_v6 = _rhs92;
        _531_v2 = _rhs93;
        _lhs32[_lhs33] = _rhs94;
        let _index95 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_571_v35).length));
        (_571_v35)[_index95] = new BigNumber(193);
      } else {
        let _599___mcc_h3 = (_source4).cf6;
        let _600_cf6 = _599___mcc_h3;
        let _601_v54;
        let _nw100 = new _module.C0();
        _nw100.__ctor(_535_v6);
        _601_v54 = _nw100;
        let _602_v55;
        let _nw101 = new _module.C1();
        _nw101.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_601_v54)).length));
        _602_v55 = _nw101;
        r0 = (_dafny.ZERO).minus((_this).f0);
        let _603_v56;
        let _nw102 = new _module.C0();
        _nw102.__ctor((_601_v54).f3);
        _603_v56 = _nw102;
        r1 = (_this).f0;
      }
      let _604_i2;
      _604_i2 = _dafny.ZERO;
      L6: {
        while (_531_v2) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_604_i2)) {
              break L6;
            }
            _604_i2 = (_604_i2).plus(_dafny.ONE);
            let _605_v57;
            let _nw103 = Array((new BigNumber(16)).toNumber());
            _605_v57 = _nw103;
            let _606_v58;
            _606_v58 = _module.D3.create_DC9(_module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ffshwyvs"))).cardinality()), new BigNumber(364)), _605_v57, _531_v2);
            let _source5 = _606_v58;
            if (_source5.is_DC8) {
              let _607___mcc_h4 = (_source5).cf8;
              let _608___mcc_h5 = (_source5).cf9;
              let _609_cf9 = _608___mcc_h5;
              let _610_cf8 = _607___mcc_h4;
              let _611_v59;
              _611_v59 = _dafny.Seq.UnicodeFromString("qpgrnxwqu");
              let _612_v60;
              let _out14;
              _out14 = _module.__default.m0((_this).fm6(_531_v2, _611_v59, _611_v59, _module.__default.fm3(_534_v5, _535_v6, new BigNumber((_534_v5).length), (_this).f0, globalState), globalState), p1, globalState);
              _612_v60 = _out14;
              r1 = _609_cf9;
              let _613_v61;
              let _nw104 = new _module.C1();
              _nw104.__ctor((_this).f0);
              _613_v61 = _nw104;
              let _614_v62;
              _614_v62 = _dafny.Map.Empty.slice().updateUnsafe(_612_v60,_612_v60);
              _614_v62 = (_614_v62).update(_531_v2, _531_v2);
            } else if (_source5.is_DC9) {
              let _615___mcc_h6 = (_source5).cf10;
              let _616___mcc_h7 = (_source5).cf11;
              let _617___mcc_h8 = (_source5).cf12;
              let _618_cf12 = _617___mcc_h8;
              let _619_cf11 = _616___mcc_h7;
              let _620_cf10 = _615___mcc_h6;
              let _621_v63;
              let _nw105 = new _module.C2();
              _nw105.__ctor(p1);
              _621_v63 = _nw105;
              let _622_v64;
              _622_v64 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f0).isLessThanOrEqualTo(p1),_621_v63);
              _622_v64 = (_622_v64).update(false, (((_622_v64).contains((_621_v63).fm5(globalState))) ? ((_622_v64).get((_621_v63).fm5(globalState))) : (_621_v63)));
              let _623_v65;
              let _init11 = ((_624_v0) => function (_625_i3) {
                return _624_v0;
              })(_529_v0);
              let _nw106 = Array((new BigNumber(19)).toNumber());
              for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw106.length); _i0_11++) {
                _nw106[_i0_11] = _init11(new BigNumber(_i0_11));
              }
              _623_v65 = _nw106;
              _623_v65 = _623_v65;
              let _626_v66;
              let _nw107 = Array((new BigNumber(26)).toNumber()).fill(false);
              _626_v66 = _nw107;
              let _627_v67;
              _627_v67 = _dafny.MultiSet.fromElements((_621_v63).f0, _620_cf10);
              let _index96 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_626_v66).length));
              (_626_v66)[_index96] = ((((_627_v67).contains(_620_cf10)) ? ((_627_v67).get(_620_cf10)) : ((_this).f0))).isLessThan((_this).f0);
              let _index97 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_626_v66).length));
              (_626_v66)[_index97] = true;
              r2 = (_dafny.ZERO).minus((p1).multipliedBy(_620_cf10));
            } else {
              let _628___mcc_h9 = (_source5).cf7;
              let _629_cf7 = _628___mcc_h9;
              let _630_v68;
              let _nw108 = new _module.C0();
              _nw108.__ctor(_535_v6);
              _630_v68 = _nw108;
              _531_v2 = _531_v2;
              r0 = (p1).plus((_this).f0);
              let _631_v69;
              let _init12 = ((_632_cf7) => function (_633_i4) {
                return (_633_i4).multipliedBy(new BigNumber((_632_cf7).length));
              })(_629_cf7);
              let _nw109 = Array((new BigNumber(7)).toNumber());
              for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw109.length); _i0_12++) {
                _nw109[_i0_12] = _init12(new BigNumber(_i0_12));
              }
              _631_v69 = _nw109;
              let _634_v70;
              _634_v70 = _dafny.MultiSet.fromElements(_631_v69, _631_v69, _631_v69, _631_v69, _631_v69);
              _634_v70 = _dafny.MultiSet.fromElements(_631_v69, _631_v69);
            }
            let _635_v71;
            let _init13 = ((_636_p1) => function (_637_i5) {
              return (_637_i5).multipliedBy(_636_p1);
            })(p1);
            let _nw110 = Array((new BigNumber(5)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw110.length); _i0_13++) {
              _nw110[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _635_v71 = _nw110;
            let _638_v72;
            _638_v72 = _dafny.Map.Empty.slice().updateUnsafe(p1,_635_v71);
            _638_v72 = _638_v72;
            let _639_v73;
            _639_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,p1);
            let _640_v74;
            _640_v74 = _dafny.Map.Empty.slice().updateUnsafe(_535_v6,p1);
            let _641_v75;
            let _nw111 = Array((new BigNumber(10)).toNumber());
            _nw111[(_dafny.ZERO).toNumber()] = _531_v2;
            _nw111[(_dafny.ONE).toNumber()] = _531_v2;
            _nw111[(new BigNumber(2)).toNumber()] = !((_639_v73).update(new BigNumber((_640_v74).length), (_this).f0)).equals(_639_v73);
            _nw111[(new BigNumber(3)).toNumber()] = true;
            _nw111[(new BigNumber(4)).toNumber()] = _531_v2;
            _nw111[(new BigNumber(5)).toNumber()] = true;
            _nw111[(new BigNumber(6)).toNumber()] = true;
            _nw111[(new BigNumber(7)).toNumber()] = false;
            _nw111[(new BigNumber(8)).toNumber()] = _531_v2;
            _nw111[(new BigNumber(9)).toNumber()] = _531_v2;
            _641_v75 = _nw111;
            let _642_v76;
            _642_v76 = _dafny.Seq.of((_this).fm13((_this).f0, false, globalState), _531_v2);
            let _index98 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_641_v75).length));
            (_641_v75)[_index98] = (_642_v76)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_639_v73).length)), new BigNumber((_642_v76).length))];
            let _index99 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_641_v75).length));
            (_641_v75)[_index99] = (_531_v2) === (_531_v2);
            let _643_v77;
            _643_v77 = _dafny.MultiSet.fromElements(_531_v2, true);
            let _644_v78;
            _644_v78 = _dafny.Seq.UnicodeFromString("dvejd");
            let _645_v79;
            _645_v79 = _dafny.MultiSet.fromElements(new BigNumber(204), (_this).f0);
            let _646_v80;
            _646_v80 = _dafny.MultiSet.fromElements(_643_v77, _643_v77);
            let _index100 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_641_v75).length));
            let _index101 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_641_v75).length));
            let _rhs95 = !(new BigNumber(((_643_v77).Union(_643_v77)).cardinality())).isEqualTo(p1);
            let _rhs96 = p1;
            let _rhs97 = _606_v58;
            let _rhs98 = new BigNumber((((_534_v5).Union(_dafny.Set.fromElements(_531_v2))).Union(_534_v5)).length);
            let _rhs99 = (new BigNumber((_dafny.Seq.Concat(_644_v78, _module.__default.fm19(_645_v79, (_this).f0, (_this).f0, (_this).f0, globalState))).length)).isEqualTo(_module.__default.fm3(_534_v5, _535_v6, (_this).f0, new BigNumber((_646_v80).cardinality()), globalState));
            let _lhs34 = _641_v75;
            let _lhs35 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_641_v75).length));
            let _lhs36 = _641_v75;
            let _lhs37 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_641_v75).length));
            _lhs34[_lhs35] = _rhs95;
            r0 = _rhs96;
            _606_v58 = _rhs97;
            r0 = _rhs98;
            _lhs36[_lhs37] = _rhs99;
            let _647_v81;
            let _nw112 = Array((new BigNumber(5)).toNumber());
            _nw112[(_dafny.ZERO).toNumber()] = _535_v6;
            _nw112[(_dafny.ONE).toNumber()] = ((_531_v2) ? (_535_v6) : (_535_v6));
            _nw112[(new BigNumber(2)).toNumber()] = _535_v6;
            _nw112[(new BigNumber(3)).toNumber()] = _535_v6;
            _nw112[(new BigNumber(4)).toNumber()] = _535_v6;
            _647_v81 = _nw112;
            let _index102 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_647_v81).length));
            (_647_v81)[_index102] = new _dafny.CodePoint('b'.codePointAt(0));
            let _index103 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_647_v81).length));
            (_647_v81)[_index103] = _535_v6;
          }
        }
      }
      let _648_i6;
      _648_i6 = _dafny.ZERO;
      L7: {
        while ((_this).fm13((_dafny.ZERO).minus((_module.__default.fm3(_534_v5, _535_v6, p1, p1, globalState)).minus(p1)), true, globalState)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_648_i6)) {
              break L7;
            }
            _648_i6 = (_648_i6).plus(_dafny.ONE);
            _531_v2 = !(_531_v2);
            let _649_v82;
            _649_v82 = _dafny.Map.Empty.slice().updateUnsafe(_531_v2,_534_v5);
            _531_v2 = (_module.__default.safeDivisionInt((_this).f0, (_module.D5.create_DC13(new BigNumber(678), (((_649_v82).contains(_531_v2)) ? ((_649_v82).get(_531_v2)) : (_534_v5)), _535_v6, (_this).f0, false)).dtor_cf17)).isLessThanOrEqualTo((_this).f0);
            _531_v2 = _531_v2;
            let _650_v83;
            _650_v83 = _dafny.Seq.of(_531_v2, _531_v2, true, !(_531_v2), true);
            let _651_v84;
            _651_v84 = _dafny.Seq.UnicodeFromString("oiil");
            let _652_v85;
            _652_v85 = _dafny.Seq.of(new BigNumber((_651_v84).length), (_dafny.ZERO).minus(p1));
            let _653_v86;
            _653_v86 = _dafny.MultiSet.fromElements(p1, p1);
            let _source6 = _module.__default.fm23((_650_v83)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f0), new BigNumber((_650_v83).length))], (_dafny.ZERO).minus((new BigNumber((_534_v5).length)).multipliedBy(p1)), (_dafny.ZERO).minus((_652_v85)[_module.__default.safeIndex((_this).f0, new BigNumber((_652_v85).length))]), _dafny.Seq.Concat(_module.__default.fm19(_653_v86, (_dafny.ZERO).minus((_this).f0), new BigNumber((_653_v86).cardinality()), new BigNumber((_651_v84).length), globalState), _651_v84), globalState);
            if (_source6.is_DC12) {
              let _654___mcc_h10 = (_source6).cf15;
              let _655___mcc_h11 = (_source6).cf16;
              let _656_cf16 = _655___mcc_h11;
              let _657_cf15 = _654___mcc_h10;
              let _658_v88;
              let _nw113 = new _module.C2();
              _nw113.__ctor(new BigNumber(((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qvaepprp")).length), new BigNumber(913))).Intersect(function () {
                let _coll18 = new _dafny.Set();
                for (const _compr_18 of _dafny.IntegerRange(new BigNumber(908), new BigNumber(602))) {
                  let _659_v87 = _compr_18;
                  if (((new BigNumber(908)).isLessThanOrEqualTo(_659_v87)) && ((_659_v87).isLessThan(new BigNumber(602)))) {
                    _coll18.add((_659_v87).minus((_dafny.ZERO).minus(p1)));
                  }
                }
                return _coll18;
              }())).length));
              _658_v88 = _nw113;
              r0 = (_652_v85)[_module.__default.safeIndex((_this).f0, new BigNumber((_652_v85).length))];
              _531_v2 = _531_v2;
              _531_v2 = _531_v2;
            } else if (_source6.is_DC13) {
              let _660___mcc_h12 = (_source6).cf17;
              let _661___mcc_h13 = (_source6).cf18;
              let _662___mcc_h14 = (_source6).cf19;
              let _663___mcc_h15 = (_source6).cf20;
              let _664___mcc_h16 = (_source6).cf21;
              let _665_cf21 = _664___mcc_h16;
              let _666_cf20 = _663___mcc_h15;
              let _667_cf19 = _662___mcc_h14;
              let _668_cf18 = _661___mcc_h13;
              let _669_cf17 = _660___mcc_h12;
              r0 = (_this).f0;
              _666_cf20 = new BigNumber(-224);
              _665_cf21 = ((_this).f0).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), ((_670_v86) => function (_671_i7) {
                return new BigNumber((_670_v86).cardinality());
              })(_653_v86))).length));
              _665_cf21 = _531_v2;
            } else {
              let _672___mcc_h17 = (_source6).cf14;
              let _673_cf14 = _672___mcc_h17;
              r1 = new BigNumber(-301);
              let _674_v89;
              _674_v89 = _module.D3.create_DC8(_531_v2, p1);
              let _675_v90;
              _675_v90 = _dafny.Map.Empty.slice().updateUnsafe(_531_v2,(_674_v89).dtor_cf9);
              let _676_v91;
              _676_v91 = _dafny.MultiSet.fromElements(_675_v90, _675_v90, _675_v90);
              let _677_v92;
              _677_v92 = _dafny.MultiSet.fromElements(_534_v5);
              let _rhs100 = !(_dafny.MultiSet.fromElements(_534_v5, _534_v5)).equals(_677_v92);
              let _rhs101 = _531_v2;
              let _rhs102 = !((p1).isLessThanOrEqualTo(new BigNumber(252)));
              let _rhs103 = ((_676_v91).Intersect(_dafny.MultiSet.fromElements(_675_v90))).Difference(_676_v91);
              let _rhs104 = (_module.__default.safeModuloInt((_this).f0, (_this).f0)).multipliedBy((_this).f0);
              _531_v2 = _rhs100;
              _531_v2 = _rhs101;
              _531_v2 = _rhs102;
              _676_v91 = _rhs103;
              r2 = _rhs104;
              _531_v2 = _531_v2;
              let _678_v93;
              let _nw114 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _678_v93 = _nw114;
              let _index104 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_678_v93).length));
              (_678_v93)[_index104] = p1;
              let _index105 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_678_v93).length));
              (_678_v93)[_index105] = (_dafny.ZERO).minus((_this).f0);
            }
          }
        }
      }
      let _679_i8;
      _679_i8 = _dafny.ZERO;
      L8: {
        while ((_module.__default.fm2(true, p1, _531_v2, _531_v2, globalState)) || (_531_v2)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_679_i8)) {
              break L8;
            }
            _679_i8 = (_679_i8).plus(_dafny.ONE);
            let _680_v94;
            let _init14 = ((_681_v2) => function (_682_i9) {
              return ((_681_v2) ? (_module.D7.create_DC17(_681_v2, _681_v2, _681_v2)) : (_module.D7.create_DC17(_681_v2, _681_v2, _681_v2)));
            })(_531_v2);
            let _nw115 = Array((new BigNumber(28)).toNumber());
            for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw115.length); _i0_14++) {
              _nw115[_i0_14] = _init14(new BigNumber(_i0_14));
            }
            _680_v94 = _nw115;
            let _nw116 = Array((new BigNumber(13)).toNumber()).fill(_module.D7.Default());
            _680_v94 = _nw116;
            let _683_v95;
            let _nw117 = Array((new BigNumber(16)).toNumber()).fill([]);
            _683_v95 = _nw117;
            let _684_v96;
            let _nw118 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
            _684_v96 = _nw118;
            let _index106 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_683_v95).length));
            (_683_v95)[_index106] = _684_v96;
            let _685_v97;
            let _nw119 = new _module.C1();
            _nw119.__ctor(p1);
            _685_v97 = _nw119;
            let _686_v98;
            _686_v98 = _dafny.Map.Empty.slice().updateUnsafe(_685_v97,_684_v96);
            let _687_v99;
            _687_v99 = (((_686_v98).contains(_685_v97)) ? ((_686_v98).get(_685_v97)) : (_684_v96));
            let _index107 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_683_v95).length));
            (_683_v95)[_index107] = (_687_v99);
            r0 = (_dafny.ZERO).minus(new BigNumber(-892));
            let _688_v100;
            _688_v100 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_534_v5, _535_v6, (_this).f0, p1, globalState),_dafny.Seq.of(_531_v2));
            let _689_v101;
            _689_v101 = _dafny.Seq.of(_531_v2, _531_v2, _531_v2);
            _688_v100 = (_688_v100).update(p1, _689_v101);
          }
        }
      }
      _531_v2 = _531_v2;
      r0 = (_this).f0;
      let _690_v102;
      _690_v102 = _dafny.Seq.of(p1);
      r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_690_v102)[_module.__default.safeIndex(p1, new BigNumber((_690_v102).length))]));
      r2 = new BigNumber((_533_v4).length);
      return [r0, r1, r2];
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _691_v0;
      let _nw120 = new _module.C1();
      _nw120.__ctor(p0);
      _691_v0 = _nw120;
      let _692_v1;
      _692_v1 = _dafny.Map.Empty.slice().updateUnsafe(_691_v0,(_691_v0).f0);
      _692_v1 = (_692_v1).update(_691_v0, (_this).f0);
      let _693_v2;
      _693_v2 = true;
      let _694_v3;
      _694_v3 = _dafny.Seq.UnicodeFromString("ejn");
      let _695_v4;
      _695_v4 = _module.D2.create_DC4(_694_v3, (_691_v0).fm5(globalState));
      let _696_v5;
      _696_v5 = _module.D3.create_DC8(_693_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fxkw"), (_695_v4).dtor_cf4)).length));
      let _source7 = _696_v5;
      if (_source7.is_DC8) {
        let _697___mcc_h0 = (_source7).cf8;
        let _698___mcc_h1 = (_source7).cf9;
        let _699_cf9 = _698___mcc_h1;
        let _700_cf8 = _697___mcc_h0;
        r1 = new BigNumber(-144);
        let _701_v6;
        let _nw121 = new _module.C1();
        _nw121.__ctor(_module.__default.safeDivisionInt(new BigNumber(276), new BigNumber((_694_v3).length)));
        _701_v6 = _nw121;
        let _702_v7;
        _702_v7 = new _dafny.CodePoint('s'.codePointAt(0));
        let _703_v8;
        let _nw122 = new _module.C0();
        _nw122.__ctor(_702_v7);
        _703_v8 = _nw122;
        let _704_v9;
        _704_v9 = _dafny.MultiSet.fromElements(_693_v2);
        let _705_v10;
        _705_v10 = _dafny.Seq.of(false);
        let _706_v11;
        _706_v11 = _dafny.Seq.of(_704_v9, _dafny.MultiSet.FromArray(_705_v10));
        let _707_v12;
        let _nw123 = new _module.C2();
        _nw123.__ctor(_module.__default.safeModuloInt(_699_cf9, new BigNumber(((_706_v11)[_module.__default.safeIndex(_699_cf9, new BigNumber((_706_v11).length))]).cardinality())));
        _707_v12 = _nw123;
      } else if (_source7.is_DC9) {
        let _708___mcc_h2 = (_source7).cf10;
        let _709___mcc_h3 = (_source7).cf11;
        let _710___mcc_h4 = (_source7).cf12;
        let _711_cf12 = _710___mcc_h4;
        let _712_cf11 = _709___mcc_h3;
        let _713_cf10 = _708___mcc_h2;
        let _714_v13;
        _714_v13 = _dafny.Seq.of(_dafny.Set.fromElements((_691_v0).f0, (_691_v0).f0));
        let _715_v14;
        _715_v14 = _dafny.Set.fromElements((_691_v0).f0);
        let _716_v15;
        _716_v15 = _dafny.MultiSet.fromElements((_dafny.Set.fromElements(p0)).Union((_714_v13)[_module.__default.safeIndex((_691_v0).f0, new BigNumber((_714_v13).length))]), _715_v14, _715_v14, (_715_v14).Union(_715_v14));
        _716_v15 = _716_v15;
        let _717_v16;
        let _nw124 = Array((new BigNumber(7)).toNumber()).fill(false);
        _717_v16 = _nw124;
        _717_v16 = _717_v16;
        let _718_v17;
        _718_v17 = _module.D6.create_DC15(new BigNumber(934), _711_cf12, (_691_v0).f0);
        _718_v17 = _718_v17;
        _694_v3 = _694_v3;
      } else {
        let _719___mcc_h5 = (_source7).cf7;
        let _720_cf7 = _719___mcc_h5;
        r1 = ((true) ? (_module.__default.safeDivisionInt(p0, p0)) : ((_691_v0).f0));
        _693_v2 = _693_v2;
        let _721_v18;
        _721_v18 = new _dafny.CodePoint('d'.codePointAt(0));
        _721_v18 = _module.__default.fm17((_691_v0).f0, globalState);
        let _722_v19;
        let _nw125 = Array((new BigNumber(11)).toNumber());
        _nw125[(_dafny.ZERO).toNumber()] = (_693_v2) && (_693_v2);
        _nw125[(_dafny.ONE).toNumber()] = _693_v2;
        _nw125[(new BigNumber(2)).toNumber()] = _693_v2;
        _nw125[(new BigNumber(3)).toNumber()] = _693_v2;
        _nw125[(new BigNumber(4)).toNumber()] = _693_v2;
        _nw125[(new BigNumber(5)).toNumber()] = false;
        _nw125[(new BigNumber(6)).toNumber()] = true;
        _nw125[(new BigNumber(7)).toNumber()] = _693_v2;
        _nw125[(new BigNumber(8)).toNumber()] = _693_v2;
        _nw125[(new BigNumber(9)).toNumber()] = _693_v2;
        _nw125[(new BigNumber(10)).toNumber()] = _693_v2;
        _722_v19 = _nw125;
        let _index108 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_722_v19).length));
        (_722_v19)[_index108] = (_this).fm5(globalState);
        let _index109 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_722_v19).length));
        (_722_v19)[_index109] = _693_v2;
        let _723_v20;
        _723_v20 = _dafny.Map.Empty.slice().updateUnsafe(_693_v2,_693_v2);
        let _724_v21;
        _724_v21 = _dafny.Seq.of(_723_v20, _723_v20, _723_v20);
        let _725_v22;
        _725_v22 = _dafny.Set.fromElements(_693_v2, _693_v2, (_691_v0).fm5(globalState));
        let _726_v23;
        _726_v23 = _module.D2.create_DC5();
        let _727_v24;
        _727_v24 = _module.D2.create_DC6(_726_v23);
        let _728_v25;
        _728_v25 = _dafny.MultiSet.fromElements(_727_v24);
        let _index110 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_722_v19).length));
        let _index111 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_722_v19).length));
        let _rhs105 = (_this).f0;
        let _rhs106 = _module.__default.fm17(_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p0)), globalState);
        let _rhs107 = ((((_693_v2) ? (_693_v2) : (_693_v2))) ? (_693_v2) : (_693_v2));
        let _rhs108 = ((_728_v25).update(_727_v24, _module.__default.abs(new BigNumber((_694_v3).length)))).IsProperSubsetOf(_module.__default.fm24(_693_v2, _693_v2, new BigNumber(((_724_v21)[_module.__default.safeIndex((_691_v0).f0, new BigNumber((_724_v21).length))]).length), _725_v22, globalState));
        let _lhs38 = _722_v19;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_722_v19).length));
        let _lhs40 = _722_v19;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_722_v19).length));
        r1 = _rhs105;
        _721_v18 = _rhs106;
        _lhs38[_lhs39] = _rhs107;
        _lhs40[_lhs41] = _rhs108;
      }
      let _729_v26;
      let _nw126 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _729_v26 = _nw126;
      let _index112 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_729_v26).length));
      (_729_v26)[_index112] = p0;
      let _index113 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_729_v26).length));
      (_729_v26)[_index113] = (p0).plus((_this).f0);
      let _730_v27;
      _730_v27 = _dafny.Map.Empty.slice().updateUnsafe(_729_v26,_729_v26);
      _730_v27 = (_730_v27).update(_729_v26, _729_v26);
      let _731_v28;
      let _init15 = function (_732_i0) {
        return true;
      };
      let _nw127 = Array((new BigNumber(28)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw127.length); _i0_15++) {
        _nw127[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _731_v28 = _nw127;
      _731_v28 = _731_v28;
      let _733_v29;
      _733_v29 = _dafny.Set.fromElements(_693_v2, _693_v2);
      r0 = (_733_v29).IsSubsetOf(_733_v29);
      r0 = _693_v2;
      let _734_v30;
      _734_v30 = _dafny.Set.fromElements(_729_v26, _729_v26);
      r1 = new BigNumber((_734_v30).length);
      return [r0, r1];
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _735_v0;
      _735_v0 = new BigNumber(-360);
      _735_v0 = _module.__default.safeModuloInt(_735_v0, p3);
      let _736_v1;
      _736_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
      if (!(((true) ? (_736_v1) : (_dafny.Map.Empty.slice().updateUnsafe(p1,p2)))).equals(_736_v1)) {
        let _737_v2;
        _737_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        _735_v0 = (new BigNumber(((_737_v2).Merge(_737_v2)).length)).multipliedBy((_this).f0);
        let _738_v3;
        _738_v3 = false;
        let _739_v4;
        _739_v4 = _dafny.Seq.of(p2, p2, p2, p2, p2);
        _738_v3 = _dafny.areEqual(_739_v4, _dafny.Seq.Concat(_739_v4, _739_v4));
        let _740_v5;
        _740_v5 = _dafny.Set.fromElements(p3);
        _738_v3 = !((_738_v3) || ((_740_v5).IsSubsetOf(_740_v5)));
        let _741_v6;
        _741_v6 = new _dafny.CodePoint('p'.codePointAt(0));
        let _742_v7;
        _742_v7 = _dafny.Seq.of(_741_v6, _741_v6);
        let _743_v8;
        let _nw128 = new _module.C2();
        _nw128.__ctor(new BigNumber((_740_v5).length));
        _743_v8 = _nw128;
        let _744_v9;
        let _nw129 = Array((new BigNumber(20)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = _743_v8;
        _nw129[(_dafny.ONE).toNumber()] = _743_v8;
        _nw129[(new BigNumber(2)).toNumber()] = ((_738_v3) ? (_743_v8) : (_743_v8));
        _nw129[(new BigNumber(3)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(4)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(5)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(6)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(7)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(8)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(9)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(10)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(11)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(12)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(13)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(14)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(15)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(16)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(17)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(18)).toNumber()] = _743_v8;
        _nw129[(new BigNumber(19)).toNumber()] = _743_v8;
        _744_v9 = _nw129;
        let _index114 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_744_v9).length));
        (_744_v9)[_index114] = _743_v8;
        let _745_v10;
        _745_v10 = _737_v2;
        let _746_v11;
        _746_v11 = _dafny.Seq.of((_743_v8).f0, new BigNumber((((_745_v10)).update(_738_v3, p1)).length));
        let _index115 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_744_v9).length));
        let _rhs109 = (_746_v11)[_module.__default.safeIndex(_module.__default.safeModuloInt((_746_v11)[_module.__default.safeIndex((_743_v8).f0, new BigNumber((_746_v11).length))], (_dafny.ZERO).minus(_735_v0)), new BigNumber((_746_v11).length))];
        let _rhs110 = p0;
        let _rhs111 = _743_v8;
        let _rhs112 = p1;
        let _lhs42 = _744_v9;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_744_v9).length));
        _735_v0 = _rhs109;
        _742_v7 = _rhs110;
        _lhs42[_lhs43] = _rhs111;
        _738_v3 = _rhs112;
        let _747_v12;
        let _init16 = ((_748_v7, _749_v3) => function (_750_i0) {
          return _module.D2.create_DC6(_module.D2.create_DC4(_748_v7, _749_v3));
        })(_742_v7, _738_v3);
        let _nw130 = Array((new BigNumber(27)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw130.length); _i0_16++) {
          _nw130[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _747_v12 = _nw130;
        let _751_v13;
        _751_v13 = _module.D2.create_DC3(_dafny.Seq.of(_module.D0.create_DC1(_738_v3)));
        let _index116 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_747_v12).length));
        (_747_v12)[_index116] = _module.D2.create_DC6(_751_v13);
        let _752_v14;
        _752_v14 = _module.D2.create_DC6(_module.D2.create_DC6(_751_v13));
        let _index117 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_747_v12).length));
        (_747_v12)[_index117] = _752_v14;
      } else {
        let _753_v15;
        _753_v15 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f0);
        let _754_v16;
        _754_v16 = _dafny.Set.fromElements(p1);
        let _755_v17;
        _755_v17 = new _dafny.CodePoint('b'.codePointAt(0));
        let _756_v18;
        _756_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_754_v16, _755_v17, _735_v0, new BigNumber(215), globalState),p1);
        let _757_v19;
        _757_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_756_v18).contains(_735_v0)) ? ((_756_v18).get(_735_v0)) : (p1)),_753_v15);
        let _758_v20;
        _758_v20 = _dafny.MultiSet.fromElements((_753_v15).Merge((((_757_v19).contains(p1)) ? ((_757_v19).get(p1)) : (_753_v15))), _753_v15, _753_v15);
        let _759_v21;
        _759_v21 = true;
        let _rhs113 = _758_v20;
        let _rhs114 = p1;
        _758_v20 = _rhs113;
        _759_v21 = _rhs114;
        let _760_v22;
        _760_v22 = _module.D5.create_DC13(_735_v0, _754_v16, _755_v17, (_this).f0, _759_v21);
        _759_v21 = (_760_v22).dtor_cf21;
        let _761_v23;
        _761_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
        let _762_v24;
        _762_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,_761_v23);
        let _763_v25;
        let _nw131 = Array((new BigNumber(6)).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = _761_v23;
        _nw131[(_dafny.ONE).toNumber()] = (_761_v23).update(p1, _759_v21);
        _nw131[(new BigNumber(2)).toNumber()] = _761_v23;
        _nw131[(new BigNumber(3)).toNumber()] = _761_v23;
        _nw131[(new BigNumber(4)).toNumber()] = (((_762_v24).contains(p1)) ? ((_762_v24).get(p1)) : (_761_v23));
        _nw131[(new BigNumber(5)).toNumber()] = _761_v23;
        _763_v25 = _nw131;
        let _764_v26;
        _764_v26 = _763_v25;
        let _source8 = _764_v26;
        let _765___mcc_h0 = _source8;
        let _766_cf31 = _765___mcc_h0;
        let _767_v27;
        let _nw132 = Array((new BigNumber(2)).toNumber()).fill(false);
        _767_v27 = _nw132;
        let _index118 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length));
        (_767_v27)[_index118] = (true) || (true);
        let _index119 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length));
        let _rhs115 = p1;
        let _rhs116 = (_754_v16).IsSubsetOf((_dafny.Set.fromElements(!(p1), true, _759_v21, false, _759_v21)).Intersect(_754_v16));
        let _rhs117 = _759_v21;
        let _lhs44 = _767_v27;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length));
        _759_v21 = _rhs115;
        _lhs44[_lhs45] = _rhs116;
        _759_v21 = _rhs117;
        _759_v21 = _759_v21;
        let _768_v28;
        _768_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus((_this).f0));
        let _769_v29;
        _769_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_768_v28).length),_755_v17);
        let _770_v30;
        _770_v30 = _dafny.Map.Empty.slice().updateUnsafe(_735_v0,_767_v27);
        _769_v29 = (_769_v29).update(((_759_v21) ? (new BigNumber((_770_v30).length)) : (_735_v0)), new _dafny.CodePoint('b'.codePointAt(0)));
        let _771_v31;
        _771_v31 = _dafny.Seq.UnicodeFromString("fdbsmm");
        let _772_v32;
        _772_v32 = _module.D0.create_DC1((_767_v27)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length))]);
        let _773_v33;
        _773_v33 = _dafny.Seq.of(_772_v32);
        let _774_v34;
        _774_v34 = _module.D2.create_DC3(_773_v33);
        let _775_v35;
        _775_v35 = _dafny.Seq.of(!(_759_v21));
        let _776_v36;
        let _nw133 = Array((new BigNumber(16)).toNumber());
        _nw133[(_dafny.ZERO).toNumber()] = _767_v27;
        _nw133[(_dafny.ONE).toNumber()] = _767_v27;
        _nw133[(new BigNumber(2)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(3)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(4)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(5)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(6)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(7)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(8)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(9)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(10)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(11)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(12)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(13)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(14)).toNumber()] = _767_v27;
        _nw133[(new BigNumber(15)).toNumber()] = _767_v27;
        _776_v36 = _nw133;
        let _index120 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_776_v36).length));
        (_776_v36)[_index120] = _767_v27;
        let _777_v37;
        let _nw134 = new _module.C1();
        _nw134.__ctor(p3);
        _777_v37 = _nw134;
        let _778_v38;
        _778_v38 = _module.D10.create_DC21(_777_v37);
        let _779_v39;
        _779_v39 = _dafny.Map.Empty.slice().updateUnsafe(_759_v21,(_778_v38).dtor_cf33);
        let _780_v40;
        _780_v40 = _dafny.Seq.of(_779_v39, _779_v39);
        let _781_v41;
        _781_v41 = _dafny.MultiSet.fromElements(new BigNumber(((_780_v40)[_module.__default.safeIndex(p3, new BigNumber((_780_v40).length))]).length));
        let _index121 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length));
        let _index122 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_776_v36).length));
        let _rhs118 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("evawlaen"), _module.__default.safeIndex((_this).f0, new BigNumber((_dafny.Seq.UnicodeFromString("evawlaen")).length)), _755_v17);
        let _rhs119 = !((_767_v27)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length))]);
        let _rhs120 = _774_v34;
        let _rhs121 = ((!(new BigNumber((_781_v41).cardinality())).isEqualTo((_this).f0)) ? (_dafny.Seq.of(p1)) : (_775_v35));
        let _rhs122 = _767_v27;
        let _lhs46 = _767_v27;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_767_v27).length));
        let _lhs48 = _776_v36;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_776_v36).length));
        _771_v31 = _rhs118;
        _lhs46[_lhs47] = _rhs119;
        _774_v34 = _rhs120;
        _775_v35 = _rhs121;
        _lhs48[_lhs49] = _rhs122;
        let _782_v42;
        let _nw135 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _782_v42 = _nw135;
        let _index123 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_782_v42).length));
        (_782_v42)[_index123] = _dafny.Seq.UnicodeFromString("vcrj");
        let _index124 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_782_v42).length));
        (_782_v42)[_index124] = _dafny.Seq.UnicodeFromString("jrrxlikqd");
        let _783_v43;
        _783_v43 = _dafny.Seq.of(_735_v0);
        let _784_v44;
        _784_v44 = _dafny.Seq.of(_759_v21);
        _759_v21 = (!_dafny.areEqual(_783_v43, _dafny.Seq.update(_783_v43, _module.__default.safeIndex((_this).f0, new BigNumber((_783_v43).length)), new BigNumber((_784_v44).length)))) && (p1);
      }
      let _785_v45;
      _785_v45 = _dafny.Set.fromElements(p1);
      _735_v0 = _module.__default.safeDivisionInt(p3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_785_v45).length))).length));
      let _786_v46;
      _786_v46 = new _dafny.CodePoint('g'.codePointAt(0));
      let _787_v47;
      _787_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
      let _788_v48;
      _788_v48 = _dafny.Map.Empty.slice().updateUnsafe(_786_v46,_787_v47);
      let _789_v49;
      let _nw136 = Array((new BigNumber(3)).toNumber());
      _nw136[(_dafny.ZERO).toNumber()] = (_788_v48).Merge(_788_v48);
      _nw136[(_dafny.ONE).toNumber()] = _788_v48;
      _nw136[(new BigNumber(2)).toNumber()] = _788_v48;
      _789_v49 = _nw136;
      let _index125 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_789_v49).length));
      (_789_v49)[_index125] = _788_v48;
      let _index126 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_789_v49).length));
      (_789_v49)[_index126] = _788_v48;
      let _hi5 = new BigNumber(-100);
      for (let _790_i1 = p3; _790_i1.isLessThan(_hi5); _790_i1 = _790_i1.plus(_dafny.ONE)) {
        let _791_v50;
        let _nw137 = Array((new BigNumber(4)).toNumber()).fill([]);
        _791_v50 = _nw137;
        _791_v50 = _791_v50;
        _735_v0 = _735_v0;
        let _792_v51;
        _792_v51 = _dafny.Set.fromElements(_dafny.Set.fromElements((_this).f0, (_dafny.ZERO).minus(p3)));
        let _793_v52;
        _793_v52 = _module.D5.create_DC11(_792_v51);
        let _source9 = _793_v52;
        if (_source9.is_DC12) {
          let _794___mcc_h1 = (_source9).cf15;
          let _795___mcc_h2 = (_source9).cf16;
          let _796_cf16 = _795___mcc_h2;
          let _797_cf15 = _794___mcc_h1;
          let _798_v53;
          _798_v53 = _module.D2.create_DC4(p0, false);
          let _799_v54;
          _799_v54 = _dafny.Map.Empty.slice().updateUnsafe((_798_v53).dtor_cf5,p1);
          _799_v54 = _799_v54;
          let _800_v55;
          _800_v55 = _dafny.MultiSet.fromElements(_790_i1, p3);
          let _801_v56;
          let _802_v57;
          let _out15;
          let _out16;
          let _outcollector1 = (_this).m6(new BigNumber(((_800_v55).update(_796_cf16, _module.__default.abs(_797_cf15))).cardinality()), globalState);
          _out15 = _outcollector1[0];
          _out16 = _outcollector1[1];
          _801_v56 = _out15;
          _802_v57 = _out16;
          let _803_v58;
          let _nw138 = new _module.C0();
          _nw138.__ctor(_786_v46);
          _803_v58 = _nw138;
          _801_v56 = !(false);
        } else if (_source9.is_DC13) {
          let _804___mcc_h3 = (_source9).cf17;
          let _805___mcc_h4 = (_source9).cf18;
          let _806___mcc_h5 = (_source9).cf19;
          let _807___mcc_h6 = (_source9).cf20;
          let _808___mcc_h7 = (_source9).cf21;
          let _809_cf21 = _808___mcc_h7;
          let _810_cf20 = _807___mcc_h6;
          let _811_cf19 = _806___mcc_h5;
          let _812_cf18 = _805___mcc_h4;
          let _813_cf17 = _804___mcc_h3;
          let _814_v59;
          _814_v59 = _dafny.Seq.of(p1, _809_cf21);
          let _815_v60;
          _815_v60 = _module.D7.create_DC16(_814_v59);
          _815_v60 = _815_v60;
          let _816_v61;
          _816_v61 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(p3)), p3, (_dafny.ZERO).minus(_810_cf20));
          let _817_v62;
          _817_v62 = _dafny.MultiSet.fromElements(new BigNumber((_816_v61).cardinality()), (_this).f0, _735_v0);
          let _index127 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((p2).length));
          (p2)[_index127] = _790_i1;
          let _index128 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((p2).length));
          let _rhs123 = ((p1) ? (_816_v61) : (_817_v62));
          let _rhs124 = new BigNumber(268);
          let _lhs50 = p2;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((p2).length));
          _817_v62 = _rhs123;
          _lhs50[_lhs51] = _rhs124;
          _735_v0 = _module.__default.safeDivisionInt(_790_i1, _810_cf20);
          let _818_v63;
          let _nw139 = Array((new BigNumber(25)).toNumber()).fill(false);
          _818_v63 = _nw139;
          let _index129 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_818_v63).length));
          (_818_v63)[_index129] = p1;
          let _index130 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_818_v63).length));
          (_818_v63)[_index130] = !(_809_cf21);
        } else {
          let _819___mcc_h8 = (_source9).cf14;
          let _820_cf14 = _819___mcc_h8;
          _735_v0 = _790_i1;
          let _821_v64;
          let _822_v65;
          let _out17;
          let _out18;
          let _outcollector2 = (_this).m6(_735_v0, globalState);
          _out17 = _outcollector2[0];
          _out18 = _outcollector2[1];
          _821_v64 = _out17;
          _822_v65 = _out18;
          _822_v65 = _790_i1;
          let _823_v66;
          _823_v66 = _dafny.Map.Empty.slice().updateUnsafe(_821_v64,_821_v64);
          let _824_v67;
          _824_v67 = _823_v66;
          _824_v67 = _module.__default.fm25(globalState);
        }
        let _825_v68;
        let _nw140 = new _module.C1();
        _nw140.__ctor((_this).f0);
        _825_v68 = _nw140;
        let _826_v69;
        _826_v69 = _module.D10.create_DC21(_825_v68);
        let _827_v70;
        _827_v70 = _module.D0.create_DC1(p1);
        let _828_v71;
        _828_v71 = _dafny.Seq.of(_827_v70);
        let _829_v72;
        _829_v72 = _dafny.Map.Empty.slice().updateUnsafe(_826_v69,_module.D2.create_DC3(_828_v71));
        let _830_v73;
        _830_v73 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_826_v69,_module.D2.create_DC3(_dafny.Seq.of(_827_v70)))).Merge(_829_v72), (_829_v72).Merge(_829_v72), ((p1) ? (_829_v72) : (_829_v72)), (_829_v72).Merge(_829_v72), _829_v72);
        _830_v73 = _830_v73;
      }
      let _831_v74;
      _831_v74 = _dafny.MultiSet.fromElements(p1, p1, p1, true, p1);
      let _832_v75;
      _832_v75 = _dafny.Seq.of((((_831_v74).contains(p1)) ? ((_831_v74).get(p1)) : (p3)), p3);
      let _833_v76;
      _833_v76 = _module.D0.create_DC1(p1);
      let _834_v77;
      _834_v77 = _dafny.Seq.of(_833_v76);
      let _835_v78;
      _835_v78 = _module.D2.create_DC3(_834_v77);
      let _836_v79;
      _836_v79 = _dafny.Map.Empty.slice().updateUnsafe(p1,_835_v78);
      let _837_v80;
      _837_v80 = _module.D2.create_DC6((((_836_v79).contains(p1)) ? ((_836_v79).get(p1)) : (_835_v78)));
      let _pat_let_tv15 = _832_v75;
      let _pat_let_tv16 = _735_v0;
      let _pat_let_tv17 = _832_v75;
      let _pat_let_tv18 = _832_v75;
      let _pat_let_tv19 = _832_v75;
      let _pat_let_tv20 = _832_v75;
      let _pat_let_tv21 = _735_v0;
      let _pat_let_tv22 = _835_v78;
      _832_v75 = function (_source10) {
        if (_source10.is_DC4) {
          let _838___mcc_h9 = (_source10).cf4;
          let _839___mcc_h10 = (_source10).cf5;
          let _840_cf5 = _839___mcc_h10;
          let _841_cf4 = _838___mcc_h9;
          return _pat_let_tv15;
        } else if (_source10.is_DC5) {
          return _dafny.Seq.Concat(_dafny.Seq.of((_this).f0, (_this).f0, _pat_let_tv16), _pat_let_tv17);
        } else if (_source10.is_DC3) {
          let _842___mcc_h11 = (_source10).cf3;
          let _843_cf3 = _842___mcc_h11;
          return _pat_let_tv18;
        } else {
          let _844___mcc_h12 = (_source10).cf6;
          let _845_cf6 = _844___mcc_h12;
          return _dafny.Seq.update(_pat_let_tv19, _module.__default.safeIndex((_this).f0, new BigNumber((_pat_let_tv20).length)), _pat_let_tv21);
        }
      }(function (_pat_let13_0) {
        return function (_846_dt__update__tmp_h0) {
          return function (_pat_let14_0) {
            return function (_847_dt__update_hcf6_h0) {
              return _module.D2.create_DC6(_847_dt__update_hcf6_h0);
            }(_pat_let14_0);
          }(_pat_let_tv22);
        }(_pat_let13_0);
      }(_837_v80));
      return;
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f0 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return !(!(!(!(_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vrjq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_848_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })), new _dafny.CodePoint('i'.codePointAt(0)))))));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(true) || ((_module.D3.create_DC8(true, new BigNumber(-634))).dtor_cf8);
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _849_v0;
      _849_v0 = false;
      let _850_v1;
      _850_v1 = _dafny.MultiSet.fromElements(_849_v0);
      r2 = _module.__default.safeDivisionInt(new BigNumber((_850_v1).cardinality()), (p1).multipliedBy(p1));
      if (_849_v0) {
        let _851_v3;
        _851_v3 = _dafny.Set.fromElements(p1, p1);
        let _852_v4;
        _852_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _853_v6;
        _853_v6 = _dafny.Seq.of((_this).f0);
        let _854_v7;
        _854_v7 = _module.D3.create_DC7(_853_v6);
        let _855_v8;
        let _nw141 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _855_v8 = _nw141;
        let _856_v9;
        _856_v9 = new _dafny.CodePoint('c'.codePointAt(0));
        let _857_v10;
        let _nw142 = new _module.C0();
        _nw142.__ctor(_856_v9);
        _857_v10 = _nw142;
        let _858_v11;
        _858_v11 = _dafny.Seq.UnicodeFromString("wgt");
        let _859_v15;
        let _nw143 = Array((new BigNumber(29)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_851_v3).Elements) {
            let _860_v2 = _compr_19;
            if ((_851_v3).contains(_860_v2)) {
              _coll19.push([(_860_v2).multipliedBy(new BigNumber((_850_v1).cardinality())),(_this).f0]);
            }
          }
          return _coll19;
        }();
        _nw143[(_dafny.ONE).toNumber()] = (_852_v4).Merge(function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of _dafny.IntegerRange(new BigNumber(859), new BigNumber(772))) {
            let _861_v5 = _compr_20;
            if (((new BigNumber(859)).isLessThanOrEqualTo(_861_v5)) && ((_861_v5).isLessThan(new BigNumber(772)))) {
              _coll20.push([(_861_v5).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_849_v0,_module.D2.create_DC5())).length)),new BigNumber((_852_v4).length)]);
            }
          }
          return _coll20;
        }());
        _nw143[(new BigNumber(2)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_854_v7).dtor_cf7).length),(_this).f0);
        _nw143[(new BigNumber(4)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(5)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(6)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(7)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(8)).toNumber()] = (_852_v4).Merge(_852_v4);
        _nw143[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,new BigNumber((_dafny.Seq.of(_855_v8, _855_v8, _855_v8)).length));
        _nw143[(new BigNumber(10)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(11)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(12)).toNumber()] = _module.__default.fm11((_this).f0, globalState);
        _nw143[(new BigNumber(13)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_857_v10), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_851_v3).length)), new BigNumber((_dafny.Seq.of(_857_v10)).length)), _857_v10)).length),new BigNumber((_858_v11).length));
        _nw143[(new BigNumber(15)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(16)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(17)).toNumber()] = (_852_v4).Merge(function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(896), new BigNumber(834))) {
            let _862_v12 = _compr_21;
            if (((new BigNumber(896)).isLessThanOrEqualTo(_862_v12)) && ((_862_v12).isLessThan(new BigNumber(834)))) {
              _coll21.push([(_862_v12).multipliedBy(new BigNumber(-918)),p1]);
            }
          }
          return _coll21;
        }());
        _nw143[(new BigNumber(18)).toNumber()] = (_852_v4).Merge(_852_v4);
        _nw143[(new BigNumber(19)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(20)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(21)).toNumber()] = function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(156), new BigNumber(62))) {
            let _863_v13 = _compr_22;
            if (((new BigNumber(156)).isLessThanOrEqualTo(_863_v13)) && ((_863_v13).isLessThan(new BigNumber(62)))) {
              _coll22.push([_module.__default.safeDivisionInt(_863_v13, p1),(_this).f0]);
            }
          }
          return _coll22;
        }();
        _nw143[(new BigNumber(22)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(23)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(24)).toNumber()] = _module.__default.fm11(p1, globalState);
        _nw143[(new BigNumber(25)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(26)).toNumber()] = function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(659), new BigNumber(736))) {
            let _864_v14 = _compr_23;
            if (((new BigNumber(659)).isLessThanOrEqualTo(_864_v14)) && ((_864_v14).isLessThan(new BigNumber(736)))) {
              _coll23.push([(_864_v14).minus(p1),new BigNumber(225)]);
            }
          }
          return _coll23;
        }();
        _nw143[(new BigNumber(27)).toNumber()] = _852_v4;
        _nw143[(new BigNumber(28)).toNumber()] = _852_v4;
        _859_v15 = _nw143;
        let _index131 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_859_v15).length));
        (_859_v15)[_index131] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f0);
        let _865_v16;
        let _init17 = ((_866_v0) => function (_867_i0) {
          return _dafny.Seq.of(_866_v0, _866_v0);
        })(_849_v0);
        let _nw144 = Array((new BigNumber(12)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw144.length); _i0_17++) {
          _nw144[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _865_v16 = _nw144;
        let _index132 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length));
        (_865_v16)[_index132] = _module.__default.fm12((_this).f0, (_this).f0, _849_v0, globalState);
        let _868_v17;
        let _init18 = ((_869_v0) => function (_870_i1) {
          return !(_869_v0) || (_869_v0);
        })(_849_v0);
        let _nw145 = Array((_dafny.ONE).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw145.length); _i0_18++) {
          _nw145[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _868_v17 = _nw145;
        let _871_v18;
        _871_v18 = _dafny.Seq.of(_849_v0, _849_v0);
        let _872_v19;
        _872_v19 = _868_v17;
        let _index133 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_859_v15).length));
        let _index134 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length));
        let _rhs125 = _852_v4;
        let _rhs126 = _858_v11;
        let _rhs127 = _871_v18;
        let _rhs128 = (_868_v17);
        let _rhs129 = _dafny.Seq.update(_dafny.Seq.Concat(_858_v11, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), ((_873_v10) => function (_874_i2) {
          return (_873_v10).f3;
        })(_857_v10)), _858_v11)), _module.__default.safeIndex(new BigNumber((_871_v18).length), new BigNumber((_dafny.Seq.Concat(_858_v11, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), ((_875_v10) => function (_876_i2) {
          return (_875_v10).f3;
        })(_857_v10)), _858_v11))).length)), (_857_v10).f3);
        let _lhs52 = _859_v15;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_859_v15).length));
        let _lhs54 = _865_v16;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length));
        _lhs52[_lhs53] = _rhs125;
        _858_v11 = _rhs126;
        _lhs54[_lhs55] = _rhs127;
        _868_v17 = _rhs128;
        _858_v11 = _rhs129;
        let _877_v20;
        let _nw146 = new _module.C0();
        _nw146.__ctor(new _dafny.CodePoint('c'.codePointAt(0)));
        _877_v20 = _nw146;
        let _pat_let_tv23 = _853_v6;
        let _pat_let_tv24 = _853_v6;
        let _source11 = function (_pat_let15_0) {
          return function (_878_dt__update__tmp_h1) {
            return function (_pat_let16_0) {
              return function (_879_dt__update_hcf7_h0) {
                return _module.D3.create_DC7(_879_dt__update_hcf7_h0);
              }(_pat_let16_0);
            }(_dafny.Seq.update(_pat_let_tv23, _module.__default.safeIndex(new BigNumber(615), new BigNumber((_pat_let_tv24).length)), (_this).f0));
          }(_pat_let15_0);
        }(_854_v7);
        if (_source11.is_DC8) {
          let _880___mcc_h0 = (_source11).cf8;
          let _881___mcc_h1 = (_source11).cf9;
          let _882_cf9 = _881___mcc_h1;
          let _883_cf8 = _880___mcc_h0;
          let _884_v21;
          let _nw147 = new _module.C3();
          _nw147.__ctor((_this).f0);
          _884_v21 = _nw147;
          let _885_v22;
          let _nw148 = Array((new BigNumber(25)).toNumber());
          _nw148[(_dafny.ZERO).toNumber()] = _884_v21;
          _nw148[(_dafny.ONE).toNumber()] = _884_v21;
          _nw148[(new BigNumber(2)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(3)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(4)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(5)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(6)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(7)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(8)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(9)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(10)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(11)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(12)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(13)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(14)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(15)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(16)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(17)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(18)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(19)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(20)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(21)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(22)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(23)).toNumber()] = _884_v21;
          _nw148[(new BigNumber(24)).toNumber()] = _884_v21;
          _885_v22 = _nw148;
          let _index135 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_885_v22).length));
          (_885_v22)[_index135] = _884_v21;
          let _886_v23;
          _886_v23 = _module.D5.create_DC13(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_858_v11).length),_849_v0)).length), _dafny.Set.fromElements(_883_cf8), (_858_v11)[_module.__default.safeIndex((_884_v21).f0, new BigNumber((_858_v11).length))], _882_cf9, _849_v0);
          let _index136 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_885_v22).length));
          (_885_v22)[_index136] = ((_883_cf8) ? ((((_886_v23).dtor_cf21) ? (_884_v21) : (_884_v21))) : (_884_v21));
          r1 = _882_cf9;
          let _887_v24;
          _887_v24 = _dafny.Set.fromElements(_849_v0, _883_cf8, _883_cf8, _883_cf8);
          r1 = (((_this).f0).multipliedBy(_module.__default.fm3(_887_v24, _856_v9, _882_cf9, new BigNumber((_858_v11).length), globalState))).multipliedBy((_884_v21).f0);
          let _888_v25;
          let _nw149 = Array((new BigNumber(6)).toNumber()).fill([]);
          _888_v25 = _nw149;
          let _index137 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_888_v25).length));
          (_888_v25)[_index137] = _868_v17;
          let _index138 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_888_v25).length));
          (_888_v25)[_index138] = _868_v17;
        } else if (_source11.is_DC9) {
          let _889___mcc_h2 = (_source11).cf10;
          let _890___mcc_h3 = (_source11).cf11;
          let _891___mcc_h4 = (_source11).cf12;
          let _892_cf12 = _891___mcc_h4;
          let _893_cf11 = _890___mcc_h3;
          let _894_cf10 = _889___mcc_h2;
          _849_v0 = _849_v0;
          let _895_v26;
          let _nw150 = new _module.C1();
          _nw150.__ctor(_module.__default.safeModuloInt(_894_cf10, new BigNumber(((_865_v16)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length))]).length)));
          _895_v26 = _nw150;
          let _896_v27;
          let _nw151 = new _module.C1();
          _nw151.__ctor((_this).f0);
          _896_v27 = _nw151;
          let _897_v28;
          _897_v28 = _dafny.Map.Empty.slice().updateUnsafe(_849_v0,p1);
          let _898_v29;
          _898_v29 = _dafny.Seq.of(_897_v28, _897_v28, _897_v28);
          let _899_v30;
          _899_v30 = _module.D3.create_DC8(_849_v0, p1);
          let _900_v31;
          _900_v31 = _dafny.Map.Empty.slice().updateUnsafe((_898_v29)[_module.__default.safeIndex(p1, new BigNumber((_898_v29).length))],(_this).fm6((_899_v30).dtor_cf8, _858_v11, _858_v11, new BigNumber(432), globalState));
          _900_v31 = (_900_v31).update(_897_v28, _892_cf12);
        } else {
          let _901___mcc_h5 = (_source11).cf7;
          let _902_cf7 = _901___mcc_h5;
          let _903_v32;
          _903_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,_868_v17);
          let _904_v33;
          _904_v33 = _dafny.Set.fromElements(_849_v0);
          let _905_v34;
          _905_v34 = _dafny.Map.Empty.slice().updateUnsafe(_849_v0,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_904_v33).length),_868_v17));
          let _906_v35;
          _906_v35 = _dafny.Map.Empty.slice().updateUnsafe(false,_849_v0);
          let _rhs130 = (((_905_v34).contains((((_906_v35).contains((_871_v18)[_module.__default.safeIndex((_this).f0, new BigNumber((_871_v18).length))])) ? ((_906_v35).get((_871_v18)[_module.__default.safeIndex((_this).f0, new BigNumber((_871_v18).length))])) : (_module.__default.fm2(_849_v0, (_this).f0, _849_v0, _849_v0, globalState))))) ? ((_905_v34).get((((_906_v35).contains((_871_v18)[_module.__default.safeIndex((_this).f0, new BigNumber((_871_v18).length))])) ? ((_906_v35).get((_871_v18)[_module.__default.safeIndex((_this).f0, new BigNumber((_871_v18).length))])) : (_module.__default.fm2(_849_v0, (_this).f0, _849_v0, _849_v0, globalState))))) : ((_dafny.Map.Empty.slice().updateUnsafe(p1,_868_v17)).update(p1, _868_v17)));
          let _rhs131 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f0, p1)));
          _903_v32 = _rhs130;
          r2 = _rhs131;
          let _907_v36;
          _907_v36 = _dafny.MultiSet.fromElements((_this).f0, _module.__default.fm3(_904_v33, (_857_v10).f3, p1, (_dafny.ZERO).minus((_this).f0), globalState));
          let _908_v37;
          let _nw152 = new _module.C2();
          _nw152.__ctor(p1);
          _908_v37 = _nw152;
          let _909_v38;
          _909_v38 = _dafny.Set.fromElements(_908_v37);
          let _910_v39;
          let _nw153 = new _module.C2();
          _nw153.__ctor(_module.__default.fm3(_904_v33, (_857_v10).f3, p1, (((_907_v36).contains(new BigNumber((_909_v38).length))) ? ((_907_v36).get(new BigNumber((_909_v38).length))) : (p1)), globalState));
          _910_v39 = _nw153;
          _910_v39 = _910_v39;
          let _911_v40;
          _911_v40 = _dafny.Map.Empty.slice().updateUnsafe(_855_v8,(_dafny.Map.Empty.slice().updateUnsafe(_849_v0,_849_v0)).update(_849_v0, _849_v0));
          let _912_v41;
          _912_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(467),_906_v35);
          let _913_v42;
          _913_v42 = _dafny.Map.Empty.slice().updateUnsafe((_857_v10).f3,(_908_v37).f0);
          let _914_v43;
          _914_v43 = _dafny.Map.Empty.slice().updateUnsafe(!(_849_v0),(((_913_v42).contains((_857_v10).f3)) ? ((_913_v42).get((_857_v10).f3)) : (new BigNumber((_858_v11).length))));
          let _915_v44;
          _915_v44 = _dafny.Map.Empty.slice().updateUnsafe(p1,_849_v0);
          _906_v35 = ((((_911_v40).contains(_855_v8)) ? ((_911_v40).get(_855_v8)) : ((((_912_v41).contains(new BigNumber((_914_v43).length))) ? ((_912_v41).get(new BigNumber((_914_v43).length))) : (_dafny.Map.Empty.slice().updateUnsafe(_849_v0,(((_915_v44).contains(p1)) ? ((_915_v44).get(p1)) : (_849_v0)))))))).update(_849_v0, _849_v0);
          let _index139 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_855_v8).length));
          (_855_v8)[_index139] = (_dafny.ZERO).minus((_908_v37).f0);
          let _916_v45;
          _916_v45 = _dafny.Map.Empty.slice().updateUnsafe(!(_849_v0),_dafny.Set.fromElements(_849_v0));
          let _index140 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_855_v8).length));
          (_855_v8)[_index140] = _module.__default.fm3(((((_916_v45).contains(_849_v0)) ? ((_916_v45).get(_849_v0)) : (_904_v33))).Difference(_904_v33), new _dafny.CodePoint('i'.codePointAt(0)), (_908_v37).f0, (_this).f0, globalState);
        }
        let _917_v46;
        _917_v46 = _dafny.Set.fromElements(!(_849_v0), _849_v0, _849_v0);
        let _918_v47;
        _918_v47 = _dafny.MultiSet.fromElements(new BigNumber(721));
        let _919_v48;
        _919_v48 = _module.D0.create_DC1(((_865_v16)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length))])[_module.__default.safeIndex(p1, new BigNumber(((_865_v16)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length))]).length))]);
        let _920_v49;
        _920_v49 = _module.D2.create_DC3(_dafny.Seq.of(_919_v48));
        let _921_v50;
        _921_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,_920_v49);
        let _922_v51;
        _922_v51 = _dafny.MultiSet.fromElements(_module.D2.create_DC6((((_921_v50).contains(p1)) ? ((_921_v50).get(p1)) : (_920_v49))), _module.D2.create_DC6(_920_v49));
        let _923_v52;
        _923_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(_module.__default.fm3(_917_v46, _856_v9, new BigNumber((_918_v47).cardinality()), (_this).f0, globalState), globalState),_922_v51);
        let _924_v53;
        let _nw154 = new _module.C3();
        _nw154.__ctor(new BigNumber((_850_v1).cardinality()));
        _924_v53 = _nw154;
        let _925_v54;
        _925_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,_924_v53);
        let _926_v55;
        _926_v55 = _dafny.Map.Empty.slice().updateUnsafe((((_923_v52).contains(_858_v11)) ? ((_923_v52).get(_858_v11)) : (_922_v51)),(((_925_v54).contains(new BigNumber(693))) ? ((_925_v54).get(new BigNumber(693))) : (_924_v53)));
        _926_v55 = (_926_v55).update(_922_v51, _924_v53);
        if ((p1).isLessThan(_module.__default.safeModuloInt(new BigNumber((_858_v11).length), (_this).f0))) {
          let _index141 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_868_v17).length));
          (_868_v17)[_index141] = _849_v0;
          let _927_v56;
          _927_v56 = _module.D2.create_DC4(_858_v11, _849_v0);
          let _index142 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_868_v17).length));
          let _rhs132 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("k"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(95)), ((_928_v20) => function (_929_i3) {
            return (_928_v20).f3;
          })(_877_v20)));
          let _rhs133 = !(_dafny.Seq.IsProperPrefixOf(_858_v11, (_927_v56).dtor_cf4));
          let _lhs56 = _868_v17;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_868_v17).length));
          _858_v11 = _rhs132;
          _lhs56[_lhs57] = _rhs133;
          (_924_v53).m7(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jljbktggr"), _858_v11), _849_v0, _855_v8, new BigNumber((((_849_v0) ? (_853_v6) : (_853_v6))).length), globalState);
          let _index143 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_868_v17).length));
          (_868_v17)[_index143] = (p1).isLessThan(p1);
          let _930_v57;
          let _nw155 = new _module.C1();
          _nw155.__ctor((_this).f0);
          _930_v57 = _nw155;
          let _931_v58;
          _931_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,_849_v0);
          let _932_v59;
          _932_v59 = _dafny.Map.Empty.slice().updateUnsafe(_930_v57,_931_v58);
          let _933_v60;
          _933_v60 = _dafny.Map.Empty.slice().updateUnsafe(_849_v0,_932_v59);
          _933_v60 = (_933_v60).update((_930_v57).fm6(_849_v0, _858_v11, _858_v11, p1, globalState), ((_849_v0) ? (_932_v59) : ((_932_v59).update(_930_v57, _931_v58))));
          let _index144 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length));
          (_865_v16)[_index144] = (_865_v16)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_865_v16).length))];
        } else {
          _856_v9 = (_857_v10).f3;
          let _934_v61;
          let _nw156 = new _module.C3();
          _nw156.__ctor(p1);
          _934_v61 = _nw156;
          let _935_v62;
          let _nw157 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
          _935_v62 = _nw157;
          let _index145 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_935_v62).length));
          (_935_v62)[_index145] = _917_v46;
          let _936_v63;
          let _nw158 = Array((new BigNumber(27)).toNumber()).fill(_dafny.MultiSet.Empty);
          _936_v63 = _nw158;
          let _index146 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_936_v63).length));
          (_936_v63)[_index146] = _918_v47;
          let _index147 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_935_v62).length));
          let _index148 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_936_v63).length));
          let _rhs134 = ((_module.__default.fm2(_849_v0, (_this).f0, !(_849_v0), false, globalState)) ? (p1) : (_module.__default.fm3(_917_v46, (_877_v20).f3, new BigNumber((_module.__default.fm11(new BigNumber((_858_v11).length), globalState)).length), new BigNumber((_851_v3).length), globalState)));
          let _rhs135 = (_dafny.ZERO).minus((_this).f0);
          let _rhs136 = _855_v8;
          let _rhs137 = _917_v46;
          let _rhs138 = _dafny.MultiSet.fromElements((_this).f0);
          let _lhs58 = _935_v62;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_935_v62).length));
          let _lhs60 = _936_v63;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_936_v63).length));
          r2 = _rhs134;
          r2 = _rhs135;
          _855_v8 = _rhs136;
          _lhs58[_lhs59] = _rhs137;
          _lhs60[_lhs61] = _rhs138;
          _856_v9 = ((_849_v0) ? (_856_v9) : ((_877_v20).f3));
          let _index149 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_868_v17).length));
          (_868_v17)[_index149] = true;
          let _index150 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_868_v17).length));
          (_868_v17)[_index150] = !(_module.__default.fm2(((_935_v62)[_module.__default.safeIndex(new BigNumber(898), new BigNumber((_935_v62).length))]).IsProperSubsetOf(_dafny.Set.fromElements(_849_v0)), p1, (true) || (_849_v0), ((_849_v0) ? (_849_v0) : (_849_v0)), globalState));
        }
      } else {
        let _937_v64;
        let _nw159 = new _module.C3();
        _nw159.__ctor(p1);
        _937_v64 = _nw159;
        if (_849_v0) {
          let _938_v65;
          let _nw160 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _938_v65 = _nw160;
          let _index151 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length));
          (_938_v65)[_index151] = p1;
          let _index152 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length));
          (_938_v65)[_index152] = p1;
          let _939_v66;
          let _init19 = ((_940_p1) => function (_941_i4) {
            return (new BigNumber(854)).isLessThanOrEqualTo(_940_p1);
          })(p1);
          let _nw161 = Array((_dafny.ONE).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw161.length); _i0_19++) {
            _nw161[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _939_v66 = _nw161;
          let _index153 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_939_v66).length));
          (_939_v66)[_index153] = _849_v0;
          let _942_v67;
          _942_v67 = _dafny.Seq.of(!(_849_v0));
          let _943_v68;
          _943_v68 = _module.D7.create_DC16(_942_v67);
          let _944_v69;
          _944_v69 = new _dafny.CodePoint('k'.codePointAt(0));
          let _945_v70;
          _945_v70 = _944_v69;
          let _946_v71;
          _946_v71 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))]),_945_v70);
          let _947_v72;
          _947_v72 = _dafny.Seq.of(_938_v65);
          let _948_v73;
          _948_v73 = _dafny.Seq.of((_this).f0);
          let _949_v74;
          _949_v74 = _dafny.Set.fromElements(_849_v0);
          let _950_v75;
          _950_v75 = _dafny.MultiSet.fromElements(new BigNumber((_948_v73).length), _module.__default.fm3(_949_v74, new _dafny.CodePoint('q'.codePointAt(0)), new BigNumber(-458), new BigNumber(-6), globalState));
          let _951_v76;
          _951_v76 = _dafny.Seq.of(p1, new BigNumber(((_module.D11.create_DC23(_947_v72)).dtor_cf37).length), new BigNumber((_dafny.Seq.UnicodeFromString("xurqdyud")).length), (_this).f0, (((_950_v75).contains((_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))])) ? ((_950_v75).get((_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))])) : (p1)));
          let _952_v77;
          _952_v77 = _dafny.Seq.of(new BigNumber((_946_v71).length), (new BigNumber((_951_v76).length)).plus(new BigNumber((_950_v75).cardinality())));
          let _953_v78;
          _953_v78 = _dafny.Seq.of(_952_v77, _951_v76, _dafny.Seq.of((_this).f0), _dafny.Seq.of((_this).f0));
          let _index154 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_939_v66).length));
          let _rhs139 = _849_v0;
          let _rhs140 = _943_v68;
          let _rhs141 = (_953_v78)[_module.__default.safeIndex(p1, new BigNumber((_953_v78).length))];
          let _rhs142 = !(_849_v0) || (_849_v0);
          let _lhs62 = _939_v66;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_939_v66).length));
          _lhs62[_lhs63] = _rhs139;
          _943_v68 = _rhs140;
          _952_v77 = _rhs141;
          _849_v0 = _rhs142;
          let _954_v79;
          _954_v79 = _module.D7.create_DC18((_dafny.ZERO).minus((_this).f0));
          let _955_v80;
          let _nw162 = new _module.C1();
          _nw162.__ctor((_this).f0);
          _955_v80 = _nw162;
          let _pat_let_tv25 = _955_v80;
          let _pat_let_tv26 = _849_v0;
          let _956_v81;
          let _nw163 = Array((new BigNumber(29)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = _module.D7.create_DC18((_dafny.ZERO).minus((((_850_v1).contains(_849_v0)) ? ((_850_v1).get(_849_v0)) : ((_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))]))));
          _nw163[(_dafny.ONE).toNumber()] = _954_v79;
          _nw163[(new BigNumber(2)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(3)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(4)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(5)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(6)).toNumber()] = _module.D7.create_DC18((_this).f0);
          _nw163[(new BigNumber(7)).toNumber()] = function (_pat_let17_0) {
            return function (_959_dt__update__tmp_h3) {
              return function (_pat_let20_0) {
                return function (_960_dt__update_hcf30_h1) {
                  return _module.D7.create_DC18(_960_dt__update_hcf30_h1);
                }(_pat_let20_0);
              }((_this).f0);
            }(_pat_let17_0);
          }(function (_pat_let18_0) {
            return function (_957_dt__update__tmp_h2) {
              return function (_pat_let19_0) {
                return function (_958_dt__update_hcf30_h0) {
                  return _module.D7.create_DC18(_958_dt__update_hcf30_h0);
                }(_pat_let19_0);
              }(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv25,_pat_let_tv26)).length));
            }(_pat_let18_0);
          }(_954_v79));
          _nw163[(new BigNumber(8)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(9)).toNumber()] = _module.D7.create_DC18(_module.__default.fm3(_949_v74, _944_v69, p1, new BigNumber((_850_v1).cardinality()), globalState));
          _nw163[(new BigNumber(10)).toNumber()] = _module.__default.fm26(globalState);
          _nw163[(new BigNumber(11)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(12)).toNumber()] = _module.D7.create_DC18(p1);
          _nw163[(new BigNumber(13)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(14)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(15)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(16)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(17)).toNumber()] = _module.D7.create_DC18(p1);
          _nw163[(new BigNumber(18)).toNumber()] = _module.__default.fm26(globalState);
          _nw163[(new BigNumber(19)).toNumber()] = _module.D7.create_DC18((_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))]);
          _nw163[(new BigNumber(20)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(21)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(22)).toNumber()] = ((_849_v0) ? (_954_v79) : (_954_v79));
          _nw163[(new BigNumber(23)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(24)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(25)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(26)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(27)).toNumber()] = _954_v79;
          _nw163[(new BigNumber(28)).toNumber()] = _954_v79;
          _956_v81 = _nw163;
          let _index155 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_956_v81).length));
          (_956_v81)[_index155] = _954_v79;
          let _index156 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_956_v81).length));
          let _rhs143 = _954_v79;
          let _rhs144 = (_949_v74).Union(_949_v74);
          let _rhs145 = (_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))];
          let _lhs64 = _956_v81;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_956_v81).length));
          _lhs64[_lhs65] = _rhs143;
          _949_v74 = _rhs144;
          r1 = _rhs145;
          let _961_v82;
          _961_v82 = _dafny.Seq.UnicodeFromString("aixybf");
          _961_v82 = _961_v82;
          r1 = (_module.__default.safeDivisionInt((_938_v65)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_938_v65).length))], (_this).f0)).plus((_dafny.ZERO).minus((_this).f0));
        } else {
          let _962_v83;
          let _init20 = ((_963_v0) => function (_964_i5) {
            return _963_v0;
          })(_849_v0);
          let _nw164 = Array((new BigNumber(4)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw164.length); _i0_20++) {
            _nw164[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _962_v83 = _nw164;
          let _965_v84;
          _965_v84 = _dafny.Seq.UnicodeFromString("elco");
          let _index157 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          (_962_v83)[_index157] = (_937_v64).fm6(_849_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(87)), function (_966_i6) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }), _965_v84, (((_850_v1).contains(!(_849_v0))) ? ((_850_v1).get(!(_849_v0))) : (p1)), globalState);
          let _967_v85;
          let _nw165 = new _module.C3();
          _nw165.__ctor(new BigNumber((_965_v84).length));
          _967_v85 = _nw165;
          let _968_v86;
          _968_v86 = _dafny.Seq.of(_849_v0, _849_v0);
          let _969_v87;
          _969_v87 = _module.D7.create_DC16(_968_v86);
          let _970_v88;
          _970_v88 = _dafny.MultiSet.fromElements(_969_v87);
          let _971_v89;
          _971_v89 = new _dafny.CodePoint('d'.codePointAt(0));
          let _972_v90;
          _972_v90 = _dafny.Set.fromElements(_849_v0, _module.__default.fm2(_849_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_849_v0,p1)).length), _849_v0, _849_v0, globalState));
          let _973_v91;
          _973_v91 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_967_v85).f0)),(_this).f0);
          let _974_v92;
          _974_v92 = _dafny.MultiSet.fromElements(p1, (_this).f0);
          let _index158 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          let _rhs146 = (_dafny.ZERO).minus(_module.__default.fm3(_module.__default.fm27(_849_v0, _970_v88, p1, _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f0), (_967_v85).f0, (_this).f0, (_967_v85).f0), globalState), _971_v89, ((!(_849_v0)) ? (_module.__default.fm3(_972_v90, _971_v89, new BigNumber((_973_v91).length), (_967_v85).f0, globalState)) : ((((_974_v92).contains((_this).f0)) ? ((_974_v92).get((_this).f0)) : ((_this).f0)))), (_dafny.ZERO).minus(new BigNumber((_850_v1).cardinality())), globalState));
          let _rhs147 = (new BigNumber((_850_v1).cardinality())).isLessThan((new BigNumber(959)).minus(_module.__default.fm3(_972_v90, _971_v89, new BigNumber(999), new BigNumber((_965_v84).length), globalState)));
          let _rhs148 = (_dafny.MultiSet.fromElements(new BigNumber((_968_v86).length), new BigNumber(191), p1)).IsDisjointFrom(((_974_v92).update((_967_v85).f0, _module.__default.abs((_967_v85).f0))).Intersect(_module.__default.fm22(globalState)));
          let _rhs149 = (p1).multipliedBy((_967_v85).f0);
          let _rhs150 = _967_v85;
          let _lhs66 = _962_v83;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          r2 = _rhs146;
          _849_v0 = _rhs147;
          _lhs66[_lhs67] = _rhs148;
          r1 = _rhs149;
          _967_v85 = _rhs150;
          let _975_v93;
          _975_v93 = _dafny.Map.Empty.slice().updateUnsafe(((_849_v0) ? (_module.__default.fm17(p1, globalState)) : (_971_v89)),(_967_v85).f0);
          _975_v93 = (function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)))).update(_971_v89, _module.__default.abs(new BigNumber((_965_v84).length)))).Elements) {
              let _976_v94 = _compr_24;
              if (((_dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)))).update(_971_v89, _module.__default.abs(new BigNumber((_965_v84).length)))).contains(_976_v94)) {
                _coll24.push([_976_v94,(_967_v85).f0]);
              }
            }
            return _coll24;
          }()).Merge((_975_v93).Merge(_975_v93));
          let _977_v95;
          let _nw166 = new _module.C2();
          _nw166.__ctor((_967_v85).f0);
          _977_v95 = _nw166;
          let _978_v96;
          _978_v96 = _dafny.Map.Empty.slice().updateUnsafe(_977_v95,_965_v84);
          _965_v84 = (((_978_v96).contains(_977_v95)) ? ((_978_v96).get(_977_v95)) : (_965_v84));
          let _979_v97;
          _979_v97 = _dafny.Seq.of((_977_v95).f0);
          let _980_v98;
          _980_v98 = _module.D3.create_DC7(_979_v97);
          let _981_v99;
          _981_v99 = _dafny.Set.fromElements(_962_v83, _962_v83);
          let _982_v100;
          let _nw167 = Array((new BigNumber(21)).toNumber());
          _nw167[(_dafny.ZERO).toNumber()] = _937_v64;
          _nw167[(_dafny.ONE).toNumber()] = _937_v64;
          _nw167[(new BigNumber(2)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(3)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(4)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(5)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(6)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(7)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(8)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(9)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(10)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(11)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(12)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(13)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(14)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(15)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(16)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(17)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(18)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(19)).toNumber()] = _937_v64;
          _nw167[(new BigNumber(20)).toNumber()] = _937_v64;
          _982_v100 = _nw167;
          let _983_v101;
          _983_v101 = _dafny.Map.Empty.slice().updateUnsafe(_982_v100,!(_849_v0));
          let _984_v102;
          _984_v102 = _dafny.Set.fromElements(_937_v64);
          let _index159 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          let _index160 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          let _index161 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          let _rhs151 = (_962_v83)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length))];
          let _rhs152 = _980_v98;
          let _rhs153 = !((_981_v99).IsProperSubsetOf(_981_v99)) || (_849_v0);
          let _rhs154 = (((_983_v101).contains(_982_v100)) ? ((_983_v101).get(_982_v100)) : ((_984_v102).IsSubsetOf(_984_v102)));
          let _lhs68 = _962_v83;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          let _lhs70 = _962_v83;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          let _lhs72 = _962_v83;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length));
          _lhs68[_lhs69] = _rhs151;
          _980_v98 = _rhs152;
          _lhs70[_lhs71] = _rhs153;
          _lhs72[_lhs73] = _rhs154;
          let _rhs155 = _849_v0;
          let _rhs156 = new BigNumber(((_dafny.MultiSet.fromElements(_965_v84)).update(_dafny.Seq.update(_dafny.Seq.Concat(_965_v84, _965_v84), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_965_v84, _965_v84)).length)), _971_v89), _module.__default.abs(((_977_v95).f0).plus((_this).f0)))).cardinality());
          let _rhs157 = _dafny.Seq.of(_849_v0, (_962_v83)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length))], (_962_v83)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v83).length))]);
          _849_v0 = _rhs155;
          r2 = _rhs156;
          _968_v86 = _rhs157;
        }
        let _985_v103;
        _985_v103 = _module.D2.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-137)), function (_986_i7) {
  return new _dafny.CodePoint('h'.codePointAt(0));
}), (p1).isLessThan(p1));
        _985_v103 = _985_v103;
        if ((_849_v0) && (_849_v0)) {
          let _987_v104;
          let _nw168 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _987_v104 = _nw168;
          let _index162 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_987_v104).length));
          (_987_v104)[_index162] = new BigNumber(-560);
          let _index163 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_987_v104).length));
          (_987_v104)[_index163] = (_dafny.ZERO).minus(p1);
          let _988_v105;
          _988_v105 = new _dafny.CodePoint('n'.codePointAt(0));
          let _989_v106;
          _989_v106 = _dafny.Map.Empty.slice().updateUnsafe(_849_v0,_988_v105);
          _989_v106 = _dafny.Map.Empty.slice().updateUnsafe(!((!(!(_849_v0))) === (_849_v0)),_988_v105);
          _849_v0 = _849_v0;
          let _990_v107;
          _990_v107 = _dafny.Seq.UnicodeFromString("bcmmxy");
          _990_v107 = _990_v107;
          _849_v0 = !(_849_v0) || (_849_v0);
        } else {
          let _991_v108;
          let _nw169 = Array((_dafny.ONE).toNumber()).fill(false);
          _991_v108 = _nw169;
          let _index164 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_991_v108).length));
          (_991_v108)[_index164] = _849_v0;
          let _index165 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_991_v108).length));
          (_991_v108)[_index165] = _849_v0;
          let _992_v109;
          let _init21 = ((_993_p1) => function (_994_i8) {
            return (_994_i8).multipliedBy(_993_p1);
          })(p1);
          let _nw170 = Array((new BigNumber(14)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw170.length); _i0_21++) {
            _nw170[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _992_v109 = _nw170;
          let _995_v110;
          _995_v110 = _dafny.Seq.of(_992_v109);
          _992_v109 = (_995_v110)[_module.__default.safeIndex((_this).f0, new BigNumber((_995_v110).length))];
          let _index166 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_991_v108).length));
          (_991_v108)[_index166] = (new BigNumber(454)).isLessThan(((!(_849_v0)) ? (new BigNumber(-114)) : ((_this).f0)));
          r0 = (_this).f0;
          let _996_v111;
          let _nw171 = Array((new BigNumber(25)).toNumber()).fill([]);
          _996_v111 = _nw171;
          let _997_v112;
          let _init22 = ((_998_v108) => function (_999_i9) {
            return _module.D0.create_DC1((_998_v108)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_998_v108).length))]);
          })(_991_v108);
          let _nw172 = Array((new BigNumber(7)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw172.length); _i0_22++) {
            _nw172[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _997_v112 = _nw172;
          let _index167 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_996_v111).length));
          (_996_v111)[_index167] = _997_v112;
          let _1000_v113;
          _1000_v113 = _997_v112;
          let _1001_v114;
          _1001_v114 = _dafny.Seq.of(_997_v112, _997_v112, (_1000_v113));
          let _index168 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_996_v111).length));
          (_996_v111)[_index168] = (_1001_v114)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1001_v114).length))];
        }
        let _1002_v115;
        _1002_v115 = _dafny.Map.Empty.slice().updateUnsafe(_849_v0,true);
        let _1003_v116;
        _1003_v116 = _dafny.Seq.of(_849_v0);
        let _1004_v117;
        _1004_v117 = _dafny.Seq.of(new BigNumber((_1002_v115).length), (_this).f0);
        if ((((_1002_v115).contains(!((_1003_v116)[_module.__default.safeIndex(new BigNumber((_1004_v117).length), new BigNumber((_1003_v116).length))]))) ? ((_1002_v115).get(!((_1003_v116)[_module.__default.safeIndex(new BigNumber((_1004_v117).length), new BigNumber((_1003_v116).length))]))) : (false))) {
          _849_v0 = _849_v0;
          _849_v0 = _849_v0;
          _849_v0 = _849_v0;
          let _1005_v118;
          _1005_v118 = _dafny.Seq.UnicodeFromString("qp");
          let _1006_v119;
          _1006_v119 = _module.D0.create_DC0(false);
          let _1007_v120;
          _1007_v120 = new _dafny.CodePoint('r'.codePointAt(0));
          _1005_v118 = _dafny.Seq.update(_1005_v118, _module.__default.safeIndex(_module.__default.safeDivisionInt((_module.D3.create_DC8((_1006_v119).dtor_cf0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_850_v1).contains(_849_v0)) ? ((_850_v1).get(_849_v0)) : (p1)),(_this).f0)).length))).dtor_cf9, new BigNumber((_1005_v118).length)), new BigNumber((_1005_v118).length)), _1007_v120);
          let _1008_v121;
          let _1009_v122;
          let _out19;
          let _out20;
          let _outcollector3 = (_937_v64).m6(_module.__default.safeModuloInt(p1, (_this).f0), globalState);
          _out19 = _outcollector3[0];
          _out20 = _outcollector3[1];
          _1008_v121 = _out19;
          _1009_v122 = _out20;
        } else {
          let _1010_v123;
          let _init23 = function (_1011_i10) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          };
          let _nw173 = Array((new BigNumber(17)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw173.length); _i0_23++) {
            _nw173[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1010_v123 = _nw173;
          let _1012_v124;
          _1012_v124 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1013_v125;
          _1013_v125 = _1012_v124;
          let _index169 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1010_v123).length));
          (_1010_v123)[_index169] = _1013_v125;
          let _1014_v126;
          _1014_v126 = _dafny.Set.fromElements((_this).f0, (_this).f0);
          let _index170 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1010_v123).length));
          let _rhs158 = _1013_v125;
          let _rhs159 = (new BigNumber((((_850_v1).update(_849_v0, _module.__default.abs((_1004_v117)[_module.__default.safeIndex(p1, new BigNumber((_1004_v117).length))]))).update(_849_v0, _module.__default.abs(new BigNumber((_1004_v117).length)))).cardinality())).multipliedBy(p1);
          let _rhs160 = ((_this).f0).isLessThan(new BigNumber((_dafny.Seq.Concat(_1003_v116, _1003_v116)).length));
          let _rhs161 = _1004_v117;
          let _rhs162 = (_1014_v126).Intersect((_dafny.Set.fromElements((_this).f0)).Difference(_1014_v126));
          let _lhs74 = _1010_v123;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1010_v123).length));
          _lhs74[_lhs75] = _rhs158;
          r0 = _rhs159;
          _849_v0 = _rhs160;
          _1004_v117 = _rhs161;
          _1014_v126 = _rhs162;
          let _1015_v127;
          let _nw174 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1015_v127 = _nw174;
          let _1016_v128;
          _1016_v128 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-949),p1);
          let _1017_v129;
          _1017_v129 = _module.D7.create_DC17(_849_v0, _849_v0, _849_v0);
          let _pat_let_tv27 = _849_v0;
          let _1018_v130;
          _1018_v130 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let21_0) {
            return function (_1019_dt__update__tmp_h4) {
              return function (_pat_let22_0) {
                return function (_1020_dt__update_hcf27_h0) {
                  return _module.D7.create_DC17(_1020_dt__update_hcf27_h0, (_1019_dt__update__tmp_h4).dtor_cf28, (_1019_dt__update__tmp_h4).dtor_cf29);
                }(_pat_let22_0);
              }(_pat_let_tv27);
            }(_pat_let21_0);
          }(_1017_v129),(p1).plus((((_1016_v128).contains(new BigNumber((_dafny.Seq.UnicodeFromString("puclwjroo")).length))) ? ((_1016_v128).get(new BigNumber((_dafny.Seq.UnicodeFromString("puclwjroo")).length))) : (new BigNumber(967)))));
          let _1021_v131;
          _1021_v131 = _dafny.Seq.UnicodeFromString("msay");
          let _1022_v132;
          _1022_v132 = _dafny.Map.Empty.slice().updateUnsafe(_849_v0,_1021_v131);
          let _1023_v133;
          _1023_v133 = _dafny.Set.fromElements(_849_v0, _849_v0);
          let _rhs163 = _1015_v127;
          let _rhs164 = _1016_v128;
          let _rhs165 = (((_849_v0) ? (p1) : ((_this).f0))).plus(p1);
          let _rhs166 = _dafny.Map.Empty.slice().updateUnsafe(_1017_v129,new BigNumber((_dafny.Seq.update((((_1022_v132).contains(_849_v0)) ? ((_1022_v132).get(_849_v0)) : (_1021_v131)), _module.__default.safeIndex(p1, new BigNumber(((((_1022_v132).contains(_849_v0)) ? ((_1022_v132).get(_849_v0)) : (_1021_v131))).length)), _1012_v124)).length));
          let _rhs167 = _module.__default.fm3(_1023_v133, _1012_v124, p1, p1, globalState);
          _1015_v127 = _rhs163;
          _1016_v128 = _rhs164;
          r2 = _rhs165;
          _1018_v130 = _rhs166;
          r1 = _rhs167;
          let _1024_v134;
          let _nw175 = new _module.C2();
          _nw175.__ctor((_this).f0);
          _1024_v134 = _nw175;
          let _1025_v135;
          _1025_v135 = _dafny.Seq.of(_1024_v134, _1024_v134);
          let _1026_v136;
          let _nw176 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1026_v136 = _nw176;
          let _rhs168 = ((_849_v0) ? ((_985_v103).dtor_cf4) : (_dafny.Seq.UnicodeFromString("hmljar")));
          let _rhs169 = (_1003_v116)[_module.__default.safeIndex(p1, new BigNumber((_1003_v116).length))];
          let _rhs170 = _1025_v135;
          let _rhs171 = _1026_v136;
          let _rhs172 = _849_v0;
          _1021_v131 = _rhs168;
          _849_v0 = _rhs169;
          _1025_v135 = _rhs170;
          _1026_v136 = _rhs171;
          _849_v0 = _rhs172;
          _1014_v126 = _1014_v126;
          _849_v0 = _849_v0;
        }
      }
      let _1027_v137;
      _1027_v137 = _dafny.Seq.of(p1);
      _1027_v137 = _1027_v137;
      if (((_849_v0) ? (_849_v0) : (_849_v0))) {
        let _1028_v138;
        _1028_v138 = _dafny.Set.fromElements(false, _849_v0);
        let _1029_v139;
        let _nw177 = new _module.C3();
        _nw177.__ctor((_dafny.ZERO).minus(new BigNumber(((_1028_v138).Intersect(_1028_v138)).length)));
        _1029_v139 = _nw177;
        let _1030_v140;
        _1030_v140 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1031_v141;
        let _nw178 = new _module.C0();
        _nw178.__ctor(_1030_v140);
        _1031_v141 = _nw178;
        if (_849_v0) {
          let _1032_v142;
          _1032_v142 = _dafny.MultiSet.fromElements((_this).f0);
          let _1033_v143;
          _1033_v143 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_1032_v142, p1, (_this).f0, (_this).f0, globalState),true);
          _1033_v143 = (_1033_v143).update(_module.__default.fm19(_1032_v142, p1, new BigNumber((_1027_v137).length), (_this).f0, globalState), !(_849_v0));
          let _1034_v144;
          _1034_v144 = _dafny.Seq.UnicodeFromString("jwm");
          _1034_v144 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), ((_1035_v140) => function (_1036_i11) {
            return _1035_v140;
          })(_1030_v140));
          let _1037_v145;
          let _nw179 = new _module.C1();
          _nw179.__ctor((_this).f0);
          _1037_v145 = _nw179;
          let _1038_v146;
          let _nw180 = new _module.C1();
          _nw180.__ctor((_dafny.ZERO).minus((_this).f0));
          _1038_v146 = _nw180;
          _849_v0 = !(_dafny.MultiSet.fromElements(_849_v0, _849_v0, _849_v0, _849_v0)).equals((_850_v1).Difference(_850_v1));
        } else {
          let _1039_v147;
          _1039_v147 = _dafny.Seq.UnicodeFromString("iltk");
          _1039_v147 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(310)), ((_1040_v140) => function (_1041_i12) {
            return _1040_v140;
          })(_1030_v140)), _1039_v147), _dafny.Seq.UnicodeFromString("gclpd"));
          let _1042_v148;
          let _nw181 = Array((new BigNumber(24)).toNumber()).fill(false);
          _1042_v148 = _nw181;
          let _1043_v149;
          _1043_v149 = _dafny.Map.Empty.slice().updateUnsafe((_849_v0) && (_849_v0),_1042_v148);
          _1043_v149 = (_1043_v149).update((_module.D3.create_DC8(_849_v0, p1)).dtor_cf8, _1042_v148);
          r0 = (p1).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_1030_v140, (_1031_v141).f3), _1039_v147)).length));
          let _1044_v150;
          let _nw182 = new _module.C1();
          _nw182.__ctor((_this).f0);
          _1044_v150 = _nw182;
          r2 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f0), (_this).f0);
        }
        _849_v0 = _849_v0;
        _849_v0 = _849_v0;
      } else {
        _849_v0 = _849_v0;
        _849_v0 = _849_v0;
        let _1045_v151;
        _1045_v151 = _dafny.Seq.UnicodeFromString("nny");
        let _1046_v152;
        _1046_v152 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1047_v153;
        _1047_v153 = _dafny.Set.fromElements(_849_v0, _849_v0);
        let _1048_v154;
        _1048_v154 = _dafny.MultiSet.fromElements(new BigNumber((_1027_v137).length), p1);
        let _1049_v155;
        _1049_v155 = _module.D2.create_DC4(_dafny.Seq.update(_1045_v151, _module.__default.safeIndex((_this).f0, new BigNumber((_1045_v151).length)), _1046_v152), ((_this).f0).isLessThanOrEqualTo(_module.__default.fm3(_1047_v153, _1046_v152, (_this).f0, new BigNumber((_1048_v154).cardinality()), globalState)));
        _1049_v155 = _1049_v155;
        let _1050_v156;
        let _nw183 = new _module.C2();
        _nw183.__ctor(p1);
        _1050_v156 = _nw183;
        let _nw184 = new _module.C2();
        _nw184.__ctor((p1).multipliedBy(p1));
        _1050_v156 = _nw184;
        _1027_v137 = _1027_v137;
      }
      let _1051_v157;
      _1051_v157 = _dafny.Set.fromElements(true);
      let _1052_v158;
      _1052_v158 = _dafny.Seq.UnicodeFromString("kw");
      let _hi6 = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1051_v157,new BigNumber((_1052_v158).length))).length), p1);
      for (let _1053_i13 = (_1027_v137)[_module.__default.safeIndex((_this).f0, new BigNumber((_1027_v137).length))]; _1053_i13.isLessThan(_hi6); _1053_i13 = _1053_i13.plus(_dafny.ONE)) {
        _849_v0 = _849_v0;
        r0 = _module.__default.safeModuloInt((_this).f0, new BigNumber((_1051_v157).length));
        _849_v0 = !((new BigNumber((_dafny.MultiSet.fromElements(p1, (_this).f0)).cardinality())).isLessThan(new BigNumber((_1027_v137).length)));
        r2 = (_this).f0;
      }
      let _1054_v159;
      _1054_v159 = _dafny.Seq.of(_dafny.Seq.Concat(_1052_v158, _dafny.Seq.UnicodeFromString("j")));
      _1054_v159 = _1054_v159;
      let _1055_v160;
      _1055_v160 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1056_v161;
      _1056_v161 = _dafny.Set.fromElements(_1055_v160, _1055_v160, _1055_v160, new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)));
      r0 = _module.__default.safeModuloInt((new BigNumber(924)).plus(p1), _module.__default.safeModuloInt((_this).f0, new BigNumber((_1056_v161).length)));
      r1 = (_this).f0;
      let _1057_v162;
      _1057_v162 = _module.D0.create_DC1(_849_v0);
      let _1058_v163;
      _1058_v163 = _dafny.Seq.of(_1057_v162);
      let _1059_v164;
      _1059_v164 = _module.D2.create_DC3(_1058_v163);
      let _pat_let_tv28 = _1052_v158;
      let _pat_let_tv29 = p1;
      let _pat_let_tv30 = p1;
      r2 = function (_source12) {
        if (_source12.is_DC4) {
          let _1060___mcc_h6 = (_source12).cf4;
          let _1061___mcc_h7 = (_source12).cf5;
          let _1062_cf5 = _1061___mcc_h7;
          let _1063_cf4 = _1060___mcc_h6;
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC4(_pat_let_tv28, false)).dtor_cf5,(_dafny.ZERO).minus(_pat_let_tv29))).length);
        } else if (_source12.is_DC5) {
          return _pat_let_tv30;
        } else if (_source12.is_DC3) {
          let _1064___mcc_h8 = (_source12).cf3;
          let _1065_cf3 = _1064___mcc_h8;
          return _module.__default.safeModuloInt((_this).f0, (_this).f0);
        } else {
          let _1066___mcc_h9 = (_source12).cf6;
          let _1067_cf6 = _1066___mcc_h9;
          return (_this).f0;
        }
      }(_1059_v164);
      return [r0, r1, r2];
    }
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1068_v0;
      let _init24 = function (_1069_i0) {
        return _dafny.Seq.UnicodeFromString("kotlm");
      };
      let _nw185 = Array((new BigNumber(19)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw185.length); _i0_24++) {
        _nw185[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _1068_v0 = _nw185;
      let _1070_v1;
      _1070_v1 = _dafny.Seq.UnicodeFromString("pxch");
      let _index171 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1068_v0).length));
      (_1068_v0)[_index171] = _1070_v1;
      let _1071_v2;
      let _nw186 = new _module.C2();
      _nw186.__ctor(new BigNumber(282));
      _1071_v2 = _nw186;
      let _1072_v3;
      _1072_v3 = _dafny.MultiSet.fromElements(_1071_v2);
      let _1073_v4;
      _1073_v4 = false;
      let _1074_v5;
      _1074_v5 = _module.D2.create_DC4(_1070_v1, _1073_v4);
      let _index172 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1068_v0).length));
      (_1068_v0)[_index172] = (((_1072_v3).contains(_1071_v2)) ? (_dafny.Seq.Concat(_module.__default.fm18((_this).f0, globalState), _1070_v1)) : ((_1074_v5).dtor_cf4));
      let _hi7 = ((_this).f0).plus(new BigNumber(-45));
      for (let _1075_i1 = _module.__default.safeDivisionInt((_this).f0, (_dafny.ZERO).minus((_this).f0)); _1075_i1.isLessThan(_hi7); _1075_i1 = _1075_i1.plus(_dafny.ONE)) {
        _1070_v1 = (_1068_v0)[_module.__default.safeIndex(new BigNumber(383), new BigNumber((_1068_v0).length))];
        let _1076_v6;
        let _init25 = ((_1077_v4) => function (_1078_i2) {
          return _1077_v4;
        })(_1073_v4);
        let _nw187 = Array((new BigNumber(25)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw187.length); _i0_25++) {
          _nw187[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1076_v6 = _nw187;
        let _index173 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1076_v6).length));
        (_1076_v6)[_index173] = _1073_v4;
        let _1079_v7;
        _1079_v7 = _dafny.Seq.of(_1073_v4);
        let _1080_v8;
        _1080_v8 = _dafny.MultiSet.fromElements(_1073_v4, _1073_v4);
        let _index174 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1076_v6).length));
        (_1076_v6)[_index174] = (_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_1079_v7, _1079_v7), _module.__default.safeIndex(_1075_i1, new BigNumber((_dafny.Seq.Concat(_1079_v7, _1079_v7)).length)), _1073_v4))).IsDisjointFrom(((_1073_v4) ? (_1080_v8) : (_1080_v8)));
        let _index175 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1076_v6).length));
        (_1076_v6)[_index175] = (_1071_v2).fm5(globalState);
        _1073_v4 = !(_1073_v4);
      }
      let _1081_v9;
      _1081_v9 = _dafny.MultiSet.fromElements((_this).f0);
      let _1082_v10;
      _1082_v10 = _dafny.Seq.of(_1073_v4);
      let _1083_v11;
      _1083_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1082_v10).length),!(_1073_v4)),_1073_v4);
      let _1084_v12;
      _1084_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(904),false);
      let _index176 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1068_v0).length));
      (_1068_v0)[_index176] = _module.__default.fm19(_1081_v9, (_this).f0, (new BigNumber((_1083_v11).length)).plus((_this).f0), new BigNumber((_1084_v12).length), globalState);
      if ((_1071_v2).fm6(_1073_v4, _1070_v1, _1070_v1, ((_this).f0).plus((_this).f0), globalState)) {
        let _1085_v13;
        _1085_v13 = new BigNumber(47);
        _1085_v13 = _1085_v13;
        if (_1073_v4) {
          let _1086_v14;
          let _nw188 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
          _1086_v14 = _nw188;
          let _1087_v15;
          _1087_v15 = _dafny.Set.fromElements(new BigNumber(636));
          let _index177 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1086_v14).length));
          (_1086_v14)[_index177] = _1087_v15;
          let _index178 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1086_v14).length));
          (_1086_v14)[_index178] = (_1087_v15).Intersect(((_1073_v4) ? (_1087_v15) : (_1087_v15)));
          let _1088_v16;
          let _nw189 = new _module.C3();
          _nw189.__ctor(_1085_v13);
          _1088_v16 = _nw189;
          let _1089_v17;
          _1089_v17 = _dafny.Set.fromElements(_dafny.Set.fromElements(_1073_v4, _1073_v4));
          let _1090_v18;
          _1090_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1088_v16,(_1089_v17).Union(_1089_v17));
          _1090_v18 = (_1090_v18).update(_1088_v16, _1089_v17);
          _1073_v4 = (_1073_v4) || (!(_1073_v4));
          _1073_v4 = _1073_v4;
          _1085_v13 = new BigNumber(976);
        } else {
          _1073_v4 = _1073_v4;
          let _1091_v19;
          _1091_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1073_v4,(_dafny.ZERO).minus(_1085_v13));
          let _1092_v20;
          _1092_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1073_v4,_1091_v19);
          _1092_v20 = (_1092_v20).update(_1073_v4, _dafny.Map.Empty.slice().updateUnsafe(_1073_v4,new BigNumber((_1070_v1).length)));
          let _1093_v21;
          let _nw190 = new _module.C3();
          _nw190.__ctor(new BigNumber((_dafny.Seq.Concat(_1070_v1, _1070_v1)).length));
          _1093_v21 = _nw190;
          let _1094_v22;
          _1094_v22 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1095_v23;
          _1095_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1085_v13,_1094_v22);
          let _1096_v24;
          _1096_v24 = _dafny.MultiSet.fromElements(_1095_v23, _1095_v23, _dafny.Map.Empty.slice().updateUnsafe((_this).f0,_1094_v22), _1095_v23, _dafny.Map.Empty.slice().updateUnsafe((_this).f0,_1094_v22));
          let _1097_v25;
          _1097_v25 = _dafny.Seq.of(_1085_v13, (_dafny.ZERO).minus(_1085_v13), _1085_v13, _1085_v13, _1085_v13);
          let _1098_v26;
          _1098_v26 = _dafny.Map.Empty.slice().updateUnsafe((_1096_v24).Union(_1096_v24),(_1097_v25)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_1097_v25).length))]);
          _1098_v26 = (_1098_v26).update((_1096_v24).Difference(_dafny.MultiSet.fromElements(_1095_v23)), (_this).f0);
          let _1099_v27;
          let _nw191 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1099_v27 = _nw191;
          let _index179 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1099_v27).length));
          (_1099_v27)[_index179] = _1085_v13;
          let _index180 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1099_v27).length));
          (_1099_v27)[_index180] = (_1085_v13).multipliedBy((_this).f0);
          let _1100_v28;
          _1100_v28 = _dafny.MultiSet.fromElements(_1073_v4, _1073_v4);
          let _1101_v29;
          _1101_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1094_v22,((_1100_v28).update((_1093_v21).fm5(globalState), _module.__default.abs((_this).f0))).IsProperSubsetOf(_1100_v28));
          let _1102_v30;
          _1102_v30 = _dafny.Seq.of(_dafny.Seq.of(_1073_v4), _1082_v10);
          let _1103_v31;
          let _nw192 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1103_v31 = _nw192;
          let _1104_v32;
          _1104_v32 = _dafny.Set.fromElements(_1103_v31);
          let _index181 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1099_v27).length));
          let _index182 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1099_v27).length));
          let _rhs173 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1082_v10, (_1102_v30)[_module.__default.safeIndex(new BigNumber((_1100_v28).cardinality()), new BigNumber((_1102_v30).length))]), _1082_v10);
          let _rhs174 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1104_v32).length), _1085_v13));
          let _rhs175 = (_1085_v13).plus((_this).f0);
          let _rhs176 = _1101_v29;
          let _rhs177 = (_dafny.ZERO).minus((_1085_v13).multipliedBy((_this).f0));
          let _lhs76 = _1099_v27;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1099_v27).length));
          let _lhs78 = _1099_v27;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1099_v27).length));
          _1082_v10 = _rhs173;
          _lhs76[_lhs77] = _rhs174;
          _lhs78[_lhs79] = _rhs175;
          _1101_v29 = _rhs176;
          _1085_v13 = _rhs177;
        }
        let _rhs178 = _1073_v4;
        let _rhs179 = _1073_v4;
        let _rhs180 = _1073_v4;
        let _rhs181 = new BigNumber(-963);
        _1073_v4 = _rhs178;
        _1073_v4 = _rhs179;
        _1073_v4 = _rhs180;
        _1085_v13 = _rhs181;
        let _1105_v33;
        _1105_v33 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1106_v34;
        let _nw193 = new _module.C0();
        _nw193.__ctor(_1105_v33);
        _1106_v34 = _nw193;
        _1106_v34 = _1106_v34;
        _1085_v13 = _module.__default.safeDivisionInt((_this).f0, (_this).f0);
      } else {
        let _1107_v35;
        _1107_v35 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f0).isLessThan((_this).f0),(_this).f0);
        _1107_v35 = (_1107_v35).update((false) || (_1073_v4), (_this).f0);
        let _1108_v36;
        let _nw194 = Array((new BigNumber(19)).toNumber()).fill([]);
        _1108_v36 = _nw194;
        let _1109_v37;
        let _1110_v38;
        let _1111_v39;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector4 = (_1071_v2).m1(_1108_v36, (_this).f0, globalState);
        _out21 = _outcollector4[0];
        _out22 = _outcollector4[1];
        _out23 = _outcollector4[2];
        _1109_v37 = _out21;
        _1110_v38 = _out22;
        _1111_v39 = _out23;
        let _1112_v40;
        _1112_v40 = new _dafny.CodePoint('y'.codePointAt(0));
        _1112_v40 = _1112_v40;
        let _1113_v41;
        let _nw195 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _1113_v41 = _nw195;
        let _1114_v42;
        let _nw196 = new _module.C3();
        _nw196.__ctor((_this).f0);
        _1114_v42 = _nw196;
        let _1115_v43;
        _1115_v43 = _module.D10.create_DC21(_1114_v42);
        let _1116_v44;
        let _nw197 = Array((new BigNumber(7)).toNumber());
        _nw197[(_dafny.ZERO).toNumber()] = _1073_v4;
        _nw197[(_dafny.ONE).toNumber()] = _1073_v4;
        _nw197[(new BigNumber(2)).toNumber()] = _1073_v4;
        _nw197[(new BigNumber(3)).toNumber()] = false;
        _nw197[(new BigNumber(4)).toNumber()] = !(false);
        _nw197[(new BigNumber(5)).toNumber()] = _1073_v4;
        _nw197[(new BigNumber(6)).toNumber()] = _1073_v4;
        _1116_v44 = _nw197;
        let _1117_v45;
        _1117_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1115_v43,_1116_v44);
        let _index183 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1113_v41).length));
        (_1113_v41)[_index183] = _1117_v45;
        let _index184 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1113_v41).length));
        (_1113_v41)[_index184] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v43,_1116_v44);
        let _1118_v46;
        let _nw198 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1118_v46 = _nw198;
        let _index185 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1118_v46).length));
        (_1118_v46)[_index185] = (_dafny.ZERO).minus(_1111_v39);
        let _1119_v47;
        _1119_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1073_v4,_1073_v4);
        let _index186 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1118_v46).length));
        (_1118_v46)[_index186] = ((_dafny.ZERO).minus(new BigNumber((((!(_1073_v4)) ? (_1119_v47) : (_dafny.Map.Empty.slice().updateUnsafe(_1073_v4,_1073_v4)))).length))).multipliedBy(_1110_v38);
      }
      let _1120_v48;
      _1120_v48 = new BigNumber(-963);
      _1120_v48 = (_dafny.ZERO).minus((_this).f0);
      let _1121_v49;
      _1121_v49 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1122_v50;
      _1122_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1121_v49,_1120_v48);
      let _hi8 = new BigNumber(((_1122_v50).Merge(_1122_v50)).length);
      for (let _1123_i3 = (_this).f0; _1123_i3.isLessThan(_hi8); _1123_i3 = _1123_i3.plus(_dafny.ONE)) {
        let _1124_v51;
        let _nw199 = new _module.C3();
        _nw199.__ctor(new BigNumber(920));
        _1124_v51 = _nw199;
        let _1125_v52;
        _1125_v52 = _dafny.Seq.of(_1123_i3);
        let _1126_v53;
        _1126_v53 = _dafny.Seq.of(_1125_v52, _1125_v52, _1125_v52, _1125_v52, _1125_v52);
        let _1127_v54;
        _1127_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1126_v53,_dafny.Seq.of(_1120_v48, new BigNumber(847), (_this).f0));
        _1127_v54 = (_1127_v54).update(_dafny.Seq.update(_1126_v53, _module.__default.safeIndex(_1123_i3, new BigNumber((_1126_v53).length)), _1125_v52), _1125_v52);
        _1120_v48 = (_1120_v48).minus((_this).f0);
        let _1128_v55;
        let _nw200 = Array((new BigNumber(2)).toNumber());
        _nw200[(_dafny.ZERO).toNumber()] = new BigNumber(11);
        _nw200[(_dafny.ONE).toNumber()] = (_1123_i3).multipliedBy(_1120_v48);
        _1128_v55 = _nw200;
        let _index187 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1128_v55).length));
        (_1128_v55)[_index187] = _1123_i3;
        let _1129_v56;
        _1129_v56 = _dafny.Set.fromElements(_1073_v4, _1073_v4);
        let _index188 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1128_v55).length));
        (_1128_v55)[_index188] = _module.__default.safeModuloInt(_module.__default.fm3(_1129_v56, _1121_v49, (_this).f0, new BigNumber((_1082_v10).length), globalState), _1120_v48);
      }
      let _1130_v57;
      _1130_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1073_v4,_1073_v4);
      let _1131_v58;
      _1131_v58 = _dafny.Seq.of(_1130_v57);
      let _1132_v59;
      _1132_v59 = _dafny.Seq.of(_1081_v9);
      let _1133_v60;
      _1133_v60 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ecqjshdfv")).length),_1120_v48)).length));
      let _1134_v61;
      _1134_v61 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f0)), _dafny.MultiSet.fromElements(new BigNumber((_1133_v60).length), _1120_v48));
      r0 = ((_1130_v57).Merge((_1131_v58)[_module.__default.safeIndex((_this).f0, new BigNumber((_1131_v58).length))])).update(_1073_v4, (_dafny.MultiSet.FromArray(_1132_v59)).IsDisjointFrom(_1134_v61));
      return r0;
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = [];
      let r3 = _dafny.MultiSet.Empty;
      let _1135_v0;
      _1135_v0 = new BigNumber(565);
      _1135_v0 = (_this).f0;
      let _1136_v1;
      _1136_v1 = new _dafny.CodePoint('g'.codePointAt(0));
      _1136_v1 = _1136_v1;
      _1135_v0 = _1135_v0;
      if (p0) {
        let _1137_v2;
        let _nw201 = new _module.C0();
        _nw201.__ctor(_1136_v1);
        _1137_v2 = _nw201;
        let _1138_v3;
        _1138_v3 = _dafny.Set.fromElements(p1, p1);
        let _1139_v4;
        _1139_v4 = _dafny.Seq.of(new BigNumber((_1138_v3).length));
        let _1140_v5;
        _1140_v5 = _module.D0.create_DC1(true);
        let _1141_v6;
        _1141_v6 = _dafny.Seq.of(_module.D0.create_DC1(_module.__default.fm2(p1, new BigNumber((_1139_v4).length), p0, p1, globalState)), _1140_v5, _1140_v5);
        let _1142_v7;
        _1142_v7 = _module.D2.create_DC3(_1141_v6);
        let _1143_v8;
        _1143_v8 = _module.D2.create_DC6(((p1) ? (_1142_v7) : (_module.D2.create_DC6(_1142_v7))));
        _1143_v8 = _1143_v8;
        _1135_v0 = (_1135_v0).multipliedBy(_1135_v0);
        _1135_v0 = (_module.__default.safeModuloInt((_this).f0, (_dafny.ZERO).minus(_module.__default.fm3(_1138_v3, (_1137_v2).f3, (_1139_v4)[_module.__default.safeIndex((_this).f0, new BigNumber((_1139_v4).length))], new BigNumber(759), globalState)))).plus(_1135_v0);
        let _1144_v9;
        let _nw202 = new _module.C0();
        _nw202.__ctor((_1137_v2).f3);
        _1144_v9 = _nw202;
      } else {
        let _1145_v10;
        let _init26 = function (_1146_i0) {
          return (_1146_i0).plus((_this).f0);
        };
        let _nw203 = Array((new BigNumber(20)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw203.length); _i0_26++) {
          _nw203[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1145_v10 = _nw203;
        let _1147_v11;
        _1147_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(301),(_this).f0);
        let _1148_v12;
        _1148_v12 = _dafny.Seq.of(_1136_v1, _1136_v1);
        let _1149_v13;
        _1149_v13 = _dafny.Seq.of(new BigNumber((_1147_v11).length), (_this).f0, (_dafny.ZERO).minus(new BigNumber((_1148_v12).length)));
        let _1150_v14;
        _1150_v14 = _dafny.Seq.of(p1);
        let _index189 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
        (_1145_v10)[_index189] = ((_1149_v13)[_module.__default.safeIndex(new BigNumber((_1150_v14).length), new BigNumber((_1149_v13).length))]).minus(_1135_v0);
        let _1151_v15;
        _1151_v15 = _module.D3.create_DC7(_1149_v13);
        let _index190 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
        let _rhs182 = new BigNumber((_dafny.Seq.of((_1151_v15).dtor_cf7, _1149_v13)).length);
        let _rhs183 = _1135_v0;
        let _rhs184 = _1148_v12;
        let _lhs80 = _1145_v10;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
        _lhs80[_lhs81] = _rhs182;
        _1135_v0 = _rhs183;
        _1148_v12 = _rhs184;
        let _1152_v16;
        let _nw204 = Array((new BigNumber(18)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = p1;
        _nw204[(_dafny.ONE).toNumber()] = p0;
        _nw204[(new BigNumber(2)).toNumber()] = true;
        _nw204[(new BigNumber(3)).toNumber()] = p1;
        _nw204[(new BigNumber(4)).toNumber()] = p1;
        _nw204[(new BigNumber(5)).toNumber()] = p1;
        _nw204[(new BigNumber(6)).toNumber()] = (_this).fm6(p0, _1148_v12, _1148_v12, _1135_v0, globalState);
        _nw204[(new BigNumber(7)).toNumber()] = p0;
        _nw204[(new BigNumber(8)).toNumber()] = p0;
        _nw204[(new BigNumber(9)).toNumber()] = !(p0);
        _nw204[(new BigNumber(10)).toNumber()] = p0;
        _nw204[(new BigNumber(11)).toNumber()] = _module.__default.fm2(p1, _1135_v0, true, p1, globalState);
        _nw204[(new BigNumber(12)).toNumber()] = p0;
        _nw204[(new BigNumber(13)).toNumber()] = p1;
        _nw204[(new BigNumber(14)).toNumber()] = (p1) && (p1);
        _nw204[(new BigNumber(15)).toNumber()] = p0;
        _nw204[(new BigNumber(16)).toNumber()] = p0;
        _nw204[(new BigNumber(17)).toNumber()] = p0;
        _1152_v16 = _nw204;
        r1 = _1152_v16;
        r0 = !(p0) || (p1);
        let _index191 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
        (_1145_v10)[_index191] = ((_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))]).plus(_1135_v0);
        let _1153_v17;
        _1153_v17 = _module.D7.create_DC17(!(p1), !(p1), p1);
        let _source13 = _1153_v17;
        if (_source13.is_DC17) {
          let _1154___mcc_h0 = (_source13).cf27;
          let _1155___mcc_h1 = (_source13).cf28;
          let _1156___mcc_h2 = (_source13).cf29;
          let _1157_cf29 = _1156___mcc_h2;
          let _1158_cf28 = _1155___mcc_h1;
          let _1159_cf27 = _1154___mcc_h0;
          let _1160_v18;
          _1160_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1148_v12,_1152_v16);
          _1160_v18 = (_1160_v18).Merge(_1160_v18);
          let _1161_v19;
          let _nw205 = Array((new BigNumber(8)).toNumber());
          _1161_v19 = _nw205;
          let _1162_v20;
          _1162_v20 = _module.D3.create_DC9(_1135_v0, _1161_v19, p0);
          _1157_cf29 = (_1162_v20).dtor_cf12;
          _1145_v10 = _1145_v10;
          let _1163_v21;
          _1163_v21 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(!(_1159_cf27),new BigNumber((_dafny.Seq.of((_this).f0)).length))).update(_1158_cf28, (_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))]),(_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))]);
          let _1164_v22;
          _1164_v22 = _module.D7.create_DC18((_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))]);
          let _1165_v23;
          _1165_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1164_v22,_1135_v0);
          let _1166_v24;
          let _nw206 = new _module.C2();
          _nw206.__ctor(_1135_v0);
          _1166_v24 = _nw206;
          let _1167_v25;
          _1167_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1166_v24,p0);
          let _1168_v26;
          _1168_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_1165_v23).update(_1164_v22, new BigNumber((_1167_v25).length))).length));
          let _1169_v27;
          _1169_v27 = _dafny.Map.Empty.slice().updateUnsafe((((_1163_v21).contains(_1168_v26)) ? ((_1163_v21).get(_1168_v26)) : ((_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))])),_1148_v12);
          _1169_v27 = (_1169_v27).update((_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))], _1148_v12);
        } else if (_source13.is_DC18) {
          let _1170___mcc_h3 = (_source13).cf30;
          let _1171_cf30 = _1170___mcc_h3;
          let _1172_v28;
          _1172_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          _1172_v28 = (_1172_v28).update(!(p0) || (p0), p1);
          let _index192 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
          (_1145_v10)[_index192] = (new BigNumber(668)).minus((_this).f0);
          let _index193 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1152_v16).length));
          (_1152_v16)[_index193] = p0;
          let _1173_v29;
          _1173_v29 = _module.D2.create_DC4(_dafny.Seq.update(_1148_v12, _module.__default.safeIndex(_1171_cf30, new BigNumber((_1148_v12).length)), _1136_v1), p1);
          let _1174_v30;
          _1174_v30 = _module.D0.create_DC0((_1173_v29).dtor_cf5);
          let _index194 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1152_v16).length));
          let _rhs185 = p0;
          let _rhs186 = _1174_v30;
          let _lhs82 = _1152_v16;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1152_v16).length));
          _lhs82[_lhs83] = _rhs185;
          _1174_v30 = _rhs186;
          let _rhs187 = _1148_v12;
          let _rhs188 = p1;
          _1148_v12 = _rhs187;
          r0 = _rhs188;
        } else {
          let _1175___mcc_h4 = (_source13).cf26;
          let _1176_cf26 = _1175___mcc_h4;
          let _1177_v31;
          _1177_v31 = _dafny.Map.Empty.slice().updateUnsafe(((_1149_v13)[_module.__default.safeIndex(new BigNumber((_1148_v12).length), new BigNumber((_1149_v13).length))]).isLessThan(_1135_v0),((!(p1)) ? (p1) : (p0)));
          let _1178_v32;
          _1178_v32 = _module.D2.create_DC4(_1148_v12, p0);
          let _1179_v33;
          _1179_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1178_v32).dtor_cf5,_1148_v12);
          _1177_v31 = (_1177_v31).update(!((_1153_v17).dtor_cf29), (_this).fm6(_module.__default.fm2(p1, (_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))], p1, false, globalState), (((_1179_v33).contains(p0)) ? ((_1179_v33).get(p0)) : (_1148_v12)), _1148_v12, _1135_v0, globalState));
          let _1180_v34;
          let _nw207 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1180_v34 = _nw207;
          let _1181_v35;
          _1181_v35 = _dafny.MultiSet.fromElements(_1136_v1);
          let _index195 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1180_v34).length));
          (_1180_v34)[_index195] = _1181_v35;
          let _1182_v36;
          let _nw208 = new _module.C3();
          _nw208.__ctor((_this).f0);
          _1182_v36 = _nw208;
          let _1183_v37;
          _1183_v37 = _module.D10.create_DC21(_1182_v36);
          let _1184_v38;
          _1184_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1181_v35).cardinality()),_1183_v37);
          let _index196 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
          let _index197 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1180_v34).length));
          let _rhs189 = !(p0) || (!((new BigNumber((_1184_v38).length)).isLessThan(new BigNumber((_1148_v12).length))));
          let _rhs190 = p0;
          let _rhs191 = (_1182_v36).f0;
          let _rhs192 = _1181_v35;
          let _lhs84 = _1145_v10;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length));
          let _lhs86 = _1180_v34;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1180_v34).length));
          r0 = _rhs189;
          r0 = _rhs190;
          _lhs84[_lhs85] = _rhs191;
          _lhs86[_lhs87] = _rhs192;
          let _1185_v39;
          let _nw209 = new _module.C1();
          _nw209.__ctor((_1145_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1145_v10).length))]);
          _1185_v39 = _nw209;
          let _1186_v40;
          _1186_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1135_v0,_1185_v39);
          let _1187_v41;
          _1187_v41 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_1149_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(88)), ((_1188_v10) => function (_1189_i1) {
            return (_1188_v10)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1188_v10).length))];
          })(_1145_v10))),new BigNumber((_1186_v40).length));
          _1187_v41 = (_dafny.Map.Empty.slice().updateUnsafe(p0,(_1182_v36).f0)).Merge(_1187_v41);
          let _1190_v42;
          _1190_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1136_v1,_1149_v13);
          _1149_v13 = (((_1190_v42).contains(_1136_v1)) ? ((_1190_v42).get(_1136_v1)) : (_1149_v13));
        }
      }
      let _1191_v43;
      _1191_v43 = _dafny.Set.fromElements(p1, p1, true);
      let _hi9 = _module.__default.fm3(_1191_v43, _1136_v1, _1135_v0, (_this).f0, globalState);
      for (let _1192_i2 = (_this).f0; _1192_i2.isLessThan(_hi9); _1192_i2 = _1192_i2.plus(_dafny.ONE)) {
        let _1193_v44;
        _1193_v44 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1135_v0);
        _1193_v44 = (_1193_v44).update(p1, _1192_i2);
        let _1194_v45;
        let _nw210 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1194_v45 = _nw210;
        let _1195_v46;
        _1195_v46 = _dafny.Seq.UnicodeFromString("k");
        let _index198 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1194_v45).length));
        (_1194_v45)[_index198] = (_1135_v0).plus(new BigNumber((_1195_v46).length));
        let _1196_v47;
        let _init27 = ((_1197_p0) => function (_1198_i3) {
          return _1197_p0;
        })(p0);
        let _nw211 = Array((new BigNumber(7)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw211.length); _i0_27++) {
          _nw211[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1196_v47 = _nw211;
        let _1199_v48;
        _1199_v48 = _module.D11.create_DC24(_1196_v47, p1, !(p0));
        let _1200_v49;
        _1200_v49 = _module.D5.create_DC13(_1192_i2, _dafny.Set.fromElements(true, true, p0, p1, p1), (_1195_v46)[_module.__default.safeIndex(_1192_i2, new BigNumber((_1195_v46).length))], _1192_i2, p0);
        let _index199 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1194_v45).length));
        let _rhs193 = ((_this).f0).multipliedBy(new BigNumber((_1191_v43).length));
        let _rhs194 = _module.D11.create_DC24(_1196_v47, (_1200_v49).dtor_cf21, p1);
        let _rhs195 = (p1) || (p0);
        let _lhs88 = _1194_v45;
        let _lhs89 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1194_v45).length));
        _lhs88[_lhs89] = _rhs193;
        _1199_v48 = _rhs194;
        r0 = _rhs195;
        let _1201_v50;
        _1201_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1194_v45)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_1194_v45).length))],!(p1));
        let _rhs196 = _1192_i2;
        let _rhs197 = !(((_this).f0).isLessThan((_1194_v45)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_1194_v45).length))]));
        let _rhs198 = (_1201_v50).Merge(_1201_v50);
        let _rhs199 = (p1) === ((_1191_v43).IsProperSubsetOf(_1191_v43));
        _1135_v0 = _rhs196;
        r0 = _rhs197;
        _1201_v50 = _rhs198;
        r0 = _rhs199;
        let _1202_v51;
        _1202_v51 = _dafny.Map.Empty.slice().updateUnsafe((_1194_v45)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_1194_v45).length))],_1192_i2);
        let _1203_v52;
        _1203_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1202_v51,_1194_v45);
        _1203_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1202_v51,_1194_v45);
      }
      let _1204_v53;
      _1204_v53 = _module.D7.create_DC17(p1, p1, (p1) === (p1));
      let _source14 = _1204_v53;
      if (_source14.is_DC17) {
        let _1205___mcc_h5 = (_source14).cf27;
        let _1206___mcc_h6 = (_source14).cf28;
        let _1207___mcc_h7 = (_source14).cf29;
        let _1208_cf29 = _1207___mcc_h7;
        let _1209_cf28 = _1206___mcc_h6;
        let _1210_cf27 = _1205___mcc_h5;
        let _1211_v54;
        _1211_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1135_v0,true);
        _1211_v54 = _1211_v54;
        _1135_v0 = _1135_v0;
        let _1212_v55;
        let _init28 = ((_1213_v0) => function (_1214_i4) {
          return _module.__default.safeDivisionInt(_1214_i4, _1213_v0);
        })(_1135_v0);
        let _nw212 = Array((new BigNumber(22)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw212.length); _i0_28++) {
          _nw212[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1212_v55 = _nw212;
        let _1215_v56;
        _1215_v56 = _dafny.Seq.of(!(p0));
        let _index200 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_1212_v55).length));
        (_1212_v55)[_index200] = new BigNumber((_dafny.Seq.Concat(_1215_v56, _1215_v56)).length);
        let _index201 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_1212_v55).length));
        (_1212_v55)[_index201] = _1135_v0;
        let _1216_v57;
        let _nw213 = new _module.C1();
        _nw213.__ctor((_this).f0);
        _1216_v57 = _nw213;
      } else if (_source14.is_DC18) {
        let _1217___mcc_h8 = (_source14).cf30;
        let _1218_cf30 = _1217___mcc_h8;
        let _1219_v58;
        let _init29 = ((_1220_v0) => function (_1221_i5) {
          return _dafny.Seq.of(_module.D5.create_DC11(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_this).f0, new BigNumber(810))).length)))), _module.D5.create_DC11(_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f0, _1220_v0, (_this).f0, new BigNumber((_dafny.Seq.UnicodeFromString("iqlvvtuaq")).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f0)).cardinality())))));
        })(_1135_v0);
        let _nw214 = Array((new BigNumber(9)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw214.length); _i0_29++) {
          _nw214[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1219_v58 = _nw214;
        let _1222_v59;
        let _nw215 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1222_v59 = _nw215;
        let _1223_v60;
        _1223_v60 = _dafny.MultiSet.fromElements(new BigNumber(-22), _1135_v0);
        let _1224_v61;
        _1224_v61 = _dafny.Seq.UnicodeFromString("kkw");
        let _rhs200 = _1222_v59;
        let _rhs201 = (((_1223_v60).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length))) ? ((_1223_v60).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length))) : (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("nqka")).length), new BigNumber((_dafny.Set.fromElements(p1, p1, p0)).length))));
        let _rhs202 = (_module.__default.safeDivisionInt(_1135_v0, (_this).f0)).minus(_module.__default.safeDivisionInt(_1135_v0, new BigNumber((_1224_v61).length)));
        let _rhs203 = _1219_v58;
        r1 = _rhs200;
        _1135_v0 = _rhs201;
        _1135_v0 = _rhs202;
        _1219_v58 = _rhs203;
        let _1225_v62;
        let _nw216 = Array((new BigNumber(18)).toNumber());
        _nw216[(_dafny.ZERO).toNumber()] = _1222_v59;
        _nw216[(_dafny.ONE).toNumber()] = ((!((_module.D0.create_DC1(p0)).dtor_cf1)) ? (_1222_v59) : (_1222_v59));
        _nw216[(new BigNumber(2)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(3)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(4)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(5)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(6)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(7)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(8)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(9)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(10)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(11)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(12)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(13)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(14)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(15)).toNumber()] = ((p0) ? (_1222_v59) : (_1222_v59));
        _nw216[(new BigNumber(16)).toNumber()] = _1222_v59;
        _nw216[(new BigNumber(17)).toNumber()] = _1222_v59;
        _1225_v62 = _nw216;
        let _index202 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1225_v62).length));
        (_1225_v62)[_index202] = _1222_v59;
        let _1226_v63;
        let _nw217 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1226_v63 = _nw217;
        let _1227_v64;
        _1227_v64 = _dafny.MultiSet.fromElements(p1, (_this).fm6((_this).fm5(globalState), _1224_v61, _1224_v61, _1218_cf30, globalState), p0);
        let _index203 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1225_v62).length));
        let _rhs204 = _dafny.Seq.Concat(_1224_v61, _1224_v61);
        let _rhs205 = _1226_v63;
        let _rhs206 = p1;
        let _rhs207 = _1222_v59;
        let _rhs208 = ((_dafny.MultiSet.fromElements(!(p0))).Union(_1227_v64)).IsProperSubsetOf(((_1227_v64).update(p1, _module.__default.abs(_1218_cf30))).Difference(_1227_v64));
        let _lhs90 = _1225_v62;
        let _lhs91 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1225_v62).length));
        _1224_v61 = _rhs204;
        r2 = _rhs205;
        r0 = _rhs206;
        _lhs90[_lhs91] = _rhs207;
        r0 = _rhs208;
        let _1228_v65;
        let _nw218 = new _module.C0();
        _nw218.__ctor(_1136_v1);
        _1228_v65 = _nw218;
        _1228_v65 = _1228_v65;
        let _index204 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_1226_v63).length));
        (_1226_v63)[_index204] = _1218_cf30;
        let _index205 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_1226_v63).length));
        (_1226_v63)[_index205] = (_1135_v0).plus((_this).f0);
      } else {
        let _1229___mcc_h9 = (_source14).cf26;
        let _1230_cf26 = _1229___mcc_h9;
        let _1231_v66;
        _1231_v66 = _dafny.Seq.UnicodeFromString("jqjpl");
        let _1232_v67;
        _1232_v67 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("vk"), _1231_v66);
        let _1233_v68;
        _1233_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,p1);
        let _1234_v69;
        _1234_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jut"),_1233_v68);
        r0 = (function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of (_1234_v69).Keys.Elements) {
            let _1235_v70 = _compr_25;
            if ((_1234_v69).contains(_1235_v70)) {
              _coll25.add(_1235_v70);
            }
          }
          return _coll25;
        }()).IsProperSubsetOf(_1232_v67);
        let _1236_v71;
        _1236_v71 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _1237_v72;
        _1237_v72 = _dafny.Set.fromElements(_1135_v0, new BigNumber((_1236_v71).length));
        r0 = !(_dafny.Set.fromElements((_this).f0, new BigNumber((_module.__default.fm0(globalState)).length), _1135_v0, _1135_v0)).equals(_1237_v72);
        _1136_v1 = _1136_v1;
        let _1238_v73;
        let _nw219 = new _module.C0();
        _nw219.__ctor(_1136_v1);
        _1238_v73 = _nw219;
      }
      r0 = ((_1191_v43).Union(_1191_v43)).IsSubsetOf(((p1) ? (_1191_v43) : (_1191_v43)));
      let _1239_v74;
      let _nw220 = Array((new BigNumber(7)).toNumber()).fill(false);
      _1239_v74 = _nw220;
      r1 = _1239_v74;
      let _1240_v75;
      _1240_v75 = _dafny.Seq.UnicodeFromString("shkq");
      let _1241_v76;
      _1241_v76 = _dafny.MultiSet.fromElements(_1135_v0);
      let _1242_v77;
      _1242_v77 = _dafny.Seq.of((_this).f0, _1135_v0);
      let _1243_v78;
      let _nw221 = Array((new BigNumber(13)).toNumber());
      _1243_v78 = _nw221;
      let _pat_let_tv31 = _1135_v0;
      let _1244_v79;
      let _nw222 = Array((new BigNumber(7)).toNumber());
      _nw222[(_dafny.ZERO).toNumber()] = (new BigNumber((_1240_v75).length)).multipliedBy(_1135_v0);
      _nw222[(_dafny.ONE).toNumber()] = _1135_v0;
      _nw222[(new BigNumber(2)).toNumber()] = _1135_v0;
      _nw222[(new BigNumber(3)).toNumber()] = new BigNumber(((_1241_v76).Difference(_dafny.MultiSet.FromArray(_1242_v77))).cardinality());
      _nw222[(new BigNumber(4)).toNumber()] = _1135_v0;
      _nw222[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((function (_pat_let23_0) {
        return function (_1245_dt__update__tmp_h0) {
          return function (_pat_let24_0) {
            return function (_1246_dt__update_hcf10_h0) {
              return _module.D3.create_DC9(_1246_dt__update_hcf10_h0, (_1245_dt__update__tmp_h0).dtor_cf11, (_1245_dt__update__tmp_h0).dtor_cf12);
            }(_pat_let24_0);
          }(_pat_let_tv31);
        }(_pat_let23_0);
      }(_module.D3.create_DC9((_this).f0, _1243_v78, p1))).dtor_cf10);
      _nw222[(new BigNumber(6)).toNumber()] = _1135_v0;
      _1244_v79 = _nw222;
      r2 = _1244_v79;
      let _1247_v80;
      _1247_v80 = _dafny.MultiSet.fromElements(false);
      r3 = ((_1247_v80).update(p0, _module.__default.abs(_1135_v0))).Intersect(_1247_v80);
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f0 = _dafny.ZERO;
      this._f1 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f2 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f1, f2, f0) {
      let _this = this;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f0 = f0;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return false;
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(573)), function (_1248_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("ilqiu")).length);
      }))).contains((_this).f0))) || ((true) === (true));
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new _dafny.CodePoint('n'.codePointAt(0));
    };
    m2(globalState) {
      let _this = this;
      let _1249_v0;
      let _nw223 = new _module.C0();
      _nw223.__ctor((_this).f1);
      _1249_v0 = _nw223;
      let _1250_v1;
      _1250_v1 = true;
      let _1251_v2;
      _1251_v2 = _dafny.Seq.of(true, _1250_v1);
      let _1252_v3;
      _1252_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1250_v1,_1251_v2);
      let _1253_v4;
      let _nw224 = Array((new BigNumber(9)).toNumber());
      _nw224[(_dafny.ZERO).toNumber()] = _1250_v1;
      _nw224[(_dafny.ONE).toNumber()] = (_1250_v1) === (_1250_v1);
      _nw224[(new BigNumber(2)).toNumber()] = _1250_v1;
      _nw224[(new BigNumber(3)).toNumber()] = _1250_v1;
      _nw224[(new BigNumber(4)).toNumber()] = _1250_v1;
      _nw224[(new BigNumber(5)).toNumber()] = _1250_v1;
      _nw224[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsPrefixOf(_1251_v2, (((_1252_v3).contains(_1250_v1)) ? ((_1252_v3).get(_1250_v1)) : (_1251_v2)));
      _nw224[(new BigNumber(7)).toNumber()] = !(_1250_v1) || (_1250_v1);
      _nw224[(new BigNumber(8)).toNumber()] = _1250_v1;
      _1253_v4 = _nw224;
      let _index206 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
      (_1253_v4)[_index206] = _1250_v1;
      let _1254_v5;
      _1254_v5 = _dafny.Map.Empty.slice().updateUnsafe(((_this).fm5(globalState)) || (_1250_v1),_dafny.MultiSet.fromElements(_1250_v1));
      let _1255_v6;
      _1255_v6 = _dafny.Seq.of(_1251_v2, _dafny.Seq.of(_1250_v1));
      let _1256_v7;
      _1256_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,(_1249_v0).f3);
      let _1257_v8;
      _1257_v8 = _dafny.Seq.UnicodeFromString("sargafyl");
      let _1258_v9;
      _1258_v9 = _dafny.Seq.of((((_1256_v7).contains(_1250_v1)) ? ((_1256_v7).get(_1250_v1)) : ((_this).f1)), (_1249_v0).f3, (_1257_v8)[_module.__default.safeIndex((_this).f2, new BigNumber((_1257_v8).length))], (_this).f1);
      let _1259_v10;
      _1259_v10 = _dafny.Seq.of((_1258_v9)[_module.__default.safeIndex((_this).f0, new BigNumber((_1258_v9).length))]);
      let _1260_v11;
      _1260_v11 = _dafny.MultiSet.fromElements(_1250_v1);
      let _index207 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
      let _rhs209 = ((_this).f0).isLessThan(new BigNumber((_dafny.Seq.update(_1255_v6, _module.__default.safeIndex((_this).f2, new BigNumber((_1255_v6).length)), _1251_v2)).length));
      let _rhs210 = _dafny.Seq.contains(_dafny.Seq.update(_1258_v9, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f0), new BigNumber((_1258_v9).length)), (_this).f1), (_this).f1);
      let _rhs211 = ((_1254_v5).update(_1250_v1, _dafny.MultiSet.fromElements(_1250_v1, !(_1250_v1), _1250_v1))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1250_v1,_1260_v11));
      let _lhs92 = _1253_v4;
      let _lhs93 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
      _lhs92[_lhs93] = _rhs209;
      _1250_v1 = _rhs210;
      _1254_v5 = _rhs211;
      let _1261_v12;
      _1261_v12 = new BigNumber(796);
      _1261_v12 = new BigNumber(234);
      let _1262_v13;
      _1262_v13 = _dafny.Seq.of(_1249_v0, _1249_v0, _1249_v0);
      let _1263_v14;
      _1263_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1260_v11,_1250_v1);
      let _1264_v15;
      _1264_v15 = _dafny.Seq.of(new BigNumber(182), (_dafny.ZERO).minus((_this).f0));
      if (!((_dafny.MultiSet.FromArray(_1264_v15)).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_1262_v13).length), new BigNumber((_1251_v2).length), new BigNumber((_dafny.Seq.UnicodeFromString("rdcu")).length), new BigNumber((_1263_v14).length)), _module.__default.safeIndex(_1261_v12, new BigNumber((_dafny.Seq.of(new BigNumber((_1262_v13).length), new BigNumber((_1251_v2).length), new BigNumber((_dafny.Seq.UnicodeFromString("rdcu")).length), new BigNumber((_1263_v14).length))).length)), (_this).f2)).length))))) {
        let _1265_v16;
        _1265_v16 = _dafny.Set.fromElements(_1261_v12, (_this).f2, (_this).f0);
        let _index208 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        let _rhs212 = !(new BigNumber(745)).isEqualTo((new BigNumber((_1264_v15).length)).minus(_1261_v12));
        let _rhs213 = _1261_v12;
        let _rhs214 = (_1250_v1) && (false);
        let _rhs215 = _1265_v16;
        let _lhs94 = _1253_v4;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        _lhs94[_lhs95] = _rhs212;
        _1261_v12 = _rhs213;
        _1250_v1 = _rhs214;
        _1265_v16 = _rhs215;
        let _1266_v17;
        let _init30 = ((_1267_v2) => function (_1268_i0) {
          return _1267_v2;
        })(_1251_v2);
        let _nw225 = Array((new BigNumber(12)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw225.length); _i0_30++) {
          _nw225[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1266_v17 = _nw225;
        let _index209 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1266_v17).length));
        (_1266_v17)[_index209] = _1251_v2;
        let _1269_v18;
        _1269_v18 = _dafny.Seq.of(_1253_v4);
        let _1270_v19;
        _1270_v19 = _module.D3.create_DC7(_1264_v15);
        let _index210 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1266_v17).length));
        let _rhs216 = (_1269_v18)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_1261_v12, new BigNumber(((_1270_v19).dtor_cf7).length)), new BigNumber((_1269_v18).length))];
        let _rhs217 = !(_1250_v1);
        let _rhs218 = ((_1250_v1) ? (_1251_v2) : (_1251_v2));
        let _rhs219 = (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))];
        let _lhs96 = _1266_v17;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1266_v17).length));
        _1253_v4 = _rhs216;
        _1250_v1 = _rhs217;
        _lhs96[_lhs97] = _rhs218;
        _1250_v1 = _rhs219;
        let _1271_v20;
        _1271_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_1253_v4);
        if (!(_1271_v20).contains((_dafny.ZERO).minus(_1261_v12))) {
          let _index211 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
          (_1253_v4)[_index211] = (((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]) ? (!(!((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]) || ((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]))) : ((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]));
          let _1272_v21;
          _1272_v21 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_1261_v12, new BigNumber(573)), ((_this).f2).plus(_1261_v12), (new BigNumber((_1257_v8).length)).plus((_this).f0));
          _1272_v21 = _1272_v21;
          let _1273_v22;
          let _nw226 = new _module.C0();
          _nw226.__ctor((_1249_v0).f3);
          _1273_v22 = _nw226;
          let _1274_v23;
          let _init31 = ((_1275_v9) => function (_1276_i1) {
            return _1275_v9;
          })(_1258_v9);
          let _nw227 = Array((new BigNumber(23)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw227.length); _i0_31++) {
            _nw227[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1274_v23 = _nw227;
          let _index212 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_1274_v23).length));
          (_1274_v23)[_index212] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("eblbebi"), _module.__default.safeIndex(new BigNumber(73), new BigNumber((_dafny.Seq.UnicodeFromString("eblbebi")).length)), (_1249_v0).f3), _1258_v9);
          let _index213 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_1274_v23).length));
          (_1274_v23)[_index213] = _dafny.Seq.Concat(_1258_v9, _1258_v9);
          let _index214 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
          (_1253_v4)[_index214] = (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))];
        } else {
          let _index215 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
          (_1253_v4)[_index215] = _1250_v1;
          let _index216 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
          (_1253_v4)[_index216] = (_this).fm5(globalState);
          _1261_v12 = _1261_v12;
          let _1277_v24;
          let _nw228 = new _module.C0();
          _nw228.__ctor((_1249_v0).f3);
          _1277_v24 = _nw228;
          let _1278_v25;
          _1278_v25 = _dafny.MultiSet.fromElements((_this).f2);
          let _index217 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
          (_1253_v4)[_index217] = !(_1278_v25).equals((_1278_v25).update((_this).f0, _module.__default.abs(new BigNumber(46))));
        }
        let _1279_v26;
        let _nw229 = Array((new BigNumber(12)).toNumber());
        _1279_v26 = _nw229;
        let _index218 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1279_v26).length));
        (_1279_v26)[_index218] = _1249_v0;
        let _1280_v27;
        _1280_v27 = (_this).f1;
        let _index219 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1279_v26).length));
        let _nw230 = new _module.C0();
        _nw230.__ctor((_1280_v27));
        (_1279_v26)[_index219] = _nw230;
        _1264_v15 = _1264_v15;
      } else {
        let _index220 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        let _index221 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        let _rhs220 = (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))];
        let _rhs221 = !(_1250_v1) || ((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]);
        let _lhs98 = _1253_v4;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        let _lhs100 = _1253_v4;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        _lhs98[_lhs99] = _rhs220;
        _lhs100[_lhs101] = _rhs221;
        let _1281_v28;
        _1281_v28 = _module.D2.create_DC5();
        let _1282_v29;
        _1282_v29 = _module.D3.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), ((_1283_v15) => function (_1284_i2) {
  return new BigNumber((_1283_v15).length);
})(_1264_v15)));
        let _1285_v30;
        _1285_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1250_v1,_1282_v29);
        let _1286_v31;
        _1286_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1261_v12, new BigNumber(-499), (_this).f0),(_1249_v0).f3);
        _1281_v28 = _module.__default.fm8((_dafny.ZERO).minus(_1261_v12), _1285_v30, _1286_v31, (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))], globalState);
        let _1287_v32;
        _1287_v32 = _module.D0.create_DC0(_1250_v1);
        _1250_v1 = (_1287_v32).dtor_cf0;
        let _index222 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
        (_1253_v4)[_index222] = !(!(((_1261_v12).minus(new BigNumber((_dafny.Seq.of(_1261_v12)).length))).isLessThanOrEqualTo(_1261_v12)));
        let _source15 = _1282_v29;
        if (_source15.is_DC8) {
          let _1288___mcc_h0 = (_source15).cf8;
          let _1289___mcc_h1 = (_source15).cf9;
          let _1290_cf9 = _1289___mcc_h1;
          let _1291_cf8 = _1288___mcc_h0;
          let _index223 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length));
          (_1253_v4)[_index223] = _1291_cf8;
          let _1292_v33;
          _1292_v33 = _module.D2.create_DC5();
          let _1293_v34;
          _1293_v34 = _module.D2.create_DC6(_1292_v33);
          _1293_v34 = _module.__default.fm9(globalState);
          let _1294_v35;
          _1294_v35 = _module.D3.create_DC8(_1250_v1, (_this).f0);
          _1261_v12 = _module.__default.safeDivisionInt((_1294_v35).dtor_cf9, _1290_cf9);
          let _1295_v36;
          _1295_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_1291_cf8),_1291_cf8);
          _1295_v36 = (_1295_v36).update(_1291_cf8, _1291_cf8);
        } else if (_source15.is_DC9) {
          let _1296___mcc_h2 = (_source15).cf10;
          let _1297___mcc_h3 = (_source15).cf11;
          let _1298___mcc_h4 = (_source15).cf12;
          let _1299_cf12 = _1298___mcc_h4;
          let _1300_cf11 = _1297___mcc_h3;
          let _1301_cf10 = _1296___mcc_h2;
          let _1302_v37;
          let _nw231 = new _module.C0();
          _nw231.__ctor((_1249_v0).f3);
          _1302_v37 = _nw231;
          let _1303_v38;
          _1303_v38 = _dafny.Seq.of(_1258_v9, _dafny.Seq.UnicodeFromString("lrbh"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xh"), _1257_v8));
          let _1304_v39;
          _1304_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1250_v1,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-525)), ((_1305_v8) => function (_1306_i4) {
            return _1305_v8;
          })(_1257_v8)));
          let _1307_v40;
          _1307_v40 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_1303_v38, _module.__default.safeIndex((_this).f0, new BigNumber((_1303_v38).length)), _1258_v9), _1303_v38), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-71)), ((_1308_v10) => function (_1309_i3) {
            return _1308_v10;
          })(_1259_v10)), (((_1304_v39).contains(!((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]))) ? ((_1304_v39).get(!((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]))) : (_1303_v38)));
          let _rhs222 = (((_1251_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(_1301_cf10), new BigNumber((_1251_v2).length))]) ? (_1302_v37) : (_1302_v37));
          let _rhs223 = (_1307_v40)[_module.__default.safeIndex(new BigNumber(-693), new BigNumber((_1307_v40).length))];
          _1302_v37 = _rhs222;
          _1303_v38 = _rhs223;
          let _1310_v41;
          _1310_v41 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), (_this).f1, (_1302_v37).f3, (_1249_v0).f3, (_1249_v0).f3);
          let _1311_v42;
          _1311_v42 = _dafny.Set.fromElements(_1299_cf12, _1299_cf12, (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))], (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]);
          let _1312_v43;
          _1312_v43 = _dafny.MultiSet.fromElements((_this).f2, (_this).f2, (_this).f0, _1261_v12);
          _1299_cf12 = ((_module.__default.fm3(_module.__default.fm10(_1310_v41, _1299_cf12, globalState), (_1249_v0).f3, (_this).f2, _module.__default.fm3(_1311_v42, (_1249_v0).f3, new BigNumber((_1258_v9).length), (_this).f2, globalState), globalState)).minus(_1301_cf10)).isLessThanOrEqualTo(new BigNumber(((_1312_v43).Difference(_dafny.MultiSet.fromElements((_this).f2, _1261_v12, (_this).f2))).cardinality()));
          let _1313_v44;
          _1313_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1299_cf12,!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("uuhjyw"), (_1302_v37).f3));
          _1313_v44 = (_1313_v44).update(!(_1250_v1), _1250_v1);
          _1261_v12 = ((_dafny.ZERO).minus(_1261_v12)).plus(_1261_v12);
        } else {
          let _1314___mcc_h5 = (_source15).cf7;
          let _1315_cf7 = _1314___mcc_h5;
          _1259_v10 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1258_v9, _1259_v10), _dafny.Seq.Concat(_1258_v9, _1259_v10));
          let _1316_v45;
          _1316_v45 = _module.D0.create_DC1(true);
          _1316_v45 = _module.D0.create_DC1((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]);
          let _1317_v46;
          _1317_v46 = _dafny.Map.Empty.slice().updateUnsafe((((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))]) ? (_1250_v1) : (_1250_v1)),false);
          _1317_v46 = (_1317_v46).update(!(_module.__default.fm2((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))], _1261_v12, (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))], (_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))], globalState)), !(!((((_1317_v46).contains(_1250_v1)) ? ((_1317_v46).get(_1250_v1)) : ((_this).fm5(globalState))))));
          let _1318_v47;
          let _nw232 = new _module.C0();
          _nw232.__ctor((_1249_v0).f3);
          _1318_v47 = _nw232;
        }
      }
      let _1319_v48;
      let _nw233 = new _module.C0();
      _nw233.__ctor((_1249_v0).f3);
      _1319_v48 = _nw233;
      let _source16 = _module.D0.create_DC0(_1250_v1);
      if (_source16.is_DC1) {
        let _1320___mcc_h6 = (_source16).cf1;
        let _1321_cf1 = _1320___mcc_h6;
        _1261_v12 = (_this).f0;
        _1261_v12 = (((!((_1253_v4)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1253_v4).length))])) ? (new BigNumber(701)) : ((_this).f2))).minus(new BigNumber((_1264_v15).length));
        let _1322_v50;
        _1322_v50 = _dafny.Set.fromElements(_1321_cf1, _1250_v1);
        let _1323_v51;
        _1323_v51 = _dafny.Seq.of(_1322_v50);
        let _1324_v52;
        _1324_v52 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),(_1323_v51)[_module.__default.safeIndex((_this).f0, new BigNumber((_1323_v51).length))]);
        let _1325_v53;
        _1325_v53 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3((((_1324_v52).contains(new _dafny.CodePoint('o'.codePointAt(0)))) ? ((_1324_v52).get(new _dafny.CodePoint('o'.codePointAt(0)))) : (_1322_v50)), (_1249_v0).f3, (_this).f0, _1261_v12, globalState),(_this).f0);
        _1261_v12 = ((_this).f2).minus(new BigNumber(((function () {
          let _coll26 = new _dafny.Map();
          for (const _compr_26 of (_1325_v53).Keys.Elements) {
            let _1326_v49 = _compr_26;
            if ((_1325_v53).contains(_1326_v49)) {
              _coll26.push([(_1326_v49).minus(new BigNumber((_dafny.Set.fromElements(false, _1250_v1)).length)),_1261_v12]);
            }
          }
          return _coll26;
        }()).Merge(_1325_v53)).length));
        _1250_v1 = true;
      } else {
        let _1327___mcc_h7 = (_source16).cf0;
        let _1328_cf0 = _1327___mcc_h7;
        let _1329_v54;
        let _nw234 = Array((new BigNumber(18)).toNumber());
        _1329_v54 = _nw234;
        let _index224 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1329_v54).length));
        (_1329_v54)[_index224] = _1249_v0;
        let _index225 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1329_v54).length));
        (_1329_v54)[_index225] = _1249_v0;
        let _1330_v55;
        _1330_v55 = _module.D0.create_DC0(_1328_cf0);
        let _1331_v56;
        _1331_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1330_v55,_1253_v4);
        _1331_v56 = (_1331_v56).update(_1330_v55, _1253_v4);
        let _1332_v57;
        let _nw235 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _1332_v57 = _nw235;
        let _1333_v58;
        _1333_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_1261_v12);
        let _index226 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1332_v57).length));
        (_1332_v57)[_index226] = new BigNumber(((_1333_v58).Merge(_1333_v58)).length);
        let _index227 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1332_v57).length));
        (_1332_v57)[_index227] = new BigNumber(929);
        _1250_v1 = (_this).fm5(globalState);
      }
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _1334_v0;
      _1334_v0 = false;
      let _1335_v1;
      _1335_v1 = _dafny.Set.fromElements(_1334_v0, _1334_v0, _1334_v0);
      let _hi10 = _module.__default.fm3(_1335_v1, (_this).f1, (_this).f2, new BigNumber(2), globalState);
      for (let _1336_i0 = p1; _1336_i0.isLessThan(_hi10); _1336_i0 = _1336_i0.plus(_dafny.ONE)) {
        if (_1334_v0) {
          let _1337_v2;
          _1337_v2 = _dafny.Set.fromElements((_this).f1);
          r0 = (_dafny.ZERO).minus((((_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).IsProperSubsetOf(_1337_v2)) ? ((_this).f0) : (p1)));
          let _1338_v3;
          _1338_v3 = _module.D3.create_DC8(true, (_this).f2);
          let _1339_v4;
          _1339_v4 = _dafny.Seq.UnicodeFromString("gmdnh");
          let _1340_v5;
          _1340_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_1338_v3).dtor_cf9, new BigNumber((_1339_v4).length))),new BigNumber(((_1335_v1).Difference(_1335_v1)).length));
          _1340_v5 = (_1340_v5).update(_1336_i0, (_this).f0);
          let _1341_v6;
          let _nw236 = new _module.C0();
          _nw236.__ctor((_this).f1);
          _1341_v6 = _nw236;
          _1334_v0 = _1334_v0;
          r1 = _module.__default.safeModuloInt(_1336_i0, (_this).f0);
        } else {
          let _1342_v7;
          let _nw237 = Array((_dafny.ONE).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = ((_1334_v0) ? (_1336_i0) : (_1336_i0));
          _1342_v7 = _nw237;
          let _1343_v8;
          let _nw238 = new _module.C3();
          _nw238.__ctor((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f2)));
          _1343_v8 = _nw238;
          let _1344_v9;
          _1344_v9 = _dafny.Seq.of(_1343_v8, _1343_v8, _1343_v8, _1343_v8);
          let _index228 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1342_v7).length));
          (_1342_v7)[_index228] = new BigNumber((_1344_v9).length);
          let _1345_v10;
          _1345_v10 = _dafny.Seq.of((_this).f2, p1, (_this).f2, (_this).f2);
          let _index229 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1342_v7).length));
          (_1342_v7)[_index229] = (_1345_v10)[_module.__default.safeIndex(((_this).f0).minus(new BigNumber(266)), new BigNumber((_1345_v10).length))];
          let _1346_v11;
          _1346_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v0,_1334_v0);
          let _1347_v12;
          _1347_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v0,new _dafny.CodePoint('p'.codePointAt(0)));
          let _index230 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1342_v7).length));
          let _rhs224 = ((_1334_v0) ? (_1334_v0) : ((_1335_v1).IsProperSubsetOf(_1335_v1)));
          let _rhs225 = false;
          let _rhs226 = (new BigNumber(((_1346_v11).Merge(_module.__default.fm20(_1334_v0, _1347_v12, _1334_v0, (_this).f1, globalState))).length)).plus(((_this).f0).multipliedBy(_1336_i0));
          let _rhs227 = new BigNumber(-863);
          let _rhs228 = _1342_v7;
          let _lhs102 = _1342_v7;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1342_v7).length));
          _1334_v0 = _rhs224;
          _1334_v0 = _rhs225;
          r0 = _rhs226;
          _lhs102[_lhs103] = _rhs227;
          _1342_v7 = _rhs228;
          let _1348_v13;
          _1348_v13 = _module.D0.create_DC1(_1334_v0);
          let _1349_v14;
          _1349_v14 = _module.D2.create_DC6(_module.D2.create_DC3(_dafny.Seq.of(_1348_v13, _1348_v13)));
          let _1350_v15;
          _1350_v15 = _module.D2.create_DC6(_1349_v14);
          let _1351_v16;
          _1351_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1350_v15,p1);
          _1334_v0 = (_dafny.Map.Empty.slice().updateUnsafe(_1350_v15,(_1343_v8).f0)).equals(_1351_v16);
          let _1352_v17;
          let _nw239 = new _module.C2();
          _nw239.__ctor(new BigNumber(409));
          _1352_v17 = _nw239;
          let _1353_v18;
          _1353_v18 = _dafny.Seq.UnicodeFromString("paqg");
          let _1354_v19;
          let _nw240 = Array((new BigNumber(22)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = _1353_v18;
          _nw240[(_dafny.ONE).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(2)).toNumber()] = _module.__default.fm18(p1, globalState);
          _nw240[(new BigNumber(3)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1353_v18, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("hibaahddh"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("hibaahddh")).length)), (_this).f1));
          _nw240[(new BigNumber(5)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(6)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("n");
          _nw240[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("gvptickwg");
          _nw240[(new BigNumber(9)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(10)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("g"), _1353_v18);
          _nw240[(new BigNumber(12)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(13)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("cidhpdr");
          _nw240[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_1355_v18) => function (_1356_i1) {
            return (_1355_v18)[_module.__default.safeIndex((_this).f2, new BigNumber((_1355_v18).length))];
          })(_1353_v18));
          _nw240[(new BigNumber(16)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("abkkeaax");
          _nw240[(new BigNumber(18)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(19)).toNumber()] = _1353_v18;
          _nw240[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_1357_i2) {
            return (_this).f1;
          });
          _nw240[(new BigNumber(21)).toNumber()] = _1353_v18;
          _1354_v19 = _nw240;
          let _index231 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1354_v19).length));
          (_1354_v19)[_index231] = _1353_v18;
          let _index232 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1354_v19).length));
          (_1354_v19)[_index232] = _dafny.Seq.Concat(_1353_v18, _1353_v18);
        }
        let _1358_v20;
        _1358_v20 = _dafny.Seq.UnicodeFromString("j");
        let _1359_v21;
        _1359_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_1358_v20);
        let _1360_v22;
        _1360_v22 = _dafny.MultiSet.fromElements(true);
        let _1361_v23;
        _1361_v23 = _dafny.MultiSet.fromElements(new BigNumber(749), p1);
        let _1362_v24;
        _1362_v24 = _dafny.Seq.of(_1334_v0);
        let _1363_v25;
        _1363_v25 = _module.D10.create_DC22(_module.__default.fm3(_module.__default.fm10(_dafny.MultiSet.fromElements((_this).f1), _1334_v0, globalState), (_this).f1, new BigNumber(((((_1359_v21).contains(new BigNumber((_1360_v22).cardinality()))) ? ((_1359_v21).get(new BigNumber((_1360_v22).cardinality()))) : (_module.__default.fm19(_1361_v23, new BigNumber((_1362_v24).length), _1336_i0, p1, globalState)))).length), _1336_i0, globalState), (_1334_v0) && (_1334_v0), _1334_v0);
        _1363_v25 = _1363_v25;
        r2 = (_dafny.ZERO).minus((_this).f2);
        _1334_v0 = ((_1334_v0) ? (!(_1334_v0)) : (_1334_v0));
      }
      let _hi11 = ((_this).f0).minus(p1);
      for (let _1364_i3 = (_this).f0; _1364_i3.isLessThan(_hi11); _1364_i3 = _1364_i3.plus(_dafny.ONE)) {
        let _1365_v26;
        let _nw241 = Array((new BigNumber(13)).toNumber());
        _1365_v26 = _nw241;
        let _1366_v27;
        _1366_v27 = _dafny.MultiSet.fromElements(_1334_v0);
        let _1367_v28;
        _1367_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1366_v27,!(_1334_v0));
        let _1368_v29;
        let _nw242 = new _module.C1();
        _nw242.__ctor(new BigNumber((_1367_v28).length));
        _1368_v29 = _nw242;
        let _index233 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_1365_v26).length));
        (_1365_v26)[_index233] = _1368_v29;
        let _index234 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_1365_v26).length));
        (_1365_v26)[_index234] = _1368_v29;
        let _1369_v30;
        _1369_v30 = _dafny.Seq.UnicodeFromString("yyrnjblm");
        _1369_v30 = _1369_v30;
        r1 = (_this).f0;
        let _1370_v31;
        _1370_v31 = _dafny.Seq.of(_1334_v0, true);
        let _1371_v32;
        _1371_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_1364_i3);
        let _1372_v33;
        _1372_v33 = _dafny.Seq.of(p1, (_this).f0);
        let _1373_v34;
        _1373_v34 = _dafny.Seq.of(_1372_v33);
        let _1374_v35;
        _1374_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v0,_1334_v0);
        let _1375_v36;
        _1375_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-707),true);
        let _1376_v37;
        let _nw243 = Array((new BigNumber(20)).toNumber());
        _nw243[(_dafny.ZERO).toNumber()] = (_1366_v27).Difference(_1366_v27);
        _nw243[(_dafny.ONE).toNumber()] = (_1366_v27).Union(_1366_v27);
        _nw243[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(_1334_v0, _1334_v0)).Union(_1366_v27);
        _nw243[(new BigNumber(3)).toNumber()] = (_1366_v27).update(_1334_v0, _module.__default.abs(_1364_i3));
        _nw243[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.FromArray(_1370_v31)).Difference(_1366_v27);
        _nw243[(new BigNumber(5)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(6)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.FromArray(_1370_v31)).update(_1334_v0, _module.__default.abs((((_1371_v32).contains(_1364_i3)) ? ((_1371_v32).get(_1364_i3)) : (new BigNumber((_1373_v34).length)))));
        _nw243[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(_1334_v0);
        _nw243[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_1334_v0, _1334_v0, _1334_v0);
        _nw243[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(_1334_v0);
        _nw243[(new BigNumber(11)).toNumber()] = (_1366_v27).update(_1334_v0, _module.__default.abs(new BigNumber(895)));
        _nw243[(new BigNumber(12)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(13)).toNumber()] = (_1366_v27).Difference(_1366_v27);
        _nw243[(new BigNumber(14)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(15)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(_1334_v0, _1334_v0, _1334_v0);
        _nw243[(new BigNumber(17)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(18)).toNumber()] = _1366_v27;
        _nw243[(new BigNumber(19)).toNumber()] = (((((_1374_v35).contains(_1334_v0)) ? ((_1374_v35).get(_1334_v0)) : (_1334_v0))) ? (_1366_v27) : (_module.__default.fm28(_1369_v30, true, new BigNumber((_1374_v35).length), _module.__default.fm2((((_1375_v36).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_1378_i4) {
          return (_this).f1;
        })).length))) ? ((_1375_v36).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_1377_i4) {
          return (_this).f1;
        })).length))) : (_1334_v0)), p1, _1334_v0, _1334_v0, globalState), globalState)));
        _1376_v37 = _nw243;
        let _index235 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1376_v37).length));
        (_1376_v37)[_index235] = _dafny.MultiSet.fromElements((_1368_v29).fm5(globalState));
        let _index236 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1376_v37).length));
        (_1376_v37)[_index236] = _1366_v27;
      }
      let _1379_v38;
      _1379_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_this).f2);
      let _1380_v39;
      _1380_v39 = _dafny.Set.fromElements(new BigNumber((_1379_v38).length), (_this).f2);
      let _1381_i5;
      _1381_i5 = _dafny.ZERO;
      L9: {
        while (!(_1334_v0) || ((_1380_v39).contains((_this).f0))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1381_i5)) {
              break L9;
            }
            _1381_i5 = (_1381_i5).plus(_dafny.ONE);
            _1334_v0 = !(((_this).f0).isLessThan((_this).f2));
            let _1382_v40;
            _1382_v40 = _dafny.Seq.UnicodeFromString("cjy");
            let _1383_v41;
            _1383_v41 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1334_v0,p1),(_dafny.ZERO).minus(new BigNumber((_1382_v40).length)));
            let _1384_v42;
            _1384_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v0,(_this).f0);
            _1383_v41 = (_1383_v41).update(_1384_v42, (_this).f2);
            _1334_v0 = _1334_v0;
            if (_1334_v0) {
              let _1385_v43;
              let _nw244 = Array((new BigNumber(3)).toNumber()).fill([]);
              _1385_v43 = _nw244;
              let _1386_v44;
              let _nw245 = Array((new BigNumber(11)).toNumber()).fill(false);
              _1386_v44 = _nw245;
              let _index237 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_1385_v43).length));
              (_1385_v43)[_index237] = _1386_v44;
              let _index238 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_1385_v43).length));
              (_1385_v43)[_index238] = _1386_v44;
              let _1387_v45;
              _1387_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1382_v40,(_this).f2);
              let _1388_v46;
              _1388_v46 = _dafny.MultiSet.fromElements((_this).f0, _module.__default.fm3(_1335_v1, new _dafny.CodePoint('p'.codePointAt(0)), (((_1387_v45).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_1390_i6) {
                return (_this).f1;
              }))) ? ((_1387_v45).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_1389_i6) {
                return (_this).f1;
              }))) : (p1)), new BigNumber(328), globalState));
              let _1391_v47;
              _1391_v47 = _dafny.MultiSet.fromElements(_1388_v46, _1388_v46);
              r0 = new BigNumber(((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f0), _1388_v46, _1388_v46, _dafny.MultiSet.fromElements((_this).f2), _1388_v46)).Difference(_1391_v47)).cardinality());
              let _1392_v48;
              let _nw246 = new _module.C0();
              _nw246.__ctor((_this).f1);
              _1392_v48 = _nw246;
              _1382_v40 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-405)), ((_1393_v48) => function (_1394_i7) {
                return (_1393_v48).f3;
              })(_1392_v48));
              _1380_v39 = function () {
                let _coll27 = new _dafny.Set();
                for (const _compr_27 of ((_dafny.Map.Empty.slice().updateUnsafe((_this).f0,p1)).Merge(function () {
                  let _coll28 = new _dafny.Map();
                  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(944), new BigNumber(-345))) {
                    let _1395_v49 = _compr_28;
                    if (((new BigNumber(944)).isLessThanOrEqualTo(_1395_v49)) && ((_1395_v49).isLessThan(new BigNumber(-345)))) {
                      _coll28.push([(_1395_v49).multipliedBy(p1),(_this).f2]);
                    }
                  }
                  return _coll28;
                }())).Keys.Elements) {
                  let _1396_v50 = _compr_27;
                  if (((_dafny.Map.Empty.slice().updateUnsafe((_this).f0,p1)).Merge(function () {
                    let _coll29 = new _dafny.Map();
                    for (const _compr_29 of _dafny.IntegerRange(new BigNumber(944), new BigNumber(-345))) {
                      let _1395_v49 = _compr_29;
                      if (((new BigNumber(944)).isLessThanOrEqualTo(_1395_v49)) && ((_1395_v49).isLessThan(new BigNumber(-345)))) {
                        _coll29.push([(_1395_v49).multipliedBy(p1),(_this).f2]);
                      }
                    }
                    return _coll29;
                  }())).contains(_1396_v50)) {
                    _coll27.add((_1396_v50).plus((_dafny.ZERO).minus(new BigNumber((((!(false)) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(204),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(872)), function (_1397_i8) {
                      return new _dafny.CodePoint('o'.codePointAt(0));
                    })).length))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(968),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("pp")).length))).length))).length))))).length))));
                  }
                }
                return _coll27;
              }();
            } else {
              let _1398_v51;
              let _nw247 = new _module.C1();
              _nw247.__ctor(new BigNumber(580));
              _1398_v51 = _nw247;
              _1398_v51 = _1398_v51;
              let _1399_v52;
              let _nw248 = new _module.C3();
              _nw248.__ctor((_this).f0);
              _1399_v52 = _nw248;
              let _1400_v53;
              _1400_v53 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_1335_v1, (_this).f1, (_this).f2, new BigNumber((_1382_v40).length), globalState),_1399_v52);
              let _1401_v54;
              _1401_v54 = _dafny.Seq.of(p1);
              let _1402_v55;
              _1402_v55 = _dafny.Seq.of(_1334_v0, _1334_v0);
              let _1403_v56;
              let _nw249 = Array((new BigNumber(11)).toNumber());
              _nw249[(_dafny.ZERO).toNumber()] = (_this).f2;
              _nw249[(_dafny.ONE).toNumber()] = new BigNumber((_1400_v53).length);
              _nw249[(new BigNumber(2)).toNumber()] = (_this).f2;
              _nw249[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_1401_v54)[_module.__default.safeIndex((_this).f0, new BigNumber((_1401_v54).length))]);
              _nw249[(new BigNumber(4)).toNumber()] = p1;
              _nw249[(new BigNumber(5)).toNumber()] = (_this).f2;
              _nw249[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber((_1379_v38).length));
              _nw249[(new BigNumber(7)).toNumber()] = (_this).f2;
              _nw249[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_this).f2);
              _nw249[(new BigNumber(9)).toNumber()] = ((_this).f0).plus(new BigNumber((_dafny.Set.fromElements(_1334_v0, _1334_v0)).length));
              _nw249[(new BigNumber(10)).toNumber()] = (p1).minus((_1398_v51).fm15(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), ((_1404_v1) => function (_1405_i9) {
                return (_module.D5.create_DC13(new BigNumber(209), _1404_v1, (_this).f1, (_this).f2, true)).dtor_cf19;
              })(_1335_v1)),(_this).fm5(globalState)), (_1402_v55)[_module.__default.safeIndex((_this).f2, new BigNumber((_1402_v55).length))], true, globalState));
              _1403_v56 = _nw249;
              let _index239 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1403_v56).length));
              (_1403_v56)[_index239] = ((_this).f2).minus(p1);
              let _1406_v57;
              _1406_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1382_v40,_1334_v0);
              let _1407_v58;
              _1407_v58 = _dafny.MultiSet.fromElements(_1334_v0, (_1402_v55)[_module.__default.safeIndex((_this).f0, new BigNumber((_1402_v55).length))], _1334_v0);
              let _index240 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1403_v56).length));
              let _rhs229 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_1401_v54)[_module.__default.safeIndex((_this).f0, new BigNumber((_1401_v54).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), function (_1408_i10) {
                return (_this).f1;
              })).length)));
              let _rhs230 = (_1398_v51).fm15(_1406_v57, false, ((_1379_v38).update((_1399_v52).f0, (_this).f0)).contains((((_1407_v58).contains(true)) ? ((_1407_v58).get(true)) : (_module.__default.fm3(_1335_v1, (_this).f1, (_this).f2, (_1399_v52).f0, globalState)))), globalState);
              let _rhs231 = (_1399_v52).f0;
              let _lhs104 = _1403_v56;
              let _lhs105 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1403_v56).length));
              r2 = _rhs229;
              _lhs104[_lhs105] = _rhs230;
              r0 = _rhs231;
              let _1409_v59;
              _1409_v59 = _module.D6.create_DC14(_1403_v56);
              _1409_v59 = _1409_v59;
              _1403_v56 = _1403_v56;
              _1334_v0 = _1334_v0;
            }
          }
        }
      }
      let _1410_v60;
      _1410_v60 = _dafny.Seq.UnicodeFromString("jiaaicmv");
      let _1411_v61;
      _1411_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1410_v60,_1334_v0);
      _1411_v61 = (_1411_v61).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), function (_1412_i11) {
        return (_this).f1;
      }), _1334_v0);
      let _1413_v62;
      let _nw250 = new _module.C2();
      _nw250.__ctor((_this).f0);
      _1413_v62 = _nw250;
      let _1414_v63;
      _1414_v63 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_1413_v62);
      let _1415_v64;
      _1415_v64 = _module.D6.create_DC15(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(299)), ((_1416_p1) => function (_1417_i12) {
  return _1416_p1;
})(p1)), _module.__default.fm0(globalState))).length), (_1414_v63).equals(_1414_v63), (_this).f0);
      let _source17 = _1415_v64;
      if (_source17.is_DC15) {
        let _1418___mcc_h0 = (_source17).cf23;
        let _1419___mcc_h1 = (_source17).cf24;
        let _1420___mcc_h2 = (_source17).cf25;
        let _1421_cf25 = _1420___mcc_h2;
        let _1422_cf24 = _1419___mcc_h1;
        let _1423_cf23 = _1418___mcc_h0;
        let _1424_v65;
        let _nw251 = new _module.C1();
        _nw251.__ctor((_1413_v62).f0);
        _1424_v65 = _nw251;
        let _source18 = (_this).f1;
        let _1425___mcc_h4 = _source18;
        let _1426_cf13 = _1425___mcc_h4;
        let _1427_v66;
        _1427_v66 = _module.D2.create_DC4(_dafny.Seq.Concat(_dafny.Seq.update(_1410_v60, _module.__default.safeIndex(_module.__default.fm3(_1335_v1, (_this).f1, (_this).f2, _1421_cf25, globalState), new BigNumber((_1410_v60).length)), (_1410_v60)[_module.__default.safeIndex((_this).f0, new BigNumber((_1410_v60).length))]), _1410_v60), !((_1334_v0) === (_1334_v0)));
        _1427_v66 = _1427_v66;
        let _pat_let_tv32 = _1410_v60;
        let _pat_let_tv33 = _1410_v60;
        let _pat_let_tv34 = _1426_cf13;
        let _rhs232 = function (_pat_let25_0) {
          return function (_1428_dt__update__tmp_h0) {
            return function (_pat_let26_0) {
              return function (_1429_dt__update_hcf4_h0) {
                return _module.D2.create_DC4(_1429_dt__update_hcf4_h0, (_1428_dt__update__tmp_h0).dtor_cf5);
              }(_pat_let26_0);
            }(_dafny.Seq.update(_pat_let_tv32, _module.__default.safeIndex((_this).f0, new BigNumber((_pat_let_tv33).length)), _pat_let_tv34));
          }(_pat_let25_0);
        }(_module.D2.create_DC4(_1410_v60, _1334_v0));
        let _rhs233 = _1334_v0;
        _1427_v66 = _rhs232;
        _1334_v0 = _rhs233;
        _1334_v0 = _1422_cf24;
        let _1430_v67;
        let _nw252 = Array((new BigNumber(26)).toNumber());
        _1430_v67 = _nw252;
        let _1431_v68;
        _1431_v68 = _module.D3.create_DC9(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_1334_v0),(_1413_v62).f0)).length), _1430_v67, _1334_v0);
        _1379_v38 = (_1379_v38).update(_1421_cf25, (_1431_v68).dtor_cf10);
        let _1432_v69;
        _1432_v69 = _dafny.MultiSet.fromElements(!(_1422_cf24));
        r2 = _module.__default.fm3((_1335_v1).Difference(_dafny.Set.fromElements(_1334_v0)), (_this).f1, (_dafny.ZERO).minus((_1424_v65).fm15(_1411_v61, _1334_v0, _1422_cf24, globalState)), ((_1334_v0) ? (new BigNumber(579)) : ((((_1432_v69).contains((_1424_v65).fm5(globalState))) ? ((_1432_v69).get((_1424_v65).fm5(globalState))) : ((_this).f2)))), globalState);
        _1379_v38 = _1379_v38;
      } else {
        let _1433___mcc_h3 = (_source17).cf22;
        let _1434_cf22 = _1433___mcc_h3;
        let _1435_v70;
        let _nw253 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
        _1435_v70 = _nw253;
        let _index241 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1434_cf22).length));
        (_1434_cf22)[_index241] = new BigNumber(-340);
        let _index242 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1434_cf22).length));
        let _rhs234 = _1435_v70;
        let _rhs235 = ((_this).f0).minus(_module.__default.safeModuloInt((_1413_v62).f0, p1));
        let _rhs236 = _dafny.Seq.UnicodeFromString("e");
        let _lhs106 = _1434_cf22;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1434_cf22).length));
        _1435_v70 = _rhs234;
        _lhs106[_lhs107] = _rhs235;
        _1410_v60 = _rhs236;
        _1334_v0 = _1334_v0;
        let _1436_v71;
        _1436_v71 = _dafny.MultiSet.fromElements((_1413_v62).f0, new BigNumber((_1410_v60).length));
        _1335_v1 = _dafny.Set.fromElements(_1334_v0, _1334_v0, (_1436_v71).IsSubsetOf(_1436_v71));
        let _1437_v72;
        _1437_v72 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f2).isLessThanOrEqualTo(p1),_1334_v0);
        _1437_v72 = ((_1334_v0) ? (_1437_v72) : (_1437_v72));
      }
      let _1438_v73;
      let _nw254 = Array((new BigNumber(8)).toNumber()).fill(false);
      _1438_v73 = _nw254;
      let _1439_v74;
      _1439_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v0,_1438_v73);
      let _1440_v75;
      _1440_v75 = _dafny.MultiSet.fromElements(p1);
      _1438_v73 = (((_1439_v74).contains(!(_1440_v75).equals(_1440_v75))) ? ((_1439_v74).get(!(_1440_v75).equals(_1440_v75))) : (_1438_v73));
      let _1441_v76;
      _1441_v76 = _dafny.Seq.of(p1);
      r0 = ((((_1379_v38).contains(new BigNumber(836))) ? ((_1379_v38).get(new BigNumber(836))) : ((_dafny.ZERO).minus((_1441_v76)[_module.__default.safeIndex(p1, new BigNumber((_1441_v76).length))])))).multipliedBy(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_1379_v38).length)), (_this).f2));
      r1 = (_this).f0;
      r2 = (_1413_v62).f0;
      return [r0, r1, r2];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _module.D2.Default();
      let _1442_v0;
      _1442_v0 = true;
      if (_1442_v0) {
        let _1443_v1;
        _1443_v1 = _dafny.Set.fromElements(_module.__default.fm2(!(_1442_v0), (p1).f0, _1442_v0, _1442_v0, globalState), false, ((_1442_v0) ? (_1442_v0) : (_1442_v0)), !(_1442_v0), _1442_v0);
        let _1444_v2;
        _1444_v2 = _dafny.Seq.of(_dafny.Set.fromElements(_1442_v0));
        let _1445_v3;
        _1445_v3 = _dafny.Seq.UnicodeFromString("hxmu");
        _1443_v1 = (_1444_v2)[_module.__default.safeIndex(_module.__default.fm3(_dafny.Set.fromElements(false, _1442_v0, _1442_v0), (_this).f1, new BigNumber((_1445_v3).length), (_this).f0, globalState), new BigNumber((_1444_v2).length))];
        let _1446_v4;
        _1446_v4 = _dafny.Set.fromElements(p0, p0);
        let _1447_v5;
        _1447_v5 = _dafny.MultiSet.fromElements((_this).f0, new BigNumber((_1446_v4).length));
        let _1448_v6;
        _1448_v6 = _dafny.Set.fromElements(_1447_v5, _dafny.MultiSet.fromElements(new BigNumber(-379)), _1447_v5);
        _1448_v6 = (_1448_v6).Difference(_1448_v6);
        let _1449_v7;
        let _nw255 = Array((new BigNumber(10)).toNumber()).fill([]);
        _1449_v7 = _nw255;
        let _1450_v10;
        let _nw256 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _1450_v10 = _nw256;
        let _1451_v11;
        _1451_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3((_1444_v2)[_module.__default.safeIndex((_this).f2, new BigNumber((_1444_v2).length))], (_this).f1, (_this).f2, new BigNumber((function () {
          let _coll30 = new _dafny.Set();
          for (const _compr_30 of ((function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(758), new BigNumber(-84))) {
              let _1452_v8 = _compr_31;
              if (((new BigNumber(758)).isLessThanOrEqualTo(_1452_v8)) && ((_1452_v8).isLessThan(new BigNumber(-84)))) {
                _coll31.push([_module.__default.safeModuloInt(_1452_v8, new BigNumber((_dafny.MultiSet.fromElements((_this).f1)).cardinality())),_1442_v0]);
              }
            }
            return _coll31;
          }()).update(new BigNumber(648), _1442_v0)).Keys.Elements) {
            let _1453_v9 = _compr_30;
            if (((function () {
              let _coll32 = new _dafny.Map();
              for (const _compr_32 of _dafny.IntegerRange(new BigNumber(758), new BigNumber(-84))) {
                let _1452_v8 = _compr_32;
                if (((new BigNumber(758)).isLessThanOrEqualTo(_1452_v8)) && ((_1452_v8).isLessThan(new BigNumber(-84)))) {
                  _coll32.push([_module.__default.safeModuloInt(_1452_v8, new BigNumber((_dafny.MultiSet.fromElements((_this).f1)).cardinality())),_1442_v0]);
                }
              }
              return _coll32;
            }()).update(new BigNumber(648), _1442_v0)).contains(_1453_v9)) {
              _coll30.add((_1453_v9).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(989), new BigNumber(-246))).length)));
            }
          }
          return _coll30;
        }()).length), globalState),_1450_v10);
        let _index243 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_1449_v7).length));
        (_1449_v7)[_index243] = (((_1451_v11).contains((p1).f0)) ? ((_1451_v11).get((p1).f0)) : (_1450_v10));
        let _index244 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_1449_v7).length));
        (_1449_v7)[_index244] = _1450_v10;
        let _1454_v12;
        _1454_v12 = new BigNumber(399);
        _1454_v12 = ((p1).f0).plus(_1454_v12);
        let _1455_v13;
        _1455_v13 = new _dafny.CodePoint('g'.codePointAt(0));
        _1455_v13 = _1455_v13;
      } else {
        let _1456_v14;
        _1456_v14 = _dafny.Seq.of(_1442_v0, _1442_v0);
        let _1457_v15;
        _1457_v15 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(_1442_v0, _1442_v0), _1456_v14), _dafny.Seq.of(_1442_v0));
        _1456_v14 = (_1457_v15)[_module.__default.safeIndex(((_1442_v0) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1442_v0,(_this).f1)).length)) : ((_dafny.ZERO).minus((_this).f0))), new BigNumber((_1457_v15).length))];
        let _1458_v16;
        _1458_v16 = _dafny.Seq.UnicodeFromString("biksd");
        let _1459_v17;
        _1459_v17 = _dafny.Seq.of(_1458_v16, _1458_v16);
        let _1460_v18;
        _1460_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1442_v0,false);
        let _1461_v19;
        _1461_v19 = _dafny.Seq.of(_1459_v17, _1459_v17);
        let _rhs237 = !(true) || ((((_1460_v18).contains((_1456_v14)[_module.__default.safeIndex((p1).f0, new BigNumber((_1456_v14).length))])) ? ((_1460_v18).get((_1456_v14)[_module.__default.safeIndex((p1).f0, new BigNumber((_1456_v14).length))])) : (true)));
        let _rhs238 = (_1461_v19)[_module.__default.safeIndex((_this).f0, new BigNumber((_1461_v19).length))];
        r0 = _rhs237;
        _1459_v17 = _rhs238;
        let _1462_v20;
        let _init32 = ((_1463_v0) => function (_1464_i0) {
          return _1463_v0;
        })(_1442_v0);
        let _nw257 = Array((new BigNumber(16)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw257.length); _i0_32++) {
          _nw257[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1462_v20 = _nw257;
        let _index245 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length));
        (_1462_v20)[_index245] = false;
        let _index246 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length));
        (_1462_v20)[_index246] = true;
        let _1465_v21;
        _1465_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f1,_1442_v0);
        let _1466_v22;
        _1466_v22 = _module.D7.create_DC17((_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))], !((((_1465_v21).contains((_this).f1)) ? ((_1465_v21).get((_this).f1)) : (!(!((_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))]))))) || ((_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))]), !((_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))]) || (true));
        let _1467_v23;
        _1467_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1460_v18).length),p1);
        _1466_v22 = _module.D7.create_DC17(false, _1442_v0, (_1467_v23).contains((_this).f0));
        let _1468_v24;
        _1468_v24 = _dafny.Set.fromElements((_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))]);
        let _1469_v25;
        _1469_v25 = _dafny.Set.fromElements((p1).f0, new BigNumber((_dafny.Seq.of((p1).f0, new BigNumber((_1468_v24).length))).length));
        let _1470_v26;
        _1470_v26 = _dafny.MultiSet.fromElements(new BigNumber((_1469_v25).length), new BigNumber((p2).length), (p1).f0);
        let _1471_v27;
        _1471_v27 = _module.D7.create_DC16(_1456_v14);
        let _1472_v28;
        _1472_v28 = _dafny.MultiSet.fromElements((_this).f1);
        let _1473_v29;
        let _nw258 = Array((new BigNumber(23)).toNumber());
        _nw258[(_dafny.ZERO).toNumber()] = _1456_v14;
        _nw258[(_dafny.ONE).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))]);
        _nw258[(new BigNumber(3)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(4)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(5)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(6)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_1442_v0, _1442_v0);
        _nw258[(new BigNumber(8)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(9)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(10)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(11)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(12)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(13)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(14)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(15)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(16)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(17)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(18)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(19)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(20)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(21)).toNumber()] = _1456_v14;
        _nw258[(new BigNumber(22)).toNumber()] = _1456_v14;
        _1473_v29 = _nw258;
        let _1474_v30;
        _1474_v30 = _dafny.Seq.of(_1473_v29);
        let _1475_v31;
        _1475_v31 = _dafny.MultiSet.fromElements((_1474_v30)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1474_v30).length))], _1473_v29);
        let _1476_v32;
        let _nw259 = new _module.C0();
        _nw259.__ctor(new _dafny.CodePoint('a'.codePointAt(0)));
        _1476_v32 = _nw259;
        let _1477_v33;
        _1477_v33 = _module.D13.create_DC26(_1476_v32);
        let _1478_v34;
        _1478_v34 = _dafny.Map.Empty.slice().updateUnsafe((p1).f0,_1476_v32);
        let _1479_v35;
        let _nw260 = Array((new BigNumber(27)).toNumber());
        _nw260[(_dafny.ZERO).toNumber()] = _1476_v32;
        _nw260[(_dafny.ONE).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(2)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(3)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(4)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(5)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(6)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(7)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(8)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(9)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(10)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(11)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(12)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(13)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(14)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(15)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(16)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(17)).toNumber()] = (_1477_v33).dtor_cf42;
        _nw260[(new BigNumber(18)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(19)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(20)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(21)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(22)).toNumber()] = (((_1478_v34).contains((p1).f0)) ? ((_1478_v34).get((p1).f0)) : (_1476_v32));
        _nw260[(new BigNumber(23)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(24)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(25)).toNumber()] = _1476_v32;
        _nw260[(new BigNumber(26)).toNumber()] = _1476_v32;
        _1479_v35 = _nw260;
        let _1480_v36;
        _1480_v36 = _module.D10.create_DC22(new BigNumber(968), (_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))], !(_1442_v0));
        let _1481_v37;
        _1481_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1442_v0,(p1).f0);
        let _1482_v38;
        _1482_v38 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1442_v0,_1442_v0));
        let _1483_v39;
        _1483_v39 = _module.D5.create_DC13((_this).f2, _1468_v24, (_this).f1, (_this).f2, _1442_v0);
        let _pat_let_tv35 = p1;
        let _pat_let_tv36 = _1476_v32;
        let _1484_v40;
        let _nw261 = Array((new BigNumber(28)).toNumber());
        _nw261[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-289), (_this).f0);
        _nw261[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1460_v18).length), (p1).f0);
        _nw261[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((new BigNumber(73)).plus((_this).f0));
        _nw261[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, (p1).f0));
        _nw261[(new BigNumber(4)).toNumber()] = (p1).f0;
        _nw261[(new BigNumber(5)).toNumber()] = (p1).f0;
        _nw261[(new BigNumber(6)).toNumber()] = ((_1442_v0) ? ((_this).f0) : ((p1).f0));
        _nw261[(new BigNumber(7)).toNumber()] = (_this).f0;
        _nw261[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt((p1).f0, new BigNumber(((_1470_v26).update((p1).f0, _module.__default.abs(new BigNumber(((_1471_v27).dtor_cf26).length)))).cardinality()));
        _nw261[(new BigNumber(9)).toNumber()] = new BigNumber(-526);
        _nw261[(new BigNumber(10)).toNumber()] = (p1).f0;
        _nw261[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm3(_module.__default.fm10(_1472_v28, (_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))], globalState), (_this).f1, (p1).f0, (_dafny.ZERO).minus((p1).f0), globalState));
        _nw261[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((p1).f0);
        _nw261[(new BigNumber(13)).toNumber()] = (p1).f0;
        _nw261[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt((_this).f2, (((_1475_v31).contains(_1473_v29)) ? ((_1475_v31).get(_1473_v29)) : ((_module.D3.create_DC9((p1).f0, _1479_v35, (_1462_v20)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1462_v20).length))])).dtor_cf10)));
        _nw261[(new BigNumber(15)).toNumber()] = (_1480_v36).dtor_cf34;
        _nw261[(new BigNumber(16)).toNumber()] = (p1).f0;
        _nw261[(new BigNumber(17)).toNumber()] = (_this).f0;
        _nw261[(new BigNumber(18)).toNumber()] = (_this).f2;
        _nw261[(new BigNumber(19)).toNumber()] = new BigNumber((_1458_v16).length);
        _nw261[(new BigNumber(20)).toNumber()] = new BigNumber((_1458_v16).length);
        _nw261[(new BigNumber(21)).toNumber()] = (_this).f0;
        _nw261[(new BigNumber(22)).toNumber()] = (_this).f0;
        _nw261[(new BigNumber(23)).toNumber()] = ((_dafny.ZERO).minus((p1).f0)).multipliedBy((_this).f2);
        _nw261[(new BigNumber(24)).toNumber()] = (p1).f0;
        _nw261[(new BigNumber(25)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1481_v37).length)), (p1).f0);
        _nw261[(new BigNumber(26)).toNumber()] = new BigNumber((_1482_v38).length);
        _nw261[(new BigNumber(27)).toNumber()] = new BigNumber(((function (_pat_let27_0) {
          return function (_1485_dt__update__tmp_h0) {
            return function (_pat_let28_0) {
              return function (_1486_dt__update_hcf17_h0) {
                return function (_pat_let29_0) {
                  return function (_1487_dt__update_hcf19_h0) {
                    return _module.D5.create_DC13(_1486_dt__update_hcf17_h0, (_1485_dt__update__tmp_h0).dtor_cf18, _1487_dt__update_hcf19_h0, (_1485_dt__update__tmp_h0).dtor_cf20, (_1485_dt__update__tmp_h0).dtor_cf21);
                  }(_pat_let29_0);
                }((_pat_let_tv36).f3);
              }(_pat_let28_0);
            }((_pat_let_tv35).f0);
          }(_pat_let27_0);
        }(_1483_v39)).dtor_cf18).length);
        _1484_v40 = _nw261;
        let _index247 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1484_v40).length));
        (_1484_v40)[_index247] = new BigNumber((_dafny.Seq.of(p0)).length);
        let _index248 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1484_v40).length));
        (_1484_v40)[_index248] = ((p1).f0).multipliedBy(new BigNumber(814));
      }
      let _1488_v41;
      let _init33 = function (_1489_i1) {
        return _module.__default.safeModuloInt(_1489_i1, new BigNumber(-498));
      };
      let _nw262 = Array((new BigNumber(16)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw262.length); _i0_33++) {
        _nw262[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1488_v41 = _nw262;
      let _1490_v42;
      _1490_v42 = _dafny.Seq.of(_1488_v41, _1488_v41);
      let _1491_v43;
      _1491_v43 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.update(p2, _module.__default.safeIndex(p0, new BigNumber((p2).length)), (_dafny.ZERO).minus((_this).f0)), _dafny.Seq.of((p1).f0, new BigNumber(531), new BigNumber((_1490_v42).length), p0, p0)));
      let _1492_v44;
      let _nw263 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1492_v44 = _nw263;
      let _index249 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1492_v44).length));
      (_1492_v44)[_index249] = (((_this).fm5(globalState)) ? (new _dafny.CodePoint('k'.codePointAt(0))) : ((_this).f1));
      let _1493_v45;
      _1493_v45 = new BigNumber(113);
      let _1494_v46;
      _1494_v46 = _dafny.Set.fromElements(true);
      let _1495_v47;
      _1495_v47 = _dafny.Seq.UnicodeFromString("davdt");
      let _index250 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1492_v44).length));
      let _rhs239 = ((_1491_v43).Union(_1491_v43)).Intersect((_dafny.MultiSet.fromElements(p2, p2)).update(p2, _module.__default.abs(_module.__default.fm3(_1494_v46, (_this).f1, (_this).f2, (p1).f0, globalState))));
      let _rhs240 = !(_1442_v0);
      let _rhs241 = (_this).f1;
      let _rhs242 = _module.__default.safeModuloInt((p1).f0, (_1493_v45).minus(new BigNumber((_1495_v47).length)));
      let _rhs243 = _1493_v45;
      let _lhs108 = _1492_v44;
      let _lhs109 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1492_v44).length));
      _1491_v43 = _rhs239;
      r0 = _rhs240;
      _lhs108[_lhs109] = _rhs241;
      _1493_v45 = _rhs242;
      _1493_v45 = _rhs243;
      let _1496_v48;
      let _nw264 = new _module.C0();
      _nw264.__ctor(new _dafny.CodePoint('h'.codePointAt(0)));
      _1496_v48 = _nw264;
      let _1497_v49;
      _1497_v49 = _module.D13.create_DC26(_1496_v48);
      let _source19 = _1497_v49;
      if (_source19.is_DC27) {
        let _1498___mcc_h0 = (_source19).cf43;
        let _1499___mcc_h1 = (_source19).cf44;
        let _1500___mcc_h2 = (_source19).cf45;
        let _1501___mcc_h3 = (_source19).cf46;
        let _1502_cf46 = _1501___mcc_h3;
        let _1503_cf45 = _1500___mcc_h2;
        let _1504_cf44 = _1499___mcc_h1;
        let _1505_cf43 = _1498___mcc_h0;
        let _1506_v50;
        let _nw265 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
        _1506_v50 = _nw265;
        let _1507_v51;
        let _nw266 = Array((new BigNumber(18)).toNumber()).fill(false);
        _1507_v51 = _nw266;
        let _1508_v52;
        _1508_v52 = _dafny.Seq.of(_1507_v51);
        let _index251 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1506_v50).length));
        (_1506_v50)[_index251] = _1508_v52;
        let _1509_v53;
        _1509_v53 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1507_v51)).length)));
        let _index252 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1506_v50).length));
        let _rhs244 = _1508_v52;
        let _rhs245 = _module.__default.fm22(globalState);
        let _lhs110 = _1506_v50;
        let _lhs111 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1506_v50).length));
        _lhs110[_lhs111] = _rhs244;
        _1509_v53 = _rhs245;
        _1502_cf46 = ((p2)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((p2).length))]).minus(((_1442_v0) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1495_v47).length)))) : ((_1504_cf44).f0)));
        let _1510_v54;
        let _init34 = ((_1511_v0) => function (_1512_i2) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_1511_v0), _dafny.Seq.of(_1511_v0));
        })(_1442_v0);
        let _nw267 = Array((new BigNumber(9)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw267.length); _i0_34++) {
          _nw267[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1510_v54 = _nw267;
        let _1513_v55;
        _1513_v55 = _dafny.Seq.of(_1510_v54, _1510_v54, _1510_v54);
        _1510_v54 = (_1513_v55)[_module.__default.safeIndex((p1).f0, new BigNumber((_1513_v55).length))];
        let _1514_v56;
        _1514_v56 = _module.D0.create_DC1(_1442_v0);
        _1514_v56 = _1514_v56;
      } else if (_source19.is_DC28) {
        let _1515___mcc_h4 = (_source19).cf47;
        let _1516___mcc_h5 = (_source19).cf48;
        let _1517_cf48 = _1516___mcc_h5;
        let _1518_cf47 = _1515___mcc_h4;
        let _1519_v57;
        let _init35 = function (_1520_i3) {
          return true;
        };
        let _nw268 = Array((new BigNumber(6)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw268.length); _i0_35++) {
          _nw268[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1519_v57 = _nw268;
        let _index253 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1519_v57).length));
        (_1519_v57)[_index253] = _1442_v0;
        let _1521_v58;
        _1521_v58 = _dafny.Seq.of(_1442_v0, _1442_v0);
        let _index254 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1519_v57).length));
        (_1519_v57)[_index254] = (_1518_cf47) && (_dafny.Seq.IsPrefixOf(_1521_v58, _1521_v58));
        _1493_v45 = (_this).f0;
        let _1522_v59;
        _1522_v59 = _dafny.MultiSet.fromElements(p0, new BigNumber(-389));
        let _1523_v60;
        _1523_v60 = _1522_v59;
        let _1524_v61;
        _1524_v61 = _dafny.Map.Empty.slice().updateUnsafe((_1523_v60),(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f2, (_dafny.ZERO).minus(p0)))).IsProperSubsetOf(_1522_v59));
        let _rhs246 = (p1).f0;
        let _rhs247 = _module.__default.fm29(_1442_v0, _1518_cf47, _1518_cf47, _1518_cf47, globalState);
        _1493_v45 = _rhs246;
        _1524_v61 = _rhs247;
        let _1525_v62;
        _1525_v62 = _dafny.Map.Empty.slice().updateUnsafe((_1521_v58)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("gisbfojc")).length), new BigNumber((_1521_v58).length))],_1518_cf47);
        let _1526_v63;
        _1526_v63 = _dafny.Map.Empty.slice().updateUnsafe((_1525_v62).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_1442_v0)),((p1).f0).isLessThan(p0));
        _1526_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1525_v62,((_this).f2).isLessThanOrEqualTo((p1).f0));
      } else if (_source19.is_DC29) {
        let _1527___mcc_h6 = (_source19).cf49;
        let _1528___mcc_h7 = (_source19).cf50;
        let _1529_cf50 = _1528___mcc_h7;
        let _1530_cf49 = _1527___mcc_h6;
        if (_1442_v0) {
          let _1531_v64;
          _1531_v64 = _dafny.Seq.of(_1442_v0);
          _1531_v64 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), ((_1532_v48) => function (_1533_i4) {
            return (_1532_v48).f3;
          })(_1496_v48)), _1495_v47), _1442_v0, _1529_cf50);
          let _1534_v65;
          _1534_v65 = _dafny.MultiSet.fromElements(new BigNumber(-707), (_this).f0);
          let _index255 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index255] = new BigNumber(((_dafny.MultiSet.FromArray(p2)).Difference(_1534_v65)).cardinality());
          let _index256 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index256] = _module.__default.safeModuloInt((_1493_v45).plus((p1).f0), (p1).f0);
          r0 = _1529_cf50;
          let _1535_v66;
          let _nw269 = new _module.C2();
          _nw269.__ctor((_1488_v41)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_1488_v41).length))]);
          _1535_v66 = _nw269;
          let _1536_v67;
          _1536_v67 = _dafny.MultiSet.fromElements(_1535_v66);
          _1529_cf50 = ((_1536_v67).IsSubsetOf(_1536_v67)) && (!(!_dafny.Seq.contains(_1495_v47, new _dafny.CodePoint('w'.codePointAt(0)))));
          _1529_cf50 = _1442_v0;
        } else {
          let _index257 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index257] = (_this).f2;
          let _index258 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index258] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.fm3(_dafny.Set.fromElements(_1442_v0), (_this).f1, (p1).f0, (p1).f0, globalState), p0));
          let _1537_v68;
          _1537_v68 = _dafny.Seq.of(false, _1529_cf50, _1442_v0);
          let _1538_v69;
          _1538_v69 = _module.D7.create_DC16(_dafny.Seq.of(_1442_v0, _1529_cf50, true, true));
          _1537_v68 = (_1538_v69).dtor_cf26;
          let _1539_v70;
          _1539_v70 = _dafny.MultiSet.fromElements(_1538_v69);
          let _1540_v71;
          _1540_v71 = _dafny.Map.Empty.slice().updateUnsafe((p1).f0,_1529_cf50);
          let _1541_v72;
          _1541_v72 = _dafny.MultiSet.fromElements(_1493_v45, (p1).f0);
          let _1542_v73;
          _1542_v73 = _dafny.Set.fromElements(p0, new BigNumber((_1541_v72).cardinality()));
          let _index259 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index259] = _module.__default.fm3(_1494_v46, (_1496_v48).f3, _module.__default.fm3(_module.__default.fm27(_1442_v0, _1539_v70, new BigNumber((_1540_v71).length), _1542_v73, globalState), (_1496_v48).f3, (p1).f0, (_this).f0, globalState), _module.__default.safeModuloInt((p1).f0, (p1).f0), globalState);
          _1493_v45 = new BigNumber((_dafny.Set.fromElements((_1488_v41)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_1488_v41).length))])).length);
          let _index260 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index260] = _module.__default.fm3((_1494_v46).Intersect(_1494_v46), new _dafny.CodePoint('f'.codePointAt(0)), new BigNumber((_1542_v73).length), (new BigNumber((_dafny.Seq.UnicodeFromString("to")).length)).plus((_this).f2), globalState);
        }
        let _index261 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_1488_v41).length));
        (_1488_v41)[_index261] = new BigNumber(86);
        let _1543_v74;
        _1543_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1530_cf49,(_1492_v44)[_module.__default.safeIndex(new BigNumber(776), new BigNumber((_1492_v44).length))]);
        let _index262 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_1488_v41).length));
        let _rhs248 = ((p0).plus(new BigNumber((_1495_v47).length))).multipliedBy((p2)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(_1493_v45, new BigNumber((p2).length)), (_this).f0)).length), new BigNumber((p2).length))]);
        let _rhs249 = new BigNumber((_1543_v74).length);
        let _rhs250 = !(_1442_v0);
        let _lhs112 = _1488_v41;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_1488_v41).length));
        _1493_v45 = _rhs248;
        _lhs112[_lhs113] = _rhs249;
        r0 = _rhs250;
        let _1544_v75;
        _1544_v75 = _dafny.Set.fromElements(p0, _1493_v45);
        _1530_cf49 = (_this).fm7((new BigNumber((_module.__default.fm18(new BigNumber((_dafny.MultiSet.fromElements(_1442_v0)).cardinality()), globalState)).length)).plus(new BigNumber((_dafny.Set.fromElements(!(!(_1529_cf50)), !(_1529_cf50))).length)), (_1488_v41)[_module.__default.safeIndex(new BigNumber(786), new BigNumber((_1488_v41).length))], _1529_cf50, (_dafny.ZERO).minus(_module.__default.fm3(_1494_v46, (_this).f1, new BigNumber((_1544_v75).length), (_dafny.ZERO).minus(p0), globalState)), globalState);
        let _1545_v76;
        _1545_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_1495_v47);
        let _1546_v77;
        _1546_v77 = _module.D2.create_DC4(_dafny.Seq.Concat(_1495_v47, (((_1545_v76).contains(p0)) ? ((_1545_v76).get(p0)) : (_dafny.Seq.update(_1495_v47, _module.__default.safeIndex((_1488_v41)[_module.__default.safeIndex(new BigNumber(786), new BigNumber((_1488_v41).length))], new BigNumber((_1495_v47).length)), (_1496_v48).f3)))), ((_this).f0).isLessThan((_1488_v41)[_module.__default.safeIndex(new BigNumber(786), new BigNumber((_1488_v41).length))]));
        r1 = _1546_v77;
      } else if (_source19.is_DC26) {
        let _1547___mcc_h8 = (_source19).cf42;
        let _1548_cf42 = _1547___mcc_h8;
        let _1549_v78;
        _1549_v78 = _dafny.MultiSet.fromElements((p1).f0);
        r0 = !((_1549_v78).IsProperSubsetOf(((_1442_v0) ? (_1549_v78) : (_1549_v78))));
        let _1550_v79;
        let _nw270 = Array((new BigNumber(6)).toNumber());
        _nw270[(_dafny.ZERO).toNumber()] = _1442_v0;
        _nw270[(_dafny.ONE).toNumber()] = _1442_v0;
        _nw270[(new BigNumber(2)).toNumber()] = _1442_v0;
        _nw270[(new BigNumber(3)).toNumber()] = _1442_v0;
        _nw270[(new BigNumber(4)).toNumber()] = _1442_v0;
        _nw270[(new BigNumber(5)).toNumber()] = _1442_v0;
        _1550_v79 = _nw270;
        let _1551_v80;
        _1551_v80 = _dafny.Seq.of(_1550_v79);
        _1493_v45 = _module.__default.safeModuloInt(new BigNumber((_1551_v80).length), (p1).f0);
        let _1552_v81;
        let _nw271 = Array((new BigNumber(9)).toNumber()).fill([]);
        _1552_v81 = _nw271;
        let _index263 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_1552_v81).length));
        (_1552_v81)[_index263] = _1488_v41;
        let _index264 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_1552_v81).length));
        let _rhs251 = _1488_v41;
        let _rhs252 = _1550_v79;
        let _rhs253 = (p1).f0;
        let _rhs254 = (_1442_v0) || (!_dafny.areEqual(_module.__default.fm19(_1549_v78, (p1).f0, (_this).f2, (p1).f0, globalState), _1495_v47));
        let _rhs255 = (((_dafny.MultiSet.FromArray(p2)).IsSubsetOf(_1549_v78)) ? (p0) : (((_this).f0).minus(new BigNumber(734))));
        let _lhs114 = _1552_v81;
        let _lhs115 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_1552_v81).length));
        _lhs114[_lhs115] = _rhs251;
        _1550_v79 = _rhs252;
        _1493_v45 = _rhs253;
        r0 = _rhs254;
        _1493_v45 = _rhs255;
        _1495_v47 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1495_v47, _1495_v47), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-630)), ((_1553_cf42) => function (_1554_i5) {
          return (_1553_cf42).f3;
        })(_1548_cf42)));
      } else {
        let _1555___mcc_h9 = (_source19).cf51;
        let _1556_cf51 = _1555___mcc_h9;
        let _1557_v82;
        _1557_v82 = _module.D0.create_DC0(_1442_v0);
        let _1558_v83;
        _1558_v83 = _dafny.Seq.of(_1557_v82);
        let _1559_v84;
        _1559_v84 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1558_v83).length)));
        let _1560_v85;
        _1560_v85 = _dafny.Seq.of(_1559_v84);
        let _1561_v86;
        _1561_v86 = _dafny.MultiSet.fromElements(new BigNumber((_1560_v85).length));
        let _1562_v87;
        _1562_v87 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Seq.UnicodeFromString("fyltafcwf")).length));
        let _1563_v88;
        _1563_v88 = _dafny.Seq.of(_1561_v86, _dafny.MultiSet.FromArray(_dafny.Seq.of(_1493_v45, p0, new BigNumber(222), (p1).f0)));
        let _1564_v90;
        _1564_v90 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll33 = new _dafny.Map();
          for (const _compr_33 of _dafny.IntegerRange(new BigNumber(798), new BigNumber(-379))) {
            let _1565_v89 = _compr_33;
            if (((new BigNumber(798)).isLessThanOrEqualTo(_1565_v89)) && ((_1565_v89).isLessThan(new BigNumber(-379)))) {
              _coll33.push([(_1565_v89).multipliedBy((_this).f0),new BigNumber((_1495_v47).length)]);
            }
          }
          return _coll33;
        }(),_dafny.MultiSet.fromElements(new BigNumber((_1495_v47).length), _module.__default.fm3(_1494_v46, (_1496_v48).f3, new BigNumber((_1559_v84).length), (_this).f0, globalState)));
        let _1566_v91;
        let _nw272 = Array((new BigNumber(24)).toNumber());
        _nw272[(_dafny.ZERO).toNumber()] = _1561_v86;
        _nw272[(_dafny.ONE).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(2)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(3)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber(-230))).update(new BigNumber((_1562_v87).length), _module.__default.abs((_this).f2));
        _nw272[(new BigNumber(5)).toNumber()] = _module.__default.fm22(globalState);
        _nw272[(new BigNumber(6)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f2));
        _nw272[(new BigNumber(8)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(9)).toNumber()] = (_1561_v86).update((_this).f0, _module.__default.abs(p0));
        _nw272[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements((p1).f0)).Intersect(_1561_v86);
        _nw272[(new BigNumber(11)).toNumber()] = (_1563_v88)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_1563_v88).length))];
        _nw272[(new BigNumber(12)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(13)).toNumber()] = (_1563_v88)[_module.__default.safeIndex((_this).f0, new BigNumber((_1563_v88).length))];
        _nw272[(new BigNumber(14)).toNumber()] = (_1561_v86).Intersect(_1561_v86);
        _nw272[(new BigNumber(15)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(16)).toNumber()] = (((_1564_v90).contains(_1562_v87)) ? ((_1564_v90).get(_1562_v87)) : (_1561_v86));
        _nw272[(new BigNumber(17)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(18)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(19)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(20)).toNumber()] = (_1561_v86).Difference(_module.__default.fm22(globalState));
        _nw272[(new BigNumber(21)).toNumber()] = _1561_v86;
        _nw272[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements((_this).f2, new BigNumber((_dafny.Seq.UnicodeFromString("cm")).length));
        _nw272[(new BigNumber(23)).toNumber()] = (_1561_v86).Intersect(_1561_v86);
        _1566_v91 = _nw272;
        let _index265 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1566_v91).length));
        (_1566_v91)[_index265] = _dafny.MultiSet.FromArray(p2);
        let _index266 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1566_v91).length));
        (_1566_v91)[_index266] = _1561_v86;
        let _1567_v92;
        let _nw273 = Array((new BigNumber(3)).toNumber());
        _1567_v92 = _nw273;
        let _1568_v93;
        let _nw274 = new _module.C3();
        _nw274.__ctor(_1493_v45);
        _1568_v93 = _nw274;
        let _index267 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1567_v92).length));
        (_1567_v92)[_index267] = _1568_v93;
        let _index268 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1567_v92).length));
        (_1567_v92)[_index268] = _1568_v93;
        let _1569_v94;
        _1569_v94 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(p2, p2),_1442_v0);
        _1569_v94 = (_1569_v94).Merge(_1569_v94);
        let _1570_v95;
        let _nw275 = Array((new BigNumber(12)).toNumber()).fill([]);
        _1570_v95 = _nw275;
        let _1571_v96;
        let _1572_v97;
        let _1573_v98;
        let _out24;
        let _out25;
        let _out26;
        let _outcollector5 = (p1).m1(_1570_v95, (_1493_v45).plus((_dafny.ZERO).minus(new BigNumber(-94))), globalState);
        _out24 = _outcollector5[0];
        _out25 = _outcollector5[1];
        _out26 = _outcollector5[2];
        _1571_v96 = _out24;
        _1572_v97 = _out25;
        _1573_v98 = _out26;
      }
      _1493_v45 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm0(globalState)).length));
      let _hi12 = _module.__default.safeDivisionInt((_this).f2, (_dafny.ZERO).minus((_this).f0));
      for (let _1574_i6 = (p1).f0; _1574_i6.isLessThan(_hi12); _1574_i6 = _1574_i6.plus(_dafny.ONE)) {
        if (!(_1442_v0)) {
          let _1575_v99;
          let _init36 = ((_1576_v0) => function (_1577_i7) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1576_v0,_dafny.Seq.of(_1576_v0, _1576_v0, _1576_v0));
          })(_1442_v0);
          let _nw276 = Array((_dafny.ONE).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw276.length); _i0_36++) {
            _nw276[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1575_v99 = _nw276;
          let _index269 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1575_v99).length));
          (_1575_v99)[_index269] = _module.__default.fm30(_1442_v0, _1442_v0, _1442_v0, globalState);
          let _1578_v100;
          _1578_v100 = _module.D13.create_DC28(_1442_v0, (_1496_v48).f3);
          let _1579_v101;
          _1579_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1442_v0,_dafny.Seq.of((_1578_v100).dtor_cf47, _1442_v0, _1442_v0, _1442_v0));
          let _index270 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1575_v99).length));
          (_1575_v99)[_index270] = _1579_v101;
          let _index271 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index271] = _module.__default.safeDivisionInt(new BigNumber((_1495_v47).length), (p1).f0);
          let _1580_v102;
          _1580_v102 = _dafny.Seq.of(_1442_v0);
          let _1581_v103;
          _1581_v103 = _dafny.MultiSet.fromElements((_1580_v102)[_module.__default.safeIndex((p1).f0, new BigNumber((_1580_v102).length))], !((_1580_v102)[_module.__default.safeIndex((p1).f0, new BigNumber((_1580_v102).length))]), _1442_v0, _1442_v0);
          let _index272 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_1488_v41).length));
          let _rhs256 = (_1581_v103).contains(_1442_v0);
          let _rhs257 = ((_1442_v0) ? ((p1).f0) : ((p1).f0));
          let _lhs116 = _1488_v41;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_1488_v41).length));
          r0 = _rhs256;
          _lhs116[_lhs117] = _rhs257;
          let _1582_v104;
          _1582_v104 = _dafny.Set.fromElements(_1495_v47);
          let _1583_v105;
          _1583_v105 = _dafny.Seq.of(p1);
          let _1584_v106;
          let _nw277 = new _module.C3();
          _nw277.__ctor((p1).f0);
          _1584_v106 = _nw277;
          let _1585_v107;
          _1585_v107 = _dafny.MultiSet.fromElements(_1584_v106);
          let _1586_v108;
          _1586_v108 = _dafny.MultiSet.fromElements((_this).f2, (p1).f0, _1493_v45, new BigNumber((_1583_v105).length), new BigNumber((_dafny.Seq.of((p1).f0, p0, new BigNumber((_1585_v107).cardinality()))).length));
          let _1587_v109;
          _1587_v109 = _module.D7.create_DC18((_1488_v41)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_1488_v41).length))]);
          _1442_v0 = (_this).fm6((_1582_v104).IsSubsetOf(_dafny.Set.fromElements(_1495_v47, _1495_v47)), _module.__default.fm19(_1586_v108, (_this).f2, (p1).f0, (_1587_v109).dtor_cf30, globalState), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vr"), _1495_v47), new BigNumber(-812), globalState);
          let _index273 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_1488_v41).length));
          (_1488_v41)[_index273] = (p1).f0;
          _1442_v0 = _1442_v0;
        } else {
          let _1588_v110;
          _1588_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1495_v47,_1442_v0);
          r0 = !(_1588_v110).equals(_1588_v110);
          let _1589_v111;
          _1589_v111 = _dafny.Map.Empty.slice().updateUnsafe((p1).f0,_module.__default.fm2(_1442_v0, (_this).f2, !(_1442_v0), _1442_v0, globalState));
          _1589_v111 = (_1589_v111).update(_1493_v45, false);
          _1493_v45 = (_1493_v45).minus(new BigNumber((_dafny.Seq.Concat(_1495_v47, _1495_v47)).length));
          let _1590_v112;
          let _nw278 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1590_v112 = _nw278;
          let _index274 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_1590_v112).length));
          (_1590_v112)[_index274] = _dafny.Seq.of((_1492_v44)[_module.__default.safeIndex(new BigNumber(776), new BigNumber((_1492_v44).length))], (_1496_v48).f3);
          let _index275 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_1590_v112).length));
          (_1590_v112)[_index275] = _1495_v47;
          let _1591_v113;
          _1591_v113 = _dafny.MultiSet.fromElements(_1574_i6);
          let _1592_v114;
          _1592_v114 = _dafny.Seq.of(_module.__default.fm19(_1591_v113, (p1).f0, (p1).f0, new BigNumber(180), globalState));
          let _1593_v115;
          _1593_v115 = _1592_v114;
          _1592_v114 = (_1593_v115);
        }
        let _1594_v116;
        _1594_v116 = _dafny.Seq.of(_1442_v0, ((_this).f0).isLessThanOrEqualTo(new BigNumber(-270)));
        _1594_v116 = _dafny.Seq.Concat(_dafny.Seq.update(_1594_v116, _module.__default.safeIndex(new BigNumber(939), new BigNumber((_1594_v116).length)), _1442_v0), _dafny.Seq.update(_dafny.Seq.of(true, _1442_v0, !(_1442_v0)), _module.__default.safeIndex((p1).f0, new BigNumber((_dafny.Seq.of(true, _1442_v0, !(_1442_v0))).length)), _1442_v0));
        let _index276 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1492_v44).length));
        (_1492_v44)[_index276] = (_1496_v48).f3;
        let _1595_v117;
        let _nw279 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Set.Empty);
        _1595_v117 = _nw279;
        let _index277 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_1595_v117).length));
        (_1595_v117)[_index277] = _module.__default.fm31(globalState);
        let _1596_v118;
        _1596_v118 = _dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)));
        let _index278 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_1595_v117).length));
        (_1595_v117)[_index278] = (_1596_v118).Union(_1596_v118);
      }
      let _1597_v119;
      _1597_v119 = _dafny.Seq.of(_1495_v47, _1495_v47, _1495_v47);
      _1495_v47 = _dafny.Seq.Concat(_1495_v47, _dafny.Seq.Concat((_1597_v119)[_module.__default.safeIndex((_this).f0, new BigNumber((_1597_v119).length))], _1495_v47));
      r0 = !_dafny.Seq.contains(_dafny.Seq.Concat(_1495_v47, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("sdiwmui"), _module.__default.safeIndex(_1493_v45, new BigNumber((_dafny.Seq.UnicodeFromString("sdiwmui")).length)), (_this).f1)), (_this).f1);
      r1 = _module.D2.create_DC4(_1495_v47, !(_1442_v0));
      return [r0, r1];
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
