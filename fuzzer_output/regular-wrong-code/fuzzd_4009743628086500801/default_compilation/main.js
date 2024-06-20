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
    static fm3(globalState) {
      return ((_dafny.Set.fromElements(_dafny.Set.fromElements(false, true), _dafny.Set.fromElements(true, true))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(false, true), _dafny.Set.fromElements(false), _dafny.Set.fromElements(true, false), _dafny.Set.fromElements(false)))).Difference((_dafny.Set.fromElements(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false))));
    };
    static fm4(p0, p1, p2, globalState) {
      return _dafny.Seq.UnicodeFromString("ed");
    };
    static fm5(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("knr"),_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), function (_0_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length), new BigNumber(771)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cvrf"),_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.MultiSet.fromElements(new BigNumber(817))).Elements) {
          let _1_v0 = _compr_0;
          if ((_dafny.MultiSet.fromElements(new BigNumber(817))).contains(_1_v0)) {
            _coll0.push([(_1_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(true)).length)),new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality())]);
          }
        }
        return _coll0;
      }()).length))).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("dujbkapsg"),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-864)), function (_2_i1) {
        return new BigNumber(733);
      })));
    };
    static fm8(p0, globalState) {
      return !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uh"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-759)), function (_3_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("snaam"), _dafny.Seq.UnicodeFromString("uupaiodb")));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), function (_4_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("bndem"));
    };
    static fm12(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), ((true) ? (_dafny.Seq.of(false)) : (_dafny.Seq.of(true))));
    };
    static fm14(p0, p1, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-535), new BigNumber(990))) {
          let _5_v0 = _compr_1;
          if (((new BigNumber(-535)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(990)))) {
            _coll1.push([(_5_v0).plus(new BigNumber(-90)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-907),new BigNumber(-667))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(521)), function (_6_i0) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            })).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), function (_7_i1) {
              return new _dafny.CodePoint('c'.codePointAt(0));
            })).length), new BigNumber((_dafny.MultiSet.fromElements(function () {
              let _coll2 = new _dafny.Map();
              for (const _compr_2 of _dafny.IntegerRange(new BigNumber(286), new BigNumber(-321))) {
                let _8_v1 = _compr_2;
                if (((new BigNumber(286)).isLessThanOrEqualTo(_8_v1)) && ((_8_v1).isLessThan(new BigNumber(-321)))) {
                  _coll2.push([_module.__default.safeDivisionInt(_8_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("srsgjqmes")).length)))).length))).length)),new BigNumber(32)]);
                }
              }
              return _coll2;
            }())).cardinality()))).length)]);
          }
        }
        return _coll1;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-841),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber(115))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D15.create_DC33(_dafny.Seq.UnicodeFromString("rcx"))).dtor_cf54).length),new BigNumber(692))))).length);
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true));
    };
    static fm16(p0, globalState) {
      return ((_dafny.Set.fromElements(false, !(!(true)))).Difference(_dafny.Set.fromElements(true, false))).Difference(_dafny.Set.fromElements(true, false));
    };
    static fm19(p0, p1, globalState) {
      let _source0 = _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hvkyd")).length),new BigNumber((_dafny.Seq.of(new BigNumber(609), new BigNumber(-451))).length)),_dafny.Set.fromElements(!(true))));
      if (_source0.is_DC19) {
        let _9___mcc_h0 = (_source0).cf29;
        let _10___mcc_h1 = (_source0).cf30;
        let _11_cf30 = _10___mcc_h1;
        let _12_cf29 = _9___mcc_h0;
        return (_dafny.ZERO).minus(_12_cf29);
      } else if (_source0.is_DC20) {
        let _13___mcc_h2 = (_source0).cf31;
        let _14___mcc_h3 = (_source0).cf32;
        let _15___mcc_h4 = (_source0).cf33;
        let _16_cf33 = _15___mcc_h4;
        let _17_cf32 = _14___mcc_h3;
        let _18_cf31 = _13___mcc_h2;
        return new BigNumber(((_dafny.Set.fromElements(_dafny.Seq.of(_dafny.ZERO), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gynrfk")).length)), new BigNumber((_dafny.Set.fromElements(!(_16_cf33))).length), new BigNumber((_dafny.Seq.UnicodeFromString("wvp")).length)))).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), function (_19_i0) {
          return new BigNumber(432);
        }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_18_cf31)).length), new BigNumber((_dafny.Seq.UnicodeFromString("gynpbladr")).length))))).length);
      } else {
        let _20___mcc_h5 = (_source0).cf28;
        let _21_cf28 = _20___mcc_h5;
        return (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("hyoldvcp")).length), new BigNumber(978), new BigNumber(255), new BigNumber((_dafny.Seq.UnicodeFromString("iflkrw")).length), new BigNumber((_dafny.Seq.of(true, true)).length))).cardinality())).plus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
      }
    };
    static fm20(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("iqunqvo")).length)),new BigNumber((_dafny.Seq.of(new BigNumber(-829))).length))).length))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), function (_22_i0) {
        return _module.D1.create_DC4(new BigNumber(82), new _dafny.CodePoint('y'.codePointAt(0)), false, new BigNumber(805));
      })).length))).Intersect(_dafny.Set.fromElements(new BigNumber(314)))).Difference(function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-820), new BigNumber(382))) {
          let _23_v0 = _compr_3;
          if (((new BigNumber(-820)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(382)))) {
            _coll3.add(_module.__default.safeDivisionInt(_23_v0, new BigNumber(599)));
          }
        }
        return _coll3;
      }());
    };
    static fm22(p0, p1, globalState) {
      return (_module.D15.create_DC33(_dafny.Seq.UnicodeFromString("xebgvpgac"))).dtor_cf54;
    };
    static fm23(p0, globalState) {
      return _dafny.Seq.of(new BigNumber(932), ((true) ? (new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(563)))).length))).length)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false)).length)))).length),new BigNumber(265))).length))), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Seq.of(false, false)).length))).length), new BigNumber(-658), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC21(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false))),false)).length));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return _module.__default.safeModuloInt(new BigNumber(918), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vkmbe"), _dafny.Seq.UnicodeFromString("r"))).length));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("vogq"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("q"), _dafny.Seq.UnicodeFromString("y")));
    };
    static fm26(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, true, false, (_module.D1.create_DC4(new BigNumber((_dafny.Set.fromElements((_module.D15.create_DC34(_dafny.Seq.UnicodeFromString("ghf"), new BigNumber(164), false, _module.D10.create_DC22(), new BigNumber(693))).dtor_cf57, true)).length), new _dafny.CodePoint('c'.codePointAt(0)), true, new BigNumber(228))).dtor_cf8), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)));
    };
    static fm27(globalState) {
      return (((!((_module.D15.create_DC32(false)).dtor_cf53)) ? (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jac"))) : (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hwf"))))).Union((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("en"))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sicipu"))));
    };
    static fm28(p0, p1, globalState) {
      if (!(new BigNumber(216)).isEqualTo(new BigNumber(-12))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(!(true),(_module.D15.create_DC33(_dafny.Seq.UnicodeFromString("fbkybdo"))).dtor_cf54)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(!(false))))),_dafny.Seq.UnicodeFromString("vudhbc")));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.UnicodeFromString("jpc"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("fgwlc")));
      }
    };
    static fm29(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("fykfcgc")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-242),new BigNumber(517)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("iglpstssc")).length),new BigNumber(-663)));
    };
    static fm30(p0, p1, p2, globalState) {
      let _source1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("bmoihfr"))).cardinality()))).length),new BigNumber(715)),true);
      let _24___mcc_h0 = _source1;
      let _25_cf18 = _24___mcc_h0;
      return new _dafny.CodePoint('e'.codePointAt(0));
    };
    static fm31(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(-55), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(585),new BigNumber(832))).length), (_dafny.ZERO).minus(new BigNumber(-24)), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())), new BigNumber(249))).Union(_dafny.MultiSet.fromElements(new BigNumber(697)))).Union((_dafny.MultiSet.fromElements(new BigNumber(747), new BigNumber((_dafny.Seq.of(false, true)).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(34), new BigNumber(873), new BigNumber(-285), new BigNumber((_dafny.Set.fromElements(true)).length))));
    };
    static fm32(p0, globalState) {
      return function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), function (_26_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(164),new BigNumber(-392));
        }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ykxkeccls")).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xrtrgqvhp")).length))), function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-447), new BigNumber(500))) {
            let _27_v1 = _compr_5;
            if (((new BigNumber(-447)).isLessThanOrEqualTo(_27_v1)) && ((_27_v1).isLessThan(new BigNumber(500)))) {
              _coll5.push([(_27_v1).multipliedBy(new BigNumber(226)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)]);
            }
          }
          return _coll5;
        }()))).Elements) {
          let _28_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), function (_26_i0) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(164),new BigNumber(-392));
          }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ykxkeccls")).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xrtrgqvhp")).length))), function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-447), new BigNumber(500))) {
              let _27_v1 = _compr_6;
              if (((new BigNumber(-447)).isLessThanOrEqualTo(_27_v1)) && ((_27_v1).isLessThan(new BigNumber(500)))) {
                _coll6.push([(_27_v1).multipliedBy(new BigNumber(226)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)]);
              }
            }
            return _coll6;
          }())), _28_v0)) {
            _coll4.push([_28_v0,_dafny.Set.fromElements(!(false), false, !(!(false)))]);
          }
        }
        return _coll4;
      }();
    };
    static fm33(p0, globalState) {
      return false;
    };
    static fm34(p0, p1, globalState) {
      let _source2 = _module.D1.create_DC2(false);
      if (_source2.is_DC3) {
        let _29___mcc_h0 = (_source2).cf5;
        let _30_cf5 = _29___mcc_h0;
        return _module.D1.create_DC2(false);
      } else if (_source2.is_DC4) {
        let _31___mcc_h1 = (_source2).cf6;
        let _32___mcc_h2 = (_source2).cf7;
        let _33___mcc_h3 = (_source2).cf8;
        let _34___mcc_h4 = (_source2).cf9;
        let _35_cf9 = _34___mcc_h4;
        let _36_cf8 = _33___mcc_h3;
        let _37_cf7 = _32___mcc_h2;
        let _38_cf6 = _31___mcc_h1;
        return _module.D1.create_DC2(_36_cf8);
      } else {
        let _39___mcc_h5 = (_source2).cf4;
        let _40_cf4 = _39___mcc_h5;
        return _module.D1.create_DC2(_40_cf4);
      }
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),new BigNumber(902))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new BigNumber(712)))).Keys.Elements) {
          let _41_v0 = _compr_7;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),new BigNumber(902))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new BigNumber(712)))).contains(_41_v0)) {
            _coll7.push([_41_v0,!(false)]);
          }
        }
        return _coll7;
      }();
    };
    static fm36(globalState) {
      return _module.D1.create_DC4((new BigNumber(522)).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-290))).length)), (_module.D1.create_DC4(new BigNumber(721), new _dafny.CodePoint('d'.codePointAt(0)), false, new BigNumber(427))).dtor_cf7, !(!(false)), ((true) ? (new BigNumber(579)) : (new BigNumber(602))));
    };
    static m6(globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = undefined;
      let _42_v0;
      let _nw0 = new _module.C8();
      _nw0.__ctor();
      _42_v0 = _nw0;
      let _nw1 = new _module.C8();
      _nw1.__ctor();
      _42_v0 = _nw1;
      let _43_v1;
      let _nw2 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
      _43_v1 = _nw2;
      let _44_v2;
      _44_v2 = false;
      let _45_v3;
      _45_v3 = _dafny.Map.Empty.slice().updateUnsafe(_44_v2,new _dafny.CodePoint('i'.codePointAt(0)));
      let _46_v4;
      _46_v4 = new _dafny.CodePoint('v'.codePointAt(0));
      let _47_v5;
      _47_v5 = _dafny.MultiSet.fromElements(_44_v2);
      let _48_v6;
      _48_v6 = _dafny.Map.Empty.slice().updateUnsafe((((_45_v3).contains(_44_v2)) ? ((_45_v3).get(_44_v2)) : (_46_v4)),_47_v5);
      let _index0 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_43_v1).length));
      (_43_v1)[_index0] = (((_48_v6).contains(_46_v4)) ? ((_48_v6).get(_46_v4)) : (_47_v5));
      let _49_v7;
      _49_v7 = _dafny.Map.Empty.slice().updateUnsafe(_44_v2,_44_v2);
      let _50_v8;
      _50_v8 = new BigNumber(94);
      let _51_v9;
      _51_v9 = _dafny.Map.Empty.slice().updateUnsafe(_50_v8,_44_v2);
      let _index1 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_43_v1).length));
      (_43_v1)[_index1] = (_42_v0).fm0((_49_v7).equals(_49_v7), (((((_51_v9).contains(_50_v8)) ? ((_51_v9).get(_50_v8)) : (false))) ? (!(_44_v2)) : (_44_v2)), _50_v8, _50_v8, globalState);
      let _52_v10;
      let _init0 = function (_53_i0) {
        return (_53_i0).multipliedBy(new BigNumber(800));
      };
      let _nw3 = Array((new BigNumber(19)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
        _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _52_v10 = _nw3;
      let _nw4 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _52_v10 = _nw4;
      _42_v0 = _42_v0;
      let _54_v11;
      _54_v11 = _dafny.Seq.UnicodeFromString("ny");
      _50_v8 = ((_44_v2) ? (new BigNumber((_dafny.Seq.update(_54_v11, _module.__default.safeIndex(_50_v8, new BigNumber((_54_v11).length)), new _dafny.CodePoint('y'.codePointAt(0)))).length)) : (_50_v8));
      let _55_v12;
      _55_v12 = _module.D4.create_DC11(_module.D4.create_DC9(_52_v10));
      let _56_v13;
      _56_v13 = _module.D4.create_DC11(_55_v12);
      let _57_v14;
      _57_v14 = _dafny.Map.Empty.slice().updateUnsafe(_56_v13,_module.__default.fm19(false, new BigNumber(961), globalState));
      let _58_v15;
      _58_v15 = _dafny.Seq.of(_50_v8);
      _44_v2 = !(((((_57_v14).contains(_56_v13)) ? ((_57_v14).get(_56_v13)) : ((_58_v15)[_module.__default.safeIndex(_50_v8, new BigNumber((_58_v15).length))]))).isLessThanOrEqualTo(_50_v8));
      r0 = _dafny.Seq.Concat(_54_v11, _54_v11);
      let _59_v16;
      let _nw5 = new _module.C8();
      _nw5.__ctor();
      _59_v16 = _nw5;
      r1 = _59_v16;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _60_v0;
      let _nw6 = Array((new BigNumber(17)).toNumber()).fill(false);
      _60_v0 = _nw6;
      let _61_globalState;
      let _nw7 = new _module.GlobalState();
      _nw7.__ctor(_60_v0, new BigNumber(-481));
      _61_globalState = _nw7;
      let _62_v1;
      _62_v1 = false;
      if (_62_v1) {
        let _63_v2;
        _63_v2 = new BigNumber(7);
        _63_v2 = new BigNumber(547);
        let _64_v3;
        _64_v3 = _dafny.MultiSet.fromElements(_62_v1, false);
        let _65_v4;
        let _nw8 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _65_v4 = _nw8;
        let _66_v5;
        let _nw9 = new _module.C6();
        _nw9.__ctor((_64_v3).update(_62_v1, _module.__default.abs(_63_v2)), _65_v4);
        _66_v5 = _nw9;
        let _67_v6;
        _67_v6 = _dafny.Seq.UnicodeFromString("yymfboxka");
        _67_v6 = _dafny.Seq.Concat(_67_v6, _dafny.Seq.Concat(_module.__default.fm4(_62_v1, _62_v1, _62_v1, _61_globalState), _67_v6));
        let _68_v7;
        _68_v7 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("kahofehgp"));
        let _69_v8;
        let _nw10 = new _module.C5();
        _nw10.__ctor((_68_v7)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("xiko")).length), new BigNumber((_68_v7).length))]);
        _69_v8 = _nw10;
        let _70_v9;
        let _nw11 = Array((new BigNumber(7)).toNumber()).fill([]);
        _70_v9 = _nw11;
        let _index2 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_70_v9).length));
        (_70_v9)[_index2] = _60_v0;
        let _index3 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_70_v9).length));
        (_70_v9)[_index3] = _60_v0;
      } else {
        let _71_v10;
        _71_v10 = new BigNumber(849);
        let _72_v11;
        _72_v11 = _dafny.Map.Empty.slice().updateUnsafe(_71_v10,_71_v10);
        let _73_v12;
        _73_v12 = _dafny.Seq.of(_62_v1, _62_v1);
        let _74_v13;
        _74_v13 = _dafny.Seq.of(_71_v10, _71_v10);
        let _75_v14;
        _75_v14 = _dafny.Map.Empty.slice().updateUnsafe(_71_v10,_62_v1);
        _72_v11 = (_72_v11).update(new BigNumber((_dafny.Seq.update(_73_v12, _module.__default.safeIndex(new BigNumber((_74_v13).length), new BigNumber((_73_v12).length)), (((_75_v14).contains(_71_v10)) ? ((_75_v14).get(_71_v10)) : (_62_v1)))).length), _71_v10);
        let _76_v15;
        _76_v15 = new _dafny.CodePoint('f'.codePointAt(0));
        let _77_v16;
        let _nw12 = Array((new BigNumber(8)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _76_v15;
        _nw12[(_dafny.ONE).toNumber()] = _76_v15;
        _nw12[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw12[(new BigNumber(3)).toNumber()] = _76_v15;
        _nw12[(new BigNumber(4)).toNumber()] = _76_v15;
        _nw12[(new BigNumber(5)).toNumber()] = _76_v15;
        _nw12[(new BigNumber(6)).toNumber()] = _76_v15;
        _nw12[(new BigNumber(7)).toNumber()] = _module.__default.fm30(_module.__default.fm22(_62_v1, (_dafny.ZERO).minus(_71_v10), _61_globalState), _71_v10, new BigNumber(581), _61_globalState);
        _77_v16 = _nw12;
        let _78_v17;
        _78_v17 = _dafny.Seq.UnicodeFromString("i");
        let _79_v18;
        _79_v18 = _dafny.Set.fromElements(_62_v1);
        let _index4 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_77_v16).length));
        (_77_v16)[_index4] = (_78_v17)[_module.__default.safeIndex(new BigNumber((_79_v18).length), new BigNumber((_78_v17).length))];
        let _index5 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_77_v16).length));
        (_77_v16)[_index5] = _76_v15;
        let _index6 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length));
        (_60_v0)[_index6] = _62_v1;
        let _80_v19;
        let _nw13 = new _module.C4();
        _nw13.__ctor(new BigNumber(232));
        _80_v19 = _nw13;
        let _index7 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length));
        let _rhs0 = _62_v1;
        let _rhs1 = _71_v10;
        let _rhs2 = _78_v17;
        let _rhs3 = _62_v1;
        let _rhs4 = _80_v19;
        let _lhs0 = _60_v0;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length));
        _62_v1 = _rhs0;
        _71_v10 = _rhs1;
        _78_v17 = _rhs2;
        _lhs0[_lhs1] = _rhs3;
        _80_v19 = _rhs4;
        let _81_v20;
        let _nw14 = Array((new BigNumber(14)).toNumber());
        _nw14[(_dafny.ZERO).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(_dafny.ONE).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(new BigNumber(2)).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(new BigNumber(3)).toNumber()] = _62_v1;
        _nw14[(new BigNumber(4)).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(new BigNumber(5)).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(new BigNumber(6)).toNumber()] = _62_v1;
        _nw14[(new BigNumber(7)).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(new BigNumber(8)).toNumber()] = (_60_v0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_60_v0).length))];
        _nw14[(new BigNumber(9)).toNumber()] = _62_v1;
        _nw14[(new BigNumber(10)).toNumber()] = _62_v1;
        _nw14[(new BigNumber(11)).toNumber()] = true;
        _nw14[(new BigNumber(12)).toNumber()] = _62_v1;
        _nw14[(new BigNumber(13)).toNumber()] = _62_v1;
        _81_v20 = _nw14;
        let _82_v21;
        _82_v21 = _dafny.Seq.of(_81_v20, _81_v20);
        let _83_v22;
        let _nw15 = Array((new BigNumber(26)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _81_v20;
        _nw15[(_dafny.ONE).toNumber()] = _81_v20;
        _nw15[(new BigNumber(2)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(3)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(4)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(5)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(6)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(7)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(8)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(9)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(10)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(11)).toNumber()] = (_82_v21)[_module.__default.safeIndex(_71_v10, new BigNumber((_82_v21).length))];
        _nw15[(new BigNumber(12)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(13)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(14)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(15)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(16)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(17)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(18)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(19)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(20)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(21)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(22)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(23)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(24)).toNumber()] = _81_v20;
        _nw15[(new BigNumber(25)).toNumber()] = _81_v20;
        _83_v22 = _nw15;
        let _84_v23;
        let _85_v24;
        let _out0;
        let _out1;
        let _outcollector0 = (_80_v19).m0(_83_v22, _61_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _84_v23 = _out0;
        _85_v24 = _out1;
        let _86_v25;
        let _nw16 = new _module.C4();
        _nw16.__ctor(_71_v10);
        _86_v25 = _nw16;
        _86_v25 = _86_v25;
      }
      let _87_v26;
      _87_v26 = new BigNumber(542);
      let _88_v27;
      _88_v27 = _dafny.Map.Empty.slice().updateUnsafe(_87_v26,_87_v26);
      let _89_v28;
      _89_v28 = _dafny.Seq.of(_87_v26);
      let _90_v29;
      _90_v29 = _dafny.Seq.of((_89_v28)[_module.__default.safeIndex(_module.__default.fm24(new BigNumber(904), _62_v1, _62_v1, _87_v26, _61_globalState), new BigNumber((_89_v28).length))]);
      let _91_v30;
      _91_v30 = _dafny.Map.Empty.slice().updateUnsafe(_87_v26,_89_v28);
      _87_v26 = (((_88_v27).contains(new BigNumber((_91_v30).length))) ? ((_88_v27).get(new BigNumber((_91_v30).length))) : (((true) ? (_87_v26) : (_87_v26))));
      let _92_v31;
      _92_v31 = new _dafny.CodePoint('a'.codePointAt(0));
      let _source3 = _module.D1.create_DC4(_87_v26, _92_v31, _62_v1, _87_v26);
      if (_source3.is_DC3) {
        let _93___mcc_h0 = (_source3).cf5;
        let _94_cf5 = _93___mcc_h0;
        let _95_v32;
        let _96_v33;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m6(_61_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _95_v32 = _out2;
        _96_v33 = _out3;
        let _97_v34;
        _97_v34 = _module.D9.create_DC19(((_89_v28)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_89_v28).length))]).plus(new BigNumber((_dafny.MultiSet.fromElements(_62_v1)).cardinality())), _94_cf5);
        let _98_v35;
        _98_v35 = _module.D2.create_DC6(_94_cf5, _62_v1);
        let _rhs5 = _97_v34;
        let _rhs6 = _module.__default.safeModuloInt(_87_v26, (_dafny.ZERO).minus(_module.__default.fm19((_98_v35).dtor_cf12, _94_cf5, _61_globalState)));
        _97_v34 = _rhs5;
        _94_cf5 = _rhs6;
        let _99_v36;
        _99_v36 = _dafny.Map.Empty.slice().updateUnsafe(_62_v1,_62_v1);
        let _100_v37;
        _100_v37 = _module.D15.create_DC31(_99_v36);
        _99_v36 = (_100_v37).dtor_cf52;
        let _101_v38;
        _101_v38 = _dafny.Seq.of(_62_v1, _62_v1);
        _101_v38 = _dafny.Seq.update(_101_v38, _module.__default.safeIndex(_87_v26, new BigNumber((_101_v38).length)), _62_v1);
      } else if (_source3.is_DC4) {
        let _102___mcc_h1 = (_source3).cf6;
        let _103___mcc_h2 = (_source3).cf7;
        let _104___mcc_h3 = (_source3).cf8;
        let _105___mcc_h4 = (_source3).cf9;
        let _106_cf9 = _105___mcc_h4;
        let _107_cf8 = _104___mcc_h3;
        let _108_cf7 = _103___mcc_h2;
        let _109_cf6 = _102___mcc_h1;
        let _110_v39;
        let _nw17 = new _module.C3();
        _nw17.__ctor();
        _110_v39 = _nw17;
        let _111_v40;
        _111_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm14(_87_v26, new BigNumber(-263), _61_globalState),((_62_v1) ? (_110_v39) : (_110_v39)));
        let _112_v41;
        _112_v41 = _dafny.Seq.of(_107_cf8, _107_cf8);
        _111_v40 = (_111_v40).update(new BigNumber((_dafny.Seq.Concat(_112_v41, _112_v41)).length), _110_v39);
        _87_v26 = _109_cf6;
        _107_cf8 = !((_106_cf9).plus(_109_cf6)).isEqualTo(_87_v26);
      } else {
        let _113___mcc_h5 = (_source3).cf4;
        let _114_cf4 = _113___mcc_h5;
        _62_v1 = !(_87_v26).isEqualTo(_87_v26);
        let _115_v42;
        _115_v42 = _dafny.Seq.UnicodeFromString("aptekw");
        let _116_v43;
        _116_v43 = _dafny.MultiSet.fromElements(new BigNumber(278), _87_v26, new BigNumber((_115_v42).length));
        _114_cf4 = _module.__default.fm8(_116_v43, _61_globalState);
        let _117_v44;
        _117_v44 = _dafny.Map.Empty.slice().updateUnsafe(_114_cf4,_module.__default.fm4(!(_114_cf4), _62_v1, _114_cf4, _61_globalState));
        let _118_v45;
        _118_v45 = _dafny.MultiSet.fromElements(_114_cf4);
        let _119_v46;
        let _nw18 = new _module.C1();
        _nw18.__ctor(_118_v45, new BigNumber((_module.__default.fm35(false, _62_v1, _62_v1, _87_v26, _61_globalState)).length));
        _119_v46 = _nw18;
        let _120_v47;
        let _nw19 = Array((new BigNumber(2)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _119_v46;
        _nw19[(_dafny.ONE).toNumber()] = _119_v46;
        _120_v47 = _nw19;
        let _121_v48;
        _121_v48 = _dafny.Seq.of(_120_v47);
        let _rhs7 = !_dafny.Seq.contains(_121_v48, _120_v47);
        let _rhs8 = (_119_v46.f11).minus((_dafny.ZERO).minus(_119_v46.f11));
        let _rhs9 = (((_117_v44).update(_62_v1, _115_v42)).update(_62_v1, _115_v42)).Merge((_117_v44).update(!(_62_v1), _115_v42));
        _114_cf4 = _rhs7;
        _87_v26 = _rhs8;
        _117_v44 = _rhs9;
        let _index8 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_60_v0).length));
        (_60_v0)[_index8] = true;
        let _index9 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_60_v0).length));
        (_60_v0)[_index9] = !(_62_v1) || (_114_cf4);
      }
      _62_v1 = _62_v1;
      _87_v26 = _module.__default.safeDivisionInt(_87_v26, _87_v26);
      let _122_v49;
      _122_v49 = _dafny.MultiSet.fromElements(_62_v1, _62_v1, _62_v1, _62_v1, _62_v1);
      let _123_v50;
      let _nw20 = new _module.C1();
      _nw20.__ctor(_122_v49, _87_v26);
      _123_v50 = _nw20;
      let _124_v51;
      _124_v51 = _dafny.Seq.UnicodeFromString("nmpkjd");
      let _125_v52;
      _125_v52 = _dafny.Set.fromElements(_87_v26);
      let _rhs10 = _62_v1;
      let _rhs11 = _62_v1;
      let _rhs12 = _123_v50;
      let _rhs13 = _module.__default.fm30(_124_v51, _87_v26, (_87_v26).plus(new BigNumber((_125_v52).length)), _61_globalState);
      let _rhs14 = ((_62_v1) ? (_62_v1) : (_62_v1));
      _62_v1 = _rhs10;
      _62_v1 = _rhs11;
      _123_v50 = _rhs12;
      _92_v31 = _rhs13;
      _62_v1 = _rhs14;
      let _126_v53;
      _126_v53 = _module.D1.create_DC4((_dafny.ZERO).minus(_87_v26), _92_v31, _62_v1, _module.__default.fm24(_87_v26, _62_v1, _62_v1, _87_v26, _61_globalState));
      let _127_v54;
      let _nw21 = Array((new BigNumber(11)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = _126_v53;
      _nw21[(_dafny.ONE).toNumber()] = _126_v53;
      _nw21[(new BigNumber(2)).toNumber()] = _126_v53;
      _nw21[(new BigNumber(3)).toNumber()] = _126_v53;
      _nw21[(new BigNumber(4)).toNumber()] = _126_v53;
      _nw21[(new BigNumber(5)).toNumber()] = _126_v53;
      _nw21[(new BigNumber(6)).toNumber()] = _module.D1.create_DC4(_87_v26, _92_v31, _62_v1, _87_v26);
      _nw21[(new BigNumber(7)).toNumber()] = _module.__default.fm36(_61_globalState);
      _nw21[(new BigNumber(8)).toNumber()] = _126_v53;
      _nw21[(new BigNumber(9)).toNumber()] = _126_v53;
      _nw21[(new BigNumber(10)).toNumber()] = _module.__default.fm36(_61_globalState);
      _127_v54 = _nw21;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_127_v54).length))) {
        let _128_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_128_i0)) && ((_128_i0).isLessThan(new BigNumber((_127_v54).length))))) {
          (_127_v54)[(_128_i0)] = _126_v53;
        }
      }
      let _129_v55;
      let _nw22 = Array((new BigNumber(13)).toNumber()).fill([]);
      _129_v55 = _nw22;
      let _130_v56;
      let _131_v57;
      let _out4;
      let _out5;
      let _outcollector2 = (_123_v50).m0(_129_v55, _61_globalState);
      _out4 = _outcollector2[0];
      _out5 = _outcollector2[1];
      _130_v56 = _out4;
      _131_v57 = _out5;
      _130_v56 = _130_v56;
      let _hi0 = _87_v26;
      for (let _132_i1 = (_87_v26).minus(_87_v26); _132_i1.isLessThan(_hi0); _132_i1 = _132_i1.plus(_dafny.ONE)) {
        let _133_v58;
        _133_v58 = _dafny.Map.Empty.slice().updateUnsafe(_92_v31,_130_v56);
        _62_v1 = ((_62_v1) ? (!((_133_v58).update(_92_v31, false)).equals(_133_v58)) : (!_dafny.areEqual(_89_v28, _dafny.Seq.of(_132_i1))));
        let _134_v59;
        let _nw23 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _134_v59 = _nw23;
        let _135_v60;
        let _nw24 = Array((new BigNumber(12)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = _134_v59;
        _nw24[(_dafny.ONE).toNumber()] = _134_v59;
        _nw24[(new BigNumber(2)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(3)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(4)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(5)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(6)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(7)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(8)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(9)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(10)).toNumber()] = _134_v59;
        _nw24[(new BigNumber(11)).toNumber()] = _134_v59;
        _135_v60 = _nw24;
        let _136_v61;
        _136_v61 = _dafny.Map.Empty.slice().updateUnsafe(!(_62_v1),_132_i1);
        let _137_v62;
        _137_v62 = _136_v61;
        _134_v59 = (_module.D14.create_DC30(new BigNumber(762), _135_v60, _137_v62, new BigNumber(-318), _134_v59)).dtor_cf51;
        let _138_v63;
        let _nw25 = new _module.C3();
        _nw25.__ctor();
        _138_v63 = _nw25;
        let _139_v64;
        _139_v64 = _module.D11.create_DC24(_138_v63);
        _139_v64 = _139_v64;
        let _140_v65;
        let _141_v66;
        let _out6;
        let _out7;
        let _outcollector3 = (_138_v63).m0(_129_v55, _61_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _140_v65 = _out6;
        _141_v66 = _out7;
      }
      let _142_v67;
      _142_v67 = _dafny.MultiSet.fromElements(_module.D10.create_DC22());
      let _143_i2;
      _143_i2 = _dafny.ZERO;
      L0: {
        while ((((_142_v67).IsSubsetOf(_142_v67)) ? (!(false)) : (_130_v56))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_143_i2)) {
              break L0;
            }
            _143_i2 = (_143_i2).plus(_dafny.ONE);
            let _144_v68;
            let _nw26 = new _module.C2();
            _nw26.__ctor(_dafny.Seq.Concat(_124_v51, _124_v51), _124_v51);
            _144_v68 = _nw26;
            _130_v56 = !(!(_62_v1) || (_62_v1));
            let _145_v69;
            let _nw27 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
            _145_v69 = _nw27;
            _145_v69 = (((_module.__default.fm20(_130_v56, (_144_v68).f9, _62_v1, _61_globalState)).IsSubsetOf(_125_v52)) ? (((_130_v56) ? (_145_v69) : (_145_v69))) : (_145_v69));
            _87_v26 = _module.__default.safeDivisionInt(_87_v26, new BigNumber(15));
          }
        }
      }
      let _146_v70;
      let _nw28 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _146_v70 = _nw28;
      let _index10 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_146_v70).length));
      (_146_v70)[_index10] = new BigNumber(72);
      let _index11 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_146_v70).length));
      (_146_v70)[_index11] = _module.__default.safeModuloInt(new BigNumber(-592), _87_v26);
      let _147_v71;
      let _148_v72;
      let _out8;
      let _out9;
      let _outcollector4 = (_123_v50).m0(_129_v55, _61_globalState);
      _out8 = _outcollector4[0];
      _out9 = _outcollector4[1];
      _147_v71 = _out8;
      _148_v72 = _out9;
      _62_v1 = !(_62_v1);
      _87_v26 = new BigNumber((((_130_v56) ? (_dafny.Seq.update(_124_v51, _module.__default.safeIndex((_146_v70)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_146_v70).length))], new BigNumber((_124_v51).length)), _92_v31)) : (_124_v51))).length);
      _124_v51 = _124_v51;
      process.stdout.write(_dafny.toString((_60_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_61_globalState).f0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_61_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_62_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_87_v26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(542),new BigNumber(542)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_89_v28, _dafny.Seq.of(new BigNumber(542)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_90_v29, _dafny.Seq.of(new BigNumber(542)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(542),_dafny.Seq.of(new BigNumber(542))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_92_v31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v49).equals(_dafny.MultiSet.fromElements(false, false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_124_v51).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v52).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v53).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v53).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v53).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v53).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ZERO]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ZERO]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ZERO]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ZERO]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ONE]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ONE]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ONE]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[_dafny.ONE]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(2)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(2)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(2)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(2)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(3)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(3)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(3)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(3)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(4)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(4)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(4)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(4)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(5)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(5)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(5)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(5)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(6)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(6)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(6)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(6)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(7)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(7)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(7)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(7)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(8)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(8)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(8)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(8)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(9)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(9)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(9)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(9)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(10)]).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(10)]).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(10)]).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v54)[new BigNumber(10)]).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_130_v56));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v57).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.MultiSet.fromElements(false, false, false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_142_v67).equals(_dafny.MultiSet.fromElements(_module.D10.create_DC22()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_143_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v70)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_v71));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_v72).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.MultiSet.fromElements(false, false, false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC1(cf3) {
      let $dt = new D0(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + this.cf1.toVerbatimString(true) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), []);
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
    static create_DC3(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf6, cf7, cf8, cf9) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
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
    get dtor_cf9() { return this.cf9; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf4 === other.cf4;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.ZERO);
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
    static create_DC6(cf11, cf12) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC5(cf10) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO, false);
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
    static create_DC7(cf13) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.ZERO);
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
    static create_DC10(cf16) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(false);
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
    static create_DC12(cf18) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf18) + ")";
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
    static create_DC13(cf19) {
      let $dt = new D6(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf19) + ")";
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
      return [];
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
    static create_DC15(cf21, cf22, cf23, cf24, cf25) {
      let $dt = new D7(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf20) {
      let $dt = new D7(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D7(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf20) + ")";
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
        return other.$tag === 0 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(false, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, null);
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
    static create_DC17(cf27) {
      let $dt = new D8(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf27) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC19(cf29, cf30) {
      let $dt = new D9(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC20(cf31, cf32, cf33) {
      let $dt = new D9(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D9(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31 && this.cf32 === other.cf32 && this.cf33 === other.cf33;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC19(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC22() {
      let $dt = new D10(0);
      return $dt;
    }
    static create_DC23(cf35, cf36, cf37, cf38) {
      let $dt = new D10(1);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC21(cf34) {
      let $dt = new D10(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22";
      } else if (this.$tag === 1) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf34) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && this.cf38 === other.cf38;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22();
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
    static create_DC25(cf40, cf41) {
      let $dt = new D11(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC26(cf42, cf43) {
      let $dt = new D11(1);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC24(cf39) {
      let $dt = new D11(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42) && this.cf43 === other.cf43;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf39 === other.cf39;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25(false, _module.D9.Default());
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
    static create_DC27(cf44) {
      let $dt = new D12(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
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
    static create_DC28(cf45) {
      let $dt = new D13(0);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC28" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf45, other.cf45);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf47, cf48, cf49, cf50, cf51) {
      let $dt = new D14(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC29(cf46) {
      let $dt = new D14(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC29" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49) && _dafny.areEqual(this.cf50, other.cf50) && this.cf51 === other.cf51;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC30(_dafny.ZERO, [], _dafny.Map.Empty, _dafny.ZERO, []);
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
    static create_DC33(cf54) {
      let $dt = new D15(1);
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC34(cf55, cf56, cf57, cf58, cf59) {
      let $dt = new D15(2);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC31(cf52) {
      let $dt = new D15(3);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC33" + "(" + this.cf54.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC34" + "(" + this.cf55.toVerbatimString(true) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC31" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf53 === other.cf53;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC32(false);
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
    static create_DC35(cf60) {
      let $dt = new D16(0);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf60, other.cf60);
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
          return D16.Default();
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
      this._f0 = [];
      this._f1 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
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

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f11 = _dafny.ZERO;
      this._f10 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f10, f11) {
      let _this = this;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f10;
    };
    fm21(p0, p1, p2, globalState) {
      let _this = this;
      return !(!(((!_dafny.areEqual(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("talxrgws")).length)), _dafny.Seq.of(_this.f11))) ? ((_dafny.Set.fromElements(!(true))).IsSubsetOf(_dafny.Set.fromElements(false, true))) : ((_this.f11).isEqualTo(_this.f11)))));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _149_v0;
      _149_v0 = true;
      let _150_v1;
      _150_v1 = _module.D1.create_DC2(!(_149_v0));
      let _151_v2;
      _151_v2 = _dafny.Seq.of(false);
      let _152_v3;
      _152_v3 = _module.D4.create_DC10(!(false));
      let _153_v4;
      _153_v4 = new _dafny.CodePoint('i'.codePointAt(0));
      let _154_v5;
      let _nw29 = Array((new BigNumber(8)).toNumber());
      _nw29[(_dafny.ZERO).toNumber()] = _149_v0;
      _nw29[(_dafny.ONE).toNumber()] = (new BigNumber(436)).isEqualTo(_this.f11);
      _nw29[(new BigNumber(2)).toNumber()] = ((_150_v1).dtor_cf4) || (_149_v0);
      _nw29[(new BigNumber(3)).toNumber()] = (_this.f11).isLessThan(new BigNumber((_dafny.Seq.update(_151_v2, _module.__default.safeIndex(_this.f11, new BigNumber((_151_v2).length)), _149_v0)).length));
      _nw29[(new BigNumber(4)).toNumber()] = !((_152_v3).dtor_cf16);
      _nw29[(new BigNumber(5)).toNumber()] = _149_v0;
      _nw29[(new BigNumber(6)).toNumber()] = !(_149_v0) || (_149_v0);
      _nw29[(new BigNumber(7)).toNumber()] = (_this).fm21(_dafny.Map.Empty.slice().updateUnsafe(_153_v4,new BigNumber(179)), _this.f11, _this.f11, globalState);
      _154_v5 = _nw29;
      let _index12 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
      (_154_v5)[_index12] = _149_v0;
      let _155_v6;
      _155_v6 = _dafny.Map.Empty.slice().updateUnsafe(_153_v4,_this.f11);
      let _index13 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
      (_154_v5)[_index13] = (_this).fm21((_155_v6).update(_153_v4, _this.f11), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(567), _this.f11)), _this.f11, globalState);
      let _156_v7;
      let _nw30 = Array((new BigNumber(2)).toNumber());
      _nw30[(_dafny.ZERO).toNumber()] = _this.f11;
      _nw30[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_this.f11);
      _156_v7 = _nw30;
      let _index14 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
      (_156_v7)[_index14] = new BigNumber((_module.__default.fm22((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], new BigNumber(-599), globalState)).length);
      let _index15 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
      (_156_v7)[_index15] = (_dafny.ZERO).minus(_this.f11);
      let _157_v8;
      _157_v8 = _dafny.Map.Empty.slice().updateUnsafe((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))],(_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]);
      let _hi1 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_151_v2)[_module.__default.safeIndex(_this.f11, new BigNumber((_151_v2).length))])).length), new BigNumber(441));
      for (let _158_i0 = new BigNumber((_157_v8).length); _158_i0.isLessThan(_hi1); _158_i0 = _158_i0.plus(_dafny.ONE)) {
        let _index16 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
        (_156_v7)[_index16] = new BigNumber((_module.__default.fm23(_158_i0, globalState)).length);
        let _159_v9;
        _159_v9 = _dafny.Seq.of((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]);
        let _160_v10;
        _160_v10 = _dafny.Seq.of(_this.f11, new BigNumber((_159_v9).length), new BigNumber((_dafny.MultiSet.fromElements(!(!(_149_v0)), _149_v0, (_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))])).cardinality()));
        let _161_v11;
        _161_v11 = _module.D2.create_DC5(_159_v9);
        _161_v11 = _module.D2.create_DC5(_160_v10);
        let _162_v12;
        _162_v12 = _dafny.Seq.UnicodeFromString("cltsjvhb");
        let _163_v13;
        let _nw31 = Array((new BigNumber(3)).toNumber()).fill([]);
        _163_v13 = _nw31;
        let _164_v14;
        _164_v14 = _module.D0.create_DC0(new BigNumber(411), _162_v12, _163_v13);
        let _165_v15;
        _165_v15 = _dafny.Map.Empty.slice().updateUnsafe(_154_v5,_164_v14);
        let _166_v16;
        _166_v16 = _dafny.Seq.of(_164_v14);
        let _167_v17;
        _167_v17 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_151_v2)).length)));
        let _168_v18;
        _168_v18 = _module.D0.create_DC1((((_165_v15).contains(_154_v5)) ? ((_165_v15).get(_154_v5)) : ((_166_v16)[_module.__default.safeIndex(new BigNumber((_167_v17).length), new BigNumber((_166_v16).length))])));
        let _source4 = _168_v18;
        if (_source4.is_DC0) {
          let _169___mcc_h0 = (_source4).cf0;
          let _170___mcc_h1 = (_source4).cf1;
          let _171___mcc_h2 = (_source4).cf2;
          let _172_cf2 = _171___mcc_h2;
          let _173_cf1 = _170___mcc_h1;
          let _174_cf0 = _169___mcc_h0;
          let _175_v19;
          let _nw32 = new _module.C0();
          _nw32.__ctor();
          _175_v19 = _nw32;
          let _index17 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
          (_154_v5)[_index17] = !(_dafny.areEqual(_160_v10, _159_v9)) || (false);
          let _176_v20;
          let _nw33 = new _module.C0();
          _nw33.__ctor();
          _176_v20 = _nw33;
          let _177_v21;
          _177_v21 = _dafny.MultiSet.fromElements(_this.f11);
          let _178_v22;
          _178_v22 = _dafny.Map.Empty.slice().updateUnsafe((_149_v0) === (!((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))])),(_177_v21).contains(_this.f11));
          _178_v22 = (_178_v22).update((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], (_150_v1).dtor_cf4);
        } else {
          let _179___mcc_h3 = (_source4).cf3;
          let _180_cf3 = _179___mcc_h3;
          let _index18 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          let _rhs15 = _this.f11;
          let _rhs16 = _this.f11;
          let _lhs2 = _156_v7;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          let _lhs4 = _156_v7;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          _lhs2[_lhs3] = _rhs15;
          _lhs4[_lhs5] = _rhs16;
          let _index20 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
          (_154_v5)[_index20] = (_this).fm21(_155_v6, ((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]).plus(_this.f11), _module.__default.fm24(_this.f11, (_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], _149_v0, new BigNumber((_159_v9).length), globalState), globalState);
          _153_v4 = _153_v4;
          _153_v4 = _153_v4;
        }
        let _index21 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
        (_154_v5)[_index21] = _149_v0;
      }
      let _181_v23;
      _181_v23 = _dafny.Seq.UnicodeFromString("vtlum");
      let _hi2 = _this.f11;
      for (let _182_i1 = _module.__default.safeDivisionInt(new BigNumber((_181_v23).length), _this.f11); _182_i1.isLessThan(_hi2); _182_i1 = _182_i1.plus(_dafny.ONE)) {
        let _183_v24;
        _183_v24 = _dafny.Seq.of(_182_i1, new BigNumber(712));
        _149_v0 = (((_this).f10).update(_149_v0, _module.__default.abs(new BigNumber((_183_v24).length)))).IsSubsetOf((_this).f10);
        let _index22 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
        let _rhs17 = _149_v0;
        let _rhs18 = (_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))];
        let _rhs19 = (_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))];
        let _lhs6 = _this;
        let _lhs7 = _156_v7;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
        r0 = _rhs17;
        _lhs6.f11 = _rhs18;
        _lhs7[_lhs8] = _rhs19;
        let _184_v25;
        let _nw34 = Array((new BigNumber(23)).toNumber()).fill(_module.D2.Default());
        _184_v25 = _nw34;
        let _185_v26;
        _185_v26 = _module.D2.create_DC6(new BigNumber((_dafny.MultiSet.FromArray(_183_v24)).cardinality()), (_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]);
        let _index23 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_184_v25).length));
        (_184_v25)[_index23] = _185_v26;
        let _index24 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_184_v25).length));
        (_184_v25)[_index24] = _185_v26;
        let _186_v27;
        let _nw35 = Array((_dafny.ONE).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-530)), ((_187_v23) => function (_188_i2) {
          return _187_v23;
        })(_181_v23));
        _186_v27 = _nw35;
        let _189_v28;
        _189_v28 = _dafny.Seq.of(_181_v23);
        let _index25 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_186_v27).length));
        (_186_v27)[_index25] = _189_v28;
        let _index26 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_186_v27).length));
        (_186_v27)[_index26] = _dafny.Seq.Concat(_189_v28, _module.__default.fm25(_181_v23, _module.__default.fm24((_dafny.ZERO).minus(_182_i1), (_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], _149_v0, new BigNumber((_dafny.MultiSet.FromArray(_183_v24)).cardinality()), globalState), _182_i1, globalState));
      }
      if (true) {
        _181_v23 = _181_v23;
        if ((_149_v0) && (((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]) && ((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]))) {
          _149_v0 = _149_v0;
          let _190_v29;
          let _nw36 = Array((new BigNumber(27)).toNumber()).fill([]);
          _190_v29 = _nw36;
          _190_v29 = p0;
          let _191_v30;
          _191_v30 = _module.D1.create_DC3(_this.f11);
          let _192_v31;
          _192_v31 = _dafny.Seq.of((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]);
          let _193_v32;
          _193_v32 = _dafny.Map.Empty.slice().updateUnsafe((((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]) ? (_dafny.Seq.of((_191_v30).dtor_cf5, _this.f11)) : (_192_v31)),(false) || ((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]));
          _193_v32 = _193_v32;
          let _index27 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          (_156_v7)[_index27] = _this.f11;
          let _index28 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
          (_154_v5)[_index28] = (false) || (_dafny.areEqual(_192_v31, _dafny.Seq.of(_this.f11)));
        } else {
          let _index29 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          (_156_v7)[_index29] = (_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))];
          _149_v0 = (_151_v2)[_module.__default.safeIndex((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], new BigNumber((_151_v2).length))];
          _150_v1 = _150_v1;
          let _194_v33;
          _194_v33 = _dafny.Seq.of(_this.f11, new BigNumber((_181_v23).length), (_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]);
          let _195_v34;
          _195_v34 = _dafny.Map.Empty.slice().updateUnsafe(_149_v0,((_dafny.MultiSet.fromElements((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), ((_196_v4) => function (_197_i3) {
            return _196_v4;
          })(_153_v4))).length))).update(_this.f11, _module.__default.abs((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]))).Difference(_dafny.MultiSet.FromArray(_194_v33)));
          let _198_v36;
          _198_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_153_v4);
          let _199_v37;
          _199_v37 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(304)), ((_200_v4) => function (_201_i4) {
            return _200_v4;
          })(_153_v4)), _module.__default.safeIndex(_this.f11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(304)), ((_202_v4) => function (_203_i4) {
            return _202_v4;
          })(_153_v4))).length)), _153_v4)).length));
          _195_v34 = (_195_v34).update(!((_this).fm21(function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe((((_198_v36).contains(_this.f11)) ? ((_198_v36).get(_this.f11)) : (_153_v4)),false)).Keys.Elements) {
              let _204_v35 = _compr_8;
              if ((_dafny.Map.Empty.slice().updateUnsafe((((_198_v36).contains(_this.f11)) ? ((_198_v36).get(_this.f11)) : (_153_v4)),false)).contains(_204_v35)) {
                _coll8.push([_204_v35,_this.f11]);
              }
            }
            return _coll8;
          }(), (_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], new BigNumber((_module.__default.fm26((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], globalState)).length), globalState)) || (_149_v0), _199_v37);
          let _index30 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
          (_154_v5)[_index30] = (true) || ((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]);
        }
        let _205_v38;
        _205_v38 = _dafny.Seq.of(_this.f11);
        let _206_v39;
        _206_v39 = _module.D2.create_DC5(_205_v38);
        _206_v39 = _206_v39;
        let _207_v40;
        let _nw37 = new _module.C0();
        _nw37.__ctor();
        _207_v40 = _nw37;
        let _208_v41;
        let _nw38 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
        _208_v41 = _nw38;
        let _209_v42;
        let _nw39 = Array((new BigNumber(28)).toNumber());
        _209_v42 = _nw39;
        let _210_v43;
        _210_v43 = _dafny.Set.fromElements(_209_v42);
        let _index31 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_208_v41).length));
        (_208_v41)[_index31] = _210_v43;
        let _211_v44;
        _211_v44 = _dafny.Seq.of(_209_v42, _209_v42, _209_v42, _209_v42);
        let _index32 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_208_v41).length));
        (_208_v41)[_index32] = (_210_v43).Difference(_dafny.Set.fromElements((_211_v44)[_module.__default.safeIndex((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], new BigNumber((_211_v44).length))], _209_v42));
      } else {
        let _index33 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
        (_156_v7)[_index33] = (_dafny.ZERO).minus(_this.f11);
        let _index34 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
        (_154_v5)[_index34] = _149_v0;
        _149_v0 = !((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]);
        let _212_v45;
        let _nw40 = new _module.C0();
        _nw40.__ctor();
        _212_v45 = _nw40;
        (_this).f11 = new BigNumber(((_this).f10).cardinality());
      }
      if ((_151_v2)[_module.__default.safeIndex(((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]).minus(_this.f11), new BigNumber((_151_v2).length))]) {
        let _213_v46;
        _213_v46 = _dafny.MultiSet.fromElements(_153_v4, _153_v4, _153_v4);
        _213_v46 = _213_v46;
        let _214_v47;
        _214_v47 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("dqbpjl"));
        let _215_v48;
        _215_v48 = _dafny.Set.fromElements(_153_v4, _153_v4, _153_v4, _153_v4, _153_v4);
        let _216_v49;
        _216_v49 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("a"), _181_v23, _181_v23, _dafny.Seq.UnicodeFromString("rmkxsg"), _181_v23);
        let _217_v50;
        _217_v50 = _dafny.Seq.of(_214_v47, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-244)), ((_218_v23) => function (_219_i5) {
          return _218_v23;
        })(_181_v23)));
        let _220_v51;
        let _nw41 = Array((new BigNumber(28)).toNumber());
        _nw41[(_dafny.ZERO).toNumber()] = _156_v7;
        _nw41[(_dafny.ONE).toNumber()] = _156_v7;
        _nw41[(new BigNumber(2)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(3)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(4)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(5)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(6)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(7)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(8)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(9)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(10)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(11)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(12)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(13)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(14)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(15)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(16)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(17)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(18)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(19)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(20)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(21)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(22)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(23)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(24)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(25)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(26)).toNumber()] = _156_v7;
        _nw41[(new BigNumber(27)).toNumber()] = _156_v7;
        _220_v51 = _nw41;
        let _221_v52;
        _221_v52 = _module.D0.create_DC0(new BigNumber(592), _dafny.Seq.UnicodeFromString("ajxyfq"), _220_v51);
        let _222_v53;
        let _nw42 = Array((new BigNumber(28)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_214_v47);
        _nw42[(_dafny.ONE).toNumber()] = (_module.__default.fm27(globalState)).update(_181_v23, _module.__default.abs(new BigNumber((_215_v48).length)));
        _nw42[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray((((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]) ? (_214_v47) : (_214_v47)));
        _nw42[(new BigNumber(3)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(4)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(5)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(6)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(7)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(_181_v23, _dafny.Seq.update(_module.__default.fm22((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], _this.f11, globalState), _module.__default.safeIndex(new BigNumber(((_this).f10).cardinality()), new BigNumber((_module.__default.fm22((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], _this.f11, globalState)).length)), _153_v4), _181_v23, _181_v23, _181_v23);
        _nw42[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray((_217_v50)[_module.__default.safeIndex((_221_v52).dtor_cf0, new BigNumber((_217_v50).length))]);
        _nw42[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements(_181_v23, _181_v23)).Difference(_216_v49);
        _nw42[(new BigNumber(11)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_214_v47, _214_v47));
        _nw42[(new BigNumber(13)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(14)).toNumber()] = (((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]) ? (_216_v49) : (_216_v49));
        _nw42[(new BigNumber(15)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(16)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(17)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(18)).toNumber()] = _module.__default.fm27(globalState);
        _nw42[(new BigNumber(19)).toNumber()] = (_216_v49).Union(_216_v49);
        _nw42[(new BigNumber(20)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(21)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(22)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(23)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(24)).toNumber()] = _216_v49;
        _nw42[(new BigNumber(25)).toNumber()] = (_dafny.MultiSet.FromArray(_214_v47)).Difference(_dafny.MultiSet.fromElements(_181_v23));
        _nw42[(new BigNumber(26)).toNumber()] = _dafny.MultiSet.fromElements(_181_v23);
        _nw42[(new BigNumber(27)).toNumber()] = _216_v49;
        _222_v53 = _nw42;
        let _index35 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_222_v53).length));
        (_222_v53)[_index35] = _216_v49;
        let _index36 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_222_v53).length));
        (_222_v53)[_index36] = _dafny.MultiSet.fromElements(_181_v23, _dafny.Seq.Concat(_181_v23, _dafny.Seq.UnicodeFromString("xce")), _181_v23);
        if ((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]) {
          (_this).f11 = (_dafny.ZERO).minus((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]);
          _149_v0 = ((_149_v0) ? (_149_v0) : (!((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]) || (_149_v0)));
          let _index37 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          (_156_v7)[_index37] = _this.f11;
          _149_v0 = !_dafny.Seq.contains(_181_v23, _153_v4);
          let _index38 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
          (_154_v5)[_index38] = _149_v0;
        } else {
          let _223_v54;
          let _init1 = ((_224_v23) => function (_225_i6) {
            return _224_v23;
          })(_181_v23);
          let _nw43 = Array((new BigNumber(15)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw43.length); _i0_1++) {
            _nw43[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _223_v54 = _nw43;
          _223_v54 = _223_v54;
          let _226_v55;
          _226_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(((_213_v46).contains(_153_v4)) ? ((_213_v46).get(_153_v4)) : ((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))])));
          (_this).f11 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-943), (((_226_v55).contains(_this.f11)) ? ((_226_v55).get(_this.f11)) : (_this.f11))));
          r0 = (_this).fm21(_dafny.Map.Empty.slice().updateUnsafe(_153_v4,(_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]), _this.f11, _this.f11, globalState);
          let _227_v56;
          _227_v56 = _dafny.Set.fromElements(false);
          let _228_v57;
          _228_v57 = _module.D1.create_DC4((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], _153_v4, _149_v0, _this.f11);
          let _index39 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length));
          (_156_v7)[_index39] = ((new BigNumber((_227_v56).length)).minus((_228_v57).dtor_cf9)).minus(_this.f11);
          _214_v47 = _dafny.Seq.Concat(_214_v47, _214_v47);
        }
        let _index40 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length));
        (_154_v5)[_index40] = !(_dafny.ONE).isEqualTo(((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]).minus(new BigNumber(-273)));
        let _229_v58;
        let _nw44 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _229_v58 = _nw44;
        let _index41 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_229_v58).length));
        (_229_v58)[_index41] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("p"), _181_v23);
        let _230_v59;
        _230_v59 = _dafny.Set.fromElements((new BigNumber(-451)).isEqualTo(_this.f11));
        let _index42 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_229_v58).length));
        let _rhs20 = _181_v23;
        let _rhs21 = _230_v59;
        let _lhs9 = _229_v58;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_229_v58).length));
        _lhs9[_lhs10] = _rhs20;
        _230_v59 = _rhs21;
      } else {
        let _231_v60;
        _231_v60 = _dafny.MultiSet.fromElements(_this.f11);
        let _232_v61;
        _232_v61 = _dafny.Map.Empty.slice().updateUnsafe((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))],_dafny.MultiSet.fromElements((((_231_v60).contains((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))])) ? ((_231_v60).get((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))])) : (new BigNumber((_181_v23).length)))));
        let _233_v62;
        _233_v62 = _module.D3.create_DC8((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))]);
        let _234_v63;
        _234_v63 = _dafny.Map.Empty.slice().updateUnsafe(_149_v0,((((_232_v61).contains((_dafny.ZERO).minus((_233_v62).dtor_cf14))) ? ((_232_v61).get((_dafny.ZERO).minus((_233_v62).dtor_cf14))) : (_231_v60))).Difference(_231_v60));
        _234_v63 = _234_v63;
        _149_v0 = (_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))];
        let _235_v64;
        let _init2 = ((_236_v23, _237_v7, _238_v4) => function (_239_i7) {
          return _dafny.Seq.update(_dafny.Seq.Concat(_236_v23, _dafny.Seq.UnicodeFromString("jtwjl")), _module.__default.safeIndex((_237_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_237_v7).length))], new BigNumber((_dafny.Seq.Concat(_236_v23, _dafny.Seq.UnicodeFromString("jtwjl"))).length)), _238_v4);
        })(_181_v23, _156_v7, _153_v4);
        let _nw45 = Array((new BigNumber(9)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw45.length); _i0_2++) {
          _nw45[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _235_v64 = _nw45;
        let _index43 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_235_v64).length));
        (_235_v64)[_index43] = _181_v23;
        let _240_v65;
        _240_v65 = _dafny.Map.Empty.slice().updateUnsafe(_149_v0,_dafny.Seq.UnicodeFromString("rotcipwa"));
        let _index44 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_235_v64).length));
        (_235_v64)[_index44] = _dafny.Seq.Concat(_module.__default.fm22(_149_v0, (_dafny.ZERO).minus(_this.f11), globalState), (((_240_v65).contains(!((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]))) ? ((_240_v65).get(!((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]))) : (_181_v23)));
        let _241_v66;
        _241_v66 = _dafny.Set.fromElements((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))], (_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))]);
        let _242_v67;
        _242_v67 = _dafny.Map.Empty.slice().updateUnsafe(_154_v5,_149_v0);
        let _243_v68;
        _243_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_241_v66).length),(_242_v67).Merge(_242_v67));
        _243_v68 = (_243_v68).update((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], _242_v67);
        let _244_v69;
        let _nw46 = new _module.C0();
        _nw46.__ctor();
        _244_v69 = _nw46;
      }
      r0 = _dafny.Seq.contains(_181_v23, _153_v4);
      let _245_v70;
      let _nw47 = Array((new BigNumber(17)).toNumber());
      _nw47[(_dafny.ZERO).toNumber()] = _156_v7;
      _nw47[(_dafny.ONE).toNumber()] = _156_v7;
      _nw47[(new BigNumber(2)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(3)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(4)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(5)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(6)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(7)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(8)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(9)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(10)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(11)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(12)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(13)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(14)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(15)).toNumber()] = _156_v7;
      _nw47[(new BigNumber(16)).toNumber()] = _156_v7;
      _245_v70 = _nw47;
      let _246_v71;
      _246_v71 = _dafny.Map.Empty.slice().updateUnsafe((_154_v5)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_154_v5).length))],_module.D0.create_DC0((_156_v7)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_156_v7).length))], _dafny.Seq.UnicodeFromString("lrqnahqbq"), _245_v70));
      let _247_v72;
      _247_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_246_v71).length), _this.f11),(_this).f10);
      r1 = _247_v72;
      return [r0, r1];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let _248_v1;
      _248_v1 = false;
      let _249_v2;
      _249_v2 = _dafny.Seq.of(_248_v1);
      let _250_i0;
      _250_i0 = _dafny.ZERO;
      L1: {
        while (!(_dafny.Seq.contains(_249_v2, !(!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), ((_256_p0) => function (_257_i1) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
            let _coll9 = new _dafny.Set();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(109), new BigNumber(909))) {
              let _258_v0 = _compr_9;
              if (((new BigNumber(109)).isLessThanOrEqualTo(_258_v0)) && ((_258_v0).isLessThan(new BigNumber(909)))) {
                _coll9.add(_module.__default.safeModuloInt(_258_v0, _256_p0));
              }
            }
            return _coll9;
          }()).length))).length);
        })(p0)), p0))))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_250_i0)) {
              break L1;
            }
            _250_i0 = (_250_i0).plus(_dafny.ONE);
            let _251_v3;
            let _nw48 = Array((new BigNumber(5)).toNumber()).fill(false);
            _251_v3 = _nw48;
            let _252_v4;
            _252_v4 = _dafny.Seq.of(_251_v3);
            _252_v4 = _252_v4;
            (_this).f11 = p0;
            let _253_v5;
            _253_v5 = _dafny.Map.Empty.slice().updateUnsafe(_248_v1,p0);
            let _254_v6;
            _254_v6 = _dafny.Map.Empty.slice().updateUnsafe(_253_v5,_248_v1);
            _254_v6 = _254_v6;
            let _255_v7;
            _255_v7 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(p0, new BigNumber((p1).length))],p0);
            _248_v1 = !(((_248_v1) ? ((_this).fm21(_255_v7, p0, p0, globalState)) : ((_this).fm21(_255_v7, p0, _module.__default.fm24(p0, _248_v1, _248_v1, new BigNumber((_249_v2).length), globalState), globalState))));
          }
        }
      }
      let _259_v8;
      _259_v8 = _dafny.Set.fromElements(_this.f11);
      _248_v1 = (_259_v8).IsSubsetOf((_259_v8).Intersect(_259_v8));
      let _260_v9;
      _260_v9 = _dafny.Seq.of(_this.f11, new BigNumber((_249_v2).length));
      let _261_v10;
      _261_v10 = _module.D2.create_DC5(_260_v9);
      let _pat_let_tv0 = _248_v1;
      _248_v1 = function (_source5) {
        if (_source5.is_DC6) {
          let _262___mcc_h0 = (_source5).cf11;
          let _263___mcc_h1 = (_source5).cf12;
          let _264_cf12 = _263___mcc_h1;
          let _265_cf11 = _262___mcc_h0;
          return _pat_let_tv0;
        } else {
          let _266___mcc_h2 = (_source5).cf10;
          let _267_cf10 = _266___mcc_h2;
          return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("xiwja"), new _dafny.CodePoint('o'.codePointAt(0)));
        }
      }(_261_v10);
      let _268_v11;
      _268_v11 = _module.D1.create_DC2(_248_v1);
      let _pat_let_tv1 = _248_v1;
      _268_v11 = function (_pat_let0_0) {
        return function (_269_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_270_dt__update_hcf4_h0) {
              return _module.D1.create_DC2(_270_dt__update_hcf4_h0);
            }(_pat_let1_0);
          }(_pat_let_tv1);
        }(_pat_let0_0);
      }(_module.D1.create_DC2(_248_v1));
      let _271_v12;
      _271_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wnc"),new BigNumber(791));
      let _272_v13;
      _272_v13 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber((_271_v12).length));
      let _273_v14;
      _273_v14 = _dafny.Set.fromElements(_248_v1);
      let _274_v15;
      _274_v15 = _dafny.Map.Empty.slice().updateUnsafe(false,_248_v1);
      let _275_v16;
      _275_v16 = _module.D1.create_DC4(new BigNumber(-768), new _dafny.CodePoint('o'.codePointAt(0)), true, new BigNumber((p1).length));
      let _276_v17;
      let _nw49 = Array((new BigNumber(22)).toNumber());
      _nw49[(_dafny.ZERO).toNumber()] = (_this).fm21(_272_v13, (_dafny.ZERO).minus(p0), p0, globalState);
      _nw49[(_dafny.ONE).toNumber()] = !_dafny.areEqual(_260_v9, _dafny.Seq.of(new BigNumber((_249_v2).length), _this.f11));
      _nw49[(new BigNumber(2)).toNumber()] = !(_273_v14).equals(_273_v14);
      _nw49[(new BigNumber(3)).toNumber()] = !(false);
      _nw49[(new BigNumber(4)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(5)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(6)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(7)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(8)).toNumber()] = true;
      _nw49[(new BigNumber(9)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(10)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(11)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(12)).toNumber()] = (new BigNumber((_274_v15).length)).isLessThanOrEqualTo(_module.__default.fm24(new BigNumber(434), (_249_v2)[_module.__default.safeIndex(p0, new BigNumber((_249_v2).length))], _248_v1, new BigNumber(((_this).f10).cardinality()), globalState));
      _nw49[(new BigNumber(13)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(14)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(15)).toNumber()] = (_this).fm21(_272_v13, p0, _this.f11, globalState);
      _nw49[(new BigNumber(16)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(17)).toNumber()] = (_this).fm21(_272_v13, _module.__default.fm24(new BigNumber(899), _248_v1, false, new BigNumber(345), globalState), new BigNumber((_dafny.Seq.of(_248_v1, false, _248_v1, _248_v1, _248_v1)).length), globalState);
      _nw49[(new BigNumber(18)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(19)).toNumber()] = _248_v1;
      _nw49[(new BigNumber(20)).toNumber()] = (_275_v16).dtor_cf8;
      _nw49[(new BigNumber(21)).toNumber()] = _248_v1;
      _276_v17 = _nw49;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_276_v17).length))) {
        let _277_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_277_i2)) && ((_277_i2).isLessThan(new BigNumber((_276_v17).length))))) {
          (_276_v17)[(_277_i2)] = ((_dafny.MultiSet.FromArray(_260_v9)).Intersect((_module.D3.create_DC7(_dafny.MultiSet.fromElements(new BigNumber(709), p0))).dtor_cf13)).IsDisjointFrom((_dafny.MultiSet.fromElements(p0, _this.f11, new BigNumber((p1).length))).Union(_dafny.MultiSet.FromArray(_260_v9)));
        }
      }
      if (_248_v1) {
        let _pat_let_tv2 = p0;
        if (((_248_v1) && ((_this).fm21(_272_v13, _this.f11, new BigNumber((_dafny.Seq.UnicodeFromString("weqkh")).length), globalState))) || ((function (_pat_let2_0) {
          return function (_278_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_279_dt__update_hcf11_h0) {
                return _module.D2.create_DC6(_279_dt__update_hcf11_h0, (_278_dt__update__tmp_h1).dtor_cf12);
              }(_pat_let3_0);
            }(_pat_let_tv2);
          }(_pat_let2_0);
        }(_module.D2.create_DC6(p0, _248_v1))).dtor_cf12)) {
          (_this).f11 = _this.f11;
          _248_v1 = !(_248_v1);
          (_this).f11 = _this.f11;
          let _280_v18;
          let _nw50 = new _module.C0();
          _nw50.__ctor();
          _280_v18 = _nw50;
          let _281_v19;
          let _nw51 = Array((new BigNumber(26)).toNumber()).fill([]);
          _281_v19 = _nw51;
          let _282_v20;
          let _nw52 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _282_v20 = _nw52;
          let _283_v21;
          _283_v21 = _dafny.Seq.of(_282_v20);
          let _index45 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_281_v19).length));
          (_281_v19)[_index45] = (_283_v21)[_module.__default.safeIndex(_this.f11, new BigNumber((_283_v21).length))];
          let _index46 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_281_v19).length));
          (_281_v19)[_index46] = _282_v20;
        } else {
          _261_v10 = _261_v10;
          _268_v11 = _268_v11;
          let _284_v22;
          let _nw53 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _284_v22 = _nw53;
          _284_v22 = _284_v22;
          _248_v1 = (_module.__default.safeDivisionInt(_this.f11, p0)).isLessThanOrEqualTo((_this.f11).minus(new BigNumber(618)));
          let _285_v23;
          _285_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_259_v8).length),p0);
          let _286_v24;
          _286_v24 = new _dafny.CodePoint('n'.codePointAt(0));
          (_this).f11 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_285_v23).length), _this.f11), new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber(779), new BigNumber((p1).length)), _286_v24))).length));
        }
        let _287_v25;
        let _nw54 = new _module.C0();
        _nw54.__ctor();
        _287_v25 = _nw54;
        let _288_v26;
        let _init3 = ((_289_p0) => function (_290_i3) {
          return (_290_i3).plus(_289_p0);
        })(p0);
        let _nw55 = Array((new BigNumber(12)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw55.length); _i0_3++) {
          _nw55[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _288_v26 = _nw55;
        let _291_v27;
        _291_v27 = _dafny.Seq.of(_288_v26);
        (_this).f11 = (new BigNumber((_291_v27).length)).plus(_this.f11);
        if (false) {
          (_this).f11 = p0;
          let _292_v28;
          _292_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _293_v29;
          _293_v29 = _dafny.Map.Empty.slice().updateUnsafe(_292_v28,_248_v1);
          let _294_v30;
          _294_v30 = _293_v29;
          let _295_v31;
          let _nw56 = Array((new BigNumber(6)).toNumber());
          _nw56[(_dafny.ZERO).toNumber()] = _288_v26;
          _nw56[(_dafny.ONE).toNumber()] = _288_v26;
          _nw56[(new BigNumber(2)).toNumber()] = _288_v26;
          _nw56[(new BigNumber(3)).toNumber()] = _288_v26;
          _nw56[(new BigNumber(4)).toNumber()] = _288_v26;
          _nw56[(new BigNumber(5)).toNumber()] = _288_v26;
          _295_v31 = _nw56;
          let _296_v32;
          _296_v32 = _module.D0.create_DC0(_this.f11, p1, _295_v31);
          let _297_v33;
          _297_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_293_v29)).length),_296_v32);
          let _298_v34;
          _298_v34 = _276_v17;
          let _299_v35;
          let _nw57 = Array((new BigNumber(28)).toNumber());
          _nw57[(_dafny.ZERO).toNumber()] = _276_v17;
          _nw57[(_dafny.ONE).toNumber()] = _276_v17;
          _nw57[(new BigNumber(2)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(3)).toNumber()] = (_298_v34);
          _nw57[(new BigNumber(4)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(5)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(6)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(7)).toNumber()] = (((_268_v11).dtor_cf4) ? (_276_v17) : (_276_v17));
          _nw57[(new BigNumber(8)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(9)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(10)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(11)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(12)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(13)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(14)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(15)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(16)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(17)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(18)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(19)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(20)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(21)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(22)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(23)).toNumber()] = ((_248_v1) ? (_276_v17) : (_276_v17));
          _nw57[(new BigNumber(24)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(25)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(26)).toNumber()] = _276_v17;
          _nw57[(new BigNumber(27)).toNumber()] = (_298_v34);
          _299_v35 = _nw57;
          let _rhs22 = _297_v33;
          let _rhs23 = _299_v35;
          let _rhs24 = ((_dafny.ZERO).minus(_this.f11)).minus(new BigNumber((_260_v9).length));
          let _lhs11 = _this;
          _297_v33 = _rhs22;
          _299_v35 = _rhs23;
          _lhs11.f11 = _rhs24;
          let _300_v36;
          _300_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_249_v2, _249_v2),_249_v2);
          _300_v36 = (_300_v36).update(_249_v2, _249_v2);
          let _301_v37;
          _301_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,_248_v1);
          let _302_v38;
          _302_v38 = _dafny.Map.Empty.slice().updateUnsafe(_276_v17,_dafny.MultiSet.fromElements(new BigNumber((_301_v37).length)));
          let _303_v39;
          _303_v39 = _dafny.MultiSet.fromElements(new BigNumber((p1).length));
          let _304_v40;
          _304_v40 = new _dafny.CodePoint('k'.codePointAt(0));
          let _305_v41;
          _305_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_304_v40);
          let _306_v42;
          _306_v42 = _dafny.Map.Empty.slice().updateUnsafe(((((_302_v38).contains(_276_v17)) ? ((_302_v38).get(_276_v17)) : (_303_v39))).Union(_dafny.MultiSet.fromElements(_module.__default.fm24(_this.f11, _248_v1, (((_301_v37).contains(new BigNumber((_260_v9).length))) ? ((_301_v37).get(new BigNumber((_260_v9).length))) : (_248_v1)), new BigNumber((_305_v41).length), globalState), _this.f11, new BigNumber(-835), _this.f11)),_module.__default.safeDivisionInt((((_303_v39).contains(p0)) ? ((_303_v39).get(p0)) : (p0)), _this.f11));
          let _307_v43;
          _307_v43 = _module.D2.create_DC6(p0, _248_v1);
          let _308_v44;
          _308_v44 = _dafny.Map.Empty.slice().updateUnsafe(_248_v1,new BigNumber(727));
          let _rhs25 = _306_v42;
          let _rhs26 = (((_308_v44).contains(_248_v1)) ? (_this.f11) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_249_v2, _module.__default.safeIndex(p0, new BigNumber((_249_v2).length)), _248_v1), _249_v2), _module.__default.safeIndex(new BigNumber((_308_v44).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_249_v2, _module.__default.safeIndex(p0, new BigNumber((_249_v2).length)), _248_v1), _249_v2)).length)), true)).length)));
          let _rhs27 = (((_259_v8).IsProperSubsetOf(_259_v8)) ? (_307_v43) : (_module.D2.create_DC6(_this.f11, _248_v1)));
          let _rhs28 = (_273_v14).IsSubsetOf(_dafny.Set.fromElements(false));
          let _lhs12 = _this;
          _306_v42 = _rhs25;
          _lhs12.f11 = _rhs26;
          _307_v43 = _rhs27;
          _248_v1 = _rhs28;
          let _309_v45;
          _309_v45 = _dafny.Seq.UnicodeFromString("aqinu");
          _309_v45 = p1;
        } else {
          (_this).f11 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((p0).minus(p0)), (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber((_dafny.MultiSet.fromElements(_248_v1)).cardinality()))));
          let _310_v46;
          _310_v46 = _module.D3.create_DC8(new BigNumber(175));
          let _311_v47;
          _311_v47 = _dafny.Map.Empty.slice().updateUnsafe(_310_v46,_248_v1);
          let _312_v48;
          _312_v48 = new _dafny.CodePoint('x'.codePointAt(0));
          _311_v47 = (((_248_v1) ? (_dafny.Map.Empty.slice().updateUnsafe(_310_v46,(_this).fm21(_dafny.Map.Empty.slice().updateUnsafe(_312_v48,_this.f11), _this.f11, p0, globalState))) : (_311_v47))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_310_v46,true));
          _276_v17 = _276_v17;
          _248_v1 = _248_v1;
          let _index47 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_276_v17).length));
          (_276_v17)[_index47] = _248_v1;
          let _index48 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_276_v17).length));
          (_276_v17)[_index48] = _248_v1;
        }
        _248_v1 = (p0).isLessThanOrEqualTo((_dafny.ZERO).minus(p0));
      } else {
        let _313_v49;
        _313_v49 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_248_v1)),p0);
        (_this).f11 = (((_313_v49).contains(_248_v1)) ? ((_313_v49).get(_248_v1)) : (_this.f11));
        let _314_v50;
        _314_v50 = new _dafny.CodePoint('y'.codePointAt(0));
        let _315_v51;
        _315_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_314_v50);
        _315_v51 = (_315_v51).update(_this.f11, ((_248_v1) ? (_314_v50) : (_314_v50)));
        _249_v2 = _dafny.Seq.Concat(_dafny.Seq.of(_248_v1), (_module.D7.create_DC14(_249_v2)).dtor_cf20);
        let _316_v52;
        _316_v52 = _dafny.Seq.of(_274_v15);
        let _317_v53;
        _317_v53 = _module.D4.create_DC10(_248_v1);
        let _318_v54;
        _318_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_317_v53).dtor_cf16,_248_v1)));
        _248_v1 = !_dafny.areEqual(_dafny.Seq.Concat(_316_v52, _316_v52), (((_318_v54).contains(p0)) ? ((_318_v54).get(p0)) : (_316_v52)));
        let _319_v55;
        _319_v55 = _dafny.MultiSet.fromElements(_module.__default.fm24(p0, _248_v1, _248_v1, p0, globalState), p0);
        let _320_v56;
        _320_v56 = _dafny.Seq.of(_319_v55, _319_v55);
        (_this).f11 = (new BigNumber(((_319_v55).Union((_320_v56)[_module.__default.safeIndex(_this.f11, new BigNumber((_320_v56).length))])).cardinality())).minus(p0);
      }
      return;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f8, f9) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), false, false)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false))));
    };
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber(-645)))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D2.create_DC6(new BigNumber(681), true)).dtor_cf12)))).cardinality())));
    };
    fm18(p0, globalState) {
      let _this = this;
      return !((function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(216),!(true)))).Elements) {
          let _321_v1 = _compr_10;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(216),!(true)))).contains(_321_v1)) {
            _coll10.add(_321_v1);
          }
        }
        return _coll10;
      }()).IsSubsetOf((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(531),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-595),true))).Difference(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f8).length),true), function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(114), new BigNumber(409))) {
          let _322_v0 = _compr_11;
          if (((new BigNumber(114)).isLessThanOrEqualTo(_322_v0)) && ((_322_v0).isLessThan(new BigNumber(409)))) {
            _coll11.push([_module.__default.safeDivisionInt(_322_v0, new BigNumber(736)),true]);
          }
        }
        return _coll11;
      }()))));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _323_v0;
      _323_v0 = new BigNumber(758);
      let _324_v1;
      _324_v1 = true;
      _323_v0 = (_323_v0).minus(_module.__default.fm19(_324_v1, _323_v0, globalState));
      let _325_v2;
      let _nw58 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _325_v2 = _nw58;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_325_v2).length))) {
        let _326_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_326_i0)) && ((_326_i0).isLessThan(new BigNumber((_325_v2).length))))) {
          (_325_v2)[(_326_i0)] = (_326_i0).minus(((_324_v1) ? (_323_v0) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(570)), function (_327_i1) {
            return (_this).f8;
          })).length))));
        }
      }
      let _328_v3;
      let _init4 = ((_329_v0) => function (_330_i2) {
        return _dafny.Set.fromElements(_329_v0);
      })(_323_v0);
      let _nw59 = Array((new BigNumber(23)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw59.length); _i0_4++) {
        _nw59[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _328_v3 = _nw59;
      let _331_v4;
      _331_v4 = _dafny.Set.fromElements((_dafny.ZERO).minus(_323_v0), _323_v0, _323_v0);
      let _index49 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_328_v3).length));
      (_328_v3)[_index49] = (_331_v4).Union(_331_v4);
      let _index50 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
      (_325_v2)[_index50] = (_dafny.ZERO).minus(_323_v0);
      let _332_v5;
      _332_v5 = new _dafny.CodePoint('g'.codePointAt(0));
      let _333_v6;
      _333_v6 = _dafny.Set.fromElements(_332_v5, _332_v5);
      let _334_v7;
      _334_v7 = _dafny.Map.Empty.slice().updateUnsafe(_333_v6,_325_v2);
      let _335_v8;
      let _nw60 = Array((new BigNumber(24)).toNumber());
      _nw60[(_dafny.ZERO).toNumber()] = ((_324_v1) ? (_325_v2) : (_325_v2));
      _nw60[(_dafny.ONE).toNumber()] = _325_v2;
      _nw60[(new BigNumber(2)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(3)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(4)).toNumber()] = (((_334_v7).contains(_333_v6)) ? ((_334_v7).get(_333_v6)) : (_325_v2));
      _nw60[(new BigNumber(5)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(6)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(7)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(8)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(9)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(10)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(11)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(12)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(13)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(14)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(15)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(16)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(17)).toNumber()] = ((_324_v1) ? (_325_v2) : (_325_v2));
      _nw60[(new BigNumber(18)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(19)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(20)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(21)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(22)).toNumber()] = _325_v2;
      _nw60[(new BigNumber(23)).toNumber()] = ((!(_324_v1)) ? (_325_v2) : (_325_v2));
      _335_v8 = _nw60;
      let _index51 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_328_v3).length));
      let _index52 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
      let _rhs29 = _module.__default.fm20(true, (_this).f9, !(_323_v0).isEqualTo(_323_v0), globalState);
      let _rhs30 = _323_v0;
      let _rhs31 = _335_v8;
      let _lhs13 = _328_v3;
      let _lhs14 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_328_v3).length));
      let _lhs15 = _325_v2;
      let _lhs16 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
      _lhs13[_lhs14] = _rhs29;
      _lhs15[_lhs16] = _rhs30;
      _335_v8 = _rhs31;
      let _336_v9;
      _336_v9 = _module.D1.create_DC2(!(_324_v1));
      let _source6 = _336_v9;
      if (_source6.is_DC3) {
        let _337___mcc_h0 = (_source6).cf5;
        let _338_cf5 = _337___mcc_h0;
        let _339_v10;
        let _nw61 = Array((new BigNumber(15)).toNumber()).fill(false);
        _339_v10 = _nw61;
        let _340_v11;
        _340_v11 = _dafny.Map.Empty.slice().updateUnsafe(_339_v10,_324_v1);
        let _341_v12;
        _341_v12 = _module.D4.create_DC10(true);
        let _index53 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        (_325_v2)[_index53] = new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_339_v10,true)).Merge(_340_v11)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_339_v10,(_341_v12).dtor_cf16))).length);
        let _index54 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        (_325_v2)[_index54] = (_338_cf5).plus(new BigNumber(908));
        let _342_v13;
        _342_v13 = _dafny.Seq.UnicodeFromString("ndj");
        let _343_v14;
        _343_v14 = _module.D0.create_DC0((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], _dafny.Seq.UnicodeFromString("flum"), _335_v8);
        let _344_v15;
        _344_v15 = _dafny.Seq.of(_324_v1, _324_v1);
        let _345_v16;
        _345_v16 = _dafny.MultiSet.fromElements(_324_v1);
        let _346_v17;
        _346_v17 = _dafny.Seq.of((_dafny.MultiSet.FromArray(_344_v15)).Intersect(_345_v16), _345_v16, _dafny.MultiSet.fromElements(_324_v1));
        let _index55 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        let _rhs32 = _dafny.Seq.update(_dafny.Seq.Concat((_343_v14).dtor_cf1, _342_v13), _module.__default.safeIndex(_323_v0, new BigNumber((_dafny.Seq.Concat((_343_v14).dtor_cf1, _342_v13)).length)), _332_v5);
        let _rhs33 = new BigNumber((_346_v17).length);
        let _lhs17 = _325_v2;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        _342_v13 = _rhs32;
        _lhs17[_lhs18] = _rhs33;
        r0 = _324_v1;
      } else if (_source6.is_DC4) {
        let _347___mcc_h1 = (_source6).cf6;
        let _348___mcc_h2 = (_source6).cf7;
        let _349___mcc_h3 = (_source6).cf8;
        let _350___mcc_h4 = (_source6).cf9;
        let _351_cf9 = _350___mcc_h4;
        let _352_cf8 = _349___mcc_h3;
        let _353_cf7 = _348___mcc_h2;
        let _354_cf6 = _347___mcc_h1;
        _354_cf6 = _323_v0;
        let _355_v18;
        _355_v18 = _dafny.Seq.of(_352_cf8);
        let _356_v19;
        _356_v19 = _dafny.Map.Empty.slice().updateUnsafe(_353_cf7,_324_v1);
        let _357_v20;
        _357_v20 = _dafny.MultiSet.fromElements(new BigNumber(121), new BigNumber((_356_v19).length), _354_cf6);
        let _358_v21;
        _358_v21 = _dafny.Map.Empty.slice().updateUnsafe(_354_cf6,!(_352_cf8));
        let _359_v22;
        let _nw62 = new _module.C1();
        _nw62.__ctor((_this).fm0(_352_cf8, !((_this).fm18(_355_v18, globalState)), new BigNumber((_357_v20).cardinality()), new BigNumber((_358_v21).length), globalState), _323_v0);
        _359_v22 = _nw62;
        let _index56 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        let _rhs34 = _dafny.Seq.IsProperPrefixOf(_355_v18, _dafny.Seq.of(_324_v1, _352_cf8, _324_v1));
        let _rhs35 = _323_v0;
        let _lhs19 = _325_v2;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        _352_cf8 = _rhs34;
        _lhs19[_lhs20] = _rhs35;
        let _360_v23;
        let _nw63 = Array((new BigNumber(28)).toNumber()).fill(_module.D2.Default());
        _360_v23 = _nw63;
        let _361_v24;
        _361_v24 = _module.D2.create_DC6(_354_cf6, _324_v1);
        let _index57 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_360_v23).length));
        (_360_v23)[_index57] = _361_v24;
        let _index58 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_360_v23).length));
        (_360_v23)[_index58] = _361_v24;
      } else {
        let _362___mcc_h5 = (_source6).cf4;
        let _363_cf4 = _362___mcc_h5;
        if (_324_v1) {
          let _364_v25;
          _364_v25 = _dafny.Map.Empty.slice().updateUnsafe(_325_v2,new BigNumber((_module.__default.fm28(new BigNumber((_module.__default.fm29(globalState)).length), (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState)).length));
          let _365_v26;
          _365_v26 = _dafny.Seq.of(_364_v25);
          _364_v25 = ((_365_v26)[_module.__default.safeIndex(_323_v0, new BigNumber((_365_v26).length))]).Merge(_364_v25);
          _323_v0 = _323_v0;
          let _366_v27;
          _366_v27 = _module.D0.create_DC0(new BigNumber(338), _dafny.Seq.Concat((_this).f8, (_this).f9), _335_v8);
          let _367_v28;
          _367_v28 = _dafny.Map.Empty.slice().updateUnsafe(_332_v5,_323_v0);
          _366_v27 = _module.D0.create_DC0(new BigNumber(((_367_v28).Merge(_367_v28)).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(515)), ((_368_v5) => function (_369_i3) {
  return _368_v5;
})(_332_v5)), _335_v8);
          let _370_v29;
          _370_v29 = _dafny.Seq.of(!(_324_v1), _324_v1);
          let _371_v30;
          _371_v30 = _dafny.Seq.of(_370_v29);
          let _372_v31;
          _372_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(51),_323_v0);
          let _373_v32;
          _373_v32 = _dafny.Seq.of((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], _323_v0, new BigNumber((_372_v31).length), (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))]);
          let _374_v33;
          _374_v33 = _dafny.Map.Empty.slice().updateUnsafe(_363_cf4,new BigNumber(345));
          let _375_v34;
          _375_v34 = _dafny.MultiSet.fromElements(new BigNumber((_374_v33).length));
          let _376_v35;
          let _nw64 = new _module.C0();
          _nw64.__ctor();
          _376_v35 = _nw64;
          let _377_v36;
          _377_v36 = _module.D7.create_DC15(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_370_v29, _dafny.Seq.of(_363_cf4, _324_v1)), _371_v30), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-459)), ((_378_v5) => function (_379_i4) {
  return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-711)), ((_380_v5) => function (_381_i5) {
    return _380_v5;
  })(_378_v5))).length);
})(_332_v5)), _373_v32)).length), new BigNumber((_375_v34).cardinality()), (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], _376_v35);
          let _382_v37;
          _382_v37 = _dafny.MultiSet.fromElements(true);
          let _383_v38;
          let _nw65 = new _module.C1();
          _nw65.__ctor(_382_v37, (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))]);
          _383_v38 = _nw65;
          let _384_v39;
          _384_v39 = _dafny.MultiSet.fromElements(_383_v38, _383_v38);
          let _pat_let_tv3 = _376_v35;
          let _index59 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          let _rhs36 = new BigNumber(((_371_v30)[_module.__default.safeIndex(_323_v0, new BigNumber((_371_v30).length))]).length);
          let _rhs37 = ((!((_dafny.MultiSet.fromElements(_383_v38)).IsDisjointFrom(_384_v39))) ? (function (_pat_let4_0) {
            return function (_385_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_386_dt__update_hcf25_h0) {
                  return _module.D7.create_DC15((_385_dt__update__tmp_h0).dtor_cf21, (_385_dt__update__tmp_h0).dtor_cf22, (_385_dt__update__tmp_h0).dtor_cf23, (_385_dt__update__tmp_h0).dtor_cf24, _386_dt__update_hcf25_h0);
                }(_pat_let5_0);
              }(_pat_let_tv3);
            }(_pat_let4_0);
          }(_377_v36)) : (_377_v36));
          let _lhs21 = _325_v2;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          _lhs21[_lhs22] = _rhs36;
          _377_v36 = _rhs37;
          _323_v0 = _module.__default.safeDivisionInt((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], _module.__default.fm19(_363_cf4, (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState));
        } else {
          _324_v1 = !(_363_cf4);
          let _387_v40;
          let _nw66 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _387_v40 = _nw66;
          let _index60 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          let _index61 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          let _rhs38 = (_module.__default.safeDivisionInt((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))])).multipliedBy(((_dafny.ZERO).minus((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))])).minus((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))]));
          let _rhs39 = new BigNumber(((_this).f9).length);
          let _rhs40 = ((_363_cf4) ? (_387_v40) : (_387_v40));
          let _rhs41 = _module.__default.safeDivisionInt(_323_v0, _module.__default.fm24((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], (_this).fm17(_332_v5, _363_cf4, (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], _module.__default.fm24(_323_v0, _324_v1, !(_363_cf4), (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState), globalState), !(!(_363_cf4)), new BigNumber(((_this).f9).length), globalState));
          let _lhs23 = _325_v2;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          let _lhs25 = _325_v2;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          _323_v0 = _rhs38;
          _lhs23[_lhs24] = _rhs39;
          _387_v40 = _rhs40;
          _lhs25[_lhs26] = _rhs41;
          let _388_v41;
          _388_v41 = _dafny.Seq.of(new BigNumber(819));
          let _index62 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_387_v40).length));
          (_387_v40)[_index62] = _388_v41;
          let _389_v42;
          _389_v42 = _dafny.Set.fromElements(_324_v1, _324_v1, _324_v1);
          let _index63 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_387_v40).length));
          (_387_v40)[_index63] = _dafny.Seq.update(_dafny.Seq.update(_module.__default.fm23((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState), _module.__default.safeIndex(new BigNumber(((_389_v42).Difference(_389_v42)).length), new BigNumber((_module.__default.fm23((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState)).length)), new BigNumber(778)), _module.__default.safeIndex((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], new BigNumber((_dafny.Seq.update(_module.__default.fm23((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState), _module.__default.safeIndex(new BigNumber(((_389_v42).Difference(_389_v42)).length), new BigNumber((_module.__default.fm23((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState)).length)), new BigNumber(778))).length)), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_this).f8).length)), _dafny.Seq.of(_323_v0, _323_v0, new BigNumber(766))), _module.__default.safeIndex(_323_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_this).f8).length)), _dafny.Seq.of(_323_v0, _323_v0, new BigNumber(766)))).length)), (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))])).length));
          let _390_v43;
          let _nw67 = new _module.C0();
          _nw67.__ctor();
          _390_v43 = _nw67;
          let _391_v44;
          _391_v44 = _dafny.Map.Empty.slice().updateUnsafe(_324_v1,_324_v1);
          let _392_v45;
          _392_v45 = _module.D3.create_DC8(new BigNumber((_391_v44).length));
          let _393_v46;
          _393_v46 = _dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC15(_324_v1, new BigNumber(((_328_v3)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_328_v3).length))]).length), (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], _323_v0, _390_v43)).dtor_cf24,(_392_v45).dtor_cf14);
          let _index64 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          let _index65 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_387_v40).length));
          let _rhs42 = (_dafny.ZERO).minus(_323_v0);
          let _rhs43 = ((_363_cf4) ? (_dafny.Seq.of((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))])) : ((_387_v40)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_387_v40).length))]));
          let _rhs44 = true;
          let _rhs45 = ((((_dafny.ZERO).minus(_323_v0)).isLessThanOrEqualTo(new BigNumber((_393_v46).length))) ? (_324_v1) : (true));
          let _lhs27 = _325_v2;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
          let _lhs29 = _387_v40;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_387_v40).length));
          _lhs27[_lhs28] = _rhs42;
          _lhs29[_lhs30] = _rhs43;
          _324_v1 = _rhs44;
          _324_v1 = _rhs45;
          _363_cf4 = !(_324_v1);
        }
        let _394_v47;
        let _init5 = ((_395_cf4) => function (_396_i6) {
          return _dafny.Seq.of(_395_cf4, _395_cf4, _395_cf4);
        })(_363_cf4);
        let _nw68 = Array((new BigNumber(10)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw68.length); _i0_5++) {
          _nw68[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _394_v47 = _nw68;
        let _397_v48;
        _397_v48 = _dafny.Seq.of(_324_v1, _363_cf4, _363_cf4, false, _324_v1);
        let _index66 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_394_v47).length));
        (_394_v47)[_index66] = _397_v48;
        let _index67 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_394_v47).length));
        (_394_v47)[_index67] = _dafny.Seq.update(_module.__default.fm26((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState), _module.__default.safeIndex(new BigNumber(577), new BigNumber((_module.__default.fm26((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], globalState)).length)), _324_v1);
        let _398_v49;
        _398_v49 = _dafny.Seq.of((_module.__default.fm24(_323_v0, _363_cf4, false, new BigNumber(((_this).f9).length), globalState)).minus(_323_v0));
        _323_v0 = new BigNumber((_398_v49).length);
        let _399_v50;
        _399_v50 = _dafny.Seq.UnicodeFromString("tsqb");
        _399_v50 = (_this).f8;
      }
      let _400_v51;
      _400_v51 = _dafny.Seq.of(_324_v1);
      let _401_v53;
      _401_v53 = _dafny.Map.Empty.slice().updateUnsafe((_400_v51)[_module.__default.safeIndex((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], new BigNumber((_400_v51).length))],function () {
        let _coll12 = new _dafny.Set();
        for (const _compr_12 of _dafny.IntegerRange(new BigNumber(-873), new BigNumber(158))) {
          let _402_v52 = _compr_12;
          if (((new BigNumber(-873)).isLessThanOrEqualTo(_402_v52)) && ((_402_v52).isLessThan(new BigNumber(158)))) {
            _coll12.add((_402_v52).plus(new BigNumber((_400_v51).length)));
          }
        }
        return _coll12;
      }());
      let _403_v55;
      _403_v55 = _dafny.Seq.of(_dafny.Set.fromElements(_323_v0, (_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))]), (((_401_v53).contains(_324_v1)) ? ((_401_v53).get(_324_v1)) : (_331_v4)), (_328_v3)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_328_v3).length))], function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of _dafny.IntegerRange(new BigNumber(53), new BigNumber(673))) {
          let _404_v54 = _compr_13;
          if (((new BigNumber(53)).isLessThanOrEqualTo(_404_v54)) && ((_404_v54).isLessThan(new BigNumber(673)))) {
            _coll13.add((_404_v54).minus(_323_v0));
          }
        }
        return _coll13;
      }());
      let _hi3 = _323_v0;
      for (let _405_i7 = new BigNumber(((_403_v55)[_module.__default.safeIndex((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))], new BigNumber((_403_v55).length))]).length); _405_i7.isLessThan(_hi3); _405_i7 = _405_i7.plus(_dafny.ONE)) {
        _324_v1 = _324_v1;
        let _index68 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        (_325_v2)[_index68] = (_323_v0).minus(_405_i7);
        let _406_v57;
        let _init6 = ((_407_i7, _408_v1, _409_v2, _410_v0) => function (_411_i8) {
          return (function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_408_v1,(_409_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_409_v2).length))]), _dafny.Map.Empty.slice().updateUnsafe(_408_v1,_410_v0), _dafny.Map.Empty.slice().updateUnsafe(_408_v1,new BigNumber((_dafny.Seq.of((_this).f8, (_this).f9)).length)))).Elements) {
              let _412_v56 = _compr_14;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_408_v1,(_409_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_409_v2).length))]), _dafny.Map.Empty.slice().updateUnsafe(_408_v1,_410_v0), _dafny.Map.Empty.slice().updateUnsafe(_408_v1,new BigNumber((_dafny.Seq.of((_this).f8, (_this).f9)).length))), _412_v56)) {
                _coll14.add(_412_v56);
              }
            }
            return _coll14;
          }()).IsDisjointFrom(_dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(_408_v1,_407_i7))));
        })(_405_i7, _324_v1, _325_v2, _323_v0);
        let _nw69 = Array((new BigNumber(26)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw69.length); _i0_6++) {
          _nw69[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _406_v57 = _nw69;
        _406_v57 = _406_v57;
        let _index69 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length));
        (_325_v2)[_index69] = _323_v0;
      }
      _323_v0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_332_v5, new _dafny.CodePoint('j'.codePointAt(0))), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), ((_413_v5) => function (_414_i9) {
        return _413_v5;
      })(_332_v5)), _module.__default.safeIndex(_323_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), ((_415_v5) => function (_416_i9) {
        return _415_v5;
      })(_332_v5))).length)), _332_v5), (_this).f8))).length);
      r0 = _324_v1;
      r1 = _dafny.Map.Empty.slice().updateUnsafe((_325_v2)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_325_v2).length))],_dafny.MultiSet.fromElements(_324_v1, _324_v1, _324_v1, _324_v1, _324_v1));
      return [r0, r1];
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

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((!(!(false))) && (!(!(false)))) {
        return _dafny.MultiSet.fromElements(!(true));
      } else {
        return _dafny.MultiSet.fromElements(false);
      }
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _417_v0;
      _417_v0 = new BigNumber(131);
      let _418_v1;
      _418_v1 = _dafny.MultiSet.fromElements(_417_v0, new BigNumber(-663), _417_v0);
      let _419_v2;
      _419_v2 = _module.D3.create_DC7(_418_v1);
      _419_v2 = _419_v2;
      let _420_v3;
      _420_v3 = true;
      let _421_v4;
      _421_v4 = _module.D4.create_DC10(_420_v3);
      _421_v4 = _421_v4;
      if (_420_v3) {
        let _422_v5;
        _422_v5 = _dafny.Set.fromElements(_420_v3);
        let _423_v6;
        _423_v6 = _dafny.Seq.of(_420_v3, _420_v3);
        let _424_v7;
        _424_v7 = _dafny.Set.fromElements(_417_v0, _module.__default.fm14(_module.__default.fm14(_417_v0, new BigNumber((_422_v5).length), globalState), _417_v0, globalState), new BigNumber((_423_v6).length), (_417_v0).multipliedBy(_module.__default.fm14(_417_v0, _417_v0, globalState)), (_dafny.ZERO).minus(_417_v0));
        _424_v7 = _424_v7;
        let _425_v8;
        let _nw70 = Array((new BigNumber(29)).toNumber()).fill(false);
        _425_v8 = _nw70;
        let _426_v9;
        _426_v9 = _dafny.Map.Empty.slice().updateUnsafe(_425_v8,_420_v3);
        _426_v9 = (_426_v9).update(_425_v8, _420_v3);
        let _427_v10;
        _427_v10 = _dafny.Seq.of(_423_v6);
        let _428_v11;
        _428_v11 = new _dafny.CodePoint('q'.codePointAt(0));
        let _429_v12;
        let _init7 = ((_430_v0) => function (_431_i1) {
          return _module.__default.safeDivisionInt(_431_i1, _430_v0);
        })(_417_v0);
        let _nw71 = Array((new BigNumber(11)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw71.length); _i0_7++) {
          _nw71[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _429_v12 = _nw71;
        let _432_v13;
        let _nw72 = Array((new BigNumber(15)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _429_v12;
        _nw72[(_dafny.ONE).toNumber()] = _429_v12;
        _nw72[(new BigNumber(2)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(3)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(4)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(5)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(6)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(7)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(8)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(9)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(10)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(11)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(12)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(13)).toNumber()] = _429_v12;
        _nw72[(new BigNumber(14)).toNumber()] = _429_v12;
        _432_v13 = _nw72;
        let _433_v14;
        _433_v14 = _module.D0.create_DC0(_417_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(642)), ((_434_v11) => function (_435_i0) {
  return _434_v11;
})(_428_v11)), _432_v13);
        let _436_v15;
        _436_v15 = _dafny.Map.Empty.slice().updateUnsafe(_428_v11,_433_v14);
        let _437_v16;
        _437_v16 = _module.D0.create_DC1((((_436_v15).contains(new _dafny.CodePoint('h'.codePointAt(0)))) ? ((_436_v15).get(new _dafny.CodePoint('h'.codePointAt(0)))) : (_433_v14)));
        let _438_v17;
        _438_v17 = _dafny.Map.Empty.slice().updateUnsafe(_427_v10,_437_v16);
        _438_v17 = (_438_v17).update(_427_v10, _437_v16);
        _417_v0 = _417_v0;
        let _439_v18;
        let _nw73 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
        _439_v18 = _nw73;
        let _440_v19;
        _440_v19 = _dafny.MultiSet.fromElements(_420_v3);
        let _441_v20;
        _441_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_440_v19).cardinality())),_428_v11);
        let _index70 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_439_v18).length));
        (_439_v18)[_index70] = (_441_v20).Merge(_441_v20);
        let _442_v22;
        _442_v22 = _dafny.Map.Empty.slice().updateUnsafe(_425_v8,function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(11), new BigNumber(664))) {
            let _443_v21 = _compr_15;
            if (((new BigNumber(11)).isLessThanOrEqualTo(_443_v21)) && ((_443_v21).isLessThan(new BigNumber(664)))) {
              _coll15.push([(_443_v21).plus(_417_v0),new _dafny.CodePoint('m'.codePointAt(0))]);
            }
          }
          return _coll15;
        }());
        let _index71 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_439_v18).length));
        (_439_v18)[_index71] = (((_442_v22).contains(((false) ? (_425_v8) : (_425_v8)))) ? ((_442_v22).get(((false) ? (_425_v8) : (_425_v8)))) : ((_441_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_417_v0),_428_v11))));
      } else {
        let _444_v23;
        let _nw74 = Array((new BigNumber(17)).toNumber()).fill(false);
        _444_v23 = _nw74;
        let _index72 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
        (_444_v23)[_index72] = _420_v3;
        let _index73 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
        (_444_v23)[_index73] = !((_418_v1).IsProperSubsetOf(_418_v1));
        let _445_v24;
        _445_v24 = new _dafny.CodePoint('a'.codePointAt(0));
        let _446_v25;
        _446_v25 = _dafny.Seq.of(_445_v24);
        let _447_v26;
        _447_v26 = _dafny.Map.Empty.slice().updateUnsafe(_444_v23,_446_v25);
        _446_v25 = (((_447_v26).contains(_444_v23)) ? ((_447_v26).get(_444_v23)) : (_446_v25));
        _417_v0 = _module.__default.fm14(_417_v0, _417_v0, globalState);
        let _448_v27;
        _448_v27 = _dafny.Map.Empty.slice().updateUnsafe((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))],_dafny.Seq.Concat(_446_v25, _dafny.Seq.update(_446_v25, _module.__default.safeIndex(_417_v0, new BigNumber((_446_v25).length)), _445_v24)));
        let _449_v28;
        _449_v28 = _dafny.Seq.of(_420_v3);
        let _450_v29;
        _450_v29 = _module.D1.create_DC3(new BigNumber((_449_v28).length));
        let _451_v30;
        _451_v30 = _dafny.MultiSet.fromElements((_421_v4).dtor_cf16);
        let _index74 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
        let _rhs46 = _448_v27;
        let _rhs47 = ((_dafny.ZERO).minus((_450_v29).dtor_cf5)).isEqualTo(new BigNumber(431));
        let _rhs48 = _dafny.Seq.UnicodeFromString("atwdm");
        let _rhs49 = ((((_451_v30).contains(true)) ? ((_451_v30).get(true)) : (_417_v0))).multipliedBy(_417_v0);
        let _lhs31 = _444_v23;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
        _448_v27 = _rhs46;
        _lhs31[_lhs32] = _rhs47;
        _446_v25 = _rhs48;
        _417_v0 = _rhs49;
        let _452_v31;
        _452_v31 = _module.D1.create_DC2(true);
        let _source7 = _452_v31;
        if (_source7.is_DC3) {
          let _453___mcc_h0 = (_source7).cf5;
          let _454_cf5 = _453___mcc_h0;
          let _index75 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          (_444_v23)[_index75] = (_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))];
          let _455_v32;
          let _init8 = function (_456_i2) {
            return (_456_i2).multipliedBy(new BigNumber(-103));
          };
          let _nw75 = Array((new BigNumber(7)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw75.length); _i0_8++) {
            _nw75[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _455_v32 = _nw75;
          _455_v32 = _455_v32;
          let _457_v33;
          _457_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(650),_417_v0);
          let _458_v34;
          _458_v34 = _dafny.Map.Empty.slice().updateUnsafe((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))],false);
          let _459_v35;
          _459_v35 = _module.D1.create_DC4((_dafny.ZERO).minus(_417_v0), _445_v24, (_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))], _417_v0);
          let _rhs50 = _417_v0;
          let _rhs51 = new BigNumber(754);
          let _rhs52 = ((((_457_v33).contains(new BigNumber((_458_v34).length))) ? ((_457_v33).get(new BigNumber((_458_v34).length))) : (_417_v0))).minus(_454_cf5);
          let _rhs53 = (_459_v35).dtor_cf9;
          _454_cf5 = _rhs50;
          _454_cf5 = _rhs51;
          _454_cf5 = _rhs52;
          _417_v0 = _rhs53;
          let _460_v36;
          _460_v36 = _dafny.Seq.of(_444_v23, _444_v23);
          let _461_v37;
          _461_v37 = _dafny.Set.fromElements(_417_v0, _417_v0, _417_v0, new BigNumber((_dafny.Seq.update(_449_v28, _module.__default.safeIndex(_417_v0, new BigNumber((_449_v28).length)), (_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))])).length));
          let _462_v38;
          _462_v38 = _dafny.Set.fromElements((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))], (_421_v4).dtor_cf16);
          let _index76 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          let _rhs54 = !(((_420_v3) ? (_461_v37) : (_461_v37))).contains((new BigNumber(423)).plus(new BigNumber((_457_v33).length)));
          let _rhs55 = (_417_v0).plus(_417_v0);
          let _rhs56 = ((_462_v38).Union(_462_v38)).IsSubsetOf(_dafny.Set.fromElements(_420_v3));
          let _rhs57 = _460_v36;
          let _lhs33 = _444_v23;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          let _lhs35 = _444_v23;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          _lhs33[_lhs34] = _rhs54;
          _417_v0 = _rhs55;
          _lhs35[_lhs36] = _rhs56;
          _460_v36 = _rhs57;
        } else if (_source7.is_DC4) {
          let _463___mcc_h1 = (_source7).cf6;
          let _464___mcc_h2 = (_source7).cf7;
          let _465___mcc_h3 = (_source7).cf8;
          let _466___mcc_h4 = (_source7).cf9;
          let _467_cf9 = _466___mcc_h4;
          let _468_cf8 = _465___mcc_h3;
          let _469_cf7 = _464___mcc_h2;
          let _470_cf6 = _463___mcc_h1;
          let _index78 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          (_444_v23)[_index78] = !(!(_468_cf8));
          let _471_v39;
          let _nw76 = Array((new BigNumber(11)).toNumber()).fill([]);
          _471_v39 = _nw76;
          let _472_v40;
          _472_v40 = _dafny.Map.Empty.slice().updateUnsafe((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))],_420_v3);
          let _473_v41;
          let _nw77 = Array((new BigNumber(21)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = _472_v40;
          _nw77[(_dafny.ONE).toNumber()] = _472_v40;
          _nw77[(new BigNumber(2)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(3)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(4)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(5)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_420_v3,(_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))]);
          _nw77[(new BigNumber(7)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(8)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(9)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))],(_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))]);
          _nw77[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_420_v3),_420_v3);
          _nw77[(new BigNumber(12)).toNumber()] = (_472_v40).update(_420_v3, _468_cf8);
          _nw77[(new BigNumber(13)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(14)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(15)).toNumber()] = (_472_v40).update(_420_v3, false);
          _nw77[(new BigNumber(16)).toNumber()] = _module.__default.fm15(_470_cf6, _468_cf8, _469_cf7, _467_cf9, globalState);
          _nw77[(new BigNumber(17)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(18)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(19)).toNumber()] = _472_v40;
          _nw77[(new BigNumber(20)).toNumber()] = _472_v40;
          _473_v41 = _nw77;
          let _index79 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_471_v39).length));
          (_471_v39)[_index79] = _473_v41;
          let _474_v42;
          let _nw78 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _474_v42 = _nw78;
          let _index80 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length));
          (_474_v42)[_index80] = new BigNumber((_446_v25).length);
          let _475_v43;
          _475_v43 = _dafny.Set.fromElements((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))], _468_cf8);
          let _476_v44;
          _476_v44 = _dafny.Seq.of(_dafny.Set.fromElements(_468_cf8, (_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))]), _475_v43);
          let _477_v45;
          let _nw79 = Array((new BigNumber(12)).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = (_475_v43).Intersect(_475_v43);
          _nw79[(_dafny.ONE).toNumber()] = (_476_v44)[_module.__default.safeIndex(new BigNumber((_449_v28).length), new BigNumber((_476_v44).length))];
          _nw79[(new BigNumber(2)).toNumber()] = _475_v43;
          _nw79[(new BigNumber(3)).toNumber()] = (_475_v43).Intersect(_475_v43);
          _nw79[(new BigNumber(4)).toNumber()] = (_475_v43).Difference(_475_v43);
          _nw79[(new BigNumber(5)).toNumber()] = _475_v43;
          _nw79[(new BigNumber(6)).toNumber()] = (_module.__default.fm16(_421_v4, globalState)).Union(_475_v43);
          _nw79[(new BigNumber(7)).toNumber()] = _475_v43;
          _nw79[(new BigNumber(8)).toNumber()] = _475_v43;
          _nw79[(new BigNumber(9)).toNumber()] = _475_v43;
          _nw79[(new BigNumber(10)).toNumber()] = _475_v43;
          _nw79[(new BigNumber(11)).toNumber()] = (_475_v43).Difference(_dafny.Set.fromElements(_420_v3));
          _477_v45 = _nw79;
          let _index81 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_477_v45).length));
          (_477_v45)[_index81] = _475_v43;
          let _478_v46;
          _478_v46 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_467_cf9),_473_v41);
          let _index82 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_471_v39).length));
          let _index83 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length));
          let _index84 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_477_v45).length));
          let _rhs58 = (((_478_v46).contains(new BigNumber(-336))) ? ((_478_v46).get(new BigNumber(-336))) : (_473_v41));
          let _rhs59 = _417_v0;
          let _rhs60 = _module.__default.fm16(_421_v4, globalState);
          let _lhs37 = _471_v39;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_471_v39).length));
          let _lhs39 = _474_v42;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length));
          let _lhs41 = _477_v45;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_477_v45).length));
          _lhs37[_lhs38] = _rhs58;
          _lhs39[_lhs40] = _rhs59;
          _lhs41[_lhs42] = _rhs60;
          let _479_v47;
          _479_v47 = _dafny.Map.Empty.slice().updateUnsafe(_469_cf7,false);
          _417_v0 = ((new BigNumber((_446_v25).length)).multipliedBy(_470_cf6)).multipliedBy((((((_479_v47).contains(_469_cf7)) ? ((_479_v47).get(_469_cf7)) : ((_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))]))) ? ((_474_v42)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length))]) : ((_474_v42)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length))])));
          let _index85 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length));
          let _rhs61 = (_474_v42)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length))];
          let _rhs62 = ((false) ? (_module.__default.safeModuloInt(_470_cf6, (_474_v42)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length))])) : (new BigNumber((_dafny.Seq.UnicodeFromString("ppyxfavq")).length)));
          let _lhs43 = _474_v42;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_474_v42).length));
          _470_cf6 = _rhs61;
          _lhs43[_lhs44] = _rhs62;
        } else {
          let _480___mcc_h5 = (_source7).cf4;
          let _481_cf4 = _480___mcc_h5;
          let _482_v48;
          let _nw80 = new _module.C2();
          _nw80.__ctor(((_420_v3) ? (_446_v25) : (_446_v25)), _446_v25);
          _482_v48 = _nw80;
          let _index86 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          let _index87 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          let _rhs63 = _420_v3;
          let _rhs64 = (_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))];
          let _rhs65 = !((_417_v0).minus(_417_v0)).isEqualTo(_417_v0);
          let _rhs66 = !(_420_v3) || (_dafny.Seq.IsProperPrefixOf((_482_v48).f9, _446_v25));
          let _lhs45 = _444_v23;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          let _lhs47 = _444_v23;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          _481_cf4 = _rhs63;
          _481_cf4 = _rhs64;
          _lhs45[_lhs46] = _rhs65;
          _lhs47[_lhs48] = _rhs66;
          let _index88 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length));
          (_444_v23)[_index88] = (_444_v23)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_444_v23).length))];
          let _483_v49;
          let _nw81 = Array((new BigNumber(6)).toNumber()).fill([]);
          _483_v49 = _nw81;
          let _484_v50;
          let _nw82 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _484_v50 = _nw82;
          let _index89 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_483_v49).length));
          (_483_v49)[_index89] = _484_v50;
          let _485_v51;
          _485_v51 = _dafny.Seq.of(_484_v50, _484_v50, _484_v50);
          let _index90 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_483_v49).length));
          (_483_v49)[_index90] = (_485_v51)[_module.__default.safeIndex(_417_v0, new BigNumber((_485_v51).length))];
        }
      }
      let _486_v52;
      _486_v52 = _dafny.Set.fromElements(_420_v3);
      let _487_i3;
      _487_i3 = _dafny.ZERO;
      L2: {
        while ((_486_v52).IsProperSubsetOf(_486_v52)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_487_i3)) {
              break L2;
            }
            _487_i3 = (_487_i3).plus(_dafny.ONE);
            let _488_v53;
            _488_v53 = _dafny.Map.Empty.slice().updateUnsafe(_420_v3,new BigNumber(-225));
            _488_v53 = _488_v53;
            let _489_v54;
            _489_v54 = _dafny.Seq.UnicodeFromString("pybwqy");
            let _490_v55;
            _490_v55 = _dafny.Map.Empty.slice().updateUnsafe(_417_v0,_489_v54);
            let _491_v56;
            let _nw83 = new _module.C2();
            _nw83.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("poqcmavv"), (((_490_v55).contains(_417_v0)) ? ((_490_v55).get(_417_v0)) : (_489_v54))), _489_v54);
            _491_v56 = _nw83;
            let _492_v57;
            _492_v57 = _dafny.Map.Empty.slice().updateUnsafe(_420_v3,_420_v3);
            _491_v56 = (((((_492_v57).contains(!(_420_v3))) ? ((_492_v57).get(!(_420_v3))) : (false))) ? (_491_v56) : (_491_v56));
            let _493_v58;
            _493_v58 = _dafny.Map.Empty.slice().updateUnsafe(_420_v3,_421_v4);
            _417_v0 = _module.__default.safeDivisionInt(_417_v0, new BigNumber((_493_v58).length));
            _492_v57 = (_492_v57).update(_420_v3, _420_v3);
          }
        }
      }
      let _494_v59;
      _494_v59 = _dafny.MultiSet.fromElements(_420_v3, _420_v3);
      let _495_v60;
      _495_v60 = _dafny.Seq.of(_420_v3, _420_v3);
      let _496_v61;
      let _nw84 = new _module.C1();
      _nw84.__ctor(_494_v59, (new BigNumber((_495_v60).length)).multipliedBy(new BigNumber((_module.__default.fm16(_421_v4, globalState)).length)));
      _496_v61 = _nw84;
      if (_420_v3) {
        let _497_v62;
        let _nw85 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _497_v62 = _nw85;
        let _498_v63;
        _498_v63 = _dafny.MultiSet.fromElements(_497_v62, _497_v62);
        _498_v63 = _498_v63;
        let _499_v64;
        _499_v64 = new _dafny.CodePoint('v'.codePointAt(0));
        let _500_v65;
        _500_v65 = _dafny.Seq.UnicodeFromString("ned");
        let _rhs67 = _module.__default.fm30(_500_v65, _496_v61.f11, _496_v61.f11, globalState);
        let _rhs68 = _499_v64;
        let _rhs69 = !(_420_v3);
        let _rhs70 = _496_v61.f11;
        let _lhs49 = _496_v61;
        _499_v64 = _rhs67;
        _499_v64 = _rhs68;
        _420_v3 = _rhs69;
        _lhs49.f11 = _rhs70;
        let _501_v66;
        _501_v66 = _dafny.Seq.of(_496_v61.f11);
        _501_v66 = _dafny.Seq.Concat(_501_v66, _dafny.Seq.Concat(_501_v66, _501_v66));
        (_496_v61).f11 = _496_v61.f11;
        (_496_v61).f11 = (_496_v61.f11).multipliedBy(_module.__default.safeModuloInt(_496_v61.f11, _417_v0));
      } else {
        let _502_v67;
        _502_v67 = _dafny.Seq.UnicodeFromString("bwgipadpu");
        (_496_v61).f11 = _module.__default.fm14(_417_v0, new BigNumber((_502_v67).length), globalState);
        let _503_v68;
        _503_v68 = new _dafny.CodePoint('a'.codePointAt(0));
        let _504_v69;
        _504_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_502_v67, _502_v67), _module.__default.safeIndex(_417_v0, new BigNumber((_dafny.Seq.Concat(_502_v67, _502_v67)).length)), _503_v68),_420_v3);
        _504_v69 = (_504_v69).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(962)), function (_505_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _420_v3);
        if (((!(_420_v3)) ? (!(_dafny.Seq.contains(_502_v67, _503_v68))) : (!(!((_module.__default.fm31(new BigNumber((_502_v67).length), globalState)).IsDisjointFrom(_418_v1)))))) {
          let _506_v70;
          let _nw86 = new _module.C0();
          _nw86.__ctor();
          _506_v70 = _nw86;
          let _507_v71;
          let _nw87 = new _module.C1();
          _nw87.__ctor((_496_v61).fm0(_420_v3, _420_v3, new BigNumber((_486_v52).length), new BigNumber(-609), globalState), _496_v61.f11);
          _507_v71 = _nw87;
          let _508_v72;
          _508_v72 = _dafny.Map.Empty.slice().updateUnsafe(_503_v68,_496_v61.f11);
          let _509_v73;
          _509_v73 = _dafny.Set.fromElements(_502_v67);
          let _510_v74;
          _510_v74 = _module.D4.create_DC10((_496_v61).fm21(_508_v72, new BigNumber((_509_v73).length), new BigNumber(324), globalState));
          let _511_v75;
          _511_v75 = _module.D4.create_DC11(_510_v74);
          _511_v75 = _511_v75;
          let _512_v76;
          let _nw88 = Array((new BigNumber(17)).toNumber());
          _512_v76 = _nw88;
          let _513_v77;
          let _nw89 = new _module.C2();
          _nw89.__ctor(_dafny.Seq.UnicodeFromString("gcchob"), _dafny.Seq.UnicodeFromString("f"));
          _513_v77 = _nw89;
          let _index91 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_512_v76).length));
          (_512_v76)[_index91] = _513_v77;
          let _index92 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_512_v76).length));
          (_512_v76)[_index92] = _513_v77;
          let _514_v78;
          _514_v78 = _dafny.Map.Empty.slice().updateUnsafe(_496_v61.f11,_496_v61.f11);
          let _515_v79;
          _515_v79 = _dafny.Map.Empty.slice().updateUnsafe(_514_v78,_486_v52);
          let _516_v80;
          _516_v80 = _dafny.Set.fromElements(_496_v61.f11, new BigNumber(((_496_v61).f10).cardinality()));
          let _517_v81;
          _517_v81 = _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_495_v60).length),new BigNumber((_516_v80).length)),_dafny.Set.fromElements(_420_v3, _420_v3)));
          let _518_v82;
          _518_v82 = _dafny.Map.Empty.slice().updateUnsafe(_507_v71.f11,_515_v79);
          let _519_v83;
          _519_v83 = _dafny.Map.Empty.slice().updateUnsafe(_420_v3,_516_v80);
          let _520_v84;
          _520_v84 = _dafny.Map.Empty.slice().updateUnsafe(_420_v3,(((_519_v83).contains(true)) ? ((_519_v83).get(true)) : (_516_v80)));
          let _521_v86;
          _521_v86 = _dafny.Map.Empty.slice().updateUnsafe(true,_486_v52);
          let _522_v87;
          _522_v87 = _dafny.Seq.of(_515_v79);
          let _523_v89;
          _523_v89 = _dafny.Set.fromElements(_514_v78);
          let _pat_let_tv4 = _515_v79;
          let _524_v90;
          let _nw90 = Array((new BigNumber(25)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _515_v79;
          _nw90[(_dafny.ONE).toNumber()] = _515_v79;
          _nw90[(new BigNumber(2)).toNumber()] = (_517_v81).dtor_cf28;
          _nw90[(new BigNumber(3)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(4)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(5)).toNumber()] = (_515_v79).Merge(_515_v79);
          _nw90[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_514_v78,_486_v52);
          _nw90[(new BigNumber(7)).toNumber()] = (_515_v79).Merge(_515_v79);
          _nw90[(new BigNumber(8)).toNumber()] = ((_420_v3) ? (_515_v79) : (_515_v79));
          _nw90[(new BigNumber(9)).toNumber()] = (_515_v79).Merge(_515_v79);
          _nw90[(new BigNumber(10)).toNumber()] = (((_518_v82).contains(new BigNumber(((((_519_v83).contains(_420_v3)) ? ((_519_v83).get(_420_v3)) : (_dafny.Set.fromElements(_496_v61.f11, _module.__default.fm19(_420_v3, _496_v61.f11, globalState))))).length))) ? ((_518_v82).get(new BigNumber(((((_519_v83).contains(_420_v3)) ? ((_519_v83).get(_420_v3)) : (_dafny.Set.fromElements(_496_v61.f11, _module.__default.fm19(_420_v3, _496_v61.f11, globalState))))).length))) : (_515_v79));
          _nw90[(new BigNumber(11)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(12)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(13)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(14)).toNumber()] = (_module.__default.fm32(_420_v3, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_417_v0,_496_v61.f11),_486_v52));
          _nw90[(new BigNumber(15)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_514_v78,_486_v52);
          _nw90[(new BigNumber(17)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_514_v78,_486_v52)).update((_514_v78).update((_dafny.ZERO).minus(_417_v0), _496_v61.f11), _486_v52)).Merge(_515_v79);
          _nw90[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(((((_496_v61).f10).contains(_420_v3)) ? (((_496_v61).f10).get(_420_v3)) : (new BigNumber((function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(187), new BigNumber(166))) {
              let _525_v85 = _compr_16;
              if (((new BigNumber(187)).isLessThanOrEqualTo(_525_v85)) && ((_525_v85).isLessThan(new BigNumber(166)))) {
                _coll16.push([_module.__default.safeModuloInt(_525_v85, _496_v61.f11),_503_v68]);
              }
            }
            return _coll16;
          }()).length))),_417_v0),(((_521_v86).contains(_420_v3)) ? ((_521_v86).get(_420_v3)) : (_486_v52)));
          _nw90[(new BigNumber(19)).toNumber()] = _515_v79;
          _nw90[(new BigNumber(20)).toNumber()] = (_515_v79).Merge((_522_v87)[_module.__default.safeIndex(_496_v61.f11, new BigNumber((_522_v87).length))]);
          _nw90[(new BigNumber(21)).toNumber()] = function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_523_v89).Elements) {
              let _526_v88 = _compr_17;
              if ((_523_v89).contains(_526_v88)) {
                _coll17.push([_526_v88,_486_v52]);
              }
            }
            return _coll17;
          }();
          _nw90[(new BigNumber(22)).toNumber()] = ((function (_pat_let6_0) {
            return function (_527_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_528_dt__update_hcf28_h0) {
                  return _module.D9.create_DC18(_528_dt__update_hcf28_h0);
                }(_pat_let7_0);
              }(_pat_let_tv4);
            }(_pat_let6_0);
          }(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_514_v78,_486_v52)))).dtor_cf28).Merge(_module.__default.fm32(_420_v3, globalState));
          _nw90[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_514_v78,_module.__default.fm16(_421_v4, globalState));
          _nw90[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_514_v78).update(_417_v0, new BigNumber(213)),_486_v52);
          _524_v90 = _nw90;
          let _index93 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_524_v90).length));
          (_524_v90)[_index93] = _515_v79;
          let _index94 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_524_v90).length));
          let _rhs71 = _module.__default.safeDivisionInt((((_514_v78).contains(_417_v0)) ? ((_514_v78).get(_417_v0)) : (_496_v61.f11)), _module.__default.safeDivisionInt(_496_v61.f11, _507_v71.f11));
          let _rhs72 = (_515_v79).Merge(_module.__default.fm32(false, globalState));
          let _lhs50 = _496_v61;
          let _lhs51 = _524_v90;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_524_v90).length));
          _lhs50.f11 = _rhs71;
          _lhs51[_lhs52] = _rhs72;
        } else {
          let _529_v91;
          let _nw91 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _529_v91 = _nw91;
          let _530_v92;
          _530_v92 = _dafny.Map.Empty.slice().updateUnsafe(_496_v61.f11,_529_v91);
          let _531_v93;
          _531_v93 = _dafny.Seq.of(_529_v91, _529_v91);
          let _532_v94;
          _532_v94 = _dafny.Map.Empty.slice().updateUnsafe(false,_496_v61.f11);
          let _533_v95;
          let _nw92 = Array((new BigNumber(11)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _529_v91;
          _nw92[(_dafny.ONE).toNumber()] = _529_v91;
          _nw92[(new BigNumber(2)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(3)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(4)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(5)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(6)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(7)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(8)).toNumber()] = _529_v91;
          _nw92[(new BigNumber(9)).toNumber()] = (((_530_v92).contains(_496_v61.f11)) ? ((_530_v92).get(_496_v61.f11)) : ((_531_v93)[_module.__default.safeIndex(new BigNumber((_532_v94).length), new BigNumber((_531_v93).length))]));
          _nw92[(new BigNumber(10)).toNumber()] = _529_v91;
          _533_v95 = _nw92;
          let _index95 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_533_v95).length));
          (_533_v95)[_index95] = (_531_v93)[_module.__default.safeIndex(_496_v61.f11, new BigNumber((_531_v93).length))];
          let _index96 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_533_v95).length));
          (_533_v95)[_index96] = (_531_v93)[_module.__default.safeIndex(_496_v61.f11, new BigNumber((_531_v93).length))];
          let _534_v96;
          _534_v96 = _dafny.Map.Empty.slice().updateUnsafe((_496_v61).f10,_417_v0);
          let _535_v97;
          _535_v97 = _dafny.Map.Empty.slice().updateUnsafe(((_534_v96).update((_496_v61).f10, _496_v61.f11)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_494_v59,_417_v0)),(_dafny.ZERO).minus(new BigNumber((_502_v67).length)));
          (_496_v61).f11 = (((_535_v97).contains(_534_v96)) ? ((_535_v97).get(_534_v96)) : (_496_v61.f11));
          _502_v67 = _502_v67;
          (_496_v61).f11 = ((!(_420_v3) || (_420_v3)) ? (_496_v61.f11) : (new BigNumber(393)));
          let _536_v98;
          let _nw93 = Array((new BigNumber(3)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _420_v3;
          _nw93[(_dafny.ONE).toNumber()] = _420_v3;
          _nw93[(new BigNumber(2)).toNumber()] = _420_v3;
          _536_v98 = _nw93;
          let _index97 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_536_v98).length));
          (_536_v98)[_index97] = _420_v3;
          let _index98 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_536_v98).length));
          (_536_v98)[_index98] = (_486_v52).IsSubsetOf(_486_v52);
        }
        let _537_v99;
        let _nw94 = new _module.C0();
        _nw94.__ctor();
        _537_v99 = _nw94;
        _417_v0 = (_417_v0).plus(new BigNumber(((_dafny.MultiSet.fromElements(_420_v3)).Union(((_496_v61).f10).update(false, _module.__default.abs(new BigNumber(506))))).cardinality()));
      }
      r0 = _420_v3;
      let _538_v100;
      _538_v100 = _module.D2.create_DC6(_496_v61.f11, _420_v3);
      let _539_v101;
      _539_v101 = _dafny.Map.Empty.slice().updateUnsafe(_496_v61.f11,_dafny.MultiSet.fromElements((_538_v100).dtor_cf12, false));
      r1 = (_539_v101).update(_module.__default.fm14(_496_v61.f11, (_dafny.ZERO).minus((_dafny.ZERO).minus(_417_v0)), globalState), (_496_v61).f10);
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this.f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f7) {
      let _this = this;
      (_this).f7 = f7;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(false, false, false, true, true))).Difference(_dafny.MultiSet.fromElements(!(!(!(true)))));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _540_v0;
      let _nw95 = new _module.C3();
      _nw95.__ctor();
      _540_v0 = _nw95;
      r0 = !_dafny.Seq.contains(_dafny.Seq.of(_this.f7), new BigNumber(471));
      let _541_v1;
      _541_v1 = _dafny.Seq.UnicodeFromString("jmcwwhb");
      let _542_v2;
      let _nw96 = new _module.C2();
      _nw96.__ctor(_541_v1, _541_v1);
      _542_v2 = _nw96;
      let _543_v3;
      _543_v3 = true;
      let _544_v4;
      _544_v4 = _dafny.Seq.of(true, true, _543_v3);
      r0 = ((_543_v3) ? (_543_v3) : ((_544_v4)[_module.__default.safeIndex(_this.f7, new BigNumber((_544_v4).length))]));
      let _545_v5;
      let _nw97 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _545_v5 = _nw97;
      let _546_v6;
      _546_v6 = _module.D4.create_DC9(_545_v5);
      _546_v6 = ((_543_v3) ? (_546_v6) : (_546_v6));
      let _547_i0;
      _547_i0 = _dafny.ZERO;
      L3: {
        while (true) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_547_i0)) {
              break L3;
            }
            _547_i0 = (_547_i0).plus(_dafny.ONE);
            _541_v1 = (_542_v2).f8;
            let _rhs73 = (_dafny.ZERO).minus((_this.f7).multipliedBy(_module.__default.fm14(_this.f7, _this.f7, globalState)));
            let _rhs74 = new BigNumber(907);
            let _rhs75 = (false) || (_543_v3);
            let _lhs53 = _this;
            let _lhs54 = _this;
            _lhs53.f7 = _rhs73;
            _lhs54.f7 = _rhs74;
            r0 = _rhs75;
            let _548_v7;
            let _nw98 = Array((new BigNumber(27)).toNumber()).fill([]);
            _548_v7 = _nw98;
            let _549_v8;
            let _nw99 = Array((new BigNumber(26)).toNumber());
            _nw99[(_dafny.ZERO).toNumber()] = _545_v5;
            _nw99[(_dafny.ONE).toNumber()] = _545_v5;
            _nw99[(new BigNumber(2)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(3)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(4)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(5)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(6)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(7)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(8)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(9)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(10)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(11)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(12)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(13)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(14)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(15)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(16)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(17)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(18)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(19)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(20)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(21)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(22)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(23)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(24)).toNumber()] = _545_v5;
            _nw99[(new BigNumber(25)).toNumber()] = _545_v5;
            _549_v8 = _nw99;
            let _index99 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_548_v7).length));
            (_548_v7)[_index99] = _549_v8;
            let _index100 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_548_v7).length));
            (_548_v7)[_index100] = _549_v8;
            let _index101 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_545_v5).length));
            (_545_v5)[_index101] = _this.f7;
            let _550_v9;
            _550_v9 = _dafny.MultiSet.fromElements(_543_v3, _543_v3);
            let _551_v10;
            _551_v10 = _dafny.MultiSet.fromElements(_this.f7, (((_550_v9).contains(_543_v3)) ? ((_550_v9).get(_543_v3)) : (_this.f7)));
            let _552_v11;
            _552_v11 = _dafny.Seq.of(_541_v1);
            let _index102 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_545_v5).length));
            let _rhs76 = new BigNumber(613);
            let _rhs77 = _this.f7;
            let _rhs78 = _543_v3;
            let _rhs79 = _module.__default.fm19((_551_v10).IsDisjointFrom(_551_v10), new BigNumber((_552_v11).length), globalState);
            let _lhs55 = _this;
            let _lhs56 = _545_v5;
            let _lhs57 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_545_v5).length));
            let _lhs58 = _this;
            _lhs55.f7 = _rhs76;
            _lhs56[_lhs57] = _rhs77;
            _543_v3 = _rhs78;
            _lhs58.f7 = _rhs79;
          }
        }
      }
      let _553_v12;
      _553_v12 = new _dafny.CodePoint('i'.codePointAt(0));
      r0 = (_542_v2).fm17(_553_v12, _543_v3, _this.f7, _this.f7, globalState);
      let _554_v13;
      _554_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_dafny.MultiSet.FromArray(_544_v4));
      let _555_v14;
      _555_v14 = _dafny.MultiSet.fromElements(_543_v3);
      r1 = ((_554_v13).Merge(_554_v13)).Merge((_554_v13).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_555_v14)));
      return [r0, r1];
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      r0 = p0;
      let _hi4 = _this.f7;
      for (let _556_i0 = p2; _556_i0.isLessThan(_hi4); _556_i0 = _556_i0.plus(_dafny.ONE)) {
        let _557_v0;
        let _nw100 = new _module.C0();
        _nw100.__ctor();
        _557_v0 = _nw100;
        let _558_v1;
        _558_v1 = _module.D7.create_DC15(p0, _this.f7, _556_i0, _this.f7, _557_v0);
        let _559_v2;
        _559_v2 = _dafny.MultiSet.fromElements(p3);
        let _560_v3;
        let _nw101 = new _module.C1();
        _nw101.__ctor(_559_v2, _556_i0);
        _560_v3 = _nw101;
        let _561_v4;
        _561_v4 = _dafny.Map.Empty.slice().updateUnsafe(_560_v3,p3);
        let _562_v5;
        _562_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_561_v4).length));
        let _563_v6;
        _563_v6 = _dafny.Map.Empty.slice().updateUnsafe((p0) === (!(p3)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_558_v1,_562_v5)).length));
        _562_v5 = _563_v6;
        let _564_v7;
        let _nw102 = Array((_dafny.ONE).toNumber()).fill(false);
        _564_v7 = _nw102;
        let _index103 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_564_v7).length));
        (_564_v7)[_index103] = p3;
        let _index104 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_564_v7).length));
        (_564_v7)[_index104] = ((((p3) ? (p3) : (p3))) ? (((_560_v3).f10).contains(p0)) : ((_this.f7).isLessThanOrEqualTo(new BigNumber(-700))));
        let _565_v8;
        let _init9 = function (_566_i1) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), function (_567_i2) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
        };
        let _nw103 = Array((new BigNumber(19)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw103.length); _i0_9++) {
          _nw103[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _565_v8 = _nw103;
        let _568_v9;
        _568_v9 = _dafny.Seq.UnicodeFromString("muqj");
        let _index105 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_565_v8).length));
        (_565_v8)[_index105] = _dafny.Seq.Concat(_568_v9, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("yifcb"), _module.__default.safeIndex(_556_i0, new BigNumber((_dafny.Seq.UnicodeFromString("yifcb")).length)), new _dafny.CodePoint('f'.codePointAt(0))));
        let _index106 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_565_v8).length));
        (_565_v8)[_index106] = ((p0) ? (_568_v9) : (_dafny.Seq.Concat(_568_v9, _568_v9)));
        (_this).f7 = (p2).multipliedBy(p2);
      }
      let _569_v10;
      let _init10 = function (_570_i4) {
        return (_570_i4).minus((_dafny.ZERO).minus(_this.f7));
      };
      let _nw104 = Array((new BigNumber(20)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw104.length); _i0_10++) {
        _nw104[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _569_v10 = _nw104;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_569_v10).length))) {
        let _571_i3 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_571_i3)) && ((_571_i3).isLessThan(new BigNumber((_569_v10).length))))) {
          (_569_v10)[(_571_i3)] = _module.__default.safeDivisionInt(_571_i3, new BigNumber(-97));
        }
      }
      let _572_v11;
      let _nw105 = Array((new BigNumber(4)).toNumber());
      _572_v11 = _nw105;
      let _573_v12;
      _573_v12 = _dafny.Seq.UnicodeFromString("jjbqnymx");
      let _574_v13;
      let _nw106 = new _module.C2();
      _nw106.__ctor(_573_v12, _573_v12);
      _574_v13 = _nw106;
      let _index107 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_572_v11).length));
      (_572_v11)[_index107] = _574_v13;
      let _index108 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_572_v11).length));
      (_572_v11)[_index108] = _574_v13;
      let _575_i5;
      _575_i5 = _dafny.ZERO;
      L4: {
        while (p0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_575_i5)) {
              break L4;
            }
            _575_i5 = (_575_i5).plus(_dafny.ONE);
            let _576_v14;
            _576_v14 = _dafny.Seq.of(p3);
            let _577_v15;
            _577_v15 = _dafny.Seq.of(_576_v14, _576_v14);
            let _578_v16;
            _578_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_577_v15, _577_v15),!(p3));
            r0 = (((_578_v16).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), ((_581_v14) => function (_582_i6) {
              return _581_v14;
            })(_576_v14)))) ? ((_578_v16).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), ((_579_v14) => function (_580_i6) {
              return _579_v14;
            })(_576_v14)))) : (_module.__default.fm33(_573_v12, globalState)));
            let _583_v17;
            let _nw107 = Array((new BigNumber(29)).toNumber()).fill([]);
            _583_v17 = _nw107;
            let _584_v18;
            let _585_v19;
            let _out10;
            let _out11;
            let _outcollector5 = (_574_v13).m0(_583_v17, globalState);
            _out10 = _outcollector5[0];
            _out11 = _outcollector5[1];
            _584_v18 = _out10;
            _585_v19 = _out11;
            let _586_v20;
            _586_v20 = new _dafny.CodePoint('v'.codePointAt(0));
            _586_v20 = _586_v20;
            let _index109 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((p1).length));
            (p1)[_index109] = new BigNumber(25);
            let _587_v21;
            let _nw108 = Array((_dafny.ONE).toNumber()).fill(_module.D4.Default());
            _587_v21 = _nw108;
            let _588_v22;
            _588_v22 = _module.D4.create_DC9(_569_v10);
            let _589_v23;
            _589_v23 = _module.D4.create_DC11(_588_v22);
            let _index110 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_587_v21).length));
            (_587_v21)[_index110] = _module.D4.create_DC11(_588_v22);
            let _590_v24;
            _590_v24 = _dafny.Seq.of(_this.f7);
            let _591_v25;
            _591_v25 = _dafny.Seq.of(new BigNumber((_590_v24).length));
            let _592_v26;
            _592_v26 = _dafny.MultiSet.fromElements(_this.f7, new BigNumber((_590_v24).length));
            let _593_v27;
            _593_v27 = _module.D1.create_DC4(new BigNumber((_dafny.Seq.of(p0, p0, _584_v18)).length), _586_v20, p0, p2);
            let _594_v28;
            let _nw109 = Array((new BigNumber(12)).toNumber());
            _nw109[(_dafny.ZERO).toNumber()] = p3;
            _nw109[(_dafny.ONE).toNumber()] = p3;
            _nw109[(new BigNumber(2)).toNumber()] = true;
            _nw109[(new BigNumber(3)).toNumber()] = _584_v18;
            _nw109[(new BigNumber(4)).toNumber()] = p3;
            _nw109[(new BigNumber(5)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_576_v14, _module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_576_v14).length)), false)).length))).IsProperSubsetOf(_592_v26);
            _nw109[(new BigNumber(6)).toNumber()] = (((_593_v27).dtor_cf8) ? (_584_v18) : (!(p0)));
            _nw109[(new BigNumber(7)).toNumber()] = ((_576_v14)[_module.__default.safeIndex(_this.f7, new BigNumber((_576_v14).length))]) && (p0);
            _nw109[(new BigNumber(8)).toNumber()] = _dafny.Seq.IsPrefixOf(_573_v12, _573_v12);
            _nw109[(new BigNumber(9)).toNumber()] = true;
            _nw109[(new BigNumber(10)).toNumber()] = _584_v18;
            _nw109[(new BigNumber(11)).toNumber()] = !(!(p3)) || (_584_v18);
            _594_v28 = _nw109;
            let _index111 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_594_v28).length));
            (_594_v28)[_index111] = ((!(p3)) ? (p0) : (!(_584_v18)));
            let _595_v29;
            _595_v29 = _dafny.MultiSet.fromElements(_584_v18);
            let _index112 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((p1).length));
            let _index113 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_587_v21).length));
            let _index114 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_594_v28).length));
            let _rhs80 = (((_595_v29).contains(p3)) ? ((_595_v29).get(p3)) : (p2));
            let _rhs81 = _module.D4.create_DC11(_589_v23);
            let _rhs82 = !_dafny.areEqual(_576_v14, _576_v14);
            let _lhs59 = p1;
            let _lhs60 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((p1).length));
            let _lhs61 = _587_v21;
            let _lhs62 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_587_v21).length));
            let _lhs63 = _594_v28;
            let _lhs64 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_594_v28).length));
            _lhs59[_lhs60] = _rhs80;
            _lhs61[_lhs62] = _rhs81;
            _lhs63[_lhs64] = _rhs82;
          }
        }
      }
      let _596_v30;
      _596_v30 = _dafny.MultiSet.fromElements(p3);
      let _597_v31;
      _597_v31 = _dafny.Seq.of(false);
      let _598_v32;
      _598_v32 = _dafny.Seq.of(_596_v30, _596_v30, _dafny.MultiSet.FromArray(_597_v31));
      let _599_v33;
      _599_v33 = _module.D2.create_DC6(_this.f7, p0);
      let _600_v34;
      let _nw110 = Array((new BigNumber(7)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = p1;
      _nw110[(_dafny.ONE).toNumber()] = _569_v10;
      _nw110[(new BigNumber(2)).toNumber()] = p1;
      _nw110[(new BigNumber(3)).toNumber()] = _569_v10;
      _nw110[(new BigNumber(4)).toNumber()] = _569_v10;
      _nw110[(new BigNumber(5)).toNumber()] = _569_v10;
      _nw110[(new BigNumber(6)).toNumber()] = p1;
      _600_v34 = _nw110;
      let _601_v35;
      _601_v35 = _module.D0.create_DC0((_dafny.ZERO).minus((_599_v33).dtor_cf11), _573_v12, _600_v34);
      let _602_v36;
      let _nw111 = new _module.C1();
      _nw111.__ctor((_module.D10.create_DC23((_598_v32)[_module.__default.safeIndex(_this.f7, new BigNumber((_598_v32).length))], p2, true, p0)).dtor_cf35, (_601_v35).dtor_cf0);
      _602_v36 = _nw111;
      let _603_v37;
      _603_v37 = _dafny.Set.fromElements(_602_v36, _602_v36, ((p3) ? (_602_v36) : (_602_v36)));
      _603_v37 = _603_v37;
      r0 = p0;
      return r0;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f6 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (!(false) || (false)) {
        return _dafny.MultiSet.fromElements(true, false, true, !(!(true)), !(false));
      } else {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(true, false, true, true));
      }
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.fromElements(new BigNumber(565), new BigNumber(((_this).f6).length), new BigNumber((_dafny.Seq.of(new BigNumber(825))).length))).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(156), new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(-408), new BigNumber((_dafny.Seq.of(true, true, false, false)).length)))) {
        return new BigNumber((_dafny.Seq.Concat((_this).f6, (_this).f6)).length);
      } else {
        return _module.__default.safeDivisionInt(new BigNumber(-330), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(602)), function (_604_i0) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), (_this).f6)).length));
      }
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _605_v0;
      _605_v0 = new BigNumber(112);
      let _606_v1;
      _606_v1 = false;
      let _607_v2;
      _607_v2 = _dafny.MultiSet.fromElements(_605_v0);
      let _608_v3;
      _608_v3 = _dafny.Seq.of(_605_v0);
      let _609_v4;
      _609_v4 = _dafny.Seq.of(_606_v1);
      let _610_v5;
      let _nw112 = Array((new BigNumber(20)).toNumber());
      _nw112[(_dafny.ZERO).toNumber()] = _605_v0;
      _nw112[(_dafny.ONE).toNumber()] = (_605_v0).minus(_605_v0);
      _nw112[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_605_v0, _605_v0));
      _nw112[(new BigNumber(3)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(4)).toNumber()] = new BigNumber((((_606_v1) ? (_607_v2) : (_dafny.MultiSet.FromArray(_608_v3)))).cardinality());
      _nw112[(new BigNumber(5)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(6)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(7)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(8)).toNumber()] = new BigNumber((_608_v3).length);
      _nw112[(new BigNumber(9)).toNumber()] = new BigNumber(686);
      _nw112[(new BigNumber(10)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(11)).toNumber()] = new BigNumber(967);
      _nw112[(new BigNumber(12)).toNumber()] = (_605_v0).minus(_605_v0);
      _nw112[(new BigNumber(13)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(14)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_606_v1, _606_v1), _609_v4)).length);
      _nw112[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_605_v0, new BigNumber(753));
      _nw112[(new BigNumber(17)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(18)).toNumber()] = _605_v0;
      _nw112[(new BigNumber(19)).toNumber()] = (_this).fm13(_606_v1, _605_v0, _606_v1, globalState);
      _610_v5 = _nw112;
      let _611_v6;
      _611_v6 = _dafny.MultiSet.fromElements((_module.D3.create_DC7(_607_v2)).dtor_cf13, _607_v2, _607_v2);
      let _index115 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
      (_610_v5)[_index115] = (((_611_v6).contains(_607_v2)) ? ((_611_v6).get(_607_v2)) : (_605_v0));
      let _index116 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
      (_610_v5)[_index116] = _605_v0;
      let _612_v7;
      _612_v7 = _dafny.Seq.of(_608_v3);
      let _613_v8;
      _613_v8 = _module.D2.create_DC5((_612_v7)[_module.__default.safeIndex((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], new BigNumber((_612_v7).length))]);
      let _614_v9;
      _614_v9 = _dafny.Map.Empty.slice().updateUnsafe((_613_v8).dtor_cf10,_606_v1);
      _614_v9 = (_614_v9).update(_608_v3, _606_v1);
      if (_606_v1) {
        let _615_v10;
        _615_v10 = _dafny.Seq.of(_607_v2);
        let _index117 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
        let _rhs83 = !(!((((_615_v10)[_module.__default.safeIndex(new BigNumber(((_this).f6).length), new BigNumber((_615_v10).length))]).Difference(_607_v2)).IsDisjointFrom(_607_v2)));
        let _rhs84 = _605_v0;
        let _lhs65 = _610_v5;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
        r0 = _rhs83;
        _lhs65[_lhs66] = _rhs84;
        let _616_v11;
        _616_v11 = _module.D4.create_DC9(_610_v5);
        let _617_v12;
        _617_v12 = _dafny.Set.fromElements((_616_v11).dtor_cf15, _610_v5, _610_v5, _610_v5);
        if (((_608_v3)[_module.__default.safeIndex(new BigNumber((_617_v12).length), new BigNumber((_608_v3).length))]).isLessThanOrEqualTo(_605_v0)) {
          let _618_v13;
          _618_v13 = new _dafny.CodePoint('a'.codePointAt(0));
          let _619_v14;
          _619_v14 = _dafny.Map.Empty.slice().updateUnsafe(_605_v0,_618_v13);
          let _620_v15;
          let _nw113 = new _module.C2();
          _nw113.__ctor(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_621_i0) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), (_this).f6), _module.__default.safeIndex(new BigNumber((_609_v4).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_622_i0) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), (_this).f6)).length)), _618_v13), _dafny.Seq.update((_this).f6, _module.__default.safeIndex(new BigNumber(((_this).f6).length), new BigNumber(((_this).f6).length)), (((_619_v14).contains((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))])) ? ((_619_v14).get((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))])) : (_618_v13))));
          _620_v15 = _nw113;
          r0 = _606_v1;
          let _623_v16;
          let _init11 = ((_624_v1) => function (_625_i1) {
            return _624_v1;
          })(_606_v1);
          let _nw114 = Array((new BigNumber(21)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw114.length); _i0_11++) {
            _nw114[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _623_v16 = _nw114;
          _623_v16 = _623_v16;
          let _626_v17;
          _626_v17 = _dafny.Seq.UnicodeFromString("nrp");
          let _627_v18;
          let _nw115 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _627_v18 = _nw115;
          let _628_v19;
          _628_v19 = _dafny.Set.fromElements(_605_v0);
          let _629_v20;
          _629_v20 = _dafny.Map.Empty.slice().updateUnsafe(_605_v0,_606_v1);
          let _630_v21;
          _630_v21 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("oexlyc"), (_620_v15).f8);
          let _index118 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_627_v18).length));
          (_627_v18)[_index118] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_628_v19).length), (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]), _dafny.Seq.of(_module.__default.fm19(_606_v1, new BigNumber((_629_v20).length), globalState), new BigNumber(129), new BigNumber(((_630_v21)[_module.__default.safeIndex(new BigNumber((_607_v2).cardinality()), new BigNumber((_630_v21).length))]).length)));
          let _index119 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_627_v18).length));
          let _index120 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          let _rhs85 = _618_v13;
          let _rhs86 = (_620_v15).fm18(_609_v4, globalState);
          let _rhs87 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(981)), ((_631_v5, _632_v13, _633_v1) => function (_634_i2) {
            return (_module.D1.create_DC4((_631_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_631_v5).length))], _632_v13, _633_v1, (_631_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_631_v5).length))])).dtor_cf7;
          })(_610_v5, _618_v13, _606_v1));
          let _rhs88 = _dafny.Seq.Concat(_608_v3, _608_v3);
          let _rhs89 = (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))];
          let _lhs67 = _627_v18;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_627_v18).length));
          let _lhs69 = _610_v5;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          _618_v13 = _rhs85;
          _606_v1 = _rhs86;
          _626_v17 = _rhs87;
          _lhs67[_lhs68] = _rhs88;
          _lhs69[_lhs70] = _rhs89;
          let _635_v22;
          let _nw116 = new _module.C0();
          _nw116.__ctor();
          _635_v22 = _nw116;
        } else {
          let _636_v23;
          let _nw117 = Array((new BigNumber(7)).toNumber()).fill([]);
          _636_v23 = _nw117;
          let _index121 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_636_v23).length));
          (_636_v23)[_index121] = _610_v5;
          let _index122 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_636_v23).length));
          (_636_v23)[_index122] = _610_v5;
          let _637_v24;
          _637_v24 = _dafny.Seq.of((_this).f6, (_this).f6);
          let _638_v25;
          _638_v25 = new _dafny.CodePoint('c'.codePointAt(0));
          let _639_v26;
          _639_v26 = _dafny.MultiSet.fromElements(_638_v25);
          let _rhs90 = _608_v3;
          let _rhs91 = ((new BigNumber((_637_v24).length)).minus(new BigNumber(((_639_v26).update(_638_v25, _module.__default.abs((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]))).cardinality()))).plus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
          _608_v3 = _rhs90;
          _605_v0 = _rhs91;
          let _640_v27;
          _640_v27 = _dafny.Map.Empty.slice().updateUnsafe(true,_605_v0);
          _640_v27 = (_640_v27).update(_606_v1, _605_v0);
          let _641_v28;
          _641_v28 = _module.D4.create_DC9((_636_v23)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_636_v23).length))]);
          let _642_v29;
          _642_v29 = _module.D4.create_DC11(_641_v28);
          let _643_v30;
          _643_v30 = _module.D4.create_DC11(_641_v28);
          let _644_v31;
          _644_v31 = _dafny.Seq.of(_643_v30, _643_v30, _643_v30, _643_v30, _643_v30);
          let _645_v32;
          _645_v32 = _dafny.Map.Empty.slice().updateUnsafe(true,_644_v31);
          let _646_v33;
          _646_v33 = _dafny.Map.Empty.slice().updateUnsafe((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))],new BigNumber((_645_v32).length));
          let _647_v36;
          let _nw118 = Array((new BigNumber(13)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _module.__default.fm29(globalState);
          _nw118[(_dafny.ONE).toNumber()] = (_646_v33).update((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], _605_v0);
          _nw118[(new BigNumber(2)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(3)).toNumber()] = function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(744), new BigNumber(9))) {
              let _648_v34 = _compr_18;
              if (((new BigNumber(744)).isLessThanOrEqualTo(_648_v34)) && ((_648_v34).isLessThan(new BigNumber(9)))) {
                _coll18.push([_module.__default.safeModuloInt(_648_v34, _605_v0),(_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]]);
              }
            }
            return _coll18;
          }();
          _nw118[(new BigNumber(4)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(5)).toNumber()] = function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(446), new BigNumber(923))) {
              let _649_v35 = _compr_19;
              if (((new BigNumber(446)).isLessThanOrEqualTo(_649_v35)) && ((_649_v35).isLessThan(new BigNumber(923)))) {
                _coll19.push([_module.__default.safeModuloInt(_649_v35, (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]),_605_v0]);
              }
            }
            return _coll19;
          }();
          _nw118[(new BigNumber(6)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(7)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(8)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))],(_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))])).update(_605_v0, (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
          _nw118[(new BigNumber(9)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(10)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(11)).toNumber()] = _646_v33;
          _nw118[(new BigNumber(12)).toNumber()] = (_646_v33).update((((_640_v27).contains(_606_v1)) ? ((_640_v27).get(_606_v1)) : (new BigNumber((_dafny.MultiSet.fromElements(_606_v1, _606_v1, _606_v1)).cardinality()))), (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
          _647_v36 = _nw118;
          let _650_v37;
          _650_v37 = _dafny.Seq.of(_647_v36);
          _647_v36 = (_650_v37)[_module.__default.safeIndex((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], new BigNumber((_650_v37).length))];
          _605_v0 = (_module.__default.fm19(_606_v1, new BigNumber(((_this).f6).length), globalState)).multipliedBy(new BigNumber(152));
        }
        _616_v11 = _616_v11;
        let _651_v38;
        _651_v38 = _module.D9.create_DC20(_606_v1, _606_v1, _606_v1);
        let _source8 = _651_v38;
        if (_source8.is_DC19) {
          let _652___mcc_h0 = (_source8).cf29;
          let _653___mcc_h1 = (_source8).cf30;
          let _654_cf30 = _653___mcc_h1;
          let _655_cf29 = _652___mcc_h0;
          let _656_v39;
          let _nw119 = new _module.C0();
          _nw119.__ctor();
          _656_v39 = _nw119;
          let _657_v40;
          let _nw120 = Array((new BigNumber(10)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _656_v39;
          _nw120[(_dafny.ONE).toNumber()] = _656_v39;
          _nw120[(new BigNumber(2)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(3)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(4)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(5)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(6)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(7)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(8)).toNumber()] = _656_v39;
          _nw120[(new BigNumber(9)).toNumber()] = _656_v39;
          _657_v40 = _nw120;
          let _658_v41;
          _658_v41 = _dafny.Seq.of(_657_v40, _657_v40, _657_v40, _657_v40);
          let _659_v42;
          let _nw121 = Array((new BigNumber(20)).toNumber());
          _nw121[(_dafny.ZERO).toNumber()] = _657_v40;
          _nw121[(_dafny.ONE).toNumber()] = _657_v40;
          _nw121[(new BigNumber(2)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(3)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(4)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(5)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(6)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(7)).toNumber()] = ((_606_v1) ? (_657_v40) : ((_658_v41)[_module.__default.safeIndex((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], new BigNumber((_658_v41).length))]));
          _nw121[(new BigNumber(8)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(9)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(10)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(11)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(12)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(13)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(14)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(15)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(16)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(17)).toNumber()] = (_658_v41)[_module.__default.safeIndex(_655_cf29, new BigNumber((_658_v41).length))];
          _nw121[(new BigNumber(18)).toNumber()] = _657_v40;
          _nw121[(new BigNumber(19)).toNumber()] = _657_v40;
          _659_v42 = _nw121;
          let _660_v43;
          _660_v43 = new _dafny.CodePoint('r'.codePointAt(0));
          let _661_v44;
          _661_v44 = _dafny.Map.Empty.slice().updateUnsafe(_605_v0,_659_v42);
          let _662_v45;
          _662_v45 = _dafny.Map.Empty.slice().updateUnsafe(_606_v1,_659_v42);
          let _rhs92 = ((_606_v1) ? ((((_661_v44).contains(_655_cf29)) ? ((_661_v44).get(_655_cf29)) : (_659_v42))) : (((_606_v1) ? ((((_662_v45).contains(_606_v1)) ? ((_662_v45).get(_606_v1)) : (_659_v42))) : (_659_v42))));
          let _rhs93 = _module.__default.fm30(_dafny.Seq.Concat((_this).f6, (_this).f6), new BigNumber((_608_v3).length), (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], globalState);
          _659_v42 = _rhs92;
          _660_v43 = _rhs93;
          let _663_v46;
          let _nw122 = new _module.C2();
          _nw122.__ctor(_dafny.Seq.Concat((_this).f6, (_this).f6), _dafny.Seq.UnicodeFromString("bjvex"));
          _663_v46 = _nw122;
          _663_v46 = _663_v46;
          let _664_v47;
          _664_v47 = _dafny.MultiSet.fromElements(_606_v1, _606_v1);
          let _665_v48;
          _665_v48 = _dafny.Set.fromElements((((_664_v47).contains(!(_606_v1))) ? ((_664_v47).get(!(_606_v1))) : (_655_cf29)));
          _606_v1 = (_665_v48).IsDisjointFrom(_665_v48);
          let _666_v49;
          _666_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_663_v46).f9, _module.__default.safeIndex((_this).fm13(_606_v1, _655_cf29, _606_v1, globalState), new BigNumber(((_663_v46).f9).length)), new _dafny.CodePoint('u'.codePointAt(0))),new BigNumber(618));
          _666_v49 = (_666_v49).update((_this).f6, _654_cf30);
        } else if (_source8.is_DC20) {
          let _667___mcc_h2 = (_source8).cf31;
          let _668___mcc_h3 = (_source8).cf32;
          let _669___mcc_h4 = (_source8).cf33;
          let _670_cf33 = _669___mcc_h4;
          let _671_cf32 = _668___mcc_h3;
          let _672_cf31 = _667___mcc_h2;
          let _index123 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          (_610_v5)[_index123] = _605_v0;
          let _index124 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          (_610_v5)[_index124] = _605_v0;
          let _673_v50;
          _673_v50 = _dafny.MultiSet.fromElements(_606_v1, _672_cf31);
          let _674_v51;
          let _nw123 = new _module.C4();
          _nw123.__ctor(_605_v0);
          _674_v51 = _nw123;
          let _675_v52;
          let _676_v53;
          let _out12;
          let _out13;
          let _outcollector6 = (_this).m3(!(((_dafny.MultiSet.fromElements(_670_cf33, _672_cf31)).update(_671_cf32, _module.__default.abs(new BigNumber(16)))).IsSubsetOf(_673_v50)), _674_v51, globalState);
          _out12 = _outcollector6[0];
          _out13 = _outcollector6[1];
          _675_v52 = _out12;
          _676_v53 = _out13;
          let _index125 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          (_610_v5)[_index125] = _675_v52;
        } else {
          let _677___mcc_h5 = (_source8).cf28;
          let _678_cf28 = _677___mcc_h5;
          let _679_v54;
          let _nw124 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _679_v54 = _nw124;
          let _680_v55;
          _680_v55 = new _dafny.CodePoint('x'.codePointAt(0));
          let _index126 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_679_v54).length));
          (_679_v54)[_index126] = _680_v55;
          let _index127 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_679_v54).length));
          (_679_v54)[_index127] = _680_v55;
          _605_v0 = _module.__default.fm14(new BigNumber((_dafny.Set.fromElements(_606_v1, _606_v1)).length), (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], globalState);
          let _681_v56;
          _681_v56 = _dafny.Set.fromElements(!(_dafny.Seq.contains((_this).f6, _680_v55)));
          _681_v56 = _681_v56;
          _605_v0 = (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))];
        }
        r0 = _606_v1;
      } else {
        let _682_v58;
        let _init12 = ((_683_v5, _684_v1) => function (_685_i3) {
          return function () {
            let _coll20 = new _dafny.Set();
            for (const _compr_20 of (_dafny.Map.Empty.slice().updateUnsafe((_683_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_683_v5).length))],_684_v1)).Keys.Elements) {
              let _686_v57 = _compr_20;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_683_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_683_v5).length))],_684_v1)).contains(_686_v57)) {
                _coll20.add((_686_v57).plus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(false))).cardinality())));
              }
            }
            return _coll20;
          }();
        })(_610_v5, _606_v1);
        let _nw125 = Array((new BigNumber(5)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw125.length); _i0_12++) {
          _nw125[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _682_v58 = _nw125;
        let _687_v59;
        _687_v59 = _dafny.Set.fromElements((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], _605_v0);
        let _index128 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_682_v58).length));
        (_682_v58)[_index128] = _687_v59;
        let _index129 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_682_v58).length));
        (_682_v58)[_index129] = _687_v59;
        _606_v1 = !(_606_v1);
        let _688_v60;
        _688_v60 = _module.D4.create_DC10(_606_v1);
        let _689_v61;
        let _nw126 = Array((new BigNumber(13)).toNumber());
        _nw126[(_dafny.ZERO).toNumber()] = _606_v1;
        _nw126[(_dafny.ONE).toNumber()] = _606_v1;
        _nw126[(new BigNumber(2)).toNumber()] = _606_v1;
        _nw126[(new BigNumber(3)).toNumber()] = false;
        _nw126[(new BigNumber(4)).toNumber()] = !(_606_v1);
        _nw126[(new BigNumber(5)).toNumber()] = _606_v1;
        _nw126[(new BigNumber(6)).toNumber()] = (_688_v60).dtor_cf16;
        _nw126[(new BigNumber(7)).toNumber()] = !(_606_v1);
        _nw126[(new BigNumber(8)).toNumber()] = _606_v1;
        _nw126[(new BigNumber(9)).toNumber()] = _606_v1;
        _nw126[(new BigNumber(10)).toNumber()] = _606_v1;
        _nw126[(new BigNumber(11)).toNumber()] = _606_v1;
        _nw126[(new BigNumber(12)).toNumber()] = _606_v1;
        _689_v61 = _nw126;
        let _690_v62;
        _690_v62 = _689_v61;
        _689_v61 = (_690_v62);
        _605_v0 = _module.__default.safeDivisionInt(_605_v0, (_dafny.ZERO).minus((_dafny.ZERO).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))])));
        let _index130 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
        (_610_v5)[_index130] = (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))];
      }
      let _hi5 = _605_v0;
      for (let _691_i4 = _module.__default.safeDivisionInt((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rgtil")).length))); _691_i4.isLessThan(_hi5); _691_i4 = _691_i4.plus(_dafny.ONE)) {
        _608_v3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_608_v3, _608_v3), _608_v3);
        r0 = _606_v1;
        let _index131 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
        (_610_v5)[_index131] = (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))];
        let _index132 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
        (_610_v5)[_index132] = _module.__default.safeDivisionInt(_691_i4, (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
      }
      let _692_v63;
      _692_v63 = _dafny.Set.fromElements((_this).f6, (_this).f6, (_this).f6);
      let _693_i5;
      _693_i5 = _dafny.ZERO;
      L5: {
        while ((_dafny.Set.fromElements((_this).f6)).IsDisjointFrom(_692_v63)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_693_i5)) {
              break L5;
            }
            _693_i5 = (_693_i5).plus(_dafny.ONE);
            let _694_v64;
            let _nw127 = new _module.C3();
            _nw127.__ctor();
            _694_v64 = _nw127;
            let _695_v65;
            _695_v65 = _module.D11.create_DC24(_694_v64);
            let _696_v66;
            _696_v66 = _dafny.Seq.of(_694_v64, _694_v64, _694_v64);
            let _697_v67;
            let _nw128 = Array((new BigNumber(11)).toNumber());
            _nw128[(_dafny.ZERO).toNumber()] = _694_v64;
            _nw128[(_dafny.ONE).toNumber()] = _694_v64;
            _nw128[(new BigNumber(2)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(3)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(4)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(5)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(6)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(7)).toNumber()] = (_695_v65).dtor_cf39;
            _nw128[(new BigNumber(8)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(9)).toNumber()] = _694_v64;
            _nw128[(new BigNumber(10)).toNumber()] = (_696_v66)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_696_v66).length))];
            _697_v67 = _nw128;
            let _index133 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_697_v67).length));
            (_697_v67)[_index133] = _694_v64;
            let _index134 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_697_v67).length));
            let _nw129 = new _module.C3();
            _nw129.__ctor();
            (_697_v67)[_index134] = _nw129;
            let _698_v68;
            _698_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]),_dafny.Seq.UnicodeFromString("mvbygep"));
            let _699_v69;
            let _nw130 = Array((new BigNumber(3)).toNumber());
            _nw130[(_dafny.ZERO).toNumber()] = _606_v1;
            _nw130[(_dafny.ONE).toNumber()] = _606_v1;
            _nw130[(new BigNumber(2)).toNumber()] = _606_v1;
            _699_v69 = _nw130;
            let _700_v70;
            _700_v70 = _699_v69;
            let _701_v71;
            _701_v71 = _dafny.Set.fromElements((_700_v70), _699_v69);
            _698_v68 = (_698_v68).update(_module.__default.safeModuloInt(_605_v0, new BigNumber((_701_v71).length)), _dafny.Seq.Concat((_this).f6, (_this).f6));
            let _702_v72;
            let _nw131 = Array((new BigNumber(27)).toNumber()).fill([]);
            _702_v72 = _nw131;
            _702_v72 = p0;
            let _703_v73;
            _703_v73 = _dafny.Set.fromElements(false, true);
            let _704_v74;
            _704_v74 = _dafny.Map.Empty.slice().updateUnsafe((_605_v0).minus(_605_v0),_703_v73);
            let _705_v75;
            _705_v75 = _dafny.Map.Empty.slice().updateUnsafe(_607_v2,((_dafny.ZERO).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))])).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]));
            let _index135 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
            let _rhs94 = (((_705_v75).contains((_dafny.MultiSet.fromElements((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], _605_v0)).Intersect(_module.__default.fm31((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], globalState)))) ? ((_705_v75).get((_dafny.MultiSet.fromElements((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], _605_v0)).Intersect(_module.__default.fm31((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], globalState)))) : ((_dafny.ZERO).minus((_this).fm13(!(_606_v1), _module.__default.fm24(_605_v0, false, _606_v1, new BigNumber(((_this).f6).length), globalState), _606_v1, globalState))));
            let _rhs95 = _704_v74;
            let _lhs71 = _610_v5;
            let _lhs72 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
            _lhs71[_lhs72] = _rhs94;
            _704_v74 = _rhs95;
          }
        }
      }
      if (_606_v1) {
        let _706_v76;
        _706_v76 = _dafny.MultiSet.fromElements(_606_v1);
        let _707_v77;
        _707_v77 = _module.D10.create_DC21(_706_v76);
        let _708_v78;
        _708_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_609_v4).length),(_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
        let _709_v80;
        _709_v80 = _dafny.Map.Empty.slice().updateUnsafe(_707_v77,new BigNumber((function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of (_708_v78).Keys.Elements) {
            let _710_v79 = _compr_21;
            if ((_708_v78).contains(_710_v79)) {
              _coll21.add(_module.__default.safeModuloInt(_710_v79, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D9.create_DC19(new BigNumber(-120), new BigNumber(390))).dtor_cf29,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(768))).length))).length)));
            }
          }
          return _coll21;
        }()).length));
        _709_v80 = (_709_v80).update(_707_v77, new BigNumber(283));
        let _711_v81;
        let _nw132 = new _module.C1();
        _nw132.__ctor((_706_v76).Intersect(_706_v76), (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
        _711_v81 = _nw132;
        if (_606_v1) {
          let _index136 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          (_610_v5)[_index136] = _605_v0;
          let _index137 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          (_610_v5)[_index137] = (_dafny.ZERO).minus(new BigNumber(-96));
          let _712_v82;
          _712_v82 = new _dafny.CodePoint('v'.codePointAt(0));
          _712_v82 = ((_this).f6)[_module.__default.safeIndex((_dafny.ZERO).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]), new BigNumber(((_this).f6).length))];
          _606_v1 = !(_module.__default.fm33((_this).f6, globalState));
          let _713_v83;
          let _nw133 = Array((new BigNumber(22)).toNumber()).fill(false);
          _713_v83 = _nw133;
          _713_v83 = _713_v83;
        } else {
          let _index138 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
          (_610_v5)[_index138] = _711_v81.f11;
          _609_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_609_v4, _dafny.Seq.of(_606_v1)), _609_v4);
          let _714_v84;
          let _nw134 = new _module.C0();
          _nw134.__ctor();
          _714_v84 = _nw134;
          let _nw135 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _610_v5 = _nw135;
          _605_v0 = (_dafny.ZERO).minus(_711_v81.f11);
        }
        let _715_v85;
        let _nw136 = new _module.C0();
        _nw136.__ctor();
        _715_v85 = _nw136;
        let _716_v86;
        _716_v86 = _module.D7.create_DC15(_606_v1, _711_v81.f11, _605_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))],_711_v81.f11)).length), _715_v85);
        let _717_v87;
        _717_v87 = _dafny.Map.Empty.slice().updateUnsafe((_716_v86).dtor_cf24,_606_v1);
        let _718_v88;
        _718_v88 = _dafny.Map.Empty.slice().updateUnsafe(_717_v87,_605_v0);
        let _719_v90;
        _719_v90 = new _dafny.CodePoint('m'.codePointAt(0));
        let _720_v91;
        _720_v91 = _dafny.Map.Empty.slice().updateUnsafe(_711_v81.f11,_dafny.Map.Empty.slice().updateUnsafe(_719_v90,_606_v1));
        _718_v88 = (_718_v88).update(function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of (_720_v91).Keys.Elements) {
            let _721_v89 = _compr_22;
            if ((_720_v91).contains(_721_v89)) {
              _coll22.push([(_721_v89).minus(_711_v81.f11),!(false)]);
            }
          }
          return _coll22;
        }(), (_this).fm13(_606_v1, (((_708_v78).contains(new BigNumber(-45))) ? ((_708_v78).get(new BigNumber(-45))) : (_605_v0)), _606_v1, globalState));
        let _722_v92;
        let _nw137 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _722_v92 = _nw137;
        _722_v92 = _722_v92;
      } else {
        let _723_v94;
        let _init13 = ((_724_v5, _725_v2) => function (_726_i6) {
          return (_dafny.Set.fromElements((_724_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_724_v5).length))], (_724_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_724_v5).length))])).Union(function () {
            let _coll23 = new _dafny.Set();
            for (const _compr_23 of (_725_v2).Elements) {
              let _727_v93 = _compr_23;
              if ((_725_v2).contains(_727_v93)) {
                _coll23.add((_727_v93).multipliedBy(new BigNumber(647)));
              }
            }
            return _coll23;
          }());
        })(_610_v5, _607_v2);
        let _nw138 = Array((new BigNumber(25)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw138.length); _i0_13++) {
          _nw138[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _723_v94 = _nw138;
        let _728_v95;
        _728_v95 = _dafny.Set.fromElements((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
        let _index139 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_723_v94).length));
        (_723_v94)[_index139] = (_728_v95).Difference(function () {
          let _coll24 = new _dafny.Set();
          for (const _compr_24 of _dafny.IntegerRange(new BigNumber(734), new BigNumber(782))) {
            let _729_v96 = _compr_24;
            if (((new BigNumber(734)).isLessThanOrEqualTo(_729_v96)) && ((_729_v96).isLessThan(new BigNumber(782)))) {
              _coll24.add((_729_v96).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]));
            }
          }
          return _coll24;
        }());
        let _index140 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_723_v94).length));
        (_723_v94)[_index140] = (_728_v95).Difference(function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), ((_730_v0) => function (_731_i7) {
            return _730_v0;
          })(_605_v0))).Elements) {
            let _732_v97 = _compr_25;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), ((_733_v0) => function (_731_i7) {
              return _733_v0;
            })(_605_v0)), _732_v97)) {
              _coll25.add((_732_v97).multipliedBy(new BigNumber(907)));
            }
          }
          return _coll25;
        }());
        let _734_v98;
        _734_v98 = _dafny.Set.fromElements(_606_v1, _606_v1);
        let _735_v99;
        _735_v99 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_734_v98).length)),(_dafny.ZERO).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))])), _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm24(_605_v0, _606_v1, _606_v1, _605_v0, globalState),_605_v0));
        let _736_v100;
        let _nw139 = Array((new BigNumber(25)).toNumber());
        _nw139[(_dafny.ZERO).toNumber()] = true;
        _nw139[(_dafny.ONE).toNumber()] = _606_v1;
        _nw139[(new BigNumber(2)).toNumber()] = true;
        _nw139[(new BigNumber(3)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(4)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(5)).toNumber()] = (_606_v1) === (!(_606_v1));
        _nw139[(new BigNumber(6)).toNumber()] = (_735_v99).IsSubsetOf(_735_v99);
        _nw139[(new BigNumber(7)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(8)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(9)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(10)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(11)).toNumber()] = ((_dafny.MultiSet.FromArray(_609_v4)).update(_606_v1, _module.__default.abs((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]))).contains(_606_v1);
        _nw139[(new BigNumber(12)).toNumber()] = !((_605_v0).isLessThanOrEqualTo(new BigNumber(317)));
        _nw139[(new BigNumber(13)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(14)).toNumber()] = !(_606_v1) || (_606_v1);
        _nw139[(new BigNumber(15)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(16)).toNumber()] = !(((_606_v1) ? (_606_v1) : (_606_v1)));
        _nw139[(new BigNumber(17)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(18)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(19)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(20)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(21)).toNumber()] = _606_v1;
        _nw139[(new BigNumber(22)).toNumber()] = _module.__default.fm33(_dafny.Seq.UnicodeFromString("lemmjobsq"), globalState);
        _nw139[(new BigNumber(23)).toNumber()] = !(_606_v1);
        _nw139[(new BigNumber(24)).toNumber()] = _606_v1;
        _736_v100 = _nw139;
        let _index141 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_736_v100).length));
        (_736_v100)[_index141] = (_606_v1) === (_606_v1);
        let _index142 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_736_v100).length));
        let _rhs96 = _609_v4;
        let _rhs97 = _610_v5;
        let _rhs98 = _606_v1;
        let _rhs99 = (_605_v0).isLessThan((_this).fm13(_606_v1, (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))], true, globalState));
        let _lhs73 = _736_v100;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_736_v100).length));
        _609_v4 = _rhs96;
        _610_v5 = _rhs97;
        _lhs73[_lhs74] = _rhs98;
        r0 = _rhs99;
        let _rhs100 = _610_v5;
        let _rhs101 = _608_v3;
        _610_v5 = _rhs100;
        _608_v3 = _rhs101;
        let _737_v101;
        let _nw140 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _737_v101 = _nw140;
        let _738_v102;
        _738_v102 = _dafny.Map.Empty.slice().updateUnsafe(_737_v101,_module.__default.safeModuloInt((_dafny.ZERO).minus((_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]), (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]));
        _738_v102 = (_738_v102).update(_737_v101, (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))]);
        let _index143 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length));
        (_610_v5)[_index143] = (_610_v5)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_610_v5).length))];
      }
      r0 = _606_v1;
      let _739_v103;
      _739_v103 = _dafny.MultiSet.fromElements(_606_v1);
      let _740_v104;
      _740_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), function (_741_i8) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length),_739_v103);
      let _742_v105;
      _742_v105 = _dafny.Seq.of(((_606_v1) ? (_740_v104) : (_740_v104)), _740_v104, _740_v104, _740_v104);
      r1 = (_742_v105)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(((_this).f6).length), _605_v0)), new BigNumber((_742_v105).length))];
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let _743_v0;
      _743_v0 = new BigNumber(-822);
      let _744_v1;
      _744_v1 = _dafny.Seq.of(_743_v0, _743_v0);
      let _745_v2;
      _745_v2 = _module.D1.create_DC4((_744_v1)[_module.__default.safeIndex(_743_v0, new BigNumber((_744_v1).length))], new _dafny.CodePoint('b'.codePointAt(0)), false, new BigNumber(((_this).f6).length));
      let _746_i0;
      _746_i0 = _dafny.ZERO;
      L6: {
        let _pat_let_tv5 = p0;
        let _pat_let_tv6 = p0;
        while (function (_source9) {
          if (_source9.is_DC3) {
            let _757___mcc_h0 = (_source9).cf5;
            let _758_cf5 = _757___mcc_h0;
            return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_dafny.Seq.of(false)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), ((_759_p0) => function (_760_i1) {
              return _dafny.Seq.of(_759_p0, false);
            })(_pat_let_tv5))));
          } else if (_source9.is_DC4) {
            let _761___mcc_h1 = (_source9).cf6;
            let _762___mcc_h2 = (_source9).cf7;
            let _763___mcc_h3 = (_source9).cf8;
            let _764___mcc_h4 = (_source9).cf9;
            let _765_cf9 = _764___mcc_h4;
            let _766_cf8 = _763___mcc_h3;
            let _767_cf7 = _762___mcc_h2;
            let _768_cf6 = _761___mcc_h1;
            return _pat_let_tv6;
          } else {
            let _769___mcc_h5 = (_source9).cf4;
            let _770_cf4 = _769___mcc_h5;
            return !(!(_770_cf4));
          }
        }(_745_v2)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_746_i0)) {
              break L6;
            }
            _746_i0 = (_746_i0).plus(_dafny.ONE);
            let _747_v3;
            _747_v3 = false;
            _747_v3 = p0;
            let _748_v4;
            let _nw141 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
            _748_v4 = _nw141;
            let _749_v5;
            _749_v5 = _module.D2.create_DC6(_743_v0, p0);
            let _index144 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            (_748_v4)[_index144] = (_749_v5).dtor_cf11;
            let _750_v7;
            _750_v7 = _dafny.Map.Empty.slice().updateUnsafe(_747_v3,_747_v3);
            let _751_v8;
            _751_v8 = _dafny.Set.fromElements(new BigNumber((_750_v7).length), (_dafny.ZERO).minus(_743_v0));
            let _752_v9;
            _752_v9 = _dafny.Seq.of(_747_v3, (_751_v8).IsProperSubsetOf(function () {
              let _coll26 = new _dafny.Set();
              for (const _compr_26 of _dafny.IntegerRange(new BigNumber(553), new BigNumber(891))) {
                let _753_v6 = _compr_26;
                if (((new BigNumber(553)).isLessThanOrEqualTo(_753_v6)) && ((_753_v6).isLessThan(new BigNumber(891)))) {
                  _coll26.add(_module.__default.safeModuloInt(_753_v6, new BigNumber((_dafny.Seq.UnicodeFromString("ba")).length)));
                }
              }
              return _coll26;
            }()), p0, !(_747_v3));
            let _754_v10;
            _754_v10 = _dafny.Seq.of(_752_v9);
            let _index145 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            let _rhs102 = (_743_v0).plus(_743_v0);
            let _rhs103 = _module.__default.safeDivisionInt(_743_v0, ((_747_v3) ? (new BigNumber(355)) : (_743_v0)));
            let _rhs104 = _743_v0;
            let _rhs105 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm26(_743_v0, globalState), _752_v9), (_754_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("aaqjua")).length), new BigNumber((_754_v10).length))]);
            let _lhs75 = _748_v4;
            let _lhs76 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            r0 = _rhs102;
            r0 = _rhs103;
            _lhs75[_lhs76] = _rhs104;
            _752_v9 = _rhs105;
            let _755_v11;
            let _nw142 = new _module.C3();
            _nw142.__ctor();
            _755_v11 = _nw142;
            let _756_v12;
            _756_v12 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_748_v4)[_module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length))], (_748_v4)[_module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length))]),(_this).f6);
            let _index146 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            let _index147 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            let _rhs106 = new BigNumber((_756_v12).length);
            let _rhs107 = new BigNumber(-470);
            let _rhs108 = _747_v3;
            let _lhs77 = _748_v4;
            let _lhs78 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            let _lhs79 = _748_v4;
            let _lhs80 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_748_v4).length));
            _lhs77[_lhs78] = _rhs106;
            _lhs79[_lhs80] = _rhs107;
            _747_v3 = _rhs108;
          }
        }
      }
      let _771_v13;
      let _init14 = ((_772_v0) => function (_773_i2) {
        return (_773_i2).plus(_772_v0);
      })(_743_v0);
      let _nw143 = Array((new BigNumber(24)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw143.length); _i0_14++) {
        _nw143[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _771_v13 = _nw143;
      _771_v13 = _771_v13;
      let _774_v14;
      let _nw144 = Array((new BigNumber(13)).toNumber()).fill(false);
      _774_v14 = _nw144;
      let _775_v15;
      _775_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f6).length),_774_v14);
      if ((_775_v15).equals(_775_v15)) {
        let _index148 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_774_v14).length));
        (_774_v14)[_index148] = p0;
        let _776_v16;
        let _nw145 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _776_v16 = _nw145;
        let _index149 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_776_v16).length));
        (_776_v16)[_index149] = _dafny.Seq.Concat((_this).f6, (_this).f6);
        let _index150 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_774_v14).length));
        let _index151 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_776_v16).length));
        let _rhs109 = p0;
        let _rhs110 = new BigNumber((_dafny.Seq.Concat((_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), function (_777_i3) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }))).length);
        let _rhs111 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f6, (_this).f6), _dafny.Seq.UnicodeFromString("npyc"));
        let _lhs81 = _774_v14;
        let _lhs82 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_774_v14).length));
        let _lhs83 = _776_v16;
        let _lhs84 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_776_v16).length));
        _lhs81[_lhs82] = _rhs109;
        r0 = _rhs110;
        _lhs83[_lhs84] = _rhs111;
        _743_v0 = (new BigNumber(-516)).plus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_743_v0,p0)).length)).minus((_744_v1)[_module.__default.safeIndex(_743_v0, new BigNumber((_744_v1).length))]));
        let _index152 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_776_v16).length));
        (_776_v16)[_index152] = (_this).f6;
        let _index153 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_774_v14).length));
        (_774_v14)[_index153] = (_774_v14)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_774_v14).length))];
        let _index154 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_774_v14).length));
        (_774_v14)[_index154] = p0;
      } else {
        r0 = new BigNumber(-976);
        let _778_v17;
        _778_v17 = false;
        let _779_v18;
        _779_v18 = _dafny.Seq.of(p0, _778_v17);
        let _780_v19;
        _780_v19 = _module.D1.create_DC2((_779_v18)[_module.__default.safeIndex(_743_v0, new BigNumber((_779_v18).length))]);
        let _781_v20;
        _781_v20 = new _dafny.CodePoint('g'.codePointAt(0));
        let _pat_let_tv7 = _743_v0;
        let _pat_let_tv8 = _781_v20;
        let _pat_let_tv9 = _778_v17;
        let _rhs112 = p0;
        let _rhs113 = !(!(false) || (false));
        let _rhs114 = (_780_v19).dtor_cf4;
        let _rhs115 = (function (_pat_let8_0) {
          return function (_785_dt__update__tmp_h1) {
            return function (_pat_let12_0) {
              return function (_786_dt__update_hcf8_h0) {
                return _module.D1.create_DC4((_785_dt__update__tmp_h1).dtor_cf6, (_785_dt__update__tmp_h1).dtor_cf7, _786_dt__update_hcf8_h0, (_785_dt__update__tmp_h1).dtor_cf9);
              }(_pat_let12_0);
            }(_pat_let_tv9);
          }(_pat_let8_0);
        }(function (_pat_let9_0) {
          return function (_782_dt__update__tmp_h0) {
            return function (_pat_let10_0) {
              return function (_783_dt__update_hcf9_h0) {
                return function (_pat_let11_0) {
                  return function (_784_dt__update_hcf7_h0) {
                    return _module.D1.create_DC4((_782_dt__update__tmp_h0).dtor_cf6, _784_dt__update_hcf7_h0, (_782_dt__update__tmp_h0).dtor_cf8, _783_dt__update_hcf9_h0);
                  }(_pat_let11_0);
                }(_pat_let_tv8);
              }(_pat_let10_0);
            }(_pat_let_tv7);
          }(_pat_let9_0);
        }(_745_v2))).dtor_cf7;
        let _rhs116 = _module.__default.fm33(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xsfeu"), _dafny.Seq.UnicodeFromString("mosbxp")), globalState);
        _778_v17 = _rhs112;
        _778_v17 = _rhs113;
        _778_v17 = _rhs114;
        r1 = _rhs115;
        _778_v17 = _rhs116;
        _744_v1 = _module.__default.fm23((new BigNumber((_dafny.Seq.UnicodeFromString("cho")).length)).multipliedBy(_743_v0), globalState);
        let _787_v21;
        _787_v21 = _dafny.Map.Empty.slice().updateUnsafe((false) === (p0),_771_v13);
        _787_v21 = (_787_v21).update(_778_v17, _771_v13);
        let _788_v22;
        _788_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
        _778_v17 = ((_788_v22).Merge(_788_v22)).contains(_778_v17);
      }
      if (p0) {
        let _789_v23;
        let _nw146 = Array((new BigNumber(8)).toNumber()).fill([]);
        _789_v23 = _nw146;
        let _790_v24;
        _790_v24 = _dafny.Map.Empty.slice().updateUnsafe(_743_v0,_module.D0.create_DC0(_743_v0, (_this).f6, _789_v23));
        let _791_v25;
        _791_v25 = _dafny.Seq.of(_790_v24);
        let _792_v26;
        let _nw147 = new _module.C3();
        _nw147.__ctor();
        _792_v26 = _nw147;
        let _793_v27;
        _793_v27 = _module.D11.create_DC26(_791_v25, _792_v26);
        let _pat_let_tv10 = _790_v24;
        let _pat_let_tv11 = _790_v24;
        let _pat_let_tv12 = _790_v24;
        let _pat_let_tv13 = _790_v24;
        let _794_v28;
        _794_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf((_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), function (_795_i4) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })),function (_pat_let13_0) {
          return function (_796_dt__update__tmp_h2) {
            return function (_pat_let14_0) {
              return function (_797_dt__update_hcf42_h0) {
                return _module.D11.create_DC26(_797_dt__update_hcf42_h0, (_796_dt__update__tmp_h2).dtor_cf43);
              }(_pat_let14_0);
            }(_dafny.Seq.of(_pat_let_tv10, _pat_let_tv11, _pat_let_tv12, _pat_let_tv13));
          }(_pat_let13_0);
        }(_793_v27));
        _794_v28 = (_794_v28).Merge((_dafny.Map.Empty.slice().updateUnsafe(p0,_793_v27)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_793_v27)));
        let _798_v29;
        _798_v29 = _dafny.Seq.of(p0, p0, !(p0) || (p0), p0);
        _798_v29 = _dafny.Seq.Concat(_798_v29, _dafny.Seq.Concat(_798_v29, _798_v29));
        _743_v0 = _743_v0;
        let _799_v30;
        _799_v30 = true;
        let _800_v31;
        _800_v31 = _module.D0.create_DC0(_743_v0, _dafny.Seq.UnicodeFromString("mqbg"), _789_v23);
        _799_v30 = _module.__default.fm33((_800_v31).dtor_cf1, globalState);
        let _801_v32;
        let _nw148 = new _module.C1();
        _nw148.__ctor(_dafny.MultiSet.FromArray(_798_v29), _module.__default.safeDivisionInt(_743_v0, _743_v0));
        _801_v32 = _nw148;
      } else {
        let _802_v33;
        let _nw149 = new _module.C0();
        _nw149.__ctor();
        _802_v33 = _nw149;
        _743_v0 = _743_v0;
        _743_v0 = new BigNumber(-933);
        let _803_v34;
        _803_v34 = false;
        let _804_v35;
        _804_v35 = _dafny.Seq.of(_803_v34);
        _803_v34 = !(_803_v34) || ((_804_v35)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_803_v34,_803_v34)).length), new BigNumber((_804_v35).length))]);
        let _index155 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_771_v13).length));
        (_771_v13)[_index155] = _743_v0;
        let _index156 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_771_v13).length));
        (_771_v13)[_index156] = _743_v0;
      }
      let _805_v36;
      _805_v36 = false;
      _805_v36 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ybptaos"), _dafny.Seq.Concat((_this).f6, _dafny.Seq.UnicodeFromString("kgjwul")));
      let _806_v37;
      _806_v37 = _dafny.Seq.of(_dafny.Seq.of(_743_v0));
      let _807_i5;
      _807_i5 = _dafny.ZERO;
      L7: {
        while (((_743_v0).plus(_743_v0)).isEqualTo(new BigNumber((_dafny.MultiSet.FromArray((_806_v37)[_module.__default.safeIndex(_743_v0, new BigNumber((_806_v37).length))])).cardinality()))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_807_i5)) {
              break L7;
            }
            _807_i5 = (_807_i5).plus(_dafny.ONE);
            let _808_v38;
            _808_v38 = _dafny.Seq.UnicodeFromString("g");
            let _809_v39;
            _809_v39 = _dafny.Set.fromElements(_743_v0, (_dafny.ZERO).minus(_743_v0));
            let _810_v40;
            _810_v40 = _dafny.Map.Empty.slice().updateUnsafe(_743_v0,_809_v39);
            let _index157 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_774_v14).length));
            (_774_v14)[_index157] = ((((_810_v40).contains(_743_v0)) ? ((_810_v40).get(_743_v0)) : (_809_v39))).IsSubsetOf(_module.__default.fm20(_805_v36, _808_v38, _module.__default.fm33(_dafny.Seq.Create(_module.__default.abs(new BigNumber(290)), function (_811_i6) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            }), globalState), globalState));
            let _index158 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_774_v14).length));
            let _rhs117 = _dafny.Seq.IsProperPrefixOf((_this).f6, _dafny.Seq.Concat(_module.__default.fm22(_805_v36, _743_v0, globalState), (_this).f6));
            let _rhs118 = _dafny.Seq.Concat((_this).f6, _dafny.Seq.UnicodeFromString("xq"));
            let _rhs119 = _805_v36;
            let _rhs120 = _805_v36;
            let _rhs121 = _743_v0;
            let _lhs85 = _774_v14;
            let _lhs86 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_774_v14).length));
            _805_v36 = _rhs117;
            _808_v38 = _rhs118;
            _lhs85[_lhs86] = _rhs119;
            _805_v36 = _rhs120;
            r0 = _rhs121;
            let _812_v41;
            _812_v41 = _dafny.MultiSet.fromElements(new BigNumber(259), _module.__default.safeModuloInt(new BigNumber(-880), (_dafny.ZERO).minus(_743_v0)));
            _812_v41 = _812_v41;
            let _813_v42;
            _813_v42 = _module.D4.create_DC9(_771_v13);
            let _814_v43;
            let _nw150 = Array((new BigNumber(2)).toNumber());
            _nw150[(_dafny.ZERO).toNumber()] = _771_v13;
            _nw150[(_dafny.ONE).toNumber()] = (_813_v42).dtor_cf15;
            _814_v43 = _nw150;
            let _index159 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_814_v43).length));
            (_814_v43)[_index159] = _771_v13;
            let _815_v44;
            _815_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,_743_v0);
            let _816_v45;
            _816_v45 = _dafny.MultiSet.fromElements((_774_v14)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_774_v14).length))], p0, false, _805_v36);
            let _index160 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_814_v43).length));
            let _nw151 = Array((new BigNumber(24)).toNumber());
            _nw151[(_dafny.ZERO).toNumber()] = _743_v0;
            _nw151[(_dafny.ONE).toNumber()] = _743_v0;
            _nw151[(new BigNumber(2)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(3)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(4)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_743_v0).minus(_743_v0));
            _nw151[(new BigNumber(6)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(7)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_743_v0, (_dafny.ZERO).minus(_743_v0));
            _nw151[(new BigNumber(9)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(10)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(11)).toNumber()] = ((_805_v36) ? (new BigNumber((_815_v44).length)) : (_743_v0));
            _nw151[(new BigNumber(12)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(13)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(14)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(15)).toNumber()] = (new BigNumber((_dafny.Set.fromElements(!(_module.__default.fm33((_this).f6, globalState)))).length)).minus(new BigNumber(902));
            _nw151[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_805_v36)).length), new BigNumber(351));
            _nw151[(new BigNumber(17)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_743_v0);
            _nw151[(new BigNumber(19)).toNumber()] = _743_v0;
            _nw151[(new BigNumber(20)).toNumber()] = _module.__default.fm14((((_816_v45).contains(p0)) ? ((_816_v45).get(p0)) : (_743_v0)), _743_v0, globalState);
            _nw151[(new BigNumber(21)).toNumber()] = new BigNumber(((_this).f6).length);
            _nw151[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(!((_774_v14)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_774_v14).length))]))).cardinality());
            _nw151[(new BigNumber(23)).toNumber()] = _743_v0;
            (_814_v43)[_index160] = _nw151;
            let _817_v46;
            _817_v46 = _dafny.Seq.of(_805_v36, true);
            let _818_v47;
            _818_v47 = _dafny.Seq.of((_743_v0).isLessThanOrEqualTo(new BigNumber((_817_v46).length)));
            let _index161 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_774_v14).length));
            (_774_v14)[_index161] = (_818_v47)[_module.__default.safeIndex((_744_v1)[_module.__default.safeIndex(_743_v0, new BigNumber((_744_v1).length))], new BigNumber((_818_v47).length))];
          }
        }
      }
      r0 = _743_v0;
      let _819_v48;
      _819_v48 = new _dafny.CodePoint('j'.codePointAt(0));
      r1 = _819_v48;
      return [r0, r1];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this.f4 = _dafny.MultiSet.Empty;
      this._f5 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f4, f5) {
      let _this = this;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_this.f4).Union(_this.f4)).Difference(_this.f4);
    };
    fm9(globalState) {
      let _this = this;
      return new BigNumber(((((false) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(788)), function (_820_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length),false)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length),false)))).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)),true)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(440),true))))).length);
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.D2.create_DC6(new BigNumber(929), false)).dtor_cf12;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _821_v0;
      let _nw152 = Array((new BigNumber(22)).toNumber()).fill([]);
      _821_v0 = _nw152;
      let _822_v1;
      _822_v1 = _dafny.Seq.UnicodeFromString("rxl");
      let _823_v2;
      _823_v2 = new _dafny.CodePoint('v'.codePointAt(0));
      let _824_v3;
      _824_v3 = new BigNumber(403);
      let _825_v4;
      _825_v4 = false;
      let _826_v5;
      _826_v5 = _module.D1.create_DC2(true);
      let _827_v6;
      let _nw153 = Array((new BigNumber(23)).toNumber());
      _nw153[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(656)), function (_828_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      });
      _nw153[(_dafny.ONE).toNumber()] = _822_v1;
      _nw153[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("pet");
      _nw153[(new BigNumber(3)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("tpfmyj");
      _nw153[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("ipd");
      _nw153[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("voauqmys");
      _nw153[(new BigNumber(7)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(8)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(9)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(10)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(11)).toNumber()] = _module.__default.fm11(_823_v2, _824_v3, _825_v4, _826_v5, globalState);
      _nw153[(new BigNumber(12)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(13)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(14)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(15)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(16)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("a");
      _nw153[(new BigNumber(18)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(19)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("usguexh");
      _nw153[(new BigNumber(21)).toNumber()] = _822_v1;
      _nw153[(new BigNumber(22)).toNumber()] = _822_v1;
      _827_v6 = _nw153;
      let _index162 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_821_v0).length));
      (_821_v0)[_index162] = _827_v6;
      let _829_v7;
      _829_v7 = _module.D2.create_DC6(_824_v3, _825_v4);
      let _830_v8;
      _830_v8 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(166)), ((_831_v3) => function (_832_i1) {
        return _dafny.Seq.of(_831_v3, new BigNumber(294));
      })(_824_v3))).length));
      let _833_v9;
      _833_v9 = _dafny.Seq.of(_830_v8, _830_v8, _830_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), ((_834_v3) => function (_835_i2) {
        return _834_v3;
      })(_824_v3)), _830_v8);
      let _836_v10;
      _836_v10 = _dafny.MultiSet.fromElements(_824_v3);
      let _837_v11;
      _837_v11 = _dafny.Map.Empty.slice().updateUnsafe(_825_v4,_836_v10);
      let _index163 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_821_v0).length));
      let _rhs122 = _827_v6;
      let _rhs123 = (_this).fm10(_829_v7, (new BigNumber(((_833_v9)[_module.__default.safeIndex(new BigNumber((_module.__default.fm12(_824_v3, _824_v3, globalState)).length), new BigNumber((_833_v9).length))]).length)).plus(new BigNumber(933)), (((_837_v11).contains(_825_v4)) ? ((_837_v11).get(_825_v4)) : (_836_v10)), globalState);
      let _rhs124 = false;
      let _lhs87 = _821_v0;
      let _lhs88 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_821_v0).length));
      _lhs87[_lhs88] = _rhs122;
      r0 = _rhs123;
      _825_v4 = _rhs124;
      let _838_i3;
      _838_i3 = _dafny.ZERO;
      L8: {
        while (!(_dafny.Seq.IsProperPrefixOf(_822_v1, _dafny.Seq.UnicodeFromString("omtjm")))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_838_i3)) {
              break L8;
            }
            _838_i3 = (_838_i3).plus(_dafny.ONE);
            let _rhs125 = _826_v5;
            let _rhs126 = (_dafny.ZERO).minus(((_824_v3).plus(_824_v3)).minus(new BigNumber(199)));
            _826_v5 = _rhs125;
            _824_v3 = _rhs126;
            let _839_v12;
            _839_v12 = _dafny.Seq.of(false, _825_v4, true, _825_v4);
            _824_v3 = _module.__default.safeModuloInt(new BigNumber((_839_v12).length), _824_v3);
            let _840_v13;
            let _841_v14;
            let _842_v15;
            let _out14;
            let _out15;
            let _out16;
            let _outcollector7 = (_this).m2(globalState);
            _out14 = _outcollector7[0];
            _out15 = _outcollector7[1];
            _out16 = _outcollector7[2];
            _840_v13 = _out14;
            _841_v14 = _out15;
            _842_v15 = _out16;
            if (_842_v15) {
              let _index164 = _module.__default.safeIndex(new BigNumber(995), new BigNumber(((_this).f5).length));
              ((_this).f5)[_index164] = (_this).fm9(globalState);
              let _index165 = _module.__default.safeIndex(new BigNumber(995), new BigNumber(((_this).f5).length));
              ((_this).f5)[_index165] = (_829_v7).dtor_cf11;
              _830_v8 = _dafny.Seq.update(_dafny.Seq.Concat(_830_v8, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_843_i4) {
                return new BigNumber(174);
              }), _module.__default.safeIndex(_840_v13, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_844_i4) {
                return new BigNumber(174);
              })).length)), new BigNumber(844))), _module.__default.safeIndex((((_836_v10).contains(((_this).f5)[_module.__default.safeIndex(new BigNumber(995), new BigNumber(((_this).f5).length))])) ? ((_836_v10).get(((_this).f5)[_module.__default.safeIndex(new BigNumber(995), new BigNumber(((_this).f5).length))])) : (new BigNumber(-687))), new BigNumber((_dafny.Seq.Concat(_830_v8, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_845_i4) {
                return new BigNumber(174);
              }), _module.__default.safeIndex(_840_v13, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_846_i4) {
                return new BigNumber(174);
              })).length)), new BigNumber(844)))).length)), new BigNumber(-367));
              let _index166 = _module.__default.safeIndex(new BigNumber(995), new BigNumber(((_this).f5).length));
              ((_this).f5)[_index166] = _840_v13;
              let _index167 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_827_v6).length));
              (_827_v6)[_index167] = _822_v1;
              let _index168 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_827_v6).length));
              (_827_v6)[_index168] = ((!(_842_v15)) ? (_dafny.Seq.UnicodeFromString("dgiljws")) : (_822_v1));
              _840_v13 = (_830_v8)[_module.__default.safeIndex(((_this).f5)[_module.__default.safeIndex(new BigNumber(995), new BigNumber(((_this).f5).length))], new BigNumber((_830_v8).length))];
            } else {
              let _847_v16;
              let _nw154 = Array((new BigNumber(22)).toNumber()).fill(false);
              _847_v16 = _nw154;
              let _index169 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_847_v16).length));
              (_847_v16)[_index169] = _842_v15;
              let _index170 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_847_v16).length));
              (_847_v16)[_index170] = !((_this).fm10(_829_v7, _824_v3, _836_v10, globalState));
              _830_v8 = _dafny.Seq.update(_830_v8, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_822_v1).length)), new BigNumber((_830_v8).length)), _824_v3);
              _824_v3 = new BigNumber((_830_v8).length);
              let _index171 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_827_v6).length));
              (_827_v6)[_index171] = _dafny.Seq.UnicodeFromString("rvrjpar");
              let _index172 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_827_v6).length));
              (_827_v6)[_index172] = _822_v1;
              let _848_v17;
              let _nw155 = new _module.C1();
              _nw155.__ctor(_this.f4, _840_v13);
              _848_v17 = _nw155;
            }
          }
        }
      }
      let _849_i5;
      _849_i5 = _dafny.ZERO;
      L9: {
        while (_825_v4) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_849_i5)) {
              break L9;
            }
            _849_i5 = (_849_i5).plus(_dafny.ONE);
            let _850_v18;
            let _nw156 = Array((new BigNumber(19)).toNumber()).fill([]);
            _850_v18 = _nw156;
            let _851_v19;
            _851_v19 = _dafny.Map.Empty.slice().updateUnsafe(_824_v3,_module.D0.create_DC0(_824_v3, _822_v1, _850_v18));
            let _852_v20;
            _852_v20 = _dafny.Seq.of(_851_v19, _851_v19, _851_v19);
            let _853_v21;
            let _nw157 = new _module.C1();
            _nw157.__ctor(_this.f4, _824_v3);
            _853_v21 = _nw157;
            let _854_v22;
            _854_v22 = _module.D11.create_DC26(_852_v20, _853_v21);
            let _855_v23;
            _855_v23 = _dafny.MultiSet.fromElements(_854_v22);
            let _856_v24;
            _856_v24 = _dafny.Seq.of(_825_v4);
            let _857_v25;
            _857_v25 = _dafny.MultiSet.fromElements(_856_v24, _module.__default.fm12(_824_v3, _824_v3, globalState));
            let _rhs127 = _855_v23;
            let _rhs128 = _module.__default.safeDivisionInt(new BigNumber((_857_v25).cardinality()), ((_825_v4) ? (_824_v3) : (new BigNumber(579))));
            _855_v23 = _rhs127;
            _824_v3 = _rhs128;
            let _858_v26;
            _858_v26 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm33(_822_v1, globalState),_dafny.Map.Empty.slice().updateUnsafe(_856_v24,_824_v3));
            let _859_v27;
            _859_v27 = _dafny.Map.Empty.slice().updateUnsafe(_856_v24,_824_v3);
            _858_v26 = (_858_v26).update(_825_v4, _859_v27);
            let _860_v28;
            let _nw158 = new _module.C0();
            _nw158.__ctor();
            _860_v28 = _nw158;
            let _861_v29;
            _861_v29 = _module.D7.create_DC15(_825_v4, _824_v3, _824_v3, _824_v3, _860_v28);
            let _862_v30;
            _862_v30 = _module.D10.create_DC22();
            let _863_v31;
            _863_v31 = _dafny.Seq.of(_862_v30);
            let _864_v32;
            _864_v32 = _dafny.Map.Empty.slice().updateUnsafe(_825_v4,(_dafny.ZERO).minus(_824_v3));
            let _865_v33;
            _865_v33 = _module.D9.create_DC19(_824_v3, new BigNumber((_864_v32).length));
            let _866_v34;
            let _nw159 = Array((new BigNumber(18)).toNumber());
            _nw159[(_dafny.ZERO).toNumber()] = _824_v3;
            _nw159[(_dafny.ONE).toNumber()] = (_861_v29).dtor_cf22;
            _nw159[(new BigNumber(2)).toNumber()] = (new BigNumber(697)).multipliedBy(_824_v3);
            _nw159[(new BigNumber(3)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(4)).toNumber()] = new BigNumber(-819);
            _nw159[(new BigNumber(5)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(6)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_824_v3);
            _nw159[(new BigNumber(8)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(9)).toNumber()] = new BigNumber(-642);
            _nw159[(new BigNumber(10)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(11)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(12)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(13)).toNumber()] = _824_v3;
            _nw159[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_862_v30, _862_v30), _863_v31)).length);
            _nw159[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt((_865_v33).dtor_cf29, _824_v3);
            _nw159[(new BigNumber(16)).toNumber()] = (_this).fm9(globalState);
            _nw159[(new BigNumber(17)).toNumber()] = (_830_v8)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm24(_824_v3, true, _825_v4, _824_v3, globalState)), new BigNumber((_830_v8).length))];
            _866_v34 = _nw159;
            _866_v34 = _866_v34;
            _830_v8 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), ((_867_v3, _868_v4) => function (_869_i6) {
              return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_867_v3,_868_v4)).length)).minus(new BigNumber(668));
            })(_824_v3, _825_v4));
          }
        }
      }
      let _870_v35;
      let _nw160 = Array((new BigNumber(26)).toNumber()).fill(false);
      _870_v35 = _nw160;
      let _rhs129 = (_829_v7).dtor_cf12;
      let _rhs130 = _870_v35;
      r0 = _rhs129;
      _870_v35 = _rhs130;
      _824_v3 = _module.__default.safeModuloInt(_824_v3, _824_v3);
      let _871_v36;
      let _nw161 = new _module.C1();
      _nw161.__ctor(_this.f4, _824_v3);
      _871_v36 = _nw161;
      r0 = _825_v4;
      let _872_v37;
      _872_v37 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-879)),_this.f4);
      r1 = (_872_v37);
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let r2 = false;
      let _873_v0;
      _873_v0 = new BigNumber(392);
      r0 = _873_v0;
      let _874_v1;
      _874_v1 = false;
      let _875_v2;
      _875_v2 = _module.D10.create_DC21(_dafny.MultiSet.fromElements(_874_v1));
      _875_v2 = _875_v2;
      r2 = !(_874_v1);
      let _876_v3;
      let _nw162 = Array((new BigNumber(24)).toNumber()).fill(false);
      _876_v3 = _nw162;
      let _index173 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length));
      (_876_v3)[_index173] = _874_v1;
      let _index174 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length));
      (_876_v3)[_index174] = _874_v1;
      let _hi6 = _873_v0;
      for (let _877_i0 = _873_v0; _877_i0.isLessThan(_hi6); _877_i0 = _877_i0.plus(_dafny.ONE)) {
        r0 = _877_i0;
        _873_v0 = (_877_i0).minus(_877_i0);
        let _878_v4;
        let _nw163 = Array((new BigNumber(2)).toNumber());
        _nw163[(_dafny.ZERO).toNumber()] = (_this).f5;
        _nw163[(_dafny.ONE).toNumber()] = (_this).f5;
        _878_v4 = _nw163;
        let _879_v5;
        _879_v5 = _module.D0.create_DC0(_873_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), function (_880_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}), _878_v4);
        r2 = _module.__default.fm33((_879_v5).dtor_cf1, globalState);
        if ((_876_v3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length))]) {
          let _881_v6;
          _881_v6 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_877_i0, new BigNumber(901)));
          _881_v6 = (_881_v6).Union((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(776)), ((_882_v0) => function (_883_i2) {
            return _882_v0;
          })(_873_v0)))).Union(_881_v6));
          let _884_v7;
          _884_v7 = new _dafny.CodePoint('t'.codePointAt(0));
          let _885_v8;
          _885_v8 = _dafny.Seq.UnicodeFromString("cu");
          let _886_v9;
          _886_v9 = _dafny.Set.fromElements(_885_v8);
          let _887_v10;
          _887_v10 = _dafny.Seq.of(!((_dafny.ZERO).minus(_873_v0)).isEqualTo(new BigNumber((_886_v9).length)));
          let _index175 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length));
          let _rhs131 = (_887_v10)[_module.__default.safeIndex((_877_i0).minus(_877_i0), new BigNumber((_887_v10).length))];
          let _rhs132 = new _dafny.CodePoint('f'.codePointAt(0));
          let _rhs133 = _873_v0;
          let _lhs89 = _876_v3;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length));
          _lhs89[_lhs90] = _rhs131;
          _884_v7 = _rhs132;
          r0 = _rhs133;
          _884_v7 = _884_v7;
          let _888_v11;
          let _nw164 = Array((new BigNumber(19)).toNumber()).fill([]);
          _888_v11 = _nw164;
          let _889_v12;
          _889_v12 = _dafny.Seq.of(_877_i0);
          let _890_v13;
          _890_v13 = _dafny.Map.Empty.slice().updateUnsafe((_876_v3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length))],_876_v3);
          let _891_v14;
          _891_v14 = _876_v3;
          let _892_v15;
          let _nw165 = Array((new BigNumber(3)).toNumber());
          _nw165[(_dafny.ZERO).toNumber()] = _876_v3;
          _nw165[(_dafny.ONE).toNumber()] = _876_v3;
          _nw165[(new BigNumber(2)).toNumber()] = (((_890_v13).contains(_874_v1)) ? ((_890_v13).get(_874_v1)) : ((_891_v14)));
          _892_v15 = _nw165;
          let _893_v16;
          _893_v16 = _dafny.Map.Empty.slice().updateUnsafe(_889_v12,_892_v15);
          let _index176 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_888_v11).length));
          (_888_v11)[_index176] = (((_893_v16).contains(_889_v12)) ? ((_893_v16).get(_889_v12)) : (_892_v15));
          let _index177 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_888_v11).length));
          (_888_v11)[_index177] = _892_v15;
          _874_v1 = (_876_v3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length))];
        } else {
          _873_v0 = _877_i0;
          let _894_v17;
          _894_v17 = _module.D3.create_DC8(((_874_v1) ? ((_dafny.ZERO).minus(_873_v0)) : (_877_i0)));
          _894_v17 = _894_v17;
          let _895_v18;
          _895_v18 = new _dafny.CodePoint('w'.codePointAt(0));
          _895_v18 = _895_v18;
          let _896_v19;
          _896_v19 = _dafny.Seq.UnicodeFromString("xjisbuee");
          let _897_v20;
          _897_v20 = _dafny.Seq.of(_896_v19, _896_v19, _896_v19);
          let _898_v21;
          _898_v21 = _897_v20;
          let _899_v22;
          _899_v22 = _dafny.Map.Empty.slice().updateUnsafe(_873_v0,new BigNumber(((_898_v21)).length));
          _899_v22 = (_899_v22).update(_873_v0, _873_v0);
          let _index178 = _module.__default.safeIndex(new BigNumber(651), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index178] = _877_i0;
          let _index179 = _module.__default.safeIndex(new BigNumber(651), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index179] = _877_i0;
        }
      }
      let _900_v23;
      _900_v23 = _module.D4.create_DC10(_874_v1);
      let _901_v24;
      _901_v24 = _dafny.Seq.UnicodeFromString("e");
      _874_v1 = (((_module.__default.fm16(_900_v23, globalState)).IsSubsetOf(_dafny.Set.fromElements(_874_v1, (_876_v3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length))]))) ? (_module.__default.fm33(_901_v24, globalState)) : ((_876_v3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length))]));
      let _902_v25;
      _902_v25 = _dafny.Seq.of(_876_v3);
      let _903_v27;
      _903_v27 = new _dafny.CodePoint('v'.codePointAt(0));
      let _904_v28;
      _904_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm33(_901_v24, globalState),new BigNumber((function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of (_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), _903_v27)).Elements) {
          let _905_v26 = _compr_27;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), _903_v27), _905_v26)) {
            _coll27.push([_905_v26,_873_v0]);
          }
        }
        return _coll27;
      }()).length));
      r0 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_902_v25).length), _873_v0), (((_904_v28).contains(_module.__default.fm33(_901_v24, globalState))) ? ((_904_v28).get(_module.__default.fm33(_901_v24, globalState))) : (_873_v0)));
      let _906_v29;
      _906_v29 = _dafny.Seq.of((_dafny.ZERO).minus(_873_v0), new BigNumber((_901_v24).length));
      let _907_v30;
      _907_v30 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC5(_906_v29),new BigNumber(101));
      let _908_v31;
      _908_v31 = _dafny.Map.Empty.slice().updateUnsafe((_907_v30).Merge(_907_v30),_874_v1);
      r1 = _908_v31;
      r2 = (_876_v3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_876_v3).length))];
      return [r0, r1, r2];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f3 = false;
      this._f2 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f2, f3) {
      let _this = this;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements((_this).f3, (new BigNumber(-476)).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f3, (_this).f3, (_this).f3)).cardinality()))));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of _dafny.IntegerRange(new BigNumber(747), new BigNumber(-512))) {
          let _909_v0 = _compr_28;
          if (((new BigNumber(747)).isLessThanOrEqualTo(_909_v0)) && ((_909_v0).isLessThan(new BigNumber(-512)))) {
            _coll28.push([(_909_v0).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gtxa")).length))),(_this).f3]);
          }
        }
        return _coll28;
      }()).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("cwbgbfuq")).length))).length))).length))).length),(_this).f3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(512),(_this).f3))).length)).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f3,!((_this).f3))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-762))).length))));
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("xxbvgut");
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _910_v0;
      _910_v0 = new BigNumber(-934);
      _910_v0 = _910_v0;
      let _911_v1;
      let _init15 = function (_912_i0) {
        return (_this).f3;
      };
      let _nw166 = Array((new BigNumber(28)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw166.length); _i0_15++) {
        _nw166[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _911_v1 = _nw166;
      let _913_v2;
      _913_v2 = new _dafny.CodePoint('o'.codePointAt(0));
      let _914_v3;
      _914_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_913_v2);
      let _915_v4;
      _915_v4 = _dafny.MultiSet.fromElements(_914_v3);
      let _index180 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
      (_911_v1)[_index180] = ((_915_v4).update((_914_v3).update((_this).f3, _913_v2), _module.__default.abs(new BigNumber(161)))).IsDisjointFrom(_915_v4);
      let _916_v5;
      _916_v5 = _dafny.MultiSet.fromElements(_910_v0, _910_v0, _910_v0);
      let _index181 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
      (_911_v1)[_index181] = _module.__default.fm8(_916_v5, globalState);
      if ((((_this).f3) ? ((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) : ((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]))) {
        r0 = (_this).f3;
        let _917_v6;
        _917_v6 = _dafny.Seq.UnicodeFromString("prubmuoud");
        let _918_v7;
        _918_v7 = _dafny.Seq.of(_910_v0);
        let _919_v8;
        _919_v8 = _module.D2.create_DC5(_918_v7);
        let _920_v9;
        _920_v9 = _module.D1.create_DC4((_dafny.ZERO).minus(_910_v0), _913_v2, false, new BigNumber((_917_v6).length));
        let _921_v10;
        let _nw167 = new _module.C2();
        _nw167.__ctor(_917_v6, (_this).fm7(_919_v8, _910_v0, _919_v8, _920_v9, globalState));
        _921_v10 = _nw167;
        let _index182 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
        (_911_v1)[_index182] = (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))];
        let _922_v11;
        _922_v11 = _dafny.Map.Empty.slice().updateUnsafe(_913_v2,_910_v0);
        let _923_v12;
        _923_v12 = _dafny.Map.Empty.slice().updateUnsafe(_913_v2,new BigNumber(((_922_v11).Merge(_922_v11)).length));
        _923_v12 = (_923_v12).update(_module.__default.fm30(_917_v6, _910_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(384)), ((_924_v2) => function (_925_i1) {
          return _924_v2;
        })(_913_v2))).length), globalState), _910_v0);
        _917_v6 = _917_v6;
      } else {
        let _926_v13;
        _926_v13 = _dafny.Seq.UnicodeFromString("snuoik");
        let _927_v14;
        let _nw168 = new _module.C5();
        _nw168.__ctor(_926_v13);
        _927_v14 = _nw168;
        _910_v0 = _910_v0;
        let _928_v15;
        _928_v15 = _module.D1.create_DC2((_this).f3);
        let _929_v16;
        _929_v16 = _dafny.MultiSet.fromElements((_this).f3, (_928_v15).dtor_cf4);
        let _930_v17;
        let _nw169 = new _module.C1();
        _nw169.__ctor(_929_v16, _910_v0);
        _930_v17 = _nw169;
        if ((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) {
          let _931_v18;
          let _nw170 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _931_v18 = _nw170;
          let _index183 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_931_v18).length));
          (_931_v18)[_index183] = _930_v17.f11;
          let _932_v19;
          _932_v19 = _dafny.Seq.of(_930_v17.f11);
          let _index184 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_931_v18).length));
          (_931_v18)[_index184] = (((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-709)), ((_933_v2) => function (_934_i2) {
            return _933_v2;
          })(_913_v2))).length)).isLessThan(new BigNumber((_932_v19).length))) ? (_930_v17.f11) : ((_910_v0).plus(new BigNumber((_dafny.MultiSet.fromElements(_930_v17.f11, _930_v17.f11, _910_v0)).cardinality()))));
          (_930_v17).f11 = (_930_v17.f11).plus(_module.__default.safeDivisionInt((_931_v18)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_931_v18).length))], _930_v17.f11));
          let _index185 = _module.__default.safeIndex(new BigNumber(715), new BigNumber(((_this).f2).length));
          ((_this).f2)[_index185] = _931_v18;
          let _index186 = _module.__default.safeIndex(new BigNumber(715), new BigNumber(((_this).f2).length));
          ((_this).f2)[_index186] = _931_v18;
          let _index187 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
          (_911_v1)[_index187] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_927_v14).f6, (_927_v14).f6), (_927_v14).f6);
          r0 = (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))];
        } else {
          let _935_v20;
          let _nw171 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _935_v20 = _nw171;
          let _936_v21;
          _936_v21 = _dafny.Seq.of((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]);
          let _index188 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_935_v20).length));
          (_935_v20)[_index188] = _dafny.Seq.update(_936_v21, _module.__default.safeIndex(_930_v17.f11, new BigNumber((_936_v21).length)), (_936_v21)[_module.__default.safeIndex(_910_v0, new BigNumber((_936_v21).length))]);
          let _937_v22;
          let _nw172 = Array((new BigNumber(8)).toNumber());
          _nw172[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_930_v17.f11);
          _nw172[(_dafny.ONE).toNumber()] = (((_916_v5).contains(_930_v17.f11)) ? ((_916_v5).get(_930_v17.f11)) : (new BigNumber((_936_v21).length)));
          _nw172[(new BigNumber(2)).toNumber()] = _930_v17.f11;
          _nw172[(new BigNumber(3)).toNumber()] = _930_v17.f11;
          _nw172[(new BigNumber(4)).toNumber()] = _910_v0;
          _nw172[(new BigNumber(5)).toNumber()] = (_927_v14).fm13((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))], _930_v17.f11, (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))], globalState);
          _nw172[(new BigNumber(6)).toNumber()] = _930_v17.f11;
          _nw172[(new BigNumber(7)).toNumber()] = _930_v17.f11;
          _937_v22 = _nw172;
          let _index189 = _module.__default.safeIndex(new BigNumber(628), new BigNumber(((_this).f2).length));
          ((_this).f2)[_index189] = _937_v22;
          let _938_v23;
          _938_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_926_v13).length),_929_v16);
          let _939_v24;
          _939_v24 = _938_v23;
          let _940_v25;
          _940_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_930_v17.f11);
          let _index190 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_935_v20).length));
          let _index191 = _module.__default.safeIndex(new BigNumber(628), new BigNumber(((_this).f2).length));
          let _rhs134 = _dafny.Seq.Concat(_936_v21, _dafny.Seq.update(_936_v21, _module.__default.safeIndex(new BigNumber((_940_v25).length), new BigNumber((_936_v21).length)), !((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))])));
          let _rhs135 = (_this).f3;
          let _rhs136 = _913_v2;
          let _rhs137 = _937_v22;
          let _rhs138 = _939_v24;
          let _lhs91 = _935_v20;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_935_v20).length));
          let _lhs93 = (_this).f2;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(628), new BigNumber(((_this).f2).length));
          _lhs91[_lhs92] = _rhs134;
          r0 = _rhs135;
          _913_v2 = _rhs136;
          _lhs93[_lhs94] = _rhs137;
          _939_v24 = _rhs138;
          let _941_v26;
          _941_v26 = _dafny.Map.Empty.slice().updateUnsafe((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))],!(!(!((_this).f3))));
          _941_v26 = (_941_v26).update((_this).f3, (_this).f3);
          _926_v13 = _dafny.Seq.Concat((_927_v14).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), ((_942_v2) => function (_943_i3) {
            return _942_v2;
          })(_913_v2)));
          let _index192 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
          (_911_v1)[_index192] = (((_941_v26).contains((((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) ? ((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) : ((_this).f3)))) ? ((_941_v26).get((((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) ? ((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) : ((_this).f3)))) : (true));
          let _index193 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
          (_911_v1)[_index193] = !(false);
        }
        let _944_v27;
        _944_v27 = _dafny.Seq.of(true);
        r0 = (_944_v27)[_module.__default.safeIndex(_930_v17.f11, new BigNumber((_944_v27).length))];
      }
      _913_v2 = _913_v2;
      let _945_v28;
      _945_v28 = _dafny.Map.Empty.slice().updateUnsafe(_910_v0,_916_v5);
      let _946_i4;
      _946_i4 = _dafny.ZERO;
      L10: {
        while ((_module.__default.fm31(new BigNumber((function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of _dafny.IntegerRange(new BigNumber(213), new BigNumber(546))) {
            let _973_v29 = _compr_31;
            if (((new BigNumber(213)).isLessThanOrEqualTo(_973_v29)) && ((_973_v29).isLessThan(new BigNumber(546)))) {
              _coll31.push([_module.__default.safeDivisionInt(_973_v29, _910_v0),_910_v0]);
            }
          }
          return _coll31;
        }()).length), globalState)).IsProperSubsetOf((((_945_v28).contains(_910_v0)) ? ((_945_v28).get(_910_v0)) : (_916_v5)))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_946_i4)) {
              break L10;
            }
            _946_i4 = (_946_i4).plus(_dafny.ONE);
            r0 = (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))];
            if ((_this).f3) {
              _910_v0 = _910_v0;
              let _947_v30;
              let _nw173 = new _module.C4();
              _nw173.__ctor((_dafny.ZERO).minus(_910_v0));
              _947_v30 = _nw173;
              let _948_v31;
              _948_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_947_v30.f7, new BigNumber(257)),(_this).f3);
              let _949_v32;
              _949_v32 = _dafny.Seq.of(new BigNumber(-445), _910_v0);
              _948_v31 = ((_948_v31).update(new BigNumber((_949_v32).length), (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))])).Merge(_948_v31);
              let _950_v33;
              let _nw174 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _950_v33 = _nw174;
              let _951_v34;
              _951_v34 = _dafny.Set.fromElements(_910_v0);
              let _952_v35;
              _952_v35 = _module.D1.create_DC4(new BigNumber((_951_v34).length), _913_v2, !((_this).f3), new BigNumber(-214));
              let _953_v36;
              _953_v36 = _dafny.Seq.of(_913_v2, _913_v2, _913_v2, (_952_v35).dtor_cf7, _913_v2);
              let _954_v37;
              _954_v37 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f3,(_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]));
              let _index194 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_950_v33).length));
              (_950_v33)[_index194] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_953_v36, _953_v36), _module.__default.safeIndex(new BigNumber((_954_v37).length), new BigNumber((_dafny.Seq.Concat(_953_v36, _953_v36)).length)), _913_v2)).length);
              let _index195 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_950_v33).length));
              (_950_v33)[_index195] = _910_v0;
              let _index196 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_950_v33).length));
              let _index197 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
              let _rhs139 = _module.__default.safeModuloInt(_947_v30.f7, (_910_v0).minus(_947_v30.f7));
              let _rhs140 = true;
              let _rhs141 = _953_v36;
              let _rhs142 = _910_v0;
              let _lhs95 = _950_v33;
              let _lhs96 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_950_v33).length));
              let _lhs97 = _911_v1;
              let _lhs98 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
              let _lhs99 = _947_v30;
              _lhs95[_lhs96] = _rhs139;
              _lhs97[_lhs98] = _rhs140;
              _953_v36 = _rhs141;
              _lhs99.f7 = _rhs142;
            } else {
              _913_v2 = _913_v2;
              let _955_v38;
              _955_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_910_v0),(_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]);
              let _956_v39;
              _956_v39 = _dafny.Seq.UnicodeFromString("ftwmf");
              let _pat_let_tv14 = _911_v1;
              let _pat_let_tv15 = _911_v1;
              let _pat_let_tv16 = _955_v38;
              let _pat_let_tv17 = globalState;
              let _957_v40;
              _957_v40 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f3) ? (function (_pat_let15_0) {
                return function (_958_dt__update__tmp_h0) {
                  return function (_pat_let16_0) {
                    return function (_959_dt__update_hcf14_h0) {
                      return _module.D3.create_DC8(_959_dt__update_hcf14_h0);
                    }(_pat_let16_0);
                  }(_module.__default.fm19((_pat_let_tv15)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_pat_let_tv14).length))], new BigNumber((_pat_let_tv16).length), _pat_let_tv17));
                }(_pat_let15_0);
              }(_module.D3.create_DC8(_910_v0))) : (_module.D3.create_DC8(new BigNumber((_dafny.MultiSet.fromElements((_this).f3)).cardinality())))),_module.__default.fm33(_956_v39, globalState));
              let _960_v42;
              _960_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,false);
              let _961_v43;
              _961_v43 = _dafny.Map.Empty.slice().updateUnsafe((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))],new BigNumber((_960_v42).length));
              let _962_v44;
              _962_v44 = _dafny.Seq.of((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]);
              let _963_v45;
              _963_v45 = _dafny.Seq.of(new BigNumber((_962_v44).length), _910_v0);
              let _964_v46;
              _964_v46 = _dafny.Set.fromElements(_961_v43, _961_v43, (_961_v43).update((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))], new BigNumber((_963_v45).length)), _961_v43);
              let _965_v47;
              _965_v47 = _module.D1.create_DC3(new BigNumber((_964_v46).length));
              let _966_v48;
              _966_v48 = _module.D3.create_DC8((_965_v47).dtor_cf5);
              let _967_v49;
              _967_v49 = _dafny.Seq.of(_966_v48);
              _957_v40 = function () {
                let _coll29 = new _dafny.Map();
                for (const _compr_29 of (_967_v49).Elements) {
                  let _968_v41 = _compr_29;
                  if (_dafny.Seq.contains(_967_v49, _968_v41)) {
                    _coll29.push([_968_v41,(_this).f3]);
                  }
                }
                return _coll29;
              }();
              let _index198 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
              (_911_v1)[_index198] = (((_916_v5).update(_910_v0, _module.__default.abs(new BigNumber(469)))).Difference(_916_v5)).IsProperSubsetOf(_916_v5);
              let _index199 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
              (_911_v1)[_index199] = (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))];
              let _index200 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
              (_911_v1)[_index200] = ((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) && ((_this).f3);
            }
            let _969_v50;
            _969_v50 = _dafny.Map.Empty.slice().updateUnsafe(!((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) || ((_this).f3),_911_v1);
            let _970_v51;
            _970_v51 = _dafny.Seq.of((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))], true);
            let _index201 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
            let _rhs143 = (((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]) ? (!_dafny.areEqual(_970_v51, _970_v51)) : (!(!((new BigNumber(314)).isLessThan(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
              let _coll30 = new _dafny.Map();
              for (const _compr_30 of _dafny.IntegerRange(new BigNumber(902), new BigNumber(637))) {
                let _971_v52 = _compr_30;
                if (((new BigNumber(902)).isLessThanOrEqualTo(_971_v52)) && ((_971_v52).isLessThan(new BigNumber(637)))) {
                  _coll30.push([_module.__default.safeDivisionInt(_971_v52, new BigNumber(470)),false]);
                }
              }
              return _coll30;
            }()).length), _910_v0, new BigNumber(119))).length))))));
            let _rhs144 = (_969_v50).update((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))], _911_v1);
            let _rhs145 = !((_this).f3) || ((_this).f3);
            let _rhs146 = new BigNumber((_dafny.Seq.UnicodeFromString("g")).length);
            let _lhs100 = _911_v1;
            let _lhs101 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
            r0 = _rhs143;
            _969_v50 = _rhs144;
            _lhs100[_lhs101] = _rhs145;
            _910_v0 = _rhs146;
            let _972_v53;
            _972_v53 = _dafny.Seq.of(_910_v0);
            let _index202 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
            let _rhs147 = (_this).f3;
            let _rhs148 = _972_v53;
            let _lhs102 = _911_v1;
            let _lhs103 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
            _lhs102[_lhs103] = _rhs147;
            _972_v53 = _rhs148;
          }
        }
      }
      let _974_v54;
      _974_v54 = _dafny.Map.Empty.slice().updateUnsafe(_910_v0,(_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]);
      let _index203 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length));
      (_911_v1)[_index203] = (((_974_v54).contains(_910_v0)) ? ((_974_v54).get(_910_v0)) : ((_this).f3));
      r0 = (_this).f3;
      let _975_v55;
      _975_v55 = _dafny.MultiSet.fromElements((_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))], (_911_v1)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_911_v1).length))]);
      let _976_v56;
      _976_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_974_v54).length),_975_v55);
      r1 = (_976_v56).Merge(_976_v56);
      return [r0, r1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(false, !(!(!(false))), !(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(78),true)).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, true)).length)))).cardinality()),false)));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length)))).Intersect((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(276))).cardinality()))).length), new BigNumber(-163))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(136)), function (_977_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-353),new BigNumber((_dafny.Seq.UnicodeFromString("nyftlaqe")).length))).length);
      }))));
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return !(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("dyqqsd"), _dafny.Seq.UnicodeFromString("mtkslws")));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _978_v0;
      _978_v0 = new BigNumber(-859);
      let _hi7 = _978_v0;
      for (let _979_i0 = _978_v0; _979_i0.isLessThan(_hi7); _979_i0 = _979_i0.plus(_dafny.ONE)) {
        let _980_v1;
        _980_v1 = true;
        let _981_v2;
        _981_v2 = _dafny.Set.fromElements(_980_v1, _980_v1);
        let _982_v3;
        _982_v3 = _dafny.Set.fromElements(_981_v2);
        _982_v3 = _module.__default.fm3(globalState);
        let _983_v4;
        _983_v4 = _dafny.Seq.UnicodeFromString("hwkemvrcr");
        let _984_v5;
        _984_v5 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), function (_985_i1) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }), _983_v4);
        _978_v0 = _module.__default.safeModuloInt(_978_v0, new BigNumber(((_984_v5)[_module.__default.safeIndex(_978_v0, new BigNumber((_984_v5).length))]).length));
        let _986_v6;
        let _nw175 = Array((new BigNumber(24)).toNumber()).fill([]);
        _986_v6 = _nw175;
        let _987_v7;
        _987_v7 = _module.D0.create_DC0(_978_v0, _983_v4, _986_v6);
        let _988_v8;
        _988_v8 = _module.D0.create_DC1(_module.D0.create_DC1(_987_v7));
        let _989_v9;
        _989_v9 = _module.D0.create_DC1(_module.D0.create_DC1(_988_v8));
        let _source10 = _989_v9;
        if (_source10.is_DC0) {
          let _990___mcc_h0 = (_source10).cf0;
          let _991___mcc_h1 = (_source10).cf1;
          let _992___mcc_h2 = (_source10).cf2;
          let _993_cf2 = _992___mcc_h2;
          let _994_cf1 = _991___mcc_h1;
          let _995_cf0 = _990___mcc_h0;
          let _996_v10;
          _996_v10 = _dafny.Map.Empty.slice().updateUnsafe(_980_v1,_983_v4);
          let _997_v11;
          _997_v11 = new _dafny.CodePoint('s'.codePointAt(0));
          let _998_v12;
          _998_v12 = _module.D1.create_DC4(_978_v0, _997_v11, _980_v1, _978_v0);
          let _999_v13;
          let _nw176 = Array((new BigNumber(23)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = (((_996_v10).contains(!(_980_v1))) ? ((_996_v10).get(!(_980_v1))) : (_983_v4));
          _nw176[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_983_v4, _module.__default.safeIndex(_978_v0, new BigNumber((_983_v4).length)), _997_v11);
          _nw176[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("p");
          _nw176[(new BigNumber(3)).toNumber()] = _983_v4;
          _nw176[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), ((_1000_v11) => function (_1001_i2) {
            return _1000_v11;
          })(_997_v11)), _994_cf1);
          _nw176[(new BigNumber(5)).toNumber()] = _983_v4;
          _nw176[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_994_cf1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), ((_1002_v11) => function (_1003_i3) {
            return _1002_v11;
          })(_997_v11))), _module.__default.safeIndex(_978_v0, new BigNumber((_dafny.Seq.Concat(_994_cf1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), ((_1004_v11) => function (_1005_i3) {
            return _1004_v11;
          })(_997_v11)))).length)), new _dafny.CodePoint('c'.codePointAt(0)));
          _nw176[(new BigNumber(7)).toNumber()] = _983_v4;
          _nw176[(new BigNumber(8)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pmh"), _dafny.Seq.UnicodeFromString("moycxky"));
          _nw176[(new BigNumber(10)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(11)).toNumber()] = _983_v4;
          _nw176[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(458)), ((_1006_v11) => function (_1007_i4) {
            return _1006_v11;
          })(_997_v11)), _983_v4);
          _nw176[(new BigNumber(13)).toNumber()] = _983_v4;
          _nw176[(new BigNumber(14)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(15)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(16)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_994_cf1, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(53)), ((_1008_v0) => function (_1009_i5) {
            return _1008_v0;
          })(_978_v0))).length), new BigNumber((_994_cf1).length)), _997_v11);
          _nw176[(new BigNumber(18)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(19)).toNumber()] = _994_cf1;
          _nw176[(new BigNumber(20)).toNumber()] = _983_v4;
          _nw176[(new BigNumber(21)).toNumber()] = (((_998_v12).dtor_cf8) ? (_994_cf1) : (_994_cf1));
          _nw176[(new BigNumber(22)).toNumber()] = _module.__default.fm4(_980_v1, _980_v1, _980_v1, globalState);
          _999_v13 = _nw176;
          _999_v13 = _999_v13;
          let _1010_v14;
          _1010_v14 = _dafny.MultiSet.fromElements(_978_v0, _995_cf0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(578)), ((_1011_v11) => function (_1012_i6) {
            return _1011_v11;
          })(_997_v11))).length));
          let _1013_v15;
          _1013_v15 = _dafny.Map.Empty.slice().updateUnsafe(_995_cf0,_1010_v14);
          _1013_v15 = (_1013_v15).update(_995_cf0, _dafny.MultiSet.fromElements(_978_v0));
          r0 = (_980_v1) === (_980_v1);
          let _1014_v16;
          _1014_v16 = _dafny.Seq.of(_980_v1);
          let _1015_v17;
          _1015_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(_980_v1, _980_v1, _980_v1), _1014_v16),_1014_v16);
          _1015_v17 = (_1015_v17).update(_1014_v16, _1014_v16);
        } else {
          let _1016___mcc_h3 = (_source10).cf3;
          let _1017_cf3 = _1016___mcc_h3;
          let _1018_v18;
          let _nw177 = Array((new BigNumber(15)).toNumber()).fill(_module.D1.Default());
          _1018_v18 = _nw177;
          let _1019_v19;
          _1019_v19 = _module.D1.create_DC3(_978_v0);
          let _index204 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1018_v18).length));
          (_1018_v18)[_index204] = _1019_v19;
          let _index205 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1018_v18).length));
          (_1018_v18)[_index205] = _1019_v19;
          let _1020_v20;
          _1020_v20 = _module.D1.create_DC2(_980_v1);
          _1020_v20 = _1020_v20;
          let _1021_v22;
          _1021_v22 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("abvray"))).update(_983_v4, _module.__default.abs(new BigNumber((_983_v4).length)))).Elements) {
              let _1022_v21 = _compr_32;
              if (((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("abvray"))).update(_983_v4, _module.__default.abs(new BigNumber((_983_v4).length)))).contains(_1022_v21)) {
                _coll32.push([_1022_v21,(_module.D2.create_DC5(_dafny.Seq.of(_978_v0))).dtor_cf10]);
              }
            }
            return _coll32;
          }(),_980_v1);
          _1021_v22 = (_1021_v22).update(_module.__default.fm5(globalState), !((_980_v1) || (_980_v1)));
          let _1023_v23;
          _1023_v23 = _dafny.MultiSet.fromElements(_980_v1);
          let _1024_v24;
          let _nw178 = new _module.C1();
          _nw178.__ctor((_1023_v23).update(_980_v1, _module.__default.abs(_978_v0)), _978_v0);
          _1024_v24 = _nw178;
        }
        let _1025_v25;
        let _nw179 = new _module.C2();
        _nw179.__ctor(_983_v4, _983_v4);
        _1025_v25 = _nw179;
        let _1026_v26;
        _1026_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1025_v25,new BigNumber(219));
        let _1027_v27;
        _1027_v27 = _dafny.Map.Empty.slice().updateUnsafe(_980_v1,_1026_v26);
        _1026_v26 = ((_980_v1) ? (_1026_v26) : ((((_1027_v27).contains(_980_v1)) ? ((_1027_v27).get(_980_v1)) : (_1026_v26))));
      }
      let _hi8 = (_978_v0).multipliedBy(_978_v0);
      for (let _1028_i7 = (_dafny.ZERO).minus(_978_v0); _1028_i7.isLessThan(_hi8); _1028_i7 = _1028_i7.plus(_dafny.ONE)) {
        r0 = (_1028_i7).isLessThan(_978_v0);
        let _1029_v28;
        _1029_v28 = _dafny.Seq.UnicodeFromString("erb");
        let _1030_v29;
        let _init16 = ((_1031_v0) => function (_1032_i8) {
          return _dafny.Seq.of(_1031_v0, _1031_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length));
        })(_978_v0);
        let _nw180 = Array((new BigNumber(10)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw180.length); _i0_16++) {
          _nw180[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _1030_v29 = _nw180;
        let _1033_v30;
        let _nw181 = new _module.C0();
        _nw181.__ctor();
        _1033_v30 = _nw181;
        let _1034_v31;
        _1034_v31 = false;
        let _1035_v32;
        _1035_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1034_v31,_1028_i7);
        let _1036_v33;
        _1036_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1034_v31,_1034_v31);
        let _rhs149 = _1029_v28;
        let _rhs150 = _1030_v29;
        let _rhs151 = (((_1035_v32).contains(!(_1034_v31) || (_1034_v31))) ? ((_1035_v32).get(!(_1034_v31) || (_1034_v31))) : (((_1034_v31) ? (_1028_i7) : (_978_v0))));
        let _rhs152 = _1033_v30;
        let _rhs153 = ((_dafny.ZERO).minus(new BigNumber(((_1036_v33).Merge(_1036_v33)).length))).multipliedBy(new BigNumber(380));
        _1029_v28 = _rhs149;
        _1030_v29 = _rhs150;
        _978_v0 = _rhs151;
        _1033_v30 = _rhs152;
        _978_v0 = _rhs153;
        let _1037_v34;
        _1037_v34 = _module.D9.create_DC20(_1034_v31, _1034_v31, _1034_v31);
        let _pat_let_tv18 = _1034_v31;
        let _1038_v35;
        _1038_v35 = _dafny.Seq.of(function (_pat_let17_0) {
          return function (_1039_dt__update__tmp_h0) {
            return function (_pat_let18_0) {
              return function (_1040_dt__update_hcf31_h0) {
                return _module.D9.create_DC20(_1040_dt__update_hcf31_h0, (_1039_dt__update__tmp_h0).dtor_cf32, (_1039_dt__update__tmp_h0).dtor_cf33);
              }(_pat_let18_0);
            }(_pat_let_tv18);
          }(_pat_let17_0);
        }(_module.D9.create_DC20(_1034_v31, _1034_v31, _1034_v31)), _module.D9.create_DC20(_1034_v31, true, _1034_v31), _1037_v34, _1037_v34, _1037_v34);
        let _1041_v36;
        _1041_v36 = _module.D14.create_DC29(_1038_v35);
        _1038_v35 = (_1041_v36).dtor_cf46;
        _1034_v31 = _1034_v31;
      }
      let _1042_v37;
      _1042_v37 = false;
      let _1043_v38;
      _1043_v38 = _dafny.Seq.of(_1042_v37);
      let _1044_i9;
      _1044_i9 = _dafny.ZERO;
      L11: {
        while (_dafny.Seq.contains(_dafny.Seq.of(_1043_v38, _1043_v38), _1043_v38)) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1044_i9)) {
              break L11;
            }
            _1044_i9 = (_1044_i9).plus(_dafny.ONE);
            _1042_v37 = _1042_v37;
            let _1045_v39;
            _1045_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_978_v0,_978_v0),(_978_v0).isLessThanOrEqualTo(_978_v0));
            let _1046_v40;
            _1046_v40 = _dafny.Map.Empty.slice().updateUnsafe(_978_v0,_978_v0);
            _1045_v39 = (_1045_v39).update(_1046_v40, _1042_v37);
            let _1047_v41;
            let _nw182 = Array((new BigNumber(16)).toNumber());
            _nw182[(_dafny.ZERO).toNumber()] = _1042_v37;
            _nw182[(_dafny.ONE).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(2)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(3)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(4)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(5)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(6)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(7)).toNumber()] = false;
            _nw182[(new BigNumber(8)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(9)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(10)).toNumber()] = false;
            _nw182[(new BigNumber(11)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(12)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(13)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(14)).toNumber()] = _1042_v37;
            _nw182[(new BigNumber(15)).toNumber()] = _1042_v37;
            _1047_v41 = _nw182;
            let _index206 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((p0).length));
            (p0)[_index206] = _1047_v41;
            let _index207 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((p0).length));
            (p0)[_index207] = _1047_v41;
            let _1048_v42;
            _1048_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1047_v41,!(!(!(_1042_v37) || (_1042_v37))));
            _1042_v37 = (((_1048_v42).contains(_1047_v41)) ? ((_1048_v42).get(_1047_v41)) : (_1042_v37));
          }
        }
      }
      let _1049_v43;
      let _nw183 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1049_v43 = _nw183;
      let _1050_v44;
      _1050_v44 = _dafny.Seq.UnicodeFromString("nhh");
      let _1051_v45;
      _1051_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1042_v37,_1050_v44);
      let _index208 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_1049_v43).length));
      (_1049_v43)[_index208] = (((_1051_v45).contains(_1042_v37)) ? ((_1051_v45).get(_1042_v37)) : (_1050_v44));
      let _1052_v47;
      _1052_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of _dafny.IntegerRange(new BigNumber(445), new BigNumber(767))) {
          let _1053_v46 = _compr_33;
          if (((new BigNumber(445)).isLessThanOrEqualTo(_1053_v46)) && ((_1053_v46).isLessThan(new BigNumber(767)))) {
            _coll33.push([(_1053_v46).plus(_978_v0),_1042_v37]);
          }
        }
        return _coll33;
      }()).length),_1042_v37);
      let _index209 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_1049_v43).length));
      (_1049_v43)[_index209] = _dafny.Seq.Concat(_1050_v44, _module.__default.fm22(_1042_v37, new BigNumber((_1052_v47).length), globalState));
      _1042_v37 = _1042_v37;
      _978_v0 = _978_v0;
      r0 = _1042_v37;
      let _1054_v48;
      let _nw184 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _1054_v48 = _nw184;
      let _1055_v49;
      _1055_v49 = _dafny.Seq.of(_1054_v48);
      let _1056_v50;
      _1056_v50 = _dafny.Seq.of((_1049_v43)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_1049_v43).length))]);
      let _1057_v51;
      _1057_v51 = _dafny.Seq.of((_1055_v49)[_module.__default.safeIndex(new BigNumber(((_1056_v50)[_module.__default.safeIndex(_978_v0, new BigNumber((_1056_v50).length))]).length), new BigNumber((_1055_v49).length))]);
      let _1058_v52;
      _1058_v52 = _dafny.MultiSet.fromElements(true);
      r1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1057_v51).length),_1058_v52);
      return [r0, r1];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = false;
      let _1059_v0;
      _1059_v0 = _dafny.Seq.UnicodeFromString("goxqvwwrh");
      _1059_v0 = _module.__default.fm4(!(p0) || (p0), p0, p0, globalState);
      let _1060_v1;
      _1060_v1 = new BigNumber(669);
      let _1061_v2;
      _1061_v2 = _dafny.Seq.of(p0);
      _1060_v1 = new BigNumber((_1061_v2).length);
      let _1062_v3;
      _1062_v3 = _dafny.MultiSet.fromElements(_1060_v1, new BigNumber(-45), _1060_v1);
      _1062_v3 = _1062_v3;
      let _1063_v4;
      _1063_v4 = _module.D1.create_DC3(_1060_v1);
      let _1064_v5;
      _1064_v5 = _module.D9.create_DC19(_1060_v1, _1060_v1);
      let _pat_let_tv19 = _1064_v5;
      let _source11 = function (_pat_let19_0) {
        return function (_1065_dt__update__tmp_h0) {
          return function (_pat_let20_0) {
            return function (_1066_dt__update_hcf5_h0) {
              return _module.D1.create_DC3(_1066_dt__update_hcf5_h0);
            }(_pat_let20_0);
          }((_pat_let_tv19).dtor_cf29);
        }(_pat_let19_0);
      }(_1063_v4);
      if (_source11.is_DC3) {
        let _1067___mcc_h0 = (_source11).cf5;
        let _1068_cf5 = _1067___mcc_h0;
        let _1069_v6;
        _1069_v6 = _dafny.Seq.of(_1060_v1, _1060_v1, _1068_cf5, _1068_cf5);
        let _1070_v7;
        _1070_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v1,_1069_v6);
        let _1071_v8;
        _1071_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v1,p0);
        let _1072_v9;
        _1072_v9 = _module.D10.create_DC21(_dafny.MultiSet.fromElements((((_1071_v8).contains(_1060_v1)) ? ((_1071_v8).get(_1060_v1)) : (false)), p0, p0));
        let _1073_v10;
        _1073_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1068_cf5);
        let _1074_v11;
        _1074_v11 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1075_v12;
        let _nw185 = Array((new BigNumber(12)).toNumber());
        _nw185[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((((_1070_v7).contains(new BigNumber(696))) ? ((_1070_v7).get(new BigNumber(696))) : (_1069_v6))).length));
        _nw185[(_dafny.ONE).toNumber()] = _1060_v1;
        _nw185[(new BigNumber(2)).toNumber()] = new BigNumber(((_1072_v9).dtor_cf34).cardinality());
        _nw185[(new BigNumber(3)).toNumber()] = _1068_cf5;
        _nw185[(new BigNumber(4)).toNumber()] = (_1068_cf5).minus(_1060_v1);
        _nw185[(new BigNumber(5)).toNumber()] = new BigNumber((_1069_v6).length);
        _nw185[(new BigNumber(6)).toNumber()] = (((_1073_v10).contains(p0)) ? ((_1073_v10).get(p0)) : (_1068_cf5));
        _nw185[(new BigNumber(7)).toNumber()] = _1060_v1;
        _nw185[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm11(_1074_v11, _1068_cf5, true, _module.__default.fm34(new BigNumber((_dafny.Seq.UnicodeFromString("qbdinbg")).length), _1068_cf5, globalState), globalState)).length));
        _nw185[(new BigNumber(9)).toNumber()] = _module.__default.fm24(new BigNumber(857), true, p0, _1060_v1, globalState);
        _nw185[(new BigNumber(10)).toNumber()] = new BigNumber((_1069_v6).length);
        _nw185[(new BigNumber(11)).toNumber()] = new BigNumber(321);
        _1075_v12 = _nw185;
        _1075_v12 = _1075_v12;
        let _1076_v13;
        let _nw186 = new _module.C1();
        _nw186.__ctor(_dafny.MultiSet.FromArray(_1061_v2), new BigNumber((_1071_v8).length));
        _1076_v13 = _nw186;
        let _1077_v14;
        _1077_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1076_v13);
        let _1078_v15;
        _1078_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1079_v16;
        _1079_v16 = _dafny.Seq.of(_1075_v12, _1075_v12);
        let _1080_v17;
        let _nw187 = Array((new BigNumber(19)).toNumber());
        _nw187[(_dafny.ZERO).toNumber()] = _1075_v12;
        _nw187[(_dafny.ONE).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(2)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(3)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(4)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(5)).toNumber()] = (_1079_v16)[_module.__default.safeIndex(_1068_cf5, new BigNumber((_1079_v16).length))];
        _nw187[(new BigNumber(6)).toNumber()] = (_1079_v16)[_module.__default.safeIndex(_1060_v1, new BigNumber((_1079_v16).length))];
        _nw187[(new BigNumber(7)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(8)).toNumber()] = (_1079_v16)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_1079_v16).length))];
        _nw187[(new BigNumber(9)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(10)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(11)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(12)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(13)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(14)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(15)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(16)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(17)).toNumber()] = _1075_v12;
        _nw187[(new BigNumber(18)).toNumber()] = _1075_v12;
        _1080_v17 = _nw187;
        let _1081_v18;
        _1081_v18 = _1073_v10;
        let _1082_v19;
        _1082_v19 = _dafny.Set.fromElements(p0);
        let _1083_v20;
        _1083_v20 = _module.D14.create_DC30(_1060_v1, _1080_v17, _1081_v18, new BigNumber((_1082_v19).length), _1075_v12);
        let _1084_v21;
        let _nw188 = Array((new BigNumber(11)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = (new BigNumber(((_1077_v14).update(!(p0), _1076_v13)).length)).isLessThan(_1068_cf5);
        _nw188[(_dafny.ONE).toNumber()] = ((p0) ? (!((((_1078_v15).contains((((_1078_v15).contains(p0)) ? ((_1078_v15).get(p0)) : (p0)))) ? ((_1078_v15).get((((_1078_v15).contains(p0)) ? ((_1078_v15).get(p0)) : (p0)))) : (p0)))) : (p0));
        _nw188[(new BigNumber(2)).toNumber()] = (_1060_v1).isLessThan(_1060_v1);
        _nw188[(new BigNumber(3)).toNumber()] = p0;
        _nw188[(new BigNumber(4)).toNumber()] = p0;
        _nw188[(new BigNumber(5)).toNumber()] = false;
        _nw188[(new BigNumber(6)).toNumber()] = !_dafny.Seq.contains(_1069_v6, (_1083_v20).dtor_cf47);
        _nw188[(new BigNumber(7)).toNumber()] = false;
        _nw188[(new BigNumber(8)).toNumber()] = p0;
        _nw188[(new BigNumber(9)).toNumber()] = p0;
        _nw188[(new BigNumber(10)).toNumber()] = false;
        _1084_v21 = _nw188;
        r1 = _1084_v21;
        _1069_v6 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), function (_1085_i0) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-426)), function (_1086_i1) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length));
        }), _dafny.Seq.Concat(_1069_v6, _1069_v6));
        let _1087_v22;
        let _nw189 = new _module.C1();
        _nw189.__ctor((_1076_v13).f10, _1076_v13.f11);
        _1087_v22 = _nw189;
      } else if (_source11.is_DC4) {
        let _1088___mcc_h1 = (_source11).cf6;
        let _1089___mcc_h2 = (_source11).cf7;
        let _1090___mcc_h3 = (_source11).cf8;
        let _1091___mcc_h4 = (_source11).cf9;
        let _1092_cf9 = _1091___mcc_h4;
        let _1093_cf8 = _1090___mcc_h3;
        let _1094_cf7 = _1089___mcc_h2;
        let _1095_cf6 = _1088___mcc_h1;
        r0 = ((_1093_cf8) ? (p0) : ((p0) === (_1093_cf8)));
        if (!((_1060_v1).isLessThan(_1092_cf9))) {
          let _1096_v23;
          let _nw190 = Array((new BigNumber(27)).toNumber()).fill([]);
          _1096_v23 = _nw190;
          let _1097_v24;
          let _nw191 = new _module.C7();
          _nw191.__ctor(_1096_v23, p0);
          _1097_v24 = _nw191;
          let _1098_v25;
          _1098_v25 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
          let _1099_v26;
          _1099_v26 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1060_v1), _1060_v1, new BigNumber((_dafny.Seq.of(new BigNumber((_1098_v25).length), _1060_v1)).length), _1060_v1);
          let _1100_v27;
          let _nw192 = Array((new BigNumber(14)).toNumber());
          _nw192[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(_1097_v24)).length);
          _nw192[(_dafny.ONE).toNumber()] = _1060_v1;
          _nw192[(new BigNumber(2)).toNumber()] = _1092_cf9;
          _nw192[(new BigNumber(3)).toNumber()] = _1095_cf6;
          _nw192[(new BigNumber(4)).toNumber()] = new BigNumber((_1099_v26).length);
          _nw192[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1062_v3).cardinality()));
          _nw192[(new BigNumber(6)).toNumber()] = _1060_v1;
          _nw192[(new BigNumber(7)).toNumber()] = _1095_cf6;
          _nw192[(new BigNumber(8)).toNumber()] = new BigNumber((_1061_v2).length);
          _nw192[(new BigNumber(9)).toNumber()] = (_1097_v24).fm6((_1097_v24).f3, _1095_cf6, _1060_v1, _1094_cf7, globalState);
          _nw192[(new BigNumber(10)).toNumber()] = new BigNumber((_1059_v0).length);
          _nw192[(new BigNumber(11)).toNumber()] = new BigNumber(210);
          _nw192[(new BigNumber(12)).toNumber()] = _1092_cf9;
          _nw192[(new BigNumber(13)).toNumber()] = new BigNumber(((_1098_v25).update((_1097_v24).f3, _1093_cf8)).length);
          _1100_v27 = _nw192;
          let _1101_v28;
          let _init17 = ((_1102_v25) => function (_1103_i2) {
            return _1102_v25;
          })(_1098_v25);
          let _nw193 = Array((new BigNumber(18)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw193.length); _i0_17++) {
            _nw193[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _1101_v28 = _nw193;
          let _1104_v29;
          _1104_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1100_v27,_1101_v28);
          _1104_v29 = (_1104_v29).update(_1100_v27, _1101_v28);
          let _1105_v30;
          let _init18 = ((_1106_v0) => function (_1107_i3) {
            return _dafny.Seq.Concat(_1106_v0, _1106_v0);
          })(_1059_v0);
          let _nw194 = Array((new BigNumber(11)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw194.length); _i0_18++) {
            _nw194[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _1105_v30 = _nw194;
          let _index210 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_1105_v30).length));
          (_1105_v30)[_index210] = _dafny.Seq.Concat(_1059_v0, _1059_v0);
          let _index211 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_1105_v30).length));
          (_1105_v30)[_index211] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), ((_1108_cf7) => function (_1109_i4) {
            return _1108_cf7;
          })(_1094_cf7)), _module.__default.fm22((_1097_v24).f3, _1095_cf6, globalState)), _module.__default.safeIndex(_1092_cf9, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), ((_1110_cf7) => function (_1111_i4) {
            return _1110_cf7;
          })(_1094_cf7)), _module.__default.fm22((_1097_v24).f3, _1095_cf6, globalState))).length)), new _dafny.CodePoint('w'.codePointAt(0)));
          let _1112_v31;
          let _nw195 = new _module.C5();
          _nw195.__ctor((_1105_v30)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_1105_v30).length))]);
          _1112_v31 = _nw195;
          _1059_v0 = _dafny.Seq.Concat((_1112_v31).f6, (_1105_v30)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_1105_v30).length))]);
          let _1113_v32;
          _1113_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1092_cf9).minus(_1060_v1),false);
          _1113_v32 = (_1113_v32).update(_1092_cf9, _1093_cf8);
        } else {
          let _1114_v33;
          let _nw196 = new _module.C4();
          _nw196.__ctor(_module.__default.safeModuloInt(new BigNumber((_1059_v0).length), _1060_v1));
          _1114_v33 = _nw196;
          _1114_v33 = _1114_v33;
          let _1115_v34;
          let _nw197 = new _module.C0();
          _nw197.__ctor();
          _1115_v34 = _nw197;
          let _1116_v35;
          let _nw198 = Array((new BigNumber(14)).toNumber()).fill(false);
          _1116_v35 = _nw198;
          let _1117_v36;
          _1117_v36 = _dafny.Set.fromElements(_1116_v35, _1116_v35);
          let _rhs154 = _1093_cf8;
          let _rhs155 = _1117_v36;
          let _rhs156 = _1114_v33.f7;
          let _rhs157 = p0;
          let _rhs158 = p0;
          r0 = _rhs154;
          _1117_v36 = _rhs155;
          _1095_cf6 = _rhs156;
          _1093_cf8 = _rhs157;
          r0 = _rhs158;
          let _1118_v37;
          let _nw199 = new _module.C2();
          _nw199.__ctor(_dafny.Seq.UnicodeFromString("sm"), _dafny.Seq.update(_dafny.Seq.update(_1059_v0, _module.__default.safeIndex(_1114_v33.f7, new BigNumber((_1059_v0).length)), _1094_cf7), _module.__default.safeIndex(_1060_v1, new BigNumber((_dafny.Seq.update(_1059_v0, _module.__default.safeIndex(_1114_v33.f7, new BigNumber((_1059_v0).length)), _1094_cf7)).length)), _1094_cf7));
          _1118_v37 = _nw199;
          let _1119_v38;
          _1119_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v1,_1118_v37);
          _1119_v38 = (_1119_v38).update(_1060_v1, _1118_v37);
          let _1120_v39;
          _1120_v39 = _dafny.Seq.of(_1116_v35);
          let _1121_v40;
          _1121_v40 = _dafny.MultiSet.fromElements(_1116_v35, ((!(_1093_cf8)) ? (_1116_v35) : ((_1120_v39)[_module.__default.safeIndex(_1060_v1, new BigNumber((_1120_v39).length))])), _1116_v35, _1116_v35, _1116_v35);
          _1092_cf9 = (((_1121_v40).contains(_1116_v35)) ? ((_1121_v40).get(_1116_v35)) : (_module.__default.safeModuloInt(_1060_v1, new BigNumber(382))));
        }
        let _1122_v41;
        let _nw200 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _1122_v41 = _nw200;
        _1122_v41 = _1122_v41;
        let _1123_v42;
        let _init19 = ((_1124_cf9, _1125_cf8, _1126_cf6) => function (_1127_i5) {
          return (_dafny.Set.fromElements(_dafny.Seq.of(_1124_cf9, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_1125_cf8))).cardinality()), _1126_cf6))).Difference(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(174)), function (_1128_i6) {
            return new BigNumber((_dafny.Seq.of(true)).length);
          })));
        })(_1092_cf9, _1093_cf8, _1095_cf6);
        let _nw201 = Array((new BigNumber(25)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw201.length); _i0_19++) {
          _nw201[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _1123_v42 = _nw201;
        let _1129_v43;
        _1129_v43 = _dafny.Seq.of(_1092_cf9);
        let _1130_v44;
        _1130_v44 = _dafny.Set.fromElements(_1129_v43);
        let _index212 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_1123_v42).length));
        (_1123_v42)[_index212] = _1130_v44;
        let _index213 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_1123_v42).length));
        (_1123_v42)[_index213] = function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of ((_dafny.Map.Empty.slice().updateUnsafe(_1129_v43,_1060_v1)).update(_dafny.Seq.of(_1095_cf6, _1092_cf9), new BigNumber(437))).Keys.Elements) {
            let _1131_v45 = _compr_34;
            if (((_dafny.Map.Empty.slice().updateUnsafe(_1129_v43,_1060_v1)).update(_dafny.Seq.of(_1095_cf6, _1092_cf9), new BigNumber(437))).contains(_1131_v45)) {
              _coll34.add(_1131_v45);
            }
          }
          return _coll34;
        }();
      } else {
        let _1132___mcc_h5 = (_source11).cf4;
        let _1133_cf4 = _1132___mcc_h5;
        let _1134_v46;
        let _nw202 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _1134_v46 = _nw202;
        let _1135_v47;
        _1135_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1134_v46);
        let _1136_v48;
        _1136_v48 = _module.D4.create_DC9(_1134_v46);
        let _1137_v49;
        _1137_v49 = _module.D4.create_DC11(_1136_v48);
        let _1138_v50;
        _1138_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1135_v47).length),_1137_v49);
        _1138_v50 = (_1138_v50).update((_dafny.ZERO).minus(_1060_v1), _1137_v49);
        let _1139_v51;
        _1139_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1134_v46,false);
        let _1140_v52;
        _1140_v52 = _module.D9.create_DC20(p0, p0, _1133_cf4);
        _1139_v51 = (_1139_v51).update(_1134_v46, (_1140_v52).dtor_cf33);
        let _index214 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length));
        (_1134_v46)[_index214] = _1060_v1;
        let _1141_v53;
        _1141_v53 = new _dafny.CodePoint('x'.codePointAt(0));
        let _1142_v54;
        _1142_v54 = _dafny.Set.fromElements(_1141_v53);
        let _1143_v55;
        _1143_v55 = _dafny.Seq.of(_1142_v54, _dafny.Set.fromElements(_1141_v53), _1142_v54);
        let _1144_v56;
        _1144_v56 = _dafny.Seq.of((_dafny.ZERO).minus(_1060_v1), _1060_v1, (_1063_v4).dtor_cf5, _1060_v1, new BigNumber((_1143_v55).length));
        let _1145_v57;
        let _nw203 = Array((new BigNumber(18)).toNumber()).fill(false);
        _1145_v57 = _nw203;
        let _1146_v58;
        _1146_v58 = _dafny.Seq.of(_1145_v57);
        let _1147_v59;
        _1147_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v1,_1060_v1);
        let _1148_v60;
        _1148_v60 = _dafny.MultiSet.fromElements(_module.__default.fm8(_1062_v3, globalState), false);
        let _1149_v61;
        let _nw204 = new _module.C6();
        _nw204.__ctor(_1148_v60, _1134_v46);
        _1149_v61 = _nw204;
        let _1150_v62;
        _1150_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1149_v61,_1145_v57);
        let _1151_v63;
        _1151_v63 = _dafny.Set.fromElements(_1145_v57, _1145_v57);
        let _index215 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length));
        let _rhs159 = _1060_v1;
        let _rhs160 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1060_v1, (_1060_v1).plus(_1060_v1)));
        let _rhs161 = ((_1133_cf4) ? (_1144_v56) : (_1144_v56));
        let _rhs162 = _1060_v1;
        let _rhs163 = (_1151_v63).IsProperSubsetOf(_dafny.Set.fromElements(_1145_v57, _1145_v57, (_1146_v58)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1060_v1), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1147_v59)).length)), new BigNumber((_dafny.Seq.of(_1060_v1)).length)), _1060_v1)).length), new BigNumber((_1146_v58).length))], (((_1150_v62).contains(_1149_v61)) ? ((_1150_v62).get(_1149_v61)) : (_1145_v57))));
        let _lhs104 = _1134_v46;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length));
        _1060_v1 = _rhs159;
        _lhs104[_lhs105] = _rhs160;
        _1144_v56 = _rhs161;
        _1060_v1 = _rhs162;
        r0 = _rhs163;
        let _1152_v64;
        _1152_v64 = _dafny.Map.Empty.slice().updateUnsafe((_1134_v46)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length))],_1061_v2);
        let _1153_v65;
        let _nw205 = new _module.C0();
        _nw205.__ctor();
        _1153_v65 = _nw205;
        let _1154_v66;
        _1154_v66 = _module.D7.create_DC15(true, new BigNumber(-529), _1060_v1, (_1134_v46)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length))], _1153_v65);
        let _1155_v67;
        _1155_v67 = _dafny.MultiSet.fromElements(_1154_v66, _1154_v66);
        let _1156_v68;
        _1156_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1155_v67,_1059_v0);
        let _1157_v69;
        _1157_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_1152_v64).contains((_1134_v46)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length))])) ? ((_1152_v64).get((_1134_v46)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length))])) : (_dafny.Seq.update(_dafny.Seq.of(_1133_cf4), _module.__default.safeIndex(new BigNumber((_1144_v56).length), new BigNumber((_dafny.Seq.of(_1133_cf4)).length)), p0)))).length),(new BigNumber((_1156_v68).length)).isLessThanOrEqualTo(_1060_v1));
        let _1158_v70;
        let _nw206 = new _module.C1();
        _nw206.__ctor(_1148_v60, (_1134_v46)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1134_v46).length))]);
        _1158_v70 = _nw206;
        let _1159_v71;
        _1159_v71 = _dafny.MultiSet.fromElements(_1158_v70);
        let _1160_v72;
        _1160_v72 = _module.D2.create_DC6(_1060_v1, _1133_cf4);
        let _1161_v73;
        _1161_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1149_v61,false);
        _1157_v69 = (_1157_v69).update((new BigNumber((_1159_v71).cardinality())).plus(new BigNumber((_1062_v3).cardinality())), (_1149_v61).fm10(_1160_v72, new BigNumber((_1059_v0).length), (_this).fm1(_1060_v1, new BigNumber((_dafny.Seq.of(new BigNumber((_1161_v73).length))).length), _1144_v56, globalState), globalState));
      }
      let _1162_v74;
      _1162_v74 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_1059_v0, _module.__default.safeIndex(new BigNumber((_1062_v3).cardinality()), new BigNumber((_1059_v0).length)), _module.__default.fm30(_1059_v0, _1060_v1, _1060_v1, globalState))).length), _1060_v1);
      let _hi9 = _1060_v1;
      for (let _1163_i7 = new BigNumber((_1162_v74).length); _1163_i7.isLessThan(_hi9); _1163_i7 = _1163_i7.plus(_dafny.ONE)) {
        let _1164_v75;
        _1164_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1163_i7,new BigNumber(223));
        _1164_v75 = (_1164_v75).update(new BigNumber(-742), _1060_v1);
        let _1165_v76;
        _1165_v76 = _dafny.Set.fromElements(_1060_v1, _1163_i7, new BigNumber((_1059_v0).length), _1163_i7, new BigNumber(924));
        let _1166_v77;
        _1166_v77 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_1165_v76) : (_1165_v76)),_1162_v74);
        let _1167_v79;
        _1167_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1162_v74);
        _1166_v77 = (_1166_v77).update(((p0) ? (_1165_v76) : (function () {
          let _coll35 = new _dafny.Set();
          for (const _compr_35 of _dafny.IntegerRange(new BigNumber(452), new BigNumber(-33))) {
            let _1168_v78 = _compr_35;
            if (((new BigNumber(452)).isLessThanOrEqualTo(_1168_v78)) && ((_1168_v78).isLessThan(new BigNumber(-33)))) {
              _coll35.add((_1168_v78).minus(new BigNumber(265)));
            }
          }
          return _coll35;
        }())), (((_1167_v79).contains(p0)) ? ((_1167_v79).get(p0)) : (_1162_v74)));
        _1060_v1 = _1060_v1;
        let _1169_v80;
        let _nw207 = Array((new BigNumber(25)).toNumber()).fill(false);
        _1169_v80 = _nw207;
        let _1170_v81;
        _1170_v81 = _dafny.Seq.of(_1169_v80);
        r2 = !_dafny.areEqual(_1170_v81, _1170_v81);
      }
      r0 = _module.__default.fm8(_1062_v3, globalState);
      let _1171_v82;
      _1171_v82 = _module.D4.create_DC10(p0);
      r0 = (_1171_v82).dtor_cf16;
      let _1172_v83;
      let _nw208 = Array((new BigNumber(9)).toNumber()).fill(false);
      _1172_v83 = _nw208;
      r1 = _1172_v83;
      let _1173_v84;
      _1173_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v1,_1162_v74);
      r2 = _module.__default.fm8(_dafny.MultiSet.FromArray((((_1173_v84).contains(_1060_v1)) ? ((_1173_v84).get(_1060_v1)) : (_1162_v74))), globalState);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
