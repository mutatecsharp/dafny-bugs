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
    static fm0(p0, p1, p2, p3, globalState) {
      return new BigNumber(-249);
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.of((false) || (!(true)), !(new BigNumber(252)).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-707)), function (_0_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length)), (new BigNumber(176)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('m'.codePointAt(0)))).length)));
    };
    static fm2(p0, globalState) {
      return (_module.D0.create_DC1(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_1_i0) {
  return new BigNumber((_dafny.Seq.of(new BigNumber(594))).length);
})).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false)).length))).length))).cardinality()), new BigNumber(-495), false)).dtor_cf3;
    };
    static fm5(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), function (_2_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)))));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),!(!(true)))),_dafny.Seq.UnicodeFromString("raqnsyr"));
    };
    static fm7(p0, p1, globalState) {
      return new _dafny.CodePoint('m'.codePointAt(0));
    };
    static fm8(p0, p1, p2, globalState) {
      let _source0 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ixyg")).length),new BigNumber((_dafny.Seq.UnicodeFromString("auue")).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))).length),new BigNumber((_dafny.MultiSet.fromElements(true, false, false, false, (_module.D0.create_DC2(false, !(true), new BigNumber(362), true)).dtor_cf5)).cardinality())));
      let _3___mcc_h0 = _source0;
      let _4_cf35 = _3___mcc_h0;
      return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.of(new BigNumber(124), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Seq.of(true)).length))).length)));
    };
    static fm9(p0, p1, p2, globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_5_i0) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("gwyvjb"))).Elements) {
          let _6_v0 = _compr_0;
          if ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_5_i0) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("gwyvjb"))).contains(_6_v0)) {
            _coll0.add(_6_v0);
          }
        }
        return _coll0;
      }());
    };
    static fm10(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), function (_7_i0) {
        return new BigNumber(-230);
      }))).cardinality())), new BigNumber(-762)),((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kl")).length), new BigNumber(977), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-882))).cardinality())), new BigNumber(983))).length))).length))).plus(new BigNumber(993)));
    };
    static fm11(p0, globalState) {
      let _source1 = _module.D2.create_DC8();
      if (_source1.is_DC8) {
        return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(116), new BigNumber(-258), new BigNumber(647), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hklxydp")).length)))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D9.create_DC27(_dafny.Set.fromElements(new BigNumber(673)))).dtor_cf37).length),true)).length)));
      } else if (_source1.is_DC9) {
        let _8___mcc_h0 = (_source1).cf13;
        let _9___mcc_h1 = (_source1).cf14;
        let _10___mcc_h2 = (_source1).cf15;
        let _11_cf15 = _10___mcc_h2;
        let _12_cf14 = _9___mcc_h1;
        let _13_cf13 = _8___mcc_h0;
        return (_dafny.Set.fromElements(_11_cf15.f19, _13_cf13, _11_cf15.f19, (_dafny.ZERO).minus(_12_cf14.f19))).Union(_dafny.Set.fromElements(_11_cf15.f19, _11_cf15.f19));
      } else if (_source1.is_DC7) {
        let _14___mcc_h3 = (_source1).cf12;
        let _15_cf12 = _14___mcc_h3;
        return (function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_16_i0) {
            return new BigNumber(631);
          })).Elements) {
            let _17_v0 = _compr_1;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_16_i0) {
              return new BigNumber(631);
            }), _17_v0)) {
              _coll1.add(_module.__default.safeDivisionInt(_17_v0, new BigNumber((_dafny.Set.fromElements(true, true, false)).length)));
            }
          }
          return _coll1;
        }()).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("nxglsvqlj")).length)));
      } else {
        let _18___mcc_h4 = (_source1).cf16;
        let _19_cf16 = _18___mcc_h4;
        return _dafny.Set.fromElements(new BigNumber(134), new BigNumber(321));
      }
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("erpx")).length)), new BigNumber(36), (false) && (true));
    };
    static fm13(p0, globalState) {
      return _module.D5.create_DC19(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("bpwtnk"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_20_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})));
    };
    static fm14(p0, globalState) {
      return _dafny.Seq.of(new BigNumber(870));
    };
    static fm15(globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, !(false), !(true), true, !(false)))).length));
    };
    static fm16(globalState) {
      return _dafny.Set.fromElements((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("slbcnr")).length),!(false))).length),new BigNumber(953))).Keys.Elements) {
          let _21_v0 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("slbcnr")).length),!(false))).length),new BigNumber(953))).contains(_21_v0)) {
            _coll2.push([(_21_v0).plus(new BigNumber((_dafny.Seq.of(true, false)).length)),new BigNumber((_dafny.Set.fromElements(new BigNumber(-705), new BigNumber((_dafny.Seq.of(true)).length))).length)]);
          }
        }
        return _coll2;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("rihcohcjh")).length),new BigNumber((_dafny.Seq.of(true, !(true))).length))));
    };
    static fm17(globalState) {
      return _dafny.Set.fromElements(true, !(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-453)), function (_22_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)).isEqualTo(new BigNumber(70)), (function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Set.fromElements(_dafny.Seq.of(!(false)), _dafny.Seq.of(false), _dafny.Seq.of(false, false), _dafny.Seq.of(true, true), _dafny.Seq.of(true))).Elements) {
          let _23_v0 = _compr_3;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(!(false)), _dafny.Seq.of(false), _dafny.Seq.of(false, false), _dafny.Seq.of(true, true), _dafny.Seq.of(true))).contains(_23_v0)) {
            _coll3.push([_23_v0,_module.D0.create_DC0(new BigNumber(566))]);
          }
        }
        return _coll3;
      }()).equals(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(true))).Elements) {
          let _24_v1 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(true)), _24_v1)) {
            _coll4.push([_24_v1,_module.D0.create_DC0(new BigNumber(421))]);
          }
        }
        return _coll4;
      }()), (_dafny.MultiSet.fromElements(true, false, !(false))).IsDisjointFrom(_dafny.MultiSet.fromElements(false)));
    };
    static fm18(p0, p1, globalState) {
      let _source2 = ((false) ? (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(722),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("h")).length))).length), new BigNumber(-185))).length)))) : (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_25_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length),false)).length),new BigNumber(265)))));
      let _26___mcc_h0 = _source2;
      let _27_cf35 = _26___mcc_h0;
      return _module.D6.create_DC23((_module.D0.create_DC2(!(true), !(true), new BigNumber((_dafny.Seq.UnicodeFromString("wmttkoul")).length), false)).dtor_cf4, new BigNumber(136), new BigNumber(541));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.ZERO;
      let r3 = _dafny.MultiSet.Empty;
      let _28_v0;
      _28_v0 = _dafny.MultiSet.fromElements(p0, p0, false);
      let _29_v1;
      _29_v1 = _dafny.Seq.of(_28_v0);
      let _30_v2;
      _30_v2 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(_28_v0, _28_v0), _29_v1), _dafny.Seq.Concat(_29_v1, _29_v1));
      let _31_v3;
      _31_v3 = _dafny.Seq.UnicodeFromString("rg");
      let _pat_let_tv0 = p2;
      let _pat_let_tv1 = p2;
      let _rhs0 = new BigNumber(((_30_v2)[_module.__default.safeIndex(p2, new BigNumber((_30_v2).length))]).length);
      let _rhs1 = _31_v3;
      let _rhs2 = function (_source3) {
        if (_source3.is_DC15) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("y")).length);
        } else if (_source3.is_DC14) {
          let _32___mcc_h0 = (_source3).cf21;
          let _33_cf21 = _32___mcc_h0;
          return _pat_let_tv0;
        } else {
          let _34___mcc_h1 = (_source3).cf22;
          let _35_cf22 = _34___mcc_h1;
          return _pat_let_tv1;
        }
      }(_module.D4.create_DC15());
      let _rhs3 = p1;
      let _lhs0 = globalState;
      let _lhs1 = globalState;
      let _lhs2 = globalState;
      _lhs0.f12 = _rhs0;
      r1 = _rhs1;
      _lhs1.f2 = _rhs2;
      _lhs2.f3 = _rhs3;
      let _36_v4;
      _36_v4 = new _dafny.CodePoint('b'.codePointAt(0));
      let _37_v5;
      _37_v5 = _module.D5.create_DC17(_36_v4);
      let _pat_let_tv2 = p2;
      let _pat_let_tv3 = p2;
      let _pat_let_tv4 = p2;
      let _pat_let_tv5 = p2;
      (globalState).f2 = function (_source4) {
        if (_source4.is_DC18) {
          let _38___mcc_h2 = (_source4).cf24;
          let _39___mcc_h3 = (_source4).cf25;
          let _40___mcc_h4 = (_source4).cf26;
          let _41___mcc_h5 = (_source4).cf27;
          let _42_cf27 = _41___mcc_h5;
          let _43_cf26 = _40___mcc_h4;
          let _44_cf25 = _39___mcc_h3;
          let _45_cf24 = _38___mcc_h2;
          return _pat_let_tv2;
        } else if (_source4.is_DC19) {
          let _46___mcc_h6 = (_source4).cf28;
          let _47_cf28 = _46___mcc_h6;
          return _pat_let_tv3;
        } else if (_source4.is_DC17) {
          let _48___mcc_h7 = (_source4).cf23;
          let _49_cf23 = _48___mcc_h7;
          return _pat_let_tv4;
        } else {
          let _50___mcc_h8 = (_source4).cf29;
          let _51_cf29 = _50___mcc_h8;
          return (_dafny.ZERO).minus(_pat_let_tv5);
        }
      }(_37_v5);
      if ((p1) && (p1)) {
        let _52_v6;
        let _init0 = ((_53_p2) => function (_54_i0) {
          return _module.__default.safeModuloInt(_54_i0, _53_p2);
        })(p2);
        let _nw0 = Array((new BigNumber(6)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _52_v6 = _nw0;
        let _55_v7;
        _55_v7 = _dafny.Seq.of(p0);
        let _56_v8;
        _56_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,_55_v7);
        let _index0 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_52_v6).length));
        (_52_v6)[_index0] = new BigNumber((_56_v8).length);
        let _57_v9;
        _57_v9 = _dafny.Seq.of(p2);
        let _index1 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_52_v6).length));
        (_52_v6)[_index1] = _module.__default.safeDivisionInt(new BigNumber((_57_v9).length), p2);
        let _58_v10;
        _58_v10 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), ((_59_p2) => function (_60_i2) {
          return _59_p2;
        })(p2)));
        let _61_v11;
        _61_v11 = _dafny.Seq.of(_58_v10, _58_v10);
        let _62_v12;
        _62_v12 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(-913));
        let _index2 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_52_v6).length));
        let _index3 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_52_v6).length));
        let _rhs4 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(812)), function (_63_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), _31_v3);
        let _rhs5 = !_dafny.Seq.contains(_57_v9, p2);
        let _rhs6 = (new BigNumber(((_61_v11)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_61_v11).length))]).length)).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_62_v12).length), (_57_v9)[_module.__default.safeIndex(p2, new BigNumber((_57_v9).length))])));
        let _rhs7 = p2;
        let _lhs3 = globalState;
        let _lhs4 = _52_v6;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_52_v6).length));
        let _lhs6 = _52_v6;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_52_v6).length));
        r1 = _rhs4;
        _lhs3.f3 = _rhs5;
        _lhs4[_lhs5] = _rhs6;
        _lhs6[_lhs7] = _rhs7;
        _36_v4 = _36_v4;
        (globalState).f3 = p0;
        let _index4 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_52_v6).length));
        (_52_v6)[_index4] = _module.__default.safeModuloInt(_module.__default.safeModuloInt(p2, new BigNumber((_31_v3).length)), new BigNumber((_dafny.Seq.UnicodeFromString("bgihoex")).length));
        let _64_v13;
        let _nw1 = Array((_dafny.ONE).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = p1;
        _64_v13 = _nw1;
        let _index5 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_64_v13).length));
        (_64_v13)[_index5] = _module.__default.fm2(p0, globalState);
        let _index6 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_64_v13).length));
        (_64_v13)[_index6] = _module.__default.fm2((_55_v7)[_module.__default.safeIndex((_52_v6)[_module.__default.safeIndex(new BigNumber(181), new BigNumber((_52_v6).length))], new BigNumber((_55_v7).length))], globalState);
      } else {
        let _65_v14;
        _65_v14 = _dafny.Map.Empty.slice().updateUnsafe(false,p3);
        _65_v14 = (_65_v14).update(p3, p3);
        let _66_v15;
        let _nw2 = new _module.C0();
        _nw2.__ctor(new BigNumber(536));
        _66_v15 = _nw2;
        let _67_v16;
        _67_v16 = _dafny.Seq.of(_66_v15);
        let _68_v17;
        _68_v17 = _module.D4.create_DC14(_67_v16);
        let _69_v18;
        _69_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC22(),_68_v17);
        let _70_v19;
        let _nw3 = Array((new BigNumber(3)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _69_v18;
        _nw3[(_dafny.ONE).toNumber()] = (_69_v18).Merge(_69_v18);
        _nw3[(new BigNumber(2)).toNumber()] = _69_v18;
        _70_v19 = _nw3;
        let _71_v20;
        _71_v20 = _module.D6.create_DC22();
        let _index7 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_70_v19).length));
        (_70_v19)[_index7] = _dafny.Map.Empty.slice().updateUnsafe(_71_v20,_68_v17);
        let _index8 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_70_v19).length));
        (_70_v19)[_index8] = _69_v18;
        let _72_v21;
        let _nw4 = Array((new BigNumber(25)).toNumber()).fill(false);
        _72_v21 = _nw4;
        let _index9 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_72_v21).length));
        (_72_v21)[_index9] = p1;
        let _index10 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_72_v21).length));
        (_72_v21)[_index10] = p1;
        let _73_v22;
        _73_v22 = _dafny.Set.fromElements(p0);
        let _index11 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_72_v21).length));
        (_72_v21)[_index11] = (((p0) ? (_66_v15.f19) : ((_dafny.ZERO).minus(_66_v15.f19)))).isLessThanOrEqualTo(new BigNumber(((_module.__default.fm17(globalState)).Intersect(_73_v22)).length));
        let _74_v23;
        _74_v23 = _dafny.Seq.of(_66_v15.f19);
        let _75_v24;
        let _nw5 = new _module.C0();
        _nw5.__ctor((_74_v23)[_module.__default.safeIndex(_66_v15.f19, new BigNumber((_74_v23).length))]);
        _75_v24 = _nw5;
      }
      let _76_v25;
      _76_v25 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _hi0 = p2;
      for (let _77_i3 = new BigNumber((_76_v25).length); _77_i3.isLessThan(_hi0); _77_i3 = _77_i3.plus(_dafny.ONE)) {
        let _78_v26;
        let _nw6 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _78_v26 = _nw6;
        let _79_v27;
        _79_v27 = _dafny.MultiSet.fromElements(_78_v26, _78_v26, _78_v26);
        (globalState).f3 = ((_79_v27).Intersect(_79_v27)).IsSubsetOf(_79_v27);
        _78_v26 = _78_v26;
        (globalState).f2 = p2;
        let _80_v28;
        let _nw7 = Array((new BigNumber(5)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = (p2).isEqualTo(_77_i3);
        _nw7[(_dafny.ONE).toNumber()] = p0;
        _nw7[(new BigNumber(2)).toNumber()] = (_module.__default.fm2(!(p1), globalState)) && (_module.__default.fm2(_module.__default.fm2(true, globalState), globalState));
        _nw7[(new BigNumber(3)).toNumber()] = p1;
        _nw7[(new BigNumber(4)).toNumber()] = (p1) && (p1);
        _80_v28 = _nw7;
        let _81_v29;
        let _nw8 = new _module.C0();
        _nw8.__ctor((new BigNumber(343)).multipliedBy(p2));
        _81_v29 = _nw8;
        let _index12 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_80_v28).length));
        (_80_v28)[_index12] = !(_77_i3).isEqualTo((_dafny.ZERO).minus(_77_i3));
        let _82_v30;
        _82_v30 = _dafny.Seq.of(p0);
        let _83_v31;
        _83_v31 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),(_dafny.ZERO).minus(_77_i3));
        let _84_v32;
        _84_v32 = _dafny.Seq.of(_77_i3);
        let _index13 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_80_v28).length));
        let _rhs8 = _80_v28;
        let _rhs9 = _81_v29;
        let _rhs10 = ((_dafny.Set.fromElements((_82_v30)[_module.__default.safeIndex(_77_i3, new BigNumber((_82_v30).length))], (_82_v30)[_module.__default.safeIndex(_81_v29.f19, new BigNumber((_82_v30).length))])).Union(_dafny.Set.fromElements(p1))).IsProperSubsetOf(_module.__default.fm17(globalState));
        let _rhs11 = ((((_83_v31).contains(p0)) ? ((_83_v31).get(p0)) : (_77_i3))).isLessThanOrEqualTo(_81_v29.f19);
        let _rhs12 = ((new BigNumber((_84_v32).length)).isLessThan(new BigNumber((_31_v3).length))) && (p1);
        let _lhs8 = _80_v28;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_80_v28).length));
        let _lhs10 = globalState;
        let _lhs11 = globalState;
        _80_v28 = _rhs8;
        _81_v29 = _rhs9;
        _lhs8[_lhs9] = _rhs10;
        _lhs10.f3 = _rhs11;
        _lhs11.f3 = _rhs12;
      }
      let _pat_let_tv6 = p3;
      let _pat_let_tv7 = p3;
      let _pat_let_tv8 = p1;
      let _pat_let_tv9 = p0;
      let _pat_let_tv10 = p1;
      let _pat_let_tv11 = p0;
      let _pat_let_tv12 = p1;
      let _pat_let_tv13 = p3;
      let _pat_let_tv14 = p0;
      if (function (_source5) {
        if (_source5.is_DC1) {
          let _85___mcc_h9 = (_source5).cf1;
          let _86___mcc_h10 = (_source5).cf2;
          let _87___mcc_h11 = (_source5).cf3;
          let _88_cf3 = _87___mcc_h11;
          let _89_cf2 = _86___mcc_h10;
          let _90_cf1 = _85___mcc_h9;
          return _dafny.areEqual(_dafny.Seq.of(_pat_let_tv6), _dafny.Seq.of(_pat_let_tv7, _pat_let_tv8, _pat_let_tv9, _pat_let_tv10));
        } else if (_source5.is_DC2) {
          let _91___mcc_h12 = (_source5).cf4;
          let _92___mcc_h13 = (_source5).cf5;
          let _93___mcc_h14 = (_source5).cf6;
          let _94___mcc_h15 = (_source5).cf7;
          let _95_cf7 = _94___mcc_h15;
          let _96_cf6 = _93___mcc_h14;
          let _97_cf5 = _92___mcc_h13;
          let _98_cf4 = _91___mcc_h12;
          return _pat_let_tv11;
        } else {
          let _99___mcc_h16 = (_source5).cf0;
          let _100_cf0 = _99___mcc_h16;
          if (_pat_let_tv12) {
            return _pat_let_tv13;
          } else {
            return _pat_let_tv14;
          }
        }
      }(_module.D0.create_DC0(new BigNumber(-878)))) {
        let _101_v33;
        let _nw9 = new _module.C0();
        _nw9.__ctor(p2);
        _101_v33 = _nw9;
        let _102_v34;
        _102_v34 = _dafny.Seq.of(_101_v33);
        let _103_v35;
        _103_v35 = _module.D4.create_DC14(_102_v34);
        let _104_v36;
        _104_v36 = _module.D4.create_DC16(_103_v35);
        let _pat_let_tv15 = _103_v35;
        let _105_v37;
        let _nw10 = Array((new BigNumber(19)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _104_v36;
        _nw10[(_dafny.ONE).toNumber()] = _104_v36;
        _nw10[(new BigNumber(2)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(3)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(4)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(5)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(6)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(7)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(8)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(9)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(10)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(11)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(12)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(13)).toNumber()] = function (_pat_let0_0) {
          return function (_106_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_107_dt__update_hcf22_h0) {
                return _module.D4.create_DC16(_107_dt__update_hcf22_h0);
              }(_pat_let1_0);
            }(_pat_let_tv15);
          }(_pat_let0_0);
        }(_104_v36);
        _nw10[(new BigNumber(14)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(15)).toNumber()] = _104_v36;
        _nw10[(new BigNumber(16)).toNumber()] = _module.D4.create_DC16(_103_v35);
        _nw10[(new BigNumber(17)).toNumber()] = _module.D4.create_DC16(_103_v35);
        _nw10[(new BigNumber(18)).toNumber()] = _104_v36;
        _105_v37 = _nw10;
        let _index14 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_105_v37).length));
        (_105_v37)[_index14] = _104_v36;
        let _index15 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_105_v37).length));
        (_105_v37)[_index15] = _104_v36;
        r0 = _101_v33.f19;
        let _108_v38;
        let _nw11 = new _module.C0();
        _nw11.__ctor(_101_v33.f19);
        _108_v38 = _nw11;
        let _109_v39;
        _109_v39 = _dafny.Seq.of(p3, _module.__default.fm2(p3, globalState));
        let _110_v40;
        _110_v40 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_module.__default.fm14(p1, globalState)).length));
        if ((_109_v39)[_module.__default.safeIndex((((_110_v40).contains(_module.__default.fm2(p0, globalState))) ? ((_110_v40).get(_module.__default.fm2(p0, globalState))) : (new BigNumber(-950))), new BigNumber((_109_v39).length))]) {
          r1 = _module.__default.fm5(new BigNumber((_31_v3).length), globalState);
          let _rhs13 = !((_28_v0).IsProperSubsetOf((_28_v0).update(p3, _module.__default.abs(p2))));
          let _rhs14 = ((!((new BigNumber(384)).isLessThan(_101_v33.f19))) ? (_101_v33) : (_108_v38));
          let _lhs12 = globalState;
          _lhs12.f3 = _rhs13;
          _108_v38 = _rhs14;
          let _111_v41;
          let _init1 = ((_112_p0) => function (_113_i4) {
            return _112_p0;
          })(p0);
          let _nw12 = Array((new BigNumber(20)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw12.length); _i0_1++) {
            _nw12[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _111_v41 = _nw12;
          let _index16 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length));
          (_111_v41)[_index16] = ((p1) ? (false) : (p3));
          let _114_v42;
          _114_v42 = _dafny.Map.Empty.slice().updateUnsafe(_36_v4,(_108_v38.f19).isLessThanOrEqualTo(new BigNumber((_31_v3).length)));
          let _index17 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length));
          (_111_v41)[_index17] = (((_114_v42).contains(_36_v4)) ? ((_114_v42).get(_36_v4)) : (p0));
          let _index18 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length));
          (_111_v41)[_index18] = !((_76_v25).equals(_76_v25)) || ((_28_v0).IsSubsetOf((_dafny.MultiSet.fromElements((_111_v41)[_module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length))], !(true))).update((_111_v41)[_module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length))], _module.__default.abs(p2))));
          let _index19 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length));
          let _rhs15 = _101_v33.f19;
          let _rhs16 = _module.__default.fm2((_111_v41)[_module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length))], globalState);
          let _lhs13 = globalState;
          let _lhs14 = _111_v41;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_111_v41).length));
          _lhs13.f2 = _rhs15;
          _lhs14[_lhs15] = _rhs16;
        } else {
          let _115_v43;
          _115_v43 = _module.D6.create_DC23(p3, _101_v33.f19, _101_v33.f19);
          let _116_v44;
          _116_v44 = _dafny.Map.Empty.slice().updateUnsafe((_115_v43).dtor_cf31,_108_v38);
          let _117_v45;
          let _nw13 = Array((new BigNumber(22)).toNumber()).fill(false);
          _117_v45 = _nw13;
          let _118_v46;
          _118_v46 = _module.D5.create_DC18(_108_v38.f19, p1, _117_v45, p1);
          let _119_v47;
          let _nw14 = Array((new BigNumber(27)).toNumber());
          _nw14[(_dafny.ZERO).toNumber()] = _108_v38;
          _nw14[(_dafny.ONE).toNumber()] = _108_v38;
          _nw14[(new BigNumber(2)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(3)).toNumber()] = _101_v33;
          _nw14[(new BigNumber(4)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(5)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(6)).toNumber()] = _101_v33;
          _nw14[(new BigNumber(7)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(8)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(9)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(10)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(11)).toNumber()] = (((_116_v44).contains((_118_v46).dtor_cf25)) ? ((_116_v44).get((_118_v46).dtor_cf25)) : (_101_v33));
          _nw14[(new BigNumber(12)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(13)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(14)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(15)).toNumber()] = _101_v33;
          _nw14[(new BigNumber(16)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(17)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(18)).toNumber()] = _101_v33;
          _nw14[(new BigNumber(19)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(20)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(21)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(22)).toNumber()] = _101_v33;
          _nw14[(new BigNumber(23)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(24)).toNumber()] = _101_v33;
          _nw14[(new BigNumber(25)).toNumber()] = _108_v38;
          _nw14[(new BigNumber(26)).toNumber()] = _101_v33;
          _119_v47 = _nw14;
          let _120_v48;
          _120_v48 = _dafny.Map.Empty.slice().updateUnsafe(_108_v38.f19,_119_v47);
          _120_v48 = (_120_v48).update(_108_v38.f19, _119_v47);
          (globalState).f3 = p1;
          let _121_v49;
          let _nw15 = new _module.C0();
          _nw15.__ctor((_dafny.ZERO).minus((_dafny.ZERO).minus((_101_v33).fm3(globalState))));
          _121_v49 = _nw15;
          let _122_v50;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_108_v38.f19);
          _122_v50 = _nw16;
          (globalState).f3 = !(p1) || (p0);
        }
        let _123_v51;
        _123_v51 = _dafny.MultiSet.fromElements(new BigNumber(103), _108_v38.f19);
        _123_v51 = _123_v51;
      } else {
        let _124_v52;
        _124_v52 = _dafny.Seq.of(p3);
        let _125_v53;
        let _nw17 = new _module.C0();
        _nw17.__ctor(_module.__default.fm0(new BigNumber((_124_v52).length), _36_v4, _dafny.Seq.of(_31_v3, _31_v3, _31_v3, _31_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), ((_126_v4) => function (_127_i5) {
          return _126_v4;
        })(_36_v4))), p2, globalState));
        _125_v53 = _nw17;
        let _128_v54;
        let _nw18 = new _module.C0();
        _nw18.__ctor(_125_v53.f19);
        _128_v54 = _nw18;
        let _129_v55;
        _129_v55 = _dafny.Set.fromElements(_128_v54.f19);
        if ((_129_v55).IsDisjointFrom((_module.__default.fm11(_128_v54.f19, globalState)).Union(_129_v55))) {
          let _130_v56;
          _130_v56 = _dafny.Set.fromElements(p0);
          let _131_v57;
          _131_v57 = _dafny.Map.Empty.slice().updateUnsafe(_36_v4,new BigNumber((_130_v56).length));
          let _rhs17 = _128_v54;
          let _rhs18 = !((_dafny.ZERO).minus((((_131_v57).contains(_36_v4)) ? ((_131_v57).get(_36_v4)) : (_125_v53.f19)))).isEqualTo(_128_v54.f19);
          let _rhs19 = _128_v54.f19;
          let _lhs16 = globalState;
          let _lhs17 = globalState;
          _128_v54 = _rhs17;
          _lhs16.f3 = _rhs18;
          _lhs17.f8 = _rhs19;
          let _132_v58;
          let _nw19 = new _module.C0();
          _nw19.__ctor((_128_v54.f19).plus(_125_v53.f19));
          _132_v58 = _nw19;
          let _133_v59;
          _133_v59 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_128_v54);
          _133_v59 = (_133_v59).Merge(_133_v59);
          (globalState).f3 = ((new BigNumber(-942)).isLessThanOrEqualTo(new BigNumber((_31_v3).length))) && (p0);
          (globalState).f2 = _125_v53.f19;
        } else {
          (globalState).f8 = _125_v53.f19;
          let _134_v60;
          _134_v60 = _module.D6.create_DC23(p1, _125_v53.f19, _125_v53.f19);
          let _135_v61;
          _135_v61 = _dafny.MultiSet.fromElements(p2, new BigNumber((_31_v3).length), _125_v53.f19);
          let _136_v62;
          _136_v62 = _dafny.Map.Empty.slice().updateUnsafe(_128_v54.f19,p3);
          let _137_v63;
          _137_v63 = _dafny.Map.Empty.slice().updateUnsafe((((_136_v62).contains(_125_v53.f19)) ? ((_136_v62).get(_125_v53.f19)) : (p0)),p0);
          let _pat_let_tv16 = _128_v54;
          let _pat_let_tv17 = _135_v61;
          let _138_v64;
          let _nw20 = Array((new BigNumber(29)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _134_v60;
          _nw20[(_dafny.ONE).toNumber()] = _134_v60;
          _nw20[(new BigNumber(2)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(3)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(4)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(5)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(6)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(7)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(8)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(9)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(10)).toNumber()] = function (_pat_let2_0) {
            return function (_139_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_140_dt__update_hcf33_h0) {
                  return _module.D6.create_DC23((_139_dt__update__tmp_h1).dtor_cf31, (_139_dt__update__tmp_h1).dtor_cf32, _140_dt__update_hcf33_h0);
                }(_pat_let3_0);
              }(_pat_let_tv16.f19);
            }(_pat_let2_0);
          }(_module.D6.create_DC23(p3, new BigNumber((_31_v3).length), _125_v53.f19));
          _nw20[(new BigNumber(11)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(12)).toNumber()] = function (_pat_let4_0) {
            return function (_141_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_142_dt__update_hcf32_h0) {
                  return _module.D6.create_DC23((_141_dt__update__tmp_h2).dtor_cf31, _142_dt__update_hcf32_h0, (_141_dt__update__tmp_h2).dtor_cf33);
                }(_pat_let5_0);
              }(new BigNumber((_pat_let_tv17).cardinality()));
            }(_pat_let4_0);
          }(_134_v60);
          _nw20[(new BigNumber(13)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(14)).toNumber()] = _module.D6.create_DC23(p1, _125_v53.f19, _125_v53.f19);
          _nw20[(new BigNumber(15)).toNumber()] = _module.D6.create_DC23((((_137_v63).contains(p0)) ? ((_137_v63).get(p0)) : (p0)), new BigNumber(-328), (_dafny.ZERO).minus(_128_v54.f19));
          _nw20[(new BigNumber(16)).toNumber()] = ((p0) ? (_134_v60) : (_134_v60));
          _nw20[(new BigNumber(17)).toNumber()] = ((false) ? (_134_v60) : (_134_v60));
          _nw20[(new BigNumber(18)).toNumber()] = _module.D6.create_DC23(p0, p2, _125_v53.f19);
          _nw20[(new BigNumber(19)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(20)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(21)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(22)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(23)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(24)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(25)).toNumber()] = ((p3) ? (_134_v60) : (_134_v60));
          _nw20[(new BigNumber(26)).toNumber()] = _module.__default.fm18(_128_v54.f19, p1, globalState);
          _nw20[(new BigNumber(27)).toNumber()] = _134_v60;
          _nw20[(new BigNumber(28)).toNumber()] = _134_v60;
          _138_v64 = _nw20;
          let _index20 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_138_v64).length));
          (_138_v64)[_index20] = ((p3) ? (_134_v60) : (_module.D6.create_DC23(p0, _128_v54.f19, (((_28_v0).contains((((_136_v62).contains(_125_v53.f19)) ? ((_136_v62).get(_125_v53.f19)) : (!(p0))))) ? ((_28_v0).get((((_136_v62).contains(_125_v53.f19)) ? ((_136_v62).get(_125_v53.f19)) : (!(p0))))) : (_128_v54.f19)))));
          let _index21 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_138_v64).length));
          (_138_v64)[_index21] = _module.__default.fm18(_125_v53.f19, false, globalState);
          (globalState).f3 = _dafny.Seq.IsProperPrefixOf(_31_v3, _dafny.Seq.UnicodeFromString("hnidaeeo"));
          let _143_v65;
          _143_v65 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_dafny.ZERO).minus(_125_v53.f19));
          _143_v65 = (_143_v65).update(p1, _125_v53.f19);
          let _144_v66;
          let _nw21 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _144_v66 = _nw21;
          let _index22 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_144_v66).length));
          (_144_v66)[_index22] = _128_v54.f19;
          let _index23 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_144_v66).length));
          (_144_v66)[_index23] = _128_v54.f19;
        }
        let _145_v67;
        let _nw22 = Array((new BigNumber(27)).toNumber());
        _145_v67 = _nw22;
        let _index24 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_145_v67).length));
        (_145_v67)[_index24] = _125_v53;
        let _index25 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_145_v67).length));
        let _nw23 = new _module.C0();
        _nw23.__ctor(_module.__default.safeDivisionInt(new BigNumber(115), (_dafny.ZERO).minus(p2)));
        (_145_v67)[_index25] = _nw23;
        let _146_v68;
        _146_v68 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(198));
        let _147_v69;
        _147_v69 = _dafny.Seq.of(_31_v3);
        _146_v68 = (_146_v68).update(!((_125_v53).fm4(p2, _128_v54.f19, p0, p0, globalState)).isEqualTo(_128_v54.f19), new BigNumber(((_147_v69)[_module.__default.safeIndex(_125_v53.f19, new BigNumber((_147_v69).length))]).length));
      }
      let _148_v70;
      _148_v70 = _dafny.Seq.of(false, p0);
      let _149_v71;
      _149_v71 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _150_v72;
      _150_v72 = _dafny.MultiSet.fromElements(_149_v71, _dafny.Map.Empty.slice().updateUnsafe(p3,p3));
      let _151_v73;
      _151_v73 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_150_v72).cardinality()));
      let _152_v74;
      _152_v74 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_151_v73).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p2)));
      let _rhs20 = _148_v70;
      let _rhs21 = _dafny.Seq.Concat(_31_v3, _dafny.Seq.Concat(_31_v3, _31_v3));
      let _rhs22 = _152_v74;
      let _rhs23 = _37_v5;
      _148_v70 = _rhs20;
      r1 = _rhs21;
      _152_v74 = _rhs22;
      _37_v5 = _rhs23;
      let _153_v75;
      let _nw24 = new _module.C0();
      _nw24.__ctor(p2);
      _153_v75 = _nw24;
      r0 = (new BigNumber((_dafny.Seq.of(_153_v75)).length)).multipliedBy(p2);
      r1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_31_v3, _dafny.Seq.UnicodeFromString("ey")), _module.__default.fm5(_153_v75.f19, globalState));
      r2 = _153_v75.f19;
      let _154_v76;
      _154_v76 = _dafny.Set.fromElements(p0, true, p1);
      r3 = (_28_v0).update((p2).isLessThan(new BigNumber((_154_v76).length)), _module.__default.abs(_module.__default.safeModuloInt(p2, p2)));
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _155_v0;
      _155_v0 = false;
      let _156_v1;
      _156_v1 = _dafny.Seq.of(_155_v0);
      let _157_globalState;
      let _nw25 = new _module.GlobalState();
      _nw25.__ctor(new BigNumber(-782), false, new BigNumber(351), false, new BigNumber(-139), new BigNumber(973), false, new BigNumber(-608), new BigNumber(117), new BigNumber(-409), _dafny.Seq.UnicodeFromString("adqetglt"), false, new BigNumber(-956), true, new BigNumber(26), false, new BigNumber(957), new BigNumber(182), _156_v1);
      _157_globalState = _nw25;
      let _158_v2;
      _158_v2 = new BigNumber(270);
      let _159_v3;
      _159_v3 = _dafny.MultiSet.fromElements(_155_v0);
      let _hi1 = (_dafny.ZERO).minus((_158_v2).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_159_v3).cardinality())))));
      for (let _160_i0 = new BigNumber(115); _160_i0.isLessThan(_hi1); _160_i0 = _160_i0.plus(_dafny.ONE)) {
        let _161_v4;
        _161_v4 = new _dafny.CodePoint('f'.codePointAt(0));
        let _162_v5;
        _162_v5 = _dafny.Seq.UnicodeFromString("oiwhio");
        let _163_v6;
        let _nw26 = Array((new BigNumber(3)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _162_v5;
        _nw26[(_dafny.ONE).toNumber()] = _162_v5;
        _nw26[(new BigNumber(2)).toNumber()] = _162_v5;
        _163_v6 = _nw26;
        let _164_v7;
        _164_v7 = _dafny.Set.fromElements(_155_v0);
        let _165_v8;
        _165_v8 = _dafny.Seq.of(_162_v5);
        let _166_v9;
        _166_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_164_v7).length), _161_v4, _165_v8, new BigNumber(591), _157_globalState),_163_v6);
        let _167_v10;
        _167_v10 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_155_v0);
        let _168_v11;
        let _nw27 = Array((new BigNumber(19)).toNumber());
        _nw27[(_dafny.ZERO).toNumber()] = _163_v6;
        _nw27[(_dafny.ONE).toNumber()] = _163_v6;
        _nw27[(new BigNumber(2)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(3)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(4)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(5)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(6)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(7)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(8)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(9)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(10)).toNumber()] = (((_166_v9).contains(new BigNumber((_167_v10).length))) ? ((_166_v9).get(new BigNumber((_167_v10).length))) : (_163_v6));
        _nw27[(new BigNumber(11)).toNumber()] = ((_155_v0) ? (_163_v6) : (_163_v6));
        _nw27[(new BigNumber(12)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(13)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(14)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(15)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(16)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(17)).toNumber()] = _163_v6;
        _nw27[(new BigNumber(18)).toNumber()] = _163_v6;
        _168_v11 = _nw27;
        let _index26 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_168_v11).length));
        (_168_v11)[_index26] = _163_v6;
        let _169_v12;
        let _nw28 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
        _169_v12 = _nw28;
        let _index27 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_169_v12).length));
        (_169_v12)[_index27] = _164_v7;
        let _170_v13;
        _170_v13 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_164_v7);
        let _index28 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_168_v11).length));
        let _index29 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_169_v12).length));
        let _rhs24 = new _dafny.CodePoint('r'.codePointAt(0));
        let _rhs25 = _163_v6;
        let _rhs26 = _155_v0;
        let _rhs27 = (_164_v7).Intersect((((_170_v13).contains(_160_i0)) ? ((_170_v13).get(_160_i0)) : (_164_v7)));
        let _lhs18 = _168_v11;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_168_v11).length));
        let _lhs20 = _157_globalState;
        let _lhs21 = _169_v12;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_169_v12).length));
        _161_v4 = _rhs24;
        _lhs18[_lhs19] = _rhs25;
        _lhs20.f3 = _rhs26;
        _lhs21[_lhs22] = _rhs27;
        let _171_v14;
        _171_v14 = _dafny.Seq.of(new BigNumber(172));
        let _172_v15;
        _172_v15 = _dafny.Map.Empty.slice().updateUnsafe((_171_v14)[_module.__default.safeIndex(new BigNumber(-542), new BigNumber((_171_v14).length))],_160_i0);
        let _rhs28 = _158_v2;
        let _rhs29 = (_dafny.ZERO).minus((((_dafny.ZERO).minus((((_172_v15).contains(_158_v2)) ? ((_172_v15).get(_158_v2)) : (_158_v2)))).plus(_158_v2)).multipliedBy(_160_i0));
        let _rhs30 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gydcso")).length))).minus(new BigNumber(-566));
        let _rhs31 = _158_v2;
        let _lhs23 = _157_globalState;
        let _lhs24 = _157_globalState;
        let _lhs25 = _157_globalState;
        let _lhs26 = _157_globalState;
        _lhs23.f12 = _rhs28;
        _lhs24.f8 = _rhs29;
        _lhs25.f2 = _rhs30;
        _lhs26.f12 = _rhs31;
        let _173_v16;
        let _init2 = ((_174_v5) => function (_175_i1) {
          return (_dafny.Set.fromElements(_174_v5)).Difference(_dafny.Set.fromElements(_174_v5, _dafny.Seq.UnicodeFromString("kvf")));
        })(_162_v5);
        let _nw29 = Array((new BigNumber(6)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw29.length); _i0_2++) {
          _nw29[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _173_v16 = _nw29;
        let _176_v17;
        _176_v17 = _dafny.Set.fromElements(_162_v5, _162_v5);
        let _index30 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_173_v16).length));
        (_173_v16)[_index30] = _176_v17;
        let _index31 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_173_v16).length));
        (_173_v16)[_index31] = ((_176_v17).Intersect(_176_v17)).Intersect(_176_v17);
        (_157_globalState).f8 = ((_dafny.ZERO).minus(_158_v2)).minus(_158_v2);
      }
      let _177_v18;
      _177_v18 = _dafny.Seq.UnicodeFromString("gcqnkipj");
      _177_v18 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_178_i2) {
        return ((true) ? (new _dafny.CodePoint('g'.codePointAt(0))) : (new _dafny.CodePoint('x'.codePointAt(0))));
      });
      if (_155_v0) {
        let _179_v19;
        let _180_v20;
        let _181_v21;
        let _182_v22;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m0(_155_v0, _155_v0, _158_v2, _155_v0, _157_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _179_v19 = _out0;
        _180_v20 = _out1;
        _181_v21 = _out2;
        _182_v22 = _out3;
        _155_v0 = !(_155_v0);
        let _183_v23;
        _183_v23 = new _dafny.CodePoint('q'.codePointAt(0));
        let _184_v24;
        _184_v24 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_183_v23);
        let _185_v25;
        _185_v25 = _dafny.Set.fromElements(_184_v24);
        let _186_v26;
        _186_v26 = _module.D0.create_DC0(_158_v2);
        let _pat_let_tv18 = _181_v21;
        let _187_v27;
        let _nw30 = Array((new BigNumber(5)).toNumber());
        _nw30[(_dafny.ZERO).toNumber()] = new BigNumber((_185_v25).length);
        _nw30[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_177_v18).length));
        _nw30[(new BigNumber(2)).toNumber()] = _179_v19;
        _nw30[(new BigNumber(3)).toNumber()] = new BigNumber(694);
        _nw30[(new BigNumber(4)).toNumber()] = (function (_pat_let6_0) {
          return function (_188_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_189_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_189_dt__update_hcf0_h0);
              }(_pat_let7_0);
            }(_pat_let_tv18);
          }(_pat_let6_0);
        }(_186_v26)).dtor_cf0;
        _187_v27 = _nw30;
        let _190_v28;
        _190_v28 = _dafny.MultiSet.fromElements(_187_v27);
        _190_v28 = (_190_v28).Intersect(_190_v28);
        let _191_v29;
        _191_v29 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_182_v22);
        let _192_v30;
        _192_v30 = _module.D0.create_DC2(_155_v0, (_158_v2).isLessThanOrEqualTo(_179_v19), _181_v21, true);
        let _rhs32 = _179_v19;
        let _rhs33 = (_191_v29).update(_179_v19, _dafny.MultiSet.fromElements(_155_v0));
        let _rhs34 = _192_v30;
        let _lhs27 = _157_globalState;
        _lhs27.f2 = _rhs32;
        _191_v29 = _rhs33;
        _192_v30 = _rhs34;
        let _193_v31;
        let _nw31 = Array((new BigNumber(23)).toNumber()).fill([]);
        _193_v31 = _nw31;
        let _194_v32;
        _194_v32 = _dafny.Seq.of(_177_v18, _177_v18);
        let _195_v33;
        let _nw32 = Array((new BigNumber(23)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_196_v23) => function (_197_i3) {
          return _196_v23;
        })(_183_v23)), _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm0(_179_v19, new _dafny.CodePoint('r'.codePointAt(0)), _194_v32, _179_v19, _157_globalState)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_198_v23) => function (_199_i3) {
          return _198_v23;
        })(_183_v23))).length)), new _dafny.CodePoint('h'.codePointAt(0)));
        _nw32[(_dafny.ONE).toNumber()] = _177_v18;
        _nw32[(new BigNumber(2)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(3)).toNumber()] = _180_v20;
        _nw32[(new BigNumber(4)).toNumber()] = _180_v20;
        _nw32[(new BigNumber(5)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), ((_200_v23) => function (_201_i4) {
          return _200_v23;
        })(_183_v23));
        _nw32[(new BigNumber(7)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(8)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(9)).toNumber()] = _180_v20;
        _nw32[(new BigNumber(10)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(11)).toNumber()] = _180_v20;
        _nw32[(new BigNumber(12)).toNumber()] = _180_v20;
        _nw32[(new BigNumber(13)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(14)).toNumber()] = _180_v20;
        _nw32[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("xeumplyd");
        _nw32[(new BigNumber(16)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(550)), ((_202_v23) => function (_203_i5) {
          return _202_v23;
        })(_183_v23));
        _nw32[(new BigNumber(18)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(19)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), ((_204_v23) => function (_205_i6) {
          return _204_v23;
        })(_183_v23));
        _nw32[(new BigNumber(21)).toNumber()] = _177_v18;
        _nw32[(new BigNumber(22)).toNumber()] = _180_v20;
        _195_v33 = _nw32;
        let _index32 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_193_v31).length));
        (_193_v31)[_index32] = _195_v33;
        let _206_v34;
        _206_v34 = _dafny.MultiSet.fromElements(_186_v26);
        let _207_v35;
        _207_v35 = _dafny.MultiSet.fromElements(_179_v19, _181_v21);
        let _index33 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_187_v27).length));
        (_187_v27)[_index33] = _module.__default.safeModuloInt(new BigNumber(((_206_v34).update(_186_v26, _module.__default.abs(new BigNumber((_182_v22).cardinality())))).cardinality()), _module.__default.fm0(_158_v2, _183_v23, _194_v32, new BigNumber((_207_v35).cardinality()), _157_globalState));
        let _index34 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_193_v31).length));
        let _index35 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_187_v27).length));
        let _rhs35 = _187_v27;
        let _rhs36 = _158_v2;
        let _rhs37 = _195_v33;
        let _rhs38 = (_dafny.ZERO).minus(_158_v2);
        let _lhs28 = _157_globalState;
        let _lhs29 = _193_v31;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_193_v31).length));
        let _lhs31 = _187_v27;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_187_v27).length));
        _187_v27 = _rhs35;
        _lhs28.f0 = _rhs36;
        _lhs29[_lhs30] = _rhs37;
        _lhs31[_lhs32] = _rhs38;
      } else {
        let _208_v36;
        _208_v36 = new _dafny.CodePoint('s'.codePointAt(0));
        _156_v1 = _module.__default.fm1(_208_v36, _155_v0, _157_globalState);
        if (_module.__default.fm2((_155_v0) === (_155_v0), _157_globalState)) {
          let _209_v37;
          let _nw33 = new _module.C0();
          _nw33.__ctor((_158_v2).plus((_dafny.ZERO).minus(_158_v2)));
          _209_v37 = _nw33;
          let _210_v38;
          _210_v38 = _dafny.Seq.of(_177_v18);
          let _211_v39;
          _211_v39 = _dafny.Set.fromElements(_155_v0);
          (_157_globalState).f0 = _module.__default.fm0((_209_v37.f19).multipliedBy(_209_v37.f19), _208_v36, _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(_208_v36), _177_v18), _210_v38), new BigNumber(((_211_v39).Intersect(_211_v39)).length), _157_globalState);
          let _212_v40;
          let _213_v41;
          let _214_v42;
          let _215_v43;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(_155_v0, _dafny.Seq.IsPrefixOf(_177_v18, _177_v18), _209_v37.f19, _155_v0, _157_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _212_v40 = _out4;
          _213_v41 = _out5;
          _214_v42 = _out6;
          _215_v43 = _out7;
          _158_v2 = (_dafny.ZERO).minus(_212_v40);
          let _216_v44;
          let _nw34 = Array((new BigNumber(14)).toNumber()).fill(false);
          _216_v44 = _nw34;
          let _index36 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_216_v44).length));
          (_216_v44)[_index36] = true;
          let _index37 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_216_v44).length));
          (_216_v44)[_index37] = _155_v0;
        } else {
          let _217_v45;
          _217_v45 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,(_155_v0) || (_155_v0));
          _217_v45 = (_217_v45).update(_158_v2, _module.__default.fm2(false, _157_globalState));
          let _218_v46;
          let _nw35 = Array((new BigNumber(6)).toNumber()).fill([]);
          _218_v46 = _nw35;
          _218_v46 = _218_v46;
          let _219_v47;
          _219_v47 = _dafny.Set.fromElements(((_155_v0) ? (_155_v0) : (_155_v0)));
          let _220_v48;
          let _nw36 = new _module.C0();
          _nw36.__ctor(_158_v2);
          _220_v48 = _nw36;
          let _221_v49;
          _221_v49 = _dafny.Map.Empty.slice().updateUnsafe(_208_v36,_module.__default.fm2(_155_v0, _157_globalState));
          let _222_v50;
          _222_v50 = _dafny.Set.fromElements(_220_v48, _220_v48);
          let _223_v51;
          _223_v51 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_208_v36);
          let _224_v52;
          _224_v52 = _dafny.Seq.of(_223_v51, _223_v51);
          let _225_v53;
          _225_v53 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,new BigNumber((_156_v1).length));
          let _226_v54;
          _226_v54 = _dafny.Set.fromElements((_224_v52)[_module.__default.safeIndex(_158_v2, new BigNumber((_224_v52).length))], (_223_v51).update(new BigNumber((_225_v53).length), _208_v36), (_223_v51).update(new BigNumber(-407), _208_v36));
          let _227_v55;
          _227_v55 = _dafny.Seq.of((_dafny.ZERO).minus(_158_v2));
          let _228_v56;
          _228_v56 = _module.D0.create_DC2(!(!(_155_v0)), _155_v0, _158_v2, _155_v0);
          let _229_v57;
          let _nw37 = Array((new BigNumber(16)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _155_v0;
          _nw37[(_dafny.ONE).toNumber()] = _module.__default.fm2((((_221_v49).contains(_208_v36)) ? ((_221_v49).get(_208_v36)) : (_155_v0)), _157_globalState);
          _nw37[(new BigNumber(2)).toNumber()] = (_222_v50).IsProperSubsetOf(_222_v50);
          _nw37[(new BigNumber(3)).toNumber()] = (_226_v54).contains(_223_v51);
          _nw37[(new BigNumber(4)).toNumber()] = (_156_v1)[_module.__default.safeIndex(new BigNumber((_227_v55).length), new BigNumber((_156_v1).length))];
          _nw37[(new BigNumber(5)).toNumber()] = _155_v0;
          _nw37[(new BigNumber(6)).toNumber()] = _155_v0;
          _nw37[(new BigNumber(7)).toNumber()] = (_228_v56).dtor_cf4;
          _nw37[(new BigNumber(8)).toNumber()] = !((_155_v0) === (_155_v0));
          _nw37[(new BigNumber(9)).toNumber()] = !_dafny.areEqual(_227_v55, _dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), ((_230_v48) => function (_231_i7) {
            return _230_v48.f19;
          })(_220_v48)));
          _nw37[(new BigNumber(10)).toNumber()] = _155_v0;
          _nw37[(new BigNumber(11)).toNumber()] = (_155_v0) || (_155_v0);
          _nw37[(new BigNumber(12)).toNumber()] = _155_v0;
          _nw37[(new BigNumber(13)).toNumber()] = true;
          _nw37[(new BigNumber(14)).toNumber()] = _155_v0;
          _nw37[(new BigNumber(15)).toNumber()] = _155_v0;
          _229_v57 = _nw37;
          let _index38 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_229_v57).length));
          (_229_v57)[_index38] = _155_v0;
          let _232_v58;
          _232_v58 = _dafny.Seq.of(_dafny.Seq.of(_208_v36), _177_v18, _177_v18);
          let _index39 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_229_v57).length));
          let _rhs39 = ((_219_v47).Intersect(_219_v47)).Union(_219_v47);
          let _rhs40 = (_158_v2).minus(_module.__default.fm0(_220_v48.f19, _208_v36, _232_v58, new BigNumber((_177_v18).length), _157_globalState));
          let _rhs41 = _220_v48;
          let _rhs42 = !((_220_v48.f19).isEqualTo(new BigNumber((_225_v53).length))) || (_155_v0);
          let _lhs33 = _157_globalState;
          let _lhs34 = _229_v57;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_229_v57).length));
          _219_v47 = _rhs39;
          _lhs33.f2 = _rhs40;
          _220_v48 = _rhs41;
          _lhs34[_lhs35] = _rhs42;
          let _233_v59;
          let _234_v60;
          let _235_v61;
          let _236_v62;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0((_220_v48.f19).isLessThan(_158_v2), !_dafny.Seq.contains(_156_v1, _module.__default.fm2(_155_v0, _157_globalState)), (_dafny.ZERO).minus(_220_v48.f19), _155_v0, _157_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _233_v59 = _out8;
          _234_v60 = _out9;
          _235_v61 = _out10;
          _236_v62 = _out11;
          (_157_globalState).f3 = (_229_v57)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_229_v57).length))];
        }
        (_157_globalState).f2 = _158_v2;
        (_157_globalState).f2 = _158_v2;
        (_157_globalState).f3 = _155_v0;
      }
      let _237_v63;
      _237_v63 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_155_v0, _155_v0, _155_v0, _155_v0),(new BigNumber((_177_v18).length)).isEqualTo(_158_v2));
      let _238_v64;
      _238_v64 = _dafny.Seq.of(_158_v2);
      let _239_v65;
      let _nw38 = new _module.C0();
      _nw38.__ctor(_158_v2);
      _239_v65 = _nw38;
      let _240_v66;
      _240_v66 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_239_v65);
      let _241_v67;
      _241_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_240_v66).length),_158_v2);
      let _242_v68;
      let _nw39 = new _module.C0();
      _nw39.__ctor((_238_v64)[_module.__default.safeIndex(new BigNumber((_241_v67).length), new BigNumber((_238_v64).length))]);
      _242_v68 = _nw39;
      let _243_v69;
      _243_v69 = _dafny.Map.Empty.slice().updateUnsafe(_239_v65,_155_v0);
      _237_v63 = (_237_v63).update(_dafny.Set.fromElements(!((((_243_v69).contains(_239_v65)) ? ((_243_v69).get(_239_v65)) : (_155_v0)))), _155_v0);
      _155_v0 = _155_v0;
      let _244_i8;
      _244_i8 = _dafny.ZERO;
      L0: {
        while (_155_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_244_i8)) {
              break L0;
            }
            _244_i8 = (_244_i8).plus(_dafny.ONE);
            if (_155_v0) {
              let _245_v70;
              _245_v70 = _module.D0.create_DC2(true, _155_v0, _242_v68.f19, false);
              let _246_v71;
              let _nw40 = Array((new BigNumber(27)).toNumber());
              _nw40[(_dafny.ZERO).toNumber()] = _155_v0;
              _nw40[(_dafny.ONE).toNumber()] = _155_v0;
              _nw40[(new BigNumber(2)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(3)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(4)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(5)).toNumber()] = (_156_v1)[_module.__default.safeIndex(_239_v65.f19, new BigNumber((_156_v1).length))];
              _nw40[(new BigNumber(6)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(7)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(8)).toNumber()] = _module.__default.fm2(!(_155_v0), _157_globalState);
              _nw40[(new BigNumber(9)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(10)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(11)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(12)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(13)).toNumber()] = (_245_v70).dtor_cf4;
              _nw40[(new BigNumber(14)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(15)).toNumber()] = !(_155_v0);
              _nw40[(new BigNumber(16)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(17)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(18)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(19)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(20)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(21)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(22)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(23)).toNumber()] = !(_155_v0);
              _nw40[(new BigNumber(24)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(25)).toNumber()] = _155_v0;
              _nw40[(new BigNumber(26)).toNumber()] = true;
              _246_v71 = _nw40;
              let _247_v72;
              _247_v72 = _dafny.Map.Empty.slice().updateUnsafe(_246_v71,_155_v0);
              _247_v72 = (_247_v72).update(_246_v71, _155_v0);
              let _248_v73;
              let _nw41 = Array((new BigNumber(27)).toNumber());
              _248_v73 = _nw41;
              let _index40 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_248_v73).length));
              (_248_v73)[_index40] = _239_v65;
              let _index41 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_248_v73).length));
              (_248_v73)[_index41] = _242_v68;
              let _249_v74;
              let _nw42 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
              _249_v74 = _nw42;
              let _index42 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_249_v74).length));
              (_249_v74)[_index42] = _158_v2;
              let _250_v75;
              let _nw43 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
              _250_v75 = _nw43;
              let _251_v76;
              _251_v76 = _module.D0.create_DC0(_239_v65.f19);
              let _252_v77;
              _252_v77 = _dafny.Seq.of(_156_v1);
              let _253_v78;
              _253_v78 = _dafny.MultiSet.fromElements(_156_v1, (_252_v77)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_252_v77).length))], _156_v1);
              let _254_v79;
              _254_v79 = new _dafny.CodePoint('s'.codePointAt(0));
              let _255_v80;
              _255_v80 = _dafny.Seq.of(_module.__default.fm5(_158_v2, _157_globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(382)), ((_256_v79) => function (_257_i9) {
                return _256_v79;
              })(_254_v79)), _177_v18);
              let _258_v81;
              _258_v81 = _dafny.Map.Empty.slice().updateUnsafe(_248_v73,_250_v75);
              let _index43 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_249_v74).length));
              let _rhs43 = !(((_155_v0) ? (_155_v0) : (_155_v0))) || (_155_v0);
              let _rhs44 = _242_v68.f19;
              let _rhs45 = _module.__default.fm0(new BigNumber(((_dafny.MultiSet.FromArray(_252_v77)).Intersect(_253_v78)).cardinality()), _254_v79, _255_v80, _239_v65.f19, _157_globalState);
              let _rhs46 = (((_258_v81).contains(_248_v73)) ? ((_258_v81).get(_248_v73)) : (_250_v75));
              let _rhs47 = _251_v76;
              let _lhs36 = _157_globalState;
              let _lhs37 = _242_v68;
              let _lhs38 = _249_v74;
              let _lhs39 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_249_v74).length));
              _lhs36.f3 = _rhs43;
              _lhs37.f19 = _rhs44;
              _lhs38[_lhs39] = _rhs45;
              _250_v75 = _rhs46;
              _251_v76 = _rhs47;
              (_239_v65).f19 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_242_v68.f19));
              let _index44 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_246_v71).length));
              (_246_v71)[_index44] = _155_v0;
              let _259_v82;
              let _init3 = ((_260_v79) => function (_261_i10) {
                return _260_v79;
              })(_254_v79);
              let _nw44 = Array((new BigNumber(9)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw44.length); _i0_3++) {
                _nw44[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _259_v82 = _nw44;
              let _262_v83;
              _262_v83 = _dafny.MultiSet.fromElements(_259_v82, _259_v82);
              let _index45 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_246_v71).length));
              let _rhs48 = _158_v2;
              let _rhs49 = ((_155_v0) ? ((new BigNumber((_262_v83).cardinality())).isLessThanOrEqualTo(_239_v65.f19)) : (_155_v0));
              let _rhs50 = (_242_v68.f19).isLessThanOrEqualTo(_242_v68.f19);
              let _lhs40 = _157_globalState;
              let _lhs41 = _246_v71;
              let _lhs42 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_246_v71).length));
              _lhs40.f8 = _rhs48;
              _lhs41[_lhs42] = _rhs49;
              _155_v0 = _rhs50;
            } else {
              _155_v0 = false;
              let _263_v84;
              let _init4 = ((_264_v0) => function (_265_i11) {
                return _264_v0;
              })(_155_v0);
              let _nw45 = Array((new BigNumber(8)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw45.length); _i0_4++) {
                _nw45[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _263_v84 = _nw45;
              let _266_v85;
              _266_v85 = _dafny.Seq.of(_263_v84, _263_v84);
              let _267_v86;
              _267_v86 = _dafny.MultiSet.fromElements(_263_v84, _263_v84, (_266_v85)[_module.__default.safeIndex(new BigNumber((_177_v18).length), new BigNumber((_266_v85).length))]);
              _267_v86 = (_267_v86).Union((_267_v86).update(_263_v84, _module.__default.abs(_239_v65.f19)));
              let _268_v87;
              _268_v87 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_155_v0);
              let _269_v88;
              _269_v88 = _dafny.Map.Empty.slice().updateUnsafe(_268_v87,_dafny.Seq.UnicodeFromString("p"));
              let _270_v89;
              _270_v89 = _module.D0.create_DC2(_155_v0, _155_v0, _242_v68.f19, _155_v0);
              let _271_v90;
              _271_v90 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_270_v89, _270_v89),_155_v0);
              let _272_v91;
              _272_v91 = _dafny.Seq.of(_269_v88);
              let _273_v92;
              _273_v92 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("v"));
              let _274_v94;
              _274_v94 = _dafny.Seq.of(_268_v87, _dafny.Map.Empty.slice().updateUnsafe(_155_v0,(((_268_v87).contains(true)) ? ((_268_v87).get(true)) : (_155_v0))));
              let _275_v96;
              _275_v96 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("qbbrxwsjm"), _177_v18);
              let _276_v97;
              let _nw46 = Array((new BigNumber(26)).toNumber());
              _nw46[(_dafny.ZERO).toNumber()] = (_269_v88).Merge(_269_v88);
              _nw46[(_dafny.ONE).toNumber()] = _269_v88;
              _nw46[(new BigNumber(2)).toNumber()] = (_269_v88).Merge(_269_v88);
              _nw46[(new BigNumber(3)).toNumber()] = _module.__default.fm6(_239_v65.f19, _158_v2, _155_v0, _271_v90, _157_globalState);
              _nw46[(new BigNumber(4)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(5)).toNumber()] = (_269_v88).update(_dafny.Map.Empty.slice().updateUnsafe(_155_v0,_module.__default.fm2(_155_v0, _157_globalState)), _177_v18);
              _nw46[(new BigNumber(6)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_268_v87,_dafny.Seq.UnicodeFromString("yv"));
              _nw46[(new BigNumber(8)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(9)).toNumber()] = (_269_v88).Merge(_269_v88);
              _nw46[(new BigNumber(10)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(11)).toNumber()] = (_272_v91)[_module.__default.safeIndex(_158_v2, new BigNumber((_272_v91).length))];
              _nw46[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_268_v87,(_273_v92).dtor_cf8)).Merge(function () {
                let _coll5 = new _dafny.Map();
                for (const _compr_5 of (_dafny.MultiSet.FromArray(_274_v94)).Elements) {
                  let _277_v93 = _compr_5;
                  if ((_dafny.MultiSet.FromArray(_274_v94)).contains(_277_v93)) {
                    _coll5.push([_277_v93,_177_v18]);
                  }
                }
                return _coll5;
              }());
              _nw46[(new BigNumber(13)).toNumber()] = (_269_v88).Merge(_269_v88);
              _nw46[(new BigNumber(14)).toNumber()] = function () {
                let _coll6 = new _dafny.Map();
                for (const _compr_6 of (_274_v94).Elements) {
                  let _278_v95 = _compr_6;
                  if (_dafny.Seq.contains(_274_v94, _278_v95)) {
                    _coll6.push([_278_v95,_177_v18]);
                  }
                }
                return _coll6;
              }();
              _nw46[(new BigNumber(15)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(16)).toNumber()] = ((_269_v88).update(_268_v87, (_275_v96)[_module.__default.safeIndex(new BigNumber((_238_v64).length), new BigNumber((_275_v96).length))])).Merge(_269_v88);
              _nw46[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_268_v87,_177_v18);
              _nw46[(new BigNumber(18)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(19)).toNumber()] = (_272_v91)[_module.__default.safeIndex(_239_v65.f19, new BigNumber((_272_v91).length))];
              _nw46[(new BigNumber(20)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_268_v87,_dafny.Seq.update(_177_v18, _module.__default.safeIndex(new BigNumber(-222), new BigNumber((_177_v18).length)), (_177_v18)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_177_v18).length))]));
              _nw46[(new BigNumber(22)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(23)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(24)).toNumber()] = _269_v88;
              _nw46[(new BigNumber(25)).toNumber()] = _269_v88;
              _276_v97 = _nw46;
              let _index46 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_276_v97).length));
              (_276_v97)[_index46] = (_269_v88).Merge(_dafny.Map.Empty.slice().updateUnsafe(_268_v87,_dafny.Seq.Create(_module.__default.abs(new BigNumber(37)), function (_279_i12) {
                return new _dafny.CodePoint('c'.codePointAt(0));
              })));
              let _index47 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_276_v97).length));
              (_276_v97)[_index47] = _269_v88;
              let _280_v98;
              let _init5 = ((_281_v0) => function (_282_i14) {
                return _dafny.Set.fromElements(_281_v0, _281_v0, false);
              })(_155_v0);
              let _nw47 = Array((new BigNumber(13)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw47.length); _i0_5++) {
                _nw47[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _280_v98 = _nw47;
              let _283_v99;
              _283_v99 = _dafny.Map.Empty.slice().updateUnsafe(_280_v98,_177_v18);
              let _284_v100;
              let _nw48 = Array((new BigNumber(27)).toNumber());
              _nw48[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-94)), function (_285_i13) {
                return new _dafny.CodePoint('w'.codePointAt(0));
              });
              _nw48[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("tjdvjrwn");
              _nw48[(new BigNumber(2)).toNumber()] = ((_155_v0) ? (_177_v18) : (_177_v18));
              _nw48[(new BigNumber(3)).toNumber()] = (((_283_v99).contains(_280_v98)) ? ((_283_v99).get(_280_v98)) : (_177_v18));
              _nw48[(new BigNumber(4)).toNumber()] = _module.__default.fm5(_242_v68.f19, _157_globalState);
              _nw48[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-417)), function (_286_i15) {
                return new _dafny.CodePoint('y'.codePointAt(0));
              });
              _nw48[(new BigNumber(6)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(7)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("vxgk");
              _nw48[(new BigNumber(9)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(10)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vb"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-384)), function (_287_i16) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              }));
              _nw48[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_177_v18, _177_v18);
              _nw48[(new BigNumber(13)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(14)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(15)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(16)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(17)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("a");
              _nw48[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), function (_288_i17) {
                return new _dafny.CodePoint('m'.codePointAt(0));
              });
              _nw48[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_177_v18, _dafny.Seq.UnicodeFromString("kkieswsc"));
              _nw48[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_177_v18, _177_v18);
              _nw48[(new BigNumber(22)).toNumber()] = _177_v18;
              _nw48[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(551)), function (_289_i18) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              });
              _nw48[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("uscna");
              _nw48[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat((_275_v96)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_275_v96).length))], _177_v18);
              _nw48[(new BigNumber(26)).toNumber()] = _177_v18;
              _284_v100 = _nw48;
              let _index48 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_284_v100).length));
              (_284_v100)[_index48] = ((_155_v0) ? (_dafny.Seq.UnicodeFromString("btynhwo")) : (_177_v18));
              let _index49 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_284_v100).length));
              (_284_v100)[_index49] = _177_v18;
              let _rhs51 = _242_v68.f19;
              let _rhs52 = (((_241_v67).contains(new BigNumber((_238_v64).length))) ? ((_241_v67).get(new BigNumber((_238_v64).length))) : (_242_v68.f19));
              let _rhs53 = !(_155_v0);
              let _rhs54 = new BigNumber(((_284_v100)[_module.__default.safeIndex(new BigNumber(571), new BigNumber((_284_v100).length))]).length);
              let _rhs55 = ((_242_v68.f19).multipliedBy(_242_v68.f19)).plus((((_159_v3).contains(_155_v0)) ? ((_159_v3).get(_155_v0)) : (_239_v65.f19)));
              let _lhs43 = _239_v65;
              let _lhs44 = _157_globalState;
              let _lhs45 = _157_globalState;
              let _lhs46 = _157_globalState;
              let _lhs47 = _157_globalState;
              _lhs43.f19 = _rhs51;
              _lhs44.f0 = _rhs52;
              _lhs45.f3 = _rhs53;
              _lhs46.f8 = _rhs54;
              _lhs47.f12 = _rhs55;
            }
            (_157_globalState).f3 = _155_v0;
            let _290_v101;
            _290_v101 = _module.D1.create_DC5((!(_155_v0)) && (_155_v0));
            let _source6 = _290_v101;
            if (_source6.is_DC4) {
              let _291___mcc_h0 = (_source6).cf9;
              let _292_cf9 = _291___mcc_h0;
              _156_v1 = _156_v1;
              _159_v3 = (_159_v3).Difference(_159_v3);
              (_239_v65).f19 = (_dafny.ZERO).minus(_239_v65.f19);
              let _293_v102;
              _293_v102 = new _dafny.CodePoint('s'.codePointAt(0));
              let _294_v103;
              _294_v103 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_238_v64);
              let _rhs56 = ((_155_v0) ? (new _dafny.CodePoint('a'.codePointAt(0))) : (_module.__default.fm7(_155_v0, _239_v65.f19, _157_globalState)));
              let _rhs57 = _155_v0;
              let _rhs58 = (_242_v68).fm4((_dafny.ZERO).minus(_292_cf9.f19), _158_v2, _155_v0, !(_294_v103).equals(_module.__default.fm8(new BigNumber((_177_v18).length), _292_cf9.f19, _155_v0, _157_globalState)), _157_globalState);
              let _lhs48 = _157_globalState;
              let _lhs49 = _157_globalState;
              _293_v102 = _rhs56;
              _lhs48.f3 = _rhs57;
              _lhs49.f8 = _rhs58;
            } else if (_source6.is_DC5) {
              let _295___mcc_h1 = (_source6).cf10;
              let _296_cf10 = _295___mcc_h1;
              (_157_globalState).f3 = _155_v0;
              _241_v67 = (_241_v67).update(_239_v65.f19, _module.__default.safeDivisionInt(_242_v68.f19, _242_v68.f19));
              let _297_v104;
              let _init6 = function (_298_i19) {
                return new _dafny.CodePoint('i'.codePointAt(0));
              };
              let _nw49 = Array((new BigNumber(26)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw49.length); _i0_6++) {
                _nw49[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _297_v104 = _nw49;
              let _index50 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_297_v104).length));
              (_297_v104)[_index50] = new _dafny.CodePoint('i'.codePointAt(0));
              let _299_v105;
              _299_v105 = new _dafny.CodePoint('l'.codePointAt(0));
              let _300_v106;
              _300_v106 = _dafny.Set.fromElements(_299_v105, _299_v105, _299_v105);
              let _301_v107;
              _301_v107 = _dafny.MultiSet.fromElements(_299_v105, _299_v105);
              let _302_v108;
              _302_v108 = _dafny.Seq.of(_177_v18, _module.__default.fm5((((_301_v107).contains(_299_v105)) ? ((_301_v107).get(_299_v105)) : (_242_v68.f19)), _157_globalState));
              let _index51 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_297_v104).length));
              let _rhs59 = !(_296_cf10) || (!(_300_v106).equals(_dafny.Set.fromElements(_299_v105, _299_v105, _299_v105)));
              let _rhs60 = _299_v105;
              let _rhs61 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_156_v1, _156_v1), _dafny.Seq.of(_155_v0));
              let _rhs62 = !_dafny.areEqual(_302_v108, _dafny.Seq.Concat(_302_v108, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("owfqc"))));
              let _lhs50 = _157_globalState;
              let _lhs51 = _297_v104;
              let _lhs52 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_297_v104).length));
              let _lhs53 = _157_globalState;
              _lhs50.f3 = _rhs59;
              _lhs51[_lhs52] = _rhs60;
              _lhs53.f3 = _rhs61;
              _296_cf10 = _rhs62;
              let _303_v110;
              let _nw50 = Array((new BigNumber(3)).toNumber());
              _nw50[(_dafny.ZERO).toNumber()] = new BigNumber((function () {
                let _coll7 = new _dafny.Map();
                for (const _compr_7 of (_238_v64).Elements) {
                  let _304_v109 = _compr_7;
                  if (_dafny.Seq.contains(_238_v64, _304_v109)) {
                    _coll7.push([_module.__default.safeDivisionInt(_304_v109, (_dafny.ZERO).minus(_158_v2)),(_module.D0.create_DC1(_242_v68.f19, new BigNumber(24), true)).dtor_cf1]);
                  }
                }
                return _coll7;
              }()).length);
              _nw50[(_dafny.ONE).toNumber()] = _242_v68.f19;
              _nw50[(new BigNumber(2)).toNumber()] = _239_v65.f19;
              _303_v110 = _nw50;
              let _index52 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_303_v110).length));
              (_303_v110)[_index52] = _158_v2;
              let _305_v111;
              _305_v111 = _dafny.MultiSet.fromElements(_241_v67);
              let _index53 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_303_v110).length));
              (_303_v110)[_index53] = (_238_v64)[_module.__default.safeIndex(new BigNumber((_305_v111).cardinality()), new BigNumber((_238_v64).length))];
            } else if (_source6.is_DC3) {
              let _306___mcc_h2 = (_source6).cf8;
              let _307_cf8 = _306___mcc_h2;
              let _308_v112;
              _308_v112 = _module.D0.create_DC2(!(_155_v0), _155_v0, _242_v68.f19, _155_v0);
              let _309_v113;
              let _nw51 = Array((_dafny.ONE).toNumber());
              _nw51[(_dafny.ZERO).toNumber()] = _308_v112;
              _309_v113 = _nw51;
              let _310_v114;
              _310_v114 = _module.D1.create_DC3(_177_v18);
              let _311_v115;
              _311_v115 = _dafny.Map.Empty.slice().updateUnsafe(_309_v113,_310_v114);
              _311_v115 = ((_311_v115).Merge(_311_v115)).Merge(_311_v115);
              (_157_globalState).f3 = !((_242_v68.f19).isEqualTo((_239_v65).fm3(_157_globalState)));
              (_157_globalState).f3 = _module.__default.fm2(_155_v0, _157_globalState);
              (_157_globalState).f8 = _239_v65.f19;
            } else {
              let _312___mcc_h3 = (_source6).cf11;
              let _313_cf11 = _312___mcc_h3;
              let _314_v116;
              let _nw52 = new _module.C0();
              _nw52.__ctor(_239_v65.f19);
              _314_v116 = _nw52;
              let _315_v117;
              _315_v117 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(true, _157_globalState),_242_v68.f19);
              _315_v117 = (_315_v117).update(true, _242_v68.f19);
              _238_v64 = _238_v64;
              let _316_v118;
              let _init7 = function (_317_i20) {
                return _module.__default.safeDivisionInt(_317_i20, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nffg")).length)));
              };
              let _nw53 = Array((new BigNumber(10)).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw53.length); _i0_7++) {
                _nw53[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _316_v118 = _nw53;
              let _318_v119;
              _318_v119 = _dafny.MultiSet.fromElements(_316_v118);
              let _319_v120;
              _319_v120 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_315_v117).length),_155_v0);
              let _320_v121;
              _320_v121 = _module.D0.create_DC1(new BigNumber((_319_v120).length), _242_v68.f19, _155_v0);
              let _321_v122;
              let _322_v123;
              let _323_v124;
              let _324_v125;
              let _out12;
              let _out13;
              let _out14;
              let _out15;
              let _outcollector3 = _module.__default.m0((new BigNumber((_318_v119).cardinality())).isLessThanOrEqualTo(_239_v65.f19), _155_v0, new BigNumber((_159_v3).cardinality()), ((_dafny.ZERO).minus((_242_v68).fm4(new BigNumber((_177_v18).length), _239_v65.f19, (_320_v121).dtor_cf3, _155_v0, _157_globalState))).isEqualTo(_314_v116.f19), _157_globalState);
              _out12 = _outcollector3[0];
              _out13 = _outcollector3[1];
              _out14 = _outcollector3[2];
              _out15 = _outcollector3[3];
              _321_v122 = _out12;
              _322_v123 = _out13;
              _323_v124 = _out14;
              _324_v125 = _out15;
            }
            (_157_globalState).f0 = _239_v65.f19;
          }
        }
      }
      let _325_v126;
      _325_v126 = _dafny.Set.fromElements(_155_v0, _155_v0);
      let _hi2 = _239_v65.f19;
      for (let _326_i21 = ((!(_155_v0)) ? ((_dafny.ZERO).minus(_158_v2)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_242_v68.f19),new BigNumber((_325_v126).length))).length))); _326_i21.isLessThan(_hi2); _326_i21 = _326_i21.plus(_dafny.ONE)) {
        _177_v18 = _177_v18;
        let _327_v127;
        _327_v127 = _module.D1.create_DC5(_155_v0);
        let _328_v128;
        let _nw54 = Array((new BigNumber(18)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = (_327_v127).dtor_cf10;
        _nw54[(_dafny.ONE).toNumber()] = _155_v0;
        _nw54[(new BigNumber(2)).toNumber()] = false;
        _nw54[(new BigNumber(3)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(4)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(5)).toNumber()] = true;
        _nw54[(new BigNumber(6)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(7)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(8)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(9)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(10)).toNumber()] = !(_155_v0);
        _nw54[(new BigNumber(11)).toNumber()] = _module.__default.fm2(_155_v0, _157_globalState);
        _nw54[(new BigNumber(12)).toNumber()] = true;
        _nw54[(new BigNumber(13)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(14)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(15)).toNumber()] = !(_155_v0);
        _nw54[(new BigNumber(16)).toNumber()] = _155_v0;
        _nw54[(new BigNumber(17)).toNumber()] = _module.__default.fm2(_155_v0, _157_globalState);
        _328_v128 = _nw54;
        let _329_v129;
        let _nw55 = Array((new BigNumber(2)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = ((_155_v0) ? (_328_v128) : (_328_v128));
        _nw55[(_dafny.ONE).toNumber()] = _328_v128;
        _329_v129 = _nw55;
        let _index54 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_329_v129).length));
        (_329_v129)[_index54] = _328_v128;
        let _index55 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_329_v129).length));
        (_329_v129)[_index55] = (_module.D2.create_DC7(_328_v128)).dtor_cf12;
        let _330_v130;
        _330_v130 = _dafny.Seq.of(_159_v3, _159_v3);
        let _331_v131;
        let _nw56 = new _module.C0();
        _nw56.__ctor(new BigNumber((_330_v130).length));
        _331_v131 = _nw56;
        let _332_v132;
        _332_v132 = _module.D2.create_DC8();
        let _source7 = _332_v132;
        if (_source7.is_DC8) {
          let _333_v133;
          let _init8 = function (_334_i22) {
            return (_334_i22).plus(new BigNumber(469));
          };
          let _nw57 = Array((new BigNumber(17)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw57.length); _i0_8++) {
            _nw57[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _333_v133 = _nw57;
          let _index56 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_333_v133).length));
          (_333_v133)[_index56] = new BigNumber((_dafny.Seq.UnicodeFromString("pyjrottw")).length);
          let _335_v134;
          _335_v134 = _dafny.Set.fromElements(new BigNumber((_module.__default.fm5(new BigNumber((_241_v67).length), _157_globalState)).length));
          let _336_v135;
          _336_v135 = _dafny.Map.Empty.slice().updateUnsafe(_242_v68.f19,_155_v0);
          let _337_v136;
          _337_v136 = _module.D3.create_DC11(_336_v135);
          let _index57 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_333_v133).length));
          let _rhs63 = new BigNumber(((_335_v134).Difference(_335_v134)).length);
          let _rhs64 = ((_238_v64)[_module.__default.safeIndex(_331_v131.f19, new BigNumber((_238_v64).length))]).plus(new BigNumber(((_337_v136).dtor_cf17).length));
          let _rhs65 = _239_v65.f19;
          let _lhs54 = _333_v133;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_333_v133).length));
          let _lhs56 = _157_globalState;
          _158_v2 = _rhs63;
          _lhs54[_lhs55] = _rhs64;
          _lhs56.f0 = _rhs65;
          _155_v0 = _155_v0;
          let _338_v137;
          _338_v137 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_155_v0);
          let _339_v138;
          let _340_v139;
          let _341_v140;
          let _342_v141;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0(_155_v0, !(_338_v137).contains(_155_v0), _326_i21, (((_338_v137).contains(_155_v0)) ? ((_338_v137).get(_155_v0)) : (_155_v0)), _157_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _339_v138 = _out16;
          _340_v139 = _out17;
          _341_v140 = _out18;
          _342_v141 = _out19;
          _340_v139 = _177_v18;
        } else if (_source7.is_DC9) {
          let _343___mcc_h4 = (_source7).cf13;
          let _344___mcc_h5 = (_source7).cf14;
          let _345___mcc_h6 = (_source7).cf15;
          let _346_cf15 = _345___mcc_h6;
          let _347_cf14 = _344___mcc_h5;
          let _348_cf13 = _343___mcc_h4;
          (_157_globalState).f0 = (_module.__default.safeModuloInt(new BigNumber(763), _239_v65.f19)).plus((_dafny.ZERO).minus(_331_v131.f19));
          _155_v0 = false;
          let _349_v142;
          _349_v142 = new _dafny.CodePoint('x'.codePointAt(0));
          let _350_v143;
          _350_v143 = _dafny.Map.Empty.slice().updateUnsafe(false,_349_v142);
          _350_v143 = (_350_v143).update(_155_v0, _349_v142);
          let _351_v144;
          _351_v144 = _module.D1.create_DC4(_347_cf14);
          let _pat_let_tv19 = _242_v68;
          let _pat_let_tv20 = _347_cf14;
          let _352_v145;
          let _nw58 = Array((new BigNumber(25)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _module.D1.create_DC4(_331_v131);
          _nw58[(_dafny.ONE).toNumber()] = ((_155_v0) ? (_351_v144) : (_module.D1.create_DC4(_242_v68)));
          _nw58[(new BigNumber(2)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(3)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(4)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(5)).toNumber()] = _module.D1.create_DC4(_239_v65);
          _nw58[(new BigNumber(6)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(7)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(8)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(9)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(10)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(11)).toNumber()] = ((_155_v0) ? (_351_v144) : (function (_pat_let8_0) {
            return function (_353_dt__update__tmp_h1) {
              return function (_pat_let9_0) {
                return function (_354_dt__update_hcf9_h0) {
                  return _module.D1.create_DC4(_354_dt__update_hcf9_h0);
                }(_pat_let9_0);
              }(_pat_let_tv19);
            }(_pat_let8_0);
          }(_351_v144)));
          _nw58[(new BigNumber(12)).toNumber()] = _module.D1.create_DC4(_242_v68);
          _nw58[(new BigNumber(13)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(14)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(15)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(16)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(17)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(18)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(19)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(20)).toNumber()] = _351_v144;
          _nw58[(new BigNumber(21)).toNumber()] = function (_pat_let10_0) {
            return function (_355_dt__update__tmp_h2) {
              return function (_pat_let11_0) {
                return function (_356_dt__update_hcf9_h1) {
                  return _module.D1.create_DC4(_356_dt__update_hcf9_h1);
                }(_pat_let11_0);
              }(_pat_let_tv20);
            }(_pat_let10_0);
          }(_351_v144);
          _nw58[(new BigNumber(22)).toNumber()] = _module.D1.create_DC4(_331_v131);
          _nw58[(new BigNumber(23)).toNumber()] = _module.D1.create_DC4(_242_v68);
          _nw58[(new BigNumber(24)).toNumber()] = _351_v144;
          _352_v145 = _nw58;
          let _index58 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_352_v145).length));
          (_352_v145)[_index58] = _351_v144;
          let _index59 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_352_v145).length));
          let _rhs66 = _351_v144;
          let _rhs67 = (_dafny.ZERO).minus(_326_i21);
          let _rhs68 = (_module.D0.create_DC2(_155_v0, _155_v0, _242_v68.f19, _155_v0)).dtor_cf6;
          let _rhs69 = (_239_v65.f19).multipliedBy(_242_v68.f19);
          let _lhs57 = _352_v145;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_352_v145).length));
          let _lhs59 = _239_v65;
          let _lhs60 = _242_v68;
          let _lhs61 = _157_globalState;
          _lhs57[_lhs58] = _rhs66;
          _lhs59.f19 = _rhs67;
          _lhs60.f19 = _rhs68;
          _lhs61.f2 = _rhs69;
        } else if (_source7.is_DC7) {
          let _357___mcc_h7 = (_source7).cf12;
          let _358_cf12 = _357___mcc_h7;
          let _359_v146;
          let _nw59 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
          _359_v146 = _nw59;
          let _360_v147;
          _360_v147 = _module.D2.create_DC9(_242_v68.f19, _331_v131, _239_v65);
          let _361_v148;
          _361_v148 = _dafny.Seq.of(_360_v147);
          let _362_v149;
          _362_v149 = _dafny.Map.Empty.slice().updateUnsafe(_242_v68.f19,_module.D2.create_DC10((_361_v148)[_module.__default.safeIndex(_331_v131.f19, new BigNumber((_361_v148).length))]));
          let _index60 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_359_v146).length));
          (_359_v146)[_index60] = _362_v149;
          let _index61 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_359_v146).length));
          (_359_v146)[_index61] = (_362_v149).Merge(_362_v149);
          let _363_v150;
          _363_v150 = _dafny.MultiSet.fromElements(_239_v65, _239_v65);
          let _364_v151;
          _364_v151 = _dafny.Map.Empty.slice().updateUnsafe(_363_v150,_326_i21);
          let _365_v152;
          _365_v152 = _dafny.Map.Empty.slice().updateUnsafe(false,_155_v0);
          let _366_v153;
          _366_v153 = _dafny.Seq.of(_242_v68);
          let _367_v154;
          _367_v154 = _module.D4.create_DC14(_366_v153);
          let _368_v155;
          _368_v155 = _dafny.MultiSet.fromElements(_158_v2, _239_v65.f19);
          let _369_v156;
          let _nw60 = Array((new BigNumber(8)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = _364_v151;
          _nw60[(_dafny.ONE).toNumber()] = (_364_v151).Merge(((_364_v151).update(_363_v150, new BigNumber((_365_v152).length))).update(_363_v150, (_238_v64)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_238_v64).length))]));
          _nw60[(new BigNumber(2)).toNumber()] = _364_v151;
          _nw60[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray((_367_v154).dtor_cf21),_326_i21);
          _nw60[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_363_v150,_239_v65.f19)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_363_v150,_326_i21));
          _nw60[(new BigNumber(5)).toNumber()] = (_364_v151).update(_363_v150, (_dafny.ZERO).minus(new BigNumber((_368_v155).cardinality())));
          _nw60[(new BigNumber(6)).toNumber()] = _364_v151;
          _nw60[(new BigNumber(7)).toNumber()] = _364_v151;
          _369_v156 = _nw60;
          let _370_v157;
          _370_v157 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_331_v131.f19);
          let _index62 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_369_v156).length));
          (_369_v156)[_index62] = (_364_v151).update(_363_v150, new BigNumber((_370_v157).length));
          let _371_v158;
          let _init9 = ((_372_v68) => function (_373_i23) {
            return (_373_i23).minus(_372_v68.f19);
          })(_242_v68);
          let _nw61 = Array((new BigNumber(2)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw61.length); _i0_9++) {
            _nw61[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _371_v158 = _nw61;
          let _374_v159;
          _374_v159 = _dafny.Map.Empty.slice().updateUnsafe((_241_v67).Merge(_241_v67),_371_v158);
          let _index63 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_369_v156).length));
          let _rhs70 = new BigNumber((_dafny.Seq.Concat(_238_v64, _238_v64)).length);
          let _rhs71 = _364_v151;
          let _rhs72 = (((_374_v159).contains(function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(333), new BigNumber(713))) {
              let _376_v160 = _compr_9;
              if (((new BigNumber(333)).isLessThanOrEqualTo(_376_v160)) && ((_376_v160).isLessThan(new BigNumber(713)))) {
                _coll9.push([(_376_v160).plus((_dafny.ZERO).minus(new BigNumber((_177_v18).length))),new BigNumber(719)]);
              }
            }
            return _coll9;
          }())) ? ((_374_v159).get(function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(333), new BigNumber(713))) {
              let _375_v160 = _compr_8;
              if (((new BigNumber(333)).isLessThanOrEqualTo(_375_v160)) && ((_375_v160).isLessThan(new BigNumber(713)))) {
                _coll8.push([(_375_v160).plus((_dafny.ZERO).minus(new BigNumber((_177_v18).length))),new BigNumber(719)]);
              }
            }
            return _coll8;
          }())) : (_371_v158));
          let _lhs62 = _157_globalState;
          let _lhs63 = _369_v156;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_369_v156).length));
          _lhs62.f12 = _rhs70;
          _lhs63[_lhs64] = _rhs71;
          _371_v158 = _rhs72;
          (_157_globalState).f8 = _module.__default.safeDivisionInt(new BigNumber(701), _239_v65.f19);
          let _index64 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_328_v128).length));
          (_328_v128)[_index64] = _155_v0;
          let _index65 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_328_v128).length));
          (_328_v128)[_index65] = _155_v0;
        } else {
          let _377___mcc_h8 = (_source7).cf16;
          let _378_cf16 = _377___mcc_h8;
          (_157_globalState).f3 = ((_module.__default.fm2(_155_v0, _157_globalState)) ? (_155_v0) : ((_242_v68.f19).isEqualTo(_331_v131.f19)));
          let _379_v161;
          _379_v161 = new _dafny.CodePoint('e'.codePointAt(0));
          let _380_v162;
          _380_v162 = _module.D5.create_DC17((_177_v18)[_module.__default.safeIndex(new BigNumber((_177_v18).length), new BigNumber((_177_v18).length))]);
          _379_v161 = (_380_v162).dtor_cf23;
          let _381_v163;
          _381_v163 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,true);
          let _382_v164;
          _382_v164 = _dafny.MultiSet.fromElements(new BigNumber((_381_v163).length));
          let _rhs73 = _242_v68;
          let _rhs74 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_379_v161, _379_v161, (_177_v18)[_module.__default.safeIndex((((_382_v164).contains(_239_v65.f19)) ? ((_382_v164).get(_239_v65.f19)) : (new BigNumber((_159_v3).cardinality()))), new BigNumber((_177_v18).length))], ((_155_v0) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (new _dafny.CodePoint('t'.codePointAt(0)))), _379_v161), _module.__default.safeIndex((_239_v65.f19).multipliedBy(_326_i21), new BigNumber((_dafny.Seq.of(_379_v161, _379_v161, (_177_v18)[_module.__default.safeIndex((((_382_v164).contains(_239_v65.f19)) ? ((_382_v164).get(_239_v65.f19)) : (new BigNumber((_159_v3).cardinality()))), new BigNumber((_177_v18).length))], ((_155_v0) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (new _dafny.CodePoint('t'.codePointAt(0)))), _379_v161)).length)), _379_v161), _module.__default.safeIndex(_239_v65.f19, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_379_v161, _379_v161, (_177_v18)[_module.__default.safeIndex((((_382_v164).contains(_239_v65.f19)) ? ((_382_v164).get(_239_v65.f19)) : (new BigNumber((_159_v3).cardinality()))), new BigNumber((_177_v18).length))], ((_155_v0) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (new _dafny.CodePoint('t'.codePointAt(0)))), _379_v161), _module.__default.safeIndex((_239_v65.f19).multipliedBy(_326_i21), new BigNumber((_dafny.Seq.of(_379_v161, _379_v161, (_177_v18)[_module.__default.safeIndex((((_382_v164).contains(_239_v65.f19)) ? ((_382_v164).get(_239_v65.f19)) : (new BigNumber((_159_v3).cardinality()))), new BigNumber((_177_v18).length))], ((_155_v0) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (new _dafny.CodePoint('t'.codePointAt(0)))), _379_v161)).length)), _379_v161)).length)), _379_v161);
          _331_v131 = _rhs73;
          _177_v18 = _rhs74;
          let _383_v165;
          let _nw62 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _383_v165 = _nw62;
          let _384_v166;
          _384_v166 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_155_v0, _155_v0), _156_v1);
          let _index66 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_383_v165).length));
          (_383_v165)[_index66] = new BigNumber((_384_v166).cardinality());
          let _index67 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_383_v165).length));
          (_383_v165)[_index67] = ((_158_v2).multipliedBy(_331_v131.f19)).minus(_326_i21);
        }
      }
      let _385_v167;
      let _nw63 = Array((new BigNumber(28)).toNumber()).fill(false);
      _385_v167 = _nw63;
      let _386_v168;
      _386_v168 = _module.D5.create_DC18(_242_v68.f19, _155_v0, _385_v167, _155_v0);
      let _387_v169;
      _387_v169 = _module.D5.create_DC20(_386_v168);
      let _388_v170;
      _388_v170 = _module.D5.create_DC20(_387_v169);
      let _source8 = _388_v170;
      if (_source8.is_DC18) {
        let _389___mcc_h9 = (_source8).cf24;
        let _390___mcc_h10 = (_source8).cf25;
        let _391___mcc_h11 = (_source8).cf26;
        let _392___mcc_h12 = (_source8).cf27;
        let _393_cf27 = _392___mcc_h12;
        let _394_cf26 = _391___mcc_h11;
        let _395_cf25 = _390___mcc_h10;
        let _396_cf24 = _389___mcc_h9;
        let _index68 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_394_cf26).length));
        (_394_cf26)[_index68] = _395_cf25;
        let _index69 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_394_cf26).length));
        (_394_cf26)[_index69] = _module.__default.fm2(!_dafny.areEqual(_156_v1, _dafny.Seq.of(_395_cf25)), _157_globalState);
        let _index70 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_394_cf26).length));
        let _rhs75 = _239_v65.f19;
        let _rhs76 = !(!(_155_v0) || (_393_cf27));
        let _rhs77 = (new BigNumber((_325_v126).length)).multipliedBy(_242_v68.f19);
        let _rhs78 = _395_cf25;
        let _lhs65 = _157_globalState;
        let _lhs66 = _394_cf26;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_394_cf26).length));
        let _lhs68 = _157_globalState;
        _lhs65.f0 = _rhs75;
        _lhs66[_lhs67] = _rhs76;
        _lhs68.f12 = _rhs77;
        _155_v0 = _rhs78;
        let _397_v171;
        _397_v171 = _dafny.MultiSet.fromElements((_240_v66).Merge(_240_v66));
        _397_v171 = _397_v171;
        let _398_v172;
        _398_v172 = _module.D2.create_DC8();
        let _source9 = _398_v172;
        if (_source9.is_DC8) {
          (_157_globalState).f3 = (_241_v67).equals(((_155_v0) ? (function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of _dafny.IntegerRange(new BigNumber(143), new BigNumber(-7))) {
              let _399_v173 = _compr_10;
              if (((new BigNumber(143)).isLessThanOrEqualTo(_399_v173)) && ((_399_v173).isLessThan(new BigNumber(-7)))) {
                _coll10.push([(_399_v173).multipliedBy(_242_v68.f19),_396_cf24]);
              }
            }
            return _coll10;
          }()) : (_241_v67)));
          (_157_globalState).f0 = new BigNumber(174);
          _177_v18 = (((_394_cf26)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_394_cf26).length))]) ? (_177_v18) : (_177_v18));
          let _400_v174;
          _400_v174 = new _dafny.CodePoint('l'.codePointAt(0));
          _400_v174 = _400_v174;
        } else if (_source9.is_DC9) {
          let _401___mcc_h16 = (_source9).cf13;
          let _402___mcc_h17 = (_source9).cf14;
          let _403___mcc_h18 = (_source9).cf15;
          let _404_cf15 = _403___mcc_h18;
          let _405_cf14 = _402___mcc_h17;
          let _406_cf13 = _401___mcc_h16;
          (_239_v65).f19 = _406_cf13;
          (_239_v65).f19 = _module.__default.safeModuloInt(_239_v65.f19, (_242_v68).fm3(_157_globalState));
          let _407_v175;
          _407_v175 = _dafny.Seq.of(_dafny.Seq.Concat(_156_v1, _156_v1), _156_v1, _156_v1, _dafny.Seq.of(_155_v0, _395_cf25, _395_cf25), _156_v1);
          _407_v175 = _dafny.Seq.Concat(_407_v175, _407_v175);
          let _408_v176;
          let _nw64 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _408_v176 = _nw64;
          let _409_v177;
          _409_v177 = new _dafny.CodePoint('n'.codePointAt(0));
          let _index71 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_408_v176).length));
          (_408_v176)[_index71] = _409_v177;
          let _410_v178;
          _410_v178 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_409_v177,_155_v0),_393_cf27);
          let _411_v179;
          _411_v179 = _module.D1.create_DC5(_395_cf25);
          let _412_v180;
          _412_v180 = _dafny.Map.Empty.slice().updateUnsafe(_409_v177,(_411_v179).dtor_cf10);
          let _index72 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_408_v176).length));
          (_408_v176)[_index72] = _module.__default.fm7((((_410_v178).contains(_412_v180)) ? ((_410_v178).get(_412_v180)) : (!(_393_cf27))), new BigNumber(318), _157_globalState);
        } else if (_source9.is_DC7) {
          let _413___mcc_h19 = (_source9).cf12;
          let _414_cf12 = _413___mcc_h19;
          _155_v0 = (_155_v0) === (_155_v0);
          (_242_v68).f19 = ((_238_v64)[_module.__default.safeIndex(_239_v65.f19, new BigNumber((_238_v64).length))]).multipliedBy(new BigNumber((_dafny.Seq.of(_242_v68.f19)).length));
          let _415_v181;
          let _416_v182;
          let _417_v183;
          let _418_v184;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector5 = _module.__default.m0(_155_v0, _395_cf25, _239_v65.f19, ((false) ? ((_156_v1)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_156_v1).length))]) : (_393_cf27)), _157_globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _out22 = _outcollector5[2];
          _out23 = _outcollector5[3];
          _415_v181 = _out20;
          _416_v182 = _out21;
          _417_v183 = _out22;
          _418_v184 = _out23;
          let _419_v185;
          _419_v185 = new _dafny.CodePoint('t'.codePointAt(0));
          let _420_v186;
          _420_v186 = _dafny.Map.Empty.slice().updateUnsafe(_239_v65,_416_v182);
          let _421_v187;
          _421_v187 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), ((_422_v185) => function (_423_i24) {
            return _422_v185;
          })(_419_v185)), _177_v18);
          let _424_v188;
          let _nw65 = Array((new BigNumber(21)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_177_v18, _416_v182), _module.__default.safeIndex(_415_v181, new BigNumber((_dafny.Seq.Concat(_177_v18, _416_v182)).length)), _419_v185);
          _nw65[(_dafny.ONE).toNumber()] = _416_v182;
          _nw65[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_177_v18, _416_v182);
          _nw65[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("rcptp");
          _nw65[(new BigNumber(4)).toNumber()] = (((_420_v186).contains(_242_v68)) ? ((_420_v186).get(_242_v68)) : (_177_v18));
          _nw65[(new BigNumber(5)).toNumber()] = (_421_v187)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_421_v187).length))];
          _nw65[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bbyhlyjx"), _416_v182);
          _nw65[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_416_v182, _416_v182);
          _nw65[(new BigNumber(8)).toNumber()] = _416_v182;
          _nw65[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm5(_242_v68.f19, _157_globalState), _dafny.Seq.UnicodeFromString("xa"));
          _nw65[(new BigNumber(10)).toNumber()] = _177_v18;
          _nw65[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("tgjqjb");
          _nw65[(new BigNumber(12)).toNumber()] = _177_v18;
          _nw65[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_177_v18, _177_v18);
          _nw65[(new BigNumber(14)).toNumber()] = _416_v182;
          _nw65[(new BigNumber(15)).toNumber()] = _177_v18;
          _nw65[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(576)), ((_425_v185) => function (_426_i25) {
            return _425_v185;
          })(_419_v185));
          _nw65[(new BigNumber(17)).toNumber()] = _416_v182;
          _nw65[(new BigNumber(18)).toNumber()] = _module.__default.fm5(_239_v65.f19, _157_globalState);
          _nw65[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_177_v18, _module.__default.safeIndex(_239_v65.f19, new BigNumber((_177_v18).length)), _419_v185);
          _nw65[(new BigNumber(20)).toNumber()] = _177_v18;
          _424_v188 = _nw65;
          let _index73 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_424_v188).length));
          (_424_v188)[_index73] = _dafny.Seq.UnicodeFromString("tb");
          let _427_v189;
          let _nw66 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
          _427_v189 = _nw66;
          let _428_v190;
          _428_v190 = _dafny.Map.Empty.slice().updateUnsafe(_395_cf25,_427_v189);
          let _index74 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_424_v188).length));
          let _rhs79 = _416_v182;
          let _rhs80 = _414_cf12;
          let _rhs81 = ((false) ? (_393_cf27) : ((_394_cf26)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_394_cf26).length))]));
          let _rhs82 = (((_428_v190).contains(_395_cf25)) ? ((_428_v190).get(_395_cf25)) : (_427_v189));
          let _lhs69 = _424_v188;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_424_v188).length));
          _lhs69[_lhs70] = _rhs79;
          _414_cf12 = _rhs80;
          _155_v0 = _rhs81;
          _427_v189 = _rhs82;
        } else {
          let _429___mcc_h20 = (_source9).cf16;
          let _430_cf16 = _429___mcc_h20;
          (_157_globalState).f3 = _393_cf27;
          (_157_globalState).f0 = _239_v65.f19;
          _395_cf25 = ((_156_v1)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_156_v1).length))]) && (true);
          let _431_v191;
          _431_v191 = _dafny.MultiSet.fromElements(_156_v1, _156_v1, _156_v1);
          _431_v191 = (_431_v191).Intersect((_431_v191).Intersect(_431_v191));
        }
      } else if (_source8.is_DC19) {
        let _432___mcc_h13 = (_source8).cf28;
        let _433_cf28 = _432___mcc_h13;
        let _434_v192;
        _434_v192 = new _dafny.CodePoint('w'.codePointAt(0));
        _434_v192 = (_177_v18)[_module.__default.safeIndex(_158_v2, new BigNumber((_177_v18).length))];
        let _435_v193;
        _435_v193 = _dafny.MultiSet.fromElements(_158_v2);
        if ((_435_v193).IsProperSubsetOf(_435_v193)) {
          _325_v126 = _325_v126;
          let _436_v194;
          let _nw67 = new _module.C0();
          _nw67.__ctor(_242_v68.f19);
          _436_v194 = _nw67;
          let _437_v195;
          _437_v195 = _dafny.Set.fromElements(_177_v18);
          _437_v195 = _module.__default.fm9((_dafny.ZERO).minus(_242_v68.f19), new _dafny.CodePoint('u'.codePointAt(0)), _module.__default.fm2(_155_v0, _157_globalState), _157_globalState);
          _239_v65 = _239_v65;
          _433_cf28 = _155_v0;
        } else {
          (_157_globalState).f3 = ((_433_cf28) || (_155_v0)) || (false);
          (_157_globalState).f3 = _433_cf28;
          let _438_v196;
          let _439_v197;
          let _440_v198;
          let _441_v199;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector6 = _module.__default.m0(_433_cf28, !_dafny.Seq.contains(_238_v64, _239_v65.f19), _239_v65.f19, (_242_v68.f19).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_242_v68.f19),_434_v192)).length)), _157_globalState);
          _out24 = _outcollector6[0];
          _out25 = _outcollector6[1];
          _out26 = _outcollector6[2];
          _out27 = _outcollector6[3];
          _438_v196 = _out24;
          _439_v197 = _out25;
          _440_v198 = _out26;
          _441_v199 = _out27;
          let _442_v200;
          let _nw68 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _442_v200 = _nw68;
          let _index75 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_442_v200).length));
          (_442_v200)[_index75] = _241_v67;
          let _443_v201;
          _443_v201 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,(_241_v67).Merge(_241_v67));
          let _index76 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_442_v200).length));
          (_442_v200)[_index76] = (((_443_v201).contains(_438_v196)) ? ((_443_v201).get(_438_v196)) : (_241_v67));
          _385_v167 = _385_v167;
        }
        let _444_v202;
        _444_v202 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_dafny.Set.fromElements(_433_cf28));
        let _445_v203;
        _445_v203 = _module.D0.create_DC1(new BigNumber(((((_444_v202).contains(true)) ? ((_444_v202).get(true)) : (_325_v126))).length), _242_v68.f19, _155_v0);
        let _446_v204;
        let _nw69 = new _module.C0();
        _nw69.__ctor((_445_v203).dtor_cf1);
        _446_v204 = _nw69;
        let _447_v205;
        let _nw70 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _447_v205 = _nw70;
        let _448_v206;
        _448_v206 = _dafny.MultiSet.fromElements(_241_v67);
        let _449_v207;
        _449_v207 = _dafny.Map.Empty.slice().updateUnsafe(_447_v205,_448_v206);
        _449_v207 = (_449_v207).update(_447_v205, (_448_v206).update(_module.__default.fm10(_157_globalState), _module.__default.abs(_158_v2)));
      } else if (_source8.is_DC17) {
        let _450___mcc_h14 = (_source8).cf23;
        let _451_cf23 = _450___mcc_h14;
        _158_v2 = (_242_v68.f19).minus(_242_v68.f19);
        let _452_v208;
        _452_v208 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,true);
        _155_v0 = (((_452_v208).contains(new BigNumber((_159_v3).cardinality()))) ? ((_452_v208).get(new BigNumber((_159_v3).cardinality()))) : (_155_v0));
        let _453_v209;
        _453_v209 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_242_v68.f19));
        if ((_module.D0.create_DC1(_158_v2, (((_453_v209).contains(new BigNumber((_177_v18).length))) ? ((_453_v209).get(new BigNumber((_177_v18).length))) : (_239_v65.f19)), _155_v0)).dtor_cf3) {
          _238_v64 = _238_v64;
          (_157_globalState).f3 = _155_v0;
          let _index77 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_385_v167).length));
          (_385_v167)[_index77] = !(_155_v0);
          let _454_v210;
          let _nw71 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _454_v210 = _nw71;
          let _index78 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_454_v210).length));
          (_454_v210)[_index78] = _module.__default.safeModuloInt(_239_v65.f19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_455_cf23) => function (_456_i26) {
            return _455_cf23;
          })(_451_cf23))).length));
          let _457_v211;
          _457_v211 = _dafny.Set.fromElements(_242_v68);
          let _index79 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_385_v167).length));
          let _index80 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_454_v210).length));
          let _rhs83 = true;
          let _rhs84 = !(_module.__default.fm2(_155_v0, _157_globalState));
          let _rhs85 = _module.__default.safeModuloInt(_239_v65.f19, new BigNumber(-237));
          let _rhs86 = _dafny.Seq.update(_dafny.Seq.update(_177_v18, _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_177_v18).length), new BigNumber(-470))), new BigNumber((_177_v18).length)), _451_cf23), _module.__default.safeIndex(new BigNumber((_453_v209).cardinality()), new BigNumber((_dafny.Seq.update(_177_v18, _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_177_v18).length), new BigNumber(-470))), new BigNumber((_177_v18).length)), _451_cf23)).length)), (_177_v18)[_module.__default.safeIndex(new BigNumber((_453_v209).cardinality()), new BigNumber((_177_v18).length))]);
          let _rhs87 = (_457_v211).IsDisjointFrom(_457_v211);
          let _lhs71 = _385_v167;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_385_v167).length));
          let _lhs73 = _157_globalState;
          let _lhs74 = _454_v210;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_454_v210).length));
          _lhs71[_lhs72] = _rhs83;
          _lhs73.f3 = _rhs84;
          _lhs74[_lhs75] = _rhs85;
          _177_v18 = _rhs86;
          _155_v0 = _rhs87;
          (_157_globalState).f3 = !(((_155_v0) ? (_155_v0) : ((_385_v167)[_module.__default.safeIndex(new BigNumber(647), new BigNumber((_385_v167).length))])));
          let _458_v212;
          _458_v212 = _module.D2.create_DC9(_239_v65.f19, _242_v68, _242_v68);
          let _459_v213;
          _459_v213 = _module.D2.create_DC10(_458_v212);
          let _pat_let_tv21 = _458_v212;
          _459_v213 = function (_pat_let12_0) {
            return function (_460_dt__update__tmp_h3) {
              return function (_pat_let13_0) {
                return function (_461_dt__update_hcf16_h0) {
                  return _module.D2.create_DC10(_461_dt__update_hcf16_h0);
                }(_pat_let13_0);
              }(_pat_let_tv21);
            }(_pat_let12_0);
          }(_459_v213);
        } else {
          let _462_v214;
          let _init10 = ((_463_v2) => function (_464_i27) {
            return _module.__default.safeDivisionInt(_464_i27, (_dafny.ZERO).minus(_463_v2));
          })(_158_v2);
          let _nw72 = Array((new BigNumber(2)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw72.length); _i0_10++) {
            _nw72[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _462_v214 = _nw72;
          _462_v214 = _462_v214;
          let _465_v215;
          let _466_v216;
          let _467_v217;
          let _468_v218;
          let _out28;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector7 = _module.__default.m0(false, _155_v0, new BigNumber(((_452_v208).Merge(_452_v208)).length), false, _157_globalState);
          _out28 = _outcollector7[0];
          _out29 = _outcollector7[1];
          _out30 = _outcollector7[2];
          _out31 = _outcollector7[3];
          _465_v215 = _out28;
          _466_v216 = _out29;
          _467_v217 = _out30;
          _468_v218 = _out31;
          let _469_v219;
          _469_v219 = _dafny.Set.fromElements(_467_v217, new BigNumber((_466_v216).length));
          let _470_v220;
          _470_v220 = _dafny.Seq.of(_module.__default.fm5(_239_v65.f19, _157_globalState), _177_v18);
          (_157_globalState).f3 = (new BigNumber((_469_v219).length)).isLessThanOrEqualTo(_module.__default.fm0(_242_v68.f19, _451_cf23, _470_v220, _239_v65.f19, _157_globalState));
          _241_v67 = (_241_v67).update(new BigNumber((_177_v18).length), _239_v65.f19);
          _385_v167 = _385_v167;
        }
        let _471_v221;
        let _init11 = ((_472_v65) => function (_473_i28) {
          return (_473_i28).plus(_472_v65.f19);
        })(_239_v65);
        let _nw73 = Array((new BigNumber(24)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw73.length); _i0_11++) {
          _nw73[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _471_v221 = _nw73;
        _471_v221 = _471_v221;
      } else {
        let _474___mcc_h15 = (_source8).cf29;
        let _475_cf29 = _474___mcc_h15;
        let _476_v222;
        _476_v222 = _dafny.MultiSet.fromElements(_239_v65);
        _476_v222 = _476_v222;
        (_157_globalState).f3 = _dafny.areEqual(_177_v18, _dafny.Seq.UnicodeFromString("eevhh"));
        let _477_v223;
        _477_v223 = _module.D2.create_DC9(_239_v65.f19, _242_v68, _242_v68);
        let _nw74 = new _module.C0();
        _nw74.__ctor((_477_v223).dtor_cf13);
        _242_v68 = _nw74;
        let _478_v224;
        _478_v224 = new _dafny.CodePoint('i'.codePointAt(0));
        let _479_v225;
        _479_v225 = _dafny.Seq.of(_177_v18, _dafny.Seq.of(_478_v224, _478_v224));
        let _480_v226;
        _480_v226 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), ((_481_v224) => function (_482_i29) {
          return _481_v224;
        })(_478_v224)),_242_v68.f19);
        _478_v224 = _module.__default.fm7(_155_v0, _module.__default.fm0(_239_v65.f19, _478_v224, _479_v225, (((_480_v226).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(889)), ((_485_v224) => function (_486_i30) {
          return _485_v224;
        })(_478_v224)))) ? ((_480_v226).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(889)), ((_483_v224) => function (_484_i30) {
          return _483_v224;
        })(_478_v224)))) : (new BigNumber((_dafny.Seq.of(_155_v0, _155_v0)).length))), _157_globalState), _157_globalState);
      }
      let _487_v227;
      _487_v227 = new _dafny.CodePoint('s'.codePointAt(0));
      let _488_v228;
      let _nw75 = Array((new BigNumber(4)).toNumber());
      _nw75[(_dafny.ZERO).toNumber()] = _487_v227;
      _nw75[(_dafny.ONE).toNumber()] = _487_v227;
      _nw75[(new BigNumber(2)).toNumber()] = (_177_v18)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_177_v18).length))];
      _nw75[(new BigNumber(3)).toNumber()] = ((false) ? (_487_v227) : (_487_v227));
      _488_v228 = _nw75;
      let _index81 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length));
      (_488_v228)[_index81] = _module.__default.fm7(true, new BigNumber(157), _157_globalState);
      let _index82 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length));
      (_488_v228)[_index82] = _487_v227;
      _158_v2 = _158_v2;
      let _hi3 = _158_v2;
      for (let _489_i31 = (new BigNumber((_238_v64).length)).plus(_158_v2); _489_i31.isLessThan(_hi3); _489_i31 = _489_i31.plus(_dafny.ONE)) {
        _177_v18 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-856)), ((_490_v228) => function (_491_i32) {
          return (_490_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_490_v228).length))];
        })(_488_v228));
        if (_155_v0) {
          let _492_v229;
          _492_v229 = _dafny.Seq.of(_177_v18);
          let _493_v230;
          _493_v230 = _module.D0.create_DC1(_module.__default.fm0(_242_v68.f19, new _dafny.CodePoint('m'.codePointAt(0)), _492_v229, _158_v2, _157_globalState), new BigNumber((_159_v3).cardinality()), _155_v0);
          let _494_v231;
          _494_v231 = _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC17((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]),_493_v230);
          let _495_v233;
          _495_v233 = _module.D5.create_DC17((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]);
          let _pat_let_tv22 = _487_v227;
          let _496_v234;
          _496_v234 = _dafny.Set.fromElements(_495_v233, _495_v233, function (_pat_let14_0) {
            return function (_497_dt__update__tmp_h4) {
              return function (_pat_let15_0) {
                return function (_498_dt__update_hcf23_h0) {
                  return _module.D5.create_DC17(_498_dt__update_hcf23_h0);
                }(_pat_let15_0);
              }(_pat_let_tv22);
            }(_pat_let14_0);
          }(_495_v233));
          let _499_v235;
          _499_v235 = _dafny.Seq.of(_494_v231, function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_496_v234).Elements) {
              let _500_v232 = _compr_11;
              if ((_496_v234).contains(_500_v232)) {
                _coll11.push([_500_v232,_493_v230]);
              }
            }
            return _coll11;
          }(), (_494_v231).Merge(_494_v231));
          _499_v235 = _dafny.Seq.Concat(_dafny.Seq.Concat(_499_v235, _dafny.Seq.Create(_module.__default.abs(new BigNumber(204)), ((_501_v233, _502_v230) => function (_503_i33) {
            return _dafny.Map.Empty.slice().updateUnsafe(_501_v233,_502_v230);
          })(_495_v233, _493_v230))), _dafny.Seq.update(_dafny.Seq.Concat(_499_v235, _499_v235), _module.__default.safeIndex(new BigNumber(583), new BigNumber((_dafny.Seq.Concat(_499_v235, _499_v235)).length)), _dafny.Map.Empty.slice().updateUnsafe(_495_v233,_493_v230)));
          (_239_v65).f19 = _158_v2;
          let _504_v236;
          _504_v236 = _module.D5.create_DC18(_158_v2, _155_v0, _385_v167, _155_v0);
          let _505_v237;
          _505_v237 = _dafny.Map.Empty.slice().updateUnsafe((_504_v236).dtor_cf26,((_155_v0) ? (_155_v0) : (true)));
          let _506_v238;
          _506_v238 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), ((_507_v68) => function (_508_i34) {
            return _507_v68.f19;
          })(_242_v68)));
          let _509_v239;
          _509_v239 = _dafny.MultiSet.fromElements((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))], _487_v227);
          let _510_v240;
          let _init12 = function (_511_i35) {
            return _dafny.Seq.UnicodeFromString("fvlqdof");
          };
          let _nw76 = Array((new BigNumber(18)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw76.length); _i0_12++) {
            _nw76[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _510_v240 = _nw76;
          let _512_v241;
          _512_v241 = _dafny.Map.Empty.slice().updateUnsafe(_510_v240,_155_v0);
          let _index83 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length));
          let _rhs88 = _dafny.Map.Empty.slice().updateUnsafe(_385_v167,(_156_v1)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_156_v1).length))]);
          let _rhs89 = ((_155_v0) ? (new _dafny.CodePoint('h'.codePointAt(0))) : ((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]));
          let _rhs90 = _dafny.areEqual((((_506_v238).contains(_155_v0)) ? ((_506_v238).get(_155_v0)) : (_dafny.Seq.of((_dafny.ZERO).minus(_242_v68.f19)))), _dafny.Seq.update(_dafny.Seq.Concat(_238_v64, _238_v64), _module.__default.safeIndex((_238_v64)[_module.__default.safeIndex((_dafny.ZERO).minus((((_509_v239).contains((_177_v18)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_177_v18).length))])) ? ((_509_v239).get((_177_v18)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_177_v18).length))])) : ((_dafny.ZERO).minus(_489_i31)))), new BigNumber((_238_v64).length))], new BigNumber((_dafny.Seq.Concat(_238_v64, _238_v64)).length)), (_dafny.ZERO).minus(new BigNumber((_159_v3).cardinality()))));
          let _rhs91 = (((_512_v241).contains(_510_v240)) ? ((_512_v241).get(_510_v240)) : (_155_v0));
          let _lhs76 = _488_v228;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length));
          let _lhs78 = _157_globalState;
          let _lhs79 = _157_globalState;
          _505_v237 = _rhs88;
          _lhs76[_lhs77] = _rhs89;
          _lhs78.f3 = _rhs90;
          _lhs79.f3 = _rhs91;
          (_157_globalState).f8 = _239_v65.f19;
          let _513_v242;
          _513_v242 = _dafny.MultiSet.fromElements(_239_v65.f19);
          let _514_v243;
          _514_v243 = _module.D0.create_DC0((_242_v68).fm4(new BigNumber((_238_v64).length), _242_v68.f19, _155_v0, _155_v0, _157_globalState));
          let _515_v244;
          let _nw77 = Array((new BigNumber(8)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = (_238_v64)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_238_v64).length))];
          _nw77[(_dafny.ONE).toNumber()] = (_242_v68).fm3(_157_globalState);
          _nw77[(new BigNumber(2)).toNumber()] = _158_v2;
          _nw77[(new BigNumber(3)).toNumber()] = new BigNumber((_159_v3).cardinality());
          _nw77[(new BigNumber(4)).toNumber()] = new BigNumber((_177_v18).length);
          _nw77[(new BigNumber(5)).toNumber()] = _239_v65.f19;
          _nw77[(new BigNumber(6)).toNumber()] = new BigNumber(((_dafny.MultiSet.FromArray(_238_v64)).Difference(_513_v242)).cardinality());
          _nw77[(new BigNumber(7)).toNumber()] = ((_155_v0) ? (_489_i31) : ((_514_v243).dtor_cf0));
          _515_v244 = _nw77;
          let _index84 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_515_v244).length));
          (_515_v244)[_index84] = _158_v2;
          let _516_v245;
          _516_v245 = _dafny.MultiSet.fromElements(_242_v68);
          let _517_v246;
          _517_v246 = _dafny.Map.Empty.slice().updateUnsafe(_516_v245,_489_i31);
          let _index85 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_515_v244).length));
          (_515_v244)[_index85] = (_dafny.ZERO).minus((((_517_v246).contains(_516_v245)) ? ((_517_v246).get(_516_v245)) : (_module.__default.safeModuloInt(new BigNumber((_238_v64).length), _158_v2))));
        } else {
          let _518_v247;
          let _nw78 = Array((new BigNumber(10)).toNumber());
          _nw78[(_dafny.ZERO).toNumber()] = _239_v65.f19;
          _nw78[(_dafny.ONE).toNumber()] = _489_i31;
          _nw78[(new BigNumber(2)).toNumber()] = _239_v65.f19;
          _nw78[(new BigNumber(3)).toNumber()] = _242_v68.f19;
          _nw78[(new BigNumber(4)).toNumber()] = (new BigNumber(586)).plus(_242_v68.f19);
          _nw78[(new BigNumber(5)).toNumber()] = _239_v65.f19;
          _nw78[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_156_v1, _156_v1)).length);
          _nw78[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.of(_489_i31, _239_v65.f19)).length);
          _nw78[(new BigNumber(8)).toNumber()] = _239_v65.f19;
          _nw78[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(_239_v65.f19, _242_v68.f19);
          _518_v247 = _nw78;
          let _index86 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_518_v247).length));
          (_518_v247)[_index86] = (_dafny.ZERO).minus(_239_v65.f19);
          let _index87 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_518_v247).length));
          (_518_v247)[_index87] = _module.__default.safeModuloInt(_489_i31, (_dafny.ZERO).minus((_489_i31).plus(new BigNumber(-138))));
          let _index88 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_385_v167).length));
          (_385_v167)[_index88] = _155_v0;
          let _519_v248;
          _519_v248 = _dafny.Seq.of(_239_v65);
          let _index89 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_385_v167).length));
          (_385_v167)[_index89] = (_156_v1)[_module.__default.safeIndex(new BigNumber((_519_v248).length), new BigNumber((_156_v1).length))];
          (_157_globalState).f8 = (new BigNumber((_dafny.Seq.of(_155_v0)).length)).plus((_238_v64)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_238_v64).length))]);
          (_157_globalState).f3 = !((((_385_v167)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_385_v167).length))]) ? ((_385_v167)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_385_v167).length))]) : ((_155_v0) && (!((_385_v167)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_385_v167).length))])))));
          let _520_v249;
          let _nw79 = new _module.C0();
          _nw79.__ctor(_242_v68.f19);
          _520_v249 = _nw79;
        }
        let _521_v250;
        _521_v250 = _dafny.Map.Empty.slice().updateUnsafe(_242_v68.f19,(_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]);
        let _522_v251;
        _522_v251 = _module.D5.create_DC18(new BigNumber((_521_v250).length), _module.__default.fm2(_155_v0, _157_globalState), _385_v167, !(_155_v0));
        let _source10 = _522_v251;
        if (_source10.is_DC18) {
          let _523___mcc_h21 = (_source10).cf24;
          let _524___mcc_h22 = (_source10).cf25;
          let _525___mcc_h23 = (_source10).cf26;
          let _526___mcc_h24 = (_source10).cf27;
          let _527_cf27 = _526___mcc_h24;
          let _528_cf26 = _525___mcc_h23;
          let _529_cf25 = _524___mcc_h22;
          let _530_cf24 = _523___mcc_h21;
          let _531_v252;
          _531_v252 = _dafny.Map.Empty.slice().updateUnsafe(_158_v2,_527_cf27);
          _531_v252 = (_531_v252).update((_239_v65).fm3(_157_globalState), _155_v0);
          let _index90 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_528_cf26).length));
          (_528_cf26)[_index90] = _529_cf25;
          let _index91 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_385_v167).length));
          (_385_v167)[_index91] = _module.__default.fm2(true, _157_globalState);
          let _532_v253;
          let _nw80 = Array((new BigNumber(7)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = _177_v18;
          _nw80[(_dafny.ONE).toNumber()] = _177_v18;
          _nw80[(new BigNumber(2)).toNumber()] = _177_v18;
          _nw80[(new BigNumber(3)).toNumber()] = _177_v18;
          _nw80[(new BigNumber(4)).toNumber()] = _177_v18;
          _nw80[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("rqsphj");
          _nw80[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), ((_533_v227) => function (_534_i36) {
            return _533_v227;
          })(_487_v227));
          _532_v253 = _nw80;
          let _535_v254;
          _535_v254 = _module.D3.create_DC12(_532_v253, _529_cf25);
          let _pat_let_tv23 = _527_cf27;
          let _index92 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_528_cf26).length));
          let _index93 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_385_v167).length));
          let _rhs92 = (((_531_v252).contains((_238_v64)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_238_v64).length))])) ? ((_531_v252).get((_238_v64)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_238_v64).length))])) : (_529_cf25));
          let _rhs93 = (_module.__default.fm2(!(_155_v0), _157_globalState)) === (!_dafny.Seq.contains(_177_v18, _487_v227));
          let _rhs94 = (_239_v65).fm4(_242_v68.f19, (_239_v65).fm4(_239_v65.f19, _242_v68.f19, (function (_pat_let16_0) {
            return function (_536_dt__update__tmp_h5) {
              return function (_pat_let17_0) {
                return function (_537_dt__update_hcf19_h0) {
                  return _module.D3.create_DC12((_536_dt__update__tmp_h5).dtor_cf18, _537_dt__update_hcf19_h0);
                }(_pat_let17_0);
              }(_pat_let_tv23);
            }(_pat_let16_0);
          }(_535_v254)).dtor_cf19, _155_v0, _157_globalState), _529_cf25, _155_v0, _157_globalState);
          let _rhs95 = _155_v0;
          let _rhs96 = _529_cf25;
          let _lhs80 = _528_cf26;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_528_cf26).length));
          let _lhs82 = _242_v68;
          let _lhs83 = _385_v167;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_385_v167).length));
          _lhs80[_lhs81] = _rhs92;
          _155_v0 = _rhs93;
          _lhs82.f19 = _rhs94;
          _lhs83[_lhs84] = _rhs95;
          _529_cf25 = _rhs96;
          (_157_globalState).f2 = _239_v65.f19;
          _240_v66 = _240_v66;
        } else if (_source10.is_DC19) {
          let _538___mcc_h25 = (_source10).cf28;
          let _539_cf28 = _538___mcc_h25;
          let _540_v255;
          _540_v255 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(!(_539_cf28), _157_globalState),_242_v68.f19);
          let _541_v256;
          let _nw81 = Array((new BigNumber(13)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_489_i31, _242_v68.f19);
          _nw81[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_489_i31);
          _nw81[(new BigNumber(2)).toNumber()] = new BigNumber(777);
          _nw81[(new BigNumber(3)).toNumber()] = _239_v65.f19;
          _nw81[(new BigNumber(4)).toNumber()] = (_489_i31).minus(_489_i31);
          _nw81[(new BigNumber(5)).toNumber()] = _239_v65.f19;
          _nw81[(new BigNumber(6)).toNumber()] = (((_540_v255).contains(_155_v0)) ? ((_540_v255).get(_155_v0)) : (_242_v68.f19));
          _nw81[(new BigNumber(7)).toNumber()] = (_239_v65).fm4(_242_v68.f19, (_dafny.ZERO).minus(_242_v68.f19), _155_v0, _155_v0, _157_globalState);
          _nw81[(new BigNumber(8)).toNumber()] = _158_v2;
          _nw81[(new BigNumber(9)).toNumber()] = new BigNumber(324);
          _nw81[(new BigNumber(10)).toNumber()] = _239_v65.f19;
          _nw81[(new BigNumber(11)).toNumber()] = (new BigNumber(55)).plus(_239_v65.f19);
          _nw81[(new BigNumber(12)).toNumber()] = ((_155_v0) ? (_239_v65.f19) : (_239_v65.f19));
          _541_v256 = _nw81;
          let _index94 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_541_v256).length));
          (_541_v256)[_index94] = new BigNumber(((_540_v255).Merge(_540_v255)).length);
          let _index95 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_541_v256).length));
          (_541_v256)[_index95] = _242_v68.f19;
          let _542_v257;
          _542_v257 = _dafny.Map.Empty.slice().updateUnsafe(_239_v65.f19,_539_cf28);
          _542_v257 = (_542_v257).update(new BigNumber(223), ((_539_cf28) ? (_539_cf28) : (_155_v0)));
          let _index96 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_385_v167).length));
          (_385_v167)[_index96] = (_242_v68.f19).isLessThan(_242_v68.f19);
          let _index97 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_385_v167).length));
          (_385_v167)[_index97] = (true) === (_539_cf28);
          let _543_v258;
          let _nw82 = Array((new BigNumber(8)).toNumber()).fill([]);
          _543_v258 = _nw82;
          let _nw83 = Array((new BigNumber(22)).toNumber()).fill([]);
          _543_v258 = _nw83;
        } else if (_source10.is_DC17) {
          let _544___mcc_h26 = (_source10).cf23;
          let _545_cf23 = _544___mcc_h26;
          let _546_v259;
          _546_v259 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("j"), _module.__default.safeIndex((_dafny.ZERO).minus(_242_v68.f19), new BigNumber((_dafny.Seq.UnicodeFromString("j")).length)), new _dafny.CodePoint('b'.codePointAt(0))),!(_155_v0));
          _546_v259 = (function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_546_v259).Keys.Elements) {
              let _547_v260 = _compr_12;
              if ((_546_v259).contains(_547_v260)) {
                _coll12.push([_547_v260,_155_v0]);
              }
            }
            return _coll12;
          }()).Merge((_546_v259).update(_dafny.Seq.UnicodeFromString("mnlnnbyv"), !(!(_155_v0))));
          let _548_v261;
          let _549_v262;
          let _550_v263;
          let _551_v264;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector8 = _module.__default.m0(_dafny.Seq.contains(_238_v64, _239_v65.f19), _155_v0, _489_i31, !(_155_v0), _157_globalState);
          _out32 = _outcollector8[0];
          _out33 = _outcollector8[1];
          _out34 = _outcollector8[2];
          _out35 = _outcollector8[3];
          _548_v261 = _out32;
          _549_v262 = _out33;
          _550_v263 = _out34;
          _551_v264 = _out35;
          let _552_v265;
          _552_v265 = _dafny.Map.Empty.slice().updateUnsafe(_548_v261,_155_v0);
          let _553_v266;
          _553_v266 = _dafny.Map.Empty.slice().updateUnsafe((_552_v265).update(_158_v2, true),(!(true)) && (_155_v0));
          _553_v266 = (_553_v266).Merge(_553_v266);
          let _554_v267;
          _554_v267 = _dafny.Set.fromElements(_385_v167);
          _554_v267 = ((_554_v267).Difference(_554_v267)).Union(_554_v267);
        } else {
          let _555___mcc_h27 = (_source10).cf29;
          let _556_cf29 = _555___mcc_h27;
          let _557_v268;
          let _558_v269;
          let _559_v270;
          let _560_v271;
          let _out36;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector9 = _module.__default.m0(!_dafny.areEqual(_module.__default.fm5((_dafny.ZERO).minus(_489_i31), _157_globalState), _dafny.Seq.UnicodeFromString("pmxvhnl")), _module.__default.fm2(_155_v0, _157_globalState), _239_v65.f19, _155_v0, _157_globalState);
          _out36 = _outcollector9[0];
          _out37 = _outcollector9[1];
          _out38 = _outcollector9[2];
          _out39 = _outcollector9[3];
          _557_v268 = _out36;
          _558_v269 = _out37;
          _559_v270 = _out38;
          _560_v271 = _out39;
          let _561_v272;
          _561_v272 = _module.D2.create_DC9(_239_v65.f19, _242_v68, _242_v68);
          let _562_v273;
          _562_v273 = _module.D2.create_DC10(_561_v272);
          let _563_v274;
          _563_v274 = _dafny.Map.Empty.slice().updateUnsafe(_239_v65.f19,_155_v0);
          let _564_v275;
          _564_v275 = _module.D3.create_DC11(_563_v274);
          let _565_v276;
          _565_v276 = _dafny.MultiSet.fromElements(_564_v275, _module.D3.create_DC11(_563_v274));
          let _rhs97 = _562_v273;
          let _rhs98 = (((_155_v0) || (_155_v0)) ? (new BigNumber(676)) : (_559_v270));
          let _rhs99 = (_565_v276).update(_564_v275, _module.__default.abs((_dafny.ZERO).minus(_239_v65.f19)));
          let _lhs85 = _239_v65;
          _562_v273 = _rhs97;
          _lhs85.f19 = _rhs98;
          _565_v276 = _rhs99;
          _239_v65 = ((_155_v0) ? (((_155_v0) ? (_239_v65) : (_239_v65))) : (_239_v65));
          let _566_v277;
          let _init13 = function (_567_i37) {
            return (_567_i37).multipliedBy(new BigNumber(936));
          };
          let _nw84 = Array((new BigNumber(25)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw84.length); _i0_13++) {
            _nw84[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _566_v277 = _nw84;
          let _index98 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_566_v277).length));
          (_566_v277)[_index98] = _158_v2;
          let _index99 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_566_v277).length));
          (_566_v277)[_index99] = _242_v68.f19;
          let _568_v278;
          _568_v278 = _dafny.Map.Empty.slice().updateUnsafe(_558_v269,_155_v0);
          let _569_v279;
          _569_v279 = _dafny.Set.fromElements((_238_v64)[_module.__default.safeIndex(new BigNumber((_568_v278).length), new BigNumber((_238_v64).length))]);
          let _570_v280;
          _570_v280 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,(_158_v2).minus(_557_v268));
          let _index100 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_566_v277).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_566_v277).length));
          let _rhs100 = _158_v2;
          let _rhs101 = (((_570_v280).contains(true)) ? ((_570_v280).get(true)) : (_242_v68.f19));
          let _rhs102 = _155_v0;
          let _rhs103 = _489_i31;
          let _rhs104 = _569_v279;
          let _lhs86 = _566_v277;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_566_v277).length));
          let _lhs88 = _157_globalState;
          let _lhs89 = _157_globalState;
          let _lhs90 = _566_v277;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_566_v277).length));
          _lhs86[_lhs87] = _rhs100;
          _lhs88.f8 = _rhs101;
          _lhs89.f3 = _rhs102;
          _lhs90[_lhs91] = _rhs103;
          _569_v279 = _rhs104;
        }
        (_157_globalState).f8 = (_242_v68.f19).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), ((_571_v227) => function (_572_i38) {
          return _571_v227;
        })(_487_v227))).length));
      }
      let _573_v281;
      _573_v281 = _module.D0.create_DC1(new BigNumber((_dafny.MultiSet.fromElements(_155_v0)).cardinality()), _158_v2, !(_155_v0));
      let _source11 = ((!(_155_v0)) ? (((_155_v0) ? (_573_v281) : (_573_v281))) : (_573_v281));
      if (_source11.is_DC1) {
        let _574___mcc_h28 = (_source11).cf1;
        let _575___mcc_h29 = (_source11).cf2;
        let _576___mcc_h30 = (_source11).cf3;
        let _577_cf3 = _576___mcc_h30;
        let _578_cf2 = _575___mcc_h29;
        let _579_cf1 = _574___mcc_h28;
        let _index102 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_385_v167).length));
        (_385_v167)[_index102] = _577_cf3;
        let _index103 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_385_v167).length));
        (_385_v167)[_index103] = _577_cf3;
        if (_577_cf3) {
          let _580_v282;
          let _nw85 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _580_v282 = _nw85;
          let _index104 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_580_v282).length));
          (_580_v282)[_index104] = _dafny.Map.Empty.slice().updateUnsafe((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))],_155_v0);
          let _581_v283;
          _581_v283 = _dafny.Seq.of(_242_v68, _242_v68);
          let _582_v284;
          _582_v284 = _dafny.Map.Empty.slice().updateUnsafe((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))],((_dafny.ZERO).minus(_579_cf1)).isLessThan(new BigNumber((_581_v283).length)));
          let _index105 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_580_v282).length));
          (_580_v282)[_index105] = _582_v284;
          (_242_v68).f19 = _module.__default.safeDivisionInt(_239_v65.f19, _239_v65.f19);
          (_157_globalState).f3 = _155_v0;
          let _583_v285;
          let _init14 = ((_584_v68, _585_v0, _586_v167) => function (_587_i39) {
            return ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_585_v0,(_586_v167)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_586_v167).length))])).length))).isLessThan(_584_v68.f19);
          })(_242_v68, _155_v0, _385_v167);
          let _nw86 = Array((new BigNumber(13)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw86.length); _i0_14++) {
            _nw86[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _583_v285 = _nw86;
          let _nw87 = Array((new BigNumber(4)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = (new BigNumber(-48)).isEqualTo(_242_v68.f19);
          _nw87[(_dafny.ONE).toNumber()] = _155_v0;
          _nw87[(new BigNumber(2)).toNumber()] = (_385_v167)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_385_v167).length))];
          _nw87[(new BigNumber(3)).toNumber()] = (_385_v167)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_385_v167).length))];
          _583_v285 = _nw87;
          let _588_v286;
          let _nw88 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _588_v286 = _nw88;
          let _589_v287;
          _589_v287 = _module.D5.create_DC17(new _dafny.CodePoint('v'.codePointAt(0)));
          let _590_v288;
          _590_v288 = _dafny.Map.Empty.slice().updateUnsafe(_239_v65.f19,(_589_v287).dtor_cf23);
          let _index106 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_588_v286).length));
          (_588_v286)[_index106] = _590_v288;
          let _index107 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_588_v286).length));
          let _rhs105 = (_590_v288).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(893),_487_v227));
          let _rhs106 = _242_v68;
          let _lhs92 = _588_v286;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_588_v286).length));
          _lhs92[_lhs93] = _rhs105;
          _242_v68 = _rhs106;
        } else {
          _158_v2 = _158_v2;
          (_157_globalState).f3 = _155_v0;
          let _591_v289;
          _591_v289 = _dafny.MultiSet.fromElements(_242_v68);
          let _592_v290;
          _592_v290 = _module.D1.create_DC4(_242_v68);
          let _593_v291;
          _593_v291 = _dafny.Seq.of((_592_v290).dtor_cf9);
          let _594_v292;
          let _nw89 = Array((new BigNumber(14)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _591_v289;
          _nw89[(_dafny.ONE).toNumber()] = (_591_v289).Difference(_dafny.MultiSet.fromElements(_242_v68, _239_v65, _239_v65, _242_v68));
          _nw89[(new BigNumber(2)).toNumber()] = (_591_v289).Union(_dafny.MultiSet.FromArray(_593_v291));
          _nw89[(new BigNumber(3)).toNumber()] = _591_v289;
          _nw89[(new BigNumber(4)).toNumber()] = _591_v289;
          _nw89[(new BigNumber(5)).toNumber()] = _591_v289;
          _nw89[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements(_239_v65)).Intersect(_591_v289);
          _nw89[(new BigNumber(7)).toNumber()] = (_591_v289).Difference(_591_v289);
          _nw89[(new BigNumber(8)).toNumber()] = (_591_v289).Intersect((_dafny.MultiSet.fromElements(_242_v68)).update(_242_v68, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("ihibjhg")).length))));
          _nw89[(new BigNumber(9)).toNumber()] = _591_v289;
          _nw89[(new BigNumber(10)).toNumber()] = (_591_v289).Difference(_591_v289);
          _nw89[(new BigNumber(11)).toNumber()] = (_591_v289).Union(_591_v289);
          _nw89[(new BigNumber(12)).toNumber()] = _591_v289;
          _nw89[(new BigNumber(13)).toNumber()] = _591_v289;
          _594_v292 = _nw89;
          let _index108 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_594_v292).length));
          (_594_v292)[_index108] = _591_v289;
          let _index109 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_594_v292).length));
          (_594_v292)[_index109] = (_591_v289).update(_239_v65, _module.__default.abs(_578_cf2));
          let _595_v293;
          let _init15 = ((_596_v281, _597_v167) => function (_598_i40) {
            return _dafny.Map.Empty.slice().updateUnsafe(_596_v281,(_597_v167)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_597_v167).length))]);
          })(_573_v281, _385_v167);
          let _nw90 = Array((new BigNumber(22)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw90.length); _i0_15++) {
            _nw90[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _595_v293 = _nw90;
          let _599_v294;
          _599_v294 = _dafny.Map.Empty.slice().updateUnsafe(_573_v281,_155_v0);
          let _index110 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_595_v293).length));
          (_595_v293)[_index110] = _599_v294;
          let _index111 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_595_v293).length));
          (_595_v293)[_index111] = (_dafny.Map.Empty.slice().updateUnsafe(_573_v281,_155_v0)).Merge(_599_v294);
          let _600_v295;
          _600_v295 = _dafny.MultiSet.fromElements(new BigNumber((_156_v1).length), _242_v68.f19);
          let _601_v296;
          _601_v296 = _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC20(_386_v168),(_600_v295).Difference(_dafny.MultiSet.fromElements(_579_cf1)));
          let _602_v297;
          _602_v297 = _dafny.Seq.of(_386_v168, _387_v169);
          _601_v296 = (_601_v296).update(_module.D5.create_DC20((_602_v297)[_module.__default.safeIndex(new BigNumber(-900), new BigNumber((_602_v297).length))]), _600_v295);
        }
        let _603_v298;
        _603_v298 = _dafny.MultiSet.fromElements(_239_v65.f19);
        (_157_globalState).f8 = (_dafny.ZERO).minus((_242_v68).fm4(_module.__default.safeModuloInt(new BigNumber((_603_v298).cardinality()), _239_v65.f19), new BigNumber(-138), _155_v0, _577_cf3, _157_globalState));
        (_157_globalState).f3 = (_385_v167)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_385_v167).length))];
      } else if (_source11.is_DC2) {
        let _604___mcc_h31 = (_source11).cf4;
        let _605___mcc_h32 = (_source11).cf5;
        let _606___mcc_h33 = (_source11).cf6;
        let _607___mcc_h34 = (_source11).cf7;
        let _608_cf7 = _607___mcc_h34;
        let _609_cf6 = _606___mcc_h33;
        let _610_cf5 = _605___mcc_h32;
        let _611_cf4 = _604___mcc_h31;
        let _612_v299;
        _612_v299 = _dafny.MultiSet.fromElements(_241_v67);
        let _613_v300;
        _613_v300 = _module.D0.create_DC0(_239_v65.f19);
        let _614_v301;
        let _nw91 = Array((new BigNumber(17)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = _242_v68;
        _nw91[(_dafny.ONE).toNumber()] = _242_v68;
        _nw91[(new BigNumber(2)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(3)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(4)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(5)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(6)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(7)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(8)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(9)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(10)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(11)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(12)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(13)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(14)).toNumber()] = _242_v68;
        _nw91[(new BigNumber(15)).toNumber()] = _239_v65;
        _nw91[(new BigNumber(16)).toNumber()] = _242_v68;
        _614_v301 = _nw91;
        let _615_v302;
        _615_v302 = _dafny.Map.Empty.slice().updateUnsafe(_614_v301,_239_v65.f19);
        let _616_v303;
        _616_v303 = _dafny.Map.Empty.slice().updateUnsafe(_611_cf4,new BigNumber((_241_v67).length));
        let _rhs107 = true;
        let _rhs108 = !(_242_v68.f19).isEqualTo(new BigNumber(((_612_v299).update(_241_v67, _module.__default.abs(_242_v68.f19))).cardinality()));
        let _rhs109 = (_239_v65.f19).isLessThan((_613_v300).dtor_cf0);
        let _rhs110 = _module.__default.safeDivisionInt((_239_v65.f19).minus(new BigNumber((_615_v302).length)), new BigNumber((_616_v303).length));
        _610_cf5 = _rhs107;
        _610_cf5 = _rhs108;
        _155_v0 = _rhs109;
        _609_cf6 = _rhs110;
        if (false) {
          let _617_v304;
          _617_v304 = _dafny.Seq.of(_dafny.Seq.of((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]));
          (_242_v68).f19 = _module.__default.safeDivisionInt(_module.__default.fm0(_158_v2, (_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))], _617_v304, _239_v65.f19, _157_globalState), _158_v2);
          let _618_v305;
          let _nw92 = Array((new BigNumber(20)).toNumber()).fill([]);
          _618_v305 = _nw92;
          let _index112 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_618_v305).length));
          (_618_v305)[_index112] = _385_v167;
          let _index113 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_618_v305).length));
          (_618_v305)[_index113] = _385_v167;
          let _619_v306;
          let _nw93 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
          _619_v306 = _nw93;
          let _620_v307;
          _620_v307 = _dafny.MultiSet.fromElements(_242_v68, _239_v65, _239_v65, _239_v65, _242_v68);
          let _index114 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_619_v306).length));
          (_619_v306)[_index114] = (_620_v307).Difference(_dafny.MultiSet.fromElements(_242_v68, _242_v68));
          let _index115 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_619_v306).length));
          (_619_v306)[_index115] = (_620_v307).Union(_620_v307);
          let _621_v308;
          let _622_v309;
          let _623_v310;
          let _624_v311;
          let _out40;
          let _out41;
          let _out42;
          let _out43;
          let _outcollector10 = _module.__default.m0(_610_cf5, !(((_611_cf4) ? (_608_cf7) : (_608_cf7))), new BigNumber(602), (_159_v3).IsSubsetOf(_159_v3), _157_globalState);
          _out40 = _outcollector10[0];
          _out41 = _outcollector10[1];
          _out42 = _outcollector10[2];
          _out43 = _outcollector10[3];
          _621_v308 = _out40;
          _622_v309 = _out41;
          _623_v310 = _out42;
          _624_v311 = _out43;
          _177_v18 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_622_v309, _dafny.Seq.UnicodeFromString("mhymh")), _module.__default.safeIndex(_242_v68.f19, new BigNumber((_dafny.Seq.Concat(_622_v309, _dafny.Seq.UnicodeFromString("mhymh"))).length)), _487_v227), _module.__default.fm5(_158_v2, _157_globalState));
        } else {
          _608_cf7 = !(_608_cf7);
          let _625_v312;
          _625_v312 = _dafny.Map.Empty.slice().updateUnsafe(_177_v18,_610_cf5);
          _625_v312 = (_625_v312).update(_177_v18, _611_cf4);
          let _626_v313;
          _626_v313 = _dafny.Map.Empty.slice().updateUnsafe(_156_v1,_610_cf5);
          (_157_globalState).f0 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_155_v0,_239_v65.f19)).update((((_626_v313).contains(_156_v1)) ? ((_626_v313).get(_156_v1)) : (_611_cf4)), _239_v65.f19)).length), _609_cf6), new BigNumber((_dafny.Seq.UnicodeFromString("hryki")).length));
          let _index116 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_614_v301).length));
          (_614_v301)[_index116] = _239_v65;
          let _index117 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_614_v301).length));
          (_614_v301)[_index117] = _242_v68;
          let _627_v314;
          let _628_v315;
          let _629_v316;
          let _630_v317;
          let _out44;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector11 = _module.__default.m0((_242_v68.f19).isLessThanOrEqualTo(_239_v65.f19), (_156_v1)[_module.__default.safeIndex(_609_cf6, new BigNumber((_156_v1).length))], _239_v65.f19, false, _157_globalState);
          _out44 = _outcollector11[0];
          _out45 = _outcollector11[1];
          _out46 = _outcollector11[2];
          _out47 = _outcollector11[3];
          _627_v314 = _out44;
          _628_v315 = _out45;
          _629_v316 = _out46;
          _630_v317 = _out47;
        }
        if (_610_cf5) {
          let _631_v318;
          let _nw94 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _631_v318 = _nw94;
          let _632_v319;
          _632_v319 = _dafny.Map.Empty.slice().updateUnsafe(_611_cf4,_608_cf7);
          let _index118 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_631_v318).length));
          (_631_v318)[_index118] = _632_v319;
          let _633_v320;
          _633_v320 = _dafny.Seq.of(_385_v167, _385_v167);
          let _index119 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_631_v318).length));
          let _rhs111 = (_632_v319).Merge((_632_v319).Merge(_632_v319));
          let _rhs112 = (_module.__default.safeModuloInt(_242_v68.f19, new BigNumber((_dafny.Seq.update(_156_v1, _module.__default.safeIndex(_239_v65.f19, new BigNumber((_156_v1).length)), _155_v0)).length))).plus(((_608_cf7) ? (_239_v65.f19) : (_242_v68.f19)));
          let _rhs113 = ((_module.__default.fm2((_573_v281).dtor_cf3, _157_globalState)) ? (new BigNumber((_177_v18).length)) : ((_dafny.ZERO).minus(_242_v68.f19)));
          let _rhs114 = _dafny.areEqual(_dafny.Seq.of(_385_v167), ((_608_cf7) ? (_633_v320) : (_633_v320)));
          let _lhs94 = _631_v318;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_631_v318).length));
          let _lhs96 = _157_globalState;
          let _lhs97 = _157_globalState;
          _lhs94[_lhs95] = _rhs111;
          _lhs96.f2 = _rhs112;
          _lhs97.f0 = _rhs113;
          _155_v0 = _rhs114;
          let _634_v321;
          _634_v321 = _dafny.Set.fromElements(_614_v301);
          _634_v321 = (_634_v321).Union(_634_v321);
          let _635_v322;
          let _636_v323;
          let _637_v324;
          let _638_v325;
          let _out48;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector12 = _module.__default.m0(_611_cf4, _610_cf5, _239_v65.f19, !(_608_cf7), _157_globalState);
          _out48 = _outcollector12[0];
          _out49 = _outcollector12[1];
          _out50 = _outcollector12[2];
          _out51 = _outcollector12[3];
          _635_v322 = _out48;
          _636_v323 = _out49;
          _637_v324 = _out50;
          _638_v325 = _out51;
          let _639_v326;
          _639_v326 = _module.D0.create_DC2(_611_cf4, _608_cf7, _239_v65.f19, _610_cf5);
          let _640_v327;
          let _641_v328;
          let _642_v329;
          let _643_v330;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector13 = _module.__default.m0((_639_v326).dtor_cf5, _module.__default.fm2(true, _157_globalState), _609_cf6, _610_cf5, _157_globalState);
          _out52 = _outcollector13[0];
          _out53 = _outcollector13[1];
          _out54 = _outcollector13[2];
          _out55 = _outcollector13[3];
          _640_v327 = _out52;
          _641_v328 = _out53;
          _642_v329 = _out54;
          _643_v330 = _out55;
          _642_v329 = _637_v324;
        } else {
          let _644_v331;
          let _nw95 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
          _644_v331 = _nw95;
          let _645_v332;
          _645_v332 = _dafny.Map.Empty.slice().updateUnsafe(_177_v18,_611_cf4);
          let _index120 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_644_v331).length));
          (_644_v331)[_index120] = _645_v332;
          let _index121 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_644_v331).length));
          (_644_v331)[_index121] = _645_v332;
          _487_v227 = _487_v227;
          let _646_v333;
          _646_v333 = _module.D2.create_DC9(_242_v68.f19, _242_v68, _239_v65);
          let _647_v334;
          _647_v334 = _module.D1.create_DC3(_module.__default.fm5(_239_v65.f19, _157_globalState));
          let _648_v335;
          _648_v335 = _dafny.Set.fromElements(_647_v334);
          let _649_v336;
          _649_v336 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let18_0) {
            return function (_650_dt__update__tmp_h6) {
              return function (_pat_let19_0) {
                return function (_651_dt__update_hcf13_h0) {
                  return _module.D2.create_DC9(_651_dt__update_hcf13_h0, (_650_dt__update__tmp_h6).dtor_cf14, (_650_dt__update__tmp_h6).dtor_cf15);
                }(_pat_let19_0);
              }(new BigNumber(961));
            }(_pat_let18_0);
          }(_646_v333),(_648_v335).Difference(_dafny.Set.fromElements(_647_v334, _module.D1.create_DC3(_177_v18), _647_v334)));
          _649_v336 = (_649_v336).update(_module.D2.create_DC9(_609_cf6, _239_v65, _239_v65), (_648_v335).Difference(_648_v335));
          _616_v303 = (_616_v303).update(_611_cf4, _242_v68.f19);
          let _652_v337;
          let _nw96 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _652_v337 = _nw96;
          let _653_v338;
          _653_v338 = _dafny.Seq.of(_177_v18, _177_v18);
          let _index122 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_652_v337).length));
          (_652_v337)[_index122] = new BigNumber((_653_v338).length);
          let _index123 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_652_v337).length));
          (_652_v337)[_index123] = _239_v65.f19;
        }
        let _654_v339;
        _654_v339 = _dafny.Seq.of(_242_v68, _239_v65, _242_v68, _242_v68, _239_v65);
        let _655_v340;
        _655_v340 = _module.D4.create_DC14(_654_v339);
        let _656_v341;
        _656_v341 = _dafny.Set.fromElements(_238_v64, _238_v64);
        let _rhs115 = !(_656_v341).equals(function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of (_dafny.Seq.of(_238_v64)).Elements) {
            let _657_v342 = _compr_13;
            if (_dafny.Seq.contains(_dafny.Seq.of(_238_v64), _657_v342)) {
              _coll13.add(_657_v342);
            }
          }
          return _coll13;
        }());
        let _rhs116 = _655_v340;
        let _rhs117 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), ((_658_v227) => function (_659_i41) {
          return _658_v227;
        })(_487_v227));
        let _rhs118 = (_609_cf6).isLessThanOrEqualTo(_242_v68.f19);
        let _lhs98 = _157_globalState;
        let _lhs99 = _157_globalState;
        _lhs98.f3 = _rhs115;
        _655_v340 = _rhs116;
        _177_v18 = _rhs117;
        _lhs99.f3 = _rhs118;
      } else {
        let _660___mcc_h35 = (_source11).cf0;
        let _661_cf0 = _660___mcc_h35;
        let _662_v343;
        let _nw97 = new _module.C0();
        _nw97.__ctor(_661_cf0);
        _662_v343 = _nw97;
        let _663_v344;
        let _nw98 = Array((new BigNumber(4)).toNumber());
        _663_v344 = _nw98;
        let _index124 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_663_v344).length));
        (_663_v344)[_index124] = _242_v68;
        let _index125 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_663_v344).length));
        (_663_v344)[_index125] = _239_v65;
        let _664_v345;
        _664_v345 = _dafny.MultiSet.fromElements(new BigNumber(-949));
        if (((_662_v343.f19).isEqualTo(_242_v68.f19)) === ((_664_v345).IsProperSubsetOf(_664_v345))) {
          let _665_v346;
          _665_v346 = _dafny.Set.fromElements(_662_v343.f19, _662_v343.f19, _662_v343.f19, _239_v65.f19);
          let _666_v349;
          _666_v349 = _dafny.MultiSet.fromElements(_665_v346, (_665_v346).Union(_665_v346), function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of _dafny.IntegerRange(new BigNumber(964), new BigNumber(-496))) {
              let _667_v347 = _compr_14;
              if (((new BigNumber(964)).isLessThanOrEqualTo(_667_v347)) && ((_667_v347).isLessThan(new BigNumber(-496)))) {
                _coll14.add((_667_v347).multipliedBy(_662_v343.f19));
              }
            }
            return _coll14;
          }(), _module.__default.fm11(_242_v68.f19, _157_globalState), function () {
            let _coll15 = new _dafny.Set();
            for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-993), new BigNumber(-987))) {
              let _668_v348 = _compr_15;
              if (((new BigNumber(-993)).isLessThanOrEqualTo(_668_v348)) && ((_668_v348).isLessThan(new BigNumber(-987)))) {
                _coll15.add((_668_v348).multipliedBy(new BigNumber((_177_v18).length)));
              }
            }
            return _coll15;
          }());
          let _669_v350;
          _669_v350 = _dafny.Seq.of(_666_v349);
          _666_v349 = (_666_v349).Difference(((_669_v350)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_669_v350).length))]).Difference(_666_v349));
          let _670_v351;
          let _nw99 = Array((new BigNumber(11)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("fnbj");
          _nw99[(_dafny.ONE).toNumber()] = _177_v18;
          _nw99[(new BigNumber(2)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(3)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(4)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("omisqr");
          _nw99[(new BigNumber(6)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(7)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(8)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(9)).toNumber()] = _177_v18;
          _nw99[(new BigNumber(10)).toNumber()] = _177_v18;
          _670_v351 = _nw99;
          let _671_v352;
          _671_v352 = _module.D3.create_DC12(_670_v351, _155_v0);
          let _672_v353;
          _672_v353 = _module.D1.create_DC5(false);
          let _673_v354;
          let _674_v355;
          let _675_v356;
          let _676_v357;
          let _out56;
          let _out57;
          let _out58;
          let _out59;
          let _outcollector14 = _module.__default.m0((_156_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_156_v1).length))], _155_v0, (_239_v65).fm4(_662_v343.f19, _239_v65.f19, (_671_v352).dtor_cf19, (_672_v353).dtor_cf10, _157_globalState), (_662_v343.f19).isLessThan((_242_v68).fm4(_158_v2, (_dafny.ZERO).minus((_dafny.ZERO).minus(_242_v68.f19)), _155_v0, _155_v0, _157_globalState)), _157_globalState);
          _out56 = _outcollector14[0];
          _out57 = _outcollector14[1];
          _out58 = _outcollector14[2];
          _out59 = _outcollector14[3];
          _673_v354 = _out56;
          _674_v355 = _out57;
          _675_v356 = _out58;
          _676_v357 = _out59;
          let _677_v358;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_242_v68.f19);
          _677_v358 = _nw100;
          let _678_v359;
          let _nw101 = new _module.C0();
          _nw101.__ctor(_239_v65.f19);
          _678_v359 = _nw101;
          let _679_v360;
          let _nw102 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _679_v360 = _nw102;
          _679_v360 = _679_v360;
        } else {
          (_157_globalState).f3 = _155_v0;
          _662_v343 = _239_v65;
          let _index126 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_385_v167).length));
          (_385_v167)[_index126] = _155_v0;
          let _index127 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_385_v167).length));
          (_385_v167)[_index127] = (_156_v1)[_module.__default.safeIndex(((_155_v0) ? (_242_v68.f19) : (_242_v68.f19)), new BigNumber((_156_v1).length))];
          _664_v345 = _664_v345;
          let _680_v361;
          _680_v361 = _dafny.Seq.of(_177_v18, _dafny.Seq.UnicodeFromString("c"), _177_v18);
          let _681_v362;
          _681_v362 = _module.D1.create_DC3((_680_v361)[_module.__default.safeIndex((_dafny.ZERO).minus(_239_v65.f19), new BigNumber((_680_v361).length))]);
          let _rhs119 = _681_v362;
          let _rhs120 = _module.__default.fm5(_158_v2, _157_globalState);
          _681_v362 = _rhs119;
          _177_v18 = _rhs120;
        }
        (_157_globalState).f3 = !(!(!(_155_v0)) || (_155_v0));
      }
      let _682_v363;
      let _nw103 = Array((new BigNumber(9)).toNumber());
      _nw103[(_dafny.ZERO).toNumber()] = _573_v281;
      _nw103[(_dafny.ONE).toNumber()] = _module.__default.fm12(_239_v65.f19, _242_v68.f19, _242_v68.f19, _155_v0, _157_globalState);
      _nw103[(new BigNumber(2)).toNumber()] = _573_v281;
      _nw103[(new BigNumber(3)).toNumber()] = _module.D0.create_DC1(_242_v68.f19, _158_v2, _155_v0);
      _nw103[(new BigNumber(4)).toNumber()] = _module.D0.create_DC1((_239_v65).fm3(_157_globalState), _239_v65.f19, _155_v0);
      _nw103[(new BigNumber(5)).toNumber()] = _573_v281;
      _nw103[(new BigNumber(6)).toNumber()] = _573_v281;
      _nw103[(new BigNumber(7)).toNumber()] = _573_v281;
      _nw103[(new BigNumber(8)).toNumber()] = _573_v281;
      _682_v363 = _nw103;
      let _index128 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_682_v363).length));
      (_682_v363)[_index128] = _573_v281;
      let _index129 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_682_v363).length));
      (_682_v363)[_index129] = _573_v281;
      let _683_v365;
      _683_v365 = _dafny.Map.Empty.slice().updateUnsafe(_385_v167,function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-50), new BigNumber(413))) {
          let _684_v364 = _compr_16;
          if (((new BigNumber(-50)).isLessThanOrEqualTo(_684_v364)) && ((_684_v364).isLessThan(new BigNumber(413)))) {
            _coll16.push([_module.__default.safeDivisionInt(_684_v364, new BigNumber((_156_v1).length)),_242_v68.f19]);
          }
        }
        return _coll16;
      }());
      _155_v0 = (_683_v365).contains(_385_v167);
      let _685_v366;
      _685_v366 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_177_v18).length),_159_v3);
      let _686_i42;
      _686_i42 = _dafny.ZERO;
      L1: {
        while ((_685_v366).contains(_242_v68.f19)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_686_i42)) {
              break L1;
            }
            _686_i42 = (_686_i42).plus(_dafny.ONE);
            let _687_v367;
            _687_v367 = _dafny.Map.Empty.slice().updateUnsafe(_242_v68.f19,_155_v0);
            let _688_v368;
            _688_v368 = _module.D3.create_DC11(_687_v367);
            let _689_v369;
            _689_v369 = _module.D3.create_DC13(_688_v368);
            let _690_v370;
            _690_v370 = _module.D3.create_DC13((_689_v369).dtor_cf20);
            let _source12 = _690_v370;
            if (_source12.is_DC12) {
              let _691___mcc_h36 = (_source12).cf18;
              let _692___mcc_h37 = (_source12).cf19;
              let _693_cf19 = _692___mcc_h37;
              let _694_cf18 = _691___mcc_h36;
              let _index130 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_385_v167).length));
              (_385_v167)[_index130] = (_693_cf19) || (_693_cf19);
              let _695_v371;
              _695_v371 = _dafny.Set.fromElements((_dafny.ZERO).minus(_242_v68.f19));
              let _index131 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_385_v167).length));
              (_385_v167)[_index131] = (((new BigNumber((_695_v371).length)).isLessThanOrEqualTo(_242_v68.f19)) ? (true) : (_155_v0));
              (_157_globalState).f8 = new BigNumber((_177_v18).length);
              let _696_v372;
              _696_v372 = _dafny.Map.Empty.slice().updateUnsafe(true,_388_v170);
              (_157_globalState).f2 = new BigNumber(((_696_v372).update(!(_242_v68.f19).isEqualTo(_158_v2), _388_v170)).length);
              let _index132 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_385_v167).length));
              (_385_v167)[_index132] = (_242_v68.f19).isLessThan(_module.__default.safeModuloInt(new BigNumber((_177_v18).length), _239_v65.f19));
            } else if (_source12.is_DC11) {
              let _697___mcc_h38 = (_source12).cf17;
              let _698_cf17 = _697___mcc_h38;
              _158_v2 = new BigNumber(554);
              _158_v2 = _239_v65.f19;
              let _699_v373;
              _699_v373 = _module.D5.create_DC19(_155_v0);
              let _pat_let_tv24 = _155_v0;
              let _700_v374;
              _700_v374 = _dafny.Seq.of(_699_v373, function (_pat_let20_0) {
                return function (_701_dt__update__tmp_h7) {
                  return function (_pat_let21_0) {
                    return function (_702_dt__update_hcf28_h0) {
                      return _module.D5.create_DC19(_702_dt__update_hcf28_h0);
                    }(_pat_let21_0);
                  }(_pat_let_tv24);
                }(_pat_let20_0);
              }(_699_v373));
              let _703_v375;
              let _nw104 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
              _703_v375 = _nw104;
              let _index133 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_703_v375).length));
              (_703_v375)[_index133] = _239_v65.f19;
              let _index134 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_703_v375).length));
              let _rhs121 = _dafny.Seq.of(_699_v373, _module.__default.fm13((_dafny.ZERO).minus(_242_v68.f19), _157_globalState), _699_v373);
              let _rhs122 = (_242_v68.f19).multipliedBy(_158_v2);
              let _rhs123 = _dafny.Seq.Concat(_dafny.Seq.Concat(_238_v64, _238_v64), _dafny.Seq.Concat(_238_v64, _238_v64));
              let _rhs124 = _155_v0;
              let _lhs100 = _703_v375;
              let _lhs101 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_703_v375).length));
              _700_v374 = _rhs121;
              _lhs100[_lhs101] = _rhs122;
              _238_v64 = _rhs123;
              _155_v0 = _rhs124;
              let _704_v376;
              _704_v376 = _dafny.MultiSet.fromElements(_242_v68.f19, _239_v65.f19, _242_v68.f19, _242_v68.f19);
              (_157_globalState).f3 = (_704_v376).contains((_242_v68.f19).multipliedBy((_239_v65).fm4(_239_v65.f19, _239_v65.f19, _155_v0, _155_v0, _157_globalState)));
            } else {
              let _705___mcc_h39 = (_source12).cf20;
              let _706_cf20 = _705___mcc_h39;
              let _707_v377;
              let _708_v378;
              let _709_v379;
              let _710_v380;
              let _out60;
              let _out61;
              let _out62;
              let _out63;
              let _outcollector15 = _module.__default.m0((_155_v0) === (_155_v0), _155_v0, _242_v68.f19, _155_v0, _157_globalState);
              _out60 = _outcollector15[0];
              _out61 = _outcollector15[1];
              _out62 = _outcollector15[2];
              _out63 = _outcollector15[3];
              _707_v377 = _out60;
              _708_v378 = _out61;
              _709_v379 = _out62;
              _710_v380 = _out63;
              _155_v0 = (_155_v0) === (_155_v0);
              let _711_v381;
              _711_v381 = _dafny.Set.fromElements(_709_v379, new BigNumber(-70));
              let _712_v382;
              let _nw105 = Array((new BigNumber(26)).toNumber());
              _nw105[(_dafny.ZERO).toNumber()] = _238_v64;
              _nw105[(_dafny.ONE).toNumber()] = ((_155_v0) ? (_238_v64) : (_238_v64));
              _nw105[(new BigNumber(2)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(3)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(4)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_238_v64, _238_v64);
              _nw105[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_238_v64, _238_v64);
              _nw105[(new BigNumber(7)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(8)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(9)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(10)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(11)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(12)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), ((_713_v68) => function (_714_i43) {
                return _713_v68.f19;
              })(_242_v68)), _module.__default.safeIndex((_dafny.ZERO).minus((_242_v68).fm4(_242_v68.f19, new BigNumber(248), _155_v0, true, _157_globalState)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), ((_715_v68) => function (_716_i43) {
                return _715_v68.f19;
              })(_242_v68))).length)), _158_v2);
              _nw105[(new BigNumber(14)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_238_v64, _238_v64);
              _nw105[(new BigNumber(16)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(809)), ((_717_v68) => function (_718_i44) {
                return (_dafny.ZERO).minus(_717_v68.f19);
              })(_242_v68)), _module.__default.safeIndex(_158_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(809)), ((_719_v68) => function (_720_i44) {
                return (_dafny.ZERO).minus(_719_v68.f19);
              })(_242_v68))).length)), _239_v65.f19);
              _nw105[(new BigNumber(18)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), ((_721_v68) => function (_722_i45) {
                return _721_v68.f19;
              })(_242_v68));
              _nw105[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_242_v68.f19, new BigNumber((_711_v381).length)), _module.__default.safeIndex(new BigNumber(136), new BigNumber((_dafny.Seq.of(_242_v68.f19, new BigNumber((_711_v381).length))).length)), _709_v379);
              _nw105[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_238_v64, _dafny.Seq.of(_707_v377));
              _nw105[(new BigNumber(22)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(23)).toNumber()] = _238_v64;
              _nw105[(new BigNumber(24)).toNumber()] = _module.__default.fm14(_155_v0, _157_globalState);
              _nw105[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(750), _242_v68.f19), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_155_v0)).length), _239_v65.f19));
              _712_v382 = _nw105;
              let _index135 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_712_v382).length));
              (_712_v382)[_index135] = _dafny.Seq.Concat(_238_v64, _238_v64);
              let _723_v383;
              let _init16 = ((_724_v2) => function (_725_i46) {
                return (_725_i46).multipliedBy(_724_v2);
              })(_158_v2);
              let _nw106 = Array((new BigNumber(26)).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw106.length); _i0_16++) {
                _nw106[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _723_v383 = _nw106;
              let _index136 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_723_v383).length));
              (_723_v383)[_index136] = _239_v65.f19;
              let _726_v384;
              _726_v384 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_dafny.Seq.of(_239_v65.f19));
              let _index137 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_712_v382).length));
              let _index138 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_723_v383).length));
              let _rhs125 = _dafny.Seq.contains(_156_v1, _155_v0);
              let _rhs126 = _dafny.Seq.Concat(_dafny.Seq.Concat(_238_v64, _dafny.Seq.of(_158_v2)), _238_v64);
              let _rhs127 = _325_v126;
              let _rhs128 = (_242_v68).fm4(new BigNumber((_726_v384).length), _239_v65.f19, ((_239_v65).fm3(_157_globalState)).isEqualTo(_242_v68.f19), _155_v0, _157_globalState);
              let _lhs102 = _157_globalState;
              let _lhs103 = _712_v382;
              let _lhs104 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_712_v382).length));
              let _lhs105 = _723_v383;
              let _lhs106 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_723_v383).length));
              _lhs102.f3 = _rhs125;
              _lhs103[_lhs104] = _rhs126;
              _325_v126 = _rhs127;
              _lhs105[_lhs106] = _rhs128;
              (_157_globalState).f3 = _module.__default.fm2(!(false), _157_globalState);
            }
            if (_155_v0) {
              let _rhs129 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-882)), ((_727_v227) => function (_728_i47) {
                return _727_v227;
              })(_487_v227)), _module.__default.safeIndex(_239_v65.f19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-882)), ((_729_v227) => function (_730_i47) {
                return _729_v227;
              })(_487_v227))).length)), (((_156_v1)[_module.__default.safeIndex(_242_v68.f19, new BigNumber((_156_v1).length))]) ? ((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]) : (_487_v227)));
              let _rhs130 = _487_v227;
              let _rhs131 = ((((_241_v67).contains(_239_v65.f19)) ? ((_241_v67).get(_239_v65.f19)) : (_239_v65.f19))).isLessThanOrEqualTo(_242_v68.f19);
              let _rhs132 = _242_v68.f19;
              let _rhs133 = _239_v65;
              let _lhs107 = _157_globalState;
              _177_v18 = _rhs129;
              _487_v227 = _rhs130;
              _155_v0 = _rhs131;
              _lhs107.f2 = _rhs132;
              _242_v68 = _rhs133;
              let _nw107 = new _module.C0();
              _nw107.__ctor(_239_v65.f19);
              _239_v65 = _nw107;
              let _731_v385;
              _731_v385 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_487_v227);
              _731_v385 = (_731_v385).update(_155_v0, new _dafny.CodePoint('m'.codePointAt(0)));
              let _732_v386;
              let _init17 = ((_733_v65) => function (_734_i48) {
                return (_734_i48).minus(_733_v65.f19);
              })(_239_v65);
              let _nw108 = Array((new BigNumber(10)).toNumber());
              for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw108.length); _i0_17++) {
                _nw108[_i0_17] = _init17(new BigNumber(_i0_17));
              }
              _732_v386 = _nw108;
              let _index139 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_732_v386).length));
              (_732_v386)[_index139] = _242_v68.f19;
              let _index140 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_732_v386).length));
              let _rhs134 = _238_v64;
              let _rhs135 = _239_v65;
              let _rhs136 = _242_v68.f19;
              let _rhs137 = new BigNumber(742);
              let _rhs138 = (((_731_v385).contains(_155_v0)) ? ((_731_v385).get(_155_v0)) : (new _dafny.CodePoint('e'.codePointAt(0))));
              let _lhs108 = _157_globalState;
              let _lhs109 = _732_v386;
              let _lhs110 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_732_v386).length));
              _238_v64 = _rhs134;
              _242_v68 = _rhs135;
              _lhs108.f8 = _rhs136;
              _lhs109[_lhs110] = _rhs137;
              _487_v227 = _rhs138;
              (_239_v65).f19 = _239_v65.f19;
            } else {
              let _index141 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_385_v167).length));
              (_385_v167)[_index141] = _dafny.Seq.contains(_177_v18, _487_v227);
              let _735_v387;
              _735_v387 = _dafny.Set.fromElements((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))], new _dafny.CodePoint('p'.codePointAt(0)));
              let _index142 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_385_v167).length));
              (_385_v167)[_index142] = (_155_v0) === ((_735_v387).IsSubsetOf(_735_v387));
              let _736_v388;
              let _nw109 = new _module.C0();
              _nw109.__ctor(_module.__default.safeDivisionInt(new BigNumber((_177_v18).length), _239_v65.f19));
              _736_v388 = _nw109;
              let _737_v389;
              _737_v389 = _dafny.Set.fromElements(new BigNumber(512), new BigNumber(-517), _242_v68.f19);
              let _rhs139 = ((!(_155_v0)) ? (_module.__default.fm11(_239_v65.f19, _157_globalState)) : (function () {
                let _coll17 = new _dafny.Set();
                for (const _compr_17 of _dafny.IntegerRange(new BigNumber(779), new BigNumber(398))) {
                  let _738_v390 = _compr_17;
                  if (((new BigNumber(779)).isLessThanOrEqualTo(_738_v390)) && ((_738_v390).isLessThan(new BigNumber(398)))) {
                    _coll17.add(_module.__default.safeModuloInt(_738_v390, _239_v65.f19));
                  }
                }
                return _coll17;
              }()));
              let _rhs140 = (_dafny.ZERO).minus((_239_v65.f19).minus(new BigNumber(168)));
              let _rhs141 = ((_dafny.ZERO).minus(_242_v68.f19)).minus((_239_v65).fm3(_157_globalState));
              let _rhs142 = _155_v0;
              let _lhs111 = _242_v68;
              let _lhs112 = _157_globalState;
              let _lhs113 = _157_globalState;
              _737_v389 = _rhs139;
              _lhs111.f19 = _rhs140;
              _lhs112.f12 = _rhs141;
              _lhs113.f3 = _rhs142;
              let _739_v391;
              let _nw110 = new _module.C0();
              _nw110.__ctor((((_385_v167)[_module.__default.safeIndex(new BigNumber(383), new BigNumber((_385_v167).length))]) ? (new BigNumber(264)) : (_242_v68.f19)));
              _739_v391 = _nw110;
              (_736_v388).f19 = _158_v2;
            }
            let _740_v392;
            _740_v392 = _module.D5.create_DC17(_487_v227);
            let _741_v393;
            _741_v393 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(458)), ((_742_v227) => function (_743_i49) {
              return _742_v227;
            })(_487_v227)), _177_v18);
            let _744_v394;
            _744_v394 = _module.D6.create_DC21(_741_v393);
            let _pat_let_tv25 = _741_v393;
            _687_v367 = (_687_v367).update((_239_v65.f19).plus(_module.__default.fm0(_242_v68.f19, (_740_v392).dtor_cf23, (function (_pat_let22_0) {
              return function (_745_dt__update__tmp_h8) {
                return function (_pat_let23_0) {
                  return function (_746_dt__update_hcf30_h0) {
                    return _module.D6.create_DC21(_746_dt__update_hcf30_h0);
                  }(_pat_let23_0);
                }(_pat_let_tv25);
              }(_pat_let22_0);
            }(_744_v394)).dtor_cf30, (_dafny.ZERO).minus((_239_v65).fm4(_239_v65.f19, new BigNumber((_177_v18).length), _155_v0, (((_687_v367).contains(_239_v65.f19)) ? ((_687_v367).get(_239_v65.f19)) : (_155_v0)), _157_globalState)), _157_globalState)), (_239_v65.f19).isLessThanOrEqualTo(_158_v2));
            if (_155_v0) {
              let _747_v395;
              let _748_v396;
              let _749_v397;
              let _750_v398;
              let _out64;
              let _out65;
              let _out66;
              let _out67;
              let _outcollector16 = _module.__default.m0(false, true, _239_v65.f19, _155_v0, _157_globalState);
              _out64 = _outcollector16[0];
              _out65 = _outcollector16[1];
              _out66 = _outcollector16[2];
              _out67 = _outcollector16[3];
              _747_v395 = _out64;
              _748_v396 = _out65;
              _749_v397 = _out66;
              _750_v398 = _out67;
              let _751_v399;
              _751_v399 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(_747_v395)), _242_v68.f19);
              let _752_v400;
              _752_v400 = _dafny.Map.Empty.slice().updateUnsafe(true,_751_v399);
              let _753_v401;
              let _nw111 = Array((new BigNumber(13)).toNumber());
              _nw111[(_dafny.ZERO).toNumber()] = (_751_v399).Intersect(_751_v399);
              _nw111[(_dafny.ONE).toNumber()] = (((_752_v400).contains(_155_v0)) ? ((_752_v400).get(_155_v0)) : (_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(new BigNumber(-279)), _module.__default.safeIndex(_239_v65.f19, new BigNumber((_dafny.Seq.of(new BigNumber(-279))).length)), _239_v65.f19))));
              _nw111[(new BigNumber(2)).toNumber()] = _751_v399;
              _nw111[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_747_v395));
              _nw111[(new BigNumber(4)).toNumber()] = ((_751_v399).update(_239_v65.f19, _module.__default.abs(_239_v65.f19))).Union(_module.__default.fm15(_157_globalState));
              _nw111[(new BigNumber(5)).toNumber()] = _751_v399;
              _nw111[(new BigNumber(6)).toNumber()] = _751_v399;
              _nw111[(new BigNumber(7)).toNumber()] = _751_v399;
              _nw111[(new BigNumber(8)).toNumber()] = ((((_752_v400).contains(_155_v0)) ? ((_752_v400).get(_155_v0)) : (_751_v399))).Difference(_751_v399);
              _nw111[(new BigNumber(9)).toNumber()] = (_751_v399).Difference(_dafny.MultiSet.fromElements(_242_v68.f19, _239_v65.f19, _239_v65.f19, _158_v2, (_dafny.ZERO).minus(_158_v2)));
              _nw111[(new BigNumber(10)).toNumber()] = _module.__default.fm15(_157_globalState);
              _nw111[(new BigNumber(11)).toNumber()] = _751_v399;
              _nw111[(new BigNumber(12)).toNumber()] = _751_v399;
              _753_v401 = _nw111;
              _753_v401 = _753_v401;
              let _754_v402;
              _754_v402 = _dafny.Map.Empty.slice().updateUnsafe(_155_v0,_242_v68.f19);
              let _755_v403;
              _755_v403 = _dafny.Set.fromElements((_754_v402).Merge(_754_v402));
              let _756_v404;
              let _nw112 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _756_v404 = _nw112;
              let _index143 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_756_v404).length));
              (_756_v404)[_index143] = ((!(_155_v0)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), ((_757_v228) => function (_758_i50) {
                return (_757_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_757_v228).length))];
              })(_488_v228))) : (_dafny.Seq.UnicodeFromString("mamgipjj")));
              let _759_v405;
              let _nw113 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
              _759_v405 = _nw113;
              let _index144 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_759_v405).length));
              (_759_v405)[_index144] = new BigNumber(-902);
              let _760_v406;
              _760_v406 = _dafny.Seq.of(_239_v65, _242_v68, _239_v65);
              let _index145 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_756_v404).length));
              let _index146 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_759_v405).length));
              let _rhs143 = _dafny.Seq.Concat(_748_v396, _module.__default.fm5(new BigNumber((_760_v406).length), _157_globalState));
              let _rhs144 = (_755_v403).Intersect(_755_v403);
              let _rhs145 = _dafny.Seq.update(_dafny.Seq.Concat(_177_v18, _748_v396), _module.__default.safeIndex((_239_v65.f19).plus(_749_v397), new BigNumber((_dafny.Seq.Concat(_177_v18, _748_v396)).length)), (_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))]);
              let _rhs146 = _158_v2;
              let _lhs114 = _756_v404;
              let _lhs115 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_756_v404).length));
              let _lhs116 = _759_v405;
              let _lhs117 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_759_v405).length));
              _748_v396 = _rhs143;
              _755_v403 = _rhs144;
              _lhs114[_lhs115] = _rhs145;
              _lhs116[_lhs117] = _rhs146;
              let _761_v407;
              _761_v407 = _dafny.Map.Empty.slice().updateUnsafe((_159_v3).Difference(_750_v398),_242_v68.f19);
              let _762_v408;
              let _nw114 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
              _762_v408 = _nw114;
              let _rhs147 = (new BigNumber(-575)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), ((_763_v227) => function (_764_i51) {
                return _763_v227;
              })(_487_v227))).length));
              let _rhs148 = _761_v407;
              let _rhs149 = _155_v0;
              let _rhs150 = _239_v65;
              let _rhs151 = _762_v408;
              let _lhs118 = _157_globalState;
              let _lhs119 = _157_globalState;
              _lhs118.f12 = _rhs147;
              _761_v407 = _rhs148;
              _lhs119.f3 = _rhs149;
              _242_v68 = _rhs150;
              _762_v408 = _rhs151;
              (_157_globalState).f3 = (_242_v68.f19).isLessThan(_239_v65.f19);
            } else {
              (_239_v65).f19 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_177_v18, _177_v18)).length), _242_v68.f19);
              _241_v67 = _module.__default.fm10(_157_globalState);
              _155_v0 = _155_v0;
              let _765_v410;
              let _init18 = ((_766_v67) => function (_767_i52) {
                return (_dafny.Set.fromElements(_766_v67)).Difference(function () {
                  let _coll18 = new _dafny.Set();
                  for (const _compr_18 of (_dafny.Set.fromElements(_766_v67)).Elements) {
                    let _768_v409 = _compr_18;
                    if ((_dafny.Set.fromElements(_766_v67)).contains(_768_v409)) {
                      _coll18.add(_768_v409);
                    }
                  }
                  return _coll18;
                }());
              })(_241_v67);
              let _nw115 = Array((new BigNumber(9)).toNumber());
              for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw115.length); _i0_18++) {
                _nw115[_i0_18] = _init18(new BigNumber(_i0_18));
              }
              _765_v410 = _nw115;
              let _769_v411;
              _769_v411 = _module.D1.create_DC5(_module.__default.fm2(_155_v0, _157_globalState));
              let _770_v412;
              _770_v412 = _dafny.Map.Empty.slice().updateUnsafe(_241_v67,_769_v411);
              let _index147 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_765_v410).length));
              (_765_v410)[_index147] = ((function () {
                let _coll19 = new _dafny.Set();
                for (const _compr_19 of (_770_v412).Keys.Elements) {
                  let _771_v413 = _compr_19;
                  if ((_770_v412).contains(_771_v413)) {
                    _coll19.add(_771_v413);
                  }
                }
                return _coll19;
              }())).Difference(_module.__default.fm16(_157_globalState));
              let _772_v414;
              _772_v414 = _dafny.Set.fromElements(_241_v67);
              let _index148 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_765_v410).length));
              (_765_v410)[_index148] = _772_v414;
              let _773_v415;
              let _nw116 = Array((new BigNumber(19)).toNumber());
              _773_v415 = _nw116;
              let _774_v418;
              let _nw117 = Array((new BigNumber(21)).toNumber());
              _nw117[(_dafny.ZERO).toNumber()] = _241_v67;
              _nw117[(_dafny.ONE).toNumber()] = _241_v67;
              _nw117[(new BigNumber(2)).toNumber()] = function () {
                let _coll20 = new _dafny.Map();
                for (const _compr_20 of _dafny.IntegerRange(new BigNumber(203), new BigNumber(-284))) {
                  let _775_v416 = _compr_20;
                  if (((new BigNumber(203)).isLessThanOrEqualTo(_775_v416)) && ((_775_v416).isLessThan(new BigNumber(-284)))) {
                    _coll20.push([(_775_v416).minus(_158_v2),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_488_v228)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_488_v228).length))],!(_155_v0))).length)]);
                  }
                }
                return _coll20;
              }();
              _nw117[(new BigNumber(3)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(4)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(5)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(6)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(7)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(8)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(9)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(10)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(11)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_242_v68.f19,new BigNumber(908));
              _nw117[(new BigNumber(13)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(14)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(15)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(16)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(17)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(18)).toNumber()] = _241_v67;
              _nw117[(new BigNumber(19)).toNumber()] = (_241_v67).update(_242_v68.f19, (_dafny.ZERO).minus(_239_v65.f19));
              _nw117[(new BigNumber(20)).toNumber()] = function () {
                let _coll21 = new _dafny.Map();
                for (const _compr_21 of (_241_v67).Keys.Elements) {
                  let _776_v417 = _compr_21;
                  if ((_241_v67).contains(_776_v417)) {
                    _coll21.push([(_776_v417).minus(_242_v68.f19),new BigNumber((_159_v3).cardinality())]);
                  }
                }
                return _coll21;
              }();
              _774_v418 = _nw117;
              let _777_v419;
              _777_v419 = _dafny.Map.Empty.slice().updateUnsafe(_774_v418,_773_v415);
              _773_v415 = (((_777_v419).contains(_774_v418)) ? ((_777_v419).get(_774_v418)) : (_773_v415));
            }
          }
        }
      }
      let _index149 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_385_v167).length));
      (_385_v167)[_index149] = _155_v0;
      let _778_v420;
      let _nw118 = Array((new BigNumber(29)).toNumber()).fill(_module.D3.Default());
      _778_v420 = _nw118;
      let _index150 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_385_v167).length));
      let _rhs152 = !(new BigNumber(519)).isEqualTo(_module.__default.safeModuloInt(_239_v65.f19, new BigNumber((_dafny.Seq.of(!(_155_v0), _155_v0)).length)));
      let _rhs153 = ((_155_v0) ? (_778_v420) : (_778_v420));
      let _lhs120 = _385_v167;
      let _lhs121 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_385_v167).length));
      _lhs120[_lhs121] = _rhs152;
      _778_v420 = _rhs153;
      process.stdout.write(_dafny.toString(_155_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_156_v1, _dafny.Seq.of(false, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_157_globalState).f10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_157_globalState).f18, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_158_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v3).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_177_v18, _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v63).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),false).updateUnsafe(_dafny.Set.fromElements(true),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_238_v64, _dafny.Seq.of(new BigNumber(270)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_239_v65.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_240_v66).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_241_v67).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(270)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_242_v68.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_243_v69).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_244_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_325_v126).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_385_v167)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_385_v167)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_386_v168).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_386_v168).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_386_v168).dtor_cf26)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_386_v168).dtor_cf26)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_386_v168).dtor_cf27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_387_v169).dtor_cf29).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_387_v169).dtor_cf29).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_387_v169).dtor_cf29).dtor_cf26)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_387_v169).dtor_cf29).dtor_cf26)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_387_v169).dtor_cf29).dtor_cf27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_388_v170).dtor_cf29).dtor_cf29).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_388_v170).dtor_cf29).dtor_cf29).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_388_v170).dtor_cf29).dtor_cf29).dtor_cf26)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_388_v170).dtor_cf29).dtor_cf29).dtor_cf26)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_388_v170).dtor_cf29).dtor_cf29).dtor_cf27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_487_v227));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_488_v228)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_488_v228)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_488_v228)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_488_v228)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_573_v281).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_573_v281).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_573_v281).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[_dafny.ZERO]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[_dafny.ZERO]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[_dafny.ZERO]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[_dafny.ONE]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[_dafny.ONE]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[_dafny.ONE]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(2)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(2)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(2)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(3)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(3)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(3)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(4)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(4)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(4)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(5)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(5)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(5)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(6)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(6)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(6)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(7)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(7)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(7)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(8)]).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(8)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_682_v363)[new BigNumber(8)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_683_v365).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_685_v366).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(650),_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_686_i42));
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC5(cf10) {
      let $dt = new D1(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC3(cf8) {
      let $dt = new D1(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC6(cf11) {
      let $dt = new D1(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + this.cf8.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf10 === other.cf10;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(null);
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
    static create_DC8() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC9(cf13, cf14, cf15) {
      let $dt = new D2(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D2(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D2(3);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8";
      } else if (this.$tag === 1) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC10" + "(" + _dafny.toString(this.cf16) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf12 === other.cf12;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8();
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
    static create_DC12(cf18, cf19) {
      let $dt = new D3(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D3(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC13" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC12([], false);
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
    static create_DC15() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC14(cf21) {
      let $dt = new D4(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC16(cf22) {
      let $dt = new D4(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC15";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf22) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC15();
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
    static create_DC18(cf24, cf25, cf26, cf27) {
      let $dt = new D5(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC19(cf28) {
      let $dt = new D5(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC17(cf23) {
      let $dt = new D5(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC20(cf29) {
      let $dt = new D5(3);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25 && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC18(_dafny.ZERO, false, [], false);
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
    static create_DC22() {
      let $dt = new D6(0);
      return $dt;
    }
    static create_DC23(cf31, cf32, cf33) {
      let $dt = new D6(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC21(cf30) {
      let $dt = new D6(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC24(cf34) {
      let $dt = new D6(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC22";
      } else if (this.$tag === 1) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC24" + "(" + _dafny.toString(this.cf34) + ")";
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
        return other.$tag === 1 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC22();
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
    static create_DC25(cf35) {
      let $dt = new D7(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC25" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf36) {
      let $dt = new D8(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36);
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

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf38, cf39) {
      let $dt = new D9(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC29(cf40, cf41) {
      let $dt = new D9(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC27(cf37) {
      let $dt = new D9(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC30(cf42) {
      let $dt = new D9(3);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get is_DC30() { return this.$tag === 3; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC29" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC30" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC28(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f2 = _dafny.ZERO;
      this.f3 = false;
      this.f8 = _dafny.ZERO;
      this.f12 = _dafny.ZERO;
      this._f1 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.Seq.UnicodeFromString("");
      this._f11 = false;
      this._f13 = false;
      this._f14 = _dafny.ZERO;
      this._f15 = false;
      this._f16 = _dafny.ZERO;
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
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
    get f11() {
      let _this = this;
      return _this._f11;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f19 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f19) {
      let _this = this;
      (_this).f19 = f19;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return _this.f19;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(_this.f19)).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length))).length));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
