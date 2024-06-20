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
    static fm0(p0, p1, p2, globalState) {
      return false;
    };
    static fm1(p0, globalState) {
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jkbuevdy"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(10)), function (_0_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }))).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), function (_1_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("irien"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(552)), function (_2_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nwfh"), _dafny.Seq.UnicodeFromString("hc")));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      if (!(false) || (false)) {
        return _module.D0.create_DC0(_dafny.Set.fromElements(!(!(true)), false, !(true), true));
      } else {
        return _module.D0.create_DC0(_dafny.Set.fromElements(true));
      }
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-33), new BigNumber(87)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(67)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xjhywi")).length), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.Seq.of(false)).length))));
    };
    static fm6(globalState) {
      return _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber(977)),false)).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),new BigNumber(786)),false)));
    };
    static fm7(globalState) {
      return _module.D2.create_DC6(_module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(!(false))),!(false))), _dafny.Seq.UnicodeFromString("jm"), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(469)), function (_3_i0) {
  return _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(false)),true));
}), _dafny.Seq.of(_module.D1.create_DC2(function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(!(false))),false)).Keys.Elements) {
    let _4_v0 = _compr_0;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(!(false))),false)).contains(_4_v0)) {
      _coll0.push([_4_v0,false]);
    }
  }
  return _coll0;
}()), _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(false)),false)), _module.D1.create_DC2(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.Seq.of(_module.D0.create_DC0(_dafny.Set.fromElements(false)), _module.D0.create_DC0(_dafny.Set.fromElements(true)))).Elements) {
    let _5_v1 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC0(_dafny.Set.fromElements(false)), _module.D0.create_DC0(_dafny.Set.fromElements(true))), _5_v1)) {
      _coll1.push([_5_v1,true]);
    }
  }
  return _coll1;
}()), _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(true, false, false, false)),!(true))), _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(true, true)),false))))).length));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      if (((!(!(false))) ? (false) : (!(false)))) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }
    };
    static fm9(p0, globalState) {
      return ((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(false, false, false, true))).Difference(_dafny.Set.fromElements(false));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-398), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(_module.D8.create_DC21(_module.D2.create_DC6(_module.D1.create_DC2(function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(false)),!(false))).Keys.Elements) {
    let _6_v0 = _compr_2;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(false)),!(false))).contains(_6_v0)) {
      _coll2.push([_6_v0,false]);
    }
  }
  return _coll2;
}()), _dafny.Seq.UnicodeFromString("gjxpwm"), new BigNumber(833)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(689)),true)))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("e")).length), new BigNumber((_dafny.Set.fromElements(false, true, true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(90),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false))).cardinality()))).length))).length)),new BigNumber(505))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(793),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("uwa")).length))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),new BigNumber((_dafny.Set.fromElements(false)).length))).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(162), new BigNumber(62))) {
          let _7_v1 = _compr_3;
          if (((new BigNumber(162)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(62)))) {
            _coll3.push([_module.__default.safeModuloInt(_7_v1, new BigNumber(859)),new BigNumber(979)]);
          }
        }
        return _coll3;
      }()));
    };
    static fm13(globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).Difference(_dafny.MultiSet.fromElements(true, !(!(false))))).Union((_dafny.MultiSet.fromElements(false, true, false, !(!(true)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))));
    };
    static fm14(p0, p1, p2, globalState) {
      if (!(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("cdpuwaavw"), _dafny.Seq.UnicodeFromString("jyc")))) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-171),new BigNumber(923))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(451),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), function (_8_i0) {
        return new BigNumber(444);
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(585))).length))).cardinality()))).length)));
    };
    static fm18(p0, p1, globalState) {
      return (_dafny.Set.fromElements(true, true, true)).Intersect(_dafny.Set.fromElements(false));
    };
    static fm19(p0, globalState) {
      return _dafny.MultiSet.fromElements(!(true));
    };
    static fm20(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('m'.codePointAt(0));
    };
    static fm21(globalState) {
      let _source0 = _module.D9.create_DC22(_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), function (_9_i0) {
  return new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality());
}));
      if (_source0.is_DC23) {
        let _10___mcc_h0 = (_source0).cf46;
        let _11___mcc_h1 = (_source0).cf47;
        let _12___mcc_h2 = (_source0).cf48;
        let _13___mcc_h3 = (_source0).cf49;
        let _14___mcc_h4 = (_source0).cf50;
        let _15_cf50 = _14___mcc_h4;
        let _16_cf49 = _13___mcc_h3;
        let _17_cf48 = _12___mcc_h2;
        let _18_cf47 = _11___mcc_h1;
        let _19_cf46 = _10___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_16_cf49, _17_cf48, !(_19_cf46)), _dafny.Seq.of(_19_cf46));
      } else {
        let _20___mcc_h5 = (_source0).cf45;
        let _21_cf45 = _20___mcc_h5;
        return _dafny.Seq.of((_module.D1.create_DC4(true)).dtor_cf8);
      }
    };
    static fm22(p0, globalState) {
      return _dafny.MultiSet.fromElements(!(false), (_dafny.Set.fromElements(new BigNumber(460))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wda")).length))), !((new BigNumber(-888)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("cbv")).length))).cardinality()))));
    };
    static fm24(p0, p1, p2, globalState) {
      return _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(true)),true));
    };
    static fm25(globalState) {
      let _source1 = _module.D7.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Seq.UnicodeFromString("rrvjgmae")).length)));
      if (_source1.is_DC17) {
        let _22___mcc_h0 = (_source1).cf31;
        let _23___mcc_h1 = (_source1).cf32;
        let _24_cf32 = _23___mcc_h1;
        let _25_cf31 = _22___mcc_h0;
        return (_25_cf31)[_module.__default.safeIndex(new BigNumber(137), new BigNumber((_25_cf31).length))];
      } else {
        let _26___mcc_h2 = (_source1).cf30;
        let _27_cf30 = _26___mcc_h2;
        return new _dafny.CodePoint('a'.codePointAt(0));
      }
    };
    static fm26(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-724))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(843)), function (_28_i0) {
        return _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length));
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(298)), function (_29_i1) {
        return _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cjmi")).length)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(53))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(756),!(true))).length), new BigNumber((_dafny.Seq.of(true, false)).length), new BigNumber(-95));
      }));
    };
    static fm27(p0, p1, p2, globalState) {
      let _source2 = _module.D12.create_DC30(_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_30_i0) {
  return _dafny.Seq.UnicodeFromString("ekdl");
}));
      if (_source2.is_DC31) {
        let _31___mcc_h0 = (_source2).cf60;
        let _32_cf60 = _31___mcc_h0;
        return _dafny.Seq.of(!(false));
      } else {
        let _33___mcc_h1 = (_source2).cf59;
        let _34_cf59 = _33___mcc_h1;
        return _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(false, !(false)));
      }
    };
    static fm28(p0, p1, globalState) {
      return _module.D5.create_DC12(new BigNumber(806));
    };
    static fm29(globalState) {
      return ((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), function (_35_i0) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length);
        }))).Elements) {
          let _36_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), function (_35_i0) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length);
          })), _36_v0)) {
            _coll4.push([_36_v0,new BigNumber((_dafny.Seq.UnicodeFromString("cvfumpis")).length)]);
          }
        }
        return _coll4;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(680),!(true))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true, false))).cardinality()))).cardinality())),new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(231), new BigNumber(-903))) {
          let _37_v1 = _compr_5;
          if (((new BigNumber(231)).isLessThanOrEqualTo(_37_v1)) && ((_37_v1).isLessThan(new BigNumber(-903)))) {
            _coll5.push([(_37_v1).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length))),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length),true)).length))]);
          }
        }
        return _coll5;
      }()).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(864)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),true)).length)));
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-996)), function (_38_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("tsneoe")),new BigNumber((_dafny.Seq.UnicodeFromString("tmfqkanh")).length));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements((_module.D10.create_DC26(true)).dtor_cf52))).Intersect(((false) ? (_dafny.Set.fromElements(!(false))) : (_dafny.Set.fromElements(false, false))));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _module.D9.create_DC22(_dafny.Seq.of(new BigNumber(998)));
    };
    static fm33(globalState) {
      return (_dafny.MultiSet.fromElements(true, false, !(false), false, true)).Union(_dafny.MultiSet.fromElements(false));
    };
    static fm34(globalState) {
      return ((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(279), new BigNumber(407)), function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)))).length), new BigNumber(-504))).Elements) {
          let _39_v0 = _compr_6;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)))).length), new BigNumber(-504)), _39_v0)) {
            _coll6.add(_module.__default.safeDivisionInt(_39_v0, new BigNumber(968)));
          }
        }
        return _coll6;
      }())).Intersect(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-84)), function (_40_i0) {
        return new BigNumber((_dafny.Set.fromElements(false, true)).length);
      })).length), new BigNumber((function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-491), new BigNumber(-155))) {
          let _41_v1 = _compr_7;
          if (((new BigNumber(-491)).isLessThanOrEqualTo(_41_v1)) && ((_41_v1).isLessThan(new BigNumber(-155)))) {
            _coll7.push([_module.__default.safeModuloInt(_41_v1, new BigNumber(778)),_dafny.Seq.of(true, true)]);
          }
        }
        return _coll7;
      }()).length))))).Difference(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of (_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(731)))).Elements) {
          let _42_v2 = _compr_8;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(731))), _42_v2)) {
            _coll8.add(_42_v2);
          }
        }
        return _coll8;
      }());
    };
    static fm35(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-469))), new BigNumber(655)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("emsbo")).length), new BigNumber(330), new BigNumber(554), new BigNumber((_dafny.Seq.of(true)).length))), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(132),true)).length), new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()))).cardinality()))));
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(41), new BigNumber((_dafny.Seq.UnicodeFromString("eqwwqd")).length), new BigNumber(79), new BigNumber(710), new BigNumber(416))).cardinality()), (new BigNumber((_dafny.Set.fromElements(false, false, true)).length)).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.Seq.of(false, true, !(!(false)), false, !(!(false)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length)), _module.__default.safeModuloInt(new BigNumber(679), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), function (_43_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })).length)));
    };
    static fm37(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(875));
    };
    static fm38(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("xeomwdehc")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-764)), function (_44_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), function (_45_i1) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        });
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), function (_46_i2) {
        return _dafny.Seq.UnicodeFromString("unkma");
      })));
    };
    static fm39(p0, p1, p2, p3, globalState) {
      if ((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(935), new BigNumber(433))) {
            let _47_v0 = _compr_10;
            if (((new BigNumber(935)).isLessThanOrEqualTo(_47_v0)) && ((_47_v0).isLessThan(new BigNumber(433)))) {
              _coll10.push([_module.__default.safeDivisionInt(_47_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(_dafny.ONE), function (_48_i0) {
                return new _dafny.CodePoint('l'.codePointAt(0));
              }),true)).length)),true]);
            }
          }
          return _coll10;
        }()).length),false)).Keys.Elements) {
          let _49_v1 = _compr_9;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of _dafny.IntegerRange(new BigNumber(935), new BigNumber(433))) {
              let _47_v0 = _compr_11;
              if (((new BigNumber(935)).isLessThanOrEqualTo(_47_v0)) && ((_47_v0).isLessThan(new BigNumber(433)))) {
                _coll11.push([_module.__default.safeDivisionInt(_47_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(_dafny.ONE), function (_48_i0) {
                  return new _dafny.CodePoint('l'.codePointAt(0));
                }),true)).length)),true]);
              }
            }
            return _coll11;
          }()).length),false)).contains(_49_v1)) {
            _coll9.add((_49_v1).multipliedBy(new BigNumber(979)));
          }
        }
        return _coll9;
      }()).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(737)))) {
        return _module.D10.create_DC24(new _dafny.CodePoint('o'.codePointAt(0)));
      } else {
        return _module.D10.create_DC24(new _dafny.CodePoint('l'.codePointAt(0)));
      }
    };
    static fm40(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(true)))).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(691),!(!(true))));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = [];
      r1 = false;
      let _50_v0;
      _50_v0 = false;
      let _51_v1;
      _51_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _52_v2;
      _52_v2 = _module.D8.create_DC20(_50_v0, _module.__default.fm0(_50_v0, (_51_v1).update(p1, p1), p1, globalState), _50_v0, (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_50_v0, _51_v1, new BigNumber((_dafny.Seq.UnicodeFromString("t")).length), globalState),new BigNumber(515))).contains(false));
      let _53_v3;
      let _nw0 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _53_v3 = _nw0;
      let _54_v4;
      _54_v4 = _dafny.MultiSet.fromElements(_53_v3);
      _52_v2 = _module.D8.create_DC20(_50_v0, (new BigNumber(-148)).isLessThanOrEqualTo(p1), true, (_54_v4).IsSubsetOf(_54_v4));
      let _55_v5;
      _55_v5 = new _dafny.CodePoint('d'.codePointAt(0));
      let _56_v6;
      let _nw1 = new _module.C1();
      _nw1.__ctor(_55_v5, p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_50_v0,_50_v0)).length));
      _56_v6 = _nw1;
      let _57_v7;
      _57_v7 = _module.D8.create_DC18(_56_v6);
      let _pat_let_tv0 = _56_v6;
      let _pat_let_tv1 = _56_v6;
      let _source3 = function (_pat_let0_0) {
        return function (_58_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_59_dt__update_hcf33_h0) {
              return _module.D8.create_DC18(_59_dt__update_hcf33_h0);
            }(_pat_let1_0);
          }(((false) ? (_pat_let_tv0) : (_pat_let_tv1)));
        }(_pat_let0_0);
      }(_57_v7);
      if (_source3.is_DC19) {
        let _60___mcc_h0 = (_source3).cf34;
        let _61___mcc_h1 = (_source3).cf35;
        let _62___mcc_h2 = (_source3).cf36;
        let _63___mcc_h3 = (_source3).cf37;
        let _64___mcc_h4 = (_source3).cf38;
        let _65_cf38 = _64___mcc_h4;
        let _66_cf37 = _63___mcc_h3;
        let _67_cf36 = _62___mcc_h2;
        let _68_cf35 = _61___mcc_h1;
        let _69_cf34 = _60___mcc_h0;
        let _70_v8;
        let _nw2 = Array((new BigNumber(22)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _69_cf34;
        _nw2[(_dafny.ONE).toNumber()] = _68_cf35;
        _nw2[(new BigNumber(2)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(3)).toNumber()] = _68_cf35;
        _nw2[(new BigNumber(4)).toNumber()] = _68_cf35;
        _nw2[(new BigNumber(5)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(6)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(7)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(8)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(9)).toNumber()] = _68_cf35;
        _nw2[(new BigNumber(10)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(11)).toNumber()] = false;
        _nw2[(new BigNumber(12)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(13)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(14)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(15)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(16)).toNumber()] = _68_cf35;
        _nw2[(new BigNumber(17)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(18)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(19)).toNumber()] = _69_cf34;
        _nw2[(new BigNumber(20)).toNumber()] = _50_v0;
        _nw2[(new BigNumber(21)).toNumber()] = false;
        _70_v8 = _nw2;
        let _71_v9;
        _71_v9 = _dafny.Map.Empty.slice().updateUnsafe(_69_cf34,_66_cf37);
        let _72_v10;
        _72_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_71_v9).length),_50_v0);
        let _73_v11;
        _73_v11 = _dafny.Map.Empty.slice().updateUnsafe(_70_v8,_72_v10);
        let _74_v12;
        let _nw3 = new _module.C4();
        _nw3.__ctor(_69_cf34, (_73_v11).Merge(_dafny.Map.Empty.slice().updateUnsafe(_70_v8,_72_v10)), (p1).multipliedBy(_56_v6.f7), (_dafny.ZERO).minus((_56_v6).f6));
        _74_v12 = _nw3;
        (globalState).f1 = (p1).minus((_56_v6).f6);
        (_56_v6).f7 = (_56_v6).f6;
        if (_50_v0) {
          let _75_v13;
          _75_v13 = _dafny.Seq.of(_66_cf37);
          _71_v9 = (_71_v9).update(_dafny.Seq.IsProperPrefixOf(_75_v13, _dafny.Seq.of(new BigNumber(93))), _65_cf38);
          _65_cf38 = ((_56_v6).f6).multipliedBy(_56_v6.f7);
          let _76_v14;
          _76_v14 = _dafny.MultiSet.fromElements((_56_v6).f6, (_56_v6).f6);
          let _77_v15;
          _77_v15 = _dafny.Seq.UnicodeFromString("ehcuiqm");
          let _78_v16;
          let _nw4 = new _module.C3();
          _nw4.__ctor(_56_v6.f7, (new BigNumber((_76_v14).cardinality())).minus(new BigNumber((_77_v15).length)));
          _78_v16 = _nw4;
          let _79_v17;
          _79_v17 = _dafny.Seq.of(((_50_v0) ? (_56_v6) : (_56_v6)));
          let _rhs0 = _78_v16;
          let _rhs1 = _55_v5;
          let _rhs2 = _79_v17;
          let _rhs3 = _56_v6.f7;
          let _lhs0 = _56_v6;
          _78_v16 = _rhs0;
          _55_v5 = _rhs1;
          _79_v17 = _rhs2;
          _lhs0.f7 = _rhs3;
          let _80_v18;
          let _nw5 = new _module.C0();
          _nw5.__ctor((_56_v6.f7).isLessThan((_56_v6).f6));
          _80_v18 = _nw5;
          let _rhs4 = _77_v15;
          let _rhs5 = new BigNumber(591);
          let _rhs6 = _80_v18;
          let _rhs7 = (_56_v6).f6;
          let _rhs8 = _56_v6;
          let _lhs1 = _56_v6;
          let _lhs2 = globalState;
          _77_v15 = _rhs4;
          _lhs1.f7 = _rhs5;
          _80_v18 = _rhs6;
          _lhs2.f1 = _rhs7;
          _56_v6 = _rhs8;
          _55_v5 = _55_v5;
        } else {
          let _81_v19;
          let _nw6 = Array((new BigNumber(27)).toNumber()).fill([]);
          _81_v19 = _nw6;
          _81_v19 = _81_v19;
          let _82_v20;
          _82_v20 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_71_v9);
          _82_v20 = (_82_v20).update(_55_v5, _71_v9);
          (globalState).f1 = _56_v6.f7;
          let _83_v21;
          _83_v21 = _module.D13.create_DC32(_51_v1);
          let _84_v22;
          _84_v22 = _dafny.Set.fromElements(false, false);
          let _85_v23;
          _85_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_74_v12).f11, (_83_v21).dtor_cf61, new BigNumber(764), globalState),_84_v22);
          let _86_v24;
          _86_v24 = _dafny.Map.Empty.slice().updateUnsafe((_74_v12).f11,(_74_v12).f11);
          _85_v23 = (_85_v23).update((((_86_v24).contains((((_72_v10).contains(_56_v6.f7)) ? ((_72_v10).get(_56_v6.f7)) : (_69_cf34)))) ? ((_86_v24).get((((_72_v10).contains(_56_v6.f7)) ? ((_72_v10).get(_56_v6.f7)) : (_69_cf34)))) : (_50_v0)), _84_v22);
          _68_cf35 = _69_cf34;
        }
      } else if (_source3.is_DC20) {
        let _87___mcc_h5 = (_source3).cf39;
        let _88___mcc_h6 = (_source3).cf40;
        let _89___mcc_h7 = (_source3).cf41;
        let _90___mcc_h8 = (_source3).cf42;
        let _91_cf42 = _90___mcc_h8;
        let _92_cf41 = _89___mcc_h7;
        let _93_cf40 = _88___mcc_h6;
        let _94_cf39 = _87___mcc_h5;
        let _95_v25;
        let _nw7 = Array((new BigNumber(19)).toNumber());
        _95_v25 = _nw7;
        let _96_v26;
        let _nw8 = new _module.C0();
        _nw8.__ctor(_94_cf39);
        _96_v26 = _nw8;
        let _index0 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_95_v25).length));
        (_95_v25)[_index0] = _96_v26;
        let _index1 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_95_v25).length));
        (_95_v25)[_index1] = _96_v26;
        if (((_56_v6).f6).isLessThanOrEqualTo((_dafny.ZERO).minus(_56_v6.f7))) {
          let _97_v27;
          _97_v27 = _dafny.Seq.UnicodeFromString("rcspe");
          (_96_v26).f10 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("bjfovcio"), _module.__default.safeIndex(_56_v6.f7, new BigNumber((_dafny.Seq.UnicodeFromString("bjfovcio")).length)), _55_v5), _dafny.Seq.Concat(_97_v27, _dafny.Seq.UnicodeFromString("xmouvrgh")));
          (_56_v6).f7 = (_56_v6.f7).plus((_56_v6).f6);
          let _98_v28;
          _98_v28 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), ((_99_v5) => function (_100_i0) {
            return _99_v5;
          })(_55_v5)));
          let _101_v29;
          _101_v29 = _dafny.Seq.of(_91_cf42, _92_cf41);
          let _102_v30;
          _102_v30 = _dafny.Seq.of((_56_v6).f6);
          let _rhs9 = _96_v26.f10;
          let _rhs10 = ((_56_v6).f6).multipliedBy(new BigNumber((_dafny.Seq.Concat(_module.__default.fm35(new BigNumber((_101_v29).length), (_56_v6).f6, globalState), _102_v30)).length));
          let _rhs11 = _98_v28;
          let _rhs12 = (_56_v6.f7).plus(((_56_v6).f6).plus(_56_v6.f7));
          let _lhs3 = globalState;
          _93_cf40 = _rhs9;
          _lhs3.f1 = _rhs10;
          _98_v28 = _rhs11;
          r0 = _rhs12;
          let _103_v31;
          let _init0 = ((_104_v6) => function (_105_i1) {
            return _dafny.Set.fromElements(_104_v6.f7, (_104_v6).f6);
          })(_56_v6);
          let _nw9 = Array((new BigNumber(2)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw9.length); _i0_0++) {
            _nw9[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _103_v31 = _nw9;
          let _106_v32;
          _106_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D11.create_DC28(_97_v27, _103_v31, _93_cf40, !(_50_v0))).dtor_cf54);
          _106_v32 = _106_v32;
          let _107_v33;
          _107_v33 = _dafny.Seq.of(_53_v3, _53_v3, _53_v3);
          _53_v3 = (_107_v33)[_module.__default.safeIndex(_56_v6.f7, new BigNumber((_107_v33).length))];
        } else {
          let _108_v34;
          _108_v34 = _dafny.Seq.UnicodeFromString("tobbymm");
          let _109_v35;
          _109_v35 = _dafny.Seq.of(new BigNumber((_108_v34).length));
          let _110_v36;
          _110_v36 = _dafny.Seq.of(_91_cf42, _50_v0, _50_v0, (new BigNumber((_dafny.Seq.of(_94_cf39)).length)).isLessThanOrEqualTo((_109_v35)[_module.__default.safeIndex(_56_v6.f7, new BigNumber((_109_v35).length))]), !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("wlngpj"), _55_v5));
          let _rhs13 = (_57_v7).dtor_cf33;
          let _rhs14 = !((_110_v36)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_110_v36).length))]);
          let _rhs15 = _108_v34;
          let _lhs4 = _96_v26;
          _56_v6 = _rhs13;
          _lhs4.f10 = _rhs14;
          _108_v34 = _rhs15;
          let _111_v37;
          let _nw10 = Array((new BigNumber(4)).toNumber()).fill(false);
          _111_v37 = _nw10;
          let _112_v38;
          _112_v38 = _dafny.Map.Empty.slice().updateUnsafe(_56_v6,_96_v26.f10);
          let _index2 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_111_v37).length));
          (_111_v37)[_index2] = _module.__default.fm0(_50_v0, (_51_v1).update(new BigNumber((_112_v38).length), _56_v6.f7), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_92_cf41)).length)), globalState);
          let _index3 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_111_v37).length));
          (_111_v37)[_index3] = _93_cf40;
          _108_v34 = _108_v34;
          let _113_v39;
          _113_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,_91_cf42);
          let _114_v40;
          _114_v40 = _dafny.Map.Empty.slice().updateUnsafe(_111_v37,_113_v39);
          let _115_v41;
          let _nw11 = new _module.C4();
          _nw11.__ctor(_96_v26.f10, _114_v40, _module.__default.safeDivisionInt(new BigNumber((_108_v34).length), (_dafny.ZERO).minus(_56_v6.f7)), _module.__default.fm1(_91_cf42, globalState));
          _115_v41 = _nw11;
          _115_v41 = _115_v41;
          (_56_v6).f7 = _56_v6.f7;
        }
        let _116_v42;
        _116_v42 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_93_cf40)),_53_v3);
        let _117_v43;
        _117_v43 = _module.D1.create_DC3(_93_cf40, p1, new BigNumber((_116_v42).length));
        let _118_v44;
        _118_v44 = _dafny.Map.Empty.slice().updateUnsafe(_92_cf41,(_56_v6).f6);
        let _119_v45;
        _119_v45 = _dafny.Map.Empty.slice().updateUnsafe(_96_v26.f10,_91_cf42);
        let _120_v46;
        _120_v46 = _dafny.Map.Empty.slice().updateUnsafe(_117_v43,(((_118_v44).contains((((_119_v45).contains(_50_v0)) ? ((_119_v45).get(_50_v0)) : (_92_cf41)))) ? ((_118_v44).get((((_119_v45).contains(_50_v0)) ? ((_119_v45).get(_50_v0)) : (_92_cf41)))) : (_56_v6.f7)));
        _120_v46 = (_120_v46).update(_117_v43, (_56_v6.f7).multipliedBy((_56_v6).f6));
        let _121_v47;
        _121_v47 = _dafny.Map.Empty.slice().updateUnsafe(_56_v6,_53_v3);
        let _122_v48;
        let _nw12 = Array((new BigNumber(26)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _53_v3;
        _nw12[(_dafny.ONE).toNumber()] = _53_v3;
        _nw12[(new BigNumber(2)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(3)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(4)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(5)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(6)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(7)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(8)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(9)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(10)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(11)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(12)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(13)).toNumber()] = (((_121_v47).contains(_56_v6)) ? ((_121_v47).get(_56_v6)) : (_53_v3));
        _nw12[(new BigNumber(14)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(15)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(16)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(17)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(18)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(19)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(20)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(21)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(22)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(23)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(24)).toNumber()] = _53_v3;
        _nw12[(new BigNumber(25)).toNumber()] = _53_v3;
        _122_v48 = _nw12;
        let _123_v49;
        _123_v49 = _dafny.Seq.of(p1, p1, p1);
        let _124_v50;
        let _nw13 = new _module.C5();
        _nw13.__ctor(_122_v48, _123_v49, (_56_v6).f6, (_56_v6).f6);
        _124_v50 = _nw13;
      } else if (_source3.is_DC21) {
        let _125___mcc_h9 = (_source3).cf43;
        let _126___mcc_h10 = (_source3).cf44;
        let _127_cf44 = _126___mcc_h10;
        let _128_cf43 = _125___mcc_h9;
        let _129_v51;
        _129_v51 = _dafny.MultiSet.fromElements(_50_v0);
        r1 = ((_129_v51).Intersect(_129_v51)).IsProperSubsetOf(_129_v51);
        if (_50_v0) {
          let _130_v52;
          let _nw14 = new _module.C2();
          _nw14.__ctor((_56_v6).f6, (_56_v6).f6);
          _130_v52 = _nw14;
          let _131_v53;
          _131_v53 = _dafny.Seq.of(p1);
          let _132_v54;
          _132_v54 = _module.D9.create_DC22(_131_v53);
          _132_v54 = _132_v54;
          let _133_v55;
          _133_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_50_v0, _50_v0, _50_v0),(_56_v6).f6);
          let _134_v56;
          _134_v56 = _dafny.Set.fromElements(_50_v0);
          _133_v55 = (_133_v55).update(_134_v56, (_56_v6).f6);
          let _135_v57;
          let _nw15 = new _module.C3();
          _nw15.__ctor(new BigNumber((_dafny.Seq.of(_56_v6.f7)).length), (_dafny.ZERO).minus((_56_v6).f6));
          _135_v57 = _nw15;
          let _136_v58;
          _136_v58 = _dafny.Seq.of(_135_v57, _135_v57, _135_v57);
          let _137_v59;
          _137_v59 = _dafny.Seq.UnicodeFromString("fxyo");
          let _index4 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_53_v3).length));
          (_53_v3)[_index4] = new BigNumber((_137_v59).length);
          let _138_v60;
          let _nw16 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _138_v60 = _nw16;
          let _index5 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_53_v3).length));
          let _rhs16 = _136_v58;
          let _rhs17 = _module.__default.safeDivisionInt((_56_v6).f6, _56_v6.f7);
          let _rhs18 = (((_56_v6.f7).isLessThanOrEqualTo(_56_v6.f7)) ? (_138_v60) : (_138_v60));
          let _lhs5 = _53_v3;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_53_v3).length));
          _136_v58 = _rhs16;
          _lhs5[_lhs6] = _rhs17;
          _138_v60 = _rhs18;
          let _139_v61;
          _139_v61 = _dafny.Map.Empty.slice().updateUnsafe(_50_v0,_50_v0);
          let _140_v62;
          _140_v62 = _module.D0.create_DC0(_134_v56);
          let _141_v63;
          _141_v63 = _dafny.Map.Empty.slice().updateUnsafe(_140_v62,false);
          let _142_v64;
          _142_v64 = _module.D1.create_DC2(_141_v63);
          (_130_v52).m8(_139_v61, _module.D2.create_DC6(_142_v64, _137_v59, new BigNumber(-757)), _137_v59, globalState);
        } else {
          let _143_v65;
          let _nw17 = Array((_dafny.ONE).toNumber());
          _143_v65 = _nw17;
          let _nw18 = Array((new BigNumber(9)).toNumber());
          _143_v65 = _nw18;
          let _144_v66;
          _144_v66 = _dafny.Seq.of((_56_v6).f6);
          let _145_v67;
          _145_v67 = _dafny.Map.Empty.slice().updateUnsafe(_51_v1,_144_v66);
          _144_v66 = (((_145_v67).contains((_51_v1).Merge(_51_v1))) ? ((_145_v67).get((_51_v1).Merge(_51_v1))) : (_144_v66));
          let _146_v68;
          _146_v68 = _dafny.Seq.of(_50_v0, _50_v0, _50_v0);
          (_56_v6).f7 = new BigNumber((_module.__default.fm2(_144_v66, _56_v6.f7, _56_v6.f7, (_dafny.MultiSet.FromArray(_146_v68)).IsProperSubsetOf(_129_v51), globalState)).length);
          _50_v0 = false;
          r1 = _50_v0;
        }
        let _index6 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_53_v3).length));
        (_53_v3)[_index6] = (_56_v6).f6;
        let _147_v69;
        _147_v69 = _dafny.Seq.UnicodeFromString("fummh");
        let _148_v70;
        _148_v70 = _dafny.Seq.of((_56_v6).f6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_56_v6.f7,_50_v0)).length));
        let _149_v71;
        _149_v71 = _dafny.Seq.of(new BigNumber((_148_v70).length));
        let _index7 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_53_v3).length));
        let _rhs19 = (_56_v6).f6;
        let _rhs20 = _dafny.Seq.UnicodeFromString("pba");
        let _rhs21 = _53_v3;
        let _rhs22 = _148_v70;
        let _rhs23 = _50_v0;
        let _lhs7 = _53_v3;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_53_v3).length));
        _lhs7[_lhs8] = _rhs19;
        _147_v69 = _rhs20;
        _53_v3 = _rhs21;
        _149_v71 = _rhs22;
        r1 = _rhs23;
        let _150_v72;
        let _init1 = ((_151_v0) => function (_152_i2) {
          return _151_v0;
        })(_50_v0);
        let _nw19 = Array((new BigNumber(18)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw19.length); _i0_1++) {
          _nw19[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _150_v72 = _nw19;
        let _index8 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_150_v72).length));
        (_150_v72)[_index8] = !(_50_v0);
        let _index9 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_150_v72).length));
        (_150_v72)[_index9] = _50_v0;
      } else {
        let _153___mcc_h11 = (_source3).cf33;
        let _154_cf33 = _153___mcc_h11;
        let _155_v73;
        _155_v73 = _module.D10.create_DC24(_55_v5);
        let _pat_let_tv2 = _55_v5;
        let _156_v74;
        _156_v74 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let2_0) {
          return function (_157_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_158_dt__update_hcf51_h0) {
                return _module.D10.create_DC24(_158_dt__update_hcf51_h0);
              }(_pat_let3_0);
            }(_pat_let_tv2);
          }(_pat_let2_0);
        }(_155_v73)).dtor_cf51,_154_cf33);
        let _159_v75;
        _159_v75 = _module.D8.create_DC19(true, !(_156_v74).contains(_55_v5), _dafny.Set.fromElements(_53_v3, _53_v3), (_154_cf33.f7).minus(p1), _module.__default.safeDivisionInt(_56_v6.f7, (_56_v6).f6));
        let _source4 = _159_v75;
        if (_source4.is_DC19) {
          let _160___mcc_h12 = (_source4).cf34;
          let _161___mcc_h13 = (_source4).cf35;
          let _162___mcc_h14 = (_source4).cf36;
          let _163___mcc_h15 = (_source4).cf37;
          let _164___mcc_h16 = (_source4).cf38;
          let _165_cf38 = _164___mcc_h16;
          let _166_cf37 = _163___mcc_h15;
          let _167_cf36 = _162___mcc_h14;
          let _168_cf35 = _161___mcc_h13;
          let _169_cf34 = _160___mcc_h12;
          let _170_v76;
          _170_v76 = _dafny.Seq.UnicodeFromString("nvmlujfv");
          let _171_v77;
          _171_v77 = _dafny.MultiSet.fromElements(_170_v76);
          _171_v77 = _171_v77;
          let _172_v78;
          let _init2 = ((_173_v0) => function (_174_i3) {
            return _173_v0;
          })(_50_v0);
          let _nw20 = Array((new BigNumber(18)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
            _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _172_v78 = _nw20;
          let _175_v79;
          _175_v79 = _dafny.Map.Empty.slice().updateUnsafe((_56_v6).f6,!(_50_v0));
          let _176_v80;
          let _nw21 = Array((new BigNumber(13)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _172_v78;
          _nw21[(_dafny.ONE).toNumber()] = _172_v78;
          _nw21[(new BigNumber(2)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(3)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(4)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(5)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(6)).toNumber()] = (((((_175_v79).contains((_154_cf33).f6)) ? ((_175_v79).get((_154_cf33).f6)) : (_169_cf34))) ? (_172_v78) : (_172_v78));
          _nw21[(new BigNumber(7)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(8)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(9)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(10)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(11)).toNumber()] = _172_v78;
          _nw21[(new BigNumber(12)).toNumber()] = _172_v78;
          _176_v80 = _nw21;
          let _index10 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_176_v80).length));
          (_176_v80)[_index10] = _172_v78;
          let _index11 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_176_v80).length));
          (_176_v80)[_index11] = _172_v78;
          let _index12 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_53_v3).length));
          (_53_v3)[_index12] = _module.__default.safeModuloInt(_56_v6.f7, p1);
          let _index13 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_53_v3).length));
          (_53_v3)[_index13] = _56_v6.f7;
          let _177_v81;
          _177_v81 = _dafny.Map.Empty.slice().updateUnsafe(_55_v5,(new BigNumber(113)).minus(new BigNumber((_170_v76).length)));
          (_154_cf33).f7 = new BigNumber((_177_v81).length);
        } else if (_source4.is_DC20) {
          let _178___mcc_h17 = (_source4).cf39;
          let _179___mcc_h18 = (_source4).cf40;
          let _180___mcc_h19 = (_source4).cf41;
          let _181___mcc_h20 = (_source4).cf42;
          let _182_cf42 = _181___mcc_h20;
          let _183_cf41 = _180___mcc_h19;
          let _184_cf40 = _179___mcc_h18;
          let _185_cf39 = _178___mcc_h17;
          let _186_v82;
          _186_v82 = _dafny.Map.Empty.slice().updateUnsafe(_182_cf42,(_56_v6).f6);
          let _187_v83;
          _187_v83 = _dafny.Map.Empty.slice().updateUnsafe(_154_cf33.f7,_185_cf39);
          _186_v82 = (_186_v82).update(((_154_cf33).f6).isLessThan(_56_v6.f7), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of((((_187_v83).contains(new BigNumber(224))) ? ((_187_v83).get(new BigNumber(224))) : (false)), _182_cf42, true)).length), (_154_cf33).f6));
          _51_v1 = (_51_v1).Merge(_dafny.Map.Empty.slice().updateUnsafe((_56_v6).f6,_56_v6.f7));
          (_56_v6).f7 = (_56_v6).f6;
          let _188_v84;
          _188_v84 = _dafny.Seq.UnicodeFromString("jiffwqiab");
          let _189_v85;
          _189_v85 = _dafny.MultiSet.fromElements(new BigNumber(((_51_v1).update((_154_cf33).f6, (_56_v6).f6)).length), new BigNumber((_186_v82).length), (_56_v6).f6);
          let _190_v86;
          _190_v86 = _dafny.Map.Empty.slice().updateUnsafe(_189_v85,_188_v84);
          let _191_v87;
          let _nw22 = Array((new BigNumber(22)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _188_v84;
          _nw22[(_dafny.ONE).toNumber()] = _188_v84;
          _nw22[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), ((_192_v5) => function (_193_i4) {
            return _192_v5;
          })(_55_v5));
          _nw22[(new BigNumber(3)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("stfj");
          _nw22[(new BigNumber(5)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(6)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(7)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("iausxosc");
          _nw22[(new BigNumber(9)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(10)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(11)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(12)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(13)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(14)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("jjnked");
          _nw22[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), function (_194_i5) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }), _module.__default.safeIndex(_56_v6.f7, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), function (_195_i5) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length)), _55_v5);
          _nw22[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("laj");
          _nw22[(new BigNumber(18)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(19)).toNumber()] = _188_v84;
          _nw22[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("xt");
          _nw22[(new BigNumber(21)).toNumber()] = (((_190_v86).contains(_189_v85)) ? ((_190_v86).get(_189_v85)) : (_188_v84));
          _191_v87 = _nw22;
          let _196_v88;
          _196_v88 = _module.D13.create_DC33(_183_cf41, (_154_cf33).f6);
          let _197_v89;
          let _nw23 = Array((new BigNumber(22)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = _191_v87;
          _nw23[(_dafny.ONE).toNumber()] = _191_v87;
          _nw23[(new BigNumber(2)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(3)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(4)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(5)).toNumber()] = ((_182_cf42) ? (_191_v87) : (_191_v87));
          _nw23[(new BigNumber(6)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(7)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(8)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(9)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(10)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(11)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(12)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(13)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(14)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(15)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(16)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(17)).toNumber()] = (((_196_v88).dtor_cf62) ? (_191_v87) : (_191_v87));
          _nw23[(new BigNumber(18)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(19)).toNumber()] = _191_v87;
          _nw23[(new BigNumber(20)).toNumber()] = ((_184_cf40) ? (_191_v87) : (_191_v87));
          _nw23[(new BigNumber(21)).toNumber()] = _191_v87;
          _197_v89 = _nw23;
          let _index14 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_197_v89).length));
          (_197_v89)[_index14] = _191_v87;
          let _index15 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_197_v89).length));
          (_197_v89)[_index15] = _191_v87;
        } else if (_source4.is_DC21) {
          let _198___mcc_h21 = (_source4).cf43;
          let _199___mcc_h22 = (_source4).cf44;
          let _200_cf44 = _199___mcc_h22;
          let _201_cf43 = _198___mcc_h21;
          _55_v5 = _55_v5;
          let _202_v90;
          _202_v90 = _module.D10.create_DC25();
          _202_v90 = ((_50_v0) ? (_202_v90) : (_module.D10.create_DC25()));
          let _203_v91;
          let _init3 = ((_204_v5) => function (_205_i6) {
            return _204_v5;
          })(_55_v5);
          let _nw24 = Array((new BigNumber(23)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw24.length); _i0_3++) {
            _nw24[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _203_v91 = _nw24;
          let _index16 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_203_v91).length));
          (_203_v91)[_index16] = ((_50_v0) ? (_55_v5) : (_55_v5));
          let _index17 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_203_v91).length));
          (_203_v91)[_index17] = _55_v5;
          r1 = (_50_v0) || (false);
        } else {
          let _206___mcc_h23 = (_source4).cf33;
          let _207_cf33 = _206___mcc_h23;
          let _208_v92;
          _208_v92 = _module.D10.create_DC26(_50_v0);
          _208_v92 = _module.D10.create_DC26(_50_v0);
          let _209_v93;
          _209_v93 = _dafny.Seq.UnicodeFromString("bpmp");
          let _210_v94;
          _210_v94 = _dafny.Set.fromElements((_56_v6).f6, (_207_cf33).f6, (_56_v6).f6, (_dafny.ZERO).minus(_56_v6.f7));
          let _211_v95;
          let _nw25 = Array((new BigNumber(4)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = _210_v94;
          _nw25[(_dafny.ONE).toNumber()] = _210_v94;
          _nw25[(new BigNumber(2)).toNumber()] = _210_v94;
          _nw25[(new BigNumber(3)).toNumber()] = _210_v94;
          _211_v95 = _nw25;
          let _212_v96;
          _212_v96 = _module.D11.create_DC28(_209_v93, _211_v95, _50_v0, !(_50_v0));
          let _213_v97;
          _213_v97 = _dafny.Set.fromElements(_50_v0, (_212_v96).dtor_cf56);
          _213_v97 = _dafny.Set.fromElements(_50_v0);
          let _214_v98;
          let _nw26 = Array((new BigNumber(29)).toNumber()).fill([]);
          _214_v98 = _nw26;
          let _index18 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_214_v98).length));
          (_214_v98)[_index18] = _53_v3;
          let _index19 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_214_v98).length));
          let _rhs24 = _53_v3;
          let _rhs25 = _209_v93;
          let _rhs26 = (!(_50_v0) || (_50_v0)) && (!(_50_v0) || (_50_v0));
          let _lhs9 = _214_v98;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_214_v98).length));
          _lhs9[_lhs10] = _rhs24;
          _209_v93 = _rhs25;
          _50_v0 = _rhs26;
          let _215_v99;
          _215_v99 = _dafny.Seq.of((_dafny.ZERO).minus((_207_cf33).f6));
          let _216_v100;
          let _nw27 = new _module.C5();
          _nw27.__ctor(_214_v98, _215_v99, (_56_v6).f6, (_56_v6).f6);
          _216_v100 = _nw27;
          _216_v100 = _216_v100;
        }
        let _217_v101;
        let _nw28 = new _module.C3();
        _nw28.__ctor(p1, (_154_cf33).f6);
        _217_v101 = _nw28;
        let _index20 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_53_v3).length));
        (_53_v3)[_index20] = ((false) ? (_56_v6.f7) : ((_56_v6).f6));
        let _218_v102;
        let _nw29 = Array((new BigNumber(13)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = _50_v0;
        _nw29[(_dafny.ONE).toNumber()] = _50_v0;
        _nw29[(new BigNumber(2)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(3)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(4)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(5)).toNumber()] = (_159_v75).dtor_cf35;
        _nw29[(new BigNumber(6)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(7)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(8)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(9)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(10)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(11)).toNumber()] = _50_v0;
        _nw29[(new BigNumber(12)).toNumber()] = !((_217_v101).fm16(globalState));
        _218_v102 = _nw29;
        let _219_v103;
        _219_v103 = _dafny.Set.fromElements(_dafny.Seq.of(_217_v101));
        let _220_v104;
        _220_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_219_v103).length),_50_v0);
        let _221_v105;
        _221_v105 = _dafny.Map.Empty.slice().updateUnsafe(_218_v102,_220_v104);
        let _222_v106;
        let _nw30 = new _module.C4();
        _nw30.__ctor(_50_v0, _221_v105, _module.__default.fm1(_50_v0, globalState), (_56_v6).f6);
        _222_v106 = _nw30;
        let _223_v107;
        _223_v107 = _dafny.Map.Empty.slice().updateUnsafe(_222_v106,true);
        let _224_v108;
        _224_v108 = _dafny.Seq.of(_56_v6);
        let _index21 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_53_v3).length));
        let _rhs27 = ((_50_v0) ? ((_56_v6.f7).isEqualTo(new BigNumber((_223_v107).length))) : ((_222_v106).f11));
        let _rhs28 = (_222_v106).f11;
        let _rhs29 = ((_154_cf33).f6).multipliedBy(_module.__default.safeModuloInt(new BigNumber((_224_v108).length), p1));
        let _rhs30 = _154_cf33.f7;
        let _lhs11 = _154_cf33;
        let _lhs12 = _53_v3;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_53_v3).length));
        _50_v0 = _rhs27;
        _50_v0 = _rhs28;
        _lhs11.f7 = _rhs29;
        _lhs12[_lhs13] = _rhs30;
        let _index22 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_218_v102).length));
        (_218_v102)[_index22] = (_222_v106).f11;
        let _225_v109;
        let _nw31 = Array((_dafny.ONE).toNumber());
        _225_v109 = _nw31;
        let _index23 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_225_v109).length));
        (_225_v109)[_index23] = _222_v106;
        let _index24 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_218_v102).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_225_v109).length));
        let _rhs31 = (_222_v106).f11;
        let _rhs32 = _222_v106;
        let _rhs33 = (_53_v3)[_module.__default.safeIndex(new BigNumber(771), new BigNumber((_53_v3).length))];
        let _rhs34 = _module.__default.safeModuloInt(p1, (_dafny.ZERO).minus((_56_v6).f6));
        let _lhs14 = _218_v102;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_218_v102).length));
        let _lhs16 = _225_v109;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_225_v109).length));
        let _lhs18 = _154_cf33;
        let _lhs19 = _56_v6;
        _lhs14[_lhs15] = _rhs31;
        _lhs16[_lhs17] = _rhs32;
        _lhs18.f7 = _rhs33;
        _lhs19.f7 = _rhs34;
      }
      let _hi0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, _56_v6.f7));
      for (let _226_i7 = p1; _226_i7.isLessThan(_hi0); _226_i7 = _226_i7.plus(_dafny.ONE)) {
        let _index26 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length));
        (_53_v3)[_index26] = _226_i7;
        let _227_v110;
        _227_v110 = _dafny.Seq.of((_56_v6).f6);
        let _index27 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length));
        (_53_v3)[_index27] = (p1).minus(new BigNumber((_227_v110).length));
        let _index28 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length));
        (_53_v3)[_index28] = _56_v6.f7;
        let _228_v111;
        let _nw32 = Array((new BigNumber(13)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = _56_v6;
        _nw32[(_dafny.ONE).toNumber()] = _56_v6;
        _nw32[(new BigNumber(2)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(3)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(4)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(5)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(6)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(7)).toNumber()] = ((_50_v0) ? (_56_v6) : (_56_v6));
        _nw32[(new BigNumber(8)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(9)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(10)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(11)).toNumber()] = _56_v6;
        _nw32[(new BigNumber(12)).toNumber()] = _56_v6;
        _228_v111 = _nw32;
        let _index29 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_228_v111).length));
        (_228_v111)[_index29] = _56_v6;
        let _229_v112;
        let _nw33 = Array((new BigNumber(28)).toNumber()).fill(false);
        _229_v112 = _nw33;
        let _230_v113;
        _230_v113 = _dafny.Map.Empty.slice().updateUnsafe(_229_v112,_dafny.Map.Empty.slice().updateUnsafe(_226_i7,!(!(_50_v0))));
        let _index30 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_228_v111).length));
        let _nw34 = new _module.C4();
        _nw34.__ctor(_50_v0, _230_v113, (p1).multipliedBy((_56_v6).f6), _module.__default.safeModuloInt((_56_v6).f6, _module.__default.fm1(_50_v0, globalState)));
        (_228_v111)[_index30] = _nw34;
        if (_50_v0) {
          let _231_v115;
          _231_v115 = _dafny.Set.fromElements((_56_v6).f6, _56_v6.f7);
          let _232_v116;
          _232_v116 = _dafny.Set.fromElements(!((function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of _dafny.IntegerRange(new BigNumber(798), new BigNumber(348))) {
              let _233_v114 = _compr_12;
              if (((new BigNumber(798)).isLessThanOrEqualTo(_233_v114)) && ((_233_v114).isLessThan(new BigNumber(348)))) {
                _coll12.add((_233_v114).plus((_53_v3)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length))]));
              }
            }
            return _coll12;
          }()).IsProperSubsetOf(_231_v115)));
          let _index31 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length));
          (_53_v3)[_index31] = new BigNumber((_232_v116).length);
          _50_v0 = (true) || (_50_v0);
          let _234_v117;
          _234_v117 = _dafny.Seq.UnicodeFromString("nbbdvvot");
          let _235_v118;
          let _nw35 = new _module.C0();
          _nw35.__ctor(_50_v0);
          _235_v118 = _nw35;
          let _236_v119;
          _236_v119 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_231_v115).length),_235_v118);
          let _237_v120;
          let _nw36 = new _module.C1();
          _nw36.__ctor(_module.__default.fm20(_dafny.Seq.Create(_module.__default.abs(new BigNumber(831)), ((_238_v5) => function (_239_i8) {
            return _238_v5;
          })(_55_v5)), false, _module.__default.fm1(_50_v0, globalState), globalState), (new BigNumber((_234_v117).length)).plus(new BigNumber((_236_v119).length)), (_56_v6).f6);
          _237_v120 = _nw36;
          let _240_v121;
          let _nw37 = Array((new BigNumber(4)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = (_237_v120).f13;
          _nw37[(_dafny.ONE).toNumber()] = ((_50_v0) ? (_55_v5) : (_55_v5));
          _nw37[(new BigNumber(2)).toNumber()] = (_237_v120).f13;
          _nw37[(new BigNumber(3)).toNumber()] = (_237_v120).f13;
          _240_v121 = _nw37;
          let _241_v122;
          let _init4 = ((_242_v6) => function (_243_i9) {
            return (_243_i9).multipliedBy((_242_v6).f6);
          })(_56_v6);
          let _nw38 = Array((new BigNumber(7)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw38.length); _i0_4++) {
            _nw38[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _241_v122 = _nw38;
          let _244_v123;
          _244_v123 = _dafny.Set.fromElements(_module.__default.fm14(new BigNumber(831), p1, _56_v6.f7, globalState), ((_235_v118.f10) ? (_55_v5) : (_55_v5)), (_237_v120).f13, (_237_v120).f13);
          let _245_v124;
          _245_v124 = _module.D2.create_DC5(_241_v122);
          let _rhs35 = _240_v121;
          let _rhs36 = _50_v0;
          let _rhs37 = (_245_v124).dtor_cf9;
          let _rhs38 = (_244_v123).Difference(_244_v123);
          _240_v121 = _rhs35;
          _50_v0 = _rhs36;
          _241_v122 = _rhs37;
          _244_v123 = _rhs38;
          let _246_v125;
          let _nw39 = new _module.C2();
          _nw39.__ctor(_56_v6.f7, new BigNumber((_234_v117).length));
          _246_v125 = _nw39;
        } else {
          let _247_v126;
          let _nw40 = new _module.C0();
          _nw40.__ctor(_50_v0);
          _247_v126 = _nw40;
          let _248_v127;
          let _nw41 = Array((new BigNumber(22)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _247_v126;
          _nw41[(_dafny.ONE).toNumber()] = _247_v126;
          _nw41[(new BigNumber(2)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(3)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(4)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(5)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(6)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(7)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(8)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(9)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(10)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(11)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(12)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(13)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(14)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(15)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(16)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(17)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(18)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(19)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(20)).toNumber()] = _247_v126;
          _nw41[(new BigNumber(21)).toNumber()] = _247_v126;
          _248_v127 = _nw41;
          let _249_v128;
          _249_v128 = _dafny.Seq.of(_247_v126, _247_v126);
          let _250_v129;
          _250_v129 = _dafny.Seq.UnicodeFromString("vrbqmsdid");
          let _251_v130;
          _251_v130 = (_249_v128)[_module.__default.safeIndex(new BigNumber((_250_v129).length), new BigNumber((_249_v128).length))];
          let _index32 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_248_v127).length));
          (_248_v127)[_index32] = (_251_v130);
          let _index33 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_248_v127).length));
          (_248_v127)[_index33] = _247_v126;
          r1 = !(_50_v0);
          let _252_v131;
          let _init5 = ((_253_v6) => function (_254_i10) {
            return _module.__default.safeDivisionInt(_254_i10, (_253_v6).f6);
          })(_56_v6);
          let _nw42 = Array((new BigNumber(23)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw42.length); _i0_5++) {
            _nw42[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _252_v131 = _nw42;
          let _255_v132;
          _255_v132 = _dafny.Set.fromElements(_252_v131);
          let _256_v133;
          _256_v133 = _module.D8.create_DC19(_50_v0, _50_v0, _255_v132, _56_v6.f7, new BigNumber(130));
          let _257_v134;
          _257_v134 = _dafny.Map.Empty.slice().updateUnsafe(_56_v6,!((_256_v133).dtor_cf34));
          let _258_v135;
          _258_v135 = _module.D1.create_DC3(_50_v0, _module.__default.fm1(_247_v126.f10, globalState), (_56_v6).f6);
          let _index34 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length));
          let _rhs39 = (_227_v110)[_module.__default.safeIndex((_53_v3)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length))], new BigNumber((_227_v110).length))];
          let _rhs40 = _247_v126.f10;
          let _rhs41 = (((((((_257_v134).contains(_56_v6)) ? ((_257_v134).get(_56_v6)) : (_50_v0))) ? (_50_v0) : (_50_v0))) ? (new BigNumber((_dafny.Seq.of((_258_v135).dtor_cf5)).length)) : ((_56_v6).f6));
          let _rhs42 = _module.__default.fm1(_247_v126.f10, globalState);
          let _lhs20 = _53_v3;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_53_v3).length));
          r0 = _rhs39;
          r1 = _rhs40;
          r0 = _rhs41;
          _lhs20[_lhs21] = _rhs42;
          let _nw43 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _252_v131 = _nw43;
          (globalState).f1 = _module.__default.safeModuloInt((new BigNumber(-750)).minus((_56_v6).f6), new BigNumber((((_50_v0) ? (_dafny.Seq.UnicodeFromString("wunqg")) : (_250_v129))).length));
        }
      }
      let _hi1 = _56_v6.f7;
      for (let _259_i11 = (_56_v6.f7).minus(_56_v6.f7); _259_i11.isLessThan(_hi1); _259_i11 = _259_i11.plus(_dafny.ONE)) {
        let _260_v136;
        let _nw44 = Array((new BigNumber(18)).toNumber()).fill([]);
        _260_v136 = _nw44;
        let _261_v137;
        _261_v137 = _dafny.Seq.of(_259_i11);
        let _262_v138;
        _262_v138 = _dafny.Map.Empty.slice().updateUnsafe(_56_v6.f7,_50_v0);
        let _263_v139;
        let _nw45 = new _module.C5();
        _nw45.__ctor(_260_v136, _261_v137, ((_50_v0) ? (p1) : (p1)), new BigNumber((_262_v138).length));
        _263_v139 = _nw45;
        _263_v139 = _263_v139;
        let _264_v140;
        let _nw46 = new _module.C0();
        _nw46.__ctor(_50_v0);
        _264_v140 = _nw46;
        let _265_v141;
        let _nw47 = Array((new BigNumber(2)).toNumber()).fill(false);
        _265_v141 = _nw47;
        (_263_v139).m3(_265_v141, (_56_v6.f7).minus(p1), globalState);
        let _266_v142;
        let _nw48 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Set.Empty);
        _266_v142 = _nw48;
        let _index35 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_266_v142).length));
        (_266_v142)[_index35] = (_dafny.Set.fromElements(_264_v140.f10, _50_v0)).Intersect(_dafny.Set.fromElements(false, (((_262_v138).contains(_259_i11)) ? ((_262_v138).get(_259_i11)) : (false))));
        let _267_v143;
        _267_v143 = _dafny.Set.fromElements(_50_v0, (_50_v0) || (true), (_264_v140.f10) && (_50_v0));
        let _index36 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_266_v142).length));
        (_266_v142)[_index36] = _267_v143;
      }
      let _268_v144;
      _268_v144 = _dafny.Seq.of(false, _50_v0, _50_v0);
      let _269_v145;
      _269_v145 = _dafny.Map.Empty.slice().updateUnsafe(_268_v144,!(_50_v0));
      let _270_v146;
      _270_v146 = _module.D5.create_DC11(_268_v144);
      _269_v145 = (_269_v145).update((_270_v146).dtor_cf20, _50_v0);
      let _271_v147;
      _271_v147 = _dafny.Set.fromElements(_56_v6.f7, _56_v6.f7);
      let _272_v148;
      _272_v148 = _dafny.Seq.of(new BigNumber((_271_v147).length), _56_v6.f7);
      r0 = _module.__default.safeDivisionInt(_56_v6.f7, ((_272_v148)[_module.__default.safeIndex((_56_v6).f6, new BigNumber((_272_v148).length))]).minus(p1));
      r1 = _50_v0;
      let _273_v149;
      let _nw49 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
      _273_v149 = _nw49;
      r2 = _273_v149;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _274_globalState;
      let _nw50 = new _module.GlobalState();
      _nw50.__ctor(true, new BigNumber(952), new BigNumber(374), new BigNumber(575), new BigNumber(555), false);
      _274_globalState = _nw50;
      let _275_v0;
      _275_v0 = new BigNumber(-379);
      (_274_globalState).f1 = _module.__default.safeDivisionInt(_275_v0, (_275_v0).minus(_275_v0));
      let _276_v1;
      _276_v1 = false;
      (_274_globalState).f1 = ((_276_v1) ? ((_275_v0).multipliedBy(_275_v0)) : (_275_v0));
      let _277_v2;
      let _init6 = function (_278_i0) {
        return !(true);
      };
      let _nw51 = Array((new BigNumber(17)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw51.length); _i0_6++) {
        _nw51[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _277_v2 = _nw51;
      let _index37 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
      (_277_v2)[_index37] = _276_v1;
      let _279_v3;
      _279_v3 = _dafny.Set.fromElements(_276_v1);
      let _280_v4;
      _280_v4 = _dafny.Map.Empty.slice().updateUnsafe((_279_v3).Union(_279_v3),_276_v1);
      let _281_v5;
      let _nw52 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _281_v5 = _nw52;
      let _282_v6;
      _282_v6 = _dafny.Set.fromElements(_281_v5, _281_v5);
      let _283_v7;
      _283_v7 = _dafny.Seq.UnicodeFromString("klhmhabdk");
      let _284_v8;
      _284_v8 = _dafny.MultiSet.fromElements(_283_v7);
      let _285_v9;
      _285_v9 = _dafny.Map.Empty.slice().updateUnsafe(_284_v8,_282_v6);
      let _index38 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
      let _rhs43 = _276_v1;
      let _rhs44 = _276_v1;
      let _rhs45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_276_v1, _276_v1),_276_v1);
      let _rhs46 = (_282_v6).IsDisjointFrom((((_285_v9).contains(_dafny.MultiSet.FromArray(_dafny.Seq.of(_283_v7, _283_v7, _283_v7, _283_v7)))) ? ((_285_v9).get(_dafny.MultiSet.FromArray(_dafny.Seq.of(_283_v7, _283_v7, _283_v7, _283_v7)))) : (_282_v6)));
      let _lhs22 = _277_v2;
      let _lhs23 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
      _276_v1 = _rhs43;
      _lhs22[_lhs23] = _rhs44;
      _280_v4 = _rhs45;
      _276_v1 = _rhs46;
      let _hi2 = _275_v0;
      for (let _286_i1 = (_275_v0).plus(_275_v0); _286_i1.isLessThan(_hi2); _286_i1 = _286_i1.plus(_dafny.ONE)) {
        let _index39 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_281_v5).length));
        (_281_v5)[_index39] = _286_i1;
        let _index40 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_281_v5).length));
        (_281_v5)[_index40] = _286_i1;
        let _287_v10;
        _287_v10 = _dafny.Map.Empty.slice().updateUnsafe((_281_v5)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_281_v5).length))],_286_i1);
        let _288_v11;
        _288_v11 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm0((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _287_v10, (_dafny.ZERO).minus(_286_i1), _274_globalState));
        _288_v11 = (_288_v11).update(_276_v1, (_286_i1).isLessThan((_281_v5)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_281_v5).length))]));
        let _index41 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_281_v5).length));
        (_281_v5)[_index41] = _275_v0;
        (_274_globalState).f1 = _275_v0;
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_281_v5).length))) {
        let _289_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_289_i2)) && ((_289_i2).isLessThan(new BigNumber((_281_v5).length))))) {
          (_281_v5)[(_289_i2)] = (_289_i2).minus(_275_v0);
        }
      }
      let _290_v12;
      _290_v12 = _dafny.Map.Empty.slice().updateUnsafe(_275_v0,_275_v0);
      if (((((_276_v1) ? (_276_v1) : (true))) ? ((_276_v1) && (_module.__default.fm0((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _290_v12, _275_v0, _274_globalState))) : ((_275_v0).isEqualTo((_dafny.ZERO).minus(_275_v0))))) {
        let _291_v13;
        _291_v13 = _module.D0.create_DC0(_279_v3);
        let _292_v14;
        _292_v14 = _dafny.MultiSet.fromElements(_279_v3, _279_v3, (_291_v13).dtor_cf0);
        let _index42 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_281_v5).length));
        (_281_v5)[_index42] = _module.__default.safeDivisionInt(_275_v0, (((_292_v14).contains(_279_v3)) ? ((_292_v14).get(_279_v3)) : (new BigNumber(996))));
        let _index43 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_281_v5).length));
        let _index44 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        let _index45 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        let _rhs47 = (_275_v0).minus(_275_v0);
        let _rhs48 = (_275_v0).plus(_275_v0);
        let _rhs49 = _276_v1;
        let _rhs50 = _277_v2;
        let _rhs51 = _276_v1;
        let _lhs24 = _281_v5;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_281_v5).length));
        let _lhs26 = _274_globalState;
        let _lhs27 = _277_v2;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        let _lhs29 = _277_v2;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        _lhs24[_lhs25] = _rhs47;
        _lhs26.f1 = _rhs48;
        _lhs27[_lhs28] = _rhs49;
        _277_v2 = _rhs50;
        _lhs29[_lhs30] = _rhs51;
        let _293_v15;
        let _nw53 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _293_v15 = _nw53;
        let _index46 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_293_v15).length));
        (_293_v15)[_index46] = _dafny.Seq.of((_281_v5)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_281_v5).length))], _module.__default.fm1(_276_v1, _274_globalState));
        let _294_v16;
        _294_v16 = _dafny.Seq.of(_275_v0);
        let _index47 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_293_v15).length));
        (_293_v15)[_index47] = _294_v16;
        let _295_v17;
        _295_v17 = _module.D0.create_DC1(_276_v1, _277_v2, (_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
        let _pat_let_tv3 = _277_v2;
        let _pat_let_tv4 = _277_v2;
        let _pat_let_tv5 = _277_v2;
        let _296_v18;
        let _297_v19;
        let _298_v20;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = _module.__default.m0(function (_pat_let4_0) {
          return function (_302_dt__update__tmp_h1) {
            return function (_pat_let8_0) {
              return function (_303_dt__update_hcf3_h1) {
                return _module.D0.create_DC1((_302_dt__update__tmp_h1).dtor_cf1, (_302_dt__update__tmp_h1).dtor_cf2, _303_dt__update_hcf3_h1);
              }(_pat_let8_0);
            }(false);
          }(_pat_let4_0);
        }(function (_pat_let5_0) {
          return function (_299_dt__update__tmp_h0) {
            return function (_pat_let6_0) {
              return function (_300_dt__update_hcf2_h0) {
                return function (_pat_let7_0) {
                  return function (_301_dt__update_hcf3_h0) {
                    return _module.D0.create_DC1((_299_dt__update__tmp_h0).dtor_cf1, _300_dt__update_hcf2_h0, _301_dt__update_hcf3_h0);
                  }(_pat_let7_0);
                }(!((_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_pat_let_tv4).length))]));
              }(_pat_let6_0);
            }(_pat_let_tv3);
          }(_pat_let5_0);
        }(_295_v17)), (_281_v5)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_281_v5).length))], _274_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _296_v18 = _out0;
        _297_v19 = _out1;
        _298_v20 = _out2;
        _295_v17 = _295_v17;
        let _304_v21;
        let _305_v22;
        let _306_v23;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = _module.__default.m0(_295_v17, _296_v18, _274_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _304_v21 = _out3;
        _305_v22 = _out4;
        _306_v23 = _out5;
      } else {
        _283_v7 = _module.__default.fm2(_dafny.Seq.of(_275_v0, _module.__default.fm1(_module.__default.fm0((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _290_v12, (_dafny.ZERO).minus(_275_v0), _274_globalState), _274_globalState)), _module.__default.safeDivisionInt(_module.__default.fm1(_276_v1, _274_globalState), _275_v0), _275_v0, !(_276_v1), _274_globalState);
        let _307_v24;
        _307_v24 = _dafny.MultiSet.fromElements(_275_v0, new BigNumber(556));
        _276_v1 = ((_307_v24).IsProperSubsetOf(_307_v24)) || (_276_v1);
        let _index48 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        (_277_v2)[_index48] = !((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]) || ((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
        (_274_globalState).f1 = _275_v0;
        let _308_v25;
        _308_v25 = _dafny.MultiSet.fromElements((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
        let _309_v26;
        _309_v26 = _dafny.Map.Empty.slice().updateUnsafe(_276_v1,(((_308_v25).contains(_276_v1)) ? ((_308_v25).get(_276_v1)) : (_275_v0)));
        let _310_v27;
        _310_v27 = _dafny.Map.Empty.slice().updateUnsafe(_283_v7,new BigNumber((_309_v26).length));
        let _311_v28;
        _311_v28 = _module.D0.create_DC0(_dafny.Set.fromElements(_276_v1));
        let _312_v29;
        _312_v29 = _dafny.MultiSet.fromElements(_module.__default.fm3(_module.__default.fm0(_276_v1, _290_v12, _275_v0, _274_globalState), _275_v0, _279_v3, _310_v27, _274_globalState), _311_v28, _module.D0.create_DC0(_279_v3), _311_v28, _311_v28);
        (_274_globalState).f1 = new BigNumber((_312_v29).cardinality());
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_277_v2).length))) {
        let _313_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_313_i3)) && ((_313_i3).isLessThan(new BigNumber((_277_v2).length))))) {
          (_277_v2)[(_313_i3)] = false;
        }
      }
      let _index49 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length));
      (_281_v5)[_index49] = new BigNumber((_283_v7).length);
      let _314_v30;
      _314_v30 = _dafny.Set.fromElements(_275_v0);
      let _315_v31;
      _315_v31 = _dafny.Map.Empty.slice().updateUnsafe(_314_v30,_275_v0);
      let _316_v32;
      _316_v32 = _dafny.Seq.of((_dafny.ZERO).minus(_275_v0), new BigNumber(405), new BigNumber(226), new BigNumber((_315_v31).length));
      let _index50 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length));
      (_281_v5)[_index50] = new BigNumber((_316_v32).length);
      let _317_v33;
      _317_v33 = _dafny.Map.Empty.slice().updateUnsafe(_275_v0,(_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
      let _318_v34;
      _318_v34 = _dafny.Map.Empty.slice().updateUnsafe(_277_v2,_317_v33);
      let _319_v35;
      let _nw54 = new _module.C4();
      _nw54.__ctor(_276_v1, _318_v34, new BigNumber(745), (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
      _319_v35 = _nw54;
      (_274_globalState).f1 = new BigNumber(-937);
      let _320_v37;
      _320_v37 = _module.D7.create_DC16(function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of (_283_v7).Elements) {
    let _321_v36 = _compr_13;
    if (_dafny.Seq.contains(_283_v7, _321_v36)) {
      _coll13.push([_321_v36,_275_v0]);
    }
  }
  return _coll13;
}());
      let _322_v38;
      _322_v38 = _dafny.Map.Empty.slice().updateUnsafe((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))],_320_v37);
      _322_v38 = _322_v38;
      let _323_v39;
      _323_v39 = _dafny.MultiSet.fromElements((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], _275_v0);
      let _hi3 = new BigNumber((_323_v39).cardinality());
      for (let _324_i4 = (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]; _324_i4.isLessThan(_hi3); _324_i4 = _324_i4.plus(_dafny.ONE)) {
        let _index51 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length));
        (_281_v5)[_index51] = (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))];
        let _325_v40;
        _325_v40 = _dafny.Seq.of((_319_v35).f11, (_319_v35).f11);
        let _326_v41;
        _326_v41 = new _dafny.CodePoint('h'.codePointAt(0));
        let _327_v42;
        let _nw55 = Array((new BigNumber(6)).toNumber()).fill([]);
        _327_v42 = _nw55;
        let _328_v43;
        _328_v43 = _dafny.Map.Empty.slice().updateUnsafe(true,(_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
        let _329_v44;
        let _nw56 = new _module.C5();
        _nw56.__ctor(_327_v42, _dafny.Seq.of(new BigNumber((_328_v43).length), (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]), _324_i4, (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
        _329_v44 = _nw56;
        let _330_v45;
        _330_v45 = _dafny.MultiSet.fromElements(_329_v44);
        let _331_v46;
        _331_v46 = _module.D5.create_DC11(_325_v40);
        let _332_v47;
        _332_v47 = _module.D8.create_DC19(_276_v1, (_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _282_v6, _275_v0, (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
        let _333_v48;
        let _nw57 = Array((new BigNumber(29)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_325_v40, _325_v40);
        _nw57[(_dafny.ONE).toNumber()] = _325_v40;
        _nw57[(new BigNumber(2)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(3)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
        _nw57[(new BigNumber(5)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(6)).toNumber()] = ((_276_v1) ? (_325_v40) : (_dafny.Seq.update(_dafny.Seq.of((_319_v35).f11, (_319_v35).f11), _module.__default.safeIndex(_275_v0, new BigNumber((_dafny.Seq.of((_319_v35).f11, (_319_v35).f11)).length)), _276_v1)));
        _nw57[(new BigNumber(7)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(8)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_module.__default.fm27(new BigNumber(-187), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-901)), function (_334_i5) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _326_v41, _274_globalState), _module.__default.safeIndex((((_330_v45).contains(_329_v44)) ? ((_330_v45).get(_329_v44)) : (new BigNumber(538))), new BigNumber((_module.__default.fm27(new BigNumber(-187), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-901)), function (_335_i5) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _326_v41, _274_globalState)).length)), (_319_v35).f11);
        _nw57[(new BigNumber(10)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(11)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(12)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(13)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(14)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(15)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(16)).toNumber()] = (_331_v46).dtor_cf20;
        _nw57[(new BigNumber(17)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_325_v40, _325_v40);
        _nw57[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_325_v40, _325_v40);
        _nw57[(new BigNumber(20)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_325_v40, _325_v40);
        _nw57[(new BigNumber(22)).toNumber()] = _module.__default.fm21(_274_globalState);
        _nw57[(new BigNumber(23)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_325_v40, _dafny.Seq.of((_332_v47).dtor_cf34));
        _nw57[(new BigNumber(25)).toNumber()] = _dafny.Seq.of((_319_v35).f11);
        _nw57[(new BigNumber(26)).toNumber()] = _325_v40;
        _nw57[(new BigNumber(27)).toNumber()] = (((_319_v35).f11) ? (_325_v40) : (_325_v40));
        _nw57[(new BigNumber(28)).toNumber()] = (_331_v46).dtor_cf20;
        _333_v48 = _nw57;
        let _index52 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_333_v48).length));
        (_333_v48)[_index52] = _325_v40;
        let _index53 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_333_v48).length));
        (_333_v48)[_index53] = _325_v40;
        let _336_v49;
        _336_v49 = _module.D0.create_DC0(_279_v3);
        let _337_v50;
        _337_v50 = _dafny.Map.Empty.slice().updateUnsafe(_336_v49,_276_v1);
        let _338_v51;
        _338_v51 = _module.D1.create_DC2(_337_v50);
        let _339_v52;
        _339_v52 = _module.D2.create_DC6(_338_v51, _dafny.Seq.UnicodeFromString("vku"), (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
        let _340_v53;
        _340_v53 = _dafny.Map.Empty.slice().updateUnsafe(_276_v1,(_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
        let _341_v54;
        let _nw58 = new _module.C1();
        _nw58.__ctor(_326_v41, _324_i4, (((_340_v53).contains((_319_v35).f11)) ? ((_340_v53).get((_319_v35).f11)) : (new BigNumber((_316_v32).length))));
        _341_v54 = _nw58;
        let _342_v55;
        _342_v55 = _dafny.Map.Empty.slice().updateUnsafe(_341_v54,_283_v7);
        _283_v7 = _dafny.Seq.Concat((_339_v52).dtor_cf11, (((_342_v55).contains(_341_v54)) ? ((_342_v55).get(_341_v54)) : (_283_v7)));
        let _init7 = ((_343_v2) => function (_344_i6) {
          return (_343_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_343_v2).length))];
        })(_277_v2);
        let _nw59 = Array((new BigNumber(22)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw59.length); _i0_7++) {
          _nw59[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _277_v2 = _nw59;
      }
      let _345_i7;
      _345_i7 = _dafny.ZERO;
      L0: {
        while ((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_345_i7)) {
              break L0;
            }
            _345_i7 = (_345_i7).plus(_dafny.ONE);
            (_274_globalState).f1 = (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))];
            let _346_v56;
            let _init8 = ((_347_v35) => function (_348_i8) {
              return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of((_347_v35).f11, (_347_v35).f11));
            })(_319_v35);
            let _nw60 = Array((new BigNumber(14)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw60.length); _i0_8++) {
              _nw60[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _346_v56 = _nw60;
            let _349_v57;
            _349_v57 = _dafny.Seq.of(true, _276_v1);
            let _index54 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_346_v56).length));
            (_346_v56)[_index54] = _dafny.Seq.Concat(_349_v57, _349_v57);
            let _350_v58;
            _350_v58 = _module.D1.create_DC4(_276_v1);
            let _351_v59;
            _351_v59 = _module.D2.create_DC7((_319_v35).f11, new BigNumber((_279_v3).length), _317_v33, _350_v58);
            let _index55 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_346_v56).length));
            (_346_v56)[_index55] = _dafny.Seq.Concat(((!(_module.__default.fm0((_351_v59).dtor_cf13, _290_v12, new BigNumber((_283_v7).length), _274_globalState))) ? (_349_v57) : (_349_v57)), _dafny.Seq.update(_dafny.Seq.of(_276_v1, (((_317_v33).contains(_275_v0)) ? ((_317_v33).get(_275_v0)) : ((_319_v35).f11))), _module.__default.safeIndex(_275_v0, new BigNumber((_dafny.Seq.of(_276_v1, (((_317_v33).contains(_275_v0)) ? ((_317_v33).get(_275_v0)) : ((_319_v35).f11)))).length)), true));
            let _352_v60;
            let _nw61 = new _module.C0();
            _nw61.__ctor((_314_v30).IsSubsetOf(_dafny.Set.fromElements((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))])));
            _352_v60 = _nw61;
            let _353_v61;
            let _nw62 = new _module.C4();
            _nw62.__ctor((_319_v35).f11, _319_v35.f12, (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], _module.__default.fm1((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _274_globalState));
            _353_v61 = _nw62;
            let _354_v62;
            _354_v62 = _dafny.MultiSet.fromElements(_353_v61, _353_v61, _353_v61, _353_v61);
            let _355_v63;
            _355_v63 = _dafny.Map.Empty.slice().updateUnsafe(_354_v62,false);
            let _356_v64;
            let _nw63 = new _module.C4();
            _nw63.__ctor(_module.__default.fm0((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _290_v12, new BigNumber((_dafny.Seq.update(_349_v57, _module.__default.safeIndex(_275_v0, new BigNumber((_349_v57).length)), (_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))])).length), _274_globalState), _318_v34, _275_v0, new BigNumber(((_355_v63).Merge(_355_v63)).length));
            _356_v64 = _nw63;
            let _357_v65;
            _357_v65 = new _dafny.CodePoint('k'.codePointAt(0));
            let _rhs52 = _dafny.Seq.Concat(_283_v7, _dafny.Seq.of(_357_v65, _module.__default.fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(828)), ((_358_v65) => function (_359_i9) {
              return _358_v65;
            })(_357_v65)), new _dafny.CodePoint('p'.codePointAt(0)), false, (_319_v35).f11, _274_globalState), _357_v65, _357_v65));
            let _rhs53 = _353_v61;
            _283_v7 = _rhs52;
            _356_v64 = _rhs53;
          }
        }
      }
      let _index56 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
      (_277_v2)[_index56] = (_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))];
      let _360_v66;
      _360_v66 = _module.D0.create_DC0(_279_v3);
      let _361_v67;
      _361_v67 = _dafny.Map.Empty.slice().updateUnsafe(_360_v66,(_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
      let _362_v68;
      _362_v68 = _module.D1.create_DC2(_361_v67);
      let _source5 = _362_v68;
      if (_source5.is_DC3) {
        let _363___mcc_h0 = (_source5).cf5;
        let _364___mcc_h1 = (_source5).cf6;
        let _365___mcc_h2 = (_source5).cf7;
        let _366_cf7 = _365___mcc_h2;
        let _367_cf6 = _364___mcc_h1;
        let _368_cf5 = _363___mcc_h0;
        let _369_v69;
        _369_v69 = _dafny.Map.Empty.slice().updateUnsafe(_368_cf5,_module.__default.fm1(false, _274_globalState));
        let _370_v70;
        let _371_v71;
        let _372_v72;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector2 = (_319_v35).m1(!(!((_366_cf7).isLessThanOrEqualTo((((_369_v69).contains(false)) ? ((_369_v69).get(false)) : (_367_cf6))))), new BigNumber((_283_v7).length), new BigNumber(103), _274_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _370_v70 = _out6;
        _371_v71 = _out7;
        _372_v72 = _out8;
        let _373_v73;
        _373_v73 = _dafny.Seq.of(_module.__default.fm0(!(_368_cf5), _290_v12, _367_cf6, _274_globalState), _276_v1, _370_v70);
        let _374_v74;
        _374_v74 = _dafny.Seq.of(_314_v30, _314_v30);
        let _375_v75;
        let _376_v76;
        let _377_v77;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector3 = (_319_v35).m1((_373_v73)[_module.__default.safeIndex(_275_v0, new BigNumber((_373_v73).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_374_v74).length)).minus(_367_cf6))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_366_cf7,_367_cf6)).length), _274_globalState);
        _out9 = _outcollector3[0];
        _out10 = _outcollector3[1];
        _out11 = _outcollector3[2];
        _375_v75 = _out9;
        _376_v76 = _out10;
        _377_v77 = _out11;
        (_319_v35).m5((_dafny.ZERO).minus(((_368_cf5) ? ((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]) : (_275_v0))), _274_globalState);
        let _378_v78;
        _378_v78 = new _dafny.CodePoint('q'.codePointAt(0));
        let _379_v79;
        _379_v79 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(_378_v78));
        let _380_v80;
        _380_v80 = _dafny.MultiSet.fromElements(_378_v78, _378_v78, _378_v78, _378_v78, _378_v78);
        _379_v79 = _dafny.Set.fromElements(_380_v80, _380_v80);
      } else if (_source5.is_DC4) {
        let _381___mcc_h3 = (_source5).cf8;
        let _382_cf8 = _381___mcc_h3;
        (_319_v35).m5(_275_v0, _274_globalState);
        let _383_v81;
        let _nw64 = new _module.C2();
        _nw64.__ctor((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], _275_v0);
        _383_v81 = _nw64;
        _277_v2 = _277_v2;
        let _384_v82;
        _384_v82 = _dafny.MultiSet.fromElements(!(false));
        if (((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]).isLessThan(new BigNumber((((_382_cf8) ? ((_384_v82).update(_276_v1, _module.__default.abs((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]))) : ((_384_v82).update((_319_v35).f11, _module.__default.abs(_275_v0))))).cardinality()))) {
          (_319_v35).m5(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kdrbwar"), _283_v7)).length), _274_globalState);
          (_274_globalState).f1 = ((!(_382_cf8)) ? (new BigNumber(785)) : (new BigNumber((_dafny.Seq.UnicodeFromString("tmolciku")).length)));
          let _385_v83;
          _385_v83 = _dafny.Seq.of(_277_v2);
          _277_v2 = (_385_v83)[_module.__default.safeIndex((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], new BigNumber((_385_v83).length))];
          let _386_v84;
          let _nw65 = Array((new BigNumber(14)).toNumber());
          _386_v84 = _nw65;
          let _387_v85;
          _387_v85 = _dafny.Seq.of(_386_v84, _386_v84, _386_v84);
          _386_v84 = (_387_v85)[_module.__default.safeIndex((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], new BigNumber((_387_v85).length))];
          let _index57 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
          (_277_v2)[_index57] = (_319_v35).f11;
        } else {
          let _388_v86;
          let _nw66 = new _module.C0();
          _nw66.__ctor((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]);
          _388_v86 = _nw66;
          let _389_v87;
          _389_v87 = _388_v86;
          _389_v87 = _389_v87;
          _382_cf8 = (_319_v35).f11;
          let _390_v88;
          _390_v88 = _dafny.Seq.of((_382_cf8) || (!(_388_v86.f10)), _382_cf8);
          _390_v88 = _dafny.Seq.Concat(_390_v88, _390_v88);
          let _391_v89;
          _391_v89 = _module.D1.create_DC3((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], (_dafny.ZERO).minus(_275_v0), new BigNumber((_323_v39).cardinality()));
          let _index58 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
          let _index59 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length));
          let _rhs54 = ((_388_v86.f10) ? ((new BigNumber((_314_v30).length)).isLessThanOrEqualTo(new BigNumber(-192))) : ((_319_v35).f11));
          let _rhs55 = (_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))];
          let _rhs56 = _module.__default.safeDivisionInt((_275_v0).minus(new BigNumber((_390_v88).length)), (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
          let _rhs57 = ((new BigNumber((_317_v33).length)).multipliedBy(_275_v0)).isEqualTo(_275_v0);
          let _rhs58 = (_391_v89).dtor_cf5;
          let _lhs31 = _277_v2;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
          let _lhs33 = _281_v5;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length));
          _276_v1 = _rhs54;
          _lhs31[_lhs32] = _rhs55;
          _lhs33[_lhs34] = _rhs56;
          _382_cf8 = _rhs57;
          _382_cf8 = _rhs58;
          (_274_globalState).f1 = _module.__default.safeModuloInt(_275_v0, (((_290_v12).contains((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))])) ? ((_290_v12).get((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))])) : (_275_v0)));
        }
      } else {
        let _392___mcc_h4 = (_source5).cf4;
        let _393_cf4 = _392___mcc_h4;
        let _394_v90;
        _394_v90 = _dafny.Seq.of(_277_v2, _277_v2);
        let _395_v91;
        let _nw67 = Array((new BigNumber(12)).toNumber());
        _nw67[(_dafny.ZERO).toNumber()] = _277_v2;
        _nw67[(_dafny.ONE).toNumber()] = (_394_v90)[_module.__default.safeIndex((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], new BigNumber((_394_v90).length))];
        _nw67[(new BigNumber(2)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(3)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(4)).toNumber()] = (_module.D0.create_DC1((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _277_v2, (_319_v35).f11)).dtor_cf2;
        _nw67[(new BigNumber(5)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(6)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(7)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(8)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(9)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(10)).toNumber()] = _277_v2;
        _nw67[(new BigNumber(11)).toNumber()] = _277_v2;
        _395_v91 = _nw67;
        let _index60 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_395_v91).length));
        (_395_v91)[_index60] = (((_319_v35).f11) ? (_277_v2) : (_277_v2));
        let _index61 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_395_v91).length));
        (_395_v91)[_index61] = _277_v2;
        _290_v12 = (_290_v12).update((_dafny.ZERO).minus((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]), (new BigNumber((_283_v7).length)).plus((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]));
        let _index62 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        (_277_v2)[_index62] = _276_v1;
        _275_v0 = (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))];
      }
      if ((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))]) {
        _319_v35 = _319_v35;
        let _396_v92;
        _396_v92 = new _dafny.CodePoint('p'.codePointAt(0));
        _396_v92 = (_283_v7)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_275_v0, new BigNumber((_279_v3).length)), new BigNumber((_283_v7).length))];
        _290_v12 = (_290_v12).update(_module.__default.fm1((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _274_globalState), _module.__default.fm1((_277_v2)[_module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length))], _274_globalState));
        _275_v0 = _module.__default.fm1(_276_v1, _274_globalState);
        let _397_v94;
        let _nw68 = Array((new BigNumber(6)).toNumber());
        _nw68[(_dafny.ZERO).toNumber()] = _314_v30;
        _nw68[(_dafny.ONE).toNumber()] = _314_v30;
        _nw68[(new BigNumber(2)).toNumber()] = _314_v30;
        _nw68[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_275_v0, _275_v0, _275_v0, (_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
        _nw68[(new BigNumber(4)).toNumber()] = function () {
          let _coll14 = new _dafny.Set();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(877), new BigNumber(434))) {
            let _398_v93 = _compr_14;
            if (((new BigNumber(877)).isLessThanOrEqualTo(_398_v93)) && ((_398_v93).isLessThan(new BigNumber(434)))) {
              _coll14.add((_398_v93).multipliedBy(_275_v0));
            }
          }
          return _coll14;
        }();
        _nw68[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_275_v0, _275_v0, _275_v0, _275_v0, _275_v0);
        _397_v94 = _nw68;
        let _399_v95;
        _399_v95 = _module.D11.create_DC28(_283_v7, _397_v94, (_319_v35).f11, _276_v1);
        let _400_v96;
        _400_v96 = _module.D11.create_DC29(_399_v95);
        _400_v96 = _400_v96;
      } else {
        let _401_v97;
        _401_v97 = _dafny.Set.fromElements(_277_v2);
        let _402_v98;
        _402_v98 = _dafny.Map.Empty.slice().updateUnsafe(_276_v1,(_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))]);
        let _index63 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_277_v2).length));
        (_277_v2)[_index63] = !(!(_module.__default.fm0((_319_v35).f11, (_290_v12).update(new BigNumber(((_323_v39).update(_275_v0, _module.__default.abs(new BigNumber((_401_v97).length)))).cardinality()), new BigNumber((_402_v98).length)), _275_v0, _274_globalState)));
        let _403_v99;
        let _nw69 = new _module.C0();
        _nw69.__ctor((_319_v35).f11);
        _403_v99 = _nw69;
        let _404_v100;
        _404_v100 = _403_v99;
        _404_v100 = _404_v100;
        (_403_v99).f10 = _module.__default.fm0(_module.__default.fm0((_319_v35).f11, _290_v12, _275_v0, _274_globalState), _290_v12, new BigNumber((_module.__default.fm40((_281_v5)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_281_v5).length))], _403_v99.f10, _274_globalState)).length), _274_globalState);
        (_274_globalState).f1 = (_dafny.ZERO).minus(_275_v0);
        (_274_globalState).f1 = new BigNumber(-252);
      }
      process.stdout.write(_dafny.toString((_274_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_274_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_275_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_276_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v2)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v3).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_282_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_283_v7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v8).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("klhmhabdk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_285_v9).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_290_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-379),new BigNumber(-379)).updateUnsafe(new BigNumber(-4),new BigNumber(567)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_314_v30).equals(_dafny.Set.fromElements(new BigNumber(-379)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v31).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(-379)),new BigNumber(-379)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_316_v32, _dafny.Seq.of(new BigNumber(379), new BigNumber(405), new BigNumber(226), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_317_v33).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-379),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_318_v34).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_319_v35).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_319_v35.f12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v37).dtor_cf30).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),new BigNumber(-379)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v38).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D7.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber(-379)).updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),new BigNumber(-379)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_323_v39).equals(_dafny.MultiSet.fromElements(new BigNumber(4), new BigNumber(-379)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_345_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_360_v66).dtor_cf0).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_361_v67).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(false)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_362_v68).dtor_cf4).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements(false)),false))));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, [], false);
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
    static create_DC3(cf5, cf6, cf7) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC6(cf10, cf11, cf12) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7(cf13, cf14, cf15, cf16) {
      let $dt = new D2(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D2(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC8(cf17) {
      let $dt = new D2(3);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ", " + this.cf11.toVerbatimString(true) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf9 === other.cf9;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_module.D1.Default(), _dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC9(cf18) {
      let $dt = new D3(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf18) + ")";
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf19) {
      let $dt = new D4(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19;
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf21) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf20) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(_dafny.ZERO);
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
    static create_DC14(cf23, cf24, cf25) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC15(cf26, cf27, cf28, cf29) {
      let $dt = new D6(1);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC13(cf22) {
      let $dt = new D6(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf22 === other.cf22;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC14(_dafny.ZERO, _dafny.MultiSet.Empty, _dafny.ZERO);
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
    static create_DC17(cf31, cf32) {
      let $dt = new D7(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC16(cf30) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + this.cf31.toVerbatimString(true) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(_dafny.Seq.UnicodeFromString(""), _dafny.Set.Empty);
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
    static create_DC19(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D8(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC20(cf39, cf40, cf41, cf42) {
      let $dt = new D8(1);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC21(cf43, cf44) {
      let $dt = new D8(2);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D8(3);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39 && this.cf40 === other.cf40 && this.cf41 === other.cf41 && this.cf42 === other.cf42;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf33 === other.cf33;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC19(false, false, _dafny.Set.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC23(cf46, cf47, cf48, cf49, cf50) {
      let $dt = new D9(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC22(cf45) {
      let $dt = new D9(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(false, _dafny.ZERO, false, false, _dafny.ZERO);
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
    static create_DC25() {
      let $dt = new D10(0);
      return $dt;
    }
    static create_DC26(cf52) {
      let $dt = new D10(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC24(cf51) {
      let $dt = new D10(2);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC25";
      } else if (this.$tag === 1) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf51) + ")";
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
        return other.$tag === 1 && this.cf52 === other.cf52;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC25();
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
    static create_DC28(cf54, cf55, cf56, cf57) {
      let $dt = new D11(0);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC27(cf53) {
      let $dt = new D11(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC29(cf58) {
      let $dt = new D11(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf58() { return this.cf58; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + this.cf54.toVerbatimString(true) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf58) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54) && this.cf55 === other.cf55 && this.cf56 === other.cf56 && this.cf57 === other.cf57;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf58, other.cf58);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC28(_dafny.Seq.UnicodeFromString(""), [], false, false);
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
    static create_DC31(cf60) {
      let $dt = new D12(0);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC30(cf59) {
      let $dt = new D12(1);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf59) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC31(_dafny.ZERO);
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
    static create_DC33(cf62, cf63) {
      let $dt = new D13(0);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC32(cf61) {
      let $dt = new D13(1);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC33(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.ZERO;
      this._f0 = false;
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
      this._f5 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f10 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f10) {
      let _this = this;
      (_this).f10 = f10;
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f7 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f13 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f13, f6, f7) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _405_v0;
      _405_v0 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
      let _406_v1;
      _406_v1 = _dafny.Seq.UnicodeFromString("sm");
      let _407_v2;
      _407_v2 = _dafny.Seq.of(_module.__default.fm0(!(p0), ((_405_v0).update(p1, _this.f7)).update(new BigNumber((_406_v1).length), (_this).f6), (_this).f6, globalState), p0, p0, true, p0);
      let _408_v3;
      let _nw70 = Array((new BigNumber(9)).toNumber());
      _nw70[(_dafny.ZERO).toNumber()] = p0;
      _nw70[(_dafny.ONE).toNumber()] = p0;
      _nw70[(new BigNumber(2)).toNumber()] = p0;
      _nw70[(new BigNumber(3)).toNumber()] = true;
      _nw70[(new BigNumber(4)).toNumber()] = p0;
      _nw70[(new BigNumber(5)).toNumber()] = p0;
      _nw70[(new BigNumber(6)).toNumber()] = p0;
      _nw70[(new BigNumber(7)).toNumber()] = p0;
      _nw70[(new BigNumber(8)).toNumber()] = p0;
      _408_v3 = _nw70;
      let _409_v4;
      _409_v4 = _dafny.Map.Empty.slice().updateUnsafe(_408_v3,_407_v2);
      let _410_v5;
      _410_v5 = _dafny.Set.fromElements(_407_v2, (((_409_v4).contains(_408_v3)) ? ((_409_v4).get(_408_v3)) : (_407_v2)), _407_v2, _407_v2, _407_v2);
      let _rhs59 = _410_v5;
      let _rhs60 = p0;
      let _rhs61 = ((_this).f6).plus(p2);
      let _lhs35 = _this;
      _410_v5 = _rhs59;
      r1 = _rhs60;
      _lhs35.f7 = _rhs61;
      let _411_v6;
      let _nw71 = new _module.C0();
      _nw71.__ctor(p0);
      _411_v6 = _nw71;
      let _412_v7;
      _412_v7 = _dafny.Map.Empty.slice().updateUnsafe(_406_v1,_411_v6);
      _412_v7 = (_412_v7).update(_406_v1, _411_v6);
      let _413_v8;
      let _nw72 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _413_v8 = _nw72;
      let _414_v9;
      _414_v9 = _module.D2.create_DC5(_413_v8);
      let _source6 = ((_411_v6.f10) ? (_414_v9) : (_414_v9));
      if (_source6.is_DC6) {
        let _415___mcc_h0 = (_source6).cf10;
        let _416___mcc_h1 = (_source6).cf11;
        let _417___mcc_h2 = (_source6).cf12;
        let _418_cf12 = _417___mcc_h2;
        let _419_cf11 = _416___mcc_h1;
        let _420_cf10 = _415___mcc_h0;
        _419_cf11 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("g"), _dafny.Seq.UnicodeFromString("yo"));
        (globalState).f1 = _this.f7;
        let _421_v10;
        _421_v10 = _module.D1.create_DC4(_411_v6.f10);
        (_411_v6).f10 = (_421_v10).dtor_cf8;
        let _422_v11;
        _422_v11 = _dafny.Seq.of(_411_v6);
        let _423_v12;
        _423_v12 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_411_v6);
        let _424_v13;
        let _nw73 = Array((new BigNumber(11)).toNumber());
        _nw73[(_dafny.ZERO).toNumber()] = _411_v6;
        _nw73[(_dafny.ONE).toNumber()] = _411_v6;
        _nw73[(new BigNumber(2)).toNumber()] = _411_v6;
        _nw73[(new BigNumber(3)).toNumber()] = _411_v6;
        _nw73[(new BigNumber(4)).toNumber()] = _411_v6;
        _nw73[(new BigNumber(5)).toNumber()] = _411_v6;
        _nw73[(new BigNumber(6)).toNumber()] = _411_v6;
        _nw73[(new BigNumber(7)).toNumber()] = _411_v6;
        _nw73[(new BigNumber(8)).toNumber()] = (_422_v11)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_422_v11).length))];
        _nw73[(new BigNumber(9)).toNumber()] = (((_423_v12).contains(_this.f7)) ? ((_423_v12).get(_this.f7)) : (_411_v6));
        _nw73[(new BigNumber(10)).toNumber()] = _411_v6;
        _424_v13 = _nw73;
        let _rhs62 = _411_v6.f10;
        let _rhs63 = _411_v6.f10;
        let _rhs64 = _411_v6.f10;
        let _rhs65 = _424_v13;
        let _rhs66 = p0;
        let _lhs36 = _411_v6;
        r1 = _rhs62;
        _lhs36.f10 = _rhs63;
        r0 = _rhs64;
        _424_v13 = _rhs65;
        r0 = _rhs66;
      } else if (_source6.is_DC7) {
        let _425___mcc_h3 = (_source6).cf13;
        let _426___mcc_h4 = (_source6).cf14;
        let _427___mcc_h5 = (_source6).cf15;
        let _428___mcc_h6 = (_source6).cf16;
        let _429_cf16 = _428___mcc_h6;
        let _430_cf15 = _427___mcc_h5;
        let _431_cf14 = _426___mcc_h4;
        let _432_cf13 = _425___mcc_h3;
        let _index64 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_413_v8).length));
        (_413_v8)[_index64] = new BigNumber(341);
        let _index65 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_413_v8).length));
        (_413_v8)[_index65] = new BigNumber(-810);
        r2 = _431_cf14;
        let _433_v14;
        _433_v14 = new _dafny.CodePoint('t'.codePointAt(0));
        let _434_v15;
        _434_v15 = _dafny.MultiSet.fromElements(_dafny.Seq.contains(_407_v2, _411_v6.f10), _411_v6.f10);
        let _435_v16;
        _435_v16 = _dafny.Seq.of((new BigNumber((_407_v2).length)).multipliedBy(_this.f7));
        let _436_v17;
        let _nw74 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _436_v17 = _nw74;
        let _437_v18;
        _437_v18 = _dafny.Map.Empty.slice().updateUnsafe(_411_v6.f10,_436_v17);
        let _438_v19;
        _438_v19 = _dafny.Set.fromElements(p0);
        let _439_v20;
        _439_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_438_v19),_411_v6.f10);
        let _440_v21;
        _440_v21 = _module.D1.create_DC2(_439_v20);
        let _441_v22;
        _441_v22 = _dafny.MultiSet.fromElements((_this).f6, new BigNumber(135), p2, (_this).f6, (_this).f6);
        let _442_v23;
        _442_v23 = _module.D2.create_DC6(_440_v21, _406_v1, (((_441_v22).contains(new BigNumber(595))) ? ((_441_v22).get(new BigNumber(595))) : (_this.f7)));
        let _443_v24;
        _443_v24 = _dafny.Map.Empty.slice().updateUnsafe((((_437_v18).contains(_411_v6.f10)) ? ((_437_v18).get(_411_v6.f10)) : (_436_v17)),_442_v23);
        let _index66 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_413_v8).length));
        let _rhs67 = _module.__default.fm25(globalState);
        let _rhs68 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_435_v16).length)));
        let _rhs69 = _dafny.MultiSet.fromElements(_432_cf13, !(_443_v24).equals(_443_v24));
        let _rhs70 = _this.f7;
        let _lhs37 = globalState;
        let _lhs38 = _413_v8;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_413_v8).length));
        _433_v14 = _rhs67;
        _lhs37.f1 = _rhs68;
        _434_v15 = _rhs69;
        _lhs38[_lhs39] = _rhs70;
        let _444_v25;
        _444_v25 = _dafny.Set.fromElements((_this).f6, p1, new BigNumber((_dafny.Set.fromElements((_this).f6, (_this).f6)).length));
        let _445_v26;
        _445_v26 = _dafny.MultiSet.fromElements(_444_v25, _444_v25, _444_v25, _444_v25, _444_v25);
        r2 = new BigNumber(((_445_v26).Difference((_445_v26).Intersect(_445_v26))).cardinality());
      } else if (_source6.is_DC5) {
        let _446___mcc_h7 = (_source6).cf9;
        let _447_cf9 = _446___mcc_h7;
        let _448_v27;
        _448_v27 = new _dafny.CodePoint('d'.codePointAt(0));
        _448_v27 = _448_v27;
        r1 = _411_v6.f10;
        let _index67 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_408_v3).length));
        (_408_v3)[_index67] = true;
        let _index68 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_408_v3).length));
        (_408_v3)[_index68] = !(_411_v6.f10);
        r0 = (_408_v3)[_module.__default.safeIndex(new BigNumber(963), new BigNumber((_408_v3).length))];
      } else {
        let _449___mcc_h8 = (_source6).cf17;
        let _450_cf17 = _449___mcc_h8;
        let _index69 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_408_v3).length));
        (_408_v3)[_index69] = p0;
        let _451_v28;
        _451_v28 = _dafny.Seq.of(_406_v1, _406_v1, _406_v1);
        let _452_v29;
        _452_v29 = _module.D0.create_DC0(_dafny.Set.fromElements(p0, p0));
        let _453_v30;
        _453_v30 = _dafny.Map.Empty.slice().updateUnsafe(_452_v29,_411_v6.f10);
        let _454_v31;
        _454_v31 = _module.D1.create_DC2(_453_v30);
        let _455_v32;
        _455_v32 = _module.D2.create_DC6(_454_v31, _dafny.Seq.UnicodeFromString("ukxgv"), (_this).f6);
        let _index70 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_408_v3).length));
        let _rhs71 = (_451_v28)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_406_v1, _406_v1)).length), new BigNumber((_451_v28).length))];
        let _rhs72 = _dafny.Seq.Concat(_dafny.Seq.Concat((_455_v32).dtor_cf11, _406_v1), _dafny.Seq.Concat(_406_v1, _406_v1));
        let _rhs73 = _module.__default.fm0(p0, _405_v0, _this.f7, globalState);
        let _lhs40 = _408_v3;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_408_v3).length));
        _406_v1 = _rhs71;
        _406_v1 = _rhs72;
        _lhs40[_lhs41] = _rhs73;
        let _456_v33;
        _456_v33 = _dafny.Map.Empty.slice().updateUnsafe(_411_v6.f10,p2);
        let _457_v34;
        _457_v34 = _dafny.MultiSet.fromElements(_413_v8);
        let _index71 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_408_v3).length));
        let _rhs74 = ((_dafny.ZERO).minus((((_456_v33).contains(_411_v6.f10)) ? ((_456_v33).get(_411_v6.f10)) : ((_dafny.ZERO).minus(p2))))).minus(new BigNumber(((_457_v34).update(_413_v8, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_406_v1).length))))).cardinality()));
        let _rhs75 = ((p2).isLessThan((_this).f6)) || (_411_v6.f10);
        let _lhs42 = _this;
        let _lhs43 = _408_v3;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_408_v3).length));
        _lhs42.f7 = _rhs74;
        _lhs43[_lhs44] = _rhs75;
        let _458_v35;
        _458_v35 = _dafny.Set.fromElements(p2);
        let _459_v37;
        _459_v37 = _dafny.Seq.of((_458_v35).Intersect(_458_v35), _458_v35, function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(118), new BigNumber(-912))) {
            let _460_v36 = _compr_15;
            if (((new BigNumber(118)).isLessThanOrEqualTo(_460_v36)) && ((_460_v36).isLessThan(new BigNumber(-912)))) {
              _coll15.add((_460_v36).plus(p1));
            }
          }
          return _coll15;
        }());
        _459_v37 = _module.__default.fm26((_dafny.ZERO).minus((_455_v32).dtor_cf12), p0, globalState);
        let _461_v39;
        let _nw75 = new _module.C0();
        _nw75.__ctor((function () {
          let _coll16 = new _dafny.Set();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-209), new BigNumber(576))) {
            let _462_v38 = _compr_16;
            if (((new BigNumber(-209)).isLessThanOrEqualTo(_462_v38)) && ((_462_v38).isLessThan(new BigNumber(576)))) {
              _coll16.add((_462_v38).plus(new BigNumber((_407_v2).length)));
            }
          }
          return _coll16;
        }()).IsProperSubsetOf(_458_v35));
        _461_v39 = _nw75;
      }
      (_this).f7 = new BigNumber((_406_v1).length);
      if (false) {
        let _index72 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_413_v8).length));
        (_413_v8)[_index72] = p2;
        let _index73 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_413_v8).length));
        (_413_v8)[_index73] = p2;
        (globalState).f1 = _this.f7;
        let _463_v40;
        let _init9 = ((_464_v6, _465_v2) => function (_466_i0) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_464_v6.f10, _464_v6.f10), _465_v2);
        })(_411_v6, _407_v2);
        let _nw76 = Array((new BigNumber(28)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw76.length); _i0_9++) {
          _nw76[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _463_v40 = _nw76;
        let _467_v41;
        _467_v41 = _module.D0.create_DC1(_411_v6.f10, _408_v3, _411_v6.f10);
        let _468_v42;
        _468_v42 = _dafny.Map.Empty.slice().updateUnsafe(_411_v6.f10,!(_411_v6.f10));
        let _469_v43;
        _469_v43 = _module.D5.create_DC11(_407_v2);
        let _470_v44;
        _470_v44 = _dafny.Seq.of(_407_v2, _407_v2);
        let _471_v45;
        _471_v45 = _dafny.Set.fromElements(_this.f7);
        let _472_v46;
        _472_v46 = _dafny.Set.fromElements(new BigNumber((_471_v45).length));
        let _pat_let_tv6 = _411_v6;
        let _nw77 = Array((new BigNumber(28)).toNumber());
        _nw77[(_dafny.ZERO).toNumber()] = _407_v2;
        _nw77[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_407_v2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_406_v1, _module.__default.safeIndex(_this.f7, new BigNumber((_406_v1).length)), (_this).f13)).length), new BigNumber((_407_v2).length)), _411_v6.f10), _407_v2);
        _nw77[(new BigNumber(2)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(3)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(p0, (_467_v41).dtor_cf3);
        _nw77[(new BigNumber(5)).toNumber()] = _dafny.Seq.of((((_468_v42).contains(_411_v6.f10)) ? ((_468_v42).get(_411_v6.f10)) : (_411_v6.f10)));
        _nw77[(new BigNumber(6)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_411_v6.f10, _411_v6.f10), _dafny.Seq.update((_module.D5.create_DC11(_407_v2)).dtor_cf20, _module.__default.safeIndex((_this).f6, new BigNumber(((_module.D5.create_DC11(_407_v2)).dtor_cf20).length)), _411_v6.f10));
        _nw77[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_411_v6.f10, _411_v6.f10);
        _nw77[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_407_v2, (_module.D5.create_DC11(_dafny.Seq.of(false))).dtor_cf20);
        _nw77[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_407_v2, _407_v2);
        _nw77[(new BigNumber(11)).toNumber()] = (_469_v43).dtor_cf20;
        _nw77[(new BigNumber(12)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false, p0, p0));
        _nw77[(new BigNumber(14)).toNumber()] = (_470_v44)[_module.__default.safeIndex(p2, new BigNumber((_470_v44).length))];
        _nw77[(new BigNumber(15)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(16)).toNumber()] = _dafny.Seq.of((_407_v2)[_module.__default.safeIndex(_this.f7, new BigNumber((_407_v2).length))], _411_v6.f10);
        _nw77[(new BigNumber(17)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_407_v2, _407_v2);
        _nw77[(new BigNumber(19)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(20)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_407_v2, _407_v2);
        _nw77[(new BigNumber(22)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(23)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(24)).toNumber()] = _dafny.Seq.of(_module.__default.fm0(_411_v6.f10, _405_v0, p1, globalState));
        _nw77[(new BigNumber(25)).toNumber()] = _407_v2;
        _nw77[(new BigNumber(26)).toNumber()] = _dafny.Seq.of((function (_pat_let9_0) {
          return function (_473_dt__update__tmp_h0) {
            return function (_pat_let10_0) {
              return function (_474_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_474_dt__update_hcf1_h0, (_473_dt__update__tmp_h0).dtor_cf2, (_473_dt__update__tmp_h0).dtor_cf3);
              }(_pat_let10_0);
            }(_pat_let_tv6.f10);
          }(_pat_let9_0);
        }(_467_v41)).dtor_cf1);
        _nw77[(new BigNumber(27)).toNumber()] = _module.__default.fm27(new BigNumber((_471_v45).length), _406_v1, (_this).f13, globalState);
        _463_v40 = _nw77;
        r0 = !(((_411_v6.f10) ? (_411_v6.f10) : ((new BigNumber((_472_v46).length)).isLessThan((_413_v8)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_413_v8).length))]))));
        let _index74 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_413_v8).length));
        (_413_v8)[_index74] = _this.f7;
      } else {
        let _475_v47;
        _475_v47 = _dafny.Map.Empty.slice().updateUnsafe(_413_v8,(_this).f13);
        _475_v47 = _475_v47;
        let _476_v48;
        _476_v48 = _dafny.Map.Empty.slice().updateUnsafe(p2,_411_v6.f10);
        let _477_v49;
        _477_v49 = _dafny.Seq.of(_476_v48);
        let _478_v50;
        _478_v50 = _module.D1.create_DC3(false, new BigNumber(((_477_v49)[_module.__default.safeIndex(new BigNumber((_406_v1).length), new BigNumber((_477_v49).length))]).length), p1);
        r0 = (((_478_v50).dtor_cf5) ? (_411_v6.f10) : (true));
        let _index75 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
        (_408_v3)[_index75] = (_411_v6.f10) || (_411_v6.f10);
        let _index76 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length));
        (_413_v8)[_index76] = p1;
        let _479_v51;
        _479_v51 = _dafny.Set.fromElements(_411_v6.f10, _411_v6.f10);
        let _480_v52;
        _480_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_479_v51).length),_479_v51);
        let _index77 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
        let _index78 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length));
        let _rhs76 = _411_v6.f10;
        let _rhs77 = _411_v6.f10;
        let _rhs78 = (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,_479_v51)).Merge((_480_v52).Merge(_480_v52))).length));
        let _rhs79 = p1;
        let _lhs45 = _411_v6;
        let _lhs46 = _408_v3;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
        let _lhs48 = globalState;
        let _lhs49 = _413_v8;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length));
        _lhs45.f10 = _rhs76;
        _lhs46[_lhs47] = _rhs77;
        _lhs48.f1 = _rhs78;
        _lhs49[_lhs50] = _rhs79;
        let _source7 = _478_v50;
        if (_source7.is_DC3) {
          let _481___mcc_h9 = (_source7).cf5;
          let _482___mcc_h10 = (_source7).cf6;
          let _483___mcc_h11 = (_source7).cf7;
          let _484_cf7 = _483___mcc_h11;
          let _485_cf6 = _482___mcc_h10;
          let _486_cf5 = _481___mcc_h9;
          let _487_v53;
          _487_v53 = new _dafny.CodePoint('u'.codePointAt(0));
          let _488_v54;
          _488_v54 = _dafny.Set.fromElements(_this.f7);
          _487_v53 = (((new BigNumber((_488_v54).length)).isLessThan(p1)) ? ((_this).f13) : (_487_v53));
          let _489_v55;
          _489_v55 = _module.D5.create_DC12(_485_cf6);
          let _490_v56;
          _490_v56 = _dafny.MultiSet.fromElements((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]);
          let _491_v57;
          _491_v57 = _dafny.Map.Empty.slice().updateUnsafe(p2,_490_v56);
          _489_v55 = _module.__default.fm28(((false) ? (_module.__default.fm27(new BigNumber((_491_v57).length), _406_v1, (_this).f13, globalState)) : (_407_v2)), p2, globalState);
          _478_v50 = _478_v50;
          let _492_v59;
          _492_v59 = _dafny.Map.Empty.slice().updateUnsafe(_485_cf6,_487_v53);
          let _493_v60;
          _493_v60 = _dafny.Set.fromElements(_492_v59, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(532),(_this).f13));
          _485_cf6 = ((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]).minus((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_493_v60).Elements) {
              let _494_v58 = _compr_17;
              if ((_493_v60).contains(_494_v58)) {
                _coll17.push([_494_v58,_485_cf6]);
              }
            }
            return _coll17;
          }()).length)));
        } else if (_source7.is_DC4) {
          let _495___mcc_h12 = (_source7).cf8;
          let _496_cf8 = _495___mcc_h12;
          let _497_v61;
          _497_v61 = _dafny.Seq.of((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))], (_this).f6);
          let _498_v62;
          let _nw78 = Array((new BigNumber(5)).toNumber());
          _nw78[(_dafny.ZERO).toNumber()] = _497_v61;
          _nw78[(_dafny.ONE).toNumber()] = _497_v61;
          _nw78[(new BigNumber(2)).toNumber()] = _497_v61;
          _nw78[(new BigNumber(3)).toNumber()] = _497_v61;
          _nw78[(new BigNumber(4)).toNumber()] = _497_v61;
          _498_v62 = _nw78;
          let _index79 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_498_v62).length));
          (_498_v62)[_index79] = _497_v61;
          let _index80 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
          let _index81 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_498_v62).length));
          let _rhs80 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(403)), function (_499_i1) {
            return (_this).f13;
          }), _dafny.Seq.UnicodeFromString("endmc")), _dafny.Seq.UnicodeFromString("gmwqy"));
          let _rhs81 = _411_v6.f10;
          let _rhs82 = _497_v61;
          let _lhs51 = _408_v3;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
          let _lhs53 = _498_v62;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_498_v62).length));
          _406_v1 = _rhs80;
          _lhs51[_lhs52] = _rhs81;
          _lhs53[_lhs54] = _rhs82;
          let _500_v63;
          let _nw79 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _500_v63 = _nw79;
          let _501_v64;
          _501_v64 = _dafny.Seq.of(_500_v63);
          let _502_v65;
          _502_v65 = _dafny.Map.Empty.slice().updateUnsafe(_406_v1,false);
          _500_v63 = (_501_v64)[_module.__default.safeIndex((new BigNumber((_502_v65).length)).multipliedBy(p1), new BigNumber((_501_v64).length))];
          let _503_v66;
          _503_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,_496_cf8);
          let _504_v67;
          _504_v67 = _dafny.Seq.of(_503_v66);
          _503_v66 = ((_504_v67)[_module.__default.safeIndex(p2, new BigNumber((_504_v67).length))]).update(!(_496_cf8), (_408_v3)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length))]);
          let _505_v68;
          _505_v68 = _dafny.MultiSet.fromElements(new BigNumber(168), (_this).f6, p1);
          let _506_v69;
          _506_v69 = _dafny.Map.Empty.slice().updateUnsafe(_496_cf8,_407_v2);
          let _507_v70;
          _507_v70 = _dafny.MultiSet.fromElements(false, _411_v6.f10);
          let _index82 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
          let _rhs83 = p0;
          let _rhs84 = ((((_405_v0).contains(new BigNumber((_506_v69).length))) ? ((_405_v0).get(new BigNumber((_506_v69).length))) : (_module.__default.fm1(_411_v6.f10, globalState)))).isLessThanOrEqualTo((_dafny.ZERO).minus(_this.f7));
          let _rhs85 = _dafny.MultiSet.fromElements(new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(_411_v6.f10, _411_v6.f10))).Union(_507_v70)).cardinality()));
          let _rhs86 = !(_411_v6.f10);
          let _lhs55 = _408_v3;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
          let _lhs57 = _411_v6;
          _lhs55[_lhs56] = _rhs83;
          _lhs57.f10 = _rhs84;
          _505_v68 = _rhs85;
          r0 = _rhs86;
        } else {
          let _508___mcc_h13 = (_source7).cf4;
          let _509_cf4 = _508___mcc_h13;
          let _510_v71;
          let _nw80 = new _module.C0();
          _nw80.__ctor(!(_411_v6.f10));
          _510_v71 = _nw80;
          let _511_v72;
          let _init10 = ((_512_v71, _513_v6) => function (_514_i2) {
            return _module.__default.safeModuloInt(_514_i2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_512_v71.f10,_dafny.Seq.of(_513_v6.f10))).length));
          })(_510_v71, _411_v6);
          let _nw81 = Array((new BigNumber(11)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw81.length); _i0_10++) {
            _nw81[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _511_v72 = _nw81;
          _511_v72 = _511_v72;
          let _515_v73;
          let _init11 = ((_516_v0) => function (_517_i3) {
            return _516_v0;
          })(_405_v0);
          let _nw82 = Array((new BigNumber(10)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw82.length); _i0_11++) {
            _nw82[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _515_v73 = _nw82;
          let _index83 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_515_v73).length));
          (_515_v73)[_index83] = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm1(_module.__default.fm0(_411_v6.f10, _405_v0, p2, globalState), globalState));
          let _518_v74;
          let _init12 = ((_519_p1) => function (_520_i4) {
            return _dafny.Seq.of(new BigNumber(959), _519_p1);
          })(p1);
          let _nw83 = Array((new BigNumber(22)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw83.length); _i0_12++) {
            _nw83[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _518_v74 = _nw83;
          let _index84 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_515_v73).length));
          let _rhs87 = _this.f7;
          let _rhs88 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))], new BigNumber(-575))), ((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]).plus((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]));
          let _rhs89 = (_405_v0).Merge(_405_v0);
          let _rhs90 = _module.__default.fm0(!(true) || (p0), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-267),_this.f7), (_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))], globalState);
          let _rhs91 = _518_v74;
          let _lhs58 = globalState;
          let _lhs59 = _515_v73;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_515_v73).length));
          let _lhs61 = _411_v6;
          _lhs58.f1 = _rhs87;
          r2 = _rhs88;
          _lhs59[_lhs60] = _rhs89;
          _lhs61.f10 = _rhs90;
          _518_v74 = _rhs91;
          (_510_v71).f10 = false;
        }
        if (p0) {
          let _index85 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length));
          (_408_v3)[_index85] = ((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]).isEqualTo((p2).plus((_dafny.ZERO).minus((_this).f6)));
          (globalState).f1 = new BigNumber(846);
          let _521_v75;
          _521_v75 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_479_v51),!(_411_v6.f10));
          let _522_v76;
          _522_v76 = _module.D1.create_DC2(_521_v75);
          let _523_v78;
          _523_v78 = _module.D0.create_DC0(_dafny.Set.fromElements((_408_v3)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length))]));
          let _pat_let_tv7 = _479_v51;
          let _pat_let_tv8 = _479_v51;
          let _pat_let_tv9 = _521_v75;
          let _524_v79;
          let _nw84 = Array((new BigNumber(23)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = _522_v76;
          _nw84[(_dafny.ONE).toNumber()] = _522_v76;
          _nw84[(new BigNumber(2)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(3)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(4)).toNumber()] = _module.D1.create_DC2(function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (_dafny.Set.fromElements(_523_v78, function (_pat_let11_0) {
    return function (_525_dt__update__tmp_h1) {
      return function (_pat_let12_0) {
        return function (_526_dt__update_hcf0_h0) {
          return _module.D0.create_DC0(_526_dt__update_hcf0_h0);
        }(_pat_let12_0);
      }(_pat_let_tv7);
    }(_pat_let11_0);
  }(_module.D0.create_DC0(_479_v51)))).Elements) {
    let _527_v77 = _compr_18;
    if ((_dafny.Set.fromElements(_523_v78, function (_pat_let13_0) {
      return function (_528_dt__update__tmp_h1) {
        return function (_pat_let14_0) {
          return function (_529_dt__update_hcf0_h0) {
            return _module.D0.create_DC0(_529_dt__update_hcf0_h0);
          }(_pat_let14_0);
        }(_pat_let_tv8);
      }(_pat_let13_0);
    }(_module.D0.create_DC0(_479_v51)))).contains(_527_v77)) {
      _coll18.push([_527_v77,true]);
    }
  }
  return _coll18;
}());
          _nw84[(new BigNumber(5)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(6)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(7)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(8)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(9)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(10)).toNumber()] = function (_pat_let15_0) {
            return function (_530_dt__update__tmp_h2) {
              return function (_pat_let16_0) {
                return function (_531_dt__update_hcf4_h0) {
                  return _module.D1.create_DC2(_531_dt__update_hcf4_h0);
                }(_pat_let16_0);
              }(_pat_let_tv9);
            }(_pat_let15_0);
          }(_module.D1.create_DC2(_521_v75));
          _nw84[(new BigNumber(11)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(12)).toNumber()] = _module.D1.create_DC2(_521_v75);
          _nw84[(new BigNumber(13)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(14)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(15)).toNumber()] = _module.D1.create_DC2(_521_v75);
          _nw84[(new BigNumber(16)).toNumber()] = _module.D1.create_DC2(_521_v75);
          _nw84[(new BigNumber(17)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(18)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(19)).toNumber()] = _module.D1.create_DC2(_521_v75);
          _nw84[(new BigNumber(20)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(21)).toNumber()] = _522_v76;
          _nw84[(new BigNumber(22)).toNumber()] = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_523_v78,false));
          _524_v79 = _nw84;
          let _532_v80;
          _532_v80 = _module.D6.create_DC13(_524_v79);
          let _533_v81;
          let _nw85 = Array((new BigNumber(23)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _524_v79;
          _nw85[(_dafny.ONE).toNumber()] = _524_v79;
          _nw85[(new BigNumber(2)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(3)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(4)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(5)).toNumber()] = (_532_v80).dtor_cf22;
          _nw85[(new BigNumber(6)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(7)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(8)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(9)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(10)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(11)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(12)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(13)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(14)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(15)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(16)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(17)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(18)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(19)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(20)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(21)).toNumber()] = _524_v79;
          _nw85[(new BigNumber(22)).toNumber()] = _524_v79;
          _533_v81 = _nw85;
          let _index86 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_533_v81).length));
          (_533_v81)[_index86] = _524_v79;
          let _534_v82;
          _534_v82 = _dafny.MultiSet.fromElements(true, _411_v6.f10, _411_v6.f10, !(p0), (_408_v3)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length))]);
          let _535_v83;
          _535_v83 = _dafny.Map.Empty.slice().updateUnsafe(((_module.__default.fm0(p0, _405_v0, _this.f7, globalState)) ? (_411_v6.f10) : ((_408_v3)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length))])),_524_v79);
          let _index87 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_533_v81).length));
          let _rhs92 = _module.__default.safeModuloInt((_this.f7).plus((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]), new BigNumber((_534_v82).cardinality()));
          let _rhs93 = (((_535_v83).contains(_411_v6.f10)) ? ((_535_v83).get(_411_v6.f10)) : (_524_v79));
          let _lhs62 = globalState;
          let _lhs63 = _533_v81;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_533_v81).length));
          _lhs62.f1 = _rhs92;
          _lhs63[_lhs64] = _rhs93;
          let _536_v84;
          let _nw86 = new _module.C0();
          _nw86.__ctor(p0);
          _536_v84 = _nw86;
          let _537_v85;
          let _nw87 = Array((new BigNumber(9)).toNumber()).fill(_module.D0.Default());
          _537_v85 = _nw87;
          let _538_v86;
          _538_v86 = _dafny.Seq.of(_537_v85, ((p0) ? (_537_v85) : (_537_v85)));
          _537_v85 = (_538_v86)[_module.__default.safeIndex((_this.f7).minus(_this.f7), new BigNumber((_538_v86).length))];
        } else {
          let _index88 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length));
          (_413_v8)[_index88] = (_dafny.ZERO).minus((_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]);
          _406_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_406_v1, _406_v1), _dafny.Seq.Concat(_406_v1, _dafny.Seq.UnicodeFromString("lvmvk")));
          (_411_v6).f10 = _411_v6.f10;
          let _539_v88;
          _539_v88 = _dafny.Set.fromElements(_406_v1);
          _479_v51 = (_module.__default.fm3(p0, p1, _479_v51, function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_539_v88).Elements) {
              let _540_v87 = _compr_19;
              if ((_539_v88).contains(_540_v87)) {
                _coll19.push([_540_v87,(_413_v8)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_413_v8).length))]]);
              }
            }
            return _coll19;
          }(), globalState)).dtor_cf0;
          (_411_v6).f10 = !(!((_408_v3)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_408_v3).length))]) || (_411_v6.f10));
        }
      }
      r1 = _module.__default.fm0(p0, (_405_v0).Merge(_405_v0), (_this).f6, globalState);
      r0 = true;
      r1 = !(_dafny.Map.Empty.slice().updateUnsafe(_408_v3,_this.f7)).equals(_dafny.Map.Empty.slice().updateUnsafe(_408_v3,p2));
      r2 = _module.__default.fm1(_411_v6.f10, globalState);
      return [r0, r1, r2];
    }
    m10(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _module.D1.Default();
      let r3 = false;
      let _541_v0;
      _541_v0 = false;
      let _542_v1;
      _542_v1 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,new BigNumber(654));
      let _hi4 = new BigNumber(((_542_v1).Merge(_542_v1)).length);
      for (let _543_i0 = (_module.__default.fm1(_541_v0, globalState)).multipliedBy(new BigNumber(-878)); _543_i0.isLessThan(_hi4); _543_i0 = _543_i0.plus(_dafny.ONE)) {
        let _544_v2;
        _544_v2 = _dafny.Set.fromElements(_543_i0);
        let _545_v3;
        _545_v3 = _dafny.Set.fromElements(new BigNumber(612), new BigNumber((_544_v2).length), (_this).f6);
        let _546_v4;
        _546_v4 = _dafny.Seq.UnicodeFromString("twoo");
        let _547_v5;
        _547_v5 = _dafny.Seq.of(_541_v0, _541_v0);
        let _548_v6;
        _548_v6 = _dafny.Set.fromElements(_541_v0);
        let _549_v7;
        _549_v7 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_546_v4);
        let _550_v8;
        let _nw88 = Array((new BigNumber(17)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = new BigNumber((_545_v3).length);
        _nw88[(_dafny.ONE).toNumber()] = p0;
        _nw88[(new BigNumber(2)).toNumber()] = (_this).f6;
        _nw88[(new BigNumber(3)).toNumber()] = p0;
        _nw88[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw88[(new BigNumber(5)).toNumber()] = _543_i0;
        _nw88[(new BigNumber(6)).toNumber()] = new BigNumber((_546_v4).length);
        _nw88[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_547_v5)).cardinality()));
        _nw88[(new BigNumber(8)).toNumber()] = new BigNumber((_548_v6).length);
        _nw88[(new BigNumber(9)).toNumber()] = _543_i0;
        _nw88[(new BigNumber(10)).toNumber()] = _543_i0;
        _nw88[(new BigNumber(11)).toNumber()] = p0;
        _nw88[(new BigNumber(12)).toNumber()] = new BigNumber(502);
        _nw88[(new BigNumber(13)).toNumber()] = _543_i0;
        _nw88[(new BigNumber(14)).toNumber()] = new BigNumber((_547_v5).length);
        _nw88[(new BigNumber(15)).toNumber()] = _543_i0;
        _nw88[(new BigNumber(16)).toNumber()] = new BigNumber(((((_549_v7).contains(_541_v0)) ? ((_549_v7).get(_541_v0)) : (_546_v4))).length);
        _550_v8 = _nw88;
        let _551_v9;
        _551_v9 = _dafny.Seq.of(_543_i0, (_dafny.ZERO).minus(_543_i0));
        let _552_v10;
        _552_v10 = _dafny.Map.Empty.slice().updateUnsafe(_551_v9,_543_i0);
        let _553_v11;
        _553_v11 = _dafny.Map.Empty.slice().updateUnsafe(_550_v8,new BigNumber((_552_v10).length));
        let _554_v12;
        _554_v12 = _dafny.Map.Empty.slice().updateUnsafe(_543_i0,_543_i0);
        _553_v11 = (_553_v11).update(((_module.__default.fm0(!(_541_v0), _554_v12, new BigNumber(432), globalState)) ? (_550_v8) : (_550_v8)), (_this).f6);
        r3 = _541_v0;
        _548_v6 = _548_v6;
        let _555_v13;
        let _nw89 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _555_v13 = _nw89;
        let _556_v14;
        _556_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p0);
        let _index89 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_555_v13).length));
        (_555_v13)[_index89] = _556_v14;
        let _557_v15;
        _557_v15 = _module.D7.create_DC16(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,p0));
        let _index90 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_555_v13).length));
        (_555_v13)[_index90] = ((_541_v0) ? (_556_v14) : ((_557_v15).dtor_cf30));
      }
      let _558_v16;
      let _init13 = function (_559_i1) {
        return _module.__default.safeDivisionInt(_559_i1, (_this).f6);
      };
      let _nw90 = Array((new BigNumber(12)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw90.length); _i0_13++) {
        _nw90[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _558_v16 = _nw90;
      let _560_v17;
      _560_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f7);
      let _index91 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      (_558_v16)[_index91] = (((_560_v17).contains(_this.f7)) ? ((_560_v17).get(_this.f7)) : ((_this).f6));
      let _561_v18;
      _561_v18 = _dafny.MultiSet.fromElements(_541_v0);
      let _index92 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      let _rhs94 = p0;
      let _rhs95 = (_561_v18).IsSubsetOf(_561_v18);
      let _lhs65 = _558_v16;
      let _lhs66 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      _lhs65[_lhs66] = _rhs94;
      r3 = _rhs95;
      let _562_v19;
      _562_v19 = _dafny.Seq.of(((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]).minus((_this).f6));
      let _index93 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      let _index94 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      let _index95 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      let _rhs96 = (_this).f6;
      let _rhs97 = (((_541_v0) ? ((_this).f6) : ((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]))).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f7, (_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))])));
      let _rhs98 = (_562_v19)[_module.__default.safeIndex(_module.__default.safeModuloInt((_this).f6, p0), new BigNumber((_562_v19).length))];
      let _lhs67 = _558_v16;
      let _lhs68 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      let _lhs69 = _558_v16;
      let _lhs70 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      let _lhs71 = _558_v16;
      let _lhs72 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
      _lhs67[_lhs68] = _rhs96;
      _lhs69[_lhs70] = _rhs97;
      _lhs71[_lhs72] = _rhs98;
      let _hi5 = p0;
      for (let _563_i2 = _module.__default.fm1(_541_v0, globalState); _563_i2.isLessThan(_hi5); _563_i2 = _563_i2.plus(_dafny.ONE)) {
        let _564_v20;
        _564_v20 = _dafny.Set.fromElements(!(_563_i2).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_563_i2))), (_541_v0) === (_541_v0), _541_v0, _module.__default.fm0(_541_v0, _560_v17, (_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))], globalState), false);
        _564_v20 = _564_v20;
        let _565_v21;
        _565_v21 = _dafny.Map.Empty.slice().updateUnsafe(_562_v19,_this.f7);
        let _index96 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        let _rhs99 = (_module.__default.fm29(globalState)).Merge((_565_v21).Merge(_565_v21));
        let _rhs100 = (_dafny.ZERO).minus((p0).minus((p0).plus(new BigNumber(410))));
        let _rhs101 = _dafny.Seq.update(_dafny.Seq.Concat(_562_v19, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-878)), function (_566_i3) {
          return (_dafny.ZERO).minus(new BigNumber(-628));
        }), _562_v19)), _module.__default.safeIndex((_module.D1.create_DC3(_541_v0, new BigNumber((_542_v1).length), (_this).f6)).dtor_cf6, new BigNumber((_dafny.Seq.Concat(_562_v19, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-878)), function (_567_i3) {
          return (_dafny.ZERO).minus(new BigNumber(-628));
        }), _562_v19))).length)), (_this).f6);
        let _lhs73 = _558_v16;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        _565_v21 = _rhs99;
        _lhs73[_lhs74] = _rhs100;
        _562_v19 = _rhs101;
        let _568_v22;
        let _nw91 = Array((new BigNumber(5)).toNumber()).fill(false);
        _568_v22 = _nw91;
        _568_v22 = _568_v22;
        let _569_v23;
        let _init14 = ((_570_v17) => function (_571_i4) {
          return _570_v17;
        })(_560_v17);
        let _nw92 = Array((new BigNumber(2)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw92.length); _i0_14++) {
          _nw92[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _569_v23 = _nw92;
        let _572_v24;
        _572_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_541_v0, globalState),_569_v23);
        _572_v24 = (_572_v24).update(_module.__default.safeModuloInt((_this).f6, (_dafny.ZERO).minus(_563_i2)), _569_v23);
      }
      let _hi6 = _this.f7;
      for (let _573_i5 = _module.__default.safeModuloInt(p0, new BigNumber((_560_v17).length)); _573_i5.isLessThan(_hi6); _573_i5 = _573_i5.plus(_dafny.ONE)) {
        let _index97 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        (_558_v16)[_index97] = _module.__default.safeDivisionInt(new BigNumber(-363), _573_i5);
        (globalState).f1 = p0;
        let _index98 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        (_558_v16)[_index98] = _this.f7;
        let _574_v25;
        let _nw93 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _574_v25 = _nw93;
        let _index99 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_574_v25).length));
        (_574_v25)[_index99] = (_this).f13;
        let _index100 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_574_v25).length));
        (_574_v25)[_index100] = (_this).f13;
      }
      let _575_v26;
      _575_v26 = _dafny.Set.fromElements(_this.f7, (_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]);
      if ((_dafny.Set.fromElements(new BigNumber(-627))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_562_v19).length), new BigNumber((_575_v26).length), new BigNumber(208)))) {
        let _576_v27;
        let _nw94 = new _module.C0();
        _nw94.__ctor(((true) ? (_541_v0) : (_541_v0)));
        _576_v27 = _nw94;
        _541_v0 = _541_v0;
        let _577_v29;
        _577_v29 = _dafny.Seq.of(_541_v0, _541_v0, _576_v27.f10);
        let _578_v30;
        _578_v30 = _dafny.Seq.of(_577_v29, _577_v29);
        (_this).f7 = ((!(_576_v27.f10)) ? ((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]) : (((_576_v27.f10) ? (new BigNumber((function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of (_578_v30).Elements) {
            let _579_v28 = _compr_20;
            if (_dafny.Seq.contains(_578_v30, _579_v28)) {
              _coll20.push([_579_v28,p0]);
            }
          }
          return _coll20;
        }()).length)) : ((_this).f6))));
        (_576_v27).f10 = _541_v0;
        let _580_v31;
        _580_v31 = _dafny.Seq.UnicodeFromString("cykvhpkly");
        let _581_v32;
        _581_v32 = _dafny.MultiSet.fromElements((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))], (_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]);
        let _index101 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        let _rhs102 = ((_576_v27.f10) ? ((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]) : (new BigNumber(200)));
        let _rhs103 = _558_v16;
        let _rhs104 = _module.__default.safeModuloInt((new BigNumber((_580_v31).length)).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_581_v32).cardinality()), new BigNumber(-162), new BigNumber(-131), new BigNumber(355))).length)), ((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))]).minus(_this.f7));
        let _rhs105 = (_this.f7).minus(_this.f7);
        let _rhs106 = _dafny.Seq.Concat(_577_v29, _dafny.Seq.Concat(_577_v29, _577_v29));
        let _lhs75 = globalState;
        let _lhs76 = _558_v16;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        let _lhs78 = globalState;
        _lhs75.f1 = _rhs102;
        _558_v16 = _rhs103;
        _lhs76[_lhs77] = _rhs104;
        _lhs78.f1 = _rhs105;
        _577_v29 = _rhs106;
      } else {
        if (_541_v0) {
          let _582_v33;
          _582_v33 = new _dafny.CodePoint('i'.codePointAt(0));
          _582_v33 = _582_v33;
          let _583_v34;
          let _nw95 = new _module.C0();
          _nw95.__ctor(_541_v0);
          _583_v34 = _nw95;
          let _584_v35;
          _584_v35 = _dafny.Set.fromElements(_558_v16, _558_v16, _558_v16, _558_v16);
          let _rhs107 = _584_v35;
          let _rhs108 = _module.__default.fm1((_dafny.MultiSet.FromArray(_562_v19)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_562_v19)), globalState);
          let _lhs79 = globalState;
          _584_v35 = _rhs107;
          _lhs79.f1 = _rhs108;
          (_583_v34).f10 = _541_v0;
          let _585_v36;
          _585_v36 = _dafny.Map.Empty.slice().updateUnsafe((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))],_558_v16);
          _585_v36 = (_585_v36).update((_this).f6, _558_v16);
        } else {
          _541_v0 = !((_541_v0) === ((false) || (_541_v0)));
          _541_v0 = ((_541_v0) ? (_541_v0) : (_541_v0));
          let _586_v37;
          _586_v37 = _module.D2.create_DC5(_558_v16);
          let _587_v38;
          _587_v38 = _dafny.Seq.of(_558_v16, _558_v16, (_586_v37).dtor_cf9, _558_v16, _558_v16);
          _558_v16 = ((_541_v0) ? (_558_v16) : ((_587_v38)[_module.__default.safeIndex(_this.f7, new BigNumber((_587_v38).length))]));
          _541_v0 = !(_541_v0);
          let _588_v39;
          let _nw96 = Array((new BigNumber(9)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _541_v0;
          _nw96[(_dafny.ONE).toNumber()] = false;
          _nw96[(new BigNumber(2)).toNumber()] = _541_v0;
          _nw96[(new BigNumber(3)).toNumber()] = _541_v0;
          _nw96[(new BigNumber(4)).toNumber()] = _541_v0;
          _nw96[(new BigNumber(5)).toNumber()] = true;
          _nw96[(new BigNumber(6)).toNumber()] = _541_v0;
          _nw96[(new BigNumber(7)).toNumber()] = _541_v0;
          _nw96[(new BigNumber(8)).toNumber()] = _541_v0;
          _588_v39 = _nw96;
          let _589_v40;
          _589_v40 = _module.D0.create_DC1(_541_v0, _588_v39, false);
          let _590_v41;
          _590_v41 = _dafny.Seq.of(_589_v40, _589_v40, _589_v40);
          let _591_v42;
          _591_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(true, globalState),!(!(_541_v0)));
          let _592_v43;
          _592_v43 = _dafny.Seq.of((_591_v42).Merge(_dafny.Map.Empty.slice().updateUnsafe((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))],true)), _591_v42, _591_v42);
          let _rhs109 = _dafny.Seq.Concat(_590_v41, _590_v41);
          let _rhs110 = _592_v43;
          let _rhs111 = ((new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_541_v0)).update(_this.f7, _541_v0)).length)).multipliedBy(p0)).multipliedBy(new BigNumber(364));
          let _rhs112 = _562_v19;
          let _lhs80 = globalState;
          _590_v41 = _rhs109;
          _592_v43 = _rhs110;
          _lhs80.f1 = _rhs111;
          _562_v19 = _rhs112;
        }
        let _index102 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length));
        (_558_v16)[_index102] = (_this).f6;
        (_this).f7 = ((_541_v0) ? (_module.__default.fm1(_541_v0, globalState)) : ((_dafny.ZERO).minus((_this).f6)));
        let _593_v44;
        let _init15 = ((_594_v0) => function (_595_i6) {
          return _594_v0;
        })(_541_v0);
        let _nw97 = Array((new BigNumber(20)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw97.length); _i0_15++) {
          _nw97[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _593_v44 = _nw97;
        let _596_v45;
        _596_v45 = _module.D0.create_DC1(_541_v0, _593_v44, _541_v0);
        let _597_v46;
        let _598_v47;
        let _599_v48;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector4 = _module.__default.m0(_596_v45, new BigNumber((_dafny.MultiSet.fromElements(_558_v16)).cardinality()), globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _597_v46 = _out12;
        _598_v47 = _out13;
        _599_v48 = _out14;
        let _600_v49;
        _600_v49 = _dafny.Map.Empty.slice().updateUnsafe(_560_v17,_this.f7);
        let _601_v51;
        _601_v51 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_541_v0);
        let _602_v52;
        _602_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,_601_v51);
        (globalState).f1 = new BigNumber(((_600_v49).Merge((_600_v49).update(function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_602_v52).Keys.Elements) {
            let _603_v50 = _compr_21;
            if ((_602_v52).contains(_603_v50)) {
              _coll21.push([(_603_v50).multipliedBy(p0),_this.f7]);
            }
          }
          return _coll21;
        }(), (_this).f6))).length);
      }
      r0 = _561_v18;
      r1 = p0;
      let _604_v54;
      _604_v54 = _dafny.MultiSet.fromElements(_575_v26);
      let _605_v55;
      _605_v55 = _dafny.Seq.of(_604_v54);
      let _606_v56;
      _606_v56 = _module.D1.create_DC4((_this.f7).isLessThanOrEqualTo(new BigNumber((function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of ((_605_v55)[_module.__default.safeIndex((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))], new BigNumber((_605_v55).length))]).Elements) {
    let _607_v53 = _compr_22;
    if (((_605_v55)[_module.__default.safeIndex((_558_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_558_v16).length))], new BigNumber((_605_v55).length))]).contains(_607_v53)) {
      _coll22.push([_607_v53,_541_v0]);
    }
  }
  return _coll22;
}()).length)));
      r2 = _606_v56;
      r3 = _541_v0;
      return [r0, r1, r2, r3];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f7 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6, f7) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm23(p0, p1, p2, globalState) {
      let _this = this;
      if (true) {
        return _module.D1.create_DC4(true);
      } else {
        return _module.D1.create_DC4(true);
      }
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _608_v0;
      _608_v0 = _module.D1.create_DC3(p0, p1, _this.f7);
      if ((((p0) ? (_608_v0) : (_608_v0))).dtor_cf5) {
        r1 = p0;
        if (p0) {
          let _609_v1;
          _609_v1 = _dafny.Seq.UnicodeFromString("ignf");
          let _610_v2;
          let _nw98 = new _module.C0();
          _nw98.__ctor(p0);
          _610_v2 = _nw98;
          let _611_v3;
          _611_v3 = _dafny.Set.fromElements(_610_v2, _610_v2, _610_v2);
          let _612_v4;
          _612_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_611_v3).length),_610_v2.f10);
          let _613_v5;
          _613_v5 = _module.D1.create_DC4(_610_v2.f10);
          let _614_v6;
          _614_v6 = _module.D2.create_DC7(p0, (_dafny.ZERO).minus(_this.f7), _612_v4, _613_v5);
          let _615_v7;
          _615_v7 = new _dafny.CodePoint('x'.codePointAt(0));
          let _616_v8;
          _616_v8 = _dafny.Map.Empty.slice().updateUnsafe(_615_v7,true);
          let _617_v9;
          _617_v9 = _dafny.Set.fromElements(_610_v2.f10, !((_module.D2.create_DC7(p0, _this.f7, _612_v4, _613_v5)).dtor_cf13));
          let _618_v11;
          _618_v11 = _dafny.Map.Empty.slice().updateUnsafe(_609_v1,_615_v7);
          let _pat_let_tv10 = _617_v9;
          let _619_v12;
          _619_v12 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let17_0) {
            return function (_621_dt__update__tmp_h0) {
              return function (_pat_let18_0) {
                return function (_622_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_622_dt__update_hcf0_h0);
                }(_pat_let18_0);
              }(_pat_let_tv10);
            }(_pat_let17_0);
          }(_module.__default.fm3((_614_v6).dtor_cf13, new BigNumber((_616_v8).length), _617_v9, function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_618_v11).Keys.Elements) {
              let _620_v10 = _compr_23;
              if ((_618_v11).contains(_620_v10)) {
                _coll23.push([_620_v10,_this.f7]);
              }
            }
            return _coll23;
          }(), globalState)),_610_v2.f10);
          let _623_v13;
          _623_v13 = _module.D1.create_DC2((_619_v12).Merge(_619_v12));
          let _rhs113 = _module.__default.fm1(_610_v2.f10, globalState);
          let _rhs114 = (_this).f6;
          let _rhs115 = _609_v1;
          let _rhs116 = _609_v1;
          let _rhs117 = _module.D1.create_DC2(_619_v12);
          let _lhs81 = _this;
          _lhs81.f7 = _rhs113;
          r2 = _rhs114;
          _609_v1 = _rhs115;
          _609_v1 = _rhs116;
          _623_v13 = _rhs117;
          let _624_v14;
          _624_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(398),_610_v2);
          let _625_v15;
          _625_v15 = _dafny.Set.fromElements(_624_v14, _dafny.Map.Empty.slice().updateUnsafe(p1,_610_v2));
          let _626_v16;
          let _nw99 = new _module.C0();
          _nw99.__ctor((_625_v15).IsSubsetOf(_dafny.Set.fromElements(_624_v14)));
          _626_v16 = _nw99;
          let _627_v17;
          _627_v17 = _dafny.Seq.of(_626_v16.f10, _610_v2.f10);
          _627_v17 = _dafny.Seq.Concat(_627_v17, _627_v17);
          let _628_v18;
          let _nw100 = new _module.C0();
          _nw100.__ctor(p0);
          _628_v18 = _nw100;
          let _629_v19;
          let _init16 = function (_630_i0) {
            return (_630_i0).multipliedBy(new BigNumber(923));
          };
          let _nw101 = Array((new BigNumber(10)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw101.length); _i0_16++) {
            _nw101[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _629_v19 = _nw101;
          let _index103 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_629_v19).length));
          (_629_v19)[_index103] = p2;
          let _631_v20;
          let _init17 = ((_632_v7) => function (_633_i1) {
            return _632_v7;
          })(_615_v7);
          let _nw102 = Array((new BigNumber(20)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw102.length); _i0_17++) {
            _nw102[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _631_v20 = _nw102;
          let _index104 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_631_v20).length));
          (_631_v20)[_index104] = _615_v7;
          let _index105 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_629_v19).length));
          let _index106 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_631_v20).length));
          let _rhs118 = _module.__default.safeModuloInt((_this).f6, ((false) ? (_this.f7) : (_this.f7)));
          let _rhs119 = _609_v1;
          let _rhs120 = _615_v7;
          let _rhs121 = _610_v2.f10;
          let _lhs82 = _629_v19;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_629_v19).length));
          let _lhs84 = _631_v20;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_631_v20).length));
          let _lhs86 = _610_v2;
          _lhs82[_lhs83] = _rhs118;
          _609_v1 = _rhs119;
          _lhs84[_lhs85] = _rhs120;
          _lhs86.f10 = _rhs121;
        } else {
          let _634_v21;
          let _nw103 = Array((new BigNumber(10)).toNumber()).fill(false);
          _634_v21 = _nw103;
          let _635_v22;
          _635_v22 = _dafny.Seq.of(p1);
          let _636_v23;
          _636_v23 = _dafny.Map.Empty.slice().updateUnsafe(_635_v22,p0);
          let _rhs122 = _634_v21;
          let _rhs123 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), function (_637_i2) {
            return _this.f7;
          }),p0);
          let _rhs124 = _634_v21;
          _634_v21 = _rhs122;
          _636_v23 = _rhs123;
          _634_v21 = _rhs124;
          let _638_v24;
          _638_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _639_v25;
          _639_v25 = _dafny.Seq.UnicodeFromString("yujxjkyll");
          let _640_v26;
          _640_v26 = _dafny.Seq.of(false);
          (_this).m8(_638_v24, _module.D2.create_DC6(_module.__default.fm24(p0, _module.D1.create_DC4(p0), _this.f7, globalState), _639_v25, new BigNumber((_dafny.Seq.update(_640_v26, _module.__default.safeIndex((_this).f6, new BigNumber((_640_v26).length)), !(p0))).length)), _639_v25, globalState);
          (globalState).f1 = new BigNumber(-712);
          r0 = false;
          _638_v24 = (_638_v24).update(p0, p0);
        }
        let _641_v27;
        let _nw104 = Array((new BigNumber(3)).toNumber()).fill(false);
        _641_v27 = _nw104;
        _641_v27 = _641_v27;
        (globalState).f1 = (((_this).f6).plus(p1)).multipliedBy((p2).minus((_this).f6));
        let _642_v28;
        let _init18 = function (_643_i3) {
          return (_643_i3).multipliedBy((_this).f6);
        };
        let _nw105 = Array((new BigNumber(2)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw105.length); _i0_18++) {
          _nw105[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _642_v28 = _nw105;
        let _644_v29;
        _644_v29 = _dafny.Set.fromElements(p1, (_this).f6, _this.f7);
        let _rhs125 = _642_v28;
        let _rhs126 = _642_v28;
        let _rhs127 = _642_v28;
        let _rhs128 = _644_v29;
        _642_v28 = _rhs125;
        _642_v28 = _rhs126;
        _642_v28 = _rhs127;
        _644_v29 = _rhs128;
      } else {
        r0 = p0;
        let _645_v30;
        _645_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Set.fromElements(p0, p0, p0, p0));
        let _646_v31;
        _646_v31 = _dafny.Set.fromElements(true);
        _645_v30 = (_645_v30).update(new BigNumber(138), (_646_v31).Difference(_646_v31));
        let _647_v32;
        let _init19 = function (_648_i4) {
          return _dafny.Seq.UnicodeFromString("yqtii");
        };
        let _nw106 = Array((new BigNumber(23)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw106.length); _i0_19++) {
          _nw106[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _647_v32 = _nw106;
        _647_v32 = _647_v32;
        let _649_v33;
        _649_v33 = _module.D1.create_DC4(!(!(p0) || (p0)));
        let _source8 = _649_v33;
        if (_source8.is_DC3) {
          let _650___mcc_h0 = (_source8).cf5;
          let _651___mcc_h1 = (_source8).cf6;
          let _652___mcc_h2 = (_source8).cf7;
          let _653_cf7 = _652___mcc_h2;
          let _654_cf6 = _651___mcc_h1;
          let _655_cf5 = _650___mcc_h0;
          let _656_v34;
          _656_v34 = _dafny.Seq.of(_655_cf5, p0);
          let _657_v35;
          _657_v35 = _dafny.MultiSet.fromElements(_656_v34, _656_v34, _656_v34, _dafny.Seq.of(_655_cf5), _dafny.Seq.of(_655_cf5));
          let _658_v36;
          _658_v36 = _dafny.Map.Empty.slice().updateUnsafe(!((new BigNumber(297)).isLessThan((_dafny.ZERO).minus(p2))),_module.__default.safeDivisionInt(_654_cf6, (((_657_v35).contains(_dafny.Seq.of(p0))) ? ((_657_v35).get(_dafny.Seq.of(p0))) : ((_this).f6))));
          _658_v36 = (_658_v36).update(false, (_dafny.ZERO).minus((_this.f7).minus(_this.f7)));
          let _659_v37;
          let _nw107 = new _module.C0();
          _nw107.__ctor(_655_cf5);
          _659_v37 = _nw107;
          let _660_v38;
          _660_v38 = _dafny.Seq.UnicodeFromString("nbw");
          let _rhs129 = _dafny.Seq.UnicodeFromString("tttqannns");
          let _rhs130 = _660_v38;
          let _rhs131 = _656_v34;
          _660_v38 = _rhs129;
          _660_v38 = _rhs130;
          _656_v34 = _rhs131;
          let _661_v39;
          let _nw108 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _661_v39 = _nw108;
          let _index107 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_661_v39).length));
          (_661_v39)[_index107] = _653_cf7;
          let _index108 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_661_v39).length));
          (_661_v39)[_index108] = p1;
        } else if (_source8.is_DC4) {
          let _662___mcc_h3 = (_source8).cf8;
          let _663_cf8 = _662___mcc_h3;
          let _664_v40;
          _664_v40 = _dafny.Seq.of(_module.__default.safeDivisionInt(_this.f7, p2), _this.f7, (_this).f6);
          _664_v40 = _664_v40;
          let _665_v41;
          _665_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p0, globalState),(true) || (p0));
          let _666_v42;
          _666_v42 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p2);
          _665_v41 = (_665_v41).update((_this).f6, _module.__default.fm0(_663_cf8, _666_v42, (_this).f6, globalState));
          let _667_v43;
          let _init20 = ((_668_cf8) => function (_669_i5) {
            return _668_cf8;
          })(_663_cf8);
          let _nw109 = Array((new BigNumber(23)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw109.length); _i0_20++) {
            _nw109[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _667_v43 = _nw109;
          let _index109 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_667_v43).length));
          (_667_v43)[_index109] = p0;
          let _index110 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_667_v43).length));
          (_667_v43)[_index110] = p0;
          let _670_v44;
          _670_v44 = _dafny.Map.Empty.slice().updateUnsafe((_667_v43)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_667_v43).length))],p0);
          let _671_v45;
          _671_v45 = _module.D0.create_DC0(_646_v31);
          let _672_v46;
          _672_v46 = _dafny.Map.Empty.slice().updateUnsafe(_671_v45,(_667_v43)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_667_v43).length))]);
          let _673_v47;
          _673_v47 = _module.D1.create_DC2(_672_v46);
          let _674_v48;
          _674_v48 = _dafny.Seq.UnicodeFromString("vpeiqsqiw");
          let _675_v49;
          _675_v49 = _module.D2.create_DC6(_673_v47, _674_v48, _module.__default.fm1(_module.__default.fm0(_663_cf8, _666_v42, _this.f7, globalState), globalState));
          (_this).m8(_670_v44, _675_v49, _dafny.Seq.Concat(_674_v48, _dafny.Seq.UnicodeFromString("sibwdulqv")), globalState);
        } else {
          let _676___mcc_h4 = (_source8).cf4;
          let _677_cf4 = _676___mcc_h4;
          let _678_v50;
          let _nw110 = Array((new BigNumber(19)).toNumber()).fill(_module.D0.Default());
          _678_v50 = _nw110;
          let _679_v51;
          let _nw111 = Array((_dafny.ONE).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = p0;
          _679_v51 = _nw111;
          let _680_v52;
          _680_v52 = _dafny.Set.fromElements(_679_v51);
          let _681_v53;
          let _nw112 = Array((new BigNumber(20)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = _680_v52;
          _nw112[(_dafny.ONE).toNumber()] = _680_v52;
          _nw112[(new BigNumber(2)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(3)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(4)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(5)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_679_v51);
          _nw112[(new BigNumber(7)).toNumber()] = (_680_v52).Difference(_680_v52);
          _nw112[(new BigNumber(8)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(9)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(10)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_679_v51);
          _nw112[(new BigNumber(12)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(13)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(14)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(15)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(16)).toNumber()] = _680_v52;
          _nw112[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(_679_v51, _679_v51);
          _nw112[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements(_679_v51, _679_v51);
          _nw112[(new BigNumber(19)).toNumber()] = (_680_v52).Union(_680_v52);
          _681_v53 = _nw112;
          let _index111 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_681_v53).length));
          (_681_v53)[_index111] = _680_v52;
          let _682_v54;
          _682_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _index112 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_681_v53).length));
          let _rhs132 = _678_v50;
          let _rhs133 = _646_v31;
          let _rhs134 = (((_682_v54).update(p0, p0)).Merge(_682_v54)).equals((_682_v54).Merge(_682_v54));
          let _rhs135 = (((true) || (p0)) ? (_680_v52) : (_dafny.Set.fromElements(_679_v51)));
          let _lhs87 = _681_v53;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_681_v53).length));
          _678_v50 = _rhs132;
          _646_v31 = _rhs133;
          r1 = _rhs134;
          _lhs87[_lhs88] = _rhs135;
          r0 = p0;
          let _683_v55;
          let _init21 = ((_684_p2) => function (_685_i6) {
            return (_685_i6).plus(_684_p2);
          })(p2);
          let _nw113 = Array((new BigNumber(8)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw113.length); _i0_21++) {
            _nw113[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _683_v55 = _nw113;
          let _686_v56;
          _686_v56 = _module.D2.create_DC5(_683_v55);
          let _687_v57;
          _687_v57 = _module.D2.create_DC8(_module.D2.create_DC5((_686_v56).dtor_cf9));
          let _688_v58;
          _688_v58 = _dafny.Map.Empty.slice().updateUnsafe(_687_v57,_module.__default.fm1(p0, globalState));
          _688_v58 = (_688_v58).update(_687_v57, _module.__default.safeDivisionInt(_this.f7, (_this).f6));
          (globalState).f1 = _this.f7;
        }
        let _689_v59;
        let _nw114 = new _module.C0();
        _nw114.__ctor((p1).isEqualTo((_this).f6));
        _689_v59 = _nw114;
      }
      r1 = (_module.__default.fm1(true, globalState)).isLessThan((_this).f6);
      let _690_v60;
      _690_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _691_v61;
      _691_v61 = _module.D1.create_DC4(p0);
      let _692_v62;
      _692_v62 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_691_v61).dtor_cf8);
      let _hi7 = (_module.D2.create_DC7(p0, (_this).f6, _692_v62, _691_v61)).dtor_cf14;
      for (let _693_i7 = (_dafny.ZERO).minus((_this.f7).plus(new BigNumber((_690_v60).length))); _693_i7.isLessThan(_hi7); _693_i7 = _693_i7.plus(_dafny.ONE)) {
        let _694_v63;
        _694_v63 = _dafny.Seq.UnicodeFromString("xgk");
        let _695_v64;
        _695_v64 = _dafny.Map.Empty.slice().updateUnsafe(p2,_694_v63);
        let _696_v65;
        _696_v65 = _dafny.Seq.of(_this.f7);
        let _697_v66;
        let _nw115 = Array((new BigNumber(19)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = (_this).f6;
        _nw115[(_dafny.ONE).toNumber()] = (((_608_v0).dtor_cf5) ? ((_608_v0).dtor_cf6) : ((_this).f6));
        _nw115[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw115[(new BigNumber(3)).toNumber()] = _this.f7;
        _nw115[(new BigNumber(4)).toNumber()] = p1;
        _nw115[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_694_v63, (((_695_v64).contains(_693_i7)) ? ((_695_v64).get(_693_i7)) : (_694_v63)))).length);
        _nw115[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus((_this).f6)).plus(p2);
        _nw115[(new BigNumber(7)).toNumber()] = p2;
        _nw115[(new BigNumber(8)).toNumber()] = p2;
        _nw115[(new BigNumber(9)).toNumber()] = p1;
        _nw115[(new BigNumber(10)).toNumber()] = _693_i7;
        _nw115[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm1(p0, globalState));
        _nw115[(new BigNumber(12)).toNumber()] = p2;
        _nw115[(new BigNumber(13)).toNumber()] = _693_i7;
        _nw115[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_693_i7);
        _nw115[(new BigNumber(15)).toNumber()] = (_696_v65)[_module.__default.safeIndex(p2, new BigNumber((_696_v65).length))];
        _nw115[(new BigNumber(16)).toNumber()] = (_608_v0).dtor_cf6;
        _nw115[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw115[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(_module.__default.fm1(p0, globalState), _this.f7);
        _697_v66 = _nw115;
        let _index113 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_697_v66).length));
        (_697_v66)[_index113] = (_this).f6;
        let _index114 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_697_v66).length));
        (_697_v66)[_index114] = new BigNumber(242);
        let _698_v67;
        _698_v67 = _dafny.Map.Empty.slice().updateUnsafe(p0,(new BigNumber((_694_v63).length)).multipliedBy(p1));
        _698_v67 = (_698_v67).update(p0, p2);
        let _699_v68;
        _699_v68 = _dafny.Map.Empty.slice().updateUnsafe((_696_v65)[_module.__default.safeIndex(new BigNumber((_696_v65).length), new BigNumber((_696_v65).length))],p1);
        _699_v68 = (_699_v68).update((_this).f6, (((_698_v67).contains(p0)) ? ((_698_v67).get(p0)) : ((_this).f6)));
        if (p0) {
          let _700_v69;
          let _nw116 = Array((new BigNumber(27)).toNumber()).fill(false);
          _700_v69 = _nw116;
          let _index115 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_700_v69).length));
          (_700_v69)[_index115] = ((_dafny.ZERO).minus(_this.f7)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(false, _699_v68, (_this).f6, globalState),(_697_v66)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_697_v66).length))])).length));
          let _index116 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_700_v69).length));
          (_700_v69)[_index116] = !(p0);
          let _index117 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_697_v66).length));
          (_697_v66)[_index117] = _module.__default.fm1(p0, globalState);
          let _701_v70;
          _701_v70 = _dafny.MultiSet.fromElements(new BigNumber((_694_v63).length));
          let _index118 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_700_v69).length));
          let _rhs136 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((new BigNumber(145)).multipliedBy(p2)), p2);
          let _rhs137 = ((new BigNumber(604)).isLessThanOrEqualTo((_this).f6)) === ((_701_v70).IsDisjointFrom(_701_v70));
          let _lhs89 = globalState;
          let _lhs90 = _700_v69;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_700_v69).length));
          _lhs89.f1 = _rhs136;
          _lhs90[_lhs91] = _rhs137;
          let _702_v71;
          _702_v71 = new _dafny.CodePoint('i'.codePointAt(0));
          let _703_v72;
          _703_v72 = _dafny.Seq.of(_694_v63);
          let _704_v73;
          _704_v73 = _dafny.Seq.of((_700_v69)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_700_v69).length))]);
          let _705_v74;
          _705_v74 = _dafny.Set.fromElements(_694_v63, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), function (_706_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), function (_707_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length)), _702_v71), (_703_v72)[_module.__default.safeIndex(new BigNumber((_704_v73).length), new BigNumber((_703_v72).length))]);
          _705_v74 = _705_v74;
          let _708_v75;
          _708_v75 = _dafny.Set.fromElements(_696_v65);
          let _rhs138 = (_708_v75).IsProperSubsetOf(_708_v75);
          let _rhs139 = (_this).f6;
          let _lhs92 = globalState;
          r0 = _rhs138;
          _lhs92.f1 = _rhs139;
        } else {
          let _709_v76;
          let _nw117 = new _module.C1();
          _nw117.__ctor(new _dafny.CodePoint('x'.codePointAt(0)), _693_i7, (new BigNumber((_690_v60).length)).minus(p1));
          _709_v76 = _nw117;
          let _710_v77;
          _710_v77 = _module.D8.create_DC18(_709_v76);
          let _711_v78;
          _711_v78 = _module.D8.create_DC18((_710_v77).dtor_cf33);
          _709_v76 = (_710_v77).dtor_cf33;
          let _712_v79;
          let _nw118 = new _module.C0();
          _nw118.__ctor((p0) || (p0));
          _712_v79 = _nw118;
          _697_v66 = _697_v66;
          let _713_v80;
          _713_v80 = _dafny.Set.fromElements(p0);
          let _714_v81;
          _714_v81 = _module.D5.create_DC12(new BigNumber((_713_v80).length));
          let _715_v82;
          _715_v82 = _dafny.Seq.of(_712_v79.f10);
          _714_v81 = _module.__default.fm28(_715_v82, (_dafny.ZERO).minus(_693_i7), globalState);
          let _716_v83;
          let _nw119 = Array((new BigNumber(12)).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = p0;
          _nw119[(_dafny.ONE).toNumber()] = p0;
          _nw119[(new BigNumber(2)).toNumber()] = _712_v79.f10;
          _nw119[(new BigNumber(3)).toNumber()] = _712_v79.f10;
          _nw119[(new BigNumber(4)).toNumber()] = p0;
          _nw119[(new BigNumber(5)).toNumber()] = p0;
          _nw119[(new BigNumber(6)).toNumber()] = false;
          _nw119[(new BigNumber(7)).toNumber()] = p0;
          _nw119[(new BigNumber(8)).toNumber()] = p0;
          _nw119[(new BigNumber(9)).toNumber()] = p0;
          _nw119[(new BigNumber(10)).toNumber()] = p0;
          _nw119[(new BigNumber(11)).toNumber()] = _712_v79.f10;
          _716_v83 = _nw119;
          let _717_v84;
          _717_v84 = _dafny.Seq.of(_716_v83);
          let _718_v85;
          let _nw120 = Array((new BigNumber(11)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _716_v83;
          _nw120[(_dafny.ONE).toNumber()] = _716_v83;
          _nw120[(new BigNumber(2)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(3)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(4)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(5)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(6)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(7)).toNumber()] = (_717_v84)[_module.__default.safeIndex(p2, new BigNumber((_717_v84).length))];
          _nw120[(new BigNumber(8)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(9)).toNumber()] = _716_v83;
          _nw120[(new BigNumber(10)).toNumber()] = _716_v83;
          _718_v85 = _nw120;
          let _719_v86;
          let _nw121 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _719_v86 = _nw121;
          let _720_v87;
          _720_v87 = new _dafny.CodePoint('o'.codePointAt(0));
          let _index119 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_719_v86).length));
          (_719_v86)[_index119] = _720_v87;
          let _index120 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_719_v86).length));
          let _rhs140 = _718_v85;
          let _rhs141 = (((_690_v60).contains(p0)) ? ((_690_v60).get(p0)) : (p0));
          let _rhs142 = _720_v87;
          let _rhs143 = _709_v76;
          let _rhs144 = (_709_v76).f6;
          let _lhs93 = _712_v79;
          let _lhs94 = _719_v86;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_719_v86).length));
          let _lhs96 = globalState;
          _718_v85 = _rhs140;
          _lhs93.f10 = _rhs141;
          _lhs94[_lhs95] = _rhs142;
          _709_v76 = _rhs143;
          _lhs96.f1 = _rhs144;
        }
      }
      let _hi8 = p2;
      for (let _721_i9 = (_this).f6; _721_i9.isLessThan(_hi8); _721_i9 = _721_i9.plus(_dafny.ONE)) {
        r0 = !(p0);
        if (!(p0)) {
          let _722_v88;
          _722_v88 = _dafny.Seq.of(_721_i9);
          let _723_v89;
          let _init22 = function (_724_i10) {
            return (_724_i10).multipliedBy(new BigNumber(397));
          };
          let _nw122 = Array((new BigNumber(12)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw122.length); _i0_22++) {
            _nw122[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _723_v89 = _nw122;
          let _725_v90;
          _725_v90 = _dafny.Set.fromElements(_722_v88);
          let _726_v91;
          _726_v91 = new _dafny.CodePoint('i'.codePointAt(0));
          let _727_v92;
          _727_v92 = _dafny.MultiSet.fromElements(_726_v91, _726_v91);
          let _728_v93;
          _728_v93 = _dafny.Seq.of(p0);
          let _729_v94;
          let _nw123 = Array((new BigNumber(23)).toNumber());
          _nw123[(_dafny.ZERO).toNumber()] = p2;
          _nw123[(_dafny.ONE).toNumber()] = p1;
          _nw123[(new BigNumber(2)).toNumber()] = p2;
          _nw123[(new BigNumber(3)).toNumber()] = new BigNumber(687);
          _nw123[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_721_i9);
          _nw123[(new BigNumber(5)).toNumber()] = p1;
          _nw123[(new BigNumber(6)).toNumber()] = _module.__default.fm1(!(p0), globalState);
          _nw123[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_722_v88).length)));
          _nw123[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.of(_723_v89, _723_v89)).length);
          _nw123[(new BigNumber(9)).toNumber()] = _721_i9;
          _nw123[(new BigNumber(10)).toNumber()] = new BigNumber((_725_v90).length);
          _nw123[(new BigNumber(11)).toNumber()] = p2;
          _nw123[(new BigNumber(12)).toNumber()] = new BigNumber((_727_v92).cardinality());
          _nw123[(new BigNumber(13)).toNumber()] = new BigNumber((_728_v93).length);
          _nw123[(new BigNumber(14)).toNumber()] = (_this).f6;
          _nw123[(new BigNumber(15)).toNumber()] = new BigNumber((_722_v88).length);
          _nw123[(new BigNumber(16)).toNumber()] = new BigNumber(503);
          _nw123[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f6)).length);
          _nw123[(new BigNumber(18)).toNumber()] = p2;
          _nw123[(new BigNumber(19)).toNumber()] = new BigNumber(835);
          _nw123[(new BigNumber(20)).toNumber()] = _this.f7;
          _nw123[(new BigNumber(21)).toNumber()] = (_this).f6;
          _nw123[(new BigNumber(22)).toNumber()] = new BigNumber(610);
          _729_v94 = _nw123;
          let _730_v95;
          _730_v95 = _dafny.Map.Empty.slice().updateUnsafe(_723_v89,p0);
          let _731_v96;
          _731_v96 = _module.D2.create_DC5(_723_v89);
          r0 = (((_730_v95).contains((_731_v96).dtor_cf9)) ? ((_730_v95).get((_731_v96).dtor_cf9)) : (p0));
          let _732_v97;
          let _nw124 = new _module.C0();
          _nw124.__ctor(p0);
          _732_v97 = _nw124;
          let _733_v98;
          _733_v98 = _dafny.MultiSet.fromElements((p1).multipliedBy(_721_i9));
          let _734_v99;
          _734_v99 = _module.D6.create_DC14(p1, _dafny.MultiSet.fromElements(p0, (((_690_v60).contains(_732_v97.f10)) ? ((_690_v60).get(_732_v97.f10)) : (_732_v97.f10))), new BigNumber((_690_v60).length));
          let _735_v100;
          _735_v100 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_734_v99).dtor_cf23, p1),p0);
          let _rhs145 = ((p0) ? (false) : (p0));
          let _rhs146 = (((_733_v98).contains(new BigNumber((_dafny.Seq.Concat(_722_v88, _722_v88)).length))) ? ((_733_v98).get(new BigNumber((_dafny.Seq.Concat(_722_v88, _722_v88)).length))) : (new BigNumber((_735_v100).length)));
          let _rhs147 = _732_v97;
          r0 = _rhs145;
          r2 = _rhs146;
          _732_v97 = _rhs147;
          let _736_v101;
          _736_v101 = _dafny.Map.Empty.slice().updateUnsafe(_732_v97.f10,_dafny.Seq.UnicodeFromString("ksw"));
          let _737_v102;
          _737_v102 = _dafny.Seq.UnicodeFromString("mqxjodfld");
          _736_v101 = (_736_v101).update(_732_v97.f10, _737_v102);
          let _738_v103;
          let _nw125 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _738_v103 = _nw125;
          let _index121 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_738_v103).length));
          (_738_v103)[_index121] = ((p0) ? (_726_v91) : (new _dafny.CodePoint('j'.codePointAt(0))));
          let _index122 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_738_v103).length));
          (_738_v103)[_index122] = _726_v91;
          let _739_v104;
          let _init23 = ((_740_v97) => function (_741_i11) {
            return _740_v97.f10;
          })(_732_v97);
          let _nw126 = Array((new BigNumber(3)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw126.length); _i0_23++) {
            _nw126[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _739_v104 = _nw126;
          let _742_v105;
          _742_v105 = _dafny.Map.Empty.slice().updateUnsafe(_739_v104,(_721_i9).isLessThan(_module.__default.fm1(_732_v97.f10, globalState)));
          _742_v105 = (_742_v105).update(_739_v104, (_728_v93)[_module.__default.safeIndex(p1, new BigNumber((_728_v93).length))]);
        } else {
          r1 = p0;
          let _743_v106;
          _743_v106 = _dafny.Set.fromElements(p0, true);
          let _744_v107;
          _744_v107 = _module.D0.create_DC0(_743_v106);
          let _745_v108;
          _745_v108 = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_744_v107,p0));
          let _746_v109;
          _746_v109 = _dafny.Seq.UnicodeFromString("fbkeac");
          let _747_v110;
          _747_v110 = _module.D2.create_DC6(_745_v108, _746_v109, p1);
          (_this).m8((_690_v60).Merge(_690_v60), _747_v110, _dafny.Seq.Concat(_746_v109, _dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), function (_748_i12) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })), globalState);
          let _749_v111;
          _749_v111 = _dafny.MultiSet.fromElements(p2, _721_i9);
          let _750_v112;
          _750_v112 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_749_v111).cardinality()),_this.f7);
          r1 = _module.__default.fm0(p0, _750_v112, _721_i9, globalState);
          let _751_v113;
          _751_v113 = _module.D9.create_DC22(_dafny.Seq.of(_this.f7));
          let _752_v114;
          _752_v114 = _dafny.Map.Empty.slice().updateUnsafe(_749_v111,!(p1).isEqualTo(new BigNumber(((_751_v113).dtor_cf45).length)));
          _752_v114 = (_752_v114).update(_749_v111, p0);
          let _753_v115;
          let _init24 = function (_754_i13) {
            return _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0))));
          };
          let _nw127 = Array((new BigNumber(15)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw127.length); _i0_24++) {
            _nw127[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _753_v115 = _nw127;
          let _755_v116;
          _755_v116 = new _dafny.CodePoint('b'.codePointAt(0));
          let _756_v117;
          _756_v117 = _dafny.MultiSet.fromElements(_755_v116, new _dafny.CodePoint('t'.codePointAt(0)), _755_v116, _755_v116, _755_v116);
          let _757_v118;
          _757_v118 = _dafny.MultiSet.fromElements(_756_v117, _dafny.MultiSet.fromElements(_755_v116), _756_v117);
          let _index123 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_753_v115).length));
          (_753_v115)[_index123] = (_757_v118).Union(_757_v118);
          let _index124 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_753_v115).length));
          (_753_v115)[_index124] = _757_v118;
        }
        (_this).f7 = p1;
        let _758_v119;
        _758_v119 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        let _759_v120;
        _759_v120 = _dafny.Seq.of(p0);
        let _760_v121;
        _760_v121 = _dafny.Seq.UnicodeFromString("kifqmnwow");
        let _761_v122;
        let _nw128 = Array((new BigNumber(17)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = p1;
        _nw128[(_dafny.ONE).toNumber()] = p2;
        _nw128[(new BigNumber(2)).toNumber()] = _module.__default.fm1(p0, globalState);
        _nw128[(new BigNumber(3)).toNumber()] = _this.f7;
        _nw128[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw128[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
        _nw128[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(p1, p2);
        _nw128[(new BigNumber(7)).toNumber()] = p2;
        _nw128[(new BigNumber(8)).toNumber()] = ((p0) ? ((_dafny.ZERO).minus(p2)) : (_721_i9));
        _nw128[(new BigNumber(9)).toNumber()] = (p1).plus((((_758_v119).contains(p0)) ? ((_758_v119).get(p0)) : (new BigNumber((_module.__default.fm30((_759_v120)[_module.__default.safeIndex(p1, new BigNumber((_759_v120).length))], p0, p1, p1, globalState)).length))));
        _nw128[(new BigNumber(10)).toNumber()] = (_721_i9).multipliedBy(p1);
        _nw128[(new BigNumber(11)).toNumber()] = _721_i9;
        _nw128[(new BigNumber(12)).toNumber()] = (_721_i9).plus(p2);
        _nw128[(new BigNumber(13)).toNumber()] = new BigNumber(-181);
        _nw128[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_762_i14) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _760_v121)).length);
        _nw128[(new BigNumber(15)).toNumber()] = new BigNumber(767);
        _nw128[(new BigNumber(16)).toNumber()] = new BigNumber(277);
        _761_v122 = _nw128;
        let _index125 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_761_v122).length));
        (_761_v122)[_index125] = _this.f7;
        let _763_v123;
        _763_v123 = _dafny.Seq.of(_721_i9);
        let _764_v124;
        _764_v124 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_721_i9),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_763_v123)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,p0)).length), new BigNumber((_763_v123).length))],true)).length));
        let _index126 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_761_v122).length));
        (_761_v122)[_index126] = _module.__default.safeDivisionInt((_module.__default.fm1(p0, globalState)).multipliedBy(new BigNumber((_690_v60).length)), ((_dafny.ZERO).minus((((_758_v119).contains(p0)) ? ((_758_v119).get(p0)) : ((((_764_v124).contains(p2)) ? ((_764_v124).get(p2)) : (_module.__default.fm1(p0, globalState))))))).multipliedBy((_this).f6));
      }
      let _765_v125;
      _765_v125 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f6);
      let _766_v126;
      _766_v126 = _dafny.Seq.of(new BigNumber((_765_v125).length));
      let _hi9 = p1;
      for (let _767_i15 = ((_766_v126)[_module.__default.safeIndex(_this.f7, new BigNumber((_766_v126).length))]).multipliedBy(p1); _767_i15.isLessThan(_hi9); _767_i15 = _767_i15.plus(_dafny.ONE)) {
        if ((p0) || ((((_608_v0).dtor_cf5) ? (p0) : (false)))) {
          let _768_v127;
          _768_v127 = _dafny.Seq.UnicodeFromString("vg");
          (_this).m9(p0, ((!(p0)) ? (new BigNumber((_768_v127).length)) : (_767_i15)), p1, new BigNumber(-200), globalState);
          let _769_v128;
          let _nw129 = Array((new BigNumber(6)).toNumber());
          _nw129[(_dafny.ZERO).toNumber()] = _this.f7;
          _nw129[(_dafny.ONE).toNumber()] = _767_i15;
          _nw129[(new BigNumber(2)).toNumber()] = _767_i15;
          _nw129[(new BigNumber(3)).toNumber()] = ((p0) ? ((_dafny.ZERO).minus(new BigNumber(-823))) : (p1));
          _nw129[(new BigNumber(4)).toNumber()] = p2;
          _nw129[(new BigNumber(5)).toNumber()] = p1;
          _769_v128 = _nw129;
          _769_v128 = _769_v128;
          let _770_v129;
          let _nw130 = new _module.C0();
          _nw130.__ctor((p0) === (p0));
          _770_v129 = _nw130;
          (_770_v129).f10 = p0;
          (_this).f7 = (_dafny.ZERO).minus((((((_692_v62).contains(_767_i15)) ? ((_692_v62).get(_767_i15)) : (p0))) ? (_767_i15) : ((_this.f7).minus(_767_i15))));
        } else {
          let _771_v130;
          let _nw131 = Array((new BigNumber(23)).toNumber()).fill([]);
          _771_v130 = _nw131;
          let _772_v131;
          _772_v131 = _dafny.Seq.UnicodeFromString("oiiryjckw");
          let _rhs148 = _dafny.Seq.IsPrefixOf(_772_v131, _772_v131);
          let _rhs149 = _771_v130;
          r1 = _rhs148;
          _771_v130 = _rhs149;
          let _773_v132;
          _773_v132 = new _dafny.CodePoint('t'.codePointAt(0));
          let _774_v133;
          _774_v133 = _module.D7.create_DC17(_772_v131, _module.__default.fm31((_this).f6, p1, _this.f7, _773_v132, globalState));
          let _775_v134;
          _775_v134 = _dafny.Seq.of(p0);
          let _776_v135;
          _776_v135 = _dafny.MultiSet.fromElements(new BigNumber((_775_v134).length), new BigNumber(885), new BigNumber((_775_v134).length));
          let _777_v136;
          _777_v136 = _dafny.MultiSet.fromElements(new BigNumber((_776_v135).cardinality()));
          let _778_v137;
          _778_v137 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_766_v126).length));
          let _779_v138;
          let _nw132 = new _module.C1();
          _nw132.__ctor(_773_v132, _this.f7, (((_777_v136).contains(_767_i15)) ? ((_777_v136).get(_767_i15)) : (new BigNumber((_778_v137).length))));
          _779_v138 = _nw132;
          let _780_v139;
          _780_v139 = _dafny.Map.Empty.slice().updateUnsafe(_774_v133,_779_v138);
          let _781_v140;
          _781_v140 = _dafny.Set.fromElements(p0, p0);
          let _782_v141;
          _782_v141 = _module.D0.create_DC0(_781_v140);
          let _783_v142;
          _783_v142 = _dafny.Map.Empty.slice().updateUnsafe(_782_v141,p0);
          let _784_v143;
          _784_v143 = _module.D1.create_DC2(_783_v142);
          let _785_v144;
          _785_v144 = _module.D2.create_DC6(_784_v143, _772_v131, (_779_v138).f6);
          let _786_v145;
          _786_v145 = _module.D2.create_DC7(!(p0), _this.f7, _692_v62, (_this).fm23(new BigNumber(260), p1, _785_v144, globalState));
          let _787_v146;
          _787_v146 = _dafny.Set.fromElements(_this.f7);
          let _rhs150 = _780_v139;
          let _rhs151 = (new BigNumber((_775_v134).length)).plus(((_this).f6).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,true)).length)));
          let _rhs152 = ((_786_v145).dtor_cf14).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_787_v146).length), _779_v138.f7));
          let _rhs153 = p0;
          let _lhs97 = _this;
          _780_v139 = _rhs150;
          _lhs97.f7 = _rhs151;
          r0 = _rhs152;
          r0 = _rhs153;
          r0 = ((p0) ? (p0) : ((new BigNumber((_772_v131).length)).isLessThan((_779_v138).f6)));
          r1 = p0;
          r0 = p0;
        }
        let _788_v147;
        let _init25 = ((_789_p2) => function (_790_i16) {
          return (_790_i16).plus(_789_p2);
        })(p2);
        let _nw133 = Array((new BigNumber(7)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw133.length); _i0_25++) {
          _nw133[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _788_v147 = _nw133;
        let _791_v148;
        _791_v148 = _dafny.Seq.UnicodeFromString("jaegurgi");
        let _792_v149;
        _792_v149 = new _dafny.CodePoint('s'.codePointAt(0));
        let _793_v150;
        _793_v150 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f7);
        let _rhs154 = _788_v147;
        let _rhs155 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("amfxpo"), _module.__default.safeIndex(new BigNumber(295), new BigNumber((_dafny.Seq.UnicodeFromString("amfxpo")).length)), _792_v149);
        let _rhs156 = _788_v147;
        let _rhs157 = p0;
        let _rhs158 = _dafny.Seq.update(_dafny.Seq.Concat(_766_v126, _766_v126), _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.Concat(_766_v126, _766_v126)).length)), _module.__default.safeDivisionInt((_this).f6, new BigNumber((_793_v150).length)));
        _788_v147 = _rhs154;
        _791_v148 = _rhs155;
        _788_v147 = _rhs156;
        r1 = _rhs157;
        _766_v126 = _rhs158;
        let _index127 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_788_v147).length));
        (_788_v147)[_index127] = ((_this).f6).multipliedBy(_767_i15);
        let _index128 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_788_v147).length));
        (_788_v147)[_index128] = _module.__default.safeModuloInt(p1, (_this).f6);
        let _794_v151;
        let _nw134 = new _module.C0();
        _nw134.__ctor(p0);
        _794_v151 = _nw134;
      }
      let _795_v152;
      _795_v152 = _dafny.Seq.UnicodeFromString("akifd");
      let _796_v153;
      _796_v153 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_795_v152).length)));
      if (_module.__default.fm0(!((_796_v153)[_module.__default.safeIndex((_this).f6, new BigNumber((_796_v153).length))]).contains((_this).f6), (_dafny.Map.Empty.slice().updateUnsafe(p2,_this.f7)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_765_v125).length),p2)), p2, globalState)) {
        let _797_v154;
        _797_v154 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(854)), ((_798_p1) => function (_799_i18) {
          return _798_p1;
        })(p1)), _766_v126, _766_v126, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-931)), function (_800_i19) {
          return _this.f7;
        }), _766_v126);
        let _801_v155;
        let _nw135 = Array((new BigNumber(3)).toNumber());
        _nw135[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), function (_802_i17) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }), _795_v152);
        _nw135[(_dafny.ONE).toNumber()] = _module.__default.fm2((_797_v154)[_module.__default.safeIndex(p1, new BigNumber((_797_v154).length))], _this.f7, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_803_i20) {
          return new BigNumber(184);
        })).length), p0, globalState);
        _nw135[(new BigNumber(2)).toNumber()] = _795_v152;
        _801_v155 = _nw135;
        let _index129 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_801_v155).length));
        (_801_v155)[_index129] = _dafny.Seq.Concat(_795_v152, _795_v152);
        let _804_v156;
        let _nw136 = Array((new BigNumber(18)).toNumber()).fill([]);
        _804_v156 = _nw136;
        let _805_v157;
        let _nw137 = Array((new BigNumber(14)).toNumber()).fill(false);
        _805_v157 = _nw137;
        let _index130 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_804_v156).length));
        (_804_v156)[_index130] = _805_v157;
        let _index131 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_801_v155).length));
        let _index132 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_804_v156).length));
        let _rhs159 = _dafny.Seq.UnicodeFromString("hv");
        let _rhs160 = _805_v157;
        let _lhs98 = _801_v155;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_801_v155).length));
        let _lhs100 = _804_v156;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_804_v156).length));
        _lhs98[_lhs99] = _rhs159;
        _lhs100[_lhs101] = _rhs160;
        (globalState).f1 = (_this).f6;
        let _806_v158;
        let _init26 = ((_807_v126) => function (_808_i21) {
          return _module.D9.create_DC22(_807_v126);
        })(_766_v126);
        let _nw138 = Array((new BigNumber(19)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw138.length); _i0_26++) {
          _nw138[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _806_v158 = _nw138;
        let _index133 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_806_v158).length));
        (_806_v158)[_index133] = _module.D9.create_DC22(_766_v126);
        let _809_v159;
        _809_v159 = new _dafny.CodePoint('u'.codePointAt(0));
        let _810_v160;
        _810_v160 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _811_v161;
        _811_v161 = _module.D5.create_DC12(new BigNumber((_dafny.Seq.of(_810_v160)).length));
        let _812_v162;
        _812_v162 = _dafny.Seq.of(_811_v161, _811_v161);
        let _813_v163;
        _813_v163 = _dafny.Seq.of(p0, p0, p0, true, p0);
        let _814_v164;
        _814_v164 = _module.D6.create_DC15(p0, p2, new BigNumber(79), p2);
        let _index134 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_806_v158).length));
        let _rhs161 = _module.__default.safeModuloInt(new BigNumber((_812_v162).length), p1);
        let _rhs162 = _module.__default.fm32(p0, !(!(_dafny.Seq.IsPrefixOf(_813_v163, _813_v163))), p0, p1, globalState);
        let _rhs163 = _809_v159;
        let _rhs164 = (_814_v164).dtor_cf26;
        let _lhs102 = globalState;
        let _lhs103 = _806_v158;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_806_v158).length));
        _lhs102.f1 = _rhs161;
        _lhs103[_lhs104] = _rhs162;
        _809_v159 = _rhs163;
        r1 = _rhs164;
        let _815_v165;
        let _nw139 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _815_v165 = _nw139;
        _815_v165 = _815_v165;
        let _index135 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_815_v165).length));
        (_815_v165)[_index135] = ((_dafny.ZERO).minus((_766_v126)[_module.__default.safeIndex(p2, new BigNumber((_766_v126).length))])).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(316)), function (_816_i22) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("p")).length);
        })).length));
        let _index136 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_815_v165).length));
        (_815_v165)[_index136] = p1;
      } else {
        _766_v126 = _766_v126;
        let _817_v166;
        _817_v166 = _dafny.MultiSet.fromElements(p0, p0, false, p0, (p0) || (p0));
        _817_v166 = _817_v166;
        let _818_v167;
        _818_v167 = _module.D10.create_DC24(new _dafny.CodePoint('y'.codePointAt(0)));
        let _819_v168;
        let _nw140 = new _module.C1();
        _nw140.__ctor((_818_v167).dtor_cf51, p2, (_this).f6);
        _819_v168 = _nw140;
        r1 = !(_dafny.areEqual(_795_v152, _795_v152)) || (p0);
        let _820_v169;
        let _nw141 = new _module.C1();
        _nw141.__ctor((_819_v168).f13, _module.__default.fm1(p0, globalState), (new BigNumber(-480)).plus(p1));
        _820_v169 = _nw141;
      }
      r0 = p0;
      let _821_v170;
      _821_v170 = _dafny.MultiSet.fromElements(p0);
      r1 = (((p0) ? (_module.__default.fm33(globalState)) : (_821_v170))).contains(((!(p0)) ? (p0) : (p0)));
      let _822_v171;
      _822_v171 = _dafny.MultiSet.fromElements((_this).f6);
      r2 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,p0)).length), _this.f7), (_dafny.ZERO).minus((((_822_v171).contains(p2)) ? ((_822_v171).get(p2)) : (new BigNumber((_822_v171).cardinality())))));
      return [r0, r1, r2];
    }
    m8(p0, p1, p2, globalState) {
      let _this = this;
      let _823_v0;
      _823_v0 = true;
      let _rhs165 = _823_v0;
      let _rhs166 = _this.f7;
      let _lhs105 = globalState;
      _823_v0 = _rhs165;
      _lhs105.f1 = _rhs166;
      let _824_v1;
      _824_v1 = _dafny.Seq.of(p2);
      let _825_v2;
      _825_v2 = _dafny.MultiSet.fromElements(new BigNumber((_824_v1).length));
      let _826_v3;
      _826_v3 = _dafny.MultiSet.fromElements(new BigNumber((_825_v2).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_823_v0,_dafny.Seq.of((_this).f6, _this.f7))).length));
      let _827_v4;
      let _nw142 = Array((new BigNumber(6)).toNumber());
      _nw142[(_dafny.ZERO).toNumber()] = new BigNumber((p2).length);
      _nw142[(_dafny.ONE).toNumber()] = (_this).f6;
      _nw142[(new BigNumber(2)).toNumber()] = (((_826_v3).contains(new BigNumber(683))) ? ((_826_v3).get(new BigNumber(683))) : ((_this).f6));
      _nw142[(new BigNumber(3)).toNumber()] = _this.f7;
      _nw142[(new BigNumber(4)).toNumber()] = _this.f7;
      _nw142[(new BigNumber(5)).toNumber()] = (_this).f6;
      _827_v4 = _nw142;
      let _828_v5;
      _828_v5 = _dafny.Set.fromElements(_827_v4, _827_v4, _827_v4, _827_v4, _827_v4);
      let _829_v6;
      _829_v6 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_828_v5).length)), (_this).f6, _this.f7);
      _829_v6 = function () {
        let _coll24 = new _dafny.Set();
        for (const _compr_24 of ((_829_v6).Difference(_829_v6)).Elements) {
          let _830_v7 = _compr_24;
          if (((_829_v6).Difference(_829_v6)).contains(_830_v7)) {
            _coll24.add(_module.__default.safeModuloInt(_830_v7, new BigNumber(264)));
          }
        }
        return _coll24;
      }();
      if (!(_823_v0)) {
        _823_v0 = _823_v0;
        let _831_v8;
        _831_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,!(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(556),(_this).f6)).length)).isEqualTo((_this).f6));
        let _832_v9;
        _832_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this).f6);
        _823_v0 = (((_831_v8).contains((_this).f6)) ? ((_831_v8).get((_this).f6)) : (_module.__default.fm0(_823_v0, _832_v9, new BigNumber(472), globalState)));
        let _833_v10;
        _833_v10 = new _dafny.CodePoint('g'.codePointAt(0));
        let _834_v11;
        let _nw143 = new _module.C1();
        _nw143.__ctor(_833_v10, _module.__default.safeModuloInt(new BigNumber(-990), new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_this.f7), _module.__default.safeIndex(new BigNumber((p2).length), new BigNumber((_dafny.Seq.of(_this.f7)).length)), new BigNumber((_dafny.Set.fromElements(new BigNumber(151), (_this).f6)).length))).length)), ((_823_v0) ? (new BigNumber(-72)) : (new BigNumber((p2).length))));
        _834_v11 = _nw143;
        (_this).f7 = _this.f7;
        let _835_v12;
        _835_v12 = _dafny.MultiSet.fromElements(_823_v0);
        let _836_v13;
        let _nw144 = new _module.C1();
        _nw144.__ctor((_834_v11).f13, _this.f7, _module.__default.safeDivisionInt(new BigNumber((_835_v12).cardinality()), _this.f7));
        _836_v13 = _nw144;
      } else {
        let _837_v14;
        _837_v14 = _dafny.Seq.UnicodeFromString("jx");
        let _838_v15;
        let _nw145 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
        _838_v15 = _nw145;
        let _index137 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_838_v15).length));
        (_838_v15)[_index137] = (_825_v2).Difference(_825_v2);
        let _839_v16;
        _839_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_823_v0);
        let _840_v19;
        _840_v19 = _dafny.Seq.of(_this.f7);
        let _841_v20;
        let _nw146 = Array((new BigNumber(5)).toNumber());
        _nw146[(_dafny.ZERO).toNumber()] = _839_v16;
        _nw146[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_823_v0);
        _nw146[(new BigNumber(2)).toNumber()] = function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(856), new BigNumber(645))) {
            let _842_v17 = _compr_25;
            if (((new BigNumber(856)).isLessThanOrEqualTo(_842_v17)) && ((_842_v17).isLessThan(new BigNumber(645)))) {
              _coll25.push([(_842_v17).plus(new BigNumber((_dafny.Seq.UnicodeFromString("xrpgkhh")).length)),_823_v0]);
            }
          }
          return _coll25;
        }();
        _nw146[(new BigNumber(3)).toNumber()] = (_839_v16).Merge(function () {
          let _coll26 = new _dafny.Map();
          for (const _compr_26 of (_840_v19).Elements) {
            let _843_v18 = _compr_26;
            if (_dafny.Seq.contains(_840_v19, _843_v18)) {
              _coll26.push([_module.__default.safeModuloInt(_843_v18, (_this).f6),_823_v0]);
            }
          }
          return _coll26;
        }());
        _nw146[(new BigNumber(4)).toNumber()] = _839_v16;
        _841_v20 = _nw146;
        let _index138 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_838_v15).length));
        let _rhs167 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(223)), function (_844_i0) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }), p2), _dafny.Seq.Create(_module.__default.abs(new BigNumber(983)), function (_845_i1) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }));
        let _rhs168 = _dafny.MultiSet.FromArray(_840_v19);
        let _rhs169 = _823_v0;
        let _rhs170 = !(!((_this).f6).isEqualTo((_this.f7).multipliedBy((_this).f6)));
        let _rhs171 = _841_v20;
        let _lhs106 = _838_v15;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_838_v15).length));
        _837_v14 = _rhs167;
        _lhs106[_lhs107] = _rhs168;
        _823_v0 = _rhs169;
        _823_v0 = _rhs170;
        _841_v20 = _rhs171;
        let _846_v22;
        let _init27 = ((_847_v19) => function (_848_i2) {
          return (_dafny.Set.fromElements(_dafny.Seq.update(_847_v19, _module.__default.safeIndex((_this).f6, new BigNumber((_847_v19).length)), (_this).f6))).Union(function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of (_dafny.Seq.of(_847_v19, _847_v19)).Elements) {
              let _849_v21 = _compr_27;
              if (_dafny.Seq.contains(_dafny.Seq.of(_847_v19, _847_v19), _849_v21)) {
                _coll27.add(_849_v21);
              }
            }
            return _coll27;
          }());
        })(_840_v19);
        let _nw147 = Array((new BigNumber(15)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw147.length); _i0_27++) {
          _nw147[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _846_v22 = _nw147;
        let _850_v23;
        _850_v23 = _dafny.Set.fromElements(_840_v19, _840_v19);
        let _index139 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_846_v22).length));
        (_846_v22)[_index139] = _850_v23;
        let _851_v24;
        _851_v24 = _dafny.Map.Empty.slice().updateUnsafe(_840_v19,false);
        let _index140 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_846_v22).length));
        (_846_v22)[_index140] = ((_823_v0) ? (function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of ((_851_v24).update(_840_v19, !(false))).Keys.Elements) {
            let _852_v25 = _compr_28;
            if (((_851_v24).update(_840_v19, !(false))).contains(_852_v25)) {
              _coll28.add(_852_v25);
            }
          }
          return _coll28;
        }()) : (_850_v23));
        if (!(_823_v0)) {
          (globalState).f1 = _this.f7;
          let _853_v26;
          let _nw148 = new _module.C0();
          _nw148.__ctor(_823_v0);
          _853_v26 = _nw148;
          let _index141 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_827_v4).length));
          (_827_v4)[_index141] = (_dafny.ZERO).minus(_this.f7);
          let _854_v27;
          _854_v27 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_826_v3).cardinality())));
          let _855_v29;
          _855_v29 = new _dafny.CodePoint('f'.codePointAt(0));
          let _index142 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_827_v4).length));
          let _rhs172 = _this.f7;
          let _rhs173 = (_dafny.ZERO).minus(new BigNumber((((_854_v27)[_module.__default.safeIndex((_this).f6, new BigNumber((_854_v27).length))]).Merge(function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), function (_856_i3) {
              return (_this).f6;
            })).Elements) {
              let _857_v28 = _compr_29;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), function (_856_i3) {
                return (_this).f6;
              }), _857_v28)) {
                _coll29.push([(_857_v28).multipliedBy(_this.f7),_this.f7]);
              }
            }
            return _coll29;
          }())).length));
          let _rhs174 = new BigNumber((_dafny.Seq.Concat(_837_v14, ((_853_v26.f10) ? (_dafny.Seq.update(_837_v14, _module.__default.safeIndex((_this).f6, new BigNumber((_837_v14).length)), _855_v29)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(950)), ((_858_v29) => function (_859_i4) {
            return _858_v29;
          })(_855_v29)))))).length);
          let _rhs175 = p2;
          let _lhs108 = globalState;
          let _lhs109 = globalState;
          let _lhs110 = _827_v4;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_827_v4).length));
          _lhs108.f1 = _rhs172;
          _lhs109.f1 = _rhs173;
          _lhs110[_lhs111] = _rhs174;
          _837_v14 = _rhs175;
          let _860_v30;
          let _nw149 = new _module.C0();
          _nw149.__ctor(_823_v0);
          _860_v30 = _nw149;
          let _index143 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_827_v4).length));
          (_827_v4)[_index143] = _module.__default.safeModuloInt(new BigNumber(222), (_this.f7).minus((_this).f6));
        } else {
          let _861_v31;
          _861_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-201)), function (_862_i5) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })).length));
          let _rhs176 = true;
          let _rhs177 = _module.__default.fm0(_823_v0, _861_v31, _module.__default.safeModuloInt(new BigNumber(734), (_this).f6), globalState);
          _823_v0 = _rhs176;
          _823_v0 = _rhs177;
          (globalState).f1 = _module.__default.safeDivisionInt(_this.f7, _this.f7);
          (globalState).f1 = new BigNumber(767);
          (_this).m9(_823_v0, _this.f7, _this.f7, (_this).f6, globalState);
          (globalState).f1 = new BigNumber((p2).length);
        }
        let _863_v32;
        let _init28 = ((_864_v0) => function (_865_i6) {
          return _864_v0;
        })(_823_v0);
        let _nw150 = Array((new BigNumber(16)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw150.length); _i0_28++) {
          _nw150[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _863_v32 = _nw150;
        let _index144 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_863_v32).length));
        (_863_v32)[_index144] = !(_823_v0) || (_823_v0);
        let _866_v33;
        _866_v33 = _module.D10.create_DC26(_823_v0);
        let _pat_let_tv11 = _823_v0;
        let _867_v34;
        _867_v34 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let19_0) {
          return function (_868_dt__update__tmp_h0) {
            return function (_pat_let20_0) {
              return function (_869_dt__update_hcf52_h0) {
                return _module.D10.create_DC26(_869_dt__update_hcf52_h0);
              }(_pat_let20_0);
            }(_pat_let_tv11);
          }(_pat_let19_0);
        }(_866_v33),_823_v0);
        let _870_v35;
        _870_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_867_v34).length),(_this).f6);
        let _871_v36;
        _871_v36 = _dafny.Set.fromElements(_823_v0, _module.__default.fm0(_823_v0, _870_v35, _this.f7, globalState));
        let _index145 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_863_v32).length));
        (_863_v32)[_index145] = (_871_v36).IsSubsetOf(_871_v36);
        let _872_v37;
        _872_v37 = _dafny.Map.Empty.slice().updateUnsafe((_863_v32)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_863_v32).length))],new BigNumber((p2).length));
        (globalState).f1 = ((((_872_v37).contains(false)) ? ((_872_v37).get(false)) : (_this.f7))).minus(_this.f7);
      }
      let _873_v39;
      _873_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(-142));
      let _874_v40;
      _874_v40 = _dafny.Set.fromElements(_823_v0, _823_v0, _module.__default.fm0(_823_v0, _873_v39, new BigNumber(685), globalState));
      let _875_v41;
      _875_v41 = _dafny.Set.fromElements(_829_v6);
      let _876_v42;
      let _nw151 = Array((new BigNumber(24)).toNumber());
      _nw151[(_dafny.ZERO).toNumber()] = _823_v0;
      _nw151[(_dafny.ONE).toNumber()] = (_823_v0) === (_823_v0);
      _nw151[(new BigNumber(2)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(3)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(4)).toNumber()] = ((!(_823_v0)) ? (_823_v0) : (_823_v0));
      _nw151[(new BigNumber(5)).toNumber()] = !((_dafny.Set.fromElements(_this.f7, _this.f7, (_this).f6, _this.f7)).contains((_this).f6));
      _nw151[(new BigNumber(6)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(7)).toNumber()] = true;
      _nw151[(new BigNumber(8)).toNumber()] = !(_823_v0);
      _nw151[(new BigNumber(9)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(10)).toNumber()] = (_825_v2).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(85)));
      _nw151[(new BigNumber(11)).toNumber()] = (new BigNumber(779)).isEqualTo(new BigNumber((function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_829_v6).length),_825_v2)).Keys.Elements) {
          let _877_v38 = _compr_30;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_829_v6).length),_825_v2)).contains(_877_v38)) {
            _coll30.push([_module.__default.safeDivisionInt(_877_v38, new BigNumber((_829_v6).length)),_823_v0]);
          }
        }
        return _coll30;
      }()).length));
      _nw151[(new BigNumber(12)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(13)).toNumber()] = false;
      _nw151[(new BigNumber(14)).toNumber()] = !_dafny.areEqual(p2, p2);
      _nw151[(new BigNumber(15)).toNumber()] = !(_874_v40).contains(_823_v0);
      _nw151[(new BigNumber(16)).toNumber()] = _dafny.Seq.IsProperPrefixOf(p2, p2);
      _nw151[(new BigNumber(17)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(18)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(19)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(20)).toNumber()] = !(!(!((_875_v41).IsDisjointFrom(_module.__default.fm34(globalState)))));
      _nw151[(new BigNumber(21)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(22)).toNumber()] = _823_v0;
      _nw151[(new BigNumber(23)).toNumber()] = _823_v0;
      _876_v42 = _nw151;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_876_v42).length))) {
        let _878_i7 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_878_i7)) && ((_878_i7).isLessThan(new BigNumber((_876_v42).length))))) {
          (_876_v42)[(_878_i7)] = (_this.f7).isLessThan((_this).f6);
        }
      }
      _823_v0 = _823_v0;
      let _879_v43;
      _879_v43 = new _dafny.CodePoint('k'.codePointAt(0));
      let _880_v44;
      let _nw152 = new _module.C1();
      _nw152.__ctor(_879_v43, _this.f7, (_dafny.ZERO).minus((_this).f6));
      _880_v44 = _nw152;
      return;
    }
    m9(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _881_v0;
      _881_v0 = true;
      let _882_v1;
      _882_v1 = _dafny.Seq.UnicodeFromString("wdfh");
      _881_v0 = (_module.__default.fm0(p0, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_882_v1).length),p2), _this.f7, globalState)) || (p0);
      let _883_v2;
      _883_v2 = new _dafny.CodePoint('a'.codePointAt(0));
      let _884_v3;
      let _nw153 = new _module.C1();
      _nw153.__ctor(_883_v2, _this.f7, p2);
      _884_v3 = _nw153;
      let _rhs178 = (_884_v3).f6;
      let _rhs179 = _884_v3;
      let _lhs112 = _this;
      _lhs112.f7 = _rhs178;
      _884_v3 = _rhs179;
      (_884_v3).f7 = _884_v3.f7;
      if (!(!(p0))) {
        let _885_v4;
        _885_v4 = _dafny.Seq.of(_this.f7);
        (_884_v3).f7 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_885_v4).length), p1), _module.__default.fm1(_881_v0, globalState));
        (globalState).f1 = (_dafny.ZERO).minus(p2);
        let _886_v5;
        _886_v5 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("tjotuokyi"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-105)), ((_887_v2) => function (_888_i0) {
          return _887_v2;
        })(_883_v2)));
        (_884_v3).f7 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_886_v5, _886_v5), _dafny.Seq.Create(_module.__default.abs(new BigNumber(135)), ((_889_v1) => function (_890_i1) {
          return _889_v1;
        })(_882_v1)))).length);
        let _891_v6;
        _891_v6 = _module.D10.create_DC24(_883_v2);
        let _892_v7;
        let _nw154 = Array((new BigNumber(7)).toNumber());
        _nw154[(_dafny.ZERO).toNumber()] = _883_v2;
        _nw154[(_dafny.ONE).toNumber()] = _883_v2;
        _nw154[(new BigNumber(2)).toNumber()] = _883_v2;
        _nw154[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
        _nw154[(new BigNumber(4)).toNumber()] = _883_v2;
        _nw154[(new BigNumber(5)).toNumber()] = _883_v2;
        _nw154[(new BigNumber(6)).toNumber()] = (_891_v6).dtor_cf51;
        _892_v7 = _nw154;
        _892_v7 = _892_v7;
        let _893_v8;
        _893_v8 = _dafny.Seq.of(_881_v0);
        let _894_v9;
        _894_v9 = _module.D5.create_DC11(_893_v8);
        (_884_v3).f7 = new BigNumber((_dafny.Seq.update((_894_v9).dtor_cf20, _module.__default.safeIndex(new BigNumber(353), new BigNumber(((_894_v9).dtor_cf20).length)), p0)).length);
      } else {
        let _895_v10;
        _895_v10 = _dafny.Map.Empty.slice().updateUnsafe(_881_v0,p2);
        let _896_v11;
        let _nw155 = Array((new BigNumber(16)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = _895_v10;
        _nw155[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,(_884_v3).f6);
        _nw155[(new BigNumber(2)).toNumber()] = _895_v10;
        _nw155[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(false,p3)).Merge(_895_v10);
        _nw155[(new BigNumber(4)).toNumber()] = _895_v10;
        _nw155[(new BigNumber(5)).toNumber()] = (_895_v10).update(!(p0), p1);
        _nw155[(new BigNumber(6)).toNumber()] = (_895_v10).Merge(_module.__default.fm30(p0, false, _884_v3.f7, p2, globalState));
        _nw155[(new BigNumber(7)).toNumber()] = (_module.__default.fm30(p0, p0, new BigNumber(-807), p2, globalState)).Merge(_895_v10);
        _nw155[(new BigNumber(8)).toNumber()] = (_895_v10).Merge(_895_v10);
        _nw155[(new BigNumber(9)).toNumber()] = _895_v10;
        _nw155[(new BigNumber(10)).toNumber()] = (_895_v10).Merge(_895_v10);
        _nw155[(new BigNumber(11)).toNumber()] = (_895_v10).update(_881_v0, new BigNumber(834));
        _nw155[(new BigNumber(12)).toNumber()] = (_895_v10).Merge(_895_v10);
        _nw155[(new BigNumber(13)).toNumber()] = _895_v10;
        _nw155[(new BigNumber(14)).toNumber()] = _895_v10;
        _nw155[(new BigNumber(15)).toNumber()] = ((_895_v10).update(p0, p3)).update(p0, new BigNumber((_module.__default.fm35((_884_v3).f6, _884_v3.f7, globalState)).length));
        _896_v11 = _nw155;
        let _index146 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_896_v11).length));
        (_896_v11)[_index146] = (_895_v10).Merge(_dafny.Map.Empty.slice().updateUnsafe(_881_v0,_884_v3.f7));
        let _index147 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_896_v11).length));
        (_896_v11)[_index147] = _module.__default.fm30(p0, p0, (p1).minus(new BigNumber((_dafny.Seq.of(_884_v3.f7)).length)), (_dafny.ZERO).minus((new BigNumber(-943)).multipliedBy(_884_v3.f7)), globalState);
        let _897_v12;
        _897_v12 = _dafny.Seq.of(_this.f7);
        let _898_v13;
        let _nw156 = new _module.C0();
        _nw156.__ctor(p0);
        _898_v13 = _nw156;
        let _899_v14;
        _899_v14 = _dafny.MultiSet.fromElements(_898_v13);
        if (!(_884_v3.f7).isEqualTo(_module.__default.safeDivisionInt(new BigNumber((_module.__default.fm2(_dafny.Seq.update(_897_v12, _module.__default.safeIndex(new BigNumber((_899_v14).cardinality()), new BigNumber((_897_v12).length)), p3), new BigNumber(986), new BigNumber(475), p0, globalState)).length), p2))) {
          let _900_v15;
          _900_v15 = _dafny.Set.fromElements(p0);
          let _901_v16;
          _901_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_900_v15).IsSubsetOf(_dafny.Set.fromElements(_898_v13.f10)));
          _901_v16 = _901_v16;
          let _902_v17;
          _902_v17 = _module.D10.create_DC24(_883_v2);
          let _903_v18;
          let _nw157 = Array((new BigNumber(22)).toNumber());
          _nw157[(_dafny.ZERO).toNumber()] = _883_v2;
          _nw157[(_dafny.ONE).toNumber()] = _883_v2;
          _nw157[(new BigNumber(2)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(3)).toNumber()] = _module.__default.fm25(globalState);
          _nw157[(new BigNumber(4)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(5)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(6)).toNumber()] = (_902_v17).dtor_cf51;
          _nw157[(new BigNumber(7)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(8)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(9)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw157[(new BigNumber(11)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(12)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(13)).toNumber()] = _module.__default.fm25(globalState);
          _nw157[(new BigNumber(14)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(15)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(16)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(17)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(18)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(19)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(20)).toNumber()] = _883_v2;
          _nw157[(new BigNumber(21)).toNumber()] = _883_v2;
          _903_v18 = _nw157;
          let _904_v19;
          _904_v19 = _dafny.Seq.of(_898_v13, _898_v13, _898_v13);
          let _905_v20;
          _905_v20 = (_904_v19)[_module.__default.safeIndex(_884_v3.f7, new BigNumber((_904_v19).length))];
          let _906_v21;
          _906_v21 = _dafny.Map.Empty.slice().updateUnsafe(_903_v18,_905_v20);
          (_884_v3).f7 = new BigNumber((_906_v21).length);
          let _907_v22;
          _907_v22 = _dafny.Map.Empty.slice().updateUnsafe(_903_v18,p2);
          _907_v22 = (_907_v22).update(_903_v18, new BigNumber((_dafny.Seq.UnicodeFromString("hh")).length));
          _900_v15 = (_900_v15).Intersect(_900_v15);
          let _908_v23;
          _908_v23 = _dafny.Seq.of(_898_v13.f10);
          let _909_v24;
          _909_v24 = _module.D10.create_DC26(_898_v13.f10);
          let _910_v25;
          let _nw158 = Array((new BigNumber(24)).toNumber());
          _nw158[(_dafny.ZERO).toNumber()] = _908_v23;
          _nw158[(_dafny.ONE).toNumber()] = _908_v23;
          _nw158[(new BigNumber(2)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(3)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(4)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(5)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(6)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(7)).toNumber()] = _module.__default.fm27(p2, _882_v1, new _dafny.CodePoint('b'.codePointAt(0)), globalState);
          _nw158[(new BigNumber(8)).toNumber()] = ((_898_v13.f10) ? (_908_v23) : (_908_v23));
          _nw158[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_908_v23, _908_v23);
          _nw158[(new BigNumber(10)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(11)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(12)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(13)).toNumber()] = ((_898_v13.f10) ? (_908_v23) : (_908_v23));
          _nw158[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_898_v13.f10, _898_v13.f10, _898_v13.f10, _898_v13.f10, (_908_v23)[_module.__default.safeIndex(p2, new BigNumber((_908_v23).length))]);
          _nw158[(new BigNumber(15)).toNumber()] = _dafny.Seq.of((_909_v24).dtor_cf52, true);
          _nw158[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_908_v23, _dafny.Seq.of(_898_v13.f10));
          _nw158[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_881_v0, _898_v13.f10);
          _nw158[(new BigNumber(18)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_908_v23, _908_v23);
          _nw158[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(true);
          _nw158[(new BigNumber(21)).toNumber()] = _908_v23;
          _nw158[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(p0, _881_v0, _881_v0);
          _nw158[(new BigNumber(23)).toNumber()] = _908_v23;
          _910_v25 = _nw158;
          let _index148 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_910_v25).length));
          (_910_v25)[_index148] = _908_v23;
          let _911_v26;
          _911_v26 = _dafny.Map.Empty.slice().updateUnsafe(_898_v13.f10,_908_v23);
          let _index149 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_910_v25).length));
          (_910_v25)[_index149] = (((_911_v26).contains(((_898_v13.f10) ? (_898_v13.f10) : (_881_v0)))) ? ((_911_v26).get(((_898_v13.f10) ? (_898_v13.f10) : (_881_v0)))) : (_908_v23));
        } else {
          let _912_v27;
          let _nw159 = Array((new BigNumber(27)).toNumber()).fill([]);
          _912_v27 = _nw159;
          let _913_v28;
          let _nw160 = Array((new BigNumber(17)).toNumber()).fill(false);
          _913_v28 = _nw160;
          let _index150 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length));
          (_913_v28)[_index150] = _898_v13.f10;
          let _index151 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length));
          let _rhs180 = _912_v27;
          let _rhs181 = _this.f7;
          let _rhs182 = _898_v13;
          let _rhs183 = p0;
          let _lhs113 = globalState;
          let _lhs114 = _913_v28;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length));
          _912_v27 = _rhs180;
          _lhs113.f1 = _rhs181;
          _898_v13 = _rhs182;
          _lhs114[_lhs115] = _rhs183;
          let _914_v29;
          _914_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), ((_915_v0) => function (_916_i2) {
            return _module.D0.create_DC0(_dafny.Set.fromElements(_915_v0));
          })(_881_v0)));
          let _917_v30;
          _917_v30 = _dafny.Map.Empty.slice().updateUnsafe((_884_v3).f6,_898_v13.f10);
          let _918_v31;
          _918_v31 = _module.D1.create_DC4(_898_v13.f10);
          let _919_v32;
          _919_v32 = _module.D2.create_DC7((_913_v28)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length))], new BigNumber(257), _917_v30, _918_v31);
          let _920_v33;
          _920_v33 = _dafny.Set.fromElements(!((_913_v28)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length))]), (_919_v32).dtor_cf13);
          let _921_v34;
          _921_v34 = _module.D0.create_DC0(_920_v33);
          let _922_v35;
          _922_v35 = _dafny.Seq.of(_921_v34);
          let _923_v36;
          _923_v36 = _dafny.Seq.of(p0);
          let _index152 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length));
          let _rhs184 = (_module.D10.create_DC26((_913_v28)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length))])).dtor_cf52;
          let _rhs185 = (((_914_v29).contains(_922_v35)) ? ((_914_v29).get(_922_v35)) : (_884_v3.f7));
          let _rhs186 = _dafny.Seq.IsProperPrefixOf(_923_v36, _923_v36);
          let _lhs116 = _913_v28;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length));
          let _lhs118 = _884_v3;
          _lhs116[_lhs117] = _rhs184;
          _lhs118.f7 = _rhs185;
          _881_v0 = _rhs186;
          (_898_v13).f10 = (_913_v28)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_913_v28).length))];
          (_898_v13).f10 = !(!(!(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("oqk"), _883_v2))));
          let _924_v37;
          _924_v37 = _dafny.MultiSet.fromElements(((_884_v3).f6).multipliedBy(_884_v3.f7));
          (_884_v3).f7 = (((_924_v37).contains((_884_v3).f6)) ? ((_924_v37).get((_884_v3).f6)) : ((new BigNumber(288)).minus(_884_v3.f7)));
        }
        let _925_v38;
        let _nw161 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _925_v38 = _nw161;
        let _index153 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_925_v38).length));
        (_925_v38)[_index153] = _884_v3.f7;
        let _index154 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_925_v38).length));
        let _rhs187 = _this.f7;
        let _rhs188 = _module.__default.safeDivisionInt((_this.f7).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-376)), ((_926_v2) => function (_927_i3) {
          return _926_v2;
        })(_883_v2))).length)), p3);
        let _lhs119 = _884_v3;
        let _lhs120 = _925_v38;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_925_v38).length));
        _lhs119.f7 = _rhs187;
        _lhs120[_lhs121] = _rhs188;
        let _928_v39;
        _928_v39 = _module.D2.create_DC5(_925_v38);
        let _pat_let_tv12 = _925_v38;
        _928_v39 = function (_pat_let21_0) {
          return function (_929_dt__update__tmp_h0) {
            return function (_pat_let22_0) {
              return function (_930_dt__update_hcf9_h0) {
                return _module.D2.create_DC5(_930_dt__update_hcf9_h0);
              }(_pat_let22_0);
            }(_pat_let_tv12);
          }(_pat_let21_0);
        }(_928_v39);
        let _931_v40;
        _931_v40 = _dafny.Seq.of(_881_v0);
        let _932_v41;
        _932_v41 = _module.D9.create_DC23((_931_v40)[_module.__default.safeIndex((_884_v3).f6, new BigNumber((_931_v40).length))], new BigNumber(20), _881_v0, _881_v0, _884_v3.f7);
        let _index155 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_925_v38).length));
        (_925_v38)[_index155] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_932_v41).dtor_cf47));
      }
      let _933_v42;
      let _init29 = ((_934_v0) => function (_935_i4) {
        return !(_934_v0);
      })(_881_v0);
      let _nw162 = Array((new BigNumber(28)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw162.length); _i0_29++) {
        _nw162[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _933_v42 = _nw162;
      _933_v42 = _933_v42;
      let _936_v43;
      _936_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_881_v0)).length), p3)).length),_881_v0);
      let _937_v44;
      _937_v44 = _dafny.Map.Empty.slice().updateUnsafe((((_936_v43).contains(p3)) ? ((_936_v43).get(p3)) : (true)),p2);
      let _938_i5;
      _938_i5 = _dafny.ZERO;
      L1: {
        let _pat_let_tv13 = _881_v0;
        let _pat_let_tv14 = _881_v0;
        let _pat_let_tv15 = _881_v0;
        let _pat_let_tv16 = p0;
        while (function (_source9) {
          if (_source9.is_DC12) {
            let _944___mcc_h0 = (_source9).cf21;
            let _945_cf21 = _944___mcc_h0;
            if (_pat_let_tv13) {
              return !(_pat_let_tv14);
            } else {
              return _pat_let_tv15;
            }
          } else {
            let _946___mcc_h1 = (_source9).cf20;
            let _947_cf20 = _946___mcc_h1;
            return _pat_let_tv16;
          }
        }(_module.D5.create_DC12((((_937_v44).contains(p0)) ? ((_937_v44).get(p0)) : (p3))))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_938_i5)) {
              break L1;
            }
            _938_i5 = (_938_i5).plus(_dafny.ONE);
            let _939_v45;
            _939_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(706));
            let _940_v46;
            _940_v46 = _dafny.Set.fromElements(_884_v3.f7, _884_v3.f7);
            let _941_v47;
            _941_v47 = _dafny.Set.fromElements(p0);
            _939_v45 = (_939_v45).update(new BigNumber(-226), _module.__default.safeModuloInt(new BigNumber((_940_v46).length), new BigNumber((_941_v47).length)));
            _882_v1 = _882_v1;
            let _942_v49;
            _942_v49 = _dafny.MultiSet.fromElements(function () {
              let _coll31 = new _dafny.Map();
              for (const _compr_31 of (_936_v43).Keys.Elements) {
                let _943_v48 = _compr_31;
                if ((_936_v43).contains(_943_v48)) {
                  _coll31.push([_module.__default.safeModuloInt(_943_v48, new BigNumber(983)),p3]);
                }
              }
              return _coll31;
            }(), _939_v45, _939_v45);
            _942_v49 = _942_v49;
            _881_v0 = _881_v0;
          }
        }
      }
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f7 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6, f7) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm15(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(-928);
    };
    fm16(globalState) {
      let _this = this;
      return !(!(((_this).f6).isLessThan(_this.f7)));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _948_v0;
      _948_v0 = _dafny.Set.fromElements(p0);
      let _949_v1;
      _949_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _950_v2;
      _950_v2 = _dafny.Seq.UnicodeFromString("gqysp");
      let _951_v3;
      _951_v3 = _dafny.Seq.of(_949_v1, _949_v1, (_949_v1).update(p0, new BigNumber((_950_v2).length)), _949_v1, _949_v1);
      let _952_v4;
      _952_v4 = _dafny.Seq.of(p0);
      let _953_v5;
      _953_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(new BigNumber((_948_v0).length)).plus(new BigNumber(((_951_v3)[_module.__default.safeIndex(new BigNumber((_952_v4).length), new BigNumber((_951_v3).length))]).length)));
      let _954_v6;
      _954_v6 = new _dafny.CodePoint('a'.codePointAt(0));
      let _955_v7;
      _955_v7 = _dafny.Seq.of(_950_v2, _950_v2);
      let _956_v8;
      _956_v8 = _dafny.Seq.of(p1, _this.f7, _this.f7);
      let _957_v9;
      _957_v9 = _dafny.Seq.of(_956_v8);
      let _958_v10;
      _958_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
      let _959_v11;
      _959_v11 = _module.D1.create_DC4(p0);
      let _960_v12;
      _960_v12 = _module.D2.create_DC7(p0, _this.f7, _958_v10, _959_v11);
      let _pat_let_tv17 = p0;
      let _pat_let_tv18 = p0;
      let _pat_let_tv19 = p0;
      let _pat_let_tv20 = p0;
      let _pat_let_tv21 = _958_v10;
      let _rhs189 = _module.__default.fm17(_dafny.MultiSet.FromArray((_955_v7)[_module.__default.safeIndex(new BigNumber((_956_v8).length), new BigNumber((_955_v7).length))]), !(p0) || (p0), _this.f7, _957_v9, globalState);
      let _rhs190 = function (_source10) {
        if (_source10.is_DC6) {
          let _961___mcc_h0 = (_source10).cf10;
          let _962___mcc_h1 = (_source10).cf11;
          let _963___mcc_h2 = (_source10).cf12;
          let _964_cf12 = _963___mcc_h2;
          let _965_cf11 = _962___mcc_h1;
          let _966_cf10 = _961___mcc_h0;
          return _pat_let_tv17;
        } else if (_source10.is_DC7) {
          let _967___mcc_h3 = (_source10).cf13;
          let _968___mcc_h4 = (_source10).cf14;
          let _969___mcc_h5 = (_source10).cf15;
          let _970___mcc_h6 = (_source10).cf16;
          let _971_cf16 = _970___mcc_h6;
          let _972_cf15 = _969___mcc_h5;
          let _973_cf14 = _968___mcc_h4;
          let _974_cf13 = _967___mcc_h3;
          return true;
        } else if (_source10.is_DC5) {
          let _975___mcc_h7 = (_source10).cf9;
          let _976_cf9 = _975___mcc_h7;
          return _pat_let_tv18;
        } else {
          let _977___mcc_h8 = (_source10).cf17;
          let _978_cf17 = _977___mcc_h8;
          return !(_pat_let_tv19) || (_pat_let_tv20);
        }
      }(function (_pat_let23_0) {
        return function (_979_dt__update__tmp_h0) {
          return function (_pat_let24_0) {
            return function (_980_dt__update_hcf15_h0) {
              return _module.D2.create_DC7((_979_dt__update__tmp_h0).dtor_cf13, (_979_dt__update__tmp_h0).dtor_cf14, _980_dt__update_hcf15_h0, (_979_dt__update__tmp_h0).dtor_cf16);
            }(_pat_let24_0);
          }(_pat_let_tv21);
        }(_pat_let23_0);
      }(_960_v12));
      let _rhs191 = _954_v6;
      let _rhs192 = _module.__default.safeDivisionInt(new BigNumber(306), ((p0) ? (_this.f7) : (new BigNumber((_module.__default.fm18(_module.__default.fm0(p0, _953_v5, (_this).f6, globalState), p1, globalState)).length))));
      let _rhs193 = p0;
      let _lhs122 = _this;
      _953_v5 = _rhs189;
      r1 = _rhs190;
      _954_v6 = _rhs191;
      _lhs122.f7 = _rhs192;
      r1 = _rhs193;
      let _981_i0;
      _981_i0 = _dafny.ZERO;
      L2: {
        while (true) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_981_i0)) {
              break L2;
            }
            _981_i0 = (_981_i0).plus(_dafny.ONE);
            let _982_v13;
            let _983_v14;
            let _984_v15;
            let _out15;
            let _out16;
            let _out17;
            let _outcollector5 = (_this).m7(p0, (((_958_v10).contains(p1)) ? ((_958_v10).get(p1)) : (p0)), p2, globalState);
            _out15 = _outcollector5[0];
            _out16 = _outcollector5[1];
            _out17 = _outcollector5[2];
            _982_v13 = _out15;
            _983_v14 = _out16;
            _984_v15 = _out17;
            let _985_v16;
            let _nw163 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
            _985_v16 = _nw163;
            let _986_v17;
            _986_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _index156 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_985_v16).length));
            (_985_v16)[_index156] = new BigNumber(((_986_v17).Merge(_986_v17)).length);
            let _987_v18;
            _987_v18 = _dafny.MultiSet.fromElements(p0);
            let _index157 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_985_v16).length));
            (_985_v16)[_index157] = (((_987_v18).contains(p0)) ? ((_987_v18).get(p0)) : ((_this).f6));
            let _988_v19;
            _988_v19 = _dafny.MultiSet.fromElements(_module.__default.fm1(p0, globalState), ((_dafny.ZERO).minus((_dafny.ZERO).minus(p1))).minus(new BigNumber(-411)), (_this.f7).minus(p1));
            let _989_v20;
            _989_v20 = _987_v18;
            let _rhs194 = (p0) && (p0);
            let _rhs195 = _988_v19;
            let _rhs196 = ((!(!(p0))) ? (_module.__default.fm19(p0, globalState)) : ((_987_v18).update(true, _module.__default.abs(p2))));
            let _rhs197 = new BigNumber(217);
            let _lhs123 = globalState;
            r0 = _rhs194;
            _988_v19 = _rhs195;
            _989_v20 = _rhs196;
            _lhs123.f1 = _rhs197;
            r1 = p0;
          }
        }
      }
      _954_v6 = _954_v6;
      let _990_i1;
      _990_i1 = _dafny.ZERO;
      L3: {
        while (false) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_990_i1)) {
              break L3;
            }
            _990_i1 = (_990_i1).plus(_dafny.ONE);
            let _991_v21;
            let _nw164 = Array((new BigNumber(13)).toNumber()).fill(false);
            _991_v21 = _nw164;
            let _992_v22;
            _992_v22 = _module.D0.create_DC1(p0, _991_v21, p0);
            let _993_v23;
            let _nw165 = Array((_dafny.ONE).toNumber());
            _nw165[(_dafny.ZERO).toNumber()] = _992_v22;
            _993_v23 = _nw165;
            let _index158 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_993_v23).length));
            (_993_v23)[_index158] = _992_v22;
            let _index159 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_993_v23).length));
            (_993_v23)[_index159] = _module.D0.create_DC1(p0, _991_v21, p0);
            let _994_v24;
            _994_v24 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_954_v6);
            let _995_v25;
            _995_v25 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_994_v24).contains(p0)),p0);
            _995_v25 = (_995_v25).update(p0, p0);
            if ((p2).isLessThan((_dafny.ZERO).minus(p1))) {
              let _996_v26;
              let _init30 = function (_997_i2) {
                return (_997_i2).plus((_this).f6);
              };
              let _nw166 = Array((new BigNumber(17)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw166.length); _i0_30++) {
                _nw166[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _996_v26 = _nw166;
              let _index160 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_996_v26).length));
              (_996_v26)[_index160] = _this.f7;
              let _998_v27;
              _998_v27 = _dafny.Set.fromElements(p1, new BigNumber(792), p1);
              let _index161 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_996_v26).length));
              let _rhs198 = p0;
              let _rhs199 = p2;
              let _rhs200 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_956_v8, _956_v8), _dafny.Seq.of(new BigNumber((_module.__default.fm2(_956_v8, new BigNumber((_998_v27).length), _this.f7, p0, globalState)).length)))).length));
              let _rhs201 = _950_v2;
              let _rhs202 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_this.f7), new BigNumber((_dafny.Seq.Concat(_956_v8, _956_v8)).length));
              let _lhs124 = _this;
              let _lhs125 = _this;
              let _lhs126 = _996_v26;
              let _lhs127 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_996_v26).length));
              r1 = _rhs198;
              _lhs124.f7 = _rhs199;
              _lhs125.f7 = _rhs200;
              _950_v2 = _rhs201;
              _lhs126[_lhs127] = _rhs202;
              r1 = !(p0);
              _950_v2 = _950_v2;
              let _999_v28;
              let _nw167 = new _module.C0();
              _nw167.__ctor(true);
              _999_v28 = _nw167;
              let _1000_v29;
              _1000_v29 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(p0, p0, p0));
              let _1001_v30;
              _1001_v30 = _dafny.MultiSet.fromElements(p0, p0);
              let _1002_v31;
              _1002_v31 = _1001_v30;
              _1000_v29 = (_1000_v29).Union((_1000_v29).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_1002_v31))));
            } else {
              let _1003_v32;
              _1003_v32 = _dafny.Set.fromElements(p2);
              let _1004_v33;
              _1004_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1003_v32,_this.f7);
              _1004_v33 = (_1004_v33).update(_1003_v32, p2);
              let _index162 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length));
              (_991_v21)[_index162] = !(new BigNumber((_958_v10).length)).isEqualTo(p2);
              let _1005_v34;
              let _init31 = ((_1006_v4) => function (_1007_i3) {
                return _1006_v4;
              })(_952_v4);
              let _nw168 = Array((new BigNumber(6)).toNumber());
              for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw168.length); _i0_31++) {
                _nw168[_i0_31] = _init31(new BigNumber(_i0_31));
              }
              _1005_v34 = _nw168;
              let _index163 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length));
              let _rhs203 = new BigNumber(-542);
              let _rhs204 = p0;
              let _rhs205 = _1005_v34;
              let _lhs128 = globalState;
              let _lhs129 = _991_v21;
              let _lhs130 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length));
              _lhs128.f1 = _rhs203;
              _lhs129[_lhs130] = _rhs204;
              _1005_v34 = _rhs205;
              let _1008_v35;
              let _nw169 = new _module.C0();
              _nw169.__ctor(!(p0) || ((_991_v21)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length))]));
              _1008_v35 = _nw169;
              let _index164 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length));
              let _rhs206 = p0;
              let _rhs207 = _1008_v35.f10;
              let _rhs208 = _module.__default.fm0((_this).fm16(globalState), _953_v5, (_this).fm15(p1, new BigNumber(285), globalState), globalState);
              let _rhs209 = _1008_v35.f10;
              let _rhs210 = (_this).fm15(p2, p1, globalState);
              let _lhs131 = _991_v21;
              let _lhs132 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length));
              let _lhs133 = _1008_v35;
              let _lhs134 = _1008_v35;
              let _lhs135 = globalState;
              r1 = _rhs206;
              _lhs131[_lhs132] = _rhs207;
              _lhs133.f10 = _rhs208;
              _lhs134.f10 = _rhs209;
              _lhs135.f1 = _rhs210;
              let _1009_v36;
              _1009_v36 = _module.D0.create_DC0(_948_v0);
              let _1010_v37;
              _1010_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1009_v36,(_991_v21)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_991_v21).length))]);
              let _1011_v38;
              _1011_v38 = _module.D1.create_DC2(_1010_v37);
              let _1012_v39;
              _1012_v39 = _dafny.Set.fromElements(_1011_v38, _1011_v38, _module.D1.create_DC2(_1010_v37), _1011_v38);
              let _1013_v40;
              let _out18;
              _out18 = (_this).m6((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_949_v1).length),_1008_v35.f10)).Merge(_958_v10), p1, (_1012_v39).Union(_1012_v39), globalState);
              _1013_v40 = _out18;
            }
            _959_v11 = _959_v11;
          }
        }
      }
      if (!(new BigNumber((_950_v2).length)).isEqualTo(p1)) {
        let _1014_v41;
        let _nw170 = new _module.C0();
        _nw170.__ctor(p0);
        _1014_v41 = _nw170;
        _954_v6 = _module.__default.fm20(_950_v2, _1014_v41.f10, _this.f7, globalState);
        let _1015_v42;
        _1015_v42 = _dafny.MultiSet.fromElements(_958_v10, _958_v10);
        r1 = !((_1015_v42).equals(_1015_v42));
        let _1016_v43;
        let _nw171 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1016_v43 = _nw171;
        let _1017_v44;
        _1017_v44 = _dafny.MultiSet.fromElements(p2, new BigNumber((_950_v2).length), new BigNumber(556), p1, new BigNumber((_958_v10).length));
        let _index165 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1016_v43).length));
        (_1016_v43)[_index165] = (_1017_v44).Difference(_dafny.MultiSet.fromElements(new BigNumber((_950_v2).length), (_this).fm15((_this).f6, new BigNumber(156), globalState), new BigNumber((_1017_v44).cardinality())));
        let _index166 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_1016_v43).length));
        (_1016_v43)[_index166] = _dafny.MultiSet.FromArray(_956_v8);
        let _1018_v45;
        let _nw172 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1018_v45 = _nw172;
        let _1019_v46;
        _1019_v46 = _dafny.MultiSet.fromElements(_1018_v45);
        let _1020_v47;
        _1020_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.UnicodeFromString("mubemdm"));
        let _1021_v48;
        let _out19;
        _out19 = (_this).m6((_958_v10).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1019_v46).update(_1018_v45, _module.__default.abs(new BigNumber((_1020_v47).length)))).cardinality()),_1014_v41.f10)), p2, _dafny.Set.fromElements(_module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_948_v0),_1014_v41.f10))), globalState);
        _1021_v48 = _out19;
      } else {
        let _1022_v49;
        let _nw173 = Array((new BigNumber(21)).toNumber());
        _nw173[(_dafny.ZERO).toNumber()] = _954_v6;
        _nw173[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw173[(new BigNumber(2)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(3)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(4)).toNumber()] = _module.__default.fm20(_950_v2, p0, new BigNumber((_958_v10).length), globalState);
        _nw173[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _nw173[(new BigNumber(6)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(7)).toNumber()] = ((p0) ? (_954_v6) : (_954_v6));
        _nw173[(new BigNumber(8)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(9)).toNumber()] = (_950_v2)[_module.__default.safeIndex(p1, new BigNumber((_950_v2).length))];
        _nw173[(new BigNumber(10)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(11)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw173[(new BigNumber(13)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(14)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(15)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw173[(new BigNumber(17)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(18)).toNumber()] = _module.__default.fm20(_dafny.Seq.UnicodeFromString("kpvixhs"), p0, p2, globalState);
        _nw173[(new BigNumber(19)).toNumber()] = _954_v6;
        _nw173[(new BigNumber(20)).toNumber()] = ((p0) ? (_954_v6) : (_954_v6));
        _1022_v49 = _nw173;
        _1022_v49 = _1022_v49;
        let _1023_v50;
        let _nw174 = new _module.C0();
        _nw174.__ctor(p0);
        _1023_v50 = _nw174;
        let _1024_v51;
        _1024_v51 = ((p0) ? (_1023_v50) : (_1023_v50));
        let _source11 = _1024_v51;
        let _1025___mcc_h9 = _source11;
        let _1026_cf19 = _1025___mcc_h9;
        let _1027_v52;
        _1027_v52 = _dafny.Set.fromElements(p1);
        let _1028_v53;
        let _nw175 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _1028_v53 = _nw175;
        let _1029_v54;
        _1029_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1028_v53,new BigNumber((_953_v5).length));
        let _1030_v55;
        let _nw176 = Array((new BigNumber(14)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = new BigNumber((_1027_v52).length);
        _nw176[(_dafny.ONE).toNumber()] = _this.f7;
        _nw176[(new BigNumber(2)).toNumber()] = p2;
        _nw176[(new BigNumber(3)).toNumber()] = (new BigNumber(278)).plus(p2);
        _nw176[(new BigNumber(4)).toNumber()] = p1;
        _nw176[(new BigNumber(5)).toNumber()] = ((_this).f6).minus(p1);
        _nw176[(new BigNumber(6)).toNumber()] = _this.f7;
        _nw176[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw176[(new BigNumber(8)).toNumber()] = (_956_v8)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_956_v8).length))];
        _nw176[(new BigNumber(9)).toNumber()] = ((_this).f6).plus(p2);
        _nw176[(new BigNumber(10)).toNumber()] = (_this).f6;
        _nw176[(new BigNumber(11)).toNumber()] = new BigNumber(((_1029_v54).Merge(_1029_v54)).length);
        _nw176[(new BigNumber(12)).toNumber()] = _this.f7;
        _nw176[(new BigNumber(13)).toNumber()] = ((_this).f6).plus(p2);
        _1030_v55 = _nw176;
        let _index167 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1028_v53).length));
        (_1028_v53)[_index167] = (_this).fm15(p2, _this.f7, globalState);
        let _1031_v57;
        let _nw177 = Array((new BigNumber(11)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = _1023_v50.f10;
        _nw177[(_dafny.ONE).toNumber()] = (_this).fm16(globalState);
        _nw177[(new BigNumber(2)).toNumber()] = _1023_v50.f10;
        _nw177[(new BigNumber(3)).toNumber()] = !(_1026_cf19.f10);
        _nw177[(new BigNumber(4)).toNumber()] = ((_1023_v50.f10) ? (_1023_v50.f10) : (p0));
        _nw177[(new BigNumber(5)).toNumber()] = !(_1023_v50.f10) || (_1026_cf19.f10);
        _nw177[(new BigNumber(6)).toNumber()] = (p0) === (p0);
        _nw177[(new BigNumber(7)).toNumber()] = !(_1023_v50.f10) || (_module.__default.fm0(_1026_cf19.f10, function () {
          let _coll32 = new _dafny.Map();
          for (const _compr_32 of _dafny.IntegerRange(new BigNumber(283), new BigNumber(-466))) {
            let _1032_v56 = _compr_32;
            if (((new BigNumber(283)).isLessThanOrEqualTo(_1032_v56)) && ((_1032_v56).isLessThan(new BigNumber(-466)))) {
              _coll32.push([_module.__default.safeModuloInt(_1032_v56, _this.f7),new BigNumber((_948_v0).length)]);
            }
          }
          return _coll32;
        }(), (_this).f6, globalState));
        _nw177[(new BigNumber(8)).toNumber()] = _1026_cf19.f10;
        _nw177[(new BigNumber(9)).toNumber()] = _1026_cf19.f10;
        _nw177[(new BigNumber(10)).toNumber()] = !(_1026_cf19.f10);
        _1031_v57 = _nw177;
        let _index168 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1031_v57).length));
        (_1031_v57)[_index168] = _1026_cf19.f10;
        let _index169 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1028_v53).length));
        let _index170 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1031_v57).length));
        let _rhs211 = p2;
        let _rhs212 = _1023_v50.f10;
        let _rhs213 = (p2).minus((p2).plus(p1));
        let _rhs214 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_1030_v55, _1030_v55, _1028_v53, _1030_v55)).length), p2);
        let _lhs136 = _1028_v53;
        let _lhs137 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1028_v53).length));
        let _lhs138 = _1031_v57;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1031_v57).length));
        let _lhs140 = globalState;
        let _lhs141 = globalState;
        _lhs136[_lhs137] = _rhs211;
        _lhs138[_lhs139] = _rhs212;
        _lhs140.f1 = _rhs213;
        _lhs141.f1 = _rhs214;
        let _index171 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1031_v57).length));
        (_1031_v57)[_index171] = (false) && (_module.__default.fm0((_1031_v57)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_1031_v57).length))], _953_v5, _this.f7, globalState));
        _1023_v50 = _1026_cf19;
        let _1033_v58;
        let _nw178 = new _module.C0();
        _nw178.__ctor(!_dafny.areEqual(_950_v2, _950_v2));
        _1033_v58 = _nw178;
        let _1034_v59;
        _1034_v59 = _module.D0.create_DC0(_948_v0);
        let _1035_v60;
        _1035_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1034_v59,_1023_v50.f10);
        let _1036_v61;
        _1036_v61 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1035_v60);
        let _1037_v62;
        _1037_v62 = _module.D1.create_DC2((((_1036_v61).contains(_1023_v50.f10)) ? ((_1036_v61).get(_1023_v50.f10)) : (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_948_v0),false))));
        _1037_v62 = _1037_v62;
        let _1038_v63;
        _1038_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1023_v50.f10);
        let _1039_v64;
        _1039_v64 = _dafny.Seq.of(_1038_v63);
        let _1040_v65;
        _1040_v65 = _dafny.MultiSet.fromElements(_this.f7, p1);
        let _1041_v66;
        let _init32 = ((_1042_v50) => function (_1043_i4) {
          return _1042_v50.f10;
        })(_1023_v50);
        let _nw179 = Array((new BigNumber(20)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw179.length); _i0_32++) {
          _nw179[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1041_v66 = _nw179;
        let _1044_v67;
        _1044_v67 = _module.D0.create_DC1(_1023_v50.f10, _1041_v66, _1023_v50.f10);
        let _1045_v68;
        _1045_v68 = _dafny.MultiSet.fromElements(_1023_v50.f10, true, p0);
        let _1046_v69;
        let _nw180 = Array((new BigNumber(27)).toNumber());
        _nw180[(_dafny.ZERO).toNumber()] = ((_1023_v50.f10) ? (p0) : (true));
        _nw180[(_dafny.ONE).toNumber()] = p0;
        _nw180[(new BigNumber(2)).toNumber()] = !((p1).isEqualTo(new BigNumber((((_1039_v64)[_module.__default.safeIndex((_this).f6, new BigNumber((_1039_v64).length))]).update(_1023_v50.f10, _1023_v50.f10)).length)));
        _nw180[(new BigNumber(3)).toNumber()] = (_960_v12).dtor_cf13;
        _nw180[(new BigNumber(4)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(5)).toNumber()] = !(p0);
        _nw180[(new BigNumber(6)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(7)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(8)).toNumber()] = false;
        _nw180[(new BigNumber(9)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(10)).toNumber()] = (_948_v0).IsSubsetOf(_948_v0);
        _nw180[(new BigNumber(11)).toNumber()] = p0;
        _nw180[(new BigNumber(12)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(13)).toNumber()] = p0;
        _nw180[(new BigNumber(14)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(15)).toNumber()] = (_1040_v65).IsProperSubsetOf(_dafny.MultiSet.FromArray(_956_v8));
        _nw180[(new BigNumber(16)).toNumber()] = _1023_v50.f10;
        _nw180[(new BigNumber(17)).toNumber()] = (p2).isLessThanOrEqualTo(p2);
        _nw180[(new BigNumber(18)).toNumber()] = p0;
        _nw180[(new BigNumber(19)).toNumber()] = (false) || (true);
        _nw180[(new BigNumber(20)).toNumber()] = !(((_1023_v50.f10) ? (_1023_v50.f10) : ((_1044_v67).dtor_cf3)));
        _nw180[(new BigNumber(21)).toNumber()] = (_1045_v68).equals(_dafny.MultiSet.fromElements(_module.__default.fm0(_1023_v50.f10, _dafny.Map.Empty.slice().updateUnsafe(p2,p2), p1, globalState)));
        _nw180[(new BigNumber(22)).toNumber()] = _module.__default.fm0(_1023_v50.f10, _953_v5, p1, globalState);
        _nw180[(new BigNumber(23)).toNumber()] = p0;
        _nw180[(new BigNumber(24)).toNumber()] = p0;
        _nw180[(new BigNumber(25)).toNumber()] = (_this).fm16(globalState);
        _nw180[(new BigNumber(26)).toNumber()] = _1023_v50.f10;
        _1046_v69 = _nw180;
        let _index172 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_1046_v69).length));
        (_1046_v69)[_index172] = p0;
        let _1047_v70;
        _1047_v70 = _dafny.Set.fromElements((_this).f6);
        let _1048_v72;
        let _nw181 = Array((new BigNumber(16)).toNumber()).fill(_module.D2.Default());
        _1048_v72 = _nw181;
        let _1049_v73;
        _1049_v73 = _dafny.Set.fromElements(_1048_v72, _1048_v72);
        let _index173 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_1046_v69).length));
        let _rhs215 = (((_1047_v70).IsDisjointFrom(function () {
          let _coll33 = new _dafny.Set();
          for (const _compr_33 of _dafny.IntegerRange(new BigNumber(288), new BigNumber(394))) {
            let _1050_v71 = _compr_33;
            if (((new BigNumber(288)).isLessThanOrEqualTo(_1050_v71)) && ((_1050_v71).isLessThan(new BigNumber(394)))) {
              _coll33.add(_module.__default.safeDivisionInt(_1050_v71, p1));
            }
          }
          return _coll33;
        }())) ? (p0) : (true));
        let _rhs216 = !(((_1049_v73).Union(_1049_v73)).IsSubsetOf((_1049_v73).Intersect(_1049_v73)));
        let _rhs217 = (_dafny.ZERO).minus((p2).minus(new BigNumber((_953_v5).length)));
        let _lhs142 = _1046_v69;
        let _lhs143 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_1046_v69).length));
        let _lhs144 = _this;
        _lhs142[_lhs143] = _rhs215;
        r1 = _rhs216;
        _lhs144.f7 = _rhs217;
        r0 = p0;
      }
      if ((_this).fm16(globalState)) {
        _949_v1 = (_949_v1).update(p0, (new BigNumber((_950_v2).length)).multipliedBy((_dafny.ZERO).minus(p2)));
        let _rhs218 = ((true) ? (_952_v4) : (_952_v4));
        let _rhs219 = _952_v4;
        _952_v4 = _rhs218;
        _952_v4 = _rhs219;
        r2 = p2;
        let _1051_v74;
        _1051_v74 = _dafny.Map.Empty.slice().updateUnsafe(_954_v6,p2);
        r1 = (_952_v4)[_module.__default.safeIndex(((((_1051_v74).contains(_954_v6)) ? ((_1051_v74).get(_954_v6)) : (new BigNumber((_956_v8).length)))).plus(p1), new BigNumber((_952_v4).length))];
        let _1052_v75;
        _1052_v75 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1053_v76;
        _1053_v76 = _dafny.Set.fromElements(_952_v4);
        let _1054_v77;
        let _nw182 = Array((new BigNumber(11)).toNumber());
        _nw182[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).fm15(p1, p1, globalState), (_dafny.ZERO).minus(new BigNumber((_1052_v75).length)));
        _nw182[(_dafny.ONE).toNumber()] = new BigNumber((_1053_v76).length);
        _nw182[(new BigNumber(2)).toNumber()] = new BigNumber(886);
        _nw182[(new BigNumber(3)).toNumber()] = p1;
        _nw182[(new BigNumber(4)).toNumber()] = p2;
        _nw182[(new BigNumber(5)).toNumber()] = p1;
        _nw182[(new BigNumber(6)).toNumber()] = (p2).multipliedBy((_956_v8)[_module.__default.safeIndex(p2, new BigNumber((_956_v8).length))]);
        _nw182[(new BigNumber(7)).toNumber()] = _this.f7;
        _nw182[(new BigNumber(8)).toNumber()] = p1;
        _nw182[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw182[(new BigNumber(10)).toNumber()] = p2;
        _1054_v77 = _nw182;
        let _index174 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1054_v77).length));
        (_1054_v77)[_index174] = p2;
        let _index175 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1054_v77).length));
        let _rhs220 = false;
        let _rhs221 = ((((_949_v1).contains(p0)) ? ((_949_v1).get(p0)) : (_this.f7))).multipliedBy(_module.__default.safeDivisionInt((_this).f6, (_this).f6));
        let _lhs145 = _1054_v77;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1054_v77).length));
        r1 = _rhs220;
        _lhs145[_lhs146] = _rhs221;
      } else {
        let _1055_v78;
        let _nw183 = new _module.C0();
        _nw183.__ctor(p0);
        _1055_v78 = _nw183;
        let _1056_v79;
        _1056_v79 = _1055_v78;
        let _rhs222 = _1055_v78;
        let _rhs223 = p0;
        _1056_v79 = _rhs222;
        r1 = _rhs223;
        let _1057_v80;
        let _init33 = ((_1058_v4) => function (_1059_i5) {
          return _1058_v4;
        })(_952_v4);
        let _nw184 = Array((new BigNumber(20)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw184.length); _i0_33++) {
          _nw184[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1057_v80 = _nw184;
        let _index176 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1057_v80).length));
        (_1057_v80)[_index176] = _952_v4;
        let _index177 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1057_v80).length));
        (_1057_v80)[_index177] = _dafny.Seq.Concat(_952_v4, _952_v4);
        let _1060_v81;
        _1060_v81 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm21(globalState));
        _958_v10 = (_958_v10).update(_this.f7, _dafny.Seq.IsProperPrefixOf((((_1060_v81).contains(_this.f7)) ? ((_1060_v81).get(_this.f7)) : ((_1057_v80)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_1057_v80).length))])), (_1057_v80)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_1057_v80).length))]));
        let _1061_v82;
        let _nw185 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1061_v82 = _nw185;
        let _index178 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1061_v82).length));
        (_1061_v82)[_index178] = p0;
        let _index179 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1061_v82).length));
        (_1061_v82)[_index179] = !((_this).f6).isEqualTo((_dafny.ZERO).minus(_module.__default.fm1(_1055_v78.f10, globalState)));
        let _1062_v83;
        let _nw186 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
        _1062_v83 = _nw186;
        let _1063_v84;
        _1063_v84 = _dafny.MultiSet.fromElements(_954_v6, _954_v6, _954_v6, _954_v6, _954_v6);
        let _index180 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_1062_v83).length));
        (_1062_v83)[_index180] = _dafny.Map.Empty.slice().updateUnsafe(_1063_v84,(_1061_v82)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1061_v82).length))]);
        let _1064_v85;
        _1064_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v84,!(_1055_v78.f10));
        let _index181 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_1062_v83).length));
        let _rhs224 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_1065_v6) => function (_1066_i6) {
          return _1065_v6;
        })(_954_v6))).length));
        let _rhs225 = _1055_v78.f10;
        let _rhs226 = _954_v6;
        let _rhs227 = (_dafny.Map.Empty.slice().updateUnsafe(_1063_v84,true)).Merge(_1064_v85);
        let _lhs147 = globalState;
        let _lhs148 = _1055_v78;
        let _lhs149 = _1062_v83;
        let _lhs150 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_1062_v83).length));
        _lhs147.f1 = _rhs224;
        _lhs148.f10 = _rhs225;
        _954_v6 = _rhs226;
        _lhs149[_lhs150] = _rhs227;
      }
      r0 = p0;
      let _1067_v86;
      _1067_v86 = _dafny.Set.fromElements((_this).f6, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("oh"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("oh")).length)), new _dafny.CodePoint('f'.codePointAt(0)))).length));
      r1 = (_1067_v86).IsDisjointFrom(_1067_v86);
      r2 = ((_this).f6).plus(new BigNumber(324));
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _1068_v0;
      _1068_v0 = false;
      let _1069_i0;
      _1069_i0 = _dafny.ZERO;
      L4: {
        while (!(_1068_v0)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1069_i0)) {
              break L4;
            }
            _1069_i0 = (_1069_i0).plus(_dafny.ONE);
            let _1070_v1;
            _1070_v1 = _dafny.Seq.UnicodeFromString("ea");
            let _1071_v2;
            let _nw187 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
            _1071_v2 = _nw187;
            let _1072_v3;
            _1072_v3 = _module.D2.create_DC5(_1071_v2);
            let _1073_v4;
            let _nw188 = Array((new BigNumber(20)).toNumber());
            _nw188[(_dafny.ZERO).toNumber()] = _1071_v2;
            _nw188[(_dafny.ONE).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(2)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(3)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(4)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(5)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(6)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(7)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(8)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(9)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(10)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(11)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(12)).toNumber()] = (_1072_v3).dtor_cf9;
            _nw188[(new BigNumber(13)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(14)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(15)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(16)).toNumber()] = (_1072_v3).dtor_cf9;
            _nw188[(new BigNumber(17)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(18)).toNumber()] = _1071_v2;
            _nw188[(new BigNumber(19)).toNumber()] = _1071_v2;
            _1073_v4 = _nw188;
            let _index182 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1073_v4).length));
            (_1073_v4)[_index182] = ((_1068_v0) ? (_1071_v2) : (_1071_v2));
            let _1074_v5;
            _1074_v5 = new _dafny.CodePoint('g'.codePointAt(0));
            let _1075_v6;
            _1075_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm18(_1068_v0, (_this).f6, globalState)).length),_1074_v5);
            let _1076_v7;
            _1076_v7 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1075_v6).length));
            let _1077_v8;
            _1077_v8 = _dafny.Seq.of(_this.f7);
            let _1078_v9;
            _1078_v9 = _dafny.Seq.of(p1, new BigNumber(938), new BigNumber((_1077_v8).length));
            let _1079_v10;
            _1079_v10 = _dafny.Seq.of(_1068_v0);
            let _index183 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1073_v4).length));
            let _rhs228 = (p1).plus((new BigNumber((_1076_v7).length)).plus((_this).f6));
            let _rhs229 = !(!(!(!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_module.__default.fm1(_1068_v0, globalState)), _dafny.Seq.of(new BigNumber((_1078_v9).length)))))));
            let _rhs230 = _dafny.Seq.Concat(_1070_v1, _1070_v1);
            let _rhs231 = _1071_v2;
            let _rhs232 = (_1079_v10)[_module.__default.safeIndex(p1, new BigNumber((_1079_v10).length))];
            let _lhs151 = globalState;
            let _lhs152 = _1073_v4;
            let _lhs153 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1073_v4).length));
            _lhs151.f1 = _rhs228;
            r0 = _rhs229;
            _1070_v1 = _rhs230;
            _lhs152[_lhs153] = _rhs231;
            r0 = _rhs232;
            _1076_v7 = (_1076_v7).update(_1068_v0, (_this).f6);
            _1068_v0 = _1068_v0;
            let _1080_v11;
            let _nw189 = new _module.C0();
            _nw189.__ctor(_1068_v0);
            _1080_v11 = _nw189;
            _1080_v11 = _1080_v11;
          }
        }
      }
      let _1081_v12;
      _1081_v12 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tayy")).length));
      let _1082_v13;
      _1082_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1068_v0,_1068_v0);
      let _1083_v14;
      _1083_v14 = _dafny.Map.Empty.slice().updateUnsafe((((_1082_v13).contains(_1068_v0)) ? ((_1082_v13).get(_1068_v0)) : ((_this).fm16(globalState))),_1081_v12);
      let _1084_v16;
      _1084_v16 = _dafny.Seq.of(!(!(_1068_v0)), (_this).fm16(globalState));
      let _1085_v17;
      let _nw190 = Array((new BigNumber(10)).toNumber()).fill(false);
      _1085_v17 = _nw190;
      let _1086_v18;
      _1086_v18 = _module.D0.create_DC1(_1068_v0, _1085_v17, !(_1068_v0));
      let _1087_v19;
      _1087_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _1088_v20;
      let _init34 = function (_1089_i1) {
        return _module.__default.safeModuloInt(_1089_i1, (_this).f6);
      };
      let _nw191 = Array((new BigNumber(6)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw191.length); _i0_34++) {
        _nw191[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1088_v20 = _nw191;
      let _1090_v21;
      _1090_v21 = _dafny.MultiSet.fromElements(_1088_v20, _1088_v20);
      let _1091_v22;
      let _nw192 = Array((new BigNumber(21)).toNumber());
      _nw192[(_dafny.ZERO).toNumber()] = _1068_v0;
      _nw192[(_dafny.ONE).toNumber()] = ((_this).f6).isLessThan(_this.f7);
      _nw192[(new BigNumber(2)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(3)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(4)).toNumber()] = (_1081_v12).IsDisjointFrom((((_1083_v14).contains(_1068_v0)) ? ((_1083_v14).get(_1068_v0)) : (function () {
        let _coll34 = new _dafny.Set();
        for (const _compr_34 of _dafny.IntegerRange(new BigNumber(-142), new BigNumber(168))) {
          let _1092_v15 = _compr_34;
          if (((new BigNumber(-142)).isLessThanOrEqualTo(_1092_v15)) && ((_1092_v15).isLessThan(new BigNumber(168)))) {
            _coll34.add((_1092_v15).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("qkk")).length)));
          }
        }
        return _coll34;
      }())));
      _nw192[(new BigNumber(5)).toNumber()] = !(_1068_v0);
      _nw192[(new BigNumber(6)).toNumber()] = (_1084_v16)[_module.__default.safeIndex(p1, new BigNumber((_1084_v16).length))];
      _nw192[(new BigNumber(7)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(8)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(9)).toNumber()] = !(!((_1086_v18).dtor_cf3)) || (!(_1068_v0));
      _nw192[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1084_v16, _1084_v16);
      _nw192[(new BigNumber(11)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(12)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(13)).toNumber()] = !((((_1087_v19).contains(p1)) ? ((_1087_v19).get(p1)) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1068_v0)))).contains(p1);
      _nw192[(new BigNumber(14)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(15)).toNumber()] = (((p0).contains(p1)) ? ((p0).get(p1)) : (_1068_v0));
      _nw192[(new BigNumber(16)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(17)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(18)).toNumber()] = _1068_v0;
      _nw192[(new BigNumber(19)).toNumber()] = (_1090_v21).IsSubsetOf(_1090_v21);
      _nw192[(new BigNumber(20)).toNumber()] = _1068_v0;
      _1091_v22 = _nw192;
      let _index184 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
      (_1091_v22)[_index184] = false;
      let _index185 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1091_v22).length));
      (_1091_v22)[_index185] = _1068_v0;
      let _1093_v23;
      _1093_v23 = _dafny.Seq.UnicodeFromString("w");
      let _1094_v24;
      _1094_v24 = _dafny.MultiSet.fromElements(_1068_v0, _1068_v0, !(!(_1068_v0)));
      let _1095_v25;
      _1095_v25 = _dafny.Set.fromElements(_1068_v0, (_1084_v16)[_module.__default.safeIndex(new BigNumber(723), new BigNumber((_1084_v16).length))]);
      let _1096_v26;
      _1096_v26 = _module.D0.create_DC0(_1095_v25);
      let _pat_let_tv22 = _1082_v13;
      let _pat_let_tv23 = _1068_v0;
      let _pat_let_tv24 = _1082_v13;
      let _pat_let_tv25 = _1068_v0;
      let _pat_let_tv26 = _1068_v0;
      let _index186 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
      let _index187 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1091_v22).length));
      let _rhs233 = !(((((_dafny.MultiSet.FromArray(_1084_v16)).update(_1068_v0, _module.__default.abs(new BigNumber((_1093_v23).length)))).update(_1068_v0, _module.__default.abs(p1))).update(_1068_v0, _module.__default.abs(new BigNumber(144)))).equals((_module.__default.fm22(p1, globalState)).Intersect(_1094_v24)));
      let _rhs234 = function (_source12) {
        if (_source12.is_DC1) {
          let _1097___mcc_h0 = (_source12).cf1;
          let _1098___mcc_h1 = (_source12).cf2;
          let _1099___mcc_h2 = (_source12).cf3;
          let _1100_cf3 = _1099___mcc_h2;
          let _1101_cf2 = _1098___mcc_h1;
          let _1102_cf1 = _1097___mcc_h0;
          if ((_pat_let_tv22).contains(_pat_let_tv23)) {
            return (_pat_let_tv24).get(_pat_let_tv25);
          } else {
            return _pat_let_tv26;
          }
        } else {
          let _1103___mcc_h3 = (_source12).cf0;
          let _1104_cf0 = _1103___mcc_h3;
          return false;
        }
      }(_1096_v26);
      let _rhs235 = (_this).f6;
      let _lhs154 = _1091_v22;
      let _lhs155 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
      let _lhs156 = _1091_v22;
      let _lhs157 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1091_v22).length));
      let _lhs158 = _this;
      _lhs154[_lhs155] = _rhs233;
      _lhs156[_lhs157] = _rhs234;
      _lhs158.f7 = _rhs235;
      let _pat_let_tv27 = _1093_v23;
      _1093_v23 = function (_source13) {
        if (_source13.is_DC1) {
          let _1105___mcc_h4 = (_source13).cf1;
          let _1106___mcc_h5 = (_source13).cf2;
          let _1107___mcc_h6 = (_source13).cf3;
          let _1108_cf3 = _1107___mcc_h6;
          let _1109_cf2 = _1106___mcc_h5;
          let _1110_cf1 = _1105___mcc_h4;
          return _pat_let_tv27;
        } else {
          let _1111___mcc_h7 = (_source13).cf0;
          let _1112_cf0 = _1111___mcc_h7;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(979)), function (_1113_i2) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          });
        }
      }(_1096_v26);
      let _hi10 = p1;
      for (let _1114_i3 = _module.__default.safeDivisionInt((_this).f6, p1); _1114_i3.isLessThan(_hi10); _1114_i3 = _1114_i3.plus(_dafny.ONE)) {
        let _1115_v27;
        _1115_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,false);
        _1115_v27 = (_1115_v27).update(new BigNumber((((true) ? (_1081_v12) : (_1081_v12))).length), (_1091_v22)[_module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length))]);
        _1068_v0 = false;
        _1068_v0 = (_1091_v22)[_module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length))];
        let _1116_v28;
        _1116_v28 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1117_v29;
        _1117_v29 = _dafny.Seq.of(_1095_v25);
        let _index188 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
        let _index189 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
        let _rhs236 = !((_1091_v22)[_module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length))]);
        let _rhs237 = !(!_dafny.Seq.contains(_1093_v23, _1116_v28));
        let _rhs238 = (_this).f6;
        let _rhs239 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_this).f6, _this.f7), _1114_i3);
        let _rhs240 = ((_1117_v29)[_module.__default.safeIndex(_1114_i3, new BigNumber((_1117_v29).length))]).IsProperSubsetOf((_1095_v25).Difference(_1095_v25));
        let _lhs159 = _1091_v22;
        let _lhs160 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
        let _lhs161 = globalState;
        let _lhs162 = globalState;
        let _lhs163 = _1091_v22;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
        r0 = _rhs236;
        _lhs159[_lhs160] = _rhs237;
        _lhs161.f1 = _rhs238;
        _lhs162.f1 = _rhs239;
        _lhs163[_lhs164] = _rhs240;
      }
      let _hi11 = p1;
      for (let _1118_i4 = p1; _1118_i4.isLessThan(_hi11); _1118_i4 = _1118_i4.plus(_dafny.ONE)) {
        let _1119_v30;
        _1119_v30 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1120_v31;
        _1120_v31 = _dafny.Seq.of(_1095_v25, _1095_v25);
        let _1121_v32;
        let _nw193 = new _module.C1();
        _nw193.__ctor(_1119_v30, _module.__default.safeModuloInt((_this).f6, new BigNumber(95)), _module.__default.safeModuloInt(new BigNumber(((_1120_v31)[_module.__default.safeIndex((_this).f6, new BigNumber((_1120_v31).length))]).length), _1118_i4));
        _1121_v32 = _nw193;
        let _1122_v33;
        _1122_v33 = _dafny.Map.Empty.slice().updateUnsafe(true,_1121_v32);
        _1121_v32 = (((_1122_v33).contains((_1121_v32.f7).isLessThan(p1))) ? ((_1122_v33).get((_1121_v32.f7).isLessThan(p1))) : (_1121_v32));
        let _index190 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1088_v20).length));
        (_1088_v20)[_index190] = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_1119_v30)).cardinality()));
        let _index191 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1088_v20).length));
        let _index192 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
        let _rhs241 = _1068_v0;
        let _rhs242 = _1118_i4;
        let _rhs243 = _1088_v20;
        let _rhs244 = _this.f7;
        let _rhs245 = (_1091_v22)[_module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length))];
        let _lhs165 = globalState;
        let _lhs166 = _1088_v20;
        let _lhs167 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1088_v20).length));
        let _lhs168 = _1091_v22;
        let _lhs169 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length));
        _1068_v0 = _rhs241;
        _lhs165.f1 = _rhs242;
        _1088_v20 = _rhs243;
        _lhs166[_lhs167] = _rhs244;
        _lhs168[_lhs169] = _rhs245;
        let _1123_v34;
        let _init35 = ((_1124_v23) => function (_1125_i5) {
          return _1124_v23;
        })(_1093_v23);
        let _nw194 = Array((new BigNumber(21)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw194.length); _i0_35++) {
          _nw194[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1123_v34 = _nw194;
        let _index193 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_1123_v34).length));
        (_1123_v34)[_index193] = _1093_v23;
        let _index194 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_1123_v34).length));
        (_1123_v34)[_index194] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xr"), _1093_v23), _dafny.Seq.Concat(_1093_v23, _1093_v23));
        let _1126_v35;
        let _nw195 = new _module.C2();
        _nw195.__ctor(_module.__default.safeModuloInt(new BigNumber(256), (_1121_v32).f6), p1);
        _1126_v35 = _nw195;
      }
      let _index195 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1088_v20).length));
      (_1088_v20)[_index195] = (_this).f6;
      let _index196 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1088_v20).length));
      (_1088_v20)[_index196] = p1;
      let _1127_v36;
      _1127_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-348),(_this).f6);
      let _1128_v37;
      _1128_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1127_v36,(_1091_v22)[_module.__default.safeIndex(new BigNumber(143), new BigNumber((_1091_v22).length))]);
      r0 = (((_1128_v37).contains(_1127_v36)) ? ((_1128_v37).get(_1127_v36)) : (_module.__default.fm0(_1068_v0, _1127_v36, p1, globalState)));
      return r0;
    }
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _1129_v0;
      let _nw196 = Array((new BigNumber(5)).toNumber()).fill(false);
      _1129_v0 = _nw196;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1129_v0).length))) {
        let _1130_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1130_i0)) && ((_1130_i0).isLessThan(new BigNumber((_1129_v0).length))))) {
          (_1129_v0)[(_1130_i0)] = p1;
        }
      }
      let _1131_i1;
      _1131_i1 = _dafny.ZERO;
      L5: {
        while (p1) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1131_i1)) {
              break L5;
            }
            _1131_i1 = (_1131_i1).plus(_dafny.ONE);
            let _1132_v1;
            _1132_v1 = false;
            let _1133_v2;
            _1133_v2 = _dafny.Seq.of(_1132_v1);
            let _rhs246 = _1129_v0;
            let _rhs247 = p0;
            let _rhs248 = _dafny.Seq.of(!(p0) || (true), !((new BigNumber((_dafny.Seq.UnicodeFromString("vgkqtd")).length)).isLessThan((_this).f6)), p0, p1);
            let _rhs249 = ((p2).plus(p2)).isEqualTo((_this).f6);
            _1129_v0 = _rhs246;
            _1132_v1 = _rhs247;
            _1133_v2 = _rhs248;
            _1132_v1 = _rhs249;
            let _1134_v3;
            let _init36 = function (_1135_i2) {
              return _module.__default.safeModuloInt(_1135_i2, (_this).f6);
            };
            let _nw197 = Array((new BigNumber(29)).toNumber());
            for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw197.length); _i0_36++) {
              _nw197[_i0_36] = _init36(new BigNumber(_i0_36));
            }
            _1134_v3 = _nw197;
            let _1136_v4;
            _1136_v4 = _dafny.Seq.of(_1134_v3);
            let _1137_v5;
            _1137_v5 = _dafny.Map.Empty.slice().updateUnsafe((_1136_v4)[_module.__default.safeIndex(new BigNumber(-925), new BigNumber((_1136_v4).length))],_1132_v1);
            _1137_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1134_v3,_1132_v1);
            let _1138_v6;
            _1138_v6 = _module.D0.create_DC1(true, _1129_v0, _1132_v1);
            let _1139_v7;
            _1139_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(273),_1132_v1);
            let _1140_v8;
            let _1141_v9;
            let _1142_v10;
            let _out20;
            let _out21;
            let _out22;
            let _outcollector6 = _module.__default.m0(_1138_v6, _module.__default.fm1(!((((_1139_v7).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("blo")).length))))) ? ((_1139_v7).get((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("blo")).length))))) : (false))), globalState), globalState);
            _out20 = _outcollector6[0];
            _out21 = _outcollector6[1];
            _out22 = _outcollector6[2];
            _1140_v8 = _out20;
            _1141_v9 = _out21;
            _1142_v10 = _out22;
            let _1143_v11;
            _1143_v11 = _dafny.MultiSet.fromElements(p0);
            let _1144_v12;
            _1144_v12 = _dafny.Seq.of(new BigNumber((_1143_v11).cardinality()), _1140_v8);
            let _1145_v13;
            _1145_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1144_v12)[_module.__default.safeIndex((_this).f6, new BigNumber((_1144_v12).length))],new BigNumber(144));
            let _1146_v14;
            _1146_v14 = _dafny.Seq.UnicodeFromString("aywp");
            if (_module.__default.fm0(_1132_v1, _1145_v13, new BigNumber((_1146_v14).length), globalState)) {
              let _1147_v15;
              _1147_v15 = _dafny.Set.fromElements((_this).f6, new BigNumber(-188));
              let _1148_v16;
              let _nw198 = new _module.C1();
              _nw198.__ctor(new _dafny.CodePoint('s'.codePointAt(0)), new BigNumber((_1147_v15).length), new BigNumber((_module.__default.fm36(_1140_v8, true, (_this).f6, _1141_v9, globalState)).length));
              _1148_v16 = _nw198;
              (globalState).f1 = new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1146_v14), _module.__default.safeIndex(_this.f7, new BigNumber((_dafny.Seq.of(_1146_v14)).length)), _1146_v14)).length);
              let _1149_v17;
              _1149_v17 = _module.D9.create_DC22(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-814)), ((_1150_v8) => function (_1151_i3) {
  return _1150_v8;
})(_1140_v8)), _1144_v12));
              _1149_v17 = _1149_v17;
              let _1152_v18;
              let _nw199 = new _module.C1();
              _nw199.__ctor(_module.__default.fm25(globalState), _module.__default.fm1(_1132_v1, globalState), (_this).f6);
              _1152_v18 = _nw199;
              _1146_v14 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), ((_1153_v16) => function (_1154_i4) {
                return (_1153_v16).f13;
              })(_1148_v16));
            } else {
              let _1155_v19;
              _1155_v19 = new _dafny.CodePoint('i'.codePointAt(0));
              let _1156_v20;
              let _nw200 = new _module.C1();
              _nw200.__ctor(_1155_v19, _1140_v8, ((_this).f6).plus(p2));
              _1156_v20 = _nw200;
              (globalState).f1 = (_this).f6;
              let _index197 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1129_v0).length));
              (_1129_v0)[_index197] = _dafny.Seq.contains(_1146_v14, _1155_v19);
              let _index198 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1134_v3).length));
              (_1134_v3)[_index198] = _this.f7;
              let _index199 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1129_v0).length));
              let _index200 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1134_v3).length));
              let _rhs250 = _1129_v0;
              let _rhs251 = _1129_v0;
              let _rhs252 = p0;
              let _rhs253 = new BigNumber((_1146_v14).length);
              let _rhs254 = new BigNumber(-972);
              let _lhs170 = _1129_v0;
              let _lhs171 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1129_v0).length));
              let _lhs172 = _1134_v3;
              let _lhs173 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1134_v3).length));
              _1129_v0 = _rhs250;
              _1129_v0 = _rhs251;
              _lhs170[_lhs171] = _rhs252;
              r0 = _rhs253;
              _lhs172[_lhs173] = _rhs254;
              let _1157_v21;
              _1157_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_1133_v2, _1133_v2), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.Concat(_1133_v2, _1133_v2)).length)), p0),((p1) ? (_1132_v1) : (_1132_v1)));
              _1157_v21 = (_1157_v21).update(_module.__default.fm21(globalState), !((_1129_v0)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_1129_v0).length))]));
              let _1158_v22;
              _1158_v22 = _dafny.Set.fromElements(p0);
              let _1159_v23;
              _1159_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1140_v8);
              let _1160_v25;
              _1160_v25 = _dafny.MultiSet.fromElements((_this).f6);
              let _1161_v28;
              let _nw201 = Array((new BigNumber(17)).toNumber());
              _nw201[(_dafny.ZERO).toNumber()] = (_module.__default.fm1(p0, globalState)).minus((_this).f6);
              _nw201[(_dafny.ONE).toNumber()] = ((_1134_v3)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1134_v3).length))]).minus(_1140_v8);
              _nw201[(new BigNumber(2)).toNumber()] = _module.__default.fm1(p1, globalState);
              _nw201[(new BigNumber(3)).toNumber()] = _1140_v8;
              _nw201[(new BigNumber(4)).toNumber()] = ((p0) ? (new BigNumber((_1158_v22).length)) : (new BigNumber((_1159_v23).length)));
              _nw201[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-436), p2);
              _nw201[(new BigNumber(6)).toNumber()] = _1140_v8;
              _nw201[(new BigNumber(7)).toNumber()] = ((_1134_v3)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1134_v3).length))]).plus(new BigNumber((_dafny.MultiSet.fromElements(_1140_v8)).cardinality()));
              _nw201[(new BigNumber(8)).toNumber()] = _module.__default.fm1(_1141_v9, globalState);
              _nw201[(new BigNumber(9)).toNumber()] = new BigNumber(((function () {
                let _coll35 = new _dafny.Set();
                for (const _compr_35 of _dafny.IntegerRange(new BigNumber(219), new BigNumber(988))) {
                  let _1162_v24 = _compr_35;
                  if (((new BigNumber(219)).isLessThanOrEqualTo(_1162_v24)) && ((_1162_v24).isLessThan(new BigNumber(988)))) {
                    _coll35.add((_1162_v24).multipliedBy(p2));
                  }
                }
                return _coll35;
              }()).Union(function () {
                let _coll36 = new _dafny.Set();
                for (const _compr_36 of (_1160_v25).Elements) {
                  let _1163_v26 = _compr_36;
                  if ((_1160_v25).contains(_1163_v26)) {
                    _coll36.add((_1163_v26).plus(new BigNumber(803)));
                  }
                }
                return _coll36;
              }())).length);
              _nw201[(new BigNumber(10)).toNumber()] = ((_1134_v3)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1134_v3).length))]).plus(new BigNumber((function () {
                let _coll37 = new _dafny.Map();
                for (const _compr_37 of _dafny.IntegerRange(new BigNumber(785), new BigNumber(192))) {
                  let _1164_v27 = _compr_37;
                  if (((new BigNumber(785)).isLessThanOrEqualTo(_1164_v27)) && ((_1164_v27).isLessThan(new BigNumber(192)))) {
                    _coll37.push([_module.__default.safeModuloInt(_1164_v27, _1140_v8),p1]);
                  }
                }
                return _coll37;
              }()).length));
              _nw201[(new BigNumber(11)).toNumber()] = p2;
              _nw201[(new BigNumber(12)).toNumber()] = p2;
              _nw201[(new BigNumber(13)).toNumber()] = p2;
              _nw201[(new BigNumber(14)).toNumber()] = (_this).fm15(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), ((_1165_v12) => function (_1166_i5) {
                return _1165_v12;
              })(_1144_v12)), _module.__default.safeIndex(new BigNumber(77), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), ((_1167_v12) => function (_1168_i5) {
                return _1167_v12;
              })(_1144_v12))).length)), _1144_v12)).length), _this.f7, globalState);
              _nw201[(new BigNumber(15)).toNumber()] = (new BigNumber(270)).multipliedBy(p2);
              _nw201[(new BigNumber(16)).toNumber()] = (_this).f6;
              _1161_v28 = _nw201;
              _1161_v28 = _1161_v28;
            }
          }
        }
      }
      let _1169_v29;
      _1169_v29 = _dafny.MultiSet.fromElements(p2);
      let _1170_v30;
      _1170_v30 = _dafny.Seq.of(_module.__default.fm35(_this.f7, new BigNumber((_1169_v29).cardinality()), globalState));
      let _1171_v31;
      _1171_v31 = _dafny.Seq.of(new BigNumber(((_1170_v30)[_module.__default.safeIndex(p2, new BigNumber((_1170_v30).length))]).length));
      let _1172_v32;
      _1172_v32 = _dafny.Seq.of((_1171_v31)[_module.__default.safeIndex(p2, new BigNumber((_1171_v31).length))], p2);
      let _1173_v33;
      _1173_v33 = _dafny.Seq.of(new BigNumber((_1172_v32).length));
      r1 = (_dafny.ZERO).minus((_1171_v31)[_module.__default.safeIndex(p2, new BigNumber((_1171_v31).length))]);
      let _1174_v34;
      _1174_v34 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1175_v35;
      let _nw202 = new _module.C1();
      _nw202.__ctor(_1174_v34, p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-485)), function (_1176_i6) {
        return (_module.D10.create_DC24(new _dafny.CodePoint('e'.codePointAt(0)))).dtor_cf51;
      })).length));
      _1175_v35 = _nw202;
      let _rhs255 = _1174_v34;
      let _rhs256 = _1175_v35;
      _1174_v34 = _rhs255;
      _1175_v35 = _rhs256;
      let _1177_v36;
      _1177_v36 = _dafny.Map.Empty.slice().updateUnsafe((p2).plus(p2),_1175_v35);
      _1177_v36 = (_1177_v36).update(p2, _1175_v35);
      let _1178_v37;
      _1178_v37 = true;
      let _1179_v38;
      _1179_v38 = _module.D10.create_DC26(_1178_v37);
      _1178_v37 = (_1179_v38).dtor_cf52;
      let _1180_v39;
      _1180_v39 = _dafny.Seq.UnicodeFromString("vdmx");
      r0 = (new BigNumber((_1180_v39).length)).multipliedBy(p2);
      r1 = ((_this).f6).plus(_module.__default.safeDivisionInt(p2, p2));
      r2 = p2;
      return [r0, r1, r2];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f7 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this.f12 = _dafny.Map.Empty;
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f11, f12, f6, f7) {
      let _this = this;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm11(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC0((_dafny.Set.fromElements((_this).f11)).Intersect(_dafny.Set.fromElements((_this).f11, false)));
    };
    fm12(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-535)), function (_1181_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(511)), function (_1182_i1) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }), _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _hi12 = _module.__default.safeDivisionInt(_module.__default.fm1(p0, globalState), p1);
      for (let _1183_i0 = (_this).f6; _1183_i0.isLessThan(_hi12); _1183_i0 = _1183_i0.plus(_dafny.ONE)) {
        let _1184_v0;
        _1184_v0 = _dafny.MultiSet.fromElements(_1183_i0);
        let _1185_v1;
        _1185_v1 = _dafny.Seq.of(new BigNumber(77));
        _1184_v0 = _dafny.MultiSet.FromArray(_1185_v1);
        let _1186_v2;
        _1186_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1183_i0,_module.__default.safeDivisionInt(p2, new BigNumber(76)));
        (_this).f7 = (((_1186_v2).contains((_dafny.ZERO).minus(_module.__default.safeModuloInt((_1185_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(_1183_i0), new BigNumber((_1185_v1).length))], (_this).f6)))) ? ((_1186_v2).get((_dafny.ZERO).minus(_module.__default.safeModuloInt((_1185_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(_1183_i0), new BigNumber((_1185_v1).length))], (_this).f6)))) : (_this.f7));
        let _1187_v3;
        _1187_v3 = _dafny.Seq.UnicodeFromString("t");
        _1187_v3 = _1187_v3;
        (_this).f7 = (p1).minus(new BigNumber(228));
      }
      let _1188_v4;
      _1188_v4 = _dafny.Seq.UnicodeFromString("uwyinhv");
      let _1189_v5;
      let _init37 = function (_1190_i1) {
        return (_1190_i1).multipliedBy((_this).f6);
      };
      let _nw203 = Array((new BigNumber(27)).toNumber());
      for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw203.length); _i0_37++) {
        _nw203[_i0_37] = _init37(new BigNumber(_i0_37));
      }
      _1189_v5 = _nw203;
      let _1191_v6;
      _1191_v6 = _dafny.Seq.of(_1189_v5, _1189_v5);
      let _1192_v7;
      let _nw204 = Array((new BigNumber(9)).toNumber());
      _nw204[(_dafny.ZERO).toNumber()] = p2;
      _nw204[(_dafny.ONE).toNumber()] = p2;
      _nw204[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
      _nw204[(new BigNumber(3)).toNumber()] = new BigNumber((_1188_v4).length);
      _nw204[(new BigNumber(4)).toNumber()] = p1;
      _nw204[(new BigNumber(5)).toNumber()] = p1;
      _nw204[(new BigNumber(6)).toNumber()] = _this.f7;
      _nw204[(new BigNumber(7)).toNumber()] = new BigNumber((_1191_v6).length);
      _nw204[(new BigNumber(8)).toNumber()] = _this.f7;
      _1192_v7 = _nw204;
      let _1193_v8;
      _1193_v8 = _module.D2.create_DC8(_module.D2.create_DC5(_1189_v5));
      _1193_v8 = _1193_v8;
      let _1194_v9;
      let _nw205 = new _module.C0();
      _nw205.__ctor(false);
      _1194_v9 = _nw205;
      let _1195_v10;
      _1195_v10 = _dafny.MultiSet.fromElements(p0, _1194_v9.f10);
      let _1196_v11;
      _1196_v11 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),_1195_v10);
      let _1197_v12;
      _1197_v12 = new _dafny.CodePoint('n'.codePointAt(0));
      let _1198_v13;
      _1198_v13 = ((((_1196_v11).contains(_1197_v12)) ? ((_1196_v11).get(_1197_v12)) : (_dafny.MultiSet.fromElements(_1194_v9.f10)))).Union(_1195_v10);
      let _source14 = _1198_v13;
      let _1199___mcc_h0 = _source14;
      let _1200_cf18 = _1199___mcc_h0;
      let _1201_v14;
      let _nw206 = Array((_dafny.ONE).toNumber()).fill(false);
      _1201_v14 = _nw206;
      let _index201 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_1201_v14).length));
      (_1201_v14)[_index201] = _1194_v9.f10;
      let _index202 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_1201_v14).length));
      (_1201_v14)[_index202] = true;
      let _index203 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_1201_v14).length));
      (_1201_v14)[_index203] = (_this).f11;
      let _1202_v15;
      _1202_v15 = _dafny.MultiSet.fromElements(_module.__default.fm1(p0, globalState));
      let _1203_v16;
      _1203_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1202_v15,p0);
      let _1204_v17;
      _1204_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_1203_v16).contains((_1202_v15).update(new BigNumber(520), _module.__default.abs((_this).f6)))) ? ((_1203_v16).get((_1202_v15).update(new BigNumber(520), _module.__default.abs((_this).f6)))) : (_1194_v9.f10)),new BigNumber(266));
      _1204_v17 = (_1204_v17).Merge((_1204_v17).Merge(_1204_v17));
      let _1205_v18;
      _1205_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1194_v9.f10);
      let _1206_v19;
      _1206_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("lbfloex"),_1197_v12);
      _1205_v18 = (_1205_v18).update(new BigNumber((_1206_v19).length), _1194_v9.f10);
      let _1207_v20;
      _1207_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),!((_this).f11));
      let _1208_v21;
      _1208_v21 = _module.D1.create_DC4(p0);
      let _1209_v22;
      _1209_v22 = _module.D2.create_DC7(p0, new BigNumber(79), _1207_v20, _1208_v21);
      let _source15 = _1209_v22;
      if (_source15.is_DC6) {
        let _1210___mcc_h1 = (_source15).cf10;
        let _1211___mcc_h2 = (_source15).cf11;
        let _1212___mcc_h3 = (_source15).cf12;
        let _1213_cf12 = _1212___mcc_h3;
        let _1214_cf11 = _1211___mcc_h2;
        let _1215_cf10 = _1210___mcc_h1;
        let _1216_v23;
        _1216_v23 = _dafny.Seq.of(_1194_v9.f10, _1194_v9.f10);
        r0 = !(((_module.__default.fm13(globalState)).update(_1194_v9.f10, _module.__default.abs(_1213_cf12))).IsSubsetOf((_module.__default.fm13(globalState)).update(p0, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_1216_v23).length))))));
        r1 = (_this).f11;
        let _1217_v24;
        let _nw207 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1217_v24 = _nw207;
        _1217_v24 = _1217_v24;
        let _1218_v25;
        let _nw208 = new _module.C0();
        _nw208.__ctor(p0);
        _1218_v25 = _nw208;
      } else if (_source15.is_DC7) {
        let _1219___mcc_h4 = (_source15).cf13;
        let _1220___mcc_h5 = (_source15).cf14;
        let _1221___mcc_h6 = (_source15).cf15;
        let _1222___mcc_h7 = (_source15).cf16;
        let _1223_cf16 = _1222___mcc_h7;
        let _1224_cf15 = _1221___mcc_h6;
        let _1225_cf14 = _1220___mcc_h5;
        let _1226_cf13 = _1219___mcc_h4;
        let _1227_v26;
        let _init38 = function (_1228_i2) {
          return (false) && (false);
        };
        let _nw209 = Array((new BigNumber(16)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw209.length); _i0_38++) {
          _nw209[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1227_v26 = _nw209;
        let _index204 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1227_v26).length));
        (_1227_v26)[_index204] = _1194_v9.f10;
        let _index205 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1227_v26).length));
        (_1227_v26)[_index205] = _1194_v9.f10;
        (_1194_v9).f10 = _1226_cf13;
        _1188_v4 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nvthhe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(504)), ((_1229_v12) => function (_1230_i3) {
          return _1229_v12;
        })(_1197_v12)));
        let _1231_v27;
        _1231_v27 = _module.D0.create_DC0(_dafny.Set.fromElements((_this).f11, (_this).f11));
        let _1232_v28;
        _1232_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1231_v27,_1225_cf14);
        _1232_v28 = (_1232_v28).update((_this).fm11(_1197_v12, globalState), p1);
      } else if (_source15.is_DC5) {
        let _1233___mcc_h8 = (_source15).cf9;
        let _1234_cf9 = _1233___mcc_h8;
        (globalState).f1 = _module.__default.fm1(_module.__default.fm0(_1194_v9.f10, _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p2), (_dafny.ZERO).minus(p1), globalState), globalState);
        if (_1194_v9.f10) {
          let _1235_v29;
          _1235_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(639),p2);
          _1235_v29 = (_1235_v29).update(((_this).f6).minus(new BigNumber(750)), new BigNumber((_1188_v4).length));
          let _1236_v30;
          _1236_v30 = _module.D2.create_DC5(_1189_v5);
          let _1237_v31;
          _1237_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v20,new BigNumber(243));
          let _rhs257 = ((p0) ? ((_dafny.ZERO).minus(p2)) : (new BigNumber((_1237_v31).length)));
          let _rhs258 = _1236_v30;
          let _lhs174 = globalState;
          _lhs174.f1 = _rhs257;
          _1236_v30 = _rhs258;
          (_1194_v9).f10 = true;
          let _1238_v32;
          let _nw210 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1238_v32 = _nw210;
          let _1239_v33;
          _1239_v33 = _module.D0.create_DC1(false, _1238_v32, (_this).f11);
          let _1240_v34;
          let _1241_v35;
          let _1242_v36;
          let _out23;
          let _out24;
          let _out25;
          let _outcollector7 = _module.__default.m0(_1239_v33, (_this).f6, globalState);
          _out23 = _outcollector7[0];
          _out24 = _outcollector7[1];
          _out25 = _outcollector7[2];
          _1240_v34 = _out23;
          _1241_v35 = _out24;
          _1242_v36 = _out25;
          _1235_v29 = (_1235_v29).update(_this.f7, (_this).f6);
        } else {
          (globalState).f1 = (_this).f6;
          let _1243_v37;
          let _nw211 = Array((new BigNumber(8)).toNumber()).fill([]);
          _1243_v37 = _nw211;
          _1243_v37 = _1243_v37;
          (_1194_v9).f10 = p0;
          _1207_v20 = (_1207_v20).update(_module.__default.safeDivisionInt(new BigNumber(-709), _this.f7), (_this).f11);
          let _1244_v38;
          _1244_v38 = _dafny.MultiSet.fromElements(p2);
          let _1245_v39;
          _1245_v39 = _dafny.Seq.of(_1244_v38);
          (globalState).f1 = new BigNumber((_1245_v39).length);
        }
        _1189_v5 = _1234_cf9;
        let _1246_v40;
        _1246_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1194_v9.f10,(_this).f11);
        let _index206 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1234_cf9).length));
        (_1234_cf9)[_index206] = new BigNumber(((_1246_v40).Merge(_1246_v40)).length);
        let _index207 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1234_cf9).length));
        (_1234_cf9)[_index207] = ((false) ? ((_this).f6) : (p1));
      } else {
        let _1247___mcc_h9 = (_source15).cf17;
        let _1248_cf17 = _1247___mcc_h9;
        r1 = _1194_v9.f10;
        let _index208 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1189_v5).length));
        (_1189_v5)[_index208] = p2;
        let _1249_v41;
        let _init39 = ((_1250_v9) => function (_1251_i4) {
          return _1250_v9.f10;
        })(_1194_v9);
        let _nw212 = Array((new BigNumber(7)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw212.length); _i0_39++) {
          _nw212[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1249_v41 = _nw212;
        let _index209 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1249_v41).length));
        (_1249_v41)[_index209] = !(p0);
        let _index210 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_1192_v7).length));
        (_1192_v7)[_index210] = _this.f7;
        let _index211 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1189_v5).length));
        let _index212 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1249_v41).length));
        let _index213 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_1192_v7).length));
        let _rhs259 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("mdoyxqucy"), _1188_v4);
        let _rhs260 = p2;
        let _rhs261 = p0;
        let _rhs262 = new BigNumber(-836);
        let _rhs263 = (_this).f11;
        let _lhs175 = _1194_v9;
        let _lhs176 = _1189_v5;
        let _lhs177 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1189_v5).length));
        let _lhs178 = _1249_v41;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1249_v41).length));
        let _lhs180 = _1192_v7;
        let _lhs181 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_1192_v7).length));
        _lhs175.f10 = _rhs259;
        _lhs176[_lhs177] = _rhs260;
        _lhs178[_lhs179] = _rhs261;
        _lhs180[_lhs181] = _rhs262;
        r1 = _rhs263;
        let _1252_v42;
        let _nw213 = Array((new BigNumber(13)).toNumber()).fill([]);
        _1252_v42 = _nw213;
        _1252_v42 = _1252_v42;
        let _1253_v43;
        _1253_v43 = _dafny.Seq.of(_1249_v41);
        let _1254_v44;
        _1254_v44 = _dafny.MultiSet.fromElements((_1253_v43)[_module.__default.safeIndex(_this.f7, new BigNumber((_1253_v43).length))]);
        _1254_v44 = (_1254_v44).update(_1249_v41, _module.__default.abs(p1));
      }
      let _1255_i5;
      _1255_i5 = _dafny.ZERO;
      L6: {
        while ((_module.__default.safeDivisionInt(new BigNumber(536), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1194_v9.f10,(_this).f11)).length))).isEqualTo(_this.f7)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1255_i5)) {
              break L6;
            }
            _1255_i5 = (_1255_i5).plus(_dafny.ONE);
            _1197_v12 = _module.__default.fm14(new BigNumber(194), new BigNumber((_1188_v4).length), new BigNumber(767), globalState);
            (_1194_v9).f10 = p0;
            let _1256_v45;
            _1256_v45 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_1188_v4, _1188_v4), _dafny.Seq.Concat(_1188_v4, _1188_v4), _dafny.Seq.Concat(_1188_v4, _dafny.Seq.UnicodeFromString("uohlhbu")));
            (globalState).f1 = new BigNumber((_1256_v45).cardinality());
            _1197_v12 = _1197_v12;
          }
        }
      }
      let _1257_v46;
      _1257_v46 = _dafny.Seq.of(p1);
      r0 = _dafny.Seq.contains(_1257_v46, p2);
      r1 = _dafny.areEqual(_1188_v4, _1188_v4);
      r2 = _this.f7;
      return [r0, r1, r2];
    }
    m4(globalState) {
      let _this = this;
      let _1258_v0;
      _1258_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(false) === ((_this).f11));
      let _1259_v1;
      _1259_v1 = _dafny.MultiSet.fromElements((_this).f11);
      _1258_v0 = (_1258_v0).update((_1259_v1).IsDisjointFrom(_1259_v1), true);
      let _1260_v2;
      _1260_v2 = true;
      _1260_v2 = (_this).f11;
      let _hi13 = new BigNumber(110);
      for (let _1261_i0 = _this.f7; _1261_i0.isLessThan(_hi13); _1261_i0 = _1261_i0.plus(_dafny.ONE)) {
        if (_1260_v2) {
          let _1262_v3;
          _1262_v3 = _1259_v1;
          let _1263_v4;
          _1263_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1262_v3);
          _1263_v4 = (_1263_v4).update(!((_this).f11) || ((_this).f11), _1262_v3);
          let _1264_v5;
          _1264_v5 = _dafny.Set.fromElements(_1260_v2, (_this).f11, _1260_v2);
          let _1265_v6;
          _1265_v6 = _module.D8.create_DC20((_this).f11, (_this).f11, _1260_v2, _1260_v2);
          let _1266_v7;
          _1266_v7 = _dafny.Set.fromElements(_module.__default.fm1((_1265_v6).dtor_cf39, globalState));
          let _1267_v8;
          let _nw214 = new _module.C2();
          _nw214.__ctor(new BigNumber((_1264_v5).length), new BigNumber((_1266_v7).length));
          _1267_v8 = _nw214;
          let _1268_v9;
          _1268_v9 = _dafny.Seq.of(_1260_v2, (_this).f11, false, _1260_v2, (_this).f11);
          let _1269_v10;
          _1269_v10 = _module.D8.create_DC18(_1267_v8);
          let _1270_v11;
          let _nw215 = Array((new BigNumber(29)).toNumber());
          _nw215[(_dafny.ZERO).toNumber()] = _1267_v8;
          _nw215[(_dafny.ONE).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(2)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(3)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(4)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(5)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(6)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(7)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(8)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(9)).toNumber()] = (((_this).f11) ? (_1267_v8) : (_1267_v8));
          _nw215[(new BigNumber(10)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(11)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(12)).toNumber()] = (((_this).f11) ? (_1267_v8) : (_1267_v8));
          _nw215[(new BigNumber(13)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(14)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(15)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(16)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(17)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(18)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(19)).toNumber()] = ((!((_1268_v9)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1260_v2, _1260_v2, _1260_v2)).length), new BigNumber((_1268_v9).length))])) ? ((_1269_v10).dtor_cf33) : (_1267_v8));
          _nw215[(new BigNumber(20)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(21)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(22)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(23)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(24)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(25)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(26)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(27)).toNumber()] = _1267_v8;
          _nw215[(new BigNumber(28)).toNumber()] = _1267_v8;
          _1270_v11 = _nw215;
          let _index214 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_1270_v11).length));
          (_1270_v11)[_index214] = _1267_v8;
          let _index215 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_1270_v11).length));
          (_1270_v11)[_index215] = (_module.D8.create_DC18(_1267_v8)).dtor_cf33;
          let _1271_v12;
          _1271_v12 = _dafny.Seq.of(_1267_v8.f7, new BigNumber(598), (_1261_i0).plus(_1261_i0));
          (_this).f7 = (_dafny.ZERO).minus((_1271_v12)[_module.__default.safeIndex(_module.__default.fm1(_1260_v2, globalState), new BigNumber((_1271_v12).length))]);
          let _1272_v13;
          let _init40 = function (_1273_i1) {
            return (_1273_i1).minus((_this).f6);
          };
          let _nw216 = Array((new BigNumber(24)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw216.length); _i0_40++) {
            _nw216[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1272_v13 = _nw216;
          let _index216 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1272_v13).length));
          (_1272_v13)[_index216] = ((_1260_v2) ? (new BigNumber(276)) : (_1267_v8.f7));
          let _index217 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1272_v13).length));
          (_1272_v13)[_index217] = (_1267_v8).f6;
          let _1274_v14;
          _1274_v14 = _dafny.Map.Empty.slice().updateUnsafe((_1270_v11)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_1270_v11).length))],_1260_v2);
          _1274_v14 = (_1274_v14).update((_1270_v11)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_1270_v11).length))], _1260_v2);
        } else {
          let _1275_v15;
          _1275_v15 = _dafny.Seq.of((_this).f11);
          let _1276_v16;
          _1276_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_this).f11, globalState),new BigNumber((_1275_v15).length));
          let _1277_v17;
          _1277_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1276_v16,(((_this).f11) ? (_1261_i0) : (_1261_i0)));
          _1277_v17 = (_1277_v17).update(_1276_v16, (_this.f7).plus(_this.f7));
          let _1278_v18;
          _1278_v18 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1279_v19;
          let _init41 = ((_1280_v2) => function (_1281_i2) {
            return _1280_v2;
          })(_1260_v2);
          let _nw217 = Array((new BigNumber(29)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw217.length); _i0_41++) {
            _nw217[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1279_v19 = _nw217;
          let _index218 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1279_v19).length));
          (_1279_v19)[_index218] = _1260_v2;
          let _1282_v20;
          let _init42 = function (_1283_i3) {
            return (_1283_i3).multipliedBy(new BigNumber(-171));
          };
          let _nw218 = Array((new BigNumber(24)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw218.length); _i0_42++) {
            _nw218[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1282_v20 = _nw218;
          let _index219 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          (_1282_v20)[_index219] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f6), (_this).f6);
          let _1284_v21;
          _1284_v21 = _dafny.Seq.UnicodeFromString("mxkhh");
          let _1285_v22;
          _1285_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1260_v2,_1284_v21);
          let _index220 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1279_v19).length));
          let _index221 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          let _rhs264 = _1278_v18;
          let _rhs265 = (!(_1260_v2)) || ((_1259_v1).contains(_1260_v2));
          let _rhs266 = _this.f7;
          let _rhs267 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-360)), _1261_i0));
          let _rhs268 = ((_module.__default.fm0((_this).f11, function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of _dafny.IntegerRange(new BigNumber(63), new BigNumber(280))) {
              let _1286_v23 = _compr_38;
              if (((new BigNumber(63)).isLessThanOrEqualTo(_1286_v23)) && ((_1286_v23).isLessThan(new BigNumber(280)))) {
                _coll38.push([(_1286_v23).minus(_this.f7),new BigNumber(689)]);
              }
            }
            return _coll38;
          }(), _this.f7, globalState)) ? ((_1285_v22).Merge(_1285_v22)) : (_dafny.Map.Empty.slice().updateUnsafe(!((_this).f11),(_this).fm12(false, (_this).f11, globalState))));
          let _lhs182 = _1279_v19;
          let _lhs183 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1279_v19).length));
          let _lhs184 = _1282_v20;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          let _lhs186 = globalState;
          _1278_v18 = _rhs264;
          _lhs182[_lhs183] = _rhs265;
          _lhs184[_lhs185] = _rhs266;
          _lhs186.f1 = _rhs267;
          _1285_v22 = _rhs268;
          _1284_v21 = _dafny.Seq.Concat(_1284_v21, _1284_v21);
          let _1287_v24;
          _1287_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1282_v20,_1261_i0);
          let _index222 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          let _index223 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          let _rhs269 = (new BigNumber((_module.__default.fm36(_this.f7, (_1279_v19)[_module.__default.safeIndex(new BigNumber(426), new BigNumber((_1279_v19).length))], (_this).f6, (_this).f11, globalState)).length)).isLessThanOrEqualTo(new BigNumber((_1287_v24).length));
          let _rhs270 = _1282_v20;
          let _rhs271 = (_this).f6;
          let _rhs272 = (_this.f7).plus(_this.f7);
          let _lhs187 = _1282_v20;
          let _lhs188 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          let _lhs189 = _1282_v20;
          let _lhs190 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1282_v20).length));
          _1260_v2 = _rhs269;
          _1282_v20 = _rhs270;
          _lhs187[_lhs188] = _rhs271;
          _lhs189[_lhs190] = _rhs272;
          let _1288_v25;
          _1288_v25 = _dafny.Set.fromElements(_1260_v2);
          _1260_v2 = !((_1288_v25).contains(true));
        }
        let _1289_v26;
        _1289_v26 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1290_v27;
        _1290_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_1260_v2) ? (new _dafny.CodePoint('i'.codePointAt(0))) : (_1289_v26)),(_this).f6);
        let _1291_v28;
        _1291_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f6);
        let _1292_v29;
        _1292_v29 = _dafny.Seq.of(_this.f7);
        let _1293_v30;
        _1293_v30 = _dafny.MultiSet.fromElements(_this.f7, (_this).f6, _1261_i0, new BigNumber((_1259_v1).cardinality()));
        let _1294_v31;
        _1294_v31 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_this.f7)).length), new BigNumber(774), new BigNumber((_1293_v30).cardinality()));
        let _1295_v32;
        let _nw219 = Array((new BigNumber(23)).toNumber());
        _nw219[(_dafny.ZERO).toNumber()] = _this.f7;
        _nw219[(_dafny.ONE).toNumber()] = (_this).f6;
        _nw219[(new BigNumber(2)).toNumber()] = (((_1291_v28).contains(_1260_v2)) ? ((_1291_v28).get(_1260_v2)) : (_this.f7));
        _nw219[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw219[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw219[(new BigNumber(5)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(6)).toNumber()] = new BigNumber(802);
        _nw219[(new BigNumber(7)).toNumber()] = new BigNumber((_1292_v29).length);
        _nw219[(new BigNumber(8)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(9)).toNumber()] = _this.f7;
        _nw219[(new BigNumber(10)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(11)).toNumber()] = (_this).f6;
        _nw219[(new BigNumber(12)).toNumber()] = _this.f7;
        _nw219[(new BigNumber(13)).toNumber()] = (_this).f6;
        _nw219[(new BigNumber(14)).toNumber()] = (_this).f6;
        _nw219[(new BigNumber(15)).toNumber()] = _this.f7;
        _nw219[(new BigNumber(16)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(17)).toNumber()] = _this.f7;
        _nw219[(new BigNumber(18)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(19)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(20)).toNumber()] = new BigNumber((_1294_v31).length);
        _nw219[(new BigNumber(21)).toNumber()] = _1261_i0;
        _nw219[(new BigNumber(22)).toNumber()] = (_this).f6;
        _1295_v32 = _nw219;
        let _1296_v33;
        _1296_v33 = _dafny.MultiSet.fromElements(_1295_v32);
        let _1297_v34;
        _1297_v34 = _dafny.Seq.UnicodeFromString("jpevpx");
        let _1298_v35;
        let _nw220 = Array((new BigNumber(9)).toNumber()).fill(false);
        _1298_v35 = _nw220;
        let _1299_v36;
        _1299_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1298_v35,_1261_i0);
        let _1300_v37;
        _1300_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(false, globalState),_this.f7);
        let _1301_v38;
        let _init43 = function (_1302_i5) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        };
        let _nw221 = Array((new BigNumber(16)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw221.length); _i0_43++) {
          _nw221[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1301_v38 = _nw221;
        let _1303_v39;
        _1303_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1301_v38);
        let _rhs273 = _module.__default.fm37((new BigNumber((_dafny.Set.fromElements(_1260_v2)).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), function (_1304_i4) {
          return _this.f7;
        })).length)), _module.D1.create_DC3(!(true), _1261_i0, new BigNumber((_1297_v34).length)), globalState);
        let _rhs274 = (_dafny.ZERO).minus(_this.f7);
        let _rhs275 = (_1299_v36).contains(_1298_v35);
        let _rhs276 = _module.__default.fm0(!(!(_1260_v2)), (_1300_v37).Merge(_1300_v37), new BigNumber((_1303_v39).length), globalState);
        let _rhs277 = _1296_v33;
        let _lhs191 = _this;
        _1290_v27 = _rhs273;
        _lhs191.f7 = _rhs274;
        _1260_v2 = _rhs275;
        _1260_v2 = _rhs276;
        _1296_v33 = _rhs277;
        _1260_v2 = !((_this).f11);
        let _1305_v40;
        let _init44 = function (_1306_i6) {
          return _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.Set.fromElements((_this).f11)),(_this).f11));
        };
        let _nw222 = Array((new BigNumber(15)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw222.length); _i0_44++) {
          _nw222[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1305_v40 = _nw222;
        let _1307_v41;
        _1307_v41 = _module.D6.create_DC13(_1305_v40);
        let _source16 = _1307_v41;
        if (_source16.is_DC14) {
          let _1308___mcc_h0 = (_source16).cf23;
          let _1309___mcc_h1 = (_source16).cf24;
          let _1310___mcc_h2 = (_source16).cf25;
          let _1311_cf25 = _1310___mcc_h2;
          let _1312_cf24 = _1309___mcc_h1;
          let _1313_cf23 = _1308___mcc_h0;
          _1313_cf23 = _module.__default.safeModuloInt(((_dafny.ZERO).minus((_this).f6)).plus(new BigNumber((_1294_v31).length)), (_1313_cf23).multipliedBy(_this.f7));
          let _1314_v42;
          _1314_v42 = _dafny.Seq.of(_1260_v2);
          let _1315_v43;
          _1315_v43 = _dafny.Seq.of((_1314_v42)[_module.__default.safeIndex(new BigNumber((_1314_v42).length), new BigNumber((_1314_v42).length))], true);
          let _index224 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1295_v32).length));
          (_1295_v32)[_index224] = (((_1291_v28).contains(_1260_v2)) ? ((_1291_v28).get(_1260_v2)) : (new BigNumber((_1315_v43).length)));
          let _1316_v44;
          _1316_v44 = _module.D2.create_DC5(_1295_v32);
          let _1317_v45;
          let _nw223 = Array((new BigNumber(2)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = _1316_v44;
          _nw223[(_dafny.ONE).toNumber()] = _1316_v44;
          _1317_v45 = _nw223;
          let _index225 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1298_v35).length));
          (_1298_v35)[_index225] = true;
          let _index226 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1298_v35).length));
          (_1298_v35)[_index226] = _1260_v2;
          let _index227 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1295_v32).length));
          let _index228 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1298_v35).length));
          let _index229 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1298_v35).length));
          let _rhs278 = _module.__default.fm1((_this).f11, globalState);
          let _rhs279 = _1317_v45;
          let _rhs280 = _1289_v26;
          let _rhs281 = (_this).f11;
          let _rhs282 = !(new BigNumber(455)).isEqualTo((((_this).f11) ? (new BigNumber((_dafny.Set.fromElements(_1260_v2, (_this).f11, _1260_v2)).length)) : (_1261_i0)));
          let _lhs192 = _1295_v32;
          let _lhs193 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1295_v32).length));
          let _lhs194 = _1298_v35;
          let _lhs195 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1298_v35).length));
          let _lhs196 = _1298_v35;
          let _lhs197 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1298_v35).length));
          _lhs192[_lhs193] = _rhs278;
          _1317_v45 = _rhs279;
          _1289_v26 = _rhs280;
          _lhs194[_lhs195] = _rhs281;
          _lhs196[_lhs197] = _rhs282;
          _1260_v2 = (_1298_v35)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_1298_v35).length))];
          let _1318_v46;
          _1318_v46 = _dafny.Set.fromElements(_1260_v2, _1260_v2, false);
          let _1319_v47;
          _1319_v47 = _module.D7.create_DC17(_dafny.Seq.UnicodeFromString("dufjuu"), _1318_v46);
          let _1320_v48;
          let _nw224 = new _module.C2();
          _nw224.__ctor(_1311_cf25, _1311_cf25);
          _1320_v48 = _nw224;
          let _1321_v49;
          _1321_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1320_v48,_1318_v46);
          let _pat_let_tv28 = _1321_v49;
          let _pat_let_tv29 = _1320_v48;
          let _pat_let_tv30 = _1318_v46;
          let _pat_let_tv31 = _1321_v49;
          let _pat_let_tv32 = _1320_v48;
          let _1322_v50;
          _1322_v50 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let25_0) {
            return function (_1323_dt__update__tmp_h0) {
              return function (_pat_let26_0) {
                return function (_1324_dt__update_hcf32_h0) {
                  return _module.D7.create_DC17((_1323_dt__update__tmp_h0).dtor_cf31, _1324_dt__update_hcf32_h0);
                }(_pat_let26_0);
              }((((_pat_let_tv31).contains(_pat_let_tv32)) ? ((_pat_let_tv28).get(_pat_let_tv29)) : (_pat_let_tv30)));
            }(_pat_let25_0);
          }(_1319_v47),new BigNumber(374));
          let _1325_v51;
          let _nw225 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _1325_v51 = _nw225;
          (_this).f7 = (((_1322_v50).contains(_module.D7.create_DC17(_dafny.Seq.UnicodeFromString("txexlpnd"), _1318_v46))) ? ((_1322_v50).get(_module.D7.create_DC17(_dafny.Seq.UnicodeFromString("txexlpnd"), _1318_v46))) : (new BigNumber((_dafny.Seq.of(_1325_v51)).length)));
        } else if (_source16.is_DC15) {
          let _1326___mcc_h3 = (_source16).cf26;
          let _1327___mcc_h4 = (_source16).cf27;
          let _1328___mcc_h5 = (_source16).cf28;
          let _1329___mcc_h6 = (_source16).cf29;
          let _1330_cf29 = _1329___mcc_h6;
          let _1331_cf28 = _1328___mcc_h5;
          let _1332_cf27 = _1327___mcc_h4;
          let _1333_cf26 = _1326___mcc_h3;
          _1333_cf26 = _1333_cf26;
          let _index230 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1295_v32).length));
          (_1295_v32)[_index230] = _1330_cf29;
          let _index231 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1295_v32).length));
          (_1295_v32)[_index231] = _1330_cf29;
          let _index232 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1295_v32).length));
          (_1295_v32)[_index232] = new BigNumber(-383);
          _1333_cf26 = (new BigNumber(-823)).isLessThanOrEqualTo(_1332_cf27);
        } else {
          let _1334___mcc_h7 = (_source16).cf22;
          let _1335_cf22 = _1334___mcc_h7;
          _1300_v37 = (_1300_v37).update((_this).f6, _this.f7);
          let _1336_v52;
          _1336_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1260_v2, _1260_v2),_1301_v38);
          let _1337_v53;
          _1337_v53 = _dafny.Seq.of(_1260_v2);
          _1336_v52 = (_1336_v52).update(_dafny.Seq.Concat(_1337_v53, _dafny.Seq.update(_1337_v53, _module.__default.safeIndex(_this.f7, new BigNumber((_1337_v53).length)), !((_this).f11))), _1301_v38);
          _1300_v37 = (_1300_v37).update((_dafny.ZERO).minus((_this).f6), _1261_i0);
          _1291_v28 = (_1291_v28).update(_1260_v2, (_this).f6);
        }
      }
      let _1338_i7;
      _1338_i7 = _dafny.ZERO;
      L7: {
        while ((_this).f11) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1338_i7)) {
              break L7;
            }
            _1338_i7 = (_1338_i7).plus(_dafny.ONE);
            (_this).f7 = _this.f7;
            let _1339_v54;
            _1339_v54 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1258_v0).length)));
            let _1340_v55;
            _1340_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_1339_v54).cardinality()));
            let _1341_v56;
            _1341_v56 = _dafny.Set.fromElements(_1340_v55);
            if ((_1341_v56).IsSubsetOf((_1341_v56).Difference(_1341_v56))) {
              let _1342_v57;
              _1342_v57 = _dafny.Set.fromElements(_1260_v2);
              let _1343_v58;
              _1343_v58 = _module.D0.create_DC0(_1342_v57);
              let _1344_v59;
              _1344_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1343_v58,(_this).f11);
              _1344_v59 = _1344_v59;
              let _1345_v60;
              let _init45 = function (_1346_i8) {
                return _module.__default.safeModuloInt(_1346_i8, _this.f7);
              };
              let _nw226 = Array((new BigNumber(12)).toNumber());
              for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw226.length); _i0_45++) {
                _nw226[_i0_45] = _init45(new BigNumber(_i0_45));
              }
              _1345_v60 = _nw226;
              let _index233 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1345_v60).length));
              (_1345_v60)[_index233] = (_dafny.ZERO).minus(_this.f7);
              let _index234 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1345_v60).length));
              (_1345_v60)[_index234] = (_this).f6;
              _1260_v2 = (_this).f11;
              _1260_v2 = (true) === (_1260_v2);
              let _1347_v61;
              let _nw227 = Array((new BigNumber(8)).toNumber()).fill(false);
              _1347_v61 = _nw227;
              let _index235 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1347_v61).length));
              (_1347_v61)[_index235] = !(_1260_v2);
              let _1348_v62;
              let _nw228 = new _module.C2();
              _nw228.__ctor(_this.f7, (_this).f6);
              _1348_v62 = _nw228;
              let _1349_v63;
              _1349_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1348_v62,(_this).f6);
              let _index236 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1345_v60).length));
              let _index237 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1347_v61).length));
              let _index238 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1345_v60).length));
              let _rhs283 = (_1259_v1).IsProperSubsetOf(_1259_v1);
              let _rhs284 = (((_1349_v63).contains(_1348_v62)) ? ((_1349_v63).get(_1348_v62)) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(((_1348_v62).f6).multipliedBy(_1348_v62.f7)))));
              let _rhs285 = (_this).f11;
              let _rhs286 = _this.f7;
              let _lhs198 = _1345_v60;
              let _lhs199 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1345_v60).length));
              let _lhs200 = _1347_v61;
              let _lhs201 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1347_v61).length));
              let _lhs202 = _1345_v60;
              let _lhs203 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1345_v60).length));
              _1260_v2 = _rhs283;
              _lhs198[_lhs199] = _rhs284;
              _lhs200[_lhs201] = _rhs285;
              _lhs202[_lhs203] = _rhs286;
            } else {
              _1260_v2 = (_this).f11;
              let _1350_v64;
              let _nw229 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
              _1350_v64 = _nw229;
              let _1351_v65;
              let _init46 = ((_1352_v2) => function (_1353_i9) {
                return _1352_v2;
              })(_1260_v2);
              let _nw230 = Array((new BigNumber(22)).toNumber());
              for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw230.length); _i0_46++) {
                _nw230[_i0_46] = _init46(new BigNumber(_i0_46));
              }
              _1351_v65 = _nw230;
              let _index239 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1350_v64).length));
              (_1350_v64)[_index239] = _dafny.Set.fromElements(_1351_v65);
              let _1354_v66;
              _1354_v66 = _dafny.Set.fromElements(_1351_v65, _1351_v65);
              let _1355_v67;
              _1355_v67 = _module.D11.create_DC27(_1354_v66);
              let _index240 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1350_v64).length));
              (_1350_v64)[_index240] = (_1355_v67).dtor_cf53;
              let _1356_v68;
              let _nw231 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
              _1356_v68 = _nw231;
              let _1357_v69;
              _1357_v69 = _dafny.Seq.of(_1258_v0);
              let _index241 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1356_v68).length));
              (_1356_v68)[_index241] = new BigNumber((_dafny.Seq.update(_1357_v69, _module.__default.safeIndex(_this.f7, new BigNumber((_1357_v69).length)), _1258_v0)).length);
              let _1358_v70;
              _1358_v70 = _dafny.Seq.UnicodeFromString("um");
              let _index242 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1356_v68).length));
              (_1356_v68)[_index242] = _module.__default.fm1(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ut"), _1358_v70), globalState);
              let _1359_v71;
              _1359_v71 = _dafny.Seq.of(_this.f7);
              _1359_v71 = _1359_v71;
              _1260_v2 = (_this).f11;
            }
            let _1360_v73;
            _1360_v73 = _dafny.Seq.of((_this).f6, new BigNumber(-596));
            let _1361_v74;
            _1361_v74 = _dafny.Seq.UnicodeFromString("utk");
            let _1362_v75;
            let _nw232 = new _module.C0();
            _nw232.__ctor(_module.__default.fm0(_module.__default.fm0(true, function () {
              let _coll39 = new _dafny.Map();
              for (const _compr_39 of _dafny.IntegerRange(new BigNumber(-881), new BigNumber(-796))) {
                let _1363_v72 = _compr_39;
                if (((new BigNumber(-881)).isLessThanOrEqualTo(_1363_v72)) && ((_1363_v72).isLessThan(new BigNumber(-796)))) {
                  _coll39.push([(_1363_v72).multipliedBy(new BigNumber((_1259_v1).cardinality())),_this.f7]);
                }
              }
              return _coll39;
            }(), (_1360_v73)[_module.__default.safeIndex(new BigNumber((_1361_v74).length), new BigNumber((_1360_v73).length))], globalState), _1340_v55, (_dafny.ZERO).minus((_this).f6), globalState));
            _1362_v75 = _nw232;
            (_this).f7 = (_this).f6;
          }
        }
      }
      let _1364_i10;
      _1364_i10 = _dafny.ZERO;
      L8: {
        while (false) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1364_i10)) {
              break L8;
            }
            _1364_i10 = (_1364_i10).plus(_dafny.ONE);
            let _1365_v76;
            let _nw233 = Array((new BigNumber(12)).toNumber()).fill(false);
            _1365_v76 = _nw233;
            let _1366_v77;
            _1366_v77 = _dafny.Seq.UnicodeFromString("kdmrd");
            let _1367_v78;
            _1367_v78 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("qpdetefa"));
            let _1368_v79;
            _1368_v79 = new _dafny.CodePoint('y'.codePointAt(0));
            let _1369_v80;
            _1369_v80 = _module.D12.create_DC30(_1367_v78);
            let _1370_v81;
            let _nw234 = Array((new BigNumber(24)).toNumber()).fill(_module.D1.Default());
            _1370_v81 = _nw234;
            let _1371_v82;
            _1371_v82 = _dafny.Seq.of(_1370_v81);
            let _1372_v83;
            _1372_v83 = _dafny.Seq.of(_this.f7, _this.f7, (_this).f6);
            let _1373_v84;
            _1373_v84 = _dafny.Seq.of(_1260_v2, (_this).f11);
            let _1374_v85;
            let _nw235 = Array((new BigNumber(23)).toNumber());
            _nw235[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_1366_v77);
            _nw235[(_dafny.ONE).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_1367_v78, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f11)).length), new BigNumber((_1367_v78).length)), _1366_v77);
            _nw235[(new BigNumber(3)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_dafny.Seq.update(_1366_v77, _module.__default.safeIndex((_this).f6, new BigNumber((_1366_v77).length)), _1368_v79), _1366_v77);
            _nw235[(new BigNumber(5)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1366_v77), _dafny.Seq.update(_1367_v78, _module.__default.safeIndex((_this).f6, new BigNumber((_1367_v78).length)), _1366_v77));
            _nw235[(new BigNumber(7)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(8)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_1366_v77, _dafny.Seq.UnicodeFromString("fdn"), _1366_v77, _1366_v77, _1366_v77);
            _nw235[(new BigNumber(10)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1367_v78, _1367_v78), _module.__default.safeIndex(_this.f7, new BigNumber((_dafny.Seq.Concat(_1367_v78, _1367_v78)).length)), _dafny.Seq.UnicodeFromString("smsluxg"));
            _nw235[(new BigNumber(12)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(13)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(196)), ((_1375_v77) => function (_1376_i11) {
              return _1375_v77;
            })(_1366_v77)), _1367_v78);
            _nw235[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), ((_1377_v77) => function (_1378_i12) {
              return _1377_v77;
            })(_1366_v77));
            _nw235[(new BigNumber(16)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(17)).toNumber()] = _dafny.Seq.update((_1369_v80).dtor_cf59, _module.__default.safeIndex(new BigNumber((_1371_v82).length), new BigNumber(((_1369_v80).dtor_cf59).length)), _module.__default.fm2(_1372_v83, _this.f7, _this.f7, (_this).f11, globalState));
            _nw235[(new BigNumber(18)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(19)).toNumber()] = _module.__default.fm38(globalState);
            _nw235[(new BigNumber(20)).toNumber()] = _1367_v78;
            _nw235[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_1367_v78, _module.__default.safeIndex(new BigNumber((_1373_v84).length), new BigNumber((_1367_v78).length)), _1366_v77);
            _nw235[(new BigNumber(22)).toNumber()] = _1367_v78;
            _1374_v85 = _nw235;
            let _index243 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_1374_v85).length));
            (_1374_v85)[_index243] = _1367_v78;
            let _1379_v86;
            _1379_v86 = _dafny.Seq.of(_dafny.Seq.of((_module.D10.create_DC26(_1260_v2)).dtor_cf52));
            let _index244 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_1374_v85).length));
            let _rhs287 = _1260_v2;
            let _rhs288 = ((_1260_v2) ? ((_1372_v83)[_module.__default.safeIndex((_this).f6, new BigNumber((_1372_v83).length))]) : (_module.__default.fm1(true, globalState)));
            let _rhs289 = (((_this.f7).isLessThan(_this.f7)) ? (_1365_v76) : (_1365_v76));
            let _rhs290 = !(((_this.f7).isLessThan(_this.f7)) && (_dafny.Seq.IsPrefixOf(_1379_v86, _dafny.Seq.of(_1373_v84))));
            let _rhs291 = _dafny.Seq.Concat(_1367_v78, _dafny.Seq.update(_1367_v78, _module.__default.safeIndex((_this).f6, new BigNumber((_1367_v78).length)), _1366_v77));
            let _lhs204 = globalState;
            let _lhs205 = _1374_v85;
            let _lhs206 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_1374_v85).length));
            _1260_v2 = _rhs287;
            _lhs204.f1 = _rhs288;
            _1365_v76 = _rhs289;
            _1260_v2 = _rhs290;
            _lhs205[_lhs206] = _rhs291;
            _1260_v2 = (_this.f7).isLessThanOrEqualTo((_this).f6);
            let _index245 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1365_v76).length));
            (_1365_v76)[_index245] = (_this).f11;
            let _index246 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1365_v76).length));
            (_1365_v76)[_index246] = (_module.__default.safeDivisionInt(_this.f7, new BigNumber((_1372_v83).length))).isLessThanOrEqualTo((_this).f6);
            (_this).f7 = _this.f7;
          }
        }
      }
      let _1380_v87;
      _1380_v87 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1381_v88;
      _1381_v88 = _module.D10.create_DC24(_1380_v87);
      _1380_v87 = (_1381_v88).dtor_cf51;
      return;
    }
    m5(p0, globalState) {
      let _this = this;
      let _1382_i0;
      _1382_i0 = _dafny.ZERO;
      L9: {
        while (((_this).f11) && ((_this).f11)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1382_i0)) {
              break L9;
            }
            _1382_i0 = (_1382_i0).plus(_dafny.ONE);
            let _1383_v0;
            let _nw236 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
            _1383_v0 = _nw236;
            let _1384_v1;
            _1384_v1 = new _dafny.CodePoint('j'.codePointAt(0));
            let _1385_v2;
            _1385_v2 = _dafny.Seq.of(_1384_v1);
            let _1386_v3;
            let _nw237 = Array((new BigNumber(4)).toNumber()).fill(false);
            _1386_v3 = _nw237;
            let _1387_v4;
            _1387_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1385_v2).length),_1386_v3);
            let _index247 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_1383_v0).length));
            (_1383_v0)[_index247] = _1387_v4;
            let _index248 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_1383_v0).length));
            (_1383_v0)[_index248] = _1387_v4;
            let _1388_v5;
            _1388_v5 = _dafny.Seq.of(_1385_v2);
            let _1389_v6;
            _1389_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f11);
            let _1390_v7;
            _1390_v7 = _dafny.Seq.of(_this.f7, p0, new BigNumber((_1389_v6).length));
            let _1391_v8;
            _1391_v8 = _dafny.Seq.of(new BigNumber((_1390_v7).length));
            let _1392_v9;
            _1392_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v5,_module.__default.safeModuloInt(new BigNumber((_1391_v8).length), _module.__default.fm1((_this).f11, globalState)));
            (_this).f7 = (((_1392_v9).contains(_1388_v5)) ? ((_1392_v9).get(_1388_v5)) : (_module.__default.fm1((_this).f11, globalState)));
            if (!(!(_this.f7).isEqualTo(_this.f7)) || ((_this).f11)) {
              let _1393_v10;
              _1393_v10 = true;
              let _1394_v11;
              _1394_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1385_v2).length),(_this).f6);
              let _1395_v12;
              _1395_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1384_v1,(((_1394_v11).contains(_this.f7)) ? ((_1394_v11).get(_this.f7)) : (_this.f7)));
              let _1396_v13;
              _1396_v13 = _module.D7.create_DC16(_1395_v12);
              let _pat_let_tv33 = _1395_v12;
              let _pat_let_tv34 = _1395_v12;
              let _rhs292 = _1385_v2;
              let _rhs293 = _1393_v10;
              let _rhs294 = _module.__default.fm1((_this).f11, globalState);
              let _rhs295 = function (_pat_let27_0) {
                return function (_1397_dt__update__tmp_h0) {
                  return function (_pat_let28_0) {
                    return function (_1398_dt__update_hcf30_h0) {
                      return _module.D7.create_DC16(_1398_dt__update_hcf30_h0);
                    }(_pat_let28_0);
                  }((_pat_let_tv33).Merge(_pat_let_tv34));
                }(_pat_let27_0);
              }(_1396_v13);
              let _lhs207 = globalState;
              _1385_v2 = _rhs292;
              _1393_v10 = _rhs293;
              _lhs207.f1 = _rhs294;
              _1396_v13 = _rhs295;
              (globalState).f1 = _this.f7;
              (globalState).f1 = ((true) ? (new BigNumber(-97)) : (_this.f7));
              let _1399_v14;
              _1399_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f11);
              _1399_v14 = _1399_v14;
              let _index249 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1386_v3).length));
              (_1386_v3)[_index249] = _1393_v10;
              let _1400_v15;
              _1400_v15 = _dafny.Set.fromElements((((_1389_v6).contains(true)) ? ((_1389_v6).get(true)) : (_1393_v10)), (_this).f11, _1393_v10);
              let _index250 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1386_v3).length));
              let _rhs296 = _1396_v13;
              let _rhs297 = (_1400_v15).contains(false);
              let _rhs298 = ((_this).f6).minus((_dafny.ZERO).minus(new BigNumber((_1385_v2).length)));
              let _lhs208 = _1386_v3;
              let _lhs209 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1386_v3).length));
              let _lhs210 = globalState;
              _1396_v13 = _rhs296;
              _lhs208[_lhs209] = _rhs297;
              _lhs210.f1 = _rhs298;
            } else {
              let _index251 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length));
              (_1386_v3)[_index251] = (_this).f11;
              let _1401_v16;
              _1401_v16 = _module.D8.create_DC20((_this).f11, (_this).f11, (_this).f11, (_this).f11);
              let _1402_v17;
              _1402_v17 = _dafny.Seq.of((function (_pat_let29_0) {
                return function (_1403_dt__update__tmp_h1) {
                  return function (_pat_let30_0) {
                    return function (_1404_dt__update_hcf39_h0) {
                      return _module.D8.create_DC20(_1404_dt__update_hcf39_h0, (_1403_dt__update__tmp_h1).dtor_cf40, (_1403_dt__update__tmp_h1).dtor_cf41, (_1403_dt__update__tmp_h1).dtor_cf42);
                    }(_pat_let30_0);
                  }((_this).f11);
                }(_pat_let29_0);
              }(_1401_v16)).dtor_cf39, (_this).f11);
              let _1405_v18;
              _1405_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1386_v3);
              let _1406_v19;
              _1406_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f11,_module.__default.fm35(new BigNumber((_1402_v17).length), new BigNumber((_1405_v18).length), globalState)),(_this).f11);
              let _1407_v20;
              _1407_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1390_v7);
              let _index252 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length));
              (_1386_v3)[_index252] = (((_1406_v19).contains(_1407_v20)) ? ((_1406_v19).get(_1407_v20)) : (false));
              let _1408_v21;
              _1408_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ctwrgg"),_this.f7);
              let _1409_v22;
              _1409_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1408_v21).length),(_1386_v3)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length))]);
              _1409_v22 = (_1409_v22).update(p0, (_1386_v3)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length))]);
              let _1410_v23;
              let _nw238 = Array((new BigNumber(19)).toNumber()).fill(_module.D6.Default());
              _1410_v23 = _nw238;
              let _index253 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1410_v23).length));
              (_1410_v23)[_index253] = _module.D6.create_DC15((_this).f11, (_this).f6, (_this).f6, p0);
              let _1411_v24;
              let _nw239 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
              _1411_v24 = _nw239;
              let _1412_v25;
              _1412_v25 = _module.D6.create_DC15((_1386_v3)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length))], _module.__default.fm1(!((_1386_v3)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length))]), globalState), _this.f7, p0);
              let _pat_let_tv35 = p0;
              let _index254 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1410_v23).length));
              let _rhs299 = function (_pat_let31_0) {
                return function (_1413_dt__update__tmp_h2) {
                  return function (_pat_let32_0) {
                    return function (_1414_dt__update_hcf29_h0) {
                      return function (_pat_let33_0) {
                        return function (_1415_dt__update_hcf27_h0) {
                          return _module.D6.create_DC15((_1413_dt__update__tmp_h2).dtor_cf26, _1415_dt__update_hcf27_h0, (_1413_dt__update__tmp_h2).dtor_cf28, _1414_dt__update_hcf29_h0);
                        }(_pat_let33_0);
                      }(_pat_let_tv35);
                    }(_pat_let32_0);
                  }((_this).f6);
                }(_pat_let31_0);
              }(_1412_v25);
              let _rhs300 = _1411_v24;
              let _lhs211 = _1410_v23;
              let _lhs212 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1410_v23).length));
              _lhs211[_lhs212] = _rhs299;
              _1411_v24 = _rhs300;
              let _index255 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1386_v3).length));
              (_1386_v3)[_index255] = (_this.f7).isEqualTo(p0);
              _1390_v7 = _dafny.Seq.Concat(_1391_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(84)), function (_1416_i1) {
                return new BigNumber(912);
              }));
            }
            let _1417_v26;
            _1417_v26 = _dafny.MultiSet.fromElements(false, (_this).f11);
            let _index256 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1386_v3).length));
            (_1386_v3)[_index256] = (_1417_v26).IsSubsetOf(_1417_v26);
            let _index257 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1386_v3).length));
            (_1386_v3)[_index257] = (_this).f11;
          }
        }
      }
      let _1418_v27;
      _1418_v27 = _dafny.Seq.UnicodeFromString("jnumj");
      let _1419_v28;
      _1419_v28 = new _dafny.CodePoint('r'.codePointAt(0));
      let _1420_v29;
      _1420_v29 = _dafny.Seq.of(_this.f7, _this.f7);
      let _1421_v30;
      _1421_v30 = _dafny.MultiSet.fromElements(_this.f7, _this.f7);
      let _1422_v31;
      let _nw240 = Array((new BigNumber(18)).toNumber());
      _nw240[(_dafny.ZERO).toNumber()] = _1418_v27;
      _nw240[(_dafny.ONE).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(2)).toNumber()] = (((_this).f11) ? (_dafny.Seq.UnicodeFromString("jeil")) : (_1418_v27));
      _nw240[(new BigNumber(3)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1418_v27, _module.__default.safeIndex(p0, new BigNumber((_1418_v27).length)), _1419_v28);
      _nw240[(new BigNumber(5)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(6)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(7)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(8)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(9)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(10)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(11)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm2(_1420_v29, (((_1421_v30).contains(p0)) ? ((_1421_v30).get(p0)) : ((_this).f6)), (_this).f6, (_this).f11, globalState), _1418_v27);
      _nw240[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1418_v27, _1418_v27);
      _nw240[(new BigNumber(14)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(15)).toNumber()] = _1418_v27;
      _nw240[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_1418_v27, _1418_v27);
      _nw240[(new BigNumber(17)).toNumber()] = _1418_v27;
      _1422_v31 = _nw240;
      let _1423_v32;
      _1423_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_1418_v27).length));
      let _index258 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1422_v31).length));
      (_1422_v31)[_index258] = (_this).fm12(_module.__default.fm0((_this).f11, _1423_v32, (_dafny.ZERO).minus(p0), globalState), (_this).f11, globalState);
      let _index259 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1422_v31).length));
      (_1422_v31)[_index259] = _1418_v27;
      let _1424_v33;
      _1424_v33 = _dafny.Set.fromElements(p0, (_1420_v29)[_module.__default.safeIndex(new BigNumber(-56), new BigNumber((_1420_v29).length))]);
      let _1425_v34;
      _1425_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_module.D12.create_DC31(new BigNumber(282)));
      let _1426_v35;
      _1426_v35 = _dafny.MultiSet.fromElements(_1425_v34, _1425_v34);
      let _1427_v36;
      _1427_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1424_v33).length),_1426_v35);
      _1427_v36 = (_1427_v36).update((((_this).f11) ? ((_this).f6) : ((_this).f6)), _1426_v35);
      let _1428_v37;
      _1428_v37 = true;
      let _1429_v38;
      _1429_v38 = _dafny.Seq.of(true);
      let _1430_v39;
      _1430_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1428_v37,new BigNumber(((_1422_v31)[_module.__default.safeIndex(new BigNumber(56), new BigNumber((_1422_v31).length))]).length));
      _1428_v37 = !((_dafny.Set.fromElements(_this.f7, (_this).f6, new BigNumber((_1429_v38).length), (_this).f6, new BigNumber((_1430_v39).length))).Difference(_1424_v33)).equals(_dafny.Set.fromElements(p0, p0, (_dafny.ZERO).minus(new BigNumber(-177)), new BigNumber(-522)));
      let _1431_v40;
      _1431_v40 = _module.D10.create_DC24(_1419_v28);
      let _source17 = _1431_v40;
      if (_source17.is_DC25) {
        let _1432_v41;
        _1432_v41 = _dafny.Set.fromElements(_dafny.areEqual(_1429_v38, _1429_v38), _1428_v37);
        _1432_v41 = _1432_v41;
        let _1433_v42;
        let _nw241 = new _module.C3();
        _nw241.__ctor(new BigNumber((_1418_v27).length), new BigNumber(-326));
        _1433_v42 = _nw241;
        let _1434_v43;
        let _nw242 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _1434_v43 = _nw242;
        let _index260 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length));
        (_1434_v43)[_index260] = (_dafny.ZERO).minus(p0);
        let _index261 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length));
        (_1434_v43)[_index261] = new BigNumber((_1420_v29).length);
        if (((p0).minus(new BigNumber(-984))).isLessThanOrEqualTo(p0)) {
          let _1435_v44;
          let _nw243 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1435_v44 = _nw243;
          let _index262 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length));
          (_1435_v44)[_index262] = _module.__default.fm0(_module.__default.fm0(true, _1423_v32, (_this).f6, globalState), _1423_v32, new BigNumber((_1432_v41).length), globalState);
          let _index263 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length));
          (_1435_v44)[_index263] = (_this).f11;
          let _1436_v45;
          let _nw244 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1436_v45 = _nw244;
          let _index264 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1436_v45).length));
          (_1436_v45)[_index264] = _1419_v28;
          let _index265 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length));
          let _index266 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1436_v45).length));
          let _rhs301 = (_1435_v44)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length))];
          let _rhs302 = p0;
          let _rhs303 = (_1435_v44)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length))];
          let _rhs304 = _1419_v28;
          let _rhs305 = (_this).f11;
          let _lhs213 = _1435_v44;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length));
          let _lhs215 = globalState;
          let _lhs216 = _1436_v45;
          let _lhs217 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1436_v45).length));
          _lhs213[_lhs214] = _rhs301;
          _lhs215.f1 = _rhs302;
          _1428_v37 = _rhs303;
          _lhs216[_lhs217] = _rhs304;
          _1428_v37 = _rhs305;
          let _1437_v46;
          _1437_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
          let _1438_v47;
          _1438_v47 = _module.D0.create_DC0(_1432_v41);
          let _1439_v48;
          _1439_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1438_v47,!((_1435_v44)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length))]));
          let _1440_v49;
          _1440_v49 = _module.D1.create_DC2(_1439_v48);
          let _pat_let_tv36 = _1439_v48;
          let _1441_v50;
          _1441_v50 = _dafny.Set.fromElements(_1440_v49, function (_pat_let34_0) {
            return function (_1442_dt__update__tmp_h3) {
              return function (_pat_let35_0) {
                return function (_1443_dt__update_hcf4_h0) {
                  return _module.D1.create_DC2(_1443_dt__update_hcf4_h0);
                }(_pat_let35_0);
              }(_pat_let_tv36);
            }(_pat_let34_0);
          }(_module.D1.create_DC2(_1439_v48)), _1440_v49, _1440_v49);
          let _1444_v51;
          let _out26;
          _out26 = (_1433_v42).m6((_1437_v46).Merge(_1437_v46), _this.f7, _1441_v50, globalState);
          _1444_v51 = _out26;
          let _1445_v52;
          let _init47 = ((_1446_v51, _1447_v44) => function (_1448_i2) {
            return _module.D6.create_DC14(_this.f7, _dafny.MultiSet.fromElements(_1446_v51, (_1447_v44)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_1447_v44).length))]), (_this).f6);
          })(_1444_v51, _1435_v44);
          let _nw245 = Array((new BigNumber(17)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw245.length); _i0_47++) {
            _nw245[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1445_v52 = _nw245;
          let _1449_v53;
          _1449_v53 = _module.D6.create_DC14(new BigNumber((_dafny.Seq.of(!((_1435_v44)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_1435_v44).length))]), _1444_v51)).length), _dafny.MultiSet.fromElements(_1428_v37), _this.f7);
          let _index267 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1445_v52).length));
          (_1445_v52)[_index267] = _1449_v53;
          let _index268 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1445_v52).length));
          (_1445_v52)[_index268] = _1449_v53;
          let _index269 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length));
          (_1434_v43)[_index269] = _module.__default.fm1((new BigNumber(135)).isLessThan((_1434_v43)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length))]), globalState);
        } else {
          let _1450_v54;
          let _nw246 = new _module.C0();
          _nw246.__ctor(_dafny.areEqual((_1422_v31)[_module.__default.safeIndex(new BigNumber(56), new BigNumber((_1422_v31).length))], _1418_v27));
          _1450_v54 = _nw246;
          let _1451_v55;
          _1451_v55 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11, (_this).f11);
          let _1452_v56;
          _1452_v56 = _module.D6.create_DC15((_1428_v37) === (_1428_v37), p0, (new BigNumber(857)).multipliedBy((_1420_v29)[_module.__default.safeIndex(p0, new BigNumber((_1420_v29).length))]), _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements((_1434_v43)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length))])).cardinality()), (_this).f6));
          let _pat_let_tv37 = _1419_v28;
          let _1453_v57;
          _1453_v57 = _dafny.Seq.of(((!(_1450_v54.f10)) ? (_1431_v40) : (function (_pat_let36_0) {
            return function (_1454_dt__update__tmp_h4) {
              return function (_pat_let37_0) {
                return function (_1455_dt__update_hcf51_h0) {
                  return _module.D10.create_DC24(_1455_dt__update_hcf51_h0);
                }(_pat_let37_0);
              }(_pat_let_tv37);
            }(_pat_let36_0);
          }(_1431_v40))), _1431_v40, _1431_v40, _1431_v40);
          let _1456_v58;
          _1456_v58 = _dafny.MultiSet.fromElements(_1450_v54.f10);
          let _rhs306 = ((_1434_v43)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length))]).minus((_this.f7).plus((_1434_v43)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length))]));
          let _rhs307 = _1456_v58;
          let _rhs308 = _1452_v56;
          let _rhs309 = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm39(new BigNumber((_dafny.Seq.of(p0)).length), new _dafny.CodePoint('h'.codePointAt(0)), p0, (_this).f6, globalState), _1431_v40), _dafny.Seq.Concat(_dafny.Seq.update(_1453_v57, _module.__default.safeIndex((_this).f6, new BigNumber((_1453_v57).length)), _1431_v40), _1453_v57));
          let _rhs310 = _dafny.Seq.UnicodeFromString("jl");
          let _lhs218 = globalState;
          _lhs218.f1 = _rhs306;
          _1451_v55 = _rhs307;
          _1452_v56 = _rhs308;
          _1453_v57 = _rhs309;
          _1418_v27 = _rhs310;
          let _rhs311 = new BigNumber(753);
          let _rhs312 = p0;
          let _rhs313 = p0;
          let _lhs219 = globalState;
          let _lhs220 = globalState;
          let _lhs221 = globalState;
          _lhs219.f1 = _rhs311;
          _lhs220.f1 = _rhs312;
          _lhs221.f1 = _rhs313;
          (globalState).f1 = ((_this.f7).plus((_1434_v43)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length))])).minus((_1434_v43)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_1434_v43).length))]);
          _1423_v32 = (_1423_v32).update((new BigNumber(356)).minus((((_1421_v30).contains(p0)) ? ((_1421_v30).get(p0)) : ((_this).f6))), _this.f7);
        }
      } else if (_source17.is_DC26) {
        let _1457___mcc_h0 = (_source17).cf52;
        let _1458_cf52 = _1457___mcc_h0;
        let _1459_v59;
        let _nw247 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1459_v59 = _nw247;
        let _1460_v60;
        _1460_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1458_cf52);
        let _index270 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1459_v59).length));
        (_1459_v59)[_index270] = (((_1460_v60).contains(_module.__default.fm1(_1458_cf52, globalState))) ? ((_1460_v60).get(_module.__default.fm1(_1458_cf52, globalState))) : ((_this).f11));
        let _index271 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1459_v59).length));
        (_1459_v59)[_index271] = (_this).f11;
        _1459_v59 = _1459_v59;
        _1428_v37 = (_this).f11;
        let _1461_v61;
        let _nw248 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _1461_v61 = _nw248;
        let _index272 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1461_v61).length));
        (_1461_v61)[_index272] = ((_1458_cf52) ? ((_this).f6) : ((_dafny.ZERO).minus(_this.f7)));
        let _index273 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1461_v61).length));
        (_1461_v61)[_index273] = ((_dafny.ZERO).minus(_this.f7)).minus((_this).f6);
      } else {
        let _1462___mcc_h1 = (_source17).cf51;
        let _1463_cf51 = _1462___mcc_h1;
        _1428_v37 = !(_1428_v37);
        let _1464_v62;
        _1464_v62 = _module.D5.create_DC11(_1429_v38);
        let _1465_v63;
        _1465_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1464_v62,(_this).f6);
        let _rhs314 = (new BigNumber(-995)).multipliedBy(_module.__default.safeModuloInt((_this).f6, new BigNumber(477)));
        let _rhs315 = _module.__default.fm0(_1428_v37, _1423_v32, (_this.f7).minus((_this).f6), globalState);
        let _rhs316 = (_1429_v38)[_module.__default.safeIndex((((_1465_v63).contains(_1464_v62)) ? ((_1465_v63).get(_1464_v62)) : ((_this).f6)), new BigNumber((_1429_v38).length))];
        let _lhs222 = globalState;
        _lhs222.f1 = _rhs314;
        _1428_v37 = _rhs315;
        _1428_v37 = _rhs316;
        let _1466_v64;
        _1466_v64 = _dafny.Set.fromElements(_1428_v37);
        let _1467_v65;
        _1467_v65 = _dafny.MultiSet.fromElements(_module.__default.fm0((_this).f11, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1466_v64).length),new BigNumber((_1429_v38).length)), new BigNumber((_1420_v29).length), globalState));
        if ((_1467_v65).IsDisjointFrom(_1467_v65)) {
          let _1468_v66;
          let _nw249 = new _module.C2();
          _nw249.__ctor((_this).f6, new BigNumber(383));
          _1468_v66 = _nw249;
          let _1469_v69;
          _1469_v69 = _module.D1.create_DC4((_this).f11);
          let _1470_v70;
          _1470_v70 = _module.D2.create_DC7((_this).f11, (_1420_v29)[_module.__default.safeIndex(_this.f7, new BigNumber((_1420_v29).length))], function () {
  let _coll40 = new _dafny.Map();
  for (const _compr_40 of _dafny.IntegerRange(new BigNumber(377), new BigNumber(-191))) {
    let _1471_v68 = _compr_40;
    if (((new BigNumber(377)).isLessThanOrEqualTo(_1471_v68)) && ((_1471_v68).isLessThan(new BigNumber(-191)))) {
      _coll40.push([_module.__default.safeDivisionInt(_1471_v68, _this.f7),_1428_v37]);
    }
  }
  return _coll40;
}(), _1469_v69);
          let _1472_v71;
          _1472_v71 = _module.D8.create_DC20(_module.__default.fm0(_1428_v37, function () {
  let _coll41 = new _dafny.Map();
  for (const _compr_41 of _dafny.IntegerRange(new BigNumber(135), new BigNumber(264))) {
    let _1473_v67 = _compr_41;
    if (((new BigNumber(135)).isLessThanOrEqualTo(_1473_v67)) && ((_1473_v67).isLessThan(new BigNumber(264)))) {
      _coll41.push([(_1473_v67).multipliedBy(new BigNumber((_1418_v27).length)),new BigNumber(277)]);
    }
  }
  return _coll41;
}(), (_this).f6, globalState), _1428_v37, (_1470_v70).dtor_cf13, false);
          let _1474_v72;
          _1474_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1423_v32).length),_1472_v71);
          _1474_v72 = (_1474_v72).update(p0, _1472_v71);
          _1428_v37 = (_1429_v38)[_module.__default.safeIndex(((_dafny.ZERO).minus(new BigNumber((_1429_v38).length))).plus(_this.f7), new BigNumber((_1429_v38).length))];
          let _1475_v73;
          let _nw250 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1475_v73 = _nw250;
          _1475_v73 = _1475_v73;
          let _1476_v74;
          _1476_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1428_v37,_1428_v37);
          let _1477_v75;
          _1477_v75 = _module.D0.create_DC0(_1466_v64);
          let _1478_v76;
          _1478_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1477_v75,_1428_v37);
          let _1479_v77;
          _1479_v77 = _module.D1.create_DC2(_1478_v76);
          let _pat_let_tv38 = _1477_v75;
          let _pat_let_tv39 = _1478_v76;
          let _1480_v78;
          _1480_v78 = _module.D2.create_DC6(function (_pat_let38_0) {
  return function (_1483_dt__update__tmp_h7) {
    return function (_pat_let41_0) {
      return function (_1484_dt__update_hcf4_h2) {
        return _module.D1.create_DC2(_1484_dt__update_hcf4_h2);
      }(_pat_let41_0);
    }(_pat_let_tv39);
  }(_pat_let38_0);
}(function (_pat_let39_0) {
  return function (_1481_dt__update__tmp_h6) {
    return function (_pat_let40_0) {
      return function (_1482_dt__update_hcf4_h1) {
        return _module.D1.create_DC2(_1482_dt__update_hcf4_h1);
      }(_pat_let40_0);
    }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv38,(_this).f11));
  }(_pat_let39_0);
}(_1479_v77)), (_1422_v31)[_module.__default.safeIndex(new BigNumber(56), new BigNumber((_1422_v31).length))], (_this).f6);
          (_1468_v66).m8(_1476_v74, _1480_v78, _dafny.Seq.Create(_module.__default.abs(new BigNumber(119)), ((_1485_cf51) => function (_1486_i3) {
            return _1485_cf51;
          })(_1463_cf51)), globalState);
        } else {
          let _1487_v79;
          _1487_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1428_v37,(_1466_v64).Difference(_1466_v64));
          let _1488_v80;
          let _nw251 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
          _1488_v80 = _nw251;
          let _1489_v81;
          _1489_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(268),(_this).f11);
          let _rhs317 = p0;
          let _rhs318 = _1487_v79;
          let _rhs319 = _module.__default.safeDivisionInt(_this.f7, ((_this).f6).plus(new BigNumber((_1489_v81).length)));
          let _rhs320 = ((_module.__default.fm0(_1428_v37, _1423_v32, new BigNumber((_1467_v65).cardinality()), globalState)) ? (_1488_v80) : (_1488_v80));
          let _lhs223 = globalState;
          let _lhs224 = globalState;
          _lhs223.f1 = _rhs317;
          _1487_v79 = _rhs318;
          _lhs224.f1 = _rhs319;
          _1488_v80 = _rhs320;
          let _1490_v82;
          let _nw252 = Array((new BigNumber(20)).toNumber());
          _nw252[(_dafny.ZERO).toNumber()] = _1428_v37;
          _nw252[(_dafny.ONE).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(2)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(3)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(4)).toNumber()] = !(_1428_v37);
          _nw252[(new BigNumber(5)).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(6)).toNumber()] = false;
          _nw252[(new BigNumber(7)).toNumber()] = !((_this).f11);
          _nw252[(new BigNumber(8)).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(9)).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(10)).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(11)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(12)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(13)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(14)).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(15)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(16)).toNumber()] = _1428_v37;
          _nw252[(new BigNumber(17)).toNumber()] = (_this).f11;
          _nw252[(new BigNumber(18)).toNumber()] = !((_this).f11);
          _nw252[(new BigNumber(19)).toNumber()] = _1428_v37;
          _1490_v82 = _nw252;
          let _1491_v83;
          let _nw253 = Array((new BigNumber(6)).toNumber());
          _nw253[(_dafny.ZERO).toNumber()] = _1490_v82;
          _nw253[(_dafny.ONE).toNumber()] = _1490_v82;
          _nw253[(new BigNumber(2)).toNumber()] = _1490_v82;
          _nw253[(new BigNumber(3)).toNumber()] = _1490_v82;
          _nw253[(new BigNumber(4)).toNumber()] = _1490_v82;
          _nw253[(new BigNumber(5)).toNumber()] = _1490_v82;
          _1491_v83 = _nw253;
          let _index274 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_1491_v83).length));
          (_1491_v83)[_index274] = _1490_v82;
          let _index275 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_1491_v83).length));
          (_1491_v83)[_index275] = _1490_v82;
          let _1492_v84;
          _1492_v84 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1490_v82);
          (globalState).f1 = new BigNumber(((_1492_v84).Merge(_1492_v84)).length);
          _1428_v37 = _module.__default.fm0(((_this).f6).isLessThanOrEqualTo((_this).f6), _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6), (_dafny.ZERO).minus(_this.f7), globalState);
          _1428_v37 = (_this).f11;
        }
        let _1493_v85;
        let _nw254 = new _module.C1();
        _nw254.__ctor(((false) ? (new _dafny.CodePoint('b'.codePointAt(0))) : (_1419_v28)), (new BigNumber((_1418_v27).length)).multipliedBy(new BigNumber((_1430_v39).length)), _this.f7);
        _1493_v85 = _nw254;
      }
      let _1494_v86;
      let _nw255 = Array((_dafny.ONE).toNumber());
      _nw255[(_dafny.ZERO).toNumber()] = (_this).f11;
      _1494_v86 = _nw255;
      _1494_v86 = _1494_v86;
      return;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f7 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f8 = [];
      this._f9 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f8, f9, f6, f7) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm4(globalState) {
      let _this = this;
      return false;
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1495_v0;
      _1495_v0 = new _dafny.CodePoint('p'.codePointAt(0));
      let _1496_v1;
      _1496_v1 = _dafny.Seq.UnicodeFromString("ybpygmgw");
      _1495_v0 = (_1496_v1)[_module.__default.safeIndex(((p0) ? (p2) : (p2)), new BigNumber((_1496_v1).length))];
      let _1497_v2;
      let _nw256 = Array((new BigNumber(12)).toNumber()).fill([]);
      _1497_v2 = _nw256;
      let _1498_v3;
      let _nw257 = Array((new BigNumber(26)).toNumber());
      _nw257[(_dafny.ZERO).toNumber()] = p0;
      _nw257[(_dafny.ONE).toNumber()] = p0;
      _nw257[(new BigNumber(2)).toNumber()] = p0;
      _nw257[(new BigNumber(3)).toNumber()] = p0;
      _nw257[(new BigNumber(4)).toNumber()] = p0;
      _nw257[(new BigNumber(5)).toNumber()] = p0;
      _nw257[(new BigNumber(6)).toNumber()] = p0;
      _nw257[(new BigNumber(7)).toNumber()] = p0;
      _nw257[(new BigNumber(8)).toNumber()] = p0;
      _nw257[(new BigNumber(9)).toNumber()] = p0;
      _nw257[(new BigNumber(10)).toNumber()] = p0;
      _nw257[(new BigNumber(11)).toNumber()] = p0;
      _nw257[(new BigNumber(12)).toNumber()] = p0;
      _nw257[(new BigNumber(13)).toNumber()] = p0;
      _nw257[(new BigNumber(14)).toNumber()] = p0;
      _nw257[(new BigNumber(15)).toNumber()] = p0;
      _nw257[(new BigNumber(16)).toNumber()] = p0;
      _nw257[(new BigNumber(17)).toNumber()] = true;
      _nw257[(new BigNumber(18)).toNumber()] = !(p0);
      _nw257[(new BigNumber(19)).toNumber()] = p0;
      _nw257[(new BigNumber(20)).toNumber()] = p0;
      _nw257[(new BigNumber(21)).toNumber()] = p0;
      _nw257[(new BigNumber(22)).toNumber()] = p0;
      _nw257[(new BigNumber(23)).toNumber()] = p0;
      _nw257[(new BigNumber(24)).toNumber()] = p0;
      _nw257[(new BigNumber(25)).toNumber()] = true;
      _1498_v3 = _nw257;
      let _index276 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length));
      (_1497_v2)[_index276] = _1498_v3;
      let _index277 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length));
      (_1497_v2)[_index277] = _1498_v3;
      let _1499_v4;
      _1499_v4 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
      let _1500_v5;
      _1500_v5 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_1499_v4).length), new BigNumber(387)),(_this).f6);
      r2 = new BigNumber((_1500_v5).length);
      let _hi14 = p2;
      for (let _1501_i0 = (_dafny.ZERO).minus((_this).f6); _1501_i0.isLessThan(_hi14); _1501_i0 = _1501_i0.plus(_dafny.ONE)) {
        let _1502_v6;
        let _nw258 = Array((new BigNumber(21)).toNumber()).fill([]);
        _1502_v6 = _nw258;
        let _1503_v7;
        _1503_v7 = _dafny.Set.fromElements(p0, p0);
        let _1504_v8;
        _1504_v8 = _module.D0.create_DC0(_1503_v7);
        let _1505_v9;
        _1505_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v8,p0);
        let _pat_let_tv40 = _1503_v7;
        let _1506_v11;
        _1506_v11 = _dafny.Set.fromElements(_1504_v8, function (_pat_let42_0) {
          return function (_1507_dt__update__tmp_h0) {
            return function (_pat_let43_0) {
              return function (_1508_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_1508_dt__update_hcf0_h0);
              }(_pat_let43_0);
            }(_pat_let_tv40);
          }(_pat_let42_0);
        }(_1504_v8), _module.D0.create_DC0(_1503_v7));
        let _1509_v13;
        _1509_v13 = _dafny.Seq.of(_1504_v8);
        let _1510_v14;
        _1510_v14 = _module.D1.create_DC2(_1505_v9);
        let _1511_v15;
        let _nw259 = Array((new BigNumber(21)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = _1505_v9;
        _nw259[(_dafny.ONE).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(2)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(3)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1504_v8,true);
        _nw259[(new BigNumber(5)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(6)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(7)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(8)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(9)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(10)).toNumber()] = function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of (_1506_v11).Elements) {
            let _1512_v10 = _compr_42;
            if ((_1506_v11).contains(_1512_v10)) {
              _coll42.push([_1512_v10,p0]);
            }
          }
          return _coll42;
        }();
        _nw259[(new BigNumber(11)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(12)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(13)).toNumber()] = function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of (_1509_v13).Elements) {
            let _1513_v12 = _compr_43;
            if (_dafny.Seq.contains(_1509_v13, _1513_v12)) {
              _coll43.push([_1513_v12,p0]);
            }
          }
          return _coll43;
        }();
        _nw259[(new BigNumber(14)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(15)).toNumber()] = (_1505_v9).update(_1504_v8, false);
        _nw259[(new BigNumber(16)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(17)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(18)).toNumber()] = _1505_v9;
        _nw259[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_1503_v7),p0);
        _nw259[(new BigNumber(20)).toNumber()] = (_1510_v14).dtor_cf4;
        _1511_v15 = _nw259;
        let _index278 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_1502_v6).length));
        (_1502_v6)[_index278] = ((p0) ? (_1511_v15) : (_1511_v15));
        let _1514_v16;
        _1514_v16 = _dafny.Seq.of(_1511_v15);
        let _index279 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_1502_v6).length));
        (_1502_v6)[_index279] = (_1514_v16)[_module.__default.safeIndex((_module.D1.create_DC3(p0, (_dafny.ZERO).minus(p1), _this.f7)).dtor_cf6, new BigNumber((_1514_v16).length))];
        let _1515_v17;
        _1515_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(65),(_this).f9);
        let _1516_v18;
        _1516_v18 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1501_i0)).length)));
        let _1517_v19;
        _1517_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm5(p0, p2, p2, p0, globalState));
        let _1518_v20;
        _1518_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f9).length),p0);
        let _1519_v21;
        _1519_v21 = _dafny.Seq.of(p0, p0);
        let _1520_v22;
        _1520_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1519_v21,_1501_i0);
        let _1521_v23;
        let _nw260 = Array((new BigNumber(18)).toNumber());
        _nw260[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw260[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((_this).f9, (_this).f9);
        _nw260[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat((_this).f9, (_this).f9);
        _nw260[(new BigNumber(3)).toNumber()] = (((_1515_v17).contains(_1501_i0)) ? ((_1515_v17).get(_1501_i0)) : ((_this).f9));
        _nw260[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw260[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw260[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_this).f9, (_1516_v18)[_module.__default.safeIndex(p2, new BigNumber((_1516_v18).length))]);
        _nw260[(new BigNumber(7)).toNumber()] = (_this).f9;
        _nw260[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw260[(new BigNumber(9)).toNumber()] = (((_1517_v19).contains(p0)) ? ((_1517_v19).get(p0)) : ((_this).f9));
        _nw260[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw260[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw260[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(p1, new BigNumber(821), (_dafny.ZERO).minus(new BigNumber((_1518_v20).length)), (_this).f6, p1);
        _nw260[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-17)), function (_1522_i1) {
          return (_this).f6;
        });
        _nw260[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat((_this).f9, (_this).f9);
        _nw260[(new BigNumber(15)).toNumber()] = (((_1515_v17).contains((((_1520_v22).contains(_1519_v21)) ? ((_1520_v22).get(_1519_v21)) : (_1501_i0)))) ? ((_1515_v17).get((((_1520_v22).contains(_1519_v21)) ? ((_1520_v22).get(_1519_v21)) : (_1501_i0)))) : ((_this).f9));
        _nw260[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat((_this).f9, (_this).f9);
        _nw260[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_1523_i2) {
          return (_this).f6;
        });
        _1521_v23 = _nw260;
        let _index280 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_1521_v23).length));
        (_1521_v23)[_index280] = _dafny.Seq.of((_this).f6);
        let _index281 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_1521_v23).length));
        (_1521_v23)[_index281] = ((p0) ? (_dafny.Seq.Concat(_module.__default.fm5(p0, p1, p2, p0, globalState), (_1516_v18)[_module.__default.safeIndex((_this).f6, new BigNumber((_1516_v18).length))])) : ((_this).f9));
        let _1524_v24;
        _1524_v24 = _dafny.MultiSet.fromElements(p2, _this.f7, (_this).f6, (_this).f6);
        let _1525_v25;
        _1525_v25 = _dafny.Set.fromElements(new BigNumber((_1518_v20).length), new BigNumber(-479), p2);
        let _1526_v26;
        let _init48 = ((_1527_p0) => function (_1528_i3) {
          return _module.__default.safeDivisionInt(_1528_i3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), ((_1529_p0) => function (_1530_i4) {
            return new BigNumber((_dafny.MultiSet.fromElements(!(_1529_p0))).cardinality());
          })(_1527_p0))).length));
        })(p0);
        let _nw261 = Array((new BigNumber(15)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw261.length); _i0_48++) {
          _nw261[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1526_v26 = _nw261;
        let _1531_v27;
        _1531_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1495_v0,_1526_v26);
        let _1532_v28;
        let _nw262 = Array((new BigNumber(23)).toNumber());
        _nw262[(_dafny.ZERO).toNumber()] = (_this).f6;
        _nw262[(_dafny.ONE).toNumber()] = p2;
        _nw262[(new BigNumber(2)).toNumber()] = new BigNumber((_1496_v1).length);
        _nw262[(new BigNumber(3)).toNumber()] = p1;
        _nw262[(new BigNumber(4)).toNumber()] = ((!(!(p0))) ? (new BigNumber(((_1521_v23)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1521_v23).length))]).length)) : (_1501_i0));
        _nw262[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("f")).length);
        _nw262[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_1501_i0);
        _nw262[(new BigNumber(7)).toNumber()] = _1501_i0;
        _nw262[(new BigNumber(8)).toNumber()] = _1501_i0;
        _nw262[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw262[(new BigNumber(10)).toNumber()] = (_this).f6;
        _nw262[(new BigNumber(11)).toNumber()] = (_1501_i0).minus((_this).f6);
        _nw262[(new BigNumber(12)).toNumber()] = ((_1521_v23)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1521_v23).length))])[_module.__default.safeIndex((_this).f6, new BigNumber(((_1521_v23)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1521_v23).length))]).length))];
        _nw262[(new BigNumber(13)).toNumber()] = _this.f7;
        _nw262[(new BigNumber(14)).toNumber()] = _1501_i0;
        _nw262[(new BigNumber(15)).toNumber()] = new BigNumber((_1517_v19).length);
        _nw262[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw262[(new BigNumber(17)).toNumber()] = (((_1524_v24).contains(new BigNumber((_1525_v25).length))) ? ((_1524_v24).get(new BigNumber((_1525_v25).length))) : (_1501_i0));
        _nw262[(new BigNumber(18)).toNumber()] = (_this).f6;
        _nw262[(new BigNumber(19)).toNumber()] = _this.f7;
        _nw262[(new BigNumber(20)).toNumber()] = new BigNumber((_1531_v27).length);
        _nw262[(new BigNumber(21)).toNumber()] = _module.__default.safeModuloInt(_this.f7, _this.f7);
        _nw262[(new BigNumber(22)).toNumber()] = p1;
        _1532_v28 = _nw262;
        _1526_v26 = (_module.D2.create_DC5(_1532_v28)).dtor_cf9;
        if (p0) {
          let _1533_v29;
          _1533_v29 = _dafny.MultiSet.fromElements(p0);
          let _1534_v30;
          let _nw263 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1534_v30 = _nw263;
          let _rhs321 = (((_1499_v4).contains(!(p0) || (p0))) ? ((_1499_v4).get(!(p0) || (p0))) : ((new BigNumber(-273)).isLessThan(p1)));
          let _rhs322 = ((_1533_v29).Union(_dafny.MultiSet.fromElements(p0))).Union(_1533_v29);
          let _rhs323 = _1534_v30;
          let _rhs324 = p0;
          r0 = _rhs321;
          _1533_v29 = _rhs322;
          _1534_v30 = _rhs323;
          r1 = _rhs324;
          let _1535_v31;
          let _nw264 = new _module.C0();
          _nw264.__ctor(!((_1525_v25).IsDisjointFrom(_1525_v25)));
          _1535_v31 = _nw264;
          _1496_v1 = _1496_v1;
          (globalState).f1 = ((_1501_i0).minus(new BigNumber(255))).plus(_this.f7);
          _1495_v0 = _1495_v0;
        } else {
          let _1536_v32;
          _1536_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1496_v1,p0);
          _1536_v32 = (_1536_v32).update(_dafny.Seq.UnicodeFromString("m"), p0);
          (_this).f7 = _module.__default.safeDivisionInt(_this.f7, _module.__default.safeDivisionInt(new BigNumber(853), new BigNumber((_1496_v1).length)));
          let _1537_v33;
          _1537_v33 = _dafny.Seq.of(_1498_v3, (_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))], _1498_v3);
          _1498_v3 = (_1537_v33)[_module.__default.safeIndex(_1501_i0, new BigNumber((_1537_v33).length))];
          r1 = p0;
          let _1538_v34;
          _1538_v34 = _dafny.MultiSet.fromElements(true, p0, p0, p0, p0);
          let _1539_v35;
          _1539_v35 = _1538_v34;
          let _1540_v36;
          _1540_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1538_v34);
          let _1541_v37;
          let _nw265 = Array((new BigNumber(24)).toNumber());
          _nw265[(_dafny.ZERO).toNumber()] = ((p0) ? (_1538_v34) : (_dafny.MultiSet.fromElements(p0, p0)));
          _nw265[(_dafny.ONE).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(2)).toNumber()] = _module.__default.fm6(globalState);
          _nw265[(new BigNumber(3)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(4)).toNumber()] = (_1539_v35);
          _nw265[(new BigNumber(5)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(!(p0), p0, p0);
          _nw265[(new BigNumber(7)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.FromArray(_1519_v21);
          _nw265[(new BigNumber(9)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(10)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs(p1));
          _nw265[(new BigNumber(12)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(13)).toNumber()] = (_1538_v34).Union(_dafny.MultiSet.fromElements(p0));
          _nw265[(new BigNumber(14)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(15)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(16)).toNumber()] = (_1538_v34).Intersect((((_1540_v36).contains(p0)) ? ((_1540_v36).get(p0)) : (_module.__default.fm6(globalState))));
          _nw265[(new BigNumber(17)).toNumber()] = (_dafny.MultiSet.fromElements(p0)).Difference(_1538_v34);
          _nw265[(new BigNumber(18)).toNumber()] = (_1538_v34).Intersect(_1538_v34);
          _nw265[(new BigNumber(19)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(20)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(21)).toNumber()] = _1538_v34;
          _nw265[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements(!(p0), p0, p0);
          _nw265[(new BigNumber(23)).toNumber()] = (_1538_v34).update(p0, _module.__default.abs(p2));
          _1541_v37 = _nw265;
          _1541_v37 = _1541_v37;
        }
      }
      let _1542_v38;
      let _nw266 = new _module.C0();
      _nw266.__ctor(p0);
      _1542_v38 = _nw266;
      let _arr0 = (_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))];
      let _index282 = _module.__default.safeIndex(new BigNumber(667), new BigNumber(((_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))]).length));
      _arr0[_index282] = _1542_v38.f10;
      let _arr1 = (_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))];
      let _index283 = _module.__default.safeIndex(new BigNumber(667), new BigNumber(((_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))]).length));
      _arr1[_index283] = p0;
      let _1543_v39;
      _1543_v39 = _dafny.MultiSet.fromElements(((_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))])[_module.__default.safeIndex(new BigNumber(667), new BigNumber(((_1497_v2)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1497_v2).length))]).length))], p0);
      let _1544_v40;
      _1544_v40 = _1543_v39;
      let _pat_let_tv41 = _1542_v38;
      r0 = function (_source18) {
        let _1545___mcc_h0 = _source18;
        let _1546_cf18 = _1545___mcc_h0;
        return _pat_let_tv41.f10;
      }(_1544_v40);
      r1 = _dafny.Seq.IsProperPrefixOf(_1496_v1, _dafny.Seq.Concat(_1496_v1, _dafny.Seq.UnicodeFromString("lrjbfvxeg")));
      r2 = p2;
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let _1547_v0;
      let _nw267 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _1547_v0 = _nw267;
      let _1548_v1;
      _1548_v1 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1549_v2;
      _1549_v2 = _dafny.Set.fromElements(_1548_v1);
      let _index284 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
      (_1547_v0)[_index284] = new BigNumber((_1549_v2).length);
      let _index285 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
      (_1547_v0)[_index285] = _this.f7;
      let _1550_v3;
      _1550_v3 = _dafny.Set.fromElements((_this).f6, _this.f7);
      _1550_v3 = (_1550_v3).Union(_1550_v3);
      let _1551_v4;
      _1551_v4 = _dafny.Seq.UnicodeFromString("wukhtqknx");
      let _1552_v5;
      _1552_v5 = _dafny.MultiSet.fromElements(_module.__default.fm2((_this).f9, _this.f7, _this.f7, !(p1), globalState), _1551_v4, _1551_v4, _1551_v4);
      _1552_v5 = _1552_v5;
      r0 = (p2) && (p1);
      let _1553_i0;
      _1553_i0 = _dafny.ZERO;
      L10: {
        while (!(_dafny.Seq.IsPrefixOf(_1551_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_1611_v1) => function (_1612_i1) {
          return _1611_v1;
        })(_1548_v1))))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1553_i0)) {
              break L10;
            }
            _1553_i0 = (_1553_i0).plus(_dafny.ONE);
            let _1554_v6;
            _1554_v6 = _dafny.Set.fromElements(p1, p1);
            let _1555_v7;
            _1555_v7 = _module.D0.create_DC0(_1554_v6);
            let _1556_v8;
            _1556_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1555_v7,p1);
            let _1557_v9;
            _1557_v9 = _module.D1.create_DC2(_1556_v8);
            let _1558_v10;
            _1558_v10 = _module.D2.create_DC6(_1557_v9, _dafny.Seq.UnicodeFromString("jgkjag"), new BigNumber((_1554_v6).length));
            let _1559_v11;
            _1559_v11 = _dafny.Seq.of(p1, false);
            let _1560_v12;
            _1560_v12 = _module.D1.create_DC3(p2, new BigNumber((_1559_v11).length), (_this).f6);
            let _pat_let_tv42 = p1;
            let _pat_let_tv43 = _1556_v8;
            let _pat_let_tv44 = _1557_v9;
            let _pat_let_tv45 = _1551_v4;
            let _1561_v13;
            let _nw268 = Array((new BigNumber(25)).toNumber());
            _nw268[(_dafny.ZERO).toNumber()] = _1558_v10;
            _nw268[(_dafny.ONE).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(2)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(3)).toNumber()] = _module.D2.create_DC6(_1557_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(759)), ((_1562_v1) => function (_1563_i2) {
  return _1562_v1;
})(_1548_v1)), (function (_pat_let44_0) {
  return function (_1564_dt__update__tmp_h0) {
    return function (_pat_let45_0) {
      return function (_1565_dt__update_hcf5_h0) {
        return function (_pat_let46_0) {
          return function (_1566_dt__update_hcf7_h0) {
            return _module.D1.create_DC3(_1565_dt__update_hcf5_h0, (_1564_dt__update__tmp_h0).dtor_cf6, _1566_dt__update_hcf7_h0);
          }(_pat_let46_0);
        }(_this.f7);
      }(_pat_let45_0);
    }(_pat_let_tv42);
  }(_pat_let44_0);
}(_1560_v12)).dtor_cf7);
            _nw268[(new BigNumber(4)).toNumber()] = _module.D2.create_DC6(_1557_v9, _1551_v4, _this.f7);
            _nw268[(new BigNumber(5)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(6)).toNumber()] = function (_pat_let47_0) {
              return function (_1567_dt__update__tmp_h1) {
                return function (_pat_let48_0) {
                  return function (_1568_dt__update_hcf12_h0) {
                    return function (_pat_let49_0) {
                      return function (_1569_dt__update_hcf10_h0) {
                        return _module.D2.create_DC6(_1569_dt__update_hcf10_h0, (_1567_dt__update__tmp_h1).dtor_cf11, _1568_dt__update_hcf12_h0);
                      }(_pat_let49_0);
                    }(_module.D1.create_DC2(_pat_let_tv43));
                  }(_pat_let48_0);
                }(_this.f7);
              }(_pat_let47_0);
            }(_1558_v10);
            _nw268[(new BigNumber(7)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(8)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(9)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(10)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(11)).toNumber()] = ((p1) ? (_1558_v10) : (_1558_v10));
            _nw268[(new BigNumber(12)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(13)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(14)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(15)).toNumber()] = function (_pat_let50_0) {
              return function (_1570_dt__update__tmp_h2) {
                return function (_pat_let51_0) {
                  return function (_1571_dt__update_hcf10_h1) {
                    return function (_pat_let52_0) {
                      return function (_1572_dt__update_hcf11_h0) {
                        return _module.D2.create_DC6(_1571_dt__update_hcf10_h1, _1572_dt__update_hcf11_h0, (_1570_dt__update__tmp_h2).dtor_cf12);
                      }(_pat_let52_0);
                    }(_pat_let_tv45);
                  }(_pat_let51_0);
                }(_pat_let_tv44);
              }(_pat_let50_0);
            }(_1558_v10);
            _nw268[(new BigNumber(16)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(17)).toNumber()] = _module.__default.fm7(globalState);
            _nw268[(new BigNumber(18)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(19)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(20)).toNumber()] = _module.__default.fm7(globalState);
            _nw268[(new BigNumber(21)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(22)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(23)).toNumber()] = _1558_v10;
            _nw268[(new BigNumber(24)).toNumber()] = _1558_v10;
            _1561_v13 = _nw268;
            _1561_v13 = _1561_v13;
            let _rhs325 = !(false);
            let _rhs326 = p2;
            r0 = _rhs325;
            r0 = _rhs326;
            let _1573_v14;
            let _nw269 = new _module.C0();
            _nw269.__ctor(p1);
            _1573_v14 = _nw269;
            _1573_v14 = _1573_v14;
            let _1574_v15;
            _1574_v15 = _module.D2.create_DC5(_1547_v0);
            let _source19 = _1574_v15;
            if (_source19.is_DC6) {
              let _1575___mcc_h0 = (_source19).cf10;
              let _1576___mcc_h1 = (_source19).cf11;
              let _1577___mcc_h2 = (_source19).cf12;
              let _1578_cf12 = _1577___mcc_h2;
              let _1579_cf11 = _1576___mcc_h1;
              let _1580_cf10 = _1575___mcc_h0;
              (_1573_v14).f10 = !_dafny.Seq.contains(_dafny.Seq.Concat((_this).f9, (_this).f9), _1578_cf12);
              _1551_v4 = _1551_v4;
              let _index286 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
              (_1547_v0)[_index286] = (new BigNumber((_dafny.Seq.of(_1573_v14, _1573_v14)).length)).plus((_1578_cf12).minus((_this).f6));
              (globalState).f1 = (_this).f6;
            } else if (_source19.is_DC7) {
              let _1581___mcc_h3 = (_source19).cf13;
              let _1582___mcc_h4 = (_source19).cf14;
              let _1583___mcc_h5 = (_source19).cf15;
              let _1584___mcc_h6 = (_source19).cf16;
              let _1585_cf16 = _1584___mcc_h6;
              let _1586_cf15 = _1583___mcc_h5;
              let _1587_cf14 = _1582___mcc_h4;
              let _1588_cf13 = _1581___mcc_h3;
              _1547_v0 = _1547_v0;
              let _1589_v16;
              _1589_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_1551_v4).length));
              let _1590_v17;
              _1590_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_1589_v16).Merge(_1589_v16));
              let _1591_v18;
              _1591_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1573_v14.f10,!(false));
              let _index287 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
              let _rhs327 = new BigNumber(((_1591_v18).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1573_v14.f10,p2))).length);
              let _rhs328 = _1588_cf13;
              let _rhs329 = _1559_v11;
              let _rhs330 = new BigNumber(-312);
              let _rhs331 = _1590_v17;
              let _lhs225 = _this;
              let _lhs226 = _1573_v14;
              let _lhs227 = _1547_v0;
              let _lhs228 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
              _lhs225.f7 = _rhs327;
              _lhs226.f10 = _rhs328;
              _1559_v11 = _rhs329;
              _lhs227[_lhs228] = _rhs330;
              _1590_v17 = _rhs331;
              let _rhs332 = p2;
              let _rhs333 = p2;
              let _lhs229 = _1573_v14;
              let _lhs230 = _1573_v14;
              _lhs229.f10 = _rhs332;
              _lhs230.f10 = _rhs333;
              let _1592_v20;
              _1592_v20 = _dafny.Set.fromElements(_1551_v4);
              _1588_cf13 = !((_1592_v20).contains(_module.__default.fm2((_this).f9, new BigNumber((function () {
                let _coll44 = new _dafny.Set();
                for (const _compr_44 of (_dafny.Map.Empty.slice().updateUnsafe(_1558_v10,_1588_cf13)).Keys.Elements) {
                  let _1593_v19 = _compr_44;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_1558_v10,_1588_cf13)).contains(_1593_v19)) {
                    _coll44.add(_1593_v19);
                  }
                }
                return _coll44;
              }()).length), _this.f7, _1573_v14.f10, globalState)));
            } else if (_source19.is_DC5) {
              let _1594___mcc_h7 = (_source19).cf9;
              let _1595_cf9 = _1594___mcc_h7;
              let _1596_v21;
              _1596_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1554_v6).IsProperSubsetOf(_1554_v6),_1547_v0);
              _1596_v21 = (_1596_v21).update(_1573_v14.f10, _1595_cf9);
              _1548_v1 = _module.__default.fm8(_dafny.Seq.Concat(_1551_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(200)), ((_1597_v1) => function (_1598_i3) {
                return _1597_v1;
              })(_1548_v1))), _1548_v1, _1573_v14.f10, (_this).fm4(globalState), globalState);
              let _1599_v22;
              _1599_v22 = _dafny.Seq.of((_1558_v10).dtor_cf11);
              let _1600_v23;
              let _nw270 = Array((new BigNumber(7)).toNumber());
              _nw270[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), ((_1601_v1) => function (_1602_i4) {
                return _1601_v1;
              })(_1548_v1));
              _nw270[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_1551_v4, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f9)).length), new BigNumber((_1551_v4).length)), _1548_v1);
              _nw270[(new BigNumber(2)).toNumber()] = _1551_v4;
              _nw270[(new BigNumber(3)).toNumber()] = _1551_v4;
              _nw270[(new BigNumber(4)).toNumber()] = _1551_v4;
              _nw270[(new BigNumber(5)).toNumber()] = (_1599_v22)[_module.__default.safeIndex(_this.f7, new BigNumber((_1599_v22).length))];
              _nw270[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("sfyudb");
              _1600_v23 = _nw270;
              let _1603_v24;
              _1603_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))],_1600_v23);
              _1603_v24 = (_1603_v24).update((_this).f6, _1600_v23);
              let _1604_v25;
              _1604_v25 = _dafny.Seq.of(((_1573_v14.f10) ? (_1573_v14) : (_1573_v14)), _1573_v14, _1573_v14);
              let _1605_v26;
              _1605_v26 = _dafny.MultiSet.fromElements(_1573_v14.f10, _1573_v14.f10);
              let _index288 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
              let _rhs334 = (_1604_v25)[_module.__default.safeIndex((_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))], new BigNumber((_1604_v25).length))];
              let _rhs335 = (_module.__default.safeModuloInt((_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))], (_this).f6)).multipliedBy((_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))]);
              let _rhs336 = _dafny.Seq.of(_1551_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_1606_v1) => function (_1607_i5) {
                return _1606_v1;
              })(_1548_v1)));
              let _rhs337 = _1548_v1;
              let _rhs338 = ((_this).f6).isEqualTo((new BigNumber((_1605_v26).cardinality())).multipliedBy((_this).f6));
              let _lhs231 = _1547_v0;
              let _lhs232 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
              _1573_v14 = _rhs334;
              _lhs231[_lhs232] = _rhs335;
              _1599_v22 = _rhs336;
              _1548_v1 = _rhs337;
              r0 = _rhs338;
            } else {
              let _1608___mcc_h8 = (_source19).cf17;
              let _1609_cf17 = _1608___mcc_h8;
              let _index289 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length));
              (_1547_v0)[_index289] = (_this).f6;
              (globalState).f1 = _this.f7;
              let _1610_v27;
              _1610_v27 = _dafny.Set.fromElements(_module.D2.create_DC5(_1547_v0), _1574_v15, _1574_v15, _1574_v15);
              _1610_v27 = ((_1573_v14.f10) ? (_1610_v27) : (_1610_v27));
              _1559_v11 = _1559_v11;
            }
          }
        }
      }
      let _1613_v28;
      _1613_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1551_v4).length),(_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))]);
      let _1614_v29;
      _1614_v29 = _dafny.Set.fromElements(p1);
      let _1615_v30;
      _1615_v30 = _module.D0.create_DC0(_1614_v29);
      let _1616_v31;
      _1616_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1615_v30,p1);
      let _1617_v32;
      _1617_v32 = _module.D1.create_DC2(_1616_v31);
      let _1618_v33;
      _1618_v33 = _module.D2.create_DC6(_1617_v32, _1551_v4, new BigNumber(229));
      let _1619_v34;
      _1619_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_1613_v28).contains((_this).f6)) ? ((_1613_v28).get((_this).f6)) : (new BigNumber((_module.__default.fm9(new BigNumber((_1551_v4).length), globalState)).length))),_module.D2.create_DC8(_1618_v33));
      let _1620_v35;
      _1620_v35 = _dafny.MultiSet.fromElements(p1);
      let _1621_v36;
      _1621_v36 = _dafny.Seq.of(p2, p2);
      let _1622_v37;
      _1622_v37 = _module.D2.create_DC8(_module.D2.create_DC8(_module.D2.create_DC8(_1618_v33)));
      _1619_v34 = (_1619_v34).update((_dafny.ZERO).minus(_module.__default.safeModuloInt((((_1620_v35).contains((_1621_v36)[_module.__default.safeIndex((_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))], new BigNumber((_1621_v36).length))])) ? ((_1620_v35).get((_1621_v36)[_module.__default.safeIndex((_1547_v0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1547_v0).length))], new BigNumber((_1621_v36).length))])) : (_this.f7)), new BigNumber((_1551_v4).length))), _1622_v37);
      r0 = p1;
      let _1623_v38;
      let _init49 = ((_1624_v29) => function (_1625_i6) {
        return _module.D0.create_DC0(_1624_v29);
      })(_1614_v29);
      let _nw271 = Array((new BigNumber(20)).toNumber());
      for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw271.length); _i0_49++) {
        _nw271[_i0_49] = _init49(new BigNumber(_i0_49));
      }
      _1623_v38 = _nw271;
      r1 = _1623_v38;
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let _1626_v0;
      _1626_v0 = true;
      let _1627_v1;
      _1627_v1 = _dafny.Seq.of(_1626_v0);
      let _1628_v2;
      _1628_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1627_v1, _module.__default.safeIndex((_this).f6, new BigNumber((_1627_v1).length)), _1626_v0),(_this).f6);
      let _1629_v3;
      _1629_v3 = _dafny.MultiSet.fromElements((_this).f6);
      (_this).f7 = (((_1628_v2).contains(_dafny.Seq.Concat(_1627_v1, _1627_v1))) ? ((_1628_v2).get(_dafny.Seq.Concat(_1627_v1, _1627_v1))) : (_module.__default.safeDivisionInt(new BigNumber(((_1629_v3).update(new BigNumber(957), _module.__default.abs((_this).f6))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("adkmbscvj")).length))));
      let _1630_i0;
      _1630_i0 = _dafny.ZERO;
      L11: {
        while (_1626_v0) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1630_i0)) {
              break L11;
            }
            _1630_i0 = (_1630_i0).plus(_dafny.ONE);
            let _1631_v4;
            _1631_v4 = _dafny.Seq.UnicodeFromString("lvswxolf");
            _1631_v4 = _module.__default.fm2((_this).f9, (((_1627_v1)[_module.__default.safeIndex(_this.f7, new BigNumber((_1627_v1).length))]) ? (_this.f7) : ((_this).f6)), (_this).f6, !(_1626_v0), globalState);
            let _1632_v5;
            let _nw272 = new _module.C0();
            _nw272.__ctor(_1626_v0);
            _1632_v5 = _nw272;
            let _1633_v6;
            _1633_v6 = _module.D0.create_DC1(true, p0, (p1).isEqualTo(p1));
            let _source20 = _1633_v6;
            if (_source20.is_DC1) {
              let _1634___mcc_h0 = (_source20).cf1;
              let _1635___mcc_h1 = (_source20).cf2;
              let _1636___mcc_h2 = (_source20).cf3;
              let _1637_cf3 = _1636___mcc_h2;
              let _1638_cf2 = _1635___mcc_h1;
              let _1639_cf1 = _1634___mcc_h0;
              let _1640_v7;
              let _nw273 = Array((_dafny.ONE).toNumber());
              _1640_v7 = _nw273;
              let _index290 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1640_v7).length));
              (_1640_v7)[_index290] = _1632_v5;
              let _1641_v8;
              _1641_v8 = _1632_v5;
              let _1642_v9;
              _1642_v9 = _dafny.Seq.of((_1641_v8), _1632_v5, _1632_v5, _1632_v5);
              let _index291 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1640_v7).length));
              (_1640_v7)[_index291] = (_1642_v9)[_module.__default.safeIndex(new BigNumber((_module.__default.fm10(_1632_v5.f10, _1626_v0, _1627_v1, true, globalState)).length), new BigNumber((_1642_v9).length))];
              let _index292 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((p0).length));
              (p0)[_index292] = _1626_v0;
              let _index293 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((p0).length));
              (p0)[_index293] = _1637_cf3;
              let _rhs339 = _1631_v4;
              let _rhs340 = _1632_v5.f10;
              let _rhs341 = !(_1629_v3).equals((_1629_v3).Difference(_1629_v3));
              let _lhs233 = _1632_v5;
              _1631_v4 = _rhs339;
              _1639_cf1 = _rhs340;
              _lhs233.f10 = _rhs341;
              let _1643_v10;
              let _nw274 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
              _1643_v10 = _nw274;
              let _index294 = _module.__default.safeIndex(new BigNumber(393), new BigNumber(((_this).f8).length));
              ((_this).f8)[_index294] = _1643_v10;
              let _index295 = _module.__default.safeIndex(new BigNumber(393), new BigNumber(((_this).f8).length));
              ((_this).f8)[_index295] = _1643_v10;
            } else {
              let _1644___mcc_h3 = (_source20).cf0;
              let _1645_cf0 = _1644___mcc_h3;
              (globalState).f1 = p1;
              _1631_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1631_v4, _1631_v4), _1631_v4);
              let _1646_v11;
              let _init50 = function (_1647_i1) {
                return _module.__default.safeDivisionInt(_1647_i1, new BigNumber(-444));
              };
              let _nw275 = Array((new BigNumber(9)).toNumber());
              for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw275.length); _i0_50++) {
                _nw275[_i0_50] = _init50(new BigNumber(_i0_50));
              }
              _1646_v11 = _nw275;
              let _1648_v12;
              let _nw276 = new _module.C2();
              _nw276.__ctor(new BigNumber((_1627_v1).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_1649_i2) {
                return (_this).f6;
              })).length));
              _1648_v12 = _nw276;
              let _1650_v13;
              _1650_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1648_v12,true);
              let _1651_v14;
              _1651_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1626_v0,new BigNumber((_1650_v13).length));
              let _index296 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1646_v11).length));
              (_1646_v11)[_index296] = new BigNumber((_1651_v14).length);
              let _index297 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
              (p0)[_index297] = _1626_v0;
              let _index298 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1646_v11).length));
              let _index299 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
              let _rhs342 = new BigNumber((_1631_v4).length);
              let _rhs343 = _1632_v5.f10;
              let _lhs234 = _1646_v11;
              let _lhs235 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1646_v11).length));
              let _lhs236 = p0;
              let _lhs237 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
              _lhs234[_lhs235] = _rhs342;
              _lhs236[_lhs237] = _rhs343;
              let _1652_v15;
              let _nw277 = new _module.C0();
              _nw277.__ctor(!(_1632_v5.f10));
              _1652_v15 = _nw277;
            }
            let _1653_v16;
            _1653_v16 = _dafny.Set.fromElements(true, _1632_v5.f10, (new BigNumber(-182)).isEqualTo(_this.f7));
            let _1654_v17;
            let _init51 = function (_1655_i3) {
              return (_1655_i3).plus((_this).f6);
            };
            let _nw278 = Array((new BigNumber(7)).toNumber());
            for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw278.length); _i0_51++) {
              _nw278[_i0_51] = _init51(new BigNumber(_i0_51));
            }
            _1654_v17 = _nw278;
            let _index300 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1654_v17).length));
            (_1654_v17)[_index300] = _this.f7;
            let _index301 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1654_v17).length));
            let _rhs344 = _1653_v16;
            let _rhs345 = p1;
            let _lhs238 = _1654_v17;
            let _lhs239 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1654_v17).length));
            _1653_v16 = _rhs344;
            _lhs238[_lhs239] = _rhs345;
          }
        }
      }
      let _1656_v18;
      let _nw279 = new _module.C0();
      _nw279.__ctor(false);
      _1656_v18 = _nw279;
      let _1657_v19;
      _1657_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_1656_v18, _1656_v18)).length)))).cardinality()))).length));
      let _1658_v20;
      _1658_v20 = _dafny.Seq.UnicodeFromString("qmtta");
      let _1659_v21;
      _1659_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v18.f10,p1);
      let _1660_v22;
      _1660_v22 = _dafny.MultiSet.fromElements(_1626_v0);
      let _1661_v23;
      _1661_v23 = _module.D6.create_DC14(_this.f7, _1660_v22, new BigNumber((_dafny.MultiSet.fromElements(_this.f7)).cardinality()));
      let _1662_v24;
      let _nw280 = Array((new BigNumber(26)).toNumber());
      _nw280[(_dafny.ZERO).toNumber()] = (_this).f9;
      _nw280[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_this).f6, ((_this).f9)[_module.__default.safeIndex(new BigNumber((_1657_v19).length), new BigNumber(((_this).f9).length))], (_this).f6);
      _nw280[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(p1);
      _nw280[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_1663_i5) {
        return (_this).f6;
      });
      _nw280[(new BigNumber(4)).toNumber()] = _dafny.Seq.update((_this).f9, _module.__default.safeIndex(_this.f7, new BigNumber(((_this).f9).length)), new BigNumber((_1658_v20).length));
      _nw280[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_this).f9, _dafny.Seq.of(_this.f7, new BigNumber(-262)));
      _nw280[(new BigNumber(6)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(7)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(8)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(9)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(10)).toNumber()] = _module.__default.fm5(_1656_v18.f10, (((_1659_v21).contains(_1626_v0)) ? ((_1659_v21).get(_1626_v0)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1626_v0,p1)).length))), _module.__default.fm1(_1656_v18.f10, globalState), _1626_v0, globalState);
      _nw280[(new BigNumber(11)).toNumber()] = _dafny.Seq.update((_this).f9, _module.__default.safeIndex(_this.f7, new BigNumber(((_this).f9).length)), (_1661_v23).dtor_cf23);
      _nw280[(new BigNumber(12)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(13)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(14)).toNumber()] = _module.__default.fm5(true, new BigNumber(-647), p1, _1656_v18.f10, globalState);
      _nw280[(new BigNumber(15)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(16)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(17)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(18)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.fm1(false, globalState)), _this.f7, (_this).f6, (_this).f6);
      _nw280[(new BigNumber(19)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(20)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(21)).toNumber()] = _module.__default.fm35(_this.f7, _this.f7, globalState);
      _nw280[(new BigNumber(22)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(23)).toNumber()] = (_this).f9;
      _nw280[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm5(_1626_v0, p1, _this.f7, _1626_v0, globalState), (_this).f9);
      _nw280[(new BigNumber(25)).toNumber()] = (_this).f9;
      _1662_v24 = _nw280;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1662_v24).length))) {
        let _1664_i4 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1664_i4)) && ((_1664_i4).isLessThan(new BigNumber((_1662_v24).length))))) {
          (_1662_v24)[(_1664_i4)] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(835)), ((_1665_v21) => function (_1666_i6) {
            return new BigNumber((_1665_v21).length);
          })(_1659_v21)), (_this).f9);
        }
      }
      let _1667_v25;
      _1667_v25 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1668_v26;
      let _nw281 = new _module.C1();
      _nw281.__ctor(_module.__default.fm8(_1658_v20, _1667_v25, _1656_v18.f10, !(_1656_v18.f10), globalState), new BigNumber((_1658_v20).length), _this.f7);
      _1668_v26 = _nw281;
      _1668_v26 = _1668_v26;
      let _1669_v27;
      let _nw282 = new _module.C3();
      _nw282.__ctor(_module.__default.safeModuloInt(new BigNumber((_module.__default.fm2(_dafny.Seq.of(new BigNumber(24)), _this.f7, (_dafny.ZERO).minus(_this.f7), _1626_v0, globalState)).length), _this.f7), new BigNumber(876));
      _1669_v27 = _nw282;
      let _1670_v28;
      let _nw283 = new _module.C3();
      _nw283.__ctor(_module.__default.safeDivisionInt((_1669_v27).fm15(_this.f7, p1, globalState), (_this).f6), ((_this).f9)[_module.__default.safeIndex(p1, new BigNumber(((_this).f9).length))]);
      _1670_v28 = _nw283;
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
