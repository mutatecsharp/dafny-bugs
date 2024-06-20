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
      if (false) {
        return new BigNumber(34);
      } else {
        return _module.__default.safeDivisionInt(new BigNumber(419), new BigNumber((_dafny.Seq.of(true, !(false), false, false, true)).length));
      }
    };
    static fm1(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(447),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(881),new BigNumber(850))).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-920), new BigNumber(636))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(-920)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(636)))) {
            _coll0.push([(_0_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ff")).length)),new BigNumber(96)]);
          }
        }
        return _coll0;
      }()));
    };
    static fm3(p0, globalState) {
      let _source0 = _module.D3.create_DC11();
      if (_source0.is_DC11) {
        if (true) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }
      } else {
        let _1___mcc_h0 = (_source0).cf18;
        let _2_cf18 = _1___mcc_h0;
        return new _dafny.CodePoint('t'.codePointAt(0));
      }
    };
    static fm4(p0, p1, globalState) {
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-89)), function (_3_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length), new BigNumber(-546))).isLessThan(_module.__default.safeDivisionInt(new BigNumber(-556), new BigNumber(756)));
    };
    static fm5(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), function (_4_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length), new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality())),false);
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(true, (new BigNumber(-765)).isLessThan(new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of((_module.D1.create_DC4(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length), false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-523),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ylexru")).length))))).dtor_cf9, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length))).length))).Elements) {
          let _5_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of((_module.D1.create_DC4(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length), false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-523),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ylexru")).length))))).dtor_cf9, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length))).length)), _5_v0)) {
            _coll1.push([(_5_v0).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("djl")).length), new BigNumber((_dafny.Seq.of(false, !(true), false)).length))).length)),new BigNumber(529)]);
          }
        }
        return _coll1;
      }()).length)), (true) || (!(true)));
    };
    static fm7(p0, globalState) {
      return _dafny.Seq.of(!((_dafny.MultiSet.fromElements(true)).IsProperSubsetOf(_dafny.MultiSet.fromElements(false, true))));
    };
    static fm8(p0, p1, globalState) {
      return _module.D1.create_DC3((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-409), new BigNumber(468))) {
    let _6_v0 = _compr_2;
    if (((new BigNumber(-409)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(468)))) {
      _coll2.push([(_6_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("uicbdh")).length))),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(687), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(550)), function (_7_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })).length))).cardinality())]);
    }
  }
  return _coll2;
}()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber(-962))), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-196), new BigNumber(427)), _dafny.Seq.of(new BigNumber(223)))).length)), ((false) ? (_module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("tlkawrtbe")).length), false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(124)), function (_8_i1) {
  return new _dafny.CodePoint('t'.codePointAt(0));
}), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(675),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-276),false)))) : (_module.D0.create_DC0(new BigNumber(205), false, _dafny.Seq.UnicodeFromString("xhtbrf"), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bq"))).length))).cardinality()),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-318),true))))));
    };
    static fm9(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(807),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wdqdlcew")).length), new BigNumber(266), new BigNumber(135))).cardinality()), _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("f")).length), false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(839)), function (_9_i0) {
  return new _dafny.CodePoint('k'.codePointAt(0));
}), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(717),function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(364), new BigNumber(851))) {
    let _10_v0 = _compr_3;
    if (((new BigNumber(364)).isLessThanOrEqualTo(_10_v0)) && ((_10_v0).isLessThan(new BigNumber(851)))) {
      _coll3.push([(_10_v0).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_module.D0.create_DC0(new BigNumber(-898), true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), function (_11_i1) {
  return new _dafny.CodePoint('v'.codePointAt(0));
}), function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(619), new BigNumber(917))) {
    let _12_v1 = _compr_4;
    if (((new BigNumber(619)).isLessThanOrEqualTo(_12_v1)) && ((_12_v1).isLessThan(new BigNumber(917)))) {
      _coll4.push([(_12_v1).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("fcscl"))).length)),function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(712), new BigNumber(236))) {
          let _13_v2 = _compr_5;
          if (((new BigNumber(712)).isLessThanOrEqualTo(_13_v2)) && ((_13_v2).isLessThan(new BigNumber(236)))) {
            _coll5.push([_module.__default.safeDivisionInt(_13_v2, new BigNumber((_dafny.Seq.UnicodeFromString("qokpwk")).length)),false]);
          }
        }
        return _coll5;
      }()]);
    }
  }
  return _coll4;
}()))).length))),false]);
    }
  }
  return _coll3;
}())))), _dafny.Seq.of(_module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-962),new BigNumber(((_module.D5.create_DC17(function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-634), new BigNumber(9))) {
    let _14_v3 = _compr_6;
    if (((new BigNumber(-634)).isLessThanOrEqualTo(_14_v3)) && ((_14_v3).isLessThan(new BigNumber(9)))) {
      _coll6.add(_module.__default.safeModuloInt(_14_v3, new BigNumber((_dafny.Seq.UnicodeFromString("bk")).length)));
    }
  }
  return _coll6;
}())).dtor_cf29).length)), new BigNumber(662), _module.D0.create_DC0(new BigNumber(164), false, _dafny.Seq.UnicodeFromString("newt"), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dxly")).length)),false)).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(447),true)))), _module.D1.create_DC3(function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(620), new BigNumber(477))) {
    let _15_v4 = _compr_7;
    if (((new BigNumber(620)).isLessThanOrEqualTo(_15_v4)) && ((_15_v4).isLessThan(new BigNumber(477)))) {
      _coll7.push([_module.__default.safeModuloInt(_15_v4, new BigNumber(77)),new BigNumber(486)]);
    }
  }
  return _coll7;
}(), new BigNumber(-215), _module.D0.create_DC0(new BigNumber(651), false, _dafny.Seq.UnicodeFromString("jhxxrjsnl"), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(295),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-593),true)))), _module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(414),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(356)), function (_16_i2) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length), _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(867))).length), true, _dafny.Seq.UnicodeFromString("ouh"), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(207),false))))));
    };
    static fm10(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, false, true, false),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), function (_17_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),new BigNumber(-742))).Merge((_module.D6.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), true)),new BigNumber(916)))).dtor_cf33));
    };
    static fm11(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(758),function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),true)).Keys.Elements) {
          let _18_v0 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),true)).contains(_18_v0)) {
            _coll8.push([_module.__default.safeDivisionInt(_18_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(623),false)).length)),true]);
          }
        }
        return _coll8;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bre")).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-798),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-689),false))));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mcyqkio"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("egwe"), _dafny.Seq.UnicodeFromString("sewnuy")));
    };
    static fm13(globalState) {
      return _module.D1.create_DC4(new BigNumber(795), !(!((new BigNumber((function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(357), new BigNumber(568))) {
    let _19_v0 = _compr_9;
    if (((new BigNumber(357)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(568)))) {
      _coll9.push([_module.__default.safeModuloInt(_19_v0, new BigNumber(793)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_20_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })).length))]);
    }
  }
  return _coll9;
}()).length)).isEqualTo(new BigNumber(907)))), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-463)), function (_21_i1) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(343))).length))).length), new BigNumber(-382), (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll10 = new _dafny.Set();
  for (const _compr_10 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), function (_22_i2) {
    return new _dafny.CodePoint('r'.codePointAt(0));
  })).Elements) {
    let _23_v1 = _compr_10;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), function (_22_i2) {
      return new _dafny.CodePoint('r'.codePointAt(0));
    }), _23_v1)) {
      _coll10.add(_23_v1);
    }
  }
  return _coll10;
}()).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),_dafny.MultiSet.fromElements(new BigNumber(726)))));
    };
    static fm14(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(-785)), _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sxmprhotv")).length))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, true, false)).length), new BigNumber(378), (_dafny.ZERO).minus(new BigNumber(-951)), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-648)), function (_24_i0) {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-950)));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_25_i1) {
        return _dafny.MultiSet.fromElements(new BigNumber(-132), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), function (_26_i2) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })).length)), new BigNumber(-96));
      })));
    };
    static fm15(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),(!(false)) && (false));
    };
    static fm16(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(new BigNumber(-10), (new BigNumber((_dafny.Set.fromElements(new BigNumber(316))).length)).isLessThan(new BigNumber(683)), ((true) ? (_dafny.Seq.UnicodeFromString("w")) : (_dafny.Seq.UnicodeFromString("jpsbupx"))), function () {
  let _coll11 = new _dafny.Map();
  for (const _compr_11 of _dafny.IntegerRange(new BigNumber(671), new BigNumber(265))) {
    let _27_v0 = _compr_11;
    if (((new BigNumber(671)).isLessThanOrEqualTo(_27_v0)) && ((_27_v0).isLessThan(new BigNumber(265)))) {
      _coll11.push([_module.__default.safeModuloInt(_27_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-506)), function (_29_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length))),function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of _dafny.IntegerRange(new BigNumber(858), new BigNumber(123))) {
          let _28_v1 = _compr_12;
          if (((new BigNumber(858)).isLessThanOrEqualTo(_28_v1)) && ((_28_v1).isLessThan(new BigNumber(123)))) {
            _coll12.push([_module.__default.safeDivisionInt(_28_v1, new BigNumber(410)),!(false)]);
          }
        }
        return _coll12;
      }()]);
    }
  }
  return _coll11;
}());
    };
    static fm17(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(642)), function (_30_i0) {
        return new BigNumber(822);
      }), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(219), new BigNumber((_dafny.Seq.UnicodeFromString("agbh")).length), new BigNumber((_dafny.Seq.UnicodeFromString("xvgfxoaqc")).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-211), new BigNumber(157), new BigNumber(330))).length), new BigNumber(31)), _dafny.Seq.of(new BigNumber(801), new BigNumber(912))));
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("vvxhidyxj"))).Difference(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(146)), function (_31_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })));
    };
    static fm19(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(878))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe((_module.D6.create_DC21(true, new _dafny.CodePoint('g'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("rjujka")).length), new BigNumber((_dafny.Seq.UnicodeFromString("lxlnkch")).length))).dtor_cf34,new BigNumber(317)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xpd")).length))).length)))));
    };
    static fm20(p0, globalState) {
      return _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber(931), new BigNumber(846)), new BigNumber(-596));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = _dafny.Set.Empty;
      let _32_v0;
      _32_v0 = true;
      let _33_v1;
      _33_v1 = new _dafny.CodePoint('c'.codePointAt(0));
      let _34_v2;
      _34_v2 = _dafny.Set.fromElements(_33_v1);
      _32_v0 = (_module.__default.safeDivisionInt(new BigNumber(-677), new BigNumber((_34_v2).length))).isLessThan(p1);
      (globalState).f6 = p1;
      r0 = ((new BigNumber(-453)).plus(p1)).plus((_dafny.ZERO).minus((p1).multipliedBy((_dafny.ZERO).minus(p1))));
      let _35_v3;
      _35_v3 = _dafny.Seq.UnicodeFromString("mhpactow");
      let _36_v4;
      _36_v4 = _module.D0.create_DC0(p1, p3, _35_v3, _module.__default.fm11(p1, _32_v0, globalState));
      let _pat_let_tv0 = _33_v1;
      let _pat_let_tv1 = p3;
      let _pat_let_tv2 = _36_v4;
      let _pat_let_tv3 = p1;
      let _pat_let_tv4 = _33_v1;
      let _pat_let_tv5 = p1;
      let _pat_let_tv6 = globalState;
      let _source1 = function (_source2) {
        let _37___mcc_h9 = (_source2).cf0;
        let _38___mcc_h10 = (_source2).cf1;
        let _39___mcc_h11 = (_source2).cf2;
        let _40___mcc_h12 = (_source2).cf3;
        let _41_cf3 = _40___mcc_h12;
        let _42_cf2 = _39___mcc_h11;
        let _43_cf1 = _38___mcc_h10;
        let _44_cf0 = _37___mcc_h9;
        return _module.D1.create_DC2(_pat_let_tv0);
      }(function (_pat_let0_0) {
        return function (_45_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_46_dt__update_hcf1_h0) {
              return function (_pat_let2_0) {
                return function (_47_dt__update_hcf2_h0) {
                  return _module.D0.create_DC0((_45_dt__update__tmp_h0).dtor_cf0, _46_dt__update_hcf1_h0, _47_dt__update_hcf2_h0, (_45_dt__update__tmp_h0).dtor_cf3);
                }(_pat_let2_0);
              }(_module.__default.fm12(_pat_let_tv2, _pat_let_tv3, _pat_let_tv4, _pat_let_tv5, _pat_let_tv6));
            }(_pat_let1_0);
          }(_pat_let_tv1);
        }(_pat_let0_0);
      }(_36_v4));
      if (_source1.is_DC2) {
        let _48___mcc_h0 = (_source1).cf5;
        let _49_cf5 = _48___mcc_h0;
        let _50_v5;
        let _nw0 = Array((new BigNumber(9)).toNumber()).fill(false);
        _50_v5 = _nw0;
        let _rhs0 = p2;
        let _rhs1 = _32_v0;
        let _rhs2 = _dafny.Seq.Concat(_35_v3, _dafny.Seq.Concat(_35_v3, _dafny.Seq.UnicodeFromString("y")));
        let _rhs3 = _50_v5;
        let _rhs4 = _35_v3;
        _32_v0 = _rhs0;
        _32_v0 = _rhs1;
        _35_v3 = _rhs2;
        _50_v5 = _rhs3;
        _35_v3 = _rhs4;
        let _51_v6;
        _51_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_module.__default.fm0(p1, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_35_v3).length),new BigNumber((_dafny.Seq.update(_35_v3, _module.__default.safeIndex(p1, new BigNumber((_35_v3).length)), _49_cf5)).length)), globalState), p1),(p1).minus(p1));
        let _52_v7;
        _52_v7 = _dafny.Set.fromElements(_35_v3, _35_v3);
        _51_v6 = (_51_v6).update(p1, new BigNumber((_52_v7).length));
        let _index0 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_50_v5).length));
        (_50_v5)[_index0] = true;
        let _index1 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_50_v5).length));
        let _rhs5 = p1;
        let _rhs6 = _32_v0;
        let _rhs7 = p1;
        let _lhs0 = globalState;
        let _lhs1 = _50_v5;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_50_v5).length));
        let _lhs3 = globalState;
        _lhs0.f11 = _rhs5;
        _lhs1[_lhs2] = _rhs6;
        _lhs3.f16 = _rhs7;
        let _53_v8;
        let _init0 = ((_54_p1) => function (_55_i0) {
          return (_55_i0).plus(_54_p1);
        })(p1);
        let _nw1 = Array((new BigNumber(26)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _53_v8 = _nw1;
        let _index2 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_53_v8).length));
        (_53_v8)[_index2] = p1;
        let _index3 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_53_v8).length));
        let _index4 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_50_v5).length));
        let _rhs8 = _33_v1;
        let _rhs9 = p1;
        let _rhs10 = _32_v0;
        let _lhs4 = _53_v8;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_53_v8).length));
        let _lhs6 = _50_v5;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_50_v5).length));
        _33_v1 = _rhs8;
        _lhs4[_lhs5] = _rhs9;
        _lhs6[_lhs7] = _rhs10;
      } else if (_source1.is_DC3) {
        let _56___mcc_h1 = (_source1).cf6;
        let _57___mcc_h2 = (_source1).cf7;
        let _58___mcc_h3 = (_source1).cf8;
        let _59_cf8 = _58___mcc_h3;
        let _60_cf7 = _57___mcc_h2;
        let _61_cf6 = _56___mcc_h1;
        let _62_v9;
        let _nw2 = new _module.C0();
        _nw2.__ctor(_32_v0);
        _62_v9 = _nw2;
        let _63_v10;
        _63_v10 = _dafny.MultiSet.fromElements(_62_v9);
        let _64_v11;
        _64_v11 = _dafny.Map.Empty.slice().updateUnsafe(_62_v9,(((_63_v10).contains(_62_v9)) ? ((_63_v10).get(_62_v9)) : (p1)));
        let _65_v12;
        _65_v12 = _dafny.Set.fromElements(_64_v11, (_64_v11).Merge(_64_v11), _64_v11, (_64_v11).Merge(_64_v11));
        _65_v12 = _65_v12;
        let _66_v14;
        _66_v14 = _dafny.Seq.of(_60_cf7);
        let _67_v15;
        _67_v15 = _dafny.Seq.of(_dafny.Seq.of((_62_v9).f17));
        let _68_v16;
        _68_v16 = _dafny.Map.Empty.slice().updateUnsafe(_60_cf7,p3);
        let _69_v17;
        _69_v17 = _dafny.Set.fromElements(new BigNumber(191));
        let _70_v18;
        _70_v18 = _dafny.Map.Empty.slice().updateUnsafe((((_68_v16).contains(new BigNumber(898))) ? ((_68_v16).get(new BigNumber(898))) : (p2)),_69_v17);
        let _71_v19;
        _71_v19 = _dafny.Set.fromElements(_60_cf7, (_dafny.ZERO).minus(new BigNumber(((((_70_v18).contains(p2)) ? ((_70_v18).get(p2)) : (_69_v17))).length)), p1, _60_cf7);
        let _72_v20;
        let _nw3 = Array((new BigNumber(28)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw3[(_dafny.ONE).toNumber()] = new BigNumber(763);
        _nw3[(new BigNumber(2)).toNumber()] = (p1).minus(_module.__default.fm0(p1, _61_cf6, globalState));
        _nw3[(new BigNumber(3)).toNumber()] = p1;
        _nw3[(new BigNumber(4)).toNumber()] = new BigNumber(947);
        _nw3[(new BigNumber(5)).toNumber()] = p1;
        _nw3[(new BigNumber(6)).toNumber()] = (_60_cf7).minus(new BigNumber((function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of (_61_cf6).Keys.Elements) {
            let _73_v13 = _compr_13;
            if ((_61_cf6).contains(_73_v13)) {
              _coll13.push([(_73_v13).multipliedBy(_60_cf7),p1]);
            }
          }
          return _coll13;
        }()).length));
        _nw3[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_35_v3, _module.__default.safeIndex(p1, new BigNumber((_35_v3).length)), _33_v1), _35_v3)).length);
        _nw3[(new BigNumber(8)).toNumber()] = p1;
        _nw3[(new BigNumber(9)).toNumber()] = (((_63_v10).contains(_62_v9)) ? ((_63_v10).get(_62_v9)) : (_60_cf7));
        _nw3[(new BigNumber(10)).toNumber()] = p1;
        _nw3[(new BigNumber(11)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(12)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(13)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(14)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(15)).toNumber()] = p1;
        _nw3[(new BigNumber(16)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(17)).toNumber()] = (_66_v14)[_module.__default.safeIndex(p1, new BigNumber((_66_v14).length))];
        _nw3[(new BigNumber(18)).toNumber()] = (_66_v14)[_module.__default.safeIndex(_60_cf7, new BigNumber((_66_v14).length))];
        _nw3[(new BigNumber(19)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(20)).toNumber()] = new BigNumber(((_67_v15)[_module.__default.safeIndex(p1, new BigNumber((_67_v15).length))]).length);
        _nw3[(new BigNumber(21)).toNumber()] = _module.__default.safeModuloInt(_60_cf7, new BigNumber(-884));
        _nw3[(new BigNumber(22)).toNumber()] = new BigNumber((_71_v19).length);
        _nw3[(new BigNumber(23)).toNumber()] = p1;
        _nw3[(new BigNumber(24)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(25)).toNumber()] = _60_cf7;
        _nw3[(new BigNumber(26)).toNumber()] = p1;
        _nw3[(new BigNumber(27)).toNumber()] = new BigNumber((_35_v3).length);
        _72_v20 = _nw3;
        let _index5 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_72_v20).length));
        (_72_v20)[_index5] = ((!(p3)) ? (new BigNumber(183)) : (p1));
        let _index6 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_72_v20).length));
        let _rhs11 = _module.__default.fm0(p1, _61_cf6, globalState);
        let _rhs12 = new BigNumber(909);
        let _rhs13 = true;
        let _rhs14 = true;
        let _rhs15 = _32_v0;
        let _lhs8 = _72_v20;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_72_v20).length));
        let _lhs10 = globalState;
        _lhs8[_lhs9] = _rhs11;
        _lhs10.f16 = _rhs12;
        _32_v0 = _rhs13;
        _32_v0 = _rhs14;
        _32_v0 = _rhs15;
        let _74_v21;
        _74_v21 = _dafny.MultiSet.fromElements(p1);
        let _75_v22;
        _75_v22 = _dafny.Set.fromElements((_62_v9).f17, (_62_v9).f17);
        let _76_v23;
        _76_v23 = _dafny.Map.Empty.slice().updateUnsafe((_74_v21).Difference(_74_v21),_75_v22);
        _76_v23 = _76_v23;
        let _77_v24;
        _77_v24 = _dafny.Seq.of(_72_v20, _72_v20, _72_v20, _72_v20);
        let _rhs16 = _66_v14;
        let _rhs17 = _62_v9;
        let _rhs18 = _77_v24;
        _66_v14 = _rhs16;
        _62_v9 = _rhs17;
        _77_v24 = _rhs18;
      } else if (_source1.is_DC4) {
        let _78___mcc_h4 = (_source1).cf9;
        let _79___mcc_h5 = (_source1).cf10;
        let _80___mcc_h6 = (_source1).cf11;
        let _81_cf11 = _80___mcc_h6;
        let _82_cf10 = _79___mcc_h5;
        let _83_cf9 = _78___mcc_h4;
        let _84_v25;
        _84_v25 = _dafny.Seq.of(_82_cf10, p2, p2, p3, p2);
        let _85_v26;
        _85_v26 = _module.D1.create_DC1(_84_v25);
        _85_v26 = _module.D1.create_DC1(_dafny.Seq.of(p2, p2));
        let _source3 = _module.D1.create_DC2(_33_v1);
        if (_source3.is_DC2) {
          let _86___mcc_h13 = (_source3).cf5;
          let _87_cf5 = _86___mcc_h13;
          let _88_v27;
          let _nw4 = new _module.C0();
          _nw4.__ctor(p3);
          _88_v27 = _nw4;
          let _89_v28;
          _89_v28 = _dafny.Seq.of(_88_v27);
          let _90_v29;
          let _nw5 = Array((new BigNumber(7)).toNumber()).fill(false);
          _90_v29 = _nw5;
          let _91_v30;
          _91_v30 = _dafny.Map.Empty.slice().updateUnsafe(_83_cf9,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_90_v29,_87_cf5)).length));
          let _92_v31;
          _92_v31 = _dafny.Map.Empty.slice().updateUnsafe(_82_cf10,_module.D1.create_DC4(_83_cf9, _module.__default.fm4(new BigNumber((_89_v28).length), (_88_v27).f17, globalState), _dafny.Map.Empty.slice().updateUnsafe(_83_cf9,_dafny.MultiSet.fromElements(_83_cf9, _module.__default.fm0(new BigNumber((_35_v3).length), _91_v30, globalState)))));
          let _93_v33;
          _93_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_89_v28).length),_32_v0);
          let _94_v34;
          _94_v34 = _module.D1.create_DC4(_83_cf9, (_module.__default.fm13(globalState)).dtor_cf10, function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of (_93_v33).Keys.Elements) {
    let _95_v32 = _compr_14;
    if ((_93_v33).contains(_95_v32)) {
      _coll14.push([(_95_v32).plus(p1),_dafny.MultiSet.fromElements(_83_cf9, p1)]);
    }
  }
  return _coll14;
}());
          _82_cf10 = !((_92_v31).Merge(_92_v31)).equals(_dafny.Map.Empty.slice().updateUnsafe(_32_v0,_94_v34));
          let _96_v37;
          let _init1 = ((_97_cf5, _98_v3, _99_p1, _100_v1, _101_cf9) => function (_102_i1) {
            return (_dafny.Set.fromElements(function () {
              let _coll15 = new _dafny.Map();
              for (const _compr_15 of (_dafny.Seq.of((_98_v3)[_module.__default.safeIndex(_99_p1, new BigNumber((_98_v3).length))], _100_v1)).Elements) {
                let _103_v35 = _compr_15;
                if (_dafny.Seq.contains(_dafny.Seq.of((_98_v3)[_module.__default.safeIndex(_99_p1, new BigNumber((_98_v3).length))], _100_v1), _103_v35)) {
                  _coll15.push([_103_v35,_99_p1]);
                }
              }
              return _coll15;
            }(), function () {
              let _coll16 = new _dafny.Map();
              for (const _compr_16 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_100_v1))).Elements) {
                let _104_v36 = _compr_16;
                if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_100_v1))).contains(_104_v36)) {
                  _coll16.push([_104_v36,(_dafny.ZERO).minus(_99_p1)]);
                }
              }
              return _coll16;
            }())).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_97_cf5,_101_cf9)));
          })(_87_cf5, _35_v3, p1, _33_v1, _83_cf9);
          let _nw6 = Array((new BigNumber(18)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
            _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _96_v37 = _nw6;
          let _105_v39;
          _105_v39 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)));
          let _106_v40;
          _106_v40 = _dafny.Map.Empty.slice().updateUnsafe(_87_cf5,_module.__default.fm0(p1, _91_v30, globalState));
          let _107_v41;
          _107_v41 = _dafny.MultiSet.fromElements(function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_105_v39).Elements) {
              let _108_v38 = _compr_17;
              if ((_105_v39).contains(_108_v38)) {
                _coll17.push([_108_v38,new BigNumber((_dafny.MultiSet.fromElements(_82_cf10)).cardinality())]);
              }
            }
            return _coll17;
          }(), _106_v40);
          let _109_v43;
          _109_v43 = _dafny.Set.fromElements(_106_v40);
          let _index7 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_96_v37).length));
          (_96_v37)[_index7] = (function () {
            let _coll18 = new _dafny.Set();
            for (const _compr_18 of (_107_v41).Elements) {
              let _110_v42 = _compr_18;
              if ((_107_v41).contains(_110_v42)) {
                _coll18.add(_110_v42);
              }
            }
            return _coll18;
          }()).Difference(_109_v43);
          let _index8 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_96_v37).length));
          (_96_v37)[_index8] = (_109_v43).Intersect(_109_v43);
          _91_v30 = (_91_v30).Merge(_91_v30);
        } else if (_source3.is_DC3) {
          let _111___mcc_h14 = (_source3).cf6;
          let _112___mcc_h15 = (_source3).cf7;
          let _113___mcc_h16 = (_source3).cf8;
          let _114_cf8 = _113___mcc_h16;
          let _115_cf7 = _112___mcc_h15;
          let _116_cf6 = _111___mcc_h14;
          let _117_v44;
          _117_v44 = _dafny.Set.fromElements(new BigNumber((_116_cf6).length));
          let _118_v45;
          let _init2 = ((_119_v1) => function (_120_i3) {
            return _119_v1;
          })(_33_v1);
          let _nw7 = Array((new BigNumber(16)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw7.length); _i0_2++) {
            _nw7[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _118_v45 = _nw7;
          _82_cf10 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), ((_121_v3) => function (_122_i2) {
            return new BigNumber((_121_v3).length);
          })(_35_v3))).length)).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_117_v44).length), new BigNumber((_dafny.MultiSet.fromElements(_118_v45)).cardinality())));
          let _123_v46;
          let _nw8 = Array((new BigNumber(8)).toNumber()).fill(false);
          _123_v46 = _nw8;
          let _124_v47;
          let _nw9 = new _module.C0();
          _nw9.__ctor(_82_cf10);
          _124_v47 = _nw9;
          let _125_v48;
          _125_v48 = _dafny.MultiSet.fromElements(p2);
          let _126_v49;
          _126_v49 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_124_v47).f17);
          let _127_v50;
          _127_v50 = _dafny.Map.Empty.slice().updateUnsafe(_83_cf9,_126_v49);
          let _128_v51;
          _128_v51 = _dafny.Map.Empty.slice().updateUnsafe(_124_v47,_module.__default.fm4(_83_cf9, (_module.D0.create_DC0((((_125_v48).contains(_module.__default.fm4(new BigNumber(-968), p3, globalState))) ? ((_125_v48).get(_module.__default.fm4(new BigNumber(-968), p3, globalState))) : (_83_cf9)), _module.__default.fm4(_83_cf9, !(_32_v0), globalState), _dafny.Seq.UnicodeFromString("suaxnms"), _127_v50)).dtor_cf1, globalState));
          let _index9 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_123_v46).length));
          (_123_v46)[_index9] = (((_128_v51).contains(_124_v47)) ? ((_128_v51).get(_124_v47)) : (_32_v0));
          let _index10 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_123_v46).length));
          (_123_v46)[_index10] = (_124_v47).f17;
          let _index11 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_123_v46).length));
          (_123_v46)[_index11] = (_117_v44).IsSubsetOf(function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(509), new BigNumber(976))) {
              let _129_v52 = _compr_19;
              if (((new BigNumber(509)).isLessThanOrEqualTo(_129_v52)) && ((_129_v52).isLessThan(new BigNumber(976)))) {
                _coll19.add((_129_v52).plus(_83_cf9));
              }
            }
            return _coll19;
          }());
          let _130_v53;
          _130_v53 = _dafny.Seq.of(_123_v46, _123_v46, _123_v46, _123_v46, _123_v46);
          _130_v53 = _130_v53;
        } else if (_source3.is_DC4) {
          let _131___mcc_h17 = (_source3).cf9;
          let _132___mcc_h18 = (_source3).cf10;
          let _133___mcc_h19 = (_source3).cf11;
          let _134_cf11 = _133___mcc_h19;
          let _135_cf10 = _132___mcc_h18;
          let _136_cf9 = _131___mcc_h17;
          _33_v1 = _33_v1;
          let _137_v54;
          _137_v54 = _dafny.MultiSet.fromElements(_32_v0);
          (globalState).f6 = (((_137_v54).contains(_32_v0)) ? ((_137_v54).get(_32_v0)) : ((new BigNumber((_dafny.Seq.of(!(p3))).length)).plus(_136_cf9)));
          let _138_v55;
          let _nw10 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _138_v55 = _nw10;
          let _139_v56;
          _139_v56 = _dafny.Map.Empty.slice().updateUnsafe(_135_cf10,_138_v55);
          let _140_v57;
          _140_v57 = _dafny.Seq.of(_136_cf9);
          let _141_v58;
          let _nw11 = Array((new BigNumber(3)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = p3;
          _nw11[(_dafny.ONE).toNumber()] = p3;
          _nw11[(new BigNumber(2)).toNumber()] = _135_cf10;
          _141_v58 = _nw11;
          let _142_v59;
          _142_v59 = _dafny.Map.Empty.slice().updateUnsafe((p1).plus((_140_v57)[_module.__default.safeIndex(p1, new BigNumber((_140_v57).length))]),_141_v58);
          let _143_v60;
          _143_v60 = _dafny.Seq.of(_139_v56, (_139_v56).Merge(_139_v56));
          let _144_v61;
          _144_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(823)), ((_145_v1) => function (_146_i4) {
            return _145_v1;
          })(_33_v1)),p1);
          let _147_v62;
          _147_v62 = _dafny.Set.fromElements(_136_cf9);
          let _148_v63;
          let _nw12 = new _module.C0();
          _nw12.__ctor(_module.__default.fm4(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_32_v0,p3)).length), _136_cf9)).length), p3, globalState));
          _148_v63 = _nw12;
          let _149_v64;
          _149_v64 = _dafny.MultiSet.fromElements(_148_v63, _148_v63, _148_v63);
          let _150_v65;
          _150_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_147_v62).length),(_dafny.ZERO).minus(new BigNumber((_149_v64).cardinality())));
          let _151_v66;
          _151_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,_150_v65);
          let _rhs19 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_136_cf9, _136_cf9)));
          let _rhs20 = (_143_v60)[_module.__default.safeIndex((_136_cf9).minus(p1), new BigNumber((_143_v60).length))];
          let _rhs21 = _module.__default.fm0((p1).multipliedBy(new BigNumber((_144_v61).length)), (((_151_v66).contains(true)) ? ((_151_v66).get(true)) : (_150_v65)), globalState);
          let _rhs22 = _142_v59;
          let _lhs11 = globalState;
          let _lhs12 = globalState;
          _lhs11.f11 = _rhs19;
          _139_v56 = _rhs20;
          _lhs12.f11 = _rhs21;
          _142_v59 = _rhs22;
          let _152_v67;
          _152_v67 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), function (_153_i5) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-300)), ((_154_v1) => function (_155_i6) {
            return _154_v1;
          })(_33_v1)), _35_v3));
          _35_v3 = _dafny.Seq.update((_152_v67)[_module.__default.safeIndex(_module.__default.safeModuloInt(p1, p1), new BigNumber((_152_v67).length))], _module.__default.safeIndex(_136_cf9, new BigNumber(((_152_v67)[_module.__default.safeIndex(_module.__default.safeModuloInt(p1, p1), new BigNumber((_152_v67).length))]).length)), new _dafny.CodePoint('j'.codePointAt(0)));
        } else if (_source3.is_DC1) {
          let _156___mcc_h20 = (_source3).cf4;
          let _157_cf4 = _156___mcc_h20;
          (globalState).f6 = _83_cf9;
          let _158_v69;
          let _nw13 = new _module.C0();
          _nw13.__ctor(!(_82_cf10));
          _158_v69 = _nw13;
          let _159_v70;
          _159_v70 = _dafny.Map.Empty.slice().updateUnsafe(_158_v69,_83_cf9);
          (globalState).f6 = (new BigNumber((function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (_35_v3).Elements) {
              let _160_v68 = _compr_20;
              if (_dafny.Seq.contains(_35_v3, _160_v68)) {
                _coll20.push([_160_v68,new BigNumber(860)]);
              }
            }
            return _coll20;
          }()).length)).minus((new BigNumber((_159_v70).length)).plus(p1));
          _158_v69 = _158_v69;
          let _161_v71;
          _161_v71 = _module.D1.create_DC2(_module.__default.fm3(_82_cf10, globalState));
          _161_v71 = function (_pat_let3_0) {
            return function (_162_dt__update__tmp_h1) {
              return function (_pat_let4_0) {
                return function (_163_dt__update_hcf5_h0) {
                  return _module.D1.create_DC2(_163_dt__update_hcf5_h0);
                }(_pat_let4_0);
              }(new _dafny.CodePoint('q'.codePointAt(0)));
            }(_pat_let3_0);
          }(_161_v71);
        } else {
          let _164___mcc_h21 = (_source3).cf12;
          let _165_cf12 = _164___mcc_h21;
          let _166_v72;
          _166_v72 = _module.D1.create_DC4(p1, true, _81_cf11);
          let _167_v73;
          _167_v73 = _dafny.Map.Empty.slice().updateUnsafe((_166_v72).dtor_cf10,_dafny.Seq.of(_83_cf9));
          let _168_v74;
          let _nw14 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _168_v74 = _nw14;
          let _169_v75;
          let _nw15 = new _module.C0();
          _nw15.__ctor(p2);
          _169_v75 = _nw15;
          let _170_v76;
          let _nw16 = Array((new BigNumber(8)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _169_v75;
          _nw16[(_dafny.ONE).toNumber()] = _169_v75;
          _nw16[(new BigNumber(2)).toNumber()] = _169_v75;
          _nw16[(new BigNumber(3)).toNumber()] = _169_v75;
          _nw16[(new BigNumber(4)).toNumber()] = _169_v75;
          _nw16[(new BigNumber(5)).toNumber()] = _169_v75;
          _nw16[(new BigNumber(6)).toNumber()] = _169_v75;
          _nw16[(new BigNumber(7)).toNumber()] = _169_v75;
          _170_v76 = _nw16;
          let _171_v77;
          _171_v77 = _dafny.Map.Empty.slice().updateUnsafe(_168_v74,_170_v76);
          let _172_v78;
          _172_v78 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_83_cf9)).cardinality()));
          let _173_v79;
          _173_v79 = _dafny.Map.Empty.slice().updateUnsafe(_83_cf9,new BigNumber((_172_v78).length));
          let _rhs23 = _83_cf9;
          let _rhs24 = new BigNumber(((_171_v77).Merge(_171_v77)).length);
          let _rhs25 = (_167_v73).Merge(_167_v73);
          let _rhs26 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(p1, _83_cf9), new BigNumber((_173_v79).length));
          let _rhs27 = _33_v1;
          let _lhs13 = globalState;
          let _lhs14 = globalState;
          r0 = _rhs23;
          _lhs13.f16 = _rhs24;
          _167_v73 = _rhs25;
          _lhs14.f6 = _rhs26;
          _33_v1 = _rhs27;
          let _174_v80;
          let _nw17 = Array((new BigNumber(28)).toNumber()).fill(false);
          _174_v80 = _nw17;
          let _index12 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_168_v74).length));
          (_168_v74)[_index12] = (((_173_v79).contains(p1)) ? ((_173_v79).get(p1)) : ((_36_v4).dtor_cf0));
          let _175_v81;
          _175_v81 = _dafny.Seq.of(_174_v80, _174_v80);
          let _index13 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_168_v74).length));
          let _rhs28 = _174_v80;
          let _rhs29 = _module.__default.safeDivisionInt(_83_cf9, p1);
          let _rhs30 = new BigNumber((_175_v81).length);
          let _rhs31 = _module.__default.safeDivisionInt(new BigNumber(258), p1);
          let _lhs15 = _168_v74;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_168_v74).length));
          let _lhs17 = globalState;
          _174_v80 = _rhs28;
          _83_cf9 = _rhs29;
          _lhs15[_lhs16] = _rhs30;
          _lhs17.f6 = _rhs31;
          let _176_v82;
          _176_v82 = _dafny.MultiSet.fromElements(p3, (((_169_v75).f17) ? (_82_cf10) : (false)), _32_v0);
          _176_v82 = _176_v82;
          let _index14 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_170_v76).length));
          (_170_v76)[_index14] = _169_v75;
          let _177_v83;
          let _nw18 = Array((new BigNumber(18)).toNumber()).fill([]);
          _177_v83 = _nw18;
          let _index15 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_177_v83).length));
          (_177_v83)[_index15] = (_175_v81)[_module.__default.safeIndex(new BigNumber(-373), new BigNumber((_175_v81).length))];
          let _index16 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_174_v80).length));
          (_174_v80)[_index16] = p3;
          let _index17 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_170_v76).length));
          let _index18 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_177_v83).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_174_v80).length));
          let _rhs32 = _169_v75;
          let _rhs33 = ((_168_v74)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_168_v74).length))]).isLessThan(_module.__default.safeModuloInt(p1, _module.__default.fm0((_168_v74)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_168_v74).length))], _173_v79, globalState)));
          let _rhs34 = _174_v80;
          let _rhs35 = p2;
          let _rhs36 = p2;
          let _lhs18 = _170_v76;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_170_v76).length));
          let _lhs20 = _177_v83;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_177_v83).length));
          let _lhs22 = _174_v80;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_174_v80).length));
          _lhs18[_lhs19] = _rhs32;
          _82_cf10 = _rhs33;
          _lhs20[_lhs21] = _rhs34;
          _lhs22[_lhs23] = _rhs35;
          _32_v0 = _rhs36;
        }
        let _178_v84;
        let _nw19 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _178_v84 = _nw19;
        _178_v84 = _178_v84;
        let _179_v85;
        _179_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_83_cf9);
        let _180_v86;
        _180_v86 = _dafny.MultiSet.fromElements((_84_v25)[_module.__default.safeIndex(p1, new BigNumber((_84_v25).length))]);
        let _pat_let_tv7 = _180_v86;
        let _source4 = function (_pat_let5_0) {
          return function (_181_dt__update__tmp_h2) {
            return function (_pat_let6_0) {
              return function (_182_dt__update_hcf7_h0) {
                return _module.D1.create_DC3((_181_dt__update__tmp_h2).dtor_cf6, _182_dt__update_hcf7_h0, (_181_dt__update__tmp_h2).dtor_cf8);
              }(_pat_let6_0);
            }(new BigNumber((_pat_let_tv7).cardinality()));
          }(_pat_let5_0);
        }(_module.D1.create_DC3(_179_v85, p1, _36_v4));
        if (_source4.is_DC2) {
          let _183___mcc_h22 = (_source4).cf5;
          let _184_cf5 = _183___mcc_h22;
          let _rhs37 = (_82_cf10) || ((!(_32_v0)) || (_32_v0));
          let _rhs38 = _dafny.Seq.UnicodeFromString("pufjq");
          _82_cf10 = _rhs37;
          _35_v3 = _rhs38;
          let _185_v87;
          let _nw20 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _185_v87 = _nw20;
          let _186_v88;
          let _nw21 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _186_v88 = _nw21;
          let _index20 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_185_v87).length));
          (_185_v87)[_index20] = _dafny.Set.fromElements(_186_v88, _186_v88, _186_v88);
          let _187_v89;
          _187_v89 = _dafny.Set.fromElements(_186_v88, _186_v88, _186_v88, _186_v88);
          let _188_v90;
          _188_v90 = _dafny.Seq.of(_187_v89, _187_v89, _dafny.Set.fromElements(_186_v88, _186_v88), _187_v89);
          let _index21 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_185_v87).length));
          (_185_v87)[_index21] = ((_187_v89).Intersect((_188_v90)[_module.__default.safeIndex(p1, new BigNumber((_188_v90).length))])).Intersect(_187_v89);
          _82_cf10 = (_180_v86).IsProperSubsetOf((_dafny.MultiSet.FromArray(_84_v25)).update(p3, _module.__default.abs(_83_cf9)));
          let _189_v91;
          let _nw22 = new _module.C0();
          _nw22.__ctor(p3);
          _189_v91 = _nw22;
        } else if (_source4.is_DC3) {
          let _190___mcc_h23 = (_source4).cf6;
          let _191___mcc_h24 = (_source4).cf7;
          let _192___mcc_h25 = (_source4).cf8;
          let _193_cf8 = _192___mcc_h25;
          let _194_cf7 = _191___mcc_h24;
          let _195_cf6 = _190___mcc_h23;
          _33_v1 = _33_v1;
          let _196_v92;
          let _nw23 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _196_v92 = _nw23;
          let _197_v93;
          _197_v93 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.update(_35_v3, _module.__default.safeIndex(new BigNumber(7), new BigNumber((_35_v3).length)), new _dafny.CodePoint('t'.codePointAt(0))));
          let _198_v94;
          _198_v94 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,p1);
          let _index22 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_196_v92).length));
          (_196_v92)[_index22] = new BigNumber(((((_197_v93).contains((((_198_v94).contains(_82_cf10)) ? ((_198_v94).get(_82_cf10)) : (_194_cf7)))) ? ((_197_v93).get((((_198_v94).contains(_82_cf10)) ? ((_198_v94).get(_82_cf10)) : (_194_cf7)))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), ((_199_v1) => function (_200_i7) {
            return _199_v1;
          })(_33_v1))))).length);
          let _201_v95;
          let _nw24 = Array((new BigNumber(18)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _82_cf10;
          _nw24[(_dafny.ONE).toNumber()] = _82_cf10;
          _nw24[(new BigNumber(2)).toNumber()] = p2;
          _nw24[(new BigNumber(3)).toNumber()] = p2;
          _nw24[(new BigNumber(4)).toNumber()] = p3;
          _nw24[(new BigNumber(5)).toNumber()] = p2;
          _nw24[(new BigNumber(6)).toNumber()] = (new BigNumber(-500)).isLessThanOrEqualTo(p1);
          _nw24[(new BigNumber(7)).toNumber()] = _82_cf10;
          _nw24[(new BigNumber(8)).toNumber()] = !((_84_v25)[_module.__default.safeIndex(_194_cf7, new BigNumber((_84_v25).length))]);
          _nw24[(new BigNumber(9)).toNumber()] = _32_v0;
          _nw24[(new BigNumber(10)).toNumber()] = ((_module.__default.fm6(p2, p3, p2, globalState)).update(p2, _module.__default.abs(_83_cf9))).IsDisjointFrom(_180_v86);
          _nw24[(new BigNumber(11)).toNumber()] = p3;
          _nw24[(new BigNumber(12)).toNumber()] = p3;
          _nw24[(new BigNumber(13)).toNumber()] = false;
          _nw24[(new BigNumber(14)).toNumber()] = (p1).isEqualTo(_194_cf7);
          _nw24[(new BigNumber(15)).toNumber()] = p2;
          _nw24[(new BigNumber(16)).toNumber()] = p2;
          _nw24[(new BigNumber(17)).toNumber()] = true;
          _201_v95 = _nw24;
          let _index23 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_201_v95).length));
          (_201_v95)[_index23] = _32_v0;
          let _index24 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_196_v92).length));
          let _index25 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_201_v95).length));
          let _rhs39 = (new BigNumber(345)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("udu"), _35_v3), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("udu"), _35_v3)).length)), _33_v1)).length));
          let _rhs40 = p3;
          let _rhs41 = _module.__default.safeModuloInt(_194_cf7, _83_cf9);
          let _rhs42 = _82_cf10;
          let _rhs43 = (_194_cf7).minus(p1);
          let _lhs24 = _196_v92;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_196_v92).length));
          let _lhs26 = _201_v95;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_201_v95).length));
          let _lhs28 = globalState;
          _32_v0 = _rhs39;
          _82_cf10 = _rhs40;
          _lhs24[_lhs25] = _rhs41;
          _lhs26[_lhs27] = _rhs42;
          _lhs28.f16 = _rhs43;
          let _202_v96;
          _202_v96 = _module.D2.create_DC6(_196_v92);
          let _203_v97;
          _203_v97 = _dafny.Seq.of((_202_v96).dtor_cf13, _196_v92, _196_v92);
          let _204_v98;
          let _nw25 = Array((new BigNumber(19)).toNumber());
          _204_v98 = _nw25;
          let _205_v99;
          let _nw26 = new _module.C0();
          _nw26.__ctor(_82_cf10);
          _205_v99 = _nw26;
          let _index26 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_204_v98).length));
          (_204_v98)[_index26] = _205_v99;
          let _index27 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_201_v95).length));
          let _index28 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_204_v98).length));
          let _rhs44 = (_84_v25)[_module.__default.safeIndex(_83_cf9, new BigNumber((_84_v25).length))];
          let _rhs45 = _dafny.Seq.update((((p1).isLessThan(_83_cf9)) ? (_dafny.Seq.Concat(_203_v97, _203_v97)) : (_203_v97)), _module.__default.safeIndex((_196_v92)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_196_v92).length))], new BigNumber(((((p1).isLessThan(_83_cf9)) ? (_dafny.Seq.Concat(_203_v97, _203_v97)) : (_203_v97))).length)), _196_v92);
          let _rhs46 = _205_v99;
          let _rhs47 = _201_v95;
          let _lhs29 = _201_v95;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_201_v95).length));
          let _lhs31 = _204_v98;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_204_v98).length));
          _lhs29[_lhs30] = _rhs44;
          _203_v97 = _rhs45;
          _lhs31[_lhs32] = _rhs46;
          _201_v95 = _rhs47;
          let _index29 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_201_v95).length));
          (_201_v95)[_index29] = (_83_cf9).isLessThanOrEqualTo(p1);
        } else if (_source4.is_DC4) {
          let _206___mcc_h26 = (_source4).cf9;
          let _207___mcc_h27 = (_source4).cf10;
          let _208___mcc_h28 = (_source4).cf11;
          let _209_cf11 = _208___mcc_h28;
          let _210_cf10 = _207___mcc_h27;
          let _211_cf9 = _206___mcc_h26;
          let _212_v100;
          let _nw27 = new _module.C0();
          _nw27.__ctor(_210_cf10);
          _212_v100 = _nw27;
          let _213_v101;
          _213_v101 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,_210_cf10);
          (globalState).f6 = new BigNumber((_213_v101).length);
          let _214_v102;
          let _init3 = ((_215_p3) => function (_216_i8) {
            return _215_p3;
          })(p3);
          let _nw28 = Array((new BigNumber(2)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw28.length); _i0_3++) {
            _nw28[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _214_v102 = _nw28;
          let _index30 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_214_v102).length));
          (_214_v102)[_index30] = p2;
          let _index31 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_214_v102).length));
          (_214_v102)[_index31] = (((_213_v101).contains(!(_dafny.Seq.IsPrefixOf(_35_v3, _35_v3)))) ? ((_213_v101).get(!(_dafny.Seq.IsPrefixOf(_35_v3, _35_v3)))) : (_210_cf10));
          (globalState).f16 = new BigNumber((_35_v3).length);
        } else if (_source4.is_DC1) {
          let _217___mcc_h29 = (_source4).cf4;
          let _218_cf4 = _217___mcc_h29;
          let _219_v103;
          let _nw29 = new _module.C0();
          _nw29.__ctor(!(p2) || (!(!(p2))));
          _219_v103 = _nw29;
          let _220_v104;
          _220_v104 = _dafny.Set.fromElements((_219_v103).f17);
          let _221_v105;
          _221_v105 = _dafny.Map.Empty.slice().updateUnsafe(_83_cf9,(_84_v25)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_219_v103).f17,(_219_v103).f17)).length), new BigNumber((_84_v25).length))]);
          _32_v0 = _module.__default.fm4(_module.__default.safeDivisionInt(p1, new BigNumber((_220_v104).length)), (((_221_v105).contains(_83_cf9)) ? ((_221_v105).get(_83_cf9)) : (p2)), globalState);
          let _222_v106;
          let _nw30 = Array((new BigNumber(2)).toNumber()).fill(false);
          _222_v106 = _nw30;
          _222_v106 = _222_v106;
          let _223_v107;
          _223_v107 = _dafny.Seq.of(_219_v103);
          let _224_v108;
          _224_v108 = _dafny.Seq.of(p1, p1, _83_cf9, p1, p1);
          let _rhs48 = _219_v103;
          let _rhs49 = _219_v103;
          let _rhs50 = !(!_dafny.areEqual(_223_v107, _223_v107));
          let _rhs51 = (_223_v107)[_module.__default.safeIndex((p1).multipliedBy((_224_v108)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_224_v108).length))]), new BigNumber((_223_v107).length))];
          _219_v103 = _rhs48;
          _219_v103 = _rhs49;
          _32_v0 = _rhs50;
          _219_v103 = _rhs51;
        } else {
          let _225___mcc_h30 = (_source4).cf12;
          let _226_cf12 = _225___mcc_h30;
          (globalState).f6 = p1;
          let _227_v109;
          let _nw31 = new _module.C0();
          _nw31.__ctor(p3);
          _227_v109 = _nw31;
          let _228_v110;
          _228_v110 = _dafny.Map.Empty.slice().updateUnsafe(_83_cf9,_227_v109);
          _228_v110 = (_228_v110).update((_dafny.ZERO).minus(_83_cf9), _227_v109);
          let _229_v111;
          let _nw32 = new _module.C0();
          _nw32.__ctor(p2);
          _229_v111 = _nw32;
          let _rhs52 = new BigNumber(-453);
          let _rhs53 = (((_dafny.ZERO).minus(p1)).minus(p1)).multipliedBy(new BigNumber((_module.__default.fm14(globalState)).length));
          let _lhs33 = globalState;
          let _lhs34 = globalState;
          _lhs33.f11 = _rhs52;
          _lhs34.f6 = _rhs53;
        }
      } else if (_source1.is_DC1) {
        let _230___mcc_h7 = (_source1).cf4;
        let _231_cf4 = _230___mcc_h7;
        let _232_v112;
        let _nw33 = new _module.C0();
        _nw33.__ctor(p2);
        _232_v112 = _nw33;
        _232_v112 = _232_v112;
        let _233_v113;
        _233_v113 = _dafny.Map.Empty.slice().updateUnsafe(_35_v3,_232_v112);
        _232_v112 = (((_233_v113).contains(((!(!(p2))) ? (_35_v3) : (_35_v3)))) ? ((_233_v113).get(((!(!(p2))) ? (_35_v3) : (_35_v3)))) : (((false) ? (_232_v112) : (_232_v112))));
        r0 = p1;
        let _234_v114;
        _234_v114 = _dafny.Map.Empty.slice().updateUnsafe((_232_v112).f17,p1);
        let _235_v115;
        _235_v115 = _dafny.Seq.of(_232_v112);
        let _236_v116;
        _236_v116 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _237_v117;
        _237_v117 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_235_v115).length), _236_v116, globalState),p1);
        let _238_v118;
        _238_v118 = _dafny.Map.Empty.slice().updateUnsafe(p1,_32_v0);
        let _239_v119;
        _239_v119 = _module.D1.create_DC4(_module.__default.safeDivisionInt(p1, p1), (_231_cf4)[_module.__default.safeIndex(p1, new BigNumber((_231_cf4).length))], _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.MultiSet.fromElements((((_234_v114).contains((_231_cf4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber(446), (((_236_v116).contains(p1)) ? ((_236_v116).get(p1)) : (p1)))).length), new BigNumber((_231_cf4).length))])) ? ((_234_v114).get((_231_cf4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber(446), (((_236_v116).contains(p1)) ? ((_236_v116).get(p1)) : (p1)))).length), new BigNumber((_231_cf4).length))])) : (new BigNumber((_238_v118).length))))));
        let _source5 = _239_v119;
        if (_source5.is_DC2) {
          let _240___mcc_h31 = (_source5).cf5;
          let _241_cf5 = _240___mcc_h31;
          let _242_v120;
          _242_v120 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(124)), ((_243_cf5) => function (_244_i9) {
            return _243_cf5;
          })(_241_cf5)));
          let _245_v122;
          _245_v122 = _dafny.Set.fromElements(!((_232_v112).f17), false, p3, (_232_v112).f17, (_module.__default.fm13(globalState)).dtor_cf10);
          let _246_v124;
          _246_v124 = _dafny.Seq.of(_238_v118, function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(92), new BigNumber(86))) {
              let _247_v123 = _compr_21;
              if (((new BigNumber(92)).isLessThanOrEqualTo(_247_v123)) && ((_247_v123).isLessThan(new BigNumber(86)))) {
                _coll21.push([(_247_v123).multipliedBy(new BigNumber(2)),p2]);
              }
            }
            return _coll21;
          }());
          let _248_v125;
          let _nw34 = Array((new BigNumber(15)).toNumber());
          _nw34[(_dafny.ZERO).toNumber()] = p2;
          _nw34[(_dafny.ONE).toNumber()] = p3;
          _nw34[(new BigNumber(2)).toNumber()] = !(p3);
          _nw34[(new BigNumber(3)).toNumber()] = (_232_v112).f17;
          _nw34[(new BigNumber(4)).toNumber()] = !(_32_v0);
          _nw34[(new BigNumber(5)).toNumber()] = (function () {
            let _coll22 = new _dafny.Set();
            for (const _compr_22 of (_242_v120).Elements) {
              let _249_v121 = _compr_22;
              if ((_242_v120).contains(_249_v121)) {
                _coll22.add(_249_v121);
              }
            }
            return _coll22;
          }()).IsProperSubsetOf(_242_v120);
          _nw34[(new BigNumber(6)).toNumber()] = (_232_v112).f17;
          _nw34[(new BigNumber(7)).toNumber()] = p3;
          _nw34[(new BigNumber(8)).toNumber()] = (p1).isLessThan(new BigNumber((_245_v122).length));
          _nw34[(new BigNumber(9)).toNumber()] = p2;
          _nw34[(new BigNumber(10)).toNumber()] = p2;
          _nw34[(new BigNumber(11)).toNumber()] = p3;
          _nw34[(new BigNumber(12)).toNumber()] = p3;
          _nw34[(new BigNumber(13)).toNumber()] = (_232_v112).f17;
          _nw34[(new BigNumber(14)).toNumber()] = _dafny.Seq.contains(_246_v124, _dafny.Map.Empty.slice().updateUnsafe(p1,true));
          _248_v125 = _nw34;
          let _250_v126;
          _250_v126 = _dafny.MultiSet.fromElements(p1);
          let _251_v127;
          _251_v127 = _dafny.Map.Empty.slice().updateUnsafe(p1,_250_v126);
          let _index32 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_248_v125).length));
          (_248_v125)[_index32] = (((_232_v112).f17) ? ((_module.D1.create_DC4(p1, p2, _251_v127)).dtor_cf10) : ((_232_v112).f17));
          let _252_v128;
          _252_v128 = _dafny.Set.fromElements(p1);
          let _index33 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_248_v125).length));
          (_248_v125)[_index33] = ((_252_v128).IsProperSubsetOf(_252_v128)) && (p3);
          let _253_v129;
          let _nw35 = new _module.C0();
          _nw35.__ctor((p1).isLessThan(p1));
          _253_v129 = _nw35;
          let _254_v131;
          let _nw36 = Array((new BigNumber(27)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = p1;
          _nw36[(_dafny.ONE).toNumber()] = p1;
          _nw36[(new BigNumber(2)).toNumber()] = p1;
          _nw36[(new BigNumber(3)).toNumber()] = p1;
          _nw36[(new BigNumber(4)).toNumber()] = p1;
          _nw36[(new BigNumber(5)).toNumber()] = p1;
          _nw36[(new BigNumber(6)).toNumber()] = p1;
          _nw36[(new BigNumber(7)).toNumber()] = p1;
          _nw36[(new BigNumber(8)).toNumber()] = p1;
          _nw36[(new BigNumber(9)).toNumber()] = p1;
          _nw36[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw36[(new BigNumber(11)).toNumber()] = new BigNumber((function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_238_v118).Keys.Elements) {
              let _255_v130 = _compr_23;
              if ((_238_v118).contains(_255_v130)) {
                _coll23.push([(_255_v130).plus(p1),(_232_v112).f17]);
              }
            }
            return _coll23;
          }()).length);
          _nw36[(new BigNumber(12)).toNumber()] = p1;
          _nw36[(new BigNumber(13)).toNumber()] = new BigNumber((_35_v3).length);
          _nw36[(new BigNumber(14)).toNumber()] = p1;
          _nw36[(new BigNumber(15)).toNumber()] = p1;
          _nw36[(new BigNumber(16)).toNumber()] = p1;
          _nw36[(new BigNumber(17)).toNumber()] = p1;
          _nw36[(new BigNumber(18)).toNumber()] = new BigNumber((_250_v126).cardinality());
          _nw36[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw36[(new BigNumber(20)).toNumber()] = new BigNumber((_35_v3).length);
          _nw36[(new BigNumber(21)).toNumber()] = p1;
          _nw36[(new BigNumber(22)).toNumber()] = p1;
          _nw36[(new BigNumber(23)).toNumber()] = p1;
          _nw36[(new BigNumber(24)).toNumber()] = new BigNumber(930);
          _nw36[(new BigNumber(25)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw36[(new BigNumber(26)).toNumber()] = p1;
          _254_v131 = _nw36;
          let _256_v132;
          _256_v132 = _dafny.Seq.of(_254_v131, _254_v131);
          let _257_v133;
          let _nw37 = Array((new BigNumber(12)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _256_v132;
          _nw37[(_dafny.ONE).toNumber()] = _256_v132;
          _nw37[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_254_v131);
          _nw37[(new BigNumber(3)).toNumber()] = _256_v132;
          _nw37[(new BigNumber(4)).toNumber()] = _256_v132;
          _nw37[(new BigNumber(5)).toNumber()] = _256_v132;
          _nw37[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_254_v131), _256_v132);
          _nw37[(new BigNumber(7)).toNumber()] = _256_v132;
          _nw37[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_256_v132, _256_v132);
          _nw37[(new BigNumber(9)).toNumber()] = ((p2) ? (_dafny.Seq.of(_254_v131)) : (_256_v132));
          _nw37[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_256_v132, _256_v132);
          _nw37[(new BigNumber(11)).toNumber()] = _256_v132;
          _257_v133 = _nw37;
          let _index34 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_257_v133).length));
          (_257_v133)[_index34] = _256_v132;
          let _index35 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_257_v133).length));
          let _rhs54 = _256_v132;
          let _rhs55 = ((new BigNumber((_245_v122).length)).plus(p1)).multipliedBy(p1);
          let _rhs56 = !(!((p2) || ((_232_v112).f17)));
          let _rhs57 = p1;
          let _rhs58 = _254_v131;
          let _lhs35 = _257_v133;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_257_v133).length));
          let _lhs37 = globalState;
          let _lhs38 = globalState;
          _lhs35[_lhs36] = _rhs54;
          _lhs37.f11 = _rhs55;
          _32_v0 = _rhs56;
          _lhs38.f11 = _rhs57;
          _254_v131 = _rhs58;
          _32_v0 = (_248_v125)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_248_v125).length))];
        } else if (_source5.is_DC3) {
          let _258___mcc_h32 = (_source5).cf6;
          let _259___mcc_h33 = (_source5).cf7;
          let _260___mcc_h34 = (_source5).cf8;
          let _261_cf8 = _260___mcc_h34;
          let _262_cf7 = _259___mcc_h33;
          let _263_cf6 = _258___mcc_h32;
          (globalState).f16 = ((_32_v0) ? (p1) : (_262_cf7));
          let _264_v134;
          _264_v134 = _dafny.Map.Empty.slice().updateUnsafe(_33_v1,!(p3));
          let _265_v135;
          _265_v135 = _dafny.Set.fromElements((_232_v112).f17, true);
          _264_v134 = _module.__default.fm15(_32_v0, _265_v135, !(p2), globalState);
          let _266_v136;
          _266_v136 = _module.D3.create_DC10(_235_v115);
          _235_v115 = (_266_v136).dtor_cf18;
          _236_v116 = (_236_v116).update(_262_cf7, (_262_cf7).multipliedBy(_module.__default.fm0(p1, _263_cf6, globalState)));
        } else if (_source5.is_DC4) {
          let _267___mcc_h35 = (_source5).cf9;
          let _268___mcc_h36 = (_source5).cf10;
          let _269___mcc_h37 = (_source5).cf11;
          let _270_cf11 = _269___mcc_h37;
          let _271_cf10 = _268___mcc_h36;
          let _272_cf9 = _267___mcc_h35;
          let _273_v137;
          let _nw38 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
          _273_v137 = _nw38;
          let _274_v138;
          _274_v138 = _dafny.Set.fromElements(_232_v112, _232_v112);
          let _275_v139;
          let _nw39 = Array((new BigNumber(4)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = new BigNumber(-142);
          _nw39[(_dafny.ONE).toNumber()] = new BigNumber(-626);
          _nw39[(new BigNumber(2)).toNumber()] = p1;
          _nw39[(new BigNumber(3)).toNumber()] = new BigNumber((_274_v138).length);
          _275_v139 = _nw39;
          let _276_v140;
          _276_v140 = _dafny.MultiSet.fromElements(_275_v139);
          let _277_v141;
          _277_v141 = _dafny.Seq.of(_module.__default.fm0(new BigNumber((_276_v140).cardinality()), _236_v116, globalState));
          let _index36 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_273_v137).length));
          (_273_v137)[_index36] = _277_v141;
          let _278_v142;
          _278_v142 = _module.D1.create_DC2(_33_v1);
          let _279_v143;
          _279_v143 = _dafny.Seq.of(_278_v142, _278_v142, _278_v142);
          let _280_v145;
          _280_v145 = _dafny.MultiSet.fromElements(function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(998), new BigNumber(785))) {
              let _281_v144 = _compr_24;
              if (((new BigNumber(998)).isLessThanOrEqualTo(_281_v144)) && ((_281_v144).isLessThan(new BigNumber(785)))) {
                _coll24.push([(_281_v144).minus(p1),_272_cf9]);
              }
            }
            return _coll24;
          }(), _236_v116, _module.__default.fm1(_236_v116, _272_cf9, _272_cf9, globalState), _236_v116);
          let _282_v146;
          _282_v146 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(541)), ((_283_cf9) => function (_284_i10) {
            return _283_cf9;
          })(_272_cf9)), _277_v141, _dafny.Seq.of(p1, new BigNumber((_dafny.Seq.of(new BigNumber(454), new BigNumber((_279_v143).length), p1, (_dafny.ZERO).minus(new BigNumber((_35_v3).length)))).length), p1), _277_v141, _dafny.Seq.update(_dafny.Seq.of((((_280_v145).contains(_236_v116)) ? ((_280_v145).get(_236_v116)) : (_272_cf9))), _module.__default.safeIndex(_272_cf9, new BigNumber((_dafny.Seq.of((((_280_v145).contains(_236_v116)) ? ((_280_v145).get(_236_v116)) : (_272_cf9)))).length)), p1));
          let _index37 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_273_v137).length));
          (_273_v137)[_index37] = (_282_v146)[_module.__default.safeIndex(p1, new BigNumber((_282_v146).length))];
          let _285_v147;
          let _nw40 = new _module.C0();
          _nw40.__ctor(_271_cf10);
          _285_v147 = _nw40;
          let _286_v148;
          _286_v148 = _module.D4.create_DC12(_234_v114);
          _234_v114 = (_234_v114).Merge((_286_v148).dtor_cf19);
          let _287_v149;
          let _nw41 = Array((new BigNumber(28)).toNumber()).fill(false);
          _287_v149 = _nw41;
          let _index38 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_287_v149).length));
          (_287_v149)[_index38] = _module.__default.fm4(_module.__default.fm0(new BigNumber((_dafny.Seq.UnicodeFromString("ns")).length), _236_v116, globalState), p2, globalState);
          let _index39 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_287_v149).length));
          (_287_v149)[_index39] = _271_cf10;
        } else if (_source5.is_DC1) {
          let _288___mcc_h38 = (_source5).cf4;
          let _289_cf4 = _288___mcc_h38;
          _236_v116 = _236_v116;
          let _290_v150;
          _290_v150 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("hhlfxth")).length));
          _290_v150 = _290_v150;
          let _291_v151;
          _291_v151 = _dafny.Seq.of(p1, p1);
          let _292_v152;
          _292_v152 = _dafny.Set.fromElements(p1, _module.__default.fm0(new BigNumber((_dafny.Seq.of(p2, p3)).length), _236_v116, globalState), p1, _module.__default.fm0((_291_v151)[_module.__default.safeIndex(p1, new BigNumber((_291_v151).length))], _dafny.Map.Empty.slice().updateUnsafe(p1,p1), globalState));
          let _293_v153;
          _293_v153 = _dafny.Map.Empty.slice().updateUnsafe(p1,_292_v152);
          let _294_v154;
          _294_v154 = _dafny.Set.fromElements((((_293_v153).contains(p1)) ? ((_293_v153).get(p1)) : (_dafny.Set.fromElements(p1, p1, p1))), _292_v152, _292_v152, _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_232_v112).f17)).length)));
          _32_v0 = (_294_v154).contains(_292_v152);
          let _rhs59 = false;
          let _rhs60 = p2;
          _32_v0 = _rhs59;
          _32_v0 = _rhs60;
        } else {
          let _295___mcc_h39 = (_source5).cf12;
          let _296_cf12 = _295___mcc_h39;
          let _297_v155;
          _297_v155 = _dafny.Map.Empty.slice().updateUnsafe(p1,_238_v118);
          let _pat_let_tv8 = _35_v3;
          let _pat_let_tv9 = p1;
          let _pat_let_tv10 = _297_v155;
          let _298_v156;
          let _nw42 = Array((new BigNumber(4)).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = function (_pat_let7_0) {
            return function (_299_dt__update__tmp_h3) {
              return function (_pat_let8_0) {
                return function (_300_dt__update_hcf2_h1) {
                  return function (_pat_let9_0) {
                    return function (_301_dt__update_hcf0_h0) {
                      return function (_pat_let10_0) {
                        return function (_302_dt__update_hcf3_h0) {
                          return _module.D0.create_DC0(_301_dt__update_hcf0_h0, (_299_dt__update__tmp_h3).dtor_cf1, _300_dt__update_hcf2_h1, _302_dt__update_hcf3_h0);
                        }(_pat_let10_0);
                      }(_pat_let_tv10);
                    }(_pat_let9_0);
                  }(_pat_let_tv9);
                }(_pat_let8_0);
              }(_pat_let_tv8);
            }(_pat_let7_0);
          }(_36_v4);
          _nw42[(_dafny.ONE).toNumber()] = _36_v4;
          _nw42[(new BigNumber(2)).toNumber()] = _36_v4;
          _nw42[(new BigNumber(3)).toNumber()] = _36_v4;
          _298_v156 = _nw42;
          let _index40 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_298_v156).length));
          (_298_v156)[_index40] = _36_v4;
          let _303_v158;
          _303_v158 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC11(),(_232_v112).f17);
          let _pat_let_tv11 = _231_cf4;
          let _pat_let_tv12 = p2;
          let _pat_let_tv13 = globalState;
          let _pat_let_tv14 = p1;
          let _index41 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_298_v156).length));
          (_298_v156)[_index41] = function (_pat_let11_0) {
            return function (_305_dt__update__tmp_h4) {
              return function (_pat_let12_0) {
                return function (_306_dt__update_hcf3_h1) {
                  return function (_pat_let13_0) {
                    return function (_307_dt__update_hcf0_h1) {
                      return _module.D0.create_DC0(_307_dt__update_hcf0_h1, (_305_dt__update__tmp_h4).dtor_cf1, (_305_dt__update__tmp_h4).dtor_cf2, _306_dt__update_hcf3_h1);
                    }(_pat_let13_0);
                  }(_pat_let_tv14);
                }(_pat_let12_0);
              }(_module.__default.fm11(new BigNumber((_pat_let_tv11).length), _pat_let_tv12, _pat_let_tv13));
            }(_pat_let11_0);
          }(_module.__default.fm16(_33_v1, _module.__default.fm17(new BigNumber(597), globalState), function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of (_303_v158).Keys.Elements) {
              let _304_v157 = _compr_25;
              if ((_303_v158).contains(_304_v157)) {
                _coll25.push([_304_v157,p1]);
              }
            }
            return _coll25;
          }(), globalState));
          let _308_v159;
          let _nw43 = new _module.C0();
          _nw43.__ctor((_232_v112).f17);
          _308_v159 = _nw43;
          let _309_v160;
          let _nw44 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _309_v160 = _nw44;
          let _310_v161;
          _310_v161 = _dafny.Seq.of(((p3) ? (_309_v160) : (_309_v160)), _309_v160, _309_v160, _309_v160);
          _310_v161 = _dafny.Seq.Concat(_dafny.Seq.of(_309_v160), _310_v161);
          _32_v0 = _32_v0;
        }
      } else {
        let _311___mcc_h8 = (_source1).cf12;
        let _312_cf12 = _311___mcc_h8;
        if (_32_v0) {
          let _313_v162;
          _313_v162 = _dafny.MultiSet.fromElements(p2, !(!(p3)));
          let _314_v163;
          let _nw45 = new _module.C0();
          _nw45.__ctor(p2);
          _314_v163 = _nw45;
          let _315_v164;
          _315_v164 = _dafny.Map.Empty.slice().updateUnsafe((_313_v162).Intersect((_313_v162).update(_32_v0, _module.__default.abs(p1))),_314_v163);
          _315_v164 = (_315_v164).update(_313_v162, _314_v163);
          _32_v0 = p3;
          let _316_v165;
          let _nw46 = Array((new BigNumber(19)).toNumber());
          _316_v165 = _nw46;
          let _index42 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_316_v165).length));
          (_316_v165)[_index42] = _314_v163;
          let _index43 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_316_v165).length));
          (_316_v165)[_index43] = _314_v163;
          let _317_v166;
          _317_v166 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,true);
          _317_v166 = (((_314_v163).f17) ? ((_317_v166).update((_314_v163).f17, (_314_v163).f17)) : ((_317_v166).Merge(_317_v166)));
          (globalState).f11 = p1;
        } else {
          let _rhs61 = p1;
          let _rhs62 = !(p2);
          let _lhs39 = globalState;
          _lhs39.f6 = _rhs61;
          _32_v0 = _rhs62;
          let _318_v167;
          let _nw47 = new _module.C0();
          _nw47.__ctor(p3);
          _318_v167 = _nw47;
          let _319_v168;
          let _init4 = function (_320_i11) {
            return true;
          };
          let _nw48 = Array((new BigNumber(28)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw48.length); _i0_4++) {
            _nw48[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _319_v168 = _nw48;
          let _index44 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_319_v168).length));
          (_319_v168)[_index44] = (((_318_v167).f17) ? ((_318_v167).f17) : (_32_v0));
          let _index45 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_319_v168).length));
          (_319_v168)[_index45] = p2;
          _32_v0 = _module.__default.fm4(p1, (_319_v168)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_319_v168).length))], globalState);
          let _321_v169;
          _321_v169 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,p1);
          let _322_v170;
          _322_v170 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_319_v168)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_319_v168).length))], (_319_v168)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_319_v168).length))])).length), (((_321_v169).contains(p2)) ? ((_321_v169).get(p2)) : (p1)), (_dafny.ZERO).minus(p1));
          let _index46 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_319_v168).length));
          (_319_v168)[_index46] = !((_dafny.Set.fromElements(p1)).IsDisjointFrom(_322_v170));
        }
        let _323_v171;
        _323_v171 = _dafny.MultiSet.fromElements(_32_v0, true);
        if (_module.__default.fm4(p1, !(new BigNumber((_323_v171).cardinality())).isEqualTo(p1), globalState)) {
          let _324_v172;
          let _nw49 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _324_v172 = _nw49;
          let _index47 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_324_v172).length));
          (_324_v172)[_index47] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kinwdfjx"), _35_v3);
          let _index48 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_324_v172).length));
          (_324_v172)[_index48] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), function (_325_i12) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          });
          let _326_v173;
          let _nw50 = new _module.C0();
          _nw50.__ctor((_323_v171).IsProperSubsetOf(_323_v171));
          _326_v173 = _nw50;
          let _327_v174;
          _327_v174 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(522),p1);
          let _328_v175;
          _328_v175 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          (globalState).f16 = (_module.__default.fm0(p1, _327_v174, globalState)).plus(new BigNumber((((p3) ? (_328_v175) : (_module.__default.fm5(globalState)))).length));
          let _index49 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_324_v172).length));
          (_324_v172)[_index49] = (_324_v172)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_324_v172).length))];
          let _329_v176;
          let _nw51 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
          _329_v176 = _nw51;
          let _index50 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_329_v176).length));
          (_329_v176)[_index50] = (_323_v171).update(_32_v0, _module.__default.abs(p1));
          let _330_v177;
          let _nw52 = Array((new BigNumber(20)).toNumber()).fill(_module.D4.Default());
          _330_v177 = _nw52;
          let _331_v178;
          let _init5 = ((_332_v1) => function (_333_i13) {
            return _332_v1;
          })(_33_v1);
          let _nw53 = Array((new BigNumber(7)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw53.length); _i0_5++) {
            _nw53[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _331_v178 = _nw53;
          let _334_v179;
          _334_v179 = _dafny.Set.fromElements(_331_v178);
          let _335_v180;
          _335_v180 = _module.D4.create_DC15(_334_v179, _33_v1);
          let _index51 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_330_v177).length));
          (_330_v177)[_index51] = _335_v180;
          let _336_v181;
          let _nw54 = Array((new BigNumber(27)).toNumber()).fill(false);
          _336_v181 = _nw54;
          let _index52 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_329_v176).length));
          let _index53 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_330_v177).length));
          let _rhs63 = ((p2) ? (_module.__default.fm6(p2, p3, p2, globalState)) : (_323_v171));
          let _rhs64 = _module.D4.create_DC15(_334_v179, _33_v1);
          let _rhs65 = (_module.D4.create_DC13(_module.__default.fm4(p1, p3, globalState), p1, _336_v181)).dtor_cf20;
          let _lhs40 = _329_v176;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_329_v176).length));
          let _lhs42 = _330_v177;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_330_v177).length));
          _lhs40[_lhs41] = _rhs63;
          _lhs42[_lhs43] = _rhs64;
          _32_v0 = _rhs65;
        } else {
          let _337_v182;
          _337_v182 = _dafny.Seq.of(!(p2));
          _32_v0 = (_337_v182)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-856))), new BigNumber((_337_v182).length))];
          let _338_v183;
          _338_v183 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _339_v184;
          let _nw55 = new _module.C0();
          _nw55.__ctor(p2);
          _339_v184 = _nw55;
          let _340_v185;
          _340_v185 = _dafny.Map.Empty.slice().updateUnsafe(p1,_339_v184);
          let _341_v186;
          _341_v186 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p1, _338_v183, globalState),(_340_v185).Merge(_340_v185));
          let _342_v187;
          _342_v187 = _dafny.Seq.of(_341_v186);
          _341_v186 = (_341_v186).Merge((_342_v187)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).length), new BigNumber((_342_v187).length))]);
          let _343_v188;
          _343_v188 = _dafny.Map.Empty.slice().updateUnsafe((_339_v184).f17,_33_v1);
          let _344_v189;
          _344_v189 = _dafny.Set.fromElements(_35_v3);
          let _345_v190;
          _345_v190 = _module.D4.create_DC14(new BigNumber((_343_v188).length), (new BigNumber((_338_v183).length)).minus(p1), _344_v189);
          _345_v190 = _345_v190;
          let _346_v191;
          let _nw56 = Array((new BigNumber(26)).toNumber()).fill(false);
          _346_v191 = _nw56;
          _346_v191 = _346_v191;
          _32_v0 = !(!(p3) || (_32_v0));
        }
        let _347_v192;
        let _init6 = ((_348_v171, _349_v3) => function (_350_i14) {
          return _module.D4.create_DC14(new BigNumber((_348_v171).cardinality()), new BigNumber(996), _dafny.Set.fromElements(_349_v3, _349_v3));
        })(_323_v171, _35_v3);
        let _nw57 = Array((_dafny.ONE).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw57.length); _i0_6++) {
          _nw57[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _347_v192 = _nw57;
        let _351_v193;
        let _nw58 = new _module.C0();
        _nw58.__ctor(!(p3));
        _351_v193 = _nw58;
        let _352_v194;
        _352_v194 = _module.D4.create_DC14(new BigNumber((_35_v3).length), new BigNumber((_dafny.Seq.of(_351_v193)).length), _module.__default.fm18(p1, p1, p1, _33_v1, globalState));
        let _353_v195;
        _353_v195 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("fpagmi"));
        let _pat_let_tv15 = _353_v195;
        let _index54 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_347_v192).length));
        (_347_v192)[_index54] = function (_pat_let14_0) {
          return function (_354_dt__update__tmp_h5) {
            return function (_pat_let15_0) {
              return function (_355_dt__update_hcf25_h0) {
                return _module.D4.create_DC14((_354_dt__update__tmp_h5).dtor_cf23, (_354_dt__update__tmp_h5).dtor_cf24, _355_dt__update_hcf25_h0);
              }(_pat_let15_0);
            }(_pat_let_tv15);
          }(_pat_let14_0);
        }(_352_v194);
        let _index55 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_347_v192).length));
        (_347_v192)[_index55] = _module.D4.create_DC14(p1, p1, _dafny.Set.fromElements(_35_v3));
        let _356_v196;
        _356_v196 = _dafny.Seq.of(_32_v0);
        let _357_v197;
        _357_v197 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_356_v196).length));
        let _358_v198;
        _358_v198 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,_357_v197);
        _358_v198 = _module.__default.fm19((_dafny.MultiSet.fromElements(_32_v0, p3, (_351_v193).f17, p2)).Difference(_323_v171), globalState);
      }
      let _359_v199;
      let _nw59 = new _module.C0();
      _nw59.__ctor(p2);
      _359_v199 = _nw59;
      let _360_i15;
      _360_i15 = _dafny.ZERO;
      L0: {
        while (p3) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_360_i15)) {
              break L0;
            }
            _360_i15 = (_360_i15).plus(_dafny.ONE);
            let _361_v200;
            let _nw60 = Array((new BigNumber(25)).toNumber()).fill([]);
            _361_v200 = _nw60;
            let _362_v201;
            _362_v201 = _dafny.Map.Empty.slice().updateUnsafe(_361_v200,_359_v199);
            _362_v201 = (_362_v201).update(_361_v200, _359_v199);
            (globalState).f16 = p1;
            if ((_359_v199).f17) {
              _32_v0 = p3;
              _33_v1 = _33_v1;
              let _363_v202;
              let _nw61 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
              _363_v202 = _nw61;
              let _index56 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_363_v202).length));
              (_363_v202)[_index56] = _dafny.Set.fromElements(p3, _module.__default.fm4(p1, (_359_v199).f17, globalState));
              let _364_v203;
              _364_v203 = _dafny.Set.fromElements(!((_359_v199).f17));
              let _index57 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_363_v202).length));
              (_363_v202)[_index57] = _364_v203;
              let _nw62 = new _module.C0();
              _nw62.__ctor((_359_v199).f17);
              _359_v199 = _nw62;
              let _365_v204;
              let _nw63 = Array((new BigNumber(13)).toNumber());
              _365_v204 = _nw63;
              let _index58 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_365_v204).length));
              (_365_v204)[_index58] = _359_v199;
              let _index59 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_365_v204).length));
              let _rhs66 = _32_v0;
              let _rhs67 = _35_v3;
              let _rhs68 = new BigNumber((_35_v3).length);
              let _rhs69 = (((_359_v199).f17) ? (_359_v199) : (_359_v199));
              let _lhs44 = globalState;
              let _lhs45 = _365_v204;
              let _lhs46 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_365_v204).length));
              _32_v0 = _rhs66;
              _35_v3 = _rhs67;
              _lhs44.f16 = _rhs68;
              _lhs45[_lhs46] = _rhs69;
            } else {
              let _366_v205;
              _366_v205 = _module.D2.create_DC7(p1, _33_v1);
              _366_v205 = _366_v205;
              let _367_v206;
              let _nw64 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
              _367_v206 = _nw64;
              let _index60 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length));
              (_361_v200)[_index60] = _367_v206;
              let _index61 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length));
              (_361_v200)[_index61] = _367_v206;
              let _368_v207;
              _368_v207 = _dafny.Map.Empty.slice().updateUnsafe(_359_v199,_33_v1);
              _32_v0 = _module.__default.fm4((_dafny.ZERO).minus(p1), !((_368_v207).update(_359_v199, _33_v1)).contains(_359_v199), globalState);
              _32_v0 = (((p1).isEqualTo(new BigNumber((_module.__default.fm20(p3, globalState)).cardinality()))) ? ((_359_v199).f17) : (p2));
              let _369_v208;
              _369_v208 = _dafny.MultiSet.fromElements(p1);
              let _index62 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((p0).length));
              (p0)[_index62] = _369_v208;
              let _arr0 = (_361_v200)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length))];
              let _index63 = _module.__default.safeIndex(new BigNumber(887), new BigNumber(((_361_v200)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length))]).length));
              _arr0[_index63] = p1;
              let _370_v209;
              let _nw65 = Array((_dafny.ONE).toNumber());
              _nw65[(_dafny.ZERO).toNumber()] = (_359_v199).f17;
              _370_v209 = _nw65;
              let _371_v210;
              _371_v210 = _dafny.Seq.of((_359_v199).f17, false, (_359_v199).f17);
              let _372_v211;
              _372_v211 = _dafny.Set.fromElements(new BigNumber((_371_v210).length), p1);
              let _index64 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_370_v209).length));
              (_370_v209)[_index64] = (_dafny.Set.fromElements(p1)).IsProperSubsetOf(_372_v211);
              let _373_v212;
              _373_v212 = _dafny.Seq.of((_dafny.MultiSet.fromElements(p1)).Difference(_369_v208), _369_v208, _369_v208, (_369_v208).Difference(_dafny.MultiSet.fromElements(p1, p1)));
              let _374_v213;
              _374_v213 = _dafny.Map.Empty.slice().updateUnsafe(p1,_359_v199);
              let _375_v214;
              _375_v214 = _dafny.Map.Empty.slice().updateUnsafe(p1,_374_v213);
              let _376_v215;
              _376_v215 = _dafny.Set.fromElements(false);
              let _377_v216;
              _377_v216 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _index65 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((p0).length));
              let _arr1 = (_361_v200)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length))];
              let _index66 = _module.__default.safeIndex(new BigNumber(887), new BigNumber(((_361_v200)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length))]).length));
              let _index67 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_370_v209).length));
              let _rhs70 = !(!((_32_v0) || ((_359_v199).f17)));
              let _rhs71 = (_373_v212)[_module.__default.safeIndex(p1, new BigNumber((_373_v212).length))];
              let _rhs72 = p1;
              let _rhs73 = (((_369_v208).contains(new BigNumber((_375_v214).length))) ? ((_369_v208).get(new BigNumber((_375_v214).length))) : (_module.__default.fm0(new BigNumber((_376_v215).length), _377_v216, globalState)));
              let _rhs74 = false;
              let _lhs47 = p0;
              let _lhs48 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((p0).length));
              let _lhs49 = (_361_v200)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length))];
              let _lhs50 = _module.__default.safeIndex(new BigNumber(887), new BigNumber(((_361_v200)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_361_v200).length))]).length));
              let _lhs51 = globalState;
              let _lhs52 = _370_v209;
              let _lhs53 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_370_v209).length));
              _32_v0 = _rhs70;
              _lhs47[_lhs48] = _rhs71;
              _lhs49[_lhs50] = _rhs72;
              _lhs51.f6 = _rhs73;
              _lhs52[_lhs53] = _rhs74;
            }
            let _378_v217;
            let _nw66 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _378_v217 = _nw66;
            let _379_v218;
            _379_v218 = _dafny.Seq.of(p1);
            let _380_v219;
            _380_v219 = _dafny.Seq.of(_379_v218);
            let _381_v220;
            _381_v220 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,p1);
            let _index68 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_378_v217).length));
            (_378_v217)[_index68] = new BigNumber((_dafny.Seq.Concat((_380_v219)[_module.__default.safeIndex((((_381_v220).contains((_359_v199).f17)) ? ((_381_v220).get((_359_v199).f17)) : (p1)), new BigNumber((_380_v219).length))], _dafny.Seq.of(p1))).length);
            let _index69 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_378_v217).length));
            (_378_v217)[_index69] = (p1).plus((p1).plus(p1));
          }
        }
      }
      let _382_v221;
      _382_v221 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(415));
      let _383_v222;
      _383_v222 = _dafny.Seq.of(p3);
      r0 = ((_module.__default.fm0(p1, _382_v221, globalState)).plus(new BigNumber(786))).minus(new BigNumber((_dafny.Seq.Concat(_383_v222, _dafny.Seq.of(p3))).length));
      let _384_v223;
      _384_v223 = _module.D1.create_DC2(_33_v1);
      let _385_v224;
      _385_v224 = _dafny.Map.Empty.slice().updateUnsafe(_384_v223,_32_v0);
      let _386_v225;
      _386_v225 = _dafny.MultiSet.fromElements((_383_v222)[_module.__default.safeIndex(p1, new BigNumber((_383_v222).length))]);
      let _387_v226;
      _387_v226 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
      let _388_v227;
      let _nw67 = Array((new BigNumber(27)).toNumber());
      _nw67[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of(_32_v0, (_359_v199).f17, (_359_v199).f17, p2, _32_v0)).length);
      _nw67[(_dafny.ONE).toNumber()] = (p1).plus((_dafny.ZERO).minus(p1));
      _nw67[(new BigNumber(2)).toNumber()] = ((p3) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(p1))) : (p1));
      _nw67[(new BigNumber(3)).toNumber()] = p1;
      _nw67[(new BigNumber(4)).toNumber()] = p1;
      _nw67[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(290), p1);
      _nw67[(new BigNumber(6)).toNumber()] = p1;
      _nw67[(new BigNumber(7)).toNumber()] = new BigNumber(752);
      _nw67[(new BigNumber(8)).toNumber()] = ((p2) ? (p1) : (new BigNumber((_385_v224).length)));
      _nw67[(new BigNumber(9)).toNumber()] = p1;
      _nw67[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_35_v3).length));
      _nw67[(new BigNumber(11)).toNumber()] = p1;
      _nw67[(new BigNumber(12)).toNumber()] = (new BigNumber(64)).plus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(552)), ((_389_v1) => function (_390_i16) {
        return _389_v1;
      })(_33_v1)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(552)), ((_391_v1) => function (_392_i16) {
        return _391_v1;
      })(_33_v1))).length)), _33_v1)).length));
      _nw67[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_386_v225).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_387_v226).contains(p2)) ? ((_387_v226).get(p2)) : (new BigNumber((_382_v221).length))),_359_v199)).length));
      _nw67[(new BigNumber(14)).toNumber()] = p1;
      _nw67[(new BigNumber(15)).toNumber()] = p1;
      _nw67[(new BigNumber(16)).toNumber()] = p1;
      _nw67[(new BigNumber(17)).toNumber()] = p1;
      _nw67[(new BigNumber(18)).toNumber()] = (p1).plus((_dafny.ZERO).minus(p1));
      _nw67[(new BigNumber(19)).toNumber()] = p1;
      _nw67[(new BigNumber(20)).toNumber()] = p1;
      _nw67[(new BigNumber(21)).toNumber()] = p1;
      _nw67[(new BigNumber(22)).toNumber()] = p1;
      _nw67[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(470)), ((_393_p1) => function (_394_i17) {
        return _393_p1;
      })(p1))).length);
      _nw67[(new BigNumber(24)).toNumber()] = p1;
      _nw67[(new BigNumber(25)).toNumber()] = p1;
      _nw67[(new BigNumber(26)).toNumber()] = p1;
      _388_v227 = _nw67;
      r1 = _388_v227;
      let _395_v228;
      _395_v228 = _dafny.Set.fromElements(((_32_v0) ? ((_359_v199).f17) : (_32_v0)), (_359_v199).f17, p3);
      r2 = _395_v228;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _396_v1;
      _396_v1 = new _dafny.CodePoint('d'.codePointAt(0));
      let _397_v2;
      _397_v2 = new BigNumber(514);
      let _398_v4;
      _398_v4 = false;
      let _399_v5;
      _399_v5 = _dafny.Seq.of(_398_v4, true);
      let _400_v6;
      _400_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_399_v5).length),true);
      let _401_v7;
      _401_v7 = _dafny.Set.fromElements((((_400_v6).contains(_397_v2)) ? ((_400_v6).get(_397_v2)) : (!(_398_v4))), _398_v4, _398_v4, _398_v4);
      let _402_v8;
      _402_v8 = _dafny.Seq.UnicodeFromString("frrp");
      let _403_v9;
      let _init7 = function (_404_i0) {
        return _module.__default.safeDivisionInt(_404_i0, new BigNumber(535));
      };
      let _nw68 = Array((new BigNumber(7)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw68.length); _i0_7++) {
        _nw68[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _403_v9 = _nw68;
      let _405_globalState;
      let _nw69 = new _module.GlobalState();
      _nw69.__ctor(new BigNumber(455), new BigNumber(119), (function () {
        let _coll26 = new _dafny.Map();
        for (const _compr_26 of (_dafny.Seq.of(_396_v1)).Elements) {
          let _406_v0 = _compr_26;
          if (_dafny.Seq.contains(_dafny.Seq.of(_396_v1), _406_v0)) {
            _coll26.push([_406_v0,new BigNumber(68)]);
          }
        }
        return _coll26;
      }()).update(_396_v1, _397_v2), true, new BigNumber(936), new BigNumber(277), new BigNumber(911), _dafny.MultiSet.fromElements(_397_v2, (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll27 = new _dafny.Set();
        for (const _compr_27 of (_dafny.Seq.of(_397_v2)).Elements) {
          let _407_v3 = _compr_27;
          if (_dafny.Seq.contains(_dafny.Seq.of(_397_v2), _407_v3)) {
            _coll27.add((_407_v3).multipliedBy((_dafny.ZERO).minus(new BigNumber(-233))));
          }
        }
        return _coll27;
      }()).length)), new BigNumber((_401_v7).length), new BigNumber((_402_v8).length)), true, new BigNumber(828), new BigNumber(622), new BigNumber(-572), new BigNumber(262), false, _403_v9, new BigNumber(726), new BigNumber(237));
      _405_globalState = _nw69;
      let _408_v10;
      _408_v10 = _dafny.Map.Empty.slice().updateUnsafe(_397_v2,_397_v2);
      let _hi0 = _397_v2;
      for (let _409_i1 = _module.__default.fm0(_397_v2, (_408_v10).update(new BigNumber(-923), _397_v2), _405_globalState); _409_i1.isLessThan(_hi0); _409_i1 = _409_i1.plus(_dafny.ONE)) {
        (_405_globalState).f16 = _409_i1;
        let _410_v11;
        _410_v11 = _dafny.MultiSet.fromElements(!(_398_v4), _398_v4, _398_v4);
        let _411_v12;
        _411_v12 = _dafny.MultiSet.fromElements(_410_v11);
        let _412_v13;
        _412_v13 = _dafny.Seq.of((((_411_v12).contains(_410_v11)) ? ((_411_v12).get(_410_v11)) : (_409_i1)));
        _412_v13 = _412_v13;
        let _413_v14;
        _413_v14 = _dafny.Map.Empty.slice().updateUnsafe(_398_v4,(_400_v6).Merge(_400_v6));
        _413_v14 = _413_v14;
        let _414_v16;
        _414_v16 = _module.D0.create_DC0(_409_i1, _398_v4, _dafny.Seq.UnicodeFromString("yicqicecd"), function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(396), new BigNumber(-156))) {
    let _415_v15 = _compr_28;
    if (((new BigNumber(396)).isLessThanOrEqualTo(_415_v15)) && ((_415_v15).isLessThan(new BigNumber(-156)))) {
      _coll28.push([(_415_v15).minus(new BigNumber((_402_v8).length)),_400_v6]);
    }
  }
  return _coll28;
}());
        let _source6 = _414_v16;
        let _416___mcc_h0 = (_source6).cf0;
        let _417___mcc_h1 = (_source6).cf1;
        let _418___mcc_h2 = (_source6).cf2;
        let _419___mcc_h3 = (_source6).cf3;
        let _420_cf3 = _419___mcc_h3;
        let _421_cf2 = _418___mcc_h2;
        let _422_cf1 = _417___mcc_h1;
        let _423_cf0 = _416___mcc_h0;
        let _index70 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_403_v9).length));
        (_403_v9)[_index70] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.fm0(_423_cf0, _408_v10, _405_globalState)).multipliedBy(_409_i1))));
        let _index71 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_403_v9).length));
        (_403_v9)[_index71] = _397_v2;
        _408_v10 = (_408_v10).update((_dafny.ZERO).minus((_403_v9)[_module.__default.safeIndex(new BigNumber(540), new BigNumber((_403_v9).length))]), (_403_v9)[_module.__default.safeIndex(new BigNumber(540), new BigNumber((_403_v9).length))]);
        let _424_v17;
        _424_v17 = _module.D1.create_DC1(_399_v5);
        let _425_v18;
        _425_v18 = _dafny.MultiSet.fromElements(_423_cf0);
        let _426_v19;
        _426_v19 = _dafny.Map.Empty.slice().updateUnsafe(_422_cf1,_398_v4);
        let _427_v20;
        _427_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_403_v9)[_module.__default.safeIndex(new BigNumber(540), new BigNumber((_403_v9).length))]),_399_v5);
        let _428_v21;
        let _nw70 = Array((new BigNumber(25)).toNumber());
        _nw70[(_dafny.ZERO).toNumber()] = _399_v5;
        _nw70[(_dafny.ONE).toNumber()] = _399_v5;
        _nw70[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat((_424_v17).dtor_cf4, _399_v5);
        _nw70[(new BigNumber(3)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(4)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(5)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(6)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_399_v5, _module.__default.safeIndex((_dafny.ZERO).minus(_423_cf0), new BigNumber((_399_v5).length)), _422_cf1);
        _nw70[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_399_v5, _module.__default.safeIndex(new BigNumber((_425_v18).cardinality()), new BigNumber((_399_v5).length)), _398_v4), _module.__default.safeIndex(_409_i1, new BigNumber((_dafny.Seq.update(_399_v5, _module.__default.safeIndex(new BigNumber((_425_v18).cardinality()), new BigNumber((_399_v5).length)), _398_v4)).length)), _422_cf1), _dafny.Seq.of(false, (((_426_v19).contains(_422_cf1)) ? ((_426_v19).get(_422_cf1)) : (_398_v4))));
        _nw70[(new BigNumber(9)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(10)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(11)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(12)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(13)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(14)).toNumber()] = (((_427_v20).contains(new BigNumber(-371))) ? ((_427_v20).get(new BigNumber(-371))) : (_dafny.Seq.of(_422_cf1)));
        _nw70[(new BigNumber(15)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_399_v5, _399_v5);
        _nw70[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(false);
        _nw70[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_399_v5, _399_v5);
        _nw70[(new BigNumber(19)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_399_v5, (_module.D1.create_DC1(_399_v5)).dtor_cf4);
        _nw70[(new BigNumber(21)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(22)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(23)).toNumber()] = _399_v5;
        _nw70[(new BigNumber(24)).toNumber()] = _399_v5;
        _428_v21 = _nw70;
        let _index72 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_428_v21).length));
        (_428_v21)[_index72] = _dafny.Seq.Concat(_dafny.Seq.of(_398_v4, false), _399_v5);
        let _429_v22;
        _429_v22 = _dafny.MultiSet.fromElements(_402_v8);
        let _430_v23;
        let _init8 = ((_431_i1, _432_cf1, _433_v8, _434_cf3) => function (_435_i2) {
          return _module.D0.create_DC0(_431_i1, _432_cf1, _433_v8, _434_cf3);
        })(_409_i1, _422_cf1, _402_v8, _420_cf3);
        let _nw71 = Array((new BigNumber(28)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw71.length); _i0_8++) {
          _nw71[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _430_v23 = _nw71;
        let _436_v24;
        _436_v24 = _dafny.Seq.of(_dafny.Set.fromElements(_398_v4));
        let _437_v25;
        let _nw72 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _437_v25 = _nw72;
        let _438_v26;
        _438_v26 = _dafny.Set.fromElements(_437_v25);
        let _index73 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_428_v21).length));
        let _rhs75 = _dafny.Seq.Concat(_399_v5, _399_v5);
        let _rhs76 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jaeue"));
        let _rhs77 = _430_v23;
        let _rhs78 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), ((_439_v7) => function (_440_i3) {
          return _439_v7;
        })(_401_v7));
        let _rhs79 = _module.__default.fm0(_397_v2, (_module.__default.fm1(_408_v10, new BigNumber((_408_v10).length), new BigNumber((_438_v26).length), _405_globalState)).Merge(_408_v10), _405_globalState);
        let _lhs54 = _428_v21;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_428_v21).length));
        let _lhs56 = _405_globalState;
        _lhs54[_lhs55] = _rhs75;
        _429_v22 = _rhs76;
        _430_v23 = _rhs77;
        _436_v24 = _rhs78;
        _lhs56.f6 = _rhs79;
        let _441_v27;
        _441_v27 = _dafny.MultiSet.fromElements(_437_v25);
        _421_cf2 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("flvuwul"), _module.__default.safeIndex(_423_cf0, new BigNumber((_dafny.Seq.UnicodeFromString("flvuwul")).length)), ((_398_v4) ? (_396_v1) : ((_402_v8)[_module.__default.safeIndex(_423_cf0, new BigNumber((_402_v8).length))]))), _module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber(240), new BigNumber((_441_v27).cardinality())), new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("flvuwul"), _module.__default.safeIndex(_423_cf0, new BigNumber((_dafny.Seq.UnicodeFromString("flvuwul")).length)), ((_398_v4) ? (_396_v1) : ((_402_v8)[_module.__default.safeIndex(_423_cf0, new BigNumber((_402_v8).length))])))).length)), _396_v1);
      }
      let _index74 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length));
      (_403_v9)[_index74] = _397_v2;
      let _index75 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length));
      (_403_v9)[_index75] = _module.__default.safeModuloInt(_397_v2, _397_v2);
      let _442_v28;
      let _nw73 = new _module.C0();
      _nw73.__ctor(_398_v4);
      _442_v28 = _nw73;
      let _443_v29;
      _443_v29 = _dafny.MultiSet.fromElements(_397_v2);
      let _444_v30;
      _444_v30 = _dafny.Map.Empty.slice().updateUnsafe((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))],_dafny.MultiSet.fromElements(new BigNumber(76)));
      let _445_v31;
      _445_v31 = _module.D1.create_DC4(_397_v2, (_442_v28).f17, _444_v30);
      let _446_v32;
      _446_v32 = _module.D1.create_DC5(_445_v31);
      let _pat_let_tv16 = _398_v4;
      let _pat_let_tv17 = _442_v28;
      let _pat_let_tv18 = _398_v4;
      let _pat_let_tv19 = _442_v28;
      let _pat_let_tv20 = _398_v4;
      let _rhs80 = _module.__default.fm3(_398_v4, _405_globalState);
      let _rhs81 = !(_398_v4) || ((_443_v29).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(187), new BigNumber((_408_v10).length))));
      let _rhs82 = function (_source7) {
        if (_source7.is_DC2) {
          let _447___mcc_h4 = (_source7).cf5;
          let _448_cf5 = _447___mcc_h4;
          return true;
        } else if (_source7.is_DC3) {
          let _449___mcc_h5 = (_source7).cf6;
          let _450___mcc_h6 = (_source7).cf7;
          let _451___mcc_h7 = (_source7).cf8;
          let _452_cf8 = _451___mcc_h7;
          let _453_cf7 = _450___mcc_h6;
          let _454_cf6 = _449___mcc_h5;
          return _pat_let_tv16;
        } else if (_source7.is_DC4) {
          let _455___mcc_h8 = (_source7).cf9;
          let _456___mcc_h9 = (_source7).cf10;
          let _457___mcc_h10 = (_source7).cf11;
          let _458_cf11 = _457___mcc_h10;
          let _459_cf10 = _456___mcc_h9;
          let _460_cf9 = _455___mcc_h8;
          return (_pat_let_tv17).f17;
        } else if (_source7.is_DC1) {
          let _461___mcc_h11 = (_source7).cf4;
          let _462_cf4 = _461___mcc_h11;
          return _pat_let_tv18;
        } else {
          let _463___mcc_h12 = (_source7).cf12;
          let _464_cf12 = _463___mcc_h12;
          return !((_pat_let_tv19).f17) || (_pat_let_tv20);
        }
      }(_446_v32);
      let _rhs83 = ((_398_v4) ? (_396_v1) : (_396_v1));
      let _rhs84 = _396_v1;
      _396_v1 = _rhs80;
      _398_v4 = _rhs81;
      _398_v4 = _rhs82;
      _396_v1 = _rhs83;
      _396_v1 = _rhs84;
      let _465_v33;
      _465_v33 = _dafny.MultiSet.fromElements(false);
      let _466_v34;
      let _nw74 = Array((new BigNumber(6)).toNumber()).fill(false);
      _466_v34 = _nw74;
      let _index76 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
      (_466_v34)[_index76] = _398_v4;
      let _index77 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_466_v34).length));
      (_466_v34)[_index77] = (_442_v28).f17;
      let _467_v35;
      _467_v35 = _dafny.Map.Empty.slice().updateUnsafe((_442_v28).f17,(_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]);
      let _468_v36;
      _468_v36 = _dafny.Seq.of(_408_v10, _408_v10);
      let _index78 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
      let _index79 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_466_v34).length));
      let _rhs85 = _465_v33;
      let _rhs86 = _module.__default.fm4(((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).plus((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]), ((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).isLessThanOrEqualTo((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]), _405_globalState);
      let _rhs87 = _module.__default.fm4(new BigNumber((_dafny.Seq.update(_402_v8, _module.__default.safeIndex((((_467_v35).contains(_module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_442_v28).f17, _405_globalState))) ? ((_467_v35).get(_module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_442_v28).f17, _405_globalState))) : (_module.__default.fm0(new BigNumber(16), (_468_v36)[_module.__default.safeIndex((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], new BigNumber((_468_v36).length))], _405_globalState))), new BigNumber((_402_v8).length)), _396_v1)).length), (_442_v28).f17, _405_globalState);
      let _rhs88 = _module.__default.fm4(_397_v2, (_442_v28).f17, _405_globalState);
      let _lhs57 = _466_v34;
      let _lhs58 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
      let _lhs59 = _466_v34;
      let _lhs60 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_466_v34).length));
      _465_v33 = _rhs85;
      _398_v4 = _rhs86;
      _lhs57[_lhs58] = _rhs87;
      _lhs59[_lhs60] = _rhs88;
      let _469_v37;
      let _nw75 = new _module.C0();
      _nw75.__ctor((_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))]);
      _469_v37 = _nw75;
      let _470_v38;
      _470_v38 = _dafny.Map.Empty.slice().updateUnsafe((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))],_402_v8);
      _470_v38 = (_470_v38).update(new BigNumber((_399_v5).length), _402_v8);
      let _index80 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
      (_466_v34)[_index80] = _dafny.Seq.IsPrefixOf(((false) ? (_402_v8) : (_dafny.Seq.update(_402_v8, _module.__default.safeIndex(_397_v2, new BigNumber((_402_v8).length)), _396_v1))), _402_v8);
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_403_v9).length))) {
        let _471_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_471_i4)) && ((_471_i4).isLessThan(new BigNumber((_403_v9).length))))) {
          (_403_v9)[(_471_i4)] = (_471_i4).multipliedBy(new BigNumber(671));
        }
      }
      let _472_v39;
      let _nw76 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _472_v39 = _nw76;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_472_v39).length))) {
        let _473_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_473_i5)) && ((_473_i5).isLessThan(new BigNumber((_472_v39).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_472_v39, (_473_i5).toNumber(), _dafny.Seq.Concat(_dafny.Seq.of((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]), _dafny.Seq.Concat(_dafny.Seq.of(_397_v2), _dafny.Seq.of(new BigNumber(129))))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      (_405_globalState).f16 = _397_v2;
      let _474_i6;
      _474_i6 = _dafny.ZERO;
      L1: {
        while ((_469_v37).f17) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_474_i6)) {
              break L1;
            }
            _474_i6 = (_474_i6).plus(_dafny.ONE);
            let _475_v40;
            _475_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(new BigNumber(861), (_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))], _444_v30),_403_v9);
            let _476_v41;
            _476_v41 = _module.D1.create_DC4(new BigNumber((_dafny.Seq.UnicodeFromString("ksjsyh")).length), (_442_v28).f17, _dafny.Map.Empty.slice().updateUnsafe((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))],_443_v29));
            let _477_v42;
            _477_v42 = _dafny.MultiSet.fromElements(_403_v9, _403_v9, (((_475_v40).contains(_476_v41)) ? ((_475_v40).get(_476_v41)) : (_403_v9)), _403_v9, _403_v9);
            _477_v42 = _477_v42;
            let _478_v43;
            _478_v43 = _dafny.Seq.of(_466_v34, _466_v34, _466_v34);
            let _index81 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
            (_466_v34)[_index81] = !(_dafny.Seq.IsPrefixOf(_478_v43, _478_v43));
            let _479_v44;
            let _nw77 = Array((new BigNumber(8)).toNumber());
            _nw77[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("fqsutwajj");
            _nw77[(_dafny.ONE).toNumber()] = _402_v8;
            _nw77[(new BigNumber(2)).toNumber()] = _402_v8;
            _nw77[(new BigNumber(3)).toNumber()] = _402_v8;
            _nw77[(new BigNumber(4)).toNumber()] = _402_v8;
            _nw77[(new BigNumber(5)).toNumber()] = _402_v8;
            _nw77[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_402_v8, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))],_397_v2)).length), new BigNumber((_402_v8).length)), _396_v1);
            _nw77[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("ibgqhjni");
            _479_v44 = _nw77;
            let _480_v45;
            _480_v45 = _dafny.Map.Empty.slice().updateUnsafe((_442_v28).f17,_479_v44);
            _480_v45 = (_480_v45).update(true, _479_v44);
            let _481_v47;
            let _nw78 = new _module.C0();
            _nw78.__ctor(!(_400_v6).equals(function () {
              let _coll29 = new _dafny.Map();
              for (const _compr_29 of (_module.__default.fm5(_405_globalState)).Keys.Elements) {
                let _482_v46 = _compr_29;
                if ((_module.__default.fm5(_405_globalState)).contains(_482_v46)) {
                  _coll29.push([(_482_v46).minus(_397_v2),(_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))]]);
                }
              }
              return _coll29;
            }()));
            _481_v47 = _nw78;
          }
        }
      }
      let _index82 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
      (_466_v34)[_index82] = (_469_v37).f17;
      let _hi1 = (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))];
      for (let _483_i7 = new BigNumber(85); _483_i7.isLessThan(_hi1); _483_i7 = _483_i7.plus(_dafny.ONE)) {
        let _484_v48;
        let _nw79 = new _module.C0();
        _nw79.__ctor((_469_v37).f17);
        _484_v48 = _nw79;
        let _485_v49;
        _485_v49 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(874)), ((_486_v8) => function (_487_i8) {
          return _486_v8;
        })(_402_v8))).length)));
        let _index83 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_472_v39).length));
        (_472_v39)[_index83] = _485_v49;
        let _index84 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_472_v39).length));
        (_472_v39)[_index84] = _485_v49;
        let _488_v50;
        _488_v50 = _dafny.Map.Empty.slice().updateUnsafe(_397_v2,_400_v6);
        let _489_v51;
        _489_v51 = _module.D0.create_DC0(new BigNumber(863), (_469_v37).f17, _402_v8, _488_v50);
        let _490_v52;
        _490_v52 = _module.D1.create_DC3(_408_v10, _397_v2, _489_v51);
        let _491_v53;
        _491_v53 = _dafny.Set.fromElements(_490_v52, _module.D1.create_DC3(_408_v10, _483_i7, _489_v51), _490_v52);
        let _index85 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length));
        (_403_v9)[_index85] = _module.__default.safeModuloInt(_483_i7, new BigNumber(((_491_v53).Intersect(_dafny.Set.fromElements(_490_v52))).length));
        _467_v35 = (_467_v35).update((_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))], (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]);
      }
      if (_dafny.Seq.IsPrefixOf(_402_v8, _402_v8)) {
        _397_v2 = (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))];
        let _492_v54;
        _492_v54 = _module.D1.create_DC2(_396_v1);
        let _493_v55;
        _493_v55 = _module.D1.create_DC4(_397_v2, (_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))], _444_v30);
        let _494_v56;
        _494_v56 = _dafny.Map.Empty.slice().updateUnsafe(_442_v28,_465_v33);
        let _495_v58;
        _495_v58 = _dafny.Set.fromElements(new BigNumber((function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-794), new BigNumber(807))) {
            let _496_v57 = _compr_30;
            if (((new BigNumber(-794)).isLessThanOrEqualTo(_496_v57)) && ((_496_v57).isLessThan(new BigNumber(807)))) {
              _coll30.push([_module.__default.safeDivisionInt(_496_v57, _397_v2),_396_v1]);
            }
          }
          return _coll30;
        }()).length));
        let _497_v59;
        let _nw80 = Array((new BigNumber(20)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = _module.__default.fm6(false, _398_v4, (_442_v28).f17, _405_globalState);
        _nw80[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((_493_v55).dtor_cf10, (_442_v28).f17, false, (_442_v28).f17, true);
        _nw80[(new BigNumber(2)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(3)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_module.__default.fm7((_492_v54).dtor_cf5, _405_globalState), _399_v5));
        _nw80[(new BigNumber(5)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(6)).toNumber()] = _module.__default.fm6(true, (_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))], (_469_v37).f17, _405_globalState);
        _nw80[(new BigNumber(7)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(8)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(9)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(10)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(11)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(12)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(13)).toNumber()] = (_465_v33).update((_469_v37).f17, _module.__default.abs(_397_v2));
        _nw80[(new BigNumber(14)).toNumber()] = (((_494_v56).contains(_469_v37)) ? ((_494_v56).get(_469_v37)) : ((_465_v33).update((_442_v28).f17, _module.__default.abs(new BigNumber((_495_v58).length)))));
        _nw80[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements(true, (_469_v37).f17);
        _nw80[(new BigNumber(16)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements((_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))], (_442_v28).f17);
        _nw80[(new BigNumber(18)).toNumber()] = _465_v33;
        _nw80[(new BigNumber(19)).toNumber()] = _465_v33;
        _497_v59 = _nw80;
        let _index86 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_497_v59).length));
        (_497_v59)[_index86] = _465_v33;
        let _index87 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_497_v59).length));
        let _rhs89 = _492_v54;
        let _rhs90 = (_dafny.MultiSet.fromElements((_469_v37).f17)).Union(_465_v33);
        let _lhs61 = _497_v59;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_497_v59).length));
        _492_v54 = _rhs89;
        _lhs61[_lhs62] = _rhs90;
        let _498_v60;
        let _nw81 = new _module.C0();
        _nw81.__ctor(((_398_v4) ? (!((_469_v37).f17)) : ((_469_v37).f17)));
        _498_v60 = _nw81;
        let _499_v61;
        _499_v61 = _dafny.Set.fromElements(_399_v5);
        if (((_499_v61).Intersect(_499_v61)).IsProperSubsetOf(_499_v61)) {
          let _500_v62;
          let _init9 = function (_501_i9) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          };
          let _nw82 = Array((new BigNumber(25)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw82.length); _i0_9++) {
            _nw82[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _500_v62 = _nw82;
          let _rhs91 = _500_v62;
          let _rhs92 = (_498_v60).f17;
          _500_v62 = _rhs91;
          _398_v4 = _rhs92;
          let _502_v63;
          let _nw83 = Array((new BigNumber(15)).toNumber());
          _502_v63 = _nw83;
          let _503_v64;
          _503_v64 = _dafny.Set.fromElements(_502_v63);
          let _index88 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          let _rhs93 = _403_v9;
          let _rhs94 = (_503_v64).IsProperSubsetOf(_dafny.Set.fromElements(_502_v63));
          let _rhs95 = _module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_469_v37).f17, _405_globalState);
          let _rhs96 = _442_v28;
          let _lhs63 = _466_v34;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          _403_v9 = _rhs93;
          _lhs63[_lhs64] = _rhs94;
          _398_v4 = _rhs95;
          _442_v28 = _rhs96;
          let _504_v65;
          _504_v65 = _dafny.Map.Empty.slice().updateUnsafe(_397_v2,_400_v6);
          let _505_v66;
          _505_v66 = _module.D0.create_DC0(_397_v2, (_498_v60).f17, _402_v8, _504_v65);
          let _506_v67;
          _506_v67 = _module.D1.create_DC3(_408_v10, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _505_v66);
          _506_v67 = _module.__default.fm8((_442_v28).f17, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _405_globalState);
          let _507_v68;
          _507_v68 = _dafny.Map.Empty.slice().updateUnsafe((_469_v37).f17,_dafny.Seq.UnicodeFromString("fyk"));
          _507_v68 = (_507_v68).update((_469_v37).f17, _402_v8);
          let _508_v69;
          _508_v69 = _dafny.Seq.of(_module.D1.create_DC3((_408_v10).update(_397_v2, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]), (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _505_v66), _506_v67, _module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_495_v58, _495_v58)).length),new BigNumber(214)), _397_v2, _module.D0.create_DC0(new BigNumber((_401_v7).length), _398_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), ((_509_v1) => function (_510_i10) {
  return _509_v1;
})(_396_v1)), _504_v65)));
          let _511_v70;
          _511_v70 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_506_v67, _506_v67), _508_v69, _508_v69, _508_v69, _module.__default.fm9(_405_globalState));
          let _index89 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index89] = (_511_v70).IsProperSubsetOf(_511_v70);
        } else {
          let _512_v71;
          _512_v71 = _dafny.Seq.of(_module.__default.safeDivisionInt((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _397_v2), _397_v2);
          let _index90 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length));
          (_403_v9)[_index90] = (_dafny.ZERO).minus((_512_v71)[_module.__default.safeIndex(((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).plus((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]), new BigNumber((_512_v71).length))]);
          let _513_v72;
          let _nw84 = new _module.C0();
          _nw84.__ctor((_401_v7).IsDisjointFrom(_401_v7));
          _513_v72 = _nw84;
          let _514_v73;
          _514_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-377),_400_v6);
          let _515_v74;
          _515_v74 = _module.D0.create_DC0(new BigNumber((_402_v8).length), (_469_v37).f17, _402_v8, _514_v73);
          let _516_v75;
          _516_v75 = _module.D1.create_DC3(_408_v10, ((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).multipliedBy(_397_v2), _515_v74);
          let _pat_let_tv21 = _512_v71;
          let _pat_let_tv22 = _512_v71;
          let _pat_let_tv23 = _408_v10;
          let _pat_let_tv24 = _403_v9;
          let _pat_let_tv25 = _403_v9;
          _516_v75 = function (_pat_let16_0) {
            return function (_517_dt__update__tmp_h0) {
              return function (_pat_let17_0) {
                return function (_519_dt__update_hcf6_h0) {
                  return _module.D1.create_DC3(_519_dt__update_hcf6_h0, (_517_dt__update__tmp_h0).dtor_cf7, (_517_dt__update__tmp_h0).dtor_cf8);
                }(_pat_let17_0);
              }(function () {
                let _coll31 = new _dafny.Map();
                for (const _compr_31 of (_pat_let_tv21).Elements) {
                  let _518_v76 = _compr_31;
                  if (_dafny.Seq.contains(_pat_let_tv22, _518_v76)) {
                    _coll31.push([(_518_v76).multipliedBy((_pat_let_tv25)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_pat_let_tv24).length))]),new BigNumber((_pat_let_tv23).length)]);
                  }
                }
                return _coll31;
              }());
            }(_pat_let16_0);
          }(_516_v75);
          let _520_v77;
          _520_v77 = _dafny.Map.Empty.slice().updateUnsafe((_442_v28).f17,(_469_v37).f17);
          let _521_v78;
          _521_v78 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("fh"),(((_520_v77).contains((_469_v37).f17)) ? ((_520_v77).get((_469_v37).f17)) : (_module.__default.fm4(_397_v2, _module.__default.fm4(_397_v2, (_513_v72).f17, _405_globalState), _405_globalState))));
          let _index91 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index91] = (((_521_v78).contains(_dafny.Seq.Concat(_402_v8, _402_v8))) ? ((_521_v78).get(_dafny.Seq.Concat(_402_v8, _402_v8))) : (((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).isLessThanOrEqualTo(new BigNumber((_399_v5).length))));
          (_405_globalState).f16 = (_dafny.ZERO).minus((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]);
        }
        let _index92 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length));
        (_403_v9)[_index92] = ((true) ? ((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]) : (_397_v2));
      } else {
        _403_v9 = _403_v9;
        if ((_401_v7).IsProperSubsetOf(_401_v7)) {
          (_405_globalState).f16 = (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))];
          _398_v4 = (_442_v28).f17;
          let _522_v79;
          let _nw85 = Array((new BigNumber(21)).toNumber());
          _522_v79 = _nw85;
          let _index93 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_522_v79).length));
          (_522_v79)[_index93] = _442_v28;
          let _index94 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_522_v79).length));
          (_522_v79)[_index94] = _442_v28;
          _397_v2 = _module.__default.safeDivisionInt((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (((_467_v35).contains(_module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_442_v28).f17, _405_globalState))) ? ((_467_v35).get(_module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_442_v28).f17, _405_globalState))) : (_module.__default.fm0((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _408_v10, _405_globalState))));
          (_405_globalState).f6 = _module.__default.safeModuloInt((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_dafny.ZERO).minus(new BigNumber((_399_v5).length)));
        } else {
          let _index95 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index95] = (_442_v28).f17;
          _396_v1 = _396_v1;
          let _523_v80;
          let _nw86 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
          _523_v80 = _nw86;
          let _index96 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_523_v80).length));
          (_523_v80)[_index96] = _443_v29;
          let _index97 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_523_v80).length));
          let _rhs97 = _dafny.Seq.update(_402_v8, _module.__default.safeIndex((new BigNumber(-462)).minus(_397_v2), new BigNumber((_402_v8).length)), _396_v1);
          let _rhs98 = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_403_v9,_396_v1)).length), _module.__default.safeModuloInt(_397_v2, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]));
          let _rhs99 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_397_v2));
          let _rhs100 = _442_v28;
          let _rhs101 = _403_v9;
          let _lhs65 = _523_v80;
          let _lhs66 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_523_v80).length));
          let _lhs67 = _405_globalState;
          _402_v8 = _rhs97;
          _397_v2 = _rhs98;
          _lhs65[_lhs66] = _rhs99;
          _442_v28 = _rhs100;
          _lhs67.f14 = _rhs101;
          let _index98 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index98] = (_469_v37).f17;
          let _index99 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length));
          (_403_v9)[_index99] = _397_v2;
        }
        if ((_399_v5)[_module.__default.safeIndex((((_408_v10).contains((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))])) ? ((_408_v10).get((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))])) : (_397_v2)), new BigNumber((_399_v5).length))]) {
          let _524_v81;
          _524_v81 = _dafny.Map.Empty.slice().updateUnsafe(_402_v8,_443_v29);
          let _525_v82;
          let _nw87 = Array((new BigNumber(27)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = _443_v29;
          _nw87[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_397_v2);
          _nw87[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]);
          _nw87[(new BigNumber(3)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(4)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(5)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(6)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(7)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(8)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]));
          _nw87[(new BigNumber(10)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(11)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(12)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(13)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(14)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(15)).toNumber()] = (((_524_v81).contains(_402_v8)) ? ((_524_v81).get(_402_v8)) : (_443_v29));
          _nw87[(new BigNumber(16)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(17)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(18)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(19)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(20)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(21)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(22)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(23)).toNumber()] = (((_444_v30).contains((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))])) ? ((_444_v30).get((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))])) : (_443_v29));
          _nw87[(new BigNumber(24)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(25)).toNumber()] = _443_v29;
          _nw87[(new BigNumber(26)).toNumber()] = _443_v29;
          _525_v82 = _nw87;
          let _526_v83;
          _526_v83 = _dafny.Map.Empty.slice().updateUnsafe(_397_v2,_dafny.Map.Empty.slice().updateUnsafe(_397_v2,_398_v4));
          let _527_v84;
          _527_v84 = _module.D0.create_DC0((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_442_v28).f17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), ((_528_v1) => function (_529_i11) {
  return _528_v1;
})(_396_v1)), _526_v83);
          let _530_v85;
          _530_v85 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(_module.D1.create_DC3(_408_v10, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _527_v84)),_469_v37);
          let _531_v86;
          _531_v86 = _dafny.Seq.of((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], new BigNumber((_530_v85).length), (_dafny.ZERO).minus((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]));
          let _532_v87;
          _532_v87 = _dafny.Map.Empty.slice().updateUnsafe(_531_v86,(_442_v28).f17);
          let _533_v88;
          let _534_v89;
          let _535_v90;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = _module.__default.m0(_525_v82, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_397_v2).isLessThan(new BigNumber((_532_v87).length)), !((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).isEqualTo(_module.__default.fm0(_397_v2, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-732),new BigNumber(-173)), _405_globalState)), _405_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _533_v88 = _out0;
          _534_v89 = _out1;
          _535_v90 = _out2;
          _399_v5 = _399_v5;
          let _536_v91;
          let _537_v92;
          let _538_v93;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = _module.__default.m0(_525_v82, _533_v88, !((new BigNumber(690)).isLessThanOrEqualTo((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))])), (_469_v37).f17, _405_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _536_v91 = _out3;
          _537_v92 = _out4;
          _538_v93 = _out5;
          let _539_v94;
          _539_v94 = _dafny.Set.fromElements(_533_v88);
          let _540_v95;
          _540_v95 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(611)), ((_541_v1) => function (_542_i13) {
            return _541_v1;
          })(_396_v1)));
          let _543_v96;
          _543_v96 = _dafny.Map.Empty.slice().updateUnsafe((_539_v94).Intersect(_539_v94),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_544_v84) => function (_545_i12) {
            return (_544_v84).dtor_cf2;
          })(_527_v84)), _540_v95));
          (_405_globalState).f16 = new BigNumber(((((_543_v96).contains((_539_v94).Intersect(_539_v94))) ? ((_543_v96).get((_539_v94).Intersect(_539_v94))) : (_540_v95))).length);
          let _pat_let_tv26 = _445_v31;
          _446_v32 = function (_pat_let18_0) {
            return function (_546_dt__update__tmp_h1) {
              return function (_pat_let19_0) {
                return function (_547_dt__update_hcf12_h0) {
                  return _module.D1.create_DC5(_547_dt__update_hcf12_h0);
                }(_pat_let19_0);
              }(_pat_let_tv26);
            }(_pat_let18_0);
          }(_446_v32);
        } else {
          _397_v2 = new BigNumber(781);
          let _548_v97;
          let _nw88 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
          _548_v97 = _nw88;
          let _549_v98;
          _549_v98 = _dafny.Seq.of(_442_v28, _442_v28);
          let _550_v99;
          let _551_v100;
          let _552_v101;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = _module.__default.m0(_548_v97, (_dafny.ZERO).minus(((!((_442_v28).f17)) ? (new BigNumber((_549_v98).length)) : (_397_v2))), (_442_v28).f17, _module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], (_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))], _405_globalState), _405_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _550_v99 = _out6;
          _551_v100 = _out7;
          _552_v101 = _out8;
          _398_v4 = (_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))];
          _442_v28 = (((_442_v28).f17) ? (_469_v37) : (_469_v37));
          let _553_v102;
          _553_v102 = _dafny.Set.fromElements(_442_v28);
          let _554_v103;
          _554_v103 = _dafny.Map.Empty.slice().updateUnsafe(_553_v102,(_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))]);
          let _555_v104;
          _555_v104 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(new BigNumber((_554_v103).length), (_469_v37).f17, _405_globalState),_466_v34);
          let _556_v105;
          _556_v105 = _dafny.Map.Empty.slice().updateUnsafe(_549_v98,_466_v34);
          _555_v104 = (_555_v104).update(_module.__default.fm4(_397_v2, (_442_v28).f17, _405_globalState), (((_556_v105).contains(_549_v98)) ? ((_556_v105).get(_549_v98)) : (_466_v34)));
        }
        (_405_globalState).f11 = _397_v2;
        if ((true) || (_dafny.Seq.IsPrefixOf(_402_v8, _402_v8))) {
          _398_v4 = (_399_v5)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_401_v7).length))), new BigNumber((_399_v5).length))];
          let _index100 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index100] = !((_469_v37).f17);
          let _557_v106;
          _557_v106 = _dafny.Map.Empty.slice().updateUnsafe((((_467_v35).contains((_442_v28).f17)) ? ((_467_v35).get((_442_v28).f17)) : (_397_v2)),_469_v37);
          let _558_v107;
          _558_v107 = _dafny.Seq.of(_557_v106);
          let _index101 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index101] = !(_dafny.Seq.contains(_dafny.Seq.Concat(_558_v107, _558_v107), (_557_v106).Merge(_557_v106)));
          let _559_v108;
          let _nw89 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
          _559_v108 = _nw89;
          let _560_v109;
          let _561_v110;
          let _562_v111;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = _module.__default.m0(_559_v108, new BigNumber((_dafny.Set.fromElements((_469_v37).f17)).length), ((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).isLessThan((((_408_v10).contains(_397_v2)) ? ((_408_v10).get(_397_v2)) : (new BigNumber((_dafny.Seq.of((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))])).length)))), (_469_v37).f17, _405_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _560_v109 = _out9;
          _561_v110 = _out10;
          _562_v111 = _out11;
          let _563_v112;
          _563_v112 = _dafny.Map.Empty.slice().updateUnsafe(_465_v33,(_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]);
          let _564_v115;
          _564_v115 = _dafny.Seq.of(_465_v33);
          let _565_v116;
          let _nw90 = Array((new BigNumber(8)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _563_v112;
          _nw90[(_dafny.ONE).toNumber()] = _563_v112;
          _nw90[(new BigNumber(2)).toNumber()] = (_563_v112).Merge(_563_v112);
          _nw90[(new BigNumber(3)).toNumber()] = function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_563_v112).Keys.Elements) {
              let _566_v113 = _compr_32;
              if ((_563_v112).contains(_566_v113)) {
                _coll32.push([_566_v113,new BigNumber(560)]);
              }
            }
            return _coll32;
          }();
          _nw90[(new BigNumber(4)).toNumber()] = _module.__default.fm10((_469_v37).f17, _405_globalState);
          _nw90[(new BigNumber(5)).toNumber()] = function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of (_dafny.MultiSet.FromArray(_564_v115)).Elements) {
              let _567_v114 = _compr_33;
              if ((_dafny.MultiSet.FromArray(_564_v115)).contains(_567_v114)) {
                _coll33.push([_567_v114,(_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]]);
              }
            }
            return _coll33;
          }();
          _nw90[(new BigNumber(6)).toNumber()] = (_module.__default.fm10((_442_v28).f17, _405_globalState)).Merge(_563_v112);
          _nw90[(new BigNumber(7)).toNumber()] = _563_v112;
          _565_v116 = _nw90;
          let _index102 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_565_v116).length));
          (_565_v116)[_index102] = _563_v112;
          let _index103 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_565_v116).length));
          (_565_v116)[_index103] = _563_v112;
        } else {
          let _568_v117;
          _568_v117 = _dafny.Seq.of(_466_v34);
          _398_v4 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_466_v34), _568_v117), _568_v117);
          _398_v4 = (_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))];
          let _index104 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          let _rhs102 = _module.__default.fm4((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], _module.__default.fm4(_397_v2, (_442_v28).f17, _405_globalState), _405_globalState);
          let _rhs103 = _469_v37;
          let _rhs104 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_397_v2), (_dafny.ZERO).minus(((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).minus(_397_v2)));
          let _lhs68 = _466_v34;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          let _lhs70 = _405_globalState;
          _lhs68[_lhs69] = _rhs102;
          _442_v28 = _rhs103;
          _lhs70.f6 = _rhs104;
          let _569_v118;
          let _nw91 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
          _569_v118 = _nw91;
          let _570_v119;
          _570_v119 = _dafny.Set.fromElements(new BigNumber((_401_v7).length), (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]);
          let _index105 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_569_v118).length));
          (_569_v118)[_index105] = _570_v119;
          let _index106 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_569_v118).length));
          (_569_v118)[_index106] = (_570_v119).Union(_570_v119);
          let _index107 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length));
          (_466_v34)[_index107] = ((_dafny.MultiSet.FromArray(_399_v5)).Difference(_465_v33)).IsSubsetOf(_465_v33);
        }
      }
      let _hi2 = _397_v2;
      for (let _571_i14 = (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]; _571_i14.isLessThan(_hi2); _571_i14 = _571_i14.plus(_dafny.ONE)) {
        _398_v4 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("to"), ((_module.__default.fm4((_dafny.ZERO).minus(_571_i14), true, _405_globalState)) ? (_402_v8) : (_402_v8)));
        let _572_v120;
        let _init10 = ((_573_v9) => function (_574_i15) {
          return _dafny.MultiSet.fromElements((_573_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_573_v9).length))]);
        })(_403_v9);
        let _nw92 = Array((new BigNumber(16)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw92.length); _i0_10++) {
          _nw92[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _572_v120 = _nw92;
        let _575_v121;
        let _576_v122;
        let _577_v123;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector4 = _module.__default.m0(_572_v120, (_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))], !((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]).isEqualTo((_403_v9)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_403_v9).length))]), _dafny.Seq.contains(_402_v8, new _dafny.CodePoint('t'.codePointAt(0))), _405_globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _575_v121 = _out12;
        _576_v122 = _out13;
        _577_v123 = _out14;
        let _578_v124;
        let _579_v125;
        let _580_v126;
        let _out15;
        let _out16;
        let _out17;
        let _outcollector5 = _module.__default.m0(_572_v120, _575_v121, (_401_v7).IsDisjointFrom(_dafny.Set.fromElements((_466_v34)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_466_v34).length))])), (_442_v28).f17, _405_globalState);
        _out15 = _outcollector5[0];
        _out16 = _outcollector5[1];
        _out17 = _outcollector5[2];
        _578_v124 = _out15;
        _579_v125 = _out16;
        _580_v126 = _out17;
        let _581_v127;
        let _582_v128;
        let _583_v129;
        let _out18;
        let _out19;
        let _out20;
        let _outcollector6 = _module.__default.m0(_572_v120, _module.__default.safeModuloInt(new BigNumber((_402_v8).length), _571_i14), (_442_v28).f17, (_442_v28).f17, _405_globalState);
        _out18 = _outcollector6[0];
        _out19 = _outcollector6[1];
        _out20 = _outcollector6[2];
        _581_v127 = _out18;
        _582_v128 = _out19;
        _583_v129 = _out20;
      }
      process.stdout.write(_dafny.toString(_396_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_397_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_398_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_399_v5, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_400_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_401_v7).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_402_v8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_403_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_405_globalState).f2).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber(514)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_405_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_405_globalState).f7).equals(_dafny.MultiSet.fromElements(new BigNumber(514), new BigNumber(-1), new BigNumber(2), new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_405_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState.f14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_405_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_405_globalState.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_408_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),new BigNumber(514)).updateUnsafe(new BigNumber(-514),new BigNumber(514)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_442_v28).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_443_v29).equals(_dafny.MultiSet.fromElements(new BigNumber(514)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_444_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.MultiSet.fromElements(new BigNumber(76))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_445_v31).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_445_v31).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_445_v31).dtor_cf11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.MultiSet.fromElements(new BigNumber(76))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_446_v32).dtor_cf12).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_446_v32).dtor_cf12).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_446_v32).dtor_cf12).dtor_cf11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.MultiSet.fromElements(new BigNumber(76))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_465_v33).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_466_v34)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_466_v34)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_467_v35).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_468_v36, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),new BigNumber(514)).updateUnsafe(new BigNumber(-514),new BigNumber(514)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),new BigNumber(514)).updateUnsafe(new BigNumber(-514),new BigNumber(514))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_469_v37).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_470_v38).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.Seq.UnicodeFromString("frrp")).updateUnsafe(new BigNumber(2),_dafny.Seq.UnicodeFromString("frrp")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[_dafny.ONE], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(7)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(8)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(9)], _dafny.Seq.of(new BigNumber(-874)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(10)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(11)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(12)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_472_v39)[new BigNumber(13)], _dafny.Seq.of(new BigNumber(3355), new BigNumber(514), new BigNumber(129)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_474_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + this.cf2.toVerbatimString(true) + ", " + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO, false, _dafny.Seq.UnicodeFromString(""), _dafny.Map.Empty);
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
    static create_DC2(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC3(cf6, cf7, cf8) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC4(cf9, cf10, cf11) {
      let $dt = new D1(2);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC1(cf4) {
      let $dt = new D1(3);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf12) {
      let $dt = new D1(4);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC1() { return this.$tag === 3; }
    get is_DC5() { return this.$tag === 4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 4) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC7(cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf16) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D2(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC6(cf13) {
      let $dt = new D2(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf16 === other.cf16;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf13 === other.cf13;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC11() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC10(cf18) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf18) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11();
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
    static create_DC13(cf20, cf21, cf22) {
      let $dt = new D4(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf23, cf24, cf25) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC15(cf26, cf27) {
      let $dt = new D4(2);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D4(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC16(cf28) {
      let $dt = new D4(4);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get is_DC16() { return this.$tag === 4; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(false, _dafny.ZERO, []);
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
    static create_DC18(cf30, cf31) {
      let $dt = new D5(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC17(cf29) {
      let $dt = new D5(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC19(cf32) {
      let $dt = new D5(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC18(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC21(cf34, cf35, cf36, cf37) {
      let $dt = new D6(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC20(cf33) {
      let $dt = new D6(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC21(false, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f6 = _dafny.ZERO;
      this.f11 = _dafny.ZERO;
      this.f14 = [];
      this.f16 = _dafny.ZERO;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = _dafny.Map.Empty;
      this._f3 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f7 = _dafny.MultiSet.Empty;
      this._f8 = false;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f13 = false;
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this).f16 = f16;
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
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f15() {
      let _this = this;
      return _this._f15;
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
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), function (_584_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jcohubbi"), _dafny.Seq.UnicodeFromString("bfaki")));
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
