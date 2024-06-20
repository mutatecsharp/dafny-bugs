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
      let _source0 = _module.D8.create_DC23();
      if (_source0.is_DC23) {
        return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), function (_0_i0) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        })).length)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(25))).length));
      } else if (_source0.is_DC24) {
        let _1___mcc_h0 = (_source0).cf45;
        let _2_cf45 = _1___mcc_h0;
        return true;
      } else if (_source0.is_DC22) {
        let _3___mcc_h1 = (_source0).cf44;
        let _4_cf44 = _3___mcc_h1;
        return _dafny.Seq.contains(_dafny.Seq.of(true, false), false);
      } else {
        let _5___mcc_h2 = (_source0).cf46;
        let _6_cf46 = _5___mcc_h2;
        return true;
      }
    };
    static fm1(globalState) {
      return _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(112),!(true))).length), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-891), new BigNumber(290))));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _module.__default.safeDivisionInt((new BigNumber(636)).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-371), new BigNumber((_dafny.Seq.UnicodeFromString("qywhj")).length), new BigNumber((_dafny.Seq.of(new BigNumber(616))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(929)), function (_7_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length))).length)), new BigNumber(503));
    };
    static fm3(p0, globalState) {
      return _module.D0.create_DC1();
    };
    static fm4(p0, p1, globalState) {
      if (_dafny.areEqual(_dafny.Seq.of(false, true, false, true, false), _dafny.Seq.of(false))) {
        return _module.D1.create_DC2(_dafny.Seq.UnicodeFromString("dudmm"));
      } else {
        return _module.D1.create_DC2(_dafny.Seq.UnicodeFromString("o"));
      }
    };
    static fm5(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(868), new BigNumber(482)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length), (_dafny.ZERO).minus(new BigNumber(-324))), _dafny.Seq.of(new BigNumber(822))));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("q"), _dafny.Seq.UnicodeFromString("rtqhqq")));
    };
    static fm7(p0, p1, globalState) {
      let _source1 = _module.D7.create_DC21(_module.D7.create_DC20(true, !(false), new BigNumber(-30), true));
      if (_source1.is_DC20) {
        let _8___mcc_h0 = (_source1).cf39;
        let _9___mcc_h1 = (_source1).cf40;
        let _10___mcc_h2 = (_source1).cf41;
        let _11___mcc_h3 = (_source1).cf42;
        let _12_cf42 = _11___mcc_h3;
        let _13_cf41 = _10___mcc_h2;
        let _14_cf40 = _9___mcc_h1;
        let _15_cf39 = _8___mcc_h0;
        return new _dafny.CodePoint('b'.codePointAt(0));
      } else if (_source1.is_DC19) {
        let _16___mcc_h4 = (_source1).cf38;
        let _17_cf38 = _16___mcc_h4;
        return new _dafny.CodePoint('w'.codePointAt(0));
      } else {
        let _18___mcc_h5 = (_source1).cf43;
        let _19_cf43 = _18___mcc_h5;
        return new _dafny.CodePoint('k'.codePointAt(0));
      }
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(702)));
    };
    static fm9(p0, globalState) {
      return ((_dafny.Set.fromElements(false, true)).Intersect(_dafny.Set.fromElements(!(!(!(true)))))).Union((_dafny.Set.fromElements(false, false)).Union(_dafny.Set.fromElements(true, true)));
    };
    static fm10(p0, p1, p2, globalState) {
      let _source2 = _module.D1.create_DC6(_module.D1.create_DC3(new BigNumber((_dafny.Seq.of(true, !(true))).length), new BigNumber(-624)));
      if (_source2.is_DC3) {
        let _20___mcc_h0 = (_source2).cf3;
        let _21___mcc_h1 = (_source2).cf4;
        let _22_cf4 = _21___mcc_h1;
        let _23_cf3 = _20___mcc_h0;
        if (!(false)) {
          return _module.D1.create_DC6(_module.D1.create_DC5(_22_cf4, true, _22_cf4, true));
        } else {
          return _module.D1.create_DC6(_module.D1.create_DC3(_22_cf4, _23_cf3));
        }
      } else if (_source2.is_DC4) {
        let _24___mcc_h2 = (_source2).cf5;
        let _25___mcc_h3 = (_source2).cf6;
        let _26_cf6 = _25___mcc_h3;
        let _27_cf5 = _24___mcc_h2;
        return _module.D1.create_DC6(_module.D1.create_DC6(_module.D1.create_DC5(_26_cf6, _27_cf5, _26_cf6, _27_cf5)));
      } else if (_source2.is_DC5) {
        let _28___mcc_h4 = (_source2).cf7;
        let _29___mcc_h5 = (_source2).cf8;
        let _30___mcc_h6 = (_source2).cf9;
        let _31___mcc_h7 = (_source2).cf10;
        let _32_cf10 = _31___mcc_h7;
        let _33_cf9 = _30___mcc_h6;
        let _34_cf8 = _29___mcc_h5;
        let _35_cf7 = _28___mcc_h4;
        return _module.D1.create_DC6(_module.D1.create_DC6(_module.D1.create_DC6(_module.D1.create_DC5((_dafny.ZERO).minus(_35_cf7), _32_cf10, new BigNumber((_dafny.Seq.UnicodeFromString("xexlenl")).length), _34_cf8))));
      } else if (_source2.is_DC2) {
        let _36___mcc_h8 = (_source2).cf2;
        let _37_cf2 = _36___mcc_h8;
        return _module.D1.create_DC6(_module.D1.create_DC6(_module.D1.create_DC2(_37_cf2)));
      } else {
        let _38___mcc_h9 = (_source2).cf11;
        let _39_cf11 = _38___mcc_h9;
        if (false) {
          return _module.D1.create_DC6(_module.D1.create_DC2(_dafny.Seq.UnicodeFromString("bno")));
        } else {
          return _module.D1.create_DC6(_module.D1.create_DC4(false, new BigNumber(755)));
        }
      }
    };
    static fm11(p0, p1, p2, globalState) {
      return _module.D1.create_DC4(((true) ? (true) : (false)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(!(true))), _dafny.Seq.of(true, false, !(false), false))).length));
    };
    static fm12(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.UnicodeFromString("dm")) : (_dafny.Seq.UnicodeFromString("yx"))), _dafny.Seq.UnicodeFromString("badgldb"));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(22), new BigNumber(-190), new BigNumber(-212))).Union((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)))).length), new BigNumber(-603))).length))).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(297),true)).length)), (_dafny.ZERO).minus(new BigNumber(-80)), new BigNumber(377), new BigNumber(-598))));
    };
    static fm14(p0, globalState) {
      let _source3 = _module.D8.create_DC23();
      if (_source3.is_DC23) {
        return _dafny.MultiSet.fromElements(true);
      } else if (_source3.is_DC24) {
        let _40___mcc_h0 = (_source3).cf45;
        let _41_cf45 = _40___mcc_h0;
        return _dafny.MultiSet.fromElements(_41_cf45, false, _41_cf45, _41_cf45);
      } else if (_source3.is_DC22) {
        let _42___mcc_h1 = (_source3).cf44;
        let _43_cf44 = _42___mcc_h1;
        return _dafny.MultiSet.fromElements((_module.D1.create_DC4(true, new BigNumber(371))).dtor_cf5, false);
      } else {
        let _44___mcc_h2 = (_source3).cf46;
        let _45_cf46 = _44___mcc_h2;
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).Difference(_dafny.MultiSet.fromElements(false));
      }
    };
    static m0(p0, globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _46_v0;
      _46_v0 = new _dafny.CodePoint('r'.codePointAt(0));
      let _47_v1;
      _47_v1 = _module.D2.create_DC10(new BigNumber(-185), _46_v0);
      let _source4 = _47_v1;
      if (_source4.is_DC8) {
        let _48___mcc_h0 = (_source4).cf13;
        let _49_cf13 = _48___mcc_h0;
        let _50_v2;
        _50_v2 = true;
        let _51_v3;
        _51_v3 = _dafny.Seq.of(_50_v2);
        let _52_v4;
        let _nw0 = new _module.C0();
        _nw0.__ctor(_dafny.Seq.IsProperPrefixOf(_51_v3, _51_v3));
        _52_v4 = _nw0;
        _52_v4 = _52_v4;
        let _53_v5;
        _53_v5 = _dafny.Seq.of(_dafny.Seq.Concat(_51_v3, _dafny.Seq.of(_50_v2)));
        let _54_v6;
        _54_v6 = new BigNumber(978);
        _53_v5 = _dafny.Seq.Concat(_dafny.Seq.update(_53_v5, _module.__default.safeIndex(_module.__default.fm2(_dafny.MultiSet.fromElements(_46_v0, new _dafny.CodePoint('e'.codePointAt(0))), _54_v6, _54_v6, false, globalState), new BigNumber((_53_v5).length)), _dafny.Seq.of((_52_v4).f29, _50_v2)), _53_v5);
        _52_v4 = _52_v4;
        (globalState).f27 = _54_v6;
      } else if (_source4.is_DC9) {
        let _55___mcc_h1 = (_source4).cf14;
        let _56___mcc_h2 = (_source4).cf15;
        let _57_cf15 = _56___mcc_h2;
        let _58_cf14 = _55___mcc_h1;
        let _59_v8;
        let _init0 = ((_60_cf15, _61_cf14) => function (_62_i0) {
          return (_62_i0).minus(new BigNumber((function () {
            let _coll0 = new _dafny.Map();
            for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(660),true)).Keys.Elements) {
              let _63_v7 = _compr_0;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(660),true)).contains(_63_v7)) {
                _coll0.push([_module.__default.safeDivisionInt(_63_v7, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(284), new BigNumber(-907))).length)),(_60_cf15).f29)).length))).length)),_61_cf14]);
              }
            }
            return _coll0;
          }()).length));
        })(_57_cf15, _58_cf14);
        let _nw1 = Array((new BigNumber(19)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _59_v8 = _nw1;
        let _64_v9;
        _64_v9 = _module.D4.create_DC12(_59_v8);
        (globalState).f28 = (_64_v9).dtor_cf19;
        if ((_57_cf15).f29) {
          let _65_v10;
          _65_v10 = new BigNumber(581);
          let _66_v11;
          _66_v11 = _dafny.Seq.of(_65_v10);
          let _67_v12;
          let _nw2 = Array((new BigNumber(5)).toNumber()).fill(false);
          _67_v12 = _nw2;
          let _68_v13;
          _68_v13 = _dafny.Map.Empty.slice().updateUnsafe(_67_v12,_57_cf15);
          _57_cf15 = ((!(((_66_v11)[_module.__default.safeIndex(new BigNumber((_68_v13).length), new BigNumber((_66_v11).length))]).isLessThanOrEqualTo(_65_v10))) ? (_57_cf15) : (_57_cf15));
          (globalState).f13 = (_66_v11)[_module.__default.safeIndex(_65_v10, new BigNumber((_66_v11).length))];
          let _69_v14;
          let _nw3 = Array((new BigNumber(5)).toNumber()).fill(_module.D2.Default());
          _69_v14 = _nw3;
          let _index0 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_69_v14).length));
          (_69_v14)[_index0] = _47_v1;
          let _index1 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_69_v14).length));
          (_69_v14)[_index1] = _47_v1;
          let _70_v15;
          _70_v15 = _dafny.Map.Empty.slice().updateUnsafe((_57_cf15).f29,(_57_cf15).f29);
          _70_v15 = (_70_v15).update(_module.__default.fm0(p0, _module.__default.fm0(p0, false, globalState), globalState), _58_cf14);
          (globalState).f13 = new BigNumber(249);
        } else {
          let _71_v16;
          let _nw4 = new _module.C0();
          _nw4.__ctor(!((_57_cf15).f29));
          _71_v16 = _nw4;
          (globalState).f14 = ((false) ? (true) : (true));
          let _72_v17;
          _72_v17 = _dafny.Set.fromElements(_58_cf14, (_57_cf15).f29);
          (globalState).f13 = new BigNumber((_72_v17).length);
          (globalState).f14 = (_57_cf15).f29;
          let _73_v18;
          _73_v18 = new BigNumber(927);
          let _74_v19;
          _74_v19 = _dafny.Seq.of(_73_v18, _73_v18, _73_v18);
          let _75_v20;
          _75_v20 = _dafny.Set.fromElements(_74_v19, _74_v19);
          (globalState).f14 = (new BigNumber((_75_v20).length)).isEqualTo(_73_v18);
        }
        let _76_v21;
        _76_v21 = new BigNumber(82);
        let _77_v22;
        _77_v22 = _dafny.Map.Empty.slice().updateUnsafe(_76_v21,(((_57_cf15).f29) ? ((_dafny.ZERO).minus(_76_v21)) : (_76_v21)));
        let _78_v23;
        _78_v23 = _dafny.Seq.of(_76_v21);
        _77_v22 = (_77_v22).update(_76_v21, (_76_v21).minus(new BigNumber((_78_v23).length)));
        let _79_v24;
        let _nw5 = new _module.C0();
        _nw5.__ctor(_58_cf14);
        _79_v24 = _nw5;
      } else if (_source4.is_DC10) {
        let _80___mcc_h3 = (_source4).cf16;
        let _81___mcc_h4 = (_source4).cf17;
        let _82_cf17 = _81___mcc_h4;
        let _83_cf16 = _80___mcc_h3;
        let _84_v25;
        _84_v25 = false;
        let _85_v26;
        let _init1 = ((_86_v25) => function (_87_i1) {
          return _86_v25;
        })(_84_v25);
        let _nw6 = Array((new BigNumber(13)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
          _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _85_v26 = _nw6;
        let _88_v27;
        _88_v27 = _dafny.Seq.of(_84_v25, _84_v25);
        let _89_v28;
        _89_v28 = _dafny.Seq.of(_88_v27);
        let _90_v29;
        _90_v29 = _dafny.Map.Empty.slice().updateUnsafe(_89_v28,_85_v26);
        let _91_v30;
        _91_v30 = _module.D5.create_DC15(p0, _84_v25, _84_v25, _84_v25, _85_v26);
        let _92_v31;
        let _nw7 = Array((new BigNumber(22)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = ((_84_v25) ? (_85_v26) : (_85_v26));
        _nw7[(_dafny.ONE).toNumber()] = _85_v26;
        _nw7[(new BigNumber(2)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(3)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(4)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(5)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(6)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(7)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(8)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(9)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(10)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(11)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(12)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(13)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(14)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(15)).toNumber()] = (((_90_v29).contains(_89_v28)) ? ((_90_v29).get(_89_v28)) : (_85_v26));
        _nw7[(new BigNumber(16)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(17)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(18)).toNumber()] = (_91_v30).dtor_cf29;
        _nw7[(new BigNumber(19)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(20)).toNumber()] = _85_v26;
        _nw7[(new BigNumber(21)).toNumber()] = _85_v26;
        _92_v31 = _nw7;
        let _93_v32;
        _93_v32 = _module.D0.create_DC0(_83_cf16, _83_cf16);
        _92_v31 = (((new BigNumber(-176)).isLessThan((function (_pat_let0_0) {
          return function (_94_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_95_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_95_dt__update_hcf0_h0, (_94_dt__update__tmp_h0).dtor_cf1);
              }(_pat_let1_0);
            }(new BigNumber((_dafny.Seq.UnicodeFromString("tfraklcw")).length));
          }(_pat_let0_0);
        }(_93_v32)).dtor_cf0)) ? (_92_v31) : (_92_v31));
        (globalState).f14 = _84_v25;
        let _96_v33;
        _96_v33 = _dafny.Seq.of(_83_cf16);
        (globalState).f27 = (_96_v33)[_module.__default.safeIndex(_83_cf16, new BigNumber((_96_v33).length))];
        let _97_v34;
        let _nw8 = new _module.C0();
        _nw8.__ctor(_84_v25);
        _97_v34 = _nw8;
        let _98_v35;
        _98_v35 = _dafny.Map.Empty.slice().updateUnsafe(!(_83_cf16).isEqualTo(_83_cf16),_97_v34);
        let _rhs0 = (_98_v35).update((_97_v34).f29, _97_v34);
        let _rhs1 = _module.__default.fm2(_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), _82_cf17, _82_cf17, _module.__default.fm7(p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(_83_cf16)), globalState)), (((_97_v34).f29) ? (_83_cf16) : (_83_cf16)), (_dafny.ZERO).minus(_83_cf16), true, globalState);
        let _rhs2 = (_97_v34).f29;
        let _lhs0 = globalState;
        _98_v35 = _rhs0;
        _83_cf16 = _rhs1;
        _lhs0.f14 = _rhs2;
      } else {
        let _99___mcc_h5 = (_source4).cf12;
        let _100_cf12 = _99___mcc_h5;
        let _101_v36;
        _101_v36 = true;
        let _102_v37;
        let _nw9 = new _module.C0();
        _nw9.__ctor(((_101_v36) ? (_101_v36) : (_101_v36)));
        _102_v37 = _nw9;
        let _103_v38;
        _103_v38 = _dafny.MultiSet.fromElements((_102_v37).f29);
        let _104_v39;
        let _nw10 = new _module.C0();
        _nw10.__ctor((_103_v38).contains(_101_v36));
        _104_v39 = _nw10;
        let _105_v40;
        let _nw11 = new _module.C0();
        _nw11.__ctor((_102_v37).f29);
        _105_v40 = _nw11;
        let _106_v41;
        _106_v41 = new BigNumber(556);
        let _107_v42;
        _107_v42 = _module.D1.create_DC4((_105_v40).f29, _106_v41);
        _101_v36 = (_107_v42).dtor_cf5;
      }
      let _108_v43;
      _108_v43 = true;
      let _109_i2;
      _109_i2 = _dafny.ZERO;
      L0: {
        while (!(_108_v43)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_109_i2)) {
              break L0;
            }
            _109_i2 = (_109_i2).plus(_dafny.ONE);
            (globalState).f14 = _108_v43;
            (globalState).f14 = ((!(_108_v43)) ? (((_108_v43) ? (_108_v43) : (_108_v43))) : (_108_v43));
            let _110_v44;
            let _nw12 = new _module.C0();
            _nw12.__ctor(false);
            _110_v44 = _nw12;
            let _111_v45;
            _111_v45 = _dafny.MultiSet.fromElements(false);
            let _112_v46;
            _112_v46 = _dafny.Seq.of((_110_v44).f29);
            let _113_v47;
            let _nw13 = Array((new BigNumber(25)).toNumber());
            _nw13[(_dafny.ZERO).toNumber()] = _111_v45;
            _nw13[(_dafny.ONE).toNumber()] = _111_v45;
            _nw13[(new BigNumber(2)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements((_110_v44).f29);
            _nw13[(new BigNumber(4)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(5)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(6)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(7)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(8)).toNumber()] = (_module.D7.create_DC19(_dafny.MultiSet.FromArray(_112_v46))).dtor_cf38;
            _nw13[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_module.__default.fm0(p0, _108_v43, globalState));
            _nw13[(new BigNumber(10)).toNumber()] = (_111_v45).update((_110_v44).f29, _module.__default.abs(new BigNumber(518)));
            _nw13[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_112_v46);
            _nw13[(new BigNumber(12)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(13)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(14)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(15)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(!(_108_v43), (_110_v44).f29, !(_108_v43), _108_v43);
            _nw13[(new BigNumber(17)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(false));
            _nw13[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements((_110_v44).f29, _108_v43);
            _nw13[(new BigNumber(20)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(21)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(22)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(23)).toNumber()] = _111_v45;
            _nw13[(new BigNumber(24)).toNumber()] = _111_v45;
            _113_v47 = _nw13;
            let _114_v48;
            _114_v48 = new BigNumber(2);
            let _115_v49;
            _115_v49 = _dafny.Seq.of(_114_v48, new BigNumber(-565), _114_v48);
            let _116_v50;
            _116_v50 = _dafny.Seq.of(_115_v49);
            let _117_v51;
            _117_v51 = _dafny.Map.Empty.slice().updateUnsafe(_113_v47,(_116_v50)[_module.__default.safeIndex(_114_v48, new BigNumber((_116_v50).length))]);
            _117_v51 = (_117_v51).update(_113_v47, _115_v49);
          }
        }
      }
      let _118_v52;
      _118_v52 = new BigNumber(366);
      let _119_v53;
      _119_v53 = _dafny.Set.fromElements(_118_v52);
      let _120_v54;
      _120_v54 = _dafny.Set.fromElements(new BigNumber((_119_v53).length));
      (globalState).f14 = ((_119_v53).equals(_119_v53)) === (_108_v43);
      let _121_v55;
      _121_v55 = _dafny.Map.Empty.slice().updateUnsafe(_108_v43,_46_v0);
      _121_v55 = (_121_v55).update(!(_108_v43), _46_v0);
      let _122_i3;
      _122_i3 = _dafny.ZERO;
      L1: {
        while (_108_v43) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_122_i3)) {
              break L1;
            }
            _122_i3 = (_122_i3).plus(_dafny.ONE);
            if (_108_v43) {
              let _123_v56;
              _123_v56 = _dafny.Map.Empty.slice().updateUnsafe(_108_v43,new BigNumber((_module.__default.fm9(_118_v52, globalState)).length));
              let _124_v57;
              let _nw14 = new _module.C0();
              _nw14.__ctor(_108_v43);
              _124_v57 = _nw14;
              let _125_v58;
              _125_v58 = _dafny.Set.fromElements(_124_v57, _124_v57, _124_v57);
              _123_v56 = (_123_v56).update(((_108_v43) ? (false) : (_108_v43)), new BigNumber((_125_v58).length));
              (globalState).f14 = (_124_v57).f29;
              let _126_v59;
              _126_v59 = _dafny.Set.fromElements((_124_v57).f29, _108_v43);
              let _127_v60;
              _127_v60 = _dafny.Seq.of(_118_v52);
              let _128_v61;
              let _nw15 = new _module.C0();
              _nw15.__ctor(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber((_module.__default.fm13((_124_v57).f29, _126_v59, _127_v60, _118_v52, globalState)).length), _118_v52), _127_v60));
              _128_v61 = _nw15;
              let _129_v62;
              _129_v62 = _dafny.Seq.of(_108_v43, (_128_v61).f29);
              let _130_v63;
              _130_v63 = _module.D4.create_DC13(_128_v61, new _dafny.CodePoint('w'.codePointAt(0)), _46_v0, _129_v62);
              let _131_v64;
              _131_v64 = _dafny.Map.Empty.slice().updateUnsafe(_118_v52,_128_v61);
              let _132_v65;
              let _nw16 = Array((new BigNumber(8)).toNumber());
              _nw16[(_dafny.ZERO).toNumber()] = (_130_v63).dtor_cf20;
              _nw16[(_dafny.ONE).toNumber()] = _128_v61;
              _nw16[(new BigNumber(2)).toNumber()] = _124_v57;
              _nw16[(new BigNumber(3)).toNumber()] = _128_v61;
              _nw16[(new BigNumber(4)).toNumber()] = _124_v57;
              _nw16[(new BigNumber(5)).toNumber()] = _124_v57;
              _nw16[(new BigNumber(6)).toNumber()] = _124_v57;
              _nw16[(new BigNumber(7)).toNumber()] = (((_128_v61).f29) ? (_124_v57) : ((((_131_v64).contains(_118_v52)) ? ((_131_v64).get(_118_v52)) : (_128_v61))));
              _132_v65 = _nw16;
              let _index2 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_132_v65).length));
              (_132_v65)[_index2] = _124_v57;
              let _index3 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_132_v65).length));
              (_132_v65)[_index3] = _124_v57;
              r2 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(932)), ((_133_v0) => function (_134_i4) {
                return _133_v0;
              })(_46_v0))).length)).multipliedBy(_118_v52);
            } else {
              let _135_v66;
              let _init2 = ((_136_v0, _137_v1) => function (_138_i5) {
                return _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(383)), ((_139_v0) => function (_140_i6) {
                  return _139_v0;
                })(_136_v0)), _module.__default.safeIndex((_137_v1).dtor_cf16, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(383)), ((_141_v0) => function (_142_i6) {
                  return _141_v0;
                })(_136_v0))).length)), _136_v0);
              })(_46_v0, _47_v1);
              let _nw17 = Array((new BigNumber(12)).toNumber());
              for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw17.length); _i0_2++) {
                _nw17[_i0_2] = _init2(new BigNumber(_i0_2));
              }
              _135_v66 = _nw17;
              let _index4 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_135_v66).length));
              (_135_v66)[_index4] = _dafny.Seq.UnicodeFromString("txf");
              let _index5 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_135_v66).length));
              (_135_v66)[_index5] = p0;
              let _143_v67;
              _143_v67 = _dafny.MultiSet.fromElements(_46_v0, (p0)[_module.__default.safeIndex((_dafny.ZERO).minus(_118_v52), new BigNumber((p0).length))]);
              (globalState).f25 = _module.__default.fm2(_143_v67, _118_v52, _118_v52, false, globalState);
              let _144_v68;
              let _nw18 = new _module.C0();
              _nw18.__ctor(_108_v43);
              _144_v68 = _nw18;
              let _rhs3 = _144_v68;
              let _rhs4 = _108_v43;
              _144_v68 = _rhs3;
              _108_v43 = _rhs4;
              let _145_v69;
              let _nw19 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
              _145_v69 = _nw19;
              let _146_v70;
              _146_v70 = _dafny.Map.Empty.slice().updateUnsafe(_145_v69,_144_v68);
              let _147_v71;
              _147_v71 = _dafny.Seq.of(_144_v68, (((_146_v70).contains(_145_v69)) ? ((_146_v70).get(_145_v69)) : (_144_v68)));
              let _148_v72;
              _148_v72 = _dafny.Map.Empty.slice().updateUnsafe(_118_v52,_144_v68);
              let _149_v73;
              _149_v73 = _module.D8.create_DC22(_147_v71);
              let _150_v74;
              _150_v74 = _dafny.Map.Empty.slice().updateUnsafe(_108_v43,_147_v71);
              let _151_v75;
              _151_v75 = _dafny.Seq.of(!(!((_144_v68).f29)), true);
              let _152_v76;
              let _nw20 = Array((new BigNumber(26)).toNumber());
              _nw20[(_dafny.ZERO).toNumber()] = _147_v71;
              _nw20[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_144_v68, _144_v68);
              _nw20[(new BigNumber(2)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_144_v68), _dafny.Seq.update(_147_v71, _module.__default.safeIndex(_118_v52, new BigNumber((_147_v71).length)), _144_v68));
              _nw20[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_144_v68, _144_v68, _144_v68);
              _nw20[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_144_v68, _144_v68, (((_148_v72).contains(_118_v52)) ? ((_148_v72).get(_118_v52)) : (_144_v68)), _144_v68, _144_v68);
              _nw20[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_147_v71, _147_v71);
              _nw20[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_147_v71, (_149_v73).dtor_cf44);
              _nw20[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_144_v68);
              _nw20[(new BigNumber(9)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(10)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(11)).toNumber()] = (((_144_v68).f29) ? (_147_v71) : (_dafny.Seq.update(_147_v71, _module.__default.safeIndex(_118_v52, new BigNumber((_147_v71).length)), _144_v68)));
              _nw20[(new BigNumber(12)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(13)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_147_v71, _147_v71);
              _nw20[(new BigNumber(15)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(16)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(17)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(18)).toNumber()] = (((_150_v74).contains(_108_v43)) ? ((_150_v74).get(_108_v43)) : (_147_v71));
              _nw20[(new BigNumber(19)).toNumber()] = (((_144_v68).f29) ? (_147_v71) : (_147_v71));
              _nw20[(new BigNumber(20)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(21)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(22)).toNumber()] = _147_v71;
              _nw20[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_147_v71, _147_v71);
              _nw20[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_147_v71, _module.__default.safeIndex(new BigNumber(((_135_v66)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_135_v66).length))]).length), new BigNumber((_147_v71).length)), (_147_v71)[_module.__default.safeIndex(new BigNumber((_151_v75).length), new BigNumber((_147_v71).length))]), _dafny.Seq.update(_dafny.Seq.of(_144_v68), _module.__default.safeIndex(_118_v52, new BigNumber((_dafny.Seq.of(_144_v68)).length)), _144_v68));
              _nw20[(new BigNumber(25)).toNumber()] = _147_v71;
              _152_v76 = _nw20;
              let _index6 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_152_v76).length));
              (_152_v76)[_index6] = _dafny.Seq.Concat(_147_v71, _147_v71);
              let _153_v77;
              _153_v77 = _dafny.Map.Empty.slice().updateUnsafe(_144_v68,_108_v43);
              let _154_v78;
              _154_v78 = _dafny.Set.fromElements(!((_144_v68).f29), (((_153_v77).contains(_144_v68)) ? ((_153_v77).get(_144_v68)) : (!(false))), (_144_v68).f29, (_144_v68).f29, (_144_v68).f29);
              let _index7 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_152_v76).length));
              let _rhs5 = _147_v71;
              let _rhs6 = ((_144_v68).f29) && (!(!(_118_v52).isEqualTo(new BigNumber((_154_v78).length))));
              let _rhs7 = false;
              let _lhs1 = _152_v76;
              let _lhs2 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_152_v76).length));
              let _lhs3 = globalState;
              _lhs1[_lhs2] = _rhs5;
              _lhs3.f14 = _rhs6;
              _108_v43 = _rhs7;
              _144_v68 = _144_v68;
            }
            let _155_v79;
            let _nw21 = new _module.C0();
            _nw21.__ctor(_108_v43);
            _155_v79 = _nw21;
            let _156_v80;
            _156_v80 = _dafny.Seq.of(_155_v79);
            let _157_v81;
            _157_v81 = _module.D8.create_DC22(_156_v80);
            let _158_v82;
            let _nw22 = Array((new BigNumber(5)).toNumber()).fill(false);
            _158_v82 = _nw22;
            let _159_v83;
            _159_v83 = _module.D5.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(418)), ((_160_v0) => function (_161_i7) {
  return _160_v0;
})(_46_v0)), _108_v43, false, (_155_v79).f29, _158_v82);
            let _162_v84;
            _162_v84 = _dafny.Map.Empty.slice().updateUnsafe(_158_v82,_159_v83);
            let _rhs8 = _46_v0;
            let _rhs9 = p0;
            let _rhs10 = _118_v52;
            let _rhs11 = _157_v81;
            let _rhs12 = new BigNumber(((_162_v84).Merge(_162_v84)).length);
            let _lhs4 = globalState;
            let _lhs5 = globalState;
            _lhs4.f2 = _rhs8;
            r0 = _rhs9;
            _118_v52 = _rhs10;
            _157_v81 = _rhs11;
            _lhs5.f17 = _rhs12;
            let _163_v85;
            _163_v85 = _dafny.MultiSet.fromElements(_108_v43);
            let _164_v86;
            _164_v86 = _dafny.Map.Empty.slice().updateUnsafe((_155_v79).f29,_118_v52);
            let _165_v87;
            _165_v87 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_118_v52),_118_v52);
            let _166_v88;
            _166_v88 = _dafny.Seq.of(_118_v52);
            let _167_v89;
            _167_v89 = _dafny.Seq.of(_108_v43, false);
            let _168_v90;
            let _nw23 = Array((new BigNumber(11)).toNumber());
            _nw23[(_dafny.ZERO).toNumber()] = new BigNumber((_164_v86).length);
            _nw23[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_118_v52);
            _nw23[(new BigNumber(2)).toNumber()] = _118_v52;
            _nw23[(new BigNumber(3)).toNumber()] = _118_v52;
            _nw23[(new BigNumber(4)).toNumber()] = new BigNumber((_165_v87).length);
            _nw23[(new BigNumber(5)).toNumber()] = _118_v52;
            _nw23[(new BigNumber(6)).toNumber()] = new BigNumber((_166_v88).length);
            _nw23[(new BigNumber(7)).toNumber()] = new BigNumber((_167_v89).length);
            _nw23[(new BigNumber(8)).toNumber()] = _118_v52;
            _nw23[(new BigNumber(9)).toNumber()] = _118_v52;
            _nw23[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_118_v52);
            _168_v90 = _nw23;
            let _169_v91;
            _169_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_163_v85).cardinality()),_168_v90);
            _169_v91 = (_169_v91).update((_118_v52).multipliedBy(new BigNumber(941)), _168_v90);
            let _170_v92;
            let _nw24 = new _module.C0();
            _nw24.__ctor((_155_v79).f29);
            _170_v92 = _nw24;
          }
        }
      }
      let _hi0 = _module.__default.safeModuloInt(new BigNumber((_120_v54).length), _118_v52);
      for (let _171_i8 = _118_v52; _171_i8.isLessThan(_hi0); _171_i8 = _171_i8.plus(_dafny.ONE)) {
        if (_108_v43) {
          let _172_v93;
          let _nw25 = new _module.C0();
          _nw25.__ctor(_module.__default.fm0(p0, _108_v43, globalState));
          _172_v93 = _nw25;
          let _173_v94;
          _173_v94 = _dafny.Seq.of(_172_v93, _172_v93, _172_v93);
          let _174_v95;
          let _nw26 = Array((new BigNumber(15)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = _172_v93;
          _nw26[(_dafny.ONE).toNumber()] = _172_v93;
          _nw26[(new BigNumber(2)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(3)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(4)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(5)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(6)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(7)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(8)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(9)).toNumber()] = (_173_v94)[_module.__default.safeIndex(_118_v52, new BigNumber((_173_v94).length))];
          _nw26[(new BigNumber(10)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(11)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(12)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(13)).toNumber()] = _172_v93;
          _nw26[(new BigNumber(14)).toNumber()] = _172_v93;
          _174_v95 = _nw26;
          let _index8 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_174_v95).length));
          (_174_v95)[_index8] = _172_v93;
          let _index9 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_174_v95).length));
          (_174_v95)[_index9] = _172_v93;
          let _175_v96;
          _175_v96 = _dafny.Seq.of(_108_v43, _module.__default.fm0(p0, _108_v43, globalState));
          let _176_v97;
          _176_v97 = _dafny.MultiSet.fromElements(_118_v52);
          let _177_v98;
          _177_v98 = _dafny.Map.Empty.slice().updateUnsafe(_118_v52,(_172_v93).f29);
          let _178_v99;
          _178_v99 = _dafny.Set.fromElements((((_177_v98).contains(_171_i8)) ? ((_177_v98).get(_171_i8)) : (!(true))));
          let _179_v100;
          let _nw27 = Array((new BigNumber(25)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _108_v43;
          _nw27[(_dafny.ONE).toNumber()] = _108_v43;
          _nw27[(new BigNumber(2)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(3)).toNumber()] = false;
          _nw27[(new BigNumber(4)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(5)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(6)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(7)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(8)).toNumber()] = true;
          _nw27[(new BigNumber(9)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(10)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(11)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(12)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(13)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(14)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(15)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(16)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(17)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(18)).toNumber()] = !(_108_v43);
          _nw27[(new BigNumber(19)).toNumber()] = true;
          _nw27[(new BigNumber(20)).toNumber()] = (_172_v93).f29;
          _nw27[(new BigNumber(21)).toNumber()] = true;
          _nw27[(new BigNumber(22)).toNumber()] = _108_v43;
          _nw27[(new BigNumber(23)).toNumber()] = true;
          _nw27[(new BigNumber(24)).toNumber()] = (_172_v93).f29;
          _179_v100 = _nw27;
          let _180_v101;
          _180_v101 = _dafny.Map.Empty.slice().updateUnsafe(_179_v100,true);
          let _181_v102;
          let _nw28 = Array((new BigNumber(28)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _108_v43;
          _nw28[(_dafny.ONE).toNumber()] = (_175_v96)[_module.__default.safeIndex(new BigNumber((_173_v94).length), new BigNumber((_175_v96).length))];
          _nw28[(new BigNumber(2)).toNumber()] = (false) && (false);
          _nw28[(new BigNumber(3)).toNumber()] = (_172_v93).f29;
          _nw28[(new BigNumber(4)).toNumber()] = !(_dafny.MultiSet.fromElements(_118_v52)).equals(_176_v97);
          _nw28[(new BigNumber(5)).toNumber()] = (true) === (_module.__default.fm0(_dafny.Seq.UnicodeFromString("shqrkmdpp"), true, globalState));
          _nw28[(new BigNumber(6)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(7)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(8)).toNumber()] = (_172_v93).f29;
          _nw28[(new BigNumber(9)).toNumber()] = (_172_v93).f29;
          _nw28[(new BigNumber(10)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(11)).toNumber()] = (new BigNumber((_178_v99).length)).isLessThan(_118_v52);
          _nw28[(new BigNumber(12)).toNumber()] = (new BigNumber((p0).length)).isLessThan(new BigNumber(-454));
          _nw28[(new BigNumber(13)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(14)).toNumber()] = _module.__default.fm0(p0, !(false), globalState);
          _nw28[(new BigNumber(15)).toNumber()] = (_172_v93).f29;
          _nw28[(new BigNumber(16)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(17)).toNumber()] = !_dafny.Seq.contains(_175_v96, !((((_180_v101).contains(_179_v100)) ? ((_180_v101).get(_179_v100)) : (_108_v43))));
          _nw28[(new BigNumber(18)).toNumber()] = ((_172_v93).f29) || (_108_v43);
          _nw28[(new BigNumber(19)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(20)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(21)).toNumber()] = ((_dafny.ZERO).minus(_118_v52)).isLessThan(_171_i8);
          _nw28[(new BigNumber(22)).toNumber()] = true;
          _nw28[(new BigNumber(23)).toNumber()] = (_172_v93).f29;
          _nw28[(new BigNumber(24)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(25)).toNumber()] = _108_v43;
          _nw28[(new BigNumber(26)).toNumber()] = (_172_v93).f29;
          _nw28[(new BigNumber(27)).toNumber()] = (_module.D6.create_DC18(p0, new BigNumber((_175_v96).length), _108_v43, new BigNumber(834))).dtor_cf36;
          _181_v102 = _nw28;
          _181_v102 = _179_v100;
          let _182_v103;
          _182_v103 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(p0, _171_i8, globalState),_dafny.Seq.Create(_module.__default.abs(new BigNumber(848)), ((_183_v52) => function (_184_i9) {
            return _dafny.Map.Empty.slice().updateUnsafe(true,_183_v52);
          })(_118_v52)));
          let _185_v104;
          _185_v104 = _dafny.Map.Empty.slice().updateUnsafe((_172_v93).f29,_171_i8);
          let _186_v105;
          _186_v105 = _dafny.Seq.of(_185_v104);
          _182_v103 = (_182_v103).update(_46_v0, _186_v105);
          let _187_v106;
          _187_v106 = _dafny.MultiSet.fromElements((_172_v93).f29);
          let _188_v107;
          _188_v107 = _dafny.Seq.of((((_187_v106).contains(!((_172_v93).f29))) ? ((_187_v106).get(!((_172_v93).f29))) : (_171_i8)), _118_v52);
          let _189_v108;
          let _nw29 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _189_v108 = _nw29;
          let _190_v109;
          _190_v109 = _dafny.Set.fromElements(_189_v108, _189_v108);
          (globalState).f27 = (_188_v107)[_module.__default.safeIndex(new BigNumber(((_190_v109).Intersect(_190_v109)).length), new BigNumber((_188_v107).length))];
          let _191_v110;
          let _nw30 = new _module.C0();
          _nw30.__ctor(true);
          _191_v110 = _nw30;
        } else {
          let _192_v111;
          _192_v111 = _dafny.MultiSet.fromElements(true, _108_v43);
          (globalState).f17 = _module.__default.safeDivisionInt(_171_i8, _module.__default.safeDivisionInt((((_192_v111).contains(_108_v43)) ? ((_192_v111).get(_108_v43)) : (_118_v52)), new BigNumber(892)));
          (globalState).f14 = _108_v43;
          let _193_v112;
          _193_v112 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("kvspt"));
          let _194_v113;
          _194_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_119_v53).length),_dafny.Seq.UnicodeFromString("yjrw"));
          let _195_v114;
          let _nw31 = Array((new BigNumber(20)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = (_193_v112)[_module.__default.safeIndex(_118_v52, new BigNumber((_193_v112).length))];
          _nw31[(_dafny.ONE).toNumber()] = p0;
          _nw31[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), ((_196_v0) => function (_197_i10) {
            return _196_v0;
          })(_46_v0));
          _nw31[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("iattnj");
          _nw31[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(p0, p0);
          _nw31[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("amks");
          _nw31[(new BigNumber(6)).toNumber()] = p0;
          _nw31[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p0, p0);
          _nw31[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("uut");
          _nw31[(new BigNumber(9)).toNumber()] = p0;
          _nw31[(new BigNumber(10)).toNumber()] = p0;
          _nw31[(new BigNumber(11)).toNumber()] = p0;
          _nw31[(new BigNumber(12)).toNumber()] = p0;
          _nw31[(new BigNumber(13)).toNumber()] = p0;
          _nw31[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("rivncvdvu");
          _nw31[(new BigNumber(15)).toNumber()] = (_193_v112)[_module.__default.safeIndex(_118_v52, new BigNumber((_193_v112).length))];
          _nw31[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("pukcc");
          _nw31[(new BigNumber(17)).toNumber()] = p0;
          _nw31[(new BigNumber(18)).toNumber()] = (((_194_v113).contains(new BigNumber((p0).length))) ? ((_194_v113).get(new BigNumber((p0).length))) : (_dafny.Seq.UnicodeFromString("mpehwcpgi")));
          _nw31[(new BigNumber(19)).toNumber()] = p0;
          _195_v114 = _nw31;
          let _index10 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_195_v114).length));
          (_195_v114)[_index10] = p0;
          let _198_v115;
          _198_v115 = _dafny.Map.Empty.slice().updateUnsafe(_108_v43,_dafny.Seq.UnicodeFromString("cugx"));
          let _index11 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_195_v114).length));
          (_195_v114)[_index11] = _dafny.Seq.Concat(p0, _dafny.Seq.Concat((((_198_v115).contains(_108_v43)) ? ((_198_v115).get(_108_v43)) : (p0)), p0));
          (globalState).f17 = ((_118_v52).minus(_171_i8)).plus(_118_v52);
          let _199_v116;
          _199_v116 = _module.D2.create_DC8(p0);
          let _200_v117;
          _200_v117 = _dafny.Seq.of(_module.__default.fm14(_199_v116, globalState), _192_v111, _192_v111, _192_v111, _192_v111);
          let _201_v118;
          _201_v118 = _dafny.Map.Empty.slice().updateUnsafe((_200_v117)[_module.__default.safeIndex(_118_v52, new BigNumber((_200_v117).length))],_108_v43);
          _201_v118 = _201_v118;
        }
        _118_v52 = new BigNumber((p0).length);
        let _202_v119;
        _202_v119 = _dafny.Seq.of(_118_v52);
        _202_v119 = _202_v119;
        let _203_v120;
        _203_v120 = _dafny.Map.Empty.slice().updateUnsafe(_108_v43,_171_i8);
        _203_v120 = _203_v120;
      }
      r0 = p0;
      let _204_v121;
      _204_v121 = _dafny.MultiSet.fromElements(false);
      let _205_v122;
      _205_v122 = _dafny.Seq.of((((_204_v121).contains(_108_v43)) ? ((_204_v121).get(_108_v43)) : (_118_v52)));
      let _206_v123;
      _206_v123 = _dafny.Seq.of(_108_v43);
      let _207_v124;
      _207_v124 = _dafny.Map.Empty.slice().updateUnsafe(_118_v52,new BigNumber((_206_v123).length));
      let _208_v125;
      _208_v125 = _dafny.Set.fromElements(_207_v124);
      let _209_v126;
      _209_v126 = _dafny.MultiSet.fromElements(_118_v52, _118_v52);
      let _210_v127;
      _210_v127 = _dafny.Map.Empty.slice().updateUnsafe(_118_v52,_108_v43);
      let _211_v128;
      _211_v128 = _dafny.Map.Empty.slice().updateUnsafe(_210_v127,false);
      let _nw32 = Array((new BigNumber(25)).toNumber());
      _nw32[(_dafny.ZERO).toNumber()] = new BigNumber((_205_v122).length);
      _nw32[(_dafny.ONE).toNumber()] = new BigNumber((_206_v123).length);
      _nw32[(new BigNumber(2)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(3)).toNumber()] = new BigNumber((_208_v125).length);
      _nw32[(new BigNumber(4)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(5)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(6)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_118_v52), new BigNumber((p0).length));
      _nw32[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_118_v52);
      _nw32[(new BigNumber(9)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(10)).toNumber()] = ((_108_v43) ? (new BigNumber(721)) : (_118_v52));
      _nw32[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(p0, p0), _module.__default.safeIndex(_118_v52, new BigNumber((_dafny.Seq.Concat(p0, p0)).length)), new _dafny.CodePoint('b'.codePointAt(0)))).length);
      _nw32[(new BigNumber(12)).toNumber()] = new BigNumber(580);
      _nw32[(new BigNumber(13)).toNumber()] = new BigNumber((_module.__default.fm9(_118_v52, globalState)).length);
      _nw32[(new BigNumber(14)).toNumber()] = new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(_118_v52, _118_v52))).Union(_209_v126)).cardinality());
      _nw32[(new BigNumber(15)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(16)).toNumber()] = (_118_v52).plus(_118_v52);
      _nw32[(new BigNumber(17)).toNumber()] = (((_204_v121).contains(_108_v43)) ? ((_204_v121).get(_108_v43)) : (_118_v52));
      _nw32[(new BigNumber(18)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(19)).toNumber()] = _118_v52;
      _nw32[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((p0).length), _118_v52);
      _nw32[(new BigNumber(21)).toNumber()] = (new BigNumber((_module.__default.fm12(_118_v52, new BigNumber(-213), _108_v43, globalState)).length)).multipliedBy(_118_v52);
      _nw32[(new BigNumber(22)).toNumber()] = new BigNumber((p0).length);
      _nw32[(new BigNumber(23)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_118_v52), new BigNumber((_211_v128).length));
      _nw32[(new BigNumber(24)).toNumber()] = (new BigNumber(-200)).minus((_dafny.ZERO).minus(_118_v52));
      r1 = _nw32;
      let _212_v129;
      _212_v129 = _module.D7.create_DC19(_204_v121);
      let _213_v130;
      _213_v130 = _module.D7.create_DC21(_212_v129);
      let _pat_let_tv0 = _118_v52;
      let _pat_let_tv1 = _209_v126;
      r2 = (_dafny.ZERO).minus(function (_source5) {
        if (_source5.is_DC20) {
          let _214___mcc_h6 = (_source5).cf39;
          let _215___mcc_h7 = (_source5).cf40;
          let _216___mcc_h8 = (_source5).cf41;
          let _217___mcc_h9 = (_source5).cf42;
          let _218_cf42 = _217___mcc_h9;
          let _219_cf41 = _216___mcc_h8;
          let _220_cf40 = _215___mcc_h7;
          let _221_cf39 = _214___mcc_h6;
          return _219_cf41;
        } else if (_source5.is_DC19) {
          let _222___mcc_h10 = (_source5).cf38;
          let _223_cf38 = _222___mcc_h10;
          return _pat_let_tv0;
        } else {
          let _224___mcc_h11 = (_source5).cf43;
          let _225_cf43 = _224___mcc_h11;
          return new BigNumber((_pat_let_tv1).cardinality());
        }
      }(_213_v130));
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _226_v0;
      _226_v0 = true;
      let _227_v1;
      _227_v1 = _dafny.MultiSet.fromElements(_226_v0, _226_v0);
      let _228_v2;
      _228_v2 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_226_v0)), _227_v1);
      let _229_v3;
      _229_v3 = _dafny.Seq.UnicodeFromString("ujuh");
      let _230_v4;
      _230_v4 = _dafny.Seq.of(_226_v0, _226_v0, _226_v0);
      let _231_v5;
      _231_v5 = _dafny.Map.Empty.slice().updateUnsafe((_230_v4)[_module.__default.safeIndex(new BigNumber((_229_v3).length), new BigNumber((_230_v4).length))],_230_v4);
      let _232_v6;
      let _nw33 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
      _232_v6 = _nw33;
      let _233_v7;
      _233_v7 = new BigNumber(773);
      let _234_v8;
      _234_v8 = _dafny.Set.fromElements(_226_v0);
      let _235_v9;
      let _nw34 = Array((new BigNumber(20)).toNumber());
      _nw34[(_dafny.ZERO).toNumber()] = _233_v7;
      _nw34[(_dafny.ONE).toNumber()] = new BigNumber(-741);
      _nw34[(new BigNumber(2)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(3)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(4)).toNumber()] = new BigNumber(-617);
      _nw34[(new BigNumber(5)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(6)).toNumber()] = new BigNumber((_234_v8).length);
      _nw34[(new BigNumber(7)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(8)).toNumber()] = new BigNumber(744);
      _nw34[(new BigNumber(9)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(10)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(11)).toNumber()] = new BigNumber(462);
      _nw34[(new BigNumber(12)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("pvc")).length);
      _nw34[(new BigNumber(14)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(15)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(16)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(17)).toNumber()] = new BigNumber(726);
      _nw34[(new BigNumber(18)).toNumber()] = _233_v7;
      _nw34[(new BigNumber(19)).toNumber()] = _233_v7;
      _235_v9 = _nw34;
      let _236_v10;
      _236_v10 = _dafny.MultiSet.fromElements(_235_v9);
      let _237_v12;
      _237_v12 = _dafny.Map.Empty.slice().updateUnsafe(_226_v0,!(_226_v0));
      let _238_v13;
      _238_v13 = _dafny.Seq.of(_237_v12, _237_v12);
      let _239_v14;
      let _init3 = ((_240_v0) => function (_241_i0) {
        return _240_v0;
      })(_226_v0);
      let _nw35 = Array((new BigNumber(16)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw35.length); _i0_3++) {
        _nw35[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _239_v14 = _nw35;
      let _242_v15;
      _242_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(155),_226_v0);
      let _243_v16;
      _243_v16 = new _dafny.CodePoint('a'.codePointAt(0));
      let _244_v17;
      _244_v17 = _dafny.Set.fromElements(_243_v16, _243_v16);
      let _245_v18;
      _245_v18 = _dafny.Map.Empty.slice().updateUnsafe(_242_v15,_244_v17);
      let _246_globalState;
      let _nw36 = new _module.GlobalState();
      _nw36.__ctor(_228_v2, new BigNumber(-624), new _dafny.CodePoint('l'.codePointAt(0)), new BigNumber(951), true, new BigNumber(868), _229_v3, _231_v5, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nwvoes"), _229_v3), _232_v6, new BigNumber(113), (_227_v1).update(_226_v0, _module.__default.abs(_233_v7)), new BigNumber(664), new BigNumber(310), false, new BigNumber(536), new BigNumber(495), new BigNumber(-530), (_236_v10).Difference(_dafny.MultiSet.fromElements(_235_v9, _235_v9, _235_v9)), function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_238_v13).Elements) {
          let _247_v11 = _compr_1;
          if (_dafny.Seq.contains(_238_v13, _247_v11)) {
            _coll1.push([_247_v11,_233_v7]);
          }
        }
        return _coll1;
      }(), _239_v14, new BigNumber(294), true, new BigNumber(119), (_245_v18).Merge(_dafny.Map.Empty.slice().updateUnsafe(_242_v15,_244_v17)), new BigNumber(-643), _227_v1, new BigNumber(239), _235_v9);
      _246_globalState = _nw36;
      let _248_i1;
      _248_i1 = _dafny.ZERO;
      L2: {
        while (!(!(_226_v0) || (_226_v0))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_248_i1)) {
              break L2;
            }
            _248_i1 = (_248_i1).plus(_dafny.ONE);
            let _249_v19;
            let _nw37 = new _module.C0();
            _nw37.__ctor((new BigNumber((_242_v15).length)).isLessThanOrEqualTo(new BigNumber(579)));
            _249_v19 = _nw37;
            let _250_v20;
            _250_v20 = _dafny.Seq.of(_229_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-469)), ((_251_v16) => function (_252_i2) {
              return _251_v16;
            })(_243_v16)));
            _226_v0 = _module.__default.fm0((_250_v20)[_module.__default.safeIndex(new BigNumber(149), new BigNumber((_250_v20).length))], _226_v0, _246_globalState);
            _235_v9 = _235_v9;
            (_246_globalState).f27 = _233_v7;
          }
        }
      }
      let _253_v21;
      let _254_v22;
      let _255_v23;
      let _out0;
      let _out1;
      let _out2;
      let _outcollector0 = _module.__default.m0(_229_v3, _246_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _253_v21 = _out0;
      _254_v22 = _out1;
      _255_v23 = _out2;
      let _hi1 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_256_v16) => function (_257_i4) {
        return _256_v16;
      })(_243_v16))).length);
      for (let _258_i3 = new BigNumber((_253_v21).length); _258_i3.isLessThan(_hi1); _258_i3 = _258_i3.plus(_dafny.ONE)) {
        _237_v12 = (_237_v12).update((_234_v8).IsProperSubsetOf(_234_v8), _226_v0);
        let _source6 = _module.__default.fm1(_246_globalState);
        if (_source6.is_DC0) {
          let _259___mcc_h0 = (_source6).cf0;
          let _260___mcc_h1 = (_source6).cf1;
          let _261_cf1 = _260___mcc_h1;
          let _262_cf0 = _259___mcc_h0;
          let _index12 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_254_v22).length));
          (_254_v22)[_index12] = _module.__default.safeDivisionInt(_255_v23, _255_v23);
          let _263_v24;
          let _nw38 = new _module.C0();
          _nw38.__ctor(_226_v0);
          _263_v24 = _nw38;
          let _264_v25;
          let _nw39 = Array((_dafny.ONE).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = ((_226_v0) ? (_263_v24) : (_263_v24));
          _264_v25 = _nw39;
          let _265_v26;
          _265_v26 = _dafny.Map.Empty.slice().updateUnsafe((_263_v24).f29,_233_v7);
          let _index13 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_254_v22).length));
          let _rhs13 = (((_226_v0) && ((_263_v24).f29)) ? ((((_265_v26).contains((_263_v24).f29)) ? ((_265_v26).get((_263_v24).f29)) : (_261_cf1))) : (_255_v23));
          let _rhs14 = _258_i3;
          let _rhs15 = _264_v25;
          let _lhs6 = _254_v22;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_254_v22).length));
          _lhs6[_lhs7] = _rhs13;
          _262_cf0 = _rhs14;
          _264_v25 = _rhs15;
          let _266_v27;
          let _nw40 = new _module.C0();
          _nw40.__ctor(_226_v0);
          _266_v27 = _nw40;
          let _267_v28;
          let _nw41 = Array((new BigNumber(18)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _236_v10;
          _nw41[(_dafny.ONE).toNumber()] = ((false) ? (_236_v10) : (_236_v10));
          _nw41[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_235_v9, _235_v9, _235_v9);
          _nw41[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_235_v9, _235_v9);
          _nw41[(new BigNumber(4)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(5)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(6)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(7)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(8)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(9)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(10)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(11)).toNumber()] = (_236_v10).update(_235_v9, _module.__default.abs(new BigNumber(437)));
          _nw41[(new BigNumber(12)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(13)).toNumber()] = (_236_v10).Difference(_236_v10);
          _nw41[(new BigNumber(14)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(15)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(16)).toNumber()] = _236_v10;
          _nw41[(new BigNumber(17)).toNumber()] = (_236_v10).Union(_dafny.MultiSet.fromElements(_235_v9));
          _267_v28 = _nw41;
          let _index14 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_267_v28).length));
          (_267_v28)[_index14] = _dafny.MultiSet.fromElements(_235_v9, _235_v9, _235_v9);
          let _268_v29;
          let _nw42 = Array((new BigNumber(17)).toNumber()).fill(_module.D0.Default());
          _268_v29 = _nw42;
          let _269_v30;
          _269_v30 = _module.D0.create_DC0(new BigNumber(781), (_dafny.ZERO).minus(_261_cf1));
          let _pat_let_tv2 = _261_cf1;
          let _index15 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_268_v29).length));
          (_268_v29)[_index15] = function (_pat_let2_0) {
            return function (_270_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_271_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_271_dt__update_hcf0_h0, (_270_dt__update__tmp_h0).dtor_cf1);
                }(_pat_let3_0);
              }(_pat_let_tv2);
            }(_pat_let2_0);
          }(_269_v30);
          let _index16 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_267_v28).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_268_v29).length));
          let _rhs16 = _dafny.MultiSet.fromElements(_235_v9, _235_v9, _235_v9, _235_v9, _235_v9);
          let _rhs17 = _module.D0.create_DC0(_262_cf0, _262_cf0);
          let _rhs18 = (_266_v27).f29;
          let _rhs19 = _263_v24;
          let _rhs20 = ((_dafny.ZERO).minus(_261_cf1)).plus((((_266_v27).f29) ? (_261_cf1) : (_262_cf0)));
          let _lhs8 = _267_v28;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_267_v28).length));
          let _lhs10 = _268_v29;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_268_v29).length));
          let _lhs12 = _246_globalState;
          let _lhs13 = _246_globalState;
          _lhs8[_lhs9] = _rhs16;
          _lhs10[_lhs11] = _rhs17;
          _lhs12.f14 = _rhs18;
          _263_v24 = _rhs19;
          _lhs13.f13 = _rhs20;
          let _272_v31;
          _272_v31 = _dafny.Map.Empty.slice().updateUnsafe(true,_266_v27);
          _272_v31 = (_272_v31).update(!((_266_v27).f29), _263_v24);
        } else {
          (_246_globalState).f27 = (_233_v7).plus(new BigNumber(316));
          let _index18 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_239_v14).length));
          (_239_v14)[_index18] = (_226_v0) === (_226_v0);
          let _index19 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_239_v14).length));
          (_239_v14)[_index19] = !(_226_v0);
          _255_v23 = _255_v23;
          let _273_v32;
          let _init4 = ((_274_v0, _275_i3) => function (_276_i5) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_274_v0,new BigNumber(889))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_274_v0,_275_i3));
          })(_226_v0, _258_i3);
          let _nw43 = Array((new BigNumber(22)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw43.length); _i0_4++) {
            _nw43[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _273_v32 = _nw43;
          let _277_v33;
          _277_v33 = _dafny.Map.Empty.slice().updateUnsafe((_239_v14)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_239_v14).length))],new BigNumber(878));
          let _index20 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_273_v32).length));
          (_273_v32)[_index20] = (((((_237_v12).contains((_239_v14)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_239_v14).length))])) ? ((_237_v12).get((_239_v14)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_239_v14).length))])) : (_226_v0))) ? (_277_v33) : (_277_v33));
          let _index21 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_273_v32).length));
          (_273_v32)[_index21] = _277_v33;
        }
        let _278_v34;
        let _nw44 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
        _278_v34 = _nw44;
        let _index22 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_278_v34).length));
        (_278_v34)[_index22] = _dafny.Seq.Concat(_230_v4, _230_v4);
        let _index23 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_278_v34).length));
        (_278_v34)[_index23] = _dafny.Seq.Concat(_230_v4, _230_v4);
        let _279_v35;
        let _nw45 = new _module.C0();
        _nw45.__ctor(_226_v0);
        _279_v35 = _nw45;
      }
      let _280_v36;
      let _nw46 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
      _280_v36 = _nw46;
      let _281_v37;
      _281_v37 = _dafny.MultiSet.fromElements(_239_v14, _239_v14, _239_v14, _239_v14, _239_v14);
      let _index24 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_280_v36).length));
      (_280_v36)[_index24] = _281_v37;
      let _index25 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_280_v36).length));
      (_280_v36)[_index25] = ((_281_v37).Difference(_281_v37)).Intersect((_281_v37).Intersect(_281_v37));
      if ((_255_v23).isLessThanOrEqualTo(_255_v23)) {
        _243_v16 = _243_v16;
        let _282_v38;
        let _283_v39;
        let _284_v40;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = _module.__default.m0(_229_v3, _246_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _282_v38 = _out3;
        _283_v39 = _out4;
        _284_v40 = _out5;
        let _285_v41;
        let _286_v42;
        let _287_v43;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector2 = _module.__default.m0(_dafny.Seq.UnicodeFromString("konvcdp"), _246_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _285_v41 = _out6;
        _286_v42 = _out7;
        _287_v43 = _out8;
        if (_226_v0) {
          (_246_globalState).f13 = _233_v7;
          let _288_v44;
          let _nw47 = new _module.C0();
          _nw47.__ctor(true);
          _288_v44 = _nw47;
          let _289_v45;
          _289_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(309),_284_v40);
          let _290_v46;
          let _nw48 = new _module.C0();
          _nw48.__ctor((_289_v45).equals(_289_v45));
          _290_v46 = _nw48;
          _229_v3 = _dafny.Seq.UnicodeFromString("i");
          (_246_globalState).f14 = (((_242_v15).contains(_287_v43)) ? ((_242_v15).get(_287_v43)) : (((_dafny.ZERO).minus((_dafny.ZERO).minus(_284_v40))).isLessThanOrEqualTo(_284_v40)));
        } else {
          (_246_globalState).f27 = new BigNumber(157);
          let _291_v47;
          _291_v47 = _dafny.Map.Empty.slice().updateUnsafe(_243_v16,_239_v14);
          let _rhs21 = ((_291_v47).Merge(_dafny.Map.Empty.slice().updateUnsafe(_243_v16,_239_v14))).update(_243_v16, _239_v14);
          let _rhs22 = _module.__default.fm0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-390)), ((_292_v16) => function (_293_i6) {
            return _292_v16;
          })(_243_v16)), !(_242_v15).equals(_242_v15), _246_globalState);
          let _rhs23 = _227_v1;
          let _lhs14 = _246_globalState;
          _291_v47 = _rhs21;
          _lhs14.f14 = _rhs22;
          _227_v1 = _rhs23;
          let _294_v48;
          _294_v48 = _dafny.Seq.of(_255_v23);
          let _295_v49;
          _295_v49 = _dafny.Map.Empty.slice().updateUnsafe(_284_v40,new BigNumber(663));
          let _296_v50;
          _296_v50 = _dafny.Seq.of(new BigNumber(745), (_294_v48)[_module.__default.safeIndex(new BigNumber((_295_v49).length), new BigNumber((_294_v48).length))]);
          (_246_globalState).f13 = _module.__default.safeDivisionInt(new BigNumber(57), (_287_v43).plus(new BigNumber((_296_v50).length)));
          let _297_v51;
          let _nw49 = new _module.C0();
          _nw49.__ctor(_dafny.Seq.IsPrefixOf(_294_v48, _296_v50));
          _297_v51 = _nw49;
          let _298_v52;
          _298_v52 = _dafny.MultiSet.fromElements(_243_v16, _243_v16);
          let _299_v53;
          _299_v53 = _dafny.Set.fromElements(_module.__default.fm2(_298_v52, _255_v23, _233_v7, _226_v0, _246_globalState));
          (_246_globalState).f25 = new BigNumber((((_242_v15).update(new BigNumber((_299_v53).length), (_297_v51).f29)).update(new BigNumber(886), _226_v0)).length);
        }
        let _300_v54;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_226_v0);
        _300_v54 = _nw50;
      } else {
        (_246_globalState).f2 = _243_v16;
        (_246_globalState).f25 = _255_v23;
        _226_v0 = (_255_v23).isLessThanOrEqualTo(_233_v7);
        (_246_globalState).f27 = (_dafny.ZERO).minus(_233_v7);
        _237_v12 = (_237_v12).update(_226_v0, ((_module.__default.fm0(_229_v3, _226_v0, _246_globalState)) ? (_226_v0) : ((((_242_v15).contains(_255_v23)) ? ((_242_v15).get(_255_v23)) : (_226_v0)))));
      }
      let _301_v55;
      let _nw51 = new _module.C0();
      _nw51.__ctor(_dafny.Seq.contains(_253_v21, _243_v16));
      _301_v55 = _nw51;
      (_246_globalState).f27 = _255_v23;
      if ((_301_v55).f29) {
        (_246_globalState).f13 = _255_v23;
        (_246_globalState).f14 = (_301_v55).f29;
        let _index26 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_239_v14).length));
        (_239_v14)[_index26] = _226_v0;
        let _302_v56;
        _302_v56 = _module.D1.create_DC2(_253_v21);
        let _index27 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_239_v14).length));
        (_239_v14)[_index27] = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.update(_229_v3, _module.__default.safeIndex(_255_v23, new BigNumber((_229_v3).length)), _243_v16), (_302_v56).dtor_cf2), _229_v3);
        (_246_globalState).f14 = !((_239_v14)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_239_v14).length))]);
        let _303_v57;
        _303_v57 = _dafny.Set.fromElements(_255_v23);
        let _304_v58;
        _304_v58 = _dafny.Seq.of(_233_v7, (_dafny.ZERO).minus(new BigNumber(-473)), _233_v7);
        let _305_v59;
        _305_v59 = _dafny.Seq.of(_303_v57, _dafny.Set.fromElements((_dafny.ZERO).minus((_304_v58)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_304_v58).length))])));
        let _306_v60;
        _306_v60 = _dafny.MultiSet.fromElements(_303_v57, (_305_v59)[_module.__default.safeIndex(_255_v23, new BigNumber((_305_v59).length))]);
        _306_v60 = _306_v60;
      } else {
        let _307_v61;
        _307_v61 = _module.D1.create_DC4((_301_v55).f29, _233_v7);
        let _index28 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length));
        (_235_v9)[_index28] = (_307_v61).dtor_cf6;
        let _308_v62;
        _308_v62 = _dafny.Map.Empty.slice().updateUnsafe((_301_v55).f29,_237_v12);
        let _309_v63;
        _309_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_308_v62).length),_243_v16);
        let _310_v64;
        _310_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_309_v63).length),_dafny.Seq.of(_226_v0, (_301_v55).f29, _226_v0, (_301_v55).f29, (_301_v55).f29));
        let _311_v65;
        _311_v65 = _dafny.Map.Empty.slice().updateUnsafe((_310_v64).Merge(_310_v64),_255_v23);
        let _index29 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length));
        (_235_v9)[_index29] = (((_311_v65).contains(_310_v64)) ? ((_311_v65).get(_310_v64)) : (_255_v23));
        let _312_v66;
        _312_v66 = _dafny.Seq.of(new BigNumber(238), (_235_v9)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length))]);
        let _313_v67;
        _313_v67 = _dafny.Seq.of(_255_v23, (_312_v66)[_module.__default.safeIndex(_233_v7, new BigNumber((_312_v66).length))]);
        let _rhs24 = _312_v66;
        let _rhs25 = (_301_v55).f29;
        let _lhs15 = _246_globalState;
        _312_v66 = _rhs24;
        _lhs15.f14 = _rhs25;
        let _314_v68;
        _314_v68 = _module.D0.create_DC1();
        _314_v68 = _module.__default.fm3(_233_v7, _246_globalState);
        if (_226_v0) {
          _226_v0 = false;
          (_246_globalState).f17 = _233_v7;
          let _315_v69;
          let _316_v70;
          let _317_v71;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = _module.__default.m0(_253_v21, _246_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _315_v69 = _out9;
          _316_v70 = _out10;
          _317_v71 = _out11;
          let _318_v72;
          _318_v72 = _dafny.Map.Empty.slice().updateUnsafe((_301_v55).f29,new BigNumber((_dafny.Seq.of(_255_v23, _317_v71)).length));
          let _rhs26 = _317_v71;
          let _rhs27 = _module.__default.fm0(_253_v21, !(_226_v0), _246_globalState);
          let _rhs28 = _226_v0;
          let _rhs29 = ((_235_v9)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length))]).plus(_module.__default.safeModuloInt((((_318_v72).contains((_301_v55).f29)) ? ((_318_v72).get((_301_v55).f29)) : (_255_v23)), (_235_v9)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length))]));
          let _lhs16 = _246_globalState;
          let _lhs17 = _246_globalState;
          let _lhs18 = _246_globalState;
          let _lhs19 = _246_globalState;
          _lhs16.f25 = _rhs26;
          _lhs17.f14 = _rhs27;
          _lhs18.f14 = _rhs28;
          _lhs19.f25 = _rhs29;
          (_246_globalState).f14 = (((!(_module.__default.fm0(_253_v21, (_301_v55).f29, _246_globalState))) ? ((_301_v55).f29) : ((_301_v55).f29))) && ((_301_v55).f29);
        } else {
          let _319_v73;
          _319_v73 = _module.D1.create_DC2(_253_v21);
          _319_v73 = _module.__default.fm4((_233_v7).isLessThanOrEqualTo(new BigNumber(509)), (_301_v55).f29, _246_globalState);
          let _320_v74;
          _320_v74 = _dafny.Set.fromElements(_253_v21);
          let _321_v75;
          _321_v75 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-172)), function (_322_i7) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(57)), function (_323_i8) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }));
          let _rhs30 = (_320_v74).IsProperSubsetOf((_320_v74).Difference(function () {
            let _coll2 = new _dafny.Set();
            for (const _compr_2 of (_321_v75).Elements) {
              let _324_v76 = _compr_2;
              if (_dafny.Seq.contains(_321_v75, _324_v76)) {
                _coll2.add(_324_v76);
              }
            }
            return _coll2;
          }()));
          let _rhs31 = (_230_v4)[_module.__default.safeIndex((_235_v9)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length))], new BigNumber((_230_v4).length))];
          let _rhs32 = _314_v68;
          _226_v0 = _rhs30;
          _226_v0 = _rhs31;
          _314_v68 = _rhs32;
          let _index30 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length));
          (_235_v9)[_index30] = (_dafny.ZERO).minus((_255_v23).minus(((_dafny.ZERO).minus(_255_v23)).minus((_235_v9)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_235_v9).length))])));
          (_246_globalState).f17 = (_dafny.ZERO).minus((_313_v67)[_module.__default.safeIndex(new BigNumber((_312_v66).length), new BigNumber((_313_v67).length))]);
          (_246_globalState).f14 = false;
        }
        let _325_v77;
        let _326_v78;
        let _327_v79;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector4 = _module.__default.m0(_dafny.Seq.UnicodeFromString("arg"), _246_globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _325_v77 = _out12;
        _326_v78 = _out13;
        _327_v79 = _out14;
      }
      if ((_233_v7).isLessThanOrEqualTo(_255_v23)) {
        let _328_v80;
        let _329_v81;
        let _330_v82;
        let _out15;
        let _out16;
        let _out17;
        let _outcollector5 = _module.__default.m0(_229_v3, _246_globalState);
        _out15 = _outcollector5[0];
        _out16 = _outcollector5[1];
        _out17 = _outcollector5[2];
        _328_v80 = _out15;
        _329_v81 = _out16;
        _330_v82 = _out17;
        let _331_v83;
        let _332_v84;
        let _333_v85;
        let _out18;
        let _out19;
        let _out20;
        let _outcollector6 = _module.__default.m0(_229_v3, _246_globalState);
        _out18 = _outcollector6[0];
        _out19 = _outcollector6[1];
        _out20 = _outcollector6[2];
        _331_v83 = _out18;
        _332_v84 = _out19;
        _333_v85 = _out20;
        (_246_globalState).f14 = (_301_v55).f29;
        let _334_v86;
        let _nw52 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _334_v86 = _nw52;
        _334_v86 = _334_v86;
        let _index31 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_239_v14).length));
        (_239_v14)[_index31] = (_301_v55).f29;
        let _index32 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_239_v14).length));
        let _rhs33 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_330_v82));
        let _rhs34 = _226_v0;
        let _lhs20 = _246_globalState;
        let _lhs21 = _239_v14;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_239_v14).length));
        _lhs20.f17 = _rhs33;
        _lhs21[_lhs22] = _rhs34;
      } else {
        let _335_v87;
        _335_v87 = _dafny.MultiSet.fromElements(_243_v16, _243_v16, _243_v16, _243_v16, _243_v16);
        let _336_v88;
        _336_v88 = _dafny.Set.fromElements(_255_v23, _module.__default.fm2(_335_v87, _233_v7, new BigNumber(331), true, _246_globalState), ((_226_v0) ? (new BigNumber((_242_v15).length)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-400)), function (_337_i9) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length))), _255_v23);
        _336_v88 = ((_336_v88).Difference(_336_v88)).Intersect(function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_module.__default.fm5(_226_v0, _255_v23, (_301_v55).f29, _246_globalState)).Elements) {
            let _338_v89 = _compr_3;
            if (_dafny.Seq.contains(_module.__default.fm5(_226_v0, _255_v23, (_301_v55).f29, _246_globalState), _338_v89)) {
              _coll3.add((_338_v89).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("khn")).length)));
            }
          }
          return _coll3;
        }());
        let _339_v90;
        let _nw53 = new _module.C0();
        _nw53.__ctor((_301_v55).f29);
        _339_v90 = _nw53;
        if ((_301_v55).f29) {
          let _340_v91;
          _340_v91 = _dafny.Seq.of(_dafny.Seq.Concat(_253_v21, _229_v3), _229_v3, _dafny.Seq.UnicodeFromString("dhyao"));
          _340_v91 = _module.__default.fm6(_253_v21, _module.D1.create_DC2(_dafny.Seq.UnicodeFromString("yuhsgwf")), (_dafny.ZERO).minus(_255_v23), (_255_v23).isLessThanOrEqualTo(_233_v7), _246_globalState);
          (_246_globalState).f17 = _255_v23;
          let _341_v92;
          _341_v92 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_233_v7)),_339_v90);
          (_246_globalState).f14 = !((_341_v92).Merge((_341_v92).update((_dafny.ZERO).minus(_255_v23), _339_v90))).contains(new BigNumber((_dafny.Seq.of((_301_v55).f29)).length));
          let _index33 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_235_v9).length));
          (_235_v9)[_index33] = (_dafny.ZERO).minus(_255_v23);
          let _index34 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_235_v9).length));
          (_235_v9)[_index34] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_255_v23, _233_v7), _255_v23);
          let _342_v93;
          let _nw54 = Array((new BigNumber(3)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _335_v87;
          _nw54[(_dafny.ONE).toNumber()] = _335_v87;
          _nw54[(new BigNumber(2)).toNumber()] = _335_v87;
          _342_v93 = _nw54;
          let _343_v94;
          _343_v94 = _dafny.Set.fromElements(_336_v88);
          let _344_v95;
          _344_v95 = _module.D2.create_DC7(_343_v94);
          let _345_v96;
          _345_v96 = _dafny.Map.Empty.slice().updateUnsafe(_342_v93,((_344_v95).dtor_cf12).IsDisjointFrom(_343_v94));
          (_246_globalState).f14 = (((_345_v96).contains(_342_v93)) ? ((_345_v96).get(_342_v93)) : ((_301_v55).f29));
        } else {
          let _346_v97;
          _346_v97 = _dafny.Map.Empty.slice().updateUnsafe(_339_v90,_339_v90);
          _339_v90 = (((_346_v97).contains(_301_v55)) ? ((_346_v97).get(_301_v55)) : (_339_v90));
          let _347_v98;
          let _nw55 = new _module.C0();
          _nw55.__ctor((_301_v55).f29);
          _347_v98 = _nw55;
          let _348_v99;
          let _349_v100;
          let _350_v101;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector7 = _module.__default.m0(_dafny.Seq.UnicodeFromString("dkmvx"), _246_globalState);
          _out21 = _outcollector7[0];
          _out22 = _outcollector7[1];
          _out23 = _outcollector7[2];
          _348_v99 = _out21;
          _349_v100 = _out22;
          _350_v101 = _out23;
          (_246_globalState).f14 = (_230_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(_233_v7), new BigNumber((_230_v4).length))];
          let _351_v102;
          let _init5 = ((_352_v3) => function (_353_i10) {
            return (_352_v3)[_module.__default.safeIndex(new BigNumber(832), new BigNumber((_352_v3).length))];
          })(_229_v3);
          let _nw56 = Array((new BigNumber(24)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw56.length); _i0_5++) {
            _nw56[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _351_v102 = _nw56;
          let _index35 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_351_v102).length));
          (_351_v102)[_index35] = _243_v16;
          let _index36 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_351_v102).length));
          let _rhs35 = _dafny.Seq.Concat(_230_v4, _230_v4);
          let _rhs36 = _module.__default.fm7((((_301_v55).f29) ? (_229_v3) : (_253_v21)), (_dafny.ZERO).minus(_233_v7), _246_globalState);
          let _lhs23 = _351_v102;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_351_v102).length));
          _230_v4 = _rhs35;
          _lhs23[_lhs24] = _rhs36;
        }
        (_246_globalState).f13 = ((_dafny.ZERO).minus((_255_v23).multipliedBy(new BigNumber(-758)))).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), ((_354_v7) => function (_355_i11) {
          return _354_v7;
        })(_233_v7))).length));
        let _356_v103;
        _356_v103 = _dafny.MultiSet.fromElements(_233_v7, new BigNumber(-464), _233_v7, _255_v23);
        if ((_356_v103).IsProperSubsetOf(_356_v103)) {
          _233_v7 = _233_v7;
          (_246_globalState).f14 = (_301_v55).f29;
          (_246_globalState).f17 = _255_v23;
          let _357_v104;
          _357_v104 = _dafny.Seq.of(_233_v7);
          let _358_v105;
          _358_v105 = _dafny.Map.Empty.slice().updateUnsafe(_357_v104,(_301_v55).f29);
          let _359_v106;
          _359_v106 = _dafny.Seq.of(_module.__default.fm2(_dafny.MultiSet.FromArray(_229_v3), new BigNumber((_358_v105).length), _255_v23, _226_v0, _246_globalState));
          let _index37 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_235_v9).length));
          (_235_v9)[_index37] = _module.__default.safeModuloInt(_255_v23, new BigNumber((_359_v106).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_235_v9).length));
          (_235_v9)[_index38] = _255_v23;
          let _index39 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_235_v9).length));
          (_235_v9)[_index39] = new BigNumber((_359_v106).length);
        } else {
          let _360_v107;
          _360_v107 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC10(_233_v7, _243_v16),_dafny.MultiSet.fromElements((_339_v90).f29, _226_v0, (_339_v90).f29, (_301_v55).f29, (_301_v55).f29));
          let _361_v108;
          _361_v108 = _module.D2.create_DC10(_255_v23, _243_v16);
          _226_v0 = !((((_360_v107).contains(_361_v108)) ? ((_360_v107).get(_361_v108)) : (_227_v1))).contains((_301_v55).f29);
          _253_v21 = _229_v3;
          let _362_v109;
          let _nw57 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _362_v109 = _nw57;
          _362_v109 = _362_v109;
          let _363_v110;
          let _nw58 = Array((_dafny.ONE).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _235_v9;
          _363_v110 = _nw58;
          let _index40 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_363_v110).length));
          (_363_v110)[_index40] = _254_v22;
          let _index41 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_239_v14).length));
          (_239_v14)[_index41] = _226_v0;
          let _364_v111;
          _364_v111 = _dafny.Map.Empty.slice().updateUnsafe((_301_v55).f29,new BigNumber((_227_v1).cardinality()));
          let _365_v112;
          _365_v112 = _dafny.Map.Empty.slice().updateUnsafe(_233_v7,_364_v111);
          let _index42 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_363_v110).length));
          let _index43 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_239_v14).length));
          let _rhs37 = _235_v9;
          let _rhs38 = _module.__default.fm0(_229_v3, ((((_365_v112).contains(_255_v23)) ? ((_365_v112).get(_255_v23)) : (_364_v111))).contains((_339_v90).f29), _246_globalState);
          let _rhs39 = (_301_v55).f29;
          let _lhs25 = _363_v110;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_363_v110).length));
          let _lhs27 = _239_v14;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_239_v14).length));
          let _lhs29 = _246_globalState;
          _lhs25[_lhs26] = _rhs37;
          _lhs27[_lhs28] = _rhs38;
          _lhs29.f14 = _rhs39;
          let _366_v113;
          let _367_v114;
          let _368_v115;
          let _out24;
          let _out25;
          let _out26;
          let _outcollector8 = _module.__default.m0(_dafny.Seq.Concat(_229_v3, _253_v21), _246_globalState);
          _out24 = _outcollector8[0];
          _out25 = _outcollector8[1];
          _out26 = _outcollector8[2];
          _366_v113 = _out24;
          _367_v114 = _out25;
          _368_v115 = _out26;
        }
      }
      if (!((_255_v23).isLessThanOrEqualTo(_255_v23))) {
        let _369_v116;
        let _nw59 = new _module.C0();
        _nw59.__ctor(true);
        _369_v116 = _nw59;
        _226_v0 = _226_v0;
        (_246_globalState).f14 = (_301_v55).f29;
        let _370_v117;
        _370_v117 = _dafny.Set.fromElements(_233_v7, _233_v7);
        let _371_v118;
        _371_v118 = _dafny.Set.fromElements(_370_v117);
        let _372_v119;
        _372_v119 = _module.D2.create_DC7(_371_v118);
        let _pat_let_tv3 = _371_v118;
        let _pat_let_tv4 = _371_v118;
        _372_v119 = function (_pat_let4_0) {
          return function (_373_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_374_dt__update_hcf12_h0) {
                return _module.D2.create_DC7(_374_dt__update_hcf12_h0);
              }(_pat_let5_0);
            }((_pat_let_tv3).Difference(_pat_let_tv4));
          }(_pat_let4_0);
        }(_module.D2.create_DC7(_371_v118));
        let _375_v120;
        _375_v120 = _dafny.Map.Empty.slice().updateUnsafe(_243_v16,(_369_v116).f29);
        let _rhs40 = _255_v23;
        let _rhs41 = _375_v120;
        let _rhs42 = _233_v7;
        let _lhs30 = _246_globalState;
        let _lhs31 = _246_globalState;
        _lhs30.f17 = _rhs40;
        _375_v120 = _rhs41;
        _lhs31.f27 = _rhs42;
      } else {
        let _376_v121;
        let _377_v122;
        let _378_v123;
        let _out27;
        let _out28;
        let _out29;
        let _outcollector9 = _module.__default.m0(_dafny.Seq.UnicodeFromString("yqs"), _246_globalState);
        _out27 = _outcollector9[0];
        _out28 = _outcollector9[1];
        _out29 = _outcollector9[2];
        _376_v121 = _out27;
        _377_v122 = _out28;
        _378_v123 = _out29;
        let _index44 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_239_v14).length));
        (_239_v14)[_index44] = (_301_v55).f29;
        let _index45 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_239_v14).length));
        (_239_v14)[_index45] = (_301_v55).f29;
        let _379_v124;
        let _nw60 = new _module.C0();
        _nw60.__ctor(_226_v0);
        _379_v124 = _nw60;
        let _index46 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_239_v14).length));
        (_239_v14)[_index46] = _226_v0;
        let _380_v125;
        _380_v125 = _dafny.Map.Empty.slice().updateUnsafe(_226_v0,_376_v121);
        let _381_v126;
        _381_v126 = _dafny.MultiSet.fromElements(_255_v23, _233_v7, new BigNumber(((((_380_v125).contains((_301_v55).f29)) ? ((_380_v125).get((_301_v55).f29)) : (_253_v21))).length), _module.__default.safeModuloInt(_233_v7, _378_v123));
        let _382_v127;
        _382_v127 = _dafny.Map.Empty.slice().updateUnsafe(_255_v23,_dafny.MultiSet.fromElements(_378_v123, _255_v23, _233_v7, _233_v7, new BigNumber(101)));
        _381_v126 = ((!((_239_v14)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_239_v14).length))])) ? (_381_v126) : ((((_379_v124).f29) ? (_381_v126) : ((((_382_v127).contains(_378_v123)) ? ((_382_v127).get(_378_v123)) : (_381_v126))))));
      }
      let _383_v128;
      let _384_v129;
      let _385_v130;
      let _out30;
      let _out31;
      let _out32;
      let _outcollector10 = _module.__default.m0(_253_v21, _246_globalState);
      _out30 = _outcollector10[0];
      _out31 = _outcollector10[1];
      _out32 = _outcollector10[2];
      _383_v128 = _out30;
      _384_v129 = _out31;
      _385_v130 = _out32;
      _301_v55 = ((false) ? (_301_v55) : (_301_v55));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_239_v14).length))) {
        let _386_i12 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_386_i12)) && ((_386_i12).isLessThan(new BigNumber((_239_v14).length))))) {
          (_239_v14)[(_386_i12)] = (_301_v55).f29;
        }
      }
      if (_226_v0) {
        let _387_v131;
        let _nw61 = Array((new BigNumber(29)).toNumber()).fill([]);
        _387_v131 = _nw61;
        _387_v131 = ((_226_v0) ? (_387_v131) : (_387_v131));
        (_246_globalState).f13 = _module.__default.safeDivisionInt(_385_v130, _module.__default.safeModuloInt(_255_v23, _233_v7));
        (_246_globalState).f27 = _385_v130;
        let _388_v132;
        let _389_v133;
        let _390_v134;
        let _out33;
        let _out34;
        let _out35;
        let _outcollector11 = _module.__default.m0(_253_v21, _246_globalState);
        _out33 = _outcollector11[0];
        _out34 = _outcollector11[1];
        _out35 = _outcollector11[2];
        _388_v132 = _out33;
        _389_v133 = _out34;
        _390_v134 = _out35;
        let _391_v135;
        _391_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_385_v130,(_301_v55).f29)).length),(_dafny.ZERO).minus(_233_v7));
        let _392_v136;
        _392_v136 = _dafny.Set.fromElements(_233_v7, (_dafny.ZERO).minus(new BigNumber((_391_v135).length)));
        let _393_v137;
        _393_v137 = _dafny.Map.Empty.slice().updateUnsafe(_255_v23,new BigNumber((_392_v136).length));
        let _index47 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_235_v9).length));
        (_235_v9)[_index47] = (((_393_v137).contains(_255_v23)) ? ((_393_v137).get(_255_v23)) : (new BigNumber(-935)));
        let _index48 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_235_v9).length));
        let _rhs43 = _226_v0;
        let _rhs44 = new BigNumber(-171);
        let _rhs45 = _module.__default.safeDivisionInt(_255_v23, _390_v134);
        let _lhs32 = _235_v9;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_235_v9).length));
        _226_v0 = _rhs43;
        _385_v130 = _rhs44;
        _lhs32[_lhs33] = _rhs45;
      } else {
        let _394_v138;
        _394_v138 = _dafny.Map.Empty.slice().updateUnsafe(_243_v16,_226_v0);
        let _395_v139;
        _395_v139 = _dafny.Seq.of(_233_v7, new BigNumber((_dafny.Set.fromElements(new BigNumber(395), _233_v7)).length), new BigNumber((_394_v138).length), new BigNumber((_237_v12).length), _255_v23);
        let _396_v140;
        _396_v140 = _dafny.Set.fromElements(_395_v139);
        let _397_v141;
        _397_v141 = _dafny.MultiSet.fromElements(_395_v139, _395_v139);
        let _398_v143;
        let _nw62 = Array((new BigNumber(13)).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = _396_v140;
        _nw62[(_dafny.ONE).toNumber()] = _396_v140;
        _nw62[(new BigNumber(2)).toNumber()] = (_396_v140).Intersect(_396_v140);
        _nw62[(new BigNumber(3)).toNumber()] = _module.__default.fm8(_233_v7, true, (_301_v55).f29, (_301_v55).f29, _246_globalState);
        _nw62[(new BigNumber(4)).toNumber()] = (_396_v140).Union(_396_v140);
        _nw62[(new BigNumber(5)).toNumber()] = function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of (_397_v141).Elements) {
            let _399_v142 = _compr_4;
            if ((_397_v141).contains(_399_v142)) {
              _coll4.add(_399_v142);
            }
          }
          return _coll4;
        }();
        _nw62[(new BigNumber(6)).toNumber()] = _396_v140;
        _nw62[(new BigNumber(7)).toNumber()] = _396_v140;
        _nw62[(new BigNumber(8)).toNumber()] = _396_v140;
        _nw62[(new BigNumber(9)).toNumber()] = _396_v140;
        _nw62[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_395_v139, _395_v139);
        _nw62[(new BigNumber(11)).toNumber()] = (_396_v140).Union(_dafny.Set.fromElements(_395_v139, _395_v139));
        _nw62[(new BigNumber(12)).toNumber()] = _396_v140;
        _398_v143 = _nw62;
        let _index49 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_398_v143).length));
        (_398_v143)[_index49] = _396_v140;
        let _index50 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_398_v143).length));
        let _rhs46 = !(_226_v0);
        let _rhs47 = (_396_v140).Intersect(_module.__default.fm8(_255_v23, false, (_301_v55).f29, _226_v0, _246_globalState));
        let _lhs34 = _246_globalState;
        let _lhs35 = _398_v143;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_398_v143).length));
        _lhs34.f14 = _rhs46;
        _lhs35[_lhs36] = _rhs47;
        _255_v23 = _255_v23;
        if ((_230_v4)[_module.__default.safeIndex(_385_v130, new BigNumber((_230_v4).length))]) {
          let _400_v144;
          _400_v144 = _module.D1.create_DC4((_301_v55).f29, _385_v130);
          _226_v0 = (_400_v144).dtor_cf5;
          let _401_v145;
          _401_v145 = _234_v8;
          let _402_v146;
          _402_v146 = _dafny.Seq.of(_dafny.Set.fromElements((_301_v55).f29, (_301_v55).f29), _module.__default.fm9(_255_v23, _246_globalState));
          let _403_v147;
          let _nw63 = Array((new BigNumber(18)).toNumber());
          _nw63[(_dafny.ZERO).toNumber()] = _234_v8;
          _nw63[(_dafny.ONE).toNumber()] = (((_301_v55).f29) ? (_dafny.Set.fromElements(_226_v0, (_400_v144).dtor_cf5)) : (_234_v8));
          _nw63[(new BigNumber(2)).toNumber()] = ((_226_v0) ? (_234_v8) : (_234_v8));
          _nw63[(new BigNumber(3)).toNumber()] = _module.__default.fm9(_255_v23, _246_globalState);
          _nw63[(new BigNumber(4)).toNumber()] = (_401_v145);
          _nw63[(new BigNumber(5)).toNumber()] = _234_v8;
          _nw63[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_226_v0, (_301_v55).f29, (_301_v55).f29, (_301_v55).f29);
          _nw63[(new BigNumber(7)).toNumber()] = (_234_v8).Intersect(_234_v8);
          _nw63[(new BigNumber(8)).toNumber()] = (_dafny.Set.fromElements((_301_v55).f29)).Intersect(_dafny.Set.fromElements((_301_v55).f29, (_301_v55).f29));
          _nw63[(new BigNumber(9)).toNumber()] = _234_v8;
          _nw63[(new BigNumber(10)).toNumber()] = (_402_v146)[_module.__default.safeIndex(_233_v7, new BigNumber((_402_v146).length))];
          _nw63[(new BigNumber(11)).toNumber()] = (_dafny.Set.fromElements(_226_v0, !((_301_v55).f29))).Difference(_module.__default.fm9(new BigNumber(629), _246_globalState));
          _nw63[(new BigNumber(12)).toNumber()] = _234_v8;
          _nw63[(new BigNumber(13)).toNumber()] = (_234_v8).Difference(_234_v8);
          _nw63[(new BigNumber(14)).toNumber()] = _234_v8;
          _nw63[(new BigNumber(15)).toNumber()] = _234_v8;
          _nw63[(new BigNumber(16)).toNumber()] = _234_v8;
          _nw63[(new BigNumber(17)).toNumber()] = (_234_v8).Union(_234_v8);
          _403_v147 = _nw63;
          let _404_v148;
          _404_v148 = _module.D1.create_DC5(_233_v7, (_301_v55).f29, _385_v130, (_301_v55).f29);
          let _405_v149;
          _405_v149 = _module.D1.create_DC6(_404_v148);
          let _406_v150;
          _406_v150 = _module.D2.create_DC9(_226_v0, _301_v55);
          let _rhs48 = _403_v147;
          let _rhs49 = !((_237_v12).Merge(_237_v12)).contains(true);
          let _rhs50 = _module.__default.fm10(_226_v0, new BigNumber(620), (_301_v55).f29, _246_globalState);
          let _rhs51 = (_406_v150).dtor_cf14;
          let _lhs37 = _246_globalState;
          let _lhs38 = _246_globalState;
          _403_v147 = _rhs48;
          _lhs37.f14 = _rhs49;
          _405_v149 = _rhs50;
          _lhs38.f14 = _rhs51;
          let _index51 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_235_v9).length));
          (_235_v9)[_index51] = _255_v23;
          let _index52 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_235_v9).length));
          (_235_v9)[_index52] = _385_v130;
          (_246_globalState).f14 = (_301_v55).f29;
          _226_v0 = (_233_v7).isLessThanOrEqualTo(_233_v7);
        } else {
          let _407_v151;
          let _nw64 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _407_v151 = _nw64;
          let _408_v152;
          _408_v152 = _dafny.MultiSet.fromElements(_255_v23);
          let _409_v153;
          _409_v153 = _module.D2.create_DC9((_301_v55).f29, _301_v55);
          let _410_v154;
          _410_v154 = _module.D2.create_DC9((_301_v55).f29, (_409_v153).dtor_cf15);
          let _411_v155;
          _411_v155 = _dafny.Map.Empty.slice().updateUnsafe(!((_module.__default.fm11(_385_v130, _408_v152, (_301_v55).f29, _246_globalState)).dtor_cf5),_409_v153);
          let _index53 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_407_v151).length));
          (_407_v151)[_index53] = _411_v155;
          let _index54 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_407_v151).length));
          (_407_v151)[_index54] = _411_v155;
          _253_v21 = _dafny.Seq.Concat(_dafny.Seq.update(_253_v21, _module.__default.safeIndex(_255_v23, new BigNumber((_253_v21).length)), _243_v16), _383_v128);
          let _412_v156;
          let _nw65 = new _module.C0();
          _nw65.__ctor(_226_v0);
          _412_v156 = _nw65;
          let _index55 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_254_v22).length));
          (_254_v22)[_index55] = _233_v7;
          let _413_v157;
          let _nw66 = Array((new BigNumber(6)).toNumber()).fill([]);
          _413_v157 = _nw66;
          let _index56 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_413_v157).length));
          (_413_v157)[_index56] = _239_v14;
          let _414_v158;
          _414_v158 = _dafny.Set.fromElements(_233_v7, new BigNumber((_237_v12).length), new BigNumber(-263), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_395_v139, _395_v139)).length)), new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(_255_v23))).length));
          let _415_v159;
          _415_v159 = _dafny.Map.Empty.slice().updateUnsafe((_412_v156).f29,_254_v22);
          let _index57 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_254_v22).length));
          let _index58 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_413_v157).length));
          let _rhs52 = (_dafny.ZERO).minus((_385_v130).multipliedBy(new BigNumber((_414_v158).length)));
          let _rhs53 = (((_415_v159).contains(!(new BigNumber((_253_v21).length)).isEqualTo(_233_v7))) ? ((_415_v159).get(!(new BigNumber((_253_v21).length)).isEqualTo(_233_v7))) : (_254_v22));
          let _rhs54 = _409_v153;
          let _rhs55 = _239_v14;
          let _rhs56 = _module.__default.fm0(_229_v3, (_301_v55).f29, _246_globalState);
          let _lhs39 = _254_v22;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_254_v22).length));
          let _lhs41 = _246_globalState;
          let _lhs42 = _413_v157;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_413_v157).length));
          let _lhs44 = _246_globalState;
          _lhs39[_lhs40] = _rhs52;
          _lhs41.f28 = _rhs53;
          _410_v154 = _rhs54;
          _lhs42[_lhs43] = _rhs55;
          _lhs44.f14 = _rhs56;
          (_246_globalState).f17 = _module.__default.safeDivisionInt(_233_v7, _233_v7);
        }
        let _416_v160;
        _416_v160 = _dafny.Set.fromElements(_255_v23);
        let _417_v161;
        _417_v161 = _module.D1.create_DC5(_233_v7, _226_v0, _module.__default.safeDivisionInt(new BigNumber((_416_v160).length), _255_v23), _226_v0);
        let _source7 = _417_v161;
        if (_source7.is_DC3) {
          let _418___mcc_h2 = (_source7).cf3;
          let _419___mcc_h3 = (_source7).cf4;
          let _420_cf4 = _419___mcc_h3;
          let _421_cf3 = _418___mcc_h2;
          (_246_globalState).f14 = !(_dafny.Seq.IsProperPrefixOf(_253_v21, _253_v21));
          let _422_v162;
          _422_v162 = _dafny.Seq.of(_229_v3, _dafny.Seq.UnicodeFromString("msee"));
          let _423_v163;
          _423_v163 = _dafny.Set.fromElements(_384_v129);
          let _424_v164;
          _424_v164 = _module.D4.create_DC12(_235_v9);
          let _rhs57 = (_255_v23).minus(new BigNumber(((_422_v162)[_module.__default.safeIndex(_255_v23, new BigNumber((_422_v162).length))]).length));
          let _rhs58 = (_423_v163).IsProperSubsetOf((_423_v163).Difference(_dafny.Set.fromElements((_424_v164).dtor_cf19, _235_v9, _384_v129)));
          let _lhs45 = _246_globalState;
          let _lhs46 = _246_globalState;
          _lhs45.f13 = _rhs57;
          _lhs46.f14 = _rhs58;
          let _425_v165;
          _425_v165 = _dafny.MultiSet.fromElements(_230_v4);
          _425_v165 = ((_425_v165).update(_230_v4, _module.__default.abs(_255_v23))).Difference(_425_v165);
          let _rhs59 = (_301_v55).f29;
          let _rhs60 = _dafny.Seq.Concat(_230_v4, _dafny.Seq.Concat(_230_v4, _230_v4));
          let _lhs47 = _246_globalState;
          _lhs47.f14 = _rhs59;
          _230_v4 = _rhs60;
        } else if (_source7.is_DC4) {
          let _426___mcc_h4 = (_source7).cf5;
          let _427___mcc_h5 = (_source7).cf6;
          let _428_cf6 = _427___mcc_h5;
          let _429_cf5 = _426___mcc_h4;
          let _430_v166;
          _430_v166 = _dafny.Seq.of(_229_v3);
          (_246_globalState).f14 = !(!(!(!_dafny.Seq.contains(_430_v166, _383_v128))));
          (_246_globalState).f27 = (_395_v139)[_module.__default.safeIndex(_255_v23, new BigNumber((_395_v139).length))];
          let _431_v167;
          let _nw67 = new _module.C0();
          _nw67.__ctor(_429_cf5);
          _431_v167 = _nw67;
          let _432_v168;
          let _nw68 = Array((new BigNumber(14)).toNumber()).fill([]);
          _432_v168 = _nw68;
          let _433_v169;
          _433_v169 = _module.D5.create_DC14(_432_v168);
          let _434_v170;
          let _nw69 = Array((new BigNumber(28)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _432_v168;
          _nw69[(_dafny.ONE).toNumber()] = _432_v168;
          _nw69[(new BigNumber(2)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(3)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(4)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(5)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(6)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(7)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(8)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(9)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(10)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(11)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(12)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(13)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(14)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(15)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(16)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(17)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(18)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(19)).toNumber()] = (((_301_v55).f29) ? (_432_v168) : (_432_v168));
          _nw69[(new BigNumber(20)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(21)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(22)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(23)).toNumber()] = (_433_v169).dtor_cf24;
          _nw69[(new BigNumber(24)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(25)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(26)).toNumber()] = _432_v168;
          _nw69[(new BigNumber(27)).toNumber()] = _432_v168;
          _434_v170 = _nw69;
          let _index59 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_434_v170).length));
          (_434_v170)[_index59] = _432_v168;
          let _index60 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_434_v170).length));
          (_434_v170)[_index60] = (((_301_v55).f29) ? (_432_v168) : (_432_v168));
        } else if (_source7.is_DC5) {
          let _435___mcc_h6 = (_source7).cf7;
          let _436___mcc_h7 = (_source7).cf8;
          let _437___mcc_h8 = (_source7).cf9;
          let _438___mcc_h9 = (_source7).cf10;
          let _439_cf10 = _438___mcc_h9;
          let _440_cf9 = _437___mcc_h8;
          let _441_cf8 = _436___mcc_h7;
          let _442_cf7 = _435___mcc_h6;
          let _443_v171;
          let _nw70 = Array((_dafny.ONE).toNumber());
          _nw70[(_dafny.ZERO).toNumber()] = _301_v55;
          _443_v171 = _nw70;
          let _index61 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_443_v171).length));
          (_443_v171)[_index61] = _301_v55;
          let _index62 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_443_v171).length));
          (_443_v171)[_index62] = _301_v55;
          let _444_v172;
          _444_v172 = _module.D1.create_DC2(_253_v21);
          let _445_v173;
          let _446_v174;
          let _447_v175;
          let _out36;
          let _out37;
          let _out38;
          let _outcollector12 = _module.__default.m0((_444_v172).dtor_cf2, _246_globalState);
          _out36 = _outcollector12[0];
          _out37 = _outcollector12[1];
          _out38 = _outcollector12[2];
          _445_v173 = _out36;
          _446_v174 = _out37;
          _447_v175 = _out38;
          (_246_globalState).f17 = (_395_v139)[_module.__default.safeIndex(_233_v7, new BigNumber((_395_v139).length))];
          let _448_v176;
          _448_v176 = _dafny.MultiSet.fromElements(_243_v16, _243_v16);
          let _449_v177;
          _449_v177 = _dafny.MultiSet.fromElements(_module.__default.fm2((_448_v176).update(_243_v16, _module.__default.abs(_233_v7)), _440_cf9, _447_v175, (_301_v55).f29, _246_globalState), new BigNumber((_395_v139).length), _440_cf9);
          _449_v177 = _449_v177;
        } else if (_source7.is_DC2) {
          let _450___mcc_h10 = (_source7).cf2;
          let _451_cf2 = _450___mcc_h10;
          let _452_v178;
          let _453_v179;
          let _454_v180;
          let _out39;
          let _out40;
          let _out41;
          let _outcollector13 = _module.__default.m0(_dafny.Seq.update(_451_cf2, _module.__default.safeIndex(new BigNumber(-918), new BigNumber((_451_cf2).length)), (_383_v128)[_module.__default.safeIndex(new BigNumber(-862), new BigNumber((_383_v128).length))]), _246_globalState);
          _out39 = _outcollector13[0];
          _out40 = _outcollector13[1];
          _out41 = _outcollector13[2];
          _452_v178 = _out39;
          _453_v179 = _out40;
          _454_v180 = _out41;
          (_246_globalState).f11 = ((_dafny.MultiSet.fromElements((_301_v55).f29)).Union((_227_v1).update(_226_v0, _module.__default.abs(_454_v180)))).Intersect(_227_v1);
          let _index63 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_239_v14).length));
          (_239_v14)[_index63] = (_301_v55).f29;
          let _index64 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_239_v14).length));
          (_239_v14)[_index64] = !(_226_v0);
          let _455_v181;
          _455_v181 = _dafny.MultiSet.fromElements(new BigNumber(-929), _233_v7);
          let _456_v182;
          _456_v182 = _module.D6.create_DC16(_455_v181);
          let _index65 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_239_v14).length));
          let _rhs61 = _301_v55;
          let _rhs62 = (((_239_v14)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((_239_v14).length))]) ? (_dafny.Seq.update(_230_v4, _module.__default.safeIndex(_385_v130, new BigNumber((_230_v4).length)), true)) : (_230_v4));
          let _rhs63 = !((_455_v181).IsSubsetOf((_456_v182).dtor_cf30));
          let _rhs64 = _230_v4;
          let _lhs48 = _239_v14;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_239_v14).length));
          _301_v55 = _rhs61;
          _230_v4 = _rhs62;
          _lhs48[_lhs49] = _rhs63;
          _230_v4 = _rhs64;
        } else {
          let _457___mcc_h11 = (_source7).cf11;
          let _458_cf11 = _457___mcc_h11;
          let _459_v183;
          let _460_v184;
          let _461_v185;
          let _out42;
          let _out43;
          let _out44;
          let _outcollector14 = _module.__default.m0(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ynbc"), _module.__default.safeIndex(_233_v7, new BigNumber((_dafny.Seq.UnicodeFromString("ynbc")).length)), _243_v16), _246_globalState);
          _out42 = _outcollector14[0];
          _out43 = _outcollector14[1];
          _out44 = _outcollector14[2];
          _459_v183 = _out42;
          _460_v184 = _out43;
          _461_v185 = _out44;
          _242_v15 = _242_v15;
          let _462_v186;
          let _463_v187;
          let _464_v188;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector15 = _module.__default.m0(_dafny.Seq.Concat(_253_v21, _253_v21), _246_globalState);
          _out45 = _outcollector15[0];
          _out46 = _outcollector15[1];
          _out47 = _outcollector15[2];
          _462_v186 = _out45;
          _463_v187 = _out46;
          _464_v188 = _out47;
          let _465_v189;
          _465_v189 = _dafny.Map.Empty.slice().updateUnsafe(_461_v185,_384_v129);
          (_246_globalState).f14 = _module.__default.fm0(_dafny.Seq.Concat(_459_v183, _module.__default.fm12(_233_v7, new BigNumber((_465_v189).length), (_301_v55).f29, _246_globalState)), (_301_v55).f29, _246_globalState);
        }
        (_246_globalState).f27 = (_dafny.ZERO).minus((_385_v130).multipliedBy(new BigNumber((_395_v139).length)));
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_239_v14).length))) {
        let _466_i13 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_466_i13)) && ((_466_i13).isLessThan(new BigNumber((_239_v14).length))))) {
          (_239_v14)[(_466_i13)] = (_301_v55).f29;
        }
      }
      _226_v0 = (_301_v55).f29;
      process.stdout.write(_dafny.toString(_226_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v1).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_228_v2, _dafny.Seq.of(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_229_v3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_230_v4, _dafny.Seq.of(true, true, true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_231_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(true, true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_233_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v8).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_236_v10).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false).updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_238_v13, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v14)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(155),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_243_v16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v17).equals(_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_245_v18).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(155),true),_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_246_globalState).f0, _dafny.Seq.of(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_246_globalState).f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f7).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(true, true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_246_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f11).equals(_dafny.MultiSet.fromElements(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_246_globalState).f18).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f19).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,false),new BigNumber(773)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f20)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f24).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(155),true),_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_globalState.f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f26).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_globalState.f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f28)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_248_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_253_v21).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v22)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_255_v23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_280_v36)[new BigNumber(2)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_281_v37).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v55).f29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_383_v128).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v129)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_385_v130));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC1() {
      let $dt = new D0(1);
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC3(cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC4(cf5, cf6) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf7, cf8, cf9, cf10) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(3);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC6(cf11) {
      let $dt = new D1(4);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC2() { return this.$tag === 3; }
    get is_DC6() { return this.$tag === 4; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC2" + "(" + this.cf2.toVerbatimString(true) + ")";
      } else if (this.$tag === 4) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC8(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC9(cf14, cf15) {
      let $dt = new D2(1);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10(cf16, cf17) {
      let $dt = new D2(2);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D2(3);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + this.cf13.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf14 === other.cf14 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.Seq.UnicodeFromString(""));
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
    static create_DC11(cf18) {
      let $dt = new D3(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf18) + ")";
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
      return _dafny.Set.Empty;
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
    static create_DC13(cf20, cf21, cf22, cf23) {
      let $dt = new D4(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(null, new _dafny.CodePoint('D'.codePointAt(0)), new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Seq.of());
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
    static create_DC15(cf25, cf26, cf27, cf28, cf29) {
      let $dt = new D5(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC14(cf24) {
      let $dt = new D5(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + this.cf25.toVerbatimString(true) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && this.cf27 === other.cf27 && this.cf28 === other.cf28 && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_dafny.Seq.UnicodeFromString(""), false, false, false, []);
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
    static create_DC17(cf31, cf32, cf33) {
      let $dt = new D6(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC18(cf34, cf35, cf36, cf37) {
      let $dt = new D6(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC16(cf30) {
      let $dt = new D6(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + this.cf34.toVerbatimString(true) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && this.cf33 === other.cf33;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(new _dafny.CodePoint('D'.codePointAt(0)), false, null);
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
    static create_DC20(cf39, cf40, cf41, cf42) {
      let $dt = new D7(0);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC19(cf38) {
      let $dt = new D7(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC21(cf43) {
      let $dt = new D7(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf39 === other.cf39 && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(false, false, _dafny.ZERO, false);
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
    static create_DC23() {
      let $dt = new D8(0);
      return $dt;
    }
    static create_DC24(cf45) {
      let $dt = new D8(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC22(cf44) {
      let $dt = new D8(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC25(cf46) {
      let $dt = new D8(3);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf46) + ")";
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
        return other.$tag === 1 && this.cf45 === other.cf45;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23();
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
      this.f2 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f11 = _dafny.MultiSet.Empty;
      this.f13 = _dafny.ZERO;
      this.f14 = false;
      this.f17 = _dafny.ZERO;
      this.f25 = _dafny.ZERO;
      this.f27 = _dafny.ZERO;
      this.f28 = [];
      this._f0 = _dafny.Seq.of();
      this._f1 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.Seq.UnicodeFromString("");
      this._f7 = _dafny.Map.Empty;
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = [];
      this._f10 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.ZERO;
      this._f18 = _dafny.MultiSet.Empty;
      this._f19 = _dafny.Map.Empty;
      this._f20 = [];
      this._f21 = _dafny.ZERO;
      this._f22 = false;
      this._f23 = _dafny.ZERO;
      this._f24 = _dafny.Map.Empty;
      this._f26 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this).f25 = f25;
      (_this)._f26 = f26;
      (_this).f27 = f27;
      (_this).f28 = f28;
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
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f15() {
      let _this = this;
      return _this._f15;
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
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f29 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f29) {
      let _this = this;
      (_this)._f29 = f29;
      return;
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
